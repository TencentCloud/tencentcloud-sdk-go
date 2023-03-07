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

package v20180330

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-03-30"

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


func NewActivateSubscribeRequest() (request *ActivateSubscribeRequest) {
    request = &ActivateSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ActivateSubscribe")
    
    
    return
}

func NewActivateSubscribeResponse() (response *ActivateSubscribeResponse) {
    response = &ActivateSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ActivateSubscribe
// 本接口用于配置数据订阅，只有在未配置状态的订阅实例才能调用此接口。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ActivateSubscribe(request *ActivateSubscribeRequest) (response *ActivateSubscribeResponse, err error) {
    return c.ActivateSubscribeWithContext(context.Background(), request)
}

// ActivateSubscribe
// 本接口用于配置数据订阅，只有在未配置状态的订阅实例才能调用此接口。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ActivateSubscribeWithContext(ctx context.Context, request *ActivateSubscribeRequest) (response *ActivateSubscribeResponse, err error) {
    if request == nil {
        request = NewActivateSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ActivateSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewActivateSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewCompleteMigrateJobRequest() (request *CompleteMigrateJobRequest) {
    request = &CompleteMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CompleteMigrateJob")
    
    
    return
}

func NewCompleteMigrateJobResponse() (response *CompleteMigrateJobResponse) {
    response = &CompleteMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CompleteMigrateJob
// 本接口（CompleteMigrateJob）用于完成数据迁移任务。
//
// 选择采用增量迁移方式的任务, 需要在迁移进度进入准备完成阶段后, 调用本接口, 停止迁移增量数据。
//
// 通过DescribeMigrateJobs接口查询到任务的状态为准备完成（status=8）时，此时可以调用本接口完成迁移任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_CELERYERROR = "InternalError.CeleryError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) CompleteMigrateJob(request *CompleteMigrateJobRequest) (response *CompleteMigrateJobResponse, err error) {
    return c.CompleteMigrateJobWithContext(context.Background(), request)
}

// CompleteMigrateJob
// 本接口（CompleteMigrateJob）用于完成数据迁移任务。
//
// 选择采用增量迁移方式的任务, 需要在迁移进度进入准备完成阶段后, 调用本接口, 停止迁移增量数据。
//
// 通过DescribeMigrateJobs接口查询到任务的状态为准备完成（status=8）时，此时可以调用本接口完成迁移任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_CELERYERROR = "InternalError.CeleryError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) CompleteMigrateJobWithContext(ctx context.Context, request *CompleteMigrateJobRequest) (response *CompleteMigrateJobResponse, err error) {
    if request == nil {
        request = NewCompleteMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CompleteMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCompleteMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMigrateCheckJobRequest() (request *CreateMigrateCheckJobRequest) {
    request = &CreateMigrateCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CreateMigrateCheckJob")
    
    
    return
}

func NewCreateMigrateCheckJobResponse() (response *CreateMigrateCheckJobResponse) {
    response = &CreateMigrateCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMigrateCheckJob
// 创建校验迁移任务
//
// 在开始迁移前, 必须调用本接口创建校验, 且校验成功后才能开始迁移. 校验的结果可以通过DescribeMigrateCheckJob查看.
//
// 校验成功后,迁移任务若有修改, 则必须重新创建校验并通过后, 才能开始迁移.
//
// 
//
// 如果是金融区链路, 请使用域名: https://dts.ap-shenzhen-fsi.tencentcloudapi.com
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_PROXYERROR = "FailedOperation.ProxyError"
//  INTERNALERROR_ADDTASKERROR = "InternalError.AddTaskError"
//  INTERNALERROR_CELERYERROR = "InternalError.CeleryError"
//  INTERNALERROR_CGWSYSTEMERROR = "InternalError.CgwSystemError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_LOCKERROR = "InternalError.LockError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDIPADDRESS = "InvalidParameter.InvalidIpAddress"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNSUPPORTEDOPERATION_INTRANETUSERNOTTAGGEDERROR = "UnsupportedOperation.IntraNetUserNotTaggedError"
func (c *Client) CreateMigrateCheckJob(request *CreateMigrateCheckJobRequest) (response *CreateMigrateCheckJobResponse, err error) {
    return c.CreateMigrateCheckJobWithContext(context.Background(), request)
}

// CreateMigrateCheckJob
// 创建校验迁移任务
//
// 在开始迁移前, 必须调用本接口创建校验, 且校验成功后才能开始迁移. 校验的结果可以通过DescribeMigrateCheckJob查看.
//
// 校验成功后,迁移任务若有修改, 则必须重新创建校验并通过后, 才能开始迁移.
//
// 
//
// 如果是金融区链路, 请使用域名: https://dts.ap-shenzhen-fsi.tencentcloudapi.com
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_PROXYERROR = "FailedOperation.ProxyError"
//  INTERNALERROR_ADDTASKERROR = "InternalError.AddTaskError"
//  INTERNALERROR_CELERYERROR = "InternalError.CeleryError"
//  INTERNALERROR_CGWSYSTEMERROR = "InternalError.CgwSystemError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_LOCKERROR = "InternalError.LockError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDIPADDRESS = "InvalidParameter.InvalidIpAddress"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNSUPPORTEDOPERATION_INTRANETUSERNOTTAGGEDERROR = "UnsupportedOperation.IntraNetUserNotTaggedError"
func (c *Client) CreateMigrateCheckJobWithContext(ctx context.Context, request *CreateMigrateCheckJobRequest) (response *CreateMigrateCheckJobResponse, err error) {
    if request == nil {
        request = NewCreateMigrateCheckJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMigrateCheckJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMigrateCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMigrateJobRequest() (request *CreateMigrateJobRequest) {
    request = &CreateMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CreateMigrateJob")
    
    
    return
}

func NewCreateMigrateJobResponse() (response *CreateMigrateJobResponse) {
    response = &CreateMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMigrateJob
// 本接口（CreateMigrateJob）用于创建数据迁移任务。
//
// 
//
// 如果是金融区链路, 请使用域名: dts.ap-shenzhen-fsi.tencentcloudapi.com
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALINNERCOMMONERROR = "InternalError.InternalInnerCommonError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameter.BizInvalidParameterValueError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameterValue.BizInvalidParameterValueError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOTENOUGHMONEYERROR = "OperationDenied.NotEnoughMoneyError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) CreateMigrateJob(request *CreateMigrateJobRequest) (response *CreateMigrateJobResponse, err error) {
    return c.CreateMigrateJobWithContext(context.Background(), request)
}

// CreateMigrateJob
// 本接口（CreateMigrateJob）用于创建数据迁移任务。
//
// 
//
// 如果是金融区链路, 请使用域名: dts.ap-shenzhen-fsi.tencentcloudapi.com
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_INTERNALINNERCOMMONERROR = "InternalError.InternalInnerCommonError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameter.BizInvalidParameterValueError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameterValue.BizInvalidParameterValueError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOTENOUGHMONEYERROR = "OperationDenied.NotEnoughMoneyError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) CreateMigrateJobWithContext(ctx context.Context, request *CreateMigrateJobRequest) (response *CreateMigrateJobResponse, err error) {
    if request == nil {
        request = NewCreateMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSubscribeRequest() (request *CreateSubscribeRequest) {
    request = &CreateSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CreateSubscribe")
    
    
    return
}

func NewCreateSubscribeResponse() (response *CreateSubscribeResponse) {
    response = &CreateSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSubscribe
// 本接口(CreateSubscribe)用于创建一个数据订阅实例。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALTRADEERROR = "InternalError.InternalTradeError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
func (c *Client) CreateSubscribe(request *CreateSubscribeRequest) (response *CreateSubscribeResponse, err error) {
    return c.CreateSubscribeWithContext(context.Background(), request)
}

// CreateSubscribe
// 本接口(CreateSubscribe)用于创建一个数据订阅实例。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALTRADEERROR = "InternalError.InternalTradeError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
func (c *Client) CreateSubscribeWithContext(ctx context.Context, request *CreateSubscribeRequest) (response *CreateSubscribeResponse, err error) {
    if request == nil {
        request = NewCreateSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMigrateJobRequest() (request *DeleteMigrateJobRequest) {
    request = &DeleteMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DeleteMigrateJob")
    
    
    return
}

func NewDeleteMigrateJobResponse() (response *DeleteMigrateJobResponse) {
    response = &DeleteMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMigrateJob
// 本接口（DeleteMigrationJob）用于删除数据迁移任务。当通过DescribeMigrateJobs接口查询到任务的状态为：检验中（status=3）、运行中（status=7）、准备完成（status=8）、撤销中（status=11）或者完成中（status=12）时，不允许删除任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CELERYERROR = "InternalError.CeleryError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_MIGRATESERVICESUPPORTERROR = "OperationDenied.MigrateServiceSupportError"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
func (c *Client) DeleteMigrateJob(request *DeleteMigrateJobRequest) (response *DeleteMigrateJobResponse, err error) {
    return c.DeleteMigrateJobWithContext(context.Background(), request)
}

// DeleteMigrateJob
// 本接口（DeleteMigrationJob）用于删除数据迁移任务。当通过DescribeMigrateJobs接口查询到任务的状态为：检验中（status=3）、运行中（status=7）、准备完成（status=8）、撤销中（status=11）或者完成中（status=12）时，不允许删除任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CELERYERROR = "InternalError.CeleryError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_MIGRATESERVICESUPPORTERROR = "OperationDenied.MigrateServiceSupportError"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
func (c *Client) DeleteMigrateJobWithContext(ctx context.Context, request *DeleteMigrateJobRequest) (response *DeleteMigrateJobResponse, err error) {
    if request == nil {
        request = NewDeleteMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAsyncRequestInfoRequest() (request *DescribeAsyncRequestInfoRequest) {
    request = &DescribeAsyncRequestInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeAsyncRequestInfo")
    
    
    return
}

func NewDescribeAsyncRequestInfoResponse() (response *DescribeAsyncRequestInfoResponse) {
    response = &DescribeAsyncRequestInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAsyncRequestInfo
// 本接口（DescribeAsyncRequestInfo）用于查询任务执行结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAsyncRequestInfo(request *DescribeAsyncRequestInfoRequest) (response *DescribeAsyncRequestInfoResponse, err error) {
    return c.DescribeAsyncRequestInfoWithContext(context.Background(), request)
}

// DescribeAsyncRequestInfo
// 本接口（DescribeAsyncRequestInfo）用于查询任务执行结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAsyncRequestInfoWithContext(ctx context.Context, request *DescribeAsyncRequestInfoRequest) (response *DescribeAsyncRequestInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAsyncRequestInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAsyncRequestInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAsyncRequestInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrateCheckJobRequest() (request *DescribeMigrateCheckJobRequest) {
    request = &DescribeMigrateCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrateCheckJob")
    
    
    return
}

func NewDescribeMigrateCheckJobResponse() (response *DescribeMigrateCheckJobResponse) {
    response = &DescribeMigrateCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMigrateCheckJob
// 本接口用于创建校验后,获取校验的结果. 能查询到当前校验的状态和进度. 
//
// 若通过校验, 则可调用'StartMigrateJob' 开始迁移.
//
// 若未通过校验, 则能查询到校验失败的原因. 请按照报错, 通过'ModifyMigrateJob'修改迁移配置或是调整源/目标实例的相关参数.
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) DescribeMigrateCheckJob(request *DescribeMigrateCheckJobRequest) (response *DescribeMigrateCheckJobResponse, err error) {
    return c.DescribeMigrateCheckJobWithContext(context.Background(), request)
}

// DescribeMigrateCheckJob
// 本接口用于创建校验后,获取校验的结果. 能查询到当前校验的状态和进度. 
//
// 若通过校验, 则可调用'StartMigrateJob' 开始迁移.
//
// 若未通过校验, 则能查询到校验失败的原因. 请按照报错, 通过'ModifyMigrateJob'修改迁移配置或是调整源/目标实例的相关参数.
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) DescribeMigrateCheckJobWithContext(ctx context.Context, request *DescribeMigrateCheckJobRequest) (response *DescribeMigrateCheckJobResponse, err error) {
    if request == nil {
        request = NewDescribeMigrateCheckJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrateCheckJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigrateCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrateJobsRequest() (request *DescribeMigrateJobsRequest) {
    request = &DescribeMigrateJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrateJobs")
    
    
    return
}

func NewDescribeMigrateJobsResponse() (response *DescribeMigrateJobsResponse) {
    response = &DescribeMigrateJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMigrateJobs
// 查询数据迁移任务.
//
// 如果是金融区链路, 请使用域名: https://dts.ap-shenzhen-fsi.tencentcloudapi.com
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameterValue.BizInvalidParameterValueError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) DescribeMigrateJobs(request *DescribeMigrateJobsRequest) (response *DescribeMigrateJobsResponse, err error) {
    return c.DescribeMigrateJobsWithContext(context.Background(), request)
}

// DescribeMigrateJobs
// 查询数据迁移任务.
//
// 如果是金融区链路, 请使用域名: https://dts.ap-shenzhen-fsi.tencentcloudapi.com
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameterValue.BizInvalidParameterValueError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) DescribeMigrateJobsWithContext(ctx context.Context, request *DescribeMigrateJobsRequest) (response *DescribeMigrateJobsResponse, err error) {
    if request == nil {
        request = NewDescribeMigrateJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrateJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigrateJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionConfRequest() (request *DescribeRegionConfRequest) {
    request = &DescribeRegionConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeRegionConf")
    
    
    return
}

func NewDescribeRegionConfResponse() (response *DescribeRegionConfResponse) {
    response = &DescribeRegionConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRegionConf
// 本接口（DescribeRegionConf）用于查询可售卖订阅实例的地域
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegionConf(request *DescribeRegionConfRequest) (response *DescribeRegionConfResponse, err error) {
    return c.DescribeRegionConfWithContext(context.Background(), request)
}

// DescribeRegionConf
// 本接口（DescribeRegionConf）用于查询可售卖订阅实例的地域
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRegionConfWithContext(ctx context.Context, request *DescribeRegionConfRequest) (response *DescribeRegionConfResponse, err error) {
    if request == nil {
        request = NewDescribeRegionConfRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegionConf require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegionConfResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscribeConfRequest() (request *DescribeSubscribeConfRequest) {
    request = &DescribeSubscribeConfRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribeConf")
    
    
    return
}

func NewDescribeSubscribeConfResponse() (response *DescribeSubscribeConfResponse) {
    response = &DescribeSubscribeConfResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSubscribeConf
// 本接口（DescribeSubscribeConf）用于查询订阅实例配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) DescribeSubscribeConf(request *DescribeSubscribeConfRequest) (response *DescribeSubscribeConfResponse, err error) {
    return c.DescribeSubscribeConfWithContext(context.Background(), request)
}

// DescribeSubscribeConf
// 本接口（DescribeSubscribeConf）用于查询订阅实例配置
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) DescribeSubscribeConfWithContext(ctx context.Context, request *DescribeSubscribeConfRequest) (response *DescribeSubscribeConfResponse, err error) {
    if request == nil {
        request = NewDescribeSubscribeConfRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubscribeConf require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubscribeConfResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscribesRequest() (request *DescribeSubscribesRequest) {
    request = &DescribeSubscribesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribes")
    
    
    return
}

func NewDescribeSubscribesResponse() (response *DescribeSubscribesResponse) {
    response = &DescribeSubscribesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSubscribes
// 本接口(DescribeSubscribes)获取数据订阅实例信息列表，默认分页，每次返回20条
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeSubscribes(request *DescribeSubscribesRequest) (response *DescribeSubscribesResponse, err error) {
    return c.DescribeSubscribesWithContext(context.Background(), request)
}

// DescribeSubscribes
// 本接口(DescribeSubscribes)获取数据订阅实例信息列表，默认分页，每次返回20条
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeSubscribesWithContext(ctx context.Context, request *DescribeSubscribesRequest) (response *DescribeSubscribesResponse, err error) {
    if request == nil {
        request = NewDescribeSubscribesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubscribes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubscribesResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateSubscribeRequest() (request *IsolateSubscribeRequest) {
    request = &IsolateSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "IsolateSubscribe")
    
    
    return
}

func NewIsolateSubscribeResponse() (response *IsolateSubscribeResponse) {
    response = &IsolateSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// IsolateSubscribe
// 本接口（IsolateSubscribe）用于隔离小时计费的订阅实例。调用后，订阅实例将不能使用，同时停止计费。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameter.BizInvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) IsolateSubscribe(request *IsolateSubscribeRequest) (response *IsolateSubscribeResponse, err error) {
    return c.IsolateSubscribeWithContext(context.Background(), request)
}

// IsolateSubscribe
// 本接口（IsolateSubscribe）用于隔离小时计费的订阅实例。调用后，订阅实例将不能使用，同时停止计费。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameter.BizInvalidParameterValueError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) IsolateSubscribeWithContext(ctx context.Context, request *IsolateSubscribeRequest) (response *IsolateSubscribeResponse, err error) {
    if request == nil {
        request = NewIsolateSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMigrateJobRequest() (request *ModifyMigrateJobRequest) {
    request = &ModifyMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifyMigrateJob")
    
    
    return
}

func NewModifyMigrateJobResponse() (response *ModifyMigrateJobResponse) {
    response = &ModifyMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMigrateJob
// 本接口（ModifyMigrateJob）用于修改数据迁移任务。
//
// 当迁移任务处于下述状态时，允许调用本接口修改迁移任务：迁移创建中（status=1）、 校验成功(status=4)、校验失败(status=5)、迁移失败(status=10)。但源实例、目标实例类型和目标实例地域不允许修改。
//
// 
//
// 如果是金融区链路, 请使用域名: dts.ap-shenzhen-fsi.tencentcloudapi.com
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameterValue.BizInvalidParameterValueError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) ModifyMigrateJob(request *ModifyMigrateJobRequest) (response *ModifyMigrateJobResponse, err error) {
    return c.ModifyMigrateJobWithContext(context.Background(), request)
}

// ModifyMigrateJob
// 本接口（ModifyMigrateJob）用于修改数据迁移任务。
//
// 当迁移任务处于下述状态时，允许调用本接口修改迁移任务：迁移创建中（status=1）、 校验成功(status=4)、校验失败(status=5)、迁移失败(status=10)。但源实例、目标实例类型和目标实例地域不允许修改。
//
// 
//
// 如果是金融区链路, 请使用域名: dts.ap-shenzhen-fsi.tencentcloudapi.com
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BIZINVALIDPARAMETERVALUEERROR = "InvalidParameterValue.BizInvalidParameterValueError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) ModifyMigrateJobWithContext(ctx context.Context, request *ModifyMigrateJobRequest) (response *ModifyMigrateJobResponse, err error) {
    if request == nil {
        request = NewModifyMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubscribeAutoRenewFlagRequest() (request *ModifySubscribeAutoRenewFlagRequest) {
    request = &ModifySubscribeAutoRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeAutoRenewFlag")
    
    
    return
}

func NewModifySubscribeAutoRenewFlagResponse() (response *ModifySubscribeAutoRenewFlagResponse) {
    response = &ModifySubscribeAutoRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySubscribeAutoRenewFlag
// 修改订阅实例自动续费标识
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySubscribeAutoRenewFlag(request *ModifySubscribeAutoRenewFlagRequest) (response *ModifySubscribeAutoRenewFlagResponse, err error) {
    return c.ModifySubscribeAutoRenewFlagWithContext(context.Background(), request)
}

// ModifySubscribeAutoRenewFlag
// 修改订阅实例自动续费标识
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySubscribeAutoRenewFlagWithContext(ctx context.Context, request *ModifySubscribeAutoRenewFlagRequest) (response *ModifySubscribeAutoRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifySubscribeAutoRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubscribeAutoRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySubscribeAutoRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubscribeConsumeTimeRequest() (request *ModifySubscribeConsumeTimeRequest) {
    request = &ModifySubscribeConsumeTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeConsumeTime")
    
    
    return
}

func NewModifySubscribeConsumeTimeResponse() (response *ModifySubscribeConsumeTimeResponse) {
    response = &ModifySubscribeConsumeTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySubscribeConsumeTime
// 本接口(ModifySubscribeConsumeTime)用于修改数据订阅通道的消费时间点
//
// 可能返回的错误码:
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  FAILEDOPERATION_STATUSINCONFLICTERROR = "FailedOperation.StatusInConflictError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySubscribeConsumeTime(request *ModifySubscribeConsumeTimeRequest) (response *ModifySubscribeConsumeTimeResponse, err error) {
    return c.ModifySubscribeConsumeTimeWithContext(context.Background(), request)
}

// ModifySubscribeConsumeTime
// 本接口(ModifySubscribeConsumeTime)用于修改数据订阅通道的消费时间点
//
// 可能返回的错误码:
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  FAILEDOPERATION_STATUSINCONFLICTERROR = "FailedOperation.StatusInConflictError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySubscribeConsumeTimeWithContext(ctx context.Context, request *ModifySubscribeConsumeTimeRequest) (response *ModifySubscribeConsumeTimeResponse, err error) {
    if request == nil {
        request = NewModifySubscribeConsumeTimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubscribeConsumeTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySubscribeConsumeTimeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubscribeNameRequest() (request *ModifySubscribeNameRequest) {
    request = &ModifySubscribeNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeName")
    
    
    return
}

func NewModifySubscribeNameResponse() (response *ModifySubscribeNameResponse) {
    response = &ModifySubscribeNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySubscribeName
// 本接口(ModifySubscribeName)用于修改数据订阅实例的名称
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySubscribeName(request *ModifySubscribeNameRequest) (response *ModifySubscribeNameResponse, err error) {
    return c.ModifySubscribeNameWithContext(context.Background(), request)
}

// ModifySubscribeName
// 本接口(ModifySubscribeName)用于修改数据订阅实例的名称
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySubscribeNameWithContext(ctx context.Context, request *ModifySubscribeNameRequest) (response *ModifySubscribeNameResponse, err error) {
    if request == nil {
        request = NewModifySubscribeNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubscribeName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySubscribeNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubscribeObjectsRequest() (request *ModifySubscribeObjectsRequest) {
    request = &ModifySubscribeObjectsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeObjects")
    
    
    return
}

func NewModifySubscribeObjectsResponse() (response *ModifySubscribeObjectsResponse) {
    response = &ModifySubscribeObjectsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySubscribeObjects
// 本接口(ModifySubscribeObjects)用于修改数据订阅通道的订阅规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) ModifySubscribeObjects(request *ModifySubscribeObjectsRequest) (response *ModifySubscribeObjectsResponse, err error) {
    return c.ModifySubscribeObjectsWithContext(context.Background(), request)
}

// ModifySubscribeObjects
// 本接口(ModifySubscribeObjects)用于修改数据订阅通道的订阅规则
//
// 可能返回的错误码:
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
func (c *Client) ModifySubscribeObjectsWithContext(ctx context.Context, request *ModifySubscribeObjectsRequest) (response *ModifySubscribeObjectsResponse, err error) {
    if request == nil {
        request = NewModifySubscribeObjectsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubscribeObjects require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySubscribeObjectsResponse()
    err = c.Send(request, response)
    return
}

func NewModifySubscribeVipVportRequest() (request *ModifySubscribeVipVportRequest) {
    request = &ModifySubscribeVipVportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifySubscribeVipVport")
    
    
    return
}

func NewModifySubscribeVipVportResponse() (response *ModifySubscribeVipVportResponse) {
    response = &ModifySubscribeVipVportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySubscribeVipVport
// 本接口(ModifySubscribeVipVport)用于修改数据订阅实例的IP和端口号
//
// 可能返回的错误码:
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySubscribeVipVport(request *ModifySubscribeVipVportRequest) (response *ModifySubscribeVipVportResponse, err error) {
    return c.ModifySubscribeVipVportWithContext(context.Background(), request)
}

// ModifySubscribeVipVport
// 本接口(ModifySubscribeVipVport)用于修改数据订阅实例的IP和端口号
//
// 可能返回的错误码:
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifySubscribeVipVportWithContext(ctx context.Context, request *ModifySubscribeVipVportRequest) (response *ModifySubscribeVipVportResponse, err error) {
    if request == nil {
        request = NewModifySubscribeVipVportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySubscribeVipVport require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySubscribeVipVportResponse()
    err = c.Send(request, response)
    return
}

func NewOfflineIsolatedSubscribeRequest() (request *OfflineIsolatedSubscribeRequest) {
    request = &OfflineIsolatedSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "OfflineIsolatedSubscribe")
    
    
    return
}

func NewOfflineIsolatedSubscribeResponse() (response *OfflineIsolatedSubscribeResponse) {
    response = &OfflineIsolatedSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// OfflineIsolatedSubscribe
// 本接口（OfflineIsolatedSubscribe）用于下线已隔离的数据订阅实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  FAILEDOPERATION_STATUSINCONFLICTERROR = "FailedOperation.StatusInConflictError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) OfflineIsolatedSubscribe(request *OfflineIsolatedSubscribeRequest) (response *OfflineIsolatedSubscribeResponse, err error) {
    return c.OfflineIsolatedSubscribeWithContext(context.Background(), request)
}

// OfflineIsolatedSubscribe
// 本接口（OfflineIsolatedSubscribe）用于下线已隔离的数据订阅实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  FAILEDOPERATION_STATUSINCONFLICTERROR = "FailedOperation.StatusInConflictError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) OfflineIsolatedSubscribeWithContext(ctx context.Context, request *OfflineIsolatedSubscribeRequest) (response *OfflineIsolatedSubscribeResponse, err error) {
    if request == nil {
        request = NewOfflineIsolatedSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("OfflineIsolatedSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewOfflineIsolatedSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewResetSubscribeRequest() (request *ResetSubscribeRequest) {
    request = &ResetSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ResetSubscribe")
    
    
    return
}

func NewResetSubscribeResponse() (response *ResetSubscribeResponse) {
    response = &ResetSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResetSubscribe
// 本接口(ResetSubscribe)用于重置数据订阅实例，已经激活的数据订阅实例，重置后可以使用ActivateSubscribe接口绑定其他的数据库实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ResetSubscribe(request *ResetSubscribeRequest) (response *ResetSubscribeResponse, err error) {
    return c.ResetSubscribeWithContext(context.Background(), request)
}

// ResetSubscribe
// 本接口(ResetSubscribe)用于重置数据订阅实例，已经激活的数据订阅实例，重置后可以使用ActivateSubscribe接口绑定其他的数据库实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ResetSubscribeWithContext(ctx context.Context, request *ResetSubscribeRequest) (response *ResetSubscribeResponse, err error) {
    if request == nil {
        request = NewResetSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewStartMigrateJobRequest() (request *StartMigrateJobRequest) {
    request = &StartMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "StartMigrateJob")
    
    
    return
}

func NewStartMigrateJobResponse() (response *StartMigrateJobResponse) {
    response = &StartMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartMigrateJob
// 本接口（StartMigrationJob）用于启动迁移任务。非定时迁移任务会在调用后立即开始迁移，定时任务则会开始倒计时。
//
// 调用此接口前，请务必先使用CreateMigrateCheckJob校验数据迁移任务，并通过DescribeMigrateJobs接口查询到任务状态为校验通过（status=4）时，才能启动数据迁移任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STARTJOBFAILED = "FailedOperation.StartJobFailed"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_ADDTASKERROR = "InternalError.AddTaskError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALINNERCOMMONERROR = "InternalError.InternalInnerCommonError"
//  INTERNALERROR_LOCKERROR = "InternalError.LockError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) StartMigrateJob(request *StartMigrateJobRequest) (response *StartMigrateJobResponse, err error) {
    return c.StartMigrateJobWithContext(context.Background(), request)
}

// StartMigrateJob
// 本接口（StartMigrationJob）用于启动迁移任务。非定时迁移任务会在调用后立即开始迁移，定时任务则会开始倒计时。
//
// 调用此接口前，请务必先使用CreateMigrateCheckJob校验数据迁移任务，并通过DescribeMigrateJobs接口查询到任务状态为校验通过（status=4）时，才能启动数据迁移任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STARTJOBFAILED = "FailedOperation.StartJobFailed"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_ADDTASKERROR = "InternalError.AddTaskError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALINNERCOMMONERROR = "InternalError.InternalInnerCommonError"
//  INTERNALERROR_LOCKERROR = "InternalError.LockError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) StartMigrateJobWithContext(ctx context.Context, request *StartMigrateJobRequest) (response *StartMigrateJobResponse, err error) {
    if request == nil {
        request = NewStartMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewStopMigrateJobRequest() (request *StopMigrateJobRequest) {
    request = &StopMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "StopMigrateJob")
    
    
    return
}

func NewStopMigrateJobResponse() (response *StopMigrateJobResponse) {
    response = &StopMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopMigrateJob
// 本接口（StopMigrateJob）用于撤销数据迁移任务。
//
// 在迁移过程中允许调用该接口撤销迁移, 撤销迁移的任务会失败。通过DescribeMigrateJobs接口查询到任务状态为运行中（status=7）或准备完成（status=8）时，才能撤销数据迁移任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) StopMigrateJob(request *StopMigrateJobRequest) (response *StopMigrateJobResponse, err error) {
    return c.StopMigrateJobWithContext(context.Background(), request)
}

// StopMigrateJob
// 本接口（StopMigrateJob）用于撤销数据迁移任务。
//
// 在迁移过程中允许调用该接口撤销迁移, 撤销迁移的任务会失败。通过DescribeMigrateJobs接口查询到任务状态为运行中（status=7）或准备完成（status=8）时，才能撤销数据迁移任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
func (c *Client) StopMigrateJobWithContext(ctx context.Context, request *StopMigrateJobRequest) (response *StopMigrateJobResponse, err error) {
    if request == nil {
        request = NewStopMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopMigrateJobResponse()
    err = c.Send(request, response)
    return
}
