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

package v20200721

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-07-21"

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


func NewDescribeStrategiesRequest() (request *DescribeStrategiesRequest) {
    request = &DescribeStrategiesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("advisor", APIVersion, "DescribeStrategies")
    
    
    return
}

func NewDescribeStrategiesResponse() (response *DescribeStrategiesResponse) {
    response = &DescribeStrategiesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeStrategies
// 用于查询评估项的信息
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeStrategies(request *DescribeStrategiesRequest) (response *DescribeStrategiesResponse, err error) {
    return c.DescribeStrategiesWithContext(context.Background(), request)
}

// DescribeStrategies
// 用于查询评估项的信息
//
// 可能返回的错误码:
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeStrategiesWithContext(ctx context.Context, request *DescribeStrategiesRequest) (response *DescribeStrategiesResponse, err error) {
    if request == nil {
        request = NewDescribeStrategiesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStrategies require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStrategiesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskStrategyRisksRequest() (request *DescribeTaskStrategyRisksRequest) {
    request = &DescribeTaskStrategyRisksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("advisor", APIVersion, "DescribeTaskStrategyRisks")
    
    
    return
}

func NewDescribeTaskStrategyRisksResponse() (response *DescribeTaskStrategyRisksResponse) {
    response = &DescribeTaskStrategyRisksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskStrategyRisks
// 查询评估项风险实例列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTaskStrategyRisks(request *DescribeTaskStrategyRisksRequest) (response *DescribeTaskStrategyRisksResponse, err error) {
    return c.DescribeTaskStrategyRisksWithContext(context.Background(), request)
}

// DescribeTaskStrategyRisks
// 查询评估项风险实例列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DEPENDSDB = "InternalError.DependsDb"
//  INTERNALERROR_SYSTEM = "InternalError.System"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeTaskStrategyRisksWithContext(ctx context.Context, request *DescribeTaskStrategyRisksRequest) (response *DescribeTaskStrategyRisksResponse, err error) {
    if request == nil {
        request = NewDescribeTaskStrategyRisksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskStrategyRisks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskStrategyRisksResponse()
    err = c.Send(request, response)
    return
}
