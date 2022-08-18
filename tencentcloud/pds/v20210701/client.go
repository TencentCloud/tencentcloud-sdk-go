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

package v20210701

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-07-01"

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


func NewDescribeNewUserAcquisitionRequest() (request *DescribeNewUserAcquisitionRequest) {
    request = &DescribeNewUserAcquisitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pds", APIVersion, "DescribeNewUserAcquisition")
    
    
    return
}

func NewDescribeNewUserAcquisitionResponse() (response *DescribeNewUserAcquisitionResponse) {
    response = &DescribeNewUserAcquisitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeNewUserAcquisition
// 判断新用户信誉值
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICETIMEOUT = "InternalError.ServiceTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_REPLAYATTACK = "LimitExceeded.ReplayAttack"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNewUserAcquisition(request *DescribeNewUserAcquisitionRequest) (response *DescribeNewUserAcquisitionResponse, err error) {
    return c.DescribeNewUserAcquisitionWithContext(context.Background(), request)
}

// DescribeNewUserAcquisition
// 判断新用户信誉值
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICETIMEOUT = "InternalError.ServiceTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_REPLAYATTACK = "LimitExceeded.ReplayAttack"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeNewUserAcquisitionWithContext(ctx context.Context, request *DescribeNewUserAcquisitionRequest) (response *DescribeNewUserAcquisitionResponse, err error) {
    if request == nil {
        request = NewDescribeNewUserAcquisitionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeNewUserAcquisition require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeNewUserAcquisitionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStockEstimationRequest() (request *DescribeStockEstimationRequest) {
    request = &DescribeStockEstimationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("pds", APIVersion, "DescribeStockEstimation")
    
    
    return
}

func NewDescribeStockEstimationResponse() (response *DescribeStockEstimationResponse) {
    response = &DescribeStockEstimationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStockEstimation
// 查询存量判断服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICETIMEOUT = "InternalError.ServiceTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_REPLAYATTACK = "LimitExceeded.ReplayAttack"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeStockEstimation(request *DescribeStockEstimationRequest) (response *DescribeStockEstimationResponse, err error) {
    return c.DescribeStockEstimationWithContext(context.Background(), request)
}

// DescribeStockEstimation
// 查询存量判断服务
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_SERVICETIMEOUT = "InternalError.ServiceTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_REPLAYATTACK = "LimitExceeded.ReplayAttack"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeStockEstimationWithContext(ctx context.Context, request *DescribeStockEstimationRequest) (response *DescribeStockEstimationResponse, err error) {
    if request == nil {
        request = NewDescribeStockEstimationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStockEstimation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStockEstimationResponse()
    err = c.Send(request, response)
    return
}
