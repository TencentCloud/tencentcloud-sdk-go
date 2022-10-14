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

package v20170312

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 云盘退还数量已达上限，不能再退还。
	INSUFFICIENTREFUNDQUOTA = "InsufficientRefundQuota"

	// 快照配额不足。
	INSUFFICIENTSNAPSHOTQUOTA = "InsufficientSnapshotQuota"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 依赖组件请求失败，请联系客服人员解决。
	INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"

	// 查询资源失败。
	INTERNALERROR_FAILQUERYRESOURCE = "InternalError.FailQueryResource"

	// 对资源的操作失败，具体错误信息请查看错误描述 Message 字段，稍后重试或者联系客服人员帮忙解决。
	INTERNALERROR_RESOURCEOPFAILED = "InternalError.ResourceOpFailed"

	// 账户余额不足。
	INVALIDACCOUNT_INSUFFICIENTBALANCE = "InvalidAccount.InsufficientBalance"

	// 输入的`AutoSnapshotPolicyId`不存在。
	INVALIDAUTOSNAPSHOTPOLICYID_NOTFOUND = "InvalidAutoSnapshotPolicyId.NotFound"

	// 云盘已绑定定期快照策略。
	INVALIDDISK_ALREADYBOUND = "InvalidDisk.AlreadyBound"

	// 云硬盘忙，请稍后重试。
	INVALIDDISK_BUSY = "InvalidDisk.Busy"

	// 云盘已过期。
	INVALIDDISK_EXPIRE = "InvalidDisk.Expire"

	// 不支持非弹性云盘。
	INVALIDDISK_NOTPORTABLE = "InvalidDisk.NotPortable"

	// 云硬盘没有快照能力。
	INVALIDDISK_NOTSUPPORTSNAPSHOT = "InvalidDisk.NotSupportSnapshot"

	// 云硬盘不支持该操作。
	INVALIDDISK_NOTSUPPORTED = "InvalidDisk.NotSupported"

	// 云硬盘正在创建快照，请稍后重试。
	INVALIDDISK_SNAPSHOTCREATING = "InvalidDisk.SnapshotCreating"

	// 云硬盘类型错误。
	INVALIDDISK_TYPEERROR = "InvalidDisk.TypeError"

	// 输入的`DiskId`不存在。
	INVALIDDISKID_NOTFOUND = "InvalidDiskId.NotFound"

	// 指定的Filter不被支持。
	INVALIDFILTER = "InvalidFilter"

	// 云服务器不支持挂载云盘。
	INVALIDINSTANCE_NOTSUPPORTED = "InvalidInstance.NotSupported"

	// 输入实例的`InstanceId`不存在。
	INVALIDINSTANCEID_NOTFOUND = "InvalidInstanceId.NotFound"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 当前地域不支持当前配置的云盘。
	INVALIDPARAMETER_DISKCONFIGNOTSUPPORTED = "InvalidParameter.DiskConfigNotSupported"

	// 云硬盘大小与快照大小不匹配。
	INVALIDPARAMETER_DISKSIZENOTMATCH = "InvalidParameter.DiskSizeNotMatch"

	// 项目ID不存在。
	INVALIDPARAMETER_PROJECTIDNOTEXIST = "InvalidParameter.ProjectIdNotExist"

	// 需要将快照转化成镜像后再执行操作。
	INVALIDPARAMETER_SHOULDCONVERTSNAPSHOTTOIMAGE = "InvalidParameter.ShouldConvertSnapshotToImage"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 标签绑定云盘超过限制。
	INVALIDPARAMETERVALUE_BINDDISKLIMITEXCEEDED = "InvalidParameterValue.BindDiskLimitExceeded"

	// 参数值数量超过限制。
	INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"

	// 快照不支持该操作。
	INVALIDSNAPSHOT_NOTSUPPORTED = "InvalidSnapshot.NotSupported"

	// 输入的`SnapshotId`不存在。
	INVALIDSNAPSHOTID_NOTFOUND = "InvalidSnapshotId.NotFound"

	// 定期快照策略数量已达到上限。
	LIMITEXCEEDED_AUTOSNAPSHOTPOLICYOUTOFQUOTA = "LimitExceeded.AutoSnapshotPolicyOutOfQuota"

	// 实例挂载云盘数量超过限制。
	LIMITEXCEEDED_INSTANCEATTACHEDDISK = "LimitExceeded.InstanceAttachedDisk"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 资源繁忙，请稍后重试。
	RESOURCEBUSY = "ResourceBusy"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 指定快照正在复制到目标地域。
	RESOURCEINUSE_COPYSNAPSHOTCONFLICT = "ResourceInUse.CopySnapshotConflict"

	// 云硬盘当前已在迁移中，请稍后重试。
	RESOURCEINUSE_DISKMIGRATING = "ResourceInUse.DiskMigrating"

	// 云硬盘正在执行快照回滚操作，请稍后重试。
	RESOURCEINUSE_DISKROLLBACKING = "ResourceInUse.DiskRollbacking"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 配额不足。
	RESOURCEINSUFFICIENT_OVERQUOTA = "ResourceInsufficient.OverQuota"

	// 云盘退还数量已达上限，不能再退还。
	RESOURCEINSUFFICIENT_OVERREFUNDQUOTA = "ResourceInsufficient.OverRefundQuota"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不存在。
	RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 云硬盘已挂载至其他云服务器上。
	RESOURCEUNAVAILABLE_ATTACHED = "ResourceUnavailable.Attached"

	// 云硬盘快照链长度过长，拒绝创建快照。
	RESOURCEUNAVAILABLE_DISKSNAPSHOTCHAINTOOLARGE = "ResourceUnavailable.DiskSnapshotChainTooLarge"

	// 云硬盘已过期。
	RESOURCEUNAVAILABLE_EXPIRE = "ResourceUnavailable.Expire"

	// 非弹性云硬盘不支持此操作。
	RESOURCEUNAVAILABLE_NOTPORTABLE = "ResourceUnavailable.NotPortable"

	// 云盘不支持退还。
	RESOURCEUNAVAILABLE_NOTSUPPORTREFUND = "ResourceUnavailable.NotSupportRefund"

	// 资源不支持此操作。
	RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"

	// 云硬盘已退还。
	RESOURCEUNAVAILABLE_REPEATREFUND = "ResourceUnavailable.RepeatRefund"

	// 快照尚未创建完成，暂时不可使用。
	RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"

	// 全网正在创建的快照数量过多。
	RESOURCEUNAVAILABLE_TOOMANYCREATINGSNAPSHOT = "ResourceUnavailable.TooManyCreatingSnapshot"

	// 云硬盘类型错误，如尝试挂载系统盘至云服务器上。
	RESOURCEUNAVAILABLE_TYPEERROR = "ResourceUnavailable.TypeError"

	// 云硬盘与实例不在同一可用区。
	RESOURCEUNAVAILABLE_ZONENOTMATCH = "ResourceUnavailable.ZoneNotMatch"

	// 订单冲突。
	TRADEDEALCONFLICT = "TradeDealConflict"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// MFA鉴权过期，请重试。
	UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"

	// 账号没有实名认证，支付失败。
	UNAUTHORIZEDOPERATION_NOTCERTIFICATION = "UnauthorizedOperation.NotCertification"

	// 没有支付权限。
	UNAUTHORIZEDOPERATION_NOTHAVEPAYMENTRIGHT = "UnauthorizedOperation.NotHavePaymentRight"

	// 授权角色不存在。
	UNAUTHORIZEDOPERATION_ROLENOTEXISTS = "UnauthorizedOperation.RoleNotExists"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 暂不支持从EKS上卸载云硬盘。
	UNSUPPORTEDOPERATION_DETACHPOD = "UnsupportedOperation.DetachPod"

	// 磁盘已加密。
	UNSUPPORTEDOPERATION_DISKENCRYPT = "UnsupportedOperation.DiskEncrypt"

	// 云盘挂载的实例未关机。
	UNSUPPORTEDOPERATION_INSTANCENOTSTOPPED = "UnsupportedOperation.InstanceNotStopped"

	// 该快照已经共享，请先解除共享。
	UNSUPPORTEDOPERATION_SNAPHASSHARED = "UnsupportedOperation.SnapHasShared"

	// 该快照创建了自定义快照，请先删除对应镜像。
	UNSUPPORTEDOPERATION_SNAPSHOTHASBINDEDIMAGE = "UnsupportedOperation.SnapshotHasBindedImage"

	// 快照不支持跨地域复制。
	UNSUPPORTEDOPERATION_SNAPSHOTNOTSUPPORTCOPY = "UnsupportedOperation.SnapshotNotSupportCopy"

	// 资源当前状态不支持该操作。
	UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"
)
