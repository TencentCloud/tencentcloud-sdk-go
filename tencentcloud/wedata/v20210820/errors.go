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

package v20210820

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 未登录或登录已过期。
	AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"

	// CAM未授权，请联系主账号到CAM中授权QcloudWeDataFullAccess策略给该账户。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 连接超时
	FAILEDOPERATION_CONNECTIONTIMEOUTERROR = "FailedOperation.ConnectionTimeOutError"

	// 当前操作的资源状态不合法。
	FAILEDOPERATION_EXECUTORCLUSTERSTATUSERROR = "FailedOperation.ExecutorClusterStatusError"

	// 操作失败
	FAILEDOPERATION_FAILEDOPERATIONWITHREASON = "FailedOperation.FailedOperationWithReason"

	// 查询数据开发资源锁状态-当前用户不持有锁
	FAILEDOPERATION_NOLOCK = "FailedOperation.NoLock"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 外部系统调用异常。
	INTERNALERROR_CALLSCHEDULERAPIERROR = "InternalError.CallSchedulerApiError"

	// 调用云API失败。
	INTERNALERROR_INTERNALCALLCLOUDAPIERROR = "InternalError.InternalCallCloudApiError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 告警接收地址不存在。
	INVALIDPARAMETER_ALARMRECEIVEADDRESSNOTEXIST = "InvalidParameter.AlarmReceiveAddressNotExist"

	// ClientIp未被授权。
	INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"

	// 数据引擎实例不存在。
	INVALIDPARAMETER_DATAENGINEINSTANCENOTEXISTS = "InvalidParameter.DataEngineInstanceNotExists"

	// 名称重复。
	INVALIDPARAMETER_DUPLICATENAME = "InvalidParameter.DuplicateName"

	// 告警接收地址不合法。
	INVALIDPARAMETER_INVALIDALARMURL = "InvalidParameter.InvalidAlarmUrl"

	// 查询过滤条件参数错误。
	INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"

	// 超出最大限额
	INVALIDPARAMETER_MAXLIMITEXCEEDED = "InvalidParameter.MaxLimitExceeded"

	// 已超过系统安全配额。
	INVALIDPARAMETER_QUOTAEXCEEDERROR = "InvalidParameter.QuotaExceedError"

	// 规则不存在。
	INVALIDPARAMETER_RULENOTEXIST = "InvalidParameter.RuleNotExist"

	// 规则模版不存在。
	INVALIDPARAMETER_RULETEMPLATENOTEXIST = "InvalidParameter.RuleTemplateNotExist"

	// 服务繁忙，请稍后重试。
	INVALIDPARAMETER_SERVICEISBUSY = "InvalidParameter.ServiceIsBusy"

	// WeData_QCSRole不存在，请进行服务授权。
	INVALIDPARAMETER_WEDATAROLENOTEXISTS = "InvalidParameter.WeDataRoleNotExists"

	// 项目名重复。
	INVALIDPARAMETER_WORKSPACENAMEDUPLICATION = "InvalidParameter.WorkspaceNameDuplication"

	// 工作空间不存在。
	INVALIDPARAMETER_WORKSPACENOTEXIST = "InvalidParameter.WorkspaceNotExist"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 数据建模实例已经被销毁
	INVALIDPARAMETERVALUE_DATAMODELDESTROYED = "InvalidParameterValue.DataModelDestroyed"

	// 重复购买错误，一个主账号在一个地域只允许购买一套数据建模
	INVALIDPARAMETERVALUE_REPEATPURCHASEERROR = "InvalidParameterValue.RepeatPurchaseError"

	// 不支持的数据建模服务提供方
	INVALIDPARAMETERVALUE_UNSUPPORTEDPROVIDER = "InvalidParameterValue.UnsupportedProvider"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 用户不在白名单。
	OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"

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

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 用户不是当前项目成员。
	UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 任务不在同一时区
	UNSUPPORTEDOPERATION_FAILEDCHECKTASKDEPENDENCETIMEZONE = "UnsupportedOperation.FailedCheckTaskDependenceTimeZone"

	// 请先配置不少于一个规则。
	UNSUPPORTEDOPERATION_NORULEINRULEGROUP = "UnsupportedOperation.NoRuleInRuleGroup"
)
