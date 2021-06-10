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

	// 查询语句运行失败。
	FAILEDOPERATION_QUERYERROR = "FailedOperation.QueryError"

	// 查询超时。
	FAILEDOPERATION_SEARCHTIMEOUT = "FailedOperation.SearchTimeout"

	// 查询语句解析错误。
	FAILEDOPERATION_SYNTAXERROR = "FailedOperation.SyntaxError"

	// 日志主题已关闭。
	FAILEDOPERATION_TOPICCLOSED = "FailedOperation.TopicClosed"

	// 日志主题已隔离。
	FAILEDOPERATION_TOPICISOLATED = "FailedOperation.TopicIsolated"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 告警策略已经存在。
	INVALIDPARAMETER_ALARMCONFLICT = "InvalidParameter.AlarmConflict"

	// 告警策略通知模板已经存在。
	INVALIDPARAMETER_ALARMNOTICECONFLICT = "InvalidParameter.AlarmNoticeConflict"

	// 指定日志主题已经存在索引规则。
	INVALIDPARAMETER_INDEXCONFLICT = "InvalidParameter.IndexConflict"

	// 同名机器组已经存在。
	INVALIDPARAMETER_MACHINEGROUPCONFLICT = "InvalidParameter.MachineGroupConflict"

	// 指定日志集下已经有同名的日志主题。
	INVALIDPARAMETER_TOPICCONFLICT = "InvalidParameter.TopicConflict"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 日志导出数量超出限制。
	LIMITEXCEEDED_EXPORT = "LimitExceeded.Export"

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

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// ACL校验失败。
	OPERATIONDENIED_ACLFAILED = "OperationDenied.ACLFailed"

	// 通知模版已绑定告警，无法删除。
	OPERATIONDENIED_NOTICEHASALARM = "OperationDenied.NoticeHasAlarm"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// agent version不存在。
	RESOURCENOTFOUND_AGENTVERSIONNOTEXIST = "ResourceNotFound.AgentVersionNotExist"

	// 告警策略不存在。
	RESOURCENOTFOUND_ALARMNOTEXIST = "ResourceNotFound.AlarmNotExist"

	// 告警策略通知模板不存在。
	RESOURCENOTFOUND_ALARMNOTICENOTEXIST = "ResourceNotFound.AlarmNoticeNotExist"

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

	// 日志主题不存在。
	RESOURCENOTFOUND_TOPICNOTEXIST = "ResourceNotFound.TopicNotExist"
)
