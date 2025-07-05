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

package v20201203

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-12-03"

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


func NewSubmitTaskEventRequest() (request *SubmitTaskEventRequest) {
    request = &SubmitTaskEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("smop", APIVersion, "SubmitTaskEvent")
    
    
    return
}

func NewSubmitTaskEventResponse() (response *SubmitTaskEventResponse) {
    response = &SubmitTaskEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitTaskEvent
// smop产品下线，接口也一起下线
//
// 
//
// 提交任务事件接口
func (c *Client) SubmitTaskEvent(request *SubmitTaskEventRequest) (response *SubmitTaskEventResponse, err error) {
    return c.SubmitTaskEventWithContext(context.Background(), request)
}

// SubmitTaskEvent
// smop产品下线，接口也一起下线
//
// 
//
// 提交任务事件接口
func (c *Client) SubmitTaskEventWithContext(ctx context.Context, request *SubmitTaskEventRequest) (response *SubmitTaskEventResponse, err error) {
    if request == nil {
        request = NewSubmitTaskEventRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitTaskEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitTaskEventResponse()
    err = c.Send(request, response)
    return
}
