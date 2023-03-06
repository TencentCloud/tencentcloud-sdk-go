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

package v20190614

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-06-14"

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


func NewCloseAsyncRecognitionTaskRequest() (request *CloseAsyncRecognitionTaskRequest) {
    request = &CloseAsyncRecognitionTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "CloseAsyncRecognitionTask")
    
    
    return
}

func NewCloseAsyncRecognitionTaskResponse() (response *CloseAsyncRecognitionTaskResponse) {
    response = &CloseAsyncRecognitionTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CloseAsyncRecognitionTask
// 本接口用于关闭语音流异步识别任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOSUCHTASK = "FailedOperation.NoSuchTask"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CloseAsyncRecognitionTask(request *CloseAsyncRecognitionTaskRequest) (response *CloseAsyncRecognitionTaskResponse, err error) {
    return c.CloseAsyncRecognitionTaskWithContext(context.Background(), request)
}

// CloseAsyncRecognitionTask
// 本接口用于关闭语音流异步识别任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOSUCHTASK = "FailedOperation.NoSuchTask"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) CloseAsyncRecognitionTaskWithContext(ctx context.Context, request *CloseAsyncRecognitionTaskRequest) (response *CloseAsyncRecognitionTaskResponse, err error) {
    if request == nil {
        request = NewCloseAsyncRecognitionTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseAsyncRecognitionTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseAsyncRecognitionTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAsrVocabRequest() (request *CreateAsrVocabRequest) {
    request = &CreateAsrVocabRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "CreateAsrVocab")
    
    
    return
}

func NewCreateAsrVocabResponse() (response *CreateAsrVocabResponse) {
    response = &CreateAsrVocabResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAsrVocab
// 用户通过本接口进行热词表的创建。
//
// <br>•   默认最多可创建30个热词表。
//
// <br>•   每个热词表最多可添加128个词，每个词最长10个字，不能超出限制。
//
// <br>•   热词表可以通过数组或者本地文件形式上传。
//
// <br>•   本地文件必须为UTF-8编码格式，每行仅添加一个热词且不能包含标点和特殊字符。
//
// <br>•   热词权重取值范围为[1,10]之间的整数，权重越大代表该词被识别出来的概率越大。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORCONTENTLENGTH = "InvalidParameter.ErrorContentlength"
//  INVALIDPARAMETER_ERRORPARAMSMISSING = "InvalidParameter.ErrorParamsMissing"
//  INVALIDPARAMETER_ERRORPARSEQUEST = "InvalidParameter.ErrorParsequest"
//  INVALIDPARAMETER_INVALIDVOCABSTATE = "InvalidParameter.InvalidVocabState"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ERRORINVALIDAPPID = "InvalidParameterValue.ErrorInvalidAppid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDCLIENTIP = "InvalidParameterValue.ErrorInvalidClientip"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERLENGTH = "InvalidParameterValue.InvalidParameterLength"
//  INVALIDPARAMETERVALUE_INVALIDVOCABID = "InvalidParameterValue.InvalidVocabId"
//  INVALIDPARAMETERVALUE_INVALIDWORDWEIGHT = "InvalidParameterValue.InvalidWordWeight"
//  INVALIDPARAMETERVALUE_INVALIDWORDWEIGHTSTR = "InvalidParameterValue.InvalidWordWeightStr"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_VOCABFULL = "LimitExceeded.VocabFull"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateAsrVocab(request *CreateAsrVocabRequest) (response *CreateAsrVocabResponse, err error) {
    return c.CreateAsrVocabWithContext(context.Background(), request)
}

// CreateAsrVocab
// 用户通过本接口进行热词表的创建。
//
// <br>•   默认最多可创建30个热词表。
//
// <br>•   每个热词表最多可添加128个词，每个词最长10个字，不能超出限制。
//
// <br>•   热词表可以通过数组或者本地文件形式上传。
//
// <br>•   本地文件必须为UTF-8编码格式，每行仅添加一个热词且不能包含标点和特殊字符。
//
// <br>•   热词权重取值范围为[1,10]之间的整数，权重越大代表该词被识别出来的概率越大。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORCONTENTLENGTH = "InvalidParameter.ErrorContentlength"
//  INVALIDPARAMETER_ERRORPARAMSMISSING = "InvalidParameter.ErrorParamsMissing"
//  INVALIDPARAMETER_ERRORPARSEQUEST = "InvalidParameter.ErrorParsequest"
//  INVALIDPARAMETER_INVALIDVOCABSTATE = "InvalidParameter.InvalidVocabState"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ERRORINVALIDAPPID = "InvalidParameterValue.ErrorInvalidAppid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDCLIENTIP = "InvalidParameterValue.ErrorInvalidClientip"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERLENGTH = "InvalidParameterValue.InvalidParameterLength"
//  INVALIDPARAMETERVALUE_INVALIDVOCABID = "InvalidParameterValue.InvalidVocabId"
//  INVALIDPARAMETERVALUE_INVALIDWORDWEIGHT = "InvalidParameterValue.InvalidWordWeight"
//  INVALIDPARAMETERVALUE_INVALIDWORDWEIGHTSTR = "InvalidParameterValue.InvalidWordWeightStr"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_VOCABFULL = "LimitExceeded.VocabFull"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateAsrVocabWithContext(ctx context.Context, request *CreateAsrVocabRequest) (response *CreateAsrVocabResponse, err error) {
    if request == nil {
        request = NewCreateAsrVocabRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAsrVocab require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAsrVocabResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAsyncRecognitionTaskRequest() (request *CreateAsyncRecognitionTaskRequest) {
    request = &CreateAsyncRecognitionTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "CreateAsyncRecognitionTask")
    
    
    return
}

