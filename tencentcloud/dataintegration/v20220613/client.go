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

package v20220613

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-06-13"

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


func NewSendMessageRequest() (request *SendMessageRequest) {
    request = &SendMessageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dataintegration", APIVersion, "SendMessage")
    
    
    return
}

func NewSendMessageResponse() (response *SendMessageResponse) {
    response = &SendMessageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SendMessage
// 使用SDK将数据上报到DIP
func (c *Client) SendMessage(request *SendMessageRequest) (response *SendMessageResponse, err error) {
    return c.SendMessageWithContext(context.Background(), request)
}

// SendMessage
// 使用SDK将数据上报到DIP
func (c *Client) SendMessageWithContext(ctx context.Context, request *SendMessageRequest) (response *SendMessageResponse, err error) {
    if request == nil {
        request = NewSendMessageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SendMessage require credential")
    }

    request.SetContext(ctx)
    
    response = NewSendMessageResponse()
    err = c.Send(request, response)
    return
}
