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

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AMQPClusterConfig struct {

	// 单Vhost TPS上限
	MaxTpsPerVHost *uint64 `json:"MaxTpsPerVHost,omitempty" name:"MaxTpsPerVHost"`

	// 单Vhost客户端连接数上限
	MaxConnNumPerVHost *uint64 `json:"MaxConnNumPerVHost,omitempty" name:"MaxConnNumPerVHost"`

	// 最大Vhost数量
	MaxVHostNum *uint64 `json:"MaxVHostNum,omitempty" name:"MaxVHostNum"`

	// 最大exchange数量
	MaxExchangeNum *uint64 `json:"MaxExchangeNum,omitempty" name:"MaxExchangeNum"`

	// 最大Queue数量
	MaxQueueNum *uint64 `json:"MaxQueueNum,omitempty" name:"MaxQueueNum"`

	// 消息最大保留时间，以毫秒为单位
	MaxRetentionTime *uint64 `json:"MaxRetentionTime,omitempty" name:"MaxRetentionTime"`

	// 已使用Vhost数量
	UsedVHostNum *uint64 `json:"UsedVHostNum,omitempty" name:"UsedVHostNum"`

	// 已使用exchange数量
	UsedExchangeNum *uint64 `json:"UsedExchangeNum,omitempty" name:"UsedExchangeNum"`

	// 已使用queue数量
	UsedQueueNum *uint64 `json:"UsedQueueNum,omitempty" name:"UsedQueueNum"`
}

type AMQPClusterDetail struct {

	// 集群基本信息
	Info *AMQPClusterInfo `json:"Info,omitempty" name:"Info"`

	// 集群配置信息
	Config *AMQPClusterConfig `json:"Config,omitempty" name:"Config"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type AMQPClusterInfo struct {

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 地域信息
	Region *string `json:"Region,omitempty" name:"Region"`

	// 创建时间，毫秒为单位
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 集群说明信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 公网接入地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicEndPoint *string `json:"PublicEndPoint,omitempty" name:"PublicEndPoint"`

	// VPC接入地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcEndPoint *string `json:"VpcEndPoint,omitempty" name:"VpcEndPoint"`
}

type AMQPClusterRecentStats struct {

	// Queue数量
	QueueNum *uint64 `json:"QueueNum,omitempty" name:"QueueNum"`

	// 消息生产数
	ProducedMsgNum *uint64 `json:"ProducedMsgNum,omitempty" name:"ProducedMsgNum"`

	// 消息堆积数
	AccumulativeMsgNum *uint64 `json:"AccumulativeMsgNum,omitempty" name:"AccumulativeMsgNum"`

	// Exchange数量
	ExchangeNum *uint64 `json:"ExchangeNum,omitempty" name:"ExchangeNum"`
}

type AMQPExchange struct {

	// Exchange名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// Exchange的类别，为枚举类型:Direct, Fanout, Topic
	Type *string `json:"Type,omitempty" name:"Type"`

	// 主绑定数
	SourceBindedNum *uint64 `json:"SourceBindedNum,omitempty" name:"SourceBindedNum"`

	// 说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 被绑定数
	DestBindedNum *uint64 `json:"DestBindedNum,omitempty" name:"DestBindedNum"`

	// 创建时间，以毫秒为单位
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 创建时间，以毫秒为单位
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 是否为内部Exchange(以amq.前缀开头的)
	Internal *bool `json:"Internal,omitempty" name:"Internal"`
}

type AMQPQueueDetail struct {

	// Queue名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 被绑定数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestBindedNum *uint64 `json:"DestBindedNum,omitempty" name:"DestBindedNum"`

	// 创建时间，以毫秒为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 创建时间，以毫秒为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 在线消费者数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OnlineConsumerNum *uint64 `json:"OnlineConsumerNum,omitempty" name:"OnlineConsumerNum"`

	// 每秒钟的事务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tps *uint64 `json:"Tps,omitempty" name:"Tps"`

	// 消息堆积数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccumulativeMsgNum *uint64 `json:"AccumulativeMsgNum,omitempty" name:"AccumulativeMsgNum"`

	// 是否自动删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoDelete *bool `json:"AutoDelete,omitempty" name:"AutoDelete"`

	// 死信交换机
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterExchange *string `json:"DeadLetterExchange,omitempty" name:"DeadLetterExchange"`

	// 死信交换机路由键
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterRoutingKey *string `json:"DeadLetterRoutingKey,omitempty" name:"DeadLetterRoutingKey"`
}

type AMQPRouteRelation struct {

	// 路由关系ID
	RouteRelationId *string `json:"RouteRelationId,omitempty" name:"RouteRelationId"`

	// 源Exchange
	SourceExchange *string `json:"SourceExchange,omitempty" name:"SourceExchange"`

	// 目标类型:Queue|Exchange
	DestType *string `json:"DestType,omitempty" name:"DestType"`

	// 目标值
	DestValue *string `json:"DestValue,omitempty" name:"DestValue"`

	// 绑定key
	RoutingKey *string `json:"RoutingKey,omitempty" name:"RoutingKey"`

	// 源路由类型:Direct|Topic|Fanout
	SourceExchangeType *string `json:"SourceExchangeType,omitempty" name:"SourceExchangeType"`

	// 创建时间，以毫秒为单位
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间，以毫秒为单位
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 说明信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type AMQPVHost struct {

	// 命名空间名称，3-64个字符，只能包含字母、数字、“-”及“_”
	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`

	// 未消费消息的保留时间，以毫秒单位，范围60秒到15天
	MsgTtl *uint64 `json:"MsgTtl,omitempty" name:"MsgTtl"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 创建时间，以毫秒为单位
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间，以毫秒为单位
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 用户名
	Username *string `json:"Username,omitempty" name:"Username"`

	// 密码
	Password *string `json:"Password,omitempty" name:"Password"`
}

type AcknowledgeMessageRequest struct {
	*tchttp.BaseRequest

	// 用作标识消息的唯一的ID（可从 receiveMessage 的返回值中获得）
	MessageId *string `json:"MessageId,omitempty" name:"MessageId"`

	// Topic 名字（可从 receiveMessage 的返回值中获得）这里尽量需要使用topic的全路径，即：tenant/namespace/topic。如果不指定，默认使用的是：public/default
	AckTopic *string `json:"AckTopic,omitempty" name:"AckTopic"`

	// 订阅者的名字，可以从receiveMessage的返回值中获取到。这里尽量与receiveMessage中的订阅者保持一致，否则没办法正确ack 接收回来的消息。
	SubName *string `json:"SubName,omitempty" name:"SubName"`
}

func (r *AcknowledgeMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcknowledgeMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MessageId")
	delete(f, "AckTopic")
	delete(f, "SubName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AcknowledgeMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AcknowledgeMessageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 如果为“”，则说明没有错误返回
	// 注意：此字段可能返回 null，表示取不到有效值。
		ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AcknowledgeMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcknowledgeMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindCluster struct {

	// 物理集群的名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type BundleSetOpt struct {
}

type ClearCmqQueueRequest struct {
	*tchttp.BaseRequest

	// 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *ClearCmqQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearCmqQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClearCmqQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ClearCmqQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearCmqQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearCmqQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClearCmqSubscriptionFilterTagsRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 订阅名字，在单个地域同一帐号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
}

func (r *ClearCmqSubscriptionFilterTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearCmqSubscriptionFilterTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "SubscriptionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClearCmqSubscriptionFilterTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ClearCmqSubscriptionFilterTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearCmqSubscriptionFilterTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearCmqSubscriptionFilterTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Cluster struct {

	// 集群Id。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 说明信息。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 接入点数量
	EndPointNum *int64 `json:"EndPointNum,omitempty" name:"EndPointNum"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 集群是否健康，1表示健康，0表示异常
	Healthy *int64 `json:"Healthy,omitempty" name:"Healthy"`

	// 集群健康信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthyInfo *string `json:"HealthyInfo,omitempty" name:"HealthyInfo"`

	// 集群状态，0:创建中，1:正常，2:删除中，3:已删除，5:创建失败，6: 删除失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 最大命名空间数量
	MaxNamespaceNum *int64 `json:"MaxNamespaceNum,omitempty" name:"MaxNamespaceNum"`

	// 最大Topic数量
	MaxTopicNum *int64 `json:"MaxTopicNum,omitempty" name:"MaxTopicNum"`

	// 最大QPS
	MaxQps *int64 `json:"MaxQps,omitempty" name:"MaxQps"`

	// 最大消息保留时间，秒为单位
	MessageRetentionTime *int64 `json:"MessageRetentionTime,omitempty" name:"MessageRetentionTime"`

	// 最大存储容量
	MaxStorageCapacity *int64 `json:"MaxStorageCapacity,omitempty" name:"MaxStorageCapacity"`

	// 集群版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitempty" name:"Version"`

	// 公网访问接入点
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicEndPoint *string `json:"PublicEndPoint,omitempty" name:"PublicEndPoint"`

	// VPC访问接入点
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcEndPoint *string `json:"VpcEndPoint,omitempty" name:"VpcEndPoint"`

	// 命名空间数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceNum *int64 `json:"NamespaceNum,omitempty" name:"NamespaceNum"`

	// 已使用存储限制，MB为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedStorageBudget *int64 `json:"UsedStorageBudget,omitempty" name:"UsedStorageBudget"`

	// 最大生产消息速率，以条数为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxPublishRateInMessages *int64 `json:"MaxPublishRateInMessages,omitempty" name:"MaxPublishRateInMessages"`

	// 最大推送消息速率，以条数为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDispatchRateInMessages *int64 `json:"MaxDispatchRateInMessages,omitempty" name:"MaxDispatchRateInMessages"`

	// 最大生产消息速率，以字节为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxPublishRateInBytes *int64 `json:"MaxPublishRateInBytes,omitempty" name:"MaxPublishRateInBytes"`

	// 最大推送消息速率，以字节为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDispatchRateInBytes *int64 `json:"MaxDispatchRateInBytes,omitempty" name:"MaxDispatchRateInBytes"`

	// 已创建主题数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicNum *int64 `json:"TopicNum,omitempty" name:"TopicNum"`

	// 最长消息延时，以秒为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMessageDelayInSeconds *int64 `json:"MaxMessageDelayInSeconds,omitempty" name:"MaxMessageDelayInSeconds"`

	// 是否开启公网访问，不填时默认开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitempty" name:"PublicAccessEnabled"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CmqDeadLetterPolicy struct {

	// 死信队列。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterQueue *string `json:"DeadLetterQueue,omitempty" name:"DeadLetterQueue"`

	// 死信队列策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Policy *uint64 `json:"Policy,omitempty" name:"Policy"`

	// 最大未消费过期时间。Policy为1时必选。范围300-43200，单位秒，需要小于消息最大保留时间MsgRetentionSeconds。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitempty" name:"MaxTimeToLive"`

	// 最大接收次数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitempty" name:"MaxReceiveCount"`
}

