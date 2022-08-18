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

package v20200710

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-07-10"

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
    
    request.Init().WithApiInfo("mmps", APIVersion, "CreateAppScanTask")
    
    
    return
}

func NewCreateAppScanTaskResponse() (response *CreateAppScanTaskResponse) {
    response = &CreateAppScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAppScanTask
// 创建小程序隐私合规诊断任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateAppScanTask(request *CreateAppScanTaskRequest) (response *CreateAppScanTaskResponse, err error) {
    return c.CreateAppScanTaskWithContext(context.Background(), request)
}

// CreateAppScanTask
// 创建小程序隐私合规诊断任务
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
    
    request.Init().WithApiInfo("mmps", APIVersion, "CreateAppScanTaskRepeat")
    
    
    return
}

func NewCreateAppScanTaskRepeatResponse() (response *CreateAppScanTaskRepeatResponse) {
    response = &CreateAppScanTaskRepeatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAppScanTaskRepeat
// 小程序隐私合规诊断重试任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateAppScanTaskRepeat(request *CreateAppScanTaskRepeatRequest) (response *CreateAppScanTaskRepeatResponse, err error) {
    return c.CreateAppScanTaskRepeatWithContext(context.Background(), request)
}

// CreateAppScanTaskRepeat
// 小程序隐私合规诊断重试任务
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

func NewCreateFlySecMiniAppProfessionalScanTaskRequest() (request *CreateFlySecMiniAppProfessionalScanTaskRequest) {
    request = &CreateFlySecMiniAppProfessionalScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mmps", APIVersion, "CreateFlySecMiniAppProfessionalScanTask")
    
    
    return
}

func NewCreateFlySecMiniAppProfessionalScanTaskResponse() (response *CreateFlySecMiniAppProfessionalScanTaskResponse) {
    response = &CreateFlySecMiniAppProfessionalScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFlySecMiniAppProfessionalScanTask
// 创建小程序安全深度诊断任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateFlySecMiniAppProfessionalScanTask(request *CreateFlySecMiniAppProfessionalScanTaskRequest) (response *CreateFlySecMiniAppProfessionalScanTaskResponse, err error) {
    return c.CreateFlySecMiniAppProfessionalScanTaskWithContext(context.Background(), request)
}

