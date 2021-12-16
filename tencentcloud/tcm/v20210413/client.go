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

package v20210413

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-04-13"

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


func NewDescribeMeshRequest() (request *DescribeMeshRequest) {
    request = &DescribeMeshRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcm", APIVersion, "DescribeMesh")
    
    
    return
}

func NewDescribeMeshResponse() (response *DescribeMeshResponse) {
    response = &DescribeMeshResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMesh
// 查询网格详情
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeMesh(request *DescribeMeshRequest) (response *DescribeMeshResponse, err error) {
    if request == nil {
        request = NewDescribeMeshRequest()
    }
    
    response = NewDescribeMeshResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMeshListRequest() (request *DescribeMeshListRequest) {
    request = &DescribeMeshListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcm", APIVersion, "DescribeMeshList")
    
    
    return
}

func NewDescribeMeshListResponse() (response *DescribeMeshListResponse) {
    response = &DescribeMeshListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMeshList
// 查询网格列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeMeshList(request *DescribeMeshListRequest) (response *DescribeMeshListResponse, err error) {
    if request == nil {
        request = NewDescribeMeshListRequest()
    }
    
    response = NewDescribeMeshListResponse()
    err = c.Send(request, response)
    return
}
