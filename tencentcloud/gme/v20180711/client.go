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

package v20180711

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-07-11"

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


func NewCreateAgeDetectTaskRequest() (request *CreateAgeDetectTaskRequest) {
    request = &CreateAgeDetectTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "CreateAgeDetectTask")
    
    
    return
}

func NewCreateAgeDetectTaskResponse() (response *CreateAgeDetectTaskResponse) {
    response = &CreateAgeDetectTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAgeDetectTask
// 用于创建年龄语音识别任务的接口，请求频率10次/秒。该接口目前通过白名单开放试用，如有需求，请提交工单申请。
//
// </br>
//
// <h4><b>接口功能说明：</b></h4>
//
// <li>支持对语音文件进行检测，判断是否为未成年人。</li>
//
// <li>支持批量提交检测子任务。检测子任务列表最多支持100个。</li>
//
// </br>
//
// <h4><b>音频文件限制说明：</b></h4>
//
// <li>音频文件大小限制：10 M</li>
//
// <li>音频文件时长限制：3分钟</li>
//
// <li>音频文件格式支持的类型：.wav、.m4a、.amr、.mp3、.aac、.wma、.ogg</li>
//
// </br>
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAgeDetectTask(request *CreateAgeDetectTaskRequest) (response *CreateAgeDetectTaskResponse, err error) {
    if request == nil {
        request = NewCreateAgeDetectTaskRequest()
    }
    
    response = NewCreateAgeDetectTaskResponse()
    err = c.Send(request, response)
    return
}

// CreateAgeDetectTask
// 用于创建年龄语音识别任务的接口，请求频率10次/秒。该接口目前通过白名单开放试用，如有需求，请提交工单申请。
//
// </br>
//
// <h4><b>接口功能说明：</b></h4>
//
// <li>支持对语音文件进行检测，判断是否为未成年人。</li>
//
// <li>支持批量提交检测子任务。检测子任务列表最多支持100个。</li>
//
// </br>
//
// <h4><b>音频文件限制说明：</b></h4>
//
// <li>音频文件大小限制：10 M</li>
//
// <li>音频文件时长限制：3分钟</li>
//
// <li>音频文件格式支持的类型：.wav、.m4a、.amr、.mp3、.aac、.wma、.ogg</li>
//
// </br>
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) CreateAgeDetectTaskWithContext(ctx context.Context, request *CreateAgeDetectTaskRequest) (response *CreateAgeDetectTaskResponse, err error) {
    if request == nil {
        request = NewCreateAgeDetectTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateAgeDetectTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAppRequest() (request *CreateAppRequest) {
    request = &CreateAppRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "CreateApp")
    
    
    return
}

func NewCreateAppResponse() (response *CreateAppResponse) {
    response = &CreateAppResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateApp
// 本接口(CreateApp)用于创建一个GME应用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGKEY = "InvalidParameter.TagKey"
//  LIMITEXCEEDED_APPLICATION = "LimitExceeded.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CREATEAPPDENIED = "UnauthorizedOperation.CreateAppDenied"
//  UNAUTHORIZEDOPERATION_UNREALNAMEAUTH = "UnauthorizedOperation.UnRealNameAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateApp(request *CreateAppRequest) (response *CreateAppResponse, err error) {
    if request == nil {
        request = NewCreateAppRequest()
    }
    
    response = NewCreateAppResponse()
    err = c.Send(request, response)
    return
}

// CreateApp
// 本接口(CreateApp)用于创建一个GME应用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGKEY = "InvalidParameter.TagKey"
//  LIMITEXCEEDED_APPLICATION = "LimitExceeded.Application"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CREATEAPPDENIED = "UnauthorizedOperation.CreateAppDenied"
//  UNAUTHORIZEDOPERATION_UNREALNAMEAUTH = "UnauthorizedOperation.UnRealNameAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAppWithContext(ctx context.Context, request *CreateAppRequest) (response *CreateAppResponse, err error) {
    if request == nil {
        request = NewCreateAppRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateAppResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAgeDetectTaskRequest() (request *DescribeAgeDetectTaskRequest) {
    request = &DescribeAgeDetectTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "DescribeAgeDetectTask")
    
    
    return
}

func NewDescribeAgeDetectTaskResponse() (response *DescribeAgeDetectTaskResponse) {
    response = &DescribeAgeDetectTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAgeDetectTask
// 查询年龄语音识别任务结果，请求频率10次/秒。该接口目前通过白名单开放试用，如有需求，请提交工单申请。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAgeDetectTask(request *DescribeAgeDetectTaskRequest) (response *DescribeAgeDetectTaskResponse, err error) {
    if request == nil {
        request = NewDescribeAgeDetectTaskRequest()
    }
    
    response = NewDescribeAgeDetectTaskResponse()
    err = c.Send(request, response)
    return
}

// DescribeAgeDetectTask
// 查询年龄语音识别任务结果，请求频率10次/秒。该接口目前通过白名单开放试用，如有需求，请提交工单申请。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeAgeDetectTaskWithContext(ctx context.Context, request *DescribeAgeDetectTaskRequest) (response *DescribeAgeDetectTaskResponse, err error) {
    if request == nil {
        request = NewDescribeAgeDetectTaskRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAgeDetectTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAppStatisticsRequest() (request *DescribeAppStatisticsRequest) {
    request = &DescribeAppStatisticsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "DescribeAppStatistics")
    
    
    return
}

func NewDescribeAppStatisticsResponse() (response *DescribeAppStatisticsResponse) {
    response = &DescribeAppStatisticsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAppStatistics
// 本接口(DescribeAppStatistics)用于获取某个GME应用的用量数据。包括实时语音，语音消息及转文本，语音分析等。最长查询周期为最近30天。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATEINVALID = "InvalidParameter.DateInvalid"
//  INVALIDPARAMETER_TIMERANGEERROR = "InvalidParameter.TimeRangeError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZIDISNOTFOUND = "ResourceNotFound.BizidIsNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAppStatistics(request *DescribeAppStatisticsRequest) (response *DescribeAppStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeAppStatisticsRequest()
    }
    
    response = NewDescribeAppStatisticsResponse()
    err = c.Send(request, response)
    return
}

