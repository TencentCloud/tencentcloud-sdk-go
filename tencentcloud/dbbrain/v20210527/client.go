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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddUserContact(request *AddUserContactRequest) (response *AddUserContactResponse, err error) {
    if request == nil {
        request = NewAddUserContactRequest()
    }
    response = NewAddUserContactResponse()
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
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) CreateDBDiagReportTask(request *CreateDBDiagReportTaskRequest) (response *CreateDBDiagReportTaskResponse, err error) {
    if request == nil {
        request = NewCreateDBDiagReportTaskRequest()
    }
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDBDiagReportUrl(request *CreateDBDiagReportUrlRequest) (response *CreateDBDiagReportUrlResponse, err error) {
    if request == nil {
        request = NewCreateDBDiagReportUrlRequest()
    }
    response = NewCreateDBDiagReportUrlResponse()
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
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateMailProfile(request *CreateMailProfileRequest) (response *CreateMailProfileResponse, err error) {
    if request == nil {
        request = NewCreateMailProfileRequest()
    }
    response = NewCreateMailProfileResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateSchedulerMailProfile(request *CreateSchedulerMailProfileRequest) (response *CreateSchedulerMailProfileResponse, err error) {
    if request == nil {
        request = NewCreateSchedulerMailProfileRequest()
    }
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
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateSecurityAuditLogExportTask(request *CreateSecurityAuditLogExportTaskRequest) (response *CreateSecurityAuditLogExportTaskResponse, err error) {
    if request == nil {
        request = NewCreateSecurityAuditLogExportTaskRequest()
    }
    response = NewCreateSecurityAuditLogExportTaskResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteSecurityAuditLogExportTasks(request *DeleteSecurityAuditLogExportTasksRequest) (response *DeleteSecurityAuditLogExportTasksResponse, err error) {
    if request == nil {
        request = NewDeleteSecurityAuditLogExportTasksRequest()
    }
    response = NewDeleteSecurityAuditLogExportTasksResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAllUserContact(request *DescribeAllUserContactRequest) (response *DescribeAllUserContactResponse, err error) {
    if request == nil {
        request = NewDescribeAllUserContactRequest()
    }
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
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAllUserGroup(request *DescribeAllUserGroupRequest) (response *DescribeAllUserGroupResponse, err error) {
    if request == nil {
        request = NewDescribeAllUserGroupRequest()
    }
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
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDBDiagEvent(request *DescribeDBDiagEventRequest) (response *DescribeDBDiagEventResponse, err error) {
    if request == nil {
        request = NewDescribeDBDiagEventRequest()
    }
    response = NewDescribeDBDiagEventResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDBDiagHistory(request *DescribeDBDiagHistoryRequest) (response *DescribeDBDiagHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeDBDiagHistoryRequest()
    }
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
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeDBDiagReportTasks(request *DescribeDBDiagReportTasksRequest) (response *DescribeDBDiagReportTasksResponse, err error) {
    if request == nil {
        request = NewDescribeDBDiagReportTasksRequest()
    }
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
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeDBSpaceStatus(request *DescribeDBSpaceStatusRequest) (response *DescribeDBSpaceStatusResponse, err error) {
    if request == nil {
        request = NewDescribeDBSpaceStatusRequest()
    }
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
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeDiagDBInstances(request *DescribeDiagDBInstancesRequest) (response *DescribeDiagDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDiagDBInstancesRequest()
    }
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeHealthScore(request *DescribeHealthScoreRequest) (response *DescribeHealthScoreResponse, err error) {
    if request == nil {
        request = NewDescribeHealthScoreRequest()
    }
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeMailProfile(request *DescribeMailProfileRequest) (response *DescribeMailProfileResponse, err error) {
    if request == nil {
        request = NewDescribeMailProfileRequest()
    }
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeMySqlProcessList(request *DescribeMySqlProcessListRequest) (response *DescribeMySqlProcessListResponse, err error) {
    if request == nil {
        request = NewDescribeMySqlProcessListRequest()
    }
    response = NewDescribeMySqlProcessListResponse()
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
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeSecurityAuditLogDownloadUrls(request *DescribeSecurityAuditLogDownloadUrlsRequest) (response *DescribeSecurityAuditLogDownloadUrlsResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityAuditLogDownloadUrlsRequest()
    }
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
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeSecurityAuditLogExportTasks(request *DescribeSecurityAuditLogExportTasksRequest) (response *DescribeSecurityAuditLogExportTasksResponse, err error) {
    if request == nil {
        request = NewDescribeSecurityAuditLogExportTasksRequest()
    }
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
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeSlowLogTimeSeriesStats(request *DescribeSlowLogTimeSeriesStatsRequest) (response *DescribeSlowLogTimeSeriesStatsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogTimeSeriesStatsRequest()
    }
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
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeSlowLogTopSqls(request *DescribeSlowLogTopSqlsRequest) (response *DescribeSlowLogTopSqlsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogTopSqlsRequest()
    }
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
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeSlowLogUserHostStats(request *DescribeSlowLogUserHostStatsRequest) (response *DescribeSlowLogUserHostStatsResponse, err error) {
    if request == nil {
        request = NewDescribeSlowLogUserHostStatsRequest()
    }
    response = NewDescribeSlowLogUserHostStatsResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeTopSpaceSchemaTimeSeries(request *DescribeTopSpaceSchemaTimeSeriesRequest) (response *DescribeTopSpaceSchemaTimeSeriesResponse, err error) {
    if request == nil {
        request = NewDescribeTopSpaceSchemaTimeSeriesRequest()
    }
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeTopSpaceSchemas(request *DescribeTopSpaceSchemasRequest) (response *DescribeTopSpaceSchemasResponse, err error) {
    if request == nil {
        request = NewDescribeTopSpaceSchemasRequest()
    }
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
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTopSpaceTableTimeSeries(request *DescribeTopSpaceTableTimeSeriesRequest) (response *DescribeTopSpaceTableTimeSeriesResponse, err error) {
    if request == nil {
        request = NewDescribeTopSpaceTableTimeSeriesRequest()
    }
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
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTopSpaceTables(request *DescribeTopSpaceTablesRequest) (response *DescribeTopSpaceTablesResponse, err error) {
    if request == nil {
        request = NewDescribeTopSpaceTablesRequest()
    }
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
// 获取SQL优化建议。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_USERHASNOSTRATEGY = "OperationDenied.UserHasNoStrategy"
func (c *Client) DescribeUserSqlAdvice(request *DescribeUserSqlAdviceRequest) (response *DescribeUserSqlAdviceResponse, err error) {
    if request == nil {
        request = NewDescribeUserSqlAdviceRequest()
    }
    response = NewDescribeUserSqlAdviceResponse()
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyDiagDBInstanceConf(request *ModifyDiagDBInstanceConfRequest) (response *ModifyDiagDBInstanceConfResponse, err error) {
    if request == nil {
        request = NewModifyDiagDBInstanceConfRequest()
    }
    response = NewModifyDiagDBInstanceConfResponse()
    err = c.Send(request, response)
    return
}
