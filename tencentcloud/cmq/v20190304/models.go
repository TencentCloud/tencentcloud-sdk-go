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
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ClearQueueRequest struct {
	*tchttp.BaseRequest

	// 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *ClearQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ClearQueueRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ClearQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ClearQueueResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ClearSubscriptionFilterTagsRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 订阅名字，在单个地域同一帐号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
}

func (r *ClearSubscriptionFilterTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ClearSubscriptionFilterTagsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ClearSubscriptionFilterTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearSubscriptionFilterTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ClearSubscriptionFilterTagsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateQueueRequest struct {
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

func (r *CreateQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateQueueRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建成功的queueId
		QueueId *string `json:"QueueId,omitempty" name:"QueueId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateQueueResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSubscribeRequest struct {
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
	FilterTag []*string `json:"FilterTag,omitempty" name:"FilterTag" list`

	// BindingKey数量不超过5个， 每个BindingKey长度不超过64字节，该字段表示订阅接收消息的过滤策略，每个BindingKey最多含有15个“.”， 即最多16个词组。
	BindingKey []*string `json:"BindingKey,omitempty" name:"BindingKey" list`

	// 推送内容的格式。取值：1）JSON；2）SIMPLIFIED，即raw格式。如果Protocol是queue，则取值必须为SIMPLIFIED。如果Protocol是http，两个值均可以，默认值是JSON。
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" name:"NotifyContentFormat"`
}

func (r *CreateSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSubscribeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// SubscriptionId
		SubscriptionId *string `json:"SubscriptionId,omitempty" name:"SubscriptionId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSubscribeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 消息最大长度。取值范围 1024-65536 Byte（即1-64K），默认值 65536。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// 用于指定主题的消息匹配策略。
	FilterType *uint64 `json:"FilterType,omitempty" name:"FilterType"`

	// 消息保存时间。取值范围60 - 86400 s（即1分钟 - 1天），默认值86400。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// 是否开启消息轨迹标识，true表示开启，false表示不开启，不填表示不开启。
	Trace *bool `json:"Trace,omitempty" name:"Trace"`
}

func (r *CreateTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// TopicName
		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeadLetterPolicy struct {

	// DeadLetterQueueName
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitempty" name:"DeadLetterQueueName"`

	// DeadLetterQueue
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterQueue *string `json:"DeadLetterQueue,omitempty" name:"DeadLetterQueue"`

	// Policy
	// 注意：此字段可能返回 null，表示取不到有效值。
	Policy *uint64 `json:"Policy,omitempty" name:"Policy"`

	// MaxTimeToLive
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitempty" name:"MaxTimeToLive"`

	// MaxReceiveCount
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitempty" name:"MaxReceiveCount"`
}

type DeadLetterSource struct {

	// QueueId
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueId *string `json:"QueueId,omitempty" name:"QueueId"`

	// QueueName
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

type DeleteQueueRequest struct {
	*tchttp.BaseRequest

	// 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *DeleteQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteQueueRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteQueueResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSubscribeRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 订阅名字，在单个地域同一帐号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
}

func (r *DeleteSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSubscribeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSubscribeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *DeleteTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeadLetterSourceQueuesRequest struct {
	*tchttp.BaseRequest

	// 死信队列名称
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitempty" name:"DeadLetterQueueName"`

	// 分页时本页获取主题列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页时本页获取主题的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤死信队列源队列名称，目前仅支持SourceQueueName过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeDeadLetterSourceQueuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeadLetterSourceQueuesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeadLetterSourceQueuesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 满足本次条件的队列个数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 死信队列源队列
		QueueSet []*DeadLetterSource `json:"QueueSet,omitempty" name:"QueueSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeadLetterSourceQueuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeadLetterSourceQueuesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeQueueDetailRequest struct {
	*tchttp.BaseRequest

	// 分页时本页获取队列列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页时本页获取队列的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选参数，目前支持QueueName筛选，且仅支持一个关键字
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 标签搜索
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 精确匹配QueueName
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *DescribeQueueDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeQueueDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeQueueDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// queue总数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// queue列表
		QueueSet []*QueueSet `json:"QueueSet,omitempty" name:"QueueSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQueueDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeQueueDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionDetailRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 分页时本页获取主题列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页时本页获取主题的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选参数，目前只支持SubscriptionName，且仅支持一个关键字。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeSubscriptionDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubscriptionDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Subscription属性集合
	// 注意：此字段可能返回 null，表示取不到有效值。
		SubscriptionSet []*Subscription `json:"SubscriptionSet,omitempty" name:"SubscriptionSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscriptionDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubscriptionDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicDetailRequest struct {
	*tchttp.BaseRequest

	// 分页时本页获取队列列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页时本页获取队列的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 目前只支持过滤TopicName ， 且只能填一个过滤值
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 标签匹配
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 精确匹配TopicName
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *DescribeTopicDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopicDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// TotalCount
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// TopicSet
		TopicSet []*TopicSet `json:"TopicSet,omitempty" name:"TopicSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopicDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 过滤参数的名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数值
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type ModifyQueueAttributeRequest struct {
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
}

func (r *ModifyQueueAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyQueueAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyQueueAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyQueueAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyQueueAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubscriptionAttributeRequest struct {
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
	FilterTags []*string `json:"FilterTags,omitempty" name:"FilterTags" list`

	// BindingKey数量不超过5个， 每个BindingKey长度不超过64字节，该字段表示订阅接收消息的过滤策略，每个BindingKey最多含有15个“.”， 即最多16个词组。
	BindingKey []*string `json:"BindingKey,omitempty" name:"BindingKey" list`
}

func (r *ModifySubscriptionAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubscriptionAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubscriptionAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubscriptionAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubscriptionAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicAttributeRequest struct {
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

func (r *ModifyTopicAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTopicAttributeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTopicAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTopicAttributeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueueSet struct {

	// QueueId
	QueueId *string `json:"QueueId,omitempty" name:"QueueId"`

	// QueueName
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// Qps
	// 注意：此字段可能返回 null，表示取不到有效值。
	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`

	// Bps
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bps *uint64 `json:"Bps,omitempty" name:"Bps"`

	// MaxDelaySeconds
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDelaySeconds *uint64 `json:"MaxDelaySeconds,omitempty" name:"MaxDelaySeconds"`

	// MaxMsgHeapNum
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitempty" name:"MaxMsgHeapNum"`

	// PollingWaitSeconds
	// 注意：此字段可能返回 null，表示取不到有效值。
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitempty" name:"PollingWaitSeconds"`

	// MsgRetentionSeconds
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// VisibilityTimeout
	// 注意：此字段可能返回 null，表示取不到有效值。
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitempty" name:"VisibilityTimeout"`

	// MaxMsgSize
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// RewindSeconds
	// 注意：此字段可能返回 null，表示取不到有效值。
	RewindSeconds *uint64 `json:"RewindSeconds,omitempty" name:"RewindSeconds"`

	// CreateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// LastModifyTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyTime *uint64 `json:"LastModifyTime,omitempty" name:"LastModifyTime"`

	// ActiveMsgNum
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveMsgNum *uint64 `json:"ActiveMsgNum,omitempty" name:"ActiveMsgNum"`

	// InactiveMsgNum
	// 注意：此字段可能返回 null，表示取不到有效值。
	InactiveMsgNum *uint64 `json:"InactiveMsgNum,omitempty" name:"InactiveMsgNum"`

	// DelayMsgNum
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelayMsgNum *uint64 `json:"DelayMsgNum,omitempty" name:"DelayMsgNum"`

	// RewindMsgNum
	// 注意：此字段可能返回 null，表示取不到有效值。
	RewindMsgNum *uint64 `json:"RewindMsgNum,omitempty" name:"RewindMsgNum"`

	// MinMsgTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinMsgTime *uint64 `json:"MinMsgTime,omitempty" name:"MinMsgTime"`

	// Transaction
	// 注意：此字段可能返回 null，表示取不到有效值。
	Transaction *bool `json:"Transaction,omitempty" name:"Transaction"`

	// DeadLetterSource
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterSource []*DeadLetterSource `json:"DeadLetterSource,omitempty" name:"DeadLetterSource" list`

	// DeadLetterPolicy
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterPolicy *DeadLetterPolicy `json:"DeadLetterPolicy,omitempty" name:"DeadLetterPolicy"`

	// TransactionPolicy
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransactionPolicy *TransactionPolicy `json:"TransactionPolicy,omitempty" name:"TransactionPolicy"`

	// 创建者uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// 消息轨迹表示，true表示开启，false表示不开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trace *bool `json:"Trace,omitempty" name:"Trace"`
}

type RewindQueueRequest struct {
	*tchttp.BaseRequest

	// 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 设定该时间，则（Batch）receiveMessage接口，会按照生产消息的先后顺序消费该时间戳以后的消息。
	StartConsumeTime *uint64 `json:"StartConsumeTime,omitempty" name:"StartConsumeTime"`
}

func (r *RewindQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RewindQueueRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RewindQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RewindQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RewindQueueResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Subscription struct {

	// SubscriptionName
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`

	// SubscriptionId
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubscriptionId *string `json:"SubscriptionId,omitempty" name:"SubscriptionId"`

	// TopicOwner
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicOwner *uint64 `json:"TopicOwner,omitempty" name:"TopicOwner"`

	// MsgCount
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgCount *uint64 `json:"MsgCount,omitempty" name:"MsgCount"`

	// LastModifyTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyTime *uint64 `json:"LastModifyTime,omitempty" name:"LastModifyTime"`

	// CreateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// BindingKey
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindingKey []*string `json:"BindingKey,omitempty" name:"BindingKey" list`

	// Endpoint
	// 注意：此字段可能返回 null，表示取不到有效值。
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// FilterTags
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterTags []*string `json:"FilterTags,omitempty" name:"FilterTags" list`

	// Protocol
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// NotifyStrategy
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" name:"NotifyStrategy"`

	// NotifyContentFormat
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" name:"NotifyContentFormat"`
}

type Tag struct {

	// 标签Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TopicSet struct {

	// TopicId
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// TopicName
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// MsgRetentionSeconds
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// MaxMsgSize
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// Qps
	// 注意：此字段可能返回 null，表示取不到有效值。
	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`

	// FilterType
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterType *uint64 `json:"FilterType,omitempty" name:"FilterType"`

	// CreateTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// LastModifyTime
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyTime *uint64 `json:"LastModifyTime,omitempty" name:"LastModifyTime"`

	// MsgCount
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgCount *uint64 `json:"MsgCount,omitempty" name:"MsgCount"`

	// CreateUin
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Tags
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// 主题是否开启消息轨迹，true表示开启，false表示不开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trace *bool `json:"Trace,omitempty" name:"Trace"`
}

type TransactionPolicy struct {

	// FirstQueryInterval
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitempty" name:"FirstQueryInterval"`

	// MaxQueryCount
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitempty" name:"MaxQueryCount"`
}

type UnbindDeadLetterRequest struct {
	*tchttp.BaseRequest

	// 死信策略源队列名称，调用本接口会清空该队列的死信队列策略。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *UnbindDeadLetterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnbindDeadLetterRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UnbindDeadLetterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindDeadLetterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UnbindDeadLetterResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
