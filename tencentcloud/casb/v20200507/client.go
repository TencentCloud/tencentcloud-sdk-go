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

package v20200507

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-05-07"

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


func NewCopyCryptoColumnPolicyRequest() (request *CopyCryptoColumnPolicyRequest) {
    request = &CopyCryptoColumnPolicyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("casb", APIVersion, "CopyCryptoColumnPolicy")
    
    
    return
}

func NewCopyCryptoColumnPolicyResponse() (response *CopyCryptoColumnPolicyResponse) {
    response = &CopyCryptoColumnPolicyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CopyCryptoColumnPolicy
// 同region下 根据用户输入的CasbId,MetaDataId 复制策略到DstCasbId,MetaDataId中。
//
// 场景1：
//
// 相同CasbId，不同MetadataId 下策略复制
//
// 场景2：
//
// 不同Casbid,不同MetaDataId 下策略复制
//
// 场景3:
//
// 相同CasbId,相同MetaDataId 且 DatabaseName不同 策略复制
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CopyCryptoColumnPolicy(request *CopyCryptoColumnPolicyRequest) (response *CopyCryptoColumnPolicyResponse, err error) {
    return c.CopyCryptoColumnPolicyWithContext(context.Background(), request)
}

// CopyCryptoColumnPolicy
// 同region下 根据用户输入的CasbId,MetaDataId 复制策略到DstCasbId,MetaDataId中。
//
// 场景1：
//
// 相同CasbId，不同MetadataId 下策略复制
//
// 场景2：
//
// 不同Casbid,不同MetaDataId 下策略复制
//
// 场景3:
//
// 相同CasbId,相同MetaDataId 且 DatabaseName不同 策略复制
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CopyCryptoColumnPolicyWithContext(ctx context.Context, request *CopyCryptoColumnPolicyRequest) (response *CopyCryptoColumnPolicyResponse, err error) {
    if request == nil {
        request = NewCopyCryptoColumnPolicyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CopyCryptoColumnPolicy require credential")
    }

    request.SetContext(ctx)
    
    response = NewCopyCryptoColumnPolicyResponse()
    err = c.Send(request, response)
    return
}
