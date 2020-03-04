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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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

// 获取云账户余额信息。
func (c *Client) DescribeAccountBalance(request *DescribeAccountBalanceRequest) (response *DescribeAccountBalanceResponse, err error) {
    if request == nil {
        request = NewDescribeAccountBalanceRequest()
    }
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

// 查询账单明细数据
func (c *Client) DescribeBillDetail(request *DescribeBillDetailRequest) (response *DescribeBillDetailResponse, err error) {
    if request == nil {
        request = NewDescribeBillDetailRequest()
    }
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

// 获取收支明细列表，支持翻页和参数过滤
func (c *Client) DescribeBillList(request *DescribeBillListRequest) (response *DescribeBillListResponse, err error) {
    if request == nil {
        request = NewDescribeBillListRequest()
    }
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

// 查询账单资源汇总数据 
func (c *Client) DescribeBillResourceSummary(request *DescribeBillResourceSummaryRequest) (response *DescribeBillResourceSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeBillResourceSummaryRequest()
    }
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

// 获取按付费模式汇总费用分布
func (c *Client) DescribeBillSummaryByPayMode(request *DescribeBillSummaryByPayModeRequest) (response *DescribeBillSummaryByPayModeResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByPayModeRequest()
    }
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

// 获取产品汇总费用分布
func (c *Client) DescribeBillSummaryByProduct(request *DescribeBillSummaryByProductRequest) (response *DescribeBillSummaryByProductResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByProductRequest()
    }
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

// 获取按项目汇总费用分布
func (c *Client) DescribeBillSummaryByProject(request *DescribeBillSummaryByProjectRequest) (response *DescribeBillSummaryByProjectResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByProjectRequest()
    }
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

// 获取按地域汇总费用分布
func (c *Client) DescribeBillSummaryByRegion(request *DescribeBillSummaryByRegionRequest) (response *DescribeBillSummaryByRegionResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByRegionRequest()
    }
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

// 获取按标签汇总费用分布
func (c *Client) DescribeBillSummaryByTag(request *DescribeBillSummaryByTagRequest) (response *DescribeBillSummaryByTagResponse, err error) {
    if request == nil {
        request = NewDescribeBillSummaryByTagRequest()
    }
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

// 查询消耗明细
func (c *Client) DescribeCostDetail(request *DescribeCostDetailRequest) (response *DescribeCostDetailResponse, err error) {
    if request == nil {
        request = NewDescribeCostDetailRequest()
    }
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

// 获取按产品汇总消耗详情
func (c *Client) DescribeCostSummaryByProduct(request *DescribeCostSummaryByProductRequest) (response *DescribeCostSummaryByProductResponse, err error) {
    if request == nil {
        request = NewDescribeCostSummaryByProductRequest()
    }
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

// 获取按项目汇总消耗详情
func (c *Client) DescribeCostSummaryByProject(request *DescribeCostSummaryByProjectRequest) (response *DescribeCostSummaryByProjectResponse, err error) {
    if request == nil {
        request = NewDescribeCostSummaryByProjectRequest()
    }
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

// 获取按地域汇总消耗详情
func (c *Client) DescribeCostSummaryByRegion(request *DescribeCostSummaryByRegionRequest) (response *DescribeCostSummaryByRegionResponse, err error) {
    if request == nil {
        request = NewDescribeCostSummaryByRegionRequest()
    }
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

// 获取按资源汇总消耗详情
func (c *Client) DescribeCostSummaryByResource(request *DescribeCostSummaryByResourceRequest) (response *DescribeCostSummaryByResourceResponse, err error) {
    if request == nil {
        request = NewDescribeCostSummaryByResourceRequest()
    }
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

// 查询订单
func (c *Client) DescribeDealsByCond(request *DescribeDealsByCondRequest) (response *DescribeDealsByCondResponse, err error) {
    if request == nil {
        request = NewDescribeDealsByCondRequest()
    }
    response = NewDescribeDealsByCondResponse()
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

// 按日期获取产品用量明细
func (c *Client) DescribeDosageDetailByDate(request *DescribeDosageDetailByDateRequest) (response *DescribeDosageDetailByDateResponse, err error) {
    if request == nil {
        request = NewDescribeDosageDetailByDateRequest()
    }
    response = NewDescribeDosageDetailByDateResponse()
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

// 支付订单
func (c *Client) PayDeals(request *PayDealsRequest) (response *PayDealsResponse, err error) {
    if request == nil {
        request = NewPayDealsRequest()
    }
    response = NewPayDealsResponse()
    err = c.Send(request, response)
    return
}
