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

package v20190121

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-01-21"

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


func NewAddCrowdPackInfoRequest() (request *AddCrowdPackInfoRequest) {
    request = &AddCrowdPackInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "AddCrowdPackInfo")
    
    
    return
}

func NewAddCrowdPackInfoResponse() (response *AddCrowdPackInfoResponse) {
    response = &AddCrowdPackInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddCrowdPackInfo
// 添加短信人群包信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddCrowdPackInfo(request *AddCrowdPackInfoRequest) (response *AddCrowdPackInfoResponse, err error) {
    if request == nil {
        request = NewAddCrowdPackInfoRequest()
    }
    
    response = NewAddCrowdPackInfoResponse()
    err = c.Send(request, response)
    return
}

func NewAddSmsSignRequest() (request *AddSmsSignRequest) {
    request = &AddSmsSignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "AddSmsSign")
    
    
    return
}

func NewAddSmsSignResponse() (response *AddSmsSignResponse) {
    response = &AddSmsSignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddSmsSign
// 创建普通短信签名信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_ROLEFAILURE = "UnauthorizedOperation.RoleFailure"
func (c *Client) AddSmsSign(request *AddSmsSignRequest) (response *AddSmsSignResponse, err error) {
    if request == nil {
        request = NewAddSmsSignRequest()
    }
    
    response = NewAddSmsSignResponse()
    err = c.Send(request, response)
    return
}

func NewAddSmsTemplateRequest() (request *AddSmsTemplateRequest) {
    request = &AddSmsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "AddSmsTemplate")
    
    
    return
}

