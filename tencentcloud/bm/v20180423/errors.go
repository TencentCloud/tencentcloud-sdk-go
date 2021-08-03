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

package v20180423

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 存在未结束故障单，操作失败。
	FAILEDOPERATION_EXISTREPAIRTASK = "FailedOperation.ExistRepairTask"

	// Tsc Agent不在线。
	FAILEDOPERATION_TSCAGENTOFFLINE = "FailedOperation.TscAgentOffline"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数值错误。
	INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 脚本数量超过上限。
	LIMITEXCEEDED_USERCMDCOUNT = "LimitExceeded.UserCmdCount"

	// 流程操作繁忙，请稍后重试。
	RESOURCEINUSE_FLOWBUSY = "ResourceInUse.FlowBusy"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 设备不足。
	RESOURCEINSUFFICIENT_DEVICEINSUFFICIENT = "ResourceInsufficient.DeviceInsufficient"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 余额不足。
	UNSUPPORTEDOPERATION_FUNDINSUFFICIENT = "UnsupportedOperation.FundInsufficient"

	// 设备不支持此操作。
	UNSUPPORTEDOPERATION_INVALIDOPERATION = "UnsupportedOperation.InvalidOperation"
)
