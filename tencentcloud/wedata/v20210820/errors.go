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

package v20210820

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 未登陆或登陆已过期。
	AUTHFAILURE_SIGNATUREEXPIRE = "AuthFailure.SignatureExpire"

	// CAM未授权，请联系主账号到CAM中授权QcloudWeDataFullAccess策略给该账户。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// DryRun 操作，代表请求将会是成功的，只是多传了 DryRun 参数。
	DRYRUNOPERATION = "DryRunOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 外部系统调用异常。
	INTERNALERROR_CALLSCHEDULERAPIERROR = "InternalError.CallSchedulerApiError"

	// 调用云API失败。
	INTERNALERROR_INTERNALCALLCLOUDAPIERROR = "InternalError.InternalCallCloudApiError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// ClientIp未被授权。
	INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"

	// 名称重复。
	INVALIDPARAMETER_DUPLICATENAME = "InvalidParameter.DuplicateName"

	// 查询过滤条件参数错误。
	INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"

	// 规则不存在。
	INVALIDPARAMETER_RULENOTEXIST = "InvalidParameter.RuleNotExist"

	// 规则模版不存在。
	INVALIDPARAMETER_RULETEMPLATENOTEXIST = "InvalidParameter.RuleTemplateNotExist"

	// 服务繁忙，请稍后重试。
	INVALIDPARAMETER_SERVICEISBUSY = "InvalidParameter.ServiceIsBusy"

	// 工作空间不存在。
	INVALIDPARAMETER_WORKSPACENOTEXIST = "InvalidParameter.WorkspaceNotExist"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

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

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 用户不是当前项目成员。
	UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 请先配置不少于一个规则。
	UNSUPPORTEDOPERATION_NORULEINRULEGROUP = "UnsupportedOperation.NoRuleInRuleGroup"
)
