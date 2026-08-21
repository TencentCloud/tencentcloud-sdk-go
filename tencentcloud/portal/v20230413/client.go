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

package v20230413

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-04-13"

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


func NewSearchDocumentsRequest() (request *SearchDocumentsRequest) {
    request = &SearchDocumentsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("portal", APIVersion, "SearchDocuments")
    
    
    return
}

func NewSearchDocumentsResponse() (response *SearchDocumentsResponse) {
    response = &SearchDocumentsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SearchDocuments
// 通过关键词搜索文档列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchDocuments(request *SearchDocumentsRequest) (response *SearchDocumentsResponse, err error) {
    return c.SearchDocumentsWithContext(context.Background(), request)
}

// SearchDocuments
// 通过关键词搜索文档列表
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) SearchDocumentsWithContext(ctx context.Context, request *SearchDocumentsRequest) (response *SearchDocumentsResponse, err error) {
    if request == nil {
        request = NewSearchDocumentsRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "portal", APIVersion, "SearchDocuments")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchDocuments require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchDocumentsResponse()
    err = c.Send(request, response)
    return
}
