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

package v20191112

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-11-12"

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


func NewAttachCcnInstancesRequest() (request *AttachCcnInstancesRequest) {
    request = &AttachCcnInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "AttachCcnInstances")
    return
}

func NewAttachCcnInstancesResponse() (response *AttachCcnInstancesResponse) {
    response = &AttachCcnInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AttachCcnInstances
// 本接口（AttachCcnInstances）用于关联云联网实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CCNLIMITEXCEEDED = "LimitExceeded.CcnLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CCNRESOURCENOTFOUND = "ResourceNotFound.CcnResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CCNATTACHED = "UnsupportedOperation.CcnAttached"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
func (c *Client) AttachCcnInstances(request *AttachCcnInstancesRequest) (response *AttachCcnInstancesResponse, err error) {
    if request == nil {
        request = NewAttachCcnInstancesRequest()
    }
    response = NewAttachCcnInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewCopyFleetRequest() (request *CopyFleetRequest) {
    request = &CopyFleetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "CopyFleet")
    return
}

func NewCopyFleetResponse() (response *CopyFleetResponse) {
    response = &CopyFleetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CopyFleet
// 本接口（CopyFleet）用于复制服务器舰队。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_FLEETLIMITEXCEEDED = "LimitExceeded.FleetLimitExceeded"
//  LIMITEXCEEDED_INSTANCELIMITEXCEEDED = "LimitExceeded.InstanceLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CopyFleet(request *CopyFleetRequest) (response *CopyFleetResponse, err error) {
    if request == nil {
        request = NewCopyFleetRequest()
    }
    response = NewCopyFleetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAliasRequest() (request *CreateAliasRequest) {
    request = &CreateAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "CreateAlias")
    return
}

func NewCreateAliasResponse() (response *CreateAliasResponse) {
    response = &CreateAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAlias
// 本接口（CreateAlias）用于创建别名。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) CreateAlias(request *CreateAliasRequest) (response *CreateAliasResponse, err error) {
    if request == nil {
        request = NewCreateAliasRequest()
    }
    response = NewCreateAliasResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAssetRequest() (request *CreateAssetRequest) {
    request = &CreateAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "CreateAsset")
    return
}

func NewCreateAssetResponse() (response *CreateAssetResponse) {
    response = &CreateAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAsset
// 本接口（CreateAsset）用于创建生成包。
//
// 通过获取上传cos的临时密钥，将文件上传至cos，然后将生成包的zip名称下发给本接口完成资源创建。  
//
// 
//
// 上传文件至 cos支持俩种方式：
//
// 
//
// - 获取预签名方式，COS 简单上传
//
//     1. [GetUploadCredentials](https://cloud.tencent.com/document/product/1165/48727) 获取预签名信息
//
//     2. 使用 COS API 进行上传([参考文档](https://cloud.tencent.com/document/product/436/7749))
//
// -  临时密钥方式，COS 简单上传或者分块上传方式
//
//     1. [GetUploadCredentials](https://cloud.tencent.com/document/product/1165/48727)（获取上传 bucket  第一次调用需要，后续可以不用调用）
//
//     2. [GetUploadFederationToken](https://cloud.tencent.com/document/product/1165/48742) 获取临时密钥
//
//     3. 使用 COS API 进行上传([参考文档](https://cloud.tencent.com/document/product/436/7742))
//
// 
//
// 具体使用场景可以参考 [GetUploadCredentials](https://cloud.tencent.com/document/product/1165/48727) ,  [GetUploadFederationToken](https://cloud.tencent.com/document/product/1165/48742)和下面 CreateAsset 示例。  
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) CreateAsset(request *CreateAssetRequest) (response *CreateAssetResponse, err error) {
    if request == nil {
        request = NewCreateAssetRequest()
    }
    response = NewCreateAssetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAssetWithImageRequest() (request *CreateAssetWithImageRequest) {
    request = &CreateAssetWithImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "CreateAssetWithImage")
    return
}

func NewCreateAssetWithImageResponse() (response *CreateAssetWithImageResponse) {
    response = &CreateAssetWithImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAssetWithImage
// 本接口（CreateAssetWithImage）用于创建生成包镜像信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAssetWithImage(request *CreateAssetWithImageRequest) (response *CreateAssetWithImageResponse, err error) {
    if request == nil {
        request = NewCreateAssetWithImageRequest()
    }
    response = NewCreateAssetWithImageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFleetRequest() (request *CreateFleetRequest) {
    request = &CreateFleetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "CreateFleet")
    return
}

func NewCreateFleetResponse() (response *CreateFleetResponse) {
    response = &CreateFleetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFleet
// 本接口（CreateFleet）用于创建服务器舰队。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_FLEETLIMITEXCEEDED = "LimitExceeded.FleetLimitExceeded"
//  LIMITEXCEEDED_INSTANCELIMITEXCEEDED = "LimitExceeded.InstanceLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateFleet(request *CreateFleetRequest) (response *CreateFleetResponse, err error) {
    if request == nil {
        request = NewCreateFleetRequest()
    }
    response = NewCreateFleetResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGameServerSessionRequest() (request *CreateGameServerSessionRequest) {
    request = &CreateGameServerSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "CreateGameServerSession")
    return
}

func NewCreateGameServerSessionResponse() (response *CreateGameServerSessionResponse) {
    response = &CreateGameServerSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateGameServerSession
// 本接口（CreateGameServerSession）用于创建游戏服务会话。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) CreateGameServerSession(request *CreateGameServerSessionRequest) (response *CreateGameServerSessionResponse, err error) {
    if request == nil {
        request = NewCreateGameServerSessionRequest()
    }
    response = NewCreateGameServerSessionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGameServerSessionQueueRequest() (request *CreateGameServerSessionQueueRequest) {
    request = &CreateGameServerSessionQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "CreateGameServerSessionQueue")
    return
}

func NewCreateGameServerSessionQueueResponse() (response *CreateGameServerSessionQueueResponse) {
    response = &CreateGameServerSessionQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateGameServerSessionQueue
// 本接口（CreateGameServerSessionQueue）用于创建游戏服务器会话队列。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) CreateGameServerSessionQueue(request *CreateGameServerSessionQueueRequest) (response *CreateGameServerSessionQueueResponse, err error) {
    if request == nil {
        request = NewCreateGameServerSessionQueueRequest()
    }
    response = NewCreateGameServerSessionQueueResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAliasRequest() (request *DeleteAliasRequest) {
    request = &DeleteAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DeleteAlias")
    return
}

func NewDeleteAliasResponse() (response *DeleteAliasResponse) {
    response = &DeleteAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAlias
// 本接口（DeleteAlias）用于删除别名。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DeleteAlias(request *DeleteAliasRequest) (response *DeleteAliasResponse, err error) {
    if request == nil {
        request = NewDeleteAliasRequest()
    }
    response = NewDeleteAliasResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAssetRequest() (request *DeleteAssetRequest) {
    request = &DeleteAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DeleteAsset")
    return
}

func NewDeleteAssetResponse() (response *DeleteAssetResponse) {
    response = &DeleteAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAsset
// 本接口（DeleteAsset）用于删除生成包。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAsset(request *DeleteAssetRequest) (response *DeleteAssetResponse, err error) {
    if request == nil {
        request = NewDeleteAssetRequest()
    }
    response = NewDeleteAssetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFleetRequest() (request *DeleteFleetRequest) {
    request = &DeleteFleetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DeleteFleet")
    return
}

func NewDeleteFleetResponse() (response *DeleteFleetResponse) {
    response = &DeleteFleetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteFleet
// 本接口（DeleteFleet）用于删除服务器舰队。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteFleet(request *DeleteFleetRequest) (response *DeleteFleetResponse, err error) {
    if request == nil {
        request = NewDeleteFleetRequest()
    }
    response = NewDeleteFleetResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGameServerSessionQueueRequest() (request *DeleteGameServerSessionQueueRequest) {
    request = &DeleteGameServerSessionQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DeleteGameServerSessionQueue")
    return
}

func NewDeleteGameServerSessionQueueResponse() (response *DeleteGameServerSessionQueueResponse) {
    response = &DeleteGameServerSessionQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteGameServerSessionQueue
// 本接口（DeleteGameServerSessionQueue）用于删除游戏服务器会话队列。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DeleteGameServerSessionQueue(request *DeleteGameServerSessionQueueRequest) (response *DeleteGameServerSessionQueueResponse, err error) {
    if request == nil {
        request = NewDeleteGameServerSessionQueueRequest()
    }
    response = NewDeleteGameServerSessionQueueResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteScalingPolicyRequest() (request *DeleteScalingPolicyRequest) {
    request = &DeleteScalingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DeleteScalingPolicy")
    return
}

func NewDeleteScalingPolicyResponse() (response *DeleteScalingPolicyResponse) {
    response = &DeleteScalingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteScalingPolicy
// 本接口（DeleteScalingPolicy）用于删除服务器舰队的扩缩容策略。
//
// 通过服务器舰队ID和策略名称删除服务器舰队的扩缩容策略，只传递服务器舰队ID时，会将这个服务器舰队下的所有策略都删除。
//
// 传递策略名称时，单独删除策略名称对应的策略。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DeleteScalingPolicy(request *DeleteScalingPolicyRequest) (response *DeleteScalingPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteScalingPolicyRequest()
    }
    response = NewDeleteScalingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTimerScalingPolicyRequest() (request *DeleteTimerScalingPolicyRequest) {
    request = &DeleteTimerScalingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DeleteTimerScalingPolicy")
    return
}

func NewDeleteTimerScalingPolicyResponse() (response *DeleteTimerScalingPolicyResponse) {
    response = &DeleteTimerScalingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteTimerScalingPolicy
// 本接口（DeleteTimerScalingPolicy）用于删除fleet下的定时器。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DeleteTimerScalingPolicy(request *DeleteTimerScalingPolicyRequest) (response *DeleteTimerScalingPolicyResponse, err error) {
    if request == nil {
        request = NewDeleteTimerScalingPolicyRequest()
    }
    response = NewDeleteTimerScalingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAliasRequest() (request *DescribeAliasRequest) {
    request = &DescribeAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeAlias")
    return
}

func NewDescribeAliasResponse() (response *DescribeAliasResponse) {
    response = &DescribeAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAlias
// 本接口（DescribeAlias）用于获取别名详情。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeAlias(request *DescribeAliasRequest) (response *DescribeAliasResponse, err error) {
    if request == nil {
        request = NewDescribeAliasRequest()
    }
    response = NewDescribeAliasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetRequest() (request *DescribeAssetRequest) {
    request = &DescribeAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeAsset")
    return
}

func NewDescribeAssetResponse() (response *DescribeAssetResponse) {
    response = &DescribeAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAsset
// 本接口（DescribeAsset）获取生成包信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeAsset(request *DescribeAssetRequest) (response *DescribeAssetResponse, err error) {
    if request == nil {
        request = NewDescribeAssetRequest()
    }
    response = NewDescribeAssetResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetSystemsRequest() (request *DescribeAssetSystemsRequest) {
    request = &DescribeAssetSystemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeAssetSystems")
    return
}

func NewDescribeAssetSystemsResponse() (response *DescribeAssetSystemsResponse) {
    response = &DescribeAssetSystemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssetSystems
// 本接口（DescribeAssetSystems）用于获取生成包支持的操作系统。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeAssetSystems(request *DescribeAssetSystemsRequest) (response *DescribeAssetSystemsResponse, err error) {
    if request == nil {
        request = NewDescribeAssetSystemsRequest()
    }
    response = NewDescribeAssetSystemsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAssetsRequest() (request *DescribeAssetsRequest) {
    request = &DescribeAssetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeAssets")
    return
}

func NewDescribeAssetsResponse() (response *DescribeAssetsResponse) {
    response = &DescribeAssetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAssets
// 本接口（DescribeAssets）用于获取生成包列表。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeAssets(request *DescribeAssetsRequest) (response *DescribeAssetsResponse, err error) {
    if request == nil {
        request = NewDescribeAssetsRequest()
    }
    response = NewDescribeAssetsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCcnInstancesRequest() (request *DescribeCcnInstancesRequest) {
    request = &DescribeCcnInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeCcnInstances")
    return
}

func NewDescribeCcnInstancesResponse() (response *DescribeCcnInstancesResponse) {
    response = &DescribeCcnInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCcnInstances
// 本接口（DescribeCcnInstances）用于查询云联网实例。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CCNLIMITEXCEEDED = "LimitExceeded.CcnLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CCNRESOURCENOTFOUND = "ResourceNotFound.CcnResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CCNATTACHED = "UnsupportedOperation.CcnAttached"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
func (c *Client) DescribeCcnInstances(request *DescribeCcnInstancesRequest) (response *DescribeCcnInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeCcnInstancesRequest()
    }
    response = NewDescribeCcnInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFleetAttributesRequest() (request *DescribeFleetAttributesRequest) {
    request = &DescribeFleetAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeFleetAttributes")
    return
}

func NewDescribeFleetAttributesResponse() (response *DescribeFleetAttributesResponse) {
    response = &DescribeFleetAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFleetAttributes
// 本接口（DescribeFleetAttributes）用于查询服务器舰队属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeFleetAttributes(request *DescribeFleetAttributesRequest) (response *DescribeFleetAttributesResponse, err error) {
    if request == nil {
        request = NewDescribeFleetAttributesRequest()
    }
    response = NewDescribeFleetAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFleetCapacityRequest() (request *DescribeFleetCapacityRequest) {
    request = &DescribeFleetCapacityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeFleetCapacity")
    return
}

func NewDescribeFleetCapacityResponse() (response *DescribeFleetCapacityResponse) {
    response = &DescribeFleetCapacityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFleetCapacity
// 本接口（DescribeFleetCapacity）用于查询服务部署容量配置。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeFleetCapacity(request *DescribeFleetCapacityRequest) (response *DescribeFleetCapacityResponse, err error) {
    if request == nil {
        request = NewDescribeFleetCapacityRequest()
    }
    response = NewDescribeFleetCapacityResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFleetEventsRequest() (request *DescribeFleetEventsRequest) {
    request = &DescribeFleetEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeFleetEvents")
    return
}

func NewDescribeFleetEventsResponse() (response *DescribeFleetEventsResponse) {
    response = &DescribeFleetEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFleetEvents
// 本接口（DescribeFleetEvents）用于查询服务器舰队相关的事件列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeFleetEvents(request *DescribeFleetEventsRequest) (response *DescribeFleetEventsResponse, err error) {
    if request == nil {
        request = NewDescribeFleetEventsRequest()
    }
    response = NewDescribeFleetEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFleetPortSettingsRequest() (request *DescribeFleetPortSettingsRequest) {
    request = &DescribeFleetPortSettingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeFleetPortSettings")
    return
}

func NewDescribeFleetPortSettingsResponse() (response *DescribeFleetPortSettingsResponse) {
    response = &DescribeFleetPortSettingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFleetPortSettings
// 本接口（DescribeFleetPortSettings）用于获取服务器舰队安全组信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFleetPortSettings(request *DescribeFleetPortSettingsRequest) (response *DescribeFleetPortSettingsResponse, err error) {
    if request == nil {
        request = NewDescribeFleetPortSettingsRequest()
    }
    response = NewDescribeFleetPortSettingsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFleetRelatedResourcesRequest() (request *DescribeFleetRelatedResourcesRequest) {
    request = &DescribeFleetRelatedResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeFleetRelatedResources")
    return
}

func NewDescribeFleetRelatedResourcesResponse() (response *DescribeFleetRelatedResourcesResponse) {
    response = &DescribeFleetRelatedResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFleetRelatedResources
// 本接口（DescribeFleetRelatedResources）用于获取与游戏服务器舰队关联的资源信息，如别名、队列
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFleetRelatedResources(request *DescribeFleetRelatedResourcesRequest) (response *DescribeFleetRelatedResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeFleetRelatedResourcesRequest()
    }
    response = NewDescribeFleetRelatedResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFleetStatisticDetailsRequest() (request *DescribeFleetStatisticDetailsRequest) {
    request = &DescribeFleetStatisticDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeFleetStatisticDetails")
    return
}

func NewDescribeFleetStatisticDetailsResponse() (response *DescribeFleetStatisticDetailsResponse) {
    response = &DescribeFleetStatisticDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFleetStatisticDetails
// 本接口（DescribeFleetStatisticDetails）用于查询服务部署统计详情。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFleetStatisticDetails(request *DescribeFleetStatisticDetailsRequest) (response *DescribeFleetStatisticDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeFleetStatisticDetailsRequest()
    }
    response = NewDescribeFleetStatisticDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFleetStatisticFlowsRequest() (request *DescribeFleetStatisticFlowsRequest) {
    request = &DescribeFleetStatisticFlowsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeFleetStatisticFlows")
    return
}

func NewDescribeFleetStatisticFlowsResponse() (response *DescribeFleetStatisticFlowsResponse) {
    response = &DescribeFleetStatisticFlowsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFleetStatisticFlows
// 本接口（DescribeFleetStatisticFlows）用于查询服务部署统计用量。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFleetStatisticFlows(request *DescribeFleetStatisticFlowsRequest) (response *DescribeFleetStatisticFlowsResponse, err error) {
    if request == nil {
        request = NewDescribeFleetStatisticFlowsRequest()
    }
    response = NewDescribeFleetStatisticFlowsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFleetStatisticSummaryRequest() (request *DescribeFleetStatisticSummaryRequest) {
    request = &DescribeFleetStatisticSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeFleetStatisticSummary")
    return
}

func NewDescribeFleetStatisticSummaryResponse() (response *DescribeFleetStatisticSummaryResponse) {
    response = &DescribeFleetStatisticSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFleetStatisticSummary
// 本接口（DescribeFleetStatisticSummary）用于查询服务部署统计汇总信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFleetStatisticSummary(request *DescribeFleetStatisticSummaryRequest) (response *DescribeFleetStatisticSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeFleetStatisticSummaryRequest()
    }
    response = NewDescribeFleetStatisticSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFleetUtilizationRequest() (request *DescribeFleetUtilizationRequest) {
    request = &DescribeFleetUtilizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeFleetUtilization")
    return
}

func NewDescribeFleetUtilizationResponse() (response *DescribeFleetUtilizationResponse) {
    response = &DescribeFleetUtilizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFleetUtilization
// 本接口（DescribeFleetUtilization）用于查询服务器舰队的利用率信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFleetUtilization(request *DescribeFleetUtilizationRequest) (response *DescribeFleetUtilizationResponse, err error) {
    if request == nil {
        request = NewDescribeFleetUtilizationRequest()
    }
    response = NewDescribeFleetUtilizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGameServerSessionDetailsRequest() (request *DescribeGameServerSessionDetailsRequest) {
    request = &DescribeGameServerSessionDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeGameServerSessionDetails")
    return
}

func NewDescribeGameServerSessionDetailsResponse() (response *DescribeGameServerSessionDetailsResponse) {
    response = &DescribeGameServerSessionDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGameServerSessionDetails
// 本接口（DescribeGameServerSessionDetails）用于查询游戏服务器会话详情列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeGameServerSessionDetails(request *DescribeGameServerSessionDetailsRequest) (response *DescribeGameServerSessionDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeGameServerSessionDetailsRequest()
    }
    response = NewDescribeGameServerSessionDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGameServerSessionPlacementRequest() (request *DescribeGameServerSessionPlacementRequest) {
    request = &DescribeGameServerSessionPlacementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeGameServerSessionPlacement")
    return
}

func NewDescribeGameServerSessionPlacementResponse() (response *DescribeGameServerSessionPlacementResponse) {
    response = &DescribeGameServerSessionPlacementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGameServerSessionPlacement
// 本接口（DescribeGameServerSessionPlacement）用于查询游戏服务器会话的放置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeGameServerSessionPlacement(request *DescribeGameServerSessionPlacementRequest) (response *DescribeGameServerSessionPlacementResponse, err error) {
    if request == nil {
        request = NewDescribeGameServerSessionPlacementRequest()
    }
    response = NewDescribeGameServerSessionPlacementResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGameServerSessionQueuesRequest() (request *DescribeGameServerSessionQueuesRequest) {
    request = &DescribeGameServerSessionQueuesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeGameServerSessionQueues")
    return
}

func NewDescribeGameServerSessionQueuesResponse() (response *DescribeGameServerSessionQueuesResponse) {
    response = &DescribeGameServerSessionQueuesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGameServerSessionQueues
// 本接口（DescribeGameServerSessionQueues）用于查询游戏服务器会话队列。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeGameServerSessionQueues(request *DescribeGameServerSessionQueuesRequest) (response *DescribeGameServerSessionQueuesResponse, err error) {
    if request == nil {
        request = NewDescribeGameServerSessionQueuesRequest()
    }
    response = NewDescribeGameServerSessionQueuesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGameServerSessionsRequest() (request *DescribeGameServerSessionsRequest) {
    request = &DescribeGameServerSessionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeGameServerSessions")
    return
}

func NewDescribeGameServerSessionsResponse() (response *DescribeGameServerSessionsResponse) {
    response = &DescribeGameServerSessionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGameServerSessions
// 本接口（DescribeGameServerSessions）用于查询游戏服务器会话列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeGameServerSessions(request *DescribeGameServerSessionsRequest) (response *DescribeGameServerSessionsResponse, err error) {
    if request == nil {
        request = NewDescribeGameServerSessionsRequest()
    }
    response = NewDescribeGameServerSessionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceLimitRequest() (request *DescribeInstanceLimitRequest) {
    request = &DescribeInstanceLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeInstanceLimit")
    return
}

func NewDescribeInstanceLimitResponse() (response *DescribeInstanceLimitResponse) {
    response = &DescribeInstanceLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceLimit
// 本接口（DescribeInstanceLimit）用于查询用户实例数限额。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeInstanceLimit(request *DescribeInstanceLimitRequest) (response *DescribeInstanceLimitResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceLimitRequest()
    }
    response = NewDescribeInstanceLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstanceTypesRequest() (request *DescribeInstanceTypesRequest) {
    request = &DescribeInstanceTypesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeInstanceTypes")
    return
}

func NewDescribeInstanceTypesResponse() (response *DescribeInstanceTypesResponse) {
    response = &DescribeInstanceTypesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstanceTypes
// 本接口（DescribeInstanceTypes）用于获取服务器实例类型列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeInstanceTypes(request *DescribeInstanceTypesRequest) (response *DescribeInstanceTypesResponse, err error) {
    if request == nil {
        request = NewDescribeInstanceTypesRequest()
    }
    response = NewDescribeInstanceTypesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstances
// 本接口（DescribeInstances）用于查询服务器实例列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesExtendRequest() (request *DescribeInstancesExtendRequest) {
    request = &DescribeInstancesExtendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeInstancesExtend")
    return
}

func NewDescribeInstancesExtendResponse() (response *DescribeInstancesExtendResponse) {
    response = &DescribeInstancesExtendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeInstancesExtend
// 本接口（DescribeInstancesExtend）用于查询实例扩展信息列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeInstancesExtend(request *DescribeInstancesExtendRequest) (response *DescribeInstancesExtendResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesExtendRequest()
    }
    response = NewDescribeInstancesExtendResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePlayerSessionsRequest() (request *DescribePlayerSessionsRequest) {
    request = &DescribePlayerSessionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribePlayerSessions")
    return
}

func NewDescribePlayerSessionsResponse() (response *DescribePlayerSessionsResponse) {
    response = &DescribePlayerSessionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePlayerSessions
// 本接口（DescribePlayerSessions）用于获取玩家会话列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribePlayerSessions(request *DescribePlayerSessionsRequest) (response *DescribePlayerSessionsResponse, err error) {
    if request == nil {
        request = NewDescribePlayerSessionsRequest()
    }
    response = NewDescribePlayerSessionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRuntimeConfigurationRequest() (request *DescribeRuntimeConfigurationRequest) {
    request = &DescribeRuntimeConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeRuntimeConfiguration")
    return
}

func NewDescribeRuntimeConfigurationResponse() (response *DescribeRuntimeConfigurationResponse) {
    response = &DescribeRuntimeConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRuntimeConfiguration
// 本接口（DescribeRuntimeConfiguration）用于获取服务器舰队运行配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeRuntimeConfiguration(request *DescribeRuntimeConfigurationRequest) (response *DescribeRuntimeConfigurationResponse, err error) {
    if request == nil {
        request = NewDescribeRuntimeConfigurationRequest()
    }
    response = NewDescribeRuntimeConfigurationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScalingPoliciesRequest() (request *DescribeScalingPoliciesRequest) {
    request = &DescribeScalingPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeScalingPolicies")
    return
}

func NewDescribeScalingPoliciesResponse() (response *DescribeScalingPoliciesResponse) {
    response = &DescribeScalingPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScalingPolicies
// 本接口（DescribeScalingPolicies）用于查询服务器舰队的动态扩缩容策略列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeScalingPolicies(request *DescribeScalingPoliciesRequest) (response *DescribeScalingPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeScalingPoliciesRequest()
    }
    response = NewDescribeScalingPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTimerScalingPoliciesRequest() (request *DescribeTimerScalingPoliciesRequest) {
    request = &DescribeTimerScalingPoliciesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeTimerScalingPolicies")
    return
}

func NewDescribeTimerScalingPoliciesResponse() (response *DescribeTimerScalingPoliciesResponse) {
    response = &DescribeTimerScalingPoliciesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTimerScalingPolicies
// 本接口（DescribeTimerScalingPolicies）用于查询fleet下的定时器列表。可以通过fleetid，定时器名称分页查询。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) DescribeTimerScalingPolicies(request *DescribeTimerScalingPoliciesRequest) (response *DescribeTimerScalingPoliciesResponse, err error) {
    if request == nil {
        request = NewDescribeTimerScalingPoliciesRequest()
    }
    response = NewDescribeTimerScalingPoliciesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserQuotaRequest() (request *DescribeUserQuotaRequest) {
    request = &DescribeUserQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeUserQuota")
    return
}

func NewDescribeUserQuotaResponse() (response *DescribeUserQuotaResponse) {
    response = &DescribeUserQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserQuota
// 本接口（DescribeUserQuota）获取用户单个模块配额。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserQuota(request *DescribeUserQuotaRequest) (response *DescribeUserQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeUserQuotaRequest()
    }
    response = NewDescribeUserQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserQuotasRequest() (request *DescribeUserQuotasRequest) {
    request = &DescribeUserQuotasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DescribeUserQuotas")
    return
}

func NewDescribeUserQuotasResponse() (response *DescribeUserQuotasResponse) {
    response = &DescribeUserQuotasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserQuotas
// 本接口（DescribeUserQuotas）用于获取用户配额
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserQuotas(request *DescribeUserQuotasRequest) (response *DescribeUserQuotasResponse, err error) {
    if request == nil {
        request = NewDescribeUserQuotasRequest()
    }
    response = NewDescribeUserQuotasResponse()
    err = c.Send(request, response)
    return
}

func NewDetachCcnInstancesRequest() (request *DetachCcnInstancesRequest) {
    request = &DetachCcnInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "DetachCcnInstances")
    return
}

func NewDetachCcnInstancesResponse() (response *DetachCcnInstancesResponse) {
    response = &DetachCcnInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetachCcnInstances
// 本接口（DetachCcnInstances）用于解关联云联网实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CCNLIMITEXCEEDED = "LimitExceeded.CcnLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_CCNRESOURCENOTFOUND = "ResourceNotFound.CcnResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CCNATTACHED = "UnsupportedOperation.CcnAttached"
//  UNSUPPORTEDOPERATION_UINNOTFOUND = "UnsupportedOperation.UinNotFound"
//  UNSUPPORTEDOPERATION_UNABLECROSSBORDER = "UnsupportedOperation.UnableCrossBorder"
func (c *Client) DetachCcnInstances(request *DetachCcnInstancesRequest) (response *DetachCcnInstancesResponse, err error) {
    if request == nil {
        request = NewDetachCcnInstancesRequest()
    }
    response = NewDetachCcnInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewEndGameServerSessionAndProcessRequest() (request *EndGameServerSessionAndProcessRequest) {
    request = &EndGameServerSessionAndProcessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "EndGameServerSessionAndProcess")
    return
}

func NewEndGameServerSessionAndProcessResponse() (response *EndGameServerSessionAndProcessResponse) {
    response = &EndGameServerSessionAndProcessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EndGameServerSessionAndProcess
// 本接口（EndGameServerSessionAndProcess）用于终止游戏服务器会话和对应的进程，适用于时限保护和不保护。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) EndGameServerSessionAndProcess(request *EndGameServerSessionAndProcessRequest) (response *EndGameServerSessionAndProcessResponse, err error) {
    if request == nil {
        request = NewEndGameServerSessionAndProcessRequest()
    }
    response = NewEndGameServerSessionAndProcessResponse()
    err = c.Send(request, response)
    return
}

func NewGetGameServerInstanceLogUrlRequest() (request *GetGameServerInstanceLogUrlRequest) {
    request = &GetGameServerInstanceLogUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "GetGameServerInstanceLogUrl")
    return
}

func NewGetGameServerInstanceLogUrlResponse() (response *GetGameServerInstanceLogUrlResponse) {
    response = &GetGameServerInstanceLogUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetGameServerInstanceLogUrl
// 本接口用于获取游戏服务器实例的日志URL。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) GetGameServerInstanceLogUrl(request *GetGameServerInstanceLogUrlRequest) (response *GetGameServerInstanceLogUrlResponse, err error) {
    if request == nil {
        request = NewGetGameServerInstanceLogUrlRequest()
    }
    response = NewGetGameServerInstanceLogUrlResponse()
    err = c.Send(request, response)
    return
}

func NewGetGameServerSessionLogUrlRequest() (request *GetGameServerSessionLogUrlRequest) {
    request = &GetGameServerSessionLogUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "GetGameServerSessionLogUrl")
    return
}

func NewGetGameServerSessionLogUrlResponse() (response *GetGameServerSessionLogUrlResponse) {
    response = &GetGameServerSessionLogUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetGameServerSessionLogUrl
// 本接口（GetGameServerSessionLogUrl）用于获取游戏服务器会话的日志URL。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) GetGameServerSessionLogUrl(request *GetGameServerSessionLogUrlRequest) (response *GetGameServerSessionLogUrlResponse, err error) {
    if request == nil {
        request = NewGetGameServerSessionLogUrlRequest()
    }
    response = NewGetGameServerSessionLogUrlResponse()
    err = c.Send(request, response)
    return
}

func NewGetInstanceAccessRequest() (request *GetInstanceAccessRequest) {
    request = &GetInstanceAccessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "GetInstanceAccess")
    return
}

func NewGetInstanceAccessResponse() (response *GetInstanceAccessResponse) {
    response = &GetInstanceAccessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetInstanceAccess
// 本接口（GetInstanceAccess）用于获取实例登录所需要的凭据。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GetInstanceAccess(request *GetInstanceAccessRequest) (response *GetInstanceAccessResponse, err error) {
    if request == nil {
        request = NewGetInstanceAccessRequest()
    }
    response = NewGetInstanceAccessResponse()
    err = c.Send(request, response)
    return
}

func NewGetUploadCredentialsRequest() (request *GetUploadCredentialsRequest) {
    request = &GetUploadCredentialsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "GetUploadCredentials")
    return
}

func NewGetUploadCredentialsResponse() (response *GetUploadCredentialsResponse) {
    response = &GetUploadCredentialsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetUploadCredentials
// 本接口（GetUploadCredentials）获取上传文件授权信息。
//
// 通过 [GetUploadCredentials](https://cloud.tencent.com/document/product/1165/48727) 接口获取临时授权信息后，调用 COS API将数据上传，根据上传的 BucketKey 信息进行生成包 [CreateAsset](https://cloud.tencent.com/document/product/1165/48731) 的创建。参考下面的示例部分。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) GetUploadCredentials(request *GetUploadCredentialsRequest) (response *GetUploadCredentialsResponse, err error) {
    if request == nil {
        request = NewGetUploadCredentialsRequest()
    }
    response = NewGetUploadCredentialsResponse()
    err = c.Send(request, response)
    return
}

func NewGetUploadFederationTokenRequest() (request *GetUploadFederationTokenRequest) {
    request = &GetUploadFederationTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "GetUploadFederationToken")
    return
}

func NewGetUploadFederationTokenResponse() (response *GetUploadFederationTokenResponse) {
    response = &GetUploadFederationTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetUploadFederationToken
// 本接口（GetUploadFederationToken）用于 获取生成包上传所需要的临时密钥。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) GetUploadFederationToken(request *GetUploadFederationTokenRequest) (response *GetUploadFederationTokenResponse, err error) {
    if request == nil {
        request = NewGetUploadFederationTokenRequest()
    }
    response = NewGetUploadFederationTokenResponse()
    err = c.Send(request, response)
    return
}

func NewJoinGameServerSessionRequest() (request *JoinGameServerSessionRequest) {
    request = &JoinGameServerSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "JoinGameServerSession")
    return
}

func NewJoinGameServerSessionResponse() (response *JoinGameServerSessionResponse) {
    response = &JoinGameServerSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// JoinGameServerSession
// 本接口（JoinGameServerSession）用于加入游戏服务器会话。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) JoinGameServerSession(request *JoinGameServerSessionRequest) (response *JoinGameServerSessionResponse, err error) {
    if request == nil {
        request = NewJoinGameServerSessionRequest()
    }
    response = NewJoinGameServerSessionResponse()
    err = c.Send(request, response)
    return
}

func NewJoinGameServerSessionBatchRequest() (request *JoinGameServerSessionBatchRequest) {
    request = &JoinGameServerSessionBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "JoinGameServerSessionBatch")
    return
}

func NewJoinGameServerSessionBatchResponse() (response *JoinGameServerSessionBatchResponse) {
    response = &JoinGameServerSessionBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// JoinGameServerSessionBatch
// 本接口（JoinGameServerSessionBatch）用于批量加入游戏服务器会话。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) JoinGameServerSessionBatch(request *JoinGameServerSessionBatchRequest) (response *JoinGameServerSessionBatchResponse, err error) {
    if request == nil {
        request = NewJoinGameServerSessionBatchRequest()
    }
    response = NewJoinGameServerSessionBatchResponse()
    err = c.Send(request, response)
    return
}

func NewListAliasesRequest() (request *ListAliasesRequest) {
    request = &ListAliasesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "ListAliases")
    return
}

func NewListAliasesResponse() (response *ListAliasesResponse) {
    response = &ListAliasesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListAliases
// 本接口（ListAliases）用于检索帐户下的所有别名。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) ListAliases(request *ListAliasesRequest) (response *ListAliasesResponse, err error) {
    if request == nil {
        request = NewListAliasesRequest()
    }
    response = NewListAliasesResponse()
    err = c.Send(request, response)
    return
}

func NewListFleetsRequest() (request *ListFleetsRequest) {
    request = &ListFleetsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "ListFleets")
    return
}

func NewListFleetsResponse() (response *ListFleetsResponse) {
    response = &ListFleetsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ListFleets
// 本接口（ListFleets）用于获取服务器舰队列表。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ListFleets(request *ListFleetsRequest) (response *ListFleetsResponse, err error) {
    if request == nil {
        request = NewListFleetsRequest()
    }
    response = NewListFleetsResponse()
    err = c.Send(request, response)
    return
}

func NewPutScalingPolicyRequest() (request *PutScalingPolicyRequest) {
    request = &PutScalingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "PutScalingPolicy")
    return
}

func NewPutScalingPolicyResponse() (response *PutScalingPolicyResponse) {
    response = &PutScalingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PutScalingPolicy
// 本接口（PutScalingPolicy）用于设置服务器舰队的动态扩缩容策略。
//
// 
//
// 通过此接口可以增加或者更新服务器舰队的扩缩容策略。
//
// 服务器舰队可以有多个扩缩容策略，但是只有一个TargetBased基于目标的策略。
//
// 
//
// ## TargetBased基于目标的策略
//
// 
//
// TargetBased策略计算的指标是PercentAvailableGameSessions，这个策略用于计算当前服务器舰队应该有多少个CVM实例来支撑和分配游戏会话。
//
// PercentAvailableGameSessions表示服务器舰队的缓冲值；用来计算服务器舰队在当前容量下可以处理的额外玩家会话数量。
//
// 如果使用基于目标的策略，可以按照业务需求设置一个期望的缓冲区大小，GSE的会按照配置的策略来扩容和缩容到这个目标要求的CVM实例数。
//
// 
//
// 例如：客户可以设置同时承载100个游戏会话的服务器舰队预留10%的缓冲区。GSE会按照这个策略执行时，若服务器舰队的可用容量低于或高于10%的游戏服务器会话时，执行扩缩容动作。
//
// GSE按照策略期望，扩容新CVM实例或缩容未使用的实例，保持在10%左右的缓冲区。
//
// 
//
// #### 请求参数取值说明
//
// 
//
// ```
//
// Name取值策略名称，
//
// FleetId取值为选择的服务器舰队ID，
//
// PolicyType取值TargetBased，
//
// MetricName取值PercentAvailableGameSessions，
//
// TargetConfiguration取值为所需的缓冲区值对象，
//
// 其他参数不用传递。
//
// 请求成功时，将返回策略名称。扩缩容策略在成功创建立即自动生效。
//
// ```
//
// 
//
// 
//
// 
//
// ## RuleBased基于规则的策略
//
// 
//
// ####  请求参数取值说明
//
// 
//
// ```
//
// Name取值策略名称，
//
// FleetId取值为选择的服务器舰队ID，
//
// PolicyType取值RuleBased，
//
// MetricName取值（AvailableGameServerSessions，AvailableCustomCount，PercentAvailableCustomCount，ActiveInstances，IdleInstances，CurrentPlayerSessions和PercentIdleInstances）说明见备注1，
//
// 其他参数不用传递。
//
// ComparisonOperator取值为 >,>=,<,<=这4个比较符号中的一个，
//
// Threshold取值为指标MetricName达到的阈值是多少，
//
// EvaluationPeriods取值为指标MetricName达到的阈值持续了多少时间，单位是分钟，
//
// ScalingAdjustmentType取值（ChangeInCapacity，ExactCapacity，PercentChangeInCapacity）说明见备注2
//
// ScalingAdjustment取值为指标MetricName达到的阈值的条件后，扩缩容多少个CVM实例。
//
// 请求成功时，将返回策略名称。扩缩容策略在成功创建立即自动生效。
//
// ```
//
// 
//
// 规则执行的条件表达式如下所示：
//
// 
//
// ```
//
// 若 [MetricName] 是 [ComparisonOperator] [Threshold] 持续 [EvaluationPeriods] 分钟, 则 [ScalingAdjustmentType] 调整 [ScalingAdjustment]个实例。
//
// ```
//
// ```
//
// if [MetricName] ComparisonOperator [Threshold] for [EvaluationPeriods] minutes, then scaling up by/to  [ScalingAdjustment]
//
// ```
//
// 例如1：如果当前AvailableCustomCount值大于等于10，持续5分钟，扩容1台CVM实例。
//
// ```
//
// ScalingAdjustmentType = ChangeInCapacity
//
// if [AvailableGameServerSessions] >= [10] for [5] minutes, then scaling up [1]
//
// ```
//
// 例如2：如果当前AvailableGameServerSessions值大于等于200，持续5分钟，扩容到2台CVM实例。
//
// ```
//
// ScalingAdjustmentType = ExactCapacity
//
// if [AvailableGameServerSessions] >= [200] for [5] minutes, then scaling to [2]
//
// ```
//
// 例如3：如果当前AvailableCustomCount值大于等于400，持续5分钟，扩容20%台CVM实例。
//
// 当前CVM实例数为10台。扩容20%台CVM实例就是增加 10*20%等于2台
//
// ```
//
// ScalingAdjustmentType = PercentChangeInCapacity
//
// if [AvailableGameServerSessions] >= [400] for [5] minutes, then scaling by [currentCVMCount * 20 %]
//
// ```
//
// **备注1**
//
// 
//
// - | 策略名称（MetricName）                                       | 计算公式                                   | 场景                                        | 场景使用举例                                                 |
//
//   | :----------------------------------------------------------- | :----------------------------------------- | :------------------------------------------ | :----------------------------------------------------------- |
//
//   | CurrentPlayerSessions<br>当前玩家数指标                      | = 当前在线的玩家数                         | CVM随着玩家会话数变化做扩缩容。             | 例如：<br>MetricName: CurrentPlayerSessions<br>ComparisonOperator: '<=' <br>Threshold: 300<br/>EvaluationPeriods: 1<br/>ScalingAdjustment: 2<br/>ScalingAdjustment: ChangeInCapacity<br>说明：若当前CurrentPlayerSessions小于等于300，持续1分钟，则扩容2台CVM。 |
//
//   | AvailableGameServerSessions<br>可用游戏服务器会话数          | = 可用游戏服务会话数                       | CVM随着可用游戏会话数变化做扩缩容。         | 例如：<br/>MetricName: AvailableGameServerSessions<br/>ComparisonOperator: '<' <br/>Threshold: 50<br/>EvaluationPeriods: 5<br/>ScalingAdjustment: 2<br/>ScalingAdjustment: ExactCapacity<br/>说明：若当前AvailableGameServerSessions小于50，持续5分钟，则扩容到2台CVM。 |
//
//   | PercentAvailableGameServerSessions<br>可用游戏服务器会话百分比 | = 空闲游戏会话数 / 所有的游戏会话数 * 100% | CVM随着可用游戏会话数百分比变化做扩缩容。   | 例如：<br/>MetricName: PercentAvailableGameServerSessions<br/>ComparisonOperator: '<' <br/>Threshold: 50<br/>EvaluationPeriods: 1<br/>ScalingAdjustment: -30<br/>ScalingAdjustment: PercentChangeInCapacity<br/>说明：若当前PercentAvailableGameServerSessions小于50%，持续1分钟，则缩容当前实例数30%台CVM。 |
//
//   | AvailableCustomCount<br>可用客户自定义数指标                 | = 客户自定义的数                           | CVM随着可用客户自定义数变化做扩缩容。       | 例如：<br/>MetricName: AvailableCustomCount<br/>ComparisonOperator: '>=' <br/>Threshold: 6<br/>EvaluationPeriods: 3<br/>ScalingAdjustment: -1<br/>ScalingAdjustment: ExactCapacity<br/>说明：若当前AvailableCustomCount大于等于6，持续3分钟，则缩容到1台CVM。 |
//
//   | PercentAvailableCustomCount<br>可用客户自定义数百分比        | = 客户自定义数 / 客户最大自定义数* 100%    | CVM随着可用客户自定义数百分比变化做扩缩容。 | 例如：<br/>MetricName: PercentAvailableCustomCount<br/>ComparisonOperator: '<' <br/>Threshold: 15<br/>EvaluationPeriods: 3<br/>ScalingAdjustment: 1<br/>ScalingAdjustment: ChangeInCapacity<br/>说明：若当前PercentAvailableCustomCount小于15%，持续3分钟，则扩容1台CVM。 |
//
//   | ActiveInstances<br>活跃实例数指标                            | = 总实例数 - 缩容中的实例数                | CVM随着活跃实例数变化做扩缩容。             | 例如：<br/>MetricName: ActiveInstances<br/>ComparisonOperator: '<' <br/>Threshold: 3<br/>EvaluationPeriods: 1<br/>ScalingAdjustment: 3<br/>ScalingAdjustment: ExactCapacity<br/>说明：若当前ActiveInstances小于3台，持续1分钟，则扩容保留到3台CVM。 |
//
//   | IdleInstances<br>空闲实例数指标                              | = 未使用的进程数 / 每实例进程数            | CVM随着空闲实例数变化做扩缩容。             | 例如：<br/>MetricName: IdleInstances<br/>ComparisonOperator: '<' <br/>Threshold: 2<br/>EvaluationPeriods: 3<br/>ScalingAdjustment: 1<br/>ScalingAdjustment: ChangeInCapacity<br/>说明：若当前IdleInstances小于2台，持续3分钟，则扩容1台CVM。 |
//
//   | PercentIdleInstances<br>空闲实例百分比                       | = IdleInstances / ActiveInstances * 100%   | CVM随着空闲实例百分比变化做扩缩容。         | 例如：<br/>MetricName: PercentIdleInstances<br/>ComparisonOperator: '<' <br/>Threshold: 50<br/>EvaluationPeriods: 3<br/>ScalingAdjustment: 1<br/>ScalingAdjustment: ChangeInCapacity<br/>说明：若当前PercentIdleInstances小于50%，持续3分钟，则扩容1台CVM。 |
//
// 
//
// 
//
// 
//
// **备注2**
//
// 
//
// **ChangeInCapacity**
//
// 
//
//     当前CVM实例个数的扩容或缩容的调整值。正值按值扩容，负值按值缩容。
//
// 
//
// **ExactCapacity**
//
// 
//
//     把当前CVM实例个数调整为ScalingAdjustment设置的CVM实例数。
//
// 
//
// **PercentChangeInCapacity**
//
// 
//
//     按比例增加或减少的百分比。正值按比例扩容，负值按比例缩容；例如，值“-10”将按10%的比例缩容CVM实例。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PutScalingPolicy(request *PutScalingPolicyRequest) (response *PutScalingPolicyResponse, err error) {
    if request == nil {
        request = NewPutScalingPolicyRequest()
    }
    response = NewPutScalingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewPutTimerScalingPolicyRequest() (request *PutTimerScalingPolicyRequest) {
    request = &PutTimerScalingPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "PutTimerScalingPolicy")
    return
}

func NewPutTimerScalingPolicyResponse() (response *PutTimerScalingPolicyResponse) {
    response = &PutTimerScalingPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PutTimerScalingPolicy
// 本接口（PutTimerScalingPolicy）用于给fleet创建或更新定时器。
//
// 
//
// 填写字段timer_id，表示更新；不填字段timer_id表示新增。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_INSTANCELIMITEXCEEDED = "LimitExceeded.InstanceLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) PutTimerScalingPolicy(request *PutTimerScalingPolicyRequest) (response *PutTimerScalingPolicyResponse, err error) {
    if request == nil {
        request = NewPutTimerScalingPolicyRequest()
    }
    response = NewPutTimerScalingPolicyResponse()
    err = c.Send(request, response)
    return
}

func NewResolveAliasRequest() (request *ResolveAliasRequest) {
    request = &ResolveAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "ResolveAlias")
    return
}

func NewResolveAliasResponse() (response *ResolveAliasResponse) {
    response = &ResolveAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResolveAlias
// 本接口（ResolveAlias）用于获取别名当前指向的fleetId。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) ResolveAlias(request *ResolveAliasRequest) (response *ResolveAliasResponse, err error) {
    if request == nil {
        request = NewResolveAliasRequest()
    }
    response = NewResolveAliasResponse()
    err = c.Send(request, response)
    return
}

func NewSearchGameServerSessionsRequest() (request *SearchGameServerSessionsRequest) {
    request = &SearchGameServerSessionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "SearchGameServerSessions")
    return
}

func NewSearchGameServerSessionsResponse() (response *SearchGameServerSessionsResponse) {
    response = &SearchGameServerSessionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchGameServerSessions
// 本接口（SearchGameServerSessions）用于搜索游戏服务器会话列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SearchGameServerSessions(request *SearchGameServerSessionsRequest) (response *SearchGameServerSessionsResponse, err error) {
    if request == nil {
        request = NewSearchGameServerSessionsRequest()
    }
    response = NewSearchGameServerSessionsResponse()
    err = c.Send(request, response)
    return
}

func NewSetServerReservedRequest() (request *SetServerReservedRequest) {
    request = &SetServerReservedRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "SetServerReserved")
    return
}

func NewSetServerReservedResponse() (response *SetServerReservedResponse) {
    response = &SetServerReservedResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetServerReserved
// 本接口（SetServerReserved）用于将异常的实例标记为保留，用于问题排查。
//
// 
//
// 字段ReserveValue：0默认值，不保留；1 保留
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) SetServerReserved(request *SetServerReservedRequest) (response *SetServerReservedResponse, err error) {
    if request == nil {
        request = NewSetServerReservedRequest()
    }
    response = NewSetServerReservedResponse()
    err = c.Send(request, response)
    return
}

func NewSetServerWeightRequest() (request *SetServerWeightRequest) {
    request = &SetServerWeightRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "SetServerWeight")
    return
}

func NewSetServerWeightResponse() (response *SetServerWeightResponse) {
    response = &SetServerWeightResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetServerWeight
// 本接口（SetServerWeight）用于设置服务器权重。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SetServerWeight(request *SetServerWeightRequest) (response *SetServerWeightResponse, err error) {
    if request == nil {
        request = NewSetServerWeightRequest()
    }
    response = NewSetServerWeightResponse()
    err = c.Send(request, response)
    return
}

func NewStartFleetActionsRequest() (request *StartFleetActionsRequest) {
    request = &StartFleetActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "StartFleetActions")
    return
}

func NewStartFleetActionsResponse() (response *StartFleetActionsResponse) {
    response = &StartFleetActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartFleetActions
// 本接口（StartFleetActions）用于启用服务器舰队自动扩缩容。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartFleetActions(request *StartFleetActionsRequest) (response *StartFleetActionsResponse, err error) {
    if request == nil {
        request = NewStartFleetActionsRequest()
    }
    response = NewStartFleetActionsResponse()
    err = c.Send(request, response)
    return
}

func NewStartGameServerSessionPlacementRequest() (request *StartGameServerSessionPlacementRequest) {
    request = &StartGameServerSessionPlacementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "StartGameServerSessionPlacement")
    return
}

func NewStartGameServerSessionPlacementResponse() (response *StartGameServerSessionPlacementResponse) {
    response = &StartGameServerSessionPlacementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartGameServerSessionPlacement
// 本接口（StartGameServerSessionPlacement）用于开始放置游戏服务器会话。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) StartGameServerSessionPlacement(request *StartGameServerSessionPlacementRequest) (response *StartGameServerSessionPlacementResponse, err error) {
    if request == nil {
        request = NewStartGameServerSessionPlacementRequest()
    }
    response = NewStartGameServerSessionPlacementResponse()
    err = c.Send(request, response)
    return
}

func NewStopFleetActionsRequest() (request *StopFleetActionsRequest) {
    request = &StopFleetActionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "StopFleetActions")
    return
}

func NewStopFleetActionsResponse() (response *StopFleetActionsResponse) {
    response = &StopFleetActionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopFleetActions
// 本接口（StopFleetActions）用于停止服务器舰队自动扩缩容，改为手动扩缩容。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopFleetActions(request *StopFleetActionsRequest) (response *StopFleetActionsResponse, err error) {
    if request == nil {
        request = NewStopFleetActionsRequest()
    }
    response = NewStopFleetActionsResponse()
    err = c.Send(request, response)
    return
}

func NewStopGameServerSessionPlacementRequest() (request *StopGameServerSessionPlacementRequest) {
    request = &StopGameServerSessionPlacementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "StopGameServerSessionPlacement")
    return
}

func NewStopGameServerSessionPlacementResponse() (response *StopGameServerSessionPlacementResponse) {
    response = &StopGameServerSessionPlacementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopGameServerSessionPlacement
// 本接口（StopGameServerSessionPlacement）用于停止放置游戏服务器会话。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) StopGameServerSessionPlacement(request *StopGameServerSessionPlacementRequest) (response *StopGameServerSessionPlacementResponse, err error) {
    if request == nil {
        request = NewStopGameServerSessionPlacementRequest()
    }
    response = NewStopGameServerSessionPlacementResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAliasRequest() (request *UpdateAliasRequest) {
    request = &UpdateAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateAlias")
    return
}

func NewUpdateAliasResponse() (response *UpdateAliasResponse) {
    response = &UpdateAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateAlias
// 本接口（UpdateAlias）用于更新别名的属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) UpdateAlias(request *UpdateAliasRequest) (response *UpdateAliasResponse, err error) {
    if request == nil {
        request = NewUpdateAliasRequest()
    }
    response = NewUpdateAliasResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAssetRequest() (request *UpdateAssetRequest) {
    request = &UpdateAssetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateAsset")
    return
}

func NewUpdateAssetResponse() (response *UpdateAssetResponse) {
    response = &UpdateAssetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateAsset
// 本接口（UpdateAsset）用于修改生成包信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) UpdateAsset(request *UpdateAssetRequest) (response *UpdateAssetResponse, err error) {
    if request == nil {
        request = NewUpdateAssetRequest()
    }
    response = NewUpdateAssetResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateBucketAccelerateOptRequest() (request *UpdateBucketAccelerateOptRequest) {
    request = &UpdateBucketAccelerateOptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateBucketAccelerateOpt")
    return
}

func NewUpdateBucketAccelerateOptResponse() (response *UpdateBucketAccelerateOptResponse) {
    response = &UpdateBucketAccelerateOptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateBucketAccelerateOpt
// 本接口（UpdateBucketAccelerateOpt）用于开启cos全球加速。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdateBucketAccelerateOpt(request *UpdateBucketAccelerateOptRequest) (response *UpdateBucketAccelerateOptResponse, err error) {
    if request == nil {
        request = NewUpdateBucketAccelerateOptRequest()
    }
    response = NewUpdateBucketAccelerateOptResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateBucketCORSOptRequest() (request *UpdateBucketCORSOptRequest) {
    request = &UpdateBucketCORSOptRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateBucketCORSOpt")
    return
}

func NewUpdateBucketCORSOptResponse() (response *UpdateBucketCORSOptResponse) {
    response = &UpdateBucketCORSOptResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateBucketCORSOpt
// 本接口（UpdateBucketCORSOpt）用于设置cos跨域访问。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) UpdateBucketCORSOpt(request *UpdateBucketCORSOptRequest) (response *UpdateBucketCORSOptResponse, err error) {
    if request == nil {
        request = NewUpdateBucketCORSOptRequest()
    }
    response = NewUpdateBucketCORSOptResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFleetAttributesRequest() (request *UpdateFleetAttributesRequest) {
    request = &UpdateFleetAttributesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateFleetAttributes")
    return
}

func NewUpdateFleetAttributesResponse() (response *UpdateFleetAttributesResponse) {
    response = &UpdateFleetAttributesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateFleetAttributes
// 本接口（UpdateFleetAttributes）用于更新服务器舰队属性。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateFleetAttributes(request *UpdateFleetAttributesRequest) (response *UpdateFleetAttributesResponse, err error) {
    if request == nil {
        request = NewUpdateFleetAttributesRequest()
    }
    response = NewUpdateFleetAttributesResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFleetCapacityRequest() (request *UpdateFleetCapacityRequest) {
    request = &UpdateFleetCapacityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateFleetCapacity")
    return
}

func NewUpdateFleetCapacityResponse() (response *UpdateFleetCapacityResponse) {
    response = &UpdateFleetCapacityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateFleetCapacity
// 本接口（UpdateFleetCapacity）用于更新服务器舰队容量配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_FLEETLIMITEXCEEDED = "LimitExceeded.FleetLimitExceeded"
//  LIMITEXCEEDED_INSTANCELIMITEXCEEDED = "LimitExceeded.InstanceLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) UpdateFleetCapacity(request *UpdateFleetCapacityRequest) (response *UpdateFleetCapacityResponse, err error) {
    if request == nil {
        request = NewUpdateFleetCapacityRequest()
    }
    response = NewUpdateFleetCapacityResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFleetNameRequest() (request *UpdateFleetNameRequest) {
    request = &UpdateFleetNameRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateFleetName")
    return
}

func NewUpdateFleetNameResponse() (response *UpdateFleetNameResponse) {
    response = &UpdateFleetNameResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateFleetName
// 本接口（UpdateFleetName）用于更新服务器舰队名称。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateFleetName(request *UpdateFleetNameRequest) (response *UpdateFleetNameResponse, err error) {
    if request == nil {
        request = NewUpdateFleetNameRequest()
    }
    response = NewUpdateFleetNameResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateFleetPortSettingsRequest() (request *UpdateFleetPortSettingsRequest) {
    request = &UpdateFleetPortSettingsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateFleetPortSettings")
    return
}

func NewUpdateFleetPortSettingsResponse() (response *UpdateFleetPortSettingsResponse) {
    response = &UpdateFleetPortSettingsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateFleetPortSettings
// 本接口（UpdateFleetPortSettings）用于更新服务器舰队安全组。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateFleetPortSettings(request *UpdateFleetPortSettingsRequest) (response *UpdateFleetPortSettingsResponse, err error) {
    if request == nil {
        request = NewUpdateFleetPortSettingsRequest()
    }
    response = NewUpdateFleetPortSettingsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGameServerSessionRequest() (request *UpdateGameServerSessionRequest) {
    request = &UpdateGameServerSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateGameServerSession")
    return
}

func NewUpdateGameServerSessionResponse() (response *UpdateGameServerSessionResponse) {
    response = &UpdateGameServerSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateGameServerSession
// 本接口（UpdateGameServerSession）用于更新游戏服务器会话。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) UpdateGameServerSession(request *UpdateGameServerSessionRequest) (response *UpdateGameServerSessionResponse, err error) {
    if request == nil {
        request = NewUpdateGameServerSessionRequest()
    }
    response = NewUpdateGameServerSessionResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateGameServerSessionQueueRequest() (request *UpdateGameServerSessionQueueRequest) {
    request = &UpdateGameServerSessionQueueRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateGameServerSessionQueue")
    return
}

func NewUpdateGameServerSessionQueueResponse() (response *UpdateGameServerSessionQueueResponse) {
    response = &UpdateGameServerSessionQueueResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateGameServerSessionQueue
// 本接口（UpdateGameServerSessionQueue）用于修改游戏服务器会话队列。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
func (c *Client) UpdateGameServerSessionQueue(request *UpdateGameServerSessionQueueRequest) (response *UpdateGameServerSessionQueueResponse, err error) {
    if request == nil {
        request = NewUpdateGameServerSessionQueueRequest()
    }
    response = NewUpdateGameServerSessionQueueResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateRuntimeConfigurationRequest() (request *UpdateRuntimeConfigurationRequest) {
    request = &UpdateRuntimeConfigurationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gse", APIVersion, "UpdateRuntimeConfiguration")
    return
}

func NewUpdateRuntimeConfigurationResponse() (response *UpdateRuntimeConfigurationResponse) {
    response = &UpdateRuntimeConfigurationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateRuntimeConfiguration
// 本接口（UpdateRuntimeConfiguration）用于更新服务器舰队配置。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SERVICENOTOPENED = "UnauthorizedOperation.ServiceNotOpened"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateRuntimeConfiguration(request *UpdateRuntimeConfigurationRequest) (response *UpdateRuntimeConfigurationResponse, err error) {
    if request == nil {
        request = NewUpdateRuntimeConfigurationRequest()
    }
    response = NewUpdateRuntimeConfigurationResponse()
    err = c.Send(request, response)
    return
}
