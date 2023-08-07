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

package v20230308

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type DescribeInstanceListRequestParams struct {
	// 查询起始位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 标签过滤器
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
}

type DescribeInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 标签过滤器
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
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
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceListResponseParams struct {
	// 查询总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例列表
	Data []*InstanceItem `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeTopicListRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeTopicListRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicListResponseParams struct {
	// 查询总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 主题列表
	Data []*TopicItem `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type Filter struct {
	// 过滤条件名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤条件的值
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type InstanceItem struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例版本
	Version *string `json:"Version,omitempty" name:"Version"`

	// 实例类型，
	// EXPERIMENT，体验版
	// BASIC，基础版
	// PRO，专业版
	// PLATINUM，铂金版
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

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
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`

	// 实例主题数上限
	TopicNumLimit *int64 `json:"TopicNumLimit,omitempty" name:"TopicNumLimit"`

	// 实例消费组数量上限
	GroupNumLimit *int64 `json:"GroupNumLimit,omitempty" name:"GroupNumLimit"`

	// 计费模式，
	// POSTPAID，按量计费
	// PREPAID，包年包月
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 到期时间，秒为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiryTime *int64 `json:"ExpiryTime,omitempty" name:"ExpiryTime"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 主题数量
	TopicNum *int64 `json:"TopicNum,omitempty" name:"TopicNum"`

	// 消费组数量
	GroupNum *int64 `json:"GroupNum,omitempty" name:"GroupNum"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*Tag `json:"TagList,omitempty" name:"TagList"`

	// 商品规格
	SkuCode *string `json:"SkuCode,omitempty" name:"SkuCode"`

	// TPS限流值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TpsLimit *int64 `json:"TpsLimit,omitempty" name:"TpsLimit"`

	// 弹性TPS限流值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaledTpsLimit *int64 `json:"ScaledTpsLimit,omitempty" name:"ScaledTpsLimit"`

	// 消息保留时间，小时为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageRetention *int64 `json:"MessageRetention,omitempty" name:"MessageRetention"`

	// 延迟消息最大时长，小时为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxMessageDelay *int64 `json:"MaxMessageDelay,omitempty" name:"MaxMessageDelay"`
}

type Tag struct {
	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagFilter struct {
	// 标签键名称
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值列表
	TagValues []*string `json:"TagValues,omitempty" name:"TagValues"`
}

type TopicItem struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// 主题类型
	TopicType *string `json:"TopicType,omitempty" name:"TopicType"`

	// 队列数量
	QueueNum *int64 `json:"QueueNum,omitempty" name:"QueueNum"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}