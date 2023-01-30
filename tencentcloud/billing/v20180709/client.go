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

package v20180709

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-07-09"

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


func NewDescribeAccountBalanceRequest() (request *DescribeAccountBalanceRequest) {
    request = &DescribeAccountBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeAccountBalance")
    
    
    return
}

func NewDescribeAccountBalanceResponse() (response *DescribeAccountBalanceResponse) {
    response = &DescribeAccountBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAccountBalance
// 获取云账户余额信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PAYPRICEERROR = "FailedOperation.PayPriceError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccountBalance(request *DescribeAccountBalanceRequest) (response *DescribeAccountBalanceResponse, err error) {
    return c.DescribeAccountBalanceWithContext(context.Background(), request)
}

// DescribeAccountBalance
// 获取云账户余额信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_PAYPRICEERROR = "FailedOperation.PayPriceError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAccountBalanceWithContext(ctx context.Context, request *DescribeAccountBalanceRequest) (response *DescribeAccountBalanceResponse, err error) {
    if request == nil {
        request = NewDescribeAccountBalanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAccountBalance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAccountBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillDetailRequest() (request *DescribeBillDetailRequest) {
    request = &DescribeBillDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillDetail")
    
    
    return
}

func NewDescribeBillDetailResponse() (response *DescribeBillDetailResponse) {
    response = &DescribeBillDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillDetail
// 查询账单明细数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYCOUNTFAILED = "FailedOperation.QueryCountFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBillDetail(request *DescribeBillDetailRequest) (response *DescribeBillDetailResponse, err error) {
    return c.DescribeBillDetailWithContext(context.Background(), request)
}

// DescribeBillDetail
// 查询账单明细数据
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYCOUNTFAILED = "FailedOperation.QueryCountFailed"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeBillDetailWithContext(ctx context.Context, request *DescribeBillDetailRequest) (response *DescribeBillDetailResponse, err error) {
    if request == nil {
        request = NewDescribeBillDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillListRequest() (request *DescribeBillListRequest) {
    request = &DescribeBillListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillList")
    
    
    return
}

func NewDescribeBillListResponse() (response *DescribeBillListResponse) {
    response = &DescribeBillListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillList
// 获取收支明细列表，支持翻页和参数过滤
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBillList(request *DescribeBillListRequest) (response *DescribeBillListResponse, err error) {
    return c.DescribeBillListWithContext(context.Background(), request)
}

// DescribeBillList
// 获取收支明细列表，支持翻页和参数过滤
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeBillListWithContext(ctx context.Context, request *DescribeBillListRequest) (response *DescribeBillListResponse, err error) {
    if request == nil {
        request = NewDescribeBillListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillResourceSummaryRequest() (request *DescribeBillResourceSummaryRequest) {
    request = &DescribeBillResourceSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillResourceSummary")
    
    
    return
}

func NewDescribeBillResourceSummaryResponse() (response *DescribeBillResourceSummaryResponse) {
    response = &DescribeBillResourceSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillResourceSummary
// 查询账单资源汇总数据 
//
// 可能返回的错误码:
//  FAILEDOPERATION_SUMMARYDATANOTREADY = "FailedOperation.SummaryDataNotReady"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeBillResourceSummary(request *DescribeBillResourceSummaryRequest) (response *DescribeBillResourceSummaryResponse, err error) {
    return c.DescribeBillResourceSummaryWithContext(context.Background(), request)
}

// DescribeBillResourceSummary
// 查询账单资源汇总数据 
//
// 可能返回的错误码:
//  FAILEDOPERATION_SUMMARYDATANOTREADY = "FailedOperation.SummaryDataNotReady"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeBillResourceSummaryWithContext(ctx context.Context, request *DescribeBillResourceSummaryRequest) (response *DescribeBillResourceSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeBillResourceSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillResourceSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillResourceSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByPayModeRequest() (request *DescribeBillSummaryByPayModeRequest) {
    request = &DescribeBillSummaryByPayModeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByPayMode")
    
    
    return
}

func NewDescribeBillSummaryByPayModeResponse() (response *DescribeBillSummaryByPayModeResponse) {
    response = &DescribeBillSummaryByPayModeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillSummaryByPayMode
// 获取按付费模式汇总费用分布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByPayMode(request *DescribeBillSummaryByPayModeRequest) (response *DescribeBillSummaryByPayModeResponse, err error) {
    return c.DescribeBillSummaryByPayModeWithContext(context.Background(), request)
}

// DescribeBillSummaryByPayMode
// 获取按付费模式汇总费用分布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByPayModeWithContext(ctx context.Context, request *DescribeBillSummaryByPayModeRequest) (response *DescribeBillSummaryByPayModeResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByPayModeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByPayMode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByPayModeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByProductRequest() (request *DescribeBillSummaryByProductRequest) {
    request = &DescribeBillSummaryByProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByProduct")
    
    
    return
}

func NewDescribeBillSummaryByProductResponse() (response *DescribeBillSummaryByProductResponse) {
    response = &DescribeBillSummaryByProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillSummaryByProduct
// 获取产品汇总费用分布
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByProduct(request *DescribeBillSummaryByProductRequest) (response *DescribeBillSummaryByProductResponse, err error) {
    return c.DescribeBillSummaryByProductWithContext(context.Background(), request)
}

// DescribeBillSummaryByProduct
// 获取产品汇总费用分布
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByProductWithContext(ctx context.Context, request *DescribeBillSummaryByProductRequest) (response *DescribeBillSummaryByProductResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByProjectRequest() (request *DescribeBillSummaryByProjectRequest) {
    request = &DescribeBillSummaryByProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByProject")
    
    
    return
}

func NewDescribeBillSummaryByProjectResponse() (response *DescribeBillSummaryByProjectResponse) {
    response = &DescribeBillSummaryByProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillSummaryByProject
// 获取按项目汇总费用分布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByProject(request *DescribeBillSummaryByProjectRequest) (response *DescribeBillSummaryByProjectResponse, err error) {
    return c.DescribeBillSummaryByProjectWithContext(context.Background(), request)
}

// DescribeBillSummaryByProject
// 获取按项目汇总费用分布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByProjectWithContext(ctx context.Context, request *DescribeBillSummaryByProjectRequest) (response *DescribeBillSummaryByProjectResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByRegionRequest() (request *DescribeBillSummaryByRegionRequest) {
    request = &DescribeBillSummaryByRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByRegion")
    
    
    return
}

func NewDescribeBillSummaryByRegionResponse() (response *DescribeBillSummaryByRegionResponse) {
    response = &DescribeBillSummaryByRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillSummaryByRegion
// 获取按地域汇总费用分布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByRegion(request *DescribeBillSummaryByRegionRequest) (response *DescribeBillSummaryByRegionResponse, err error) {
    return c.DescribeBillSummaryByRegionWithContext(context.Background(), request)
}

// DescribeBillSummaryByRegion
// 获取按地域汇总费用分布
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByRegionWithContext(ctx context.Context, request *DescribeBillSummaryByRegionRequest) (response *DescribeBillSummaryByRegionResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBillSummaryByTagRequest() (request *DescribeBillSummaryByTagRequest) {
    request = &DescribeBillSummaryByTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeBillSummaryByTag")
    
    
    return
}

func NewDescribeBillSummaryByTagResponse() (response *DescribeBillSummaryByTagResponse) {
    response = &DescribeBillSummaryByTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBillSummaryByTag
// 获取按标签汇总费用分布
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByTag(request *DescribeBillSummaryByTagRequest) (response *DescribeBillSummaryByTagResponse, err error) {
    return c.DescribeBillSummaryByTagWithContext(context.Background(), request)
}

// DescribeBillSummaryByTag
// 获取按标签汇总费用分布
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeBillSummaryByTagWithContext(ctx context.Context, request *DescribeBillSummaryByTagRequest) (response *DescribeBillSummaryByTagResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByTagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBillSummaryByTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBillSummaryByTagResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCostDetailRequest() (request *DescribeCostDetailRequest) {
    request = &DescribeCostDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeCostDetail")
    
    
    return
}

func NewDescribeCostDetailResponse() (response *DescribeCostDetailResponse) {
    response = &DescribeCostDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCostDetail
// 查询消耗明细
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeCostDetail(request *DescribeCostDetailRequest) (response *DescribeCostDetailResponse, err error) {
    return c.DescribeCostDetailWithContext(context.Background(), request)
}

// DescribeCostDetail
// 查询消耗明细
//
// 可能返回的错误码:
//  FAILEDOPERATION_TAGKEYNOTEXIST = "FailedOperation.TagKeyNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeCostDetailWithContext(ctx context.Context, request *DescribeCostDetailRequest) (response *DescribeCostDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCostDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCostDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCostDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCostSummaryByProductRequest() (request *DescribeCostSummaryByProductRequest) {
    request = &DescribeCostSummaryByProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeCostSummaryByProduct")
    
    
    return
}

func NewDescribeCostSummaryByProductResponse() (response *DescribeCostSummaryByProductResponse) {
    response = &DescribeCostSummaryByProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCostSummaryByProduct
// 获取按产品汇总消耗详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByProduct(request *DescribeCostSummaryByProductRequest) (response *DescribeCostSummaryByProductResponse, err error) {
    return c.DescribeCostSummaryByProductWithContext(context.Background(), request)
}

// DescribeCostSummaryByProduct
// 获取按产品汇总消耗详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByProductWithContext(ctx context.Context, request *DescribeCostSummaryByProductRequest) (response *DescribeCostSummaryByProductResponse, err error) {
    if request == nil {
        request = NewDescribeCostSummaryByProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCostSummaryByProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCostSummaryByProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCostSummaryByProjectRequest() (request *DescribeCostSummaryByProjectRequest) {
    request = &DescribeCostSummaryByProjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeCostSummaryByProject")
    
    
    return
}

func NewDescribeCostSummaryByProjectResponse() (response *DescribeCostSummaryByProjectResponse) {
    response = &DescribeCostSummaryByProjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCostSummaryByProject
// 获取按项目汇总消耗详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByProject(request *DescribeCostSummaryByProjectRequest) (response *DescribeCostSummaryByProjectResponse, err error) {
    return c.DescribeCostSummaryByProjectWithContext(context.Background(), request)
}

// DescribeCostSummaryByProject
// 获取按项目汇总消耗详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByProjectWithContext(ctx context.Context, request *DescribeCostSummaryByProjectRequest) (response *DescribeCostSummaryByProjectResponse, err error) {
    if request == nil {
        request = NewDescribeCostSummaryByProjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCostSummaryByProject require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCostSummaryByProjectResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCostSummaryByRegionRequest() (request *DescribeCostSummaryByRegionRequest) {
    request = &DescribeCostSummaryByRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeCostSummaryByRegion")
    
    
    return
}

func NewDescribeCostSummaryByRegionResponse() (response *DescribeCostSummaryByRegionResponse) {
    response = &DescribeCostSummaryByRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCostSummaryByRegion
// 获取按地域汇总消耗详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByRegion(request *DescribeCostSummaryByRegionRequest) (response *DescribeCostSummaryByRegionResponse, err error) {
    return c.DescribeCostSummaryByRegionWithContext(context.Background(), request)
}

// DescribeCostSummaryByRegion
// 获取按地域汇总消耗详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByRegionWithContext(ctx context.Context, request *DescribeCostSummaryByRegionRequest) (response *DescribeCostSummaryByRegionResponse, err error) {
    if request == nil {
        request = NewDescribeCostSummaryByRegionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCostSummaryByRegion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCostSummaryByRegionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCostSummaryByResourceRequest() (request *DescribeCostSummaryByResourceRequest) {
    request = &DescribeCostSummaryByResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeCostSummaryByResource")
    
    
    return
}

func NewDescribeCostSummaryByResourceResponse() (response *DescribeCostSummaryByResourceResponse) {
    response = &DescribeCostSummaryByResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCostSummaryByResource
// 获取按资源汇总消耗详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByResource(request *DescribeCostSummaryByResourceRequest) (response *DescribeCostSummaryByResourceResponse, err error) {
    return c.DescribeCostSummaryByResourceWithContext(context.Background(), request)
}

// DescribeCostSummaryByResource
// 获取按资源汇总消耗详情
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeCostSummaryByResourceWithContext(ctx context.Context, request *DescribeCostSummaryByResourceRequest) (response *DescribeCostSummaryByResourceResponse, err error) {
    if request == nil {
        request = NewDescribeCostSummaryByResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCostSummaryByResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCostSummaryByResourceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDealsByCondRequest() (request *DescribeDealsByCondRequest) {
    request = &DescribeDealsByCondRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeDealsByCond")
    
    
    return
}

func NewDescribeDealsByCondResponse() (response *DescribeDealsByCondResponse) {
    response = &DescribeDealsByCondResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDealsByCond
// 查询订单
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDealsByCond(request *DescribeDealsByCondRequest) (response *DescribeDealsByCondResponse, err error) {
    return c.DescribeDealsByCondWithContext(context.Background(), request)
}

// DescribeDealsByCond
// 查询订单
//
// 可能返回的错误码:
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeDealsByCondWithContext(ctx context.Context, request *DescribeDealsByCondRequest) (response *DescribeDealsByCondResponse, err error) {
    if request == nil {
        request = NewDescribeDealsByCondRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDealsByCond require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDealsByCondResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDosageCosDetailByDateRequest() (request *DescribeDosageCosDetailByDateRequest) {
    request = &DescribeDosageCosDetailByDateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeDosageCosDetailByDate")
    
    
    return
}

func NewDescribeDosageCosDetailByDateResponse() (response *DescribeDosageCosDetailByDateResponse) {
    response = &DescribeDosageCosDetailByDateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDosageCosDetailByDate
// 获取COS产品用量明细
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeDosageCosDetailByDate(request *DescribeDosageCosDetailByDateRequest) (response *DescribeDosageCosDetailByDateResponse, err error) {
    return c.DescribeDosageCosDetailByDateWithContext(context.Background(), request)
}

// DescribeDosageCosDetailByDate
// 获取COS产品用量明细
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeDosageCosDetailByDateWithContext(ctx context.Context, request *DescribeDosageCosDetailByDateRequest) (response *DescribeDosageCosDetailByDateResponse, err error) {
    if request == nil {
        request = NewDescribeDosageCosDetailByDateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDosageCosDetailByDate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDosageCosDetailByDateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDosageDetailByDateRequest() (request *DescribeDosageDetailByDateRequest) {
    request = &DescribeDosageDetailByDateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeDosageDetailByDate")
    
    
    return
}

func NewDescribeDosageDetailByDateResponse() (response *DescribeDosageDetailByDateResponse) {
    response = &DescribeDosageDetailByDateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDosageDetailByDate
// 按日期获取产品用量明细
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeDosageDetailByDate(request *DescribeDosageDetailByDateRequest) (response *DescribeDosageDetailByDateResponse, err error) {
    return c.DescribeDosageDetailByDateWithContext(context.Background(), request)
}

// DescribeDosageDetailByDate
// 按日期获取产品用量明细
//
// 可能返回的错误码:
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
func (c *Client) DescribeDosageDetailByDateWithContext(ctx context.Context, request *DescribeDosageDetailByDateRequest) (response *DescribeDosageDetailByDateResponse, err error) {
    if request == nil {
        request = NewDescribeDosageDetailByDateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDosageDetailByDate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDosageDetailByDateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVoucherInfoRequest() (request *DescribeVoucherInfoRequest) {
    request = &DescribeVoucherInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeVoucherInfo")
    
    
    return
}

func NewDescribeVoucherInfoResponse() (response *DescribeVoucherInfoResponse) {
    response = &DescribeVoucherInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVoucherInfo
// 获取代金券相关信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDAPPID = "FailedOperation.InvalidAppId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeVoucherInfo(request *DescribeVoucherInfoRequest) (response *DescribeVoucherInfoResponse, err error) {
    return c.DescribeVoucherInfoWithContext(context.Background(), request)
}

// DescribeVoucherInfo
// 获取代金券相关信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDAPPID = "FailedOperation.InvalidAppId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeVoucherInfoWithContext(ctx context.Context, request *DescribeVoucherInfoRequest) (response *DescribeVoucherInfoResponse, err error) {
    if request == nil {
        request = NewDescribeVoucherInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVoucherInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVoucherInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVoucherUsageDetailsRequest() (request *DescribeVoucherUsageDetailsRequest) {
    request = &DescribeVoucherUsageDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "DescribeVoucherUsageDetails")
    
    
    return
}

func NewDescribeVoucherUsageDetailsResponse() (response *DescribeVoucherUsageDetailsResponse) {
    response = &DescribeVoucherUsageDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVoucherUsageDetails
// 获取代金券使用记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDAPPID = "FailedOperation.InvalidAppId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeVoucherUsageDetails(request *DescribeVoucherUsageDetailsRequest) (response *DescribeVoucherUsageDetailsResponse, err error) {
    return c.DescribeVoucherUsageDetailsWithContext(context.Background(), request)
}

// DescribeVoucherUsageDetails
// 获取代金券使用记录
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDAPPID = "FailedOperation.InvalidAppId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GATEWAYERROR = "InternalError.GatewayError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeVoucherUsageDetailsWithContext(ctx context.Context, request *DescribeVoucherUsageDetailsRequest) (response *DescribeVoucherUsageDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeVoucherUsageDetailsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVoucherUsageDetails require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVoucherUsageDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewPayDealsRequest() (request *PayDealsRequest) {
    request = &PayDealsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("billing", APIVersion, "PayDeals")
    
    
    return
}

func NewPayDealsResponse() (response *PayDealsResponse) {
    response = &PayDealsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PayDeals
// 支付订单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTPAYDEALCANNOTDOWN = "FailedOperation.AgentPayDealCannotDown"
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_INVALIDDEAL = "FailedOperation.InvalidDeal"
//  FAILEDOPERATION_INVALIDVOUCHER = "FailedOperation.InvalidVoucher"
//  FAILEDOPERATION_NEEDPAYTOGETER = "FailedOperation.NeedPayTogeter"
//  FAILEDOPERATION_NEEDPAYTOGETHER = "FailedOperation.NeedPayTogether"
//  FAILEDOPERATION_PAYPRICEERROR = "FailedOperation.PayPriceError"
//  FAILEDOPERATION_PAYSUCCDELIVERFAILED = "FailedOperation.PaySuccDeliverFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CERTIFICATIONNEEDUPGRADE = "UnauthorizedOperation.CertificationNeedUpgrade"
//  UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"
func (c *Client) PayDeals(request *PayDealsRequest) (response *PayDealsResponse, err error) {
    return c.PayDealsWithContext(context.Background(), request)
}

// PayDeals
// 支付订单
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_AGENTPAYDEALCANNOTDOWN = "FailedOperation.AgentPayDealCannotDown"
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_INVALIDDEAL = "FailedOperation.InvalidDeal"
//  FAILEDOPERATION_INVALIDVOUCHER = "FailedOperation.InvalidVoucher"
//  FAILEDOPERATION_NEEDPAYTOGETER = "FailedOperation.NeedPayTogeter"
//  FAILEDOPERATION_NEEDPAYTOGETHER = "FailedOperation.NeedPayTogether"
//  FAILEDOPERATION_PAYPRICEERROR = "FailedOperation.PayPriceError"
//  FAILEDOPERATION_PAYSUCCDELIVERFAILED = "FailedOperation.PaySuccDeliverFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
//  UNAUTHORIZEDOPERATION_CERTIFICATIONNEEDUPGRADE = "UnauthorizedOperation.CertificationNeedUpgrade"
//  UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"
func (c *Client) PayDealsWithContext(ctx context.Context, request *PayDealsRequest) (response *PayDealsResponse, err error) {
    if request == nil {
        request = NewPayDealsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PayDeals require credential")
    }

    request.SetContext(ctx)
    
    response = NewPayDealsResponse()
    err = c.Send(request, response)
    return
}
