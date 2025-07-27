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

package v20250508

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2025-05-08"

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


func NewSearchProRequest() (request *SearchProRequest) {
    request = &SearchProRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wsa", APIVersion, "SearchPro")
    
    
    return
}

func NewSearchProResponse() (response *SearchProResponse) {
    response = &SearchProResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchPro
// 联网搜索API，以json形式向客户提供搜索结果数据，包含不仅限于标题、摘要、内容来源url等信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SearchPro(request *SearchProRequest) (response *SearchProResponse, err error) {
    return c.SearchProWithContext(context.Background(), request)
}

// SearchPro
// 联网搜索API，以json形式向客户提供搜索结果数据，包含不仅限于标题、摘要、内容来源url等信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SearchProWithContext(ctx context.Context, request *SearchProRequest) (response *SearchProResponse, err error) {
    if request == nil {
        request = NewSearchProRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchPro require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchProResponse()
    err = c.Send(request, response)
    return
}
