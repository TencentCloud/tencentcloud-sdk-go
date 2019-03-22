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

package v20180420

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-20"

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


func NewDescribeDbauditInstanceTypeRequest() (request *DescribeDbauditInstanceTypeRequest) {
    request = &DescribeDbauditInstanceTypeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cds", APIVersion, "DescribeDbauditInstanceType")
    return
}

func NewDescribeDbauditInstanceTypeResponse() (response *DescribeDbauditInstanceTypeResponse) {
    response = &DescribeDbauditInstanceTypeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeDbauditInstanceType) 用于查询可售卖的产品规格列表。
func (c *Client) DescribeDbauditInstanceType(request *DescribeDbauditInstanceTypeRequest) (response *DescribeDbauditInstanceTypeResponse, err error) {
    if request == nil {
        request = NewDescribeDbauditInstanceTypeRequest()
    }
    response = NewDescribeDbauditInstanceTypeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDbauditInstancesRequest() (request *DescribeDbauditInstancesRequest) {
    request = &DescribeDbauditInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cds", APIVersion, "DescribeDbauditInstances")
    return
}

func NewDescribeDbauditInstancesResponse() (response *DescribeDbauditInstancesResponse) {
    response = &DescribeDbauditInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeDbauditInstances) 用于查询数据安全审计实例列表
func (c *Client) DescribeDbauditInstances(request *DescribeDbauditInstancesRequest) (response *DescribeDbauditInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDbauditInstancesRequest()
    }
    response = NewDescribeDbauditInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDbauditUsedRegionsRequest() (request *DescribeDbauditUsedRegionsRequest) {
    request = &DescribeDbauditUsedRegionsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cds", APIVersion, "DescribeDbauditUsedRegions")
    return
}

func NewDescribeDbauditUsedRegionsResponse() (response *DescribeDbauditUsedRegionsResponse) {
    response = &DescribeDbauditUsedRegionsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeDbauditUsedRegions) 用于查询可售卖地域列表。
func (c *Client) DescribeDbauditUsedRegions(request *DescribeDbauditUsedRegionsRequest) (response *DescribeDbauditUsedRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeDbauditUsedRegionsRequest()
    }
    response = NewDescribeDbauditUsedRegionsResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceDbauditInstanceRequest() (request *InquiryPriceDbauditInstanceRequest) {
    request = &InquiryPriceDbauditInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cds", APIVersion, "InquiryPriceDbauditInstance")
    return
}

func NewInquiryPriceDbauditInstanceResponse() (response *InquiryPriceDbauditInstanceResponse) {
    response = &InquiryPriceDbauditInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于查询数据安全审计产品实例价格
func (c *Client) InquiryPriceDbauditInstance(request *InquiryPriceDbauditInstanceRequest) (response *InquiryPriceDbauditInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceDbauditInstanceRequest()
    }
    response = NewInquiryPriceDbauditInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDbauditInstancesRenewFlagRequest() (request *ModifyDbauditInstancesRenewFlagRequest) {
    request = &ModifyDbauditInstancesRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cds", APIVersion, "ModifyDbauditInstancesRenewFlag")
    return
}

func NewModifyDbauditInstancesRenewFlagResponse() (response *ModifyDbauditInstancesRenewFlagResponse) {
    response = &ModifyDbauditInstancesRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (ModifyDbauditInstancesRenewFlag) 用于修改数据安全审计产品实例续费标识
func (c *Client) ModifyDbauditInstancesRenewFlag(request *ModifyDbauditInstancesRenewFlagRequest) (response *ModifyDbauditInstancesRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyDbauditInstancesRenewFlagRequest()
    }
    response = NewModifyDbauditInstancesRenewFlagResponse()
    err = c.Send(request, response)
    return
}
