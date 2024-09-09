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

package v20211206

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-12-06"

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
// 通过DescribeMigrationJobs接口查询到任务的状态为准备完成（Status="readyComplete"）时，此时可以调用本接口完成迁移任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_CELERYERROR = "InternalError.CeleryError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CompleteMigrateJob(request *CompleteMigrateJobRequest) (response *CompleteMigrateJobResponse, err error) {
    return c.CompleteMigrateJobWithContext(context.Background(), request)
}

// CompleteMigrateJob
// 本接口（CompleteMigrateJob）用于完成数据迁移任务。
//
// 选择采用增量迁移方式的任务, 需要在迁移进度进入准备完成阶段后, 调用本接口, 停止迁移增量数据。
//
// 通过DescribeMigrationJobs接口查询到任务的状态为准备完成（Status="readyComplete"）时，此时可以调用本接口完成迁移任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_CELERYERROR = "InternalError.CeleryError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
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

func NewConfigureSubscribeJobRequest() (request *ConfigureSubscribeJobRequest) {
    request = &ConfigureSubscribeJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ConfigureSubscribeJob")
    
    
    return
}

func NewConfigureSubscribeJobResponse() (response *ConfigureSubscribeJobResponse) {
    response = &ConfigureSubscribeJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ConfigureSubscribeJob
// 本接口(ConfigureSubscribeJob)用于配置数据订阅实例。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ConfigureSubscribeJob(request *ConfigureSubscribeJobRequest) (response *ConfigureSubscribeJobResponse, err error) {
    return c.ConfigureSubscribeJobWithContext(context.Background(), request)
}

// ConfigureSubscribeJob
// 本接口(ConfigureSubscribeJob)用于配置数据订阅实例。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ConfigureSubscribeJobWithContext(ctx context.Context, request *ConfigureSubscribeJobRequest) (response *ConfigureSubscribeJobResponse, err error) {
    if request == nil {
        request = NewConfigureSubscribeJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ConfigureSubscribeJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewConfigureSubscribeJobResponse()
    err = c.Send(request, response)
    return
}

func NewConfigureSyncJobRequest() (request *ConfigureSyncJobRequest) {
    request = &ConfigureSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ConfigureSyncJob")
    
    
    return
}

func NewConfigureSyncJobResponse() (response *ConfigureSyncJobResponse) {
    response = &ConfigureSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ConfigureSyncJob
// 配置一个同步任务
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) ConfigureSyncJob(request *ConfigureSyncJobRequest) (response *ConfigureSyncJobResponse, err error) {
    return c.ConfigureSyncJobWithContext(context.Background(), request)
}

