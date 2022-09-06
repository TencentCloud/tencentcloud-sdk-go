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

	// 添加私有链接失败。
	FAILEDOPERATION_ADDPRIVATELINK = "FailedOperation.AddPrivateLink"

	// 用户鉴权错误。
	FAILEDOPERATION_AUTHENTICATEUSERFAILED = "FailedOperation.AuthenticateUserFailed"

	// 投递目标创建失败，请检查函数状态。
	FAILEDOPERATION_CREATETRIGGER = "FailedOperation.CreateTrigger"

	// 删除连接器失败，请检查资源状态。
	FAILEDOPERATION_DELETECONNECTION = "FailedOperation.DeleteConnection"

	// 删除私有链接失败。
	FAILEDOPERATION_DELETEPRIVATELINK = "FailedOperation.DeletePrivateLink"

	// 规则删除失败，请检查资源信息，确认是资源否存在或状态正常。
	FAILEDOPERATION_DELETERULE = "FailedOperation.DeleteRule"

	// ES集群内部错误。
	FAILEDOPERATION_ESINTERNALERROR = "FailedOperation.ESInternalError"

	// ES集群操作失败。
	FAILEDOPERATION_ESREQUESTFAILED = "FailedOperation.ESRequestFailed"

	// ES索引模板冲突错误。
	FAILEDOPERATION_ESTEMPLATECONFLICT = "FailedOperation.ESTemplateConflict"

	// 规则与事件不匹配，请修正后再试。
	FAILEDOPERATION_ERRORFILTER = "FailedOperation.ErrorFilter"

	// 注册CLS服务失败。
	FAILEDOPERATION_REGISTERCLSSERVICE = "FailedOperation.RegisterCLSService"

	// ServiceError操作失败，请检查资源信息。
	FAILEDOPERATION_SERVICEERROR = "FailedOperation.ServiceError"

	// 标签绑定操作失败。
	FAILEDOPERATION_TAGRESOURCE = "FailedOperation.TagResource"

	// 申请标签配额失败。
	FAILEDOPERATION_TAGRESOURCEALLOCATEQUOTAS = "FailedOperation.TagResourceAllocateQuotas"

	// 标签解绑操作失败。
	FAILEDOPERATION_UNTAGRESOURCE = "FailedOperation.UnTagResource"

	// 连接器更新失败，请检查资源信息，确认是资源否存在或状态正常。
	FAILEDOPERATION_UPDATECONNECTION = "FailedOperation.UpdateConnection"

	// 规则更新失败，请检查资源信息，确认是资源否存在或状态正常。
	FAILEDOPERATION_UPDATERULE = "FailedOperation.UpdateRule"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 服务处理出错，请稍后重试。若无法解决，请联系智能客服或提交工单。
	INTERNALERROR_OSSEXCEPTION = "InternalError.OssException"

	// 服务处理出错，请稍后重试。若无法解决，请联系智能客服或提交工单。
	INTERNALERROR_SYSTEM = "InternalError.System"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// APIGateway连接器不支持开启或关闭操作，请修正后重试。
	INVALIDPARAMETER_ENABLEAPIGATEWAY = "InvalidParameter.EnableAPIGateway"

	// Payload与规范不符，请修正后再试。
	INVALIDPARAMETER_PAYLOAD = "InvalidParameter.Payload"

	// AMPParams取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_AMPPARAMS = "InvalidParameterValue.AMPParams"

	// BatchEventCount取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_BATCHEVENTCOUNT = "InvalidParameterValue.BatchEventCount"

	// BatchTimeout取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_BATCHTIMEOUT = "InvalidParameterValue.BatchTimeout"

	// Ckafka 目标配置参数取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_CKAFKATARGETPARAMS = "InvalidParameterValue.CKafkaTargetParams"

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

	// DTSParams取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_DTSPARAMS = "InvalidParameterValue.DTSParams"

	// 死信队列配置参数取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_DEADLETTERCONFIG = "InvalidParameterValue.DeadLetterConfig"

	// Description取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_DESCRIPTION = "InvalidParameterValue.Description"

	// ES目标参数错误。
	INVALIDPARAMETERVALUE_ELASTICSEARCHTARGETPARAMS = "InvalidParameterValue.ElasticSearchTargetParams"

	// EventBusId取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_EVENTBUSID = "InvalidParameterValue.EventBusId"

	// EventBusName取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_EVENTBUSNAME = "InvalidParameterValue.EventBusName"

	// EventPattern取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_EVENTPATTERN = "InvalidParameterValue.EventPattern"

	// 链路追踪配置参数取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_EVENTTRACECONFIG = "InvalidParameterValue.EventTraceConfig"

	// Filters取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_FILTERS = "InvalidParameterValue.Filters"

	// InvalidApiRequestConfig取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_INVALIDAPIREQUESTCONFIG = "InvalidParameterValue.InvalidApiRequestConfig"

	// 事件格式非法，请修正后再试。
	INVALIDPARAMETERVALUE_INVALIDEVENT = "InvalidParameterValue.InvalidEvent"

	// 非法的事件集，请检查后重试。
	INVALIDPARAMETERVALUE_INVALIDEVENTBUS = "InvalidParameterValue.InvalidEventBus"

	// 事件模式格式错误，请修正后再试。
	INVALIDPARAMETERVALUE_INVALIDFILTERRULE = "InvalidParameterValue.InvalidFilterRule"

	// 非法的匹配规则，请检查后重试。
	INVALIDPARAMETERVALUE_INVALIDPATTERN = "InvalidParameterValue.InvalidPattern"

	// Limit取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"

	// LinkMode取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_LINKMODE = "InvalidParameterValue.LinkMode"

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

	// 标签参数错误。
	INVALIDPARAMETERVALUE_TAGS = "InvalidParameterValue.Tags"

	// TargetDescription取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_TARGETDESCRIPTION = "InvalidParameterValue.TargetDescription"

	// TargetId取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_TARGETID = "InvalidParameterValue.TargetId"

	// 数据转换配置参数取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_TRANSFORMATIONID = "InvalidParameterValue.TransformationID"

	// 数据转换任务创建失败，配置参数取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_TRANSFORMATIONS = "InvalidParameterValue.Transformations"

	// Type取值与规范不符，请修正后再试。
	INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"

	// 事件投递失败，因欠费或违规等原因，账号已被禁用，请联系官网账户端客服处理。
	LIMITEXCEEDED_BANNEDACCOUNT = "LimitExceeded.BannedAccount"

	// 集群私有链接超限。
	LIMITEXCEEDED_CLUSTERPRIVATELINKEXCEEDED = "LimitExceeded.ClusterPrivateLinkExceeded"

	// Connection数量达到限制，可提交工单申请提升限制。
	LIMITEXCEEDED_CONNECTION = "LimitExceeded.Connection"

	// EventBus数量达到限制，可提交工单申请提升限制。
	LIMITEXCEEDED_EVENTBUS = "LimitExceeded.EventBus"

	// 资源创建失败，可冻结余额不足，请充值后重新创建。
	LIMITEXCEEDED_INSUFFICIENTBALANCE = "LimitExceeded.InsufficientBalance"

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

	// 用户私有链接超限。
	LIMITEXCEEDED_USERPRIVATELINKEXCEEDED = "LimitExceeded.UserPrivateLinkExceeded"

	// 投递目标创建失败，检测到当前账号不存在，请确认您的账号状态。
	OPERATIONDENIED_ACCOUNTNOTEXISTS = "OperationDenied.AccountNotExists"

	// 不允许在自定义事件集上创建基于EB默认日志集的CLS事件目标。
	OPERATIONDENIED_DEFAULTCLSRESOURCEUNSUPPORTED = "OperationDenied.DefaultCLSResourceUnsupported"

	// 不支持操作当前ES版本。
	OPERATIONDENIED_ESVERSIONUNSUPPORTED = "OperationDenied.ESVersionUnsupported"

	// 技术架构升级，该资源临时锁定中，预计持续3~5分钟，事件推送流程无影响。
	OPERATIONDENIED_EVENTBUSRESOURCEISLOCKED = "OperationDenied.EventBusResourceIsLocked"

	// 操作失败，该资源不可修改或删除。
	OPERATIONDENIED_RESOURCEIMMUTABLE = "OperationDenied.ResourceImmutable"

	// 当前用户账号类型暂不支持操作，请提交工单处理。
	OPERATIONDENIED_UNSUPPORTEDOPERATION = "OperationDenied.UnsupportedOperation"

	// 同一个地域下，云服务默认事件集不允许重复创建。
	RESOURCEINUSE_DEFAULTEVENTBUS = "ResourceInUse.DefaultEventBus"

	// 无法删除，因事件集下存在规则、连接器或归档，请删除所有规则、连接器和归档后重试。
	RESOURCEINUSE_EVENTBUS = "ResourceInUse.EventBus"

	// 规则下有目标，无法删除，请删除所有目标后重试。
	RESOURCEINUSE_RULE = "ResourceInUse.Rule"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 未找到指定的Connection，请创建后再试。
	RESOURCENOTFOUND_CONNECTION = "ResourceNotFound.Connection"

	// 未找到指定事件集，请创建后再试。
	RESOURCENOTFOUND_EVENTBUS = "ResourceNotFound.EventBus"

	// 事件集不存在或未配置规则，请检查后再试。
	RESOURCENOTFOUND_EVENTBUSNOTFOUND = "ResourceNotFound.EventBusNotFound"

	// 未找到指定的Function，请创建后再试。
	RESOURCENOTFOUND_FUNCTION = "ResourceNotFound.Function"

	// 连接器删除失败，未找到指定 API 。
	RESOURCENOTFOUND_INVALIDAPI = "ResourceNotFound.InvalidApi"

	// 连接器操作失败，未找到指定 API 网关服务。
	RESOURCENOTFOUND_INVALIDSERVICE = "ResourceNotFound.InvalidService"

	// 投递目标创建失败，未找到指定的命名空间，请创建后再试。
	RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"

	// 未找到网络资源关联记录。
	RESOURCENOTFOUND_NETASSOCIATION = "ResourceNotFound.NetAssociation"

	// 未找到privatelink记录。
	RESOURCENOTFOUND_PRIVATELINKRESOURCE = "ResourceNotFound.PrivateLinkResource"

	// 未找到指定的服务角色，请创建后再试。
	RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"

	// 未找到指定的Rule，请创建后再试。
	RESOURCENOTFOUND_RULE = "ResourceNotFound.Rule"

	// 未找到指定标签。
	RESOURCENOTFOUND_TAG = "ResourceNotFound.Tag"

	// 未找到指定的Target，请创建后再试。
	RESOURCENOTFOUND_TARGET = "ResourceNotFound.Target"

	// 未找到指定的转换任务，请创建后再试。
	RESOURCENOTFOUND_TRANSFORMATION = "ResourceNotFound.Transformation"

	// 投递目标创建失败，未找到指定的服务版本，请创建后再试。
	RESOURCENOTFOUND_VERSION = "ResourceNotFound.Version"

	// 连接器创建失败，当前资源不可用。
	RESOURCEUNAVAILABLE_CONNECTION = "ResourceUnavailable.Connection"

	// ES集群状态异常。
	RESOURCEUNAVAILABLE_ESUNHEALTH = "ResourceUnavailable.ESUnhealth"

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
