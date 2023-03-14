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

package v20180416

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-16"

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


func NewCreateIndexRequest() (request *CreateIndexRequest) {
    request = &CreateIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "CreateIndex")
    
    
    return
}

func NewCreateIndexResponse() (response *CreateIndexResponse) {
    response = &CreateIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateIndex
// 创建索引
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIndex(request *CreateIndexRequest) (response *CreateIndexResponse, err error) {
    return c.CreateIndexWithContext(context.Background(), request)
}

// CreateIndex
// 创建索引
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateIndexWithContext(ctx context.Context, request *CreateIndexRequest) (response *CreateIndexResponse, err error) {
    if request == nil {
        request = NewCreateIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateIndexResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceRequest() (request *CreateInstanceRequest) {
    request = &CreateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "CreateInstance")
    
    
    return
}

func NewCreateInstanceResponse() (response *CreateInstanceResponse) {
    response = &CreateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstance
// 创建指定规格的ES集群实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
func (c *Client) CreateInstance(request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    return c.CreateInstanceWithContext(context.Background(), request)
}

// CreateInstance
// 创建指定规格的ES集群实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
func (c *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest) (response *CreateInstanceResponse, err error) {
    if request == nil {
        request = NewCreateInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLogstashInstanceRequest() (request *CreateLogstashInstanceRequest) {
    request = &CreateLogstashInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "CreateLogstashInstance")
    
    
    return
}

func NewCreateLogstashInstanceResponse() (response *CreateLogstashInstanceResponse) {
    response = &CreateLogstashInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLogstashInstance
// 用于创建Logstash实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
func (c *Client) CreateLogstashInstance(request *CreateLogstashInstanceRequest) (response *CreateLogstashInstanceResponse, err error) {
    return c.CreateLogstashInstanceWithContext(context.Background(), request)
}

// CreateLogstashInstance
// 用于创建Logstash实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
func (c *Client) CreateLogstashInstanceWithContext(ctx context.Context, request *CreateLogstashInstanceRequest) (response *CreateLogstashInstanceResponse, err error) {
    if request == nil {
        request = NewCreateLogstashInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLogstashInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLogstashInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteIndexRequest() (request *DeleteIndexRequest) {
    request = &DeleteIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DeleteIndex")
    
    
    return
}

func NewDeleteIndexResponse() (response *DeleteIndexResponse) {
    response = &DeleteIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteIndex
// 删除索引
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteIndex(request *DeleteIndexRequest) (response *DeleteIndexResponse, err error) {
    return c.DeleteIndexWithContext(context.Background(), request)
}

// DeleteIndex
// 删除索引
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteIndexWithContext(ctx context.Context, request *DeleteIndexRequest) (response *DeleteIndexResponse, err error) {
    if request == nil {
        request = NewDeleteIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteIndexResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
    request = &DeleteInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DeleteInstance")
    
    
    return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
    response = &DeleteInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteInstance
// 销毁集群实例 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    return c.DeleteInstanceWithContext(context.Background(), request)
}

// DeleteInstance
// 销毁集群实例 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLogstashInstanceRequest() (request *DeleteLogstashInstanceRequest) {
    request = &DeleteLogstashInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DeleteLogstashInstance")
    
    
    return
}

func NewDeleteLogstashInstanceResponse() (response *DeleteLogstashInstanceResponse) {
    response = &DeleteLogstashInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLogstashInstance
// 用于删除Logstash实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteLogstashInstance(request *DeleteLogstashInstanceRequest) (response *DeleteLogstashInstanceResponse, err error) {
    return c.DeleteLogstashInstanceWithContext(context.Background(), request)
}

// DeleteLogstashInstance
// 用于删除Logstash实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteLogstashInstanceWithContext(ctx context.Context, request *DeleteLogstashInstanceRequest) (response *DeleteLogstashInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteLogstashInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLogstashInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLogstashInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteLogstashPipelinesRequest() (request *DeleteLogstashPipelinesRequest) {
    request = &DeleteLogstashPipelinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DeleteLogstashPipelines")
    
    
    return
}

func NewDeleteLogstashPipelinesResponse() (response *DeleteLogstashPipelinesResponse) {
    response = &DeleteLogstashPipelinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteLogstashPipelines
// 用于批量删除Logstash管道
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteLogstashPipelines(request *DeleteLogstashPipelinesRequest) (response *DeleteLogstashPipelinesResponse, err error) {
    return c.DeleteLogstashPipelinesWithContext(context.Background(), request)
}

// DeleteLogstashPipelines
// 用于批量删除Logstash管道
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DeleteLogstashPipelinesWithContext(ctx context.Context, request *DeleteLogstashPipelinesRequest) (response *DeleteLogstashPipelinesResponse, err error) {
    if request == nil {
        request = NewDeleteLogstashPipelinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteLogstashPipelines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteLogstashPipelinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIndexListRequest() (request *DescribeIndexListRequest) {
    request = &DescribeIndexListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeIndexList")
    
    
    return
}

func NewDescribeIndexListResponse() (response *DescribeIndexListResponse) {
    response = &DescribeIndexListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIndexList
// 获取索引列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIndexList(request *DescribeIndexListRequest) (response *DescribeIndexListResponse, err error) {
    return c.DescribeIndexListWithContext(context.Background(), request)
}

// DescribeIndexList
// 获取索引列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIndexListWithContext(ctx context.Context, request *DescribeIndexListRequest) (response *DescribeIndexListResponse, err error) {
    if request == nil {
        request = NewDescribeIndexListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIndexList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIndexListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIndexMetaRequest() (request *DescribeIndexMetaRequest) {
    request = &DescribeIndexMetaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeIndexMeta")
    
    
    return
}

func NewDescribeIndexMetaResponse() (response *DescribeIndexMetaResponse) {
    response = &DescribeIndexMetaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIndexMeta
// 获取索引元数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIndexMeta(request *DescribeIndexMetaRequest) (response *DescribeIndexMetaResponse, err error) {
    return c.DescribeIndexMetaWithContext(context.Background(), request)
}

// DescribeIndexMeta
// 获取索引元数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeIndexMetaWithContext(ctx context.Context, request *DescribeIndexMetaRequest) (response *DescribeIndexMetaResponse, err error) {
    if request == nil {
        request = NewDescribeIndexMetaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIndexMeta require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIndexMetaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceLogsRequest() (request *DescribeInstanceLogsRequest) {
    request = &DescribeInstanceLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeInstanceLogs")
    
    
    return
}

func NewDescribeInstanceLogsResponse() (response *DescribeInstanceLogsResponse) {
    response = &DescribeInstanceLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceLogs
// 查询用户该地域下符合条件的ES集群的日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeInstanceLogs(request *DescribeInstanceLogsRequest) (response *DescribeInstanceLogsResponse, err error) {
    return c.DescribeInstanceLogsWithContext(context.Background(), request)
}

// DescribeInstanceLogs
// 查询用户该地域下符合条件的ES集群的日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeInstanceLogsWithContext(ctx context.Context, request *DescribeInstanceLogsRequest) (response *DescribeInstanceLogsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceOperationsRequest() (request *DescribeInstanceOperationsRequest) {
    request = &DescribeInstanceOperationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeInstanceOperations")
    
    
    return
}

func NewDescribeInstanceOperationsResponse() (response *DescribeInstanceOperationsResponse) {
    response = &DescribeInstanceOperationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceOperations
// 查询实例指定条件下的操作记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeInstanceOperations(request *DescribeInstanceOperationsRequest) (response *DescribeInstanceOperationsResponse, err error) {
    return c.DescribeInstanceOperationsWithContext(context.Background(), request)
}

// DescribeInstanceOperations
// 查询实例指定条件下的操作记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeInstanceOperationsWithContext(ctx context.Context, request *DescribeInstanceOperationsRequest) (response *DescribeInstanceOperationsResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceOperationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceOperations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceOperationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstances
// 查询用户该地域下符合条件的所有实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// 查询用户该地域下符合条件的所有实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogstashInstanceLogsRequest() (request *DescribeLogstashInstanceLogsRequest) {
    request = &DescribeLogstashInstanceLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeLogstashInstanceLogs")
    
    
    return
}

func NewDescribeLogstashInstanceLogsResponse() (response *DescribeLogstashInstanceLogsResponse) {
    response = &DescribeLogstashInstanceLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogstashInstanceLogs
// 查询用户该地域下符合条件的Logstash实例的日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLogstashInstanceLogs(request *DescribeLogstashInstanceLogsRequest) (response *DescribeLogstashInstanceLogsResponse, err error) {
    return c.DescribeLogstashInstanceLogsWithContext(context.Background(), request)
}

// DescribeLogstashInstanceLogs
// 查询用户该地域下符合条件的Logstash实例的日志
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLogstashInstanceLogsWithContext(ctx context.Context, request *DescribeLogstashInstanceLogsRequest) (response *DescribeLogstashInstanceLogsResponse, err error) {
    if request == nil {
        request = NewDescribeLogstashInstanceLogsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogstashInstanceLogs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogstashInstanceLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogstashInstanceOperationsRequest() (request *DescribeLogstashInstanceOperationsRequest) {
    request = &DescribeLogstashInstanceOperationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeLogstashInstanceOperations")
    
    
    return
}

func NewDescribeLogstashInstanceOperationsResponse() (response *DescribeLogstashInstanceOperationsResponse) {
    response = &DescribeLogstashInstanceOperationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogstashInstanceOperations
// 查询实例指定条件下的操作记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeLogstashInstanceOperations(request *DescribeLogstashInstanceOperationsRequest) (response *DescribeLogstashInstanceOperationsResponse, err error) {
    return c.DescribeLogstashInstanceOperationsWithContext(context.Background(), request)
}

// DescribeLogstashInstanceOperations
// 查询实例指定条件下的操作记录
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeLogstashInstanceOperationsWithContext(ctx context.Context, request *DescribeLogstashInstanceOperationsRequest) (response *DescribeLogstashInstanceOperationsResponse, err error) {
    if request == nil {
        request = NewDescribeLogstashInstanceOperationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogstashInstanceOperations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogstashInstanceOperationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogstashInstancesRequest() (request *DescribeLogstashInstancesRequest) {
    request = &DescribeLogstashInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeLogstashInstances")
    
    
    return
}

func NewDescribeLogstashInstancesResponse() (response *DescribeLogstashInstancesResponse) {
    response = &DescribeLogstashInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogstashInstances
// 查询用户该地域下符合条件的所有Logstash实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLogstashInstances(request *DescribeLogstashInstancesRequest) (response *DescribeLogstashInstancesResponse, err error) {
    return c.DescribeLogstashInstancesWithContext(context.Background(), request)
}

// DescribeLogstashInstances
// 查询用户该地域下符合条件的所有Logstash实例
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLogstashInstancesWithContext(ctx context.Context, request *DescribeLogstashInstancesRequest) (response *DescribeLogstashInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeLogstashInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogstashInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogstashInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLogstashPipelinesRequest() (request *DescribeLogstashPipelinesRequest) {
    request = &DescribeLogstashPipelinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeLogstashPipelines")
    
    
    return
}

func NewDescribeLogstashPipelinesResponse() (response *DescribeLogstashPipelinesResponse) {
    response = &DescribeLogstashPipelinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLogstashPipelines
// 用于获取Logstash实例管道列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLogstashPipelines(request *DescribeLogstashPipelinesRequest) (response *DescribeLogstashPipelinesResponse, err error) {
    return c.DescribeLogstashPipelinesWithContext(context.Background(), request)
}

// DescribeLogstashPipelines
// 用于获取Logstash实例管道列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeLogstashPipelinesWithContext(ctx context.Context, request *DescribeLogstashPipelinesRequest) (response *DescribeLogstashPipelinesResponse, err error) {
    if request == nil {
        request = NewDescribeLogstashPipelinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLogstashPipelines require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLogstashPipelinesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeViewsRequest() (request *DescribeViewsRequest) {
    request = &DescribeViewsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DescribeViews")
    
    
    return
}

func NewDescribeViewsResponse() (response *DescribeViewsResponse) {
    response = &DescribeViewsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeViews
// 查询集群各视图数据，包括集群维度、节点维度、Kibana维度
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeViews(request *DescribeViewsRequest) (response *DescribeViewsResponse, err error) {
    return c.DescribeViewsWithContext(context.Background(), request)
}

// DescribeViews
// 查询集群各视图数据，包括集群维度、节点维度、Kibana维度
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) DescribeViewsWithContext(ctx context.Context, request *DescribeViewsRequest) (response *DescribeViewsResponse, err error) {
    if request == nil {
        request = NewDescribeViewsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeViews require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeViewsResponse()
    err = c.Send(request, response)
    return
}

func NewDiagnoseInstanceRequest() (request *DiagnoseInstanceRequest) {
    request = &DiagnoseInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "DiagnoseInstance")
    
    
    return
}

func NewDiagnoseInstanceResponse() (response *DiagnoseInstanceResponse) {
    response = &DiagnoseInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DiagnoseInstance
// 智能运维诊断集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DiagnoseInstance(request *DiagnoseInstanceRequest) (response *DiagnoseInstanceResponse, err error) {
    return c.DiagnoseInstanceWithContext(context.Background(), request)
}

// DiagnoseInstance
// 智能运维诊断集群
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DiagnoseInstanceWithContext(ctx context.Context, request *DiagnoseInstanceRequest) (response *DiagnoseInstanceResponse, err error) {
    if request == nil {
        request = NewDiagnoseInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DiagnoseInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDiagnoseInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewGetRequestTargetNodeTypesRequest() (request *GetRequestTargetNodeTypesRequest) {
    request = &GetRequestTargetNodeTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "GetRequestTargetNodeTypes")
    
    
    return
}

func NewGetRequestTargetNodeTypesResponse() (response *GetRequestTargetNodeTypesResponse) {
    response = &GetRequestTargetNodeTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRequestTargetNodeTypes
// 获取接收客户端请求的节点类型
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetRequestTargetNodeTypes(request *GetRequestTargetNodeTypesRequest) (response *GetRequestTargetNodeTypesResponse, err error) {
    return c.GetRequestTargetNodeTypesWithContext(context.Background(), request)
}

// GetRequestTargetNodeTypes
// 获取接收客户端请求的节点类型
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetRequestTargetNodeTypesWithContext(ctx context.Context, request *GetRequestTargetNodeTypesRequest) (response *GetRequestTargetNodeTypesResponse, err error) {
    if request == nil {
        request = NewGetRequestTargetNodeTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRequestTargetNodeTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRequestTargetNodeTypesResponse()
    err = c.Send(request, response)
    return
}

func NewRestartInstanceRequest() (request *RestartInstanceRequest) {
    request = &RestartInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "RestartInstance")
    
    
    return
}

func NewRestartInstanceResponse() (response *RestartInstanceResponse) {
    response = &RestartInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestartInstance
// 重启ES集群实例(用于系统版本更新等操作) 
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) RestartInstance(request *RestartInstanceRequest) (response *RestartInstanceResponse, err error) {
    return c.RestartInstanceWithContext(context.Background(), request)
}

// RestartInstance
// 重启ES集群实例(用于系统版本更新等操作) 
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) RestartInstanceWithContext(ctx context.Context, request *RestartInstanceRequest) (response *RestartInstanceResponse, err error) {
    if request == nil {
        request = NewRestartInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRestartKibanaRequest() (request *RestartKibanaRequest) {
    request = &RestartKibanaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "RestartKibana")
    
    
    return
}

func NewRestartKibanaResponse() (response *RestartKibanaResponse) {
    response = &RestartKibanaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestartKibana
// 重启Kibana 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) RestartKibana(request *RestartKibanaRequest) (response *RestartKibanaResponse, err error) {
    return c.RestartKibanaWithContext(context.Background(), request)
}

// RestartKibana
// 重启Kibana 
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) RestartKibanaWithContext(ctx context.Context, request *RestartKibanaRequest) (response *RestartKibanaResponse, err error) {
    if request == nil {
        request = NewRestartKibanaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartKibana require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartKibanaResponse()
    err = c.Send(request, response)
    return
}

func NewRestartLogstashInstanceRequest() (request *RestartLogstashInstanceRequest) {
    request = &RestartLogstashInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "RestartLogstashInstance")
    
    
    return
}

func NewRestartLogstashInstanceResponse() (response *RestartLogstashInstanceResponse) {
    response = &RestartLogstashInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestartLogstashInstance
// 用于重启Logstash实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestartLogstashInstance(request *RestartLogstashInstanceRequest) (response *RestartLogstashInstanceResponse, err error) {
    return c.RestartLogstashInstanceWithContext(context.Background(), request)
}

// RestartLogstashInstance
// 用于重启Logstash实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestartLogstashInstanceWithContext(ctx context.Context, request *RestartLogstashInstanceRequest) (response *RestartLogstashInstanceResponse, err error) {
    if request == nil {
        request = NewRestartLogstashInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartLogstashInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartLogstashInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewRestartNodesRequest() (request *RestartNodesRequest) {
    request = &RestartNodesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "RestartNodes")
    
    
    return
}

func NewRestartNodesResponse() (response *RestartNodesResponse) {
    response = &RestartNodesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RestartNodes
// 用于重启集群节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestartNodes(request *RestartNodesRequest) (response *RestartNodesResponse, err error) {
    return c.RestartNodesWithContext(context.Background(), request)
}

// RestartNodes
// 用于重启集群节点
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) RestartNodesWithContext(ctx context.Context, request *RestartNodesRequest) (response *RestartNodesResponse, err error) {
    if request == nil {
        request = NewRestartNodesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RestartNodes require credential")
    }

    request.SetContext(ctx)
    
    response = NewRestartNodesResponse()
    err = c.Send(request, response)
    return
}

func NewSaveAndDeployLogstashPipelineRequest() (request *SaveAndDeployLogstashPipelineRequest) {
    request = &SaveAndDeployLogstashPipelineRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "SaveAndDeployLogstashPipeline")
    
    
    return
}

func NewSaveAndDeployLogstashPipelineResponse() (response *SaveAndDeployLogstashPipelineResponse) {
    response = &SaveAndDeployLogstashPipelineResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SaveAndDeployLogstashPipeline
// 用于下发并且部署管道
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) SaveAndDeployLogstashPipeline(request *SaveAndDeployLogstashPipelineRequest) (response *SaveAndDeployLogstashPipelineResponse, err error) {
    return c.SaveAndDeployLogstashPipelineWithContext(context.Background(), request)
}

// SaveAndDeployLogstashPipeline
// 用于下发并且部署管道
//
// 可能返回的错误码:
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) SaveAndDeployLogstashPipelineWithContext(ctx context.Context, request *SaveAndDeployLogstashPipelineRequest) (response *SaveAndDeployLogstashPipelineResponse, err error) {
    if request == nil {
        request = NewSaveAndDeployLogstashPipelineRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SaveAndDeployLogstashPipeline require credential")
    }

    request.SetContext(ctx)
    
    response = NewSaveAndDeployLogstashPipelineResponse()
    err = c.Send(request, response)
    return
}

func NewStartLogstashPipelinesRequest() (request *StartLogstashPipelinesRequest) {
    request = &StartLogstashPipelinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "StartLogstashPipelines")
    
    
    return
}

func NewStartLogstashPipelinesResponse() (response *StartLogstashPipelinesResponse) {
    response = &StartLogstashPipelinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartLogstashPipelines
// 用于启动Logstash管道
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) StartLogstashPipelines(request *StartLogstashPipelinesRequest) (response *StartLogstashPipelinesResponse, err error) {
    return c.StartLogstashPipelinesWithContext(context.Background(), request)
}

// StartLogstashPipelines
// 用于启动Logstash管道
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) StartLogstashPipelinesWithContext(ctx context.Context, request *StartLogstashPipelinesRequest) (response *StartLogstashPipelinesResponse, err error) {
    if request == nil {
        request = NewStartLogstashPipelinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartLogstashPipelines require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartLogstashPipelinesResponse()
    err = c.Send(request, response)
    return
}

func NewStopLogstashPipelinesRequest() (request *StopLogstashPipelinesRequest) {
    request = &StopLogstashPipelinesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "StopLogstashPipelines")
    
    
    return
}

func NewStopLogstashPipelinesResponse() (response *StopLogstashPipelinesResponse) {
    response = &StopLogstashPipelinesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopLogstashPipelines
// 用于批量停止Logstash管道
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) StopLogstashPipelines(request *StopLogstashPipelinesRequest) (response *StopLogstashPipelinesResponse, err error) {
    return c.StopLogstashPipelinesWithContext(context.Background(), request)
}

// StopLogstashPipelines
// 用于批量停止Logstash管道
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) StopLogstashPipelinesWithContext(ctx context.Context, request *StopLogstashPipelinesRequest) (response *StopLogstashPipelinesResponse, err error) {
    if request == nil {
        request = NewStopLogstashPipelinesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopLogstashPipelines require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopLogstashPipelinesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDiagnoseSettingsRequest() (request *UpdateDiagnoseSettingsRequest) {
    request = &UpdateDiagnoseSettingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdateDiagnoseSettings")
    
    
    return
}

func NewUpdateDiagnoseSettingsResponse() (response *UpdateDiagnoseSettingsResponse) {
    response = &UpdateDiagnoseSettingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDiagnoseSettings
// 更新智能运维配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateDiagnoseSettings(request *UpdateDiagnoseSettingsRequest) (response *UpdateDiagnoseSettingsResponse, err error) {
    return c.UpdateDiagnoseSettingsWithContext(context.Background(), request)
}

// UpdateDiagnoseSettings
// 更新智能运维配置
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateDiagnoseSettingsWithContext(ctx context.Context, request *UpdateDiagnoseSettingsRequest) (response *UpdateDiagnoseSettingsResponse, err error) {
    if request == nil {
        request = NewUpdateDiagnoseSettingsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDiagnoseSettings require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDiagnoseSettingsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateDictionariesRequest() (request *UpdateDictionariesRequest) {
    request = &UpdateDictionariesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdateDictionaries")
    
    
    return
}

func NewUpdateDictionariesResponse() (response *UpdateDictionariesResponse) {
    response = &UpdateDictionariesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateDictionaries
// 更新ES集群词典
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateDictionaries(request *UpdateDictionariesRequest) (response *UpdateDictionariesResponse, err error) {
    return c.UpdateDictionariesWithContext(context.Background(), request)
}

// UpdateDictionaries
// 更新ES集群词典
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) UpdateDictionariesWithContext(ctx context.Context, request *UpdateDictionariesRequest) (response *UpdateDictionariesResponse, err error) {
    if request == nil {
        request = NewUpdateDictionariesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateDictionaries require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateDictionariesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateIndexRequest() (request *UpdateIndexRequest) {
    request = &UpdateIndexRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdateIndex")
    
    
    return
}

func NewUpdateIndexResponse() (response *UpdateIndexResponse) {
    response = &UpdateIndexResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateIndex
// 更新索引
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateIndex(request *UpdateIndexRequest) (response *UpdateIndexResponse, err error) {
    return c.UpdateIndexWithContext(context.Background(), request)
}

// UpdateIndex
// 更新索引
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateIndexWithContext(ctx context.Context, request *UpdateIndexRequest) (response *UpdateIndexResponse, err error) {
    if request == nil {
        request = NewUpdateIndexRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateIndex require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateIndexResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateInstanceRequest() (request *UpdateInstanceRequest) {
    request = &UpdateInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdateInstance")
    
    
    return
}

func NewUpdateInstanceResponse() (response *UpdateInstanceResponse) {
    response = &UpdateInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateInstance
// 对集群进行节点规格变更，修改实例名称，修改配置，重置密码， 添加Kibana黑白名单等操作。参数中InstanceId为必传参数，ForceRestart为选填参数，剩余参数传递组合及含义如下：
//
// - InstanceName：修改实例名称(仅用于标识实例)
//
// - NodeInfoList: 修改节点配置（节点横向扩缩容，纵向扩缩容，增加主节点，增加冷节点等）
//
// - EsConfig：修改集群配置
//
// - Password：修改默认用户elastic的密码
//
// - EsAcl：修改访问控制列表
//
// - CosBackUp: 设置集群COS自动备份信息
//
// 以上参数组合只能传递一种，多传或少传均会导致请求失败
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_UNSUPPORTRESETNODETYPEANDSCALEOUTDISK = "FailedOperation.UnsupportResetNodeTypeAndScaleoutDisk"
//  FAILEDOPERATION_UNSUPPORTRESETSCALEDOWNANDMODIFYDISK = "FailedOperation.UnsupportResetScaledownAndModifyDisk"
//  FAILEDOPERATION_UNSUPPORTREVERSEREGULATIONNODETYPEANDDISK = "FailedOperation.UnsupportReverseRegulationNodeTypeAndDisk"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateInstance(request *UpdateInstanceRequest) (response *UpdateInstanceResponse, err error) {
    return c.UpdateInstanceWithContext(context.Background(), request)
}

// UpdateInstance
// 对集群进行节点规格变更，修改实例名称，修改配置，重置密码， 添加Kibana黑白名单等操作。参数中InstanceId为必传参数，ForceRestart为选填参数，剩余参数传递组合及含义如下：
//
// - InstanceName：修改实例名称(仅用于标识实例)
//
// - NodeInfoList: 修改节点配置（节点横向扩缩容，纵向扩缩容，增加主节点，增加冷节点等）
//
// - EsConfig：修改集群配置
//
// - Password：修改默认用户elastic的密码
//
// - EsAcl：修改访问控制列表
//
// - CosBackUp: 设置集群COS自动备份信息
//
// 以上参数组合只能传递一种，多传或少传均会导致请求失败
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  FAILEDOPERATION_UNSUPPORTRESETNODETYPEANDSCALEOUTDISK = "FailedOperation.UnsupportResetNodeTypeAndScaleoutDisk"
//  FAILEDOPERATION_UNSUPPORTRESETSCALEDOWNANDMODIFYDISK = "FailedOperation.UnsupportResetScaledownAndModifyDisk"
//  FAILEDOPERATION_UNSUPPORTREVERSEREGULATIONNODETYPEANDDISK = "FailedOperation.UnsupportReverseRegulationNodeTypeAndDisk"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateInstanceWithContext(ctx context.Context, request *UpdateInstanceRequest) (response *UpdateInstanceResponse, err error) {
    if request == nil {
        request = NewUpdateInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateJdkRequest() (request *UpdateJdkRequest) {
    request = &UpdateJdkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdateJdk")
    
    
    return
}

func NewUpdateJdkResponse() (response *UpdateJdkResponse) {
    response = &UpdateJdkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateJdk
// 更新实例Jdk配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATENOREPLICATION = "FailedOperation.ErrorClusterStateNoReplication"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateJdk(request *UpdateJdkRequest) (response *UpdateJdkResponse, err error) {
    return c.UpdateJdkWithContext(context.Background(), request)
}

// UpdateJdk
// 更新实例Jdk配置
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATENOREPLICATION = "FailedOperation.ErrorClusterStateNoReplication"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateJdkWithContext(ctx context.Context, request *UpdateJdkRequest) (response *UpdateJdkResponse, err error) {
    if request == nil {
        request = NewUpdateJdkRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateJdk require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateJdkResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateLogstashInstanceRequest() (request *UpdateLogstashInstanceRequest) {
    request = &UpdateLogstashInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdateLogstashInstance")
    
    
    return
}

func NewUpdateLogstashInstanceResponse() (response *UpdateLogstashInstanceResponse) {
    response = &UpdateLogstashInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateLogstashInstance
// 对集群进行节点规格变更，修改实例名称，修改配置，等操作。参数中InstanceId为必传参数，参数传递组合及含义如下：
//
// - InstanceName：修改实例名称(仅用于标识实例)
//
// - NodeNum: 修改实例节点数量（节点横向扩缩容，纵向扩缩容等）
//
// - YMLConfig: 修改实例YML配置
//
// - BindedES：修改绑定的ES集群配置
//
// 以上参数组合只能传递一种，多传或少传均会导致请求失败
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateLogstashInstance(request *UpdateLogstashInstanceRequest) (response *UpdateLogstashInstanceResponse, err error) {
    return c.UpdateLogstashInstanceWithContext(context.Background(), request)
}

// UpdateLogstashInstance
// 对集群进行节点规格变更，修改实例名称，修改配置，等操作。参数中InstanceId为必传参数，参数传递组合及含义如下：
//
// - InstanceName：修改实例名称(仅用于标识实例)
//
// - NodeNum: 修改实例节点数量（节点横向扩缩容，纵向扩缩容等）
//
// - YMLConfig: 修改实例YML配置
//
// - BindedES：修改绑定的ES集群配置
//
// 以上参数组合只能传递一种，多传或少传均会导致请求失败
//
// 可能返回的错误码:
//  FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateLogstashInstanceWithContext(ctx context.Context, request *UpdateLogstashInstanceRequest) (response *UpdateLogstashInstanceResponse, err error) {
    if request == nil {
        request = NewUpdateLogstashInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateLogstashInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateLogstashInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateLogstashPipelineDescRequest() (request *UpdateLogstashPipelineDescRequest) {
    request = &UpdateLogstashPipelineDescRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdateLogstashPipelineDesc")
    
    
    return
}

func NewUpdateLogstashPipelineDescResponse() (response *UpdateLogstashPipelineDescResponse) {
    response = &UpdateLogstashPipelineDescResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateLogstashPipelineDesc
// 用于更新管道描述信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) UpdateLogstashPipelineDesc(request *UpdateLogstashPipelineDescRequest) (response *UpdateLogstashPipelineDescResponse, err error) {
    return c.UpdateLogstashPipelineDescWithContext(context.Background(), request)
}

// UpdateLogstashPipelineDesc
// 用于更新管道描述信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
func (c *Client) UpdateLogstashPipelineDescWithContext(ctx context.Context, request *UpdateLogstashPipelineDescRequest) (response *UpdateLogstashPipelineDescResponse, err error) {
    if request == nil {
        request = NewUpdateLogstashPipelineDescRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateLogstashPipelineDesc require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateLogstashPipelineDescResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePluginsRequest() (request *UpdatePluginsRequest) {
    request = &UpdatePluginsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdatePlugins")
    
    
    return
}

func NewUpdatePluginsResponse() (response *UpdatePluginsResponse) {
    response = &UpdatePluginsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdatePlugins
// 变更插件列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATENOREPLICATION = "FailedOperation.ErrorClusterStateNoReplication"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdatePlugins(request *UpdatePluginsRequest) (response *UpdatePluginsResponse, err error) {
    return c.UpdatePluginsWithContext(context.Background(), request)
}

// UpdatePlugins
// 变更插件列表
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATENOREPLICATION = "FailedOperation.ErrorClusterStateNoReplication"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdatePluginsWithContext(ctx context.Context, request *UpdatePluginsRequest) (response *UpdatePluginsResponse, err error) {
    if request == nil {
        request = NewUpdatePluginsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePlugins require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePluginsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRequestTargetNodeTypesRequest() (request *UpdateRequestTargetNodeTypesRequest) {
    request = &UpdateRequestTargetNodeTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpdateRequestTargetNodeTypes")
    
    
    return
}

func NewUpdateRequestTargetNodeTypesResponse() (response *UpdateRequestTargetNodeTypesResponse) {
    response = &UpdateRequestTargetNodeTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateRequestTargetNodeTypes
// 更新接收客户端请求的节点类型
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateRequestTargetNodeTypes(request *UpdateRequestTargetNodeTypesRequest) (response *UpdateRequestTargetNodeTypesResponse, err error) {
    return c.UpdateRequestTargetNodeTypesWithContext(context.Background(), request)
}

// UpdateRequestTargetNodeTypes
// 更新接收客户端请求的节点类型
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateRequestTargetNodeTypesWithContext(ctx context.Context, request *UpdateRequestTargetNodeTypesRequest) (response *UpdateRequestTargetNodeTypesResponse, err error) {
    if request == nil {
        request = NewUpdateRequestTargetNodeTypesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateRequestTargetNodeTypes require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateRequestTargetNodeTypesResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeInstanceRequest() (request *UpgradeInstanceRequest) {
    request = &UpgradeInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpgradeInstance")
    
    
    return
}

func NewUpgradeInstanceResponse() (response *UpgradeInstanceResponse) {
    response = &UpgradeInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeInstance
// 升级ES集群版本
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpgradeInstance(request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    return c.UpgradeInstanceWithContext(context.Background(), request)
}

// UpgradeInstance
// 升级ES集群版本
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"
//  FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpgradeInstanceWithContext(ctx context.Context, request *UpgradeInstanceRequest) (response *UpgradeInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeLicenseRequest() (request *UpgradeLicenseRequest) {
    request = &UpgradeLicenseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("es", APIVersion, "UpgradeLicense")
    
    
    return
}

func NewUpgradeLicenseResponse() (response *UpgradeLicenseResponse) {
    response = &UpgradeLicenseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpgradeLicense
// 升级ES商业特性
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpgradeLicense(request *UpgradeLicenseRequest) (response *UpgradeLicenseResponse, err error) {
    return c.UpgradeLicenseWithContext(context.Background(), request)
}

// UpgradeLicense
// 升级ES商业特性
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpgradeLicenseWithContext(ctx context.Context, request *UpgradeLicenseRequest) (response *UpgradeLicenseResponse, err error) {
    if request == nil {
        request = NewUpgradeLicenseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpgradeLicense require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpgradeLicenseResponse()
    err = c.Send(request, response)
    return
}
