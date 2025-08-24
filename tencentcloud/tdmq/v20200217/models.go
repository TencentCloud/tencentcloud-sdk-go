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

package v20200217

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AMQPClusterConfig struct {
	// 单Vhost TPS上限
	MaxTpsPerVHost *uint64 `json:"MaxTpsPerVHost,omitnil,omitempty" name:"MaxTpsPerVHost"`

	// 单Vhost客户端连接数上限
	MaxConnNumPerVHost *uint64 `json:"MaxConnNumPerVHost,omitnil,omitempty" name:"MaxConnNumPerVHost"`

	// 最大Vhost数量
	MaxVHostNum *uint64 `json:"MaxVHostNum,omitnil,omitempty" name:"MaxVHostNum"`

	// 最大exchange数量
	MaxExchangeNum *uint64 `json:"MaxExchangeNum,omitnil,omitempty" name:"MaxExchangeNum"`

	// 最大Queue数量
	MaxQueueNum *uint64 `json:"MaxQueueNum,omitnil,omitempty" name:"MaxQueueNum"`

	// 消息最大保留时间，以毫秒为单位
	MaxRetentionTime *uint64 `json:"MaxRetentionTime,omitnil,omitempty" name:"MaxRetentionTime"`

	// 已使用Vhost数量
	UsedVHostNum *uint64 `json:"UsedVHostNum,omitnil,omitempty" name:"UsedVHostNum"`

	// 已使用exchange数量
	UsedExchangeNum *uint64 `json:"UsedExchangeNum,omitnil,omitempty" name:"UsedExchangeNum"`

	// 已使用queue数量
	UsedQueueNum *uint64 `json:"UsedQueueNum,omitnil,omitempty" name:"UsedQueueNum"`
}

type AMQPClusterDetail struct {
	// 集群基本信息
	Info *AMQPClusterInfo `json:"Info,omitnil,omitempty" name:"Info"`

	// 集群配置信息
	Config *AMQPClusterConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 集群状态，0:创建中，1:正常，2:销毁中，3:已删除，4: 隔离中，5:创建失败，6: 删除失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type AMQPClusterInfo struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 地域信息
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 创建时间，毫秒为单位
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 集群说明信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 公网接入地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicEndPoint *string `json:"PublicEndPoint,omitnil,omitempty" name:"PublicEndPoint"`

	// VPC接入地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcEndPoint *string `json:"VpcEndPoint,omitnil,omitempty" name:"VpcEndPoint"`
}

// Predefined struct for user
type AcknowledgeMessageRequestParams struct {
	// 用作标识消息的唯一的ID（可从 receiveMessage 的返回值中获得）
	MessageId *string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// Topic 名字（可从 receiveMessage 的返回值中获得）这里尽量需要使用topic的全路径，即：tenant/namespace/topic。如果不指定，默认使用的是：public/default
	AckTopic *string `json:"AckTopic,omitnil,omitempty" name:"AckTopic"`

	// 订阅者的名字，可以从receiveMessage的返回值中获取到。这里尽量与receiveMessage中的订阅者保持一致，否则没办法正确ack 接收回来的消息。
	SubName *string `json:"SubName,omitnil,omitempty" name:"SubName"`
}