// DescribeAppStatistics
// 本接口(DescribeAppStatistics)用于获取某个GME应用的用量数据。包括实时语音，语音消息及转文本，语音分析等。最长查询周期为最近30天。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_DATEINVALID = "InvalidParameter.DateInvalid"
//  INVALIDPARAMETER_TIMERANGEERROR = "InvalidParameter.TimeRangeError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZIDISNOTFOUND = "ResourceNotFound.BizidIsNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAppStatisticsWithContext(ctx context.Context, request *DescribeAppStatisticsRequest) (response *DescribeAppStatisticsResponse, err error) {
    if request == nil {
        request = NewDescribeAppStatisticsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeAppStatisticsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationDataRequest() (request *DescribeApplicationDataRequest) {
    request = &DescribeApplicationDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "DescribeApplicationData")
    
    
    return
}

func NewDescribeApplicationDataResponse() (response *DescribeApplicationDataResponse) {
    response = &DescribeApplicationDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeApplicationData
// 本接口(DescribeApplicationData)用于获取数据详情信息，最多可拉取最近90天的数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApplicationData(request *DescribeApplicationDataRequest) (response *DescribeApplicationDataResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationDataRequest()
    }
    
    response = NewDescribeApplicationDataResponse()
    err = c.Send(request, response)
    return
}

// DescribeApplicationData
// 本接口(DescribeApplicationData)用于获取数据详情信息，最多可拉取最近90天的数据。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeApplicationDataWithContext(ctx context.Context, request *DescribeApplicationDataRequest) (response *DescribeApplicationDataResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationDataRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeApplicationDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFilterResultRequest() (request *DescribeFilterResultRequest) {
    request = &DescribeFilterResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "DescribeFilterResult")
    
    
    return
}

func NewDescribeFilterResultResponse() (response *DescribeFilterResultResponse) {
    response = &DescribeFilterResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFilterResult
// 根据应用ID和文件ID查询识别结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFilterResult(request *DescribeFilterResultRequest) (response *DescribeFilterResultResponse, err error) {
    if request == nil {
        request = NewDescribeFilterResultRequest()
    }
    
    response = NewDescribeFilterResultResponse()
    err = c.Send(request, response)
    return
}

