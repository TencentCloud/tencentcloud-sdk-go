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

package v20180724

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-07-24"

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


func NewInitOralProcessRequest() (request *InitOralProcessRequest) {
    request = &InitOralProcessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("soe", APIVersion, "InitOralProcess")
    
    
    return
}

func NewInitOralProcessResponse() (response *InitOralProcessResponse) {
    response = &InitOralProcessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// InitOralProcess
// 初始化发音评估过程，每一轮评估前进行调用。语音输入模式分为流式模式和非流式模式，流式模式支持数据分片传输，可以加快评估响应速度。评估模式分为词模式和句子模式，词模式会标注每个音节的详细信息；句子模式会有完整度和流利度的评估。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_ACCOUNTUNAVAILABLE = "AuthFailure.AccountUnavailable"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EVALUATETIMEOUT = "FailedOperation.EvaluateTimeout"
//  FAILEDOPERATION_EVALUATEUNKNOWNERROR = "FailedOperation.EvaluateUnknownError"
//  FAILEDOPERATION_FAILEDGETENGINEIP = "FailedOperation.FailedGetEngineIP"
//  FAILEDOPERATION_FAILEDGETRESULT = "FailedOperation.FailedGetResult"
//  FAILEDOPERATION_FAILEDGETSESSION = "FailedOperation.FailedGetSession"
//  FAILEDOPERATION_FAILEDGETSESSIONSEQID = "FailedOperation.FailedGetSessionSeqID"
//  FAILEDOPERATION_FAILEDGETUSER = "FailedOperation.FailedGetUser"
//  FAILEDOPERATION_FAILEDINIT = "FailedOperation.FailedInit"
//  FAILEDOPERATION_FAILEDSETRESULT = "FailedOperation.FailedSetResult"
//  FAILEDOPERATION_FAILEDSETSESSION = "FailedOperation.FailedSetSession"
//  FAILEDOPERATION_FAILEDSETSESSIONSEQID = "FailedOperation.FailedSetSessionSeqID"
//  FAILEDOPERATION_FAILEDSETUSER = "FailedOperation.FailedSetUser"
//  FAILEDOPERATION_INTERNALSERVERERROR = "FailedOperation.InternalServerError"
//  FAILEDOPERATION_INVALIDPARAMETERVALUE = "FailedOperation.InvalidParameterValue"
//  FAILEDOPERATION_JSONCODECERROR = "FailedOperation.JsonCodecError"
//  FAILEDOPERATION_NEEDINITBEFOREEVALUATION = "FailedOperation.NeedInitBeforeEvaluation"
//  FAILEDOPERATION_PASTSEQIDLOSE = "FailedOperation.PastSeqIdLose"
//  FAILEDOPERATION_RESULTEXPIRED = "FailedOperation.ResultExpired"
//  FAILEDOPERATION_SEQIDEXPIRED = "FailedOperation.SeqIdExpired"
//  FAILEDOPERATION_SERVEROVERLOAD = "FailedOperation.ServerOverload"
//  FAILEDOPERATION_SERVICETIMEOUT = "FailedOperation.ServiceTimeout"
//  FAILEDOPERATION_SESSIONEXPIRED = "FailedOperation.SessionExpired"
//  FAILEDOPERATION_WAITPASTSEQIDTIMEOUT = "FailedOperation.WaitPastSeqIdTimeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_AUDIOPROCESSINGFAILED = "InternalError.AudioProcessingFailed"
//  INTERNALERROR_BASE64DECODEFAILED = "InternalError.BASE64DecodeFailed"
//  INTERNALERROR_ILEGALSERVERRESPONSE = "InternalError.IlegalServerResponse"
//  INTERNALERROR_INITIALPARAMETERERROR = "InternalError.InitialParameterError"
//  INTERNALERROR_INVALIDSEQID = "InternalError.InvalidSeqId"
//  INTERNALERROR_INVALIDWAVHEADER = "InternalError.InvalidWAVHeader"
//  INTERNALERROR_NOCONVERSATIONFOUND = "InternalError.NoConversationFound"
//  INTERNALERROR_REFTXTOOV = "InternalError.RefTxtOov"
//  INTERNALERROR_SERVERINTERNALERROR = "InternalError.ServerInternalError"
//  INTERNALERROR_SERVEROVERLOAD = "InternalError.ServerOverload"
//  INTERNALERROR_SERVICETIMEOUT = "InternalError.ServiceTimeout"
//  INTERNALERROR_VOICEMSGOVERSIZED = "InternalError.VoiceMsgOversized"
//  INTERNALERROR_WORDLENGTHTOOLONG = "InternalError.WordLengthTooLong"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTHORIZEERROR = "InvalidParameter.AuthorizeError"
//  INVALIDPARAMETER_ERRORPHONEME = "InvalidParameter.ErrorPhoneme"
//  INVALIDPARAMETER_INITIALPARAMETERERROR = "InvalidParameter.InitialParameterError"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_SEQIDLIMITEXCEEDED = "InvalidParameter.SeqIdLimitExceeded"
//  INVALIDPARAMETER_VOICEMSGOVERSIZED = "InvalidParameter.VoiceMsgOversized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUDIODATASIZELIMITEXCEEDED = "InvalidParameterValue.AudioDataSizeLimitExceeded"
//  INVALIDPARAMETERVALUE_AUDIODECODEFAILED = "InvalidParameterValue.AudioDecodeFailed"
//  INVALIDPARAMETERVALUE_AUDIOLIMITEXCEEDED = "InvalidParameterValue.AudioLimitExceeded"
//  INVALIDPARAMETERVALUE_AUDIOSIZEMUSTBEEVEN = "InvalidParameterValue.AudioSizeMustBeEven"
//  INVALIDPARAMETERVALUE_BASEDECODEFAILED = "InvalidParameterValue.BASEDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDSEQID = "InvalidParameterValue.InvalidSeqId"
//  INVALIDPARAMETERVALUE_INVALIDWAVHEADER = "InvalidParameterValue.InvalidWAVHeader"
//  INVALIDPARAMETERVALUE_NODOCINLIST = "InvalidParameterValue.NoDocInList"
//  INVALIDPARAMETERVALUE_PARAMETERINVALID = "InvalidParameterValue.ParameterInvalid"
//  INVALIDPARAMETERVALUE_REFTEXTEMPTY = "InvalidParameterValue.RefTextEmpty"
//  INVALIDPARAMETERVALUE_REFTEXTGRAMMARERROR = "InvalidParameterValue.RefTextGrammarError"
//  INVALIDPARAMETERVALUE_REFTEXTLIMITEXCEEDED = "InvalidParameterValue.RefTextLimitExceeded"
//  INVALIDPARAMETERVALUE_REFTEXTOOV = "InvalidParameterValue.RefTextOOV"
//  INVALIDPARAMETERVALUE_REFTXTEMPTY = "InvalidParameterValue.RefTxtEmpty"
//  INVALIDPARAMETERVALUE_REFTXTTOOLANG = "InvalidParameterValue.RefTxtTooLang"
//  INVALIDPARAMETERVALUE_SENSITIVEWORDS = "InvalidParameterValue.SensitiveWords"
//  INVALIDPARAMETERVALUE_SESSIONIDINUSE = "InvalidParameterValue.SessionIdInUse"
//  INVALIDPARAMETERVALUE_SHARDNOSTARTWITHONE = "InvalidParameterValue.ShardNoStartWithOne"
//  INVALIDPARAMETERVALUE_STREAMINGVOICEPKGTIMEOUT = "InvalidParameterValue.StreamingvoicepkgTimeout"
//  INVALIDPARAMETERVALUE_VADNOTDETECTEDSPEAK = "InvalidParameterValue.VadNotDetectedSpeak"
//  INVALIDPARAMETERVALUE_VOICEFILETYPENOTFOUND = "InvalidParameterValue.VoiceFileTypeNotFound"
//  INVALIDPARAMETERVALUE_WAVHEADERDECODEFAILED = "InvalidParameterValue.WAVHeaderDecodeFailed"
//  INVALIDPARAMETERVALUE_WORDLENGTHTOOLONG = "InvalidParameterValue.WordLengthTooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CONCURRENCYLIMITEXCEEDED = "LimitExceeded.ConcurrencyLimitExceeded"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT_SERVERTIMEOUT = "ResourceInsufficient.ServerTimeout"
//  RESOURCEUNAVAILABLE_AUTHORIZEERROR = "ResourceUnavailable.AuthorizeError"
//  RESOURCEUNAVAILABLE_CANNOTFINDSESSION = "ResourceUnavailable.CannotFindSession"
//  RESOURCEUNAVAILABLE_CONCURRENCYLIMIT = "ResourceUnavailable.ConcurrencyLimit"
//  RESOURCEUNAVAILABLE_INITSTREAMNOTSUPPORT = "ResourceUnavailable.InitStreamNotSupport"
//  RESOURCEUNAVAILABLE_INITSTREAMUNFINISHED = "ResourceUnavailable.InitStreamUnfinished"
//  RESOURCEUNAVAILABLE_LASTSEQUNFINISHED = "ResourceUnavailable.LastSeqUnfinished"
//  RESOURCEUNAVAILABLE_NOCONVERSATIONFOUND = "ResourceUnavailable.NoConversationFound"
//  RESOURCEUNAVAILABLE_NOINITBEFOREEVALUATION = "ResourceUnavailable.NoInitBeforeEvaluation"
func (c *Client) InitOralProcess(request *InitOralProcessRequest) (response *InitOralProcessResponse, err error) {
    return c.InitOralProcessWithContext(context.Background(), request)
}

// InitOralProcess
// 初始化发音评估过程，每一轮评估前进行调用。语音输入模式分为流式模式和非流式模式，流式模式支持数据分片传输，可以加快评估响应速度。评估模式分为词模式和句子模式，词模式会标注每个音节的详细信息；句子模式会有完整度和流利度的评估。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_ACCOUNTUNAVAILABLE = "AuthFailure.AccountUnavailable"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_EVALUATETIMEOUT = "FailedOperation.EvaluateTimeout"
//  FAILEDOPERATION_EVALUATEUNKNOWNERROR = "FailedOperation.EvaluateUnknownError"
//  FAILEDOPERATION_FAILEDGETENGINEIP = "FailedOperation.FailedGetEngineIP"
//  FAILEDOPERATION_FAILEDGETRESULT = "FailedOperation.FailedGetResult"
//  FAILEDOPERATION_FAILEDGETSESSION = "FailedOperation.FailedGetSession"
//  FAILEDOPERATION_FAILEDGETSESSIONSEQID = "FailedOperation.FailedGetSessionSeqID"
//  FAILEDOPERATION_FAILEDGETUSER = "FailedOperation.FailedGetUser"
//  FAILEDOPERATION_FAILEDINIT = "FailedOperation.FailedInit"
//  FAILEDOPERATION_FAILEDSETRESULT = "FailedOperation.FailedSetResult"
//  FAILEDOPERATION_FAILEDSETSESSION = "FailedOperation.FailedSetSession"
//  FAILEDOPERATION_FAILEDSETSESSIONSEQID = "FailedOperation.FailedSetSessionSeqID"
//  FAILEDOPERATION_FAILEDSETUSER = "FailedOperation.FailedSetUser"
//  FAILEDOPERATION_INTERNALSERVERERROR = "FailedOperation.InternalServerError"
//  FAILEDOPERATION_INVALIDPARAMETERVALUE = "FailedOperation.InvalidParameterValue"
//  FAILEDOPERATION_JSONCODECERROR = "FailedOperation.JsonCodecError"
//  FAILEDOPERATION_NEEDINITBEFOREEVALUATION = "FailedOperation.NeedInitBeforeEvaluation"
//  FAILEDOPERATION_PASTSEQIDLOSE = "FailedOperation.PastSeqIdLose"
//  FAILEDOPERATION_RESULTEXPIRED = "FailedOperation.ResultExpired"
//  FAILEDOPERATION_SEQIDEXPIRED = "FailedOperation.SeqIdExpired"
//  FAILEDOPERATION_SERVEROVERLOAD = "FailedOperation.ServerOverload"
//  FAILEDOPERATION_SERVICETIMEOUT = "FailedOperation.ServiceTimeout"
//  FAILEDOPERATION_SESSIONEXPIRED = "FailedOperation.SessionExpired"
//  FAILEDOPERATION_WAITPASTSEQIDTIMEOUT = "FailedOperation.WaitPastSeqIdTimeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_AUDIOPROCESSINGFAILED = "InternalError.AudioProcessingFailed"
//  INTERNALERROR_BASE64DECODEFAILED = "InternalError.BASE64DecodeFailed"
//  INTERNALERROR_ILEGALSERVERRESPONSE = "InternalError.IlegalServerResponse"
//  INTERNALERROR_INITIALPARAMETERERROR = "InternalError.InitialParameterError"
//  INTERNALERROR_INVALIDSEQID = "InternalError.InvalidSeqId"
//  INTERNALERROR_INVALIDWAVHEADER = "InternalError.InvalidWAVHeader"
//  INTERNALERROR_NOCONVERSATIONFOUND = "InternalError.NoConversationFound"
//  INTERNALERROR_REFTXTOOV = "InternalError.RefTxtOov"
//  INTERNALERROR_SERVERINTERNALERROR = "InternalError.ServerInternalError"
//  INTERNALERROR_SERVEROVERLOAD = "InternalError.ServerOverload"
//  INTERNALERROR_SERVICETIMEOUT = "InternalError.ServiceTimeout"
//  INTERNALERROR_VOICEMSGOVERSIZED = "InternalError.VoiceMsgOversized"
//  INTERNALERROR_WORDLENGTHTOOLONG = "InternalError.WordLengthTooLong"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTHORIZEERROR = "InvalidParameter.AuthorizeError"
//  INVALIDPARAMETER_ERRORPHONEME = "InvalidParameter.ErrorPhoneme"
//  INVALIDPARAMETER_INITIALPARAMETERERROR = "InvalidParameter.InitialParameterError"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_SEQIDLIMITEXCEEDED = "InvalidParameter.SeqIdLimitExceeded"
//  INVALIDPARAMETER_VOICEMSGOVERSIZED = "InvalidParameter.VoiceMsgOversized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUDIODATASIZELIMITEXCEEDED = "InvalidParameterValue.AudioDataSizeLimitExceeded"
//  INVALIDPARAMETERVALUE_AUDIODECODEFAILED = "InvalidParameterValue.AudioDecodeFailed"
//  INVALIDPARAMETERVALUE_AUDIOLIMITEXCEEDED = "InvalidParameterValue.AudioLimitExceeded"
//  INVALIDPARAMETERVALUE_AUDIOSIZEMUSTBEEVEN = "InvalidParameterValue.AudioSizeMustBeEven"
//  INVALIDPARAMETERVALUE_BASEDECODEFAILED = "InvalidParameterValue.BASEDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDSEQID = "InvalidParameterValue.InvalidSeqId"
//  INVALIDPARAMETERVALUE_INVALIDWAVHEADER = "InvalidParameterValue.InvalidWAVHeader"
//  INVALIDPARAMETERVALUE_NODOCINLIST = "InvalidParameterValue.NoDocInList"
//  INVALIDPARAMETERVALUE_PARAMETERINVALID = "InvalidParameterValue.ParameterInvalid"
//  INVALIDPARAMETERVALUE_REFTEXTEMPTY = "InvalidParameterValue.RefTextEmpty"
//  INVALIDPARAMETERVALUE_REFTEXTGRAMMARERROR = "InvalidParameterValue.RefTextGrammarError"
//  INVALIDPARAMETERVALUE_REFTEXTLIMITEXCEEDED = "InvalidParameterValue.RefTextLimitExceeded"
//  INVALIDPARAMETERVALUE_REFTEXTOOV = "InvalidParameterValue.RefTextOOV"
//  INVALIDPARAMETERVALUE_REFTXTEMPTY = "InvalidParameterValue.RefTxtEmpty"
//  INVALIDPARAMETERVALUE_REFTXTTOOLANG = "InvalidParameterValue.RefTxtTooLang"
//  INVALIDPARAMETERVALUE_SENSITIVEWORDS = "InvalidParameterValue.SensitiveWords"
//  INVALIDPARAMETERVALUE_SESSIONIDINUSE = "InvalidParameterValue.SessionIdInUse"
//  INVALIDPARAMETERVALUE_SHARDNOSTARTWITHONE = "InvalidParameterValue.ShardNoStartWithOne"
//  INVALIDPARAMETERVALUE_STREAMINGVOICEPKGTIMEOUT = "InvalidParameterValue.StreamingvoicepkgTimeout"
//  INVALIDPARAMETERVALUE_VADNOTDETECTEDSPEAK = "InvalidParameterValue.VadNotDetectedSpeak"
//  INVALIDPARAMETERVALUE_VOICEFILETYPENOTFOUND = "InvalidParameterValue.VoiceFileTypeNotFound"
//  INVALIDPARAMETERVALUE_WAVHEADERDECODEFAILED = "InvalidParameterValue.WAVHeaderDecodeFailed"
//  INVALIDPARAMETERVALUE_WORDLENGTHTOOLONG = "InvalidParameterValue.WordLengthTooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CONCURRENCYLIMITEXCEEDED = "LimitExceeded.ConcurrencyLimitExceeded"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT_SERVERTIMEOUT = "ResourceInsufficient.ServerTimeout"
//  RESOURCEUNAVAILABLE_AUTHORIZEERROR = "ResourceUnavailable.AuthorizeError"
//  RESOURCEUNAVAILABLE_CANNOTFINDSESSION = "ResourceUnavailable.CannotFindSession"
//  RESOURCEUNAVAILABLE_CONCURRENCYLIMIT = "ResourceUnavailable.ConcurrencyLimit"
//  RESOURCEUNAVAILABLE_INITSTREAMNOTSUPPORT = "ResourceUnavailable.InitStreamNotSupport"
//  RESOURCEUNAVAILABLE_INITSTREAMUNFINISHED = "ResourceUnavailable.InitStreamUnfinished"
//  RESOURCEUNAVAILABLE_LASTSEQUNFINISHED = "ResourceUnavailable.LastSeqUnfinished"
//  RESOURCEUNAVAILABLE_NOCONVERSATIONFOUND = "ResourceUnavailable.NoConversationFound"
//  RESOURCEUNAVAILABLE_NOINITBEFOREEVALUATION = "ResourceUnavailable.NoInitBeforeEvaluation"
func (c *Client) InitOralProcessWithContext(ctx context.Context, request *InitOralProcessRequest) (response *InitOralProcessResponse, err error) {
    if request == nil {
        request = NewInitOralProcessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("InitOralProcess require credential")
    }

    request.SetContext(ctx)
    
    response = NewInitOralProcessResponse()
    err = c.Send(request, response)
    return
}

func NewKeywordEvaluateRequest() (request *KeywordEvaluateRequest) {
    request = &KeywordEvaluateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("soe", APIVersion, "KeywordEvaluate")
    
    
    return
}

func NewKeywordEvaluateResponse() (response *KeywordEvaluateResponse) {
    response = &KeywordEvaluateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// KeywordEvaluate
// 指定主题关键词词汇评估，分析语音与关键词的切合程度，可指定多个关键词，支持中文英文同时评测。分片传输时，尽量保证纯异步调用，即不等待上一个分片的传输结果边录边传，这样可以尽可能早的提供音频数据。音频源目前仅支持16k采样率16bit单声道编码方式，如有不一致可能导致评估不准确或失败。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_ACCOUNTUNAVAILABLE = "AuthFailure.AccountUnavailable"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORGETSESSION = "FailedOperation.ErrorGetSession"
//  FAILEDOPERATION_EVALUATETIMEOUT = "FailedOperation.EvaluateTimeout"
//  FAILEDOPERATION_EVALUATEUNKNOWNERROR = "FailedOperation.EvaluateUnknownError"
//  FAILEDOPERATION_FAILEDGETENGINEIP = "FailedOperation.FailedGetEngineIP"
//  FAILEDOPERATION_FAILEDGETRESULT = "FailedOperation.FailedGetResult"
//  FAILEDOPERATION_FAILEDGETSESSION = "FailedOperation.FailedGetSession"
//  FAILEDOPERATION_FAILEDGETSESSIONSEQID = "FailedOperation.FailedGetSessionSeqID"
//  FAILEDOPERATION_FAILEDGETUSER = "FailedOperation.FailedGetUser"
//  FAILEDOPERATION_FAILEDINIT = "FailedOperation.FailedInit"
//  FAILEDOPERATION_FAILEDSETRESULT = "FailedOperation.FailedSetResult"
//  FAILEDOPERATION_FAILEDSETSESSION = "FailedOperation.FailedSetSession"
//  FAILEDOPERATION_FAILEDSETSESSIONSEQID = "FailedOperation.FailedSetSessionSeqID"
//  FAILEDOPERATION_FAILEDSETUSER = "FailedOperation.FailedSetUser"
//  FAILEDOPERATION_INTERNALSERVERERROR = "FailedOperation.InternalServerError"
//  FAILEDOPERATION_INVALIDPARAMETERVALUE = "FailedOperation.InvalidParameterValue"
//  FAILEDOPERATION_JSONCODECERROR = "FailedOperation.JsonCodecError"
//  FAILEDOPERATION_NEEDINITBEFOREEVALUATION = "FailedOperation.NeedInitBeforeEvaluation"
//  FAILEDOPERATION_PASTSEQIDLOSE = "FailedOperation.PastSeqIdLose"
//  FAILEDOPERATION_RESULTEXPIRED = "FailedOperation.ResultExpired"
//  FAILEDOPERATION_SEQIDEXPIRED = "FailedOperation.SeqIdExpired"
//  FAILEDOPERATION_SERVEROVERLOAD = "FailedOperation.ServerOverload"
//  FAILEDOPERATION_SERVICETIMEOUT = "FailedOperation.ServiceTimeout"
//  FAILEDOPERATION_SESSIONEXPIRED = "FailedOperation.SessionExpired"
//  FAILEDOPERATION_WAITPASTSEQIDTIMEOUT = "FailedOperation.WaitPastSeqIdTimeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_AUTHORIZEERROR = "InternalError.AuthorizeError"
//  INTERNALERROR_CANNOTFINDSESSION = "InternalError.CannotFindSession"
//  INTERNALERROR_INITSTREAMUNFINISHED = "InternalError.InitStreamUnfinished"
//  INTERNALERROR_INITIALPARAMETERERROR = "InternalError.InitialParameterError"
//  INTERNALERROR_NOINITBEFOREEVALUATION = "InternalError.NoInitBeforeEvaluation"
//  INTERNALERROR_SERVERINTERNALERROR = "InternalError.ServerInternalError"
//  INTERNALERROR_SERVEROVERLOAD = "InternalError.ServerOverload"
//  INTERNALERROR_SERVICETIMEOUT = "InternalError.ServiceTimeout"
//  INTERNALERROR_VOICEMSGOVERSIZED = "InternalError.VoiceMsgOversized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORPHONEME = "InvalidParameter.ErrorPhoneme"
//  INVALIDPARAMETER_INITIALPARAMETERERROR = "InvalidParameter.InitialParameterError"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_SEQIDLIMITEXCEEDED = "InvalidParameter.SeqIdLimitExceeded"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUDIODATASIZELIMITEXCEEDED = "InvalidParameterValue.AudioDataSizeLimitExceeded"
//  INVALIDPARAMETERVALUE_AUDIODECODEFAILED = "InvalidParameterValue.AudioDecodeFailed"
//  INVALIDPARAMETERVALUE_AUDIOLIMITEXCEEDED = "InvalidParameterValue.AudioLimitExceeded"
//  INVALIDPARAMETERVALUE_AUDIOSIZEMUSTBEEVEN = "InvalidParameterValue.AudioSizeMustBeEven"
//  INVALIDPARAMETERVALUE_BASEDECODEFAILED = "InvalidParameterValue.BASEDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDSEQID = "InvalidParameterValue.InvalidSeqId"
//  INVALIDPARAMETERVALUE_INVALIDWAVHEADER = "InvalidParameterValue.InvalidWAVHeader"
//  INVALIDPARAMETERVALUE_REFTEXTEMPTY = "InvalidParameterValue.RefTextEmpty"
//  INVALIDPARAMETERVALUE_REFTEXTGRAMMARERROR = "InvalidParameterValue.RefTextGrammarError"
//  INVALIDPARAMETERVALUE_REFTEXTLIMITEXCEEDED = "InvalidParameterValue.RefTextLimitExceeded"
//  INVALIDPARAMETERVALUE_REFTEXTOOV = "InvalidParameterValue.RefTextOOV"
//  INVALIDPARAMETERVALUE_REFTXTEMPTY = "InvalidParameterValue.RefTxtEmpty"
//  INVALIDPARAMETERVALUE_REFTXTTOOLANG = "InvalidParameterValue.RefTxtTooLang"
//  INVALIDPARAMETERVALUE_SENSITIVEWORDS = "InvalidParameterValue.SensitiveWords"
//  INVALIDPARAMETERVALUE_SESSIONIDINUSE = "InvalidParameterValue.SessionIdInUse"
//  INVALIDPARAMETERVALUE_VOICEFILETYPENOTFOUND = "InvalidParameterValue.VoiceFileTypeNotFound"
//  INVALIDPARAMETERVALUE_WAVHEADERDECODEFAILED = "InvalidParameterValue.WAVHeaderDecodeFailed"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CONCURRENCYLIMITEXCEEDED = "LimitExceeded.ConcurrencyLimitExceeded"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT_SERVERTIMEOUT = "ResourceInsufficient.ServerTimeout"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_AUTHORIZEERROR = "ResourceUnavailable.AuthorizeError"
//  RESOURCEUNAVAILABLE_CANNOTFINDSESSION = "ResourceUnavailable.CannotFindSession"
//  RESOURCEUNAVAILABLE_CONCURRENCYLIMIT = "ResourceUnavailable.ConcurrencyLimit"
//  RESOURCEUNAVAILABLE_NOINITBEFOREEVALUATION = "ResourceUnavailable.NoInitBeforeEvaluation"
func (c *Client) KeywordEvaluate(request *KeywordEvaluateRequest) (response *KeywordEvaluateResponse, err error) {
    return c.KeywordEvaluateWithContext(context.Background(), request)
}

// KeywordEvaluate
// 指定主题关键词词汇评估，分析语音与关键词的切合程度，可指定多个关键词，支持中文英文同时评测。分片传输时，尽量保证纯异步调用，即不等待上一个分片的传输结果边录边传，这样可以尽可能早的提供音频数据。音频源目前仅支持16k采样率16bit单声道编码方式，如有不一致可能导致评估不准确或失败。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_ACCOUNTUNAVAILABLE = "AuthFailure.AccountUnavailable"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORGETSESSION = "FailedOperation.ErrorGetSession"
//  FAILEDOPERATION_EVALUATETIMEOUT = "FailedOperation.EvaluateTimeout"
//  FAILEDOPERATION_EVALUATEUNKNOWNERROR = "FailedOperation.EvaluateUnknownError"
//  FAILEDOPERATION_FAILEDGETENGINEIP = "FailedOperation.FailedGetEngineIP"
//  FAILEDOPERATION_FAILEDGETRESULT = "FailedOperation.FailedGetResult"
//  FAILEDOPERATION_FAILEDGETSESSION = "FailedOperation.FailedGetSession"
//  FAILEDOPERATION_FAILEDGETSESSIONSEQID = "FailedOperation.FailedGetSessionSeqID"
//  FAILEDOPERATION_FAILEDGETUSER = "FailedOperation.FailedGetUser"
//  FAILEDOPERATION_FAILEDINIT = "FailedOperation.FailedInit"
//  FAILEDOPERATION_FAILEDSETRESULT = "FailedOperation.FailedSetResult"
//  FAILEDOPERATION_FAILEDSETSESSION = "FailedOperation.FailedSetSession"
//  FAILEDOPERATION_FAILEDSETSESSIONSEQID = "FailedOperation.FailedSetSessionSeqID"
//  FAILEDOPERATION_FAILEDSETUSER = "FailedOperation.FailedSetUser"
//  FAILEDOPERATION_INTERNALSERVERERROR = "FailedOperation.InternalServerError"
//  FAILEDOPERATION_INVALIDPARAMETERVALUE = "FailedOperation.InvalidParameterValue"
//  FAILEDOPERATION_JSONCODECERROR = "FailedOperation.JsonCodecError"
//  FAILEDOPERATION_NEEDINITBEFOREEVALUATION = "FailedOperation.NeedInitBeforeEvaluation"
//  FAILEDOPERATION_PASTSEQIDLOSE = "FailedOperation.PastSeqIdLose"
//  FAILEDOPERATION_RESULTEXPIRED = "FailedOperation.ResultExpired"
//  FAILEDOPERATION_SEQIDEXPIRED = "FailedOperation.SeqIdExpired"
//  FAILEDOPERATION_SERVEROVERLOAD = "FailedOperation.ServerOverload"
//  FAILEDOPERATION_SERVICETIMEOUT = "FailedOperation.ServiceTimeout"
//  FAILEDOPERATION_SESSIONEXPIRED = "FailedOperation.SessionExpired"
//  FAILEDOPERATION_WAITPASTSEQIDTIMEOUT = "FailedOperation.WaitPastSeqIdTimeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_AUTHORIZEERROR = "InternalError.AuthorizeError"
//  INTERNALERROR_CANNOTFINDSESSION = "InternalError.CannotFindSession"
//  INTERNALERROR_INITSTREAMUNFINISHED = "InternalError.InitStreamUnfinished"
//  INTERNALERROR_INITIALPARAMETERERROR = "InternalError.InitialParameterError"
//  INTERNALERROR_NOINITBEFOREEVALUATION = "InternalError.NoInitBeforeEvaluation"
//  INTERNALERROR_SERVERINTERNALERROR = "InternalError.ServerInternalError"
//  INTERNALERROR_SERVEROVERLOAD = "InternalError.ServerOverload"
//  INTERNALERROR_SERVICETIMEOUT = "InternalError.ServiceTimeout"
//  INTERNALERROR_VOICEMSGOVERSIZED = "InternalError.VoiceMsgOversized"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ERRORPHONEME = "InvalidParameter.ErrorPhoneme"
//  INVALIDPARAMETER_INITIALPARAMETERERROR = "InvalidParameter.InitialParameterError"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_SEQIDLIMITEXCEEDED = "InvalidParameter.SeqIdLimitExceeded"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUDIODATASIZELIMITEXCEEDED = "InvalidParameterValue.AudioDataSizeLimitExceeded"
//  INVALIDPARAMETERVALUE_AUDIODECODEFAILED = "InvalidParameterValue.AudioDecodeFailed"
//  INVALIDPARAMETERVALUE_AUDIOLIMITEXCEEDED = "InvalidParameterValue.AudioLimitExceeded"
//  INVALIDPARAMETERVALUE_AUDIOSIZEMUSTBEEVEN = "InvalidParameterValue.AudioSizeMustBeEven"
//  INVALIDPARAMETERVALUE_BASEDECODEFAILED = "InvalidParameterValue.BASEDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDSEQID = "InvalidParameterValue.InvalidSeqId"
//  INVALIDPARAMETERVALUE_INVALIDWAVHEADER = "InvalidParameterValue.InvalidWAVHeader"
//  INVALIDPARAMETERVALUE_REFTEXTEMPTY = "InvalidParameterValue.RefTextEmpty"
//  INVALIDPARAMETERVALUE_REFTEXTGRAMMARERROR = "InvalidParameterValue.RefTextGrammarError"
//  INVALIDPARAMETERVALUE_REFTEXTLIMITEXCEEDED = "InvalidParameterValue.RefTextLimitExceeded"
//  INVALIDPARAMETERVALUE_REFTEXTOOV = "InvalidParameterValue.RefTextOOV"
//  INVALIDPARAMETERVALUE_REFTXTEMPTY = "InvalidParameterValue.RefTxtEmpty"
//  INVALIDPARAMETERVALUE_REFTXTTOOLANG = "InvalidParameterValue.RefTxtTooLang"
//  INVALIDPARAMETERVALUE_SENSITIVEWORDS = "InvalidParameterValue.SensitiveWords"
//  INVALIDPARAMETERVALUE_SESSIONIDINUSE = "InvalidParameterValue.SessionIdInUse"
//  INVALIDPARAMETERVALUE_VOICEFILETYPENOTFOUND = "InvalidParameterValue.VoiceFileTypeNotFound"
//  INVALIDPARAMETERVALUE_WAVHEADERDECODEFAILED = "InvalidParameterValue.WAVHeaderDecodeFailed"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CONCURRENCYLIMITEXCEEDED = "LimitExceeded.ConcurrencyLimitExceeded"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT_SERVERTIMEOUT = "ResourceInsufficient.ServerTimeout"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_AUTHORIZEERROR = "ResourceUnavailable.AuthorizeError"
//  RESOURCEUNAVAILABLE_CANNOTFINDSESSION = "ResourceUnavailable.CannotFindSession"
//  RESOURCEUNAVAILABLE_CONCURRENCYLIMIT = "ResourceUnavailable.ConcurrencyLimit"
//  RESOURCEUNAVAILABLE_NOINITBEFOREEVALUATION = "ResourceUnavailable.NoInitBeforeEvaluation"
func (c *Client) KeywordEvaluateWithContext(ctx context.Context, request *KeywordEvaluateRequest) (response *KeywordEvaluateResponse, err error) {
    if request == nil {
        request = NewKeywordEvaluateRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("KeywordEvaluate require credential")
    }

    request.SetContext(ctx)
    
    response = NewKeywordEvaluateResponse()
    err = c.Send(request, response)
    return
}

func NewTransmitOralProcessRequest() (request *TransmitOralProcessRequest) {
    request = &TransmitOralProcessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("soe", APIVersion, "TransmitOralProcess")
    
    
    return
}

func NewTransmitOralProcessResponse() (response *TransmitOralProcessResponse) {
    response = &TransmitOralProcessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TransmitOralProcess
// 本接口可用于中英文发音评测数据传输。在使用本接口时需要注意：传输音频数据，必须在完成发音评估初始化接口之后调用，且SessonId要与初始化接口保持一致。分片传输时，尽量保证SeqId顺序传输（请确认SeqId由1开始）。音频源目前仅支持16k采样率16bit单声道编码方式，如有不一致可能导致评估不准确或失败。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_ACCOUNTUNAVAILABLE = "AuthFailure.AccountUnavailable"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORGETSESSION = "FailedOperation.ErrorGetSession"
//  FAILEDOPERATION_EVALUATETIMEOUT = "FailedOperation.EvaluateTimeout"
//  FAILEDOPERATION_EVALUATEUNKNOWNERROR = "FailedOperation.EvaluateUnknownError"
//  FAILEDOPERATION_FAILEDGETENGINEIP = "FailedOperation.FailedGetEngineIP"
//  FAILEDOPERATION_FAILEDGETRESULT = "FailedOperation.FailedGetResult"
//  FAILEDOPERATION_FAILEDGETSESSION = "FailedOperation.FailedGetSession"
//  FAILEDOPERATION_FAILEDGETSESSIONSEQID = "FailedOperation.FailedGetSessionSeqID"
//  FAILEDOPERATION_FAILEDGETUSER = "FailedOperation.FailedGetUser"
//  FAILEDOPERATION_FAILEDINIT = "FailedOperation.FailedInit"
//  FAILEDOPERATION_FAILEDSETRESULT = "FailedOperation.FailedSetResult"
//  FAILEDOPERATION_FAILEDSETSESSION = "FailedOperation.FailedSetSession"
//  FAILEDOPERATION_FAILEDSETSESSIONSEQID = "FailedOperation.FailedSetSessionSeqID"
//  FAILEDOPERATION_FAILEDSETUSER = "FailedOperation.FailedSetUser"
//  FAILEDOPERATION_INTERNALSERVERERROR = "FailedOperation.InternalServerError"
//  FAILEDOPERATION_INVALIDPARAMETERVALUE = "FailedOperation.InvalidParameterValue"
//  FAILEDOPERATION_JSONCODECERROR = "FailedOperation.JsonCodecError"
//  FAILEDOPERATION_NEEDINITBEFOREEVALUATION = "FailedOperation.NeedInitBeforeEvaluation"
//  FAILEDOPERATION_PASTSEQIDLOSE = "FailedOperation.PastSeqIdLose"
//  FAILEDOPERATION_RESULTEXPIRED = "FailedOperation.ResultExpired"
//  FAILEDOPERATION_SEQIDEXPIRED = "FailedOperation.SeqIdExpired"
//  FAILEDOPERATION_SERVEROVERLOAD = "FailedOperation.ServerOverload"
//  FAILEDOPERATION_SERVICETIMEOUT = "FailedOperation.ServiceTimeout"
//  FAILEDOPERATION_SESSIONEXPIRED = "FailedOperation.SessionExpired"
//  FAILEDOPERATION_WAITPASTSEQIDTIMEOUT = "FailedOperation.WaitPastSeqIdTimeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_AUDIOPROCESSINGFAILED = "InternalError.AudioProcessingFailed"
//  INTERNALERROR_AUTHORIZEERROR = "InternalError.AuthorizeError"
//  INTERNALERROR_BASE64DECODEFAILED = "InternalError.BASE64DecodeFailed"
//  INTERNALERROR_CANNOTFINDSESSION = "InternalError.CannotFindSession"
//  INTERNALERROR_FAILTODECODEVOICE = "InternalError.FailToDecodeVoice"
//  INTERNALERROR_ILEGALSERVERRESPONSE = "InternalError.IlegalServerResponse"
//  INTERNALERROR_INITIALPARAMETERERROR = "InternalError.InitialParameterError"
//  INTERNALERROR_INVALIDSEQID = "InternalError.InvalidSeqId"
//  INTERNALERROR_INVALIDWAVHEADER = "InternalError.InvalidWAVHeader"
//  INTERNALERROR_LASTSEQUNFINISHED = "InternalError.LastSeqUnfinished"
//  INTERNALERROR_MP3DECODEFAILED = "InternalError.MP3DecodeFailed"
//  INTERNALERROR_NEEDTOINIT = "InternalError.NeedToInit"
//  INTERNALERROR_NOCONVERSATIONFOUND = "InternalError.NoConversationFound"
//  INTERNALERROR_NOINITBEFOREEVALUATION = "InternalError.NoInitBeforeEvaluation"
//  INTERNALERROR_REFTXTEMPTY = "InternalError.RefTxtEmpty"
//  INTERNALERROR_REFTXTOOV = "InternalError.RefTxtOov"
//  INTERNALERROR_REFTXTTOOLANG = "InternalError.RefTxtTooLang"
//  INTERNALERROR_SERVERINTERNALERROR = "InternalError.ServerInternalError"
//  INTERNALERROR_SERVEROVERLOAD = "InternalError.ServerOverload"
//  INTERNALERROR_SERVICETIMEOUT = "InternalError.ServiceTimeout"
//  INTERNALERROR_SHARDNOSTARTWITHONE = "InternalError.ShardNoStartWithOne"
//  INTERNALERROR_STREAMINGVOICEPKGTIMEOUT = "InternalError.StreamingvoicepkgTimeout"
//  INTERNALERROR_VOICEMSGOVERSIZED = "InternalError.VoiceMsgOversized"
//  INTERNALERROR_WORDLENGTHTOOLONG = "InternalError.WordLengthTooLong"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTHORIZEERROR = "InvalidParameter.AuthorizeError"
//  INVALIDPARAMETER_ERRORPHONEME = "InvalidParameter.ErrorPhoneme"
//  INVALIDPARAMETER_INITIALPARAMETERERROR = "InvalidParameter.InitialParameterError"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_SEQIDLIMITEXCEEDED = "InvalidParameter.SeqIdLimitExceeded"
//  INVALIDPARAMETER_VOICEMSGOVERSIZED = "InvalidParameter.VoiceMsgOversized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUDIODATASIZELIMITEXCEEDED = "InvalidParameterValue.AudioDataSizeLimitExceeded"
//  INVALIDPARAMETERVALUE_AUDIODECODEFAILED = "InvalidParameterValue.AudioDecodeFailed"
//  INVALIDPARAMETERVALUE_AUDIOLIMITEXCEEDED = "InvalidParameterValue.AudioLimitExceeded"
//  INVALIDPARAMETERVALUE_AUDIOSIZEMUSTBEEVEN = "InvalidParameterValue.AudioSizeMustBeEven"
//  INVALIDPARAMETERVALUE_BASEDECODEFAILED = "InvalidParameterValue.BASEDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDSEQID = "InvalidParameterValue.InvalidSeqId"
//  INVALIDPARAMETERVALUE_INVALIDWAVHEADER = "InvalidParameterValue.InvalidWAVHeader"
//  INVALIDPARAMETERVALUE_NODOCINLIST = "InvalidParameterValue.NoDocInList"
//  INVALIDPARAMETERVALUE_PARAMETERINVALID = "InvalidParameterValue.ParameterInvalid"
//  INVALIDPARAMETERVALUE_REFTEXTEMPTY = "InvalidParameterValue.RefTextEmpty"
//  INVALIDPARAMETERVALUE_REFTEXTGRAMMARERROR = "InvalidParameterValue.RefTextGrammarError"
//  INVALIDPARAMETERVALUE_REFTEXTLIMITEXCEEDED = "InvalidParameterValue.RefTextLimitExceeded"
//  INVALIDPARAMETERVALUE_REFTEXTOOV = "InvalidParameterValue.RefTextOOV"
//  INVALIDPARAMETERVALUE_REFTEXTPOLYPHONICLIMITEXCEEDED = "InvalidParameterValue.RefTextPolyphonicLimitExceeded"
//  INVALIDPARAMETERVALUE_REFTXTEMPTY = "InvalidParameterValue.RefTxtEmpty"
//  INVALIDPARAMETERVALUE_REFTXTTOOLANG = "InvalidParameterValue.RefTxtTooLang"
//  INVALIDPARAMETERVALUE_SENSITIVEWORDS = "InvalidParameterValue.SensitiveWords"
//  INVALIDPARAMETERVALUE_SESSIONIDINUSE = "InvalidParameterValue.SessionIdInUse"
//  INVALIDPARAMETERVALUE_SHARDNOSTARTWITHONE = "InvalidParameterValue.ShardNoStartWithOne"
//  INVALIDPARAMETERVALUE_STREAMINGVOICEPKGTIMEOUT = "InvalidParameterValue.StreamingvoicepkgTimeout"
//  INVALIDPARAMETERVALUE_VADNOTDETECTEDSPEAK = "InvalidParameterValue.VadNotDetectedSpeak"
//  INVALIDPARAMETERVALUE_VOICEFILETYPENOTFOUND = "InvalidParameterValue.VoiceFileTypeNotFound"
//  INVALIDPARAMETERVALUE_WAVHEADERDECODEFAILED = "InvalidParameterValue.WAVHeaderDecodeFailed"
//  INVALIDPARAMETERVALUE_WORDLENGTHTOOLONG = "InvalidParameterValue.WordLengthTooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CONCURRENCYLIMITEXCEEDED = "LimitExceeded.ConcurrencyLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT_SERVERTIMEOUT = "ResourceInsufficient.ServerTimeout"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_AUTHORIZEERROR = "ResourceUnavailable.AuthorizeError"
//  RESOURCEUNAVAILABLE_CANNOTFINDSESSION = "ResourceUnavailable.CannotFindSession"
//  RESOURCEUNAVAILABLE_CONCURRENCYLIMIT = "ResourceUnavailable.ConcurrencyLimit"
//  RESOURCEUNAVAILABLE_INITSTREAMNOTSUPPORT = "ResourceUnavailable.InitStreamNotSupport"
//  RESOURCEUNAVAILABLE_INITSTREAMUNFINISHED = "ResourceUnavailable.InitStreamUnfinished"
//  RESOURCEUNAVAILABLE_LASTSEQUNFINISHED = "ResourceUnavailable.LastSeqUnfinished"
//  RESOURCEUNAVAILABLE_NOCONVERSATIONFOUND = "ResourceUnavailable.NoConversationFound"
//  RESOURCEUNAVAILABLE_NOINITBEFOREEVALUATION = "ResourceUnavailable.NoInitBeforeEvaluation"
func (c *Client) TransmitOralProcess(request *TransmitOralProcessRequest) (response *TransmitOralProcessResponse, err error) {
    return c.TransmitOralProcessWithContext(context.Background(), request)
}

// TransmitOralProcess
// 本接口可用于中英文发音评测数据传输。在使用本接口时需要注意：传输音频数据，必须在完成发音评估初始化接口之后调用，且SessonId要与初始化接口保持一致。分片传输时，尽量保证SeqId顺序传输（请确认SeqId由1开始）。音频源目前仅支持16k采样率16bit单声道编码方式，如有不一致可能导致评估不准确或失败。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_ACCOUNTUNAVAILABLE = "AuthFailure.AccountUnavailable"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORGETSESSION = "FailedOperation.ErrorGetSession"
//  FAILEDOPERATION_EVALUATETIMEOUT = "FailedOperation.EvaluateTimeout"
//  FAILEDOPERATION_EVALUATEUNKNOWNERROR = "FailedOperation.EvaluateUnknownError"
//  FAILEDOPERATION_FAILEDGETENGINEIP = "FailedOperation.FailedGetEngineIP"
//  FAILEDOPERATION_FAILEDGETRESULT = "FailedOperation.FailedGetResult"
//  FAILEDOPERATION_FAILEDGETSESSION = "FailedOperation.FailedGetSession"
//  FAILEDOPERATION_FAILEDGETSESSIONSEQID = "FailedOperation.FailedGetSessionSeqID"
//  FAILEDOPERATION_FAILEDGETUSER = "FailedOperation.FailedGetUser"
//  FAILEDOPERATION_FAILEDINIT = "FailedOperation.FailedInit"
//  FAILEDOPERATION_FAILEDSETRESULT = "FailedOperation.FailedSetResult"
//  FAILEDOPERATION_FAILEDSETSESSION = "FailedOperation.FailedSetSession"
//  FAILEDOPERATION_FAILEDSETSESSIONSEQID = "FailedOperation.FailedSetSessionSeqID"
//  FAILEDOPERATION_FAILEDSETUSER = "FailedOperation.FailedSetUser"
//  FAILEDOPERATION_INTERNALSERVERERROR = "FailedOperation.InternalServerError"
//  FAILEDOPERATION_INVALIDPARAMETERVALUE = "FailedOperation.InvalidParameterValue"
//  FAILEDOPERATION_JSONCODECERROR = "FailedOperation.JsonCodecError"
//  FAILEDOPERATION_NEEDINITBEFOREEVALUATION = "FailedOperation.NeedInitBeforeEvaluation"
//  FAILEDOPERATION_PASTSEQIDLOSE = "FailedOperation.PastSeqIdLose"
//  FAILEDOPERATION_RESULTEXPIRED = "FailedOperation.ResultExpired"
//  FAILEDOPERATION_SEQIDEXPIRED = "FailedOperation.SeqIdExpired"
//  FAILEDOPERATION_SERVEROVERLOAD = "FailedOperation.ServerOverload"
//  FAILEDOPERATION_SERVICETIMEOUT = "FailedOperation.ServiceTimeout"
//  FAILEDOPERATION_SESSIONEXPIRED = "FailedOperation.SessionExpired"
//  FAILEDOPERATION_WAITPASTSEQIDTIMEOUT = "FailedOperation.WaitPastSeqIdTimeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_AUDIOPROCESSINGFAILED = "InternalError.AudioProcessingFailed"
//  INTERNALERROR_AUTHORIZEERROR = "InternalError.AuthorizeError"
//  INTERNALERROR_BASE64DECODEFAILED = "InternalError.BASE64DecodeFailed"
//  INTERNALERROR_CANNOTFINDSESSION = "InternalError.CannotFindSession"
//  INTERNALERROR_FAILTODECODEVOICE = "InternalError.FailToDecodeVoice"
//  INTERNALERROR_ILEGALSERVERRESPONSE = "InternalError.IlegalServerResponse"
//  INTERNALERROR_INITIALPARAMETERERROR = "InternalError.InitialParameterError"
//  INTERNALERROR_INVALIDSEQID = "InternalError.InvalidSeqId"
//  INTERNALERROR_INVALIDWAVHEADER = "InternalError.InvalidWAVHeader"
//  INTERNALERROR_LASTSEQUNFINISHED = "InternalError.LastSeqUnfinished"
//  INTERNALERROR_MP3DECODEFAILED = "InternalError.MP3DecodeFailed"
//  INTERNALERROR_NEEDTOINIT = "InternalError.NeedToInit"
//  INTERNALERROR_NOCONVERSATIONFOUND = "InternalError.NoConversationFound"
//  INTERNALERROR_NOINITBEFOREEVALUATION = "InternalError.NoInitBeforeEvaluation"
//  INTERNALERROR_REFTXTEMPTY = "InternalError.RefTxtEmpty"
//  INTERNALERROR_REFTXTOOV = "InternalError.RefTxtOov"
//  INTERNALERROR_REFTXTTOOLANG = "InternalError.RefTxtTooLang"
//  INTERNALERROR_SERVERINTERNALERROR = "InternalError.ServerInternalError"
//  INTERNALERROR_SERVEROVERLOAD = "InternalError.ServerOverload"
//  INTERNALERROR_SERVICETIMEOUT = "InternalError.ServiceTimeout"
//  INTERNALERROR_SHARDNOSTARTWITHONE = "InternalError.ShardNoStartWithOne"
//  INTERNALERROR_STREAMINGVOICEPKGTIMEOUT = "InternalError.StreamingvoicepkgTimeout"
//  INTERNALERROR_VOICEMSGOVERSIZED = "InternalError.VoiceMsgOversized"
//  INTERNALERROR_WORDLENGTHTOOLONG = "InternalError.WordLengthTooLong"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTHORIZEERROR = "InvalidParameter.AuthorizeError"
//  INVALIDPARAMETER_ERRORPHONEME = "InvalidParameter.ErrorPhoneme"
//  INVALIDPARAMETER_INITIALPARAMETERERROR = "InvalidParameter.InitialParameterError"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_SEQIDLIMITEXCEEDED = "InvalidParameter.SeqIdLimitExceeded"
//  INVALIDPARAMETER_VOICEMSGOVERSIZED = "InvalidParameter.VoiceMsgOversized"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUDIODATASIZELIMITEXCEEDED = "InvalidParameterValue.AudioDataSizeLimitExceeded"
//  INVALIDPARAMETERVALUE_AUDIODECODEFAILED = "InvalidParameterValue.AudioDecodeFailed"
//  INVALIDPARAMETERVALUE_AUDIOLIMITEXCEEDED = "InvalidParameterValue.AudioLimitExceeded"
//  INVALIDPARAMETERVALUE_AUDIOSIZEMUSTBEEVEN = "InvalidParameterValue.AudioSizeMustBeEven"
//  INVALIDPARAMETERVALUE_BASEDECODEFAILED = "InvalidParameterValue.BASEDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDSEQID = "InvalidParameterValue.InvalidSeqId"
//  INVALIDPARAMETERVALUE_INVALIDWAVHEADER = "InvalidParameterValue.InvalidWAVHeader"
//  INVALIDPARAMETERVALUE_NODOCINLIST = "InvalidParameterValue.NoDocInList"
//  INVALIDPARAMETERVALUE_PARAMETERINVALID = "InvalidParameterValue.ParameterInvalid"
//  INVALIDPARAMETERVALUE_REFTEXTEMPTY = "InvalidParameterValue.RefTextEmpty"
//  INVALIDPARAMETERVALUE_REFTEXTGRAMMARERROR = "InvalidParameterValue.RefTextGrammarError"
//  INVALIDPARAMETERVALUE_REFTEXTLIMITEXCEEDED = "InvalidParameterValue.RefTextLimitExceeded"
//  INVALIDPARAMETERVALUE_REFTEXTOOV = "InvalidParameterValue.RefTextOOV"
//  INVALIDPARAMETERVALUE_REFTEXTPOLYPHONICLIMITEXCEEDED = "InvalidParameterValue.RefTextPolyphonicLimitExceeded"
//  INVALIDPARAMETERVALUE_REFTXTEMPTY = "InvalidParameterValue.RefTxtEmpty"
//  INVALIDPARAMETERVALUE_REFTXTTOOLANG = "InvalidParameterValue.RefTxtTooLang"
//  INVALIDPARAMETERVALUE_SENSITIVEWORDS = "InvalidParameterValue.SensitiveWords"
//  INVALIDPARAMETERVALUE_SESSIONIDINUSE = "InvalidParameterValue.SessionIdInUse"
//  INVALIDPARAMETERVALUE_SHARDNOSTARTWITHONE = "InvalidParameterValue.ShardNoStartWithOne"
//  INVALIDPARAMETERVALUE_STREAMINGVOICEPKGTIMEOUT = "InvalidParameterValue.StreamingvoicepkgTimeout"
//  INVALIDPARAMETERVALUE_VADNOTDETECTEDSPEAK = "InvalidParameterValue.VadNotDetectedSpeak"
//  INVALIDPARAMETERVALUE_VOICEFILETYPENOTFOUND = "InvalidParameterValue.VoiceFileTypeNotFound"
//  INVALIDPARAMETERVALUE_WAVHEADERDECODEFAILED = "InvalidParameterValue.WAVHeaderDecodeFailed"
//  INVALIDPARAMETERVALUE_WORDLENGTHTOOLONG = "InvalidParameterValue.WordLengthTooLong"
//  LIMITEXCEEDED = "LimitExceeded"
//  LIMITEXCEEDED_CONCURRENCYLIMITEXCEEDED = "LimitExceeded.ConcurrencyLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT_SERVERTIMEOUT = "ResourceInsufficient.ServerTimeout"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_AUTHORIZEERROR = "ResourceUnavailable.AuthorizeError"
//  RESOURCEUNAVAILABLE_CANNOTFINDSESSION = "ResourceUnavailable.CannotFindSession"
//  RESOURCEUNAVAILABLE_CONCURRENCYLIMIT = "ResourceUnavailable.ConcurrencyLimit"
//  RESOURCEUNAVAILABLE_INITSTREAMNOTSUPPORT = "ResourceUnavailable.InitStreamNotSupport"
//  RESOURCEUNAVAILABLE_INITSTREAMUNFINISHED = "ResourceUnavailable.InitStreamUnfinished"
//  RESOURCEUNAVAILABLE_LASTSEQUNFINISHED = "ResourceUnavailable.LastSeqUnfinished"
//  RESOURCEUNAVAILABLE_NOCONVERSATIONFOUND = "ResourceUnavailable.NoConversationFound"
//  RESOURCEUNAVAILABLE_NOINITBEFOREEVALUATION = "ResourceUnavailable.NoInitBeforeEvaluation"
func (c *Client) TransmitOralProcessWithContext(ctx context.Context, request *TransmitOralProcessRequest) (response *TransmitOralProcessResponse, err error) {
    if request == nil {
        request = NewTransmitOralProcessRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TransmitOralProcess require credential")
    }

    request.SetContext(ctx)
    
    response = NewTransmitOralProcessResponse()
    err = c.Send(request, response)
    return
}

func NewTransmitOralProcessWithInitRequest() (request *TransmitOralProcessWithInitRequest) {
    request = &TransmitOralProcessWithInitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("soe", APIVersion, "TransmitOralProcessWithInit")
    
    
    return
}

func NewTransmitOralProcessWithInitResponse() (response *TransmitOralProcessWithInitResponse) {
    response = &TransmitOralProcessWithInitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TransmitOralProcessWithInit
// 本接口可用于中英文发音评测。在使用本接口时需要注意：初始化并传输音频数据，分片传输时，尽量保证SeqId顺序传输（请确认SeqId由1开始）。音频源目前仅支持16k采样率16bit单声道编码方式，如有不一致可能导致评估不准确或失败。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_ACCOUNTUNAVAILABLE = "AuthFailure.AccountUnavailable"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORGETSESSION = "FailedOperation.ErrorGetSession"
//  FAILEDOPERATION_ERRORGETUSER = "FailedOperation.ErrorGetUser"
//  FAILEDOPERATION_EVALUATETIMEOUT = "FailedOperation.EvaluateTimeout"
//  FAILEDOPERATION_EVALUATEUNKNOWNERROR = "FailedOperation.EvaluateUnknownError"
//  FAILEDOPERATION_FAILEDGETENGINEIP = "FailedOperation.FailedGetEngineIP"
//  FAILEDOPERATION_FAILEDGETRESULT = "FailedOperation.FailedGetResult"
//  FAILEDOPERATION_FAILEDGETSESSION = "FailedOperation.FailedGetSession"
//  FAILEDOPERATION_FAILEDGETSESSIONSEQID = "FailedOperation.FailedGetSessionSeqID"
//  FAILEDOPERATION_FAILEDGETUSER = "FailedOperation.FailedGetUser"
//  FAILEDOPERATION_FAILEDINIT = "FailedOperation.FailedInit"
//  FAILEDOPERATION_FAILEDSETRESULT = "FailedOperation.FailedSetResult"
//  FAILEDOPERATION_FAILEDSETSESSION = "FailedOperation.FailedSetSession"
//  FAILEDOPERATION_FAILEDSETSESSIONSEQID = "FailedOperation.FailedSetSessionSeqID"
//  FAILEDOPERATION_FAILEDSETUSER = "FailedOperation.FailedSetUser"
//  FAILEDOPERATION_INTERNALSERVERERROR = "FailedOperation.InternalServerError"
//  FAILEDOPERATION_INVALIDPARAMETERVALUE = "FailedOperation.InvalidParameterValue"
//  FAILEDOPERATION_JSONCODECERROR = "FailedOperation.JsonCodecError"
//  FAILEDOPERATION_NEEDINITBEFOREEVALUATION = "FailedOperation.NeedInitBeforeEvaluation"
//  FAILEDOPERATION_PASTSEQIDLOSE = "FailedOperation.PastSeqIdLose"
//  FAILEDOPERATION_RESULTEXPIRED = "FailedOperation.ResultExpired"
//  FAILEDOPERATION_SEQIDEXPIRED = "FailedOperation.SeqIdExpired"
//  FAILEDOPERATION_SERVEROVERLOAD = "FailedOperation.ServerOverload"
//  FAILEDOPERATION_SERVICETIMEOUT = "FailedOperation.ServiceTimeout"
//  FAILEDOPERATION_SESSIONEXPIRED = "FailedOperation.SessionExpired"
//  FAILEDOPERATION_WAITPASTSEQIDTIMEOUT = "FailedOperation.WaitPastSeqIdTimeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_AUDIOPROCESSINGFAILED = "InternalError.AudioProcessingFailed"
//  INTERNALERROR_BASE64DECODEFAILED = "InternalError.BASE64DecodeFailed"
//  INTERNALERROR_CANNOTFINDSESSION = "InternalError.CannotFindSession"
//  INTERNALERROR_FAILTODECODEVOICE = "InternalError.FailToDecodeVoice"
//  INTERNALERROR_ILEGALSERVERRESPONSE = "InternalError.IlegalServerResponse"
//  INTERNALERROR_INITSTREAMNOTSUPPORT = "InternalError.InitStreamNotSupport"
//  INTERNALERROR_INITSTREAMUNFINISHED = "InternalError.InitStreamUnfinished"
//  INTERNALERROR_INITIALPARAMETERERROR = "InternalError.InitialParameterError"
//  INTERNALERROR_INVALIDSEQID = "InternalError.InvalidSeqId"
//  INTERNALERROR_INVALIDWAVHEADER = "InternalError.InvalidWAVHeader"
//  INTERNALERROR_LASTSEQUNFINISHED = "InternalError.LastSeqUnfinished"
//  INTERNALERROR_MP3DECODEFAILED = "InternalError.MP3DecodeFailed"
//  INTERNALERROR_NEEDTOINIT = "InternalError.NeedToInit"
//  INTERNALERROR_NOCONVERSATIONFOUND = "InternalError.NoConversationFound"
//  INTERNALERROR_NODOCINLIST = "InternalError.NoDocInList"
//  INTERNALERROR_NOERROR = "InternalError.NoError"
//  INTERNALERROR_NOINITBEFOREEVALUATION = "InternalError.NoInitBeforeEvaluation"
//  INTERNALERROR_REFTXTEMPTY = "InternalError.RefTxtEmpty"
//  INTERNALERROR_REFTXTOOV = "InternalError.RefTxtOov"
//  INTERNALERROR_REFTXTTOOLANG = "InternalError.RefTxtTooLang"
//  INTERNALERROR_SERVERINTERNALERROR = "InternalError.ServerInternalError"
//  INTERNALERROR_SERVEROVERLOAD = "InternalError.ServerOverload"
//  INTERNALERROR_SERVICETIMEOUT = "InternalError.ServiceTimeout"
//  INTERNALERROR_SHARDNOSTARTWITHONE = "InternalError.ShardNoStartWithOne"
//  INTERNALERROR_STREAMPROCESSFAIL = "InternalError.StreamProcessFail"
//  INTERNALERROR_STREAMPROCESSTIMEOUT = "InternalError.StreamProcessTimeOut"
//  INTERNALERROR_STREAMINGVOICEPKGTIMEOUT = "InternalError.StreamingvoicepkgTimeout"
//  INTERNALERROR_TIMEOUT = "InternalError.TimeOut"
//  INTERNALERROR_TOOLONGPACKAGE = "InternalError.TooLongPackage"
//  INTERNALERROR_VADNOTDETECTEDSPEAK = "InternalError.VadNotDetectedSpeak"
//  INTERNALERROR_VOICEMSGOVERSIZED = "InternalError.VoiceMsgOversized"
//  INTERNALERROR_VOICEMSGTOOSHORT = "InternalError.VoiceMsgTooShort"
//  INTERNALERROR_WORDLENGTHTOOLONG = "InternalError.WordLengthTooLong"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTHORIZEERROR = "InvalidParameter.AuthorizeError"
//  INVALIDPARAMETER_ERRORPHONEME = "InvalidParameter.ErrorPhoneme"
//  INVALIDPARAMETER_INITIALPARAMETERERROR = "InvalidParameter.InitialParameterError"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_SEQIDLIMITEXCEEDED = "InvalidParameter.SeqIdLimitExceeded"
//  INVALIDPARAMETER_VOICEMSGOVERSIZED = "InvalidParameter.VoiceMsgOversized"
//  INVALIDPARAMETER_WITHOUTREALNAME = "InvalidParameter.WithoutRealName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUDIODATASIZELIMITEXCEEDED = "InvalidParameterValue.AudioDataSizeLimitExceeded"
//  INVALIDPARAMETERVALUE_AUDIODECODEFAILED = "InvalidParameterValue.AudioDecodeFailed"
//  INVALIDPARAMETERVALUE_AUDIOLIMITEXCEEDED = "InvalidParameterValue.AudioLimitExceeded"
//  INVALIDPARAMETERVALUE_AUDIOSIZEMUSTBEEVEN = "InvalidParameterValue.AudioSizeMustBeEven"
//  INVALIDPARAMETERVALUE_BASEDECODEFAILED = "InvalidParameterValue.BASEDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDSEQID = "InvalidParameterValue.InvalidSeqId"
//  INVALIDPARAMETERVALUE_INVALIDWAVHEADER = "InvalidParameterValue.InvalidWAVHeader"
//  INVALIDPARAMETERVALUE_NODOCINLIST = "InvalidParameterValue.NoDocInList"
//  INVALIDPARAMETERVALUE_PARAMETERINVALID = "InvalidParameterValue.ParameterInvalid"
//  INVALIDPARAMETERVALUE_REFTEXTEMPTY = "InvalidParameterValue.RefTextEmpty"
//  INVALIDPARAMETERVALUE_REFTEXTGRAMMARERROR = "InvalidParameterValue.RefTextGrammarError"
//  INVALIDPARAMETERVALUE_REFTEXTLIMITEXCEEDED = "InvalidParameterValue.RefTextLimitExceeded"
//  INVALIDPARAMETERVALUE_REFTEXTOOV = "InvalidParameterValue.RefTextOOV"
//  INVALIDPARAMETERVALUE_REFTEXTPOLYPHONICLIMITEXCEEDED = "InvalidParameterValue.RefTextPolyphonicLimitExceeded"
//  INVALIDPARAMETERVALUE_REFTXTEMPTY = "InvalidParameterValue.RefTxtEmpty"
//  INVALIDPARAMETERVALUE_REFTXTTOOLANG = "InvalidParameterValue.RefTxtTooLang"
//  INVALIDPARAMETERVALUE_SENSITIVEWORDS = "InvalidParameterValue.SensitiveWords"
//  INVALIDPARAMETERVALUE_SESSIONIDINUSE = "InvalidParameterValue.SessionIdInUse"
//  INVALIDPARAMETERVALUE_SHARDNOSTARTWITHONE = "InvalidParameterValue.ShardNoStartWithOne"
//  INVALIDPARAMETERVALUE_STREAMINGVOICEPKGTIMEOUT = "InvalidParameterValue.StreamingvoicepkgTimeout"
//  INVALIDPARAMETERVALUE_VADNOTDETECTEDSPEAK = "InvalidParameterValue.VadNotDetectedSpeak"
//  INVALIDPARAMETERVALUE_VOICEFILETYPENOTFOUND = "InvalidParameterValue.VoiceFileTypeNotFound"
//  INVALIDPARAMETERVALUE_VOICELENGTHTOOLONG = "InvalidParameterValue.VoiceLengthTooLong"
//  INVALIDPARAMETERVALUE_WAVHEADERDECODEFAILED = "InvalidParameterValue.WAVHeaderDecodeFailed"
//  INVALIDPARAMETERVALUE_WORDLENGTHTOOLONG = "InvalidParameterValue.WordLengthTooLong"
//  LIMITEXCEEDED_CONCURRENCYLIMITEXCEEDED = "LimitExceeded.ConcurrencyLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT_SERVERTIMEOUT = "ResourceInsufficient.ServerTimeout"
//  RESOURCENOTFOUND_APPIDNOTFOUNT = "ResourceNotFound.AppidNotFount"
//  RESOURCENOTFOUND_INTERFACENOTFOUNT = "ResourceNotFound.InterfaceNotFount"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_AUTHORIZEERROR = "ResourceUnavailable.AuthorizeError"
//  RESOURCEUNAVAILABLE_CANNOTFINDSESSION = "ResourceUnavailable.CannotFindSession"
//  RESOURCEUNAVAILABLE_CONCURRENCYLIMIT = "ResourceUnavailable.ConcurrencyLimit"
//  RESOURCEUNAVAILABLE_INITSTREAMUNFINISHED = "ResourceUnavailable.InitStreamUnfinished"
//  RESOURCEUNAVAILABLE_LASTSEQUNFINISHED = "ResourceUnavailable.LastSeqUnfinished"
//  RESOURCEUNAVAILABLE_NOCONVERSATIONFOUND = "ResourceUnavailable.NoConversationFound"
//  RESOURCEUNAVAILABLE_NOINITBEFOREEVALUATION = "ResourceUnavailable.NoInitBeforeEvaluation"
func (c *Client) TransmitOralProcessWithInit(request *TransmitOralProcessWithInitRequest) (response *TransmitOralProcessWithInitResponse, err error) {
    return c.TransmitOralProcessWithInitWithContext(context.Background(), request)
}

// TransmitOralProcessWithInit
// 本接口可用于中英文发音评测。在使用本接口时需要注意：初始化并传输音频数据，分片传输时，尽量保证SeqId顺序传输（请确认SeqId由1开始）。音频源目前仅支持16k采样率16bit单声道编码方式，如有不一致可能导致评估不准确或失败。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  AUTHFAILURE_ACCOUNTUNAVAILABLE = "AuthFailure.AccountUnavailable"
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ERRORGETSESSION = "FailedOperation.ErrorGetSession"
//  FAILEDOPERATION_ERRORGETUSER = "FailedOperation.ErrorGetUser"
//  FAILEDOPERATION_EVALUATETIMEOUT = "FailedOperation.EvaluateTimeout"
//  FAILEDOPERATION_EVALUATEUNKNOWNERROR = "FailedOperation.EvaluateUnknownError"
//  FAILEDOPERATION_FAILEDGETENGINEIP = "FailedOperation.FailedGetEngineIP"
//  FAILEDOPERATION_FAILEDGETRESULT = "FailedOperation.FailedGetResult"
//  FAILEDOPERATION_FAILEDGETSESSION = "FailedOperation.FailedGetSession"
//  FAILEDOPERATION_FAILEDGETSESSIONSEQID = "FailedOperation.FailedGetSessionSeqID"
//  FAILEDOPERATION_FAILEDGETUSER = "FailedOperation.FailedGetUser"
//  FAILEDOPERATION_FAILEDINIT = "FailedOperation.FailedInit"
//  FAILEDOPERATION_FAILEDSETRESULT = "FailedOperation.FailedSetResult"
//  FAILEDOPERATION_FAILEDSETSESSION = "FailedOperation.FailedSetSession"
//  FAILEDOPERATION_FAILEDSETSESSIONSEQID = "FailedOperation.FailedSetSessionSeqID"
//  FAILEDOPERATION_FAILEDSETUSER = "FailedOperation.FailedSetUser"
//  FAILEDOPERATION_INTERNALSERVERERROR = "FailedOperation.InternalServerError"
//  FAILEDOPERATION_INVALIDPARAMETERVALUE = "FailedOperation.InvalidParameterValue"
//  FAILEDOPERATION_JSONCODECERROR = "FailedOperation.JsonCodecError"
//  FAILEDOPERATION_NEEDINITBEFOREEVALUATION = "FailedOperation.NeedInitBeforeEvaluation"
//  FAILEDOPERATION_PASTSEQIDLOSE = "FailedOperation.PastSeqIdLose"
//  FAILEDOPERATION_RESULTEXPIRED = "FailedOperation.ResultExpired"
//  FAILEDOPERATION_SEQIDEXPIRED = "FailedOperation.SeqIdExpired"
//  FAILEDOPERATION_SERVEROVERLOAD = "FailedOperation.ServerOverload"
//  FAILEDOPERATION_SERVICETIMEOUT = "FailedOperation.ServiceTimeout"
//  FAILEDOPERATION_SESSIONEXPIRED = "FailedOperation.SessionExpired"
//  FAILEDOPERATION_WAITPASTSEQIDTIMEOUT = "FailedOperation.WaitPastSeqIdTimeout"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_AUDIOPROCESSINGFAILED = "InternalError.AudioProcessingFailed"
//  INTERNALERROR_BASE64DECODEFAILED = "InternalError.BASE64DecodeFailed"
//  INTERNALERROR_CANNOTFINDSESSION = "InternalError.CannotFindSession"
//  INTERNALERROR_FAILTODECODEVOICE = "InternalError.FailToDecodeVoice"
//  INTERNALERROR_ILEGALSERVERRESPONSE = "InternalError.IlegalServerResponse"
//  INTERNALERROR_INITSTREAMNOTSUPPORT = "InternalError.InitStreamNotSupport"
//  INTERNALERROR_INITSTREAMUNFINISHED = "InternalError.InitStreamUnfinished"
//  INTERNALERROR_INITIALPARAMETERERROR = "InternalError.InitialParameterError"
//  INTERNALERROR_INVALIDSEQID = "InternalError.InvalidSeqId"
//  INTERNALERROR_INVALIDWAVHEADER = "InternalError.InvalidWAVHeader"
//  INTERNALERROR_LASTSEQUNFINISHED = "InternalError.LastSeqUnfinished"
//  INTERNALERROR_MP3DECODEFAILED = "InternalError.MP3DecodeFailed"
//  INTERNALERROR_NEEDTOINIT = "InternalError.NeedToInit"
//  INTERNALERROR_NOCONVERSATIONFOUND = "InternalError.NoConversationFound"
//  INTERNALERROR_NODOCINLIST = "InternalError.NoDocInList"
//  INTERNALERROR_NOERROR = "InternalError.NoError"
//  INTERNALERROR_NOINITBEFOREEVALUATION = "InternalError.NoInitBeforeEvaluation"
//  INTERNALERROR_REFTXTEMPTY = "InternalError.RefTxtEmpty"
//  INTERNALERROR_REFTXTOOV = "InternalError.RefTxtOov"
//  INTERNALERROR_REFTXTTOOLANG = "InternalError.RefTxtTooLang"
//  INTERNALERROR_SERVERINTERNALERROR = "InternalError.ServerInternalError"
//  INTERNALERROR_SERVEROVERLOAD = "InternalError.ServerOverload"
//  INTERNALERROR_SERVICETIMEOUT = "InternalError.ServiceTimeout"
//  INTERNALERROR_SHARDNOSTARTWITHONE = "InternalError.ShardNoStartWithOne"
//  INTERNALERROR_STREAMPROCESSFAIL = "InternalError.StreamProcessFail"
//  INTERNALERROR_STREAMPROCESSTIMEOUT = "InternalError.StreamProcessTimeOut"
//  INTERNALERROR_STREAMINGVOICEPKGTIMEOUT = "InternalError.StreamingvoicepkgTimeout"
//  INTERNALERROR_TIMEOUT = "InternalError.TimeOut"
//  INTERNALERROR_TOOLONGPACKAGE = "InternalError.TooLongPackage"
//  INTERNALERROR_VADNOTDETECTEDSPEAK = "InternalError.VadNotDetectedSpeak"
//  INTERNALERROR_VOICEMSGOVERSIZED = "InternalError.VoiceMsgOversized"
//  INTERNALERROR_VOICEMSGTOOSHORT = "InternalError.VoiceMsgTooShort"
//  INTERNALERROR_WORDLENGTHTOOLONG = "InternalError.WordLengthTooLong"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTHORIZEERROR = "InvalidParameter.AuthorizeError"
//  INVALIDPARAMETER_ERRORPHONEME = "InvalidParameter.ErrorPhoneme"
//  INVALIDPARAMETER_INITIALPARAMETERERROR = "InvalidParameter.InitialParameterError"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_SEQIDLIMITEXCEEDED = "InvalidParameter.SeqIdLimitExceeded"
//  INVALIDPARAMETER_VOICEMSGOVERSIZED = "InvalidParameter.VoiceMsgOversized"
//  INVALIDPARAMETER_WITHOUTREALNAME = "InvalidParameter.WithoutRealName"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_AUDIODATASIZELIMITEXCEEDED = "InvalidParameterValue.AudioDataSizeLimitExceeded"
//  INVALIDPARAMETERVALUE_AUDIODECODEFAILED = "InvalidParameterValue.AudioDecodeFailed"
//  INVALIDPARAMETERVALUE_AUDIOLIMITEXCEEDED = "InvalidParameterValue.AudioLimitExceeded"
//  INVALIDPARAMETERVALUE_AUDIOSIZEMUSTBEEVEN = "InvalidParameterValue.AudioSizeMustBeEven"
//  INVALIDPARAMETERVALUE_BASEDECODEFAILED = "InvalidParameterValue.BASEDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDSEQID = "InvalidParameterValue.InvalidSeqId"
//  INVALIDPARAMETERVALUE_INVALIDWAVHEADER = "InvalidParameterValue.InvalidWAVHeader"
//  INVALIDPARAMETERVALUE_NODOCINLIST = "InvalidParameterValue.NoDocInList"
//  INVALIDPARAMETERVALUE_PARAMETERINVALID = "InvalidParameterValue.ParameterInvalid"
//  INVALIDPARAMETERVALUE_REFTEXTEMPTY = "InvalidParameterValue.RefTextEmpty"
//  INVALIDPARAMETERVALUE_REFTEXTGRAMMARERROR = "InvalidParameterValue.RefTextGrammarError"
//  INVALIDPARAMETERVALUE_REFTEXTLIMITEXCEEDED = "InvalidParameterValue.RefTextLimitExceeded"
//  INVALIDPARAMETERVALUE_REFTEXTOOV = "InvalidParameterValue.RefTextOOV"
//  INVALIDPARAMETERVALUE_REFTEXTPOLYPHONICLIMITEXCEEDED = "InvalidParameterValue.RefTextPolyphonicLimitExceeded"
//  INVALIDPARAMETERVALUE_REFTXTEMPTY = "InvalidParameterValue.RefTxtEmpty"
//  INVALIDPARAMETERVALUE_REFTXTTOOLANG = "InvalidParameterValue.RefTxtTooLang"
//  INVALIDPARAMETERVALUE_SENSITIVEWORDS = "InvalidParameterValue.SensitiveWords"
//  INVALIDPARAMETERVALUE_SESSIONIDINUSE = "InvalidParameterValue.SessionIdInUse"
//  INVALIDPARAMETERVALUE_SHARDNOSTARTWITHONE = "InvalidParameterValue.ShardNoStartWithOne"
//  INVALIDPARAMETERVALUE_STREAMINGVOICEPKGTIMEOUT = "InvalidParameterValue.StreamingvoicepkgTimeout"
//  INVALIDPARAMETERVALUE_VADNOTDETECTEDSPEAK = "InvalidParameterValue.VadNotDetectedSpeak"
//  INVALIDPARAMETERVALUE_VOICEFILETYPENOTFOUND = "InvalidParameterValue.VoiceFileTypeNotFound"
//  INVALIDPARAMETERVALUE_VOICELENGTHTOOLONG = "InvalidParameterValue.VoiceLengthTooLong"
//  INVALIDPARAMETERVALUE_WAVHEADERDECODEFAILED = "InvalidParameterValue.WAVHeaderDecodeFailed"
//  INVALIDPARAMETERVALUE_WORDLENGTHTOOLONG = "InvalidParameterValue.WordLengthTooLong"
//  LIMITEXCEEDED_CONCURRENCYLIMITEXCEEDED = "LimitExceeded.ConcurrencyLimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT_SERVERTIMEOUT = "ResourceInsufficient.ServerTimeout"
//  RESOURCENOTFOUND_APPIDNOTFOUNT = "ResourceNotFound.AppidNotFount"
//  RESOURCENOTFOUND_INTERFACENOTFOUNT = "ResourceNotFound.InterfaceNotFount"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_AUTHORIZEERROR = "ResourceUnavailable.AuthorizeError"
//  RESOURCEUNAVAILABLE_CANNOTFINDSESSION = "ResourceUnavailable.CannotFindSession"
//  RESOURCEUNAVAILABLE_CONCURRENCYLIMIT = "ResourceUnavailable.ConcurrencyLimit"
//  RESOURCEUNAVAILABLE_INITSTREAMUNFINISHED = "ResourceUnavailable.InitStreamUnfinished"
//  RESOURCEUNAVAILABLE_LASTSEQUNFINISHED = "ResourceUnavailable.LastSeqUnfinished"
//  RESOURCEUNAVAILABLE_NOCONVERSATIONFOUND = "ResourceUnavailable.NoConversationFound"
//  RESOURCEUNAVAILABLE_NOINITBEFOREEVALUATION = "ResourceUnavailable.NoInitBeforeEvaluation"
func (c *Client) TransmitOralProcessWithInitWithContext(ctx context.Context, request *TransmitOralProcessWithInitRequest) (response *TransmitOralProcessWithInitResponse, err error) {
    if request == nil {
        request = NewTransmitOralProcessWithInitRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TransmitOralProcessWithInit require credential")
    }

    request.SetContext(ctx)
    
    response = NewTransmitOralProcessWithInitResponse()
    err = c.Send(request, response)
    return
}
