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

package v20230303

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-03-03"

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


func NewCreate5GInstanceRequest() (request *Create5GInstanceRequest) {
    request = &Create5GInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csxg", APIVersion, "Create5GInstance")
    
    
    return
}

func NewCreate5GInstanceResponse() (response *Create5GInstanceResponse) {
    response = &Create5GInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// Create5GInstance
// 创建5G入云服务接口
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) Create5GInstance(request *Create5GInstanceRequest) (response *Create5GInstanceResponse, err error) {
    return c.Create5GInstanceWithContext(context.Background(), request)
}

// Create5GInstance
// 创建5G入云服务接口
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) Create5GInstanceWithContext(ctx context.Context, request *Create5GInstanceRequest) (response *Create5GInstanceResponse, err error) {
    if request == nil {
        request = NewCreate5GInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("Create5GInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreate5GInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDelete5GInstanceRequest() (request *Delete5GInstanceRequest) {
    request = &Delete5GInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csxg", APIVersion, "Delete5GInstance")
    
    
    return
}

func NewDelete5GInstanceResponse() (response *Delete5GInstanceResponse) {
    response = &Delete5GInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// Delete5GInstance
// 删除5G入云服务
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) Delete5GInstance(request *Delete5GInstanceRequest) (response *Delete5GInstanceResponse, err error) {
    return c.Delete5GInstanceWithContext(context.Background(), request)
}

// Delete5GInstance
// 删除5G入云服务
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) Delete5GInstanceWithContext(ctx context.Context, request *Delete5GInstanceRequest) (response *Delete5GInstanceResponse, err error) {
    if request == nil {
        request = NewDelete5GInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("Delete5GInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDelete5GInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribe5GAPNsRequest() (request *Describe5GAPNsRequest) {
    request = &Describe5GAPNsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csxg", APIVersion, "Describe5GAPNs")
    
    
    return
}

func NewDescribe5GAPNsResponse() (response *Describe5GAPNsResponse) {
    response = &Describe5GAPNsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// Describe5GAPNs
// 查询APN信息
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) Describe5GAPNs(request *Describe5GAPNsRequest) (response *Describe5GAPNsResponse, err error) {
    return c.Describe5GAPNsWithContext(context.Background(), request)
}

// Describe5GAPNs
// 查询APN信息
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) Describe5GAPNsWithContext(ctx context.Context, request *Describe5GAPNsRequest) (response *Describe5GAPNsResponse, err error) {
    if request == nil {
        request = NewDescribe5GAPNsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("Describe5GAPNs require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribe5GAPNsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribe5GInstancesRequest() (request *Describe5GInstancesRequest) {
    request = &Describe5GInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csxg", APIVersion, "Describe5GInstances")
    
    
    return
}

func NewDescribe5GInstancesResponse() (response *Describe5GInstancesResponse) {
    response = &Describe5GInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// Describe5GInstances
// 查询5G入云服务
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) Describe5GInstances(request *Describe5GInstancesRequest) (response *Describe5GInstancesResponse, err error) {
    return c.Describe5GInstancesWithContext(context.Background(), request)
}

// Describe5GInstances
// 查询5G入云服务
//
// 可能返回的错误码:
//  LIMITEXCEEDED = "LimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) Describe5GInstancesWithContext(ctx context.Context, request *Describe5GInstancesRequest) (response *Describe5GInstancesResponse, err error) {
    if request == nil {
        request = NewDescribe5GInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("Describe5GInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribe5GInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewModify5GInstanceAttributeRequest() (request *Modify5GInstanceAttributeRequest) {
    request = &Modify5GInstanceAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("csxg", APIVersion, "Modify5GInstanceAttribute")
    
    
    return
}

func NewModify5GInstanceAttributeResponse() (response *Modify5GInstanceAttributeResponse) {
    response = &Modify5GInstanceAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// Modify5GInstanceAttribute
// 修改5G入云服务
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) Modify5GInstanceAttribute(request *Modify5GInstanceAttributeRequest) (response *Modify5GInstanceAttributeResponse, err error) {
    return c.Modify5GInstanceAttributeWithContext(context.Background(), request)
}

// Modify5GInstanceAttribute
// 修改5G入云服务
//
// 可能返回的错误码:
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) Modify5GInstanceAttributeWithContext(ctx context.Context, request *Modify5GInstanceAttributeRequest) (response *Modify5GInstanceAttributeResponse, err error) {
    if request == nil {
        request = NewModify5GInstanceAttributeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("Modify5GInstanceAttribute require credential")
    }

    request.SetContext(ctx)
    
    response = NewModify5GInstanceAttributeResponse()
    err = c.Send(request, response)
    return
}
