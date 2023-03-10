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

package v20220805

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-08-05"

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


func NewDescribeGoodsRecommendRequest() (request *DescribeGoodsRecommendRequest) {
    request = &DescribeGoodsRecommendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("irp", APIVersion, "DescribeGoodsRecommend")
    
    
    return
}

func NewDescribeGoodsRecommendResponse() (response *DescribeGoodsRecommendResponse) {
    response = &DescribeGoodsRecommendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGoodsRecommend
// 获取电商类推荐结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeGoodsRecommend(request *DescribeGoodsRecommendRequest) (response *DescribeGoodsRecommendResponse, err error) {
    return c.DescribeGoodsRecommendWithContext(context.Background(), request)
}

// DescribeGoodsRecommend
// 获取电商类推荐结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeGoodsRecommendWithContext(ctx context.Context, request *DescribeGoodsRecommendRequest) (response *DescribeGoodsRecommendResponse, err error) {
    if request == nil {
        request = NewDescribeGoodsRecommendRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGoodsRecommend require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeGoodsRecommendResponse()
    err = c.Send(request, response)
    return
}

func NewFeedRecommendRequest() (request *FeedRecommendRequest) {
    request = &FeedRecommendRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("irp", APIVersion, "FeedRecommend")
    
    
    return
}

func NewFeedRecommendResponse() (response *FeedRecommendResponse) {
    response = &FeedRecommendResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FeedRecommend
// 获取信息流推荐结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) FeedRecommend(request *FeedRecommendRequest) (response *FeedRecommendResponse, err error) {
    return c.FeedRecommendWithContext(context.Background(), request)
}

// FeedRecommend
// 获取信息流推荐结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) FeedRecommendWithContext(ctx context.Context, request *FeedRecommendRequest) (response *FeedRecommendResponse, err error) {
    if request == nil {
        request = NewFeedRecommendRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FeedRecommend require credential")
    }

    request.SetContext(ctx)
    
    response = NewFeedRecommendResponse()
    err = c.Send(request, response)
    return
}

func NewReportFeedBehaviorRequest() (request *ReportFeedBehaviorRequest) {
    request = &ReportFeedBehaviorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("irp", APIVersion, "ReportFeedBehavior")
    
    
    return
}

func NewReportFeedBehaviorResponse() (response *ReportFeedBehaviorResponse) {
    response = &ReportFeedBehaviorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReportFeedBehavior
// 上报信息流场景内的行为数据，随着数据的积累，模型的效果会逐渐稳定。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReportFeedBehavior(request *ReportFeedBehaviorRequest) (response *ReportFeedBehaviorResponse, err error) {
    return c.ReportFeedBehaviorWithContext(context.Background(), request)
}

// ReportFeedBehavior
// 上报信息流场景内的行为数据，随着数据的积累，模型的效果会逐渐稳定。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReportFeedBehaviorWithContext(ctx context.Context, request *ReportFeedBehaviorRequest) (response *ReportFeedBehaviorResponse, err error) {
    if request == nil {
        request = NewReportFeedBehaviorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReportFeedBehavior require credential")
    }

    request.SetContext(ctx)
    
    response = NewReportFeedBehaviorResponse()
    err = c.Send(request, response)
    return
}

func NewReportFeedItemRequest() (request *ReportFeedItemRequest) {
    request = &ReportFeedItemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("irp", APIVersion, "ReportFeedItem")
    
    
    return
}

func NewReportFeedItemResponse() (response *ReportFeedItemResponse) {
    response = &ReportFeedItemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReportFeedItem
// 上报被用于推荐的信息流内容信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReportFeedItem(request *ReportFeedItemRequest) (response *ReportFeedItemResponse, err error) {
    return c.ReportFeedItemWithContext(context.Background(), request)
}

// ReportFeedItem
// 上报被用于推荐的信息流内容信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReportFeedItemWithContext(ctx context.Context, request *ReportFeedItemRequest) (response *ReportFeedItemResponse, err error) {
    if request == nil {
        request = NewReportFeedItemRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReportFeedItem require credential")
    }

    request.SetContext(ctx)
    
    response = NewReportFeedItemResponse()
    err = c.Send(request, response)
    return
}

func NewReportFeedUserRequest() (request *ReportFeedUserRequest) {
    request = &ReportFeedUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("irp", APIVersion, "ReportFeedUser")
    
    
    return
}

func NewReportFeedUserResponse() (response *ReportFeedUserResponse) {
    response = &ReportFeedUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReportFeedUser
// 上报信息流用户信息，请务必确认用户的唯一性，并在请求推荐结果时指定用户的唯一标识信息（UserId），否则将无法进行千人千面的推荐
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReportFeedUser(request *ReportFeedUserRequest) (response *ReportFeedUserResponse, err error) {
    return c.ReportFeedUserWithContext(context.Background(), request)
}

// ReportFeedUser
// 上报信息流用户信息，请务必确认用户的唯一性，并在请求推荐结果时指定用户的唯一标识信息（UserId），否则将无法进行千人千面的推荐
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReportFeedUserWithContext(ctx context.Context, request *ReportFeedUserRequest) (response *ReportFeedUserResponse, err error) {
    if request == nil {
        request = NewReportFeedUserRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReportFeedUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewReportFeedUserResponse()
    err = c.Send(request, response)
    return
}

func NewReportGoodsBehaviorRequest() (request *ReportGoodsBehaviorRequest) {
    request = &ReportGoodsBehaviorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("irp", APIVersion, "ReportGoodsBehavior")
    
    
    return
}

func NewReportGoodsBehaviorResponse() (response *ReportGoodsBehaviorResponse) {
    response = &ReportGoodsBehaviorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReportGoodsBehavior
// 上报电商类行为数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReportGoodsBehavior(request *ReportGoodsBehaviorRequest) (response *ReportGoodsBehaviorResponse, err error) {
    return c.ReportGoodsBehaviorWithContext(context.Background(), request)
}

// ReportGoodsBehavior
// 上报电商类行为数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReportGoodsBehaviorWithContext(ctx context.Context, request *ReportGoodsBehaviorRequest) (response *ReportGoodsBehaviorResponse, err error) {
    if request == nil {
        request = NewReportGoodsBehaviorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReportGoodsBehavior require credential")
    }

    request.SetContext(ctx)
    
    response = NewReportGoodsBehaviorResponse()
    err = c.Send(request, response)
    return
}

func NewReportGoodsInfoRequest() (request *ReportGoodsInfoRequest) {
    request = &ReportGoodsInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("irp", APIVersion, "ReportGoodsInfo")
    
    
    return
}

func NewReportGoodsInfoResponse() (response *ReportGoodsInfoResponse) {
    response = &ReportGoodsInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReportGoodsInfo
// 上报电商类商品信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReportGoodsInfo(request *ReportGoodsInfoRequest) (response *ReportGoodsInfoResponse, err error) {
    return c.ReportGoodsInfoWithContext(context.Background(), request)
}

// ReportGoodsInfo
// 上报电商类商品信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) ReportGoodsInfoWithContext(ctx context.Context, request *ReportGoodsInfoRequest) (response *ReportGoodsInfoResponse, err error) {
    if request == nil {
        request = NewReportGoodsInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReportGoodsInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewReportGoodsInfoResponse()
    err = c.Send(request, response)
    return
}
