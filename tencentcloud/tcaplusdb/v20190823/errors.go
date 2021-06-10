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

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 接口操作鉴权错误。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 旧密码已经失效。
	FAILEDOPERATION_OLDPASSWORDHASEXPIRED = "FailedOperation.OldPasswordHasExpired"

	// 旧密码还未失效。
	FAILEDOPERATION_OLDPASSWORDINUSE = "FailedOperation.OldPasswordInUse"

	// 密码错误。
	FAILEDOPERATION_PASSWORDFAILURE = "FailedOperation.PasswordFailure"

	// 请求的地域不匹配。
	FAILEDOPERATION_REGIONMISMATCH = "FailedOperation.RegionMismatch"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 应用名不合法。
	INVALIDPARAMETERVALUE_INVALIDAPPNAME = "InvalidParameterValue.InvalidAppName"

	// 集群名称不合法。
	INVALIDPARAMETERVALUE_INVALIDCLUSTERNAME = "InvalidParameterValue.InvalidClusterName"

	// 非法的表格组名称。
	INVALIDPARAMETERVALUE_INVALIDTABLEGROUPNAME = "InvalidParameterValue.InvalidTableGroupName"

	// 非法的时间格式。
	INVALIDPARAMETERVALUE_INVALIDTIMEVALUE = "InvalidParameterValue.InvalidTimeValue"

	// 不支持的应用数据描述类型。
	INVALIDPARAMETERVALUE_UNSUPPORTIDLTYPE = "InvalidParameterValue.UnsupportIdlType"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 余额不足。
	RESOURCEINSUFFICIENT_BALANCEERROR = "ResourceInsufficient.BalanceError"

	// 没有可用的应用资源。
	RESOURCEINSUFFICIENT_NOAVAILABLEAPP = "ResourceInsufficient.NoAvailableApp"

	// 没有可用的集群资源。
	RESOURCEINSUFFICIENT_NOAVAILABLECLUSTER = "ResourceInsufficient.NoAvailableCluster"

	// 私有网络中没有可用的vip资源。
	RESOURCEINSUFFICIENT_NOENOUGHVIPINVPC = "ResourceInsufficient.NoEnoughVipInVPC"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 集群名称重复。
	RESOURCEUNAVAILABLE_DUPLICATECLUSTERNAME = "ResourceUnavailable.DuplicateClusterName"

	// 重复的表格组id或名称。
	RESOURCEUNAVAILABLE_DUPLICATETABLEGROUPINFO = "ResourceUnavailable.DuplicateTableGroupInfo"

	// 重复的表格组名称。
	RESOURCEUNAVAILABLE_DUPLICATETABLEGROUPNAME = "ResourceUnavailable.DuplicateTableGroupName"

	// 没有可用的表格组资源。
	RESOURCEUNAVAILABLE_NOAVAILABLETABLEGROUP = "ResourceUnavailable.NoAvailableTableGroup"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
