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
    if request == nil {
        request = NewCreateChannelCodeRequest()
    }
    
    response = NewCreateChannelCodeResponse()
    err = c.Send(request, response)
    return
}

// CreateChannelCode
// 新增渠道活码接口
func (c *Client) CreateChannelCodeWithContext(ctx context.Context, request *CreateChannelCodeRequest) (response *CreateChannelCodeResponse, err error) {
    if request == nil {
        request = NewCreateChannelCodeRequest()
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
    if request == nil {
        request = NewCreateCorpTagRequest()
    }
    
    response = NewCreateCorpTagResponse()
    err = c.Send(request, response)
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
func (c *Client) CreateCorpTagWithContext(ctx context.Context, request *CreateCorpTagRequest) (response *CreateCorpTagResponse, err error) {
    if request == nil {
        request = NewCreateCorpTagRequest()
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
    if request == nil {
        request = NewCreateLeadRequest()
    }
    
    response = NewCreateLeadResponse()
    err = c.Send(request, response)
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
func (c *Client) CreateLeadWithContext(ctx context.Context, request *CreateLeadRequest) (response *CreateLeadResponse, err error) {
    if request == nil {
        request = NewCreateLeadRequest()
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
    if request == nil {
        request = NewQueryActivityJoinListRequest()
    }
    
    response = NewQueryActivityJoinListResponse()
    err = c.Send(request, response)
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
func (c *Client) QueryActivityJoinListWithContext(ctx context.Context, request *QueryActivityJoinListRequest) (response *QueryActivityJoinListResponse, err error) {
    if request == nil {
        request = NewQueryActivityJoinListRequest()
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
    if request == nil {
        request = NewQueryActivityListRequest()
    }
    
    response = NewQueryActivityListResponse()
    err = c.Send(request, response)
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
func (c *Client) QueryActivityListWithContext(ctx context.Context, request *QueryActivityListRequest) (response *QueryActivityListResponse, err error) {
    if request == nil {
        request = NewQueryActivityListRequest()
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
    if request == nil {
        request = NewQueryActivityLiveCodeListRequest()
    }
    
    response = NewQueryActivityLiveCodeListResponse()
    err = c.Send(request, response)
    return
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
    if request == nil {
        request = NewQueryChannelCodeListRequest()
    }
    
    response = NewQueryChannelCodeListResponse()
    err = c.Send(request, response)
    return
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
    if request == nil {
        request = NewQueryChatArchivingListRequest()
    }
    
    response = NewQueryChatArchivingListResponse()
    err = c.Send(request, response)
    return
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
    if request == nil {
        request = NewQueryClueInfoListRequest()
    }
    
    response = NewQueryClueInfoListResponse()
    err = c.Send(request, response)
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
func (c *Client) QueryClueInfoListWithContext(ctx context.Context, request *QueryClueInfoListRequest) (response *QueryClueInfoListResponse, err error) {
    if request == nil {
        request = NewQueryClueInfoListRequest()
    }
    request.SetContext(ctx)
    
    response = NewQueryClueInfoListResponse()
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
    if request == nil {
        request = NewQueryDealerInfoListRequest()
    }
    
    response = NewQueryDealerInfoListResponse()
    err = c.Send(request, response)
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
func (c *Client) QueryDealerInfoListWithContext(ctx context.Context, request *QueryDealerInfoListRequest) (response *QueryDealerInfoListResponse, err error) {
    if request == nil {
        request = NewQueryDealerInfoListRequest()
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
    if request == nil {
        request = NewQueryExternalContactDetailRequest()
    }
    
    response = NewQueryExternalContactDetailResponse()
    err = c.Send(request, response)
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
func (c *Client) QueryExternalContactDetailWithContext(ctx context.Context, request *QueryExternalContactDetailRequest) (response *QueryExternalContactDetailResponse, err error) {
    if request == nil {
        request = NewQueryExternalContactDetailRequest()
    }
    request.SetContext(ctx)
    
    response = NewQueryExternalContactDetailResponse()
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
    if request == nil {
        request = NewQueryExternalContactListRequest()
    }
    
    response = NewQueryExternalContactListResponse()
    err = c.Send(request, response)
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
func (c *Client) QueryExternalContactListWithContext(ctx context.Context, request *QueryExternalContactListRequest) (response *QueryExternalContactListResponse, err error) {
    if request == nil {
        request = NewQueryExternalContactListRequest()
    }
    request.SetContext(ctx)
    
    response = NewQueryExternalContactListResponse()
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
    if request == nil {
        request = NewQueryExternalUserMappingInfoRequest()
    }
    
    response = NewQueryExternalUserMappingInfoResponse()
    err = c.Send(request, response)
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
func (c *Client) QueryExternalUserMappingInfoWithContext(ctx context.Context, request *QueryExternalUserMappingInfoRequest) (response *QueryExternalUserMappingInfoResponse, err error) {
    if request == nil {
        request = NewQueryExternalUserMappingInfoRequest()
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
    if request == nil {
        request = NewQueryLicenseInfoRequest()
    }
    
    response = NewQueryLicenseInfoResponse()
    err = c.Send(request, response)
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
func (c *Client) QueryLicenseInfoWithContext(ctx context.Context, request *QueryLicenseInfoRequest) (response *QueryLicenseInfoResponse, err error) {
    if request == nil {
        request = NewQueryLicenseInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewQueryLicenseInfoResponse()
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
    if request == nil {
        request = NewQueryMiniAppCodeListRequest()
    }
    
    response = NewQueryMiniAppCodeListResponse()
    err = c.Send(request, response)
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
func (c *Client) QueryMiniAppCodeListWithContext(ctx context.Context, request *QueryMiniAppCodeListRequest) (response *QueryMiniAppCodeListResponse, err error) {
    if request == nil {
        request = NewQueryMiniAppCodeListRequest()
    }
    request.SetContext(ctx)
    
    response = NewQueryMiniAppCodeListResponse()
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
    if request == nil {
        request = NewQueryVehicleInfoListRequest()
    }
    
    response = NewQueryVehicleInfoListResponse()
    err = c.Send(request, response)
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
func (c *Client) QueryVehicleInfoListWithContext(ctx context.Context, request *QueryVehicleInfoListRequest) (response *QueryVehicleInfoListResponse, err error) {
    if request == nil {
        request = NewQueryVehicleInfoListRequest()
    }
    request.SetContext(ctx)
    
    response = NewQueryVehicleInfoListResponse()
    err = c.Send(request, response)
    return
}
