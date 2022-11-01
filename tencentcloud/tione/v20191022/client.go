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

package v20191022

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-10-22"

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


func NewCreateCodeRepositoryRequest() (request *CreateCodeRepositoryRequest) {
    request = &CreateCodeRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "CreateCodeRepository")
    
    
    return
}

func NewCreateCodeRepositoryResponse() (response *CreateCodeRepositoryResponse) {
    response = &CreateCodeRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCodeRepository
// 创建存储库
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_KMSKEYNOTFOUND = "InvalidParameterValue.KmsKeyNotFound"
//  INVALIDPARAMETERVALUE_REPOSITORYURLINVALID = "InvalidParameterValue.RepositoryUrlInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE_BILLNOTACTIVATED = "ResourceUnavailable.BillNotActivated"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateCodeRepository(request *CreateCodeRepositoryRequest) (response *CreateCodeRepositoryResponse, err error) {
    return c.CreateCodeRepositoryWithContext(context.Background(), request)
}

// CreateCodeRepository
// 创建存储库
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_KMSKEYNOTFOUND = "InvalidParameterValue.KmsKeyNotFound"
//  INVALIDPARAMETERVALUE_REPOSITORYURLINVALID = "InvalidParameterValue.RepositoryUrlInvalid"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE_BILLNOTACTIVATED = "ResourceUnavailable.BillNotActivated"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateCodeRepositoryWithContext(ctx context.Context, request *CreateCodeRepositoryRequest) (response *CreateCodeRepositoryResponse, err error) {
    if request == nil {
        request = NewCreateCodeRepositoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCodeRepository require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCodeRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNotebookInstanceRequest() (request *CreateNotebookInstanceRequest) {
    request = &CreateNotebookInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "CreateNotebookInstance")
    
    
    return
}

func NewCreateNotebookInstanceResponse() (response *CreateNotebookInstanceResponse) {
    response = &CreateNotebookInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNotebookInstance
// 创建Notebook实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLSCONFIGREQUIRED = "InvalidParameterValue.ClsConfigRequired"
//  INVALIDPARAMETERVALUE_CODEREPONOTFOUND = "InvalidParameterValue.CodeRepoNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_LOGSETNOTFOUND = "InvalidParameterValue.LogSetNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_BILLNOTACTIVATED = "ResourceUnavailable.BillNotActivated"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateNotebookInstance(request *CreateNotebookInstanceRequest) (response *CreateNotebookInstanceResponse, err error) {
    return c.CreateNotebookInstanceWithContext(context.Background(), request)
}

// CreateNotebookInstance
// 创建Notebook实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLSCONFIGREQUIRED = "InvalidParameterValue.ClsConfigRequired"
//  INVALIDPARAMETERVALUE_CODEREPONOTFOUND = "InvalidParameterValue.CodeRepoNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_LOGSETNOTFOUND = "InvalidParameterValue.LogSetNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_BILLNOTACTIVATED = "ResourceUnavailable.BillNotActivated"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateNotebookInstanceWithContext(ctx context.Context, request *CreateNotebookInstanceRequest) (response *CreateNotebookInstanceResponse, err error) {
    if request == nil {
        request = NewCreateNotebookInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNotebookInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNotebookInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNotebookLifecycleScriptRequest() (request *CreateNotebookLifecycleScriptRequest) {
    request = &CreateNotebookLifecycleScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "CreateNotebookLifecycleScript")
    
    
    return
}

