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

package v20230308

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type ChangeMigratingTopicToNextStageRequestParams struct {
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 主题名称列表，主题名称可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	TopicNameList []*string `json:"TopicNameList,omitnil,omitempty" name:"TopicNameList"`

	// 命名空间列表，仅4.x集群有效，与TopicNameList一一对应，可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	NamespaceList []*string `json:"NamespaceList,omitnil,omitempty" name:"NamespaceList"`
}

type ChangeMigratingTopicToNextStageRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 主题名称列表，主题名称可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	TopicNameList []*string `json:"TopicNameList,omitnil,omitempty" name:"TopicNameList"`

	// 命名空间列表，仅4.x集群有效，与TopicNameList一一对应，可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	NamespaceList []*string `json:"NamespaceList,omitnil,omitempty" name:"NamespaceList"`
}

func (r *ChangeMigratingTopicToNextStageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeMigratingTopicToNextStageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "TopicNameList")
	delete(f, "NamespaceList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChangeMigratingTopicToNextStageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChangeMigratingTopicToNextStageResponseParams struct {
	// 迁移主题状态修改的结果列表
	Results []*TopicStageChangeResult `json:"Results,omitnil,omitempty" name:"Results"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChangeMigratingTopicToNextStageResponse struct {
	*tchttp.BaseResponse
	Response *ChangeMigratingTopicToNextStageResponseParams `json:"Response"`
}

func (r *ChangeMigratingTopicToNextStageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChangeMigratingTopicToNextStageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClientSubscriptionInfo struct {
	// 客户端ID
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 客户端地址
	ClientAddr *string `json:"ClientAddr,omitnil,omitempty" name:"ClientAddr"`

	// 订阅主题
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 订阅表达式
	SubString *string `json:"SubString,omitnil,omitempty" name:"SubString"`

	// 订阅方式
	ExpressionType *string `json:"ExpressionType,omitnil,omitempty" name:"ExpressionType"`
}

type ConsumeGroupItem struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费组名称
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`

	// 是否开启消费
	ConsumeEnable *bool `json:"ConsumeEnable,omitnil,omitempty" name:"ConsumeEnable"`

	// 顺序投递：true
	// 并发投递：false
	ConsumeMessageOrderly *bool `json:"ConsumeMessageOrderly,omitnil,omitempty" name:"ConsumeMessageOrderly"`

	// 最大重试次数
	MaxRetryTimes *int64 `json:"MaxRetryTimes,omitnil,omitempty" name:"MaxRetryTimes"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 4.x的集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterIdV4 *string `json:"ClusterIdV4,omitnil,omitempty" name:"ClusterIdV4"`

	// 4.x的命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceV4 *string `json:"NamespaceV4,omitnil,omitempty" name:"NamespaceV4"`

	// 4.x的消费组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerGroupV4 *string `json:"ConsumerGroupV4,omitnil,omitempty" name:"ConsumerGroupV4"`

	// 4.x的完整命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullNamespaceV4 *string `json:"FullNamespaceV4,omitnil,omitempty" name:"FullNamespaceV4"`
}

type ConsumerClient struct {
	// 客户端ID
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 客户端地址
	ClientAddr *string `json:"ClientAddr,omitnil,omitempty" name:"ClientAddr"`

	// 客户端SDK语言
	Language *string `json:"Language,omitnil,omitempty" name:"Language"`

	// 客户端SDK版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 客户端消费堆积
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerLag *int64 `json:"ConsumerLag,omitnil,omitempty" name:"ConsumerLag"`

	// 消费者客户端类型，枚举值如下：
	// 
	// - grpc：GRPC协议
	// - remoting：Remoting协议
	// - http：HTTP协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelProtocol *string `json:"ChannelProtocol,omitnil,omitempty" name:"ChannelProtocol"`
}

// Predefined struct for user
type CreateConsumerGroupRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 最大重试次数，取值范围0～1000
	MaxRetryTimes *int64 `json:"MaxRetryTimes,omitnil,omitempty" name:"MaxRetryTimes"`

	// 是否开启消费
	ConsumeEnable *bool `json:"ConsumeEnable,omitnil,omitempty" name:"ConsumeEnable"`

	// 顺序投递：true
	// 并发投递：false
	ConsumeMessageOrderly *bool `json:"ConsumeMessageOrderly,omitnil,omitempty" name:"ConsumeMessageOrderly"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`

	// 备注信息，最多 128 个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type CreateConsumerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 最大重试次数，取值范围0～1000
	MaxRetryTimes *int64 `json:"MaxRetryTimes,omitnil,omitempty" name:"MaxRetryTimes"`

	// 是否开启消费
	ConsumeEnable *bool `json:"ConsumeEnable,omitnil,omitempty" name:"ConsumeEnable"`

	// 顺序投递：true
	// 并发投递：false
	ConsumeMessageOrderly *bool `json:"ConsumeMessageOrderly,omitnil,omitempty" name:"ConsumeMessageOrderly"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`

	// 备注信息，最多 128 个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`
}

func (r *CreateConsumerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConsumerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "MaxRetryTimes")
	delete(f, "ConsumeEnable")
	delete(f, "ConsumeMessageOrderly")
	delete(f, "ConsumerGroup")
	delete(f, "Remark")
	delete(f, "TagList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConsumerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConsumerGroupResponseParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费组名称
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConsumerGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateConsumerGroupResponseParams `json:"Response"`
}

