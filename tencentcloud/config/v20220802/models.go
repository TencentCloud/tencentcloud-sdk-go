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

package v20220802

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Annotation struct {
	// 资源当前实际配置。长度为0~256位字符，即资源不合规配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Configuration *string `json:"Configuration,omitnil" name:"Configuration"`

	// 资源期望配置。长度为0~256位字符，即资源合规配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	DesiredValue *string `json:"DesiredValue,omitnil" name:"DesiredValue"`

	// 资源当前配置和期望配置之间的比较运算符。长度为0~16位字符，自定义规则上报评估结果此字段可能为空
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 当前配置在资源属性结构体中的JSON路径。长度为0~256位字符，自定义规则上报评估结果此字段可能为空
	Property *string `json:"Property,omitnil" name:"Property"`
}

type ConfigRule struct {
	// 规则标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Identifier *string `json:"Identifier,omitnil" name:"Identifier"`

	// 规则名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 规则参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputParameter []*InputParameter `json:"InputParameter,omitnil" name:"InputParameter"`

	// 规则触发条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceCondition []*SourceConditionForManage `json:"SourceCondition,omitnil" name:"SourceCondition"`

	// 规则支持的资源类型，规则仅对指定资源类型的资源生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType []*string `json:"ResourceType,omitnil" name:"ResourceType"`

	// 规则所属标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*string `json:"Labels,omitnil" name:"Labels"`

	// 规则风险等级
	// 1:低风险
	// 2:中风险
	// 3:高风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *int64 `json:"RiskLevel,omitnil" name:"RiskLevel"`

	// 规则对应的函数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceFunction *string `json:"ServiceFunction,omitnil" name:"ServiceFunction"`

	// 创建时间
	// 格式：YYYY-MM-DD h:i:s
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 规则描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// ACTIVE：启用
	// NO_ACTIVE：停止
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 合规： 'COMPLIANT'
	// 不合规： 'NON_COMPLIANT'
	// 无法应用规则： 'NOT_APPLICABLE'
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplianceResult *string `json:"ComplianceResult,omitnil" name:"ComplianceResult"`

	// ["",""]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Annotation *Annotation `json:"Annotation,omitnil" name:"Annotation"`

	// 规则评估时间
	// 格式：YYYY-MM-DD h:i:s
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigRuleInvokedTime *string `json:"ConfigRuleInvokedTime,omitnil" name:"ConfigRuleInvokedTime"`

	// 规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigRuleId *string `json:"ConfigRuleId,omitnil" name:"ConfigRuleId"`

	// CUSTOMIZE：自定义规则、
	// SYSTEM：托管规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentifierType *string `json:"IdentifierType,omitnil" name:"IdentifierType"`

	// 合规包ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompliancePackId *string `json:"CompliancePackId,omitnil" name:"CompliancePackId"`

	// 触发类型
	// ScheduledNotification：周期触发、
	// ConfigurationItemChangeNotification：变更触发
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerType []*TriggerType `json:"TriggerType,omitnil" name:"TriggerType"`

	// 参数详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManageInputParameter []*InputParameterForManage `json:"ManageInputParameter,omitnil" name:"ManageInputParameter"`

	// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompliancePackName *string `json:"CompliancePackName,omitnil" name:"CompliancePackName"`

	// 关联地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionsScope []*string `json:"RegionsScope,omitnil" name:"RegionsScope"`

	// 关联标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagsScope []*Tag `json:"TagsScope,omitnil" name:"TagsScope"`

	//  规则对指定资源ID无效，即不对该资源执行评估。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludeResourceIdsScope []*string `json:"ExcludeResourceIdsScope,omitnil" name:"ExcludeResourceIdsScope"`

	// 账号组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountGroupId *string `json:"AccountGroupId,omitnil" name:"AccountGroupId"`

	// 账号组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountGroupName *string `json:"AccountGroupName,omitnil" name:"AccountGroupName"`

	// 规则所属用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleOwnerId *uint64 `json:"RuleOwnerId,omitnil" name:"RuleOwnerId"`

	// 预设规则支持的触发方式
	// ScheduledNotification：周期触发
	// ConfigurationItemChangeNotification：变更触发
	ManageTriggerType []*string `json:"ManageTriggerType,omitnil" name:"ManageTriggerType"`
}

type InputParameter struct {
	// 参数名
	ParameterKey *string `json:"ParameterKey,omitnil" name:"ParameterKey"`

	// 参数类型。必填类型：Require，可选类型：Optional。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 参数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type InputParameterForManage struct {
	// 值类型。数值：Integer， 字符串：String
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueType *string `json:"ValueType,omitnil" name:"ValueType"`

	// 参数Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParameterKey *string `json:"ParameterKey,omitnil" name:"ParameterKey"`

	// 参数类型。必填类型：Require，可选类型：Optional。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 默认值
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultValue *string `json:"DefaultValue,omitnil" name:"DefaultValue"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`
}

