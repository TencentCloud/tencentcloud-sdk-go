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

package v20220105

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-01-05"

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


func NewCreateAppScanTaskRequest() (request *CreateAppScanTaskRequest) {
    request = &CreateAppScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("acp", APIVersion, "CreateAppScanTask")
    
    
    return
}

func NewCreateAppScanTaskResponse() (response *CreateAppScanTaskResponse) {
    response = &CreateAppScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAppScanTask
// 创建应用合规隐私诊断任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateAppScanTask(request *CreateAppScanTaskRequest) (response *CreateAppScanTaskResponse, err error) {
    return c.CreateAppScanTaskWithContext(context.Background(), request)
}

// CreateAppScanTask
// 创建应用合规隐私诊断任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateAppScanTaskWithContext(ctx context.Context, request *CreateAppScanTaskRequest) (response *CreateAppScanTaskResponse, err error) {
    if request == nil {
        request = NewCreateAppScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAppScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAppScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAppScanTaskRepeatRequest() (request *CreateAppScanTaskRepeatRequest) {
    request = &CreateAppScanTaskRepeatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("acp", APIVersion, "CreateAppScanTaskRepeat")
    
    
    return
}

func NewCreateAppScanTaskRepeatResponse() (response *CreateAppScanTaskRepeatResponse) {
    response = &CreateAppScanTaskRepeatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAppScanTaskRepeat
// App应用合规隐私诊断重试任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateAppScanTaskRepeat(request *CreateAppScanTaskRepeatRequest) (response *CreateAppScanTaskRepeatResponse, err error) {
    return c.CreateAppScanTaskRepeatWithContext(context.Background(), request)
}

// CreateAppScanTaskRepeat
// App应用合规隐私诊断重试任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateAppScanTaskRepeatWithContext(ctx context.Context, request *CreateAppScanTaskRepeatRequest) (response *CreateAppScanTaskRepeatResponse, err error) {
    if request == nil {
        request = NewCreateAppScanTaskRepeatRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAppScanTaskRepeat require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAppScanTaskRepeatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChannelTaskReportUrlRequest() (request *DescribeChannelTaskReportUrlRequest) {
    request = &DescribeChannelTaskReportUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("acp", APIVersion, "DescribeChannelTaskReportUrl")
    
    
    return
}

func NewDescribeChannelTaskReportUrlResponse() (response *DescribeChannelTaskReportUrlResponse) {
    response = &DescribeChannelTaskReportUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeChannelTaskReportUrl
// 获取子渠道的App合规诊断任务报告url
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeChannelTaskReportUrl(request *DescribeChannelTaskReportUrlRequest) (response *DescribeChannelTaskReportUrlResponse, err error) {
    return c.DescribeChannelTaskReportUrlWithContext(context.Background(), request)
}

// DescribeChannelTaskReportUrl
// 获取子渠道的App合规诊断任务报告url
//
// 可能返回的错误码:
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeChannelTaskReportUrlWithContext(ctx context.Context, request *DescribeChannelTaskReportUrlRequest) (response *DescribeChannelTaskReportUrlResponse, err error) {
    if request == nil {
        request = NewDescribeChannelTaskReportUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeChannelTaskReportUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeChannelTaskReportUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileTicketRequest() (request *DescribeFileTicketRequest) {
    request = &DescribeFileTicketRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("acp", APIVersion, "DescribeFileTicket")
    
    
    return
}

func NewDescribeFileTicketResponse() (response *DescribeFileTicketResponse) {
    response = &DescribeFileTicketResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFileTicket
// 获取应用合规文件上传凭证，用于上传诊断文件
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFileTicket(request *DescribeFileTicketRequest) (response *DescribeFileTicketResponse, err error) {
    return c.DescribeFileTicketWithContext(context.Background(), request)
}

// DescribeFileTicket
// 获取应用合规文件上传凭证，用于上传诊断文件
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFileTicketWithContext(ctx context.Context, request *DescribeFileTicketRequest) (response *DescribeFileTicketResponse, err error) {
    if request == nil {
        request = NewDescribeFileTicketRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFileTicket require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFileTicketResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceUsageInfoRequest() (request *DescribeResourceUsageInfoRequest) {
    request = &DescribeResourceUsageInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("acp", APIVersion, "DescribeResourceUsageInfo")
    
    
    return
}

func NewDescribeResourceUsageInfoResponse() (response *DescribeResourceUsageInfoResponse) {
    response = &DescribeResourceUsageInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResourceUsageInfo
// 查询应用合规平台用户资源的使用情况
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeResourceUsageInfo(request *DescribeResourceUsageInfoRequest) (response *DescribeResourceUsageInfoResponse, err error) {
    return c.DescribeResourceUsageInfoWithContext(context.Background(), request)
}

// DescribeResourceUsageInfo
// 查询应用合规平台用户资源的使用情况
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeResourceUsageInfoWithContext(ctx context.Context, request *DescribeResourceUsageInfoRequest) (response *DescribeResourceUsageInfoResponse, err error) {
    if request == nil {
        request = NewDescribeResourceUsageInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeResourceUsageInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeResourceUsageInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanTaskListRequest() (request *DescribeScanTaskListRequest) {
    request = &DescribeScanTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("acp", APIVersion, "DescribeScanTaskList")
    
    
    return
}

func NewDescribeScanTaskListResponse() (response *DescribeScanTaskListResponse) {
    response = &DescribeScanTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScanTaskList
// 获取App隐私合规诊断任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeScanTaskList(request *DescribeScanTaskListRequest) (response *DescribeScanTaskListResponse, err error) {
    return c.DescribeScanTaskListWithContext(context.Background(), request)
}

// DescribeScanTaskList
// 获取App隐私合规诊断任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeScanTaskListWithContext(ctx context.Context, request *DescribeScanTaskListRequest) (response *DescribeScanTaskListResponse, err error) {
    if request == nil {
        request = NewDescribeScanTaskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanTaskReportUrlRequest() (request *DescribeScanTaskReportUrlRequest) {
    request = &DescribeScanTaskReportUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("acp", APIVersion, "DescribeScanTaskReportUrl")
    
    
    return
}

func NewDescribeScanTaskReportUrlResponse() (response *DescribeScanTaskReportUrlResponse) {
    response = &DescribeScanTaskReportUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScanTaskReportUrl
// 获取App合规诊断任务报告url
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScanTaskReportUrl(request *DescribeScanTaskReportUrlRequest) (response *DescribeScanTaskReportUrlResponse, err error) {
    return c.DescribeScanTaskReportUrlWithContext(context.Background(), request)
}

// DescribeScanTaskReportUrl
// 获取App合规诊断任务报告url
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScanTaskReportUrlWithContext(ctx context.Context, request *DescribeScanTaskReportUrlRequest) (response *DescribeScanTaskReportUrlResponse, err error) {
    if request == nil {
        request = NewDescribeScanTaskReportUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanTaskReportUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanTaskReportUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanTaskStatusRequest() (request *DescribeScanTaskStatusRequest) {
    request = &DescribeScanTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("acp", APIVersion, "DescribeScanTaskStatus")
    
    
    return
}

func NewDescribeScanTaskStatusResponse() (response *DescribeScanTaskStatusResponse) {
    response = &DescribeScanTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScanTaskStatus
// 查询App隐私合规诊断任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScanTaskStatus(request *DescribeScanTaskStatusRequest) (response *DescribeScanTaskStatusResponse, err error) {
    return c.DescribeScanTaskStatusWithContext(context.Background(), request)
}

// DescribeScanTaskStatus
// 查询App隐私合规诊断任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScanTaskStatusWithContext(ctx context.Context, request *DescribeScanTaskStatusRequest) (response *DescribeScanTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeScanTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScanTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScanTaskStatusResponse()
    err = c.Send(request, response)
    return
}
