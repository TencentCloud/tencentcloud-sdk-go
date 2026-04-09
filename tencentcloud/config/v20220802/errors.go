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

package v20220802

const (
	// 此产品的特有错误码

	// 合规包不存在。
	FAILEDOPERATION_COMPLIANCEPACKISNOTEXIST = "FailedOperation.CompliancePackIsNotExist"

	// 合规包未停止。
	FAILEDOPERATION_COMPLIANCEPACKISNOTSTOP = "FailedOperation.CompliancePackIsNotStop"

	// 合规包名称已存在。
	FAILEDOPERATION_COMPLIANCEPACKNAMEISNOTEXIST = "FailedOperation.CompliancePackNameIsNotExist"

	// 投递的成员账号不是委派管理员或者管理员
	FAILEDOPERATION_MEMBERNOTASSIGNMANAGER = "FailedOperation.MemberNotAssignManager"

	// 不允许操作控制中心合规包
	FAILEDOPERATION_NOTALLOWOPERATECONTROLCENTERGOVERNANCEPACKAGE = "FailedOperation.NotAllowOperateControlCenterGovernancePackage"

	// 不允许操作控制中心合规包规则
	FAILEDOPERATION_NOTALLOWOPERATECONTROLCENTERGOVERNANCEPACKAGERULE = "FailedOperation.NotAllowOperateControlCenterGovernancePackageRule"

	// 监控已关闭。
	FAILEDOPERATION_RECORDERISCLOSE = "FailedOperation.RecorderIsClose"

	// 规则未停用。
	FAILEDOPERATION_RULEISNOTSTOP = "FailedOperation.RuleIsNotStop"

	// 规则名称已存在。
	FAILEDOPERATION_RULENAMEISEXIST = "FailedOperation.RuleNameIsExist"

	// 相同成员的账号组已存在
	FAILEDOPERATION_SAMEMEMBERACCOUNTGROUPALREADYEXIST = "FailedOperation.SameMemberAccountGroupAlreadyExist"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 用户组数量超过上限。
	LIMITEXCEEDED_ACCOUNTGROUPLIMITOVER = "LimitExceeded.AccountGroupLimitOver"

	// 合规包数量到达上限。
	LIMITEXCEEDED_COMPLIANCEPACKMAXNUM = "LimitExceeded.CompliancePackMaxNum"

	// 当日打开监控次数达到上限。
	LIMITEXCEEDED_OPENRECORDERLIMITOVER = "LimitExceeded.OpenRecorderLimitOver"

	// 创建修复设置超过上限。
	LIMITEXCEEDED_REMEDIATIONMAXNUM = "LimitExceeded.RemediationMaxNum"

	// 当日修改监控资源类型次数达到上限。
	LIMITEXCEEDED_UPDATERECORDERLIMITOVER = "LimitExceeded.UpdateRecorderLimitOver"

	// 账号组不存在。
	RESOURCENOTFOUND_ACCOUNTGROUPISNOTEXIST = "ResourceNotFound.AccountGroupIsNotExist"

	// 合规包不存在。
	RESOURCENOTFOUND_COMPLIANCEPACKISNOTEXIST = "ResourceNotFound.CompliancePackIsNotExist"

	// 合规包不存在。
	RESOURCENOTFOUND_COMPLIANCEPACKNOTEXIST = "ResourceNotFound.CompliancePackNotExist"

	// 投递的成员账号的COS 桶不存在
	RESOURCENOTFOUND_DELIVERBUCKETNOTEXIST = "ResourceNotFound.DeliverBucketNotExist"

	// 投递类型不存在。
	RESOURCENOTFOUND_DELIVEREVENTNOTEXIST = "ResourceNotFound.DeliverEventNotExist"

	// 投递信息不存在。
	RESOURCENOTFOUND_DELIVERISNOTEXIST = "ResourceNotFound.DeliverIsNotExist"

	// 成员不存在。
	RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"

	// 监控信息不存在。
	RESOURCENOTFOUND_RECORDERISNOTEXIST = "ResourceNotFound.RecorderIsNotExist"

	// 修正设置不存在。
	RESOURCENOTFOUND_REMEDIATIONISNOTEXIST = "ResourceNotFound.RemediationIsNotExist"

	// 资源不存在。
	RESOURCENOTFOUND_RESOURCENOTEXIST = "ResourceNotFound.ResourceNotExist"

	// 规则不存在。
	RESOURCENOTFOUND_RULEISNOTEXIST = "ResourceNotFound.RuleIsNotExist"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
)
