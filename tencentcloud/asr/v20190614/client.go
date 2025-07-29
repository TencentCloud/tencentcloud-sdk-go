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
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "CloseAsyncRecognitionTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CloseAsyncRecognitionTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCloseAsyncRecognitionTaskResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAsrKeyWordLibRequest() (request *CreateAsrKeyWordLibRequest) {
    request = &CreateAsrKeyWordLibRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "CreateAsrKeyWordLib")
    
    
    return
}

func NewCreateAsrKeyWordLibResponse() (response *CreateAsrKeyWordLibResponse) {
    response = &CreateAsrKeyWordLibResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAsrKeyWordLib
// 用户通过本接口进行关键字词表的创建。
//
// <br>•   默认每个用户最多可创建30个关键字词表。
//
// <br>•   每个关键词词表最多可添加100个词，每个词最多5个汉字或15个字符。
//
// <br>•   词表通过本地文件形式上传。
//
// <br>•   本地文件必须为UTF-8编码格式，每行仅添加一个词且不能包含标点和特殊字符。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORCONTENTLENGTH = "InvalidParameter.ErrorContentlength"
//  INVALIDPARAMETER_ERRORPARAMSMISSING = "InvalidParameter.ErrorParamsMissing"
//  INVALIDPARAMETER_KEYWORDLIBNAMEEXIST = "InvalidParameter.KeyWordLibNameExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ERRORINVALIDAPPID = "InvalidParameterValue.ErrorInvalidAppid"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERLENGTH = "InvalidParameterValue.InvalidParameterLength"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateAsrKeyWordLib(request *CreateAsrKeyWordLibRequest) (response *CreateAsrKeyWordLibResponse, err error) {
    return c.CreateAsrKeyWordLibWithContext(context.Background(), request)
}

// CreateAsrKeyWordLib
// 用户通过本接口进行关键字词表的创建。
//
// <br>•   默认每个用户最多可创建30个关键字词表。
//
// <br>•   每个关键词词表最多可添加100个词，每个词最多5个汉字或15个字符。
//
// <br>•   词表通过本地文件形式上传。
//
// <br>•   本地文件必须为UTF-8编码格式，每行仅添加一个词且不能包含标点和特殊字符。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORCONTENTLENGTH = "InvalidParameter.ErrorContentlength"
//  INVALIDPARAMETER_ERRORPARAMSMISSING = "InvalidParameter.ErrorParamsMissing"
//  INVALIDPARAMETER_KEYWORDLIBNAMEEXIST = "InvalidParameter.KeyWordLibNameExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ERRORINVALIDAPPID = "InvalidParameterValue.ErrorInvalidAppid"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERLENGTH = "InvalidParameterValue.InvalidParameterLength"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) CreateAsrKeyWordLibWithContext(ctx context.Context, request *CreateAsrKeyWordLibRequest) (response *CreateAsrKeyWordLibResponse, err error) {
    if request == nil {
        request = NewCreateAsrKeyWordLibRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "CreateAsrKeyWordLib")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAsrKeyWordLib require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAsrKeyWordLibResponse()
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
// <br>•   每个热词表最多可添加1000个词，每个词最长10个汉字或30个英文字符，不能超出限制。
//
// <br>•   热词表可以通过数组或者本地文件形式上传。
//
// <br>•   本地文件必须为UTF-8编码格式，每行仅添加一个热词且不能包含标点和特殊字符。
//
// <br>•   热词权重取值范围为[1,11]之间的整数或者100，权重越大代表该词被识别出来的概率越大。
//
// <br>• 注意:  热词权重设置为11时，当前热词将升级为超级热词，建议仅将重要且必须生效的热词设置到11，设置过多权重为11的热词将影响整体字准率。
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
// <br>•   每个热词表最多可添加1000个词，每个词最长10个汉字或30个英文字符，不能超出限制。
//
// <br>•   热词表可以通过数组或者本地文件形式上传。
//
// <br>•   本地文件必须为UTF-8编码格式，每行仅添加一个热词且不能包含标点和特殊字符。
//
// <br>•   热词权重取值范围为[1,11]之间的整数或者100，权重越大代表该词被识别出来的概率越大。
//
// <br>• 注意:  热词权重设置为11时，当前热词将升级为超级热词，建议仅将重要且必须生效的热词设置到11，设置过多权重为11的热词将影响整体字准率。
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
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "CreateAsrVocab")
    
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
//  FAILEDOPERATION_USERHASNOAMOUNT = "FailedOperation.UserHasNoAmount"
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
//  FAILEDOPERATION_USERHASNOAMOUNT = "FailedOperation.UserHasNoAmount"
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
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "CreateAsyncRecognitionTask")
    
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
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "CreateCustomization")
    
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
// 本接口可对较长的录音文件进行识别。如希望直接使用带界面的语音识别产品，请访问[产品体验中心](https://console.cloud.tencent.com/asr/demonstrate)。产品计费标准请查阅 [计费概述（在线版）](https://cloud.tencent.com/document/product/1093/35686)
//
// • 接口默认限频：20次/秒。此处仅限制任务提交频次，与识别结果返回时效无关
//
// • 返回时效：异步回调，非实时返回。最长3小时返回识别结果，**大多数情况下，1小时的音频1-3分钟即可完成识别**。请注意：上述返回时长不含音频下载时延，且30分钟内发送超过1000小时录音或2万条任务的情况除外
//
// • 音频格式：wav、mp3、m4a、flv、mp4、wma、3gp、amr、aac、ogg-opus、flac
//
// • 支持语言：在本页面上搜索 **EngineModelType**，或前往 [产品功能](https://cloud.tencent.com/document/product/1093/35682) 查看
//
// • 音频提交方式：本接口支持**音频 URL 、本地音频文件**两种请求方式。推荐使用 [腾讯云COS](https://cloud.tencent.com/document/product/436/38484) 来存储、生成URL并提交任务，此种方式将不产生外网和流量下行费用，可节约成本、提升任务速度（可参考COS预签名指南：[使用预签名 URL 访问 COS](https://cloud.tencent.com/document/product/436/68284) ，获取COS预签名url）
//
// • 音频限制：音频 URL 时长不能大于5小时，文件大小不超过1GB；本地音频文件不能大于5MB
//
// • 如何获取识别结果：支持**回调或轮询**的方式获取结果，具体请参考 [录音文件识别结果查询](https://cloud.tencent.com/document/product/1093/37822)
//
// • 识别结果有效时间：识别结果在服务端保存24小时
//
// • 签名方法参考 [公共参数](https://cloud.tencent.com/document/api/1093/35640) 中签名方法 v3
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CHECKAUTHINFOFAILED = "FailedOperation.CheckAuthInfoFailed"
//  FAILEDOPERATION_ERRORDOWNFILE = "FailedOperation.ErrorDownFile"
//  FAILEDOPERATION_ERRORRECOGNIZE = "FailedOperation.ErrorRecognize"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERHASNOAMOUNT = "FailedOperation.UserHasNoAmount"
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
// 本接口可对较长的录音文件进行识别。如希望直接使用带界面的语音识别产品，请访问[产品体验中心](https://console.cloud.tencent.com/asr/demonstrate)。产品计费标准请查阅 [计费概述（在线版）](https://cloud.tencent.com/document/product/1093/35686)
//
// • 接口默认限频：20次/秒。此处仅限制任务提交频次，与识别结果返回时效无关
//
// • 返回时效：异步回调，非实时返回。最长3小时返回识别结果，**大多数情况下，1小时的音频1-3分钟即可完成识别**。请注意：上述返回时长不含音频下载时延，且30分钟内发送超过1000小时录音或2万条任务的情况除外
//
// • 音频格式：wav、mp3、m4a、flv、mp4、wma、3gp、amr、aac、ogg-opus、flac
//
// • 支持语言：在本页面上搜索 **EngineModelType**，或前往 [产品功能](https://cloud.tencent.com/document/product/1093/35682) 查看
//
// • 音频提交方式：本接口支持**音频 URL 、本地音频文件**两种请求方式。推荐使用 [腾讯云COS](https://cloud.tencent.com/document/product/436/38484) 来存储、生成URL并提交任务，此种方式将不产生外网和流量下行费用，可节约成本、提升任务速度（可参考COS预签名指南：[使用预签名 URL 访问 COS](https://cloud.tencent.com/document/product/436/68284) ，获取COS预签名url）
//
// • 音频限制：音频 URL 时长不能大于5小时，文件大小不超过1GB；本地音频文件不能大于5MB
//
// • 如何获取识别结果：支持**回调或轮询**的方式获取结果，具体请参考 [录音文件识别结果查询](https://cloud.tencent.com/document/product/1093/37822)
//
// • 识别结果有效时间：识别结果在服务端保存24小时
//
// • 签名方法参考 [公共参数](https://cloud.tencent.com/document/api/1093/35640) 中签名方法 v3
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_CHECKAUTHINFOFAILED = "FailedOperation.CheckAuthInfoFailed"
//  FAILEDOPERATION_ERRORDOWNFILE = "FailedOperation.ErrorDownFile"
//  FAILEDOPERATION_ERRORRECOGNIZE = "FailedOperation.ErrorRecognize"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERHASNOAMOUNT = "FailedOperation.UserHasNoAmount"
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
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "CreateRecTask")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateRecTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateRecTaskResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAsrKeyWordLibRequest() (request *DeleteAsrKeyWordLibRequest) {
    request = &DeleteAsrKeyWordLibRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "DeleteAsrKeyWordLib")
    
    
    return
}

