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

package v20200820

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 匹配日志开通失败，请在线咨询。
	FAILEDOPERATION_CLSFAILEDOPERATION = "FailedOperation.CLSFailedOperation"

	// 频率限制。玩家重复请求匹配的间隔不小于100ms。
	FAILEDOPERATION_FREQUENCYSAMEPLAYERLIMITED = "FailedOperation.FrequencySamePlayerLimited"

	// 创建的匹配数超过限制。
	FAILEDOPERATION_LIMITMATCHCOUNT = "FailedOperation.LimitMatchCount"

	// 创建的规则数超过限制。
	FAILEDOPERATION_LIMITRULECOUNT = "FailedOperation.LimitRuleCount"

	// 系统错误。若无法解决，请在线咨询。
	FAILEDOPERATION_SERVICEUNAVAILABLE = "FailedOperation.ServiceUnavailable"

	// 标签操作失败。
	FAILEDOPERATION_TAGSOPEARTIONFAILED = "FailedOperation.TagsOpeartionFailed"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 接口不存在。
	INVALIDACTION = "InvalidAction"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// GameServerSession正在匹配中，无法再发起匹配。
	INVALIDPARAMETERVALUE_GAMESERVERSESSIONREPEATED = "InvalidParameterValue.GameServerSessionRepeated"

	// 规则脚本有误。
	INVALIDPARAMETERVALUE_INVALIDRULESCRIPT = "InvalidParameterValue.InvalidRuleScript"

	// 匹配 Code 不存在。
	INVALIDPARAMETERVALUE_MATCHCODENOTFOUND = "InvalidParameterValue.MatchCodeNotFound"

	// 字段值长度超过限制。
	INVALIDPARAMETERVALUE_MATCHFEILDVALUELIMIT = "InvalidParameterValue.MatchFeildValueLimit"

	// 字段中包含非法的字符。
	INVALIDPARAMETERVALUE_MATCHINVALIDCHARACTERS = "InvalidParameterValue.MatchInvalidCharacters"

	// 匹配不存在。
	INVALIDPARAMETERVALUE_MATCHNOTFOUND = "InvalidParameterValue.MatchNotFound"

	// Players 长度超过限制。
	INVALIDPARAMETERVALUE_MATCHPLAYERSLIMIT = "InvalidParameterValue.MatchPlayersLimit"

	// 重复的玩家Id。
	INVALIDPARAMETERVALUE_MATCHPLAYERSREPEATED = "InvalidParameterValue.MatchPlayersRepeated"

	// 当前状态不能取消。
	INVALIDPARAMETERVALUE_MATCHSTATUSNOTPERMITCANCEL = "InvalidParameterValue.MatchStatusNotPermitCancel"

	// 匹配 TicketId 不存在。
	INVALIDPARAMETERVALUE_MATCHTICKETIDNOTFOUND = "InvalidParameterValue.MatchTicketIdNotFound"

	// 匹配 TicketId 重复。
	INVALIDPARAMETERVALUE_MATCHTICKETIDREPEATED = "InvalidParameterValue.MatchTicketIdRepeated"

	// 匹配票据列表长度超过限制。
	INVALIDPARAMETERVALUE_MATCHTICKETLIMIT = "InvalidParameterValue.MatchTicketLimit"

	// 规则存在关联的匹配。
	INVALIDPARAMETERVALUE_RULEMATCHEXISTENT = "InvalidParameterValue.RuleMatchExistent"

	// 规则名称已存在。
	INVALIDPARAMETERVALUE_RULENAMEDUPLICATED = "InvalidParameterValue.RuleNameDuplicated"

	// 规则不存在。
	INVALIDPARAMETERVALUE_RULENOTFOUND = "InvalidParameterValue.RuleNotFound"

	// 没有该标签权限。
	INVALIDPARAMETERVALUE_TAGSLIMITPERMISSION = "InvalidParameterValue.TagsLimitPermission"

	// Token CompatibleSpan 取值超出限制。
	INVALIDPARAMETERVALUE_TOKENCOMPATIBLESPANINVALID = "InvalidParameterValue.TokenCompatibleSpanInvalid"

	// Token长度超出限制。
	INVALIDPARAMETERVALUE_TOKENLIMIT = "InvalidParameterValue.TokenLimit"

	// 取值超出范围。
	INVALIDPARAMETERVALUE_VALUERANGELIMIT = "InvalidParameterValue.ValueRangeLimit"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// Token更新的频率超出限制。
	LIMITEXCEEDED_TOKENUPDATEEXCEED = "LimitExceeded.TokenUpdateExceed"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 接口版本不存在。
	NOSUCHVERSION = "NoSuchVersion"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// CLS日志服务未开通。
	RESOURCEUNAVAILABLE_CLSSERVICENOTACTIVATED = "ResourceUnavailable.CLSServiceNotActivated"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 需要授权GPM角色权限。
	UNAUTHORIZEDOPERATION_CAMUNAUTHORIZEDOPERATION = "UnauthorizedOperation.CAMUnauthorizedOperation"

	// 需要授权接口权限。
	UNAUTHORIZEDOPERATION_UNAUTHORIZEDACTION = "UnauthorizedOperation.UnauthorizedAction"

	// 当前账号未完成实名认证，使用服务前请先进行实名认证。
	UNAUTHORIZEDOPERATION_USERUNAUTH = "UnauthorizedOperation.UserUnAuth"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 接口不支持所传地域。
	UNSUPPORTEDREGION = "UnsupportedRegion"
)
