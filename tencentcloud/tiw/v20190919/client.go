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


func NewApplyTiwTrialRequest() (request *ApplyTiwTrialRequest) {
    request = &ApplyTiwTrialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "ApplyTiwTrial")
    
    
    return
}

func NewApplyTiwTrialResponse() (response *ApplyTiwTrialResponse) {
    response = &ApplyTiwTrialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ApplyTiwTrial
// 申请互动白板试用，默认15天
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ALREADYENABLED = "FailedOperation.AlreadyEnabled"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) ApplyTiwTrial(request *ApplyTiwTrialRequest) (response *ApplyTiwTrialResponse, err error) {
    return c.ApplyTiwTrialWithContext(context.Background(), request)
}

// ApplyTiwTrial
// 申请互动白板试用，默认15天
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ALREADYENABLED = "FailedOperation.AlreadyEnabled"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) ApplyTiwTrialWithContext(ctx context.Context, request *ApplyTiwTrialRequest) (response *ApplyTiwTrialResponse, err error) {
    if request == nil {
        request = NewApplyTiwTrialRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ApplyTiwTrial require credential")
    }

    request.SetContext(ctx)
    
    response = NewApplyTiwTrialResponse()
    err = c.Send(request, response)
    return
}

func NewCreateApplicationRequest() (request *CreateApplicationRequest) {
    request = &CreateApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "CreateApplication")
    
    
    return
}