// DescribeFilterResult
// 根据应用ID和文件ID查询识别结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFilterResultWithContext(ctx context.Context, request *DescribeFilterResultRequest) (response *DescribeFilterResultResponse, err error) {
    if request == nil {
        request = NewDescribeFilterResultRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeFilterResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFilterResultListRequest() (request *DescribeFilterResultListRequest) {
    request = &DescribeFilterResultListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "DescribeFilterResultList")
    
    
    return
}

func NewDescribeFilterResultListResponse() (response *DescribeFilterResultListResponse) {
    response = &DescribeFilterResultListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeFilterResultList
// 根据日期查询识别结果列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFilterResultList(request *DescribeFilterResultListRequest) (response *DescribeFilterResultListResponse, err error) {
    if request == nil {
        request = NewDescribeFilterResultListRequest()
    }
    
    response = NewDescribeFilterResultListResponse()
    err = c.Send(request, response)
    return
}

// DescribeFilterResultList
// 根据日期查询识别结果列表
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeFilterResultListWithContext(ctx context.Context, request *DescribeFilterResultListRequest) (response *DescribeFilterResultListResponse, err error) {
    if request == nil {
        request = NewDescribeFilterResultListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeFilterResultListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRealtimeScanConfigRequest() (request *DescribeRealtimeScanConfigRequest) {
    request = &DescribeRealtimeScanConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "DescribeRealtimeScanConfig")
    
    
    return
}

func NewDescribeRealtimeScanConfigResponse() (response *DescribeRealtimeScanConfigResponse) {
    response = &DescribeRealtimeScanConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRealtimeScanConfig
// 获取用户自定义送检信息
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeRealtimeScanConfig(request *DescribeRealtimeScanConfigRequest) (response *DescribeRealtimeScanConfigResponse, err error) {
    if request == nil {
        request = NewDescribeRealtimeScanConfigRequest()
    }
    
    response = NewDescribeRealtimeScanConfigResponse()
    err = c.Send(request, response)
    return
}

// DescribeRealtimeScanConfig
// 获取用户自定义送检信息
//
// 可能返回的错误码:
//  OPERATIONDENIED = "OperationDenied"
func (c *Client) DescribeRealtimeScanConfigWithContext(ctx context.Context, request *DescribeRealtimeScanConfigRequest) (response *DescribeRealtimeScanConfigResponse, err error) {
    if request == nil {
        request = NewDescribeRealtimeScanConfigRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeRealtimeScanConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoomInfoRequest() (request *DescribeRoomInfoRequest) {
    request = &DescribeRoomInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "DescribeRoomInfo")
    
    
    return
}

func NewDescribeRoomInfoResponse() (response *DescribeRoomInfoResponse) {
    response = &DescribeRoomInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeRoomInfo
// 获取房间内用户信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRoomInfo(request *DescribeRoomInfoRequest) (response *DescribeRoomInfoResponse, err error) {
    if request == nil {
        request = NewDescribeRoomInfoRequest()
    }
    
    response = NewDescribeRoomInfoResponse()
    err = c.Send(request, response)
    return
}

// DescribeRoomInfo
// 获取房间内用户信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) DescribeRoomInfoWithContext(ctx context.Context, request *DescribeRoomInfoRequest) (response *DescribeRoomInfoResponse, err error) {
    if request == nil {
        request = NewDescribeRoomInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeRoomInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScanResultListRequest() (request *DescribeScanResultListRequest) {
    request = &DescribeScanResultListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "DescribeScanResultList")
    
    
    return
}

func NewDescribeScanResultListResponse() (response *DescribeScanResultListResponse) {
    response = &DescribeScanResultListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeScanResultList
// 本接口(DescribeScanResultList)用于查询语音检测结果，查询任务列表最多支持100个。
//
// <p style="color:red">如果在提交语音检测任务时未设置 Callback 字段，则需要通过本接口获取检测结果</p>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGKEY = "InvalidParameter.TagKey"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScanResultList(request *DescribeScanResultListRequest) (response *DescribeScanResultListResponse, err error) {
    if request == nil {
        request = NewDescribeScanResultListRequest()
    }
    
    response = NewDescribeScanResultListResponse()
    err = c.Send(request, response)
    return
}

// DescribeScanResultList
// 本接口(DescribeScanResultList)用于查询语音检测结果，查询任务列表最多支持100个。
//
// <p style="color:red">如果在提交语音检测任务时未设置 Callback 字段，则需要通过本接口获取检测结果</p>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TAGKEY = "InvalidParameter.TagKey"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeScanResultListWithContext(ctx context.Context, request *DescribeScanResultListRequest) (response *DescribeScanResultListResponse, err error) {
    if request == nil {
        request = NewDescribeScanResultListRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeScanResultListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserInAndOutTimeRequest() (request *DescribeUserInAndOutTimeRequest) {
    request = &DescribeUserInAndOutTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "DescribeUserInAndOutTime")
    
    
    return
}

func NewDescribeUserInAndOutTimeResponse() (response *DescribeUserInAndOutTimeResponse) {
    response = &DescribeUserInAndOutTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeUserInAndOutTime
// 拉取用户在房间得进出时间
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CALLBACKADDRESS = "InvalidParameter.CallbackAddress"
//  INVALIDPARAMETER_DATEINVALID = "InvalidParameter.DateInvalid"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CREATEAPPDENIED = "UnauthorizedOperation.CreateAppDenied"
//  UNAUTHORIZEDOPERATION_UNREALNAMEAUTH = "UnauthorizedOperation.UnRealNameAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserInAndOutTime(request *DescribeUserInAndOutTimeRequest) (response *DescribeUserInAndOutTimeResponse, err error) {
    if request == nil {
        request = NewDescribeUserInAndOutTimeRequest()
    }
    
    response = NewDescribeUserInAndOutTimeResponse()
    err = c.Send(request, response)
    return
}

// DescribeUserInAndOutTime
// 拉取用户在房间得进出时间
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_LOGINFAILED = "FailedOperation.LoginFailed"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CALLBACKADDRESS = "InvalidParameter.CallbackAddress"
//  INVALIDPARAMETER_DATEINVALID = "InvalidParameter.DateInvalid"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_TASKNOTFOUND = "ResourceNotFound.TaskNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_CREATEAPPDENIED = "UnauthorizedOperation.CreateAppDenied"
//  UNAUTHORIZEDOPERATION_UNREALNAMEAUTH = "UnauthorizedOperation.UnRealNameAuth"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeUserInAndOutTimeWithContext(ctx context.Context, request *DescribeUserInAndOutTimeRequest) (response *DescribeUserInAndOutTimeResponse, err error) {
    if request == nil {
        request = NewDescribeUserInAndOutTimeRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeUserInAndOutTimeResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAppStatusRequest() (request *ModifyAppStatusRequest) {
    request = &ModifyAppStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "ModifyAppStatus")
    
    
    return
}

func NewModifyAppStatusResponse() (response *ModifyAppStatusResponse) {
    response = &ModifyAppStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyAppStatus
// 本接口(ModifyAppStatus)用于修改应用总开关状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZIDISNOTFOUND = "ResourceNotFound.BizidIsNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAppStatus(request *ModifyAppStatusRequest) (response *ModifyAppStatusResponse, err error) {
    if request == nil {
        request = NewModifyAppStatusRequest()
    }
    
    response = NewModifyAppStatusResponse()
    err = c.Send(request, response)
    return
}

// ModifyAppStatus
// 本接口(ModifyAppStatus)用于修改应用总开关状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_USERFEENEGATIVE = "FailedOperation.UserFeeNegative"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCENOTFOUND_BIZIDISNOTFOUND = "ResourceNotFound.BizidIsNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ModifyAppStatusWithContext(ctx context.Context, request *ModifyAppStatusRequest) (response *ModifyAppStatusResponse, err error) {
    if request == nil {
        request = NewModifyAppStatusRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyAppStatusResponse()
    err = c.Send(request, response)
    return
}

func NewModifyRoomInfoRequest() (request *ModifyRoomInfoRequest) {
    request = &ModifyRoomInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "ModifyRoomInfo")
    
    
    return
}

func NewModifyRoomInfoResponse() (response *ModifyRoomInfoResponse) {
    response = &ModifyRoomInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyRoomInfo
// 修改房间信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyRoomInfo(request *ModifyRoomInfoRequest) (response *ModifyRoomInfoResponse, err error) {
    if request == nil {
        request = NewModifyRoomInfoRequest()
    }
    
    response = NewModifyRoomInfoResponse()
    err = c.Send(request, response)
    return
}

// ModifyRoomInfo
// 修改房间信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) ModifyRoomInfoWithContext(ctx context.Context, request *ModifyRoomInfoRequest) (response *ModifyRoomInfoResponse, err error) {
    if request == nil {
        request = NewModifyRoomInfoRequest()
    }
    request.SetContext(ctx)
    
    response = NewModifyRoomInfoResponse()
    err = c.Send(request, response)
    return
}

func NewScanVoiceRequest() (request *ScanVoiceRequest) {
    request = &ScanVoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "ScanVoice")
    
    
    return
}

