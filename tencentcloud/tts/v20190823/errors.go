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

package v20190823

const (
	// 此产品的特有错误码

	// 授权无效。
	AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 任务不存在。
	FAILEDOPERATION_NOSUCHTASK = "FailedOperation.NoSuchTask"

	// 路由错误。
	INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"

	// 负载限流。
	INTERNALERROR_EXCEEDMAXLIMIT = "InternalError.ExceedMaxLimit"

	// 数据库存取失败。
	INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"

	// Redis存储失败。
	INTERNALERROR_FAILACCESSREDIS = "InternalError.FailAccessRedis"

	// 内部错误。
	INTERNALERROR_INTERNALERROR = "InternalError.InternalError"

	// 请求文本含有非法字符。
	INVALIDPARAMETER_INVALIDTEXT = "InvalidParameter.InvalidText"

	// status 不合法。
	INVALIDPARAMETER_STATUS = "InvalidParameter.Status"

	// AppId非法，请参考AppId参数说明。
	INVALIDPARAMETERVALUE_APPID = "InvalidParameterValue.AppId"

	// APPID未注册，请在语音合成主页   https://console.cloud.tencent.com/tts  开通使用。
	INVALIDPARAMETERVALUE_APPIDNOTREGISTERED = "InvalidParameterValue.AppIdNotRegistered"

	// CallbackUrl非法或不可访问。
	INVALIDPARAMETERVALUE_CALLBACKURL = "InvalidParameterValue.CallbackUrl"

	// Codec非法，请参考Codec参数说明。
	INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"

	// ssml的say-as标签属性为cardinal、currency、address时，数字部分非有效常数，仅允许包含数字、“,”、“.”和空格。
	INVALIDPARAMETERVALUE_ERRORCARDINALFORMAT = "InvalidParameterValue.ErrorCardinalFormat"

	// 请求文本含有非法字符，或请求文本没有有效字符。
	INVALIDPARAMETERVALUE_INVALIDTEXT = "InvalidParameterValue.InvalidText"

	// 缺少参数。
	INVALIDPARAMETERVALUE_MISSPARAMETERS = "InvalidParameterValue.MissParameters"

	// ModelType非法。
	INVALIDPARAMETERVALUE_MODELTYPE = "InvalidParameterValue.ModelType"

	// PrimaryLanguage非法，请参考PrimaryLanguage参数说明。
	INVALIDPARAMETERVALUE_PRIMARYLANGUAGE = "InvalidParameterValue.PrimaryLanguage"

	// SampleRate非法，请参考SampleRate参数说明。
	INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"

	// SessionId非法，请参考Volume参数说明。
	INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"

	// Speed非法，请参考Speed参数说明。
	INVALIDPARAMETERVALUE_SPEED = "InvalidParameterValue.Speed"

	// 文本缺失。
	INVALIDPARAMETERVALUE_TEXT = "InvalidParameterValue.Text"

	// Text为空。
	INVALIDPARAMETERVALUE_TEXTEMPTY = "InvalidParameterValue.TextEmpty"

	// 文本不是 UTF8 格式。
	INVALIDPARAMETERVALUE_TEXTNOTUTF8 = "InvalidParameterValue.TextNotUtf8"

	// Text参数SSML语法错误，请参考SSML文档说明。
	INVALIDPARAMETERVALUE_TEXTSSMLPARSEERROR = "InvalidParameterValue.TextSsmlParseError"

	// 合成文本字符过长。
	INVALIDPARAMETERVALUE_TEXTTOOLONG = "InvalidParameterValue.TextTooLong"

	// Type 类型非法。
	INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"

	// VoiceType非法，请参考VoiceType参数说明。
	INVALIDPARAMETERVALUE_VOICETYPE = "InvalidParameterValue.VoiceType"

	// Volume非法，请参考Volume参数说明。
	INVALIDPARAMETERVALUE_VOLUME = "InvalidParameterValue.Volume"

	// 请求超过限制频率。
	LIMITEXCEEDED_ACCESSLIMIT = "LimitExceeded.AccessLimit"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 欠费。
	UNSUPPORTEDOPERATION_ACCOUNTARREARS = "UnsupportedOperation.AccountArrears"

	// 鉴权已过期。
	UNSUPPORTEDOPERATION_AUTHORIZATIONEXPIRED = "UnsupportedOperation.AuthorizationExpired"

	// 鉴权失败。
	UNSUPPORTEDOPERATION_AUTHORIZATIONFAILED = "UnsupportedOperation.AuthorizationFailed"

	// 服务禁止使用。
	UNSUPPORTEDOPERATION_FORBIDDENUSE = "UnsupportedOperation.ForbiddenUse"

	// 没有余额。
	UNSUPPORTEDOPERATION_NOBANLANCE = "UnsupportedOperation.NoBanlance"

	// 客户免费额度已用完。
	UNSUPPORTEDOPERATION_NOFREEACCOUNT = "UnsupportedOperation.NoFreeAccount"

	// 服务器已打开。
	UNSUPPORTEDOPERATION_SERVERALREADYOPEN = "UnsupportedOperation.ServerAlreadyOpen"

	// 服务已销毁。
	UNSUPPORTEDOPERATION_SERVERDESTORYED = "UnsupportedOperation.ServerDestoryed"

	// 服务未开通使用。
	UNSUPPORTEDOPERATION_SERVERNOTOPEN = "UnsupportedOperation.ServerNotOpen"

	// 服务已停止使用。
	UNSUPPORTEDOPERATION_SERVERSTOPPED = "UnsupportedOperation.ServerStopped"

	// 文本过长，请参考请求参数Text的说明。
	UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
)
