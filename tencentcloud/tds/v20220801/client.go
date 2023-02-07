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

package v20220801

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-08-01"

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


func NewDescribeFraudBaseRequest() (request *DescribeFraudBaseRequest) {
    request = &DescribeFraudBaseRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tds", APIVersion, "DescribeFraudBase")
    
    
    return
}

func NewDescribeFraudBaseResponse() (response *DescribeFraudBaseResponse) {
    response = &DescribeFraudBaseResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFraudBase
// 查询设备风险
//
// 可能返回的错误码:
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) DescribeFraudBase(request *DescribeFraudBaseRequest) (response *DescribeFraudBaseResponse, err error) {
    return c.DescribeFraudBaseWithContext(context.Background(), request)
}

// DescribeFraudBase
// 查询设备风险
//
// 可能返回的错误码:
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) DescribeFraudBaseWithContext(ctx context.Context, request *DescribeFraudBaseRequest) (response *DescribeFraudBaseResponse, err error) {
    if request == nil {
        request = NewDescribeFraudBaseRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFraudBase require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFraudBaseResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFraudPremiumRequest() (request *DescribeFraudPremiumRequest) {
    request = &DescribeFraudPremiumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tds", APIVersion, "DescribeFraudPremium")
    
    
    return
}

func NewDescribeFraudPremiumResponse() (response *DescribeFraudPremiumResponse) {
    response = &DescribeFraudPremiumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFraudPremium
// 查询设备标识及风险
//
// 可能返回的错误码:
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) DescribeFraudPremium(request *DescribeFraudPremiumRequest) (response *DescribeFraudPremiumResponse, err error) {
    return c.DescribeFraudPremiumWithContext(context.Background(), request)
}

// DescribeFraudPremium
// 查询设备标识及风险
//
// 可能返回的错误码:
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) DescribeFraudPremiumWithContext(ctx context.Context, request *DescribeFraudPremiumRequest) (response *DescribeFraudPremiumResponse, err error) {
    if request == nil {
        request = NewDescribeFraudPremiumRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFraudPremium require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFraudPremiumResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFraudUltimateRequest() (request *DescribeFraudUltimateRequest) {
    request = &DescribeFraudUltimateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tds", APIVersion, "DescribeFraudUltimate")
    
    
    return
}

func NewDescribeFraudUltimateResponse() (response *DescribeFraudUltimateResponse) {
    response = &DescribeFraudUltimateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFraudUltimate
// 查询设备标识及风险（旗舰版）
//
// 可能返回的错误码:
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) DescribeFraudUltimate(request *DescribeFraudUltimateRequest) (response *DescribeFraudUltimateResponse, err error) {
    return c.DescribeFraudUltimateWithContext(context.Background(), request)
}

// DescribeFraudUltimate
// 查询设备标识及风险（旗舰版）
//
// 可能返回的错误码:
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) DescribeFraudUltimateWithContext(ctx context.Context, request *DescribeFraudUltimateRequest) (response *DescribeFraudUltimateResponse, err error) {
    if request == nil {
        request = NewDescribeFraudUltimateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFraudUltimate require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFraudUltimateResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrustedIDRequest() (request *DescribeTrustedIDRequest) {
    request = &DescribeTrustedIDRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tds", APIVersion, "DescribeTrustedID")
    
    
    return
}

func NewDescribeTrustedIDResponse() (response *DescribeTrustedIDResponse) {
    response = &DescribeTrustedIDResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTrustedID
// 查询设备标识
//
// 可能返回的错误码:
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) DescribeTrustedID(request *DescribeTrustedIDRequest) (response *DescribeTrustedIDResponse, err error) {
    return c.DescribeTrustedIDWithContext(context.Background(), request)
}

// DescribeTrustedID
// 查询设备标识
//
// 可能返回的错误码:
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
func (c *Client) DescribeTrustedIDWithContext(ctx context.Context, request *DescribeTrustedIDRequest) (response *DescribeTrustedIDResponse, err error) {
    if request == nil {
        request = NewDescribeTrustedIDRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrustedID require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrustedIDResponse()
    err = c.Send(request, response)
    return
}
