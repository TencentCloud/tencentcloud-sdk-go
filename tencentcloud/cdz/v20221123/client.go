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

package v20221123

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-11-23"

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


func NewDescribeCloudDedicatedZoneHostsRequest() (request *DescribeCloudDedicatedZoneHostsRequest) {
    request = &DescribeCloudDedicatedZoneHostsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdz", APIVersion, "DescribeCloudDedicatedZoneHosts")
    
    
    return
}

func NewDescribeCloudDedicatedZoneHostsResponse() (response *DescribeCloudDedicatedZoneHostsResponse) {
    response = &DescribeCloudDedicatedZoneHostsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudDedicatedZoneHosts
// 查询可用区的Host和Host上部署的实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDACTION = "FailedOperation.InvalidAction"
//  INVALIDPARAMETER_HOSTUUIDSANDINSIDSCANNOTAPPEARSAMETIME = "InvalidParameter.HostUuidsAndInsIdsCannotAppearSameTime"
//  RESOURCENOTFOUND_CDZIDNOTFOUND = "ResourceNotFound.CdzIdNotFound"
func (c *Client) DescribeCloudDedicatedZoneHosts(request *DescribeCloudDedicatedZoneHostsRequest) (response *DescribeCloudDedicatedZoneHostsResponse, err error) {
    return c.DescribeCloudDedicatedZoneHostsWithContext(context.Background(), request)
}

// DescribeCloudDedicatedZoneHosts
// 查询可用区的Host和Host上部署的实例
//
// 可能返回的错误码:
//  FAILEDOPERATION_INVALIDACTION = "FailedOperation.InvalidAction"
//  INVALIDPARAMETER_HOSTUUIDSANDINSIDSCANNOTAPPEARSAMETIME = "InvalidParameter.HostUuidsAndInsIdsCannotAppearSameTime"
//  RESOURCENOTFOUND_CDZIDNOTFOUND = "ResourceNotFound.CdzIdNotFound"
func (c *Client) DescribeCloudDedicatedZoneHostsWithContext(ctx context.Context, request *DescribeCloudDedicatedZoneHostsRequest) (response *DescribeCloudDedicatedZoneHostsResponse, err error) {
    if request == nil {
        request = NewDescribeCloudDedicatedZoneHostsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudDedicatedZoneHosts require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudDedicatedZoneHostsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudDedicatedZoneResourceSummaryRequest() (request *DescribeCloudDedicatedZoneResourceSummaryRequest) {
    request = &DescribeCloudDedicatedZoneResourceSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cdz", APIVersion, "DescribeCloudDedicatedZoneResourceSummary")
    
    
    return
}

func NewDescribeCloudDedicatedZoneResourceSummaryResponse() (response *DescribeCloudDedicatedZoneResourceSummaryResponse) {
    response = &DescribeCloudDedicatedZoneResourceSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudDedicatedZoneResourceSummary
// 查询专属可用区各个垂直产品的资源使用情况
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CDZIDNOTFOUND = "ResourceNotFound.CdzIdNotFound"
func (c *Client) DescribeCloudDedicatedZoneResourceSummary(request *DescribeCloudDedicatedZoneResourceSummaryRequest) (response *DescribeCloudDedicatedZoneResourceSummaryResponse, err error) {
    return c.DescribeCloudDedicatedZoneResourceSummaryWithContext(context.Background(), request)
}

// DescribeCloudDedicatedZoneResourceSummary
// 查询专属可用区各个垂直产品的资源使用情况
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_CDZIDNOTFOUND = "ResourceNotFound.CdzIdNotFound"
func (c *Client) DescribeCloudDedicatedZoneResourceSummaryWithContext(ctx context.Context, request *DescribeCloudDedicatedZoneResourceSummaryRequest) (response *DescribeCloudDedicatedZoneResourceSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeCloudDedicatedZoneResourceSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudDedicatedZoneResourceSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudDedicatedZoneResourceSummaryResponse()
    err = c.Send(request, response)
    return
}
