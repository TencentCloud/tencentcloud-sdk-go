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

package v20190916

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-09-16"

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


func NewBatchDescribeKTVMusicDetailsRequest() (request *BatchDescribeKTVMusicDetailsRequest) {
    request = &BatchDescribeKTVMusicDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "BatchDescribeKTVMusicDetails")
    
    
    return
}

func NewBatchDescribeKTVMusicDetailsResponse() (response *BatchDescribeKTVMusicDetailsResponse) {
    response = &BatchDescribeKTVMusicDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchDescribeKTVMusicDetails
// 根据 Id 列表查询歌曲的详细信息，包含基础信息及播放信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) BatchDescribeKTVMusicDetails(request *BatchDescribeKTVMusicDetailsRequest) (response *BatchDescribeKTVMusicDetailsResponse, err error) {
    return c.BatchDescribeKTVMusicDetailsWithContext(context.Background(), request)
}

// BatchDescribeKTVMusicDetails
// 根据 Id 列表查询歌曲的详细信息，包含基础信息及播放信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) BatchDescribeKTVMusicDetailsWithContext(ctx context.Context, request *BatchDescribeKTVMusicDetailsRequest) (response *BatchDescribeKTVMusicDetailsResponse, err error) {
    if request == nil {
        request = NewBatchDescribeKTVMusicDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BatchDescribeKTVMusicDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewBatchDescribeKTVMusicDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateKTVRobotRequest() (request *CreateKTVRobotRequest) {
    request = &CreateKTVRobotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "CreateKTVRobot")
    
    
    return
}

func NewCreateKTVRobotResponse() (response *CreateKTVRobotResponse) {
    response = &CreateKTVRobotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateKTVRobot
// 创建机器人，支持进入 RTC 房间，播放直播互动曲库歌曲。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateKTVRobot(request *CreateKTVRobotRequest) (response *CreateKTVRobotResponse, err error) {
    return c.CreateKTVRobotWithContext(context.Background(), request)
}