func NewDeleteAsrKeyWordLibResponse() (response *DeleteAsrKeyWordLibResponse) {
    response = &DeleteAsrKeyWordLibResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAsrKeyWordLib
// 用户通过本接口进行关键词表的删除。
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
func (c *Client) DeleteAsrKeyWordLib(request *DeleteAsrKeyWordLibRequest) (response *DeleteAsrKeyWordLibResponse, err error) {
    return c.DeleteAsrKeyWordLibWithContext(context.Background(), request)
}

// DeleteAsrKeyWordLib
// 用户通过本接口进行关键词表的删除。
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
func (c *Client) DeleteAsrKeyWordLibWithContext(ctx context.Context, request *DeleteAsrKeyWordLibRequest) (response *DeleteAsrKeyWordLibResponse, err error) {
    if request == nil {
        request = NewDeleteAsrKeyWordLibRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "DeleteAsrKeyWordLib")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAsrKeyWordLib require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAsrKeyWordLibResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "DeleteAsrVocab")
    
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
//  INTERNALERROR_TAGREQUESTERROR = "InternalError.TagRequestError"
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
//  INTERNALERROR_TAGREQUESTERROR = "InternalError.TagRequestError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_MODELSTATE = "InvalidParameter.ModelState"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_MODELID = "InvalidParameterValue.ModelId"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DeleteCustomizationWithContext(ctx context.Context, request *DeleteCustomizationRequest) (response *DeleteCustomizationResponse, err error) {
    if request == nil {
        request = NewDeleteCustomizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "DeleteCustomization")
    
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
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "DescribeAsyncRecognitionTasks")
    
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
// 调用录音文件识别请求接口后，有回调和轮询两种方式获取识别结果。
//
// • **注意任务有效期为24小时，超过24小时的任务请不要再查询，且不要依赖TaskId作为业务唯一ID，不同日期可能出现重复TaskId。**
//
// • 当采用回调方式时，识别完成后会将结果通过 POST 请求的形式通知到用户在请求时填写的回调 URL，具体请参见[ 录音识别结果回调 ](https://cloud.tencent.com/document/product/1093/52632)。
//
// • 当采用轮询方式时，需要主动提交任务ID来轮询识别结果，共有任务成功、等待、执行中和失败四种结果，具体信息请参见下文说明。
//
// •   请求方法为 HTTP POST , Content-Type为"application/json; charset=utf-8"
//
// •   签名方法参考 [公共参数](https://cloud.tencent.com/document/api/1093/35640) 中签名方法v3。
//
// •   默认接口请求频率限制：50次/秒，如您有提高请求频率限制的需求，请提[工单](https://console.cloud.tencent.com/workorder/category)进行咨询。
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
// 调用录音文件识别请求接口后，有回调和轮询两种方式获取识别结果。
//
// • **注意任务有效期为24小时，超过24小时的任务请不要再查询，且不要依赖TaskId作为业务唯一ID，不同日期可能出现重复TaskId。**
//
// • 当采用回调方式时，识别完成后会将结果通过 POST 请求的形式通知到用户在请求时填写的回调 URL，具体请参见[ 录音识别结果回调 ](https://cloud.tencent.com/document/product/1093/52632)。
//
// • 当采用轮询方式时，需要主动提交任务ID来轮询识别结果，共有任务成功、等待、执行中和失败四种结果，具体信息请参见下文说明。
//
// •   请求方法为 HTTP POST , Content-Type为"application/json; charset=utf-8"
//
// •   签名方法参考 [公共参数](https://cloud.tencent.com/document/api/1093/35640) 中签名方法v3。
//
// •   默认接口请求频率限制：50次/秒，如您有提高请求频率限制的需求，请提[工单](https://console.cloud.tencent.com/workorder/category)进行咨询。
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
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "DescribeTaskStatus")
    
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
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "DownloadAsrVocab")
    
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
//  FAILEDOPERATION_ERRORDOWNFILE = "FailedOperation.ErrorDownFile"
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
//  FAILEDOPERATION_ERRORDOWNFILE = "FailedOperation.ErrorDownFile"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_MODELID = "InvalidParameterValue.ModelId"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) DownloadCustomizationWithContext(ctx context.Context, request *DownloadCustomizationRequest) (response *DownloadCustomizationResponse, err error) {
    if request == nil {
        request = NewDownloadCustomizationRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "DownloadCustomization")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DownloadCustomization require credential")
    }

    request.SetContext(ctx)
    
    response = NewDownloadCustomizationResponse()
    err = c.Send(request, response)
    return
}

