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

package v20230215

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-02-15"

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


func NewDescribeDrawResourceListRequest() (request *DescribeDrawResourceListRequest) {
    request = &DescribeDrawResourceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tourism", APIVersion, "DescribeDrawResourceList")
    
    
    return
}

func NewDescribeDrawResourceListResponse() (response *DescribeDrawResourceListResponse) {
    response = &DescribeDrawResourceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeDrawResourceList
// 依据客户的Uin查询开通的资源列表
func (c *Client) DescribeDrawResourceList(request *DescribeDrawResourceListRequest) (response *DescribeDrawResourceListResponse, err error) {
    return c.DescribeDrawResourceListWithContext(context.Background(), request)
}

// DescribeDrawResourceList
// 依据客户的Uin查询开通的资源列表
func (c *Client) DescribeDrawResourceListWithContext(ctx context.Context, request *DescribeDrawResourceListRequest) (response *DescribeDrawResourceListResponse, err error) {
    if request == nil {
        request = NewDescribeDrawResourceListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "tourism", APIVersion, "DescribeDrawResourceList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDrawResourceList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDrawResourceListResponse()
    err = c.Send(request, response)
    return
}
