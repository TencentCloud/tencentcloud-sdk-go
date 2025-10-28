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

package v20201016

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 请求未授权。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 告警策略通知模板已经绑定到了某个告警策略上。
	FAILEDOPERATION_BINDEDALARM = "FailedOperation.BindedAlarm"

	// 桶内无相应前缀文件，请使用正确的桶、文件前缀和压缩方式。
	FAILEDOPERATION_BUCKETNOFILE = "FailedOperation.BucketNoFile"

	// 调用云产品接口异常
	FAILEDOPERATION_CLOUDPRODUCTINVOCATIONERROR = "FailedOperation.CloudProductInvocationError"

	// 文件解压缩失败，请选择正确的压缩方式。
	FAILEDOPERATION_DECOMPRESSFILE = "FailedOperation.DecompressFile"

	// 文件下载失败，请稍后再试。
	FAILEDOPERATION_DOWNLOADFILE = "FailedOperation.DownLoadFile"

	// 获取文件列表失败，请稍后再试。
	FAILEDOPERATION_GETLISTFILE = "FailedOperation.GetListFile"

	// 检索日志触发最大条数限制。
	FAILEDOPERATION_GETLOGREACHLIMIT = "FailedOperation.GetlogReachLimit"

	// 低频不支持配置kv和tag索引。
	FAILEDOPERATION_INVALIDINDEXRULEFORSEARCHLOW = "FailedOperation.InValidIndexRuleForSearchLow"

	// 该告警策略状态异常，请检查下日志主题ID是否都存在。
	FAILEDOPERATION_INVALIDALARM = "FailedOperation.InvalidAlarm"

	// 检索游标已失效或不存在。
	FAILEDOPERATION_INVALIDCONTEXT = "FailedOperation.InvalidContext"

	// 离线存储保存时间不得少于7天。
	FAILEDOPERATION_INVALIDPERIOD = "FailedOperation.InvalidPeriod"

	// 相同的日志集已存在。
	FAILEDOPERATION_LOGSETCONFLICT = "FailedOperation.LogsetConflict"

	// 日志集下存在日志主题。
	FAILEDOPERATION_LOGSETNOTEMPTY = "FailedOperation.LogsetNotEmpty"

	// 无效的Content。
	FAILEDOPERATION_MISSINGCONTENT = "FailedOperation.MissingContent"

	// 修改的生命周期被禁止。
	FAILEDOPERATION_PERIODMODIFYFORBIDDEN = "FailedOperation.PeriodModifyForbidden"

	// 文件预览失败，请稍后再试。
	FAILEDOPERATION_PREVIEWFILE = "FailedOperation.PreviewFile"

	// 查询语句运行失败。
	FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"

	// 读取文件内容失败，请确认文件可读。
	FAILEDOPERATION_READFILE = "FailedOperation.ReadFile"

	// 读qps超过限制。
	FAILEDOPERATION_READQPSLIMIT = "FailedOperation.ReadQpsLimit"

	// 查询超时。
	FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"

	// 投递任务不允许重试。
	FAILEDOPERATION_SHIPPERTASKNOTTORETRY = "FailedOperation.ShipperTaskNotToRetry"

	// 查询语句解析错误。
	FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"

	// 请求标签服务限频。
	FAILEDOPERATION_TAGQPSLIMIT = "FailedOperation.TagQpsLimit"

	// 操作超时
	FAILEDOPERATION_TIMEOUT = "FailedOperation.Timeout"

	// 日志主题已关闭。
	FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"

	// topic创建中
	FAILEDOPERATION_TOPICCREATING = "FailedOperation.TopicCreating"

	// 日志主题已隔离。
	FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"

	// 写qps超过限制。
	FAILEDOPERATION_WRITEQPSLIMIT = "FailedOperation.WriteQpsLimit"

	// 写流量超过限制。
	FAILEDOPERATION_WRITETRAFFICLIMIT = "FailedOperation.WriteTrafficLimit"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// DB错误。
	INTERNALERROR_DBERROR = "InternalError.DbError"

	// 角色非法。
	INTERNALERROR_ILLEGALROLE = "InternalError.IllegalRole"

	// 检索错误
	INTERNALERROR_SEARCHERROR = "InternalError.SearchError"

	// 检索失败
	INTERNALERROR_SEARCHFAILED = "InternalError.SearchFailed"

	// 内部错误服务器繁忙
	INTERNALERROR_SERVERBUSY = "InternalError.ServerBusy"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 告警策略已经存在。
	INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"

	// 告警策略通知模板已经存在。
	INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"

	// 相同的采集配置规则已经存在。
	INVALIDPARAMETER_CONFIGCONFLICT = "InvalidParameter.ConfigConflict"

	// 无效的Content。
	INVALIDPARAMETER_CONTENT = "InvalidParameter.Content"

	// 仪表盘命名冲突。
	INVALIDPARAMETER_DASHBOARDNAMECONFLICT = "InvalidParameter.DashboardNameConflict"

	// 数据加工任务冲突。
	INVALIDPARAMETER_DATAFROMTASKCONFLICT = "InvalidParameter.DataFromTaskConflict"

	// 数据加工任务不存在。
	INVALIDPARAMETER_DATAFROMTASKNOTEXIST = "InvalidParameter.DataFromTaskNotExist"

	// 数据库唯一键冲突。
	INVALIDPARAMETER_DBDUPLICATION = "InvalidParameter.DbDuplication"

	// 导出任务已经存在。
	INVALIDPARAMETER_EXPORTCONFLICT = "InvalidParameter.ExportConflict"

	// 低频不支持配置kv和tag索引。
	INVALIDPARAMETER_INVALIDINDEXRULEFORSEARCHLOW = "InvalidParameter.InValidIndexRuleForSearchLow"

	// 指定日志主题已经存在索引规则。
	INVALIDPARAMETER_INDEXCONFLICT = "InvalidParameter.IndexConflict"

	// 无效的数据加工语句。
	INVALIDPARAMETER_INVALIDETLCONTENT = "InvalidParameter.InvalidEtlContent"

	// 相同的日志集已存在。
	INVALIDPARAMETER_LOGSETCONFLICT = "InvalidParameter.LogsetConflict"

	// 同名机器组已经存在。
	INVALIDPARAMETER_MACHINEGROUPCONFLICT = "InvalidParameter.MachineGroupConflict"

	// 参数错误。
	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"

	// 投递规则命名冲突。
	INVALIDPARAMETER_SHIPPERCONFLICT = "InvalidParameter.ShipperConflict"

	// 指定日志集下已经有同名的日志主题。
	INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 一个billing topic只能创建一个cos采集配置
	LIMITEXCEEDED_BILLINGCOSRECHARGEOUTOFLIMIT = "LimitExceeded.BillingCosRechargeOutOfLimit"

	// 一个uin只能创建一个billing topic
	LIMITEXCEEDED_BILLINGTOPICOUTOFLIMIT = "LimitExceeded.BillingTopicOutOfLimit"

	// 采集规则配置超过最大值限制。
	LIMITEXCEEDED_CONFIG = "LimitExceeded.Config"

	// 创建日志导出任务数量超出限制。
	LIMITEXCEEDED_EXPORT = "LimitExceeded.Export"

	// 索引操作超过频率限制。
	LIMITEXCEEDED_INDEXOPERATING = "LimitExceeded.IndexOperating"

	// 并发查询超过限制，单topic并发最大值15。
	LIMITEXCEEDED_LOGSEARCH = "LimitExceeded.LogSearch"

	// 日志大小超过限制。
	LIMITEXCEEDED_LOGSIZE = "LimitExceeded.LogSize"

	// 日志集数量超出限制。
	LIMITEXCEEDED_LOGSET = "LimitExceeded.Logset"

	// 机器组超过限制。
	LIMITEXCEEDED_MACHINEGROUP = "LimitExceeded.MachineGroup"

	// 机器组IP超过限制。
	LIMITEXCEEDED_MACHINEGROUPIP = "LimitExceeded.MachineGroupIp"

	// 机器组Label超过限制。
	LIMITEXCEEDED_MACHINEGROUPIPLABELS = "LimitExceeded.MachineGroupIpLabels"

	// 分区超过限制。
	LIMITEXCEEDED_PARTITION = "LimitExceeded.Partition"

	// 记录超过限制
	LIMITEXCEEDED_RECORDOUTOFLIMIT = "LimitExceeded.RecordOutOfLimit"

	// 检索内存超限。
	LIMITEXCEEDED_SEARCHRESOURCES = "LimitExceeded.SearchResources"

	// 检索接口返回的日志量太大， 超过20MB限制。
	LIMITEXCEEDED_SEARCHRESULTTOOLARGE = "LimitExceeded.SearchResultTooLarge"

	// 投递规则超出限制。
	LIMITEXCEEDED_SHIPPER = "LimitExceeded.Shipper"

	// tag超过限制。
	LIMITEXCEEDED_TAG = "LimitExceeded.Tag"

	// 日志主题数目超过限制。
	LIMITEXCEEDED_TOPIC = "LimitExceeded.Topic"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// ACL校验失败。
	OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"

	// 账户已销毁。
	OPERATIONDENIED_ACCOUNTDESTROY = "OperationDenied.AccountDestroy"

	// 账户欠费。
	OPERATIONDENIED_ACCOUNTISOLATE = "OperationDenied.AccountIsolate"

	// 账户不存在。
	OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"

	// 低频不支持告警。
	OPERATIONDENIED_ALARMNOTSUPPORTFORSEARCHLOW = "OperationDenied.AlarmNotSupportForSearchLow"

	// 字段没有开启分析功能。
	OPERATIONDENIED_ANALYSISSWITCHCLOSE = "OperationDenied.AnalysisSwitchClose"

	// 不支持新语法。
	OPERATIONDENIED_NEWSYNTAXNOTSUPPORTED = "OperationDenied.NewSyntaxNotSupported"

	// 通知模板已绑定告警，无法删除。
	OPERATIONDENIED_NOTICEHASALARM = "OperationDenied.NoticeHasAlarm"

	// 操作低频检索不支持。
	OPERATIONDENIED_OPERATIONNOTSUPPORTINSEARCHLOW = "OperationDenied.OperationNotSupportInSearchLow"

	// topic绑定了数据加工。
	OPERATIONDENIED_TOPICHASDATAFORMTASK = "OperationDenied.TopicHasDataFormTask"

	// topic绑定了函数投递。
	OPERATIONDENIED_TOPICHASDELIVERFUNCTION = "OperationDenied.TopicHasDeliverFunction"

	// topic配置有外部数据源配置，不能删除
	OPERATIONDENIED_TOPICHASEXTERNALDATASOURCECONFIG = "OperationDenied.TopicHasExternalDatasourceConfig"

	// topic绑定了scheduleSql任务。
	OPERATIONDENIED_TOPICHASSCHEDULESQLTASK = "OperationDenied.TopicHasScheduleSqlTask"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// agent version不存在。
	RESOURCENOTFOUND_AGENTVERSIONNOTEXIST = "ResourceNotFound.AgentVersionNotExist"

	// 告警策略不存在。
	RESOURCENOTFOUND_ALARMNOTEXIST = "ResourceNotFound.AlarmNotExist"

	// 告警策略通知模板不存在。
	RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"

	// 指定的采集规则配置不存在。
	RESOURCENOTFOUND_CONFIGNOTEXIST = "ResourceNotFound.ConfigNotExist"

	// 仪表盘记录不存在。
	RESOURCENOTFOUND_DASHBOARDRECORDNOTEXIST = "ResourceNotFound.DashboardRecordNotExist"

	// 仪表盘订阅记录不存在。
	RESOURCENOTFOUND_DASHBOARDSUBSCRIBERECORDNOTEXIST = "ResourceNotFound.DashboardSubscribeRecordNotExist"

	// 数据加工任务不存在。
	RESOURCENOTFOUND_DATAFROMTASKNOTEXIST = "ResourceNotFound.DataFromTaskNotExist"

	// 日志导出不存在。
	RESOURCENOTFOUND_EXPORTNOTEXIST = "ResourceNotFound.ExportNotExist"

	// 索引不存在。
	RESOURCENOTFOUND_INDEXNOTEXIST = "ResourceNotFound.IndexNotExist"

	// 指定的日志集不存在。
	RESOURCENOTFOUND_LOGSETNOTEXIST = "ResourceNotFound.LogsetNotExist"

	// 机器组不存在。
	RESOURCENOTFOUND_MACHINEGROUPNOTEXIST = "ResourceNotFound.MachineGroupNotExist"

	// 分区不存在。
	RESOURCENOTFOUND_PARTITIONNOTEXIST = "ResourceNotFound.PartitionNotExist"

	// 记录不存在
	RESOURCENOTFOUND_RECORDNOTEXIST = "ResourceNotFound.RecordNotExist"

	// 投递规则不存在。
	RESOURCENOTFOUND_SHIPPERNOTEXIST = "ResourceNotFound.ShipperNotExist"

	// 投递任务不存在。
	RESOURCENOTFOUND_SHIPPERTASKNOTEXIST = "ResourceNotFound.ShipperTaskNotExist"

	// 日志主题不存在。
	RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// billing topic不允许修改cos导入配置
	UNSUPPORTEDOPERATION_MODIFYBILLINGCOSRECHARGENOSUPPORT = "UnsupportedOperation.ModifyBillingCosRechargeNoSupport"
)
