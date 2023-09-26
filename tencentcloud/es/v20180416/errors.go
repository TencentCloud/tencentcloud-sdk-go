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

package v20180416

const (
	// 此产品的特有错误码

	// 当前用户对查询集群实例操作未授权，请添加CAM权限。
	AUTHFAILURE_UNAUTHDESCRIBEINSTANCES = "AuthFailure.UnAuthDescribeInstances"

	// 操作未授权。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"

	// 集群资源配额限制错误。
	FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"

	// 节点磁盘块数参数检查失败。
	FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"

	// 集群状态错误。
	FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"

	// 由于EsDictionaryInfo错误，不允许操纵。
	FAILEDOPERATION_ESDICTIONARYINFOERROR = "FailedOperation.EsDictionaryInfoError"

	// 文件名不符合规范。
	FAILEDOPERATION_FILENAMEERROR = "FailedOperation.FileNameError"

	// 文件大小不符合要求。
	FAILEDOPERATION_FILESIZEERROR = "FailedOperation.FileSizeError"

	// 获取实例的标签列表错误。
	FAILEDOPERATION_GETTAGINFOERROR = "FailedOperation.GetTagInfoError"

	// 账户未绑定信用卡或paypal，无法支付。
	FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"

	// 用户未实名认证。
	FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"

	// 该实例不满足退费条件。
	FAILEDOPERATION_REFUNDERROR = "FailedOperation.RefundError"

	// 不支持对本地盘集群通过滚动重启方式扩缩容。
	FAILEDOPERATION_UNSUPPORTEDLOCALDISKROLLUPSCALEUPORDOWN = "FailedOperation.UnsupportedLocalDiskRollUpScaleUpOrDown"

	// 不支持在滚动重启扩容计算资源同时扩容磁盘数量。
	FAILEDOPERATION_UNSUPPORTEDRESETNODETYPEANDSCALEOUTDISK = "FailedOperation.UnsupportedResetNodeTypeAndScaleOutDisk"

	// 不支持在滚动重启缩容计算资源同时修改磁盘大小
	FAILEDOPERATION_UNSUPPORTEDRESTSCALEDOWNANDMODIFYDISK = "FailedOperation.UnsupportedRestScaleDownAndModifyDisk"

	// 不支持反向调节节点配置和磁盘容量。
	FAILEDOPERATION_UNSUPPORTEDREVERSEREGULATIONNODETYPEANDDISK = "FailedOperation.UnsupportedReverseRegulationNodeTypeAndDisk"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// Gc类型的取值和预期不符。
	INVALIDPARAMETER_GC = "InvalidParameter.GC"

	// AppId的取值和预期不符。
	INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"

	// 是否自动使用代金券AutoVoucher的取值和预期不符。
	INVALIDPARAMETER_INVALIDAUTOVOUCHER = "InvalidParameter.InvalidAutoVoucher"

	// ClusterName的取值和预期不符。
	INVALIDPARAMETER_INVALIDCLUSTERNAME = "InvalidParameter.InvalidClusterName"

	// cos自动备份信息与预期不符。
	INVALIDPARAMETER_INVALIDCOSBACKUPINFO = "InvalidParameter.InvalidCosBackupInfo"

	// 集群部署方式InvalidDeployMode的取值和预期不符。
	INVALIDPARAMETER_INVALIDDEPLOYMODE = "InvalidParameter.InvalidDeployMode"

	// 节点磁盘块数DiskCount取值和预期不符。
	INVALIDPARAMETER_INVALIDDISKCOUNT = "InvalidParameter.InvalidDiskCount"

	// 节点磁盘加密信息DiskEncrypt的取值与预期不符。
	INVALIDPARAMETER_INVALIDDISKENCRYPT = "InvalidParameter.InvalidDiskEncrypt"

	// 是否启用增强型ssd云盘DiskEnhance的取值和预期不符。
	INVALIDPARAMETER_INVALIDDISKENHANCE = "InvalidParameter.InvalidDiskEnhance"

	// 节点磁盘容量DiskSize的取值和预期不符。
	INVALIDPARAMETER_INVALIDDISKSIZE = "InvalidParameter.InvalidDiskSize"

	// 节点磁盘类型DiskType的取值和预期不符。
	INVALIDPARAMETER_INVALIDDISKTYPE = "InvalidParameter.InvalidDiskType"

	// EsAcl的取值和预期不符。
	INVALIDPARAMETER_INVALIDESACL = "InvalidParameter.InvalidEsACL"

	// 实例版本EsVersion的取值和预期不符。
	INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"

	// 创建的索引元数据JSON IndexMetaJson的取值和预期不符。
	INVALIDPARAMETER_INVALIDINDEXMETAJSON = "InvalidParameter.InvalidIndexMetaJson"

	// 索引名IndexName的取值和预期不符。
	INVALIDPARAMETER_INVALIDINDEXNAME = "InvalidParameter.InvalidIndexName"

	// 索引类型IndexType的取值和预期不符。
	INVALIDPARAMETER_INVALIDINDEXTYPE = "InvalidParameter.InvalidIndexType"

	// 无效的InstanceId，没有找到对应资源。
	INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"

	// Ip地址取值和预期不符。
	INVALIDPARAMETER_INVALIDIP = "InvalidParameter.InvalidIp"

	// 私有网络vip列表IpList的取值和预期不符。
	INVALIDPARAMETER_INVALIDIPLIST = "InvalidParameter.InvalidIpList"

	// Kibana内网访问端口的取值和预期不符。
	INVALIDPARAMETER_INVALIDKIBANAPRIVATEPORT = "InvalidParameter.InvalidKibanaPrivatePort"

	// 分页大小Limit的取值和预期不符。
	INVALIDPARAMETER_INVALIDLIMIT = "InvalidParameter.InvalidLimit"

	// 日志类型LogType的取值和预期不符。
	INVALIDPARAMETER_INVALIDLOGTYPE = "InvalidParameter.InvalidLogType"

	// 多可用区部署ZoneDetail的信息与预期不符。
	INVALIDPARAMETER_INVALIDMULTIZONEINFO = "InvalidParameter.InvalidMultiZoneInfo"

	// 节点名称列表NodeNames的取值和预期不符。
	INVALIDPARAMETER_INVALIDNODENAMES = "InvalidParameter.InvalidNodeNames"

	// 节点数量NodeNum的取值和预期不符。
	INVALIDPARAMETER_INVALIDNODENUM = "InvalidParameter.InvalidNodeNum"

	// 节点规格NodeType的取值和预期不符。
	INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"

	// 分页起始值Offset的取值和预期不符。
	INVALIDPARAMETER_INVALIDOFFSET = "InvalidParameter.InvalidOffset"

	// 操作类型OpType的取值和预期不符。
	INVALIDPARAMETER_INVALIDOPTYPE = "InvalidParameter.InvalidOpType"

	// 可维护时间段OperationDuration的取值和预期不符。
	INVALIDPARAMETER_INVALIDOPERATIONDURATION = "InvalidParameter.InvalidOperationDuration"

	// 排序字段OrderByKey的取值和预期不符。
	INVALIDPARAMETER_INVALIDORDERBYKEY = "InvalidParameter.InvalidOrderByKey"

	// 排序方式OrderByType的取值和预期不符。
	INVALIDPARAMETER_INVALIDORDERBYTYPE = "InvalidParameter.InvalidOrderByType"

	// 内网访问PrivateAccess的取值和预期不符。
	INVALIDPARAMETER_INVALIDPRIVATEACCESS = "InvalidParameter.InvalidPrivateAccess"

	// 外网访问PublicAccess的取值和预期不符。
	INVALIDPARAMETER_INVALIDPUBLICACCESS = "InvalidParameter.InvalidPublicAccess"

	// 地域Region的取值与预期不符。
	INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"

	// 可选重启模式RestartMode的取值和预期不符。
	INVALIDPARAMETER_INVALIDRESTARTMODE = "InvalidParameter.InvalidRestartMode"

	// 重启方式RestartType的取值和预期不符。
	INVALIDPARAMETER_INVALIDRESTARTTYPE = "InvalidParameter.InvalidRestartType"

	// 安全组id列表SecurityGroupIds的取值和预期不符。
	INVALIDPARAMETER_INVALIDSECURITYGROUPIDS = "InvalidParameter.InvalidSecurityGroupIds"

	// 子网ID SubnetId的取值与预期不符。
	INVALIDPARAMETER_INVALIDSUBNETID = "InvalidParameter.InvalidSubnetId"

	// 虚拟子网络统一ID列表的取值和预期不符。
	INVALIDPARAMETER_INVALIDSUBNETUIDLIST = "InvalidParameter.InvalidSubnetUidList"

	// 节点标签信息TagInfo的取值和预期不符。
	INVALIDPARAMETER_INVALIDTAGINFO = "InvalidParameter.InvalidTagInfo"

	// 节点标签信息列表TagList的取值和预期不符。
	INVALIDPARAMETER_INVALIDTAGLIST = "InvalidParameter.InvalidTagList"

	// 时间相关参数的取值或格式和预期不符。
	INVALIDPARAMETER_INVALIDTIMEPARAM = "InvalidParameter.InvalidTimeParam"

	// 节点类型Type的取值和预期不符。
	INVALIDPARAMETER_INVALIDTYPE = "InvalidParameter.InvalidType"

	// 更新的索引元数据json UpdateMetaJson的取值和预期不符。
	INVALIDPARAMETER_INVALIDUPDATEMETAJSON = "InvalidParameter.InvalidUpdateMetaJson"

	// 更新类型UpdateType的取值和预期不符。
	INVALIDPARAMETER_INVALIDUPDATETYPE = "InvalidParameter.InvalidUpdateType"

	// 代金券ID列表VoucherIds的取值和预期不符。
	INVALIDPARAMETER_INVALIDVOUCHERIDS = "InvalidParameter.InvalidVoucherIds"

	// 可用区Zone的信息与预期不符。
	INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 计费类型ChargeType的参数取值有误。
	INVALIDPARAMETERVALUE_CHARGETYPE = "InvalidParameterValue.ChargeType"

	// 组件配置值有误。
	INVALIDPARAMETERVALUE_CONFIGINFO = "InvalidParameterValue.ConfigInfo"

	// ES配置类型取值有误。
	INVALIDPARAMETERVALUE_ESCONFIGTYPE = "InvalidParameterValue.EsConfigType"

	// 为集群稳定起见，当前只支持同时安装1个用户插件。
	INVALIDPARAMETERVALUE_INSTALLPLUGINLIST = "InvalidParameterValue.InstallPluginList"

	// 实例名校验失败，按规则输入。
	INVALIDPARAMETERVALUE_INSTANCENAME = "InvalidParameterValue.InstanceName"

	// 截止日期冲突。
	INVALIDPARAMETERVALUE_INVALIDDEADLINE = "InvalidParameterValue.InvalidDeadline"

	// Jdk类型的取值和预期不符。
	INVALIDPARAMETERVALUE_INVALIDJDK = "InvalidParameterValue.InvalidJDK"

	// 密码校验不合法。
	INVALIDPARAMETERVALUE_PASSWORD = "InvalidParameterValue.Password"

	// 插件类型PluginType的取值有误。
	INVALIDPARAMETERVALUE_PLUGINTYPE = "InvalidParameterValue.PluginType"

	// 自动续费标识RenewFlag的参数取值有误。
	INVALIDPARAMETERVALUE_RENEWFLAG = "InvalidParameterValue.RenewFlag"

	// 升级方式UpgradeMode的取值有误。
	INVALIDPARAMETERVALUE_UPGRADEMODE = "InvalidParameterValue.UpgradeMode"

	// 该账号下的集群数超过限额。
	LIMITEXCEEDED_CLUSTERNUM = "LimitExceeded.ClusterNum"

	// 超过当天智能诊断最大次数。
	LIMITEXCEEDED_DIAGNOSECOUNT = "LimitExceeded.DiagnoseCount"

	// 节点数量或索引存储量超过限额。
	LIMITEXCEEDED_NODENUMORINDICES = "LimitExceeded.NodeNumOrIndices"

	// 超过最大自定义插件数量，请先删除部分自定义插件。
	LIMITEXCEEDED_PLUGININSTALL = "LimitExceeded.PluginInstall"

	// 超出一个用户可以拥有的最大资源限制。
	LIMITEXCEEDED_RESOURCELIMIT = "LimitExceeded.ResourceLimit"

	// 更新参数值过多，超过限制。
	LIMITEXCEEDED_UPDATEITEMLIMIT = "LimitExceeded.UpdateItemLimit"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 存在诊断中的作业，请等待作业诊断完成后重试。
	RESOURCEINUSE_DIAGNOSE = "ResourceInUse.Diagnose"

	// 订单被锁定。
	RESOURCEINUSE_ORDER = "ResourceInUse.Order"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 账户余额不足。
	RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"

	// 计算资源不足，可更换可用区或调整机型。
	RESOURCEINSUFFICIENT_CVM = "ResourceInsufficient.CVM"

	// 隐藏可用区专用主节点资源不足。
	RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"

	// 子网剩余ip数量不足。
	RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"

	// 可用的子网IP不足。
	RESOURCEINSUFFICIENT_SUBNETIP = "ResourceInsufficient.SubnetIp"

	// 可用区资源不足。
	RESOURCEINSUFFICIENT_ZONE = "ResourceInsufficient.Zone"

	// CAM资源未找到。
	RESOURCENOTFOUND_CAMINFONOTFOUND = "ResourceNotFound.CAMInfoNotFound"

	// 集群资源获取失败。
	RESOURCENOTFOUND_CLUSTERINFONOTFOUND = "ResourceNotFound.ClusterInfoNotFound"

	// COS相关信息获取失败。
	RESOURCENOTFOUND_COSINFONOTFOUND = "ResourceNotFound.CosInfoNotFound"

	// 数据库资源获取失败。
	RESOURCENOTFOUND_DBINFONOTFOUND = "ResourceNotFound.DBInfoNotFound"

	// 数据节点信息未找到。
	RESOURCENOTFOUND_DATANODENOTFOUND = "ResourceNotFound.DataNodeNotFound"

	// 智能诊断相关资源未找到。
	RESOURCENOTFOUND_DIAGNOSENOTFOUND = "ResourceNotFound.DiagnoseNotFound"

	// 磁盘相关资源获取失败。
	RESOURCENOTFOUND_DISKINFONOTFOUND = "ResourceNotFound.DiskInfoNotFound"

	// 获取oss资源失败。
	RESOURCENOTFOUND_OSSINFONOTFOUND = "ResourceNotFound.OssInfoNotFound"

	// 安全组信息获取失败。
	RESOURCENOTFOUND_SECURITYGROUPNOTFOUND = "ResourceNotFound.SecurityGroupNotFound"

	// 获取计费资源失败。
	RESOURCENOTFOUND_TRADECGWNOTFOUND = "ResourceNotFound.TradeCgwNotFound"

	// VPC资源获取失败。
	RESOURCENOTFOUND_VPCINFONOTFOUND = "ResourceNotFound.VPCInfoNotFound"

	// 白名单资源获取失败。
	RESOURCENOTFOUND_WHITELISTNOTFOUND = "ResourceNotFound.WhiteListNotFound"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// Uin不在白名单中。
	UNAUTHORIZEDOPERATION_UINNOTINWHITELIST = "UnauthorizedOperation.UinNotInWhiteList"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 不支持该认证方式。
	UNSUPPORTEDOPERATION_BASICSECURITYTYPE = "UnsupportedOperation.BasicSecurityType"

	// 不支持变配节点或磁盘的配置。
	UNSUPPORTEDOPERATION_CHANGENODETYPE = "UnsupportedOperation.ChangeNodeType"

	// 集群中存在部分索引处于关闭状态。
	UNSUPPORTEDOPERATION_CLUSTERSTATECLOSE = "UnsupportedOperation.ClusterStateClose"

	// 集群中存在部分索引无备份。
	UNSUPPORTEDOPERATION_CLUSTERSTATENOREPLICATION = "UnsupportedOperation.ClusterStateNoReplication"

	// 集群的状态不正常（不为绿）。
	UNSUPPORTEDOPERATION_CLUSTERSTATEUNHEALTH = "UnsupportedOperation.ClusterStateUnHealth"

	// 该版本不支持cos备份。集群需要重启来允许cos备份。
	UNSUPPORTEDOPERATION_COSBACKUP = "UnsupportedOperation.CosBackUp"

	// 不支持该智能诊断任务。
	UNSUPPORTEDOPERATION_DIAGNOSEJOB = "UnsupportedOperation.DiagnoseJob"

	// 智能诊断未开启。
	UNSUPPORTEDOPERATION_DIAGNOSENOTOPEN = "UnsupportedOperation.DiagnoseNotOpen"

	// 不支持该操作，服务类型有误。
	UNSUPPORTEDOPERATION_INSTANCETYPEERROR = "UnsupportedOperation.InstanceTypeError"

	// 不支持该操作，license的类型有误。
	UNSUPPORTEDOPERATION_LICENSEERROR = "UnsupportedOperation.LicenseError"

	// 不支持本地盘机型。
	UNSUPPORTEDOPERATION_LOCALDISK = "UnsupportedOperation.LocalDisk"

	// 不支持多可用区变配。
	UNSUPPORTEDOPERATION_MULTIZONESUPGRADE = "UnsupportedOperation.MultiZonesUpgrade"

	// 不支持该插件。
	UNSUPPORTEDOPERATION_PLUGIN = "UnsupportedOperation.Plugin"

	// 离线节点蓝绿变更存在风险，请直接选择重启“in-place”。
	UNSUPPORTEDOPERATION_RESTARTMODE = "UnsupportedOperation.RestartMode"

	// 不支持该操作，当前实例运行状态不为正常。
	UNSUPPORTEDOPERATION_STATUSNOTNORMAL = "UnsupportedOperation.StatusNotNormal"

	// 不支持该操作，实例状态有误。
	UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"

	// VPC资源获取失败。
	UNSUPPORTEDOPERATION_VPCINFONOTFOUND = "UnsupportedOperation.VPCInfoNotFound"

	// 不支持的web组件。
	UNSUPPORTEDOPERATION_WEBSERVICETYPE = "UnsupportedOperation.WebServiceType"
)
