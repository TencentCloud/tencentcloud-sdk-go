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

package v20180228

const (
	// 此产品的特有错误码

	// 客户端已离线。
	FAILEDOPERATION_AGENTOFFLINE = "FailedOperation.AgentOffline"

	// 关闭专业版失败。
	FAILEDOPERATION_CLOSEPROVERSION = "FailedOperation.CloseProVersion"

	// 创建获取端口任务失败。
	FAILEDOPERATION_CREATEOPENPORTTASK = "FailedOperation.CreateOpenPortTask"

	// 创建获取实时进程任务失败。
	FAILEDOPERATION_CREATEPROCESSTASK = "FailedOperation.CreateProcessTask"

	// 导出文件失败。
	FAILEDOPERATION_EXPORT = "FailedOperation.Export"

	// 询价失败。
	FAILEDOPERATION_INQUIRYPRICE = "FailedOperation.InquiryPrice"

	// 卸载主机。
	FAILEDOPERATION_MACHINEDELETE = "FailedOperation.MachineDelete"

	// 无专业版主机。
	FAILEDOPERATION_NOPROFESSIONHOST = "FailedOperation.NoProfessionHost"

	// 拉取实时端口任务不存在。
	FAILEDOPERATION_OPENPORTTASKNOTFOUND = "FailedOperation.OpenPortTaskNotFound"

	// 开通专业版失败。
	FAILEDOPERATION_OPENPROVERSION = "FailedOperation.OpenProVersion"

	// 多台主机隔离，部分或全部失败。
	FAILEDOPERATION_PARTSEPARATE = "FailedOperation.PartSeparate"

	// 不能关闭预付费模式专业版。
	FAILEDOPERATION_PREPAYMODE = "FailedOperation.PrePayMode"

	// 拉取实时进程任务不存在。
	FAILEDOPERATION_PROCESSTASKNOTFOUND = "FailedOperation.ProcessTaskNotFound"

	// 回复木马失败。
	FAILEDOPERATION_RECOVER = "FailedOperation.Recover"

	// 重新检测漏洞失败。
	FAILEDOPERATION_RESCANVUL = "FailedOperation.RescanVul"

	// 该主机已有重新检测任务正在执行中。
	FAILEDOPERATION_RESCANVULPROCESSINUSE = "FailedOperation.RescanVulProcessInUse"

	// 单台主机隔离失败。
	FAILEDOPERATION_SINGLESEPARATE = "FailedOperation.SingleSeparate"

	// 创建基线策略超过上限。
	FAILEDOPERATION_TOOMANYSTRATEGY = "FailedOperation.TooManyStrategy"

	// 交易失败。
	FAILEDOPERATION_TRADEERROR = "FailedOperation.TradeError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 操作数据失败。
	INTERNALERROR_MAINDBFAIL = "InternalError.MainDBFail"

	// 时间区间格式错误。
	INVALIDPARAMETER_DATERANGE = "InvalidParameter.DateRange"

	// 非法请求。
	INVALIDPARAMETER_ILLEGALREQUEST = "InvalidParameter.IllegalRequest"

	// 参数格式错误。
	INVALIDPARAMETER_INVALIDFORMAT = "InvalidParameter.InvalidFormat"

	// IP格式不合法。
	INVALIDPARAMETER_IPNOVALID = "InvalidParameter.IpNoValid"

	// 参数缺失。
	INVALIDPARAMETER_MISSINGPARAMETER = "InvalidParameter.MissingParameter"

	// 参数解析错误。
	INVALIDPARAMETER_PARSINGERROR = "InvalidParameter.ParsingError"

	// 端口格式不合法。
	INVALIDPARAMETER_PORTNOVALID = "InvalidParameter.PortNoValid"

	// 正则参数格式错误。
	INVALIDPARAMETER_REGEXRULEERROR = "InvalidParameter.RegexRuleError"

	// 进程名/目标IP/目标端口，不能同时为空。
	INVALIDPARAMETER_REVERSHELLKEYFIELDALLEMPTY = "InvalidParameter.ReverShellKeyFieldAllEmpty"

	// 规则类接口，主机IP不正确。
	INVALIDPARAMETER_RULEHOSTIPERR = "InvalidParameter.RuleHostipErr"

	// 标签名称长度不能超过15个字符。
	INVALIDPARAMETERVALUE_TAGNAMELENGTHLIMIT = "InvalidParameterValue.TagNameLengthLimit"

	// 超出批量添加数量。
	LIMITEXCEEDED_AREAQUOTA = "LimitExceeded.AreaQuota"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"
)
