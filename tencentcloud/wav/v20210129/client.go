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

package v20210129

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-01-29"

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


func NewCreateChannelCodeRequest() (request *CreateChannelCodeRequest) {
    request = &CreateChannelCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "CreateChannelCode")
    
    
    return
}

func NewCreateChannelCodeResponse() (response *CreateChannelCodeResponse) {
    response = &CreateChannelCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateChannelCode
// 新增渠道活码接口
func (c *Client) CreateChannelCode(request *CreateChannelCodeRequest) (response *CreateChannelCodeResponse, err error) {
    return c.CreateChannelCodeWithContext(context.Background(), request)
}

// CreateChannelCode
// 新增渠道活码接口
func (c *Client) CreateChannelCodeWithContext(ctx context.Context, request *CreateChannelCodeRequest) (response *CreateChannelCodeResponse, err error) {
    if request == nil {
        request = NewCreateChannelCodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateChannelCode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateChannelCodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCorpTagRequest() (request *CreateCorpTagRequest) {
    request = &CreateCorpTagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "CreateCorpTag")
    
    
    return
}

func NewCreateCorpTagResponse() (response *CreateCorpTagResponse) {
    response = &CreateCorpTagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCorpTag
// 该接口用户设置标签库, 每个企业最多可配置3000个企业标签。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  AUTHFAILURE_MISSINGACCESSTOKEN = "AuthFailure.MissingAccessToken"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateCorpTag(request *CreateCorpTagRequest) (response *CreateCorpTagResponse, err error) {
    return c.CreateCorpTagWithContext(context.Background(), request)
}

// CreateCorpTag
// 该接口用户设置标签库, 每个企业最多可配置3000个企业标签。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  AUTHFAILURE_MISSINGACCESSTOKEN = "AuthFailure.MissingAccessToken"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateCorpTagWithContext(ctx context.Context, request *CreateCorpTagRequest) (response *CreateCorpTagResponse, err error) {
    if request == nil {
        request = NewCreateCorpTagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCorpTag require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCorpTagResponse()
    err = c.Send(request, response)
    return
}

func NewCreateLeadRequest() (request *CreateLeadRequest) {
    request = &CreateLeadRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "CreateLead")
    
    
    return
}

func NewCreateLeadResponse() (response *CreateLeadResponse) {
    response = &CreateLeadResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateLead
// 线索回收接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_OPENPLATFORMERROR = "FailedOperation.OpenPlatformError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateLead(request *CreateLeadRequest) (response *CreateLeadResponse, err error) {
    return c.CreateLeadWithContext(context.Background(), request)
}

// CreateLead
// 线索回收接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_OPENPLATFORMERROR = "FailedOperation.OpenPlatformError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreateLeadWithContext(ctx context.Context, request *CreateLeadRequest) (response *CreateLeadResponse, err error) {
    if request == nil {
        request = NewCreateLeadRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateLead require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateLeadResponse()
    err = c.Send(request, response)
    return
}

func NewQueryActivityJoinListRequest() (request *QueryActivityJoinListRequest) {
    request = &QueryActivityJoinListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryActivityJoinList")
    
    
    return
}

func NewQueryActivityJoinListResponse() (response *QueryActivityJoinListResponse) {
    response = &QueryActivityJoinListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryActivityJoinList
// 根据游标拉取活动参与列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  FAILEDOPERATION_OPENPLATFORMERROR = "FailedOperation.OpenPlatformError"
//  INTERNALERROR = "InternalError"
func (c *Client) QueryActivityJoinList(request *QueryActivityJoinListRequest) (response *QueryActivityJoinListResponse, err error) {
    return c.QueryActivityJoinListWithContext(context.Background(), request)
}

// QueryActivityJoinList
// 根据游标拉取活动参与列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  FAILEDOPERATION_OPENPLATFORMERROR = "FailedOperation.OpenPlatformError"
//  INTERNALERROR = "InternalError"
func (c *Client) QueryActivityJoinListWithContext(ctx context.Context, request *QueryActivityJoinListRequest) (response *QueryActivityJoinListResponse, err error) {
    if request == nil {
        request = NewQueryActivityJoinListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryActivityJoinList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryActivityJoinListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryActivityListRequest() (request *QueryActivityListRequest) {
    request = &QueryActivityListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryActivityList")
    
    
    return
}

func NewQueryActivityListResponse() (response *QueryActivityListResponse) {
    response = &QueryActivityListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryActivityList
// 根据游标拉取活动列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  FAILEDOPERATION_OPENPLATFORMERROR = "FailedOperation.OpenPlatformError"
//  INTERNALERROR = "InternalError"
func (c *Client) QueryActivityList(request *QueryActivityListRequest) (response *QueryActivityListResponse, err error) {
    return c.QueryActivityListWithContext(context.Background(), request)
}

// QueryActivityList
// 根据游标拉取活动列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  FAILEDOPERATION_OPENPLATFORMERROR = "FailedOperation.OpenPlatformError"
//  INTERNALERROR = "InternalError"
func (c *Client) QueryActivityListWithContext(ctx context.Context, request *QueryActivityListRequest) (response *QueryActivityListResponse, err error) {
    if request == nil {
        request = NewQueryActivityListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryActivityList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryActivityListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryActivityLiveCodeListRequest() (request *QueryActivityLiveCodeListRequest) {
    request = &QueryActivityLiveCodeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryActivityLiveCodeList")
    
    
    return
}

func NewQueryActivityLiveCodeListResponse() (response *QueryActivityLiveCodeListResponse) {
    response = &QueryActivityLiveCodeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryActivityLiveCodeList
// 根据游标拉取活动活码列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
func (c *Client) QueryActivityLiveCodeList(request *QueryActivityLiveCodeListRequest) (response *QueryActivityLiveCodeListResponse, err error) {
    return c.QueryActivityLiveCodeListWithContext(context.Background(), request)
}

// QueryActivityLiveCodeList
// 根据游标拉取活动活码列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
func (c *Client) QueryActivityLiveCodeListWithContext(ctx context.Context, request *QueryActivityLiveCodeListRequest) (response *QueryActivityLiveCodeListResponse, err error) {
    if request == nil {
        request = NewQueryActivityLiveCodeListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryActivityLiveCodeList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryActivityLiveCodeListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryChannelCodeListRequest() (request *QueryChannelCodeListRequest) {
    request = &QueryChannelCodeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryChannelCodeList")
    
    
    return
}

func NewQueryChannelCodeListResponse() (response *QueryChannelCodeListResponse) {
    response = &QueryChannelCodeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryChannelCodeList
// 根据游标拉取渠道活码列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
func (c *Client) QueryChannelCodeList(request *QueryChannelCodeListRequest) (response *QueryChannelCodeListResponse, err error) {
    return c.QueryChannelCodeListWithContext(context.Background(), request)
}

// QueryChannelCodeList
// 根据游标拉取渠道活码列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
func (c *Client) QueryChannelCodeListWithContext(ctx context.Context, request *QueryChannelCodeListRequest) (response *QueryChannelCodeListResponse, err error) {
    if request == nil {
        request = NewQueryChannelCodeListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryChannelCodeList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryChannelCodeListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryChatArchivingListRequest() (request *QueryChatArchivingListRequest) {
    request = &QueryChatArchivingListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryChatArchivingList")
    
    
    return
}

func NewQueryChatArchivingListResponse() (response *QueryChatArchivingListResponse) {
    response = &QueryChatArchivingListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryChatArchivingList
// 根据游标拉取会话存档列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
func (c *Client) QueryChatArchivingList(request *QueryChatArchivingListRequest) (response *QueryChatArchivingListResponse, err error) {
    return c.QueryChatArchivingListWithContext(context.Background(), request)
}

// QueryChatArchivingList
// 根据游标拉取会话存档列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
func (c *Client) QueryChatArchivingListWithContext(ctx context.Context, request *QueryChatArchivingListRequest) (response *QueryChatArchivingListResponse, err error) {
    if request == nil {
        request = NewQueryChatArchivingListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryChatArchivingList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryChatArchivingListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryClueInfoListRequest() (request *QueryClueInfoListRequest) {
    request = &QueryClueInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryClueInfoList")
    
    
    return
}

func NewQueryClueInfoListResponse() (response *QueryClueInfoListResponse) {
    response = &QueryClueInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryClueInfoList
// 企业可通过此接口获取线索列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) QueryClueInfoList(request *QueryClueInfoListRequest) (response *QueryClueInfoListResponse, err error) {
    return c.QueryClueInfoListWithContext(context.Background(), request)
}

// QueryClueInfoList
// 企业可通过此接口获取线索列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) QueryClueInfoListWithContext(ctx context.Context, request *QueryClueInfoListRequest) (response *QueryClueInfoListResponse, err error) {
    if request == nil {
        request = NewQueryClueInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryClueInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryClueInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCrmStatisticsRequest() (request *QueryCrmStatisticsRequest) {
    request = &QueryCrmStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryCrmStatistics")
    
    
    return
}

func NewQueryCrmStatisticsResponse() (response *QueryCrmStatisticsResponse) {
    response = &QueryCrmStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCrmStatistics
// 通过接口拉取租户/指定成员/部门在指定日期范围内的CRM跟进统计数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryCrmStatistics(request *QueryCrmStatisticsRequest) (response *QueryCrmStatisticsResponse, err error) {
    return c.QueryCrmStatisticsWithContext(context.Background(), request)
}

// QueryCrmStatistics
// 通过接口拉取租户/指定成员/部门在指定日期范围内的CRM跟进统计数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryCrmStatisticsWithContext(ctx context.Context, request *QueryCrmStatisticsRequest) (response *QueryCrmStatisticsResponse, err error) {
    if request == nil {
        request = NewQueryCrmStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCrmStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCrmStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCustomerEventDetailStatisticsRequest() (request *QueryCustomerEventDetailStatisticsRequest) {
    request = &QueryCustomerEventDetailStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryCustomerEventDetailStatistics")
    
    
    return
}

func NewQueryCustomerEventDetailStatisticsResponse() (response *QueryCustomerEventDetailStatisticsResponse) {
    response = &QueryCustomerEventDetailStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCustomerEventDetailStatistics
// 通过接口拉取SaaS内C端外部联系人在指定时间范围内的行为事件明细。此接口提供的数据以天为维度，查询的时间范围为[start_time,end_time]，即前后均为闭区间，支持的最大查询跨度为365天。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryCustomerEventDetailStatistics(request *QueryCustomerEventDetailStatisticsRequest) (response *QueryCustomerEventDetailStatisticsResponse, err error) {
    return c.QueryCustomerEventDetailStatisticsWithContext(context.Background(), request)
}

// QueryCustomerEventDetailStatistics
// 通过接口拉取SaaS内C端外部联系人在指定时间范围内的行为事件明细。此接口提供的数据以天为维度，查询的时间范围为[start_time,end_time]，即前后均为闭区间，支持的最大查询跨度为365天。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryCustomerEventDetailStatisticsWithContext(ctx context.Context, request *QueryCustomerEventDetailStatisticsRequest) (response *QueryCustomerEventDetailStatisticsResponse, err error) {
    if request == nil {
        request = NewQueryCustomerEventDetailStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCustomerEventDetailStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCustomerEventDetailStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCustomerProfileListRequest() (request *QueryCustomerProfileListRequest) {
    request = &QueryCustomerProfileListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryCustomerProfileList")
    
    
    return
}

func NewQueryCustomerProfileListResponse() (response *QueryCustomerProfileListResponse) {
    response = &QueryCustomerProfileListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryCustomerProfileList
// 通过接口拉取租户已有潜客客户档案列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) QueryCustomerProfileList(request *QueryCustomerProfileListRequest) (response *QueryCustomerProfileListResponse, err error) {
    return c.QueryCustomerProfileListWithContext(context.Background(), request)
}

// QueryCustomerProfileList
// 通过接口拉取租户已有潜客客户档案列表信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) QueryCustomerProfileListWithContext(ctx context.Context, request *QueryCustomerProfileListRequest) (response *QueryCustomerProfileListResponse, err error) {
    if request == nil {
        request = NewQueryCustomerProfileListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryCustomerProfileList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryCustomerProfileListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDealerInfoListRequest() (request *QueryDealerInfoListRequest) {
    request = &QueryDealerInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryDealerInfoList")
    
    
    return
}

func NewQueryDealerInfoListResponse() (response *QueryDealerInfoListResponse) {
    response = &QueryDealerInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryDealerInfoList
// 企业可通过此接口获取录入在企微SaaS平台上的经销商信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) QueryDealerInfoList(request *QueryDealerInfoListRequest) (response *QueryDealerInfoListResponse, err error) {
    return c.QueryDealerInfoListWithContext(context.Background(), request)
}

// QueryDealerInfoList
// 企业可通过此接口获取录入在企微SaaS平台上的经销商信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) QueryDealerInfoListWithContext(ctx context.Context, request *QueryDealerInfoListRequest) (response *QueryDealerInfoListResponse, err error) {
    if request == nil {
        request = NewQueryDealerInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryDealerInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryDealerInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryExternalContactDetailRequest() (request *QueryExternalContactDetailRequest) {
    request = &QueryExternalContactDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryExternalContactDetail")
    
    
    return
}

func NewQueryExternalContactDetailResponse() (response *QueryExternalContactDetailResponse) {
    response = &QueryExternalContactDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryExternalContactDetail
// 企业可通过此接口，根据外部联系人的userid，拉取外部联系人详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryExternalContactDetail(request *QueryExternalContactDetailRequest) (response *QueryExternalContactDetailResponse, err error) {
    return c.QueryExternalContactDetailWithContext(context.Background(), request)
}

// QueryExternalContactDetail
// 企业可通过此接口，根据外部联系人的userid，拉取外部联系人详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryExternalContactDetailWithContext(ctx context.Context, request *QueryExternalContactDetailRequest) (response *QueryExternalContactDetailResponse, err error) {
    if request == nil {
        request = NewQueryExternalContactDetailRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryExternalContactDetail require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryExternalContactDetailResponse()
    err = c.Send(request, response)
    return
}

func NewQueryExternalContactDetailByDateRequest() (request *QueryExternalContactDetailByDateRequest) {
    request = &QueryExternalContactDetailByDateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryExternalContactDetailByDate")
    
    
    return
}

func NewQueryExternalContactDetailByDateResponse() (response *QueryExternalContactDetailByDateResponse) {
    response = &QueryExternalContactDetailByDateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryExternalContactDetailByDate
// 企业可通过传入起始和结束时间，获取该时间段的外部联系人详情列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryExternalContactDetailByDate(request *QueryExternalContactDetailByDateRequest) (response *QueryExternalContactDetailByDateResponse, err error) {
    return c.QueryExternalContactDetailByDateWithContext(context.Background(), request)
}

// QueryExternalContactDetailByDate
// 企业可通过传入起始和结束时间，获取该时间段的外部联系人详情列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryExternalContactDetailByDateWithContext(ctx context.Context, request *QueryExternalContactDetailByDateRequest) (response *QueryExternalContactDetailByDateResponse, err error) {
    if request == nil {
        request = NewQueryExternalContactDetailByDateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryExternalContactDetailByDate require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryExternalContactDetailByDateResponse()
    err = c.Send(request, response)
    return
}

func NewQueryExternalContactListRequest() (request *QueryExternalContactListRequest) {
    request = &QueryExternalContactListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryExternalContactList")
    
    
    return
}

func NewQueryExternalContactListResponse() (response *QueryExternalContactListResponse) {
    response = &QueryExternalContactListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryExternalContactList
// 企业可通过此接口基于外部联系人获取指定成员添加的客户列表。客户是指配置了客户联系功能的成员所添加的外部联系人。没有配置客户联系功能的成员，所添加的外部联系人将不会作为客户返回。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryExternalContactList(request *QueryExternalContactListRequest) (response *QueryExternalContactListResponse, err error) {
    return c.QueryExternalContactListWithContext(context.Background(), request)
}

// QueryExternalContactList
// 企业可通过此接口基于外部联系人获取指定成员添加的客户列表。客户是指配置了客户联系功能的成员所添加的外部联系人。没有配置客户联系功能的成员，所添加的外部联系人将不会作为客户返回。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryExternalContactListWithContext(ctx context.Context, request *QueryExternalContactListRequest) (response *QueryExternalContactListResponse, err error) {
    if request == nil {
        request = NewQueryExternalContactListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryExternalContactList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryExternalContactListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryExternalUserEventListRequest() (request *QueryExternalUserEventListRequest) {
    request = &QueryExternalUserEventListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryExternalUserEventList")
    
    
    return
}

func NewQueryExternalUserEventListResponse() (response *QueryExternalUserEventListResponse) {
    response = &QueryExternalUserEventListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryExternalUserEventList
// 通过接口拉取租户在指定时间范围内的外部联系人添加/删除明细，此接口提供的数据以天为维度，查询的时间范围为[StarTime, EndTime]，即前后均为闭区间，支持的最大查询跨度为365天；
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  AUTHFAILURE_MISSINGACCESSTOKEN = "AuthFailure.MissingAccessToken"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TIMESPANLIMITEXCEEDED = "InvalidParameterValue.TimeSpanLimitExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryExternalUserEventList(request *QueryExternalUserEventListRequest) (response *QueryExternalUserEventListResponse, err error) {
    return c.QueryExternalUserEventListWithContext(context.Background(), request)
}

// QueryExternalUserEventList
// 通过接口拉取租户在指定时间范围内的外部联系人添加/删除明细，此接口提供的数据以天为维度，查询的时间范围为[StarTime, EndTime]，即前后均为闭区间，支持的最大查询跨度为365天；
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  AUTHFAILURE_MISSINGACCESSTOKEN = "AuthFailure.MissingAccessToken"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TIMESPANLIMITEXCEEDED = "InvalidParameterValue.TimeSpanLimitExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryExternalUserEventListWithContext(ctx context.Context, request *QueryExternalUserEventListRequest) (response *QueryExternalUserEventListResponse, err error) {
    if request == nil {
        request = NewQueryExternalUserEventListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryExternalUserEventList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryExternalUserEventListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryExternalUserMappingInfoRequest() (request *QueryExternalUserMappingInfoRequest) {
    request = &QueryExternalUserMappingInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryExternalUserMappingInfo")
    
    
    return
}

func NewQueryExternalUserMappingInfoResponse() (response *QueryExternalUserMappingInfoResponse) {
    response = &QueryExternalUserMappingInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryExternalUserMappingInfo
// 企业可通过此接口将企业主体对应的外部联系人id转换为乐销车应用主体对应的外部联系人。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryExternalUserMappingInfo(request *QueryExternalUserMappingInfoRequest) (response *QueryExternalUserMappingInfoResponse, err error) {
    return c.QueryExternalUserMappingInfoWithContext(context.Background(), request)
}

// QueryExternalUserMappingInfo
// 企业可通过此接口将企业主体对应的外部联系人id转换为乐销车应用主体对应的外部联系人。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryExternalUserMappingInfoWithContext(ctx context.Context, request *QueryExternalUserMappingInfoRequest) (response *QueryExternalUserMappingInfoResponse, err error) {
    if request == nil {
        request = NewQueryExternalUserMappingInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryExternalUserMappingInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryExternalUserMappingInfoResponse()
    err = c.Send(request, response)
    return
}

func NewQueryLicenseInfoRequest() (request *QueryLicenseInfoRequest) {
    request = &QueryLicenseInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryLicenseInfo")
    
    
    return
}

func NewQueryLicenseInfoResponse() (response *QueryLicenseInfoResponse) {
    response = &QueryLicenseInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryLicenseInfo
// 该接口获取license对应的详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryLicenseInfo(request *QueryLicenseInfoRequest) (response *QueryLicenseInfoResponse, err error) {
    return c.QueryLicenseInfoWithContext(context.Background(), request)
}

// QueryLicenseInfo
// 该接口获取license对应的详细信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryLicenseInfoWithContext(ctx context.Context, request *QueryLicenseInfoRequest) (response *QueryLicenseInfoResponse, err error) {
    if request == nil {
        request = NewQueryLicenseInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryLicenseInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryLicenseInfoResponse()
    err = c.Send(request, response)
    return
}

func NewQueryMaterialListRequest() (request *QueryMaterialListRequest) {
    request = &QueryMaterialListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryMaterialList")
    
    
    return
}

func NewQueryMaterialListResponse() (response *QueryMaterialListResponse) {
    response = &QueryMaterialListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryMaterialList
// 通过接口按类型拉取租户当前的素材列表及关键信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  AUTHFAILURE_MISSINGACCESSTOKEN = "AuthFailure.MissingAccessToken"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TIMESPANLIMITEXCEEDED = "InvalidParameterValue.TimeSpanLimitExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryMaterialList(request *QueryMaterialListRequest) (response *QueryMaterialListResponse, err error) {
    return c.QueryMaterialListWithContext(context.Background(), request)
}

// QueryMaterialList
// 通过接口按类型拉取租户当前的素材列表及关键信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  AUTHFAILURE_MISSINGACCESSTOKEN = "AuthFailure.MissingAccessToken"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_TIMESPANLIMITEXCEEDED = "InvalidParameterValue.TimeSpanLimitExceeded"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
func (c *Client) QueryMaterialListWithContext(ctx context.Context, request *QueryMaterialListRequest) (response *QueryMaterialListResponse, err error) {
    if request == nil {
        request = NewQueryMaterialListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryMaterialList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryMaterialListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryMiniAppCodeListRequest() (request *QueryMiniAppCodeListRequest) {
    request = &QueryMiniAppCodeListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryMiniAppCodeList")
    
    
    return
}

func NewQueryMiniAppCodeListResponse() (response *QueryMiniAppCodeListResponse) {
    response = &QueryMiniAppCodeListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryMiniAppCodeList
// 查询小程序码列表接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  AUTHFAILURE_MISSINGACCESSTOKEN = "AuthFailure.MissingAccessToken"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) QueryMiniAppCodeList(request *QueryMiniAppCodeListRequest) (response *QueryMiniAppCodeListResponse, err error) {
    return c.QueryMiniAppCodeListWithContext(context.Background(), request)
}

// QueryMiniAppCodeList
// 查询小程序码列表接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  AUTHFAILURE_MISSINGACCESSTOKEN = "AuthFailure.MissingAccessToken"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) QueryMiniAppCodeListWithContext(ctx context.Context, request *QueryMiniAppCodeListRequest) (response *QueryMiniAppCodeListResponse, err error) {
    if request == nil {
        request = NewQueryMiniAppCodeListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryMiniAppCodeList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryMiniAppCodeListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryStaffEventDetailStatisticsRequest() (request *QueryStaffEventDetailStatisticsRequest) {
    request = &QueryStaffEventDetailStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryStaffEventDetailStatistics")
    
    
    return
}

func NewQueryStaffEventDetailStatisticsResponse() (response *QueryStaffEventDetailStatisticsResponse) {
    response = &QueryStaffEventDetailStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryStaffEventDetailStatistics
// 通过接口拉取SaaS内企业成员在指定时间范围内的行为事件明细。此接口提供的数据以天为维度，查询的时间范围为[start_time,end_time]，即前后均为闭区间，支持的最大查询跨度为365天。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryStaffEventDetailStatistics(request *QueryStaffEventDetailStatisticsRequest) (response *QueryStaffEventDetailStatisticsResponse, err error) {
    return c.QueryStaffEventDetailStatisticsWithContext(context.Background(), request)
}

// QueryStaffEventDetailStatistics
// 通过接口拉取SaaS内企业成员在指定时间范围内的行为事件明细。此接口提供的数据以天为维度，查询的时间范围为[start_time,end_time]，即前后均为闭区间，支持的最大查询跨度为365天。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) QueryStaffEventDetailStatisticsWithContext(ctx context.Context, request *QueryStaffEventDetailStatisticsRequest) (response *QueryStaffEventDetailStatisticsResponse, err error) {
    if request == nil {
        request = NewQueryStaffEventDetailStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryStaffEventDetailStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryStaffEventDetailStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewQueryUserInfoListRequest() (request *QueryUserInfoListRequest) {
    request = &QueryUserInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryUserInfoList")
    
    
    return
}

func NewQueryUserInfoListResponse() (response *QueryUserInfoListResponse) {
    response = &QueryUserInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryUserInfoList
// 查询企业成员信息列表接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  AUTHFAILURE_MISSINGACCESSTOKEN = "AuthFailure.MissingAccessToken"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) QueryUserInfoList(request *QueryUserInfoListRequest) (response *QueryUserInfoListResponse, err error) {
    return c.QueryUserInfoListWithContext(context.Background(), request)
}

// QueryUserInfoList
// 查询企业成员信息列表接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  AUTHFAILURE_MISSINGACCESSTOKEN = "AuthFailure.MissingAccessToken"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) QueryUserInfoListWithContext(ctx context.Context, request *QueryUserInfoListRequest) (response *QueryUserInfoListResponse, err error) {
    if request == nil {
        request = NewQueryUserInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryUserInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryUserInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryVehicleInfoListRequest() (request *QueryVehicleInfoListRequest) {
    request = &QueryVehicleInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wav", APIVersion, "QueryVehicleInfoList")
    
    
    return
}

func NewQueryVehicleInfoListResponse() (response *QueryVehicleInfoListResponse) {
    response = &QueryVehicleInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryVehicleInfoList
// 企业可通过此接口获取企微SaaS平台上的车系车型信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) QueryVehicleInfoList(request *QueryVehicleInfoListRequest) (response *QueryVehicleInfoListResponse, err error) {
    return c.QueryVehicleInfoListWithContext(context.Background(), request)
}

// QueryVehicleInfoList
// 企业可通过此接口获取企微SaaS平台上的车系车型信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_INVALIDAUTHORIZATIONCODE = "AuthFailure.InvalidAuthorizationCode"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) QueryVehicleInfoListWithContext(ctx context.Context, request *QueryVehicleInfoListRequest) (response *QueryVehicleInfoListResponse, err error) {
    if request == nil {
        request = NewQueryVehicleInfoListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryVehicleInfoList require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryVehicleInfoListResponse()
    err = c.Send(request, response)
    return
}
