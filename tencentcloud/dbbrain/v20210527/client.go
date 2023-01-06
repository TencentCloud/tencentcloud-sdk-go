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
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddUserContact require credential")
    }

    request.SetContext(ctx)
    
    response = NewAddUserContactResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelKillTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelKillTaskResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) CreateProxySessionKillTask(request *CreateProxySessionKillTaskRequest) (response *CreateProxySessionKillTaskResponse, err error) {
    return c.CreateProxySessionKillTaskWithContext(context.Background(), request)
}

// CreateProxySessionKillTask
// 创建中止所有代理节点连接会话的异步任务。当前仅支持 Redis。得到的返回值为异步任务 id，可以作为参数传入接口 DescribeProxySessionKillTasks 查询kill会话任务执行状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) CreateProxySessionKillTaskWithContext(ctx context.Context, request *CreateProxySessionKillTaskRequest) (response *CreateProxySessionKillTaskResponse, err error) {
    if request == nil {
        request = NewCreateProxySessionKillTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProxySessionKillTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProxySessionKillTaskResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSqlFilter require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSqlFilterResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteSqlFilters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteSqlFiltersResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAllUserGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAllUserGroupResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBDiagHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBDiagHistoryResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDBDiagReportTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDBDiagReportTasksResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHealthScore require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHealthScoreResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMailProfile require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMailProfileResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeProxySessionKillTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeProxySessionKillTasksResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRedisTopBigKeys require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRedisTopBigKeysResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRedisTopKeyPrefixList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRedisTopKeyPrefixListResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSecurityAuditLogExportTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSecurityAuditLogExportTasksResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSlowLogUserHostStats require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSlowLogUserHostStatsResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopSpaceTables require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopSpaceTablesResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("KillMySqlThreads require credential")
    }

    request.SetContext(ctx)
    
    response = NewKillMySqlThreadsResponse()
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
// 修改实例巡检开关。
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
// 修改实例巡检开关。
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySqlFilters require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySqlFiltersResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("VerifyUserAccount require credential")
    }

    request.SetContext(ctx)
    
    response = NewVerifyUserAccountResponse()
    err = c.Send(request, response)
    return
}
