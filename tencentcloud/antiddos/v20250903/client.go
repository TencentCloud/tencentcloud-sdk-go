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

package v20250903

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2025-09-03"

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


func NewDescribeDDoSBlockRecordsRequest() (request *DescribeDDoSBlockRecordsRequest) {
    request = &DescribeDDoSBlockRecordsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("antiddos", APIVersion, "DescribeDDoSBlockRecords")
    
    
    return
}

func NewDescribeDDoSBlockRecordsResponse() (response *DescribeDDoSBlockRecordsResponse) {
    response = &DescribeDDoSBlockRecordsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDDoSBlockRecords
// 查询封堵解封记录和解封配额信息。
func (c *Client) DescribeDDoSBlockRecords(request *DescribeDDoSBlockRecordsRequest) (response *DescribeDDoSBlockRecordsResponse, err error) {
    return c.DescribeDDoSBlockRecordsWithContext(context.Background(), request)
}

// DescribeDDoSBlockRecords
// 查询封堵解封记录和解封配额信息。
func (c *Client) DescribeDDoSBlockRecordsWithContext(ctx context.Context, request *DescribeDDoSBlockRecordsRequest) (response *DescribeDDoSBlockRecordsResponse, err error) {
    if request == nil {
        request = NewDescribeDDoSBlockRecordsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "antiddos", APIVersion, "DescribeDDoSBlockRecords")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDDoSBlockRecords require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDDoSBlockRecordsResponse()
    err = c.Send(request, response)
    return
}

func NewUnblockResourcesRequest() (request *UnblockResourcesRequest) {
    request = &UnblockResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("antiddos", APIVersion, "UnblockResources")
    
    
    return
}

func NewUnblockResourcesResponse() (response *UnblockResourcesResponse) {
    response = &UnblockResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UnblockResources
// 申请解封资源，可通过 DescribeDDoSBlockRecords 接口获取资源的封堵解封状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) UnblockResources(request *UnblockResourcesRequest) (response *UnblockResourcesResponse, err error) {
    return c.UnblockResourcesWithContext(context.Background(), request)
}

// UnblockResources
// 申请解封资源，可通过 DescribeDDoSBlockRecords 接口获取资源的封堵解封状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) UnblockResourcesWithContext(ctx context.Context, request *UnblockResourcesRequest) (response *UnblockResourcesResponse, err error) {
    if request == nil {
        request = NewUnblockResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "antiddos", APIVersion, "UnblockResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UnblockResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewUnblockResourcesResponse()
    err = c.Send(request, response)
    return
}
