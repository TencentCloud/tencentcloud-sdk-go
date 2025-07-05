// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20240125

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2024-01-25"

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


func NewCreateSavingPlanOrderRequest() (request *CreateSavingPlanOrderRequest) {
    request = &CreateSavingPlanOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("svp", APIVersion, "CreateSavingPlanOrder")
    
    
    return
}

func NewCreateSavingPlanOrderResponse() (response *CreateSavingPlanOrderResponse) {
    response = &CreateSavingPlanOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSavingPlanOrder
// 创建节省计划订单
func (c *Client) CreateSavingPlanOrder(request *CreateSavingPlanOrderRequest) (response *CreateSavingPlanOrderResponse, err error) {
    return c.CreateSavingPlanOrderWithContext(context.Background(), request)
}

// CreateSavingPlanOrder
// 创建节省计划订单
func (c *Client) CreateSavingPlanOrderWithContext(ctx context.Context, request *CreateSavingPlanOrderRequest) (response *CreateSavingPlanOrderResponse, err error) {
    if request == nil {
        request = NewCreateSavingPlanOrderRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSavingPlanOrder require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSavingPlanOrderResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSavingPlanCoverageRequest() (request *DescribeSavingPlanCoverageRequest) {
    request = &DescribeSavingPlanCoverageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("svp", APIVersion, "DescribeSavingPlanCoverage")
    
    
    return
}

func NewDescribeSavingPlanCoverageResponse() (response *DescribeSavingPlanCoverageResponse) {
    response = &DescribeSavingPlanCoverageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSavingPlanCoverage
// 查询当前用户节省计划覆盖率明细数据，如无特别说明，金额单位均为元（国内站）或者美元（国际站）。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeSavingPlanCoverage(request *DescribeSavingPlanCoverageRequest) (response *DescribeSavingPlanCoverageResponse, err error) {
    return c.DescribeSavingPlanCoverageWithContext(context.Background(), request)
}

// DescribeSavingPlanCoverage
// 查询当前用户节省计划覆盖率明细数据，如无特别说明，金额单位均为元（国内站）或者美元（国际站）。
//
// 可能返回的错误码:
//  INTERNALERROR_UNKNOWNERROR = "InternalError.UnknownError"
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeSavingPlanCoverageWithContext(ctx context.Context, request *DescribeSavingPlanCoverageRequest) (response *DescribeSavingPlanCoverageResponse, err error) {
    if request == nil {
        request = NewDescribeSavingPlanCoverageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSavingPlanCoverage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSavingPlanCoverageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSavingPlanDeductRequest() (request *DescribeSavingPlanDeductRequest) {
    request = &DescribeSavingPlanDeductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("svp", APIVersion, "DescribeSavingPlanDeduct")
    
    
    return
}

func NewDescribeSavingPlanDeductResponse() (response *DescribeSavingPlanDeductResponse) {
    response = &DescribeSavingPlanDeductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSavingPlanDeduct
// 查询节省计划抵扣明细
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeSavingPlanDeduct(request *DescribeSavingPlanDeductRequest) (response *DescribeSavingPlanDeductResponse, err error) {
    return c.DescribeSavingPlanDeductWithContext(context.Background(), request)
}

// DescribeSavingPlanDeduct
// 查询节省计划抵扣明细
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeSavingPlanDeductWithContext(ctx context.Context, request *DescribeSavingPlanDeductRequest) (response *DescribeSavingPlanDeductResponse, err error) {
    if request == nil {
        request = NewDescribeSavingPlanDeductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSavingPlanDeduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSavingPlanDeductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSavingPlanOverviewRequest() (request *DescribeSavingPlanOverviewRequest) {
    request = &DescribeSavingPlanOverviewRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("svp", APIVersion, "DescribeSavingPlanOverview")
    
    
    return
}

func NewDescribeSavingPlanOverviewResponse() (response *DescribeSavingPlanOverviewResponse) {
    response = &DescribeSavingPlanOverviewResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSavingPlanOverview
// 查用当前用户明细节省计划总览查询时段内的使用情况
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeSavingPlanOverview(request *DescribeSavingPlanOverviewRequest) (response *DescribeSavingPlanOverviewResponse, err error) {
    return c.DescribeSavingPlanOverviewWithContext(context.Background(), request)
}

// DescribeSavingPlanOverview
// 查用当前用户明细节省计划总览查询时段内的使用情况
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeSavingPlanOverviewWithContext(ctx context.Context, request *DescribeSavingPlanOverviewRequest) (response *DescribeSavingPlanOverviewResponse, err error) {
    if request == nil {
        request = NewDescribeSavingPlanOverviewRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSavingPlanOverview require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSavingPlanOverviewResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSavingPlanUsageRequest() (request *DescribeSavingPlanUsageRequest) {
    request = &DescribeSavingPlanUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("svp", APIVersion, "DescribeSavingPlanUsage")
    
    
    return
}

func NewDescribeSavingPlanUsageResponse() (response *DescribeSavingPlanUsageResponse) {
    response = &DescribeSavingPlanUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSavingPlanUsage
// 查用当前用户明细节省计划查询时段内的使用情况
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeSavingPlanUsage(request *DescribeSavingPlanUsageRequest) (response *DescribeSavingPlanUsageResponse, err error) {
    return c.DescribeSavingPlanUsageWithContext(context.Background(), request)
}

// DescribeSavingPlanUsage
// 查用当前用户明细节省计划查询时段内的使用情况
//
// 可能返回的错误码:
//  UNAUTHORIZEDOPERATION_CAMNOAUTH = "UnauthorizedOperation.CamNoAuth"
func (c *Client) DescribeSavingPlanUsageWithContext(ctx context.Context, request *DescribeSavingPlanUsageRequest) (response *DescribeSavingPlanUsageResponse, err error) {
    if request == nil {
        request = NewDescribeSavingPlanUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSavingPlanUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSavingPlanUsageResponse()
    err = c.Send(request, response)
    return
}
