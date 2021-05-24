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

package v20180808

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-08-08"

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


func NewBatchModifyDomainInfoRequest() (request *BatchModifyDomainInfoRequest) {
    request = &BatchModifyDomainInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "BatchModifyDomainInfo")
    return
}

func NewBatchModifyDomainInfoResponse() (response *BatchModifyDomainInfoResponse) {
    response = &BatchModifyDomainInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 ( BatchModifyDomainInfo ) 用于批量域名信息修改 。
func (c *Client) BatchModifyDomainInfo(request *BatchModifyDomainInfoRequest) (response *BatchModifyDomainInfoResponse, err error) {
    if request == nil {
        request = NewBatchModifyDomainInfoRequest()
    }
    response = NewBatchModifyDomainInfoResponse()
    err = c.Send(request, response)
    return
}

func NewCheckBatchStatusRequest() (request *CheckBatchStatusRequest) {
    request = &CheckBatchStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "CheckBatchStatus")
    return
}

func NewCheckBatchStatusResponse() (response *CheckBatchStatusResponse) {
    response = &CheckBatchStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 ( CheckBatchStatus ) 用于查询批量操作日志状态 。
func (c *Client) CheckBatchStatus(request *CheckBatchStatusRequest) (response *CheckBatchStatusResponse, err error) {
    if request == nil {
        request = NewCheckBatchStatusRequest()
    }
    response = NewCheckBatchStatusResponse()
    err = c.Send(request, response)
    return
}

func NewCheckDomainRequest() (request *CheckDomainRequest) {
    request = &CheckDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "CheckDomain")
    return
}

