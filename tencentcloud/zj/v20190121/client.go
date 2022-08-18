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
    "context"
    "errors"
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
    return c.AddCrowdPackInfoWithContext(context.Background(), request)
}

// AddCrowdPackInfo
// 添加短信人群包信息
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) AddCrowdPackInfoWithContext(ctx context.Context, request *AddCrowdPackInfoRequest) (response *AddCrowdPackInfoResponse, err error) {
    if request == nil {
        request = NewAddCrowdPackInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddCrowdPackInfo require credential")
    }

    request.SetContext(ctx)
    
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
    return c.AddSmsSignWithContext(context.Background(), request)
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
func (c *Client) AddSmsSignWithContext(ctx context.Context, request *AddSmsSignRequest) (response *AddSmsSignResponse, err error) {
    if request == nil {
        request = NewAddSmsSignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddSmsSign require credential")
    }

    request.SetContext(ctx)
    
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
    return c.AddSmsTemplateWithContext(context.Background(), request)
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
func (c *Client) AddSmsTemplateWithContext(ctx context.Context, request *AddSmsTemplateRequest) (response *AddSmsTemplateResponse, err error) {
    if request == nil {
        request = NewAddSmsTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("AddSmsTemplate require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CancelCampaignWithContext(context.Background(), request)
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
func (c *Client) CancelCampaignWithContext(ctx context.Context, request *CancelCampaignRequest) (response *CancelCampaignResponse, err error) {
    if request == nil {
        request = NewCancelCampaignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelCampaign require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateCampaignWithContext(context.Background(), request)
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
func (c *Client) CreateCampaignWithContext(ctx context.Context, request *CreateCampaignRequest) (response *CreateCampaignResponse, err error) {
    if request == nil {
        request = NewCreateCampaignRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCampaign require credential")
    }

    request.SetContext(ctx)
    
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
    return c.CreateMmsInstanceWithContext(context.Background(), request)
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
func (c *Client) CreateMmsInstanceWithContext(ctx context.Context, request *CreateMmsInstanceRequest) (response *CreateMmsInstanceResponse, err error) {
    if request == nil {
        request = NewCreateMmsInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateMmsInstance require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DelCrowdPackWithContext(context.Background(), request)
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
func (c *Client) DelCrowdPackWithContext(ctx context.Context, request *DelCrowdPackRequest) (response *DelCrowdPackResponse, err error) {
    if request == nil {
        request = NewDelCrowdPackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DelCrowdPack require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DelTemplateWithContext(context.Background(), request)
}

// DelTemplate
// 删除短信模板
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) DelTemplateWithContext(ctx context.Context, request *DelTemplateRequest) (response *DelTemplateResponse, err error) {
    if request == nil {
        request = NewDelTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DelTemplate require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DeleteMmsInstanceWithContext(context.Background(), request)
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
func (c *Client) DeleteMmsInstanceWithContext(ctx context.Context, request *DeleteMmsInstanceRequest) (response *DeleteMmsInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteMmsInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteMmsInstance require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeMmsInstanceInfoWithContext(context.Background(), request)
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
func (c *Client) DescribeMmsInstanceInfoWithContext(ctx context.Context, request *DescribeMmsInstanceInfoRequest) (response *DescribeMmsInstanceInfoResponse, err error) {
    if request == nil {
        request = NewDescribeMmsInstanceInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMmsInstanceInfo require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeMmsInstanceListWithContext(context.Background(), request)
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
func (c *Client) DescribeMmsInstanceListWithContext(ctx context.Context, request *DescribeMmsInstanceListRequest) (response *DescribeMmsInstanceListResponse, err error) {
    if request == nil {
        request = NewDescribeMmsInstanceListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMmsInstanceList require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeSmsCampaignStatisticsWithContext(context.Background(), request)
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
func (c *Client) DescribeSmsCampaignStatisticsWithContext(ctx context.Context, request *DescribeSmsCampaignStatisticsRequest) (response *DescribeSmsCampaignStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeSmsCampaignStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSmsCampaignStatistics require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeSmsSignListWithContext(context.Background(), request)
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
func (c *Client) DescribeSmsSignListWithContext(ctx context.Context, request *DescribeSmsSignListRequest) (response *DescribeSmsSignListResponse, err error) {
    if request == nil {
        request = NewDescribeSmsSignListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSmsSignList require credential")
    }

    request.SetContext(ctx)
    
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
    return c.DescribeSmsTemplateListWithContext(context.Background(), request)
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
func (c *Client) DescribeSmsTemplateListWithContext(ctx context.Context, request *DescribeSmsTemplateListRequest) (response *DescribeSmsTemplateListResponse, err error) {
    if request == nil {
        request = NewDescribeSmsTemplateListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSmsTemplateList require credential")
    }

    request.SetContext(ctx)
    
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
    return c.GetCrowdPackListWithContext(context.Background(), request)
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
func (c *Client) GetCrowdPackListWithContext(ctx context.Context, request *GetCrowdPackListRequest) (response *GetCrowdPackListResponse, err error) {
    if request == nil {
        request = NewGetCrowdPackListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCrowdPackList require credential")
    }

    request.SetContext(ctx)
    
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
    return c.GetCrowdUploadInfoWithContext(context.Background(), request)
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
func (c *Client) GetCrowdUploadInfoWithContext(ctx context.Context, request *GetCrowdUploadInfoRequest) (response *GetCrowdUploadInfoResponse, err error) {
    if request == nil {
        request = NewGetCrowdUploadInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCrowdUploadInfo require credential")
    }

    request.SetContext(ctx)
    
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
    return c.GetSmsAmountInfoWithContext(context.Background(), request)
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
func (c *Client) GetSmsAmountInfoWithContext(ctx context.Context, request *GetSmsAmountInfoRequest) (response *GetSmsAmountInfoResponse, err error) {
    if request == nil {
        request = NewGetSmsAmountInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSmsAmountInfo require credential")
    }

    request.SetContext(ctx)
    
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
    return c.GetSmsCampaignStatusWithContext(context.Background(), request)
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
func (c *Client) GetSmsCampaignStatusWithContext(ctx context.Context, request *GetSmsCampaignStatusRequest) (response *GetSmsCampaignStatusResponse, err error) {
    if request == nil {
        request = NewGetSmsCampaignStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetSmsCampaignStatus require credential")
    }

    request.SetContext(ctx)
    
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
    return c.ModifySmsTemplateWithContext(context.Background(), request)
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
func (c *Client) ModifySmsTemplateWithContext(ctx context.Context, request *ModifySmsTemplateRequest) (response *ModifySmsTemplateResponse, err error) {
    if request == nil {
        request = NewModifySmsTemplateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifySmsTemplate require credential")
    }

    request.SetContext(ctx)
    
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
    return c.PushMmsContentWithContext(context.Background(), request)
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
func (c *Client) PushMmsContentWithContext(ctx context.Context, request *PushMmsContentRequest) (response *PushMmsContentResponse, err error) {
    if request == nil {
        request = NewPushMmsContentRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PushMmsContent require credential")
    }

    request.SetContext(ctx)
    
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
    return c.SendSmsWithContext(context.Background(), request)
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
func (c *Client) SendSmsWithContext(ctx context.Context, request *SendSmsRequest) (response *SendSmsResponse, err error) {
    if request == nil {
        request = NewSendSmsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendSms require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendSmsResponse()
    err = c.Send(request, response)
    return
}