type CmqDeadLetterSource struct {

	// 消息队列ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueId *string `json:"QueueId,omitempty" name:"QueueId"`

	// 消息队列名字。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

type CmqQueue struct {

	// 消息队列ID。
	QueueId *string `json:"QueueId,omitempty" name:"QueueId"`

	// 消息队列名字。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 每秒钟生产消息条数的限制，消费消息的大小是该值的1.1倍。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`

	// 带宽限制。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bps *uint64 `json:"Bps,omitempty" name:"Bps"`

	// 飞行消息最大保留时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDelaySeconds *uint64 `json:"MaxDelaySeconds,omitempty" name:"MaxDelaySeconds"`

	// 最大堆积消息数。取值范围在公测期间为 1,000,000 - 10,000,000，正式上线后范围可达到 1000,000-1000,000,000。默认取值在公测期间为 10,000,000，正式上线后为 100,000,000。
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitempty" name:"MaxMsgHeapNum"`

	// 消息接收长轮询等待时间。取值范围0 - 30秒，默认值0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitempty" name:"PollingWaitSeconds"`

	// 消息保留周期。取值范围60-1296000秒（1min-15天），默认值345600秒（4 天）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// 消息可见性超时。取值范围1 - 43200秒（即12小时内），默认值30。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitempty" name:"VisibilityTimeout"`

	// 消息最大长度。取值范围1024 - 1048576 Byte（即1K - 1024K），默认值65536。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// 回溯队列的消息回溯时间最大值，取值范围0 - 43200秒，0表示不开启消息回溯。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RewindSeconds *uint64 `json:"RewindSeconds,omitempty" name:"RewindSeconds"`

	// 队列的创建时间。返回 Unix 时间戳，精确到毫秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最后一次修改队列属性的时间。返回 Unix 时间戳，精确到毫秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyTime *uint64 `json:"LastModifyTime,omitempty" name:"LastModifyTime"`

	// 在队列中处于 Active 状态（不处于被消费状态）的消息总数，为近似值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveMsgNum *uint64 `json:"ActiveMsgNum,omitempty" name:"ActiveMsgNum"`

	// 在队列中处于 Inactive 状态（正处于被消费状态）的消息总数，为近似值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InactiveMsgNum *uint64 `json:"InactiveMsgNum,omitempty" name:"InactiveMsgNum"`

	// 延迟消息数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelayMsgNum *uint64 `json:"DelayMsgNum,omitempty" name:"DelayMsgNum"`

	// 已调用 DelMsg 接口删除，但还在回溯保留时间内的消息数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RewindMsgNum *uint64 `json:"RewindMsgNum,omitempty" name:"RewindMsgNum"`

	// 消息最小未消费时间，单位为秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinMsgTime *uint64 `json:"MinMsgTime,omitempty" name:"MinMsgTime"`

	// 事务消息队列。true表示是事务消息，false表示不是事务消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Transaction *bool `json:"Transaction,omitempty" name:"Transaction"`

	// 死信队列。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterSource []*CmqDeadLetterSource `json:"DeadLetterSource,omitempty" name:"DeadLetterSource"`

	// 死信队列策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterPolicy *CmqDeadLetterPolicy `json:"DeadLetterPolicy,omitempty" name:"DeadLetterPolicy"`

	// 事务消息策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransactionPolicy *CmqTransactionPolicy `json:"TransactionPolicy,omitempty" name:"TransactionPolicy"`

	// 创建者Uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 关联的标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 消息轨迹。true表示开启，false表示不开启。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trace *bool `json:"Trace,omitempty" name:"Trace"`

	// 租户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
}

type CmqSubscription struct {

	// 订阅名字，在单个地域同一帐号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`

	// 订阅 ID。订阅 ID 在拉取监控数据时会用到。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubscriptionId *string `json:"SubscriptionId,omitempty" name:"SubscriptionId"`

	// 订阅拥有者的 APPID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicOwner *uint64 `json:"TopicOwner,omitempty" name:"TopicOwner"`

	// 该订阅待投递的消息数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgCount *uint64 `json:"MsgCount,omitempty" name:"MsgCount"`

	// 最后一次修改订阅属性的时间。返回 Unix 时间戳，精确到毫秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyTime *uint64 `json:"LastModifyTime,omitempty" name:"LastModifyTime"`

	// 订阅的创建时间。返回 Unix 时间戳，精确到毫秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 表示订阅接收消息的过滤策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindingKey []*string `json:"BindingKey,omitempty" name:"BindingKey"`

	// 接收通知的 endpoint，根据协议 protocol 区分：对于 HTTP，endpoint 必须以http://开头，host 可以是域名或 IP；对于 queue，则填 queueName。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// 描述用户创建订阅时选择的过滤策略：
	// filterType = 1表示用户使用 filterTag 标签过滤
	// filterType = 2表示用户使用 bindingKey 过滤。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterTags []*string `json:"FilterTags,omitempty" name:"FilterTags"`

	// 订阅的协议，目前支持两种协议：HTTP、queue。使用 HTTP 协议，用户需自己搭建接受消息的 Web Server。使用 queue，消息会自动推送到 CMQ queue，用户可以并发地拉取消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 向 endpoint 推送消息出现错误时，CMQ 推送服务器的重试策略。取值有：
	// （1）BACKOFF_RETRY，退避重试。每隔一定时间重试一次，重试够一定次数后，就把该消息丢弃，继续推送下一条消息；
	// （2）EXPONENTIAL_DECAY_RETRY，指数衰退重试。每次重试的间隔是指数递增的，例如开始 1s，后面是 2s，4s，8s...由于 Topic 消息的周期是一天，所以最多重试一天就把消息丢弃。默认值是 EXPONENTIAL_DECAY_RETRY。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" name:"NotifyStrategy"`

	// 推送内容的格式。取值：（1）JSON；（2）SIMPLIFIED，即 raw 格式。如果 protocol 是 queue，则取值必须为 SIMPLIFIED。如果 protocol 是 HTTP，两个值均可以，默认值是 JSON。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" name:"NotifyContentFormat"`
}

type CmqTopic struct {

	// 主题的 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 主题名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 消息在主题中最长存活时间，从发送到该主题开始经过此参数指定的时间后，不论消息是否被成功推送给用户都将被删除，单位为秒。固定为一天（86400秒），该属性不能修改。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// 消息最大长度。取值范围1024 - 1048576Byte（即1 - 1024K），默认值为65536。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// 每秒钟发布消息的条数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`

	// 描述用户创建订阅时选择的过滤策略：
	// FilterType = 1表示用户使用 FilterTag 标签过滤;
	// FilterType = 2表示用户使用 BindingKey 过滤。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterType *uint64 `json:"FilterType,omitempty" name:"FilterType"`

	// 主题的创建时间。返回 Unix 时间戳，精确到毫秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最后一次修改主题属性的时间。返回 Unix 时间戳，精确到毫秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyTime *uint64 `json:"LastModifyTime,omitempty" name:"LastModifyTime"`

	// 当前该主题中消息数目（消息堆积数）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgCount *uint64 `json:"MsgCount,omitempty" name:"MsgCount"`

	// 创建者 Uin，CAM 鉴权 resource 由该字段组合而成。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 关联的标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 消息轨迹。true表示开启，false表示不开启。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trace *bool `json:"Trace,omitempty" name:"Trace"`

	// 租户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`

	// 命名空间名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
}

type CmqTransactionPolicy struct {

	// 第一次回查时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitempty" name:"FirstQueryInterval"`

	// 最大查询次数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitempty" name:"MaxQueryCount"`
}

type Connection struct {

	// 生产者地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitempty" name:"Address"`

	// 主题分区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions *int64 `json:"Partitions,omitempty" name:"Partitions"`

	// 生产者版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientVersion *string `json:"ClientVersion,omitempty" name:"ClientVersion"`

	// 生产者名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`

	// 生产者ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProducerId *string `json:"ProducerId,omitempty" name:"ProducerId"`

	// 消息平均大小(byte)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AverageMsgSize *string `json:"AverageMsgSize,omitempty" name:"AverageMsgSize"`

	// 生成速率(byte/秒)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputIn *string `json:"MsgThroughputIn,omitempty" name:"MsgThroughputIn"`
}

type Consumer struct {

	// 消费者开始连接的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectedSince *string `json:"ConnectedSince,omitempty" name:"ConnectedSince"`

	// 消费者地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerAddr *string `json:"ConsumerAddr,omitempty" name:"ConsumerAddr"`

	// 消费者名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerName *string `json:"ConsumerName,omitempty" name:"ConsumerName"`

	// 消费者版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientVersion *string `json:"ClientVersion,omitempty" name:"ClientVersion"`
}

type ConsumersSchedule struct {

	// 当前分区id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`

	// 消息数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumberOfEntries *uint64 `json:"NumberOfEntries,omitempty" name:"NumberOfEntries"`

	// 消息积压数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgBacklog *uint64 `json:"MsgBacklog,omitempty" name:"MsgBacklog"`

	// 消费者每秒分发消息的数量之和。
	MsgRateOut *string `json:"MsgRateOut,omitempty" name:"MsgRateOut"`

	// 消费者每秒消息的byte。
	MsgThroughputOut *string `json:"MsgThroughputOut,omitempty" name:"MsgThroughputOut"`

	// 超时丢弃比例。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateExpired *string `json:"MsgRateExpired,omitempty" name:"MsgRateExpired"`
}

type CreateAMQPClusterRequest struct {
	*tchttp.BaseRequest

	// 3-64个字符，只能包含字母、数字、“-”及“_”
	Name *string `json:"Name,omitempty" name:"Name"`

	// 集群描述，128个字符以内
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateAMQPClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAMQPClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAMQPClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAMQPClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAMQPClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAMQPClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAMQPExchangeRequest struct {
	*tchttp.BaseRequest

	// 交换机名称，3-64个字符，只能包含字母、数字、“-”及“_”
	Exchange *string `json:"Exchange,omitempty" name:"Exchange"`

	// 交换机所在的vhost，目前支持在单个vhost下创建主题
	VHosts []*string `json:"VHosts,omitempty" name:"VHosts"`

	// 交换机类型，可选值为Direct, Fanout, Topic
	Type *string `json:"Type,omitempty" name:"Type"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 交换机说明，最大128个字符
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 备用交换机名称
	AlternateExchange *string `json:"AlternateExchange,omitempty" name:"AlternateExchange"`
}

func (r *CreateAMQPExchangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAMQPExchangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Exchange")
	delete(f, "VHosts")
	delete(f, "Type")
	delete(f, "ClusterId")
	delete(f, "Remark")
	delete(f, "AlternateExchange")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAMQPExchangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAMQPExchangeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAMQPExchangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAMQPExchangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAMQPQueueRequest struct {
	*tchttp.BaseRequest

	// 队列名称，3-64个字符，只能包含字母、数字、“-”及“_”
	Queue *string `json:"Queue,omitempty" name:"Queue"`

	// 队列所在的vhost名称
	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`

	// 是否自动清除
	AutoDelete *bool `json:"AutoDelete,omitempty" name:"AutoDelete"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 队列说明，最大128个字符
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 死信exchange
	DeadLetterExchange *string `json:"DeadLetterExchange,omitempty" name:"DeadLetterExchange"`

	// 路由键
	DeadLetterRoutingKey *string `json:"DeadLetterRoutingKey,omitempty" name:"DeadLetterRoutingKey"`
}

func (r *CreateAMQPQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAMQPQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Queue")
	delete(f, "VHostId")
	delete(f, "AutoDelete")
	delete(f, "ClusterId")
	delete(f, "Remark")
	delete(f, "DeadLetterExchange")
	delete(f, "DeadLetterRoutingKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAMQPQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAMQPQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAMQPQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAMQPQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAMQPRouteRelationRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 交换机所在的vhost
	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`

	// 源Exchange名称
	SourceExchange *string `json:"SourceExchange,omitempty" name:"SourceExchange"`

	// 目标类型:Queue|Exchange
	DestType *string `json:"DestType,omitempty" name:"DestType"`

	// 目标值
	DestValue *string `json:"DestValue,omitempty" name:"DestValue"`

	// 交换机说明，最大128个字符
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 绑定key,缺省值为default
	RoutingKey *string `json:"RoutingKey,omitempty" name:"RoutingKey"`
}

func (r *CreateAMQPRouteRelationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAMQPRouteRelationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "VHostId")
	delete(f, "SourceExchange")
	delete(f, "DestType")
	delete(f, "DestValue")
	delete(f, "Remark")
	delete(f, "RoutingKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAMQPRouteRelationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAMQPRouteRelationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAMQPRouteRelationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAMQPRouteRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAMQPVHostRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// vhost名称，3-64个字符，只能包含字母、数字、“-”及“_”
	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`

	// 未消费消息的保留时间，以毫秒为单位，60秒-15天
	MsgTtl *uint64 `json:"MsgTtl,omitempty" name:"MsgTtl"`

	// 说明，最大128个字符
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateAMQPVHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAMQPVHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "VHostId")
	delete(f, "MsgTtl")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAMQPVHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAMQPVHostResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAMQPVHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAMQPVHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest

	// 集群名称，不支持中字以及除了短线和下划线外的特殊字符且不超过16个字符。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 用户专享物理集群ID，如果不传，则默认在公共集群上创建用户集群资源。
	BindClusterId *uint64 `json:"BindClusterId,omitempty" name:"BindClusterId"`

	// 说明，128个字符以内。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 集群的标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 是否开启公网访问，不填时默认开启
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitempty" name:"PublicAccessEnabled"`
}

func (r *CreateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterName")
	delete(f, "BindClusterId")
	delete(f, "Remark")
	delete(f, "Tags")
	delete(f, "PublicAccessEnabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCmqQueueRequest struct {
	*tchttp.BaseRequest

	// 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 最大堆积消息数。取值范围在公测期间为 1,000,000 - 10,000,000，正式上线后范围可达到 1000,000-1000,000,000。默认取值在公测期间为 10,000,000，正式上线后为 100,000,000。
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitempty" name:"MaxMsgHeapNum"`

	// 消息接收长轮询等待时间。取值范围 0-30 秒，默认值 0。
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitempty" name:"PollingWaitSeconds"`

	// 消息可见性超时。取值范围 1-43200 秒（即12小时内），默认值 30。
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitempty" name:"VisibilityTimeout"`

	// 消息最大长度。取值范围 1024-65536 Byte（即1-64K），默认值 65536。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// 消息保留周期。取值范围 60-1296000 秒（1min-15天），默认值 345600 (4 天)。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// 队列是否开启回溯消息能力，该参数取值范围0-msgRetentionSeconds,即最大的回溯时间为消息在队列中的保留周期，0表示不开启。
	RewindSeconds *uint64 `json:"RewindSeconds,omitempty" name:"RewindSeconds"`

	// 1 表示事务队列，0 表示普通队列
	Transaction *uint64 `json:"Transaction,omitempty" name:"Transaction"`

	// 第一次回查间隔
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitempty" name:"FirstQueryInterval"`

	// 最大回查次数
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitempty" name:"MaxQueryCount"`

	// 死信队列名称
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitempty" name:"DeadLetterQueueName"`

	// 死信策略。0为消息被多次消费未删除，1为Time-To-Live过期
	Policy *uint64 `json:"Policy,omitempty" name:"Policy"`

	// 最大接收次数 1-1000
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitempty" name:"MaxReceiveCount"`

	// policy为1时必选。最大未消费过期时间。范围300-43200，单位秒，需要小于消息最大保留时间msgRetentionSeconds
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitempty" name:"MaxTimeToLive"`

	// 是否开启消息轨迹追踪，当不设置字段时，默认为不开启，该字段为true表示开启，为false表示不开启
	Trace *bool `json:"Trace,omitempty" name:"Trace"`
}

func (r *CreateCmqQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCmqQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	delete(f, "MaxMsgHeapNum")
	delete(f, "PollingWaitSeconds")
	delete(f, "VisibilityTimeout")
	delete(f, "MaxMsgSize")
	delete(f, "MsgRetentionSeconds")
	delete(f, "RewindSeconds")
	delete(f, "Transaction")
	delete(f, "FirstQueryInterval")
	delete(f, "MaxQueryCount")
	delete(f, "DeadLetterQueueName")
	delete(f, "Policy")
	delete(f, "MaxReceiveCount")
	delete(f, "MaxTimeToLive")
	delete(f, "Trace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCmqQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCmqQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建成功的queueId
		QueueId *string `json:"QueueId,omitempty" name:"QueueId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCmqQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCmqQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCmqSubscribeRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 订阅名字，在单个地域同一帐号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`

	// 订阅的协议，目前支持两种协议：http、queue。使用http协议，用户需自己搭建接受消息的web server。使用queue，消息会自动推送到CMQ queue，用户可以并发地拉取消息。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 接收通知的Endpoint，根据协议Protocol区分：对于http，Endpoint必须以“`http://`”开头，host可以是域名或IP；对于Queue，则填QueueName。 请注意，目前推送服务不能推送到私有网络中，因此Endpoint填写为私有网络域名或地址将接收不到推送的消息，目前支持推送到公网和基础网络。
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// 向Endpoint推送消息出现错误时，CMQ推送服务器的重试策略。取值有：1）BACKOFF_RETRY，退避重试。每隔一定时间重试一次，重试够一定次数后，就把该消息丢弃，继续推送下一条消息；2）EXPONENTIAL_DECAY_RETRY，指数衰退重试。每次重试的间隔是指数递增的，例如开始1s，后面是2s，4s，8s...由于Topic消息的周期是一天，所以最多重试一天就把消息丢弃。默认值是EXPONENTIAL_DECAY_RETRY。
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" name:"NotifyStrategy"`

	// 消息正文。消息标签（用于消息过滤)。标签数量不能超过5个，每个标签不超过16个字符。与(Batch)PublishMessage的MsgTag参数配合使用，规则：1）如果FilterTag没有设置，则无论MsgTag是否有设置，订阅接收所有发布到Topic的消息；2）如果FilterTag数组有值，则只有数组中至少有一个值在MsgTag数组中也存在时（即FilterTag和MsgTag有交集），订阅才接收该发布到Topic的消息；3）如果FilterTag数组有值，但MsgTag没设置，则不接收任何发布到Topic的消息，可以认为是2）的一种特例，此时FilterTag和MsgTag没有交集。规则整体的设计思想是以订阅者的意愿为主。
	FilterTag []*string `json:"FilterTag,omitempty" name:"FilterTag"`

	// BindingKey数量不超过5个， 每个BindingKey长度不超过64字节，该字段表示订阅接收消息的过滤策略，每个BindingKey最多含有15个“.”， 即最多16个词组。
	BindingKey []*string `json:"BindingKey,omitempty" name:"BindingKey"`

	// 推送内容的格式。取值：1）JSON；2）SIMPLIFIED，即raw格式。如果Protocol是queue，则取值必须为SIMPLIFIED。如果Protocol是http，两个值均可以，默认值是JSON。
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" name:"NotifyContentFormat"`
}

func (r *CreateCmqSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCmqSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "SubscriptionName")
	delete(f, "Protocol")
	delete(f, "Endpoint")
	delete(f, "NotifyStrategy")
	delete(f, "FilterTag")
	delete(f, "BindingKey")
	delete(f, "NotifyContentFormat")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCmqSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCmqSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订阅id
		SubscriptionId *string `json:"SubscriptionId,omitempty" name:"SubscriptionId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCmqSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCmqSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCmqTopicRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 消息最大长度。取值范围 1024-65536 Byte（即1-64K），默认值 65536。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// 用于指定主题的消息匹配策略。1：表示标签匹配策略；2：表示路由匹配策略，默认值为标签匹配策略。
	FilterType *uint64 `json:"FilterType,omitempty" name:"FilterType"`

	// 消息保存时间。取值范围60 - 86400 s（即1分钟 - 1天），默认值86400。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// 是否开启消息轨迹标识，true表示开启，false表示不开启，不填表示不开启。
	Trace *bool `json:"Trace,omitempty" name:"Trace"`
}

func (r *CreateCmqTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCmqTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "MaxMsgSize")
	delete(f, "FilterType")
	delete(f, "MsgRetentionSeconds")
	delete(f, "Trace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCmqTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCmqTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 主题id
		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCmqTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCmqTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEnvironmentRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称，不支持中字以及除了短线和下划线外的特殊字符且不超过16个字符。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 未消费消息过期时间，单位：秒，最小60，最大1296000，（15天）。
	MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`

	// 说明，128个字符以内。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 消息保留策略
	RetentionPolicy *RetentionPolicy `json:"RetentionPolicy,omitempty" name:"RetentionPolicy"`
}

func (r *CreateEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "MsgTTL")
	delete(f, "Remark")
	delete(f, "ClusterId")
	delete(f, "RetentionPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 环境（命名空间）名称。
		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

		// 未消费消息过期时间，单位：秒。
		MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`

		// 说明，128个字符以内。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 命名空间ID
		NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEnvironmentRoleRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 角色名称。
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 授权项，最多只能包含produce、consume两项的非空字符串数组。
	Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`

	// 必填字段，集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateEnvironmentRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvironmentRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "RoleName")
	delete(f, "Permissions")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEnvironmentRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateEnvironmentRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEnvironmentRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvironmentRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQClusterRequest struct {
	*tchttp.BaseRequest

	// 集群名称，3-64个字符，只能包含字母、数字、“-”及“_”
	Name *string `json:"Name,omitempty" name:"Name"`

	// 集群描述，128个字符以内
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateRocketMQClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRocketMQClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRocketMQClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQGroupRequest struct {
	*tchttp.BaseRequest

	// Group名称，8~64个字符
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 命名空间，目前只支持单个命名空间
	Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces"`

	// 是否开启消费
	ReadEnable *bool `json:"ReadEnable,omitempty" name:"ReadEnable"`

	// 是否开启广播消费
	BroadcastEnable *bool `json:"BroadcastEnable,omitempty" name:"BroadcastEnable"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 说明信息，最长128个字符
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateRocketMQGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Namespaces")
	delete(f, "ReadEnable")
	delete(f, "BroadcastEnable")
	delete(f, "ClusterId")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRocketMQGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRocketMQGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQNamespaceRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间名称，3-64个字符，只能包含字母、数字、“-”及“_”
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 未消费消息的保留时间，以毫秒为单位，60秒-15天
	Ttl *uint64 `json:"Ttl,omitempty" name:"Ttl"`

	// 消息持久化后保留的时间，以毫秒为单位
	RetentionTime *uint64 `json:"RetentionTime,omitempty" name:"RetentionTime"`

	// 说明，最大128个字符
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateRocketMQNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "Ttl")
	delete(f, "RetentionTime")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRocketMQNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRocketMQNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQTopicRequest struct {
	*tchttp.BaseRequest

	// 主题名称，3-64个字符，只能包含字母、数字、“-”及“_”
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 主题所在的命名空间，目前支持在单个命名空间下创建主题
	Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces"`

	// 主题类型，可选值为Normal, GlobalOrder, PartitionedOrder
	Type *string `json:"Type,omitempty" name:"Type"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 主题说明，最大128个字符
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 分区数，全局顺序无效
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`
}

func (r *CreateRocketMQTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Topic")
	delete(f, "Namespaces")
	delete(f, "Type")
	delete(f, "ClusterId")
	delete(f, "Remark")
	delete(f, "PartitionNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRocketMQTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRocketMQTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRoleRequest struct {
	*tchttp.BaseRequest

	// 角色名称，不支持中字以及除了短线和下划线外的特殊字符且长度必须大于0且小等于32。
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 备注说明，长度必须大等于0且小等于128。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleName")
	delete(f, "Remark")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 角色名称
		RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

		// 角色token
		Token *string `json:"Token,omitempty" name:"Token"`

		// 备注说明
	// 注意：此字段可能返回 null，表示取不到有效值。
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubscriptionRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 订阅者名称，不支持中字以及除了短线和下划线外的特殊字符且不超过150个字符。
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`

	// 是否幂等创建，若否不允许创建同名的订阅关系。
	IsIdempotent *bool `json:"IsIdempotent,omitempty" name:"IsIdempotent"`

	// 备注，128个字符以内。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 是否自动创建死信和重试主题，True 表示创建，False表示不创建，默认自动创建死信和重试主题。
	AutoCreatePolicyTopic *bool `json:"AutoCreatePolicyTopic,omitempty" name:"AutoCreatePolicyTopic"`

	// 指定死信和重试主题名称规范，LEGACY表示历史命名规则，COMMUNITY表示Pulsar社区命名规范
	PostFixPattern *string `json:"PostFixPattern,omitempty" name:"PostFixPattern"`
}

func (r *CreateSubscriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubscriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "SubscriptionName")
	delete(f, "IsIdempotent")
	delete(f, "Remark")
	delete(f, "ClusterId")
	delete(f, "AutoCreatePolicyTopic")
	delete(f, "PostFixPattern")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubscriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubscriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建结果。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSubscriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名，不支持中字以及除了短线和下划线外的特殊字符且不超过64个字符。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 0：非分区topic，无分区；非0：具体分区topic的分区数，最大不允许超过128。
	Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`

	// 0： 普通消息；
	// 1 ：全局顺序消息；
	// 2 ：局部顺序消息；
	// 3 ：重试队列；
	// 4 ：死信队列。
	TopicType *uint64 `json:"TopicType,omitempty" name:"TopicType"`

	// 备注，128字符以内。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "Partitions")
	delete(f, "TopicType")
	delete(f, "Remark")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 环境（命名空间）名称。
		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

		// 主题名。
		TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

		// 0：非分区topic，无分区；非0：具体分区topic的分区数。
		Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`

		// 备注，128字符以内。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 0： 普通消息；
	// 1 ：全局顺序消息；
	// 2 ：局部顺序消息；
	// 3 ：重试队列；
	// 4 ：死信队列；
	// 5 ：事务消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TopicType *uint64 `json:"TopicType,omitempty" name:"TopicType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAMQPClusterRequest struct {
	*tchttp.BaseRequest

	// 待删除的集群Id。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteAMQPClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAMQPClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAMQPClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAMQPClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAMQPClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAMQPClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAMQPExchangeRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Vhost名称
	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`

	// 交换机名称
	Exchange *string `json:"Exchange,omitempty" name:"Exchange"`
}

func (r *DeleteAMQPExchangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAMQPExchangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "VHostId")
	delete(f, "Exchange")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAMQPExchangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAMQPExchangeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAMQPExchangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAMQPExchangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAMQPQueueRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Vhost名称
	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`

	// 队列名称
	Queue *string `json:"Queue,omitempty" name:"Queue"`
}

func (r *DeleteAMQPQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAMQPQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "VHostId")
	delete(f, "Queue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAMQPQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAMQPQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAMQPQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAMQPQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAMQPRouteRelationRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Vhost名称
	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`

	// 路由关系ID
	RouteRelationId *string `json:"RouteRelationId,omitempty" name:"RouteRelationId"`
}

func (r *DeleteAMQPRouteRelationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAMQPRouteRelationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "VHostId")
	delete(f, "RouteRelationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAMQPRouteRelationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAMQPRouteRelationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAMQPRouteRelationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAMQPRouteRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAMQPVHostRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// vhost名称
	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`
}

func (r *DeleteAMQPVHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAMQPVHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "VHostId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAMQPVHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAMQPVHostResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAMQPVHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAMQPVHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest

	// 集群Id，传入需要删除的集群Id。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群的ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmqQueueRequest struct {
	*tchttp.BaseRequest

	// 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *DeleteCmqQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCmqQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCmqQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmqQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCmqQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCmqQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmqSubscribeRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 订阅名字，在单个地域同一帐号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
}

func (r *DeleteCmqSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCmqSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "SubscriptionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCmqSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmqSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCmqSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCmqSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmqTopicRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *DeleteCmqTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCmqTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCmqTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmqTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCmqTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCmqTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEnvironmentRolesRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 角色名称数组。
	RoleNames []*string `json:"RoleNames,omitempty" name:"RoleNames"`

	// 必填字段，集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteEnvironmentRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEnvironmentRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "RoleNames")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEnvironmentRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEnvironmentRolesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteEnvironmentRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEnvironmentRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEnvironmentsRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）数组，每次最多删除20个。
	EnvironmentIds []*string `json:"EnvironmentIds,omitempty" name:"EnvironmentIds"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEnvironmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentIds")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功删除的环境（命名空间）数组。
		EnvironmentIds []*string `json:"EnvironmentIds,omitempty" name:"EnvironmentIds"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQClusterRequest struct {
	*tchttp.BaseRequest

	// 待删除的集群Id。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteRocketMQClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRocketMQClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRocketMQClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQGroupRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 消费组名称
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteRocketMQGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRocketMQGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRocketMQGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQNamespaceRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *DeleteRocketMQNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRocketMQNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRocketMQNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQTopicRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 主题名称
	Topic *string `json:"Topic,omitempty" name:"Topic"`
}

func (r *DeleteRocketMQTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "Topic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRocketMQTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRocketMQTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRolesRequest struct {
	*tchttp.BaseRequest

	// 角色名称数组。
	RoleNames []*string `json:"RoleNames,omitempty" name:"RoleNames"`

	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleNames")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRolesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功删除的角色名称数组。
		RoleNames []*string `json:"RoleNames,omitempty" name:"RoleNames"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSubscriptionsRequest struct {
	*tchttp.BaseRequest

	// 订阅关系集合，每次最多删除20个。
	SubscriptionTopicSets []*SubscriptionTopic `json:"SubscriptionTopicSets,omitempty" name:"SubscriptionTopicSets"`

	// pulsar集群Id。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 是否强制删除，默认为false
	Force *bool `json:"Force,omitempty" name:"Force"`
}

func (r *DeleteSubscriptionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSubscriptionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscriptionTopicSets")
	delete(f, "ClusterId")
	delete(f, "EnvironmentId")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSubscriptionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSubscriptionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功删除的订阅关系数组。
		SubscriptionTopicSets []*SubscriptionTopic `json:"SubscriptionTopicSets,omitempty" name:"SubscriptionTopicSets"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSubscriptionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSubscriptionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicsRequest struct {
	*tchttp.BaseRequest

	// 主题集合，每次最多删除20个。
	TopicSets []*TopicRecord `json:"TopicSets,omitempty" name:"TopicSets"`

	// pulsar集群Id。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 是否强制删除，默认为false
	Force *bool `json:"Force,omitempty" name:"Force"`
}

func (r *DeleteTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicSets")
	delete(f, "ClusterId")
	delete(f, "EnvironmentId")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 被删除的主题数组。
		TopicSets []*TopicRecord `json:"TopicSets,omitempty" name:"TopicSets"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPClusterRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeAMQPClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAMQPClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAMQPClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群信息
		ClusterInfo *AMQPClusterInfo `json:"ClusterInfo,omitempty" name:"ClusterInfo"`

		// 集群配置
		ClusterConfig *AMQPClusterConfig `json:"ClusterConfig,omitempty" name:"ClusterConfig"`

		// 集群最近使用量
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClusterStats *AMQPClusterRecentStats `json:"ClusterStats,omitempty" name:"ClusterStats"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAMQPClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAMQPClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPClustersRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 按照集群ID关键字搜索
	IdKeyword *string `json:"IdKeyword,omitempty" name:"IdKeyword"`

	// 按照集群名称关键字搜索
	NameKeyword *string `json:"NameKeyword,omitempty" name:"NameKeyword"`

	// 集群ID列表过滤
	ClusterIdList []*string `json:"ClusterIdList,omitempty" name:"ClusterIdList"`

	// 标签过滤查找时，需要设置为true
	IsTagFilter *bool `json:"IsTagFilter,omitempty" name:"IsTagFilter"`

	// 过滤器。目前支持按标签过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeAMQPClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAMQPClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "IdKeyword")
	delete(f, "NameKeyword")
	delete(f, "ClusterIdList")
	delete(f, "IsTagFilter")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAMQPClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPClustersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClusterList []*AMQPClusterDetail `json:"ClusterList,omitempty" name:"ClusterList"`

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAMQPClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAMQPClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPCreateQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAMQPCreateQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAMQPCreateQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAMQPCreateQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPCreateQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 租户总共可使用集群数量
		MaxClusterNum *uint64 `json:"MaxClusterNum,omitempty" name:"MaxClusterNum"`

		// 租户已创建集群数量
		UsedClusterNum *uint64 `json:"UsedClusterNum,omitempty" name:"UsedClusterNum"`

		// Exchange容量
		ExchangeCapacity *uint64 `json:"ExchangeCapacity,omitempty" name:"ExchangeCapacity"`

		// Queue容量
		QueueCapacity *uint64 `json:"QueueCapacity,omitempty" name:"QueueCapacity"`

		// 单Vhost TPS
		MaxTpsPerVHost *uint64 `json:"MaxTpsPerVHost,omitempty" name:"MaxTpsPerVHost"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAMQPCreateQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAMQPCreateQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPExchangesRequest struct {
	*tchttp.BaseRequest

	// 查询偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询限制数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Vhost ID
	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`

	// 按路由类型过滤查询结果，可选择Direct, Fanout, Topic
	FilterType []*string `json:"FilterType,omitempty" name:"FilterType"`

	// 按exchange名称搜索，支持模糊查询
	FilterName *string `json:"FilterName,omitempty" name:"FilterName"`

	// 过滤查询内部或者外部exchange
	FilterInternal *bool `json:"FilterInternal,omitempty" name:"FilterInternal"`
}

func (r *DescribeAMQPExchangesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAMQPExchangesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClusterId")
	delete(f, "VHostId")
	delete(f, "FilterType")
	delete(f, "FilterName")
	delete(f, "FilterInternal")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAMQPExchangesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPExchangesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 主题信息列表
		Exchanges []*AMQPExchange `json:"Exchanges,omitempty" name:"Exchanges"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAMQPExchangesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAMQPExchangesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPQueuesRequest struct {
	*tchttp.BaseRequest

	// 查询偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询限制数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Vhost名称
	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`

	// 按队列名称搜索，支持模糊查询
	NameKeyword *string `json:"NameKeyword,omitempty" name:"NameKeyword"`

	// 查询结果排序规则，ASC为升序，DESC为降序
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`

	// 对查询结果排序，此为排序字段，目前支持Accumulative（消息堆积量）、Tps
	SortedBy *string `json:"SortedBy,omitempty" name:"SortedBy"`

	// 队列名称，指定此参数后将只返回该队列信息
	FilterOneQueue *string `json:"FilterOneQueue,omitempty" name:"FilterOneQueue"`
}

func (r *DescribeAMQPQueuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAMQPQueuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClusterId")
	delete(f, "VHostId")
	delete(f, "NameKeyword")
	delete(f, "SortOrder")
	delete(f, "SortedBy")
	delete(f, "FilterOneQueue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAMQPQueuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPQueuesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 队列信息列表
		Queues []*AMQPQueueDetail `json:"Queues,omitempty" name:"Queues"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAMQPQueuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAMQPQueuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPRouteRelationsRequest struct {
	*tchttp.BaseRequest

	// 查询偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询限制数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Vhost名称
	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`

	// 按源exchange名称过滤查询结果，支持模糊查询
	FilterSourceExchange *string `json:"FilterSourceExchange,omitempty" name:"FilterSourceExchange"`

	// 按绑定的目标类型过滤查询结果，可选值:Exchange、Queue
	FilterDestType *string `json:"FilterDestType,omitempty" name:"FilterDestType"`

	// 按目标名称过滤查询结果，支持模糊查询
	FilterDestValue *string `json:"FilterDestValue,omitempty" name:"FilterDestValue"`
}

func (r *DescribeAMQPRouteRelationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAMQPRouteRelationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClusterId")
	delete(f, "VHostId")
	delete(f, "FilterSourceExchange")
	delete(f, "FilterDestType")
	delete(f, "FilterDestValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAMQPRouteRelationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPRouteRelationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 路由关系列表
		RouteRelations []*AMQPRouteRelation `json:"RouteRelations,omitempty" name:"RouteRelations"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAMQPRouteRelationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAMQPRouteRelationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPVHostsRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 按名称搜索
	NameKeyword *string `json:"NameKeyword,omitempty" name:"NameKeyword"`
}

func (r *DescribeAMQPVHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAMQPVHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NameKeyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAMQPVHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPVHostsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Vhost 列表
		VHosts []*AMQPVHost `json:"VHosts,omitempty" name:"VHosts"`

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAMQPVHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAMQPVHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllTenantsRequest struct {
	*tchttp.BaseRequest

	// 查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 物理集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 虚拟集群ID
	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`

	// 虚拟集群名称
	TenantName *string `json:"TenantName,omitempty" name:"TenantName"`

	// 协议类型数组
	Types []*string `json:"Types,omitempty" name:"Types"`

	// 排序字段名，支持createTime，updateTime
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 升序排列ASC，降序排列DESC
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

func (r *DescribeAllTenantsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllTenantsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClusterName")
	delete(f, "TenantId")
	delete(f, "TenantName")
	delete(f, "Types")
	delete(f, "SortBy")
	delete(f, "SortOrder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllTenantsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllTenantsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总条数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 虚拟集群列表
		Tenants []*InternalTenant `json:"Tenants,omitempty" name:"Tenants"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllTenantsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllTenantsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBindClustersRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBindClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBindClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBindClustersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 专享集群的数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 专享集群的列表
		ClusterSet []*BindCluster `json:"ClusterSet,omitempty" name:"ClusterSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBindClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBindVpcsRequest struct {
	*tchttp.BaseRequest

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeBindVpcsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindVpcsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBindVpcsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBindVpcsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Vpc集合。
		VpcSets []*VpcBindRecord `json:"VpcSets,omitempty" name:"VpcSets"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBindVpcsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindVpcsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterDetailRequest struct {
	*tchttp.BaseRequest

	// 集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群的详细信息
		ClusterSet *Cluster `json:"ClusterSet,omitempty" name:"ClusterSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 集群ID列表过滤
	ClusterIdList []*string `json:"ClusterIdList,omitempty" name:"ClusterIdList"`

	// 是否标签过滤
	IsTagFilter *bool `json:"IsTagFilter,omitempty" name:"IsTagFilter"`

	// 过滤器。目前支持按标签过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClusterIdList")
	delete(f, "IsTagFilter")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群列表数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 集群信息列表
		ClusterSet []*Cluster `json:"ClusterSet,omitempty" name:"ClusterSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqDeadLetterSourceQueuesRequest struct {
	*tchttp.BaseRequest

	// 死信队列名称
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitempty" name:"DeadLetterQueueName"`

	// 分页时本页获取主题列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页时本页获取主题的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 根据SourceQueueName过滤
	SourceQueueName *string `json:"SourceQueueName,omitempty" name:"SourceQueueName"`
}

func (r *DescribeCmqDeadLetterSourceQueuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqDeadLetterSourceQueuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeadLetterQueueName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SourceQueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCmqDeadLetterSourceQueuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqDeadLetterSourceQueuesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 满足本次条件的队列个数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 死信队列源队列
		QueueSet []*CmqDeadLetterSource `json:"QueueSet,omitempty" name:"QueueSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCmqDeadLetterSourceQueuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqDeadLetterSourceQueuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqQueueDetailRequest struct {
	*tchttp.BaseRequest

	// 精确匹配QueueName
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *DescribeCmqQueueDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqQueueDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCmqQueueDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqQueueDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 队列详情列表。
		QueueDescribe *CmqQueue `json:"QueueDescribe,omitempty" name:"QueueDescribe"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCmqQueueDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqQueueDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqQueuesRequest struct {
	*tchttp.BaseRequest

	// 分页时本页获取队列列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页时本页获取队列的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 根据QueueName进行过滤
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// CMQ 队列名称列表过滤
	QueueNameList []*string `json:"QueueNameList,omitempty" name:"QueueNameList"`

	// 标签过滤查找时，需要设置为 true
	IsTagFilter *bool `json:"IsTagFilter,omitempty" name:"IsTagFilter"`

	// 过滤器。目前支持按标签过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCmqQueuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqQueuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "QueueName")
	delete(f, "QueueNameList")
	delete(f, "IsTagFilter")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCmqQueuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqQueuesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 队列列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		QueueList []*CmqQueue `json:"QueueList,omitempty" name:"QueueList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCmqQueuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqQueuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqSubscriptionDetailRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 分页时本页获取主题列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页时本页获取主题的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 根据SubscriptionName进行模糊搜索
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
}

func (r *DescribeCmqSubscriptionDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqSubscriptionDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SubscriptionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCmqSubscriptionDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqSubscriptionDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Subscription属性集合
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubscriptionSet []*CmqSubscription `json:"SubscriptionSet,omitempty" name:"SubscriptionSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCmqSubscriptionDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqSubscriptionDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqTopicDetailRequest struct {
	*tchttp.BaseRequest

	// 精确匹配TopicName。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *DescribeCmqTopicDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqTopicDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCmqTopicDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqTopicDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 主题详情
		TopicDescribe *CmqTopic `json:"TopicDescribe,omitempty" name:"TopicDescribe"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCmqTopicDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqTopicDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqTopicsRequest struct {
	*tchttp.BaseRequest

	// 分页时本页获取队列列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页时本页获取队列的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 根据TopicName进行模糊搜索
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// CMQ 主题名称列表过滤
	TopicNameList []*string `json:"TopicNameList,omitempty" name:"TopicNameList"`

	// 标签过滤查找时，需要设置为 true
	IsTagFilter *bool `json:"IsTagFilter,omitempty" name:"IsTagFilter"`

	// 过滤器。目前支持按标签过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeCmqTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TopicName")
	delete(f, "TopicNameList")
	delete(f, "IsTagFilter")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCmqTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqTopicsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 主题列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		TopicList []*CmqTopic `json:"TopicList,omitempty" name:"TopicList"`

		// 全量主题数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCmqTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvironmentAttributesRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeEnvironmentAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvironmentAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 未消费消息过期时间，单位：秒，最大1296000（15天）。
		MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`

		// 消费速率限制，单位：byte/秒，0：不限速。
		RateInByte *uint64 `json:"RateInByte,omitempty" name:"RateInByte"`

		// 消费速率限制，单位：个数/秒，0：不限速。
		RateInSize *uint64 `json:"RateInSize,omitempty" name:"RateInSize"`

		// 已消费消息保存策略，单位：小时，0：消费完马上删除。
		RetentionHours *uint64 `json:"RetentionHours,omitempty" name:"RetentionHours"`

		// 已消费消息保存策略，单位：G，0：消费完马上删除。
		RetentionSize *uint64 `json:"RetentionSize,omitempty" name:"RetentionSize"`

		// 环境（命名空间）名称。
		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

		// 副本数。
		Replicas *uint64 `json:"Replicas,omitempty" name:"Replicas"`

		// 备注。
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEnvironmentAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvironmentRolesRequest struct {
	*tchttp.BaseRequest

	// 必填字段，环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 起始下标，不填默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 必填字段，Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 角色名称
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// * RoleName
	// 按照角色名进行过滤，精确查询。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeEnvironmentRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClusterId")
	delete(f, "RoleName")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvironmentRolesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 命名空间角色集合。
		EnvironmentRoleSets []*EnvironmentRole `json:"EnvironmentRoleSets,omitempty" name:"EnvironmentRoleSets"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEnvironmentRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvironmentsRequest struct {
	*tchttp.BaseRequest

	// 命名空间名称，模糊搜索。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// * EnvironmentId
	// 按照名称空间进行过滤，精确查询。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClusterId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 命名空间记录数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 命名空间集合数组。
		EnvironmentSet []*Environment `json:"EnvironmentSet,omitempty" name:"EnvironmentSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespaceBundlesOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 虚拟集群（租户）ID
	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`

	// 命名空间名
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// 是否需要监控指标，若传false，则不需要传Limit和Offset分页参数
	NeedMetrics *bool `json:"NeedMetrics,omitempty" name:"NeedMetrics"`

	// 查询限制条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeNamespaceBundlesOptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespaceBundlesOptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterName")
	delete(f, "TenantId")
	delete(f, "NamespaceName")
	delete(f, "NeedMetrics")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNamespaceBundlesOptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespaceBundlesOptResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录条数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// bundle列表
		BundleSet []*BundleSetOpt `json:"BundleSet,omitempty" name:"BundleSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNamespaceBundlesOptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNamespaceBundlesOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeHealthOptRequest struct {
	*tchttp.BaseRequest

	// 节点实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeNodeHealthOptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeHealthOptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNodeHealthOptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeHealthOptResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0-异常；1-正常
		NodeState *int64 `json:"NodeState,omitempty" name:"NodeState"`

		// 最近一次健康检查的时间
		LatestHealthCheckTime *string `json:"LatestHealthCheckTime,omitempty" name:"LatestHealthCheckTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNodeHealthOptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeHealthOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProducersRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 生产者名称，模糊匹配。
	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeProducersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProducersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProducerName")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProducersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProducersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 生产者集合数组。
		ProducerSets []*Producer `json:"ProducerSets,omitempty" name:"ProducerSets"`

		// 记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProducersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProducersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublisherSummaryRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间名称
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 主题名称
	Topic *string `json:"Topic,omitempty" name:"Topic"`
}

func (r *DescribePublisherSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublisherSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Namespace")
	delete(f, "Topic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublisherSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublisherSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 生产速率（条/秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
		MsgRateIn *float64 `json:"MsgRateIn,omitempty" name:"MsgRateIn"`

		// 生产速率（字节/秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
		MsgThroughputIn *float64 `json:"MsgThroughputIn,omitempty" name:"MsgThroughputIn"`

		// 生产者数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		PublisherCount *int64 `json:"PublisherCount,omitempty" name:"PublisherCount"`

		// 消息存储大小，以字节为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
		StorageSize *int64 `json:"StorageSize,omitempty" name:"StorageSize"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublisherSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublisherSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublishersRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间名称
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 主题名称
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 参数过滤器，支持ProducerName，Address字段
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 查询偏移量，默认为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询条数，默认为20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序器
	Sort *Sort `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribePublishersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublishersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Namespace")
	delete(f, "Topic")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublishersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublishersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总条数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 生产者信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Publishers []*Publisher `json:"Publishers,omitempty" name:"Publishers"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublishersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublishersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQClusterRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeRocketMQClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群信息
		ClusterInfo *RocketMQClusterInfo `json:"ClusterInfo,omitempty" name:"ClusterInfo"`

		// 集群配置
		ClusterConfig *RocketMQClusterConfig `json:"ClusterConfig,omitempty" name:"ClusterConfig"`

		// 集群最近使用量
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClusterStats *RocketMQClusterRecentStats `json:"ClusterStats,omitempty" name:"ClusterStats"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQClustersRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 按照集群ID关键字搜索
	IdKeyword *string `json:"IdKeyword,omitempty" name:"IdKeyword"`

	// 按照集群名称关键字搜索
	NameKeyword *string `json:"NameKeyword,omitempty" name:"NameKeyword"`

	// 集群ID列表过滤
	ClusterIdList []*string `json:"ClusterIdList,omitempty" name:"ClusterIdList"`

	// 标签过滤查找时，需要设置为true
	IsTagFilter *bool `json:"IsTagFilter,omitempty" name:"IsTagFilter"`

	// 过滤器。目前支持标签过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRocketMQClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "IdKeyword")
	delete(f, "NameKeyword")
	delete(f, "ClusterIdList")
	delete(f, "IsTagFilter")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQClustersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 集群信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClusterList []*RocketMQClusterDetail `json:"ClusterList,omitempty" name:"ClusterList"`

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQGroupsRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 主题名称，输入此参数可查询该主题下所有的订阅组
	FilterTopic *string `json:"FilterTopic,omitempty" name:"FilterTopic"`

	// 按消费组名称查询消费组，支持模糊查询
	FilterGroup *string `json:"FilterGroup,omitempty" name:"FilterGroup"`

	// 按照指定字段排序，可选值为tps，accumulative
	SortedBy *string `json:"SortedBy,omitempty" name:"SortedBy"`

	// 按升序或降序排列，可选值为asc，desc
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`

	// 订阅组名称，指定此参数后将只返回该订阅组信息
	FilterOneGroup *string `json:"FilterOneGroup,omitempty" name:"FilterOneGroup"`
}

func (r *DescribeRocketMQGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterTopic")
	delete(f, "FilterGroup")
	delete(f, "SortedBy")
	delete(f, "SortOrder")
	delete(f, "FilterOneGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 订阅组列表
		Groups []*RocketMQGroup `json:"Groups,omitempty" name:"Groups"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQNamespacesRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 按名称搜索
	NameKeyword *string `json:"NameKeyword,omitempty" name:"NameKeyword"`
}

func (r *DescribeRocketMQNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQNamespacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NameKeyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQNamespacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 命名空间列表
		Namespaces []*RocketMQNamespace `json:"Namespaces,omitempty" name:"Namespaces"`

		// 总条数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQTopicsRequest struct {
	*tchttp.BaseRequest

	// 查询偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询限制数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 按主题类型过滤查询结果，可选择Normal, GlobalOrder, PartitionedOrder, Transaction
	FilterType []*string `json:"FilterType,omitempty" name:"FilterType"`

	// 按主题名称搜索，支持模糊查询
	FilterName *string `json:"FilterName,omitempty" name:"FilterName"`
}

func (r *DescribeRocketMQTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "FilterType")
	delete(f, "FilterName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQTopicsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总记录数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 主题信息列表
		Topics []*RocketMQTopic `json:"Topics,omitempty" name:"Topics"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRolesRequest struct {
	*tchttp.BaseRequest

	// 角色名称，模糊查询
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 起始下标，不填默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// * RoleName
	// 按照角色名进行过滤，精确查询。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClusterId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRolesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 记录数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 角色数组。
		RoleSets []*Role `json:"RoleSets,omitempty" name:"RoleSets"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionsRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 订阅者名称，模糊匹配。
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`

	// 数据过滤条件。
	Filters []*FilterSubscription `json:"Filters,omitempty" name:"Filters"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeSubscriptionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscriptionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SubscriptionName")
	delete(f, "Filters")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubscriptionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订阅者集合数组。
		SubscriptionSets []*Subscription `json:"SubscriptionSets,omitempty" name:"SubscriptionSets"`

		// 数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscriptionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscriptionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicsRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名模糊匹配。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// topic类型描述：
	// 0：普通消息；
	// 1：全局顺序消息；
	// 2：局部顺序消息；
	// 3：重试队列；
	// 4：死信队列；
	// 5：事务消息。
	TopicType *uint64 `json:"TopicType,omitempty" name:"TopicType"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// * TopicName
	// 按照主题名字查询，精确查询。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TopicType")
	delete(f, "ClusterId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 主题集合数组。
		TopicSets []*Topic `json:"TopicSets,omitempty" name:"TopicSets"`

		// 主题数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Environment struct {

	// 命名空间名称
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 说明
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 未消费消息过期时间，单位：秒，最大1296000（15天）
	MsgTTL *int64 `json:"MsgTTL,omitempty" name:"MsgTTL"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最近修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// Topic数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicNum *int64 `json:"TopicNum,omitempty" name:"TopicNum"`

	// 消息保留策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetentionPolicy *RetentionPolicy `json:"RetentionPolicy,omitempty" name:"RetentionPolicy"`
}

type EnvironmentRole struct {

	// 环境（命名空间）。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 角色名称。
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 授权项，最多只能包含produce、consume两项的非空字符串数组。
	Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`

	// 角色描述。
	RoleDescribe *string `json:"RoleDescribe,omitempty" name:"RoleDescribe"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type Filter struct {

	// 过滤参数的名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数值
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type FilterSubscription struct {

	// 是否仅展示包含真实消费者的订阅。
	ConsumerHasCount *bool `json:"ConsumerHasCount,omitempty" name:"ConsumerHasCount"`

	// 是否仅展示消息堆积的订阅。
	ConsumerHasBacklog *bool `json:"ConsumerHasBacklog,omitempty" name:"ConsumerHasBacklog"`

	// 是否仅展示存在消息超期丢弃的订阅。
	ConsumerHasExpired *bool `json:"ConsumerHasExpired,omitempty" name:"ConsumerHasExpired"`

	// 按照订阅名过滤，精确查询。
	SubscriptionNames []*string `json:"SubscriptionNames,omitempty" name:"SubscriptionNames"`
}

type InternalTenant struct {

	// 虚拟集群ID
	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`

	// 虚拟集群名称
	TenantName *string `json:"TenantName,omitempty" name:"TenantName"`

	// 客户UIN
	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`

	// 客户的APPID
	CustomerAppId *string `json:"CustomerAppId,omitempty" name:"CustomerAppId"`

	// 物理集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群协议类型，支持的值为TDMQ，ROCKETMQ，AMQP，CMQ
	Type *string `json:"Type,omitempty" name:"Type"`

	// 命名空间配额
	MaxNamespaces *int64 `json:"MaxNamespaces,omitempty" name:"MaxNamespaces"`

	// 已使用命名空间配额
	UsedNamespaces *int64 `json:"UsedNamespaces,omitempty" name:"UsedNamespaces"`

	// Topic配额
	MaxTopics *int64 `json:"MaxTopics,omitempty" name:"MaxTopics"`

	// 已使用Topic配额
	UsedTopics *int64 `json:"UsedTopics,omitempty" name:"UsedTopics"`

	// Topic分区数配额
	MaxPartitions *int64 `json:"MaxPartitions,omitempty" name:"MaxPartitions"`

	// 已使用Topic分区数配额
	UsedPartitions *int64 `json:"UsedPartitions,omitempty" name:"UsedPartitions"`

	// 存储配额, byte为单位
	MaxMsgBacklogSize *uint64 `json:"MaxMsgBacklogSize,omitempty" name:"MaxMsgBacklogSize"`

	// 命名空间最大生产TPS
	MaxPublishTps *uint64 `json:"MaxPublishTps,omitempty" name:"MaxPublishTps"`

	// 消息最大保留时间，秒为单位
	MaxRetention *uint64 `json:"MaxRetention,omitempty" name:"MaxRetention"`

	// 创建时间，毫秒为单位
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间，毫秒为单位
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 命名空间最大消费TPS
	MaxDispatchTps *uint64 `json:"MaxDispatchTps,omitempty" name:"MaxDispatchTps"`

	// 命名空间最大消费带宽，byte为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDispatchRateInBytes *uint64 `json:"MaxDispatchRateInBytes,omitempty" name:"MaxDispatchRateInBytes"`

	// 命名空间最大生产带宽，byte为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxPublishRateInBytes *uint64 `json:"MaxPublishRateInBytes,omitempty" name:"MaxPublishRateInBytes"`

	// 消息最大保留空间，MB为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRetentionSizeInMB *uint64 `json:"MaxRetentionSizeInMB,omitempty" name:"MaxRetentionSizeInMB"`
}

type ModifyAMQPClusterRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 3-64个字符，只能包含字母、数字、“-”及“_”
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 说明信息，不超过128个字符
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyAMQPClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAMQPClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterName")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAMQPClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAMQPClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAMQPClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAMQPClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAMQPExchangeRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Vhost间名称
	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`

	// 交换机名称
	Exchange *string `json:"Exchange,omitempty" name:"Exchange"`

	// 说明信息，最大128个字符
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyAMQPExchangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAMQPExchangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "VHostId")
	delete(f, "Exchange")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAMQPExchangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAMQPExchangeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAMQPExchangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAMQPExchangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAMQPQueueRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Vhost名称
	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`

	// 队列名称
	Queue *string `json:"Queue,omitempty" name:"Queue"`

	// 是否自动清除
	AutoDelete *bool `json:"AutoDelete,omitempty" name:"AutoDelete"`

	// 说明信息，最大128个字符
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 死信exchange
	DeadLetterExchange *string `json:"DeadLetterExchange,omitempty" name:"DeadLetterExchange"`

	// 路由键
	DeadLetterRoutingKey *string `json:"DeadLetterRoutingKey,omitempty" name:"DeadLetterRoutingKey"`
}

func (r *ModifyAMQPQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAMQPQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "VHostId")
	delete(f, "Queue")
	delete(f, "AutoDelete")
	delete(f, "Remark")
	delete(f, "DeadLetterExchange")
	delete(f, "DeadLetterRoutingKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAMQPQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAMQPQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAMQPQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAMQPQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAMQPVHostRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// vhost名称，3-64个字符，只能包含字母、数字、“-”及“_”
	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`

	// 未消费消息的保留时间，以毫秒为单位，60秒-15天
	MsgTtl *uint64 `json:"MsgTtl,omitempty" name:"MsgTtl"`

	// 说明，最大128个字符
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyAMQPVHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAMQPVHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "VHostId")
	delete(f, "MsgTtl")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAMQPVHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAMQPVHostResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAMQPVHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAMQPVHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterRequest struct {
	*tchttp.BaseRequest

	// Pulsar 集群的ID，需要更新的集群Id。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 更新后的集群名称。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 说明信息。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 开启公网访问，只能为true
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitempty" name:"PublicAccessEnabled"`
}

func (r *ModifyClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterName")
	delete(f, "Remark")
	delete(f, "PublicAccessEnabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Pulsar 集群的ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmqQueueAttributeRequest struct {
	*tchttp.BaseRequest

	// 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 最大堆积消息数。取值范围在公测期间为 1,000,000 - 10,000,000，正式上线后范围可达到 1000,000-1000,000,000。默认取值在公测期间为 10,000,000，正式上线后为 100,000,000。
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitempty" name:"MaxMsgHeapNum"`

	// 消息接收长轮询等待时间。取值范围 0-30 秒，默认值 0。
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitempty" name:"PollingWaitSeconds"`

	// 消息可见性超时。取值范围 1-43200 秒（即12小时内），默认值 30。
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitempty" name:"VisibilityTimeout"`

	// 消息最大长度。取值范围 1024-65536 Byte（即1-64K），默认值 65536。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// 消息保留周期。取值范围 60-1296000 秒（1min-15天），默认值 345600 (4 天)。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// 消息最长回溯时间，取值范围0-msgRetentionSeconds，消息的最大回溯之间为消息在队列中的保存周期，0表示不开启消息回溯。
	RewindSeconds *uint64 `json:"RewindSeconds,omitempty" name:"RewindSeconds"`

	// 第一次查询时间
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitempty" name:"FirstQueryInterval"`

	// 最大查询次数
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitempty" name:"MaxQueryCount"`

	// 死信队列名称
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitempty" name:"DeadLetterQueueName"`

	// MaxTimeToLivepolicy为1时必选。最大未消费过期时间。范围300-43200，单位秒，需要小于消息最大保留时间MsgRetentionSeconds
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitempty" name:"MaxTimeToLive"`

	// 最大接收次数
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitempty" name:"MaxReceiveCount"`

	// 死信队列策略
	Policy *uint64 `json:"Policy,omitempty" name:"Policy"`

	// 是否开启消息轨迹标识，true表示开启，false表示不开启，不填表示不开启。
	Trace *bool `json:"Trace,omitempty" name:"Trace"`

	// 是否开启事务，1开启，0不开启
	Transaction *uint64 `json:"Transaction,omitempty" name:"Transaction"`
}

func (r *ModifyCmqQueueAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCmqQueueAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	delete(f, "MaxMsgHeapNum")
	delete(f, "PollingWaitSeconds")
	delete(f, "VisibilityTimeout")
	delete(f, "MaxMsgSize")
	delete(f, "MsgRetentionSeconds")
	delete(f, "RewindSeconds")
	delete(f, "FirstQueryInterval")
	delete(f, "MaxQueryCount")
	delete(f, "DeadLetterQueueName")
	delete(f, "MaxTimeToLive")
	delete(f, "MaxReceiveCount")
	delete(f, "Policy")
	delete(f, "Trace")
	delete(f, "Transaction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCmqQueueAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmqQueueAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCmqQueueAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCmqQueueAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmqSubscriptionAttributeRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 订阅名字，在单个地域同一帐号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`

	// 向 Endpoint 推送消息出现错误时，CMQ 推送服务器的重试策略。取值如下：
	// （1）BACKOFF_RETRY，退避重试。每隔一定时间重试一次，重试够一定次数后，就把该消息丢弃，继续推送下一条消息。
	// （2）EXPONENTIAL_DECAY_RETRY，指数衰退重试。每次重试的间隔是指数递增的，例如开始1s，后面是2s，4s，8s···由于 Topic 消息的周期是一天，所以最多重试一天就把消息丢弃。默认值是 EXPONENTIAL_DECAY_RETRY。
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" name:"NotifyStrategy"`

	// 推送内容的格式。取值：（1）JSON；（2）SIMPLIFIED，即 raw 格式。如果 Protocol 是 queue，则取值必须为 SIMPLIFIED。如果 Protocol 是 HTTP，两个值均可以，默认值是 JSON。
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" name:"NotifyContentFormat"`

	// 消息正文。消息标签（用于消息过滤)。标签数量不能超过5个，每个标签不超过16个字符。与(Batch)PublishMessage的MsgTag参数配合使用，规则：1）如果FilterTag没有设置，则无论MsgTag是否有设置，订阅接收所有发布到Topic的消息；2）如果FilterTag数组有值，则只有数组中至少有一个值在MsgTag数组中也存在时（即FilterTag和MsgTag有交集），订阅才接收该发布到Topic的消息；3）如果FilterTag数组有值，但MsgTag没设置，则不接收任何发布到Topic的消息，可以认为是2）的一种特例，此时FilterTag和MsgTag没有交集。规则整体的设计思想是以订阅者的意愿为主。
	FilterTags []*string `json:"FilterTags,omitempty" name:"FilterTags"`

	// BindingKey数量不超过5个， 每个BindingKey长度不超过64字节，该字段表示订阅接收消息的过滤策略，每个BindingKey最多含有15个“.”， 即最多16个词组。
	BindingKey []*string `json:"BindingKey,omitempty" name:"BindingKey"`
}

func (r *ModifyCmqSubscriptionAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCmqSubscriptionAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "SubscriptionName")
	delete(f, "NotifyStrategy")
	delete(f, "NotifyContentFormat")
	delete(f, "FilterTags")
	delete(f, "BindingKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCmqSubscriptionAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmqSubscriptionAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCmqSubscriptionAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCmqSubscriptionAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmqTopicAttributeRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 消息最大长度。取值范围1024 - 65536 Byte（即1 - 64K），默认值65536。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// 消息保存时间。取值范围60 - 86400 s（即1分钟 - 1天），默认值86400。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// 是否开启消息轨迹标识，true表示开启，false表示不开启，不填表示不开启。
	Trace *bool `json:"Trace,omitempty" name:"Trace"`
}

func (r *ModifyCmqTopicAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCmqTopicAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "MaxMsgSize")
	delete(f, "MsgRetentionSeconds")
	delete(f, "Trace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCmqTopicAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmqTopicAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCmqTopicAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCmqTopicAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvironmentAttributesRequest struct {
	*tchttp.BaseRequest

	// 命名空间名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 未消费消息过期时间，单位：秒，最大1296000。
	MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`

	// 备注，字符串最长不超过128。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 消息保留策略
	RetentionPolicy *RetentionPolicy `json:"RetentionPolicy,omitempty" name:"RetentionPolicy"`
}

func (r *ModifyEnvironmentAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvironmentAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "MsgTTL")
	delete(f, "Remark")
	delete(f, "ClusterId")
	delete(f, "RetentionPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEnvironmentAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvironmentAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 命名空间名称。
		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

		// 未消费消息过期时间，单位：秒。
		MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`

		// 备注，字符串最长不超过128。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEnvironmentAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvironmentAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvironmentRoleRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 角色名称。
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 授权项，最多只能包含produce、consume两项的非空字符串数组。
	Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`

	// 必填字段，集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ModifyEnvironmentRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvironmentRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "RoleName")
	delete(f, "Permissions")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEnvironmentRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvironmentRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEnvironmentRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvironmentRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQClusterRequest struct {
	*tchttp.BaseRequest

	// RocketMQ集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 3-64个字符，只能包含字母、数字、“-”及“_”
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 说明信息，不超过128个字符
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyRocketMQClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterName")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRocketMQClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRocketMQClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQGroupRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 消费组名称
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 说明信息，最长128个字符
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 是否开启消费
	ReadEnable *bool `json:"ReadEnable,omitempty" name:"ReadEnable"`

	// 是否开启广播消费
	BroadcastEnable *bool `json:"BroadcastEnable,omitempty" name:"BroadcastEnable"`
}

func (r *ModifyRocketMQGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "GroupId")
	delete(f, "Remark")
	delete(f, "ReadEnable")
	delete(f, "BroadcastEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRocketMQGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRocketMQGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQNamespaceRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间名称，3-64个字符，只能包含字母、数字、“-”及“_”
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 未消费消息的保留时间，以毫秒为单位，60秒-15天
	Ttl *uint64 `json:"Ttl,omitempty" name:"Ttl"`

	// 消息持久化后保留的时间，以毫秒为单位
	RetentionTime *uint64 `json:"RetentionTime,omitempty" name:"RetentionTime"`

	// 说明，最大128个字符
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyRocketMQNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "Ttl")
	delete(f, "RetentionTime")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRocketMQNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRocketMQNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQTopicRequest struct {
	*tchttp.BaseRequest

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 主题名称
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 说明信息，最大128个字符
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 分区数，全局类型无效，不可小于当前分区数
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`
}

func (r *ModifyRocketMQTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "Topic")
	delete(f, "Remark")
	delete(f, "PartitionNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRocketMQTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRocketMQTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRoleRequest struct {
	*tchttp.BaseRequest

	// 角色名称，不支持中字以及除了短线和下划线外的特殊字符且长度必须大于0且小等于32。
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 备注说明，长度必须大等于0且小等于128。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ModifyRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleName")
	delete(f, "Remark")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 角色名称
		RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

		// 备注说明
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 分区数，必须大于或者等于原分区数，若想维持原分区数请输入原数目，修改分区数仅对非全局顺序消息起效果，不允许超过128个分区。
	Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`

	// 备注，128字符以内。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ModifyTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "Partitions")
	delete(f, "Remark")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分区数
		Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`

		// 备注，128字符以内。
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PartitionsTopic struct {

	// 最后一次间隔内发布消息的平均byte大小。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AverageMsgSize *string `json:"AverageMsgSize,omitempty" name:"AverageMsgSize"`

	// 消费者数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerCount *string `json:"ConsumerCount,omitempty" name:"ConsumerCount"`

	// 被记录下来的消息总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastConfirmedEntry *string `json:"LastConfirmedEntry,omitempty" name:"LastConfirmedEntry"`

	// 最后一个ledger创建的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastLedgerCreatedTimestamp *string `json:"LastLedgerCreatedTimestamp,omitempty" name:"LastLedgerCreatedTimestamp"`

	// 本地和复制的发布者每秒发布消息的速率。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateIn *string `json:"MsgRateIn,omitempty" name:"MsgRateIn"`

	// 本地和复制的消费者每秒分发消息的数量之和。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateOut *string `json:"MsgRateOut,omitempty" name:"MsgRateOut"`

	// 本地和复制的发布者每秒发布消息的byte。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputIn *string `json:"MsgThroughputIn,omitempty" name:"MsgThroughputIn"`

	// 本地和复制的消费者每秒分发消息的byte。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputOut *string `json:"MsgThroughputOut,omitempty" name:"MsgThroughputOut"`

	// 被记录下来的消息总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumberOfEntries *string `json:"NumberOfEntries,omitempty" name:"NumberOfEntries"`

	// 子分区id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions *int64 `json:"Partitions,omitempty" name:"Partitions"`

	// 生产者数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProducerCount *string `json:"ProducerCount,omitempty" name:"ProducerCount"`

	// 以byte计算的所有消息存储总量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSize *string `json:"TotalSize,omitempty" name:"TotalSize"`

	// topic类型描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicType *uint64 `json:"TopicType,omitempty" name:"TopicType"`
}

type Producer struct {

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 连接数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountConnect *int64 `json:"CountConnect,omitempty" name:"CountConnect"`

	// 连接集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectionSets []*Connection `json:"ConnectionSets,omitempty" name:"ConnectionSets"`
}

type PublishCmqMsgRequest struct {
	*tchttp.BaseRequest

	// 主题名
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 消息内容
	MsgContent *string `json:"MsgContent,omitempty" name:"MsgContent"`

	// 消息标签
	MsgTag []*string `json:"MsgTag,omitempty" name:"MsgTag"`
}

func (r *PublishCmqMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishCmqMsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "MsgContent")
	delete(f, "MsgTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PublishCmqMsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type PublishCmqMsgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true表示发送成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 消息id
		MsgId *string `json:"MsgId,omitempty" name:"MsgId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishCmqMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishCmqMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Publisher struct {

	// 生产者id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProducerId *int64 `json:"ProducerId,omitempty" name:"ProducerId"`

	// 生产者名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`

	// 生产者地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitempty" name:"Address"`

	// 客户端版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientVersion *string `json:"ClientVersion,omitempty" name:"ClientVersion"`

	// 消息生产速率（条/秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateIn *float64 `json:"MsgRateIn,omitempty" name:"MsgRateIn"`

	// 消息生产吞吐速率（字节/秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputIn *float64 `json:"MsgThroughputIn,omitempty" name:"MsgThroughputIn"`

	// 平均消息大小（字节）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AverageMsgSize *float64 `json:"AverageMsgSize,omitempty" name:"AverageMsgSize"`

	// 连接时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectedSince *string `json:"ConnectedSince,omitempty" name:"ConnectedSince"`

	// 生产者连接的主题分区号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`
}

type ReceiveMessageRequest struct {
	*tchttp.BaseRequest

	// 接收消息的topic的名字, 这里尽量需要使用topic的全路径，如果不指定，即：tenant/namespace/topic。默认使用的是：public/default
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 订阅者的名字
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`

	// 默认值为1000，consumer接收的消息会首先存储到receiverQueueSize这个队列中，用作调优接收消息的速率
	ReceiverQueueSize *int64 `json:"ReceiverQueueSize,omitempty" name:"ReceiverQueueSize"`

	// 默认值为：Latest。用作判定consumer初始接收消息的位置，可选参数为：Earliest, Latest
	SubInitialPosition *string `json:"SubInitialPosition,omitempty" name:"SubInitialPosition"`
}

func (r *ReceiveMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReceiveMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Topic")
	delete(f, "SubscriptionName")
	delete(f, "ReceiverQueueSize")
	delete(f, "SubInitialPosition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReceiveMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ReceiveMessageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用作标识消息的唯一主键
		MessageID *string `json:"MessageID,omitempty" name:"MessageID"`

		// 接收消息的内容
		MessagePayload *string `json:"MessagePayload,omitempty" name:"MessagePayload"`

		// 提供给 Ack 接口，用来Ack哪一个topic中的消息
		AckTopic *string `json:"AckTopic,omitempty" name:"AckTopic"`

		// 返回的错误信息，如果为空，说明没有错误
	// 注意：此字段可能返回 null，表示取不到有效值。
		ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

		// 返回订阅者的名字，用来创建 ack consumer时使用
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubName *string `json:"SubName,omitempty" name:"SubName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReceiveMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReceiveMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetMsgSubOffsetByTimestampRequest struct {
	*tchttp.BaseRequest

	// 命名空间名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 订阅者名称。
	Subscription *string `json:"Subscription,omitempty" name:"Subscription"`

	// 时间戳，精确到毫秒。
	ToTimestamp *uint64 `json:"ToTimestamp,omitempty" name:"ToTimestamp"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ResetMsgSubOffsetByTimestampRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetMsgSubOffsetByTimestampRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "Subscription")
	delete(f, "ToTimestamp")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetMsgSubOffsetByTimestampRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResetMsgSubOffsetByTimestampResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetMsgSubOffsetByTimestampResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetMsgSubOffsetByTimestampResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetentionPolicy struct {

	// 消息保留时长
	TimeInMinutes *int64 `json:"TimeInMinutes,omitempty" name:"TimeInMinutes"`

	// 消息保留大小
	SizeInMB *int64 `json:"SizeInMB,omitempty" name:"SizeInMB"`
}

type RewindCmqQueueRequest struct {
	*tchttp.BaseRequest

	// 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 设定该时间，则（Batch）receiveMessage接口，会按照生产消息的先后顺序消费该时间戳以后的消息。
	StartConsumeTime *uint64 `json:"StartConsumeTime,omitempty" name:"StartConsumeTime"`
}

func (r *RewindCmqQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RewindCmqQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	delete(f, "StartConsumeTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RewindCmqQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RewindCmqQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RewindCmqQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RewindCmqQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RocketMQClusterConfig struct {

	// 单命名空间TPS上线
	MaxTpsPerNamespace *uint64 `json:"MaxTpsPerNamespace,omitempty" name:"MaxTpsPerNamespace"`

	// 最大命名空间数量
	MaxNamespaceNum *uint64 `json:"MaxNamespaceNum,omitempty" name:"MaxNamespaceNum"`

	// 已使用命名空间数量
	UsedNamespaceNum *uint64 `json:"UsedNamespaceNum,omitempty" name:"UsedNamespaceNum"`

	// 最大Topic数量
	MaxTopicNum *uint64 `json:"MaxTopicNum,omitempty" name:"MaxTopicNum"`

	// 已使用Topic数量
	UsedTopicNum *uint64 `json:"UsedTopicNum,omitempty" name:"UsedTopicNum"`

	// 最大Group数量
	MaxGroupNum *uint64 `json:"MaxGroupNum,omitempty" name:"MaxGroupNum"`

	// 已使用Group数量
	UsedGroupNum *uint64 `json:"UsedGroupNum,omitempty" name:"UsedGroupNum"`

	// 消息最大保留时间，以毫秒为单位
	MaxRetentionTime *uint64 `json:"MaxRetentionTime,omitempty" name:"MaxRetentionTime"`

	// 消息最长延时，以毫秒为单位
	MaxLatencyTime *uint64 `json:"MaxLatencyTime,omitempty" name:"MaxLatencyTime"`
}

type RocketMQClusterDetail struct {

	// 集群基本信息
	Info *RocketMQClusterInfo `json:"Info,omitempty" name:"Info"`

	// 集群配置信息
	Config *RocketMQClusterConfig `json:"Config,omitempty" name:"Config"`

	// 集群状态，0:创建中，1:正常，2:销毁中，3:已删除，4: 隔离中，5:创建失败，6: 删除失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type RocketMQClusterInfo struct {

	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 地域信息
	Region *string `json:"Region,omitempty" name:"Region"`

	// 创建时间，毫秒为单位
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 集群说明信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 公网接入地址
	PublicEndPoint *string `json:"PublicEndPoint,omitempty" name:"PublicEndPoint"`

	// VPC接入地址
	VpcEndPoint *string `json:"VpcEndPoint,omitempty" name:"VpcEndPoint"`
}

type RocketMQClusterRecentStats struct {

	// Topic数量
	TopicNum *uint64 `json:"TopicNum,omitempty" name:"TopicNum"`

	// 消息生产数
	ProducedMsgNum *uint64 `json:"ProducedMsgNum,omitempty" name:"ProducedMsgNum"`

	// 消息消费数
	ConsumedMsgNum *uint64 `json:"ConsumedMsgNum,omitempty" name:"ConsumedMsgNum"`

	// 消息堆积数
	AccumulativeMsgNum *uint64 `json:"AccumulativeMsgNum,omitempty" name:"AccumulativeMsgNum"`
}

type RocketMQGroup struct {

	// 消费组名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 在线消费者数量
	ConsumerNum *uint64 `json:"ConsumerNum,omitempty" name:"ConsumerNum"`

	// 消费TPS
	TPS *uint64 `json:"TPS,omitempty" name:"TPS"`

	// 总堆积数量
	TotalAccumulative *int64 `json:"TotalAccumulative,omitempty" name:"TotalAccumulative"`

	// 0表示集群消费模式，1表示广播消费模式，-1表示未知
	ConsumptionMode *int64 `json:"ConsumptionMode,omitempty" name:"ConsumptionMode"`

	// 是否允许消费
	ReadEnabled *bool `json:"ReadEnabled,omitempty" name:"ReadEnabled"`

	// 重试队列分区数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryPartitionNum *uint64 `json:"RetryPartitionNum,omitempty" name:"RetryPartitionNum"`

	// 创建时间，以毫秒为单位
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 修改时间，以毫秒为单位
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 客户端协议
	ClientProtocol *string `json:"ClientProtocol,omitempty" name:"ClientProtocol"`

	// 说明信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 消费者类型，枚举值ACTIVELY, PASSIVELY
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerType *string `json:"ConsumerType,omitempty" name:"ConsumerType"`

	// 是否开启广播消费
	BroadcastEnabled *bool `json:"BroadcastEnabled,omitempty" name:"BroadcastEnabled"`
}

type RocketMQNamespace struct {

	// 命名空间名称，3-64个字符，只能包含字母、数字、“-”及“_”
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// 未消费消息的保留时间，以毫秒单位，范围60秒到15天
	Ttl *uint64 `json:"Ttl,omitempty" name:"Ttl"`

	// 消息持久化后保留的时间，以毫秒单位
	RetentionTime *uint64 `json:"RetentionTime,omitempty" name:"RetentionTime"`

	// 说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type RocketMQTopic struct {

	// 主题名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 读写分区数
	PartitionNum *uint64 `json:"PartitionNum,omitempty" name:"PartitionNum"`

	// 创建时间，以毫秒为单位
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 创建时间，以毫秒为单位
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type Role struct {

	// 角色名称。
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 角色token值。
	Token *string `json:"Token,omitempty" name:"Token"`

	// 备注说明。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type SendBatchMessagesRequest struct {
	*tchttp.BaseRequest

	// 消息要发送的topic的名字, 这里尽量需要使用topic的全路径，即：tenant/namespace/topic。如果不指定，默认使用的是：public/default
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 需要发送消息的内容
	Payload *string `json:"Payload,omitempty" name:"Payload"`

	// String 类型的 token，可以不填，系统会自动获取
	StringToken *string `json:"StringToken,omitempty" name:"StringToken"`

	// producer 的名字，要求全局是唯一的，如果不设置，系统会自动生成
	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`

	// 单位：s。消息发送的超时时间。默认值为：30s
	SendTimeout *int64 `json:"SendTimeout,omitempty" name:"SendTimeout"`

	// 内存中允许缓存的生产消息的最大数量，默认值：1000条
	MaxPendingMessages *int64 `json:"MaxPendingMessages,omitempty" name:"MaxPendingMessages"`

	// 每一个batch中消息的最大数量，默认值：1000条/batch
	BatchingMaxMessages *int64 `json:"BatchingMaxMessages,omitempty" name:"BatchingMaxMessages"`

	// 每一个batch最大等待的时间，超过这个时间，不管是否达到指定的batch中消息的数量和大小，都会将该batch发送出去，默认：10ms
	BatchingMaxPublishDelay *int64 `json:"BatchingMaxPublishDelay,omitempty" name:"BatchingMaxPublishDelay"`

	// 每一个batch中最大允许的消息的大小，默认：128KB
	BatchingMaxBytes *int64 `json:"BatchingMaxBytes,omitempty" name:"BatchingMaxBytes"`
}

func (r *SendBatchMessagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendBatchMessagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Topic")
	delete(f, "Payload")
	delete(f, "StringToken")
	delete(f, "ProducerName")
	delete(f, "SendTimeout")
	delete(f, "MaxPendingMessages")
	delete(f, "BatchingMaxMessages")
	delete(f, "BatchingMaxPublishDelay")
	delete(f, "BatchingMaxBytes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendBatchMessagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SendBatchMessagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 消息的唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
		MessageId *string `json:"MessageId,omitempty" name:"MessageId"`

		// 错误消息，返回为 ""，代表没有错误
	// 注意：此字段可能返回 null，表示取不到有效值。
		ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendBatchMessagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendBatchMessagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendCmqMsgRequest struct {
	*tchttp.BaseRequest

	// 队列名
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 消息内容
	MsgContent *string `json:"MsgContent,omitempty" name:"MsgContent"`

	// 延迟时间
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" name:"DelaySeconds"`
}

func (r *SendCmqMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendCmqMsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	delete(f, "MsgContent")
	delete(f, "DelaySeconds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendCmqMsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SendCmqMsgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// true表示发送成功
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 消息id
		MsgId *string `json:"MsgId,omitempty" name:"MsgId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendCmqMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendCmqMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendMessagesRequest struct {
	*tchttp.BaseRequest

	// 消息要发送的topic的名字, 这里尽量需要使用topic的全路径，即：tenant/namespace/topic。如果不指定，默认使用的是：public/default
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 要发送的消息的内容
	Payload *string `json:"Payload,omitempty" name:"Payload"`

	// Token 是用来做鉴权使用的，可以不填，系统会自动获取
	StringToken *string `json:"StringToken,omitempty" name:"StringToken"`

	// 设置 producer 的名字，要求全局唯一，用户不配置，系统会随机生成
	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`

	// 设置消息发送的超时时间，默认为30s
	SendTimeout *int64 `json:"SendTimeout,omitempty" name:"SendTimeout"`

	// 内存中缓存的最大的生产消息的数量，默认为1000条
	MaxPendingMessages *int64 `json:"MaxPendingMessages,omitempty" name:"MaxPendingMessages"`
}

func (r *SendMessagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendMessagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Topic")
	delete(f, "Payload")
	delete(f, "StringToken")
	delete(f, "ProducerName")
	delete(f, "SendTimeout")
	delete(f, "MaxPendingMessages")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendMessagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SendMessagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 消息的messageID, 是全局唯一的，用来标识消息的元数据信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		MessageId *string `json:"MessageId,omitempty" name:"MessageId"`

		// 返回的错误消息，如果返回为 “”，说明没有错误
	// 注意：此字段可能返回 null，表示取不到有效值。
		ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendMessagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendMessagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendMsgRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名称，如果是分区topic需要指定具体分区，如果没有指定则默认发到0分区，例如：my_topic-partition-0。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 消息内容，不能为空且大小不得大于5242880个byte。
	MsgContent *string `json:"MsgContent,omitempty" name:"MsgContent"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *SendMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendMsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "MsgContent")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendMsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SendMsgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Sort struct {

	// 排序字段
	Name *string `json:"Name,omitempty" name:"Name"`

	// 升序ASC，降序DESC
	Order *string `json:"Order,omitempty" name:"Order"`
}

type Subscription struct {

	// 主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 消费者开始连接的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectedSince *string `json:"ConnectedSince,omitempty" name:"ConnectedSince"`

	// 消费者地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerAddr *string `json:"ConsumerAddr,omitempty" name:"ConsumerAddr"`

	// 消费者数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerCount *string `json:"ConsumerCount,omitempty" name:"ConsumerCount"`

	// 消费者名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerName *string `json:"ConsumerName,omitempty" name:"ConsumerName"`

	// 堆积的消息数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgBacklog *string `json:"MsgBacklog,omitempty" name:"MsgBacklog"`

	// 于TTL，此订阅下没有被发送而是被丢弃的比例。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateExpired *string `json:"MsgRateExpired,omitempty" name:"MsgRateExpired"`

	// 消费者每秒分发消息的数量之和。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateOut *string `json:"MsgRateOut,omitempty" name:"MsgRateOut"`

	// 消费者每秒消息的byte。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputOut *string `json:"MsgThroughputOut,omitempty" name:"MsgThroughputOut"`

	// 订阅名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`

	// 消费者集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerSets []*Consumer `json:"ConsumerSets,omitempty" name:"ConsumerSets"`

	// 是否在线。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsOnline *bool `json:"IsOnline,omitempty" name:"IsOnline"`

	// 消费进度集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumersScheduleSets []*ConsumersSchedule `json:"ConsumersScheduleSets,omitempty" name:"ConsumersScheduleSets"`

	// 备注。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最近修改时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type SubscriptionTopic struct {

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 订阅名称。
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
}

type Tag struct {

	// 标签的key的值
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签的Value的值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type Topic struct {

	// 最后一次间隔内发布消息的平均byte大小。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AverageMsgSize *string `json:"AverageMsgSize,omitempty" name:"AverageMsgSize"`

	// 消费者数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerCount *string `json:"ConsumerCount,omitempty" name:"ConsumerCount"`

	// 被记录下来的消息总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastConfirmedEntry *string `json:"LastConfirmedEntry,omitempty" name:"LastConfirmedEntry"`

	// 最后一个ledger创建的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastLedgerCreatedTimestamp *string `json:"LastLedgerCreatedTimestamp,omitempty" name:"LastLedgerCreatedTimestamp"`

	// 本地和复制的发布者每秒发布消息的速率。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateIn *string `json:"MsgRateIn,omitempty" name:"MsgRateIn"`

	// 本地和复制的消费者每秒分发消息的数量之和。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateOut *string `json:"MsgRateOut,omitempty" name:"MsgRateOut"`

	// 本地和复制的发布者每秒发布消息的byte。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputIn *string `json:"MsgThroughputIn,omitempty" name:"MsgThroughputIn"`

	// 本地和复制的消费者每秒分发消息的byte。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputOut *string `json:"MsgThroughputOut,omitempty" name:"MsgThroughputOut"`

	// 被记录下来的消息总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumberOfEntries *string `json:"NumberOfEntries,omitempty" name:"NumberOfEntries"`

	// 分区数<=0：topic下无子分区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions *int64 `json:"Partitions,omitempty" name:"Partitions"`

	// 生产者数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProducerCount *string `json:"ProducerCount,omitempty" name:"ProducerCount"`

	// 以byte计算的所有消息存储总量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSize *string `json:"TotalSize,omitempty" name:"TotalSize"`

	// 分区topic里面的子分区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubTopicSets []*PartitionsTopic `json:"SubTopicSets,omitempty" name:"SubTopicSets"`

	// topic类型描述：
	// 0：普通消息；
	// 1：全局顺序消息；
	// 2：局部顺序消息；
	// 3：重试队列；
	// 4：死信队列；
	// 5：事务消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicType *uint64 `json:"TopicType,omitempty" name:"TopicType"`

	// 环境（命名空间）名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 说明，128个字符以内。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最近修改时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 生产者上限。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProducerLimit *string `json:"ProducerLimit,omitempty" name:"ProducerLimit"`

	// 消费者上限。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerLimit *string `json:"ConsumerLimit,omitempty" name:"ConsumerLimit"`
}

type TopicRecord struct {

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type UnbindCmqDeadLetterRequest struct {
	*tchttp.BaseRequest

	// 死信策略源队列名称，调用本接口会清空该队列的死信队列策略。
	SourceQueueName *string `json:"SourceQueueName,omitempty" name:"SourceQueueName"`
}

func (r *UnbindCmqDeadLetterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindCmqDeadLetterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceQueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindCmqDeadLetterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UnbindCmqDeadLetterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindCmqDeadLetterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindCmqDeadLetterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcBindRecord struct {

	// 租户Vpc Id
	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`

	// 租户Vpc子网Id
	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`

	// 路由Id
	RouterId *string `json:"RouterId,omitempty" name:"RouterId"`

	// Vpc的Id
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// Vpc的Port
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 说明，128个字符以内
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}
