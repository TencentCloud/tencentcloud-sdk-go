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

package v20210331

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-03-31"

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


func NewBindOrganizationMemberAuthAccountRequest() (request *BindOrganizationMemberAuthAccountRequest) {
    request = &BindOrganizationMemberAuthAccountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("organization", APIVersion, "BindOrganizationMemberAuthAccount")
    return
}

func NewBindOrganizationMemberAuthAccountResponse() (response *BindOrganizationMemberAuthAccountResponse) {
    response = &BindOrganizationMemberAuthAccountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BindOrganizationMemberAuthAccount
// 绑定组织成员和子账号的授权关系
//
// 可能返回的错误码:
//  FAILEDOPERATION_OPERATEPOLICY = "FailedOperation.OperatePolicy"
//  FAILEDOPERATION_SUBACCOUNTNOTEXIST = "FailedOperation.SubAccountNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) BindOrganizationMemberAuthAccount(request *BindOrganizationMemberAuthAccountRequest) (response *BindOrganizationMemberAuthAccountResponse, err error) {
    if request == nil {
        request = NewBindOrganizationMemberAuthAccountRequest()
    }
    response = NewBindOrganizationMemberAuthAccountResponse()
    err = c.Send(request, response)
    return
}
