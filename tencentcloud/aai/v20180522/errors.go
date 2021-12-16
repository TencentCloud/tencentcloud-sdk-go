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

package v20180522

const (
	// 此产品的特有错误码

	// 账号因为欠费停止服务，请在腾讯云账户充值。
	FAILEDOPERATION_SERVICEISOLATE = "FailedOperation.ServiceIsolate"

	// 账号本月免费额度已用完。
	FAILEDOPERATION_USERHASNOFREEAMOUNT = "FailedOperation.UserHasNoFreeAmount"

	// 服务未开通，请在腾讯云官网语音识别控制台开通服务。
	FAILEDOPERATION_USERNOTREGISTERED = "FailedOperation.UserNotRegistered"

	// 内部错误
	INTERNALERROR = "InternalError"

	// 初始化配置失败。
	INTERNALERROR_ERRORCONFIGURE = "InternalError.ErrorConfigure"

	// 创建日志失败。
	INTERNALERROR_ERRORCREATELOG = "InternalError.ErrorCreateLog"

	// 下载音频文件失败。
	INTERNALERROR_ERRORDOWNFILE = "InternalError.ErrorDownFile"

	// 新建数组失败。
	INTERNALERROR_ERRORFAILNEWPREQUEST = "InternalError.ErrorFailNewprequest"

	// 写入数据库失败。
	INTERNALERROR_ERRORFAILWRITETODB = "InternalError.ErrorFailWritetodb"

	// 文件无法打开。
	INTERNALERROR_ERRORFILECANNOTOPEN = "InternalError.ErrorFileCannotopen"

	// 获取路由失败。
	INTERNALERROR_ERRORGETROUTE = "InternalError.ErrorGetRoute"

	// 创建日志路径失败。
	INTERNALERROR_ERRORMAKELOGPATH = "InternalError.ErrorMakeLogpath"

	// 识别失败。
	INTERNALERROR_ERRORRECOGNIZE = "InternalError.ErrorRecognize"

	// 机器翻译失败。
	INTERNALERROR_ERRORTRANSLATE = "InternalError.ErrorTranslate"

	// 文本转换失败，请检查文本格式。
	INTERNALERROR_TEXTCONVERTFAILED = "InternalError.TextConvertFailed"

	// 错误的聊天输入文本。
	INVALIDPARAMETER_ERRORCHATTEXT = "InvalidParameter.ErrorChatText"

	// 错误的User参数。
	INVALIDPARAMETER_ERRORCHATUSER = "InvalidParameter.ErrorChatUser"

	// 请求数据长度无效。
	INVALIDPARAMETER_ERRORCONTENTLENGTH = "InvalidParameter.ErrorContentlength"

	// 没有body数据。
	INVALIDPARAMETER_ERRORNOBODYDATA = "InvalidParameter.ErrorNoBodydata"

	// 参数不全。
	INVALIDPARAMETER_ERRORPARAMSMISSING = "InvalidParameter.ErrorParamsMissing"

	// 解析请求数据失败。
	INVALIDPARAMETER_ERRORPARSEQUEST = "InvalidParameter.ErrorParsequest"

	// 参数取值错误
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// appid未注册。
	INVALIDPARAMETERVALUE_APPIDNOTREGISTERED = "InvalidParameterValue.AppIdNotRegistered"

	// AppId无效。
	INVALIDPARAMETERVALUE_ERRORINVALIDAPPID = "InvalidParameterValue.ErrorInvalidAppid"

	// ClientIp无效。
	INVALIDPARAMETERVALUE_ERRORINVALIDCLIENTIP = "InvalidParameterValue.ErrorInvalidClientip"

	// EngSerViceType无效。
	INVALIDPARAMETERVALUE_ERRORINVALIDENGSERVICE = "InvalidParameterValue.ErrorInvalidEngservice"

	// 是否进行翻译。
	INVALIDPARAMETERVALUE_ERRORINVALIDOPENTRANSLATE = "InvalidParameterValue.ErrorInvalidOpenTranslate"

	// ProjectId无效。
	INVALIDPARAMETERVALUE_ERRORINVALIDPROJECTID = "InvalidParameterValue.ErrorInvalidProjectid"

	// RequestId无效。
	INVALIDPARAMETERVALUE_ERRORINVALIDREQUESTID = "InvalidParameterValue.ErrorInvalidRequestid"

	// 翻译的源语言类型不支持。
	INVALIDPARAMETERVALUE_ERRORINVALIDSOURCELANGUAGE = "InvalidParameterValue.ErrorInvalidSourceLanguage"

	// SourceType无效。
	INVALIDPARAMETERVALUE_ERRORINVALIDSOURCETYPE = "InvalidParameterValue.ErrorInvalidSourcetype"

	// SubserviceType无效。
	INVALIDPARAMETERVALUE_ERRORINVALIDSUBSERVICETYPE = "InvalidParameterValue.ErrorInvalidSubservicetype"

	// 翻译的目标语言类型不支持。
	INVALIDPARAMETERVALUE_ERRORINVALIDTARGETLANGUAGE = "InvalidParameterValue.ErrorInvalidTargetLanguage"

	// Url无效。
	INVALIDPARAMETERVALUE_ERRORINVALIDURL = "InvalidParameterValue.ErrorInvalidUrl"

	// UsrAudioKey无效。
	INVALIDPARAMETERVALUE_ERRORINVALIDUSERAUDIOKEY = "InvalidParameterValue.ErrorInvalidUseraudiokey"

	// 音频编码格式不支持。
	INVALIDPARAMETERVALUE_ERRORINVALIDVOICEFORMAT = "InvalidParameterValue.ErrorInvalidVoiceFormat"

	// 音频数据无效。
	INVALIDPARAMETERVALUE_ERRORINVALIDVOICEDATA = "InvalidParameterValue.ErrorInvalidVoicedata"

	// 文本过长，请参考请求参数Text的说明。
	UNSUPPORTEDOPERATION_TEXTTOOLONG = "UnsupportedOperation.TextTooLong"
)
