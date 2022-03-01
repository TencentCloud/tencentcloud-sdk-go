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

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 账号未开通口语评测服务或账号已欠费隔离，请开通服务或检查账号状态。
	AUTHFAILURE_ACCOUNTUNAVAILABLE = "AuthFailure.AccountUnavailable"

	// 鉴权失败。
	AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 会话缓存保存失败，请重新初始化。
	FAILEDOPERATION_ERRORGETSESSION = "FailedOperation.ErrorGetSession"

	// 获取用户信息失败。
	FAILEDOPERATION_ERRORGETUSER = "FailedOperation.ErrorGetUser"

	// 评测时间超出限制，请检查音频时间是否过长后重试。
	FAILEDOPERATION_EVALUATETIMEOUT = "FailedOperation.EvaluateTimeout"

	// 引擎未知错误，请检查一下RefText是否正常后重试。
	FAILEDOPERATION_EVALUATEUNKNOWNERROR = "FailedOperation.EvaluateUnknownError"

	// 获取评测引擎IP失败，请稍后重试。
	FAILEDOPERATION_FAILEDGETENGINEIP = "FailedOperation.FailedGetEngineIP"

	// 结果缓存获取失败，请稍后重试。
	FAILEDOPERATION_FAILEDGETRESULT = "FailedOperation.FailedGetResult"

	// 会话缓存获取失败，请稍后重试。
	FAILEDOPERATION_FAILEDGETSESSION = "FailedOperation.FailedGetSession"

	// 会话分片序号缓存获取失败，请稍后重试。
	FAILEDOPERATION_FAILEDGETSESSIONSEQID = "FailedOperation.FailedGetSessionSeqID"

	// 用户信息缓存获取失败，请稍后重试。
	FAILEDOPERATION_FAILEDGETUSER = "FailedOperation.FailedGetUser"

	// 请求初始化失败，请检查参数后重新初始化。
	FAILEDOPERATION_FAILEDINIT = "FailedOperation.FailedInit"

	// 结果缓存保存失败，请稍后重试。
	FAILEDOPERATION_FAILEDSETRESULT = "FailedOperation.FailedSetResult"

	// 会话缓存保存失败，请重新初始化。
	FAILEDOPERATION_FAILEDSETSESSION = "FailedOperation.FailedSetSession"

	// 会话分片序号缓存保存失败，请重新初始化。
	FAILEDOPERATION_FAILEDSETSESSIONSEQID = "FailedOperation.FailedSetSessionSeqID"

	// 用户信息缓存保存失败，请稍后重试。
	FAILEDOPERATION_FAILEDSETUSER = "FailedOperation.FailedSetUser"

	// 服务内部错误，请稍后重试或联系我们。
	FAILEDOPERATION_INTERNALSERVERERROR = "FailedOperation.InternalServerError"

	// 引擎参数错误，请稍后重试。
	FAILEDOPERATION_INVALIDPARAMETERVALUE = "FailedOperation.InvalidParameterValue"

	// Json编解码失败，请稍后重试。
	FAILEDOPERATION_JSONCODECERROR = "FailedOperation.JsonCodecError"

	// 引擎评估之前没有初始化，请重新初始化成功之后重新传输数据。
	FAILEDOPERATION_NEEDINITBEFOREEVALUATION = "FailedOperation.NeedInitBeforeEvaluation"

	// 前序分片缺失，请重新补发前序分片。
	FAILEDOPERATION_PASTSEQIDLOSE = "FailedOperation.PastSeqIdLose"

	// 结果缓存已过期，请重新初始化成功之后重新传输数据。
	FAILEDOPERATION_RESULTEXPIRED = "FailedOperation.ResultExpired"

	// 分片序号缓存已过期，请重新初始化成功之后重新传输数据。
	FAILEDOPERATION_SEQIDEXPIRED = "FailedOperation.SeqIdExpired"

	// 引擎服务器过载，请稍后重试。
	FAILEDOPERATION_SERVEROVERLOAD = "FailedOperation.ServerOverload"

	// 评测超时，请通过轮询查询评测结果，后续请使用分片传输或减少单次传输音频时长。
	FAILEDOPERATION_SERVICETIMEOUT = "FailedOperation.ServiceTimeout"

	// 会话缓存已过期，请重新初始化成功之后重新传输数据。
	FAILEDOPERATION_SESSIONEXPIRED = "FailedOperation.SessionExpired"

	// 引擎等待前序分片超时，请重新补发前序分片。
	FAILEDOPERATION_WAITPASTSEQIDTIMEOUT = "FailedOperation.WaitPastSeqIdTimeout"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 音频处理出错。
	INTERNALERROR_AUDIOPROCESSINGFAILED = "InternalError.AudioProcessingFailed"

	// 服务未开通或已欠费。
	INTERNALERROR_AUTHORIZEERROR = "InternalError.AuthorizeError"

	// BASE64解码错误。
	INTERNALERROR_BASE64DECODEFAILED = "InternalError.BASE64DecodeFailed"

	// 评估之前没有初始化或已过期。
	INTERNALERROR_CANNOTFINDSESSION = "InternalError.CannotFindSession"

	// 语音解码失败。
	INTERNALERROR_FAILTODECODEVOICE = "InternalError.FailToDecodeVoice"

	// 服务器应答非法 。
	INTERNALERROR_ILEGALSERVERRESPONSE = "InternalError.IlegalServerResponse"

	// 该接口不支持init_stream。
	INTERNALERROR_INITSTREAMNOTSUPPORT = "InternalError.InitStreamNotSupport"

	// 初始化请求未完成，请稍后重试。
	INTERNALERROR_INITSTREAMUNFINISHED = "InternalError.InitStreamUnfinished"

	// 初始化参数错误  。
	INTERNALERROR_INITIALPARAMETERERROR = "InternalError.InitialParameterError"

	// 分片序号错误。
	INTERNALERROR_INVALIDSEQID = "InternalError.InvalidSeqId"

	// WAV头部格式非法或不在同一分片内。
	INTERNALERROR_INVALIDWAVHEADER = "InternalError.InvalidWAVHeader"

	// 前一个分片未处理完，请稍后重试。
	INTERNALERROR_LASTSEQUNFINISHED = "InternalError.LastSeqUnfinished"

	// MP3转码发生错误。
	INTERNALERROR_MP3DECODEFAILED = "InternalError.MP3DecodeFailed"

	// 进行评估之前没有进行初始化。
	INTERNALERROR_NEEDTOINIT = "InternalError.NeedToInit"

	// 使用的会话没有找到或已经被释放。
	INTERNALERROR_NOCONVERSATIONFOUND = "InternalError.NoConversationFound"

	// 表单中没有文件。
	INTERNALERROR_NODOCINLIST = "InternalError.NoDocInList"

	// 没有错误。
	INTERNALERROR_NOERROR = "InternalError.NoError"

	// 评估之前没有初始化。
	INTERNALERROR_NOINITBEFOREEVALUATION = "InternalError.NoInitBeforeEvaluation"

	// 检测到不支持的字符在输入文本。
	INTERNALERROR_REFTXTEMPTY = "InternalError.RefTxtEmpty"

	// 检测到不支持的字符在输入文本。
	INTERNALERROR_REFTXTOOV = "InternalError.RefTxtOov"

	// 输入文本太长。
	INTERNALERROR_REFTXTTOOLANG = "InternalError.RefTxtTooLang"

	// 服务器内部错误。
	INTERNALERROR_SERVERINTERNALERROR = "InternalError.ServerInternalError"

	// 服务器过载。
	INTERNALERROR_SERVEROVERLOAD = "InternalError.ServerOverload"

	// 服务超时。
	INTERNALERROR_SERVICETIMEOUT = "InternalError.ServiceTimeout"

	// 分片序号错误应该从1开始。
	INTERNALERROR_SHARDNOSTARTWITHONE = "InternalError.ShardNoStartWithOne"

	// 流式模式数据包处理过程中间失败。
	INTERNALERROR_STREAMPROCESSFAIL = "InternalError.StreamProcessFail"

	// 流式模式数据包处理超时。
	INTERNALERROR_STREAMPROCESSTIMEOUT = "InternalError.StreamProcessTimeOut"

	// 流式语音包超时。
	INTERNALERROR_STREAMINGVOICEPKGTIMEOUT = "InternalError.StreamingvoicepkgTimeout"

	// 获得结果超时。
	INTERNALERROR_TIMEOUT = "InternalError.TimeOut"

	// 语音数据包长度超过 1MB。
	INTERNALERROR_TOOLONGPACKAGE = "InternalError.TooLongPackage"

	// 没有检测到语音。
	INTERNALERROR_VADNOTDETECTEDSPEAK = "InternalError.VadNotDetectedSpeak"

	// 语音数据大于1MB。
	INTERNALERROR_VOICEMSGOVERSIZED = "InternalError.VoiceMsgOversized"

	// 语音时长太短 。
	INTERNALERROR_VOICEMSGTOOSHORT = "InternalError.VoiceMsgTooShort"

	// 文本单词超过限制 。
	INTERNALERROR_WORDLENGTHTOOLONG = "InternalError.WordLengthTooLong"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 服务未开通或已欠费。
	INVALIDPARAMETER_AUTHORIZEERROR = "InvalidParameter.AuthorizeError"

	// 请求参数RefText的音素Json解码失败，请参考API文档使用标准的Json格式。
	INVALIDPARAMETER_ERRORPHONEME = "InvalidParameter.ErrorPhoneme"

	// 初始化参数错误。
	INVALIDPARAMETER_INITIALPARAMETERERROR = "InvalidParameter.InitialParameterError"

	// 请求参数Action不合法，请参考API文档检查参数Action的有效性。
	INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"

	// 请求参数不合法，请参考API文档检查参数的有效性。
	INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"

	// 请求参数SeqId超过最大值限制，请参考API文档检查参数SeqId是否小于3000。
	INVALIDPARAMETER_SEQIDLIMITEXCEEDED = "InvalidParameter.SeqIdLimitExceeded"

	// 语音数据大于1MB。
	INVALIDPARAMETER_VOICEMSGOVERSIZED = "InvalidParameter.VoiceMsgOversized"

	// 用户未实名制认证。
	INVALIDPARAMETER_WITHOUTREALNAME = "InvalidParameter.WithoutRealName"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 输入分片音频大小超过最大限制，请调整分片大小后重新传输数据。
	INVALIDPARAMETERVALUE_AUDIODATASIZELIMITEXCEEDED = "InvalidParameterValue.AudioDataSizeLimitExceeded"

	// 音频数据解码失败，请参考API文档中音频要求检查音频数据格式设置是否正确后重新传输数据。
	INVALIDPARAMETERVALUE_AUDIODECODEFAILED = "InvalidParameterValue.AudioDecodeFailed"

	// 输入音频时长超过限制，请结束本次评测，后续请根据评测模式设置音频时长限制。
	INVALIDPARAMETERVALUE_AUDIOLIMITEXCEEDED = "InvalidParameterValue.AudioLimitExceeded"

	// 输入音频异常，音频数据指针或音频⻓度必须为偶数，请参考API文档检查音频数据是否正确后重新传输数据。
	INVALIDPARAMETERVALUE_AUDIOSIZEMUSTBEEVEN = "InvalidParameterValue.AudioSizeMustBeEven"

	// BASE64解码错误。
	INVALIDPARAMETERVALUE_BASEDECODEFAILED = "InvalidParameterValue.BASEDecodeFailed"

	// 分片序号错误。
	INVALIDPARAMETERVALUE_INVALIDSEQID = "InvalidParameterValue.InvalidSeqId"

	// WAV头部格式非法或不在同一分片内。
	INVALIDPARAMETERVALUE_INVALIDWAVHEADER = "InvalidParameterValue.InvalidWAVHeader"

	// 表单中没有文件。
	INVALIDPARAMETERVALUE_NODOCINLIST = "InvalidParameterValue.NoDocInList"

	// 参数值无效，请检查ScoreCoeff参数输入是否在限制内。
	INVALIDPARAMETERVALUE_PARAMETERINVALID = "InvalidParameterValue.ParameterInvalid"

	// 请求参数RefText无效或参考文本为空，请检查RefText是否为空。
	INVALIDPARAMETERVALUE_REFTEXTEMPTY = "InvalidParameterValue.RefTextEmpty"

	// 请求参数RefText语法错误，请参考API文档检查文本格式，尤其是指定发音格式是否正确。
	INVALIDPARAMETERVALUE_REFTEXTGRAMMARERROR = "InvalidParameterValue.RefTextGrammarError"

	// 请求参数RefText的字数超过最大限制，请根据评测模式调整字数后重新初始化。
	INVALIDPARAMETERVALUE_REFTEXTLIMITEXCEEDED = "InvalidParameterValue.RefTextLimitExceeded"

	// 请求参数RefText包含OOV词汇，请使用指定发音或联系我们处理。
	INVALIDPARAMETERVALUE_REFTEXTOOV = "InvalidParameterValue.RefTextOOV"

	// 请检查参考文本中是否包含大量多音字，可通过发音描述块指定标准发音解决。
	INVALIDPARAMETERVALUE_REFTEXTPOLYPHONICLIMITEXCEEDED = "InvalidParameterValue.RefTextPolyphonicLimitExceeded"

	// 输入文本为空。
	INVALIDPARAMETERVALUE_REFTXTEMPTY = "InvalidParameterValue.RefTxtEmpty"

	// 输入文本太长。
	INVALIDPARAMETERVALUE_REFTXTTOOLANG = "InvalidParameterValue.RefTxtTooLang"

	// 请求内容包含违禁词汇，请检查后重试。
	INVALIDPARAMETERVALUE_SENSITIVEWORDS = "InvalidParameterValue.SensitiveWords"

	// SessionId已存在，建议使用uuid作为SessionId重新初始化。
	INVALIDPARAMETERVALUE_SESSIONIDINUSE = "InvalidParameterValue.SessionIdInUse"

	// 分片序号错误应该从1开始。
	INVALIDPARAMETERVALUE_SHARDNOSTARTWITHONE = "InvalidParameterValue.ShardNoStartWithOne"

	// 流式语音包超时。
	INVALIDPARAMETERVALUE_STREAMINGVOICEPKGTIMEOUT = "InvalidParameterValue.StreamingvoicepkgTimeout"

	// 没有检测到语音。
	INVALIDPARAMETERVALUE_VADNOTDETECTEDSPEAK = "InvalidParameterValue.VadNotDetectedSpeak"

	// 语音文件格式参数VoiceFileType取值错误，请参考API文档检查语音文件格式VoiceFileType是否正确后重新传输数据。
	INVALIDPARAMETERVALUE_VOICEFILETYPENOTFOUND = "InvalidParameterValue.VoiceFileTypeNotFound"

	// 音频超过长度限制，要求音频大小不能超过3Mb。
	INVALIDPARAMETERVALUE_VOICELENGTHTOOLONG = "InvalidParameterValue.VoiceLengthTooLong"

	// WAV格式的音频数据第一个分片的数据长度小于44，头部数据不合法，请检查后重试。
	INVALIDPARAMETERVALUE_WAVHEADERDECODEFAILED = "InvalidParameterValue.WAVHeaderDecodeFailed"

	// 文本单词超过限制。
	INVALIDPARAMETERVALUE_WORDLENGTHTOOLONG = "InvalidParameterValue.WordLengthTooLong"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 请求并发数超过配额限制，请减少并发数或联系我们调大并发限额。
	LIMITEXCEEDED_CONCURRENCYLIMITEXCEEDED = "LimitExceeded.ConcurrencyLimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 评测超时，请检查语音数据大小。
	RESOURCEINSUFFICIENT_SERVERTIMEOUT = "ResourceInsufficient.ServerTimeout"

	// appid不存在。
	RESOURCENOTFOUND_APPIDNOTFOUNT = "ResourceNotFound.AppidNotFount"

	// 接口不存在。
	RESOURCENOTFOUND_INTERFACENOTFOUNT = "ResourceNotFound.InterfaceNotFount"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 服务未开通或已欠费。
	RESOURCEUNAVAILABLE_AUTHORIZEERROR = "ResourceUnavailable.AuthorizeError"

	// 评估之前没有初始化或已过期。
	RESOURCEUNAVAILABLE_CANNOTFINDSESSION = "ResourceUnavailable.CannotFindSession"

	// 使用并发超出限制。
	RESOURCEUNAVAILABLE_CONCURRENCYLIMIT = "ResourceUnavailable.ConcurrencyLimit"

	// 该接口不支持init_stream。
	RESOURCEUNAVAILABLE_INITSTREAMNOTSUPPORT = "ResourceUnavailable.InitStreamNotSupport"

	// 初始化请求未完成，请稍后重试。
	RESOURCEUNAVAILABLE_INITSTREAMUNFINISHED = "ResourceUnavailable.InitStreamUnfinished"

	// 前一个分片未处理完，请稍后重试。
	RESOURCEUNAVAILABLE_LASTSEQUNFINISHED = "ResourceUnavailable.LastSeqUnfinished"

	// 使用的会话没有找到或已经被释放。
	RESOURCEUNAVAILABLE_NOCONVERSATIONFOUND = "ResourceUnavailable.NoConversationFound"

	// 评估之前没有初始化。
	RESOURCEUNAVAILABLE_NOINITBEFOREEVALUATION = "ResourceUnavailable.NoInitBeforeEvaluation"
)
