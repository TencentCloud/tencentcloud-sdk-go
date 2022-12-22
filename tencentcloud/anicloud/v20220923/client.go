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

package v20220923

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-09-23"

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


func NewCheckAppidExistRequest() (request *CheckAppidExistRequest) {
    request = &CheckAppidExistRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("anicloud", APIVersion, "CheckAppidExist")
    
    
    return
}

func NewCheckAppidExistResponse() (response *CheckAppidExistResponse) {
    response = &CheckAppidExistResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CheckAppidExist
// 查看appid是否存在
func (c *Client) CheckAppidExist(request *CheckAppidExistRequest) (response *CheckAppidExistResponse, err error) {
    return c.CheckAppidExistWithContext(context.Background(), request)
}

// CheckAppidExist
// 查看appid是否存在
func (c *Client) CheckAppidExistWithContext(ctx context.Context, request *CheckAppidExistRequest) (response *CheckAppidExistResponse, err error) {
    if request == nil {
        request = NewCheckAppidExistRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckAppidExist require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckAppidExistResponse()
    err = c.Send(request, response)
    return
}

func NewQueryResourceRequest() (request *QueryResourceRequest) {
    request = &QueryResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("anicloud", APIVersion, "QueryResource")
    
    
    return
}

func NewQueryResourceResponse() (response *QueryResourceResponse) {
    response = &QueryResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryResource
// 查询购买资源
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QueryResource(request *QueryResourceRequest) (response *QueryResourceResponse, err error) {
    return c.QueryResourceWithContext(context.Background(), request)
}

// QueryResource
// 查询购买资源
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QueryResourceWithContext(ctx context.Context, request *QueryResourceRequest) (response *QueryResourceResponse, err error) {
    if request == nil {
        request = NewQueryResourceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryResource require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryResourceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryResourceInfoRequest() (request *QueryResourceInfoRequest) {
    request = &QueryResourceInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("anicloud", APIVersion, "QueryResourceInfo")
    
    
    return
}

func NewQueryResourceInfoResponse() (response *QueryResourceInfoResponse) {
    response = &QueryResourceInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryResourceInfo
// 查询资源信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QueryResourceInfo(request *QueryResourceInfoRequest) (response *QueryResourceInfoResponse, err error) {
    return c.QueryResourceInfoWithContext(context.Background(), request)
}

// QueryResourceInfo
// 查询资源信息
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) QueryResourceInfoWithContext(ctx context.Context, request *QueryResourceInfoRequest) (response *QueryResourceInfoResponse, err error) {
    if request == nil {
        request = NewQueryResourceInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryResourceInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryResourceInfoResponse()
    err = c.Send(request, response)
    return
}
