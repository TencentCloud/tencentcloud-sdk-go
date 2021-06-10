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

	// 鉴权失败。
	AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"

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

	// 初始化参数错误。
	INVALIDPARAMETER_INITIALPARAMETERERROR = "InvalidParameter.InitialParameterError"

	// 语音数据大于1MB。
	INVALIDPARAMETER_VOICEMSGOVERSIZED = "InvalidParameter.VoiceMsgOversized"

	// 用户未实名制认证。
	INVALIDPARAMETER_WITHOUTREALNAME = "InvalidParameter.WithoutRealName"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// BASE64解码错误。
	INVALIDPARAMETERVALUE_BASEDECODEFAILED = "InvalidParameterValue.BASEDecodeFailed"

	// 分片序号错误。
	INVALIDPARAMETERVALUE_INVALIDSEQID = "InvalidParameterValue.InvalidSeqId"

	// WAV头部格式非法或不在同一分片内。
	INVALIDPARAMETERVALUE_INVALIDWAVHEADER = "InvalidParameterValue.InvalidWAVHeader"

	// 表单中没有文件。
	INVALIDPARAMETERVALUE_NODOCINLIST = "InvalidParameterValue.NoDocInList"

	// 输入文本为空。
	INVALIDPARAMETERVALUE_REFTXTEMPTY = "InvalidParameterValue.RefTxtEmpty"

	// 输入文本太长。
	INVALIDPARAMETERVALUE_REFTXTTOOLANG = "InvalidParameterValue.RefTxtTooLang"

	// 分片序号错误应该从1开始。
	INVALIDPARAMETERVALUE_SHARDNOSTARTWITHONE = "InvalidParameterValue.ShardNoStartWithOne"

	// 流式语音包超时。
	INVALIDPARAMETERVALUE_STREAMINGVOICEPKGTIMEOUT = "InvalidParameterValue.StreamingvoicepkgTimeout"

	// 没有检测到语音。
	INVALIDPARAMETERVALUE_VADNOTDETECTEDSPEAK = "InvalidParameterValue.VadNotDetectedSpeak"

	// 文本单词超过限制。
	INVALIDPARAMETERVALUE_WORDLENGTHTOOLONG = "InvalidParameterValue.WordLengthTooLong"

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
