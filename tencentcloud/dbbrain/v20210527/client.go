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

package v20210527

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-05-27"

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


func NewAddUserContactRequest() (request *AddUserContactRequest) {
    request = &AddUserContactRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "AddUserContact")
    
    
    return
}

func NewAddUserContactResponse() (response *AddUserContactResponse) {
    response = &AddUserContactResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// AddUserContact
// 添加邮件接收联系人的姓名， 邮件地址，返回值为添加成功的联系人id。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddUserContact(request *AddUserContactRequest) (response *AddUserContactResponse, err error) {
    return c.AddUserContactWithContext(context.Background(), request)
}

// AddUserContact
// 添加邮件接收联系人的姓名， 邮件地址，返回值为添加成功的联系人id。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddUserContactWithContext(ctx context.Context, request *AddUserContactRequest) (response *AddUserContactResponse, err error) {
    if request == nil {
        request = NewAddUserContactRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "AddUserContact")
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddUserContact require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddUserContactResponse()
    err = c.Send(request, response)
    return
}

func NewCancelDBAutonomyActionRequest() (request *CancelDBAutonomyActionRequest) {
    request = &CancelDBAutonomyActionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CancelDBAutonomyAction")
    
    
    return
}

func NewCancelDBAutonomyActionResponse() (response *CancelDBAutonomyActionResponse) {
    response = &CancelDBAutonomyActionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelDBAutonomyAction
// 自治中心-终止自治任务（单次）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelDBAutonomyAction(request *CancelDBAutonomyActionRequest) (response *CancelDBAutonomyActionResponse, err error) {
    return c.CancelDBAutonomyActionWithContext(context.Background(), request)
}

// CancelDBAutonomyAction
// 自治中心-终止自治任务（单次）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelDBAutonomyActionWithContext(ctx context.Context, request *CancelDBAutonomyActionRequest) (response *CancelDBAutonomyActionResponse, err error) {
    if request == nil {
        request = NewCancelDBAutonomyActionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "CancelDBAutonomyAction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelDBAutonomyAction require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelDBAutonomyActionResponse()
    err = c.Send(request, response)
    return
}

func NewCancelDBAutonomyEventRequest() (request *CancelDBAutonomyEventRequest) {
    request = &CancelDBAutonomyEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CancelDBAutonomyEvent")
    
    
    return
}

func NewCancelDBAutonomyEventResponse() (response *CancelDBAutonomyEventResponse) {
    response = &CancelDBAutonomyEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelDBAutonomyEvent
// 自治中心-终止自治事件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelDBAutonomyEvent(request *CancelDBAutonomyEventRequest) (response *CancelDBAutonomyEventResponse, err error) {
    return c.CancelDBAutonomyEventWithContext(context.Background(), request)
}

// CancelDBAutonomyEvent
// 自治中心-终止自治事件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelDBAutonomyEventWithContext(ctx context.Context, request *CancelDBAutonomyEventRequest) (response *CancelDBAutonomyEventResponse, err error) {
    if request == nil {
        request = NewCancelDBAutonomyEventRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "CancelDBAutonomyEvent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelDBAutonomyEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelDBAutonomyEventResponse()
    err = c.Send(request, response)
    return
}

func NewCancelKillTaskRequest() (request *CancelKillTaskRequest) {
    request = &CancelKillTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CancelKillTask")
    
    
    return
}

func NewCancelKillTaskResponse() (response *CancelKillTaskResponse) {
    response = &CancelKillTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelKillTask
// 终止中断会话任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelKillTask(request *CancelKillTaskRequest) (response *CancelKillTaskResponse, err error) {
    return c.CancelKillTaskWithContext(context.Background(), request)
}

// CancelKillTask
// 终止中断会话任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelKillTaskWithContext(ctx context.Context, request *CancelKillTaskRequest) (response *CancelKillTaskResponse, err error) {
    if request == nil {
        request = NewCancelKillTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "CancelKillTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelKillTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelKillTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCancelRedisBigKeyAnalysisTasksRequest() (request *CancelRedisBigKeyAnalysisTasksRequest) {
    request = &CancelRedisBigKeyAnalysisTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CancelRedisBigKeyAnalysisTasks")
    
    
    return
}

func NewCancelRedisBigKeyAnalysisTasksResponse() (response *CancelRedisBigKeyAnalysisTasksResponse) {
    response = &CancelRedisBigKeyAnalysisTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CancelRedisBigKeyAnalysisTasks
// 自治中心-终止自治任务（单次）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelRedisBigKeyAnalysisTasks(request *CancelRedisBigKeyAnalysisTasksRequest) (response *CancelRedisBigKeyAnalysisTasksResponse, err error) {
    return c.CancelRedisBigKeyAnalysisTasksWithContext(context.Background(), request)
}

// CancelRedisBigKeyAnalysisTasks
// 自治中心-终止自治任务（单次）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CancelRedisBigKeyAnalysisTasksWithContext(ctx context.Context, request *CancelRedisBigKeyAnalysisTasksRequest) (response *CancelRedisBigKeyAnalysisTasksResponse, err error) {
    if request == nil {
        request = NewCancelRedisBigKeyAnalysisTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "CancelRedisBigKeyAnalysisTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelRedisBigKeyAnalysisTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelRedisBigKeyAnalysisTasksResponse()
    err = c.Send(request, response)
    return
}

func NewCloseAuditServiceRequest() (request *CloseAuditServiceRequest) {
    request = &CloseAuditServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CloseAuditService")
    
    
    return
}

