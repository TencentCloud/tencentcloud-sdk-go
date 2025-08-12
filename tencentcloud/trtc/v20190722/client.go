// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20190722

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-07-22"

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


func NewControlAIConversationRequest() (request *ControlAIConversationRequest) {
    request = &ControlAIConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "ControlAIConversation")
    
    
    return
}

func NewControlAIConversationResponse() (response *ControlAIConversationResponse) {
    response = &ControlAIConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ControlAIConversation
// 提供服务端控制机器人的功能
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) ControlAIConversation(request *ControlAIConversationRequest) (response *ControlAIConversationResponse, err error) {
    return c.ControlAIConversationWithContext(context.Background(), request)
}

// ControlAIConversation
// 提供服务端控制机器人的功能
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) ControlAIConversationWithContext(ctx context.Context, request *ControlAIConversationRequest) (response *ControlAIConversationResponse, err error) {
    if request == nil {
        request = NewControlAIConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "ControlAIConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ControlAIConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewControlAIConversationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBasicModerationRequest() (request *CreateBasicModerationRequest) {
    request = &CreateBasicModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "CreateBasicModeration")
    
    
    return
}

func NewCreateBasicModerationResponse() (response *CreateBasicModerationResponse) {
    response = &CreateBasicModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateBasicModeration
// 接口说明：
//
// 启动终端审核功能，完成房间内的音频审核。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ACCESSKEY = "MissingParameter.AccessKey"
//  MISSINGPARAMETER_BUCKET = "MissingParameter.Bucket"
//  MISSINGPARAMETER_REGION = "MissingParameter.Region"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_SECRETKEY = "MissingParameter.SecretKey"
//  MISSINGPARAMETER_SLICEPARAMS = "MissingParameter.SliceParams"
//  MISSINGPARAMETER_SLICESTORAGEPARAMS = "MissingParameter.SliceStorageParams"
//  MISSINGPARAMETER_SLICETYPE = "MissingParameter.SliceType"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  MISSINGPARAMETER_USERSIG = "MissingParameter.UserSig"
//  MISSINGPARAMETER_VENDOR = "MissingParameter.Vendor"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBasicModeration(request *CreateBasicModerationRequest) (response *CreateBasicModerationResponse, err error) {
    return c.CreateBasicModerationWithContext(context.Background(), request)
}

// CreateBasicModeration
// 接口说明：
//
// 启动终端审核功能，完成房间内的音频审核。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ACCESSKEY = "MissingParameter.AccessKey"
//  MISSINGPARAMETER_BUCKET = "MissingParameter.Bucket"
//  MISSINGPARAMETER_REGION = "MissingParameter.Region"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_SECRETKEY = "MissingParameter.SecretKey"
//  MISSINGPARAMETER_SLICEPARAMS = "MissingParameter.SliceParams"
//  MISSINGPARAMETER_SLICESTORAGEPARAMS = "MissingParameter.SliceStorageParams"
//  MISSINGPARAMETER_SLICETYPE = "MissingParameter.SliceType"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  MISSINGPARAMETER_USERSIG = "MissingParameter.UserSig"
//  MISSINGPARAMETER_VENDOR = "MissingParameter.Vendor"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateBasicModerationWithContext(ctx context.Context, request *CreateBasicModerationRequest) (response *CreateBasicModerationResponse, err error) {
    if request == nil {
        request = NewCreateBasicModerationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "CreateBasicModeration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateBasicModeration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateBasicModerationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudModerationRequest() (request *CreateCloudModerationRequest) {
    request = &CreateCloudModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "CreateCloudModeration")
    
    
    return
}

func NewCreateCloudModerationResponse() (response *CreateCloudModerationResponse) {
    response = &CreateCloudModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudModeration
// 接口说明：
//
// 启动云端审核功能，完成房间内的音视频切片，视频截帧，或者录制音频流，送审到指定的审核商，完成审核。
//
// 
//
// 您可以通过此接口实现如下目标：
//
// * 指定审核参数（ModerationParams）来指定审核需要的详细参数。
//
// * 指定存储参数（ModerationStorageParams）将命中的审核文件指定上传到您希望的云存储，目前支持腾讯云（对象存储COS）和第三方AWS
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ACCESSKEY = "MissingParameter.AccessKey"
//  MISSINGPARAMETER_BUCKET = "MissingParameter.Bucket"
//  MISSINGPARAMETER_REGION = "MissingParameter.Region"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_SECRETKEY = "MissingParameter.SecretKey"
//  MISSINGPARAMETER_SLICEPARAMS = "MissingParameter.SliceParams"
//  MISSINGPARAMETER_SLICESTORAGEPARAMS = "MissingParameter.SliceStorageParams"
//  MISSINGPARAMETER_SLICETYPE = "MissingParameter.SliceType"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  MISSINGPARAMETER_USERSIG = "MissingParameter.UserSig"
//  MISSINGPARAMETER_VENDOR = "MissingParameter.Vendor"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloudModeration(request *CreateCloudModerationRequest) (response *CreateCloudModerationResponse, err error) {
    return c.CreateCloudModerationWithContext(context.Background(), request)
}

// CreateCloudModeration
// 接口说明：
//
// 启动云端审核功能，完成房间内的音视频切片，视频截帧，或者录制音频流，送审到指定的审核商，完成审核。
//
// 
//
// 您可以通过此接口实现如下目标：
//
// * 指定审核参数（ModerationParams）来指定审核需要的详细参数。
//
// * 指定存储参数（ModerationStorageParams）将命中的审核文件指定上传到您希望的云存储，目前支持腾讯云（对象存储COS）和第三方AWS
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ACCESSKEY = "MissingParameter.AccessKey"
//  MISSINGPARAMETER_BUCKET = "MissingParameter.Bucket"
//  MISSINGPARAMETER_REGION = "MissingParameter.Region"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_SECRETKEY = "MissingParameter.SecretKey"
//  MISSINGPARAMETER_SLICEPARAMS = "MissingParameter.SliceParams"
//  MISSINGPARAMETER_SLICESTORAGEPARAMS = "MissingParameter.SliceStorageParams"
//  MISSINGPARAMETER_SLICETYPE = "MissingParameter.SliceType"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  MISSINGPARAMETER_USERSIG = "MissingParameter.UserSig"
//  MISSINGPARAMETER_VENDOR = "MissingParameter.Vendor"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloudModerationWithContext(ctx context.Context, request *CreateCloudModerationRequest) (response *CreateCloudModerationResponse, err error) {
    if request == nil {
        request = NewCreateCloudModerationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "CreateCloudModeration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudModeration require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudModerationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudRecordingRequest() (request *CreateCloudRecordingRequest) {
    request = &CreateCloudRecordingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "CreateCloudRecording")
    
    
    return
}

func NewCreateCloudRecordingResponse() (response *CreateCloudRecordingResponse) {
    response = &CreateCloudRecordingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudRecording
// 接口说明：
//
// 启动云端录制功能，完成房间内的音视频录制，并上传到指定的云存储。您可以通过此 API 接口把TRTC 房间中的每一路音视频流做单独的录制又或者多路视频画面合流混成一路。
//
// 
//
// 您可以通过此接口实现如下目标：
//
// * 指定订阅流参数（RecordParams）来指定需要录制的主播的黑名单或者白名单。
//
// * 指定录制存储参数（StorageParams）来指定上传到您希望的云存储，目前支持腾讯云（云点播VOD、对象存储COS）和第三方AWS
//
// * 指定合流模式下的音视频转码详细参数（MixTranscodeParams），包括视频分辨率、视频码率、视频帧率、以及声音质量等
//
// * 指定合流模式各路画面的位置和布局或者也可以指定自动模板的方式来配置。
//
// 
//
// 关键名词：
//
// * 单流录制：分别录制房间的订阅UserId的音频和视频，录制服务会实时将录制文件上传至您指定的云存储。
//
// * 合流录制：将房间内订阅UserId的音视频混录成一个视频文件，并将录制文件上传至您指定的云存储。（录制结束后可前往云点播控制台https://console.cloud.tencent.com/vod/media 或 对象存储COS控制台https://console.cloud.tencent.com/cos/bucket查看文件）。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ACCESSKEY = "MissingParameter.AccessKey"
//  MISSINGPARAMETER_BUCKET = "MissingParameter.Bucket"
//  MISSINGPARAMETER_CLOUDSTORAGE = "MissingParameter.CloudStorage"
//  MISSINGPARAMETER_RECORDMODE = "MissingParameter.RecordMode"
//  MISSINGPARAMETER_RECORDPARAMS = "MissingParameter.RecordParams"
//  MISSINGPARAMETER_REGION = "MissingParameter.Region"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_SECRETKEY = "MissingParameter.SecretKey"
//  MISSINGPARAMETER_STORAGEPARAMS = "MissingParameter.StorageParams"
//  MISSINGPARAMETER_STREAMTYPE = "MissingParameter.StreamType"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  MISSINGPARAMETER_USERSIG = "MissingParameter.UserSig"
//  MISSINGPARAMETER_VENDOR = "MissingParameter.Vendor"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloudRecording(request *CreateCloudRecordingRequest) (response *CreateCloudRecordingResponse, err error) {
    return c.CreateCloudRecordingWithContext(context.Background(), request)
}

// CreateCloudRecording
// 接口说明：
//
// 启动云端录制功能，完成房间内的音视频录制，并上传到指定的云存储。您可以通过此 API 接口把TRTC 房间中的每一路音视频流做单独的录制又或者多路视频画面合流混成一路。
//
// 
//
// 您可以通过此接口实现如下目标：
//
// * 指定订阅流参数（RecordParams）来指定需要录制的主播的黑名单或者白名单。
//
// * 指定录制存储参数（StorageParams）来指定上传到您希望的云存储，目前支持腾讯云（云点播VOD、对象存储COS）和第三方AWS
//
// * 指定合流模式下的音视频转码详细参数（MixTranscodeParams），包括视频分辨率、视频码率、视频帧率、以及声音质量等
//
// * 指定合流模式各路画面的位置和布局或者也可以指定自动模板的方式来配置。
//
// 
//
// 关键名词：
//
// * 单流录制：分别录制房间的订阅UserId的音频和视频，录制服务会实时将录制文件上传至您指定的云存储。
//
// * 合流录制：将房间内订阅UserId的音视频混录成一个视频文件，并将录制文件上传至您指定的云存储。（录制结束后可前往云点播控制台https://console.cloud.tencent.com/vod/media 或 对象存储COS控制台https://console.cloud.tencent.com/cos/bucket查看文件）。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ACCESSKEY = "MissingParameter.AccessKey"
//  MISSINGPARAMETER_BUCKET = "MissingParameter.Bucket"
//  MISSINGPARAMETER_CLOUDSTORAGE = "MissingParameter.CloudStorage"
//  MISSINGPARAMETER_RECORDMODE = "MissingParameter.RecordMode"
//  MISSINGPARAMETER_RECORDPARAMS = "MissingParameter.RecordParams"
//  MISSINGPARAMETER_REGION = "MissingParameter.Region"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_SECRETKEY = "MissingParameter.SecretKey"
//  MISSINGPARAMETER_STORAGEPARAMS = "MissingParameter.StorageParams"
//  MISSINGPARAMETER_STREAMTYPE = "MissingParameter.StreamType"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  MISSINGPARAMETER_USERSIG = "MissingParameter.UserSig"
//  MISSINGPARAMETER_VENDOR = "MissingParameter.Vendor"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloudRecordingWithContext(ctx context.Context, request *CreateCloudRecordingRequest) (response *CreateCloudRecordingResponse, err error) {
    if request == nil {
        request = NewCreateCloudRecordingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "CreateCloudRecording")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudRecordingResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCloudSliceTaskRequest() (request *CreateCloudSliceTaskRequest) {
    request = &CreateCloudSliceTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "CreateCloudSliceTask")
    
    
    return
}

func NewCreateCloudSliceTaskResponse() (response *CreateCloudSliceTaskResponse) {
    response = &CreateCloudSliceTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateCloudSliceTask
// 接口说明：
//
// 启动云端切片功能，完成房间内的音视频切片，并上传到指定的云存储。
//
// 
//
// 您可以通过此接口实现如下目标：
//
// * 指定切片参数（SliceParams）来指定需要切片的主播的黑名单或者白名单。
//
// * 指定存储参数（SliceStorageParams）来指定上传到您希望的云存储，目前支持腾讯云（对象存储COS）和第三方AWS
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ACCESSKEY = "MissingParameter.AccessKey"
//  MISSINGPARAMETER_BUCKET = "MissingParameter.Bucket"
//  MISSINGPARAMETER_REGION = "MissingParameter.Region"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_SECRETKEY = "MissingParameter.SecretKey"
//  MISSINGPARAMETER_SLICEPARAMS = "MissingParameter.SliceParams"
//  MISSINGPARAMETER_SLICESTORAGEPARAMS = "MissingParameter.SliceStorageParams"
//  MISSINGPARAMETER_SLICETYPE = "MissingParameter.SliceType"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  MISSINGPARAMETER_USERSIG = "MissingParameter.UserSig"
//  MISSINGPARAMETER_VENDOR = "MissingParameter.Vendor"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloudSliceTask(request *CreateCloudSliceTaskRequest) (response *CreateCloudSliceTaskResponse, err error) {
    return c.CreateCloudSliceTaskWithContext(context.Background(), request)
}

// CreateCloudSliceTask
// 接口说明：
//
// 启动云端切片功能，完成房间内的音视频切片，并上传到指定的云存储。
//
// 
//
// 您可以通过此接口实现如下目标：
//
// * 指定切片参数（SliceParams）来指定需要切片的主播的黑名单或者白名单。
//
// * 指定存储参数（SliceStorageParams）来指定上传到您希望的云存储，目前支持腾讯云（对象存储COS）和第三方AWS
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ACCESSKEY = "MissingParameter.AccessKey"
//  MISSINGPARAMETER_BUCKET = "MissingParameter.Bucket"
//  MISSINGPARAMETER_REGION = "MissingParameter.Region"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_SECRETKEY = "MissingParameter.SecretKey"
//  MISSINGPARAMETER_SLICEPARAMS = "MissingParameter.SliceParams"
//  MISSINGPARAMETER_SLICESTORAGEPARAMS = "MissingParameter.SliceStorageParams"
//  MISSINGPARAMETER_SLICETYPE = "MissingParameter.SliceType"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  MISSINGPARAMETER_USERSIG = "MissingParameter.UserSig"
//  MISSINGPARAMETER_VENDOR = "MissingParameter.Vendor"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateCloudSliceTaskWithContext(ctx context.Context, request *CreateCloudSliceTaskRequest) (response *CreateCloudSliceTaskResponse, err error) {
    if request == nil {
        request = NewCreateCloudSliceTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "CreateCloudSliceTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudSliceTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudSliceTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePictureRequest() (request *CreatePictureRequest) {
    request = &CreatePictureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "CreatePicture")
    
    
    return
}

func NewCreatePictureResponse() (response *CreatePictureResponse) {
    response = &CreatePictureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreatePicture
// 如果您需要在 [云端混流转码](https://cloud.tencent.com/document/product/647/16827) 时频繁新增自定义背景图或水印，可通过此接口上传新的图片素材。无需频繁新增图片的场景，建议直接在 [控制台 > 应用管理 > 素材管理](https://cloud.tencent.com/document/product/647/50769) 中操作。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_CHECKCONTENTFAILED = "InvalidParameter.CheckContentFailed"
//  INVALIDPARAMETER_CHECKSUFFIXFAILED = "InvalidParameter.CheckSuffixFailed"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreatePicture(request *CreatePictureRequest) (response *CreatePictureResponse, err error) {
    return c.CreatePictureWithContext(context.Background(), request)
}

// CreatePicture
// 如果您需要在 [云端混流转码](https://cloud.tencent.com/document/product/647/16827) 时频繁新增自定义背景图或水印，可通过此接口上传新的图片素材。无需频繁新增图片的场景，建议直接在 [控制台 > 应用管理 > 素材管理](https://cloud.tencent.com/document/product/647/50769) 中操作。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_CHECKCONTENTFAILED = "InvalidParameter.CheckContentFailed"
//  INVALIDPARAMETER_CHECKSUFFIXFAILED = "InvalidParameter.CheckSuffixFailed"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreatePictureWithContext(ctx context.Context, request *CreatePictureRequest) (response *CreatePictureResponse, err error) {
    if request == nil {
        request = NewCreatePictureRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "CreatePicture")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePicture require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePictureResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBasicModerationRequest() (request *DeleteBasicModerationRequest) {
    request = &DeleteBasicModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DeleteBasicModeration")
    
    
    return
}

func NewDeleteBasicModerationResponse() (response *DeleteBasicModerationResponse) {
    response = &DeleteBasicModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteBasicModeration
// 成功开启审核任务后，可以使用此接口来停止任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBasicModeration(request *DeleteBasicModerationRequest) (response *DeleteBasicModerationResponse, err error) {
    return c.DeleteBasicModerationWithContext(context.Background(), request)
}

