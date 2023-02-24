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

package v20190819

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-08-19"

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


func NewAuthorizeTokenRequest() (request *AuthorizeTokenRequest) {
    request = &AuthorizeTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "AuthorizeToken")
    
    
    return
}

func NewAuthorizeTokenResponse() (response *AuthorizeTokenResponse) {
    response = &AuthorizeTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AuthorizeToken
// 给实例授权token
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) AuthorizeToken(request *AuthorizeTokenRequest) (response *AuthorizeTokenResponse, err error) {
    return c.AuthorizeTokenWithContext(context.Background(), request)
}

// AuthorizeToken
// 给实例授权token
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) AuthorizeTokenWithContext(ctx context.Context, request *AuthorizeTokenRequest) (response *AuthorizeTokenResponse, err error) {
    if request == nil {
        request = NewAuthorizeTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AuthorizeToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewAuthorizeTokenResponse()
    err = c.Send(request, response)
    return
}

func NewBatchCreateAclRequest() (request *BatchCreateAclRequest) {
    request = &BatchCreateAclRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "BatchCreateAcl")
    
    
    return
}

func NewBatchCreateAclResponse() (response *BatchCreateAclResponse) {
    response = &BatchCreateAclResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchCreateAcl
// 批量添加ACL策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) BatchCreateAcl(request *BatchCreateAclRequest) (response *BatchCreateAclResponse, err error) {
    return c.BatchCreateAclWithContext(context.Background(), request)
}

// BatchCreateAcl
// 批量添加ACL策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) BatchCreateAclWithContext(ctx context.Context, request *BatchCreateAclRequest) (response *BatchCreateAclResponse, err error) {
    if request == nil {
        request = NewBatchCreateAclRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchCreateAcl require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchCreateAclResponse()
    err = c.Send(request, response)
    return
}

func NewBatchModifyGroupOffsetsRequest() (request *BatchModifyGroupOffsetsRequest) {
    request = &BatchModifyGroupOffsetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "BatchModifyGroupOffsets")
    
    
    return
}