func NewAddSmsTemplateResponse() (response *AddSmsTemplateResponse) {
    response = &AddSmsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AddSmsTemplate
// 根据短信标题、模板内容等创建短信模板
//
// 可能返回的错误码:
//  FAILEDOPERATION_SMSTEMPLATEEXISTS = "FailedOperation.SmsTemplateExists"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_ROLEFAILURE = "UnauthorizedOperation.RoleFailure"
func (c *Client) AddSmsTemplate(request *AddSmsTemplateRequest) (response *AddSmsTemplateResponse, err error) {
    if request == nil {
        request = NewAddSmsTemplateRequest()
    }
    
    response = NewAddSmsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewCancelCampaignRequest() (request *CancelCampaignRequest) {
    request = &CancelCampaignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "CancelCampaign")
    
    
    return
}

func NewCancelCampaignResponse() (response *CancelCampaignResponse) {
    response = &CancelCampaignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelCampaign
// 取消短信推送活动
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_ROLEFAILURE = "UnauthorizedOperation.RoleFailure"
func (c *Client) CancelCampaign(request *CancelCampaignRequest) (response *CancelCampaignResponse, err error) {
    if request == nil {
        request = NewCancelCampaignRequest()
    }
    
    response = NewCancelCampaignResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCampaignRequest() (request *CreateCampaignRequest) {
    request = &CreateCampaignRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "CreateCampaign")
    
    
    return
}

func NewCreateCampaignResponse() (response *CreateCampaignResponse) {
    response = &CreateCampaignResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCampaign
// 创建短信推送活动
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION_ROLEFAILURE = "UnauthorizedOperation.RoleFailure"
func (c *Client) CreateCampaign(request *CreateCampaignRequest) (response *CreateCampaignResponse, err error) {
    if request == nil {
        request = NewCreateCampaignRequest()
    }
    
    response = NewCreateCampaignResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMmsInstanceRequest() (request *CreateMmsInstanceRequest) {
    request = &CreateMmsInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "CreateMmsInstance")
    
    
    return
}

func NewCreateMmsInstanceResponse() (response *CreateMmsInstanceResponse) {
    response = &CreateMmsInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateMmsInstance
// 创建超级短信的素材样例内容
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_ROLEFAILURE = "UnauthorizedOperation.RoleFailure"
func (c *Client) CreateMmsInstance(request *CreateMmsInstanceRequest) (response *CreateMmsInstanceResponse, err error) {
    if request == nil {
        request = NewCreateMmsInstanceRequest()
    }
    
    response = NewCreateMmsInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDelCrowdPackRequest() (request *DelCrowdPackRequest) {
    request = &DelCrowdPackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "DelCrowdPack")
    
    
    return
}

func NewDelCrowdPackResponse() (response *DelCrowdPackResponse) {
    response = &DelCrowdPackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DelCrowdPack
// 删除人群包
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DelCrowdPack(request *DelCrowdPackRequest) (response *DelCrowdPackResponse, err error) {
    if request == nil {
        request = NewDelCrowdPackRequest()
    }
    
    response = NewDelCrowdPackResponse()
    err = c.Send(request, response)
    return
}

func NewDelTemplateRequest() (request *DelTemplateRequest) {
    request = &DelTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "DelTemplate")
    
    
    return
}

func NewDelTemplateResponse() (response *DelTemplateResponse) {
    response = &DelTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DelTemplate
// 删除短信模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DelTemplate(request *DelTemplateRequest) (response *DelTemplateResponse, err error) {
    if request == nil {
        request = NewDelTemplateRequest()
    }
    
    response = NewDelTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMmsInstanceRequest() (request *DeleteMmsInstanceRequest) {
    request = &DeleteMmsInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "DeleteMmsInstance")
    
    
    return
}

func NewDeleteMmsInstanceResponse() (response *DeleteMmsInstanceResponse) {
    response = &DeleteMmsInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteMmsInstance
// 删除超级短信样例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_ROLEFAILURE = "UnauthorizedOperation.RoleFailure"
func (c *Client) DeleteMmsInstance(request *DeleteMmsInstanceRequest) (response *DeleteMmsInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteMmsInstanceRequest()
    }
    
    response = NewDeleteMmsInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMmsInstanceInfoRequest() (request *DescribeMmsInstanceInfoRequest) {
    request = &DescribeMmsInstanceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "DescribeMmsInstanceInfo")
    
    
    return
}

func NewDescribeMmsInstanceInfoResponse() (response *DescribeMmsInstanceInfoResponse) {
    response = &DescribeMmsInstanceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMmsInstanceInfo
// 获取彩信实例信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMmsInstanceInfo(request *DescribeMmsInstanceInfoRequest) (response *DescribeMmsInstanceInfoResponse, err error) {
    if request == nil {
        request = NewDescribeMmsInstanceInfoRequest()
    }
    
    response = NewDescribeMmsInstanceInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMmsInstanceListRequest() (request *DescribeMmsInstanceListRequest) {
    request = &DescribeMmsInstanceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "DescribeMmsInstanceList")
    
    
    return
}

func NewDescribeMmsInstanceListResponse() (response *DescribeMmsInstanceListResponse) {
    response = &DescribeMmsInstanceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMmsInstanceList
// 获取彩信实例列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMmsInstanceList(request *DescribeMmsInstanceListRequest) (response *DescribeMmsInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeMmsInstanceListRequest()
    }
    
    response = NewDescribeMmsInstanceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSmsCampaignStatisticsRequest() (request *DescribeSmsCampaignStatisticsRequest) {
    request = &DescribeSmsCampaignStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "DescribeSmsCampaignStatistics")
    
    
    return
}

func NewDescribeSmsCampaignStatisticsResponse() (response *DescribeSmsCampaignStatisticsResponse) {
    response = &DescribeSmsCampaignStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSmsCampaignStatistics
// 获取短信超短活动统计数据
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_ROLEFAILURE = "UnauthorizedOperation.RoleFailure"
func (c *Client) DescribeSmsCampaignStatistics(request *DescribeSmsCampaignStatisticsRequest) (response *DescribeSmsCampaignStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeSmsCampaignStatisticsRequest()
    }
    
    response = NewDescribeSmsCampaignStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSmsSignListRequest() (request *DescribeSmsSignListRequest) {
    request = &DescribeSmsSignListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "DescribeSmsSignList")
    
    
    return
}

func NewDescribeSmsSignListResponse() (response *DescribeSmsSignListResponse) {
    response = &DescribeSmsSignListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSmsSignList
// 获取普通短信签名信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_ROLEFAILURE = "UnauthorizedOperation.RoleFailure"
func (c *Client) DescribeSmsSignList(request *DescribeSmsSignListRequest) (response *DescribeSmsSignListResponse, err error) {
    if request == nil {
        request = NewDescribeSmsSignListRequest()
    }
    
    response = NewDescribeSmsSignListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSmsTemplateListRequest() (request *DescribeSmsTemplateListRequest) {
    request = &DescribeSmsTemplateListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "DescribeSmsTemplateList")
    
    
    return
}

func NewDescribeSmsTemplateListResponse() (response *DescribeSmsTemplateListResponse) {
    response = &DescribeSmsTemplateListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSmsTemplateList
// 获取模板信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_ROLEFAILURE = "UnauthorizedOperation.RoleFailure"
func (c *Client) DescribeSmsTemplateList(request *DescribeSmsTemplateListRequest) (response *DescribeSmsTemplateListResponse, err error) {
    if request == nil {
        request = NewDescribeSmsTemplateListRequest()
    }
    
    response = NewDescribeSmsTemplateListResponse()
    err = c.Send(request, response)
    return
}

func NewGetCrowdPackListRequest() (request *GetCrowdPackListRequest) {
    request = &GetCrowdPackListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "GetCrowdPackList")
    
    
    return
}

func NewGetCrowdPackListResponse() (response *GetCrowdPackListResponse) {
    response = &GetCrowdPackListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetCrowdPackList
// 获取人群包列表接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetCrowdPackList(request *GetCrowdPackListRequest) (response *GetCrowdPackListResponse, err error) {
    if request == nil {
        request = NewGetCrowdPackListRequest()
    }
    
    response = NewGetCrowdPackListResponse()
    err = c.Send(request, response)
    return
}

func NewGetCrowdUploadInfoRequest() (request *GetCrowdUploadInfoRequest) {
    request = &GetCrowdUploadInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "GetCrowdUploadInfo")
    
    
    return
}

func NewGetCrowdUploadInfoResponse() (response *GetCrowdUploadInfoResponse) {
    response = &GetCrowdUploadInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetCrowdUploadInfo
// 获取短信人群包cos上传需要的信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetCrowdUploadInfo(request *GetCrowdUploadInfoRequest) (response *GetCrowdUploadInfoResponse, err error) {
    if request == nil {
        request = NewGetCrowdUploadInfoRequest()
    }
    
    response = NewGetCrowdUploadInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetSmsAmountInfoRequest() (request *GetSmsAmountInfoRequest) {
    request = &GetSmsAmountInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "GetSmsAmountInfo")
    
    
    return
}

func NewGetSmsAmountInfoResponse() (response *GetSmsAmountInfoResponse) {
    response = &GetSmsAmountInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetSmsAmountInfo
// 获取账号短信额度配置信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetSmsAmountInfo(request *GetSmsAmountInfoRequest) (response *GetSmsAmountInfoResponse, err error) {
    if request == nil {
        request = NewGetSmsAmountInfoRequest()
    }
    
    response = NewGetSmsAmountInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetSmsCampaignStatusRequest() (request *GetSmsCampaignStatusRequest) {
    request = &GetSmsCampaignStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "GetSmsCampaignStatus")
    
    
    return
}

func NewGetSmsCampaignStatusResponse() (response *GetSmsCampaignStatusResponse) {
    response = &GetSmsCampaignStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetSmsCampaignStatus
// 获取短信活动状态信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetSmsCampaignStatus(request *GetSmsCampaignStatusRequest) (response *GetSmsCampaignStatusResponse, err error) {
    if request == nil {
        request = NewGetSmsCampaignStatusRequest()
    }
    
    response = NewGetSmsCampaignStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifySmsTemplateRequest() (request *ModifySmsTemplateRequest) {
    request = &ModifySmsTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "ModifySmsTemplate")
    
    
    return
}

func NewModifySmsTemplateResponse() (response *ModifySmsTemplateResponse) {
    response = &ModifySmsTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifySmsTemplate
// 对未审核或者审核未通过的短信模板内容进行编辑修改
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SMSTEMPLATEEXISTS = "FailedOperation.SmsTemplateExists"
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
//  UNAUTHORIZEDOPERATION_ROLEFAILURE = "UnauthorizedOperation.RoleFailure"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifySmsTemplate(request *ModifySmsTemplateRequest) (response *ModifySmsTemplateResponse, err error) {
    if request == nil {
        request = NewModifySmsTemplateRequest()
    }
    
    response = NewModifySmsTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewPushMmsContentRequest() (request *PushMmsContentRequest) {
    request = &PushMmsContentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "PushMmsContent")
    
    
    return
}

func NewPushMmsContentResponse() (response *PushMmsContentResponse) {
    response = &PushMmsContentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PushMmsContent
// 推送超级短信
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_ROLEFAILURE = "UnauthorizedOperation.RoleFailure"
func (c *Client) PushMmsContent(request *PushMmsContentRequest) (response *PushMmsContentResponse, err error) {
    if request == nil {
        request = NewPushMmsContentRequest()
    }
    
    response = NewPushMmsContentResponse()
    err = c.Send(request, response)
    return
}

func NewSendSmsRequest() (request *SendSmsRequest) {
    request = &SendSmsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("zj", APIVersion, "SendSms")
    
    
    return
}

func NewSendSmsResponse() (response *SendSmsResponse) {
    response = &SendSmsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendSms
// 发送短信
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_ROLEFAILURE = "UnauthorizedOperation.RoleFailure"
func (c *Client) SendSms(request *SendSmsRequest) (response *SendSmsResponse, err error) {
    if request == nil {
        request = NewSendSmsRequest()
    }
    
    response = NewSendSmsResponse()
    err = c.Send(request, response)
    return
}
