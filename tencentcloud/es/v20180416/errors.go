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

package v20180416

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 当前用户对创建集群实例操作未授权，请添加CAM权限
	AUTHFAILURE_UNAUTHCREATEINSTANCE = "AuthFailure.UnAuthCreateInstance"

	// 当前用户对查询集群实例操作未授权，请添加CAM权限。
	AUTHFAILURE_UNAUTHDESCRIBEINSTANCES = "AuthFailure.UnAuthDescribeInstances"

	// 操作未授权。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnAuthorizedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 更新ACL失败
	FAILEDOPERATION_ACCESSCONTROLLIST = "FailedOperation.AccessControlList"

	// 关闭数据层失败，至少要保留热层和温层其中一层。
	FAILEDOPERATION_CLOSEDATATIER = "FailedOperation.CloseDataTier"

	// 集群资源配额限制错误。
	FAILEDOPERATION_CLUSTERRESOURCELIMITERROR = "FailedOperation.ClusterResourceLimitError"

	// cos融合迁移失败
	FAILEDOPERATION_COSMIGRATE = "FailedOperation.CosMigrate"

	// 节点磁盘块数参数检查失败。
	FAILEDOPERATION_DISKCOUNTPARAMERROR = "FailedOperation.DiskCountParamError"

	// 集群状态错误。
	FAILEDOPERATION_ERRORCLUSTERSTATE = "FailedOperation.ErrorClusterState"

	// 集群状态不健康。
	FAILEDOPERATION_ERRORCLUSTERSTATEUNHEALTH = "FailedOperation.ErrorClusterStateUnhealth"

	// 由于EsDictionaryInfo错误，不允许操纵。
	FAILEDOPERATION_ESDICTIONARYINFOERROR = "FailedOperation.EsDictionaryInfoError"

	// 文件名不符合规范。
	FAILEDOPERATION_FILENAMEERROR = "FailedOperation.FileNameError"

	// 文件大小不符合要求。
	FAILEDOPERATION_FILESIZEERROR = "FailedOperation.FileSizeError"

	// 获取实例的标签列表错误。
	FAILEDOPERATION_GETTAGINFOERROR = "FailedOperation.GetTagInfoError"

	// 获取user失败
	FAILEDOPERATION_GETUSER = "FailedOperation.GetUser"

	// 获取快照列表失败
	FAILEDOPERATION_GETUSERCOSSNAPSHOTLIST = "FailedOperation.GetUserCosSnapshotList"

	// 账户未绑定信用卡或paypal，无法支付。
	FAILEDOPERATION_NOPAYMENT = "FailedOperation.NoPayment"

	// 节点数和可用区的信息不符合对应操作的规则。
	FAILEDOPERATION_NODENUMANDZONEERROR = "FailedOperation.NodeNumAndZoneError"

	// 用户未实名认证。
	FAILEDOPERATION_NOTAUTHENTICATED = "FailedOperation.NotAuthenticated"

	// 当前集群已开启销毁保护，无法直接销毁，请解除该保护后再尝试销毁
	FAILEDOPERATION_OPENDESTROYPROTECTION = "FailedOperation.OpenDestroyProtection"

	// 解码Password时出错。
	FAILEDOPERATION_PASSWORD = "FailedOperation.Password"

	// 该实例不满足退费条件。
	FAILEDOPERATION_REFUNDERROR = "FailedOperation.RefundError"

	// 请求超时。
	FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeOut"

	// 变更space的kibana语言信息失败
	FAILEDOPERATION_SPACELANGUAGE = "FailedOperation.SpaceLanguage"

	// 生成询价或者下单签名有误。
	FAILEDOPERATION_TRADESIGNERROR = "FailedOperation.TradeSignError"

	// 不支持对本地盘集群通过滚动重启方式扩缩容。
	FAILEDOPERATION_UNSUPPORTEDLOCALDISKROLLUPSCALEUPORDOWN = "FailedOperation.UnsupportedLocalDiskRollUpScaleUpOrDown"

	// 不支持开启低频存储
	FAILEDOPERATION_UNSUPPORTEDOPENLOWFREQUENCY = "FailedOperation.UnsupportedOpenLowFrequency"

	// 不支持在滚动重启扩容计算资源同时扩容磁盘数量。
	FAILEDOPERATION_UNSUPPORTEDRESETNODETYPEANDSCALEOUTDISK = "FailedOperation.UnsupportedResetNodeTypeAndScaleOutDisk"

	// 不支持在滚动重启缩容计算资源同时修改磁盘大小
	FAILEDOPERATION_UNSUPPORTEDRESTSCALEDOWNANDMODIFYDISK = "FailedOperation.UnsupportedRestScaleDownAndModifyDisk"

	// 不支持反向调节节点配置和磁盘容量。
	FAILEDOPERATION_UNSUPPORTEDREVERSEREGULATIONNODETYPEANDDISK = "FailedOperation.UnsupportedReverseRegulationNodeTypeAndDisk"

	// 变更space的vpc信息失败
	FAILEDOPERATION_UPDATEVPC = "FailedOperation.UpdateVpc"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// Gc类型的取值和预期不符。
	INVALIDPARAMETER_GC = "InvalidParameter.GC"

	// 索引名已存在。
	INVALIDPARAMETER_INDEXNAMEEXIST = "InvalidParameter.IndexNameExist"

	// 索引名不合法。
	INVALIDPARAMETER_INDEXNAMEINVALID = "InvalidParameter.IndexNameInvalid"

	// AppId的取值和预期不符。
	INVALIDPARAMETER_INVALIDAPPID = "InvalidParameter.InvalidAppId"

	// 是否自动使用代金券AutoVoucher的取值和预期不符。
	INVALIDPARAMETER_INVALIDAUTOVOUCHER = "InvalidParameter.InvalidAutoVoucher"

	// Backends的取值和预期不符。
	INVALIDPARAMETER_INVALIDBACKENDS = "InvalidParameter.InvalidBackends"

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

	// 内网域名解析服务器IP Dns的取值和预期不符。
	INVALIDPARAMETER_INVALIDDNS = "InvalidParameter.InvalidDns"

	// EsAcl的取值和预期不符。
	INVALIDPARAMETER_INVALIDESACL = "InvalidParameter.InvalidEsACL"

	// 实例版本EsVersion的取值和预期不符。
	INVALIDPARAMETER_INVALIDESVERSION = "InvalidParameter.InvalidEsVersion"

	// 创建的索引元数据JSON IndexMetaJson的取值和预期不符。
	INVALIDPARAMETER_INVALIDINDEXMETAJSON = "InvalidParameter.InvalidIndexMetaJson"

	// 索引名IndexName的取值和预期不符。
	INVALIDPARAMETER_INVALIDINDEXNAME = "InvalidParameter.InvalidIndexName"

	// 索引自治字段IndexOptionsField的取值和预期不符。
	INVALIDPARAMETER_INVALIDINDEXOPTIONSFIELD = "InvalidParameter.InvalidIndexOptionsField"

	// 索引生命周期字段IndexPolicyField的取值和预期不符。
	INVALIDPARAMETER_INVALIDINDEXPOLICYFIELD = "InvalidParameter.InvalidIndexPolicyField"

	// 索引配置字段IndexSettingsField的取值和预期不符。
	INVALIDPARAMETER_INVALIDINDEXSETTINGSFIELD = "InvalidParameter.InvalidIndexSettingsField"

	// 索引类型IndexType的取值和预期不符。
	INVALIDPARAMETER_INVALIDINDEXTYPE = "InvalidParameter.InvalidIndexType"

	// 无效的InstanceId，没有找到对应资源。
	INVALIDPARAMETER_INVALIDINSTANCEID = "InvalidParameter.InvalidInstanceId"

	// 集群可维护时间段和集群ID InstanceOperationDurations的取值和预期不符。
	INVALIDPARAMETER_INVALIDINSTANCEOPERATIONDURATIONS = "InvalidParameter.InvalidInstanceOperationDurations"

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

	// 实例版本LogstashVersion的取值和预期不符。
	INVALIDPARAMETER_INVALIDLOGSTASHVERSION = "InvalidParameter.InvalidLogstashVersion"

	// 多可用区部署ZoneDetail的信息与预期不符。
	INVALIDPARAMETER_INVALIDMULTIZONEINFO = "InvalidParameter.InvalidMultiZoneInfo"

	// 节点名称列表NodeNames的取值和预期不符。
	INVALIDPARAMETER_INVALIDNODENAMES = "InvalidParameter.InvalidNodeNames"

	// 节点数量NodeNum的取值和预期不符。
	INVALIDPARAMETER_INVALIDNODENUM = "InvalidParameter.InvalidNodeNum"

	// 节点规格NodeType的取值和预期不符。
	INVALIDPARAMETER_INVALIDNODETYPE = "InvalidParameter.InvalidNodeType"

	// 温热层数据结构NodeTypeInfos的取值和预期不符。
	INVALIDPARAMETER_INVALIDNODETYPESTORAGEINFO = "InvalidParameter.InvalidNodeTypeStorageInfo"

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

	// 需要签名的参数Para的取值和预期不符。
	INVALIDPARAMETER_INVALIDPARA = "InvalidParameter.InvalidPara"

	// 内网访问PrivateAccess的取值和预期不符。
	INVALIDPARAMETER_INVALIDPRIVATEACCESS = "InvalidParameter.InvalidPrivateAccess"

	// 外网访问PublicAccess的取值和预期不符。
	INVALIDPARAMETER_INVALIDPUBLICACCESS = "InvalidParameter.InvalidPublicAccess"

	// 无效的lucene语法
	INVALIDPARAMETER_INVALIDQUERYSTRING = "InvalidParameter.InvalidQueryString"

	// 地域Region的取值与预期不符。
	INVALIDPARAMETER_INVALIDREGION = "InvalidParameter.InvalidRegion"

	// 可选重启模式RestartMode的取值和预期不符。
	INVALIDPARAMETER_INVALIDRESTARTMODE = "InvalidParameter.InvalidRestartMode"

	// 重启方式RestartType的取值和预期不符。
	INVALIDPARAMETER_INVALIDRESTARTTYPE = "InvalidParameter.InvalidRestartType"

	// 数据样例JSON SampleJson的取值和预期不符。
	INVALIDPARAMETER_INVALIDSAMPLEJSON = "InvalidParameter.InvalidSampleJson"

	// 安全组id列表SecurityGroupIds的取值和预期不符。
	INVALIDPARAMETER_INVALIDSECURITYGROUPIDS = "InvalidParameter.InvalidSecurityGroupIds"

	// ES serverless的serverlessId未找到
	INVALIDPARAMETER_INVALIDSERVERLESSID = "InvalidParameter.InvalidServerlessId"

	// ES Serverless的spaceId未找到
	INVALIDPARAMETER_INVALIDSPACEID = "InvalidParameter.InvalidSpaceId"

	// space的名称错误
	INVALIDPARAMETER_INVALIDSPACENAMES = "InvalidParameter.InvalidSpaceNames"

	// 子网ID SubnetId的取值与预期不符。
	INVALIDPARAMETER_INVALIDSUBNETID = "InvalidParameter.InvalidSubnetId"

	// 虚拟子网络统一ID列表的取值和预期不符。
	INVALIDPARAMETER_INVALIDSUBNETUIDLIST = "InvalidParameter.InvalidSubnetUidList"

	// 节点标签信息TagInfo的取值和预期不符。
	INVALIDPARAMETER_INVALIDTAGINFO = "InvalidParameter.InvalidTagInfo"

	// 节点标签信息列表TagList的取值和预期不符。
	INVALIDPARAMETER_INVALIDTAGLIST = "InvalidParameter.InvalidTagList"

	// 目标索引名TargetIndexName的取值和预期不符。
	INVALIDPARAMETER_INVALIDTARGETINDEXNAME = "InvalidParameter.InvalidTargetIndexName"

	// 接收请求的目标节点类型列表InvalidTargetNodeTypes的取值和预期不符。
	INVALIDPARAMETER_INVALIDTARGETNODETYPES = "InvalidParameter.InvalidTargetNodeTypes"

	// 时间相关参数的取值或格式和预期不符。
	INVALIDPARAMETER_INVALIDTIMEPARAM = "InvalidParameter.InvalidTimeParam"

	// 调用链Id TraceUuid的取值和预期不符。
	INVALIDPARAMETER_INVALIDTRACEUUID = "InvalidParameter.InvalidTraceUuid"

	// 节点类型Type的取值和预期不符。
	INVALIDPARAMETER_INVALIDTYPE = "InvalidParameter.InvalidType"

	// Uin的取值和预期不符。
	INVALIDPARAMETER_INVALIDUIN = "InvalidParameter.InvalidUin"

	// 更新的索引元数据json UpdateMetaJson的取值和预期不符。
	INVALIDPARAMETER_INVALIDUPDATEMETAJSON = "InvalidParameter.InvalidUpdateMetaJson"

	// 更新类型UpdateType的取值和预期不符。
	INVALIDPARAMETER_INVALIDUPDATETYPE = "InvalidParameter.InvalidUpdateType"

	// 代金券ID列表VoucherIds的取值和预期不符。
	INVALIDPARAMETER_INVALIDVOUCHERIDS = "InvalidParameter.InvalidVoucherIds"

	// 可用区Zone的信息与预期不符。
	INVALIDPARAMETER_INVALIDZONE = "InvalidParameter.InvalidZone"

	// 空间名已存在。
	INVALIDPARAMETER_SPACENAMEEXIST = "InvalidParameter.SpaceNameExist"

	// 空间名不合法。
	INVALIDPARAMETER_SPACENAMEINVALID = "InvalidParameter.SpaceNameInvalid"

	// 用户名已存在。
	INVALIDPARAMETER_USERNAMEEXIST = "InvalidParameter.UsernameExist"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 包年包月购买时长ChargePeriod的取值有误。
	INVALIDPARAMETERVALUE_CHARGEPERIOD = "InvalidParameterValue.ChargePeriod"

	// 计费类型ChargeType的参数取值有误。
	INVALIDPARAMETERVALUE_CHARGETYPE = "InvalidParameterValue.ChargeType"

	// 是否只做升级检查CheckOnly的取值有误。
	INVALIDPARAMETERVALUE_CHECKONLY = "InvalidParameterValue.CheckOnly"

	// 组件配置值有误。
	INVALIDPARAMETERVALUE_CONFIGINFO = "InvalidParameterValue.ConfigInfo"

	// 是否允许双网卡EnableDoubleEni的值有误。
	INVALIDPARAMETERVALUE_ENABLEDOUBLEENI = "InvalidParameterValue.EnableDoubleEni"

	// ES配置类型取值有误。
	INVALIDPARAMETERVALUE_ESCONFIGTYPE = "InvalidParameterValue.EsConfigType"

	// 端口EsPort的取值有误。
	INVALIDPARAMETERVALUE_ESPORT = "InvalidParameterValue.EsPort"

	// EsVip的取值有误。
	INVALIDPARAMETERVALUE_ESVIP = "InvalidParameterValue.EsVip"

	// 要安装的bundle信息列表InstallBundleList的取值有误。
	INVALIDPARAMETERVALUE_INSTALLBUNDLELIST = "InvalidParameterValue.InstallBundleList"

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

	// 要卸载的bundle名列表RemoveBundleList的取值有误。
	INVALIDPARAMETERVALUE_REMOVEBUNDLELIST = "InvalidParameterValue.RemoveBundleList"

	// 为集群稳定起见，当前只支持同时卸载1个用户插件。
	INVALIDPARAMETERVALUE_REMOVEPLUGINLIST = "InvalidParameterValue.RemovePluginList"

	// 自动续费标识RenewFlag的参数取值有误。
	INVALIDPARAMETERVALUE_RENEWFLAG = "InvalidParameterValue.RenewFlag"

	// 要安装的bundle的zip包解压所在的路径TargetPath的取值有误。
	INVALIDPARAMETERVALUE_TARGETPATH = "InvalidParameterValue.TargetPath"

	// 计费时长单位TimeUnit的取值有误。
	INVALIDPARAMETERVALUE_TIMEUNIT = "InvalidParameterValue.TimeUnit"

	// 升级方式UpgradeMode的取值有误。
	INVALIDPARAMETERVALUE_UPGRADEMODE = "InvalidParameterValue.UpgradeMode"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 该账号下的集群数超过限额。
	LIMITEXCEEDED_CLUSTERNUM = "LimitExceeded.ClusterNum"

	// 超过当天智能诊断最大次数。
	LIMITEXCEEDED_DIAGNOSECOUNT = "LimitExceeded.DiagnoseCount"

	// 节点磁盘块数超过最大限额。
	LIMITEXCEEDED_DISKCOUNT = "LimitExceeded.DiskCount"

	// 空间索引数超过配额。
	LIMITEXCEEDED_INDEXCOUNT = "LimitExceeded.IndexCount"

	// 空间索引每日创建数超过配额。
	LIMITEXCEEDED_INDEXCREATE = "LimitExceeded.IndexCreate"

	// 节点数量或索引存储量超过限额。
	LIMITEXCEEDED_NODENUMORINDICES = "LimitExceeded.NodeNumOrIndices"

	// 超过最大自定义插件数量，请先删除部分自定义插件。
	LIMITEXCEEDED_PLUGININSTALL = "LimitExceeded.PluginInstall"

	// 超出一个用户可以拥有的最大资源限制。
	LIMITEXCEEDED_RESOURCELIMIT = "LimitExceeded.ResourceLimit"

	// 空间数超过配额。
	LIMITEXCEEDED_SPACECOUNT = "LimitExceeded.SpaceCount"

	// 空间每日创建数超过配额。
	LIMITEXCEEDED_SPACECREATE = "LimitExceeded.SpaceCreate"

	// 更新参数值过多，超过限制。
	LIMITEXCEEDED_UPDATEITEMLIMIT = "LimitExceeded.UpdateItemLimit"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 地域错误
	REGIONERROR = "RegionError"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 存在诊断中的作业，请等待作业诊断完成后重试。
	RESOURCEINUSE_DIAGNOSE = "ResourceInUse.Diagnose"

	// 索引状态处于隔离状态，不能进行其他操作。
	RESOURCEINUSE_ISOLATEDSTATUS = "ResourceInUse.IsolatedStatus"

	// 订单被锁定。
	RESOURCEINUSE_ORDER = "ResourceInUse.Order"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 账户余额不足。
	RESOURCEINSUFFICIENT_BALANCE = "ResourceInsufficient.Balance"

	// 计算资源不足，可更换可用区或调整机型。
	RESOURCEINSUFFICIENT_CVM = "ResourceInsufficient.CVM"

	// cvm份额不足。
	RESOURCEINSUFFICIENT_CVMQUOTA = "ResourceInsufficient.CVMQuota"

	// clb资源不足
	RESOURCEINSUFFICIENT_CLB = "ResourceInsufficient.Clb"

	// es资源不足
	RESOURCEINSUFFICIENT_ES = "ResourceInsufficient.Es"

	// 隐藏可用区专用主节点资源不足。
	RESOURCEINSUFFICIENT_HIDDENZONE = "ResourceInsufficient.HiddenZone"

	// 子网剩余ip数量不足。
	RESOURCEINSUFFICIENT_SUBNET = "ResourceInsufficient.Subnet"

	// 可用的子网IP不足。
	RESOURCEINSUFFICIENT_SUBNETIP = "ResourceInsufficient.SubnetIp"

	// 可用区资源不足。
	RESOURCEINSUFFICIENT_ZONE = "ResourceInsufficient.Zone"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 账户信息未找到。
	RESOURCENOTFOUND_ACCOUNTINFONOTFOUND = "ResourceNotFound.AccountInfoNotFound"

	// 巴拉多（barad）指标资源未找到。
	RESOURCENOTFOUND_BARADMETRICNOTFOUND = "ResourceNotFound.BaradMetricNotFound"

	// CAM资源未找到。
	RESOURCENOTFOUND_CAMINFONOTFOUND = "ResourceNotFound.CAMInfoNotFound"

	// 负载均衡信息未找到。
	RESOURCENOTFOUND_CLBNOTFOUND = "ResourceNotFound.CLBNotFound"

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

	// 集群信息获取失败
	RESOURCENOTFOUND_ESINFONOTFOUND = "ResourceNotFound.EsInfoNotFound"

	// 密钥管理系统资源获取失败。
	RESOURCENOTFOUND_KMSNOTFOUND = "ResourceNotFound.KMSNotFound"

	// 获取oss资源失败。
	RESOURCENOTFOUND_OSSINFONOTFOUND = "ResourceNotFound.OssInfoNotFound"

	// 获取安全凭证服务资源失败。
	RESOURCENOTFOUND_STSNOTFOUND = "ResourceNotFound.STSNotFound"

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

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// Uin不在白名单中。
	UNAUTHORIZEDOPERATION_UINNOTINWHITELIST = "UnauthorizedOperation.UinNotInWhiteList"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 审计日志的状态冲突。
	UNSUPPORTEDOPERATION_AUDITSTATUSCONFLICT = "UnsupportedOperation.AuditStatusConflict"

	// 不支持该认证方式。
	UNSUPPORTEDOPERATION_BASICSECURITYTYPE = "UnsupportedOperation.BasicSecurityType"

	// 不支持变配节点或磁盘的配置。
	UNSUPPORTEDOPERATION_CHANGENODETYPE = "UnsupportedOperation.ChangeNodeType"

	// 不支持该操作，在集群中存在分配的分片。
	UNSUPPORTEDOPERATION_CLUSTERSHARDNUM = "UnsupportedOperation.ClusterShardNum"

	// 集群中存在部分索引处于关闭状态。
	UNSUPPORTEDOPERATION_CLUSTERSTATECLOSE = "UnsupportedOperation.ClusterStateClose"

	// 集群中存在部分索引无备份。
	UNSUPPORTEDOPERATION_CLUSTERSTATENOREPLICATION = "UnsupportedOperation.ClusterStateNoReplication"

	// 集群的状态不正常（不为绿）。
	UNSUPPORTEDOPERATION_CLUSTERSTATEUNHEALTH = "UnsupportedOperation.ClusterStateUnHealth"

	// 该版本不支持cos备份。集群需要重启来允许cos备份。
	UNSUPPORTEDOPERATION_COSBACKUP = "UnsupportedOperation.CosBackUp"

	// 有关机的cvm
	UNSUPPORTEDOPERATION_CVMSTATUSSTOPPED = "UnsupportedOperation.CvmStatusStopped"

	// 不支持该智能诊断任务。
	UNSUPPORTEDOPERATION_DIAGNOSEJOB = "UnsupportedOperation.DiagnoseJob"

	// 智能诊断未开启。
	UNSUPPORTEDOPERATION_DIAGNOSENOTOPEN = "UnsupportedOperation.DiagnoseNotOpen"

	// 节点磁盘可用空间不足
	UNSUPPORTEDOPERATION_DISKFREESPACEINSUFFICIENT = "UnsupportedOperation.DiskFreeSpaceInsufficient"

	// 不支持的读硬盘配置。
	UNSUPPORTEDOPERATION_DISKUSE = "UnsupportedOperation.DiskUse"

	// 不支持这个长度的editList的操作。
	UNSUPPORTEDOPERATION_EDITLISTLENGTH = "UnsupportedOperation.EditListLength"

	// 存在未完成的异构扩缩容流程
	UNSUPPORTEDOPERATION_HETEROSCALINGEXECUTING = "UnsupportedOperation.HeteroScalingExecuting"

	// 集群有索引设置index.routing.allocation.require._name,or _ip, or _id
	UNSUPPORTEDOPERATION_INDEXSETTINGSREQUIRESET = "UnsupportedOperation.IndexSettingsRequireSet"

	// 不支持该操作，服务类型有误。
	UNSUPPORTEDOPERATION_INSTANCETYPEERROR = "UnsupportedOperation.InstanceTypeError"

	// 不支持该操作，license的类型有误。
	UNSUPPORTEDOPERATION_LICENSEERROR = "UnsupportedOperation.LicenseError"

	// 不支持本地盘机型。
	UNSUPPORTEDOPERATION_LOCALDISK = "UnsupportedOperation.LocalDisk"

	// 不支持迁移操作。
	UNSUPPORTEDOPERATION_MIGRATE = "UnsupportedOperation.Migrate"

	// 不支持多可用区变配。
	UNSUPPORTEDOPERATION_MULTIZONESUPGRADE = "UnsupportedOperation.MultiZonesUpgrade"

	// 抱歉，您需要前往ES控制台购买页完成确认认证，我们才能为您提供所选的腾讯云服务。如果对于确认流程有任何疑问，可以联系我们：https://www.tencentcloud.com/zh/contact-us。
	UNSUPPORTEDOPERATION_NOTAGREEMENT = "UnsupportedOperation.NotAgreement"

	// 用户同步日志的配置passLogstashId为空，不支持该操作。
	UNSUPPORTEDOPERATION_PASSLOGSTASHID = "UnsupportedOperation.PassLogstashId"

	// 不支持该插件。
	UNSUPPORTEDOPERATION_PLUGIN = "UnsupportedOperation.Plugin"

	// 不支持对应Protocol的操作。
	UNSUPPORTEDOPERATION_PROTOCOL = "UnsupportedOperation.Protocol"

	// 不支持的读取速度。
	UNSUPPORTEDOPERATION_READRATE = "UnsupportedOperation.ReadRate"

	// 离线节点蓝绿变更存在风险，请直接选择重启“in-place”。
	UNSUPPORTEDOPERATION_RESTARTMODE = "UnsupportedOperation.RestartMode"

	// 纵向缩容太多资源。
	UNSUPPORTEDOPERATION_SCALEDOWNTOOMUCH = "UnsupportedOperation.ScaleDownTooMuch"

	// 不支持该操作，当前实例运行状态不为正常。
	UNSUPPORTEDOPERATION_STATUSNOTNORMAL = "UnsupportedOperation.StatusNotNormal"

	// 不支持该操作，实例状态有误。
	UNSUPPORTEDOPERATION_STATUSNOTSUPPORT = "UnsupportedOperation.StatusNotSupport"

	// 节点或磁盘数变更时不支持更新磁盘加密信息。
	UNSUPPORTEDOPERATION_UPDATEDISKENCRYPT = "UnsupportedOperation.UpdateDiskEncrypt"

	// 不支持更新磁盘类型。
	UNSUPPORTEDOPERATION_UPDATEDISKTYPE = "UnsupportedOperation.UpdateDiskType"

	// 不支持同时更新节点个数（NodeNum）和节点规格（NodeType）。
	UNSUPPORTEDOPERATION_UPDATENODENUMANDNODETYPE = "UnsupportedOperation.UpdateNodeNumAndNodeType"

	// VPC资源获取失败。
	UNSUPPORTEDOPERATION_VPCINFONOTFOUND = "UnsupportedOperation.VPCInfoNotFound"

	// 不支持的web组件。
	UNSUPPORTEDOPERATION_WEBSERVICETYPE = "UnsupportedOperation.WebServiceType"

	// 不支持的写入速度。
	UNSUPPORTEDOPERATION_WRITERATE = "UnsupportedOperation.WriteRate"
)