func NewScanVoiceResponse() (response *ScanVoiceResponse) {
    response = &ScanVoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ScanVoice
// 本接口(ScanVoice)用于提交语音检测任务，检测任务列表最多支持100个。使用前请您登录[控制台 - 服务配置](https://console.cloud.tencent.com/gamegme/conf)开启语音分析服务。
//
// </br></br>
//
// 
//
// <h4><b>功能试用说明：</b></h4>
//
// <li>打开前往<a href="https://console.cloud.tencent.com/gamegme/tryout">控制台 - 产品试用</a>免费试用语音分析服务。</li>
//
// </br>
//
// 
//
// <h4><b>接口功能说明：</b></h4>
//
// <li>支持对语音流或语音文件进行检测，判断其中是否包含违规内容。</li>
//
// <li>支持设置回调地址 Callback 获取检测结果，同时支持通过接口(查询语音检测结果)主动轮询获取检测结果。</li>
//
// <li>支持场景输入，包括：谩骂、色情等场景</li>
//
// <li>支持批量提交检测任务。检测任务列表最多支持100个。</li>
//
// </br>
//
// <h4><b>音频文件限制说明：</b></h4>
//
// <li>音频文件大小限制：100 M</li>
//
// <li>音频文件时长限制：30分钟</li>
//
// <li>音频文件格式支持的类型：.wav、.m4a、.amr、.mp3、.aac、.wma、.ogg</li>
//
// </br>
//
// <h4><b>语音流限制说明：</b></h4>
//
// <li>语音流格式支持的类型：.m3u8、.flv</li>
//
// <li>语音流支持的传输协议：RTMP、HTTP、HTTPS</li>
//
// <li>语音流时长限制：4小时</li>
//
// <li>支持音视频流分离并对音频流进行分析</li>
//
// </br>
//
// <h4 id="Label_Value"><b>Scenes 与 Label 参数说明：</b></h4>
//
// <p>提交语音检测任务时，需要指定 Scenes 场景参数，<font color="red">目前要求您设置 Scenes 参数值为：["default"]</font>；而在检测结果中，则包含请求时指定的场景，以及对应类型的检测结果。</p>
//
// <table>
//
// <thread>
//
// <tr>
//
// <th>场景</th>
//
// <th>描述</th>
//
// <th>Label</th>
//
// </tr>
//
// </thread>
//
// <tbody>
//
// <tr>
//
// <td>语音检测</td>
//
// <td>语音检测的检测类型</td>
//
// <td>
//
// <p>normal:正常文本</p>
//
// <p>porn:色情</p>
//
// <p>abuse:谩骂</p>
//
// <p>ad :广告</p>
//
// <p>contraband :违禁</p>
//
// <p>customized:自定义词库。目前白名单开放，如有需要请<a href="https://cloud.tencent.com/apply/p/8809fjcik56">联系我们</a>。</p>
//
// </td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// </br>
//
// <h4 id="Callback_Declare"><b>回调相关说明：</b></h4>
//
// <li>如果在请求参数中指定了回调地址参数 Callback，即一个 HTTP(S) 协议接口的 URL，则需要支持 POST 方法，传输数据编码采用 UTF-8。</li>
//
// <li>在推送回调数据后，接收到的 HTTP 状态码为 200 时，表示推送成功。</li>
//
// <li>HTTP 头参数说明：</li>
//
// <table>
//
// <thread>
//
// <tr>
//
// <th>名称</th>
//
// <th>类型</th>
//
// <th>是否必需</th>
//
// <th>描述</th>
//
// </tr>
//
// </thread>
//
// <tbody>
//
// <tr>
//
// <td>Signatue</td>
//
// <td>string</td>
//
// <td>是</td>
//
// <td>签名，具体见<a href="#Callback_Signatue">签名生成说明</a></td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// <ul  id="Callback_Signatue">
//
// 	<li>签名生成说明：</li>
//
// 	<ul>
//
// 		<li>使用 HMAC-SH1 算法, 最终结果做 BASE64 编码;</li>
//
// 		<li>签名原文串为 POST+body 的整个json内容(长度以 Content-Length 为准);</li>
//
// 		<li>签名key为应用的 SecretKey，可以通过控制台查看。</li>
//
// 	</ul>
//
// </ul>
//
// 
//
// <li>回调示例如下<font color="red">（详细字段说明见结构：
//
// <a href="https://cloud.tencent.com/document/api/607/35375#DescribeScanResult" target="_blank">DescribeScanResult</a>）</font>：</li>
//
// <pre><code>{
//
// 	"Code": 0,
//
// 	"DataId": "1400000000_test_data_id",
//
// 	"ScanFinishTime": 1566720906,
//
// 	"HitFlag": true,
//
// 	"Live": false,
//
// 	"Msg": "",
//
// 	"ScanPiece": [{
//
// 		"DumpUrl": "",
//
// 		"HitFlag": true,
//
// 		"MainType": "abuse",
//
// 		"RoomId": "123",
//
// 		"OpenId": "xxx",
//
// 		"Info":"",
//
// 		"Offset": 0,
//
// 		"Duration": 3400,
//
// 		"PieceStartTime":1574684231,
//
// 		"ScanDetail": [{
//
// 			"EndTime": 1110,
//
// 			"KeyWord": "xxx",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 1110
//
// 		}, {
//
// 			"EndTime": 1380,
//
// 			"KeyWord": "xxx",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 930
//
// 		}, {
//
// 			"EndTime": 1560,
//
// 			"KeyWord": "xxx",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 930
//
// 		}, {
//
// 			"EndTime": 2820,
//
// 			"KeyWord": "xxx",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 2490
//
// 		}]
//
// 	}],
//
// 	"ScanStartTime": 1566720905,
//
// 	"Scenes": [
//
// 		"default"
//
// 	],
//
// 	"Status": "Success",
//
// 	"TaskId": "xxx",
//
// 	"Url": "https://xxx/xxx.m4a"
//
// }
//
// </code></pre>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CALLBACKADDRESS = "InvalidParameter.CallbackAddress"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ScanVoice(request *ScanVoiceRequest) (response *ScanVoiceResponse, err error) {
    if request == nil {
        request = NewScanVoiceRequest()
    }
    
    response = NewScanVoiceResponse()
    err = c.Send(request, response)
    return
}

