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

package v20190422

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-04-22"

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


func NewCreateJobRequest() (request *CreateJobRequest) {
    request = &CreateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "CreateJob")
    return
}

func NewCreateJobResponse() (response *CreateJobResponse) {
    response = &CreateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateJob
// 新建作业接口，一个 AppId 最多允许创建1000个作业
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CLUSTERNOTFOUND = "FailedOperation.ClusterNotFound"
//  FAILEDOPERATION_DUPLICATEDJOBNAME = "FailedOperation.DuplicatedJobName"
//  FAILEDOPERATION_RESOURCEINSUFFICIENT = "FailedOperation.ResourceInsufficient"
//  FAILEDOPERATION_USERNOTAUTHENTICATED = "FailedOperation.UserNotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CLUSTERID = "InvalidParameterValue.ClusterId"
//  INVALIDPARAMETERVALUE_CUMEM = "InvalidParameterValue.CuMem"
//  INVALIDPARAMETERVALUE_JOBNAME = "InvalidParameterValue.JobName"
//  INVALIDPARAMETERVALUE_JOBTYPECOMBINEWITHCLUSTERTYPE = "InvalidParameterValue.JobTypeCombineWithClusterType"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCEUNAVAILABLE_CLUSTER = "ResourceUnavailable.Cluster"
func (c *Client) CreateJob(request *CreateJobRequest) (response *CreateJobResponse, err error) {
    if request == nil {
        request = NewCreateJobRequest()
    }
    response = NewCreateJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateJobConfigRequest() (request *CreateJobConfigRequest) {
    request = &CreateJobConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "CreateJobConfig")
    return
}

func NewCreateJobConfigResponse() (response *CreateJobConfigResponse) {
    response = &CreateJobConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateJobConfig
// 创建作业配置，一个作业最多有100个配置版本
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_COSCLIENT = "InternalError.COSClient"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INTERNALERROR_SQLCODENOTFOUND = "InternalError.SqlCodeNotFound"
//  INTERNALERROR_STSNEWCLIENT = "InternalError.StsNewClient"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETER_UNSUPPORTEDFLINKCONF = "InvalidParameter.UnsupportedFlinkConf"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_JOBTYPECOMBINEWITHENTRYPOINTCLASS = "InvalidParameterValue.JobTypeCombineWithEntrypointClass"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_JOBCONFIG = "LimitExceeded.JobConfig"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_COSBUCKET = "ResourceNotFound.COSBucket"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  RESOURCENOTFOUND_JOBID = "ResourceNotFound.JobId"
//  UNSUPPORTEDOPERATION_INVALIDCHECKPOINTINTERVALERROR = "UnsupportedOperation.InvalidCheckpointIntervalError"
func (c *Client) CreateJobConfig(request *CreateJobConfigRequest) (response *CreateJobConfigResponse, err error) {
    if request == nil {
        request = NewCreateJobConfigRequest()
    }
    response = NewCreateJobConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateResourceRequest() (request *CreateResourceRequest) {
    request = &CreateResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "CreateResource")
    return
}

func NewCreateResourceResponse() (response *CreateResourceResponse) {
    response = &CreateResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateResource
// 创建资源接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSCLIENT = "InternalError.COSClient"
//  INTERNALERROR_DB = "InternalError.DB"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE_RESOURCENAMEALREADYEXISTS = "ResourceInUse.ResourceNameAlreadyExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CHECKRESOURCELOCEXISTS = "ResourceUnavailable.CheckResourceLocExists"
//  RESOURCEUNAVAILABLE_RESOURCELOCNOTEXISTS = "ResourceUnavailable.ResourceLocNotExists"
func (c *Client) CreateResource(request *CreateResourceRequest) (response *CreateResourceResponse, err error) {
    if request == nil {
        request = NewCreateResourceRequest()
    }
    response = NewCreateResourceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateResourceConfigRequest() (request *CreateResourceConfigRequest) {
    request = &CreateResourceConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "CreateResourceConfig")
    return
}

func NewCreateResourceConfigResponse() (response *CreateResourceConfigResponse) {
    response = &CreateResourceConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateResourceConfig
// 创建资源配置接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_RESOURCENOTEXIST = "InternalError.ResourceNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND_RESOURCENOTEXIST = "ResourceNotFound.ResourceNotExist"
func (c *Client) CreateResourceConfig(request *CreateResourceConfigRequest) (response *CreateResourceConfigResponse, err error) {
    if request == nil {
        request = NewCreateResourceConfigRequest()
    }
    response = NewCreateResourceConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteResourceConfigsRequest() (request *DeleteResourceConfigsRequest) {
    request = &DeleteResourceConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "DeleteResourceConfigs")
    return
}

func NewDeleteResourceConfigsResponse() (response *DeleteResourceConfigsResponse) {
    response = &DeleteResourceConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteResourceConfigs
// 删除资源版本
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_RESOURCECONFIGCANNOTDELETE = "InternalError.ResourceConfigCanNotDelete"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  INVALIDPARAMETERVALUE_RESOURCEIDSNOTFOUND = "InvalidParameterValue.ResourceIdsNotFound"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE_BEUSEBYSOMEJOBCONFIG = "ResourceUnavailable.BeUseBySomeJobConfig"
func (c *Client) DeleteResourceConfigs(request *DeleteResourceConfigsRequest) (response *DeleteResourceConfigsResponse, err error) {
    if request == nil {
        request = NewDeleteResourceConfigsRequest()
    }
    response = NewDeleteResourceConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteResourcesRequest() (request *DeleteResourcesRequest) {
    request = &DeleteResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "DeleteResources")
    return
}

func NewDeleteResourcesResponse() (response *DeleteResourcesResponse) {
    response = &DeleteResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteResources
// 删除资源接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteResources(request *DeleteResourcesRequest) (response *DeleteResourcesResponse, err error) {
    if request == nil {
        request = NewDeleteResourcesRequest()
    }
    response = NewDeleteResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTableConfigRequest() (request *DeleteTableConfigRequest) {
    request = &DeleteTableConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "DeleteTableConfig")
    return
}

func NewDeleteTableConfigResponse() (response *DeleteTableConfigResponse) {
    response = &DeleteTableConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTableConfig
// 删除作业表配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
func (c *Client) DeleteTableConfig(request *DeleteTableConfigRequest) (response *DeleteTableConfigResponse, err error) {
    if request == nil {
        request = NewDeleteTableConfigRequest()
    }
    response = NewDeleteTableConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobConfigsRequest() (request *DescribeJobConfigsRequest) {
    request = &DescribeJobConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "DescribeJobConfigs")
    return
}

func NewDescribeJobConfigsResponse() (response *DescribeJobConfigsResponse) {
    response = &DescribeJobConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeJobConfigs
// 查询作业配置列表，一次最多查询100个
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INVALIDPARAMETER_ILLEGALMAXPARALLELISM = "InvalidParameter.IllegalMaxParallelism"
//  INVALIDPARAMETER_MAXPARALLELISMTOOLARGE = "InvalidParameter.MaxParallelismTooLarge"
//  INVALIDPARAMETER_MAXPARALLELISMTOOSMALL = "InvalidParameter.MaxParallelismTooSmall"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
func (c *Client) DescribeJobConfigs(request *DescribeJobConfigsRequest) (response *DescribeJobConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeJobConfigsRequest()
    }
    response = NewDescribeJobConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobsRequest() (request *DescribeJobsRequest) {
    request = &DescribeJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "DescribeJobs")
    return
}

func NewDescribeJobsResponse() (response *DescribeJobsResponse) {
    response = &DescribeJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeJobs
// 查询作业
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  RESOURCENOTFOUND_JOBID = "ResourceNotFound.JobId"
//  RESOURCEUNAVAILABLE_GETJOBPUBLISHEDJOBCONFIG = "ResourceUnavailable.GetJobPublishedJobConfig"
func (c *Client) DescribeJobs(request *DescribeJobsRequest) (response *DescribeJobsResponse, err error) {
    if request == nil {
        request = NewDescribeJobsRequest()
    }
    response = NewDescribeJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceConfigsRequest() (request *DescribeResourceConfigsRequest) {
    request = &DescribeResourceConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "DescribeResourceConfigs")
    return
}

func NewDescribeResourceConfigsResponse() (response *DescribeResourceConfigsResponse) {
    response = &DescribeResourceConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResourceConfigs
// 描述资源配置接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INVALIDPARAMETER_APPIDRESOURCENOTMATCH = "InvalidParameter.AppIdResourceNotMatch"
//  INVALIDPARAMETER_UINRESOURCENOTMATCH = "InvalidParameter.UinResourceNotMatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_JOBIDVALUEERROR = "InvalidParameterValue.JobIdValueError"
//  RESOURCENOTFOUND_RESOURCECONFIG = "ResourceNotFound.ResourceConfig"
func (c *Client) DescribeResourceConfigs(request *DescribeResourceConfigsRequest) (response *DescribeResourceConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeResourceConfigsRequest()
    }
    response = NewDescribeResourceConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceRelatedJobsRequest() (request *DescribeResourceRelatedJobsRequest) {
    request = &DescribeResourceRelatedJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "DescribeResourceRelatedJobs")
    return
}

func NewDescribeResourceRelatedJobsResponse() (response *DescribeResourceRelatedJobsResponse) {
    response = &DescribeResourceRelatedJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResourceRelatedJobs
// 获取资源关联作业信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeResourceRelatedJobs(request *DescribeResourceRelatedJobsRequest) (response *DescribeResourceRelatedJobsResponse, err error) {
    if request == nil {
        request = NewDescribeResourceRelatedJobsRequest()
    }
    response = NewDescribeResourceRelatedJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcesRequest() (request *DescribeResourcesRequest) {
    request = &DescribeResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "DescribeResources")
    return
}

func NewDescribeResourcesResponse() (response *DescribeResourcesResponse) {
    response = &DescribeResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResources
// 描述资源接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCECONFIG = "ResourceNotFound.ResourceConfig"
func (c *Client) DescribeResources(request *DescribeResourcesRequest) (response *DescribeResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeResourcesRequest()
    }
    response = NewDescribeResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSystemResourcesRequest() (request *DescribeSystemResourcesRequest) {
    request = &DescribeSystemResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "DescribeSystemResources")
    return
}

func NewDescribeSystemResourcesResponse() (response *DescribeSystemResourcesResponse) {
    response = &DescribeSystemResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSystemResources
// 描述系统资源接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_FAILEDTOBESCRIBERESOURCES = "InternalError.FailedToBescribeResources"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE_FAILEDTOBESCRIBERESOURCES = "ResourceUnavailable.FailedToBescribeResources"
func (c *Client) DescribeSystemResources(request *DescribeSystemResourcesRequest) (response *DescribeSystemResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeSystemResourcesRequest()
    }
    response = NewDescribeSystemResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewRunJobsRequest() (request *RunJobsRequest) {
    request = &RunJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "RunJobs")
    return
}

func NewRunJobsResponse() (response *RunJobsResponse) {
    response = &RunJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RunJobs
// 批量启动或者恢复作业，批量操作数量上限20
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_JOBIDVALUEERROR = "InvalidParameterValue.JobIdValueError"
func (c *Client) RunJobs(request *RunJobsRequest) (response *RunJobsResponse, err error) {
    if request == nil {
        request = NewRunJobsRequest()
    }
    response = NewRunJobsResponse()
    err = c.Send(request, response)
    return
}

func NewStopJobsRequest() (request *StopJobsRequest) {
    request = &StopJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "StopJobs")
    return
}

func NewStopJobsResponse() (response *StopJobsResponse) {
    response = &StopJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopJobs
// 批量停止作业，批量操作数量上限为20
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  RESOURCENOTFOUND_JOBID = "ResourceNotFound.JobId"
//  RESOURCEUNAVAILABLE_NOTALLOWEDTOBESTOPORPAUSE = "ResourceUnavailable.NotAllowedToBeStopOrPause"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopJobs(request *StopJobsRequest) (response *StopJobsResponse, err error) {
    if request == nil {
        request = NewStopJobsRequest()
    }
    response = NewStopJobsResponse()
    err = c.Send(request, response)
    return
}
