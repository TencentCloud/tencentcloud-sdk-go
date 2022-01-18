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

	// 创建镜像失败。
	FAILEDOPERATION_CREATEBLUEPRINTFAILED = "FailedOperation.CreateBlueprintFailed"

	// 对密钥对的创建操作失败。
	FAILEDOPERATION_CREATEKEYPAIRFAILED = "FailedOperation.CreateKeyPairFailed"

	// 对密钥对的删除操作失败。
	FAILEDOPERATION_DELETEKEYPAIRFAILED = "FailedOperation.DeleteKeyPairFailed"

	// 对防火墙规则的操作失败。
	FAILEDOPERATION_FIREWALLRULESOPERATIONFAILED = "FailedOperation.FirewallRulesOperationFailed"

	// 对密钥对的导入操作失败。
	FAILEDOPERATION_IMPORTKEYPAIRFAILED = "FailedOperation.ImportKeyPairFailed"

	// 对实例的操作失败。
	FAILEDOPERATION_INSTANCEOPERATIONFAILED = "FailedOperation.InstanceOperationFailed"

	// 快照操作失败。
	FAILEDOPERATION_SNAPSHOTOPERATIONFAILED = "FailedOperation.SnapshotOperationFailed"

	// 操作失败，不能创建自定义镜像。
	FAILEDOPERATION_UNABLETOCREATEBLUEPRINT = "FailedOperation.UnableToCreateBlueprint"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 查询磁盘返回了不合法内容。
	INTERNALERROR_DESCRIBEDISKSRETURNABLEERROR = "InternalError.DescribeDisksReturnableError"

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

	// 快照配额锁获取失败。
	INTERNALERROR_GETSNAPSHOTALLOCQUOTALOCKERROR = "InternalError.GetSnapshotAllocQuotaLockError"

	// 套餐价格错误。
	INTERNALERROR_INVALIDBUNDLEPRICE = "InternalError.InvalidBundlePrice"

	// 命令 `DescribeInstanceLoginKeyPair` 无法找到。
	INTERNALERROR_INVALIDCOMMANDNOTFOUND = "InternalError.InvalidCommandNotFound"

	// 请求出现错误。
	INTERNALERROR_REQUESTERROR = "InternalError.RequestError"

	// 调用计费网关服务失败。
	INTERNALERROR_TRADECALLBILLINGGATEWAYFAILED = "InternalError.TradeCallBillingGatewayFailed"

	// 价格获取失败。
	INTERNALERROR_TRADEGETPRICEFAILED = "InternalError.TradeGetPriceFailed"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 产品未定义的套餐 ID。
	INVALIDPARAMETER_BUNDLEIDNOTFOUND = "InvalidParameter.BundleIdNotFound"

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

	// 参数非法，Filter 参数不是字典。
	INVALIDPARAMETER_INVALIDFILTERNOTDICT = "InvalidParameter.InvalidFilterNotDict"

	// 参数非法，Filter 参数中有不支持的 Name。
	INVALIDPARAMETER_INVALIDFILTERNOTSUPPORTEDNAME = "InvalidParameter.InvalidFilterNotSupportedName"

	// 参数非法，每次只能修改一个属性。
	INVALIDPARAMETER_ONLYALLOWMODIFYONEATTRIBUTE = "InvalidParameter.OnlyAllowModifyOneAttribute"

	// 参数非法，参数冲突。
	INVALIDPARAMETER_PARAMETERCONFLICT = "InvalidParameter.ParameterConflict"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 该实例配置不符合指定镜像的要求。
	INVALIDPARAMETERVALUE_BLUEPRINTCONFIGNOTMATCH = "InvalidParameterValue.BlueprintConfigNotMatch"

	// 镜像 ID 不合法，重装实例不允许切换操作系统类型。
	INVALIDPARAMETERVALUE_BLUEPRINTID = "InvalidParameterValue.BlueprintId"

	// 参数值非法，镜像 ID 格式非法。
	INVALIDPARAMETERVALUE_BLUEPRINTIDMALFORMED = "InvalidParameterValue.BlueprintIdMalformed"

	// 云联网实例ID格式非法。
	INVALIDPARAMETERVALUE_CCNIDMALFORMED = "InvalidParameterValue.CcnIdMalformed"

	// 磁盘大小发生改变。
	INVALIDPARAMETERVALUE_DISKSIZENOTMATCH = "InvalidParameterValue.DiskSizeNotMatch"

	// 参数 `KeyName` 已经存在且重复。
	INVALIDPARAMETERVALUE_DUPLICATEPARAMETERVALUE = "InvalidParameterValue.DuplicateParameterValue"

	// 参数值非法，不允许包含重复的值。
	INVALIDPARAMETERVALUE_DUPLICATED = "InvalidParameterValue.Duplicated"

	// 防火墙规则描述长度超出限制。
	INVALIDPARAMETERVALUE_FIREWALLRULEDESCRIPTIONTOOLONG = "InvalidParameterValue.FirewallRuleDescriptionTooLong"

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

	// 控制台显示类型不合法。
	INVALIDPARAMETERVALUE_INVALIDCONSOLEDISPLAYTYPES = "InvalidParameterValue.InvalidConsoleDisplayTypes"

	// 参数值非法，磁盘 ID 格式非法。
	INVALIDPARAMETERVALUE_INVALIDDISKIDMALFORMED = "InvalidParameterValue.InvalidDiskIdMalformed"

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

	// 参数中的密码不合法。
	INVALIDPARAMETERVALUE_INVALIDPASSWORD = "InvalidParameterValue.InvalidPassword"

	// 不正确的配额资源名称。
	INVALIDPARAMETERVALUE_INVALIDRESOURCEQUOTARESOURCENAME = "InvalidParameterValue.InvalidResourceQuotaResourceName"

	// 参数Zone的取值不合法。
	INVALIDPARAMETERVALUE_INVALIDZONE = "InvalidParameterValue.InvalidZone"

	// 参数值非法，密钥对 ID 格式非法。
	INVALIDPARAMETERVALUE_KEYPAIRIDMALFORMED = "InvalidParameterValue.KeyPairIdMalformed"

	// 该密钥对中的公钥在系统中已存在，无法重复使用 。
	INVALIDPARAMETERVALUE_KEYPAIRPUBLICKEYDUPLICATED = "InvalidParameterValue.KeyPairPublicKeyDuplicated"

	// 指定的公钥格式错误。
	INVALIDPARAMETERVALUE_KEYPAIRPUBLICKEYMALFORMED = "InvalidParameterValue.KeyPairPublicKeyMalformed"

	// 参数值非法，参数值的数量超过最大限制。
	INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"

	// 参数值非法，不能为负值。
	INVALIDPARAMETERVALUE_NEGATIVE = "InvalidParameterValue.Negative"

	// 不允许改变平台类型。
	INVALIDPARAMETERVALUE_NOTALLOWTOCHANGEPLATFORMTYPE = "InvalidParameterValue.NotAllowToChangePlatformType"

	// 参数值非法，不在合法范围内。
	INVALIDPARAMETERVALUE_OUTOFRANGE = "InvalidParameterValue.OutOfRange"

	// 参数值非法，快照 ID 格式非法。
	INVALIDPARAMETERVALUE_SNAPSHOTIDMALFORMED = "InvalidParameterValue.SnapshotIdMalformed"

	// 参数值非法，快照名称超过允许的最大长度。
	INVALIDPARAMETERVALUE_SNAPSHOTNAMETOOLONG = "InvalidParameterValue.SnapshotNameTooLong"

	// 参数取值过长，超过最大长度。
	INVALIDPARAMETERVALUE_TOOLONG = "InvalidParameterValue.TooLong"

	// 非法的可用区。
	INVALIDPARAMETERVALUE_ZONEINVALID = "InvalidParameterValue.ZoneInvalid"

	// 实例挂载数据盘配额不足，无法挂载磁盘。
	LIMITEXCEEDED_ATTACHDATADISKQUOTALIMITEXCEEDED = "LimitExceeded.AttachDataDiskQuotaLimitExceeded"

	// 超过防火墙规则配额。
	LIMITEXCEEDED_FIREWALLRULESLIMITEXCEEDED = "LimitExceeded.FirewallRulesLimitExceeded"

	// 超过密钥对配额。
	LIMITEXCEEDED_KEYPAIRLIMITEXCEEDED = "LimitExceeded.KeyPairLimitExceeded"

	// 超过快照配额。
	LIMITEXCEEDED_SNAPSHOTQUOTALIMITEXCEEDED = "LimitExceeded.SnapshotQuotaLimitExceeded"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 该实例不支持升级套餐操作。
	OPERATIONDENIED_BUNDLENOTSUPPORTMODIFY = "OperationDenied.BundleNotSupportModify"

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

	// 使用存储型套餐的实例不支持创建快照。
	OPERATIONDENIED_OPERATIONDENIEDCREATESNAPSHOTFORSTORAGEBUNDLE = "OperationDenied.OperationDeniedCreateSnapshotForStorageBundle"

	// 密钥对正在使用中。
	RESOURCEINUSE_KEYPAIRINUSE = "ResourceInUse.KeyPairInUse"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 镜像 ID 不存在。
	RESOURCENOTFOUND_BLUEPRINTIDNOTFOUND = "ResourceNotFound.BlueprintIdNotFound"

	// 指定的镜像不存在。请检查镜像的BlueprintId是否正确。
	RESOURCENOTFOUND_BLUEPRINTNOTFOUND = "ResourceNotFound.BlueprintNotFound"

	// 磁盘 ID 不存在。
	RESOURCENOTFOUND_DISKIDNOTFOUND = "ResourceNotFound.DiskIdNotFound"

	// 磁盘不存在。
	RESOURCENOTFOUND_DISKNOTFOUND = "ResourceNotFound.DiskNotFound"

	// 防火墙不存在。
	RESOURCENOTFOUND_FIREWALLNOTFOUND = "ResourceNotFound.FirewallNotFound"

	// 防火墙规则不存在。
	RESOURCENOTFOUND_FIREWALLRULESNOTFOUND = "ResourceNotFound.FirewallRulesNotFound"

	// 实例 ID 不存在。
	RESOURCENOTFOUND_INSTANCEIDNOTFOUND = "ResourceNotFound.InstanceIdNotFound"

	// 实例不存在。
	RESOURCENOTFOUND_INSTANCENOTFOUND = "ResourceNotFound.InstanceNotFound"

	// 密钥对 ID 不存在。
	RESOURCENOTFOUND_KEYIDNOTFOUND = "ResourceNotFound.KeyIdNotFound"

	// 快照 ID 不存在。
	RESOURCENOTFOUND_SNAPSHOTIDNOTFOUND = "ResourceNotFound.SnapshotIdNotFound"

	// 快照不存在。
	RESOURCENOTFOUND_SNAPSHOTNOTFOUND = "ResourceNotFound.SnapshotNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 镜像资源不可用。
	RESOURCEUNAVAILABLE_BLUEPRINTUNAVAILABLE = "ResourceUnavailable.BlueprintUnavailable"

	// MFA 已过期。
	UNAUTHORIZEDOPERATION_MFAEXPIRED = "UnauthorizedOperation.MFAExpired"

	// MFA 不存在。
	UNAUTHORIZEDOPERATION_MFANOTFOUND = "UnauthorizedOperation.MFANotFound"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 没有实例不支持关联到云联网。
	UNSUPPORTEDOPERATION_ATTACHCCNCONDITIONUNSATISFIED = "UnsupportedOperation.AttachCcnConditionUnsatisfied"

	// 关联云联网失败。请检查云联网状态并稍后再试。
	UNSUPPORTEDOPERATION_ATTACHCCNFAILED = "UnsupportedOperation.AttachCcnFailed"

	// 镜像当前状态不支持该操作。
	UNSUPPORTEDOPERATION_BLUEPRINTCURSTATEINVALID = "UnsupportedOperation.BlueprintCurStateInvalid"

	// 镜像被使用中，不支持该操作。
	UNSUPPORTEDOPERATION_BLUEPRINTOCCUPIED = "UnsupportedOperation.BlueprintOccupied"

	// 已经关联云联网，不支持再次关联。
	UNSUPPORTEDOPERATION_CCNALREADYATTACHED = "UnsupportedOperation.CcnAlreadyAttached"

	// 云联网 尚未关联。不支持此操作。
	UNSUPPORTEDOPERATION_CCNNOTATTACHED = "UnsupportedOperation.CcnNotAttached"

	// 查询云联网关联的实例状态失败。请稍后再试。
	UNSUPPORTEDOPERATION_DESCRIBECCNATTACHEDINSTANCESFAILED = "UnsupportedOperation.DescribeCcnAttachedInstancesFailed"

	// 解关联云联网失败。请检查云联网状态并稍后再试。
	UNSUPPORTEDOPERATION_DETACHCCNFAILED = "UnsupportedOperation.DetachCcnFailed"

	// 磁盘忙。
	UNSUPPORTEDOPERATION_DISKBUSY = "UnsupportedOperation.DiskBusy"

	// 不支持的操作，磁盘最近一次的操作尚未完成。
	UNSUPPORTEDOPERATION_DISKLATESTOPERATIONUNFINISHED = "UnsupportedOperation.DiskLatestOperationUnfinished"

	// 防火墙忙。
	UNSUPPORTEDOPERATION_FIREWALLBUSY = "UnsupportedOperation.FirewallBusy"

	// 指定的防火墙版本号和当前版本不一致。
	UNSUPPORTEDOPERATION_FIREWALLVERSIONMISMATCH = "UnsupportedOperation.FirewallVersionMismatch"

	// 实例已到期，不支持该操作。
	UNSUPPORTEDOPERATION_INSTANCEEXPIRED = "UnsupportedOperation.InstanceExpired"

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

	// 计费资源中心删除资源失败。
	UNSUPPORTEDOPERATION_POSTDESTROYRESOURCEFAILED = "UnsupportedOperation.PostDestroyResourceFailed"

	// 重新申请关联云联网失败。请检查云联网状态并稍后再试。
	UNSUPPORTEDOPERATION_RESETATTACHCCNFAILED = "UnsupportedOperation.ResetAttachCcnFailed"

	// 快照忙。
	UNSUPPORTEDOPERATION_SNAPSHOTBUSY = "UnsupportedOperation.SnapshotBusy"

	// Windows实例不支持绑定密钥对。
	UNSUPPORTEDOPERATION_WINDOWSNOTALLOWTOASSOCIATEKEYPAIR = "UnsupportedOperation.WindowsNotAllowToAssociateKeyPair"
)
