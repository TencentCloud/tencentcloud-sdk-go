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

package v20260330

const (
	// 此产品的特有错误码

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 删除备份点时计费退还流水写入失败，余额未退还，已拒绝删除以保证账务一致。
	INTERNALERROR_BALANCERECORDFAILED = "InternalError.BalanceRecordFailed"

	// 计费订单已创建，但在发货等待窗口内未确认备份创建。备份可能仍在异步创建中，请稍后查询备份列表确认；请勿立即重试，以免重复下单。
	INTERNALERROR_BILLINGDELIVERYTIMEOUT = "InternalError.BillingDeliveryTimeout"

	// 云API超时，请重试
	INTERNALERROR_CLOUDAPITIMEOUT = "InternalError.CloudApiTimeout"

	// 第三方组件错误
	INTERNALERROR_COMPONENTERROR = "InternalError.ComponentError"

	// 校验对象存储（COS）存储桶时发生内部异常或网络超时（非桶不存在/无权限类入参错误）。
	INTERNALERROR_COSBUCKETCHECKFAILED = "InternalError.CosBucketCheckFailed"

	// COS 桶连通性检测失败
	INTERNALERROR_COSCONNECTIVITYERROR = "InternalError.CosConnectivityError"

	// COS 相关操作失败
	INTERNALERROR_COSERROR = "InternalError.CosError"

	// 查询CVM实例信息失败，请稍后重试
	INTERNALERROR_CVMQUERYFAILED = "InternalError.CvmQueryFailed"

	// 数据库并发操作繁忙，本次请求未能完成。
	INTERNALERROR_DATABASEDEADLOCK = "InternalError.DatabaseDeadlock"

	// 初始化备份库失败，请稍后重试
	INTERNALERROR_INITREPOSITORYFAILED = "InternalError.InitRepositoryFailed"

	// 部分备份点删除失败，请稍后重试
	INTERNALERROR_PARTIALDELETEFAILED = "InternalError.PartialDeleteFailed"

	// 恢复任务已受理并落库，任务投递结果正在确认中。
	INTERNALERROR_RESTORETASKENQUEUEPENDING = "InternalError.RestoreTaskEnqueuePending"

	// 查询备份内容失败，请稍后重试
	INTERNALERROR_SNAPSHOTTREEQUERYFAILED = "InternalError.SnapshotTreeQueryFailed"

	// 计费系统调用失败，请稍后重试
	INTERNALERROR_TRADEERROR = "InternalError.TradeError"

	// 无效的过滤器
	INVALIDFILTER = "InvalidFilter"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 大小不匹配
	INVALIDPARAMETER_DISKSIZENOTMATCH = "InvalidParameter.DiskSizeNotMatch"

	// 参数不合法
	INVALIDPARAMETER_INVALID = "InvalidParameter.Invalid"

	// 过滤参数非法
	INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"

	// 参数错误
	INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"

	// 参数错误
	INVALIDPARAMETER_INVALIDVALUE = "InvalidParameter.InvalidValue"

	// 至少需要提供一个要修改的字段
	INVALIDPARAMETER_MISSINGFIELD = "InvalidParameter.MissingField"

	// 参数缺失
	INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"

	// 参数错误
	INVALIDPARAMETER_PARAMETERERROR = "InvalidParameter.ParameterError"

	// 备份路径存在嵌套关系，请确保路径之间无父子目录包含关系
	INVALIDPARAMETER_PATHNESTED = "InvalidParameter.PathNested"

	// 备份策略不支持备份库模式
	INVALIDPARAMETER_POLICYNOTSUPPORTVAULT = "InvalidParameter.PolicyNotSupportVault"

	// 指定的策略不是文件备份策略，请选择正确的策略类型
	INVALIDPARAMETER_POLICYTYPEMISMATCH = "InvalidParameter.PolicyTypeMismatch"

	// 策略已绑定其他备份库，无法使用不同的备份库
	INVALIDPARAMETER_POLICYVAULTCONFLICT = "InvalidParameter.PolicyVaultConflict"

	// 盘与备份格式不匹配
	INVALIDPARAMETER_SHOULDCONVERTBACKUPTOIMAGE = "InvalidParameter.ShouldConvertBackupToImage"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// CVM实例不存在，请确认实例ID是否正确
	INVALIDPARAMETERVALUE_CVMINSTANCENOTEXIST = "InvalidParameterValue.CvmInstanceNotExist"

	// 冲突处理策略参数无效
	INVALIDPARAMETERVALUE_INVALIDCONFLICTSTRATEGY = "InvalidParameterValue.InvalidConflictStrategy"

	// 无效参数值。参数值格式错误或者参数值不被支持等
	INVALIDPARAMETERVALUE_INVALIDFORMAT = "InvalidParameterValue.InvalidFormat"

	// 网关实例ID格式错误
	INVALIDPARAMETERVALUE_INVALIDINSTANCEID = "InvalidParameterValue.InvalidInstanceId"

	// 排序方向无效，仅支持asc或desc
	INVALIDPARAMETERVALUE_INVALIDORDERDIRECTION = "InvalidParameterValue.InvalidOrderDirection"

	// 排序字段无效，当前仅支持按name排序
	INVALIDPARAMETERVALUE_INVALIDORDERFIELD = "InvalidParameterValue.InvalidOrderField"

	// 进度值超出有效范围[0.00-100.00]
	INVALIDPARAMETERVALUE_INVALIDPROGRESS = "InvalidParameterValue.InvalidProgress"

	// 作业状态转换无效，例如从终态转换到其他状态
	INVALIDPARAMETERVALUE_INVALIDSTATUSTRANSITION = "InvalidParameterValue.InvalidStatusTransition"

	// 参数值非法
	INVALIDPARAMETERVALUE_INVALIDVALUE = "InvalidParameterValue.InvalidValue"

	// 不合法的参数值，超过限制
	INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"

	// 恢复路径列表中存在超过系统长度限制的路径
	INVALIDPARAMETERVALUE_PATHTOOLONG = "InvalidParameterValue.PathTooLong"

	// 备份策略不存在或不可绑定
	INVALIDPARAMETERVALUE_POLICYNOTAVAILABLE = "InvalidParameterValue.PolicyNotAvailable"

	// 备份库当前状态不可写，请检查备份库是否正常
	INVALIDPARAMETERVALUE_VAULTSTATEERROR = "InvalidParameterValue.VaultStateError"

	// 备份库数量超过配额限制
	LIMITEXCEEDED_BACKUPVAULTOUTOFQUOTA = "LimitExceeded.BackupVaultOutOfQuota"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 必填参数缺失
	MISSINGPARAMETER_MISSINGREQUIRED = "MissingParameter.MissingRequired"

	// 备份桶正在删除中，请稍等。
	OPERATIONDENIED_BACKUPVAULTDELETING = "OperationDenied.BackupVaultDeleting"

	// 备份库中仍存在备份点
	OPERATIONDENIED_BACKUPVAULTHASFILEBACKUP = "OperationDenied.BackupVaultHasFileBackup"

	// 当前备份库存在备份策略绑定
	OPERATIONDENIED_BACKUPVAULTHASPOLICY = "OperationDenied.BackupVaultHasPolicy"

	// 资源回滚中
	RESOURCEINUSE_DISKROLLBACKING = "ResourceInUse.DiskRollbacking"

	// 资源繁忙
	RESOURCEINUSE_RESOURCEBUSY = "ResourceInUse.ResourceBusy"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 指定实例未找到对应的Agent，请确认Agent是否已安装
	RESOURCENOTFOUND_AGENTNOTINSTALLED = "ResourceNotFound.AgentNotInstalled"

	// 定期备份策略不存在
	RESOURCENOTFOUND_AUTOBACKUPPOLICYNOTFOUND = "ResourceNotFound.AutoBackupPolicyNotFound"

	// 指定的备份计划不存在
	RESOURCENOTFOUND_BACKUPPLANNOTEXIST = "ResourceNotFound.BackupPlanNotExist"

	// 备份源不存在，请确认资源 ID 或备份源视图 ID 是否正确。
	RESOURCENOTFOUND_BACKUPSOURCENOTEXIST = "ResourceNotFound.BackupSourceNotExist"

	// 指定的备份库不存在
	RESOURCENOTFOUND_BACKUPVAULTNOTFOUND = "ResourceNotFound.BackupVaultNotFound"

	// CFS 文件系统不存在或不可达。
	RESOURCENOTFOUND_CFSFILESYSTEMNOTFOUND = "ResourceNotFound.CfsFilesystemNotFound"

	// 复制对不存在
	RESOURCENOTFOUND_COPYPAIRNOTEXIST = "ResourceNotFound.CopyPairNotExist"

	// COS 存储桶不存在，请确认桶名称是否正确。
	RESOURCENOTFOUND_COSBUCKETNOTFOUND = "ResourceNotFound.CosBucketNotFound"

	// 保护组不存在
	RESOURCENOTFOUND_DISASTERRECOVERYPROTECTGROUPNOTEXIST = "ResourceNotFound.DisasterRecoveryProtectGroupNotExist"

	// 站点对ID不存在
	RESOURCENOTFOUND_DISASTERRECOVERYSITEPAIRNOTEXIST = "ResourceNotFound.DisasterRecoverySitePairNotExist"

	// 演练不存在
	RESOURCENOTFOUND_DRILLNOTEXIST = "ResourceNotFound.DrillNotExist"

	// 当前可用区域中的ENI池尚未配置
	RESOURCENOTFOUND_ENIPOOLNOTCONFIGURED = "ResourceNotFound.EniPoolNotConfigured"

	// 指定的备份点不存在，请确认备份点ID是否正确
	RESOURCENOTFOUND_FILEBACKUPNOTEXIST = "ResourceNotFound.FileBackupNotExist"

	// 指定的网关不存在或未注册
	RESOURCENOTFOUND_GATEWAY = "ResourceNotFound.Gateway"

	// 实例不存在
	RESOURCENOTFOUND_INSTANCENOTEXIST = "ResourceNotFound.InstanceNotExist"

	// 实例不存在
	RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"

	// 指定的作业不存在
	RESOURCENOTFOUND_JOB = "ResourceNotFound.Job"

	// 资源不存在
	RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"

	// 恢复目标不存在（或未授权为恢复目标），请确认后重试。
	RESOURCENOTFOUND_RESTORETARGETNOTEXIST = "ResourceNotFound.RestoreTargetNotExist"

	// 源端云服务器不存在
	RESOURCENOTFOUND_SOURCECVMNOTFOUND = "ResourceNotFound.SourceCVMNotFound"

	// 备份库不存在
	RESOURCENOTFOUND_VAULTNOTEXIST = "ResourceNotFound.VaultNotExist"

	// VPC映射不存在
	RESOURCENOTFOUND_VPCMAPPINGNOTEXIST = "ResourceNotFound.VpcMappingNotExist"

	// 资源正在创建备份
	RESOURCEUNAVAILABLE_BACKUPCREATING = "ResourceUnavailable.BackupCreating"

	// 当前端不支持此操作
	RESOURCEUNAVAILABLE_NOTSUPPORTINCURRENTSIDE = "ResourceUnavailable.NotSupportInCurrentSide"

	// 资源不支持当前操作
	RESOURCEUNAVAILABLE_NOTSUPPORTED = "ResourceUnavailable.NotSupported"

	// 云硬盘正在创建快照，请稍后重试
	RESOURCEUNAVAILABLE_SNAPSHOTCREATING = "ResourceUnavailable.SnapshotCreating"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 无 COS 存储桶访问权限，请确认已授权备份服务访问该桶。
	UNAUTHORIZEDOPERATION_COSBUCKETFORBIDDEN = "UnauthorizedOperation.CosBucketForbidden"

	// 未授权操作
	UNAUTHORIZEDOPERATION_HASNOSTRATEGY = "UnauthorizedOperation.HasNoStrategy"

	// 用户资质不满足要求，无法创建备份资源
	UNAUTHORIZEDOPERATION_LACKOFQUALIFICATION = "UnauthorizedOperation.LackOfQualification"

	// MFA认证过期
	UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"

	// 服务角色 BDRC_QCSLinkedRoleInCFSBackup 未创建
	UNAUTHORIZEDOPERATION_ROLENOTEXISTS = "UnauthorizedOperation.RoleNotExists"

	// Agent当前状态不是活跃状态，无法创建备份
	UNSUPPORTEDOPERATION_AGENTNOTACTIVE = "UnsupportedOperation.AgentNotActive"

	// 盘挂载在不同子机上
	UNSUPPORTEDOPERATION_BACKUPGROUPDISKATTACHMULTIINSTANCE = "UnsupportedOperation.BackupGroupDiskAttachMultiInstance"

	// 已有在途备份任务，请等待完成后再试。
	UNSUPPORTEDOPERATION_BACKUPJOBCONFLICT = "UnsupportedOperation.BackupJobConflict"

	// 同资源同策略已存在备份计划且备份目录不同，请通过修改计划实现目录调整。
	UNSUPPORTEDOPERATION_BACKUPPLANPATHSCONFLICT = "UnsupportedOperation.BackupPlanPathsConflict"

	// 备份桶状态异常，请稍后重试。
	UNSUPPORTEDOPERATION_BACKUPVAULTSTATEERROR = "UnsupportedOperation.BackupVaultStateError"

	// 复制对当前状态不支持该操作
	UNSUPPORTEDOPERATION_COPYPAIRSTATEERROR = "UnsupportedOperation.CopyPairStateError"

	// 容灾保护组已绑定资源
	UNSUPPORTEDOPERATION_DISASTERRECOVERYPROTECTGROUPBINDRESOURCE = "UnsupportedOperation.DisasterRecoveryProtectGroupBindResource"

	// 容灾保护组状态异常
	UNSUPPORTEDOPERATION_DISASTERRECOVERYPROTECTGROUPSTATEERROR = "UnsupportedOperation.DisasterRecoveryProtectGroupStateError"

	// 容灾站点对已绑定资源
	UNSUPPORTEDOPERATION_DISASTERRECOVERYSITEPAIRBINDRESOURCE = "UnsupportedOperation.DisasterRecoverySitePairBindResource"

	// 站点对状态错误
	UNSUPPORTEDOPERATION_DISASTERRECOVERYSITEPAIRSTATEERROR = "UnsupportedOperation.DisasterRecoverySitePairStateError"

	// 演练对状态异常
	UNSUPPORTEDOPERATION_DRILLPAIRSTATEERROR = "UnsupportedOperation.DrillPairStateError"

	// 路径冲突
	UNSUPPORTEDOPERATION_FILEBACKUPRESTORECONFLICT = "UnsupportedOperation.FileBackupRestoreConflict"

	// 备份点当前状态不支持恢复操作，请等待备份完成后再试
	UNSUPPORTEDOPERATION_FILEBACKUPSTATEERROR = "UnsupportedOperation.FileBackupStateError"

	// 指定路径已有创建中的备份任务，请等待完成后再试
	UNSUPPORTEDOPERATION_FILEBACKUPTASKDUPLICATE = "UnsupportedOperation.FileBackupTaskDuplicate"

	// 没有可用的保护点
	UNSUPPORTEDOPERATION_HAVENOPROTECTIONPOINT = "UnsupportedOperation.HaveNoProtectionPoint"

	// 实例存在恢复中的整机备份任务
	UNSUPPORTEDOPERATION_INSTANCESNAPSHOTROLLBACKING = "UnsupportedOperation.InstanceSnapshotRollbacking"

	// 恢复任务冲突
	UNSUPPORTEDOPERATION_RESTOREJOBCONFLICT = "UnsupportedOperation.RestoreJobConflict"

	// 备份快照不可用，无法浏览内容
	UNSUPPORTEDOPERATION_SNAPSHOTUNAVAILABLE = "UnsupportedOperation.SnapshotUnavailable"

	// 资源状态无法进行当前操作
	UNSUPPORTEDOPERATION_STATEERROR = "UnsupportedOperation.StateError"

	// 目标端资源正在回滚中
	UNSUPPORTEDOPERATION_TARGETRESOURCEROLLBACKING = "UnsupportedOperation.TargetResourceRollbacking"
)
