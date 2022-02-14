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

package v20200210

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-02-10"

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


func NewBindStaffSkillGroupListRequest() (request *BindStaffSkillGroupListRequest) {
    request = &BindStaffSkillGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "BindStaffSkillGroupList")
    
    
    return
}

func NewBindStaffSkillGroupListResponse() (response *BindStaffSkillGroupListResponse) {
    response = &BindStaffSkillGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindStaffSkillGroupList
// 绑定坐席所属技能组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) BindStaffSkillGroupList(request *BindStaffSkillGroupListRequest) (response *BindStaffSkillGroupListResponse, err error) {
    if request == nil {
        request = NewBindStaffSkillGroupListRequest()
    }
    
    response = NewBindStaffSkillGroupListResponse()
    err = c.Send(request, response)
    return
}

// BindStaffSkillGroupList
// 绑定坐席所属技能组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) BindStaffSkillGroupListWithContext(ctx context.Context, request *BindStaffSkillGroupListRequest) (response *BindStaffSkillGroupListResponse, err error) {
    if request == nil {
        request = NewBindStaffSkillGroupListRequest()
    }
    request.SetContext(ctx)
    
    response = NewBindStaffSkillGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAutoCalloutTaskRequest() (request *CreateAutoCalloutTaskRequest) {
    request = &CreateAutoCalloutTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "CreateAutoCalloutTask")
    
    
    return
}

func NewCreateAutoCalloutTaskResponse() (response *CreateAutoCalloutTaskResponse) {
    response = &CreateAutoCalloutTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAutoCalloutTask
// 创建自动外呼任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateAutoCalloutTask(request *CreateAutoCalloutTaskRequest) (response *CreateAutoCalloutTaskResponse, err error) {
    if request == nil {
        request = NewCreateAutoCalloutTaskRequest()
    }
    
    response = NewCreateAutoCalloutTaskResponse()
    err = c.Send(request, response)
    return
}

// CreateAutoCalloutTask
// 创建自动外呼任务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) CreateAutoCalloutTaskWithContext(ctx context.Context, request *CreateAutoCalloutTaskRequest) (response *CreateAutoCalloutTaskResponse, err error) {
    if request == nil {
        request = NewCreateAutoCalloutTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateAutoCalloutTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCallOutSessionRequest() (request *CreateCallOutSessionRequest) {
    request = &CreateCallOutSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "CreateCallOutSession")
    
    
    return
}

