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
// 该接口支持与自行部署的大模型的聊天。
//
// 
//
// 使用该接口调用时需要携带腾讯云的密钥信息用于身份信息鉴权，建议通过腾讯云的云 API SDK调用，具体可以参考
//
// https://cloud.tencent.com/document/product/1278/85305
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDAUTHTOKEN = "FailedOperation.InvalidAuthToken"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MODELISOFFLINE = "InvalidParameter.ModelIsOffline"
//  INVALIDPARAMETER_TGWINVALIDINTERFACE = "InvalidParameter.TgwInvalidInterface"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ChatCompletion(request *ChatCompletionRequest) (response *ChatCompletionResponse, err error) {
    return c.ChatCompletionWithContext(context.Background(), request)
}

// ChatCompletion
// 该接口支持与自行部署的大模型的聊天。
//
// 
//
// 使用该接口调用时需要携带腾讯云的密钥信息用于身份信息鉴权，建议通过腾讯云的云 API SDK调用，具体可以参考
//
// https://cloud.tencent.com/document/product/1278/85305
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INVALIDAUTHTOKEN = "FailedOperation.InvalidAuthToken"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MODELISOFFLINE = "InvalidParameter.ModelIsOffline"
//  INVALIDPARAMETER_TGWINVALIDINTERFACE = "InvalidParameter.TgwInvalidInterface"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ChatCompletionWithContext(ctx context.Context, request *ChatCompletionRequest) (response *ChatCompletionResponse, err error) {
    if request == nil {
        request = NewChatCompletionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "ChatCompletion")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChatCompletion require credential")
    }

    request.SetContext(ctx)
    
    response = NewChatCompletionResponse()
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
//  INVALIDPARAMETERVALUE_CFSPATHCOLLISION = "InvalidParameterValue.CfsPathCollision"
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
//  INVALIDPARAMETERVALUE_CFSPATHCOLLISION = "InvalidParameterValue.CfsPathCollision"
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "CreateDataset")
    
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
//  INVALIDPARAMETER_VALIDATEERROR = "InvalidParameter.ValidateError"
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
//  INVALIDPARAMETER_VALIDATEERROR = "InvalidParameter.ValidateError"
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "CreateModelService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateModelService require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateModelServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateModelServiceAuthTokenRequest() (request *CreateModelServiceAuthTokenRequest) {
    request = &CreateModelServiceAuthTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "CreateModelServiceAuthToken")
    
    
    return
}