func NewGetAsrKeyWordLibListRequest() (request *GetAsrKeyWordLibListRequest) {
    request = &GetAsrKeyWordLibListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "GetAsrKeyWordLibList")
    
    
    return
}

func NewGetAsrKeyWordLibListResponse() (response *GetAsrKeyWordLibListResponse) {
    response = &GetAsrKeyWordLibListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetAsrKeyWordLibList
// 用户通过该接口，可获得所有的关键词表及其信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_CHECKRESOURCERESPONSECODEERROR = "AuthFailure.CheckResourceResponseCodeError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetAsrKeyWordLibList(request *GetAsrKeyWordLibListRequest) (response *GetAsrKeyWordLibListResponse, err error) {
    return c.GetAsrKeyWordLibListWithContext(context.Background(), request)
}

// GetAsrKeyWordLibList
// 用户通过该接口，可获得所有的关键词表及其信息。
//
// 可能返回的错误码:
//  AUTHFAILURE_CHECKRESOURCERESPONSECODEERROR = "AuthFailure.CheckResourceResponseCodeError"
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetAsrKeyWordLibListWithContext(ctx context.Context, request *GetAsrKeyWordLibListRequest) (response *GetAsrKeyWordLibListResponse, err error) {
    if request == nil {
        request = NewGetAsrKeyWordLibListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "GetAsrKeyWordLibList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetAsrKeyWordLibList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetAsrKeyWordLibListResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "GetAsrVocab")
    
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
//  INTERNALERROR_TAGREQUESTERROR = "InternalError.TagRequestError"
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
//  INTERNALERROR_TAGREQUESTERROR = "InternalError.TagRequestError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) GetAsrVocabListWithContext(ctx context.Context, request *GetAsrVocabListRequest) (response *GetAsrVocabListResponse, err error) {
    if request == nil {
        request = NewGetAsrVocabListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "GetAsrVocabList")
    
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
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "GetCustomizationList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetCustomizationList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetCustomizationListResponse()
    err = c.Send(request, response)
    return
}

func NewGetModelInfoRequest() (request *GetModelInfoRequest) {
    request = &GetModelInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "GetModelInfo")
    
    
    return
}

func NewGetModelInfoResponse() (response *GetModelInfoResponse) {
    response = &GetModelInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetModelInfo
// 通过自学习模型id获取自学习模型详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_MODELID = "InvalidParameterValue.ModelId"
func (c *Client) GetModelInfo(request *GetModelInfoRequest) (response *GetModelInfoResponse, err error) {
    return c.GetModelInfoWithContext(context.Background(), request)
}

// GetModelInfo
// 通过自学习模型id获取自学习模型详细信息
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_MODELID = "InvalidParameterValue.ModelId"
func (c *Client) GetModelInfoWithContext(ctx context.Context, request *GetModelInfoRequest) (response *GetModelInfoResponse, err error) {
    if request == nil {
        request = NewGetModelInfoRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "GetModelInfo")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetModelInfo require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetModelInfoResponse()
    err = c.Send(request, response)
    return
}

func NewGetUsageByDateRequest() (request *GetUsageByDateRequest) {
    request = &GetUsageByDateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "GetUsageByDate")
    
    
    return
}

