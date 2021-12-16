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
    if request == nil {
        request = NewDescribeAuthInfoRequest()
    }
    
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
    if request == nil {
        request = NewDescribeCloudMusicRequest()
    }
    
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
    if request == nil {
        request = NewDescribeCloudMusicPurchasedRequest()
    }
    
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
    if request == nil {
        request = NewDescribeItemByIdRequest()
    }
    
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
    if request == nil {
        request = NewDescribeItemsRequest()
    }
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKTVMusicDetail(request *DescribeKTVMusicDetailRequest) (response *DescribeKTVMusicDetailResponse, err error) {
    if request == nil {
        request = NewDescribeKTVMusicDetailRequest()
    }
    
    response = NewDescribeKTVMusicDetailResponse()
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
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKTVPlaylistDetail(request *DescribeKTVPlaylistDetailRequest) (response *DescribeKTVPlaylistDetailResponse, err error) {
    if request == nil {
        request = NewDescribeKTVPlaylistDetailRequest()
    }
    
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
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeKTVPlaylists(request *DescribeKTVPlaylistsRequest) (response *DescribeKTVPlaylistsResponse, err error) {
    if request == nil {
        request = NewDescribeKTVPlaylistsRequest()
    }
    
    response = NewDescribeKTVPlaylistsResponse()
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
    if request == nil {
        request = NewDescribeLyricRequest()
    }
    
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
    if request == nil {
        request = NewDescribeMusicRequest()
    }
    
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
    if request == nil {
        request = NewDescribeMusicSaleStatusRequest()
    }
    
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
    if request == nil {
        request = NewDescribePackageItemsRequest()
    }
    
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
    if request == nil {
        request = NewDescribePackagesRequest()
    }
    
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
    if request == nil {
        request = NewDescribePkgOfflineMusicRequest()
    }
    
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
    if request == nil {
        request = NewDescribeStationsRequest()
    }
    
    response = NewDescribeStationsResponse()
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
    if request == nil {
        request = NewModifyMusicOnShelvesRequest()
    }
    
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
    if request == nil {
        request = NewPutMusicOnTheShelvesRequest()
    }
    
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
    if request == nil {
        request = NewReportDataRequest()
    }
    
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
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SearchKTVMusics(request *SearchKTVMusicsRequest) (response *SearchKTVMusicsResponse, err error) {
    if request == nil {
        request = NewSearchKTVMusicsRequest()
    }
    
    response = NewSearchKTVMusicsResponse()
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
    if request == nil {
        request = NewTakeMusicOffShelvesRequest()
    }
    
    response = NewTakeMusicOffShelvesResponse()
    err = c.Send(request, response)
    return
}
