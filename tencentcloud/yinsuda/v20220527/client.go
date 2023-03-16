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

package v20220527

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-05-27"

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


func NewApplyChorusRequest() (request *ApplyChorusRequest) {
    request = &ApplyChorusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yinsuda", APIVersion, "ApplyChorus")
    
    
    return
}

func NewApplyChorusResponse() (response *ApplyChorusResponse) {
    response = &ApplyChorusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ApplyChorus
// 申请合唱相关信息，用于标记用户的演唱是在合唱场景下。
func (c *Client) ApplyChorus(request *ApplyChorusRequest) (response *ApplyChorusResponse, err error) {
    return c.ApplyChorusWithContext(context.Background(), request)
}

// ApplyChorus
// 申请合唱相关信息，用于标记用户的演唱是在合唱场景下。
func (c *Client) ApplyChorusWithContext(ctx context.Context, request *ApplyChorusRequest) (response *ApplyChorusResponse, err error) {
    if request == nil {
        request = NewApplyChorusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyChorus require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyChorusResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDescribeKTVMusicDetailsRequest() (request *BatchDescribeKTVMusicDetailsRequest) {
    request = &BatchDescribeKTVMusicDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yinsuda", APIVersion, "BatchDescribeKTVMusicDetails")
    
    
    return
}

func NewBatchDescribeKTVMusicDetailsResponse() (response *BatchDescribeKTVMusicDetailsResponse) {
    response = &BatchDescribeKTVMusicDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchDescribeKTVMusicDetails
// 批量获取歌曲详细信息，包括：歌词下载链接、播放凭证、音高数据下载链接信息等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_USERLIVEVIPTIMEEXPIRE = "FailedOperation.UserLiveVipTimeExpire"
//  FAILEDOPERATION_USERNOTLIVEVIP = "FailedOperation.UserNotLiveVip"
func (c *Client) BatchDescribeKTVMusicDetails(request *BatchDescribeKTVMusicDetailsRequest) (response *BatchDescribeKTVMusicDetailsResponse, err error) {
    return c.BatchDescribeKTVMusicDetailsWithContext(context.Background(), request)
}

// BatchDescribeKTVMusicDetails
// 批量获取歌曲详细信息，包括：歌词下载链接、播放凭证、音高数据下载链接信息等。
//
// 可能返回的错误码:
//  FAILEDOPERATION_USERLIVEVIPTIMEEXPIRE = "FailedOperation.UserLiveVipTimeExpire"
//  FAILEDOPERATION_USERNOTLIVEVIP = "FailedOperation.UserNotLiveVip"
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
    
    request.Init().WithApiInfo("yinsuda", APIVersion, "CreateKTVRobot")
    
    
    return
}

func NewCreateKTVRobotResponse() (response *CreateKTVRobotResponse) {
    response = &CreateKTVRobotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateKTVRobot
// 创建机器人，支持进入 RTC 房间，播放曲库歌曲。
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
// 创建机器人，支持进入 RTC 房间，播放曲库歌曲。
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

func NewDescribeKTVMatchMusicsRequest() (request *DescribeKTVMatchMusicsRequest) {
    request = &DescribeKTVMatchMusicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yinsuda", APIVersion, "DescribeKTVMatchMusics")
    
    
    return
}

func NewDescribeKTVMatchMusicsResponse() (response *DescribeKTVMatchMusicsResponse) {
    response = &DescribeKTVMatchMusicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKTVMatchMusics
// 根据输入的规则匹配曲库中的歌曲。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeKTVMatchMusics(request *DescribeKTVMatchMusicsRequest) (response *DescribeKTVMatchMusicsResponse, err error) {
    return c.DescribeKTVMatchMusicsWithContext(context.Background(), request)
}

// DescribeKTVMatchMusics
// 根据输入的规则匹配曲库中的歌曲。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeKTVMatchMusicsWithContext(ctx context.Context, request *DescribeKTVMatchMusicsRequest) (response *DescribeKTVMatchMusicsResponse, err error) {
    if request == nil {
        request = NewDescribeKTVMatchMusicsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKTVMatchMusics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKTVMatchMusicsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKTVMusicsByTagRequest() (request *DescribeKTVMusicsByTagRequest) {
    request = &DescribeKTVMusicsByTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yinsuda", APIVersion, "DescribeKTVMusicsByTag")
    
    
    return
}

func NewDescribeKTVMusicsByTagResponse() (response *DescribeKTVMusicsByTagResponse) {
    response = &DescribeKTVMusicsByTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKTVMusicsByTag
// 通过标签过滤歌曲列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeKTVMusicsByTag(request *DescribeKTVMusicsByTagRequest) (response *DescribeKTVMusicsByTagResponse, err error) {
    return c.DescribeKTVMusicsByTagWithContext(context.Background(), request)
}

// DescribeKTVMusicsByTag
// 通过标签过滤歌曲列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeKTVMusicsByTagWithContext(ctx context.Context, request *DescribeKTVMusicsByTagRequest) (response *DescribeKTVMusicsByTagResponse, err error) {
    if request == nil {
        request = NewDescribeKTVMusicsByTagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKTVMusicsByTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKTVMusicsByTagResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeKTVPlaylistDetailRequest() (request *DescribeKTVPlaylistDetailRequest) {
    request = &DescribeKTVPlaylistDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yinsuda", APIVersion, "DescribeKTVPlaylistDetail")
    
    
    return
}

func NewDescribeKTVPlaylistDetailResponse() (response *DescribeKTVPlaylistDetailResponse) {
    response = &DescribeKTVPlaylistDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKTVPlaylistDetail
// 根据歌单 Id 获取歌单详情。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeKTVPlaylistDetail(request *DescribeKTVPlaylistDetailRequest) (response *DescribeKTVPlaylistDetailResponse, err error) {
    return c.DescribeKTVPlaylistDetailWithContext(context.Background(), request)
}

// DescribeKTVPlaylistDetail
// 根据歌单 Id 获取歌单详情。
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
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
    
    request.Init().WithApiInfo("yinsuda", APIVersion, "DescribeKTVPlaylists")
    
    
    return
}

func NewDescribeKTVPlaylistsResponse() (response *DescribeKTVPlaylistsResponse) {
    response = &DescribeKTVPlaylistsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKTVPlaylists
// 获取歌单列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeKTVPlaylists(request *DescribeKTVPlaylistsRequest) (response *DescribeKTVPlaylistsResponse, err error) {
    return c.DescribeKTVPlaylistsWithContext(context.Background(), request)
}

// DescribeKTVPlaylists
// 获取歌单列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
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
    
    request.Init().WithApiInfo("yinsuda", APIVersion, "DescribeKTVRobots")
    
    
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

func NewDescribeKTVSuggestionsRequest() (request *DescribeKTVSuggestionsRequest) {
    request = &DescribeKTVSuggestionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yinsuda", APIVersion, "DescribeKTVSuggestions")
    
    
    return
}

func NewDescribeKTVSuggestionsResponse() (response *DescribeKTVSuggestionsResponse) {
    response = &DescribeKTVSuggestionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKTVSuggestions
// 根据关键词获取联想词列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeKTVSuggestions(request *DescribeKTVSuggestionsRequest) (response *DescribeKTVSuggestionsResponse, err error) {
    return c.DescribeKTVSuggestionsWithContext(context.Background(), request)
}

// DescribeKTVSuggestions
// 根据关键词获取联想词列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
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

func NewDescribeKTVTagsRequest() (request *DescribeKTVTagsRequest) {
    request = &DescribeKTVTagsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yinsuda", APIVersion, "DescribeKTVTags")
    
    
    return
}

func NewDescribeKTVTagsResponse() (response *DescribeKTVTagsResponse) {
    response = &DescribeKTVTagsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeKTVTags
// 获取标签分组及分组下的标签列表信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeKTVTags(request *DescribeKTVTagsRequest) (response *DescribeKTVTagsResponse, err error) {
    return c.DescribeKTVTagsWithContext(context.Background(), request)
}

// DescribeKTVTags
// 获取标签分组及分组下的标签列表信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeKTVTagsWithContext(ctx context.Context, request *DescribeKTVTagsRequest) (response *DescribeKTVTagsResponse, err error) {
    if request == nil {
        request = NewDescribeKTVTagsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeKTVTags require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeKTVTagsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeLiveVipTradeInfosRequest() (request *DescribeLiveVipTradeInfosRequest) {
    request = &DescribeLiveVipTradeInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yinsuda", APIVersion, "DescribeLiveVipTradeInfos")
    
    
    return
}

func NewDescribeLiveVipTradeInfosResponse() (response *DescribeLiveVipTradeInfosResponse) {
    response = &DescribeLiveVipTradeInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeLiveVipTradeInfos
// 批量获取直播会员充值流水详细信息，包括：流水号，订单状态，下订单时间等
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLiveVipTradeInfos(request *DescribeLiveVipTradeInfosRequest) (response *DescribeLiveVipTradeInfosResponse, err error) {
    return c.DescribeLiveVipTradeInfosWithContext(context.Background(), request)
}

// DescribeLiveVipTradeInfos
// 批量获取直播会员充值流水详细信息，包括：流水号，订单状态，下订单时间等
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DescribeLiveVipTradeInfosWithContext(ctx context.Context, request *DescribeLiveVipTradeInfosRequest) (response *DescribeLiveVipTradeInfosResponse, err error) {
    if request == nil {
        request = NewDescribeLiveVipTradeInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeLiveVipTradeInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeLiveVipTradeInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserInfoRequest() (request *DescribeUserInfoRequest) {
    request = &DescribeUserInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yinsuda", APIVersion, "DescribeUserInfo")
    
    
    return
}

func NewDescribeUserInfoResponse() (response *DescribeUserInfoResponse) {
    response = &DescribeUserInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserInfo
// 获取用户信息，包括是否为直播会员，及直播会员信息等
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeUserInfo(request *DescribeUserInfoRequest) (response *DescribeUserInfoResponse, err error) {
    return c.DescribeUserInfoWithContext(context.Background(), request)
}

// DescribeUserInfo
// 获取用户信息，包括是否为直播会员，及直播会员信息等
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeUserInfoWithContext(ctx context.Context, request *DescribeUserInfoRequest) (response *DescribeUserInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUserInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyKTVRobotRequest() (request *DestroyKTVRobotRequest) {
    request = &DestroyKTVRobotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yinsuda", APIVersion, "DestroyKTVRobot")
    
    
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

func NewRechargeLiveVipRequest() (request *RechargeLiveVipRequest) {
    request = &RechargeLiveVipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yinsuda", APIVersion, "RechargeLiveVip")
    
    
    return
}

func NewRechargeLiveVipResponse() (response *RechargeLiveVipResponse) {
    response = &RechargeLiveVipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RechargeLiveVip
// 充值直播会员，使该用户可以在直播场景使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATETRADESERIALNO = "FailedOperation.DuplicateTradeSerialNo"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RechargeLiveVip(request *RechargeLiveVipRequest) (response *RechargeLiveVipResponse, err error) {
    return c.RechargeLiveVipWithContext(context.Background(), request)
}

// RechargeLiveVip
// 充值直播会员，使该用户可以在直播场景使用
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATETRADESERIALNO = "FailedOperation.DuplicateTradeSerialNo"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RechargeLiveVipWithContext(ctx context.Context, request *RechargeLiveVipRequest) (response *RechargeLiveVipResponse, err error) {
    if request == nil {
        request = NewRechargeLiveVipRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RechargeLiveVip require credential")
    }

    request.SetContext(ctx)
    
    response = NewRechargeLiveVipResponse()
    err = c.Send(request, response)
    return
}

func NewSearchKTVMusicsRequest() (request *SearchKTVMusicsRequest) {
    request = &SearchKTVMusicsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("yinsuda", APIVersion, "SearchKTVMusics")
    
    
    return
}

func NewSearchKTVMusicsResponse() (response *SearchKTVMusicsResponse) {
    response = &SearchKTVMusicsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchKTVMusics
// 根据关键词搜索歌曲，返回相关歌曲列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATETRADESERIALNO = "FailedOperation.DuplicateTradeSerialNo"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) SearchKTVMusics(request *SearchKTVMusicsRequest) (response *SearchKTVMusicsResponse, err error) {
    return c.SearchKTVMusicsWithContext(context.Background(), request)
}

// SearchKTVMusics
// 根据关键词搜索歌曲，返回相关歌曲列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_DUPLICATETRADESERIALNO = "FailedOperation.DuplicateTradeSerialNo"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
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
    
    request.Init().WithApiInfo("yinsuda", APIVersion, "SyncKTVRobotCommand")
    
    
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
