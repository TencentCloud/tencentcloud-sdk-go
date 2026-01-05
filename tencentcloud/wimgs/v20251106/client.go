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

package v20251106

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2025-11-06"

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


func NewSearchByTextRequest() (request *SearchByTextRequest) {
    request = &SearchByTextRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("wimgs", APIVersion, "SearchByText")
    
    
    return
}

func NewSearchByTextResponse() (response *SearchByTextResponse) {
    response = &SearchByTextResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchByText
// 文搜图接口，本服务将针对您输入的搜索关键词，检索并以JSON形式返回相关图片的标题、宽高、缩略图、内容来源url等信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SearchByText(request *SearchByTextRequest) (response *SearchByTextResponse, err error) {
    return c.SearchByTextWithContext(context.Background(), request)
}

// SearchByText
// 文搜图接口，本服务将针对您输入的搜索关键词，检索并以JSON形式返回相关图片的标题、宽高、缩略图、内容来源url等信息。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SearchByTextWithContext(ctx context.Context, request *SearchByTextRequest) (response *SearchByTextResponse, err error) {
    if request == nil {
        request = NewSearchByTextRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wimgs", APIVersion, "SearchByText")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchByText require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchByTextResponse()
    err = c.Send(request, response)
    return
}