func NewCreateAsyncRecognitionTaskResponse() (response *CreateAsyncRecognitionTaskResponse) {
    response = &CreateAsyncRecognitionTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateAsyncRecognitionTask
// 本接口用于对语音流进行准实时识别，通过异步回调来返回识别结果。适用于直播审核等场景。
//
// <br>• 支持rtmp、rtsp等流媒体协议，以及各类基于http协议的直播流(不支持hls)
//
// <br>• 音频流时长无限制，服务会自动拉取音频流数据，若连续10分钟拉不到流数据时，服务会终止识别任务
//
// <br>• 服务通过回调的方式来提供识别结果，用户需要提供CallbackUrl。回调时机为一小段话(最长15秒)回调一次。
//
// <br>• 签名方法参考 [公共参数](https://cloud.tencent.com/document/api/1093/35640) 中签名方法v3。
//
// <br>• 默认单账号限制并发数为20路，如您有提高并发限制的需求，请提[工单](https://console.cloud.tencent.com/workorder/category)进行咨询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERHASNOFREEAMOUNT = "FailedOperation.UserHasNoFreeAmount"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateAsyncRecognitionTask(request *CreateAsyncRecognitionTaskRequest) (response *CreateAsyncRecognitionTaskResponse, err error) {
    return c.CreateAsyncRecognitionTaskWithContext(context.Background(), request)
}

// CreateAsyncRecognitionTask
// 本接口用于对语音流进行准实时识别，通过异步回调来返回识别结果。适用于直播审核等场景。
//
// <br>• 支持rtmp、rtsp等流媒体协议，以及各类基于http协议的直播流(不支持hls)
//
// <br>• 音频流时长无限制，服务会自动拉取音频流数据，若连续10分钟拉不到流数据时，服务会终止识别任务
//
// <br>• 服务通过回调的方式来提供识别结果，用户需要提供CallbackUrl。回调时机为一小段话(最长15秒)回调一次。
//
// <br>• 签名方法参考 [公共参数](https://cloud.tencent.com/document/api/1093/35640) 中签名方法v3。
//
// <br>• 默认单账号限制并发数为20路，如您有提高并发限制的需求，请提[工单](https://console.cloud.tencent.com/workorder/category)进行咨询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERHASNOFREEAMOUNT = "FailedOperation.UserHasNoFreeAmount"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateAsyncRecognitionTaskWithContext(ctx context.Context, request *CreateAsyncRecognitionTaskRequest) (response *CreateAsyncRecognitionTaskResponse, err error) {
    if request == nil {
        request = NewCreateAsyncRecognitionTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAsyncRecognitionTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAsyncRecognitionTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCustomizationRequest() (request *CreateCustomizationRequest) {
    request = &CreateCustomizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "CreateCustomization")
    
    
    return
}