// ScanVoice
// 本接口(ScanVoice)用于提交语音检测任务，检测任务列表最多支持100个。使用前请您登录[控制台 - 服务配置](https://console.cloud.tencent.com/gamegme/conf)开启语音分析服务。
//
// </br></br>
//
// 
//
// <h4><b>功能试用说明：</b></h4>
//
// <li>打开前往<a href="https://console.cloud.tencent.com/gamegme/tryout">控制台 - 产品试用</a>免费试用语音分析服务。</li>
//
// </br>
//
// 
//
// <h4><b>接口功能说明：</b></h4>
//
// <li>支持对语音流或语音文件进行检测，判断其中是否包含违规内容。</li>
//
// <li>支持设置回调地址 Callback 获取检测结果，同时支持通过接口(查询语音检测结果)主动轮询获取检测结果。</li>
//
// <li>支持场景输入，包括：谩骂、色情等场景</li>
//
// <li>支持批量提交检测任务。检测任务列表最多支持100个。</li>
//
// </br>
//
// <h4><b>音频文件限制说明：</b></h4>
//
// <li>音频文件大小限制：100 M</li>
//
// <li>音频文件时长限制：30分钟</li>
//
// <li>音频文件格式支持的类型：.wav、.m4a、.amr、.mp3、.aac、.wma、.ogg</li>
//
// </br>
//
// <h4><b>语音流限制说明：</b></h4>
//
// <li>语音流格式支持的类型：.m3u8、.flv</li>
//
// <li>语音流支持的传输协议：RTMP、HTTP、HTTPS</li>
//
// <li>语音流时长限制：4小时</li>
//
// <li>支持音视频流分离并对音频流进行分析</li>
//
// </br>
//
// <h4 id="Label_Value"><b>Scenes 与 Label 参数说明：</b></h4>
//
// <p>提交语音检测任务时，需要指定 Scenes 场景参数，<font color="red">目前要求您设置 Scenes 参数值为：["default"]</font>；而在检测结果中，则包含请求时指定的场景，以及对应类型的检测结果。</p>
//
// <table>
//
// <thread>
//
// <tr>
//
// <th>场景</th>
//
// <th>描述</th>
//
// <th>Label</th>
//
// </tr>
//
// </thread>
//
// <tbody>
//
// <tr>
//
// <td>语音检测</td>
//
// <td>语音检测的检测类型</td>
//
// <td>
//
// <p>normal:正常文本</p>
//
// <p>porn:色情</p>
//
// <p>abuse:谩骂</p>
//
// <p>ad :广告</p>
//
// <p>contraband :违禁</p>
//
// <p>customized:自定义词库。目前白名单开放，如有需要请<a href="https://cloud.tencent.com/apply/p/8809fjcik56">联系我们</a>。</p>
//
// </td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// </br>
//
// <h4 id="Callback_Declare"><b>回调相关说明：</b></h4>
//
// <li>如果在请求参数中指定了回调地址参数 Callback，即一个 HTTP(S) 协议接口的 URL，则需要支持 POST 方法，传输数据编码采用 UTF-8。</li>
//
// <li>在推送回调数据后，接收到的 HTTP 状态码为 200 时，表示推送成功。</li>
//
// <li>HTTP 头参数说明：</li>
//
// <table>
//
// <thread>
//
// <tr>
//
// <th>名称</th>
//
// <th>类型</th>
//
// <th>是否必需</th>
//
// <th>描述</th>
//
// </tr>
//
// </thread>
//
// <tbody>
//
// <tr>
//
// <td>Signatue</td>
//
// <td>string</td>
//
// <td>是</td>
//
// <td>签名，具体见<a href="#Callback_Signatue">签名生成说明</a></td>
//
// </tr>
//
// </tbody>
//
// </table>
//
// <ul  id="Callback_Signatue">
//
// 	<li>签名生成说明：</li>
//
// 	<ul>
//
// 		<li>使用 HMAC-SH1 算法, 最终结果做 BASE64 编码;</li>
//
// 		<li>签名原文串为 POST+body 的整个json内容(长度以 Content-Length 为准);</li>
//
// 		<li>签名key为应用的 SecretKey，可以通过控制台查看。</li>
//
// 	</ul>
//
// </ul>
//
// 
//
// <li>回调示例如下<font color="red">（详细字段说明见结构：
//
// <a href="https://cloud.tencent.com/document/api/607/35375#DescribeScanResult" target="_blank">DescribeScanResult</a>）</font>：</li>
//
// <pre><code>{
//
// 	"Code": 0,
//
// 	"DataId": "1400000000_test_data_id",
//
// 	"ScanFinishTime": 1566720906,
//
// 	"HitFlag": true,
//
// 	"Live": false,
//
// 	"Msg": "",
//
// 	"ScanPiece": [{
//
// 		"DumpUrl": "",
//
// 		"HitFlag": true,
//
// 		"MainType": "abuse",
//
// 		"RoomId": "123",
//
// 		"OpenId": "xxx",
//
// 		"Info":"",
//
// 		"Offset": 0,
//
// 		"Duration": 3400,
//
// 		"PieceStartTime":1574684231,
//
// 		"ScanDetail": [{
//
// 			"EndTime": 1110,
//
// 			"KeyWord": "xxx",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 1110
//
// 		}, {
//
// 			"EndTime": 1380,
//
// 			"KeyWord": "xxx",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 930
//
// 		}, {
//
// 			"EndTime": 1560,
//
// 			"KeyWord": "xxx",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 930
//
// 		}, {
//
// 			"EndTime": 2820,
//
// 			"KeyWord": "xxx",
//
// 			"Label": "abuse",
//
// 			"Rate": "90.00",
//
// 			"StartTime": 2490
//
// 		}]
//
// 	}],
//
// 	"ScanStartTime": 1566720905,
//
// 	"Scenes": [
//
// 		"default"
//
// 	],
//
// 	"Status": "Success",
//
// 	"TaskId": "xxx",
//
// 	"Url": "https://xxx/xxx.m4a"
//
// }
//
// </code></pre>
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CALLBACKADDRESS = "InvalidParameter.CallbackAddress"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ScanVoiceWithContext(ctx context.Context, request *ScanVoiceRequest) (response *ScanVoiceResponse, err error) {
    if request == nil {
        request = NewScanVoiceRequest()
    }
    request.SetContext(ctx)
    
    response = NewScanVoiceResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateScanRoomsRequest() (request *UpdateScanRoomsRequest) {
    request = &UpdateScanRoomsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "UpdateScanRooms")
    
    
    return
}