func (r *CreateConsumerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceRequestParams struct {
	// 实例类型，枚举值如下：
	// 
	// - EXPERIMENT：体验版
	// 
	// - BASIC：基础版
	// 
	// - PRO：专业版
	// 
	// - PLATINUM：铂金版
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 集群名称，不能为空， 3-64个字符，只能包含数字、字母、“-”和“_”
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 商品规格，从 [DescribeProductSKUs](https://cloud.tencent.com/document/api/1493/107676) 接口中的 [ProductSKU](https://cloud.tencent.com/document/api/1493/96031#ProductSKU) 出参获得。
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 集群绑定的VPC信息，必填
	VpcList []*VpcInfo `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// 是否开启公网，默认值为false表示不开启
	EnablePublic *bool `json:"EnablePublic,omitnil,omitempty" name:"EnablePublic"`

	// 公网是否按流量计费，默认值为false表示不按流量计费
	BillingFlow *bool `json:"BillingFlow,omitnil,omitempty" name:"BillingFlow"`

	// 公网带宽（单位：兆），默认值为0。如果开启公网，该字段必须为大于0的正整数
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网访问白名单，不填表示拒绝所有 IP 访问
	IpRules []*IpRule `json:"IpRules,omitnil,omitempty" name:"IpRules"`

	// 消息保留时长（单位：小时），取值范围参考 [DescribeProductSKUs](https://cloud.tencent.com/document/api/1493/107676) 接口中的 [ProductSKU](https://cloud.tencent.com/document/api/1493/96031#ProductSKU) 出参：
	// 
	// - 默认值：DefaultRetention 参数
	// - 最小值：RetentionLowerLimit 参数
	// - 最大值：RetentionUpperLimit 参数
	MessageRetention *int64 `json:"MessageRetention,omitnil,omitempty" name:"MessageRetention"`

	// 付费模式（0: 后付费；1: 预付费），默认值为0
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 预付费集群是否自动续费（0: 不自动续费；1: 自动续费），默认值为0
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 预付费集群的购买时长（单位：月），取值范围为1～60，默认值为1
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 最大可创建主题数，从 [DescribeProductSKUs](https://cloud.tencent.com/document/api/1493/107676) 接口中的 [ProductSKU](https://cloud.tencent.com/document/api/1493/96031#ProductSKU) 出参：
	// 
	// - 默认值和最小值：TopicNumLimit 参数
	// - 最大值：TopicNumUpperLimit 参数
	MaxTopicNum *int64 `json:"MaxTopicNum,omitnil,omitempty" name:"MaxTopicNum"`

	// 部署可用区列表，从 [DescribeZones](https://cloud.tencent.com/document/product/1596/77929) 接口返回中的 [ZoneInfo](https://cloud.tencent.com/document/api/1596/77932#ZoneInfo) 数据结构中获得。
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例类型，枚举值如下：
	// 
	// - EXPERIMENT：体验版
	// 
	// - BASIC：基础版
	// 
	// - PRO：专业版
	// 
	// - PLATINUM：铂金版
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 集群名称，不能为空， 3-64个字符，只能包含数字、字母、“-”和“_”
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 商品规格，从 [DescribeProductSKUs](https://cloud.tencent.com/document/api/1493/107676) 接口中的 [ProductSKU](https://cloud.tencent.com/document/api/1493/96031#ProductSKU) 出参获得。
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 集群绑定的VPC信息，必填
	VpcList []*VpcInfo `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// 是否开启公网，默认值为false表示不开启
	EnablePublic *bool `json:"EnablePublic,omitnil,omitempty" name:"EnablePublic"`

	// 公网是否按流量计费，默认值为false表示不按流量计费
	BillingFlow *bool `json:"BillingFlow,omitnil,omitempty" name:"BillingFlow"`

	// 公网带宽（单位：兆），默认值为0。如果开启公网，该字段必须为大于0的正整数
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网访问白名单，不填表示拒绝所有 IP 访问
	IpRules []*IpRule `json:"IpRules,omitnil,omitempty" name:"IpRules"`

	// 消息保留时长（单位：小时），取值范围参考 [DescribeProductSKUs](https://cloud.tencent.com/document/api/1493/107676) 接口中的 [ProductSKU](https://cloud.tencent.com/document/api/1493/96031#ProductSKU) 出参：
	// 
	// - 默认值：DefaultRetention 参数
	// - 最小值：RetentionLowerLimit 参数
	// - 最大值：RetentionUpperLimit 参数
	MessageRetention *int64 `json:"MessageRetention,omitnil,omitempty" name:"MessageRetention"`

	// 付费模式（0: 后付费；1: 预付费），默认值为0
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 预付费集群是否自动续费（0: 不自动续费；1: 自动续费），默认值为0
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 预付费集群的购买时长（单位：月），取值范围为1～60，默认值为1
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 最大可创建主题数，从 [DescribeProductSKUs](https://cloud.tencent.com/document/api/1493/107676) 接口中的 [ProductSKU](https://cloud.tencent.com/document/api/1493/96031#ProductSKU) 出参：
	// 
	// - 默认值和最小值：TopicNumLimit 参数
	// - 最大值：TopicNumUpperLimit 参数
	MaxTopicNum *int64 `json:"MaxTopicNum,omitnil,omitempty" name:"MaxTopicNum"`

	// 部署可用区列表，从 [DescribeZones](https://cloud.tencent.com/document/product/1596/77929) 接口返回中的 [ZoneInfo](https://cloud.tencent.com/document/api/1596/77932#ZoneInfo) 数据结构中获得。
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceType")
	delete(f, "Name")
	delete(f, "SkuCode")
	delete(f, "Remark")
	delete(f, "TagList")
	delete(f, "VpcList")
	delete(f, "EnablePublic")
	delete(f, "BillingFlow")
	delete(f, "Bandwidth")
	delete(f, "IpRules")
	delete(f, "MessageRetention")
	delete(f, "PayMode")
	delete(f, "RenewFlag")
	delete(f, "TimeSpan")
	delete(f, "MaxTopicNum")
	delete(f, "ZoneIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceResponseParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceResponseParams `json:"Response"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMQTTInsPublicEndpointRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 带宽
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网访问规则
	Rules []*PublicAccessRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type CreateMQTTInsPublicEndpointRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 带宽
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网访问规则
	Rules []*PublicAccessRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *CreateMQTTInsPublicEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMQTTInsPublicEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Bandwidth")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMQTTInsPublicEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMQTTInsPublicEndpointResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMQTTInsPublicEndpointResponse struct {
	*tchttp.BaseResponse
	Response *CreateMQTTInsPublicEndpointResponseParams `json:"Response"`
}

func (r *CreateMQTTInsPublicEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMQTTInsPublicEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMQTTInstanceRequestParams struct {
	// 实例类型，
	// EXPERIMENT 体验版
	// BASIC 基础版
	// PRO  专业版
	// PLATINUM 铂金版
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 商品规格，可用规格如下：
	// basic_1k,
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 实例绑定的VPC信息
	VpcList []*VpcInfo `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// 是否开启公网
	EnablePublic *bool `json:"EnablePublic,omitnil,omitempty" name:"EnablePublic"`

	// 公网带宽（单位：兆）
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网访问白名单
	IpRules []*IpRule `json:"IpRules,omitnil,omitempty" name:"IpRules"`

	// 是否自动续费（0: 不自动续费；1: 自动续费）
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 购买时长（单位：月）
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`
}

type CreateMQTTInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例类型，
	// EXPERIMENT 体验版
	// BASIC 基础版
	// PRO  专业版
	// PLATINUM 铂金版
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 商品规格，可用规格如下：
	// basic_1k,
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 实例绑定的VPC信息
	VpcList []*VpcInfo `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// 是否开启公网
	EnablePublic *bool `json:"EnablePublic,omitnil,omitempty" name:"EnablePublic"`

	// 公网带宽（单位：兆）
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网访问白名单
	IpRules []*IpRule `json:"IpRules,omitnil,omitempty" name:"IpRules"`

	// 是否自动续费（0: 不自动续费；1: 自动续费）
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 购买时长（单位：月）
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`
}

func (r *CreateMQTTInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMQTTInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceType")
	delete(f, "Name")
	delete(f, "SkuCode")
	delete(f, "Remark")
	delete(f, "TagList")
	delete(f, "VpcList")
	delete(f, "EnablePublic")
	delete(f, "Bandwidth")
	delete(f, "IpRules")
	delete(f, "RenewFlag")
	delete(f, "TimeSpan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMQTTInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMQTTInstanceResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMQTTInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateMQTTInstanceResponseParams `json:"Response"`
}

func (r *CreateMQTTInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMQTTInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMQTTTopicRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateMQTTTopicRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreateMQTTTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMQTTTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMQTTTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMQTTTopicResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMQTTTopicResponse struct {
	*tchttp.BaseResponse
	Response *CreateMQTTTopicResponseParams `json:"Response"`
}

func (r *CreateMQTTTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMQTTTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMQTTUserRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否开启生产权限
	PermWrite *bool `json:"PermWrite,omitnil,omitempty" name:"PermWrite"`

	// 是否开启消费权限
	PermRead *bool `json:"PermRead,omitnil,omitempty" name:"PermRead"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 密码，该字段为空时候则后端会默认生成
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type CreateMQTTUserRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否开启生产权限
	PermWrite *bool `json:"PermWrite,omitnil,omitempty" name:"PermWrite"`

	// 是否开启消费权限
	PermRead *bool `json:"PermRead,omitnil,omitempty" name:"PermRead"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 密码，该字段为空时候则后端会默认生成
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

func (r *CreateMQTTUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMQTTUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Remark")
	delete(f, "PermWrite")
	delete(f, "PermRead")
	delete(f, "Username")
	delete(f, "Password")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMQTTUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMQTTUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMQTTUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateMQTTUserResponseParams `json:"Response"`
}

func (r *CreateMQTTUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMQTTUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoleRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 角色名称，不能为空，只支持数字 大小写字母 分隔符("_","-")，不能超过 32 个字符
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 是否开启生产权限
	PermWrite *bool `json:"PermWrite,omitnil,omitempty" name:"PermWrite"`

	// 是否开启消费权限
	PermRead *bool `json:"PermRead,omitnil,omitempty" name:"PermRead"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 权限类型，默认按集群授权（Cluster：集群级别；TopicAndGroup：主题&消费组级别）
	PermType *string `json:"PermType,omitnil,omitempty" name:"PermType"`

	// Topic&Group维度权限配置，权限类型为 TopicAndGroup 时必填
	DetailedPerms []*DetailedRolePerm `json:"DetailedPerms,omitnil,omitempty" name:"DetailedPerms"`
}

type CreateRoleRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 角色名称，不能为空，只支持数字 大小写字母 分隔符("_","-")，不能超过 32 个字符
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 是否开启生产权限
	PermWrite *bool `json:"PermWrite,omitnil,omitempty" name:"PermWrite"`

	// 是否开启消费权限
	PermRead *bool `json:"PermRead,omitnil,omitempty" name:"PermRead"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 权限类型，默认按集群授权（Cluster：集群级别；TopicAndGroup：主题&消费组级别）
	PermType *string `json:"PermType,omitnil,omitempty" name:"PermType"`

	// Topic&Group维度权限配置，权限类型为 TopicAndGroup 时必填
	DetailedPerms []*DetailedRolePerm `json:"DetailedPerms,omitnil,omitempty" name:"DetailedPerms"`
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
	delete(f, "InstanceId")
	delete(f, "Role")
	delete(f, "PermWrite")
	delete(f, "PermRead")
	delete(f, "Remark")
	delete(f, "PermType")
	delete(f, "DetailedPerms")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRoleResponseParams struct {
	// 角色名
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

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
type CreateTopicRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，从 [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030) 接口返回的 [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem) 或控制台获得。
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 主题类型，枚举值如下：
	// 
	// - NORMAL: 普通消息
	// - FIFO: 顺序消息
	// - DELAY: 延时消息
	// - TRANSACTION: 事务消息
	TopicType *string `json:"TopicType,omitnil,omitempty" name:"TopicType"`

	// 队列数量，取值范围3～16
	QueueNum *int64 `json:"QueueNum,omitnil,omitempty" name:"QueueNum"`

	// 备注信息，最多 128 个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 消息保留时长（单位：小时）
	MsgTTL *int64 `json:"MsgTTL,omitnil,omitempty" name:"MsgTTL"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，从 [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030) 接口返回的 [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem) 或控制台获得。
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 主题类型，枚举值如下：
	// 
	// - NORMAL: 普通消息
	// - FIFO: 顺序消息
	// - DELAY: 延时消息
	// - TRANSACTION: 事务消息
	TopicType *string `json:"TopicType,omitnil,omitempty" name:"TopicType"`

	// 队列数量，取值范围3～16
	QueueNum *int64 `json:"QueueNum,omitnil,omitempty" name:"QueueNum"`

	// 备注信息，最多 128 个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 消息保留时长（单位：小时）
	MsgTTL *int64 `json:"MsgTTL,omitnil,omitempty" name:"MsgTTL"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`
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
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "TopicType")
	delete(f, "QueueNum")
	delete(f, "Remark")
	delete(f, "MsgTTL")
	delete(f, "TagList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicResponseParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

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

type CustomMapEntry struct {
	// key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type DeleteConsumerGroupRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`
}

type DeleteConsumerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`
}

func (r *DeleteConsumerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConsumerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConsumerGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConsumerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConsumerGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteConsumerGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConsumerGroupResponseParams `json:"Response"`
}

func (r *DeleteConsumerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceResponseParams `json:"Response"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMQTTInsPublicEndpointRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteMQTTInsPublicEndpointRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteMQTTInsPublicEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMQTTInsPublicEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMQTTInsPublicEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMQTTInsPublicEndpointResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMQTTInsPublicEndpointResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMQTTInsPublicEndpointResponseParams `json:"Response"`
}

func (r *DeleteMQTTInsPublicEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMQTTInsPublicEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMQTTInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteMQTTInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteMQTTInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMQTTInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMQTTInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMQTTInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMQTTInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMQTTInstanceResponseParams `json:"Response"`
}

func (r *DeleteMQTTInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMQTTInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMQTTTopicRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
}

type DeleteMQTTTopicRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
}

func (r *DeleteMQTTTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMQTTTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMQTTTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMQTTTopicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMQTTTopicResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMQTTTopicResponseParams `json:"Response"`
}

func (r *DeleteMQTTTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMQTTTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMQTTUserRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`
}

type DeleteMQTTUserRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`
}

func (r *DeleteMQTTUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMQTTUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Username")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMQTTUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMQTTUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMQTTUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMQTTUserResponseParams `json:"Response"`
}

func (r *DeleteMQTTUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMQTTUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoleRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 角色名称，从 [DescribeRoleList](https://cloud.tencent.com/document/api/1493/98862) 接口响应中的 [RoleItem](https://cloud.tencent.com/document/api/1493/96031#RoleItem) 或控制台获得。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`
}

type DeleteRoleRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 角色名称，从 [DescribeRoleList](https://cloud.tencent.com/document/api/1493/98862) 接口响应中的 [RoleItem](https://cloud.tencent.com/document/api/1493/96031#RoleItem) 或控制台获得。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`
}

func (r *DeleteRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Role")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRoleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRoleResponseParams `json:"Response"`
}

func (r *DeleteRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSmoothMigrationTaskRequestParams struct {
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteSmoothMigrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteSmoothMigrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSmoothMigrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSmoothMigrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSmoothMigrationTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSmoothMigrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSmoothMigrationTaskResponseParams `json:"Response"`
}

func (r *DeleteSmoothMigrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSmoothMigrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，从 [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030) 接口返回的 [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem) 或控制台获得。
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
}

type DeleteTopicRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，从 [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030) 接口返回的 [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem) 或控制台获得。
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
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
	delete(f, "InstanceId")
	delete(f, "Topic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeConsumerClientListRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeConsumerClientListRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeConsumerClientListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerClientListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConsumerGroup")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConsumerClientListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumerClientListResponseParams struct {
	// 查询总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 消费客户端
	Data []*ConsumerClient `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConsumerClientListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConsumerClientListResponseParams `json:"Response"`
}

func (r *DescribeConsumerClientListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerClientListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumerClientRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 客户端ID，从 [DescribeConsumerClientList](https://cloud.tencent.com/document/api/1493/120140) 接口中的 [ConsumerClient](https://cloud.tencent.com/document/api/1493/96031#ConsumerClient) 出参中获得。
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`
}

type DescribeConsumerClientRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 客户端ID，从 [DescribeConsumerClientList](https://cloud.tencent.com/document/api/1493/120140) 接口中的 [ConsumerClient](https://cloud.tencent.com/document/api/1493/96031#ConsumerClient) 出参中获得。
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`
}

func (r *DescribeConsumerClientRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerClientRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClientId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ConsumerGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConsumerClientRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumerClientResponseParams struct {
	// 客户端详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Client *ConsumerClient `json:"Client,omitnil,omitempty" name:"Client"`

	// 主题消费信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicList []*TopicConsumeStats `json:"TopicList,omitnil,omitempty" name:"TopicList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConsumerClientResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConsumerClientResponseParams `json:"Response"`
}

func (r *DescribeConsumerClientResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerClientResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumerGroupListRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询指定主题下的消费组
	FromTopic *string `json:"FromTopic,omitnil,omitempty" name:"FromTopic"`
}

type DescribeConsumerGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询指定主题下的消费组
	FromTopic *string `json:"FromTopic,omitnil,omitempty" name:"FromTopic"`
}

func (r *DescribeConsumerGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FromTopic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConsumerGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumerGroupListResponseParams struct {
	// 查询总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 消费组列表
	Data []*ConsumeGroupItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConsumerGroupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConsumerGroupListResponseParams `json:"Response"`
}

func (r *DescribeConsumerGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumerGroupRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`
}

type DescribeConsumerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`
}

func (r *DescribeConsumerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConsumerGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConsumerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumerGroupResponseParams struct {
	// 在线消费者数量
	ConsumerNum *int64 `json:"ConsumerNum,omitnil,omitempty" name:"ConsumerNum"`

	// TPS
	Tps *int64 `json:"Tps,omitnil,omitempty" name:"Tps"`

	// 消息堆积数量
	ConsumerLag *int64 `json:"ConsumerLag,omitnil,omitempty" name:"ConsumerLag"`

	// 消费类型，枚举值如下：
	// 
	// - PULL：PULL 消费类型
	// - PUSH：PUSH 消费类型
	// - POP：POP 消费类型
	ConsumeType *string `json:"ConsumeType,omitnil,omitempty" name:"ConsumeType"`

	// 创建时间，**Unix时间戳（毫秒）**
	CreatedTime *int64 `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 顺序投递：true
	// 并发投递：false
	ConsumeMessageOrderly *bool `json:"ConsumeMessageOrderly,omitnil,omitempty" name:"ConsumeMessageOrderly"`

	// 是否开启消费
	ConsumeEnable *bool `json:"ConsumeEnable,omitnil,omitempty" name:"ConsumeEnable"`

	// 最大重试次数
	MaxRetryTimes *int64 `json:"MaxRetryTimes,omitnil,omitempty" name:"MaxRetryTimes"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 消费模式：
	// BROADCASTING 广播模式
	// CLUSTERING 集群模式
	MessageModel *string `json:"MessageModel,omitnil,omitempty" name:"MessageModel"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConsumerGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConsumerGroupResponseParams `json:"Response"`
}

func (r *DescribeConsumerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumerLagRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`

	// 命名空间，4.x集群必填，从 [DescribeRocketMQNamespaces](https://cloud.tencent.com/document/api/1179/63419) 接口返回的 [RocketMQNamespace](https://cloud.tencent.com/document/api/1179/46089#RocketMQNamespace) 或控制台获得。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 订阅主题，不为空则查询订阅了该主题的消费组的堆积，从 [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030) 接口返回的 [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem) 或控制台获得。
	SubscribeTopic *string `json:"SubscribeTopic,omitnil,omitempty" name:"SubscribeTopic"`
}

type DescribeConsumerLagRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`

	// 命名空间，4.x集群必填，从 [DescribeRocketMQNamespaces](https://cloud.tencent.com/document/api/1179/63419) 接口返回的 [RocketMQNamespace](https://cloud.tencent.com/document/api/1179/46089#RocketMQNamespace) 或控制台获得。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 订阅主题，不为空则查询订阅了该主题的消费组的堆积，从 [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030) 接口返回的 [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem) 或控制台获得。
	SubscribeTopic *string `json:"SubscribeTopic,omitnil,omitempty" name:"SubscribeTopic"`
}

func (r *DescribeConsumerLagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerLagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConsumerGroup")
	delete(f, "Namespace")
	delete(f, "SubscribeTopic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConsumerLagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumerLagResponseParams struct {
	// 堆积数
	ConsumerLag *int64 `json:"ConsumerLag,omitnil,omitempty" name:"ConsumerLag"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConsumerLagResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConsumerLagResponseParams `json:"Response"`
}

func (r *DescribeConsumerLagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumerLagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFusionInstanceListRequestParams struct {
	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤器
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

type DescribeFusionInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤器
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

func (r *DescribeFusionInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFusionInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFusionInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFusionInstanceListResponseParams struct {
	// 查询总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例列表
	Data []*FusionInstanceItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFusionInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFusionInstanceListResponseParams `json:"Response"`
}

func (r *DescribeFusionInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFusionInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceListRequestParams struct {
	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤器
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签过滤器
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "TagFilters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceListResponseParams struct {
	// 查询总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例列表
	Data []*InstanceItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceListResponseParams `json:"Response"`
}

func (r *DescribeInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceResponseParams struct {
	// 实例类型，
	// EXPERIMENT 体验版
	// BASIC 基础版
	// PRO  专业版
	// PLATINUM 铂金版
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 主题数量
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// 实例最大主题数量
	TopicNumLimit *int64 `json:"TopicNumLimit,omitnil,omitempty" name:"TopicNumLimit"`

	// 消费组数量
	GroupNum *int64 `json:"GroupNum,omitnil,omitempty" name:"GroupNum"`

	// 实例最大消费组数量
	GroupNumLimit *int64 `json:"GroupNumLimit,omitnil,omitempty" name:"GroupNumLimit"`

	// 消息保留时间，小时为单位
	MessageRetention *int64 `json:"MessageRetention,omitnil,omitempty" name:"MessageRetention"`

	// 最大可调整消息保留时间，小时为单位
	RetentionUpperLimit *int64 `json:"RetentionUpperLimit,omitnil,omitempty" name:"RetentionUpperLimit"`

	// 最小可调整消息保留时间，小时为单位
	RetentionLowerLimit *int64 `json:"RetentionLowerLimit,omitnil,omitempty" name:"RetentionLowerLimit"`

	// TPS限流值
	TpsLimit *int64 `json:"TpsLimit,omitnil,omitempty" name:"TpsLimit"`

	// 弹性TPS限流值
	ScaledTpsLimit *int64 `json:"ScaledTpsLimit,omitnil,omitempty" name:"ScaledTpsLimit"`

	// 延迟消息最长时间，小时为单位
	MaxMessageDelay *int64 `json:"MaxMessageDelay,omitnil,omitempty" name:"MaxMessageDelay"`

	// 创建时间，**Unix时间戳（毫秒）**
	CreatedTime *int64 `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 消息发送接收比例
	SendReceiveRatio *float64 `json:"SendReceiveRatio,omitnil,omitempty" name:"SendReceiveRatio"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 接入点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndpointList []*Endpoint `json:"EndpointList,omitnil,omitempty" name:"EndpointList"`

	// 主题队列数上限
	TopicQueueNumUpperLimit *int64 `json:"TopicQueueNumUpperLimit,omitnil,omitempty" name:"TopicQueueNumUpperLimit"`

	// 主题队列数下限
	TopicQueueNumLowerLimit *int64 `json:"TopicQueueNumLowerLimit,omitnil,omitempty" name:"TopicQueueNumLowerLimit"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 实例状态，枚举值如下：
	// 
	// - RUNNING：运行中
	// - ABNORMAL：异常
	// - OVERDUE：隔离中
	// - DESTROYED：已销毁
	// - CREATING：创建中
	// - MODIFYING：变配中
	// - CREATE_FAILURE：创建失败
	// - MODIFY_FAILURE：变配失败
	// - DELETING：删除中
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 实例规格
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// 计费模式，枚举值如下：
	// 
	// - POSTPAID：后付费按量计费
	// - PREPAID：预付费包年包月
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 是否开启弹性TPS
	ScaledTpsEnabled *bool `json:"ScaledTpsEnabled,omitnil,omitempty" name:"ScaledTpsEnabled"`

	// 预付费集群是否自动续费，枚举值如下：
	// 
	// - 0: 不自动续费
	// - 1: 自动续费
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 到期时间，**Unix时间戳（毫秒）**
	ExpiryTime *int64 `json:"ExpiryTime,omitnil,omitempty" name:"ExpiryTime"`

	// 角色数量限制
	RoleNumLimit *int64 `json:"RoleNumLimit,omitnil,omitempty" name:"RoleNumLimit"`

	// 是否开启 ACL
	// 注意：此字段可能返回 null，表示取不到有效值。
	AclEnabled *bool `json:"AclEnabled,omitnil,omitempty" name:"AclEnabled"`

	// topic个数免费额度
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicNumLowerLimit *int64 `json:"TopicNumLowerLimit,omitnil,omitempty" name:"TopicNumLowerLimit"`

	// 最大可设置的topic个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicNumUpperLimit *int64 `json:"TopicNumUpperLimit,omitnil,omitempty" name:"TopicNumUpperLimit"`

	// 所属可用区列表，参考 [DescribeZones](https://cloud.tencent.com/document/product/1596/77929) 接口返回中的 [ZoneInfo](https://cloud.tencent.com/document/api/1596/77932#ZoneInfo) 数据结构。
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceResponseParams `json:"Response"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTClientRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 客户端详情
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`
}

type DescribeMQTTClientRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 客户端详情
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`
}

func (r *DescribeMQTTClientRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTClientRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClientId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMQTTClientRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTClientResponseParams struct {
	// 客户端唯一标识
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 客户端网络地址
	ClientAddress *string `json:"ClientAddress,omitnil,omitempty" name:"ClientAddress"`

	// MQTT 协议版本，4 表示 MQTT 3.1.1
	ProtocolVersion *int64 `json:"ProtocolVersion,omitnil,omitempty" name:"ProtocolVersion"`

	// 保持连接时间，单位：秒
	Keepalive *int64 `json:"Keepalive,omitnil,omitempty" name:"Keepalive"`

	// 连接状态，CONNECTED 已连接，DISCONNECTED 未连接
	ConnectionStatus *string `json:"ConnectionStatus,omitnil,omitempty" name:"ConnectionStatus"`

	// 客户端创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 上次建立连接时间
	ConnectTime *int64 `json:"ConnectTime,omitnil,omitempty" name:"ConnectTime"`

	// 上次断开连接时间，仅对持久会话（cleanSession=false）并且客户端当前未连接时有意义
	DisconnectTime *int64 `json:"DisconnectTime,omitnil,omitempty" name:"DisconnectTime"`

	// 客户端的订阅列表
	MQTTClientSubscriptions []*MQTTClientSubscription `json:"MQTTClientSubscriptions,omitnil,omitempty" name:"MQTTClientSubscriptions"`

	// 服务端到客户端的流量统计
	Inbound *StatisticsReport `json:"Inbound,omitnil,omitempty" name:"Inbound"`

	// 客户端到服务端的流量统计
	OutBound *StatisticsReport `json:"OutBound,omitnil,omitempty" name:"OutBound"`

	// cleansession标志
	CleanSession *bool `json:"CleanSession,omitnil,omitempty" name:"CleanSession"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMQTTClientResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMQTTClientResponseParams `json:"Response"`
}

func (r *DescribeMQTTClientResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTClientResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTInsPublicEndpointsRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeMQTTInsPublicEndpointsRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeMQTTInsPublicEndpointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTInsPublicEndpointsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMQTTInsPublicEndpointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTInsPublicEndpointsResponseParams struct {
	// 接入点
	Endpoints []*MQTTEndpointItem `json:"Endpoints,omitnil,omitempty" name:"Endpoints"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 带宽
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网访问规则
	Rules []*PublicAccessRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 公网状态：
	//     NORMAL-正常
	//     CLOSING-关闭中
	//     MODIFYING-修改中
	//     CREATING-开启中
	//     CLOSE-关闭
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMQTTInsPublicEndpointsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMQTTInsPublicEndpointsResponseParams `json:"Response"`
}

func (r *DescribeMQTTInsPublicEndpointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTInsPublicEndpointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTInsVPCEndpointsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeMQTTInsVPCEndpointsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeMQTTInsVPCEndpointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTInsVPCEndpointsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMQTTInsVPCEndpointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTInsVPCEndpointsResponseParams struct {
	// 接入点
	Endpoints []*MQTTEndpointItem `json:"Endpoints,omitnil,omitempty" name:"Endpoints"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMQTTInsVPCEndpointsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMQTTInsVPCEndpointsResponseParams `json:"Response"`
}

func (r *DescribeMQTTInsVPCEndpointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTInsVPCEndpointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTInstanceCertRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeMQTTInstanceCertRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeMQTTInstanceCertRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTInstanceCertRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMQTTInstanceCertRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTInstanceCertResponseParams struct {
	// 集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务端证书id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SSLServerCertId *string `json:"SSLServerCertId,omitnil,omitempty" name:"SSLServerCertId"`

	// CA证书id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SSLCaCertId *string `json:"SSLCaCertId,omitnil,omitempty" name:"SSLCaCertId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMQTTInstanceCertResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMQTTInstanceCertResponseParams `json:"Response"`
}

func (r *DescribeMQTTInstanceCertResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTInstanceCertResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTInstanceListRequestParams struct {
	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否包含新控制台集群：默认为包含
	IncludeNew *bool `json:"IncludeNew,omitnil,omitempty" name:"IncludeNew"`
}

type DescribeMQTTInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否包含新控制台集群：默认为包含
	IncludeNew *bool `json:"IncludeNew,omitnil,omitempty" name:"IncludeNew"`
}

func (r *DescribeMQTTInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "IncludeNew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMQTTInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTInstanceListResponseParams struct {
	// 查询总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例列表
	Data []*MQTTInstanceItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMQTTInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMQTTInstanceListResponseParams `json:"Response"`
}

func (r *DescribeMQTTInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTInstanceRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeMQTTInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeMQTTInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMQTTInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTInstanceResponseParams struct {
	// 实例类型，
	// EXPERIMENT 体验版
	// BASIC 基础版
	// PRO  专业版
	// PLATINUM 铂金版
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 主题数量
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// 实例最大主题数量
	TopicNumLimit *int64 `json:"TopicNumLimit,omitnil,omitempty" name:"TopicNumLimit"`

	// TPS限流值
	TpsLimit *int64 `json:"TpsLimit,omitnil,omitempty" name:"TpsLimit"`

	// 创建时间，秒为单位
	CreatedTime *int64 `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 实例状态
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 实例规格
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// 订阅数上限
	SubscriptionNumLimit *int64 `json:"SubscriptionNumLimit,omitnil,omitempty" name:"SubscriptionNumLimit"`

	// 客户端数量上限
	ClientNumLimit *int64 `json:"ClientNumLimit,omitnil,omitempty" name:"ClientNumLimit"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMQTTInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMQTTInstanceResponseParams `json:"Response"`
}

func (r *DescribeMQTTInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTMessageListRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 请求任务id
	TaskRequestId *string `json:"TaskRequestId,omitnil,omitempty" name:"TaskRequestId"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeMQTTMessageListRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 请求任务id
	TaskRequestId *string `json:"TaskRequestId,omitnil,omitempty" name:"TaskRequestId"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeMQTTMessageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTMessageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskRequestId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMQTTMessageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTMessageListResponseParams struct {
	// 查询总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 消息记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*MQTTMessageItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 请求任务id
	TaskRequestId *string `json:"TaskRequestId,omitnil,omitempty" name:"TaskRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMQTTMessageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMQTTMessageListResponseParams `json:"Response"`
}

func (r *DescribeMQTTMessageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTMessageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTMessageRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 消息ID
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`
}

type DescribeMQTTMessageRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 消息ID
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`
}

func (r *DescribeMQTTMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "MsgId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMQTTMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTMessageResponseParams struct {
	// 消息体
	Body *string `json:"Body,omitnil,omitempty" name:"Body"`

	// 详情参数
	Properties []*CustomMapEntry `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 生产时间
	ProduceTime *string `json:"ProduceTime,omitnil,omitempty" name:"ProduceTime"`

	// 消息ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageId *string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// 生产者地址
	ProducerAddr *string `json:"ProducerAddr,omitnil,omitempty" name:"ProducerAddr"`

	// Topic
	ShowTopicName *string `json:"ShowTopicName,omitnil,omitempty" name:"ShowTopicName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMQTTMessageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMQTTMessageResponseParams `json:"Response"`
}

func (r *DescribeMQTTMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTProductSKUListRequestParams struct {

}

type DescribeMQTTProductSKUListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeMQTTProductSKUListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTProductSKUListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMQTTProductSKUListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTProductSKUListResponseParams struct {
	// 查询总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// mqtt商品配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MQTTProductSkuList []*MQTTProductSkuItem `json:"MQTTProductSkuList,omitnil,omitempty" name:"MQTTProductSkuList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMQTTProductSKUListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMQTTProductSKUListResponseParams `json:"Response"`
}

func (r *DescribeMQTTProductSKUListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTProductSKUListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTTopicListRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeMQTTTopicListRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeMQTTTopicListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTTopicListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMQTTTopicListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTTopicListResponseParams struct {
	// 查询总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 主题列表
	Data []*MQTTTopicItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMQTTTopicListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMQTTTopicListResponseParams `json:"Response"`
}

func (r *DescribeMQTTTopicListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTTopicListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTTopicRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
}

type DescribeMQTTTopicRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
}

func (r *DescribeMQTTTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMQTTTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTTopicResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 创建时间，秒为单位
	CreatedTime *int64 `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMQTTTopicResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMQTTTopicResponseParams `json:"Response"`
}

func (r *DescribeMQTTTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTUserListRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeMQTTUserListRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeMQTTUserListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTUserListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMQTTUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMQTTUserListResponseParams struct {
	// 查询总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 角色信息列表
	Data []*MQTTUserItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMQTTUserListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMQTTUserListResponseParams `json:"Response"`
}

func (r *DescribeMQTTUserListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMQTTUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMessageListRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，从 [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030) 接口返回的 [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem) 或控制台获得。
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 要查询消息的开始时间，**Unix时间戳（毫秒）**
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 要查询消息的结束时间，**Unix时间戳（毫秒）**
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 一次查询标识。第一次查询可传空字符串，当查询结果涉及分页，请求下一页数据时该入参的值取上一次请求响应中的出参TaskRequestId 值即可。
	TaskRequestId *string `json:"TaskRequestId,omitnil,omitempty" name:"TaskRequestId"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`

	// 消息 ID，从 [DescribeMessageList](https://cloud.tencent.com/document/api/1493/114593) 接口返回的 [MessageItem](https://cloud.tencent.com/document/api/1493/96031#MessageItem) 或业务日志中获得。
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 消息 Key，从 [DescribeMessageList](https://cloud.tencent.com/document/api/1493/114593) 接口返回的 [MessageItem](https://cloud.tencent.com/document/api/1493/96031#MessageItem) 或业务日志中获得。
	MsgKey *string `json:"MsgKey,omitnil,omitempty" name:"MsgKey"`

	// 查询最近N条消息 最大不超过1024，默认-1为其他查询条件
	RecentMessageNum *int64 `json:"RecentMessageNum,omitnil,omitempty" name:"RecentMessageNum"`

	// 是否查询死信消息，默认为false
	QueryDeadLetterMessage *bool `json:"QueryDeadLetterMessage,omitnil,omitempty" name:"QueryDeadLetterMessage"`

	// 消息 Tag，从 [DescribeMessageList](https://cloud.tencent.com/document/api/1493/114593) 接口返回的 [MessageItem](https://cloud.tencent.com/document/api/1493/96031#MessageItem) 或业务日志中获得。
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`
}

type DescribeMessageListRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，从 [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030) 接口返回的 [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem) 或控制台获得。
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 要查询消息的开始时间，**Unix时间戳（毫秒）**
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 要查询消息的结束时间，**Unix时间戳（毫秒）**
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 一次查询标识。第一次查询可传空字符串，当查询结果涉及分页，请求下一页数据时该入参的值取上一次请求响应中的出参TaskRequestId 值即可。
	TaskRequestId *string `json:"TaskRequestId,omitnil,omitempty" name:"TaskRequestId"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`

	// 消息 ID，从 [DescribeMessageList](https://cloud.tencent.com/document/api/1493/114593) 接口返回的 [MessageItem](https://cloud.tencent.com/document/api/1493/96031#MessageItem) 或业务日志中获得。
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 消息 Key，从 [DescribeMessageList](https://cloud.tencent.com/document/api/1493/114593) 接口返回的 [MessageItem](https://cloud.tencent.com/document/api/1493/96031#MessageItem) 或业务日志中获得。
	MsgKey *string `json:"MsgKey,omitnil,omitempty" name:"MsgKey"`

	// 查询最近N条消息 最大不超过1024，默认-1为其他查询条件
	RecentMessageNum *int64 `json:"RecentMessageNum,omitnil,omitempty" name:"RecentMessageNum"`

	// 是否查询死信消息，默认为false
	QueryDeadLetterMessage *bool `json:"QueryDeadLetterMessage,omitnil,omitempty" name:"QueryDeadLetterMessage"`

	// 消息 Tag，从 [DescribeMessageList](https://cloud.tencent.com/document/api/1493/114593) 接口返回的 [MessageItem](https://cloud.tencent.com/document/api/1493/96031#MessageItem) 或业务日志中获得。
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`
}

func (r *DescribeMessageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskRequestId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ConsumerGroup")
	delete(f, "MsgId")
	delete(f, "MsgKey")
	delete(f, "RecentMessageNum")
	delete(f, "QueryDeadLetterMessage")
	delete(f, "Tag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMessageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMessageListResponseParams struct {
	// 查询总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 消息记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*MessageItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 一次查询ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskRequestId *string `json:"TaskRequestId,omitnil,omitempty" name:"TaskRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMessageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMessageListResponseParams `json:"Response"`
}

func (r *DescribeMessageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMessageRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，从 [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030) 接口返回的 [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem) 或控制台获得。
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 消息 ID，从 [DescribeMessageList](https://cloud.tencent.com/document/api/1493/114593) 接口或业务日志中获得。
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否是死信消息，默认为false
	QueryDeadLetterMessage *bool `json:"QueryDeadLetterMessage,omitnil,omitempty" name:"QueryDeadLetterMessage"`

	// 是否是延时消息，默认为false
	QueryDelayMessage *bool `json:"QueryDelayMessage,omitnil,omitempty" name:"QueryDelayMessage"`
}

type DescribeMessageRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，从 [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030) 接口返回的 [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem) 或控制台获得。
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 消息 ID，从 [DescribeMessageList](https://cloud.tencent.com/document/api/1493/114593) 接口或业务日志中获得。
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否是死信消息，默认为false
	QueryDeadLetterMessage *bool `json:"QueryDeadLetterMessage,omitnil,omitempty" name:"QueryDeadLetterMessage"`

	// 是否是延时消息，默认为false
	QueryDelayMessage *bool `json:"QueryDelayMessage,omitnil,omitempty" name:"QueryDelayMessage"`
}

func (r *DescribeMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "MsgId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "QueryDeadLetterMessage")
	delete(f, "QueryDelayMessage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMessageResponseParams struct {
	// 消息体
	Body *string `json:"Body,omitnil,omitempty" name:"Body"`

	// 详情参数
	Properties *string `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 生产时间
	ProduceTime *string `json:"ProduceTime,omitnil,omitempty" name:"ProduceTime"`

	// 消息ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageId *string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// 生产者地址
	ProducerAddr *string `json:"ProducerAddr,omitnil,omitempty" name:"ProducerAddr"`

	// 消息消费情况列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageTracks []*MessageTrackItem `json:"MessageTracks,omitnil,omitempty" name:"MessageTracks"`

	// 主题名称
	ShowTopicName *string `json:"ShowTopicName,omitnil,omitempty" name:"ShowTopicName"`

	// 消息消费情况列表总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageTracksCount *int64 `json:"MessageTracksCount,omitnil,omitempty" name:"MessageTracksCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMessageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMessageResponseParams `json:"Response"`
}

func (r *DescribeMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMessageTraceRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，从 [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030) 接口返回的 [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem) 或控制台获得。
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 消息 ID，从 [DescribeMessageList](https://cloud.tencent.com/document/api/1493/114593) 接口返回的 [MessageItem](https://cloud.tencent.com/document/api/1493/96031#MessageItem) 或业务日志中获得。
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 是否是死信消息，默认为false
	QueryDeadLetterMessage *bool `json:"QueryDeadLetterMessage,omitnil,omitempty" name:"QueryDeadLetterMessage"`

	// 是否是延时消息，默认为false
	QueryDelayMessage *bool `json:"QueryDelayMessage,omitnil,omitempty" name:"QueryDelayMessage"`
}

type DescribeMessageTraceRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，从 [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030) 接口返回的 [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem) 或控制台获得。
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 消息 ID，从 [DescribeMessageList](https://cloud.tencent.com/document/api/1493/114593) 接口返回的 [MessageItem](https://cloud.tencent.com/document/api/1493/96031#MessageItem) 或业务日志中获得。
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 是否是死信消息，默认为false
	QueryDeadLetterMessage *bool `json:"QueryDeadLetterMessage,omitnil,omitempty" name:"QueryDeadLetterMessage"`

	// 是否是延时消息，默认为false
	QueryDelayMessage *bool `json:"QueryDelayMessage,omitnil,omitempty" name:"QueryDelayMessage"`
}

func (r *DescribeMessageTraceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageTraceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "MsgId")
	delete(f, "QueryDeadLetterMessage")
	delete(f, "QueryDelayMessage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMessageTraceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMessageTraceResponseParams struct {
	// 主题名称
	ShowTopicName *string `json:"ShowTopicName,omitnil,omitempty" name:"ShowTopicName"`

	// 轨迹详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*MessageTraceItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMessageTraceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMessageTraceResponseParams `json:"Response"`
}

func (r *DescribeMessageTraceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageTraceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigratingGroupStatsRequestParams struct {
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 消费组名称，可在[DescribeSourceClusterGroupList](https://cloud.tencent.com/document/api/1493/118006)接口返回的[SourceClusterGroupConfig](https://cloud.tencent.com/document/api/1493/96031#SourceClusterGroupConfig)或控制台中获取。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 命名空间，仅迁移至4.x集群有效，可在[DescribeSourceClusterGroupList](https://cloud.tencent.com/document/api/1493/118006)接口返回的[SourceClusterGroupConfig](https://cloud.tencent.com/document/api/1493/96031#SourceClusterGroupConfig)或控制台中获取。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type DescribeMigratingGroupStatsRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 消费组名称，可在[DescribeSourceClusterGroupList](https://cloud.tencent.com/document/api/1493/118006)接口返回的[SourceClusterGroupConfig](https://cloud.tencent.com/document/api/1493/96031#SourceClusterGroupConfig)或控制台中获取。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 命名空间，仅迁移至4.x集群有效，可在[DescribeSourceClusterGroupList](https://cloud.tencent.com/document/api/1493/118006)接口返回的[SourceClusterGroupConfig](https://cloud.tencent.com/document/api/1493/96031#SourceClusterGroupConfig)或控制台中获取。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *DescribeMigratingGroupStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigratingGroupStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "GroupName")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigratingGroupStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigratingGroupStatsResponseParams struct {
	// 源集群消费组堆积
	SourceConsumeLag *int64 `json:"SourceConsumeLag,omitnil,omitempty" name:"SourceConsumeLag"`

	// 目标集群消费组堆积
	TargetConsumeLag *int64 `json:"TargetConsumeLag,omitnil,omitempty" name:"TargetConsumeLag"`

	// 源集群连接客户端列表
	SourceConsumerClients []*ConsumerClient `json:"SourceConsumerClients,omitnil,omitempty" name:"SourceConsumerClients"`

	// 目标集群连接客户端列表
	TargetConsumerClients []*ConsumerClient `json:"TargetConsumerClients,omitnil,omitempty" name:"TargetConsumerClients"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMigratingGroupStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigratingGroupStatsResponseParams `json:"Response"`
}

func (r *DescribeMigratingGroupStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigratingGroupStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigratingTopicListRequestParams struct {
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeMigratingTopicListRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeMigratingTopicListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigratingTopicListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigratingTopicListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigratingTopicListResponseParams struct {
	// 查询总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 主题列表
	MigrateTopics []*MigratingTopic `json:"MigrateTopics,omitnil,omitempty" name:"MigrateTopics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMigratingTopicListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigratingTopicListResponseParams `json:"Response"`
}

func (r *DescribeMigratingTopicListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigratingTopicListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigratingTopicStatsRequestParams struct {
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 主题名称，可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 命名空间，仅迁移至4.x集群有效，可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type DescribeMigratingTopicStatsRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 主题名称，可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 命名空间，仅迁移至4.x集群有效，可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *DescribeMigratingTopicStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigratingTopicStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "TopicName")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigratingTopicStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigratingTopicStatsResponseParams struct {
	// 源集群的消费者数量
	SourceClusterConsumerCount *int64 `json:"SourceClusterConsumerCount,omitnil,omitempty" name:"SourceClusterConsumerCount"`

	// 目标集群的消费者数量
	TargetClusterConsumerCount *int64 `json:"TargetClusterConsumerCount,omitnil,omitempty" name:"TargetClusterConsumerCount"`

	// 源集群消费组列表
	SourceClusterConsumerGroups []*string `json:"SourceClusterConsumerGroups,omitnil,omitempty" name:"SourceClusterConsumerGroups"`

	// 目标集群消费组列表
	TargetClusterConsumerGroups []*string `json:"TargetClusterConsumerGroups,omitnil,omitempty" name:"TargetClusterConsumerGroups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMigratingTopicStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigratingTopicStatsResponseParams `json:"Response"`
}

func (r *DescribeMigratingTopicStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigratingTopicStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationTaskListRequestParams struct {
	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeMigrationTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeMigrationTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrationTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationTaskListResponseParams struct {
	// 查询总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 迁移任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*MigrationTaskItem `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMigrationTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrationTaskListResponseParams `json:"Response"`
}

func (r *DescribeMigrationTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductSKUsRequestParams struct {

}

type DescribeProductSKUsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeProductSKUsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductSKUsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductSKUsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductSKUsResponseParams struct {
	// 商品配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ProductSKU `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProductSKUsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductSKUsResponseParams `json:"Response"`
}

func (r *DescribeProductSKUsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductSKUsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoleListRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRoleListRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeRoleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoleListResponseParams struct {
	// 查询总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 角色信息列表
	Data []*RoleItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRoleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoleListResponseParams `json:"Response"`
}

func (r *DescribeRoleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmoothMigrationTaskListRequestParams struct {
	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeSmoothMigrationTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeSmoothMigrationTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmoothMigrationTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSmoothMigrationTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmoothMigrationTaskListResponseParams struct {
	// 查询总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务列表	
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*SmoothMigrationTaskItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSmoothMigrationTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSmoothMigrationTaskListResponseParams `json:"Response"`
}

func (r *DescribeSmoothMigrationTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmoothMigrationTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSourceClusterGroupListRequestParams struct {
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台上获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeSourceClusterGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台上获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeSourceClusterGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceClusterGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSourceClusterGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSourceClusterGroupListResponseParams struct {
	// 查询总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 消费组配置列表
	Groups []*SourceClusterGroupConfig `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSourceClusterGroupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSourceClusterGroupListResponseParams `json:"Response"`
}

func (r *DescribeSourceClusterGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceClusterGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicListByGroupRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTopicListByGroupRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeTopicListByGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicListByGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ConsumerGroup")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicListByGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicListByGroupResponseParams struct {
	// 查询总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 主题列表
	Data []*SubscriptionData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicListByGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicListByGroupResponseParams `json:"Response"`
}

func (r *DescribeTopicListByGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicListByGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicListRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTopicListRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeTopicListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicListResponseParams struct {
	// 查询总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 主题列表
	Data []*TopicItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicListResponseParams `json:"Response"`
}

func (r *DescribeTopicListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，从 [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030) 接口返回的 [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem) 或控制台获得。
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTopicRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，从 [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030) 接口返回的 [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem) 或控制台获得。
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 过滤查询条件列表，请在引用此参数的API说明中了解使用方法。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 主题类型
	// UNSPECIFIED:未指定,
	// NORMAL:普通消息,
	// FIFO:顺序消息,
	// DELAY:延时消息,
	// TRANSACTION:事务消息
	TopicType *string `json:"TopicType,omitnil,omitempty" name:"TopicType"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 创建时间，**Unix时间戳（毫秒）**
	CreatedTime *int64 `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 最后写入时间，**Unix时间戳（毫秒）**
	LastUpdateTime *int64 `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// 订阅数量
	SubscriptionCount *int64 `json:"SubscriptionCount,omitnil,omitempty" name:"SubscriptionCount"`

	// 订阅关系列表
	SubscriptionData []*SubscriptionData `json:"SubscriptionData,omitnil,omitempty" name:"SubscriptionData"`

	// 消息保留时长，单位：小时
	MsgTTL *int64 `json:"MsgTTL,omitnil,omitempty" name:"MsgTTL"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicResponseParams `json:"Response"`
}

func (r *DescribeTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailedRolePerm struct {
	// 权限对应的资源
	// 可以是主题名称，从 [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030) 接口返回的 [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem) 或控制台获得。
	// 可以是消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
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

// Predefined struct for user
type DoHealthCheckOnMigratingTopicRequestParams struct {
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 主题名称，可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 必填，是否忽略当前检查
	IgnoreCheck *bool `json:"IgnoreCheck,omitnil,omitempty" name:"IgnoreCheck"`

	// 命名空间，仅迁移至4.x集群有效，可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type DoHealthCheckOnMigratingTopicRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 主题名称，可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 必填，是否忽略当前检查
	IgnoreCheck *bool `json:"IgnoreCheck,omitnil,omitempty" name:"IgnoreCheck"`

	// 命名空间，仅迁移至4.x集群有效，可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *DoHealthCheckOnMigratingTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DoHealthCheckOnMigratingTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "TopicName")
	delete(f, "IgnoreCheck")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DoHealthCheckOnMigratingTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DoHealthCheckOnMigratingTopicResponseParams struct {
	// 是否通过	
	Passed *bool `json:"Passed,omitnil,omitempty" name:"Passed"`

	// 健康检查返回的错误信息
	// NotChecked 未执行检查， 
	// Unknown 未知错误, 
	// TopicNotImported 主题未导入,
	// TopicNotExistsInSourceCluster 主题在源集群中不存在, 
	// TopicNotExistsInTargetCluster 主题在目标集群中不存在, 
	// ConsumerConnectedOnTarget 目标集群上存在消费者连接, 
	// SourceTopicHasNewMessagesIn5Minutes 源集群主题前5分钟内有新消息写入, 
	// TargetTopicHasNewMessagesIn5Minutes 目标集群主题前5分钟内有新消息写入, 
	// SourceTopicHasNoMessagesIn5Minutes 源集群前5分钟内没有新消息写入, 
	// TargetTopicHasNoMessagesIn5Minutes 源集群前5分钟内没有新消息写入, 
	// ConsumerGroupCountNotMatch 订阅组数量不一致, 
	// SourceTopicHasUnconsumedMessages 源集群主题存在未消费消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 健康检查返回的错误信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReasonList []*string `json:"ReasonList,omitnil,omitempty" name:"ReasonList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DoHealthCheckOnMigratingTopicResponse struct {
	*tchttp.BaseResponse
	Response *DoHealthCheckOnMigratingTopicResponseParams `json:"Response"`
}

func (r *DoHealthCheckOnMigratingTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DoHealthCheckOnMigratingTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Endpoint struct {
	// 接入点类型，枚举值如下：
	// 
	// - VPC：VPC 网络
	// 
	// - PUBLIC：公网
	// 
	// - INTERNAL：支撑网
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 状态，
	// OPEN 开启，
	// CLOSE 关闭，
	// PROCESSING 操作中，
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 付费类型，仅公网
	// PREPAID 包年包月
	// POSTPAID 按量付费
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 接入点地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndpointUrl *string `json:"EndpointUrl,omitnil,omitempty" name:"EndpointUrl"`

	// VPC ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 公网带宽，Mbps为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网放通规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpRules []*IpRule `json:"IpRules,omitnil,omitempty" name:"IpRules"`

	// 公网是否按流量计费
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingFlow *bool `json:"BillingFlow,omitnil,omitempty" name:"BillingFlow"`
}

type Filter struct {
	// 过滤条件参数名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤条件的值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FusionInstanceItem struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 实例类型，
	// EXPERIMENT，体验版
	// BASIC，基础版
	// PRO，专业版
	// PLATINUM，铂金版
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例状态，
	// RUNNING, 运行中
	// MAINTAINING，维护中
	// ABNORMAL，异常
	// OVERDUE，欠费
	// DESTROYED，已删除
	// CREATING，创建中
	// MODIFYING，变配中
	// CREATE_FAILURE，创建失败
	// MODIFY_FAILURE，变配失败
	// DELETING，删除中
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 实例主题数上限
	TopicNumLimit *int64 `json:"TopicNumLimit,omitnil,omitempty" name:"TopicNumLimit"`

	// 实例消费组数量上限
	GroupNumLimit *int64 `json:"GroupNumLimit,omitnil,omitempty" name:"GroupNumLimit"`

	// 计费模式，枚举值如下：
	// 
	// - POSTPAID：按量计费
	// 
	// - PREPAID：包年包月
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 到期时间，**Unix时间戳（毫秒）**
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiryTime *int64 `json:"ExpiryTime,omitnil,omitempty" name:"ExpiryTime"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 主题数量
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// 消费组数量
	GroupNum *int64 `json:"GroupNum,omitnil,omitempty" name:"GroupNum"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 商品规格
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// TPS限流值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TpsLimit *int64 `json:"TpsLimit,omitnil,omitempty" name:"TpsLimit"`

	// 弹性TPS限流值
	ScaledTpsLimit *int64 `json:"ScaledTpsLimit,omitnil,omitempty" name:"ScaledTpsLimit"`

	// 消息保留时间，小时为单位
	MessageRetention *int64 `json:"MessageRetention,omitnil,omitempty" name:"MessageRetention"`

	// 延迟消息最大时长，小时为单位
	MaxMessageDelay *int64 `json:"MaxMessageDelay,omitnil,omitempty" name:"MaxMessageDelay"`

	// 预付费集群是否自动续费，枚举值如下：
	// 
	// - 0: 不自动续费
	// - 1: 自动续费
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 4.x独有数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceItemExtraInfo *InstanceItemExtraInfo `json:"InstanceItemExtraInfo,omitnil,omitempty" name:"InstanceItemExtraInfo"`

	// 预销毁时间，**Unix时间戳（毫秒）**
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestroyTime *int64 `json:"DestroyTime,omitnil,omitempty" name:"DestroyTime"`

	// 所属可用区列表，参考 [DescribeZones](https://cloud.tencent.com/document/product/1596/77929) 接口返回中的 [ZoneInfo](https://cloud.tencent.com/document/api/1596/77932#ZoneInfo) 数据结构。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 是否开启删除保护
	EnableDeletionProtection *bool `json:"EnableDeletionProtection,omitnil,omitempty" name:"EnableDeletionProtection"`
}

// Predefined struct for user
type ImportSourceClusterConsumerGroupsRequestParams struct {
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 待导入的消费组列表
	GroupList []*SourceClusterGroupConfig `json:"GroupList,omitnil,omitempty" name:"GroupList"`
}

type ImportSourceClusterConsumerGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 待导入的消费组列表
	GroupList []*SourceClusterGroupConfig `json:"GroupList,omitnil,omitempty" name:"GroupList"`
}

func (r *ImportSourceClusterConsumerGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportSourceClusterConsumerGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "GroupList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportSourceClusterConsumerGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportSourceClusterConsumerGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImportSourceClusterConsumerGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ImportSourceClusterConsumerGroupsResponseParams `json:"Response"`
}

func (r *ImportSourceClusterConsumerGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportSourceClusterConsumerGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportSourceClusterTopicsRequestParams struct {
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 待导入的主题列表
	TopicList []*SourceClusterTopicConfig `json:"TopicList,omitnil,omitempty" name:"TopicList"`
}

type ImportSourceClusterTopicsRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 待导入的主题列表
	TopicList []*SourceClusterTopicConfig `json:"TopicList,omitnil,omitempty" name:"TopicList"`
}

func (r *ImportSourceClusterTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportSourceClusterTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "TopicList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportSourceClusterTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportSourceClusterTopicsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImportSourceClusterTopicsResponse struct {
	*tchttp.BaseResponse
	Response *ImportSourceClusterTopicsResponseParams `json:"Response"`
}

func (r *ImportSourceClusterTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportSourceClusterTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceItem struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 实例类型，
	// EXPERIMENT，体验版
	// BASIC，基础版
	// PRO，专业版
	// PLATINUM，铂金版
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例状态，
	// RUNNING, 运行中
	// MAINTAINING，维护中
	// ABNORMAL，异常
	// OVERDUE，欠费
	// DESTROYED，已删除
	// CREATING，创建中
	// MODIFYING，变配中
	// CREATE_FAILURE，创建失败
	// MODIFY_FAILURE，变配失败
	// DELETING，删除中
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 实例主题数上限
	TopicNumLimit *int64 `json:"TopicNumLimit,omitnil,omitempty" name:"TopicNumLimit"`

	// 实例消费组数量上限
	GroupNumLimit *int64 `json:"GroupNumLimit,omitnil,omitempty" name:"GroupNumLimit"`

	// 计费模式，枚举值如下：
	// 
	// - POSTPAID：按量计费
	// 
	// - PREPAID：包年包月
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 到期时间戳，**Unix时间戳（毫秒）**
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiryTime *int64 `json:"ExpiryTime,omitnil,omitempty" name:"ExpiryTime"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 主题数量
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// 消费组数量
	GroupNum *int64 `json:"GroupNum,omitnil,omitempty" name:"GroupNum"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 商品规格
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// TPS限流值
	TpsLimit *int64 `json:"TpsLimit,omitnil,omitempty" name:"TpsLimit"`

	// 弹性TPS限流值
	ScaledTpsLimit *int64 `json:"ScaledTpsLimit,omitnil,omitempty" name:"ScaledTpsLimit"`

	// 消息保留时间，小时为单位
	MessageRetention *int64 `json:"MessageRetention,omitnil,omitempty" name:"MessageRetention"`

	// 延迟消息最大时长，小时为单位
	MaxMessageDelay *int64 `json:"MaxMessageDelay,omitnil,omitempty" name:"MaxMessageDelay"`

	// 是否自动续费，仅针对预付费集群（0: 不自动续费；1:自动续费）
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type InstanceItemExtraInfo struct {
	// 是否vip
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsVip *bool `json:"IsVip,omitnil,omitempty" name:"IsVip"`

	// 4.x专享集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	VipInstanceStatus *int64 `json:"VipInstanceStatus,omitnil,omitempty" name:"VipInstanceStatus"`

	// 专享集群峰值带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxBandWidth *int64 `json:"MaxBandWidth,omitnil,omitempty" name:"MaxBandWidth"`

	// 专享集群规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecName *string `json:"SpecName,omitnil,omitempty" name:"SpecName"`

	// 专享集群节点数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeCount *int64 `json:"NodeCount,omitnil,omitempty" name:"NodeCount"`

	// 专享集群最大存储
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxStorage *int64 `json:"MaxStorage,omitnil,omitempty" name:"MaxStorage"`

	// 专享集群最大保留时间，单位：小时
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRetention *int64 `json:"MaxRetention,omitnil,omitempty" name:"MaxRetention"`

	// 专项集群最大保留时间，单位：小时
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinRetention *int64 `json:"MinRetention,omitnil,omitempty" name:"MinRetention"`

	// 4.0共享集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatus *int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 是否已冻结
	IsFrozen *bool `json:"IsFrozen,omitnil,omitempty" name:"IsFrozen"`
}

type IpRule struct {
	// IP地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 是否允许放行，默认为false表示拒绝
	Allow *bool `json:"Allow,omitnil,omitempty" name:"Allow"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type MQTTClientSubscription struct {
	// topic 订阅
	TopicFilter *string `json:"TopicFilter,omitnil,omitempty" name:"TopicFilter"`

	// 服务质量等级
	Qos *int64 `json:"Qos,omitnil,omitempty" name:"Qos"`
}

type MQTTEndpointItem struct {
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 接入点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// vpc信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 主机
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 接入点ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`
}

type MQTTInstanceItem struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 实例类型，
	// BASIC，基础版
	// PRO，专业版
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例状态，
	// RUNNING, 运行中
	// MAINTAINING，维护中
	// ABNORMAL，异常
	// OVERDUE，欠费
	// DESTROYED，已删除
	// CREATING，创建中
	// MODIFYING，变配中
	// CREATE_FAILURE，创建失败
	// MODIFY_FAILURE，变配失败
	// DELETING，删除中
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 实例主题数上限
	TopicNumLimit *int64 `json:"TopicNumLimit,omitnil,omitempty" name:"TopicNumLimit"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 主题数量
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// 商品规格
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// 弹性TPS限流值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TpsLimit *int64 `json:"TpsLimit,omitnil,omitempty" name:"TpsLimit"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 订阅关系上限
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubscriptionNumLimit *int64 `json:"SubscriptionNumLimit,omitnil,omitempty" name:"SubscriptionNumLimit"`

	// 客户端连接数上线
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientNumLimit *int64 `json:"ClientNumLimit,omitnil,omitempty" name:"ClientNumLimit"`
}

type MQTTMessageItem struct {
	// 消息ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 消息tag
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 消息key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keys *string `json:"Keys,omitnil,omitempty" name:"Keys"`

	// 客户端地址	
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProducerAddr *string `json:"ProducerAddr,omitnil,omitempty" name:"ProducerAddr"`

	// 消息发送时间	
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProduceTime *string `json:"ProduceTime,omitnil,omitempty" name:"ProduceTime"`

	// 死信重发次数	
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterResendTimes *int64 `json:"DeadLetterResendTimes,omitnil,omitempty" name:"DeadLetterResendTimes"`

	// 死信重发成功次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterResendSuccessTimes *int64 `json:"DeadLetterResendSuccessTimes,omitnil,omitempty" name:"DeadLetterResendSuccessTimes"`

	// 子topic
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubTopic *string `json:"SubTopic,omitnil,omitempty" name:"SubTopic"`

	// 消息质量等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Qos *string `json:"Qos,omitnil,omitempty" name:"Qos"`
}

type MQTTProductSkuItem struct {
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// cide
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// sale
	// 注意：此字段可能返回 null，表示取不到有效值。
	OnSale *bool `json:"OnSale,omitnil,omitempty" name:"OnSale"`

	// topic num限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicNumLimit *int64 `json:"TopicNumLimit,omitnil,omitempty" name:"TopicNumLimit"`

	// tps
	// 注意：此字段可能返回 null，表示取不到有效值。
	TpsLimit *int64 `json:"TpsLimit,omitnil,omitempty" name:"TpsLimit"`

	// 客户端连接数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientNumLimit *int64 `json:"ClientNumLimit,omitnil,omitempty" name:"ClientNumLimit"`

	// 订阅数限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubscriptionNumLimit *int64 `json:"SubscriptionNumLimit,omitnil,omitempty" name:"SubscriptionNumLimit"`

	// 代理核
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxySpecCore *int64 `json:"ProxySpecCore,omitnil,omitempty" name:"ProxySpecCore"`

	// 代理内存
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxySpecMemory *int64 `json:"ProxySpecMemory,omitnil,omitempty" name:"ProxySpecMemory"`

	// 代理总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxySpecCount *int64 `json:"ProxySpecCount,omitnil,omitempty" name:"ProxySpecCount"`
}

type MQTTTopicItem struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 主题描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type MQTTUserItem struct {
	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 是否开启消费
	PermRead *bool `json:"PermRead,omitnil,omitempty" name:"PermRead"`

	// 是否开启生产
	PermWrite *bool `json:"PermWrite,omitnil,omitempty" name:"PermWrite"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 创建时间，秒为单位
	CreatedTime *int64 `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 修改时间，秒为单位
	ModifiedTime *int64 `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`
}

type MessageItem struct {
	// 消息ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 消息tag
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 消息key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keys *string `json:"Keys,omitnil,omitempty" name:"Keys"`

	// 客户端地址	
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProducerAddr *string `json:"ProducerAddr,omitnil,omitempty" name:"ProducerAddr"`

	// 消息发送时间	
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProduceTime *string `json:"ProduceTime,omitnil,omitempty" name:"ProduceTime"`

	// 死信重发次数	
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterResendTimes *int64 `json:"DeadLetterResendTimes,omitnil,omitempty" name:"DeadLetterResendTimes"`

	// 死信重发成功次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeadLetterResendSuccessTimes *int64 `json:"DeadLetterResendSuccessTimes,omitnil,omitempty" name:"DeadLetterResendSuccessTimes"`
}

type MessageTraceItem struct {
	// 消息处理阶段，枚举值如下：
	// 
	// - produce：消息生产
	// 
	// - persist：消息存储
	// 
	// - consume：消息消费
	// 注意：此字段可能返回 null，表示取不到有效值。
	Stage *string `json:"Stage,omitnil,omitempty" name:"Stage"`

	// 轨迹详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`
}

type MessageTrackItem struct {
	// 消费组名称
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`

	// 消费状态, CONSUMED: 已消费 CONSUMED_BUT_FILTERED: 已过滤 NOT_CONSUME: 未消费 ENTER_RETRY: 进入重试队列 ENTER_DLQ: 进入死信队列 UNKNOWN: 查询不到消费状态
	ConsumeStatus *string `json:"ConsumeStatus,omitnil,omitempty" name:"ConsumeStatus"`

	// track类型
	TrackType *string `json:"TrackType,omitnil,omitempty" name:"TrackType"`

	// 异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceptionDesc *string `json:"ExceptionDesc,omitnil,omitempty" name:"ExceptionDesc"`
}

type MigratingTopic struct {
	// 主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 迁移状态 
	// S_RW_D_NA 源集群读写，
	// S_RW_D_R 源集群读写目标集群读，
	// S_RW_D_RW 源集群读写目标集群读写，
	// S_R_D_RW 源集群读目标集群读写，
	// S_NA_D_RW 目标集群读写
	MigrationStatus *string `json:"MigrationStatus,omitnil,omitempty" name:"MigrationStatus"`

	// 是否完成健康检查	
	HealthCheckPassed *bool `json:"HealthCheckPassed,omitnil,omitempty" name:"HealthCheckPassed"`

	// 上次健康检查返回的错误信息，仅在HealthCheckPassed为false时有效。 NotChecked 未执行检查， Unknown 未知错误, TopicNotImported 主题未导入, TopicNotExistsInSourceCluster 主题在源集群中不存在, TopicNotExistsInTargetCluster 主题在目标集群中不存在, ConsumerConnectedOnTarget 目标集群上存在消费者连接, SourceTopicHasNewMessagesIn5Minutes 源集群主题前5分钟内有新消息写入, TargetTopicHasNewMessagesIn5Minutes 目标集群主题前5分钟内有新消息写入, SourceTopicHasNoMessagesIn5Minutes 源集群前5分钟内没有新消息写入, TargetTopicHasNoMessagesIn5Minutes 源集群前5分钟内没有新消息写入, ConsumerGroupCountNotMatch 订阅组数量不一致, SourceTopicHasUnconsumedMessages 源集群主题存在未消费消息,
	HealthCheckError *string `json:"HealthCheckError,omitnil,omitempty" name:"HealthCheckError"`

	// 命名空间，仅4.x集群有效
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 4.x的命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceV4 *string `json:"NamespaceV4,omitnil,omitempty" name:"NamespaceV4"`

	// 4.x的主题名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicNameV4 *string `json:"TopicNameV4,omitnil,omitempty" name:"TopicNameV4"`

	// 4.x的完整命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullNamespaceV4 *string `json:"FullNamespaceV4,omitnil,omitempty" name:"FullNamespaceV4"`

	// 上次健康检查返回的错误列表
	HealthCheckErrorList []*string `json:"HealthCheckErrorList,omitnil,omitempty" name:"HealthCheckErrorList"`
}

type MigrationTaskItem struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 0 - 未指定（存量）
	// 1 - 元数据导入
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 主题总数
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// 消费组总数
	GroupNum *int64 `json:"GroupNum,omitnil,omitempty" name:"GroupNum"`

	// 任务状态： 0，迁移中 1，迁移成功 2，迁移完成，只有部分数据完成迁移
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type ModifyConsumerGroupRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`

	// 是否开启消费
	ConsumeEnable *bool `json:"ConsumeEnable,omitnil,omitempty" name:"ConsumeEnable"`

	// 顺序投递：true
	// 并发投递：false
	ConsumeMessageOrderly *bool `json:"ConsumeMessageOrderly,omitnil,omitempty" name:"ConsumeMessageOrderly"`

	// 最大重试次数，取值范围0～1000
	MaxRetryTimes *int64 `json:"MaxRetryTimes,omitnil,omitempty" name:"MaxRetryTimes"`

	// 备注信息，最多 128 个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyConsumerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`

	// 是否开启消费
	ConsumeEnable *bool `json:"ConsumeEnable,omitnil,omitempty" name:"ConsumeEnable"`

	// 顺序投递：true
	// 并发投递：false
	ConsumeMessageOrderly *bool `json:"ConsumeMessageOrderly,omitnil,omitempty" name:"ConsumeMessageOrderly"`

	// 最大重试次数，取值范围0～1000
	MaxRetryTimes *int64 `json:"MaxRetryTimes,omitnil,omitempty" name:"MaxRetryTimes"`

	// 备注信息，最多 128 个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyConsumerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConsumerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConsumerGroup")
	delete(f, "ConsumeEnable")
	delete(f, "ConsumeMessageOrderly")
	delete(f, "MaxRetryTimes")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConsumerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConsumerGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyConsumerGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConsumerGroupResponseParams `json:"Response"`
}

func (r *ModifyConsumerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceEndpointRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 接入点类型，
	// PUBLIC 公网
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 公网带宽，Mbps为单位
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网安全组信息
	IpRules []*IpRule `json:"IpRules,omitnil,omitempty" name:"IpRules"`

	// 公网是否按流量计费
	BillingFlow *bool `json:"BillingFlow,omitnil,omitempty" name:"BillingFlow"`
}

type ModifyInstanceEndpointRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 接入点类型，
	// PUBLIC 公网
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 公网带宽，Mbps为单位
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网安全组信息
	IpRules []*IpRule `json:"IpRules,omitnil,omitempty" name:"IpRules"`

	// 公网是否按流量计费
	BillingFlow *bool `json:"BillingFlow,omitnil,omitempty" name:"BillingFlow"`
}

func (r *ModifyInstanceEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	delete(f, "Bandwidth")
	delete(f, "IpRules")
	delete(f, "BillingFlow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceEndpointResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceEndpointResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceEndpointResponseParams `json:"Response"`
}

func (r *ModifyInstanceEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称，不能为空, 3-64个字符，只能包含数字、字母、“-”和“_”
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注信息，最多 128 个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 消息发送和接收的比例
	SendReceiveRatio *float64 `json:"SendReceiveRatio,omitnil,omitempty" name:"SendReceiveRatio"`

	// 商品规格，从 [DescribeProductSKUs](https://cloud.tencent.com/document/api/1493/107676) 接口中的 [ProductSKU](https://cloud.tencent.com/document/api/1493/96031#ProductSKU) 出参获得。
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// 消息保留时长（单位：小时），取值范围参考 [DescribeProductSKUs](https://cloud.tencent.com/document/api/1493/107676) 接口中的 [ProductSKU](https://cloud.tencent.com/document/api/1493/96031#ProductSKU) 出参：
	// 
	// - 默认值：DefaultRetention 参数
	// - 最小值：RetentionLowerLimit 参数
	// - 最大值：RetentionUpperLimit 参数
	MessageRetention *int64 `json:"MessageRetention,omitnil,omitempty" name:"MessageRetention"`

	// 是否开启弹性TPS
	ScaledTpsEnabled *bool `json:"ScaledTpsEnabled,omitnil,omitempty" name:"ScaledTpsEnabled"`

	// 是否开启ACL
	AclEnabled *bool `json:"AclEnabled,omitnil,omitempty" name:"AclEnabled"`

	// 最大可创建主题数，取值范围参考 [DescribeProductSKUs](https://cloud.tencent.com/document/api/1493/107676) 接口中的 [ProductSKU](https://cloud.tencent.com/document/api/1493/96031#ProductSKU) 出参：
	// 
	// - 最小值和默认值：TopicNumLimit 参数
	// - 最大值：TopicNumUpperLimit 参数
	MaxTopicNum *int64 `json:"MaxTopicNum,omitnil,omitempty" name:"MaxTopicNum"`

	// 免费额度之外的主题个数，免费额度参考 [DescribeProductSKUs](https://cloud.tencent.com/document/api/1493/107676) 接口中的 [ProductSKU](https://cloud.tencent.com/document/api/1493/96031#ProductSKU) 出参中的 TopicNumLimit 参数。
	ExtraTopicNum *string `json:"ExtraTopicNum,omitnil,omitempty" name:"ExtraTopicNum"`

	// 是否开启删除保护
	EnableDeletionProtection *bool `json:"EnableDeletionProtection,omitnil,omitempty" name:"EnableDeletionProtection"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称，不能为空, 3-64个字符，只能包含数字、字母、“-”和“_”
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注信息，最多 128 个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 消息发送和接收的比例
	SendReceiveRatio *float64 `json:"SendReceiveRatio,omitnil,omitempty" name:"SendReceiveRatio"`

	// 商品规格，从 [DescribeProductSKUs](https://cloud.tencent.com/document/api/1493/107676) 接口中的 [ProductSKU](https://cloud.tencent.com/document/api/1493/96031#ProductSKU) 出参获得。
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// 消息保留时长（单位：小时），取值范围参考 [DescribeProductSKUs](https://cloud.tencent.com/document/api/1493/107676) 接口中的 [ProductSKU](https://cloud.tencent.com/document/api/1493/96031#ProductSKU) 出参：
	// 
	// - 默认值：DefaultRetention 参数
	// - 最小值：RetentionLowerLimit 参数
	// - 最大值：RetentionUpperLimit 参数
	MessageRetention *int64 `json:"MessageRetention,omitnil,omitempty" name:"MessageRetention"`

	// 是否开启弹性TPS
	ScaledTpsEnabled *bool `json:"ScaledTpsEnabled,omitnil,omitempty" name:"ScaledTpsEnabled"`

	// 是否开启ACL
	AclEnabled *bool `json:"AclEnabled,omitnil,omitempty" name:"AclEnabled"`

	// 最大可创建主题数，取值范围参考 [DescribeProductSKUs](https://cloud.tencent.com/document/api/1493/107676) 接口中的 [ProductSKU](https://cloud.tencent.com/document/api/1493/96031#ProductSKU) 出参：
	// 
	// - 最小值和默认值：TopicNumLimit 参数
	// - 最大值：TopicNumUpperLimit 参数
	MaxTopicNum *int64 `json:"MaxTopicNum,omitnil,omitempty" name:"MaxTopicNum"`

	// 免费额度之外的主题个数，免费额度参考 [DescribeProductSKUs](https://cloud.tencent.com/document/api/1493/107676) 接口中的 [ProductSKU](https://cloud.tencent.com/document/api/1493/96031#ProductSKU) 出参中的 TopicNumLimit 参数。
	ExtraTopicNum *string `json:"ExtraTopicNum,omitnil,omitempty" name:"ExtraTopicNum"`

	// 是否开启删除保护
	EnableDeletionProtection *bool `json:"EnableDeletionProtection,omitnil,omitempty" name:"EnableDeletionProtection"`
}

func (r *ModifyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "Remark")
	delete(f, "SendReceiveRatio")
	delete(f, "SkuCode")
	delete(f, "MessageRetention")
	delete(f, "ScaledTpsEnabled")
	delete(f, "AclEnabled")
	delete(f, "MaxTopicNum")
	delete(f, "ExtraTopicNum")
	delete(f, "EnableDeletionProtection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceResponseParams `json:"Response"`
}

func (r *ModifyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMQTTInsPublicEndpointRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 带宽
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网访问规则
	Rules []*PublicAccessRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type ModifyMQTTInsPublicEndpointRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 带宽
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网访问规则
	Rules []*PublicAccessRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *ModifyMQTTInsPublicEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMQTTInsPublicEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Bandwidth")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMQTTInsPublicEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMQTTInsPublicEndpointResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMQTTInsPublicEndpointResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMQTTInsPublicEndpointResponseParams `json:"Response"`
}

func (r *ModifyMQTTInsPublicEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMQTTInsPublicEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMQTTInstanceCertBindingRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务端证书id
	SSLServerCertId *string `json:"SSLServerCertId,omitnil,omitempty" name:"SSLServerCertId"`

	// CA证书id
	SSLCaCertId *string `json:"SSLCaCertId,omitnil,omitempty" name:"SSLCaCertId"`
}

type ModifyMQTTInstanceCertBindingRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务端证书id
	SSLServerCertId *string `json:"SSLServerCertId,omitnil,omitempty" name:"SSLServerCertId"`

	// CA证书id
	SSLCaCertId *string `json:"SSLCaCertId,omitnil,omitempty" name:"SSLCaCertId"`
}

func (r *ModifyMQTTInstanceCertBindingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMQTTInstanceCertBindingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SSLServerCertId")
	delete(f, "SSLCaCertId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMQTTInstanceCertBindingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMQTTInstanceCertBindingResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMQTTInstanceCertBindingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMQTTInstanceCertBindingResponseParams `json:"Response"`
}

func (r *ModifyMQTTInstanceCertBindingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMQTTInstanceCertBindingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMQTTInstanceRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyMQTTInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyMQTTInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMQTTInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMQTTInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMQTTInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMQTTInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMQTTInstanceResponseParams `json:"Response"`
}

func (r *ModifyMQTTInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMQTTInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMQTTTopicRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyMQTTTopicRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyMQTTTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMQTTTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMQTTTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMQTTTopicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMQTTTopicResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMQTTTopicResponseParams `json:"Response"`
}

func (r *ModifyMQTTTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMQTTTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMQTTUserRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 是否开启消费
	PermRead *bool `json:"PermRead,omitnil,omitempty" name:"PermRead"`

	// 是否开启生产
	PermWrite *bool `json:"PermWrite,omitnil,omitempty" name:"PermWrite"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyMQTTUserRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 是否开启消费
	PermRead *bool `json:"PermRead,omitnil,omitempty" name:"PermRead"`

	// 是否开启生产
	PermWrite *bool `json:"PermWrite,omitnil,omitempty" name:"PermWrite"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyMQTTUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMQTTUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Username")
	delete(f, "PermRead")
	delete(f, "PermWrite")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMQTTUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMQTTUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMQTTUserResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMQTTUserResponseParams `json:"Response"`
}

func (r *ModifyMQTTUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMQTTUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoleRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 角色名称，从 [DescribeRoleList](https://cloud.tencent.com/document/api/1493/98862) 接口中返回的 [RoleItem](https://cloud.tencent.com/document/api/1493/96031#RoleItem) 或控制台获得。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 是否开启消费
	PermRead *bool `json:"PermRead,omitnil,omitempty" name:"PermRead"`

	// 是否开启生产
	PermWrite *bool `json:"PermWrite,omitnil,omitempty" name:"PermWrite"`

	// 权限类型，默认按集群授权（Cluster：集群维度；TopicAndGroup：主题和消费组维度）
	PermType *string `json:"PermType,omitnil,omitempty" name:"PermType"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Topic&Group维度权限配置，权限类型为 TopicAndGroup 时必填
	DetailedPerms []*DetailedRolePerm `json:"DetailedPerms,omitnil,omitempty" name:"DetailedPerms"`
}

type ModifyRoleRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 角色名称，从 [DescribeRoleList](https://cloud.tencent.com/document/api/1493/98862) 接口中返回的 [RoleItem](https://cloud.tencent.com/document/api/1493/96031#RoleItem) 或控制台获得。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 是否开启消费
	PermRead *bool `json:"PermRead,omitnil,omitempty" name:"PermRead"`

	// 是否开启生产
	PermWrite *bool `json:"PermWrite,omitnil,omitempty" name:"PermWrite"`

	// 权限类型，默认按集群授权（Cluster：集群维度；TopicAndGroup：主题和消费组维度）
	PermType *string `json:"PermType,omitnil,omitempty" name:"PermType"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Topic&Group维度权限配置，权限类型为 TopicAndGroup 时必填
	DetailedPerms []*DetailedRolePerm `json:"DetailedPerms,omitnil,omitempty" name:"DetailedPerms"`
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
	delete(f, "InstanceId")
	delete(f, "Role")
	delete(f, "PermRead")
	delete(f, "PermWrite")
	delete(f, "PermType")
	delete(f, "Remark")
	delete(f, "DetailedPerms")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRoleResponseParams struct {
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
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，从 [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030) 接口返回的 [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem) 或控制台获得。
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 队列数量，取值范围3～16
	QueueNum *int64 `json:"QueueNum,omitnil,omitempty" name:"QueueNum"`

	// 备注信息，最多 128 个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 消息保留时长（单位：小时）
	MsgTTL *int64 `json:"MsgTTL,omitnil,omitempty" name:"MsgTTL"`
}

type ModifyTopicRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，从 [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030) 接口返回的 [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem) 或控制台获得。
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 队列数量，取值范围3～16
	QueueNum *int64 `json:"QueueNum,omitnil,omitempty" name:"QueueNum"`

	// 备注信息，最多 128 个字符
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 消息保留时长（单位：小时）
	MsgTTL *int64 `json:"MsgTTL,omitnil,omitempty" name:"MsgTTL"`
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
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "QueueNum")
	delete(f, "Remark")
	delete(f, "MsgTTL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicResponseParams struct {
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

type PacketStatistics struct {
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageType *string `json:"MessageType,omitnil,omitempty" name:"MessageType"`

	// 服务质量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Qos *int64 `json:"Qos,omitnil,omitempty" name:"Qos"`

	// 指标值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type PriceTag struct {
	// 计价名称（枚举值：tps：TPS基础价；stepTps：TPS步长）
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 计费项对应的步长数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Step *int64 `json:"Step,omitnil,omitempty" name:"Step"`
}

type ProductSKU struct {
	// 产品类型，
	// EXPERIMENT，体验版
	// BASIC，基础版
	// PRO，专业版
	// PLATINUM，铂金版
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 规格代码
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// TPS上限
	TpsLimit *int64 `json:"TpsLimit,omitnil,omitempty" name:"TpsLimit"`

	// 弹性TPS上限
	ScaledTpsLimit *int64 `json:"ScaledTpsLimit,omitnil,omitempty" name:"ScaledTpsLimit"`

	// 主题数量上限默认值
	TopicNumLimit *int64 `json:"TopicNumLimit,omitnil,omitempty" name:"TopicNumLimit"`

	// 消费组数量上限
	GroupNumLimit *int64 `json:"GroupNumLimit,omitnil,omitempty" name:"GroupNumLimit"`

	// 默认消息保留时间，小时为单位
	DefaultRetention *int64 `json:"DefaultRetention,omitnil,omitempty" name:"DefaultRetention"`

	// 可调整消息保留时间上限，小时为单位
	RetentionUpperLimit *int64 `json:"RetentionUpperLimit,omitnil,omitempty" name:"RetentionUpperLimit"`

	// 可调整消息保留时间下限，小时为单位
	RetentionLowerLimit *int64 `json:"RetentionLowerLimit,omitnil,omitempty" name:"RetentionLowerLimit"`

	// 延时消息最大时长，小时为单位
	MaxMessageDelay *int64 `json:"MaxMessageDelay,omitnil,omitempty" name:"MaxMessageDelay"`

	// 是否可购买
	OnSale *bool `json:"OnSale,omitnil,omitempty" name:"OnSale"`

	// 计费项信息
	PriceTags []*PriceTag `json:"PriceTags,omitnil,omitempty" name:"PriceTags"`

	// 主题数量上限默认最大值
	TopicNumUpperLimit *int64 `json:"TopicNumUpperLimit,omitnil,omitempty" name:"TopicNumUpperLimit"`
}

type PublicAccessRule struct {
	// ip网段信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpRule *string `json:"IpRule,omitnil,omitempty" name:"IpRule"`

	// 允许或者拒绝
	// 注意：此字段可能返回 null，表示取不到有效值。
	Allow *bool `json:"Allow,omitnil,omitempty" name:"Allow"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

// Predefined struct for user
type RemoveMigratingTopicRequestParams struct {
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 主题名称，可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 命名空间，仅迁移至4.x集群有效，可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type RemoveMigratingTopicRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 主题名称，可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 命名空间，仅迁移至4.x集群有效，可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *RemoveMigratingTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveMigratingTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "TopicName")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveMigratingTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveMigratingTopicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveMigratingTopicResponse struct {
	*tchttp.BaseResponse
	Response *RemoveMigratingTopicResponseParams `json:"Response"`
}

func (r *RemoveMigratingTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveMigratingTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResendDeadLetterMessageRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 死信消息ID列表
	MessageIds []*string `json:"MessageIds,omitnil,omitempty" name:"MessageIds"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`
}

type ResendDeadLetterMessageRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 死信消息ID列表
	MessageIds []*string `json:"MessageIds,omitnil,omitempty" name:"MessageIds"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`
}

func (r *ResendDeadLetterMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResendDeadLetterMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "MessageIds")
	delete(f, "ConsumerGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResendDeadLetterMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResendDeadLetterMessageResponseParams struct {
	// 重发消息结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResendResult *bool `json:"ResendResult,omitnil,omitempty" name:"ResendResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResendDeadLetterMessageResponse struct {
	*tchttp.BaseResponse
	Response *ResendDeadLetterMessageResponseParams `json:"Response"`
}

func (r *ResendDeadLetterMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResendDeadLetterMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetConsumerGroupOffsetRequestParams struct {
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，从 [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030) 接口返回的 [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem) 或控制台获得。
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 重置位点的时间戳（单位：毫秒），指定为 -1 时表示重置到最新位点
	ResetTimestamp *int64 `json:"ResetTimestamp,omitnil,omitempty" name:"ResetTimestamp"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`
}

type ResetConsumerGroupOffsetRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云 RocketMQ 实例 ID，从 [DescribeFusionInstanceList](https://cloud.tencent.com/document/api/1493/106745) 接口或控制台获得。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称，从 [DescribeTopicList](https://cloud.tencent.com/document/api/1493/96030) 接口返回的 [TopicItem](https://cloud.tencent.com/document/api/1493/96031#TopicItem) 或控制台获得。
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 重置位点的时间戳（单位：毫秒），指定为 -1 时表示重置到最新位点
	ResetTimestamp *int64 `json:"ResetTimestamp,omitnil,omitempty" name:"ResetTimestamp"`

	// 消费组名称，从 [DescribeConsumerGroupList](https://cloud.tencent.com/document/api/1493/101535) 接口返回的 [ConsumeGroupItem](https://cloud.tencent.com/document/api/1493/96031#ConsumeGroupItem) 或控制台获得。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`
}

func (r *ResetConsumerGroupOffsetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetConsumerGroupOffsetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "ResetTimestamp")
	delete(f, "ConsumerGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetConsumerGroupOffsetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetConsumerGroupOffsetResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetConsumerGroupOffsetResponse struct {
	*tchttp.BaseResponse
	Response *ResetConsumerGroupOffsetResponseParams `json:"Response"`
}

func (r *ResetConsumerGroupOffsetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetConsumerGroupOffsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RoleItem struct {
	// 角色名称
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 是否开启消费
	PermRead *bool `json:"PermRead,omitnil,omitempty" name:"PermRead"`

	// 是否开启生产
	PermWrite *bool `json:"PermWrite,omitnil,omitempty" name:"PermWrite"`

	// Access Key
	AccessKey *string `json:"AccessKey,omitnil,omitempty" name:"AccessKey"`

	// Secret Key
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 角色的创建时间，**Unix时间戳（毫秒）**
	CreatedTime *int64 `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 角色的更新时间，**Unix时间戳（毫秒）**
	ModifiedTime *int64 `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// 权限类型，默认按集群授权（Cluster：集群级别；TopicAndGroup：主题&消费组级别）
	PermType *string `json:"PermType,omitnil,omitempty" name:"PermType"`

	// Topic和Group维度权限配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetailedRolePerms []*DetailedRolePerm `json:"DetailedRolePerms,omitnil,omitempty" name:"DetailedRolePerms"`
}

// Predefined struct for user
type RollbackMigratingTopicStageRequestParams struct {
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 主题名称，可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 命名空间，仅迁移至4.x集群有效，可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type RollbackMigratingTopicStageRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID，可在[DescribeSmoothMigrationTaskList](https://cloud.tencent.com/document/api/1493/119997)接口返回的[SmoothMigrationTaskItem](https://cloud.tencent.com/document/api/1493/96031#SmoothMigrationTaskItem)或控制台中获得。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 主题名称，可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 命名空间，仅迁移至4.x集群有效，可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

func (r *RollbackMigratingTopicStageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackMigratingTopicStageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "TopicName")
	delete(f, "Namespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackMigratingTopicStageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackMigratingTopicStageResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RollbackMigratingTopicStageResponse struct {
	*tchttp.BaseResponse
	Response *RollbackMigratingTopicStageResponseParams `json:"Response"`
}

func (r *RollbackMigratingTopicStageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackMigratingTopicStageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SmoothMigrationTaskItem struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称	
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 源集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceClusterName *string `json:"SourceClusterName,omitnil,omitempty" name:"SourceClusterName"`

	// 目标集群实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 网络连接类型， 
	// PUBLIC 公网 
	// VPC 私有网络 
	// OTHER 其他
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectionType *string `json:"ConnectionType,omitnil,omitempty" name:"ConnectionType"`

	// 源集群NameServer地址	
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceNameServer *string `json:"SourceNameServer,omitnil,omitempty" name:"SourceNameServer"`

	// 任务状态:
	// Configuration 迁移配置,
	// SourceConnecting 连接源集群中,
	//  MetaDataImport 元数据导入,
	// EndpointSetup 切换接入点,
	// ServiceMigration 切流中,
	// Completed 已完成,
	// Cancelled 已取消
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 目标集群实例版本，
	// 4 表示4.x版本
	// 5 表示5.x版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceVersion *string `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`
}

type SourceClusterGroupConfig struct {
	// 消费组名称，可在[DescribeSourceClusterGroupList](https://cloud.tencent.com/document/api/1493/118006)接口返回的[SourceClusterGroupConfig](https://cloud.tencent.com/document/api/1493/96031#SourceClusterGroupConfig)数据中获取。
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否已导入，作为入参时无效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Imported *bool `json:"Imported,omitnil,omitempty" name:"Imported"`

	// 命名空间，仅4.x集群有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 导入状态
	// Unknown 未知
	// Success 成功
	// Failure 失败
	// AlreadyExists 已存在
	// 
	// 仅作为出参时使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImportStatus *string `json:"ImportStatus,omitnil,omitempty" name:"ImportStatus"`

	// 4.x的命名空间，出参使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceV4 *string `json:"NamespaceV4,omitnil,omitempty" name:"NamespaceV4"`

	// 4.x的消费组名，出参使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupNameV4 *string `json:"GroupNameV4,omitnil,omitempty" name:"GroupNameV4"`

	// 4.x的完整命名空间，出参使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullNamespaceV4 *string `json:"FullNamespaceV4,omitnil,omitempty" name:"FullNamespaceV4"`

	// 是否为顺序投递，5.0有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumeMessageOrderly *bool `json:"ConsumeMessageOrderly,omitnil,omitempty" name:"ConsumeMessageOrderly"`
}

type SourceClusterTopicConfig struct {
	// 主题名称，可在[DescribeMigratingTopicList](https://cloud.tencent.com/document/api/1493/118007)接口返回的[MigratingTopic](https://cloud.tencent.com/document/api/1493/96031#MigratingTopic)数据结构中获得。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 主题类型，
	// 5.x版本
	// UNSPECIFIED 未指定
	// NORMAL 普通消息
	// FIFO 顺序消息
	// DELAY 延迟消息
	// TRANSACTION 事务消息
	// 
	// 4.x版本
	// Normal 普通消息
	// PartitionedOrder 分区顺序消息
	// Transaction 事务消息
	// DelayScheduled 延时消息
	TopicType *string `json:"TopicType,omitnil,omitempty" name:"TopicType"`

	// 队列数
	QueueNum *int64 `json:"QueueNum,omitnil,omitempty" name:"QueueNum"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否已导入，作为入参时无效
	Imported *bool `json:"Imported,omitnil,omitempty" name:"Imported"`

	// 命名空间，仅4.x集群有效
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 导入状态，
	// Unknown 未知，
	// AlreadyExists 已存在，
	// Success 成功，
	// Failure 失败
	// 
	// 仅作为出参可用
	ImportStatus *string `json:"ImportStatus,omitnil,omitempty" name:"ImportStatus"`

	// 4.x的命名空间，出参使用
	NamespaceV4 *string `json:"NamespaceV4,omitnil,omitempty" name:"NamespaceV4"`

	// 4.x的主题名，出参使用
	TopicNameV4 *string `json:"TopicNameV4,omitnil,omitempty" name:"TopicNameV4"`

	// 4.x的完整命名空间，出参使用
	FullNamespaceV4 *string `json:"FullNamespaceV4,omitnil,omitempty" name:"FullNamespaceV4"`
}

type StatisticsReport struct {
	// 字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bytes *int64 `json:"Bytes,omitnil,omitempty" name:"Bytes"`

	// 监控指标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*PacketStatistics `json:"Items,omitnil,omitempty" name:"Items"`
}

type SubscriptionData struct {
	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 主题类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicType *string `json:"TopicType,omitnil,omitempty" name:"TopicType"`

	// 单个节点上主题队列数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicQueueNum *int64 `json:"TopicQueueNum,omitnil,omitempty" name:"TopicQueueNum"`

	// 消费组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerGroup *string `json:"ConsumerGroup,omitnil,omitempty" name:"ConsumerGroup"`

	// 是否在线
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsOnline *bool `json:"IsOnline,omitnil,omitempty" name:"IsOnline"`

	// 消费类型，枚举值如下：
	// 
	// - PULL：PULL 消费类型
	// - PUSH：PUSH 消费类型
	// - POP：POP 消费类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumeType *string `json:"ConsumeType,omitnil,omitempty" name:"ConsumeType"`

	// 订阅规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubString *string `json:"SubString,omitnil,omitempty" name:"SubString"`

	// 过滤类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpressionType *string `json:"ExpressionType,omitnil,omitempty" name:"ExpressionType"`

	// 订阅一致性，枚举如下：
	// 
	// - 0: 订阅一致
	// - 1: 订阅不一致
	// 注意：此字段可能返回 null，表示取不到有效值。
	Consistency *int64 `json:"Consistency,omitnil,omitempty" name:"Consistency"`

	// 消费堆积
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerLag *int64 `json:"ConsumerLag,omitnil,omitempty" name:"ConsumerLag"`

	// 最后消费进度更新时间，**Unix时间戳（毫秒）**
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *int64 `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// 最大重试次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxRetryTimes *int64 `json:"MaxRetryTimes,omitnil,omitempty" name:"MaxRetryTimes"`

	// 是否顺序消费
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumeMessageOrderly *bool `json:"ConsumeMessageOrderly,omitnil,omitempty" name:"ConsumeMessageOrderly"`

	// 消费模式: 
	// BROADCASTING 广播模式;
	// CLUSTERING 集群模式;
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageModel *string `json:"MessageModel,omitnil,omitempty" name:"MessageModel"`

	// 订阅不一致的客户端列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientSubscriptionInfos []*ClientSubscriptionInfo `json:"ClientSubscriptionInfos,omitnil,omitempty" name:"ClientSubscriptionInfos"`
}

type Tag struct {
	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TagFilter struct {
	// 标签键名称
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值列表
	TagValues []*string `json:"TagValues,omitnil,omitempty" name:"TagValues"`
}

type TopicConsumeStats struct {
	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 主题类型，枚举值如下：
	// 
	// - UNSPECIFIED：未指定
	// - NORMAL：普通消息
	// - FIFO：顺序消息
	// - DELAY：延时消息
	// - TRANSACTION：事务消息
	TopicType *string `json:"TopicType,omitnil,omitempty" name:"TopicType"`

	// 单节点主题队列数量
	QueueNum *int64 `json:"QueueNum,omitnil,omitempty" name:"QueueNum"`

	// 消费堆积
	ConsumerLag *int64 `json:"ConsumerLag,omitnil,omitempty" name:"ConsumerLag"`

	// 订阅规则，`*`表示订阅全部TAG
	SubString *string `json:"SubString,omitnil,omitempty" name:"SubString"`

	// 最后消费进度更新时间，**Unix时间戳（毫秒）**
	LastUpdateTime *int64 `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`
}

type TopicItem struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 主题类型
	// NORMAL:普通消息,
	// FIFO:顺序消息,
	// DELAY:延时消息,
	// TRANSACTION:事务消息
	TopicType *string `json:"TopicType,omitnil,omitempty" name:"TopicType"`

	// 队列数量
	QueueNum *int64 `json:"QueueNum,omitnil,omitempty" name:"QueueNum"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 4.x的集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterIdV4 *string `json:"ClusterIdV4,omitnil,omitempty" name:"ClusterIdV4"`

	// 4.x的命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	NamespaceV4 *string `json:"NamespaceV4,omitnil,omitempty" name:"NamespaceV4"`

	// 4.x的主题名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicV4 *string `json:"TopicV4,omitnil,omitempty" name:"TopicV4"`

	// 4.x的完整命名空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullNamespaceV4 *string `json:"FullNamespaceV4,omitnil,omitempty" name:"FullNamespaceV4"`

	// 消息保留时长
	MsgTTL *int64 `json:"MsgTTL,omitnil,omitempty" name:"MsgTTL"`
}

type TopicStageChangeResult struct {
	// 主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 是否成功
	Success *bool `json:"Success,omitnil,omitempty" name:"Success"`

	// 命名空间，仅4.x有效
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`
}

type VpcInfo struct {
	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}