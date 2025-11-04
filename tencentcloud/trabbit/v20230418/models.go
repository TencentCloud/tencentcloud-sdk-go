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

package v20230418

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type CreateRabbitMQServerlessBindingRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 源exchange
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标类型,取值queue或exchange
	DestinationType *string `json:"DestinationType,omitnil,omitempty" name:"DestinationType"`

	// 目标队列或者交换机
	Destination *string `json:"Destination,omitnil,omitempty" name:"Destination"`

	// 绑定key
	RoutingKey *string `json:"RoutingKey,omitnil,omitempty" name:"RoutingKey"`
}

type CreateRabbitMQServerlessBindingRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 源exchange
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标类型,取值queue或exchange
	DestinationType *string `json:"DestinationType,omitnil,omitempty" name:"DestinationType"`

	// 目标队列或者交换机
	Destination *string `json:"Destination,omitnil,omitempty" name:"Destination"`

	// 绑定key
	RoutingKey *string `json:"RoutingKey,omitnil,omitempty" name:"RoutingKey"`
}

func (r *CreateRabbitMQServerlessBindingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQServerlessBindingRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRabbitMQServerlessBindingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQServerlessBindingResponseParams struct {
	// 队列名称
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 路由关系Id
	BindingId *int64 `json:"BindingId,omitnil,omitempty" name:"BindingId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRabbitMQServerlessBindingResponse struct {
	*tchttp.BaseResponse
	Response *CreateRabbitMQServerlessBindingResponseParams `json:"Response"`
}

func (r *CreateRabbitMQServerlessBindingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQServerlessBindingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQServerlessExchangeRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VHost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// exchange 名称
	ExchangeName *string `json:"ExchangeName,omitnil,omitempty" name:"ExchangeName"`

	// exchange 类型, 支持 "fanout","direct","topic","headers"
	ExchangeType *string `json:"ExchangeType,omitnil,omitempty" name:"ExchangeType"`

	// exchange 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否为持久化 exchange, 当集群重启时,将会清除所有该字段为"false"的 exchange
	Durable *bool `json:"Durable,omitnil,omitempty" name:"Durable"`

	// 是否自动删除该 exchange, 如果为 "true", 当解绑所有当前 exchange 上的路由关系时, 该 exchange 将会被自动删除
	AutoDelete *bool `json:"AutoDelete,omitnil,omitempty" name:"AutoDelete"`

	// 是否为内部 exchange, 如果为 "true", 则无法直接投递消息到该 exchange, 需要在路由设置中通过其他 exchange 进行转发
	Internal *bool `json:"Internal,omitnil,omitempty" name:"Internal"`

	// 替代 exchange, 如果消息无法发送到当前 exchange, 就会发送到该替代 exchange
	AlternateExchange *string `json:"AlternateExchange,omitnil,omitempty" name:"AlternateExchange"`

	// 延迟类型的exchange背后对应的exchange类型, 支持 "fanout","direct","topic","headers"
	DelayedExchangeType *string `json:"DelayedExchangeType,omitnil,omitempty" name:"DelayedExchangeType"`
}

type CreateRabbitMQServerlessExchangeRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VHost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// exchange 名称
	ExchangeName *string `json:"ExchangeName,omitnil,omitempty" name:"ExchangeName"`

	// exchange 类型, 支持 "fanout","direct","topic","headers"
	ExchangeType *string `json:"ExchangeType,omitnil,omitempty" name:"ExchangeType"`

	// exchange 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否为持久化 exchange, 当集群重启时,将会清除所有该字段为"false"的 exchange
	Durable *bool `json:"Durable,omitnil,omitempty" name:"Durable"`

	// 是否自动删除该 exchange, 如果为 "true", 当解绑所有当前 exchange 上的路由关系时, 该 exchange 将会被自动删除
	AutoDelete *bool `json:"AutoDelete,omitnil,omitempty" name:"AutoDelete"`

	// 是否为内部 exchange, 如果为 "true", 则无法直接投递消息到该 exchange, 需要在路由设置中通过其他 exchange 进行转发
	Internal *bool `json:"Internal,omitnil,omitempty" name:"Internal"`

	// 替代 exchange, 如果消息无法发送到当前 exchange, 就会发送到该替代 exchange
	AlternateExchange *string `json:"AlternateExchange,omitnil,omitempty" name:"AlternateExchange"`

	// 延迟类型的exchange背后对应的exchange类型, 支持 "fanout","direct","topic","headers"
	DelayedExchangeType *string `json:"DelayedExchangeType,omitnil,omitempty" name:"DelayedExchangeType"`
}

func (r *CreateRabbitMQServerlessExchangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQServerlessExchangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "ExchangeName")
	delete(f, "ExchangeType")
	delete(f, "Remark")
	delete(f, "Durable")
	delete(f, "AutoDelete")
	delete(f, "Internal")
	delete(f, "AlternateExchange")
	delete(f, "DelayedExchangeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRabbitMQServerlessExchangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQServerlessExchangeResponseParams struct {
	// exchange 名称
	ExchangeName *string `json:"ExchangeName,omitnil,omitempty" name:"ExchangeName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRabbitMQServerlessExchangeResponse struct {
	*tchttp.BaseResponse
	Response *CreateRabbitMQServerlessExchangeResponseParams `json:"Response"`
}

func (r *CreateRabbitMQServerlessExchangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQServerlessExchangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQServerlessQueueRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VHost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 队列名称
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 只支持 classic
	QueueType *string `json:"QueueType,omitnil,omitempty" name:"QueueType"`

	// 持久标记,classic类型必传,quorum类型无需传入固定为true
	Durable *bool `json:"Durable,omitnil,omitempty" name:"Durable"`

	// 自动清除,classic类型必传,quorum类型无需传入固定为false
	AutoDelete *bool `json:"AutoDelete,omitnil,omitempty" name:"AutoDelete"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// MessageTTL参数,classic类型专用
	MessageTTL *int64 `json:"MessageTTL,omitnil,omitempty" name:"MessageTTL"`

	// AutoExpire参数，单位为 ms，队列在指定时间内没有被使用，将会被删除
	AutoExpire *int64 `json:"AutoExpire,omitnil,omitempty" name:"AutoExpire"`

	// MaxLength参数。队列可以容纳的最大条数。若超出上限，将根据 overview behavior 处理
	MaxLength *int64 `json:"MaxLength,omitnil,omitempty" name:"MaxLength"`

	// MaxLengthBytes参数。若超出上限，将根据 overview behavior 处理。
	MaxLengthBytes *int64 `json:"MaxLengthBytes,omitnil,omitempty" name:"MaxLengthBytes"`

	// DeliveryLimit参数,quorum类型专用
	DeliveryLimit *int64 `json:"DeliveryLimit,omitnil,omitempty" name:"DeliveryLimit"`

	// OverflowBehaviour参数,取值为drop-head, reject-publish或reject-publish-dlx
	OverflowBehaviour *string `json:"OverflowBehaviour,omitnil,omitempty" name:"OverflowBehaviour"`

	// DeadLetterExchange参数。可将过期或被拒绝的消息投往指定的死信 exchange。
	DeadLetterExchange *string `json:"DeadLetterExchange,omitnil,omitempty" name:"DeadLetterExchange"`

	// DeadLetterRoutingKey参数。只能包含字母、数字、"."、"-"，"@"，"_"
	DeadLetterRoutingKey *string `json:"DeadLetterRoutingKey,omitnil,omitempty" name:"DeadLetterRoutingKey"`

	// SingleActiveConsumer参数。若开启，需确保每次有且只有一个消费者从队列中消费
	SingleActiveConsumer *bool `json:"SingleActiveConsumer,omitnil,omitempty" name:"SingleActiveConsumer"`

	// MaximumPriority参数,classic类型专用
	MaximumPriority *int64 `json:"MaximumPriority,omitnil,omitempty" name:"MaximumPriority"`

	// LazyMode参数,classic类型专用
	LazyMode *bool `json:"LazyMode,omitnil,omitempty" name:"LazyMode"`

	// MasterLocator参数,classic类型专用,取值为min-masters,client-local或random
	MasterLocator *string `json:"MasterLocator,omitnil,omitempty" name:"MasterLocator"`

	// MaxInMemoryLength参数，quorum类型专用。quorum 队列的内存中最大消息数量
	MaxInMemoryLength *int64 `json:"MaxInMemoryLength,omitnil,omitempty" name:"MaxInMemoryLength"`

	// MaxInMemoryBytes参数，quorum类型专用。quorum 队列的内存中最大数总消息大小
	MaxInMemoryBytes *int64 `json:"MaxInMemoryBytes,omitnil,omitempty" name:"MaxInMemoryBytes"`

	// Node参数，非必填，指定创建 queue 所在节点
	Node *string `json:"Node,omitnil,omitempty" name:"Node"`

	// 仲裁队列死信一致性策略，at-most-once、at-least-once，默认是at-most-once
	DeadLetterStrategy *string `json:"DeadLetterStrategy,omitnil,omitempty" name:"DeadLetterStrategy"`

	// 仲裁队列的领导者选举策略，client-local、balanced，默认是client-local
	QueueLeaderLocator *string `json:"QueueLeaderLocator,omitnil,omitempty" name:"QueueLeaderLocator"`

	// 仲裁队列的初始副本组大小，默认3
	QuorumInitialGroupSize *int64 `json:"QuorumInitialGroupSize,omitnil,omitempty" name:"QuorumInitialGroupSize"`
}

type CreateRabbitMQServerlessQueueRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VHost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 队列名称
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 只支持 classic
	QueueType *string `json:"QueueType,omitnil,omitempty" name:"QueueType"`

	// 持久标记,classic类型必传,quorum类型无需传入固定为true
	Durable *bool `json:"Durable,omitnil,omitempty" name:"Durable"`

	// 自动清除,classic类型必传,quorum类型无需传入固定为false
	AutoDelete *bool `json:"AutoDelete,omitnil,omitempty" name:"AutoDelete"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// MessageTTL参数,classic类型专用
	MessageTTL *int64 `json:"MessageTTL,omitnil,omitempty" name:"MessageTTL"`

	// AutoExpire参数，单位为 ms，队列在指定时间内没有被使用，将会被删除
	AutoExpire *int64 `json:"AutoExpire,omitnil,omitempty" name:"AutoExpire"`

	// MaxLength参数。队列可以容纳的最大条数。若超出上限，将根据 overview behavior 处理
	MaxLength *int64 `json:"MaxLength,omitnil,omitempty" name:"MaxLength"`

	// MaxLengthBytes参数。若超出上限，将根据 overview behavior 处理。
	MaxLengthBytes *int64 `json:"MaxLengthBytes,omitnil,omitempty" name:"MaxLengthBytes"`

	// DeliveryLimit参数,quorum类型专用
	DeliveryLimit *int64 `json:"DeliveryLimit,omitnil,omitempty" name:"DeliveryLimit"`

	// OverflowBehaviour参数,取值为drop-head, reject-publish或reject-publish-dlx
	OverflowBehaviour *string `json:"OverflowBehaviour,omitnil,omitempty" name:"OverflowBehaviour"`

	// DeadLetterExchange参数。可将过期或被拒绝的消息投往指定的死信 exchange。
	DeadLetterExchange *string `json:"DeadLetterExchange,omitnil,omitempty" name:"DeadLetterExchange"`

	// DeadLetterRoutingKey参数。只能包含字母、数字、"."、"-"，"@"，"_"
	DeadLetterRoutingKey *string `json:"DeadLetterRoutingKey,omitnil,omitempty" name:"DeadLetterRoutingKey"`

	// SingleActiveConsumer参数。若开启，需确保每次有且只有一个消费者从队列中消费
	SingleActiveConsumer *bool `json:"SingleActiveConsumer,omitnil,omitempty" name:"SingleActiveConsumer"`

	// MaximumPriority参数,classic类型专用
	MaximumPriority *int64 `json:"MaximumPriority,omitnil,omitempty" name:"MaximumPriority"`

	// LazyMode参数,classic类型专用
	LazyMode *bool `json:"LazyMode,omitnil,omitempty" name:"LazyMode"`

	// MasterLocator参数,classic类型专用,取值为min-masters,client-local或random
	MasterLocator *string `json:"MasterLocator,omitnil,omitempty" name:"MasterLocator"`

	// MaxInMemoryLength参数，quorum类型专用。quorum 队列的内存中最大消息数量
	MaxInMemoryLength *int64 `json:"MaxInMemoryLength,omitnil,omitempty" name:"MaxInMemoryLength"`

	// MaxInMemoryBytes参数，quorum类型专用。quorum 队列的内存中最大数总消息大小
	MaxInMemoryBytes *int64 `json:"MaxInMemoryBytes,omitnil,omitempty" name:"MaxInMemoryBytes"`

	// Node参数，非必填，指定创建 queue 所在节点
	Node *string `json:"Node,omitnil,omitempty" name:"Node"`

	// 仲裁队列死信一致性策略，at-most-once、at-least-once，默认是at-most-once
	DeadLetterStrategy *string `json:"DeadLetterStrategy,omitnil,omitempty" name:"DeadLetterStrategy"`

	// 仲裁队列的领导者选举策略，client-local、balanced，默认是client-local
	QueueLeaderLocator *string `json:"QueueLeaderLocator,omitnil,omitempty" name:"QueueLeaderLocator"`

	// 仲裁队列的初始副本组大小，默认3
	QuorumInitialGroupSize *int64 `json:"QuorumInitialGroupSize,omitnil,omitempty" name:"QuorumInitialGroupSize"`
}

func (r *CreateRabbitMQServerlessQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQServerlessQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "QueueName")
	delete(f, "QueueType")
	delete(f, "Durable")
	delete(f, "AutoDelete")
	delete(f, "Remark")
	delete(f, "MessageTTL")
	delete(f, "AutoExpire")
	delete(f, "MaxLength")
	delete(f, "MaxLengthBytes")
	delete(f, "DeliveryLimit")
	delete(f, "OverflowBehaviour")
	delete(f, "DeadLetterExchange")
	delete(f, "DeadLetterRoutingKey")
	delete(f, "SingleActiveConsumer")
	delete(f, "MaximumPriority")
	delete(f, "LazyMode")
	delete(f, "MasterLocator")
	delete(f, "MaxInMemoryLength")
	delete(f, "MaxInMemoryBytes")
	delete(f, "Node")
	delete(f, "DeadLetterStrategy")
	delete(f, "QueueLeaderLocator")
	delete(f, "QuorumInitialGroupSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRabbitMQServerlessQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQServerlessQueueResponseParams struct {
	// 队列名称
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRabbitMQServerlessQueueResponse struct {
	*tchttp.BaseResponse
	Response *CreateRabbitMQServerlessQueueResponseParams `json:"Response"`
}

func (r *CreateRabbitMQServerlessQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQServerlessQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQServerlessUserRequestParams struct {
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，登录时使用
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 密码，登录时使用
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// serverless 实例该字段无效
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 该用户的最大连接数，不填写则不限制
	MaxConnections *int64 `json:"MaxConnections,omitnil,omitempty" name:"MaxConnections"`

	// 该用户的最大channel数，不填写则不限制
	MaxChannels *int64 `json:"MaxChannels,omitnil,omitempty" name:"MaxChannels"`
}

type CreateRabbitMQServerlessUserRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，登录时使用
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 密码，登录时使用
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// serverless 实例该字段无效
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 该用户的最大连接数，不填写则不限制
	MaxConnections *int64 `json:"MaxConnections,omitnil,omitempty" name:"MaxConnections"`

	// 该用户的最大channel数，不填写则不限制
	MaxChannels *int64 `json:"MaxChannels,omitnil,omitempty" name:"MaxChannels"`
}

func (r *CreateRabbitMQServerlessUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQServerlessUserRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRabbitMQServerlessUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQServerlessUserResponseParams struct {
	// 用户名，登录时使用
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRabbitMQServerlessUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateRabbitMQServerlessUserResponseParams `json:"Response"`
}

func (r *CreateRabbitMQServerlessUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQServerlessUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQServerlessVirtualHostRequestParams struct {
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost名称
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 消息轨迹开关,true打开,false关闭,默认关闭
	TraceFlag *bool `json:"TraceFlag,omitnil,omitempty" name:"TraceFlag"`

	// 是否创建镜像队列策略，默认值 true
	MirrorQueuePolicyFlag *bool `json:"MirrorQueuePolicyFlag,omitnil,omitempty" name:"MirrorQueuePolicyFlag"`
}

type CreateRabbitMQServerlessVirtualHostRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost名称
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 消息轨迹开关,true打开,false关闭,默认关闭
	TraceFlag *bool `json:"TraceFlag,omitnil,omitempty" name:"TraceFlag"`

	// 是否创建镜像队列策略，默认值 true
	MirrorQueuePolicyFlag *bool `json:"MirrorQueuePolicyFlag,omitnil,omitempty" name:"MirrorQueuePolicyFlag"`
}

func (r *CreateRabbitMQServerlessVirtualHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQServerlessVirtualHostRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRabbitMQServerlessVirtualHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQServerlessVirtualHostResponseParams struct {
	// vhost名称
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRabbitMQServerlessVirtualHostResponse struct {
	*tchttp.BaseResponse
	Response *CreateRabbitMQServerlessVirtualHostResponseParams `json:"Response"`
}

func (r *CreateRabbitMQServerlessVirtualHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQServerlessVirtualHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQServerlessBindingRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 路由关系Id
	BindingId *int64 `json:"BindingId,omitnil,omitempty" name:"BindingId"`
}

type DeleteRabbitMQServerlessBindingRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 路由关系Id
	BindingId *int64 `json:"BindingId,omitnil,omitempty" name:"BindingId"`
}

func (r *DeleteRabbitMQServerlessBindingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQServerlessBindingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "BindingId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRabbitMQServerlessBindingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQServerlessBindingResponseParams struct {
	// 队列名称
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 路由关系Id
	BindingId *int64 `json:"BindingId,omitnil,omitempty" name:"BindingId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRabbitMQServerlessBindingResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRabbitMQServerlessBindingResponseParams `json:"Response"`
}

func (r *DeleteRabbitMQServerlessBindingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQServerlessBindingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQServerlessExchangeRequestParams struct {
	// 实例 id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost 参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// exchange 名称
	ExchangeName *string `json:"ExchangeName,omitnil,omitempty" name:"ExchangeName"`
}

type DeleteRabbitMQServerlessExchangeRequest struct {
	*tchttp.BaseRequest
	
	// 实例 id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost 参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// exchange 名称
	ExchangeName *string `json:"ExchangeName,omitnil,omitempty" name:"ExchangeName"`
}

func (r *DeleteRabbitMQServerlessExchangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQServerlessExchangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "ExchangeName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRabbitMQServerlessExchangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQServerlessExchangeResponseParams struct {
	// exchange 名称
	ExchangeName *string `json:"ExchangeName,omitnil,omitempty" name:"ExchangeName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRabbitMQServerlessExchangeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRabbitMQServerlessExchangeResponseParams `json:"Response"`
}

func (r *DeleteRabbitMQServerlessExchangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQServerlessExchangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQServerlessPermissionRequestParams struct {
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，登录时使用
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// vhost名
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`
}

type DeleteRabbitMQServerlessPermissionRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，登录时使用
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// vhost名
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`
}

func (r *DeleteRabbitMQServerlessPermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQServerlessPermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "VirtualHost")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRabbitMQServerlessPermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQServerlessPermissionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRabbitMQServerlessPermissionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRabbitMQServerlessPermissionResponseParams `json:"Response"`
}

func (r *DeleteRabbitMQServerlessPermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQServerlessPermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQServerlessQueueRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 队列名称
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`
}

type DeleteRabbitMQServerlessQueueRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 队列名称
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`
}

func (r *DeleteRabbitMQServerlessQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQServerlessQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRabbitMQServerlessQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQServerlessQueueResponseParams struct {
	// 队列名称
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRabbitMQServerlessQueueResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRabbitMQServerlessQueueResponseParams `json:"Response"`
}

func (r *DeleteRabbitMQServerlessQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQServerlessQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQServerlessUserRequestParams struct {
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，登录时使用
	User *string `json:"User,omitnil,omitempty" name:"User"`
}

type DeleteRabbitMQServerlessUserRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，登录时使用
	User *string `json:"User,omitnil,omitempty" name:"User"`
}

func (r *DeleteRabbitMQServerlessUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQServerlessUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRabbitMQServerlessUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQServerlessUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRabbitMQServerlessUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRabbitMQServerlessUserResponseParams `json:"Response"`
}

func (r *DeleteRabbitMQServerlessUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQServerlessUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQServerlessVirtualHostRequestParams struct {
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost名
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`
}

type DeleteRabbitMQServerlessVirtualHostRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost名
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`
}

func (r *DeleteRabbitMQServerlessVirtualHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQServerlessVirtualHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRabbitMQServerlessVirtualHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQServerlessVirtualHostResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRabbitMQServerlessVirtualHostResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRabbitMQServerlessVirtualHostResponseParams `json:"Response"`
}

func (r *DeleteRabbitMQServerlessVirtualHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQServerlessVirtualHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessBindingsRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 分页offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键词，根据源exchange名称/目标资源名称/绑定key进行模糊搜索
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 根据源Exchange精准搜索过滤
	SourceExchange *string `json:"SourceExchange,omitnil,omitempty" name:"SourceExchange"`

	// 根据目标QueueName精准搜索过滤，和DestinationExchange过滤不可同时设置
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 根据目标Exchange精准搜索过滤，和QueueName过滤不可同时设置
	DestinationExchange *string `json:"DestinationExchange,omitnil,omitempty" name:"DestinationExchange"`
}

type DescribeRabbitMQServerlessBindingsRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 分页offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键词，根据源exchange名称/目标资源名称/绑定key进行模糊搜索
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 根据源Exchange精准搜索过滤
	SourceExchange *string `json:"SourceExchange,omitnil,omitempty" name:"SourceExchange"`

	// 根据目标QueueName精准搜索过滤，和DestinationExchange过滤不可同时设置
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 根据目标Exchange精准搜索过滤，和QueueName过滤不可同时设置
	DestinationExchange *string `json:"DestinationExchange,omitnil,omitempty" name:"DestinationExchange"`
}

func (r *DescribeRabbitMQServerlessBindingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessBindingsRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQServerlessBindingsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessBindingsResponseParams struct {
	// 路由关系列表
	BindingInfoList []*RabbitMQBindingListInfo `json:"BindingInfoList,omitnil,omitempty" name:"BindingInfoList"`

	// 数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQServerlessBindingsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQServerlessBindingsResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQServerlessBindingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessBindingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessConnectionRequestParams struct {
	// 集群实例Id	
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost名
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 按哪个字段排序，支持：channel(channel数),incoming_bytes(入流量大小),outgoing_bytes(出流量大小)
	SortElement *string `json:"SortElement,omitnil,omitempty" name:"SortElement"`

	// 排序方式：ASC,DESC
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 分页参数，从第几条数据开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 一页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 连接名模糊搜索
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeRabbitMQServerlessConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例Id	
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost名
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 按哪个字段排序，支持：channel(channel数),incoming_bytes(入流量大小),outgoing_bytes(出流量大小)
	SortElement *string `json:"SortElement,omitnil,omitempty" name:"SortElement"`

	// 排序方式：ASC,DESC
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 分页参数，从第几条数据开始
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 一页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 连接名模糊搜索
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DescribeRabbitMQServerlessConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "SortElement")
	delete(f, "SortType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQServerlessConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessConnectionResponseParams struct {
	// 返回连接数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 连接详情列表
	Connections []*RabbitMQConnection `json:"Connections,omitnil,omitempty" name:"Connections"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQServerlessConnectionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQServerlessConnectionResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQServerlessConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessConsumersRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 队列名
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 搜索关键词
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// channelId
	Channel *string `json:"Channel,omitnil,omitempty" name:"Channel"`
}

type DescribeRabbitMQServerlessConsumersRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 队列名
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 搜索关键词
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// channelId
	Channel *string `json:"Channel,omitnil,omitempty" name:"Channel"`
}

func (r *DescribeRabbitMQServerlessConsumersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessConsumersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "QueueName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SearchWord")
	delete(f, "Channel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQServerlessConsumersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessConsumersResponseParams struct {
	// 消费者列表信息
	ConsumerInfoList []*RabbitMQConsumersListInfo `json:"ConsumerInfoList,omitnil,omitempty" name:"ConsumerInfoList"`

	// 数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQServerlessConsumersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQServerlessConsumersResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQServerlessConsumersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessConsumersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessExchangeDetailRequestParams struct {
	// 实例 id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost 参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// exchange 名称
	ExchangeName *string `json:"ExchangeName,omitnil,omitempty" name:"ExchangeName"`
}

type DescribeRabbitMQServerlessExchangeDetailRequest struct {
	*tchttp.BaseRequest
	
	// 实例 id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost 参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// exchange 名称
	ExchangeName *string `json:"ExchangeName,omitnil,omitempty" name:"ExchangeName"`
}

func (r *DescribeRabbitMQServerlessExchangeDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessExchangeDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "ExchangeName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQServerlessExchangeDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessExchangeDetailResponseParams struct {
	// exchange 名
	ExchangeName *string `json:"ExchangeName,omitnil,omitempty" name:"ExchangeName"`

	// 备注说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否为持久化 exchange, 当集群重启时, 将会清除所有该字段为 "false" 的 exchange
	Durable *bool `json:"Durable,omitnil,omitempty" name:"Durable"`

	// 是否自动删除该 exchange, 如果为 "true", 当解绑所有当前 exchange 上的路由关系时, 该 exchange 将会被自动删除
	AutoDelete *bool `json:"AutoDelete,omitnil,omitempty" name:"AutoDelete"`

	// 是否为内部 exchange, 如果为 "true", 则无法直接投递消息到该 exchange, 需要在路由设置中通过其他 exchange 进行转发
	Internal *bool `json:"Internal,omitnil,omitempty" name:"Internal"`

	// 替代 exchange, 如果消息没有匹配当前 exchange 绑定的所有 queue 或 exchange, 就会发送到该替代 exchange
	AlternateExchange *string `json:"AlternateExchange,omitnil,omitempty" name:"AlternateExchange"`

	// exchange 类型, 支持 "fanout","direct","topic","headers"
	ExchangeType *string `json:"ExchangeType,omitnil,omitempty" name:"ExchangeType"`

	// VHost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// exchange 创建者, "system":"系统创建", "user":"用户创建"
	ExchangeCreator *string `json:"ExchangeCreator,omitnil,omitempty" name:"ExchangeCreator"`

	// 扩展参数 key-value 字符串
	Arguments *string `json:"Arguments,omitnil,omitempty" name:"Arguments"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQServerlessExchangeDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQServerlessExchangeDetailResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQServerlessExchangeDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessExchangeDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessExchangesRequestParams struct {
	// 实例 id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost 参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 分页 offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页 limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键词, 支持模糊匹配 
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 筛选 exchange 类型, 数组中每个元素为选中的过滤类型
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
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

type DescribeRabbitMQServerlessExchangesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost 参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 分页 offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页 limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键词, 支持模糊匹配 
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 筛选 exchange 类型, 数组中每个元素为选中的过滤类型
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
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

func (r *DescribeRabbitMQServerlessExchangesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessExchangesRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQServerlessExchangesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessExchangesResponseParams struct {
	// 交换机列表
	ExchangeInfoList []*RabbitMQExchangeListInfo `json:"ExchangeInfoList,omitnil,omitempty" name:"ExchangeInfoList"`

	// 交换机总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQServerlessExchangesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQServerlessExchangesResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQServerlessExchangesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessExchangesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessInstanceRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeRabbitMQServerlessInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeRabbitMQServerlessInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQServerlessInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessInstanceResponseParams struct {
	// 集群信息
	ClusterInfo *RabbitMQClusterInfo `json:"ClusterInfo,omitnil,omitempty" name:"ClusterInfo"`

	// 集群规格信息
	ClusterSpecInfo *RabbitMQClusterSpecInfo `json:"ClusterSpecInfo,omitnil,omitempty" name:"ClusterSpecInfo"`

	// vhost配额信息
	VirtualHostQuota *VirtualHostQuota `json:"VirtualHostQuota,omitnil,omitempty" name:"VirtualHostQuota"`

	// exchange配额信息
	ExchangeQuota *ExchangeQuota `json:"ExchangeQuota,omitnil,omitempty" name:"ExchangeQuota"`

	// queue配额信息
	QueueQuota *QueueQuota `json:"QueueQuota,omitnil,omitempty" name:"QueueQuota"`

	// 网络信息
	ClusterNetInfo *RabbitMQServerlessAccessInfo `json:"ClusterNetInfo,omitnil,omitempty" name:"ClusterNetInfo"`

	// 公网白名单信息
	ClusterWhiteListInfo *RabbitMQServerlessWhiteListInfo `json:"ClusterWhiteListInfo,omitnil,omitempty" name:"ClusterWhiteListInfo"`

	// user配额信息
	UserQuota *UserQuota `json:"UserQuota,omitnil,omitempty" name:"UserQuota"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQServerlessInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQServerlessInstanceResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQServerlessInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessPermissionRequestParams struct {
	// 集群实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，用于查询过滤，不传则查询全部
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// vhost名，用于查询过滤，不传则查询全部
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeRabbitMQServerlessPermissionRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，用于查询过滤，不传则查询全部
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// vhost名，用于查询过滤，不传则查询全部
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeRabbitMQServerlessPermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessPermissionRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQServerlessPermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessPermissionResponseParams struct {
	// 返回权限数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 权限详情列表
	RabbitMQPermissionList []*RabbitMQPermission `json:"RabbitMQPermissionList,omitnil,omitempty" name:"RabbitMQPermissionList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQServerlessPermissionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQServerlessPermissionResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQServerlessPermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessPermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessQueueDetailRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 队列名称
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`
}

type DescribeRabbitMQServerlessQueueDetailRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 队列名称
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`
}

func (r *DescribeRabbitMQServerlessQueueDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessQueueDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQServerlessQueueDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessQueueDetailResponseParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 队列名称
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 队列类型,取值classic或quorum
	QueueType *string `json:"QueueType,omitnil,omitempty" name:"QueueType"`

	// 在线消费者数量
	Consumers *int64 `json:"Consumers,omitnil,omitempty" name:"Consumers"`

	// 持久标记
	Durable *bool `json:"Durable,omitnil,omitempty" name:"Durable"`

	// 自动清除
	AutoDelete *bool `json:"AutoDelete,omitnil,omitempty" name:"AutoDelete"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// MessageTTL参数,classic类型专用
	MessageTTL *int64 `json:"MessageTTL,omitnil,omitempty" name:"MessageTTL"`

	// AutoExpire参数
	AutoExpire *int64 `json:"AutoExpire,omitnil,omitempty" name:"AutoExpire"`

	// MaxLength参数
	MaxLength *int64 `json:"MaxLength,omitnil,omitempty" name:"MaxLength"`

	// MaxLengthBytes参数
	MaxLengthBytes *int64 `json:"MaxLengthBytes,omitnil,omitempty" name:"MaxLengthBytes"`

	// DeliveryLimit参数,quorum类型专用
	DeliveryLimit *int64 `json:"DeliveryLimit,omitnil,omitempty" name:"DeliveryLimit"`

	// OverflowBehaviour参数,取值为drop-head, reject-publish或reject-publish-dlx
	OverflowBehaviour *string `json:"OverflowBehaviour,omitnil,omitempty" name:"OverflowBehaviour"`

	// DeadLetterExchange参数
	DeadLetterExchange *string `json:"DeadLetterExchange,omitnil,omitempty" name:"DeadLetterExchange"`

	// DeadLetterRoutingKey参数
	DeadLetterRoutingKey *string `json:"DeadLetterRoutingKey,omitnil,omitempty" name:"DeadLetterRoutingKey"`

	// SingleActiveConsumer参数
	SingleActiveConsumer *bool `json:"SingleActiveConsumer,omitnil,omitempty" name:"SingleActiveConsumer"`

	// MaximumPriority参数,classic类型专用
	MaximumPriority *int64 `json:"MaximumPriority,omitnil,omitempty" name:"MaximumPriority"`

	// LazyMode参数,classic类型专用
	LazyMode *bool `json:"LazyMode,omitnil,omitempty" name:"LazyMode"`

	// MasterLocator参数,classic类型专用
	MasterLocator *string `json:"MasterLocator,omitnil,omitempty" name:"MasterLocator"`

	// MaxInMemoryLength参数,quorum类型专用
	MaxInMemoryLength *int64 `json:"MaxInMemoryLength,omitnil,omitempty" name:"MaxInMemoryLength"`

	// MaxInMemoryBytes参数,quorum类型专用
	MaxInMemoryBytes *int64 `json:"MaxInMemoryBytes,omitnil,omitempty" name:"MaxInMemoryBytes"`

	// 创建时间戳,单位秒
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 节点
	Node *string `json:"Node,omitnil,omitempty" name:"Node"`

	// 仲裁队列死信一致性策略
	DeadLetterStrategy *string `json:"DeadLetterStrategy,omitnil,omitempty" name:"DeadLetterStrategy"`

	// 仲裁队列的领导者选举策略
	QueueLeaderLocator *string `json:"QueueLeaderLocator,omitnil,omitempty" name:"QueueLeaderLocator"`

	// 仲裁队列的初始副本组大小
	QuorumInitialGroupSize *int64 `json:"QuorumInitialGroupSize,omitnil,omitempty" name:"QuorumInitialGroupSize"`

	// 是否为独占队列
	Exclusive *bool `json:"Exclusive,omitnil,omitempty" name:"Exclusive"`

	// 生效的策略名
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 扩展参数 key-value
	Arguments *string `json:"Arguments,omitnil,omitempty" name:"Arguments"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQServerlessQueueDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQServerlessQueueDetailResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQServerlessQueueDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessQueueDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessQueuesRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键词
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 队列类型筛选，不填或 "all"：classic 和 quorum 队列；"classic"：筛选 classic 队列；"quorum"：筛选 quorum 队列
	QueueType *string `json:"QueueType,omitnil,omitempty" name:"QueueType"`

	// 排序依据的字段：
	// ConsumerNumber - 在线消费者数量；
	// MessageHeapCount - 消息堆积数；
	// MessageRateInOut - 生产消费速率之和；
	// MessageRateIn - 生产速率；
	// MessageRateOut - 消费速率；
	SortElement *string `json:"SortElement,omitnil,omitempty" name:"SortElement"`

	// 排序顺序，ascend 或 descend
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

type DescribeRabbitMQServerlessQueuesRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索关键词
	SearchWord *string `json:"SearchWord,omitnil,omitempty" name:"SearchWord"`

	// 队列类型筛选，不填或 "all"：classic 和 quorum 队列；"classic"：筛选 classic 队列；"quorum"：筛选 quorum 队列
	QueueType *string `json:"QueueType,omitnil,omitempty" name:"QueueType"`

	// 排序依据的字段：
	// ConsumerNumber - 在线消费者数量；
	// MessageHeapCount - 消息堆积数；
	// MessageRateInOut - 生产消费速率之和；
	// MessageRateIn - 生产速率；
	// MessageRateOut - 消费速率；
	SortElement *string `json:"SortElement,omitnil,omitempty" name:"SortElement"`

	// 排序顺序，ascend 或 descend
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`
}

func (r *DescribeRabbitMQServerlessQueuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessQueuesRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQServerlessQueuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessQueuesResponseParams struct {
	// 队列列表信息
	QueueInfoList []*RabbitMQQueueListInfo `json:"QueueInfoList,omitnil,omitempty" name:"QueueInfoList"`

	// 数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQServerlessQueuesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQServerlessQueuesResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQServerlessQueuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessQueuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessUserRequestParams struct {
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名检索，支持前缀匹配，后缀匹配
	SearchUser *string `json:"SearchUser,omitnil,omitempty" name:"SearchUser"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 用户名，精确查询
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 用户标签，根据标签过滤列表
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type DescribeRabbitMQServerlessUserRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名检索，支持前缀匹配，后缀匹配
	SearchUser *string `json:"SearchUser,omitnil,omitempty" name:"SearchUser"`

	// 分页Offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 用户名，精确查询
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 用户标签，根据标签过滤列表
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *DescribeRabbitMQServerlessUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessUserRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQServerlessUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessUserResponseParams struct {
	// 返回的User数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 当前已创建的RabbitMQ用户列表
	RabbitMQUserList []*RabbitMQUser `json:"RabbitMQUserList,omitnil,omitempty" name:"RabbitMQUserList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQServerlessUserResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQServerlessUserResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQServerlessUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessVirtualHostRequestParams struct {
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

type DescribeRabbitMQServerlessVirtualHostRequest struct {
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

func (r *DescribeRabbitMQServerlessVirtualHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessVirtualHostRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQServerlessVirtualHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQServerlessVirtualHostResponseParams struct {
	// 返回vhost数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// vhost详情列表
	VirtualHostList []*RabbitMQVirtualHostInfo `json:"VirtualHostList,omitnil,omitempty" name:"VirtualHostList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRabbitMQServerlessVirtualHostResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQServerlessVirtualHostResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQServerlessVirtualHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQServerlessVirtualHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExchangeQuota struct {
	// 可创建最大exchange数
	MaxExchange *int64 `json:"MaxExchange,omitnil,omitempty" name:"MaxExchange"`

	// 已创建exchange数
	UsedExchange *int64 `json:"UsedExchange,omitnil,omitempty" name:"UsedExchange"`
}

type Filter struct {
	// 过滤参数的名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type ListRabbitMQServerlessInstancesRequestParams struct {
	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 翻页的起始索引值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type ListRabbitMQServerlessInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 翻页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 翻页的起始索引值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *ListRabbitMQServerlessInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRabbitMQServerlessInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRabbitMQServerlessInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRabbitMQServerlessInstancesResponseParams struct {
	// 实例列表
	Instances []*RabbitMQServerlessInstance `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRabbitMQServerlessInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ListRabbitMQServerlessInstancesResponseParams `json:"Response"`
}

func (r *ListRabbitMQServerlessInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRabbitMQServerlessInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQServerlessExchangeRequestParams struct {
	// 实例 id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost 参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// exchange 名称
	ExchangeName *string `json:"ExchangeName,omitnil,omitempty" name:"ExchangeName"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyRabbitMQServerlessExchangeRequest struct {
	*tchttp.BaseRequest
	
	// 实例 id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost 参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// exchange 名称
	ExchangeName *string `json:"ExchangeName,omitnil,omitempty" name:"ExchangeName"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyRabbitMQServerlessExchangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQServerlessExchangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "ExchangeName")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRabbitMQServerlessExchangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQServerlessExchangeResponseParams struct {
	// exchange 名称
	ExchangeName *string `json:"ExchangeName,omitnil,omitempty" name:"ExchangeName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRabbitMQServerlessExchangeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRabbitMQServerlessExchangeResponseParams `json:"Response"`
}

func (r *ModifyRabbitMQServerlessExchangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQServerlessExchangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQServerlessInstanceRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群名
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否开启trace
	TraceFlag *bool `json:"TraceFlag,omitnil,omitempty" name:"TraceFlag"`

	// 限流生产消费比例
	SendReceiveRatio *float64 `json:"SendReceiveRatio,omitnil,omitempty" name:"SendReceiveRatio"`

	// 是否删除所有标签，默认为false
	DeleteAllTags *bool `json:"DeleteAllTags,omitnil,omitempty" name:"DeleteAllTags"`

	// 修改的实例标签列表
	InstanceTags []*RabbitMQServerlessTag `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`
}

type ModifyRabbitMQServerlessInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 集群名
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否开启trace
	TraceFlag *bool `json:"TraceFlag,omitnil,omitempty" name:"TraceFlag"`

	// 限流生产消费比例
	SendReceiveRatio *float64 `json:"SendReceiveRatio,omitnil,omitempty" name:"SendReceiveRatio"`

	// 是否删除所有标签，默认为false
	DeleteAllTags *bool `json:"DeleteAllTags,omitnil,omitempty" name:"DeleteAllTags"`

	// 修改的实例标签列表
	InstanceTags []*RabbitMQServerlessTag `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`
}

func (r *ModifyRabbitMQServerlessInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQServerlessInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClusterName")
	delete(f, "Remark")
	delete(f, "TraceFlag")
	delete(f, "SendReceiveRatio")
	delete(f, "DeleteAllTags")
	delete(f, "InstanceTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRabbitMQServerlessInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQServerlessInstanceResponseParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRabbitMQServerlessInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRabbitMQServerlessInstanceResponseParams `json:"Response"`
}

func (r *ModifyRabbitMQServerlessInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQServerlessInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQServerlessPermissionRequestParams struct {
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，权限关联的用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// vhost名
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 权限类型，declare相关操作，该用户可操作该vhost下的资源名称正则表达式
	ConfigRegexp *string `json:"ConfigRegexp,omitnil,omitempty" name:"ConfigRegexp"`

	// 权限类型，消息写入相关操作，该用户可操作该vhost下的资源名称正则表达式
	WriteRegexp *string `json:"WriteRegexp,omitnil,omitempty" name:"WriteRegexp"`

	// 权限类型，消息读取相关操作，该用户可操作该vhost下的资源名称正则表达式
	ReadRegexp *string `json:"ReadRegexp,omitnil,omitempty" name:"ReadRegexp"`
}

type ModifyRabbitMQServerlessPermissionRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，权限关联的用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// vhost名
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 权限类型，declare相关操作，该用户可操作该vhost下的资源名称正则表达式
	ConfigRegexp *string `json:"ConfigRegexp,omitnil,omitempty" name:"ConfigRegexp"`

	// 权限类型，消息写入相关操作，该用户可操作该vhost下的资源名称正则表达式
	WriteRegexp *string `json:"WriteRegexp,omitnil,omitempty" name:"WriteRegexp"`

	// 权限类型，消息读取相关操作，该用户可操作该vhost下的资源名称正则表达式
	ReadRegexp *string `json:"ReadRegexp,omitnil,omitempty" name:"ReadRegexp"`
}

func (r *ModifyRabbitMQServerlessPermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQServerlessPermissionRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRabbitMQServerlessPermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQServerlessPermissionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRabbitMQServerlessPermissionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRabbitMQServerlessPermissionResponseParams `json:"Response"`
}

func (r *ModifyRabbitMQServerlessPermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQServerlessPermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQServerlessQueueRequestParams struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 队列名称
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 新修改的备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// MessageTTL参数单位ms,classic类型专用	
	MessageTTL *int64 `json:"MessageTTL,omitnil,omitempty" name:"MessageTTL"`

	// DeadLetterExchange参数。可将过期或被拒绝的消息投往指定的死信 exchange。
	DeadLetterExchange *string `json:"DeadLetterExchange,omitnil,omitempty" name:"DeadLetterExchange"`

	// DeadLetterRoutingKey参数。只能包含字母、数字、"."、"-"，"@"，"_"
	DeadLetterRoutingKey *string `json:"DeadLetterRoutingKey,omitnil,omitempty" name:"DeadLetterRoutingKey"`
}

type ModifyRabbitMQServerlessQueueRequest struct {
	*tchttp.BaseRequest
	
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Vhost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// 队列名称
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 新修改的备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// MessageTTL参数单位ms,classic类型专用	
	MessageTTL *int64 `json:"MessageTTL,omitnil,omitempty" name:"MessageTTL"`

	// DeadLetterExchange参数。可将过期或被拒绝的消息投往指定的死信 exchange。
	DeadLetterExchange *string `json:"DeadLetterExchange,omitnil,omitempty" name:"DeadLetterExchange"`

	// DeadLetterRoutingKey参数。只能包含字母、数字、"."、"-"，"@"，"_"
	DeadLetterRoutingKey *string `json:"DeadLetterRoutingKey,omitnil,omitempty" name:"DeadLetterRoutingKey"`
}

func (r *ModifyRabbitMQServerlessQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQServerlessQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "QueueName")
	delete(f, "Remark")
	delete(f, "MessageTTL")
	delete(f, "DeadLetterExchange")
	delete(f, "DeadLetterRoutingKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRabbitMQServerlessQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQServerlessQueueResponseParams struct {
	// 队列名称
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRabbitMQServerlessQueueResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRabbitMQServerlessQueueResponseParams `json:"Response"`
}

func (r *ModifyRabbitMQServerlessQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQServerlessQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQServerlessUserRequestParams struct {
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 描述，不传则不修改
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户标签，用于决定改用户访问RabbitMQ Management的权限范围，不传则不修改
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 该用户的最大连接数，不传则不修改
	MaxConnections *int64 `json:"MaxConnections,omitnil,omitempty" name:"MaxConnections"`

	// 该用户的最大channel数，不传则不修改
	MaxChannels *int64 `json:"MaxChannels,omitnil,omitempty" name:"MaxChannels"`
}

type ModifyRabbitMQServerlessUserRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 描述，不传则不修改
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用户标签，用于决定改用户访问RabbitMQ Management的权限范围，不传则不修改
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 该用户的最大连接数，不传则不修改
	MaxConnections *int64 `json:"MaxConnections,omitnil,omitempty" name:"MaxConnections"`

	// 该用户的最大channel数，不传则不修改
	MaxChannels *int64 `json:"MaxChannels,omitnil,omitempty" name:"MaxChannels"`
}

func (r *ModifyRabbitMQServerlessUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQServerlessUserRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRabbitMQServerlessUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQServerlessUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRabbitMQServerlessUserResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRabbitMQServerlessUserResponseParams `json:"Response"`
}

func (r *ModifyRabbitMQServerlessUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQServerlessUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQServerlessVirtualHostRequestParams struct {
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost名
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// vhost描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 消息轨迹开关,true打开,false关闭
	TraceFlag *bool `json:"TraceFlag,omitnil,omitempty" name:"TraceFlag"`
}

type ModifyRabbitMQServerlessVirtualHostRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// vhost名
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// vhost描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 消息轨迹开关,true打开,false关闭
	TraceFlag *bool `json:"TraceFlag,omitnil,omitempty" name:"TraceFlag"`
}

func (r *ModifyRabbitMQServerlessVirtualHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQServerlessVirtualHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "Description")
	delete(f, "TraceFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRabbitMQServerlessVirtualHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQServerlessVirtualHostResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRabbitMQServerlessVirtualHostResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRabbitMQServerlessVirtualHostResponseParams `json:"Response"`
}

func (r *ModifyRabbitMQServerlessVirtualHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQServerlessVirtualHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	// Vhost参数
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
}

type RabbitMQClusterInfo struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 地域信息
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 创建时间，毫秒为单位
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

	// 过期时间
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

	// 集群类型
	InstanceType *uint64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 消息保留时间，单位小时
	MessageRetainTime *uint64 `json:"MessageRetainTime,omitnil,omitempty" name:"MessageRetainTime"`

	// 发送消息流量比例
	SendReceiveRatio *float64 `json:"SendReceiveRatio,omitnil,omitempty" name:"SendReceiveRatio"`

	// 消息轨迹保留时间，单位小时
	TraceTime *uint64 `json:"TraceTime,omitnil,omitempty" name:"TraceTime"`

	// 实例标签列表
	Tags []*RabbitMQServerlessTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type RabbitMQClusterSpecInfo struct {
	// 集群规格名称
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 峰值tps
	MaxTps *uint64 `json:"MaxTps,omitnil,omitempty" name:"MaxTps"`

	// 最大队列数
	MaxQueueNum *uint64 `json:"MaxQueueNum,omitnil,omitempty" name:"MaxQueueNum"`

	// 最大交换机数
	MaxExchangeNum *uint64 `json:"MaxExchangeNum,omitnil,omitempty" name:"MaxExchangeNum"`

	// 最大vhost数
	MaxVhostNum *uint64 `json:"MaxVhostNum,omitnil,omitempty" name:"MaxVhostNum"`

	// 最大连接数
	MaxConnNum *uint64 `json:"MaxConnNum,omitnil,omitempty" name:"MaxConnNum"`

	// 最大用户数
	MaxUserNum *uint64 `json:"MaxUserNum,omitnil,omitempty" name:"MaxUserNum"`

	// 峰值带宽，已废弃
	MaxBandWidth *uint64 `json:"MaxBandWidth,omitnil,omitempty" name:"MaxBandWidth"`

	// 公网带宽，已废弃
	PublicNetworkTps *uint64 `json:"PublicNetworkTps,omitnil,omitempty" name:"PublicNetworkTps"`

	// 实例对应的功能列表，true表示支持，false 表示不支持
	Features *string `json:"Features,omitnil,omitempty" name:"Features"`
}

type RabbitMQConnection struct {
	// 连接名称
	ConnectionName *string `json:"ConnectionName,omitnil,omitempty" name:"ConnectionName"`

	// 客户端ip
	PeerHost *string `json:"PeerHost,omitnil,omitempty" name:"PeerHost"`

	// 连接状态，包括 starting、tuning、opening、running、flow、blocking、blocked、closing 和 closed
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 连接使用用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 是否开启ssl
	SSL *bool `json:"SSL,omitnil,omitempty" name:"SSL"`

	// 连接协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 连接下的channel数
	Channels *int64 `json:"Channels,omitnil,omitempty" name:"Channels"`

	// 入流量大小，单位 bytes
	IncomingBytes *float64 `json:"IncomingBytes,omitnil,omitempty" name:"IncomingBytes"`

	// 出流量大小，单位bytes
	OutgoingBytes *float64 `json:"OutgoingBytes,omitnil,omitempty" name:"OutgoingBytes"`

	// 心跳间隔时间，默认60s
	Heartbeat *uint64 `json:"Heartbeat,omitnil,omitempty" name:"Heartbeat"`

	// 一个链接最大的channel数，默认1024
	MaxChannel *uint64 `json:"MaxChannel,omitnil,omitempty" name:"MaxChannel"`

	// 空闲时间点
	IdleSince *string `json:"IdleSince,omitnil,omitempty" name:"IdleSince"`
}

type RabbitMQConsumersListInfo struct {
	// 客户端Ip
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// 消费者Tag
	ConsumerTag *string `json:"ConsumerTag,omitnil,omitempty" name:"ConsumerTag"`

	// 消费目标队列
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 是否需要消费者手动 ack
	AckRequired *bool `json:"AckRequired,omitnil,omitempty" name:"AckRequired"`

	// 消费者 qos 值
	PrefetchCount *uint64 `json:"PrefetchCount,omitnil,omitempty" name:"PrefetchCount"`

	// 消费者状态
	Active *string `json:"Active,omitnil,omitempty" name:"Active"`

	// 最后一次投递消息时间
	LastDeliveredTime *string `json:"LastDeliveredTime,omitnil,omitempty" name:"LastDeliveredTime"`

	// 消费者未确认消息数
	UnAckMsgCount *int64 `json:"UnAckMsgCount,omitnil,omitempty" name:"UnAckMsgCount"`

	// consumer 所属的 channel 
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`
}

type RabbitMQExchangeListInfo struct {
	// exchange 名
	ExchangeName *string `json:"ExchangeName,omitnil,omitempty" name:"ExchangeName"`

	// 备注说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// exchange 类型, 支持 "fanout","direct","topic","headers"
	ExchangeType *string `json:"ExchangeType,omitnil,omitempty" name:"ExchangeType"`

	// VHost参数
	VirtualHost *string `json:"VirtualHost,omitnil,omitempty" name:"VirtualHost"`

	// exchange 创建者, "system":"系统创建", "user":"用户创建"
	ExchangeCreator *string `json:"ExchangeCreator,omitnil,omitempty" name:"ExchangeCreator"`

	// exchange 创建时间
	CreateTimeStamp *string `json:"CreateTimeStamp,omitnil,omitempty" name:"CreateTimeStamp"`

	// exchange 修改时间
	ModTimeStamp *string `json:"ModTimeStamp,omitnil,omitempty" name:"ModTimeStamp"`

	// 输入消息速率
	MessageRateIn *float64 `json:"MessageRateIn,omitnil,omitempty" name:"MessageRateIn"`

	// 输出消息速率
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
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 扩展参数 key-value 对象
	Arguments *string `json:"Arguments,omitnil,omitempty" name:"Arguments"`

	// 未调度的延时消息数量
	MessagesDelayed *uint64 `json:"MessagesDelayed,omitnil,omitempty" name:"MessagesDelayed"`
}

type RabbitMQPermission struct {
	// 集群实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名，权限关联的用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// vhost名
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
}

type RabbitMQQueueListConsumerDetailInfo struct {
	// 消费者数量
	ConsumersNumber *int64 `json:"ConsumersNumber,omitnil,omitempty" name:"ConsumersNumber"`
}

type RabbitMQQueueListInfo struct {
	// 队列名
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 备注说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 消费者信息
	ConsumerDetail *RabbitMQQueueListConsumerDetailInfo `json:"ConsumerDetail,omitnil,omitempty" name:"ConsumerDetail"`

	// 队列类型，取值 "classic"，"quorum"
	QueueType *string `json:"QueueType,omitnil,omitempty" name:"QueueType"`

	// 消息堆积数
	MessageHeapCount *int64 `json:"MessageHeapCount,omitnil,omitempty" name:"MessageHeapCount"`

	// 消息生产速率，每秒
	MessageRateIn *float64 `json:"MessageRateIn,omitnil,omitempty" name:"MessageRateIn"`

	// 消息消费速率，每秒
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
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 扩展参数 key-value 对象
	Arguments *string `json:"Arguments,omitnil,omitempty" name:"Arguments"`

	// 是否独占队列
	Exclusive *bool `json:"Exclusive,omitnil,omitempty" name:"Exclusive"`
}

type RabbitMQServerlessAccessInfo struct {
	// 公网域名
	PublicAccessEndpoint *string `json:"PublicAccessEndpoint,omitnil,omitempty" name:"PublicAccessEndpoint"`

	// 公网状态
	PublicDataStreamStatus *string `json:"PublicDataStreamStatus,omitnil,omitempty" name:"PublicDataStreamStatus"`

	// 公网CLB实例ID
	PublicClbId *string `json:"PublicClbId,omitnil,omitempty" name:"PublicClbId"`
}

type RabbitMQServerlessEndpoint struct {
	// vpc id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// subnet id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 接入地址
	VpcEndpoint *string `json:"VpcEndpoint,omitnil,omitempty" name:"VpcEndpoint"`

	// 接入地址状态
	VpcDataStreamEndpointStatus *string `json:"VpcDataStreamEndpointStatus,omitnil,omitempty" name:"VpcDataStreamEndpointStatus"`

	// 是否是公网
	PublicNetwork *bool `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`

	// 访问策略
	AccessStrategy *string `json:"AccessStrategy,omitnil,omitempty" name:"AccessStrategy"`

	// 带宽
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`
}

type RabbitMQServerlessInstance struct {
	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例版本号
	InstanceVersion *string `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`

	// 实例状态，0表示创建中，1表示正常，2表示隔离中，3表示已销毁，4 - 异常, 5 - 发货失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// TPS
	MaxTps *int64 `json:"MaxTps,omitnil,omitempty" name:"MaxTps"`

	// 带宽
	MaxBandWidth *int64 `json:"MaxBandWidth,omitnil,omitempty" name:"MaxBandWidth"`

	// 集群过期时间
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 自动续费标记，0表示默认状态(用户未设置，即初始状态即手动续费)， 1表示自动续费，2表示明确不自动续费(用户设置)
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 0-后付费，1-预付费
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 集群规格
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 异常信息
	ExceptionInformation *string `json:"ExceptionInformation,omitnil,omitempty" name:"ExceptionInformation"`

	// 公网接入点
	PublicAccessEndpoint *string `json:"PublicAccessEndpoint,omitnil,omitempty" name:"PublicAccessEndpoint"`

	// 私有网络接入点
	Vpcs []*RabbitMQServerlessEndpoint `json:"Vpcs,omitnil,omitempty" name:"Vpcs"`

	// 实例状态，0表示创建中，1表示正常，2表示隔离中，3表示已销毁，4 - 异常, 5 - 发货失败
	ClusterStatus *int64 `json:"ClusterStatus,omitnil,omitempty" name:"ClusterStatus"`

	// 集群类型：1
	InstanceType *uint64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 过期时间
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 为了兼容托管版，固定值 0
	NodeCount *int64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 为了兼容托管版，固定值 0
	MaxStorage *int64 `json:"MaxStorage,omitnil,omitempty" name:"MaxStorage"`

	// 隔离时间
	IsolatedTime *uint64 `json:"IsolatedTime,omitnil,omitempty" name:"IsolatedTime"`

	// Serverless 扩展字段
	ServerlessExt *string `json:"ServerlessExt,omitnil,omitempty" name:"ServerlessExt"`

	// 实例标签列表
	Tags []*RabbitMQServerlessTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type RabbitMQServerlessTag struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type RabbitMQServerlessWhiteListInfo struct {
	// 公网数据流白名单
	PublicDataStreamWhiteList *string `json:"PublicDataStreamWhiteList,omitnil,omitempty" name:"PublicDataStreamWhiteList"`

	// 公网数据流白名单状态
	PublicDataStreamWhiteListStatus *string `json:"PublicDataStreamWhiteListStatus,omitnil,omitempty" name:"PublicDataStreamWhiteListStatus"`
}

type RabbitMQUser struct {
	// 集群实例Id
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

	// 该用户所能允许的最大连接数
	MaxConnections *int64 `json:"MaxConnections,omitnil,omitempty" name:"MaxConnections"`

	// 该用户所能允许的最大通道数
	MaxChannels *int64 `json:"MaxChannels,omitnil,omitempty" name:"MaxChannels"`
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

type UserQuota struct {
	// 最大用户数
	MaxUser *int64 `json:"MaxUser,omitnil,omitempty" name:"MaxUser"`

	// 已用用户数
	UsedUser *int64 `json:"UsedUser,omitnil,omitempty" name:"UsedUser"`
}

type VirtualHostQuota struct {
	// 最大虚拟主机数
	MaxVirtualHost *uint64 `json:"MaxVirtualHost,omitnil,omitempty" name:"MaxVirtualHost"`

	// 已经使用的虚拟主机数
	UsedVirtualHost *uint64 `json:"UsedVirtualHost,omitnil,omitempty" name:"UsedVirtualHost"`
}

type VpcEndpointInfo struct {
	// vpc的id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// vpc接入点信息
	VpcEndpoint *string `json:"VpcEndpoint,omitnil,omitempty" name:"VpcEndpoint"`

	// vpc接入点状态
	// OFF/ON/CREATING/DELETING
	VpcDataStreamEndpointStatus *string `json:"VpcDataStreamEndpointStatus,omitnil,omitempty" name:"VpcDataStreamEndpointStatus"`
}