func NewGetUsageByDateResponse() (response *GetUsageByDateResponse) {
    response = &GetUsageByDateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GetUsageByDate
// 查询用户用量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) GetUsageByDate(request *GetUsageByDateRequest) (response *GetUsageByDateResponse, err error) {
    return c.GetUsageByDateWithContext(context.Background(), request)
}

// GetUsageByDate
// 查询用户用量
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  LIMITEXCEEDED = "LimitExceeded"
func (c *Client) GetUsageByDateWithContext(ctx context.Context, request *GetUsageByDateRequest) (response *GetUsageByDateResponse, err error) {
    if request == nil {
        request = NewGetUsageByDateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "GetUsageByDate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetUsageByDate require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetUsageByDateResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "ModifyCustomization")
    
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
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "ModifyCustomizationState")
    
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
// 本接口用于对60秒之内的短音频文件进行识别。
//
// •   支持中文普通话、英语、粤语、日语、越南语、马来语、印度尼西亚语、菲律宾语、泰语、葡萄牙语、土耳其语、阿拉伯语、印地语、法语、德语、上海话、四川话、武汉话、贵阳话、昆明话、西安话、郑州话、太原话、兰州话、银川话、西宁话、南京话、合肥话、南昌话、长沙话、苏州话、杭州话、济南话、天津话、石家庄话、黑龙江话、吉林话、辽宁话。
//
// •   支持本地语音文件上传和语音URL上传两种请求方式，音频时长不能超过60s，音频文件大小不能超过3MB。推荐使用 [腾讯云COS](https://cloud.tencent.com/document/product/436/38484) 来存储音频、生成URL并提交请求，此种方式会走内网下载音频，极大降低整体请求时延；并且不会产生外网和流量下行费用，可节约成本（可参考COS预签名指南：[使用预签名 URL 访问 COS](https://cloud.tencent.com/document/product/436/68284) ，获取COS预签名url）
//
// •   音频格式支持wav、pcm、ogg-opus、speex、silk、mp3、m4a、aac、 amr。
//
// •   请求方法为 HTTP POST , Content-Type为"application/json; charset=utf-8"
//
// •   签名方法参考 [公共参数](https://cloud.tencent.com/document/api/1093/35640) 中签名方法v3。
//
// •   默认接口请求频率限制：30次/秒，如您有提高请求频率限制的需求，请[前往购买](https://buy.cloud.tencent.com/asr)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORRECOGNIZE = "FailedOperation.ErrorRecognize"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERHASNOAMOUNT = "FailedOperation.UserHasNoAmount"
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
// 本接口用于对60秒之内的短音频文件进行识别。
//
// •   支持中文普通话、英语、粤语、日语、越南语、马来语、印度尼西亚语、菲律宾语、泰语、葡萄牙语、土耳其语、阿拉伯语、印地语、法语、德语、上海话、四川话、武汉话、贵阳话、昆明话、西安话、郑州话、太原话、兰州话、银川话、西宁话、南京话、合肥话、南昌话、长沙话、苏州话、杭州话、济南话、天津话、石家庄话、黑龙江话、吉林话、辽宁话。
//
// •   支持本地语音文件上传和语音URL上传两种请求方式，音频时长不能超过60s，音频文件大小不能超过3MB。推荐使用 [腾讯云COS](https://cloud.tencent.com/document/product/436/38484) 来存储音频、生成URL并提交请求，此种方式会走内网下载音频，极大降低整体请求时延；并且不会产生外网和流量下行费用，可节约成本（可参考COS预签名指南：[使用预签名 URL 访问 COS](https://cloud.tencent.com/document/product/436/68284) ，获取COS预签名url）
//
// •   音频格式支持wav、pcm、ogg-opus、speex、silk、mp3、m4a、aac、 amr。
//
// •   请求方法为 HTTP POST , Content-Type为"application/json; charset=utf-8"
//
// •   签名方法参考 [公共参数](https://cloud.tencent.com/document/api/1093/35640) 中签名方法v3。
//
// •   默认接口请求频率限制：30次/秒，如您有提高请求频率限制的需求，请[前往购买](https://buy.cloud.tencent.com/asr)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ERRORRECOGNIZE = "FailedOperation.ErrorRecognize"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  FAILEDOPERATION_USERHASNOAMOUNT = "FailedOperation.UserHasNoAmount"
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
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "SentenceRecognition")
    
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
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "SetVocabState")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SetVocabState require credential")
    }

    request.SetContext(ctx)
    
    response = NewSetVocabStateResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateAsrKeyWordLibRequest() (request *UpdateAsrKeyWordLibRequest) {
    request = &UpdateAsrKeyWordLibRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "UpdateAsrKeyWordLib")
    
    
    return
}

func NewUpdateAsrKeyWordLibResponse() (response *UpdateAsrKeyWordLibResponse) {
    response = &UpdateAsrKeyWordLibResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UpdateAsrKeyWordLib
// 用户通过本接口进行对应的关键词表信息更新。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_KEYWORDLIBNAMEEXIST = "InvalidParameter.KeyWordLibNameExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERLENGTH = "InvalidParameterValue.InvalidParameterLength"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) UpdateAsrKeyWordLib(request *UpdateAsrKeyWordLibRequest) (response *UpdateAsrKeyWordLibResponse, err error) {
    return c.UpdateAsrKeyWordLibWithContext(context.Background(), request)
}

