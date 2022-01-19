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

package v20180412

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 认证无效。
	AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"

	// 实例安全组信息添加失败。
	FAILEDOPERATION_ADDINSTANCEINFOFAILED = "FailedOperation.AddInstanceInfoFailed"

	// 绑定安全组失败。
	FAILEDOPERATION_ASSOCIATESECURITYGROUPSFAILED = "FailedOperation.AssociateSecurityGroupsFailed"

	// 实例安全组信息清除失败。
	FAILEDOPERATION_CLEARINSTANCEINFOFAILED = "FailedOperation.ClearInstanceInfoFailed"

	// 提交工作流失败。
	FAILEDOPERATION_COMMITFLOWERROR = "FailedOperation.CommitFlowError"

	// 解绑安全组失败。
	FAILEDOPERATION_DISASSOCIATESECURITYGROUPSFAILED = "FailedOperation.DisassociateSecurityGroupsFailed"

	// 当前DTS状态下不允许该操作。
	FAILEDOPERATION_DTSSTATUSABNORMAL = "FailedOperation.DtsStatusAbnormal"

	// 流表不存在。
	FAILEDOPERATION_FLOWNOTEXISTS = "FailedOperation.FlowNotExists"

	// 获取安全组详情失败。
	FAILEDOPERATION_GETSECURITYGROUPDETAILFAILED = "FailedOperation.GetSecurityGroupDetailFailed"

	// 支付失败。
	FAILEDOPERATION_PAYFAILED = "FailedOperation.PayFailed"

	// 暂时笼统的定义这个错误码。
	FAILEDOPERATION_REDOFLOWFAILED = "FailedOperation.RedoFlowFailed"

	// 设置规则失败。
	FAILEDOPERATION_SETRULELOCATIONFAILED = "FailedOperation.SetRuleLocationFailed"

	// 内部系统错误，和业务无关。
	FAILEDOPERATION_SYSTEMERROR = "FailedOperation.SystemError"

	// 实例不支持该接口。
	FAILEDOPERATION_UNSUPPORTERROR = "FailedOperation.UnSupportError"

	// weekday输入无效数据。
	FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"

	// 实例安全组信息更新失败。
	FAILEDOPERATION_UPDATEINSTANCEINFOFAILED = "FailedOperation.UpdateInstanceInfoFailed"

	// 更新安全组失败。
	FAILEDOPERATION_UPDATESECURITYGROUPSFAILED = "FailedOperation.UpdateSecurityGroupsFailed"

	// cam鉴权错误。
	INTERNALERROR_CAMAUTHOSSRESPONSERETURNCODEERROR = "InternalError.CamAuthOssResponseReturnCodeError"

	// 统一的 DB 操作错误，可以是 update insert select..。
	INTERNALERROR_DBOPERATIONFAILED = "InternalError.DbOperationFailed"

	// 执行Http请求失败。
	INTERNALERROR_EXECHTTPREQUESTERROR = "InternalError.ExecHttpRequestError"

	// 内部错误。
	INTERNALERROR_INTERNALERROR = "InternalError.InternalError"

	// 获取实例列表出错。
	INTERNALERROR_LISTINSTANCESERROR = "InternalError.ListInstancesError"

	// 网络错误。
	INTERNALERROR_NETWORKERR = "InternalError.NetWorkErr"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 参数为空。
	INVALIDPARAMETER_EMPTYPARAM = "InvalidParameter.EmptyParam"

	// 非法参数错误。
	INVALIDPARAMETER_ILLEGALPARAMETERERROR = "InvalidParameter.IllegalParameterError"

	// 业务参数错误。
	INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"

	// 参数错误，不支持操作。
	INVALIDPARAMETER_NOTSUPPORTED = "InvalidParameter.NotSupported"

	// 上海金融只提供vpc网络。
	INVALIDPARAMETER_ONLYVPCONSPECZONEID = "InvalidParameter.OnlyVPCOnSpecZoneId"

	// 时间格式或者范围不符合要求。
	INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"

	// 接口没有cam权限。
	INVALIDPARAMETER_PERMISSIONDENIED = "InvalidParameter.PermissionDenied"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 备份不存在。
	INVALIDPARAMETERVALUE_BACKUPNOTEXISTS = "InvalidParameterValue.BackupNotExists"

	// 不是vpc网络下实例。
	INVALIDPARAMETERVALUE_BASENETWORKACCESSDENY = "InvalidParameterValue.BaseNetWorkAccessDeny"

	// 业务校验不通过。
	INVALIDPARAMETERVALUE_CHECKNOTPASS = "InvalidParameterValue.CheckNotPass"

	// 重命名，命名规则错误。
	INVALIDPARAMETERVALUE_INSTANCENAMERULEERROR = "InvalidParameterValue.InstanceNameRuleError"

	// 请求购买的实例类型错误（TypeId 1:集群版；2:主从版,即原主从版)。
	INVALIDPARAMETERVALUE_INVALIDINSTANCETYPEID = "InvalidParameterValue.InvalidInstanceTypeId"

	// vpc网络下，vpcid 子网id 非法。
	INVALIDPARAMETERVALUE_INVALIDSUBNETID = "InvalidParameterValue.InvalidSubnetId"

	// 请求的容量不在售卖容量范围内。
	INVALIDPARAMETERVALUE_MEMSIZENOTINRANGE = "InvalidParameterValue.MemSizeNotInRange"

	// 实例不能重复绑定。
	INVALIDPARAMETERVALUE_NOTREPEATBIND = "InvalidParameterValue.NotRepeatBind"

	// 密码为空。
	INVALIDPARAMETERVALUE_PASSWORDEMPTY = "InvalidParameterValue.PasswordEmpty"

	// 密码校验出错，密码错误。
	INVALIDPARAMETERVALUE_PASSWORDERROR = "InvalidParameterValue.PasswordError"

	// 腾讯集团内部账号禁止使用免密实例。
	INVALIDPARAMETERVALUE_PASSWORDFREEDENIED = "InvalidParameterValue.PasswordFreeDenied"

	// 设置密码时，MC 传入的 old password 与先前设定密码不同。
	INVALIDPARAMETERVALUE_PASSWORDRULEERROR = "InvalidParameterValue.PasswordRuleError"

	// 请求容量偏小，不支持缩容。
	INVALIDPARAMETERVALUE_REDUCECAPACITYNOTALLOWED = "InvalidParameterValue.ReduceCapacityNotAllowed"

	// 复制组不存在。
	INVALIDPARAMETERVALUE_REPLICATIONGROUPNOTEXISTS = "InvalidParameterValue.ReplicationGroupNotExists"

	// 请求参数错误，安全组id错误。
	INVALIDPARAMETERVALUE_SECURITYGROUPIDSNOTEXISTS = "InvalidParameterValue.SecurityGroupIdsNotExists"

	// 实例规格不存在。
	INVALIDPARAMETERVALUE_SPECNOTEXIST = "InvalidParameterValue.SpecNotExist"

	// 实例类型不支持。
	INVALIDPARAMETERVALUE_UNSUPPORTEDTYPE = "InvalidParameterValue.UnSupportedType"

	// vpc网络下，uniqVpcId 子网id 非法。
	INVALIDPARAMETERVALUE_UNVPCIDNOTEXISTS = "InvalidParameterValue.UnVpcIdNotExists"

	// weekday输入无效数据。
	INVALIDPARAMETERVALUE_WEEKDAYSISINVALID = "InvalidParameterValue.WeekDaysIsInvalid"

	// 绑定超过上限。
	LIMITEXCEEDED_EXCEEDUPPERLIMIT = "LimitExceeded.ExceedUpperLimit"

	// 绑定实例必须为空。
	LIMITEXCEEDED_INSTANCENOTEMPTY = "LimitExceeded.InstanceNotEmpty"

	// 请求的容量不在售卖规格中（memSize应为1024的整数倍，单位：MB）。
	LIMITEXCEEDED_INVALIDMEMSIZE = "LimitExceeded.InvalidMemSize"

	// 一次请求购买的实例数不在售卖数量限制范围内。
	LIMITEXCEEDED_INVALIDPARAMETERGOODSNUMNOTINRANGE = "LimitExceeded.InvalidParameterGoodsNumNotInRange"

	// 请求的容量不在售卖容量范围内。
	LIMITEXCEEDED_MEMSIZENOTINRANGE = "LimitExceeded.MemSizeNotInRange"

	// 购买时长超过3年,请求时长超过最大时长。
	LIMITEXCEEDED_PERIODEXCEEDMAXLIMIT = "LimitExceeded.PeriodExceedMaxLimit"

	// 购买时长非法，时长最少1个月。
	LIMITEXCEEDED_PERIODLESSTHANMINLIMIT = "LimitExceeded.PeriodLessThanMinLimit"

	// 复制组已锁定。
	LIMITEXCEEDED_REPLICATIONGROUPLOCKED = "LimitExceeded.ReplicationGroupLocked"

	// 实例被其它流程锁住。
	RESOURCEINUSE_INSTANCEBEENLOCKED = "ResourceInUse.InstanceBeenLocked"

	// uin 值为空。
	RESOURCENOTFOUND_ACCOUNTDOESNOTEXISTS = "ResourceNotFound.AccountDoesNotExists"

	// 根据 serialId 没有找到对应 redis。
	RESOURCENOTFOUND_INSTANCENOTEXISTS = "ResourceNotFound.InstanceNotExists"

	// 找不到该实例。
	RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"

	// 请求订单号不存在。
	RESOURCEUNAVAILABLE_ACCOUNTBALANCENOTENOUGH = "ResourceUnavailable.AccountBalanceNotEnough"

	// 备份已被其它任务锁住，暂时不能执行该操作。
	RESOURCEUNAVAILABLE_BACKUPLOCKEDERROR = "ResourceUnavailable.BackupLockedError"

	// 备份状态异常，暂不能执行该操作。备份可能已过期或已被删除。
	RESOURCEUNAVAILABLE_BACKUPSTATUSABNORMAL = "ResourceUnavailable.BackupStatusAbnormal"

	// 调用后端接口失败。
	RESOURCEUNAVAILABLE_CALLOSSERROR = "ResourceUnavailable.CallOssError"

	// 获取安全组信息失败。
	RESOURCEUNAVAILABLE_GETSECURITYERROR = "ResourceUnavailable.GetSecurityError"

	// 实例配置错误。
	RESOURCEUNAVAILABLE_INSTANCECONFERROR = "ResourceUnavailable.InstanceConfError"

	// 实例已经被回收了。
	RESOURCEUNAVAILABLE_INSTANCEDELETED = "ResourceUnavailable.InstanceDeleted"

	// 实例已过期。
	RESOURCEUNAVAILABLE_INSTANCEISOLATED = "ResourceUnavailable.InstanceIsolated"

	// redis 已经被其它流程锁定。
	RESOURCEUNAVAILABLE_INSTANCELOCKEDERROR = "ResourceUnavailable.InstanceLockedError"

	// 订单不存在。
	RESOURCEUNAVAILABLE_INSTANCENODEAL = "ResourceUnavailable.InstanceNoDeal"

	// 实例状态不支持操作。
	RESOURCEUNAVAILABLE_INSTANCENOTSUPPORTOPERATION = "ResourceUnavailable.InstanceNotSupportOperation"

	// 实例状态错误。
	RESOURCEUNAVAILABLE_INSTANCESTATEERROR = "ResourceUnavailable.InstanceStateError"

	// redis 状态异常，不能执行对应流程。
	RESOURCEUNAVAILABLE_INSTANCESTATUSABNORMAL = "ResourceUnavailable.InstanceStatusAbnormal"

	// 实例状态异常，不能执行对应操作。
	RESOURCEUNAVAILABLE_INSTANCESTATUSERROR = "ResourceUnavailable.InstanceStatusError"

	// vpc网络IP资源不足。
	RESOURCEUNAVAILABLE_NOENOUGHVIPINVPC = "ResourceUnavailable.NoEnoughVipInVPC"

	// 请求的区域暂时不提供redis服务。
	RESOURCEUNAVAILABLE_NOREDISSERVICE = "ResourceUnavailable.NoRedisService"

	// 请求的区域暂时不提供请求类型的redis服务。
	RESOURCEUNAVAILABLE_NOTYPEIDREDISSERVICE = "ResourceUnavailable.NoTypeIdRedisService"

	// 产品还没有接入安全组。
	RESOURCEUNAVAILABLE_SECURITYGROUPNOTSUPPORTED = "ResourceUnavailable.SecurityGroupNotSupported"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 无cam 权限。
	UNAUTHORIZEDOPERATION_NOCAMAUTHED = "UnauthorizedOperation.NoCAMAuthed"

	// 用户不在白名单中。
	UNAUTHORIZEDOPERATION_USERNOTINWHITELIST = "UnauthorizedOperation.UserNotInWhiteList"

	// redis 集群版不允许接入安全组。
	UNSUPPORTEDOPERATION_CLUSTERINSTANCEACCESSEDDENY = "UnsupportedOperation.ClusterInstanceAccessedDeny"

	// 复制组巡检中。
	UNSUPPORTEDOPERATION_INSPECTION = "UnsupportedOperation.Inspection"

	// 不支持当前操作。
	UNSUPPORTEDOPERATION_INSTANCENOTOPERATION = "UnsupportedOperation.InstanceNotOperation"

	// 自动续费标识错误。
	UNSUPPORTEDOPERATION_ISAUTORENEWERROR = "UnsupportedOperation.IsAutoRenewError"

	// 只有集群版实例支持导出备份。
	UNSUPPORTEDOPERATION_ONLYCLUSTERINSTANCECANEXPORTBACKUP = "UnsupportedOperation.OnlyClusterInstanceCanExportBackup"
)
