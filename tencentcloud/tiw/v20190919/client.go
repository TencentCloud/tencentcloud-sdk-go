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

package v20190919

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-09-19"

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


func NewCreateSnapshotTaskRequest() (request *CreateSnapshotTaskRequest) {
    request = &CreateSnapshotTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "CreateSnapshotTask")
    
    
    return
}

func NewCreateSnapshotTaskResponse() (response *CreateSnapshotTaskResponse) {
    response = &CreateSnapshotTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateSnapshotTask
// 创建白板板书生成任务, 在任务结束后，如果提供了回调地址，将通过回调地址通知板书生成结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateSnapshotTask(request *CreateSnapshotTaskRequest) (response *CreateSnapshotTaskResponse, err error) {
    return c.CreateSnapshotTaskWithContext(context.Background(), request)
}

// CreateSnapshotTask
// 创建白板板书生成任务, 在任务结束后，如果提供了回调地址，将通过回调地址通知板书生成结果
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateSnapshotTaskWithContext(ctx context.Context, request *CreateSnapshotTaskRequest) (response *CreateSnapshotTaskResponse, err error) {
    if request == nil {
        request = NewCreateSnapshotTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSnapshotTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSnapshotTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTranscodeRequest() (request *CreateTranscodeRequest) {
    request = &CreateTranscodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "CreateTranscode")
    
    
    return
}

func NewCreateTranscodeResponse() (response *CreateTranscodeResponse) {
    response = &CreateTranscodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTranscode
// 创建一个文档转码任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_TRANSCODE = "FailedOperation.Transcode"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_FILEFORMATUNSUPPORTED = "InvalidParameter.FileFormatUnsupported"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateTranscode(request *CreateTranscodeRequest) (response *CreateTranscodeResponse, err error) {
    return c.CreateTranscodeWithContext(context.Background(), request)
}

// CreateTranscode
// 创建一个文档转码任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_TRANSCODE = "FailedOperation.Transcode"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_FILEFORMATUNSUPPORTED = "InvalidParameter.FileFormatUnsupported"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateTranscodeWithContext(ctx context.Context, request *CreateTranscodeRequest) (response *CreateTranscodeResponse, err error) {
    if request == nil {
        request = NewCreateTranscodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateTranscode require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateTranscodeResponse()
    err = c.Send(request, response)
    return
}

func NewCreateVideoGenerationTaskRequest() (request *CreateVideoGenerationTaskRequest) {
    request = &CreateVideoGenerationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "CreateVideoGenerationTask")
    
    
    return
}

func NewCreateVideoGenerationTaskResponse() (response *CreateVideoGenerationTaskResponse) {
    response = &CreateVideoGenerationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateVideoGenerationTask
// 创建视频生成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDEXTRA = "InvalidParameter.InvalidExtra"
//  INVALIDPARAMETER_RECORDPARAMETER = "InvalidParameter.RecordParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEINUSE_RECORDUSERID = "ResourceInUse.RecordUserId"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateVideoGenerationTask(request *CreateVideoGenerationTaskRequest) (response *CreateVideoGenerationTaskResponse, err error) {
    return c.CreateVideoGenerationTaskWithContext(context.Background(), request)
}