// UpdateAsrKeyWordLib
// 用户通过本接口进行对应的关键词表信息更新。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_KEYWORDLIBNAMEEXIST = "InvalidParameter.KeyWordLibNameExist"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERLENGTH = "InvalidParameterValue.InvalidParameterLength"
//  MISSINGPARAMETER = "MissingParameter"
//  UNKNOWNPARAMETER = "UnknownParameter"
func (c *Client) UpdateAsrKeyWordLibWithContext(ctx context.Context, request *UpdateAsrKeyWordLibRequest) (response *UpdateAsrKeyWordLibResponse, err error) {
    if request == nil {
        request = NewUpdateAsrKeyWordLibRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "UpdateAsrKeyWordLib")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAsrKeyWordLib require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAsrKeyWordLibResponse()
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
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "UpdateAsrVocab")
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateAsrVocab require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateAsrVocabResponse()
    err = c.Send(request, response)
    return
}

func NewVoicePrintCompareRequest() (request *VoicePrintCompareRequest) {
    request = &VoicePrintCompareRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "VoicePrintCompare")
    
    
    return
}

func NewVoicePrintCompareResponse() (response *VoicePrintCompareResponse) {
    response = &VoicePrintCompareResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VoicePrintCompare
// 通过比对两段音频内说话人的声纹，得到一个打分，可通过打分判断两段音频声纹相似度,  打分区间[0 - 100]。 音频要求：16k采样率， 16bit位深，pcm或者wav格式， 单声道，总时长不超过30秒的音频，base64编码数据大小不超过2M，音频内容只有一个说话人声音，并且尽可能清晰，这样结果更加准确。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_VOICEPRINTAUDIOFAILED = "InternalError.VoicePrintAudioFailed"
//  INTERNALERROR_VOICEPRINTDECODEFAILED = "InternalError.VoicePrintDecodeFailed"
//  INTERNALERROR_VOICEPRINTVERIFYFAILED = "InternalError.VoicePrintVerifyFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ERRORINVALIDVOICEDATA = "InvalidParameterValue.ErrorInvalidVoicedata"
//  INVALIDPARAMETERVALUE_NOHUMANVOICE = "InvalidParameterValue.NoHumanVoice"
func (c *Client) VoicePrintCompare(request *VoicePrintCompareRequest) (response *VoicePrintCompareResponse, err error) {
    return c.VoicePrintCompareWithContext(context.Background(), request)
}

// VoicePrintCompare
// 通过比对两段音频内说话人的声纹，得到一个打分，可通过打分判断两段音频声纹相似度,  打分区间[0 - 100]。 音频要求：16k采样率， 16bit位深，pcm或者wav格式， 单声道，总时长不超过30秒的音频，base64编码数据大小不超过2M，音频内容只有一个说话人声音，并且尽可能清晰，这样结果更加准确。
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_VOICEPRINTAUDIOFAILED = "InternalError.VoicePrintAudioFailed"
//  INTERNALERROR_VOICEPRINTDECODEFAILED = "InternalError.VoicePrintDecodeFailed"
//  INTERNALERROR_VOICEPRINTVERIFYFAILED = "InternalError.VoicePrintVerifyFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ERRORINVALIDVOICEDATA = "InvalidParameterValue.ErrorInvalidVoicedata"
//  INVALIDPARAMETERVALUE_NOHUMANVOICE = "InvalidParameterValue.NoHumanVoice"
func (c *Client) VoicePrintCompareWithContext(ctx context.Context, request *VoicePrintCompareRequest) (response *VoicePrintCompareResponse, err error) {
    if request == nil {
        request = NewVoicePrintCompareRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "VoicePrintCompare")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VoicePrintCompare require credential")
    }

    request.SetContext(ctx)
    
    response = NewVoicePrintCompareResponse()
    err = c.Send(request, response)
    return
}

func NewVoicePrintCountRequest() (request *VoicePrintCountRequest) {
    request = &VoicePrintCountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "VoicePrintCount")
    
    
    return
}

func NewVoicePrintCountResponse() (response *VoicePrintCountResponse) {
    response = &VoicePrintCountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VoicePrintCount
// 统计并返回注册的说话人id总数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) VoicePrintCount(request *VoicePrintCountRequest) (response *VoicePrintCountResponse, err error) {
    return c.VoicePrintCountWithContext(context.Background(), request)
}

// VoicePrintCount
// 统计并返回注册的说话人id总数
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
func (c *Client) VoicePrintCountWithContext(ctx context.Context, request *VoicePrintCountRequest) (response *VoicePrintCountResponse, err error) {
    if request == nil {
        request = NewVoicePrintCountRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "VoicePrintCount")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VoicePrintCount require credential")
    }

    request.SetContext(ctx)
    
    response = NewVoicePrintCountResponse()
    err = c.Send(request, response)
    return
}

func NewVoicePrintDeleteRequest() (request *VoicePrintDeleteRequest) {
    request = &VoicePrintDeleteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "VoicePrintDelete")
    
    
    return
}

func NewVoicePrintDeleteResponse() (response *VoicePrintDeleteResponse) {
    response = &VoicePrintDeleteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VoicePrintDelete
// 本接口用于以删除已经注册的说话人信息（删除之后，原有的说话人ID和说话人音频数据都会失效）
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXISTENTVOICEPRINTID = "FailedOperation.NotExistentVoicePrintId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_ERRORRECOGNIZE = "InternalError.ErrorRecognize"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_VOICEPRINTAUDIOFAILED = "InternalError.VoicePrintAudioFailed"
//  INTERNALERROR_VOICEPRINTDECODEFAILED = "InternalError.VoicePrintDecodeFailed"
//  INTERNALERROR_VOICEPRINTENROLLFAILED = "InternalError.VoicePrintEnrollFailed"
//  INTERNALERROR_VOICEPRINTVERIFYFAILED = "InternalError.VoicePrintVerifyFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) VoicePrintDelete(request *VoicePrintDeleteRequest) (response *VoicePrintDeleteResponse, err error) {
    return c.VoicePrintDeleteWithContext(context.Background(), request)
}

