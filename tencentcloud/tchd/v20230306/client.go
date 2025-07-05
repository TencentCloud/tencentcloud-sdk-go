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

package v20230306

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-03-06"

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


func NewDescribeEventStatisticsRequest() (request *DescribeEventStatisticsRequest) {
    request = &DescribeEventStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tchd", APIVersion, "DescribeEventStatistics")
    
    
    return
}

func NewDescribeEventStatisticsResponse() (response *DescribeEventStatisticsResponse) {
    response = &DescribeEventStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEventStatistics
// 本接口用于查询腾讯云健康看板的实时可用性事件信息，可以通过产品列表、地域进行过滤查询。
//
// 可以参考健康看板历史事件页面来获取查询案例（链接：https://status.cloud.tencent.com/history）。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeEventStatistics(request *DescribeEventStatisticsRequest) (response *DescribeEventStatisticsResponse, err error) {
    return c.DescribeEventStatisticsWithContext(context.Background(), request)
}

// DescribeEventStatistics
// 本接口用于查询腾讯云健康看板的实时可用性事件信息，可以通过产品列表、地域进行过滤查询。
//
// 可以参考健康看板历史事件页面来获取查询案例（链接：https://status.cloud.tencent.com/history）。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeEventStatisticsWithContext(ctx context.Context, request *DescribeEventStatisticsRequest) (response *DescribeEventStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeEventStatisticsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEventStatistics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEventStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEventsRequest() (request *DescribeEventsRequest) {
    request = &DescribeEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tchd", APIVersion, "DescribeEvents")
    
    
    return
}

func NewDescribeEventsResponse() (response *DescribeEventsResponse) {
    response = &DescribeEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeEvents
// 本接口用于查询腾讯云健康看板的可用性事件信息，可以通过产品列表、地域列表和事件发生日期进行过滤查询。
//
// 当查询的产品对应时间内无事件时将返回空结果。
//
// 可以参考健康看板历史事件页面来获取查询案例（链接：https://status.cloud.tencent.com/history）。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeEvents(request *DescribeEventsRequest) (response *DescribeEventsResponse, err error) {
    return c.DescribeEventsWithContext(context.Background(), request)
}

// DescribeEvents
// 本接口用于查询腾讯云健康看板的可用性事件信息，可以通过产品列表、地域列表和事件发生日期进行过滤查询。
//
// 当查询的产品对应时间内无事件时将返回空结果。
//
// 可以参考健康看板历史事件页面来获取查询案例（链接：https://status.cloud.tencent.com/history）。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeEventsWithContext(ctx context.Context, request *DescribeEventsRequest) (response *DescribeEventsResponse, err error) {
    if request == nil {
        request = NewDescribeEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeEventsResponse()
    err = c.Send(request, response)
    return
}
