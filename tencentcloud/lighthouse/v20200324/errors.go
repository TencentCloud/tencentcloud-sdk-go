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

package v20200324

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 该地域尚未开放，请选择其他地域。
	AUTHFAILURE_INVALIDREGION = "AuthFailure.InvalidRegion"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 创建镜像失败。
	FAILEDOPERATION_CREATEBLUEPRINTFAILED = "FailedOperation.CreateBlueprintFailed"

	// 创建云硬盘失败。
	FAILEDOPERATION_CREATEDISKSFAILED = "FailedOperation.CreateDisksFailed"

	// 创建实例失败。
	FAILEDOPERATION_CREATEINSTANCESFAILED = "FailedOperation.CreateInstancesFailed"

	// 对密钥对的创建操作失败。
	FAILEDOPERATION_CREATEKEYPAIRFAILED = "FailedOperation.CreateKeyPairFailed"

	// 对密钥对的删除操作失败。
	FAILEDOPERATION_DELETEKEYPAIRFAILED = "FailedOperation.DeleteKeyPairFailed"

	// 查询镜像失败，请稍后再试。
	FAILEDOPERATION_DESCRIBEBLUEPRINTSFAILED = "FailedOperation.DescribeBlueprintsFailed"

	// 查询套餐折扣失败
	FAILEDOPERATION_DESCRIBEBUNDLEDISCOUNTFAILED = "FailedOperation.DescribeBundleDiscountFailed"

	// 查询套餐失败
	FAILEDOPERATION_DESCRIBEBUNDLESFAILED = "FailedOperation.DescribeBundlesFailed"

	// 查询云硬盘是否可以退还操作失败。
	FAILEDOPERATION_DESCRIBEDISKSRETURNABLEERROR = "FailedOperation.DescribeDisksReturnableError"

	// 查询实例状态错误。
	FAILEDOPERATION_DESCRIBEINSTANCESTATUS = "FailedOperation.DescribeInstanceStatus"

	// 查询实例变配套餐失败。
	FAILEDOPERATION_DESCRIBEINSTANCESMODIFICATIONERROR = "FailedOperation.DescribeInstancesModificationError"

	// 查询实例退还错误。
	FAILEDOPERATION_DESCRIBEINSTANCESRETURNABLEERROR = "FailedOperation.DescribeInstancesReturnableError"

	// 查询流量包失败。
	FAILEDOPERATION_DESCRIBEINSTANCESTRAFFICPACKAGESFAILED = "FailedOperation.DescribeInstancesTrafficPackagesFailed"

	// 查询资源返回了不符合要求内容。
	FAILEDOPERATION_DESCRIBERESOURCESRETURNABLEERROR = "FailedOperation.DescribeResourcesReturnableError"

	// 销毁资源失败，请稍后重新操作。
	FAILEDOPERATION_DESTROYRESOURCESFAILED = "FailedOperation.DestroyResourcesFailed"

	// 容器列表过长。
	FAILEDOPERATION_DOCKERCONTAINERSLISTTOOLARGE = "FailedOperation.DockerContainersListTooLarge"

	// 指定Docker环境操作失败, 请检查Docker环境。
	FAILEDOPERATION_DOCKEROPERATIONFAILED = "FailedOperation.DockerOperationFailed"

	// 对防火墙规则的操作失败。
	FAILEDOPERATION_FIREWALLRULESOPERATIONFAILED = "FailedOperation.FirewallRulesOperationFailed"

	// 对密钥对的导入操作失败。
	FAILEDOPERATION_IMPORTKEYPAIRFAILED = "FailedOperation.ImportKeyPairFailed"

	// 对实例的操作失败。
	FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"

	// 账户余额不足, 请及时充值。
	FAILEDOPERATION_INSUFFICIENTBALANCE = "FailedOperation.InsufficientBalance"

	// 命令无法找到。
	FAILEDOPERATION_INVALIDCOMMANDNOTFOUND = "FailedOperation.InvalidCommandNotFound"

	// 退还资源失败。
	FAILEDOPERATION_ISOLATERESOURCESFAILED = "FailedOperation.IsolateResourcesFailed"

	// 变更实例套餐失败。
	FAILEDOPERATION_MODIFYINSTANCESBUNDLEFAILED = "FailedOperation.ModifyInstancesBundleFailed"

	// 变更资源属性失败，请稍后重新操作。
	FAILEDOPERATION_MODIFYRESOURCESATTRIBUTEFAILED = "FailedOperation.ModifyResourcesAttributeFailed"

	// 修改资源自动续费失败
	FAILEDOPERATION_MODIFYRESOURCESRENEWFLAGFAILED = "FailedOperation.ModifyResourcesRenewFlagFailed"

	// 续费资源失败。
	FAILEDOPERATION_RENEWRESOURCESFAILED = "FailedOperation.RenewResourcesFailed"

	// 请求错误。
	FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"

	// 扩容云硬盘失败，请稍后重新操作。
	FAILEDOPERATION_RESIZEDISKSFAILED = "FailedOperation.ResizeDisksFailed"

	// 跨账号共享镜像失败，请稍后重试。
	FAILEDOPERATION_SHAREBLUEPRINTACROSSACCOUNTFAILED = "FailedOperation.ShareBlueprintAcrossAccountFailed"

	// 快照操作失败。
	FAILEDOPERATION_SNAPSHOTOPERATIONFAILED = "FailedOperation.SnapshotOperationFailed"

	// TAT命令未完成。
	FAILEDOPERATION_TATINVOCATIONNOTFINISHED = "FailedOperation.TATInvocationNotFinished"

	// 调用计费网关服务失败，请稍后重新操作。
	FAILEDOPERATION_TRADECALLBILLINGGATEWAYFAILED = "FailedOperation.TradeCallBillingGatewayFailed"

	// 计费询价失败。
	FAILEDOPERATION_TRADEGETPRICEFAILED = "FailedOperation.TradeGetPriceFailed"

	// 操作失败，不能创建自定义镜像。
	FAILEDOPERATION_UNABLETOCREATEBLUEPRINT = "FailedOperation.UnableToCreateBlueprint"

	// 无法创建实例。
	FAILEDOPERATION_UNABLETOCREATEINSTANCES = "FailedOperation.UnableToCreateInstances"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 查询实例状态失败，请稍后重试。
	INTERNALERROR_DESCRIBEINSTANCESTATUS = "InternalError.DescribeInstanceStatus"

	// 查询实例是否可变配失败。
	INTERNALERROR_DESCRIBEINSTANCESMODIFICATION = "InternalError.DescribeInstancesModification"

	// 查询实例是否可变配失败。
	INTERNALERROR_DESCRIBEINSTANCESMODIFICATIONERROR = "InternalError.DescribeInstancesModificationError"

	// 查询实例是否可退还失败。
	INTERNALERROR_DESCRIBEINSTANCESRETURNABLEERROR = "InternalError.DescribeInstancesReturnableError"

	// 查询实例流量包错误。
	INTERNALERROR_DESCRIBEINSTANCESTRAFFICPACKAGESFAILED = "InternalError.DescribeInstancesTrafficPackagesFailed"

	// 查询资源返回了不符合要求内容。
	INTERNALERROR_DESCRIBERESOURCESRETURNABLEERROR = "InternalError.DescribeResourcesReturnableError"

	// 快照配额锁获取失败。
	INTERNALERROR_GETSNAPSHOTALLOCQUOTALOCKERROR = "InternalError.GetSnapshotAllocQuotaLockError"

	// 无法找到此接口。
	INTERNALERROR_INVALIDACTIONNOTFOUND = "InternalError.InvalidActionNotFound"

	// 套餐价格错误。
	INTERNALERROR_INVALIDBUNDLEPRICE = "InternalError.InvalidBundlePrice"

	// 命令无法找到。
	INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"

	// 请求出现错误。
	INTERNALERROR_REQUESTERROR = "InternalError.RequestError"

	// 调用计费网关服务失败。
	INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"

	// 价格获取失败。
	INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 套餐和镜像不匹配。
	INVALIDPARAMETER_BUNDLEANDBLUEPRINTNOTMATCH = "InvalidParameter.BundleAndBlueprintNotMatch"

	// 产品未定义的套餐 ID。
	INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"

	// 参数冲突。
	INVALIDPARAMETER_CONFLICT = "InvalidParameter.Conflict"

	// 参数非法，Filter 参数中的 Values 取值数量超过允许的最大数量。
	INVALIDPARAMETER_FILTERVALUELIMITEXCEEDED = "InvalidParameter.FilterValueLimitExceeded"

	// 参数非法，防火墙规则重复。
	INVALIDPARAMETER_FIREWALLRULESDUPLICATED = "InvalidParameter.FirewallRulesDuplicated"

	// 参数非法，防火墙规则已存在。
	INVALIDPARAMETER_FIREWALLRULESEXIST = "InvalidParameter.FirewallRulesExist"

	// 参数非法，Filter 参数非法。
	INVALIDPARAMETER_INVALIDFILTER = "InvalidParameter.InvalidFilter"

	// 参数非法，Filter 参数中的 Name 取值非法。
	INVALIDPARAMETER_INVALIDFILTERINVALIDKEY = "InvalidParameter.InvalidFilterInvalidKey"

	// 参数非法，Filter 参数中的 Name 取值不是字符串。
	INVALIDPARAMETER_INVALIDFILTERINVALIDNAMENOTSTR = "InvalidParameter.InvalidFilterInvalidNameNotStr"

	// 参数非法，Filter 参数中的 Values 不是列表。
	INVALIDPARAMETER_INVALIDFILTERINVALIDVALUESNOTLIST = "InvalidParameter.InvalidFilterInvalidValuesNotList"

	// Filter参数名称不正确。
	INVALIDPARAMETER_INVALIDFILTERNAME = "InvalidParameter.InvalidFilterName"

	// 参数非法，Filter 参数不是字典。
	INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"

	// 参数非法，Filter 参数中有不支持的 Name。
	INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"

	// 必须要指定一个要修改的属性。
	INVALIDPARAMETER_MUSTSPECIFYONEATTRIBUTETOMODIFY = "InvalidParameter.MustSpecifyOneAttributeToModify"

	// 参数非法，每次只能修改一个属性。
	INVALIDPARAMETER_ONLYALLOWMODIFYONEATTRIBUTE = "InvalidParameter.OnlyAllowModifyOneAttribute"

	// 参数非法，参数冲突。
	INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 镜像不允许共享至不同站点的账号。
	INVALIDPARAMETERVALUE_ACCOUNTIDINVALIDACCOUNTAREA = "InvalidParameterValue.AccountIdInvalidAccountArea"

	// 账号为当前用户。
	INVALIDPARAMETERVALUE_ACCOUNTIDSAMEWITHUIN = "InvalidParameterValue.AccountIdSameWithUin"

	// 账号ID不存在。
	INVALIDPARAMETERVALUE_ACCOUNTIDSNOTEXIST = "InvalidParameterValue.AccountIdsNotExist"

	// 账号ID不为主账号。
	INVALIDPARAMETERVALUE_ACCOUNTIDSNOTOWNERACCOUNT = "InvalidParameterValue.AccountIdsNotOwnerAccount"

	// 该实例配置不符合指定镜像的要求。
	INVALIDPARAMETERVALUE_BLUEPRINTCONFIGNOTMATCH = "InvalidParameterValue.BlueprintConfigNotMatch"

	// 镜像 ID 不合法，重装实例不允许切换操作系统类型。
	INVALIDPARAMETERVALUE_BLUEPRINTID = "InvalidParameterValue.BlueprintId"

	// 参数值非法，镜像 ID 格式非法。
	INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"

	// 套餐和镜像不匹配。
	INVALIDPARAMETERVALUE_BUNDLEANDBLUEPRINTNOTMATCH = "InvalidParameterValue.BundleAndBlueprintNotMatch"

	// 所选套餐不支持镜像的操作系统平台类型。
	INVALIDPARAMETERVALUE_BUNDLENOTSUPPORTBLUEPRINTPLATFORM = "InvalidParameterValue.BundleNotSupportBlueprintPlatform"

	// 云联网实例ID格式非法。
	INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"

	// 客户令牌长度超出限制。
	INVALIDPARAMETERVALUE_CLIENTTOKENTOOLONG = "InvalidParameterValue.ClientTokenTooLong"

	// 参数值非法，云硬盘备份点 ID 格式非法。
	INVALIDPARAMETERVALUE_DISKBACKUPIDMALFORMED = "InvalidParameterValue.DiskBackupIdMalformed"

	// 指定的云盘备份点名称不可大于最大长度。
	INVALIDPARAMETERVALUE_DISKBACKUPNAMETOOLONG = "InvalidParameterValue.DiskBackupNameTooLong"

	// 云硬盘备份点配额小于当前云硬盘备份点数量。
	INVALIDPARAMETERVALUE_DISKBACKUPQUOTALESSTHENCURRENTDISKBACKUPNUM = "InvalidParameterValue.DiskBackupQuotaLessThenCurrentDiskBackupNum"

	// 云硬盘的可用区与实例的可用区不匹配。
	INVALIDPARAMETERVALUE_DISKINSTANCEZONENOTMATCH = "InvalidParameterValue.DiskInstanceZoneNotMatch"

	// 磁盘名称长度超出限制。
	INVALIDPARAMETERVALUE_DISKNAMETOOLONG = "InvalidParameterValue.DiskNameTooLong"

	// 磁盘大小发生改变。
	INVALIDPARAMETERVALUE_DISKSIZENOTMATCH = "InvalidParameterValue.DiskSizeNotMatch"

	// 指定云硬盘大小小于当前云硬盘大小。
	INVALIDPARAMETERVALUE_DISKSIZESMALLERTHANCURRENTDISKSIZE = "InvalidParameterValue.DiskSizeSmallerThanCurrentDiskSize"

	// 参数 `KeyName` 已经存在且重复。
	INVALIDPARAMETERVALUE_DUPLICATEPARAMETERVALUE = "InvalidParameterValue.DuplicateParameterValue"

	// 参数值非法，不允许包含重复的值。
	INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"

	// 防火墙模板规则已存在
	INVALIDPARAMETERVALUE_DUPLICATEDFIREWALLTEMPLATERULE = "InvalidParameterValue.DuplicatedFirewallTemplateRule"

	// 列值不正确。
	INVALIDPARAMETERVALUE_FIELDSCOMPARE = "InvalidParameterValue.FieldsCompare"

	// 防火墙规则描述长度超出限制。
	INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"

	// 防火墙模板ID格式非法。
	INVALIDPARAMETERVALUE_FIREWALLTEMPLATEIDMALFORMED = "InvalidParameterValue.FirewallTemplateIdMalformed"

	// 参数值非法，实例 ID 格式非法。
	INVALIDPARAMETERVALUE_INSTANCEIDMALFORMED = "InvalidParameterValue.InstanceIdMalformed"

	// 参数值非法，实例名称超过允许的最大长度。
	INVALIDPARAMETERVALUE_INSTANCENAMETOOLONG = "InvalidParameterValue.InstanceNameTooLong"

	// 镜像 ID 不合法。
	INVALIDPARAMETERVALUE_INVALIDBLUEPRINTID = "InvalidParameterValue.InvalidBlueprintId"

	// 镜像操作系统类型不合法。
	INVALIDPARAMETERVALUE_INVALIDBLUEPRINTPLATFORMTYPE = "InvalidParameterValue.InvalidBlueprintPlatformType"

	// 镜像状态取值非法。
	INVALIDPARAMETERVALUE_INVALIDBLUEPRINTSTATE = "InvalidParameterValue.InvalidBlueprintState"

	// 镜像类型不合法。
	INVALIDPARAMETERVALUE_INVALIDBLUEPRINTTYPE = "InvalidParameterValue.InvalidBlueprintType"

	// 非法的套餐参数。
	INVALIDPARAMETERVALUE_INVALIDBUNDLE = "InvalidParameterValue.InvalidBundle"

	// 套餐和镜像的组合非法。
	INVALIDPARAMETERVALUE_INVALIDBUNDLEBLUEPRINTCOMBINATION = "InvalidParameterValue.InvalidBundleBlueprintCombination"

	// 控制台显示类型不合法。
	INVALIDPARAMETERVALUE_INVALIDCONSOLEDISPLAYTYPES = "InvalidParameterValue.InvalidConsoleDisplayTypes"

	// 当前实例到期时间不能早于云硬盘到期时间。
	INVALIDPARAMETERVALUE_INVALIDCURINSTANCEDEADLINE = "InvalidParameterValue.InvalidCurInstanceDeadline"

	// 参数值非法，磁盘 ID 格式非法。
	INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"

	// 指定云硬盘大小不符合要求。
	INVALIDPARAMETERVALUE_INVALIDDISKSIZE = "InvalidParameterValue.InvalidDiskSize"

	// 云硬盘类型非法。
	INVALIDPARAMETERVALUE_INVALIDDISKTYPE = "InvalidParameterValue.InvalidDiskType"

	// 设置是否使用默认密钥对登录的值不正确。
	INVALIDPARAMETERVALUE_INVALIDINSTANCELOGINKEYPAIRPERMITLOGIN = "InvalidParameterValue.InvalidInstanceLoginKeyPairPermitLogin"

	// 参数值非法，IP 地址格式非法。
	INVALIDPARAMETERVALUE_INVALIDIPFORMAT = "InvalidParameterValue.InvalidIpFormat"

	// 参数值非法。
	INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMEEMPTY = "InvalidParameterValue.InvalidKeyPairNameEmpty"

	// 非法的密钥对名称。
	INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMEINCLUDEILLEGALCHAR = "InvalidParameterValue.InvalidKeyPairNameIncludeIllegalChar"

	// 参数长度非法。
	INVALIDPARAMETERVALUE_INVALIDKEYPAIRNAMETOOLONG = "InvalidParameterValue.InvalidKeyPairNameTooLong"

	// 参数组合非法。
	INVALIDPARAMETERVALUE_INVALIDPARAMETERCOMBINATION = "InvalidParameterValue.InvalidParameterCombination"

	// 参数中的密码不合法。
	INVALIDPARAMETERVALUE_INVALIDPASSWORD = "InvalidParameterValue.InvalidPassword"

	// 不正确的配额资源名称。
	INVALIDPARAMETERVALUE_INVALIDRESOURCEQUOTARESOURCENAME = "InvalidParameterValue.InvalidResourceQuotaResourceName"

	// 使用场景Id不合法。
	INVALIDPARAMETERVALUE_INVALIDSCENEIDMALFORMED = "InvalidParameterValue.InvalidSceneIdMalformed"

	// 参数Zone的取值不合法。
	INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"

	// 参数值非法，密钥对 ID 格式非法。
	INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"

	// 该密钥对中的公钥在系统中已存在，无法重复使用 。
	INVALIDPARAMETERVALUE_KEYPAIRPUBLICKEYDUPLICATED = "InvalidParameterValue.KeyPairPublicKeyDuplicated"

	// 指定的公钥格式错误。
	INVALIDPARAMETERVALUE_KEYPAIRPUBLICKEYMALFORMED = "InvalidParameterValue.KeyPairPublicKeyMalformed"

	// 参数取值长度不合法。
	INVALIDPARAMETERVALUE_LENGTH = "InvalidParameterValue.Length"

	// 参数值非法，参数值的数量超过最大限制。
	INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"

	// 参数值非法，不能为负值。
	INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"

	// 不允许改变平台类型。
	INVALIDPARAMETERVALUE_NOTALLOWTOCHANGEPLATFORMTYPE = "InvalidParameterValue.NotAllowToChangePlatformType"

	// 实例套餐的套餐类型不支持表更至新套餐。
	INVALIDPARAMETERVALUE_NOTSUPPORTMODIFYINSTANCEBUNDLETYPE = "InvalidParameterValue.NotSupportModifyInstanceBundleType"

	// 参数值非法，不在合法范围内。
	INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"

	// 实例操作系统不支持该文件系统。
	INVALIDPARAMETERVALUE_PLATFORMTYPENOTSUPPORTFILESYSTEM = "InvalidParameterValue.PlatformTypeNotSupportFileSystem"

	// 实例操作系统不支持指定挂载点。
	INVALIDPARAMETERVALUE_PLATFORMTYPENOTSUPPORTMOUNTPOINT = "InvalidParameterValue.PlatformTypeNotSupportMountPoint"

	// 地域不存在。
	INVALIDPARAMETERVALUE_REGIONNOTFOUND = "InvalidParameterValue.RegionNotFound"

	// 地域不匹配。
	INVALIDPARAMETERVALUE_REGIONNOTMATCH = "InvalidParameterValue.RegionNotMatch"

	// 不支持的地域。
	INVALIDPARAMETERVALUE_REGIONNOTSUPPORTED = "InvalidParameterValue.RegionNotSupported"

	// 该地域不可用。
	INVALIDPARAMETERVALUE_REGIONUNAVAILABLE = "InvalidParameterValue.RegionUnavailable"

	// 参数值非法，快照 ID 格式非法。
	INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"

	// 参数值非法，快照名称超过允许的最大长度。
	INVALIDPARAMETERVALUE_SNAPSHOTNAMETOOLONG = "InvalidParameterValue.SnapshotNameTooLong"

	// 参数值非法，大于有效值。
	INVALIDPARAMETERVALUE_TOOLARGE = "InvalidParameterValue.TooLarge"

	// 参数取值过长，超过最大长度。
	INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"

	// 参数值非法，小于有效值。
	INVALIDPARAMETERVALUE_TOOSMALL = "InvalidParameterValue.TooSmall"

	// 非法的可用区。
	INVALIDPARAMETERVALUE_ZONEINVALID = "InvalidParameterValue.ZoneInvalid"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 实例挂载数据盘配额不足，无法挂载云硬盘。
	LIMITEXCEEDED_ATTACHDATADISKQUOTALIMITEXCEEDED = "LimitExceeded.AttachDataDiskQuotaLimitExceeded"

	// 配额不足，当前自定义镜像配额不允许创建新的自定义镜像。
	LIMITEXCEEDED_BLUEPRINTQUOTALIMITEXCEEDED = "LimitExceeded.BlueprintQuotaLimitExceeded"

	// 超过磁盘备份点配额限制。
	LIMITEXCEEDED_DISKBACKUPQUOTALIMITEXCEEDED = "LimitExceeded.DiskBackupQuotaLimitExceeded"

	// 当前配额不足，无法创建新的云硬盘。
	LIMITEXCEEDED_DISKQUOTALIMITEXCEEDED = "LimitExceeded.DiskQuotaLimitExceeded"

	// 超过防火墙规则配额。
	LIMITEXCEEDED_FIREWALLRULESLIMITEXCEEDED = "LimitExceeded.FirewallRulesLimitExceeded"

	// 防火墙模板规则超出配额
	LIMITEXCEEDED_FIREWALLTEMPLATERULEQUOTALIMITEXCEEDED = "LimitExceeded.FirewallTemplateRuleQuotaLimitExceeded"

	// 超过实例配额。
	LIMITEXCEEDED_INSTANCEQUOTALIMITEXCEEDED = "LimitExceeded.InstanceQuotaLimitExceeded"

	// 退还资源数量超出限制。
	LIMITEXCEEDED_ISOLATERESOURCESLIMITEXCEEDED = "LimitExceeded.IsolateResourcesLimitExceeded"

	// 超过密钥对配额。
	LIMITEXCEEDED_KEYPAIRLIMITEXCEEDED = "LimitExceeded.KeyPairLimitExceeded"

	// 镜像当前配额不足，无法共享至指定账号。
	LIMITEXCEEDED_SHAREBLUEPRINTACROSSACCOUNTQUOTALIMITEXCEEDED = "LimitExceeded.ShareBlueprintAcrossAccountQuotaLimitExceeded"

	// 超过快照配额。
	LIMITEXCEEDED_SNAPSHOTQUOTALIMITEXCEEDED = "LimitExceeded.SnapshotQuotaLimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 套餐缺少参数镜像ID。
	MISSINGPARAMETER_BUNDLEMISSINGPARAMETERBLUEPRINTID = "MissingParameter.BundleMissingParameterBlueprintId"

	// 必须传入参数Period或CurInstanceDeadline。
	MISSINGPARAMETER_MISSINGPARAMETERPERIODCURINSTANCEDEADLINE = "MissingParameter.MissingParameterPeriodCurInstanceDeadline"

	// 指定参数`Order`时，必须传入参数 `OrderField`。
	MISSINGPARAMETER_ORDERFIELDREQUIRED = "MissingParameter.OrderFieldRequired"

	// 镜像在操作中。请稍后再试。
	OPERATIONDENIED_BLUEPRINTOPERATIONINPROGRESS = "OperationDenied.BlueprintOperationInProgress"

	// 该实例不支持升级套餐操作。
	OPERATIONDENIED_BUNDLENOTSUPPORTMODIFY = "OperationDenied.BundleNotSupportModify"

	// 磁盘备份点忙，请稍后重新操作。
	OPERATIONDENIED_DISKBACKUPBUSY = "OperationDenied.DiskBackupBusy"

	// 磁盘备份点正在操作过程中，请稍后重试。
	OPERATIONDENIED_DISKBACKUPOPERATIONINPROGRESS = "OperationDenied.DiskBackupOperationInProgress"

	// 磁盘正在操作备份点过程中，请稍后重新操作。
	OPERATIONDENIED_DISKBUSYFORBACKUPOPERATION = "OperationDenied.DiskBusyForBackupOperation"

	// 磁盘处于创建过程中。
	OPERATIONDENIED_DISKCREATING = "OperationDenied.DiskCreating"

	// 磁盘正在操作过程中，请稍后重试。
	OPERATIONDENIED_DISKOPERATIONINPROGRESS = "OperationDenied.DiskOperationInProgress"

	// 磁盘的云盘类型不支持该操作。
	OPERATIONDENIED_DISKUSAGENOTSUPPORTOPERATION = "OperationDenied.DiskUsageNotSupportOperation"

	// 禁止对实例进行操作，实例在创建中，不允许进行该操作。
	OPERATIONDENIED_INSTANCECREATING = "OperationDenied.InstanceCreating"

	// 禁止对实例进行操作，实例最近一次的操作尚在进行中。
	OPERATIONDENIED_INSTANCEOPERATIONINPROGRESS = "OperationDenied.InstanceOperationInProgress"

	// 禁止创建快照。
	OPERATIONDENIED_OPERATIONDENIEDCREATESNAPSHOT = "OperationDenied.OperationDeniedCreateSnapshot"

	// 使用存储型套餐的实例不支持创建快照。
	OPERATIONDENIED_OPERATIONDENIEDCREATESNAPSHOTFORSTORAGEBUNDLE = "OperationDenied.OperationDeniedCreateSnapshotForStorageBundle"

	// 镜像正在修改共享属性操作中。不支持此操作。
	RESOURCEINUSE_BLUEPRINTMODIFYINGSHAREPERMISSION = "ResourceInUse.BlueprintModifyingSharePermission"

	// 磁盘备份点正在使用中，不支持此操作。
	RESOURCEINUSE_DISKBACKUPINUSE = "ResourceInUse.DiskBackupInUse"

	// 密钥对正在使用中。
	RESOURCEINUSE_KEYPAIRINUSE = "ResourceInUse.KeyPairInUse"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 镜像 ID 不存在。
	RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"

	// 指定的镜像不存在。请检查镜像的BlueprintId是否正确。
	RESOURCENOTFOUND_BLUEPRINTNOTFOUND = "ResourceNotFound.BlueprintNotFound"

	// 处于已挂载状态的磁盘关联实例不存在。
	RESOURCENOTFOUND_DISKATTACHEDHASNOINSTANCEID = "ResourceNotFound.DiskAttachedHasNoInstanceId"

	// 磁盘备份点ID不存在。
	RESOURCENOTFOUND_DISKBACKUPIDNOTFOUND = "ResourceNotFound.DiskBackupIdNotFound"

	// 磁盘备份点不存在。
	RESOURCENOTFOUND_DISKBACKUPNOTEXISTS = "ResourceNotFound.DiskBackupNotExists"

	// 用户指定磁盘备份点不存在。
	RESOURCENOTFOUND_DISKBACKUPNOTFOUND = "ResourceNotFound.DiskBackupNotFound"

	// 磁盘 ID 不存在。
	RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"

	// 磁盘不存在。
	RESOURCENOTFOUND_DISKNOTEXISTS = "ResourceNotFound.DiskNotExists"

	// 磁盘不存在。
	RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"

	// 防火墙不存在。
	RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"

	// 防火墙规则不存在。
	RESOURCENOTFOUND_FIREWALLRULESNOTFOUND = "ResourceNotFound.FirewallRulesNotFound"

	// 防火墙模板不存在
	RESOURCENOTFOUND_FIREWALLTEMPLATENOTFOUND = "ResourceNotFound.FirewallTemplateNotFound"

	// 防火墙模板规则不存在
	RESOURCENOTFOUND_FIREWALLTEMPLATERULENOTFOUND = "ResourceNotFound.FirewallTemplateRuleNotFound"

	// 实例不存在挂载的数据盘。
	RESOURCENOTFOUND_INSTANCEDATADISKNOTFOUND = "ResourceNotFound.InstanceDataDiskNotFound"

	// 实例 ID 不存在。
	RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"

	// 实例不存在。
	RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"

	// 密钥对 ID 不存在。
	RESOURCENOTFOUND_KEYIDNOTFOUND = "ResourceNotFound.KeyIdNotFound"

	// 密钥对不存在。
	RESOURCENOTFOUND_KEYPAIRNOTFOUND = "ResourceNotFound.KeyPairNotFound"

	// 自定义镜像不存在。
	RESOURCENOTFOUND_PRIVATEBLUEPRINTNOTFOUND = "ResourceNotFound.PrivateBlueprintNotFound"

	// 服务角色不存在, 请为账号添加这个角色。
	RESOURCENOTFOUND_ROLENOTFOUND = "ResourceNotFound.RoleNotFound"

	// 未查询到场景ID
	RESOURCENOTFOUND_SCENEIDNOTFOUND = "ResourceNotFound.SceneIdNotFound"

	// 快照 ID 不存在。
	RESOURCENOTFOUND_SNAPSHOTIDNOTFOUND = "ResourceNotFound.SnapshotIdNotFound"

	// 快照不存在。
	RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 指定实例镜像不符合要求。
	RESOURCEUNAVAILABLE_BLUEPRINTINVALID = "ResourceUnavailable.BlueprintInvalid"

	// 镜像资源不可用。
	RESOURCEUNAVAILABLE_BLUEPRINTUNAVAILABLE = "ResourceUnavailable.BlueprintUnavailable"

	// 套餐不可用。
	RESOURCEUNAVAILABLE_BUNDLEUNAVAILABLE = "ResourceUnavailable.BundleUnavailable"

	// 不能应用该防火墙模板
	RESOURCEUNAVAILABLE_CANNOTAPPLYEMPTYFIREWALLTEMPLATE = "ResourceUnavailable.CannotApplyEmptyFirewallTemplate"

	// Docker资源不可用。
	RESOURCEUNAVAILABLE_DOCKERUNAVAILABLE = "ResourceUnavailable.DockerUnavailable"

	// 防火墙模板在使用中，不支持该操作。
	RESOURCEUNAVAILABLE_FIREWALLTEMPLATEINUSE = "ResourceUnavailable.FirewallTemplateInUse"

	// 当前套餐不支持通过API购买。
	RESOURCEUNAVAILABLE_INVALIDPURCHASEREQUESTSOURCE = "ResourceUnavailable.InvalidPurchaseRequestSource"

	// TAT agent不可用。
	RESOURCEUNAVAILABLE_TATAGENTUNAVAILABLE = "ResourceUnavailable.TATAgentUnavailable"

	// TAT 服务错误。
	RESOURCEUNAVAILABLE_TATSERVICEERROR = "ResourceUnavailable.TATServiceError"

	// 套餐已售罄。
	RESOURCESSOLDOUT_BUNDLESOLDOUT = "ResourcesSoldOut.BundleSoldOut"

	// 套餐无可用配置。
	RESOURCESSOLDOUT_PURCHASESOURCEHASNOBUNDLECONFIGS = "ResourcesSoldOut.PurchaseSourceHasNoBundleConfigs"

	// 套餐无可用配置。
	RESOURCESSOLDOUT_ZONESHASNOBUNDLECONFIGS = "ResourcesSoldOut.ZonesHasNoBundleConfigs"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 无效 Token。
	UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"

	// MFA 已过期。
	UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"

	// MFA 不存在。
	UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"

	// 无权限。
	UNAUTHORIZEDOPERATION_NOPERMISSION = "UnauthorizedOperation.NoPermission"

	// 无权限进行此操作，请求中token不合法。
	UNAUTHORIZEDOPERATION_TOKENINVALID = "UnauthorizedOperation.TokenInvalid"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 将磁盘备份点回滚到其他盘，不支持该操作。
	UNSUPPORTEDOPERATION_APPLYDISKBACKUPTOANOTHERDISK = "UnsupportedOperation.ApplyDiskBackupToAnotherDisk"

	// 没有实例不支持关联到云联网。
	UNSUPPORTEDOPERATION_ATTACHCCNCONDITIONUNSATISFIED = "UnsupportedOperation.AttachCcnConditionUnsatisfied"

	// 关联云联网失败。请检查云联网状态并稍后再试。
	UNSUPPORTEDOPERATION_ATTACHCCNFAILED = "UnsupportedOperation.AttachCcnFailed"

	// 镜像已经被共享。不支持此操作。
	UNSUPPORTEDOPERATION_BLUEPRINTALREADYSHARED = "UnsupportedOperation.BlueprintAlreadyShared"

	// 镜像当前状态不支持该操作。
	UNSUPPORTEDOPERATION_BLUEPRINTCURSTATEINVALID = "UnsupportedOperation.BlueprintCurStateInvalid"

	// 镜像没有被共享。不支持此操作。
	UNSUPPORTEDOPERATION_BLUEPRINTHASNOTSHARED = "UnsupportedOperation.BlueprintHasNotShared"

	// 镜像最近一次的操作尚未完成。
	UNSUPPORTEDOPERATION_BLUEPRINTLATESTOPERATIONUNFINISHED = "UnsupportedOperation.BlueprintLatestOperationUnfinished"

	// 镜像被使用中，不支持该操作。
	UNSUPPORTEDOPERATION_BLUEPRINTOCCUPIED = "UnsupportedOperation.BlueprintOccupied"

	// 该镜像的镜像类型不支持该操作。
	UNSUPPORTEDOPERATION_BLUEPRINTTYPENOTSUPPORTOPERATION = "UnsupportedOperation.BlueprintTypeNotSupportOperation"

	// 已经关联云联网，不支持再次关联。
	UNSUPPORTEDOPERATION_CCNALREADYATTACHED = "UnsupportedOperation.CcnAlreadyAttached"

	// 云联网 尚未关联。不支持此操作。
	UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"

	// 查询云联网关联的实例状态失败。请稍后再试。
	UNSUPPORTEDOPERATION_DESCRIBECCNATTACHEDINSTANCESFAILED = "UnsupportedOperation.DescribeCcnAttachedInstancesFailed"

	// 解关联云联网失败。请检查云联网状态并稍后再试。
	UNSUPPORTEDOPERATION_DETACHCCNFAILED = "UnsupportedOperation.DetachCcnFailed"

	// 磁盘备份点上一次操作未结束，不支持当前操作。
	UNSUPPORTEDOPERATION_DISKBACKUPLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskBackupLatestOperationUnfinished"

	// 磁盘忙。
	UNSUPPORTEDOPERATION_DISKBUSY = "UnsupportedOperation.DiskBusy"

	// 不支持的操作，磁盘最近一次的操作尚未完成。
	UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"

	// 防火墙忙。
	UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"

	// 指定的防火墙版本号和当前版本不一致。
	UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"

	// 实例展示区域不支持该操作。
	UNSUPPORTEDOPERATION_INSTANCEDISPLAYAREANOTSUPPORTOPERATION = "UnsupportedOperation.InstanceDisplayAreaNotSupportOperation"

	// 实例已到期，不支持该操作。
	UNSUPPORTEDOPERATION_INSTANCEEXPIRED = "UnsupportedOperation.InstanceExpired"

	// LinuxUnix实例在创建时不支持设置密码。
	UNSUPPORTEDOPERATION_INSTANCELINUXUNIXCREATINGNOTSUPPORTPASSWORD = "UnsupportedOperation.InstanceLinuxUnixCreatingNotSupportPassword"

	// 磁盘备份点状态不支持该操作。
	UNSUPPORTEDOPERATION_INVALIDDISKBACKUPSTATE = "UnsupportedOperation.InvalidDiskBackupState"

	// 磁盘状态不支持该操作。
	UNSUPPORTEDOPERATION_INVALIDDISKSTATE = "UnsupportedOperation.InvalidDiskState"

	// 不支持的操作，实例状态不合法。
	UNSUPPORTEDOPERATION_INVALIDINSTANCESTATE = "UnsupportedOperation.InvalidInstanceState"

	// 不支持的操作，快照状态不合法。
	UNSUPPORTEDOPERATION_INVALIDSNAPSHOTSTATE = "UnsupportedOperation.InvalidSnapshotState"

	// 不支持的操作，不支持将同一个密钥对重复绑定到同一个实例上。
	UNSUPPORTEDOPERATION_KEYPAIRBINDDUPLICATE = "UnsupportedOperation.KeyPairBindDuplicate"

	// 不支持该操作。KeyPair 与镜像存在绑定关系。在进行该操作前请删除与密钥对有绑定关系的自定义镜像。
	UNSUPPORTEDOPERATION_KEYPAIRBINDTOBLUEPRINTS = "UnsupportedOperation.KeyPairBindToBlueprints"

	// 不支持的操作，不支持将未绑定到实例的密钥对从实例解绑。
	UNSUPPORTEDOPERATION_KEYPAIRNOTBOUNDTOINSTANCE = "UnsupportedOperation.KeyPairNotBoundToInstance"

	// 不支持的操作，实例最近一次的操作尚未完成。
	UNSUPPORTEDOPERATION_LATESTOPERATIONUNFINISHED = "UnsupportedOperation.LatestOperationUnfinished"

	// 共享镜像不支持此操作。
	UNSUPPORTEDOPERATION_NOTSUPPORTSHAREDBLUEPRINT = "UnsupportedOperation.NotSupportSharedBlueprint"

	// 国际站不支持该操作。
	UNSUPPORTEDOPERATION_OPERATIONNOTSUPPORTEDININTLSITE = "UnsupportedOperation.OperationNotSupportedInIntlSite"

	// 此接口已下线。
	UNSUPPORTEDOPERATION_OPERATIONOFFLINE = "UnsupportedOperation.OperationOffline"

	// 计费资源中心删除资源失败。
	UNSUPPORTEDOPERATION_POSTDESTROYRESOURCEFAILED = "UnsupportedOperation.PostDestroyResourceFailed"

	// 重新申请关联云联网失败。请检查云联网状态并稍后再试。
	UNSUPPORTEDOPERATION_RESETATTACHCCNFAILED = "UnsupportedOperation.ResetAttachCcnFailed"

	// 资源不支持退换。
	UNSUPPORTEDOPERATION_RESOURCENOTRETURNABLE = "UnsupportedOperation.ResourceNotReturnable"

	// 资源变配操作中新旧配置一样，不支持此操作。
	UNSUPPORTEDOPERATION_SAMEWITHOLDCONFIG = "UnsupportedOperation.SameWithOldConfig"

	// 快照忙。
	UNSUPPORTEDOPERATION_SNAPSHOTBUSY = "UnsupportedOperation.SnapshotBusy"

	// 系统忙。
	UNSUPPORTEDOPERATION_SYSTEMBUSY = "UnsupportedOperation.SystemBusy"

	// 实例上腾讯云助手 agent 不在线。
	UNSUPPORTEDOPERATION_TATAGENTNOTONLINE = "UnsupportedOperation.TatAgentNotOnline"

	// Windows实例不支持绑定密钥对。
	UNSUPPORTEDOPERATION_WINDOWSNOTALLOWTOASSOCIATEKEYPAIR = "UnsupportedOperation.WindowsNotAllowToAssociateKeyPair"

	// windows类型实例不支持密钥对功能。
	UNSUPPORTEDOPERATION_WINDOWSNOTSUPPORTKEYPAIR = "UnsupportedOperation.WindowsNotSupportKeyPair"
)