func NewCreateCallOutSessionResponse() (response *CreateCallOutSessionResponse) {
    response = &CreateCallOutSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCallOutSession
// 创建外呼会话
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOUTFAILED = "FailedOperation.CallOutFailed"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_NOTINWHITELIST = "OperationDenied.NotInWhiteList"
func (c *Client) CreateCallOutSession(request *CreateCallOutSessionRequest) (response *CreateCallOutSessionResponse, err error) {
    if request == nil {
        request = NewCreateCallOutSessionRequest()
    }
    
    response = NewCreateCallOutSessionResponse()
    err = c.Send(request, response)
    return
}

// CreateCallOutSession
// 创建外呼会话
//
// 可能返回的错误码:
//  FAILEDOPERATION_CALLOUTFAILED = "FailedOperation.CallOutFailed"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED_NOTINWHITELIST = "OperationDenied.NotInWhiteList"
func (c *Client) CreateCallOutSessionWithContext(ctx context.Context, request *CreateCallOutSessionRequest) (response *CreateCallOutSessionResponse, err error) {
    if request == nil {
        request = NewCreateCallOutSessionRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateCallOutSessionResponse()
    err = c.Send(request, response)
    return
}

func NewCreateSDKLoginTokenRequest() (request *CreateSDKLoginTokenRequest) {
    request = &CreateSDKLoginTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "CreateSDKLoginToken")
    
    
    return
}

func NewCreateSDKLoginTokenResponse() (response *CreateSDKLoginTokenResponse) {
    response = &CreateSDKLoginTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSDKLoginToken
// 创建 SDK 登录 Token。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) CreateSDKLoginToken(request *CreateSDKLoginTokenRequest) (response *CreateSDKLoginTokenResponse, err error) {
    if request == nil {
        request = NewCreateSDKLoginTokenRequest()
    }
    
    response = NewCreateSDKLoginTokenResponse()
    err = c.Send(request, response)
    return
}

// CreateSDKLoginToken
// 创建 SDK 登录 Token。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) CreateSDKLoginTokenWithContext(ctx context.Context, request *CreateSDKLoginTokenRequest) (response *CreateSDKLoginTokenResponse, err error) {
    if request == nil {
        request = NewCreateSDKLoginTokenRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateSDKLoginTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStaffRequest() (request *CreateStaffRequest) {
    request = &CreateStaffRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "CreateStaff")
    
    
    return
}

func NewCreateStaffResponse() (response *CreateStaffResponse) {
    response = &CreateStaffResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateStaff
// 创建客服账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEDACCOUNT = "FailedOperation.DuplicatedAccount"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateStaff(request *CreateStaffRequest) (response *CreateStaffResponse, err error) {
    if request == nil {
        request = NewCreateStaffRequest()
    }
    
    response = NewCreateStaffResponse()
    err = c.Send(request, response)
    return
}

// CreateStaff
// 创建客服账号。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DUPLICATEDACCOUNT = "FailedOperation.DuplicatedAccount"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) CreateStaffWithContext(ctx context.Context, request *CreateStaffRequest) (response *CreateStaffResponse, err error) {
    if request == nil {
        request = NewCreateStaffRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateStaffResponse()
    err = c.Send(request, response)
    return
}

func NewCreateUserSigRequest() (request *CreateUserSigRequest) {
    request = &CreateUserSigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "CreateUserSig")
    
    
    return
}

func NewCreateUserSigResponse() (response *CreateUserSigResponse) {
    response = &CreateUserSigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateUserSig
// 创建用户数据签名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) CreateUserSig(request *CreateUserSigRequest) (response *CreateUserSigResponse, err error) {
    if request == nil {
        request = NewCreateUserSigRequest()
    }
    
    response = NewCreateUserSigResponse()
    err = c.Send(request, response)
    return
}

