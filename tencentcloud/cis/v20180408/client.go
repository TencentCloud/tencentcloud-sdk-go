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

package v20180408

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-08"

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


func NewCreateContainerInstanceRequest() (request *CreateContainerInstanceRequest) {
    request = &CreateContainerInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cis", APIVersion, "CreateContainerInstance")
    
    
    return
}

func NewCreateContainerInstanceResponse() (response *CreateContainerInstanceResponse) {
    response = &CreateContainerInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateContainerInstance
// 此接口（CreateContainerInstance）用于创建容器实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateContainerInstance(request *CreateContainerInstanceRequest) (response *CreateContainerInstanceResponse, err error) {
    return c.CreateContainerInstanceWithContext(context.Background(), request)
}

// CreateContainerInstance
// 此接口（CreateContainerInstance）用于创建容器实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) CreateContainerInstanceWithContext(ctx context.Context, request *CreateContainerInstanceRequest) (response *CreateContainerInstanceResponse, err error) {
    if request == nil {
        request = NewCreateContainerInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateContainerInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateContainerInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteContainerInstanceRequest() (request *DeleteContainerInstanceRequest) {
    request = &DeleteContainerInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cis", APIVersion, "DeleteContainerInstance")
    
    
    return
}

func NewDeleteContainerInstanceResponse() (response *DeleteContainerInstanceResponse) {
    response = &DeleteContainerInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteContainerInstance
// 此接口（DeleteContainerInstance）用于删除容器实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteContainerInstance(request *DeleteContainerInstanceRequest) (response *DeleteContainerInstanceResponse, err error) {
    return c.DeleteContainerInstanceWithContext(context.Background(), request)
}

// DeleteContainerInstance
// 此接口（DeleteContainerInstance）用于删除容器实例
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DeleteContainerInstanceWithContext(ctx context.Context, request *DeleteContainerInstanceRequest) (response *DeleteContainerInstanceResponse, err error) {
    if request == nil {
        request = NewDeleteContainerInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteContainerInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteContainerInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerInstanceRequest() (request *DescribeContainerInstanceRequest) {
    request = &DescribeContainerInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cis", APIVersion, "DescribeContainerInstance")
    
    
    return
}

func NewDescribeContainerInstanceResponse() (response *DescribeContainerInstanceResponse) {
    response = &DescribeContainerInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeContainerInstance
// 此接口（DescribeContainerInstance）用于获取容器实例详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeContainerInstance(request *DescribeContainerInstanceRequest) (response *DescribeContainerInstanceResponse, err error) {
    return c.DescribeContainerInstanceWithContext(context.Background(), request)
}

// DescribeContainerInstance
// 此接口（DescribeContainerInstance）用于获取容器实例详情
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeContainerInstanceWithContext(ctx context.Context, request *DescribeContainerInstanceRequest) (response *DescribeContainerInstanceResponse, err error) {
    if request == nil {
        request = NewDescribeContainerInstanceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContainerInstance require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContainerInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerInstanceEventsRequest() (request *DescribeContainerInstanceEventsRequest) {
    request = &DescribeContainerInstanceEventsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cis", APIVersion, "DescribeContainerInstanceEvents")
    
    
    return
}

func NewDescribeContainerInstanceEventsResponse() (response *DescribeContainerInstanceEventsResponse) {
    response = &DescribeContainerInstanceEventsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeContainerInstanceEvents
// 此接口（DescribeContainerInstanceEvents）用于查询容器实例事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeContainerInstanceEvents(request *DescribeContainerInstanceEventsRequest) (response *DescribeContainerInstanceEventsResponse, err error) {
    return c.DescribeContainerInstanceEventsWithContext(context.Background(), request)
}

// DescribeContainerInstanceEvents
// 此接口（DescribeContainerInstanceEvents）用于查询容器实例事件列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeContainerInstanceEventsWithContext(ctx context.Context, request *DescribeContainerInstanceEventsRequest) (response *DescribeContainerInstanceEventsResponse, err error) {
    if request == nil {
        request = NewDescribeContainerInstanceEventsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContainerInstanceEvents require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContainerInstanceEventsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerInstancesRequest() (request *DescribeContainerInstancesRequest) {
    request = &DescribeContainerInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cis", APIVersion, "DescribeContainerInstances")
    
    
    return
}

func NewDescribeContainerInstancesResponse() (response *DescribeContainerInstancesResponse) {
    response = &DescribeContainerInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeContainerInstances
// 此接口（DescribeContainerInstances）查询容器实例列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeContainerInstances(request *DescribeContainerInstancesRequest) (response *DescribeContainerInstancesResponse, err error) {
    return c.DescribeContainerInstancesWithContext(context.Background(), request)
}

// DescribeContainerInstances
// 此接口（DescribeContainerInstances）查询容器实例列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeContainerInstancesWithContext(ctx context.Context, request *DescribeContainerInstancesRequest) (response *DescribeContainerInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeContainerInstancesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContainerInstances require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContainerInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeContainerLogRequest() (request *DescribeContainerLogRequest) {
    request = &DescribeContainerLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cis", APIVersion, "DescribeContainerLog")
    
    
    return
}

func NewDescribeContainerLogResponse() (response *DescribeContainerLogResponse) {
    response = &DescribeContainerLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeContainerLog
// 此接口（DescribeContainerLog）用于获取容器日志信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeContainerLog(request *DescribeContainerLogRequest) (response *DescribeContainerLogResponse, err error) {
    return c.DescribeContainerLogWithContext(context.Background(), request)
}

// DescribeContainerLog
// 此接口（DescribeContainerLog）用于获取容器日志信息
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeContainerLogWithContext(ctx context.Context, request *DescribeContainerLogRequest) (response *DescribeContainerLogResponse, err error) {
    if request == nil {
        request = NewDescribeContainerLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeContainerLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeContainerLogResponse()
    err = c.Send(request, response)
    return
}

func NewInquiryPriceCreateCisRequest() (request *InquiryPriceCreateCisRequest) {
    request = &InquiryPriceCreateCisRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("cis", APIVersion, "InquiryPriceCreateCis")
    
    
    return
}

func NewInquiryPriceCreateCisResponse() (response *InquiryPriceCreateCisResponse) {
    response = &InquiryPriceCreateCisResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InquiryPriceCreateCis
// 此接口（InquiryPriceCreateCis）用于查询容器实例价格
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) InquiryPriceCreateCis(request *InquiryPriceCreateCisRequest) (response *InquiryPriceCreateCisResponse, err error) {
    return c.InquiryPriceCreateCisWithContext(context.Background(), request)
}

// InquiryPriceCreateCis
// 此接口（InquiryPriceCreateCis）用于查询容器实例价格
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) InquiryPriceCreateCisWithContext(ctx context.Context, request *InquiryPriceCreateCisRequest) (response *InquiryPriceCreateCisResponse, err error) {
    if request == nil {
        request = NewInquiryPriceCreateCisRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InquiryPriceCreateCis require credential")
    }

    request.SetContext(ctx)
    
    response = NewInquiryPriceCreateCisResponse()
    err = c.Send(request, response)
    return
}
