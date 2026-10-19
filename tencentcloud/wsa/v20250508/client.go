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
// 元宝搜索（Tencent Yuanbao Search，wsa）以互联网全网公开资源为基础，叠加腾讯优质内容生态，结合元宝 App 联网搜索应用实践，从数据收录到智能检索召回，全链路构建 AI 友好搜索引擎。合作伙伴通过检索词发起请求，接口以 JSON 形式返回搜索结果对应的排序信息、数据字段，合作伙伴可在业务场景中结合搜索结果的内容，扩充信息内容源，丰富展现效果，提升用户的查询满意度。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SearchPro(request *SearchProRequest) (response *SearchProResponse, err error) {
    return c.SearchProWithContext(context.Background(), request)
}

// SearchPro
// 元宝搜索（Tencent Yuanbao Search，wsa）以互联网全网公开资源为基础，叠加腾讯优质内容生态，结合元宝 App 联网搜索应用实践，从数据收录到智能检索召回，全链路构建 AI 友好搜索引擎。合作伙伴通过检索词发起请求，接口以 JSON 形式返回搜索结果对应的排序信息、数据字段，合作伙伴可在业务场景中结合搜索结果的内容，扩充信息内容源，丰富展现效果，提升用户的查询满意度。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) SearchProWithContext(ctx context.Context, request *SearchProRequest) (response *SearchProResponse, err error) {
    if request == nil {
        request = NewSearchProRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "wsa", APIVersion, "SearchPro")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchPro require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchProResponse()
    err = c.Send(request, response)
    return
}
