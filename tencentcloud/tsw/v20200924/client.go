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

package v20200924

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-09-24"

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


func NewDescribeAgentShellRequest() (request *DescribeAgentShellRequest) {
    request = &DescribeAgentShellRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tsw", APIVersion, "DescribeAgentShell")
    
    
    return
}

func NewDescribeAgentShellResponse() (response *DescribeAgentShellResponse) {
    response = &DescribeAgentShellResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAgentShell
// 获取服务接入信息
func (c *Client) DescribeAgentShell(request *DescribeAgentShellRequest) (response *DescribeAgentShellResponse, err error) {
    return c.DescribeAgentShellWithContext(context.Background(), request)
}

// DescribeAgentShell
// 获取服务接入信息
func (c *Client) DescribeAgentShellWithContext(ctx context.Context, request *DescribeAgentShellRequest) (response *DescribeAgentShellResponse, err error) {
    if request == nil {
        request = NewDescribeAgentShellRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAgentShell require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAgentShellResponse()
    err = c.Send(request, response)
    return
}
