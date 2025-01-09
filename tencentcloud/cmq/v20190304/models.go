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

package v20190304

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type DeadLetterPolicy struct {
	// 死信队列。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterQueue *string `json:"DeadLetterQueue,omitnil,omitempty" name:"DeadLetterQueue"`

	// 死信队列名字。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitnil,omitempty" name:"DeadLetterQueueName"`

	// 最大未消费过期时间。Policy为1时必选。范围300-43200，单位秒，需要小于消息最大保留时间MsgRetentionSeconds。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitnil,omitempty" name:"MaxTimeToLive"`

	// 死信队列策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Policy *uint64 `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 最大接收次数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitnil,omitempty" name:"MaxReceiveCount"`
}

type DeadLetterSource struct {
	// 消息队列ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueId *string `json:"QueueId,omitnil,omitempty" name:"QueueId"`

	// 消息队列名字。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`
}

// Predefined struct for user
type DescribeQueueDetailRequestParams struct {
	// 标签搜索
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 分页时本页获取队列的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 精确匹配QueueName
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 筛选参数，目前支持QueueName筛选，且仅支持一个关键字
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页时本页获取队列列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeQueueDetailRequest struct {
	*tchttp.BaseRequest
	
	// 标签搜索
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 分页时本页获取队列的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 精确匹配QueueName
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 筛选参数，目前支持QueueName筛选，且仅支持一个关键字
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页时本页获取队列列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeQueueDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQueueDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	delete(f, "Limit")
	delete(f, "QueueName")
	delete(f, "Filters")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQueueDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQueueDetailResponseParams struct {
	// 总队列数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 队列详情列表。
	QueueSet []*QueueSet `json:"QueueSet,omitnil,omitempty" name:"QueueSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeQueueDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQueueDetailResponseParams `json:"Response"`
}

func (r *DescribeQueueDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQueueDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicDetailRequestParams struct {
	// 标签匹配。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 分页时本页获取队列的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 精确匹配TopicName。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 目前只支持过滤TopicName ， 且只能填一个过滤值。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页时本页获取队列列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeTopicDetailRequest struct {
	*tchttp.BaseRequest
	
	// 标签匹配。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 分页时本页获取队列的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 精确匹配TopicName。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 目前只支持过滤TopicName ， 且只能填一个过滤值。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页时本页获取队列列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeTopicDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	delete(f, "Limit")
	delete(f, "TopicName")
	delete(f, "Filters")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicDetailResponseParams struct {
	// 主题列表总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 主题详情列表。
	TopicSet []*TopicSet `json:"TopicSet,omitnil,omitempty" name:"TopicSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicDetailResponseParams `json:"Response"`
}

func (r *DescribeTopicDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 数值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 过滤参数的名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type QueueSet struct {
	// 消息队列ID。
	QueueId *string `json:"QueueId,omitnil,omitempty" name:"QueueId"`

	// 回溯队列的消息回溯时间最大值，取值范围0 - 43200秒，0表示不开启消息回溯。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RewindSeconds *uint64 `json:"RewindSeconds,omitnil,omitempty" name:"RewindSeconds"`

	// 创建者Uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUin *uint64 `json:"CreateUin,omitnil,omitempty" name:"CreateUin"`

	// 最后一次修改队列属性的时间。返回 Unix 时间戳，精确到秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyTime *uint64 `json:"LastModifyTime,omitnil,omitempty" name:"LastModifyTime"`

	// 消息可见性超时。取值范围1 - 43200秒（即12小时内），默认值30。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitnil,omitempty" name:"VisibilityTimeout"`

	// 消息队列名字。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 消息轨迹。true表示开启，false表示不开启。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trace *bool `json:"Trace,omitnil,omitempty" name:"Trace"`

	// 关联的标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 已调用 DelMsg 接口删除，但还在回溯保留时间内的消息数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RewindMsgNum *uint64 `json:"RewindMsgNum,omitnil,omitempty" name:"RewindMsgNum"`

	// 飞行消息最大保留时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDelaySeconds *uint64 `json:"MaxDelaySeconds,omitnil,omitempty" name:"MaxDelaySeconds"`

	// 事务消息策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransactionPolicy *TransactionPolicy `json:"TransactionPolicy,omitnil,omitempty" name:"TransactionPolicy"`

	// 消息保留周期。取值范围60-1296000秒（1min-15天），默认值345600秒（4 天）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil,omitempty" name:"MsgRetentionSeconds"`

	// 延迟消息数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelayMsgNum *uint64 `json:"DelayMsgNum,omitnil,omitempty" name:"DelayMsgNum"`

	// 最大堆积消息数。取值范围在公测期间为 1,000,000 - 10,000,000，正式上线后范围可达到 1000,000-1000,000,000。默认取值在公测期间为 10,000,000，正式上线后为 100,000,000。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitnil,omitempty" name:"MaxMsgHeapNum"`

	// 消息接收长轮询等待时间。取值范围0 - 30秒，默认值0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitnil,omitempty" name:"PollingWaitSeconds"`

	// 带宽限制。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bps *uint64 `json:"Bps,omitnil,omitempty" name:"Bps"`

	// 在队列中处于 Inactive 状态（正处于被消费状态）的消息总数，为近似值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InactiveMsgNum *uint64 `json:"InactiveMsgNum,omitnil,omitempty" name:"InactiveMsgNum"`

	// 死信队列策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterPolicy *DeadLetterPolicy `json:"DeadLetterPolicy,omitnil,omitempty" name:"DeadLetterPolicy"`

	// 在队列中处于 Active 状态（不处于被消费状态）的消息总数，为近似值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveMsgNum *uint64 `json:"ActiveMsgNum,omitnil,omitempty" name:"ActiveMsgNum"`

	// 消息最大长度。取值范围1024 - 1048576 Byte（即1K - 1024K），默认值65536。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil,omitempty" name:"MaxMsgSize"`

	// 消息最小未消费时间，单位为秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinMsgTime *uint64 `json:"MinMsgTime,omitnil,omitempty" name:"MinMsgTime"`

	// 死信队列。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterSource []*DeadLetterSource `json:"DeadLetterSource,omitnil,omitempty" name:"DeadLetterSource"`

	// 事务消息队列。true表示是事务消息，false表示不是事务消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Transaction *bool `json:"Transaction,omitnil,omitempty" name:"Transaction"`

	// 每秒钟生产消息条数的限制，消费消息的大小是该值的1.1倍。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Qps *uint64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// 队列的创建时间。返回 Unix 时间戳，精确到秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否迁移到新版本。0 表示仅同步元数据，1 表示迁移中，2 表示已经迁移完毕，3 表示回切状态，曾经迁移过，4 未迁移。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Migrate *int64 `json:"Migrate,omitnil,omitempty" name:"Migrate"`
}

type Tag struct {
	// 标签Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TopicSet struct {
	// 当前该主题中消息数目（消息堆积数）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgCount *uint64 `json:"MsgCount,omitnil,omitempty" name:"MsgCount"`

	// 主题的 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 消息最大长度。取值范围1024 - 1048576Byte（即1 - 1024K），默认值为65536。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil,omitempty" name:"MaxMsgSize"`

	// 消息轨迹。true表示开启，false表示不开启。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trace *bool `json:"Trace,omitnil,omitempty" name:"Trace"`

	// 关联的标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 创建者 Uin，CAM 鉴权 resource 由该字段组合而成。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUin *uint64 `json:"CreateUin,omitnil,omitempty" name:"CreateUin"`

	// 描述用户创建订阅时选择的过滤策略：
	// FilterType = 1表示用户使用 FilterTag 标签过滤;
	// FilterType = 2表示用户使用 BindingKey 过滤。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterType *uint64 `json:"FilterType,omitnil,omitempty" name:"FilterType"`

	// 主题名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 最后一次修改主题属性的时间。返回 Unix 时间戳，精确到秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyTime *uint64 `json:"LastModifyTime,omitnil,omitempty" name:"LastModifyTime"`

	// 消息在主题中最长存活时间，从发送到该主题开始经过此参数指定的时间后，不论消息是否被成功推送给用户都将被删除，单位为秒。固定为一天（86400秒），该属性不能修改。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil,omitempty" name:"MsgRetentionSeconds"`

	// 每秒钟发布消息的条数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Qps *uint64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// 主题的创建时间。返回 Unix 时间戳，精确到秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否迁移到新版本。0 表示未迁移，1 表示迁移中，2 表示已经迁移完毕，3 表示回切状态，曾经迁移过，4 未知状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Migrate *int64 `json:"Migrate,omitnil,omitempty" name:"Migrate"`
}

type TransactionPolicy struct {
	// 最大查询次数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitnil,omitempty" name:"MaxQueryCount"`

	// 第一次回查时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitnil,omitempty" name:"FirstQueryInterval"`
}