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

package v20210622

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ApmAgentInfo struct {

	// Agent下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentDownloadURL *string `json:"AgentDownloadURL,omitempty" name:"AgentDownloadURL"`

	// Collector上报地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	CollectorURL *string `json:"CollectorURL,omitempty" name:"CollectorURL"`

	// Token信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Token *string `json:"Token,omitempty" name:"Token"`

	// 外网上报地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicCollectorURL *string `json:"PublicCollectorURL,omitempty" name:"PublicCollectorURL"`

	// 自研VPC上报地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerCollectorURL *string `json:"InnerCollectorURL,omitempty" name:"InnerCollectorURL"`

	// 内网上报地址(Private Link上报地址)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateLinkCollectorURL *string `json:"PrivateLinkCollectorURL,omitempty" name:"PrivateLinkCollectorURL"`
}

type ApmInstanceDetail struct {

	// 存储使用量(MB)
	// 注意：此字段可能返回 null，表示取不到有效值。
	AmountOfUsedStorage *float64 `json:"AmountOfUsedStorage,omitempty" name:"AmountOfUsedStorage"`

	// 实例名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 实例所属tag列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 创建人Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`

	// 该实例已上报的服务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceCount *int64 `json:"ServiceCount,omitempty" name:"ServiceCount"`

	// 日均上报Span数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountOfReportSpanPerDay *int64 `json:"CountOfReportSpanPerDay,omitempty" name:"CountOfReportSpanPerDay"`

	// AppId信息
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// Trace数据保存时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	TraceDuration *int64 `json:"TraceDuration,omitempty" name:"TraceDuration"`

	// 实例描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例所属地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 实例上报额度
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpanDailyCounters *int64 `json:"SpanDailyCounters,omitempty" name:"SpanDailyCounters"`

	// 实例是否开通计费
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingInstance *int64 `json:"BillingInstance,omitempty" name:"BillingInstance"`

	// 错误率阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrRateThreshold *int64 `json:"ErrRateThreshold,omitempty" name:"ErrRateThreshold"`

	// 采样率阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleRate *int64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// 是否开启错误采样 0  关 1 开
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorSample *int64 `json:"ErrorSample,omitempty" name:"ErrorSample"`

	// 慢调用保存阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlowRequestSavedThreshold *int64 `json:"SlowRequestSavedThreshold,omitempty" name:"SlowRequestSavedThreshold"`
}

type ApmTag struct {

	// 维度Key(列名，标签Key)
	Key *string `json:"Key,omitempty" name:"Key"`

	// 维度值（标签值）
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeApmAgentRequest struct {
	*tchttp.BaseRequest

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 接入方式
	AgentType *string `json:"AgentType,omitempty" name:"AgentType"`

	// 环境
	NetworkMode *string `json:"NetworkMode,omitempty" name:"NetworkMode"`

	// 语言
	LanguageEnvironment *string `json:"LanguageEnvironment,omitempty" name:"LanguageEnvironment"`
}

func (r *DescribeApmAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AgentType")
	delete(f, "NetworkMode")
	delete(f, "LanguageEnvironment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmAgentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Agent信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		ApmAgent *ApmAgentInfo `json:"ApmAgent,omitempty" name:"ApmAgent"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmInstancesRequest struct {
	*tchttp.BaseRequest

	// Tag列表
	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`

	// 搜索实例名
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 过滤实例ID
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeApmInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tags")
	delete(f, "InstanceName")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// apm实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Instances []*ApmInstanceDetail `json:"Instances,omitempty" name:"Instances"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