// CreateFlySecMiniAppProfessionalScanTask
// 创建小程序安全深度诊断任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateFlySecMiniAppProfessionalScanTaskWithContext(ctx context.Context, request *CreateFlySecMiniAppProfessionalScanTaskRequest) (response *CreateFlySecMiniAppProfessionalScanTaskResponse, err error) {
    if request == nil {
        request = NewCreateFlySecMiniAppProfessionalScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlySecMiniAppProfessionalScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlySecMiniAppProfessionalScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlySecMiniAppScanTaskRequest() (request *CreateFlySecMiniAppScanTaskRequest) {
    request = &CreateFlySecMiniAppScanTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mmps", APIVersion, "CreateFlySecMiniAppScanTask")
    
    
    return
}

func NewCreateFlySecMiniAppScanTaskResponse() (response *CreateFlySecMiniAppScanTaskResponse) {
    response = &CreateFlySecMiniAppScanTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFlySecMiniAppScanTask
// 创建小程序翼扬安全的基础或深度诊断任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateFlySecMiniAppScanTask(request *CreateFlySecMiniAppScanTaskRequest) (response *CreateFlySecMiniAppScanTaskResponse, err error) {
    return c.CreateFlySecMiniAppScanTaskWithContext(context.Background(), request)
}

// CreateFlySecMiniAppScanTask
// 创建小程序翼扬安全的基础或深度诊断任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateFlySecMiniAppScanTaskWithContext(ctx context.Context, request *CreateFlySecMiniAppScanTaskRequest) (response *CreateFlySecMiniAppScanTaskResponse, err error) {
    if request == nil {
        request = NewCreateFlySecMiniAppScanTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlySecMiniAppScanTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlySecMiniAppScanTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFlySecMiniAppScanTaskRepeatRequest() (request *CreateFlySecMiniAppScanTaskRepeatRequest) {
    request = &CreateFlySecMiniAppScanTaskRepeatRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mmps", APIVersion, "CreateFlySecMiniAppScanTaskRepeat")
    
    
    return
}

func NewCreateFlySecMiniAppScanTaskRepeatResponse() (response *CreateFlySecMiniAppScanTaskRepeatResponse) {
    response = &CreateFlySecMiniAppScanTaskRepeatResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateFlySecMiniAppScanTaskRepeat
// 重新提交基础诊断任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateFlySecMiniAppScanTaskRepeat(request *CreateFlySecMiniAppScanTaskRepeatRequest) (response *CreateFlySecMiniAppScanTaskRepeatResponse, err error) {
    return c.CreateFlySecMiniAppScanTaskRepeatWithContext(context.Background(), request)
}

// CreateFlySecMiniAppScanTaskRepeat
// 重新提交基础诊断任务
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) CreateFlySecMiniAppScanTaskRepeatWithContext(ctx context.Context, request *CreateFlySecMiniAppScanTaskRepeatRequest) (response *CreateFlySecMiniAppScanTaskRepeatResponse, err error) {
    if request == nil {
        request = NewCreateFlySecMiniAppScanTaskRepeatRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateFlySecMiniAppScanTaskRepeat require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateFlySecMiniAppScanTaskRepeatResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBasicDiagnosisResourceUsageInfoRequest() (request *DescribeBasicDiagnosisResourceUsageInfoRequest) {
    request = &DescribeBasicDiagnosisResourceUsageInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mmps", APIVersion, "DescribeBasicDiagnosisResourceUsageInfo")
    
    
    return
}

func NewDescribeBasicDiagnosisResourceUsageInfoResponse() (response *DescribeBasicDiagnosisResourceUsageInfoResponse) {
    response = &DescribeBasicDiagnosisResourceUsageInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeBasicDiagnosisResourceUsageInfo
// 查询翼扬安全基础诊断资源使用情况
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeBasicDiagnosisResourceUsageInfo(request *DescribeBasicDiagnosisResourceUsageInfoRequest) (response *DescribeBasicDiagnosisResourceUsageInfoResponse, err error) {
    return c.DescribeBasicDiagnosisResourceUsageInfoWithContext(context.Background(), request)
}

// DescribeBasicDiagnosisResourceUsageInfo
// 查询翼扬安全基础诊断资源使用情况
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeBasicDiagnosisResourceUsageInfoWithContext(ctx context.Context, request *DescribeBasicDiagnosisResourceUsageInfoRequest) (response *DescribeBasicDiagnosisResourceUsageInfoResponse, err error) {
    if request == nil {
        request = NewDescribeBasicDiagnosisResourceUsageInfoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBasicDiagnosisResourceUsageInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBasicDiagnosisResourceUsageInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlySecMiniAppReportUrlRequest() (request *DescribeFlySecMiniAppReportUrlRequest) {
    request = &DescribeFlySecMiniAppReportUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mmps", APIVersion, "DescribeFlySecMiniAppReportUrl")
    
    
    return
}

func NewDescribeFlySecMiniAppReportUrlResponse() (response *DescribeFlySecMiniAppReportUrlResponse) {
    response = &DescribeFlySecMiniAppReportUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlySecMiniAppReportUrl
// 获取翼扬诊断任务报告链接地址
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFlySecMiniAppReportUrl(request *DescribeFlySecMiniAppReportUrlRequest) (response *DescribeFlySecMiniAppReportUrlResponse, err error) {
    return c.DescribeFlySecMiniAppReportUrlWithContext(context.Background(), request)
}

// DescribeFlySecMiniAppReportUrl
// 获取翼扬诊断任务报告链接地址
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFlySecMiniAppReportUrlWithContext(ctx context.Context, request *DescribeFlySecMiniAppReportUrlRequest) (response *DescribeFlySecMiniAppReportUrlResponse, err error) {
    if request == nil {
        request = NewDescribeFlySecMiniAppReportUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlySecMiniAppReportUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlySecMiniAppReportUrlResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlySecMiniAppScanReportListRequest() (request *DescribeFlySecMiniAppScanReportListRequest) {
    request = &DescribeFlySecMiniAppScanReportListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mmps", APIVersion, "DescribeFlySecMiniAppScanReportList")
    
    
    return
}

func NewDescribeFlySecMiniAppScanReportListResponse() (response *DescribeFlySecMiniAppScanReportListResponse) {
    response = &DescribeFlySecMiniAppScanReportListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlySecMiniAppScanReportList
// 查询指定小程序版本的翼扬诊断安全得分
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFlySecMiniAppScanReportList(request *DescribeFlySecMiniAppScanReportListRequest) (response *DescribeFlySecMiniAppScanReportListResponse, err error) {
    return c.DescribeFlySecMiniAppScanReportListWithContext(context.Background(), request)
}

// DescribeFlySecMiniAppScanReportList
// 查询指定小程序版本的翼扬诊断安全得分
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFlySecMiniAppScanReportListWithContext(ctx context.Context, request *DescribeFlySecMiniAppScanReportListRequest) (response *DescribeFlySecMiniAppScanReportListResponse, err error) {
    if request == nil {
        request = NewDescribeFlySecMiniAppScanReportListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlySecMiniAppScanReportList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlySecMiniAppScanReportListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlySecMiniAppScanTaskListRequest() (request *DescribeFlySecMiniAppScanTaskListRequest) {
    request = &DescribeFlySecMiniAppScanTaskListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mmps", APIVersion, "DescribeFlySecMiniAppScanTaskList")
    
    
    return
}

func NewDescribeFlySecMiniAppScanTaskListResponse() (response *DescribeFlySecMiniAppScanTaskListResponse) {
    response = &DescribeFlySecMiniAppScanTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlySecMiniAppScanTaskList
// 获取翼扬安全诊断任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFlySecMiniAppScanTaskList(request *DescribeFlySecMiniAppScanTaskListRequest) (response *DescribeFlySecMiniAppScanTaskListResponse, err error) {
    return c.DescribeFlySecMiniAppScanTaskListWithContext(context.Background(), request)
}

// DescribeFlySecMiniAppScanTaskList
// 获取翼扬安全诊断任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFlySecMiniAppScanTaskListWithContext(ctx context.Context, request *DescribeFlySecMiniAppScanTaskListRequest) (response *DescribeFlySecMiniAppScanTaskListResponse, err error) {
    if request == nil {
        request = NewDescribeFlySecMiniAppScanTaskListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlySecMiniAppScanTaskList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlySecMiniAppScanTaskListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlySecMiniAppScanTaskParamRequest() (request *DescribeFlySecMiniAppScanTaskParamRequest) {
    request = &DescribeFlySecMiniAppScanTaskParamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mmps", APIVersion, "DescribeFlySecMiniAppScanTaskParam")
    
    
    return
}

func NewDescribeFlySecMiniAppScanTaskParamResponse() (response *DescribeFlySecMiniAppScanTaskParamResponse) {
    response = &DescribeFlySecMiniAppScanTaskParamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlySecMiniAppScanTaskParam
// 获取用户提交的基础诊断任务参数信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFlySecMiniAppScanTaskParam(request *DescribeFlySecMiniAppScanTaskParamRequest) (response *DescribeFlySecMiniAppScanTaskParamResponse, err error) {
    return c.DescribeFlySecMiniAppScanTaskParamWithContext(context.Background(), request)
}

// DescribeFlySecMiniAppScanTaskParam
// 获取用户提交的基础诊断任务参数信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFlySecMiniAppScanTaskParamWithContext(ctx context.Context, request *DescribeFlySecMiniAppScanTaskParamRequest) (response *DescribeFlySecMiniAppScanTaskParamResponse, err error) {
    if request == nil {
        request = NewDescribeFlySecMiniAppScanTaskParamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlySecMiniAppScanTaskParam require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlySecMiniAppScanTaskParamResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFlySecMiniAppScanTaskStatusRequest() (request *DescribeFlySecMiniAppScanTaskStatusRequest) {
    request = &DescribeFlySecMiniAppScanTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mmps", APIVersion, "DescribeFlySecMiniAppScanTaskStatus")
    
    
    return
}

func NewDescribeFlySecMiniAppScanTaskStatusResponse() (response *DescribeFlySecMiniAppScanTaskStatusResponse) {
    response = &DescribeFlySecMiniAppScanTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFlySecMiniAppScanTaskStatus
// 查询翼扬安全诊断任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFlySecMiniAppScanTaskStatus(request *DescribeFlySecMiniAppScanTaskStatusRequest) (response *DescribeFlySecMiniAppScanTaskStatusResponse, err error) {
    return c.DescribeFlySecMiniAppScanTaskStatusWithContext(context.Background(), request)
}

// DescribeFlySecMiniAppScanTaskStatus
// 查询翼扬安全诊断任务状态
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeFlySecMiniAppScanTaskStatusWithContext(ctx context.Context, request *DescribeFlySecMiniAppScanTaskStatusRequest) (response *DescribeFlySecMiniAppScanTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeFlySecMiniAppScanTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeFlySecMiniAppScanTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeFlySecMiniAppScanTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceUsageInfoRequest() (request *DescribeResourceUsageInfoRequest) {
    request = &DescribeResourceUsageInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mmps", APIVersion, "DescribeResourceUsageInfo")
    
    
    return
}

func NewDescribeResourceUsageInfoResponse() (response *DescribeResourceUsageInfoResponse) {
    response = &DescribeResourceUsageInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeResourceUsageInfo
// 查询翼扬安全资源使用情况
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeResourceUsageInfo(request *DescribeResourceUsageInfoRequest) (response *DescribeResourceUsageInfoResponse, err error) {
    return c.DescribeResourceUsageInfoWithContext(context.Background(), request)
}

// DescribeResourceUsageInfo
// 查询翼扬安全资源使用情况
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
    
    request.Init().WithApiInfo("mmps", APIVersion, "DescribeScanTaskList")
    
    
    return
}

func NewDescribeScanTaskListResponse() (response *DescribeScanTaskListResponse) {
    response = &DescribeScanTaskListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScanTaskList
// 获取小程序隐私合规诊断任务列表
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeScanTaskList(request *DescribeScanTaskListRequest) (response *DescribeScanTaskListResponse, err error) {
    return c.DescribeScanTaskListWithContext(context.Background(), request)
}

// DescribeScanTaskList
// 获取小程序隐私合规诊断任务列表
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
    
    request.Init().WithApiInfo("mmps", APIVersion, "DescribeScanTaskReportUrl")
    
    
    return
}

func NewDescribeScanTaskReportUrlResponse() (response *DescribeScanTaskReportUrlResponse) {
    response = &DescribeScanTaskReportUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScanTaskReportUrl
// 获取小程序合规诊断任务报告url
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
func (c *Client) DescribeScanTaskReportUrl(request *DescribeScanTaskReportUrlRequest) (response *DescribeScanTaskReportUrlResponse, err error) {
    return c.DescribeScanTaskReportUrlWithContext(context.Background(), request)
}

// DescribeScanTaskReportUrl
// 获取小程序合规诊断任务报告url
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
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
    
    request.Init().WithApiInfo("mmps", APIVersion, "DescribeScanTaskStatus")
    
    
    return
}

func NewDescribeScanTaskStatusResponse() (response *DescribeScanTaskStatusResponse) {
    response = &DescribeScanTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScanTaskStatus
// 查询小程序隐私合规诊断任务状态
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScanTaskStatus(request *DescribeScanTaskStatusRequest) (response *DescribeScanTaskStatusResponse, err error) {
    return c.DescribeScanTaskStatusWithContext(context.Background(), request)
}

// DescribeScanTaskStatus
// 查询小程序隐私合规诊断任务状态
//
// 可能返回的错误码:
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