// VoicePrintDelete
// 本接口用于以删除已经注册的说话人信息（删除之后，原有的说话人ID和说话人音频数据都会失效）
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXISTENTVOICEPRINTID = "FailedOperation.NotExistentVoicePrintId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_ERRORRECOGNIZE = "InternalError.ErrorRecognize"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_VOICEPRINTAUDIOFAILED = "InternalError.VoicePrintAudioFailed"
//  INTERNALERROR_VOICEPRINTDECODEFAILED = "InternalError.VoicePrintDecodeFailed"
//  INTERNALERROR_VOICEPRINTENROLLFAILED = "InternalError.VoicePrintEnrollFailed"
//  INTERNALERROR_VOICEPRINTVERIFYFAILED = "InternalError.VoicePrintVerifyFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) VoicePrintDeleteWithContext(ctx context.Context, request *VoicePrintDeleteRequest) (response *VoicePrintDeleteResponse, err error) {
    if request == nil {
        request = NewVoicePrintDeleteRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "VoicePrintDelete")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VoicePrintDelete require credential")
    }

    request.SetContext(ctx)
    
    response = NewVoicePrintDeleteResponse()
    err = c.Send(request, response)
    return
}

func NewVoicePrintEnrollRequest() (request *VoicePrintEnrollRequest) {
    request = &VoicePrintEnrollRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "VoicePrintEnroll")
    
    
    return
}

func NewVoicePrintEnrollResponse() (response *VoicePrintEnrollResponse) {
    response = &VoicePrintEnrollResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VoicePrintEnroll
// 说话人注册接口用于注册一个指定音频，生成一个唯一的说话人id，后续可通过说话人验证接口验证其它音频和已有的说话人ID匹配度，注册时可指定说话人昵称，方便标识说话人ID，  说话人昵称可重复配置。 
//
// （注: 一个appid最多可以注册1000个说话人ID，一个说话人ID仅支持一条音频注册，后续可通过更新接口进行更新）
//
// 
//
// 使用须知
//
// 支持的输入格式：编码文件(PCM, WAV)、16 bit采样位数、单声道（mono）。
//
// 
//
// 支持的音频采样率：16000 Hz。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXISTENTVOICEPRINTID = "FailedOperation.NotExistentVoicePrintId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORRECOGNIZE = "InternalError.ErrorRecognize"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_FAILEDVOICEPRINTDECODE = "InternalError.FailedVoicePrintDecode"
//  INTERNALERROR_FAILEDVOICEPRINTENROLL = "InternalError.FailedVoicePrintEnroll"
//  INTERNALERROR_VOICEPRINTAUDIOFAILED = "InternalError.VoicePrintAudioFailed"
//  INTERNALERROR_VOICEPRINTDECODEFAILED = "InternalError.VoicePrintDecodeFailed"
//  INTERNALERROR_VOICEPRINTENROLLFAILED = "InternalError.VoicePrintEnrollFailed"
//  INTERNALERROR_VOICEPRINTVERIFYFAILED = "InternalError.VoicePrintVerifyFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ERRORVOICEDATATOOLONG = "InvalidParameterValue.ErrorVoicedataTooLong"
//  INVALIDPARAMETERVALUE_NOHUMANVOICE = "InvalidParameterValue.NoHumanVoice"
//  LIMITEXCEEDED_VOICEPRINTFULL = "LimitExceeded.VoicePrintFull"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) VoicePrintEnroll(request *VoicePrintEnrollRequest) (response *VoicePrintEnrollResponse, err error) {
    return c.VoicePrintEnrollWithContext(context.Background(), request)
}

// VoicePrintEnroll
// 说话人注册接口用于注册一个指定音频，生成一个唯一的说话人id，后续可通过说话人验证接口验证其它音频和已有的说话人ID匹配度，注册时可指定说话人昵称，方便标识说话人ID，  说话人昵称可重复配置。 
//
// （注: 一个appid最多可以注册1000个说话人ID，一个说话人ID仅支持一条音频注册，后续可通过更新接口进行更新）
//
// 
//
// 使用须知
//
// 支持的输入格式：编码文件(PCM, WAV)、16 bit采样位数、单声道（mono）。
//
// 
//
// 支持的音频采样率：16000 Hz。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXISTENTVOICEPRINTID = "FailedOperation.NotExistentVoicePrintId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORRECOGNIZE = "InternalError.ErrorRecognize"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_FAILEDVOICEPRINTDECODE = "InternalError.FailedVoicePrintDecode"
//  INTERNALERROR_FAILEDVOICEPRINTENROLL = "InternalError.FailedVoicePrintEnroll"
//  INTERNALERROR_VOICEPRINTAUDIOFAILED = "InternalError.VoicePrintAudioFailed"
//  INTERNALERROR_VOICEPRINTDECODEFAILED = "InternalError.VoicePrintDecodeFailed"
//  INTERNALERROR_VOICEPRINTENROLLFAILED = "InternalError.VoicePrintEnrollFailed"
//  INTERNALERROR_VOICEPRINTVERIFYFAILED = "InternalError.VoicePrintVerifyFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ERRORVOICEDATATOOLONG = "InvalidParameterValue.ErrorVoicedataTooLong"
//  INVALIDPARAMETERVALUE_NOHUMANVOICE = "InvalidParameterValue.NoHumanVoice"
//  LIMITEXCEEDED_VOICEPRINTFULL = "LimitExceeded.VoicePrintFull"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) VoicePrintEnrollWithContext(ctx context.Context, request *VoicePrintEnrollRequest) (response *VoicePrintEnrollResponse, err error) {
    if request == nil {
        request = NewVoicePrintEnrollRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "VoicePrintEnroll")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VoicePrintEnroll require credential")
    }

    request.SetContext(ctx)
    
    response = NewVoicePrintEnrollResponse()
    err = c.Send(request, response)
    return
}

