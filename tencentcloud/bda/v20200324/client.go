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

package v20200324

import (
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


func NewCreateGroupRequest() (request *CreateGroupRequest) {
    request = &CreateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bda", APIVersion, "CreateGroup")
    return
}

func NewCreateGroupResponse() (response *CreateGroupResponse) {
    response = &CreateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateGroup
// 用于创建一个空的人体库，如果人体库已存在返回错误。
//
// 
//
// 1个APPID下最多有2000W个人体轨迹（Trace），最多1W个人体库（Group）。
//
// 
//
// 单个人体库（Group）最多10W个人体轨迹（Trace）。
//
// 
//
// 单个人员（Person）最多添加 5 个人体轨迹（Trace）。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTTRACENUMEXCEED = "InvalidParameterValue.AccountTraceNumExceed"
//  INVALIDPARAMETERVALUE_BODYMODELVERSIONILLEGAL = "InvalidParameterValue.BodyModelVersionIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDALREADYEXIST = "InvalidParameterValue.GroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPNUMEXCEED = "InvalidParameterValue.GroupNumExceed"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  INVALIDPARAMETERVALUE_GROUPTRACENUMEXCEED = "InvalidParameterValue.GroupTraceNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    if request == nil {
        request = NewCreateGroupRequest()
    }
    response = NewCreateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePersonRequest() (request *CreatePersonRequest) {
    request = &CreatePersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bda", APIVersion, "CreatePerson")
    return
}

func NewCreatePersonResponse() (response *CreatePersonResponse) {
    response = &CreatePersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreatePerson
// 创建人员，添加对应人员的人体轨迹信息。
//
// 
//
// 请注意：
//
// - 我们希望您的输入为 严格符合轨迹图片 要求的图片。如果您输入的图片不符合轨迹图片要求，会对最终效果产生较大负面影响。请您尽量保证一个Trace中的图片人体清晰、无遮挡、连贯；
//
// - 一个人体轨迹（Trace）可以包含1-5张人体图片。提供越多质量高的人体图片有助于提升最终识别结果；
//
// - 无论您在单个Trace中提供了多少张人体图片，我们都将生成一个对应的轨迹（Trace）信息。即，Trace仅和本次输入的图片序列相关，和图片的个数无关；
//
// - 输入的图片组中，若有部分图片输入不合法（如图片大小过大、分辨率过大、无法解码等），我们将舍弃这部分图片，确保合法图片被正确搜索。即，我们将尽可能保证请求成功，去除不合法的输入；
//
// - 构成人体轨迹单张图片大小不得超过2M，分辨率不得超过1920*1080。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BODYRECTILLEGAL = "FailedOperation.BodyRectIllegal"
//  FAILEDOPERATION_BODYRECTNUMILLEGAL = "FailedOperation.BodyRectNumIllegal"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTTRACENUMEXCEED = "InvalidParameterValue.AccountTraceNumExceed"
//  INVALIDPARAMETERVALUE_BODYMODELVERSIONILLEGAL = "InvalidParameterValue.BodyModelVersionIllegal"
//  INVALIDPARAMETERVALUE_BODYRECTSEXCEED = "InvalidParameterValue.BodyRectsExceed"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPTRACENUMEXCEED = "InvalidParameterValue.GroupTraceNumExceed"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PERSONIDALREADYEXIST = "InvalidParameterValue.PersonIdAlreadyExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  INVALIDPARAMETERVALUE_TRACEBODYNUMEXCEED = "InvalidParameterValue.TraceBodyNumExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) CreatePerson(request *CreatePersonRequest) (response *CreatePersonResponse, err error) {
    if request == nil {
        request = NewCreatePersonRequest()
    }
    response = NewCreatePersonResponse()
    err = c.Send(request, response)
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
// 本接口为离线人像分割处理接口组中的提交任务接口，可以对提交的资源进行处理视频流/图片流识别视频作品中的人像区域，进行一键抠像、背景替换、人像虚化等后期处理。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBQUEUEFULL = "FailedOperation.JobQueueFull"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_TASKLIMITEXCEEDED = "FailedOperation.TaskLimitExceeded"
//  FAILEDOPERATION_TOOLARGEFILEERROR = "FailedOperation.TooLargeFileError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) CreateSegmentationTask(request *CreateSegmentationTaskRequest) (response *CreateSegmentationTaskResponse, err error) {
    if request == nil {
        request = NewCreateSegmentationTaskRequest()
    }
    response = NewCreateSegmentationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTraceRequest() (request *CreateTraceRequest) {
    request = &CreateTraceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bda", APIVersion, "CreateTrace")
    return
}

func NewCreateTraceResponse() (response *CreateTraceResponse) {
    response = &CreateTraceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateTrace
// 将一个人体轨迹添加到一个人员中。一个人员最多允许包含 5 个人体轨迹。同一人的人体轨迹越多，搜索识别效果越好。
//
// 
//
// >请注意：
//
// - 我们希望您的输入为 严格符合轨迹图片 要求的图片。如果您输入的图片不符合轨迹图片要求，会对最终效果产生较大负面影响。请您尽量保证一个Trace中的图片人体清晰、无遮挡、连贯。
//
// - 一个人体轨迹（Trace）可以包含1-5张人体图片。提供越多质量高的人体图片有助于提升最终识别结果。
//
// - 无论您在单个Trace中提供了多少张人体图片，我们都将生成一个对应的轨迹（Trace）信息。即，Trace仅和本次输入的图片序列相关，和图片的个数无关。
//
// - 输入的图片组中，若有部分图片输入不合法（如图片大小过大、分辨率过大、无法解码等），我们将舍弃这部分图片，确保合法图片被正确搜索。即，我们将尽可能保证请求成功，去除不合法的输入；
//
// - 构成人体轨迹单张图片大小限制为2M，分辨率限制为1920*1080。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BODYRECTILLEGAL = "FailedOperation.BodyRectIllegal"
//  FAILEDOPERATION_BODYRECTNUMILLEGAL = "FailedOperation.BodyRectNumIllegal"
//  FAILEDOPERATION_CREATETRACEFAILED = "FailedOperation.CreateTraceFailed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ACCOUNTTRACENUMEXCEED = "InvalidParameterValue.AccountTraceNumExceed"
//  INVALIDPARAMETERVALUE_BODYRECTSEXCEED = "InvalidParameterValue.BodyRectsExceed"
//  INVALIDPARAMETERVALUE_GROUPTRACENUMEXCEED = "InvalidParameterValue.GroupTraceNumExceed"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONTRACENUMEXCEED = "InvalidParameterValue.PersonTraceNumExceed"
//  INVALIDPARAMETERVALUE_TRACEBODYNUMEXCEED = "InvalidParameterValue.TraceBodyNumExceed"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) CreateTrace(request *CreateTraceRequest) (response *CreateTraceResponse, err error) {
    if request == nil {
        request = NewCreateTraceRequest()
    }
    response = NewCreateTraceResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
    request = &DeleteGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bda", APIVersion, "DeleteGroup")
    return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
    response = &DeleteGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteGroup
// 删除该人体库及包含的所有的人员。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
    if request == nil {
        request = NewDeleteGroupRequest()
    }
    response = NewDeleteGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeletePersonRequest() (request *DeletePersonRequest) {
    request = &DeletePersonRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bda", APIVersion, "DeletePerson")
    return
}

func NewDeletePersonResponse() (response *DeletePersonResponse) {
    response = &DeletePersonResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeletePerson
// 删除人员。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DeletePerson(request *DeletePersonRequest) (response *DeletePersonResponse, err error) {
    if request == nil {
        request = NewDeletePersonRequest()
    }
    response = NewDeletePersonResponse()
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
    if request == nil {
        request = NewDescribeSegmentationTaskRequest()
    }
    response = NewDescribeSegmentationTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDetectBodyRequest() (request *DetectBodyRequest) {
    request = &DetectBodyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bda", APIVersion, "DetectBody")
    return
}

func NewDetectBodyResponse() (response *DetectBodyResponse) {
    response = &DetectBodyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectBody
// 检测给定图片中的人体（Body）的位置信息及属性信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BODYQUALITYNOTQUALIFIED = "FailedOperation.BodyQualityNotQualified"
//  FAILEDOPERATION_IMAGEBODYDETECTFAILED = "FailedOperation.ImageBodyDetectFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NOBODYINPHOTO = "FailedOperation.NoBodyInPhoto"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DetectBody(request *DetectBodyRequest) (response *DetectBodyResponse, err error) {
    if request == nil {
        request = NewDetectBodyRequest()
    }
    response = NewDetectBodyResponse()
    err = c.Send(request, response)
    return
}

func NewDetectBodyJointsRequest() (request *DetectBodyJointsRequest) {
    request = &DetectBodyJointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bda", APIVersion, "DetectBodyJoints")
    return
}

func NewDetectBodyJointsResponse() (response *DetectBodyJointsResponse) {
    response = &DetectBodyJointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectBodyJoints
// 检测图片中人体的14个关键点。建议用于人体图像清晰、无遮挡的场景。支持一张图片中存在多个人体的识别。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BODYFEATUREFAIL = "FailedOperation.BodyFeatureFail"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NOBODYINPHOTO = "FailedOperation.NoBodyInPhoto"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) DetectBodyJoints(request *DetectBodyJointsRequest) (response *DetectBodyJointsResponse, err error) {
    if request == nil {
        request = NewDetectBodyJointsRequest()
    }
    response = NewDetectBodyJointsResponse()
    err = c.Send(request, response)
    return
}

func NewGetGroupListRequest() (request *GetGroupListRequest) {
    request = &GetGroupListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bda", APIVersion, "GetGroupList")
    return
}

func NewGetGroupListResponse() (response *GetGroupListResponse) {
    response = &GetGroupListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetGroupList
// 获取人体库列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) GetGroupList(request *GetGroupListRequest) (response *GetGroupListResponse, err error) {
    if request == nil {
        request = NewGetGroupListRequest()
    }
    response = NewGetGroupListResponse()
    err = c.Send(request, response)
    return
}

func NewGetPersonListRequest() (request *GetPersonListRequest) {
    request = &GetPersonListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bda", APIVersion, "GetPersonList")
    return
}

func NewGetPersonListResponse() (response *GetPersonListResponse) {
    response = &GetPersonListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetPersonList
// 获取指定人体库中的人员列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  INVALIDPARAMETERVALUE_OFFSETEXCEED = "InvalidParameterValue.OffsetExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) GetPersonList(request *GetPersonListRequest) (response *GetPersonListResponse, err error) {
    if request == nil {
        request = NewGetPersonListRequest()
    }
    response = NewGetPersonListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyGroupRequest() (request *ModifyGroupRequest) {
    request = &ModifyGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bda", APIVersion, "ModifyGroup")
    return
}

func NewModifyGroupResponse() (response *ModifyGroupResponse) {
    response = &ModifyGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyGroup
// 修改人体库名称、备注。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_GROUPNAMEALREADYEXIST = "InvalidParameterValue.GroupNameAlreadyExist"
//  INVALIDPARAMETERVALUE_GROUPNAMETOOLONG = "InvalidParameterValue.GroupNameTooLong"
//  INVALIDPARAMETERVALUE_GROUPTAGTOOLONG = "InvalidParameterValue.GroupTagTooLong"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) ModifyGroup(request *ModifyGroupRequest) (response *ModifyGroupResponse, err error) {
    if request == nil {
        request = NewModifyGroupRequest()
    }
    response = NewModifyGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyPersonInfoRequest() (request *ModifyPersonInfoRequest) {
    request = &ModifyPersonInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bda", APIVersion, "ModifyPersonInfo")
    return
}

func NewModifyPersonInfoResponse() (response *ModifyPersonInfoResponse) {
    response = &ModifyPersonInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyPersonInfo
// 修改人员信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_PERSONIDILLEGAL = "InvalidParameterValue.PersonIdIllegal"
//  INVALIDPARAMETERVALUE_PERSONIDNOTEXIST = "InvalidParameterValue.PersonIdNotExist"
//  INVALIDPARAMETERVALUE_PERSONIDTOOLONG = "InvalidParameterValue.PersonIdTooLong"
//  INVALIDPARAMETERVALUE_PERSONNAMEILLEGAL = "InvalidParameterValue.PersonNameIllegal"
//  INVALIDPARAMETERVALUE_PERSONNAMETOOLONG = "InvalidParameterValue.PersonNameTooLong"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) ModifyPersonInfo(request *ModifyPersonInfoRequest) (response *ModifyPersonInfoResponse, err error) {
    if request == nil {
        request = NewModifyPersonInfoRequest()
    }
    response = NewModifyPersonInfoResponse()
    err = c.Send(request, response)
    return
}

func NewSearchTraceRequest() (request *SearchTraceRequest) {
    request = &SearchTraceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bda", APIVersion, "SearchTrace")
    return
}

func NewSearchTraceResponse() (response *SearchTraceResponse) {
    response = &SearchTraceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchTrace
// 本接口用于对一组待识别的人体轨迹（Trace）图片，在人体库中识别出最相似的 TopK 人体，按照相似度从大到小排列。
//
// 
//
// 人体轨迹（Trace）图片要求：图片中当且仅包含一个人体。人体完整、无遮挡。
//
// 
//
// > 请注意：
//
// - 我们希望您的输入为严格符合轨迹图片要求的图片。如果您输入的图片不符合轨迹图片要求，会对最终效果产生较大负面影响；
//
// - 人体轨迹，是一个包含1-5张图片的图片序列。您可以输入1张图片作为轨迹，也可以输入多张。单个轨迹中包含越多符合质量的图片，搜索效果越好。
//
// - 构成人体轨迹单张图片大小不得超过2M，分辨率不得超过1920*1080。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BODYRECTILLEGAL = "FailedOperation.BodyRectIllegal"
//  FAILEDOPERATION_BODYRECTNUMILLEGAL = "FailedOperation.BodyRectNumIllegal"
//  FAILEDOPERATION_GROUPEMPTY = "FailedOperation.GroupEmpty"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_BODYRECTSEXCEED = "InvalidParameterValue.BodyRectsExceed"
//  INVALIDPARAMETERVALUE_GROUPIDILLEGAL = "InvalidParameterValue.GroupIdIllegal"
//  INVALIDPARAMETERVALUE_GROUPIDNOTEXIST = "InvalidParameterValue.GroupIdNotExist"
//  INVALIDPARAMETERVALUE_GROUPIDTOOLONG = "InvalidParameterValue.GroupIdTooLong"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_SEARCHPERSONSEXCEED = "InvalidParameterValue.SearchPersonsExceed"
//  INVALIDPARAMETERVALUE_TRACEBODYNUMEXCEED = "InvalidParameterValue.TraceBodyNumExceed"
//  INVALIDPARAMETERVALUE_TRACEMATCHTHRESHOLDILLEGAL = "InvalidParameterValue.TraceMatchThresholdIllegal"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) SearchTrace(request *SearchTraceRequest) (response *SearchTraceResponse, err error) {
    if request == nil {
        request = NewSearchTraceRequest()
    }
    response = NewSearchTraceResponse()
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
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFacedetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_PROFILENUMEXCEED = "FailedOperation.ProfileNumExceed"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
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
    if request == nil {
        request = NewSegmentCustomizedPortraitPicRequest()
    }
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
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONINSUFFICIENT = "FailedOperation.ImageResolutionInsufficient"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_PROFILENUMEXCEED = "FailedOperation.ProfileNumExceed"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SEGMENTFAILED = "FailedOperation.SegmentFailed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) SegmentPortraitPic(request *SegmentPortraitPicRequest) (response *SegmentPortraitPicResponse, err error) {
    if request == nil {
        request = NewSegmentPortraitPicRequest()
    }
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
    if request == nil {
        request = NewTerminateSegmentationTaskRequest()
    }
    response = NewTerminateSegmentationTaskResponse()
    err = c.Send(request, response)
    return
}