func NewCloseAuditServiceResponse() (response *CloseAuditServiceResponse) {
    response = &CloseAuditServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CloseAuditService
// 不用审计日志时，关闭数据库审计
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CloseAuditService(request *CloseAuditServiceRequest) (response *CloseAuditServiceResponse, err error) {
    return c.CloseAuditServiceWithContext(context.Background(), request)
}

// CloseAuditService
// 不用审计日志时，关闭数据库审计
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CloseAuditServiceWithContext(ctx context.Context, request *CloseAuditServiceRequest) (response *CloseAuditServiceResponse, err error) {
    if request == nil {
        request = NewCloseAuditServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "CloseAuditService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseAuditService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseAuditServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAuditLogFileRequest() (request *CreateAuditLogFileRequest) {
    request = &CreateAuditLogFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CreateAuditLogFile")
    
    
    return
}

func NewCreateAuditLogFileResponse() (response *CreateAuditLogFileResponse) {
    response = &CreateAuditLogFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAuditLogFile
// 用于创建云数据库实例的审计日志文件，最多下载600w审计日志。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUDITNOTOPENED = "FailedOperation.AuditNotOpened"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_HASDUPLICATEDTASK = "UnsupportedOperation.HasDuplicatedTask"
func (c *Client) CreateAuditLogFile(request *CreateAuditLogFileRequest) (response *CreateAuditLogFileResponse, err error) {
    return c.CreateAuditLogFileWithContext(context.Background(), request)
}

// CreateAuditLogFile
// 用于创建云数据库实例的审计日志文件，最多下载600w审计日志。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUDITNOTOPENED = "FailedOperation.AuditNotOpened"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_HASDUPLICATEDTASK = "UnsupportedOperation.HasDuplicatedTask"
func (c *Client) CreateAuditLogFileWithContext(ctx context.Context, request *CreateAuditLogFileRequest) (response *CreateAuditLogFileResponse, err error) {
    if request == nil {
        request = NewCreateAuditLogFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "CreateAuditLogFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAuditLogFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAuditLogFileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBDiagReportTaskRequest() (request *CreateDBDiagReportTaskRequest) {
    request = &CreateDBDiagReportTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CreateDBDiagReportTask")
    
    
    return
}

func NewCreateDBDiagReportTaskResponse() (response *CreateDBDiagReportTaskResponse) {
    response = &CreateDBDiagReportTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDBDiagReportTask
// 创建健康报告，并可以选择是否发送邮件。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDBDiagReportTask(request *CreateDBDiagReportTaskRequest) (response *CreateDBDiagReportTaskResponse, err error) {
    return c.CreateDBDiagReportTaskWithContext(context.Background(), request)
}

// CreateDBDiagReportTask
// 创建健康报告，并可以选择是否发送邮件。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDBDiagReportTaskWithContext(ctx context.Context, request *CreateDBDiagReportTaskRequest) (response *CreateDBDiagReportTaskResponse, err error) {
    if request == nil {
        request = NewCreateDBDiagReportTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "CreateDBDiagReportTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBDiagReportTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBDiagReportTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDBDiagReportUrlRequest() (request *CreateDBDiagReportUrlRequest) {
    request = &CreateDBDiagReportUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CreateDBDiagReportUrl")
    
    
    return
}

func NewCreateDBDiagReportUrlResponse() (response *CreateDBDiagReportUrlResponse) {
    response = &CreateDBDiagReportUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateDBDiagReportUrl
// 创建健康报告的浏览地址。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDBDiagReportUrl(request *CreateDBDiagReportUrlRequest) (response *CreateDBDiagReportUrlResponse, err error) {
    return c.CreateDBDiagReportUrlWithContext(context.Background(), request)
}

// CreateDBDiagReportUrl
// 创建健康报告的浏览地址。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateDBDiagReportUrlWithContext(ctx context.Context, request *CreateDBDiagReportUrlRequest) (response *CreateDBDiagReportUrlResponse, err error) {
    if request == nil {
        request = NewCreateDBDiagReportUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "CreateDBDiagReportUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDBDiagReportUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDBDiagReportUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateKillTaskRequest() (request *CreateKillTaskRequest) {
    request = &CreateKillTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CreateKillTask")
    
    
    return
}

func NewCreateKillTaskResponse() (response *CreateKillTaskResponse) {
    response = &CreateKillTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateKillTask
// 创建中断会话的任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateKillTask(request *CreateKillTaskRequest) (response *CreateKillTaskResponse, err error) {
    return c.CreateKillTaskWithContext(context.Background(), request)
}

// CreateKillTask
// 创建中断会话的任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateKillTaskWithContext(ctx context.Context, request *CreateKillTaskRequest) (response *CreateKillTaskResponse, err error) {
    if request == nil {
        request = NewCreateKillTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "CreateKillTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateKillTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateKillTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMailProfileRequest() (request *CreateMailProfileRequest) {
    request = &CreateMailProfileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CreateMailProfile")
    
    
    return
}

func NewCreateMailProfileResponse() (response *CreateMailProfileResponse) {
    response = &CreateMailProfileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMailProfile
// 创建邮件配置。其中入参ProfileType表示所创建配置的类型，ProfileType 取值包括：dbScan_mail_configuration（数据库巡检邮件配置）、scheduler_mail_configuration（定期生成健康报告的邮件发送配置）。Region统一选择广州，和实例所属地域无关。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateMailProfile(request *CreateMailProfileRequest) (response *CreateMailProfileResponse, err error) {
    return c.CreateMailProfileWithContext(context.Background(), request)
}

// CreateMailProfile
// 创建邮件配置。其中入参ProfileType表示所创建配置的类型，ProfileType 取值包括：dbScan_mail_configuration（数据库巡检邮件配置）、scheduler_mail_configuration（定期生成健康报告的邮件发送配置）。Region统一选择广州，和实例所属地域无关。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateMailProfileWithContext(ctx context.Context, request *CreateMailProfileRequest) (response *CreateMailProfileResponse, err error) {
    if request == nil {
        request = NewCreateMailProfileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "CreateMailProfile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMailProfile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMailProfileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProxySessionKillTaskRequest() (request *CreateProxySessionKillTaskRequest) {
    request = &CreateProxySessionKillTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CreateProxySessionKillTask")
    
    
    return
}

func NewCreateProxySessionKillTaskResponse() (response *CreateProxySessionKillTaskResponse) {
    response = &CreateProxySessionKillTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateProxySessionKillTask
// 创建中止所有代理节点连接会话的异步任务。当前仅支持 Redis。得到的返回值为异步任务 id，可以作为参数传入接口 DescribeProxySessionKillTasks 查询kill会话任务执行状态。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateProxySessionKillTask(request *CreateProxySessionKillTaskRequest) (response *CreateProxySessionKillTaskResponse, err error) {
    return c.CreateProxySessionKillTaskWithContext(context.Background(), request)
}

// CreateProxySessionKillTask
// 创建中止所有代理节点连接会话的异步任务。当前仅支持 Redis。得到的返回值为异步任务 id，可以作为参数传入接口 DescribeProxySessionKillTasks 查询kill会话任务执行状态。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateProxySessionKillTaskWithContext(ctx context.Context, request *CreateProxySessionKillTaskRequest) (response *CreateProxySessionKillTaskResponse, err error) {
    if request == nil {
        request = NewCreateProxySessionKillTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "CreateProxySessionKillTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProxySessionKillTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProxySessionKillTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRedisBigKeyAnalysisTaskRequest() (request *CreateRedisBigKeyAnalysisTaskRequest) {
    request = &CreateRedisBigKeyAnalysisTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CreateRedisBigKeyAnalysisTask")
    
    
    return
}

func NewCreateRedisBigKeyAnalysisTaskResponse() (response *CreateRedisBigKeyAnalysisTaskResponse) {
    response = &CreateRedisBigKeyAnalysisTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRedisBigKeyAnalysisTask
// 即时创建redis实例大key分析任务，限制正在运行的即时分析任务数量默认为5。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRedisBigKeyAnalysisTask(request *CreateRedisBigKeyAnalysisTaskRequest) (response *CreateRedisBigKeyAnalysisTaskResponse, err error) {
    return c.CreateRedisBigKeyAnalysisTaskWithContext(context.Background(), request)
}

// CreateRedisBigKeyAnalysisTask
// 即时创建redis实例大key分析任务，限制正在运行的即时分析任务数量默认为5。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateRedisBigKeyAnalysisTaskWithContext(ctx context.Context, request *CreateRedisBigKeyAnalysisTaskRequest) (response *CreateRedisBigKeyAnalysisTaskResponse, err error) {
    if request == nil {
        request = NewCreateRedisBigKeyAnalysisTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "CreateRedisBigKeyAnalysisTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRedisBigKeyAnalysisTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRedisBigKeyAnalysisTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSchedulerMailProfileRequest() (request *CreateSchedulerMailProfileRequest) {
    request = &CreateSchedulerMailProfileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CreateSchedulerMailProfile")
    
    
    return
}

func NewCreateSchedulerMailProfileResponse() (response *CreateSchedulerMailProfileResponse) {
    response = &CreateSchedulerMailProfileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSchedulerMailProfile
// 该接口用于创建定期生成健康报告并邮件发送的配置，将健康报告的定期生成时间作为参数传入（周一至周日），用于设置健康报告的定期生成时间，同时保存相应的定期邮件发送的配置。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSchedulerMailProfile(request *CreateSchedulerMailProfileRequest) (response *CreateSchedulerMailProfileResponse, err error) {
    return c.CreateSchedulerMailProfileWithContext(context.Background(), request)
}

// CreateSchedulerMailProfile
// 该接口用于创建定期生成健康报告并邮件发送的配置，将健康报告的定期生成时间作为参数传入（周一至周日），用于设置健康报告的定期生成时间，同时保存相应的定期邮件发送的配置。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSchedulerMailProfileWithContext(ctx context.Context, request *CreateSchedulerMailProfileRequest) (response *CreateSchedulerMailProfileResponse, err error) {
    if request == nil {
        request = NewCreateSchedulerMailProfileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "CreateSchedulerMailProfile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSchedulerMailProfile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSchedulerMailProfileResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSecurityAuditLogExportTaskRequest() (request *CreateSecurityAuditLogExportTaskRequest) {
    request = &CreateSecurityAuditLogExportTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CreateSecurityAuditLogExportTask")
    
    
    return
}

func NewCreateSecurityAuditLogExportTaskResponse() (response *CreateSecurityAuditLogExportTaskResponse) {
    response = &CreateSecurityAuditLogExportTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSecurityAuditLogExportTask
// 创建安全审计日志导出任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSecurityAuditLogExportTask(request *CreateSecurityAuditLogExportTaskRequest) (response *CreateSecurityAuditLogExportTaskResponse, err error) {
    return c.CreateSecurityAuditLogExportTaskWithContext(context.Background(), request)
}

// CreateSecurityAuditLogExportTask
// 创建安全审计日志导出任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSecurityAuditLogExportTaskWithContext(ctx context.Context, request *CreateSecurityAuditLogExportTaskRequest) (response *CreateSecurityAuditLogExportTaskResponse, err error) {
    if request == nil {
        request = NewCreateSecurityAuditLogExportTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "CreateSecurityAuditLogExportTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSecurityAuditLogExportTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSecurityAuditLogExportTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSqlFilterRequest() (request *CreateSqlFilterRequest) {
    request = &CreateSqlFilterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CreateSqlFilter")
    
    
    return
}

func NewCreateSqlFilterResponse() (response *CreateSqlFilterResponse) {
    response = &CreateSqlFilterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSqlFilter
// 创建实例SQL限流任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSqlFilter(request *CreateSqlFilterRequest) (response *CreateSqlFilterResponse, err error) {
    return c.CreateSqlFilterWithContext(context.Background(), request)
}

// CreateSqlFilter
// 创建实例SQL限流任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSqlFilterWithContext(ctx context.Context, request *CreateSqlFilterRequest) (response *CreateSqlFilterResponse, err error) {
    if request == nil {
        request = NewCreateSqlFilterRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "CreateSqlFilter")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSqlFilter require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSqlFilterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserAutonomyProfileRequest() (request *CreateUserAutonomyProfileRequest) {
    request = &CreateUserAutonomyProfileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "CreateUserAutonomyProfile")
    
    
    return
}

func NewCreateUserAutonomyProfileResponse() (response *CreateUserAutonomyProfileResponse) {
    response = &CreateUserAutonomyProfileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateUserAutonomyProfile
// 自治中心-终止自治任务（单次）；注意：接口需要加白名单。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUserAutonomyProfile(request *CreateUserAutonomyProfileRequest) (response *CreateUserAutonomyProfileResponse, err error) {
    return c.CreateUserAutonomyProfileWithContext(context.Background(), request)
}

// CreateUserAutonomyProfile
// 自治中心-终止自治任务（单次）；注意：接口需要加白名单。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateUserAutonomyProfileWithContext(ctx context.Context, request *CreateUserAutonomyProfileRequest) (response *CreateUserAutonomyProfileResponse, err error) {
    if request == nil {
        request = NewCreateUserAutonomyProfileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "CreateUserAutonomyProfile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUserAutonomyProfile require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserAutonomyProfileResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuditLogFileRequest() (request *DeleteAuditLogFileRequest) {
    request = &DeleteAuditLogFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DeleteAuditLogFile")
    
    
    return
}

func NewDeleteAuditLogFileResponse() (response *DeleteAuditLogFileResponse) {
    response = &DeleteAuditLogFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAuditLogFile
// 用于删除云数据库实例的审计日志文件。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAuditLogFile(request *DeleteAuditLogFileRequest) (response *DeleteAuditLogFileResponse, err error) {
    return c.DeleteAuditLogFileWithContext(context.Background(), request)
}

// DeleteAuditLogFile
// 用于删除云数据库实例的审计日志文件。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAuditLogFileWithContext(ctx context.Context, request *DeleteAuditLogFileRequest) (response *DeleteAuditLogFileResponse, err error) {
    if request == nil {
        request = NewDeleteAuditLogFileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DeleteAuditLogFile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAuditLogFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAuditLogFileResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDBDiagReportTasksRequest() (request *DeleteDBDiagReportTasksRequest) {
    request = &DeleteDBDiagReportTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DeleteDBDiagReportTasks")
    
    
    return
}

func NewDeleteDBDiagReportTasksResponse() (response *DeleteDBDiagReportTasksResponse) {
    response = &DeleteDBDiagReportTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteDBDiagReportTasks
// 根据任务id删除健康报告生成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDBDiagReportTasks(request *DeleteDBDiagReportTasksRequest) (response *DeleteDBDiagReportTasksResponse, err error) {
    return c.DeleteDBDiagReportTasksWithContext(context.Background(), request)
}

// DeleteDBDiagReportTasks
// 根据任务id删除健康报告生成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteDBDiagReportTasksWithContext(ctx context.Context, request *DeleteDBDiagReportTasksRequest) (response *DeleteDBDiagReportTasksResponse, err error) {
    if request == nil {
        request = NewDeleteDBDiagReportTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DeleteDBDiagReportTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDBDiagReportTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDBDiagReportTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRedisBigKeyAnalysisTasksRequest() (request *DeleteRedisBigKeyAnalysisTasksRequest) {
    request = &DeleteRedisBigKeyAnalysisTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DeleteRedisBigKeyAnalysisTasks")
    
    
    return
}

func NewDeleteRedisBigKeyAnalysisTasksResponse() (response *DeleteRedisBigKeyAnalysisTasksResponse) {
    response = &DeleteRedisBigKeyAnalysisTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteRedisBigKeyAnalysisTasks
// 删除Redis实例的大key分析任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRedisBigKeyAnalysisTasks(request *DeleteRedisBigKeyAnalysisTasksRequest) (response *DeleteRedisBigKeyAnalysisTasksResponse, err error) {
    return c.DeleteRedisBigKeyAnalysisTasksWithContext(context.Background(), request)
}

// DeleteRedisBigKeyAnalysisTasks
// 删除Redis实例的大key分析任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteRedisBigKeyAnalysisTasksWithContext(ctx context.Context, request *DeleteRedisBigKeyAnalysisTasksRequest) (response *DeleteRedisBigKeyAnalysisTasksResponse, err error) {
    if request == nil {
        request = NewDeleteRedisBigKeyAnalysisTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DeleteRedisBigKeyAnalysisTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRedisBigKeyAnalysisTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRedisBigKeyAnalysisTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSecurityAuditLogExportTasksRequest() (request *DeleteSecurityAuditLogExportTasksRequest) {
    request = &DeleteSecurityAuditLogExportTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DeleteSecurityAuditLogExportTasks")
    
    
    return
}

func NewDeleteSecurityAuditLogExportTasksResponse() (response *DeleteSecurityAuditLogExportTasksResponse) {
    response = &DeleteSecurityAuditLogExportTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSecurityAuditLogExportTasks
// 删除安全审计日志导出任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSecurityAuditLogExportTasks(request *DeleteSecurityAuditLogExportTasksRequest) (response *DeleteSecurityAuditLogExportTasksResponse, err error) {
    return c.DeleteSecurityAuditLogExportTasksWithContext(context.Background(), request)
}

// DeleteSecurityAuditLogExportTasks
// 删除安全审计日志导出任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSecurityAuditLogExportTasksWithContext(ctx context.Context, request *DeleteSecurityAuditLogExportTasksRequest) (response *DeleteSecurityAuditLogExportTasksResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityAuditLogExportTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DeleteSecurityAuditLogExportTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSecurityAuditLogExportTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSecurityAuditLogExportTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSqlFiltersRequest() (request *DeleteSqlFiltersRequest) {
    request = &DeleteSqlFiltersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DeleteSqlFilters")
    
    
    return
}

func NewDeleteSqlFiltersResponse() (response *DeleteSqlFiltersResponse) {
    response = &DeleteSqlFiltersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteSqlFilters
// 删除实例SQL限流任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSqlFilters(request *DeleteSqlFiltersRequest) (response *DeleteSqlFiltersResponse, err error) {
    return c.DeleteSqlFiltersWithContext(context.Background(), request)
}

// DeleteSqlFilters
// 删除实例SQL限流任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteSqlFiltersWithContext(ctx context.Context, request *DeleteSqlFiltersRequest) (response *DeleteSqlFiltersResponse, err error) {
    if request == nil {
        request = NewDeleteSqlFiltersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DeleteSqlFilters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSqlFilters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSqlFiltersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAlarmTemplateRequest() (request *DescribeAlarmTemplateRequest) {
    request = &DescribeAlarmTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeAlarmTemplate")
    
    
    return
}

func NewDescribeAlarmTemplateResponse() (response *DescribeAlarmTemplateResponse) {
    response = &DescribeAlarmTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAlarmTemplate
// 通知模板查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAlarmTemplate(request *DescribeAlarmTemplateRequest) (response *DescribeAlarmTemplateResponse, err error) {
    return c.DescribeAlarmTemplateWithContext(context.Background(), request)
}

// DescribeAlarmTemplate
// 通知模板查询
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAlarmTemplateWithContext(ctx context.Context, request *DescribeAlarmTemplateRequest) (response *DescribeAlarmTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeAlarmTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeAlarmTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAlarmTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAlarmTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllUserContactRequest() (request *DescribeAllUserContactRequest) {
    request = &DescribeAllUserContactRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeAllUserContact")
    
    
    return
}

func NewDescribeAllUserContactResponse() (response *DescribeAllUserContactResponse) {
    response = &DescribeAllUserContactResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllUserContact
// 获取邮件发送中联系人的相关信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAllUserContact(request *DescribeAllUserContactRequest) (response *DescribeAllUserContactResponse, err error) {
    return c.DescribeAllUserContactWithContext(context.Background(), request)
}

// DescribeAllUserContact
// 获取邮件发送中联系人的相关信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAllUserContactWithContext(ctx context.Context, request *DescribeAllUserContactRequest) (response *DescribeAllUserContactResponse, err error) {
    if request == nil {
        request = NewDescribeAllUserContactRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeAllUserContact")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllUserContact require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllUserContactResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAllUserGroupRequest() (request *DescribeAllUserGroupRequest) {
    request = &DescribeAllUserGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeAllUserGroup")
    
    
    return
}

func NewDescribeAllUserGroupResponse() (response *DescribeAllUserGroupResponse) {
    response = &DescribeAllUserGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAllUserGroup
// 获取邮件发送联系组的相关信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAllUserGroup(request *DescribeAllUserGroupRequest) (response *DescribeAllUserGroupResponse, err error) {
    return c.DescribeAllUserGroupWithContext(context.Background(), request)
}

// DescribeAllUserGroup
// 获取邮件发送联系组的相关信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAllUserGroupWithContext(ctx context.Context, request *DescribeAllUserGroupRequest) (response *DescribeAllUserGroupResponse, err error) {
    if request == nil {
        request = NewDescribeAllUserGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeAllUserGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllUserGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditInstanceListRequest() (request *DescribeAuditInstanceListRequest) {
    request = &DescribeAuditInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeAuditInstanceList")
    
    
    return
}

func NewDescribeAuditInstanceListResponse() (response *DescribeAuditInstanceListResponse) {
    response = &DescribeAuditInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditInstanceList
// 查询实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAuditInstanceList(request *DescribeAuditInstanceListRequest) (response *DescribeAuditInstanceListResponse, err error) {
    return c.DescribeAuditInstanceListWithContext(context.Background(), request)
}

// DescribeAuditInstanceList
// 查询实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAuditInstanceListWithContext(ctx context.Context, request *DescribeAuditInstanceListRequest) (response *DescribeAuditInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeAuditInstanceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeAuditInstanceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditInstanceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditLogFilesRequest() (request *DescribeAuditLogFilesRequest) {
    request = &DescribeAuditLogFilesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeAuditLogFiles")
    
    
    return
}

func NewDescribeAuditLogFilesResponse() (response *DescribeAuditLogFilesResponse) {
    response = &DescribeAuditLogFilesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAuditLogFiles
// 用于创建云数据库实例的审计日志文件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUDITNOTOPENED = "FailedOperation.AuditNotOpened"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAuditLogFiles(request *DescribeAuditLogFilesRequest) (response *DescribeAuditLogFilesResponse, err error) {
    return c.DescribeAuditLogFilesWithContext(context.Background(), request)
}

// DescribeAuditLogFiles
// 用于创建云数据库实例的审计日志文件
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AUDITNOTOPENED = "FailedOperation.AuditNotOpened"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAuditLogFilesWithContext(ctx context.Context, request *DescribeAuditLogFilesRequest) (response *DescribeAuditLogFilesResponse, err error) {
    if request == nil {
        request = NewDescribeAuditLogFilesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeAuditLogFiles")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuditLogFiles require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuditLogFilesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBAutonomyActionRequest() (request *DescribeDBAutonomyActionRequest) {
    request = &DescribeDBAutonomyActionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeDBAutonomyAction")
    
    
    return
}

func NewDescribeDBAutonomyActionResponse() (response *DescribeDBAutonomyActionResponse) {
    response = &DescribeDBAutonomyActionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBAutonomyAction
// 自治中心-查询自治事件任务详情。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBAutonomyAction(request *DescribeDBAutonomyActionRequest) (response *DescribeDBAutonomyActionResponse, err error) {
    return c.DescribeDBAutonomyActionWithContext(context.Background(), request)
}

// DescribeDBAutonomyAction
// 自治中心-查询自治事件任务详情。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBAutonomyActionWithContext(ctx context.Context, request *DescribeDBAutonomyActionRequest) (response *DescribeDBAutonomyActionResponse, err error) {
    if request == nil {
        request = NewDescribeDBAutonomyActionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeDBAutonomyAction")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBAutonomyAction require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBAutonomyActionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBAutonomyActionsRequest() (request *DescribeDBAutonomyActionsRequest) {
    request = &DescribeDBAutonomyActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeDBAutonomyActions")
    
    
    return
}

func NewDescribeDBAutonomyActionsResponse() (response *DescribeDBAutonomyActionsResponse) {
    response = &DescribeDBAutonomyActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBAutonomyActions
// 自治中心-终止自治任务（单次）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBAutonomyActions(request *DescribeDBAutonomyActionsRequest) (response *DescribeDBAutonomyActionsResponse, err error) {
    return c.DescribeDBAutonomyActionsWithContext(context.Background(), request)
}

// DescribeDBAutonomyActions
// 自治中心-终止自治任务（单次）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBAutonomyActionsWithContext(ctx context.Context, request *DescribeDBAutonomyActionsRequest) (response *DescribeDBAutonomyActionsResponse, err error) {
    if request == nil {
        request = NewDescribeDBAutonomyActionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeDBAutonomyActions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBAutonomyActions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBAutonomyActionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBAutonomyEventsRequest() (request *DescribeDBAutonomyEventsRequest) {
    request = &DescribeDBAutonomyEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeDBAutonomyEvents")
    
    
    return
}

func NewDescribeDBAutonomyEventsResponse() (response *DescribeDBAutonomyEventsResponse) {
    response = &DescribeDBAutonomyEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBAutonomyEvents
// 自治中心-终止自治任务（单次）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBAutonomyEvents(request *DescribeDBAutonomyEventsRequest) (response *DescribeDBAutonomyEventsResponse, err error) {
    return c.DescribeDBAutonomyEventsWithContext(context.Background(), request)
}

// DescribeDBAutonomyEvents
// 自治中心-终止自治任务（单次）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBAutonomyEventsWithContext(ctx context.Context, request *DescribeDBAutonomyEventsRequest) (response *DescribeDBAutonomyEventsResponse, err error) {
    if request == nil {
        request = NewDescribeDBAutonomyEventsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeDBAutonomyEvents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBAutonomyEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBAutonomyEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBDiagEventRequest() (request *DescribeDBDiagEventRequest) {
    request = &DescribeDBDiagEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeDBDiagEvent")
    
    
    return
}

func NewDescribeDBDiagEventResponse() (response *DescribeDBDiagEventResponse) {
    response = &DescribeDBDiagEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBDiagEvent
// 获取实例异常诊断事件的详情信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBDiagEvent(request *DescribeDBDiagEventRequest) (response *DescribeDBDiagEventResponse, err error) {
    return c.DescribeDBDiagEventWithContext(context.Background(), request)
}

// DescribeDBDiagEvent
// 获取实例异常诊断事件的详情信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBDiagEventWithContext(ctx context.Context, request *DescribeDBDiagEventRequest) (response *DescribeDBDiagEventResponse, err error) {
    if request == nil {
        request = NewDescribeDBDiagEventRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeDBDiagEvent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBDiagEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBDiagEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBDiagEventsRequest() (request *DescribeDBDiagEventsRequest) {
    request = &DescribeDBDiagEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeDBDiagEvents")
    
    
    return
}

func NewDescribeDBDiagEventsResponse() (response *DescribeDBDiagEventsResponse) {
    response = &DescribeDBDiagEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBDiagEvents
// 获取指定时间段内的诊断事件列表，支持依据风险等级、实例ID等条件过滤。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBDiagEvents(request *DescribeDBDiagEventsRequest) (response *DescribeDBDiagEventsResponse, err error) {
    return c.DescribeDBDiagEventsWithContext(context.Background(), request)
}

// DescribeDBDiagEvents
// 获取指定时间段内的诊断事件列表，支持依据风险等级、实例ID等条件过滤。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBDiagEventsWithContext(ctx context.Context, request *DescribeDBDiagEventsRequest) (response *DescribeDBDiagEventsResponse, err error) {
    if request == nil {
        request = NewDescribeDBDiagEventsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeDBDiagEvents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBDiagEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBDiagEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBDiagHistoryRequest() (request *DescribeDBDiagHistoryRequest) {
    request = &DescribeDBDiagHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeDBDiagHistory")
    
    
    return
}

func NewDescribeDBDiagHistoryResponse() (response *DescribeDBDiagHistoryResponse) {
    response = &DescribeDBDiagHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBDiagHistory
// 获取实例诊断事件的列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBDiagHistory(request *DescribeDBDiagHistoryRequest) (response *DescribeDBDiagHistoryResponse, err error) {
    return c.DescribeDBDiagHistoryWithContext(context.Background(), request)
}

// DescribeDBDiagHistory
// 获取实例诊断事件的列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBDiagHistoryWithContext(ctx context.Context, request *DescribeDBDiagHistoryRequest) (response *DescribeDBDiagHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeDBDiagHistoryRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeDBDiagHistory")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBDiagHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBDiagHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBDiagReportContentRequest() (request *DescribeDBDiagReportContentRequest) {
    request = &DescribeDBDiagReportContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeDBDiagReportContent")
    
    
    return
}

func NewDescribeDBDiagReportContentResponse() (response *DescribeDBDiagReportContentResponse) {
    response = &DescribeDBDiagReportContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBDiagReportContent
// 健康报告内容。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDBDiagReportContent(request *DescribeDBDiagReportContentRequest) (response *DescribeDBDiagReportContentResponse, err error) {
    return c.DescribeDBDiagReportContentWithContext(context.Background(), request)
}

// DescribeDBDiagReportContent
// 健康报告内容。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDBDiagReportContentWithContext(ctx context.Context, request *DescribeDBDiagReportContentRequest) (response *DescribeDBDiagReportContentResponse, err error) {
    if request == nil {
        request = NewDescribeDBDiagReportContentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeDBDiagReportContent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBDiagReportContent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBDiagReportContentResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBDiagReportTasksRequest() (request *DescribeDBDiagReportTasksRequest) {
    request = &DescribeDBDiagReportTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeDBDiagReportTasks")
    
    
    return
}

func NewDescribeDBDiagReportTasksResponse() (response *DescribeDBDiagReportTasksResponse) {
    response = &DescribeDBDiagReportTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBDiagReportTasks
// 查询健康报告生成任务列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBDiagReportTasks(request *DescribeDBDiagReportTasksRequest) (response *DescribeDBDiagReportTasksResponse, err error) {
    return c.DescribeDBDiagReportTasksWithContext(context.Background(), request)
}

// DescribeDBDiagReportTasks
// 查询健康报告生成任务列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBDiagReportTasksWithContext(ctx context.Context, request *DescribeDBDiagReportTasksRequest) (response *DescribeDBDiagReportTasksResponse, err error) {
    if request == nil {
        request = NewDescribeDBDiagReportTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeDBDiagReportTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBDiagReportTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBDiagReportTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBPerfTimeSeriesRequest() (request *DescribeDBPerfTimeSeriesRequest) {
    request = &DescribeDBPerfTimeSeriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeDBPerfTimeSeries")
    
    
    return
}

func NewDescribeDBPerfTimeSeriesResponse() (response *DescribeDBPerfTimeSeriesResponse) {
    response = &DescribeDBPerfTimeSeriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBPerfTimeSeries
// 根据实例ID获取指定时间段的性能趋势。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBPerfTimeSeries(request *DescribeDBPerfTimeSeriesRequest) (response *DescribeDBPerfTimeSeriesResponse, err error) {
    return c.DescribeDBPerfTimeSeriesWithContext(context.Background(), request)
}

// DescribeDBPerfTimeSeries
// 根据实例ID获取指定时间段的性能趋势。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBPerfTimeSeriesWithContext(ctx context.Context, request *DescribeDBPerfTimeSeriesRequest) (response *DescribeDBPerfTimeSeriesResponse, err error) {
    if request == nil {
        request = NewDescribeDBPerfTimeSeriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeDBPerfTimeSeries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBPerfTimeSeries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBPerfTimeSeriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDBSpaceStatusRequest() (request *DescribeDBSpaceStatusRequest) {
    request = &DescribeDBSpaceStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeDBSpaceStatus")
    
    
    return
}

func NewDescribeDBSpaceStatusResponse() (response *DescribeDBSpaceStatusResponse) {
    response = &DescribeDBSpaceStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDBSpaceStatus
// 获取指定时间段内的实例空间使用概览，包括磁盘增长量(MB)、磁盘剩余(MB)、磁盘总量(MB)及预计可用天数。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEGMENTLOADING = "FailedOperation.SegmentLoading"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBSpaceStatus(request *DescribeDBSpaceStatusRequest) (response *DescribeDBSpaceStatusResponse, err error) {
    return c.DescribeDBSpaceStatusWithContext(context.Background(), request)
}

// DescribeDBSpaceStatus
// 获取指定时间段内的实例空间使用概览，包括磁盘增长量(MB)、磁盘剩余(MB)、磁盘总量(MB)及预计可用天数。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEGMENTLOADING = "FailedOperation.SegmentLoading"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDBSpaceStatusWithContext(ctx context.Context, request *DescribeDBSpaceStatusRequest) (response *DescribeDBSpaceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeDBSpaceStatusRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeDBSpaceStatus")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBSpaceStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBSpaceStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDiagDBInstancesRequest() (request *DescribeDiagDBInstancesRequest) {
    request = &DescribeDiagDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeDiagDBInstances")
    
    
    return
}

func NewDescribeDiagDBInstancesResponse() (response *DescribeDiagDBInstancesResponse) {
    response = &DescribeDiagDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDiagDBInstances
// 获取实例信息列表。Region统一选择广州。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDiagDBInstances(request *DescribeDiagDBInstancesRequest) (response *DescribeDiagDBInstancesResponse, err error) {
    return c.DescribeDiagDBInstancesWithContext(context.Background(), request)
}

// DescribeDiagDBInstances
// 获取实例信息列表。Region统一选择广州。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDiagDBInstancesWithContext(ctx context.Context, request *DescribeDiagDBInstancesRequest) (response *DescribeDiagDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDiagDBInstancesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeDiagDBInstances")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDiagDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDiagDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHealthScoreRequest() (request *DescribeHealthScoreRequest) {
    request = &DescribeHealthScoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeHealthScore")
    
    
    return
}

func NewDescribeHealthScoreResponse() (response *DescribeHealthScoreResponse) {
    response = &DescribeHealthScoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHealthScore
// 根据实例ID获取指定时间段（30分钟）的健康得分，以及异常扣分项。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeHealthScore(request *DescribeHealthScoreRequest) (response *DescribeHealthScoreResponse, err error) {
    return c.DescribeHealthScoreWithContext(context.Background(), request)
}

// DescribeHealthScore
// 根据实例ID获取指定时间段（30分钟）的健康得分，以及异常扣分项。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeHealthScoreWithContext(ctx context.Context, request *DescribeHealthScoreRequest) (response *DescribeHealthScoreResponse, err error) {
    if request == nil {
        request = NewDescribeHealthScoreRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeHealthScore")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHealthScore require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHealthScoreResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHealthScoreTimeSeriesRequest() (request *DescribeHealthScoreTimeSeriesRequest) {
    request = &DescribeHealthScoreTimeSeriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeHealthScoreTimeSeries")
    
    
    return
}

func NewDescribeHealthScoreTimeSeriesResponse() (response *DescribeHealthScoreTimeSeriesResponse) {
    response = &DescribeHealthScoreTimeSeriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHealthScoreTimeSeries
// 获取指定时间段内的健康得分趋势
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeHealthScoreTimeSeries(request *DescribeHealthScoreTimeSeriesRequest) (response *DescribeHealthScoreTimeSeriesResponse, err error) {
    return c.DescribeHealthScoreTimeSeriesWithContext(context.Background(), request)
}

// DescribeHealthScoreTimeSeries
// 获取指定时间段内的健康得分趋势
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeHealthScoreTimeSeriesWithContext(ctx context.Context, request *DescribeHealthScoreTimeSeriesRequest) (response *DescribeHealthScoreTimeSeriesResponse, err error) {
    if request == nil {
        request = NewDescribeHealthScoreTimeSeriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeHealthScoreTimeSeries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHealthScoreTimeSeries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHealthScoreTimeSeriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIndexRecommendAggregationSlowLogsRequest() (request *DescribeIndexRecommendAggregationSlowLogsRequest) {
    request = &DescribeIndexRecommendAggregationSlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeIndexRecommendAggregationSlowLogs")
    
    
    return
}

func NewDescribeIndexRecommendAggregationSlowLogsResponse() (response *DescribeIndexRecommendAggregationSlowLogsResponse) {
    response = &DescribeIndexRecommendAggregationSlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIndexRecommendAggregationSlowLogs
// 查询某张表的慢查模板概览，这个接口是对用户点击对应的推荐索引后，展示慢日志用的
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIndexRecommendAggregationSlowLogs(request *DescribeIndexRecommendAggregationSlowLogsRequest) (response *DescribeIndexRecommendAggregationSlowLogsResponse, err error) {
    return c.DescribeIndexRecommendAggregationSlowLogsWithContext(context.Background(), request)
}

// DescribeIndexRecommendAggregationSlowLogs
// 查询某张表的慢查模板概览，这个接口是对用户点击对应的推荐索引后，展示慢日志用的
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIndexRecommendAggregationSlowLogsWithContext(ctx context.Context, request *DescribeIndexRecommendAggregationSlowLogsRequest) (response *DescribeIndexRecommendAggregationSlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeIndexRecommendAggregationSlowLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeIndexRecommendAggregationSlowLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIndexRecommendAggregationSlowLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIndexRecommendAggregationSlowLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIndexRecommendInfoRequest() (request *DescribeIndexRecommendInfoRequest) {
    request = &DescribeIndexRecommendInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeIndexRecommendInfo")
    
    
    return
}

func NewDescribeIndexRecommendInfoResponse() (response *DescribeIndexRecommendInfoResponse) {
    response = &DescribeIndexRecommendInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIndexRecommendInfo
// 查询实例的索引推荐信息，包括索引统计相关信息，推荐索引列表，无效索引列表等。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIndexRecommendInfo(request *DescribeIndexRecommendInfoRequest) (response *DescribeIndexRecommendInfoResponse, err error) {
    return c.DescribeIndexRecommendInfoWithContext(context.Background(), request)
}

// DescribeIndexRecommendInfo
// 查询实例的索引推荐信息，包括索引统计相关信息，推荐索引列表，无效索引列表等。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIndexRecommendInfoWithContext(ctx context.Context, request *DescribeIndexRecommendInfoRequest) (response *DescribeIndexRecommendInfoResponse, err error) {
    if request == nil {
        request = NewDescribeIndexRecommendInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeIndexRecommendInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIndexRecommendInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIndexRecommendInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMailProfileRequest() (request *DescribeMailProfileRequest) {
    request = &DescribeMailProfileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeMailProfile")
    
    
    return
}

func NewDescribeMailProfileResponse() (response *DescribeMailProfileResponse) {
    response = &DescribeMailProfileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMailProfile
// 获取发送邮件的配置， 包括数据库巡检的邮件配置以及定期生成健康报告的邮件发送配置。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMailProfile(request *DescribeMailProfileRequest) (response *DescribeMailProfileResponse, err error) {
    return c.DescribeMailProfileWithContext(context.Background(), request)
}

// DescribeMailProfile
// 获取发送邮件的配置， 包括数据库巡检的邮件配置以及定期生成健康报告的邮件发送配置。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMailProfileWithContext(ctx context.Context, request *DescribeMailProfileRequest) (response *DescribeMailProfileResponse, err error) {
    if request == nil {
        request = NewDescribeMailProfileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeMailProfile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMailProfile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMailProfileResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMetricTopProxiesRequest() (request *DescribeMetricTopProxiesRequest) {
    request = &DescribeMetricTopProxiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeMetricTopProxies")
    
    
    return
}

func NewDescribeMetricTopProxiesResponse() (response *DescribeMetricTopProxiesResponse) {
    response = &DescribeMetricTopProxiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMetricTopProxies
// 获取指定时间段内Redis Proxy 指标
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMetricTopProxies(request *DescribeMetricTopProxiesRequest) (response *DescribeMetricTopProxiesResponse, err error) {
    return c.DescribeMetricTopProxiesWithContext(context.Background(), request)
}

// DescribeMetricTopProxies
// 获取指定时间段内Redis Proxy 指标
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMetricTopProxiesWithContext(ctx context.Context, request *DescribeMetricTopProxiesRequest) (response *DescribeMetricTopProxiesResponse, err error) {
    if request == nil {
        request = NewDescribeMetricTopProxiesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeMetricTopProxies")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMetricTopProxies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMetricTopProxiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMongoDBProcessListRequest() (request *DescribeMongoDBProcessListRequest) {
    request = &DescribeMongoDBProcessListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeMongoDBProcessList")
    
    
    return
}

func NewDescribeMongoDBProcessListResponse() (response *DescribeMongoDBProcessListResponse) {
    response = &DescribeMongoDBProcessListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMongoDBProcessList
// 查询MongoDB实时会话列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeMongoDBProcessList(request *DescribeMongoDBProcessListRequest) (response *DescribeMongoDBProcessListResponse, err error) {
    return c.DescribeMongoDBProcessListWithContext(context.Background(), request)
}

// DescribeMongoDBProcessList
// 查询MongoDB实时会话列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeMongoDBProcessListWithContext(ctx context.Context, request *DescribeMongoDBProcessListRequest) (response *DescribeMongoDBProcessListResponse, err error) {
    if request == nil {
        request = NewDescribeMongoDBProcessListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeMongoDBProcessList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMongoDBProcessList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMongoDBProcessListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMySqlProcessListRequest() (request *DescribeMySqlProcessListRequest) {
    request = &DescribeMySqlProcessListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeMySqlProcessList")
    
    
    return
}

func NewDescribeMySqlProcessListResponse() (response *DescribeMySqlProcessListResponse) {
    response = &DescribeMySqlProcessListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMySqlProcessList
// 查询关系型数据库的实时线程列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMySqlProcessList(request *DescribeMySqlProcessListRequest) (response *DescribeMySqlProcessListResponse, err error) {
    return c.DescribeMySqlProcessListWithContext(context.Background(), request)
}

// DescribeMySqlProcessList
// 查询关系型数据库的实时线程列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMySqlProcessListWithContext(ctx context.Context, request *DescribeMySqlProcessListRequest) (response *DescribeMySqlProcessListResponse, err error) {
    if request == nil {
        request = NewDescribeMySqlProcessListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeMySqlProcessList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMySqlProcessList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMySqlProcessListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNoPrimaryKeyTablesRequest() (request *DescribeNoPrimaryKeyTablesRequest) {
    request = &DescribeNoPrimaryKeyTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeNoPrimaryKeyTables")
    
    
    return
}

func NewDescribeNoPrimaryKeyTablesResponse() (response *DescribeNoPrimaryKeyTablesResponse) {
    response = &DescribeNoPrimaryKeyTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeNoPrimaryKeyTables
// 查询实例无主键表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNoPrimaryKeyTables(request *DescribeNoPrimaryKeyTablesRequest) (response *DescribeNoPrimaryKeyTablesResponse, err error) {
    return c.DescribeNoPrimaryKeyTablesWithContext(context.Background(), request)
}

// DescribeNoPrimaryKeyTables
// 查询实例无主键表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeNoPrimaryKeyTablesWithContext(ctx context.Context, request *DescribeNoPrimaryKeyTablesRequest) (response *DescribeNoPrimaryKeyTablesResponse, err error) {
    if request == nil {
        request = NewDescribeNoPrimaryKeyTablesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeNoPrimaryKeyTables")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNoPrimaryKeyTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNoPrimaryKeyTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxyProcessStatisticsRequest() (request *DescribeProxyProcessStatisticsRequest) {
    request = &DescribeProxyProcessStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeProxyProcessStatistics")
    
    
    return
}

func NewDescribeProxyProcessStatisticsResponse() (response *DescribeProxyProcessStatisticsResponse) {
    response = &DescribeProxyProcessStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProxyProcessStatistics
// 获取当前实例下的单个proxy的会话统计详情信息， 返回数据为单个 proxy 的会话统计信息。【注意】该接口仅限部分环境调用。
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
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProxyProcessStatistics(request *DescribeProxyProcessStatisticsRequest) (response *DescribeProxyProcessStatisticsResponse, err error) {
    return c.DescribeProxyProcessStatisticsWithContext(context.Background(), request)
}

// DescribeProxyProcessStatistics
// 获取当前实例下的单个proxy的会话统计详情信息， 返回数据为单个 proxy 的会话统计信息。【注意】该接口仅限部分环境调用。
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
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProxyProcessStatisticsWithContext(ctx context.Context, request *DescribeProxyProcessStatisticsRequest) (response *DescribeProxyProcessStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeProxyProcessStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeProxyProcessStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProxyProcessStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProxyProcessStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeProxySessionKillTasksRequest() (request *DescribeProxySessionKillTasksRequest) {
    request = &DescribeProxySessionKillTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeProxySessionKillTasks")
    
    
    return
}

func NewDescribeProxySessionKillTasksResponse() (response *DescribeProxySessionKillTasksResponse) {
    response = &DescribeProxySessionKillTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeProxySessionKillTasks
// 用于查询 redis 执行 kill 会话任务后代理节点的执行结果，入参异步任务 ID 从接口 CreateProxySessionKillTask 调用成功后取得。当前 product 只支持：redis。
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
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProxySessionKillTasks(request *DescribeProxySessionKillTasksRequest) (response *DescribeProxySessionKillTasksResponse, err error) {
    return c.DescribeProxySessionKillTasksWithContext(context.Background(), request)
}

// DescribeProxySessionKillTasks
// 用于查询 redis 执行 kill 会话任务后代理节点的执行结果，入参异步任务 ID 从接口 CreateProxySessionKillTask 调用成功后取得。当前 product 只支持：redis。
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
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeProxySessionKillTasksWithContext(ctx context.Context, request *DescribeProxySessionKillTasksRequest) (response *DescribeProxySessionKillTasksResponse, err error) {
    if request == nil {
        request = NewDescribeProxySessionKillTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeProxySessionKillTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProxySessionKillTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProxySessionKillTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRedisBigKeyAnalysisTasksRequest() (request *DescribeRedisBigKeyAnalysisTasksRequest) {
    request = &DescribeRedisBigKeyAnalysisTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeRedisBigKeyAnalysisTasks")
    
    
    return
}

func NewDescribeRedisBigKeyAnalysisTasksResponse() (response *DescribeRedisBigKeyAnalysisTasksResponse) {
    response = &DescribeRedisBigKeyAnalysisTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRedisBigKeyAnalysisTasks
// 查询redis大key分析任务列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisBigKeyAnalysisTasks(request *DescribeRedisBigKeyAnalysisTasksRequest) (response *DescribeRedisBigKeyAnalysisTasksResponse, err error) {
    return c.DescribeRedisBigKeyAnalysisTasksWithContext(context.Background(), request)
}

// DescribeRedisBigKeyAnalysisTasks
// 查询redis大key分析任务列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisBigKeyAnalysisTasksWithContext(ctx context.Context, request *DescribeRedisBigKeyAnalysisTasksRequest) (response *DescribeRedisBigKeyAnalysisTasksResponse, err error) {
    if request == nil {
        request = NewDescribeRedisBigKeyAnalysisTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeRedisBigKeyAnalysisTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRedisBigKeyAnalysisTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRedisBigKeyAnalysisTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRedisCmdPerfTimeSeriesRequest() (request *DescribeRedisCmdPerfTimeSeriesRequest) {
    request = &DescribeRedisCmdPerfTimeSeriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeRedisCmdPerfTimeSeries")
    
    
    return
}

func NewDescribeRedisCmdPerfTimeSeriesResponse() (response *DescribeRedisCmdPerfTimeSeriesResponse) {
    response = &DescribeRedisCmdPerfTimeSeriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRedisCmdPerfTimeSeries
// 延迟分析-命令字分析-查询命令延迟趋势
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisCmdPerfTimeSeries(request *DescribeRedisCmdPerfTimeSeriesRequest) (response *DescribeRedisCmdPerfTimeSeriesResponse, err error) {
    return c.DescribeRedisCmdPerfTimeSeriesWithContext(context.Background(), request)
}

// DescribeRedisCmdPerfTimeSeries
// 延迟分析-命令字分析-查询命令延迟趋势
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisCmdPerfTimeSeriesWithContext(ctx context.Context, request *DescribeRedisCmdPerfTimeSeriesRequest) (response *DescribeRedisCmdPerfTimeSeriesResponse, err error) {
    if request == nil {
        request = NewDescribeRedisCmdPerfTimeSeriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeRedisCmdPerfTimeSeries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRedisCmdPerfTimeSeries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRedisCmdPerfTimeSeriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRedisCommandCostStatisticsRequest() (request *DescribeRedisCommandCostStatisticsRequest) {
    request = &DescribeRedisCommandCostStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeRedisCommandCostStatistics")
    
    
    return
}

func NewDescribeRedisCommandCostStatisticsResponse() (response *DescribeRedisCommandCostStatisticsResponse) {
    response = &DescribeRedisCommandCostStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRedisCommandCostStatistics
// 延迟分析-查询命令延迟分布
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisCommandCostStatistics(request *DescribeRedisCommandCostStatisticsRequest) (response *DescribeRedisCommandCostStatisticsResponse, err error) {
    return c.DescribeRedisCommandCostStatisticsWithContext(context.Background(), request)
}

// DescribeRedisCommandCostStatistics
// 延迟分析-查询命令延迟分布
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisCommandCostStatisticsWithContext(ctx context.Context, request *DescribeRedisCommandCostStatisticsRequest) (response *DescribeRedisCommandCostStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeRedisCommandCostStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeRedisCommandCostStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRedisCommandCostStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRedisCommandCostStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRedisCommandOverviewRequest() (request *DescribeRedisCommandOverviewRequest) {
    request = &DescribeRedisCommandOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeRedisCommandOverview")
    
    
    return
}

func NewDescribeRedisCommandOverviewResponse() (response *DescribeRedisCommandOverviewResponse) {
    response = &DescribeRedisCommandOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRedisCommandOverview
// 延迟分析-查询实例访问命令统计
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisCommandOverview(request *DescribeRedisCommandOverviewRequest) (response *DescribeRedisCommandOverviewResponse, err error) {
    return c.DescribeRedisCommandOverviewWithContext(context.Background(), request)
}

// DescribeRedisCommandOverview
// 延迟分析-查询实例访问命令统计
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisCommandOverviewWithContext(ctx context.Context, request *DescribeRedisCommandOverviewRequest) (response *DescribeRedisCommandOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeRedisCommandOverviewRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeRedisCommandOverview")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRedisCommandOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRedisCommandOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRedisProcessListRequest() (request *DescribeRedisProcessListRequest) {
    request = &DescribeRedisProcessListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeRedisProcessList")
    
    
    return
}

func NewDescribeRedisProcessListResponse() (response *DescribeRedisProcessListResponse) {
    response = &DescribeRedisProcessListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRedisProcessList
// 获取 Redis 实例所有 proxy 节点的实时会话详情列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisProcessList(request *DescribeRedisProcessListRequest) (response *DescribeRedisProcessListResponse, err error) {
    return c.DescribeRedisProcessListWithContext(context.Background(), request)
}

// DescribeRedisProcessList
// 获取 Redis 实例所有 proxy 节点的实时会话详情列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisProcessListWithContext(ctx context.Context, request *DescribeRedisProcessListRequest) (response *DescribeRedisProcessListResponse, err error) {
    if request == nil {
        request = NewDescribeRedisProcessListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeRedisProcessList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRedisProcessList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRedisProcessListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRedisSlowLogTopSqlsRequest() (request *DescribeRedisSlowLogTopSqlsRequest) {
    request = &DescribeRedisSlowLogTopSqlsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeRedisSlowLogTopSqls")
    
    
    return
}

func NewDescribeRedisSlowLogTopSqlsResponse() (response *DescribeRedisSlowLogTopSqlsResponse) {
    response = &DescribeRedisSlowLogTopSqlsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRedisSlowLogTopSqls
// 统计排序指定时间段内的top慢sql。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEGMENTLOADING = "FailedOperation.SegmentLoading"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisSlowLogTopSqls(request *DescribeRedisSlowLogTopSqlsRequest) (response *DescribeRedisSlowLogTopSqlsResponse, err error) {
    return c.DescribeRedisSlowLogTopSqlsWithContext(context.Background(), request)
}

// DescribeRedisSlowLogTopSqls
// 统计排序指定时间段内的top慢sql。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEGMENTLOADING = "FailedOperation.SegmentLoading"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisSlowLogTopSqlsWithContext(ctx context.Context, request *DescribeRedisSlowLogTopSqlsRequest) (response *DescribeRedisSlowLogTopSqlsResponse, err error) {
    if request == nil {
        request = NewDescribeRedisSlowLogTopSqlsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeRedisSlowLogTopSqls")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRedisSlowLogTopSqls require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRedisSlowLogTopSqlsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRedisTopBigKeysRequest() (request *DescribeRedisTopBigKeysRequest) {
    request = &DescribeRedisTopBigKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeRedisTopBigKeys")
    
    
    return
}

func NewDescribeRedisTopBigKeysResponse() (response *DescribeRedisTopBigKeysResponse) {
    response = &DescribeRedisTopBigKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRedisTopBigKeys
// 查询redis实例大key列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisTopBigKeys(request *DescribeRedisTopBigKeysRequest) (response *DescribeRedisTopBigKeysResponse, err error) {
    return c.DescribeRedisTopBigKeysWithContext(context.Background(), request)
}

// DescribeRedisTopBigKeys
// 查询redis实例大key列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisTopBigKeysWithContext(ctx context.Context, request *DescribeRedisTopBigKeysRequest) (response *DescribeRedisTopBigKeysResponse, err error) {
    if request == nil {
        request = NewDescribeRedisTopBigKeysRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeRedisTopBigKeys")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRedisTopBigKeys require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRedisTopBigKeysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRedisTopCostCommandsRequest() (request *DescribeRedisTopCostCommandsRequest) {
    request = &DescribeRedisTopCostCommandsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeRedisTopCostCommands")
    
    
    return
}

func NewDescribeRedisTopCostCommandsResponse() (response *DescribeRedisTopCostCommandsResponse) {
    response = &DescribeRedisTopCostCommandsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRedisTopCostCommands
// 获取指定时间段内Redis 访问命令 cost top N
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisTopCostCommands(request *DescribeRedisTopCostCommandsRequest) (response *DescribeRedisTopCostCommandsResponse, err error) {
    return c.DescribeRedisTopCostCommandsWithContext(context.Background(), request)
}

// DescribeRedisTopCostCommands
// 获取指定时间段内Redis 访问命令 cost top N
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisTopCostCommandsWithContext(ctx context.Context, request *DescribeRedisTopCostCommandsRequest) (response *DescribeRedisTopCostCommandsResponse, err error) {
    if request == nil {
        request = NewDescribeRedisTopCostCommandsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeRedisTopCostCommands")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRedisTopCostCommands require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRedisTopCostCommandsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRedisTopHotKeysRequest() (request *DescribeRedisTopHotKeysRequest) {
    request = &DescribeRedisTopHotKeysRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeRedisTopHotKeys")
    
    
    return
}

func NewDescribeRedisTopHotKeysResponse() (response *DescribeRedisTopHotKeysResponse) {
    response = &DescribeRedisTopHotKeysResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRedisTopHotKeys
// 热Key分析
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisTopHotKeys(request *DescribeRedisTopHotKeysRequest) (response *DescribeRedisTopHotKeysResponse, err error) {
    return c.DescribeRedisTopHotKeysWithContext(context.Background(), request)
}

// DescribeRedisTopHotKeys
// 热Key分析
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisTopHotKeysWithContext(ctx context.Context, request *DescribeRedisTopHotKeysRequest) (response *DescribeRedisTopHotKeysResponse, err error) {
    if request == nil {
        request = NewDescribeRedisTopHotKeysRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeRedisTopHotKeys")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRedisTopHotKeys require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRedisTopHotKeysResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRedisTopKeyPrefixListRequest() (request *DescribeRedisTopKeyPrefixListRequest) {
    request = &DescribeRedisTopKeyPrefixListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeRedisTopKeyPrefixList")
    
    
    return
}

func NewDescribeRedisTopKeyPrefixListResponse() (response *DescribeRedisTopKeyPrefixListResponse) {
    response = &DescribeRedisTopKeyPrefixListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRedisTopKeyPrefixList
// 查询redis实例top key前缀列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisTopKeyPrefixList(request *DescribeRedisTopKeyPrefixListRequest) (response *DescribeRedisTopKeyPrefixListResponse, err error) {
    return c.DescribeRedisTopKeyPrefixListWithContext(context.Background(), request)
}

// DescribeRedisTopKeyPrefixList
// 查询redis实例top key前缀列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisTopKeyPrefixListWithContext(ctx context.Context, request *DescribeRedisTopKeyPrefixListRequest) (response *DescribeRedisTopKeyPrefixListResponse, err error) {
    if request == nil {
        request = NewDescribeRedisTopKeyPrefixListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeRedisTopKeyPrefixList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRedisTopKeyPrefixList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRedisTopKeyPrefixListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRedisUnExpiredKeyStatisticsRequest() (request *DescribeRedisUnExpiredKeyStatisticsRequest) {
    request = &DescribeRedisUnExpiredKeyStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeRedisUnExpiredKeyStatistics")
    
    
    return
}

func NewDescribeRedisUnExpiredKeyStatisticsResponse() (response *DescribeRedisUnExpiredKeyStatisticsResponse) {
    response = &DescribeRedisUnExpiredKeyStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRedisUnExpiredKeyStatistics
// 查询Redis全量Key的内存分布情况。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisUnExpiredKeyStatistics(request *DescribeRedisUnExpiredKeyStatisticsRequest) (response *DescribeRedisUnExpiredKeyStatisticsResponse, err error) {
    return c.DescribeRedisUnExpiredKeyStatisticsWithContext(context.Background(), request)
}

// DescribeRedisUnExpiredKeyStatistics
// 查询Redis全量Key的内存分布情况。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRedisUnExpiredKeyStatisticsWithContext(ctx context.Context, request *DescribeRedisUnExpiredKeyStatisticsRequest) (response *DescribeRedisUnExpiredKeyStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeRedisUnExpiredKeyStatisticsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeRedisUnExpiredKeyStatistics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRedisUnExpiredKeyStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRedisUnExpiredKeyStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityAuditLogDownloadUrlsRequest() (request *DescribeSecurityAuditLogDownloadUrlsRequest) {
    request = &DescribeSecurityAuditLogDownloadUrlsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeSecurityAuditLogDownloadUrls")
    
    
    return
}

func NewDescribeSecurityAuditLogDownloadUrlsResponse() (response *DescribeSecurityAuditLogDownloadUrlsResponse) {
    response = &DescribeSecurityAuditLogDownloadUrlsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityAuditLogDownloadUrls
// 查询安全审计日志导出文件下载链接。目前日志文件下载仅提供腾讯云内网地址，请通过广州地域的腾讯云服务器进行下载。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityAuditLogDownloadUrls(request *DescribeSecurityAuditLogDownloadUrlsRequest) (response *DescribeSecurityAuditLogDownloadUrlsResponse, err error) {
    return c.DescribeSecurityAuditLogDownloadUrlsWithContext(context.Background(), request)
}

// DescribeSecurityAuditLogDownloadUrls
// 查询安全审计日志导出文件下载链接。目前日志文件下载仅提供腾讯云内网地址，请通过广州地域的腾讯云服务器进行下载。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityAuditLogDownloadUrlsWithContext(ctx context.Context, request *DescribeSecurityAuditLogDownloadUrlsRequest) (response *DescribeSecurityAuditLogDownloadUrlsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityAuditLogDownloadUrlsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeSecurityAuditLogDownloadUrls")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityAuditLogDownloadUrls require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityAuditLogDownloadUrlsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSecurityAuditLogExportTasksRequest() (request *DescribeSecurityAuditLogExportTasksRequest) {
    request = &DescribeSecurityAuditLogExportTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeSecurityAuditLogExportTasks")
    
    
    return
}

func NewDescribeSecurityAuditLogExportTasksResponse() (response *DescribeSecurityAuditLogExportTasksResponse) {
    response = &DescribeSecurityAuditLogExportTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSecurityAuditLogExportTasks
// 查询安全审计日志导出任务列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityAuditLogExportTasks(request *DescribeSecurityAuditLogExportTasksRequest) (response *DescribeSecurityAuditLogExportTasksResponse, err error) {
    return c.DescribeSecurityAuditLogExportTasksWithContext(context.Background(), request)
}

// DescribeSecurityAuditLogExportTasks
// 查询安全审计日志导出任务列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSecurityAuditLogExportTasksWithContext(ctx context.Context, request *DescribeSecurityAuditLogExportTasksRequest) (response *DescribeSecurityAuditLogExportTasksResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityAuditLogExportTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeSecurityAuditLogExportTasks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityAuditLogExportTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityAuditLogExportTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogQueryTimeStatsRequest() (request *DescribeSlowLogQueryTimeStatsRequest) {
    request = &DescribeSlowLogQueryTimeStatsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeSlowLogQueryTimeStats")
    
    
    return
}

func NewDescribeSlowLogQueryTimeStatsResponse() (response *DescribeSlowLogQueryTimeStatsResponse) {
    response = &DescribeSlowLogQueryTimeStatsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowLogQueryTimeStats
// 统计排序指定时间段内的top慢sql。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEGMENTLOADING = "FailedOperation.SegmentLoading"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSlowLogQueryTimeStats(request *DescribeSlowLogQueryTimeStatsRequest) (response *DescribeSlowLogQueryTimeStatsResponse, err error) {
    return c.DescribeSlowLogQueryTimeStatsWithContext(context.Background(), request)
}

// DescribeSlowLogQueryTimeStats
// 统计排序指定时间段内的top慢sql。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEGMENTLOADING = "FailedOperation.SegmentLoading"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSlowLogQueryTimeStatsWithContext(ctx context.Context, request *DescribeSlowLogQueryTimeStatsRequest) (response *DescribeSlowLogQueryTimeStatsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogQueryTimeStatsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeSlowLogQueryTimeStats")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLogQueryTimeStats require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogQueryTimeStatsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogTimeSeriesStatsRequest() (request *DescribeSlowLogTimeSeriesStatsRequest) {
    request = &DescribeSlowLogTimeSeriesStatsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeSlowLogTimeSeriesStats")
    
    
    return
}

func NewDescribeSlowLogTimeSeriesStatsResponse() (response *DescribeSlowLogTimeSeriesStatsResponse) {
    response = &DescribeSlowLogTimeSeriesStatsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowLogTimeSeriesStats
// 获取慢日志统计柱状图。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEGMENTLOADING = "FailedOperation.SegmentLoading"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSlowLogTimeSeriesStats(request *DescribeSlowLogTimeSeriesStatsRequest) (response *DescribeSlowLogTimeSeriesStatsResponse, err error) {
    return c.DescribeSlowLogTimeSeriesStatsWithContext(context.Background(), request)
}

// DescribeSlowLogTimeSeriesStats
// 获取慢日志统计柱状图。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEGMENTLOADING = "FailedOperation.SegmentLoading"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSlowLogTimeSeriesStatsWithContext(ctx context.Context, request *DescribeSlowLogTimeSeriesStatsRequest) (response *DescribeSlowLogTimeSeriesStatsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogTimeSeriesStatsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeSlowLogTimeSeriesStats")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLogTimeSeriesStats require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogTimeSeriesStatsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogTopSqlsRequest() (request *DescribeSlowLogTopSqlsRequest) {
    request = &DescribeSlowLogTopSqlsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeSlowLogTopSqls")
    
    
    return
}

func NewDescribeSlowLogTopSqlsResponse() (response *DescribeSlowLogTopSqlsResponse) {
    response = &DescribeSlowLogTopSqlsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowLogTopSqls
// 按照Sql模板+schema的聚合方式，统计排序指定时间段内的top慢sql。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEGMENTLOADING = "FailedOperation.SegmentLoading"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSlowLogTopSqls(request *DescribeSlowLogTopSqlsRequest) (response *DescribeSlowLogTopSqlsResponse, err error) {
    return c.DescribeSlowLogTopSqlsWithContext(context.Background(), request)
}

// DescribeSlowLogTopSqls
// 按照Sql模板+schema的聚合方式，统计排序指定时间段内的top慢sql。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEGMENTLOADING = "FailedOperation.SegmentLoading"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSlowLogTopSqlsWithContext(ctx context.Context, request *DescribeSlowLogTopSqlsRequest) (response *DescribeSlowLogTopSqlsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogTopSqlsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeSlowLogTopSqls")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLogTopSqls require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogTopSqlsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogUserHostStatsRequest() (request *DescribeSlowLogUserHostStatsRequest) {
    request = &DescribeSlowLogUserHostStatsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeSlowLogUserHostStats")
    
    
    return
}

func NewDescribeSlowLogUserHostStatsResponse() (response *DescribeSlowLogUserHostStatsResponse) {
    response = &DescribeSlowLogUserHostStatsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowLogUserHostStats
// 获取慢日志来源地址统计分布图。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEGMENTLOADING = "FailedOperation.SegmentLoading"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSlowLogUserHostStats(request *DescribeSlowLogUserHostStatsRequest) (response *DescribeSlowLogUserHostStatsResponse, err error) {
    return c.DescribeSlowLogUserHostStatsWithContext(context.Background(), request)
}

// DescribeSlowLogUserHostStats
// 获取慢日志来源地址统计分布图。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEGMENTLOADING = "FailedOperation.SegmentLoading"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSlowLogUserHostStatsWithContext(ctx context.Context, request *DescribeSlowLogUserHostStatsRequest) (response *DescribeSlowLogUserHostStatsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogUserHostStatsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeSlowLogUserHostStats")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLogUserHostStats require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogUserHostStatsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSlowLogsRequest() (request *DescribeSlowLogsRequest) {
    request = &DescribeSlowLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeSlowLogs")
    
    
    return
}

func NewDescribeSlowLogsResponse() (response *DescribeSlowLogsResponse) {
    response = &DescribeSlowLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSlowLogs
// 获取指定时间内某个sql模板的慢日志明细
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEGMENTLOADING = "FailedOperation.SegmentLoading"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSlowLogs(request *DescribeSlowLogsRequest) (response *DescribeSlowLogsResponse, err error) {
    return c.DescribeSlowLogsWithContext(context.Background(), request)
}

// DescribeSlowLogs
// 获取指定时间内某个sql模板的慢日志明细
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SEGMENTLOADING = "FailedOperation.SegmentLoading"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSlowLogsWithContext(ctx context.Context, request *DescribeSlowLogsRequest) (response *DescribeSlowLogsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeSlowLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSqlFiltersRequest() (request *DescribeSqlFiltersRequest) {
    request = &DescribeSqlFiltersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeSqlFilters")
    
    
    return
}

func NewDescribeSqlFiltersResponse() (response *DescribeSqlFiltersResponse) {
    response = &DescribeSqlFiltersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSqlFilters
// 查询实例SQL限流任务列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSqlFilters(request *DescribeSqlFiltersRequest) (response *DescribeSqlFiltersResponse, err error) {
    return c.DescribeSqlFiltersWithContext(context.Background(), request)
}

// DescribeSqlFilters
// 查询实例SQL限流任务列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSqlFiltersWithContext(ctx context.Context, request *DescribeSqlFiltersRequest) (response *DescribeSqlFiltersResponse, err error) {
    if request == nil {
        request = NewDescribeSqlFiltersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeSqlFilters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSqlFilters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSqlFiltersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSqlTemplateRequest() (request *DescribeSqlTemplateRequest) {
    request = &DescribeSqlTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeSqlTemplate")
    
    
    return
}

func NewDescribeSqlTemplateResponse() (response *DescribeSqlTemplateResponse) {
    response = &DescribeSqlTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSqlTemplate
// 查询SQL模板。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSqlTemplate(request *DescribeSqlTemplateRequest) (response *DescribeSqlTemplateResponse, err error) {
    return c.DescribeSqlTemplateWithContext(context.Background(), request)
}

// DescribeSqlTemplate
// 查询SQL模板。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSqlTemplateWithContext(ctx context.Context, request *DescribeSqlTemplateRequest) (response *DescribeSqlTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeSqlTemplateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeSqlTemplate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSqlTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSqlTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopSpaceSchemaTimeSeriesRequest() (request *DescribeTopSpaceSchemaTimeSeriesRequest) {
    request = &DescribeTopSpaceSchemaTimeSeriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeTopSpaceSchemaTimeSeries")
    
    
    return
}

func NewDescribeTopSpaceSchemaTimeSeriesResponse() (response *DescribeTopSpaceSchemaTimeSeriesResponse) {
    response = &DescribeTopSpaceSchemaTimeSeriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopSpaceSchemaTimeSeries
// 获取实例占用空间最大的前几个库在指定时间段内的每日由DBbrain定时采集的空间数据，默认返回按大小排序。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopSpaceSchemaTimeSeries(request *DescribeTopSpaceSchemaTimeSeriesRequest) (response *DescribeTopSpaceSchemaTimeSeriesResponse, err error) {
    return c.DescribeTopSpaceSchemaTimeSeriesWithContext(context.Background(), request)
}

// DescribeTopSpaceSchemaTimeSeries
// 获取实例占用空间最大的前几个库在指定时间段内的每日由DBbrain定时采集的空间数据，默认返回按大小排序。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopSpaceSchemaTimeSeriesWithContext(ctx context.Context, request *DescribeTopSpaceSchemaTimeSeriesRequest) (response *DescribeTopSpaceSchemaTimeSeriesResponse, err error) {
    if request == nil {
        request = NewDescribeTopSpaceSchemaTimeSeriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeTopSpaceSchemaTimeSeries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopSpaceSchemaTimeSeries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopSpaceSchemaTimeSeriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopSpaceSchemasRequest() (request *DescribeTopSpaceSchemasRequest) {
    request = &DescribeTopSpaceSchemasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeTopSpaceSchemas")
    
    
    return
}

func NewDescribeTopSpaceSchemasResponse() (response *DescribeTopSpaceSchemasResponse) {
    response = &DescribeTopSpaceSchemasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopSpaceSchemas
// 获取实例Top库的实时空间统计信息，默认返回按大小排序。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopSpaceSchemas(request *DescribeTopSpaceSchemasRequest) (response *DescribeTopSpaceSchemasResponse, err error) {
    return c.DescribeTopSpaceSchemasWithContext(context.Background(), request)
}

// DescribeTopSpaceSchemas
// 获取实例Top库的实时空间统计信息，默认返回按大小排序。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopSpaceSchemasWithContext(ctx context.Context, request *DescribeTopSpaceSchemasRequest) (response *DescribeTopSpaceSchemasResponse, err error) {
    if request == nil {
        request = NewDescribeTopSpaceSchemasRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeTopSpaceSchemas")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopSpaceSchemas require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopSpaceSchemasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopSpaceTableTimeSeriesRequest() (request *DescribeTopSpaceTableTimeSeriesRequest) {
    request = &DescribeTopSpaceTableTimeSeriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeTopSpaceTableTimeSeries")
    
    
    return
}

func NewDescribeTopSpaceTableTimeSeriesResponse() (response *DescribeTopSpaceTableTimeSeriesResponse) {
    response = &DescribeTopSpaceTableTimeSeriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopSpaceTableTimeSeries
// 获取实例占用空间最大的前几张表在指定时间段内的每日由DBbrain定时采集的空间数据，默认返回按大小排序。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopSpaceTableTimeSeries(request *DescribeTopSpaceTableTimeSeriesRequest) (response *DescribeTopSpaceTableTimeSeriesResponse, err error) {
    return c.DescribeTopSpaceTableTimeSeriesWithContext(context.Background(), request)
}

// DescribeTopSpaceTableTimeSeries
// 获取实例占用空间最大的前几张表在指定时间段内的每日由DBbrain定时采集的空间数据，默认返回按大小排序。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopSpaceTableTimeSeriesWithContext(ctx context.Context, request *DescribeTopSpaceTableTimeSeriesRequest) (response *DescribeTopSpaceTableTimeSeriesResponse, err error) {
    if request == nil {
        request = NewDescribeTopSpaceTableTimeSeriesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeTopSpaceTableTimeSeries")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopSpaceTableTimeSeries require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopSpaceTableTimeSeriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopSpaceTablesRequest() (request *DescribeTopSpaceTablesRequest) {
    request = &DescribeTopSpaceTablesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeTopSpaceTables")
    
    
    return
}

func NewDescribeTopSpaceTablesResponse() (response *DescribeTopSpaceTablesResponse) {
    response = &DescribeTopSpaceTablesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTopSpaceTables
// 获取实例Top表的实时空间统计信息，默认返回按大小排序。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopSpaceTables(request *DescribeTopSpaceTablesRequest) (response *DescribeTopSpaceTablesResponse, err error) {
    return c.DescribeTopSpaceTablesWithContext(context.Background(), request)
}

// DescribeTopSpaceTables
// 获取实例Top表的实时空间统计信息，默认返回按大小排序。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopSpaceTablesWithContext(ctx context.Context, request *DescribeTopSpaceTablesRequest) (response *DescribeTopSpaceTablesResponse, err error) {
    if request == nil {
        request = NewDescribeTopSpaceTablesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeTopSpaceTables")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopSpaceTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopSpaceTablesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserAutonomyProfileRequest() (request *DescribeUserAutonomyProfileRequest) {
    request = &DescribeUserAutonomyProfileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeUserAutonomyProfile")
    
    
    return
}

func NewDescribeUserAutonomyProfileResponse() (response *DescribeUserAutonomyProfileResponse) {
    response = &DescribeUserAutonomyProfileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserAutonomyProfile
// 自治中心-终止自治任务（单次）；注意： 接口调用需要加白名单。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserAutonomyProfile(request *DescribeUserAutonomyProfileRequest) (response *DescribeUserAutonomyProfileResponse, err error) {
    return c.DescribeUserAutonomyProfileWithContext(context.Background(), request)
}

// DescribeUserAutonomyProfile
// 自治中心-终止自治任务（单次）；注意： 接口调用需要加白名单。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserAutonomyProfileWithContext(ctx context.Context, request *DescribeUserAutonomyProfileRequest) (response *DescribeUserAutonomyProfileResponse, err error) {
    if request == nil {
        request = NewDescribeUserAutonomyProfileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeUserAutonomyProfile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserAutonomyProfile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserAutonomyProfileResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserSqlAdviceRequest() (request *DescribeUserSqlAdviceRequest) {
    request = &DescribeUserSqlAdviceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "DescribeUserSqlAdvice")
    
    
    return
}

func NewDescribeUserSqlAdviceResponse() (response *DescribeUserSqlAdviceResponse) {
    response = &DescribeUserSqlAdviceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserSqlAdvice
// 获取SQL优化建议。【产品用户回馈，此接口限免开放，后续将并入dbbrain专业版】
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserSqlAdvice(request *DescribeUserSqlAdviceRequest) (response *DescribeUserSqlAdviceResponse, err error) {
    return c.DescribeUserSqlAdviceWithContext(context.Background(), request)
}

// DescribeUserSqlAdvice
// 获取SQL优化建议。【产品用户回馈，此接口限免开放，后续将并入dbbrain专业版】
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserSqlAdviceWithContext(ctx context.Context, request *DescribeUserSqlAdviceRequest) (response *DescribeUserSqlAdviceResponse, err error) {
    if request == nil {
        request = NewDescribeUserSqlAdviceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "DescribeUserSqlAdvice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserSqlAdvice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserSqlAdviceResponse()
    err = c.Send(request, response)
    return
}

func NewKillMySqlThreadsRequest() (request *KillMySqlThreadsRequest) {
    request = &KillMySqlThreadsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "KillMySqlThreads")
    
    
    return
}

func NewKillMySqlThreadsResponse() (response *KillMySqlThreadsResponse) {
    response = &KillMySqlThreadsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// KillMySqlThreads
// 根据会话ID中断当前会话，该接口分为两次提交：第一次为预提交阶段，Stage为"Prepare"，得到的返回值包含SqlExecId；第二次为确认提交， Stage为"Commit"， 将SqlExecId的值作为参数传入，最终终止会话进程。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) KillMySqlThreads(request *KillMySqlThreadsRequest) (response *KillMySqlThreadsResponse, err error) {
    return c.KillMySqlThreadsWithContext(context.Background(), request)
}

// KillMySqlThreads
// 根据会话ID中断当前会话，该接口分为两次提交：第一次为预提交阶段，Stage为"Prepare"，得到的返回值包含SqlExecId；第二次为确认提交， Stage为"Commit"， 将SqlExecId的值作为参数传入，最终终止会话进程。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) KillMySqlThreadsWithContext(ctx context.Context, request *KillMySqlThreadsRequest) (response *KillMySqlThreadsResponse, err error) {
    if request == nil {
        request = NewKillMySqlThreadsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "KillMySqlThreads")
    
    if c.GetCredential() == nil {
        return nil, errors.New("KillMySqlThreads require credential")
    }

    request.SetContext(ctx)
    
    response = NewKillMySqlThreadsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAlarmPolicyRequest() (request *ModifyAlarmPolicyRequest) {
    request = &ModifyAlarmPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "ModifyAlarmPolicy")
    
    
    return
}

func NewModifyAlarmPolicyResponse() (response *ModifyAlarmPolicyResponse) {
    response = &ModifyAlarmPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAlarmPolicy
// 修改告警策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyAlarmPolicy(request *ModifyAlarmPolicyRequest) (response *ModifyAlarmPolicyResponse, err error) {
    return c.ModifyAlarmPolicyWithContext(context.Background(), request)
}

// ModifyAlarmPolicy
// 修改告警策略
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyAlarmPolicyWithContext(ctx context.Context, request *ModifyAlarmPolicyRequest) (response *ModifyAlarmPolicyResponse, err error) {
    if request == nil {
        request = NewModifyAlarmPolicyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "ModifyAlarmPolicy")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAlarmPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAlarmPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAuditServiceRequest() (request *ModifyAuditServiceRequest) {
    request = &ModifyAuditServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "ModifyAuditService")
    
    
    return
}

func NewModifyAuditServiceResponse() (response *ModifyAuditServiceResponse) {
    response = &ModifyAuditServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAuditService
// 修改审计配置相关信息，如高频存储时长等
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyAuditService(request *ModifyAuditServiceRequest) (response *ModifyAuditServiceResponse, err error) {
    return c.ModifyAuditServiceWithContext(context.Background(), request)
}

// ModifyAuditService
// 修改审计配置相关信息，如高频存储时长等
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyAuditServiceWithContext(ctx context.Context, request *ModifyAuditServiceRequest) (response *ModifyAuditServiceResponse, err error) {
    if request == nil {
        request = NewModifyAuditServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "ModifyAuditService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAuditService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAuditServiceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDiagDBInstanceConfRequest() (request *ModifyDiagDBInstanceConfRequest) {
    request = &ModifyDiagDBInstanceConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "ModifyDiagDBInstanceConf")
    
    
    return
}

func NewModifyDiagDBInstanceConfResponse() (response *ModifyDiagDBInstanceConfResponse) {
    response = &ModifyDiagDBInstanceConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyDiagDBInstanceConf
// 修改实例的配置信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDiagDBInstanceConf(request *ModifyDiagDBInstanceConfRequest) (response *ModifyDiagDBInstanceConfResponse, err error) {
    return c.ModifyDiagDBInstanceConfWithContext(context.Background(), request)
}

// ModifyDiagDBInstanceConf
// 修改实例的配置信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyDiagDBInstanceConfWithContext(ctx context.Context, request *ModifyDiagDBInstanceConfRequest) (response *ModifyDiagDBInstanceConfResponse, err error) {
    if request == nil {
        request = NewModifyDiagDBInstanceConfRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "ModifyDiagDBInstanceConf")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDiagDBInstanceConf require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDiagDBInstanceConfResponse()
    err = c.Send(request, response)
    return
}

func NewModifySqlFiltersRequest() (request *ModifySqlFiltersRequest) {
    request = &ModifySqlFiltersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "ModifySqlFilters")
    
    
    return
}

func NewModifySqlFiltersResponse() (response *ModifySqlFiltersResponse) {
    response = &ModifySqlFiltersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySqlFilters
// 更改实例限流任务状态，目前仅用于终止限流。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySqlFilters(request *ModifySqlFiltersRequest) (response *ModifySqlFiltersResponse, err error) {
    return c.ModifySqlFiltersWithContext(context.Background(), request)
}

// ModifySqlFilters
// 更改实例限流任务状态，目前仅用于终止限流。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySqlFiltersWithContext(ctx context.Context, request *ModifySqlFiltersRequest) (response *ModifySqlFiltersResponse, err error) {
    if request == nil {
        request = NewModifySqlFiltersRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "ModifySqlFilters")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySqlFilters require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySqlFiltersResponse()
    err = c.Send(request, response)
    return
}

func NewModifyUserAutonomyProfileRequest() (request *ModifyUserAutonomyProfileRequest) {
    request = &ModifyUserAutonomyProfileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "ModifyUserAutonomyProfile")
    
    
    return
}

func NewModifyUserAutonomyProfileResponse() (response *ModifyUserAutonomyProfileResponse) {
    response = &ModifyUserAutonomyProfileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyUserAutonomyProfile
// 自治中心-终止自治任务（单次）；注意：接口需要加白名单。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUserAutonomyProfile(request *ModifyUserAutonomyProfileRequest) (response *ModifyUserAutonomyProfileResponse, err error) {
    return c.ModifyUserAutonomyProfileWithContext(context.Background(), request)
}

// ModifyUserAutonomyProfile
// 自治中心-终止自治任务（单次）；注意：接口需要加白名单。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyUserAutonomyProfileWithContext(ctx context.Context, request *ModifyUserAutonomyProfileRequest) (response *ModifyUserAutonomyProfileResponse, err error) {
    if request == nil {
        request = NewModifyUserAutonomyProfileRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "ModifyUserAutonomyProfile")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyUserAutonomyProfile require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyUserAutonomyProfileResponse()
    err = c.Send(request, response)
    return
}

func NewOpenAuditServiceRequest() (request *OpenAuditServiceRequest) {
    request = &OpenAuditServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "OpenAuditService")
    
    
    return
}

func NewOpenAuditServiceResponse() (response *OpenAuditServiceResponse) {
    response = &OpenAuditServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// OpenAuditService
// 开启数据库审计服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PAASAUDITNOTOPENED = "FailedOperation.PaasAuditNotOpened"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) OpenAuditService(request *OpenAuditServiceRequest) (response *OpenAuditServiceResponse, err error) {
    return c.OpenAuditServiceWithContext(context.Background(), request)
}

// OpenAuditService
// 开启数据库审计服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PAASAUDITNOTOPENED = "FailedOperation.PaasAuditNotOpened"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) OpenAuditServiceWithContext(ctx context.Context, request *OpenAuditServiceRequest) (response *OpenAuditServiceResponse, err error) {
    if request == nil {
        request = NewOpenAuditServiceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "OpenAuditService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("OpenAuditService require credential")
    }

    request.SetContext(ctx)
    
    response = NewOpenAuditServiceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAgentSwitchRequest() (request *UpdateAgentSwitchRequest) {
    request = &UpdateAgentSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "UpdateAgentSwitch")
    
    
    return
}

func NewUpdateAgentSwitchResponse() (response *UpdateAgentSwitchResponse) {
    response = &UpdateAgentSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAgentSwitch
// 更新agent状态（停止或重连Agent）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PAASAUDITNOTOPENED = "FailedOperation.PaasAuditNotOpened"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateAgentSwitch(request *UpdateAgentSwitchRequest) (response *UpdateAgentSwitchResponse, err error) {
    return c.UpdateAgentSwitchWithContext(context.Background(), request)
}

// UpdateAgentSwitch
// 更新agent状态（停止或重连Agent）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PAASAUDITNOTOPENED = "FailedOperation.PaasAuditNotOpened"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateAgentSwitchWithContext(ctx context.Context, request *UpdateAgentSwitchRequest) (response *UpdateAgentSwitchResponse, err error) {
    if request == nil {
        request = NewUpdateAgentSwitchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "UpdateAgentSwitch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAgentSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAgentSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateMonitorSwitchRequest() (request *UpdateMonitorSwitchRequest) {
    request = &UpdateMonitorSwitchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "UpdateMonitorSwitch")
    
    
    return
}

func NewUpdateMonitorSwitchResponse() (response *UpdateMonitorSwitchResponse) {
    response = &UpdateMonitorSwitchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateMonitorSwitch
// 更新Agent实例状态（停止或重连实例）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PAASAUDITNOTOPENED = "FailedOperation.PaasAuditNotOpened"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateMonitorSwitch(request *UpdateMonitorSwitchRequest) (response *UpdateMonitorSwitchResponse, err error) {
    return c.UpdateMonitorSwitchWithContext(context.Background(), request)
}

// UpdateMonitorSwitch
// 更新Agent实例状态（停止或重连实例）
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PAASAUDITNOTOPENED = "FailedOperation.PaasAuditNotOpened"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateMonitorSwitchWithContext(ctx context.Context, request *UpdateMonitorSwitchRequest) (response *UpdateMonitorSwitchResponse, err error) {
    if request == nil {
        request = NewUpdateMonitorSwitchRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "UpdateMonitorSwitch")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateMonitorSwitch require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateMonitorSwitchResponse()
    err = c.Send(request, response)
    return
}

func NewVerifyUserAccountRequest() (request *VerifyUserAccountRequest) {
    request = &VerifyUserAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dbbrain", APIVersion, "VerifyUserAccount")
    
    
    return
}

func NewVerifyUserAccountResponse() (response *VerifyUserAccountResponse) {
    response = &VerifyUserAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VerifyUserAccount
// 验证用户数据库账号权限，获取会话token。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) VerifyUserAccount(request *VerifyUserAccountRequest) (response *VerifyUserAccountResponse, err error) {
    return c.VerifyUserAccountWithContext(context.Background(), request)
}

// VerifyUserAccount
// 验证用户数据库账号权限，获取会话token。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) VerifyUserAccountWithContext(ctx context.Context, request *VerifyUserAccountRequest) (response *VerifyUserAccountResponse, err error) {
    if request == nil {
        request = NewVerifyUserAccountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "dbbrain", APIVersion, "VerifyUserAccount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyUserAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyUserAccountResponse()
    err = c.Send(request, response)
    return
}
