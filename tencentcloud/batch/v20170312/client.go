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

package v20170312

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2017-03-12"

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


func NewAttachInstancesRequest() (request *AttachInstancesRequest) {
    request = &AttachInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "AttachInstances")
    
    
    return
}

func NewAttachInstancesResponse() (response *AttachInstancesResponse) {
    response = &AttachInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachInstances
// 此接口可将已存在实例添加到计算环境中。
//
// 实例需要满足如下条件：<br/>
//
// 1.实例不在批量计算系统中。<br/>
//
// 2.实例状态要求处于运行中。<br/>
//
// 3.支持预付费实例，按小时后付费实例，专享子机实例。不支持竞价实例。<br/>
//
// 
//
// 此接口会将加入到计算环境中的实例重设UserData和重装操作系统。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCEIDDUPLICATED = "InvalidParameterValue.InstanceIdDuplicated"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCESNOTALLOWTOATTACH = "UnsupportedOperation.InstancesNotAllowToAttach"
func (c *Client) AttachInstances(request *AttachInstancesRequest) (response *AttachInstancesResponse, err error) {
    return c.AttachInstancesWithContext(context.Background(), request)
}

// AttachInstances
// 此接口可将已存在实例添加到计算环境中。
//
// 实例需要满足如下条件：<br/>
//
// 1.实例不在批量计算系统中。<br/>
//
// 2.实例状态要求处于运行中。<br/>
//
// 3.支持预付费实例，按小时后付费实例，专享子机实例。不支持竞价实例。<br/>
//
// 
//
// 此接口会将加入到计算环境中的实例重设UserData和重装操作系统。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCEIDDUPLICATED = "InvalidParameterValue.InstanceIdDuplicated"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_INSTANCESNOTALLOWTOATTACH = "UnsupportedOperation.InstancesNotAllowToAttach"
func (c *Client) AttachInstancesWithContext(ctx context.Context, request *AttachInstancesRequest) (response *AttachInstancesResponse, err error) {
    if request == nil {
        request = NewAttachInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AttachInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewAttachInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateComputeEnvRequest() (request *CreateComputeEnvRequest) {
    request = &CreateComputeEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "CreateComputeEnv")
    
    
    return
}

func NewCreateComputeEnvResponse() (response *CreateComputeEnvResponse) {
    response = &CreateComputeEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateComputeEnv
// 用于创建计算环境
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INTERNALERROR_CALLTAGAPI = "InternalError.CallTagAPI"
//  INVALIDPARAMETER_CVMPARAMETERS = "InvalidParameter.CvmParameters"
//  INVALIDPARAMETER_ENVDESCRIPTIONTOOLONG = "InvalidParameter.EnvDescriptionTooLong"
//  INVALIDPARAMETER_ENVNAMETOOLONG = "InvalidParameter.EnvNameTooLong"
//  INVALIDPARAMETER_IMAGEIDMALFORMED = "InvalidParameter.ImageIdMalformed"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETER_NOTIFICATIONEVENTNAMEDUPLICATE = "InvalidParameter.NotificationEventNameDuplicate"
//  INVALIDPARAMETER_NOTIFICATIONTOPICNAME = "InvalidParameter.NotificationTopicName"
//  INVALIDPARAMETER_NOTIFICATIONTOPICNAMETOOLONG = "InvalidParameter.NotificationTopicNameTooLong"
//  INVALIDPARAMETERVALUE_COMPUTEENV = "InvalidParameterValue.ComputeEnv"
//  INVALIDPARAMETERVALUE_INSTANCETYPE = "InvalidParameterValue.InstanceType"
//  INVALIDPARAMETERVALUE_INVALIDDATATYPEANY = "InvalidParameterValue.InvalidDataTypeAny"
//  INVALIDPARAMETERVALUE_INVALIDZONEMISMATCHREGION = "InvalidParameterValue.InvalidZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_LOCALPATH = "InvalidParameterValue.LocalPath"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_REMOTESTORAGEPATH = "InvalidParameterValue.RemoteStoragePath"
//  INVALIDPARAMETERVALUE_REMOTESTORAGESCHEMETYPE = "InvalidParameterValue.RemoteStorageSchemeType"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  LIMITEXCEEDED_COMPUTEENVQUOTA = "LimitExceeded.ComputeEnvQuota"
//  LIMITEXCEEDED_CPUQUOTA = "LimitExceeded.CpuQuota"
//  UNAUTHORIZEDOPERATION_USERNOTALLOWEDTOUSEBATCH = "UnauthorizedOperation.UserNotAllowedToUseBatch"
func (c *Client) CreateComputeEnv(request *CreateComputeEnvRequest) (response *CreateComputeEnvResponse, err error) {
    return c.CreateComputeEnvWithContext(context.Background(), request)
}

// CreateComputeEnv
// 用于创建计算环境
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INTERNALERROR_CALLTAGAPI = "InternalError.CallTagAPI"
//  INVALIDPARAMETER_CVMPARAMETERS = "InvalidParameter.CvmParameters"
//  INVALIDPARAMETER_ENVDESCRIPTIONTOOLONG = "InvalidParameter.EnvDescriptionTooLong"
//  INVALIDPARAMETER_ENVNAMETOOLONG = "InvalidParameter.EnvNameTooLong"
//  INVALIDPARAMETER_IMAGEIDMALFORMED = "InvalidParameter.ImageIdMalformed"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETER_NOTIFICATIONEVENTNAMEDUPLICATE = "InvalidParameter.NotificationEventNameDuplicate"
//  INVALIDPARAMETER_NOTIFICATIONTOPICNAME = "InvalidParameter.NotificationTopicName"
//  INVALIDPARAMETER_NOTIFICATIONTOPICNAMETOOLONG = "InvalidParameter.NotificationTopicNameTooLong"
//  INVALIDPARAMETERVALUE_COMPUTEENV = "InvalidParameterValue.ComputeEnv"
//  INVALIDPARAMETERVALUE_INSTANCETYPE = "InvalidParameterValue.InstanceType"
//  INVALIDPARAMETERVALUE_INVALIDDATATYPEANY = "InvalidParameterValue.InvalidDataTypeAny"
//  INVALIDPARAMETERVALUE_INVALIDZONEMISMATCHREGION = "InvalidParameterValue.InvalidZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_LOCALPATH = "InvalidParameterValue.LocalPath"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_REMOTESTORAGEPATH = "InvalidParameterValue.RemoteStoragePath"
//  INVALIDPARAMETERVALUE_REMOTESTORAGESCHEMETYPE = "InvalidParameterValue.RemoteStorageSchemeType"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  LIMITEXCEEDED_COMPUTEENVQUOTA = "LimitExceeded.ComputeEnvQuota"
//  LIMITEXCEEDED_CPUQUOTA = "LimitExceeded.CpuQuota"
//  UNAUTHORIZEDOPERATION_USERNOTALLOWEDTOUSEBATCH = "UnauthorizedOperation.UserNotAllowedToUseBatch"
func (c *Client) CreateComputeEnvWithContext(ctx context.Context, request *CreateComputeEnvRequest) (response *CreateComputeEnvResponse, err error) {
    if request == nil {
        request = NewCreateComputeEnvRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateComputeEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateComputeEnvResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCpmComputeEnvRequest() (request *CreateCpmComputeEnvRequest) {
    request = &CreateCpmComputeEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "CreateCpmComputeEnv")
    
    
    return
}

func NewCreateCpmComputeEnvResponse() (response *CreateCpmComputeEnvResponse) {
    response = &CreateCpmComputeEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCpmComputeEnv
// 创建黑石计算环境
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCPMAPI = "InternalError.CallCpmAPI"
//  INTERNALERROR_CPMRESPONSEDATAEMPTY = "InternalError.CpmResponseDataEmpty"
//  INVALIDPARAMETER_ENVDESCRIPTIONTOOLONG = "InvalidParameter.EnvDescriptionTooLong"
//  INVALIDPARAMETER_ENVNAMETOOLONG = "InvalidParameter.EnvNameTooLong"
//  INVALIDPARAMETER_NOTIFICATIONEVENTNAMEDUPLICATE = "InvalidParameter.NotificationEventNameDuplicate"
//  INVALIDPARAMETER_NOTIFICATIONTOPICNAME = "InvalidParameter.NotificationTopicName"
//  INVALIDPARAMETER_NOTIFICATIONTOPICNAMETOOLONG = "InvalidParameter.NotificationTopicNameTooLong"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OSTYPEID = "InvalidParameterValue.OsTypeId"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTCPM = "InvalidParameterValue.RegionNotSupportCpm"
//  UNAUTHORIZEDOPERATION_USERNOTALLOWEDTOUSEBATCH = "UnauthorizedOperation.UserNotAllowedToUseBatch"
func (c *Client) CreateCpmComputeEnv(request *CreateCpmComputeEnvRequest) (response *CreateCpmComputeEnvResponse, err error) {
    return c.CreateCpmComputeEnvWithContext(context.Background(), request)
}

// CreateCpmComputeEnv
// 创建黑石计算环境
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCPMAPI = "InternalError.CallCpmAPI"
//  INTERNALERROR_CPMRESPONSEDATAEMPTY = "InternalError.CpmResponseDataEmpty"
//  INVALIDPARAMETER_ENVDESCRIPTIONTOOLONG = "InvalidParameter.EnvDescriptionTooLong"
//  INVALIDPARAMETER_ENVNAMETOOLONG = "InvalidParameter.EnvNameTooLong"
//  INVALIDPARAMETER_NOTIFICATIONEVENTNAMEDUPLICATE = "InvalidParameter.NotificationEventNameDuplicate"
//  INVALIDPARAMETER_NOTIFICATIONTOPICNAME = "InvalidParameter.NotificationTopicName"
//  INVALIDPARAMETER_NOTIFICATIONTOPICNAMETOOLONG = "InvalidParameter.NotificationTopicNameTooLong"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_OSTYPEID = "InvalidParameterValue.OsTypeId"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTCPM = "InvalidParameterValue.RegionNotSupportCpm"
//  UNAUTHORIZEDOPERATION_USERNOTALLOWEDTOUSEBATCH = "UnauthorizedOperation.UserNotAllowedToUseBatch"
func (c *Client) CreateCpmComputeEnvWithContext(ctx context.Context, request *CreateCpmComputeEnvRequest) (response *CreateCpmComputeEnvResponse, err error) {
    if request == nil {
        request = NewCreateCpmComputeEnvRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCpmComputeEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCpmComputeEnvResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTaskTemplateRequest() (request *CreateTaskTemplateRequest) {
    request = &CreateTaskTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "CreateTaskTemplate")
    
    
    return
}

func NewCreateTaskTemplateResponse() (response *CreateTaskTemplateResponse) {
    response = &CreateTaskTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTaskTemplate
// 用于创建任务模板
//
// 可能返回的错误码:
//  ALLOWEDONEATTRIBUTEINENVIDANDCOMPUTEENV = "AllowedOneAttributeInEnvIdAndComputeEnv"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLTAGAPI = "InternalError.CallTagAPI"
//  INVALIDPARAMETER_CVMPARAMETERS = "InvalidParameter.CvmParameters"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETER_IMAGEIDMALFORMED = "InvalidParameter.ImageIdMalformed"
//  INVALIDPARAMETER_TASKNAME = "InvalidParameter.TaskName"
//  INVALIDPARAMETER_TASKNAMETOOLONG = "InvalidParameter.TaskNameTooLong"
//  INVALIDPARAMETER_TASKTEMPLATENAME = "InvalidParameter.TaskTemplateName"
//  INVALIDPARAMETER_TASKTEMPLATENAMETOOLONG = "InvalidParameter.TaskTemplateNameTooLong"
//  INVALIDPARAMETERVALUE_COMPUTEENV = "InvalidParameterValue.ComputeEnv"
//  INVALIDPARAMETERVALUE_DEPENDENCENOTFOUNDTASKNAME = "InvalidParameterValue.DependenceNotFoundTaskName"
//  INVALIDPARAMETERVALUE_DEPENDENCEUNFEASIBLE = "InvalidParameterValue.DependenceUnfeasible"
//  INVALIDPARAMETERVALUE_INVALIDDATATYPEANY = "InvalidParameterValue.InvalidDataTypeAny"
//  INVALIDPARAMETERVALUE_LOCALPATH = "InvalidParameterValue.LocalPath"
//  INVALIDPARAMETERVALUE_MAXRETRYCOUNT = "InvalidParameterValue.MaxRetryCount"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_REMOTESTORAGEPATH = "InvalidParameterValue.RemoteStoragePath"
//  INVALIDPARAMETERVALUE_REMOTESTORAGESCHEMETYPE = "InvalidParameterValue.RemoteStorageSchemeType"
//  LIMITEXCEEDED_TASKTEMPLATEQUOTA = "LimitExceeded.TaskTemplateQuota"
func (c *Client) CreateTaskTemplate(request *CreateTaskTemplateRequest) (response *CreateTaskTemplateResponse, err error) {
    return c.CreateTaskTemplateWithContext(context.Background(), request)
}

// CreateTaskTemplate
// 用于创建任务模板
//
// 可能返回的错误码:
//  ALLOWEDONEATTRIBUTEINENVIDANDCOMPUTEENV = "AllowedOneAttributeInEnvIdAndComputeEnv"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLTAGAPI = "InternalError.CallTagAPI"
//  INVALIDPARAMETER_CVMPARAMETERS = "InvalidParameter.CvmParameters"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETER_IMAGEIDMALFORMED = "InvalidParameter.ImageIdMalformed"
//  INVALIDPARAMETER_TASKNAME = "InvalidParameter.TaskName"
//  INVALIDPARAMETER_TASKNAMETOOLONG = "InvalidParameter.TaskNameTooLong"
//  INVALIDPARAMETER_TASKTEMPLATENAME = "InvalidParameter.TaskTemplateName"
//  INVALIDPARAMETER_TASKTEMPLATENAMETOOLONG = "InvalidParameter.TaskTemplateNameTooLong"
//  INVALIDPARAMETERVALUE_COMPUTEENV = "InvalidParameterValue.ComputeEnv"
//  INVALIDPARAMETERVALUE_DEPENDENCENOTFOUNDTASKNAME = "InvalidParameterValue.DependenceNotFoundTaskName"
//  INVALIDPARAMETERVALUE_DEPENDENCEUNFEASIBLE = "InvalidParameterValue.DependenceUnfeasible"
//  INVALIDPARAMETERVALUE_INVALIDDATATYPEANY = "InvalidParameterValue.InvalidDataTypeAny"
//  INVALIDPARAMETERVALUE_LOCALPATH = "InvalidParameterValue.LocalPath"
//  INVALIDPARAMETERVALUE_MAXRETRYCOUNT = "InvalidParameterValue.MaxRetryCount"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_REMOTESTORAGEPATH = "InvalidParameterValue.RemoteStoragePath"
//  INVALIDPARAMETERVALUE_REMOTESTORAGESCHEMETYPE = "InvalidParameterValue.RemoteStorageSchemeType"
//  LIMITEXCEEDED_TASKTEMPLATEQUOTA = "LimitExceeded.TaskTemplateQuota"
func (c *Client) CreateTaskTemplateWithContext(ctx context.Context, request *CreateTaskTemplateRequest) (response *CreateTaskTemplateResponse, err error) {
    if request == nil {
        request = NewCreateTaskTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTaskTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTaskTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteComputeEnvRequest() (request *DeleteComputeEnvRequest) {
    request = &DeleteComputeEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DeleteComputeEnv")
    
    
    return
}

func NewDeleteComputeEnvResponse() (response *DeleteComputeEnvResponse) {
    response = &DeleteComputeEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteComputeEnv
// 用于删除计算环境
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCEPTOTHERREQUEST = "UnsupportedOperation.AcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVACCEPTOTHERREQUEST = "UnsupportedOperation.ComputeEnvAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVOPERATION = "UnsupportedOperation.ComputeEnvOperation"
//  UNSUPPORTEDOPERATION_COMPUTENODEFORBIDTERMINATE = "UnsupportedOperation.ComputeNodeForbidTerminate"
func (c *Client) DeleteComputeEnv(request *DeleteComputeEnvRequest) (response *DeleteComputeEnvResponse, err error) {
    return c.DeleteComputeEnvWithContext(context.Background(), request)
}

// DeleteComputeEnv
// 用于删除计算环境
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCEPTOTHERREQUEST = "UnsupportedOperation.AcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVACCEPTOTHERREQUEST = "UnsupportedOperation.ComputeEnvAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVOPERATION = "UnsupportedOperation.ComputeEnvOperation"
//  UNSUPPORTEDOPERATION_COMPUTENODEFORBIDTERMINATE = "UnsupportedOperation.ComputeNodeForbidTerminate"
func (c *Client) DeleteComputeEnvWithContext(ctx context.Context, request *DeleteComputeEnvRequest) (response *DeleteComputeEnvResponse, err error) {
    if request == nil {
        request = NewDeleteComputeEnvRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteComputeEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteComputeEnvResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteJobRequest() (request *DeleteJobRequest) {
    request = &DeleteJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DeleteJob")
    
    
    return
}

func NewDeleteJobResponse() (response *DeleteJobResponse) {
    response = &DeleteJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteJob
// 用于删除作业记录。
//
// 删除作业的效果相当于删除作业相关的所有信息。删除成功后，作业相关的所有信息都无法查询。
//
// 待删除的作业必须处于完结状态，且其内部包含的所有任务实例也必须处于完结状态，否则会禁止操作。完结状态，是指处于 SUCCEED 或 FAILED 状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCEINUSE_JOB = "ResourceInUse.Job"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteJob(request *DeleteJobRequest) (response *DeleteJobResponse, err error) {
    return c.DeleteJobWithContext(context.Background(), request)
}

// DeleteJob
// 用于删除作业记录。
//
// 删除作业的效果相当于删除作业相关的所有信息。删除成功后，作业相关的所有信息都无法查询。
//
// 待删除的作业必须处于完结状态，且其内部包含的所有任务实例也必须处于完结状态，否则会禁止操作。完结状态，是指处于 SUCCEED 或 FAILED 状态。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCEINUSE_JOB = "ResourceInUse.Job"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteJobWithContext(ctx context.Context, request *DeleteJobRequest) (response *DeleteJobResponse, err error) {
    if request == nil {
        request = NewDeleteJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteJobResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTaskTemplatesRequest() (request *DeleteTaskTemplatesRequest) {
    request = &DeleteTaskTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DeleteTaskTemplates")
    
    
    return
}

func NewDeleteTaskTemplatesResponse() (response *DeleteTaskTemplatesResponse) {
    response = &DeleteTaskTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTaskTemplates
// 用于删除任务模板信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TASKTEMPLATEIDMALFORMED = "InvalidParameter.TaskTemplateIdMalformed"
//  RESOURCENOTFOUND_TASKTEMPLATE = "ResourceNotFound.TaskTemplate"
func (c *Client) DeleteTaskTemplates(request *DeleteTaskTemplatesRequest) (response *DeleteTaskTemplatesResponse, err error) {
    return c.DeleteTaskTemplatesWithContext(context.Background(), request)
}

// DeleteTaskTemplates
// 用于删除任务模板信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TASKTEMPLATEIDMALFORMED = "InvalidParameter.TaskTemplateIdMalformed"
//  RESOURCENOTFOUND_TASKTEMPLATE = "ResourceNotFound.TaskTemplate"
func (c *Client) DeleteTaskTemplatesWithContext(ctx context.Context, request *DeleteTaskTemplatesRequest) (response *DeleteTaskTemplatesResponse, err error) {
    if request == nil {
        request = NewDeleteTaskTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTaskTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTaskTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAvailableCvmInstanceTypesRequest() (request *DescribeAvailableCvmInstanceTypesRequest) {
    request = &DescribeAvailableCvmInstanceTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeAvailableCvmInstanceTypes")
    
    
    return
}

func NewDescribeAvailableCvmInstanceTypesResponse() (response *DescribeAvailableCvmInstanceTypesResponse) {
    response = &DescribeAvailableCvmInstanceTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAvailableCvmInstanceTypes
// 查看可用的CVM机型配置信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDZONEMISMATCHREGION = "InvalidParameterValue.InvalidZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
func (c *Client) DescribeAvailableCvmInstanceTypes(request *DescribeAvailableCvmInstanceTypesRequest) (response *DescribeAvailableCvmInstanceTypesResponse, err error) {
    return c.DescribeAvailableCvmInstanceTypesWithContext(context.Background(), request)
}

// DescribeAvailableCvmInstanceTypes
// 查看可用的CVM机型配置信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDZONEMISMATCHREGION = "InvalidParameterValue.InvalidZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
func (c *Client) DescribeAvailableCvmInstanceTypesWithContext(ctx context.Context, request *DescribeAvailableCvmInstanceTypesRequest) (response *DescribeAvailableCvmInstanceTypesResponse, err error) {
    if request == nil {
        request = NewDescribeAvailableCvmInstanceTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAvailableCvmInstanceTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAvailableCvmInstanceTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComputeEnvRequest() (request *DescribeComputeEnvRequest) {
    request = &DescribeComputeEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeComputeEnv")
    
    
    return
}

func NewDescribeComputeEnvResponse() (response *DescribeComputeEnvResponse) {
    response = &DescribeComputeEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComputeEnv
// 用于查询计算环境的详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
func (c *Client) DescribeComputeEnv(request *DescribeComputeEnvRequest) (response *DescribeComputeEnvResponse, err error) {
    return c.DescribeComputeEnvWithContext(context.Background(), request)
}

// DescribeComputeEnv
// 用于查询计算环境的详细信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
func (c *Client) DescribeComputeEnvWithContext(ctx context.Context, request *DescribeComputeEnvRequest) (response *DescribeComputeEnvResponse, err error) {
    if request == nil {
        request = NewDescribeComputeEnvRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComputeEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComputeEnvResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComputeEnvActivitiesRequest() (request *DescribeComputeEnvActivitiesRequest) {
    request = &DescribeComputeEnvActivitiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeComputeEnvActivities")
    
    
    return
}

func NewDescribeComputeEnvActivitiesResponse() (response *DescribeComputeEnvActivitiesResponse) {
    response = &DescribeComputeEnvActivitiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComputeEnvActivities
// 用于查询计算环境的活动信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_COMPUTENODEIDMALFORMED = "InvalidParameter.ComputeNodeIdMalformed"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeComputeEnvActivities(request *DescribeComputeEnvActivitiesRequest) (response *DescribeComputeEnvActivitiesResponse, err error) {
    return c.DescribeComputeEnvActivitiesWithContext(context.Background(), request)
}

// DescribeComputeEnvActivities
// 用于查询计算环境的活动信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_COMPUTENODEIDMALFORMED = "InvalidParameter.ComputeNodeIdMalformed"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeComputeEnvActivitiesWithContext(ctx context.Context, request *DescribeComputeEnvActivitiesRequest) (response *DescribeComputeEnvActivitiesResponse, err error) {
    if request == nil {
        request = NewDescribeComputeEnvActivitiesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComputeEnvActivities require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComputeEnvActivitiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComputeEnvCreateInfoRequest() (request *DescribeComputeEnvCreateInfoRequest) {
    request = &DescribeComputeEnvCreateInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeComputeEnvCreateInfo")
    
    
    return
}

func NewDescribeComputeEnvCreateInfoResponse() (response *DescribeComputeEnvCreateInfoResponse) {
    response = &DescribeComputeEnvCreateInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComputeEnvCreateInfo
// 查看计算环境的创建信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
func (c *Client) DescribeComputeEnvCreateInfo(request *DescribeComputeEnvCreateInfoRequest) (response *DescribeComputeEnvCreateInfoResponse, err error) {
    return c.DescribeComputeEnvCreateInfoWithContext(context.Background(), request)
}

// DescribeComputeEnvCreateInfo
// 查看计算环境的创建信息。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
func (c *Client) DescribeComputeEnvCreateInfoWithContext(ctx context.Context, request *DescribeComputeEnvCreateInfoRequest) (response *DescribeComputeEnvCreateInfoResponse, err error) {
    if request == nil {
        request = NewDescribeComputeEnvCreateInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComputeEnvCreateInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComputeEnvCreateInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComputeEnvCreateInfosRequest() (request *DescribeComputeEnvCreateInfosRequest) {
    request = &DescribeComputeEnvCreateInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeComputeEnvCreateInfos")
    
    
    return
}

func NewDescribeComputeEnvCreateInfosResponse() (response *DescribeComputeEnvCreateInfosResponse) {
    response = &DescribeComputeEnvCreateInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComputeEnvCreateInfos
// 用于查看计算环境创建信息列表，包括名称、描述、类型、环境参数、通知及期望节点数等。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeComputeEnvCreateInfos(request *DescribeComputeEnvCreateInfosRequest) (response *DescribeComputeEnvCreateInfosResponse, err error) {
    return c.DescribeComputeEnvCreateInfosWithContext(context.Background(), request)
}

// DescribeComputeEnvCreateInfos
// 用于查看计算环境创建信息列表，包括名称、描述、类型、环境参数、通知及期望节点数等。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeComputeEnvCreateInfosWithContext(ctx context.Context, request *DescribeComputeEnvCreateInfosRequest) (response *DescribeComputeEnvCreateInfosResponse, err error) {
    if request == nil {
        request = NewDescribeComputeEnvCreateInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComputeEnvCreateInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComputeEnvCreateInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeComputeEnvsRequest() (request *DescribeComputeEnvsRequest) {
    request = &DescribeComputeEnvsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeComputeEnvs")
    
    
    return
}

func NewDescribeComputeEnvsResponse() (response *DescribeComputeEnvsResponse) {
    response = &DescribeComputeEnvsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeComputeEnvs
// 用于查看计算环境列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_RESOURCETYPE = "InvalidParameterValue.ResourceType"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeComputeEnvs(request *DescribeComputeEnvsRequest) (response *DescribeComputeEnvsResponse, err error) {
    return c.DescribeComputeEnvsWithContext(context.Background(), request)
}

// DescribeComputeEnvs
// 用于查看计算环境列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_RESOURCETYPE = "InvalidParameterValue.ResourceType"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeComputeEnvsWithContext(ctx context.Context, request *DescribeComputeEnvsRequest) (response *DescribeComputeEnvsResponse, err error) {
    if request == nil {
        request = NewDescribeComputeEnvsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeComputeEnvs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeComputeEnvsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCpmOsInfoRequest() (request *DescribeCpmOsInfoRequest) {
    request = &DescribeCpmOsInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeCpmOsInfo")
    
    
    return
}

func NewDescribeCpmOsInfoResponse() (response *DescribeCpmOsInfoResponse) {
    response = &DescribeCpmOsInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCpmOsInfo
// 创建黑石计算环境时，查询批量计算环境支持的黑石操作系统信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCPMAPI = "InternalError.CallCpmAPI"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTCPM = "InvalidParameterValue.RegionNotSupportCpm"
func (c *Client) DescribeCpmOsInfo(request *DescribeCpmOsInfoRequest) (response *DescribeCpmOsInfoResponse, err error) {
    return c.DescribeCpmOsInfoWithContext(context.Background(), request)
}

// DescribeCpmOsInfo
// 创建黑石计算环境时，查询批量计算环境支持的黑石操作系统信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCPMAPI = "InternalError.CallCpmAPI"
//  INVALIDPARAMETERVALUE_REGIONNOTSUPPORTCPM = "InvalidParameterValue.RegionNotSupportCpm"
func (c *Client) DescribeCpmOsInfoWithContext(ctx context.Context, request *DescribeCpmOsInfoRequest) (response *DescribeCpmOsInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCpmOsInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCpmOsInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCpmOsInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCvmZoneInstanceConfigInfosRequest() (request *DescribeCvmZoneInstanceConfigInfosRequest) {
    request = &DescribeCvmZoneInstanceConfigInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeCvmZoneInstanceConfigInfos")
    
    
    return
}

func NewDescribeCvmZoneInstanceConfigInfosResponse() (response *DescribeCvmZoneInstanceConfigInfosResponse) {
    response = &DescribeCvmZoneInstanceConfigInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCvmZoneInstanceConfigInfos
// 获取批量计算可用区机型配置信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDZONEMISMATCHREGION = "InvalidParameterValue.InvalidZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDBATCHINSTANCECHARGETYPE = "InvalidParameterValue.UnsupportedBatchInstanceChargeType"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
func (c *Client) DescribeCvmZoneInstanceConfigInfos(request *DescribeCvmZoneInstanceConfigInfosRequest) (response *DescribeCvmZoneInstanceConfigInfosResponse, err error) {
    return c.DescribeCvmZoneInstanceConfigInfosWithContext(context.Background(), request)
}

// DescribeCvmZoneInstanceConfigInfos
// 获取批量计算可用区机型配置信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_INVALIDZONEMISMATCHREGION = "InvalidParameterValue.InvalidZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDBATCHINSTANCECHARGETYPE = "InvalidParameterValue.UnsupportedBatchInstanceChargeType"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
func (c *Client) DescribeCvmZoneInstanceConfigInfosWithContext(ctx context.Context, request *DescribeCvmZoneInstanceConfigInfosRequest) (response *DescribeCvmZoneInstanceConfigInfosResponse, err error) {
    if request == nil {
        request = NewDescribeCvmZoneInstanceConfigInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCvmZoneInstanceConfigInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCvmZoneInstanceConfigInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceCategoriesRequest() (request *DescribeInstanceCategoriesRequest) {
    request = &DescribeInstanceCategoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeInstanceCategories")
    
    
    return
}

func NewDescribeInstanceCategoriesResponse() (response *DescribeInstanceCategoriesResponse) {
    response = &DescribeInstanceCategoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceCategories
// 目前对CVM现有实例族分类，每一类包含若干实例族。该接口用于查询实例分类信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceCategories(request *DescribeInstanceCategoriesRequest) (response *DescribeInstanceCategoriesResponse, err error) {
    return c.DescribeInstanceCategoriesWithContext(context.Background(), request)
}

// DescribeInstanceCategories
// 目前对CVM现有实例族分类，每一类包含若干实例族。该接口用于查询实例分类信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceCategoriesWithContext(ctx context.Context, request *DescribeInstanceCategoriesRequest) (response *DescribeInstanceCategoriesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceCategoriesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceCategories require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceCategoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobRequest() (request *DescribeJobRequest) {
    request = &DescribeJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeJob")
    
    
    return
}

func NewDescribeJobResponse() (response *DescribeJobResponse) {
    response = &DescribeJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeJob
// 用于查看一个作业的详细信息，包括内部任务（Task）和依赖（Dependence）信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
func (c *Client) DescribeJob(request *DescribeJobRequest) (response *DescribeJobResponse, err error) {
    return c.DescribeJobWithContext(context.Background(), request)
}

// DescribeJob
// 用于查看一个作业的详细信息，包括内部任务（Task）和依赖（Dependence）信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
func (c *Client) DescribeJobWithContext(ctx context.Context, request *DescribeJobRequest) (response *DescribeJobResponse, err error) {
    if request == nil {
        request = NewDescribeJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobSubmitInfoRequest() (request *DescribeJobSubmitInfoRequest) {
    request = &DescribeJobSubmitInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeJobSubmitInfo")
    
    
    return
}

func NewDescribeJobSubmitInfoResponse() (response *DescribeJobSubmitInfoResponse) {
    response = &DescribeJobSubmitInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeJobSubmitInfo
// 用于查询指定作业的提交信息，其返回内容包括 JobId 和 SubmitJob 接口中作为输入参数的作业提交信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
func (c *Client) DescribeJobSubmitInfo(request *DescribeJobSubmitInfoRequest) (response *DescribeJobSubmitInfoResponse, err error) {
    return c.DescribeJobSubmitInfoWithContext(context.Background(), request)
}

// DescribeJobSubmitInfo
// 用于查询指定作业的提交信息，其返回内容包括 JobId 和 SubmitJob 接口中作为输入参数的作业提交信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
func (c *Client) DescribeJobSubmitInfoWithContext(ctx context.Context, request *DescribeJobSubmitInfoRequest) (response *DescribeJobSubmitInfoResponse, err error) {
    if request == nil {
        request = NewDescribeJobSubmitInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJobSubmitInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobSubmitInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobsRequest() (request *DescribeJobsRequest) {
    request = &DescribeJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeJobs")
    
    
    return
}

func NewDescribeJobsResponse() (response *DescribeJobsResponse) {
    response = &DescribeJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeJobs
// 用于查询若干个作业的概览信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeJobs(request *DescribeJobsRequest) (response *DescribeJobsResponse, err error) {
    return c.DescribeJobsWithContext(context.Background(), request)
}

// DescribeJobs
// 用于查询若干个作业的概览信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeJobsWithContext(ctx context.Context, request *DescribeJobsRequest) (response *DescribeJobsResponse, err error) {
    if request == nil {
        request = NewDescribeJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
    request = &DescribeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeTask")
    
    
    return
}

func NewDescribeTaskResponse() (response *DescribeTaskResponse) {
    response = &DescribeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTask
// 用于查询指定任务的详细信息，包括任务内部的任务实例信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_TASK = "ResourceNotFound.Task"
func (c *Client) DescribeTask(request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
    return c.DescribeTaskWithContext(context.Background(), request)
}

// DescribeTask
// 用于查询指定任务的详细信息，包括任务内部的任务实例信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_TASK = "ResourceNotFound.Task"
func (c *Client) DescribeTaskWithContext(ctx context.Context, request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
    if request == nil {
        request = NewDescribeTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskLogsRequest() (request *DescribeTaskLogsRequest) {
    request = &DescribeTaskLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeTaskLogs")
    
    
    return
}

func NewDescribeTaskLogsResponse() (response *DescribeTaskLogsResponse) {
    response = &DescribeTaskLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskLogs
// 用于获取任务多个实例标准输出和标准错误日志。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  INVALIDPARAMETER_TASKNAME = "InvalidParameter.TaskName"
//  INVALIDPARAMETER_TASKNAMETOOLONG = "InvalidParameter.TaskNameTooLong"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  RESOURCENOTFOUND_TASK = "ResourceNotFound.Task"
//  RESOURCENOTFOUND_TASKINSTANCE = "ResourceNotFound.TaskInstance"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTaskLogs(request *DescribeTaskLogsRequest) (response *DescribeTaskLogsResponse, err error) {
    return c.DescribeTaskLogsWithContext(context.Background(), request)
}

// DescribeTaskLogs
// 用于获取任务多个实例标准输出和标准错误日志。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  INVALIDPARAMETER_TASKNAME = "InvalidParameter.TaskName"
//  INVALIDPARAMETER_TASKNAMETOOLONG = "InvalidParameter.TaskNameTooLong"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  RESOURCENOTFOUND_TASK = "ResourceNotFound.Task"
//  RESOURCENOTFOUND_TASKINSTANCE = "ResourceNotFound.TaskInstance"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTaskLogsWithContext(ctx context.Context, request *DescribeTaskLogsRequest) (response *DescribeTaskLogsResponse, err error) {
    if request == nil {
        request = NewDescribeTaskLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskTemplatesRequest() (request *DescribeTaskTemplatesRequest) {
    request = &DescribeTaskTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DescribeTaskTemplates")
    
    
    return
}

func NewDescribeTaskTemplatesResponse() (response *DescribeTaskTemplatesResponse) {
    response = &DescribeTaskTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskTemplates
// 用于查询任务模板信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETER_TASKTEMPLATEIDMALFORMED = "InvalidParameter.TaskTemplateIdMalformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTaskTemplates(request *DescribeTaskTemplatesRequest) (response *DescribeTaskTemplatesResponse, err error) {
    return c.DescribeTaskTemplatesWithContext(context.Background(), request)
}

// DescribeTaskTemplates
// 用于查询任务模板信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETER_TASKTEMPLATEIDMALFORMED = "InvalidParameter.TaskTemplateIdMalformed"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTaskTemplatesWithContext(ctx context.Context, request *DescribeTaskTemplatesRequest) (response *DescribeTaskTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeTaskTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDetachInstancesRequest() (request *DetachInstancesRequest) {
    request = &DetachInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "DetachInstances")
    
    
    return
}

func NewDetachInstancesResponse() (response *DetachInstancesResponse) {
    response = &DetachInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachInstances
// 将添加到计算环境中的实例从计算环境中移出。若是由批量计算自动创建的计算节点实例则不允许移出。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCEIDDUPLICATED = "InvalidParameterValue.InstanceIdDuplicated"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DetachInstances(request *DetachInstancesRequest) (response *DetachInstancesResponse, err error) {
    return c.DetachInstancesWithContext(context.Background(), request)
}

// DetachInstances
// 将添加到计算环境中的实例从计算环境中移出。若是由批量计算自动创建的计算节点实例则不允许移出。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETERVALUE_INSTANCEIDDUPLICATED = "InvalidParameterValue.InstanceIdDuplicated"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DetachInstancesWithContext(ctx context.Context, request *DetachInstancesRequest) (response *DetachInstancesResponse, err error) {
    if request == nil {
        request = NewDetachInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetachInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetachInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyComputeEnvRequest() (request *ModifyComputeEnvRequest) {
    request = &ModifyComputeEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "ModifyComputeEnv")
    
    
    return
}

func NewModifyComputeEnvResponse() (response *ModifyComputeEnvResponse) {
    response = &ModifyComputeEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyComputeEnv
// 用于修改计算环境属性
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCPMAPI = "InternalError.CallCpmAPI"
//  INVALIDPARAMETER_ENVDESCRIPTIONTOOLONG = "InvalidParameter.EnvDescriptionTooLong"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETER_ENVNAMETOOLONG = "InvalidParameter.EnvNameTooLong"
//  INVALIDPARAMETERVALUE_INSTANCETYPE = "InvalidParameterValue.InstanceType"
//  INVALIDPARAMETERVALUE_INSTANCETYPEDUPLICATE = "InvalidParameterValue.InstanceTypeDuplicate"
//  INVALIDPARAMETERVALUE_INSTANCETYPESEMPTY = "InvalidParameterValue.InstanceTypesEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  LIMITEXCEEDED_CPUQUOTA = "LimitExceeded.CpuQuota"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNSUPPORTEDOPERATION_ACCEPTOTHERREQUEST = "UnsupportedOperation.AcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVACCEPTOTHERREQUEST = "UnsupportedOperation.ComputeEnvAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTENODEFORBIDTERMINATE = "UnsupportedOperation.ComputeNodeForbidTerminate"
//  UNSUPPORTEDOPERATION_NOTENOUGHCOMPUTENODESTOTERMINATE = "UnsupportedOperation.NotEnoughComputeNodesToTerminate"
func (c *Client) ModifyComputeEnv(request *ModifyComputeEnvRequest) (response *ModifyComputeEnvResponse, err error) {
    return c.ModifyComputeEnvWithContext(context.Background(), request)
}

// ModifyComputeEnv
// 用于修改计算环境属性
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCPMAPI = "InternalError.CallCpmAPI"
//  INVALIDPARAMETER_ENVDESCRIPTIONTOOLONG = "InvalidParameter.EnvDescriptionTooLong"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETER_ENVNAMETOOLONG = "InvalidParameter.EnvNameTooLong"
//  INVALIDPARAMETERVALUE_INSTANCETYPE = "InvalidParameterValue.InstanceType"
//  INVALIDPARAMETERVALUE_INSTANCETYPEDUPLICATE = "InvalidParameterValue.InstanceTypeDuplicate"
//  INVALIDPARAMETERVALUE_INSTANCETYPESEMPTY = "InvalidParameterValue.InstanceTypesEmpty"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  LIMITEXCEEDED_CPUQUOTA = "LimitExceeded.CpuQuota"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNSUPPORTEDOPERATION_ACCEPTOTHERREQUEST = "UnsupportedOperation.AcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVACCEPTOTHERREQUEST = "UnsupportedOperation.ComputeEnvAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTENODEFORBIDTERMINATE = "UnsupportedOperation.ComputeNodeForbidTerminate"
//  UNSUPPORTEDOPERATION_NOTENOUGHCOMPUTENODESTOTERMINATE = "UnsupportedOperation.NotEnoughComputeNodesToTerminate"
func (c *Client) ModifyComputeEnvWithContext(ctx context.Context, request *ModifyComputeEnvRequest) (response *ModifyComputeEnvResponse, err error) {
    if request == nil {
        request = NewModifyComputeEnvRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyComputeEnv require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyComputeEnvResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTaskTemplateRequest() (request *ModifyTaskTemplateRequest) {
    request = &ModifyTaskTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "ModifyTaskTemplate")
    
    
    return
}

func NewModifyTaskTemplateResponse() (response *ModifyTaskTemplateResponse) {
    response = &ModifyTaskTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTaskTemplate
// 用于修改任务模板
//
// 可能返回的错误码:
//  ALLOWEDONEATTRIBUTEINENVIDANDCOMPUTEENV = "AllowedOneAttributeInEnvIdAndComputeEnv"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CVMPARAMETERS = "InvalidParameter.CvmParameters"
//  INVALIDPARAMETER_TASKNAME = "InvalidParameter.TaskName"
//  INVALIDPARAMETER_TASKNAMETOOLONG = "InvalidParameter.TaskNameTooLong"
//  INVALIDPARAMETER_TASKTEMPLATEDESCRIPTIONTOOLONG = "InvalidParameter.TaskTemplateDescriptionTooLong"
//  INVALIDPARAMETER_TASKTEMPLATEIDMALFORMED = "InvalidParameter.TaskTemplateIdMalformed"
//  INVALIDPARAMETER_TASKTEMPLATENAME = "InvalidParameter.TaskTemplateName"
//  INVALIDPARAMETER_TASKTEMPLATENAMETOOLONG = "InvalidParameter.TaskTemplateNameTooLong"
//  INVALIDPARAMETERATLEASTONEATTRIBUTE = "InvalidParameterAtLeastOneAttribute"
//  INVALIDPARAMETERVALUE_COMPUTEENV = "InvalidParameterValue.ComputeEnv"
//  INVALIDPARAMETERVALUE_DEPENDENCENOTFOUNDTASKNAME = "InvalidParameterValue.DependenceNotFoundTaskName"
//  INVALIDPARAMETERVALUE_DEPENDENCEUNFEASIBLE = "InvalidParameterValue.DependenceUnfeasible"
//  INVALIDPARAMETERVALUE_LOCALPATH = "InvalidParameterValue.LocalPath"
//  INVALIDPARAMETERVALUE_MAXRETRYCOUNT = "InvalidParameterValue.MaxRetryCount"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_REMOTESTORAGEPATH = "InvalidParameterValue.RemoteStoragePath"
//  INVALIDPARAMETERVALUE_REMOTESTORAGESCHEMETYPE = "InvalidParameterValue.RemoteStorageSchemeType"
//  RESOURCENOTFOUND_TASKTEMPLATE = "ResourceNotFound.TaskTemplate"
func (c *Client) ModifyTaskTemplate(request *ModifyTaskTemplateRequest) (response *ModifyTaskTemplateResponse, err error) {
    return c.ModifyTaskTemplateWithContext(context.Background(), request)
}

// ModifyTaskTemplate
// 用于修改任务模板
//
// 可能返回的错误码:
//  ALLOWEDONEATTRIBUTEINENVIDANDCOMPUTEENV = "AllowedOneAttributeInEnvIdAndComputeEnv"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_CVMPARAMETERS = "InvalidParameter.CvmParameters"
//  INVALIDPARAMETER_TASKNAME = "InvalidParameter.TaskName"
//  INVALIDPARAMETER_TASKNAMETOOLONG = "InvalidParameter.TaskNameTooLong"
//  INVALIDPARAMETER_TASKTEMPLATEDESCRIPTIONTOOLONG = "InvalidParameter.TaskTemplateDescriptionTooLong"
//  INVALIDPARAMETER_TASKTEMPLATEIDMALFORMED = "InvalidParameter.TaskTemplateIdMalformed"
//  INVALIDPARAMETER_TASKTEMPLATENAME = "InvalidParameter.TaskTemplateName"
//  INVALIDPARAMETER_TASKTEMPLATENAMETOOLONG = "InvalidParameter.TaskTemplateNameTooLong"
//  INVALIDPARAMETERATLEASTONEATTRIBUTE = "InvalidParameterAtLeastOneAttribute"
//  INVALIDPARAMETERVALUE_COMPUTEENV = "InvalidParameterValue.ComputeEnv"
//  INVALIDPARAMETERVALUE_DEPENDENCENOTFOUNDTASKNAME = "InvalidParameterValue.DependenceNotFoundTaskName"
//  INVALIDPARAMETERVALUE_DEPENDENCEUNFEASIBLE = "InvalidParameterValue.DependenceUnfeasible"
//  INVALIDPARAMETERVALUE_LOCALPATH = "InvalidParameterValue.LocalPath"
//  INVALIDPARAMETERVALUE_MAXRETRYCOUNT = "InvalidParameterValue.MaxRetryCount"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_REMOTESTORAGEPATH = "InvalidParameterValue.RemoteStoragePath"
//  INVALIDPARAMETERVALUE_REMOTESTORAGESCHEMETYPE = "InvalidParameterValue.RemoteStorageSchemeType"
//  RESOURCENOTFOUND_TASKTEMPLATE = "ResourceNotFound.TaskTemplate"
func (c *Client) ModifyTaskTemplateWithContext(ctx context.Context, request *ModifyTaskTemplateRequest) (response *ModifyTaskTemplateResponse, err error) {
    if request == nil {
        request = NewModifyTaskTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTaskTemplate require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTaskTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewRetryJobsRequest() (request *RetryJobsRequest) {
    request = &RetryJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "RetryJobs")
    
    
    return
}

func NewRetryJobsResponse() (response *RetryJobsResponse) {
    response = &RetryJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RetryJobs
// 用于重试作业中失败的任务实例。
//
// 当且仅当作业处于“FAILED”状态，支持重试操作。重试操作成功后，作业会按照“DAG”中指定的任务依赖关系，依次重试各个任务中失败的任务实例。任务实例的历史信息将被重置，如同首次运行一样，参与后续的调度和执行。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RetryJobs(request *RetryJobsRequest) (response *RetryJobsResponse, err error) {
    return c.RetryJobsWithContext(context.Background(), request)
}

// RetryJobs
// 用于重试作业中失败的任务实例。
//
// 当且仅当作业处于“FAILED”状态，支持重试操作。重试操作成功后，作业会按照“DAG”中指定的任务依赖关系，依次重试各个任务中失败的任务实例。任务实例的历史信息将被重置，如同首次运行一样，参与后续的调度和执行。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RetryJobsWithContext(ctx context.Context, request *RetryJobsRequest) (response *RetryJobsResponse, err error) {
    if request == nil {
        request = NewRetryJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RetryJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewRetryJobsResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitJobRequest() (request *SubmitJobRequest) {
    request = &SubmitJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "SubmitJob")
    
    
    return
}

func NewSubmitJobResponse() (response *SubmitJobResponse) {
    response = &SubmitJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SubmitJob
// 用于提交一个作业
//
// 可能返回的错误码:
//  ALLOWEDONEATTRIBUTEINENVIDANDCOMPUTEENV = "AllowedOneAttributeInEnvIdAndComputeEnv"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INTERNALERROR_CALLTAGAPI = "InternalError.CallTagAPI"
//  INVALIDPARAMETER_CVMPARAMETERS = "InvalidParameter.CvmParameters"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETER_IMAGEIDMALFORMED = "InvalidParameter.ImageIdMalformed"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETER_JOBDESCRIPTIONTOOLONG = "InvalidParameter.JobDescriptionTooLong"
//  INVALIDPARAMETER_JOBNAMETOOLONG = "InvalidParameter.JobNameTooLong"
//  INVALIDPARAMETER_NOTIFICATIONEVENTNAMEDUPLICATE = "InvalidParameter.NotificationEventNameDuplicate"
//  INVALIDPARAMETER_NOTIFICATIONTOPICNAME = "InvalidParameter.NotificationTopicName"
//  INVALIDPARAMETER_NOTIFICATIONTOPICNAMETOOLONG = "InvalidParameter.NotificationTopicNameTooLong"
//  INVALIDPARAMETER_TASKNAME = "InvalidParameter.TaskName"
//  INVALIDPARAMETER_TASKNAMETOOLONG = "InvalidParameter.TaskNameTooLong"
//  INVALIDPARAMETERVALUE_COMPUTEENV = "InvalidParameterValue.ComputeEnv"
//  INVALIDPARAMETERVALUE_DEPENDENCENOTFOUNDTASKNAME = "InvalidParameterValue.DependenceNotFoundTaskName"
//  INVALIDPARAMETERVALUE_DEPENDENCEUNFEASIBLE = "InvalidParameterValue.DependenceUnfeasible"
//  INVALIDPARAMETERVALUE_INSTANCETYPEDUPLICATE = "InvalidParameterValue.InstanceTypeDuplicate"
//  INVALIDPARAMETERVALUE_INVALIDDATATYPEANY = "InvalidParameterValue.InvalidDataTypeAny"
//  INVALIDPARAMETERVALUE_INVALIDZONEMISMATCHREGION = "InvalidParameterValue.InvalidZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_LOCALPATH = "InvalidParameterValue.LocalPath"
//  INVALIDPARAMETERVALUE_MAXRETRYCOUNT = "InvalidParameterValue.MaxRetryCount"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_NOTFLOAT = "InvalidParameterValue.NotFloat"
//  INVALIDPARAMETERVALUE_REMOTESTORAGEPATH = "InvalidParameterValue.RemoteStoragePath"
//  INVALIDPARAMETERVALUE_REMOTESTORAGESCHEMETYPE = "InvalidParameterValue.RemoteStorageSchemeType"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  LIMITEXCEEDED_JOBQUOTA = "LimitExceeded.JobQuota"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNAUTHORIZEDOPERATION_USERNOTALLOWEDTOUSEBATCH = "UnauthorizedOperation.UserNotAllowedToUseBatch"
func (c *Client) SubmitJob(request *SubmitJobRequest) (response *SubmitJobResponse, err error) {
    return c.SubmitJobWithContext(context.Background(), request)
}

// SubmitJob
// 用于提交一个作业
//
// 可能返回的错误码:
//  ALLOWEDONEATTRIBUTEINENVIDANDCOMPUTEENV = "AllowedOneAttributeInEnvIdAndComputeEnv"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CALLCVM = "InternalError.CallCvm"
//  INTERNALERROR_CALLTAGAPI = "InternalError.CallTagAPI"
//  INVALIDPARAMETER_CVMPARAMETERS = "InvalidParameter.CvmParameters"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  INVALIDPARAMETER_IMAGEIDMALFORMED = "InvalidParameter.ImageIdMalformed"
//  INVALIDPARAMETER_INVALIDPARAMETERCOMBINATION = "InvalidParameter.InvalidParameterCombination"
//  INVALIDPARAMETER_JOBDESCRIPTIONTOOLONG = "InvalidParameter.JobDescriptionTooLong"
//  INVALIDPARAMETER_JOBNAMETOOLONG = "InvalidParameter.JobNameTooLong"
//  INVALIDPARAMETER_NOTIFICATIONEVENTNAMEDUPLICATE = "InvalidParameter.NotificationEventNameDuplicate"
//  INVALIDPARAMETER_NOTIFICATIONTOPICNAME = "InvalidParameter.NotificationTopicName"
//  INVALIDPARAMETER_NOTIFICATIONTOPICNAMETOOLONG = "InvalidParameter.NotificationTopicNameTooLong"
//  INVALIDPARAMETER_TASKNAME = "InvalidParameter.TaskName"
//  INVALIDPARAMETER_TASKNAMETOOLONG = "InvalidParameter.TaskNameTooLong"
//  INVALIDPARAMETERVALUE_COMPUTEENV = "InvalidParameterValue.ComputeEnv"
//  INVALIDPARAMETERVALUE_DEPENDENCENOTFOUNDTASKNAME = "InvalidParameterValue.DependenceNotFoundTaskName"
//  INVALIDPARAMETERVALUE_DEPENDENCEUNFEASIBLE = "InvalidParameterValue.DependenceUnfeasible"
//  INVALIDPARAMETERVALUE_INSTANCETYPEDUPLICATE = "InvalidParameterValue.InstanceTypeDuplicate"
//  INVALIDPARAMETERVALUE_INVALIDDATATYPEANY = "InvalidParameterValue.InvalidDataTypeAny"
//  INVALIDPARAMETERVALUE_INVALIDZONEMISMATCHREGION = "InvalidParameterValue.InvalidZoneMismatchRegion"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_LOCALPATH = "InvalidParameterValue.LocalPath"
//  INVALIDPARAMETERVALUE_MAXRETRYCOUNT = "InvalidParameterValue.MaxRetryCount"
//  INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"
//  INVALIDPARAMETERVALUE_NOTFLOAT = "InvalidParameterValue.NotFloat"
//  INVALIDPARAMETERVALUE_REMOTESTORAGEPATH = "InvalidParameterValue.RemoteStoragePath"
//  INVALIDPARAMETERVALUE_REMOTESTORAGESCHEMETYPE = "InvalidParameterValue.RemoteStorageSchemeType"
//  INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  LIMITEXCEEDED_JOBQUOTA = "LimitExceeded.JobQuota"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  UNAUTHORIZEDOPERATION_USERNOTALLOWEDTOUSEBATCH = "UnauthorizedOperation.UserNotAllowedToUseBatch"
func (c *Client) SubmitJobWithContext(ctx context.Context, request *SubmitJobRequest) (response *SubmitJobResponse, err error) {
    if request == nil {
        request = NewSubmitJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitJobResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateComputeNodeRequest() (request *TerminateComputeNodeRequest) {
    request = &TerminateComputeNodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "TerminateComputeNode")
    
    
    return
}

func NewTerminateComputeNodeResponse() (response *TerminateComputeNodeResponse) {
    response = &TerminateComputeNodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TerminateComputeNode
// 用于销毁计算节点。
//
// 对于状态为CREATED、CREATION_FAILED、RUNNING和ABNORMAL的节点，允许销毁处理。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_COMPUTENODEIDMALFORMED = "InvalidParameter.ComputeNodeIdMalformed"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  RESOURCENOTFOUND_COMPUTENODE = "ResourceNotFound.ComputeNode"
//  UNSUPPORTEDOPERATION_ACCEPTOTHERREQUEST = "UnsupportedOperation.AcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVACCEPTOTHERREQUEST = "UnsupportedOperation.ComputeEnvAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVOPERATION = "UnsupportedOperation.ComputeEnvOperation"
//  UNSUPPORTEDOPERATION_COMPUTENODEFORBIDTERMINATE = "UnsupportedOperation.ComputeNodeForbidTerminate"
//  UNSUPPORTEDOPERATION_COMPUTENODEISTERMINATING = "UnsupportedOperation.ComputeNodeIsTerminating"
func (c *Client) TerminateComputeNode(request *TerminateComputeNodeRequest) (response *TerminateComputeNodeResponse, err error) {
    return c.TerminateComputeNodeWithContext(context.Background(), request)
}

// TerminateComputeNode
// 用于销毁计算节点。
//
// 对于状态为CREATED、CREATION_FAILED、RUNNING和ABNORMAL的节点，允许销毁处理。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_COMPUTENODEIDMALFORMED = "InvalidParameter.ComputeNodeIdMalformed"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  RESOURCENOTFOUND_COMPUTENODE = "ResourceNotFound.ComputeNode"
//  UNSUPPORTEDOPERATION_ACCEPTOTHERREQUEST = "UnsupportedOperation.AcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVACCEPTOTHERREQUEST = "UnsupportedOperation.ComputeEnvAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVOPERATION = "UnsupportedOperation.ComputeEnvOperation"
//  UNSUPPORTEDOPERATION_COMPUTENODEFORBIDTERMINATE = "UnsupportedOperation.ComputeNodeForbidTerminate"
//  UNSUPPORTEDOPERATION_COMPUTENODEISTERMINATING = "UnsupportedOperation.ComputeNodeIsTerminating"
func (c *Client) TerminateComputeNodeWithContext(ctx context.Context, request *TerminateComputeNodeRequest) (response *TerminateComputeNodeResponse, err error) {
    if request == nil {
        request = NewTerminateComputeNodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateComputeNode require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateComputeNodeResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateComputeNodesRequest() (request *TerminateComputeNodesRequest) {
    request = &TerminateComputeNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "TerminateComputeNodes")
    
    
    return
}

func NewTerminateComputeNodesResponse() (response *TerminateComputeNodesResponse) {
    response = &TerminateComputeNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TerminateComputeNodes
// 用于批量销毁计算节点，不允许重复销毁同一个节点。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_COMPUTENODEIDMALFORMED = "InvalidParameter.ComputeNodeIdMalformed"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  RESOURCENOTFOUND_COMPUTENODE = "ResourceNotFound.ComputeNode"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCEPTOTHERREQUEST = "UnsupportedOperation.AcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVACCEPTOTHERREQUEST = "UnsupportedOperation.ComputeEnvAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTENODEFORBIDTERMINATE = "UnsupportedOperation.ComputeNodeForbidTerminate"
func (c *Client) TerminateComputeNodes(request *TerminateComputeNodesRequest) (response *TerminateComputeNodesResponse, err error) {
    return c.TerminateComputeNodesWithContext(context.Background(), request)
}

// TerminateComputeNodes
// 用于批量销毁计算节点，不允许重复销毁同一个节点。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_COMPUTENODEIDMALFORMED = "InvalidParameter.ComputeNodeIdMalformed"
//  INVALIDPARAMETER_ENVIDMALFORMED = "InvalidParameter.EnvIdMalformed"
//  RESOURCENOTFOUND_COMPUTEENV = "ResourceNotFound.ComputeEnv"
//  RESOURCENOTFOUND_COMPUTENODE = "ResourceNotFound.ComputeNode"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_ACCEPTOTHERREQUEST = "UnsupportedOperation.AcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTEENVACCEPTOTHERREQUEST = "UnsupportedOperation.ComputeEnvAcceptOtherRequest"
//  UNSUPPORTEDOPERATION_COMPUTENODEFORBIDTERMINATE = "UnsupportedOperation.ComputeNodeForbidTerminate"
func (c *Client) TerminateComputeNodesWithContext(ctx context.Context, request *TerminateComputeNodesRequest) (response *TerminateComputeNodesResponse, err error) {
    if request == nil {
        request = NewTerminateComputeNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateComputeNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateComputeNodesResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateJobRequest() (request *TerminateJobRequest) {
    request = &TerminateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "TerminateJob")
    
    
    return
}

func NewTerminateJobResponse() (response *TerminateJobResponse) {
    response = &TerminateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TerminateJob
// 用于终止作业。
//
// 当作业处于“SUBMITTED”状态时，禁止终止操作；当作业处于“SUCCEED”状态时，终止操作不会生效。
//
// 终止作业是一个异步过程。整个终止过程的耗时和任务总数成正比。终止的效果相当于所含的所有任务实例进行TerminateTaskInstance操作。具体效果和用法可参考TerminateTaskInstance。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCEINUSE_JOB = "ResourceInUse.Job"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_TERMINATEOPERATIONWITHENVID = "UnsupportedOperation.TerminateOperationWithEnvId"
func (c *Client) TerminateJob(request *TerminateJobRequest) (response *TerminateJobResponse, err error) {
    return c.TerminateJobWithContext(context.Background(), request)
}

// TerminateJob
// 用于终止作业。
//
// 当作业处于“SUBMITTED”状态时，禁止终止操作；当作业处于“SUCCEED”状态时，终止操作不会生效。
//
// 终止作业是一个异步过程。整个终止过程的耗时和任务总数成正比。终止的效果相当于所含的所有任务实例进行TerminateTaskInstance操作。具体效果和用法可参考TerminateTaskInstance。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCEINUSE_JOB = "ResourceInUse.Job"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_TERMINATEOPERATIONWITHENVID = "UnsupportedOperation.TerminateOperationWithEnvId"
func (c *Client) TerminateJobWithContext(ctx context.Context, request *TerminateJobRequest) (response *TerminateJobResponse, err error) {
    if request == nil {
        request = NewTerminateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateJobResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateTaskInstanceRequest() (request *TerminateTaskInstanceRequest) {
    request = &TerminateTaskInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("batch", APIVersion, "TerminateTaskInstance")
    
    
    return
}

func NewTerminateTaskInstanceResponse() (response *TerminateTaskInstanceResponse) {
    response = &TerminateTaskInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TerminateTaskInstance
// 用于终止任务实例。
//
// 对于状态已经为“SUCCEED”和“FAILED”的任务实例，不做处理。
//
// 对于状态为“SUBMITTED”、“PENDING”、“RUNNABLE”的任务实例，状态将置为“FAILED”状态。
//
// 对于状态为“STARTING”、“RUNNING”、“FAILED_INTERRUPTED”的任务实例，分区两种情况：如果未显示指定计算环境，会先销毁CVM服务器，然后将状态置为“FAILED”，具有一定耗时；如果指定了计算环境EnvId，任务实例状态置为“FAILED”，并重启执行该任务的CVM服务器，具有一定的耗时。
//
// 对于状态为“FAILED_INTERRUPTED”的任务实例，终止操作实际成功之后，相关资源和配额才会释放。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_TASKINSTANCE = "ResourceNotFound.TaskInstance"
//  UNSUPPORTEDOPERATION_TERMINATEOPERATIONWITHENVID = "UnsupportedOperation.TerminateOperationWithEnvId"
func (c *Client) TerminateTaskInstance(request *TerminateTaskInstanceRequest) (response *TerminateTaskInstanceResponse, err error) {
    return c.TerminateTaskInstanceWithContext(context.Background(), request)
}

// TerminateTaskInstance
// 用于终止任务实例。
//
// 对于状态已经为“SUCCEED”和“FAILED”的任务实例，不做处理。
//
// 对于状态为“SUBMITTED”、“PENDING”、“RUNNABLE”的任务实例，状态将置为“FAILED”状态。
//
// 对于状态为“STARTING”、“RUNNING”、“FAILED_INTERRUPTED”的任务实例，分区两种情况：如果未显示指定计算环境，会先销毁CVM服务器，然后将状态置为“FAILED”，具有一定耗时；如果指定了计算环境EnvId，任务实例状态置为“FAILED”，并重启执行该任务的CVM服务器，具有一定的耗时。
//
// 对于状态为“FAILED_INTERRUPTED”的任务实例，终止操作实际成功之后，相关资源和配额才会释放。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_JOBIDMALFORMED = "InvalidParameter.JobIdMalformed"
//  RESOURCENOTFOUND_TASKINSTANCE = "ResourceNotFound.TaskInstance"
//  UNSUPPORTEDOPERATION_TERMINATEOPERATIONWITHENVID = "UnsupportedOperation.TerminateOperationWithEnvId"
func (c *Client) TerminateTaskInstanceWithContext(ctx context.Context, request *TerminateTaskInstanceRequest) (response *TerminateTaskInstanceResponse, err error) {
    if request == nil {
        request = NewTerminateTaskInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateTaskInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateTaskInstanceResponse()
    err = c.Send(request, response)
    return
}
