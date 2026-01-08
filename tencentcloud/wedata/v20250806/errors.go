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

package v20250806

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 当前操作的资源状态不合法。
	FAILEDOPERATION_EXECUTORCLUSTERSTATUSERROR = "FailedOperation.ExecutorClusterStatusError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// ClientIp未被授权。
	INVALIDPARAMETER_CLIENTIPNOTAUTHORIZED = "InvalidParameter.ClientIpNotAuthorized"

	// 名称重复。
	INVALIDPARAMETER_DUPLICATENAME = "InvalidParameter.DuplicateName"

	// 查询过滤条件参数错误。
	INVALIDPARAMETER_INVALIDFILTERPARAMETER = "InvalidParameter.InvalidFilterParameter"

	// Missing Servlet Request Parameter
	INVALIDPARAMETER_MISSINGREQUESTPARAMETER = "InvalidParameter.MissingRequestParameter"

	// 已超过系统安全配额。
	INVALIDPARAMETER_QUOTAEXCEEDERROR = "InvalidParameter.QuotaExceedError"

	// 规则不存在。
	INVALIDPARAMETER_RULENOTEXIST = "InvalidParameter.RuleNotExist"

	// 服务繁忙，请稍后重试。
	INVALIDPARAMETER_SERVICEISBUSY = "InvalidParameter.ServiceIsBusy"

	// 项目名重复。
	INVALIDPARAMETER_WORKSPACENAMEDUPLICATION = "InvalidParameter.WorkspaceNameDuplication"

	// 工作空间不存在。
	INVALIDPARAMETER_WORKSPACENOTEXIST = "InvalidParameter.WorkspaceNotExist"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 告警规则名称重复
	INVALIDPARAMETERVALUE_RULENAMEREPEATED = "InvalidParameterValue.RuleNameRepeated"

	// Param Validation Error
	INVALIDPARAMETERVALUE_VALIDATIONERROR = "InvalidParameterValue.ValidationError"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 用户不在白名单。
	OPERATIONDENIED_USERNOTINWHITELISTERROR = "OperationDenied.UserNotInWhitelistError"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 用户不是当前项目成员。
	UNAUTHORIZEDOPERATION_USERNOTINPROJECT = "UnauthorizedOperation.UserNotInProject"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
