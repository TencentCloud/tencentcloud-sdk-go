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

package v20211111

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-11-11"

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


func NewChatCompletionRequest() (request *ChatCompletionRequest) {
    request = &ChatCompletionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "ChatCompletion")
    
    
    return
}

func NewChatCompletionResponse() (response *ChatCompletionResponse) {
    response = &ChatCompletionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChatCompletion
// 该接口支持与两种类型大模型的聊天。
//
// 1. 与多行业多场景大模型的在线体验聊天
//
// 2. 与自行部署的开源大模型的聊天
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ChatCompletion(request *ChatCompletionRequest) (response *ChatCompletionResponse, err error) {
    return c.ChatCompletionWithContext(context.Background(), request)
}

// ChatCompletion
// 该接口支持与两种类型大模型的聊天。
//
// 1. 与多行业多场景大模型的在线体验聊天
//
// 2. 与自行部署的开源大模型的聊天
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ChatCompletionWithContext(ctx context.Context, request *ChatCompletionRequest) (response *ChatCompletionResponse, err error) {
    if request == nil {
        request = NewChatCompletionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChatCompletion require credential")
    }

    request.SetContext(ctx)
    
    response = NewChatCompletionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBatchModelAccTasksRequest() (request *CreateBatchModelAccTasksRequest) {
    request = &CreateBatchModelAccTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "CreateBatchModelAccTasks")
    
    
    return
}

func NewCreateBatchModelAccTasksResponse() (response *CreateBatchModelAccTasksResponse) {
    response = &CreateBatchModelAccTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBatchModelAccTasks
// 批量创建模型加速任务
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBatchModelAccTasks(request *CreateBatchModelAccTasksRequest) (response *CreateBatchModelAccTasksResponse, err error) {
    return c.CreateBatchModelAccTasksWithContext(context.Background(), request)
}

// CreateBatchModelAccTasks
// 批量创建模型加速任务
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBatchModelAccTasksWithContext(ctx context.Context, request *CreateBatchModelAccTasksRequest) (response *CreateBatchModelAccTasksResponse, err error) {
    if request == nil {
        request = NewCreateBatchModelAccTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBatchModelAccTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBatchModelAccTasksResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBatchTaskRequest() (request *CreateBatchTaskRequest) {
    request = &CreateBatchTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "CreateBatchTask")
    
    
    return
}

func NewCreateBatchTaskResponse() (response *CreateBatchTaskResponse) {
    response = &CreateBatchTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateBatchTask
// 创建跑批任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateBatchTask(request *CreateBatchTaskRequest) (response *CreateBatchTaskResponse, err error) {
    return c.CreateBatchTaskWithContext(context.Background(), request)
}

// CreateBatchTask
// 创建跑批任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateBatchTaskWithContext(ctx context.Context, request *CreateBatchTaskRequest) (response *CreateBatchTaskResponse, err error) {
    if request == nil {
        request = NewCreateBatchTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBatchTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBatchTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDatasetRequest() (request *CreateDatasetRequest) {
    request = &CreateDatasetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "CreateDataset")
    
    
    return
}

func NewCreateDatasetResponse() (response *CreateDatasetResponse) {
    response = &CreateDatasetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDataset
// 创建数据集
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VALIDATEERROR = "InvalidParameter.ValidateError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateDataset(request *CreateDatasetRequest) (response *CreateDatasetResponse, err error) {
    return c.CreateDatasetWithContext(context.Background(), request)
}

// CreateDataset
// 创建数据集
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VALIDATEERROR = "InvalidParameter.ValidateError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateDatasetWithContext(ctx context.Context, request *CreateDatasetRequest) (response *CreateDatasetResponse, err error) {
    if request == nil {
        request = NewCreateDatasetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDataset require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDatasetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateModelServiceRequest() (request *CreateModelServiceRequest) {
    request = &CreateModelServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "CreateModelService")
    
    
    return
}

func NewCreateModelServiceResponse() (response *CreateModelServiceResponse) {
    response = &CreateModelServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateModelService
// 用于创建、发布一个新的模型服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_INSUFFICIENTWHITELISTQUOTA = "FailedOperation.InsufficientWhitelistQuota"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MODELFILEINVALID = "InvalidParameter.ModelFileInvalid"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateModelService(request *CreateModelServiceRequest) (response *CreateModelServiceResponse, err error) {
    return c.CreateModelServiceWithContext(context.Background(), request)
}

// CreateModelService
// 用于创建、发布一个新的模型服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_INSUFFICIENTWHITELISTQUOTA = "FailedOperation.InsufficientWhitelistQuota"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MODELFILEINVALID = "InvalidParameter.ModelFileInvalid"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateModelServiceWithContext(ctx context.Context, request *CreateModelServiceRequest) (response *CreateModelServiceResponse, err error) {
    if request == nil {
        request = NewCreateModelServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateModelService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateModelServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNotebookRequest() (request *CreateNotebookRequest) {
    request = &CreateNotebookRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "CreateNotebook")
    
    
    return
}

func NewCreateNotebookResponse() (response *CreateNotebookResponse) {
    response = &CreateNotebookResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNotebook
// 创建Notebook
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDINGTAGSFAILED = "FailedOperation.BindingTagsFailed"
//  FAILEDOPERATION_FREEZEBILLFAILED = "FailedOperation.FreezeBillFailed"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYRESOURCEGROUPFAILED = "InternalError.QueryResourceGroupFailed"
//  INTERNALERROR_QUERYSUBNETINFOFAILED = "InternalError.QuerySubnetInfoFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CODEREPODUPLICATED = "InvalidParameterValue.CodeRepoDuplicated"
//  INVALIDPARAMETERVALUE_CODEREPONOTFOUND = "InvalidParameterValue.CodeRepoNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_LIFECYCLENOTFOUND = "InvalidParameterValue.LifecycleNotFound"
//  INVALIDPARAMETERVALUE_NOTALLOWEDTOCREATENOTEBOOKWITHBAREMETALRESOURCEGROUP = "InvalidParameterValue.NotAllowedToCreateNotebookWithBareMetalResourceGroup"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_BILLINGSTATUSRESOURCEINSUFFICIENT = "OperationDenied.BillingStatusResourceInsufficient"
//  OPERATIONDENIED_IPILLEGAL = "OperationDenied.IpIllegal"
//  OPERATIONDENIED_MIYINGBALANCEINSUFFICIENT = "OperationDenied.MIYINGBalanceInsufficient"
//  OPERATIONDENIED_NETWORKCIDRILLEGAL = "OperationDenied.NetworkCidrIllegal"
//  OPERATIONDENIED_RESOURCEGROUPINSUFFICIENT = "OperationDenied.ResourceGroupInsufficient"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateNotebook(request *CreateNotebookRequest) (response *CreateNotebookResponse, err error) {
    return c.CreateNotebookWithContext(context.Background(), request)
}

// CreateNotebook
// 创建Notebook
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDINGTAGSFAILED = "FailedOperation.BindingTagsFailed"
//  FAILEDOPERATION_FREEZEBILLFAILED = "FailedOperation.FreezeBillFailed"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYRESOURCEGROUPFAILED = "InternalError.QueryResourceGroupFailed"
//  INTERNALERROR_QUERYSUBNETINFOFAILED = "InternalError.QuerySubnetInfoFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CODEREPODUPLICATED = "InvalidParameterValue.CodeRepoDuplicated"
//  INVALIDPARAMETERVALUE_CODEREPONOTFOUND = "InvalidParameterValue.CodeRepoNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_LIFECYCLENOTFOUND = "InvalidParameterValue.LifecycleNotFound"
//  INVALIDPARAMETERVALUE_NOTALLOWEDTOCREATENOTEBOOKWITHBAREMETALRESOURCEGROUP = "InvalidParameterValue.NotAllowedToCreateNotebookWithBareMetalResourceGroup"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_BILLINGSTATUSRESOURCEINSUFFICIENT = "OperationDenied.BillingStatusResourceInsufficient"
//  OPERATIONDENIED_IPILLEGAL = "OperationDenied.IpIllegal"
//  OPERATIONDENIED_MIYINGBALANCEINSUFFICIENT = "OperationDenied.MIYINGBalanceInsufficient"
//  OPERATIONDENIED_NETWORKCIDRILLEGAL = "OperationDenied.NetworkCidrIllegal"
//  OPERATIONDENIED_RESOURCEGROUPINSUFFICIENT = "OperationDenied.ResourceGroupInsufficient"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateNotebookWithContext(ctx context.Context, request *CreateNotebookRequest) (response *CreateNotebookResponse, err error) {
    if request == nil {
        request = NewCreateNotebookRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNotebook require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNotebookResponse()
    err = c.Send(request, response)
    return
}

func NewCreateNotebookImageRequest() (request *CreateNotebookImageRequest) {
    request = &CreateNotebookImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "CreateNotebookImage")
    
    
    return
}

func NewCreateNotebookImageResponse() (response *CreateNotebookImageResponse) {
    response = &CreateNotebookImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateNotebookImage
// 保存镜像
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEIMAGEFAILED = "FailedOperation.CreateImageFailed"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_NOTSUPPORTEDTOCREATEIMAGE = "FailedOperation.NotSupportedToCreateImage"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateNotebookImage(request *CreateNotebookImageRequest) (response *CreateNotebookImageResponse, err error) {
    return c.CreateNotebookImageWithContext(context.Background(), request)
}

// CreateNotebookImage
// 保存镜像
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEIMAGEFAILED = "FailedOperation.CreateImageFailed"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_NOTSUPPORTEDTOCREATEIMAGE = "FailedOperation.NotSupportedToCreateImage"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateNotebookImageWithContext(ctx context.Context, request *CreateNotebookImageRequest) (response *CreateNotebookImageResponse, err error) {
    if request == nil {
        request = NewCreateNotebookImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNotebookImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNotebookImageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOptimizedModelRequest() (request *CreateOptimizedModelRequest) {
    request = &CreateOptimizedModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "CreateOptimizedModel")
    
    
    return
}

func NewCreateOptimizedModelResponse() (response *CreateOptimizedModelResponse) {
    response = &CreateOptimizedModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOptimizedModel
// 保存优化模型
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateOptimizedModel(request *CreateOptimizedModelRequest) (response *CreateOptimizedModelResponse, err error) {
    return c.CreateOptimizedModelWithContext(context.Background(), request)
}

// CreateOptimizedModel
// 保存优化模型
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateOptimizedModelWithContext(ctx context.Context, request *CreateOptimizedModelRequest) (response *CreateOptimizedModelResponse, err error) {
    if request == nil {
        request = NewCreateOptimizedModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOptimizedModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOptimizedModelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTrainingModelRequest() (request *CreateTrainingModelRequest) {
    request = &CreateTrainingModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "CreateTrainingModel")
    
    
    return
}

func NewCreateTrainingModelResponse() (response *CreateTrainingModelResponse) {
    response = &CreateTrainingModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTrainingModel
// 导入模型
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDINGTAGSFAILED = "FailedOperation.BindingTagsFailed"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_MOVEMODELDIRFAILED = "FailedOperation.MoveModelDirFailed"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateTrainingModel(request *CreateTrainingModelRequest) (response *CreateTrainingModelResponse, err error) {
    return c.CreateTrainingModelWithContext(context.Background(), request)
}

// CreateTrainingModel
// 导入模型
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDINGTAGSFAILED = "FailedOperation.BindingTagsFailed"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_MOVEMODELDIRFAILED = "FailedOperation.MoveModelDirFailed"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateTrainingModelWithContext(ctx context.Context, request *CreateTrainingModelRequest) (response *CreateTrainingModelResponse, err error) {
    if request == nil {
        request = NewCreateTrainingModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTrainingModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTrainingModelResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTrainingTaskRequest() (request *CreateTrainingTaskRequest) {
    request = &CreateTrainingTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "CreateTrainingTask")
    
    
    return
}

func NewCreateTrainingTaskResponse() (response *CreateTrainingTaskResponse) {
    response = &CreateTrainingTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTrainingTask
// 创建模型训练任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATENAMETASKISCREATING = "FailedOperation.DuplicateNameTaskIsCreating"
//  FAILEDOPERATION_FREEZEBILLFAILED = "FailedOperation.FreezeBillFailed"
//  FAILEDOPERATION_QUERYRESOURCESPECFAILED = "FailedOperation.QueryResourceSpecFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BINDINGTAGSFAILED = "InternalError.BindingTagsFailed"
//  INTERNALERROR_CREATETCRINSTANCETOKENFAILED = "InternalError.CreateTcrInstanceTokenFailed"
//  INTERNALERROR_QUERYHDFSINFOFAILED = "InternalError.QueryHDFSInfoFailed"
//  INTERNALERROR_QUERYRESOURCEGROUPFAILED = "InternalError.QueryResourceGroupFailed"
//  INTERNALERROR_QUERYRESOURCESPECFAILED = "InternalError.QueryResourceSpecFailed"
//  INTERNALERROR_QUERYSUBNETINFOFAILED = "InternalError.QuerySubnetInfoFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AIMARKETPUBLICALGOVERSIONNOTEXIST = "InvalidParameterValue.AIMarketPublicAlgoVersionNotExist"
//  INVALIDPARAMETERVALUE_DATASETNUMLIMITEXCEEDED = "InvalidParameterValue.DatasetNumLimitExceeded"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"
//  INVALIDPARAMETERVALUE_GETCFSMOUNTIPFAILED = "InvalidParameterValue.GetCFSMountIPFailed"
//  INVALIDPARAMETERVALUE_IMAGENOTFOUND = "InvalidParameterValue.ImageNotFound"
//  INVALIDPARAMETERVALUE_NOTALLOW = "InvalidParameterValue.NotAllow"
//  INVALIDPARAMETERVALUE_RDMACONFIGILLEGAL = "InvalidParameterValue.RDMAConfigIllegal"
//  INVALIDPARAMETERVALUE_RESOURCECONFIGILLEGAL = "InvalidParameterValue.ResourceConfigIllegal"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDDATACONFIG = "InvalidParameterValue.UnsupportedDataConfig"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_BILLINGEXCEPTION = "OperationDenied.BillingException"
//  OPERATIONDENIED_BILLINGSTATUSRESOURCEINSUFFICIENT = "OperationDenied.BillingStatusResourceInsufficient"
//  OPERATIONDENIED_IPILLEGAL = "OperationDenied.IpIllegal"
//  OPERATIONDENIED_MIYINGBALANCEINSUFFICIENT = "OperationDenied.MIYINGBalanceInsufficient"
//  OPERATIONDENIED_NETWORKCIDRILLEGAL = "OperationDenied.NetworkCidrIllegal"
//  OPERATIONDENIED_NOTALLOW = "OperationDenied.NotAllow"
//  OPERATIONDENIED_RESOURCEGROUPINSUFFICIENT = "OperationDenied.ResourceGroupInsufficient"
//  OPERATIONDENIED_SUBNETILLEGAL = "OperationDenied.SubnetIllegal"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateTrainingTask(request *CreateTrainingTaskRequest) (response *CreateTrainingTaskResponse, err error) {
    return c.CreateTrainingTaskWithContext(context.Background(), request)
}

// CreateTrainingTask
// 创建模型训练任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATENAMETASKISCREATING = "FailedOperation.DuplicateNameTaskIsCreating"
//  FAILEDOPERATION_FREEZEBILLFAILED = "FailedOperation.FreezeBillFailed"
//  FAILEDOPERATION_QUERYRESOURCESPECFAILED = "FailedOperation.QueryResourceSpecFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BINDINGTAGSFAILED = "InternalError.BindingTagsFailed"
//  INTERNALERROR_CREATETCRINSTANCETOKENFAILED = "InternalError.CreateTcrInstanceTokenFailed"
//  INTERNALERROR_QUERYHDFSINFOFAILED = "InternalError.QueryHDFSInfoFailed"
//  INTERNALERROR_QUERYRESOURCEGROUPFAILED = "InternalError.QueryResourceGroupFailed"
//  INTERNALERROR_QUERYRESOURCESPECFAILED = "InternalError.QueryResourceSpecFailed"
//  INTERNALERROR_QUERYSUBNETINFOFAILED = "InternalError.QuerySubnetInfoFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AIMARKETPUBLICALGOVERSIONNOTEXIST = "InvalidParameterValue.AIMarketPublicAlgoVersionNotExist"
//  INVALIDPARAMETERVALUE_DATASETNUMLIMITEXCEEDED = "InvalidParameterValue.DatasetNumLimitExceeded"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"
//  INVALIDPARAMETERVALUE_GETCFSMOUNTIPFAILED = "InvalidParameterValue.GetCFSMountIPFailed"
//  INVALIDPARAMETERVALUE_IMAGENOTFOUND = "InvalidParameterValue.ImageNotFound"
//  INVALIDPARAMETERVALUE_NOTALLOW = "InvalidParameterValue.NotAllow"
//  INVALIDPARAMETERVALUE_RDMACONFIGILLEGAL = "InvalidParameterValue.RDMAConfigIllegal"
//  INVALIDPARAMETERVALUE_RESOURCECONFIGILLEGAL = "InvalidParameterValue.ResourceConfigIllegal"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDDATACONFIG = "InvalidParameterValue.UnsupportedDataConfig"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_BILLINGEXCEPTION = "OperationDenied.BillingException"
//  OPERATIONDENIED_BILLINGSTATUSRESOURCEINSUFFICIENT = "OperationDenied.BillingStatusResourceInsufficient"
//  OPERATIONDENIED_IPILLEGAL = "OperationDenied.IpIllegal"
//  OPERATIONDENIED_MIYINGBALANCEINSUFFICIENT = "OperationDenied.MIYINGBalanceInsufficient"
//  OPERATIONDENIED_NETWORKCIDRILLEGAL = "OperationDenied.NetworkCidrIllegal"
//  OPERATIONDENIED_NOTALLOW = "OperationDenied.NotAllow"
//  OPERATIONDENIED_RESOURCEGROUPINSUFFICIENT = "OperationDenied.ResourceGroupInsufficient"
//  OPERATIONDENIED_SUBNETILLEGAL = "OperationDenied.SubnetIllegal"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateTrainingTaskWithContext(ctx context.Context, request *CreateTrainingTaskRequest) (response *CreateTrainingTaskResponse, err error) {
    if request == nil {
        request = NewCreateTrainingTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTrainingTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTrainingTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBatchTaskRequest() (request *DeleteBatchTaskRequest) {
    request = &DeleteBatchTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DeleteBatchTask")
    
    
    return
}

func NewDeleteBatchTaskResponse() (response *DeleteBatchTaskResponse) {
    response = &DeleteBatchTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteBatchTask
// 删除跑批任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_NOPERMISSION = "AuthFailure.NoPermission"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteBatchTask(request *DeleteBatchTaskRequest) (response *DeleteBatchTaskResponse, err error) {
    return c.DeleteBatchTaskWithContext(context.Background(), request)
}

// DeleteBatchTask
// 删除跑批任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_NOPERMISSION = "AuthFailure.NoPermission"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteBatchTaskWithContext(ctx context.Context, request *DeleteBatchTaskRequest) (response *DeleteBatchTaskResponse, err error) {
    if request == nil {
        request = NewDeleteBatchTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBatchTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBatchTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDatasetRequest() (request *DeleteDatasetRequest) {
    request = &DeleteDatasetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DeleteDataset")
    
    
    return
}

func NewDeleteDatasetResponse() (response *DeleteDatasetResponse) {
    response = &DeleteDatasetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDataset
// 删除数据集
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteDataset(request *DeleteDatasetRequest) (response *DeleteDatasetResponse, err error) {
    return c.DeleteDatasetWithContext(context.Background(), request)
}

// DeleteDataset
// 删除数据集
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteDatasetWithContext(ctx context.Context, request *DeleteDatasetRequest) (response *DeleteDatasetResponse, err error) {
    if request == nil {
        request = NewDeleteDatasetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDataset require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDatasetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteModelAccelerateTaskRequest() (request *DeleteModelAccelerateTaskRequest) {
    request = &DeleteModelAccelerateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DeleteModelAccelerateTask")
    
    
    return
}

func NewDeleteModelAccelerateTaskResponse() (response *DeleteModelAccelerateTaskResponse) {
    response = &DeleteModelAccelerateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteModelAccelerateTask
// 删除模型加速任务
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteModelAccelerateTask(request *DeleteModelAccelerateTaskRequest) (response *DeleteModelAccelerateTaskResponse, err error) {
    return c.DeleteModelAccelerateTaskWithContext(context.Background(), request)
}

// DeleteModelAccelerateTask
// 删除模型加速任务
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DeleteModelAccelerateTaskWithContext(ctx context.Context, request *DeleteModelAccelerateTaskRequest) (response *DeleteModelAccelerateTaskResponse, err error) {
    if request == nil {
        request = NewDeleteModelAccelerateTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteModelAccelerateTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteModelAccelerateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteModelServiceRequest() (request *DeleteModelServiceRequest) {
    request = &DeleteModelServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DeleteModelService")
    
    
    return
}

func NewDeleteModelServiceResponse() (response *DeleteModelServiceResponse) {
    response = &DeleteModelServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteModelService
// 根据服务id删除模型服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteModelService(request *DeleteModelServiceRequest) (response *DeleteModelServiceResponse, err error) {
    return c.DeleteModelServiceWithContext(context.Background(), request)
}

// DeleteModelService
// 根据服务id删除模型服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteModelServiceWithContext(ctx context.Context, request *DeleteModelServiceRequest) (response *DeleteModelServiceResponse, err error) {
    if request == nil {
        request = NewDeleteModelServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteModelService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteModelServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteModelServiceGroupRequest() (request *DeleteModelServiceGroupRequest) {
    request = &DeleteModelServiceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DeleteModelServiceGroup")
    
    
    return
}

func NewDeleteModelServiceGroupResponse() (response *DeleteModelServiceGroupResponse) {
    response = &DeleteModelServiceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteModelServiceGroup
// 根据服务组id删除服务组下所有模型服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteModelServiceGroup(request *DeleteModelServiceGroupRequest) (response *DeleteModelServiceGroupResponse, err error) {
    return c.DeleteModelServiceGroupWithContext(context.Background(), request)
}

// DeleteModelServiceGroup
// 根据服务组id删除服务组下所有模型服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteModelServiceGroupWithContext(ctx context.Context, request *DeleteModelServiceGroupRequest) (response *DeleteModelServiceGroupResponse, err error) {
    if request == nil {
        request = NewDeleteModelServiceGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteModelServiceGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteModelServiceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNotebookRequest() (request *DeleteNotebookRequest) {
    request = &DeleteNotebookRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DeleteNotebook")
    
    
    return
}

func NewDeleteNotebookResponse() (response *DeleteNotebookResponse) {
    response = &DeleteNotebookResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNotebook
// 删除Notebook
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETENOTEBOOKSTORAGEFAILED = "FailedOperation.DeleteNotebookStorageFailed"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_UNBINDINGTAGSFAILED = "FailedOperation.UnBindingTagsFailed"
//  FAILEDOPERATION_UNSUBMITNOTALLOWTOSTOP = "FailedOperation.UnSubmitNotAllowToStop"
//  FAILEDOPERATION_UNFREEZEBILLFAILED = "FailedOperation.UnfreezeBillFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteNotebook(request *DeleteNotebookRequest) (response *DeleteNotebookResponse, err error) {
    return c.DeleteNotebookWithContext(context.Background(), request)
}

// DeleteNotebook
// 删除Notebook
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETENOTEBOOKSTORAGEFAILED = "FailedOperation.DeleteNotebookStorageFailed"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_UNBINDINGTAGSFAILED = "FailedOperation.UnBindingTagsFailed"
//  FAILEDOPERATION_UNSUBMITNOTALLOWTOSTOP = "FailedOperation.UnSubmitNotAllowToStop"
//  FAILEDOPERATION_UNFREEZEBILLFAILED = "FailedOperation.UnfreezeBillFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteNotebookWithContext(ctx context.Context, request *DeleteNotebookRequest) (response *DeleteNotebookResponse, err error) {
    if request == nil {
        request = NewDeleteNotebookRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNotebook require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNotebookResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteNotebookImageRecordRequest() (request *DeleteNotebookImageRecordRequest) {
    request = &DeleteNotebookImageRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DeleteNotebookImageRecord")
    
    
    return
}

func NewDeleteNotebookImageRecordResponse() (response *DeleteNotebookImageRecordResponse) {
    response = &DeleteNotebookImageRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteNotebookImageRecord
// 删除notebook镜像保存记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETENOTEBOOKSTORAGEFAILED = "FailedOperation.DeleteNotebookStorageFailed"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_UNBINDINGTAGSFAILED = "FailedOperation.UnBindingTagsFailed"
//  FAILEDOPERATION_UNSUBMITNOTALLOWTOSTOP = "FailedOperation.UnSubmitNotAllowToStop"
//  FAILEDOPERATION_UNFREEZEBILLFAILED = "FailedOperation.UnfreezeBillFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteNotebookImageRecord(request *DeleteNotebookImageRecordRequest) (response *DeleteNotebookImageRecordResponse, err error) {
    return c.DeleteNotebookImageRecordWithContext(context.Background(), request)
}

// DeleteNotebookImageRecord
// 删除notebook镜像保存记录
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DELETENOTEBOOKSTORAGEFAILED = "FailedOperation.DeleteNotebookStorageFailed"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_UNBINDINGTAGSFAILED = "FailedOperation.UnBindingTagsFailed"
//  FAILEDOPERATION_UNSUBMITNOTALLOWTOSTOP = "FailedOperation.UnSubmitNotAllowToStop"
//  FAILEDOPERATION_UNFREEZEBILLFAILED = "FailedOperation.UnfreezeBillFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteNotebookImageRecordWithContext(ctx context.Context, request *DeleteNotebookImageRecordRequest) (response *DeleteNotebookImageRecordResponse, err error) {
    if request == nil {
        request = NewDeleteNotebookImageRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNotebookImageRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNotebookImageRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTrainingModelRequest() (request *DeleteTrainingModelRequest) {
    request = &DeleteTrainingModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DeleteTrainingModel")
    
    
    return
}

func NewDeleteTrainingModelResponse() (response *DeleteTrainingModelResponse) {
    response = &DeleteTrainingModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTrainingModel
// 删除模型
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOMODEL = "ResourceNotFound.NoModel"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteTrainingModel(request *DeleteTrainingModelRequest) (response *DeleteTrainingModelResponse, err error) {
    return c.DeleteTrainingModelWithContext(context.Background(), request)
}

// DeleteTrainingModel
// 删除模型
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOMODEL = "ResourceNotFound.NoModel"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteTrainingModelWithContext(ctx context.Context, request *DeleteTrainingModelRequest) (response *DeleteTrainingModelResponse, err error) {
    if request == nil {
        request = NewDeleteTrainingModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTrainingModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTrainingModelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTrainingModelVersionRequest() (request *DeleteTrainingModelVersionRequest) {
    request = &DeleteTrainingModelVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DeleteTrainingModelVersion")
    
    
    return
}

func NewDeleteTrainingModelVersionResponse() (response *DeleteTrainingModelVersionResponse) {
    response = &DeleteTrainingModelVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTrainingModelVersion
// 删除模型版本
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteTrainingModelVersion(request *DeleteTrainingModelVersionRequest) (response *DeleteTrainingModelVersionResponse, err error) {
    return c.DeleteTrainingModelVersionWithContext(context.Background(), request)
}

// DeleteTrainingModelVersion
// 删除模型版本
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteTrainingModelVersionWithContext(ctx context.Context, request *DeleteTrainingModelVersionRequest) (response *DeleteTrainingModelVersionResponse, err error) {
    if request == nil {
        request = NewDeleteTrainingModelVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTrainingModelVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTrainingModelVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTrainingTaskRequest() (request *DeleteTrainingTaskRequest) {
    request = &DeleteTrainingTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DeleteTrainingTask")
    
    
    return
}

func NewDeleteTrainingTaskResponse() (response *DeleteTrainingTaskResponse) {
    response = &DeleteTrainingTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTrainingTask
// 删除训练任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NOTALLOW = "InternalError.NotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteTrainingTask(request *DeleteTrainingTaskRequest) (response *DeleteTrainingTaskResponse, err error) {
    return c.DeleteTrainingTaskWithContext(context.Background(), request)
}

// DeleteTrainingTask
// 删除训练任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NOTALLOW = "InternalError.NotAllow"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeleteTrainingTaskWithContext(ctx context.Context, request *DeleteTrainingTaskRequest) (response *DeleteTrainingTaskResponse, err error) {
    if request == nil {
        request = NewDeleteTrainingTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTrainingTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTrainingTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAPIConfigsRequest() (request *DescribeAPIConfigsRequest) {
    request = &DescribeAPIConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeAPIConfigs")
    
    
    return
}

func NewDescribeAPIConfigsResponse() (response *DescribeAPIConfigsResponse) {
    response = &DescribeAPIConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAPIConfigs
// 列举API
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAPIConfigs(request *DescribeAPIConfigsRequest) (response *DescribeAPIConfigsResponse, err error) {
    return c.DescribeAPIConfigsWithContext(context.Background(), request)
}

// DescribeAPIConfigs
// 列举API
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAPIConfigsWithContext(ctx context.Context, request *DescribeAPIConfigsRequest) (response *DescribeAPIConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeAPIConfigsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAPIConfigs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAPIConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBatchTaskRequest() (request *DescribeBatchTaskRequest) {
    request = &DescribeBatchTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeBatchTask")
    
    
    return
}

func NewDescribeBatchTaskResponse() (response *DescribeBatchTaskResponse) {
    response = &DescribeBatchTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBatchTask
// 查询跑批任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBatchTask(request *DescribeBatchTaskRequest) (response *DescribeBatchTaskResponse, err error) {
    return c.DescribeBatchTaskWithContext(context.Background(), request)
}

// DescribeBatchTask
// 查询跑批任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBatchTaskWithContext(ctx context.Context, request *DescribeBatchTaskRequest) (response *DescribeBatchTaskResponse, err error) {
    if request == nil {
        request = NewDescribeBatchTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBatchTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBatchTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBatchTaskInstancesRequest() (request *DescribeBatchTaskInstancesRequest) {
    request = &DescribeBatchTaskInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeBatchTaskInstances")
    
    
    return
}

func NewDescribeBatchTaskInstancesResponse() (response *DescribeBatchTaskInstancesResponse) {
    response = &DescribeBatchTaskInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBatchTaskInstances
// 查询跑批实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBatchTaskInstances(request *DescribeBatchTaskInstancesRequest) (response *DescribeBatchTaskInstancesResponse, err error) {
    return c.DescribeBatchTaskInstancesWithContext(context.Background(), request)
}

// DescribeBatchTaskInstances
// 查询跑批实例列表
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBatchTaskInstancesWithContext(ctx context.Context, request *DescribeBatchTaskInstancesRequest) (response *DescribeBatchTaskInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeBatchTaskInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBatchTaskInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBatchTaskInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBatchTasksRequest() (request *DescribeBatchTasksRequest) {
    request = &DescribeBatchTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeBatchTasks")
    
    
    return
}

func NewDescribeBatchTasksResponse() (response *DescribeBatchTasksResponse) {
    response = &DescribeBatchTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBatchTasks
// 批量预测任务列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBatchTasks(request *DescribeBatchTasksRequest) (response *DescribeBatchTasksResponse, err error) {
    return c.DescribeBatchTasksWithContext(context.Background(), request)
}

// DescribeBatchTasks
// 批量预测任务列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBatchTasksWithContext(ctx context.Context, request *DescribeBatchTasksRequest) (response *DescribeBatchTasksResponse, err error) {
    if request == nil {
        request = NewDescribeBatchTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBatchTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBatchTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillingResourceGroupsRequest() (request *DescribeBillingResourceGroupsRequest) {
    request = &DescribeBillingResourceGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeBillingResourceGroups")
    
    
    return
}

func NewDescribeBillingResourceGroupsResponse() (response *DescribeBillingResourceGroupsResponse) {
    response = &DescribeBillingResourceGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillingResourceGroups
// 查询资源组详情
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLCLUSTERFAIL = "FailedOperation.CallClusterFail"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VALIDATEERROR = "InvalidParameter.ValidateError"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBillingResourceGroups(request *DescribeBillingResourceGroupsRequest) (response *DescribeBillingResourceGroupsResponse, err error) {
    return c.DescribeBillingResourceGroupsWithContext(context.Background(), request)
}

// DescribeBillingResourceGroups
// 查询资源组详情
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLCLUSTERFAIL = "FailedOperation.CallClusterFail"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_VALIDATEERROR = "InvalidParameter.ValidateError"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBillingResourceGroupsWithContext(ctx context.Context, request *DescribeBillingResourceGroupsRequest) (response *DescribeBillingResourceGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeBillingResourceGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillingResourceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillingResourceGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillingSpecsRequest() (request *DescribeBillingSpecsRequest) {
    request = &DescribeBillingSpecsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeBillingSpecs")
    
    
    return
}

func NewDescribeBillingSpecsResponse() (response *DescribeBillingSpecsResponse) {
    response = &DescribeBillingSpecsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillingSpecs
// 本接口(DescribeBillingSpecs)用于查询计费项列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QUERYSPECSFAILED = "FailedOperation.QuerySpecsFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBillingSpecs(request *DescribeBillingSpecsRequest) (response *DescribeBillingSpecsResponse, err error) {
    return c.DescribeBillingSpecsWithContext(context.Background(), request)
}

// DescribeBillingSpecs
// 本接口(DescribeBillingSpecs)用于查询计费项列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QUERYSPECSFAILED = "FailedOperation.QuerySpecsFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBillingSpecsWithContext(ctx context.Context, request *DescribeBillingSpecsRequest) (response *DescribeBillingSpecsResponse, err error) {
    if request == nil {
        request = NewDescribeBillingSpecsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillingSpecs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillingSpecsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillingSpecsPriceRequest() (request *DescribeBillingSpecsPriceRequest) {
    request = &DescribeBillingSpecsPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeBillingSpecsPrice")
    
    
    return
}

func NewDescribeBillingSpecsPriceResponse() (response *DescribeBillingSpecsPriceResponse) {
    response = &DescribeBillingSpecsPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillingSpecsPrice
// 本接口(DescribeBillingSpecsPrice)用于查询按量计费计费项价格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillingSpecsPrice(request *DescribeBillingSpecsPriceRequest) (response *DescribeBillingSpecsPriceResponse, err error) {
    return c.DescribeBillingSpecsPriceWithContext(context.Background(), request)
}

// DescribeBillingSpecsPrice
// 本接口(DescribeBillingSpecsPrice)用于查询按量计费计费项价格。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillingSpecsPriceWithContext(ctx context.Context, request *DescribeBillingSpecsPriceRequest) (response *DescribeBillingSpecsPriceResponse, err error) {
    if request == nil {
        request = NewDescribeBillingSpecsPriceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillingSpecsPrice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillingSpecsPriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatasetDetailStructuredRequest() (request *DescribeDatasetDetailStructuredRequest) {
    request = &DescribeDatasetDetailStructuredRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeDatasetDetailStructured")
    
    
    return
}

func NewDescribeDatasetDetailStructuredResponse() (response *DescribeDatasetDetailStructuredResponse) {
    response = &DescribeDatasetDetailStructuredResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatasetDetailStructured
// 查询结构化数据集详情
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDatasetDetailStructured(request *DescribeDatasetDetailStructuredRequest) (response *DescribeDatasetDetailStructuredResponse, err error) {
    return c.DescribeDatasetDetailStructuredWithContext(context.Background(), request)
}

// DescribeDatasetDetailStructured
// 查询结构化数据集详情
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDatasetDetailStructuredWithContext(ctx context.Context, request *DescribeDatasetDetailStructuredRequest) (response *DescribeDatasetDetailStructuredResponse, err error) {
    if request == nil {
        request = NewDescribeDatasetDetailStructuredRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatasetDetailStructured require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatasetDetailStructuredResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatasetDetailUnstructuredRequest() (request *DescribeDatasetDetailUnstructuredRequest) {
    request = &DescribeDatasetDetailUnstructuredRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeDatasetDetailUnstructured")
    
    
    return
}

func NewDescribeDatasetDetailUnstructuredResponse() (response *DescribeDatasetDetailUnstructuredResponse) {
    response = &DescribeDatasetDetailUnstructuredResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatasetDetailUnstructured
// 查询非结构化数据集详情
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDatasetDetailUnstructured(request *DescribeDatasetDetailUnstructuredRequest) (response *DescribeDatasetDetailUnstructuredResponse, err error) {
    return c.DescribeDatasetDetailUnstructuredWithContext(context.Background(), request)
}

// DescribeDatasetDetailUnstructured
// 查询非结构化数据集详情
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDatasetDetailUnstructuredWithContext(ctx context.Context, request *DescribeDatasetDetailUnstructuredRequest) (response *DescribeDatasetDetailUnstructuredResponse, err error) {
    if request == nil {
        request = NewDescribeDatasetDetailUnstructuredRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatasetDetailUnstructured require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatasetDetailUnstructuredResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatasetsRequest() (request *DescribeDatasetsRequest) {
    request = &DescribeDatasetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeDatasets")
    
    
    return
}

func NewDescribeDatasetsResponse() (response *DescribeDatasetsResponse) {
    response = &DescribeDatasetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatasets
// 查询数据集列表
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDatasets(request *DescribeDatasetsRequest) (response *DescribeDatasetsResponse, err error) {
    return c.DescribeDatasetsWithContext(context.Background(), request)
}

// DescribeDatasets
// 查询数据集列表
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  FAILEDOPERATION_DCCREATEASYNCTASKERROR = "FailedOperation.DCCreateAsyncTaskError"
//  FAILEDOPERATION_DCCREATEUSERCOSCLIENTERR = "FailedOperation.DCCreateUserCosClientErr"
//  FAILEDOPERATION_DCDATAANNOTATIONRPCERR = "FailedOperation.DCDataAnnotationRpcErr"
//  FAILEDOPERATION_DCDATAREPORPCERR = "FailedOperation.DCDatarepoRpcErr"
//  FAILEDOPERATION_DCDATASETEXCEEDLIMIT = "FailedOperation.DCDatasetExceedLimit"
//  FAILEDOPERATION_DCDATASETSTATUSNOTREADY = "FailedOperation.DCDatasetStatusNotReady"
//  FAILEDOPERATION_DCGETUSERTEMPORARYSECRETERR = "FailedOperation.DCGetUserTemporarySecretErr"
//  FAILEDOPERATION_DCMARSHALDATAERR = "FailedOperation.DCMarshalDataErr"
//  FAILEDOPERATION_DCQUERYDATASETCONTENTERR = "FailedOperation.DCQueryDatasetContentErr"
//  FAILEDOPERATION_DCUNMARSHALDATAERR = "FailedOperation.DCUnmarshalDataErr"
//  FAILEDOPERATION_DCUNSUPPORTEDOPERATION = "FailedOperation.DCUnsupportedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_DCANNOTATIONTYPE = "InvalidParameterValue.DCAnnotationType"
//  INVALIDPARAMETERVALUE_DCCOSPATHINFO = "InvalidParameterValue.DCCosPathInfo"
//  INVALIDPARAMETERVALUE_DCDATASETANNOTATIONNOTMATCH = "InvalidParameterValue.DCDatasetAnnotationNotMatch"
//  INVALIDPARAMETERVALUE_DCDATASETIDNOTEXIST = "InvalidParameterValue.DCDatasetIdNotExist"
//  INVALIDPARAMETERVALUE_DCDATASETNAMEEXIST = "InvalidParameterValue.DCDatasetNameExist"
//  INVALIDPARAMETERVALUE_DCDATASETTYPE = "InvalidParameterValue.DCDatasetType"
//  INVALIDPARAMETERVALUE_DCFILTERVALUES = "InvalidParameterValue.DCFilterValues"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeDatasetsWithContext(ctx context.Context, request *DescribeDatasetsRequest) (response *DescribeDatasetsResponse, err error) {
    if request == nil {
        request = NewDescribeDatasetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatasets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatasetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInferTemplatesRequest() (request *DescribeInferTemplatesRequest) {
    request = &DescribeInferTemplatesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeInferTemplates")
    
    
    return
}

func NewDescribeInferTemplatesResponse() (response *DescribeInferTemplatesResponse) {
    response = &DescribeInferTemplatesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInferTemplates
// 查询推理镜像模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInferTemplates(request *DescribeInferTemplatesRequest) (response *DescribeInferTemplatesResponse, err error) {
    return c.DescribeInferTemplatesWithContext(context.Background(), request)
}

// DescribeInferTemplates
// 查询推理镜像模板
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeInferTemplatesWithContext(ctx context.Context, request *DescribeInferTemplatesRequest) (response *DescribeInferTemplatesResponse, err error) {
    if request == nil {
        request = NewDescribeInferTemplatesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInferTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInferTemplatesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLatestTrainingMetricsRequest() (request *DescribeLatestTrainingMetricsRequest) {
    request = &DescribeLatestTrainingMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeLatestTrainingMetrics")
    
    
    return
}

func NewDescribeLatestTrainingMetricsResponse() (response *DescribeLatestTrainingMetricsResponse) {
    response = &DescribeLatestTrainingMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLatestTrainingMetrics
// 查询最近上报的训练自定义指标
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_NOPERMISSION = "AuthFailure.NoPermission"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_BINDINGTAGSFAILED = "FailedOperation.BindingTagsFailed"
//  FAILEDOPERATION_CAMFAILURE = "FailedOperation.CAMFailure"
//  FAILEDOPERATION_CALLCLUSTERFAIL = "FailedOperation.CallClusterFail"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_CREATEJOBINSTANCEFAILED = "FailedOperation.CreateJobInstanceFailed"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FREEZEBILLFAILED = "InternalError.FreezeBillFailed"
//  INTERNALERROR_INSUFFICIENTBALANCE = "InternalError.InsufficientBalance"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLatestTrainingMetrics(request *DescribeLatestTrainingMetricsRequest) (response *DescribeLatestTrainingMetricsResponse, err error) {
    return c.DescribeLatestTrainingMetricsWithContext(context.Background(), request)
}

// DescribeLatestTrainingMetrics
// 查询最近上报的训练自定义指标
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_NOPERMISSION = "AuthFailure.NoPermission"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_BINDINGTAGSFAILED = "FailedOperation.BindingTagsFailed"
//  FAILEDOPERATION_CAMFAILURE = "FailedOperation.CAMFailure"
//  FAILEDOPERATION_CALLCLUSTERFAIL = "FailedOperation.CallClusterFail"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_CREATEJOBINSTANCEFAILED = "FailedOperation.CreateJobInstanceFailed"
//  FAILEDOPERATION_DCCOSCLIENTERR = "FailedOperation.DCCosClientErr"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FREEZEBILLFAILED = "InternalError.FreezeBillFailed"
//  INTERNALERROR_INSUFFICIENTBALANCE = "InternalError.InsufficientBalance"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLatestTrainingMetricsWithContext(ctx context.Context, request *DescribeLatestTrainingMetricsRequest) (response *DescribeLatestTrainingMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeLatestTrainingMetricsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLatestTrainingMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLatestTrainingMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogsRequest() (request *DescribeLogsRequest) {
    request = &DescribeLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeLogs")
    
    
    return
}

func NewDescribeLogsResponse() (response *DescribeLogsResponse) {
    response = &DescribeLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogs
// 获取训练、推理、Notebook服务的日志 API
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeLogs(request *DescribeLogsRequest) (response *DescribeLogsResponse, err error) {
    return c.DescribeLogsWithContext(context.Background(), request)
}

// DescribeLogs
// 获取训练、推理、Notebook服务的日志 API
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeLogsWithContext(ctx context.Context, request *DescribeLogsRequest) (response *DescribeLogsResponse, err error) {
    if request == nil {
        request = NewDescribeLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelAccEngineVersionsRequest() (request *DescribeModelAccEngineVersionsRequest) {
    request = &DescribeModelAccEngineVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeModelAccEngineVersions")
    
    
    return
}

func NewDescribeModelAccEngineVersionsResponse() (response *DescribeModelAccEngineVersionsResponse) {
    response = &DescribeModelAccEngineVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeModelAccEngineVersions
// 查询模型加速引擎版本列表
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TGWINVALIDINTERFACE = "InvalidParameter.TgwInvalidInterface"
//  INVALIDPARAMETER_TGWINVALIDREQUESTBODY = "InvalidParameter.TgwInvalidRequestBody"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeModelAccEngineVersions(request *DescribeModelAccEngineVersionsRequest) (response *DescribeModelAccEngineVersionsResponse, err error) {
    return c.DescribeModelAccEngineVersionsWithContext(context.Background(), request)
}

// DescribeModelAccEngineVersions
// 查询模型加速引擎版本列表
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TGWINVALIDINTERFACE = "InvalidParameter.TgwInvalidInterface"
//  INVALIDPARAMETER_TGWINVALIDREQUESTBODY = "InvalidParameter.TgwInvalidRequestBody"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeModelAccEngineVersionsWithContext(ctx context.Context, request *DescribeModelAccEngineVersionsRequest) (response *DescribeModelAccEngineVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeModelAccEngineVersionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelAccEngineVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelAccEngineVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelAccelerateTaskRequest() (request *DescribeModelAccelerateTaskRequest) {
    request = &DescribeModelAccelerateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeModelAccelerateTask")
    
    
    return
}

func NewDescribeModelAccelerateTaskResponse() (response *DescribeModelAccelerateTaskResponse) {
    response = &DescribeModelAccelerateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeModelAccelerateTask
// 查询模型优化任务详情
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeModelAccelerateTask(request *DescribeModelAccelerateTaskRequest) (response *DescribeModelAccelerateTaskResponse, err error) {
    return c.DescribeModelAccelerateTaskWithContext(context.Background(), request)
}

// DescribeModelAccelerateTask
// 查询模型优化任务详情
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeModelAccelerateTaskWithContext(ctx context.Context, request *DescribeModelAccelerateTaskRequest) (response *DescribeModelAccelerateTaskResponse, err error) {
    if request == nil {
        request = NewDescribeModelAccelerateTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelAccelerateTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelAccelerateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelAccelerateTasksRequest() (request *DescribeModelAccelerateTasksRequest) {
    request = &DescribeModelAccelerateTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeModelAccelerateTasks")
    
    
    return
}

func NewDescribeModelAccelerateTasksResponse() (response *DescribeModelAccelerateTasksResponse) {
    response = &DescribeModelAccelerateTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeModelAccelerateTasks
// 查询模型加速任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeModelAccelerateTasks(request *DescribeModelAccelerateTasksRequest) (response *DescribeModelAccelerateTasksResponse, err error) {
    return c.DescribeModelAccelerateTasksWithContext(context.Background(), request)
}

// DescribeModelAccelerateTasks
// 查询模型加速任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeModelAccelerateTasksWithContext(ctx context.Context, request *DescribeModelAccelerateTasksRequest) (response *DescribeModelAccelerateTasksResponse, err error) {
    if request == nil {
        request = NewDescribeModelAccelerateTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelAccelerateTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelAccelerateTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelServiceRequest() (request *DescribeModelServiceRequest) {
    request = &DescribeModelServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeModelService")
    
    
    return
}

func NewDescribeModelServiceResponse() (response *DescribeModelServiceResponse) {
    response = &DescribeModelServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeModelService
// 查询单个服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_INVALIDUSERTYPE = "FailedOperation.InvalidUserType"
//  FAILEDOPERATION_KMSKEYNOTOPEN = "FailedOperation.KmsKeyNotOpen"
//  FAILEDOPERATION_NOFREEBUCKET = "FailedOperation.NoFreeBucket"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_PROCESSING = "FailedOperation.Processing"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_REPOBINDBYINSTANCE = "FailedOperation.RepoBindByInstance"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_TIMEDOUT = "FailedOperation.Timedout"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"
//  INVALIDPARAMETERVALUE_CLSCONFIGREQUIRED = "InvalidParameterValue.ClsConfigRequired"
//  INVALIDPARAMETERVALUE_CODEREPONOTFOUND = "InvalidParameterValue.CodeRepoNotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeModelService(request *DescribeModelServiceRequest) (response *DescribeModelServiceResponse, err error) {
    return c.DescribeModelServiceWithContext(context.Background(), request)
}

// DescribeModelService
// 查询单个服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_INVALIDUSERTYPE = "FailedOperation.InvalidUserType"
//  FAILEDOPERATION_KMSKEYNOTOPEN = "FailedOperation.KmsKeyNotOpen"
//  FAILEDOPERATION_NOFREEBUCKET = "FailedOperation.NoFreeBucket"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_PROCESSING = "FailedOperation.Processing"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_REPOBINDBYINSTANCE = "FailedOperation.RepoBindByInstance"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_TIMEDOUT = "FailedOperation.Timedout"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"
//  INVALIDPARAMETERVALUE_CLSCONFIGREQUIRED = "InvalidParameterValue.ClsConfigRequired"
//  INVALIDPARAMETERVALUE_CODEREPONOTFOUND = "InvalidParameterValue.CodeRepoNotFound"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeModelServiceWithContext(ctx context.Context, request *DescribeModelServiceRequest) (response *DescribeModelServiceResponse, err error) {
    if request == nil {
        request = NewDescribeModelServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelServiceCallInfoRequest() (request *DescribeModelServiceCallInfoRequest) {
    request = &DescribeModelServiceCallInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeModelServiceCallInfo")
    
    
    return
}

func NewDescribeModelServiceCallInfoResponse() (response *DescribeModelServiceCallInfoResponse) {
    response = &DescribeModelServiceCallInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeModelServiceCallInfo
// 展示服务的调用信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_INVALIDUSERTYPE = "FailedOperation.InvalidUserType"
//  FAILEDOPERATION_KMSKEYNOTOPEN = "FailedOperation.KmsKeyNotOpen"
//  FAILEDOPERATION_NOFREEBUCKET = "FailedOperation.NoFreeBucket"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_PROCESSING = "FailedOperation.Processing"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_REPOBINDBYINSTANCE = "FailedOperation.RepoBindByInstance"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_TIMEDOUT = "FailedOperation.Timedout"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeModelServiceCallInfo(request *DescribeModelServiceCallInfoRequest) (response *DescribeModelServiceCallInfoResponse, err error) {
    return c.DescribeModelServiceCallInfoWithContext(context.Background(), request)
}

// DescribeModelServiceCallInfo
// 展示服务的调用信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_INVALIDUSERTYPE = "FailedOperation.InvalidUserType"
//  FAILEDOPERATION_KMSKEYNOTOPEN = "FailedOperation.KmsKeyNotOpen"
//  FAILEDOPERATION_NOFREEBUCKET = "FailedOperation.NoFreeBucket"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_PROCESSING = "FailedOperation.Processing"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_REPOBINDBYINSTANCE = "FailedOperation.RepoBindByInstance"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_TIMEDOUT = "FailedOperation.Timedout"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeModelServiceCallInfoWithContext(ctx context.Context, request *DescribeModelServiceCallInfoRequest) (response *DescribeModelServiceCallInfoResponse, err error) {
    if request == nil {
        request = NewDescribeModelServiceCallInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelServiceCallInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelServiceCallInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelServiceGroupRequest() (request *DescribeModelServiceGroupRequest) {
    request = &DescribeModelServiceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeModelServiceGroup")
    
    
    return
}

func NewDescribeModelServiceGroupResponse() (response *DescribeModelServiceGroupResponse) {
    response = &DescribeModelServiceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeModelServiceGroup
// 查询单个服务组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeModelServiceGroup(request *DescribeModelServiceGroupRequest) (response *DescribeModelServiceGroupResponse, err error) {
    return c.DescribeModelServiceGroupWithContext(context.Background(), request)
}

// DescribeModelServiceGroup
// 查询单个服务组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeModelServiceGroupWithContext(ctx context.Context, request *DescribeModelServiceGroupRequest) (response *DescribeModelServiceGroupResponse, err error) {
    if request == nil {
        request = NewDescribeModelServiceGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelServiceGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelServiceGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelServiceGroupsRequest() (request *DescribeModelServiceGroupsRequest) {
    request = &DescribeModelServiceGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeModelServiceGroups")
    
    
    return
}

func NewDescribeModelServiceGroupsResponse() (response *DescribeModelServiceGroupsResponse) {
    response = &DescribeModelServiceGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeModelServiceGroups
// 列举在线推理服务组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeModelServiceGroups(request *DescribeModelServiceGroupsRequest) (response *DescribeModelServiceGroupsResponse, err error) {
    return c.DescribeModelServiceGroupsWithContext(context.Background(), request)
}

// DescribeModelServiceGroups
// 列举在线推理服务组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeModelServiceGroupsWithContext(ctx context.Context, request *DescribeModelServiceGroupsRequest) (response *DescribeModelServiceGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeModelServiceGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelServiceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelServiceGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelServiceHistoryRequest() (request *DescribeModelServiceHistoryRequest) {
    request = &DescribeModelServiceHistoryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeModelServiceHistory")
    
    
    return
}

func NewDescribeModelServiceHistoryResponse() (response *DescribeModelServiceHistoryResponse) {
    response = &DescribeModelServiceHistoryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeModelServiceHistory
// 展示服务的历史版本
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_INVALIDUSERTYPE = "FailedOperation.InvalidUserType"
//  FAILEDOPERATION_KMSKEYNOTOPEN = "FailedOperation.KmsKeyNotOpen"
//  FAILEDOPERATION_NOFREEBUCKET = "FailedOperation.NoFreeBucket"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_PROCESSING = "FailedOperation.Processing"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_REPOBINDBYINSTANCE = "FailedOperation.RepoBindByInstance"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_TIMEDOUT = "FailedOperation.Timedout"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeModelServiceHistory(request *DescribeModelServiceHistoryRequest) (response *DescribeModelServiceHistoryResponse, err error) {
    return c.DescribeModelServiceHistoryWithContext(context.Background(), request)
}

// DescribeModelServiceHistory
// 展示服务的历史版本
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_INVALIDUSERTYPE = "FailedOperation.InvalidUserType"
//  FAILEDOPERATION_KMSKEYNOTOPEN = "FailedOperation.KmsKeyNotOpen"
//  FAILEDOPERATION_NOFREEBUCKET = "FailedOperation.NoFreeBucket"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_PROCESSING = "FailedOperation.Processing"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_REPOBINDBYINSTANCE = "FailedOperation.RepoBindByInstance"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_TIMEDOUT = "FailedOperation.Timedout"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeModelServiceHistoryWithContext(ctx context.Context, request *DescribeModelServiceHistoryRequest) (response *DescribeModelServiceHistoryResponse, err error) {
    if request == nil {
        request = NewDescribeModelServiceHistoryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelServiceHistory require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelServiceHistoryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelServiceHotUpdatedRequest() (request *DescribeModelServiceHotUpdatedRequest) {
    request = &DescribeModelServiceHotUpdatedRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeModelServiceHotUpdated")
    
    
    return
}

func NewDescribeModelServiceHotUpdatedResponse() (response *DescribeModelServiceHotUpdatedResponse) {
    response = &DescribeModelServiceHotUpdatedResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeModelServiceHotUpdated
// 用于查询模型服务能否开启热更新
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_INSUFFICIENTWHITELISTQUOTA = "FailedOperation.InsufficientWhitelistQuota"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeModelServiceHotUpdated(request *DescribeModelServiceHotUpdatedRequest) (response *DescribeModelServiceHotUpdatedResponse, err error) {
    return c.DescribeModelServiceHotUpdatedWithContext(context.Background(), request)
}

// DescribeModelServiceHotUpdated
// 用于查询模型服务能否开启热更新
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_INSUFFICIENTWHITELISTQUOTA = "FailedOperation.InsufficientWhitelistQuota"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeModelServiceHotUpdatedWithContext(ctx context.Context, request *DescribeModelServiceHotUpdatedRequest) (response *DescribeModelServiceHotUpdatedResponse, err error) {
    if request == nil {
        request = NewDescribeModelServiceHotUpdatedRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelServiceHotUpdated require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelServiceHotUpdatedResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelServicesRequest() (request *DescribeModelServicesRequest) {
    request = &DescribeModelServicesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeModelServices")
    
    
    return
}

func NewDescribeModelServicesResponse() (response *DescribeModelServicesResponse) {
    response = &DescribeModelServicesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeModelServices
// 查询多个服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_INSUFFICIENTWHITELISTQUOTA = "FailedOperation.InsufficientWhitelistQuota"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeModelServices(request *DescribeModelServicesRequest) (response *DescribeModelServicesResponse, err error) {
    return c.DescribeModelServicesWithContext(context.Background(), request)
}

// DescribeModelServices
// 查询多个服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_INSUFFICIENTWHITELISTQUOTA = "FailedOperation.InsufficientWhitelistQuota"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeModelServicesWithContext(ctx context.Context, request *DescribeModelServicesRequest) (response *DescribeModelServicesResponse, err error) {
    if request == nil {
        request = NewDescribeModelServicesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelServices require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelServicesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotebookRequest() (request *DescribeNotebookRequest) {
    request = &DescribeNotebookRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeNotebook")
    
    
    return
}

func NewDescribeNotebookResponse() (response *DescribeNotebookResponse) {
    response = &DescribeNotebookResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNotebook
// Notebook详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QUERYBINDINGTAGSFAILED = "FailedOperation.QueryBindingTagsFailed"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYRESOURCEGROUPNAMESFAILED = "FailedOperation.QueryResourceGroupNamesFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeNotebook(request *DescribeNotebookRequest) (response *DescribeNotebookResponse, err error) {
    return c.DescribeNotebookWithContext(context.Background(), request)
}

// DescribeNotebook
// Notebook详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_QUERYBINDINGTAGSFAILED = "FailedOperation.QueryBindingTagsFailed"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYRESOURCEGROUPNAMESFAILED = "FailedOperation.QueryResourceGroupNamesFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeNotebookWithContext(ctx context.Context, request *DescribeNotebookRequest) (response *DescribeNotebookResponse, err error) {
    if request == nil {
        request = NewDescribeNotebookRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNotebook require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNotebookResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotebookImageKernelsRequest() (request *DescribeNotebookImageKernelsRequest) {
    request = &DescribeNotebookImageKernelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeNotebookImageKernels")
    
    
    return
}

func NewDescribeNotebookImageKernelsResponse() (response *DescribeNotebookImageKernelsResponse) {
    response = &DescribeNotebookImageKernelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNotebookImageKernels
// 查询镜像kernel
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEIMAGEFAILED = "FailedOperation.CreateImageFailed"
//  OPERATIONDENIED_NOTSUPPORTSAVEIMAGE = "OperationDenied.NotSupportSaveImage"
func (c *Client) DescribeNotebookImageKernels(request *DescribeNotebookImageKernelsRequest) (response *DescribeNotebookImageKernelsResponse, err error) {
    return c.DescribeNotebookImageKernelsWithContext(context.Background(), request)
}

// DescribeNotebookImageKernels
// 查询镜像kernel
//
// 可能返回的错误码:
//  FAILEDOPERATION_CREATEIMAGEFAILED = "FailedOperation.CreateImageFailed"
//  OPERATIONDENIED_NOTSUPPORTSAVEIMAGE = "OperationDenied.NotSupportSaveImage"
func (c *Client) DescribeNotebookImageKernelsWithContext(ctx context.Context, request *DescribeNotebookImageKernelsRequest) (response *DescribeNotebookImageKernelsResponse, err error) {
    if request == nil {
        request = NewDescribeNotebookImageKernelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNotebookImageKernels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNotebookImageKernelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotebookImageRecordsRequest() (request *DescribeNotebookImageRecordsRequest) {
    request = &DescribeNotebookImageRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeNotebookImageRecords")
    
    
    return
}

func NewDescribeNotebookImageRecordsResponse() (response *DescribeNotebookImageRecordsResponse) {
    response = &DescribeNotebookImageRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNotebookImageRecords
// 查看notebook镜像保存记录
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNotebookImageRecords(request *DescribeNotebookImageRecordsRequest) (response *DescribeNotebookImageRecordsResponse, err error) {
    return c.DescribeNotebookImageRecordsWithContext(context.Background(), request)
}

// DescribeNotebookImageRecords
// 查看notebook镜像保存记录
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeNotebookImageRecordsWithContext(ctx context.Context, request *DescribeNotebookImageRecordsRequest) (response *DescribeNotebookImageRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeNotebookImageRecordsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNotebookImageRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNotebookImageRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeNotebooksRequest() (request *DescribeNotebooksRequest) {
    request = &DescribeNotebooksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeNotebooks")
    
    
    return
}

func NewDescribeNotebooksResponse() (response *DescribeNotebooksResponse) {
    response = &DescribeNotebooksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNotebooks
// Notebook列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CHECKBILLINGWHITELISTFAILED = "FailedOperation.CheckBillingWhiteListFailed"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYRESOURCEGROUPNAMESFAILED = "FailedOperation.QueryResourceGroupNamesFailed"
//  FAILEDOPERATION_QUERYRESOURCESPECFAILED = "FailedOperation.QueryResourceSpecFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYBINDINGTAGSFAILED = "InternalError.QueryBindingTagsFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FILTERKEYSTATUSCONFLICTWITHCHARGESTATUS = "InvalidParameter.FilterKeyStatusConflictWithChargeStatus"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeNotebooks(request *DescribeNotebooksRequest) (response *DescribeNotebooksResponse, err error) {
    return c.DescribeNotebooksWithContext(context.Background(), request)
}

// DescribeNotebooks
// Notebook列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_CHECKBILLINGWHITELISTFAILED = "FailedOperation.CheckBillingWhiteListFailed"
//  FAILEDOPERATION_QUERYPRICEFAILED = "FailedOperation.QueryPriceFailed"
//  FAILEDOPERATION_QUERYRESOURCEGROUPNAMESFAILED = "FailedOperation.QueryResourceGroupNamesFailed"
//  FAILEDOPERATION_QUERYRESOURCESPECFAILED = "FailedOperation.QueryResourceSpecFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYBINDINGTAGSFAILED = "InternalError.QueryBindingTagsFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FILTERKEYSTATUSCONFLICTWITHCHARGESTATUS = "InvalidParameter.FilterKeyStatusConflictWithChargeStatus"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeNotebooksWithContext(ctx context.Context, request *DescribeNotebooksRequest) (response *DescribeNotebooksResponse, err error) {
    if request == nil {
        request = NewDescribeNotebooksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNotebooks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNotebooksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrainingFrameworksRequest() (request *DescribeTrainingFrameworksRequest) {
    request = &DescribeTrainingFrameworksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeTrainingFrameworks")
    
    
    return
}

func NewDescribeTrainingFrameworksResponse() (response *DescribeTrainingFrameworksResponse) {
    response = &DescribeTrainingFrameworksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrainingFrameworks
// 训练框架列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_NOTALLOW = "OperationDenied.NotAllow"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingFrameworks(request *DescribeTrainingFrameworksRequest) (response *DescribeTrainingFrameworksResponse, err error) {
    return c.DescribeTrainingFrameworksWithContext(context.Background(), request)
}

// DescribeTrainingFrameworks
// 训练框架列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_NOTALLOW = "OperationDenied.NotAllow"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingFrameworksWithContext(ctx context.Context, request *DescribeTrainingFrameworksRequest) (response *DescribeTrainingFrameworksResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingFrameworksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrainingFrameworks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrainingFrameworksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrainingMetricsRequest() (request *DescribeTrainingMetricsRequest) {
    request = &DescribeTrainingMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeTrainingMetrics")
    
    
    return
}

func NewDescribeTrainingMetricsResponse() (response *DescribeTrainingMetricsResponse) {
    response = &DescribeTrainingMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrainingMetrics
// 查询训练自定义指标
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTrainingMetrics(request *DescribeTrainingMetricsRequest) (response *DescribeTrainingMetricsResponse, err error) {
    return c.DescribeTrainingMetricsWithContext(context.Background(), request)
}

// DescribeTrainingMetrics
// 查询训练自定义指标
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeTrainingMetricsWithContext(ctx context.Context, request *DescribeTrainingMetricsRequest) (response *DescribeTrainingMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingMetricsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrainingMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrainingMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrainingModelVersionRequest() (request *DescribeTrainingModelVersionRequest) {
    request = &DescribeTrainingModelVersionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeTrainingModelVersion")
    
    
    return
}

func NewDescribeTrainingModelVersionResponse() (response *DescribeTrainingModelVersionResponse) {
    response = &DescribeTrainingModelVersionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrainingModelVersion
// 查询模型版本
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingModelVersion(request *DescribeTrainingModelVersionRequest) (response *DescribeTrainingModelVersionResponse, err error) {
    return c.DescribeTrainingModelVersionWithContext(context.Background(), request)
}

// DescribeTrainingModelVersion
// 查询模型版本
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingModelVersionWithContext(ctx context.Context, request *DescribeTrainingModelVersionRequest) (response *DescribeTrainingModelVersionResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingModelVersionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrainingModelVersion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrainingModelVersionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrainingModelVersionsRequest() (request *DescribeTrainingModelVersionsRequest) {
    request = &DescribeTrainingModelVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeTrainingModelVersions")
    
    
    return
}

func NewDescribeTrainingModelVersionsResponse() (response *DescribeTrainingModelVersionsResponse) {
    response = &DescribeTrainingModelVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrainingModelVersions
// 模型版本列表
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOMODEL = "ResourceNotFound.NoModel"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingModelVersions(request *DescribeTrainingModelVersionsRequest) (response *DescribeTrainingModelVersionsResponse, err error) {
    return c.DescribeTrainingModelVersionsWithContext(context.Background(), request)
}

// DescribeTrainingModelVersions
// 模型版本列表
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_NOMODEL = "ResourceNotFound.NoModel"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingModelVersionsWithContext(ctx context.Context, request *DescribeTrainingModelVersionsRequest) (response *DescribeTrainingModelVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingModelVersionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrainingModelVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrainingModelVersionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrainingModelsRequest() (request *DescribeTrainingModelsRequest) {
    request = &DescribeTrainingModelsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeTrainingModels")
    
    
    return
}

func NewDescribeTrainingModelsResponse() (response *DescribeTrainingModelsResponse) {
    response = &DescribeTrainingModelsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrainingModels
// 模型列表
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYBINDINGTAGSFAILED = "FailedOperation.QueryBindingTagsFailed"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_QUERYMODELSBYTAGSFAILED = "FailedOperation.QueryModelsByTagsFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingModels(request *DescribeTrainingModelsRequest) (response *DescribeTrainingModelsResponse, err error) {
    return c.DescribeTrainingModelsWithContext(context.Background(), request)
}

// DescribeTrainingModels
// 模型列表
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYBINDINGTAGSFAILED = "FailedOperation.QueryBindingTagsFailed"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_QUERYMODELSBYTAGSFAILED = "FailedOperation.QueryModelsByTagsFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingModelsWithContext(ctx context.Context, request *DescribeTrainingModelsRequest) (response *DescribeTrainingModelsResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingModelsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrainingModels require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrainingModelsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrainingTaskRequest() (request *DescribeTrainingTaskRequest) {
    request = &DescribeTrainingTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeTrainingTask")
    
    
    return
}

func NewDescribeTrainingTaskResponse() (response *DescribeTrainingTaskResponse) {
    response = &DescribeTrainingTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrainingTask
// 训练任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYRESOURCEGROUPNAMESFAILED = "FailedOperation.QueryResourceGroupNamesFailed"
//  INTERNALERROR_NOPERMISSION = "InternalError.NoPermission"
//  INTERNALERROR_QUERYBINDINGTAGSFAILED = "InternalError.QueryBindingTagsFailed"
//  INTERNALERROR_QUERYRESOURCEGROUPNAMESFAILED = "InternalError.QueryResourceGroupNamesFailed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTrainingTask(request *DescribeTrainingTaskRequest) (response *DescribeTrainingTaskResponse, err error) {
    return c.DescribeTrainingTaskWithContext(context.Background(), request)
}

// DescribeTrainingTask
// 训练任务详情
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYRESOURCEGROUPNAMESFAILED = "FailedOperation.QueryResourceGroupNamesFailed"
//  INTERNALERROR_NOPERMISSION = "InternalError.NoPermission"
//  INTERNALERROR_QUERYBINDINGTAGSFAILED = "InternalError.QueryBindingTagsFailed"
//  INTERNALERROR_QUERYRESOURCEGROUPNAMESFAILED = "InternalError.QueryResourceGroupNamesFailed"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTrainingTaskWithContext(ctx context.Context, request *DescribeTrainingTaskRequest) (response *DescribeTrainingTaskResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrainingTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrainingTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrainingTaskPodsRequest() (request *DescribeTrainingTaskPodsRequest) {
    request = &DescribeTrainingTaskPodsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeTrainingTaskPods")
    
    
    return
}

func NewDescribeTrainingTaskPodsResponse() (response *DescribeTrainingTaskPodsResponse) {
    response = &DescribeTrainingTaskPodsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrainingTaskPods
// 训练任务pod列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYRESOURCEGROUPNAMESFAILED = "FailedOperation.QueryResourceGroupNamesFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NOPERMISSION = "InternalError.NoPermission"
//  INTERNALERROR_QUERYRESOURCEGROUPNAMESFAILED = "InternalError.QueryResourceGroupNamesFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingTaskPods(request *DescribeTrainingTaskPodsRequest) (response *DescribeTrainingTaskPodsResponse, err error) {
    return c.DescribeTrainingTaskPodsWithContext(context.Background(), request)
}

// DescribeTrainingTaskPods
// 训练任务pod列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYRESOURCEGROUPNAMESFAILED = "FailedOperation.QueryResourceGroupNamesFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NOPERMISSION = "InternalError.NoPermission"
//  INTERNALERROR_QUERYRESOURCEGROUPNAMESFAILED = "InternalError.QueryResourceGroupNamesFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTrainingTaskPodsWithContext(ctx context.Context, request *DescribeTrainingTaskPodsRequest) (response *DescribeTrainingTaskPodsResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingTaskPodsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrainingTaskPods require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrainingTaskPodsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrainingTasksRequest() (request *DescribeTrainingTasksRequest) {
    request = &DescribeTrainingTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeTrainingTasks")
    
    
    return
}

func NewDescribeTrainingTasksResponse() (response *DescribeTrainingTasksResponse) {
    response = &DescribeTrainingTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrainingTasks
// 训练任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYRESOURCESPECFAILED = "FailedOperation.QueryResourceSpecFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYRESOURCESPECFAILED = "InternalError.QueryResourceSpecFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_PAGELIMITEXCEEDED = "InvalidParameterValue.PageLimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTrainingTasks(request *DescribeTrainingTasksRequest) (response *DescribeTrainingTasksResponse, err error) {
    return c.DescribeTrainingTasksWithContext(context.Background(), request)
}

// DescribeTrainingTasks
// 训练任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYRESOURCESPECFAILED = "FailedOperation.QueryResourceSpecFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYRESOURCESPECFAILED = "InternalError.QueryResourceSpecFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_PAGELIMITEXCEEDED = "InvalidParameterValue.PageLimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTrainingTasksWithContext(ctx context.Context, request *DescribeTrainingTasksRequest) (response *DescribeTrainingTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrainingTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrainingTasksResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModelServiceRequest() (request *ModifyModelServiceRequest) {
    request = &ModifyModelServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "ModifyModelService")
    
    
    return
}

func NewModifyModelServiceResponse() (response *ModifyModelServiceResponse) {
    response = &ModifyModelServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyModelService
// 用于更新模型服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_INSUFFICIENTWHITELISTQUOTA = "FailedOperation.InsufficientWhitelistQuota"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MODELFILEINVALID = "InvalidParameter.ModelFileInvalid"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyModelService(request *ModifyModelServiceRequest) (response *ModifyModelServiceResponse, err error) {
    return c.ModifyModelServiceWithContext(context.Background(), request)
}

// ModifyModelService
// 用于更新模型服务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_INSUFFICIENTWHITELISTQUOTA = "FailedOperation.InsufficientWhitelistQuota"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MODELFILEINVALID = "InvalidParameter.ModelFileInvalid"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyModelServiceWithContext(ctx context.Context, request *ModifyModelServiceRequest) (response *ModifyModelServiceResponse, err error) {
    if request == nil {
        request = NewModifyModelServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyModelService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyModelServiceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModelServicePartialConfigRequest() (request *ModifyModelServicePartialConfigRequest) {
    request = &ModifyModelServicePartialConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "ModifyModelServicePartialConfig")
    
    
    return
}

func NewModifyModelServicePartialConfigResponse() (response *ModifyModelServicePartialConfigResponse) {
    response = &ModifyModelServicePartialConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyModelServicePartialConfig
// 增量更新在线推理服务的部分配置，不更新的配置项不需要传入
//
// 可能返回的错误码:
//  AUTHFAILURE_NOPERMISSION = "AuthFailure.NoPermission"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyModelServicePartialConfig(request *ModifyModelServicePartialConfigRequest) (response *ModifyModelServicePartialConfigResponse, err error) {
    return c.ModifyModelServicePartialConfigWithContext(context.Background(), request)
}

// ModifyModelServicePartialConfig
// 增量更新在线推理服务的部分配置，不更新的配置项不需要传入
//
// 可能返回的错误码:
//  AUTHFAILURE_NOPERMISSION = "AuthFailure.NoPermission"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyModelServicePartialConfigWithContext(ctx context.Context, request *ModifyModelServicePartialConfigRequest) (response *ModifyModelServicePartialConfigResponse, err error) {
    if request == nil {
        request = NewModifyModelServicePartialConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyModelServicePartialConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyModelServicePartialConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNotebookRequest() (request *ModifyNotebookRequest) {
    request = &ModifyNotebookRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "ModifyNotebook")
    
    
    return
}

func NewModifyNotebookResponse() (response *ModifyNotebookResponse) {
    response = &ModifyNotebookResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNotebook
// 修改Notebook
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDINGTAGSFAILED = "FailedOperation.BindingTagsFailed"
//  FAILEDOPERATION_MODIFYRESOURCEBILLINGTAGSFAILED = "FailedOperation.ModifyResourceBillingTagsFailed"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_UNBINDINGTAGSFAILED = "FailedOperation.UnBindingTagsFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BINDINGTAGSFAILED = "InternalError.BindingTagsFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CODEREPONOTFOUND = "InvalidParameterValue.CodeRepoNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_LIFECYCLENOTFOUND = "InvalidParameterValue.LifecycleNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
//  INVALIDPARAMETERVALUE_VOLUMESHRINKNOTALLOW = "InvalidParameterValue.VolumeShrinkNotAllow"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyNotebook(request *ModifyNotebookRequest) (response *ModifyNotebookResponse, err error) {
    return c.ModifyNotebookWithContext(context.Background(), request)
}

// ModifyNotebook
// 修改Notebook
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_BINDINGTAGSFAILED = "FailedOperation.BindingTagsFailed"
//  FAILEDOPERATION_MODIFYRESOURCEBILLINGTAGSFAILED = "FailedOperation.ModifyResourceBillingTagsFailed"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_UNBINDINGTAGSFAILED = "FailedOperation.UnBindingTagsFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_BINDINGTAGSFAILED = "InternalError.BindingTagsFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_CODEREPONOTFOUND = "InvalidParameterValue.CodeRepoNotFound"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_LIFECYCLENOTFOUND = "InvalidParameterValue.LifecycleNotFound"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
//  INVALIDPARAMETERVALUE_VOLUMESHRINKNOTALLOW = "InvalidParameterValue.VolumeShrinkNotAllow"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyNotebookWithContext(ctx context.Context, request *ModifyNotebookRequest) (response *ModifyNotebookResponse, err error) {
    if request == nil {
        request = NewModifyNotebookRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNotebook require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNotebookResponse()
    err = c.Send(request, response)
    return
}

func NewModifyNotebookTagsRequest() (request *ModifyNotebookTagsRequest) {
    request = &ModifyNotebookTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "ModifyNotebookTags")
    
    
    return
}

func NewModifyNotebookTagsResponse() (response *ModifyNotebookTagsResponse) {
    response = &ModifyNotebookTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyNotebookTags
// 修改Notebook标签
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BINDINGTAGSFAILED = "FailedOperation.BindingTagsFailed"
//  INTERNALERROR_BINDINGTAGSFAILED = "InternalError.BindingTagsFailed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyNotebookTags(request *ModifyNotebookTagsRequest) (response *ModifyNotebookTagsResponse, err error) {
    return c.ModifyNotebookTagsWithContext(context.Background(), request)
}

// ModifyNotebookTags
// 修改Notebook标签
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BINDINGTAGSFAILED = "FailedOperation.BindingTagsFailed"
//  INTERNALERROR_BINDINGTAGSFAILED = "InternalError.BindingTagsFailed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyNotebookTagsWithContext(ctx context.Context, request *ModifyNotebookTagsRequest) (response *ModifyNotebookTagsResponse, err error) {
    if request == nil {
        request = NewModifyNotebookTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNotebookTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNotebookTagsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyServiceGroupWeightsRequest() (request *ModifyServiceGroupWeightsRequest) {
    request = &ModifyServiceGroupWeightsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "ModifyServiceGroupWeights")
    
    
    return
}

func NewModifyServiceGroupWeightsResponse() (response *ModifyServiceGroupWeightsResponse) {
    response = &ModifyServiceGroupWeightsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyServiceGroupWeights
// 更新推理服务组流量分配
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_INVALIDUSERTYPE = "FailedOperation.InvalidUserType"
//  FAILEDOPERATION_KMSKEYNOTOPEN = "FailedOperation.KmsKeyNotOpen"
//  FAILEDOPERATION_NOFREEBUCKET = "FailedOperation.NoFreeBucket"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_PROCESSING = "FailedOperation.Processing"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_REPOBINDBYINSTANCE = "FailedOperation.RepoBindByInstance"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_TIMEDOUT = "FailedOperation.Timedout"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyServiceGroupWeights(request *ModifyServiceGroupWeightsRequest) (response *ModifyServiceGroupWeightsResponse, err error) {
    return c.ModifyServiceGroupWeightsWithContext(context.Background(), request)
}

// ModifyServiceGroupWeights
// 更新推理服务组流量分配
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_APIGATEWAYQUERYFAILED = "FailedOperation.ApiGatewayQueryFailed"
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  FAILEDOPERATION_CLUSTERQUERYFAILED = "FailedOperation.ClusterQueryFailed"
//  FAILEDOPERATION_DUPLICATENAME = "FailedOperation.DuplicateName"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_INVALIDUSERTYPE = "FailedOperation.InvalidUserType"
//  FAILEDOPERATION_KMSKEYNOTOPEN = "FailedOperation.KmsKeyNotOpen"
//  FAILEDOPERATION_NOFREEBUCKET = "FailedOperation.NoFreeBucket"
//  FAILEDOPERATION_NOPERMISSION = "FailedOperation.NoPermission"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_PROCESSING = "FailedOperation.Processing"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_REPOBINDBYINSTANCE = "FailedOperation.RepoBindByInstance"
//  FAILEDOPERATION_STSQUERYFAILED = "FailedOperation.StsQueryFailed"
//  FAILEDOPERATION_TIMEDOUT = "FailedOperation.Timedout"
//  FAILEDOPERATION_UNKNOWNINSTANCETYPE = "FailedOperation.UnknownInstanceType"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyServiceGroupWeightsWithContext(ctx context.Context, request *ModifyServiceGroupWeightsRequest) (response *ModifyServiceGroupWeightsResponse, err error) {
    if request == nil {
        request = NewModifyServiceGroupWeightsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyServiceGroupWeights require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyServiceGroupWeightsResponse()
    err = c.Send(request, response)
    return
}

func NewPushTrainingMetricsRequest() (request *PushTrainingMetricsRequest) {
    request = &PushTrainingMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "PushTrainingMetrics")
    
    
    return
}

func NewPushTrainingMetricsResponse() (response *PushTrainingMetricsResponse) {
    response = &PushTrainingMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PushTrainingMetrics
// 上报训练自定义指标
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) PushTrainingMetrics(request *PushTrainingMetricsRequest) (response *PushTrainingMetricsResponse, err error) {
    return c.PushTrainingMetricsWithContext(context.Background(), request)
}

// PushTrainingMetrics
// 上报训练自定义指标
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CLSSERVICENOTACTIVED = "FailedOperation.ClsServiceNotActived"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) PushTrainingMetricsWithContext(ctx context.Context, request *PushTrainingMetricsRequest) (response *PushTrainingMetricsResponse, err error) {
    if request == nil {
        request = NewPushTrainingMetricsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PushTrainingMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewPushTrainingMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewRestartModelAccelerateTaskRequest() (request *RestartModelAccelerateTaskRequest) {
    request = &RestartModelAccelerateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "RestartModelAccelerateTask")
    
    
    return
}

func NewRestartModelAccelerateTaskResponse() (response *RestartModelAccelerateTaskResponse) {
    response = &RestartModelAccelerateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestartModelAccelerateTask
// 重启模型加速任务
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) RestartModelAccelerateTask(request *RestartModelAccelerateTaskRequest) (response *RestartModelAccelerateTaskResponse, err error) {
    return c.RestartModelAccelerateTaskWithContext(context.Background(), request)
}

// RestartModelAccelerateTask
// 重启模型加速任务
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) RestartModelAccelerateTaskWithContext(ctx context.Context, request *RestartModelAccelerateTaskRequest) (response *RestartModelAccelerateTaskResponse, err error) {
    if request == nil {
        request = NewRestartModelAccelerateTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartModelAccelerateTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartModelAccelerateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSendChatMessageRequest() (request *SendChatMessageRequest) {
    request = &SendChatMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "SendChatMessage")
    
    
    return
}

func NewSendChatMessageResponse() (response *SendChatMessageResponse) {
    response = &SendChatMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendChatMessage
// 这是一个供您体验大模型聊天的接口。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SendChatMessage(request *SendChatMessageRequest) (response *SendChatMessageResponse, err error) {
    return c.SendChatMessageWithContext(context.Background(), request)
}

// SendChatMessage
// 这是一个供您体验大模型聊天的接口。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) SendChatMessageWithContext(ctx context.Context, request *SendChatMessageRequest) (response *SendChatMessageResponse, err error) {
    if request == nil {
        request = NewSendChatMessageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendChatMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendChatMessageResponse()
    err = c.Send(request, response)
    return
}

func NewStartNotebookRequest() (request *StartNotebookRequest) {
    request = &StartNotebookRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "StartNotebook")
    
    
    return
}

func NewStartNotebookResponse() (response *StartNotebookResponse) {
    response = &StartNotebookResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartNotebook
// 启动Notebook
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MODIFYBILLINGINSTANCEBATCHFAILED = "FailedOperation.ModifyBillingInstanceBatchFailed"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_QUERYRESOURCESPECFAILED = "FailedOperation.QueryResourceSpecFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GETCFSMOUNTIPFAILED = "InvalidParameterValue.GetCFSMountIPFailed"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_BILLINGSTATUSRESOURCEINSUFFICIENT = "OperationDenied.BillingStatusResourceInsufficient"
//  OPERATIONDENIED_IPILLEGAL = "OperationDenied.IpIllegal"
//  OPERATIONDENIED_MIYINGBALANCEINSUFFICIENT = "OperationDenied.MIYINGBalanceInsufficient"
//  OPERATIONDENIED_NOTALLOW = "OperationDenied.NotAllow"
//  OPERATIONDENIED_RESOURCEGROUPINSUFFICIENT = "OperationDenied.ResourceGroupInsufficient"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) StartNotebook(request *StartNotebookRequest) (response *StartNotebookResponse, err error) {
    return c.StartNotebookWithContext(context.Background(), request)
}

// StartNotebook
// 启动Notebook
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_MODIFYBILLINGINSTANCEBATCHFAILED = "FailedOperation.ModifyBillingInstanceBatchFailed"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_QUERYRESOURCESPECFAILED = "FailedOperation.QueryResourceSpecFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_GETCFSMOUNTIPFAILED = "InvalidParameterValue.GetCFSMountIPFailed"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_BILLINGSTATUSRESOURCEINSUFFICIENT = "OperationDenied.BillingStatusResourceInsufficient"
//  OPERATIONDENIED_IPILLEGAL = "OperationDenied.IpIllegal"
//  OPERATIONDENIED_MIYINGBALANCEINSUFFICIENT = "OperationDenied.MIYINGBalanceInsufficient"
//  OPERATIONDENIED_NOTALLOW = "OperationDenied.NotAllow"
//  OPERATIONDENIED_RESOURCEGROUPINSUFFICIENT = "OperationDenied.ResourceGroupInsufficient"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) StartNotebookWithContext(ctx context.Context, request *StartNotebookRequest) (response *StartNotebookResponse, err error) {
    if request == nil {
        request = NewStartNotebookRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartNotebook require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartNotebookResponse()
    err = c.Send(request, response)
    return
}

func NewStartTrainingTaskRequest() (request *StartTrainingTaskRequest) {
    request = &StartTrainingTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "StartTrainingTask")
    
    
    return
}

func NewStartTrainingTaskResponse() (response *StartTrainingTaskResponse) {
    response = &StartTrainingTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartTrainingTask
// 启动模型训练任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYRESOURCESPECFAILED = "FailedOperation.QueryResourceSpecFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEJOBINSTANCEFAILED = "InternalError.CreateJobInstanceFailed"
//  INTERNALERROR_CREATETCRINSTANCETOKENFAILED = "InternalError.CreateTcrInstanceTokenFailed"
//  INTERNALERROR_NOTALLOW = "InternalError.NotAllow"
//  INTERNALERROR_QUERYHDFSINFOFAILED = "InternalError.QueryHDFSInfoFailed"
//  INTERNALERROR_QUERYRESOURCESPECFAILED = "InternalError.QueryResourceSpecFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"
//  INVALIDPARAMETERVALUE_GETCFSMOUNTIPFAILED = "InvalidParameterValue.GetCFSMountIPFailed"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_BILLINGEXCEPTION = "OperationDenied.BillingException"
//  OPERATIONDENIED_BILLINGSTATUSRESOURCEINSUFFICIENT = "OperationDenied.BillingStatusResourceInsufficient"
//  OPERATIONDENIED_MIYINGBALANCEINSUFFICIENT = "OperationDenied.MIYINGBalanceInsufficient"
//  OPERATIONDENIED_NOTALLOW = "OperationDenied.NotAllow"
//  OPERATIONDENIED_RESOURCEGROUPINSUFFICIENT = "OperationDenied.ResourceGroupInsufficient"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) StartTrainingTask(request *StartTrainingTaskRequest) (response *StartTrainingTaskResponse, err error) {
    return c.StartTrainingTaskWithContext(context.Background(), request)
}

// StartTrainingTask
// 启动模型训练任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYRESOURCESPECFAILED = "FailedOperation.QueryResourceSpecFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATEJOBINSTANCEFAILED = "InternalError.CreateJobInstanceFailed"
//  INTERNALERROR_CREATETCRINSTANCETOKENFAILED = "InternalError.CreateTcrInstanceTokenFailed"
//  INTERNALERROR_NOTALLOW = "InternalError.NotAllow"
//  INTERNALERROR_QUERYHDFSINFOFAILED = "InternalError.QueryHDFSInfoFailed"
//  INTERNALERROR_QUERYRESOURCESPECFAILED = "InternalError.QueryResourceSpecFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"
//  INVALIDPARAMETERVALUE_GETCFSMOUNTIPFAILED = "InvalidParameterValue.GetCFSMountIPFailed"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_BILLINGEXCEPTION = "OperationDenied.BillingException"
//  OPERATIONDENIED_BILLINGSTATUSRESOURCEINSUFFICIENT = "OperationDenied.BillingStatusResourceInsufficient"
//  OPERATIONDENIED_MIYINGBALANCEINSUFFICIENT = "OperationDenied.MIYINGBalanceInsufficient"
//  OPERATIONDENIED_NOTALLOW = "OperationDenied.NotAllow"
//  OPERATIONDENIED_RESOURCEGROUPINSUFFICIENT = "OperationDenied.ResourceGroupInsufficient"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) StartTrainingTaskWithContext(ctx context.Context, request *StartTrainingTaskRequest) (response *StartTrainingTaskResponse, err error) {
    if request == nil {
        request = NewStartTrainingTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartTrainingTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartTrainingTaskResponse()
    err = c.Send(request, response)
    return
}

func NewStopBatchTaskRequest() (request *StopBatchTaskRequest) {
    request = &StopBatchTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "StopBatchTask")
    
    
    return
}

func NewStopBatchTaskResponse() (response *StopBatchTaskResponse) {
    response = &StopBatchTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopBatchTask
// 停止跑批任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) StopBatchTask(request *StopBatchTaskRequest) (response *StopBatchTaskResponse, err error) {
    return c.StopBatchTaskWithContext(context.Background(), request)
}

// StopBatchTask
// 停止跑批任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) StopBatchTaskWithContext(ctx context.Context, request *StopBatchTaskRequest) (response *StopBatchTaskResponse, err error) {
    if request == nil {
        request = NewStopBatchTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopBatchTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopBatchTaskResponse()
    err = c.Send(request, response)
    return
}

func NewStopCreatingImageRequest() (request *StopCreatingImageRequest) {
    request = &StopCreatingImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "StopCreatingImage")
    
    
    return
}

func NewStopCreatingImageResponse() (response *StopCreatingImageResponse) {
    response = &StopCreatingImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopCreatingImage
// 停止保存镜像
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) StopCreatingImage(request *StopCreatingImageRequest) (response *StopCreatingImageResponse, err error) {
    return c.StopCreatingImageWithContext(context.Background(), request)
}

// StopCreatingImage
// 停止保存镜像
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECTAGFAIL = "FailedOperation.ExecTagFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) StopCreatingImageWithContext(ctx context.Context, request *StopCreatingImageRequest) (response *StopCreatingImageResponse, err error) {
    if request == nil {
        request = NewStopCreatingImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopCreatingImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopCreatingImageResponse()
    err = c.Send(request, response)
    return
}

func NewStopModelAccelerateTaskRequest() (request *StopModelAccelerateTaskRequest) {
    request = &StopModelAccelerateTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "StopModelAccelerateTask")
    
    
    return
}

func NewStopModelAccelerateTaskResponse() (response *StopModelAccelerateTaskResponse) {
    response = &StopModelAccelerateTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopModelAccelerateTask
// 停止模型加速任务
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) StopModelAccelerateTask(request *StopModelAccelerateTaskRequest) (response *StopModelAccelerateTaskResponse, err error) {
    return c.StopModelAccelerateTaskWithContext(context.Background(), request)
}

// StopModelAccelerateTask
// 停止模型加速任务
//
// 可能返回的错误码:
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) StopModelAccelerateTaskWithContext(ctx context.Context, request *StopModelAccelerateTaskRequest) (response *StopModelAccelerateTaskResponse, err error) {
    if request == nil {
        request = NewStopModelAccelerateTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopModelAccelerateTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopModelAccelerateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewStopNotebookRequest() (request *StopNotebookRequest) {
    request = &StopNotebookRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "StopNotebook")
    
    
    return
}

func NewStopNotebookResponse() (response *StopNotebookResponse) {
    response = &StopNotebookResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopNotebook
// 停止Notebook
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_STOPJOBINSTANCEFAILED = "FailedOperation.StopJobInstanceFailed"
//  FAILEDOPERATION_UNSUBMITNOTALLOWTOSTOP = "FailedOperation.UnSubmitNotAllowToStop"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) StopNotebook(request *StopNotebookRequest) (response *StopNotebookResponse, err error) {
    return c.StopNotebookWithContext(context.Background(), request)
}

// StopNotebook
// 停止Notebook
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_STOPJOBINSTANCEFAILED = "FailedOperation.StopJobInstanceFailed"
//  FAILEDOPERATION_UNSUBMITNOTALLOWTOSTOP = "FailedOperation.UnSubmitNotAllowToStop"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) StopNotebookWithContext(ctx context.Context, request *StopNotebookRequest) (response *StopNotebookResponse, err error) {
    if request == nil {
        request = NewStopNotebookRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopNotebook require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopNotebookResponse()
    err = c.Send(request, response)
    return
}

func NewStopTrainingTaskRequest() (request *StopTrainingTaskRequest) {
    request = &StopTrainingTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "StopTrainingTask")
    
    
    return
}

func NewStopTrainingTaskResponse() (response *StopTrainingTaskResponse) {
    response = &StopTrainingTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopTrainingTask
// 停止模型训练任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NOTALLOW = "InternalError.NotAllow"
//  INTERNALERROR_STOPJOBINSTANCEFAILED = "InternalError.StopJobInstanceFailed"
//  INTERNALERROR_UNSUBMITTEDSTATUSNOTALLOWSTOP = "InternalError.UnSubmittedStatusNotAllowStop"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) StopTrainingTask(request *StopTrainingTaskRequest) (response *StopTrainingTaskResponse, err error) {
    return c.StopTrainingTaskWithContext(context.Background(), request)
}

// StopTrainingTask
// 停止模型训练任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_NOTALLOW = "InternalError.NotAllow"
//  INTERNALERROR_STOPJOBINSTANCEFAILED = "InternalError.StopJobInstanceFailed"
//  INTERNALERROR_UNSUBMITTEDSTATUSNOTALLOWSTOP = "InternalError.UnSubmittedStatusNotAllowStop"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) StopTrainingTaskWithContext(ctx context.Context, request *StopTrainingTaskRequest) (response *StopTrainingTaskResponse, err error) {
    if request == nil {
        request = NewStopTrainingTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopTrainingTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopTrainingTaskResponse()
    err = c.Send(request, response)
    return
}
