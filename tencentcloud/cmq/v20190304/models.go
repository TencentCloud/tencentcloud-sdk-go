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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type ClearQueueRequestParams struct {
	// 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

type ClearQueueRequest struct {
	*tchttp.BaseRequest
	
	// 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *ClearQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClearQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearQueueResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ClearQueueResponse struct {
	*tchttp.BaseResponse
	Response *ClearQueueResponseParams `json:"Response"`
}

func (r *ClearQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearSubscriptionFilterTagsRequestParams struct {
	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 订阅名字，在单个地域同一帐号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearSubscriptionFilterTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "SubscriptionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClearSubscriptionFilterTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearSubscriptionFilterTagsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ClearSubscriptionFilterTagsResponse struct {
	*tchttp.BaseResponse
	Response *ClearSubscriptionFilterTagsResponseParams `json:"Response"`
}

func (r *ClearSubscriptionFilterTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearSubscriptionFilterTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQueueRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQueueRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQueueResponseParams struct {
	// 创建成功的queueId
	QueueId *string `json:"QueueId,omitempty" name:"QueueId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateQueueResponse struct {
	*tchttp.BaseResponse
	Response *CreateQueueResponseParams `json:"Response"`
}

func (r *CreateQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubscribeRequestParams struct {
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
	FilterTag []*string `json:"FilterTag,omitempty" name:"FilterTag"`

	// BindingKey数量不超过5个， 每个BindingKey长度不超过64字节，该字段表示订阅接收消息的过滤策略，每个BindingKey最多含有15个“.”， 即最多16个词组。
	BindingKey []*string `json:"BindingKey,omitempty" name:"BindingKey"`

	// 推送内容的格式。取值：1）JSON；2）SIMPLIFIED，即raw格式。如果Protocol是queue，则取值必须为SIMPLIFIED。如果Protocol是http，两个值均可以，默认值是JSON。
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" name:"NotifyContentFormat"`
}

func (r *CreateSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubscribeRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubscribeResponseParams struct {
	// SubscriptionId
	SubscriptionId *string `json:"SubscriptionId,omitempty" name:"SubscriptionId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *CreateSubscribeResponseParams `json:"Response"`
}

func (r *CreateSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicRequestParams struct {
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

type CreateTopicRequest struct {
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
	delete(f, "TopicName")
	delete(f, "MaxMsgSize")
	delete(f, "FilterType")
	delete(f, "MsgRetentionSeconds")
	delete(f, "Trace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicResponseParams struct {
	// TopicName
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type DeadLetterPolicy struct {
	// 死信队列。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterQueue *string `json:"DeadLetterQueue,omitempty" name:"DeadLetterQueue"`

	// 死信队列名字。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitempty" name:"DeadLetterQueueName"`

	// 最大未消费过期时间。Policy为1时必选。范围300-43200，单位秒，需要小于消息最大保留时间MsgRetentionSeconds。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitempty" name:"MaxTimeToLive"`

	// 死信队列策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Policy *uint64 `json:"Policy,omitempty" name:"Policy"`

	// 最大接收次数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitempty" name:"MaxReceiveCount"`
}

type DeadLetterSource struct {
	// 消息队列ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueId *string `json:"QueueId,omitempty" name:"QueueId"`

	// 消息队列名字。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

// Predefined struct for user
type DeleteQueueRequestParams struct {
	// 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQueueResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteQueueResponse struct {
	*tchttp.BaseResponse
	Response *DeleteQueueResponseParams `json:"Response"`
}

func (r *DeleteQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSubscribeRequestParams struct {
	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 订阅名字，在单个地域同一帐号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "SubscriptionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSubscribeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSubscribeResponseParams `json:"Response"`
}

func (r *DeleteSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicRequestParams struct {
	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTopicResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTopicResponseParams `json:"Response"`
}

func (r *DeleteTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeadLetterSourceQueuesRequestParams struct {
	// 死信队列名称
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitempty" name:"DeadLetterQueueName"`

	// 分页时本页获取主题列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页时本页获取主题的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤死信队列源队列名称，目前仅支持SourceQueueName过滤
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDeadLetterSourceQueuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeadLetterSourceQueuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeadLetterQueueName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeadLetterSourceQueuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeadLetterSourceQueuesResponseParams struct {
	// 满足本次条件的队列个数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 死信队列源队列
	QueueSet []*DeadLetterSource `json:"QueueSet,omitempty" name:"QueueSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDeadLetterSourceQueuesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeadLetterSourceQueuesResponseParams `json:"Response"`
}

func (r *DescribeDeadLetterSourceQueuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeadLetterSourceQueuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQueueDetailRequestParams struct {
	// 分页时本页获取队列列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页时本页获取队列的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选参数，目前支持QueueName筛选，且仅支持一个关键字
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 标签搜索
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 精确匹配QueueName
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

type DescribeQueueDetailRequest struct {
	*tchttp.BaseRequest
	
	// 分页时本页获取队列列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页时本页获取队列的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选参数，目前支持QueueName筛选，且仅支持一个关键字
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 标签搜索
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 精确匹配QueueName
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
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
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "TagKey")
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQueueDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQueueDetailResponseParams struct {
	// 总队列数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 队列详情列表。
	QueueSet []*QueueSet `json:"QueueSet,omitempty" name:"QueueSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeSubscriptionDetailRequestParams struct {
	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 分页时本页获取主题列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页时本页获取主题的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选参数，目前只支持SubscriptionName，且仅支持一个关键字。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeSubscriptionDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscriptionDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubscriptionDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubscriptionDetailResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// Subscription属性集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubscriptionSet []*Subscription `json:"SubscriptionSet,omitempty" name:"SubscriptionSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSubscriptionDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubscriptionDetailResponseParams `json:"Response"`
}

func (r *DescribeSubscriptionDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscriptionDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicDetailRequestParams struct {
	// 分页时本页获取队列列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页时本页获取队列的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 目前只支持过滤TopicName ， 且只能填一个过滤值。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 标签匹配。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 精确匹配TopicName。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type DescribeTopicDetailRequest struct {
	*tchttp.BaseRequest
	
	// 分页时本页获取队列列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 分页时本页获取队列的个数，如果不传递该参数，则该参数默认为20，最大值为50。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 目前只支持过滤TopicName ， 且只能填一个过滤值。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 标签匹配。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 精确匹配TopicName。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
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
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "TagKey")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicDetailResponseParams struct {
	// 主题列表总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 主题详情列表。
	TopicSet []*TopicSet `json:"TopicSet,omitempty" name:"TopicSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 过滤参数的名字
	Name *string `json:"Name,omitempty" name:"Name"`
}

// Predefined struct for user
type ModifyQueueAttributeRequestParams struct {
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQueueAttributeRequest) FromJsonString(s string) error {
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyQueueAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQueueAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyQueueAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyQueueAttributeResponseParams `json:"Response"`
}

func (r *ModifyQueueAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQueueAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscriptionAttributeRequestParams struct {
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
	FilterTags []*string `json:"FilterTags,omitempty" name:"FilterTags"`

	// BindingKey数量不超过5个， 每个BindingKey长度不超过64字节，该字段表示订阅接收消息的过滤策略，每个BindingKey最多含有15个“.”， 即最多16个词组。
	BindingKey []*string `json:"BindingKey,omitempty" name:"BindingKey"`
}

func (r *ModifySubscriptionAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscriptionAttributeRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySubscriptionAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySubscriptionAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifySubscriptionAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifySubscriptionAttributeResponseParams `json:"Response"`
}

func (r *ModifySubscriptionAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySubscriptionAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicAttributeRequestParams struct {
	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 消息最大长度。取值范围1024 - 65536 Byte（即1 - 64K），默认值65536。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// 消息保存时间。取值范围60 - 86400 s（即1分钟 - 1天），默认值86400。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// 是否开启消息轨迹标识，true表示开启，false表示不开启，不填表示不开启。
	Trace *bool `json:"Trace,omitempty" name:"Trace"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "MaxMsgSize")
	delete(f, "MsgRetentionSeconds")
	delete(f, "Trace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTopicAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTopicAttributeResponseParams `json:"Response"`
}

func (r *ModifyTopicAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueueSet struct {
	// 消息队列ID。
	QueueId *string `json:"QueueId,omitempty" name:"QueueId"`

	// 回溯队列的消息回溯时间最大值，取值范围0 - 43200秒，0表示不开启消息回溯。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RewindSeconds *uint64 `json:"RewindSeconds,omitempty" name:"RewindSeconds"`

	// 创建者Uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 最后一次修改队列属性的时间。返回 Unix 时间戳，精确到秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyTime *uint64 `json:"LastModifyTime,omitempty" name:"LastModifyTime"`

	// 消息可见性超时。取值范围1 - 43200秒（即12小时内），默认值30。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitempty" name:"VisibilityTimeout"`

	// 消息队列名字。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 消息轨迹。true表示开启，false表示不开启。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trace *bool `json:"Trace,omitempty" name:"Trace"`

	// 关联的标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 已调用 DelMsg 接口删除，但还在回溯保留时间内的消息数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RewindMsgNum *uint64 `json:"RewindMsgNum,omitempty" name:"RewindMsgNum"`

	// 飞行消息最大保留时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDelaySeconds *uint64 `json:"MaxDelaySeconds,omitempty" name:"MaxDelaySeconds"`

	// 事务消息策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransactionPolicy *TransactionPolicy `json:"TransactionPolicy,omitempty" name:"TransactionPolicy"`

	// 消息保留周期。取值范围60-1296000秒（1min-15天），默认值345600秒（4 天）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// 延迟消息数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelayMsgNum *uint64 `json:"DelayMsgNum,omitempty" name:"DelayMsgNum"`

	// 最大堆积消息数。取值范围在公测期间为 1,000,000 - 10,000,000，正式上线后范围可达到 1000,000-1000,000,000。默认取值在公测期间为 10,000,000，正式上线后为 100,000,000。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitempty" name:"MaxMsgHeapNum"`

	// 消息接收长轮询等待时间。取值范围0 - 30秒，默认值0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitempty" name:"PollingWaitSeconds"`

	// 带宽限制。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bps *uint64 `json:"Bps,omitempty" name:"Bps"`

	// 在队列中处于 Inactive 状态（正处于被消费状态）的消息总数，为近似值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InactiveMsgNum *uint64 `json:"InactiveMsgNum,omitempty" name:"InactiveMsgNum"`

	// 死信队列策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterPolicy *DeadLetterPolicy `json:"DeadLetterPolicy,omitempty" name:"DeadLetterPolicy"`

	// 在队列中处于 Active 状态（不处于被消费状态）的消息总数，为近似值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveMsgNum *uint64 `json:"ActiveMsgNum,omitempty" name:"ActiveMsgNum"`

	// 消息最大长度。取值范围1024 - 1048576 Byte（即1K - 1024K），默认值65536。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// 消息最小未消费时间，单位为秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinMsgTime *uint64 `json:"MinMsgTime,omitempty" name:"MinMsgTime"`

	// 死信队列。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterSource []*DeadLetterSource `json:"DeadLetterSource,omitempty" name:"DeadLetterSource"`

	// 事务消息队列。true表示是事务消息，false表示不是事务消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Transaction *bool `json:"Transaction,omitempty" name:"Transaction"`

	// 每秒钟生产消息条数的限制，消费消息的大小是该值的1.1倍。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`

	// 队列的创建时间。返回 Unix 时间戳，精确到秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 是否迁移到新版本。0 表示仅同步元数据，1 表示迁移中，2 表示已经迁移完毕，3 表示回切状态，曾经迁移过，4 未迁移。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Migrate *int64 `json:"Migrate,omitempty" name:"Migrate"`
}

// Predefined struct for user
type RewindQueueRequestParams struct {
	// 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 设定该时间，则（Batch）receiveMessage接口，会按照生产消息的先后顺序消费该时间戳以后的消息。
	StartConsumeTime *uint64 `json:"StartConsumeTime,omitempty" name:"StartConsumeTime"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RewindQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	delete(f, "StartConsumeTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RewindQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RewindQueueResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RewindQueueResponse struct {
	*tchttp.BaseResponse
	Response *RewindQueueResponseParams `json:"Response"`
}

func (r *RewindQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RewindQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Subscription struct {
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

	// 最后一次修改订阅属性的时间。返回 Unix 时间戳，精确到秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyTime *uint64 `json:"LastModifyTime,omitempty" name:"LastModifyTime"`

	// 订阅的创建时间。返回 Unix 时间戳，精确到秒。
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

type Tag struct {
	// 标签Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TopicSet struct {
	// 当前该主题中消息数目（消息堆积数）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgCount *uint64 `json:"MsgCount,omitempty" name:"MsgCount"`

	// 主题的 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 消息最大长度。取值范围1024 - 1048576Byte（即1 - 1024K），默认值为65536。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// 消息轨迹。true表示开启，false表示不开启。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trace *bool `json:"Trace,omitempty" name:"Trace"`

	// 关联的标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 创建者 Uin，CAM 鉴权 resource 由该字段组合而成。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 描述用户创建订阅时选择的过滤策略：
	// FilterType = 1表示用户使用 FilterTag 标签过滤;
	// FilterType = 2表示用户使用 BindingKey 过滤。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterType *uint64 `json:"FilterType,omitempty" name:"FilterType"`

	// 主题名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 最后一次修改主题属性的时间。返回 Unix 时间戳，精确到秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifyTime *uint64 `json:"LastModifyTime,omitempty" name:"LastModifyTime"`

	// 消息在主题中最长存活时间，从发送到该主题开始经过此参数指定的时间后，不论消息是否被成功推送给用户都将被删除，单位为秒。固定为一天（86400秒），该属性不能修改。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// 每秒钟发布消息的条数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`

	// 主题的创建时间。返回 Unix 时间戳，精确到秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 是否迁移到新版本。0 表示未迁移，1 表示迁移中，2 表示已经迁移完毕，3 表示回切状态，曾经迁移过，4 未知状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Migrate *int64 `json:"Migrate,omitempty" name:"Migrate"`
}

type TransactionPolicy struct {
	// 最大查询次数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitempty" name:"MaxQueryCount"`

	// 第一次回查时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitempty" name:"FirstQueryInterval"`
}

// Predefined struct for user
type UnbindDeadLetterRequestParams struct {
	// 死信策略源队列名称，调用本接口会清空该队列的死信队列策略。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindDeadLetterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindDeadLetterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindDeadLetterResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnbindDeadLetterResponse struct {
	*tchttp.BaseResponse
	Response *UnbindDeadLetterResponseParams `json:"Response"`
}

func (r *UnbindDeadLetterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindDeadLetterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}