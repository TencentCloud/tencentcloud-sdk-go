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

package v20211014

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-10-14"

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


func NewGetIndustryV1HomeMembersRequest() (request *GetIndustryV1HomeMembersRequest) {
    request = &GetIndustryV1HomeMembersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("icr", APIVersion, "GetIndustryV1HomeMembers")
    
    
    return
}

func NewGetIndustryV1HomeMembersResponse() (response *GetIndustryV1HomeMembersResponse) {
    response = &GetIndustryV1HomeMembersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetIndustryV1HomeMembers
// 获取成员列表接口
func (c *Client) GetIndustryV1HomeMembers(request *GetIndustryV1HomeMembersRequest) (response *GetIndustryV1HomeMembersResponse, err error) {
    return c.GetIndustryV1HomeMembersWithContext(context.Background(), request)
}

// GetIndustryV1HomeMembers
// 获取成员列表接口
func (c *Client) GetIndustryV1HomeMembersWithContext(ctx context.Context, request *GetIndustryV1HomeMembersRequest) (response *GetIndustryV1HomeMembersResponse, err error) {
    if request == nil {
        request = NewGetIndustryV1HomeMembersRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetIndustryV1HomeMembers require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetIndustryV1HomeMembersResponse()
    err = c.Send(request, response)
    return
}
