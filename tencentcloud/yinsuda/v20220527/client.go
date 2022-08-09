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
func (c *Client) BatchDescribeKTVMusicDetails(request *BatchDescribeKTVMusicDetailsRequest) (response *BatchDescribeKTVMusicDetailsResponse, err error) {
    return c.BatchDescribeKTVMusicDetailsWithContext(context.Background(), request)
}

// BatchDescribeKTVMusicDetails
// 批量获取歌曲详细信息，包括：歌词下载链接、播放凭证、音高数据下载链接信息等。
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
//  INTERNALERROR = "InternalError"
func (c *Client) DescribeKTVSuggestions(request *DescribeKTVSuggestionsRequest) (response *DescribeKTVSuggestionsResponse, err error) {
    return c.DescribeKTVSuggestionsWithContext(context.Background(), request)
}

// DescribeKTVSuggestions
// 根据关键词获取联想词列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
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
//  INTERNALERROR = "InternalError"
func (c *Client) SearchKTVMusics(request *SearchKTVMusicsRequest) (response *SearchKTVMusicsResponse, err error) {
    return c.SearchKTVMusicsWithContext(context.Background(), request)
}

// SearchKTVMusics
// 根据关键词搜索歌曲，返回相关歌曲列表。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
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
