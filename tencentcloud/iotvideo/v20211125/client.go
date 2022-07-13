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

package v20211125

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-11-25"

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


func NewCallDeviceActionAsyncRequest() (request *CallDeviceActionAsyncRequest) {
    request = &CallDeviceActionAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CallDeviceActionAsync")
    
    
    return
}

func NewCallDeviceActionAsyncResponse() (response *CallDeviceActionAsyncResponse) {
    response = &CallDeviceActionAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CallDeviceActionAsync
// 异步调用设备行为
func (c *Client) CallDeviceActionAsync(request *CallDeviceActionAsyncRequest) (response *CallDeviceActionAsyncResponse, err error) {
    return c.CallDeviceActionAsyncWithContext(context.Background(), request)
}

// CallDeviceActionAsync
// 异步调用设备行为
func (c *Client) CallDeviceActionAsyncWithContext(ctx context.Context, request *CallDeviceActionAsyncRequest) (response *CallDeviceActionAsyncResponse, err error) {
    if request == nil {
        request = NewCallDeviceActionAsyncRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CallDeviceActionAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewCallDeviceActionAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewCallDeviceActionSyncRequest() (request *CallDeviceActionSyncRequest) {
    request = &CallDeviceActionSyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CallDeviceActionSync")
    
    
    return
}

func NewCallDeviceActionSyncResponse() (response *CallDeviceActionSyncResponse) {
    response = &CallDeviceActionSyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CallDeviceActionSync
// 同步调用设备行为
func (c *Client) CallDeviceActionSync(request *CallDeviceActionSyncRequest) (response *CallDeviceActionSyncResponse, err error) {
    return c.CallDeviceActionSyncWithContext(context.Background(), request)
}

// CallDeviceActionSync
// 同步调用设备行为
func (c *Client) CallDeviceActionSyncWithContext(ctx context.Context, request *CallDeviceActionSyncRequest) (response *CallDeviceActionSyncResponse, err error) {
    if request == nil {
        request = NewCallDeviceActionSyncRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CallDeviceActionSync require credential")
    }

    request.SetContext(ctx)
    
    response = NewCallDeviceActionSyncResponse()
    err = c.Send(request, response)
    return
}

func NewCreateProductRequest() (request *CreateProductRequest) {
    request = &CreateProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "CreateProduct")
    
    
    return
}

func NewCreateProductResponse() (response *CreateProductResponse) {
    response = &CreateProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateProduct
// 创建产品
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateProduct(request *CreateProductRequest) (response *CreateProductResponse, err error) {
    return c.CreateProductWithContext(context.Background(), request)
}

// CreateProduct
// 创建产品
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateProductWithContext(ctx context.Context, request *CreateProductRequest) (response *CreateProductResponse, err error) {
    if request == nil {
        request = NewCreateProductRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateProduct require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateProductResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDeviceDataStatsRequest() (request *DescribeDeviceDataStatsRequest) {
    request = &DescribeDeviceDataStatsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeDeviceDataStats")
    
    
    return
}

func NewDescribeDeviceDataStatsResponse() (response *DescribeDeviceDataStatsResponse) {
    response = &DescribeDeviceDataStatsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeDeviceDataStats
// 查询设备数据统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceDataStats(request *DescribeDeviceDataStatsRequest) (response *DescribeDeviceDataStatsResponse, err error) {
    return c.DescribeDeviceDataStatsWithContext(context.Background(), request)
}

// DescribeDeviceDataStats
// 查询设备数据统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeDeviceDataStatsWithContext(ctx context.Context, request *DescribeDeviceDataStatsRequest) (response *DescribeDeviceDataStatsResponse, err error) {
    if request == nil {
        request = NewDescribeDeviceDataStatsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeDeviceDataStats require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeDeviceDataStatsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMessageDataStatsRequest() (request *DescribeMessageDataStatsRequest) {
    request = &DescribeMessageDataStatsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "DescribeMessageDataStats")
    
    
    return
}

func NewDescribeMessageDataStatsResponse() (response *DescribeMessageDataStatsResponse) {
    response = &DescribeMessageDataStatsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeMessageDataStats
// 查询设备消息数量统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMessageDataStats(request *DescribeMessageDataStatsRequest) (response *DescribeMessageDataStatsResponse, err error) {
    return c.DescribeMessageDataStatsWithContext(context.Background(), request)
}

// DescribeMessageDataStats
// 查询设备消息数量统计
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeMessageDataStatsWithContext(ctx context.Context, request *DescribeMessageDataStatsRequest) (response *DescribeMessageDataStatsResponse, err error) {
    if request == nil {
        request = NewDescribeMessageDataStatsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMessageDataStats require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMessageDataStatsResponse()
    err = c.Send(request, response)
    return
}

func NewGenSingleDeviceSignatureOfPublicRequest() (request *GenSingleDeviceSignatureOfPublicRequest) {
    request = &GenSingleDeviceSignatureOfPublicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("iotvideo", APIVersion, "GenSingleDeviceSignatureOfPublic")
    
    
    return
}

func NewGenSingleDeviceSignatureOfPublicResponse() (response *GenSingleDeviceSignatureOfPublicResponse) {
    response = &GenSingleDeviceSignatureOfPublicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GenSingleDeviceSignatureOfPublic
// 获取设备的绑定签名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GenSingleDeviceSignatureOfPublic(request *GenSingleDeviceSignatureOfPublicRequest) (response *GenSingleDeviceSignatureOfPublicResponse, err error) {
    return c.GenSingleDeviceSignatureOfPublicWithContext(context.Background(), request)
}

// GenSingleDeviceSignatureOfPublic
// 获取设备的绑定签名
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) GenSingleDeviceSignatureOfPublicWithContext(ctx context.Context, request *GenSingleDeviceSignatureOfPublicRequest) (response *GenSingleDeviceSignatureOfPublicResponse, err error) {
    if request == nil {
        request = NewGenSingleDeviceSignatureOfPublicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenSingleDeviceSignatureOfPublic require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenSingleDeviceSignatureOfPublicResponse()
    err = c.Send(request, response)
    return
}
