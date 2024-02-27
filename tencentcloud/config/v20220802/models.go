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
	Configuration *string `json:"Configuration,omitnil,omitempty" name:"Configuration"`

	// 资源期望配置。长度为0~256位字符，即资源合规配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	DesiredValue *string `json:"DesiredValue,omitnil,omitempty" name:"DesiredValue"`

	// 资源当前配置和期望配置之间的比较运算符。长度为0~16位字符，自定义规则上报评估结果此字段可能为空
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 当前配置在资源属性结构体中的JSON路径。长度为0~256位字符，自定义规则上报评估结果此字段可能为空
	Property *string `json:"Property,omitnil,omitempty" name:"Property"`
}

type ConfigRule struct {
	// 规则标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// 规则名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputParameter []*InputParameter `json:"InputParameter,omitnil,omitempty" name:"InputParameter"`

	// 规则触发条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceCondition []*SourceConditionForManage `json:"SourceCondition,omitnil,omitempty" name:"SourceCondition"`

	// 规则支持的资源类型，规则仅对指定资源类型的资源生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType []*string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 规则所属标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*string `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 规则风险等级
	// 1:低风险
	// 2:中风险
	// 3:高风险
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 规则对应的函数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceFunction *string `json:"ServiceFunction,omitnil,omitempty" name:"ServiceFunction"`

	// 创建时间
	// 格式：YYYY-MM-DD h:i:s
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 规则描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// ACTIVE：启用
	// NO_ACTIVE：停止
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 合规： 'COMPLIANT'
	// 不合规： 'NON_COMPLIANT'
	// 无法应用规则： 'NOT_APPLICABLE'
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplianceResult *string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// ["",""]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Annotation *Annotation `json:"Annotation,omitnil,omitempty" name:"Annotation"`

	// 规则评估时间
	// 格式：YYYY-MM-DD h:i:s
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigRuleInvokedTime *string `json:"ConfigRuleInvokedTime,omitnil,omitempty" name:"ConfigRuleInvokedTime"`

	// 规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigRuleId *string `json:"ConfigRuleId,omitnil,omitempty" name:"ConfigRuleId"`

	// CUSTOMIZE：自定义规则、
	// SYSTEM：托管规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdentifierType *string `json:"IdentifierType,omitnil,omitempty" name:"IdentifierType"`

	// 合规包ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 触发类型
	// ScheduledNotification：周期触发、
	// ConfigurationItemChangeNotification：变更触发
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerType []*TriggerType `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 参数详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManageInputParameter []*InputParameterForManage `json:"ManageInputParameter,omitnil,omitempty" name:"ManageInputParameter"`

	// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// 关联地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionsScope []*string `json:"RegionsScope,omitnil,omitempty" name:"RegionsScope"`

	// 关联标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagsScope []*Tag `json:"TagsScope,omitnil,omitempty" name:"TagsScope"`

	//  规则对指定资源ID无效，即不对该资源执行评估。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludeResourceIdsScope []*string `json:"ExcludeResourceIdsScope,omitnil,omitempty" name:"ExcludeResourceIdsScope"`

	// 账号组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 账号组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountGroupName *string `json:"AccountGroupName,omitnil,omitempty" name:"AccountGroupName"`

	// 规则所属用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleOwnerId *uint64 `json:"RuleOwnerId,omitnil,omitempty" name:"RuleOwnerId"`

	// 预设规则支持的触发方式
	// ScheduledNotification：周期触发
	// ConfigurationItemChangeNotification：变更触发
	ManageTriggerType []*string `json:"ManageTriggerType,omitnil,omitempty" name:"ManageTriggerType"`
}

type Evaluation struct {
	// 已评估资源ID。长度为0~256个字符
	ComplianceResourceId *string `json:"ComplianceResourceId,omitnil,omitempty" name:"ComplianceResourceId"`

	// 已评估资源类型。
	// 支持:
	// QCS::CVM::Instance、 QCS::CBS::Disk、QCS::VPC::Vpc、QCS::VPC::Subnet、QCS::VPC::SecurityGroup、 QCS::CAM::User、QCS::CAM::Group、QCS::CAM::Policy、QCS::CAM::Role、QCS::COS::Bucket
	ComplianceResourceType *string `json:"ComplianceResourceType,omitnil,omitempty" name:"ComplianceResourceType"`

	// 已评估资源地域。
	// 长度为0~32个字符
	ComplianceRegion *string `json:"ComplianceRegion,omitnil,omitempty" name:"ComplianceRegion"`

	// 合规类型。取值：
	// COMPLIANT：合规、
	// NON_COMPLIANT：不合规
	ComplianceType *string `json:"ComplianceType,omitnil,omitempty" name:"ComplianceType"`

	// 不合规资源的补充信息。
	Annotation *Annotation `json:"Annotation,omitnil,omitempty" name:"Annotation"`
}

type InputParameter struct {
	// 参数名
	ParameterKey *string `json:"ParameterKey,omitnil,omitempty" name:"ParameterKey"`

	// 参数类型。必填类型：Require，可选类型：Optional。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 参数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type InputParameterForManage struct {
	// 值类型。数值：Integer， 字符串：String
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// 参数Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParameterKey *string `json:"ParameterKey,omitnil,omitempty" name:"ParameterKey"`

	// 参数类型。必填类型：Require，可选类型：Optional。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 默认值
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

// Predefined struct for user
type ListAggregateConfigRulesRequestParams struct {
	// 每页限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 排序类型, 倒序：desc，顺序：asc
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 风险等级
	// 1：高风险。
	// 2：中风险。
	// 3：低风险。
	RiskLevel []*uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 规则状态
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 评估结果
	ComplianceResult []*string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// 规则名
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则所属账号ID
	RuleOwnerId *uint64 `json:"RuleOwnerId,omitnil,omitempty" name:"RuleOwnerId"`
}

type ListAggregateConfigRulesRequest struct {
	*tchttp.BaseRequest
	
	// 每页限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 排序类型, 倒序：desc，顺序：asc
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 风险等级
	// 1：高风险。
	// 2：中风险。
	// 3：低风险。
	RiskLevel []*uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 规则状态
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 评估结果
	ComplianceResult []*string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// 规则名
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则所属账号ID
	RuleOwnerId *uint64 `json:"RuleOwnerId,omitnil,omitempty" name:"RuleOwnerId"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 详情
	Items []*ConfigRule `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型, 倒序：desc，顺序：asc
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 风险等级
	// 1：高风险。
	// 2：中风险。
	// 3：低风险。
	RiskLevel []*uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 规则状态
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 评估结果
	ComplianceResult []*string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// 规则名
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

type ListConfigRulesRequest struct {
	*tchttp.BaseRequest
	
	// 每页限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型, 倒序：desc，顺序：asc
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 风险等级
	// 1：高风险。
	// 2：中风险。
	// 3：低风险。
	RiskLevel []*uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 规则状态
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 评估结果
	ComplianceResult []*string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// 规则名
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 详情
	Items []*ConfigRule `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type PutEvaluationsRequestParams struct {
	// 回调令牌。从自定义规则所选的scf云函数Context中取参数ResultToken值
	ResultToken *string `json:"ResultToken,omitnil,omitempty" name:"ResultToken"`

	// 自定义规则评估结果信息。
	Evaluations []*Evaluation `json:"Evaluations,omitnil,omitempty" name:"Evaluations"`
}

type PutEvaluationsRequest struct {
	*tchttp.BaseRequest
	
	// 回调令牌。从自定义规则所选的scf云函数Context中取参数ResultToken值
	ResultToken *string `json:"ResultToken,omitnil,omitempty" name:"ResultToken"`

	// 自定义规则评估结果信息。
	Evaluations []*Evaluation `json:"Evaluations,omitnil,omitempty" name:"Evaluations"`
}

func (r *PutEvaluationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutEvaluationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResultToken")
	delete(f, "Evaluations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PutEvaluationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutEvaluationsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PutEvaluationsResponse struct {
	*tchttp.BaseResponse
	Response *PutEvaluationsResponseParams `json:"Response"`
}

func (r *PutEvaluationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PutEvaluationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SourceConditionForManage struct {
	// 条件为空，合规：COMPLIANT，不合规：NON_COMPLIANT，无法应用：NOT_APPLICABLE
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmptyAs *string `json:"EmptyAs,omitnil,omitempty" name:"EmptyAs"`

	// 配置路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelectPath *string `json:"SelectPath,omitnil,omitempty" name:"SelectPath"`

	// 操作运算符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 是否必须
	// 注意：此字段可能返回 null，表示取不到有效值。
	Required *bool `json:"Required,omitnil,omitempty" name:"Required"`

	// 期望值
	// 注意：此字段可能返回 null，表示取不到有效值。
	DesiredValue *string `json:"DesiredValue,omitnil,omitempty" name:"DesiredValue"`
}

type Tag struct {
	// 标签key
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签value
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TriggerType struct {
	// 触发类型
	MessageType *string `json:"MessageType,omitnil,omitempty" name:"MessageType"`

	// 触发时间周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitnil,omitempty" name:"MaximumExecutionFrequency"`
}