type AcknowledgeMessageRequest struct {
	*tchttp.BaseRequest
	
	// 用作标识消息的唯一的ID（可从 receiveMessage 的返回值中获得）
	MessageId *string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// Topic 名字（可从 receiveMessage 的返回值中获得）这里尽量需要使用topic的全路径，即：tenant/namespace/topic。如果不指定，默认使用的是：public/default
	AckTopic *string `json:"AckTopic,omitnil,omitempty" name:"AckTopic"`

	// 订阅者的名字，可以从receiveMessage的返回值中获取到。这里尽量与receiveMessage中的订阅者保持一致，否则没办法正确ack 接收回来的消息。
	SubName *string `json:"SubName,omitnil,omitempty" name:"SubName"`
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

// Predefined struct for user
type AcknowledgeMessageResponseParams struct {
	// 如果为""，则说明没有错误返回，否则返回具体的错误信息。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AcknowledgeMessageResponse struct {
	*tchttp.BaseResponse
	Response *AcknowledgeMessageResponseParams `json:"Response"`
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
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`
}

type CertificateInfo struct {
	// SSL证书管理中的id
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 证书到期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 证书绑定的域名
	DomainName *string `json:"DomainName,omitnil,omitempty" name:"DomainName"`

	// 证书状态：0 已签发
	// 1 即将过期
	// 2 未启用
	// 3 已过期
	// 4 不可用
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 证书类型：0：根证书，1：服务端证书
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// TencentCloud：SSL证书；Default：TDMQ官方默认证书
	Origin *string `json:"Origin,omitnil,omitempty" name:"Origin"`

	// 证书添加/更新时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`
}

// Predefined struct for user
type ClearCmqQueueRequestParams struct {
	// 队列名字，在单个地域同一账号下唯一。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`
}

type ClearCmqQueueRequest struct {
	*tchttp.BaseRequest
	
	// 队列名字，在单个地域同一账号下唯一。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`
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

// Predefined struct for user
type ClearCmqQueueResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ClearCmqQueueResponse struct {
	*tchttp.BaseResponse
	Response *ClearCmqQueueResponseParams `json:"Response"`
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

// Predefined struct for user
type ClearCmqSubscriptionFilterTagsRequestParams struct {
	// 主题名字，在单个地域同一账号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 订阅名字，在单个地域同一账号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`
}

type ClearCmqSubscriptionFilterTagsRequest struct {
	*tchttp.BaseRequest
	
	// 主题名字，在单个地域同一账号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 订阅名字，在单个地域同一账号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`
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

// Predefined struct for user
type ClearCmqSubscriptionFilterTagsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ClearCmqSubscriptionFilterTagsResponse struct {
	*tchttp.BaseResponse
	Response *ClearCmqSubscriptionFilterTagsResponseParams `json:"Response"`
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

type ClientSubscriptionInfo struct {
	// 客户端ID
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 客户端IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientAddr *string `json:"ClientAddr,omitnil,omitempty" name:"ClientAddr"`

	// 订阅主题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 订阅表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubString *string `json:"SubString,omitnil,omitempty" name:"SubString"`

	// 订阅方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpressionType *string `json:"ExpressionType,omitnil,omitempty" name:"ExpressionType"`
}

type Cluster struct {
	// 集群Id。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 说明信息。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 接入点数量
	EndPointNum *int64 `json:"EndPointNum,omitnil,omitempty" name:"EndPointNum"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 集群是否健康，1表示健康，0表示异常
	Healthy *int64 `json:"Healthy,omitnil,omitempty" name:"Healthy"`

	// 集群健康信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthyInfo *string `json:"HealthyInfo,omitnil,omitempty" name:"HealthyInfo"`

	// 集群状态，0:创建中，1:正常，2:销毁中，3:已删除，4: 隔离中，5:创建失败，6: 删除失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 最大命名空间数量
	MaxNamespaceNum *int64 `json:"MaxNamespaceNum,omitnil,omitempty" name:"MaxNamespaceNum"`

	// 最大Topic数量
	MaxTopicNum *int64 `json:"MaxTopicNum,omitnil,omitempty" name:"MaxTopicNum"`

	// 最大QPS
	MaxQps *int64 `json:"MaxQps,omitnil,omitempty" name:"MaxQps"`

	// 最大消息保留时间，秒为单位
	MessageRetentionTime *int64 `json:"MessageRetentionTime,omitnil,omitempty" name:"MessageRetentionTime"`

	// 最大存储容量
	MaxStorageCapacity *int64 `json:"MaxStorageCapacity,omitnil,omitempty" name:"MaxStorageCapacity"`

	// 集群版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 公网访问接入点
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicEndPoint *string `json:"PublicEndPoint,omitnil,omitempty" name:"PublicEndPoint"`

	// VPC访问接入点
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcEndPoint *string `json:"VpcEndPoint,omitnil,omitempty" name:"VpcEndPoint"`

	// 命名空间数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceNum *int64 `json:"NamespaceNum,omitnil,omitempty" name:"NamespaceNum"`

	// 已使用存储限制，MB为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedStorageBudget *int64 `json:"UsedStorageBudget,omitnil,omitempty" name:"UsedStorageBudget"`

	// 最大生产消息速率，以条数为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxPublishRateInMessages *int64 `json:"MaxPublishRateInMessages,omitnil,omitempty" name:"MaxPublishRateInMessages"`

	// 最大推送消息速率，以条数为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDispatchRateInMessages *int64 `json:"MaxDispatchRateInMessages,omitnil,omitempty" name:"MaxDispatchRateInMessages"`

	// 最大生产消息速率，以字节为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxPublishRateInBytes *int64 `json:"MaxPublishRateInBytes,omitnil,omitempty" name:"MaxPublishRateInBytes"`

	// 最大推送消息速率，以字节为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDispatchRateInBytes *int64 `json:"MaxDispatchRateInBytes,omitnil,omitempty" name:"MaxDispatchRateInBytes"`

	// 已创建主题数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// 最长消息延时，以秒为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMessageDelayInSeconds *int64 `json:"MaxMessageDelayInSeconds,omitnil,omitempty" name:"MaxMessageDelayInSeconds"`

	// 是否开启公网访问，不填时默认开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitnil,omitempty" name:"PublicAccessEnabled"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 计费模式：
	// 0: 按量计费
	// 1: 包年包月
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 是否支持升级专业版实例
	UpgradeProInstance *bool `json:"UpgradeProInstance,omitnil,omitempty" name:"UpgradeProInstance"`
}

type CmqDeadLetterPolicy struct {
	// 死信队列。
	DeadLetterQueue *string `json:"DeadLetterQueue,omitnil,omitempty" name:"DeadLetterQueue"`

	// 死信队列策略。0:最大接收次数;1:最大未消费时间
	Policy *uint64 `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 最大未消费过期时间。Policy为1时必选。范围300-43200，单位秒，需要小于消息最大保留时间MsgRetentionSeconds。
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitnil,omitempty" name:"MaxTimeToLive"`

	// 最大接收次数。Policy为0时必选，范围在1到1000。
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitnil,omitempty" name:"MaxReceiveCount"`
}

type CmqDeadLetterSource struct {
	// 消息队列ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueId *string `json:"QueueId,omitnil,omitempty" name:"QueueId"`

	// 消息队列名字。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`
}

type CmqQueue struct {
	// 消息队列ID。
	QueueId *string `json:"QueueId,omitnil,omitempty" name:"QueueId"`

	// 消息队列名字。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 每秒钟生产消息条数的限制，消费消息的大小是该值的1.1倍。
	Qps *uint64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// 带宽限制。
	Bps *uint64 `json:"Bps,omitnil,omitempty" name:"Bps"`

	// 飞行消息最大保留时间，需要小于消息保留周期。
	MaxDelaySeconds *uint64 `json:"MaxDelaySeconds,omitnil,omitempty" name:"MaxDelaySeconds"`

	// 最大堆积消息数。取值范围在公测期间为 1,000,000 - 10,000,000，正式上线后范围可达到 1000,000-1000,000,000。默认取值在公测期间为 10,000,000，正式上线后为 100,000,000。
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitnil,omitempty" name:"MaxMsgHeapNum"`

	// 消息接收长轮询等待时间。取值范围0 - 30秒，默认值0。
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitnil,omitempty" name:"PollingWaitSeconds"`

	// 消息保留周期。取值范围60-1296000秒（1min-15天），默认值345600秒（4 天）。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil,omitempty" name:"MsgRetentionSeconds"`

	// 消息可见性超时。取值范围1 - 43200秒（即12小时内），默认值30。
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitnil,omitempty" name:"VisibilityTimeout"`

	// 消息最大长度。取值范围1024 - 1048576 Byte（即1K - 1024K），默认值65536。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil,omitempty" name:"MaxMsgSize"`

	// 回溯队列的消息回溯时间最大值，取值范围0 - 43200秒，0表示不开启消息回溯。
	RewindSeconds *uint64 `json:"RewindSeconds,omitnil,omitempty" name:"RewindSeconds"`

	// 队列的创建时间。返回 Unix 时间戳，精确到毫秒。
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后一次修改队列属性的时间。返回 Unix 时间戳，精确到毫秒。
	LastModifyTime *uint64 `json:"LastModifyTime,omitnil,omitempty" name:"LastModifyTime"`

	// 在队列中处于 Active 状态（不处于被消费状态）的消息总数，为近似值。
	ActiveMsgNum *uint64 `json:"ActiveMsgNum,omitnil,omitempty" name:"ActiveMsgNum"`

	// 在队列中处于 Inactive 状态（正处于被消费状态）的消息总数，为近似值。
	InactiveMsgNum *uint64 `json:"InactiveMsgNum,omitnil,omitempty" name:"InactiveMsgNum"`

	// 延迟消息数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelayMsgNum *uint64 `json:"DelayMsgNum,omitnil,omitempty" name:"DelayMsgNum"`

	// 已调用 DelMsg 接口删除，但还在回溯保留时间内的消息数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RewindMsgNum *uint64 `json:"RewindMsgNum,omitnil,omitempty" name:"RewindMsgNum"`

	// 消息最小未消费时间，单位为秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinMsgTime *uint64 `json:"MinMsgTime,omitnil,omitempty" name:"MinMsgTime"`

	// 事务消息队列。true表示是事务消息，false表示不是事务消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Transaction *bool `json:"Transaction,omitnil,omitempty" name:"Transaction"`

	// 死信队列。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterSource []*CmqDeadLetterSource `json:"DeadLetterSource,omitnil,omitempty" name:"DeadLetterSource"`

	// 死信队列策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterPolicy *CmqDeadLetterPolicy `json:"DeadLetterPolicy,omitnil,omitempty" name:"DeadLetterPolicy"`

	// 事务消息策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransactionPolicy *CmqTransactionPolicy `json:"TransactionPolicy,omitnil,omitempty" name:"TransactionPolicy"`

	// 创建者Uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUin *uint64 `json:"CreateUin,omitnil,omitempty" name:"CreateUin"`

	// 关联的标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 消息轨迹。true表示开启，false表示不开启。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trace *bool `json:"Trace,omitnil,omitempty" name:"Trace"`

	// 租户id
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 集群状态，0:创建中，1:正常，2:销毁中，3:已删除，4: 隔离中，5:创建失败，6: 删除失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 最大未确认消息数量
	MaxUnackedMsgNum *int64 `json:"MaxUnackedMsgNum,omitnil,omitempty" name:"MaxUnackedMsgNum"`

	// 最大消息堆积大小（字节）
	MaxMsgBacklogSize *int64 `json:"MaxMsgBacklogSize,omitnil,omitempty" name:"MaxMsgBacklogSize"`

	// 队列可回溯存储空间，取值范围1024MB - 10240MB，0表示不开启
	RetentionSizeInMB *uint64 `json:"RetentionSizeInMB,omitnil,omitempty" name:"RetentionSizeInMB"`
}

type CmqSubscription struct {
	// 订阅名字，在单个地域同一账号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`

	// 订阅 ID。订阅 ID 在拉取监控数据时会用到。
	SubscriptionId *string `json:"SubscriptionId,omitnil,omitempty" name:"SubscriptionId"`

	// 订阅拥有者的 APPID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicOwner *uint64 `json:"TopicOwner,omitnil,omitempty" name:"TopicOwner"`

	// 该订阅待投递的消息数。
	MsgCount *uint64 `json:"MsgCount,omitnil,omitempty" name:"MsgCount"`

	// 最后一次修改订阅属性的时间。返回 Unix 时间戳，精确到毫秒。
	LastModifyTime *uint64 `json:"LastModifyTime,omitnil,omitempty" name:"LastModifyTime"`

	// 订阅的创建时间。返回 Unix 时间戳，精确到毫秒。
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 表示订阅接收消息的过滤策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindingKey []*string `json:"BindingKey,omitnil,omitempty" name:"BindingKey"`

	// 接收通知的 endpoint，根据协议 protocol 区分：对于 HTTP，endpoint 必须以http://开头，host 可以是域名或 IP；对于 queue，则填 queueName。
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 描述用户创建订阅时选择的过滤策略：
	// filterType = 1表示用户使用 filterTag 标签过滤
	// filterType = 2表示用户使用 bindingKey 过滤。
	FilterTags []*string `json:"FilterTags,omitnil,omitempty" name:"FilterTags"`

	// 订阅的协议，目前支持两种协议：HTTP、queue。使用 HTTP 协议，用户需自己搭建接受消息的 Web Server。使用 queue，消息会自动推送到 CMQ queue，用户可以并发地拉取消息。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 向 endpoint 推送消息出现错误时，CMQ 推送服务器的重试策略。取值有：
	// （1）BACKOFF_RETRY，退避重试。每隔一定时间重试一次，重试够一定次数后，就把该消息丢弃，继续推送下一条消息；
	// （2）EXPONENTIAL_DECAY_RETRY，指数衰退重试。每次重试的间隔是指数递增的，例如开始 1s，后面是 2s，4s，8s...由于 Topic 消息的周期是一天，所以最多重试一天就把消息丢弃。默认值是 EXPONENTIAL_DECAY_RETRY。
	NotifyStrategy *string `json:"NotifyStrategy,omitnil,omitempty" name:"NotifyStrategy"`

	// 推送内容的格式。取值：（1）JSON；（2）SIMPLIFIED，即 raw 格式。如果 protocol 是 queue，则取值必须为 SIMPLIFIED。如果 protocol 是 HTTP，两个值均可以，默认值是 JSON。
	NotifyContentFormat *string `json:"NotifyContentFormat,omitnil,omitempty" name:"NotifyContentFormat"`

	// 订阅所属的主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type CmqTopic struct {
	// 主题的 ID。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 消息在主题中最长存活时间，从发送到该主题开始经过此参数指定的时间后，不论消息是否被成功推送给用户都将被删除，单位为秒。固定为一天（86400秒），该属性不能修改。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil,omitempty" name:"MsgRetentionSeconds"`

	// 消息最大长度。取值范围1024 - 1048576Byte（即1 - 1024K），默认值为1048576。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil,omitempty" name:"MaxMsgSize"`

	// 每秒钟发布消息的条数。
	Qps *uint64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// 描述用户创建订阅时选择的过滤策略：
	// FilterType = 1表示用户使用 FilterTag 标签过滤;
	// FilterType = 2表示用户使用 BindingKey 过滤。
	FilterType *uint64 `json:"FilterType,omitnil,omitempty" name:"FilterType"`

	// 主题的创建时间。返回 Unix 时间戳，精确到毫秒。
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后一次修改主题属性的时间。返回 Unix 时间戳，精确到毫秒。
	LastModifyTime *uint64 `json:"LastModifyTime,omitnil,omitempty" name:"LastModifyTime"`

	// 当前该主题中消息数目（消息堆积数）。
	MsgCount *uint64 `json:"MsgCount,omitnil,omitempty" name:"MsgCount"`

	// 创建者 Uin，CAM 鉴权 resource 由该字段组合而成。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUin *uint64 `json:"CreateUin,omitnil,omitempty" name:"CreateUin"`

	// 关联的标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 消息轨迹。true表示开启，false表示不开启。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trace *bool `json:"Trace,omitnil,omitempty" name:"Trace"`

	// 租户id
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 集群状态，0:创建中，1:正常，2:销毁中，3:已删除，4: 隔离中，5:创建失败，6: 删除失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 0表示pulsar，1表示rocketmq
	BrokerType *int64 `json:"BrokerType,omitnil,omitempty" name:"BrokerType"`

	// 订阅数量
	SubscriptionCount *int64 `json:"SubscriptionCount,omitnil,omitempty" name:"SubscriptionCount"`
}

type CmqTransactionPolicy struct {
	// 第一次回查时间。
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitnil,omitempty" name:"FirstQueryInterval"`

	// 最大查询次数。
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitnil,omitempty" name:"MaxQueryCount"`
}

type Consumer struct {
	// 消费者开始连接的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectedSince *string `json:"ConnectedSince,omitnil,omitempty" name:"ConnectedSince"`

	// 消费者地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerAddr *string `json:"ConsumerAddr,omitnil,omitempty" name:"ConsumerAddr"`

	// 消费者名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerName *string `json:"ConsumerName,omitnil,omitempty" name:"ConsumerName"`

	// 消费者版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientVersion *string `json:"ClientVersion,omitnil,omitempty" name:"ClientVersion"`

	// 消费者连接的主题分区号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`
}

type ConsumerLog struct {
	// 消息ID。
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 消费组。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`

	// 消费者名称。
	ConsumerName *string `json:"ConsumerName,omitnil,omitempty" name:"ConsumerName"`

	// 消费时间。
	ConsumeTime *string `json:"ConsumeTime,omitnil,omitempty" name:"ConsumeTime"`

	// 消费者客户端地址。
	ConsumerAddr *string `json:"ConsumerAddr,omitnil,omitempty" name:"ConsumerAddr"`

	// 消费耗时（毫秒）。
	ConsumeUseTime *uint64 `json:"ConsumeUseTime,omitnil,omitempty" name:"ConsumeUseTime"`

	// 消费状态。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ConsumerLogs struct {
	// 记录数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 消费日志。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerLogSets []*ConsumerLog `json:"ConsumerLogSets,omitnil,omitempty" name:"ConsumerLogSets"`
}

type ConsumerStats struct {
	// 主题名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 所属Broker
	// 注意：此字段可能返回 null，表示取不到有效值。
	BrokerName *string `json:"BrokerName,omitnil,omitempty" name:"BrokerName"`

	// 队列编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueId *int64 `json:"QueueId,omitnil,omitempty" name:"QueueId"`

	// 消费者ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerClientId *string `json:"ConsumerClientId,omitnil,omitempty" name:"ConsumerClientId"`

	// 消费位点
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerOffset *int64 `json:"ConsumerOffset,omitnil,omitempty" name:"ConsumerOffset"`

	// 服务端位点
	// 注意：此字段可能返回 null，表示取不到有效值。
	BrokerOffset *int64 `json:"BrokerOffset,omitnil,omitempty" name:"BrokerOffset"`

	// 消息堆积条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiffTotal *int64 `json:"DiffTotal,omitnil,omitempty" name:"DiffTotal"`

	// 最近消费时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTimestamp *int64 `json:"LastTimestamp,omitnil,omitempty" name:"LastTimestamp"`
}

type ConsumersSchedule struct {
	// 当前分区id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions *uint64 `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 消息数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumberOfEntries *uint64 `json:"NumberOfEntries,omitnil,omitempty" name:"NumberOfEntries"`

	// 消息积压数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgBacklog *uint64 `json:"MsgBacklog,omitnil,omitempty" name:"MsgBacklog"`

	// 消费者每秒分发消息的数量之和。
	MsgRateOut *string `json:"MsgRateOut,omitnil,omitempty" name:"MsgRateOut"`

	// 消费者每秒消息的byte。
	MsgThroughputOut *string `json:"MsgThroughputOut,omitnil,omitempty" name:"MsgThroughputOut"`

	// 超时丢弃比例。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateExpired *string `json:"MsgRateExpired,omitnil,omitempty" name:"MsgRateExpired"`
}

// Predefined struct for user
type CreateClusterRequestParams struct {
	// 集群名称，不支持中字以及除了短线和下划线外的特殊字符且不超过64个字符。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 用户专享物理集群ID，如果不传，则默认在公共集群上创建用户集群资源。
	BindClusterId *uint64 `json:"BindClusterId,omitnil,omitempty" name:"BindClusterId"`

	// 说明，128个字符以内。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 集群的标签列表(已废弃)
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否开启公网访问，不填时默认开启
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitnil,omitempty" name:"PublicAccessEnabled"`
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群名称，不支持中字以及除了短线和下划线外的特殊字符且不超过64个字符。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 用户专享物理集群ID，如果不传，则默认在公共集群上创建用户集群资源。
	BindClusterId *uint64 `json:"BindClusterId,omitnil,omitempty" name:"BindClusterId"`

	// 说明，128个字符以内。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 集群的标签列表(已废弃)
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否开启公网访问，不填时默认开启
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitnil,omitempty" name:"PublicAccessEnabled"`
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

// Predefined struct for user
type CreateClusterResponseParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateCmqQueueRequestParams struct {
	// 队列名字，在单个地域同一账号下唯一。队列名称是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 最大堆积消息数。取值范围在公测期间为 1,000,000 - 10,000,000，正式上线后范围可达到 1000,000-1000,000,000。默认取值在公测期间为 10,000,000，正式上线后为 100,000,000。
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitnil,omitempty" name:"MaxMsgHeapNum"`

	// 消息接收长轮询等待时间。取值范围 0-30 秒，默认值 0。
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitnil,omitempty" name:"PollingWaitSeconds"`

	// 消息可见性超时。取值范围 1-43200 秒（即12小时内），默认值 30。
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitnil,omitempty" name:"VisibilityTimeout"`

	// 消息最大长度。取值范围 1024-1048576 Byte（即1-1024K），默认值 1048576。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil,omitempty" name:"MaxMsgSize"`

	// 消息最长未确认时间。取值范围 30-43200 秒（30秒~12小时），默认值 3600 (1 小时)。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil,omitempty" name:"MsgRetentionSeconds"`

	// 队列是否开启回溯消息能力，该参数取值范围0-1296000，0表示不开启。
	RewindSeconds *uint64 `json:"RewindSeconds,omitnil,omitempty" name:"RewindSeconds"`

	// 1 表示事务队列，0 表示普通队列
	Transaction *uint64 `json:"Transaction,omitnil,omitempty" name:"Transaction"`

	// 第一次回查间隔
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitnil,omitempty" name:"FirstQueryInterval"`

	// 最大回查次数
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitnil,omitempty" name:"MaxQueryCount"`

	// 死信队列名称
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitnil,omitempty" name:"DeadLetterQueueName"`

	// 死信策略。0为消息被多次消费未删除，1为Time-To-Live过期
	Policy *uint64 `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 最大接收次数 1-1000
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitnil,omitempty" name:"MaxReceiveCount"`

	// policy为1时必选。最大未消费过期时间。范围300-43200，单位秒，需要小于消息最大保留时间msgRetentionSeconds
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitnil,omitempty" name:"MaxTimeToLive"`

	// 是否开启消息轨迹追踪，当不设置字段时，默认为不开启，该字段为true表示开启，为false表示不开启
	Trace *bool `json:"Trace,omitnil,omitempty" name:"Trace"`

	// 标签数组
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 队列可回溯存储空间：若开启消息回溯，取值范围：10240MB - 512000MB，若不开启消息回溯，取值：0
	RetentionSizeInMB *uint64 `json:"RetentionSizeInMB,omitnil,omitempty" name:"RetentionSizeInMB"`
}

type CreateCmqQueueRequest struct {
	*tchttp.BaseRequest
	
	// 队列名字，在单个地域同一账号下唯一。队列名称是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 最大堆积消息数。取值范围在公测期间为 1,000,000 - 10,000,000，正式上线后范围可达到 1000,000-1000,000,000。默认取值在公测期间为 10,000,000，正式上线后为 100,000,000。
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitnil,omitempty" name:"MaxMsgHeapNum"`

	// 消息接收长轮询等待时间。取值范围 0-30 秒，默认值 0。
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitnil,omitempty" name:"PollingWaitSeconds"`

	// 消息可见性超时。取值范围 1-43200 秒（即12小时内），默认值 30。
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitnil,omitempty" name:"VisibilityTimeout"`

	// 消息最大长度。取值范围 1024-1048576 Byte（即1-1024K），默认值 1048576。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil,omitempty" name:"MaxMsgSize"`

	// 消息最长未确认时间。取值范围 30-43200 秒（30秒~12小时），默认值 3600 (1 小时)。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil,omitempty" name:"MsgRetentionSeconds"`

	// 队列是否开启回溯消息能力，该参数取值范围0-1296000，0表示不开启。
	RewindSeconds *uint64 `json:"RewindSeconds,omitnil,omitempty" name:"RewindSeconds"`

	// 1 表示事务队列，0 表示普通队列
	Transaction *uint64 `json:"Transaction,omitnil,omitempty" name:"Transaction"`

	// 第一次回查间隔
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitnil,omitempty" name:"FirstQueryInterval"`

	// 最大回查次数
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitnil,omitempty" name:"MaxQueryCount"`

	// 死信队列名称
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitnil,omitempty" name:"DeadLetterQueueName"`

	// 死信策略。0为消息被多次消费未删除，1为Time-To-Live过期
	Policy *uint64 `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 最大接收次数 1-1000
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitnil,omitempty" name:"MaxReceiveCount"`

	// policy为1时必选。最大未消费过期时间。范围300-43200，单位秒，需要小于消息最大保留时间msgRetentionSeconds
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitnil,omitempty" name:"MaxTimeToLive"`

	// 是否开启消息轨迹追踪，当不设置字段时，默认为不开启，该字段为true表示开启，为false表示不开启
	Trace *bool `json:"Trace,omitnil,omitempty" name:"Trace"`

	// 标签数组
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 队列可回溯存储空间：若开启消息回溯，取值范围：10240MB - 512000MB，若不开启消息回溯，取值：0
	RetentionSizeInMB *uint64 `json:"RetentionSizeInMB,omitnil,omitempty" name:"RetentionSizeInMB"`
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
	delete(f, "Tags")
	delete(f, "RetentionSizeInMB")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCmqQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCmqQueueResponseParams struct {
	// 创建成功的queueId
	QueueId *string `json:"QueueId,omitnil,omitempty" name:"QueueId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCmqQueueResponse struct {
	*tchttp.BaseResponse
	Response *CreateCmqQueueResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateCmqSubscribeRequestParams struct {
	// 主题名字，在单个地域同一账号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 订阅名字，在单个地域同一账号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`

	// 订阅的协议，目前支持两种协议：http、queue。使用http协议，用户需自己搭建接受消息的web server。使用queue，消息会自动推送到CMQ queue，用户可以并发地拉取消息。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 接收通知的Endpoint，根据协议Protocol区分：对于http，Endpoint必须以“`http://`”开头，host可以是域名或IP；对于Queue，则填QueueName。 请注意，目前推送服务不能推送到私有网络中，因此Endpoint填写为私有网络域名或地址将接收不到推送的消息，目前支持推送到公网和基础网络。
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 向Endpoint推送消息出现错误时，CMQ推送服务器的重试策略。取值有：1）BACKOFF_RETRY，退避重试。每隔一定时间重试一次，重试够一定次数后，就把该消息丢弃，继续推送下一条消息；2）EXPONENTIAL_DECAY_RETRY，指数衰退重试。每次重试的间隔是指数递增的，例如开始1s，后面是2s，4s，8s...由于Topic消息的周期是一天，所以最多重试一天就把消息丢弃。默认值是EXPONENTIAL_DECAY_RETRY。
	NotifyStrategy *string `json:"NotifyStrategy,omitnil,omitempty" name:"NotifyStrategy"`

	// 消息正文。消息标签（用于消息过滤)。标签数量不能超过5个，每个标签不超过16个字符。与(Batch)PublishMessage的MsgTag参数配合使用，规则：1）如果FilterTag没有设置，则无论MsgTag是否有设置，订阅接收所有发布到Topic的消息；2）如果FilterTag数组有值，则只有数组中至少有一个值在MsgTag数组中也存在时（即FilterTag和MsgTag有交集），订阅才接收该发布到Topic的消息；3）如果FilterTag数组有值，但MsgTag没设置，则不接收任何发布到Topic的消息，可以认为是2）的一种特例，此时FilterTag和MsgTag没有交集。规则整体的设计思想是以订阅者的意愿为主。
	FilterTag []*string `json:"FilterTag,omitnil,omitempty" name:"FilterTag"`

	// BindingKey数量不超过5个， 每个BindingKey长度不超过64字节，该字段表示订阅接收消息的过滤策略，每个BindingKey最多含有15个“.”， 即最多16个词组。
	BindingKey []*string `json:"BindingKey,omitnil,omitempty" name:"BindingKey"`

	// 推送内容的格式。取值：1）JSON；2）SIMPLIFIED，即raw格式。如果Protocol是queue，则取值必须为SIMPLIFIED。如果Protocol是http，两个值均可以，默认值是JSON。
	NotifyContentFormat *string `json:"NotifyContentFormat,omitnil,omitempty" name:"NotifyContentFormat"`
}

type CreateCmqSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// 主题名字，在单个地域同一账号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 订阅名字，在单个地域同一账号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`

	// 订阅的协议，目前支持两种协议：http、queue。使用http协议，用户需自己搭建接受消息的web server。使用queue，消息会自动推送到CMQ queue，用户可以并发地拉取消息。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 接收通知的Endpoint，根据协议Protocol区分：对于http，Endpoint必须以“`http://`”开头，host可以是域名或IP；对于Queue，则填QueueName。 请注意，目前推送服务不能推送到私有网络中，因此Endpoint填写为私有网络域名或地址将接收不到推送的消息，目前支持推送到公网和基础网络。
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 向Endpoint推送消息出现错误时，CMQ推送服务器的重试策略。取值有：1）BACKOFF_RETRY，退避重试。每隔一定时间重试一次，重试够一定次数后，就把该消息丢弃，继续推送下一条消息；2）EXPONENTIAL_DECAY_RETRY，指数衰退重试。每次重试的间隔是指数递增的，例如开始1s，后面是2s，4s，8s...由于Topic消息的周期是一天，所以最多重试一天就把消息丢弃。默认值是EXPONENTIAL_DECAY_RETRY。
	NotifyStrategy *string `json:"NotifyStrategy,omitnil,omitempty" name:"NotifyStrategy"`

	// 消息正文。消息标签（用于消息过滤)。标签数量不能超过5个，每个标签不超过16个字符。与(Batch)PublishMessage的MsgTag参数配合使用，规则：1）如果FilterTag没有设置，则无论MsgTag是否有设置，订阅接收所有发布到Topic的消息；2）如果FilterTag数组有值，则只有数组中至少有一个值在MsgTag数组中也存在时（即FilterTag和MsgTag有交集），订阅才接收该发布到Topic的消息；3）如果FilterTag数组有值，但MsgTag没设置，则不接收任何发布到Topic的消息，可以认为是2）的一种特例，此时FilterTag和MsgTag没有交集。规则整体的设计思想是以订阅者的意愿为主。
	FilterTag []*string `json:"FilterTag,omitnil,omitempty" name:"FilterTag"`

	// BindingKey数量不超过5个， 每个BindingKey长度不超过64字节，该字段表示订阅接收消息的过滤策略，每个BindingKey最多含有15个“.”， 即最多16个词组。
	BindingKey []*string `json:"BindingKey,omitnil,omitempty" name:"BindingKey"`

	// 推送内容的格式。取值：1）JSON；2）SIMPLIFIED，即raw格式。如果Protocol是queue，则取值必须为SIMPLIFIED。如果Protocol是http，两个值均可以，默认值是JSON。
	NotifyContentFormat *string `json:"NotifyContentFormat,omitnil,omitempty" name:"NotifyContentFormat"`
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

// Predefined struct for user
type CreateCmqSubscribeResponseParams struct {
	// 订阅id
	SubscriptionId *string `json:"SubscriptionId,omitnil,omitempty" name:"SubscriptionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCmqSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *CreateCmqSubscribeResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateCmqTopicRequestParams struct {
	// 主题名字，在单个地域同一账号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 消息最大长度。取值范围 1024-65536 Byte（即1-64K），默认值 65536。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil,omitempty" name:"MaxMsgSize"`

	// 用于指定主题的消息匹配策略。1：表示标签匹配策略；2：表示路由匹配策略，默认值为标签匹配策略。
	FilterType *uint64 `json:"FilterType,omitnil,omitempty" name:"FilterType"`

	// 消息保存时间。取值范围60 - 86400 s（即1分钟 - 1天），默认值86400。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil,omitempty" name:"MsgRetentionSeconds"`

	// 是否开启消息轨迹标识，true表示开启，false表示不开启，不填表示不开启。
	Trace *bool `json:"Trace,omitnil,omitempty" name:"Trace"`

	// 标签数组
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateCmqTopicRequest struct {
	*tchttp.BaseRequest
	
	// 主题名字，在单个地域同一账号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 消息最大长度。取值范围 1024-65536 Byte（即1-64K），默认值 65536。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil,omitempty" name:"MaxMsgSize"`

	// 用于指定主题的消息匹配策略。1：表示标签匹配策略；2：表示路由匹配策略，默认值为标签匹配策略。
	FilterType *uint64 `json:"FilterType,omitnil,omitempty" name:"FilterType"`

	// 消息保存时间。取值范围60 - 86400 s（即1分钟 - 1天），默认值86400。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil,omitempty" name:"MsgRetentionSeconds"`

	// 是否开启消息轨迹标识，true表示开启，false表示不开启，不填表示不开启。
	Trace *bool `json:"Trace,omitnil,omitempty" name:"Trace"`

	// 标签数组
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCmqTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCmqTopicResponseParams struct {
	// 主题id
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCmqTopicResponse struct {
	*tchttp.BaseResponse
	Response *CreateCmqTopicResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateEnvironmentRequestParams struct {
	// 环境（命名空间）名称，不支持中字以及除了短线和下划线外的特殊字符且不超过16个字符。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 未消费消息过期时间，单位：秒，取值范围：60秒~15天。
	MsgTTL *uint64 `json:"MsgTTL,omitnil,omitempty" name:"MsgTTL"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 说明，128个字符以内。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 消息保留策略
	RetentionPolicy *RetentionPolicy `json:"RetentionPolicy,omitnil,omitempty" name:"RetentionPolicy"`

	// 是否开启自动创建订阅
	AutoSubscriptionCreation *bool `json:"AutoSubscriptionCreation,omitnil,omitempty" name:"AutoSubscriptionCreation"`
}

type CreateEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// 环境（命名空间）名称，不支持中字以及除了短线和下划线外的特殊字符且不超过16个字符。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 未消费消息过期时间，单位：秒，取值范围：60秒~15天。
	MsgTTL *uint64 `json:"MsgTTL,omitnil,omitempty" name:"MsgTTL"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 说明，128个字符以内。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 消息保留策略
	RetentionPolicy *RetentionPolicy `json:"RetentionPolicy,omitnil,omitempty" name:"RetentionPolicy"`

	// 是否开启自动创建订阅
	AutoSubscriptionCreation *bool `json:"AutoSubscriptionCreation,omitnil,omitempty" name:"AutoSubscriptionCreation"`
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
	delete(f, "ClusterId")
	delete(f, "Remark")
	delete(f, "RetentionPolicy")
	delete(f, "AutoSubscriptionCreation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnvironmentResponseParams struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 未消费消息过期时间，单位：秒。
	MsgTTL *uint64 `json:"MsgTTL,omitnil,omitempty" name:"MsgTTL"`

	// 说明，128个字符以内。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *CreateEnvironmentResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateEnvironmentRoleRequestParams struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 角色名称。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 授权项，最多只能包含produce、consume两项的非空字符串数组。
	Permissions []*string `json:"Permissions,omitnil,omitempty" name:"Permissions"`

	// 必填字段，集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type CreateEnvironmentRoleRequest struct {
	*tchttp.BaseRequest
	
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 角色名称。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 授权项，最多只能包含produce、consume两项的非空字符串数组。
	Permissions []*string `json:"Permissions,omitnil,omitempty" name:"Permissions"`

	// 必填字段，集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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

// Predefined struct for user
type CreateEnvironmentRoleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEnvironmentRoleResponse struct {
	*tchttp.BaseResponse
	Response *CreateEnvironmentRoleResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateProClusterRequestParams struct {
	// 多可用区部署选择三个可用区，示例[200002,200003,200004]
	// 
	// 单可用区部署选择一个可用区，示例[200002]
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 集群规格代号
	// 参考 [专业集群规格](https://cloud.tencent.com/document/product/1179/83705)
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 存储规格
	// 参考 [专业集群规格](https://cloud.tencent.com/document/product/1179/83705)
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 1: true，开启自动按月续费
	// 
	// 0: false，关闭自动按月续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 购买时长，取值范围：1～50
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 集群名称，不支持中字以及除了短线和下划线外的特殊字符且不超过64个字符。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// vpc网络标签
	Vpc *VpcInfo `json:"Vpc,omitnil,omitempty" name:"Vpc"`

	// 集群的标签列表(已废弃)
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateProClusterRequest struct {
	*tchttp.BaseRequest
	
	// 多可用区部署选择三个可用区，示例[200002,200003,200004]
	// 
	// 单可用区部署选择一个可用区，示例[200002]
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 集群规格代号
	// 参考 [专业集群规格](https://cloud.tencent.com/document/product/1179/83705)
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 存储规格
	// 参考 [专业集群规格](https://cloud.tencent.com/document/product/1179/83705)
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 1: true，开启自动按月续费
	// 
	// 0: false，关闭自动按月续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 购买时长，取值范围：1～50
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 集群名称，不支持中字以及除了短线和下划线外的特殊字符且不超过64个字符。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 是否自动选择代金券 1是 0否 默认为0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// vpc网络标签
	Vpc *VpcInfo `json:"Vpc,omitnil,omitempty" name:"Vpc"`

	// 集群的标签列表(已废弃)
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateProClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneIds")
	delete(f, "ProductName")
	delete(f, "StorageSize")
	delete(f, "AutoRenewFlag")
	delete(f, "TimeSpan")
	delete(f, "ClusterName")
	delete(f, "AutoVoucher")
	delete(f, "Vpc")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProClusterResponseParams struct {
	// 子订单号
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 订单号
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`

	// 集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateProClusterResponseParams `json:"Response"`
}

func (r *CreateProClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQBindingRequestParams struct {
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到Vhost名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 源exchange
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标类型,取值queue或exchange
	DestinationType *string `json:"DestinationType,omitnil,omitempty" name:"DestinationType"`

	// 目标
	Destination *string `json:"Destination,omitnil,omitempty" name:"Destination"`

	// 路由键
	RoutingKey *string `json:"RoutingKey,omitnil,omitempty" name:"RoutingKey"`
}

type CreateRabbitMQBindingRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到Vhost名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 源exchange
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标类型,取值queue或exchange
	DestinationType *string `json:"DestinationType,omitnil,omitempty" name:"DestinationType"`

	// 目标
	Destination *string `json:"Destination,omitnil,omitempty" name:"Destination"`

	// 路由键
	RoutingKey *string `json:"RoutingKey,omitnil,omitempty" name:"RoutingKey"`
}

func (r *CreateRabbitMQBindingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQBindingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "Source")
	delete(f, "DestinationType")
	delete(f, "Destination")
	delete(f, "RoutingKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRabbitMQBindingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQBindingResponseParams struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost名称
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 路由关系Id
	BindingId *int64 `json:"BindingId,omitnil,omitempty" name:"BindingId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRabbitMQBindingResponse struct {
	*tchttp.BaseResponse
	Response *CreateRabbitMQBindingResponseParams `json:"Response"`
}

func (r *CreateRabbitMQBindingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQBindingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQUserRequestParams struct {
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，登录时使用
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 密码，登录时使用。规范：不能为空，8-64个字符，至少要包含小写字母、大写字母、数字、特殊字符【()`~!@#$%^&*_=|{}[]:;',.?/】中的两项
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户标签，用于决定改用户访问RabbitMQ Management的权限范围
	// management：普通控制台用户，monitoring：管理型控制台用户，其他值：非控制台用户
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 该用户的最大连接数，不填写则不限制
	MaxConnections *int64 `json:"MaxConnections,omitnil,omitempty" name:"MaxConnections"`

	// 该用户的最大channel数，不填写则不限制
	MaxChannels *int64 `json:"MaxChannels,omitnil,omitempty" name:"MaxChannels"`
}

type CreateRabbitMQUserRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，登录时使用
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 密码，登录时使用。规范：不能为空，8-64个字符，至少要包含小写字母、大写字母、数字、特殊字符【()`~!@#$%^&*_=|{}[]:;',.?/】中的两项
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户标签，用于决定改用户访问RabbitMQ Management的权限范围
	// management：普通控制台用户，monitoring：管理型控制台用户，其他值：非控制台用户
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 该用户的最大连接数，不填写则不限制
	MaxConnections *int64 `json:"MaxConnections,omitnil,omitempty" name:"MaxConnections"`

	// 该用户的最大channel数，不填写则不限制
	MaxChannels *int64 `json:"MaxChannels,omitnil,omitempty" name:"MaxChannels"`
}

func (r *CreateRabbitMQUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "Password")
	delete(f, "Description")
	delete(f, "Tags")
	delete(f, "MaxConnections")
	delete(f, "MaxChannels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRabbitMQUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQUserResponseParams struct {
	// 用户名，登录时使用
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRabbitMQUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateRabbitMQUserResponseParams `json:"Response"`
}

func (r *CreateRabbitMQUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQVipInstanceRequestParams struct {
	// 可用区
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 私有网络ID，形如 vpc-xxx。有效的 VpcId 可通过登录[私有网络](https://console.cloud.tencent.com/vpc/vpc?rid=1)控制台查询；也可以调用接口 [DescribeVpcEx](https://cloud.tencent.com/document/api/215/1372)，从接口返回中的 unVpcId 字段获取。若在创建子机时 VpcId 与 SubnetId 同时传入 DEFAULT，则强制使用默认 vpc 网络。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络子网 ID，形如 subnet-xxx。有效的私有网络子网 ID 可通过登录[子网控制台](https://console.cloud.tencent.com/vpc/subnet?rid=1)查询；也可以调用接口 [DescribeSubnets](https://cloud.tencent.com/document/api/215/15784)，从接口返回中的 unSubnetId 字段获取。若在创建子机时 SubnetId 与 VpcId 同时传入 DEFAULT，则强制使用默认 vpc 网络。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 集群的节点规格，需要输入对应的规格标识：
	// 2C8G：rabbit-vip-basic-2c8g
	// 4C16G：rabbit-vip-basic-4c16g
	// 8C32G：rabbit-vip-basic-8c32g
	// 16C32G：rabbit-vip-basic-4
	// 16C64G：rabbit-vip-basic-16c64g
	// 2C4G：rabbit-vip-basic-5
	// 4C8G：rabbit-vip-basic-1
	// 8C16G（已售罄）：rabbit-vip-basic-2
	// 不传默认为 4C8G：rabbit-vip-basic-1
	NodeSpec *string `json:"NodeSpec,omitnil,omitempty" name:"NodeSpec"`

	// 节点数量,多可用区最少为3节点。不传默认单可用区为1,多可用区为3
	NodeNum *int64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 单节点存储规格,不传默认为200G
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 是否开启默认镜像队列，true 表示为开启，false 表示为不开启。不传默认为 false
	EnableCreateDefaultHaMirrorQueue *bool `json:"EnableCreateDefaultHaMirrorQueue,omitnil,omitempty" name:"EnableCreateDefaultHaMirrorQueue"`

	// 仅预付费集群（PayMode 参数为 1 时）使用该参数，表示是否自动续费，true 表示打开自动续费。不传默认为 true
	AutoRenewFlag *bool `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 购买时长,不传默认为1(月)
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 付费方式，0 为后付费，即按量计费；1 为预付费，即包年包月。默认包年包月
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 集群版本，不传默认为 3.8.30，可选值为 3.8.30 和 3.11.8
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`

	// 是否国际站请求，默认 false
	IsIntl *bool `json:"IsIntl,omitnil,omitempty" name:"IsIntl"`

	// 资源标签列表
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 公网带宽大小，单位 Mbps
	Bandwidth *uint64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 是否打开公网接入，不传默认为false
	EnablePublicAccess *bool `json:"EnablePublicAccess,omitnil,omitempty" name:"EnablePublicAccess"`

	// 是否打开集群删除保护，不传默认为 false
	EnableDeletionProtection *bool `json:"EnableDeletionProtection,omitnil,omitempty" name:"EnableDeletionProtection"`
}

type CreateRabbitMQVipInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 可用区
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 私有网络ID，形如 vpc-xxx。有效的 VpcId 可通过登录[私有网络](https://console.cloud.tencent.com/vpc/vpc?rid=1)控制台查询；也可以调用接口 [DescribeVpcEx](https://cloud.tencent.com/document/api/215/1372)，从接口返回中的 unVpcId 字段获取。若在创建子机时 VpcId 与 SubnetId 同时传入 DEFAULT，则强制使用默认 vpc 网络。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络子网 ID，形如 subnet-xxx。有效的私有网络子网 ID 可通过登录[子网控制台](https://console.cloud.tencent.com/vpc/subnet?rid=1)查询；也可以调用接口 [DescribeSubnets](https://cloud.tencent.com/document/api/215/15784)，从接口返回中的 unSubnetId 字段获取。若在创建子机时 SubnetId 与 VpcId 同时传入 DEFAULT，则强制使用默认 vpc 网络。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 集群的节点规格，需要输入对应的规格标识：
	// 2C8G：rabbit-vip-basic-2c8g
	// 4C16G：rabbit-vip-basic-4c16g
	// 8C32G：rabbit-vip-basic-8c32g
	// 16C32G：rabbit-vip-basic-4
	// 16C64G：rabbit-vip-basic-16c64g
	// 2C4G：rabbit-vip-basic-5
	// 4C8G：rabbit-vip-basic-1
	// 8C16G（已售罄）：rabbit-vip-basic-2
	// 不传默认为 4C8G：rabbit-vip-basic-1
	NodeSpec *string `json:"NodeSpec,omitnil,omitempty" name:"NodeSpec"`

	// 节点数量,多可用区最少为3节点。不传默认单可用区为1,多可用区为3
	NodeNum *int64 `json:"NodeNum,omitnil,omitempty" name:"NodeNum"`

	// 单节点存储规格,不传默认为200G
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 是否开启默认镜像队列，true 表示为开启，false 表示为不开启。不传默认为 false
	EnableCreateDefaultHaMirrorQueue *bool `json:"EnableCreateDefaultHaMirrorQueue,omitnil,omitempty" name:"EnableCreateDefaultHaMirrorQueue"`

	// 仅预付费集群（PayMode 参数为 1 时）使用该参数，表示是否自动续费，true 表示打开自动续费。不传默认为 true
	AutoRenewFlag *bool `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 购买时长,不传默认为1(月)
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 付费方式，0 为后付费，即按量计费；1 为预付费，即包年包月。默认包年包月
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 集群版本，不传默认为 3.8.30，可选值为 3.8.30 和 3.11.8
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`

	// 是否国际站请求，默认 false
	IsIntl *bool `json:"IsIntl,omitnil,omitempty" name:"IsIntl"`

	// 资源标签列表
	ResourceTags []*Tag `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`

	// 公网带宽大小，单位 Mbps
	Bandwidth *uint64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 是否打开公网接入，不传默认为false
	EnablePublicAccess *bool `json:"EnablePublicAccess,omitnil,omitempty" name:"EnablePublicAccess"`

	// 是否打开集群删除保护，不传默认为 false
	EnableDeletionProtection *bool `json:"EnableDeletionProtection,omitnil,omitempty" name:"EnableDeletionProtection"`
}

func (r *CreateRabbitMQVipInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQVipInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneIds")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ClusterName")
	delete(f, "NodeSpec")
	delete(f, "NodeNum")
	delete(f, "StorageSize")
	delete(f, "EnableCreateDefaultHaMirrorQueue")
	delete(f, "AutoRenewFlag")
	delete(f, "TimeSpan")
	delete(f, "PayMode")
	delete(f, "ClusterVersion")
	delete(f, "IsIntl")
	delete(f, "ResourceTags")
	delete(f, "Bandwidth")
	delete(f, "EnablePublicAccess")
	delete(f, "EnableDeletionProtection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRabbitMQVipInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQVipInstanceResponseParams struct {
	// 订单号 ID
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// 实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRabbitMQVipInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateRabbitMQVipInstanceResponseParams `json:"Response"`
}

func (r *CreateRabbitMQVipInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQVipInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQVirtualHostRequestParams struct {
	// 实例 ID，形如amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost名
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 消息轨迹开关,true打开,false关闭,默认关闭
	TraceFlag *bool `json:"TraceFlag,omitnil,omitempty" name:"TraceFlag"`

	// 是否创建镜像队列策略，默认值 true
	MirrorQueuePolicyFlag *bool `json:"MirrorQueuePolicyFlag,omitnil,omitempty" name:"MirrorQueuePolicyFlag"`
}

type CreateRabbitMQVirtualHostRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost名
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 消息轨迹开关,true打开,false关闭,默认关闭
	TraceFlag *bool `json:"TraceFlag,omitnil,omitempty" name:"TraceFlag"`

	// 是否创建镜像队列策略，默认值 true
	MirrorQueuePolicyFlag *bool `json:"MirrorQueuePolicyFlag,omitnil,omitempty" name:"MirrorQueuePolicyFlag"`
}

func (r *CreateRabbitMQVirtualHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQVirtualHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "Description")
	delete(f, "TraceFlag")
	delete(f, "MirrorQueuePolicyFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRabbitMQVirtualHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQVirtualHostResponseParams struct {
	// vhost名
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRabbitMQVirtualHostResponse struct {
	*tchttp.BaseResponse
	Response *CreateRabbitMQVirtualHostResponseParams `json:"Response"`
}

func (r *CreateRabbitMQVirtualHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQVirtualHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRocketMQClusterRequestParams struct {
	// 集群名称，3-64个字符，只能包含字母、数字、“-”及“_”
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 集群描述，128个字符以内
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type CreateRocketMQClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群名称，3-64个字符，只能包含字母、数字、“-”及“_”
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 集群描述，128个字符以内
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`
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
	delete(f, "TagList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRocketMQClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRocketMQClusterResponseParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRocketMQClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateRocketMQClusterResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateRocketMQEnvironmentRoleRequestParams struct {
	// 命名空间
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 角色名称。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 授权项，最多只能包含produce、consume两项的非空字符串数组。
	Permissions []*string `json:"Permissions,omitnil,omitempty" name:"Permissions"`

	// 必填字段，集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Topic&Group维度权限配置
	DetailedPerms []*DetailedRolePerm `json:"DetailedPerms,omitnil,omitempty" name:"DetailedPerms"`
}

type CreateRocketMQEnvironmentRoleRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 角色名称。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 授权项，最多只能包含produce、consume两项的非空字符串数组。
	Permissions []*string `json:"Permissions,omitnil,omitempty" name:"Permissions"`

	// 必填字段，集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Topic&Group维度权限配置
	DetailedPerms []*DetailedRolePerm `json:"DetailedPerms,omitnil,omitempty" name:"DetailedPerms"`
}

func (r *CreateRocketMQEnvironmentRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQEnvironmentRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "RoleName")
	delete(f, "Permissions")
	delete(f, "ClusterId")
	delete(f, "DetailedPerms")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRocketMQEnvironmentRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRocketMQEnvironmentRoleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRocketMQEnvironmentRoleResponse struct {
	*tchttp.BaseResponse
	Response *CreateRocketMQEnvironmentRoleResponseParams `json:"Response"`
}

func (r *CreateRocketMQEnvironmentRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQEnvironmentRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRocketMQGroupRequestParams struct {
	// Group名称，8~64个字符
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 命名空间，目前只支持单个命名空间
	Namespaces []*string `json:"Namespaces,omitnil,omitempty" name:"Namespaces"`

	// 是否开启消费
	ReadEnable *bool `json:"ReadEnable,omitnil,omitempty" name:"ReadEnable"`

	// 是否开启广播消费
	BroadcastEnable *bool `json:"BroadcastEnable,omitnil,omitempty" name:"BroadcastEnable"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 说明信息，最长128个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Group类型（TCP/HTTP）
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// Group最大重试次数
	RetryMaxTimes *uint64 `json:"RetryMaxTimes,omitnil,omitempty" name:"RetryMaxTimes"`
}

type CreateRocketMQGroupRequest struct {
	*tchttp.BaseRequest
	
	// Group名称，8~64个字符
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 命名空间，目前只支持单个命名空间
	Namespaces []*string `json:"Namespaces,omitnil,omitempty" name:"Namespaces"`

	// 是否开启消费
	ReadEnable *bool `json:"ReadEnable,omitnil,omitempty" name:"ReadEnable"`

	// 是否开启广播消费
	BroadcastEnable *bool `json:"BroadcastEnable,omitnil,omitempty" name:"BroadcastEnable"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 说明信息，最长128个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Group类型（TCP/HTTP）
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// Group最大重试次数
	RetryMaxTimes *uint64 `json:"RetryMaxTimes,omitnil,omitempty" name:"RetryMaxTimes"`
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
	delete(f, "GroupType")
	delete(f, "RetryMaxTimes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRocketMQGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRocketMQGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRocketMQGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateRocketMQGroupResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateRocketMQNamespaceRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称，3-64个字符，只能包含字母、数字、“-”及“_”
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 已废弃
	Ttl *uint64 `json:"Ttl,omitnil,omitempty" name:"Ttl"`

	// 已废弃
	RetentionTime *uint64 `json:"RetentionTime,omitnil,omitempty" name:"RetentionTime"`

	// 说明，最大128个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateRocketMQNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称，3-64个字符，只能包含字母、数字、“-”及“_”
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 已废弃
	Ttl *uint64 `json:"Ttl,omitnil,omitempty" name:"Ttl"`

	// 已废弃
	RetentionTime *uint64 `json:"RetentionTime,omitnil,omitempty" name:"RetentionTime"`

	// 说明，最大128个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
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

// Predefined struct for user
type CreateRocketMQNamespaceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRocketMQNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateRocketMQNamespaceResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateRocketMQRoleRequestParams struct {
	// 角色名称，不支持中字以及除了短线和下划线外的特殊字符且长度必须大于0且小等于32。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备注说明，长度必须大等于0且小等于128。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 角色授权类型（集群：Cluster; 主题或消费组：TopicAndGroup）
	PermType *string `json:"PermType,omitnil,omitempty" name:"PermType"`
}

type CreateRocketMQRoleRequest struct {
	*tchttp.BaseRequest
	
	// 角色名称，不支持中字以及除了短线和下划线外的特殊字符且长度必须大于0且小等于32。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备注说明，长度必须大等于0且小等于128。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 角色授权类型（集群：Cluster; 主题或消费组：TopicAndGroup）
	PermType *string `json:"PermType,omitnil,omitempty" name:"PermType"`
}

func (r *CreateRocketMQRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleName")
	delete(f, "ClusterId")
	delete(f, "Remark")
	delete(f, "PermType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRocketMQRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRocketMQRoleResponseParams struct {
	// 角色名称
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 角色token
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 备注说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRocketMQRoleResponse struct {
	*tchttp.BaseResponse
	Response *CreateRocketMQRoleResponseParams `json:"Response"`
}

func (r *CreateRocketMQRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRocketMQTopicRequestParams struct {
	// 主题名称，3-64个字符，只能包含字母、数字、“-”及“_”
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 主题所在的命名空间，目前支持在单个命名空间下创建主题
	Namespaces []*string `json:"Namespaces,omitnil,omitempty" name:"Namespaces"`

	// 主题类型，可选值为Normal, GlobalOrder, PartitionedOrder, Transaction, DelayScheduled。Transaction仅在专享版支持。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 主题说明，最大128个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 分区数，全局顺序无效
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`
}

type CreateRocketMQTopicRequest struct {
	*tchttp.BaseRequest
	
	// 主题名称，3-64个字符，只能包含字母、数字、“-”及“_”
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 主题所在的命名空间，目前支持在单个命名空间下创建主题
	Namespaces []*string `json:"Namespaces,omitnil,omitempty" name:"Namespaces"`

	// 主题类型，可选值为Normal, GlobalOrder, PartitionedOrder, Transaction, DelayScheduled。Transaction仅在专享版支持。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 主题说明，最大128个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 分区数，全局顺序无效
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`
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

// Predefined struct for user
type CreateRocketMQTopicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRocketMQTopicResponse struct {
	*tchttp.BaseResponse
	Response *CreateRocketMQTopicResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateRocketMQVipInstanceRequestParams struct {
	// 实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 集群规格，支持规格有 1.通用型:rocket-vip-basic-0; 2.基础型:rocket-vip-basic-1; 3.标准型:rocket-vip-basic-2; 4.高阶Ⅰ型:rocket-vip-basic-3; 5.高阶Ⅱ型:rocket-vip-basic-4
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// 节点数量，最小2，最大20
	NodeCount *int64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 单节点存储空间，GB为单位，最低200GB
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 节点部署的区域ID列表，如广州一区，则是100001，具体可查询腾讯云官网
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// VPC信息
	VpcInfo *VpcInfo `json:"VpcInfo,omitnil,omitempty" name:"VpcInfo"`

	// 购买时长，月为单位
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 是否用于迁移上云，默认为false
	SupportsMigrateToCloud *bool `json:"SupportsMigrateToCloud,omitnil,omitempty" name:"SupportsMigrateToCloud"`

	// 是否开启公网
	EnablePublic *bool `json:"EnablePublic,omitnil,omitempty" name:"EnablePublic"`

	// 公网带宽，在开启公网情况下为必传字段
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网白名单
	IpRules []*PublicAccessRule `json:"IpRules,omitnil,omitempty" name:"IpRules"`

	// 标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateRocketMQVipInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 集群规格，支持规格有 1.通用型:rocket-vip-basic-0; 2.基础型:rocket-vip-basic-1; 3.标准型:rocket-vip-basic-2; 4.高阶Ⅰ型:rocket-vip-basic-3; 5.高阶Ⅱ型:rocket-vip-basic-4
	Spec *string `json:"Spec,omitnil,omitempty" name:"Spec"`

	// 节点数量，最小2，最大20
	NodeCount *int64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 单节点存储空间，GB为单位，最低200GB
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 节点部署的区域ID列表，如广州一区，则是100001，具体可查询腾讯云官网
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// VPC信息
	VpcInfo *VpcInfo `json:"VpcInfo,omitnil,omitempty" name:"VpcInfo"`

	// 购买时长，月为单位
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 是否用于迁移上云，默认为false
	SupportsMigrateToCloud *bool `json:"SupportsMigrateToCloud,omitnil,omitempty" name:"SupportsMigrateToCloud"`

	// 是否开启公网
	EnablePublic *bool `json:"EnablePublic,omitnil,omitempty" name:"EnablePublic"`

	// 公网带宽，在开启公网情况下为必传字段
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网白名单
	IpRules []*PublicAccessRule `json:"IpRules,omitnil,omitempty" name:"IpRules"`

	// 标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateRocketMQVipInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQVipInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Spec")
	delete(f, "NodeCount")
	delete(f, "StorageSize")
	delete(f, "ZoneIds")
	delete(f, "VpcInfo")
	delete(f, "TimeSpan")
	delete(f, "SupportsMigrateToCloud")
	delete(f, "EnablePublic")
	delete(f, "Bandwidth")
	delete(f, "IpRules")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRocketMQVipInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRocketMQVipInstanceResponseParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRocketMQVipInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateRocketMQVipInstanceResponseParams `json:"Response"`
}

func (r *CreateRocketMQVipInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQVipInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoleRequestParams struct {
	// 角色名称，不支持中字以及除了短线和下划线外的特殊字符且长度必须大于0且小等于32。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 备注说明，长度必须大等于0且小等于128。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type CreateRoleRequest struct {
	*tchttp.BaseRequest
	
	// 角色名称，不支持中字以及除了短线和下划线外的特殊字符且长度必须大于0且小等于32。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 备注说明，长度必须大等于0且小等于128。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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

// Predefined struct for user
type CreateRoleResponseParams struct {
	// 角色名称
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 角色token
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 备注说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 批量绑定名字空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentRoleSets []*EnvironmentRoleSet `json:"EnvironmentRoleSets,omitnil,omitempty" name:"EnvironmentRoleSets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRoleResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoleResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateSubscriptionRequestParams struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 订阅者名称，不超过128个字符。
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`

	// 是否幂等创建，若否不允许创建同名的订阅关系。
	IsIdempotent *bool `json:"IsIdempotent,omitnil,omitempty" name:"IsIdempotent"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备注，128个字符以内。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否自动创建死信和重试主题，True 表示创建，False表示不创建，默认自动创建死信和重试主题。
	AutoCreatePolicyTopic *bool `json:"AutoCreatePolicyTopic,omitnil,omitempty" name:"AutoCreatePolicyTopic"`

	// 指定死信和重试主题名称规范，LEGACY表示历史命名规则，COMMUNITY表示Pulsar社区命名规范
	PostFixPattern *string `json:"PostFixPattern,omitnil,omitempty" name:"PostFixPattern"`
}

type CreateSubscriptionRequest struct {
	*tchttp.BaseRequest
	
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 订阅者名称，不超过128个字符。
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`

	// 是否幂等创建，若否不允许创建同名的订阅关系。
	IsIdempotent *bool `json:"IsIdempotent,omitnil,omitempty" name:"IsIdempotent"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备注，128个字符以内。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否自动创建死信和重试主题，True 表示创建，False表示不创建，默认自动创建死信和重试主题。
	AutoCreatePolicyTopic *bool `json:"AutoCreatePolicyTopic,omitnil,omitempty" name:"AutoCreatePolicyTopic"`

	// 指定死信和重试主题名称规范，LEGACY表示历史命名规则，COMMUNITY表示Pulsar社区命名规范
	PostFixPattern *string `json:"PostFixPattern,omitnil,omitempty" name:"PostFixPattern"`
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
	delete(f, "ClusterId")
	delete(f, "Remark")
	delete(f, "AutoCreatePolicyTopic")
	delete(f, "PostFixPattern")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubscriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubscriptionResponseParams struct {
	// 创建结果。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSubscriptionResponse struct {
	*tchttp.BaseResponse
	Response *CreateSubscriptionResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateTopicRequestParams struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名，不支持中字以及除了短线和下划线外的特殊字符且不超过64个字符。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 入参为1，即是创建非分区topic，无分区；入参大于1，表示分区topic的分区数，最大不允许超过32。
	Partitions *uint64 `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备注，128字符以内。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 该入参将逐步弃用，可切换至PulsarTopicType参数
	// 0： 普通消息；
	// 1 ：全局顺序消息；
	// 2 ：局部顺序消息；
	// 3 ：重试队列；
	// 4 ：死信队列。
	TopicType *uint64 `json:"TopicType,omitnil,omitempty" name:"TopicType"`

	// Pulsar 主题类型
	// 0: 非持久非分区
	// 1: 非持久分区
	// 2: 持久非分区
	// 3: 持久分区
	PulsarTopicType *int64 `json:"PulsarTopicType,omitnil,omitempty" name:"PulsarTopicType"`

	// 未消费消息过期时间，单位：秒，取值范围：60秒~15天。
	MsgTTL *uint64 `json:"MsgTTL,omitnil,omitempty" name:"MsgTTL"`

	// 不传默认是原生策略，DefaultPolicy表示当订阅下达到最大未确认消息数 5000 时，服务端将不再向当前订阅下的所有消费者推送消息，DynamicPolicy表示动态调整订阅下的最大未确认消息数，具体配额是在 5000 和消费者数量*20之间取最大值。每个消费者默认最大 unack 消息数为 20，超过该限制时仅影响该消费者，不影响其他消费者。
	UnackPolicy *string `json:"UnackPolicy,omitnil,omitempty" name:"UnackPolicy"`

	// 是否开启异常消费者隔离
	IsolateConsumerEnable *bool `json:"IsolateConsumerEnable,omitnil,omitempty" name:"IsolateConsumerEnable"`

	// 消费者 Ack 超时时间，单位：秒，范围60-（3600*24）
	AckTimeOut *int64 `json:"AckTimeOut,omitnil,omitempty" name:"AckTimeOut"`
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest
	
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名，不支持中字以及除了短线和下划线外的特殊字符且不超过64个字符。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 入参为1，即是创建非分区topic，无分区；入参大于1，表示分区topic的分区数，最大不允许超过32。
	Partitions *uint64 `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备注，128字符以内。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 该入参将逐步弃用，可切换至PulsarTopicType参数
	// 0： 普通消息；
	// 1 ：全局顺序消息；
	// 2 ：局部顺序消息；
	// 3 ：重试队列；
	// 4 ：死信队列。
	TopicType *uint64 `json:"TopicType,omitnil,omitempty" name:"TopicType"`

	// Pulsar 主题类型
	// 0: 非持久非分区
	// 1: 非持久分区
	// 2: 持久非分区
	// 3: 持久分区
	PulsarTopicType *int64 `json:"PulsarTopicType,omitnil,omitempty" name:"PulsarTopicType"`

	// 未消费消息过期时间，单位：秒，取值范围：60秒~15天。
	MsgTTL *uint64 `json:"MsgTTL,omitnil,omitempty" name:"MsgTTL"`

	// 不传默认是原生策略，DefaultPolicy表示当订阅下达到最大未确认消息数 5000 时，服务端将不再向当前订阅下的所有消费者推送消息，DynamicPolicy表示动态调整订阅下的最大未确认消息数，具体配额是在 5000 和消费者数量*20之间取最大值。每个消费者默认最大 unack 消息数为 20，超过该限制时仅影响该消费者，不影响其他消费者。
	UnackPolicy *string `json:"UnackPolicy,omitnil,omitempty" name:"UnackPolicy"`

	// 是否开启异常消费者隔离
	IsolateConsumerEnable *bool `json:"IsolateConsumerEnable,omitnil,omitempty" name:"IsolateConsumerEnable"`

	// 消费者 Ack 超时时间，单位：秒，范围60-（3600*24）
	AckTimeOut *int64 `json:"AckTimeOut,omitnil,omitempty" name:"AckTimeOut"`
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
	delete(f, "ClusterId")
	delete(f, "Remark")
	delete(f, "TopicType")
	delete(f, "PulsarTopicType")
	delete(f, "MsgTTL")
	delete(f, "UnackPolicy")
	delete(f, "IsolateConsumerEnable")
	delete(f, "AckTimeOut")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicResponseParams struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 0或1：非分区topic，无分区；大于1：具体分区topic的分区数。（存量非分区主题返回0，增量非分区主题返回1）
	Partitions *uint64 `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 备注，128字符以内。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 0： 普通消息；
	// 1 ：全局顺序消息；
	// 2 ：局部顺序消息；
	// 3 ：重试队列；
	// 4 ：死信队列；
	TopicType *uint64 `json:"TopicType,omitnil,omitempty" name:"TopicType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse
	Response *CreateTopicResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteClusterRequestParams struct {
	// 集群Id，传入需要删除的集群Id。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id，传入需要删除的集群Id。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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

// Predefined struct for user
type DeleteClusterResponseParams struct {
	// 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteClusterResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteCmqQueueRequestParams struct {
	// 队列名字，在单个地域同一账号下唯一。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`
}

type DeleteCmqQueueRequest struct {
	*tchttp.BaseRequest
	
	// 队列名字，在单个地域同一账号下唯一。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`
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

// Predefined struct for user
type DeleteCmqQueueResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCmqQueueResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCmqQueueResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteCmqSubscribeRequestParams struct {
	// 主题名字，在单个地域同一账号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 订阅名字，在单个地域同一账号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`
}

type DeleteCmqSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// 主题名字，在单个地域同一账号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 订阅名字，在单个地域同一账号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`
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

// Predefined struct for user
type DeleteCmqSubscribeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCmqSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCmqSubscribeResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteCmqTopicRequestParams struct {
	// 主题名字，在单个地域同一账号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type DeleteCmqTopicRequest struct {
	*tchttp.BaseRequest
	
	// 主题名字，在单个地域同一账号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
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

// Predefined struct for user
type DeleteCmqTopicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCmqTopicResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCmqTopicResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteEnvironmentRolesRequestParams struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 角色名称数组。
	RoleNames []*string `json:"RoleNames,omitnil,omitempty" name:"RoleNames"`

	// 必填字段，集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DeleteEnvironmentRolesRequest struct {
	*tchttp.BaseRequest
	
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 角色名称数组。
	RoleNames []*string `json:"RoleNames,omitnil,omitempty" name:"RoleNames"`

	// 必填字段，集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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

// Predefined struct for user
type DeleteEnvironmentRolesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEnvironmentRolesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEnvironmentRolesResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteEnvironmentsRequestParams struct {
	// 环境（命名空间）数组，每次最多删除20个。
	EnvironmentIds []*string `json:"EnvironmentIds,omitnil,omitempty" name:"EnvironmentIds"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DeleteEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// 环境（命名空间）数组，每次最多删除20个。
	EnvironmentIds []*string `json:"EnvironmentIds,omitnil,omitempty" name:"EnvironmentIds"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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

// Predefined struct for user
type DeleteEnvironmentsResponseParams struct {
	// 成功删除的环境（命名空间）数组。
	EnvironmentIds []*string `json:"EnvironmentIds,omitnil,omitempty" name:"EnvironmentIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEnvironmentsResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteProClusterRequestParams struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DeleteProClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DeleteProClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProClusterResponseParams struct {
	// 退还实例订单号
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteProClusterResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProClusterResponseParams `json:"Response"`
}

func (r *DeleteProClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQBindingRequestParams struct {
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到Vhost名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 路由关系Id
	BindingId *int64 `json:"BindingId,omitnil,omitempty" name:"BindingId"`
}

type DeleteRabbitMQBindingRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到Vhost名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 路由关系Id
	BindingId *int64 `json:"BindingId,omitnil,omitempty" name:"BindingId"`
}

func (r *DeleteRabbitMQBindingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQBindingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "BindingId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRabbitMQBindingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQBindingResponseParams struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 路由关系Id
	BindingId *int64 `json:"BindingId,omitnil,omitempty" name:"BindingId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRabbitMQBindingResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRabbitMQBindingResponseParams `json:"Response"`
}

func (r *DeleteRabbitMQBindingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQBindingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQPermissionRequestParams struct {
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，形如 admin。有效的 User 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，点击集群列表中的集群，进入集群详情，并在用户与权限页签中找到用户列表，从而找到用户名称。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到 Vhost 名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`
}

type DeleteRabbitMQPermissionRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，形如 admin。有效的 User 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，点击集群列表中的集群，进入集群详情，并在用户与权限页签中找到用户列表，从而找到用户名称。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到 Vhost 名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`
}

func (r *DeleteRabbitMQPermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQPermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "VirtualHost")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRabbitMQPermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQPermissionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRabbitMQPermissionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRabbitMQPermissionResponseParams `json:"Response"`
}

func (r *DeleteRabbitMQPermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQPermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQUserRequestParams struct {
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，形如 admin。有效的 User 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，点击集群列表中的集群，进入集群详情，并在用户与权限页签中找到用户列表，从而找到用户名称。
	User *string `json:"User,omitnil,omitempty" name:"User"`
}

type DeleteRabbitMQUserRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，形如 admin。有效的 User 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，点击集群列表中的集群，进入集群详情，并在用户与权限页签中找到用户列表，从而找到用户名称。
	User *string `json:"User,omitnil,omitempty" name:"User"`
}

func (r *DeleteRabbitMQUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRabbitMQUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRabbitMQUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRabbitMQUserResponseParams `json:"Response"`
}

func (r *DeleteRabbitMQUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQVipInstanceRequestParams struct {
	// 实例 ID，形如amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否国际站请求，默认 false
	IsIntl *bool `json:"IsIntl,omitnil,omitempty" name:"IsIntl"`
}

type DeleteRabbitMQVipInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否国际站请求，默认 false
	IsIntl *bool `json:"IsIntl,omitnil,omitempty" name:"IsIntl"`
}

func (r *DeleteRabbitMQVipInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQVipInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "IsIntl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRabbitMQVipInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQVipInstanceResponseParams struct {
	// 订单号 ID
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// 实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRabbitMQVipInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRabbitMQVipInstanceResponseParams `json:"Response"`
}

func (r *DeleteRabbitMQVipInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQVipInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQVirtualHostRequestParams struct {
	// 实例 ID，形如amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到Vhost名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`
}

type DeleteRabbitMQVirtualHostRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到Vhost名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`
}

func (r *DeleteRabbitMQVirtualHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQVirtualHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRabbitMQVirtualHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQVirtualHostResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRabbitMQVirtualHostResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRabbitMQVirtualHostResponseParams `json:"Response"`
}

func (r *DeleteRabbitMQVirtualHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQVirtualHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRocketMQClusterRequestParams struct {
	// 待删除的集群Id。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DeleteRocketMQClusterRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的集群Id。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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

// Predefined struct for user
type DeleteRocketMQClusterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRocketMQClusterResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRocketMQClusterResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteRocketMQEnvironmentRolesRequestParams struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 角色名称数组。
	RoleNames []*string `json:"RoleNames,omitnil,omitempty" name:"RoleNames"`

	// 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DeleteRocketMQEnvironmentRolesRequest struct {
	*tchttp.BaseRequest
	
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 角色名称数组。
	RoleNames []*string `json:"RoleNames,omitnil,omitempty" name:"RoleNames"`

	// 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DeleteRocketMQEnvironmentRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQEnvironmentRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "RoleNames")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRocketMQEnvironmentRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRocketMQEnvironmentRolesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRocketMQEnvironmentRolesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRocketMQEnvironmentRolesResponseParams `json:"Response"`
}

func (r *DeleteRocketMQEnvironmentRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQEnvironmentRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRocketMQGroupRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 消费组名称
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DeleteRocketMQGroupRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 消费组名称
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
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

// Predefined struct for user
type DeleteRocketMQGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRocketMQGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRocketMQGroupResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteRocketMQNamespaceRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`
}

type DeleteRocketMQNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`
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

// Predefined struct for user
type DeleteRocketMQNamespaceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRocketMQNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRocketMQNamespaceResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteRocketMQRolesRequestParams struct {
	// 角色名称数组。
	RoleNames []*string `json:"RoleNames,omitnil,omitempty" name:"RoleNames"`

	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DeleteRocketMQRolesRequest struct {
	*tchttp.BaseRequest
	
	// 角色名称数组。
	RoleNames []*string `json:"RoleNames,omitnil,omitempty" name:"RoleNames"`

	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DeleteRocketMQRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleNames")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRocketMQRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRocketMQRolesResponseParams struct {
	// 成功删除的角色名称数组。
	RoleNames []*string `json:"RoleNames,omitnil,omitempty" name:"RoleNames"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRocketMQRolesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRocketMQRolesResponseParams `json:"Response"`
}

func (r *DeleteRocketMQRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRocketMQTopicRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
}

type DeleteRocketMQTopicRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
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

// Predefined struct for user
type DeleteRocketMQTopicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRocketMQTopicResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRocketMQTopicResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteRocketMQVipInstanceRequestParams struct {
	// 实例的集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DeleteRocketMQVipInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例的集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DeleteRocketMQVipInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQVipInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRocketMQVipInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRocketMQVipInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRocketMQVipInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRocketMQVipInstanceResponseParams `json:"Response"`
}

func (r *DeleteRocketMQVipInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQVipInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRolesRequestParams struct {
	// 角色名称数组。
	RoleNames []*string `json:"RoleNames,omitnil,omitempty" name:"RoleNames"`

	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DeleteRolesRequest struct {
	*tchttp.BaseRequest
	
	// 角色名称数组。
	RoleNames []*string `json:"RoleNames,omitnil,omitempty" name:"RoleNames"`

	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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

// Predefined struct for user
type DeleteRolesResponseParams struct {
	// 成功删除的角色名称数组。
	RoleNames []*string `json:"RoleNames,omitnil,omitempty" name:"RoleNames"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRolesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRolesResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteSubscriptionsRequestParams struct {
	// 订阅关系集合，每次最多删除20个。
	SubscriptionTopicSets []*SubscriptionTopic `json:"SubscriptionTopicSets,omitnil,omitempty" name:"SubscriptionTopicSets"`

	// pulsar集群Id。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 是否强制删除，默认为false
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

type DeleteSubscriptionsRequest struct {
	*tchttp.BaseRequest
	
	// 订阅关系集合，每次最多删除20个。
	SubscriptionTopicSets []*SubscriptionTopic `json:"SubscriptionTopicSets,omitnil,omitempty" name:"SubscriptionTopicSets"`

	// pulsar集群Id。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 是否强制删除，默认为false
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
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

// Predefined struct for user
type DeleteSubscriptionsResponseParams struct {
	// 成功删除的订阅关系数组。
	SubscriptionTopicSets []*SubscriptionTopic `json:"SubscriptionTopicSets,omitnil,omitempty" name:"SubscriptionTopicSets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSubscriptionsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSubscriptionsResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteTopicsRequestParams struct {
	// 主题集合，每次最多删除20个。
	TopicSets []*TopicRecord `json:"TopicSets,omitnil,omitempty" name:"TopicSets"`

	// pulsar集群Id。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 是否强制删除，默认为false
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

type DeleteTopicsRequest struct {
	*tchttp.BaseRequest
	
	// 主题集合，每次最多删除20个。
	TopicSets []*TopicRecord `json:"TopicSets,omitnil,omitempty" name:"TopicSets"`

	// pulsar集群Id。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 是否强制删除，默认为false
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
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

// Predefined struct for user
type DeleteTopicsResponseParams struct {
	// 被删除的主题数组。
	TopicSets []*TopicRecord `json:"TopicSets,omitnil,omitempty" name:"TopicSets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTopicsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTopicsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAMQPClustersRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 按照集群ID关键字搜索
	IdKeyword *string `json:"IdKeyword,omitnil,omitempty" name:"IdKeyword"`

	// 按照集群名称关键字搜索
	NameKeyword *string `json:"NameKeyword,omitnil,omitempty" name:"NameKeyword"`

	// 集群ID列表过滤
	ClusterIdList []*string `json:"ClusterIdList,omitnil,omitempty" name:"ClusterIdList"`

	// 标签过滤查找时，需要设置为true
	IsTagFilter *bool `json:"IsTagFilter,omitnil,omitempty" name:"IsTagFilter"`

	// 过滤器。目前支持按标签过滤。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAMQPClustersRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 按照集群ID关键字搜索
	IdKeyword *string `json:"IdKeyword,omitnil,omitempty" name:"IdKeyword"`

	// 按照集群名称关键字搜索
	NameKeyword *string `json:"NameKeyword,omitnil,omitempty" name:"NameKeyword"`

	// 集群ID列表过滤
	ClusterIdList []*string `json:"ClusterIdList,omitnil,omitempty" name:"ClusterIdList"`

	// 标签过滤查找时，需要设置为true
	IsTagFilter *bool `json:"IsTagFilter,omitnil,omitempty" name:"IsTagFilter"`

	// 过滤器。目前支持按标签过滤。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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

// Predefined struct for user
type DescribeAMQPClustersResponseParams struct {
	// 集群信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterList []*AMQPClusterDetail `json:"ClusterList,omitnil,omitempty" name:"ClusterList"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAMQPClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAMQPClustersResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeAllTenantsRequestParams struct {
	// 查询偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询限制条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 物理集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 虚拟集群ID
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 虚拟集群名称
	TenantName *string `json:"TenantName,omitnil,omitempty" name:"TenantName"`

	// 协议类型数组
	Types []*string `json:"Types,omitnil,omitempty" name:"Types"`

	// 排序字段名，支持createTime，updateTime
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 升序排列ASC，降序排列DESC
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

type DescribeAllTenantsRequest struct {
	*tchttp.BaseRequest
	
	// 查询偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询限制条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 物理集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 虚拟集群ID
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 虚拟集群名称
	TenantName *string `json:"TenantName,omitnil,omitempty" name:"TenantName"`

	// 协议类型数组
	Types []*string `json:"Types,omitnil,omitempty" name:"Types"`

	// 排序字段名，支持createTime，updateTime
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 升序排列ASC，降序排列DESC
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
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

// Predefined struct for user
type DescribeAllTenantsResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 虚拟集群列表
	Tenants []*InternalTenant `json:"Tenants,omitnil,omitempty" name:"Tenants"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllTenantsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllTenantsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBindClustersRequestParams struct {

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

// Predefined struct for user
type DescribeBindClustersResponseParams struct {
	// 专享集群的数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 专享集群的列表
	ClusterSet []*BindCluster `json:"ClusterSet,omitnil,omitempty" name:"ClusterSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBindClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBindClustersResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeBindVpcsRequestParams struct {
	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeBindVpcsRequest struct {
	*tchttp.BaseRequest
	
	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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

// Predefined struct for user
type DescribeBindVpcsResponseParams struct {
	// 记录数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Vpc集合。
	VpcSets []*VpcBindRecord `json:"VpcSets,omitnil,omitempty" name:"VpcSets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBindVpcsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBindVpcsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeClusterDetailRequestParams struct {
	// 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeClusterDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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

// Predefined struct for user
type DescribeClusterDetailResponseParams struct {
	// 集群的详细信息
	ClusterSet *Cluster `json:"ClusterSet,omitnil,omitempty" name:"ClusterSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeClustersRequestParams struct {
	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 集群ID列表过滤
	ClusterIdList []*string `json:"ClusterIdList,omitnil,omitempty" name:"ClusterIdList"`

	// 是否标签过滤
	IsTagFilter *bool `json:"IsTagFilter,omitnil,omitempty" name:"IsTagFilter"`

	// 过滤器。目前支持按标签过滤。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest
	
	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 集群ID列表过滤
	ClusterIdList []*string `json:"ClusterIdList,omitnil,omitempty" name:"ClusterIdList"`

	// 是否标签过滤
	IsTagFilter *bool `json:"IsTagFilter,omitnil,omitempty" name:"IsTagFilter"`

	// 过滤器。目前支持按标签过滤。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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

// Predefined struct for user
type DescribeClustersResponseParams struct {
	// 集群列表数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集群信息列表
	ClusterSet []*Cluster `json:"ClusterSet,omitnil,omitempty" name:"ClusterSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClustersResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCmqQueueDetailRequestParams struct {
	// 精确匹配QueueName
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`
}

type DescribeCmqQueueDetailRequest struct {
	*tchttp.BaseRequest
	
	// 精确匹配QueueName
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`
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

// Predefined struct for user
type DescribeCmqQueueDetailResponseParams struct {
	// 队列详情列表。
	QueueDescribe *CmqQueue `json:"QueueDescribe,omitnil,omitempty" name:"QueueDescribe"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCmqQueueDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCmqQueueDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCmqQueuesRequestParams struct {
	// 分页时本页获取队列列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页时本页获取队列的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据QueueName进行过滤
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// CMQ 队列名称列表过滤
	QueueNameList []*string `json:"QueueNameList,omitnil,omitempty" name:"QueueNameList"`

	// 标签过滤查找时，需要设置为 true
	IsTagFilter *bool `json:"IsTagFilter,omitnil,omitempty" name:"IsTagFilter"`

	// 过滤器。目前支持按标签过滤，标签的Name需要加前缀“tag:”，例如：tag:负责人、tag:环境、tag:业务
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeCmqQueuesRequest struct {
	*tchttp.BaseRequest
	
	// 分页时本页获取队列列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页时本页获取队列的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据QueueName进行过滤
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// CMQ 队列名称列表过滤
	QueueNameList []*string `json:"QueueNameList,omitnil,omitempty" name:"QueueNameList"`

	// 标签过滤查找时，需要设置为 true
	IsTagFilter *bool `json:"IsTagFilter,omitnil,omitempty" name:"IsTagFilter"`

	// 过滤器。目前支持按标签过滤，标签的Name需要加前缀“tag:”，例如：tag:负责人、tag:环境、tag:业务
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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

// Predefined struct for user
type DescribeCmqQueuesResponseParams struct {
	// 数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 队列列表
	QueueList []*CmqQueue `json:"QueueList,omitnil,omitempty" name:"QueueList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCmqQueuesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCmqQueuesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCmqSubscriptionDetailRequestParams struct {
	// 主题名字，在单个地域同一账号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 分页时本页获取主题列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页时本页获取主题的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据SubscriptionName进行模糊搜索
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`

	// 队列名称，订阅绑定的endpoint
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 查询类型。取值：（1）topic；（2）queue。
	// 默认值是topic。如果 queryType 是 topic，则查询主题下的订阅列表；如果 queryType 是 queue，则查询队列绑定的订阅列表。
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`
}

type DescribeCmqSubscriptionDetailRequest struct {
	*tchttp.BaseRequest
	
	// 主题名字，在单个地域同一账号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 分页时本页获取主题列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页时本页获取主题的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据SubscriptionName进行模糊搜索
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`

	// 队列名称，订阅绑定的endpoint
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 查询类型。取值：（1）topic；（2）queue。
	// 默认值是topic。如果 queryType 是 topic，则查询主题下的订阅列表；如果 queryType 是 queue，则查询队列绑定的订阅列表。
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`
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
	delete(f, "QueueName")
	delete(f, "QueryType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCmqSubscriptionDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCmqSubscriptionDetailResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Subscription属性集合
	SubscriptionSet []*CmqSubscription `json:"SubscriptionSet,omitnil,omitempty" name:"SubscriptionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCmqSubscriptionDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCmqSubscriptionDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCmqTopicDetailRequestParams struct {
	// 精确匹配TopicName。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type DescribeCmqTopicDetailRequest struct {
	*tchttp.BaseRequest
	
	// 精确匹配TopicName。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
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

// Predefined struct for user
type DescribeCmqTopicDetailResponseParams struct {
	// 主题详情
	TopicDescribe *CmqTopic `json:"TopicDescribe,omitnil,omitempty" name:"TopicDescribe"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCmqTopicDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCmqTopicDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCmqTopicsRequestParams struct {
	// 分页时本页获取队列列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页时本页获取队列的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据TopicName进行模糊搜索
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// CMQ 主题名称列表过滤
	TopicNameList []*string `json:"TopicNameList,omitnil,omitempty" name:"TopicNameList"`

	// 标签过滤查找时，需要设置为 true
	IsTagFilter *bool `json:"IsTagFilter,omitnil,omitempty" name:"IsTagFilter"`

	// 过滤器。目前支持按标签过滤，标签的Name需要加前缀“tag:”，例如：tag:负责人、tag:环境、tag:业务
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeCmqTopicsRequest struct {
	*tchttp.BaseRequest
	
	// 分页时本页获取队列列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页时本页获取队列的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据TopicName进行模糊搜索
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// CMQ 主题名称列表过滤
	TopicNameList []*string `json:"TopicNameList,omitnil,omitempty" name:"TopicNameList"`

	// 标签过滤查找时，需要设置为 true
	IsTagFilter *bool `json:"IsTagFilter,omitnil,omitempty" name:"IsTagFilter"`

	// 过滤器。目前支持按标签过滤，标签的Name需要加前缀“tag:”，例如：tag:负责人、tag:环境、tag:业务
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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

// Predefined struct for user
type DescribeCmqTopicsResponseParams struct {
	// 主题列表
	TopicList []*CmqTopic `json:"TopicList,omitnil,omitempty" name:"TopicList"`

	// 全量主题数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCmqTopicsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCmqTopicsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeEnvironmentAttributesRequestParams struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeEnvironmentAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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

// Predefined struct for user
type DescribeEnvironmentAttributesResponseParams struct {
	// 未消费消息过期时间，单位：秒，最大1296000（15天）。
	MsgTTL *uint64 `json:"MsgTTL,omitnil,omitempty" name:"MsgTTL"`

	// 消费速率限制，单位：byte/秒，0：不限速。
	RateInByte *uint64 `json:"RateInByte,omitnil,omitempty" name:"RateInByte"`

	// 消费速率限制，单位：个数/秒，0：不限速。
	RateInSize *uint64 `json:"RateInSize,omitnil,omitempty" name:"RateInSize"`

	// 已消费消息保存策略，单位：小时，0：消费完马上删除。
	RetentionHours *uint64 `json:"RetentionHours,omitnil,omitempty" name:"RetentionHours"`

	// 已消费消息保存策略，单位：G，0：消费完马上删除。
	RetentionSize *uint64 `json:"RetentionSize,omitnil,omitempty" name:"RetentionSize"`

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 副本数。
	Replicas *uint64 `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEnvironmentAttributesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvironmentAttributesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeEnvironmentRolesRequestParams struct {
	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 起始下标，不填默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 角色名称
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// * RoleName
	// 按照角色名进行过滤，精确查询。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeEnvironmentRolesRequest struct {
	*tchttp.BaseRequest
	
	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 起始下标，不填默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 角色名称
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// * RoleName
	// 按照角色名进行过滤，精确查询。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	delete(f, "ClusterId")
	delete(f, "EnvironmentId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "RoleName")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentRolesResponseParams struct {
	// 记录数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 命名空间角色集合。
	EnvironmentRoleSets []*EnvironmentRole `json:"EnvironmentRoleSets,omitnil,omitempty" name:"EnvironmentRoleSets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEnvironmentRolesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvironmentRolesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeEnvironmentsRequestParams struct {
	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称，模糊搜索。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// * EnvironmentId
	// 按照名称空间进行过滤，精确查询。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称，模糊搜索。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// * EnvironmentId
	// 按照名称空间进行过滤，精确查询。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	delete(f, "ClusterId")
	delete(f, "EnvironmentId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsResponseParams struct {
	// 命名空间记录数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 命名空间集合数组。
	EnvironmentSet []*Environment `json:"EnvironmentSet,omitnil,omitempty" name:"EnvironmentSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvironmentsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeMqMsgTraceRequestParams struct {
	// pulsar、rocketmq、rabbitmq、cmq
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 消息id
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 集群id，cmq为空
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间，cmq为空
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题，cmq为空，rocketmq查询死信时值为groupId
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// cmq必填，其他协议填空
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 消费组、订阅
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 查询死信时该值为true，只对Rocketmq有效
	QueryDlqMsg *bool `json:"QueryDlqMsg,omitnil,omitempty" name:"QueryDlqMsg"`

	// 生产时间
	ProduceTime *string `json:"ProduceTime,omitnil,omitempty" name:"ProduceTime"`
}

type DescribeMqMsgTraceRequest struct {
	*tchttp.BaseRequest
	
	// pulsar、rocketmq、rabbitmq、cmq
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 消息id
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 集群id，cmq为空
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间，cmq为空
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题，cmq为空，rocketmq查询死信时值为groupId
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// cmq必填，其他协议填空
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 消费组、订阅
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 查询死信时该值为true，只对Rocketmq有效
	QueryDlqMsg *bool `json:"QueryDlqMsg,omitnil,omitempty" name:"QueryDlqMsg"`

	// 生产时间
	ProduceTime *string `json:"ProduceTime,omitnil,omitempty" name:"ProduceTime"`
}

func (r *DescribeMqMsgTraceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMqMsgTraceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Protocol")
	delete(f, "MsgId")
	delete(f, "ClusterId")
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "QueueName")
	delete(f, "GroupName")
	delete(f, "QueryDlqMsg")
	delete(f, "ProduceTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMqMsgTraceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMqMsgTraceResponseParams struct {
	// 消息内容
	Result []*TraceResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 消息轨迹页展示的topic名称
	ShowTopicName *string `json:"ShowTopicName,omitnil,omitempty" name:"ShowTopicName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMqMsgTraceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMqMsgTraceResponseParams `json:"Response"`
}

func (r *DescribeMqMsgTraceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMqMsgTraceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMsgRequestParams struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 消息ID。
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 主题名。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeMsgRequest struct {
	*tchttp.BaseRequest
	
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 消息ID。
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 主题名。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "MsgId")
	delete(f, "TopicName")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMsgResponseParams struct {
	// 消息属性。
	Properties *string `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 消息体。
	Body *string `json:"Body,omitnil,omitempty" name:"Body"`

	// 批次ID。
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 生产时间。
	ProduceTime *string `json:"ProduceTime,omitnil,omitempty" name:"ProduceTime"`

	// 消息ID。
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 生产者名称。
	ProducerName *string `json:"ProducerName,omitnil,omitempty" name:"ProducerName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMsgResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMsgResponseParams `json:"Response"`
}

func (r *DescribeMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMsgTraceRequestParams struct {
	// 环境（命名空间）。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 消息ID。
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 消息生产时间。
	ProduceTime *string `json:"ProduceTime,omitnil,omitempty" name:"ProduceTime"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 消费组名称模糊匹配。
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeMsgTraceRequest struct {
	*tchttp.BaseRequest
	
	// 环境（命名空间）。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 消息ID。
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 消息生产时间。
	ProduceTime *string `json:"ProduceTime,omitnil,omitempty" name:"ProduceTime"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 消费组名称模糊匹配。
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeMsgTraceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMsgTraceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "MsgId")
	delete(f, "ProduceTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SubscriptionName")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMsgTraceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMsgTraceResponseParams struct {
	// 生产信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProducerLog *ProducerLog `json:"ProducerLog,omitnil,omitempty" name:"ProducerLog"`

	// 服务方信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServerLog *ServerLog `json:"ServerLog,omitnil,omitempty" name:"ServerLog"`

	// 消费信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerLogs *ConsumerLogs `json:"ConsumerLogs,omitnil,omitempty" name:"ConsumerLogs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMsgTraceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMsgTraceResponseParams `json:"Response"`
}

func (r *DescribeMsgTraceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMsgTraceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNamespaceBundlesOptRequestParams struct {
	// 物理集群名
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 虚拟集群（租户）ID
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 命名空间名
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 是否需要监控指标，若传false，则不需要传Limit和Offset分页参数
	NeedMetrics *bool `json:"NeedMetrics,omitnil,omitempty" name:"NeedMetrics"`

	// 查询限制条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤的 bundle
	Bundle *string `json:"Bundle,omitnil,omitempty" name:"Bundle"`

	// bundle 所属的 broker IP 地址，支持模糊查询
	OwnerBroker *string `json:"OwnerBroker,omitnil,omitempty" name:"OwnerBroker"`

	// 租户(如果没有自定义租户名称，和 tenantId 相同；如果有配置自定义租户名称，则为自定义租户名 user_tenant)
	Tenant *string `json:"Tenant,omitnil,omitempty" name:"Tenant"`
}

type DescribeNamespaceBundlesOptRequest struct {
	*tchttp.BaseRequest
	
	// 物理集群名
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 虚拟集群（租户）ID
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 命名空间名
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// 是否需要监控指标，若传false，则不需要传Limit和Offset分页参数
	NeedMetrics *bool `json:"NeedMetrics,omitnil,omitempty" name:"NeedMetrics"`

	// 查询限制条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤的 bundle
	Bundle *string `json:"Bundle,omitnil,omitempty" name:"Bundle"`

	// bundle 所属的 broker IP 地址，支持模糊查询
	OwnerBroker *string `json:"OwnerBroker,omitnil,omitempty" name:"OwnerBroker"`

	// 租户(如果没有自定义租户名称，和 tenantId 相同；如果有配置自定义租户名称，则为自定义租户名 user_tenant)
	Tenant *string `json:"Tenant,omitnil,omitempty" name:"Tenant"`
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
	delete(f, "Bundle")
	delete(f, "OwnerBroker")
	delete(f, "Tenant")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNamespaceBundlesOptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNamespaceBundlesOptResponseParams struct {
	// 记录条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNamespaceBundlesOptResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNamespaceBundlesOptResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeNodeHealthOptRequestParams struct {
	// 节点实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeNodeHealthOptRequest struct {
	*tchttp.BaseRequest
	
	// 节点实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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

// Predefined struct for user
type DescribeNodeHealthOptResponseParams struct {
	// 0-异常；1-正常
	NodeState *int64 `json:"NodeState,omitnil,omitempty" name:"NodeState"`

	// 最近一次健康检查的时间
	LatestHealthCheckTime *string `json:"LatestHealthCheckTime,omitnil,omitempty" name:"LatestHealthCheckTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNodeHealthOptResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNodeHealthOptResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribePublisherSummaryRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
}

type DescribePublisherSummaryRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
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

// Predefined struct for user
type DescribePublisherSummaryResponseParams struct {
	// 生产速率（条/秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateIn *float64 `json:"MsgRateIn,omitnil,omitempty" name:"MsgRateIn"`

	// 生产速率（字节/秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputIn *float64 `json:"MsgThroughputIn,omitnil,omitempty" name:"MsgThroughputIn"`

	// 生产者数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublisherCount *int64 `json:"PublisherCount,omitnil,omitempty" name:"PublisherCount"`

	// 消息存储大小，以字节为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePublisherSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublisherSummaryResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribePublishersRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 参数过滤器，支持ProducerName，Address字段
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询条数，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序器
	Sort *Sort `json:"Sort,omitnil,omitempty" name:"Sort"`
}

type DescribePublishersRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 参数过滤器，支持ProducerName，Address字段
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询条数，默认为20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序器
	Sort *Sort `json:"Sort,omitnil,omitempty" name:"Sort"`
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

// Predefined struct for user
type DescribePublishersResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 生产者信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Publishers []*Publisher `json:"Publishers,omitnil,omitempty" name:"Publishers"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePublishersResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublishersResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribePulsarProInstanceDetailRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribePulsarProInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribePulsarProInstanceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePulsarProInstanceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePulsarProInstanceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePulsarProInstanceDetailResponseParams struct {
	// 集群信息
	ClusterInfo *PulsarProClusterInfo `json:"ClusterInfo,omitnil,omitempty" name:"ClusterInfo"`

	// 集群网络接入点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkAccessPointInfos []*PulsarNetworkAccessPointInfo `json:"NetworkAccessPointInfos,omitnil,omitempty" name:"NetworkAccessPointInfos"`

	// 集群规格信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterSpecInfo *PulsarProClusterSpecInfo `json:"ClusterSpecInfo,omitnil,omitempty" name:"ClusterSpecInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePulsarProInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribePulsarProInstanceDetailResponseParams `json:"Response"`
}

func (r *DescribePulsarProInstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePulsarProInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePulsarProInstancesRequestParams struct {
	// 查询条件过滤器
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询数目上限，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询起始位置
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribePulsarProInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 查询条件过滤器
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询数目上限，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询起始位置
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribePulsarProInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePulsarProInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePulsarProInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePulsarProInstancesResponseParams struct {
	// 未分页的总数目
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例信息列表
	Instances []*PulsarProInstance `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePulsarProInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePulsarProInstancesResponseParams `json:"Response"`
}

func (r *DescribePulsarProInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePulsarProInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQBindingsRequestParams struct {
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到Vhost名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 分页 offset，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页 limit，默认 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键词，根据源exchange名称/目标资源名称/绑定key进行模糊搜索
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 根据源Exchange精准搜索过滤
	SourceExchange *string `json:"SourceExchange,omitnil,omitempty" name:"SourceExchange"`

	// 根据目标队列名精准搜索过滤，和 DestinationExchange 过滤不可同时设置
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 根据目标Exchange精准搜索过滤，和QueueName过滤不可同时设置
	DestinationExchange *string `json:"DestinationExchange,omitnil,omitempty" name:"DestinationExchange"`
}

type DescribeRabbitMQBindingsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到Vhost名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 分页 offset，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页 limit，默认 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键词，根据源exchange名称/目标资源名称/绑定key进行模糊搜索
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 根据源Exchange精准搜索过滤
	SourceExchange *string `json:"SourceExchange,omitnil,omitempty" name:"SourceExchange"`

	// 根据目标队列名精准搜索过滤，和 DestinationExchange 过滤不可同时设置
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 根据目标Exchange精准搜索过滤，和QueueName过滤不可同时设置
	DestinationExchange *string `json:"DestinationExchange,omitnil,omitempty" name:"DestinationExchange"`
}

func (r *DescribeRabbitMQBindingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQBindingsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	delete(f, "SourceExchange")
	delete(f, "QueueName")
	delete(f, "DestinationExchange")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQBindingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQBindingsResponseParams struct {
	// 路由关系列表
	BindingInfoList []*RabbitMQBindingListInfo `json:"BindingInfoList,omitnil,omitempty" name:"BindingInfoList"`

	// 路由关系数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQBindingsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQBindingsResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQBindingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQBindingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQExchangesRequestParams struct {
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到Vhost名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 分页 offset，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页 limit，默认 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键词, 支持模糊匹配 
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 筛选 exchange 类型, 数组中每个元素为选中的过滤类型，仅支持 direct、fanout、topic、header
	ExchangeTypeFilters []*string `json:"ExchangeTypeFilters,omitnil,omitempty" name:"ExchangeTypeFilters"`

	// 筛选 exchange 创建来源,  "system":"系统创建", "user":"用户创建"
	ExchangeCreatorFilters []*string `json:"ExchangeCreatorFilters,omitnil,omitempty" name:"ExchangeCreatorFilters"`

	// exchange 名称，用于精确匹配
	ExchangeName *string `json:"ExchangeName,omitnil,omitempty" name:"ExchangeName"`

	// 排序依据的字段：
	// MessageRateInOut - 生产消费速率之和；
	// MessageRateIn - 生产速率；
	// MessageRateOut - 消费速率；
	SortElement *string `json:"SortElement,omitnil,omitempty" name:"SortElement"`

	// 排序顺序，ascend 或 descend
	// ascend：升序
	// descend：降序
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

type DescribeRabbitMQExchangesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到Vhost名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 分页 offset，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页 limit，默认 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键词, 支持模糊匹配 
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 筛选 exchange 类型, 数组中每个元素为选中的过滤类型，仅支持 direct、fanout、topic、header
	ExchangeTypeFilters []*string `json:"ExchangeTypeFilters,omitnil,omitempty" name:"ExchangeTypeFilters"`

	// 筛选 exchange 创建来源,  "system":"系统创建", "user":"用户创建"
	ExchangeCreatorFilters []*string `json:"ExchangeCreatorFilters,omitnil,omitempty" name:"ExchangeCreatorFilters"`

	// exchange 名称，用于精确匹配
	ExchangeName *string `json:"ExchangeName,omitnil,omitempty" name:"ExchangeName"`

	// 排序依据的字段：
	// MessageRateInOut - 生产消费速率之和；
	// MessageRateIn - 生产速率；
	// MessageRateOut - 消费速率；
	SortElement *string `json:"SortElement,omitnil,omitempty" name:"SortElement"`

	// 排序顺序，ascend 或 descend
	// ascend：升序
	// descend：降序
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

func (r *DescribeRabbitMQExchangesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQExchangesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	delete(f, "ExchangeTypeFilters")
	delete(f, "ExchangeCreatorFilters")
	delete(f, "ExchangeName")
	delete(f, "SortElement")
	delete(f, "SortOrder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQExchangesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQExchangesResponseParams struct {
	// 策略列表信息
	ExchangeInfoList []*RabbitMQExchangeListInfo `json:"ExchangeInfoList,omitnil,omitempty" name:"ExchangeInfoList"`

	// 策略结果总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQExchangesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQExchangesResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQExchangesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQExchangesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQNodeListRequestParams struct {
	// 实例 ID，形如amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 偏移量，默认值 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 一页限制，默认值 20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模糊搜索节点名字
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 过滤参数的名字和数值，当前仅支持根据节点状态筛选。
	// "Name": "nodeStatus"
	// "Value": running or down
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 按指定元素排序，现在只有2个
	// cpuUsage：节点CPU利用率
	// diskUsage：节点磁盘利用率
	SortElement *string `json:"SortElement,omitnil,omitempty" name:"SortElement"`

	// 升序/降序
	// ascend/descend
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

type DescribeRabbitMQNodeListRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 偏移量，默认值 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 一页限制，默认值 20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模糊搜索节点名字
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 过滤参数的名字和数值，当前仅支持根据节点状态筛选。
	// "Name": "nodeStatus"
	// "Value": running or down
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 按指定元素排序，现在只有2个
	// cpuUsage：节点CPU利用率
	// diskUsage：节点磁盘利用率
	SortElement *string `json:"SortElement,omitnil,omitempty" name:"SortElement"`

	// 升序/降序
	// ascend/descend
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

func (r *DescribeRabbitMQNodeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQNodeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NodeName")
	delete(f, "Filters")
	delete(f, "SortElement")
	delete(f, "SortOrder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQNodeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQNodeListResponseParams struct {
	// 集群节点数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集群节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeList []*RabbitMQPrivateNode `json:"NodeList,omitnil,omitempty" name:"NodeList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQNodeListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQNodeListResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQNodeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQNodeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQPermissionRequestParams struct {
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，形如 admin。有效的 User 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，点击集群列表中的集群，进入集群详情，并在用户与权限页签中找到用户列表，从而找到用户名称。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到 Vhost 名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 分页 Offset，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页 Limit，默认 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeRabbitMQPermissionRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，形如 admin。有效的 User 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，点击集群列表中的集群，进入集群详情，并在用户与权限页签中找到用户列表，从而找到用户名称。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到 Vhost 名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 分页 Offset，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页 Limit，默认 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeRabbitMQPermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQPermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "VirtualHost")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQPermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQPermissionResponseParams struct {
	// 返回权限数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 权限详情列表
	RabbitMQPermissionList []*RabbitMQPermission `json:"RabbitMQPermissionList,omitnil,omitempty" name:"RabbitMQPermissionList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQPermissionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQPermissionResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQPermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQPermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQQueueDetailRequestParams struct {
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到 Vhost 名称
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 队列名称
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`
}

type DescribeRabbitMQQueueDetailRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到 Vhost 名称
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 队列名称
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`
}

func (r *DescribeRabbitMQQueueDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQQueueDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQQueueDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQQueueDetailResponseParams struct {
	// 实例名称
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 队列名称
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 队列类型,取值classic或quorum
	QueueType *string `json:"QueueType,omitnil,omitempty" name:"QueueType"`

	// 在线消费者数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Consumers *int64 `json:"Consumers,omitnil,omitempty" name:"Consumers"`

	// 持久标记
	Durable *bool `json:"Durable,omitnil,omitempty" name:"Durable"`

	// 自动清除
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoDelete *bool `json:"AutoDelete,omitnil,omitempty" name:"AutoDelete"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// MessageTTL参数,classic类型专用
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageTTL *int64 `json:"MessageTTL,omitnil,omitempty" name:"MessageTTL"`

	// AutoExpire参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoExpire *int64 `json:"AutoExpire,omitnil,omitempty" name:"AutoExpire"`

	// MaxLength参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxLength *int64 `json:"MaxLength,omitnil,omitempty" name:"MaxLength"`

	// MaxLengthBytes参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxLengthBytes *int64 `json:"MaxLengthBytes,omitnil,omitempty" name:"MaxLengthBytes"`

	// DeliveryLimit参数,quorum类型专用
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeliveryLimit *int64 `json:"DeliveryLimit,omitnil,omitempty" name:"DeliveryLimit"`

	// OverflowBehaviour参数,取值为drop-head, reject-publish或reject-publish-dlx
	// 注意：此字段可能返回 null，表示取不到有效值。
	OverflowBehaviour *string `json:"OverflowBehaviour,omitnil,omitempty" name:"OverflowBehaviour"`

	// DeadLetterExchange参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterExchange *string `json:"DeadLetterExchange,omitnil,omitempty" name:"DeadLetterExchange"`

	// DeadLetterRoutingKey参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterRoutingKey *string `json:"DeadLetterRoutingKey,omitnil,omitempty" name:"DeadLetterRoutingKey"`

	// SingleActiveConsumer参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SingleActiveConsumer *bool `json:"SingleActiveConsumer,omitnil,omitempty" name:"SingleActiveConsumer"`

	// MaximumPriority参数,classic类型专用
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaximumPriority *int64 `json:"MaximumPriority,omitnil,omitempty" name:"MaximumPriority"`

	// LazyMode参数,classic类型专用
	// 注意：此字段可能返回 null，表示取不到有效值。
	LazyMode *bool `json:"LazyMode,omitnil,omitempty" name:"LazyMode"`

	// MasterLocator参数,classic类型专用
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterLocator *string `json:"MasterLocator,omitnil,omitempty" name:"MasterLocator"`

	// MaxInMemoryLength参数,quorum类型专用
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxInMemoryLength *int64 `json:"MaxInMemoryLength,omitnil,omitempty" name:"MaxInMemoryLength"`

	// MaxInMemoryBytes参数,quorum类型专用
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxInMemoryBytes *int64 `json:"MaxInMemoryBytes,omitnil,omitempty" name:"MaxInMemoryBytes"`

	// 创建时间戳,单位秒
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Node *string `json:"Node,omitnil,omitempty" name:"Node"`

	// 仲裁队列死信一致性策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterStrategy *string `json:"DeadLetterStrategy,omitnil,omitempty" name:"DeadLetterStrategy"`

	// 仲裁队列的领导者选举策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueLeaderLocator *string `json:"QueueLeaderLocator,omitnil,omitempty" name:"QueueLeaderLocator"`

	// 仲裁队列的初始副本组大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuorumInitialGroupSize *int64 `json:"QuorumInitialGroupSize,omitnil,omitempty" name:"QuorumInitialGroupSize"`

	// 是否为独占队列
	Exclusive *bool `json:"Exclusive,omitnil,omitempty" name:"Exclusive"`

	// 生效的策略名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 扩展参数 key-value
	Arguments *string `json:"Arguments,omitnil,omitempty" name:"Arguments"`

	// 创建时间时间戳
	CreateTs *uint64 `json:"CreateTs,omitnil,omitempty" name:"CreateTs"`

	// 修改时间时间戳
	ModifyTs *uint64 `json:"ModifyTs,omitnil,omitempty" name:"ModifyTs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQQueueDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQQueueDetailResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQQueueDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQQueueDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQQueuesRequestParams struct {
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到Vhost名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 分页 Offset，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页 Limit，默认 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键词
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 队列类型筛选，不填或 "all"：筛选普通队列 和 quorum 队列；"classic"：筛选 classic(普通) 队列；"quorum"：筛选 quorum 队列
	QueueType *string `json:"QueueType,omitnil,omitempty" name:"QueueType"`

	// 排序依据的字段：
	// ConsumerNumber - 在线消费者数量；
	// MessageHeapCount - 消息堆积数；
	// MessageRateInOut - 生产消费速率之和；
	// MessageRateIn - 生产速率；
	// MessageRateOut - 消费速率；
	SortElement *string `json:"SortElement,omitnil,omitempty" name:"SortElement"`

	// 排序顺序，ascend 或 descend
	// ascend：升序
	// descend：降序
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

type DescribeRabbitMQQueuesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到Vhost名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 分页 Offset，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页 Limit，默认 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键词
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 队列类型筛选，不填或 "all"：筛选普通队列 和 quorum 队列；"classic"：筛选 classic(普通) 队列；"quorum"：筛选 quorum 队列
	QueueType *string `json:"QueueType,omitnil,omitempty" name:"QueueType"`

	// 排序依据的字段：
	// ConsumerNumber - 在线消费者数量；
	// MessageHeapCount - 消息堆积数；
	// MessageRateInOut - 生产消费速率之和；
	// MessageRateIn - 生产速率；
	// MessageRateOut - 消费速率；
	SortElement *string `json:"SortElement,omitnil,omitempty" name:"SortElement"`

	// 排序顺序，ascend 或 descend
	// ascend：升序
	// descend：降序
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

func (r *DescribeRabbitMQQueuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQQueuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchWord")
	delete(f, "QueueType")
	delete(f, "SortElement")
	delete(f, "SortOrder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQQueuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQQueuesResponseParams struct {
	// 队列列表信息
	QueueInfoList []*RabbitMQQueueListInfo `json:"QueueInfoList,omitnil,omitempty" name:"QueueInfoList"`

	// 队列数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQQueuesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQQueuesResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQQueuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQQueuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQUserRequestParams struct {
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名检索，支持前缀匹配，后缀匹配
	SearchUser *string `json:"SearchUser,omitnil,omitempty" name:"SearchUser"`

	// 分页 Offset，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页 Limit，默认 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 用户名，精确查询
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 用户标签，用于决定改用户访问 RabbitMQ Management 的权限范围
	// management：普通控制台用户，monitoring：管理型控制台用户，其他值：非控制台用户
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRabbitMQUserRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名检索，支持前缀匹配，后缀匹配
	SearchUser *string `json:"SearchUser,omitnil,omitempty" name:"SearchUser"`

	// 分页 Offset，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页 Limit，默认 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 用户名，精确查询
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 用户标签，用于决定改用户访问 RabbitMQ Management 的权限范围
	// management：普通控制台用户，monitoring：管理型控制台用户，其他值：非控制台用户
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRabbitMQUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SearchUser")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "User")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQUserResponseParams struct {
	// 返回的User数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 当前已创建的RabbitMQ用户列表
	RabbitMQUserList []*RabbitMQUser `json:"RabbitMQUserList,omitnil,omitempty" name:"RabbitMQUserList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQUserResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQUserResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQVipInstanceRequestParams struct {
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeRabbitMQVipInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeRabbitMQVipInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQVipInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQVipInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQVipInstanceResponseParams struct {
	// 集群信息
	ClusterInfo *RabbitMQClusterInfo `json:"ClusterInfo,omitnil,omitempty" name:"ClusterInfo"`

	// 集群规格信息
	ClusterSpecInfo *RabbitMQClusterSpecInfo `json:"ClusterSpecInfo,omitnil,omitempty" name:"ClusterSpecInfo"`

	// 集群访问
	ClusterNetInfo *RabbitMQClusterAccessInfo `json:"ClusterNetInfo,omitnil,omitempty" name:"ClusterNetInfo"`

	// 集群白名单
	ClusterWhiteListInfo *RabbitMQClusterWhiteListInfo `json:"ClusterWhiteListInfo,omitnil,omitempty" name:"ClusterWhiteListInfo"`

	// vhost配额信息
	VirtualHostQuota *VirtualHostQuota `json:"VirtualHostQuota,omitnil,omitempty" name:"VirtualHostQuota"`

	// exchange配额信息
	ExchangeQuota *ExchangeQuota `json:"ExchangeQuota,omitnil,omitempty" name:"ExchangeQuota"`

	// queue配额信息
	QueueQuota *QueueQuota `json:"QueueQuota,omitnil,omitempty" name:"QueueQuota"`

	// 用户配额信息
	UserQuota *RabbitMQUserQuota `json:"UserQuota,omitnil,omitempty" name:"UserQuota"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQVipInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQVipInstanceResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQVipInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQVipInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQVipInstancesRequestParams struct {
	// 查询条件过滤器
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询数目上限，默认 20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询起始位置，默认 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeRabbitMQVipInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 查询条件过滤器
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询数目上限，默认 20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询起始位置，默认 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeRabbitMQVipInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQVipInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQVipInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQVipInstancesResponseParams struct {
	// 未分页的总数目
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例信息列表
	Instances []*RabbitMQVipInstance `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQVipInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQVipInstancesResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQVipInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQVipInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQVirtualHostRequestParams struct {
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost名,不传则查询全部
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// search-virtual-host：vhost名称模糊查询，之前前缀和后缀匹配
	Filters *Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序依据的字段：
	// MessageHeapCount - 消息堆积数；
	// MessageRateInOut - 生产消费速率之和；
	// MessageRateIn - 生产速率；
	// MessageRateOut - 消费速率；
	SortElement *string `json:"SortElement,omitnil,omitempty" name:"SortElement"`

	// 排序顺序，ascend 或 descend
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

type DescribeRabbitMQVirtualHostRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost名,不传则查询全部
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// search-virtual-host：vhost名称模糊查询，之前前缀和后缀匹配
	Filters *Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序依据的字段：
	// MessageHeapCount - 消息堆积数；
	// MessageRateInOut - 生产消费速率之和；
	// MessageRateIn - 生产速率；
	// MessageRateOut - 消费速率；
	SortElement *string `json:"SortElement,omitnil,omitempty" name:"SortElement"`

	// 排序顺序，ascend 或 descend
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

func (r *DescribeRabbitMQVirtualHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQVirtualHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "SortElement")
	delete(f, "SortOrder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQVirtualHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQVirtualHostResponseParams struct {
	// 返回vhost数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// vhost详情列表
	VirtualHostList []*RabbitMQVirtualHostInfo `json:"VirtualHostList,omitnil,omitempty" name:"VirtualHostList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQVirtualHostResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQVirtualHostResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQVirtualHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQVirtualHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQClusterRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeRocketMQClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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

// Predefined struct for user
type DescribeRocketMQClusterResponseParams struct {
	// 集群信息
	ClusterInfo *RocketMQClusterInfo `json:"ClusterInfo,omitnil,omitempty" name:"ClusterInfo"`

	// 集群配置
	ClusterConfig *RocketMQClusterConfig `json:"ClusterConfig,omitnil,omitempty" name:"ClusterConfig"`

	// 集群最近使用量，即将废弃，请使用腾讯云可观测平台获取相关数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterStats *RocketMQClusterRecentStats `json:"ClusterStats,omitnil,omitempty" name:"ClusterStats"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQClusterResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQClusterResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeRocketMQClustersRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 按照集群ID关键字搜索
	IdKeyword *string `json:"IdKeyword,omitnil,omitempty" name:"IdKeyword"`

	// 按照集群名称关键字搜索
	NameKeyword *string `json:"NameKeyword,omitnil,omitempty" name:"NameKeyword"`

	// 集群ID列表过滤
	ClusterIdList []*string `json:"ClusterIdList,omitnil,omitempty" name:"ClusterIdList"`

	// 标签过滤查找时，需要设置为true
	IsTagFilter *bool `json:"IsTagFilter,omitnil,omitempty" name:"IsTagFilter"`

	// 过滤器。目前支持标签过滤。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRocketMQClustersRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 按照集群ID关键字搜索
	IdKeyword *string `json:"IdKeyword,omitnil,omitempty" name:"IdKeyword"`

	// 按照集群名称关键字搜索
	NameKeyword *string `json:"NameKeyword,omitnil,omitempty" name:"NameKeyword"`

	// 集群ID列表过滤
	ClusterIdList []*string `json:"ClusterIdList,omitnil,omitempty" name:"ClusterIdList"`

	// 标签过滤查找时，需要设置为true
	IsTagFilter *bool `json:"IsTagFilter,omitnil,omitempty" name:"IsTagFilter"`

	// 过滤器。目前支持标签过滤。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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

// Predefined struct for user
type DescribeRocketMQClustersResponseParams struct {
	// 集群信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterList []*RocketMQClusterDetail `json:"ClusterList,omitnil,omitempty" name:"ClusterList"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQClustersResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeRocketMQConsumeStatsRequestParams struct {
	// 实例ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 消费组
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`
}

type DescribeRocketMQConsumeStatsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 消费组
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`
}

func (r *DescribeRocketMQConsumeStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQConsumeStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "ConsumerGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQConsumeStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQConsumeStatsResponseParams struct {
	// 消费详情列表
	ConsumerStatsList []*ConsumerStats `json:"ConsumerStatsList,omitnil,omitempty" name:"ConsumerStatsList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQConsumeStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQConsumeStatsResponseParams `json:"Response"`
}

func (r *DescribeRocketMQConsumeStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQConsumeStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQConsumerConnectionDetailRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 消费组名称
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 消费端实例ID
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 按主题类型过滤查询结果，可选择Normal, GlobalOrder, PartitionedOrder, Retry, Transaction, DeadLetter
	FilterType []*string `json:"FilterType,omitnil,omitempty" name:"FilterType"`
}

type DescribeRocketMQConsumerConnectionDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 消费组名称
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 消费端实例ID
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 按主题类型过滤查询结果，可选择Normal, GlobalOrder, PartitionedOrder, Retry, Transaction, DeadLetter
	FilterType []*string `json:"FilterType,omitnil,omitempty" name:"FilterType"`
}

func (r *DescribeRocketMQConsumerConnectionDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQConsumerConnectionDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "GroupId")
	delete(f, "ClientId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQConsumerConnectionDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQConsumerConnectionDetailResponseParams struct {
	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 消费端主题信息列表
	Details []*RocketMQConsumerTopic `json:"Details,omitnil,omitempty" name:"Details"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQConsumerConnectionDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQConsumerConnectionDetailResponseParams `json:"Response"`
}

func (r *DescribeRocketMQConsumerConnectionDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQConsumerConnectionDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQConsumerConnectionsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 消费组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 对查询结果排序，此为排序字段，目前支持Accumulative（消息堆积量）
	SortedBy *string `json:"SortedBy,omitnil,omitempty" name:"SortedBy"`

	// 查询结果排序规则，ASC为升序，DESC为降序
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

type DescribeRocketMQConsumerConnectionsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 消费组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 对查询结果排序，此为排序字段，目前支持Accumulative（消息堆积量）
	SortedBy *string `json:"SortedBy,omitnil,omitempty" name:"SortedBy"`

	// 查询结果排序规则，ASC为升序，DESC为降序
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

func (r *DescribeRocketMQConsumerConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQConsumerConnectionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "GroupId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SortedBy")
	delete(f, "SortOrder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQConsumerConnectionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQConsumerConnectionsResponseParams struct {
	// 总数目
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 在线消费者信息
	Connections []*RocketMQConsumerConnection `json:"Connections,omitnil,omitempty" name:"Connections"`

	// 订阅组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupDetail *RocketMQGroup `json:"GroupDetail,omitnil,omitempty" name:"GroupDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQConsumerConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQConsumerConnectionsResponseParams `json:"Response"`
}

func (r *DescribeRocketMQConsumerConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQConsumerConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQEnvironmentRolesRequestParams struct {
	// 必填字段，RocketMQ集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 起始下标，不填默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 角色名称
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// RoleName按照角色名进行过滤，精确查询。类型：String必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRocketMQEnvironmentRolesRequest struct {
	*tchttp.BaseRequest
	
	// 必填字段，RocketMQ集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 起始下标，不填默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 角色名称
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// RoleName按照角色名进行过滤，精确查询。类型：String必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeRocketMQEnvironmentRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQEnvironmentRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "EnvironmentId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "RoleName")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQEnvironmentRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQEnvironmentRolesResponseParams struct {
	// 记录数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 命名空间角色集合。
	EnvironmentRoleSets []*EnvironmentRole `json:"EnvironmentRoleSets,omitnil,omitempty" name:"EnvironmentRoleSets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQEnvironmentRolesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQEnvironmentRolesResponseParams `json:"Response"`
}

func (r *DescribeRocketMQEnvironmentRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQEnvironmentRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQGroupsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 主题名称，输入此参数可查询该主题下所有的订阅组
	FilterTopic *string `json:"FilterTopic,omitnil,omitempty" name:"FilterTopic"`

	// 按消费组名称查询消费组，支持模糊查询
	FilterGroup *string `json:"FilterGroup,omitnil,omitempty" name:"FilterGroup"`

	// 按照指定字段排序，可选值为 subscribeNum: 订阅 Topic 个数
	SortedBy *string `json:"SortedBy,omitnil,omitempty" name:"SortedBy"`

	// 按升序或降序排列，可选值为asc，desc
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// 订阅组名称，指定此参数后将只返回该订阅组信息
	FilterOneGroup *string `json:"FilterOneGroup,omitnil,omitempty" name:"FilterOneGroup"`

	// group类型
	Types []*string `json:"Types,omitnil,omitempty" name:"Types"`
}

type DescribeRocketMQGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 主题名称，输入此参数可查询该主题下所有的订阅组
	FilterTopic *string `json:"FilterTopic,omitnil,omitempty" name:"FilterTopic"`

	// 按消费组名称查询消费组，支持模糊查询
	FilterGroup *string `json:"FilterGroup,omitnil,omitempty" name:"FilterGroup"`

	// 按照指定字段排序，可选值为 subscribeNum: 订阅 Topic 个数
	SortedBy *string `json:"SortedBy,omitnil,omitempty" name:"SortedBy"`

	// 按升序或降序排列，可选值为asc，desc
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// 订阅组名称，指定此参数后将只返回该订阅组信息
	FilterOneGroup *string `json:"FilterOneGroup,omitnil,omitempty" name:"FilterOneGroup"`

	// group类型
	Types []*string `json:"Types,omitnil,omitempty" name:"Types"`
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
	delete(f, "Types")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQGroupsResponseParams struct {
	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 订阅组列表
	Groups []*RocketMQGroup `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQGroupsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeRocketMQMigratingTopicListRequestParams struct {
	// 迁移任务名称
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询过滤器，支持topicname、MigrationStatus查询
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRocketMQMigratingTopicListRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务名称
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询过滤器，支持topicname、MigrationStatus查询
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeRocketMQMigratingTopicListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQMigratingTopicListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQMigratingTopicListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQMigratingTopicListResponseParams struct {
	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 迁移topic列表
	MigrateTopics []*MigrateTopic `json:"MigrateTopics,omitnil,omitempty" name:"MigrateTopics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQMigratingTopicListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQMigratingTopicListResponseParams `json:"Response"`
}

func (r *DescribeRocketMQMigratingTopicListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQMigratingTopicListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQMsgRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题，查询死信时传groupId
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 消息id
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// pulsar消息id
	PulsarMsgId *string `json:"PulsarMsgId,omitnil,omitempty" name:"PulsarMsgId"`

	// 查询死信时该值为true，只对Rocketmq有效
	//
	// Deprecated: QueryDlqMsg is deprecated.
	QueryDlqMsg *bool `json:"QueryDlqMsg,omitnil,omitempty" name:"QueryDlqMsg"`

	// 查询死信时该值为true，只对Rocketmq有效
	QueryDeadLetterMessage *bool `json:"QueryDeadLetterMessage,omitnil,omitempty" name:"QueryDeadLetterMessage"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据消费组名称过滤消费详情
	FilterTrackGroup *string `json:"FilterTrackGroup,omitnil,omitempty" name:"FilterTrackGroup"`
}

type DescribeRocketMQMsgRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题，查询死信时传groupId
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 消息id
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// pulsar消息id
	PulsarMsgId *string `json:"PulsarMsgId,omitnil,omitempty" name:"PulsarMsgId"`

	// 查询死信时该值为true，只对Rocketmq有效
	QueryDlqMsg *bool `json:"QueryDlqMsg,omitnil,omitempty" name:"QueryDlqMsg"`

	// 查询死信时该值为true，只对Rocketmq有效
	QueryDeadLetterMessage *bool `json:"QueryDeadLetterMessage,omitnil,omitempty" name:"QueryDeadLetterMessage"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据消费组名称过滤消费详情
	FilterTrackGroup *string `json:"FilterTrackGroup,omitnil,omitempty" name:"FilterTrackGroup"`
}

func (r *DescribeRocketMQMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQMsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "MsgId")
	delete(f, "PulsarMsgId")
	delete(f, "QueryDlqMsg")
	delete(f, "QueryDeadLetterMessage")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterTrackGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQMsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQMsgResponseParams struct {
	// 消息体
	Body *string `json:"Body,omitnil,omitempty" name:"Body"`

	// 详情参数
	Properties *string `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 生产时间
	ProduceTime *string `json:"ProduceTime,omitnil,omitempty" name:"ProduceTime"`

	// 消息id
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 生产者地址
	ProducerAddr *string `json:"ProducerAddr,omitnil,omitempty" name:"ProducerAddr"`

	// 消费组消费情况列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageTracks []*RocketMQMessageTrack `json:"MessageTracks,omitnil,omitempty" name:"MessageTracks"`

	// 详情页展示的topic名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowTopicName *string `json:"ShowTopicName,omitnil,omitempty" name:"ShowTopicName"`

	// 消费组消费情况列表总数
	MessageTracksCount *int64 `json:"MessageTracksCount,omitnil,omitempty" name:"MessageTracksCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQMsgResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQMsgResponseParams `json:"Response"`
}

func (r *DescribeRocketMQMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQMsgTraceRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题，rocketmq查询死信时值为groupId
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 消息id
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 消费组、订阅
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 查询死信时该值为true
	//
	// Deprecated: QueryDLQMsg is deprecated.
	QueryDLQMsg *bool `json:"QueryDLQMsg,omitnil,omitempty" name:"QueryDLQMsg"`

	// 查询死信时该值为true
	QueryDeadLetterMessage *string `json:"QueryDeadLetterMessage,omitnil,omitempty" name:"QueryDeadLetterMessage"`
}

type DescribeRocketMQMsgTraceRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题，rocketmq查询死信时值为groupId
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 消息id
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 消费组、订阅
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 查询死信时该值为true
	QueryDLQMsg *bool `json:"QueryDLQMsg,omitnil,omitempty" name:"QueryDLQMsg"`

	// 查询死信时该值为true
	QueryDeadLetterMessage *string `json:"QueryDeadLetterMessage,omitnil,omitempty" name:"QueryDeadLetterMessage"`
}

func (r *DescribeRocketMQMsgTraceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQMsgTraceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "MsgId")
	delete(f, "GroupName")
	delete(f, "QueryDLQMsg")
	delete(f, "QueryDeadLetterMessage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQMsgTraceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQMsgTraceResponseParams struct {
	// 轨迹详情列表
	Result []*TraceResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 消息轨迹页展示的topic名称
	ShowTopicName *string `json:"ShowTopicName,omitnil,omitempty" name:"ShowTopicName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQMsgTraceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQMsgTraceResponseParams `json:"Response"`
}

func (r *DescribeRocketMQMsgTraceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQMsgTraceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQNamespacesRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 按名称搜索
	NameKeyword *string `json:"NameKeyword,omitnil,omitempty" name:"NameKeyword"`
}

type DescribeRocketMQNamespacesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 按名称搜索
	NameKeyword *string `json:"NameKeyword,omitnil,omitempty" name:"NameKeyword"`
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

// Predefined struct for user
type DescribeRocketMQNamespacesResponseParams struct {
	// 命名空间列表
	Namespaces []*RocketMQNamespace `json:"Namespaces,omitnil,omitempty" name:"Namespaces"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQNamespacesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeRocketMQProducersRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 主题名
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 分页offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤查询条件列表，支持以下过滤参数：
	// 
	// - ClientId：生产者客户端ID
	// - ClientIp：生产者客户端IP
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRocketMQProducersRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 主题名
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 分页offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤查询条件列表，支持以下过滤参数：
	// 
	// - ClientId：生产者客户端ID
	// - ClientIp：生产者客户端IP
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeRocketMQProducersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQProducersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "Topic")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQProducersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQProducersResponseParams struct {
	// 生产者客户端列表
	Producers []*ProducerInfo `json:"Producers,omitnil,omitempty" name:"Producers"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQProducersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQProducersResponseParams `json:"Response"`
}

func (r *DescribeRocketMQProducersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQProducersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQPublicAccessMonitorDataRequestParams struct {
	// 专享集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指标名称，仅支持单指标拉取。目前仅支持：ClientIntraffic; ClientOuttraffic
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，默认为当前时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 监控统计周期，如60。默认为取值为300，单位为s。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`
}

type DescribeRocketMQPublicAccessMonitorDataRequest struct {
	*tchttp.BaseRequest
	
	// 专享集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指标名称，仅支持单指标拉取。目前仅支持：ClientIntraffic; ClientOuttraffic
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，默认为当前时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 监控统计周期，如60。默认为取值为300，单位为s。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`
}

func (r *DescribeRocketMQPublicAccessMonitorDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQPublicAccessMonitorDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "MetricName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQPublicAccessMonitorDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQPublicAccessMonitorDataResponseParams struct {
	// 指标名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 监控统计周期，如60。默认为取值为300，单位为s。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 起始时间，如2018-09-22T19:51:23+08:00
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如2018-09-22T20:51:23+08:00，默认为当前时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 数据点数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataPoints []*RocketMQDataPoint `json:"DataPoints,omitnil,omitempty" name:"DataPoints"`

	// 返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQPublicAccessMonitorDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQPublicAccessMonitorDataResponseParams `json:"Response"`
}

func (r *DescribeRocketMQPublicAccessMonitorDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQPublicAccessMonitorDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQPublicAccessPointRequestParams struct {
	// 集群ID，当前只支持专享集群
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeRocketMQPublicAccessPointRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，当前只支持专享集群
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeRocketMQPublicAccessPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQPublicAccessPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQPublicAccessPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQPublicAccessPointResponseParams struct {
	// 公网接入点状态：
	// 0， 已开启
	// 1， 已关闭
	// 2，开启中
	// 3，关闭中
	// 4，修改中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 支付状态：
	// 0, 未知
	// 1，正常
	// 2，欠费
	PayStatus *int64 `json:"PayStatus,omitnil,omitempty" name:"PayStatus"`

	// 接入点地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessUrl *string `json:"AccessUrl,omitnil,omitempty" name:"AccessUrl"`

	// 安全访问规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*PublicAccessRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 付费模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 公网是否按流量计费
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingFlow *bool `json:"BillingFlow,omitnil,omitempty" name:"BillingFlow"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQPublicAccessPointResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQPublicAccessPointResponseParams `json:"Response"`
}

func (r *DescribeRocketMQPublicAccessPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQPublicAccessPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQRolesRequestParams struct {
	// 起始下标，不填默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 角色名称，模糊查询
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// RoleName按照角色名进行过滤，精确查询。类型：String必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRocketMQRolesRequest struct {
	*tchttp.BaseRequest
	
	// 起始下标，不填默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 角色名称，模糊查询
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// RoleName按照角色名进行过滤，精确查询。类型：String必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeRocketMQRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClusterId")
	delete(f, "RoleName")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQRolesResponseParams struct {
	// 记录数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 角色数组。
	RoleSets []*Role `json:"RoleSets,omitnil,omitempty" name:"RoleSets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQRolesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQRolesResponseParams `json:"Response"`
}

func (r *DescribeRocketMQRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQSmoothMigrationTaskListRequestParams struct {
	// 查询起始偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询最大数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询过滤器，
	// 支持的字段如下
	// TaskStatus, 支持多选
	// ConnectionType，支持多选
	// ClusterId，精确搜索
	// TaskName，支持模糊搜索
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRocketMQSmoothMigrationTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询最大数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询过滤器，
	// 支持的字段如下
	// TaskStatus, 支持多选
	// ConnectionType，支持多选
	// ClusterId，精确搜索
	// TaskName，支持模糊搜索
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeRocketMQSmoothMigrationTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQSmoothMigrationTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQSmoothMigrationTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQSmoothMigrationTaskListResponseParams struct {
	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务列表
	Data []*RocketMQSmoothMigrationTaskItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQSmoothMigrationTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQSmoothMigrationTaskListResponseParams `json:"Response"`
}

func (r *DescribeRocketMQSmoothMigrationTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQSmoothMigrationTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQSmoothMigrationTaskRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeRocketMQSmoothMigrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeRocketMQSmoothMigrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQSmoothMigrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQSmoothMigrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQSmoothMigrationTaskResponseParams struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 目标集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 源集群名称
	SourceClusterName *string `json:"SourceClusterName,omitnil,omitempty" name:"SourceClusterName"`

	// 网络连接类型，
	// PUBLIC 公网
	// VPC 私有网络
	// OTHER 其它
	ConnectionType *string `json:"ConnectionType,omitnil,omitempty" name:"ConnectionType"`

	// 源集群NameServer地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceClusterNameServer *string `json:"SourceClusterNameServer,omitnil,omitempty" name:"SourceClusterNameServer"`

	// 源集群所在私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 源集群所在子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 是否开启ACL
	EnableACL *bool `json:"EnableACL,omitnil,omitempty" name:"EnableACL"`

	// 源集群AccessKey
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// 元集群SecretKey
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// 配置源集群时发生的错误
	// TIMEOUT 连接超时，
	// SERVER_ERROR 服务错误，
	// INTERNAL_ERROR 内部错误，
	// CONNECT_NAMESERVER_ERROR 连接nameserver错误
	// CONNECT_BROKER_ERROR 连接broker错误
	// ACL_WRONG ACL信息不正确
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskError *string `json:"TaskError,omitnil,omitempty" name:"TaskError"`

	// 任务状态
	// Configuration 迁移配置
	// SourceConnecting 连接源集群中
	// SourceConnectionFailure 连接源集群失败
	// MetaDataImport 元数据导入
	// EndpointSetup 切换接入点
	// ServiceMigration 切流中
	// Completed 已完成
	// Cancelled 已取消
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 主题类型分布情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicTypeDistribution []*RocketMQTopicDistribution `json:"TopicTypeDistribution,omitnil,omitempty" name:"TopicTypeDistribution"`

	// 主题迁移进度分布情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicStageDistribution []*RocketMQMigrationTopicDistribution `json:"TopicStageDistribution,omitnil,omitempty" name:"TopicStageDistribution"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQSmoothMigrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQSmoothMigrationTaskResponseParams `json:"Response"`
}

func (r *DescribeRocketMQSmoothMigrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQSmoothMigrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQSourceClusterGroupListRequestParams struct {
	// 页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 迁移任务名称
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 查询过滤器，支持字段groupName，imported
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRocketMQSourceClusterGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 迁移任务名称
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 查询过滤器，支持字段groupName，imported
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeRocketMQSourceClusterGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQSourceClusterGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "TaskId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQSourceClusterGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQSourceClusterGroupListResponseParams struct {
	// group列表
	Groups []*RocketMQGroupConfigOutput `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQSourceClusterGroupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQSourceClusterGroupListResponseParams `json:"Response"`
}

func (r *DescribeRocketMQSourceClusterGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQSourceClusterGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQSourceClusterTopicListRequestParams struct {
	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 迁移任务名
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 查询过滤器，支持字段如下
	// TopicName,
	// Type，Imported
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRocketMQSourceClusterTopicListRequest struct {
	*tchttp.BaseRequest
	
	// 分页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 迁移任务名
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 查询过滤器，支持字段如下
	// TopicName,
	// Type，Imported
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeRocketMQSourceClusterTopicListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQSourceClusterTopicListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "TaskId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQSourceClusterTopicListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQSourceClusterTopicListResponseParams struct {
	// topic层列表
	Topics []*RocketMQTopicConfigOutput `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQSourceClusterTopicListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQSourceClusterTopicListResponseParams `json:"Response"`
}

func (r *DescribeRocketMQSourceClusterTopicListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQSourceClusterTopicListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQSubscriptionsRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 消费组名称
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询限制条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeRocketMQSubscriptionsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 消费组名称
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询限制条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeRocketMQSubscriptionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQSubscriptionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Namespace")
	delete(f, "Group")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQSubscriptionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQSubscriptionsResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 订阅关系列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Subscriptions []*RocketMQSubscription `json:"Subscriptions,omitnil,omitempty" name:"Subscriptions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQSubscriptionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQSubscriptionsResponseParams `json:"Response"`
}

func (r *DescribeRocketMQSubscriptionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQSubscriptionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQTopUsagesRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 指标名称，支持以下：
	// consumeLag，消费组堆积数量
	// deadLetterCount，死信数量
	// topicRateIn,   Topic生产速率
	// topicRateOut，Topic消费速率
	// topicStorageSize，Topic存储空间
	// topicApiCalls，Topic API调用次数
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 排序数量，最大20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeRocketMQTopUsagesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 指标名称，支持以下：
	// consumeLag，消费组堆积数量
	// deadLetterCount，死信数量
	// topicRateIn,   Topic生产速率
	// topicRateOut，Topic消费速率
	// topicStorageSize，Topic存储空间
	// topicApiCalls，Topic API调用次数
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 排序数量，最大20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeRocketMQTopUsagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQTopUsagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "MetricName")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQTopUsagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQTopUsagesResponseParams struct {
	// 指标值列表
	Values []*int64 `json:"Values,omitnil,omitempty" name:"Values"`

	// 指标值对应的维度组合，本接口存在以下几个维度：
	// tenant，namespace，group，topic
	Dimensions []*DimensionInstance `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQTopUsagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQTopUsagesResponseParams `json:"Response"`
}

func (r *DescribeRocketMQTopUsagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQTopUsagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQTopicMsgsRequestParams struct {
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名称，查询死信时为groupId
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 消息 ID
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 消息 key
	MsgKey *string `json:"MsgKey,omitnil,omitempty" name:"MsgKey"`

	// 查询偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询限额
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 标志一次分页事务
	TaskRequestId *string `json:"TaskRequestId,omitnil,omitempty" name:"TaskRequestId"`

	// 死信查询时该值为true，只对Rocketmq有效
	//
	// Deprecated: QueryDlqMsg is deprecated.
	QueryDlqMsg *bool `json:"QueryDlqMsg,omitnil,omitempty" name:"QueryDlqMsg"`

	// 查询最近N条消息 最大不超过1024，默认-1为其他查询条件
	NumOfLatestMsg *int64 `json:"NumOfLatestMsg,omitnil,omitempty" name:"NumOfLatestMsg"`

	// TAG表达式
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 死信查询时该值为true，只对Rocketmq有效
	QueryDeadLetterMessage *bool `json:"QueryDeadLetterMessage,omitnil,omitempty" name:"QueryDeadLetterMessage"`
}

type DescribeRocketMQTopicMsgsRequest struct {
	*tchttp.BaseRequest
	
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名称，查询死信时为groupId
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 消息 ID
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 消息 key
	MsgKey *string `json:"MsgKey,omitnil,omitempty" name:"MsgKey"`

	// 查询偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询限额
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 标志一次分页事务
	TaskRequestId *string `json:"TaskRequestId,omitnil,omitempty" name:"TaskRequestId"`

	// 死信查询时该值为true，只对Rocketmq有效
	QueryDlqMsg *bool `json:"QueryDlqMsg,omitnil,omitempty" name:"QueryDlqMsg"`

	// 查询最近N条消息 最大不超过1024，默认-1为其他查询条件
	NumOfLatestMsg *int64 `json:"NumOfLatestMsg,omitnil,omitempty" name:"NumOfLatestMsg"`

	// TAG表达式
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 死信查询时该值为true，只对Rocketmq有效
	QueryDeadLetterMessage *bool `json:"QueryDeadLetterMessage,omitnil,omitempty" name:"QueryDeadLetterMessage"`
}

func (r *DescribeRocketMQTopicMsgsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQTopicMsgsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MsgId")
	delete(f, "MsgKey")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TaskRequestId")
	delete(f, "QueryDlqMsg")
	delete(f, "NumOfLatestMsg")
	delete(f, "Tag")
	delete(f, "QueryDeadLetterMessage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQTopicMsgsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQTopicMsgsResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 消息列表
	TopicMsgLogSets []*RocketMQMsgLog `json:"TopicMsgLogSets,omitnil,omitempty" name:"TopicMsgLogSets"`

	// 标志一次分页事务
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskRequestId *string `json:"TaskRequestId,omitnil,omitempty" name:"TaskRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQTopicMsgsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQTopicMsgsResponseParams `json:"Response"`
}

func (r *DescribeRocketMQTopicMsgsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQTopicMsgsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQTopicStatsRequestParams struct {
	// 实例ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 主题名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type DescribeRocketMQTopicStatsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 主题名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

func (r *DescribeRocketMQTopicStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQTopicStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQTopicStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQTopicStatsResponseParams struct {
	// 生产详情列表
	TopicStatsList []*TopicStats `json:"TopicStatsList,omitnil,omitempty" name:"TopicStatsList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQTopicStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQTopicStatsResponseParams `json:"Response"`
}

func (r *DescribeRocketMQTopicStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQTopicStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQTopicsByGroupRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 消费组名称
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeRocketMQTopicsByGroupRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 消费组名称
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeRocketMQTopicsByGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQTopicsByGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "GroupId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQTopicsByGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQTopicsByGroupResponseParams struct {
	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 主题列表
	Topics []*string `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQTopicsByGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQTopicsByGroupResponseParams `json:"Response"`
}

func (r *DescribeRocketMQTopicsByGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQTopicsByGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQTopicsRequestParams struct {
	// 查询偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询限制数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 按主题类型过滤查询结果，可选择Normal, GlobalOrder, PartitionedOrder, Transaction
	FilterType []*string `json:"FilterType,omitnil,omitempty" name:"FilterType"`

	// 按主题名称搜索，支持模糊查询
	FilterName *string `json:"FilterName,omitnil,omitempty" name:"FilterName"`

	// 按订阅消费组名称过滤
	FilterGroup *string `json:"FilterGroup,omitnil,omitempty" name:"FilterGroup"`
}

type DescribeRocketMQTopicsRequest struct {
	*tchttp.BaseRequest
	
	// 查询偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询限制数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 按主题类型过滤查询结果，可选择Normal, GlobalOrder, PartitionedOrder, Transaction
	FilterType []*string `json:"FilterType,omitnil,omitempty" name:"FilterType"`

	// 按主题名称搜索，支持模糊查询
	FilterName *string `json:"FilterName,omitnil,omitempty" name:"FilterName"`

	// 按订阅消费组名称过滤
	FilterGroup *string `json:"FilterGroup,omitnil,omitempty" name:"FilterGroup"`
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
	delete(f, "FilterGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQTopicsResponseParams struct {
	// 总记录数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 主题信息列表
	Topics []*RocketMQTopic `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQTopicsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQTopicsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeRocketMQVipInstanceDetailRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeRocketMQVipInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeRocketMQVipInstanceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQVipInstanceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQVipInstanceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQVipInstanceDetailResponseParams struct {
	// 集群信息
	ClusterInfo *RocketMQClusterInfo `json:"ClusterInfo,omitnil,omitempty" name:"ClusterInfo"`

	// 集群配置
	InstanceConfig *RocketMQInstanceConfig `json:"InstanceConfig,omitnil,omitempty" name:"InstanceConfig"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQVipInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQVipInstanceDetailResponseParams `json:"Response"`
}

func (r *DescribeRocketMQVipInstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQVipInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQVipInstancesRequestParams struct {
	// 查询条件过滤器，支持的查询条件如下：
	// instanceIds - 实例ID
	// instanceName - 实例名称
	// status - 实例状态
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询数目上限，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询起始位置
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeRocketMQVipInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 查询条件过滤器，支持的查询条件如下：
	// instanceIds - 实例ID
	// instanceName - 实例名称
	// status - 实例状态
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询数目上限，默认20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询起始位置
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeRocketMQVipInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQVipInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQVipInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQVipInstancesResponseParams struct {
	// 未分页的总数目
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例信息列表
	Instances []*RocketMQVipInstance `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRocketMQVipInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQVipInstancesResponseParams `json:"Response"`
}

func (r *DescribeRocketMQVipInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQVipInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRolesRequestParams struct {
	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 角色名称，模糊查询
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 起始下标，不填默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// * RoleName
	// 按照角色名进行过滤，精确查询。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRolesRequest struct {
	*tchttp.BaseRequest
	
	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 角色名称，模糊查询
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 起始下标，不填默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// * RoleName
	// 按照角色名进行过滤，精确查询。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	delete(f, "ClusterId")
	delete(f, "RoleName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRolesResponseParams struct {
	// 记录数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 角色数组。
	RoleSets []*Role `json:"RoleSets,omitnil,omitempty" name:"RoleSets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRolesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRolesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeSubscriptionsRequestParams struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 订阅者名称，模糊匹配。
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`

	// 数据过滤条件。
	Filters []*FilterSubscription `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeSubscriptionsRequest struct {
	*tchttp.BaseRequest
	
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 订阅者名称，模糊匹配。
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`

	// 数据过滤条件。
	Filters []*FilterSubscription `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SubscriptionName")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubscriptionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubscriptionsResponseParams struct {
	// 订阅者集合数组。
	SubscriptionSets []*Subscription `json:"SubscriptionSets,omitnil,omitempty" name:"SubscriptionSets"`

	// 数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSubscriptionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubscriptionsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeTopicMsgsRequestParams struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 消息ID。
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type DescribeTopicMsgsRequest struct {
	*tchttp.BaseRequest
	
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 消息ID。
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

func (r *DescribeTopicMsgsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicMsgsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "MsgId")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicMsgsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicMsgsResponseParams struct {
	// 总记录数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 消息日志列表。
	TopicMsgLogSets []*MsgLog `json:"TopicMsgLogSets,omitnil,omitempty" name:"TopicMsgLogSets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicMsgsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicMsgsResponseParams `json:"Response"`
}

func (r *DescribeTopicMsgsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicMsgsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicsRequestParams struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 主题名模糊匹配。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// topic类型描述：
	// 0：非持久非分区主题类型；
	// 1：非持久分区主题类型；
	// 2：持久非分区主题类型；
	// 3：持久分区主题类型；
	TopicType *uint64 `json:"TopicType,omitnil,omitempty" name:"TopicType"`

	// * TopicName
	// 按照主题名字查询，精确查询。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 创建来源：
	// 1：用户创建
	// 2：系统创建
	TopicCreator *uint64 `json:"TopicCreator,omitnil,omitempty" name:"TopicCreator"`
}

type DescribeTopicsRequest struct {
	*tchttp.BaseRequest
	
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 主题名模糊匹配。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// topic类型描述：
	// 0：非持久非分区主题类型；
	// 1：非持久分区主题类型；
	// 2：持久非分区主题类型；
	// 3：持久分区主题类型；
	TopicType *uint64 `json:"TopicType,omitnil,omitempty" name:"TopicType"`

	// * TopicName
	// 按照主题名字查询，精确查询。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 创建来源：
	// 1：用户创建
	// 2：系统创建
	TopicCreator *uint64 `json:"TopicCreator,omitnil,omitempty" name:"TopicCreator"`
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
	delete(f, "ClusterId")
	delete(f, "TopicName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TopicType")
	delete(f, "Filters")
	delete(f, "TopicCreator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicsResponseParams struct {
	// 主题集合数组。
	TopicSets []*Topic `json:"TopicSets,omitnil,omitempty" name:"TopicSets"`

	// 主题数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicsResponseParams `json:"Response"`
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

type DetailedRolePerm struct {
	// 权限对应的资源
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 是否开启生产权限
	PermWrite *bool `json:"PermWrite,omitnil,omitempty" name:"PermWrite"`

	// 是否开启消费权限
	PermRead *bool `json:"PermRead,omitnil,omitempty" name:"PermRead"`

	// 授权资源类型（Topic:主题; Group:消费组）
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type DimensionInstance struct {
	// 实例的维度组合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dimensions []*DimensionOpt `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`
}

type DimensionOpt struct {
	// 查询的维度名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 查询维度的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Environment struct {
	// 命名空间名称
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 未消费消息过期时间，单位：秒，最大1296000（15天）
	MsgTTL *int64 `json:"MsgTTL,omitnil,omitempty" name:"MsgTTL"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最近修改时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 命名空间ID
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitnil,omitempty" name:"NamespaceName"`

	// Topic数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// 消息保留策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetentionPolicy *RetentionPolicy `json:"RetentionPolicy,omitnil,omitempty" name:"RetentionPolicy"`

	// 是否自动创建订阅
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoSubscriptionCreation *bool `json:"AutoSubscriptionCreation,omitnil,omitempty" name:"AutoSubscriptionCreation"`
}

type EnvironmentRole struct {
	// 环境（命名空间）。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 角色名称。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 授权项，最多只能包含produce、consume两项的非空字符串数组。
	Permissions []*string `json:"Permissions,omitnil,omitempty" name:"Permissions"`

	// 角色描述。
	RoleDescribe *string `json:"RoleDescribe,omitnil,omitempty" name:"RoleDescribe"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type EnvironmentRoleSet struct {
	// 需要绑定的命名空间Id，不重复且存在资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 名字空间需要绑定的权限，枚举为 "consume" "produce" 组合，但是不为空
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permissions []*string `json:"Permissions,omitnil,omitempty" name:"Permissions"`
}

type ExchangeQuota struct {
	// 可创建最大exchange数
	MaxExchange *int64 `json:"MaxExchange,omitnil,omitempty" name:"MaxExchange"`

	// 已创建exchange数
	UsedExchange *int64 `json:"UsedExchange,omitnil,omitempty" name:"UsedExchange"`
}

// Predefined struct for user
type ExecuteDisasterRecoveryRequestParams struct {

}

type ExecuteDisasterRecoveryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ExecuteDisasterRecoveryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteDisasterRecoveryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecuteDisasterRecoveryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteDisasterRecoveryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExecuteDisasterRecoveryResponse struct {
	*tchttp.BaseResponse
	Response *ExecuteDisasterRecoveryResponseParams `json:"Response"`
}

func (r *ExecuteDisasterRecoveryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteDisasterRecoveryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportRocketMQMessageDetailRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 应用命名空间
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Topic名称
	// 如果是死信消息 isDlqMsg=true
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 消息id
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 是否包含消息体
	IncludeMsgBody *bool `json:"IncludeMsgBody,omitnil,omitempty" name:"IncludeMsgBody"`

	// 是否死信消息
	DeadLetterMsg *bool `json:"DeadLetterMsg,omitnil,omitempty" name:"DeadLetterMsg"`
}

type ExportRocketMQMessageDetailRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 应用命名空间
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Topic名称
	// 如果是死信消息 isDlqMsg=true
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 消息id
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 是否包含消息体
	IncludeMsgBody *bool `json:"IncludeMsgBody,omitnil,omitempty" name:"IncludeMsgBody"`

	// 是否死信消息
	DeadLetterMsg *bool `json:"DeadLetterMsg,omitnil,omitempty" name:"DeadLetterMsg"`
}

func (r *ExportRocketMQMessageDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportRocketMQMessageDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "MsgId")
	delete(f, "IncludeMsgBody")
	delete(f, "DeadLetterMsg")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportRocketMQMessageDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportRocketMQMessageDetailResponseParams struct {
	// 消息id
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 消息生成时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	BornTimestamp *int64 `json:"BornTimestamp,omitnil,omitempty" name:"BornTimestamp"`

	// 消息存储时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoreTimestamp *int64 `json:"StoreTimestamp,omitnil,omitempty" name:"StoreTimestamp"`

	// 消息生产客户端地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	BornHost *string `json:"BornHost,omitnil,omitempty" name:"BornHost"`

	// 消息Tag
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgTag *string `json:"MsgTag,omitnil,omitempty" name:"MsgTag"`

	// 消息Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgKey *string `json:"MsgKey,omitnil,omitempty" name:"MsgKey"`

	// 消息属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties *string `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 消息重试次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReConsumeTimes *uint64 `json:"ReConsumeTimes,omitnil,omitempty" name:"ReConsumeTimes"`

	// Base64编码格式字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgBody *string `json:"MsgBody,omitnil,omitempty" name:"MsgBody"`

	// 消息内容的CRC32 Code
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgBodyCRC *int64 `json:"MsgBodyCRC,omitnil,omitempty" name:"MsgBodyCRC"`

	// 消息体大小（单位K）
	// 当大于2048时不返回消息
	MsgBodySize *uint64 `json:"MsgBodySize,omitnil,omitempty" name:"MsgBodySize"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportRocketMQMessageDetailResponse struct {
	*tchttp.BaseResponse
	Response *ExportRocketMQMessageDetailResponseParams `json:"Response"`
}

func (r *ExportRocketMQMessageDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportRocketMQMessageDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 过滤参数的名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FilterSubscription struct {
	// 是否仅展示包含真实消费者的订阅。
	ConsumerHasCount *bool `json:"ConsumerHasCount,omitnil,omitempty" name:"ConsumerHasCount"`

	// 是否仅展示消息堆积的订阅。
	ConsumerHasBacklog *bool `json:"ConsumerHasBacklog,omitnil,omitempty" name:"ConsumerHasBacklog"`

	// 是否仅展示存在消息超期丢弃的订阅。
	ConsumerHasExpired *bool `json:"ConsumerHasExpired,omitnil,omitempty" name:"ConsumerHasExpired"`

	// 按照订阅名过滤，精确查询。
	SubscriptionNames []*string `json:"SubscriptionNames,omitnil,omitempty" name:"SubscriptionNames"`
}

// Predefined struct for user
type GetTopicListRequestParams struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type GetTopicListRequest struct {
	*tchttp.BaseRequest
	
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *GetTopicListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTopicListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTopicListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTopicListResponseParams struct {
	// 主题数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 主题列表
	TopicList []*Topic_Simplification `json:"TopicList,omitnil,omitempty" name:"TopicList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTopicListResponse struct {
	*tchttp.BaseResponse
	Response *GetTopicListResponseParams `json:"Response"`
}

func (r *GetTopicListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTopicListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportRocketMQConsumerGroupsRequestParams struct {
	// 导入topic
	Groups []*RocketMQGroupConfig `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type ImportRocketMQConsumerGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 导入topic
	Groups []*RocketMQGroupConfig `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *ImportRocketMQConsumerGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportRocketMQConsumerGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Groups")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportRocketMQConsumerGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportRocketMQConsumerGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImportRocketMQConsumerGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ImportRocketMQConsumerGroupsResponseParams `json:"Response"`
}

func (r *ImportRocketMQConsumerGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportRocketMQConsumerGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportRocketMQTopicsRequestParams struct {
	// 导入topic
	Topics []*RocketMQTopicConfig `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type ImportRocketMQTopicsRequest struct {
	*tchttp.BaseRequest
	
	// 导入topic
	Topics []*RocketMQTopicConfig `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *ImportRocketMQTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportRocketMQTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Topics")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportRocketMQTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportRocketMQTopicsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImportRocketMQTopicsResponse struct {
	*tchttp.BaseResponse
	Response *ImportRocketMQTopicsResponseParams `json:"Response"`
}

func (r *ImportRocketMQTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportRocketMQTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceNodeDistribution struct {
	// 可用区
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 可用区id
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 节点数
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 有调度任务且没有切回的可用区，此标识为true
	NodePermWipeFlag *bool `json:"NodePermWipeFlag,omitnil,omitempty" name:"NodePermWipeFlag"`

	// 可用区状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneStatus *string `json:"ZoneStatus,omitnil,omitempty" name:"ZoneStatus"`
}

type InternalTenant struct {
	// 虚拟集群ID
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 虚拟集群名称
	TenantName *string `json:"TenantName,omitnil,omitempty" name:"TenantName"`

	// 客户UIN
	CustomerUin *string `json:"CustomerUin,omitnil,omitempty" name:"CustomerUin"`

	// 客户的APPID
	CustomerAppId *string `json:"CustomerAppId,omitnil,omitempty" name:"CustomerAppId"`

	// 物理集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 集群协议类型，支持的值为TDMQ，ROCKETMQ，AMQP，CMQ
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 命名空间配额
	MaxNamespaces *int64 `json:"MaxNamespaces,omitnil,omitempty" name:"MaxNamespaces"`

	// 已使用命名空间配额
	UsedNamespaces *int64 `json:"UsedNamespaces,omitnil,omitempty" name:"UsedNamespaces"`

	// Topic配额
	MaxTopics *int64 `json:"MaxTopics,omitnil,omitempty" name:"MaxTopics"`

	// 已使用Topic配额
	UsedTopics *int64 `json:"UsedTopics,omitnil,omitempty" name:"UsedTopics"`

	// Topic分区数配额
	MaxPartitions *int64 `json:"MaxPartitions,omitnil,omitempty" name:"MaxPartitions"`

	// 已使用Topic分区数配额
	UsedPartitions *int64 `json:"UsedPartitions,omitnil,omitempty" name:"UsedPartitions"`

	// 存储配额, byte为单位
	MaxMsgBacklogSize *uint64 `json:"MaxMsgBacklogSize,omitnil,omitempty" name:"MaxMsgBacklogSize"`

	// 命名空间最大生产TPS
	MaxPublishTps *uint64 `json:"MaxPublishTps,omitnil,omitempty" name:"MaxPublishTps"`

	// 消息最大保留时间，秒为单位
	MaxRetention *uint64 `json:"MaxRetention,omitnil,omitempty" name:"MaxRetention"`

	// 创建时间，毫秒为单位
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间，毫秒为单位
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 命名空间最大消费TPS
	MaxDispatchTps *uint64 `json:"MaxDispatchTps,omitnil,omitempty" name:"MaxDispatchTps"`

	// 命名空间最大消费带宽，byte为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDispatchRateInBytes *uint64 `json:"MaxDispatchRateInBytes,omitnil,omitempty" name:"MaxDispatchRateInBytes"`

	// 命名空间最大生产带宽，byte为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxPublishRateInBytes *uint64 `json:"MaxPublishRateInBytes,omitnil,omitempty" name:"MaxPublishRateInBytes"`

	// 消息最大保留空间，MB为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRetentionSizeInMB *uint64 `json:"MaxRetentionSizeInMB,omitnil,omitempty" name:"MaxRetentionSizeInMB"`

	// public Access Enabled
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitnil,omitempty" name:"PublicAccessEnabled"`
}

type MigrateTopic struct {
	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// topic名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 迁移状态
	// S_RW_D_NA 源集群读写
	// S_RW_D_R 源集群读写目标集群读
	// S_RW_D_RW 源集群读写目标集群读写
	// S_R_D_RW 源集群读目标集群读写
	// S_NA_D_RW 目标集群读写
	// 注意：此字段可能返回 null，表示取不到有效值。
	MigrationStatus *string `json:"MigrationStatus,omitnil,omitempty" name:"MigrationStatus"`

	// 是否完成健康检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheckPassed *bool `json:"HealthCheckPassed,omitnil,omitempty" name:"HealthCheckPassed"`

	// 上次健康检查返回的错误信息，仅在HealthCheckPassed为false时有效。
	// NotChecked 未执行检查，
	// Unknown 未知错误,
	// TopicNotImported 主题未导入,
	//  TopicNotExistsInSourceCluster  主题在源集群中不存在,
	//     TopicNotExistsInTargetCluster 主题在目标集群中不存在,
	//     ConsumerConnectedOnTarget 目标集群上存在消费者连接,
	//     SourceTopicHasNewMessagesIn5Minutes 源集群主题前5分钟内有新消息写入,
	// TargetTopicHasNewMessagesIn5Minutes 目标集群主题前5分钟内有新消息写入,
	//     SourceTopicHasNoMessagesIn5Minutes 源集群前5分钟内没有新消息写入,
	// TargetTopicHasNoMessagesIn5Minutes 源集群前5分钟内没有新消息写入,
	//     ConsumerGroupCountNotMatch 订阅组数量不一致,
	//     SourceTopicHasUnconsumedMessages 源集群主题存在未消费消息,
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheckError *string `json:"HealthCheckError,omitnil,omitempty" name:"HealthCheckError"`
}

// Predefined struct for user
type ModifyClusterRequestParams struct {
	// Pulsar 集群的ID，需要更新的集群Id。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 更新后的集群名称。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 说明信息。长度限制为 128 字节
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 开启公网访问，只能为true
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitnil,omitempty" name:"PublicAccessEnabled"`
}

type ModifyClusterRequest struct {
	*tchttp.BaseRequest
	
	// Pulsar 集群的ID，需要更新的集群Id。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 更新后的集群名称。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 说明信息。长度限制为 128 字节
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 开启公网访问，只能为true
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitnil,omitempty" name:"PublicAccessEnabled"`
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

// Predefined struct for user
type ModifyClusterResponseParams struct {
	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyCmqQueueAttributeRequestParams struct {
	// 队列名字，在单个地域同一账号下唯一。队列名称是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 最大堆积消息数。取值范围在公测期间为 1,000,000 - 10,000,000，正式上线后范围可达到 1000,000-1000,000,000。默认取值在公测期间为 10,000,000，正式上线后为 100,000,000。
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitnil,omitempty" name:"MaxMsgHeapNum"`

	// 消息接收长轮询等待时间。取值范围 0-30 秒，默认值 0。
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitnil,omitempty" name:"PollingWaitSeconds"`

	// 消息可见性超时。取值范围 1-43200 秒（即12小时内），默认值 30。
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitnil,omitempty" name:"VisibilityTimeout"`

	// 消息最大长度，新版CMQ新建的队列默认1024KB，不支持修改
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil,omitempty" name:"MaxMsgSize"`

	// 消息最长未确认时间。取值范围 30-43200 秒（30秒~12小时），默认值 3600 (1 小时)。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil,omitempty" name:"MsgRetentionSeconds"`

	// 队列是否开启回溯消息能力，该参数取值范围0-1296000，0表示不开启。
	RewindSeconds *uint64 `json:"RewindSeconds,omitnil,omitempty" name:"RewindSeconds"`

	// 第一次查询时间
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitnil,omitempty" name:"FirstQueryInterval"`

	// 最大查询次数
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitnil,omitempty" name:"MaxQueryCount"`

	// 死信队列名称
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitnil,omitempty" name:"DeadLetterQueueName"`

	// policy为1时必选。最大未消费过期时间。范围300-43200，单位秒，需要小于消息最大保留时间MsgRetentionSeconds
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitnil,omitempty" name:"MaxTimeToLive"`

	// 最大接收次数
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitnil,omitempty" name:"MaxReceiveCount"`

	// 死信队列策略
	Policy *uint64 `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 是否开启消息轨迹标识，true表示开启，false表示不开启，不填表示不开启。
	Trace *bool `json:"Trace,omitnil,omitempty" name:"Trace"`

	// 是否开启事务，1开启，0不开启
	Transaction *uint64 `json:"Transaction,omitnil,omitempty" name:"Transaction"`

	// 队列可回溯存储空间：若开启消息回溯，取值范围：10240MB - 512000MB，若不开启消息回溯，取值：0
	RetentionSizeInMB *uint64 `json:"RetentionSizeInMB,omitnil,omitempty" name:"RetentionSizeInMB"`
}

type ModifyCmqQueueAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 队列名字，在单个地域同一账号下唯一。队列名称是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 最大堆积消息数。取值范围在公测期间为 1,000,000 - 10,000,000，正式上线后范围可达到 1000,000-1000,000,000。默认取值在公测期间为 10,000,000，正式上线后为 100,000,000。
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitnil,omitempty" name:"MaxMsgHeapNum"`

	// 消息接收长轮询等待时间。取值范围 0-30 秒，默认值 0。
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitnil,omitempty" name:"PollingWaitSeconds"`

	// 消息可见性超时。取值范围 1-43200 秒（即12小时内），默认值 30。
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitnil,omitempty" name:"VisibilityTimeout"`

	// 消息最大长度，新版CMQ新建的队列默认1024KB，不支持修改
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil,omitempty" name:"MaxMsgSize"`

	// 消息最长未确认时间。取值范围 30-43200 秒（30秒~12小时），默认值 3600 (1 小时)。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil,omitempty" name:"MsgRetentionSeconds"`

	// 队列是否开启回溯消息能力，该参数取值范围0-1296000，0表示不开启。
	RewindSeconds *uint64 `json:"RewindSeconds,omitnil,omitempty" name:"RewindSeconds"`

	// 第一次查询时间
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitnil,omitempty" name:"FirstQueryInterval"`

	// 最大查询次数
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitnil,omitempty" name:"MaxQueryCount"`

	// 死信队列名称
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitnil,omitempty" name:"DeadLetterQueueName"`

	// policy为1时必选。最大未消费过期时间。范围300-43200，单位秒，需要小于消息最大保留时间MsgRetentionSeconds
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitnil,omitempty" name:"MaxTimeToLive"`

	// 最大接收次数
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitnil,omitempty" name:"MaxReceiveCount"`

	// 死信队列策略
	Policy *uint64 `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 是否开启消息轨迹标识，true表示开启，false表示不开启，不填表示不开启。
	Trace *bool `json:"Trace,omitnil,omitempty" name:"Trace"`

	// 是否开启事务，1开启，0不开启
	Transaction *uint64 `json:"Transaction,omitnil,omitempty" name:"Transaction"`

	// 队列可回溯存储空间：若开启消息回溯，取值范围：10240MB - 512000MB，若不开启消息回溯，取值：0
	RetentionSizeInMB *uint64 `json:"RetentionSizeInMB,omitnil,omitempty" name:"RetentionSizeInMB"`
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
	delete(f, "RetentionSizeInMB")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCmqQueueAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCmqQueueAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCmqQueueAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCmqQueueAttributeResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyCmqSubscriptionAttributeRequestParams struct {
	// 主题名字，在单个地域同一账号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 订阅名字，在单个地域同一账号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`

	// 向 Endpoint 推送消息出现错误时，CMQ 推送服务器的重试策略。取值如下：
	// （1）BACKOFF_RETRY，退避重试。每隔一定时间重试一次，重试够一定次数后，就把该消息丢弃，继续推送下一条消息。
	// （2）EXPONENTIAL_DECAY_RETRY，指数衰退重试。每次重试的间隔是指数递增的，例如开始1s，后面是2s，4s，8s···由于 Topic 消息的周期是一天，所以最多重试一天就把消息丢弃。默认值是 EXPONENTIAL_DECAY_RETRY。
	NotifyStrategy *string `json:"NotifyStrategy,omitnil,omitempty" name:"NotifyStrategy"`

	// 推送内容的格式。取值：（1）JSON；（2）SIMPLIFIED，即 raw 格式。如果 Protocol 是 queue，则取值必须为 SIMPLIFIED。如果 Protocol 是 HTTP，两个值均可以，默认值是 JSON。
	NotifyContentFormat *string `json:"NotifyContentFormat,omitnil,omitempty" name:"NotifyContentFormat"`

	// 消息正文。消息标签（用于消息过滤)。标签数量不能超过5个，每个标签不超过16个字符。与(Batch)PublishMessage的MsgTag参数配合使用，规则：1）如果FilterTag没有设置，则无论MsgTag是否有设置，订阅接收所有发布到Topic的消息；2）如果FilterTag数组有值，则只有数组中至少有一个值在MsgTag数组中也存在时（即FilterTag和MsgTag有交集），订阅才接收该发布到Topic的消息；3）如果FilterTag数组有值，但MsgTag没设置，则不接收任何发布到Topic的消息，可以认为是2）的一种特例，此时FilterTag和MsgTag没有交集。规则整体的设计思想是以订阅者的意愿为主。
	FilterTags []*string `json:"FilterTags,omitnil,omitempty" name:"FilterTags"`

	// BindingKey数量不超过5个， 每个BindingKey长度不超过64字节，该字段表示订阅接收消息的过滤策略，每个BindingKey最多含有15个“.”， 即最多16个词组。
	BindingKey []*string `json:"BindingKey,omitnil,omitempty" name:"BindingKey"`
}

type ModifyCmqSubscriptionAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 主题名字，在单个地域同一账号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 订阅名字，在单个地域同一账号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`

	// 向 Endpoint 推送消息出现错误时，CMQ 推送服务器的重试策略。取值如下：
	// （1）BACKOFF_RETRY，退避重试。每隔一定时间重试一次，重试够一定次数后，就把该消息丢弃，继续推送下一条消息。
	// （2）EXPONENTIAL_DECAY_RETRY，指数衰退重试。每次重试的间隔是指数递增的，例如开始1s，后面是2s，4s，8s···由于 Topic 消息的周期是一天，所以最多重试一天就把消息丢弃。默认值是 EXPONENTIAL_DECAY_RETRY。
	NotifyStrategy *string `json:"NotifyStrategy,omitnil,omitempty" name:"NotifyStrategy"`

	// 推送内容的格式。取值：（1）JSON；（2）SIMPLIFIED，即 raw 格式。如果 Protocol 是 queue，则取值必须为 SIMPLIFIED。如果 Protocol 是 HTTP，两个值均可以，默认值是 JSON。
	NotifyContentFormat *string `json:"NotifyContentFormat,omitnil,omitempty" name:"NotifyContentFormat"`

	// 消息正文。消息标签（用于消息过滤)。标签数量不能超过5个，每个标签不超过16个字符。与(Batch)PublishMessage的MsgTag参数配合使用，规则：1）如果FilterTag没有设置，则无论MsgTag是否有设置，订阅接收所有发布到Topic的消息；2）如果FilterTag数组有值，则只有数组中至少有一个值在MsgTag数组中也存在时（即FilterTag和MsgTag有交集），订阅才接收该发布到Topic的消息；3）如果FilterTag数组有值，但MsgTag没设置，则不接收任何发布到Topic的消息，可以认为是2）的一种特例，此时FilterTag和MsgTag没有交集。规则整体的设计思想是以订阅者的意愿为主。
	FilterTags []*string `json:"FilterTags,omitnil,omitempty" name:"FilterTags"`

	// BindingKey数量不超过5个， 每个BindingKey长度不超过64字节，该字段表示订阅接收消息的过滤策略，每个BindingKey最多含有15个“.”， 即最多16个词组。
	BindingKey []*string `json:"BindingKey,omitnil,omitempty" name:"BindingKey"`
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

// Predefined struct for user
type ModifyCmqSubscriptionAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCmqSubscriptionAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCmqSubscriptionAttributeResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyCmqTopicAttributeRequestParams struct {
	// 主题名字，在单个地域同一账号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 消息最大长度。取值范围1024 - 65536 Byte（即1 - 64K），默认值65536。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil,omitempty" name:"MaxMsgSize"`

	// 消息保存时间。取值范围60 - 86400 s（即1分钟 - 1天），默认值86400。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil,omitempty" name:"MsgRetentionSeconds"`

	// 是否开启消息轨迹标识，true表示开启，false表示不开启，不填表示不开启。
	Trace *bool `json:"Trace,omitnil,omitempty" name:"Trace"`
}

type ModifyCmqTopicAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 主题名字，在单个地域同一账号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 消息最大长度。取值范围1024 - 65536 Byte（即1 - 64K），默认值65536。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil,omitempty" name:"MaxMsgSize"`

	// 消息保存时间。取值范围60 - 86400 s（即1分钟 - 1天），默认值86400。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil,omitempty" name:"MsgRetentionSeconds"`

	// 是否开启消息轨迹标识，true表示开启，false表示不开启，不填表示不开启。
	Trace *bool `json:"Trace,omitnil,omitempty" name:"Trace"`
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

// Predefined struct for user
type ModifyCmqTopicAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCmqTopicAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCmqTopicAttributeResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyEnvironmentAttributesRequestParams struct {
	// 命名空间名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 未消费消息过期时间，单位：秒，范围60秒~15天。
	MsgTTL *uint64 `json:"MsgTTL,omitnil,omitempty" name:"MsgTTL"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备注，字符串最长不超过128。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 消息保留策略
	RetentionPolicy *RetentionPolicy `json:"RetentionPolicy,omitnil,omitempty" name:"RetentionPolicy"`

	// 是否开启自动创建订阅
	AutoSubscriptionCreation *bool `json:"AutoSubscriptionCreation,omitnil,omitempty" name:"AutoSubscriptionCreation"`
}

type ModifyEnvironmentAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 未消费消息过期时间，单位：秒，范围60秒~15天。
	MsgTTL *uint64 `json:"MsgTTL,omitnil,omitempty" name:"MsgTTL"`

	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备注，字符串最长不超过128。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 消息保留策略
	RetentionPolicy *RetentionPolicy `json:"RetentionPolicy,omitnil,omitempty" name:"RetentionPolicy"`

	// 是否开启自动创建订阅
	AutoSubscriptionCreation *bool `json:"AutoSubscriptionCreation,omitnil,omitempty" name:"AutoSubscriptionCreation"`
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
	delete(f, "ClusterId")
	delete(f, "Remark")
	delete(f, "RetentionPolicy")
	delete(f, "AutoSubscriptionCreation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEnvironmentAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnvironmentAttributesResponseParams struct {
	// 命名空间名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 未消费消息过期时间，单位：秒。
	MsgTTL *uint64 `json:"MsgTTL,omitnil,omitempty" name:"MsgTTL"`

	// 备注，字符串最长不超过128。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 命名空间ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEnvironmentAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEnvironmentAttributesResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyEnvironmentRoleRequestParams struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 角色名称。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 授权项，最多只能包含produce、consume两项的非空字符串数组。
	Permissions []*string `json:"Permissions,omitnil,omitempty" name:"Permissions"`

	// 必填字段，集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type ModifyEnvironmentRoleRequest struct {
	*tchttp.BaseRequest
	
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 角色名称。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 授权项，最多只能包含produce、consume两项的非空字符串数组。
	Permissions []*string `json:"Permissions,omitnil,omitempty" name:"Permissions"`

	// 必填字段，集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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

// Predefined struct for user
type ModifyEnvironmentRoleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEnvironmentRoleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEnvironmentRoleResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyPublicNetworkSecurityPolicyRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略列表
	PolicyList []*SecurityPolicy `json:"PolicyList,omitnil,omitempty" name:"PolicyList"`
}

type ModifyPublicNetworkSecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略列表
	PolicyList []*SecurityPolicy `json:"PolicyList,omitnil,omitempty" name:"PolicyList"`
}

func (r *ModifyPublicNetworkSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPublicNetworkSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PolicyList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPublicNetworkSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPublicNetworkSecurityPolicyResponseParams struct {
	// SUCCESS或者FAILURE
	ModifyResult *string `json:"ModifyResult,omitnil,omitempty" name:"ModifyResult"`

	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPublicNetworkSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPublicNetworkSecurityPolicyResponseParams `json:"Response"`
}

func (r *ModifyPublicNetworkSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPublicNetworkSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQPermissionRequestParams struct {
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，形如 admin。有效的 User 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，点击集群列表中的集群，进入集群详情，并在用户与权限页签中找到用户列表，从而找到用户名称。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到Vhost名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 权限类型，declare相关操作，该用户可操作该vhost下的资源名称正则表达式
	ConfigRegexp *string `json:"ConfigRegexp,omitnil,omitempty" name:"ConfigRegexp"`

	// 权限类型，消息写入相关操作，该用户可操作该vhost下的资源名称正则表达式
	WriteRegexp *string `json:"WriteRegexp,omitnil,omitempty" name:"WriteRegexp"`

	// 权限类型，消息读取相关操作，该用户可操作该vhost下的资源名称正则表达式
	ReadRegexp *string `json:"ReadRegexp,omitnil,omitempty" name:"ReadRegexp"`
}

type ModifyRabbitMQPermissionRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，形如 admin。有效的 User 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，点击集群列表中的集群，进入集群详情，并在用户与权限页签中找到用户列表，从而找到用户名称。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到Vhost名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 权限类型，declare相关操作，该用户可操作该vhost下的资源名称正则表达式
	ConfigRegexp *string `json:"ConfigRegexp,omitnil,omitempty" name:"ConfigRegexp"`

	// 权限类型，消息写入相关操作，该用户可操作该vhost下的资源名称正则表达式
	WriteRegexp *string `json:"WriteRegexp,omitnil,omitempty" name:"WriteRegexp"`

	// 权限类型，消息读取相关操作，该用户可操作该vhost下的资源名称正则表达式
	ReadRegexp *string `json:"ReadRegexp,omitnil,omitempty" name:"ReadRegexp"`
}

func (r *ModifyRabbitMQPermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQPermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "VirtualHost")
	delete(f, "ConfigRegexp")
	delete(f, "WriteRegexp")
	delete(f, "ReadRegexp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRabbitMQPermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQPermissionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRabbitMQPermissionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRabbitMQPermissionResponseParams `json:"Response"`
}

func (r *ModifyRabbitMQPermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQPermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQUserRequestParams struct {
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，形如 admin。有效的 User 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，点击集群列表中的集群，进入集群详情，并在用户与权限页签中找到用户列表，从而找到用户名称。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 密码，登录时使用。规范：不能为空，8-64个字符，至少要包含小写字母、大写字母、数字、特殊字符【()`~!@#$%^&*_=|{}[]:;',.?/】中的两项
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 描述，不传则不修改
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户标签，用于决定改用户访问 RabbitMQ Management 的权限范围
	// management：普通控制台用户，monitoring：管理型控制台用户，其他值：非控制台用户
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 该用户的最大连接数，不传则不修改
	MaxConnections *int64 `json:"MaxConnections,omitnil,omitempty" name:"MaxConnections"`

	// 该用户的最大channel数，不传则不修改
	MaxChannels *int64 `json:"MaxChannels,omitnil,omitempty" name:"MaxChannels"`
}

type ModifyRabbitMQUserRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，形如 admin。有效的 User 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，点击集群列表中的集群，进入集群详情，并在用户与权限页签中找到用户列表，从而找到用户名称。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 密码，登录时使用。规范：不能为空，8-64个字符，至少要包含小写字母、大写字母、数字、特殊字符【()`~!@#$%^&*_=|{}[]:;',.?/】中的两项
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 描述，不传则不修改
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户标签，用于决定改用户访问 RabbitMQ Management 的权限范围
	// management：普通控制台用户，monitoring：管理型控制台用户，其他值：非控制台用户
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 该用户的最大连接数，不传则不修改
	MaxConnections *int64 `json:"MaxConnections,omitnil,omitempty" name:"MaxConnections"`

	// 该用户的最大channel数，不传则不修改
	MaxChannels *int64 `json:"MaxChannels,omitnil,omitempty" name:"MaxChannels"`
}

func (r *ModifyRabbitMQUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "Password")
	delete(f, "Description")
	delete(f, "Tags")
	delete(f, "MaxConnections")
	delete(f, "MaxChannels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRabbitMQUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRabbitMQUserResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRabbitMQUserResponseParams `json:"Response"`
}

func (r *ModifyRabbitMQUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQVipInstanceRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群名称，不填则不修改。非空字符串时必须 3-64 个字符，只能包含数字、字母、“-”和“_”
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 备注，不填则不修改
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否开启删除保护，不填则不修改
	EnableDeletionProtection *bool `json:"EnableDeletionProtection,omitnil,omitempty" name:"EnableDeletionProtection"`
}

type ModifyRabbitMQVipInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群名称，不填则不修改。非空字符串时必须 3-64 个字符，只能包含数字、字母、“-”和“_”
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 备注，不填则不修改
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否开启删除保护，不填则不修改
	EnableDeletionProtection *bool `json:"EnableDeletionProtection,omitnil,omitempty" name:"EnableDeletionProtection"`
}

func (r *ModifyRabbitMQVipInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQVipInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterName")
	delete(f, "Remark")
	delete(f, "EnableDeletionProtection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRabbitMQVipInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQVipInstanceResponseParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRabbitMQVipInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRabbitMQVipInstanceResponseParams `json:"Response"`
}

func (r *ModifyRabbitMQVipInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQVipInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQVirtualHostRequestParams struct {
	// 实例 ID，形如amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到Vhost名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// Virtual Host 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 消息轨迹开关,true打开,false关闭
	TraceFlag *bool `json:"TraceFlag,omitnil,omitempty" name:"TraceFlag"`
}

type ModifyRabbitMQVirtualHostRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID，形如amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到Vhost名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// Virtual Host 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 消息轨迹开关,true打开,false关闭
	TraceFlag *bool `json:"TraceFlag,omitnil,omitempty" name:"TraceFlag"`
}

func (r *ModifyRabbitMQVirtualHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQVirtualHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "Description")
	delete(f, "TraceFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRabbitMQVirtualHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQVirtualHostResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRabbitMQVirtualHostResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRabbitMQVirtualHostResponseParams `json:"Response"`
}

func (r *ModifyRabbitMQVirtualHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQVirtualHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRocketMQClusterRequestParams struct {
	// RocketMQ集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 3-64个字符，只能包含字母、数字、“-”及“_”
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 说明信息，不超过128个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否开启HTTP公网访问
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitnil,omitempty" name:"PublicAccessEnabled"`
}

type ModifyRocketMQClusterRequest struct {
	*tchttp.BaseRequest
	
	// RocketMQ集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 3-64个字符，只能包含字母、数字、“-”及“_”
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 说明信息，不超过128个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否开启HTTP公网访问
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitnil,omitempty" name:"PublicAccessEnabled"`
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
	delete(f, "PublicAccessEnabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRocketMQClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRocketMQClusterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRocketMQClusterResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRocketMQClusterResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyRocketMQEnvironmentRoleRequestParams struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 角色名称。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 授权项，最多只能包含produce、consume两项的非空字符串数组。
	Permissions []*string `json:"Permissions,omitnil,omitempty" name:"Permissions"`

	// 必填字段，集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Topic&Group维度权限配置
	DetailedPerms []*DetailedRolePerm `json:"DetailedPerms,omitnil,omitempty" name:"DetailedPerms"`
}

type ModifyRocketMQEnvironmentRoleRequest struct {
	*tchttp.BaseRequest
	
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 角色名称。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 授权项，最多只能包含produce、consume两项的非空字符串数组。
	Permissions []*string `json:"Permissions,omitnil,omitempty" name:"Permissions"`

	// 必填字段，集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// Topic&Group维度权限配置
	DetailedPerms []*DetailedRolePerm `json:"DetailedPerms,omitnil,omitempty" name:"DetailedPerms"`
}

func (r *ModifyRocketMQEnvironmentRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQEnvironmentRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "RoleName")
	delete(f, "Permissions")
	delete(f, "ClusterId")
	delete(f, "DetailedPerms")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRocketMQEnvironmentRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRocketMQEnvironmentRoleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRocketMQEnvironmentRoleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRocketMQEnvironmentRoleResponseParams `json:"Response"`
}

func (r *ModifyRocketMQEnvironmentRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQEnvironmentRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRocketMQGroupRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 消费组名称
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 说明信息，最长128个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否开启消费
	ReadEnable *bool `json:"ReadEnable,omitnil,omitempty" name:"ReadEnable"`

	// 是否开启广播消费
	BroadcastEnable *bool `json:"BroadcastEnable,omitnil,omitempty" name:"BroadcastEnable"`

	// 最大重试次数
	RetryMaxTimes *uint64 `json:"RetryMaxTimes,omitnil,omitempty" name:"RetryMaxTimes"`
}

type ModifyRocketMQGroupRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 消费组名称
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 说明信息，最长128个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否开启消费
	ReadEnable *bool `json:"ReadEnable,omitnil,omitempty" name:"ReadEnable"`

	// 是否开启广播消费
	BroadcastEnable *bool `json:"BroadcastEnable,omitnil,omitempty" name:"BroadcastEnable"`

	// 最大重试次数
	RetryMaxTimes *uint64 `json:"RetryMaxTimes,omitnil,omitempty" name:"RetryMaxTimes"`
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
	delete(f, "RetryMaxTimes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRocketMQGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRocketMQGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRocketMQGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRocketMQGroupResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyRocketMQInstanceRequestParams struct {
	// 专享实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 实例消息保留时间，小时为单位
	MessageRetention *int64 `json:"MessageRetention,omitnil,omitempty" name:"MessageRetention"`

	// 是否开启删除保护
	EnableDeletionProtection *bool `json:"EnableDeletionProtection,omitnil,omitempty" name:"EnableDeletionProtection"`
}

type ModifyRocketMQInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 专享实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 实例消息保留时间，小时为单位
	MessageRetention *int64 `json:"MessageRetention,omitnil,omitempty" name:"MessageRetention"`

	// 是否开启删除保护
	EnableDeletionProtection *bool `json:"EnableDeletionProtection,omitnil,omitempty" name:"EnableDeletionProtection"`
}

func (r *ModifyRocketMQInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "Remark")
	delete(f, "MessageRetention")
	delete(f, "EnableDeletionProtection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRocketMQInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRocketMQInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRocketMQInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRocketMQInstanceResponseParams `json:"Response"`
}

func (r *ModifyRocketMQInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRocketMQInstanceSpecRequestParams struct {
	// 专享实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例规格，
	// rocket-vip-basic-1 基础型
	// rocket-vip-basic-2 标准型
	// rocket-vip-basic-3 高阶Ⅰ型
	// rocket-vip-basic-4 高阶Ⅱ型
	Specification *string `json:"Specification,omitnil,omitempty" name:"Specification"`

	// 节点数量
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 存储空间，GB为单位
	StorageSize *uint64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`
}

type ModifyRocketMQInstanceSpecRequest struct {
	*tchttp.BaseRequest
	
	// 专享实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例规格，
	// rocket-vip-basic-1 基础型
	// rocket-vip-basic-2 标准型
	// rocket-vip-basic-3 高阶Ⅰ型
	// rocket-vip-basic-4 高阶Ⅱ型
	Specification *string `json:"Specification,omitnil,omitempty" name:"Specification"`

	// 节点数量
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 存储空间，GB为单位
	StorageSize *uint64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`
}

func (r *ModifyRocketMQInstanceSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQInstanceSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Specification")
	delete(f, "NodeCount")
	delete(f, "StorageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRocketMQInstanceSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRocketMQInstanceSpecResponseParams struct {
	// 订单号
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRocketMQInstanceSpecResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRocketMQInstanceSpecResponseParams `json:"Response"`
}

func (r *ModifyRocketMQInstanceSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQInstanceSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRocketMQNamespaceRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称，3-64个字符，只能包含字母、数字、“-”及“_”
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 已废弃
	Ttl *uint64 `json:"Ttl,omitnil,omitempty" name:"Ttl"`

	// 已废弃
	RetentionTime *uint64 `json:"RetentionTime,omitnil,omitempty" name:"RetentionTime"`

	// 说明，最大128个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否开启公网访问
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitnil,omitempty" name:"PublicAccessEnabled"`
}

type ModifyRocketMQNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称，3-64个字符，只能包含字母、数字、“-”及“_”
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 已废弃
	Ttl *uint64 `json:"Ttl,omitnil,omitempty" name:"Ttl"`

	// 已废弃
	RetentionTime *uint64 `json:"RetentionTime,omitnil,omitempty" name:"RetentionTime"`

	// 说明，最大128个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否开启公网访问
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitnil,omitempty" name:"PublicAccessEnabled"`
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
	delete(f, "PublicAccessEnabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRocketMQNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRocketMQNamespaceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRocketMQNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRocketMQNamespaceResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyRocketMQRoleRequestParams struct {
	// 角色名称，不支持中字以及除了短线和下划线外的特殊字符且长度必须大于0且小等于32。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备注说明，长度必须大等于0且小等于128。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 权限类型，默认按集群授权（Cluster：集群级别；TopicAndGroup：主题&消费组级别）
	PermType *string `json:"PermType,omitnil,omitempty" name:"PermType"`
}

type ModifyRocketMQRoleRequest struct {
	*tchttp.BaseRequest
	
	// 角色名称，不支持中字以及除了短线和下划线外的特殊字符且长度必须大于0且小等于32。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备注说明，长度必须大等于0且小等于128。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 权限类型，默认按集群授权（Cluster：集群级别；TopicAndGroup：主题&消费组级别）
	PermType *string `json:"PermType,omitnil,omitempty" name:"PermType"`
}

func (r *ModifyRocketMQRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleName")
	delete(f, "ClusterId")
	delete(f, "Remark")
	delete(f, "PermType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRocketMQRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRocketMQRoleResponseParams struct {
	// 角色名称
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 备注说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRocketMQRoleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRocketMQRoleResponseParams `json:"Response"`
}

func (r *ModifyRocketMQRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRocketMQTopicRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 说明信息，最大128个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 分区数，全局类型无效，不可小于当前分区数
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`
}

type ModifyRocketMQTopicRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 说明信息，最大128个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 分区数，全局类型无效，不可小于当前分区数
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`
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

// Predefined struct for user
type ModifyRocketMQTopicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRocketMQTopicResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRocketMQTopicResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyRoleRequestParams struct {
	// 角色名称，不支持中字以及除了短线和下划线外的特殊字符且长度必须大于0且小等于32。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备注说明，长度必须大等于0且小等于128。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 批量绑定名字空间信息
	EnvironmentRoleSets []*EnvironmentRoleSet `json:"EnvironmentRoleSets,omitnil,omitempty" name:"EnvironmentRoleSets"`

	// 全部解绑名字空间，设置为 true
	UnbindAllEnvironment *bool `json:"UnbindAllEnvironment,omitnil,omitempty" name:"UnbindAllEnvironment"`
}

type ModifyRoleRequest struct {
	*tchttp.BaseRequest
	
	// 角色名称，不支持中字以及除了短线和下划线外的特殊字符且长度必须大于0且小等于32。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 必填字段，集群Id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备注说明，长度必须大等于0且小等于128。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 批量绑定名字空间信息
	EnvironmentRoleSets []*EnvironmentRoleSet `json:"EnvironmentRoleSets,omitnil,omitempty" name:"EnvironmentRoleSets"`

	// 全部解绑名字空间，设置为 true
	UnbindAllEnvironment *bool `json:"UnbindAllEnvironment,omitnil,omitempty" name:"UnbindAllEnvironment"`
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
	delete(f, "ClusterId")
	delete(f, "Remark")
	delete(f, "EnvironmentRoleSets")
	delete(f, "UnbindAllEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoleResponseParams struct {
	// 角色名称
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 备注说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRoleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRoleResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyTopicRequestParams struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 分区数，必须大于或者等于原分区数，若想维持原分区数请输入原数目，修改分区数仅对非全局顺序消息起效果，不允许超过32个分区。
	Partitions *uint64 `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备注，128字符以内。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 未消费消息过期时间，单位：秒，取值范围：60秒~15天。
	MsgTTL *uint64 `json:"MsgTTL,omitnil,omitempty" name:"MsgTTL"`

	// 不传默认是原生策略，DefaultPolicy表示当订阅下达到最大未确认消息数 5000 时，服务端将不再向当前订阅下的所有消费者推送消息，DynamicPolicy表示动态调整订阅下的最大未确认消息数，具体配额是在 5000 和消费者数量*20之间取最大值。每个消费者默认最大 unack 消息数为 20，超过该限制时仅影响该消费者，不影响其他消费者。
	UnackPolicy *string `json:"UnackPolicy,omitnil,omitempty" name:"UnackPolicy"`

	// 是否开启异常消费者隔离
	IsolateConsumerEnable *bool `json:"IsolateConsumerEnable,omitnil,omitempty" name:"IsolateConsumerEnable"`

	// 消费者 Ack 超时时间，单位：秒，范围60-（3600*24
	AckTimeOut *int64 `json:"AckTimeOut,omitnil,omitempty" name:"AckTimeOut"`
}

type ModifyTopicRequest struct {
	*tchttp.BaseRequest
	
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 分区数，必须大于或者等于原分区数，若想维持原分区数请输入原数目，修改分区数仅对非全局顺序消息起效果，不允许超过32个分区。
	Partitions *uint64 `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 备注，128字符以内。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 未消费消息过期时间，单位：秒，取值范围：60秒~15天。
	MsgTTL *uint64 `json:"MsgTTL,omitnil,omitempty" name:"MsgTTL"`

	// 不传默认是原生策略，DefaultPolicy表示当订阅下达到最大未确认消息数 5000 时，服务端将不再向当前订阅下的所有消费者推送消息，DynamicPolicy表示动态调整订阅下的最大未确认消息数，具体配额是在 5000 和消费者数量*20之间取最大值。每个消费者默认最大 unack 消息数为 20，超过该限制时仅影响该消费者，不影响其他消费者。
	UnackPolicy *string `json:"UnackPolicy,omitnil,omitempty" name:"UnackPolicy"`

	// 是否开启异常消费者隔离
	IsolateConsumerEnable *bool `json:"IsolateConsumerEnable,omitnil,omitempty" name:"IsolateConsumerEnable"`

	// 消费者 Ack 超时时间，单位：秒，范围60-（3600*24
	AckTimeOut *int64 `json:"AckTimeOut,omitnil,omitempty" name:"AckTimeOut"`
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
	delete(f, "ClusterId")
	delete(f, "Remark")
	delete(f, "MsgTTL")
	delete(f, "UnackPolicy")
	delete(f, "IsolateConsumerEnable")
	delete(f, "AckTimeOut")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicResponseParams struct {
	// 分区数
	Partitions *uint64 `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 备注，128字符以内。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTopicResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTopicResponseParams `json:"Response"`
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

type MsgLog struct {
	// 消息ID。
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 生产者名称。
	ProducerName *string `json:"ProducerName,omitnil,omitempty" name:"ProducerName"`

	// 生产时间。
	ProduceTime *string `json:"ProduceTime,omitnil,omitempty" name:"ProduceTime"`

	// 生产客户端地址。
	ProducerAddr *string `json:"ProducerAddr,omitnil,omitempty" name:"ProducerAddr"`
}

type PartitionsTopic struct {
	// 最后一次间隔内发布消息的平均byte大小。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AverageMsgSize *string `json:"AverageMsgSize,omitnil,omitempty" name:"AverageMsgSize"`

	// 消费者数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerCount *string `json:"ConsumerCount,omitnil,omitempty" name:"ConsumerCount"`

	// 被记录下来的消息总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastConfirmedEntry *string `json:"LastConfirmedEntry,omitnil,omitempty" name:"LastConfirmedEntry"`

	// 最后一个ledger创建的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastLedgerCreatedTimestamp *string `json:"LastLedgerCreatedTimestamp,omitnil,omitempty" name:"LastLedgerCreatedTimestamp"`

	// 本地和复制的发布者每秒发布消息的速率。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateIn *string `json:"MsgRateIn,omitnil,omitempty" name:"MsgRateIn"`

	// 本地和复制的消费者每秒分发消息的数量之和。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateOut *string `json:"MsgRateOut,omitnil,omitempty" name:"MsgRateOut"`

	// 本地和复制的发布者每秒发布消息的byte。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputIn *string `json:"MsgThroughputIn,omitnil,omitempty" name:"MsgThroughputIn"`

	// 本地和复制的消费者每秒分发消息的byte。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputOut *string `json:"MsgThroughputOut,omitnil,omitempty" name:"MsgThroughputOut"`

	// 被记录下来的消息总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumberOfEntries *string `json:"NumberOfEntries,omitnil,omitempty" name:"NumberOfEntries"`

	// 子分区id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions *int64 `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 生产者数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProducerCount *string `json:"ProducerCount,omitnil,omitempty" name:"ProducerCount"`

	// 以byte计算的所有消息存储总量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSize *string `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`

	// topic类型描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicType *uint64 `json:"TopicType,omitnil,omitempty" name:"TopicType"`
}

type ProducerInfo struct {
	// 客户端ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 客户端IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// 客户端语言
	// JAVA((byte) 0),
	//     CPP((byte) 1),
	//     DOTNET((byte) 2),
	//     PYTHON((byte) 3),
	//     DELPHI((byte) 4),
	//     ERLANG((byte) 5),
	//     RUBY((byte) 6),
	//     OTHER((byte) 7),
	//     HTTP((byte) 8),
	//     GO((byte) 9),
	//     PHP((byte) 10),
	//     OMS((byte) 11);
	// 注意：此字段可能返回 null，表示取不到有效值。
	Language *string `json:"Language,omitnil,omitempty" name:"Language"`

	// 客户端版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 最后生产时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTimestamp *int64 `json:"LastUpdateTimestamp,omitnil,omitempty" name:"LastUpdateTimestamp"`
}

type ProducerLog struct {
	// 消息ID。
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 生产者名称。
	ProducerName *string `json:"ProducerName,omitnil,omitempty" name:"ProducerName"`

	// 消息生产时间。
	ProduceTime *string `json:"ProduceTime,omitnil,omitempty" name:"ProduceTime"`

	// 生产者客户端。
	ProducerAddr *string `json:"ProducerAddr,omitnil,omitempty" name:"ProducerAddr"`

	// 生产耗时（秒）。
	ProduceUseTime *uint64 `json:"ProduceUseTime,omitnil,omitempty" name:"ProduceUseTime"`

	// 状态。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type PrometheusEndpointInfo struct {
	// Prometheus开关的状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrometheusEndpointStatus *string `json:"PrometheusEndpointStatus,omitnil,omitempty" name:"PrometheusEndpointStatus"`

	// prometheus信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcPrometheusEndpoint []*string `json:"VpcPrometheusEndpoint,omitnil,omitempty" name:"VpcPrometheusEndpoint"`

	// 节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodePrometheusAddress []*string `json:"NodePrometheusAddress,omitnil,omitempty" name:"NodePrometheusAddress"`

	// vpc信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcEndpointInfo *VpcEndpointInfo `json:"VpcEndpointInfo,omitnil,omitempty" name:"VpcEndpointInfo"`
}

type PublicAccessRule struct {
	// ip网段信息
	IpRule *string `json:"IpRule,omitnil,omitempty" name:"IpRule"`

	// 允许或者拒绝
	Allow *bool `json:"Allow,omitnil,omitempty" name:"Allow"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

// Predefined struct for user
type PublishCmqMsgRequestParams struct {
	// 主题名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 消息内容，消息总大小需不大于1024K
	MsgContent *string `json:"MsgContent,omitnil,omitempty" name:"MsgContent"`

	// 消息标签，支持传递多标签或单路由，单个标签、路由长度不能超过64个字符。
	MsgTag []*string `json:"MsgTag,omitnil,omitempty" name:"MsgTag"`
}

type PublishCmqMsgRequest struct {
	*tchttp.BaseRequest
	
	// 主题名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 消息内容，消息总大小需不大于1024K
	MsgContent *string `json:"MsgContent,omitnil,omitempty" name:"MsgContent"`

	// 消息标签，支持传递多标签或单路由，单个标签、路由长度不能超过64个字符。
	MsgTag []*string `json:"MsgTag,omitnil,omitempty" name:"MsgTag"`
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

// Predefined struct for user
type PublishCmqMsgResponseParams struct {
	// true表示发送成功
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 消息id
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PublishCmqMsgResponse struct {
	*tchttp.BaseResponse
	Response *PublishCmqMsgResponseParams `json:"Response"`
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
	ProducerId *int64 `json:"ProducerId,omitnil,omitempty" name:"ProducerId"`

	// 生产者名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProducerName *string `json:"ProducerName,omitnil,omitempty" name:"ProducerName"`

	// 生产者地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 客户端版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientVersion *string `json:"ClientVersion,omitnil,omitempty" name:"ClientVersion"`

	// 消息生产速率（条/秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateIn *float64 `json:"MsgRateIn,omitnil,omitempty" name:"MsgRateIn"`

	// 消息生产吞吐速率（字节/秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputIn *float64 `json:"MsgThroughputIn,omitnil,omitempty" name:"MsgThroughputIn"`

	// 平均消息大小（字节）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AverageMsgSize *float64 `json:"AverageMsgSize,omitnil,omitempty" name:"AverageMsgSize"`

	// 连接时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectedSince *string `json:"ConnectedSince,omitnil,omitempty" name:"ConnectedSince"`

	// 生产者连接的主题分区号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partition *int64 `json:"Partition,omitnil,omitempty" name:"Partition"`
}

type PulsarNetworkAccessPointInfo struct {
	// vpc的id，支撑网和公网接入点，该字段为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网id，支撑网和公网接入点，该字段为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 接入地址
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 接入点类型：
	// 0：支撑网接入点 
	// 1：VPC接入点 
	// 2：公网接入点
	RouteType *uint64 `json:"RouteType,omitnil,omitempty" name:"RouteType"`

	// 0：本地域访问，由于并没有配置跨地域容灾，所该类型的接入点，无法进行异地切换、异地访问切回；
	// 1：本地域访问，由于配置了跨地域容灾，随时可以进行异地切换，该状态用于主集群的接入点
	// 2：跨地域访问，已经完成了异地切换，该状态用于源集群的接入点，该状态下的接入点不可删除
	// 3：跨地域访问，随时可以进行异地访问切回，该状态用于目标集群的接入点，该状态下的接入点不可删除
	// 4:跨地域访问，目标集群已经完成异地切回，等待删除状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationType *uint64 `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 接入点类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessPointsType *string `json:"AccessPointsType,omitnil,omitempty" name:"AccessPointsType"`

	// 带宽，目前只有公网会有这个值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 类
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityPolicy []*SecurityPolicy `json:"SecurityPolicy,omitnil,omitempty" name:"SecurityPolicy"`

	// 是否是标准的接入点 true是标准的 false不是标准的
	// 注意：此字段可能返回 null，表示取不到有效值。
	StandardAccessPoint *bool `json:"StandardAccessPoint,omitnil,omitempty" name:"StandardAccessPoint"`

	// 可用区信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 是否开启TLS加密
	Tls *bool `json:"Tls,omitnil,omitempty" name:"Tls"`

	// 接入点自定义域名
	CustomUrl *string `json:"CustomUrl,omitnil,omitempty" name:"CustomUrl"`
}

type PulsarProClusterInfo struct {
	// 集群Id。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 说明信息。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 集群状态，0:创建中，1:正常，2:隔离
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 集群版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 节点分布情况
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeDistribution []*InstanceNodeDistribution `json:"NodeDistribution,omitnil,omitempty" name:"NodeDistribution"`

	// 最大储存容量，单位：MB
	MaxStorage *uint64 `json:"MaxStorage,omitnil,omitempty" name:"MaxStorage"`

	// 是否可以修改路由
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanEditRoute *bool `json:"CanEditRoute,omitnil,omitempty" name:"CanEditRoute"`

	// 代表是专业版和小规格专业版的不同计费规格PULSAR.P1固定存储PULSAR.P2弹性存储
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingLabelVersion *string `json:"BillingLabelVersion,omitnil,omitempty" name:"BillingLabelVersion"`

	// 实例到期时间戳，毫秒级精度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 是否开启自动创建主题
	// true就是开启了，false是关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoCreateTopicStatus *bool `json:"AutoCreateTopicStatus,omitnil,omitempty" name:"AutoCreateTopicStatus"`

	// 自动创建主题的默认分区数，如果没开启就是0
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultPartitionNumber *int64 `json:"DefaultPartitionNumber,omitnil,omitempty" name:"DefaultPartitionNumber"`

	// 用户自定义的租户别名，如果没有，会复用专业集群 ID
	Tenant *string `json:"Tenant,omitnil,omitempty" name:"Tenant"`

	// 删除保护开关标识
	DeleteProtection *int64 `json:"DeleteProtection,omitnil,omitempty" name:"DeleteProtection"`
}

type PulsarProClusterSpecInfo struct {
	// 集群规格名称
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 峰值tps
	MaxTps *uint64 `json:"MaxTps,omitnil,omitempty" name:"MaxTps"`

	// 峰值带宽。单位：mbps
	MaxBandWidth *uint64 `json:"MaxBandWidth,omitnil,omitempty" name:"MaxBandWidth"`

	// 最大命名空间个数
	MaxNamespaces *uint64 `json:"MaxNamespaces,omitnil,omitempty" name:"MaxNamespaces"`

	// 最大主题分区数
	MaxTopics *uint64 `json:"MaxTopics,omitnil,omitempty" name:"MaxTopics"`

	// 规格外弹性TPS
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScalableTps *uint64 `json:"ScalableTps,omitnil,omitempty" name:"ScalableTps"`

	// 32或者128
	// 当前集群topic的最大分区数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxPartitions *uint64 `json:"MaxPartitions,omitnil,omitempty" name:"MaxPartitions"`

	// 商品最大延迟消息数量。0代表没有限制	
	MaxDelayedMessages *int64 `json:"MaxDelayedMessages,omitnil,omitempty" name:"MaxDelayedMessages"`
}

type PulsarProInstance struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例版本
	InstanceVersion *string `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`

	// 实例状态，0-创建中，1-正常，2-隔离中，3-已销毁，4 - 异常, 5 - 发货失败，6-变配中，7-变配失败
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例配置规格名称
	ConfigDisplay *string `json:"ConfigDisplay,omitnil,omitempty" name:"ConfigDisplay"`

	// 峰值TPS
	MaxTps *uint64 `json:"MaxTps,omitnil,omitempty" name:"MaxTps"`

	// 存储容量，GB为单位
	MaxStorage *uint64 `json:"MaxStorage,omitnil,omitempty" name:"MaxStorage"`

	// 实例到期时间，毫秒为单位
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 自动续费标记，0表示默认状态(用户未设置，即初始状态即手动续费)， 1表示自动续费，2表示明确不自动续费(用户设置)
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 0-后付费，1-预付费
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 实例配置ID
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 规格外弹性TPS
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScalableTps *uint64 `json:"ScalableTps,omitnil,omitempty" name:"ScalableTps"`

	// VPC的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 峰值带宽。单位：mbps
	MaxBandWidth *uint64 `json:"MaxBandWidth,omitnil,omitempty" name:"MaxBandWidth"`

	// 集群的标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 集群创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 代表是专业版和小规格专业版的不同计费规格PULSAR.P1固定存储PULSAR.P2弹性存储
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingLabelVersion *string `json:"BillingLabelVersion,omitnil,omitempty" name:"BillingLabelVersion"`

	// 自定义租户
	Tenant *string `json:"Tenant,omitnil,omitempty" name:"Tenant"`

	// 集群的证书列表
	CertificateList []*CertificateInfo `json:"CertificateList,omitnil,omitempty" name:"CertificateList"`
}

type QueueQuota struct {
	// 可创建最大Queue数
	MaxQueue *int64 `json:"MaxQueue,omitnil,omitempty" name:"MaxQueue"`

	// 已创建Queue数
	UsedQueue *int64 `json:"UsedQueue,omitnil,omitempty" name:"UsedQueue"`
}

type RabbitMQBindingListInfo struct {
	// 路由关系id
	BindingId *int64 `json:"BindingId,omitnil,omitempty" name:"BindingId"`

	// VhostName
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 源exchange名称
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标类型,queue或exchange
	DestinationType *string `json:"DestinationType,omitnil,omitempty" name:"DestinationType"`

	// 目标资源名称
	Destination *string `json:"Destination,omitnil,omitempty" name:"Destination"`

	// 绑定key
	RoutingKey *string `json:"RoutingKey,omitnil,omitempty" name:"RoutingKey"`

	// 源exchange类型
	SourceExchangeType *string `json:"SourceExchangeType,omitnil,omitempty" name:"SourceExchangeType"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 创建时间时间戳
	CreateTs *uint64 `json:"CreateTs,omitnil,omitempty" name:"CreateTs"`

	// 修改时间时间戳
	ModifyTs *uint64 `json:"ModifyTs,omitnil,omitempty" name:"ModifyTs"`
}

type RabbitMQClusterAccessInfo struct {
	// 集群公网接入地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicAccessEndpoint *string `json:"PublicAccessEndpoint,omitnil,omitempty" name:"PublicAccessEndpoint"`

	// 集群控制台访问地址
	WebConsoleEndpoint *string `json:"WebConsoleEndpoint,omitnil,omitempty" name:"WebConsoleEndpoint"`

	// 集群控制台登录用户名
	WebConsoleUsername *string `json:"WebConsoleUsername,omitnil,omitempty" name:"WebConsoleUsername"`

	// 集群控制台登录密码
	WebConsolePassword *string `json:"WebConsolePassword,omitnil,omitempty" name:"WebConsolePassword"`

	// 已废弃
	PublicAccessEndpointStatus *bool `json:"PublicAccessEndpointStatus,omitnil,omitempty" name:"PublicAccessEndpointStatus"`

	// 已废弃
	PublicControlConsoleSwitchStatus *bool `json:"PublicControlConsoleSwitchStatus,omitnil,omitempty" name:"PublicControlConsoleSwitchStatus"`

	// 已废弃
	VpcControlConsoleSwitchStatus *bool `json:"VpcControlConsoleSwitchStatus,omitnil,omitempty" name:"VpcControlConsoleSwitchStatus"`

	// Vpc管控台访问地址，示例值，http://1.1.1.1:15672
	VpcWebConsoleEndpoint *string `json:"VpcWebConsoleEndpoint,omitnil,omitempty" name:"VpcWebConsoleEndpoint"`

	// 公网管控台开关状态，示例值，OFF/ON/CREATING/DELETING
	PublicWebConsoleSwitchStatus *string `json:"PublicWebConsoleSwitchStatus,omitnil,omitempty" name:"PublicWebConsoleSwitchStatus"`

	// Vpc管控台开关状态，示例值，
	// OFF/ON/CREATING/DELETING
	VpcWebConsoleSwitchStatus *string `json:"VpcWebConsoleSwitchStatus,omitnil,omitempty" name:"VpcWebConsoleSwitchStatus"`

	// 公网管控台开关状态，示例值，OFF/ON/CREATING/DELETING
	PublicDataStreamStatus *string `json:"PublicDataStreamStatus,omitnil,omitempty" name:"PublicDataStreamStatus"`

	// Prometheus信息
	PrometheusEndpointInfo *PrometheusEndpointInfo `json:"PrometheusEndpointInfo,omitnil,omitempty" name:"PrometheusEndpointInfo"`

	// 公网域名接入点
	WebConsoleDomainEndpoint *string `json:"WebConsoleDomainEndpoint,omitnil,omitempty" name:"WebConsoleDomainEndpoint"`

	// 控制面所使用的VPC信息
	ControlPlaneEndpointInfo *VpcEndpointInfo `json:"ControlPlaneEndpointInfo,omitnil,omitempty" name:"ControlPlaneEndpointInfo"`

	// TLS加密的数据流公网接入点
	PublicTlsAccessEndpoint *string `json:"PublicTlsAccessEndpoint,omitnil,omitempty" name:"PublicTlsAccessEndpoint"`
}

type RabbitMQClusterInfo struct {
	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 地域信息
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 创建时间，毫秒为单位。unix 时间戳
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 集群说明信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// VPC及网络信息
	Vpcs []*VpcEndpointInfo `json:"Vpcs,omitnil,omitempty" name:"Vpcs"`

	// 可用区信息
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 虚拟主机数量
	VirtualHostNumber *int64 `json:"VirtualHostNumber,omitnil,omitempty" name:"VirtualHostNumber"`

	// 队列数量
	QueueNumber *int64 `json:"QueueNumber,omitnil,omitempty" name:"QueueNumber"`

	// 每秒生产消息数 单位：条/秒
	MessagePublishRate *float64 `json:"MessagePublishRate,omitnil,omitempty" name:"MessagePublishRate"`

	// 堆积消息数 单位：条
	MessageStackNumber *int64 `json:"MessageStackNumber,omitnil,omitempty" name:"MessageStackNumber"`

	// 实例到期时间，按量付费的资源该值为 0，毫秒为单位。unix 时间戳
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// Channel数量
	ChannelNumber *int64 `json:"ChannelNumber,omitnil,omitempty" name:"ChannelNumber"`

	// Connection数量
	ConnectionNumber *int64 `json:"ConnectionNumber,omitnil,omitempty" name:"ConnectionNumber"`

	// Consumer数量
	ConsumerNumber *int64 `json:"ConsumerNumber,omitnil,omitempty" name:"ConsumerNumber"`

	// Exchang数量
	ExchangeNumber *int64 `json:"ExchangeNumber,omitnil,omitempty" name:"ExchangeNumber"`

	// 集群异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceptionInformation *string `json:"ExceptionInformation,omitnil,omitempty" name:"ExceptionInformation"`

	// 实例状态，0表示创建中，1表示正常，2表示隔离中，3表示已销毁，4 - 异常, 5 - 发货失败
	ClusterStatus *int64 `json:"ClusterStatus,omitnil,omitempty" name:"ClusterStatus"`

	// 自动续费标记，0表示默认状态(用户未设置，即初始状态即手动续费)， 1表示自动续费，2表示明确不自动续费(用户设置)
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 是否开启镜像队列策略。1表示开启，0表示没开启。
	MirrorQueuePolicyFlag *int64 `json:"MirrorQueuePolicyFlag,omitnil,omitempty" name:"MirrorQueuePolicyFlag"`

	// 每秒消费消息数 单位：条/秒
	MessageConsumeRate *float64 `json:"MessageConsumeRate,omitnil,omitempty" name:"MessageConsumeRate"`

	// 集群版本信息
	ClusterVersion *string `json:"ClusterVersion,omitnil,omitempty" name:"ClusterVersion"`

	// 计费模式，0-后付费，1-预付费
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 实例类型，0 专享版、1 Serverless 版
	InstanceType *uint64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 开始隔离时间。unix 时间戳
	IsolatedTime *int64 `json:"IsolatedTime,omitnil,omitempty" name:"IsolatedTime"`

	// 是否为容器实例，默认 true
	Container *bool `json:"Container,omitnil,omitempty" name:"Container"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否已开启删除保护
	EnableDeletionProtection *bool `json:"EnableDeletionProtection,omitnil,omitempty" name:"EnableDeletionProtection"`
}

type RabbitMQClusterSpecInfo struct {
	// 集群规格名称
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 节点数量
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 峰值tps
	MaxTps *uint64 `json:"MaxTps,omitnil,omitempty" name:"MaxTps"`

	// 峰值带宽。单位：mbps
	MaxBandWidth *uint64 `json:"MaxBandWidth,omitnil,omitempty" name:"MaxBandWidth"`

	// 存储容量。单位：GB
	MaxStorage *uint64 `json:"MaxStorage,omitnil,omitempty" name:"MaxStorage"`

	// 公网带宽tps。单位：Mbps
	PublicNetworkTps *uint64 `json:"PublicNetworkTps,omitnil,omitempty" name:"PublicNetworkTps"`
}

type RabbitMQClusterWhiteListInfo struct {
	// 废弃
	WhiteList *string `json:"WhiteList,omitnil,omitempty" name:"WhiteList"`

	// 公网管控台白名单
	PublicControlConsoleWhiteList *string `json:"PublicControlConsoleWhiteList,omitnil,omitempty" name:"PublicControlConsoleWhiteList"`

	// 公网数据流白名单
	PublicDataStreamWhiteList *string `json:"PublicDataStreamWhiteList,omitnil,omitempty" name:"PublicDataStreamWhiteList"`

	// 公网管控台白名单状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicControlConsoleWhiteListStatus *string `json:"PublicControlConsoleWhiteListStatus,omitnil,omitempty" name:"PublicControlConsoleWhiteListStatus"`

	// 公网数据流白名单状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicDataStreamWhiteListStatus *string `json:"PublicDataStreamWhiteListStatus,omitnil,omitempty" name:"PublicDataStreamWhiteListStatus"`
}

type RabbitMQExchangeListInfo struct {
	// exchange 名称
	ExchangeName *string `json:"ExchangeName,omitnil,omitempty" name:"ExchangeName"`

	// 备注说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// exchange 类型, 支持 "fanout","direct","topic","headers"
	ExchangeType *string `json:"ExchangeType,omitnil,omitempty" name:"ExchangeType"`

	// 交换机所属 Virtual Host 名称
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// exchange 创建者, "system":"系统创建", "user":"用户创建"
	ExchangeCreator *string `json:"ExchangeCreator,omitnil,omitempty" name:"ExchangeCreator"`

	// exchange 创建时间
	CreateTimeStamp *string `json:"CreateTimeStamp,omitnil,omitempty" name:"CreateTimeStamp"`

	// exchange 修改时间
	ModTimeStamp *string `json:"ModTimeStamp,omitnil,omitempty" name:"ModTimeStamp"`

	// 输入消息速率
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageRateIn *float64 `json:"MessageRateIn,omitnil,omitempty" name:"MessageRateIn"`

	// 输出消息速率
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageRateOut *float64 `json:"MessageRateOut,omitnil,omitempty" name:"MessageRateOut"`

	// 是否为持久化交换机，true 为持久化，false 为非持久化
	Durable *bool `json:"Durable,omitnil,omitempty" name:"Durable"`

	// 是否为自动删除交换机，true 为自动删除，false 为非自动删除
	AutoDelete *bool `json:"AutoDelete,omitnil,omitempty" name:"AutoDelete"`

	// 是否为内部交换机，true 为内部交换机
	Internal *bool `json:"Internal,omitnil,omitempty" name:"Internal"`

	// 交换机所属实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 生效的策略名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 扩展参数 key-value 对象
	Arguments *string `json:"Arguments,omitnil,omitempty" name:"Arguments"`

	// 未调度的延时消息数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessagesDelayed *uint64 `json:"MessagesDelayed,omitnil,omitempty" name:"MessagesDelayed"`

	// 创建时间时间戳
	CreateTs *uint64 `json:"CreateTs,omitnil,omitempty" name:"CreateTs"`

	// 修改时间时间戳
	ModifyTs *uint64 `json:"ModifyTs,omitnil,omitempty" name:"ModifyTs"`
}

type RabbitMQPermission struct {
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，形如 admin。有效的 User 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，点击集群列表中的集群，进入集群详情，并在用户与权限页签中找到用户列表，从而找到用户名称。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// VirtualHost 名称，形如 testvhost。有效的 VirtualHost 名称可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询，在左侧导航栏点击 Vhost，并在 Vhost 列表中找到 Vhost 名称。
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 权限类型，declare相关操作，该用户可操作该vhost下的资源名称正则表达式
	ConfigRegexp *string `json:"ConfigRegexp,omitnil,omitempty" name:"ConfigRegexp"`

	// 权限类型，消息写入相关操作，该用户可操作该vhost下的资源名称正则表达式
	WriteRegexp *string `json:"WriteRegexp,omitnil,omitempty" name:"WriteRegexp"`

	// 权限类型，消息读取相关操作，该用户可操作该vhost下的资源名称正则表达式
	ReadRegexp *string `json:"ReadRegexp,omitnil,omitempty" name:"ReadRegexp"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 创建时间时间戳
	CreateTs *uint64 `json:"CreateTs,omitnil,omitempty" name:"CreateTs"`

	// 修改时间时间戳
	ModifyTs *uint64 `json:"ModifyTs,omitnil,omitempty" name:"ModifyTs"`
}

type RabbitMQPrivateNode struct {
	// 节点名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 节点状态，running 运行中，down 异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeStatus *string `json:"NodeStatus,omitnil,omitempty" name:"NodeStatus"`

	// CPU使用率
	// 注意：此字段可能返回 null，表示取不到有效值。
	CPUUsage *string `json:"CPUUsage,omitnil,omitempty" name:"CPUUsage"`

	// 内存使用情况，单位MB
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *uint64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 磁盘使用率
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskUsage *string `json:"DiskUsage,omitnil,omitempty" name:"DiskUsage"`

	// Rabbitmq的Erlang进程数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessNumber *uint64 `json:"ProcessNumber,omitnil,omitempty" name:"ProcessNumber"`
}

type RabbitMQQueueListConsumerDetailInfo struct {
	// 消费者数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumersNumber *int64 `json:"ConsumersNumber,omitnil,omitempty" name:"ConsumersNumber"`
}

type RabbitMQQueueListInfo struct {
	// 队列名
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 备注说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 消费者信息
	ConsumerDetail *RabbitMQQueueListConsumerDetailInfo `json:"ConsumerDetail,omitnil,omitempty" name:"ConsumerDetail"`

	// 队列类型，取值 "classic"，"quorum"
	QueueType *string `json:"QueueType,omitnil,omitempty" name:"QueueType"`

	// 消息堆积数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageHeapCount *int64 `json:"MessageHeapCount,omitnil,omitempty" name:"MessageHeapCount"`

	// 消息生产速率，每秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageRateIn *float64 `json:"MessageRateIn,omitnil,omitempty" name:"MessageRateIn"`

	// 消息消费速率，每秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageRateOut *float64 `json:"MessageRateOut,omitnil,omitempty" name:"MessageRateOut"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 队列是否持久化，true 为持久化，false 为非持久化
	Durable *bool `json:"Durable,omitnil,omitempty" name:"Durable"`

	// 队列是否为自动删除队列，true 为自动删除，false 为非自动删除
	AutoDelete *bool `json:"AutoDelete,omitnil,omitempty" name:"AutoDelete"`

	// 队列所属实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 队列所属虚拟主机名称
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 队列所在主节点名称
	Node *string `json:"Node,omitnil,omitempty" name:"Node"`

	// 生效的策略名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 扩展参数 key-value 对象
	Arguments *string `json:"Arguments,omitnil,omitempty" name:"Arguments"`

	// 是否独占队列
	Exclusive *bool `json:"Exclusive,omitnil,omitempty" name:"Exclusive"`

	// 创建时间时间戳
	CreateTs *uint64 `json:"CreateTs,omitnil,omitempty" name:"CreateTs"`

	// 修改时间时间戳
	ModifyTs *uint64 `json:"ModifyTs,omitnil,omitempty" name:"ModifyTs"`
}

type RabbitMQUser struct {
	// 实例 ID，形如 amqp-xxxxxxxx。有效的 InstanceId 可通过登录 [TDMQ RabbitMQ 控制台](https://console.cloud.tencent.com/trabbitmq/cluster?rid=1)查询。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，登录时使用
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 密码，登录时使用
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 用户描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户标签，用于决定改用户访问RabbitMQ Management的权限范围
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 用户创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 用户最后修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 用户类型，System：系统创建，User：用户创建
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 单个用户最大可用连接数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxConnections *int64 `json:"MaxConnections,omitnil,omitempty" name:"MaxConnections"`

	// 单个用户最大可用通道数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxChannels *int64 `json:"MaxChannels,omitnil,omitempty" name:"MaxChannels"`

	// 创建时间时间戳
	CreateTs *uint64 `json:"CreateTs,omitnil,omitempty" name:"CreateTs"`

	// 修改时间时间戳
	ModifyTs *uint64 `json:"ModifyTs,omitnil,omitempty" name:"ModifyTs"`
}

type RabbitMQUserQuota struct {
	// 最大可创建用户数
	MaxUser *int64 `json:"MaxUser,omitnil,omitempty" name:"MaxUser"`

	// 已使用用户数
	UsedUser *int64 `json:"UsedUser,omitnil,omitempty" name:"UsedUser"`
}

type RabbitMQVipInstance struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例版本
	InstanceVersion *string `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`

	// 实例状态，0表示创建中，1表示正常，2表示隔离中，3表示已销毁，4 - 异常, 5 - 发货失败
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 节点数量
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 实例配置规格名称
	ConfigDisplay *string `json:"ConfigDisplay,omitnil,omitempty" name:"ConfigDisplay"`

	// 峰值TPS
	MaxTps *uint64 `json:"MaxTps,omitnil,omitempty" name:"MaxTps"`

	// 峰值带宽，Mbps为单位
	MaxBandWidth *uint64 `json:"MaxBandWidth,omitnil,omitempty" name:"MaxBandWidth"`

	// 存储容量，GB为单位
	MaxStorage *uint64 `json:"MaxStorage,omitnil,omitempty" name:"MaxStorage"`

	// 实例到期时间，按量付费的资源该值为 0，毫秒为单位。unix 时间戳
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 自动续费标记，0表示默认状态(用户未设置，即初始状态即手动续费)， 1表示自动续费，2表示明确不自动续费(用户设置)
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 1 表示预付费，0 表示后付费
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 集群的节点规格，需要输入对应的规格标识：
	// 2C8G：rabbit-vip-basic-2c8g
	// 4C16G：rabbit-vip-basic-4c16g
	// 8C32G：rabbit-vip-basic-8c32g
	// 16C32G：rabbit-vip-basic-4
	// 16C64G：rabbit-vip-basic-16c64g
	// 2C4G：rabbit-vip-basic-5
	// 4C8G：rabbit-vip-basic-1
	// 8C16G（已售罄）：rabbit-vip-basic-2
	// 不传默认为4C8G：rabbit-vip-basic-1
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 集群异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceptionInformation *string `json:"ExceptionInformation,omitnil,omitempty" name:"ExceptionInformation"`

	// 实例状态，0表示创建中，1表示正常，2表示隔离中，3表示已销毁，4 - 异常, 5 - 发货失败
	// 为了和计费区分开，额外开启一个状态位，用于显示。
	ClusterStatus *int64 `json:"ClusterStatus,omitnil,omitempty" name:"ClusterStatus"`

	// 公网接入点
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicAccessEndpoint *string `json:"PublicAccessEndpoint,omitnil,omitempty" name:"PublicAccessEndpoint"`

	// VPC 接入点列表
	Vpcs []*VpcEndpointInfo `json:"Vpcs,omitnil,omitempty" name:"Vpcs"`

	// 创建时间，毫秒为单位。unix 时间戳
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例类型，0 托管版、1 Serverless 版
	InstanceType *uint64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 隔离时间，毫秒为单位。unix 时间戳
	IsolatedTime *uint64 `json:"IsolatedTime,omitnil,omitempty" name:"IsolatedTime"`

	// 是否已开启删除保护
	EnableDeletionProtection *bool `json:"EnableDeletionProtection,omitnil,omitempty" name:"EnableDeletionProtection"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type RabbitMQVirtualHostInfo struct {
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost名
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// vhost描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// vhost标签
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// vhost概览统计信息
	VirtualHostStatistics *RabbitMQVirtualHostStatistics `json:"VirtualHostStatistics,omitnil,omitempty" name:"VirtualHostStatistics"`

	// 消息轨迹开关,true打开,false关闭
	TraceFlag *bool `json:"TraceFlag,omitnil,omitempty" name:"TraceFlag"`

	// vhost状态，与原生控制台对应，有running、partial、stopped、unknown
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 消息堆积数
	MessageHeapCount *int64 `json:"MessageHeapCount,omitnil,omitempty" name:"MessageHeapCount"`

	// 输入消息速率
	MessageRateIn *float64 `json:"MessageRateIn,omitnil,omitempty" name:"MessageRateIn"`

	// 输出消息速率
	MessageRateOut *float64 `json:"MessageRateOut,omitnil,omitempty" name:"MessageRateOut"`

	// 是否存在镜像队列策略，true 为存在，false 为不存
	MirrorQueuePolicyFlag *bool `json:"MirrorQueuePolicyFlag,omitnil,omitempty" name:"MirrorQueuePolicyFlag"`

	// 创建时间时间戳
	CreateTs *uint64 `json:"CreateTs,omitnil,omitempty" name:"CreateTs"`

	// 修改时间时间戳
	ModifyTs *uint64 `json:"ModifyTs,omitnil,omitempty" name:"ModifyTs"`
}

type RabbitMQVirtualHostStatistics struct {
	// 当前vhost的queue数量
	CurrentQueues *int64 `json:"CurrentQueues,omitnil,omitempty" name:"CurrentQueues"`

	// 当前vhost的exchange数量
	CurrentExchanges *int64 `json:"CurrentExchanges,omitnil,omitempty" name:"CurrentExchanges"`

	// 当前vhost的连接数量
	CurrentConnections *int64 `json:"CurrentConnections,omitnil,omitempty" name:"CurrentConnections"`

	// 当前vhost的channel数量
	CurrentChannels *int64 `json:"CurrentChannels,omitnil,omitempty" name:"CurrentChannels"`

	// 当前vhost的用户数量
	CurrentUsers *int64 `json:"CurrentUsers,omitnil,omitempty" name:"CurrentUsers"`
}

// Predefined struct for user
type ReceiveMessageRequestParams struct {
	// 接收消息的topic的名字, 这里尽量需要使用topic的全路径，如果不指定，即：tenant/namespace/topic。默认使用的是：public/default
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 订阅者的名字
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`

	// 默认值为1000，consumer接收的消息会首先存储到receiverQueueSize这个队列中，用作调优接收消息的速率
	ReceiverQueueSize *int64 `json:"ReceiverQueueSize,omitnil,omitempty" name:"ReceiverQueueSize"`

	// 默认值为：Earliest。用作判定consumer初始接收消息的位置，可选参数为：Earliest, Latest
	SubInitialPosition *string `json:"SubInitialPosition,omitnil,omitempty" name:"SubInitialPosition"`

	// 用于设置BatchReceivePolicy，指在一次batch中最多接收多少条消息，默认是 0。即不开启BatchReceivePolicy
	MaxNumMessages *int64 `json:"MaxNumMessages,omitnil,omitempty" name:"MaxNumMessages"`

	// 用于设置BatchReceivePolicy，指在一次batch中最多接收的消息体有多大，单位是 bytes。默认是 0，即不开启BatchReceivePolicy
	MaxNumBytes *int64 `json:"MaxNumBytes,omitnil,omitempty" name:"MaxNumBytes"`

	// 用于设置BatchReceivePolicy，指在一次batch消息的接收z中最多等待的超时时间，单位是毫秒。默认是 0，即不开启BatchReceivePolicy
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`
}

type ReceiveMessageRequest struct {
	*tchttp.BaseRequest
	
	// 接收消息的topic的名字, 这里尽量需要使用topic的全路径，如果不指定，即：tenant/namespace/topic。默认使用的是：public/default
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 订阅者的名字
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`

	// 默认值为1000，consumer接收的消息会首先存储到receiverQueueSize这个队列中，用作调优接收消息的速率
	ReceiverQueueSize *int64 `json:"ReceiverQueueSize,omitnil,omitempty" name:"ReceiverQueueSize"`

	// 默认值为：Earliest。用作判定consumer初始接收消息的位置，可选参数为：Earliest, Latest
	SubInitialPosition *string `json:"SubInitialPosition,omitnil,omitempty" name:"SubInitialPosition"`

	// 用于设置BatchReceivePolicy，指在一次batch中最多接收多少条消息，默认是 0。即不开启BatchReceivePolicy
	MaxNumMessages *int64 `json:"MaxNumMessages,omitnil,omitempty" name:"MaxNumMessages"`

	// 用于设置BatchReceivePolicy，指在一次batch中最多接收的消息体有多大，单位是 bytes。默认是 0，即不开启BatchReceivePolicy
	MaxNumBytes *int64 `json:"MaxNumBytes,omitnil,omitempty" name:"MaxNumBytes"`

	// 用于设置BatchReceivePolicy，指在一次batch消息的接收z中最多等待的超时时间，单位是毫秒。默认是 0，即不开启BatchReceivePolicy
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`
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
	delete(f, "MaxNumMessages")
	delete(f, "MaxNumBytes")
	delete(f, "Timeout")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReceiveMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReceiveMessageResponseParams struct {
	// 用作标识消息的唯一主键
	MessageID *string `json:"MessageID,omitnil,omitempty" name:"MessageID"`

	// 接收消息的内容
	MessagePayload *string `json:"MessagePayload,omitnil,omitempty" name:"MessagePayload"`

	// 提供给 Ack 接口，用来Ack哪一个topic中的消息
	AckTopic *string `json:"AckTopic,omitnil,omitempty" name:"AckTopic"`

	// 返回的错误信息，如果为空，说明没有错误
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 返回订阅者的名字，用来创建 ack consumer时使用
	SubName *string `json:"SubName,omitnil,omitempty" name:"SubName"`

	// BatchReceivePolicy 一次性返回的多条消息的 MessageID，用 ‘###’ 来区分不同的 MessageID
	MessageIDList *string `json:"MessageIDList,omitnil,omitempty" name:"MessageIDList"`

	// BatchReceivePolicy 一次性返回的多条消息的消息内容，用 ‘###’ 来区分不同的消息内容
	MessagesPayload *string `json:"MessagesPayload,omitnil,omitempty" name:"MessagesPayload"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReceiveMessageResponse struct {
	*tchttp.BaseResponse
	Response *ReceiveMessageResponseParams `json:"Response"`
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

// Predefined struct for user
type ResetMsgSubOffsetByTimestampRequestParams struct {
	// 命名空间名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 订阅者名称。
	Subscription *string `json:"Subscription,omitnil,omitempty" name:"Subscription"`

	// 时间戳，精确到毫秒。
	ToTimestamp *uint64 `json:"ToTimestamp,omitnil,omitempty" name:"ToTimestamp"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type ResetMsgSubOffsetByTimestampRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 订阅者名称。
	Subscription *string `json:"Subscription,omitnil,omitempty" name:"Subscription"`

	// 时间戳，精确到毫秒。
	ToTimestamp *uint64 `json:"ToTimestamp,omitnil,omitempty" name:"ToTimestamp"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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

// Predefined struct for user
type ResetMsgSubOffsetByTimestampResponseParams struct {
	// 结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetMsgSubOffsetByTimestampResponse struct {
	*tchttp.BaseResponse
	Response *ResetMsgSubOffsetByTimestampResponseParams `json:"Response"`
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

// Predefined struct for user
type ResetRocketMQConsumerOffSetRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 消费组名称
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 重置方式，0表示从最新位点开始，1表示从指定时间点开始
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 重置指定的时间戳，仅在 Type 为1是生效，以毫秒为单位
	ResetTimestamp *uint64 `json:"ResetTimestamp,omitnil,omitempty" name:"ResetTimestamp"`

	// 重置的是否是retry topic
	RetryFlag *bool `json:"RetryFlag,omitnil,omitempty" name:"RetryFlag"`
}

type ResetRocketMQConsumerOffSetRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 消费组名称
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 重置方式，0表示从最新位点开始，1表示从指定时间点开始
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 重置指定的时间戳，仅在 Type 为1是生效，以毫秒为单位
	ResetTimestamp *uint64 `json:"ResetTimestamp,omitnil,omitempty" name:"ResetTimestamp"`

	// 重置的是否是retry topic
	RetryFlag *bool `json:"RetryFlag,omitnil,omitempty" name:"RetryFlag"`
}

func (r *ResetRocketMQConsumerOffSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetRocketMQConsumerOffSetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "GroupId")
	delete(f, "Type")
	delete(f, "Topic")
	delete(f, "ResetTimestamp")
	delete(f, "RetryFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetRocketMQConsumerOffSetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetRocketMQConsumerOffSetResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetRocketMQConsumerOffSetResponse struct {
	*tchttp.BaseResponse
	Response *ResetRocketMQConsumerOffSetResponseParams `json:"Response"`
}

func (r *ResetRocketMQConsumerOffSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetRocketMQConsumerOffSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetentionPolicy struct {
	// 消息保留时长
	TimeInMinutes *int64 `json:"TimeInMinutes,omitnil,omitempty" name:"TimeInMinutes"`

	// 消息保留大小
	SizeInMB *int64 `json:"SizeInMB,omitnil,omitempty" name:"SizeInMB"`
}

// Predefined struct for user
type RetryRocketMQDlqMessageRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// group名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 死信消息ID
	MessageIds []*string `json:"MessageIds,omitnil,omitempty" name:"MessageIds"`
}

type RetryRocketMQDlqMessageRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间名称
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// group名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 死信消息ID
	MessageIds []*string `json:"MessageIds,omitnil,omitempty" name:"MessageIds"`
}

func (r *RetryRocketMQDlqMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryRocketMQDlqMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "GroupName")
	delete(f, "MessageIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryRocketMQDlqMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryRocketMQDlqMessageResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RetryRocketMQDlqMessageResponse struct {
	*tchttp.BaseResponse
	Response *RetryRocketMQDlqMessageResponseParams `json:"Response"`
}

func (r *RetryRocketMQDlqMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryRocketMQDlqMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RewindCmqQueueRequestParams struct {
	// 队列名字，在单个地域同一账号下唯一。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 设定该时间，则（Batch）receiveMessage接口，会按照生产消息的先后顺序消费该时间戳以后的消息。
	StartConsumeTime *uint64 `json:"StartConsumeTime,omitnil,omitempty" name:"StartConsumeTime"`
}

type RewindCmqQueueRequest struct {
	*tchttp.BaseRequest
	
	// 队列名字，在单个地域同一账号下唯一。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 设定该时间，则（Batch）receiveMessage接口，会按照生产消息的先后顺序消费该时间戳以后的消息。
	StartConsumeTime *uint64 `json:"StartConsumeTime,omitnil,omitempty" name:"StartConsumeTime"`
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

// Predefined struct for user
type RewindCmqQueueResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RewindCmqQueueResponse struct {
	*tchttp.BaseResponse
	Response *RewindCmqQueueResponseParams `json:"Response"`
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
	// 单命名空间TPS上限
	//
	// Deprecated: MaxTpsPerNamespace is deprecated.
	MaxTpsPerNamespace *uint64 `json:"MaxTpsPerNamespace,omitnil,omitempty" name:"MaxTpsPerNamespace"`

	// 最大命名空间数量
	MaxNamespaceNum *uint64 `json:"MaxNamespaceNum,omitnil,omitempty" name:"MaxNamespaceNum"`

	// 已使用命名空间数量
	UsedNamespaceNum *uint64 `json:"UsedNamespaceNum,omitnil,omitempty" name:"UsedNamespaceNum"`

	// 最大Topic数量
	MaxTopicNum *uint64 `json:"MaxTopicNum,omitnil,omitempty" name:"MaxTopicNum"`

	// 已使用Topic数量
	UsedTopicNum *uint64 `json:"UsedTopicNum,omitnil,omitempty" name:"UsedTopicNum"`

	// 最大Group数量
	MaxGroupNum *uint64 `json:"MaxGroupNum,omitnil,omitempty" name:"MaxGroupNum"`

	// 已使用Group数量
	UsedGroupNum *uint64 `json:"UsedGroupNum,omitnil,omitempty" name:"UsedGroupNum"`

	// 消息最大保留时间，以毫秒为单位
	MaxRetentionTime *uint64 `json:"MaxRetentionTime,omitnil,omitempty" name:"MaxRetentionTime"`

	// 消息最长延时，以毫秒为单位
	MaxLatencyTime *uint64 `json:"MaxLatencyTime,omitnil,omitempty" name:"MaxLatencyTime"`

	// 单个主题最大队列数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxQueuesPerTopic *uint64 `json:"MaxQueuesPerTopic,omitnil,omitempty" name:"MaxQueuesPerTopic"`

	// topic分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicDistribution []*RocketMQTopicDistribution `json:"TopicDistribution,omitnil,omitempty" name:"TopicDistribution"`

	// 最大角色数量
	MaxRoleNum *int64 `json:"MaxRoleNum,omitnil,omitempty" name:"MaxRoleNum"`

	// TPS限额
	MaxTpsLimit *int64 `json:"MaxTpsLimit,omitnil,omitempty" name:"MaxTpsLimit"`
}

type RocketMQClusterDetail struct {
	// 集群基本信息
	Info *RocketMQClusterInfo `json:"Info,omitnil,omitempty" name:"Info"`

	// 集群配置信息
	Config *RocketMQClusterConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 集群状态，0:创建中，1:正常，2:销毁中，3:已删除，4: 隔离中，5:创建失败，6: 删除失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type RocketMQClusterInfo struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 地域信息
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 创建时间，毫秒为单位
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 集群说明信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 公网接入地址
	PublicEndPoint *string `json:"PublicEndPoint,omitnil,omitempty" name:"PublicEndPoint"`

	// VPC接入地址
	VpcEndPoint *string `json:"VpcEndPoint,omitnil,omitempty" name:"VpcEndPoint"`

	// 是否支持命名空间接入点
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportNamespaceEndpoint *bool `json:"SupportNamespaceEndpoint,omitnil,omitempty" name:"SupportNamespaceEndpoint"`

	// VPC信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vpcs []*VpcConfig `json:"Vpcs,omitnil,omitempty" name:"Vpcs"`

	// 是否为专享实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsVip *bool `json:"IsVip,omitnil,omitempty" name:"IsVip"`

	// Rocketmq集群标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	RocketMQFlag *bool `json:"RocketMQFlag,omitnil,omitempty" name:"RocketMQFlag"`

	// 计费状态，1表示正常，2表示已停服，3表示已销毁
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 欠费停服时间，毫秒为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolateTime *int64 `json:"IsolateTime,omitnil,omitempty" name:"IsolateTime"`

	// HTTP协议公网接入地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpPublicEndpoint *string `json:"HttpPublicEndpoint,omitnil,omitempty" name:"HttpPublicEndpoint"`

	// HTTP协议VPC接入地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpVpcEndpoint *string `json:"HttpVpcEndpoint,omitnil,omitempty" name:"HttpVpcEndpoint"`

	// TCP内部接入地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternalEndpoint *string `json:"InternalEndpoint,omitnil,omitempty" name:"InternalEndpoint"`

	// HTTP协议内部接入地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpInternalEndpoint *string `json:"HttpInternalEndpoint,omitnil,omitempty" name:"HttpInternalEndpoint"`

	// 是否开启ACL鉴权，专享实例支持关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	AclEnabled *bool `json:"AclEnabled,omitnil,omitempty" name:"AclEnabled"`

	// 公网CLB实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicClbId *string `json:"PublicClbId,omitnil,omitempty" name:"PublicClbId"`

	// vip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 所属VPC
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 是否支持迁移
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportMigration *bool `json:"SupportMigration,omitnil,omitempty" name:"SupportMigration"`

	// 实例状态，0表示创建中，1表示正常，2表示隔离中，3表示已销毁，4 - 异常, 5 - 发货失败，6 - 变配中，7 - 变配失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatus *int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 集群所属可用区，表明集群归属的可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 集群节点所在的可用区，若该集群为跨可用区集群，则包含该集群节点所在的多个可用区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 是否已冻结
	IsFrozen *bool `json:"IsFrozen,omitnil,omitempty" name:"IsFrozen"`

	// 是否开启自动创建主题
	AutoCreateTopicEnabled *bool `json:"AutoCreateTopicEnabled,omitnil,omitempty" name:"AutoCreateTopicEnabled"`

	// 是否开启集群Admin能力
	AdminFeatureEnabled *bool `json:"AdminFeatureEnabled,omitnil,omitempty" name:"AdminFeatureEnabled"`

	// Admin AK
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminAccessKey *string `json:"AdminAccessKey,omitnil,omitempty" name:"AdminAccessKey"`

	// Admin SK
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminSecretKey *string `json:"AdminSecretKey,omitnil,omitempty" name:"AdminSecretKey"`

	// 是否开启删除保护
	EnableDeletionProtection *bool `json:"EnableDeletionProtection,omitnil,omitempty" name:"EnableDeletionProtection"`
}

type RocketMQClusterRecentStats struct {
	// Topic数量
	TopicNum *uint64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// 消息生产数
	ProducedMsgNum *uint64 `json:"ProducedMsgNum,omitnil,omitempty" name:"ProducedMsgNum"`

	// 消息消费数
	ConsumedMsgNum *uint64 `json:"ConsumedMsgNum,omitnil,omitempty" name:"ConsumedMsgNum"`

	// 消息堆积数
	AccumulativeMsgNum *uint64 `json:"AccumulativeMsgNum,omitnil,omitempty" name:"AccumulativeMsgNum"`
}

type RocketMQConsumerConnection struct {
	// 消费者实例ID
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 消费者实例的地址和端口
	ClientAddr *string `json:"ClientAddr,omitnil,omitempty" name:"ClientAddr"`

	// 消费者应用的语言版本
	Language *string `json:"Language,omitnil,omitempty" name:"Language"`

	// 消息堆积量
	Accumulative *int64 `json:"Accumulative,omitnil,omitempty" name:"Accumulative"`

	// 消费端版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`
}

type RocketMQConsumerTopic struct {
	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 主题类型，Normal表示普通，GlobalOrder表示全局顺序，PartitionedOrder表示局部顺序，Transaction表示事务，Retry表示重试，DeadLetter表示死信
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分区数
	PartitionNum *uint64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// 消息堆积数
	Accumulative *int64 `json:"Accumulative,omitnil,omitempty" name:"Accumulative"`

	// 最后消费时间，以毫秒为单位
	LastConsumptionTime *uint64 `json:"LastConsumptionTime,omitnil,omitempty" name:"LastConsumptionTime"`

	// 订阅规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubRule *string `json:"SubRule,omitnil,omitempty" name:"SubRule"`
}

type RocketMQDataPoint struct {
	// 监控值数组，该数组和Timestamps一一对应
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamps []*int64 `json:"Timestamps,omitnil,omitempty" name:"Timestamps"`

	// 监控数据点位置，比如一天按分钟划分有1440个点，每个点的序号是0 - 1439之间的一个数，当某个序号不在该数组中，说明掉点了
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*float64 `json:"Values,omitnil,omitempty" name:"Values"`
}

type RocketMQGroup struct {
	// 消费组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 在线消费者数量
	ConsumerNum *uint64 `json:"ConsumerNum,omitnil,omitempty" name:"ConsumerNum"`

	// 消费TPS
	//
	// Deprecated: TPS is deprecated.
	TPS *uint64 `json:"TPS,omitnil,omitempty" name:"TPS"`

	// 总堆积数量
	//
	// Deprecated: TotalAccumulative is deprecated.
	TotalAccumulative *int64 `json:"TotalAccumulative,omitnil,omitempty" name:"TotalAccumulative"`

	// 0表示集群消费模式，1表示广播消费模式，-1表示未知
	ConsumptionMode *int64 `json:"ConsumptionMode,omitnil,omitempty" name:"ConsumptionMode"`

	// 是否允许消费
	ReadEnabled *bool `json:"ReadEnabled,omitnil,omitempty" name:"ReadEnabled"`

	// 重试队列分区数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryPartitionNum *uint64 `json:"RetryPartitionNum,omitnil,omitempty" name:"RetryPartitionNum"`

	// 创建时间，以毫秒为单位
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间，以毫秒为单位
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 客户端协议
	ClientProtocol *string `json:"ClientProtocol,omitnil,omitempty" name:"ClientProtocol"`

	// 说明信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 消费者类型，枚举值ACTIVELY, PASSIVELY
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerType *string `json:"ConsumerType,omitnil,omitempty" name:"ConsumerType"`

	// 是否开启广播消费
	BroadcastEnabled *bool `json:"BroadcastEnabled,omitnil,omitempty" name:"BroadcastEnabled"`

	// Group类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 重试次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryMaxTimes *uint64 `json:"RetryMaxTimes,omitnil,omitempty" name:"RetryMaxTimes"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 订阅的主题个数
	SubscribeTopicNum *int64 `json:"SubscribeTopicNum,omitnil,omitempty" name:"SubscribeTopicNum"`
}

type RocketMQGroupConfig struct {
	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 消费组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 是否开启广播消费
	ConsumeBroadcastEnable *bool `json:"ConsumeBroadcastEnable,omitnil,omitempty" name:"ConsumeBroadcastEnable"`

	// 是否开启消费
	ConsumeEnable *bool `json:"ConsumeEnable,omitnil,omitempty" name:"ConsumeEnable"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 协议类型，支持以下枚举值
	// TCP;
	// HTTP;
	ConsumerGroupType *string `json:"ConsumerGroupType,omitnil,omitempty" name:"ConsumerGroupType"`
}

type RocketMQGroupConfigOutput struct {
	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 消费组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 是否已导入
	// 注意：此字段可能返回 null，表示取不到有效值。
	Imported *bool `json:"Imported,omitnil,omitempty" name:"Imported"`

	// remark
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type RocketMQInstanceConfig struct {
	// 单命名空间TPS上线
	MaxTpsPerNamespace *uint64 `json:"MaxTpsPerNamespace,omitnil,omitempty" name:"MaxTpsPerNamespace"`

	// 最大命名空间数量
	MaxNamespaceNum *uint64 `json:"MaxNamespaceNum,omitnil,omitempty" name:"MaxNamespaceNum"`

	// 已使用命名空间数量
	UsedNamespaceNum *uint64 `json:"UsedNamespaceNum,omitnil,omitempty" name:"UsedNamespaceNum"`

	// 最大Topic数量
	MaxTopicNum *uint64 `json:"MaxTopicNum,omitnil,omitempty" name:"MaxTopicNum"`

	// 已使用Topic数量
	UsedTopicNum *uint64 `json:"UsedTopicNum,omitnil,omitempty" name:"UsedTopicNum"`

	// 最大Group数量
	MaxGroupNum *uint64 `json:"MaxGroupNum,omitnil,omitempty" name:"MaxGroupNum"`

	// 已使用Group数量
	UsedGroupNum *uint64 `json:"UsedGroupNum,omitnil,omitempty" name:"UsedGroupNum"`

	// 集群类型
	ConfigDisplay *string `json:"ConfigDisplay,omitnil,omitempty" name:"ConfigDisplay"`

	// 集群节点数
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 节点分布情况
	NodeDistribution []*InstanceNodeDistribution `json:"NodeDistribution,omitnil,omitempty" name:"NodeDistribution"`

	// topic分布情况
	TopicDistribution []*RocketMQTopicDistribution `json:"TopicDistribution,omitnil,omitempty" name:"TopicDistribution"`

	// 每个主题最大队列数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxQueuesPerTopic *uint64 `json:"MaxQueuesPerTopic,omitnil,omitempty" name:"MaxQueuesPerTopic"`

	// 最大可设置消息保留时间，小时为单位	
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRetention *int64 `json:"MaxRetention,omitnil,omitempty" name:"MaxRetention"`

	// 最小可设置消息保留时间，小时为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinRetention *int64 `json:"MinRetention,omitnil,omitempty" name:"MinRetention"`

	// 实例消息保留时间，小时为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Retention *int64 `json:"Retention,omitnil,omitempty" name:"Retention"`

	// Topic个数最小配额，即免费额度，默认为集群规格单节点最小配额*节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicNumLowerLimit *int64 `json:"TopicNumLowerLimit,omitnil,omitempty" name:"TopicNumLowerLimit"`

	// Topic个数最大配额，默认为集群规格单节点最大配额*节点个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicNumUpperLimit *int64 `json:"TopicNumUpperLimit,omitnil,omitempty" name:"TopicNumUpperLimit"`
}

type RocketMQMessageTrack struct {
	// 消费者组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 消费状态,
	// CONSUMED: 已消费
	// CONSUMED_BUT_FILTERED: 已过滤
	// NOT_CONSUME: 未消费
	// ENTER_RETRY: 进入重试队列
	// ENTER_DLQ: 进入死信队列
	// UNKNOWN: 查询不到消费状态
	ConsumeStatus *string `json:"ConsumeStatus,omitnil,omitempty" name:"ConsumeStatus"`

	// 消息track类型
	TrackType *string `json:"TrackType,omitnil,omitempty" name:"TrackType"`

	// 异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceptionDesc *string `json:"ExceptionDesc,omitnil,omitempty" name:"ExceptionDesc"`
}

type RocketMQMigrationTopicDistribution struct {
	// 迁移主题阶段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Stage *string `json:"Stage,omitnil,omitempty" name:"Stage"`

	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type RocketMQMsgLog struct {
	// 消息id
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 消息tag
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgTag *string `json:"MsgTag,omitnil,omitempty" name:"MsgTag"`

	// 消息key
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgKey *string `json:"MsgKey,omitnil,omitempty" name:"MsgKey"`

	// 客户端地址
	ProducerAddr *string `json:"ProducerAddr,omitnil,omitempty" name:"ProducerAddr"`

	// 消息发送时间
	ProduceTime *string `json:"ProduceTime,omitnil,omitempty" name:"ProduceTime"`

	// pulsar消息id
	PulsarMsgId *string `json:"PulsarMsgId,omitnil,omitempty" name:"PulsarMsgId"`

	// 死信重发次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterResendTimes *int64 `json:"DeadLetterResendTimes,omitnil,omitempty" name:"DeadLetterResendTimes"`

	// 死信重发成功次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResendSuccessCount *int64 `json:"ResendSuccessCount,omitnil,omitempty" name:"ResendSuccessCount"`
}

type RocketMQNamespace struct {
	// 命名空间名称，3-64个字符，只能包含字母、数字、“-”及“_”
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 已废弃，未消费消息的保留时间，以毫秒单位，范围60秒到15天
	Ttl *uint64 `json:"Ttl,omitnil,omitempty" name:"Ttl"`

	// 消息持久化后保留的时间，以毫秒单位
	RetentionTime *uint64 `json:"RetentionTime,omitnil,omitempty" name:"RetentionTime"`

	// 说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 公网接入点地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicEndpoint *string `json:"PublicEndpoint,omitnil,omitempty" name:"PublicEndpoint"`

	// VPC接入点地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcEndpoint *string `json:"VpcEndpoint,omitnil,omitempty" name:"VpcEndpoint"`

	// 内部接入点地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternalEndpoint *string `json:"InternalEndpoint,omitnil,omitempty" name:"InternalEndpoint"`
}

type RocketMQSmoothMigrationTaskItem struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 源集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceClusterName *string `json:"SourceClusterName,omitnil,omitempty" name:"SourceClusterName"`

	// 目标集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 网络连接类型，
	// PUBLIC 公网
	// VPC 私有网络
	// OTHER 其他
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectionType *string `json:"ConnectionType,omitnil,omitempty" name:"ConnectionType"`

	// 源集群NameServer地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceNameServer *string `json:"SourceNameServer,omitnil,omitempty" name:"SourceNameServer"`

	// 任务状态
	// Configuration 迁移配置
	// SourceConnecting 连接源集群中
	// MetaDataImport 元数据导入
	// EndpointSetup 切换接入点
	// ServiceMigration 切流中
	// Completed 已完成
	// Cancelled 已取消
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`
}

type RocketMQSubscription struct {
	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 主题类型：
	// Normal 普通,
	// GlobalOrder 全局顺序,
	// PartitionedOrder 局部顺序,
	// Transaction 事务消息,
	// DelayScheduled 延时消息,
	// Retry 重试,
	// DeadLetter 死信
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分区数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionNum *int64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// 过滤模式，TAG，SQL
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpressionType *string `json:"ExpressionType,omitnil,omitempty" name:"ExpressionType"`

	// 过滤表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubString *string `json:"SubString,omitnil,omitempty" name:"SubString"`

	// 订阅状态：
	// 0，订阅关系一致
	// 1，订阅关系不一致
	// 2，未知
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 消费堆积数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerLag *int64 `json:"ConsumerLag,omitnil,omitempty" name:"ConsumerLag"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 消费组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`

	// 是否在线
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsOnline *bool `json:"IsOnline,omitnil,omitempty" name:"IsOnline"`

	// 消费类型
	// 0: 广播消费
	// 1: 集群消费
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumeType *int64 `json:"ConsumeType,omitnil,omitempty" name:"ConsumeType"`

	// 订阅一致性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Consistency *int64 `json:"Consistency,omitnil,omitempty" name:"Consistency"`

	// 最后消费进度更新时间，秒为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *int64 `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// 最大重试次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRetryTimes *int64 `json:"MaxRetryTimes,omitnil,omitempty" name:"MaxRetryTimes"`

	// 协议类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientProtocol *string `json:"ClientProtocol,omitnil,omitempty" name:"ClientProtocol"`

	// 客户端订阅详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientSubscriptionInfos []*ClientSubscriptionInfo `json:"ClientSubscriptionInfos,omitnil,omitempty" name:"ClientSubscriptionInfos"`
}

type RocketMQTopic struct {
	// 主题名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 主题的类别，为枚举类型，Normal，GlobalOrder，PartitionedOrder，Transaction，Retry及DeadLetter
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 订阅组数量
	GroupNum *uint64 `json:"GroupNum,omitnil,omitempty" name:"GroupNum"`

	// 说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 读写分区数
	PartitionNum *uint64 `json:"PartitionNum,omitnil,omitempty" name:"PartitionNum"`

	// 创建时间，以毫秒为单位
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 创建时间，以毫秒为单位
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 最后写入时间，单位为秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *int64 `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// 订阅数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubscriptionCount *int64 `json:"SubscriptionCount,omitnil,omitempty" name:"SubscriptionCount"`

	// 订阅关系列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubscriptionData []*RocketMQSubscription `json:"SubscriptionData,omitnil,omitempty" name:"SubscriptionData"`
}

type RocketMQTopicConfig struct {
	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 主题类型：
	// Normal，普通
	// PartitionedOrder, 分区顺序
	// Transaction，事务消息
	// DelayScheduled，延迟/定时消息
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分区个数
	Partitions *int64 `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type RocketMQTopicConfigOutput struct {
	// 命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 主题名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 主题类型：
	// Normal，普通
	// GlobalOrder， 全局顺序
	// PartitionedOrder, 分区顺序
	// Transaction，事务消息
	// DelayScheduled，延迟/定时消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 分区个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions *int64 `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否导入
	// 注意：此字段可能返回 null，表示取不到有效值。
	Imported *bool `json:"Imported,omitnil,omitempty" name:"Imported"`
}

type RocketMQTopicDistribution struct {
	// topic类型
	TopicType *string `json:"TopicType,omitnil,omitempty" name:"TopicType"`

	// topic数量
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type RocketMQVipInstance struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例版本
	InstanceVersion *string `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`

	// 实例状态，0表示创建中，1表示正常，2表示隔离中，3表示已销毁，4 - 异常, 5 - 发货失败，6 - 变配中，7 - 变配失败
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 节点数量
	NodeCount *uint64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 实例配置规格名称
	ConfigDisplay *string `json:"ConfigDisplay,omitnil,omitempty" name:"ConfigDisplay"`

	// 峰值TPS
	MaxTps *uint64 `json:"MaxTps,omitnil,omitempty" name:"MaxTps"`

	// 峰值带宽，Mbps为单位
	MaxBandWidth *uint64 `json:"MaxBandWidth,omitnil,omitempty" name:"MaxBandWidth"`

	// 存储容量，GB为单位
	MaxStorage *uint64 `json:"MaxStorage,omitnil,omitempty" name:"MaxStorage"`

	// 实例到期时间，毫秒为单位
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 自动续费标记，0表示默认状态(用户未设置，即初始状态即手动续费)， 1表示自动续费，2表示明确不自动续费(用户设置)
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 0-后付费，1-预付费
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 实例配置ID
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 最大可设置消息保留时间，小时为单位
	MaxRetention *int64 `json:"MaxRetention,omitnil,omitempty" name:"MaxRetention"`

	// 最小可设置消息保留时间，小时为单位
	MinRetention *int64 `json:"MinRetention,omitnil,omitempty" name:"MinRetention"`

	// 实例消息保留时间，小时为单位
	Retention *int64 `json:"Retention,omitnil,omitempty" name:"Retention"`

	// 是否开启ACL鉴权
	AclEnabled *bool `json:"AclEnabled,omitnil,omitempty" name:"AclEnabled"`

	// 销毁时间
	DestroyTime *uint64 `json:"DestroyTime,omitnil,omitempty" name:"DestroyTime"`
}

type Role struct {
	// 角色名称。
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 角色token值。
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 备注说明。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 授权类型（Cluster：集群；TopicAndGroup：主题或消费组）
	PermType *string `json:"PermType,omitnil,omitempty" name:"PermType"`
}

type SecurityPolicy struct {
	// ip或者网段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Route *string `json:"Route,omitnil,omitempty" name:"Route"`

	// 策略 true就是允许，白名单或者 false 拒绝 黑名单
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Policy *bool `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

// Predefined struct for user
type SendBatchMessagesRequestParams struct {
	// 消息要发送的topic的名字, 这里尽量需要使用topic的全路径，即：tenant/namespace/topic。如果不指定，默认使用的是：public/default
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 需要发送消息的内容
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// String 类型的 token，可以不填，系统会自动获取
	StringToken *string `json:"StringToken,omitnil,omitempty" name:"StringToken"`

	// producer 的名字，要求全局是唯一的，如果不设置，系统会自动生成
	ProducerName *string `json:"ProducerName,omitnil,omitempty" name:"ProducerName"`

	// 单位：s。消息发送的超时时间。默认值为：30s
	SendTimeout *int64 `json:"SendTimeout,omitnil,omitempty" name:"SendTimeout"`

	// 内存中允许缓存的生产消息的最大数量，默认值：1000条
	MaxPendingMessages *int64 `json:"MaxPendingMessages,omitnil,omitempty" name:"MaxPendingMessages"`

	// 每一个batch中消息的最大数量，默认值：1000条/batch
	BatchingMaxMessages *int64 `json:"BatchingMaxMessages,omitnil,omitempty" name:"BatchingMaxMessages"`

	// 每一个batch最大等待的时间，超过这个时间，不管是否达到指定的batch中消息的数量和大小，都会将该batch发送出去，默认：10ms
	BatchingMaxPublishDelay *int64 `json:"BatchingMaxPublishDelay,omitnil,omitempty" name:"BatchingMaxPublishDelay"`

	// 每一个batch中最大允许的消息的大小，默认：128KB
	BatchingMaxBytes *int64 `json:"BatchingMaxBytes,omitnil,omitempty" name:"BatchingMaxBytes"`
}

type SendBatchMessagesRequest struct {
	*tchttp.BaseRequest
	
	// 消息要发送的topic的名字, 这里尽量需要使用topic的全路径，即：tenant/namespace/topic。如果不指定，默认使用的是：public/default
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 需要发送消息的内容
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// String 类型的 token，可以不填，系统会自动获取
	StringToken *string `json:"StringToken,omitnil,omitempty" name:"StringToken"`

	// producer 的名字，要求全局是唯一的，如果不设置，系统会自动生成
	ProducerName *string `json:"ProducerName,omitnil,omitempty" name:"ProducerName"`

	// 单位：s。消息发送的超时时间。默认值为：30s
	SendTimeout *int64 `json:"SendTimeout,omitnil,omitempty" name:"SendTimeout"`

	// 内存中允许缓存的生产消息的最大数量，默认值：1000条
	MaxPendingMessages *int64 `json:"MaxPendingMessages,omitnil,omitempty" name:"MaxPendingMessages"`

	// 每一个batch中消息的最大数量，默认值：1000条/batch
	BatchingMaxMessages *int64 `json:"BatchingMaxMessages,omitnil,omitempty" name:"BatchingMaxMessages"`

	// 每一个batch最大等待的时间，超过这个时间，不管是否达到指定的batch中消息的数量和大小，都会将该batch发送出去，默认：10ms
	BatchingMaxPublishDelay *int64 `json:"BatchingMaxPublishDelay,omitnil,omitempty" name:"BatchingMaxPublishDelay"`

	// 每一个batch中最大允许的消息的大小，默认：128KB
	BatchingMaxBytes *int64 `json:"BatchingMaxBytes,omitnil,omitempty" name:"BatchingMaxBytes"`
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

// Predefined struct for user
type SendBatchMessagesResponseParams struct {
	// 消息的唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageId *string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// 错误消息，返回为 ""，代表没有错误
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SendBatchMessagesResponse struct {
	*tchttp.BaseResponse
	Response *SendBatchMessagesResponseParams `json:"Response"`
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

// Predefined struct for user
type SendCmqMsgRequestParams struct {
	// 队列名
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 消息内容
	MsgContent *string `json:"MsgContent,omitnil,omitempty" name:"MsgContent"`

	// 延迟时间。单位为秒，默认值为0秒，最大不能超过队列配置的消息最长未确认时间。
	DelaySeconds *int64 `json:"DelaySeconds,omitnil,omitempty" name:"DelaySeconds"`
}

type SendCmqMsgRequest struct {
	*tchttp.BaseRequest
	
	// 队列名
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 消息内容
	MsgContent *string `json:"MsgContent,omitnil,omitempty" name:"MsgContent"`

	// 延迟时间。单位为秒，默认值为0秒，最大不能超过队列配置的消息最长未确认时间。
	DelaySeconds *int64 `json:"DelaySeconds,omitnil,omitempty" name:"DelaySeconds"`
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

// Predefined struct for user
type SendCmqMsgResponseParams struct {
	// true表示发送成功
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 消息id
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SendCmqMsgResponse struct {
	*tchttp.BaseResponse
	Response *SendCmqMsgResponseParams `json:"Response"`
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

// Predefined struct for user
type SendMessagesRequestParams struct {
	// 消息要发送的topic的名字, 这里尽量需要使用topic的全路径，即：tenant/namespace/topic。如果不指定，默认使用的是：public/default
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 要发送的消息的内容
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// Token 是用来做鉴权使用的，可以不填，系统会自动获取
	StringToken *string `json:"StringToken,omitnil,omitempty" name:"StringToken"`

	// 设置 producer 的名字，要求全局唯一。该参数建议用户无需手动配置，此时系统会随机生成，如果手动设置有可能会造成创建 Producer 失败进而导致消息发送失败。
	// 该参数主要用于某些特定场景下，只允许特定的 Producer 生产消息时设置，用户的大部分场景使用不到该特性。
	ProducerName *string `json:"ProducerName,omitnil,omitempty" name:"ProducerName"`

	// 设置消息发送的超时时间，默认为30s
	SendTimeout *int64 `json:"SendTimeout,omitnil,omitempty" name:"SendTimeout"`

	// 内存中缓存的最大的生产消息的数量，默认为1000条
	MaxPendingMessages *int64 `json:"MaxPendingMessages,omitnil,omitempty" name:"MaxPendingMessages"`
}

type SendMessagesRequest struct {
	*tchttp.BaseRequest
	
	// 消息要发送的topic的名字, 这里尽量需要使用topic的全路径，即：tenant/namespace/topic。如果不指定，默认使用的是：public/default
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 要发送的消息的内容
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// Token 是用来做鉴权使用的，可以不填，系统会自动获取
	StringToken *string `json:"StringToken,omitnil,omitempty" name:"StringToken"`

	// 设置 producer 的名字，要求全局唯一。该参数建议用户无需手动配置，此时系统会随机生成，如果手动设置有可能会造成创建 Producer 失败进而导致消息发送失败。
	// 该参数主要用于某些特定场景下，只允许特定的 Producer 生产消息时设置，用户的大部分场景使用不到该特性。
	ProducerName *string `json:"ProducerName,omitnil,omitempty" name:"ProducerName"`

	// 设置消息发送的超时时间，默认为30s
	SendTimeout *int64 `json:"SendTimeout,omitnil,omitempty" name:"SendTimeout"`

	// 内存中缓存的最大的生产消息的数量，默认为1000条
	MaxPendingMessages *int64 `json:"MaxPendingMessages,omitnil,omitempty" name:"MaxPendingMessages"`
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

// Predefined struct for user
type SendMessagesResponseParams struct {
	// 消息的messageID, 是全局唯一的，用来标识消息的元数据信息
	MessageId *string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// 返回的错误消息，如果返回为 “”，说明没有错误
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SendMessagesResponse struct {
	*tchttp.BaseResponse
	Response *SendMessagesResponseParams `json:"Response"`
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

// Predefined struct for user
type SendMsgRequestParams struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名称，如果是分区topic需要指定具体分区，如果没有指定则默认发到0分区，例如：my_topic-partition-0。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 消息内容，不能为空且大小不得大于5242880个byte。
	MsgContent *string `json:"MsgContent,omitnil,omitempty" name:"MsgContent"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
}

type SendMsgRequest struct {
	*tchttp.BaseRequest
	
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名称，如果是分区topic需要指定具体分区，如果没有指定则默认发到0分区，例如：my_topic-partition-0。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 消息内容，不能为空且大小不得大于5242880个byte。
	MsgContent *string `json:"MsgContent,omitnil,omitempty" name:"MsgContent"`

	// Pulsar 集群的ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`
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

// Predefined struct for user
type SendMsgResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SendMsgResponse struct {
	*tchttp.BaseResponse
	Response *SendMsgResponseParams `json:"Response"`
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

// Predefined struct for user
type SendRocketMQMessageRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// topic名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 信息内容
	MsgBody *string `json:"MsgBody,omitnil,omitempty" name:"MsgBody"`

	// 消息key信息
	MsgKey *string `json:"MsgKey,omitnil,omitempty" name:"MsgKey"`

	// 消息tag信息
	MsgTag *string `json:"MsgTag,omitnil,omitempty" name:"MsgTag"`
}

type SendRocketMQMessageRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// topic名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 信息内容
	MsgBody *string `json:"MsgBody,omitnil,omitempty" name:"MsgBody"`

	// 消息key信息
	MsgKey *string `json:"MsgKey,omitnil,omitempty" name:"MsgKey"`

	// 消息tag信息
	MsgTag *string `json:"MsgTag,omitnil,omitempty" name:"MsgTag"`
}

func (r *SendRocketMQMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendRocketMQMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "TopicName")
	delete(f, "MsgBody")
	delete(f, "MsgKey")
	delete(f, "MsgTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendRocketMQMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendRocketMQMessageResponseParams struct {
	// 发送结果
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 消息ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SendRocketMQMessageResponse struct {
	*tchttp.BaseResponse
	Response *SendRocketMQMessageResponseParams `json:"Response"`
}

func (r *SendRocketMQMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendRocketMQMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerLog struct {
	// 存储时间。
	SaveTime *string `json:"SaveTime,omitnil,omitempty" name:"SaveTime"`

	// 状态。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type SetRocketMQPublicAccessPointRequestParams struct {
	// 集群ID，当前只支持专享集群
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开启或关闭访问
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 带宽大小，开启或者调整公网时必须指定，Mbps为单位
	Bandwidth *uint64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 付费模式，开启公网时必须指定，0为按小时计费，1为包年包月，当前只支持按小时计费
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 公网访问安全规则列表，Enabled为true时必须传入
	Rules []*PublicAccessRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 公网是否按流量计费
	BillingFlow *bool `json:"BillingFlow,omitnil,omitempty" name:"BillingFlow"`
}

type SetRocketMQPublicAccessPointRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID，当前只支持专享集群
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开启或关闭访问
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 带宽大小，开启或者调整公网时必须指定，Mbps为单位
	Bandwidth *uint64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 付费模式，开启公网时必须指定，0为按小时计费，1为包年包月，当前只支持按小时计费
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 公网访问安全规则列表，Enabled为true时必须传入
	Rules []*PublicAccessRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 公网是否按流量计费
	BillingFlow *bool `json:"BillingFlow,omitnil,omitempty" name:"BillingFlow"`
}

func (r *SetRocketMQPublicAccessPointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetRocketMQPublicAccessPointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Enabled")
	delete(f, "Bandwidth")
	delete(f, "PayMode")
	delete(f, "Rules")
	delete(f, "BillingFlow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetRocketMQPublicAccessPointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetRocketMQPublicAccessPointResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetRocketMQPublicAccessPointResponse struct {
	*tchttp.BaseResponse
	Response *SetRocketMQPublicAccessPointResponseParams `json:"Response"`
}

func (r *SetRocketMQPublicAccessPointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetRocketMQPublicAccessPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Sort struct {
	// 排序字段
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 升序ASC，降序DESC
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type Subscription struct {
	// 主题名称。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 消费者开始连接的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectedSince *string `json:"ConnectedSince,omitnil,omitempty" name:"ConnectedSince"`

	// 消费者地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerAddr *string `json:"ConsumerAddr,omitnil,omitempty" name:"ConsumerAddr"`

	// 消费者数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerCount *string `json:"ConsumerCount,omitnil,omitempty" name:"ConsumerCount"`

	// 消费者名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerName *string `json:"ConsumerName,omitnil,omitempty" name:"ConsumerName"`

	// 堆积的消息数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgBacklog *string `json:"MsgBacklog,omitnil,omitempty" name:"MsgBacklog"`

	// 于TTL，此订阅下没有被发送而是被丢弃的比例。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateExpired *string `json:"MsgRateExpired,omitnil,omitempty" name:"MsgRateExpired"`

	// 消费者每秒分发消息的数量之和。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateOut *string `json:"MsgRateOut,omitnil,omitempty" name:"MsgRateOut"`

	// 消费者每秒消息的byte。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputOut *string `json:"MsgThroughputOut,omitnil,omitempty" name:"MsgThroughputOut"`

	// 订阅名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`

	// 消费者集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerSets []*Consumer `json:"ConsumerSets,omitnil,omitempty" name:"ConsumerSets"`

	// 是否在线。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsOnline *bool `json:"IsOnline,omitnil,omitempty" name:"IsOnline"`

	// 消费进度集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumersScheduleSets []*ConsumersSchedule `json:"ConsumersScheduleSets,omitnil,omitempty" name:"ConsumersScheduleSets"`

	// 备注。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最近修改时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 订阅类型，Exclusive，Shared，Failover， Key_Shared，空或NULL表示未知，
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubType *string `json:"SubType,omitnil,omitempty" name:"SubType"`

	// 是否由于未 ack 数到达上限而被 block
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlockedSubscriptionOnUnackedMsgs *bool `json:"BlockedSubscriptionOnUnackedMsgs,omitnil,omitempty" name:"BlockedSubscriptionOnUnackedMsgs"`

	// 未 ack 消息数上限
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxUnackedMsgNum *int64 `json:"MaxUnackedMsgNum,omitnil,omitempty" name:"MaxUnackedMsgNum"`
}

type SubscriptionTopic struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 订阅名称。
	SubscriptionName *string `json:"SubscriptionName,omitnil,omitempty" name:"SubscriptionName"`
}

type Tag struct {
	// 标签的key的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签的Value的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type Topic struct {
	// 最后一次间隔内发布消息的平均byte大小。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AverageMsgSize *string `json:"AverageMsgSize,omitnil,omitempty" name:"AverageMsgSize"`

	// 消费者数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerCount *string `json:"ConsumerCount,omitnil,omitempty" name:"ConsumerCount"`

	// 被记录下来的消息总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastConfirmedEntry *string `json:"LastConfirmedEntry,omitnil,omitempty" name:"LastConfirmedEntry"`

	// 最后一个ledger创建的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastLedgerCreatedTimestamp *string `json:"LastLedgerCreatedTimestamp,omitnil,omitempty" name:"LastLedgerCreatedTimestamp"`

	// 本地和复制的发布者每秒发布消息的速率。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateIn *string `json:"MsgRateIn,omitnil,omitempty" name:"MsgRateIn"`

	// 本地和复制的消费者每秒分发消息的数量之和。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateOut *string `json:"MsgRateOut,omitnil,omitempty" name:"MsgRateOut"`

	// 本地和复制的发布者每秒发布消息的byte。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputIn *string `json:"MsgThroughputIn,omitnil,omitempty" name:"MsgThroughputIn"`

	// 本地和复制的消费者每秒分发消息的byte。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputOut *string `json:"MsgThroughputOut,omitnil,omitempty" name:"MsgThroughputOut"`

	// 被记录下来的消息总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumberOfEntries *string `json:"NumberOfEntries,omitnil,omitempty" name:"NumberOfEntries"`

	// 分区数<=0：topic下无子分区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions *int64 `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 生产者数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProducerCount *string `json:"ProducerCount,omitnil,omitempty" name:"ProducerCount"`

	// 以byte计算的所有消息存储总量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSize *string `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`

	// 分区topic里面的子分区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubTopicSets []*PartitionsTopic `json:"SubTopicSets,omitnil,omitempty" name:"SubTopicSets"`

	// topic类型描述：
	// 0：普通消息；
	// 1：全局顺序消息；
	// 2：局部顺序消息；
	// 3：重试队列；
	// 4：死信队列；
	// 5：事务消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicType *uint64 `json:"TopicType,omitnil,omitempty" name:"TopicType"`

	// 环境（命名空间）名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 说明，128个字符以内。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最近修改时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 生产者上限。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProducerLimit *string `json:"ProducerLimit,omitnil,omitempty" name:"ProducerLimit"`

	// 消费者上限。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerLimit *string `json:"ConsumerLimit,omitnil,omitempty" name:"ConsumerLimit"`

	// 0: 非持久非分区
	// 1: 非持久分区
	// 2: 持久非分区
	// 3: 持久分区
	// 注意：此字段可能返回 null，表示取不到有效值。
	PulsarTopicType *int64 `json:"PulsarTopicType,omitnil,omitempty" name:"PulsarTopicType"`

	// 未消费消息过期时间，单位：秒
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgTTL *uint64 `json:"MsgTTL,omitnil,omitempty" name:"MsgTTL"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 用户自定义的租户别名，如果没有，会复用专业集群 ID
	Tenant *string `json:"Tenant,omitnil,omitempty" name:"Tenant"`

	// 是否开启异常消费者隔离
	IsolateConsumerEnable *bool `json:"IsolateConsumerEnable,omitnil,omitempty" name:"IsolateConsumerEnable"`

	// 消费者 Ack 超时时间，单位：秒
	AckTimeOut *int64 `json:"AckTimeOut,omitnil,omitempty" name:"AckTimeOut"`
}

type TopicRecord struct {
	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitnil,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type TopicStats struct {
	// 所属Broker节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	BrokerName *string `json:"BrokerName,omitnil,omitempty" name:"BrokerName"`

	// 队列编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueId *int64 `json:"QueueId,omitnil,omitempty" name:"QueueId"`

	// 最小位点
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinOffset *int64 `json:"MinOffset,omitnil,omitempty" name:"MinOffset"`

	// 最大位点
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxOffset *int64 `json:"MaxOffset,omitnil,omitempty" name:"MaxOffset"`

	// 消息条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageCount *int64 `json:"MessageCount,omitnil,omitempty" name:"MessageCount"`

	// 消息最后写入时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTimestamp *int64 `json:"LastUpdateTimestamp,omitnil,omitempty" name:"LastUpdateTimestamp"`
}

type Topic_Simplification struct {
	// 主题名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 0: 非持久非分区
	// 1: 非持久分区
	// 2: 持久非分区
	// 3: 持久分区
	// 注意：此字段可能返回 null，表示取不到有效值。
	PulsarTopicType *int64 `json:"PulsarTopicType,omitnil,omitempty" name:"PulsarTopicType"`
}

type TraceResult struct {
	// 阶段
	Stage *string `json:"Stage,omitnil,omitempty" name:"Stage"`

	// 内容详情
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`
}

// Predefined struct for user
type UnbindCmqDeadLetterRequestParams struct {
	// 死信策略源队列名称，调用本接口会清空该队列的死信队列策略。
	SourceQueueName *string `json:"SourceQueueName,omitnil,omitempty" name:"SourceQueueName"`
}

type UnbindCmqDeadLetterRequest struct {
	*tchttp.BaseRequest
	
	// 死信策略源队列名称，调用本接口会清空该队列的死信队列策略。
	SourceQueueName *string `json:"SourceQueueName,omitnil,omitempty" name:"SourceQueueName"`
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

// Predefined struct for user
type UnbindCmqDeadLetterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindCmqDeadLetterResponse struct {
	*tchttp.BaseResponse
	Response *UnbindCmqDeadLetterResponseParams `json:"Response"`
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

// Predefined struct for user
type VerifyRocketMQConsumeRequestParams struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 消费组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 消息id
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 客户端ID
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type VerifyRocketMQConsumeRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 命名空间
	NamespaceId *string `json:"NamespaceId,omitnil,omitempty" name:"NamespaceId"`

	// 消费组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 消息id
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 客户端ID
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

func (r *VerifyRocketMQConsumeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyRocketMQConsumeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "GroupId")
	delete(f, "MsgId")
	delete(f, "ClientId")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyRocketMQConsumeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyRocketMQConsumeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VerifyRocketMQConsumeResponse struct {
	*tchttp.BaseResponse
	Response *VerifyRocketMQConsumeResponseParams `json:"Response"`
}

func (r *VerifyRocketMQConsumeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyRocketMQConsumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VirtualHostQuota struct {
	// 允许创建最大vhost数
	MaxVirtualHost *int64 `json:"MaxVirtualHost,omitnil,omitempty" name:"MaxVirtualHost"`

	// 已创建vhost数
	UsedVirtualHost *int64 `json:"UsedVirtualHost,omitnil,omitempty" name:"UsedVirtualHost"`

	// 单个 vhost 下允许的最大连接数
	MaxConnectionPerVhost *int64 `json:"MaxConnectionPerVhost,omitnil,omitempty" name:"MaxConnectionPerVhost"`

	// 单个 vhost 下允许的最大交换机数
	MaxExchangePerVhost *int64 `json:"MaxExchangePerVhost,omitnil,omitempty" name:"MaxExchangePerVhost"`

	// 单个 vhost 下允许的最大队列机数
	MaxQueuePerVhost *int64 `json:"MaxQueuePerVhost,omitnil,omitempty" name:"MaxQueuePerVhost"`
}

type VpcBindRecord struct {
	// 租户Vpc Id
	UniqueVpcId *string `json:"UniqueVpcId,omitnil,omitempty" name:"UniqueVpcId"`

	// 租户Vpc子网Id
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil,omitempty" name:"UniqueSubnetId"`

	// 路由Id
	RouterId *string `json:"RouterId,omitnil,omitempty" name:"RouterId"`

	// Vpc的Id
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// Vpc的Port
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 说明，128个字符以内
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type VpcConfig struct {
	// vpc的id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

type VpcEndpointInfo struct {
	// vpc的id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// vpc接入点信息
	VpcEndpoint *string `json:"VpcEndpoint,omitnil,omitempty" name:"VpcEndpoint"`

	// vpc接入点状态 OFF/ON/CREATING/DELETING
	VpcDataStreamEndpointStatus *string `json:"VpcDataStreamEndpointStatus,omitnil,omitempty" name:"VpcDataStreamEndpointStatus"`

	// TLS加密的数据流接入点
	VpcTlsEndpoint *string `json:"VpcTlsEndpoint,omitnil,omitempty" name:"VpcTlsEndpoint"`
}

type VpcInfo struct {
	// vpc信息
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网信息
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}