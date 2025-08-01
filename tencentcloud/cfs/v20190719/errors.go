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

package v20190719

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 获取CFS服务角色错误
	AUTHFAILURE_GETROLEFAILED = "AuthFailure.GetRoleFailed"

	// 请求未CAM授权。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 调整频率超限
	FAILEDOPERATION_ADJUSTFREQUENCYLIMIT = "FailedOperation.AdjustFrequencyLimit"

	// 文件系统绑定资源包失败。
	FAILEDOPERATION_BINDRESOURCEPKGFAILED = "FailedOperation.BindResourcePkgFailed"

	// 资源正在创建中。
	FAILEDOPERATION_CLIENTTOKENINUSE = "FailedOperation.ClientTokenInUse"

	// 文件系统存在挂载点。
	FAILEDOPERATION_MOUNTTARGETEXISTS = "FailedOperation.MountTargetExists"

	// 权限组已绑定文件系统。
	FAILEDOPERATION_PGROUPINUSE = "FailedOperation.PgroupInUse"

	// 权限组正在更新中。
	FAILEDOPERATION_PGROUPISUPDATING = "FailedOperation.PgroupIsUpdating"

	// 该权限组关联了旧版本实例，请解除关联后重试。
	FAILEDOPERATION_PGROUPLINKCFSV10 = "FailedOperation.PgroupLinkCfsv10"

	// 解绑资源标签失败。
	FAILEDOPERATION_UNTAGRESOURCEFAILED = "FailedOperation.UntagResourceFailed"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 创建文件系统失败。
	INTERNALERROR_CREATEFSFAILED = "InternalError.CreateFsFailed"

	// 获取用户费用状态失败。
	INTERNALERROR_GETACCOUNTSTATUSFAILED = "InternalError.GetAccountStatusFailed"

	// 超时。
	INTERNALERROR_TIMEOUT = "InternalError.Timeout"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 快照策略未找到。
	INVALIDPARAMETER_AUTOPOLICYNOTFOUND = "InvalidParameter.AutoPolicyNotFound"

	// 无效的快照保留时间。
	INVALIDPARAMETER_INVALIDALIVEDDAYS = "InvalidParameter.InvalidAlivedDays"

	// 定期星期参数无效。
	INVALIDPARAMETER_INVALIDPARAMDAYOFWEEK = "InvalidParameter.InvalidParamDayofWeek"

	// 定期小时 参数值错误。
	INVALIDPARAMETER_INVALIDPARAMHOUR = "InvalidParameter.InvalidParamHour"

	// 无效的快照策略状态。
	INVALIDPARAMETER_INVALIDSNAPPOLICYSTATUS = "InvalidParameter.InvalidSnapPolicyStatus"

	// 无效的文件系统快照参数名称 。
	INVALIDPARAMETER_INVALIDSNAPSHOTNAME = "InvalidParameter.InvalidSnapshotName"

	// 无效的文件系统快照策略名称。
	INVALIDPARAMETER_INVALIDSNAPSHOTPOLICYNAME = "InvalidParameter.InvalidSnapshotPolicyName"

	// 缺少策略相关参数。
	INVALIDPARAMETER_MISSINGPOLICYPARAM = "InvalidParameter.MissingPolicyParam"

	// 文件系统快照名称超出上限。
	INVALIDPARAMETER_SNAPSHOTNAMELIMITEXCEEDED = "InvalidParameter.SnapshotNameLimitExceeded"

	// 文件系统快照策略名称超过限制。
	INVALIDPARAMETER_SNAPSHOTPOLICYNAMELIMITEXCEEDED = "InvalidParameter.SnapshotPolicyNameLimitExceeded"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 自动扩容策略不存在
	INVALIDPARAMETERVALUE_AUTOPOLICYNOTFOUND = "InvalidParameterValue.AutoPolicyNotFound"

	// 参数值错误: 绑定归档规则前需绑定沉降规则
	INVALIDPARAMETERVALUE_BINDINFREQUENTACCESSFIRST = "InvalidParameterValue.BindInfrequentaccessFirst"

	// 无效的文件系统路径
	INVALIDPARAMETERVALUE_CFSPATH = "InvalidParameterValue.CfsPath"

	// 用于保证请求幂等性的字符串长度超过限制（不能超过64字节）。
	INVALIDPARAMETERVALUE_CLIENTTOKENLIMITEXCEEDED = "InvalidParameterValue.ClientTokenLimitExceeded"

	// 权限组名称重复。
	INVALIDPARAMETERVALUE_DUPLICATEDPGROUPNAME = "InvalidParameterValue.DuplicatedPgroupName"

	// 规则IP重复。
	INVALIDPARAMETERVALUE_DUPLICATEDRULEAUTHCLIENTIP = "InvalidParameterValue.DuplicatedRuleAuthClientIp"

	// 标签键重复
	INVALIDPARAMETERVALUE_DUPLICATEDTAGKEY = "InvalidParameterValue.DuplicatedTagKey"

	// 用户自定义名称过长（超过64字节)。
	INVALIDPARAMETERVALUE_FSNAMELIMITEXCEEDED = "InvalidParameterValue.FsNameLimitExceeded"

	// 文件系统配额设置超出上限。
	INVALIDPARAMETERVALUE_FSSIZELIMITEXCEEDED = "InvalidParameterValue.FsSizeLimitExceeded"

	// 无效的AliveDays值错误
	INVALIDPARAMETERVALUE_INVALIDALIVEDAYS = "InvalidParameterValue.InvalidAliveDays"

	// 规则IP错误。
	INVALIDPARAMETERVALUE_INVALIDAUTHCLIENTIP = "InvalidParameterValue.InvalidAuthClientIp"

	// 扩容策略参数无效
	INVALIDPARAMETERVALUE_INVALIDAUTOSCALEUPPARAMS = "InvalidParameterValue.InvalidAutoScaleUpParams"

	// 用于保证请求幂等性的字符串错误。
	INVALIDPARAMETERVALUE_INVALIDCLIENTTOKEN = "InvalidParameterValue.InvalidClientToken"

	// 参数值错误：数据流动ID无效
	INVALIDPARAMETERVALUE_INVALIDDATAFLOWID = "InvalidParameterValue.InvalidDataFlowId"

	// 无效的数据流动名称
	INVALIDPARAMETERVALUE_INVALIDDATAFLOWNAME = "InvalidParameterValue.InvalidDataFlowName"

	// 参数值错误：数据流动源信息无效,请检查密钥及路径
	INVALIDPARAMETERVALUE_INVALIDDATAFLOWSOURCEINFO = "InvalidParameterValue.InvalidDataFlowSourceInfo"

	// 参数值错误：数据流动源存储类型无效
	INVALIDPARAMETERVALUE_INVALIDDATAFLOWSOURCESTORAGETYPE = "InvalidParameterValue.InvalidDataFlowSourceStorageType"

	// 参数值错误：数据流动目标路径无效
	INVALIDPARAMETERVALUE_INVALIDDATAFLOWTARGETPATH = "InvalidParameterValue.InvalidDataFlowTargetPath"

	// 快照跨地域复制参数不支持该地域
	INVALIDPARAMETERVALUE_INVALIDDESTINATIONREGIONS = "InvalidParameterValue.InvalidDestinationRegions"

	// 加密参数错误。
	INVALIDPARAMETERVALUE_INVALIDENCRYPTED = "InvalidParameterValue.InvalidEncrypted"

	// FileSystemId无效。
	INVALIDPARAMETERVALUE_INVALIDFILESYSTEMID = "InvalidParameterValue.InvalidFileSystemId"

	// 无效的自定义名称。
	INVALIDPARAMETERVALUE_INVALIDFSNAME = "InvalidParameterValue.InvalidFsName"

	// 无效的文件系统配额。
	INVALIDPARAMETERVALUE_INVALIDFSSIZELIMIT = "InvalidParameterValue.InvalidFsSizeLimit"

	// 无效的文件系统状态。
	INVALIDPARAMETERVALUE_INVALIDFSSTATUS = "InvalidParameterValue.InvalidFsStatus"

	// 参数值错误：生命周期任务类型无效
	INVALIDPARAMETERVALUE_INVALIDLIFECYCLEDATATASKTYPE = "InvalidParameterValue.InvalidLifecycleDataTaskType"

	// 生命周期任务id 无效
	INVALIDPARAMETERVALUE_INVALIDLIFECYCLETASKID = "InvalidParameterValue.InvalidLifecycleTaskId"

	// MetaType 值无效
	INVALIDPARAMETERVALUE_INVALIDMETATYPE = "InvalidParameterValue.InvalidMetaType"

	// 错误的挂载点IP。
	INVALIDPARAMETERVALUE_INVALIDMOUNTTARGETIP = "InvalidParameterValue.InvalidMountTargetIp"

	// 无效的网络类型。
	INVALIDPARAMETERVALUE_INVALIDNETINTERFACE = "InvalidParameterValue.InvalidNetInterface"

	// DayOfMonth 参数错误
	INVALIDPARAMETERVALUE_INVALIDPARAMDAYOFMONTH = "InvalidParameterValue.InvalidParamDayOfMonth"

	// 无效的参数，DayOfWeek
	INVALIDPARAMETERVALUE_INVALIDPARAMDAYOFWEEK = "InvalidParameterValue.InvalidParamDayOfWeek"

	// IntervalDays 参数值错误
	INVALIDPARAMETERVALUE_INVALIDPARAMINTERVALDAYS = "InvalidParameterValue.InvalidParamIntervalDays"

	// 权限组不属于该用户。
	INVALIDPARAMETERVALUE_INVALIDPGROUP = "InvalidParameterValue.InvalidPgroup"

	// 权限组ID无效。
	INVALIDPARAMETERVALUE_INVALIDPGROUPID = "InvalidParameterValue.InvalidPgroupId"

	// 无效的权限组名称。
	INVALIDPARAMETERVALUE_INVALIDPGROUPNAME = "InvalidParameterValue.InvalidPgroupName"

	// 目录路径中包含无效字符。
	INVALIDPARAMETERVALUE_INVALIDPOLICYFSPATH = "InvalidParameterValue.InvalidPolicyFsPath"

	// 优先级设置错误。
	INVALIDPARAMETERVALUE_INVALIDPRIORITY = "InvalidParameterValue.InvalidPriority"

	// 协议参数错误。
	INVALIDPARAMETERVALUE_INVALIDPROTOCOL = "InvalidParameterValue.InvalidProtocol"

	// 用户区域选择错误 (ZoneName) 或 (ZoneId, Region)二者必选一。
	INVALIDPARAMETERVALUE_INVALIDREGIONZONEINFO = "InvalidParameterValue.InvalidRegionZoneInfo"

	// 参数值错误: 资源标签值错误。
	INVALIDPARAMETERVALUE_INVALIDRESOURCETAGS = "InvalidParameterValue.InvalidResourceTags"

	// 读写权限设置错误。
	INVALIDPARAMETERVALUE_INVALIDRWPERMISSION = "InvalidParameterValue.InvalidRwPermission"

	// 扩容目标容量参数错误
	INVALIDPARAMETERVALUE_INVALIDSCALEUPTARGETCAPACITY = "InvalidParameterValue.InvalidScaleupTargetCapacity"

	// 无效的快照策略状态
	INVALIDPARAMETERVALUE_INVALIDSNAPPOLICYSTATUS = "InvalidParameterValue.InvalidSnapPolicyStatus"

	// 无效的快照名称
	INVALIDPARAMETERVALUE_INVALIDSNAPSHOTNAME = "InvalidParameterValue.InvalidSnapshotName"

	// 无效的快照策略名称
	INVALIDPARAMETERVALUE_INVALIDSNAPSHOTPOLICYNAME = "InvalidParameterValue.InvalidSnapshotPolicyName"

	// 无效的快照状态。
	INVALIDPARAMETERVALUE_INVALIDSNAPSHOTSTATUS = "InvalidParameterValue.InvalidSnapshotStatus"

	// 资源包不存在或已绑定。
	INVALIDPARAMETERVALUE_INVALIDSTORAGERESOURCEPKG = "InvalidParameterValue.InvalidStorageResourcePkg"

	// 存储类型参数错误。
	INVALIDPARAMETERVALUE_INVALIDSTORAGETYPE = "InvalidParameterValue.InvalidStorageType"

	// 无效的子网ID。
	INVALIDPARAMETERVALUE_INVALIDSUBNETID = "InvalidParameterValue.InvalidSubnetId"

	// 标签键不能为空。
	INVALIDPARAMETERVALUE_INVALIDTAGKEY = "InvalidParameterValue.InvalidTagKey"

	// 标签值为空或字符无效。
	INVALIDPARAMETERVALUE_INVALIDTAGVALUE = "InvalidParameterValue.InvalidTagValue"

	// 无效的容量值。
	INVALIDPARAMETERVALUE_INVALIDTURBOCAPACITY = "InvalidParameterValue.InvalidTurboCapacity"

	// 用户权限设置错误。
	INVALIDPARAMETERVALUE_INVALIDUSERPERMISSION = "InvalidParameterValue.InvalidUserPermission"

	// 用户指定不可用的vip。
	INVALIDPARAMETERVALUE_INVALIDVIP = "InvalidParameterValue.InvalidVip"

	// 无效的VPCID。
	INVALIDPARAMETERVALUE_INVALIDVPCID = "InvalidParameterValue.InvalidVpcId"

	// VPC参数错误。
	INVALIDPARAMETERVALUE_INVALIDVPCPARAMETER = "InvalidParameterValue.InvalidVpcParameter"

	// 无效的可用区。
	INVALIDPARAMETERVALUE_INVALIDZONEID = "InvalidParameterValue.InvalidZoneId"

	// 无效的可用区或可用区ID。
	INVALIDPARAMETERVALUE_INVALIDZONEORZONEID = "InvalidParameterValue.InvalidZoneOrZoneId"

	// FileSystemId缺失。
	INVALIDPARAMETERVALUE_MISSINGFILESYSTEMID = "InvalidParameterValue.MissingFileSystemId"

	// 用户参数选择错误 (FileSystemId) 或 (Region)二者必选一。
	INVALIDPARAMETERVALUE_MISSINGFILESYSTEMIDORREGION = "InvalidParameterValue.MissingFileSystemIdOrRegion"

	// FileSystem参数缺失。
	INVALIDPARAMETERVALUE_MISSINGFSPARAMETER = "InvalidParameterValue.MissingFsParameter"

	// 密钥ID或ARN参数缺失。
	INVALIDPARAMETERVALUE_MISSINGKMSKEYID = "InvalidParameterValue.MissingKmsKeyId"

	// 权限组名称和描述不能同时为空。
	INVALIDPARAMETERVALUE_MISSINGNAMEORDESCINFO = "InvalidParameterValue.MissingNameOrDescinfo"

	// 权限组名称不能为空。
	INVALIDPARAMETERVALUE_MISSINGPGROUPNAME = "InvalidParameterValue.MissingPgroupName"

	// 缺少快照策略参数
	INVALIDPARAMETERVALUE_MISSINGPOLICYPARAM = "InvalidParameterValue.MissingPolicyParam"

	// 缺少参数Snapshot Name，或者aliveDays
	INVALIDPARAMETERVALUE_MISSINGSNAPNAMEORALIVEDAY = "InvalidParameterValue.MissingSnapNameOrAliveDay"

	// 未绑定存储包。
	INVALIDPARAMETERVALUE_MISSINGSTORAGERESOURCEPKG = "InvalidParameterValue.MissingStorageResourcePkg"

	// SUBNETID和UNSUBNETID不能同时为空。
	INVALIDPARAMETERVALUE_MISSINGSUBNETIDORUNSUBNETID = "InvalidParameterValue.MissingSubnetidOrUnsubnetid"

	// 参数错误，缺少taskid 或者region 至少输入一个
	INVALIDPARAMETERVALUE_MISSINGTASKIDORREGION = "InvalidParameterValue.MissingTaskIdOrRegion"

	// VPC相关参数缺失。
	INVALIDPARAMETERVALUE_MISSINGVPCPARAMETER = "InvalidParameterValue.MissingVpcParameter"

	// VPCID和UNVPCID不能同时为空。
	INVALIDPARAMETERVALUE_MISSINGVPCIDORUNVPCID = "InvalidParameterValue.MissingVpcidOrUnvpcid"

	// ZoneID缺失。
	INVALIDPARAMETERVALUE_MISSINGZONEID = "InvalidParameterValue.MissingZoneId"

	// 用户区域选择错误 (Zone) 或 (Zone_id)二者必选一。
	INVALIDPARAMETERVALUE_MISSINGZONEORZONEID = "InvalidParameterValue.MissingZoneOrZoneId"

	// 权限组描述长度超过限制（不能超过255字节）。
	INVALIDPARAMETERVALUE_PGROUPDESCINFOLIMITEXCEEDED = "InvalidParameterValue.PgroupDescinfoLimitExceeded"

	// 权限组名称长度超过限制（不能超过64字节）。
	INVALIDPARAMETERVALUE_PGROUPNAMELIMITEXCEEDED = "InvalidParameterValue.PgroupNameLimitExceeded"

	// 生命周期规则 类型错误
	INVALIDPARAMETERVALUE_POLICYRULESTORAGETYPEINVALID = "InvalidParameterValue.PolicyRuleStorageTypeInvalid"

	// 目录配额设置的目录嵌套
	INVALIDPARAMETERVALUE_PROJECTQUOTASDIRNESTED = "InvalidParameterValue.ProjectQuotasDirNested"

	// 目录配额数量超限
	INVALIDPARAMETERVALUE_PROJECTQUOTASEXCEEDED = "InvalidParameterValue.ProjectQuotasExceeded"

	// 容量硬限制取值范围错误。
	INVALIDPARAMETERVALUE_QUOTACAPLIMITERROR = "InvalidParameterValue.QuotaCapLimitError"

	// 文件硬限制取值范围错误。
	INVALIDPARAMETERVALUE_QUOTAFILELIMITERROR = "InvalidParameterValue.QuotaFileLimitError"

	// USER ID类型错误。
	INVALIDPARAMETERVALUE_QUOTAUSERIDERROR = "InvalidParameterValue.QuotaUserIdError"

	// 配额类型错误。
	INVALIDPARAMETERVALUE_QUOTAUSERTYPEERROR = "InvalidParameterValue.QuotaUserTypeError"

	// 权限组规则和权限组不匹配。
	INVALIDPARAMETERVALUE_RULENOTMATCHPGROUP = "InvalidParameterValue.RuleNotMatchPgroup"

	// 快照名称超限制，名称长度不超过64位字符
	INVALIDPARAMETERVALUE_SNAPSHOTNAMELIMITEXCEEDED = "InvalidParameterValue.SnapshotNameLimitExceeded"

	// 快照策略名称超过64位限制
	INVALIDPARAMETERVALUE_SNAPSHOTPOLICYNAMELIMITEXCEEDED = "InvalidParameterValue.SnapshotPolicyNameLimitExceeded"

	// 生命周期规则存储类型错误
	INVALIDPARAMETERVALUE_STORAGETYPEINVALID = "InvalidParameterValue.StorageTypeInvalid"

	// 参数值错误: 标签键个数超过上限（6个）。
	INVALIDPARAMETERVALUE_TAGKEYFILTERLIMITEXCEEDED = "InvalidParameterValue.TagKeyFilterLimitExceeded"

	// 标签键长度超过限制（不能超过127字节）。
	INVALIDPARAMETERVALUE_TAGKEYLIMITEXCEEDED = "InvalidParameterValue.TagKeyLimitExceeded"

	// 标签值个数超过上限（10个）。
	INVALIDPARAMETERVALUE_TAGVALUEFILTERLIMITEXCEEDED = "InvalidParameterValue.TagValueFilterLimitExceeded"

	// 标签值长度超过限制（不能超过255字节）。
	INVALIDPARAMETERVALUE_TAGVALUELIMITEXCEEDED = "InvalidParameterValue.TagValueLimitExceeded"

	// 该可用区无法提供服务。
	INVALIDPARAMETERVALUE_UNAVAILABLEREGION = "InvalidParameterValue.UnavailableRegion"

	// 该地域无法提供服务。
	INVALIDPARAMETERVALUE_UNAVAILABLEZONE = "InvalidParameterValue.UnavailableZone"

	// 生命周期数据任务数超限
	INVALIDPARAMETERVALUE_WAITINGTASKLIMITEXCEEDED = "InvalidParameterValue.WaitingTaskLimitExceeded"

	// ZoneId和Region不匹配。
	INVALIDPARAMETERVALUE_ZONEIDREGIONNOTMATCH = "InvalidParameterValue.ZoneIdRegionNotMatch"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 资源不足：数据流动数超限
	RESOURCEINSUFFICIENT_DATAFLOWLIMITEXCEEDED = "ResourceInsufficient.DataFlowLimitExceeded"

	// 文件系统数量达到上限。
	RESOURCEINSUFFICIENT_FILESYSTEMLIMITEXCEEDED = "ResourceInsufficient.FileSystemLimitExceeded"

	// 权限组数量达到上限。
	RESOURCEINSUFFICIENT_PGROUPNUMBERLIMITEXCEEDED = "ResourceInsufficient.PgroupNumberLimitExceeded"

	// 策略绑定文件系统数量超过限制。
	RESOURCEINSUFFICIENT_POLICYFSLIMITEXCEEDED = "ResourceInsufficient.PolicyFsLimitExceeded"

	// 策略绑定路径数量超过限制。
	RESOURCEINSUFFICIENT_POLICYFSPATHLIMITEXCEEDED = "ResourceInsufficient.PolicyFsPathLimitExceeded"

	// 该用户的生命周期策略个数达到最大限制。
	RESOURCEINSUFFICIENT_POLICYLIMITEXCEEDED = "ResourceInsufficient.PolicyLimitExceeded"

	// 区域资源售罄。
	RESOURCEINSUFFICIENT_REGIONSOLDOUT = "ResourceInsufficient.RegionSoldOut"

	// 规则条数超过上限。
	RESOURCEINSUFFICIENT_RULELIMITEXCEEDED = "ResourceInsufficient.RuleLimitExceeded"

	// 文件系统过大，不支持快照
	RESOURCEINSUFFICIENT_SNAPSHOTSIZELIMITEXCEEDED = "ResourceInsufficient.SnapshotSizeLimitExceeded"

	// 该子网下IP数量不足
	RESOURCEINSUFFICIENT_SUBNETIPALLOCCUPIED = "ResourceInsufficient.SubnetIpAllOccupied"

	// 该资源的标签个数达到最大限制。
	RESOURCEINSUFFICIENT_TAGLIMITEXCEEDED = "ResourceInsufficient.TagLimitExceeded"

	// 标签限额不足。
	RESOURCEINSUFFICIENT_TAGQUOTASEXCEEDED = "ResourceInsufficient.TagQuotasExceeded"

	// Turbo类型特殊容量文件系统数量限制
	RESOURCEINSUFFICIENT_TURBOSPECIALCAPACITYFILESYSTEMCOUNTLIMIT = "ResourceInsufficient.TurboSpecialCapacityFileSystemCountLimit"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 该文件系统不存在。
	RESOURCENOTFOUND_FILESYSTEMNOTFOUND = "ResourceNotFound.FileSystemNotFound"

	// 文件系统不存在。
	RESOURCENOTFOUND_FSNOTEXIST = "ResourceNotFound.FsNotExist"

	// 挂载点不存在。
	RESOURCENOTFOUND_MOUNTTARGETNOTFOUND = "ResourceNotFound.MountTargetNotFound"

	// 权限组不存在。
	RESOURCENOTFOUND_PGROUPNOTFOUND = "ResourceNotFound.PgroupNotFound"

	// 资源不存在：资源ID 不存在。
	RESOURCENOTFOUND_RESOURCEPACKAGENOTFOUND = "ResourceNotFound.ResourcePackageNotFound"

	// 权限规则不存在。
	RESOURCENOTFOUND_RULENOTFOUND = "ResourceNotFound.RuleNotFound"

	// 快照ID 不存在。
	RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 该可用区不支持基础网络。
	UNSUPPORTEDOPERATION_BASICNETINTERFACENOTSUPPORTED = "UnsupportedOperation.BasicNetInterfaceNotSupported"

	// 该数据流动不支持删除，关联该数据流动的数据管理任务正在执行中，请等待执行完成或终止执行。
	UNSUPPORTEDOPERATION_INVALIDLIFECYCLEDATATASKSTATUS = "UnsupportedOperation.InvalidLifecycleDataTaskStatus"

	// appid不在kms白名单(kms目前是白名单开启状态)。
	UNSUPPORTEDOPERATION_MISSINGKMSACCESSPERMISSION = "UnsupportedOperation.MissingKmsAccessPermission"

	// 用户已欠费, 请充值后重试。
	UNSUPPORTEDOPERATION_OUTOFSERVICE = "UnsupportedOperation.OutOfService"

	// cfs角色未被授权。
	UNSUPPORTEDOPERATION_UNAUTHORIZEDCFSQCSROLE = "UnsupportedOperation.UnauthorizedCfsQcsRole"

	// 用户未经过实名认证。
	UNSUPPORTEDOPERATION_UNVERIFIEDUSER = "UnsupportedOperation.UnverifiedUser"
)