func NewVoicePrintGroupVerifyRequest() (request *VoicePrintGroupVerifyRequest) {
    request = &VoicePrintGroupVerifyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "VoicePrintGroupVerify")
    
    
    return
}

func NewVoicePrintGroupVerifyResponse() (response *VoicePrintGroupVerifyResponse) {
    response = &VoicePrintGroupVerifyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VoicePrintGroupVerify
// 说话人验证1:N接口，可以通过传入一段说话人音频，并且指定已存在的groupId, 和返回topN,  接口返回groupId内所有声纹和传入音频声纹比对打分TopN的结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_USERHASNOFREEAMOUNT = "FailedOperation.UserHasNoFreeAmount"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_FAILEDVOICEPRINTDECODE = "InternalError.FailedVoicePrintDecode"
//  INTERNALERROR_FAILEDVOICEPRINTVERIFY = "InternalError.FailedVoicePrintVerify"
//  INTERNALERROR_VOICEPRINTAUDIOFAILED = "InternalError.VoicePrintAudioFailed"
//  INTERNALERROR_VOICEPRINTDECODEFAILED = "InternalError.VoicePrintDecodeFailed"
//  INTERNALERROR_VOICEPRINTVERIFYFAILED = "InternalError.VoicePrintVerifyFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FAILEDVOICEPRINTDECODE = "InvalidParameter.FailedVoicePrintDecode"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ERRORVOICEDATATOOLONG = "InvalidParameterValue.ErrorVoicedataTooLong"
//  INVALIDPARAMETERVALUE_NOHUMANVOICE = "InvalidParameterValue.NoHumanVoice"
func (c *Client) VoicePrintGroupVerify(request *VoicePrintGroupVerifyRequest) (response *VoicePrintGroupVerifyResponse, err error) {
    return c.VoicePrintGroupVerifyWithContext(context.Background(), request)
}

// VoicePrintGroupVerify
// 说话人验证1:N接口，可以通过传入一段说话人音频，并且指定已存在的groupId, 和返回topN,  接口返回groupId内所有声纹和传入音频声纹比对打分TopN的结果。
//
// 可能返回的错误码:
//  FAILEDOPERATION_USERHASNOFREEAMOUNT = "FailedOperation.UserHasNoFreeAmount"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_FAILEDVOICEPRINTDECODE = "InternalError.FailedVoicePrintDecode"
//  INTERNALERROR_FAILEDVOICEPRINTVERIFY = "InternalError.FailedVoicePrintVerify"
//  INTERNALERROR_VOICEPRINTAUDIOFAILED = "InternalError.VoicePrintAudioFailed"
//  INTERNALERROR_VOICEPRINTDECODEFAILED = "InternalError.VoicePrintDecodeFailed"
//  INTERNALERROR_VOICEPRINTVERIFYFAILED = "InternalError.VoicePrintVerifyFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_FAILEDVOICEPRINTDECODE = "InvalidParameter.FailedVoicePrintDecode"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ERRORVOICEDATATOOLONG = "InvalidParameterValue.ErrorVoicedataTooLong"
//  INVALIDPARAMETERVALUE_NOHUMANVOICE = "InvalidParameterValue.NoHumanVoice"
func (c *Client) VoicePrintGroupVerifyWithContext(ctx context.Context, request *VoicePrintGroupVerifyRequest) (response *VoicePrintGroupVerifyResponse, err error) {
    if request == nil {
        request = NewVoicePrintGroupVerifyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "VoicePrintGroupVerify")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VoicePrintGroupVerify require credential")
    }

    request.SetContext(ctx)
    
    response = NewVoicePrintGroupVerifyResponse()
    err = c.Send(request, response)
    return
}

func NewVoicePrintUpdateRequest() (request *VoicePrintUpdateRequest) {
    request = &VoicePrintUpdateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "VoicePrintUpdate")
    
    
    return
}

func NewVoicePrintUpdateResponse() (response *VoicePrintUpdateResponse) {
    response = &VoicePrintUpdateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VoicePrintUpdate
// 本接口用于更新和覆盖已注册的音频数据和说话人昵称，更新后原有的音频数据将失效。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXISTENTVOICEPRINTID = "FailedOperation.NotExistentVoicePrintId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_ERRORRECOGNIZE = "InternalError.ErrorRecognize"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_FAILEDVOICEPRINTENROLL = "InternalError.FailedVoicePrintEnroll"
//  INTERNALERROR_VOICEPRINTAUDIOFAILED = "InternalError.VoicePrintAudioFailed"
//  INTERNALERROR_VOICEPRINTDECODEFAILED = "InternalError.VoicePrintDecodeFailed"
//  INTERNALERROR_VOICEPRINTENROLLFAILED = "InternalError.VoicePrintEnrollFailed"
//  INTERNALERROR_VOICEPRINTVERIFYFAILED = "InternalError.VoicePrintVerifyFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ERRORVOICEDATATOOLONG = "InvalidParameterValue.ErrorVoicedataTooLong"
//  INVALIDPARAMETERVALUE_NOHUMANVOICE = "InvalidParameterValue.NoHumanVoice"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) VoicePrintUpdate(request *VoicePrintUpdateRequest) (response *VoicePrintUpdateResponse, err error) {
    return c.VoicePrintUpdateWithContext(context.Background(), request)
}

