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

package v20210416

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 投递目标创建失败，请检查函数状态。
	FAILEDOPERATION_CREATETRIGGER = "FailedOperation.CreateTrigger"

	// 删除连接器失败，请检查资源状态。
	FAILEDOPERATION_DELETECONNECTION = "FailedOperation.DeleteConnection"

	// ServiceError操作失败，请检查资源信息。
	FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 服务处理出错，请稍后重试。若无法解决，请联系智能客服或提交工单。
	INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"

	// 服务处理出错，请稍后重试。若无法解决，请联系智能客服或提交工单。
	INTERNALERROR_SYSTEM = "InternalError.System"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// Payload与规范不符，请修正后再试。
	INVALIDPARAMETER_PAYLOAD = "InvalidParameter.Payload"

	// AMPParams取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_AMPPARAMS = "InvalidParameterValue.AMPParams"

	// BatchEventCount取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_BATCHEVENTCOUNT = "InvalidParameterValue.BatchEventCount"

	// BatchTimeout取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_BATCHTIMEOUT = "InvalidParameterValue.BatchTimeout"

	// CallbackType取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_CALLBACKTYPE = "InvalidParameterValue.CallbackType"

	// CallbackWeComURL取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_CALLBACKWECOMURL = "InvalidParameterValue.CallbackWeComURL"

	// ConnectionDescription取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_CONNECTIONDESCRIPTION = "InvalidParameterValue.ConnectionDescription"

	// ConnectionId取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_CONNECTIONID = "InvalidParameterValue.ConnectionId"

	// ConnectionName取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_CONNECTIONNAME = "InvalidParameterValue.ConnectionName"

	// Description取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"

	// EventBusId取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"

	// EventBusName取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_EVENTBUSNAME = "InvalidParameterValue.EventBusName"

	// EventPattern取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_EVENTPATTERN = "InvalidParameterValue.EventPattern"

	// Filters取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_FILTERS = "InvalidParameterValue.Filters"

	// InvalidApiRequestConfig取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_INVALIDAPIREQUESTCONFIG = "InvalidParameterValue.InvalidApiRequestConfig"

	// Limit取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"

	// NoticeReceiverChannel取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_NOTICERECEIVERCHANNEL = "InvalidParameterValue.NoticeReceiverChannel"

	// NoticeReceiverTimeWindow取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_NOTICERECEIVERTIMEWINDOW = "InvalidParameterValue.NoticeReceiverTimeWindow"

	// NoticeReceiverUserIds取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_NOTICERECEIVERUSERIDS = "InvalidParameterValue.NoticeReceiverUserIds"

	// NoticeReceiverUserType取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_NOTICERECEIVERUSERTYPE = "InvalidParameterValue.NoticeReceiverUserType"

	// Offset取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_OFFSET = "InvalidParameterValue.Offset"

	// Order取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_ORDER = "InvalidParameterValue.Order"

	// OrderBy取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_ORDERBY = "InvalidParameterValue.OrderBy"

	// Qualifier取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_QUALIFIER = "InvalidParameterValue.Qualifier"

	// RuleId取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_RULEID = "InvalidParameterValue.RuleId"

	// RuleName取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_RULENAME = "InvalidParameterValue.RuleName"

	// TargetDescription取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_TARGETDESCRIPTION = "InvalidParameterValue.TargetDescription"

	// TargetId取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_TARGETID = "InvalidParameterValue.TargetId"

	// Type取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"

	// Connection数量达到限制，可提交工单申请提升限制。
	LIMITEXCEEDED_CONNECTION = "LimitExceeded.Connection"

	// EventBus数量达到限制，可提交工单申请提升限制。
	LIMITEXCEEDED_EVENTBUS = "LimitExceeded.EventBus"

	// Logset数量达到限制，可提交工单申请提升限制。
	LIMITEXCEEDED_LOGSET = "LimitExceeded.Logset"

	// ResourceLimit数量达到限制，可提交工单申请提升限制。
	LIMITEXCEEDED_RESOURCELIMIT = "LimitExceeded.ResourceLimit"

	// RouteOverLimit数量达到限制，可提交工单申请提升限制。
	LIMITEXCEEDED_ROUTEOVERLIMIT = "LimitExceeded.RouteOverLimit"

	// Rule数量达到限制，可提交工单申请提升限制。
	LIMITEXCEEDED_RULE = "LimitExceeded.Rule"

	// Target数量达到限制，可提交工单申请提升限制。
	LIMITEXCEEDED_TARGET = "LimitExceeded.Target"

	// 投递目标创建失败，函数触发器数量达到限制，可提交工单申请提升限制。
	LIMITEXCEEDED_TRIGGER = "LimitExceeded.Trigger"

	// 投递目标创建失败，检测到当前账号不存在，请确认您的账号状态。
	OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"

	// 操作失败，该资源不可修改或删除。
	OPERATIONDENIED_RESOURCEIMMUTABLE = "OperationDenied.ResourceImmutable"

	// 同一个地域下，云服务默认事件集不允许重复创建。
	RESOURCEINUSE_DEFAULTEVENTBUS = "ResourceInUse.DefaultEventBus"

	// 事件集下有规则或者连接器，无法删除，请删除所有规则和连接器后重试。
	RESOURCEINUSE_EVENTBUS = "ResourceInUse.EventBus"

	// 规则下有目标，无法删除，请删除所有目标后重试。
	RESOURCEINUSE_RULE = "ResourceInUse.Rule"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 未找到指定的Connection，请创建后再试。
	RESOURCENOTFOUND_CONNECTION = "ResourceNotFound.Connection"

	// 未找到指定事件集，请创建后再试。
	RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"

	// 未找到指定的Function，请创建后再试。
	RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"

	// 连接器删除失败，未找到指定 API 。
	RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"

	// 连接器操作失败，未找到指定 API 网关服务。
	RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"

	// 投递目标创建失败，未找到指定的命名空间，请创建后再试。
	RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"

	// 未找到指定的服务角色，请创建后再试。
	RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"

	// 未找到指定的Rule，请创建后再试。
	RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"

	// 未找到指定的Target，请创建后再试。
	RESOURCENOTFOUND_TARGET = "ResourceNotFound.Target"

	// 投递目标创建失败，未找到指定的服务版本，请创建后再试。
	RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"

	// 连接器创建失败，当前资源不可用。
	RESOURCEUNAVAILABLE_CONNECTION = "ResourceUnavailable.Connection"

	// 投递目标创建失败，当前资源不可用。
	RESOURCEUNAVAILABLE_TARGET = "ResourceUnavailable.Target"

	// 当前账号缺少 EB 操作权限，请登录 CAM 控制台进行授权。
	UNAUTHORIZEDOPERATION_CAM = "UnauthorizedOperation.CAM"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 连接器操作失败，接口不支持当前操作。
	UNSUPPORTEDOPERATION_INVALIDACTION = "UnsupportedOperation.InvalidAction"

	// 连接器创建失败，不支持当前后端服务类型。
	UNSUPPORTEDOPERATION_INVALIDENDPOINTTYPE = "UnsupportedOperation.InvalidEndpointType"
)
