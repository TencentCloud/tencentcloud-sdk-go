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
    "context"
    "errors"
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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCheckSavepointRequest() (request *CheckSavepointRequest) {
    request = &CheckSavepointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("oceanus", APIVersion, "CheckSavepoint")
    
    
    return
}

func NewCheckSavepointResponse() (response *CheckSavepointResponse) {
    response = &CheckSavepointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckSavepoint
// 检查快照是否可用
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) CheckSavepoint(request *CheckSavepointRequest) (response *CheckSavepointResponse, err error) {
    return c.CheckSavepointWithContext(context.Background(), request)
}

// CheckSavepoint
// 检查快照是否可用
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) CheckSavepointWithContext(ctx context.Context, request *CheckSavepointRequest) (response *CheckSavepointResponse, err error) {
    if request == nil {
        request = NewCheckSavepointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckSavepoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckSavepointResponse()
    err = c.Send(request, response)
    return
}

func NewCopyJobsRequest() (request *CopyJobsRequest) {
    request = &CopyJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("oceanus", APIVersion, "CopyJobs")
    
    
    return
}

func NewCopyJobsResponse() (response *CopyJobsResponse) {
    response = &CopyJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CopyJobs
// 单条和批量复制作业
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLS = "InternalError.CLS"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_STSNEWCLIENT = "InternalError.StsNewClient"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) CopyJobs(request *CopyJobsRequest) (response *CopyJobsResponse, err error) {
    return c.CopyJobsWithContext(context.Background(), request)
}

// CopyJobs
// 单条和批量复制作业
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CLS = "InternalError.CLS"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_STSNEWCLIENT = "InternalError.StsNewClient"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) CopyJobsWithContext(ctx context.Context, request *CopyJobsRequest) (response *CopyJobsResponse, err error) {
    if request == nil {
        request = NewCopyJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewCopyJobsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFolderRequest() (request *CreateFolderRequest) {
    request = &CreateFolderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("oceanus", APIVersion, "CreateFolder")
    
    
    return
}

func NewCreateFolderResponse() (response *CreateFolderResponse) {
    response = &CreateFolderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFolder
// 作业列表页面新建文件夹请求
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) CreateFolder(request *CreateFolderRequest) (response *CreateFolderResponse, err error) {
    return c.CreateFolderWithContext(context.Background(), request)
}

// CreateFolder
// 作业列表页面新建文件夹请求
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) CreateFolderWithContext(ctx context.Context, request *CreateFolderRequest) (response *CreateFolderResponse, err error) {
    if request == nil {
        request = NewCreateFolderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFolder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFolderResponse()
    err = c.Send(request, response)
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEDJOBNAME = "FailedOperation.DuplicatedJobName"
//  FAILEDOPERATION_USERNOTAUTHENTICATED = "FailedOperation.UserNotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETERVALUE_CLUSTERID = "InvalidParameterValue.ClusterId"
//  INVALIDPARAMETERVALUE_CUMEM = "InvalidParameterValue.CuMem"
//  INVALIDPARAMETERVALUE_JOBNAME = "InvalidParameterValue.JobName"
//  INVALIDPARAMETERVALUE_JOBNAMEEXISTED = "InvalidParameterValue.JobNameExisted"
//  INVALIDPARAMETERVALUE_JOBTYPECOMBINEWITHCLUSTERTYPE = "InvalidParameterValue.JobTypeCombineWithClusterType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDCOMPOSITE = "InvalidParameterValue.UnSupportedComposite"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_JOB = "LimitExceeded.Job"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCEUNAVAILABLE_CLUSTER = "ResourceUnavailable.Cluster"
//  RESOURCEUNAVAILABLE_CLUSTERGROUPSTATUS = "ResourceUnavailable.ClusterGroupStatus"
//  RESOURCEUNAVAILABLE_REQCUMEM = "ResourceUnavailable.ReqCuMem"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) CreateJob(request *CreateJobRequest) (response *CreateJobResponse, err error) {
    return c.CreateJobWithContext(context.Background(), request)
}

// CreateJob
// 新建作业接口，一个 AppId 最多允许创建1000个作业
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATEDJOBNAME = "FailedOperation.DuplicatedJobName"
//  FAILEDOPERATION_USERNOTAUTHENTICATED = "FailedOperation.UserNotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INVALIDPARAMETER_INVALIDNAME = "InvalidParameter.InvalidName"
//  INVALIDPARAMETERVALUE_CLUSTERID = "InvalidParameterValue.ClusterId"
//  INVALIDPARAMETERVALUE_CUMEM = "InvalidParameterValue.CuMem"
//  INVALIDPARAMETERVALUE_JOBNAME = "InvalidParameterValue.JobName"
//  INVALIDPARAMETERVALUE_JOBNAMEEXISTED = "InvalidParameterValue.JobNameExisted"
//  INVALIDPARAMETERVALUE_JOBTYPECOMBINEWITHCLUSTERTYPE = "InvalidParameterValue.JobTypeCombineWithClusterType"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDCOMPOSITE = "InvalidParameterValue.UnSupportedComposite"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_JOB = "LimitExceeded.Job"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCEUNAVAILABLE_CLUSTER = "ResourceUnavailable.Cluster"
//  RESOURCEUNAVAILABLE_CLUSTERGROUPSTATUS = "ResourceUnavailable.ClusterGroupStatus"
//  RESOURCEUNAVAILABLE_REQCUMEM = "ResourceUnavailable.ReqCuMem"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) CreateJobWithContext(ctx context.Context, request *CreateJobRequest) (response *CreateJobResponse, err error) {
    if request == nil {
        request = NewCreateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateJob require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GRAMMARCHECKFAILURE = "FailedOperation.GrammarCheckFailure"
//  FAILEDOPERATION_PARSESQL = "FailedOperation.ParseSql"
//  INTERNALERROR_COSCLIENT = "InternalError.COSClient"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INTERNALERROR_SQLCODENOTFOUND = "InternalError.SqlCodeNotFound"
//  INTERNALERROR_STSNEWCLIENT = "InternalError.StsNewClient"
//  INVALIDPARAMETER_ILLEGALMAXPARALLELISM = "InvalidParameter.IllegalMaxParallelism"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETER_JOBCONFIGLOGCOLLECTPARAMERROR = "InvalidParameter.JobConfigLogCollectParamError"
//  INVALIDPARAMETER_MAXPARALLELISMTOOLARGE = "InvalidParameter.MaxParallelismTooLarge"
//  INVALIDPARAMETER_MAXPARALLELISMTOOSMALL = "InvalidParameter.MaxParallelismTooSmall"
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
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) CreateJobConfig(request *CreateJobConfigRequest) (response *CreateJobConfigResponse, err error) {
    return c.CreateJobConfigWithContext(context.Background(), request)
}

// CreateJobConfig
// 创建作业配置，一个作业最多有100个配置版本
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GRAMMARCHECKFAILURE = "FailedOperation.GrammarCheckFailure"
//  FAILEDOPERATION_PARSESQL = "FailedOperation.ParseSql"
//  INTERNALERROR_COSCLIENT = "InternalError.COSClient"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INTERNALERROR_SQLCODENOTFOUND = "InternalError.SqlCodeNotFound"
//  INTERNALERROR_STSNEWCLIENT = "InternalError.StsNewClient"
//  INVALIDPARAMETER_ILLEGALMAXPARALLELISM = "InvalidParameter.IllegalMaxParallelism"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETER_JOBCONFIGLOGCOLLECTPARAMERROR = "InvalidParameter.JobConfigLogCollectParamError"
//  INVALIDPARAMETER_MAXPARALLELISMTOOLARGE = "InvalidParameter.MaxParallelismTooLarge"
//  INVALIDPARAMETER_MAXPARALLELISMTOOSMALL = "InvalidParameter.MaxParallelismTooSmall"
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
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) CreateJobConfigWithContext(ctx context.Context, request *CreateJobConfigRequest) (response *CreateJobConfigResponse, err error) {
    if request == nil {
        request = NewCreateJobConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateJobConfig require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSCLIENT = "InternalError.COSClient"
//  INTERNALERROR_DB = "InternalError.DB"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE_RESOURCENAMEALREADYEXISTS = "ResourceInUse.ResourceNameAlreadyExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CHECKRESOURCELOCEXISTS = "ResourceUnavailable.CheckResourceLocExists"
//  RESOURCEUNAVAILABLE_RESOURCELOCNOTEXISTS = "ResourceUnavailable.ResourceLocNotExists"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) CreateResource(request *CreateResourceRequest) (response *CreateResourceResponse, err error) {
    return c.CreateResourceWithContext(context.Background(), request)
}

// CreateResource
// 创建资源接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_COSCLIENT = "InternalError.COSClient"
//  INTERNALERROR_DB = "InternalError.DB"
//  INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE_RESOURCENAMEALREADYEXISTS = "ResourceInUse.ResourceNameAlreadyExists"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_CHECKRESOURCELOCEXISTS = "ResourceUnavailable.CheckResourceLocExists"
//  RESOURCEUNAVAILABLE_RESOURCELOCNOTEXISTS = "ResourceUnavailable.ResourceLocNotExists"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) CreateResourceWithContext(ctx context.Context, request *CreateResourceRequest) (response *CreateResourceResponse, err error) {
    if request == nil {
        request = NewCreateResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateResource require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTEXIST = "ResourceNotFound.ResourceNotExist"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) CreateResourceConfig(request *CreateResourceConfigRequest) (response *CreateResourceConfigResponse, err error) {
    return c.CreateResourceConfigWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCENOTEXIST = "ResourceNotFound.ResourceNotExist"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) CreateResourceConfigWithContext(ctx context.Context, request *CreateResourceConfigRequest) (response *CreateResourceConfigResponse, err error) {
    if request == nil {
        request = NewCreateResourceConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateResourceConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateResourceConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteJobsRequest() (request *DeleteJobsRequest) {
    request = &DeleteJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("oceanus", APIVersion, "DeleteJobs")
    
    
    return
}

func NewDeleteJobsResponse() (response *DeleteJobsResponse) {
    response = &DeleteJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteJobs
// 批量删除作业接口，批量操作数量上限20
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INVALIDPARAMETERVALUE_JOBIDVALUEERROR = "InvalidParameterValue.JobIdValueError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  RESOURCENOTFOUND_JOBID = "ResourceNotFound.JobId"
//  RESOURCEUNAVAILABLE_NOTALLOWEDTOBEDELETED = "ResourceUnavailable.NotAllowedToBeDeleted"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DeleteJobs(request *DeleteJobsRequest) (response *DeleteJobsResponse, err error) {
    return c.DeleteJobsWithContext(context.Background(), request)
}

// DeleteJobs
// 批量删除作业接口，批量操作数量上限20
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INVALIDPARAMETERVALUE_JOBIDVALUEERROR = "InvalidParameterValue.JobIdValueError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  RESOURCENOTFOUND_JOBID = "ResourceNotFound.JobId"
//  RESOURCEUNAVAILABLE_NOTALLOWEDTOBEDELETED = "ResourceUnavailable.NotAllowedToBeDeleted"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DeleteJobsWithContext(ctx context.Context, request *DeleteJobsRequest) (response *DeleteJobsResponse, err error) {
    if request == nil {
        request = NewDeleteJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteJobsResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_RESOURCECONFIGCANNOTDELETE = "InternalError.ResourceConfigCanNotDelete"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  INVALIDPARAMETERVALUE_RESOURCEIDSNOTFOUND = "InvalidParameterValue.ResourceIdsNotFound"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_BEUSEBYSOMEJOBCONFIG = "ResourceUnavailable.BeUseBySomeJobConfig"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DeleteResourceConfigs(request *DeleteResourceConfigsRequest) (response *DeleteResourceConfigsResponse, err error) {
    return c.DeleteResourceConfigsWithContext(context.Background(), request)
}

// DeleteResourceConfigs
// 删除资源版本
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_RESOURCECONFIGCANNOTDELETE = "InternalError.ResourceConfigCanNotDelete"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  INVALIDPARAMETERVALUE_RESOURCEIDSNOTFOUND = "InvalidParameterValue.ResourceIdsNotFound"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_BEUSEBYSOMEJOBCONFIG = "ResourceUnavailable.BeUseBySomeJobConfig"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DeleteResourceConfigsWithContext(ctx context.Context, request *DeleteResourceConfigsRequest) (response *DeleteResourceConfigsResponse, err error) {
    if request == nil {
        request = NewDeleteResourceConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteResourceConfigs require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DeleteResources(request *DeleteResourcesRequest) (response *DeleteResourcesResponse, err error) {
    return c.DeleteResourcesWithContext(context.Background(), request)
}

// DeleteResources
// 删除资源接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INVALIDPARAMETER_INVALIDRESOURCEIDS = "InvalidParameter.InvalidResourceIds"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DeleteResourcesWithContext(ctx context.Context, request *DeleteResourcesRequest) (response *DeleteResourcesResponse, err error) {
    if request == nil {
        request = NewDeleteResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteResources require credential")
    }

    request.SetContext(ctx)
    
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
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
func (c *Client) DeleteTableConfig(request *DeleteTableConfigRequest) (response *DeleteTableConfigResponse, err error) {
    return c.DeleteTableConfigWithContext(context.Background(), request)
}

// DeleteTableConfig
// 删除作业表配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
func (c *Client) DeleteTableConfigWithContext(ctx context.Context, request *DeleteTableConfigRequest) (response *DeleteTableConfigResponse, err error) {
    if request == nil {
        request = NewDeleteTableConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTableConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTableConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
    request = &DescribeClustersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("oceanus", APIVersion, "DescribeClusters")
    
    
    return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
    response = &DescribeClustersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeClusters
// 查询集群
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETRESOURCETAGSBYRESOURCEIDS = "FailedOperation.GetResourceTagsByResourceIds"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CLUSTERIDS = "InvalidParameterValue.ClusterIds"
//  INVALIDPARAMETERVALUE_ORDERTYPE = "InvalidParameterValue.OrderType"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    return c.DescribeClustersWithContext(context.Background(), request)
}

// DescribeClusters
// 查询集群
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_GETRESOURCETAGSBYRESOURCEIDS = "FailedOperation.GetResourceTagsByResourceIds"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_CLUSTERIDS = "InvalidParameterValue.ClusterIds"
//  INVALIDPARAMETERVALUE_ORDERTYPE = "InvalidParameterValue.OrderType"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DescribeClustersWithContext(ctx context.Context, request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
    if request == nil {
        request = NewDescribeClustersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeClusters require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeClustersResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INVALIDPARAMETER_ILLEGALMAXPARALLELISM = "InvalidParameter.IllegalMaxParallelism"
//  INVALIDPARAMETER_MAXPARALLELISMTOOLARGE = "InvalidParameter.MaxParallelismTooLarge"
//  INVALIDPARAMETER_MAXPARALLELISMTOOSMALL = "InvalidParameter.MaxParallelismTooSmall"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DescribeJobConfigs(request *DescribeJobConfigsRequest) (response *DescribeJobConfigsResponse, err error) {
    return c.DescribeJobConfigsWithContext(context.Background(), request)
}

// DescribeJobConfigs
// 查询作业配置列表，一次最多查询100个
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INVALIDPARAMETER_ILLEGALMAXPARALLELISM = "InvalidParameter.IllegalMaxParallelism"
//  INVALIDPARAMETER_MAXPARALLELISMTOOLARGE = "InvalidParameter.MaxParallelismTooLarge"
//  INVALIDPARAMETER_MAXPARALLELISMTOOSMALL = "InvalidParameter.MaxParallelismTooSmall"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DescribeJobConfigsWithContext(ctx context.Context, request *DescribeJobConfigsRequest) (response *DescribeJobConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeJobConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJobConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobSavepointRequest() (request *DescribeJobSavepointRequest) {
    request = &DescribeJobSavepointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("oceanus", APIVersion, "DescribeJobSavepoint")
    
    
    return
}

func NewDescribeJobSavepointResponse() (response *DescribeJobSavepointResponse) {
    response = &DescribeJobSavepointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeJobSavepoint
// 查找Savepoint列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_JOBID = "ResourceNotFound.JobId"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DescribeJobSavepoint(request *DescribeJobSavepointRequest) (response *DescribeJobSavepointResponse, err error) {
    return c.DescribeJobSavepointWithContext(context.Background(), request)
}

// DescribeJobSavepoint
// 查找Savepoint列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  RESOURCENOTFOUND_JOBID = "ResourceNotFound.JobId"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DescribeJobSavepointWithContext(ctx context.Context, request *DescribeJobSavepointRequest) (response *DescribeJobSavepointResponse, err error) {
    if request == nil {
        request = NewDescribeJobSavepointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeJobSavepoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeJobSavepointResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_DB = "InternalError.DB"
//  RESOURCENOTFOUND_JOBID = "ResourceNotFound.JobId"
//  RESOURCEUNAVAILABLE_GETJOBPUBLISHEDJOBCONFIG = "ResourceUnavailable.GetJobPublishedJobConfig"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DescribeJobs(request *DescribeJobsRequest) (response *DescribeJobsResponse, err error) {
    return c.DescribeJobsWithContext(context.Background(), request)
}

// DescribeJobs
// 查询作业
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CAMCGWERROR = "InternalError.CamCgwError"
//  INTERNALERROR_DB = "InternalError.DB"
//  RESOURCENOTFOUND_JOBID = "ResourceNotFound.JobId"
//  RESOURCEUNAVAILABLE_GETJOBPUBLISHEDJOBCONFIG = "ResourceUnavailable.GetJobPublishedJobConfig"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INVALIDPARAMETER_APPIDRESOURCENOTMATCH = "InvalidParameter.AppIdResourceNotMatch"
//  INVALIDPARAMETER_UINRESOURCENOTMATCH = "InvalidParameter.UinResourceNotMatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_JOBIDVALUEERROR = "InvalidParameterValue.JobIdValueError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  RESOURCENOTFOUND_RESOURCECONFIG = "ResourceNotFound.ResourceConfig"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DescribeResourceConfigs(request *DescribeResourceConfigsRequest) (response *DescribeResourceConfigsResponse, err error) {
    return c.DescribeResourceConfigsWithContext(context.Background(), request)
}

// DescribeResourceConfigs
// 描述资源配置接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INVALIDPARAMETER_APPIDRESOURCENOTMATCH = "InvalidParameter.AppIdResourceNotMatch"
//  INVALIDPARAMETER_UINRESOURCENOTMATCH = "InvalidParameter.UinResourceNotMatch"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_JOBIDVALUEERROR = "InvalidParameterValue.JobIdValueError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  RESOURCENOTFOUND_RESOURCECONFIG = "ResourceNotFound.ResourceConfig"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DescribeResourceConfigsWithContext(ctx context.Context, request *DescribeResourceConfigsRequest) (response *DescribeResourceConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeResourceConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceConfigs require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DescribeResourceRelatedJobs(request *DescribeResourceRelatedJobsRequest) (response *DescribeResourceRelatedJobsResponse, err error) {
    return c.DescribeResourceRelatedJobsWithContext(context.Background(), request)
}

// DescribeResourceRelatedJobs
// 获取资源关联作业信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DescribeResourceRelatedJobsWithContext(ctx context.Context, request *DescribeResourceRelatedJobsRequest) (response *DescribeResourceRelatedJobsResponse, err error) {
    if request == nil {
        request = NewDescribeResourceRelatedJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceRelatedJobs require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCECONFIG = "ResourceNotFound.ResourceConfig"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DescribeResources(request *DescribeResourcesRequest) (response *DescribeResourcesResponse, err error) {
    return c.DescribeResourcesWithContext(context.Background(), request)
}

// DescribeResources
// 描述资源接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCECONFIG = "ResourceNotFound.ResourceConfig"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DescribeResourcesWithContext(ctx context.Context, request *DescribeResourcesRequest) (response *DescribeResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResources require credential")
    }

    request.SetContext(ctx)
    
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
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCEUNAVAILABLE_FAILEDTOBESCRIBERESOURCES = "ResourceUnavailable.FailedToBescribeResources"
func (c *Client) DescribeSystemResources(request *DescribeSystemResourcesRequest) (response *DescribeSystemResourcesResponse, err error) {
    return c.DescribeSystemResourcesWithContext(context.Background(), request)
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
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCEUNAVAILABLE_FAILEDTOBESCRIBERESOURCES = "ResourceUnavailable.FailedToBescribeResources"
func (c *Client) DescribeSystemResourcesWithContext(ctx context.Context, request *DescribeSystemResourcesRequest) (response *DescribeSystemResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeSystemResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSystemResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSystemResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTreeJobsRequest() (request *DescribeTreeJobsRequest) {
    request = &DescribeTreeJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("oceanus", APIVersion, "DescribeTreeJobs")
    
    
    return
}

func NewDescribeTreeJobsResponse() (response *DescribeTreeJobsResponse) {
    response = &DescribeTreeJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTreeJobs
// 生成树状作业显示结构
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DescribeTreeJobs(request *DescribeTreeJobsRequest) (response *DescribeTreeJobsResponse, err error) {
    return c.DescribeTreeJobsWithContext(context.Background(), request)
}

// DescribeTreeJobs
// 生成树状作业显示结构
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DescribeTreeJobsWithContext(ctx context.Context, request *DescribeTreeJobsRequest) (response *DescribeTreeJobsResponse, err error) {
    if request == nil {
        request = NewDescribeTreeJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTreeJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTreeJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTreeResourcesRequest() (request *DescribeTreeResourcesRequest) {
    request = &DescribeTreeResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("oceanus", APIVersion, "DescribeTreeResources")
    
    
    return
}

func NewDescribeTreeResourcesResponse() (response *DescribeTreeResourcesResponse) {
    response = &DescribeTreeResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTreeResources
// 查询树状结构资源列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATASOURCECONNECTIONFAILED = "FailedOperation.DataSourceConnectionFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DescribeTreeResources(request *DescribeTreeResourcesRequest) (response *DescribeTreeResourcesResponse, err error) {
    return c.DescribeTreeResourcesWithContext(context.Background(), request)
}

// DescribeTreeResources
// 查询树状结构资源列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DATASOURCECONNECTIONFAILED = "FailedOperation.DataSourceConnectionFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) DescribeTreeResourcesWithContext(ctx context.Context, request *DescribeTreeResourcesRequest) (response *DescribeTreeResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeTreeResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTreeResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTreeResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyJobRequest() (request *ModifyJobRequest) {
    request = &ModifyJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("oceanus", APIVersion, "ModifyJob")
    
    
    return
}

func NewModifyJobResponse() (response *ModifyJobResponse) {
    response = &ModifyJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyJob
// 更新作业属性，仅允许以下3种操作，不支持组合操作：
//
// (1)	更新作业名称
//
// (2)	更新作业备注 
//
// (3)	更新作业最大并行度
//
// 变更前提：WorkerCuNum<=MaxParallelism
//
// 如果MaxParallelism变小，不重启作业，待下一次重启生效
//
// 如果MaxParallelism变大，则要求入参RestartAllowed必须为True
//
// 假设作业运行状态，则先停止作业，再启动作业，中间状态丢失
//
// 假设作业暂停状态，则将作业更改为停止状态，中间状态丢失
//
// 
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DUPLICATEDJOBNAME = "FailedOperation.DuplicatedJobName"
//  INTERNALERROR_JOBINSTANCENOTFOUND = "InternalError.JobInstanceNotFound"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_JOBNAME = "InvalidParameterValue.JobName"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) ModifyJob(request *ModifyJobRequest) (response *ModifyJobResponse, err error) {
    return c.ModifyJobWithContext(context.Background(), request)
}

// ModifyJob
// 更新作业属性，仅允许以下3种操作，不支持组合操作：
//
// (1)	更新作业名称
//
// (2)	更新作业备注 
//
// (3)	更新作业最大并行度
//
// 变更前提：WorkerCuNum<=MaxParallelism
//
// 如果MaxParallelism变小，不重启作业，待下一次重启生效
//
// 如果MaxParallelism变大，则要求入参RestartAllowed必须为True
//
// 假设作业运行状态，则先停止作业，再启动作业，中间状态丢失
//
// 假设作业暂停状态，则将作业更改为停止状态，中间状态丢失
//
// 
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DUPLICATEDJOBNAME = "FailedOperation.DuplicatedJobName"
//  INTERNALERROR_JOBINSTANCENOTFOUND = "InternalError.JobInstanceNotFound"
//  INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"
//  INVALIDPARAMETERVALUE_JOBNAME = "InvalidParameterValue.JobName"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) ModifyJobWithContext(ctx context.Context, request *ModifyJobRequest) (response *ModifyJobResponse, err error) {
    if request == nil {
        request = NewModifyJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyJobResponse()
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_COSCLIENT = "InternalError.COSClient"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_FAILEDTOUPDATEJOB = "InternalError.FailedToUpdateJob"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INTERNALERROR_STSNEWCLIENT = "InternalError.StsNewClient"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDSTARTMODE = "InvalidParameterValue.InvalidStartMode"
//  INVALIDPARAMETERVALUE_JOBIDVALUEERROR = "InvalidParameterValue.JobIdValueError"
//  INVALIDPARAMETERVALUE_RUNJOBDESCRIPTIONSCOUNT = "InvalidParameterValue.RunJobDescriptionsCount"
//  INVALIDPARAMETERVALUE_RUNTYPE = "InvalidParameterValue.RunType"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCEINSUFFICIENT_CU = "ResourceInsufficient.CU"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_COSBUCKET = "ResourceNotFound.COSBucket"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  RESOURCENOTFOUND_JOBCONFIG = "ResourceNotFound.JobConfig"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_JOBRESOURCECONFIGNOTREADY = "ResourceUnavailable.JobResourceConfigNotReady"
//  RESOURCEUNAVAILABLE_NORUNNINGJOBINSTANCESFOUNDFORJOBID = "ResourceUnavailable.NoRunningJobInstancesFoundForJobId"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDSTARTMODE = "UnsupportedOperation.UnsupportedStartMode"
func (c *Client) RunJobs(request *RunJobsRequest) (response *RunJobsResponse, err error) {
    return c.RunJobsWithContext(context.Background(), request)
}

// RunJobs
// 批量启动或者恢复作业，批量操作数量上限20
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_COSCLIENT = "InternalError.COSClient"
//  INTERNALERROR_DB = "InternalError.DB"
//  INTERNALERROR_FAILEDTOUPDATEJOB = "InternalError.FailedToUpdateJob"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INTERNALERROR_STSNEWCLIENT = "InternalError.StsNewClient"
//  INVALIDPARAMETER_INVALIDCLUSTERID = "InvalidParameter.InvalidClusterId"
//  INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"
//  INVALIDPARAMETERVALUE_INVALIDSTARTMODE = "InvalidParameterValue.InvalidStartMode"
//  INVALIDPARAMETERVALUE_JOBIDVALUEERROR = "InvalidParameterValue.JobIdValueError"
//  INVALIDPARAMETERVALUE_RUNJOBDESCRIPTIONSCOUNT = "InvalidParameterValue.RunJobDescriptionsCount"
//  INVALIDPARAMETERVALUE_RUNTYPE = "InvalidParameterValue.RunType"
//  RESOURCEINUSE_INSTANCEINPROCESS = "ResourceInUse.InstanceInProcess"
//  RESOURCEINSUFFICIENT_CU = "ResourceInsufficient.CU"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_COSBUCKET = "ResourceNotFound.COSBucket"
//  RESOURCENOTFOUND_CLUSTERID = "ResourceNotFound.ClusterId"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  RESOURCENOTFOUND_JOBCONFIG = "ResourceNotFound.JobConfig"
//  RESOURCENOTFOUND_RESOURCE = "ResourceNotFound.Resource"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_JOBRESOURCECONFIGNOTREADY = "ResourceUnavailable.JobResourceConfigNotReady"
//  RESOURCEUNAVAILABLE_NORUNNINGJOBINSTANCESFOUNDFORJOBID = "ResourceUnavailable.NoRunningJobInstancesFoundForJobId"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
//  UNSUPPORTEDOPERATION_UNSUPPORTEDSTARTMODE = "UnsupportedOperation.UnsupportedStartMode"
func (c *Client) RunJobsWithContext(ctx context.Context, request *RunJobsRequest) (response *RunJobsResponse, err error) {
    if request == nil {
        request = NewRunJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunJobs require credential")
    }

    request.SetContext(ctx)
    
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
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INTERNALERROR_RESOURCENOTEXIST = "InternalError.ResourceNotExist"
//  INVALIDPARAMETERVALUE_UNKNOWNSTOPTYPE = "InvalidParameterValue.UnknownStopType"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_COSBUCKET = "ResourceNotFound.COSBucket"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  RESOURCENOTFOUND_RESOURCENOTEXIST = "ResourceNotFound.ResourceNotExist"
//  RESOURCEUNAVAILABLE_NORUNNINGJOBINSTANCESFOUNDFORJOBID = "ResourceUnavailable.NoRunningJobInstancesFoundForJobId"
//  RESOURCEUNAVAILABLE_NOTALLOWEDTOBESTOPORPAUSE = "ResourceUnavailable.NotAllowedToBeStopOrPause"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) StopJobs(request *StopJobsRequest) (response *StopJobsResponse, err error) {
    return c.StopJobsWithContext(context.Background(), request)
}

// StopJobs
// 批量停止作业，批量操作数量上限为20
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_LOGICERROR = "InternalError.LogicError"
//  INTERNALERROR_RESOURCENOTEXIST = "InternalError.ResourceNotExist"
//  INVALIDPARAMETERVALUE_UNKNOWNSTOPTYPE = "InvalidParameterValue.UnknownStopType"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND_COSBUCKET = "ResourceNotFound.COSBucket"
//  RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"
//  RESOURCENOTFOUND_RESOURCENOTEXIST = "ResourceNotFound.ResourceNotExist"
//  RESOURCEUNAVAILABLE_NORUNNINGJOBINSTANCESFOUNDFORJOBID = "ResourceUnavailable.NoRunningJobInstancesFoundForJobId"
//  RESOURCEUNAVAILABLE_NOTALLOWEDTOBESTOPORPAUSE = "ResourceUnavailable.NotAllowedToBeStopOrPause"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) StopJobsWithContext(ctx context.Context, request *StopJobsRequest) (response *StopJobsResponse, err error) {
    if request == nil {
        request = NewStopJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopJobsResponse()
    err = c.Send(request, response)
    return
}

func NewTriggerJobSavepointRequest() (request *TriggerJobSavepointRequest) {
    request = &TriggerJobSavepointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("oceanus", APIVersion, "TriggerJobSavepoint")
    
    
    return
}

func NewTriggerJobSavepointResponse() (response *TriggerJobSavepointResponse) {
    response = &TriggerJobSavepointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TriggerJobSavepoint
// 触发Savepoint
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) TriggerJobSavepoint(request *TriggerJobSavepointRequest) (response *TriggerJobSavepointResponse, err error) {
    return c.TriggerJobSavepointWithContext(context.Background(), request)
}

// TriggerJobSavepoint
// 触发Savepoint
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  UNSUPPORTEDOPERATION_NOPERMISSIONACCESS = "UnsupportedOperation.NoPermissionAccess"
func (c *Client) TriggerJobSavepointWithContext(ctx context.Context, request *TriggerJobSavepointRequest) (response *TriggerJobSavepointResponse, err error) {
    if request == nil {
        request = NewTriggerJobSavepointRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TriggerJobSavepoint require credential")
    }

    request.SetContext(ctx)
    
    response = NewTriggerJobSavepointResponse()
    err = c.Send(request, response)
    return
}