func NewCreateModelServiceAuthTokenResponse() (response *CreateModelServiceAuthTokenResponse) {
    response = &CreateModelServiceAuthTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateModelServiceAuthToken
// 创建一个 AuthToken
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateModelServiceAuthToken(request *CreateModelServiceAuthTokenRequest) (response *CreateModelServiceAuthTokenResponse, err error) {
    return c.CreateModelServiceAuthTokenWithContext(context.Background(), request)
}

// CreateModelServiceAuthToken
// 创建一个 AuthToken
//
// 可能返回的错误码:
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) CreateModelServiceAuthTokenWithContext(ctx context.Context, request *CreateModelServiceAuthTokenRequest) (response *CreateModelServiceAuthTokenResponse, err error) {
    if request == nil {
        request = NewCreateModelServiceAuthTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "CreateModelServiceAuthToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateModelServiceAuthToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateModelServiceAuthTokenResponse()
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
//  FAILEDOPERATION_UNAUTHORIZEDOPERATION = "FailedOperation.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATECUSTOMSECRETFAILED = "InternalError.CreateCustomSecretFailed"
//  INTERNALERROR_DELETECUSTOMSECRETFAILED = "InternalError.DeleteCustomSecretFailed"
//  INTERNALERROR_DESCRIBECUSTOMSECRETFAILED = "InternalError.DescribeCustomSecretFailed"
//  INTERNALERROR_DESCRIBESUBUINDEPENDENCYACCESSIBILITYFAILED = "InternalError.DescribeSubuinDependencyAccessibilityFailed"
//  INTERNALERROR_MODIFYCUSTOMSECRETFAILED = "InternalError.ModifyCustomSecretFailed"
//  INTERNALERROR_QUERYRESOURCEGROUPFAILED = "InternalError.QueryResourceGroupFailed"
//  INTERNALERROR_QUERYSUBNETINFOFAILED = "InternalError.QuerySubnetInfoFailed"
//  INTERNALERROR_QUERYUSERTMPCREDENTIALFAILED = "InternalError.QueryUserTMPCredentialFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTEDMULTILOCALDISK = "InvalidParameter.UnsupportedMultiLocalDisk"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CCNHASNOATTACHEDVPC = "InvalidParameterValue.CcnHasNoAttachedVpc"
//  INVALIDPARAMETERVALUE_CODEREPODUPLICATED = "InvalidParameterValue.CodeRepoDuplicated"
//  INVALIDPARAMETERVALUE_CODEREPONOTFOUND = "InvalidParameterValue.CodeRepoNotFound"
//  INVALIDPARAMETERVALUE_DATACONFIGNUMLIMITEXCEEDED = "InvalidParameterValue.DataConfigNumLimitExceeded"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_GETCFSMOUNTIPFAILED = "InvalidParameterValue.GetCFSMountIPFailed"
//  INVALIDPARAMETERVALUE_GETGOOSEFSFAILED = "InvalidParameterValue.GetGooseFSFailed"
//  INVALIDPARAMETERVALUE_GOOSEFSCONFIGCANNOTEMPTY = "InvalidParameterValue.GooseFSConfigCannotEmpty"
//  INVALIDPARAMETERVALUE_GOOSEFSNOTEXIST = "InvalidParameterValue.GooseFSNotExist"
//  INVALIDPARAMETERVALUE_IMAGEADDRESSILLEGAL = "InvalidParameterValue.ImageAddressIllegal"
//  INVALIDPARAMETERVALUE_INVALIDIMAGENAME = "InvalidParameterValue.InvalidImageName"
//  INVALIDPARAMETERVALUE_INVALIDRESOURCESPEC = "InvalidParameterValue.InvalidResourceSpec"
//  INVALIDPARAMETERVALUE_LIFECYCLENOTFOUND = "InvalidParameterValue.LifecycleNotFound"
//  INVALIDPARAMETERVALUE_NOTALLOWEDTOCREATEFREEVOLUMENOTEBOOKWITHBAREMETALRESOURCEGROUP = "InvalidParameterValue.NotAllowedToCreateFreeVolumeNotebookWithBareMetalResourceGroup"
//  INVALIDPARAMETERVALUE_NOTALLOWEDTOCREATEFREEVOLUMENOTEBOOKWITHSWRESOURCEGROUP = "InvalidParameterValue.NotAllowedToCreateFreeVolumeNotebookWithSWResourceGroup"
//  INVALIDPARAMETERVALUE_NOTALLOWEDTOCREATENOTEBOOKWITHBAREMETALRESOURCEGROUP = "InvalidParameterValue.NotAllowedToCreateNotebookWithBareMetalResourceGroup"
//  INVALIDPARAMETERVALUE_PATHILLEGAL = "InvalidParameterValue.PathIllegal"
//  INVALIDPARAMETERVALUE_RESOURCECONFIGILLEGAL = "InvalidParameterValue.ResourceConfigIllegal"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDGOOSEFSCONFIG = "InvalidParameterValue.UnsupportedGooseFSConfig"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDMULTICBSSTORAGE = "InvalidParameterValue.UnsupportedMultiCBSStorage"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_BILLINGSTATUSNOTSUPPORTWHITELISTGPUTYPE = "OperationDenied.BillingStatusNotSupportWhitelistGPUType"
//  OPERATIONDENIED_BILLINGSTATUSRESOURCEINSUFFICIENT = "OperationDenied.BillingStatusResourceInsufficient"
//  OPERATIONDENIED_IPILLEGAL = "OperationDenied.IpIllegal"
//  OPERATIONDENIED_MIYINGBALANCEINSUFFICIENT = "OperationDenied.MIYINGBalanceInsufficient"
//  OPERATIONDENIED_NETWORKCIDRILLEGAL = "OperationDenied.NetworkCidrIllegal"
//  OPERATIONDENIED_NOTALLOW = "OperationDenied.NotAllow"
//  OPERATIONDENIED_NOTSUPPORTCUSTOMRESOURCE = "OperationDenied.NotSupportCustomResource"
//  OPERATIONDENIED_RESOURCEGROUPINSUFFICIENT = "OperationDenied.ResourceGroupInsufficient"
//  OPERATIONDENIED_UNAUTHORIZEDOPERATION = "OperationDenied.UnauthorizedOperation"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  RESOURCEINSUFFICIENT_SSHPORTISCONSUMEDUP = "ResourceInsufficient.SSHPortIsConsumedUp"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCEGROUPNOTFOUND = "ResourceNotFound.ResourceGroupNotFound"
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
//  FAILEDOPERATION_UNAUTHORIZEDOPERATION = "FailedOperation.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CREATECUSTOMSECRETFAILED = "InternalError.CreateCustomSecretFailed"
//  INTERNALERROR_DELETECUSTOMSECRETFAILED = "InternalError.DeleteCustomSecretFailed"
//  INTERNALERROR_DESCRIBECUSTOMSECRETFAILED = "InternalError.DescribeCustomSecretFailed"
//  INTERNALERROR_DESCRIBESUBUINDEPENDENCYACCESSIBILITYFAILED = "InternalError.DescribeSubuinDependencyAccessibilityFailed"
//  INTERNALERROR_MODIFYCUSTOMSECRETFAILED = "InternalError.ModifyCustomSecretFailed"
//  INTERNALERROR_QUERYRESOURCEGROUPFAILED = "InternalError.QueryResourceGroupFailed"
//  INTERNALERROR_QUERYSUBNETINFOFAILED = "InternalError.QuerySubnetInfoFailed"
//  INTERNALERROR_QUERYUSERTMPCREDENTIALFAILED = "InternalError.QueryUserTMPCredentialFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTEDMULTILOCALDISK = "InvalidParameter.UnsupportedMultiLocalDisk"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CCNHASNOATTACHEDVPC = "InvalidParameterValue.CcnHasNoAttachedVpc"
//  INVALIDPARAMETERVALUE_CODEREPODUPLICATED = "InvalidParameterValue.CodeRepoDuplicated"
//  INVALIDPARAMETERVALUE_CODEREPONOTFOUND = "InvalidParameterValue.CodeRepoNotFound"
//  INVALIDPARAMETERVALUE_DATACONFIGNUMLIMITEXCEEDED = "InvalidParameterValue.DataConfigNumLimitExceeded"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_GETCFSMOUNTIPFAILED = "InvalidParameterValue.GetCFSMountIPFailed"
//  INVALIDPARAMETERVALUE_GETGOOSEFSFAILED = "InvalidParameterValue.GetGooseFSFailed"
//  INVALIDPARAMETERVALUE_GOOSEFSCONFIGCANNOTEMPTY = "InvalidParameterValue.GooseFSConfigCannotEmpty"
//  INVALIDPARAMETERVALUE_GOOSEFSNOTEXIST = "InvalidParameterValue.GooseFSNotExist"
//  INVALIDPARAMETERVALUE_IMAGEADDRESSILLEGAL = "InvalidParameterValue.ImageAddressIllegal"
//  INVALIDPARAMETERVALUE_INVALIDIMAGENAME = "InvalidParameterValue.InvalidImageName"
//  INVALIDPARAMETERVALUE_INVALIDRESOURCESPEC = "InvalidParameterValue.InvalidResourceSpec"
//  INVALIDPARAMETERVALUE_LIFECYCLENOTFOUND = "InvalidParameterValue.LifecycleNotFound"
//  INVALIDPARAMETERVALUE_NOTALLOWEDTOCREATEFREEVOLUMENOTEBOOKWITHBAREMETALRESOURCEGROUP = "InvalidParameterValue.NotAllowedToCreateFreeVolumeNotebookWithBareMetalResourceGroup"
//  INVALIDPARAMETERVALUE_NOTALLOWEDTOCREATEFREEVOLUMENOTEBOOKWITHSWRESOURCEGROUP = "InvalidParameterValue.NotAllowedToCreateFreeVolumeNotebookWithSWResourceGroup"
//  INVALIDPARAMETERVALUE_NOTALLOWEDTOCREATENOTEBOOKWITHBAREMETALRESOURCEGROUP = "InvalidParameterValue.NotAllowedToCreateNotebookWithBareMetalResourceGroup"
//  INVALIDPARAMETERVALUE_PATHILLEGAL = "InvalidParameterValue.PathIllegal"
//  INVALIDPARAMETERVALUE_RESOURCECONFIGILLEGAL = "InvalidParameterValue.ResourceConfigIllegal"
//  INVALIDPARAMETERVALUE_SUBNETNOTFOUND = "InvalidParameterValue.SubnetNotFound"
//  INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDGOOSEFSCONFIG = "InvalidParameterValue.UnsupportedGooseFSConfig"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDMULTICBSSTORAGE = "InvalidParameterValue.UnsupportedMultiCBSStorage"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_BILLINGSTATUSNOTSUPPORTWHITELISTGPUTYPE = "OperationDenied.BillingStatusNotSupportWhitelistGPUType"
//  OPERATIONDENIED_BILLINGSTATUSRESOURCEINSUFFICIENT = "OperationDenied.BillingStatusResourceInsufficient"
//  OPERATIONDENIED_IPILLEGAL = "OperationDenied.IpIllegal"
//  OPERATIONDENIED_MIYINGBALANCEINSUFFICIENT = "OperationDenied.MIYINGBalanceInsufficient"
//  OPERATIONDENIED_NETWORKCIDRILLEGAL = "OperationDenied.NetworkCidrIllegal"
//  OPERATIONDENIED_NOTALLOW = "OperationDenied.NotAllow"
//  OPERATIONDENIED_NOTSUPPORTCUSTOMRESOURCE = "OperationDenied.NotSupportCustomResource"
//  OPERATIONDENIED_RESOURCEGROUPINSUFFICIENT = "OperationDenied.ResourceGroupInsufficient"
//  OPERATIONDENIED_UNAUTHORIZEDOPERATION = "OperationDenied.UnauthorizedOperation"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  RESOURCEINSUFFICIENT_SSHPORTISCONSUMEDUP = "ResourceInsufficient.SSHPortIsConsumedUp"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCEGROUPNOTFOUND = "ResourceNotFound.ResourceGroupNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateNotebookWithContext(ctx context.Context, request *CreateNotebookRequest) (response *CreateNotebookResponse, err error) {
    if request == nil {
        request = NewCreateNotebookRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "CreateNotebook")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateNotebook require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateNotebookResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePresignedNotebookUrlRequest() (request *CreatePresignedNotebookUrlRequest) {
    request = &CreatePresignedNotebookUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "CreatePresignedNotebookUrl")
    
    
    return
}

func NewCreatePresignedNotebookUrlResponse() (response *CreatePresignedNotebookUrlResponse) {
    response = &CreatePresignedNotebookUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePresignedNotebookUrl
// 生成Notebook访问链接
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_PODIPNOTFOUND = "FailedOperation.PodIpNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYUSERTMPCREDENTIALFAILED = "InternalError.QueryUserTMPCredentialFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreatePresignedNotebookUrl(request *CreatePresignedNotebookUrlRequest) (response *CreatePresignedNotebookUrlResponse, err error) {
    return c.CreatePresignedNotebookUrlWithContext(context.Background(), request)
}

// CreatePresignedNotebookUrl
// 生成Notebook访问链接
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  FAILEDOPERATION_PODIPNOTFOUND = "FailedOperation.PodIpNotFound"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_QUERYUSERTMPCREDENTIALFAILED = "InternalError.QueryUserTMPCredentialFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) CreatePresignedNotebookUrlWithContext(ctx context.Context, request *CreatePresignedNotebookUrlRequest) (response *CreatePresignedNotebookUrlResponse, err error) {
    if request == nil {
        request = NewCreatePresignedNotebookUrlRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "CreatePresignedNotebookUrl")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePresignedNotebookUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePresignedNotebookUrlResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "CreateTrainingModel")
    
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
//  INTERNALERROR_CFSNOTFOUND = "InternalError.CFSNotFound"
//  INTERNALERROR_CHECKFSPATHACCESSIBILITYFAILED = "InternalError.CheckFSPathAccessibilityFailed"
//  INTERNALERROR_CREATETCRINSTANCETOKENFAILED = "InternalError.CreateTcrInstanceTokenFailed"
//  INTERNALERROR_GETCFSFILESYSTEMSFAILED = "InternalError.GetCFSFileSystemsFailed"
//  INTERNALERROR_GETCFSMOUNTINFOFAILED = "InternalError.GetCFSMountInfoFailed"
//  INTERNALERROR_QUERYHDFSINFOFAILED = "InternalError.QueryHDFSInfoFailed"
//  INTERNALERROR_QUERYRESOURCEGROUPFAILED = "InternalError.QueryResourceGroupFailed"
//  INTERNALERROR_QUERYRESOURCESPECFAILED = "InternalError.QueryResourceSpecFailed"
//  INTERNALERROR_QUERYSUBNETINFOFAILED = "InternalError.QuerySubnetInfoFailed"
//  INTERNALERROR_QUERYVPCINFOFAILED = "InternalError.QueryVPCInfoFailed"
//  INTERNALERROR_VALIDATECREATETASKFAILED = "InternalError.ValidateCreateTaskFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AIMARKETOUTPUTCONFIGEMPTY = "InvalidParameterValue.AIMarketOutputConfigEmpty"
//  INVALIDPARAMETERVALUE_AIMARKETPUBLICALGOVERSIONNOTEXIST = "InvalidParameterValue.AIMarketPublicAlgoVersionNotExist"
//  INVALIDPARAMETERVALUE_BACKOFFLIMITILLEGAL = "InvalidParameterValue.BackOffLimitIllegal"
//  INVALIDPARAMETERVALUE_BACKOFFLIMITNOTSUPPORT = "InvalidParameterValue.BackOffLimitNotSupport"
//  INVALIDPARAMETERVALUE_COSPATHNOTEXIST = "InvalidParameterValue.CosPathNotExist"
//  INVALIDPARAMETERVALUE_DATASETNUMLIMITEXCEEDED = "InvalidParameterValue.DatasetNumLimitExceeded"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_FSPATHINACCESSIBLE = "InvalidParameterValue.FSPathInaccessible"
//  INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"
//  INVALIDPARAMETERVALUE_GETCFSMOUNTIPFAILED = "InvalidParameterValue.GetCFSMountIPFailed"
//  INVALIDPARAMETERVALUE_GETGOOSEFSFAILED = "InvalidParameterValue.GetGooseFSFailed"
//  INVALIDPARAMETERVALUE_GOOSEFSNOTEXIST = "InvalidParameterValue.GooseFSNotExist"
//  INVALIDPARAMETERVALUE_IMAGEILLEGAL = "InvalidParameterValue.ImageIllegal"
//  INVALIDPARAMETERVALUE_IMAGENOTFOUND = "InvalidParameterValue.ImageNotFound"
//  INVALIDPARAMETERVALUE_NOTALLOW = "InvalidParameterValue.NotAllow"
//  INVALIDPARAMETERVALUE_PARAMLENGTHEXCEEDLIMIT = "InvalidParameterValue.ParamLengthExceedLimit"
//  INVALIDPARAMETERVALUE_PATHILLEGAL = "InvalidParameterValue.PathIllegal"
//  INVALIDPARAMETERVALUE_QUERYVPCINFOFAILED = "InvalidParameterValue.QueryVPCInfoFailed"
//  INVALIDPARAMETERVALUE_RDMACONFIGILLEGAL = "InvalidParameterValue.RDMAConfigIllegal"
//  INVALIDPARAMETERVALUE_RESOURCECONFIGILLEGAL = "InvalidParameterValue.ResourceConfigIllegal"
//  INVALIDPARAMETERVALUE_TAIJIRESOURCECONFIGILLEGAL = "InvalidParameterValue.TAIJIResourceConfigIllegal"
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
//  OPERATIONDENIED_TAIJIAPPLICATIONGROUPINSUFFICIENT = "OperationDenied.TAIJIApplicationGroupInsufficient"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  RESOURCENOTFOUND_CFSNOTFOUND = "ResourceNotFound.CfsNotFound"
//  RESOURCENOTFOUND_VPCNOTFOUND = "ResourceNotFound.VPCNotFound"
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
//  INTERNALERROR_CFSNOTFOUND = "InternalError.CFSNotFound"
//  INTERNALERROR_CHECKFSPATHACCESSIBILITYFAILED = "InternalError.CheckFSPathAccessibilityFailed"
//  INTERNALERROR_CREATETCRINSTANCETOKENFAILED = "InternalError.CreateTcrInstanceTokenFailed"
//  INTERNALERROR_GETCFSFILESYSTEMSFAILED = "InternalError.GetCFSFileSystemsFailed"
//  INTERNALERROR_GETCFSMOUNTINFOFAILED = "InternalError.GetCFSMountInfoFailed"
//  INTERNALERROR_QUERYHDFSINFOFAILED = "InternalError.QueryHDFSInfoFailed"
//  INTERNALERROR_QUERYRESOURCEGROUPFAILED = "InternalError.QueryResourceGroupFailed"
//  INTERNALERROR_QUERYRESOURCESPECFAILED = "InternalError.QueryResourceSpecFailed"
//  INTERNALERROR_QUERYSUBNETINFOFAILED = "InternalError.QuerySubnetInfoFailed"
//  INTERNALERROR_QUERYVPCINFOFAILED = "InternalError.QueryVPCInfoFailed"
//  INTERNALERROR_VALIDATECREATETASKFAILED = "InternalError.ValidateCreateTaskFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AIMARKETOUTPUTCONFIGEMPTY = "InvalidParameterValue.AIMarketOutputConfigEmpty"
//  INVALIDPARAMETERVALUE_AIMARKETPUBLICALGOVERSIONNOTEXIST = "InvalidParameterValue.AIMarketPublicAlgoVersionNotExist"
//  INVALIDPARAMETERVALUE_BACKOFFLIMITILLEGAL = "InvalidParameterValue.BackOffLimitIllegal"
//  INVALIDPARAMETERVALUE_BACKOFFLIMITNOTSUPPORT = "InvalidParameterValue.BackOffLimitNotSupport"
//  INVALIDPARAMETERVALUE_COSPATHNOTEXIST = "InvalidParameterValue.CosPathNotExist"
//  INVALIDPARAMETERVALUE_DATASETNUMLIMITEXCEEDED = "InvalidParameterValue.DatasetNumLimitExceeded"
//  INVALIDPARAMETERVALUE_DUPLICATENAME = "InvalidParameterValue.DuplicateName"
//  INVALIDPARAMETERVALUE_FSPATHINACCESSIBLE = "InvalidParameterValue.FSPathInaccessible"
//  INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"
//  INVALIDPARAMETERVALUE_GETCFSMOUNTIPFAILED = "InvalidParameterValue.GetCFSMountIPFailed"
//  INVALIDPARAMETERVALUE_GETGOOSEFSFAILED = "InvalidParameterValue.GetGooseFSFailed"
//  INVALIDPARAMETERVALUE_GOOSEFSNOTEXIST = "InvalidParameterValue.GooseFSNotExist"
//  INVALIDPARAMETERVALUE_IMAGEILLEGAL = "InvalidParameterValue.ImageIllegal"
//  INVALIDPARAMETERVALUE_IMAGENOTFOUND = "InvalidParameterValue.ImageNotFound"
//  INVALIDPARAMETERVALUE_NOTALLOW = "InvalidParameterValue.NotAllow"
//  INVALIDPARAMETERVALUE_PARAMLENGTHEXCEEDLIMIT = "InvalidParameterValue.ParamLengthExceedLimit"
//  INVALIDPARAMETERVALUE_PATHILLEGAL = "InvalidParameterValue.PathIllegal"
//  INVALIDPARAMETERVALUE_QUERYVPCINFOFAILED = "InvalidParameterValue.QueryVPCInfoFailed"
//  INVALIDPARAMETERVALUE_RDMACONFIGILLEGAL = "InvalidParameterValue.RDMAConfigIllegal"
//  INVALIDPARAMETERVALUE_RESOURCECONFIGILLEGAL = "InvalidParameterValue.ResourceConfigIllegal"
//  INVALIDPARAMETERVALUE_TAIJIRESOURCECONFIGILLEGAL = "InvalidParameterValue.TAIJIResourceConfigIllegal"
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
//  OPERATIONDENIED_TAIJIAPPLICATIONGROUPINSUFFICIENT = "OperationDenied.TAIJIApplicationGroupInsufficient"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  RESOURCENOTFOUND_CFSNOTFOUND = "ResourceNotFound.CfsNotFound"
//  RESOURCENOTFOUND_VPCNOTFOUND = "ResourceNotFound.VPCNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateTrainingTaskWithContext(ctx context.Context, request *CreateTrainingTaskRequest) (response *CreateTrainingTaskResponse, err error) {
    if request == nil {
        request = NewCreateTrainingTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "CreateTrainingTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTrainingTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTrainingTaskResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DeleteDataset")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDataset require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDatasetResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DeleteModelService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteModelService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteModelServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteModelServiceAuthTokenRequest() (request *DeleteModelServiceAuthTokenRequest) {
    request = &DeleteModelServiceAuthTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DeleteModelServiceAuthToken")
    
    
    return
}

func NewDeleteModelServiceAuthTokenResponse() (response *DeleteModelServiceAuthTokenResponse) {
    response = &DeleteModelServiceAuthTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteModelServiceAuthToken
// 删除一个 AuthToken
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteModelServiceAuthToken(request *DeleteModelServiceAuthTokenRequest) (response *DeleteModelServiceAuthTokenResponse, err error) {
    return c.DeleteModelServiceAuthTokenWithContext(context.Background(), request)
}

// DeleteModelServiceAuthToken
// 删除一个 AuthToken
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteModelServiceAuthTokenWithContext(ctx context.Context, request *DeleteModelServiceAuthTokenRequest) (response *DeleteModelServiceAuthTokenResponse, err error) {
    if request == nil {
        request = NewDeleteModelServiceAuthTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DeleteModelServiceAuthToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteModelServiceAuthToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteModelServiceAuthTokenResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DeleteModelServiceGroup")
    
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
//  FAILEDOPERATION_RELEASESSHPORTFAILED = "FailedOperation.ReleaseSSHPortFailed"
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
//  FAILEDOPERATION_RELEASESSHPORTFAILED = "FailedOperation.ReleaseSSHPortFailed"
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DeleteNotebook")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteNotebook require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteNotebookResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DeleteTrainingModel")
    
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DeleteTrainingModelVersion")
    
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DeleteTrainingTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTrainingTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTrainingTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillingResourceGroupRequest() (request *DescribeBillingResourceGroupRequest) {
    request = &DescribeBillingResourceGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeBillingResourceGroup")
    
    
    return
}

func NewDescribeBillingResourceGroupResponse() (response *DescribeBillingResourceGroupResponse) {
    response = &DescribeBillingResourceGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillingResourceGroup
// 查询资源组节点列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_NOPERMISSION = "AuthFailure.NoPermission"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLCLUSTERFAIL = "FailedOperation.CallClusterFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_SERVERROOMQUERYFAILED = "FailedOperation.ServerRoomQueryFailed"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_RESOURCEGROUPNOTFOUND = "ResourceNotFound.ResourceGroupNotFound"
func (c *Client) DescribeBillingResourceGroup(request *DescribeBillingResourceGroupRequest) (response *DescribeBillingResourceGroupResponse, err error) {
    return c.DescribeBillingResourceGroupWithContext(context.Background(), request)
}

// DescribeBillingResourceGroup
// 查询资源组节点列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_NOPERMISSION = "AuthFailure.NoPermission"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLCLUSTERFAIL = "FailedOperation.CallClusterFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_QUERYTAGFAIL = "FailedOperation.QueryTagFail"
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
//  FAILEDOPERATION_SERVERROOMQUERYFAILED = "FailedOperation.ServerRoomQueryFailed"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDFILTER = "InvalidParameterValue.InvalidFilter"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_RESOURCEGROUPNOTFOUND = "ResourceNotFound.ResourceGroupNotFound"
func (c *Client) DescribeBillingResourceGroupWithContext(ctx context.Context, request *DescribeBillingResourceGroupRequest) (response *DescribeBillingResourceGroupResponse, err error) {
    if request == nil {
        request = NewDescribeBillingResourceGroupRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeBillingResourceGroup")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillingResourceGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillingResourceGroupResponse()
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
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
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
//  FAILEDOPERATION_RECORDNOTFOUND = "FailedOperation.RecordNotFound"
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeBillingResourceGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillingResourceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillingResourceGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillingResourceInstanceRunningJobsRequest() (request *DescribeBillingResourceInstanceRunningJobsRequest) {
    request = &DescribeBillingResourceInstanceRunningJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeBillingResourceInstanceRunningJobs")
    
    
    return
}

func NewDescribeBillingResourceInstanceRunningJobsResponse() (response *DescribeBillingResourceInstanceRunningJobsResponse) {
    response = &DescribeBillingResourceInstanceRunningJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBillingResourceInstanceRunningJobs
// 查询资源组节点运行中的任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLCLUSTERFAIL = "FailedOperation.CallClusterFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TGWINVALIDINTERFACE = "InvalidParameter.TgwInvalidInterface"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBillingResourceInstanceRunningJobs(request *DescribeBillingResourceInstanceRunningJobsRequest) (response *DescribeBillingResourceInstanceRunningJobsResponse, err error) {
    return c.DescribeBillingResourceInstanceRunningJobsWithContext(context.Background(), request)
}

// DescribeBillingResourceInstanceRunningJobs
// 查询资源组节点运行中的任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CALLCLUSTERFAIL = "FailedOperation.CallClusterFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TGWINVALIDINTERFACE = "InvalidParameter.TgwInvalidInterface"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeBillingResourceInstanceRunningJobsWithContext(ctx context.Context, request *DescribeBillingResourceInstanceRunningJobsRequest) (response *DescribeBillingResourceInstanceRunningJobsResponse, err error) {
    if request == nil {
        request = NewDescribeBillingResourceInstanceRunningJobsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeBillingResourceInstanceRunningJobs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillingResourceInstanceRunningJobs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillingResourceInstanceRunningJobsResponse()
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
// 本接口(DescribeBillingSpecs) 提供查询计费项列表
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
// 本接口(DescribeBillingSpecs) 提供查询计费项列表
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeBillingSpecs")
    
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeBillingSpecsPrice")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillingSpecsPrice require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillingSpecsPriceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBuildInImagesRequest() (request *DescribeBuildInImagesRequest) {
    request = &DescribeBuildInImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeBuildInImages")
    
    
    return
}

func NewDescribeBuildInImagesResponse() (response *DescribeBuildInImagesResponse) {
    response = &DescribeBuildInImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBuildInImages
// 获取内置镜像列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_QUERYBINDINGTAGSFAILED = "FailedOperation.QueryBindingTagsFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeBuildInImages(request *DescribeBuildInImagesRequest) (response *DescribeBuildInImagesResponse, err error) {
    return c.DescribeBuildInImagesWithContext(context.Background(), request)
}

// DescribeBuildInImages
// 获取内置镜像列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_QUERYBINDINGTAGSFAILED = "FailedOperation.QueryBindingTagsFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeBuildInImagesWithContext(ctx context.Context, request *DescribeBuildInImagesRequest) (response *DescribeBuildInImagesResponse, err error) {
    if request == nil {
        request = NewDescribeBuildInImagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeBuildInImages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBuildInImages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBuildInImagesResponse()
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
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
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
//  FAILEDOPERATION_BILLINGQUERYFAILED = "FailedOperation.BillingQueryFailed"
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeDatasets")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatasets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatasetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventsRequest() (request *DescribeEventsRequest) {
    request = &DescribeEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeEvents")
    
    
    return
}

func NewDescribeEventsResponse() (response *DescribeEventsResponse) {
    response = &DescribeEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEvents
// 获取任务式建模训练任务，Notebook，在线服务和批量预测任务的事件API
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEvents(request *DescribeEventsRequest) (response *DescribeEventsResponse, err error) {
    return c.DescribeEventsWithContext(context.Background(), request)
}

// DescribeEvents
// 获取任务式建模训练任务，Notebook，在线服务和批量预测任务的事件API
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeEventsWithContext(ctx context.Context, request *DescribeEventsRequest) (response *DescribeEventsResponse, err error) {
    if request == nil {
        request = NewDescribeEventsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeEvents")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEventsResponse()
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
// 已废弃，收敛到统一接口
//
// 
//
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
// 已废弃，收敛到统一接口
//
// 
//
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeInferTemplates")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInferTemplates require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInferTemplatesResponse()
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
// 获取任务式建模训练任务，Notebook，在线服务和批量预测任务的日志API
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONTEXTLIMITERROR = "FailedOperation.ContextLimitError"
func (c *Client) DescribeLogs(request *DescribeLogsRequest) (response *DescribeLogsResponse, err error) {
    return c.DescribeLogsWithContext(context.Background(), request)
}

// DescribeLogs
// 获取任务式建模训练任务，Notebook，在线服务和批量预测任务的日志API
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CONTEXTLIMITERROR = "FailedOperation.ContextLimitError"
func (c *Client) DescribeLogsWithContext(ctx context.Context, request *DescribeLogsRequest) (response *DescribeLogsResponse, err error) {
    if request == nil {
        request = NewDescribeLogsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeLogs")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogsResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeModelAccelerateTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelAccelerateTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelAccelerateTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeModelAccelerateVersionsRequest() (request *DescribeModelAccelerateVersionsRequest) {
    request = &DescribeModelAccelerateVersionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribeModelAccelerateVersions")
    
    
    return
}

func NewDescribeModelAccelerateVersionsResponse() (response *DescribeModelAccelerateVersionsResponse) {
    response = &DescribeModelAccelerateVersionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeModelAccelerateVersions
// 模型加速之后的模型版本列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_NOPERMISSION = "AuthFailure.NoPermission"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMFAILURE = "FailedOperation.CAMFailure"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeModelAccelerateVersions(request *DescribeModelAccelerateVersionsRequest) (response *DescribeModelAccelerateVersionsResponse, err error) {
    return c.DescribeModelAccelerateVersionsWithContext(context.Background(), request)
}

// DescribeModelAccelerateVersions
// 模型加速之后的模型版本列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CAMEXCEPTION = "AuthFailure.CamException"
//  AUTHFAILURE_NOPERMISSION = "AuthFailure.NoPermission"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CAMFAILURE = "FailedOperation.CAMFailure"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  FAILEDOPERATION_UNMARSHALDATA = "FailedOperation.UnmarshalData"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeModelAccelerateVersionsWithContext(ctx context.Context, request *DescribeModelAccelerateVersionsRequest) (response *DescribeModelAccelerateVersionsResponse, err error) {
    if request == nil {
        request = NewDescribeModelAccelerateVersionsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeModelAccelerateVersions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelAccelerateVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelAccelerateVersionsResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeModelService")
    
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeModelServiceCallInfo")
    
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeModelServiceGroup")
    
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeModelServiceGroups")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelServiceGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelServiceGroupsResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeModelServiceHotUpdated")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeModelServiceHotUpdated require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeModelServiceHotUpdatedResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeNotebook")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNotebook require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNotebookResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeNotebooks")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNotebooks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNotebooksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePlatformImagesRequest() (request *DescribePlatformImagesRequest) {
    request = &DescribePlatformImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "DescribePlatformImages")
    
    
    return
}

func NewDescribePlatformImagesResponse() (response *DescribePlatformImagesResponse) {
    response = &DescribePlatformImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePlatformImages
// 查询平台镜像信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"
func (c *Client) DescribePlatformImages(request *DescribePlatformImagesRequest) (response *DescribePlatformImagesResponse, err error) {
    return c.DescribePlatformImagesWithContext(context.Background(), request)
}

// DescribePlatformImages
// 查询平台镜像信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_NOTALLOW = "FailedOperation.NotAllow"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE_BADNAME = "InvalidParameterValue.BadName"
func (c *Client) DescribePlatformImagesWithContext(ctx context.Context, request *DescribePlatformImagesRequest) (response *DescribePlatformImagesResponse, err error) {
    if request == nil {
        request = NewDescribePlatformImagesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribePlatformImages")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePlatformImages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePlatformImagesResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeTrainingModelVersion")
    
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeTrainingModelVersions")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrainingModelVersions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrainingModelVersionsResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeTrainingTask")
    
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeTrainingTaskPods")
    
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
//  INTERNALERROR_QUERYBINDINGTAGSFAILED = "InternalError.QueryBindingTagsFailed"
//  INTERNALERROR_QUERYPRICEFAILED = "InternalError.QueryPriceFailed"
//  INTERNALERROR_QUERYRESOURCEGROUPNAMESFAILED = "InternalError.QueryResourceGroupNamesFailed"
//  INTERNALERROR_QUERYRESOURCESPECFAILED = "InternalError.QueryResourceSpecFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"
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
//  INTERNALERROR_QUERYBINDINGTAGSFAILED = "InternalError.QueryBindingTagsFailed"
//  INTERNALERROR_QUERYPRICEFAILED = "InternalError.QueryPriceFailed"
//  INTERNALERROR_QUERYRESOURCEGROUPNAMESFAILED = "InternalError.QueryResourceGroupNamesFailed"
//  INTERNALERROR_QUERYRESOURCESPECFAILED = "InternalError.QueryResourceSpecFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"
//  INVALIDPARAMETERVALUE_PAGELIMITEXCEEDED = "InvalidParameterValue.PageLimitExceeded"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTrainingTasksWithContext(ctx context.Context, request *DescribeTrainingTasksRequest) (response *DescribeTrainingTasksResponse, err error) {
    if request == nil {
        request = NewDescribeTrainingTasksRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "DescribeTrainingTasks")
    
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "ModifyModelService")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyModelService require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyModelServiceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModelServiceAuthTokenRequest() (request *ModifyModelServiceAuthTokenRequest) {
    request = &ModifyModelServiceAuthTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "ModifyModelServiceAuthToken")
    
    
    return
}

func NewModifyModelServiceAuthTokenResponse() (response *ModifyModelServiceAuthTokenResponse) {
    response = &ModifyModelServiceAuthTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyModelServiceAuthToken
// 修改一个 AuthToken
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyModelServiceAuthToken(request *ModifyModelServiceAuthTokenRequest) (response *ModifyModelServiceAuthTokenResponse, err error) {
    return c.ModifyModelServiceAuthTokenWithContext(context.Background(), request)
}

// ModifyModelServiceAuthToken
// 修改一个 AuthToken
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyModelServiceAuthTokenWithContext(ctx context.Context, request *ModifyModelServiceAuthTokenRequest) (response *ModifyModelServiceAuthTokenResponse, err error) {
    if request == nil {
        request = NewModifyModelServiceAuthTokenRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "ModifyModelServiceAuthToken")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyModelServiceAuthToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyModelServiceAuthTokenResponse()
    err = c.Send(request, response)
    return
}

func NewModifyModelServiceAuthorizationRequest() (request *ModifyModelServiceAuthorizationRequest) {
    request = &ModifyModelServiceAuthorizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tione", APIVersion, "ModifyModelServiceAuthorization")
    
    
    return
}

func NewModifyModelServiceAuthorizationResponse() (response *ModifyModelServiceAuthorizationResponse) {
    response = &ModifyModelServiceAuthorizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyModelServiceAuthorization
// 修改服务鉴权配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyModelServiceAuthorization(request *ModifyModelServiceAuthorizationRequest) (response *ModifyModelServiceAuthorizationResponse, err error) {
    return c.ModifyModelServiceAuthorizationWithContext(context.Background(), request)
}

// ModifyModelServiceAuthorization
// 修改服务鉴权配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EXECDATABASEFAIL = "FailedOperation.ExecDatabaseFail"
//  FAILEDOPERATION_QUERYDATABASEFAIL = "FailedOperation.QueryDatabaseFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyModelServiceAuthorizationWithContext(ctx context.Context, request *ModifyModelServiceAuthorizationRequest) (response *ModifyModelServiceAuthorizationResponse, err error) {
    if request == nil {
        request = NewModifyModelServiceAuthorizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "ModifyModelServiceAuthorization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyModelServiceAuthorization require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyModelServiceAuthorizationResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "ModifyNotebookTags")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyNotebookTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyNotebookTagsResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "PushTrainingMetrics")
    
    if c.GetCredential() == nil {
        return nil, errors.New("PushTrainingMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewPushTrainingMetricsResponse()
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
//  INTERNALERROR_DESCRIBECUSTOMSECRETFAILED = "InternalError.DescribeCustomSecretFailed"
//  INTERNALERROR_DESCRIBESUBUINDEPENDENCYACCESSIBILITYFAILED = "InternalError.DescribeSubuinDependencyAccessibilityFailed"
//  INTERNALERROR_MODIFYCUSTOMSECRETFAILED = "InternalError.ModifyCustomSecretFailed"
//  INTERNALERROR_QUERYRESOURCEGROUPFAILED = "InternalError.QueryResourceGroupFailed"
//  INTERNALERROR_QUERYSUBNETINFOFAILED = "InternalError.QuerySubnetInfoFailed"
//  INTERNALERROR_QUERYUSERTMPCREDENTIALFAILED = "InternalError.QueryUserTMPCredentialFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTEDMULTILOCALDISK = "InvalidParameter.UnsupportedMultiLocalDisk"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CCNHASNOATTACHEDVPC = "InvalidParameterValue.CcnHasNoAttachedVpc"
//  INVALIDPARAMETERVALUE_DATACONFIGNUMLIMITEXCEEDED = "InvalidParameterValue.DataConfigNumLimitExceeded"
//  INVALIDPARAMETERVALUE_GETCFSMOUNTIPFAILED = "InvalidParameterValue.GetCFSMountIPFailed"
//  INVALIDPARAMETERVALUE_IMAGEADDRESSILLEGAL = "InvalidParameterValue.ImageAddressIllegal"
//  INVALIDPARAMETERVALUE_INVALIDRESOURCESPEC = "InvalidParameterValue.InvalidResourceSpec"
//  INVALIDPARAMETERVALUE_NOTALLOWEDTOCREATEFREEVOLUMENOTEBOOKWITHBAREMETALRESOURCEGROUP = "InvalidParameterValue.NotAllowedToCreateFreeVolumeNotebookWithBareMetalResourceGroup"
//  INVALIDPARAMETERVALUE_NOTALLOWEDTOCREATEFREEVOLUMENOTEBOOKWITHSWRESOURCEGROUP = "InvalidParameterValue.NotAllowedToCreateFreeVolumeNotebookWithSWResourceGroup"
//  INVALIDPARAMETERVALUE_NOTALLOWEDTOCREATENOTEBOOKWITHBAREMETALRESOURCEGROUP = "InvalidParameterValue.NotAllowedToCreateNotebookWithBareMetalResourceGroup"
//  INVALIDPARAMETERVALUE_PATHILLEGAL = "InvalidParameterValue.PathIllegal"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDMULTICBSSTORAGE = "InvalidParameterValue.UnsupportedMultiCBSStorage"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_BILLINGSTATUSNOTSUPPORTWHITELISTGPUTYPE = "OperationDenied.BillingStatusNotSupportWhitelistGPUType"
//  OPERATIONDENIED_BILLINGSTATUSRESOURCEINSUFFICIENT = "OperationDenied.BillingStatusResourceInsufficient"
//  OPERATIONDENIED_IPILLEGAL = "OperationDenied.IpIllegal"
//  OPERATIONDENIED_MIYINGBALANCEINSUFFICIENT = "OperationDenied.MIYINGBalanceInsufficient"
//  OPERATIONDENIED_NOTALLOW = "OperationDenied.NotAllow"
//  OPERATIONDENIED_NOTSUPPORTCUSTOMRESOURCE = "OperationDenied.NotSupportCustomResource"
//  OPERATIONDENIED_RESOURCEGROUPINSUFFICIENT = "OperationDenied.ResourceGroupInsufficient"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  RESOURCEINSUFFICIENT_SSHPORTISCONSUMEDUP = "ResourceInsufficient.SSHPortIsConsumedUp"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCEGROUPNOTFOUND = "ResourceNotFound.ResourceGroupNotFound"
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
//  INTERNALERROR_DESCRIBECUSTOMSECRETFAILED = "InternalError.DescribeCustomSecretFailed"
//  INTERNALERROR_DESCRIBESUBUINDEPENDENCYACCESSIBILITYFAILED = "InternalError.DescribeSubuinDependencyAccessibilityFailed"
//  INTERNALERROR_MODIFYCUSTOMSECRETFAILED = "InternalError.ModifyCustomSecretFailed"
//  INTERNALERROR_QUERYRESOURCEGROUPFAILED = "InternalError.QueryResourceGroupFailed"
//  INTERNALERROR_QUERYSUBNETINFOFAILED = "InternalError.QuerySubnetInfoFailed"
//  INTERNALERROR_QUERYUSERTMPCREDENTIALFAILED = "InternalError.QueryUserTMPCredentialFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNSUPPORTEDMULTILOCALDISK = "InvalidParameter.UnsupportedMultiLocalDisk"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_CCNHASNOATTACHEDVPC = "InvalidParameterValue.CcnHasNoAttachedVpc"
//  INVALIDPARAMETERVALUE_DATACONFIGNUMLIMITEXCEEDED = "InvalidParameterValue.DataConfigNumLimitExceeded"
//  INVALIDPARAMETERVALUE_GETCFSMOUNTIPFAILED = "InvalidParameterValue.GetCFSMountIPFailed"
//  INVALIDPARAMETERVALUE_IMAGEADDRESSILLEGAL = "InvalidParameterValue.ImageAddressIllegal"
//  INVALIDPARAMETERVALUE_INVALIDRESOURCESPEC = "InvalidParameterValue.InvalidResourceSpec"
//  INVALIDPARAMETERVALUE_NOTALLOWEDTOCREATEFREEVOLUMENOTEBOOKWITHBAREMETALRESOURCEGROUP = "InvalidParameterValue.NotAllowedToCreateFreeVolumeNotebookWithBareMetalResourceGroup"
//  INVALIDPARAMETERVALUE_NOTALLOWEDTOCREATEFREEVOLUMENOTEBOOKWITHSWRESOURCEGROUP = "InvalidParameterValue.NotAllowedToCreateFreeVolumeNotebookWithSWResourceGroup"
//  INVALIDPARAMETERVALUE_NOTALLOWEDTOCREATENOTEBOOKWITHBAREMETALRESOURCEGROUP = "InvalidParameterValue.NotAllowedToCreateNotebookWithBareMetalResourceGroup"
//  INVALIDPARAMETERVALUE_PATHILLEGAL = "InvalidParameterValue.PathIllegal"
//  INVALIDPARAMETERVALUE_UNSUPPORTEDMULTICBSSTORAGE = "InvalidParameterValue.UnsupportedMultiCBSStorage"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_BILLINGSTATUSNOTSUPPORTWHITELISTGPUTYPE = "OperationDenied.BillingStatusNotSupportWhitelistGPUType"
//  OPERATIONDENIED_BILLINGSTATUSRESOURCEINSUFFICIENT = "OperationDenied.BillingStatusResourceInsufficient"
//  OPERATIONDENIED_IPILLEGAL = "OperationDenied.IpIllegal"
//  OPERATIONDENIED_MIYINGBALANCEINSUFFICIENT = "OperationDenied.MIYINGBalanceInsufficient"
//  OPERATIONDENIED_NOTALLOW = "OperationDenied.NotAllow"
//  OPERATIONDENIED_NOTSUPPORTCUSTOMRESOURCE = "OperationDenied.NotSupportCustomResource"
//  OPERATIONDENIED_RESOURCEGROUPINSUFFICIENT = "OperationDenied.ResourceGroupInsufficient"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  RESOURCEINSUFFICIENT_SSHPORTISCONSUMEDUP = "ResourceInsufficient.SSHPortIsConsumedUp"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_RESOURCEGROUPNOTFOUND = "ResourceNotFound.ResourceGroupNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) StartNotebookWithContext(ctx context.Context, request *StartNotebookRequest) (response *StartNotebookResponse, err error) {
    if request == nil {
        request = NewStartNotebookRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "StartNotebook")
    
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
//  INTERNALERROR_CFSNOTFOUND = "InternalError.CFSNotFound"
//  INTERNALERROR_CREATEJOBINSTANCEFAILED = "InternalError.CreateJobInstanceFailed"
//  INTERNALERROR_CREATETCRINSTANCETOKENFAILED = "InternalError.CreateTcrInstanceTokenFailed"
//  INTERNALERROR_GETCFSFILESYSTEMSFAILED = "InternalError.GetCFSFileSystemsFailed"
//  INTERNALERROR_GETCFSMOUNTINFOFAILED = "InternalError.GetCFSMountInfoFailed"
//  INTERNALERROR_NOTALLOW = "InternalError.NotAllow"
//  INTERNALERROR_QUERYHDFSINFOFAILED = "InternalError.QueryHDFSInfoFailed"
//  INTERNALERROR_QUERYRESOURCEGROUPFAILED = "InternalError.QueryResourceGroupFailed"
//  INTERNALERROR_QUERYRESOURCESPECFAILED = "InternalError.QueryResourceSpecFailed"
//  INTERNALERROR_QUERYSUBNETINFOFAILED = "InternalError.QuerySubnetInfoFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_COSPATHNOTEXIST = "InvalidParameterValue.CosPathNotExist"
//  INVALIDPARAMETERVALUE_DATASETNOTEXIST = "InvalidParameterValue.DatasetNotExist"
//  INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"
//  INVALIDPARAMETERVALUE_GETCFSMOUNTIPFAILED = "InvalidParameterValue.GetCFSMountIPFailed"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_BILLINGEXCEPTION = "OperationDenied.BillingException"
//  OPERATIONDENIED_BILLINGSTATUSRESOURCEINSUFFICIENT = "OperationDenied.BillingStatusResourceInsufficient"
//  OPERATIONDENIED_MIYINGBALANCEINSUFFICIENT = "OperationDenied.MIYINGBalanceInsufficient"
//  OPERATIONDENIED_NOTALLOW = "OperationDenied.NotAllow"
//  OPERATIONDENIED_RESOURCEGROUPINSUFFICIENT = "OperationDenied.ResourceGroupInsufficient"
//  OPERATIONDENIED_TAIJIAPPLICATIONGROUPINSUFFICIENT = "OperationDenied.TAIJIApplicationGroupInsufficient"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CFSNOTFOUND = "ResourceNotFound.CfsNotFound"
//  RESOURCENOTFOUND_VPCNOTFOUND = "ResourceNotFound.VPCNotFound"
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
//  INTERNALERROR_CFSNOTFOUND = "InternalError.CFSNotFound"
//  INTERNALERROR_CREATEJOBINSTANCEFAILED = "InternalError.CreateJobInstanceFailed"
//  INTERNALERROR_CREATETCRINSTANCETOKENFAILED = "InternalError.CreateTcrInstanceTokenFailed"
//  INTERNALERROR_GETCFSFILESYSTEMSFAILED = "InternalError.GetCFSFileSystemsFailed"
//  INTERNALERROR_GETCFSMOUNTINFOFAILED = "InternalError.GetCFSMountInfoFailed"
//  INTERNALERROR_NOTALLOW = "InternalError.NotAllow"
//  INTERNALERROR_QUERYHDFSINFOFAILED = "InternalError.QueryHDFSInfoFailed"
//  INTERNALERROR_QUERYRESOURCEGROUPFAILED = "InternalError.QueryResourceGroupFailed"
//  INTERNALERROR_QUERYRESOURCESPECFAILED = "InternalError.QueryResourceSpecFailed"
//  INTERNALERROR_QUERYSUBNETINFOFAILED = "InternalError.QuerySubnetInfoFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_COSPATHNOTEXIST = "InvalidParameterValue.CosPathNotExist"
//  INVALIDPARAMETERVALUE_DATASETNOTEXIST = "InvalidParameterValue.DatasetNotExist"
//  INVALIDPARAMETERVALUE_FRAMEWORKVERSIONNOTSUPPORT = "InvalidParameterValue.FrameworkVersionNotSupport"
//  INVALIDPARAMETERVALUE_GETCFSMOUNTIPFAILED = "InvalidParameterValue.GetCFSMountIPFailed"
//  OPERATIONDENIED_BALANCEINSUFFICIENT = "OperationDenied.BalanceInsufficient"
//  OPERATIONDENIED_BILLINGEXCEPTION = "OperationDenied.BillingException"
//  OPERATIONDENIED_BILLINGSTATUSRESOURCEINSUFFICIENT = "OperationDenied.BillingStatusResourceInsufficient"
//  OPERATIONDENIED_MIYINGBALANCEINSUFFICIENT = "OperationDenied.MIYINGBalanceInsufficient"
//  OPERATIONDENIED_NOTALLOW = "OperationDenied.NotAllow"
//  OPERATIONDENIED_RESOURCEGROUPINSUFFICIENT = "OperationDenied.ResourceGroupInsufficient"
//  OPERATIONDENIED_TAIJIAPPLICATIONGROUPINSUFFICIENT = "OperationDenied.TAIJIApplicationGroupInsufficient"
//  OPERATIONDENIED_WHITELISTQUOTAEXCEED = "OperationDenied.WhitelistQuotaExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CFSNOTFOUND = "ResourceNotFound.CfsNotFound"
//  RESOURCENOTFOUND_VPCNOTFOUND = "ResourceNotFound.VPCNotFound"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) StartTrainingTaskWithContext(ctx context.Context, request *StartTrainingTaskRequest) (response *StartTrainingTaskResponse, err error) {
    if request == nil {
        request = NewStartTrainingTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "StartTrainingTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartTrainingTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartTrainingTaskResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "StopModelAccelerateTask")
    
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "StopNotebook")
    
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
//  INTERNALERROR_NOPERMISSION = "InternalError.NoPermission"
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
//  INTERNALERROR_NOPERMISSION = "InternalError.NoPermission"
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
    c.InitBaseRequest(&request.BaseRequest, "tione", APIVersion, "StopTrainingTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopTrainingTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopTrainingTaskResponse()
    err = c.Send(request, response)
    return
}
