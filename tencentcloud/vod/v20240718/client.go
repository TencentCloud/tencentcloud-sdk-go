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

package v20240718

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2024-07-18"

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


func NewCreateStorageCredentialsRequest() (request *CreateStorageCredentialsRequest) {
    request = &CreateStorageCredentialsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vod", APIVersion, "CreateStorageCredentials")
    
    
    return
}

func NewCreateStorageCredentialsResponse() (response *CreateStorageCredentialsResponse) {
    response = &CreateStorageCredentialsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateStorageCredentials
// 用于按指定策略，生成专业版应用的临时访问凭证，比如生成用于客户端上传的临时凭证。
func (c *Client) CreateStorageCredentials(request *CreateStorageCredentialsRequest) (response *CreateStorageCredentialsResponse, err error) {
    return c.CreateStorageCredentialsWithContext(context.Background(), request)
}

// CreateStorageCredentials
// 用于按指定策略，生成专业版应用的临时访问凭证，比如生成用于客户端上传的临时凭证。
func (c *Client) CreateStorageCredentialsWithContext(ctx context.Context, request *CreateStorageCredentialsRequest) (response *CreateStorageCredentialsResponse, err error) {
    if request == nil {
        request = NewCreateStorageCredentialsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateStorageCredentials require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateStorageCredentialsResponse()
    err = c.Send(request, response)
    return
}
