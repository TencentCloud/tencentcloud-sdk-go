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

package v20180129

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-01-29"

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


func NewDescribeBrandCommentCountRequest() (request *DescribeBrandCommentCountRequest) {
    request = &DescribeBrandCommentCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbm", APIVersion, "DescribeBrandCommentCount")
    
    
    return
}

func NewDescribeBrandCommentCountResponse() (response *DescribeBrandCommentCountResponse) {
    response = &DescribeBrandCommentCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBrandCommentCount
// 通过分析用户在评价品牌时用词的正负面情绪评分，返回品牌好评与差评评价条数，按天输出结果。
//
// 可能返回的错误码:
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBrandCommentCount(request *DescribeBrandCommentCountRequest) (response *DescribeBrandCommentCountResponse, err error) {
    return c.DescribeBrandCommentCountWithContext(context.Background(), request)
}

// DescribeBrandCommentCount
// 通过分析用户在评价品牌时用词的正负面情绪评分，返回品牌好评与差评评价条数，按天输出结果。
//
// 可能返回的错误码:
//  INTERNALERROR_INNERSERVERFAILED = "InternalError.InnerServerFailed"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBrandCommentCountWithContext(ctx context.Context, request *DescribeBrandCommentCountRequest) (response *DescribeBrandCommentCountResponse, err error) {
    if request == nil {
        request = NewDescribeBrandCommentCountRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBrandCommentCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBrandCommentCountResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBrandExposureRequest() (request *DescribeBrandExposureRequest) {
    request = &DescribeBrandExposureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbm", APIVersion, "DescribeBrandExposure")
    
    
    return
}

func NewDescribeBrandExposureResponse() (response *DescribeBrandExposureResponse) {
    response = &DescribeBrandExposureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBrandExposure
// 监测品牌关键词命中文章标题或全文的文章篇数，按天输出数据。
//
// 可能返回的错误码:
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBrandExposure(request *DescribeBrandExposureRequest) (response *DescribeBrandExposureResponse, err error) {
    return c.DescribeBrandExposureWithContext(context.Background(), request)
}

// DescribeBrandExposure
// 监测品牌关键词命中文章标题或全文的文章篇数，按天输出数据。
//
// 可能返回的错误码:
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBrandExposureWithContext(ctx context.Context, request *DescribeBrandExposureRequest) (response *DescribeBrandExposureResponse, err error) {
    if request == nil {
        request = NewDescribeBrandExposureRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBrandExposure require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBrandExposureResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBrandMediaReportRequest() (request *DescribeBrandMediaReportRequest) {
    request = &DescribeBrandMediaReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbm", APIVersion, "DescribeBrandMediaReport")
    
    
    return
}

func NewDescribeBrandMediaReportResponse() (response *DescribeBrandMediaReportResponse) {
    response = &DescribeBrandMediaReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBrandMediaReport
// 监测品牌关键词出现在媒体网站（新闻媒体、网络门户、政府网站、微信公众号、天天快报等）发布资讯标题和正文中的报道数。按天输出结果。
//
// 可能返回的错误码:
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBrandMediaReport(request *DescribeBrandMediaReportRequest) (response *DescribeBrandMediaReportResponse, err error) {
    return c.DescribeBrandMediaReportWithContext(context.Background(), request)
}

// DescribeBrandMediaReport
// 监测品牌关键词出现在媒体网站（新闻媒体、网络门户、政府网站、微信公众号、天天快报等）发布资讯标题和正文中的报道数。按天输出结果。
//
// 可能返回的错误码:
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBrandMediaReportWithContext(ctx context.Context, request *DescribeBrandMediaReportRequest) (response *DescribeBrandMediaReportResponse, err error) {
    if request == nil {
        request = NewDescribeBrandMediaReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBrandMediaReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBrandMediaReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBrandNegCommentsRequest() (request *DescribeBrandNegCommentsRequest) {
    request = &DescribeBrandNegCommentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbm", APIVersion, "DescribeBrandNegComments")
    
    
    return
}

func NewDescribeBrandNegCommentsResponse() (response *DescribeBrandNegCommentsResponse) {
    response = &DescribeBrandNegCommentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBrandNegComments
// 通过分析用户在评价品牌时用词的正负面情绪评分，返回品牌热门差评观点列表。
//
// 可能返回的错误码:
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBrandNegComments(request *DescribeBrandNegCommentsRequest) (response *DescribeBrandNegCommentsResponse, err error) {
    return c.DescribeBrandNegCommentsWithContext(context.Background(), request)
}

// DescribeBrandNegComments
// 通过分析用户在评价品牌时用词的正负面情绪评分，返回品牌热门差评观点列表。
//
// 可能返回的错误码:
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBrandNegCommentsWithContext(ctx context.Context, request *DescribeBrandNegCommentsRequest) (response *DescribeBrandNegCommentsResponse, err error) {
    if request == nil {
        request = NewDescribeBrandNegCommentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBrandNegComments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBrandNegCommentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBrandPosCommentsRequest() (request *DescribeBrandPosCommentsRequest) {
    request = &DescribeBrandPosCommentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbm", APIVersion, "DescribeBrandPosComments")
    
    
    return
}

func NewDescribeBrandPosCommentsResponse() (response *DescribeBrandPosCommentsResponse) {
    response = &DescribeBrandPosCommentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBrandPosComments
// 通过分析用户在评价品牌时用词的正负面情绪评分，返回品牌热门好评观点列表。
//
// 可能返回的错误码:
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBrandPosComments(request *DescribeBrandPosCommentsRequest) (response *DescribeBrandPosCommentsResponse, err error) {
    return c.DescribeBrandPosCommentsWithContext(context.Background(), request)
}

// DescribeBrandPosComments
// 通过分析用户在评价品牌时用词的正负面情绪评分，返回品牌热门好评观点列表。
//
// 可能返回的错误码:
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBrandPosCommentsWithContext(ctx context.Context, request *DescribeBrandPosCommentsRequest) (response *DescribeBrandPosCommentsResponse, err error) {
    if request == nil {
        request = NewDescribeBrandPosCommentsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBrandPosComments require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBrandPosCommentsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBrandSocialOpinionRequest() (request *DescribeBrandSocialOpinionRequest) {
    request = &DescribeBrandSocialOpinionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbm", APIVersion, "DescribeBrandSocialOpinion")
    
    
    return
}

func NewDescribeBrandSocialOpinionResponse() (response *DescribeBrandSocialOpinionResponse) {
    response = &DescribeBrandSocialOpinionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBrandSocialOpinion
// 检测品牌关键词出现在微博、QQ兴趣部落、论坛、博客等个人公开贡献资讯中的内容，每天聚合近30天热度最高的观点列表。
//
// 可能返回的错误码:
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBrandSocialOpinion(request *DescribeBrandSocialOpinionRequest) (response *DescribeBrandSocialOpinionResponse, err error) {
    return c.DescribeBrandSocialOpinionWithContext(context.Background(), request)
}

// DescribeBrandSocialOpinion
// 检测品牌关键词出现在微博、QQ兴趣部落、论坛、博客等个人公开贡献资讯中的内容，每天聚合近30天热度最高的观点列表。
//
// 可能返回的错误码:
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBrandSocialOpinionWithContext(ctx context.Context, request *DescribeBrandSocialOpinionRequest) (response *DescribeBrandSocialOpinionResponse, err error) {
    if request == nil {
        request = NewDescribeBrandSocialOpinionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBrandSocialOpinion require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBrandSocialOpinionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBrandSocialReportRequest() (request *DescribeBrandSocialReportRequest) {
    request = &DescribeBrandSocialReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbm", APIVersion, "DescribeBrandSocialReport")
    
    
    return
}

func NewDescribeBrandSocialReportResponse() (response *DescribeBrandSocialReportResponse) {
    response = &DescribeBrandSocialReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBrandSocialReport
// 监测品牌关键词出现在微博、QQ兴趣部落、论坛、博客等个人公开贡献资讯中的条数。按天输出数据结果。
//
// 可能返回的错误码:
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBrandSocialReport(request *DescribeBrandSocialReportRequest) (response *DescribeBrandSocialReportResponse, err error) {
    return c.DescribeBrandSocialReportWithContext(context.Background(), request)
}

// DescribeBrandSocialReport
// 监测品牌关键词出现在微博、QQ兴趣部落、论坛、博客等个人公开贡献资讯中的条数。按天输出数据结果。
//
// 可能返回的错误码:
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeBrandSocialReportWithContext(ctx context.Context, request *DescribeBrandSocialReportRequest) (response *DescribeBrandSocialReportResponse, err error) {
    if request == nil {
        request = NewDescribeBrandSocialReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBrandSocialReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBrandSocialReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIndustryNewsRequest() (request *DescribeIndustryNewsRequest) {
    request = &DescribeIndustryNewsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbm", APIVersion, "DescribeIndustryNews")
    
    
    return
}

func NewDescribeIndustryNewsResponse() (response *DescribeIndustryNewsResponse) {
    response = &DescribeIndustryNewsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIndustryNews
// 根据客户定制的行业关键词，监测关键词出现在媒体网站（新闻媒体、网络门户、政府网站、微信公众号、天天快报等）发布资讯标题和正文中的报道数，以及文章列表、来源渠道、作者、发布时间等。
//
// 可能返回的错误码:
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeIndustryNews(request *DescribeIndustryNewsRequest) (response *DescribeIndustryNewsResponse, err error) {
    return c.DescribeIndustryNewsWithContext(context.Background(), request)
}

// DescribeIndustryNews
// 根据客户定制的行业关键词，监测关键词出现在媒体网站（新闻媒体、网络门户、政府网站、微信公众号、天天快报等）发布资讯标题和正文中的报道数，以及文章列表、来源渠道、作者、发布时间等。
//
// 可能返回的错误码:
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeIndustryNewsWithContext(ctx context.Context, request *DescribeIndustryNewsRequest) (response *DescribeIndustryNewsResponse, err error) {
    if request == nil {
        request = NewDescribeIndustryNewsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIndustryNews require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIndustryNewsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserPortraitRequest() (request *DescribeUserPortraitRequest) {
    request = &DescribeUserPortraitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tbm", APIVersion, "DescribeUserPortrait")
    
    
    return
}

func NewDescribeUserPortraitResponse() (response *DescribeUserPortraitResponse) {
    response = &DescribeUserPortraitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserPortrait
// 通过分析洞察参与过品牌媒体互动的用户，比如公开发表品牌的新闻评论、在公开社交渠道发表过对品牌的评价观点等用户，返回用户的画像属性分布，例如性别、年龄、地域、喜爱的明星、喜爱的影视。
//
// 可能返回的错误码:
//  INTERNALERROR_DATAINPROCESSING = "InternalError.DataInProcessing"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeUserPortrait(request *DescribeUserPortraitRequest) (response *DescribeUserPortraitResponse, err error) {
    return c.DescribeUserPortraitWithContext(context.Background(), request)
}

// DescribeUserPortrait
// 通过分析洞察参与过品牌媒体互动的用户，比如公开发表品牌的新闻评论、在公开社交渠道发表过对品牌的评价观点等用户，返回用户的画像属性分布，例如性别、年龄、地域、喜爱的明星、喜爱的影视。
//
// 可能返回的错误码:
//  INTERNALERROR_DATAINPROCESSING = "InternalError.DataInProcessing"
//  INTERNALERROR_METADATAOPFAILED = "InternalError.MetaDataOpFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeUserPortraitWithContext(ctx context.Context, request *DescribeUserPortraitRequest) (response *DescribeUserPortraitResponse, err error) {
    if request == nil {
        request = NewDescribeUserPortraitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserPortrait require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserPortraitResponse()
    err = c.Send(request, response)
    return
}