// VoicePrintUpdate
// 本接口用于更新和覆盖已注册的音频数据和说话人昵称，更新后原有的音频数据将失效。
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXISTENTVOICEPRINTID = "FailedOperation.NotExistentVoicePrintId"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_ERRORRECOGNIZE = "InternalError.ErrorRecognize"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_FAILEDVOICEPRINTENROLL = "InternalError.FailedVoicePrintEnroll"
//  INTERNALERROR_VOICEPRINTAUDIOFAILED = "InternalError.VoicePrintAudioFailed"
//  INTERNALERROR_VOICEPRINTDECODEFAILED = "InternalError.VoicePrintDecodeFailed"
//  INTERNALERROR_VOICEPRINTENROLLFAILED = "InternalError.VoicePrintEnrollFailed"
//  INTERNALERROR_VOICEPRINTVERIFYFAILED = "InternalError.VoicePrintVerifyFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE_ERRORVOICEDATATOOLONG = "InvalidParameterValue.ErrorVoicedataTooLong"
//  INVALIDPARAMETERVALUE_NOHUMANVOICE = "InvalidParameterValue.NoHumanVoice"
//  MISSINGPARAMETER = "MissingParameter"
func (c *Client) VoicePrintUpdateWithContext(ctx context.Context, request *VoicePrintUpdateRequest) (response *VoicePrintUpdateResponse, err error) {
    if request == nil {
        request = NewVoicePrintUpdateRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "VoicePrintUpdate")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VoicePrintUpdate require credential")
    }

    request.SetContext(ctx)
    
    response = NewVoicePrintUpdateResponse()
    err = c.Send(request, response)
    return
}

func NewVoicePrintVerifyRequest() (request *VoicePrintVerifyRequest) {
    request = &VoicePrintVerifyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("asr", APIVersion, "VoicePrintVerify")
    
    
    return
}

func NewVoicePrintVerifyResponse() (response *VoicePrintVerifyResponse) {
    response = &VoicePrintVerifyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// VoicePrintVerify
// 本接口用于校验传入音频与已注册音频的匹配程度，通过指定说话人ID（VoicePrintId）和一段音频进行音频和说话人的匹配度判断
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXISTENTVOICEPRINTID = "FailedOperation.NotExistentVoicePrintId"
//  FAILEDOPERATION_USERHASNOFREEAMOUNT = "FailedOperation.UserHasNoFreeAmount"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_ERRORRECOGNIZE = "InternalError.ErrorRecognize"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_FAILEDVOICEPRINTDECODE = "InternalError.FailedVoicePrintDecode"
//  INTERNALERROR_FAILEDVOICEPRINTVERIFY = "InternalError.FailedVoicePrintVerify"
//  INTERNALERROR_VOICEPRINTAUDIOFAILED = "InternalError.VoicePrintAudioFailed"
//  INTERNALERROR_VOICEPRINTDECODEFAILED = "InternalError.VoicePrintDecodeFailed"
//  INTERNALERROR_VOICEPRINTENROLLFAILED = "InternalError.VoicePrintEnrollFailed"
//  INTERNALERROR_VOICEPRINTVERIFYFAILED = "InternalError.VoicePrintVerifyFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ERRORVOICEDATATOOLONG = "InvalidParameterValue.ErrorVoicedataTooLong"
//  INVALIDPARAMETERVALUE_NOHUMANVOICE = "InvalidParameterValue.NoHumanVoice"
func (c *Client) VoicePrintVerify(request *VoicePrintVerifyRequest) (response *VoicePrintVerifyResponse, err error) {
    return c.VoicePrintVerifyWithContext(context.Background(), request)
}

// VoicePrintVerify
// 本接口用于校验传入音频与已注册音频的匹配程度，通过指定说话人ID（VoicePrintId）和一段音频进行音频和说话人的匹配度判断
//
// 可能返回的错误码:
//  FAILEDOPERATION_NOTEXISTENTVOICEPRINTID = "FailedOperation.NotExistentVoicePrintId"
//  FAILEDOPERATION_USERHASNOFREEAMOUNT = "FailedOperation.UserHasNoFreeAmount"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"
//  INTERNALERROR_ERRORRECOGNIZE = "InternalError.ErrorRecognize"
//  INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"
//  INTERNALERROR_FAILEDVOICEPRINTDECODE = "InternalError.FailedVoicePrintDecode"
//  INTERNALERROR_FAILEDVOICEPRINTVERIFY = "InternalError.FailedVoicePrintVerify"
//  INTERNALERROR_VOICEPRINTAUDIOFAILED = "InternalError.VoicePrintAudioFailed"
//  INTERNALERROR_VOICEPRINTDECODEFAILED = "InternalError.VoicePrintDecodeFailed"
//  INTERNALERROR_VOICEPRINTENROLLFAILED = "InternalError.VoicePrintEnrollFailed"
//  INTERNALERROR_VOICEPRINTVERIFYFAILED = "InternalError.VoicePrintVerifyFailed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_ERRORVOICEDATATOOLONG = "InvalidParameterValue.ErrorVoicedataTooLong"
//  INVALIDPARAMETERVALUE_NOHUMANVOICE = "InvalidParameterValue.NoHumanVoice"
func (c *Client) VoicePrintVerifyWithContext(ctx context.Context, request *VoicePrintVerifyRequest) (response *VoicePrintVerifyResponse, err error) {
    if request == nil {
        request = NewVoicePrintVerifyRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "asr", APIVersion, "VoicePrintVerify")
    
    if c.GetCredential() == nil {
        return nil, errors.New("VoicePrintVerify require credential")
    }

    request.SetContext(ctx)
    
    response = NewVoicePrintVerifyResponse()
    err = c.Send(request, response)
    return
}
