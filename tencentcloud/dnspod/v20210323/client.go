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

package v20210323

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-03-23"

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


func NewCreateDomainRequest() (request *CreateDomainRequest) {
    request = &CreateDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "CreateDomain")
    return
}

func NewCreateDomainResponse() (response *CreateDomainResponse) {
    response = &CreateDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加域名
func (c *Client) CreateDomain(request *CreateDomainRequest) (response *CreateDomainResponse, err error) {
    if request == nil {
        request = NewCreateDomainRequest()
    }
    response = NewCreateDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDomainAliasRequest() (request *CreateDomainAliasRequest) {
    request = &CreateDomainAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "CreateDomainAlias")
    return
}

func NewCreateDomainAliasResponse() (response *CreateDomainAliasResponse) {
    response = &CreateDomainAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建域名别名
func (c *Client) CreateDomainAlias(request *CreateDomainAliasRequest) (response *CreateDomainAliasResponse, err error) {
    if request == nil {
        request = NewCreateDomainAliasRequest()
    }
    response = NewCreateDomainAliasResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDomainBatchRequest() (request *CreateDomainBatchRequest) {
    request = &CreateDomainBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "CreateDomainBatch")
    return
}

func NewCreateDomainBatchResponse() (response *CreateDomainBatchResponse) {
    response = &CreateDomainBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量添加域名
func (c *Client) CreateDomainBatch(request *CreateDomainBatchRequest) (response *CreateDomainBatchResponse, err error) {
    if request == nil {
        request = NewCreateDomainBatchRequest()
    }
    response = NewCreateDomainBatchResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDomainGroupRequest() (request *CreateDomainGroupRequest) {
    request = &CreateDomainGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "CreateDomainGroup")
    return
}

func NewCreateDomainGroupResponse() (response *CreateDomainGroupResponse) {
    response = &CreateDomainGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建域名分组
func (c *Client) CreateDomainGroup(request *CreateDomainGroupRequest) (response *CreateDomainGroupResponse, err error) {
    if request == nil {
        request = NewCreateDomainGroupRequest()
    }
    response = NewCreateDomainGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRecordRequest() (request *CreateRecordRequest) {
    request = &CreateRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "CreateRecord")
    return
}

func NewCreateRecordResponse() (response *CreateRecordResponse) {
    response = &CreateRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加记录
func (c *Client) CreateRecord(request *CreateRecordRequest) (response *CreateRecordResponse, err error) {
    if request == nil {
        request = NewCreateRecordRequest()
    }
    response = NewCreateRecordResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRecordBatchRequest() (request *CreateRecordBatchRequest) {
    request = &CreateRecordBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "CreateRecordBatch")
    return
}

func NewCreateRecordBatchResponse() (response *CreateRecordBatchResponse) {
    response = &CreateRecordBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量添加记录
func (c *Client) CreateRecordBatch(request *CreateRecordBatchRequest) (response *CreateRecordBatchResponse, err error) {
    if request == nil {
        request = NewCreateRecordBatchRequest()
    }
    response = NewCreateRecordBatchResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDomainRequest() (request *DeleteDomainRequest) {
    request = &DeleteDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DeleteDomain")
    return
}

func NewDeleteDomainResponse() (response *DeleteDomainResponse) {
    response = &DeleteDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除域名
func (c *Client) DeleteDomain(request *DeleteDomainRequest) (response *DeleteDomainResponse, err error) {
    if request == nil {
        request = NewDeleteDomainRequest()
    }
    response = NewDeleteDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteDomainAliasRequest() (request *DeleteDomainAliasRequest) {
    request = &DeleteDomainAliasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DeleteDomainAlias")
    return
}

func NewDeleteDomainAliasResponse() (response *DeleteDomainAliasResponse) {
    response = &DeleteDomainAliasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除域名别名
func (c *Client) DeleteDomainAlias(request *DeleteDomainAliasRequest) (response *DeleteDomainAliasResponse, err error) {
    if request == nil {
        request = NewDeleteDomainAliasRequest()
    }
    response = NewDeleteDomainAliasResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteRecordRequest() (request *DeleteRecordRequest) {
    request = &DeleteRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DeleteRecord")
    return
}

func NewDeleteRecordResponse() (response *DeleteRecordResponse) {
    response = &DeleteRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除记录
func (c *Client) DeleteRecord(request *DeleteRecordRequest) (response *DeleteRecordResponse, err error) {
    if request == nil {
        request = NewDeleteRecordRequest()
    }
    response = NewDeleteRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteShareDomainRequest() (request *DeleteShareDomainRequest) {
    request = &DeleteShareDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DeleteShareDomain")
    return
}

func NewDeleteShareDomainResponse() (response *DeleteShareDomainResponse) {
    response = &DeleteShareDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除域名共享
func (c *Client) DeleteShareDomain(request *DeleteShareDomainRequest) (response *DeleteShareDomainResponse, err error) {
    if request == nil {
        request = NewDeleteShareDomainRequest()
    }
    response = NewDeleteShareDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBatchTaskRequest() (request *DescribeBatchTaskRequest) {
    request = &DescribeBatchTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeBatchTask")
    return
}

func NewDescribeBatchTaskResponse() (response *DescribeBatchTaskResponse) {
    response = &DescribeBatchTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取任务详情
func (c *Client) DescribeBatchTask(request *DescribeBatchTaskRequest) (response *DescribeBatchTaskResponse, err error) {
    if request == nil {
        request = NewDescribeBatchTaskRequest()
    }
    response = NewDescribeBatchTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainRequest() (request *DescribeDomainRequest) {
    request = &DescribeDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeDomain")
    return
}

func NewDescribeDomainResponse() (response *DescribeDomainResponse) {
    response = &DescribeDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取域名信息
func (c *Client) DescribeDomain(request *DescribeDomainRequest) (response *DescribeDomainResponse, err error) {
    if request == nil {
        request = NewDescribeDomainRequest()
    }
    response = NewDescribeDomainResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainListRequest() (request *DescribeDomainListRequest) {
    request = &DescribeDomainListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeDomainList")
    return
}

func NewDescribeDomainListResponse() (response *DescribeDomainListResponse) {
    response = &DescribeDomainListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取域名列表
func (c *Client) DescribeDomainList(request *DescribeDomainListRequest) (response *DescribeDomainListResponse, err error) {
    if request == nil {
        request = NewDescribeDomainListRequest()
    }
    response = NewDescribeDomainListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainLogListRequest() (request *DescribeDomainLogListRequest) {
    request = &DescribeDomainLogListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeDomainLogList")
    return
}

func NewDescribeDomainLogListResponse() (response *DescribeDomainLogListResponse) {
    response = &DescribeDomainLogListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取域名日志
func (c *Client) DescribeDomainLogList(request *DescribeDomainLogListRequest) (response *DescribeDomainLogListResponse, err error) {
    if request == nil {
        request = NewDescribeDomainLogListRequest()
    }
    response = NewDescribeDomainLogListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainPurviewRequest() (request *DescribeDomainPurviewRequest) {
    request = &DescribeDomainPurviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeDomainPurview")
    return
}

func NewDescribeDomainPurviewResponse() (response *DescribeDomainPurviewResponse) {
    response = &DescribeDomainPurviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取域名权限
func (c *Client) DescribeDomainPurview(request *DescribeDomainPurviewRequest) (response *DescribeDomainPurviewResponse, err error) {
    if request == nil {
        request = NewDescribeDomainPurviewRequest()
    }
    response = NewDescribeDomainPurviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainShareInfoRequest() (request *DescribeDomainShareInfoRequest) {
    request = &DescribeDomainShareInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeDomainShareInfo")
    return
}

func NewDescribeDomainShareInfoResponse() (response *DescribeDomainShareInfoResponse) {
    response = &DescribeDomainShareInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取域名共享信息
func (c *Client) DescribeDomainShareInfo(request *DescribeDomainShareInfoRequest) (response *DescribeDomainShareInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDomainShareInfoRequest()
    }
    response = NewDescribeDomainShareInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordRequest() (request *DescribeRecordRequest) {
    request = &DescribeRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeRecord")
    return
}

func NewDescribeRecordResponse() (response *DescribeRecordResponse) {
    response = &DescribeRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取记录信息
func (c *Client) DescribeRecord(request *DescribeRecordRequest) (response *DescribeRecordResponse, err error) {
    if request == nil {
        request = NewDescribeRecordRequest()
    }
    response = NewDescribeRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordLineListRequest() (request *DescribeRecordLineListRequest) {
    request = &DescribeRecordLineListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeRecordLineList")
    return
}

func NewDescribeRecordLineListResponse() (response *DescribeRecordLineListResponse) {
    response = &DescribeRecordLineListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取等级允许的线路
func (c *Client) DescribeRecordLineList(request *DescribeRecordLineListRequest) (response *DescribeRecordLineListResponse, err error) {
    if request == nil {
        request = NewDescribeRecordLineListRequest()
    }
    response = NewDescribeRecordLineListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordListRequest() (request *DescribeRecordListRequest) {
    request = &DescribeRecordListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeRecordList")
    return
}

func NewDescribeRecordListResponse() (response *DescribeRecordListResponse) {
    response = &DescribeRecordListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取某个域名下的解析记录
func (c *Client) DescribeRecordList(request *DescribeRecordListRequest) (response *DescribeRecordListResponse, err error) {
    if request == nil {
        request = NewDescribeRecordListRequest()
    }
    response = NewDescribeRecordListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordTypeRequest() (request *DescribeRecordTypeRequest) {
    request = &DescribeRecordTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeRecordType")
    return
}

func NewDescribeRecordTypeResponse() (response *DescribeRecordTypeResponse) {
    response = &DescribeRecordTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取等级允许的记录类型
func (c *Client) DescribeRecordType(request *DescribeRecordTypeRequest) (response *DescribeRecordTypeResponse, err error) {
    if request == nil {
        request = NewDescribeRecordTypeRequest()
    }
    response = NewDescribeRecordTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserDetailRequest() (request *DescribeUserDetailRequest) {
    request = &DescribeUserDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "DescribeUserDetail")
    return
}

func NewDescribeUserDetailResponse() (response *DescribeUserDetailResponse) {
    response = &DescribeUserDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取帐户信息
func (c *Client) DescribeUserDetail(request *DescribeUserDetailRequest) (response *DescribeUserDetailResponse, err error) {
    if request == nil {
        request = NewDescribeUserDetailRequest()
    }
    response = NewDescribeUserDetailResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainLockRequest() (request *ModifyDomainLockRequest) {
    request = &ModifyDomainLockRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyDomainLock")
    return
}

func NewModifyDomainLockResponse() (response *ModifyDomainLockResponse) {
    response = &ModifyDomainLockResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 锁定域名
func (c *Client) ModifyDomainLock(request *ModifyDomainLockRequest) (response *ModifyDomainLockResponse, err error) {
    if request == nil {
        request = NewModifyDomainLockRequest()
    }
    response = NewModifyDomainLockResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainOwnerRequest() (request *ModifyDomainOwnerRequest) {
    request = &ModifyDomainOwnerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyDomainOwner")
    return
}

func NewModifyDomainOwnerResponse() (response *ModifyDomainOwnerResponse) {
    response = &ModifyDomainOwnerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 域名过户
func (c *Client) ModifyDomainOwner(request *ModifyDomainOwnerRequest) (response *ModifyDomainOwnerResponse, err error) {
    if request == nil {
        request = NewModifyDomainOwnerRequest()
    }
    response = NewModifyDomainOwnerResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainRemarkRequest() (request *ModifyDomainRemarkRequest) {
    request = &ModifyDomainRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyDomainRemark")
    return
}

func NewModifyDomainRemarkResponse() (response *ModifyDomainRemarkResponse) {
    response = &ModifyDomainRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置域名备注
func (c *Client) ModifyDomainRemark(request *ModifyDomainRemarkRequest) (response *ModifyDomainRemarkResponse, err error) {
    if request == nil {
        request = NewModifyDomainRemarkRequest()
    }
    response = NewModifyDomainRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainStatusRequest() (request *ModifyDomainStatusRequest) {
    request = &ModifyDomainStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyDomainStatus")
    return
}

func NewModifyDomainStatusResponse() (response *ModifyDomainStatusResponse) {
    response = &ModifyDomainStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改域名状态
func (c *Client) ModifyDomainStatus(request *ModifyDomainStatusRequest) (response *ModifyDomainStatusResponse, err error) {
    if request == nil {
        request = NewModifyDomainStatusRequest()
    }
    response = NewModifyDomainStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDomainUnlockRequest() (request *ModifyDomainUnlockRequest) {
    request = &ModifyDomainUnlockRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyDomainUnlock")
    return
}

func NewModifyDomainUnlockResponse() (response *ModifyDomainUnlockResponse) {
    response = &ModifyDomainUnlockResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 域名锁定解锁
func (c *Client) ModifyDomainUnlock(request *ModifyDomainUnlockRequest) (response *ModifyDomainUnlockResponse, err error) {
    if request == nil {
        request = NewModifyDomainUnlockRequest()
    }
    response = NewModifyDomainUnlockResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDynamicDNSRequest() (request *ModifyDynamicDNSRequest) {
    request = &ModifyDynamicDNSRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyDynamicDNS")
    return
}

func NewModifyDynamicDNSResponse() (response *ModifyDynamicDNSResponse) {
    response = &ModifyDynamicDNSResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新动态 DNS 记录
func (c *Client) ModifyDynamicDNS(request *ModifyDynamicDNSRequest) (response *ModifyDynamicDNSResponse, err error) {
    if request == nil {
        request = NewModifyDynamicDNSRequest()
    }
    response = NewModifyDynamicDNSResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRecordRequest() (request *ModifyRecordRequest) {
    request = &ModifyRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyRecord")
    return
}

func NewModifyRecordResponse() (response *ModifyRecordResponse) {
    response = &ModifyRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改记录
func (c *Client) ModifyRecord(request *ModifyRecordRequest) (response *ModifyRecordResponse, err error) {
    if request == nil {
        request = NewModifyRecordRequest()
    }
    response = NewModifyRecordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRecordBatchRequest() (request *ModifyRecordBatchRequest) {
    request = &ModifyRecordBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyRecordBatch")
    return
}

func NewModifyRecordBatchResponse() (response *ModifyRecordBatchResponse) {
    response = &ModifyRecordBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量修改记录
func (c *Client) ModifyRecordBatch(request *ModifyRecordBatchRequest) (response *ModifyRecordBatchResponse, err error) {
    if request == nil {
        request = NewModifyRecordBatchRequest()
    }
    response = NewModifyRecordBatchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRecordRemarkRequest() (request *ModifyRecordRemarkRequest) {
    request = &ModifyRecordRemarkRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyRecordRemark")
    return
}

func NewModifyRecordRemarkResponse() (response *ModifyRecordRemarkResponse) {
    response = &ModifyRecordRemarkResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 设置记录备注
func (c *Client) ModifyRecordRemark(request *ModifyRecordRemarkRequest) (response *ModifyRecordRemarkResponse, err error) {
    if request == nil {
        request = NewModifyRecordRemarkRequest()
    }
    response = NewModifyRecordRemarkResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRecordStatusRequest() (request *ModifyRecordStatusRequest) {
    request = &ModifyRecordStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("dnspod", APIVersion, "ModifyRecordStatus")
    return
}

func NewModifyRecordStatusResponse() (response *ModifyRecordStatusResponse) {
    response = &ModifyRecordStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改解析记录的状态
func (c *Client) ModifyRecordStatus(request *ModifyRecordStatusRequest) (response *ModifyRecordStatusResponse, err error) {
    if request == nil {
        request = NewModifyRecordStatusRequest()
    }
    response = NewModifyRecordStatusResponse()
    err = c.Send(request, response)
    return
}