func NewCreateCustomizationResponse() (response *CreateCustomizationResponse) {
    response = &CreateCustomizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateCustomization
// 用户使用该接口可以创建自学习模型，以供识别调用。
//
// 
//
// 注意：调用该接口后，模型会自动训练。新建模型成功后，调用ModifyCustomizationState接口修改为上线状态，即可在识别请求中使用对应模型ID。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CHECKRESOURCERESPONSECODEERROR = "AuthFailure.CheckResourceResponseCodeError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ERRORDOWNFILE = "FailedOperation.ErrorDownFile"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FILEENCODE = "InvalidParameter.FileEncode"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_CUSTOMIZATIONFULL = "LimitExceeded.CustomizationFull"
func (c *Client) CreateCustomization(request *CreateCustomizationRequest) (response *CreateCustomizationResponse, err error) {
    return c.CreateCustomizationWithContext(context.Background(), request)
}

// CreateCustomization
// 用户使用该接口可以创建自学习模型，以供识别调用。
//
// 
//
// 注意：调用该接口后，模型会自动训练。新建模型成功后，调用ModifyCustomizationState接口修改为上线状态，即可在识别请求中使用对应模型ID。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CHECKRESOURCERESPONSECODEERROR = "AuthFailure.CheckResourceResponseCodeError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ERRORDOWNFILE = "FailedOperation.ErrorDownFile"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FILEENCODE = "InvalidParameter.FileEncode"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED_CUSTOMIZATIONFULL = "LimitExceeded.CustomizationFull"
func (c *Client) CreateCustomizationWithContext(ctx context.Context, request *CreateCustomizationRequest) (response *CreateCustomizationResponse, err error) {
    if request == nil {
        request = NewCreateCustomizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateCustomization require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateCustomizationResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRecTaskRequest() (request *CreateRecTaskRequest) {
    request = &CreateRecTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "CreateRecTask")
    
    
    return
}

func NewCreateRecTaskResponse() (response *CreateRecTaskResponse) {
    response = &CreateRecTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateRecTask
// 本接口服务对时长5小时以内的录音文件进行识别，异步返回识别全部结果。
//
// • 支持中文普通话、英语、粤语、日语、越南语、马来语、上海话、四川话、武汉话、贵阳话、昆明话、西安话、郑州话、太原话、兰州话、银川话、西宁话、南京话、合肥话、南昌话、长沙话、苏州话、杭州话、济南话、天津话、石家庄话、黑龙江话、吉林话、辽宁话。
//
// • 支持wav、mp3、m4a、flv、mp4、wma、3gp、amr、aac、ogg-opus、flac格式。
//
// • 支持语音 URL 和本地语音文件两种请求方式。语音 URL 的音频时长不能长于5小时，文件大小不超过1GB。本地语音文件调用不能大于5MB。
//
// • 提交录音文件识别请求后，在3小时内完成识别（大多数情况下1小时音频约3分钟以内完成识别，半小时内发送超过1000小时录音或者2万条识别任务的除外），识别结果在服务端可保存7天。
//
// • 支持回调或轮询的方式获取结果，结果获取请参考[ 录音文件识别结果查询](https://cloud.tencent.com/document/product/1093/37822)。
//
// •   生成字幕场景可设置参数ResTextFormat为3，解析ResultDetail结构生成字幕，可参考 [生成字幕最佳实践](https://cloud.tencent.com/document/product/1093/84291)。
//
// •   签名方法参考 [公共参数](https://cloud.tencent.com/document/api/1093/35640) 中签名方法v3。
//
// • 默认接口请求频率限制：20次/秒，如您有提高请求频率限制的需求，请提[工单](https://console.cloud.tencent.com/workorder/category)进行咨询。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CHECKAUTHINFOFAILED = "FailedOperation.CheckAuthInfoFailed"
//  FAILEDOPERATION_ERRORDOWNFILE = "FailedOperation.ErrorDownFile"
//  FAILEDOPERATION_ERRORRECOGNIZE = "FailedOperation.ErrorRecognize"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERHASNOFREEAMOUNT = "FailedOperation.UserHasNoFreeAmount"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR_ERRORDOWNFILE = "InternalError.ErrorDownFile"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_FAILACCESSREDIS = "InternalError.FailAccessRedis"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_UINLIMITEXCEEDED = "RequestLimitExceeded.UinLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateRecTask(request *CreateRecTaskRequest) (response *CreateRecTaskResponse, err error) {
    return c.CreateRecTaskWithContext(context.Background(), request)
}

// CreateRecTask
// 本接口服务对时长5小时以内的录音文件进行识别，异步返回识别全部结果。
//
// • 支持中文普通话、英语、粤语、日语、越南语、马来语、上海话、四川话、武汉话、贵阳话、昆明话、西安话、郑州话、太原话、兰州话、银川话、西宁话、南京话、合肥话、南昌话、长沙话、苏州话、杭州话、济南话、天津话、石家庄话、黑龙江话、吉林话、辽宁话。
//
// • 支持wav、mp3、m4a、flv、mp4、wma、3gp、amr、aac、ogg-opus、flac格式。
//
// • 支持语音 URL 和本地语音文件两种请求方式。语音 URL 的音频时长不能长于5小时，文件大小不超过1GB。本地语音文件调用不能大于5MB。
//
// • 提交录音文件识别请求后，在3小时内完成识别（大多数情况下1小时音频约3分钟以内完成识别，半小时内发送超过1000小时录音或者2万条识别任务的除外），识别结果在服务端可保存7天。
//
// • 支持回调或轮询的方式获取结果，结果获取请参考[ 录音文件识别结果查询](https://cloud.tencent.com/document/product/1093/37822)。
//
// •   生成字幕场景可设置参数ResTextFormat为3，解析ResultDetail结构生成字幕，可参考 [生成字幕最佳实践](https://cloud.tencent.com/document/product/1093/84291)。
//
// •   签名方法参考 [公共参数](https://cloud.tencent.com/document/api/1093/35640) 中签名方法v3。
//
// • 默认接口请求频率限制：20次/秒，如您有提高请求频率限制的需求，请提[工单](https://console.cloud.tencent.com/workorder/category)进行咨询。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CHECKAUTHINFOFAILED = "FailedOperation.CheckAuthInfoFailed"
//  FAILEDOPERATION_ERRORDOWNFILE = "FailedOperation.ErrorDownFile"
//  FAILEDOPERATION_ERRORRECOGNIZE = "FailedOperation.ErrorRecognize"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERHASNOFREEAMOUNT = "FailedOperation.UserHasNoFreeAmount"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR_ERRORDOWNFILE = "InternalError.ErrorDownFile"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_FAILACCESSREDIS = "InternalError.FailAccessRedis"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_UINLIMITEXCEEDED = "RequestLimitExceeded.UinLimitExceeded"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateRecTaskWithContext(ctx context.Context, request *CreateRecTaskRequest) (response *CreateRecTaskResponse, err error) {
    if request == nil {
        request = NewCreateRecTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRecTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRecTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAsrVocabRequest() (request *DeleteAsrVocabRequest) {
    request = &DeleteAsrVocabRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "DeleteAsrVocab")
    
    
    return
}

func NewDeleteAsrVocabResponse() (response *DeleteAsrVocabResponse) {
    response = &DeleteAsrVocabResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteAsrVocab
// 用户通过本接口进行热词表的删除。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORCONTENTLENGTH = "InvalidParameter.ErrorContentlength"
//  INVALIDPARAMETER_ERRORPARAMSMISSING = "InvalidParameter.ErrorParamsMissing"
//  INVALIDPARAMETER_ERRORPARSEQUEST = "InvalidParameter.ErrorParsequest"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ERRORINVALIDAPPID = "InvalidParameterValue.ErrorInvalidAppid"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERLENGTH = "InvalidParameterValue.InvalidParameterLength"
//  INVALIDPARAMETERVALUE_INVALIDVOCABID = "InvalidParameterValue.InvalidVocabId"
//  INVALIDPARAMETERVALUE_INVALIDWORDWEIGHT = "InvalidParameterValue.InvalidWordWeight"
//  INVALIDPARAMETERVALUE_INVALIDWORDWEIGHTSTR = "InvalidParameterValue.InvalidWordWeightStr"
func (c *Client) DeleteAsrVocab(request *DeleteAsrVocabRequest) (response *DeleteAsrVocabResponse, err error) {
    return c.DeleteAsrVocabWithContext(context.Background(), request)
}

// DeleteAsrVocab
// 用户通过本接口进行热词表的删除。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORCONTENTLENGTH = "InvalidParameter.ErrorContentlength"
//  INVALIDPARAMETER_ERRORPARAMSMISSING = "InvalidParameter.ErrorParamsMissing"
//  INVALIDPARAMETER_ERRORPARSEQUEST = "InvalidParameter.ErrorParsequest"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ERRORINVALIDAPPID = "InvalidParameterValue.ErrorInvalidAppid"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERLENGTH = "InvalidParameterValue.InvalidParameterLength"
//  INVALIDPARAMETERVALUE_INVALIDVOCABID = "InvalidParameterValue.InvalidVocabId"
//  INVALIDPARAMETERVALUE_INVALIDWORDWEIGHT = "InvalidParameterValue.InvalidWordWeight"
//  INVALIDPARAMETERVALUE_INVALIDWORDWEIGHTSTR = "InvalidParameterValue.InvalidWordWeightStr"
func (c *Client) DeleteAsrVocabWithContext(ctx context.Context, request *DeleteAsrVocabRequest) (response *DeleteAsrVocabResponse, err error) {
    if request == nil {
        request = NewDeleteAsrVocabRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAsrVocab require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAsrVocabResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteCustomizationRequest() (request *DeleteCustomizationRequest) {
    request = &DeleteCustomizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "DeleteCustomization")
    
    
    return
}

func NewDeleteCustomizationResponse() (response *DeleteCustomizationResponse) {
    response = &DeleteCustomizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteCustomization
// 用户通过该接口可以删除自学习模型
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CHECKRESOURCERESPONSECODEERROR = "AuthFailure.CheckResourceResponseCodeError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MODELSTATE = "InvalidParameter.ModelState"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODELID = "InvalidParameterValue.ModelId"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DeleteCustomization(request *DeleteCustomizationRequest) (response *DeleteCustomizationResponse, err error) {
    return c.DeleteCustomizationWithContext(context.Background(), request)
}

// DeleteCustomization
// 用户通过该接口可以删除自学习模型
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CHECKRESOURCERESPONSECODEERROR = "AuthFailure.CheckResourceResponseCodeError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MODELSTATE = "InvalidParameter.ModelState"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODELID = "InvalidParameterValue.ModelId"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DeleteCustomizationWithContext(ctx context.Context, request *DeleteCustomizationRequest) (response *DeleteCustomizationResponse, err error) {
    if request == nil {
        request = NewDeleteCustomizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteCustomization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteCustomizationResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAsyncRecognitionTasksRequest() (request *DescribeAsyncRecognitionTasksRequest) {
    request = &DescribeAsyncRecognitionTasksRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "DescribeAsyncRecognitionTasks")
    
    
    return
}

func NewDescribeAsyncRecognitionTasksResponse() (response *DescribeAsyncRecognitionTasksResponse) {
    response = &DescribeAsyncRecognitionTasksResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeAsyncRecognitionTasks
// 本接口用于查询当前在运行的语音流异步识别任务列表。
//
// <br>•   签名方法参考 [公共参数](https://cloud.tencent.com/document/api/1093/35640) 中签名方法v3。
//
// 可能返回的错误码:
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
func (c *Client) DescribeAsyncRecognitionTasks(request *DescribeAsyncRecognitionTasksRequest) (response *DescribeAsyncRecognitionTasksResponse, err error) {
    return c.DescribeAsyncRecognitionTasksWithContext(context.Background(), request)
}

// DescribeAsyncRecognitionTasks
// 本接口用于查询当前在运行的语音流异步识别任务列表。
//
// <br>•   签名方法参考 [公共参数](https://cloud.tencent.com/document/api/1093/35640) 中签名方法v3。
//
// 可能返回的错误码:
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
func (c *Client) DescribeAsyncRecognitionTasksWithContext(ctx context.Context, request *DescribeAsyncRecognitionTasksRequest) (response *DescribeAsyncRecognitionTasksResponse, err error) {
    if request == nil {
        request = NewDescribeAsyncRecognitionTasksRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAsyncRecognitionTasks require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAsyncRecognitionTasksResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskStatusRequest() (request *DescribeTaskStatusRequest) {
    request = &DescribeTaskStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "DescribeTaskStatus")
    
    
    return
}

func NewDescribeTaskStatusResponse() (response *DescribeTaskStatusResponse) {
    response = &DescribeTaskStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTaskStatus
// 在调用录音文件识别请求接口后，有回调和轮询两种方式获取识别结果。
//
// <br>• 当采用回调方式时，识别完成后会将结果通过 POST 请求的形式通知到用户在请求时填写的回调 URL，具体请参见[ 录音识别结果回调 ](https://cloud.tencent.com/document/product/1093/52632)。
//
// <br>• 当采用轮询方式时，需要主动提交任务ID来轮询识别结果，共有任务成功、等待、执行中和失败四种结果，具体信息请参见下文说明。
//
// <br>•   请求方法为 HTTP POST , Content-Type为"application/json; charset=utf-8"
//
// <br>•   签名方法参考 [公共参数](https://cloud.tencent.com/document/api/1093/35640) 中签名方法v3。
//
// <br>•   默认接口请求频率限制：50次/秒，如您有提高请求频率限制的需求，请提[工单](https://console.cloud.tencent.com/workorder/category)进行咨询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORDOWNFILE = "FailedOperation.ErrorDownFile"
//  FAILEDOPERATION_ERRORRECOGNIZE = "FailedOperation.ErrorRecognize"
//  FAILEDOPERATION_NOSUCHTASK = "FailedOperation.NoSuchTask"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERHASNOFREEAMOUNT = "FailedOperation.UserHasNoFreeAmount"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_FAILACCESSREDIS = "InternalError.FailAccessRedis"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTaskStatus(request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    return c.DescribeTaskStatusWithContext(context.Background(), request)
}

// DescribeTaskStatus
// 在调用录音文件识别请求接口后，有回调和轮询两种方式获取识别结果。
//
// <br>• 当采用回调方式时，识别完成后会将结果通过 POST 请求的形式通知到用户在请求时填写的回调 URL，具体请参见[ 录音识别结果回调 ](https://cloud.tencent.com/document/product/1093/52632)。
//
// <br>• 当采用轮询方式时，需要主动提交任务ID来轮询识别结果，共有任务成功、等待、执行中和失败四种结果，具体信息请参见下文说明。
//
// <br>•   请求方法为 HTTP POST , Content-Type为"application/json; charset=utf-8"
//
// <br>•   签名方法参考 [公共参数](https://cloud.tencent.com/document/api/1093/35640) 中签名方法v3。
//
// <br>•   默认接口请求频率限制：50次/秒，如您有提高请求频率限制的需求，请提[工单](https://console.cloud.tencent.com/workorder/category)进行咨询。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORDOWNFILE = "FailedOperation.ErrorDownFile"
//  FAILEDOPERATION_ERRORRECOGNIZE = "FailedOperation.ErrorRecognize"
//  FAILEDOPERATION_NOSUCHTASK = "FailedOperation.NoSuchTask"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERHASNOFREEAMOUNT = "FailedOperation.UserHasNoFreeAmount"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_FAILACCESSREDIS = "InternalError.FailAccessRedis"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) DescribeTaskStatusWithContext(ctx context.Context, request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
    if request == nil {
        request = NewDescribeTaskStatusRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTaskStatus require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadAsrVocabRequest() (request *DownloadAsrVocabRequest) {
    request = &DownloadAsrVocabRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "DownloadAsrVocab")
    
    
    return
}

func NewDownloadAsrVocabResponse() (response *DownloadAsrVocabResponse) {
    response = &DownloadAsrVocabResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DownloadAsrVocab
// 用户通过本接口进行热词表的下载，获得词表权重文件形式的 base64 值，文件形式为通过 “|” 分割的词和权重，即 word|weight 的形式。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_FAILACCESSREDIS = "InternalError.FailAccessRedis"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDVOCABID = "InvalidParameterValue.InvalidVocabId"
func (c *Client) DownloadAsrVocab(request *DownloadAsrVocabRequest) (response *DownloadAsrVocabResponse, err error) {
    return c.DownloadAsrVocabWithContext(context.Background(), request)
}

// DownloadAsrVocab
// 用户通过本接口进行热词表的下载，获得词表权重文件形式的 base64 值，文件形式为通过 “|” 分割的词和权重，即 word|weight 的形式。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_FAILACCESSREDIS = "InternalError.FailAccessRedis"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDVOCABID = "InvalidParameterValue.InvalidVocabId"
func (c *Client) DownloadAsrVocabWithContext(ctx context.Context, request *DownloadAsrVocabRequest) (response *DownloadAsrVocabResponse, err error) {
    if request == nil {
        request = NewDownloadAsrVocabRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadAsrVocab require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadAsrVocabResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadCustomizationRequest() (request *DownloadCustomizationRequest) {
    request = &DownloadCustomizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "DownloadCustomization")
    
    
    return
}

func NewDownloadCustomizationResponse() (response *DownloadCustomizationResponse) {
    response = &DownloadCustomizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DownloadCustomization
// 用户通过该接口可以下载自学习模型的语料
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CHECKRESOURCERESPONSECODEERROR = "AuthFailure.CheckResourceResponseCodeError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_MODELID = "InvalidParameterValue.ModelId"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DownloadCustomization(request *DownloadCustomizationRequest) (response *DownloadCustomizationResponse, err error) {
    return c.DownloadCustomizationWithContext(context.Background(), request)
}

// DownloadCustomization
// 用户通过该接口可以下载自学习模型的语料
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CHECKRESOURCERESPONSECODEERROR = "AuthFailure.CheckResourceResponseCodeError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_MODELID = "InvalidParameterValue.ModelId"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DownloadCustomizationWithContext(ctx context.Context, request *DownloadCustomizationRequest) (response *DownloadCustomizationResponse, err error) {
    if request == nil {
        request = NewDownloadCustomizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadCustomization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadCustomizationResponse()
    err = c.Send(request, response)
    return
}

func NewGetAsrVocabRequest() (request *GetAsrVocabRequest) {
    request = &GetAsrVocabRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "GetAsrVocab")
    
    
    return
}

func NewGetAsrVocabResponse() (response *GetAsrVocabResponse) {
    response = &GetAsrVocabResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetAsrVocab
// 用户根据词表的ID可以获取对应的热词表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDVOCABID = "InvalidParameterValue.InvalidVocabId"
func (c *Client) GetAsrVocab(request *GetAsrVocabRequest) (response *GetAsrVocabResponse, err error) {
    return c.GetAsrVocabWithContext(context.Background(), request)
}

// GetAsrVocab
// 用户根据词表的ID可以获取对应的热词表信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDVOCABID = "InvalidParameterValue.InvalidVocabId"
func (c *Client) GetAsrVocabWithContext(ctx context.Context, request *GetAsrVocabRequest) (response *GetAsrVocabResponse, err error) {
    if request == nil {
        request = NewGetAsrVocabRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAsrVocab require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAsrVocabResponse()
    err = c.Send(request, response)
    return
}

func NewGetAsrVocabListRequest() (request *GetAsrVocabListRequest) {
    request = &GetAsrVocabListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "GetAsrVocabList")
    
    
    return
}

func NewGetAsrVocabListResponse() (response *GetAsrVocabListResponse) {
    response = &GetAsrVocabListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetAsrVocabList
// 用户通过该接口，可获得所有的热词表及其信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_CHECKRESOURCERESPONSECODEERROR = "AuthFailure.CheckResourceResponseCodeError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_FAILACCESSREDIS = "InternalError.FailAccessRedis"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetAsrVocabList(request *GetAsrVocabListRequest) (response *GetAsrVocabListResponse, err error) {
    return c.GetAsrVocabListWithContext(context.Background(), request)
}

// GetAsrVocabList
// 用户通过该接口，可获得所有的热词表及其信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_CHECKRESOURCERESPONSECODEERROR = "AuthFailure.CheckResourceResponseCodeError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_FAILACCESSREDIS = "InternalError.FailAccessRedis"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetAsrVocabListWithContext(ctx context.Context, request *GetAsrVocabListRequest) (response *GetAsrVocabListResponse, err error) {
    if request == nil {
        request = NewGetAsrVocabListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAsrVocabList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAsrVocabListResponse()
    err = c.Send(request, response)
    return
}

func NewGetCustomizationListRequest() (request *GetCustomizationListRequest) {
    request = &GetCustomizationListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "GetCustomizationList")
    
    
    return
}

func NewGetCustomizationListResponse() (response *GetCustomizationListResponse) {
    response = &GetCustomizationListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetCustomizationList
// 查询自学习模型列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) GetCustomizationList(request *GetCustomizationListRequest) (response *GetCustomizationListResponse, err error) {
    return c.GetCustomizationListWithContext(context.Background(), request)
}

// GetCustomizationList
// 查询自学习模型列表
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
func (c *Client) GetCustomizationListWithContext(ctx context.Context, request *GetCustomizationListRequest) (response *GetCustomizationListResponse, err error) {
    if request == nil {
        request = NewGetCustomizationListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCustomizationList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCustomizationListResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomizationRequest() (request *ModifyCustomizationRequest) {
    request = &ModifyCustomizationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "ModifyCustomization")
    
    
    return
}

func NewModifyCustomizationResponse() (response *ModifyCustomizationResponse) {
    response = &ModifyCustomizationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCustomization
// 用户通过该接口可以更新自学习模型，如模型名称、模型类型、模型语料。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CHECKRESOURCERESPONSECODEERROR = "AuthFailure.CheckResourceResponseCodeError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ERRORDOWNFILE = "FailedOperation.ErrorDownFile"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FILEENCODE = "InvalidParameter.FileEncode"
//  INVALIDPARAMETER_MODELSTATE = "InvalidParameter.ModelState"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODELID = "InvalidParameterValue.ModelId"
func (c *Client) ModifyCustomization(request *ModifyCustomizationRequest) (response *ModifyCustomizationResponse, err error) {
    return c.ModifyCustomizationWithContext(context.Background(), request)
}

// ModifyCustomization
// 用户通过该接口可以更新自学习模型，如模型名称、模型类型、模型语料。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CHECKRESOURCERESPONSECODEERROR = "AuthFailure.CheckResourceResponseCodeError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_ERRORDOWNFILE = "FailedOperation.ErrorDownFile"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FILEENCODE = "InvalidParameter.FileEncode"
//  INVALIDPARAMETER_MODELSTATE = "InvalidParameter.ModelState"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODELID = "InvalidParameterValue.ModelId"
func (c *Client) ModifyCustomizationWithContext(ctx context.Context, request *ModifyCustomizationRequest) (response *ModifyCustomizationResponse, err error) {
    if request == nil {
        request = NewModifyCustomizationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomization require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomizationResponse()
    err = c.Send(request, response)
    return
}

func NewModifyCustomizationStateRequest() (request *ModifyCustomizationStateRequest) {
    request = &ModifyCustomizationStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "ModifyCustomizationState")
    
    
    return
}

func NewModifyCustomizationStateResponse() (response *ModifyCustomizationStateResponse) {
    response = &ModifyCustomizationStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ModifyCustomizationState
// 通过该接口，用户可以修改自学习模型状态，上下线自学习模型
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CHECKRESOURCERESPONSECODEERROR = "AuthFailure.CheckResourceResponseCodeError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODELID = "InvalidParameterValue.ModelId"
//  INVALIDPARAMETERVALUE_TOSTATE = "InvalidParameterValue.ToState"
//  LIMITEXCEEDED_ONLINEFULL = "LimitExceeded.OnlineFull"
func (c *Client) ModifyCustomizationState(request *ModifyCustomizationStateRequest) (response *ModifyCustomizationStateResponse, err error) {
    return c.ModifyCustomizationStateWithContext(context.Background(), request)
}

// ModifyCustomizationState
// 通过该接口，用户可以修改自学习模型状态，上下线自学习模型
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_CHECKRESOURCERESPONSECODEERROR = "AuthFailure.CheckResourceResponseCodeError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODELID = "InvalidParameterValue.ModelId"
//  INVALIDPARAMETERVALUE_TOSTATE = "InvalidParameterValue.ToState"
//  LIMITEXCEEDED_ONLINEFULL = "LimitExceeded.OnlineFull"
func (c *Client) ModifyCustomizationStateWithContext(ctx context.Context, request *ModifyCustomizationStateRequest) (response *ModifyCustomizationStateResponse, err error) {
    if request == nil {
        request = NewModifyCustomizationStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ModifyCustomizationState require credential")
    }

    request.SetContext(ctx)
    
    response = NewModifyCustomizationStateResponse()
    err = c.Send(request, response)
    return
}

func NewSentenceRecognitionRequest() (request *SentenceRecognitionRequest) {
    request = &SentenceRecognitionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "SentenceRecognition")
    
    
    return
}

func NewSentenceRecognitionResponse() (response *SentenceRecognitionResponse) {
    response = &SentenceRecognitionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SentenceRecognition
// 本接口用于对60秒之内的短音频文件进行识别。<br>•   支持中文普通话、英语、粤语、日语、越南语、马来语、上海话、四川话、武汉话、贵阳话、昆明话、西安话、郑州话、太原话、兰州话、银川话、西宁话、南京话、合肥话、南昌话、长沙话、苏州话、杭州话、济南话、天津话、石家庄话、黑龙江话、吉林话、辽宁话。<br>•   支持本地语音文件上传和语音URL上传两种请求方式，音频时长不能超过60s，音频文件大小不能超过3MB。<br>•   音频格式支持wav、pcm、ogg-opus、speex、silk、mp3、m4a、aac。<br>•   请求方法为 HTTP POST , Content-Type为"application/json; charset=utf-8"<br>•   签名方法参考 [公共参数](https://cloud.tencent.com/document/api/1093/35640) 中签名方法v3。<br>•   默认接口请求频率限制：30次/秒，如您有提高请求频率限制的需求，请[前往购买](https://buy.cloud.tencent.com/asr)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORRECOGNIZE = "FailedOperation.ErrorRecognize"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERHASNOFREEAMOUNT = "FailedOperation.UserHasNoFreeAmount"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORCONFIGURE = "InternalError.ErrorConfigure"
//  INTERNALERROR_ERRORCREATELOG = "InternalError.ErrorCreateLog"
//  INTERNALERROR_ERRORDOWNFILE = "InternalError.ErrorDownFile"
//  INTERNALERROR_ERRORFAILNEWPREQUEST = "InternalError.ErrorFailNewprequest"
//  INTERNALERROR_ERRORFAILWRITETODB = "InternalError.ErrorFailWritetodb"
//  INTERNALERROR_ERRORFILECANNOTOPEN = "InternalError.ErrorFileCannotopen"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_ERRORMAKELOGPATH = "InternalError.ErrorMakeLogpath"
//  INTERNALERROR_ERRORRECOGNIZE = "InternalError.ErrorRecognize"
//  INVALIDPARAMETER_ERRORCONTENTLENGTH = "InvalidParameter.ErrorContentlength"
//  INVALIDPARAMETER_ERRORPARAMSMISSING = "InvalidParameter.ErrorParamsMissing"
//  INVALIDPARAMETER_ERRORPARSEQUEST = "InvalidParameter.ErrorParsequest"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ERRORINVALIDAPPID = "InvalidParameterValue.ErrorInvalidAppid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDCLIENTIP = "InvalidParameterValue.ErrorInvalidClientip"
//  INVALIDPARAMETERVALUE_ERRORINVALIDENGSERVICE = "InvalidParameterValue.ErrorInvalidEngservice"
//  INVALIDPARAMETERVALUE_ERRORINVALIDPROJECTID = "InvalidParameterValue.ErrorInvalidProjectid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDREQUESTID = "InvalidParameterValue.ErrorInvalidRequestid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDSOURCETYPE = "InvalidParameterValue.ErrorInvalidSourcetype"
//  INVALIDPARAMETERVALUE_ERRORINVALIDSUBSERVICETYPE = "InvalidParameterValue.ErrorInvalidSubservicetype"
//  INVALIDPARAMETERVALUE_ERRORINVALIDURL = "InvalidParameterValue.ErrorInvalidUrl"
//  INVALIDPARAMETERVALUE_ERRORINVALIDUSERAUDIOKEY = "InvalidParameterValue.ErrorInvalidUseraudiokey"
//  INVALIDPARAMETERVALUE_ERRORINVALIDVOICEFORMAT = "InvalidParameterValue.ErrorInvalidVoiceFormat"
//  INVALIDPARAMETERVALUE_ERRORINVALIDVOICEDATA = "InvalidParameterValue.ErrorInvalidVoicedata"
//  INVALIDPARAMETERVALUE_ERRORVOICEDATATOOLONG = "InvalidParameterValue.ErrorVoicedataTooLong"
func (c *Client) SentenceRecognition(request *SentenceRecognitionRequest) (response *SentenceRecognitionResponse, err error) {
    return c.SentenceRecognitionWithContext(context.Background(), request)
}

// SentenceRecognition
// 本接口用于对60秒之内的短音频文件进行识别。<br>•   支持中文普通话、英语、粤语、日语、越南语、马来语、上海话、四川话、武汉话、贵阳话、昆明话、西安话、郑州话、太原话、兰州话、银川话、西宁话、南京话、合肥话、南昌话、长沙话、苏州话、杭州话、济南话、天津话、石家庄话、黑龙江话、吉林话、辽宁话。<br>•   支持本地语音文件上传和语音URL上传两种请求方式，音频时长不能超过60s，音频文件大小不能超过3MB。<br>•   音频格式支持wav、pcm、ogg-opus、speex、silk、mp3、m4a、aac。<br>•   请求方法为 HTTP POST , Content-Type为"application/json; charset=utf-8"<br>•   签名方法参考 [公共参数](https://cloud.tencent.com/document/api/1093/35640) 中签名方法v3。<br>•   默认接口请求频率限制：30次/秒，如您有提高请求频率限制的需求，请[前往购买](https://buy.cloud.tencent.com/asr)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORRECOGNIZE = "FailedOperation.ErrorRecognize"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERHASNOFREEAMOUNT = "FailedOperation.UserHasNoFreeAmount"
//  FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORCONFIGURE = "InternalError.ErrorConfigure"
//  INTERNALERROR_ERRORCREATELOG = "InternalError.ErrorCreateLog"
//  INTERNALERROR_ERRORDOWNFILE = "InternalError.ErrorDownFile"
//  INTERNALERROR_ERRORFAILNEWPREQUEST = "InternalError.ErrorFailNewprequest"
//  INTERNALERROR_ERRORFAILWRITETODB = "InternalError.ErrorFailWritetodb"
//  INTERNALERROR_ERRORFILECANNOTOPEN = "InternalError.ErrorFileCannotopen"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_ERRORMAKELOGPATH = "InternalError.ErrorMakeLogpath"
//  INTERNALERROR_ERRORRECOGNIZE = "InternalError.ErrorRecognize"
//  INVALIDPARAMETER_ERRORCONTENTLENGTH = "InvalidParameter.ErrorContentlength"
//  INVALIDPARAMETER_ERRORPARAMSMISSING = "InvalidParameter.ErrorParamsMissing"
//  INVALIDPARAMETER_ERRORPARSEQUEST = "InvalidParameter.ErrorParsequest"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ERRORINVALIDAPPID = "InvalidParameterValue.ErrorInvalidAppid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDCLIENTIP = "InvalidParameterValue.ErrorInvalidClientip"
//  INVALIDPARAMETERVALUE_ERRORINVALIDENGSERVICE = "InvalidParameterValue.ErrorInvalidEngservice"
//  INVALIDPARAMETERVALUE_ERRORINVALIDPROJECTID = "InvalidParameterValue.ErrorInvalidProjectid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDREQUESTID = "InvalidParameterValue.ErrorInvalidRequestid"
//  INVALIDPARAMETERVALUE_ERRORINVALIDSOURCETYPE = "InvalidParameterValue.ErrorInvalidSourcetype"
//  INVALIDPARAMETERVALUE_ERRORINVALIDSUBSERVICETYPE = "InvalidParameterValue.ErrorInvalidSubservicetype"
//  INVALIDPARAMETERVALUE_ERRORINVALIDURL = "InvalidParameterValue.ErrorInvalidUrl"
//  INVALIDPARAMETERVALUE_ERRORINVALIDUSERAUDIOKEY = "InvalidParameterValue.ErrorInvalidUseraudiokey"
//  INVALIDPARAMETERVALUE_ERRORINVALIDVOICEFORMAT = "InvalidParameterValue.ErrorInvalidVoiceFormat"
//  INVALIDPARAMETERVALUE_ERRORINVALIDVOICEDATA = "InvalidParameterValue.ErrorInvalidVoicedata"
//  INVALIDPARAMETERVALUE_ERRORVOICEDATATOOLONG = "InvalidParameterValue.ErrorVoicedataTooLong"
func (c *Client) SentenceRecognitionWithContext(ctx context.Context, request *SentenceRecognitionRequest) (response *SentenceRecognitionResponse, err error) {
    if request == nil {
        request = NewSentenceRecognitionRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SentenceRecognition require credential")
    }

    request.SetContext(ctx)
    
    response = NewSentenceRecognitionResponse()
    err = c.Send(request, response)
    return
}

func NewSetVocabStateRequest() (request *SetVocabStateRequest) {
    request = &SetVocabStateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "SetVocabState")
    
    
    return
}

func NewSetVocabStateResponse() (response *SetVocabStateResponse) {
    response = &SetVocabStateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SetVocabState
// 用户通过该接口可以设置热词表的默认状态。初始状态为0，用户可设置状态为1，即为默认状态。默认状态表示用户在请求识别时，如不设置热词表ID，则默认使用状态为1的热词表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_FAILACCESSREDIS = "InternalError.FailAccessRedis"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDVOCABID = "InvalidParameterValue.InvalidVocabId"
//  INVALIDPARAMETERVALUE_INVALIDVOCABSTATE = "InvalidParameterValue.InvalidVocabState"
func (c *Client) SetVocabState(request *SetVocabStateRequest) (response *SetVocabStateResponse, err error) {
    return c.SetVocabStateWithContext(context.Background(), request)
}

// SetVocabState
// 用户通过该接口可以设置热词表的默认状态。初始状态为0，用户可设置状态为1，即为默认状态。默认状态表示用户在请求识别时，如不设置热词表ID，则默认使用状态为1的热词表。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_FAILACCESSREDIS = "InternalError.FailAccessRedis"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDVOCABID = "InvalidParameterValue.InvalidVocabId"
//  INVALIDPARAMETERVALUE_INVALIDVOCABSTATE = "InvalidParameterValue.InvalidVocabState"
func (c *Client) SetVocabStateWithContext(ctx context.Context, request *SetVocabStateRequest) (response *SetVocabStateResponse, err error) {
    if request == nil {
        request = NewSetVocabStateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetVocabState require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetVocabStateResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAsrVocabRequest() (request *UpdateAsrVocabRequest) {
    request = &UpdateAsrVocabRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "UpdateAsrVocab")
    
    
    return
}

func NewUpdateAsrVocabResponse() (response *UpdateAsrVocabResponse) {
    response = &UpdateAsrVocabResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateAsrVocab
// 用户通过本接口进行对应的词表信息更新。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERLENGTH = "InvalidParameterValue.InvalidParameterLength"
//  INVALIDPARAMETERVALUE_INVALIDVOCABID = "InvalidParameterValue.InvalidVocabId"
//  INVALIDPARAMETERVALUE_INVALIDWORDWEIGHT = "InvalidParameterValue.InvalidWordWeight"
//  INVALIDPARAMETERVALUE_INVALIDWORDWEIGHTSTR = "InvalidParameterValue.InvalidWordWeightStr"
//  LIMITEXCEEDED_VOCABFULL = "LimitExceeded.VocabFull"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) UpdateAsrVocab(request *UpdateAsrVocabRequest) (response *UpdateAsrVocabResponse, err error) {
    return c.UpdateAsrVocabWithContext(context.Background(), request)
}

// UpdateAsrVocab
// 用户通过本接口进行对应的词表信息更新。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERLENGTH = "InvalidParameterValue.InvalidParameterLength"
//  INVALIDPARAMETERVALUE_INVALIDVOCABID = "InvalidParameterValue.InvalidVocabId"
//  INVALIDPARAMETERVALUE_INVALIDWORDWEIGHT = "InvalidParameterValue.InvalidWordWeight"
//  INVALIDPARAMETERVALUE_INVALIDWORDWEIGHTSTR = "InvalidParameterValue.InvalidWordWeightStr"
//  LIMITEXCEEDED_VOCABFULL = "LimitExceeded.VocabFull"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) UpdateAsrVocabWithContext(ctx context.Context, request *UpdateAsrVocabRequest) (response *UpdateAsrVocabResponse, err error) {
    if request == nil {
        request = NewUpdateAsrVocabRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAsrVocab require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAsrVocabResponse()
    err = c.Send(request, response)
    return
}
