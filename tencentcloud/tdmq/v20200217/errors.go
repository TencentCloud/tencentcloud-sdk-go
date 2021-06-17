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

package v20200217

const (
	// 此产品的特有错误码

	// CAM鉴权不通过。
	AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 环境创建失败。
	FAILEDOPERATION_CREATEENVIRONMENT = "FailedOperation.CreateEnvironment"

	// 创建producer出错。
	FAILEDOPERATION_CREATEPRODUCERERROR = "FailedOperation.CreateProducerError"

	// 创建TDMQ client的出错。
	FAILEDOPERATION_CREATEPULSARCLIENTERROR = "FailedOperation.CreatePulsarClientError"

	// 创建订阅关系失败。
	FAILEDOPERATION_CREATESUBSCRIPTION = "FailedOperation.CreateSubscription"

	// 主题创建失败。
	FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"

	// 环境删除失败。
	FAILEDOPERATION_DELETEENVIRONMENTS = "FailedOperation.DeleteEnvironments"

	// 删除订阅关系失败。
	FAILEDOPERATION_DELETESUBSCRIPTIONS = "FailedOperation.DeleteSubscriptions"

	// 主题删除失败。
	FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"

	// 获取环境属性失败。
	FAILEDOPERATION_GETENVIRONMENTATTRIBUTESFAILED = "FailedOperation.GetEnvironmentAttributesFailed"

	// 获取主题分区数失败。
	FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"

	// 消息回溯设置失败。
	FAILEDOPERATION_RESETMSGSUBOFFSETBYTIMESTAMPFAILED = "FailedOperation.ResetMsgSubOffsetByTimestampFailed"

	// 发送消息失败。
	FAILEDOPERATION_SENDMSGFAILED = "FailedOperation.SendMsgFailed"

	// 环境更新失败。
	FAILEDOPERATION_UPDATEENVIRONMENT = "FailedOperation.UpdateEnvironment"

	// 主题更新失败。
	FAILEDOPERATION_UPDATETOPIC = "FailedOperation.UpdateTopic"

	// 获取属性失败。
	INTERNALERROR_GETATTRIBUTESFAILED = "InternalError.GetAttributesFailed"

	// 重试可以成功。
	INTERNALERROR_RETRY = "InternalError.Retry"

	// 系统错误。
	INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"

	// 上传的 tenant name 错误。
	INVALIDPARAMETER_TENANTNOTFOUND = "InvalidParameter.TenantNotFound"

	// 没有获取到正确的 token。
	INVALIDPARAMETER_TOKENNOTFOUND = "InvalidParameter.TokenNotFound"

	// 参数值不在允许范围内。
	INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"

	// 必要参数没有传递。
	INVALIDPARAMETERVALUE_NEEDMOREPARAMS = "InvalidParameterValue.NeedMoreParams"

	// 上传的topic name错误。
	INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"

	// 实例下环境数量超过限制。
	LIMITEXCEEDED_ENVIRONMENTS = "LimitExceeded.Environments"

	// 实例下订阅者数量超过限制。
	LIMITEXCEEDED_SUBSCRIPTIONS = "LimitExceeded.Subscriptions"

	// 实例下主题数量超过限制。
	LIMITEXCEEDED_TOPICS = "LimitExceeded.Topics"

	// 必要参数没有传递。
	MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"

	// 默认环境不允许操作。
	OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"

	// 不允许创建以-dlq、-retry结尾的主题，-dql为死信队列，-retry为重试队列。
	OPERATIONDENIED_DLQORRETRYTOPIC = "OperationDenied.DlqOrRetryTopic"

	// 重名，环境已存在。
	RESOURCEINUSE_ENVIRONMENT = "ResourceInUse.Environment"

	// 重名，订阅关系已存在。
	RESOURCEINUSE_SUBSCRIPTION = "ResourceInUse.Subscription"

	// 重名，主题已存在。
	RESOURCEINUSE_TOPIC = "ResourceInUse.Topic"

	// 服务的集群不存在。
	RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"

	// 环境不存在。
	RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"

	// 角色不存在。
	RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"

	// 订阅关系不存在。
	RESOURCENOTFOUND_SUBSCRIPTION = "ResourceNotFound.Subscription"

	// 主题不存在。
	RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 系统升级。
	RESOURCEUNAVAILABLE_SYSTEMUPGRADE = "ResourceUnavailable.SystemUpgrade"

	// 资源售罄。
	RESOURCESSOLDOUT = "ResourcesSoldOut"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"
)
