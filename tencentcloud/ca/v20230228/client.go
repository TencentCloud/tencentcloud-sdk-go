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

package v20230228

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2023-02-28"

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


func NewCreateVerifyReportRequest() (request *CreateVerifyReportRequest) {
    request = &CreateVerifyReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ca", APIVersion, "CreateVerifyReport")
    
    
    return
}

func NewCreateVerifyReportResponse() (response *CreateVerifyReportResponse) {
    response = &CreateVerifyReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateVerifyReport
// 创建签名验证报告任务，此接口为异步盖章接口，盖章时效24小时。
//
// 可能返回的错误码:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateVerifyReport(request *CreateVerifyReportRequest) (response *CreateVerifyReportResponse, err error) {
    return c.CreateVerifyReportWithContext(context.Background(), request)
}

// CreateVerifyReport
// 创建签名验证报告任务，此接口为异步盖章接口，盖章时效24小时。
//
// 可能返回的错误码:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CreateVerifyReportWithContext(ctx context.Context, request *CreateVerifyReportRequest) (response *CreateVerifyReportResponse, err error) {
    if request == nil {
        request = NewCreateVerifyReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVerifyReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVerifyReportResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVerifyReportRequest() (request *DescribeVerifyReportRequest) {
    request = &DescribeVerifyReportRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ca", APIVersion, "DescribeVerifyReport")
    
    
    return
}

func NewDescribeVerifyReportResponse() (response *DescribeVerifyReportResponse) {
    response = &DescribeVerifyReportResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVerifyReport
// 下载验签报告url，url有效期默认12小时
//
// 可能返回的错误码:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeVerifyReport(request *DescribeVerifyReportRequest) (response *DescribeVerifyReportResponse, err error) {
    return c.DescribeVerifyReportWithContext(context.Background(), request)
}

// DescribeVerifyReport
// 下载验签报告url，url有效期默认12小时
//
// 可能返回的错误码:
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
func (c *Client) DescribeVerifyReportWithContext(ctx context.Context, request *DescribeVerifyReportRequest) (response *DescribeVerifyReportResponse, err error) {
    if request == nil {
        request = NewDescribeVerifyReportRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVerifyReport require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVerifyReportResponse()
    err = c.Send(request, response)
    return
}

func NewUploadFileRequest() (request *UploadFileRequest) {
    request = &UploadFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ca", APIVersion, "UploadFile")
    
    
    return
}

func NewUploadFileResponse() (response *UploadFileResponse) {
    response = &UploadFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadFile
// 文件上传接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UploadFile(request *UploadFileRequest) (response *UploadFileResponse, err error) {
    return c.UploadFileWithContext(context.Background(), request)
}

// UploadFile
// 文件上传接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) UploadFileWithContext(ctx context.Context, request *UploadFileRequest) (response *UploadFileResponse, err error) {
    if request == nil {
        request = NewUploadFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadFileResponse()
    err = c.Send(request, response)
    return
}
