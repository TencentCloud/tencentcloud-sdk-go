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

package v20201028

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-10-28"

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


func NewCreatePrivateZoneRequest() (request *CreatePrivateZoneRequest) {
    request = &CreatePrivateZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "CreatePrivateZone")
    return
}

func NewCreatePrivateZoneResponse() (response *CreatePrivateZoneResponse) {
    response = &CreatePrivateZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建私有域
func (c *Client) CreatePrivateZone(request *CreatePrivateZoneRequest) (response *CreatePrivateZoneResponse, err error) {
    if request == nil {
        request = NewCreatePrivateZoneRequest()
    }
    response = NewCreatePrivateZoneResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePrivateZoneRecordRequest() (request *CreatePrivateZoneRecordRequest) {
    request = &CreatePrivateZoneRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "CreatePrivateZoneRecord")
    return
}

func NewCreatePrivateZoneRecordResponse() (response *CreatePrivateZoneRecordResponse) {
    response = &CreatePrivateZoneRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 添加私有域解析记录
func (c *Client) CreatePrivateZoneRecord(request *CreatePrivateZoneRecordRequest) (response *CreatePrivateZoneRecordResponse, err error) {
    if request == nil {
        request = NewCreatePrivateZoneRecordRequest()
    }
    response = NewCreatePrivateZoneRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrivateZoneRequest() (request *DeletePrivateZoneRequest) {
    request = &DeletePrivateZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "DeletePrivateZone")
    return
}

func NewDeletePrivateZoneResponse() (response *DeletePrivateZoneResponse) {
    response = &DeletePrivateZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除私有域并停止解析
func (c *Client) DeletePrivateZone(request *DeletePrivateZoneRequest) (response *DeletePrivateZoneResponse, err error) {
    if request == nil {
        request = NewDeletePrivateZoneRequest()
    }
    response = NewDeletePrivateZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePrivateZoneRecordRequest() (request *DeletePrivateZoneRecordRequest) {
    request = &DeletePrivateZoneRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "DeletePrivateZoneRecord")
    return
}

func NewDeletePrivateZoneRecordResponse() (response *DeletePrivateZoneRecordResponse) {
    response = &DeletePrivateZoneRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除私有域解析记录
func (c *Client) DeletePrivateZoneRecord(request *DeletePrivateZoneRecordRequest) (response *DeletePrivateZoneRecordResponse, err error) {
    if request == nil {
        request = NewDeletePrivateZoneRecordRequest()
    }
    response = NewDeletePrivateZoneRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuditLogRequest() (request *DescribeAuditLogRequest) {
    request = &DescribeAuditLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribeAuditLog")
    return
}

func NewDescribeAuditLogResponse() (response *DescribeAuditLogResponse) {
    response = &DescribeAuditLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取操作日志列表
func (c *Client) DescribeAuditLog(request *DescribeAuditLogRequest) (response *DescribeAuditLogResponse, err error) {
    if request == nil {
        request = NewDescribeAuditLogRequest()
    }
    response = NewDescribeAuditLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDashboardRequest() (request *DescribeDashboardRequest) {
    request = &DescribeDashboardRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribeDashboard")
    return
}

func NewDescribeDashboardResponse() (response *DescribeDashboardResponse) {
    response = &DescribeDashboardResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取私有域解析概览
func (c *Client) DescribeDashboard(request *DescribeDashboardRequest) (response *DescribeDashboardResponse, err error) {
    if request == nil {
        request = NewDescribeDashboardRequest()
    }
    response = NewDescribeDashboardResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateZoneRequest() (request *DescribePrivateZoneRequest) {
    request = &DescribePrivateZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribePrivateZone")
    return
}

func NewDescribePrivateZoneResponse() (response *DescribePrivateZoneResponse) {
    response = &DescribePrivateZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取私有域信息
func (c *Client) DescribePrivateZone(request *DescribePrivateZoneRequest) (response *DescribePrivateZoneResponse, err error) {
    if request == nil {
        request = NewDescribePrivateZoneRequest()
    }
    response = NewDescribePrivateZoneResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateZoneListRequest() (request *DescribePrivateZoneListRequest) {
    request = &DescribePrivateZoneListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribePrivateZoneList")
    return
}

func NewDescribePrivateZoneListResponse() (response *DescribePrivateZoneListResponse) {
    response = &DescribePrivateZoneListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取私有域列表
func (c *Client) DescribePrivateZoneList(request *DescribePrivateZoneListRequest) (response *DescribePrivateZoneListResponse, err error) {
    if request == nil {
        request = NewDescribePrivateZoneListRequest()
    }
    response = NewDescribePrivateZoneListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateZoneRecordListRequest() (request *DescribePrivateZoneRecordListRequest) {
    request = &DescribePrivateZoneRecordListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribePrivateZoneRecordList")
    return
}

func NewDescribePrivateZoneRecordListResponse() (response *DescribePrivateZoneRecordListResponse) {
    response = &DescribePrivateZoneRecordListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取私有域记录列表
func (c *Client) DescribePrivateZoneRecordList(request *DescribePrivateZoneRecordListRequest) (response *DescribePrivateZoneRecordListResponse, err error) {
    if request == nil {
        request = NewDescribePrivateZoneRecordListRequest()
    }
    response = NewDescribePrivateZoneRecordListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePrivateZoneServiceRequest() (request *DescribePrivateZoneServiceRequest) {
    request = &DescribePrivateZoneServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribePrivateZoneService")
    return
}

func NewDescribePrivateZoneServiceResponse() (response *DescribePrivateZoneServiceResponse) {
    response = &DescribePrivateZoneServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询私有域解析开通状态
func (c *Client) DescribePrivateZoneService(request *DescribePrivateZoneServiceRequest) (response *DescribePrivateZoneServiceResponse, err error) {
    if request == nil {
        request = NewDescribePrivateZoneServiceRequest()
    }
    response = NewDescribePrivateZoneServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRequestDataRequest() (request *DescribeRequestDataRequest) {
    request = &DescribeRequestDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "DescribeRequestData")
    return
}

func NewDescribeRequestDataResponse() (response *DescribeRequestDataResponse) {
    response = &DescribeRequestDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取私有域解析请求量
func (c *Client) DescribeRequestData(request *DescribeRequestDataRequest) (response *DescribeRequestDataResponse, err error) {
    if request == nil {
        request = NewDescribeRequestDataRequest()
    }
    response = NewDescribeRequestDataResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrivateZoneRequest() (request *ModifyPrivateZoneRequest) {
    request = &ModifyPrivateZoneRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "ModifyPrivateZone")
    return
}

func NewModifyPrivateZoneResponse() (response *ModifyPrivateZoneResponse) {
    response = &ModifyPrivateZoneResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改私有域信息
func (c *Client) ModifyPrivateZone(request *ModifyPrivateZoneRequest) (response *ModifyPrivateZoneResponse, err error) {
    if request == nil {
        request = NewModifyPrivateZoneRequest()
    }
    response = NewModifyPrivateZoneResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrivateZoneRecordRequest() (request *ModifyPrivateZoneRecordRequest) {
    request = &ModifyPrivateZoneRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "ModifyPrivateZoneRecord")
    return
}

func NewModifyPrivateZoneRecordResponse() (response *ModifyPrivateZoneRecordResponse) {
    response = &ModifyPrivateZoneRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改私有域解析记录
func (c *Client) ModifyPrivateZoneRecord(request *ModifyPrivateZoneRecordRequest) (response *ModifyPrivateZoneRecordResponse, err error) {
    if request == nil {
        request = NewModifyPrivateZoneRecordRequest()
    }
    response = NewModifyPrivateZoneRecordResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPrivateZoneVpcRequest() (request *ModifyPrivateZoneVpcRequest) {
    request = &ModifyPrivateZoneVpcRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "ModifyPrivateZoneVpc")
    return
}

func NewModifyPrivateZoneVpcResponse() (response *ModifyPrivateZoneVpcResponse) {
    response = &ModifyPrivateZoneVpcResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改私有域关联的VPC
func (c *Client) ModifyPrivateZoneVpc(request *ModifyPrivateZoneVpcRequest) (response *ModifyPrivateZoneVpcResponse, err error) {
    if request == nil {
        request = NewModifyPrivateZoneVpcRequest()
    }
    response = NewModifyPrivateZoneVpcResponse()
    err = c.Send(request, response)
    return
}

func NewSubscribePrivateZoneServiceRequest() (request *SubscribePrivateZoneServiceRequest) {
    request = &SubscribePrivateZoneServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("privatedns", APIVersion, "SubscribePrivateZoneService")
    return
}

func NewSubscribePrivateZoneServiceResponse() (response *SubscribePrivateZoneServiceResponse) {
    response = &SubscribePrivateZoneServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 开通私有域解析
func (c *Client) SubscribePrivateZoneService(request *SubscribePrivateZoneServiceRequest) (response *SubscribePrivateZoneServiceResponse, err error) {
    if request == nil {
        request = NewSubscribePrivateZoneServiceRequest()
    }
    response = NewSubscribePrivateZoneServiceResponse()
    err = c.Send(request, response)
    return
}
