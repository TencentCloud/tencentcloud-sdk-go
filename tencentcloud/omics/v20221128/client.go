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

package v20221128

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-11-28"

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


func NewDescribeRunGroupsRequest() (request *DescribeRunGroupsRequest) {
    request = &DescribeRunGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "DescribeRunGroups")
    
    
    return
}

func NewDescribeRunGroupsResponse() (response *DescribeRunGroupsResponse) {
    response = &DescribeRunGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRunGroups
// 查询任务批次列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRunGroups(request *DescribeRunGroupsRequest) (response *DescribeRunGroupsResponse, err error) {
    return c.DescribeRunGroupsWithContext(context.Background(), request)
}

// DescribeRunGroups
// 查询任务批次列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRunGroupsWithContext(ctx context.Context, request *DescribeRunGroupsRequest) (response *DescribeRunGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeRunGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRunGroups require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRunGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRunsRequest() (request *DescribeRunsRequest) {
    request = &DescribeRunsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "DescribeRuns")
    
    
    return
}

func NewDescribeRunsResponse() (response *DescribeRunsResponse) {
    response = &DescribeRunsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRuns
// 查询任务列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRuns(request *DescribeRunsRequest) (response *DescribeRunsResponse, err error) {
    return c.DescribeRunsWithContext(context.Background(), request)
}

// DescribeRuns
// 查询任务列表。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeRunsWithContext(ctx context.Context, request *DescribeRunsRequest) (response *DescribeRunsResponse, err error) {
    if request == nil {
        request = NewDescribeRunsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRuns require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRunsResponse()
    err = c.Send(request, response)
    return
}

func NewGetRunCallsRequest() (request *GetRunCallsRequest) {
    request = &GetRunCallsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "GetRunCalls")
    
    
    return
}

func NewGetRunCallsResponse() (response *GetRunCallsResponse) {
    response = &GetRunCallsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRunCalls
// 查询作业详情。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetRunCalls(request *GetRunCallsRequest) (response *GetRunCallsResponse, err error) {
    return c.GetRunCallsWithContext(context.Background(), request)
}

// GetRunCalls
// 查询作业详情。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetRunCallsWithContext(ctx context.Context, request *GetRunCallsRequest) (response *GetRunCallsResponse, err error) {
    if request == nil {
        request = NewGetRunCallsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRunCalls require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRunCallsResponse()
    err = c.Send(request, response)
    return
}

func NewGetRunStatusRequest() (request *GetRunStatusRequest) {
    request = &GetRunStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "GetRunStatus")
    
    
    return
}

func NewGetRunStatusResponse() (response *GetRunStatusResponse) {
    response = &GetRunStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetRunStatus
// 查询任务详情。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetRunStatus(request *GetRunStatusRequest) (response *GetRunStatusResponse, err error) {
    return c.GetRunStatusWithContext(context.Background(), request)
}

// GetRunStatus
// 查询任务详情。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) GetRunStatusWithContext(ctx context.Context, request *GetRunStatusRequest) (response *GetRunStatusResponse, err error) {
    if request == nil {
        request = NewGetRunStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetRunStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetRunStatusResponse()
    err = c.Send(request, response)
    return
}

func NewImportTableFileRequest() (request *ImportTableFileRequest) {
    request = &ImportTableFileRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "ImportTableFile")
    
    
    return
}

func NewImportTableFileResponse() (response *ImportTableFileResponse) {
    response = &ImportTableFileResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ImportTableFile
// 导入表格文件。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ImportTableFile(request *ImportTableFileRequest) (response *ImportTableFileResponse, err error) {
    return c.ImportTableFileWithContext(context.Background(), request)
}

// ImportTableFile
// 导入表格文件。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ImportTableFileWithContext(ctx context.Context, request *ImportTableFileRequest) (response *ImportTableFileResponse, err error) {
    if request == nil {
        request = NewImportTableFileRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImportTableFile require credential")
    }

    request.SetContext(ctx)
    
    response = NewImportTableFileResponse()
    err = c.Send(request, response)
    return
}

func NewRunApplicationRequest() (request *RunApplicationRequest) {
    request = &RunApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("omics", APIVersion, "RunApplication")
    
    
    return
}

func NewRunApplicationResponse() (response *RunApplicationResponse) {
    response = &RunApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RunApplication
// 运行应用。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RunApplication(request *RunApplicationRequest) (response *RunApplicationResponse, err error) {
    return c.RunApplicationWithContext(context.Background(), request)
}

// RunApplication
// 运行应用。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) RunApplicationWithContext(ctx context.Context, request *RunApplicationRequest) (response *RunApplicationResponse, err error) {
    if request == nil {
        request = NewRunApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RunApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewRunApplicationResponse()
    err = c.Send(request, response)
    return
}