func NewBatchModifyGroupOffsetsResponse() (response *BatchModifyGroupOffsetsResponse) {
    response = &BatchModifyGroupOffsetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchModifyGroupOffsets
// 批量修改消费组offset
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
func (c *Client) BatchModifyGroupOffsets(request *BatchModifyGroupOffsetsRequest) (response *BatchModifyGroupOffsetsResponse, err error) {
    return c.BatchModifyGroupOffsetsWithContext(context.Background(), request)
}

// BatchModifyGroupOffsets
// 批量修改消费组offset
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
func (c *Client) BatchModifyGroupOffsetsWithContext(ctx context.Context, request *BatchModifyGroupOffsetsRequest) (response *BatchModifyGroupOffsetsResponse, err error) {
    if request == nil {
        request = NewBatchModifyGroupOffsetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchModifyGroupOffsets require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchModifyGroupOffsetsResponse()
    err = c.Send(request, response)
    return
}

func NewBatchModifyTopicAttributesRequest() (request *BatchModifyTopicAttributesRequest) {
    request = &BatchModifyTopicAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "BatchModifyTopicAttributes")
    
    
    return
}

func NewBatchModifyTopicAttributesResponse() (response *BatchModifyTopicAttributesResponse) {
    response = &BatchModifyTopicAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchModifyTopicAttributes
// 批量设置主题属性
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) BatchModifyTopicAttributes(request *BatchModifyTopicAttributesRequest) (response *BatchModifyTopicAttributesResponse, err error) {
    return c.BatchModifyTopicAttributesWithContext(context.Background(), request)
}

// BatchModifyTopicAttributes
// 批量设置主题属性
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) BatchModifyTopicAttributesWithContext(ctx context.Context, request *BatchModifyTopicAttributesRequest) (response *BatchModifyTopicAttributesResponse, err error) {
    if request == nil {
        request = NewBatchModifyTopicAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchModifyTopicAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchModifyTopicAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewCancelAuthorizationTokenRequest() (request *CancelAuthorizationTokenRequest) {
    request = &CancelAuthorizationTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CancelAuthorizationToken")
    
    
    return
}

func NewCancelAuthorizationTokenResponse() (response *CancelAuthorizationTokenResponse) {
    response = &CancelAuthorizationTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelAuthorizationToken
// 取消授权token
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) CancelAuthorizationToken(request *CancelAuthorizationTokenRequest) (response *CancelAuthorizationTokenResponse, err error) {
    return c.CancelAuthorizationTokenWithContext(context.Background(), request)
}

// CancelAuthorizationToken
// 取消授权token
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) CancelAuthorizationTokenWithContext(ctx context.Context, request *CancelAuthorizationTokenRequest) (response *CancelAuthorizationTokenResponse, err error) {
    if request == nil {
        request = NewCancelAuthorizationTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelAuthorizationToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelAuthorizationTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCheckCdcClusterRequest() (request *CheckCdcClusterRequest) {
    request = &CheckCdcClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CheckCdcCluster")
    
    
    return
}

func NewCheckCdcClusterResponse() (response *CheckCdcClusterResponse) {
    response = &CheckCdcClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckCdcCluster
// 用于查询cdc-ckafka任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) CheckCdcCluster(request *CheckCdcClusterRequest) (response *CheckCdcClusterResponse, err error) {
    return c.CheckCdcClusterWithContext(context.Background(), request)
}

// CheckCdcCluster
// 用于查询cdc-ckafka任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) CheckCdcClusterWithContext(ctx context.Context, request *CheckCdcClusterRequest) (response *CheckCdcClusterResponse, err error) {
    if request == nil {
        request = NewCheckCdcClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckCdcCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckCdcClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAclRequest() (request *CreateAclRequest) {
    request = &CreateAclRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateAcl")
    
    
    return
}

func NewCreateAclResponse() (response *CreateAclResponse) {
    response = &CreateAclResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAcl
// 添加 ACL 策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateAcl(request *CreateAclRequest) (response *CreateAclResponse, err error) {
    return c.CreateAclWithContext(context.Background(), request)
}

// CreateAcl
// 添加 ACL 策略
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateAclWithContext(ctx context.Context, request *CreateAclRequest) (response *CreateAclResponse, err error) {
    if request == nil {
        request = NewCreateAclRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAcl require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAclResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCdcClusterRequest() (request *CreateCdcClusterRequest) {
    request = &CreateCdcClusterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateCdcCluster")
    
    
    return
}

func NewCreateCdcClusterResponse() (response *CreateCdcClusterResponse) {
    response = &CreateCdcClusterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCdcCluster
// 用于cdc的专用ckafka集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateCdcCluster(request *CreateCdcClusterRequest) (response *CreateCdcClusterResponse, err error) {
    return c.CreateCdcClusterWithContext(context.Background(), request)
}

// CreateCdcCluster
// 用于cdc的专用ckafka集群
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateCdcClusterWithContext(ctx context.Context, request *CreateCdcClusterRequest) (response *CreateCdcClusterResponse, err error) {
    if request == nil {
        request = NewCreateCdcClusterRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCdcCluster require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCdcClusterResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConnectResourceRequest() (request *CreateConnectResourceRequest) {
    request = &CreateConnectResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateConnectResource")
    
    
    return
}

func NewCreateConnectResourceResponse() (response *CreateConnectResourceResponse) {
    response = &CreateConnectResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateConnectResource
// 创建Datahub连接源
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
func (c *Client) CreateConnectResource(request *CreateConnectResourceRequest) (response *CreateConnectResourceResponse, err error) {
    return c.CreateConnectResourceWithContext(context.Background(), request)
}

// CreateConnectResource
// 创建Datahub连接源
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
func (c *Client) CreateConnectResourceWithContext(ctx context.Context, request *CreateConnectResourceRequest) (response *CreateConnectResourceResponse, err error) {
    if request == nil {
        request = NewCreateConnectResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConnectResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConnectResourceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateConsumerRequest() (request *CreateConsumerRequest) {
    request = &CreateConsumerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateConsumer")
    
    
    return
}

func NewCreateConsumerResponse() (response *CreateConsumerResponse) {
    response = &CreateConsumerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateConsumer
// 创建消费者组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateConsumer(request *CreateConsumerRequest) (response *CreateConsumerResponse, err error) {
    return c.CreateConsumerWithContext(context.Background(), request)
}

// CreateConsumer
// 创建消费者组
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateConsumerWithContext(ctx context.Context, request *CreateConsumerRequest) (response *CreateConsumerResponse, err error) {
    if request == nil {
        request = NewCreateConsumerRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateConsumer require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateConsumerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDatahubTaskRequest() (request *CreateDatahubTaskRequest) {
    request = &CreateDatahubTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateDatahubTask")
    
    
    return
}

func NewCreateDatahubTaskResponse() (response *CreateDatahubTaskResponse) {
    response = &CreateDatahubTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDatahubTask
// 创建DIP转储任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDatahubTask(request *CreateDatahubTaskRequest) (response *CreateDatahubTaskResponse, err error) {
    return c.CreateDatahubTaskWithContext(context.Background(), request)
}

// CreateDatahubTask
// 创建DIP转储任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDatahubTaskWithContext(ctx context.Context, request *CreateDatahubTaskRequest) (response *CreateDatahubTaskResponse, err error) {
    if request == nil {
        request = NewCreateDatahubTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDatahubTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDatahubTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstancePostRequest() (request *CreateInstancePostRequest) {
    request = &CreateInstancePostRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateInstancePost")
    
    
    return
}

func NewCreateInstancePostResponse() (response *CreateInstancePostResponse) {
    response = &CreateInstancePostResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstancePost
// 创建按量计费实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateInstancePost(request *CreateInstancePostRequest) (response *CreateInstancePostResponse, err error) {
    return c.CreateInstancePostWithContext(context.Background(), request)
}

// CreateInstancePost
// 创建按量计费实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateInstancePostWithContext(ctx context.Context, request *CreateInstancePostRequest) (response *CreateInstancePostResponse, err error) {
    if request == nil {
        request = NewCreateInstancePostRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstancePost require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstancePostResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstancePreRequest() (request *CreateInstancePreRequest) {
    request = &CreateInstancePreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateInstancePre")
    
    
    return
}

func NewCreateInstancePreResponse() (response *CreateInstancePreResponse) {
    response = &CreateInstancePreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateInstancePre
// 创建实例(预付费包年包月)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateInstancePre(request *CreateInstancePreRequest) (response *CreateInstancePreResponse, err error) {
    return c.CreateInstancePreWithContext(context.Background(), request)
}

// CreateInstancePre
// 创建实例(预付费包年包月)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateInstancePreWithContext(ctx context.Context, request *CreateInstancePreRequest) (response *CreateInstancePreResponse, err error) {
    if request == nil {
        request = NewCreateInstancePreRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateInstancePre require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateInstancePreResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePartitionRequest() (request *CreatePartitionRequest) {
    request = &CreatePartitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreatePartition")
    
    
    return
}

func NewCreatePartitionResponse() (response *CreatePartitionResponse) {
    response = &CreatePartitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePartition
// 本接口用于增加主题中的分区
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreatePartition(request *CreatePartitionRequest) (response *CreatePartitionResponse, err error) {
    return c.CreatePartitionWithContext(context.Background(), request)
}

// CreatePartition
// 本接口用于增加主题中的分区
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreatePartitionWithContext(ctx context.Context, request *CreatePartitionRequest) (response *CreatePartitionResponse, err error) {
    if request == nil {
        request = NewCreatePartitionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePartition require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePartitionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRouteRequest() (request *CreateRouteRequest) {
    request = &CreateRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateRoute")
    
    
    return
}

func NewCreateRouteResponse() (response *CreateRouteResponse) {
    response = &CreateRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRoute
// 添加实例路由
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  LIMITEXCEEDED_ROUTEOVERLIMIT = "LimitExceeded.RouteOverLimit"
//  LIMITEXCEEDED_ROUTESASLOVERLIMIT = "LimitExceeded.RouteSASLOverLimit"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateRoute(request *CreateRouteRequest) (response *CreateRouteResponse, err error) {
    return c.CreateRouteWithContext(context.Background(), request)
}

// CreateRoute
// 添加实例路由
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  LIMITEXCEEDED_ROUTEOVERLIMIT = "LimitExceeded.RouteOverLimit"
//  LIMITEXCEEDED_ROUTESASLOVERLIMIT = "LimitExceeded.RouteSASLOverLimit"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateRouteWithContext(ctx context.Context, request *CreateRouteRequest) (response *CreateRouteResponse, err error) {
    if request == nil {
        request = NewCreateRouteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRouteResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTokenRequest() (request *CreateTokenRequest) {
    request = &CreateTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateToken")
    
    
    return
}

func NewCreateTokenResponse() (response *CreateTokenResponse) {
    response = &CreateTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateToken
// 创建最高权限的token
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) CreateToken(request *CreateTokenRequest) (response *CreateTokenResponse, err error) {
    return c.CreateTokenWithContext(context.Background(), request)
}

// CreateToken
// 创建最高权限的token
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) CreateTokenWithContext(ctx context.Context, request *CreateTokenRequest) (response *CreateTokenResponse, err error) {
    if request == nil {
        request = NewCreateTokenRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateToken require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopicRequest() (request *CreateTopicRequest) {
    request = &CreateTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateTopic")
    
    
    return
}

func NewCreateTopicResponse() (response *CreateTopicResponse) {
    response = &CreateTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTopic
// 创建ckafka主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    return c.CreateTopicWithContext(context.Background(), request)
}

// CreateTopic
// 创建ckafka主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateTopicWithContext(ctx context.Context, request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
    if request == nil {
        request = NewCreateTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTopicResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTopicIpWhiteListRequest() (request *CreateTopicIpWhiteListRequest) {
    request = &CreateTopicIpWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateTopicIpWhiteList")
    
    
    return
}

func NewCreateTopicIpWhiteListResponse() (response *CreateTopicIpWhiteListResponse) {
    response = &CreateTopicIpWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTopicIpWhiteList
// 创建主题ip白名单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateTopicIpWhiteList(request *CreateTopicIpWhiteListRequest) (response *CreateTopicIpWhiteListResponse, err error) {
    return c.CreateTopicIpWhiteListWithContext(context.Background(), request)
}

// CreateTopicIpWhiteList
// 创建主题ip白名单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateTopicIpWhiteListWithContext(ctx context.Context, request *CreateTopicIpWhiteListRequest) (response *CreateTopicIpWhiteListResponse, err error) {
    if request == nil {
        request = NewCreateTopicIpWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTopicIpWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTopicIpWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
    request = &CreateUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "CreateUser")
    
    
    return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
    response = &CreateUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUser
// 添加用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
    return c.CreateUserWithContext(context.Background(), request)
}

// CreateUser
// 添加用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest) (response *CreateUserResponse, err error) {
    if request == nil {
        request = NewCreateUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAclRequest() (request *DeleteAclRequest) {
    request = &DeleteAclRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteAcl")
    
    
    return
}

func NewDeleteAclResponse() (response *DeleteAclResponse) {
    response = &DeleteAclResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAcl
// 删除ACL
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteAcl(request *DeleteAclRequest) (response *DeleteAclResponse, err error) {
    return c.DeleteAclWithContext(context.Background(), request)
}

// DeleteAcl
// 删除ACL
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteAclWithContext(ctx context.Context, request *DeleteAclRequest) (response *DeleteAclResponse, err error) {
    if request == nil {
        request = NewDeleteAclRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAcl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAclResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAclRuleRequest() (request *DeleteAclRuleRequest) {
    request = &DeleteAclRuleRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteAclRule")
    
    
    return
}

func NewDeleteAclRuleResponse() (response *DeleteAclRuleResponse) {
    response = &DeleteAclRuleResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAclRule
// 删除ACL规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteAclRule(request *DeleteAclRuleRequest) (response *DeleteAclRuleResponse, err error) {
    return c.DeleteAclRuleWithContext(context.Background(), request)
}

// DeleteAclRule
// 删除ACL规则
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteAclRuleWithContext(ctx context.Context, request *DeleteAclRuleRequest) (response *DeleteAclRuleResponse, err error) {
    if request == nil {
        request = NewDeleteAclRuleRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAclRule require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAclRuleResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteConnectResourceRequest() (request *DeleteConnectResourceRequest) {
    request = &DeleteConnectResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteConnectResource")
    
    
    return
}

func NewDeleteConnectResourceResponse() (response *DeleteConnectResourceResponse) {
    response = &DeleteConnectResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteConnectResource
// 删除Datahub连接源
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
func (c *Client) DeleteConnectResource(request *DeleteConnectResourceRequest) (response *DeleteConnectResourceResponse, err error) {
    return c.DeleteConnectResourceWithContext(context.Background(), request)
}

// DeleteConnectResource
// 删除Datahub连接源
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
func (c *Client) DeleteConnectResourceWithContext(ctx context.Context, request *DeleteConnectResourceRequest) (response *DeleteConnectResourceResponse, err error) {
    if request == nil {
        request = NewDeleteConnectResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteConnectResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteConnectResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDatahubTaskRequest() (request *DeleteDatahubTaskRequest) {
    request = &DeleteDatahubTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteDatahubTask")
    
    
    return
}

func NewDeleteDatahubTaskResponse() (response *DeleteDatahubTaskResponse) {
    response = &DeleteDatahubTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDatahubTask
// 删除Datahub任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
func (c *Client) DeleteDatahubTask(request *DeleteDatahubTaskRequest) (response *DeleteDatahubTaskResponse, err error) {
    return c.DeleteDatahubTaskWithContext(context.Background(), request)
}

// DeleteDatahubTask
// 删除Datahub任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
func (c *Client) DeleteDatahubTaskWithContext(ctx context.Context, request *DeleteDatahubTaskRequest) (response *DeleteDatahubTaskResponse, err error) {
    if request == nil {
        request = NewDeleteDatahubTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDatahubTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDatahubTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDatahubTopicRequest() (request *DeleteDatahubTopicRequest) {
    request = &DeleteDatahubTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteDatahubTopic")
    
    
    return
}

func NewDeleteDatahubTopicResponse() (response *DeleteDatahubTopicResponse) {
    response = &DeleteDatahubTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteDatahubTopic
// 删除Datahub主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteDatahubTopic(request *DeleteDatahubTopicRequest) (response *DeleteDatahubTopicResponse, err error) {
    return c.DeleteDatahubTopicWithContext(context.Background(), request)
}

// DeleteDatahubTopic
// 删除Datahub主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteDatahubTopicWithContext(ctx context.Context, request *DeleteDatahubTopicRequest) (response *DeleteDatahubTopicResponse, err error) {
    if request == nil {
        request = NewDeleteDatahubTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteDatahubTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteDatahubTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
    request = &DeleteGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteGroup")
    
    
    return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
    response = &DeleteGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteGroup
// 删除消费组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    return c.DeleteGroupWithContext(context.Background(), request)
}

// DeleteGroup
// 删除消费组
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteGroupWithContext(ctx context.Context, request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteInstancePreRequest() (request *DeleteInstancePreRequest) {
    request = &DeleteInstancePreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteInstancePre")
    
    
    return
}

func NewDeleteInstancePreResponse() (response *DeleteInstancePreResponse) {
    response = &DeleteInstancePreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteInstancePre
// 删除预付费实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstancePre(request *DeleteInstancePreRequest) (response *DeleteInstancePreResponse, err error) {
    return c.DeleteInstancePreWithContext(context.Background(), request)
}

// DeleteInstancePre
// 删除预付费实例
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteInstancePreWithContext(ctx context.Context, request *DeleteInstancePreRequest) (response *DeleteInstancePreResponse, err error) {
    if request == nil {
        request = NewDeleteInstancePreRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteInstancePre require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteInstancePreResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRouteRequest() (request *DeleteRouteRequest) {
    request = &DeleteRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteRoute")
    
    
    return
}

func NewDeleteRouteResponse() (response *DeleteRouteResponse) {
    response = &DeleteRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRoute
// 删除路由
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteRoute(request *DeleteRouteRequest) (response *DeleteRouteResponse, err error) {
    return c.DeleteRouteWithContext(context.Background(), request)
}

// DeleteRoute
// 删除路由
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteRouteWithContext(ctx context.Context, request *DeleteRouteRequest) (response *DeleteRouteResponse, err error) {
    if request == nil {
        request = NewDeleteRouteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRouteTriggerTimeRequest() (request *DeleteRouteTriggerTimeRequest) {
    request = &DeleteRouteTriggerTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteRouteTriggerTime")
    
    
    return
}

func NewDeleteRouteTriggerTimeResponse() (response *DeleteRouteTriggerTimeResponse) {
    response = &DeleteRouteTriggerTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteRouteTriggerTime
// 修改删除路由延迟触发时间
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRouteTriggerTime(request *DeleteRouteTriggerTimeRequest) (response *DeleteRouteTriggerTimeResponse, err error) {
    return c.DeleteRouteTriggerTimeWithContext(context.Background(), request)
}

// DeleteRouteTriggerTime
// 修改删除路由延迟触发时间
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteRouteTriggerTimeWithContext(ctx context.Context, request *DeleteRouteTriggerTimeRequest) (response *DeleteRouteTriggerTimeResponse, err error) {
    if request == nil {
        request = NewDeleteRouteTriggerTimeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteRouteTriggerTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteRouteTriggerTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicRequest() (request *DeleteTopicRequest) {
    request = &DeleteTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteTopic")
    
    
    return
}

func NewDeleteTopicResponse() (response *DeleteTopicResponse) {
    response = &DeleteTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTopic
// 删除ckafka主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    return c.DeleteTopicWithContext(context.Background(), request)
}

// DeleteTopic
// 删除ckafka主题
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteTopicWithContext(ctx context.Context, request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
    if request == nil {
        request = NewDeleteTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTopicIpWhiteListRequest() (request *DeleteTopicIpWhiteListRequest) {
    request = &DeleteTopicIpWhiteListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteTopicIpWhiteList")
    
    
    return
}

func NewDeleteTopicIpWhiteListResponse() (response *DeleteTopicIpWhiteListResponse) {
    response = &DeleteTopicIpWhiteListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTopicIpWhiteList
// 删除主题IP白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteTopicIpWhiteList(request *DeleteTopicIpWhiteListRequest) (response *DeleteTopicIpWhiteListResponse, err error) {
    return c.DeleteTopicIpWhiteListWithContext(context.Background(), request)
}

// DeleteTopicIpWhiteList
// 删除主题IP白名单
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteTopicIpWhiteListWithContext(ctx context.Context, request *DeleteTopicIpWhiteListRequest) (response *DeleteTopicIpWhiteListResponse, err error) {
    if request == nil {
        request = NewDeleteTopicIpWhiteListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteTopicIpWhiteList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteTopicIpWhiteListResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteUserRequest() (request *DeleteUserRequest) {
    request = &DeleteUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DeleteUser")
    
    
    return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
    response = &DeleteUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteUser
// 删除用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    return c.DeleteUserWithContext(context.Background(), request)
}

// DeleteUser
// 删除用户
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DeleteUserWithContext(ctx context.Context, request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
    if request == nil {
        request = NewDeleteUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeACLRequest() (request *DescribeACLRequest) {
    request = &DescribeACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeACL")
    
    
    return
}

func NewDescribeACLResponse() (response *DescribeACLResponse) {
    response = &DescribeACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeACL
// 枚举ACL
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeACL(request *DescribeACLRequest) (response *DescribeACLResponse, err error) {
    return c.DescribeACLWithContext(context.Background(), request)
}

// DescribeACL
// 枚举ACL
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeACLWithContext(ctx context.Context, request *DescribeACLRequest) (response *DescribeACLResponse, err error) {
    if request == nil {
        request = NewDescribeACLRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeACL require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeACLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppInfoRequest() (request *DescribeAppInfoRequest) {
    request = &DescribeAppInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeAppInfo")
    
    
    return
}

func NewDescribeAppInfoResponse() (response *DescribeAppInfoResponse) {
    response = &DescribeAppInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAppInfo
// 查询用户列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeAppInfo(request *DescribeAppInfoRequest) (response *DescribeAppInfoResponse, err error) {
    return c.DescribeAppInfoWithContext(context.Background(), request)
}

// DescribeAppInfo
// 查询用户列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeAppInfoWithContext(ctx context.Context, request *DescribeAppInfoRequest) (response *DescribeAppInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAppInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAppInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAppInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCkafkaZoneRequest() (request *DescribeCkafkaZoneRequest) {
    request = &DescribeCkafkaZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeCkafkaZone")
    
    
    return
}

func NewDescribeCkafkaZoneResponse() (response *DescribeCkafkaZoneResponse) {
    response = &DescribeCkafkaZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCkafkaZone
// 用于查看ckafka的可用区列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeCkafkaZone(request *DescribeCkafkaZoneRequest) (response *DescribeCkafkaZoneResponse, err error) {
    return c.DescribeCkafkaZoneWithContext(context.Background(), request)
}

// DescribeCkafkaZone
// 用于查看ckafka的可用区列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeCkafkaZoneWithContext(ctx context.Context, request *DescribeCkafkaZoneRequest) (response *DescribeCkafkaZoneResponse, err error) {
    if request == nil {
        request = NewDescribeCkafkaZoneRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCkafkaZone require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCkafkaZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConnectResourceRequest() (request *DescribeConnectResourceRequest) {
    request = &DescribeConnectResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeConnectResource")
    
    
    return
}

func NewDescribeConnectResourceResponse() (response *DescribeConnectResourceResponse) {
    response = &DescribeConnectResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeConnectResource
// 查询Datahub连接源
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
func (c *Client) DescribeConnectResource(request *DescribeConnectResourceRequest) (response *DescribeConnectResourceResponse, err error) {
    return c.DescribeConnectResourceWithContext(context.Background(), request)
}

// DescribeConnectResource
// 查询Datahub连接源
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
func (c *Client) DescribeConnectResourceWithContext(ctx context.Context, request *DescribeConnectResourceRequest) (response *DescribeConnectResourceResponse, err error) {
    if request == nil {
        request = NewDescribeConnectResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConnectResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConnectResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConnectResourcesRequest() (request *DescribeConnectResourcesRequest) {
    request = &DescribeConnectResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeConnectResources")
    
    
    return
}

func NewDescribeConnectResourcesResponse() (response *DescribeConnectResourcesResponse) {
    response = &DescribeConnectResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeConnectResources
// 查询Datahub连接源列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
func (c *Client) DescribeConnectResources(request *DescribeConnectResourcesRequest) (response *DescribeConnectResourcesResponse, err error) {
    return c.DescribeConnectResourcesWithContext(context.Background(), request)
}

// DescribeConnectResources
// 查询Datahub连接源列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
func (c *Client) DescribeConnectResourcesWithContext(ctx context.Context, request *DescribeConnectResourcesRequest) (response *DescribeConnectResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeConnectResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConnectResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConnectResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeConsumerGroupRequest() (request *DescribeConsumerGroupRequest) {
    request = &DescribeConsumerGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeConsumerGroup")
    
    
    return
}

func NewDescribeConsumerGroupResponse() (response *DescribeConsumerGroupResponse) {
    response = &DescribeConsumerGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeConsumerGroup
// 查询消费分组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeConsumerGroup(request *DescribeConsumerGroupRequest) (response *DescribeConsumerGroupResponse, err error) {
    return c.DescribeConsumerGroupWithContext(context.Background(), request)
}

// DescribeConsumerGroup
// 查询消费分组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeConsumerGroupWithContext(ctx context.Context, request *DescribeConsumerGroupRequest) (response *DescribeConsumerGroupResponse, err error) {
    if request == nil {
        request = NewDescribeConsumerGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeConsumerGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeConsumerGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatahubGroupOffsetsRequest() (request *DescribeDatahubGroupOffsetsRequest) {
    request = &DescribeDatahubGroupOffsetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeDatahubGroupOffsets")
    
    
    return
}

func NewDescribeDatahubGroupOffsetsResponse() (response *DescribeDatahubGroupOffsetsResponse) {
    response = &DescribeDatahubGroupOffsetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatahubGroupOffsets
// 获取Datahub消费分组offset
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeDatahubGroupOffsets(request *DescribeDatahubGroupOffsetsRequest) (response *DescribeDatahubGroupOffsetsResponse, err error) {
    return c.DescribeDatahubGroupOffsetsWithContext(context.Background(), request)
}

// DescribeDatahubGroupOffsets
// 获取Datahub消费分组offset
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeDatahubGroupOffsetsWithContext(ctx context.Context, request *DescribeDatahubGroupOffsetsRequest) (response *DescribeDatahubGroupOffsetsResponse, err error) {
    if request == nil {
        request = NewDescribeDatahubGroupOffsetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatahubGroupOffsets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatahubGroupOffsetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatahubTaskRequest() (request *DescribeDatahubTaskRequest) {
    request = &DescribeDatahubTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeDatahubTask")
    
    
    return
}

func NewDescribeDatahubTaskResponse() (response *DescribeDatahubTaskResponse) {
    response = &DescribeDatahubTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatahubTask
// 查询Datahub任务信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDatahubTask(request *DescribeDatahubTaskRequest) (response *DescribeDatahubTaskResponse, err error) {
    return c.DescribeDatahubTaskWithContext(context.Background(), request)
}

// DescribeDatahubTask
// 查询Datahub任务信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDatahubTaskWithContext(ctx context.Context, request *DescribeDatahubTaskRequest) (response *DescribeDatahubTaskResponse, err error) {
    if request == nil {
        request = NewDescribeDatahubTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatahubTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatahubTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatahubTasksRequest() (request *DescribeDatahubTasksRequest) {
    request = &DescribeDatahubTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeDatahubTasks")
    
    
    return
}

func NewDescribeDatahubTasksResponse() (response *DescribeDatahubTasksResponse) {
    response = &DescribeDatahubTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatahubTasks
// 查询Datahub任务列表 
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
func (c *Client) DescribeDatahubTasks(request *DescribeDatahubTasksRequest) (response *DescribeDatahubTasksResponse, err error) {
    return c.DescribeDatahubTasksWithContext(context.Background(), request)
}

// DescribeDatahubTasks
// 查询Datahub任务列表 
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
func (c *Client) DescribeDatahubTasksWithContext(ctx context.Context, request *DescribeDatahubTasksRequest) (response *DescribeDatahubTasksResponse, err error) {
    if request == nil {
        request = NewDescribeDatahubTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatahubTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatahubTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatahubTopicRequest() (request *DescribeDatahubTopicRequest) {
    request = &DescribeDatahubTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeDatahubTopic")
    
    
    return
}

func NewDescribeDatahubTopicResponse() (response *DescribeDatahubTopicResponse) {
    response = &DescribeDatahubTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatahubTopic
// 获取Datahub主题属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeDatahubTopic(request *DescribeDatahubTopicRequest) (response *DescribeDatahubTopicResponse, err error) {
    return c.DescribeDatahubTopicWithContext(context.Background(), request)
}

// DescribeDatahubTopic
// 获取Datahub主题属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeDatahubTopicWithContext(ctx context.Context, request *DescribeDatahubTopicRequest) (response *DescribeDatahubTopicResponse, err error) {
    if request == nil {
        request = NewDescribeDatahubTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatahubTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatahubTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatahubTopicsRequest() (request *DescribeDatahubTopicsRequest) {
    request = &DescribeDatahubTopicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeDatahubTopics")
    
    
    return
}

func NewDescribeDatahubTopicsResponse() (response *DescribeDatahubTopicsResponse) {
    response = &DescribeDatahubTopicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDatahubTopics
// 查询DIP主题列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeDatahubTopics(request *DescribeDatahubTopicsRequest) (response *DescribeDatahubTopicsResponse, err error) {
    return c.DescribeDatahubTopicsWithContext(context.Background(), request)
}

// DescribeDatahubTopics
// 查询DIP主题列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeDatahubTopicsWithContext(ctx context.Context, request *DescribeDatahubTopicsRequest) (response *DescribeDatahubTopicsResponse, err error) {
    if request == nil {
        request = NewDescribeDatahubTopicsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDatahubTopics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDatahubTopicsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupRequest() (request *DescribeGroupRequest) {
    request = &DescribeGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeGroup")
    
    
    return
}

func NewDescribeGroupResponse() (response *DescribeGroupResponse) {
    response = &DescribeGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroup
// 枚举消费分组(精简版)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeGroup(request *DescribeGroupRequest) (response *DescribeGroupResponse, err error) {
    return c.DescribeGroupWithContext(context.Background(), request)
}

// DescribeGroup
// 枚举消费分组(精简版)
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeGroupWithContext(ctx context.Context, request *DescribeGroupRequest) (response *DescribeGroupResponse, err error) {
    if request == nil {
        request = NewDescribeGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupInfoRequest() (request *DescribeGroupInfoRequest) {
    request = &DescribeGroupInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeGroupInfo")
    
    
    return
}

func NewDescribeGroupInfoResponse() (response *DescribeGroupInfoResponse) {
    response = &DescribeGroupInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupInfo
// 获取消费分组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeGroupInfo(request *DescribeGroupInfoRequest) (response *DescribeGroupInfoResponse, err error) {
    return c.DescribeGroupInfoWithContext(context.Background(), request)
}

// DescribeGroupInfo
// 获取消费分组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeGroupInfoWithContext(ctx context.Context, request *DescribeGroupInfoRequest) (response *DescribeGroupInfoResponse, err error) {
    if request == nil {
        request = NewDescribeGroupInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupOffsetsRequest() (request *DescribeGroupOffsetsRequest) {
    request = &DescribeGroupOffsetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeGroupOffsets")
    
    
    return
}

func NewDescribeGroupOffsetsResponse() (response *DescribeGroupOffsetsResponse) {
    response = &DescribeGroupOffsetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroupOffsets
// 获取消费分组offset
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeGroupOffsets(request *DescribeGroupOffsetsRequest) (response *DescribeGroupOffsetsResponse, err error) {
    return c.DescribeGroupOffsetsWithContext(context.Background(), request)
}

// DescribeGroupOffsets
// 获取消费分组offset
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeGroupOffsetsWithContext(ctx context.Context, request *DescribeGroupOffsetsRequest) (response *DescribeGroupOffsetsResponse, err error) {
    if request == nil {
        request = NewDescribeGroupOffsetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroupOffsets require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGroupOffsetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceAttributesRequest() (request *DescribeInstanceAttributesRequest) {
    request = &DescribeInstanceAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeInstanceAttributes")
    
    
    return
}

func NewDescribeInstanceAttributesResponse() (response *DescribeInstanceAttributesResponse) {
    response = &DescribeInstanceAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceAttributes
// 获取实例属性 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeInstanceAttributes(request *DescribeInstanceAttributesRequest) (response *DescribeInstanceAttributesResponse, err error) {
    return c.DescribeInstanceAttributesWithContext(context.Background(), request)
}

// DescribeInstanceAttributes
// 获取实例属性 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeInstanceAttributesWithContext(ctx context.Context, request *DescribeInstanceAttributesRequest) (response *DescribeInstanceAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstanceAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstanceAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeInstances")
    
    
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstances
// 本接口（DescribeInstance）用于在用户账户下获取消息队列 CKafka 实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    return c.DescribeInstancesWithContext(context.Background(), request)
}

// DescribeInstances
// 本接口（DescribeInstance）用于在用户账户下获取消息队列 CKafka 实例列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
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

func NewDescribeInstancesDetailRequest() (request *DescribeInstancesDetailRequest) {
    request = &DescribeInstancesDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeInstancesDetail")
    
    
    return
}

func NewDescribeInstancesDetailResponse() (response *DescribeInstancesDetailResponse) {
    response = &DescribeInstancesDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesDetail
// 用户账户下获取实例列表详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeInstancesDetail(request *DescribeInstancesDetailRequest) (response *DescribeInstancesDetailResponse, err error) {
    return c.DescribeInstancesDetailWithContext(context.Background(), request)
}

// DescribeInstancesDetail
// 用户账户下获取实例列表详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeInstancesDetailWithContext(ctx context.Context, request *DescribeInstancesDetailRequest) (response *DescribeInstancesDetailResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeInstancesDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeInstancesDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRegionRequest() (request *DescribeRegionRequest) {
    request = &DescribeRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeRegion")
    
    
    return
}

func NewDescribeRegionResponse() (response *DescribeRegionResponse) {
    response = &DescribeRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRegion
// 枚举地域,只支持广州地域
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeRegion(request *DescribeRegionRequest) (response *DescribeRegionResponse, err error) {
    return c.DescribeRegionWithContext(context.Background(), request)
}

// DescribeRegion
// 枚举地域,只支持广州地域
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeRegionWithContext(ctx context.Context, request *DescribeRegionRequest) (response *DescribeRegionResponse, err error) {
    if request == nil {
        request = NewDescribeRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRouteRequest() (request *DescribeRouteRequest) {
    request = &DescribeRouteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeRoute")
    
    
    return
}

func NewDescribeRouteResponse() (response *DescribeRouteResponse) {
    response = &DescribeRouteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRoute
// 查看路由信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeRoute(request *DescribeRouteRequest) (response *DescribeRouteResponse, err error) {
    return c.DescribeRouteWithContext(context.Background(), request)
}

// DescribeRoute
// 查看路由信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeRouteWithContext(ctx context.Context, request *DescribeRouteRequest) (response *DescribeRouteResponse, err error) {
    if request == nil {
        request = NewDescribeRouteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoute require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRouteResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicRequest() (request *DescribeTopicRequest) {
    request = &DescribeTopicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTopic")
    
    
    return
}

func NewDescribeTopicResponse() (response *DescribeTopicResponse) {
    response = &DescribeTopicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopic
// 接口请求域名：https://ckafka.tencentcloudapi.com
//
// 本接口（DescribeTopic）用于在用户获取消息队列 CKafka 实例的主题列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTopic(request *DescribeTopicRequest) (response *DescribeTopicResponse, err error) {
    return c.DescribeTopicWithContext(context.Background(), request)
}

// DescribeTopic
// 接口请求域名：https://ckafka.tencentcloudapi.com
//
// 本接口（DescribeTopic）用于在用户获取消息队列 CKafka 实例的主题列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTopicWithContext(ctx context.Context, request *DescribeTopicRequest) (response *DescribeTopicResponse, err error) {
    if request == nil {
        request = NewDescribeTopicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicAttributesRequest() (request *DescribeTopicAttributesRequest) {
    request = &DescribeTopicAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTopicAttributes")
    
    
    return
}

func NewDescribeTopicAttributesResponse() (response *DescribeTopicAttributesResponse) {
    response = &DescribeTopicAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopicAttributes
// 获取主题属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTopicAttributes(request *DescribeTopicAttributesRequest) (response *DescribeTopicAttributesResponse, err error) {
    return c.DescribeTopicAttributesWithContext(context.Background(), request)
}

// DescribeTopicAttributes
// 获取主题属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTopicAttributesWithContext(ctx context.Context, request *DescribeTopicAttributesRequest) (response *DescribeTopicAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeTopicAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicDetailRequest() (request *DescribeTopicDetailRequest) {
    request = &DescribeTopicDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTopicDetail")
    
    
    return
}

func NewDescribeTopicDetailResponse() (response *DescribeTopicDetailResponse) {
    response = &DescribeTopicDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopicDetail
// 获取主题列表详情（仅控制台调用）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTopicDetail(request *DescribeTopicDetailRequest) (response *DescribeTopicDetailResponse, err error) {
    return c.DescribeTopicDetailWithContext(context.Background(), request)
}

// DescribeTopicDetail
// 获取主题列表详情（仅控制台调用）
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTopicDetailWithContext(ctx context.Context, request *DescribeTopicDetailRequest) (response *DescribeTopicDetailResponse, err error) {
    if request == nil {
        request = NewDescribeTopicDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicProduceConnectionRequest() (request *DescribeTopicProduceConnectionRequest) {
    request = &DescribeTopicProduceConnectionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTopicProduceConnection")
    
    
    return
}

func NewDescribeTopicProduceConnectionResponse() (response *DescribeTopicProduceConnectionResponse) {
    response = &DescribeTopicProduceConnectionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopicProduceConnection
// 查询topic 生产端连接信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
func (c *Client) DescribeTopicProduceConnection(request *DescribeTopicProduceConnectionRequest) (response *DescribeTopicProduceConnectionResponse, err error) {
    return c.DescribeTopicProduceConnectionWithContext(context.Background(), request)
}

// DescribeTopicProduceConnection
// 查询topic 生产端连接信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
func (c *Client) DescribeTopicProduceConnectionWithContext(ctx context.Context, request *DescribeTopicProduceConnectionRequest) (response *DescribeTopicProduceConnectionResponse, err error) {
    if request == nil {
        request = NewDescribeTopicProduceConnectionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicProduceConnection require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicProduceConnectionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicSubscribeGroupRequest() (request *DescribeTopicSubscribeGroupRequest) {
    request = &DescribeTopicSubscribeGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTopicSubscribeGroup")
    
    
    return
}

func NewDescribeTopicSubscribeGroupResponse() (response *DescribeTopicSubscribeGroupResponse) {
    response = &DescribeTopicSubscribeGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopicSubscribeGroup
// 查询订阅某主题消息分组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTopicSubscribeGroup(request *DescribeTopicSubscribeGroupRequest) (response *DescribeTopicSubscribeGroupResponse, err error) {
    return c.DescribeTopicSubscribeGroupWithContext(context.Background(), request)
}

// DescribeTopicSubscribeGroup
// 查询订阅某主题消息分组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeTopicSubscribeGroupWithContext(ctx context.Context, request *DescribeTopicSubscribeGroupRequest) (response *DescribeTopicSubscribeGroupResponse, err error) {
    if request == nil {
        request = NewDescribeTopicSubscribeGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicSubscribeGroup require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicSubscribeGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTopicSyncReplicaRequest() (request *DescribeTopicSyncReplicaRequest) {
    request = &DescribeTopicSyncReplicaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTopicSyncReplica")
    
    
    return
}

func NewDescribeTopicSyncReplicaResponse() (response *DescribeTopicSyncReplicaResponse) {
    response = &DescribeTopicSyncReplicaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTopicSyncReplica
// 获取Topic 副本详情信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopicSyncReplica(request *DescribeTopicSyncReplicaRequest) (response *DescribeTopicSyncReplicaResponse, err error) {
    return c.DescribeTopicSyncReplicaWithContext(context.Background(), request)
}

// DescribeTopicSyncReplica
// 获取Topic 副本详情信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTopicSyncReplicaWithContext(ctx context.Context, request *DescribeTopicSyncReplicaRequest) (response *DescribeTopicSyncReplicaResponse, err error) {
    if request == nil {
        request = NewDescribeTopicSyncReplicaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTopicSyncReplica require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTopicSyncReplicaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserRequest() (request *DescribeUserRequest) {
    request = &DescribeUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "DescribeUser")
    
    
    return
}

func NewDescribeUserResponse() (response *DescribeUserResponse) {
    response = &DescribeUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUser
// 查询用户信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeUser(request *DescribeUserRequest) (response *DescribeUserResponse, err error) {
    return c.DescribeUserWithContext(context.Background(), request)
}

// DescribeUser
// 查询用户信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) DescribeUserWithContext(ctx context.Context, request *DescribeUserRequest) (response *DescribeUserResponse, err error) {
    if request == nil {
        request = NewDescribeUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserResponse()
    err = c.Send(request, response)
    return
}

func NewFetchDatahubMessageByOffsetRequest() (request *FetchDatahubMessageByOffsetRequest) {
    request = &FetchDatahubMessageByOffsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "FetchDatahubMessageByOffset")
    
    
    return
}

func NewFetchDatahubMessageByOffsetResponse() (response *FetchDatahubMessageByOffsetResponse) {
    response = &FetchDatahubMessageByOffsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FetchDatahubMessageByOffset
// 根据指定offset位置的消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FetchDatahubMessageByOffset(request *FetchDatahubMessageByOffsetRequest) (response *FetchDatahubMessageByOffsetResponse, err error) {
    return c.FetchDatahubMessageByOffsetWithContext(context.Background(), request)
}

// FetchDatahubMessageByOffset
// 根据指定offset位置的消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FetchDatahubMessageByOffsetWithContext(ctx context.Context, request *FetchDatahubMessageByOffsetRequest) (response *FetchDatahubMessageByOffsetResponse, err error) {
    if request == nil {
        request = NewFetchDatahubMessageByOffsetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FetchDatahubMessageByOffset require credential")
    }

    request.SetContext(ctx)
    
    response = NewFetchDatahubMessageByOffsetResponse()
    err = c.Send(request, response)
    return
}

func NewFetchLatestDatahubMessageListRequest() (request *FetchLatestDatahubMessageListRequest) {
    request = &FetchLatestDatahubMessageListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "FetchLatestDatahubMessageList")
    
    
    return
}

func NewFetchLatestDatahubMessageListResponse() (response *FetchLatestDatahubMessageListResponse) {
    response = &FetchLatestDatahubMessageListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FetchLatestDatahubMessageList
// 查询最新消息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) FetchLatestDatahubMessageList(request *FetchLatestDatahubMessageListRequest) (response *FetchLatestDatahubMessageListResponse, err error) {
    return c.FetchLatestDatahubMessageListWithContext(context.Background(), request)
}

// FetchLatestDatahubMessageList
// 查询最新消息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) FetchLatestDatahubMessageListWithContext(ctx context.Context, request *FetchLatestDatahubMessageListRequest) (response *FetchLatestDatahubMessageListResponse, err error) {
    if request == nil {
        request = NewFetchLatestDatahubMessageListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FetchLatestDatahubMessageList require credential")
    }

    request.SetContext(ctx)
    
    response = NewFetchLatestDatahubMessageListResponse()
    err = c.Send(request, response)
    return
}

func NewFetchMessageByOffsetRequest() (request *FetchMessageByOffsetRequest) {
    request = &FetchMessageByOffsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "FetchMessageByOffset")
    
    
    return
}

func NewFetchMessageByOffsetResponse() (response *FetchMessageByOffsetResponse) {
    response = &FetchMessageByOffsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FetchMessageByOffset
// 根据指定offset位置的消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FetchMessageByOffset(request *FetchMessageByOffsetRequest) (response *FetchMessageByOffsetResponse, err error) {
    return c.FetchMessageByOffsetWithContext(context.Background(), request)
}

// FetchMessageByOffset
// 根据指定offset位置的消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) FetchMessageByOffsetWithContext(ctx context.Context, request *FetchMessageByOffsetRequest) (response *FetchMessageByOffsetResponse, err error) {
    if request == nil {
        request = NewFetchMessageByOffsetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FetchMessageByOffset require credential")
    }

    request.SetContext(ctx)
    
    response = NewFetchMessageByOffsetResponse()
    err = c.Send(request, response)
    return
}

func NewFetchMessageListByOffsetRequest() (request *FetchMessageListByOffsetRequest) {
    request = &FetchMessageListByOffsetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "FetchMessageListByOffset")
    
    
    return
}

func NewFetchMessageListByOffsetResponse() (response *FetchMessageListByOffsetResponse) {
    response = &FetchMessageListByOffsetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FetchMessageListByOffset
// 根据位点查询消息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) FetchMessageListByOffset(request *FetchMessageListByOffsetRequest) (response *FetchMessageListByOffsetResponse, err error) {
    return c.FetchMessageListByOffsetWithContext(context.Background(), request)
}

// FetchMessageListByOffset
// 根据位点查询消息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) FetchMessageListByOffsetWithContext(ctx context.Context, request *FetchMessageListByOffsetRequest) (response *FetchMessageListByOffsetResponse, err error) {
    if request == nil {
        request = NewFetchMessageListByOffsetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FetchMessageListByOffset require credential")
    }

    request.SetContext(ctx)
    
    response = NewFetchMessageListByOffsetResponse()
    err = c.Send(request, response)
    return
}

func NewInquireCkafkaPriceRequest() (request *InquireCkafkaPriceRequest) {
    request = &InquireCkafkaPriceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "InquireCkafkaPrice")
    
    
    return
}

func NewInquireCkafkaPriceResponse() (response *InquireCkafkaPriceResponse) {
    response = &InquireCkafkaPriceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquireCkafkaPrice
// Ckafka实例购买/续费询价
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) InquireCkafkaPrice(request *InquireCkafkaPriceRequest) (response *InquireCkafkaPriceResponse, err error) {
    return c.InquireCkafkaPriceWithContext(context.Background(), request)
}

// InquireCkafkaPrice
// Ckafka实例购买/续费询价
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) InquireCkafkaPriceWithContext(ctx context.Context, request *InquireCkafkaPriceRequest) (response *InquireCkafkaPriceResponse, err error) {
    if request == nil {
        request = NewInquireCkafkaPriceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquireCkafkaPrice require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquireCkafkaPriceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyConnectResourceRequest() (request *ModifyConnectResourceRequest) {
    request = &ModifyConnectResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "ModifyConnectResource")
    
    
    return
}

func NewModifyConnectResourceResponse() (response *ModifyConnectResourceResponse) {
    response = &ModifyConnectResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyConnectResource
// 编辑Datahub连接源
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
func (c *Client) ModifyConnectResource(request *ModifyConnectResourceRequest) (response *ModifyConnectResourceResponse, err error) {
    return c.ModifyConnectResourceWithContext(context.Background(), request)
}

// ModifyConnectResource
// 编辑Datahub连接源
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
func (c *Client) ModifyConnectResourceWithContext(ctx context.Context, request *ModifyConnectResourceRequest) (response *ModifyConnectResourceResponse, err error) {
    if request == nil {
        request = NewModifyConnectResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyConnectResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyConnectResourceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDatahubTaskRequest() (request *ModifyDatahubTaskRequest) {
    request = &ModifyDatahubTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "ModifyDatahubTask")
    
    
    return
}

func NewModifyDatahubTaskResponse() (response *ModifyDatahubTaskResponse) {
    response = &ModifyDatahubTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyDatahubTask
// 修改Datahub任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
func (c *Client) ModifyDatahubTask(request *ModifyDatahubTaskRequest) (response *ModifyDatahubTaskResponse, err error) {
    return c.ModifyDatahubTaskWithContext(context.Background(), request)
}

// ModifyDatahubTask
// 修改Datahub任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
func (c *Client) ModifyDatahubTaskWithContext(ctx context.Context, request *ModifyDatahubTaskRequest) (response *ModifyDatahubTaskResponse, err error) {
    if request == nil {
        request = NewModifyDatahubTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDatahubTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDatahubTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGroupOffsetsRequest() (request *ModifyGroupOffsetsRequest) {
    request = &ModifyGroupOffsetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "ModifyGroupOffsets")
    
    
    return
}

func NewModifyGroupOffsetsResponse() (response *ModifyGroupOffsetsResponse) {
    response = &ModifyGroupOffsetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyGroupOffsets
// 设置Groups 消费分组offset
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyGroupOffsets(request *ModifyGroupOffsetsRequest) (response *ModifyGroupOffsetsResponse, err error) {
    return c.ModifyGroupOffsetsWithContext(context.Background(), request)
}

// ModifyGroupOffsets
// 设置Groups 消费分组offset
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyGroupOffsetsWithContext(ctx context.Context, request *ModifyGroupOffsetsRequest) (response *ModifyGroupOffsetsResponse, err error) {
    if request == nil {
        request = NewModifyGroupOffsetsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyGroupOffsets require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyGroupOffsetsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstanceAttributesRequest() (request *ModifyInstanceAttributesRequest) {
    request = &ModifyInstanceAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "ModifyInstanceAttributes")
    
    
    return
}

func NewModifyInstanceAttributesResponse() (response *ModifyInstanceAttributesResponse) {
    response = &ModifyInstanceAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstanceAttributes
// 设置实例属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyInstanceAttributes(request *ModifyInstanceAttributesRequest) (response *ModifyInstanceAttributesResponse, err error) {
    return c.ModifyInstanceAttributesWithContext(context.Background(), request)
}

// ModifyInstanceAttributes
// 设置实例属性
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyInstanceAttributesWithContext(ctx context.Context, request *ModifyInstanceAttributesRequest) (response *ModifyInstanceAttributesResponse, err error) {
    if request == nil {
        request = NewModifyInstanceAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstanceAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstanceAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyInstancePreRequest() (request *ModifyInstancePreRequest) {
    request = &ModifyInstancePreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "ModifyInstancePre")
    
    
    return
}

func NewModifyInstancePreResponse() (response *ModifyInstancePreResponse) {
    response = &ModifyInstancePreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyInstancePre
// 预付费实例变配接口，调整磁盘，带宽
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyInstancePre(request *ModifyInstancePreRequest) (response *ModifyInstancePreResponse, err error) {
    return c.ModifyInstancePreWithContext(context.Background(), request)
}

// ModifyInstancePre
// 预付费实例变配接口，调整磁盘，带宽
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_NOTALLOWEDEMPTY = "InvalidParameterValue.NotAllowedEmpty"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyInstancePreWithContext(ctx context.Context, request *ModifyInstancePreRequest) (response *ModifyInstancePreResponse, err error) {
    if request == nil {
        request = NewModifyInstancePreRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyInstancePre require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyInstancePreResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPasswordRequest() (request *ModifyPasswordRequest) {
    request = &ModifyPasswordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "ModifyPassword")
    
    
    return
}

func NewModifyPasswordResponse() (response *ModifyPasswordResponse) {
    response = &ModifyPasswordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPassword
// 修改密码
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyPassword(request *ModifyPasswordRequest) (response *ModifyPasswordResponse, err error) {
    return c.ModifyPasswordWithContext(context.Background(), request)
}

// ModifyPassword
// 修改密码
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyPasswordWithContext(ctx context.Context, request *ModifyPasswordRequest) (response *ModifyPasswordResponse, err error) {
    if request == nil {
        request = NewModifyPasswordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPassword require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPasswordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyTopicAttributesRequest() (request *ModifyTopicAttributesRequest) {
    request = &ModifyTopicAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "ModifyTopicAttributes")
    
    
    return
}

func NewModifyTopicAttributesResponse() (response *ModifyTopicAttributesResponse) {
    response = &ModifyTopicAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyTopicAttributes
// 本接口用于修改主题属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyTopicAttributes(request *ModifyTopicAttributesRequest) (response *ModifyTopicAttributesResponse, err error) {
    return c.ModifyTopicAttributesWithContext(context.Background(), request)
}

// ModifyTopicAttributes
// 本接口用于修改主题属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_REPETITIONVALUE = "InvalidParameterValue.RepetitionValue"
//  INVALIDPARAMETERVALUE_SUBNETIDINVALID = "InvalidParameterValue.SubnetIdInvalid"
//  INVALIDPARAMETERVALUE_SUBNETNOTBELONGTOZONE = "InvalidParameterValue.SubnetNotBelongToZone"
//  INVALIDPARAMETERVALUE_VPCIDINVALID = "InvalidParameterValue.VpcIdInvalid"
//  INVALIDPARAMETERVALUE_WRONGACTION = "InvalidParameterValue.WrongAction"
//  INVALIDPARAMETERVALUE_ZONENOTSUPPORT = "InvalidParameterValue.ZoneNotSupport"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION_BATCHDELINSTANCELIMIT = "UnsupportedOperation.BatchDelInstanceLimit"
//  UNSUPPORTEDOPERATION_OSSREJECT = "UnsupportedOperation.OssReject"
func (c *Client) ModifyTopicAttributesWithContext(ctx context.Context, request *ModifyTopicAttributesRequest) (response *ModifyTopicAttributesResponse, err error) {
    if request == nil {
        request = NewModifyTopicAttributesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyTopicAttributes require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyTopicAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewSendMessageRequest() (request *SendMessageRequest) {
    request = &SendMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ckafka", APIVersion, "SendMessage")
    
    
    return
}

func NewSendMessageResponse() (response *SendMessageResponse) {
    response = &SendMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendMessage
// 通过HTTP接入层发送消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_RESOURCETASKPAUSED = "OperationDenied.ResourceTaskPaused"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SendMessage(request *SendMessageRequest) (response *SendMessageResponse, err error) {
    return c.SendMessageWithContext(context.Background(), request)
}

// SendMessage
// 通过HTTP接入层发送消息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  OPERATIONDENIED = "OperationDenied"
//  OPERATIONDENIED_RESOURCETASKPAUSED = "OperationDenied.ResourceTaskPaused"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SendMessageWithContext(ctx context.Context, request *SendMessageRequest) (response *SendMessageResponse, err error) {
    if request == nil {
        request = NewSendMessageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendMessageResponse()
    err = c.Send(request, response)
    return
}