// ConfigureSyncJob
// 配置一个同步任务
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) ConfigureSyncJobWithContext(ctx context.Context, request *ConfigureSyncJobRequest) (response *ConfigureSyncJobResponse, err error) {
    if request == nil {
        request = NewConfigureSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ConfigureSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewConfigureSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewContinueMigrateJobRequest() (request *ContinueMigrateJobRequest) {
    request = &ContinueMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ContinueMigrateJob")
    
    
    return
}

func NewContinueMigrateJobResponse() (response *ContinueMigrateJobResponse) {
    response = &ContinueMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ContinueMigrateJob
// 恢复一个暂停中的迁移任务。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) ContinueMigrateJob(request *ContinueMigrateJobRequest) (response *ContinueMigrateJobResponse, err error) {
    return c.ContinueMigrateJobWithContext(context.Background(), request)
}

// ContinueMigrateJob
// 恢复一个暂停中的迁移任务。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) ContinueMigrateJobWithContext(ctx context.Context, request *ContinueMigrateJobRequest) (response *ContinueMigrateJobResponse, err error) {
    if request == nil {
        request = NewContinueMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ContinueMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewContinueMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewContinueSyncJobRequest() (request *ContinueSyncJobRequest) {
    request = &ContinueSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ContinueSyncJob")
    
    
    return
}

func NewContinueSyncJobResponse() (response *ContinueSyncJobResponse) {
    response = &ContinueSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ContinueSyncJob
// 恢复处于已暂停状态的数据同步任务。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) ContinueSyncJob(request *ContinueSyncJobRequest) (response *ContinueSyncJobResponse, err error) {
    return c.ContinueSyncJobWithContext(context.Background(), request)
}

// ContinueSyncJob
// 恢复处于已暂停状态的数据同步任务。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) ContinueSyncJobWithContext(ctx context.Context, request *ContinueSyncJobRequest) (response *ContinueSyncJobResponse, err error) {
    if request == nil {
        request = NewContinueSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ContinueSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewContinueSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCheckSyncJobRequest() (request *CreateCheckSyncJobRequest) {
    request = &CreateCheckSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CreateCheckSyncJob")
    
    
    return
}

func NewCreateCheckSyncJobResponse() (response *CreateCheckSyncJobResponse) {
    response = &CreateCheckSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCheckSyncJob
// 校验同步任务，检查必要参数和周边配置。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INTRANETUSERNOTTAGGEDERROR = "UnsupportedOperation.IntraNetUserNotTaggedError"
//  UNSUPPORTEDOPERATION_INTRANETUSERNOTTAGGEDFORSYNCJOBERROR = "UnsupportedOperation.IntraNetUserNotTaggedForSyncJobError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateCheckSyncJob(request *CreateCheckSyncJobRequest) (response *CreateCheckSyncJobResponse, err error) {
    return c.CreateCheckSyncJobWithContext(context.Background(), request)
}

// CreateCheckSyncJob
// 校验同步任务，检查必要参数和周边配置。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INTRANETUSERNOTTAGGEDERROR = "UnsupportedOperation.IntraNetUserNotTaggedError"
//  UNSUPPORTEDOPERATION_INTRANETUSERNOTTAGGEDFORSYNCJOBERROR = "UnsupportedOperation.IntraNetUserNotTaggedForSyncJobError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateCheckSyncJobWithContext(ctx context.Context, request *CreateCheckSyncJobRequest) (response *CreateCheckSyncJobResponse, err error) {
    if request == nil {
        request = NewCreateCheckSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCheckSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCheckSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCompareTaskRequest() (request *CreateCompareTaskRequest) {
    request = &CreateCompareTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CreateCompareTask")
    
    
    return
}

func NewCreateCompareTaskResponse() (response *CreateCompareTaskResponse) {
    response = &CreateCompareTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCompareTask
// 本接口用于创建数据对比任务，创建成功后会返回数据对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9，创建成功后可通过StartCompare启动一致性校验任务
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateCompareTask(request *CreateCompareTaskRequest) (response *CreateCompareTaskResponse, err error) {
    return c.CreateCompareTaskWithContext(context.Background(), request)
}

// CreateCompareTask
// 本接口用于创建数据对比任务，创建成功后会返回数据对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9，创建成功后可通过StartCompare启动一致性校验任务
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateCompareTaskWithContext(ctx context.Context, request *CreateCompareTaskRequest) (response *CreateCompareTaskResponse, err error) {
    if request == nil {
        request = NewCreateCompareTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCompareTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCompareTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConsumerGroupRequest() (request *CreateConsumerGroupRequest) {
    request = &CreateConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CreateConsumerGroup")
    
    
    return
}

func NewCreateConsumerGroupResponse() (response *CreateConsumerGroupResponse) {
    response = &CreateConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateConsumerGroup
// 为订阅实例创建消费者组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateConsumerGroup(request *CreateConsumerGroupRequest) (response *CreateConsumerGroupResponse, err error) {
    return c.CreateConsumerGroupWithContext(context.Background(), request)
}

// CreateConsumerGroup
// 为订阅实例创建消费者组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateConsumerGroupWithContext(ctx context.Context, request *CreateConsumerGroupRequest) (response *CreateConsumerGroupResponse, err error) {
    if request == nil {
        request = NewCreateConsumerGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConsumerGroupResponse()
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
// 校验迁移任务，
//
// 在开始迁移前, 必须调用本接口创建校验迁移任务, 且校验成功后才能开始迁移. 校验的结果可以通过DescribeMigrationCheckJob查看，
//
// 校验成功后,迁移任务若有修改, 则必须重新校验并通过后, 才能开始迁移
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ADDTASKERROR = "InternalError.AddTaskError"
//  INTERNALERROR_CELERYERROR = "InternalError.CeleryError"
//  INTERNALERROR_CGWSYSTEMERROR = "InternalError.CgwSystemError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_LOCKERROR = "InternalError.LockError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_INTRANETUSERNOTTAGGEDERROR = "UnsupportedOperation.IntraNetUserNotTaggedError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateMigrateCheckJob(request *CreateMigrateCheckJobRequest) (response *CreateMigrateCheckJobResponse, err error) {
    return c.CreateMigrateCheckJobWithContext(context.Background(), request)
}

// CreateMigrateCheckJob
// 校验迁移任务，
//
// 在开始迁移前, 必须调用本接口创建校验迁移任务, 且校验成功后才能开始迁移. 校验的结果可以通过DescribeMigrationCheckJob查看，
//
// 校验成功后,迁移任务若有修改, 则必须重新校验并通过后, 才能开始迁移
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ADDTASKERROR = "InternalError.AddTaskError"
//  INTERNALERROR_CELERYERROR = "InternalError.CeleryError"
//  INTERNALERROR_CGWSYSTEMERROR = "InternalError.CgwSystemError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_LOCKERROR = "InternalError.LockError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNDEFINEDERROR = "InternalError.UndefinedError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTFOUND = "InvalidParameter.InstanceNotFound"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_INTRANETUSERNOTTAGGEDERROR = "UnsupportedOperation.IntraNetUserNotTaggedError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
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

func NewCreateMigrationServiceRequest() (request *CreateMigrationServiceRequest) {
    request = &CreateMigrationServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CreateMigrationService")
    
    
    return
}

func NewCreateMigrationServiceResponse() (response *CreateMigrationServiceResponse) {
    response = &CreateMigrationServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateMigrationService
// 购买迁移任务。购买成功后会返回随机生成的迁移任务id列表，也可以通过查询迁移任务任务列表接口`DescribeMigrationJobs`看到购买成功的实例Id。注意，一旦购买成功后源及目标数据库类型，源及目标实例地域不可修改。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateMigrationService(request *CreateMigrationServiceRequest) (response *CreateMigrationServiceResponse, err error) {
    return c.CreateMigrationServiceWithContext(context.Background(), request)
}

// CreateMigrationService
// 购买迁移任务。购买成功后会返回随机生成的迁移任务id列表，也可以通过查询迁移任务任务列表接口`DescribeMigrationJobs`看到购买成功的实例Id。注意，一旦购买成功后源及目标数据库类型，源及目标实例地域不可修改。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateMigrationServiceWithContext(ctx context.Context, request *CreateMigrationServiceRequest) (response *CreateMigrationServiceResponse, err error) {
    if request == nil {
        request = NewCreateMigrationServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMigrationService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateMigrationServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateModifyCheckSyncJobRequest() (request *CreateModifyCheckSyncJobRequest) {
    request = &CreateModifyCheckSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CreateModifyCheckSyncJob")
    
    
    return
}

func NewCreateModifyCheckSyncJobResponse() (response *CreateModifyCheckSyncJobResponse) {
    response = &CreateModifyCheckSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateModifyCheckSyncJob
// 在修改同步任务的配置后、通过该接口校验当前任务是否支持修改对象操作
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateModifyCheckSyncJob(request *CreateModifyCheckSyncJobRequest) (response *CreateModifyCheckSyncJobResponse, err error) {
    return c.CreateModifyCheckSyncJobWithContext(context.Background(), request)
}

// CreateModifyCheckSyncJob
// 在修改同步任务的配置后、通过该接口校验当前任务是否支持修改对象操作
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateModifyCheckSyncJobWithContext(ctx context.Context, request *CreateModifyCheckSyncJobRequest) (response *CreateModifyCheckSyncJobResponse, err error) {
    if request == nil {
        request = NewCreateModifyCheckSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateModifyCheckSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateModifyCheckSyncJobResponse()
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
// 本接口(CreateSubscribe)用于创建一个数据订阅任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_USERNOTAUTHED = "FailedOperation.UserNotAuthed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOTENOUGHMONEYERROR = "OperationDenied.NotEnoughMoneyError"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateSubscribe(request *CreateSubscribeRequest) (response *CreateSubscribeResponse, err error) {
    return c.CreateSubscribeWithContext(context.Background(), request)
}

// CreateSubscribe
// 本接口(CreateSubscribe)用于创建一个数据订阅任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_USERNOTAUTHED = "FailedOperation.UserNotAuthed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_NOTENOUGHMONEYERROR = "OperationDenied.NotEnoughMoneyError"
//  OPERATIONDENIED_OPERATIONDENIED = "OperationDenied.OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
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

func NewCreateSubscribeCheckJobRequest() (request *CreateSubscribeCheckJobRequest) {
    request = &CreateSubscribeCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CreateSubscribeCheckJob")
    
    
    return
}

func NewCreateSubscribeCheckJobResponse() (response *CreateSubscribeCheckJobResponse) {
    response = &CreateSubscribeCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSubscribeCheckJob
// 本接口(CreateSubscribeCheckJob)用于创建一个订阅校验任务。任务必须已经成功调用ConfigureSubscribeJob接口配置了所有的必要信息才能启动校验。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateSubscribeCheckJob(request *CreateSubscribeCheckJobRequest) (response *CreateSubscribeCheckJobResponse, err error) {
    return c.CreateSubscribeCheckJobWithContext(context.Background(), request)
}

// CreateSubscribeCheckJob
// 本接口(CreateSubscribeCheckJob)用于创建一个订阅校验任务。任务必须已经成功调用ConfigureSubscribeJob接口配置了所有的必要信息才能启动校验。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) CreateSubscribeCheckJobWithContext(ctx context.Context, request *CreateSubscribeCheckJobRequest) (response *CreateSubscribeCheckJobResponse, err error) {
    if request == nil {
        request = NewCreateSubscribeCheckJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSubscribeCheckJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSubscribeCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSyncJobRequest() (request *CreateSyncJobRequest) {
    request = &CreateSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "CreateSyncJob")
    
    
    return
}

func NewCreateSyncJobResponse() (response *CreateSyncJobResponse) {
    response = &CreateSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSyncJob
// 创建一个同步任务
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALINNERCOMMONERROR = "InternalError.InternalInnerCommonError"
//  INTERNALERROR_INTERNALTRADEERROR = "InternalError.InternalTradeError"
//  INTERNALERROR_NOTENOUGHMONEYERROR = "InternalError.NotEnoughMoneyError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_NOTENOUGHMONEYERROR = "OperationDenied.NotEnoughMoneyError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
func (c *Client) CreateSyncJob(request *CreateSyncJobRequest) (response *CreateSyncJobResponse, err error) {
    return c.CreateSyncJobWithContext(context.Background(), request)
}

// CreateSyncJob
// 创建一个同步任务
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALINNERCOMMONERROR = "InternalError.InternalInnerCommonError"
//  INTERNALERROR_INTERNALTRADEERROR = "InternalError.InternalTradeError"
//  INTERNALERROR_NOTENOUGHMONEYERROR = "InternalError.NotEnoughMoneyError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_NOTENOUGHMONEYERROR = "OperationDenied.NotEnoughMoneyError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
func (c *Client) CreateSyncJobWithContext(ctx context.Context, request *CreateSyncJobRequest) (response *CreateSyncJobResponse, err error) {
    if request == nil {
        request = NewCreateSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCompareTaskRequest() (request *DeleteCompareTaskRequest) {
    request = &DeleteCompareTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DeleteCompareTask")
    
    
    return
}

func NewDeleteCompareTaskResponse() (response *DeleteCompareTaskResponse) {
    response = &DeleteCompareTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCompareTask
// 删除一致性校验任务。当一致性校验任务状态为success、failed、canceled 时可以执行此操作。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DeleteCompareTask(request *DeleteCompareTaskRequest) (response *DeleteCompareTaskResponse, err error) {
    return c.DeleteCompareTaskWithContext(context.Background(), request)
}

// DeleteCompareTask
// 删除一致性校验任务。当一致性校验任务状态为success、failed、canceled 时可以执行此操作。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DeleteCompareTaskWithContext(ctx context.Context, request *DeleteCompareTaskRequest) (response *DeleteCompareTaskResponse, err error) {
    if request == nil {
        request = NewDeleteCompareTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCompareTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCompareTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConsumerGroupRequest() (request *DeleteConsumerGroupRequest) {
    request = &DeleteConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DeleteConsumerGroup")
    
    
    return
}

func NewDeleteConsumerGroupResponse() (response *DeleteConsumerGroupResponse) {
    response = &DeleteConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteConsumerGroup
// 本接口(DeleteConsumerGroup)用于删除一个订阅任务的消费组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DeleteConsumerGroup(request *DeleteConsumerGroupRequest) (response *DeleteConsumerGroupResponse, err error) {
    return c.DeleteConsumerGroupWithContext(context.Background(), request)
}

// DeleteConsumerGroup
// 本接口(DeleteConsumerGroup)用于删除一个订阅任务的消费组。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DeleteConsumerGroupWithContext(ctx context.Context, request *DeleteConsumerGroupRequest) (response *DeleteConsumerGroupResponse, err error) {
    if request == nil {
        request = NewDeleteConsumerGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCheckSyncJobResultRequest() (request *DescribeCheckSyncJobResultRequest) {
    request = &DescribeCheckSyncJobResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeCheckSyncJobResult")
    
    
    return
}

func NewDescribeCheckSyncJobResultResponse() (response *DescribeCheckSyncJobResultResponse) {
    response = &DescribeCheckSyncJobResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCheckSyncJobResult
// 查询同步校验任务结果，检查必要参数和周边配置
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) DescribeCheckSyncJobResult(request *DescribeCheckSyncJobResultRequest) (response *DescribeCheckSyncJobResultResponse, err error) {
    return c.DescribeCheckSyncJobResultWithContext(context.Background(), request)
}

// DescribeCheckSyncJobResult
// 查询同步校验任务结果，检查必要参数和周边配置
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) DescribeCheckSyncJobResultWithContext(ctx context.Context, request *DescribeCheckSyncJobResultRequest) (response *DescribeCheckSyncJobResultResponse, err error) {
    if request == nil {
        request = NewDescribeCheckSyncJobResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCheckSyncJobResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCheckSyncJobResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCompareReportRequest() (request *DescribeCompareReportRequest) {
    request = &DescribeCompareReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeCompareReport")
    
    
    return
}

func NewDescribeCompareReportResponse() (response *DescribeCompareReportResponse) {
    response = &DescribeCompareReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCompareReport
// 查询一致性校验任务详情
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeCompareReport(request *DescribeCompareReportRequest) (response *DescribeCompareReportResponse, err error) {
    return c.DescribeCompareReportWithContext(context.Background(), request)
}

// DescribeCompareReport
// 查询一致性校验任务详情
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeCompareReportWithContext(ctx context.Context, request *DescribeCompareReportRequest) (response *DescribeCompareReportResponse, err error) {
    if request == nil {
        request = NewDescribeCompareReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCompareReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCompareReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCompareTasksRequest() (request *DescribeCompareTasksRequest) {
    request = &DescribeCompareTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeCompareTasks")
    
    
    return
}

func NewDescribeCompareTasksResponse() (response *DescribeCompareTasksResponse) {
    response = &DescribeCompareTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCompareTasks
// 查询一致性校验任务列表，调用该接口后可通过接口`DescribeCompareTasks` 查询一致性校验任务列表来获得启动后的状态。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeCompareTasks(request *DescribeCompareTasksRequest) (response *DescribeCompareTasksResponse, err error) {
    return c.DescribeCompareTasksWithContext(context.Background(), request)
}

// DescribeCompareTasks
// 查询一致性校验任务列表，调用该接口后可通过接口`DescribeCompareTasks` 查询一致性校验任务列表来获得启动后的状态。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeCompareTasksWithContext(ctx context.Context, request *DescribeCompareTasksRequest) (response *DescribeCompareTasksResponse, err error) {
    if request == nil {
        request = NewDescribeCompareTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCompareTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCompareTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsumerGroupsRequest() (request *DescribeConsumerGroupsRequest) {
    request = &DescribeConsumerGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeConsumerGroups")
    
    
    return
}

func NewDescribeConsumerGroupsResponse() (response *DescribeConsumerGroupsResponse) {
    response = &DescribeConsumerGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeConsumerGroups
// 本接口(DescribeConsumerGroups)用于获取订阅实例配置的消费者组详情。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeConsumerGroups(request *DescribeConsumerGroupsRequest) (response *DescribeConsumerGroupsResponse, err error) {
    return c.DescribeConsumerGroupsWithContext(context.Background(), request)
}

// DescribeConsumerGroups
// 本接口(DescribeConsumerGroups)用于获取订阅实例配置的消费者组详情。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeConsumerGroupsWithContext(ctx context.Context, request *DescribeConsumerGroupsRequest) (response *DescribeConsumerGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeConsumerGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsumerGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsumerGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrateDBInstancesRequest() (request *DescribeMigrateDBInstancesRequest) {
    request = &DescribeMigrateDBInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrateDBInstances")
    
    
    return
}

func NewDescribeMigrateDBInstancesResponse() (response *DescribeMigrateDBInstancesResponse) {
    response = &DescribeMigrateDBInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMigrateDBInstances
// 本接口用于查询支持迁移的云数据库实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeMigrateDBInstances(request *DescribeMigrateDBInstancesRequest) (response *DescribeMigrateDBInstancesResponse, err error) {
    return c.DescribeMigrateDBInstancesWithContext(context.Background(), request)
}

// DescribeMigrateDBInstances
// 本接口用于查询支持迁移的云数据库实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIED = "UnauthorizedOperation.PermissionDenied"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeMigrateDBInstancesWithContext(ctx context.Context, request *DescribeMigrateDBInstancesRequest) (response *DescribeMigrateDBInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeMigrateDBInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrateDBInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigrateDBInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrationCheckJobRequest() (request *DescribeMigrationCheckJobRequest) {
    request = &DescribeMigrationCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrationCheckJob")
    
    
    return
}

func NewDescribeMigrationCheckJobResponse() (response *DescribeMigrationCheckJobResponse) {
    response = &DescribeMigrationCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMigrationCheckJob
// 本接口用于创建校验后,获取校验的结果. 能查询到当前校验的状态和进度. 
//
// 若通过校验, 则可调用'StartMigrateJob' 开始迁移.
//
// 若未通过校验, 则能查询到校验失败的原因. 请按照报错, 通过'ModifyMigrationJob'修改迁移配置或是调整源/目标实例的相关参数.
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeMigrationCheckJob(request *DescribeMigrationCheckJobRequest) (response *DescribeMigrationCheckJobResponse, err error) {
    return c.DescribeMigrationCheckJobWithContext(context.Background(), request)
}

// DescribeMigrationCheckJob
// 本接口用于创建校验后,获取校验的结果. 能查询到当前校验的状态和进度. 
//
// 若通过校验, 则可调用'StartMigrateJob' 开始迁移.
//
// 若未通过校验, 则能查询到校验失败的原因. 请按照报错, 通过'ModifyMigrationJob'修改迁移配置或是调整源/目标实例的相关参数.
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeMigrationCheckJobWithContext(ctx context.Context, request *DescribeMigrationCheckJobRequest) (response *DescribeMigrationCheckJobResponse, err error) {
    if request == nil {
        request = NewDescribeMigrationCheckJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrationCheckJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigrationCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrationDetailRequest() (request *DescribeMigrationDetailRequest) {
    request = &DescribeMigrationDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrationDetail")
    
    
    return
}

func NewDescribeMigrationDetailResponse() (response *DescribeMigrationDetailResponse) {
    response = &DescribeMigrationDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMigrationDetail
// 查询某个迁移任务详情
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeMigrationDetail(request *DescribeMigrationDetailRequest) (response *DescribeMigrationDetailResponse, err error) {
    return c.DescribeMigrationDetailWithContext(context.Background(), request)
}

// DescribeMigrationDetail
// 查询某个迁移任务详情
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeMigrationDetailWithContext(ctx context.Context, request *DescribeMigrationDetailRequest) (response *DescribeMigrationDetailResponse, err error) {
    if request == nil {
        request = NewDescribeMigrationDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrationDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigrationDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMigrationJobsRequest() (request *DescribeMigrationJobsRequest) {
    request = &DescribeMigrationJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeMigrationJobs")
    
    
    return
}

func NewDescribeMigrationJobsResponse() (response *DescribeMigrationJobsResponse) {
    response = &DescribeMigrationJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMigrationJobs
// 查询数据迁移任务列表
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeMigrationJobs(request *DescribeMigrationJobsRequest) (response *DescribeMigrationJobsResponse, err error) {
    return c.DescribeMigrationJobsWithContext(context.Background(), request)
}

// DescribeMigrationJobs
// 查询数据迁移任务列表
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_CONTROLLERNOTFOUNDERROR = "InvalidParameter.ControllerNotFoundError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeMigrationJobsWithContext(ctx context.Context, request *DescribeMigrationJobsRequest) (response *DescribeMigrationJobsResponse, err error) {
    if request == nil {
        request = NewDescribeMigrationJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMigrationJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMigrationJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModifyCheckSyncJobResultRequest() (request *DescribeModifyCheckSyncJobResultRequest) {
    request = &DescribeModifyCheckSyncJobResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeModifyCheckSyncJobResult")
    
    
    return
}

func NewDescribeModifyCheckSyncJobResultResponse() (response *DescribeModifyCheckSyncJobResultResponse) {
    response = &DescribeModifyCheckSyncJobResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModifyCheckSyncJobResult
// 在创建修改对象的校验任务后、通过该接口查看校验任务的结果
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) DescribeModifyCheckSyncJobResult(request *DescribeModifyCheckSyncJobResultRequest) (response *DescribeModifyCheckSyncJobResultResponse, err error) {
    return c.DescribeModifyCheckSyncJobResultWithContext(context.Background(), request)
}

// DescribeModifyCheckSyncJobResult
// 在创建修改对象的校验任务后、通过该接口查看校验任务的结果
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
func (c *Client) DescribeModifyCheckSyncJobResultWithContext(ctx context.Context, request *DescribeModifyCheckSyncJobResultRequest) (response *DescribeModifyCheckSyncJobResultResponse, err error) {
    if request == nil {
        request = NewDescribeModifyCheckSyncJobResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModifyCheckSyncJobResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModifyCheckSyncJobResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOffsetByTimeRequest() (request *DescribeOffsetByTimeRequest) {
    request = &DescribeOffsetByTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeOffsetByTime")
    
    
    return
}

func NewDescribeOffsetByTimeResponse() (response *DescribeOffsetByTimeResponse) {
    response = &DescribeOffsetByTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOffsetByTime
// 本接口(DescribeOffsetByTime)查询KafkaTopic中指定时间前最近的offset。
//
// 接口输出的offset是离这个时间最近的offset。
//
// 如果输入时间比当前时间晚的多，相当于输出的就是最新的offset；
//
// 如果输入时间比当前时间早的多，相当于输出的就是最老的offset；
//
// 如果输入空，默认0时间，也就是查询最老的offset。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeOffsetByTime(request *DescribeOffsetByTimeRequest) (response *DescribeOffsetByTimeResponse, err error) {
    return c.DescribeOffsetByTimeWithContext(context.Background(), request)
}

// DescribeOffsetByTime
// 本接口(DescribeOffsetByTime)查询KafkaTopic中指定时间前最近的offset。
//
// 接口输出的offset是离这个时间最近的offset。
//
// 如果输入时间比当前时间晚的多，相当于输出的就是最新的offset；
//
// 如果输入时间比当前时间早的多，相当于输出的就是最老的offset；
//
// 如果输入空，默认0时间，也就是查询最老的offset。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeOffsetByTimeWithContext(ctx context.Context, request *DescribeOffsetByTimeRequest) (response *DescribeOffsetByTimeResponse, err error) {
    if request == nil {
        request = NewDescribeOffsetByTimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOffsetByTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOffsetByTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscribeCheckJobRequest() (request *DescribeSubscribeCheckJobRequest) {
    request = &DescribeSubscribeCheckJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribeCheckJob")
    
    
    return
}

func NewDescribeSubscribeCheckJobResponse() (response *DescribeSubscribeCheckJobResponse) {
    response = &DescribeSubscribeCheckJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSubscribeCheckJob
// 本接口(DescribeSubscribeCheckJob)用于查询订阅校验任务结果。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeSubscribeCheckJob(request *DescribeSubscribeCheckJobRequest) (response *DescribeSubscribeCheckJobResponse, err error) {
    return c.DescribeSubscribeCheckJobWithContext(context.Background(), request)
}

// DescribeSubscribeCheckJob
// 本接口(DescribeSubscribeCheckJob)用于查询订阅校验任务结果。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeSubscribeCheckJobWithContext(ctx context.Context, request *DescribeSubscribeCheckJobRequest) (response *DescribeSubscribeCheckJobResponse, err error) {
    if request == nil {
        request = NewDescribeSubscribeCheckJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubscribeCheckJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubscribeCheckJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscribeDetailRequest() (request *DescribeSubscribeDetailRequest) {
    request = &DescribeSubscribeDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribeDetail")
    
    
    return
}

func NewDescribeSubscribeDetailResponse() (response *DescribeSubscribeDetailResponse) {
    response = &DescribeSubscribeDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSubscribeDetail
// 本接口(DescribeSubscribeDetail)获取数据订阅实例的配置信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeSubscribeDetail(request *DescribeSubscribeDetailRequest) (response *DescribeSubscribeDetailResponse, err error) {
    return c.DescribeSubscribeDetailWithContext(context.Background(), request)
}

// DescribeSubscribeDetail
// 本接口(DescribeSubscribeDetail)获取数据订阅实例的配置信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeSubscribeDetailWithContext(ctx context.Context, request *DescribeSubscribeDetailRequest) (response *DescribeSubscribeDetailResponse, err error) {
    if request == nil {
        request = NewDescribeSubscribeDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubscribeDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubscribeDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscribeJobsRequest() (request *DescribeSubscribeJobsRequest) {
    request = &DescribeSubscribeJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribeJobs")
    
    
    return
}

func NewDescribeSubscribeJobsResponse() (response *DescribeSubscribeJobsResponse) {
    response = &DescribeSubscribeJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSubscribeJobs
// 本接口(DescribeSubscribes)获取数据订阅实例信息列表，默认分页，每次返回20条
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeSubscribeJobs(request *DescribeSubscribeJobsRequest) (response *DescribeSubscribeJobsResponse, err error) {
    return c.DescribeSubscribeJobsWithContext(context.Background(), request)
}

// DescribeSubscribeJobs
// 本接口(DescribeSubscribes)获取数据订阅实例信息列表，默认分页，每次返回20条
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeSubscribeJobsWithContext(ctx context.Context, request *DescribeSubscribeJobsRequest) (response *DescribeSubscribeJobsResponse, err error) {
    if request == nil {
        request = NewDescribeSubscribeJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubscribeJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubscribeJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSubscribeReturnableRequest() (request *DescribeSubscribeReturnableRequest) {
    request = &DescribeSubscribeReturnableRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSubscribeReturnable")
    
    
    return
}

func NewDescribeSubscribeReturnableResponse() (response *DescribeSubscribeReturnableResponse) {
    response = &DescribeSubscribeReturnableResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSubscribeReturnable
// 本接口(DescribeSubscribeReturnable)用于查询订阅任务是否可以销毁和退货。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeSubscribeReturnable(request *DescribeSubscribeReturnableRequest) (response *DescribeSubscribeReturnableResponse, err error) {
    return c.DescribeSubscribeReturnableWithContext(context.Background(), request)
}

// DescribeSubscribeReturnable
// 本接口(DescribeSubscribeReturnable)用于查询订阅任务是否可以销毁和退货。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DescribeSubscribeReturnableWithContext(ctx context.Context, request *DescribeSubscribeReturnableRequest) (response *DescribeSubscribeReturnableResponse, err error) {
    if request == nil {
        request = NewDescribeSubscribeReturnableRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSubscribeReturnable require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSubscribeReturnableResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSyncJobsRequest() (request *DescribeSyncJobsRequest) {
    request = &DescribeSyncJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DescribeSyncJobs")
    
    
    return
}

func NewDescribeSyncJobsResponse() (response *DescribeSyncJobsResponse) {
    response = &DescribeSyncJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSyncJobs
// 查询同步任务信息
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeSyncJobs(request *DescribeSyncJobsRequest) (response *DescribeSyncJobsResponse, err error) {
    return c.DescribeSyncJobsWithContext(context.Background(), request)
}

// DescribeSyncJobs
// 查询同步任务信息
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
func (c *Client) DescribeSyncJobsWithContext(ctx context.Context, request *DescribeSyncJobsRequest) (response *DescribeSyncJobsResponse, err error) {
    if request == nil {
        request = NewDescribeSyncJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSyncJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSyncJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyIsolatedSubscribeRequest() (request *DestroyIsolatedSubscribeRequest) {
    request = &DestroyIsolatedSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DestroyIsolatedSubscribe")
    
    
    return
}

func NewDestroyIsolatedSubscribeResponse() (response *DestroyIsolatedSubscribeResponse) {
    response = &DestroyIsolatedSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyIsolatedSubscribe
// 本接口（DestroyIsolatedSubscribe）用于下线已隔离的数据订阅实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DestroyIsolatedSubscribe(request *DestroyIsolatedSubscribeRequest) (response *DestroyIsolatedSubscribeResponse, err error) {
    return c.DestroyIsolatedSubscribeWithContext(context.Background(), request)
}

// DestroyIsolatedSubscribe
// 本接口（DestroyIsolatedSubscribe）用于下线已隔离的数据订阅实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DestroyIsolatedSubscribeWithContext(ctx context.Context, request *DestroyIsolatedSubscribeRequest) (response *DestroyIsolatedSubscribeResponse, err error) {
    if request == nil {
        request = NewDestroyIsolatedSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyIsolatedSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyIsolatedSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyMigrateJobRequest() (request *DestroyMigrateJobRequest) {
    request = &DestroyMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DestroyMigrateJob")
    
    
    return
}

func NewDestroyMigrateJobResponse() (response *DestroyMigrateJobResponse) {
    response = &DestroyMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroyMigrateJob
// 下线数据迁移任务。计费任务必须先调用隔离(IsolateMigrateJob)接口，且只有是**已隔离**状态下，才能调用此接口销毁任务。对于不计费任务，调用隔离(IsolateMigrateJob)接口删除任务操作。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DestroyMigrateJob(request *DestroyMigrateJobRequest) (response *DestroyMigrateJobResponse, err error) {
    return c.DestroyMigrateJobWithContext(context.Background(), request)
}

// DestroyMigrateJob
// 下线数据迁移任务。计费任务必须先调用隔离(IsolateMigrateJob)接口，且只有是**已隔离**状态下，才能调用此接口销毁任务。对于不计费任务，调用隔离(IsolateMigrateJob)接口删除任务操作。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) DestroyMigrateJobWithContext(ctx context.Context, request *DestroyMigrateJobRequest) (response *DestroyMigrateJobResponse, err error) {
    if request == nil {
        request = NewDestroyMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewDestroySyncJobRequest() (request *DestroySyncJobRequest) {
    request = &DestroySyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "DestroySyncJob")
    
    
    return
}

func NewDestroySyncJobResponse() (response *DestroySyncJobResponse) {
    response = &DestroySyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DestroySyncJob
// 下线同步任务，任务在已隔离状态下可以通过此操作进行任务下线，即彻底删除任务。下线操作后可通过查询同步任务信息接口DescribeSyncJobs获取任务列表查看状态，此操作成功后无法看到此任务表示下线成功。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALTRADEERROR = "InternalError.InternalTradeError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DestroySyncJob(request *DestroySyncJobRequest) (response *DestroySyncJobResponse, err error) {
    return c.DestroySyncJobWithContext(context.Background(), request)
}

// DestroySyncJob
// 下线同步任务，任务在已隔离状态下可以通过此操作进行任务下线，即彻底删除任务。下线操作后可通过查询同步任务信息接口DescribeSyncJobs获取任务列表查看状态，此操作成功后无法看到此任务表示下线成功。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALTRADEERROR = "InternalError.InternalTradeError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) DestroySyncJobWithContext(ctx context.Context, request *DestroySyncJobRequest) (response *DestroySyncJobResponse, err error) {
    if request == nil {
        request = NewDestroySyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroySyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroySyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateMigrateJobRequest() (request *IsolateMigrateJobRequest) {
    request = &IsolateMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "IsolateMigrateJob")
    
    
    return
}

func NewIsolateMigrateJobResponse() (response *IsolateMigrateJobResponse) {
    response = &IsolateMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IsolateMigrateJob
//  隔离退还数据迁移服务。调用此接口后可通过查询迁移服务列表接口`DescribeMigrationJobs`来查询当前任务状态。对于计费任务，在任务隔离后可进行解除隔离(RecoverMigrationJob)操作或直接进行下线销毁(DestroyMigrateJob)操作。对于不计费任务，调用此接口会直接销毁任务，无法进行恢复操作。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) IsolateMigrateJob(request *IsolateMigrateJobRequest) (response *IsolateMigrateJobResponse, err error) {
    return c.IsolateMigrateJobWithContext(context.Background(), request)
}

// IsolateMigrateJob
//  隔离退还数据迁移服务。调用此接口后可通过查询迁移服务列表接口`DescribeMigrationJobs`来查询当前任务状态。对于计费任务，在任务隔离后可进行解除隔离(RecoverMigrationJob)操作或直接进行下线销毁(DestroyMigrateJob)操作。对于不计费任务，调用此接口会直接销毁任务，无法进行恢复操作。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) IsolateMigrateJobWithContext(ctx context.Context, request *IsolateMigrateJobRequest) (response *IsolateMigrateJobResponse, err error) {
    if request == nil {
        request = NewIsolateMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateMigrateJobResponse()
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
// 本接口（IsolateSubscribe）用于隔离订阅任务。调用后，订阅任务将不能使用。按量计费的任务会停止计费，包年包月的任务会自动退费
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) IsolateSubscribe(request *IsolateSubscribeRequest) (response *IsolateSubscribeResponse, err error) {
    return c.IsolateSubscribeWithContext(context.Background(), request)
}

// IsolateSubscribe
// 本接口（IsolateSubscribe）用于隔离订阅任务。调用后，订阅任务将不能使用。按量计费的任务会停止计费，包年包月的任务会自动退费
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
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

func NewIsolateSyncJobRequest() (request *IsolateSyncJobRequest) {
    request = &IsolateSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "IsolateSyncJob")
    
    
    return
}

func NewIsolateSyncJobResponse() (response *IsolateSyncJobResponse) {
    response = &IsolateSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// IsolateSyncJob
// 隔离同步任务，隔离后可通过查询同步任务信息接口DescribeSyncJobs获取隔离后状态。在任务隔离后可进行解除隔离(RecoverSyncJob)操作或直接进行下线操作。对于不计费任务，调用此接口后会直接删除任务，无法进行恢复操作。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALTRADEERROR = "InternalError.InternalTradeError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) IsolateSyncJob(request *IsolateSyncJobRequest) (response *IsolateSyncJobResponse, err error) {
    return c.IsolateSyncJobWithContext(context.Background(), request)
}

// IsolateSyncJob
// 隔离同步任务，隔离后可通过查询同步任务信息接口DescribeSyncJobs获取隔离后状态。在任务隔离后可进行解除隔离(RecoverSyncJob)操作或直接进行下线操作。对于不计费任务，调用此接口后会直接删除任务，无法进行恢复操作。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALTRADEERROR = "InternalError.InternalTradeError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) IsolateSyncJobWithContext(ctx context.Context, request *IsolateSyncJobRequest) (response *IsolateSyncJobResponse, err error) {
    if request == nil {
        request = NewIsolateSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("IsolateSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewIsolateSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCompareTaskRequest() (request *ModifyCompareTaskRequest) {
    request = &ModifyCompareTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifyCompareTask")
    
    
    return
}

func NewModifyCompareTaskResponse() (response *ModifyCompareTaskResponse) {
    response = &ModifyCompareTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCompareTask
// 修改一致性校验任务，在任务创建后启动之前，可修改一致性校验参数
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyCompareTask(request *ModifyCompareTaskRequest) (response *ModifyCompareTaskResponse, err error) {
    return c.ModifyCompareTaskWithContext(context.Background(), request)
}

// ModifyCompareTask
// 修改一致性校验任务，在任务创建后启动之前，可修改一致性校验参数
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyCompareTaskWithContext(ctx context.Context, request *ModifyCompareTaskRequest) (response *ModifyCompareTaskResponse, err error) {
    if request == nil {
        request = NewModifyCompareTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCompareTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCompareTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCompareTaskNameRequest() (request *ModifyCompareTaskNameRequest) {
    request = &ModifyCompareTaskNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifyCompareTaskName")
    
    
    return
}

func NewModifyCompareTaskNameResponse() (response *ModifyCompareTaskNameResponse) {
    response = &ModifyCompareTaskNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCompareTaskName
// 修改一致性校验任务名称
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyCompareTaskName(request *ModifyCompareTaskNameRequest) (response *ModifyCompareTaskNameResponse, err error) {
    return c.ModifyCompareTaskNameWithContext(context.Background(), request)
}

// ModifyCompareTaskName
// 修改一致性校验任务名称
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyCompareTaskNameWithContext(ctx context.Context, request *ModifyCompareTaskNameRequest) (response *ModifyCompareTaskNameResponse, err error) {
    if request == nil {
        request = NewModifyCompareTaskNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCompareTaskName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCompareTaskNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConsumerGroupDescriptionRequest() (request *ModifyConsumerGroupDescriptionRequest) {
    request = &ModifyConsumerGroupDescriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifyConsumerGroupDescription")
    
    
    return
}

func NewModifyConsumerGroupDescriptionResponse() (response *ModifyConsumerGroupDescriptionResponse) {
    response = &ModifyConsumerGroupDescriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConsumerGroupDescription
// 本接口(ModifyConsumerGroupDescription)用于修改指定订阅消费组备注。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyConsumerGroupDescription(request *ModifyConsumerGroupDescriptionRequest) (response *ModifyConsumerGroupDescriptionResponse, err error) {
    return c.ModifyConsumerGroupDescriptionWithContext(context.Background(), request)
}

// ModifyConsumerGroupDescription
// 本接口(ModifyConsumerGroupDescription)用于修改指定订阅消费组备注。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyConsumerGroupDescriptionWithContext(ctx context.Context, request *ModifyConsumerGroupDescriptionRequest) (response *ModifyConsumerGroupDescriptionResponse, err error) {
    if request == nil {
        request = NewModifyConsumerGroupDescriptionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConsumerGroupDescription require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConsumerGroupDescriptionResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConsumerGroupPasswordRequest() (request *ModifyConsumerGroupPasswordRequest) {
    request = &ModifyConsumerGroupPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifyConsumerGroupPassword")
    
    
    return
}

func NewModifyConsumerGroupPasswordResponse() (response *ModifyConsumerGroupPasswordResponse) {
    response = &ModifyConsumerGroupPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyConsumerGroupPassword
// 本接口(ModifyConsumerGroupPassword)用于修改指定订阅消费组密码。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyConsumerGroupPassword(request *ModifyConsumerGroupPasswordRequest) (response *ModifyConsumerGroupPasswordResponse, err error) {
    return c.ModifyConsumerGroupPasswordWithContext(context.Background(), request)
}

// ModifyConsumerGroupPassword
// 本接口(ModifyConsumerGroupPassword)用于修改指定订阅消费组密码。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyConsumerGroupPasswordWithContext(ctx context.Context, request *ModifyConsumerGroupPasswordRequest) (response *ModifyConsumerGroupPasswordResponse, err error) {
    if request == nil {
        request = NewModifyConsumerGroupPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConsumerGroupPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConsumerGroupPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMigrateJobSpecRequest() (request *ModifyMigrateJobSpecRequest) {
    request = &ModifyMigrateJobSpecRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifyMigrateJobSpec")
    
    
    return
}

func NewModifyMigrateJobSpecResponse() (response *ModifyMigrateJobSpecResponse) {
    response = &ModifyMigrateJobSpecResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMigrateJobSpec
// 调整实例规格，此接口只支持按量计费任务的调整。调用此接口后可通过查询迁移服务列表接口`DescribeMigrationJobs`来查询当前任务状态。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyMigrateJobSpec(request *ModifyMigrateJobSpecRequest) (response *ModifyMigrateJobSpecResponse, err error) {
    return c.ModifyMigrateJobSpecWithContext(context.Background(), request)
}

// ModifyMigrateJobSpec
// 调整实例规格，此接口只支持按量计费任务的调整。调用此接口后可通过查询迁移服务列表接口`DescribeMigrationJobs`来查询当前任务状态。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyMigrateJobSpecWithContext(ctx context.Context, request *ModifyMigrateJobSpecRequest) (response *ModifyMigrateJobSpecResponse, err error) {
    if request == nil {
        request = NewModifyMigrateJobSpecRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMigrateJobSpec require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMigrateJobSpecResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMigrateNameRequest() (request *ModifyMigrateNameRequest) {
    request = &ModifyMigrateNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifyMigrateName")
    
    
    return
}

func NewModifyMigrateNameResponse() (response *ModifyMigrateNameResponse) {
    response = &ModifyMigrateNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMigrateName
// 修改迁移任务名
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyMigrateName(request *ModifyMigrateNameRequest) (response *ModifyMigrateNameResponse, err error) {
    return c.ModifyMigrateNameWithContext(context.Background(), request)
}

// ModifyMigrateName
// 修改迁移任务名
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyMigrateNameWithContext(ctx context.Context, request *ModifyMigrateNameRequest) (response *ModifyMigrateNameResponse, err error) {
    if request == nil {
        request = NewModifyMigrateNameRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMigrateName require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMigrateNameResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMigrateRateLimitRequest() (request *ModifyMigrateRateLimitRequest) {
    request = &ModifyMigrateRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifyMigrateRateLimit")
    
    
    return
}

func NewModifyMigrateRateLimitResponse() (response *ModifyMigrateRateLimitResponse) {
    response = &ModifyMigrateRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMigrateRateLimit
// 用户在发现迁移任务对用户的数据库的负载影响较大时、可通过该接口限制任务的传输速率
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyMigrateRateLimit(request *ModifyMigrateRateLimitRequest) (response *ModifyMigrateRateLimitResponse, err error) {
    return c.ModifyMigrateRateLimitWithContext(context.Background(), request)
}

// ModifyMigrateRateLimit
// 用户在发现迁移任务对用户的数据库的负载影响较大时、可通过该接口限制任务的传输速率
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyMigrateRateLimitWithContext(ctx context.Context, request *ModifyMigrateRateLimitRequest) (response *ModifyMigrateRateLimitResponse, err error) {
    if request == nil {
        request = NewModifyMigrateRateLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMigrateRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMigrateRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMigrateRuntimeAttributeRequest() (request *ModifyMigrateRuntimeAttributeRequest) {
    request = &ModifyMigrateRuntimeAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifyMigrateRuntimeAttribute")
    
    
    return
}

func NewModifyMigrateRuntimeAttributeResponse() (response *ModifyMigrateRuntimeAttributeResponse) {
    response = &ModifyMigrateRuntimeAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMigrateRuntimeAttribute
// 修改任务运行时属性，此接口不同于配置类接口，不会进行状态机判断。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
func (c *Client) ModifyMigrateRuntimeAttribute(request *ModifyMigrateRuntimeAttributeRequest) (response *ModifyMigrateRuntimeAttributeResponse, err error) {
    return c.ModifyMigrateRuntimeAttributeWithContext(context.Background(), request)
}

// ModifyMigrateRuntimeAttribute
// 修改任务运行时属性，此接口不同于配置类接口，不会进行状态机判断。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
func (c *Client) ModifyMigrateRuntimeAttributeWithContext(ctx context.Context, request *ModifyMigrateRuntimeAttributeRequest) (response *ModifyMigrateRuntimeAttributeResponse, err error) {
    if request == nil {
        request = NewModifyMigrateRuntimeAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMigrateRuntimeAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMigrateRuntimeAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMigrationJobRequest() (request *ModifyMigrationJobRequest) {
    request = &ModifyMigrationJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifyMigrationJob")
    
    
    return
}

func NewModifyMigrationJobResponse() (response *ModifyMigrationJobResponse) {
    response = &ModifyMigrationJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyMigrationJob
// 配置迁移服务，配置成功后可通过`CreateMigrationCheckJob` 创建迁移校验任务接口发起校验任务，只有校验通过才能启动迁移任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_AUTHORIZEDOPERATIONDENYERROR = "AuthFailure.AuthorizedOperationDenyError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyMigrationJob(request *ModifyMigrationJobRequest) (response *ModifyMigrationJobResponse, err error) {
    return c.ModifyMigrationJobWithContext(context.Background(), request)
}

// ModifyMigrationJob
// 配置迁移服务，配置成功后可通过`CreateMigrationCheckJob` 创建迁移校验任务接口发起校验任务，只有校验通过才能启动迁移任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  AUTHFAILURE_AUTHORIZEDOPERATIONDENYERROR = "AuthFailure.AuthorizedOperationDenyError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATIONERROR = "AuthFailure.UnauthorizedOperationError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_DATACONVERTERROR = "InvalidParameterValue.DataConvertError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifyMigrationJobWithContext(ctx context.Context, request *ModifyMigrationJobRequest) (response *ModifyMigrationJobResponse, err error) {
    if request == nil {
        request = NewModifyMigrationJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMigrationJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMigrationJobResponse()
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
// 修改订阅实例自动续费标识。只有包年包月的任务修改才有意义，按量计费任务修改后无影响。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifySubscribeAutoRenewFlag(request *ModifySubscribeAutoRenewFlagRequest) (response *ModifySubscribeAutoRenewFlagResponse, err error) {
    return c.ModifySubscribeAutoRenewFlagWithContext(context.Background(), request)
}

// ModifySubscribeAutoRenewFlag
// 修改订阅实例自动续费标识。只有包年包月的任务修改才有意义，按量计费任务修改后无影响。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
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
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifySubscribeName(request *ModifySubscribeNameRequest) (response *ModifySubscribeNameResponse, err error) {
    return c.ModifySubscribeNameWithContext(context.Background(), request)
}

// ModifySubscribeName
// 本接口(ModifySubscribeName)用于修改数据订阅实例的名称
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
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
// 本接口(ModifySubscribeObjects)用于修改数据订阅对象和kafka分区规则，如果是mongo订阅，还可以修改输出聚合规则。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ModifySubscribeObjects(request *ModifySubscribeObjectsRequest) (response *ModifySubscribeObjectsResponse, err error) {
    return c.ModifySubscribeObjectsWithContext(context.Background(), request)
}

// ModifySubscribeObjects
// 本接口(ModifySubscribeObjects)用于修改数据订阅对象和kafka分区规则，如果是mongo订阅，还可以修改输出聚合规则。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
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

func NewModifySyncJobConfigRequest() (request *ModifySyncJobConfigRequest) {
    request = &ModifySyncJobConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifySyncJobConfig")
    
    
    return
}

func NewModifySyncJobConfigResponse() (response *ModifySyncJobConfigResponse) {
    response = &ModifySyncJobConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySyncJobConfig
// 该接口支持在同步任务启动后修改任务的配置
//
// 修改同步配置的完整流程：修改同步任务配置->创建修改同步任务配置的校验任务->查询修改配置的校验任务的结果->启动修改配置任务
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
func (c *Client) ModifySyncJobConfig(request *ModifySyncJobConfigRequest) (response *ModifySyncJobConfigResponse, err error) {
    return c.ModifySyncJobConfigWithContext(context.Background(), request)
}

// ModifySyncJobConfig
// 该接口支持在同步任务启动后修改任务的配置
//
// 修改同步配置的完整流程：修改同步任务配置->创建修改同步任务配置的校验任务->查询修改配置的校验任务的结果->启动修改配置任务
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
func (c *Client) ModifySyncJobConfigWithContext(ctx context.Context, request *ModifySyncJobConfigRequest) (response *ModifySyncJobConfigResponse, err error) {
    if request == nil {
        request = NewModifySyncJobConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySyncJobConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySyncJobConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifySyncRateLimitRequest() (request *ModifySyncRateLimitRequest) {
    request = &ModifySyncRateLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ModifySyncRateLimit")
    
    
    return
}

func NewModifySyncRateLimitResponse() (response *ModifySyncRateLimitResponse) {
    response = &ModifySyncRateLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifySyncRateLimit
// 用户在发现同步任务对用户的数据库的负载影响较大时、可通过该接口限制任务的传输速率
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) ModifySyncRateLimit(request *ModifySyncRateLimitRequest) (response *ModifySyncRateLimitResponse, err error) {
    return c.ModifySyncRateLimitWithContext(context.Background(), request)
}

// ModifySyncRateLimit
// 用户在发现同步任务对用户的数据库的负载影响较大时、可通过该接口限制任务的传输速率
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) ModifySyncRateLimitWithContext(ctx context.Context, request *ModifySyncRateLimitRequest) (response *ModifySyncRateLimitResponse, err error) {
    if request == nil {
        request = NewModifySyncRateLimitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySyncRateLimit require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifySyncRateLimitResponse()
    err = c.Send(request, response)
    return
}

func NewPauseMigrateJobRequest() (request *PauseMigrateJobRequest) {
    request = &PauseMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "PauseMigrateJob")
    
    
    return
}

func NewPauseMigrateJobResponse() (response *PauseMigrateJobResponse) {
    response = &PauseMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PauseMigrateJob
// 暂停一个迁移任务。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) PauseMigrateJob(request *PauseMigrateJobRequest) (response *PauseMigrateJobResponse, err error) {
    return c.PauseMigrateJobWithContext(context.Background(), request)
}

// PauseMigrateJob
// 暂停一个迁移任务。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) PauseMigrateJobWithContext(ctx context.Context, request *PauseMigrateJobRequest) (response *PauseMigrateJobResponse, err error) {
    if request == nil {
        request = NewPauseMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PauseMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewPauseMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewPauseSyncJobRequest() (request *PauseSyncJobRequest) {
    request = &PauseSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "PauseSyncJob")
    
    
    return
}

func NewPauseSyncJobResponse() (response *PauseSyncJobResponse) {
    response = &PauseSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// PauseSyncJob
// 暂停处于同步中的数据同步任务。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) PauseSyncJob(request *PauseSyncJobRequest) (response *PauseSyncJobResponse, err error) {
    return c.PauseSyncJobWithContext(context.Background(), request)
}

// PauseSyncJob
// 暂停处于同步中的数据同步任务。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZRESOURCENOTFOUNDERROR = "ResourceNotFound.BizResourceNotFoundError"
//  RESOURCENOTFOUND_RESOURCENOTFOUND = "ResourceNotFound.ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) PauseSyncJobWithContext(ctx context.Context, request *PauseSyncJobRequest) (response *PauseSyncJobResponse, err error) {
    if request == nil {
        request = NewPauseSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PauseSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewPauseSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewRecoverMigrateJobRequest() (request *RecoverMigrateJobRequest) {
    request = &RecoverMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "RecoverMigrateJob")
    
    
    return
}

func NewRecoverMigrateJobResponse() (response *RecoverMigrateJobResponse) {
    response = &RecoverMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecoverMigrateJob
// 解除隔离数据迁移任务，用户手动发起隔离后的手动解隔离，只有任务状态为已隔离(手动操作)状态下才能触发此操作。调用此接口后可通过查询迁移服务列表接口`DescribeMigrationJobs`来查询当前任务状态。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) RecoverMigrateJob(request *RecoverMigrateJobRequest) (response *RecoverMigrateJobResponse, err error) {
    return c.RecoverMigrateJobWithContext(context.Background(), request)
}

// RecoverMigrateJob
// 解除隔离数据迁移任务，用户手动发起隔离后的手动解隔离，只有任务状态为已隔离(手动操作)状态下才能触发此操作。调用此接口后可通过查询迁移服务列表接口`DescribeMigrationJobs`来查询当前任务状态。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALHTTPSERVERERROR = "InternalError.InternalHttpServerError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) RecoverMigrateJobWithContext(ctx context.Context, request *RecoverMigrateJobRequest) (response *RecoverMigrateJobResponse, err error) {
    if request == nil {
        request = NewRecoverMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecoverMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecoverMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewRecoverSyncJobRequest() (request *RecoverSyncJobRequest) {
    request = &RecoverSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "RecoverSyncJob")
    
    
    return
}

func NewRecoverSyncJobResponse() (response *RecoverSyncJobResponse) {
    response = &RecoverSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RecoverSyncJob
// 解除隔离同步任务，任务在已隔离状态下可调用该接口解除隔离状态任务，同时可通过查询同步任务信息接口DescribeSyncJobs，获取操作后状态。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) RecoverSyncJob(request *RecoverSyncJobRequest) (response *RecoverSyncJobResponse, err error) {
    return c.RecoverSyncJobWithContext(context.Background(), request)
}

// RecoverSyncJob
// 解除隔离同步任务，任务在已隔离状态下可调用该接口解除隔离状态任务，同时可通过查询同步任务信息接口DescribeSyncJobs，获取操作后状态。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) RecoverSyncJobWithContext(ctx context.Context, request *RecoverSyncJobRequest) (response *RecoverSyncJobResponse, err error) {
    if request == nil {
        request = NewRecoverSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecoverSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecoverSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewResetConsumerGroupOffsetRequest() (request *ResetConsumerGroupOffsetRequest) {
    request = &ResetConsumerGroupOffsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ResetConsumerGroupOffset")
    
    
    return
}

func NewResetConsumerGroupOffsetResponse() (response *ResetConsumerGroupOffsetResponse) {
    response = &ResetConsumerGroupOffsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResetConsumerGroupOffset
// 本接口(ResetConsumerGroupOffset)用于重置订阅消费组的offset。调用DescribeConsumerGroups接口查询消费组状态，只有消费组状态为 Dead 或 Empty 才可以执行重置该操作。否则重置不会生效，接口也不会报错。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ResetConsumerGroupOffset(request *ResetConsumerGroupOffsetRequest) (response *ResetConsumerGroupOffsetResponse, err error) {
    return c.ResetConsumerGroupOffsetWithContext(context.Background(), request)
}

// ResetConsumerGroupOffset
// 本接口(ResetConsumerGroupOffset)用于重置订阅消费组的offset。调用DescribeConsumerGroups接口查询消费组状态，只有消费组状态为 Dead 或 Empty 才可以执行重置该操作。否则重置不会生效，接口也不会报错。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ResetConsumerGroupOffsetWithContext(ctx context.Context, request *ResetConsumerGroupOffsetRequest) (response *ResetConsumerGroupOffsetResponse, err error) {
    if request == nil {
        request = NewResetConsumerGroupOffsetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResetConsumerGroupOffset require credential")
    }

    request.SetContext(ctx)
    
    response = NewResetConsumerGroupOffsetResponse()
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
// 本接口(ResetSubscribe)用于重置订阅实例，重置后，可以重新配置订阅任务。
//
// 可以调用 DescribeSubscribeDetail 查询订阅信息判断是否置成功。当SubsStatus变为notStarted时，表示重置成功。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ResetSubscribe(request *ResetSubscribeRequest) (response *ResetSubscribeResponse, err error) {
    return c.ResetSubscribeWithContext(context.Background(), request)
}

// ResetSubscribe
// 本接口(ResetSubscribe)用于重置订阅实例，重置后，可以重新配置订阅任务。
//
// 可以调用 DescribeSubscribeDetail 查询订阅信息判断是否置成功。当SubsStatus变为notStarted时，表示重置成功。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
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

func NewResizeSyncJobRequest() (request *ResizeSyncJobRequest) {
    request = &ResizeSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ResizeSyncJob")
    
    
    return
}

func NewResizeSyncJobResponse() (response *ResizeSyncJobResponse) {
    response = &ResizeSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResizeSyncJob
// 调整同步任务规格，此接口只支持按量计费任务的调整，调用此接口后不会立即生效，后台调整时间大概为3~5分钟。调用此接口后可通过查询同步任务信息接口DescribeSyncJobs，获取变配后的状态。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALTRADEERROR = "InternalError.InternalTradeError"
//  INTERNALERROR_NOTENOUGHMONEYERROR = "InternalError.NotEnoughMoneyError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ResizeSyncJob(request *ResizeSyncJobRequest) (response *ResizeSyncJobResponse, err error) {
    return c.ResizeSyncJobWithContext(context.Background(), request)
}

// ResizeSyncJob
// 调整同步任务规格，此接口只支持按量计费任务的调整，调用此接口后不会立即生效，后台调整时间大概为3~5分钟。调用此接口后可通过查询同步任务信息接口DescribeSyncJobs，获取变配后的状态。
//
// 可能返回的错误码:
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_INTERNALTRADEERROR = "InternalError.InternalTradeError"
//  INTERNALERROR_NOTENOUGHMONEYERROR = "InternalError.NotEnoughMoneyError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ResizeSyncJobWithContext(ctx context.Context, request *ResizeSyncJobRequest) (response *ResizeSyncJobResponse, err error) {
    if request == nil {
        request = NewResizeSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResizeSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewResizeSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewResumeMigrateJobRequest() (request *ResumeMigrateJobRequest) {
    request = &ResumeMigrateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ResumeMigrateJob")
    
    
    return
}

func NewResumeMigrateJobResponse() (response *ResumeMigrateJobResponse) {
    response = &ResumeMigrateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumeMigrateJob
// 重试数据迁移任务，针对异常情况可进行重试，对于redis在失败时也可重试。注意：此操作跳过校验阶段，直接重新发起任务，相当于从StartMigrationJob开始执行。调用此接口后可通过查询迁移服务列表接口`DescribeMigrationJobs`来查询当前任务状态。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ResumeMigrateJob(request *ResumeMigrateJobRequest) (response *ResumeMigrateJobResponse, err error) {
    return c.ResumeMigrateJobWithContext(context.Background(), request)
}

// ResumeMigrateJob
// 重试数据迁移任务，针对异常情况可进行重试，对于redis在失败时也可重试。注意：此操作跳过校验阶段，直接重新发起任务，相当于从StartMigrationJob开始执行。调用此接口后可通过查询迁移服务列表接口`DescribeMigrationJobs`来查询当前任务状态。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUE = "InvalidParameterValue.InvalidParameterValue"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ResumeMigrateJobWithContext(ctx context.Context, request *ResumeMigrateJobRequest) (response *ResumeMigrateJobResponse, err error) {
    if request == nil {
        request = NewResumeMigrateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeMigrateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeMigrateJobResponse()
    err = c.Send(request, response)
    return
}

func NewResumeSubscribeRequest() (request *ResumeSubscribeRequest) {
    request = &ResumeSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ResumeSubscribe")
    
    
    return
}

func NewResumeSubscribeResponse() (response *ResumeSubscribeResponse) {
    response = &ResumeSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumeSubscribe
// 本接口(ResumeSubscribe) 用于恢复报错的订阅任务。当订阅任务的状态为error时，可通过本接口尝试对任务进行恢复。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ResumeSubscribe(request *ResumeSubscribeRequest) (response *ResumeSubscribeResponse, err error) {
    return c.ResumeSubscribeWithContext(context.Background(), request)
}

// ResumeSubscribe
// 本接口(ResumeSubscribe) 用于恢复报错的订阅任务。当订阅任务的状态为error时，可通过本接口尝试对任务进行恢复。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) ResumeSubscribeWithContext(ctx context.Context, request *ResumeSubscribeRequest) (response *ResumeSubscribeResponse, err error) {
    if request == nil {
        request = NewResumeSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewResumeSyncJobRequest() (request *ResumeSyncJobRequest) {
    request = &ResumeSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "ResumeSyncJob")
    
    
    return
}

func NewResumeSyncJobResponse() (response *ResumeSyncJobResponse) {
    response = &ResumeSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ResumeSyncJob
// 重试同步任务，部分可恢复报错情况下，可通过该接口重试同步任务，可通过查询同步任务信息接口DescribeSyncJobs，获取操作后状态。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ResumeSyncJob(request *ResumeSyncJobRequest) (response *ResumeSyncJobResponse, err error) {
    return c.ResumeSyncJobWithContext(context.Background(), request)
}

// ResumeSyncJob
// 重试同步任务，部分可恢复报错情况下，可通过该接口重试同步任务，可通过查询同步任务信息接口DescribeSyncJobs，获取操作后状态。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
func (c *Client) ResumeSyncJobWithContext(ctx context.Context, request *ResumeSyncJobRequest) (response *ResumeSyncJobResponse, err error) {
    if request == nil {
        request = NewResumeSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewSkipCheckItemRequest() (request *SkipCheckItemRequest) {
    request = &SkipCheckItemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "SkipCheckItem")
    
    
    return
}

func NewSkipCheckItemResponse() (response *SkipCheckItemResponse) {
    response = &SkipCheckItemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SkipCheckItem
// 本接口用于校验检查项不通过后，可进行跳过此校验项操作，后端将不再校验该项。任何校验步骤都是不应该跳过的，通过校验是能正确执行的前置条件。支持跳过的产品及链路的校验项可 [参考文档](https://cloud.tencent.com/document/product/571/61639)。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) SkipCheckItem(request *SkipCheckItemRequest) (response *SkipCheckItemResponse, err error) {
    return c.SkipCheckItemWithContext(context.Background(), request)
}

// SkipCheckItem
// 本接口用于校验检查项不通过后，可进行跳过此校验项操作，后端将不再校验该项。任何校验步骤都是不应该跳过的，通过校验是能正确执行的前置条件。支持跳过的产品及链路的校验项可 [参考文档](https://cloud.tencent.com/document/product/571/61639)。
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) SkipCheckItemWithContext(ctx context.Context, request *SkipCheckItemRequest) (response *SkipCheckItemResponse, err error) {
    if request == nil {
        request = NewSkipCheckItemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SkipCheckItem require credential")
    }

    request.SetContext(ctx)
    
    response = NewSkipCheckItemResponse()
    err = c.Send(request, response)
    return
}

func NewSkipSyncCheckItemRequest() (request *SkipSyncCheckItemRequest) {
    request = &SkipSyncCheckItemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "SkipSyncCheckItem")
    
    
    return
}

func NewSkipSyncCheckItemResponse() (response *SkipSyncCheckItemResponse) {
    response = &SkipSyncCheckItemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SkipSyncCheckItem
// 本接口用于校验检查项不通过后，可进行跳过此校验项操作，后端将不再校验该项。任何校验步骤都是不应该跳过的，通过校验是能正确执行的前置条件。支持跳过的产品及链路的校验项可 [参考文档](https://cloud.tencent.com/document/product/571/61639)。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) SkipSyncCheckItem(request *SkipSyncCheckItemRequest) (response *SkipSyncCheckItemResponse, err error) {
    return c.SkipSyncCheckItemWithContext(context.Background(), request)
}

// SkipSyncCheckItem
// 本接口用于校验检查项不通过后，可进行跳过此校验项操作，后端将不再校验该项。任何校验步骤都是不应该跳过的，通过校验是能正确执行的前置条件。支持跳过的产品及链路的校验项可 [参考文档](https://cloud.tencent.com/document/product/571/61639)。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) SkipSyncCheckItemWithContext(ctx context.Context, request *SkipSyncCheckItemRequest) (response *SkipSyncCheckItemResponse, err error) {
    if request == nil {
        request = NewSkipSyncCheckItemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SkipSyncCheckItem require credential")
    }

    request.SetContext(ctx)
    
    response = NewSkipSyncCheckItemResponse()
    err = c.Send(request, response)
    return
}

func NewStartCompareRequest() (request *StartCompareRequest) {
    request = &StartCompareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "StartCompare")
    
    
    return
}

func NewStartCompareResponse() (response *StartCompareResponse) {
    response = &StartCompareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartCompare
// 启动一致性校验任务，启动之前需要先通过接口`CreateCompareTask` 创建一致性校验任务，启动后可通过接口`DescribeCompareTasks` 查询一致性校验任务列表来获得启动后的状态
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) StartCompare(request *StartCompareRequest) (response *StartCompareResponse, err error) {
    return c.StartCompareWithContext(context.Background(), request)
}

// StartCompare
// 启动一致性校验任务，启动之前需要先通过接口`CreateCompareTask` 创建一致性校验任务，启动后可通过接口`DescribeCompareTasks` 查询一致性校验任务列表来获得启动后的状态
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) StartCompareWithContext(ctx context.Context, request *StartCompareRequest) (response *StartCompareResponse, err error) {
    if request == nil {
        request = NewStartCompareRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartCompare require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartCompareResponse()
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
// 本接口（StartMigrationJob）用于启动迁移任务。调用此接口后可通过查询迁移服务列表接口`DescribeMigrationJobs`来查询当前任务状态。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
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
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) StartMigrateJob(request *StartMigrateJobRequest) (response *StartMigrateJobResponse, err error) {
    return c.StartMigrateJobWithContext(context.Background(), request)
}

// StartMigrateJob
// 本接口（StartMigrationJob）用于启动迁移任务。调用此接口后可通过查询迁移服务列表接口`DescribeMigrationJobs`来查询当前任务状态。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
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
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  LIMITEXCEEDED_MAXUNUSEDJOBS = "LimitExceeded.MaxUnusedJobs"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
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

func NewStartModifySyncJobRequest() (request *StartModifySyncJobRequest) {
    request = &StartModifySyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "StartModifySyncJob")
    
    
    return
}

func NewStartModifySyncJobResponse() (response *StartModifySyncJobResponse) {
    response = &StartModifySyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartModifySyncJob
// 在查询修改对象的校验任务的结果中的status为success后、通过该接口开始修改配置流程
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) StartModifySyncJob(request *StartModifySyncJobRequest) (response *StartModifySyncJobResponse, err error) {
    return c.StartModifySyncJobWithContext(context.Background(), request)
}

// StartModifySyncJob
// 在查询修改对象的校验任务的结果中的status为success后、通过该接口开始修改配置流程
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) StartModifySyncJobWithContext(ctx context.Context, request *StartModifySyncJobRequest) (response *StartModifySyncJobResponse, err error) {
    if request == nil {
        request = NewStartModifySyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartModifySyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartModifySyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewStartSubscribeRequest() (request *StartSubscribeRequest) {
    request = &StartSubscribeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "StartSubscribe")
    
    
    return
}

func NewStartSubscribeResponse() (response *StartSubscribeResponse) {
    response = &StartSubscribeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartSubscribe
// 本接口(StartSubscribe)用于启动一个kafka版本的数据订阅实例。只有当订阅任务的状态为checkPass时，才能调用本接口。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) StartSubscribe(request *StartSubscribeRequest) (response *StartSubscribeResponse, err error) {
    return c.StartSubscribeWithContext(context.Background(), request)
}

// StartSubscribe
// 本接口(StartSubscribe)用于启动一个kafka版本的数据订阅实例。只有当订阅任务的状态为checkPass时，才能调用本接口。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) StartSubscribeWithContext(ctx context.Context, request *StartSubscribeRequest) (response *StartSubscribeResponse, err error) {
    if request == nil {
        request = NewStartSubscribeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartSubscribe require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartSubscribeResponse()
    err = c.Send(request, response)
    return
}

func NewStartSyncJobRequest() (request *StartSyncJobRequest) {
    request = &StartSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "StartSyncJob")
    
    
    return
}

func NewStartSyncJobResponse() (response *StartSyncJobResponse) {
    response = &StartSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartSyncJob
// 启动同步任务
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) StartSyncJob(request *StartSyncJobRequest) (response *StartSyncJobResponse, err error) {
    return c.StartSyncJobWithContext(context.Background(), request)
}

// StartSyncJob
// 启动同步任务
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  OPERATIONDENIED_BIZOPERATIONDENIEDERROR = "OperationDenied.BizOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) StartSyncJobWithContext(ctx context.Context, request *StartSyncJobRequest) (response *StartSyncJobResponse, err error) {
    if request == nil {
        request = NewStartSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartSyncJobResponse()
    err = c.Send(request, response)
    return
}

func NewStopCompareRequest() (request *StopCompareRequest) {
    request = &StopCompareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "StopCompare")
    
    
    return
}

func NewStopCompareResponse() (response *StopCompareResponse) {
    response = &StopCompareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopCompare
// 终止一致性校验任务
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) StopCompare(request *StopCompareRequest) (response *StopCompareResponse, err error) {
    return c.StopCompareWithContext(context.Background(), request)
}

// StopCompare
// 终止一致性校验任务
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  MISSINGPARAMETER_MISSINGPARAMETERERROR = "MissingParameter.MissingParameterError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) StopCompareWithContext(ctx context.Context, request *StopCompareRequest) (response *StopCompareResponse, err error) {
    if request == nil {
        request = NewStopCompareRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopCompare require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopCompareResponse()
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
// 本接口（StopMigrateJob）用于终止数据迁移任务。
//
// 调用此接口后可通过查询迁移服务列表接口`DescribeMigrationJobs`来查询当前任务状态。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
func (c *Client) StopMigrateJob(request *StopMigrateJobRequest) (response *StopMigrateJobResponse, err error) {
    return c.StopMigrateJobWithContext(context.Background(), request)
}

// StopMigrateJob
// 本接口（StopMigrateJob）用于终止数据迁移任务。
//
// 调用此接口后可通过查询迁移服务列表接口`DescribeMigrationJobs`来查询当前任务状态。
//
// 可能返回的错误码:
//  AUTHFAILURE_AUTHFAILUREERROR = "AuthFailure.AuthFailureError"
//  DRYRUNOPERATION_DRYRUNOPERATIONERROR = "DryRunOperation.DryRunOperationError"
//  FAILEDOPERATION_FAILEDOPERATIONERROR = "FailedOperation.FailedOperationError"
//  FAILEDOPERATION_NOTALLOWOPERATION = "FailedOperation.NotAllowOperation"
//  FAILEDOPERATION_STATUSINCONFLICT = "FailedOperation.StatusInConflict"
//  INTERNALERROR_DATABASEERROR = "InternalError.DatabaseError"
//  INTERNALERROR_DUPLICATEJOB = "InternalError.DuplicateJob"
//  INTERNALERROR_INTERNALERRORERROR = "InternalError.InternalErrorError"
//  INTERNALERROR_PROTOCOLERROR = "InternalError.ProtocolError"
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETERERROR = "InvalidParameter.InvalidParameterError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUEERROR = "InvalidParameterValue.InvalidParameterValueError"
//  LIMITEXCEEDED_LIMITEXCEEDEDERROR = "LimitExceeded.LimitExceededError"
//  OPERATIONDENIED_JOBOPERATIONDENIEDERROR = "OperationDenied.JobOperationDeniedError"
//  OPERATIONDENIED_OPERATIONDENIEDERROR = "OperationDenied.OperationDeniedError"
//  REQUESTLIMITEXCEEDED_REQUESTLIMITEXCEEDEDERROR = "RequestLimitExceeded.RequestLimitExceededError"
//  RESOURCEINUSE_RESOURCEINUSEERROR = "ResourceInUse.ResourceInUseError"
//  RESOURCEINSUFFICIENT_RESOURCEINSUFFICIENTERROR = "ResourceInsufficient.ResourceInsufficientError"
//  RESOURCENOTFOUND_JOBNOTEXIST = "ResourceNotFound.JobNotExist"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
//  RESOURCEUNAVAILABLE_RESOURCEUNAVAILABLEERROR = "ResourceUnavailable.ResourceUnavailableError"
//  RESOURCESSOLDOUT_RESOURCESSOLDOUTERROR = "ResourcesSoldOut.ResourcesSoldOutError"
//  UNAUTHORIZEDOPERATION_NOTENOUGHPRIVILEGES = "UnauthorizedOperation.NotEnoughPrivileges"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZEDOPERATIONERROR = "UnauthorizedOperation.UnauthorizedOperationError"
//  UNKNOWNPARAMETER_UNKNOWNPARAMETERERROR = "UnknownParameter.UnknownParameterError"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDOPERATIONERROR = "UnsupportedOperation.UnsupportedOperationError"
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

func NewStopSyncJobRequest() (request *StopSyncJobRequest) {
    request = &StopSyncJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dts", APIVersion, "StopSyncJob")
    
    
    return
}

func NewStopSyncJobResponse() (response *StopSyncJobResponse) {
    response = &StopSyncJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopSyncJob
// 结束同步任务，操作后可通过查询同步任务信息接口DescribeSyncJobs，获取操作后的状态。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) StopSyncJob(request *StopSyncJobRequest) (response *StopSyncJobResponse, err error) {
    return c.StopSyncJobWithContext(context.Background(), request)
}

// StopSyncJob
// 结束同步任务，操作后可通过查询同步任务信息接口DescribeSyncJobs，获取操作后的状态。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  RESOURCENOTFOUND_RESOURCENOTFOUNDERROR = "ResourceNotFound.ResourceNotFoundError"
func (c *Client) StopSyncJobWithContext(ctx context.Context, request *StopSyncJobRequest) (response *StopSyncJobResponse, err error) {
    if request == nil {
        request = NewStopSyncJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopSyncJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopSyncJobResponse()
    err = c.Send(request, response)
    return
}