// CreateVideoGenerationTask
// 创建视频生成任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDEXTRA = "InvalidParameter.InvalidExtra"
//  INVALIDPARAMETER_RECORDPARAMETER = "InvalidParameter.RecordParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEINUSE_RECORDUSERID = "ResourceInUse.RecordUserId"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateVideoGenerationTaskWithContext(ctx context.Context, request *CreateVideoGenerationTaskRequest) (response *CreateVideoGenerationTaskResponse, err error) {
    if request == nil {
        request = NewCreateVideoGenerationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateVideoGenerationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateVideoGenerationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOnlineRecordRequest() (request *DescribeOnlineRecordRequest) {
    request = &DescribeOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeOnlineRecord")
    
    
    return
}

func NewDescribeOnlineRecordResponse() (response *DescribeOnlineRecordResponse) {
    response = &DescribeOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOnlineRecord
// 查询录制任务状态与结果
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RECORD = "FailedOperation.Record"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOnlineRecord(request *DescribeOnlineRecordRequest) (response *DescribeOnlineRecordResponse, err error) {
    return c.DescribeOnlineRecordWithContext(context.Background(), request)
}

// DescribeOnlineRecord
// 查询录制任务状态与结果
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RECORD = "FailedOperation.Record"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOnlineRecordWithContext(ctx context.Context, request *DescribeOnlineRecordRequest) (response *DescribeOnlineRecordResponse, err error) {
    if request == nil {
        request = NewDescribeOnlineRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOnlineRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOnlineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOnlineRecordCallbackRequest() (request *DescribeOnlineRecordCallbackRequest) {
    request = &DescribeOnlineRecordCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeOnlineRecordCallback")
    
    
    return
}

func NewDescribeOnlineRecordCallbackResponse() (response *DescribeOnlineRecordCallbackResponse) {
    response = &DescribeOnlineRecordCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeOnlineRecordCallback
// 查询实时录制回调地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOnlineRecordCallback(request *DescribeOnlineRecordCallbackRequest) (response *DescribeOnlineRecordCallbackResponse, err error) {
    return c.DescribeOnlineRecordCallbackWithContext(context.Background(), request)
}

// DescribeOnlineRecordCallback
// 查询实时录制回调地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOnlineRecordCallbackWithContext(ctx context.Context, request *DescribeOnlineRecordCallbackRequest) (response *DescribeOnlineRecordCallbackResponse, err error) {
    if request == nil {
        request = NewDescribeOnlineRecordCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOnlineRecordCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOnlineRecordCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQualityMetricsRequest() (request *DescribeQualityMetricsRequest) {
    request = &DescribeQualityMetricsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeQualityMetrics")
    
    
    return
}

func NewDescribeQualityMetricsResponse() (response *DescribeQualityMetricsResponse) {
    response = &DescribeQualityMetricsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeQualityMetrics
// 查询互动白板质量数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeQualityMetrics(request *DescribeQualityMetricsRequest) (response *DescribeQualityMetricsResponse, err error) {
    return c.DescribeQualityMetricsWithContext(context.Background(), request)
}

// DescribeQualityMetrics
// 查询互动白板质量数据
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeQualityMetricsWithContext(ctx context.Context, request *DescribeQualityMetricsRequest) (response *DescribeQualityMetricsResponse, err error) {
    if request == nil {
        request = NewDescribeQualityMetricsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeQualityMetrics require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeQualityMetricsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotTaskRequest() (request *DescribeSnapshotTaskRequest) {
    request = &DescribeSnapshotTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeSnapshotTask")
    
    
    return
}

func NewDescribeSnapshotTaskResponse() (response *DescribeSnapshotTaskResponse) {
    response = &DescribeSnapshotTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeSnapshotTask
// 获取指定白板板书生成任务信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSnapshotTask(request *DescribeSnapshotTaskRequest) (response *DescribeSnapshotTaskResponse, err error) {
    return c.DescribeSnapshotTaskWithContext(context.Background(), request)
}

// DescribeSnapshotTask
// 获取指定白板板书生成任务信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeSnapshotTaskWithContext(ctx context.Context, request *DescribeSnapshotTaskRequest) (response *DescribeSnapshotTaskResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSnapshotTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSnapshotTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTIWDailyUsageRequest() (request *DescribeTIWDailyUsageRequest) {
    request = &DescribeTIWDailyUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeTIWDailyUsage")
    
    
    return
}

func NewDescribeTIWDailyUsageResponse() (response *DescribeTIWDailyUsageResponse) {
    response = &DescribeTIWDailyUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTIWDailyUsage
// 查询互动白板天维度计费用量。
//
// 1. 单次查询统计区间最多不能超过31天。
//
// 2. 由于统计延迟等原因，暂时不支持查询当天数据，建议在次日上午7点以后再来查询前一天的用量，例如在10月27日上午7点后，再来查询到10月26日整天的用量
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTIWDailyUsage(request *DescribeTIWDailyUsageRequest) (response *DescribeTIWDailyUsageResponse, err error) {
    return c.DescribeTIWDailyUsageWithContext(context.Background(), request)
}

// DescribeTIWDailyUsage
// 查询互动白板天维度计费用量。
//
// 1. 单次查询统计区间最多不能超过31天。
//
// 2. 由于统计延迟等原因，暂时不支持查询当天数据，建议在次日上午7点以后再来查询前一天的用量，例如在10月27日上午7点后，再来查询到10月26日整天的用量
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTIWDailyUsageWithContext(ctx context.Context, request *DescribeTIWDailyUsageRequest) (response *DescribeTIWDailyUsageResponse, err error) {
    if request == nil {
        request = NewDescribeTIWDailyUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTIWDailyUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTIWDailyUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTIWRoomDailyUsageRequest() (request *DescribeTIWRoomDailyUsageRequest) {
    request = &DescribeTIWRoomDailyUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeTIWRoomDailyUsage")
    
    
    return
}

func NewDescribeTIWRoomDailyUsageResponse() (response *DescribeTIWRoomDailyUsageResponse) {
    response = &DescribeTIWRoomDailyUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTIWRoomDailyUsage
// 查询互动白板房间维度每天计费用量。
//
// 1. 单次查询统计区间最多不能超过31天。
//
// 2. 由于统计延迟等原因，暂时不支持查询当天数据，建议在次日上午7点以后再来查询前一天的用量，例如在10月27日上午7点后，再来查询到10月26日整天的用量
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTIWRoomDailyUsage(request *DescribeTIWRoomDailyUsageRequest) (response *DescribeTIWRoomDailyUsageResponse, err error) {
    return c.DescribeTIWRoomDailyUsageWithContext(context.Background(), request)
}

// DescribeTIWRoomDailyUsage
// 查询互动白板房间维度每天计费用量。
//
// 1. 单次查询统计区间最多不能超过31天。
//
// 2. 由于统计延迟等原因，暂时不支持查询当天数据，建议在次日上午7点以后再来查询前一天的用量，例如在10月27日上午7点后，再来查询到10月26日整天的用量
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTIWRoomDailyUsageWithContext(ctx context.Context, request *DescribeTIWRoomDailyUsageRequest) (response *DescribeTIWRoomDailyUsageResponse, err error) {
    if request == nil {
        request = NewDescribeTIWRoomDailyUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTIWRoomDailyUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTIWRoomDailyUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTranscodeRequest() (request *DescribeTranscodeRequest) {
    request = &DescribeTranscodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeTranscode")
    
    
    return
}

func NewDescribeTranscodeResponse() (response *DescribeTranscodeResponse) {
    response = &DescribeTranscodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTranscode
// 查询文档转码任务的执行进度与转码结果
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_FILEDOWNLOADFAIL = "FailedOperation.FileDownloadFail"
//  FAILEDOPERATION_FILEFORMATERROR = "FailedOperation.FileFormatError"
//  FAILEDOPERATION_FILEOPENFAIL = "FailedOperation.FileOpenFail"
//  FAILEDOPERATION_FILEUPLOADFAIL = "FailedOperation.FileUploadFail"
//  FAILEDOPERATION_TRANSCODE = "FailedOperation.Transcode"
//  FAILEDOPERATION_TRANSCODESERVERERROR = "FailedOperation.TranscodeServerError"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_FILEFORMATUNSUPPORTED = "InvalidParameter.FileFormatUnsupported"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  INVALIDPARAMETER_URLFORMATERROR = "InvalidParameter.UrlFormatError"
//  LIMITEXCEEDED_TRANSCODEPAGESLIMITATION = "LimitExceeded.TranscodePagesLimitation"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTranscode(request *DescribeTranscodeRequest) (response *DescribeTranscodeResponse, err error) {
    return c.DescribeTranscodeWithContext(context.Background(), request)
}

// DescribeTranscode
// 查询文档转码任务的执行进度与转码结果
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_FILEDOWNLOADFAIL = "FailedOperation.FileDownloadFail"
//  FAILEDOPERATION_FILEFORMATERROR = "FailedOperation.FileFormatError"
//  FAILEDOPERATION_FILEOPENFAIL = "FailedOperation.FileOpenFail"
//  FAILEDOPERATION_FILEUPLOADFAIL = "FailedOperation.FileUploadFail"
//  FAILEDOPERATION_TRANSCODE = "FailedOperation.Transcode"
//  FAILEDOPERATION_TRANSCODESERVERERROR = "FailedOperation.TranscodeServerError"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_FILEFORMATUNSUPPORTED = "InvalidParameter.FileFormatUnsupported"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  INVALIDPARAMETER_TRANSCODEPARAMETER = "InvalidParameter.TranscodeParameter"
//  INVALIDPARAMETER_URLFORMATERROR = "InvalidParameter.UrlFormatError"
//  LIMITEXCEEDED_TRANSCODEPAGESLIMITATION = "LimitExceeded.TranscodePagesLimitation"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTranscodeWithContext(ctx context.Context, request *DescribeTranscodeRequest) (response *DescribeTranscodeResponse, err error) {
    if request == nil {
        request = NewDescribeTranscodeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTranscode require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTranscodeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTranscodeCallbackRequest() (request *DescribeTranscodeCallbackRequest) {
    request = &DescribeTranscodeCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeTranscodeCallback")
    
    
    return
}

func NewDescribeTranscodeCallbackResponse() (response *DescribeTranscodeCallbackResponse) {
    response = &DescribeTranscodeCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTranscodeCallback
// 查询文档转码回调地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTranscodeCallback(request *DescribeTranscodeCallbackRequest) (response *DescribeTranscodeCallbackResponse, err error) {
    return c.DescribeTranscodeCallbackWithContext(context.Background(), request)
}

// DescribeTranscodeCallback
// 查询文档转码回调地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTranscodeCallbackWithContext(ctx context.Context, request *DescribeTranscodeCallbackRequest) (response *DescribeTranscodeCallbackResponse, err error) {
    if request == nil {
        request = NewDescribeTranscodeCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTranscodeCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTranscodeCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoGenerationTaskRequest() (request *DescribeVideoGenerationTaskRequest) {
    request = &DescribeVideoGenerationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeVideoGenerationTask")
    
    
    return
}

func NewDescribeVideoGenerationTaskResponse() (response *DescribeVideoGenerationTaskResponse) {
    response = &DescribeVideoGenerationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVideoGenerationTask
// 查询录制视频生成任务状态与结果
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RECORD = "FailedOperation.Record"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeVideoGenerationTask(request *DescribeVideoGenerationTaskRequest) (response *DescribeVideoGenerationTaskResponse, err error) {
    return c.DescribeVideoGenerationTaskWithContext(context.Background(), request)
}

// DescribeVideoGenerationTask
// 查询录制视频生成任务状态与结果
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RECORD = "FailedOperation.Record"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeVideoGenerationTaskWithContext(ctx context.Context, request *DescribeVideoGenerationTaskRequest) (response *DescribeVideoGenerationTaskResponse, err error) {
    if request == nil {
        request = NewDescribeVideoGenerationTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoGenerationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoGenerationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoGenerationTaskCallbackRequest() (request *DescribeVideoGenerationTaskCallbackRequest) {
    request = &DescribeVideoGenerationTaskCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeVideoGenerationTaskCallback")
    
    
    return
}

func NewDescribeVideoGenerationTaskCallbackResponse() (response *DescribeVideoGenerationTaskCallbackResponse) {
    response = &DescribeVideoGenerationTaskCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeVideoGenerationTaskCallback
// 查询录制视频生成回调地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeVideoGenerationTaskCallback(request *DescribeVideoGenerationTaskCallbackRequest) (response *DescribeVideoGenerationTaskCallbackResponse, err error) {
    return c.DescribeVideoGenerationTaskCallbackWithContext(context.Background(), request)
}

// DescribeVideoGenerationTaskCallback
// 查询录制视频生成回调地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeVideoGenerationTaskCallbackWithContext(ctx context.Context, request *DescribeVideoGenerationTaskCallbackRequest) (response *DescribeVideoGenerationTaskCallbackResponse, err error) {
    if request == nil {
        request = NewDescribeVideoGenerationTaskCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoGenerationTaskCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoGenerationTaskCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteboardPushRequest() (request *DescribeWhiteboardPushRequest) {
    request = &DescribeWhiteboardPushRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeWhiteboardPush")
    
    
    return
}

func NewDescribeWhiteboardPushResponse() (response *DescribeWhiteboardPushResponse) {
    response = &DescribeWhiteboardPushResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWhiteboardPush
// 查询推流任务状态与结果
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_WHITEBOARDPUSH = "FailedOperation.WhiteboardPush"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardPush(request *DescribeWhiteboardPushRequest) (response *DescribeWhiteboardPushResponse, err error) {
    return c.DescribeWhiteboardPushWithContext(context.Background(), request)
}

// DescribeWhiteboardPush
// 查询推流任务状态与结果
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_WHITEBOARDPUSH = "FailedOperation.WhiteboardPush"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardPushWithContext(ctx context.Context, request *DescribeWhiteboardPushRequest) (response *DescribeWhiteboardPushResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteboardPushRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWhiteboardPush require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWhiteboardPushResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteboardPushCallbackRequest() (request *DescribeWhiteboardPushCallbackRequest) {
    request = &DescribeWhiteboardPushCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeWhiteboardPushCallback")
    
    
    return
}

func NewDescribeWhiteboardPushCallbackResponse() (response *DescribeWhiteboardPushCallbackResponse) {
    response = &DescribeWhiteboardPushCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeWhiteboardPushCallback
// 查询白板推流回调地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardPushCallback(request *DescribeWhiteboardPushCallbackRequest) (response *DescribeWhiteboardPushCallbackResponse, err error) {
    return c.DescribeWhiteboardPushCallbackWithContext(context.Background(), request)
}

// DescribeWhiteboardPushCallback
// 查询白板推流回调地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardPushCallbackWithContext(ctx context.Context, request *DescribeWhiteboardPushCallbackRequest) (response *DescribeWhiteboardPushCallbackResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteboardPushCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWhiteboardPushCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWhiteboardPushCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewPauseOnlineRecordRequest() (request *PauseOnlineRecordRequest) {
    request = &PauseOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "PauseOnlineRecord")
    
    
    return
}

func NewPauseOnlineRecordResponse() (response *PauseOnlineRecordResponse) {
    response = &PauseOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// PauseOnlineRecord
// 暂停实时录制
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_INVALIDTASKSTATUS = "UnsupportedOperation.InvalidTaskStatus"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) PauseOnlineRecord(request *PauseOnlineRecordRequest) (response *PauseOnlineRecordResponse, err error) {
    return c.PauseOnlineRecordWithContext(context.Background(), request)
}

// PauseOnlineRecord
// 暂停实时录制
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_INVALIDTASKSTATUS = "UnsupportedOperation.InvalidTaskStatus"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) PauseOnlineRecordWithContext(ctx context.Context, request *PauseOnlineRecordRequest) (response *PauseOnlineRecordResponse, err error) {
    if request == nil {
        request = NewPauseOnlineRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("PauseOnlineRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewPauseOnlineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewResumeOnlineRecordRequest() (request *ResumeOnlineRecordRequest) {
    request = &ResumeOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "ResumeOnlineRecord")
    
    
    return
}

func NewResumeOnlineRecordResponse() (response *ResumeOnlineRecordResponse) {
    response = &ResumeOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ResumeOnlineRecord
// 恢复实时录制
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_INVALIDTASKSTATUS = "UnsupportedOperation.InvalidTaskStatus"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) ResumeOnlineRecord(request *ResumeOnlineRecordRequest) (response *ResumeOnlineRecordResponse, err error) {
    return c.ResumeOnlineRecordWithContext(context.Background(), request)
}

// ResumeOnlineRecord
// 恢复实时录制
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_INVALIDTASKSTATUS = "UnsupportedOperation.InvalidTaskStatus"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) ResumeOnlineRecordWithContext(ctx context.Context, request *ResumeOnlineRecordRequest) (response *ResumeOnlineRecordResponse, err error) {
    if request == nil {
        request = NewResumeOnlineRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ResumeOnlineRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewResumeOnlineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewSetOnlineRecordCallbackRequest() (request *SetOnlineRecordCallbackRequest) {
    request = &SetOnlineRecordCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetOnlineRecordCallback")
    
    
    return
}

func NewSetOnlineRecordCallbackResponse() (response *SetOnlineRecordCallbackResponse) {
    response = &SetOnlineRecordCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetOnlineRecordCallback
// 设置实时录制回调地址，回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40258
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetOnlineRecordCallback(request *SetOnlineRecordCallbackRequest) (response *SetOnlineRecordCallbackResponse, err error) {
    return c.SetOnlineRecordCallbackWithContext(context.Background(), request)
}

// SetOnlineRecordCallback
// 设置实时录制回调地址，回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40258
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetOnlineRecordCallbackWithContext(ctx context.Context, request *SetOnlineRecordCallbackRequest) (response *SetOnlineRecordCallbackResponse, err error) {
    if request == nil {
        request = NewSetOnlineRecordCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetOnlineRecordCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetOnlineRecordCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewSetOnlineRecordCallbackKeyRequest() (request *SetOnlineRecordCallbackKeyRequest) {
    request = &SetOnlineRecordCallbackKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetOnlineRecordCallbackKey")
    
    
    return
}

func NewSetOnlineRecordCallbackKeyResponse() (response *SetOnlineRecordCallbackKeyResponse) {
    response = &SetOnlineRecordCallbackKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetOnlineRecordCallbackKey
// 设置实时录制回调鉴权密钥，回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetOnlineRecordCallbackKey(request *SetOnlineRecordCallbackKeyRequest) (response *SetOnlineRecordCallbackKeyResponse, err error) {
    return c.SetOnlineRecordCallbackKeyWithContext(context.Background(), request)
}

// SetOnlineRecordCallbackKey
// 设置实时录制回调鉴权密钥，回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetOnlineRecordCallbackKeyWithContext(ctx context.Context, request *SetOnlineRecordCallbackKeyRequest) (response *SetOnlineRecordCallbackKeyResponse, err error) {
    if request == nil {
        request = NewSetOnlineRecordCallbackKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetOnlineRecordCallbackKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetOnlineRecordCallbackKeyResponse()
    err = c.Send(request, response)
    return
}

func NewSetTranscodeCallbackRequest() (request *SetTranscodeCallbackRequest) {
    request = &SetTranscodeCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetTranscodeCallback")
    
    
    return
}

func NewSetTranscodeCallbackResponse() (response *SetTranscodeCallbackResponse) {
    response = &SetTranscodeCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetTranscodeCallback
// 设置文档转码回调地址，回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40260
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetTranscodeCallback(request *SetTranscodeCallbackRequest) (response *SetTranscodeCallbackResponse, err error) {
    return c.SetTranscodeCallbackWithContext(context.Background(), request)
}

// SetTranscodeCallback
// 设置文档转码回调地址，回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40260
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetTranscodeCallbackWithContext(ctx context.Context, request *SetTranscodeCallbackRequest) (response *SetTranscodeCallbackResponse, err error) {
    if request == nil {
        request = NewSetTranscodeCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetTranscodeCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetTranscodeCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewSetTranscodeCallbackKeyRequest() (request *SetTranscodeCallbackKeyRequest) {
    request = &SetTranscodeCallbackKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetTranscodeCallbackKey")
    
    
    return
}

func NewSetTranscodeCallbackKeyResponse() (response *SetTranscodeCallbackKeyResponse) {
    response = &SetTranscodeCallbackKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetTranscodeCallbackKey
// 设置文档转码回调鉴权密钥，回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetTranscodeCallbackKey(request *SetTranscodeCallbackKeyRequest) (response *SetTranscodeCallbackKeyResponse, err error) {
    return c.SetTranscodeCallbackKeyWithContext(context.Background(), request)
}

// SetTranscodeCallbackKey
// 设置文档转码回调鉴权密钥，回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetTranscodeCallbackKeyWithContext(ctx context.Context, request *SetTranscodeCallbackKeyRequest) (response *SetTranscodeCallbackKeyResponse, err error) {
    if request == nil {
        request = NewSetTranscodeCallbackKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetTranscodeCallbackKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetTranscodeCallbackKeyResponse()
    err = c.Send(request, response)
    return
}

func NewSetVideoGenerationTaskCallbackRequest() (request *SetVideoGenerationTaskCallbackRequest) {
    request = &SetVideoGenerationTaskCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetVideoGenerationTaskCallback")
    
    
    return
}

func NewSetVideoGenerationTaskCallbackResponse() (response *SetVideoGenerationTaskCallbackResponse) {
    response = &SetVideoGenerationTaskCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetVideoGenerationTaskCallback
// 设置录制视频生成回调地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetVideoGenerationTaskCallback(request *SetVideoGenerationTaskCallbackRequest) (response *SetVideoGenerationTaskCallbackResponse, err error) {
    return c.SetVideoGenerationTaskCallbackWithContext(context.Background(), request)
}

// SetVideoGenerationTaskCallback
// 设置录制视频生成回调地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetVideoGenerationTaskCallbackWithContext(ctx context.Context, request *SetVideoGenerationTaskCallbackRequest) (response *SetVideoGenerationTaskCallbackResponse, err error) {
    if request == nil {
        request = NewSetVideoGenerationTaskCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetVideoGenerationTaskCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetVideoGenerationTaskCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewSetVideoGenerationTaskCallbackKeyRequest() (request *SetVideoGenerationTaskCallbackKeyRequest) {
    request = &SetVideoGenerationTaskCallbackKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetVideoGenerationTaskCallbackKey")
    
    
    return
}

func NewSetVideoGenerationTaskCallbackKeyResponse() (response *SetVideoGenerationTaskCallbackKeyResponse) {
    response = &SetVideoGenerationTaskCallbackKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetVideoGenerationTaskCallbackKey
// 设置视频生成回调鉴权密钥
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetVideoGenerationTaskCallbackKey(request *SetVideoGenerationTaskCallbackKeyRequest) (response *SetVideoGenerationTaskCallbackKeyResponse, err error) {
    return c.SetVideoGenerationTaskCallbackKeyWithContext(context.Background(), request)
}

// SetVideoGenerationTaskCallbackKey
// 设置视频生成回调鉴权密钥
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetVideoGenerationTaskCallbackKeyWithContext(ctx context.Context, request *SetVideoGenerationTaskCallbackKeyRequest) (response *SetVideoGenerationTaskCallbackKeyResponse, err error) {
    if request == nil {
        request = NewSetVideoGenerationTaskCallbackKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetVideoGenerationTaskCallbackKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetVideoGenerationTaskCallbackKeyResponse()
    err = c.Send(request, response)
    return
}

func NewSetWhiteboardPushCallbackRequest() (request *SetWhiteboardPushCallbackRequest) {
    request = &SetWhiteboardPushCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetWhiteboardPushCallback")
    
    
    return
}

func NewSetWhiteboardPushCallbackResponse() (response *SetWhiteboardPushCallbackResponse) {
    response = &SetWhiteboardPushCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetWhiteboardPushCallback
// 设置白板推流回调地址，回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40257
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetWhiteboardPushCallback(request *SetWhiteboardPushCallbackRequest) (response *SetWhiteboardPushCallbackResponse, err error) {
    return c.SetWhiteboardPushCallbackWithContext(context.Background(), request)
}

// SetWhiteboardPushCallback
// 设置白板推流回调地址，回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40257
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetWhiteboardPushCallbackWithContext(ctx context.Context, request *SetWhiteboardPushCallbackRequest) (response *SetWhiteboardPushCallbackResponse, err error) {
    if request == nil {
        request = NewSetWhiteboardPushCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetWhiteboardPushCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetWhiteboardPushCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewSetWhiteboardPushCallbackKeyRequest() (request *SetWhiteboardPushCallbackKeyRequest) {
    request = &SetWhiteboardPushCallbackKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetWhiteboardPushCallbackKey")
    
    
    return
}

func NewSetWhiteboardPushCallbackKeyResponse() (response *SetWhiteboardPushCallbackKeyResponse) {
    response = &SetWhiteboardPushCallbackKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetWhiteboardPushCallbackKey
// 设置白板推流回调鉴权密钥，回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetWhiteboardPushCallbackKey(request *SetWhiteboardPushCallbackKeyRequest) (response *SetWhiteboardPushCallbackKeyResponse, err error) {
    return c.SetWhiteboardPushCallbackKeyWithContext(context.Background(), request)
}

// SetWhiteboardPushCallbackKey
// 设置白板推流回调鉴权密钥，回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetWhiteboardPushCallbackKeyWithContext(ctx context.Context, request *SetWhiteboardPushCallbackKeyRequest) (response *SetWhiteboardPushCallbackKeyResponse, err error) {
    if request == nil {
        request = NewSetWhiteboardPushCallbackKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetWhiteboardPushCallbackKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetWhiteboardPushCallbackKeyResponse()
    err = c.Send(request, response)
    return
}

func NewStartOnlineRecordRequest() (request *StartOnlineRecordRequest) {
    request = &StartOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "StartOnlineRecord")
    
    
    return
}

func NewStartOnlineRecordResponse() (response *StartOnlineRecordResponse) {
    response = &StartOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartOnlineRecord
// 发起一个实时录制任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDEXTRA = "InvalidParameter.InvalidExtra"
//  INVALIDPARAMETER_RECORDPARAMETER = "InvalidParameter.RecordParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEINUSE_RECORDUSERID = "ResourceInUse.RecordUserId"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartOnlineRecord(request *StartOnlineRecordRequest) (response *StartOnlineRecordResponse, err error) {
    return c.StartOnlineRecordWithContext(context.Background(), request)
}

// StartOnlineRecord
// 发起一个实时录制任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDEXTRA = "InvalidParameter.InvalidExtra"
//  INVALIDPARAMETER_RECORDPARAMETER = "InvalidParameter.RecordParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEINUSE_RECORDUSERID = "ResourceInUse.RecordUserId"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartOnlineRecordWithContext(ctx context.Context, request *StartOnlineRecordRequest) (response *StartOnlineRecordResponse, err error) {
    if request == nil {
        request = NewStartOnlineRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartOnlineRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartOnlineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewStartWhiteboardPushRequest() (request *StartWhiteboardPushRequest) {
    request = &StartWhiteboardPushRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "StartWhiteboardPush")
    
    
    return
}

func NewStartWhiteboardPushResponse() (response *StartWhiteboardPushResponse) {
    response = &StartWhiteboardPushResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StartWhiteboardPush
// 发起一个白板推流任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDEXTRA = "InvalidParameter.InvalidExtra"
//  INVALIDPARAMETER_RECORDPARAMETER = "InvalidParameter.RecordParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEINUSE_RECORDUSERID = "ResourceInUse.RecordUserId"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartWhiteboardPush(request *StartWhiteboardPushRequest) (response *StartWhiteboardPushResponse, err error) {
    return c.StartWhiteboardPushWithContext(context.Background(), request)
}

// StartWhiteboardPush
// 发起一个白板推流任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDEXTRA = "InvalidParameter.InvalidExtra"
//  INVALIDPARAMETER_RECORDPARAMETER = "InvalidParameter.RecordParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEINUSE_RECORDUSERID = "ResourceInUse.RecordUserId"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartWhiteboardPushWithContext(ctx context.Context, request *StartWhiteboardPushRequest) (response *StartWhiteboardPushResponse, err error) {
    if request == nil {
        request = NewStartWhiteboardPushRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartWhiteboardPush require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartWhiteboardPushResponse()
    err = c.Send(request, response)
    return
}

func NewStopOnlineRecordRequest() (request *StopOnlineRecordRequest) {
    request = &StopOnlineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "StopOnlineRecord")
    
    
    return
}

func NewStopOnlineRecordResponse() (response *StopOnlineRecordResponse) {
    response = &StopOnlineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopOnlineRecord
// 停止实时录制
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) StopOnlineRecord(request *StopOnlineRecordRequest) (response *StopOnlineRecordResponse, err error) {
    return c.StopOnlineRecordWithContext(context.Background(), request)
}

// StopOnlineRecord
// 停止实时录制
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) StopOnlineRecordWithContext(ctx context.Context, request *StopOnlineRecordRequest) (response *StopOnlineRecordResponse, err error) {
    if request == nil {
        request = NewStopOnlineRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopOnlineRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopOnlineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewStopWhiteboardPushRequest() (request *StopWhiteboardPushRequest) {
    request = &StopWhiteboardPushRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "StopWhiteboardPush")
    
    
    return
}

func NewStopWhiteboardPushResponse() (response *StopWhiteboardPushResponse) {
    response = &StopWhiteboardPushResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StopWhiteboardPush
// 停止白板推流任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) StopWhiteboardPush(request *StopWhiteboardPushRequest) (response *StopWhiteboardPushResponse, err error) {
    return c.StopWhiteboardPushWithContext(context.Background(), request)
}

// StopWhiteboardPush
// 停止白板推流任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
//  UNSUPPORTEDOPERATION_TASKHASALREADYSTOPPED = "UnsupportedOperation.TaskHasAlreadyStopped"
func (c *Client) StopWhiteboardPushWithContext(ctx context.Context, request *StopWhiteboardPushRequest) (response *StopWhiteboardPushResponse, err error) {
    if request == nil {
        request = NewStopWhiteboardPushRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopWhiteboardPush require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopWhiteboardPushResponse()
    err = c.Send(request, response)
    return
}