func NewUpdateScanRoomsResponse() (response *UpdateScanRoomsResponse) {
    response = &UpdateScanRoomsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateScanRooms
// 更新自定义送检房间号
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CALLBACKADDRESS = "InvalidParameter.CallbackAddress"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateScanRooms(request *UpdateScanRoomsRequest) (response *UpdateScanRoomsResponse, err error) {
    if request == nil {
        request = NewUpdateScanRoomsRequest()
    }
    
    response = NewUpdateScanRoomsResponse()
    err = c.Send(request, response)
    return
}

// UpdateScanRooms
// 更新自定义送检房间号
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CALLBACKADDRESS = "InvalidParameter.CallbackAddress"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateScanRoomsWithContext(ctx context.Context, request *UpdateScanRoomsRequest) (response *UpdateScanRoomsResponse, err error) {
    if request == nil {
        request = NewUpdateScanRoomsRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateScanRoomsResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateScanUsersRequest() (request *UpdateScanUsersRequest) {
    request = &UpdateScanUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "UpdateScanUsers")
    
    
    return
}

func NewUpdateScanUsersResponse() (response *UpdateScanUsersResponse) {
    response = &UpdateScanUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateScanUsers
// 更新自定义送检用户号
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CALLBACKADDRESS = "InvalidParameter.CallbackAddress"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateScanUsers(request *UpdateScanUsersRequest) (response *UpdateScanUsersResponse, err error) {
    if request == nil {
        request = NewUpdateScanUsersRequest()
    }
    
    response = NewUpdateScanUsersResponse()
    err = c.Send(request, response)
    return
}

// UpdateScanUsers
// 更新自定义送检用户号
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CALLBACKADDRESS = "InvalidParameter.CallbackAddress"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateScanUsersWithContext(ctx context.Context, request *UpdateScanUsersRequest) (response *UpdateScanUsersResponse, err error) {
    if request == nil {
        request = NewUpdateScanUsersRequest()
    }
    request.SetContext(ctx)
    
    response = NewUpdateScanUsersResponse()
    err = c.Send(request, response)
    return
}

func NewVoiceFilterRequest() (request *VoiceFilterRequest) {
    request = &VoiceFilterRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("gme", APIVersion, "VoiceFilter")
    
    
    return
}

func NewVoiceFilterResponse() (response *VoiceFilterResponse) {
    response = &VoiceFilterResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// VoiceFilter
// 本接口用于识别涉黄等违规音频，成功会回调配置在应用的回调地址。回调示例如下：
//
// {"BizId":0,"FileId":"test_file_id","FileName":"test_file_name","FileUrl":"test_file_url","OpenId":"test_open_id","TimeStamp":"0000-00-00 00:00:00","Data":[{"Type":1,"Word":"xx"}]}
//
// Type表示过滤类型，1：色情，2：谩骂
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) VoiceFilter(request *VoiceFilterRequest) (response *VoiceFilterResponse, err error) {
    if request == nil {
        request = NewVoiceFilterRequest()
    }
    
    response = NewVoiceFilterResponse()
    err = c.Send(request, response)
    return
}

// VoiceFilter
// 本接口用于识别涉黄等违规音频，成功会回调配置在应用的回调地址。回调示例如下：
//
// {"BizId":0,"FileId":"test_file_id","FileName":"test_file_name","FileUrl":"test_file_url","OpenId":"test_open_id","TimeStamp":"0000-00-00 00:00:00","Data":[{"Type":1,"Word":"xx"}]}
//
// Type表示过滤类型，1：色情，2：谩骂
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) VoiceFilterWithContext(ctx context.Context, request *VoiceFilterRequest) (response *VoiceFilterResponse, err error) {
    if request == nil {
        request = NewVoiceFilterRequest()
    }
    request.SetContext(ctx)
    
    response = NewVoiceFilterResponse()
    err = c.Send(request, response)
    return
}