// CreateUserSig
// 创建用户数据签名
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) CreateUserSigWithContext(ctx context.Context, request *CreateUserSigRequest) (response *CreateUserSigResponse, err error) {
    if request == nil {
        request = NewCreateUserSigRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateUserSigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteStaffRequest() (request *DeleteStaffRequest) {
    request = &DeleteStaffRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "DeleteStaff")
    
    
    return
}

func NewDeleteStaffResponse() (response *DeleteStaffResponse) {
    response = &DeleteStaffResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteStaff
// 删除坐席信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
func (c *Client) DeleteStaff(request *DeleteStaffRequest) (response *DeleteStaffResponse, err error) {
    if request == nil {
        request = NewDeleteStaffRequest()
    }
    
    response = NewDeleteStaffResponse()
    err = c.Send(request, response)
    return
}

// DeleteStaff
// 删除坐席信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
func (c *Client) DeleteStaffWithContext(ctx context.Context, request *DeleteStaffRequest) (response *DeleteStaffResponse, err error) {
    if request == nil {
        request = NewDeleteStaffRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteStaffResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoCalloutTaskRequest() (request *DescribeAutoCalloutTaskRequest) {
    request = &DescribeAutoCalloutTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeAutoCalloutTask")
    
    
    return
}

func NewDescribeAutoCalloutTaskResponse() (response *DescribeAutoCalloutTaskResponse) {
    response = &DescribeAutoCalloutTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAutoCalloutTask
// 查询自动外呼任务详情
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAutoCalloutTask(request *DescribeAutoCalloutTaskRequest) (response *DescribeAutoCalloutTaskResponse, err error) {
    if request == nil {
        request = NewDescribeAutoCalloutTaskRequest()
    }
    
    response = NewDescribeAutoCalloutTaskResponse()
    err = c.Send(request, response)
    return
}

// DescribeAutoCalloutTask
// 查询自动外呼任务详情
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAutoCalloutTaskWithContext(ctx context.Context, request *DescribeAutoCalloutTaskRequest) (response *DescribeAutoCalloutTaskResponse, err error) {
    if request == nil {
        request = NewDescribeAutoCalloutTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAutoCalloutTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAutoCalloutTasksRequest() (request *DescribeAutoCalloutTasksRequest) {
    request = &DescribeAutoCalloutTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeAutoCalloutTasks")
    
    
    return
}

func NewDescribeAutoCalloutTasksResponse() (response *DescribeAutoCalloutTasksResponse) {
    response = &DescribeAutoCalloutTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAutoCalloutTasks
// 批量查询自动任务外呼
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAutoCalloutTasks(request *DescribeAutoCalloutTasksRequest) (response *DescribeAutoCalloutTasksResponse, err error) {
    if request == nil {
        request = NewDescribeAutoCalloutTasksRequest()
    }
    
    response = NewDescribeAutoCalloutTasksResponse()
    err = c.Send(request, response)
    return
}

// DescribeAutoCalloutTasks
// 批量查询自动任务外呼
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeAutoCalloutTasksWithContext(ctx context.Context, request *DescribeAutoCalloutTasksRequest) (response *DescribeAutoCalloutTasksResponse, err error) {
    if request == nil {
        request = NewDescribeAutoCalloutTasksRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAutoCalloutTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCCCBuyInfoListRequest() (request *DescribeCCCBuyInfoListRequest) {
    request = &DescribeCCCBuyInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeCCCBuyInfoList")
    
    
    return
}

func NewDescribeCCCBuyInfoListResponse() (response *DescribeCCCBuyInfoListResponse) {
    response = &DescribeCCCBuyInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCCCBuyInfoList
// 获取用户购买信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCCCBuyInfoList(request *DescribeCCCBuyInfoListRequest) (response *DescribeCCCBuyInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeCCCBuyInfoListRequest()
    }
    
    response = NewDescribeCCCBuyInfoListResponse()
    err = c.Send(request, response)
    return
}

// DescribeCCCBuyInfoList
// 获取用户购买信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTNOTEXIST = "InvalidParameterValue.AccountNotExist"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeCCCBuyInfoListWithContext(ctx context.Context, request *DescribeCCCBuyInfoListRequest) (response *DescribeCCCBuyInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeCCCBuyInfoListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCCCBuyInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCallInMetricsRequest() (request *DescribeCallInMetricsRequest) {
    request = &DescribeCallInMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeCallInMetrics")
    
    
    return
}

func NewDescribeCallInMetricsResponse() (response *DescribeCallInMetricsResponse) {
    response = &DescribeCallInMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeCallInMetrics
// 获取呼入实时数据统计指标
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeCallInMetrics(request *DescribeCallInMetricsRequest) (response *DescribeCallInMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeCallInMetricsRequest()
    }
    
    response = NewDescribeCallInMetricsResponse()
    err = c.Send(request, response)
    return
}

// DescribeCallInMetrics
// 获取呼入实时数据统计指标
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeCallInMetricsWithContext(ctx context.Context, request *DescribeCallInMetricsRequest) (response *DescribeCallInMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeCallInMetricsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeCallInMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChatMessagesRequest() (request *DescribeChatMessagesRequest) {
    request = &DescribeChatMessagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeChatMessages")
    
    
    return
}

func NewDescribeChatMessagesResponse() (response *DescribeChatMessagesResponse) {
    response = &DescribeChatMessagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeChatMessages
// 包括具体聊天内容
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
func (c *Client) DescribeChatMessages(request *DescribeChatMessagesRequest) (response *DescribeChatMessagesResponse, err error) {
    if request == nil {
        request = NewDescribeChatMessagesRequest()
    }
    
    response = NewDescribeChatMessagesResponse()
    err = c.Send(request, response)
    return
}

// DescribeChatMessages
// 包括具体聊天内容
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
//  INVALIDPARAMETERVALUE_RECORDNOTEXIST = "InvalidParameterValue.RecordNotExist"
func (c *Client) DescribeChatMessagesWithContext(ctx context.Context, request *DescribeChatMessagesRequest) (response *DescribeChatMessagesResponse, err error) {
    if request == nil {
        request = NewDescribeChatMessagesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeChatMessagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIMCdrsRequest() (request *DescribeIMCdrsRequest) {
    request = &DescribeIMCdrsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeIMCdrs")
    
    
    return
}

func NewDescribeIMCdrsResponse() (response *DescribeIMCdrsResponse) {
    response = &DescribeIMCdrsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeIMCdrs
// 包括全媒体和文本两种类型
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIMCdrs(request *DescribeIMCdrsRequest) (response *DescribeIMCdrsResponse, err error) {
    if request == nil {
        request = NewDescribeIMCdrsRequest()
    }
    
    response = NewDescribeIMCdrsResponse()
    err = c.Send(request, response)
    return
}

// DescribeIMCdrs
// 包括全媒体和文本两种类型
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeIMCdrsWithContext(ctx context.Context, request *DescribeIMCdrsRequest) (response *DescribeIMCdrsResponse, err error) {
    if request == nil {
        request = NewDescribeIMCdrsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeIMCdrsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePSTNActiveSessionListRequest() (request *DescribePSTNActiveSessionListRequest) {
    request = &DescribePSTNActiveSessionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "DescribePSTNActiveSessionList")
    
    
    return
}

func NewDescribePSTNActiveSessionListResponse() (response *DescribePSTNActiveSessionListResponse) {
    response = &DescribePSTNActiveSessionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribePSTNActiveSessionList
// 获取当前正在通话的会话列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_OUTOFCOUNTLIMIT = "LimitExceeded.OutOfCountLimit"
func (c *Client) DescribePSTNActiveSessionList(request *DescribePSTNActiveSessionListRequest) (response *DescribePSTNActiveSessionListResponse, err error) {
    if request == nil {
        request = NewDescribePSTNActiveSessionListRequest()
    }
    
    response = NewDescribePSTNActiveSessionListResponse()
    err = c.Send(request, response)
    return
}

// DescribePSTNActiveSessionList
// 获取当前正在通话的会话列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  LIMITEXCEEDED_OUTOFCOUNTLIMIT = "LimitExceeded.OutOfCountLimit"
func (c *Client) DescribePSTNActiveSessionListWithContext(ctx context.Context, request *DescribePSTNActiveSessionListRequest) (response *DescribePSTNActiveSessionListResponse, err error) {
    if request == nil {
        request = NewDescribePSTNActiveSessionListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribePSTNActiveSessionListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSeatUserListRequest() (request *DescribeSeatUserListRequest) {
    request = &DescribeSeatUserListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeSeatUserList")
    
    
    return
}

func NewDescribeSeatUserListResponse() (response *DescribeSeatUserListResponse) {
    response = &DescribeSeatUserListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSeatUserList
// 废弃接口下架
//
// 
//
// 获取坐席用户列表（废弃）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeSeatUserList(request *DescribeSeatUserListRequest) (response *DescribeSeatUserListResponse, err error) {
    if request == nil {
        request = NewDescribeSeatUserListRequest()
    }
    
    response = NewDescribeSeatUserListResponse()
    err = c.Send(request, response)
    return
}

// DescribeSeatUserList
// 废弃接口下架
//
// 
//
// 获取坐席用户列表（废弃）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeSeatUserListWithContext(ctx context.Context, request *DescribeSeatUserListRequest) (response *DescribeSeatUserListResponse, err error) {
    if request == nil {
        request = NewDescribeSeatUserListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeSeatUserListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSkillGroupInfoListRequest() (request *DescribeSkillGroupInfoListRequest) {
    request = &DescribeSkillGroupInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeSkillGroupInfoList")
    
    
    return
}

func NewDescribeSkillGroupInfoListResponse() (response *DescribeSkillGroupInfoListResponse) {
    response = &DescribeSkillGroupInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSkillGroupInfoList
// 获取技能组信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeSkillGroupInfoList(request *DescribeSkillGroupInfoListRequest) (response *DescribeSkillGroupInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeSkillGroupInfoListRequest()
    }
    
    response = NewDescribeSkillGroupInfoListResponse()
    err = c.Send(request, response)
    return
}

// DescribeSkillGroupInfoList
// 获取技能组信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeSkillGroupInfoListWithContext(ctx context.Context, request *DescribeSkillGroupInfoListRequest) (response *DescribeSkillGroupInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeSkillGroupInfoListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeSkillGroupInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStaffInfoListRequest() (request *DescribeStaffInfoListRequest) {
    request = &DescribeStaffInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeStaffInfoList")
    
    
    return
}

func NewDescribeStaffInfoListResponse() (response *DescribeStaffInfoListResponse) {
    response = &DescribeStaffInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStaffInfoList
// 获取坐席信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeStaffInfoList(request *DescribeStaffInfoListRequest) (response *DescribeStaffInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStaffInfoListRequest()
    }
    
    response = NewDescribeStaffInfoListResponse()
    err = c.Send(request, response)
    return
}

// DescribeStaffInfoList
// 获取坐席信息列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeStaffInfoListWithContext(ctx context.Context, request *DescribeStaffInfoListRequest) (response *DescribeStaffInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStaffInfoListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeStaffInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStaffStatusMetricsRequest() (request *DescribeStaffStatusMetricsRequest) {
    request = &DescribeStaffStatusMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeStaffStatusMetrics")
    
    
    return
}

func NewDescribeStaffStatusMetricsResponse() (response *DescribeStaffStatusMetricsResponse) {
    response = &DescribeStaffStatusMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStaffStatusMetrics
// 获取坐席实时状态统计指标
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
func (c *Client) DescribeStaffStatusMetrics(request *DescribeStaffStatusMetricsRequest) (response *DescribeStaffStatusMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeStaffStatusMetricsRequest()
    }
    
    response = NewDescribeStaffStatusMetricsResponse()
    err = c.Send(request, response)
    return
}

// DescribeStaffStatusMetrics
// 获取坐席实时状态统计指标
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
func (c *Client) DescribeStaffStatusMetricsWithContext(ctx context.Context, request *DescribeStaffStatusMetricsRequest) (response *DescribeStaffStatusMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeStaffStatusMetricsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeStaffStatusMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTelCallInfoRequest() (request *DescribeTelCallInfoRequest) {
    request = &DescribeTelCallInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeTelCallInfo")
    
    
    return
}

func NewDescribeTelCallInfoResponse() (response *DescribeTelCallInfoResponse) {
    response = &DescribeTelCallInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTelCallInfo
// 按实例获取电话消耗统计
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeTelCallInfo(request *DescribeTelCallInfoRequest) (response *DescribeTelCallInfoResponse, err error) {
    if request == nil {
        request = NewDescribeTelCallInfoRequest()
    }
    
    response = NewDescribeTelCallInfoResponse()
    err = c.Send(request, response)
    return
}

// DescribeTelCallInfo
// 按实例获取电话消耗统计
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeTelCallInfoWithContext(ctx context.Context, request *DescribeTelCallInfoRequest) (response *DescribeTelCallInfoResponse, err error) {
    if request == nil {
        request = NewDescribeTelCallInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTelCallInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTelCdrRequest() (request *DescribeTelCdrRequest) {
    request = &DescribeTelCdrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeTelCdr")
    
    
    return
}

func NewDescribeTelCdrResponse() (response *DescribeTelCdrResponse) {
    response = &DescribeTelCdrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTelCdr
// 获取电话服务记录与录音
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeTelCdr(request *DescribeTelCdrRequest) (response *DescribeTelCdrResponse, err error) {
    if request == nil {
        request = NewDescribeTelCdrRequest()
    }
    
    response = NewDescribeTelCdrResponse()
    err = c.Send(request, response)
    return
}

// DescribeTelCdr
// 获取电话服务记录与录音
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeTelCdrWithContext(ctx context.Context, request *DescribeTelCdrRequest) (response *DescribeTelCdrResponse, err error) {
    if request == nil {
        request = NewDescribeTelCdrRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTelCdrResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTelSessionRequest() (request *DescribeTelSessionRequest) {
    request = &DescribeTelSessionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeTelSession")
    
    
    return
}

func NewDescribeTelSessionResponse() (response *DescribeTelSessionResponse) {
    response = &DescribeTelSessionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTelSession
// 获取 PSTN 会话信息
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeTelSession(request *DescribeTelSessionRequest) (response *DescribeTelSessionResponse, err error) {
    if request == nil {
        request = NewDescribeTelSessionRequest()
    }
    
    response = NewDescribeTelSessionResponse()
    err = c.Send(request, response)
    return
}

// DescribeTelSession
// 获取 PSTN 会话信息
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_INSTANCENOTEXIST = "InvalidParameter.InstanceNotExist"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) DescribeTelSessionWithContext(ctx context.Context, request *DescribeTelSessionRequest) (response *DescribeTelSessionResponse, err error) {
    if request == nil {
        request = NewDescribeTelSessionRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeTelSessionResponse()
    err = c.Send(request, response)
    return
}

func NewStopAutoCalloutTaskRequest() (request *StopAutoCalloutTaskRequest) {
    request = &StopAutoCalloutTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "StopAutoCalloutTask")
    
    
    return
}

func NewStopAutoCalloutTaskResponse() (response *StopAutoCalloutTaskResponse) {
    response = &StopAutoCalloutTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopAutoCalloutTask
// 停止自动外呼任务
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopAutoCalloutTask(request *StopAutoCalloutTaskRequest) (response *StopAutoCalloutTaskResponse, err error) {
    if request == nil {
        request = NewStopAutoCalloutTaskRequest()
    }
    
    response = NewStopAutoCalloutTaskResponse()
    err = c.Send(request, response)
    return
}

// StopAutoCalloutTask
// 停止自动外呼任务
//
// 可能返回的错误码:
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StopAutoCalloutTaskWithContext(ctx context.Context, request *StopAutoCalloutTaskRequest) (response *StopAutoCalloutTaskResponse, err error) {
    if request == nil {
        request = NewStopAutoCalloutTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewStopAutoCalloutTaskResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindStaffSkillGroupListRequest() (request *UnbindStaffSkillGroupListRequest) {
    request = &UnbindStaffSkillGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "UnbindStaffSkillGroupList")
    
    
    return
}

func NewUnbindStaffSkillGroupListResponse() (response *UnbindStaffSkillGroupListResponse) {
    response = &UnbindStaffSkillGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UnbindStaffSkillGroupList
// 解绑坐席所属技能组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) UnbindStaffSkillGroupList(request *UnbindStaffSkillGroupListRequest) (response *UnbindStaffSkillGroupListResponse, err error) {
    if request == nil {
        request = NewUnbindStaffSkillGroupListRequest()
    }
    
    response = NewUnbindStaffSkillGroupListResponse()
    err = c.Send(request, response)
    return
}

// UnbindStaffSkillGroupList
// 解绑坐席所属技能组
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INSTANCENOTEXIST = "InvalidParameterValue.InstanceNotExist"
func (c *Client) UnbindStaffSkillGroupListWithContext(ctx context.Context, request *UnbindStaffSkillGroupListRequest) (response *UnbindStaffSkillGroupListResponse, err error) {
    if request == nil {
        request = NewUnbindStaffSkillGroupListRequest()
    }
    request.SetContext(ctx)
    
    response = NewUnbindStaffSkillGroupListResponse()
    err = c.Send(request, response)
    return
}
