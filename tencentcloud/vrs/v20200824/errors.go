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

package v20200824

const (
	// 此产品的特有错误码

	// CAM权限错误。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 任务状态错误。
	FAILEDOPERATION_ERRORINVALIDTASKSTATUS = "FailedOperation.ErrorInvalidTaskStatus"

	// 任务不存在。
	FAILEDOPERATION_NOSUCHTASK = "FailedOperation.NoSuchTask"

	// 检测失败。
	FAILEDOPERATION_VOICEEVALUATEFAILED = "FailedOperation.VoiceEvaluateFailed"

	// 音频质量差。
	FAILEDOPERATION_VOICENOTQUALIFIED = "FailedOperation.VoiceNotQualified"

	// 内部错误。
	INTERNALERROR_FAILACCESSDATABASE = "InternalError.FailAccessDatabase"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 音频数据错误。
	INVALIDPARAMETERVALUE_AUDIODATA = "InvalidParameterValue.AudioData"

	// 音频时长超限，时长不可超过15s
	INVALIDPARAMETERVALUE_AUDIODURATIONEXCEEDSLIMIT = "InvalidParameterValue.AudioDurationExceedsLimit"

	// Codec非法，请参考Codec参数说明。
	INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"

	// 无效TaskId。
	INVALIDPARAMETERVALUE_ERRORINVALIDTASKID = "InvalidParameterValue.ErrorInvalidTaskId"

	// SampleRate非法，请参考SampleRate参数说明。
	INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"

	// 音色性别错误。
	INVALIDPARAMETERVALUE_VOICEGENDER = "InvalidParameterValue.VoiceGender"

	// 音色语言错误。
	INVALIDPARAMETERVALUE_VOICELANGUAGE = "InvalidParameterValue.VoiceLanguage"

	// 音色名称错误。
	INVALIDPARAMETERVALUE_VOICENAME = "InvalidParameterValue.VoiceName"

	// 无声音复刻任务配额
	UNSUPPORTEDOPERATION_VRSQUOTAEXHAUSTED = "UnsupportedOperation.VRSQuotaExhausted"
)
