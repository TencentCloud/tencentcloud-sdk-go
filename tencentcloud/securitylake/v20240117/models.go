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

package v20240117

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type DescribeSecurityAlarmTableListRequestParams struct {
	// 实例ID
	SdlId *string `json:"SdlId,omitnil,omitempty" name:"SdlId"`

	// 过滤条件
	Filters []*WebSearchFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 长度
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 开始时间,毫秒
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeSecurityAlarmTableListRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	SdlId *string `json:"SdlId,omitnil,omitempty" name:"SdlId"`

	// 过滤条件
	Filters []*WebSearchFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 长度
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 开始时间,毫秒
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeSecurityAlarmTableListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAlarmTableListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdlId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityAlarmTableListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityAlarmTableListResponseParams struct {
	// 字段列表
	AlarmList []*SecurityAlarmTable `json:"AlarmList,omitnil,omitempty" name:"AlarmList"`

	// 数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityAlarmTableListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityAlarmTableListResponseParams `json:"Response"`
}

func (r *DescribeSecurityAlarmTableListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAlarmTableListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityAlarmTable struct {
	// 时间
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 告警名称
	AlarmName *string `json:"AlarmName,omitnil,omitempty" name:"AlarmName"`

	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 告警id
	AlarmId *int64 `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// 安全性
	Severity *int64 `json:"Severity,omitnil,omitempty" name:"Severity"`

	// 评分
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 分类
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 子分类
	SubCategory *string `json:"SubCategory,omitnil,omitempty" name:"SubCategory"`

	// 标签
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 有效载荷
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// 结果
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 可信度
	Confidence *int64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 规则主题
	RuleTopic *string `json:"RuleTopic,omitnil,omitempty" name:"RuleTopic"`

	// 处理时间
	HandleTime *string `json:"HandleTime,omitnil,omitempty" name:"HandleTime"`

	// 建议
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 来源名称
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// APPID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 规则ID
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 事件时间
	EventTime *string `json:"EventTime,omitnil,omitempty" name:"EventTime"`

	// 规则类型
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 攻击次数
	AttackNum *int64 `json:"AttackNum,omitnil,omitempty" name:"AttackNum"`

	// 告警数量
	AlarmCount *int64 `json:"AlarmCount,omitnil,omitempty" name:"AlarmCount"`

	// ATT&CK子技术
	AttackSubTechnique *string `json:"AttackSubTechnique,omitnil,omitempty" name:"AttackSubTechnique"`

	// ATT&CK技术
	AttackTechnique *string `json:"AttackTechnique,omitnil,omitempty" name:"AttackTechnique"`

	// ATT&CK战术
	AttackTactic *string `json:"AttackTactic,omitnil,omitempty" name:"AttackTactic"`

	// ATT&CK子技术名称
	AttackSubTechniqueName *string `json:"AttackSubTechniqueName,omitnil,omitempty" name:"AttackSubTechniqueName"`

	// ATT&CK技术名称
	AttackTechniqueName *string `json:"AttackTechniqueName,omitnil,omitempty" name:"AttackTechniqueName"`

	// 凭证访问
	AttackTacticName *string `json:"AttackTacticName,omitnil,omitempty" name:"AttackTacticName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 规则表达式
	RuleExpression *string `json:"RuleExpression,omitnil,omitempty" name:"RuleExpression"`

	// 表达式类型
	ExpressionType *string `json:"ExpressionType,omitnil,omitempty" name:"ExpressionType"`

	// 下钻表达式
	DrillDownExpression *string `json:"DrillDownExpression,omitnil,omitempty" name:"DrillDownExpression"`

	// 源IP
	SrcIp *string `json:"SrcIp,omitnil,omitempty" name:"SrcIp"`

	// 源端口
	SrcPort *int64 `json:"SrcPort,omitnil,omitempty" name:"SrcPort"`

	// 目的IP
	DstIp *string `json:"DstIp,omitnil,omitempty" name:"DstIp"`

	// 目的端口
	DstPort *int64 `json:"DstPort,omitnil,omitempty" name:"DstPort"`

	// 主机IP
	HostIp *string `json:"HostIp,omitnil,omitempty" name:"HostIp"`

	// 主机资产
	HostAsset *string `json:"HostAsset,omitnil,omitempty" name:"HostAsset"`

	// 实例id
	SdlId *string `json:"SdlId,omitnil,omitempty" name:"SdlId"`

	// 自定义富化字段信息
	RichCustomInfos []*string `json:"RichCustomInfos,omitnil,omitempty" name:"RichCustomInfos"`

	// 攻击者ip
	AttackerIp *string `json:"AttackerIp,omitnil,omitempty" name:"AttackerIp"`

	// 攻击者资产ID
	AttackerAsset *string `json:"AttackerAsset,omitnil,omitempty" name:"AttackerAsset"`

	// 受害者ip
	VictimIp *string `json:"VictimIp,omitnil,omitempty" name:"VictimIp"`

	// 受害者资产ID
	VictimAsset *string `json:"VictimAsset,omitnil,omitempty" name:"VictimAsset"`

	// 攻击方向
	AttackDirection *string `json:"AttackDirection,omitnil,omitempty" name:"AttackDirection"`

	// 流量方向
	TrafficDirection *string `json:"TrafficDirection,omitnil,omitempty" name:"TrafficDirection"`

	// 测试
	SecurityGroupAlertInfos []*SecurityGroupAlertInfo `json:"SecurityGroupAlertInfos,omitnil,omitempty" name:"SecurityGroupAlertInfos"`
}

type SecurityGroupAlertInfo struct {
	// 告警Uuid
	AlarmUuid *string `json:"AlarmUuid,omitnil,omitempty" name:"AlarmUuid"`

	// 告警生成时间
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`
}

type WebSearchFilter struct {
	// 过滤字段
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 是否全匹配
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`
}