// DeleteBasicModeration
// 成功开启审核任务后，可以使用此接口来停止任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteBasicModerationWithContext(ctx context.Context, request *DeleteBasicModerationRequest) (response *DeleteBasicModerationResponse, err error) {
    if request == nil {
        request = NewDeleteBasicModerationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DeleteBasicModeration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteBasicModeration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteBasicModerationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudModerationRequest() (request *DeleteCloudModerationRequest) {
    request = &DeleteCloudModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DeleteCloudModeration")
    
    
    return
}

func NewDeleteCloudModerationResponse() (response *DeleteCloudModerationResponse) {
    response = &DeleteCloudModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudModeration
// 成功开启云端审核任务后，可以使用此接口来停止送审。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCloudModeration(request *DeleteCloudModerationRequest) (response *DeleteCloudModerationResponse, err error) {
    return c.DeleteCloudModerationWithContext(context.Background(), request)
}

// DeleteCloudModeration
// 成功开启云端审核任务后，可以使用此接口来停止送审。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCloudModerationWithContext(ctx context.Context, request *DeleteCloudModerationRequest) (response *DeleteCloudModerationResponse, err error) {
    if request == nil {
        request = NewDeleteCloudModerationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DeleteCloudModeration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudModeration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudModerationResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudRecordingRequest() (request *DeleteCloudRecordingRequest) {
    request = &DeleteCloudRecordingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DeleteCloudRecording")
    
    
    return
}

func NewDeleteCloudRecordingResponse() (response *DeleteCloudRecordingResponse) {
    response = &DeleteCloudRecordingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudRecording
// 成功开启录制后，可以使用此接口来停止录制任务。停止录制成功后不代表文件全部传输完成，如果未完成后台将会继续上传文件，成功后通过事件回调通知客户文件全部传输完成状态。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCloudRecording(request *DeleteCloudRecordingRequest) (response *DeleteCloudRecordingResponse, err error) {
    return c.DeleteCloudRecordingWithContext(context.Background(), request)
}

// DeleteCloudRecording
// 成功开启录制后，可以使用此接口来停止录制任务。停止录制成功后不代表文件全部传输完成，如果未完成后台将会继续上传文件，成功后通过事件回调通知客户文件全部传输完成状态。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCloudRecordingWithContext(ctx context.Context, request *DeleteCloudRecordingRequest) (response *DeleteCloudRecordingResponse, err error) {
    if request == nil {
        request = NewDeleteCloudRecordingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DeleteCloudRecording")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudRecordingResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCloudSliceTaskRequest() (request *DeleteCloudSliceTaskRequest) {
    request = &DeleteCloudSliceTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DeleteCloudSliceTask")
    
    
    return
}

func NewDeleteCloudSliceTaskResponse() (response *DeleteCloudSliceTaskResponse) {
    response = &DeleteCloudSliceTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteCloudSliceTask
// 成功开启切片任务后，可以使用此接口来停止任务。停止切片成功后不代表文件全部传输完成，如果未完成后台将会继续上传文件，成功后通过事件回调通知客户文件全部传输完成状态。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCloudSliceTask(request *DeleteCloudSliceTaskRequest) (response *DeleteCloudSliceTaskResponse, err error) {
    return c.DeleteCloudSliceTaskWithContext(context.Background(), request)
}

// DeleteCloudSliceTask
// 成功开启切片任务后，可以使用此接口来停止任务。停止切片成功后不代表文件全部传输完成，如果未完成后台将会继续上传文件，成功后通过事件回调通知客户文件全部传输完成状态。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DeleteCloudSliceTaskWithContext(ctx context.Context, request *DeleteCloudSliceTaskRequest) (response *DeleteCloudSliceTaskResponse, err error) {
    if request == nil {
        request = NewDeleteCloudSliceTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DeleteCloudSliceTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudSliceTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudSliceTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePictureRequest() (request *DeletePictureRequest) {
    request = &DeletePictureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DeletePicture")
    
    
    return
}

func NewDeletePictureResponse() (response *DeletePictureResponse) {
    response = &DeletePictureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeletePicture
// 如果您需要在 [云端混流转码](https://cloud.tencent.com/document/product/647/16827) 时频繁删除自定义背景图或水印，可通过此接口删除已上传的图片。无需频繁删除图片的场景，建议直接在 [控制台 > 应用管理 > 素材管理](https://cloud.tencent.com/document/product/647/50769) 中操作。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PICTURENOTFOUND = "InvalidParameter.PictureNotFound"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeletePicture(request *DeletePictureRequest) (response *DeletePictureResponse, err error) {
    return c.DeletePictureWithContext(context.Background(), request)
}

// DeletePicture
// 如果您需要在 [云端混流转码](https://cloud.tencent.com/document/product/647/16827) 时频繁删除自定义背景图或水印，可通过此接口删除已上传的图片。无需频繁删除图片的场景，建议直接在 [控制台 > 应用管理 > 素材管理](https://cloud.tencent.com/document/product/647/50769) 中操作。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_PICTURENOTFOUND = "InvalidParameter.PictureNotFound"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DeletePictureWithContext(ctx context.Context, request *DeletePictureRequest) (response *DeletePictureResponse, err error) {
    if request == nil {
        request = NewDeletePictureRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DeletePicture")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePicture require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePictureResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteVoicePrintRequest() (request *DeleteVoicePrintRequest) {
    request = &DeleteVoicePrintRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DeleteVoicePrint")
    
    
    return
}

func NewDeleteVoicePrintResponse() (response *DeleteVoicePrintResponse) {
    response = &DeleteVoicePrintResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteVoicePrint
// 传入声纹ID，删除之前注册的声纹信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONFAILED = "FailedOperation.DbOperationFailed"
//  FAILEDOPERATION_UPDATEVOICEPRINTIDNOTFOUND = "FailedOperation.UpdateVoicePrintIdNotFound"
func (c *Client) DeleteVoicePrint(request *DeleteVoicePrintRequest) (response *DeleteVoicePrintResponse, err error) {
    return c.DeleteVoicePrintWithContext(context.Background(), request)
}

// DeleteVoicePrint
// 传入声纹ID，删除之前注册的声纹信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONFAILED = "FailedOperation.DbOperationFailed"
//  FAILEDOPERATION_UPDATEVOICEPRINTIDNOTFOUND = "FailedOperation.UpdateVoicePrintIdNotFound"
func (c *Client) DeleteVoicePrintWithContext(ctx context.Context, request *DeleteVoicePrintRequest) (response *DeleteVoicePrintResponse, err error) {
    if request == nil {
        request = NewDeleteVoicePrintRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DeleteVoicePrint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteVoicePrint require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteVoicePrintResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAIConversationRequest() (request *DescribeAIConversationRequest) {
    request = &DescribeAIConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeAIConversation")
    
    
    return
}

func NewDescribeAIConversationResponse() (response *DescribeAIConversationResponse) {
    response = &DescribeAIConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAIConversation
// 查询AI对话任务状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) DescribeAIConversation(request *DescribeAIConversationRequest) (response *DescribeAIConversationResponse, err error) {
    return c.DescribeAIConversationWithContext(context.Background(), request)
}

// DescribeAIConversation
// 查询AI对话任务状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) DescribeAIConversationWithContext(ctx context.Context, request *DescribeAIConversationRequest) (response *DescribeAIConversationResponse, err error) {
    if request == nil {
        request = NewDescribeAIConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeAIConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAIConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAIConversationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAITranscriptionRequest() (request *DescribeAITranscriptionRequest) {
    request = &DescribeAITranscriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeAITranscription")
    
    
    return
}

func NewDescribeAITranscriptionResponse() (response *DescribeAITranscriptionResponse) {
    response = &DescribeAITranscriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAITranscription
// 查询AI转录任务状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) DescribeAITranscription(request *DescribeAITranscriptionRequest) (response *DescribeAITranscriptionResponse, err error) {
    return c.DescribeAITranscriptionWithContext(context.Background(), request)
}

// DescribeAITranscription
// 查询AI转录任务状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) DescribeAITranscriptionWithContext(ctx context.Context, request *DescribeAITranscriptionRequest) (response *DescribeAITranscriptionResponse, err error) {
    if request == nil {
        request = NewDescribeAITranscriptionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeAITranscription")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAITranscription require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAITranscriptionResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCallDetailInfoRequest() (request *DescribeCallDetailInfoRequest) {
    request = &DescribeCallDetailInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeCallDetailInfo")
    
    
    return
}

func NewDescribeCallDetailInfoResponse() (response *DescribeCallDetailInfoResponse) {
    response = &DescribeCallDetailInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCallDetailInfo
// 查询指定时间内的用户列表及用户通话质量数据，最大可查询14天内数据。DataType 不为null，查询起止时间不超过1个小时，查询用户不超过6个，支持跨天查询。DataType为null时，查询起止时间不超过4个小时， 默认查询6个用户，同时支持每页查询100以内用户个数（PageSize不超过100）。接口用于查询质量问题，不推荐作为计费使用。（同老接口DescribeCallDetail）
//
// **注意**：
//
// 1.该接口只用于历史数据统计或核对数据使用，实时类关键业务逻辑不能使用。
//
// 2.该接口自2024年4月1日起正式商业化，需订阅套餐解锁调用能力，提供以下两种解锁方式，可任选其一解锁：
//
// 方式一：通过订阅[包月套餐](https://cloud.tencent.com/document/product/647/85386)「尊享版」（可查近7天）和「旗舰版」（可查近14天），[前往订阅](https://buy.cloud.tencent.com/trtc?trtcversion=top)。
//
// 方式二：通过订阅[监控仪表盘](https://cloud.tencent.com/document/product/647/81331)商业套餐包「基础版」（可查近7天）和「进阶版」（可查近14天），[前往订阅](https://buy.cloud.tencent.com/trtc_monitor)。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_PAGENUMBER = "InvalidParameter.PageNumber"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIMEOVERSIZE = "InvalidParameter.StartTimeOversize"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_USERIDSMORETHANSIX = "InvalidParameter.UserIdsMorethanSix"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeCallDetailInfo(request *DescribeCallDetailInfoRequest) (response *DescribeCallDetailInfoResponse, err error) {
    return c.DescribeCallDetailInfoWithContext(context.Background(), request)
}

// DescribeCallDetailInfo
// 查询指定时间内的用户列表及用户通话质量数据，最大可查询14天内数据。DataType 不为null，查询起止时间不超过1个小时，查询用户不超过6个，支持跨天查询。DataType为null时，查询起止时间不超过4个小时， 默认查询6个用户，同时支持每页查询100以内用户个数（PageSize不超过100）。接口用于查询质量问题，不推荐作为计费使用。（同老接口DescribeCallDetail）
//
// **注意**：
//
// 1.该接口只用于历史数据统计或核对数据使用，实时类关键业务逻辑不能使用。
//
// 2.该接口自2024年4月1日起正式商业化，需订阅套餐解锁调用能力，提供以下两种解锁方式，可任选其一解锁：
//
// 方式一：通过订阅[包月套餐](https://cloud.tencent.com/document/product/647/85386)「尊享版」（可查近7天）和「旗舰版」（可查近14天），[前往订阅](https://buy.cloud.tencent.com/trtc?trtcversion=top)。
//
// 方式二：通过订阅[监控仪表盘](https://cloud.tencent.com/document/product/647/81331)商业套餐包「基础版」（可查近7天）和「进阶版」（可查近14天），[前往订阅](https://buy.cloud.tencent.com/trtc_monitor)。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_PAGENUMBER = "InvalidParameter.PageNumber"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIMEOVERSIZE = "InvalidParameter.StartTimeOversize"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_USERIDSMORETHANSIX = "InvalidParameter.UserIdsMorethanSix"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeCallDetailInfoWithContext(ctx context.Context, request *DescribeCallDetailInfoRequest) (response *DescribeCallDetailInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCallDetailInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeCallDetailInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCallDetailInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCallDetailInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudModerationRequest() (request *DescribeCloudModerationRequest) {
    request = &DescribeCloudModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeCloudModeration")
    
    
    return
}

func NewDescribeCloudModerationResponse() (response *DescribeCloudModerationResponse) {
    response = &DescribeCloudModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudModeration
// 成功开启审核任务后，可以使用此接口来查询审核任务状态和订阅的黑白名单信息。仅在任务进行时有效，任务退出后查询将会返回错误。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudModeration(request *DescribeCloudModerationRequest) (response *DescribeCloudModerationResponse, err error) {
    return c.DescribeCloudModerationWithContext(context.Background(), request)
}

// DescribeCloudModeration
// 成功开启审核任务后，可以使用此接口来查询审核任务状态和订阅的黑白名单信息。仅在任务进行时有效，任务退出后查询将会返回错误。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudModerationWithContext(ctx context.Context, request *DescribeCloudModerationRequest) (response *DescribeCloudModerationResponse, err error) {
    if request == nil {
        request = NewDescribeCloudModerationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeCloudModeration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudModeration require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudModerationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudRecordingRequest() (request *DescribeCloudRecordingRequest) {
    request = &DescribeCloudRecordingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeCloudRecording")
    
    
    return
}

func NewDescribeCloudRecordingResponse() (response *DescribeCloudRecordingResponse) {
    response = &DescribeCloudRecordingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudRecording
// 成功开启录制后，可以使用此接口来查询录制状态。仅在录制任务进行时有效，录制退出后查询将会返回错误。
//
// 录制文件上传到云点播VOD时，StorageFileList中不会返回录制文件信息，请订阅相关录制文件回调事件，获取录制文件信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudRecording(request *DescribeCloudRecordingRequest) (response *DescribeCloudRecordingResponse, err error) {
    return c.DescribeCloudRecordingWithContext(context.Background(), request)
}

// DescribeCloudRecording
// 成功开启录制后，可以使用此接口来查询录制状态。仅在录制任务进行时有效，录制退出后查询将会返回错误。
//
// 录制文件上传到云点播VOD时，StorageFileList中不会返回录制文件信息，请订阅相关录制文件回调事件，获取录制文件信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudRecordingWithContext(ctx context.Context, request *DescribeCloudRecordingRequest) (response *DescribeCloudRecordingResponse, err error) {
    if request == nil {
        request = NewDescribeCloudRecordingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeCloudRecording")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudRecordingResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudSliceTaskRequest() (request *DescribeCloudSliceTaskRequest) {
    request = &DescribeCloudSliceTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeCloudSliceTask")
    
    
    return
}

func NewDescribeCloudSliceTaskResponse() (response *DescribeCloudSliceTaskResponse) {
    response = &DescribeCloudSliceTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeCloudSliceTask
// 成功开启切片后，可以使用此接口来查询切片任务状态。仅在任务进行时有效，任务退出后查询将会返回错误。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudSliceTask(request *DescribeCloudSliceTaskRequest) (response *DescribeCloudSliceTaskResponse, err error) {
    return c.DescribeCloudSliceTaskWithContext(context.Background(), request)
}

// DescribeCloudSliceTask
// 成功开启切片后，可以使用此接口来查询切片任务状态。仅在任务进行时有效，任务退出后查询将会返回错误。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) DescribeCloudSliceTaskWithContext(ctx context.Context, request *DescribeCloudSliceTaskRequest) (response *DescribeCloudSliceTaskResponse, err error) {
    if request == nil {
        request = NewDescribeCloudSliceTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeCloudSliceTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudSliceTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudSliceTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMixTranscodingUsageRequest() (request *DescribeMixTranscodingUsageRequest) {
    request = &DescribeMixTranscodingUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeMixTranscodingUsage")
    
    
    return
}

func NewDescribeMixTranscodingUsageResponse() (response *DescribeMixTranscodingUsageResponse) {
    response = &DescribeMixTranscodingUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMixTranscodingUsage
// 获取TRTC混流转码的用量明细。
//
// - 查询时间小于等于1天时，返回每5分钟粒度的数据；查询时间大于1天时，返回按天汇总的数据。
//
// - 单次查询统计区间最多不能超过31天。
//
// - 若查询当天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 该接口只用于历史用量数据统计或核对数据使用，关键业务逻辑不能使用。
//
// - 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeMixTranscodingUsage(request *DescribeMixTranscodingUsageRequest) (response *DescribeMixTranscodingUsageResponse, err error) {
    return c.DescribeMixTranscodingUsageWithContext(context.Background(), request)
}

// DescribeMixTranscodingUsage
// 获取TRTC混流转码的用量明细。
//
// - 查询时间小于等于1天时，返回每5分钟粒度的数据；查询时间大于1天时，返回按天汇总的数据。
//
// - 单次查询统计区间最多不能超过31天。
//
// - 若查询当天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 该接口只用于历史用量数据统计或核对数据使用，关键业务逻辑不能使用。
//
// - 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeMixTranscodingUsageWithContext(ctx context.Context, request *DescribeMixTranscodingUsageRequest) (response *DescribeMixTranscodingUsageResponse, err error) {
    if request == nil {
        request = NewDescribeMixTranscodingUsageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeMixTranscodingUsage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMixTranscodingUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMixTranscodingUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePictureRequest() (request *DescribePictureRequest) {
    request = &DescribePictureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribePicture")
    
    
    return
}

func NewDescribePictureResponse() (response *DescribePictureResponse) {
    response = &DescribePictureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePicture
// 如果您需要在 [云端混流转码](https://cloud.tencent.com/document/product/647/16827) 时频繁查找自定义背景图或水印信息，可通过此接口查找已上传的图片信息。无需频繁查找图片信息的场景，建议直接在 [控制台 > 应用管理 > 素材管理](https://cloud.tencent.com/document/product/647/50769) 中查看。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribePicture(request *DescribePictureRequest) (response *DescribePictureResponse, err error) {
    return c.DescribePictureWithContext(context.Background(), request)
}

// DescribePicture
// 如果您需要在 [云端混流转码](https://cloud.tencent.com/document/product/647/16827) 时频繁查找自定义背景图或水印信息，可通过此接口查找已上传的图片信息。无需频繁查找图片信息的场景，建议直接在 [控制台 > 应用管理 > 素材管理](https://cloud.tencent.com/document/product/647/50769) 中查看。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribePictureWithContext(ctx context.Context, request *DescribePictureRequest) (response *DescribePictureResponse, err error) {
    if request == nil {
        request = NewDescribePictureRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribePicture")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePicture require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePictureResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordStatisticRequest() (request *DescribeRecordStatisticRequest) {
    request = &DescribeRecordStatisticRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeRecordStatistic")
    
    
    return
}

func NewDescribeRecordStatisticResponse() (response *DescribeRecordStatisticResponse) {
    response = &DescribeRecordStatisticResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordStatistic
// 查询云端录制计费时长。
//
// 
//
// - 查询时间小于等于1天时，返回每5分钟粒度的数据；查询时间大于1天时，返回按天汇总的数据。
//
// - 单次查询统计区间最多不能超过31天。
//
// - 若查询当天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 日结后付费将于次日上午推送账单，建议次日上午9点以后再来查询前一天的用量。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPID = "InvalidParameter.AppId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeRecordStatistic(request *DescribeRecordStatisticRequest) (response *DescribeRecordStatisticResponse, err error) {
    return c.DescribeRecordStatisticWithContext(context.Background(), request)
}

// DescribeRecordStatistic
// 查询云端录制计费时长。
//
// 
//
// - 查询时间小于等于1天时，返回每5分钟粒度的数据；查询时间大于1天时，返回按天汇总的数据。
//
// - 单次查询统计区间最多不能超过31天。
//
// - 若查询当天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 日结后付费将于次日上午推送账单，建议次日上午9点以后再来查询前一天的用量。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_APPID = "InvalidParameter.AppId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeRecordStatisticWithContext(ctx context.Context, request *DescribeRecordStatisticRequest) (response *DescribeRecordStatisticResponse, err error) {
    if request == nil {
        request = NewDescribeRecordStatisticRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeRecordStatistic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordStatistic require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordStatisticResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRecordingUsageRequest() (request *DescribeRecordingUsageRequest) {
    request = &DescribeRecordingUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeRecordingUsage")
    
    
    return
}

func NewDescribeRecordingUsageResponse() (response *DescribeRecordingUsageResponse) {
    response = &DescribeRecordingUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRecordingUsage
// 获取TRTC录制的用量明细。
//
// - 查询时间小于等于1天时，返回每5分钟粒度的数据；查询时间大于1天时，返回按天汇总的数据。
//
// - 单次查询统计区间最多不能超过31天。
//
// - 若查询当天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 该接口只用于历史用量数据统计或核对数据使用，关键业务逻辑不能使用。
//
// - 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeRecordingUsage(request *DescribeRecordingUsageRequest) (response *DescribeRecordingUsageResponse, err error) {
    return c.DescribeRecordingUsageWithContext(context.Background(), request)
}

// DescribeRecordingUsage
// 获取TRTC录制的用量明细。
//
// - 查询时间小于等于1天时，返回每5分钟粒度的数据；查询时间大于1天时，返回按天汇总的数据。
//
// - 单次查询统计区间最多不能超过31天。
//
// - 若查询当天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 该接口只用于历史用量数据统计或核对数据使用，关键业务逻辑不能使用。
//
// - 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeRecordingUsageWithContext(ctx context.Context, request *DescribeRecordingUsageRequest) (response *DescribeRecordingUsageResponse, err error) {
    if request == nil {
        request = NewDescribeRecordingUsageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeRecordingUsage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRecordingUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRecordingUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRelayUsageRequest() (request *DescribeRelayUsageRequest) {
    request = &DescribeRelayUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeRelayUsage")
    
    
    return
}

func NewDescribeRelayUsageResponse() (response *DescribeRelayUsageResponse) {
    response = &DescribeRelayUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRelayUsage
// 获取TRTC旁路转推的用量明细。
//
// - 查询时间小于等于1天时，返回每5分钟粒度的数据；查询时间大于1天时，返回按天汇总的数据。
//
// - 单次查询统计区间最多不能超过31天。
//
// - 若查询当天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 该接口只用于历史用量数据统计或核对数据使用，关键业务逻辑不能使用。
//
// - 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeRelayUsage(request *DescribeRelayUsageRequest) (response *DescribeRelayUsageResponse, err error) {
    return c.DescribeRelayUsageWithContext(context.Background(), request)
}

// DescribeRelayUsage
// 获取TRTC旁路转推的用量明细。
//
// - 查询时间小于等于1天时，返回每5分钟粒度的数据；查询时间大于1天时，返回按天汇总的数据。
//
// - 单次查询统计区间最多不能超过31天。
//
// - 若查询当天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 该接口只用于历史用量数据统计或核对数据使用，关键业务逻辑不能使用。
//
// - 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeRelayUsageWithContext(ctx context.Context, request *DescribeRelayUsageRequest) (response *DescribeRelayUsageResponse, err error) {
    if request == nil {
        request = NewDescribeRelayUsageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeRelayUsage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRelayUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRelayUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeRoomInfoRequest() (request *DescribeRoomInfoRequest) {
    request = &DescribeRoomInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeRoomInfo")
    
    
    return
}

func NewDescribeRoomInfoResponse() (response *DescribeRoomInfoResponse) {
    response = &DescribeRoomInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeRoomInfo
// 查询SdkAppId下的房间列表。默认返回10条通话，一次最多返回100条通话。最大可查询14天内的数据。（同老接口DescribeRoomInformation）
//
// **注意**：
//
// 1.该接口只用于历史数据统计或核对数据使用，实时类关键业务逻辑不能使用。
//
// 2.该接口自2024年4月1日起正式商业化，需订阅套餐解锁调用能力，提供以下两种解锁方式，可任意其一解锁：
//
// 方式一：通过订阅[包月套餐](https://cloud.tencent.com/document/product/647/85386)「尊享版」（可查近7天）和「旗舰版」（可查近14天），[前往订阅](https://buy.cloud.tencent.com/trtc?trtcversion=top)。
//
// 方式二：通过订阅[监控仪表盘](https://cloud.tencent.com/document/product/647/81331)商业套餐包「基础版」（可查近7天）和「进阶版」（可查近14天），[前往订阅](https://buy.cloud.tencent.com/trtc_monitor)。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_PAGENUMBER = "InvalidParameter.PageNumber"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIMEOVERSIZE = "InvalidParameter.StartTimeOversize"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_URLPARAMSERROR = "InvalidParameter.UrlParamsError"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_ROOMNUM = "MissingParameter.RoomNum"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeRoomInfo(request *DescribeRoomInfoRequest) (response *DescribeRoomInfoResponse, err error) {
    return c.DescribeRoomInfoWithContext(context.Background(), request)
}

// DescribeRoomInfo
// 查询SdkAppId下的房间列表。默认返回10条通话，一次最多返回100条通话。最大可查询14天内的数据。（同老接口DescribeRoomInformation）
//
// **注意**：
//
// 1.该接口只用于历史数据统计或核对数据使用，实时类关键业务逻辑不能使用。
//
// 2.该接口自2024年4月1日起正式商业化，需订阅套餐解锁调用能力，提供以下两种解锁方式，可任意其一解锁：
//
// 方式一：通过订阅[包月套餐](https://cloud.tencent.com/document/product/647/85386)「尊享版」（可查近7天）和「旗舰版」（可查近14天），[前往订阅](https://buy.cloud.tencent.com/trtc?trtcversion=top)。
//
// 方式二：通过订阅[监控仪表盘](https://cloud.tencent.com/document/product/647/81331)商业套餐包「基础版」（可查近7天）和「进阶版」（可查近14天），[前往订阅](https://buy.cloud.tencent.com/trtc_monitor)。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_PAGENUMBER = "InvalidParameter.PageNumber"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIMEOVERSIZE = "InvalidParameter.StartTimeOversize"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_URLPARAMSERROR = "InvalidParameter.UrlParamsError"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_ROOMNUM = "MissingParameter.RoomNum"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeRoomInfoWithContext(ctx context.Context, request *DescribeRoomInfoRequest) (response *DescribeRoomInfoResponse, err error) {
    if request == nil {
        request = NewDescribeRoomInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeRoomInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeRoomInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeRoomInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeScaleInfoRequest() (request *DescribeScaleInfoRequest) {
    request = &DescribeScaleInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeScaleInfo")
    
    
    return
}

func NewDescribeScaleInfoResponse() (response *DescribeScaleInfoResponse) {
    response = &DescribeScaleInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeScaleInfo
// 可查询SdkAppId每天的房间数和用户数，按天统计，可查询最近14天的数据。当天未结束，数据未统计完成，无法查到当天的房间数与用户数。（同老接口DescribeHistoryScale）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_USERIDSMORETHANSIX = "InvalidParameter.UserIdsMorethanSix"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeScaleInfo(request *DescribeScaleInfoRequest) (response *DescribeScaleInfoResponse, err error) {
    return c.DescribeScaleInfoWithContext(context.Background(), request)
}

// DescribeScaleInfo
// 可查询SdkAppId每天的房间数和用户数，按天统计，可查询最近14天的数据。当天未结束，数据未统计完成，无法查到当天的房间数与用户数。（同老接口DescribeHistoryScale）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_USERIDSMORETHANSIX = "InvalidParameter.UserIdsMorethanSix"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeScaleInfoWithContext(ctx context.Context, request *DescribeScaleInfoRequest) (response *DescribeScaleInfoResponse, err error) {
    if request == nil {
        request = NewDescribeScaleInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeScaleInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeScaleInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeScaleInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeStreamIngestRequest() (request *DescribeStreamIngestRequest) {
    request = &DescribeStreamIngestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeStreamIngest")
    
    
    return
}

func NewDescribeStreamIngestResponse() (response *DescribeStreamIngestResponse) {
    response = &DescribeStreamIngestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeStreamIngest
// 您可以查询输入在线媒体流任务的状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYTASKINFOFAILED = "FailedOperation.QueryTaskInfoFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
func (c *Client) DescribeStreamIngest(request *DescribeStreamIngestRequest) (response *DescribeStreamIngestResponse, err error) {
    return c.DescribeStreamIngestWithContext(context.Background(), request)
}

// DescribeStreamIngest
// 您可以查询输入在线媒体流任务的状态。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYTASKINFOFAILED = "FailedOperation.QueryTaskInfoFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
func (c *Client) DescribeStreamIngestWithContext(ctx context.Context, request *DescribeStreamIngestRequest) (response *DescribeStreamIngestResponse, err error) {
    if request == nil {
        request = NewDescribeStreamIngestRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeStreamIngest")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeStreamIngest require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeStreamIngestResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTRTCMarketQualityDataRequest() (request *DescribeTRTCMarketQualityDataRequest) {
    request = &DescribeTRTCMarketQualityDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTRTCMarketQualityData")
    
    
    return
}

func NewDescribeTRTCMarketQualityDataResponse() (response *DescribeTRTCMarketQualityDataResponse) {
    response = &DescribeTRTCMarketQualityDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTRTCMarketQualityData
// 查询TRTC监控仪表盘-数据大盘质量指标（包括下列指标）
//
// joinSuccessRate：加入频道成功率。
//
// joinSuccessIn5sRate：5s内加入频道成功率。
//
// audioFreezeRate：音频卡顿率。
//
// videoFreezeRate：视频卡顿率。
//
// networkDelay ：网络延迟率。
//
// 注意：
//
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，监控仪表盘[版本功能和计费说明](https://cloud.tencent.com/document/product/647/81331)。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，【基础版】可查近30天，【进阶版】可查近60天。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYTASKINFOFAILED = "FailedOperation.QueryTaskInfoFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
func (c *Client) DescribeTRTCMarketQualityData(request *DescribeTRTCMarketQualityDataRequest) (response *DescribeTRTCMarketQualityDataResponse, err error) {
    return c.DescribeTRTCMarketQualityDataWithContext(context.Background(), request)
}

// DescribeTRTCMarketQualityData
// 查询TRTC监控仪表盘-数据大盘质量指标（包括下列指标）
//
// joinSuccessRate：加入频道成功率。
//
// joinSuccessIn5sRate：5s内加入频道成功率。
//
// audioFreezeRate：音频卡顿率。
//
// videoFreezeRate：视频卡顿率。
//
// networkDelay ：网络延迟率。
//
// 注意：
//
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，监控仪表盘[版本功能和计费说明](https://cloud.tencent.com/document/product/647/81331)。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，【基础版】可查近30天，【进阶版】可查近60天。
//
// 可能返回的错误码:
//  FAILEDOPERATION_QUERYTASKINFOFAILED = "FailedOperation.QueryTaskInfoFailed"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
func (c *Client) DescribeTRTCMarketQualityDataWithContext(ctx context.Context, request *DescribeTRTCMarketQualityDataRequest) (response *DescribeTRTCMarketQualityDataResponse, err error) {
    if request == nil {
        request = NewDescribeTRTCMarketQualityDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeTRTCMarketQualityData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTRTCMarketQualityData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTRTCMarketQualityDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTRTCMarketQualityMetricDataRequest() (request *DescribeTRTCMarketQualityMetricDataRequest) {
    request = &DescribeTRTCMarketQualityMetricDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTRTCMarketQualityMetricData")
    
    
    return
}

func NewDescribeTRTCMarketQualityMetricDataResponse() (response *DescribeTRTCMarketQualityMetricDataResponse) {
    response = &DescribeTRTCMarketQualityMetricDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTRTCMarketQualityMetricData
// 查询TRTC监控仪表盘-数据大盘质量指标（包括下列指标）
//
// joinSuccessRate：加入频道成功率。
//
// joinSuccessIn5sRate：5s内加入频道成功率。
//
// audioFreezeRate：音频卡顿率。
//
// videoFreezeRate：视频卡顿率。
//
// networkDelay ：网络延迟率。
//
// 注意：
//
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，监控仪表盘版本功能和计费说明：https://cloud.tencent.com/document/product/647/81331。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，【基础版】可查近30天，【进阶版】可查近60天。
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTRTCMarketQualityMetricData(request *DescribeTRTCMarketQualityMetricDataRequest) (response *DescribeTRTCMarketQualityMetricDataResponse, err error) {
    return c.DescribeTRTCMarketQualityMetricDataWithContext(context.Background(), request)
}

// DescribeTRTCMarketQualityMetricData
// 查询TRTC监控仪表盘-数据大盘质量指标（包括下列指标）
//
// joinSuccessRate：加入频道成功率。
//
// joinSuccessIn5sRate：5s内加入频道成功率。
//
// audioFreezeRate：音频卡顿率。
//
// videoFreezeRate：视频卡顿率。
//
// networkDelay ：网络延迟率。
//
// 注意：
//
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，监控仪表盘版本功能和计费说明：https://cloud.tencent.com/document/product/647/81331。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，【基础版】可查近30天，【进阶版】可查近60天。
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTRTCMarketQualityMetricDataWithContext(ctx context.Context, request *DescribeTRTCMarketQualityMetricDataRequest) (response *DescribeTRTCMarketQualityMetricDataResponse, err error) {
    if request == nil {
        request = NewDescribeTRTCMarketQualityMetricDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeTRTCMarketQualityMetricData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTRTCMarketQualityMetricData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTRTCMarketQualityMetricDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTRTCMarketScaleDataRequest() (request *DescribeTRTCMarketScaleDataRequest) {
    request = &DescribeTRTCMarketScaleDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTRTCMarketScaleData")
    
    
    return
}

func NewDescribeTRTCMarketScaleDataResponse() (response *DescribeTRTCMarketScaleDataResponse) {
    response = &DescribeTRTCMarketScaleDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTRTCMarketScaleData
// 查询TRTC监控仪表盘-数据大盘规模指标（会返回通话人数，通话房间数，峰值同时在线人数，峰值同时在线频道数）
//
// userCount：通话人数，
//
// roomCount：通话房间数，从有用户加入频道到所有用户离开频道计为一个通话频道。
//
// peakCurrentChannels：峰值同时在线频道数。
//
// peakCurrentUsers：峰值同时在线人数。
//
// 注意：
//
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，监控仪表盘[版本功能和计费说明](https://cloud.tencent.com/document/product/647/81331)。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，【基础版】可查近30天，【进阶版】可查近60天。
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTRTCMarketScaleData(request *DescribeTRTCMarketScaleDataRequest) (response *DescribeTRTCMarketScaleDataResponse, err error) {
    return c.DescribeTRTCMarketScaleDataWithContext(context.Background(), request)
}

// DescribeTRTCMarketScaleData
// 查询TRTC监控仪表盘-数据大盘规模指标（会返回通话人数，通话房间数，峰值同时在线人数，峰值同时在线频道数）
//
// userCount：通话人数，
//
// roomCount：通话房间数，从有用户加入频道到所有用户离开频道计为一个通话频道。
//
// peakCurrentChannels：峰值同时在线频道数。
//
// peakCurrentUsers：峰值同时在线人数。
//
// 注意：
//
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，监控仪表盘[版本功能和计费说明](https://cloud.tencent.com/document/product/647/81331)。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，【基础版】可查近30天，【进阶版】可查近60天。
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTRTCMarketScaleDataWithContext(ctx context.Context, request *DescribeTRTCMarketScaleDataRequest) (response *DescribeTRTCMarketScaleDataResponse, err error) {
    if request == nil {
        request = NewDescribeTRTCMarketScaleDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeTRTCMarketScaleData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTRTCMarketScaleData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTRTCMarketScaleDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTRTCMarketScaleMetricDataRequest() (request *DescribeTRTCMarketScaleMetricDataRequest) {
    request = &DescribeTRTCMarketScaleMetricDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTRTCMarketScaleMetricData")
    
    
    return
}

func NewDescribeTRTCMarketScaleMetricDataResponse() (response *DescribeTRTCMarketScaleMetricDataResponse) {
    response = &DescribeTRTCMarketScaleMetricDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTRTCMarketScaleMetricData
// 查询TRTC监控仪表盘-数据大盘规模指标（会返回通话人数，通话房间数，峰值同时在线人数，峰值同时在线频道数）
//
// userCount：通话人数，
//
// roomCount：通话房间数，从有用户加入频道到所有用户离开频道计为一个通话频道。
//
// peakCurrentChannels：峰值同时在线频道数。
//
// peakCurrentUsers：峰值同时在线人数。
//
// 注意：
//
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，监控仪表盘版本功能和计费说明：https://cloud.tencent.com/document/product/647/81331。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，【基础版】可查近30天，【进阶版】可查近60天。
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTRTCMarketScaleMetricData(request *DescribeTRTCMarketScaleMetricDataRequest) (response *DescribeTRTCMarketScaleMetricDataResponse, err error) {
    return c.DescribeTRTCMarketScaleMetricDataWithContext(context.Background(), request)
}

// DescribeTRTCMarketScaleMetricData
// 查询TRTC监控仪表盘-数据大盘规模指标（会返回通话人数，通话房间数，峰值同时在线人数，峰值同时在线频道数）
//
// userCount：通话人数，
//
// roomCount：通话房间数，从有用户加入频道到所有用户离开频道计为一个通话频道。
//
// peakCurrentChannels：峰值同时在线频道数。
//
// peakCurrentUsers：峰值同时在线人数。
//
// 注意：
//
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，监控仪表盘版本功能和计费说明：https://cloud.tencent.com/document/product/647/81331。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，【基础版】可查近30天，【进阶版】可查近60天。
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTRTCMarketScaleMetricDataWithContext(ctx context.Context, request *DescribeTRTCMarketScaleMetricDataRequest) (response *DescribeTRTCMarketScaleMetricDataResponse, err error) {
    if request == nil {
        request = NewDescribeTRTCMarketScaleMetricDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeTRTCMarketScaleMetricData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTRTCMarketScaleMetricData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTRTCMarketScaleMetricDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTRTCRealTimeQualityDataRequest() (request *DescribeTRTCRealTimeQualityDataRequest) {
    request = &DescribeTRTCRealTimeQualityDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTRTCRealTimeQualityData")
    
    
    return
}

func NewDescribeTRTCRealTimeQualityDataResponse() (response *DescribeTRTCRealTimeQualityDataResponse) {
    response = &DescribeTRTCRealTimeQualityDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTRTCRealTimeQualityData
// 查询TRTC监控仪表盘-实时监控质量指标（会返回下列指标）
//
// -视频卡顿率
//
// -音频卡顿率
//
// 注意：
//
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，详情参考[监控仪表盘](https://cloud.tencent.com/document/product/647/81331)。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDSDKAPPID = "InvalidParameter.InvalidSdkAppId"
func (c *Client) DescribeTRTCRealTimeQualityData(request *DescribeTRTCRealTimeQualityDataRequest) (response *DescribeTRTCRealTimeQualityDataResponse, err error) {
    return c.DescribeTRTCRealTimeQualityDataWithContext(context.Background(), request)
}

// DescribeTRTCRealTimeQualityData
// 查询TRTC监控仪表盘-实时监控质量指标（会返回下列指标）
//
// -视频卡顿率
//
// -音频卡顿率
//
// 注意：
//
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，详情参考[监控仪表盘](https://cloud.tencent.com/document/product/647/81331)。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDSDKAPPID = "InvalidParameter.InvalidSdkAppId"
func (c *Client) DescribeTRTCRealTimeQualityDataWithContext(ctx context.Context, request *DescribeTRTCRealTimeQualityDataRequest) (response *DescribeTRTCRealTimeQualityDataResponse, err error) {
    if request == nil {
        request = NewDescribeTRTCRealTimeQualityDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeTRTCRealTimeQualityData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTRTCRealTimeQualityData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTRTCRealTimeQualityDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTRTCRealTimeQualityMetricDataRequest() (request *DescribeTRTCRealTimeQualityMetricDataRequest) {
    request = &DescribeTRTCRealTimeQualityMetricDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTRTCRealTimeQualityMetricData")
    
    
    return
}

func NewDescribeTRTCRealTimeQualityMetricDataResponse() (response *DescribeTRTCRealTimeQualityMetricDataResponse) {
    response = &DescribeTRTCRealTimeQualityMetricDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTRTCRealTimeQualityMetricData
// 查询TRTC监控仪表盘-实时监控质量指标（会返回下列指标）
//
// -视频卡顿率
//
// -音频卡顿率
//
// 注意：
//
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，监控仪表盘版本功能和计费说明：https://cloud.tencent.com/document/product/647/81331。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时。
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTRTCRealTimeQualityMetricData(request *DescribeTRTCRealTimeQualityMetricDataRequest) (response *DescribeTRTCRealTimeQualityMetricDataResponse, err error) {
    return c.DescribeTRTCRealTimeQualityMetricDataWithContext(context.Background(), request)
}

// DescribeTRTCRealTimeQualityMetricData
// 查询TRTC监控仪表盘-实时监控质量指标（会返回下列指标）
//
// -视频卡顿率
//
// -音频卡顿率
//
// 注意：
//
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，监控仪表盘版本功能和计费说明：https://cloud.tencent.com/document/product/647/81331。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时。
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTRTCRealTimeQualityMetricDataWithContext(ctx context.Context, request *DescribeTRTCRealTimeQualityMetricDataRequest) (response *DescribeTRTCRealTimeQualityMetricDataResponse, err error) {
    if request == nil {
        request = NewDescribeTRTCRealTimeQualityMetricDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeTRTCRealTimeQualityMetricData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTRTCRealTimeQualityMetricData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTRTCRealTimeQualityMetricDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTRTCRealTimeScaleDataRequest() (request *DescribeTRTCRealTimeScaleDataRequest) {
    request = &DescribeTRTCRealTimeScaleDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTRTCRealTimeScaleData")
    
    
    return
}

func NewDescribeTRTCRealTimeScaleDataResponse() (response *DescribeTRTCRealTimeScaleDataResponse) {
    response = &DescribeTRTCRealTimeScaleDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTRTCRealTimeScaleData
// 查询TRTC监控仪表盘-实时监控规模指标（会返回下列指标）
//
// -userCount（在线用户数）
//
// -roomCount（在线房间数）
//
// 注意：
//
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，详情参考[监控仪表盘](https://cloud.tencent.com/document/product/647/81331)。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时。
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTRTCRealTimeScaleData(request *DescribeTRTCRealTimeScaleDataRequest) (response *DescribeTRTCRealTimeScaleDataResponse, err error) {
    return c.DescribeTRTCRealTimeScaleDataWithContext(context.Background(), request)
}

// DescribeTRTCRealTimeScaleData
// 查询TRTC监控仪表盘-实时监控规模指标（会返回下列指标）
//
// -userCount（在线用户数）
//
// -roomCount（在线房间数）
//
// 注意：
//
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，详情参考[监控仪表盘](https://cloud.tencent.com/document/product/647/81331)。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时。
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTRTCRealTimeScaleDataWithContext(ctx context.Context, request *DescribeTRTCRealTimeScaleDataRequest) (response *DescribeTRTCRealTimeScaleDataResponse, err error) {
    if request == nil {
        request = NewDescribeTRTCRealTimeScaleDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeTRTCRealTimeScaleData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTRTCRealTimeScaleData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTRTCRealTimeScaleDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTRTCRealTimeScaleMetricDataRequest() (request *DescribeTRTCRealTimeScaleMetricDataRequest) {
    request = &DescribeTRTCRealTimeScaleMetricDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTRTCRealTimeScaleMetricData")
    
    
    return
}

func NewDescribeTRTCRealTimeScaleMetricDataResponse() (response *DescribeTRTCRealTimeScaleMetricDataResponse) {
    response = &DescribeTRTCRealTimeScaleMetricDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTRTCRealTimeScaleMetricData
// 查询TRTC监控仪表盘-实时监控规模指标（会返回下列指标）
//
// -userCount（在线用户数）
//
// -roomCount（在线房间数）
//
// 注意：
//
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，监控仪表盘版本功能和计费说明：https://cloud.tencent.com/document/product/647/81331。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时。
//
// xa0
//
// 3.除此之外您也可以通过订阅TRTC包月套餐(https://buy.cloud.tencent.com/trtc)尊享版或旗舰版解锁此接口（DescribeTRTCRealTimeScaleMetricData）的调用能力，请在开通包月套餐后，请提交工单联系售后解锁调用能力https://console.cloud.tencent.com/workorder/category
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTRTCRealTimeScaleMetricData(request *DescribeTRTCRealTimeScaleMetricDataRequest) (response *DescribeTRTCRealTimeScaleMetricDataResponse, err error) {
    return c.DescribeTRTCRealTimeScaleMetricDataWithContext(context.Background(), request)
}

// DescribeTRTCRealTimeScaleMetricData
// 查询TRTC监控仪表盘-实时监控规模指标（会返回下列指标）
//
// -userCount（在线用户数）
//
// -roomCount（在线房间数）
//
// 注意：
//
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，监控仪表盘版本功能和计费说明：https://cloud.tencent.com/document/product/647/81331。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时。
//
// xa0
//
// 3.除此之外您也可以通过订阅TRTC包月套餐(https://buy.cloud.tencent.com/trtc)尊享版或旗舰版解锁此接口（DescribeTRTCRealTimeScaleMetricData）的调用能力，请在开通包月套餐后，请提交工单联系售后解锁调用能力https://console.cloud.tencent.com/workorder/category
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTRTCRealTimeScaleMetricDataWithContext(ctx context.Context, request *DescribeTRTCRealTimeScaleMetricDataRequest) (response *DescribeTRTCRealTimeScaleMetricDataResponse, err error) {
    if request == nil {
        request = NewDescribeTRTCRealTimeScaleMetricDataRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeTRTCRealTimeScaleMetricData")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTRTCRealTimeScaleMetricData require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTRTCRealTimeScaleMetricDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrtcMcuTranscodeTimeRequest() (request *DescribeTrtcMcuTranscodeTimeRequest) {
    request = &DescribeTrtcMcuTranscodeTimeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTrtcMcuTranscodeTime")
    
    
    return
}

func NewDescribeTrtcMcuTranscodeTimeResponse() (response *DescribeTrtcMcuTranscodeTimeResponse) {
    response = &DescribeTrtcMcuTranscodeTimeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTrtcMcuTranscodeTime
// 查询旁路转码计费时长。
//
// - 查询时间小于等于1天时，返回每5分钟粒度的数据；查询时间大于1天时，返回按天汇总的数据。
//
// - 单次查询统计区间最多不能超过31天。
//
// - 若查询当天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 日结后付费将于次日上午推送账单，建议次日上午9点以后再来查询前一天的用量。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeTrtcMcuTranscodeTime(request *DescribeTrtcMcuTranscodeTimeRequest) (response *DescribeTrtcMcuTranscodeTimeResponse, err error) {
    return c.DescribeTrtcMcuTranscodeTimeWithContext(context.Background(), request)
}

// DescribeTrtcMcuTranscodeTime
// 查询旁路转码计费时长。
//
// - 查询时间小于等于1天时，返回每5分钟粒度的数据；查询时间大于1天时，返回按天汇总的数据。
//
// - 单次查询统计区间最多不能超过31天。
//
// - 若查询当天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 日结后付费将于次日上午推送账单，建议次日上午9点以后再来查询前一天的用量。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeTrtcMcuTranscodeTimeWithContext(ctx context.Context, request *DescribeTrtcMcuTranscodeTimeRequest) (response *DescribeTrtcMcuTranscodeTimeResponse, err error) {
    if request == nil {
        request = NewDescribeTrtcMcuTranscodeTimeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeTrtcMcuTranscodeTime")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrtcMcuTranscodeTime require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrtcMcuTranscodeTimeResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrtcRoomUsageRequest() (request *DescribeTrtcRoomUsageRequest) {
    request = &DescribeTrtcRoomUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTrtcRoomUsage")
    
    
    return
}

func NewDescribeTrtcRoomUsageResponse() (response *DescribeTrtcRoomUsageResponse) {
    response = &DescribeTrtcRoomUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTrtcRoomUsage
// 查询TRTC音视频房间维度用量。
//
// - 单次只能查询一天数据，返回查询时间段内的汇总数据；通过多次查询可以查不同天数据。若查询跨天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 该接口只用于历史用量数据统计或核对数据使用，关键业务逻辑不能使用，不可用于账单核对，如需对账请使用账号/应用维度用量API：DescribeTrtcUsage。
//
// - 默认接口请求频率限制：1次/15秒。
//
// - 数据最早可查日期为2023年4月1日0点，最大可查范围近3个月。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SDKAPPIDNOTEXIST = "FailedOperation.SdkAppIdNotExist"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_FREQOVERLIMIT = "OperationDenied.FreqOverLimit"
func (c *Client) DescribeTrtcRoomUsage(request *DescribeTrtcRoomUsageRequest) (response *DescribeTrtcRoomUsageResponse, err error) {
    return c.DescribeTrtcRoomUsageWithContext(context.Background(), request)
}

// DescribeTrtcRoomUsage
// 查询TRTC音视频房间维度用量。
//
// - 单次只能查询一天数据，返回查询时间段内的汇总数据；通过多次查询可以查不同天数据。若查询跨天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 该接口只用于历史用量数据统计或核对数据使用，关键业务逻辑不能使用，不可用于账单核对，如需对账请使用账号/应用维度用量API：DescribeTrtcUsage。
//
// - 默认接口请求频率限制：1次/15秒。
//
// - 数据最早可查日期为2023年4月1日0点，最大可查范围近3个月。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SDKAPPIDNOTEXIST = "FailedOperation.SdkAppIdNotExist"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  OPERATIONDENIED_FREQOVERLIMIT = "OperationDenied.FreqOverLimit"
func (c *Client) DescribeTrtcRoomUsageWithContext(ctx context.Context, request *DescribeTrtcRoomUsageRequest) (response *DescribeTrtcRoomUsageResponse, err error) {
    if request == nil {
        request = NewDescribeTrtcRoomUsageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeTrtcRoomUsage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrtcRoomUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrtcRoomUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTrtcUsageRequest() (request *DescribeTrtcUsageRequest) {
    request = &DescribeTrtcUsageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeTrtcUsage")
    
    
    return
}

func NewDescribeTrtcUsageResponse() (response *DescribeTrtcUsageResponse) {
    response = &DescribeTrtcUsageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTrtcUsage
// 获取TRTC音视频互动的用量明细，单位:分钟。
//
// - 查询时间小于等于1天时，返回每5分钟粒度的数据；查询时间大于1天时，返回按天汇总的数据。
//
// - 单次查询统计区间最多不能超过31天。
//
// - 若查询当天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 该接口只用于历史用量数据统计或核对数据使用，关键业务逻辑不能使用。
//
// - 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeTrtcUsage(request *DescribeTrtcUsageRequest) (response *DescribeTrtcUsageResponse, err error) {
    return c.DescribeTrtcUsageWithContext(context.Background(), request)
}

// DescribeTrtcUsage
// 获取TRTC音视频互动的用量明细，单位:分钟。
//
// - 查询时间小于等于1天时，返回每5分钟粒度的数据；查询时间大于1天时，返回按天汇总的数据。
//
// - 单次查询统计区间最多不能超过31天。
//
// - 若查询当天用量，由于统计延迟等原因，返回数据可能不够准确。
//
// - 该接口只用于历史用量数据统计或核对数据使用，关键业务逻辑不能使用。
//
// - 默认接口请求频率限制：5次/秒。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
func (c *Client) DescribeTrtcUsageWithContext(ctx context.Context, request *DescribeTrtcUsageRequest) (response *DescribeTrtcUsageResponse, err error) {
    if request == nil {
        request = NewDescribeTrtcUsageRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeTrtcUsage")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTrtcUsage require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTrtcUsageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUnusualEventRequest() (request *DescribeUnusualEventRequest) {
    request = &DescribeUnusualEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeUnusualEvent")
    
    
    return
}

func NewDescribeUnusualEventResponse() (response *DescribeUnusualEventResponse) {
    response = &DescribeUnusualEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUnusualEvent
// 查询SdkAppId下任意20条异常体验事件，返回异常体验ID与可能产生异常体验的原因。可查询14天内数据，查询起止时间不超过1个小时。支持跨天查询。（同老接口DescribeAbnormalEvent）
//
// 异常体验ID映射见：https://cloud.tencent.com/document/product/647/44916
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIMEEXPIRE = "InvalidParameter.StartTimeExpire"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeUnusualEvent(request *DescribeUnusualEventRequest) (response *DescribeUnusualEventResponse, err error) {
    return c.DescribeUnusualEventWithContext(context.Background(), request)
}

// DescribeUnusualEvent
// 查询SdkAppId下任意20条异常体验事件，返回异常体验ID与可能产生异常体验的原因。可查询14天内数据，查询起止时间不超过1个小时。支持跨天查询。（同老接口DescribeAbnormalEvent）
//
// 异常体验ID映射见：https://cloud.tencent.com/document/product/647/44916
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIMEEXPIRE = "InvalidParameter.StartTimeExpire"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeUnusualEventWithContext(ctx context.Context, request *DescribeUnusualEventRequest) (response *DescribeUnusualEventResponse, err error) {
    if request == nil {
        request = NewDescribeUnusualEventRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeUnusualEvent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUnusualEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUnusualEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserEventRequest() (request *DescribeUserEventRequest) {
    request = &DescribeUserEventRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeUserEvent")
    
    
    return
}

func NewDescribeUserEventResponse() (response *DescribeUserEventResponse) {
    response = &DescribeUserEventResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserEvent
// 查询用户某次通话内的进退房，视频开关等详细事件。可查询14天内数据。（同接口DescribeDetailEvent）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_URLPARAMSERROR = "InvalidParameter.UrlParamsError"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
func (c *Client) DescribeUserEvent(request *DescribeUserEventRequest) (response *DescribeUserEventResponse, err error) {
    return c.DescribeUserEventWithContext(context.Background(), request)
}

// DescribeUserEvent
// 查询用户某次通话内的进退房，视频开关等详细事件。可查询14天内数据。（同接口DescribeDetailEvent）
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ESQUERYERROR = "InternalError.EsQueryError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENDTS = "InvalidParameter.EndTs"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_URLPARAMSERROR = "InvalidParameter.UrlParamsError"
//  INVALIDPARAMETER_USERID = "InvalidParameter.UserId"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_APPID = "MissingParameter.AppId"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
func (c *Client) DescribeUserEventWithContext(ctx context.Context, request *DescribeUserEventRequest) (response *DescribeUserEventResponse, err error) {
    if request == nil {
        request = NewDescribeUserEventRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeUserEvent")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserEvent require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserEventResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeUserInfoRequest() (request *DescribeUserInfoRequest) {
    request = &DescribeUserInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeUserInfo")
    
    
    return
}

func NewDescribeUserInfoResponse() (response *DescribeUserInfoResponse) {
    response = &DescribeUserInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeUserInfo
// 查询指定时间内的用户列表，最大可查询14天内数据，查询起止时间不超过4小时。默认每页查询6个用户，支持每页最大查询100个用户PageSize不超过100）。（同老接口DescribeUserInformation）
//
// **注意**：
//
// 1.该接口只用于历史数据统计或核对数据使用，实时类关键业务逻辑不能使用。
//
// 2.该接口自2024年4月1日起正式商业化，需订阅套餐解锁调用能力，提供以下两种解锁方式，可任选其一解锁：
//
// 方式一：通过订阅[包月套餐](https://cloud.tencent.com/document/product/647/85386)「尊享版」（可查近7天）和「旗舰版」（可查近14天），[前往订阅](https://buy.cloud.tencent.com/trtc?trtcversion=top)。
//
// 方式二：通过订阅[监控仪表盘](https://cloud.tencent.com/document/product/647/81331)商业套餐包「基础版」（可查近7天）和「进阶版」（可查近14天），[前往订阅](https://buy.cloud.tencent.com/trtc_monitor)。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_PAGENUMBER = "InvalidParameter.PageNumber"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIMEOVERSIZE = "InvalidParameter.StartTimeOversize"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_USERIDSMORETHANSIX = "InvalidParameter.UserIdsMorethanSix"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeUserInfo(request *DescribeUserInfoRequest) (response *DescribeUserInfoResponse, err error) {
    return c.DescribeUserInfoWithContext(context.Background(), request)
}

// DescribeUserInfo
// 查询指定时间内的用户列表，最大可查询14天内数据，查询起止时间不超过4小时。默认每页查询6个用户，支持每页最大查询100个用户PageSize不超过100）。（同老接口DescribeUserInformation）
//
// **注意**：
//
// 1.该接口只用于历史数据统计或核对数据使用，实时类关键业务逻辑不能使用。
//
// 2.该接口自2024年4月1日起正式商业化，需订阅套餐解锁调用能力，提供以下两种解锁方式，可任选其一解锁：
//
// 方式一：通过订阅[包月套餐](https://cloud.tencent.com/document/product/647/85386)「尊享版」（可查近7天）和「旗舰版」（可查近14天），[前往订阅](https://buy.cloud.tencent.com/trtc?trtcversion=top)。
//
// 方式二：通过订阅[监控仪表盘](https://cloud.tencent.com/document/product/647/81331)商业套餐包「基础版」（可查近7天）和「进阶版」（可查近14天），[前往订阅](https://buy.cloud.tencent.com/trtc_monitor)。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INTERNALERROR_HTTPPARASEFALIED = "InternalError.HttpParaseFalied"
//  INTERNALERROR_INTERFACEERR = "InternalError.InterfaceErr"
//  INTERNALERROR_METHODERR = "InternalError.MethodErr"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_PAGENUMBER = "InvalidParameter.PageNumber"
//  INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"
//  INVALIDPARAMETER_PAGESIZEOVERSIZE = "InvalidParameter.PageSizeOversize"
//  INVALIDPARAMETER_QUERYSCALEOVERSIZE = "InvalidParameter.QueryScaleOversize"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STARTTIMEOVERSIZE = "InvalidParameter.StartTimeOversize"
//  INVALIDPARAMETER_STARTTS = "InvalidParameter.StartTs"
//  INVALIDPARAMETER_STARTTSOVERSIZE = "InvalidParameter.StartTsOversize"
//  INVALIDPARAMETER_USERIDSMORETHANSIX = "InvalidParameter.UserIdsMorethanSix"
//  MISSINGPARAMETER = "MissingParameter"
//  MISSINGPARAMETER_COMMID = "MissingParameter.CommId"
//  MISSINGPARAMETER_COMMIDORSDKAPPID = "MissingParameter.CommIdOrSdkAppId"
//  MISSINGPARAMETER_ENDTS = "MissingParameter.EndTs"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STARTTS = "MissingParameter.StartTs"
func (c *Client) DescribeUserInfoWithContext(ctx context.Context, request *DescribeUserInfoRequest) (response *DescribeUserInfoResponse, err error) {
    if request == nil {
        request = NewDescribeUserInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeUserInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVoicePrintRequest() (request *DescribeVoicePrintRequest) {
    request = &DescribeVoicePrintRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeVoicePrint")
    
    
    return
}

func NewDescribeVoicePrintResponse() (response *DescribeVoicePrintResponse) {
    response = &DescribeVoicePrintResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVoicePrint
// 查询先前注册的声纹信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONFAILED = "FailedOperation.DbOperationFailed"
//  INVALIDPARAMETER_INVALIDVOICEPRINTIDLIST = "InvalidParameter.InvalidVoicePrintIdList"
func (c *Client) DescribeVoicePrint(request *DescribeVoicePrintRequest) (response *DescribeVoicePrintResponse, err error) {
    return c.DescribeVoicePrintWithContext(context.Background(), request)
}

// DescribeVoicePrint
// 查询先前注册的声纹信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_DBOPERATIONFAILED = "FailedOperation.DbOperationFailed"
//  INVALIDPARAMETER_INVALIDVOICEPRINTIDLIST = "InvalidParameter.InvalidVoicePrintIdList"
func (c *Client) DescribeVoicePrintWithContext(ctx context.Context, request *DescribeVoicePrintRequest) (response *DescribeVoicePrintResponse, err error) {
    if request == nil {
        request = NewDescribeVoicePrintRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeVoicePrint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVoicePrint require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVoicePrintResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeWebRecordRequest() (request *DescribeWebRecordRequest) {
    request = &DescribeWebRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DescribeWebRecord")
    
    
    return
}

func NewDescribeWebRecordResponse() (response *DescribeWebRecordResponse) {
    response = &DescribeWebRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeWebRecord
// 查询页面录制任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) DescribeWebRecord(request *DescribeWebRecordRequest) (response *DescribeWebRecordResponse, err error) {
    return c.DescribeWebRecordWithContext(context.Background(), request)
}

// DescribeWebRecord
// 查询页面录制任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) DescribeWebRecordWithContext(ctx context.Context, request *DescribeWebRecordRequest) (response *DescribeWebRecordResponse, err error) {
    if request == nil {
        request = NewDescribeWebRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DescribeWebRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeWebRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeWebRecordResponse()
    err = c.Send(request, response)
    return
}

func NewDismissRoomRequest() (request *DismissRoomRequest) {
    request = &DismissRoomRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DismissRoom")
    
    
    return
}

func NewDismissRoomResponse() (response *DismissRoomResponse) {
    response = &DismissRoomResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DismissRoom
// 接口说明：把房间所有用户从房间移出，解散房间。支持所有平台，Android、iOS、Windows 和 macOS 需升级到 TRTC SDK 6.6及以上版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DismissRoom(request *DismissRoomRequest) (response *DismissRoomResponse, err error) {
    return c.DismissRoomWithContext(context.Background(), request)
}

// DismissRoom
// 接口说明：把房间所有用户从房间移出，解散房间。支持所有平台，Android、iOS、Windows 和 macOS 需升级到 TRTC SDK 6.6及以上版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DismissRoomWithContext(ctx context.Context, request *DismissRoomRequest) (response *DismissRoomResponse, err error) {
    if request == nil {
        request = NewDismissRoomRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DismissRoom")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DismissRoom require credential")
    }

    request.SetContext(ctx)
    
    response = NewDismissRoomResponse()
    err = c.Send(request, response)
    return
}

func NewDismissRoomByStrRoomIdRequest() (request *DismissRoomByStrRoomIdRequest) {
    request = &DismissRoomByStrRoomIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "DismissRoomByStrRoomId")
    
    
    return
}

func NewDismissRoomByStrRoomIdResponse() (response *DismissRoomByStrRoomIdResponse) {
    response = &DismissRoomByStrRoomIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DismissRoomByStrRoomId
// 接口说明：把房间所有用户从房间移出，解散房间。支持所有平台，Android、iOS、Windows 和 macOS 需升级到 TRTC SDK 6.6及以上版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DismissRoomByStrRoomId(request *DismissRoomByStrRoomIdRequest) (response *DismissRoomByStrRoomIdResponse, err error) {
    return c.DismissRoomByStrRoomIdWithContext(context.Background(), request)
}

// DismissRoomByStrRoomId
// 接口说明：把房间所有用户从房间移出，解散房间。支持所有平台，Android、iOS、Windows 和 macOS 需升级到 TRTC SDK 6.6及以上版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) DismissRoomByStrRoomIdWithContext(ctx context.Context, request *DismissRoomByStrRoomIdRequest) (response *DismissRoomByStrRoomIdResponse, err error) {
    if request == nil {
        request = NewDismissRoomByStrRoomIdRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "DismissRoomByStrRoomId")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DismissRoomByStrRoomId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDismissRoomByStrRoomIdResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudModerationRequest() (request *ModifyCloudModerationRequest) {
    request = &ModifyCloudModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "ModifyCloudModeration")
    
    
    return
}

func NewModifyCloudModerationResponse() (response *ModifyCloudModerationResponse) {
    response = &ModifyCloudModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudModeration
// 成功开启云端审核任务后，可以使用此接口来更新订阅黑白名单。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCloudModeration(request *ModifyCloudModerationRequest) (response *ModifyCloudModerationResponse, err error) {
    return c.ModifyCloudModerationWithContext(context.Background(), request)
}

// ModifyCloudModeration
// 成功开启云端审核任务后，可以使用此接口来更新订阅黑白名单。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCloudModerationWithContext(ctx context.Context, request *ModifyCloudModerationRequest) (response *ModifyCloudModerationResponse, err error) {
    if request == nil {
        request = NewModifyCloudModerationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "ModifyCloudModeration")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudModeration require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudModerationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudRecordingRequest() (request *ModifyCloudRecordingRequest) {
    request = &ModifyCloudRecordingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "ModifyCloudRecording")
    
    
    return
}

func NewModifyCloudRecordingResponse() (response *ModifyCloudRecordingResponse) {
    response = &ModifyCloudRecordingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudRecording
// 成功开启录制后，可以使用此接口来更新录制任务。仅在录制任务进行时有效，录制退出后更新将会返回错误。更新操作是全量覆盖，并不是增量更新的模式，也就是说每次更新都需要携带全量的信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCloudRecording(request *ModifyCloudRecordingRequest) (response *ModifyCloudRecordingResponse, err error) {
    return c.ModifyCloudRecordingWithContext(context.Background(), request)
}

// ModifyCloudRecording
// 成功开启录制后，可以使用此接口来更新录制任务。仅在录制任务进行时有效，录制退出后更新将会返回错误。更新操作是全量覆盖，并不是增量更新的模式，也就是说每次更新都需要携带全量的信息。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR_CRINTERNALERROR = "InternalError.CRInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCloudRecordingWithContext(ctx context.Context, request *ModifyCloudRecordingRequest) (response *ModifyCloudRecordingResponse, err error) {
    if request == nil {
        request = NewModifyCloudRecordingRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "ModifyCloudRecording")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudRecordingResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCloudSliceTaskRequest() (request *ModifyCloudSliceTaskRequest) {
    request = &ModifyCloudSliceTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "ModifyCloudSliceTask")
    
    
    return
}

func NewModifyCloudSliceTaskResponse() (response *ModifyCloudSliceTaskResponse) {
    response = &ModifyCloudSliceTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyCloudSliceTask
// 成功开启切片任务后，可以使用此接口来更新任务。用于更新指定订阅流白名单或者黑名单。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCloudSliceTask(request *ModifyCloudSliceTaskRequest) (response *ModifyCloudSliceTaskResponse, err error) {
    return c.ModifyCloudSliceTaskWithContext(context.Background(), request)
}

// ModifyCloudSliceTask
// 成功开启切片任务后，可以使用此接口来更新任务。用于更新指定订阅流白名单或者黑名单。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_CSUNSUPPORTMETHOD = "FailedOperation.CSUnsupportMethod"
//  INTERNALERROR_CSINTERNALERROR = "InternalError.CSInternalError"
//  INVALIDPARAMETER_OUTOFRANGE = "InvalidParameter.OutOfRange"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  MISSINGPARAMETER_USERID = "MissingParameter.UserId"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) ModifyCloudSliceTaskWithContext(ctx context.Context, request *ModifyCloudSliceTaskRequest) (response *ModifyCloudSliceTaskResponse, err error) {
    if request == nil {
        request = NewModifyCloudSliceTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "ModifyCloudSliceTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudSliceTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudSliceTaskResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPictureRequest() (request *ModifyPictureRequest) {
    request = &ModifyPictureRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "ModifyPicture")
    
    
    return
}

func NewModifyPictureResponse() (response *ModifyPictureResponse) {
    response = &ModifyPictureResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ModifyPicture
// 如果您需要在 [云端混流转码](https://cloud.tencent.com/document/product/647/16827) 时频繁修改自定义背景图或水印素材，可通过此接口修改已上传的图片。无需频繁修改图片素材的场景，建议直接在 [控制台 > 应用管理 > 素材管理](https://cloud.tencent.com/document/product/647/50769) 中操作。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyPicture(request *ModifyPictureRequest) (response *ModifyPictureResponse, err error) {
    return c.ModifyPictureWithContext(context.Background(), request)
}

// ModifyPicture
// 如果您需要在 [云端混流转码](https://cloud.tencent.com/document/product/647/16827) 时频繁修改自定义背景图或水印素材，可通过此接口修改已上传的图片。无需频繁修改图片素材的场景，建议直接在 [控制台 > 应用管理 > 素材管理](https://cloud.tencent.com/document/product/647/50769) 中操作。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) ModifyPictureWithContext(ctx context.Context, request *ModifyPictureRequest) (response *ModifyPictureResponse, err error) {
    if request == nil {
        request = NewModifyPictureRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "ModifyPicture")
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPicture require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPictureResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterVoicePrintRequest() (request *RegisterVoicePrintRequest) {
    request = &RegisterVoicePrintRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "RegisterVoicePrint")
    
    
    return
}

func NewRegisterVoicePrintResponse() (response *RegisterVoicePrintResponse) {
    response = &RegisterVoicePrintResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RegisterVoicePrint
// 传入音频base64串，注册声纹信息，返回声纹ID
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSOPERATIONFAILED = "FailedOperation.CosOperationFailed"
//  FAILEDOPERATION_DBOPERATIONFAILED = "FailedOperation.DbOperationFailed"
//  FAILEDOPERATION_GENVOICEPRINTIDFAILED = "FailedOperation.GenVoicePrintIdFailed"
//  FAILEDOPERATION_VOICEPRINTAUDIOCHECKFAILED = "FailedOperation.VoicePrintAudioCheckFailed"
//  FAILEDOPERATION_VOICEPRINTREGISTRATIONLIMIT = "FailedOperation.VoicePrintRegistrationLimit"
//  INVALIDPARAMETER_INVALIDAUDIOINPUT = "InvalidParameter.InvalidAudioInput"
func (c *Client) RegisterVoicePrint(request *RegisterVoicePrintRequest) (response *RegisterVoicePrintResponse, err error) {
    return c.RegisterVoicePrintWithContext(context.Background(), request)
}

// RegisterVoicePrint
// 传入音频base64串，注册声纹信息，返回声纹ID
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSOPERATIONFAILED = "FailedOperation.CosOperationFailed"
//  FAILEDOPERATION_DBOPERATIONFAILED = "FailedOperation.DbOperationFailed"
//  FAILEDOPERATION_GENVOICEPRINTIDFAILED = "FailedOperation.GenVoicePrintIdFailed"
//  FAILEDOPERATION_VOICEPRINTAUDIOCHECKFAILED = "FailedOperation.VoicePrintAudioCheckFailed"
//  FAILEDOPERATION_VOICEPRINTREGISTRATIONLIMIT = "FailedOperation.VoicePrintRegistrationLimit"
//  INVALIDPARAMETER_INVALIDAUDIOINPUT = "InvalidParameter.InvalidAudioInput"
func (c *Client) RegisterVoicePrintWithContext(ctx context.Context, request *RegisterVoicePrintRequest) (response *RegisterVoicePrintResponse, err error) {
    if request == nil {
        request = NewRegisterVoicePrintRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "RegisterVoicePrint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RegisterVoicePrint require credential")
    }

    request.SetContext(ctx)
    
    response = NewRegisterVoicePrintResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveUserRequest() (request *RemoveUserRequest) {
    request = &RemoveUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "RemoveUser")
    
    
    return
}

func NewRemoveUserResponse() (response *RemoveUserResponse) {
    response = &RemoveUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveUser
// 接口说明：将用户从房间移出，适用于主播/房主/管理员踢人等场景。支持所有平台，Android、iOS、Windows 和 macOS 需升级到 TRTC SDK 6.6及以上版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_USERIDS = "InvalidParameter.UserIds"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_USERIDS = "MissingParameter.UserIds"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) RemoveUser(request *RemoveUserRequest) (response *RemoveUserResponse, err error) {
    return c.RemoveUserWithContext(context.Background(), request)
}

// RemoveUser
// 接口说明：将用户从房间移出，适用于主播/房主/管理员踢人等场景。支持所有平台，Android、iOS、Windows 和 macOS 需升级到 TRTC SDK 6.6及以上版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_USERIDS = "InvalidParameter.UserIds"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_USERIDS = "MissingParameter.UserIds"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) RemoveUserWithContext(ctx context.Context, request *RemoveUserRequest) (response *RemoveUserResponse, err error) {
    if request == nil {
        request = NewRemoveUserRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "RemoveUser")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveUser require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveUserResponse()
    err = c.Send(request, response)
    return
}

func NewRemoveUserByStrRoomIdRequest() (request *RemoveUserByStrRoomIdRequest) {
    request = &RemoveUserByStrRoomIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "RemoveUserByStrRoomId")
    
    
    return
}

func NewRemoveUserByStrRoomIdResponse() (response *RemoveUserByStrRoomIdResponse) {
    response = &RemoveUserByStrRoomIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RemoveUserByStrRoomId
// 接口说明：将用户从房间移出，适用于主播/房主/管理员踢人等场景。支持所有平台，Android、iOS、Windows 和 macOS 需升级到 TRTC SDK 6.6及以上版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_USERIDS = "InvalidParameter.UserIds"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_USERIDS = "MissingParameter.UserIds"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) RemoveUserByStrRoomId(request *RemoveUserByStrRoomIdRequest) (response *RemoveUserByStrRoomIdResponse, err error) {
    return c.RemoveUserByStrRoomIdWithContext(context.Background(), request)
}

// RemoveUserByStrRoomId
// 接口说明：将用户从房间移出，适用于主播/房主/管理员踢人等场景。支持所有平台，Android、iOS、Windows 和 macOS 需升级到 TRTC SDK 6.6及以上版本。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMCACHEIPERROR = "InternalError.GetRoomCacheIpError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_USERIDS = "InvalidParameter.UserIds"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_USERIDS = "MissingParameter.UserIds"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) RemoveUserByStrRoomIdWithContext(ctx context.Context, request *RemoveUserByStrRoomIdRequest) (response *RemoveUserByStrRoomIdResponse, err error) {
    if request == nil {
        request = NewRemoveUserByStrRoomIdRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "RemoveUserByStrRoomId")
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveUserByStrRoomId require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveUserByStrRoomIdResponse()
    err = c.Send(request, response)
    return
}

func NewStartAIConversationRequest() (request *StartAIConversationRequest) {
    request = &StartAIConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StartAIConversation")
    
    
    return
}

func NewStartAIConversationResponse() (response *StartAIConversationResponse) {
    response = &StartAIConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartAIConversation
// 启动AI对话任务，AI通道机器人进入TRTC房间，与房间内指定的成员进行AI对话，适用于智能客服，AI口语教师等场景
//
// 
//
// TRTC AI对话功能内置语音转文本能力，同时提供通道服务，即客户可灵活指定第三方AI模型（LLM）服务和文本转音频（TTS)服务，更多[功能说明](https://cloud.tencent.com/document/product/647/108901)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTABILITY = "FailedOperation.NotAbility"
//  FAILEDOPERATION_NOTALLOWED = "FailedOperation.NotAllowed"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  INVALIDPARAMETER_USERSIG = "InvalidParameter.UserSig"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
func (c *Client) StartAIConversation(request *StartAIConversationRequest) (response *StartAIConversationResponse, err error) {
    return c.StartAIConversationWithContext(context.Background(), request)
}

// StartAIConversation
// 启动AI对话任务，AI通道机器人进入TRTC房间，与房间内指定的成员进行AI对话，适用于智能客服，AI口语教师等场景
//
// 
//
// TRTC AI对话功能内置语音转文本能力，同时提供通道服务，即客户可灵活指定第三方AI模型（LLM）服务和文本转音频（TTS)服务，更多[功能说明](https://cloud.tencent.com/document/product/647/108901)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTABILITY = "FailedOperation.NotAbility"
//  FAILEDOPERATION_NOTALLOWED = "FailedOperation.NotAllowed"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  INVALIDPARAMETER_USERSIG = "InvalidParameter.UserSig"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
func (c *Client) StartAIConversationWithContext(ctx context.Context, request *StartAIConversationRequest) (response *StartAIConversationResponse, err error) {
    if request == nil {
        request = NewStartAIConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StartAIConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartAIConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartAIConversationResponse()
    err = c.Send(request, response)
    return
}

func NewStartAITranscriptionRequest() (request *StartAITranscriptionRequest) {
    request = &StartAITranscriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StartAITranscription")
    
    
    return
}

func NewStartAITranscriptionResponse() (response *StartAITranscriptionResponse) {
    response = &StartAITranscriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartAITranscription
// 启动转录机器人，后台会通过机器人拉流进行实时进行语音识别并下发字幕和转录消息。
//
// 转录机器人支持两种拉流方式，通过TranscriptionMode字段控制：
//
// - 拉取全房间的流。
//
// - 拉取特定用户的流。
//
// 
//
// 服务端通过TRTC的自定义消息实时下发字幕以及转录消息，CmdId固定是1。客户端只需监听自定义消息的回调即可，比如[c++回调](https://cloud.tencent.com/document/product/647/79637#4cd82f4edb24992a15a25187089e1565)。其他客户端比如安卓、Web等同样可在该链接处找到。
//
// 
//
// 
//
// **注意：**
//
// TranscriptionMode为0时，需要保证一个房间内只发起一个任务，如果发起多个任务，则机器人之间会相互订阅，除非主动停止任务，否则只有10小时后任务才会超时退出，这种情况下建议填写SessionId，保证后续重复发起的任务失败。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTABILITY = "FailedOperation.NotAbility"
//  FAILEDOPERATION_NOTALLOWED = "FailedOperation.NotAllowed"
//  FAILEDOPERATION_SDKAPPIDNOTUNDERAPPID = "FailedOperation.SdkAppIdNotUnderAppId"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  INVALIDPARAMETER_USERSIG = "InvalidParameter.UserSig"
//  INVALIDPARAMETER_USERSIGNOTADMIN = "InvalidParameter.UserSigNotAdmin"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
func (c *Client) StartAITranscription(request *StartAITranscriptionRequest) (response *StartAITranscriptionResponse, err error) {
    return c.StartAITranscriptionWithContext(context.Background(), request)
}

// StartAITranscription
// 启动转录机器人，后台会通过机器人拉流进行实时进行语音识别并下发字幕和转录消息。
//
// 转录机器人支持两种拉流方式，通过TranscriptionMode字段控制：
//
// - 拉取全房间的流。
//
// - 拉取特定用户的流。
//
// 
//
// 服务端通过TRTC的自定义消息实时下发字幕以及转录消息，CmdId固定是1。客户端只需监听自定义消息的回调即可，比如[c++回调](https://cloud.tencent.com/document/product/647/79637#4cd82f4edb24992a15a25187089e1565)。其他客户端比如安卓、Web等同样可在该链接处找到。
//
// 
//
// 
//
// **注意：**
//
// TranscriptionMode为0时，需要保证一个房间内只发起一个任务，如果发起多个任务，则机器人之间会相互订阅，除非主动停止任务，否则只有10小时后任务才会超时退出，这种情况下建议填写SessionId，保证后续重复发起的任务失败。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTABILITY = "FailedOperation.NotAbility"
//  FAILEDOPERATION_NOTALLOWED = "FailedOperation.NotAllowed"
//  FAILEDOPERATION_SDKAPPIDNOTUNDERAPPID = "FailedOperation.SdkAppIdNotUnderAppId"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  INVALIDPARAMETER_USERSIG = "InvalidParameter.UserSig"
//  INVALIDPARAMETER_USERSIGNOTADMIN = "InvalidParameter.UserSigNotAdmin"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
func (c *Client) StartAITranscriptionWithContext(ctx context.Context, request *StartAITranscriptionRequest) (response *StartAITranscriptionResponse, err error) {
    if request == nil {
        request = NewStartAITranscriptionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StartAITranscription")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartAITranscription require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartAITranscriptionResponse()
    err = c.Send(request, response)
    return
}

func NewStartMCUMixTranscodeRequest() (request *StartMCUMixTranscodeRequest) {
    request = &StartMCUMixTranscodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StartMCUMixTranscode")
    
    
    return
}

func NewStartMCUMixTranscodeResponse() (response *StartMCUMixTranscodeResponse) {
    response = &StartMCUMixTranscodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartMCUMixTranscode
// 接口说明：启动云端混流，并指定混流画面中各路画面的布局位置。
//
// 
//
// TRTC 的一个房间中可能会同时存在多路音视频流，您可以通过此 API 接口，通知腾讯云服务端将多路视频画面合成一路，并指定每一路画面的位置，同时将多路声音进行混音，最终形成一路音视频流，以便用于录制和直播观看。房间销毁后混流自动结束。
//
// 
//
// 您可以通过此接口实现如下目标：
//
// - 设置最终直播流的画质和音质，包括视频分辨率、视频码率、视频帧率、以及声音质量等。
//
// - 设置各路画面的位置和布局，您只需要在启动时设置一次，排版引擎会自动完成后续的画面排布。
//
// - 设置录制文件名，用于二次回放。
//
// - 设置 CDN 直播流 ID，用于在 CDN 进行直播观看。
//
// 
//
// 目前已经支持了如下几种布局模板：
//
// - 悬浮模板：第一个进入房间的用户的视频画面会铺满整个屏幕，其他用户的视频画面从左下角依次水平排列，显示为小画面，最多4行，每行4个，小画面悬浮于大画面之上。最多支持1个大画面和15个小画面，如果用户只发送音频，仍然会占用画面位置。
//
// - 九宫格模板：所有用户的视频画面大小一致，平分整个屏幕，人数越多，每个画面的尺寸越小。最多支持16个画面，如果用户只发送音频，仍然会占用画面位置。
//
// - 屏幕分享模板：适合视频会议和在线教育场景的布局，屏幕分享（或者主讲的摄像头）始终占据屏幕左侧的大画面位置，其他用户依次垂直排列于右侧，最多两列，每列最多8个小画面。最多支持1个大画面和15个小画面。若上行分辨率宽高比与画面输出宽高比不一致时，左侧大画面为了保持内容的完整性采用缩放方式处理，右侧小画面采用裁剪方式处理。
//
// - 画中画模板：适用于混合大小两路视频画面和其他用户混音，或者混合一路大画面和其他用户混音的场景。小画面悬浮于大画面之上，可以指定大小画面的用户以及小画面的显示位置，最多支持2个画面。
//
// - 自定义模板：适用于在混流中指定用户的画面位置，或者预设视频画面位置的场景。当预设位置指定用户时，排版引擎会为该用户预留位置；当预设位置未指定用户时，排版引擎会根据进房间顺序自动填充。预设位置填满时，不再混合其他用户的画面和声音。自定义模板启用占位图功能时（LayoutParams中的PlaceHolderMode设置成1），在预设位置的用户没有上行视频时可显示对应的占位图（PlaceImageId）。
//
// 
//
// 注意：
//
// 1、**混流转码为收费功能，调用接口将产生云端混流转码费用，详见[云端混流转码计费说明](https://cloud.tencent.com/document/product/647/49446)。**
//
// 2、2020年1月9号及以后创建的应用才能直接调用此接口。2020年1月9日之前创建的应用默认使用云直播的云端混流，如需切换至MCU混流，请[提交工单](https://console.cloud.tencent.com/workorder/category)寻求帮助。
//
// 3、客户端混流和服务端混流不能混用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  FAILEDOPERATION_TASKFINISHED = "FailedOperation.TaskFinished"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_AUDIOENCODEPARAMS = "InvalidParameter.AudioEncodeParams"
//  INVALIDPARAMETER_BACKGROUNDIMAGEURL = "InvalidParameter.BackgroundImageUrl"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_MAINVIDEORIGHTALIGN = "InvalidParameter.MainVideoRightAlign"
//  INVALIDPARAMETER_MAINVIDEOSTREAMTYPE = "InvalidParameter.MainVideoStreamType"
//  INVALIDPARAMETER_OUTPUTPARAMS = "InvalidParameter.OutputParams"
//  INVALIDPARAMETER_PRESETLAYOUTCONFIG = "InvalidParameter.PresetLayoutConfig"
//  INVALIDPARAMETER_PUBLISHCDNURLS = "InvalidParameter.PublishCdnUrls"
//  INVALIDPARAMETER_PUREAUDIOSTREAM = "InvalidParameter.PureAudioStream"
//  INVALIDPARAMETER_RECORDAUDIOONLY = "InvalidParameter.RecordAudioOnly"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_SMALLVIDEOLAYOUTPARAMS = "InvalidParameter.SmallVideoLayoutParams"
//  INVALIDPARAMETER_SMALLVIDEOSTREAMTYPE = "InvalidParameter.SmallVideoStreamType"
//  INVALIDPARAMETER_STREAMID = "InvalidParameter.StreamId"
//  INVALIDPARAMETER_VIDEORESOLUTION = "InvalidParameter.VideoResolution"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_AUDIOENCODEPARAMS = "MissingParameter.AudioEncodeParams"
//  MISSINGPARAMETER_BIZID = "MissingParameter.BizId"
//  MISSINGPARAMETER_ENCODEPARAMS = "MissingParameter.EncodeParams"
//  MISSINGPARAMETER_OUTPUTPARAMS = "MissingParameter.OutputParams"
//  MISSINGPARAMETER_PRESETLAYOUTCONFIG = "MissingParameter.PresetLayoutConfig"
//  MISSINGPARAMETER_PUBLISHCDNPARAMS = "MissingParameter.PublishCdnParams"
//  MISSINGPARAMETER_PUBLISHCDNURLS = "MissingParameter.PublishCdnUrls"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STREAMID = "MissingParameter.StreamId"
//  MISSINGPARAMETER_VIDEOENCODEPARAMS = "MissingParameter.VideoEncodeParams"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartMCUMixTranscode(request *StartMCUMixTranscodeRequest) (response *StartMCUMixTranscodeResponse, err error) {
    return c.StartMCUMixTranscodeWithContext(context.Background(), request)
}

// StartMCUMixTranscode
// 接口说明：启动云端混流，并指定混流画面中各路画面的布局位置。
//
// 
//
// TRTC 的一个房间中可能会同时存在多路音视频流，您可以通过此 API 接口，通知腾讯云服务端将多路视频画面合成一路，并指定每一路画面的位置，同时将多路声音进行混音，最终形成一路音视频流，以便用于录制和直播观看。房间销毁后混流自动结束。
//
// 
//
// 您可以通过此接口实现如下目标：
//
// - 设置最终直播流的画质和音质，包括视频分辨率、视频码率、视频帧率、以及声音质量等。
//
// - 设置各路画面的位置和布局，您只需要在启动时设置一次，排版引擎会自动完成后续的画面排布。
//
// - 设置录制文件名，用于二次回放。
//
// - 设置 CDN 直播流 ID，用于在 CDN 进行直播观看。
//
// 
//
// 目前已经支持了如下几种布局模板：
//
// - 悬浮模板：第一个进入房间的用户的视频画面会铺满整个屏幕，其他用户的视频画面从左下角依次水平排列，显示为小画面，最多4行，每行4个，小画面悬浮于大画面之上。最多支持1个大画面和15个小画面，如果用户只发送音频，仍然会占用画面位置。
//
// - 九宫格模板：所有用户的视频画面大小一致，平分整个屏幕，人数越多，每个画面的尺寸越小。最多支持16个画面，如果用户只发送音频，仍然会占用画面位置。
//
// - 屏幕分享模板：适合视频会议和在线教育场景的布局，屏幕分享（或者主讲的摄像头）始终占据屏幕左侧的大画面位置，其他用户依次垂直排列于右侧，最多两列，每列最多8个小画面。最多支持1个大画面和15个小画面。若上行分辨率宽高比与画面输出宽高比不一致时，左侧大画面为了保持内容的完整性采用缩放方式处理，右侧小画面采用裁剪方式处理。
//
// - 画中画模板：适用于混合大小两路视频画面和其他用户混音，或者混合一路大画面和其他用户混音的场景。小画面悬浮于大画面之上，可以指定大小画面的用户以及小画面的显示位置，最多支持2个画面。
//
// - 自定义模板：适用于在混流中指定用户的画面位置，或者预设视频画面位置的场景。当预设位置指定用户时，排版引擎会为该用户预留位置；当预设位置未指定用户时，排版引擎会根据进房间顺序自动填充。预设位置填满时，不再混合其他用户的画面和声音。自定义模板启用占位图功能时（LayoutParams中的PlaceHolderMode设置成1），在预设位置的用户没有上行视频时可显示对应的占位图（PlaceImageId）。
//
// 
//
// 注意：
//
// 1、**混流转码为收费功能，调用接口将产生云端混流转码费用，详见[云端混流转码计费说明](https://cloud.tencent.com/document/product/647/49446)。**
//
// 2、2020年1月9号及以后创建的应用才能直接调用此接口。2020年1月9日之前创建的应用默认使用云直播的云端混流，如需切换至MCU混流，请[提交工单](https://console.cloud.tencent.com/workorder/category)寻求帮助。
//
// 3、客户端混流和服务端混流不能混用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  FAILEDOPERATION_TASKFINISHED = "FailedOperation.TaskFinished"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_AUDIOENCODEPARAMS = "InvalidParameter.AudioEncodeParams"
//  INVALIDPARAMETER_BACKGROUNDIMAGEURL = "InvalidParameter.BackgroundImageUrl"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_MAINVIDEORIGHTALIGN = "InvalidParameter.MainVideoRightAlign"
//  INVALIDPARAMETER_MAINVIDEOSTREAMTYPE = "InvalidParameter.MainVideoStreamType"
//  INVALIDPARAMETER_OUTPUTPARAMS = "InvalidParameter.OutputParams"
//  INVALIDPARAMETER_PRESETLAYOUTCONFIG = "InvalidParameter.PresetLayoutConfig"
//  INVALIDPARAMETER_PUBLISHCDNURLS = "InvalidParameter.PublishCdnUrls"
//  INVALIDPARAMETER_PUREAUDIOSTREAM = "InvalidParameter.PureAudioStream"
//  INVALIDPARAMETER_RECORDAUDIOONLY = "InvalidParameter.RecordAudioOnly"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_SMALLVIDEOLAYOUTPARAMS = "InvalidParameter.SmallVideoLayoutParams"
//  INVALIDPARAMETER_SMALLVIDEOSTREAMTYPE = "InvalidParameter.SmallVideoStreamType"
//  INVALIDPARAMETER_STREAMID = "InvalidParameter.StreamId"
//  INVALIDPARAMETER_VIDEORESOLUTION = "InvalidParameter.VideoResolution"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_AUDIOENCODEPARAMS = "MissingParameter.AudioEncodeParams"
//  MISSINGPARAMETER_BIZID = "MissingParameter.BizId"
//  MISSINGPARAMETER_ENCODEPARAMS = "MissingParameter.EncodeParams"
//  MISSINGPARAMETER_OUTPUTPARAMS = "MissingParameter.OutputParams"
//  MISSINGPARAMETER_PRESETLAYOUTCONFIG = "MissingParameter.PresetLayoutConfig"
//  MISSINGPARAMETER_PUBLISHCDNPARAMS = "MissingParameter.PublishCdnParams"
//  MISSINGPARAMETER_PUBLISHCDNURLS = "MissingParameter.PublishCdnUrls"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_STREAMID = "MissingParameter.StreamId"
//  MISSINGPARAMETER_VIDEOENCODEPARAMS = "MissingParameter.VideoEncodeParams"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartMCUMixTranscodeWithContext(ctx context.Context, request *StartMCUMixTranscodeRequest) (response *StartMCUMixTranscodeResponse, err error) {
    if request == nil {
        request = NewStartMCUMixTranscodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StartMCUMixTranscode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartMCUMixTranscode require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartMCUMixTranscodeResponse()
    err = c.Send(request, response)
    return
}

func NewStartMCUMixTranscodeByStrRoomIdRequest() (request *StartMCUMixTranscodeByStrRoomIdRequest) {
    request = &StartMCUMixTranscodeByStrRoomIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StartMCUMixTranscodeByStrRoomId")
    
    
    return
}

func NewStartMCUMixTranscodeByStrRoomIdResponse() (response *StartMCUMixTranscodeByStrRoomIdResponse) {
    response = &StartMCUMixTranscodeByStrRoomIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartMCUMixTranscodeByStrRoomId
// 接口说明：启动云端混流，并指定混流画面中各路画面的布局位置。
//
// 
//
// TRTC 的一个房间中可能会同时存在多路音视频流，您可以通过此 API 接口，通知腾讯云服务端将多路视频画面合成一路，并指定每一路画面的位置，同时将多路声音进行混音，最终形成一路音视频流，以便用于录制和直播观看。
//
// 
//
// 您可以通过此接口实现如下目标：
//
// - 设置最终直播流的画质和音质，包括视频分辨率、视频码率、视频帧率、以及声音质量等。
//
// - 设置各路画面的位置和布局，您只需要在启动时设置一次，排版引擎会自动完成后续的画面排布。
//
// - 设置录制文件名，用于二次回放。
//
// - 设置 CDN 直播流 ID，用于在 CDN 进行直播观看。
//
// 
//
// 目前已经支持了如下几种布局模板：
//
// - 悬浮模板：第一个进入房间的用户的视频画面会铺满整个屏幕，其他用户的视频画面从左下角依次水平排列，显示为小画面，最多4行，每行4个，小画面悬浮于大画面之上。最多支持1个大画面和15个小画面，如果用户只发送音频，仍然会占用画面位置。
//
// - 九宫格模板：所有用户的视频画面大小一致，平分整个屏幕，人数越多，每个画面的尺寸越小。最多支持16个画面，如果用户只发送音频，仍然会占用画面位置。
//
// - 屏幕分享模板：适合视频会议和在线教育场景的布局，屏幕分享（或者主讲的摄像头）始终占据屏幕左侧的大画面位置，其他用户依次垂直排列于右侧，最多两列，每列最多8个小画面。最多支持1个大画面和15个小画面。若上行分辨率宽高比与画面输出宽高比不一致时，左侧大画面为了保持内容的完整性采用缩放方式处理，右侧小画面采用裁剪方式处理。
//
// - 画中画模板：适用于混合大小两路视频画面和其他用户混音，或者混合一路大画面和其他用户混音的场景。小画面悬浮于大画面之上，可以指定大小画面的用户以及小画面的显示位置。
//
// - 自定义模板：适用于在混流中指定用户的画面位置，或者预设视频画面位置的场景。当预设位置指定用户时，排版引擎会为该用户预留位置；当预设位置未指定用户时，排版引擎会根据进房间顺序自动填充。预设位置填满时，不再混合其他用户的画面和声音。自定义模板启用占位图功能时（LayoutParams中的PlaceHolderMode设置成1），在预设位置的用户没有上行视频时可显示对应的占位图（PlaceImageId）。
//
// 
//
// 注意：
//
// 1、**混流转码为收费功能，调用接口将产生云端混流转码费用，详见[云端混流转码计费说明](https://cloud.tencent.com/document/product/647/49446)。**
//
// 2、2020年1月9号及以后创建的应用才能直接调用此接口。2020年1月9日之前创建的应用默认使用云直播的云端混流，如需切换至MCU混流，请[提交工单](https://console.cloud.tencent.com/workorder/category)寻求帮助。
//
// 3、客户端混流和服务端混流不能混用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  FAILEDOPERATION_TASKFINISHED = "FailedOperation.TaskFinished"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_AUDIOENCODEPARAMS = "InvalidParameter.AudioEncodeParams"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_MAINVIDEOSTREAMTYPE = "InvalidParameter.MainVideoStreamType"
//  INVALIDPARAMETER_OUTPUTPARAMS = "InvalidParameter.OutputParams"
//  INVALIDPARAMETER_PRESETLAYOUTCONFIG = "InvalidParameter.PresetLayoutConfig"
//  INVALIDPARAMETER_PUBLISHCDNURLS = "InvalidParameter.PublishCdnUrls"
//  INVALIDPARAMETER_PUREAUDIOSTREAM = "InvalidParameter.PureAudioStream"
//  INVALIDPARAMETER_RECORDAUDIOONLY = "InvalidParameter.RecordAudioOnly"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_SMALLVIDEOLAYOUTPARAMS = "InvalidParameter.SmallVideoLayoutParams"
//  INVALIDPARAMETER_SMALLVIDEOSTREAMTYPE = "InvalidParameter.SmallVideoStreamType"
//  INVALIDPARAMETER_STRROOMID = "InvalidParameter.StrRoomId"
//  INVALIDPARAMETER_STREAMID = "InvalidParameter.StreamId"
//  INVALIDPARAMETER_VIDEORESOLUTION = "InvalidParameter.VideoResolution"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_AUDIOENCODEPARAMS = "MissingParameter.AudioEncodeParams"
//  MISSINGPARAMETER_BIZID = "MissingParameter.BizId"
//  MISSINGPARAMETER_ENCODEPARAMS = "MissingParameter.EncodeParams"
//  MISSINGPARAMETER_OUTPUTPARAMS = "MissingParameter.OutputParams"
//  MISSINGPARAMETER_PRESETLAYOUTCONFIG = "MissingParameter.PresetLayoutConfig"
//  MISSINGPARAMETER_PUBLISHCDNPARAMS = "MissingParameter.PublishCdnParams"
//  MISSINGPARAMETER_PUBLISHCDNURLS = "MissingParameter.PublishCdnUrls"
//  MISSINGPARAMETER_STREAMID = "MissingParameter.StreamId"
//  MISSINGPARAMETER_VIDEOENCODEPARAMS = "MissingParameter.VideoEncodeParams"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartMCUMixTranscodeByStrRoomId(request *StartMCUMixTranscodeByStrRoomIdRequest) (response *StartMCUMixTranscodeByStrRoomIdResponse, err error) {
    return c.StartMCUMixTranscodeByStrRoomIdWithContext(context.Background(), request)
}

// StartMCUMixTranscodeByStrRoomId
// 接口说明：启动云端混流，并指定混流画面中各路画面的布局位置。
//
// 
//
// TRTC 的一个房间中可能会同时存在多路音视频流，您可以通过此 API 接口，通知腾讯云服务端将多路视频画面合成一路，并指定每一路画面的位置，同时将多路声音进行混音，最终形成一路音视频流，以便用于录制和直播观看。
//
// 
//
// 您可以通过此接口实现如下目标：
//
// - 设置最终直播流的画质和音质，包括视频分辨率、视频码率、视频帧率、以及声音质量等。
//
// - 设置各路画面的位置和布局，您只需要在启动时设置一次，排版引擎会自动完成后续的画面排布。
//
// - 设置录制文件名，用于二次回放。
//
// - 设置 CDN 直播流 ID，用于在 CDN 进行直播观看。
//
// 
//
// 目前已经支持了如下几种布局模板：
//
// - 悬浮模板：第一个进入房间的用户的视频画面会铺满整个屏幕，其他用户的视频画面从左下角依次水平排列，显示为小画面，最多4行，每行4个，小画面悬浮于大画面之上。最多支持1个大画面和15个小画面，如果用户只发送音频，仍然会占用画面位置。
//
// - 九宫格模板：所有用户的视频画面大小一致，平分整个屏幕，人数越多，每个画面的尺寸越小。最多支持16个画面，如果用户只发送音频，仍然会占用画面位置。
//
// - 屏幕分享模板：适合视频会议和在线教育场景的布局，屏幕分享（或者主讲的摄像头）始终占据屏幕左侧的大画面位置，其他用户依次垂直排列于右侧，最多两列，每列最多8个小画面。最多支持1个大画面和15个小画面。若上行分辨率宽高比与画面输出宽高比不一致时，左侧大画面为了保持内容的完整性采用缩放方式处理，右侧小画面采用裁剪方式处理。
//
// - 画中画模板：适用于混合大小两路视频画面和其他用户混音，或者混合一路大画面和其他用户混音的场景。小画面悬浮于大画面之上，可以指定大小画面的用户以及小画面的显示位置。
//
// - 自定义模板：适用于在混流中指定用户的画面位置，或者预设视频画面位置的场景。当预设位置指定用户时，排版引擎会为该用户预留位置；当预设位置未指定用户时，排版引擎会根据进房间顺序自动填充。预设位置填满时，不再混合其他用户的画面和声音。自定义模板启用占位图功能时（LayoutParams中的PlaceHolderMode设置成1），在预设位置的用户没有上行视频时可显示对应的占位图（PlaceImageId）。
//
// 
//
// 注意：
//
// 1、**混流转码为收费功能，调用接口将产生云端混流转码费用，详见[云端混流转码计费说明](https://cloud.tencent.com/document/product/647/49446)。**
//
// 2、2020年1月9号及以后创建的应用才能直接调用此接口。2020年1月9日之前创建的应用默认使用云直播的云端混流，如需切换至MCU混流，请[提交工单](https://console.cloud.tencent.com/workorder/category)寻求帮助。
//
// 3、客户端混流和服务端混流不能混用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  FAILEDOPERATION_TASKFINISHED = "FailedOperation.TaskFinished"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_AUDIOENCODEPARAMS = "InvalidParameter.AudioEncodeParams"
//  INVALIDPARAMETER_ENCODEPARAMS = "InvalidParameter.EncodeParams"
//  INVALIDPARAMETER_MAINVIDEOSTREAMTYPE = "InvalidParameter.MainVideoStreamType"
//  INVALIDPARAMETER_OUTPUTPARAMS = "InvalidParameter.OutputParams"
//  INVALIDPARAMETER_PRESETLAYOUTCONFIG = "InvalidParameter.PresetLayoutConfig"
//  INVALIDPARAMETER_PUBLISHCDNURLS = "InvalidParameter.PublishCdnUrls"
//  INVALIDPARAMETER_PUREAUDIOSTREAM = "InvalidParameter.PureAudioStream"
//  INVALIDPARAMETER_RECORDAUDIOONLY = "InvalidParameter.RecordAudioOnly"
//  INVALIDPARAMETER_RECORDID = "InvalidParameter.RecordId"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_SMALLVIDEOLAYOUTPARAMS = "InvalidParameter.SmallVideoLayoutParams"
//  INVALIDPARAMETER_SMALLVIDEOSTREAMTYPE = "InvalidParameter.SmallVideoStreamType"
//  INVALIDPARAMETER_STRROOMID = "InvalidParameter.StrRoomId"
//  INVALIDPARAMETER_STREAMID = "InvalidParameter.StreamId"
//  INVALIDPARAMETER_VIDEORESOLUTION = "InvalidParameter.VideoResolution"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_AUDIOENCODEPARAMS = "MissingParameter.AudioEncodeParams"
//  MISSINGPARAMETER_BIZID = "MissingParameter.BizId"
//  MISSINGPARAMETER_ENCODEPARAMS = "MissingParameter.EncodeParams"
//  MISSINGPARAMETER_OUTPUTPARAMS = "MissingParameter.OutputParams"
//  MISSINGPARAMETER_PRESETLAYOUTCONFIG = "MissingParameter.PresetLayoutConfig"
//  MISSINGPARAMETER_PUBLISHCDNPARAMS = "MissingParameter.PublishCdnParams"
//  MISSINGPARAMETER_PUBLISHCDNURLS = "MissingParameter.PublishCdnUrls"
//  MISSINGPARAMETER_STREAMID = "MissingParameter.StreamId"
//  MISSINGPARAMETER_VIDEOENCODEPARAMS = "MissingParameter.VideoEncodeParams"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StartMCUMixTranscodeByStrRoomIdWithContext(ctx context.Context, request *StartMCUMixTranscodeByStrRoomIdRequest) (response *StartMCUMixTranscodeByStrRoomIdResponse, err error) {
    if request == nil {
        request = NewStartMCUMixTranscodeByStrRoomIdRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StartMCUMixTranscodeByStrRoomId")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartMCUMixTranscodeByStrRoomId require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartMCUMixTranscodeByStrRoomIdResponse()
    err = c.Send(request, response)
    return
}

func NewStartPublishCdnStreamRequest() (request *StartPublishCdnStreamRequest) {
    request = &StartPublishCdnStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StartPublishCdnStream")
    
    
    return
}

func NewStartPublishCdnStreamResponse() (response *StartPublishCdnStreamResponse) {
    response = &StartPublishCdnStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartPublishCdnStream
// 接口说明：  
//
// 启动一个混流转推任务，将  TRTC 房间的多路音视频流混成一路音视频流，编码后推到直播 CDN 或者回推到 TRTC 房间。也支持不转码直接转推 TRTC 房间的单路流。启动成功后，会返回一个 SdkAppid 维度唯一的任务 Id（TaskId）。您需要保存该 TaskId，后续需要依赖此 TaskId 更新和结束任务。可以参考文档： [功能说明](https://cloud.tencent.com/document/product/647/84721#b9a855f4-e38c-4616-9b07-fc44e0e8282a) 和 [常见问题](https://cloud.tencent.com/document/product/647/62620)
//
// 
//
// 注意：
//
// 您可以在控制台开通旁路转推回调功能，对转推 CDN 状态的事件进行监控，回调请参考文档：[旁路转推回调说明](https://cloud.tencent.com/document/product/647/88552)  
//
// 您发起混流转推任务时，可能会产生如下费用：  
//
// MCU 混流转码费用，请参考文档：[云端混流转码计费说明](https://cloud.tencent.com/document/product/647/49446)  
//
// 转推非腾讯云 CDN 费用，请参考文档：[云端转推计费说明](https://cloud.tencent.com/document/product/647/82155)
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartPublishCdnStream(request *StartPublishCdnStreamRequest) (response *StartPublishCdnStreamResponse, err error) {
    return c.StartPublishCdnStreamWithContext(context.Background(), request)
}

// StartPublishCdnStream
// 接口说明：  
//
// 启动一个混流转推任务，将  TRTC 房间的多路音视频流混成一路音视频流，编码后推到直播 CDN 或者回推到 TRTC 房间。也支持不转码直接转推 TRTC 房间的单路流。启动成功后，会返回一个 SdkAppid 维度唯一的任务 Id（TaskId）。您需要保存该 TaskId，后续需要依赖此 TaskId 更新和结束任务。可以参考文档： [功能说明](https://cloud.tencent.com/document/product/647/84721#b9a855f4-e38c-4616-9b07-fc44e0e8282a) 和 [常见问题](https://cloud.tencent.com/document/product/647/62620)
//
// 
//
// 注意：
//
// 您可以在控制台开通旁路转推回调功能，对转推 CDN 状态的事件进行监控，回调请参考文档：[旁路转推回调说明](https://cloud.tencent.com/document/product/647/88552)  
//
// 您发起混流转推任务时，可能会产生如下费用：  
//
// MCU 混流转码费用，请参考文档：[云端混流转码计费说明](https://cloud.tencent.com/document/product/647/49446)  
//
// 转推非腾讯云 CDN 费用，请参考文档：[云端转推计费说明](https://cloud.tencent.com/document/product/647/82155)
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) StartPublishCdnStreamWithContext(ctx context.Context, request *StartPublishCdnStreamRequest) (response *StartPublishCdnStreamResponse, err error) {
    if request == nil {
        request = NewStartPublishCdnStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StartPublishCdnStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartPublishCdnStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartPublishCdnStreamResponse()
    err = c.Send(request, response)
    return
}

func NewStartStreamIngestRequest() (request *StartStreamIngestRequest) {
    request = &StartStreamIngestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StartStreamIngest")
    
    
    return
}

func NewStartStreamIngestResponse() (response *StartStreamIngestResponse) {
    response = &StartStreamIngestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartStreamIngest
// 将一个在线媒体流推到TRTC房间，更多功能说明见[输入媒体流进房](https://cloud.tencent.com/document/product/647/102957#50940aad-d90f-4473-9f46-d5dd46917653)。
//
// 使用输入在线媒体流功能需先订阅 [尊享版或旗舰版套餐包](https://cloud.tencent.com/document/product/647/85386) 解锁能力位。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWED = "FailedOperation.NotAllowed"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  INTERNALERROR_HTTPPARSEFAILED = "InternalError.HttpParseFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STRROOMID = "InvalidParameter.StrRoomId"
//  INVALIDPARAMETER_STREAMURL = "InvalidParameter.StreamUrl"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  INVALIDPARAMETER_USERSIG = "InvalidParameter.UserSig"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
func (c *Client) StartStreamIngest(request *StartStreamIngestRequest) (response *StartStreamIngestResponse, err error) {
    return c.StartStreamIngestWithContext(context.Background(), request)
}

// StartStreamIngest
// 将一个在线媒体流推到TRTC房间，更多功能说明见[输入媒体流进房](https://cloud.tencent.com/document/product/647/102957#50940aad-d90f-4473-9f46-d5dd46917653)。
//
// 使用输入在线媒体流功能需先订阅 [尊享版或旗舰版套餐包](https://cloud.tencent.com/document/product/647/85386) 解锁能力位。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTALLOWED = "FailedOperation.NotAllowed"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  INTERNALERROR_HTTPPARSEFAILED = "InternalError.HttpParseFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STRROOMID = "InvalidParameter.StrRoomId"
//  INVALIDPARAMETER_STREAMURL = "InvalidParameter.StreamUrl"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  INVALIDPARAMETER_USERSIG = "InvalidParameter.UserSig"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
func (c *Client) StartStreamIngestWithContext(ctx context.Context, request *StartStreamIngestRequest) (response *StartStreamIngestResponse, err error) {
    if request == nil {
        request = NewStartStreamIngestRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StartStreamIngest")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartStreamIngest require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartStreamIngestResponse()
    err = c.Send(request, response)
    return
}

func NewStartWebRecordRequest() (request *StartWebRecordRequest) {
    request = &StartWebRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StartWebRecord")
    
    
    return
}

func NewStartWebRecordResponse() (response *StartWebRecordResponse) {
    response = &StartWebRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StartWebRecord
// 通过此接口可以发起 WEB 页面录制任务，在接口参数中指定录制 URL，录制分辨率，录制结果存储等参数。
//
// 因为参数或API逻辑问题会立即返回结果。而因为页面问题，如页面无法访问，会在回调中返回结果，请关注。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTABILITY = "FailedOperation.NotAbility"
//  FAILEDOPERATION_NOTALLOWED = "FailedOperation.NotAllowed"
//  FAILEDOPERATION_SDKAPPIDNOTUNDERAPPID = "FailedOperation.SdkAppIdNotUnderAppId"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
func (c *Client) StartWebRecord(request *StartWebRecordRequest) (response *StartWebRecordResponse, err error) {
    return c.StartWebRecordWithContext(context.Background(), request)
}

// StartWebRecord
// 通过此接口可以发起 WEB 页面录制任务，在接口参数中指定录制 URL，录制分辨率，录制结果存储等参数。
//
// 因为参数或API逻辑问题会立即返回结果。而因为页面问题，如页面无法访问，会在回调中返回结果，请关注。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTABILITY = "FailedOperation.NotAbility"
//  FAILEDOPERATION_NOTALLOWED = "FailedOperation.NotAllowed"
//  FAILEDOPERATION_SDKAPPIDNOTUNDERAPPID = "FailedOperation.SdkAppIdNotUnderAppId"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
func (c *Client) StartWebRecordWithContext(ctx context.Context, request *StartWebRecordRequest) (response *StartWebRecordResponse, err error) {
    if request == nil {
        request = NewStartWebRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StartWebRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartWebRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartWebRecordResponse()
    err = c.Send(request, response)
    return
}

func NewStopAIConversationRequest() (request *StopAIConversationRequest) {
    request = &StopAIConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StopAIConversation")
    
    
    return
}

func NewStopAIConversationResponse() (response *StopAIConversationResponse) {
    response = &StopAIConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopAIConversation
// 停止AI对话任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) StopAIConversation(request *StopAIConversationRequest) (response *StopAIConversationResponse, err error) {
    return c.StopAIConversationWithContext(context.Background(), request)
}

// StopAIConversation
// 停止AI对话任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) StopAIConversationWithContext(ctx context.Context, request *StopAIConversationRequest) (response *StopAIConversationResponse, err error) {
    if request == nil {
        request = NewStopAIConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StopAIConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopAIConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopAIConversationResponse()
    err = c.Send(request, response)
    return
}

func NewStopAITranscriptionRequest() (request *StopAITranscriptionRequest) {
    request = &StopAITranscriptionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StopAITranscription")
    
    
    return
}

func NewStopAITranscriptionResponse() (response *StopAITranscriptionResponse) {
    response = &StopAITranscriptionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopAITranscription
// 停止AI转录任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) StopAITranscription(request *StopAITranscriptionRequest) (response *StopAITranscriptionResponse, err error) {
    return c.StopAITranscriptionWithContext(context.Background(), request)
}

// StopAITranscription
// 停止AI转录任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) StopAITranscriptionWithContext(ctx context.Context, request *StopAITranscriptionRequest) (response *StopAITranscriptionResponse, err error) {
    if request == nil {
        request = NewStopAITranscriptionRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StopAITranscription")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopAITranscription require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopAITranscriptionResponse()
    err = c.Send(request, response)
    return
}

func NewStopMCUMixTranscodeRequest() (request *StopMCUMixTranscodeRequest) {
    request = &StopMCUMixTranscodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StopMCUMixTranscode")
    
    
    return
}

func NewStopMCUMixTranscodeResponse() (response *StopMCUMixTranscodeResponse) {
    response = &StopMCUMixTranscodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopMCUMixTranscode
// 接口说明：结束云端混流
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIXSESSIONNOTEXIST = "FailedOperation.MixSessionNotExist"
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StopMCUMixTranscode(request *StopMCUMixTranscodeRequest) (response *StopMCUMixTranscodeResponse, err error) {
    return c.StopMCUMixTranscodeWithContext(context.Background(), request)
}

// StopMCUMixTranscode
// 接口说明：结束云端混流
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIXSESSIONNOTEXIST = "FailedOperation.MixSessionNotExist"
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_GETROOMFROMCACHEERROR = "InternalError.GetRoomFromCacheError"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StopMCUMixTranscodeWithContext(ctx context.Context, request *StopMCUMixTranscodeRequest) (response *StopMCUMixTranscodeResponse, err error) {
    if request == nil {
        request = NewStopMCUMixTranscodeRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StopMCUMixTranscode")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopMCUMixTranscode require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopMCUMixTranscodeResponse()
    err = c.Send(request, response)
    return
}

func NewStopMCUMixTranscodeByStrRoomIdRequest() (request *StopMCUMixTranscodeByStrRoomIdRequest) {
    request = &StopMCUMixTranscodeByStrRoomIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StopMCUMixTranscodeByStrRoomId")
    
    
    return
}

func NewStopMCUMixTranscodeByStrRoomIdResponse() (response *StopMCUMixTranscodeByStrRoomIdResponse) {
    response = &StopMCUMixTranscodeByStrRoomIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopMCUMixTranscodeByStrRoomId
// 接口说明：结束云端混流
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIXSESSIONNOTEXIST = "FailedOperation.MixSessionNotExist"
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  FAILEDOPERATION_SDKAPPIDNOTEXIST = "FailedOperation.SdkAppIdNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STRROOMID = "InvalidParameter.StrRoomId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StopMCUMixTranscodeByStrRoomId(request *StopMCUMixTranscodeByStrRoomIdRequest) (response *StopMCUMixTranscodeByStrRoomIdResponse, err error) {
    return c.StopMCUMixTranscodeByStrRoomIdWithContext(context.Background(), request)
}

// StopMCUMixTranscodeByStrRoomId
// 接口说明：结束云端混流
//
// 可能返回的错误码:
//  FAILEDOPERATION_MIXSESSIONNOTEXIST = "FailedOperation.MixSessionNotExist"
//  FAILEDOPERATION_REQUESTREJECTION = "FailedOperation.RequestRejection"
//  FAILEDOPERATION_ROOMNOTEXIST = "FailedOperation.RoomNotExist"
//  FAILEDOPERATION_SDKAPPIDNOTEXIST = "FailedOperation.SdkAppIdNotExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STRROOMID = "InvalidParameter.StrRoomId"
//  INVALIDPARAMETERVALUE_ROOMID = "InvalidParameterValue.RoomId"
//  MISSINGPARAMETER_ROOMID = "MissingParameter.RoomId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNAUTHORIZEDOPERATION_SDKAPPID = "UnauthorizedOperation.SdkAppId"
func (c *Client) StopMCUMixTranscodeByStrRoomIdWithContext(ctx context.Context, request *StopMCUMixTranscodeByStrRoomIdRequest) (response *StopMCUMixTranscodeByStrRoomIdResponse, err error) {
    if request == nil {
        request = NewStopMCUMixTranscodeByStrRoomIdRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StopMCUMixTranscodeByStrRoomId")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopMCUMixTranscodeByStrRoomId require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopMCUMixTranscodeByStrRoomIdResponse()
    err = c.Send(request, response)
    return
}

func NewStopPublishCdnStreamRequest() (request *StopPublishCdnStreamRequest) {
    request = &StopPublishCdnStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StopPublishCdnStream")
    
    
    return
}

func NewStopPublishCdnStreamResponse() (response *StopPublishCdnStreamResponse) {
    response = &StopPublishCdnStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopPublishCdnStream
// 接口说明：
//
// 停止指定的混流转推任务。如果没有调用 Stop 接口停止任务，所有参与混流转推的主播离开 TRTC 房间超过 AgentParams.MaxIdleTime 设置的时间后，任务也会自动停止。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopPublishCdnStream(request *StopPublishCdnStreamRequest) (response *StopPublishCdnStreamResponse, err error) {
    return c.StopPublishCdnStreamWithContext(context.Background(), request)
}

// StopPublishCdnStream
// 接口说明：
//
// 停止指定的混流转推任务。如果没有调用 Stop 接口停止任务，所有参与混流转推的主播离开 TRTC 房间超过 AgentParams.MaxIdleTime 设置的时间后，任务也会自动停止。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION_CRUNSUPPORTMETHOD = "FailedOperation.CRUnsupportMethod"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
func (c *Client) StopPublishCdnStreamWithContext(ctx context.Context, request *StopPublishCdnStreamRequest) (response *StopPublishCdnStreamResponse, err error) {
    if request == nil {
        request = NewStopPublishCdnStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StopPublishCdnStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopPublishCdnStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopPublishCdnStreamResponse()
    err = c.Send(request, response)
    return
}

func NewStopStreamIngestRequest() (request *StopStreamIngestRequest) {
    request = &StopStreamIngestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StopStreamIngest")
    
    
    return
}

func NewStopStreamIngestResponse() (response *StopStreamIngestResponse) {
    response = &StopStreamIngestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopStreamIngest
// 停止一个输入在线媒体流任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKFINISHED = "FailedOperation.TaskFinished"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) StopStreamIngest(request *StopStreamIngestRequest) (response *StopStreamIngestResponse, err error) {
    return c.StopStreamIngestWithContext(context.Background(), request)
}

// StopStreamIngest
// 停止一个输入在线媒体流任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKFINISHED = "FailedOperation.TaskFinished"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) StopStreamIngestWithContext(ctx context.Context, request *StopStreamIngestRequest) (response *StopStreamIngestResponse, err error) {
    if request == nil {
        request = NewStopStreamIngestRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StopStreamIngest")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopStreamIngest require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopStreamIngestResponse()
    err = c.Send(request, response)
    return
}

func NewStopWebRecordRequest() (request *StopWebRecordRequest) {
    request = &StopWebRecordRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "StopWebRecord")
    
    
    return
}

func NewStopWebRecordResponse() (response *StopWebRecordResponse) {
    response = &StopWebRecordResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// StopWebRecord
// 停止页面录制任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) StopWebRecord(request *StopWebRecordRequest) (response *StopWebRecordResponse, err error) {
    return c.StopWebRecordWithContext(context.Background(), request)
}

// StopWebRecord
// 停止页面录制任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
func (c *Client) StopWebRecordWithContext(ctx context.Context, request *StopWebRecordRequest) (response *StopWebRecordResponse, err error) {
    if request == nil {
        request = NewStopWebRecordRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "StopWebRecord")
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopWebRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopWebRecordResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAIConversationRequest() (request *UpdateAIConversationRequest) {
    request = &UpdateAIConversationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "UpdateAIConversation")
    
    
    return
}

func NewUpdateAIConversationResponse() (response *UpdateAIConversationResponse) {
    response = &UpdateAIConversationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAIConversation
// 更新AIConversation参数
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
func (c *Client) UpdateAIConversation(request *UpdateAIConversationRequest) (response *UpdateAIConversationResponse, err error) {
    return c.UpdateAIConversationWithContext(context.Background(), request)
}

// UpdateAIConversation
// 更新AIConversation参数
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
func (c *Client) UpdateAIConversationWithContext(ctx context.Context, request *UpdateAIConversationRequest) (response *UpdateAIConversationResponse, err error) {
    if request == nil {
        request = NewUpdateAIConversationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "UpdateAIConversation")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAIConversation require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAIConversationResponse()
    err = c.Send(request, response)
    return
}

func NewUpdatePublishCdnStreamRequest() (request *UpdatePublishCdnStreamRequest) {
    request = &UpdatePublishCdnStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "UpdatePublishCdnStream")
    
    
    return
}

func NewUpdatePublishCdnStreamResponse() (response *UpdatePublishCdnStreamResponse) {
    response = &UpdatePublishCdnStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdatePublishCdnStream
// 接口说明：
//
// 成功发起混流转推后，可以使用此接口来更新任务。仅在任务进行时有效，任务退出后更新将会返回错误。更新操作为增量更新模式。
//
// 注意：为了保障推流的稳定性，更新不支持任务在纯音频、音视频、纯视频之间进行切换。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTOUTDATED = "FailedOperation.RequestOutdated"
//  FAILEDOPERATION_TASKFINISHED = "FailedOperation.TaskFinished"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdatePublishCdnStream(request *UpdatePublishCdnStreamRequest) (response *UpdatePublishCdnStreamResponse, err error) {
    return c.UpdatePublishCdnStreamWithContext(context.Background(), request)
}

// UpdatePublishCdnStream
// 接口说明：
//
// 成功发起混流转推后，可以使用此接口来更新任务。仅在任务进行时有效，任务退出后更新将会返回错误。更新操作为增量更新模式。
//
// 注意：为了保障推流的稳定性，更新不支持任务在纯音频、音视频、纯视频之间进行切换。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_REQUESTOUTDATED = "FailedOperation.RequestOutdated"
//  FAILEDOPERATION_TASKFINISHED = "FailedOperation.TaskFinished"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdatePublishCdnStreamWithContext(ctx context.Context, request *UpdatePublishCdnStreamRequest) (response *UpdatePublishCdnStreamResponse, err error) {
    if request == nil {
        request = NewUpdatePublishCdnStreamRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "UpdatePublishCdnStream")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePublishCdnStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePublishCdnStreamResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateStreamIngestRequest() (request *UpdateStreamIngestRequest) {
    request = &UpdateStreamIngestRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "UpdateStreamIngest")
    
    
    return
}

func NewUpdateStreamIngestResponse() (response *UpdateStreamIngestResponse) {
    response = &UpdateStreamIngestResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateStreamIngest
// 更新输入在线媒体流任务的StreamUrl
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_STREAMURL = "InvalidParameter.StreamUrl"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
func (c *Client) UpdateStreamIngest(request *UpdateStreamIngestRequest) (response *UpdateStreamIngestResponse, err error) {
    return c.UpdateStreamIngestWithContext(context.Background(), request)
}

// UpdateStreamIngest
// 更新输入在线媒体流任务的StreamUrl
//
// 可能返回的错误码:
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  INTERNALERROR_DBERROR = "InternalError.DBError"
//  INVALIDPARAMETER_STREAMURL = "InvalidParameter.StreamUrl"
//  INVALIDPARAMETER_TASKID = "InvalidParameter.TaskId"
//  MISSINGPARAMETER_SDKAPPID = "MissingParameter.SdkAppId"
//  MISSINGPARAMETER_TASKID = "MissingParameter.TaskId"
func (c *Client) UpdateStreamIngestWithContext(ctx context.Context, request *UpdateStreamIngestRequest) (response *UpdateStreamIngestResponse, err error) {
    if request == nil {
        request = NewUpdateStreamIngestRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "UpdateStreamIngest")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateStreamIngest require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateStreamIngestResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateVoicePrintRequest() (request *UpdateVoicePrintRequest) {
    request = &UpdateVoicePrintRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("trtc", APIVersion, "UpdateVoicePrint")
    
    
    return
}

func NewUpdateVoicePrintResponse() (response *UpdateVoicePrintResponse) {
    response = &UpdateVoicePrintResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateVoicePrint
// 传入声纹ID以及对应音频信息，更新对应声纹信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSOPERATIONFAILED = "FailedOperation.CosOperationFailed"
//  FAILEDOPERATION_DBOPERATIONFAILED = "FailedOperation.DbOperationFailed"
//  FAILEDOPERATION_UPDATEVOICEPRINTIDNOTFOUND = "FailedOperation.UpdateVoicePrintIdNotFound"
//  FAILEDOPERATION_VOICEPRINTAUDIOCHECKFAILED = "FailedOperation.VoicePrintAudioCheckFailed"
//  INVALIDPARAMETER_INVALIDAUDIOINPUT = "InvalidParameter.InvalidAudioInput"
//  MISSINGPARAMETER_MISSINGVOICEPRINTUPDATEPARAMS = "MissingParameter.MissingVoicePrintUpdateParams"
func (c *Client) UpdateVoicePrint(request *UpdateVoicePrintRequest) (response *UpdateVoicePrintResponse, err error) {
    return c.UpdateVoicePrintWithContext(context.Background(), request)
}

// UpdateVoicePrint
// 传入声纹ID以及对应音频信息，更新对应声纹信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_COSOPERATIONFAILED = "FailedOperation.CosOperationFailed"
//  FAILEDOPERATION_DBOPERATIONFAILED = "FailedOperation.DbOperationFailed"
//  FAILEDOPERATION_UPDATEVOICEPRINTIDNOTFOUND = "FailedOperation.UpdateVoicePrintIdNotFound"
//  FAILEDOPERATION_VOICEPRINTAUDIOCHECKFAILED = "FailedOperation.VoicePrintAudioCheckFailed"
//  INVALIDPARAMETER_INVALIDAUDIOINPUT = "InvalidParameter.InvalidAudioInput"
//  MISSINGPARAMETER_MISSINGVOICEPRINTUPDATEPARAMS = "MissingParameter.MissingVoicePrintUpdateParams"
func (c *Client) UpdateVoicePrintWithContext(ctx context.Context, request *UpdateVoicePrintRequest) (response *UpdateVoicePrintResponse, err error) {
    if request == nil {
        request = NewUpdateVoicePrintRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "trtc", APIVersion, "UpdateVoicePrint")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateVoicePrint require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateVoicePrintResponse()
    err = c.Send(request, response)
    return
}
