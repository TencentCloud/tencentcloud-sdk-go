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

package v20250217

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2025-02-17"

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


func NewCompleteApprovalRequest() (request *CompleteApprovalRequest) {
    request = &CompleteApprovalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("evt", APIVersion, "CompleteApproval")
    
    
    return
}

func NewCompleteApprovalResponse() (response *CompleteApprovalResponse) {
    response = &CompleteApprovalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CompleteApproval
// 执行审批
func (c *Client) CompleteApproval(request *CompleteApprovalRequest) (response *CompleteApprovalResponse, err error) {
    return c.CompleteApprovalWithContext(context.Background(), request)
}

// CompleteApproval
// 执行审批
func (c *Client) CompleteApprovalWithContext(ctx context.Context, request *CompleteApprovalRequest) (response *CompleteApprovalResponse, err error) {
    if request == nil {
        request = NewCompleteApprovalRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "evt", APIVersion, "CompleteApproval")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CompleteApproval require credential")
    }

    request.SetContext(ctx)
    
    response = NewCompleteApprovalResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRoleUserRequest() (request *CreateRoleUserRequest) {
    request = &CreateRoleUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("evt", APIVersion, "CreateRoleUser")
    
    
    return
}

func NewCreateRoleUserResponse() (response *CreateRoleUserResponse) {
    response = &CreateRoleUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateRoleUser
// 创建人员
func (c *Client) CreateRoleUser(request *CreateRoleUserRequest) (response *CreateRoleUserResponse, err error) {
    return c.CreateRoleUserWithContext(context.Background(), request)
}

// CreateRoleUser
// 创建人员
func (c *Client) CreateRoleUserWithContext(ctx context.Context, request *CreateRoleUserRequest) (response *CreateRoleUserResponse, err error) {
    if request == nil {
        request = NewCreateRoleUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "evt", APIVersion, "CreateRoleUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRoleUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRoleUserResponse()
    err = c.Send(request, response)
    return
}
