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

package v20201101

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 授权数不足。
	FAILEDOPERATION_AUTHORIZEDNOTENOUGH = "FailedOperation.AuthorizedNotEnough"

	// 响应数据值不正确。
	FAILEDOPERATION_DATAVALUENOTCORRECT = "FailedOperation.DataValueNotCorrect"

	// 在扫描中或无扫描权限建议授权后扫描。
	FAILEDOPERATION_ERRALREADYSCANNING = "FailedOperation.ErrAlreadyScanning"

	// 通知策略变更失败。
	FAILEDOPERATION_NOTIFYPOLICYCHANGEFAILED = "FailedOperation.NotifyPolicyChangeFailed"

	// 子规则配置过多。
	FAILEDOPERATION_RULECONFIGTOOMANY = "FailedOperation.RuleConfigTooMany"

	// 规则信息存在重复。
	FAILEDOPERATION_RULEINFOREPEAT = "FailedOperation.RuleInfoRepeat"

	// 规则名字存在重复。
	FAILEDOPERATION_RULENAMEREPEAT = "FailedOperation.RuleNameRepeat"

	// 当前规则信息未找到。
	FAILEDOPERATION_RULENOTFIND = "FailedOperation.RuleNotFind"

	// 选择镜像数量过多。
	FAILEDOPERATION_RULESELECTIMAGEOUTRANGE = "FailedOperation.RuleSelectImageOutRange"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 用户未授权。
	INTERNALERROR_ERRROLENOTEXIST = "InternalError.ErrRoleNotExist"

	// 操作数据库失败。
	INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// IP格式不合法。
	INVALIDPARAMETER_ERRIPNOVALID = "InvalidParameter.ErrIpNoValid"

	// 参数格式错误。
	INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"

	// 缺少必须参数。
	INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"

	// 参数解析错误。
	INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"

	// 端口格式不合法。
	INVALIDPARAMETER_PORTNOVALID = "InvalidParameter.PortNoValid"

	// 进程名/目标IP/目标端口，不能同时为空。
	INVALIDPARAMETER_REVERSHELLKEYFIELDALLEMPTY = "InvalidParameter.ReverShellKeyFieldAllEmpty"

	// 前规则信息参数非法。
	INVALIDPARAMETER_RULEINFOINVALID = "InvalidParameter.RuleInfoInValid"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 当前数据未能查询到。
	INVALIDPARAMETERVALUE_DATANOTFOUND = "InvalidParameterValue.DataNotFound"

	// 当前数据区间错误。
	INVALIDPARAMETERVALUE_DATARANGE = "InvalidParameterValue.DataRange"

	// 参数长度受限。
	INVALIDPARAMETERVALUE_LENGTHLIMIT = "InvalidParameterValue.LengthLimit"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
)
