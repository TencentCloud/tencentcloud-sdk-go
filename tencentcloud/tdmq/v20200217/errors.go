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

	// CMQ 后台服务错误。
	FAILEDOPERATION_CMQBACKENDERROR = "FailedOperation.CmqBackendError"

	// 创建vpc绑定关系失败。
	FAILEDOPERATION_CREATEBINDVPC = "FailedOperation.CreateBindVpc"

	// 创建集群失败。
	FAILEDOPERATION_CREATECLUSTER = "FailedOperation.CreateCluster"

	// 环境创建失败。
	FAILEDOPERATION_CREATEENVIRONMENT = "FailedOperation.CreateEnvironment"

	// 创建环境角色失败。
	FAILEDOPERATION_CREATEENVIRONMENTROLE = "FailedOperation.CreateEnvironmentRole"

	// 创建命名空间失败。
	FAILEDOPERATION_CREATENAMESPACE = "FailedOperation.CreateNamespace"

	// 创建producer出错。
	FAILEDOPERATION_CREATEPRODUCERERROR = "FailedOperation.CreateProducerError"

	// 创建TDMQ client的出错。
	FAILEDOPERATION_CREATEPULSARCLIENTERROR = "FailedOperation.CreatePulsarClientError"

	// 角色创建失败。
	FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"

	// 创建秘钥失败。
	FAILEDOPERATION_CREATESECRETKEY = "FailedOperation.CreateSecretKey"

	// 创建订阅关系失败。
	FAILEDOPERATION_CREATESUBSCRIPTION = "FailedOperation.CreateSubscription"

	// 主题创建失败。
	FAILEDOPERATION_CREATETOPIC = "FailedOperation.CreateTopic"

	// 删除集群失败。
	FAILEDOPERATION_DELETECLUSTER = "FailedOperation.DeleteCluster"

	// 删除环境角色失败。
	FAILEDOPERATION_DELETEENVIRONMENTROLES = "FailedOperation.DeleteEnvironmentRoles"

	// 环境删除失败。
	FAILEDOPERATION_DELETEENVIRONMENTS = "FailedOperation.DeleteEnvironments"

	// 删除命名空间失败。
	FAILEDOPERATION_DELETENAMESPACE = "FailedOperation.DeleteNamespace"

	// 角色删除失败。
	FAILEDOPERATION_DELETEROLES = "FailedOperation.DeleteRoles"

	// 删除订阅关系失败。
	FAILEDOPERATION_DELETESUBSCRIPTIONS = "FailedOperation.DeleteSubscriptions"

	// 主题删除失败。
	FAILEDOPERATION_DELETETOPICS = "FailedOperation.DeleteTopics"

	// 查询订阅数据失败。
	FAILEDOPERATION_DESCRIBESUBSCRIPTION = "FailedOperation.DescribeSubscription"

	// 获取环境属性失败。
	FAILEDOPERATION_GETENVIRONMENTATTRIBUTESFAILED = "FailedOperation.GetEnvironmentAttributesFailed"

	// 获取主题分区数失败。
	FAILEDOPERATION_GETTOPICPARTITIONSFAILED = "FailedOperation.GetTopicPartitionsFailed"

	// 最大消息超过1MB。
	FAILEDOPERATION_MAXMESSAGESIZEERROR = "FailedOperation.MaxMessageSizeError"

	// 上传的msgID错误。
	FAILEDOPERATION_MESSAGEIDERROR = "FailedOperation.MessageIDError"

	// 必须先清除关联命名空间才能继续操作。
	FAILEDOPERATION_NAMESPACEINUSE = "FailedOperation.NamespaceInUse"

	// PulsarAdminClient错误。
	FAILEDOPERATION_PULSARADMINERROR = "FailedOperation.PulsarAdminError"

	// 接收消息出错。
	FAILEDOPERATION_RECEIVEERROR = "FailedOperation.ReceiveError"

	// 接收消息超时，请重试。
	FAILEDOPERATION_RECEIVETIMEOUT = "FailedOperation.ReceiveTimeout"

	// 消息回溯设置失败。
	FAILEDOPERATION_RESETMSGSUBOFFSETBYTIMESTAMPFAILED = "FailedOperation.ResetMsgSubOffsetByTimestampFailed"

	// 必须先清除关联角色数据才能继续操作。
	FAILEDOPERATION_ROLEINUSE = "FailedOperation.RoleInUse"

	// 保存秘钥失败。
	FAILEDOPERATION_SAVESECRETKEY = "FailedOperation.SaveSecretKey"

	// 消息发送超时。
	FAILEDOPERATION_SENDMESSAGETIMEOUTERROR = "FailedOperation.SendMessageTimeoutError"

	// 发送消息失败。
	FAILEDOPERATION_SENDMSGFAILED = "FailedOperation.SendMsgFailed"

	// 设置消息保留策略失败。
	FAILEDOPERATION_SETRETENTIONPOLICY = "FailedOperation.SetRetentionPolicy"

	// 设置消息TTL失败。
	FAILEDOPERATION_SETTTL = "FailedOperation.SetTTL"

	// 必须先清除关联主题数据才能继续操作。
	FAILEDOPERATION_TOPICINUSE = "FailedOperation.TopicInUse"

	// 请使用partition topic。
	FAILEDOPERATION_TOPICTYPEERROR = "FailedOperation.TopicTypeError"

	// 环境更新失败。
	FAILEDOPERATION_UPDATEENVIRONMENT = "FailedOperation.UpdateEnvironment"

	// 更新环境角色失败。
	FAILEDOPERATION_UPDATEENVIRONMENTROLE = "FailedOperation.UpdateEnvironmentRole"

	// 角色更新失败。
	FAILEDOPERATION_UPDATEROLE = "FailedOperation.UpdateRole"

	// 主题更新失败。
	FAILEDOPERATION_UPDATETOPIC = "FailedOperation.UpdateTopic"

	// 必须先清除关联VPC路由数据才能继续操作。
	FAILEDOPERATION_VPCINUSE = "FailedOperation.VpcInUse"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// Broker服务异常。
	INTERNALERROR_BROKERSERVICE = "InternalError.BrokerService"

	// 获取属性失败。
	INTERNALERROR_GETATTRIBUTESFAILED = "InternalError.GetAttributesFailed"

	// 内部错误。
	INTERNALERROR_ILLEGALMESSAGE = "InternalError.IllegalMessage"

	// 重试可以成功。
	INTERNALERROR_RETRY = "InternalError.Retry"

	// 系统错误。
	INTERNALERROR_SYSTEMERROR = "InternalError.SystemError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 错误的分区数。
	INVALIDPARAMETER_PARTITION = "InvalidParameter.Partition"

	// 上传的 tenant name 错误。
	INVALIDPARAMETER_TENANTNOTFOUND = "InvalidParameter.TenantNotFound"

	// 没有获取到正确的 token。
	INVALIDPARAMETER_TOKENNOTFOUND = "InvalidParameter.TokenNotFound"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 与现有集群名称重复。
	INVALIDPARAMETERVALUE_CLUSTERNAMEDUPLICATION = "InvalidParameterValue.ClusterNameDuplication"

	// 参数值不在允许范围内。
	INVALIDPARAMETERVALUE_INVALIDPARAMS = "InvalidParameterValue.InvalidParams"

	// 必要参数没有传递。
	INVALIDPARAMETERVALUE_NEEDMOREPARAMS = "InvalidParameterValue.NeedMoreParams"

	// 无效的消息TTL值。
	INVALIDPARAMETERVALUE_TTL = "InvalidParameterValue.TTL"

	// 上传的topic name错误。
	INVALIDPARAMETERVALUE_TOPICNOTFOUND = "InvalidParameterValue.TopicNotFound"

	// 超过配额限制。
	LIMITEXCEEDED = "LimitExceeded"

	// 实例下集群数量超过限制。
	LIMITEXCEEDED_CLUSTERS = "LimitExceeded.Clusters"

	// 实例下环境数量超过限制。
	LIMITEXCEEDED_ENVIRONMENTS = "LimitExceeded.Environments"

	// 实例下命名空间数量超过限额。
	LIMITEXCEEDED_NAMESPACES = "LimitExceeded.Namespaces"

	// 超过剩余额度，请重新调整。
	LIMITEXCEEDED_RETENTIONSIZE = "LimitExceeded.RetentionSize"

	// 超过保留时间限制，请重新调整。
	LIMITEXCEEDED_RETENTIONTIME = "LimitExceeded.RetentionTime"

	// 实例下订阅者数量超过限制。
	LIMITEXCEEDED_SUBSCRIPTIONS = "LimitExceeded.Subscriptions"

	// 实例下主题数量超过限制。
	LIMITEXCEEDED_TOPICS = "LimitExceeded.Topics"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 必要参数没有传递。
	MISSINGPARAMETER_NEEDMOREPARAMS = "MissingParameter.NeedMoreParams"

	// 订阅仍在消费中。
	OPERATIONDENIED_CONSUMERRUNNING = "OperationDenied.ConsumerRunning"

	// 默认环境不允许操作。
	OPERATIONDENIED_DEFAULTENVIRONMENT = "OperationDenied.DefaultEnvironment"

	// 资源被占用。
	RESOURCEINUSE = "ResourceInUse"

	// 集群已存在。
	RESOURCEINUSE_CLUSTER = "ResourceInUse.Cluster"

	// 环境角色已存在。
	RESOURCEINUSE_ENVIRONMENTROLE = "ResourceInUse.EnvironmentRole"

	// 重名，命名空间已存在。
	RESOURCEINUSE_NAMESPACE = "ResourceInUse.Namespace"

	// 队列已存在。
	RESOURCEINUSE_QUEUE = "ResourceInUse.Queue"

	// 角色已存在。
	RESOURCEINUSE_ROLE = "ResourceInUse.Role"

	// 重名，订阅关系已存在。
	RESOURCEINUSE_SUBSCRIPTION = "ResourceInUse.Subscription"

	// 重名，主题已存在。
	RESOURCEINUSE_TOPIC = "ResourceInUse.Topic"

	// 资源不足。
	RESOURCEINSUFFICIENT = "ResourceInsufficient"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 服务的集群不存在。
	RESOURCENOTFOUND_BROKERCLUSTER = "ResourceNotFound.BrokerCluster"

	// 集群不存在。
	RESOURCENOTFOUND_CLUSTER = "ResourceNotFound.Cluster"

	// 环境不存在。
	RESOURCENOTFOUND_ENVIRONMENT = "ResourceNotFound.Environment"

	// 环境角色不存在。
	RESOURCENOTFOUND_ENVIRONMENTROLE = "ResourceNotFound.EnvironmentRole"

	// 命名空间不存在。
	RESOURCENOTFOUND_NAMESPACE = "ResourceNotFound.Namespace"

	// 角色不存在。
	RESOURCENOTFOUND_ROLE = "ResourceNotFound.Role"

	// 订阅关系不存在。
	RESOURCENOTFOUND_SUBSCRIPTION = "ResourceNotFound.Subscription"

	// 主题不存在。
	RESOURCENOTFOUND_TOPIC = "ResourceNotFound.Topic"

	// 资源不可用。
	RESOURCEUNAVAILABLE = "ResourceUnavailable"

	// 发货异常。
	RESOURCEUNAVAILABLE_CREATEFAILED = "ResourceUnavailable.CreateFailed"

	// 需要充值才可继续操作。
	RESOURCEUNAVAILABLE_FUNDREQUIRED = "ResourceUnavailable.FundRequired"

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
