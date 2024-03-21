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
// 启动云端录制功能，完成房间内的音视频录制，并上传到指定的云存储。您可以通过此 API 接口把TRTC 房间中的每一路音视频流做单独的录制又或者多路视频画面混流一路。
//
// 
//
// 您可以通过此接口实现如下目标：
//
// * 指定订阅流参数（RecordParams）来指定需要录制的主播的黑名单或者白名单。
//
// * 指定第三方存储的参数（StorageParams）来指定上传到您希望的云存储，目前支持云点播VOD和对象存储COS
//
// * 指定混流模式下的音视频转码详细参数（MixTranscodeParams），包括视频分辨率、视频码率、视频帧率、以及声音质量等
//
// * 指定混流模式各路画面的位置和布局或者也可以指定自动模板的方式来配置。
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
// 启动云端录制功能，完成房间内的音视频录制，并上传到指定的云存储。您可以通过此 API 接口把TRTC 房间中的每一路音视频流做单独的录制又或者多路视频画面混流一路。
//
// 
//
// 您可以通过此接口实现如下目标：
//
// * 指定订阅流参数（RecordParams）来指定需要录制的主播的黑名单或者白名单。
//
// * 指定第三方存储的参数（StorageParams）来指定上传到您希望的云存储，目前支持云点播VOD和对象存储COS
//
// * 指定混流模式下的音视频转码详细参数（MixTranscodeParams），包括视频分辨率、视频码率、视频帧率、以及声音质量等
//
// * 指定混流模式各路画面的位置和布局或者也可以指定自动模板的方式来配置。
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCloudRecordingResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreatePicture require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreatePictureResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCloudRecordingResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeletePicture require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeletePictureResponse()
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
// 查询指定时间内的用户列表及用户通话质量数据，可查询14天内数据。DataType 不为null，查询起止时间不超过1个小时，查询用户不超过6个，支持跨天查询。DataType为null时，查询起止时间不超过4个小时， 默认查询6个用户，同时支持每页查询100以内用户个数（PageSize不超过100）。接口用于查询质量问题，不推荐作为计费使用。（同老接口DescribeCallDetail）
//
// **注意**：
//
// 1.该接口只用于历史数据统计或核对数据使用，实时类关键业务逻辑不能使用。
//
// 2.该接口目前免费提供中，监控仪表盘商业化计费后该接口需要订阅付费版后方可调用，仪表盘商业化说明请见：https://cloud.tencent.com/document/product/647/77735
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
// 查询指定时间内的用户列表及用户通话质量数据，可查询14天内数据。DataType 不为null，查询起止时间不超过1个小时，查询用户不超过6个，支持跨天查询。DataType为null时，查询起止时间不超过4个小时， 默认查询6个用户，同时支持每页查询100以内用户个数（PageSize不超过100）。接口用于查询质量问题，不推荐作为计费使用。（同老接口DescribeCallDetail）
//
// **注意**：
//
// 1.该接口只用于历史数据统计或核对数据使用，实时类关键业务逻辑不能使用。
//
// 2.该接口目前免费提供中，监控仪表盘商业化计费后该接口需要订阅付费版后方可调用，仪表盘商业化说明请见：https://cloud.tencent.com/document/product/647/77735
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCallDetailInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCallDetailInfoResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeCloudRecordingResponse()
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
// 查询SdkAppId下的房间列表。默认返回10条通话，一次最多返回100条通话。可查询14天内的数据。（同老接口DescribeRoomInformation）
//
// **注意**：
//
// 1.该接口只用于历史数据统计或核对数据使用，实时类关键业务逻辑不能使用。
//
// 2.该接口目前免费提供中，监控仪表盘商业化计费后该接口需要订阅付费版后方可调用，仪表盘商业化说明请见：https://cloud.tencent.com/document/product/647/77735
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
// 查询SdkAppId下的房间列表。默认返回10条通话，一次最多返回100条通话。可查询14天内的数据。（同老接口DescribeRoomInformation）
//
// **注意**：
//
// 1.该接口只用于历史数据统计或核对数据使用，实时类关键业务逻辑不能使用。
//
// 2.该接口目前免费提供中，监控仪表盘商业化计费后该接口需要订阅付费版后方可调用，仪表盘商业化说明请见：https://cloud.tencent.com/document/product/647/77735
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
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，监控仪表盘[版本功能和计费说明]（https://cloud.tencent.com/document/product/647/81331）。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时。
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，监控仪表盘[版本功能和计费说明]（https://cloud.tencent.com/document/product/647/81331）。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时。
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTRTCRealTimeQualityDataWithContext(ctx context.Context, request *DescribeTRTCRealTimeQualityDataRequest) (response *DescribeTRTCRealTimeQualityDataResponse, err error) {
    if request == nil {
        request = NewDescribeTRTCRealTimeQualityDataRequest()
    }
    
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
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，监控仪表盘[版本功能和计费说明](https://cloud.tencent.com/document/product/647/81331)。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时。
//
// 3.除此之外您也可以通过[订阅TRTC包月套餐](https://buy.cloud.tencent.com/trtc)尊享版或旗舰版解锁此接口的调用能力，请在开通包月套餐后，请[提交工单](https://console.cloud.tencent.com/workorder/category)联系售后解锁调用能力
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
// 1.调用接口需开通监控仪表盘【基础版】和【进阶版】，监控仪表盘【免费版】不支持调用，监控仪表盘[版本功能和计费说明](https://cloud.tencent.com/document/product/647/81331)。
//
// 2.查询时间范围根据监控仪表盘功能版本而定，基础版可查近3小时，进阶版可查近12小时。
//
// 3.除此之外您也可以通过[订阅TRTC包月套餐](https://buy.cloud.tencent.com/trtc)尊享版或旗舰版解锁此接口的调用能力，请在开通包月套餐后，请[提交工单](https://console.cloud.tencent.com/workorder/category)联系售后解锁调用能力
//
// 可能返回的错误码:
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeTRTCRealTimeScaleDataWithContext(ctx context.Context, request *DescribeTRTCRealTimeScaleDataRequest) (response *DescribeTRTCRealTimeScaleDataResponse, err error) {
    if request == nil {
        request = NewDescribeTRTCRealTimeScaleDataRequest()
    }
    
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
func (c *Client) DescribeTrtcRoomUsageWithContext(ctx context.Context, request *DescribeTrtcRoomUsageRequest) (response *DescribeTrtcRoomUsageResponse, err error) {
    if request == nil {
        request = NewDescribeTrtcRoomUsageRequest()
    }
    
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
// 获取TRTC音视频互动的用量明细。
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
// 获取TRTC音视频互动的用量明细。
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
// 查询指定时间内的用户列表，可查询14天内数据，查询起止时间不超过4小时。默认每页查询6个用户，支持每页最大查询100个用户PageSize不超过100）。（同老接口DescribeUserInformation）
//
// **注意**：
//
// 1.该接口只用于历史数据统计或核对数据使用，实时类关键业务逻辑不能使用。
//
// 2.该接口目前免费提供中，监控仪表盘商业化计费后该接口需要订阅付费版后方可调用，仪表盘商业化说明请见：https://cloud.tencent.com/document/product/647/77735
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
// 查询指定时间内的用户列表，可查询14天内数据，查询起止时间不超过4小时。默认每页查询6个用户，支持每页最大查询100个用户PageSize不超过100）。（同老接口DescribeUserInformation）
//
// **注意**：
//
// 1.该接口只用于历史数据统计或核对数据使用，实时类关键业务逻辑不能使用。
//
// 2.该接口目前免费提供中，监控仪表盘商业化计费后该接口需要订阅付费版后方可调用，仪表盘商业化说明请见：https://cloud.tencent.com/document/product/647/77735
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeUserInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeUserInfoResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DismissRoomByStrRoomId require credential")
    }

    request.SetContext(ctx)
    
    response = NewDismissRoomByStrRoomIdResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCloudRecording require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCloudRecordingResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyPicture require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyPictureResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("RemoveUserByStrRoomId require credential")
    }

    request.SetContext(ctx)
    
    response = NewRemoveUserByStrRoomIdResponse()
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
// TRTC房间中可能会同时存在多路音视频流，您可以通过混流转推API，通知腾讯云服务端将同个房间或者多个房间的多路视频画面混合到一起，并指定每一路画面的位置，同时将多路声音进行混音，最终形成一路音视频流，以便用于录制和直播观看，同时可以支持将这一路音视频的混流回推到TRTC房间内部。
//
// 
//
// 云api混流转推功能包含三个接口：  
//
// 1、StartPublishCdnStream：启动一个混流转推任务，此接口会发起一个新的混流转推任务，启动成功后会返回一个sdkappid维度唯一的任务id（TaskId）。您需要保存该TaskId，后续需要依赖此TaskId更新任务和结束任务。  
//
// 2、UpdatePublishCdnStream：更新指定的混流转推任务，包括：更新视频画面的布局、更新混音列表、更新转推cdn的地址列表以及更新回推房间列表。  
//
// 3、StopPublishCdnStream：停止指定的混流转推任务。
//
// 
//
// 您可以通过此套接口实现目标：  
//
// 1、设置最终混流的画质和音质，包括视频分辨率、视频帧率、视频码率，和音频质量。  
//
// 2、设置每一路画面的布局位置，您只需要设置一次，后续您指定的混流用户进房上行音视频时，排版引擎会自动将该用户的画面混合到您指定的布局位置。同时您可以使用更新接口调整画面的布局位置。  
//
// 3、设置多个cdn转推的目的地址，一个转推任务最多可以同时设置10个cdn转推地址，如果您需要转推非腾讯云的cdn地址时，需要联系腾讯云开通能力。  
//
// 4、设置多个回推房间列表，一个转推任务最多可以同时将混流回推到10个TRTC房间。
//
// 
//
// 目前已经支持如下几种布局模版，其中动态布局模版（悬浮模板、九宫格模板、屏幕分享模板）只支持单个TRTC房间，自定义模版支持混合多个TRTC房间内的音视频流。具体说明如下：  
//
// 1、悬浮模板：第一个进入房间的用户的视频画面会铺满整个屏幕，其他用户的视频画面从左下角依次水平排列，显示为小画面，最多4行，每行4个，小画面悬浮于大画面之上。最多支持1个大画面和15个小画面，如果用户只发送音频默认不占布局配置，也支持设置。每个子画面默认采用居中裁剪的方式进行渲染，也支持统一设置子画面的渲染方式。   
//
// 2、九宫格模板：所有用户的视频画面大小一致，平分整个屏幕，人数越多，每个画面的尺寸越小。最多支持16个画面，如果用户只发送音频，默认不占布局配置，也支持设置。每个子画面默认采用居中裁剪的方式进行渲染，也支持统一设置子画面的渲染方式。  
//
// 3、屏幕分享模板：适合视频会议和在线教育场景的布局，屏幕分享（或者主讲的摄像头）始终占据屏幕左侧的大画面位置，需要您明确设置占据大画面的混流用户信息。其他用户依次垂直排列于右侧，最多两列，每列最多8个小画面。最多支持1个大画面和15个小画面。若上行分辨率宽高比与画面输出宽高比不一致时，左侧大画面为了保持内容的完整性采用缩放方式处理，右侧小画面采用裁剪方式处理，也支持统一设置子画面的渲染方式。  
//
// 4、自定义布局模版：支持您主动根据业务需要设置布局位置，每个预设的布局位置支持具名设置（具名设置需要明确指定房间号和用户名）和不具名设置。当一个子画面具名设置时，该位置就为该用户预留，用户进房且上行音视频数据时会自动占据该位置，其它用户不会占据该位置。当预设的布局位置未具名时，排版引擎会根据进房间顺序自动填充，预设位置填满时，不再混合其它用户的画面和声音。每个子画面位置支持设置占位图（BackgroundImageUrl），当用户未进房或者只上行音频数据时，该位置画面可显示对应的占位图画面。
//
// 
//
// 您可以控制台开通旁路转推回调功能实现转推cdn状态的事件监控，具体说明请参考官网文档：[旁路转推回调说明](https://cloud.tencent.com/document/product/647/88552)  
//
// 您使用转推api时根据使用特性可能会产生如下费用：  
//
// MCU混流转码费用请参考文档：[云端混流转码计费说明](https://cloud.tencent.com/document/product/647/49446)  
//
// 转推非腾讯云CDN费用请参考文档：[云端转推计费说明](https://cloud.tencent.com/document/product/647/82155)
//
// 
//
// 参数的使用说明：  
//
// 1、AgentParams：每个转推任务会拉起一个机器人用户进入TRTC房间进行拉流，您需要通过AgentParams.UserId参数进行设置，这个机器人id不能和房间中的普通用户id冲突，否则会导致转推任务由于机器人用户被踢出TRTC房间而异常结束，您可以通过增加特殊前缀的方式规避。您可以通过设置AgentParams.MaxIdleTime控制转推任务的自动结束，当设置此参数时，所有参与混流转推的主播持续离开TRTC房间超过MaxIdleTime的时长，自动停止转推任务。注意：参与混流转推的主播仅是停止音视频上行，转推任务不会自动停止。  
//
// 2、WithTranscoding：如果需要将多路音视频流混合到一路时，WithTranscoding必须设置为1。  
//
// 3、AudioParams：转推任务音频参数和视频参数是分开设置的，若您想要将指定的用户音频进行混音时，需要明确设置AudioParams.SubscribeAudioList。若您不设置AudioParams.SubscribeAudioList，混音引擎会自动将TRTC房间中所有用户的音频混合。若您想要混合TRTC房间除指定用户之外的所有用户的声音，可以通过AudioParams.UnSubscribeAudioList设置音频黑名单列表。
//
// 4、VideoParams：若您想要将用户的视频混合，可以通过VideoParams设置，若只想要混纯音频，则不用设置VideoParams参数。您可以通过VideoParams.LayoutParams.MixLayoutMode设置画面的布局模式，包括：动态布局（1：悬浮布局（默认），2：屏幕分享布局，3：九宫格布局）和自定义布局。动态布局模式由排版引擎按照固定的布局自动混合，不需要设置VideoParams.LayoutParams.MixLayoutList。当使用悬浮布局和屏幕分享布局时，您可以通过设置VideoParams.LayoutParams.MaxVideoUser参数指定大画面用户。自定义布局模式提供给您自主布局画面的能力，可以通过VideoParams.LayoutParams.MixLayoutList参数指定每个用户所在的布局位置。在每个布局参数中，您可以通过用户媒体流参数（UserMediaStream）指定这个布局位置为指定的用户预留，也可以不设置UserMediaStream，由排版引擎按照用户进入TRTC房间的顺序自动填充，另外您也可以设置每个布局位置的渲染方式（RenderMode）和裁剪方式（CustomCrop）。  
//
// 5、VideoParams.WaterMarkList：若您想要混流画面中叠加水印，可以通过VideoParams.WaterMarkList参数设置，支持图片水印和文字水印，支持透明通道。  
//
// 6、SingleSubscribeParams：若您想要将TRTC房间中的一路单流推到CDN，可以使用SingleSubscribeParams参数设置，此时需要将WithTranscoding参数设置为0。  
//
// 7、PublishCdnParams.N：若您想要推流到CDN，可以使用PublishCdnParams.N参数设置，支持最多同时推流到10个CDN地址。若转推地址是腾讯云CDN时，请将IsTencentCdn明确设置为1；若您有转推非腾讯云CDN的需求时，请联系腾讯云技术支持开通，转推非腾讯云会产生转推费用，费用说明请参考官网文档：[云端转推计费说明](https://cloud.tencent.com/document/product/647/82155)。    
//
// 8、FeedBackRoomParams.N：若您想要将混合的音视频流回推到TRTC房间，可以使用FeedBackRoomParams.N参数设置，支持最多同时推10路流回TRTC房间。您需要指定回推的TRTC房间号和机器人ID（UserId），机器人ID不能与普通用户ID冲突，否则会导致转推任务由于机器人用户被踢出TRTC房间而异常结束，您可以通过增加特殊前缀的方式规避。  
//
// 9、SeiParams：若您想要混合的音视频流中增加SEI信息时，可以使用SeiParams参数设置，支持音量布局SEI和叠加转推请求中的SEI，其中音量布局SEI的内容是固定的json结构，具体请看本章节后续的SEI说明。您可以通过FollowIdr参数设置SEI跟随关键帧一起发送。音量布局sei的说明如下：  
//
// 如果您的CDN观众端需要识别到参与混流的主播位置，以及需要识别到参与混流的主播的音量信息，可以通过音量布局sei实现。音量布局sei的payload内容及参数说明如下：
//
// 
//
// ```
//
// {
//
//     "app_data":"",
//
//     "canvas":{
//
//         "w":1080,
//
//         "h":960
//
//     },
//
//     "regions":[
//
//         {
//
//             "uid":"65949987242835883c",
//
//             "zorder":2,
//
//             "volume":45,
//
//             "x":270,
//
//             "y":480,
//
//             "w":540,
//
//             "h":480
//
//         },
//
//         {
//
//             "uid":"659c9d8d242b328d31",
//
//             "zorder":2,
//
//             "volume":0,
//
//             "x":0,
//
//             "y":0,
//
//             "w":540,
//
//             "h":480
//
//         },
//
//         {
//
//             "uid":"64989a82272b308c",
//
//             "zorder":2,
//
//             "volume":91,
//
//             "x":540,
//
//             "y":0,
//
//             "w":540,
//
//             "h":480
//
//         }
//
//     ],
//
//     "ver":"1.0",
//
//     "ts":1648544726
//
// }
//
// ```
//
// canvas：这个是混流信令中VideoEncode设置的宽高，即：混流输出的整个画布宽高。  
//
// regions：里面是真实混入的用户id和对应的子画面位置，若参与混流的用户未进入TRTC房间或者未开启视频上行，regions中不会包含该用户。  
//
// uid：代表参与混流的用户ID。  
//
// zorder：参与混流userid在混流输出的层级。  
//
// x/y：参与混流userid的子画面在画布的坐标。  
//
// w/h：参与混流userid的子画面的大小。  
//
// volume：代表混流用户的音量，取值范围为0-100，值越大，代表该用户参与混流时的音量越大。  
//
// ts：是输出sei的服务器本地秒级时戳。   
//
// ver：可以忽略。
//
// 
//
// 其它使用注意事项：  
//
// 1、使用混流转推接口时，您需要先调用启动转推任务接口（StartPublishCdnStream），获取启动转推任务响应中的任务ID标识（TaskId）。后续传入任务ID标识（TaskId）来更新转推任务（UpdatePublishCdnStream）和停止转推任务（StopPublishCdnStream）。  
//
// 2、转推API不支持发起到TRTC控制台配置的自动旁路任务，以及通过TRTC SDK进房接口中设置的自定义流ID的旁路任务。  
//
// 3、为了确保转推链接的稳定，同一个转推任务不支持纯音频、音视频、纯视频之间的切换。  
//
// 4、为了确保转推链接的稳定，不支持update时更新视频参数（codec）和音频参数（codec、采样率、码率、声道数）。  
//
// 5、发起单流旁路任务时，AudioParams和VideoParams都填写表示音视频旁路，如果仅填写AudioParams表示纯音频旁路，任务进行过程中不支持纯音频到音视频的切换。音视频旁路时，VideoParams中的Width、Height、Fps、BitRate、Gop需要按照真实上行参数填写。  
//
// 6、更新请求中必须携带SequenceNumber参数，用于防止请求乱序。客户保证对同一个任务更新时的SequenceNumber参数递增，否则会导致混流任务更新失败。  
//
// 7、调用api时region选择说明：如果应用id是1400xxx时，region可填北京、上海、广州、香港，如果您的cdn观众主要在海外，请选择香港；如果应用id是200xxx或400xxx时，region请选择新加坡。  
//
// 8、回推到TRTC房间的流不会参与其他回推房间任务的混流，满足下面条件之一时，可以参与其他转推cdn任务的混流: (1) 转推cdn任务在视频参数中通过具名方式指定推流机器人参与混流；(2) 转推cdn任务在音频参数中通过白名单方式指定推流机器人参与混流；(3) 转推cdn任务参与混流用户的房间号与回推机器人对应混流用户的房间号完全不同。  
//
// 9、您可以在主播进房前，提前创建转推任务，结束转推任务时需要主动调用停止接口。如果您没有调用停止转推任务接口时，腾讯云后台会按照所有参与混流的用户没有任何数据上行的时间算起，直到超过启动转推任务时设置的超时时间（AgentParams.MaxIdleTime）为止，自动停止混流转推任务。
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
// TRTC房间中可能会同时存在多路音视频流，您可以通过混流转推API，通知腾讯云服务端将同个房间或者多个房间的多路视频画面混合到一起，并指定每一路画面的位置，同时将多路声音进行混音，最终形成一路音视频流，以便用于录制和直播观看，同时可以支持将这一路音视频的混流回推到TRTC房间内部。
//
// 
//
// 云api混流转推功能包含三个接口：  
//
// 1、StartPublishCdnStream：启动一个混流转推任务，此接口会发起一个新的混流转推任务，启动成功后会返回一个sdkappid维度唯一的任务id（TaskId）。您需要保存该TaskId，后续需要依赖此TaskId更新任务和结束任务。  
//
// 2、UpdatePublishCdnStream：更新指定的混流转推任务，包括：更新视频画面的布局、更新混音列表、更新转推cdn的地址列表以及更新回推房间列表。  
//
// 3、StopPublishCdnStream：停止指定的混流转推任务。
//
// 
//
// 您可以通过此套接口实现目标：  
//
// 1、设置最终混流的画质和音质，包括视频分辨率、视频帧率、视频码率，和音频质量。  
//
// 2、设置每一路画面的布局位置，您只需要设置一次，后续您指定的混流用户进房上行音视频时，排版引擎会自动将该用户的画面混合到您指定的布局位置。同时您可以使用更新接口调整画面的布局位置。  
//
// 3、设置多个cdn转推的目的地址，一个转推任务最多可以同时设置10个cdn转推地址，如果您需要转推非腾讯云的cdn地址时，需要联系腾讯云开通能力。  
//
// 4、设置多个回推房间列表，一个转推任务最多可以同时将混流回推到10个TRTC房间。
//
// 
//
// 目前已经支持如下几种布局模版，其中动态布局模版（悬浮模板、九宫格模板、屏幕分享模板）只支持单个TRTC房间，自定义模版支持混合多个TRTC房间内的音视频流。具体说明如下：  
//
// 1、悬浮模板：第一个进入房间的用户的视频画面会铺满整个屏幕，其他用户的视频画面从左下角依次水平排列，显示为小画面，最多4行，每行4个，小画面悬浮于大画面之上。最多支持1个大画面和15个小画面，如果用户只发送音频默认不占布局配置，也支持设置。每个子画面默认采用居中裁剪的方式进行渲染，也支持统一设置子画面的渲染方式。   
//
// 2、九宫格模板：所有用户的视频画面大小一致，平分整个屏幕，人数越多，每个画面的尺寸越小。最多支持16个画面，如果用户只发送音频，默认不占布局配置，也支持设置。每个子画面默认采用居中裁剪的方式进行渲染，也支持统一设置子画面的渲染方式。  
//
// 3、屏幕分享模板：适合视频会议和在线教育场景的布局，屏幕分享（或者主讲的摄像头）始终占据屏幕左侧的大画面位置，需要您明确设置占据大画面的混流用户信息。其他用户依次垂直排列于右侧，最多两列，每列最多8个小画面。最多支持1个大画面和15个小画面。若上行分辨率宽高比与画面输出宽高比不一致时，左侧大画面为了保持内容的完整性采用缩放方式处理，右侧小画面采用裁剪方式处理，也支持统一设置子画面的渲染方式。  
//
// 4、自定义布局模版：支持您主动根据业务需要设置布局位置，每个预设的布局位置支持具名设置（具名设置需要明确指定房间号和用户名）和不具名设置。当一个子画面具名设置时，该位置就为该用户预留，用户进房且上行音视频数据时会自动占据该位置，其它用户不会占据该位置。当预设的布局位置未具名时，排版引擎会根据进房间顺序自动填充，预设位置填满时，不再混合其它用户的画面和声音。每个子画面位置支持设置占位图（BackgroundImageUrl），当用户未进房或者只上行音频数据时，该位置画面可显示对应的占位图画面。
//
// 
//
// 您可以控制台开通旁路转推回调功能实现转推cdn状态的事件监控，具体说明请参考官网文档：[旁路转推回调说明](https://cloud.tencent.com/document/product/647/88552)  
//
// 您使用转推api时根据使用特性可能会产生如下费用：  
//
// MCU混流转码费用请参考文档：[云端混流转码计费说明](https://cloud.tencent.com/document/product/647/49446)  
//
// 转推非腾讯云CDN费用请参考文档：[云端转推计费说明](https://cloud.tencent.com/document/product/647/82155)
//
// 
//
// 参数的使用说明：  
//
// 1、AgentParams：每个转推任务会拉起一个机器人用户进入TRTC房间进行拉流，您需要通过AgentParams.UserId参数进行设置，这个机器人id不能和房间中的普通用户id冲突，否则会导致转推任务由于机器人用户被踢出TRTC房间而异常结束，您可以通过增加特殊前缀的方式规避。您可以通过设置AgentParams.MaxIdleTime控制转推任务的自动结束，当设置此参数时，所有参与混流转推的主播持续离开TRTC房间超过MaxIdleTime的时长，自动停止转推任务。注意：参与混流转推的主播仅是停止音视频上行，转推任务不会自动停止。  
//
// 2、WithTranscoding：如果需要将多路音视频流混合到一路时，WithTranscoding必须设置为1。  
//
// 3、AudioParams：转推任务音频参数和视频参数是分开设置的，若您想要将指定的用户音频进行混音时，需要明确设置AudioParams.SubscribeAudioList。若您不设置AudioParams.SubscribeAudioList，混音引擎会自动将TRTC房间中所有用户的音频混合。若您想要混合TRTC房间除指定用户之外的所有用户的声音，可以通过AudioParams.UnSubscribeAudioList设置音频黑名单列表。
//
// 4、VideoParams：若您想要将用户的视频混合，可以通过VideoParams设置，若只想要混纯音频，则不用设置VideoParams参数。您可以通过VideoParams.LayoutParams.MixLayoutMode设置画面的布局模式，包括：动态布局（1：悬浮布局（默认），2：屏幕分享布局，3：九宫格布局）和自定义布局。动态布局模式由排版引擎按照固定的布局自动混合，不需要设置VideoParams.LayoutParams.MixLayoutList。当使用悬浮布局和屏幕分享布局时，您可以通过设置VideoParams.LayoutParams.MaxVideoUser参数指定大画面用户。自定义布局模式提供给您自主布局画面的能力，可以通过VideoParams.LayoutParams.MixLayoutList参数指定每个用户所在的布局位置。在每个布局参数中，您可以通过用户媒体流参数（UserMediaStream）指定这个布局位置为指定的用户预留，也可以不设置UserMediaStream，由排版引擎按照用户进入TRTC房间的顺序自动填充，另外您也可以设置每个布局位置的渲染方式（RenderMode）和裁剪方式（CustomCrop）。  
//
// 5、VideoParams.WaterMarkList：若您想要混流画面中叠加水印，可以通过VideoParams.WaterMarkList参数设置，支持图片水印和文字水印，支持透明通道。  
//
// 6、SingleSubscribeParams：若您想要将TRTC房间中的一路单流推到CDN，可以使用SingleSubscribeParams参数设置，此时需要将WithTranscoding参数设置为0。  
//
// 7、PublishCdnParams.N：若您想要推流到CDN，可以使用PublishCdnParams.N参数设置，支持最多同时推流到10个CDN地址。若转推地址是腾讯云CDN时，请将IsTencentCdn明确设置为1；若您有转推非腾讯云CDN的需求时，请联系腾讯云技术支持开通，转推非腾讯云会产生转推费用，费用说明请参考官网文档：[云端转推计费说明](https://cloud.tencent.com/document/product/647/82155)。    
//
// 8、FeedBackRoomParams.N：若您想要将混合的音视频流回推到TRTC房间，可以使用FeedBackRoomParams.N参数设置，支持最多同时推10路流回TRTC房间。您需要指定回推的TRTC房间号和机器人ID（UserId），机器人ID不能与普通用户ID冲突，否则会导致转推任务由于机器人用户被踢出TRTC房间而异常结束，您可以通过增加特殊前缀的方式规避。  
//
// 9、SeiParams：若您想要混合的音视频流中增加SEI信息时，可以使用SeiParams参数设置，支持音量布局SEI和叠加转推请求中的SEI，其中音量布局SEI的内容是固定的json结构，具体请看本章节后续的SEI说明。您可以通过FollowIdr参数设置SEI跟随关键帧一起发送。音量布局sei的说明如下：  
//
// 如果您的CDN观众端需要识别到参与混流的主播位置，以及需要识别到参与混流的主播的音量信息，可以通过音量布局sei实现。音量布局sei的payload内容及参数说明如下：
//
// 
//
// ```
//
// {
//
//     "app_data":"",
//
//     "canvas":{
//
//         "w":1080,
//
//         "h":960
//
//     },
//
//     "regions":[
//
//         {
//
//             "uid":"65949987242835883c",
//
//             "zorder":2,
//
//             "volume":45,
//
//             "x":270,
//
//             "y":480,
//
//             "w":540,
//
//             "h":480
//
//         },
//
//         {
//
//             "uid":"659c9d8d242b328d31",
//
//             "zorder":2,
//
//             "volume":0,
//
//             "x":0,
//
//             "y":0,
//
//             "w":540,
//
//             "h":480
//
//         },
//
//         {
//
//             "uid":"64989a82272b308c",
//
//             "zorder":2,
//
//             "volume":91,
//
//             "x":540,
//
//             "y":0,
//
//             "w":540,
//
//             "h":480
//
//         }
//
//     ],
//
//     "ver":"1.0",
//
//     "ts":1648544726
//
// }
//
// ```
//
// canvas：这个是混流信令中VideoEncode设置的宽高，即：混流输出的整个画布宽高。  
//
// regions：里面是真实混入的用户id和对应的子画面位置，若参与混流的用户未进入TRTC房间或者未开启视频上行，regions中不会包含该用户。  
//
// uid：代表参与混流的用户ID。  
//
// zorder：参与混流userid在混流输出的层级。  
//
// x/y：参与混流userid的子画面在画布的坐标。  
//
// w/h：参与混流userid的子画面的大小。  
//
// volume：代表混流用户的音量，取值范围为0-100，值越大，代表该用户参与混流时的音量越大。  
//
// ts：是输出sei的服务器本地秒级时戳。   
//
// ver：可以忽略。
//
// 
//
// 其它使用注意事项：  
//
// 1、使用混流转推接口时，您需要先调用启动转推任务接口（StartPublishCdnStream），获取启动转推任务响应中的任务ID标识（TaskId）。后续传入任务ID标识（TaskId）来更新转推任务（UpdatePublishCdnStream）和停止转推任务（StopPublishCdnStream）。  
//
// 2、转推API不支持发起到TRTC控制台配置的自动旁路任务，以及通过TRTC SDK进房接口中设置的自定义流ID的旁路任务。  
//
// 3、为了确保转推链接的稳定，同一个转推任务不支持纯音频、音视频、纯视频之间的切换。  
//
// 4、为了确保转推链接的稳定，不支持update时更新视频参数（codec）和音频参数（codec、采样率、码率、声道数）。  
//
// 5、发起单流旁路任务时，AudioParams和VideoParams都填写表示音视频旁路，如果仅填写AudioParams表示纯音频旁路，任务进行过程中不支持纯音频到音视频的切换。音视频旁路时，VideoParams中的Width、Height、Fps、BitRate、Gop需要按照真实上行参数填写。  
//
// 6、更新请求中必须携带SequenceNumber参数，用于防止请求乱序。客户保证对同一个任务更新时的SequenceNumber参数递增，否则会导致混流任务更新失败。  
//
// 7、调用api时region选择说明：如果应用id是1400xxx时，region可填北京、上海、广州、香港，如果您的cdn观众主要在海外，请选择香港；如果应用id是200xxx或400xxx时，region请选择新加坡。  
//
// 8、回推到TRTC房间的流不会参与其他回推房间任务的混流，满足下面条件之一时，可以参与其他转推cdn任务的混流: (1) 转推cdn任务在视频参数中通过具名方式指定推流机器人参与混流；(2) 转推cdn任务在音频参数中通过白名单方式指定推流机器人参与混流；(3) 转推cdn任务参与混流用户的房间号与回推机器人对应混流用户的房间号完全不同。  
//
// 9、您可以在主播进房前，提前创建转推任务，结束转推任务时需要主动调用停止接口。如果您没有调用停止转推任务接口时，腾讯云后台会按照所有参与混流的用户没有任何数据上行的时间算起，直到超过启动转推任务时设置的超时时间（AgentParams.MaxIdleTime）为止，自动停止混流转推任务。
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
// 将一个在线媒体流推到TRTC房间。
//
// 使用此接口的前提条件：需要购买 [尊享版或旗舰版套餐包](https://cloud.tencent.com/document/product/647/85386
//
// )
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTRTMPFUNCTION = "FailedOperation.NotRtmpFunction"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  INTERNALERROR_HTTPPARSEFAILED = "InternalError.HttpParseFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STRROOMID = "InvalidParameter.StrRoomId"
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
// 将一个在线媒体流推到TRTC房间。
//
// 使用此接口的前提条件：需要购买 [尊享版或旗舰版套餐包](https://cloud.tencent.com/document/product/647/85386
//
// )
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTRTMPFUNCTION = "FailedOperation.NotRtmpFunction"
//  FAILEDOPERATION_RESTRICTEDCONCURRENCY = "FailedOperation.RestrictedConcurrency"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  INTERNALERROR_HTTPPARSEFAILED = "InternalError.HttpParseFailed"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER_BODYPARAMSERROR = "InvalidParameter.BodyParamsError"
//  INVALIDPARAMETER_ROOMID = "InvalidParameter.RoomId"
//  INVALIDPARAMETER_SDKAPPID = "InvalidParameter.SdkAppId"
//  INVALIDPARAMETER_STRROOMID = "InvalidParameter.StrRoomId"
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
//  FAILEDOPERATION_SDKAPPIDNOTUNDERAPPID = "FailedOperation.SdkAppIdNotUnderAppId"
//  FAILEDOPERATION_SDKAPPIDNOTWEBRECORDABILITY = "FailedOperation.SdkAppIdNotWebRecordAbility"
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
//  FAILEDOPERATION_SDKAPPIDNOTUNDERAPPID = "FailedOperation.SdkAppIdNotUnderAppId"
//  FAILEDOPERATION_SDKAPPIDNOTWEBRECORDABILITY = "FailedOperation.SdkAppIdNotWebRecordAbility"
//  FAILEDOPERATION_TASKEXIST = "FailedOperation.TaskExist"
//  RESOURCEINSUFFICIENT_REQUESTREJECTION = "ResourceInsufficient.RequestRejection"
func (c *Client) StartWebRecordWithContext(ctx context.Context, request *StartWebRecordRequest) (response *StartWebRecordResponse, err error) {
    if request == nil {
        request = NewStartWebRecordRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StartWebRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStartWebRecordResponse()
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
// 停止转推任务。
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
// 停止转推任务。
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("StopWebRecord require credential")
    }

    request.SetContext(ctx)
    
    response = NewStopWebRecordResponse()
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
// 更新转推任务。
//
// 注：请参见启动转推任务的接口说明和使用说明。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdatePublishCdnStream(request *UpdatePublishCdnStreamRequest) (response *UpdatePublishCdnStreamResponse, err error) {
    return c.UpdatePublishCdnStreamWithContext(context.Background(), request)
}

// UpdatePublishCdnStream
// 更新转推任务。
//
// 注：请参见启动转推任务的接口说明和使用说明。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNREALNAMEAUTHENTICATED = "AuthFailure.UnRealNameAuthenticated"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  AUTHFAILURE_UNSUPPORTEDOPERATION = "AuthFailure.UnsupportedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdatePublishCdnStreamWithContext(ctx context.Context, request *UpdatePublishCdnStreamRequest) (response *UpdatePublishCdnStreamResponse, err error) {
    if request == nil {
        request = NewUpdatePublishCdnStreamRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdatePublishCdnStream require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdatePublishCdnStreamResponse()
    err = c.Send(request, response)
    return
}