// CreateKTVRobot
// 创建机器人，支持进入 RTC 房间，播放直播互动曲库歌曲。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateKTVRobotWithContext(ctx context.Context, request *CreateKTVRobotRequest) (response *CreateKTVRobotResponse, err error) {
    if request == nil {
        request = NewCreateKTVRobotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateKTVRobot require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateKTVRobotResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuthInfoRequest() (request *DescribeAuthInfoRequest) {
    request = &DescribeAuthInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribeAuthInfo")
    
    
    return
}

func NewDescribeAuthInfoResponse() (response *DescribeAuthInfoResponse) {
    response = &DescribeAuthInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAuthInfo
// 获取授权项目信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAuthInfo(request *DescribeAuthInfoRequest) (response *DescribeAuthInfoResponse, err error) {
    return c.DescribeAuthInfoWithContext(context.Background(), request)
}

// DescribeAuthInfo
// 获取授权项目信息列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeAuthInfoWithContext(ctx context.Context, request *DescribeAuthInfoRequest) (response *DescribeAuthInfoResponse, err error) {
    if request == nil {
        request = NewDescribeAuthInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAuthInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAuthInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudMusicRequest() (request *DescribeCloudMusicRequest) {
    request = &DescribeCloudMusicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribeCloudMusic")
    
    
    return
}

func NewDescribeCloudMusicResponse() (response *DescribeCloudMusicResponse) {
    response = &DescribeCloudMusicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudMusic
// 获取云音乐播放信息接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCloudMusic(request *DescribeCloudMusicRequest) (response *DescribeCloudMusicResponse, err error) {
    return c.DescribeCloudMusicWithContext(context.Background(), request)
}

// DescribeCloudMusic
// 获取云音乐播放信息接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeCloudMusicWithContext(ctx context.Context, request *DescribeCloudMusicRequest) (response *DescribeCloudMusicResponse, err error) {
    if request == nil {
        request = NewDescribeCloudMusicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudMusic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudMusicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudMusicPurchasedRequest() (request *DescribeCloudMusicPurchasedRequest) {
    request = &DescribeCloudMusicPurchasedRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribeCloudMusicPurchased")
    
    
    return
}

func NewDescribeCloudMusicPurchasedResponse() (response *DescribeCloudMusicPurchasedResponse) {
    response = &DescribeCloudMusicPurchasedResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCloudMusicPurchased
// 获取授权项目下已购云音乐列表
//
// 可能返回的错误码:
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
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudMusicPurchased(request *DescribeCloudMusicPurchasedRequest) (response *DescribeCloudMusicPurchasedResponse, err error) {
    return c.DescribeCloudMusicPurchasedWithContext(context.Background(), request)
}

// DescribeCloudMusicPurchased
// 获取授权项目下已购云音乐列表
//
// 可能返回的错误码:
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
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeCloudMusicPurchasedWithContext(ctx context.Context, request *DescribeCloudMusicPurchasedRequest) (response *DescribeCloudMusicPurchasedResponse, err error) {
    if request == nil {
        request = NewDescribeCloudMusicPurchasedRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudMusicPurchased require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudMusicPurchasedResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeItemByIdRequest() (request *DescribeItemByIdRequest) {
    request = &DescribeItemByIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribeItemById")
    
    
    return
}

func NewDescribeItemByIdResponse() (response *DescribeItemByIdResponse) {
    response = &DescribeItemByIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeItemById
// 根据歌曲ID查询歌曲信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeItemById(request *DescribeItemByIdRequest) (response *DescribeItemByIdResponse, err error) {
    return c.DescribeItemByIdWithContext(context.Background(), request)
}

// DescribeItemById
// 根据歌曲ID查询歌曲信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeItemByIdWithContext(ctx context.Context, request *DescribeItemByIdRequest) (response *DescribeItemByIdResponse, err error) {
    if request == nil {
        request = NewDescribeItemByIdRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeItemById require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeItemByIdResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeItemsRequest() (request *DescribeItemsRequest) {
    request = &DescribeItemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribeItems")
    
    
    return
}

func NewDescribeItemsResponse() (response *DescribeItemsResponse) {
    response = &DescribeItemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeItems
// 该服务后续会停用，不再建议使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeItems(request *DescribeItemsRequest) (response *DescribeItemsResponse, err error) {
    return c.DescribeItemsWithContext(context.Background(), request)
}

// DescribeItems
// 该服务后续会停用，不再建议使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeItemsWithContext(ctx context.Context, request *DescribeItemsRequest) (response *DescribeItemsResponse, err error) {
    if request == nil {
        request = NewDescribeItemsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeItems require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeItemsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKTVMusicDetailRequest() (request *DescribeKTVMusicDetailRequest) {
    request = &DescribeKTVMusicDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribeKTVMusicDetail")
    
    
    return
}

func NewDescribeKTVMusicDetailResponse() (response *DescribeKTVMusicDetailResponse) {
    response = &DescribeKTVMusicDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKTVMusicDetail
// 根据 Id 查询歌曲的详细信息，包含基础信息及播放信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKTVMusicDetail(request *DescribeKTVMusicDetailRequest) (response *DescribeKTVMusicDetailResponse, err error) {
    return c.DescribeKTVMusicDetailWithContext(context.Background(), request)
}

// DescribeKTVMusicDetail
// 根据 Id 查询歌曲的详细信息，包含基础信息及播放信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKTVMusicDetailWithContext(ctx context.Context, request *DescribeKTVMusicDetailRequest) (response *DescribeKTVMusicDetailResponse, err error) {
    if request == nil {
        request = NewDescribeKTVMusicDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKTVMusicDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKTVMusicDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKTVMusicTagsRequest() (request *DescribeKTVMusicTagsRequest) {
    request = &DescribeKTVMusicTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribeKTVMusicTags")
    
    
    return
}

func NewDescribeKTVMusicTagsResponse() (response *DescribeKTVMusicTagsResponse) {
    response = &DescribeKTVMusicTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKTVMusicTags
// 获取直播互动曲库标签分组信息和标签信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKTVMusicTags(request *DescribeKTVMusicTagsRequest) (response *DescribeKTVMusicTagsResponse, err error) {
    return c.DescribeKTVMusicTagsWithContext(context.Background(), request)
}

// DescribeKTVMusicTags
// 获取直播互动曲库标签分组信息和标签信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKTVMusicTagsWithContext(ctx context.Context, request *DescribeKTVMusicTagsRequest) (response *DescribeKTVMusicTagsResponse, err error) {
    if request == nil {
        request = NewDescribeKTVMusicTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKTVMusicTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKTVMusicTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKTVPlaylistDetailRequest() (request *DescribeKTVPlaylistDetailRequest) {
    request = &DescribeKTVPlaylistDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribeKTVPlaylistDetail")
    
    
    return
}

func NewDescribeKTVPlaylistDetailResponse() (response *DescribeKTVPlaylistDetailResponse) {
    response = &DescribeKTVPlaylistDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKTVPlaylistDetail
// 根据歌单 Id 获取歌单详情，包括歌单的基础信息以及歌曲列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKTVPlaylistDetail(request *DescribeKTVPlaylistDetailRequest) (response *DescribeKTVPlaylistDetailResponse, err error) {
    return c.DescribeKTVPlaylistDetailWithContext(context.Background(), request)
}

// DescribeKTVPlaylistDetail
// 根据歌单 Id 获取歌单详情，包括歌单的基础信息以及歌曲列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKTVPlaylistDetailWithContext(ctx context.Context, request *DescribeKTVPlaylistDetailRequest) (response *DescribeKTVPlaylistDetailResponse, err error) {
    if request == nil {
        request = NewDescribeKTVPlaylistDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKTVPlaylistDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKTVPlaylistDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKTVPlaylistsRequest() (request *DescribeKTVPlaylistsRequest) {
    request = &DescribeKTVPlaylistsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribeKTVPlaylists")
    
    
    return
}

func NewDescribeKTVPlaylistsResponse() (response *DescribeKTVPlaylistsResponse) {
    response = &DescribeKTVPlaylistsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKTVPlaylists
// 获取直播互动曲库推荐歌单列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKTVPlaylists(request *DescribeKTVPlaylistsRequest) (response *DescribeKTVPlaylistsResponse, err error) {
    return c.DescribeKTVPlaylistsWithContext(context.Background(), request)
}

// DescribeKTVPlaylists
// 获取直播互动曲库推荐歌单列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKTVPlaylistsWithContext(ctx context.Context, request *DescribeKTVPlaylistsRequest) (response *DescribeKTVPlaylistsResponse, err error) {
    if request == nil {
        request = NewDescribeKTVPlaylistsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKTVPlaylists require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKTVPlaylistsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKTVRobotsRequest() (request *DescribeKTVRobotsRequest) {
    request = &DescribeKTVRobotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribeKTVRobots")
    
    
    return
}

func NewDescribeKTVRobotsResponse() (response *DescribeKTVRobotsResponse) {
    response = &DescribeKTVRobotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKTVRobots
// 获取机器人列表，支持 Id、状态等过滤条件。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeKTVRobots(request *DescribeKTVRobotsRequest) (response *DescribeKTVRobotsResponse, err error) {
    return c.DescribeKTVRobotsWithContext(context.Background(), request)
}

// DescribeKTVRobots
// 获取机器人列表，支持 Id、状态等过滤条件。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeKTVRobotsWithContext(ctx context.Context, request *DescribeKTVRobotsRequest) (response *DescribeKTVRobotsResponse, err error) {
    if request == nil {
        request = NewDescribeKTVRobotsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKTVRobots require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKTVRobotsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKTVSingerCategoriesRequest() (request *DescribeKTVSingerCategoriesRequest) {
    request = &DescribeKTVSingerCategoriesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribeKTVSingerCategories")
    
    
    return
}

func NewDescribeKTVSingerCategoriesResponse() (response *DescribeKTVSingerCategoriesResponse) {
    response = &DescribeKTVSingerCategoriesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKTVSingerCategories
// 获取直播互动曲库歌手分类信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKTVSingerCategories(request *DescribeKTVSingerCategoriesRequest) (response *DescribeKTVSingerCategoriesResponse, err error) {
    return c.DescribeKTVSingerCategoriesWithContext(context.Background(), request)
}

// DescribeKTVSingerCategories
// 获取直播互动曲库歌手分类信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKTVSingerCategoriesWithContext(ctx context.Context, request *DescribeKTVSingerCategoriesRequest) (response *DescribeKTVSingerCategoriesResponse, err error) {
    if request == nil {
        request = NewDescribeKTVSingerCategoriesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKTVSingerCategories require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKTVSingerCategoriesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKTVSingerMusicsRequest() (request *DescribeKTVSingerMusicsRequest) {
    request = &DescribeKTVSingerMusicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribeKTVSingerMusics")
    
    
    return
}

func NewDescribeKTVSingerMusicsResponse() (response *DescribeKTVSingerMusicsResponse) {
    response = &DescribeKTVSingerMusicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKTVSingerMusics
// 根据歌手id，返回该歌手下歌曲列表。
//
// 
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKTVSingerMusics(request *DescribeKTVSingerMusicsRequest) (response *DescribeKTVSingerMusicsResponse, err error) {
    return c.DescribeKTVSingerMusicsWithContext(context.Background(), request)
}

// DescribeKTVSingerMusics
// 根据歌手id，返回该歌手下歌曲列表。
//
// 
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKTVSingerMusicsWithContext(ctx context.Context, request *DescribeKTVSingerMusicsRequest) (response *DescribeKTVSingerMusicsResponse, err error) {
    if request == nil {
        request = NewDescribeKTVSingerMusicsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKTVSingerMusics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKTVSingerMusicsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKTVSingersRequest() (request *DescribeKTVSingersRequest) {
    request = &DescribeKTVSingersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribeKTVSingers")
    
    
    return
}

func NewDescribeKTVSingersResponse() (response *DescribeKTVSingersResponse) {
    response = &DescribeKTVSingersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKTVSingers
// 根据过滤条件，返回匹配的歌手列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKTVSingers(request *DescribeKTVSingersRequest) (response *DescribeKTVSingersResponse, err error) {
    return c.DescribeKTVSingersWithContext(context.Background(), request)
}

// DescribeKTVSingers
// 根据过滤条件，返回匹配的歌手列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKTVSingersWithContext(ctx context.Context, request *DescribeKTVSingersRequest) (response *DescribeKTVSingersResponse, err error) {
    if request == nil {
        request = NewDescribeKTVSingersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKTVSingers require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKTVSingersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKTVSuggestionsRequest() (request *DescribeKTVSuggestionsRequest) {
    request = &DescribeKTVSuggestionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribeKTVSuggestions")
    
    
    return
}

func NewDescribeKTVSuggestionsResponse() (response *DescribeKTVSuggestionsResponse) {
    response = &DescribeKTVSuggestionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKTVSuggestions
// 获取直播互动曲库联想词
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeKTVSuggestions(request *DescribeKTVSuggestionsRequest) (response *DescribeKTVSuggestionsResponse, err error) {
    return c.DescribeKTVSuggestionsWithContext(context.Background(), request)
}

// DescribeKTVSuggestions
// 获取直播互动曲库联想词
//
// 可能返回的错误码:
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeKTVSuggestionsWithContext(ctx context.Context, request *DescribeKTVSuggestionsRequest) (response *DescribeKTVSuggestionsResponse, err error) {
    if request == nil {
        request = NewDescribeKTVSuggestionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKTVSuggestions require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKTVSuggestionsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKTVTopListRequest() (request *DescribeKTVTopListRequest) {
    request = &DescribeKTVTopListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribeKTVTopList")
    
    
    return
}

func NewDescribeKTVTopListResponse() (response *DescribeKTVTopListResponse) {
    response = &DescribeKTVTopListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKTVTopList
// 获取直播互动曲库歌曲的周榜和月榜
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeKTVTopList(request *DescribeKTVTopListRequest) (response *DescribeKTVTopListResponse, err error) {
    return c.DescribeKTVTopListWithContext(context.Background(), request)
}

// DescribeKTVTopList
// 获取直播互动曲库歌曲的周榜和月榜
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeKTVTopListWithContext(ctx context.Context, request *DescribeKTVTopListRequest) (response *DescribeKTVTopListResponse, err error) {
    if request == nil {
        request = NewDescribeKTVTopListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKTVTopList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKTVTopListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLyricRequest() (request *DescribeLyricRequest) {
    request = &DescribeLyricRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribeLyric")
    
    
    return
}

func NewDescribeLyricResponse() (response *DescribeLyricResponse) {
    response = &DescribeLyricResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLyric
// 根据接口的模式及歌曲ID来取得歌词信息或者波形图信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLyric(request *DescribeLyricRequest) (response *DescribeLyricResponse, err error) {
    return c.DescribeLyricWithContext(context.Background(), request)
}

// DescribeLyric
// 根据接口的模式及歌曲ID来取得歌词信息或者波形图信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeLyricWithContext(ctx context.Context, request *DescribeLyricRequest) (response *DescribeLyricResponse, err error) {
    if request == nil {
        request = NewDescribeLyricRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLyric require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLyricResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMusicRequest() (request *DescribeMusicRequest) {
    request = &DescribeMusicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribeMusic")
    
    
    return
}

func NewDescribeMusicResponse() (response *DescribeMusicResponse) {
    response = &DescribeMusicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMusic
// 获取曲库包歌曲播放信息接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMusic(request *DescribeMusicRequest) (response *DescribeMusicResponse, err error) {
    return c.DescribeMusicWithContext(context.Background(), request)
}

// DescribeMusic
// 获取曲库包歌曲播放信息接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMusicWithContext(ctx context.Context, request *DescribeMusicRequest) (response *DescribeMusicResponse, err error) {
    if request == nil {
        request = NewDescribeMusicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMusic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMusicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMusicSaleStatusRequest() (request *DescribeMusicSaleStatusRequest) {
    request = &DescribeMusicSaleStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribeMusicSaleStatus")
    
    
    return
}

func NewDescribeMusicSaleStatusResponse() (response *DescribeMusicSaleStatusResponse) {
    response = &DescribeMusicSaleStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMusicSaleStatus
// 根据音乐信息查询音乐是否在售
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeMusicSaleStatus(request *DescribeMusicSaleStatusRequest) (response *DescribeMusicSaleStatusResponse, err error) {
    return c.DescribeMusicSaleStatusWithContext(context.Background(), request)
}

// DescribeMusicSaleStatus
// 根据音乐信息查询音乐是否在售
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeMusicSaleStatusWithContext(ctx context.Context, request *DescribeMusicSaleStatusRequest) (response *DescribeMusicSaleStatusResponse, err error) {
    if request == nil {
        request = NewDescribeMusicSaleStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMusicSaleStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMusicSaleStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePackageItemsRequest() (request *DescribePackageItemsRequest) {
    request = &DescribePackageItemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribePackageItems")
    
    
    return
}

func NewDescribePackageItemsResponse() (response *DescribePackageItemsResponse) {
    response = &DescribePackageItemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePackageItems
// 获取曲库包下已核销歌曲列表接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePackageItems(request *DescribePackageItemsRequest) (response *DescribePackageItemsResponse, err error) {
    return c.DescribePackageItemsWithContext(context.Background(), request)
}

// DescribePackageItems
// 获取曲库包下已核销歌曲列表接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePackageItemsWithContext(ctx context.Context, request *DescribePackageItemsRequest) (response *DescribePackageItemsResponse, err error) {
    if request == nil {
        request = NewDescribePackageItemsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePackageItems require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePackageItemsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePackagesRequest() (request *DescribePackagesRequest) {
    request = &DescribePackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribePackages")
    
    
    return
}

func NewDescribePackagesResponse() (response *DescribePackagesResponse) {
    response = &DescribePackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePackages
// 获取已购曲库包列表接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePackages(request *DescribePackagesRequest) (response *DescribePackagesResponse, err error) {
    return c.DescribePackagesWithContext(context.Background(), request)
}

// DescribePackages
// 获取已购曲库包列表接口
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePackagesWithContext(ctx context.Context, request *DescribePackagesRequest) (response *DescribePackagesResponse, err error) {
    if request == nil {
        request = NewDescribePackagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePackages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePackagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePkgOfflineMusicRequest() (request *DescribePkgOfflineMusicRequest) {
    request = &DescribePkgOfflineMusicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribePkgOfflineMusic")
    
    
    return
}

func NewDescribePkgOfflineMusicResponse() (response *DescribePkgOfflineMusicResponse) {
    response = &DescribePkgOfflineMusicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePkgOfflineMusic
// 根据购买曲库包用户可查询已回退的歌曲信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePkgOfflineMusic(request *DescribePkgOfflineMusicRequest) (response *DescribePkgOfflineMusicResponse, err error) {
    return c.DescribePkgOfflineMusicWithContext(context.Background(), request)
}

// DescribePkgOfflineMusic
// 根据购买曲库包用户可查询已回退的歌曲信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribePkgOfflineMusicWithContext(ctx context.Context, request *DescribePkgOfflineMusicRequest) (response *DescribePkgOfflineMusicResponse, err error) {
    if request == nil {
        request = NewDescribePkgOfflineMusicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePkgOfflineMusic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePkgOfflineMusicResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStationsRequest() (request *DescribeStationsRequest) {
    request = &DescribeStationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DescribeStations")
    
    
    return
}

func NewDescribeStationsResponse() (response *DescribeStationsResponse) {
    response = &DescribeStationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStations
// 该服务后续会停用，不再建议使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeStations(request *DescribeStationsRequest) (response *DescribeStationsResponse, err error) {
    return c.DescribeStationsWithContext(context.Background(), request)
}

// DescribeStations
// 该服务后续会停用，不再建议使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeStationsWithContext(ctx context.Context, request *DescribeStationsRequest) (response *DescribeStationsResponse, err error) {
    if request == nil {
        request = NewDescribeStationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStations require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStationsResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyKTVRobotRequest() (request *DestroyKTVRobotRequest) {
    request = &DestroyKTVRobotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "DestroyKTVRobot")
    
    
    return
}

func NewDestroyKTVRobotResponse() (response *DestroyKTVRobotResponse) {
    response = &DestroyKTVRobotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DestroyKTVRobot
// 销毁机器人，机器人退出 RTC 房间。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DestroyKTVRobot(request *DestroyKTVRobotRequest) (response *DestroyKTVRobotResponse, err error) {
    return c.DestroyKTVRobotWithContext(context.Background(), request)
}

// DestroyKTVRobot
// 销毁机器人，机器人退出 RTC 房间。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DestroyKTVRobotWithContext(ctx context.Context, request *DestroyKTVRobotRequest) (response *DestroyKTVRobotResponse, err error) {
    if request == nil {
        request = NewDestroyKTVRobotRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DestroyKTVRobot require credential")
    }

    request.SetContext(ctx)
    
    response = NewDestroyKTVRobotResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMusicOnShelvesRequest() (request *ModifyMusicOnShelvesRequest) {
    request = &ModifyMusicOnShelvesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "ModifyMusicOnShelves")
    
    
    return
}

func NewModifyMusicOnShelvesResponse() (response *ModifyMusicOnShelvesResponse) {
    response = &ModifyMusicOnShelvesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyMusicOnShelves
// 根据资源方，需要变更的参数，请求该接口进行变更，为空的参数默认为无变更
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyMusicOnShelves(request *ModifyMusicOnShelvesRequest) (response *ModifyMusicOnShelvesResponse, err error) {
    return c.ModifyMusicOnShelvesWithContext(context.Background(), request)
}

// ModifyMusicOnShelves
// 根据资源方，需要变更的参数，请求该接口进行变更，为空的参数默认为无变更
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyMusicOnShelvesWithContext(ctx context.Context, request *ModifyMusicOnShelvesRequest) (response *ModifyMusicOnShelvesResponse, err error) {
    if request == nil {
        request = NewModifyMusicOnShelvesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyMusicOnShelves require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyMusicOnShelvesResponse()
    err = c.Send(request, response)
    return
}

func NewPutMusicOnTheShelvesRequest() (request *PutMusicOnTheShelvesRequest) {
    request = &PutMusicOnTheShelvesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "PutMusicOnTheShelves")
    
    
    return
}

func NewPutMusicOnTheShelvesResponse() (response *PutMusicOnTheShelvesResponse) {
    response = &PutMusicOnTheShelvesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PutMusicOnTheShelves
// 根据资源方所传歌曲信息，进行歌曲上架，多个歌曲同时请求时，需构造复合结构进行请求
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PutMusicOnTheShelves(request *PutMusicOnTheShelvesRequest) (response *PutMusicOnTheShelvesResponse, err error) {
    return c.PutMusicOnTheShelvesWithContext(context.Background(), request)
}

// PutMusicOnTheShelves
// 根据资源方所传歌曲信息，进行歌曲上架，多个歌曲同时请求时，需构造复合结构进行请求
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) PutMusicOnTheShelvesWithContext(ctx context.Context, request *PutMusicOnTheShelvesRequest) (response *PutMusicOnTheShelvesResponse, err error) {
    if request == nil {
        request = NewPutMusicOnTheShelvesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PutMusicOnTheShelves require credential")
    }

    request.SetContext(ctx)
    
    response = NewPutMusicOnTheShelvesResponse()
    err = c.Send(request, response)
    return
}

func NewReportDataRequest() (request *ReportDataRequest) {
    request = &ReportDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "ReportData")
    
    
    return
}

func NewReportDataResponse() (response *ReportDataResponse) {
    response = &ReportDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReportData
// 客户上报用户数据功能，为了更好地为用户提供优质服务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReportData(request *ReportDataRequest) (response *ReportDataResponse, err error) {
    return c.ReportDataWithContext(context.Background(), request)
}

// ReportData
// 客户上报用户数据功能，为了更好地为用户提供优质服务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ReportDataWithContext(ctx context.Context, request *ReportDataRequest) (response *ReportDataResponse, err error) {
    if request == nil {
        request = NewReportDataRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReportData require credential")
    }

    request.SetContext(ctx)
    
    response = NewReportDataResponse()
    err = c.Send(request, response)
    return
}

func NewSearchKTVMusicsRequest() (request *SearchKTVMusicsRequest) {
    request = &SearchKTVMusicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "SearchKTVMusics")
    
    
    return
}

func NewSearchKTVMusicsResponse() (response *SearchKTVMusicsResponse) {
    response = &SearchKTVMusicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchKTVMusics
// 根据搜索条件，返回匹配的歌曲列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SearchKTVMusics(request *SearchKTVMusicsRequest) (response *SearchKTVMusicsResponse, err error) {
    return c.SearchKTVMusicsWithContext(context.Background(), request)
}

// SearchKTVMusics
// 根据搜索条件，返回匹配的歌曲列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SearchKTVMusicsWithContext(ctx context.Context, request *SearchKTVMusicsRequest) (response *SearchKTVMusicsResponse, err error) {
    if request == nil {
        request = NewSearchKTVMusicsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchKTVMusics require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchKTVMusicsResponse()
    err = c.Send(request, response)
    return
}

func NewSyncKTVRobotCommandRequest() (request *SyncKTVRobotCommandRequest) {
    request = &SyncKTVRobotCommandRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "SyncKTVRobotCommand")
    
    
    return
}

func NewSyncKTVRobotCommandResponse() (response *SyncKTVRobotCommandResponse) {
    response = &SyncKTVRobotCommandResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SyncKTVRobotCommand
// 下发操作机器人指令，支持播放、暂停、恢复、歌单设置等操作指令，实现对机器人行为的控制。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) SyncKTVRobotCommand(request *SyncKTVRobotCommandRequest) (response *SyncKTVRobotCommandResponse, err error) {
    return c.SyncKTVRobotCommandWithContext(context.Background(), request)
}

// SyncKTVRobotCommand
// 下发操作机器人指令，支持播放、暂停、恢复、歌单设置等操作指令，实现对机器人行为的控制。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) SyncKTVRobotCommandWithContext(ctx context.Context, request *SyncKTVRobotCommandRequest) (response *SyncKTVRobotCommandResponse, err error) {
    if request == nil {
        request = NewSyncKTVRobotCommandRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SyncKTVRobotCommand require credential")
    }

    request.SetContext(ctx)
    
    response = NewSyncKTVRobotCommandResponse()
    err = c.Send(request, response)
    return
}

func NewTakeMusicOffShelvesRequest() (request *TakeMusicOffShelvesRequest) {
    request = &TakeMusicOffShelvesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ame", APIVersion, "TakeMusicOffShelves")
    
    
    return
}

func NewTakeMusicOffShelvesResponse() (response *TakeMusicOffShelvesResponse) {
    response = &TakeMusicOffShelvesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TakeMusicOffShelves
// 根据资源方所传MusicId进行将歌曲进行下架，多个MusicId使用逗号隔开
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) TakeMusicOffShelves(request *TakeMusicOffShelvesRequest) (response *TakeMusicOffShelvesResponse, err error) {
    return c.TakeMusicOffShelvesWithContext(context.Background(), request)
}

// TakeMusicOffShelves
// 根据资源方所传MusicId进行将歌曲进行下架，多个MusicId使用逗号隔开
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) TakeMusicOffShelvesWithContext(ctx context.Context, request *TakeMusicOffShelvesRequest) (response *TakeMusicOffShelvesResponse, err error) {
    if request == nil {
        request = NewTakeMusicOffShelvesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TakeMusicOffShelves require credential")
    }

    request.SetContext(ctx)
    
    response = NewTakeMusicOffShelvesResponse()
    err = c.Send(request, response)
    return
}