// Predefined struct for user
type ListAggregateConfigRulesRequestParams struct {
	// 每页限制
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil" name:"AccountGroupId"`

	// 排序类型, 倒序：desc，顺序：asc
	OrderType *string `json:"OrderType,omitnil" name:"OrderType"`

	// 风险等级
	// 1：高风险。
	// 2：中风险。
	// 3：低风险。
	RiskLevel []*uint64 `json:"RiskLevel,omitnil" name:"RiskLevel"`

	// 规则状态
	State *string `json:"State,omitnil" name:"State"`

	// 评估结果
	ComplianceResult []*string `json:"ComplianceResult,omitnil" name:"ComplianceResult"`

	// 规则名
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 规则所属账号ID
	RuleOwnerId *uint64 `json:"RuleOwnerId,omitnil" name:"RuleOwnerId"`
}

type ListAggregateConfigRulesRequest struct {
	*tchttp.BaseRequest
	
	// 每页限制
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil" name:"AccountGroupId"`

	// 排序类型, 倒序：desc，顺序：asc
	OrderType *string `json:"OrderType,omitnil" name:"OrderType"`

	// 风险等级
	// 1：高风险。
	// 2：中风险。
	// 3：低风险。
	RiskLevel []*uint64 `json:"RiskLevel,omitnil" name:"RiskLevel"`

	// 规则状态
	State *string `json:"State,omitnil" name:"State"`

	// 评估结果
	ComplianceResult []*string `json:"ComplianceResult,omitnil" name:"ComplianceResult"`

	// 规则名
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 规则所属账号ID
	RuleOwnerId *uint64 `json:"RuleOwnerId,omitnil" name:"RuleOwnerId"`
}

func (r *ListAggregateConfigRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAggregateConfigRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "AccountGroupId")
	delete(f, "OrderType")
	delete(f, "RiskLevel")
	delete(f, "State")
	delete(f, "ComplianceResult")
	delete(f, "RuleName")
	delete(f, "RuleOwnerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAggregateConfigRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAggregateConfigRulesResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 详情
	Items []*ConfigRule `json:"Items,omitnil" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListAggregateConfigRulesResponse struct {
	*tchttp.BaseResponse
	Response *ListAggregateConfigRulesResponseParams `json:"Response"`
}

func (r *ListAggregateConfigRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAggregateConfigRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListConfigRulesRequestParams struct {
	// 每页限制
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序类型, 倒序：desc，顺序：asc
	OrderType *string `json:"OrderType,omitnil" name:"OrderType"`

	// 风险等级
	// 1：高风险。
	// 2：中风险。
	// 3：低风险。
	RiskLevel []*uint64 `json:"RiskLevel,omitnil" name:"RiskLevel"`

	// 规则状态
	State *string `json:"State,omitnil" name:"State"`

	// 评估结果
	ComplianceResult []*string `json:"ComplianceResult,omitnil" name:"ComplianceResult"`

	// 规则名
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`
}

type ListConfigRulesRequest struct {
	*tchttp.BaseRequest
	
	// 每页限制
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 排序类型, 倒序：desc，顺序：asc
	OrderType *string `json:"OrderType,omitnil" name:"OrderType"`

	// 风险等级
	// 1：高风险。
	// 2：中风险。
	// 3：低风险。
	RiskLevel []*uint64 `json:"RiskLevel,omitnil" name:"RiskLevel"`

	// 规则状态
	State *string `json:"State,omitnil" name:"State"`

	// 评估结果
	ComplianceResult []*string `json:"ComplianceResult,omitnil" name:"ComplianceResult"`

	// 规则名
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`
}

func (r *ListConfigRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListConfigRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderType")
	delete(f, "RiskLevel")
	delete(f, "State")
	delete(f, "ComplianceResult")
	delete(f, "RuleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListConfigRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListConfigRulesResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 详情
	Items []*ConfigRule `json:"Items,omitnil" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ListConfigRulesResponse struct {
	*tchttp.BaseResponse
	Response *ListConfigRulesResponseParams `json:"Response"`
}

func (r *ListConfigRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListConfigRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SourceConditionForManage struct {
	// 条件为空，合规：COMPLIANT，不合规：NON_COMPLIANT，无法应用：NOT_APPLICABLE
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmptyAs *string `json:"EmptyAs,omitnil" name:"EmptyAs"`

	// 配置路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelectPath *string `json:"SelectPath,omitnil" name:"SelectPath"`

	// 操作运算符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 是否必须
	// 注意：此字段可能返回 null，表示取不到有效值。
	Required *bool `json:"Required,omitnil" name:"Required"`

	// 期望值
	// 注意：此字段可能返回 null，表示取不到有效值。
	DesiredValue *string `json:"DesiredValue,omitnil" name:"DesiredValue"`
}

type Tag struct {
	// 标签key
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签value
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type TriggerType struct {
	// 触发类型
	MessageType *string `json:"MessageType,omitnil" name:"MessageType"`

	// 触发时间周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitnil" name:"MaximumExecutionFrequency"`
}