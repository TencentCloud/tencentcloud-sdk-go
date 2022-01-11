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

package v20201016

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 告警策略通知模板已经绑定到了某个告警策略上。
	FAILEDOPERATION_BINDEDALARM = "FailedOperation.BindedAlarm"

	// 检索日志触发最大条数限制。
	FAILEDOPERATION_GETLOGREACHLIMIT = "FailedOperation.GetlogReachLimit"

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

	// 查询语句运行失败。
	FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"

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

	// 日志主题已关闭。
	FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"

	// 日志主题已隔离。
	FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"

	// 写qps超过限制。
	FAILEDOPERATION_WRITEQPSLIMIT = "FailedOperation.WriteQpsLimit"

	// 写流量超过限制。
	FAILEDOPERATION_WRITETRAFFICLIMIT = "FailedOperation.WriteTrafficLimit"

	// 内部错误。
	INTERNALERROR = "InternalError"

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

	// 低频不支持配置kv和tag索引。
	INVALIDPARAMETER_INVALIDINDEXRULEFORSEARCHLOW = "InvalidParameter.InValidIndexRuleForSearchLow"

	// 指定日志主题已经存在索引规则。
	INVALIDPARAMETER_INDEXCONFLICT = "InvalidParameter.IndexConflict"

	// 相同的日志集已存在。
	INVALIDPARAMETER_LOGSETCONFLICT = "InvalidParameter.LogsetConflict"

	// 同名机器组已经存在。
	INVALIDPARAMETER_MACHINEGROUPCONFLICT = "InvalidParameter.MachineGroupConflict"

	// 投递规则命名冲突。
	INVALIDPARAMETER_SHIPPERCONFLICT = "InvalidParameter.ShipperConflict"

	// 指定日志集下已经有同名的日志主题。
	INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 采集规则配置超过最大值限制。
	LIMITEXCEEDED_CONFIG = "LimitExceeded.Config"

	// 日志导出数量超出限制。
	LIMITEXCEEDED_EXPORT = "LimitExceeded.Export"

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
	LIMITEXCEEDED_MACHINEGROUPLABELS = "LimitExceeded.MachineGroupLabels"

	// 分区超过限制。
	LIMITEXCEEDED_PARTITION = "LimitExceeded.Partition"

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

	// 通知模板已绑定告警，无法删除。
	OPERATIONDENIED_NOTICEHASALARM = "OperationDenied.NoticeHasAlarm"

	// topic绑定了数据加工。
	OPERATIONDENIED_TOPICHASDATAFORMTASK = "OperationDenied.TopicHasDataFormTask"

	// topic绑定了函数投递。
	OPERATIONDENIED_TOPICHASDELIVERFUNCTION = "OperationDenied.TopicHasDeliverFunction"

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

	// 投递规则不存在。
	RESOURCENOTFOUND_SHIPPERNOTEXIST = "ResourceNotFound.ShipperNotExist"

	// 投递任务不存在。
	RESOURCENOTFOUND_SHIPPERTASKNOTEXIST = "ResourceNotFound.ShipperTaskNotExist"

	// 日志主题不存在。
	RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