func NewCheckDomainResponse() (response *CheckDomainResponse) {
    response = &CheckDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 检查域名是否可以注册。
func (c *Client) CheckDomain(request *CheckDomainRequest) (response *CheckDomainResponse, err error) {
    if request == nil {
        request = NewCheckDomainRequest()
    }
    response = NewCheckDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDomainBatchRequest() (request *CreateDomainBatchRequest) {
    request = &CreateDomainBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "CreateDomainBatch")
    return
}

func NewCreateDomainBatchResponse() (response *CreateDomainBatchResponse) {
    response = &CreateDomainBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 ( CreateDomainBatch ) 用于批量域名注册 。
func (c *Client) CreateDomainBatch(request *CreateDomainBatchRequest) (response *CreateDomainBatchResponse, err error) {
    if request == nil {
        request = NewCreateDomainBatchRequest()
    }
    response = NewCreateDomainBatchResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTemplateRequest() (request *CreateTemplateRequest) {
    request = &CreateTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "CreateTemplate")
    return
}

func NewCreateTemplateResponse() (response *CreateTemplateResponse) {
    response = &CreateTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 ( CreateTemplate ) 用于添加域名信息模板 。
func (c *Client) CreateTemplate(request *CreateTemplateRequest) (response *CreateTemplateResponse, err error) {
    if request == nil {
        request = NewCreateTemplateRequest()
    }
    response = NewCreateTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTemplateRequest() (request *DeleteTemplateRequest) {
    request = &DeleteTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "DeleteTemplate")
    return
}

func NewDeleteTemplateResponse() (response *DeleteTemplateResponse) {
    response = &DeleteTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 ( DeleteTemplate ) 用于删除信息模板。
func (c *Client) DeleteTemplate(request *DeleteTemplateRequest) (response *DeleteTemplateResponse, err error) {
    if request == nil {
        request = NewDeleteTemplateRequest()
    }
    response = NewDeleteTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBatchOperationLogDetailsRequest() (request *DescribeBatchOperationLogDetailsRequest) {
    request = &DescribeBatchOperationLogDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "DescribeBatchOperationLogDetails")
    return
}

func NewDescribeBatchOperationLogDetailsResponse() (response *DescribeBatchOperationLogDetailsResponse) {
    response = &DescribeBatchOperationLogDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 ( DescribeBatchOperationLogDetails ) 用于获取批量操作日志详情。
func (c *Client) DescribeBatchOperationLogDetails(request *DescribeBatchOperationLogDetailsRequest) (response *DescribeBatchOperationLogDetailsResponse, err error) {
    if request == nil {
        request = NewDescribeBatchOperationLogDetailsRequest()
    }
    response = NewDescribeBatchOperationLogDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBatchOperationLogsRequest() (request *DescribeBatchOperationLogsRequest) {
    request = &DescribeBatchOperationLogsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "DescribeBatchOperationLogs")
    return
}

func NewDescribeBatchOperationLogsResponse() (response *DescribeBatchOperationLogsResponse) {
    response = &DescribeBatchOperationLogsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 ( DescribeBatchOperationLogs ) 用于获取批量操作日志 。
func (c *Client) DescribeBatchOperationLogs(request *DescribeBatchOperationLogsRequest) (response *DescribeBatchOperationLogsResponse, err error) {
    if request == nil {
        request = NewDescribeBatchOperationLogsRequest()
    }
    response = NewDescribeBatchOperationLogsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainBaseInfoRequest() (request *DescribeDomainBaseInfoRequest) {
    request = &DescribeDomainBaseInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "DescribeDomainBaseInfo")
    return
}

func NewDescribeDomainBaseInfoResponse() (response *DescribeDomainBaseInfoResponse) {
    response = &DescribeDomainBaseInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (  DescribeDomainBaseInfo) 获取域名基本信息。
func (c *Client) DescribeDomainBaseInfo(request *DescribeDomainBaseInfoRequest) (response *DescribeDomainBaseInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDomainBaseInfoRequest()
    }
    response = NewDescribeDomainBaseInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainNameListRequest() (request *DescribeDomainNameListRequest) {
    request = &DescribeDomainNameListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "DescribeDomainNameList")
    return
}

func NewDescribeDomainNameListResponse() (response *DescribeDomainNameListResponse) {
    response = &DescribeDomainNameListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (  DescribeDomainNameList ) 我的域名列表。
func (c *Client) DescribeDomainNameList(request *DescribeDomainNameListRequest) (response *DescribeDomainNameListResponse, err error) {
    if request == nil {
        request = NewDescribeDomainNameListRequest()
    }
    response = NewDescribeDomainNameListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainPriceListRequest() (request *DescribeDomainPriceListRequest) {
    request = &DescribeDomainPriceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "DescribeDomainPriceList")
    return
}

func NewDescribeDomainPriceListResponse() (response *DescribeDomainPriceListResponse) {
    response = &DescribeDomainPriceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 按照域名后缀获取对应的价格列表
func (c *Client) DescribeDomainPriceList(request *DescribeDomainPriceListRequest) (response *DescribeDomainPriceListResponse, err error) {
    if request == nil {
        request = NewDescribeDomainPriceListRequest()
    }
    response = NewDescribeDomainPriceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTemplateRequest() (request *DescribeTemplateRequest) {
    request = &DescribeTemplateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "DescribeTemplate")
    return
}

func NewDescribeTemplateResponse() (response *DescribeTemplateResponse) {
    response = &DescribeTemplateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeTemplate) 用于获取模板信息。
func (c *Client) DescribeTemplate(request *DescribeTemplateRequest) (response *DescribeTemplateResponse, err error) {
    if request == nil {
        request = NewDescribeTemplateRequest()
    }
    response = NewDescribeTemplateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTemplateListRequest() (request *DescribeTemplateListRequest) {
    request = &DescribeTemplateListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "DescribeTemplateList")
    return
}

func NewDescribeTemplateListResponse() (response *DescribeTemplateListResponse) {
    response = &DescribeTemplateListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeTemplateList) 用于获取信息模板列表。
func (c *Client) DescribeTemplateList(request *DescribeTemplateListRequest) (response *DescribeTemplateListResponse, err error) {
    if request == nil {
        request = NewDescribeTemplateListRequest()
    }
    response = NewDescribeTemplateListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainDNSBatchRequest() (request *ModifyDomainDNSBatchRequest) {
    request = &ModifyDomainDNSBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "ModifyDomainDNSBatch")
    return
}

func NewModifyDomainDNSBatchResponse() (response *ModifyDomainDNSBatchResponse) {
    response = &ModifyDomainDNSBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 ( ModifyDomainDNSBatch) 用于批量域名 DNS 修改 。
func (c *Client) ModifyDomainDNSBatch(request *ModifyDomainDNSBatchRequest) (response *ModifyDomainDNSBatchResponse, err error) {
    if request == nil {
        request = NewModifyDomainDNSBatchRequest()
    }
    response = NewModifyDomainDNSBatchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainOwnerBatchRequest() (request *ModifyDomainOwnerBatchRequest) {
    request = &ModifyDomainOwnerBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "ModifyDomainOwnerBatch")
    return
}

func NewModifyDomainOwnerBatchResponse() (response *ModifyDomainOwnerBatchResponse) {
    response = &ModifyDomainOwnerBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 ( ModifyDomainOwnerBatch) 用于域名批量账号间转移 。
func (c *Client) ModifyDomainOwnerBatch(request *ModifyDomainOwnerBatchRequest) (response *ModifyDomainOwnerBatchResponse, err error) {
    if request == nil {
        request = NewModifyDomainOwnerBatchRequest()
    }
    response = NewModifyDomainOwnerBatchResponse()
    err = c.Send(request, response)
    return
}

func NewRenewDomainBatchRequest() (request *RenewDomainBatchRequest) {
    request = &RenewDomainBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "RenewDomainBatch")
    return
}

func NewRenewDomainBatchResponse() (response *RenewDomainBatchResponse) {
    response = &RenewDomainBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 ( RenewDomainBatch ) 用于批量续费域名 。
func (c *Client) RenewDomainBatch(request *RenewDomainBatchRequest) (response *RenewDomainBatchResponse, err error) {
    if request == nil {
        request = NewRenewDomainBatchRequest()
    }
    response = NewRenewDomainBatchResponse()
    err = c.Send(request, response)
    return
}

func NewSetDomainAutoRenewRequest() (request *SetDomainAutoRenewRequest) {
    request = &SetDomainAutoRenewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "SetDomainAutoRenew")
    return
}

func NewSetDomainAutoRenewResponse() (response *SetDomainAutoRenewResponse) {
    response = &SetDomainAutoRenewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 ( SetDomainAutoRenew ) 用于设置域名自动续费。
func (c *Client) SetDomainAutoRenew(request *SetDomainAutoRenewRequest) (response *SetDomainAutoRenewResponse, err error) {
    if request == nil {
        request = NewSetDomainAutoRenewRequest()
    }
    response = NewSetDomainAutoRenewResponse()
    err = c.Send(request, response)
    return
}

func NewTransferInDomainBatchRequest() (request *TransferInDomainBatchRequest) {
    request = &TransferInDomainBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "TransferInDomainBatch")
    return
}

func NewTransferInDomainBatchResponse() (response *TransferInDomainBatchResponse) {
    response = &TransferInDomainBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 ( TransferInDomainBatch ) 用于批量转入域名 。
func (c *Client) TransferInDomainBatch(request *TransferInDomainBatchRequest) (response *TransferInDomainBatchResponse, err error) {
    if request == nil {
        request = NewTransferInDomainBatchRequest()
    }
    response = NewTransferInDomainBatchResponse()
    err = c.Send(request, response)
    return
}

func NewTransferProhibitionBatchRequest() (request *TransferProhibitionBatchRequest) {
    request = &TransferProhibitionBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "TransferProhibitionBatch")
    return
}

func NewTransferProhibitionBatchResponse() (response *TransferProhibitionBatchResponse) {
    response = &TransferProhibitionBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 ( TransferProhibitionBatch ) 用于批量禁止域名转移 。
func (c *Client) TransferProhibitionBatch(request *TransferProhibitionBatchRequest) (response *TransferProhibitionBatchResponse, err error) {
    if request == nil {
        request = NewTransferProhibitionBatchRequest()
    }
    response = NewTransferProhibitionBatchResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateProhibitionBatchRequest() (request *UpdateProhibitionBatchRequest) {
    request = &UpdateProhibitionBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "UpdateProhibitionBatch")
    return
}

func NewUpdateProhibitionBatchResponse() (response *UpdateProhibitionBatchResponse) {
    response = &UpdateProhibitionBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 ( UpdateProhibitionBatch ) 用于批量禁止更新锁。
func (c *Client) UpdateProhibitionBatch(request *UpdateProhibitionBatchRequest) (response *UpdateProhibitionBatchResponse, err error) {
    if request == nil {
        request = NewUpdateProhibitionBatchRequest()
    }
    response = NewUpdateProhibitionBatchResponse()
    err = c.Send(request, response)
    return
}

func NewUploadImageRequest() (request *UploadImageRequest) {
    request = &UploadImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "UploadImage")
    return
}

func NewUploadImageResponse() (response *UploadImageResponse) {
    response = &UploadImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 ( UploadImage ) 用于证件图片上传 。
func (c *Client) UploadImage(request *UploadImageRequest) (response *UploadImageResponse, err error) {
    if request == nil {
        request = NewUploadImageRequest()
    }
    response = NewUploadImageResponse()
    err = c.Send(request, response)
    return
}