func NewCreateApplicationResponse() (response *CreateApplicationResponse) {
    response = &CreateApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateApplication
// 创建白板应用
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATIONALREADYEXISTS = "InvalidParameter.ApplicationAlreadyExists"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateApplication(request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    return c.CreateApplicationWithContext(context.Background(), request)
}

// CreateApplication
// 创建白板应用
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPLICATIONALREADYEXISTS = "InvalidParameter.ApplicationAlreadyExists"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateApplicationWithContext(ctx context.Context, request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
    if request == nil {
        request = NewCreateApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateOfflineRecordRequest() (request *CreateOfflineRecordRequest) {
    request = &CreateOfflineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "CreateOfflineRecord")
    
    
    return
}

func NewCreateOfflineRecordResponse() (response *CreateOfflineRecordResponse) {
    response = &CreateOfflineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateOfflineRecord
// 课后录制服务已下线
//
// 
//
// 创建课后录制任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_RECORDPARAMETER = "InvalidParameter.RecordParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateOfflineRecord(request *CreateOfflineRecordRequest) (response *CreateOfflineRecordResponse, err error) {
    return c.CreateOfflineRecordWithContext(context.Background(), request)
}

// CreateOfflineRecord
// 课后录制服务已下线
//
// 
//
// 创建课后录制任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_RECORDPARAMETER = "InvalidParameter.RecordParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreateOfflineRecordWithContext(ctx context.Context, request *CreateOfflineRecordRequest) (response *CreateOfflineRecordResponse, err error) {
    if request == nil {
        request = NewCreateOfflineRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateOfflineRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateOfflineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePPTCheckTaskRequest() (request *CreatePPTCheckTaskRequest) {
    request = &CreatePPTCheckTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "CreatePPTCheckTask")
    
    
    return
}

func NewCreatePPTCheckTaskResponse() (response *CreatePPTCheckTaskResponse) {
    response = &CreatePPTCheckTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePPTCheckTask
// 检测PPT文件，识别PPT中包含的动态转码任务（Transcode）不支持的元素
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_PREPROCESS = "FailedOperation.Preprocess"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_FILEFORMATUNSUPPORTED = "InvalidParameter.FileFormatUnsupported"
//  INVALIDPARAMETER_PREPROCESSPARAMETER = "InvalidParameter.PreprocessParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreatePPTCheckTask(request *CreatePPTCheckTaskRequest) (response *CreatePPTCheckTaskResponse, err error) {
    return c.CreatePPTCheckTaskWithContext(context.Background(), request)
}

// CreatePPTCheckTask
// 检测PPT文件，识别PPT中包含的动态转码任务（Transcode）不支持的元素
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_PREPROCESS = "FailedOperation.Preprocess"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_FILEFORMATUNSUPPORTED = "InvalidParameter.FileFormatUnsupported"
//  INVALIDPARAMETER_PREPROCESSPARAMETER = "InvalidParameter.PreprocessParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  LIMITEXCEEDED_TASKCONCURRENCY = "LimitExceeded.TaskConcurrency"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) CreatePPTCheckTaskWithContext(ctx context.Context, request *CreatePPTCheckTaskRequest) (response *CreatePPTCheckTaskResponse, err error) {
    if request == nil {
        request = NewCreatePPTCheckTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePPTCheckTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePPTCheckTaskResponse()
    err = c.Send(request, response)
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

func NewDescribeAPIServiceRequest() (request *DescribeAPIServiceRequest) {
    request = &DescribeAPIServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeAPIService")
    
    
    return
}

func NewDescribeAPIServiceResponse() (response *DescribeAPIServiceResponse) {
    response = &DescribeAPIServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAPIService
// 通过服务角色调用其他云产品API接口获取信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAPIService(request *DescribeAPIServiceRequest) (response *DescribeAPIServiceResponse, err error) {
    return c.DescribeAPIServiceWithContext(context.Background(), request)
}

// DescribeAPIService
// 通过服务角色调用其他云产品API接口获取信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeAPIServiceWithContext(ctx context.Context, request *DescribeAPIServiceRequest) (response *DescribeAPIServiceResponse, err error) {
    if request == nil {
        request = NewDescribeAPIServiceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAPIService require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAPIServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationInfosRequest() (request *DescribeApplicationInfosRequest) {
    request = &DescribeApplicationInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeApplicationInfos")
    
    
    return
}

func NewDescribeApplicationInfosResponse() (response *DescribeApplicationInfosResponse) {
    response = &DescribeApplicationInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationInfos
// 查询白板应用详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeApplicationInfos(request *DescribeApplicationInfosRequest) (response *DescribeApplicationInfosResponse, err error) {
    return c.DescribeApplicationInfosWithContext(context.Background(), request)
}

// DescribeApplicationInfos
// 查询白板应用详情
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeApplicationInfosWithContext(ctx context.Context, request *DescribeApplicationInfosRequest) (response *DescribeApplicationInfosResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationInfosRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationInfos require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeApplicationUsageRequest() (request *DescribeApplicationUsageRequest) {
    request = &DescribeApplicationUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeApplicationUsage")
    
    
    return
}

func NewDescribeApplicationUsageResponse() (response *DescribeApplicationUsageResponse) {
    response = &DescribeApplicationUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeApplicationUsage
// 查询互动白板各个子产品用量
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TIMEFORMAT = "InvalidParameter.TimeFormat"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeApplicationUsage(request *DescribeApplicationUsageRequest) (response *DescribeApplicationUsageResponse, err error) {
    return c.DescribeApplicationUsageWithContext(context.Background(), request)
}

// DescribeApplicationUsage
// 查询互动白板各个子产品用量
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TIMEFORMAT = "InvalidParameter.TimeFormat"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeApplicationUsageWithContext(ctx context.Context, request *DescribeApplicationUsageRequest) (response *DescribeApplicationUsageResponse, err error) {
    if request == nil {
        request = NewDescribeApplicationUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeApplicationUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeApplicationUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBoardSDKLogRequest() (request *DescribeBoardSDKLogRequest) {
    request = &DescribeBoardSDKLogRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeBoardSDKLog")
    
    
    return
}

func NewDescribeBoardSDKLogResponse() (response *DescribeBoardSDKLogResponse) {
    response = &DescribeBoardSDKLogResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeBoardSDKLog
// 查询客户端白板日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeBoardSDKLog(request *DescribeBoardSDKLogRequest) (response *DescribeBoardSDKLogResponse, err error) {
    return c.DescribeBoardSDKLogWithContext(context.Background(), request)
}

// DescribeBoardSDKLog
// 查询客户端白板日志
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeBoardSDKLogWithContext(ctx context.Context, request *DescribeBoardSDKLogRequest) (response *DescribeBoardSDKLogResponse, err error) {
    if request == nil {
        request = NewDescribeBoardSDKLogRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeBoardSDKLog require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeBoardSDKLogResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIMApplicationsRequest() (request *DescribeIMApplicationsRequest) {
    request = &DescribeIMApplicationsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeIMApplications")
    
    
    return
}

func NewDescribeIMApplicationsResponse() (response *DescribeIMApplicationsResponse) {
    response = &DescribeIMApplicationsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeIMApplications
// 查询可用于创建白板应用的IM应用列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeIMApplications(request *DescribeIMApplicationsRequest) (response *DescribeIMApplicationsResponse, err error) {
    return c.DescribeIMApplicationsWithContext(context.Background(), request)
}

// DescribeIMApplications
// 查询可用于创建白板应用的IM应用列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeIMApplicationsWithContext(ctx context.Context, request *DescribeIMApplicationsRequest) (response *DescribeIMApplicationsResponse, err error) {
    if request == nil {
        request = NewDescribeIMApplicationsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeIMApplications require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeIMApplicationsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOfflineRecordRequest() (request *DescribeOfflineRecordRequest) {
    request = &DescribeOfflineRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeOfflineRecord")
    
    
    return
}

func NewDescribeOfflineRecordResponse() (response *DescribeOfflineRecordResponse) {
    response = &DescribeOfflineRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOfflineRecord
// 课后录制服务已下线
//
// 
//
// 查询课后录制任务的进度与录制结果等相关信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RECORD = "FailedOperation.Record"
//  FAILEDOPERATION_TRANSCODE = "FailedOperation.Transcode"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOfflineRecord(request *DescribeOfflineRecordRequest) (response *DescribeOfflineRecordResponse, err error) {
    return c.DescribeOfflineRecordWithContext(context.Background(), request)
}

// DescribeOfflineRecord
// 课后录制服务已下线
//
// 
//
// 查询课后录制任务的进度与录制结果等相关信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_RECORD = "FailedOperation.Record"
//  FAILEDOPERATION_TRANSCODE = "FailedOperation.Transcode"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOfflineRecordWithContext(ctx context.Context, request *DescribeOfflineRecordRequest) (response *DescribeOfflineRecordResponse, err error) {
    if request == nil {
        request = NewDescribeOfflineRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOfflineRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOfflineRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOfflineRecordCallbackRequest() (request *DescribeOfflineRecordCallbackRequest) {
    request = &DescribeOfflineRecordCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeOfflineRecordCallback")
    
    
    return
}

func NewDescribeOfflineRecordCallbackResponse() (response *DescribeOfflineRecordCallbackResponse) {
    response = &DescribeOfflineRecordCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeOfflineRecordCallback
// 课后录制服务已下线
//
// 
//
// 查询课后录制回调地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOfflineRecordCallback(request *DescribeOfflineRecordCallbackRequest) (response *DescribeOfflineRecordCallbackResponse, err error) {
    return c.DescribeOfflineRecordCallbackWithContext(context.Background(), request)
}

// DescribeOfflineRecordCallback
// 课后录制服务已下线
//
// 
//
// 查询课后录制回调地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeOfflineRecordCallbackWithContext(ctx context.Context, request *DescribeOfflineRecordCallbackRequest) (response *DescribeOfflineRecordCallbackResponse, err error) {
    if request == nil {
        request = NewDescribeOfflineRecordCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeOfflineRecordCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeOfflineRecordCallbackResponse()
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

func NewDescribePPTCheckRequest() (request *DescribePPTCheckRequest) {
    request = &DescribePPTCheckRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribePPTCheck")
    
    
    return
}

func NewDescribePPTCheckResponse() (response *DescribePPTCheckResponse) {
    response = &DescribePPTCheckResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePPTCheck
// 查询PPT检测任务的执行进度或结果
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_FILEDOWNLOADFAIL = "FailedOperation.FileDownloadFail"
//  FAILEDOPERATION_FILEFORMATERROR = "FailedOperation.FileFormatError"
//  FAILEDOPERATION_FILEOPENFAIL = "FailedOperation.FileOpenFail"
//  FAILEDOPERATION_FILEUPLOADFAIL = "FailedOperation.FileUploadFail"
//  FAILEDOPERATION_PREPROCESS = "FailedOperation.Preprocess"
//  FAILEDOPERATION_PREPROCESSSERVERERROR = "FailedOperation.PreprocessServerError"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_FILEFORMATUNSUPPORTED = "InvalidParameter.FileFormatUnsupported"
//  INVALIDPARAMETER_PREPROCESSPARAMETER = "InvalidParameter.PreprocessParameter"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  INVALIDPARAMETER_URLFORMATERROR = "InvalidParameter.UrlFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribePPTCheck(request *DescribePPTCheckRequest) (response *DescribePPTCheckResponse, err error) {
    return c.DescribePPTCheckWithContext(context.Background(), request)
}

// DescribePPTCheck
// 查询PPT检测任务的执行进度或结果
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_FILEDOWNLOADFAIL = "FailedOperation.FileDownloadFail"
//  FAILEDOPERATION_FILEFORMATERROR = "FailedOperation.FileFormatError"
//  FAILEDOPERATION_FILEOPENFAIL = "FailedOperation.FileOpenFail"
//  FAILEDOPERATION_FILEUPLOADFAIL = "FailedOperation.FileUploadFail"
//  FAILEDOPERATION_PREPROCESS = "FailedOperation.Preprocess"
//  FAILEDOPERATION_PREPROCESSSERVERERROR = "FailedOperation.PreprocessServerError"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_FILEFORMATUNSUPPORTED = "InvalidParameter.FileFormatUnsupported"
//  INVALIDPARAMETER_PREPROCESSPARAMETER = "InvalidParameter.PreprocessParameter"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  INVALIDPARAMETER_URLFORMATERROR = "InvalidParameter.UrlFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribePPTCheckWithContext(ctx context.Context, request *DescribePPTCheckRequest) (response *DescribePPTCheckResponse, err error) {
    if request == nil {
        request = NewDescribePPTCheckRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePPTCheck require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePPTCheckResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePPTCheckCallbackRequest() (request *DescribePPTCheckCallbackRequest) {
    request = &DescribePPTCheckCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribePPTCheckCallback")
    
    
    return
}

func NewDescribePPTCheckCallbackResponse() (response *DescribePPTCheckCallbackResponse) {
    response = &DescribePPTCheckCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePPTCheckCallback
// 查询PPT检测任务回调地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribePPTCheckCallback(request *DescribePPTCheckCallbackRequest) (response *DescribePPTCheckCallbackResponse, err error) {
    return c.DescribePPTCheckCallbackWithContext(context.Background(), request)
}

// DescribePPTCheckCallback
// 查询PPT检测任务回调地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribePPTCheckCallbackWithContext(ctx context.Context, request *DescribePPTCheckCallbackRequest) (response *DescribePPTCheckCallbackResponse, err error) {
    if request == nil {
        request = NewDescribePPTCheckCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePPTCheckCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePPTCheckCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePostpaidUsageRequest() (request *DescribePostpaidUsageRequest) {
    request = &DescribePostpaidUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribePostpaidUsage")
    
    
    return
}

func NewDescribePostpaidUsageResponse() (response *DescribePostpaidUsageResponse) {
    response = &DescribePostpaidUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePostpaidUsage
// 查询用户后付费用量
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TIMEFORMAT = "InvalidParameter.TimeFormat"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribePostpaidUsage(request *DescribePostpaidUsageRequest) (response *DescribePostpaidUsageResponse, err error) {
    return c.DescribePostpaidUsageWithContext(context.Background(), request)
}

// DescribePostpaidUsage
// 查询用户后付费用量
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_TIMEFORMAT = "InvalidParameter.TimeFormat"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribePostpaidUsageWithContext(ctx context.Context, request *DescribePostpaidUsageRequest) (response *DescribePostpaidUsageResponse, err error) {
    if request == nil {
        request = NewDescribePostpaidUsageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePostpaidUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePostpaidUsageResponse()
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

func NewDescribeRecordSearchRequest() (request *DescribeRecordSearchRequest) {
    request = &DescribeRecordSearchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeRecordSearch")
    
    
    return
}

func NewDescribeRecordSearchResponse() (response *DescribeRecordSearchResponse) {
    response = &DescribeRecordSearchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordSearch
// 根据房间号搜索实时录制任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeRecordSearch(request *DescribeRecordSearchRequest) (response *DescribeRecordSearchResponse, err error) {
    return c.DescribeRecordSearchWithContext(context.Background(), request)
}

// DescribeRecordSearch
// 根据房间号搜索实时录制任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeRecordSearchWithContext(ctx context.Context, request *DescribeRecordSearchRequest) (response *DescribeRecordSearchResponse, err error) {
    if request == nil {
        request = NewDescribeRecordSearchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordSearch require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordSearchResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoomListRequest() (request *DescribeRoomListRequest) {
    request = &DescribeRoomListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeRoomList")
    
    
    return
}

func NewDescribeRoomListResponse() (response *DescribeRoomListResponse) {
    response = &DescribeRoomListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRoomList
// 查询白板房间列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeRoomList(request *DescribeRoomListRequest) (response *DescribeRoomListResponse, err error) {
    return c.DescribeRoomListWithContext(context.Background(), request)
}

// DescribeRoomList
// 查询白板房间列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeRoomListWithContext(ctx context.Context, request *DescribeRoomListRequest) (response *DescribeRoomListResponse, err error) {
    if request == nil {
        request = NewDescribeRoomListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoomList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoomListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRunningTasksRequest() (request *DescribeRunningTasksRequest) {
    request = &DescribeRunningTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeRunningTasks")
    
    
    return
}

func NewDescribeRunningTasksResponse() (response *DescribeRunningTasksResponse) {
    response = &DescribeRunningTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRunningTasks
// 根据指定的任务类型，获取当前正在执行中的任务列表。只能查询最近3天内创建的任务。
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
func (c *Client) DescribeRunningTasks(request *DescribeRunningTasksRequest) (response *DescribeRunningTasksResponse, err error) {
    return c.DescribeRunningTasksWithContext(context.Background(), request)
}

// DescribeRunningTasks
// 根据指定的任务类型，获取当前正在执行中的任务列表。只能查询最近3天内创建的任务。
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
func (c *Client) DescribeRunningTasksWithContext(ctx context.Context, request *DescribeRunningTasksRequest) (response *DescribeRunningTasksResponse, err error) {
    if request == nil {
        request = NewDescribeRunningTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRunningTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRunningTasksResponse()
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

func NewDescribeTranscodeByUrlRequest() (request *DescribeTranscodeByUrlRequest) {
    request = &DescribeTranscodeByUrlRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeTranscodeByUrl")
    
    
    return
}

func NewDescribeTranscodeByUrlResponse() (response *DescribeTranscodeByUrlResponse) {
    response = &DescribeTranscodeByUrlResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTranscodeByUrl
// 通过文档URL查询转码任务，返回最近一次的转码任务状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_TRANSCODESERVERERROR = "FailedOperation.TranscodeServerError"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTranscodeByUrl(request *DescribeTranscodeByUrlRequest) (response *DescribeTranscodeByUrlResponse, err error) {
    return c.DescribeTranscodeByUrlWithContext(context.Background(), request)
}

// DescribeTranscodeByUrl
// 通过文档URL查询转码任务，返回最近一次的转码任务状态
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_TRANSCODESERVERERROR = "FailedOperation.TranscodeServerError"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTranscodeByUrlWithContext(ctx context.Context, request *DescribeTranscodeByUrlRequest) (response *DescribeTranscodeByUrlResponse, err error) {
    if request == nil {
        request = NewDescribeTranscodeByUrlRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTranscodeByUrl require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTranscodeByUrlResponse()
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

func NewDescribeTranscodeSearchRequest() (request *DescribeTranscodeSearchRequest) {
    request = &DescribeTranscodeSearchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeTranscodeSearch")
    
    
    return
}

func NewDescribeTranscodeSearchResponse() (response *DescribeTranscodeSearchResponse) {
    response = &DescribeTranscodeSearchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTranscodeSearch
// 按文档名称搜索转码任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTranscodeSearch(request *DescribeTranscodeSearchRequest) (response *DescribeTranscodeSearchResponse, err error) {
    return c.DescribeTranscodeSearchWithContext(context.Background(), request)
}

// DescribeTranscodeSearch
// 按文档名称搜索转码任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeTranscodeSearchWithContext(ctx context.Context, request *DescribeTranscodeSearchRequest) (response *DescribeTranscodeSearchResponse, err error) {
    if request == nil {
        request = NewDescribeTranscodeSearchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTranscodeSearch require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTranscodeSearchResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUsageSummaryRequest() (request *DescribeUsageSummaryRequest) {
    request = &DescribeUsageSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeUsageSummary")
    
    
    return
}

func NewDescribeUsageSummaryResponse() (response *DescribeUsageSummaryResponse) {
    response = &DescribeUsageSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUsageSummary
// 查询指定时间段内子产品的用量汇总
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TIMEFORMAT = "InvalidParameter.TimeFormat"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeUsageSummary(request *DescribeUsageSummaryRequest) (response *DescribeUsageSummaryResponse, err error) {
    return c.DescribeUsageSummaryWithContext(context.Background(), request)
}

// DescribeUsageSummary
// 查询指定时间段内子产品的用量汇总
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_TIMEFORMAT = "InvalidParameter.TimeFormat"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeUsageSummaryWithContext(ctx context.Context, request *DescribeUsageSummaryRequest) (response *DescribeUsageSummaryResponse, err error) {
    if request == nil {
        request = NewDescribeUsageSummaryRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUsageSummary require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUsageSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserListRequest() (request *DescribeUserListRequest) {
    request = &DescribeUserListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeUserList")
    
    
    return
}

func NewDescribeUserListResponse() (response *DescribeUserListResponse) {
    response = &DescribeUserListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserList
// 查询白板用户列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeUserList(request *DescribeUserListRequest) (response *DescribeUserListResponse, err error) {
    return c.DescribeUserListWithContext(context.Background(), request)
}

// DescribeUserList
// 查询白板用户列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeUserListWithContext(ctx context.Context, request *DescribeUserListRequest) (response *DescribeUserListResponse, err error) {
    if request == nil {
        request = NewDescribeUserListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserResourcesRequest() (request *DescribeUserResourcesRequest) {
    request = &DescribeUserResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeUserResources")
    
    
    return
}

func NewDescribeUserResourcesResponse() (response *DescribeUserResourcesResponse) {
    response = &DescribeUserResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserResources
// 查询客户资源列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeUserResources(request *DescribeUserResourcesRequest) (response *DescribeUserResourcesResponse, err error) {
    return c.DescribeUserResourcesWithContext(context.Background(), request)
}

// DescribeUserResources
// 查询客户资源列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeUserResourcesWithContext(ctx context.Context, request *DescribeUserResourcesRequest) (response *DescribeUserResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeUserResourcesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserResources require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserStatusRequest() (request *DescribeUserStatusRequest) {
    request = &DescribeUserStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeUserStatus")
    
    
    return
}

func NewDescribeUserStatusResponse() (response *DescribeUserStatusResponse) {
    response = &DescribeUserStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserStatus
// 查询互动白板用户详情，包括是否开通了互动白板，当前互动白板服务有效期等信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ALREADYENABLED = "FailedOperation.AlreadyEnabled"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeUserStatus(request *DescribeUserStatusRequest) (response *DescribeUserStatusResponse, err error) {
    return c.DescribeUserStatusWithContext(context.Background(), request)
}

// DescribeUserStatus
// 查询互动白板用户详情，包括是否开通了互动白板，当前互动白板服务有效期等信息
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_ALREADYENABLED = "FailedOperation.AlreadyEnabled"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) DescribeUserStatusWithContext(ctx context.Context, request *DescribeUserStatusRequest) (response *DescribeUserStatusResponse, err error) {
    if request == nil {
        request = NewDescribeUserStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserStatusResponse()
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

func NewDescribeWarningCallbackRequest() (request *DescribeWarningCallbackRequest) {
    request = &DescribeWarningCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeWarningCallback")
    
    
    return
}

func NewDescribeWarningCallbackResponse() (response *DescribeWarningCallbackResponse) {
    response = &DescribeWarningCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWarningCallback
// 查询告警回调地址。此功能需要申请白名单使用。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWarningCallback(request *DescribeWarningCallbackRequest) (response *DescribeWarningCallbackResponse, err error) {
    return c.DescribeWarningCallbackWithContext(context.Background(), request)
}

// DescribeWarningCallback
// 查询告警回调地址。此功能需要申请白名单使用。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWarningCallbackWithContext(ctx context.Context, request *DescribeWarningCallbackRequest) (response *DescribeWarningCallbackResponse, err error) {
    if request == nil {
        request = NewDescribeWarningCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWarningCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWarningCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteboardApplicationConfigRequest() (request *DescribeWhiteboardApplicationConfigRequest) {
    request = &DescribeWhiteboardApplicationConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeWhiteboardApplicationConfig")
    
    
    return
}

func NewDescribeWhiteboardApplicationConfigResponse() (response *DescribeWhiteboardApplicationConfigResponse) {
    response = &DescribeWhiteboardApplicationConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWhiteboardApplicationConfig
// 查询白板应用任务相关的配置，包括存储桶、回调等
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardApplicationConfig(request *DescribeWhiteboardApplicationConfigRequest) (response *DescribeWhiteboardApplicationConfigResponse, err error) {
    return c.DescribeWhiteboardApplicationConfigWithContext(context.Background(), request)
}

// DescribeWhiteboardApplicationConfig
// 查询白板应用任务相关的配置，包括存储桶、回调等
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardApplicationConfigWithContext(ctx context.Context, request *DescribeWhiteboardApplicationConfigRequest) (response *DescribeWhiteboardApplicationConfigResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteboardApplicationConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWhiteboardApplicationConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWhiteboardApplicationConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWhiteboardBucketConfigRequest() (request *DescribeWhiteboardBucketConfigRequest) {
    request = &DescribeWhiteboardBucketConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeWhiteboardBucketConfig")
    
    
    return
}

func NewDescribeWhiteboardBucketConfigResponse() (response *DescribeWhiteboardBucketConfigResponse) {
    response = &DescribeWhiteboardBucketConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWhiteboardBucketConfig
// 查询文档转码，实时录制存储桶的配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardBucketConfig(request *DescribeWhiteboardBucketConfigRequest) (response *DescribeWhiteboardBucketConfigResponse, err error) {
    return c.DescribeWhiteboardBucketConfigWithContext(context.Background(), request)
}

// DescribeWhiteboardBucketConfig
// 查询文档转码，实时录制存储桶的配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardBucketConfigWithContext(ctx context.Context, request *DescribeWhiteboardBucketConfigRequest) (response *DescribeWhiteboardBucketConfigResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteboardBucketConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWhiteboardBucketConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWhiteboardBucketConfigResponse()
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

func NewDescribeWhiteboardPushSearchRequest() (request *DescribeWhiteboardPushSearchRequest) {
    request = &DescribeWhiteboardPushSearchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "DescribeWhiteboardPushSearch")
    
    
    return
}

func NewDescribeWhiteboardPushSearchResponse() (response *DescribeWhiteboardPushSearchResponse) {
    response = &DescribeWhiteboardPushSearchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWhiteboardPushSearch
// 根据房间号搜索白板推流任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardPushSearch(request *DescribeWhiteboardPushSearchRequest) (response *DescribeWhiteboardPushSearchResponse, err error) {
    return c.DescribeWhiteboardPushSearchWithContext(context.Background(), request)
}

// DescribeWhiteboardPushSearch
// 根据房间号搜索白板推流任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DescribeWhiteboardPushSearchWithContext(ctx context.Context, request *DescribeWhiteboardPushSearchRequest) (response *DescribeWhiteboardPushSearchResponse, err error) {
    if request == nil {
        request = NewDescribeWhiteboardPushSearchRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWhiteboardPushSearch require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWhiteboardPushSearchResponse()
    err = c.Send(request, response)
    return
}

func NewModifyApplicationRequest() (request *ModifyApplicationRequest) {
    request = &ModifyApplicationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "ModifyApplication")
    
    
    return
}

func NewModifyApplicationResponse() (response *ModifyApplicationResponse) {
    response = &ModifyApplicationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyApplication
// 修改白板应用
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) ModifyApplication(request *ModifyApplicationRequest) (response *ModifyApplicationResponse, err error) {
    return c.ModifyApplicationWithContext(context.Background(), request)
}

// ModifyApplication
// 修改白板应用
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) ModifyApplicationWithContext(ctx context.Context, request *ModifyApplicationRequest) (response *ModifyApplicationResponse, err error) {
    if request == nil {
        request = NewModifyApplicationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyApplication require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyApplicationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAutoRenewFlagRequest() (request *ModifyAutoRenewFlagRequest) {
    request = &ModifyAutoRenewFlagRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "ModifyAutoRenewFlag")
    
    
    return
}

func NewModifyAutoRenewFlagResponse() (response *ModifyAutoRenewFlagResponse) {
    response = &ModifyAutoRenewFlagResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyAutoRenewFlag
// 设置白板月功能费自动续费
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) ModifyAutoRenewFlag(request *ModifyAutoRenewFlagRequest) (response *ModifyAutoRenewFlagResponse, err error) {
    return c.ModifyAutoRenewFlagWithContext(context.Background(), request)
}

// ModifyAutoRenewFlag
// 设置白板月功能费自动续费
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
func (c *Client) ModifyAutoRenewFlagWithContext(ctx context.Context, request *ModifyAutoRenewFlagRequest) (response *ModifyAutoRenewFlagResponse, err error) {
    if request == nil {
        request = NewModifyAutoRenewFlagRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyAutoRenewFlag require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyAutoRenewFlagResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWhiteboardApplicationConfigRequest() (request *ModifyWhiteboardApplicationConfigRequest) {
    request = &ModifyWhiteboardApplicationConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "ModifyWhiteboardApplicationConfig")
    
    
    return
}

func NewModifyWhiteboardApplicationConfigResponse() (response *ModifyWhiteboardApplicationConfigResponse) {
    response = &ModifyWhiteboardApplicationConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWhiteboardApplicationConfig
// 修改白板应用任务相关的配置，包括存储桶、回调等
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COSBUCKETINVALID = "FailedOperation.CosBucketInvalid"
//  FAILEDOPERATION_GETCREDENTIALFAIL = "FailedOperation.GetCredentialFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  INVALIDPARAMETER_CDNDOMAINNOTFOUND = "InvalidParameter.CdnDomainNotFound"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) ModifyWhiteboardApplicationConfig(request *ModifyWhiteboardApplicationConfigRequest) (response *ModifyWhiteboardApplicationConfigResponse, err error) {
    return c.ModifyWhiteboardApplicationConfigWithContext(context.Background(), request)
}

// ModifyWhiteboardApplicationConfig
// 修改白板应用任务相关的配置，包括存储桶、回调等
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_COSBUCKETINVALID = "FailedOperation.CosBucketInvalid"
//  FAILEDOPERATION_GETCREDENTIALFAIL = "FailedOperation.GetCredentialFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  INVALIDPARAMETER_CDNDOMAINNOTFOUND = "InvalidParameter.CdnDomainNotFound"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  INVALIDPARAMETER_UNMARSHALJSONBODYFAIL = "InvalidParameter.UnmarshalJSONBodyFail"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) ModifyWhiteboardApplicationConfigWithContext(ctx context.Context, request *ModifyWhiteboardApplicationConfigRequest) (response *ModifyWhiteboardApplicationConfigResponse, err error) {
    if request == nil {
        request = NewModifyWhiteboardApplicationConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWhiteboardApplicationConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWhiteboardApplicationConfigResponse()
    err = c.Send(request, response)
    return
}

func NewModifyWhiteboardBucketConfigRequest() (request *ModifyWhiteboardBucketConfigRequest) {
    request = &ModifyWhiteboardBucketConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "ModifyWhiteboardBucketConfig")
    
    
    return
}

func NewModifyWhiteboardBucketConfigResponse() (response *ModifyWhiteboardBucketConfigResponse) {
    response = &ModifyWhiteboardBucketConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyWhiteboardBucketConfig
// 设置文档转码，实时录制存储桶的配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_COSBUCKETINVALID = "FailedOperation.CosBucketInvalid"
//  FAILEDOPERATION_GETCREDENTIALFAIL = "FailedOperation.GetCredentialFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CDNDOMAINNOTFOUND = "InvalidParameter.CdnDomainNotFound"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) ModifyWhiteboardBucketConfig(request *ModifyWhiteboardBucketConfigRequest) (response *ModifyWhiteboardBucketConfigResponse, err error) {
    return c.ModifyWhiteboardBucketConfigWithContext(context.Background(), request)
}

// ModifyWhiteboardBucketConfig
// 设置文档转码，实时录制存储桶的配置
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_COSBUCKETINVALID = "FailedOperation.CosBucketInvalid"
//  FAILEDOPERATION_GETCREDENTIALFAIL = "FailedOperation.GetCredentialFail"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CDNDOMAINNOTFOUND = "InvalidParameter.CdnDomainNotFound"
//  INVALIDPARAMETER_INVALIDTASKTYPE = "InvalidParameter.InvalidTaskType"
//  INVALIDPARAMETER_SDKAPPIDNOTFOUND = "InvalidParameter.SdkAppIdNotFound"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) ModifyWhiteboardBucketConfigWithContext(ctx context.Context, request *ModifyWhiteboardBucketConfigRequest) (response *ModifyWhiteboardBucketConfigResponse, err error) {
    if request == nil {
        request = NewModifyWhiteboardBucketConfigRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyWhiteboardBucketConfig require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyWhiteboardBucketConfigResponse()
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

func NewSetOfflineRecordCallbackRequest() (request *SetOfflineRecordCallbackRequest) {
    request = &SetOfflineRecordCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetOfflineRecordCallback")
    
    
    return
}

func NewSetOfflineRecordCallbackResponse() (response *SetOfflineRecordCallbackResponse) {
    response = &SetOfflineRecordCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetOfflineRecordCallback
// 课后录制服务已下线
//
// 
//
// 设置课后录制回调地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetOfflineRecordCallback(request *SetOfflineRecordCallbackRequest) (response *SetOfflineRecordCallbackResponse, err error) {
    return c.SetOfflineRecordCallbackWithContext(context.Background(), request)
}

// SetOfflineRecordCallback
// 课后录制服务已下线
//
// 
//
// 设置课后录制回调地址
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetOfflineRecordCallbackWithContext(ctx context.Context, request *SetOfflineRecordCallbackRequest) (response *SetOfflineRecordCallbackResponse, err error) {
    if request == nil {
        request = NewSetOfflineRecordCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetOfflineRecordCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetOfflineRecordCallbackResponse()
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

func NewSetPPTCheckCallbackRequest() (request *SetPPTCheckCallbackRequest) {
    request = &SetPPTCheckCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetPPTCheckCallback")
    
    
    return
}

func NewSetPPTCheckCallbackResponse() (response *SetPPTCheckCallbackResponse) {
    response = &SetPPTCheckCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetPPTCheckCallback
// 设置PPT检测任务回调地址，回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40260#c9cbe05f-fe1a-4410-b4dc-40cc301c7b81
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetPPTCheckCallback(request *SetPPTCheckCallbackRequest) (response *SetPPTCheckCallbackResponse, err error) {
    return c.SetPPTCheckCallbackWithContext(context.Background(), request)
}

// SetPPTCheckCallback
// 设置PPT检测任务回调地址，回调数据格式请参考文档：https://cloud.tencent.com/document/product/1137/40260#c9cbe05f-fe1a-4410-b4dc-40cc301c7b81
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetPPTCheckCallbackWithContext(ctx context.Context, request *SetPPTCheckCallbackRequest) (response *SetPPTCheckCallbackResponse, err error) {
    if request == nil {
        request = NewSetPPTCheckCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetPPTCheckCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetPPTCheckCallbackResponse()
    err = c.Send(request, response)
    return
}

func NewSetPPTCheckCallbackKeyRequest() (request *SetPPTCheckCallbackKeyRequest) {
    request = &SetPPTCheckCallbackKeyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetPPTCheckCallbackKey")
    
    
    return
}

func NewSetPPTCheckCallbackKeyResponse() (response *SetPPTCheckCallbackKeyResponse) {
    response = &SetPPTCheckCallbackKeyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetPPTCheckCallbackKey
// 设置PPT检测任务回调密钥，回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetPPTCheckCallbackKey(request *SetPPTCheckCallbackKeyRequest) (response *SetPPTCheckCallbackKeyResponse, err error) {
    return c.SetPPTCheckCallbackKeyWithContext(context.Background(), request)
}

// SetPPTCheckCallbackKey
// 设置PPT检测任务回调密钥，回调鉴权方式请参考文档：https://cloud.tencent.com/document/product/1137/40257
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetPPTCheckCallbackKeyWithContext(ctx context.Context, request *SetPPTCheckCallbackKeyRequest) (response *SetPPTCheckCallbackKeyResponse, err error) {
    if request == nil {
        request = NewSetPPTCheckCallbackKeyRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetPPTCheckCallbackKey require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetPPTCheckCallbackKeyResponse()
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

func NewSetWarningCallbackRequest() (request *SetWarningCallbackRequest) {
    request = &SetWarningCallbackRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiw", APIVersion, "SetWarningCallback")
    
    
    return
}

func NewSetWarningCallbackResponse() (response *SetWarningCallbackResponse) {
    response = &SetWarningCallbackResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SetWarningCallback
// 设置告警回调地址。此功能需要申请白名单使用。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetWarningCallback(request *SetWarningCallbackRequest) (response *SetWarningCallbackResponse, err error) {
    return c.SetWarningCallbackWithContext(context.Background(), request)
}

// SetWarningCallback
// 设置告警回调地址。此功能需要申请白名单使用。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  INVALIDPARAMETER_BODYPARAMETERTYPEUNMATCHED = "InvalidParameter.BodyParameterTypeUnmatched"
//  INVALIDPARAMETER_CALLBACKADDRESSFORMATERROR = "InvalidParameter.CallbackAddressFormatError"
//  RESOURCEUNAVAILABLE_NOTREGISTERED = "ResourceUnavailable.NotRegistered"
//  RESOURCEUNAVAILABLE_SERVICEEXPIRED = "ResourceUnavailable.ServiceExpired"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) SetWarningCallbackWithContext(ctx context.Context, request *SetWarningCallbackRequest) (response *SetWarningCallbackResponse, err error) {
    if request == nil {
        request = NewSetWarningCallbackRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetWarningCallback require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetWarningCallbackResponse()
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
