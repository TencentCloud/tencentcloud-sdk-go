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

package v20200513

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-05-13"

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


func NewBatchDescribeOrderCertificateRequest() (request *BatchDescribeOrderCertificateRequest) {
    request = &BatchDescribeOrderCertificateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ape", APIVersion, "BatchDescribeOrderCertificate")
    
    
    return
}

func NewBatchDescribeOrderCertificateResponse() (response *BatchDescribeOrderCertificateResponse) {
    response = &BatchDescribeOrderCertificateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchDescribeOrderCertificate
// 批量获取授权书下载地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND_SENSITIVEIMAGE = "ResourceNotFound.SensitiveImage"
//  RESOURCENOTFOUND_SENSITIVESEARCH = "ResourceNotFound.SensitiveSearch"
func (c *Client) BatchDescribeOrderCertificate(request *BatchDescribeOrderCertificateRequest) (response *BatchDescribeOrderCertificateResponse, err error) {
    if request == nil {
        request = NewBatchDescribeOrderCertificateRequest()
    }
    
    response = NewBatchDescribeOrderCertificateResponse()
    err = c.Send(request, response)
    return
}

func NewBatchDescribeOrderImageRequest() (request *BatchDescribeOrderImageRequest) {
    request = &BatchDescribeOrderImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ape", APIVersion, "BatchDescribeOrderImage")
    
    
    return
}

func NewBatchDescribeOrderImageResponse() (response *BatchDescribeOrderImageResponse) {
    response = &BatchDescribeOrderImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BatchDescribeOrderImage
// 批量获取图片下载地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_ORDEREXPIREDERROR = "LimitExceeded.OrderExpiredError"
//  RESOURCENOTFOUND_SENSITIVEIMAGE = "ResourceNotFound.SensitiveImage"
//  RESOURCENOTFOUND_SENSITIVESEARCH = "ResourceNotFound.SensitiveSearch"
func (c *Client) BatchDescribeOrderImage(request *BatchDescribeOrderImageRequest) (response *BatchDescribeOrderImageResponse, err error) {
    if request == nil {
        request = NewBatchDescribeOrderImageRequest()
    }
    
    response = NewBatchDescribeOrderImageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrderAndDownloadsRequest() (request *CreateOrderAndDownloadsRequest) {
    request = &CreateOrderAndDownloadsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ape", APIVersion, "CreateOrderAndDownloads")
    
    
    return
}

func NewCreateOrderAndDownloadsResponse() (response *CreateOrderAndDownloadsResponse) {
    response = &CreateOrderAndDownloadsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOrderAndDownloads
// 核销图片，获取原图URL地址
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) CreateOrderAndDownloads(request *CreateOrderAndDownloadsRequest) (response *CreateOrderAndDownloadsResponse, err error) {
    if request == nil {
        request = NewCreateOrderAndDownloadsRequest()
    }
    
    response = NewCreateOrderAndDownloadsResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOrderAndPayRequest() (request *CreateOrderAndPayRequest) {
    request = &CreateOrderAndPayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ape", APIVersion, "CreateOrderAndPay")
    
    
    return
}

func NewCreateOrderAndPayResponse() (response *CreateOrderAndPayResponse) {
    response = &CreateOrderAndPayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateOrderAndPay
// 购买一张图片并且支付
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_ORDERLIMITERROR = "LimitExceeded.OrderLimitError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateOrderAndPay(request *CreateOrderAndPayRequest) (response *CreateOrderAndPayResponse, err error) {
    if request == nil {
        request = NewCreateOrderAndPayRequest()
    }
    
    response = NewCreateOrderAndPayResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuthUsersRequest() (request *DescribeAuthUsersRequest) {
    request = &DescribeAuthUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ape", APIVersion, "DescribeAuthUsers")
    
    
    return
}

func NewDescribeAuthUsersResponse() (response *DescribeAuthUsersResponse) {
    response = &DescribeAuthUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAuthUsers
// 分页查询授权人列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAuthUsers(request *DescribeAuthUsersRequest) (response *DescribeAuthUsersResponse, err error) {
    if request == nil {
        request = NewDescribeAuthUsersRequest()
    }
    
    response = NewDescribeAuthUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDownloadInfosRequest() (request *DescribeDownloadInfosRequest) {
    request = &DescribeDownloadInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ape", APIVersion, "DescribeDownloadInfos")
    
    
    return
}

func NewDescribeDownloadInfosResponse() (response *DescribeDownloadInfosResponse) {
    response = &DescribeDownloadInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDownloadInfos
// 获取用户图片下载记录
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDownloadInfos(request *DescribeDownloadInfosRequest) (response *DescribeDownloadInfosResponse, err error) {
    if request == nil {
        request = NewDescribeDownloadInfosRequest()
    }
    
    response = NewDescribeDownloadInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageRequest() (request *DescribeImageRequest) {
    request = &DescribeImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ape", APIVersion, "DescribeImage")
    
    
    return
}

func NewDescribeImageResponse() (response *DescribeImageResponse) {
    response = &DescribeImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImage
// 根据ID查询一张图片的详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_SENSITIVEIMAGE = "ResourceNotFound.SensitiveImage"
//  RESOURCENOTFOUND_SENSITIVESEARCH = "ResourceNotFound.SensitiveSearch"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeImage(request *DescribeImageRequest) (response *DescribeImageResponse, err error) {
    if request == nil {
        request = NewDescribeImageRequest()
    }
    
    response = NewDescribeImageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
    request = &DescribeImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ape", APIVersion, "DescribeImages")
    
    
    return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
    response = &DescribeImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImages
// 根据关键字搜索图片列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_SENSITIVESEARCH = "ResourceNotFound.SensitiveSearch"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    if request == nil {
        request = NewDescribeImagesRequest()
    }
    
    response = NewDescribeImagesResponse()
    err = c.Send(request, response)
    return
}
