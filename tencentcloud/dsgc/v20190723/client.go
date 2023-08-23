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

package v20190723

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-07-23"

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


func NewCreateDSPADbMetaResourcesRequest() (request *CreateDSPADbMetaResourcesRequest) {
    request = &CreateDSPADbMetaResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("dsgc", APIVersion, "CreateDSPADbMetaResources")
    
    
    return
}

func NewCreateDSPADbMetaResourcesResponse() (response *CreateDSPADbMetaResourcesResponse) {
    response = &CreateDSPADbMetaResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateDSPADbMetaResources
// 添加用户云上数据库类型资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDSPADbMetaResources(request *CreateDSPADbMetaResourcesRequest) (response *CreateDSPADbMetaResourcesResponse, err error) {
    return c.CreateDSPADbMetaResourcesWithContext(context.Background(), request)
}

// CreateDSPADbMetaResources
// 添加用户云上数据库类型资源
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateDSPADbMetaResourcesWithContext(ctx context.Context, request *CreateDSPADbMetaResourcesRequest) (response *CreateDSPADbMetaResourcesResponse, err error) {
    if request == nil {
        request = NewCreateDSPADbMetaResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateDSPADbMetaResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateDSPADbMetaResourcesResponse()
    err = c.Send(request, response)
    return
}
