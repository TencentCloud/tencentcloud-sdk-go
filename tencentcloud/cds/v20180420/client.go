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
    "context"
    "errors"
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

func NewClient(credential common.CredentialIface, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewDescribeDasbImageIdsRequest() (request *DescribeDasbImageIdsRequest) {
    request = &DescribeDasbImageIdsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cds", APIVersion, "DescribeDasbImageIds")
    
    
    return
}

func NewDescribeDasbImageIdsResponse() (response *DescribeDasbImageIdsResponse) {
    response = &DescribeDasbImageIdsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDasbImageIds
// 获取镜像列表
func (c *Client) DescribeDasbImageIds(request *DescribeDasbImageIdsRequest) (response *DescribeDasbImageIdsResponse, err error) {
    return c.DescribeDasbImageIdsWithContext(context.Background(), request)
}

// DescribeDasbImageIds
// 获取镜像列表
func (c *Client) DescribeDasbImageIdsWithContext(ctx context.Context, request *DescribeDasbImageIdsRequest) (response *DescribeDasbImageIdsResponse, err error) {
    if request == nil {
        request = NewDescribeDasbImageIdsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDasbImageIds require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDasbImageIdsResponse()
    err = c.Send(request, response)
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

// DescribeDbauditInstanceType
// 本接口 (DescribeDbauditInstanceType) 用于查询可售卖的产品规格列表。
func (c *Client) DescribeDbauditInstanceType(request *DescribeDbauditInstanceTypeRequest) (response *DescribeDbauditInstanceTypeResponse, err error) {
    return c.DescribeDbauditInstanceTypeWithContext(context.Background(), request)
}

// DescribeDbauditInstanceType
// 本接口 (DescribeDbauditInstanceType) 用于查询可售卖的产品规格列表。
func (c *Client) DescribeDbauditInstanceTypeWithContext(ctx context.Context, request *DescribeDbauditInstanceTypeRequest) (response *DescribeDbauditInstanceTypeResponse, err error) {
    if request == nil {
        request = NewDescribeDbauditInstanceTypeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDbauditInstanceType require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDbauditInstances
// 本接口 (DescribeDbauditInstances) 用于查询数据安全审计实例列表
func (c *Client) DescribeDbauditInstances(request *DescribeDbauditInstancesRequest) (response *DescribeDbauditInstancesResponse, err error) {
    return c.DescribeDbauditInstancesWithContext(context.Background(), request)
}

// DescribeDbauditInstances
// 本接口 (DescribeDbauditInstances) 用于查询数据安全审计实例列表
func (c *Client) DescribeDbauditInstancesWithContext(ctx context.Context, request *DescribeDbauditInstancesRequest) (response *DescribeDbauditInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeDbauditInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDbauditInstances require credential")
    }

    request.SetContext(ctx)
    
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

// DescribeDbauditUsedRegions
// 本接口 (DescribeDbauditUsedRegions) 用于查询可售卖地域列表。
func (c *Client) DescribeDbauditUsedRegions(request *DescribeDbauditUsedRegionsRequest) (response *DescribeDbauditUsedRegionsResponse, err error) {
    return c.DescribeDbauditUsedRegionsWithContext(context.Background(), request)
}

// DescribeDbauditUsedRegions
// 本接口 (DescribeDbauditUsedRegions) 用于查询可售卖地域列表。
func (c *Client) DescribeDbauditUsedRegionsWithContext(ctx context.Context, request *DescribeDbauditUsedRegionsRequest) (response *DescribeDbauditUsedRegionsResponse, err error) {
    if request == nil {
        request = NewDescribeDbauditUsedRegionsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDbauditUsedRegions require credential")
    }

    request.SetContext(ctx)
    
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

// InquiryPriceDbauditInstance
// 用于查询数据安全审计产品实例价格
func (c *Client) InquiryPriceDbauditInstance(request *InquiryPriceDbauditInstanceRequest) (response *InquiryPriceDbauditInstanceResponse, err error) {
    return c.InquiryPriceDbauditInstanceWithContext(context.Background(), request)
}

// InquiryPriceDbauditInstance
// 用于查询数据安全审计产品实例价格
func (c *Client) InquiryPriceDbauditInstanceWithContext(ctx context.Context, request *InquiryPriceDbauditInstanceRequest) (response *InquiryPriceDbauditInstanceResponse, err error) {
    if request == nil {
        request = NewInquiryPriceDbauditInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceDbauditInstance require credential")
    }

    request.SetContext(ctx)
    
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

// ModifyDbauditInstancesRenewFlag
// 本接口 (ModifyDbauditInstancesRenewFlag) 用于修改数据安全审计产品实例续费标识
//
// 可能返回的错误码:
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETER_RESOURCEIDERROR = "InvalidParameter.ResourceIdError"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyDbauditInstancesRenewFlag(request *ModifyDbauditInstancesRenewFlagRequest) (response *ModifyDbauditInstancesRenewFlagResponse, err error) {
    return c.ModifyDbauditInstancesRenewFlagWithContext(context.Background(), request)
}

// ModifyDbauditInstancesRenewFlag
// 本接口 (ModifyDbauditInstancesRenewFlag) 用于修改数据安全审计产品实例续费标识
//
// 可能返回的错误码:
//  INVALIDPARAMETER_RESOURCEID = "InvalidParameter.ResourceId"
//  INVALIDPARAMETER_RESOURCEIDERROR = "InvalidParameter.ResourceIdError"
//  UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"
func (c *Client) ModifyDbauditInstancesRenewFlagWithContext(ctx context.Context, request *ModifyDbauditInstancesRenewFlagRequest) (response *ModifyDbauditInstancesRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyDbauditInstancesRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyDbauditInstancesRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyDbauditInstancesRenewFlagResponse()
    err = c.Send(request, response)
    return
}
