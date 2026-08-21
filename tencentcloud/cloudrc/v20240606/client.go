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

package v20240606

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2024-06-06"

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


func NewDescribeResourceRequest() (request *DescribeResourceRequest) {
    request = &DescribeResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudrc", APIVersion, "DescribeResource")
    
    
    return
}

func NewDescribeResourceResponse() (response *DescribeResourceResponse) {
    response = &DescribeResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeResource
// 查询资源详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_RESOURCEIDNOTFOUND = "ResourceNotFound.ResourceIdNotFound"
//  RESOURCENOTFOUND_VIEWIDNOTFOUND = "ResourceNotFound.ViewIdNotFound"
func (c *Client) DescribeResource(request *DescribeResourceRequest) (response *DescribeResourceResponse, err error) {
    return c.DescribeResourceWithContext(context.Background(), request)
}

// DescribeResource
// 查询资源详情
//
// 可能返回的错误码:
//  RESOURCENOTFOUND_RESOURCEIDNOTFOUND = "ResourceNotFound.ResourceIdNotFound"
//  RESOURCENOTFOUND_VIEWIDNOTFOUND = "ResourceNotFound.ViewIdNotFound"
func (c *Client) DescribeResourceWithContext(ctx context.Context, request *DescribeResourceRequest) (response *DescribeResourceResponse, err error) {
    if request == nil {
        request = NewDescribeResourceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cloudrc", APIVersion, "DescribeResource")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceResponse()
    err = c.Send(request, response)
    return
}

func NewSearchResourcesRequest() (request *SearchResourcesRequest) {
    request = &SearchResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudrc", APIVersion, "SearchResources")
    
    
    return
}

func NewSearchResourcesResponse() (response *SearchResourcesResponse) {
    response = &SearchResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchResources
// 搜索资源
//
// 可能返回的错误码:
//  INVALIDPARAMETER_APIPARAMERROR = "InvalidParameter.ApiParamError"
//  RESOURCENOTFOUND_VIEWIDNOTFOUND = "ResourceNotFound.ViewIdNotFound"
func (c *Client) SearchResources(request *SearchResourcesRequest) (response *SearchResourcesResponse, err error) {
    return c.SearchResourcesWithContext(context.Background(), request)
}

// SearchResources
// 搜索资源
//
// 可能返回的错误码:
//  INVALIDPARAMETER_APIPARAMERROR = "InvalidParameter.ApiParamError"
//  RESOURCENOTFOUND_VIEWIDNOTFOUND = "ResourceNotFound.ViewIdNotFound"
func (c *Client) SearchResourcesWithContext(ctx context.Context, request *SearchResourcesRequest) (response *SearchResourcesResponse, err error) {
    if request == nil {
        request = NewSearchResourcesRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cloudrc", APIVersion, "SearchResources")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchResourcesResponse()
    err = c.Send(request, response)
    return
}
