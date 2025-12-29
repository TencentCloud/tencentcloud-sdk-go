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

package v20251030

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2025-10-30"

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


func NewCloudMateAgentRequest() (request *CloudMateAgentRequest) {
    request = &CloudMateAgentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cloudmate", APIVersion, "CloudMateAgent")
    
    
    return
}

func NewCloudMateAgentResponse() (response *CloudMateAgentResponse) {
    response = &CloudMateAgentResponse{} 
    return

}

// CloudMateAgent
// 发起智能诊断对话
//
// 注意：使用该API时，请务必设置请求域名（Endpoint）为 cloudmate.ai.tencentcloudapi.com
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKSPACENOTEXIST = "ResourceNotFound.WorkspaceNotExist"
func (c *Client) CloudMateAgent(request *CloudMateAgentRequest) (response *CloudMateAgentResponse, err error) {
    return c.CloudMateAgentWithContext(context.Background(), request)
}

// CloudMateAgent
// 发起智能诊断对话
//
// 注意：使用该API时，请务必设置请求域名（Endpoint）为 cloudmate.ai.tencentcloudapi.com
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_WORKSPACENOTEXIST = "ResourceNotFound.WorkspaceNotExist"
func (c *Client) CloudMateAgentWithContext(ctx context.Context, request *CloudMateAgentRequest) (response *CloudMateAgentResponse, err error) {
    if request == nil {
        request = NewCloudMateAgentRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "cloudmate", APIVersion, "CloudMateAgent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloudMateAgent require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloudMateAgentResponse()
    err = c.Send(request, response)
    return
}