func NewCreateNotebookLifecycleScriptResponse() (response *CreateNotebookLifecycleScriptResponse) {
    response = &CreateNotebookLifecycleScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNotebookLifecycleScript
// 创建Notebook生命周期脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
func (c *Client) CreateNotebookLifecycleScript(request *CreateNotebookLifecycleScriptRequest) (response *CreateNotebookLifecycleScriptResponse, err error) {
    return c.CreateNotebookLifecycleScriptWithContext(context.Background(), request)
}

// CreateNotebookLifecycleScript
// 创建Notebook生命周期脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
func (c *Client) CreateNotebookLifecycleScriptWithContext(ctx context.Context, request *CreateNotebookLifecycleScriptRequest) (response *CreateNotebookLifecycleScriptResponse, err error) {
    if request == nil {
        request = NewCreateNotebookLifecycleScriptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNotebookLifecycleScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNotebookLifecycleScriptResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePresignedNotebookInstanceUrlRequest() (request *CreatePresignedNotebookInstanceUrlRequest) {
    request = &CreatePresignedNotebookInstanceUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "CreatePresignedNotebookInstanceUrl")
    
    
    return
}

func NewCreatePresignedNotebookInstanceUrlResponse() (response *CreatePresignedNotebookInstanceUrlResponse) {
    response = &CreatePresignedNotebookInstanceUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePresignedNotebookInstanceUrl
// 创建Notebook授权Url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTALIVE = "ResourceUnavailable.NotAlive"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreatePresignedNotebookInstanceUrl(request *CreatePresignedNotebookInstanceUrlRequest) (response *CreatePresignedNotebookInstanceUrlResponse, err error) {
    return c.CreatePresignedNotebookInstanceUrlWithContext(context.Background(), request)
}

// CreatePresignedNotebookInstanceUrl
// 创建Notebook授权Url
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTALIVE = "ResourceUnavailable.NotAlive"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreatePresignedNotebookInstanceUrlWithContext(ctx context.Context, request *CreatePresignedNotebookInstanceUrlRequest) (response *CreatePresignedNotebookInstanceUrlResponse, err error) {
    if request == nil {
        request = NewCreatePresignedNotebookInstanceUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePresignedNotebookInstanceUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePresignedNotebookInstanceUrlResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTrainingJobRequest() (request *CreateTrainingJobRequest) {
    request = &CreateTrainingJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "CreateTrainingJob")
    
    
    return
}

func NewCreateTrainingJobResponse() (response *CreateTrainingJobResponse) {
    response = &CreateTrainingJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTrainingJob
// 创建训练任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_FILESYSTEMNEEDVPCCONFIGSUPPORT = "InvalidParameterValue.FileSystemNeedVpcConfigSupport"
//  INVALIDPARAMETERVALUE_FILESYSTEMNUMLIMIT = "InvalidParameterValue.FileSystemNumLimit"
//  INVALIDPARAMETERVALUE_FILESYSTEMVPCNOTMATCH = "InvalidParameterValue.FileSystemVpcNotMatch"
//  INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"
//  INVALIDPARAMETERVALUE_IMAGENOTFOUND = "InvalidParameterValue.ImageNotFound"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPE = "InvalidParameterValue.InvalidInstanceType"
//  INVALIDPARAMETERVALUE_INVALIDTRAININGIMAGENAME = "InvalidParameterValue.InvalidTrainingImageName"
//  INVALIDPARAMETERVALUE_MPIPROCESSESPERHOSTTOOMUCH = "InvalidParameterValue.MpiProcessesPerHostTooMuch"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_TRAINCODENOTFOUND = "InvalidParameterValue.TrainCodeNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateTrainingJob(request *CreateTrainingJobRequest) (response *CreateTrainingJobResponse, err error) {
    return c.CreateTrainingJobWithContext(context.Background(), request)
}

// CreateTrainingJob
// 创建训练任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_FILESYSTEMNEEDVPCCONFIGSUPPORT = "InvalidParameterValue.FileSystemNeedVpcConfigSupport"
//  INVALIDPARAMETERVALUE_FILESYSTEMNUMLIMIT = "InvalidParameterValue.FileSystemNumLimit"
//  INVALIDPARAMETERVALUE_FILESYSTEMVPCNOTMATCH = "InvalidParameterValue.FileSystemVpcNotMatch"
//  INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"
//  INVALIDPARAMETERVALUE_IMAGENOTFOUND = "InvalidParameterValue.ImageNotFound"
//  INVALIDPARAMETERVALUE_INVALIDINSTANCETYPE = "InvalidParameterValue.InvalidInstanceType"
//  INVALIDPARAMETERVALUE_INVALIDTRAININGIMAGENAME = "InvalidParameterValue.InvalidTrainingImageName"
//  INVALIDPARAMETERVALUE_MPIPROCESSESPERHOSTTOOMUCH = "InvalidParameterValue.MpiProcessesPerHostTooMuch"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_TRAINCODENOTFOUND = "InvalidParameterValue.TrainCodeNotFound"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateTrainingJobWithContext(ctx context.Context, request *CreateTrainingJobRequest) (response *CreateTrainingJobResponse, err error) {
    if request == nil {
        request = NewCreateTrainingJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTrainingJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTrainingJobResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCodeRepositoryRequest() (request *DeleteCodeRepositoryRequest) {
    request = &DeleteCodeRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DeleteCodeRepository")
    
    
    return
}

func NewDeleteCodeRepositoryResponse() (response *DeleteCodeRepositoryResponse) {
    response = &DeleteCodeRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCodeRepository
// 删除存储库
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REPOBINDBYINSTANCE = "FailedOperation.RepoBindByInstance"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_BILLNOTACTIVATED = "ResourceUnavailable.BillNotActivated"
func (c *Client) DeleteCodeRepository(request *DeleteCodeRepositoryRequest) (response *DeleteCodeRepositoryResponse, err error) {
    return c.DeleteCodeRepositoryWithContext(context.Background(), request)
}

// DeleteCodeRepository
// 删除存储库
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REPOBINDBYINSTANCE = "FailedOperation.RepoBindByInstance"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_BILLNOTACTIVATED = "ResourceUnavailable.BillNotActivated"
func (c *Client) DeleteCodeRepositoryWithContext(ctx context.Context, request *DeleteCodeRepositoryRequest) (response *DeleteCodeRepositoryResponse, err error) {
    if request == nil {
        request = NewDeleteCodeRepositoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCodeRepository require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCodeRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNotebookInstanceRequest() (request *DeleteNotebookInstanceRequest) {
    request = &DeleteNotebookInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DeleteNotebookInstance")
    
    
    return
}

func NewDeleteNotebookInstanceResponse() (response *DeleteNotebookInstanceResponse) {
    response = &DeleteNotebookInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNotebookInstance
// 删除notebook实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteNotebookInstance(request *DeleteNotebookInstanceRequest) (response *DeleteNotebookInstanceResponse, err error) {
    return c.DeleteNotebookInstanceWithContext(context.Background(), request)
}

// DeleteNotebookInstance
// 删除notebook实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteNotebookInstanceWithContext(ctx context.Context, request *DeleteNotebookInstanceRequest) (response *DeleteNotebookInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteNotebookInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNotebookInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNotebookInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNotebookLifecycleScriptRequest() (request *DeleteNotebookLifecycleScriptRequest) {
    request = &DeleteNotebookLifecycleScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DeleteNotebookLifecycleScript")
    
    
    return
}

func NewDeleteNotebookLifecycleScriptResponse() (response *DeleteNotebookLifecycleScriptResponse) {
    response = &DeleteNotebookLifecycleScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNotebookLifecycleScript
// 删除Notebook生命周期脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteNotebookLifecycleScript(request *DeleteNotebookLifecycleScriptRequest) (response *DeleteNotebookLifecycleScriptResponse, err error) {
    return c.DeleteNotebookLifecycleScriptWithContext(context.Background(), request)
}

// DeleteNotebookLifecycleScript
// 删除Notebook生命周期脚本
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteNotebookLifecycleScriptWithContext(ctx context.Context, request *DeleteNotebookLifecycleScriptRequest) (response *DeleteNotebookLifecycleScriptResponse, err error) {
    if request == nil {
        request = NewDeleteNotebookLifecycleScriptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNotebookLifecycleScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNotebookLifecycleScriptResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCodeRepositoriesRequest() (request *DescribeCodeRepositoriesRequest) {
    request = &DescribeCodeRepositoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeCodeRepositories")
    
    
    return
}

func NewDescribeCodeRepositoriesResponse() (response *DescribeCodeRepositoriesResponse) {
    response = &DescribeCodeRepositoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCodeRepositories
// 查询存储库列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_BILLNOTACTIVATED = "ResourceUnavailable.BillNotActivated"
func (c *Client) DescribeCodeRepositories(request *DescribeCodeRepositoriesRequest) (response *DescribeCodeRepositoriesResponse, err error) {
    return c.DescribeCodeRepositoriesWithContext(context.Background(), request)
}

// DescribeCodeRepositories
// 查询存储库列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_BILLNOTACTIVATED = "ResourceUnavailable.BillNotActivated"
func (c *Client) DescribeCodeRepositoriesWithContext(ctx context.Context, request *DescribeCodeRepositoriesRequest) (response *DescribeCodeRepositoriesResponse, err error) {
    if request == nil {
        request = NewDescribeCodeRepositoriesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCodeRepositories require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCodeRepositoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCodeRepositoryRequest() (request *DescribeCodeRepositoryRequest) {
    request = &DescribeCodeRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeCodeRepository")
    
    
    return
}

func NewDescribeCodeRepositoryResponse() (response *DescribeCodeRepositoryResponse) {
    response = &DescribeCodeRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCodeRepository
// 查询存储库详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_BILLNOTACTIVATED = "ResourceUnavailable.BillNotActivated"
func (c *Client) DescribeCodeRepository(request *DescribeCodeRepositoryRequest) (response *DescribeCodeRepositoryResponse, err error) {
    return c.DescribeCodeRepositoryWithContext(context.Background(), request)
}

// DescribeCodeRepository
// 查询存储库详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_BILLNOTACTIVATED = "ResourceUnavailable.BillNotActivated"
func (c *Client) DescribeCodeRepositoryWithContext(ctx context.Context, request *DescribeCodeRepositoryRequest) (response *DescribeCodeRepositoryResponse, err error) {
    if request == nil {
        request = NewDescribeCodeRepositoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCodeRepository require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCodeRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotebookInstanceRequest() (request *DescribeNotebookInstanceRequest) {
    request = &DescribeNotebookInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeNotebookInstance")
    
    
    return
}

func NewDescribeNotebookInstanceResponse() (response *DescribeNotebookInstanceResponse) {
    response = &DescribeNotebookInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNotebookInstance
// 查询Notebook实例详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNotebookInstance(request *DescribeNotebookInstanceRequest) (response *DescribeNotebookInstanceResponse, err error) {
    return c.DescribeNotebookInstanceWithContext(context.Background(), request)
}

// DescribeNotebookInstance
// 查询Notebook实例详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNotebookInstanceWithContext(ctx context.Context, request *DescribeNotebookInstanceRequest) (response *DescribeNotebookInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeNotebookInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNotebookInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNotebookInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotebookInstancesRequest() (request *DescribeNotebookInstancesRequest) {
    request = &DescribeNotebookInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeNotebookInstances")
    
    
    return
}

func NewDescribeNotebookInstancesResponse() (response *DescribeNotebookInstancesResponse) {
    response = &DescribeNotebookInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNotebookInstances
// 查询Notebook实例列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeNotebookInstances(request *DescribeNotebookInstancesRequest) (response *DescribeNotebookInstancesResponse, err error) {
    return c.DescribeNotebookInstancesWithContext(context.Background(), request)
}

// DescribeNotebookInstances
// 查询Notebook实例列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeNotebookInstancesWithContext(ctx context.Context, request *DescribeNotebookInstancesRequest) (response *DescribeNotebookInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeNotebookInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNotebookInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNotebookInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotebookLifecycleScriptRequest() (request *DescribeNotebookLifecycleScriptRequest) {
    request = &DescribeNotebookLifecycleScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeNotebookLifecycleScript")
    
    
    return
}

func NewDescribeNotebookLifecycleScriptResponse() (response *DescribeNotebookLifecycleScriptResponse) {
    response = &DescribeNotebookLifecycleScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNotebookLifecycleScript
// 查看notebook生命周期脚本详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNotebookLifecycleScript(request *DescribeNotebookLifecycleScriptRequest) (response *DescribeNotebookLifecycleScriptResponse, err error) {
    return c.DescribeNotebookLifecycleScriptWithContext(context.Background(), request)
}

// DescribeNotebookLifecycleScript
// 查看notebook生命周期脚本详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNotebookLifecycleScriptWithContext(ctx context.Context, request *DescribeNotebookLifecycleScriptRequest) (response *DescribeNotebookLifecycleScriptResponse, err error) {
    if request == nil {
        request = NewDescribeNotebookLifecycleScriptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNotebookLifecycleScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNotebookLifecycleScriptResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotebookLifecycleScriptsRequest() (request *DescribeNotebookLifecycleScriptsRequest) {
    request = &DescribeNotebookLifecycleScriptsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeNotebookLifecycleScripts")
    
    
    return
}

func NewDescribeNotebookLifecycleScriptsResponse() (response *DescribeNotebookLifecycleScriptsResponse) {
    response = &DescribeNotebookLifecycleScriptsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNotebookLifecycleScripts
// 查看notebook生命周期脚本列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNotebookLifecycleScripts(request *DescribeNotebookLifecycleScriptsRequest) (response *DescribeNotebookLifecycleScriptsResponse, err error) {
    return c.DescribeNotebookLifecycleScriptsWithContext(context.Background(), request)
}

// DescribeNotebookLifecycleScripts
// 查看notebook生命周期脚本列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNotebookLifecycleScriptsWithContext(ctx context.Context, request *DescribeNotebookLifecycleScriptsRequest) (response *DescribeNotebookLifecycleScriptsResponse, err error) {
    if request == nil {
        request = NewDescribeNotebookLifecycleScriptsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNotebookLifecycleScripts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNotebookLifecycleScriptsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotebookSummaryRequest() (request *DescribeNotebookSummaryRequest) {
    request = &DescribeNotebookSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeNotebookSummary")
    
    
    return
}

func NewDescribeNotebookSummaryResponse() (response *DescribeNotebookSummaryResponse) {
    response = &DescribeNotebookSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNotebookSummary
// 查询Notebook概览数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeNotebookSummary(request *DescribeNotebookSummaryRequest) (response *DescribeNotebookSummaryResponse, err error) {
    return c.DescribeNotebookSummaryWithContext(context.Background(), request)
}

// DescribeNotebookSummary
// 查询Notebook概览数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeNotebookSummaryWithContext(ctx context.Context, request *DescribeNotebookSummaryRequest) (response *DescribeNotebookSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeNotebookSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNotebookSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNotebookSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrainingJobRequest() (request *DescribeTrainingJobRequest) {
    request = &DescribeTrainingJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeTrainingJob")
    
    
    return
}

func NewDescribeTrainingJobResponse() (response *DescribeTrainingJobResponse) {
    response = &DescribeTrainingJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrainingJob
// 查询训练任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeTrainingJob(request *DescribeTrainingJobRequest) (response *DescribeTrainingJobResponse, err error) {
    return c.DescribeTrainingJobWithContext(context.Background(), request)
}

// DescribeTrainingJob
// 查询训练任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeTrainingJobWithContext(ctx context.Context, request *DescribeTrainingJobRequest) (response *DescribeTrainingJobResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrainingJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrainingJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrainingJobsRequest() (request *DescribeTrainingJobsRequest) {
    request = &DescribeTrainingJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeTrainingJobs")
    
    
    return
}

func NewDescribeTrainingJobsResponse() (response *DescribeTrainingJobsResponse) {
    response = &DescribeTrainingJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrainingJobs
// 查询训练任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeTrainingJobs(request *DescribeTrainingJobsRequest) (response *DescribeTrainingJobsResponse, err error) {
    return c.DescribeTrainingJobsWithContext(context.Background(), request)
}

// DescribeTrainingJobs
// 查询训练任务列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeTrainingJobsWithContext(ctx context.Context, request *DescribeTrainingJobsRequest) (response *DescribeTrainingJobsResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingJobsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrainingJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrainingJobsResponse()
    err = c.Send(request, response)
    return
}

func NewStartNotebookInstanceRequest() (request *StartNotebookInstanceRequest) {
    request = &StartNotebookInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "StartNotebookInstance")
    
    
    return
}

func NewStartNotebookInstanceResponse() (response *StartNotebookInstanceResponse) {
    response = &StartNotebookInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartNotebookInstance
// 启动Notebook实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) StartNotebookInstance(request *StartNotebookInstanceRequest) (response *StartNotebookInstanceResponse, err error) {
    return c.StartNotebookInstanceWithContext(context.Background(), request)
}

// StartNotebookInstance
// 启动Notebook实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) StartNotebookInstanceWithContext(ctx context.Context, request *StartNotebookInstanceRequest) (response *StartNotebookInstanceResponse, err error) {
    if request == nil {
        request = NewStartNotebookInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartNotebookInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartNotebookInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewStopNotebookInstanceRequest() (request *StopNotebookInstanceRequest) {
    request = &StopNotebookInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "StopNotebookInstance")
    
    
    return
}

func NewStopNotebookInstanceResponse() (response *StopNotebookInstanceResponse) {
    response = &StopNotebookInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopNotebookInstance
// 停止Notebook实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopNotebookInstance(request *StopNotebookInstanceRequest) (response *StopNotebookInstanceResponse, err error) {
    return c.StopNotebookInstanceWithContext(context.Background(), request)
}

// StopNotebookInstance
// 停止Notebook实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopNotebookInstanceWithContext(ctx context.Context, request *StopNotebookInstanceRequest) (response *StopNotebookInstanceResponse, err error) {
    if request == nil {
        request = NewStopNotebookInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopNotebookInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopNotebookInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewStopTrainingJobRequest() (request *StopTrainingJobRequest) {
    request = &StopTrainingJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "StopTrainingJob")
    
    
    return
}

func NewStopTrainingJobResponse() (response *StopTrainingJobResponse) {
    response = &StopTrainingJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopTrainingJob
// 停止训练任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopTrainingJob(request *StopTrainingJobRequest) (response *StopTrainingJobResponse, err error) {
    return c.StopTrainingJobWithContext(context.Background(), request)
}

// StopTrainingJob
// 停止训练任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopTrainingJobWithContext(ctx context.Context, request *StopTrainingJobRequest) (response *StopTrainingJobResponse, err error) {
    if request == nil {
        request = NewStopTrainingJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopTrainingJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopTrainingJobResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCodeRepositoryRequest() (request *UpdateCodeRepositoryRequest) {
    request = &UpdateCodeRepositoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "UpdateCodeRepository")
    
    
    return
}

func NewUpdateCodeRepositoryResponse() (response *UpdateCodeRepositoryResponse) {
    response = &UpdateCodeRepositoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateCodeRepository
// 更新存储库
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_KMSKEYNOTFOUND = "InvalidParameterValue.KmsKeyNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_BILLNOTACTIVATED = "ResourceUnavailable.BillNotActivated"
func (c *Client) UpdateCodeRepository(request *UpdateCodeRepositoryRequest) (response *UpdateCodeRepositoryResponse, err error) {
    return c.UpdateCodeRepositoryWithContext(context.Background(), request)
}

// UpdateCodeRepository
// 更新存储库
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_KMSKEYNOTFOUND = "InvalidParameterValue.KmsKeyNotFound"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_BILLNOTACTIVATED = "ResourceUnavailable.BillNotActivated"
func (c *Client) UpdateCodeRepositoryWithContext(ctx context.Context, request *UpdateCodeRepositoryRequest) (response *UpdateCodeRepositoryResponse, err error) {
    if request == nil {
        request = NewUpdateCodeRepositoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateCodeRepository require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateCodeRepositoryResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateNotebookInstanceRequest() (request *UpdateNotebookInstanceRequest) {
    request = &UpdateNotebookInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "UpdateNotebookInstance")
    
    
    return
}

func NewUpdateNotebookInstanceResponse() (response *UpdateNotebookInstanceResponse) {
    response = &UpdateNotebookInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateNotebookInstance
// 更新Notebook实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLSCONFIGREQUIRED = "InvalidParameterValue.ClsConfigRequired"
//  INVALIDPARAMETERVALUE_CODEREPONOTFOUND = "InvalidParameterValue.CodeRepoNotFound"
//  INVALIDPARAMETERVALUE_LOGSETNOTFOUND = "InvalidParameterValue.LogSetNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
//  INVALIDPARAMETERVALUE_VOLUMESHRINKNOTALLOW = "InvalidParameterValue.VolumeShrinkNotAllow"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_BILLNOTACTIVATED = "ResourceUnavailable.BillNotActivated"
func (c *Client) UpdateNotebookInstance(request *UpdateNotebookInstanceRequest) (response *UpdateNotebookInstanceResponse, err error) {
    return c.UpdateNotebookInstanceWithContext(context.Background(), request)
}

// UpdateNotebookInstance
// 更新Notebook实例
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CLSCONFIGREQUIRED = "InvalidParameterValue.ClsConfigRequired"
//  INVALIDPARAMETERVALUE_CODEREPONOTFOUND = "InvalidParameterValue.CodeRepoNotFound"
//  INVALIDPARAMETERVALUE_LOGSETNOTFOUND = "InvalidParameterValue.LogSetNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
//  INVALIDPARAMETERVALUE_VOLUMESHRINKNOTALLOW = "InvalidParameterValue.VolumeShrinkNotAllow"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_BILLNOTACTIVATED = "ResourceUnavailable.BillNotActivated"
func (c *Client) UpdateNotebookInstanceWithContext(ctx context.Context, request *UpdateNotebookInstanceRequest) (response *UpdateNotebookInstanceResponse, err error) {
    if request == nil {
        request = NewUpdateNotebookInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateNotebookInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateNotebookInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateNotebookLifecycleScriptRequest() (request *UpdateNotebookLifecycleScriptRequest) {
    request = &UpdateNotebookLifecycleScriptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "UpdateNotebookLifecycleScript")
    
    
    return
}

func NewUpdateNotebookLifecycleScriptResponse() (response *UpdateNotebookLifecycleScriptResponse) {
    response = &UpdateNotebookLifecycleScriptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateNotebookLifecycleScript
// 更新notebook生命周期脚本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateNotebookLifecycleScript(request *UpdateNotebookLifecycleScriptRequest) (response *UpdateNotebookLifecycleScriptResponse, err error) {
    return c.UpdateNotebookLifecycleScriptWithContext(context.Background(), request)
}

// UpdateNotebookLifecycleScript
// 更新notebook生命周期脚本
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateNotebookLifecycleScriptWithContext(ctx context.Context, request *UpdateNotebookLifecycleScriptRequest) (response *UpdateNotebookLifecycleScriptResponse, err error) {
    if request == nil {
        request = NewUpdateNotebookLifecycleScriptRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateNotebookLifecycleScript require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateNotebookLifecycleScriptResponse()
    err = c.Send(request, response)
    return
}
