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

package v20200324

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-03-24"

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


func NewCreateSegmentationTaskRequest() (request *CreateSegmentationTaskRequest) {
    request = &CreateSegmentationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bda", APIVersion, "CreateSegmentationTask")
    
    
    return
}

func NewCreateSegmentationTaskResponse() (response *CreateSegmentationTaskResponse) {
    response = &CreateSegmentationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateSegmentationTask
// 本接口为人像分割在线处理接口组中的提交任务接口，可以对提交的资源进行处理视频流/图片流识别视频作品中的人像区域，进行一键抠像、背景替换、人像虚化等后期处理。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBCONFLICT = "FailedOperation.JobConflict"
//  FAILEDOPERATION_JOBQUEUEFULL = "FailedOperation.JobQueueFull"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TASKLIMITEXCEEDED = "FailedOperation.TaskLimitExceeded"
//  FAILEDOPERATION_TOOLARGEFILEERROR = "FailedOperation.TooLargeFileError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) CreateSegmentationTask(request *CreateSegmentationTaskRequest) (response *CreateSegmentationTaskResponse, err error) {
    return c.CreateSegmentationTaskWithContext(context.Background(), request)
}

// CreateSegmentationTask
// 本接口为人像分割在线处理接口组中的提交任务接口，可以对提交的资源进行处理视频流/图片流识别视频作品中的人像区域，进行一键抠像、背景替换、人像虚化等后期处理。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBCONFLICT = "FailedOperation.JobConflict"
//  FAILEDOPERATION_JOBQUEUEFULL = "FailedOperation.JobQueueFull"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TASKLIMITEXCEEDED = "FailedOperation.TaskLimitExceeded"
//  FAILEDOPERATION_TOOLARGEFILEERROR = "FailedOperation.TooLargeFileError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) CreateSegmentationTaskWithContext(ctx context.Context, request *CreateSegmentationTaskRequest) (response *CreateSegmentationTaskResponse, err error) {
    if request == nil {
        request = NewCreateSegmentationTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bda", APIVersion, "CreateSegmentationTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateSegmentationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateSegmentationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSegmentationTaskRequest() (request *DescribeSegmentationTaskRequest) {
    request = &DescribeSegmentationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bda", APIVersion, "DescribeSegmentationTask")
    
    
    return
}

func NewDescribeSegmentationTaskResponse() (response *DescribeSegmentationTaskResponse) {
    response = &DescribeSegmentationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeSegmentationTask
// 可以查看单条任务的处理情况，包括处理状态，处理结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDIODECODEFAILED = "FailedOperation.AudioDecodeFailed"
//  FAILEDOPERATION_AUDIOENCODEFAILED = "FailedOperation.AudioEncodeFailed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  FAILEDOPERATION_VIDEODECODEFAILED = "FailedOperation.VideoDecodeFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribeSegmentationTask(request *DescribeSegmentationTaskRequest) (response *DescribeSegmentationTaskResponse, err error) {
    return c.DescribeSegmentationTaskWithContext(context.Background(), request)
}

// DescribeSegmentationTask
// 可以查看单条任务的处理情况，包括处理状态，处理结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDIODECODEFAILED = "FailedOperation.AudioDecodeFailed"
//  FAILEDOPERATION_AUDIOENCODEFAILED = "FailedOperation.AudioEncodeFailed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  FAILEDOPERATION_VIDEODECODEFAILED = "FailedOperation.VideoDecodeFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) DescribeSegmentationTaskWithContext(ctx context.Context, request *DescribeSegmentationTaskRequest) (response *DescribeSegmentationTaskResponse, err error) {
    if request == nil {
        request = NewDescribeSegmentationTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bda", APIVersion, "DescribeSegmentationTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeSegmentationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeSegmentationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewSegmentCustomizedPortraitPicRequest() (request *SegmentCustomizedPortraitPicRequest) {
    request = &SegmentCustomizedPortraitPicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bda", APIVersion, "SegmentCustomizedPortraitPic")
    
    
    return
}

func NewSegmentCustomizedPortraitPicResponse() (response *SegmentCustomizedPortraitPicResponse) {
    response = &SegmentCustomizedPortraitPicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SegmentCustomizedPortraitPic
// 在前后景分割的基础上优化多分类分割，支持对头发、五官等的分割，既作为换发型、挂件等底层技术，也可用于抠人头、抠人脸等玩法
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONINSUFFICIENT = "FailedOperation.ImageResolutionInsufficient"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_PROFILENUMEXCEED = "FailedOperation.ProfileNumExceed"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SEGMENTFAILED = "FailedOperation.SegmentFailed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) SegmentCustomizedPortraitPic(request *SegmentCustomizedPortraitPicRequest) (response *SegmentCustomizedPortraitPicResponse, err error) {
    return c.SegmentCustomizedPortraitPicWithContext(context.Background(), request)
}

// SegmentCustomizedPortraitPic
// 在前后景分割的基础上优化多分类分割，支持对头发、五官等的分割，既作为换发型、挂件等底层技术，也可用于抠人头、抠人脸等玩法
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONINSUFFICIENT = "FailedOperation.ImageResolutionInsufficient"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_PROFILENUMEXCEED = "FailedOperation.ProfileNumExceed"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SEGMENTFAILED = "FailedOperation.SegmentFailed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) SegmentCustomizedPortraitPicWithContext(ctx context.Context, request *SegmentCustomizedPortraitPicRequest) (response *SegmentCustomizedPortraitPicResponse, err error) {
    if request == nil {
        request = NewSegmentCustomizedPortraitPicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bda", APIVersion, "SegmentCustomizedPortraitPic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SegmentCustomizedPortraitPic require credential")
    }

    request.SetContext(ctx)
    
    response = NewSegmentCustomizedPortraitPicResponse()
    err = c.Send(request, response)
    return
}

func NewSegmentPortraitPicRequest() (request *SegmentPortraitPicRequest) {
    request = &SegmentPortraitPicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bda", APIVersion, "SegmentPortraitPic")
    
    
    return
}

func NewSegmentPortraitPicResponse() (response *SegmentPortraitPicResponse) {
    response = &SegmentPortraitPicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SegmentPortraitPic
// 即二分类人像分割，识别传入图片中人体的完整轮廓，进行抠像。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGENOTFOREGROUND = "FailedOperation.ImageNotForeground"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONINSUFFICIENT = "FailedOperation.ImageResolutionInsufficient"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_PROFILENUMEXCEED = "FailedOperation.ProfileNumExceed"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SEGMENTFAILED = "FailedOperation.SegmentFailed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) SegmentPortraitPic(request *SegmentPortraitPicRequest) (response *SegmentPortraitPicResponse, err error) {
    return c.SegmentPortraitPicWithContext(context.Background(), request)
}

// SegmentPortraitPic
// 即二分类人像分割，识别传入图片中人体的完整轮廓，进行抠像。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGENOTFOREGROUND = "FailedOperation.ImageNotForeground"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONINSUFFICIENT = "FailedOperation.ImageResolutionInsufficient"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_PROFILENUMEXCEED = "FailedOperation.ProfileNumExceed"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SEGMENTFAILED = "FailedOperation.SegmentFailed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) SegmentPortraitPicWithContext(ctx context.Context, request *SegmentPortraitPicRequest) (response *SegmentPortraitPicResponse, err error) {
    if request == nil {
        request = NewSegmentPortraitPicRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bda", APIVersion, "SegmentPortraitPic")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SegmentPortraitPic require credential")
    }

    request.SetContext(ctx)
    
    response = NewSegmentPortraitPicResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateSegmentationTaskRequest() (request *TerminateSegmentationTaskRequest) {
    request = &TerminateSegmentationTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("bda", APIVersion, "TerminateSegmentationTask")
    
    
    return
}

func NewTerminateSegmentationTaskResponse() (response *TerminateSegmentationTaskResponse) {
    response = &TerminateSegmentationTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TerminateSegmentationTask
// 终止指定视频人像分割处理任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  FAILEDOPERATION_TERMINATETASKFAILED = "FailedOperation.TerminateTaskFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) TerminateSegmentationTask(request *TerminateSegmentationTaskRequest) (response *TerminateSegmentationTaskResponse, err error) {
    return c.TerminateSegmentationTaskWithContext(context.Background(), request)
}

// TerminateSegmentationTask
// 终止指定视频人像分割处理任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  FAILEDOPERATION_TERMINATETASKFAILED = "FailedOperation.TerminateTaskFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) TerminateSegmentationTaskWithContext(ctx context.Context, request *TerminateSegmentationTaskRequest) (response *TerminateSegmentationTaskResponse, err error) {
    if request == nil {
        request = NewTerminateSegmentationTaskRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "bda", APIVersion, "TerminateSegmentationTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("TerminateSegmentationTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewTerminateSegmentationTaskResponse()
    err = c.Send(request, response)
    return
}
