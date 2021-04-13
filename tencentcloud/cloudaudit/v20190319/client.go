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

package v20190319

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-03-19"

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


func NewCreateAuditRequest() (request *CreateAuditRequest) {
    request = &CreateAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "CreateAudit")
    return
}

func NewCreateAuditResponse() (response *CreateAuditResponse) {
    response = &CreateAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 参数要求：
// 1、如果IsCreateNewBucket的值存在的话，cosRegion和cosBucketName都是必填参数。
// 2、如果IsEnableCmqNotify的值是1的话，IsCreateNewQueue、CmqRegion和CmqQueueName都是必填参数。
// 3、如果IsEnableCmqNotify的值是0的话，IsCreateNewQueue、CmqRegion和CmqQueueName都不能传。
// 4、如果IsEnableKmsEncry的值是1的话，KmsRegion和KeyId属于必填项
func (c *Client) CreateAudit(request *CreateAuditRequest) (response *CreateAuditResponse, err error) {
    if request == nil {
        request = NewCreateAuditRequest()
    }
    response = NewCreateAuditResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAuditRequest() (request *DeleteAuditRequest) {
    request = &DeleteAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DeleteAudit")
    return
}

func NewDeleteAuditResponse() (response *DeleteAuditResponse) {
    response = &DeleteAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除跟踪集
func (c *Client) DeleteAudit(request *DeleteAuditRequest) (response *DeleteAuditResponse, err error) {
    if request == nil {
        request = NewDeleteAuditRequest()
    }
    response = NewDeleteAuditResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditRequest() (request *DescribeAuditRequest) {
    request = &DescribeAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DescribeAudit")
    return
}

func NewDescribeAuditResponse() (response *DescribeAuditResponse) {
    response = &DescribeAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询跟踪集详情
func (c *Client) DescribeAudit(request *DescribeAuditRequest) (response *DescribeAuditResponse, err error) {
    if request == nil {
        request = NewDescribeAuditRequest()
    }
    response = NewDescribeAuditResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventsRequest() (request *DescribeEventsRequest) {
    request = &DescribeEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "DescribeEvents")
    return
}

func NewDescribeEventsResponse() (response *DescribeEventsResponse) {
    response = &DescribeEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询云审计日志
func (c *Client) DescribeEvents(request *DescribeEventsRequest) (response *DescribeEventsResponse, err error) {
    if request == nil {
        request = NewDescribeEventsRequest()
    }
    response = NewDescribeEventsResponse()
    err = c.Send(request, response)
    return
}

func NewGetAttributeKeyRequest() (request *GetAttributeKeyRequest) {
    request = &GetAttributeKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "GetAttributeKey")
    return
}

func NewGetAttributeKeyResponse() (response *GetAttributeKeyResponse) {
    response = &GetAttributeKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询AttributeKey的有效取值范围
func (c *Client) GetAttributeKey(request *GetAttributeKeyRequest) (response *GetAttributeKeyResponse, err error) {
    if request == nil {
        request = NewGetAttributeKeyRequest()
    }
    response = NewGetAttributeKeyResponse()
    err = c.Send(request, response)
    return
}

func NewInquireAuditCreditRequest() (request *InquireAuditCreditRequest) {
    request = &InquireAuditCreditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "InquireAuditCredit")
    return
}

func NewInquireAuditCreditResponse() (response *InquireAuditCreditResponse) {
    response = &InquireAuditCreditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询用户可创建跟踪集的数量
func (c *Client) InquireAuditCredit(request *InquireAuditCreditRequest) (response *InquireAuditCreditResponse, err error) {
    if request == nil {
        request = NewInquireAuditCreditRequest()
    }
    response = NewInquireAuditCreditResponse()
    err = c.Send(request, response)
    return
}

func NewListAuditsRequest() (request *ListAuditsRequest) {
    request = &ListAuditsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "ListAudits")
    return
}

func NewListAuditsResponse() (response *ListAuditsResponse) {
    response = &ListAuditsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询跟踪集概要
func (c *Client) ListAudits(request *ListAuditsRequest) (response *ListAuditsResponse, err error) {
    if request == nil {
        request = NewListAuditsRequest()
    }
    response = NewListAuditsResponse()
    err = c.Send(request, response)
    return
}

func NewListCmqEnableRegionRequest() (request *ListCmqEnableRegionRequest) {
    request = &ListCmqEnableRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "ListCmqEnableRegion")
    return
}

func NewListCmqEnableRegionResponse() (response *ListCmqEnableRegionResponse) {
    response = &ListCmqEnableRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询云审计支持的cmq的可用区
func (c *Client) ListCmqEnableRegion(request *ListCmqEnableRegionRequest) (response *ListCmqEnableRegionResponse, err error) {
    if request == nil {
        request = NewListCmqEnableRegionRequest()
    }
    response = NewListCmqEnableRegionResponse()
    err = c.Send(request, response)
    return
}

func NewListCosEnableRegionRequest() (request *ListCosEnableRegionRequest) {
    request = &ListCosEnableRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "ListCosEnableRegion")
    return
}

func NewListCosEnableRegionResponse() (response *ListCosEnableRegionResponse) {
    response = &ListCosEnableRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询云审计支持的cos可用区
func (c *Client) ListCosEnableRegion(request *ListCosEnableRegionRequest) (response *ListCosEnableRegionResponse, err error) {
    if request == nil {
        request = NewListCosEnableRegionRequest()
    }
    response = NewListCosEnableRegionResponse()
    err = c.Send(request, response)
    return
}

func NewListKeyAliasByRegionRequest() (request *ListKeyAliasByRegionRequest) {
    request = &ListKeyAliasByRegionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "ListKeyAliasByRegion")
    return
}

func NewListKeyAliasByRegionResponse() (response *ListKeyAliasByRegionResponse) {
    response = &ListKeyAliasByRegionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据地域获取KMS密钥别名
func (c *Client) ListKeyAliasByRegion(request *ListKeyAliasByRegionRequest) (response *ListKeyAliasByRegionResponse, err error) {
    if request == nil {
        request = NewListKeyAliasByRegionRequest()
    }
    response = NewListKeyAliasByRegionResponse()
    err = c.Send(request, response)
    return
}

func NewLookUpEventsRequest() (request *LookUpEventsRequest) {
    request = &LookUpEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "LookUpEvents")
    return
}

func NewLookUpEventsResponse() (response *LookUpEventsResponse) {
    response = &LookUpEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于对操作日志进行检索，便于用户进行查询相关的操作信息。
func (c *Client) LookUpEvents(request *LookUpEventsRequest) (response *LookUpEventsResponse, err error) {
    if request == nil {
        request = NewLookUpEventsRequest()
    }
    response = NewLookUpEventsResponse()
    err = c.Send(request, response)
    return
}

func NewStartLoggingRequest() (request *StartLoggingRequest) {
    request = &StartLoggingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "StartLogging")
    return
}

func NewStartLoggingResponse() (response *StartLoggingResponse) {
    response = &StartLoggingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 开启跟踪集
func (c *Client) StartLogging(request *StartLoggingRequest) (response *StartLoggingResponse, err error) {
    if request == nil {
        request = NewStartLoggingRequest()
    }
    response = NewStartLoggingResponse()
    err = c.Send(request, response)
    return
}

func NewStopLoggingRequest() (request *StopLoggingRequest) {
    request = &StopLoggingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "StopLogging")
    return
}

func NewStopLoggingResponse() (response *StopLoggingResponse) {
    response = &StopLoggingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 关闭跟踪集
func (c *Client) StopLogging(request *StopLoggingRequest) (response *StopLoggingResponse, err error) {
    if request == nil {
        request = NewStopLoggingRequest()
    }
    response = NewStopLoggingResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAuditRequest() (request *UpdateAuditRequest) {
    request = &UpdateAuditRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cloudaudit", APIVersion, "UpdateAudit")
    return
}

func NewUpdateAuditResponse() (response *UpdateAuditResponse) {
    response = &UpdateAuditResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 参数要求：
// 1、如果IsCreateNewBucket的值存在的话，cosRegion和cosBucketName都是必填参数。
// 2、如果IsEnableCmqNotify的值是1的话，IsCreateNewQueue、CmqRegion和CmqQueueName都是必填参数。
// 3、如果IsEnableCmqNotify的值是0的话，IsCreateNewQueue、CmqRegion和CmqQueueName都不能传。
// 4、如果IsEnableKmsEncry的值是1的话，KmsRegion和KeyId属于必填项
func (c *Client) UpdateAudit(request *UpdateAuditRequest) (response *UpdateAuditResponse, err error) {
    if request == nil {
        request = NewUpdateAuditRequest()
    }
    response = NewUpdateAuditResponse()
    err = c.Send(request, response)
    return
}
