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

package v20220802

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddAggregateCompliancePackRequestParams struct {
	// <p>合规包规则</p>
	ConfigRules []*CompliancePackRule `json:"ConfigRules,omitnil,omitempty" name:"ConfigRules"`

	// <p>风险等级<br>1：高风险。<br>2：中风险。<br>3：低风险。</p>
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// <p>合规包名称</p>
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// <p>账号组ID</p>
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type AddAggregateCompliancePackRequest struct {
	*tchttp.BaseRequest
	
	// <p>合规包规则</p>
	ConfigRules []*CompliancePackRule `json:"ConfigRules,omitnil,omitempty" name:"ConfigRules"`

	// <p>风险等级<br>1：高风险。<br>2：中风险。<br>3：低风险。</p>
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// <p>合规包名称</p>
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// <p>账号组ID</p>
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *AddAggregateCompliancePackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAggregateCompliancePackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigRules")
	delete(f, "RiskLevel")
	delete(f, "CompliancePackName")
	delete(f, "AccountGroupId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddAggregateCompliancePackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAggregateCompliancePackResponseParams struct {
	// <p>合规包Id</p>
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddAggregateCompliancePackResponse struct {
	*tchttp.BaseResponse
	Response *AddAggregateCompliancePackResponseParams `json:"Response"`
}

func (r *AddAggregateCompliancePackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAggregateCompliancePackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAggregateConfigRuleRequestParams struct {
	// 规则模板标识，预设规则模板为Identifier, 自定义规则为云函数arn（region:functionName）
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// 规则模板类型，SYSTEM, CUSTOMIZE
	IdentifierType *string `json:"IdentifierType,omitnil,omitempty" name:"IdentifierType"`

	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则支持的资源
	ResourceType []*string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 触发类型，最多支持两种
	TriggerType []*TriggerType `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 风险等级
	// 1：高风险。
	// 2：中风险。
	// 3：低风险。
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 账号组Id
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 入参
	InputParameter []*InputParameter `json:"InputParameter,omitnil,omitempty" name:"InputParameter"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 关联地域
	RegionScope []*string `json:"RegionScope,omitnil,omitempty" name:"RegionScope"`

	// 关联标签
	TagsScope []*Tag `json:"TagsScope,omitnil,omitempty" name:"TagsScope"`

	// 排除的资源ID
	ExcludeResourceIdsScope []*string `json:"ExcludeResourceIdsScope,omitnil,omitempty" name:"ExcludeResourceIdsScope"`
}

type AddAggregateConfigRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则模板标识，预设规则模板为Identifier, 自定义规则为云函数arn（region:functionName）
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// 规则模板类型，SYSTEM, CUSTOMIZE
	IdentifierType *string `json:"IdentifierType,omitnil,omitempty" name:"IdentifierType"`

	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则支持的资源
	ResourceType []*string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 触发类型，最多支持两种
	TriggerType []*TriggerType `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 风险等级
	// 1：高风险。
	// 2：中风险。
	// 3：低风险。
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 账号组Id
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 入参
	InputParameter []*InputParameter `json:"InputParameter,omitnil,omitempty" name:"InputParameter"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 关联地域
	RegionScope []*string `json:"RegionScope,omitnil,omitempty" name:"RegionScope"`

	// 关联标签
	TagsScope []*Tag `json:"TagsScope,omitnil,omitempty" name:"TagsScope"`

	// 排除的资源ID
	ExcludeResourceIdsScope []*string `json:"ExcludeResourceIdsScope,omitnil,omitempty" name:"ExcludeResourceIdsScope"`
}

func (r *AddAggregateConfigRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAggregateConfigRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Identifier")
	delete(f, "IdentifierType")
	delete(f, "RuleName")
	delete(f, "ResourceType")
	delete(f, "TriggerType")
	delete(f, "RiskLevel")
	delete(f, "AccountGroupId")
	delete(f, "InputParameter")
	delete(f, "Description")
	delete(f, "RegionScope")
	delete(f, "TagsScope")
	delete(f, "ExcludeResourceIdsScope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddAggregateConfigRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAggregateConfigRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddAggregateConfigRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddAggregateConfigRuleResponseParams `json:"Response"`
}

func (r *AddAggregateConfigRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAggregateConfigRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAlarmPolicyRequestParams struct {
	// 告警策略名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 事件范围  1：当前账号  2：多账号
	EventScope []*int64 `json:"EventScope,omitnil,omitempty" name:"EventScope"`

	// 风险等级 1：高风险  2：中风险 3：低风险
	RiskLevel []*int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 通知时间段
	NoticeTime *string `json:"NoticeTime,omitnil,omitempty" name:"NoticeTime"`

	// 通知机制
	NotificationMechanism *string `json:"NotificationMechanism,omitnil,omitempty" name:"NotificationMechanism"`

	// 状态 1：启用 2：停用
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 通知周期
	NoticePeriod []*int64 `json:"NoticePeriod,omitnil,omitempty" name:"NoticePeriod"`

	// 策略描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type AddAlarmPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 告警策略名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 事件范围  1：当前账号  2：多账号
	EventScope []*int64 `json:"EventScope,omitnil,omitempty" name:"EventScope"`

	// 风险等级 1：高风险  2：中风险 3：低风险
	RiskLevel []*int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 通知时间段
	NoticeTime *string `json:"NoticeTime,omitnil,omitempty" name:"NoticeTime"`

	// 通知机制
	NotificationMechanism *string `json:"NotificationMechanism,omitnil,omitempty" name:"NotificationMechanism"`

	// 状态 1：启用 2：停用
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 通知周期
	NoticePeriod []*int64 `json:"NoticePeriod,omitnil,omitempty" name:"NoticePeriod"`

	// 策略描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *AddAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAlarmPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "EventScope")
	delete(f, "RiskLevel")
	delete(f, "NoticeTime")
	delete(f, "NotificationMechanism")
	delete(f, "Status")
	delete(f, "NoticePeriod")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddAlarmPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAlarmPolicyResponseParams struct {
	// 告警策略唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmPolicyId *uint64 `json:"AlarmPolicyId,omitnil,omitempty" name:"AlarmPolicyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *AddAlarmPolicyResponseParams `json:"Response"`
}

func (r *AddAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCompliancePackRequestParams struct {
	// <p>合规包规则</p>
	ConfigRules []*CompliancePackRule `json:"ConfigRules,omitnil,omitempty" name:"ConfigRules"`

	// <p>风险等级<br>1：高风险。<br>2：中风险。<br>3：低风险。</p>
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// <p>合规包名称</p>
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type AddCompliancePackRequest struct {
	*tchttp.BaseRequest
	
	// <p>合规包规则</p>
	ConfigRules []*CompliancePackRule `json:"ConfigRules,omitnil,omitempty" name:"ConfigRules"`

	// <p>风险等级<br>1：高风险。<br>2：中风险。<br>3：低风险。</p>
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// <p>合规包名称</p>
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *AddCompliancePackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCompliancePackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigRules")
	delete(f, "RiskLevel")
	delete(f, "CompliancePackName")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCompliancePackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCompliancePackResponseParams struct {
	// <p>合规包Id</p>
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddCompliancePackResponse struct {
	*tchttp.BaseResponse
	Response *AddCompliancePackResponseParams `json:"Response"`
}

func (r *AddCompliancePackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCompliancePackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddConfigRuleRequestParams struct {
	// <p>规则模板标识，预设规则模板为Identifier, 自定义规则为云函数arn（region:functionName）</p>
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// <p>规则模板类型，SYSTEM, CUSTOMIZE</p>
	IdentifierType *string `json:"IdentifierType,omitnil,omitempty" name:"IdentifierType"`

	// <p>规则名称</p>
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// <p>规则支持的资源</p>
	ResourceType []*string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>触发类型，最多支持两种</p>
	TriggerType []*TriggerType `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>风险等级<br>1：高风险。<br>2：中风险。<br>3：低风险。</p>
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// <p>入参</p>
	InputParameter []*InputParameter `json:"InputParameter,omitnil,omitempty" name:"InputParameter"`

	// <p>规则描述。长度范围0~1024字符</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>规则评估地域范围，规则仅对指定地域中的资源生效。<br>支持的地域范围config:ListResourceRegions返回的地域</p>
	RegionsScope []*string `json:"RegionsScope,omitnil,omitempty" name:"RegionsScope"`

	// <p>规则评估标签范围，规则仅对绑定指定标签的资源生效。</p>
	TagsScope []*Tag `json:"TagsScope,omitnil,omitempty" name:"TagsScope"`

	// <p>规则对指定资源ID无效，即不对该资源执行评估。</p>
	ExcludeResourceIdsScope []*string `json:"ExcludeResourceIdsScope,omitnil,omitempty" name:"ExcludeResourceIdsScope"`
}

type AddConfigRuleRequest struct {
	*tchttp.BaseRequest
	
	// <p>规则模板标识，预设规则模板为Identifier, 自定义规则为云函数arn（region:functionName）</p>
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// <p>规则模板类型，SYSTEM, CUSTOMIZE</p>
	IdentifierType *string `json:"IdentifierType,omitnil,omitempty" name:"IdentifierType"`

	// <p>规则名称</p>
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// <p>规则支持的资源</p>
	ResourceType []*string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>触发类型，最多支持两种</p>
	TriggerType []*TriggerType `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>风险等级<br>1：高风险。<br>2：中风险。<br>3：低风险。</p>
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// <p>入参</p>
	InputParameter []*InputParameter `json:"InputParameter,omitnil,omitempty" name:"InputParameter"`

	// <p>规则描述。长度范围0~1024字符</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>规则评估地域范围，规则仅对指定地域中的资源生效。<br>支持的地域范围config:ListResourceRegions返回的地域</p>
	RegionsScope []*string `json:"RegionsScope,omitnil,omitempty" name:"RegionsScope"`

	// <p>规则评估标签范围，规则仅对绑定指定标签的资源生效。</p>
	TagsScope []*Tag `json:"TagsScope,omitnil,omitempty" name:"TagsScope"`

	// <p>规则对指定资源ID无效，即不对该资源执行评估。</p>
	ExcludeResourceIdsScope []*string `json:"ExcludeResourceIdsScope,omitnil,omitempty" name:"ExcludeResourceIdsScope"`
}

func (r *AddConfigRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddConfigRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Identifier")
	delete(f, "IdentifierType")
	delete(f, "RuleName")
	delete(f, "ResourceType")
	delete(f, "TriggerType")
	delete(f, "RiskLevel")
	delete(f, "InputParameter")
	delete(f, "Description")
	delete(f, "RegionsScope")
	delete(f, "TagsScope")
	delete(f, "ExcludeResourceIdsScope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddConfigRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddConfigRuleResponseParams struct {
	// <p>规则ID</p>
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddConfigRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddConfigRuleResponseParams `json:"Response"`
}

func (r *AddConfigRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddConfigRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AggregateEvaluationResult struct {
	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源地域
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`

	// 资源名称
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 规则id
	ConfigRuleId *string `json:"ConfigRuleId,omitnil,omitempty" name:"ConfigRuleId"`

	// 规则名称
	ConfigRuleName *string `json:"ConfigRuleName,omitnil,omitempty" name:"ConfigRuleName"`

	// 合规包id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 风险等级
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 评估结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Annotation *Annotation `json:"Annotation,omitnil,omitempty" name:"Annotation"`

	// 合规类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplianceType *string `json:"ComplianceType,omitnil,omitempty" name:"ComplianceType"`

	// 规则发起类型
	InvokingEventMessageType *string `json:"InvokingEventMessageType,omitnil,omitempty" name:"InvokingEventMessageType"`

	// 评估发起时间
	ConfigRuleInvokedTime *string `json:"ConfigRuleInvokedTime,omitnil,omitempty" name:"ConfigRuleInvokedTime"`

	// 评估实际时间
	ResultRecordedTime *string `json:"ResultRecordedTime,omitnil,omitempty" name:"ResultRecordedTime"`

	// 资源所属用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceOwnerId *uint64 `json:"ResourceOwnerId,omitnil,omitempty" name:"ResourceOwnerId"`

	// 资源所属用户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceOwnerName *string `json:"ResourceOwnerName,omitnil,omitempty" name:"ResourceOwnerName"`
}

type AggregateResourceInfo struct {
	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源名称
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 地域
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`

	// 资源状态
	ResourceStatus *string `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`

	// 是否删除 1:已删除 0:未删除
	ResourceDelete *uint64 `json:"ResourceDelete,omitnil,omitempty" name:"ResourceDelete"`

	// 资源创建时间
	ResourceCreateTime *string `json:"ResourceCreateTime,omitnil,omitempty" name:"ResourceCreateTime"`

	// 标签信息
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 可用区
	ResourceZone *string `json:"ResourceZone,omitnil,omitempty" name:"ResourceZone"`

	// 合规状态
	ComplianceResult *string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// 资源所属用户ID
	ResourceOwnerId *uint64 `json:"ResourceOwnerId,omitnil,omitempty" name:"ResourceOwnerId"`

	// 用户昵称
	ResourceOwnerName *string `json:"ResourceOwnerName,omitnil,omitempty" name:"ResourceOwnerName"`
}

type Aggregator struct {
	// 账号组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 账号组描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建者用户ID
	OwnerUin *uint64 `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 账号组成员数量
	AccountCount *uint64 `json:"AccountCount,omitnil,omitempty" name:"AccountCount"`

	// RD:全局账号组
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 账号组状态
	AggregatorStatus *uint64 `json:"AggregatorStatus,omitnil,omitempty" name:"AggregatorStatus"`

	// 账号组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberName *string `json:"MemberName,omitnil,omitempty" name:"MemberName"`
}

type AggregatorAccount struct {
	// 成员ID
	MemberUin *uint64 `json:"MemberUin,omitnil,omitempty" name:"MemberUin"`

	// 成员名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberName *string `json:"MemberName,omitnil,omitempty" name:"MemberName"`
}

type AlarmPolicyRsp struct {
	// 告警策略唯一标识id
	AlarmPolicyId *uint64 `json:"AlarmPolicyId,omitnil,omitempty" name:"AlarmPolicyId"`

	// 策略名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 事件类型
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 1：当前账号  2：多账号
	EventScope []*int64 `json:"EventScope,omitnil,omitempty" name:"EventScope"`

	// 1：高风险  2：中风险 3：低风险
	RiskLevel []*int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 通知周期，1-7数字分别代表周一至周日
	NoticePeriod []*int64 `json:"NoticePeriod,omitnil,omitempty" name:"NoticePeriod"`

	// 通知时间段
	NoticeTime *string `json:"NoticeTime,omitnil,omitempty" name:"NoticeTime"`

	// 通知机制
	NotificationMechanism *string `json:"NotificationMechanism,omitnil,omitempty" name:"NotificationMechanism"`

	// 策略状态 1：启动  2：停止
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 策略描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type Annotation struct {
	// 资源当前实际配置。长度为0~256位字符，即资源不合规配置
	Configuration *string `json:"Configuration,omitnil,omitempty" name:"Configuration"`

	// 资源期望配置。长度为0~256位字符，即资源合规配置
	DesiredValue *string `json:"DesiredValue,omitnil,omitempty" name:"DesiredValue"`

	// 资源当前配置和期望配置之间的比较运算符。长度为0~16位字符，自定义规则上报评估结果此字段可能为空
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 当前配置在资源属性结构体中的JSON路径。长度为0~256位字符，自定义规则上报评估结果此字段可能为空
	Property *string `json:"Property,omitnil,omitempty" name:"Property"`
}

// Predefined struct for user
type CloseAggregateConfigRuleRequestParams struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 账号组Id
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

type CloseAggregateConfigRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 账号组Id
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

func (r *CloseAggregateConfigRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseAggregateConfigRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "AccountGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseAggregateConfigRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseAggregateConfigRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseAggregateConfigRuleResponse struct {
	*tchttp.BaseResponse
	Response *CloseAggregateConfigRuleResponseParams `json:"Response"`
}

func (r *CloseAggregateConfigRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseAggregateConfigRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseConfigRecorderRequestParams struct {

}

type CloseConfigRecorderRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CloseConfigRecorderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseConfigRecorderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseConfigRecorderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseConfigRecorderResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseConfigRecorderResponse struct {
	*tchttp.BaseResponse
	Response *CloseConfigRecorderResponseParams `json:"Response"`
}

func (r *CloseConfigRecorderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseConfigRecorderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseConfigRuleRequestParams struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type CloseConfigRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

func (r *CloseConfigRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseConfigRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseConfigRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseConfigRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseConfigRuleResponse struct {
	*tchttp.BaseResponse
	Response *CloseConfigRuleResponseParams `json:"Response"`
}

func (r *CloseConfigRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseConfigRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComplianceConfigRule struct {
	// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 风险等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 规则状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigRuleId *string `json:"ConfigRuleId,omitnil,omitempty" name:"ConfigRuleId"`

	// 评估结果
	// 合规： 'COMPLIANT'
	// 不合规： 'NON_COMPLIANT'
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplianceResult *string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// 关键字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*string `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 入参
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputParameter []*InputParameter `json:"InputParameter,omitnil,omitempty" name:"InputParameter"`

	// 参数格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceCondition []*SourceConditionForManage `json:"SourceCondition,omitnil,omitempty" name:"SourceCondition"`

	// 规则标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// 创建方式 user、compliancePack、system
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateByType *string `json:"CreateByType,omitnil,omitempty" name:"CreateByType"`

	// 规则描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 参数描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManageInputParameter []*InputParameterForManage `json:"ManageInputParameter,omitnil,omitempty" name:"ManageInputParameter"`
}

type CompliancePackRule struct {
	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 风险等级
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 入参
	InputParameter []*InputParameter `json:"InputParameter,omitnil,omitempty" name:"InputParameter"`

	// 规则身份标识
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 预设规则身份标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManagedRuleIdentifier *string `json:"ManagedRuleIdentifier,omitnil,omitempty" name:"ManagedRuleIdentifier"`

	// 规则ID
	ConfigRuleId *string `json:"ConfigRuleId,omitnil,omitempty" name:"ConfigRuleId"`

	// 合规包ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`
}

type CompliancePackRuleForManage struct {
	// <p>规则唯一身份标识</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// <p>规则名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// <p>描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>风险等级</p><p>1：高风险 2：中风险  3：低风险</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// <p>创建时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>合规包规则</p>
	CompliancePackRules []*CompliancePackRules `json:"CompliancePackRules,omitnil,omitempty" name:"CompliancePackRules"`

	// <p>规则编号信息</p>
	Control []*Control `json:"Control,omitnil,omitempty" name:"Control"`

	// <p>资源类型</p>
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`
}

type CompliancePackRules struct {
	// <p>规则标识</p>
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// <p>规则编号信息</p>
	Control []*Control `json:"Control,omitnil,omitempty" name:"Control"`

	// <p>资源类型</p>
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`
}

type ConfigCompliancePack struct {
	// 合规包状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 合规包名称
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// 风险等级
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 评估结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplianceResult *string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 不合规规则名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoCompliantNames []*string `json:"NoCompliantNames,omitnil,omitempty" name:"NoCompliantNames"`

	// 合规包规则数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleCount *uint64 `json:"RuleCount,omitnil,omitempty" name:"RuleCount"`
}

type ConfigResource struct {
	// 产品
	// 注意：此字段可能返回 null，表示取不到有效值。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 产品名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源类型名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceTypeName *string `json:"ResourceTypeName,omitnil,omitempty" name:"ResourceTypeName"`
}

type ConfigRule struct {
	// 规则标识
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// 规则名
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则参数
	InputParameter []*InputParameter `json:"InputParameter,omitnil,omitempty" name:"InputParameter"`

	// 规则触发条件
	SourceCondition []*SourceConditionForManage `json:"SourceCondition,omitnil,omitempty" name:"SourceCondition"`

	// 规则支持的资源类型，规则仅对指定资源类型的资源生效。
	ResourceType []*string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 规则所属标签
	Labels []*string `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 规则风险等级
	// 1:低风险
	// 2:中风险
	// 3:高风险
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 规则对应的函数
	ServiceFunction *string `json:"ServiceFunction,omitnil,omitempty" name:"ServiceFunction"`

	// 创建时间
	// 格式：YYYY-MM-DD h:i:s
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// ACTIVE：启用
	// NO_ACTIVE：停止
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 合规： 'COMPLIANT'
	// 不合规： 'NON_COMPLIANT'
	// 无法应用规则： 'NOT_APPLICABLE'
	ComplianceResult *string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// ["",""]
	Annotation *Annotation `json:"Annotation,omitnil,omitempty" name:"Annotation"`

	// 规则评估时间
	// 格式：YYYY-MM-DD h:i:s
	ConfigRuleInvokedTime *string `json:"ConfigRuleInvokedTime,omitnil,omitempty" name:"ConfigRuleInvokedTime"`

	// 规则ID
	ConfigRuleId *string `json:"ConfigRuleId,omitnil,omitempty" name:"ConfigRuleId"`

	// CUSTOMIZE：自定义规则、
	// SYSTEM：托管规则
	IdentifierType *string `json:"IdentifierType,omitnil,omitempty" name:"IdentifierType"`

	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 触发类型
	// ScheduledNotification：周期触发、
	// ConfigurationItemChangeNotification：变更触发
	TriggerType []*TriggerType `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 参数详情
	ManageInputParameter []*InputParameterForManage `json:"ManageInputParameter,omitnil,omitempty" name:"ManageInputParameter"`

	// 合规包名称
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// 关联地域
	RegionsScope []*string `json:"RegionsScope,omitnil,omitempty" name:"RegionsScope"`

	// 关联标签
	TagsScope []*Tag `json:"TagsScope,omitnil,omitempty" name:"TagsScope"`

	//  规则对指定资源ID无效，即不对该资源执行评估。
	ExcludeResourceIdsScope []*string `json:"ExcludeResourceIdsScope,omitnil,omitempty" name:"ExcludeResourceIdsScope"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 账号组名称
	AccountGroupName *string `json:"AccountGroupName,omitnil,omitempty" name:"AccountGroupName"`

	// 规则所属用户ID
	RuleOwnerId *uint64 `json:"RuleOwnerId,omitnil,omitempty" name:"RuleOwnerId"`

	// 预设规则支持的触发方式
	// ScheduledNotification：周期触发
	// ConfigurationItemChangeNotification：变更触发
	ManageTriggerType []*string `json:"ManageTriggerType,omitnil,omitempty" name:"ManageTriggerType"`
}

type Control struct {
	// <p>规则编号</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>编号描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

// Predefined struct for user
type CreateAggregatorRequestParams struct {
	// <p>账号组名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>账号组描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>账号组类型 </p><p>枚举值：</p><ul><li>RD： 全局账号组</li><li>CUSTOM： 自定义账号组</li></ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>账号组成员信息列表，最多100个</p>
	AggregatorAccounts []*AggregatorAccount `json:"AggregatorAccounts,omitnil,omitempty" name:"AggregatorAccounts"`
}

type CreateAggregatorRequest struct {
	*tchttp.BaseRequest
	
	// <p>账号组名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>账号组描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>账号组类型 </p><p>枚举值：</p><ul><li>RD： 全局账号组</li><li>CUSTOM： 自定义账号组</li></ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>账号组成员信息列表，最多100个</p>
	AggregatorAccounts []*AggregatorAccount `json:"AggregatorAccounts,omitnil,omitempty" name:"AggregatorAccounts"`
}

func (r *CreateAggregatorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAggregatorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Type")
	delete(f, "AggregatorAccounts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAggregatorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAggregatorResponseParams struct {
	// <p>账号组Id</p>
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAggregatorResponse struct {
	*tchttp.BaseResponse
	Response *CreateAggregatorResponseParams `json:"Response"`
}

func (r *CreateAggregatorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAggregatorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRemediationRequestParams struct {
	// <p>规则 ID</p>
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// <p>修正类型取值：   SCF：云函数（自定义修正）。</p>
	RemediationType *string `json:"RemediationType,omitnil,omitempty" name:"RemediationType"`

	// <p>修正模板 ID</p>
	RemediationTemplateId *string `json:"RemediationTemplateId,omitnil,omitempty" name:"RemediationTemplateId"`

	// <p>修正执行方式。<br>取值：   MANUAL_EXECUTION：手动执行</p>
	InvokeType *string `json:"InvokeType,omitnil,omitempty" name:"InvokeType"`

	// <p>执行修正的模板来源<br>CUSTOM：自定义模板。</p>
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`
}

type CreateRemediationRequest struct {
	*tchttp.BaseRequest
	
	// <p>规则 ID</p>
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// <p>修正类型取值：   SCF：云函数（自定义修正）。</p>
	RemediationType *string `json:"RemediationType,omitnil,omitempty" name:"RemediationType"`

	// <p>修正模板 ID</p>
	RemediationTemplateId *string `json:"RemediationTemplateId,omitnil,omitempty" name:"RemediationTemplateId"`

	// <p>修正执行方式。<br>取值：   MANUAL_EXECUTION：手动执行</p>
	InvokeType *string `json:"InvokeType,omitnil,omitempty" name:"InvokeType"`

	// <p>执行修正的模板来源<br>CUSTOM：自定义模板。</p>
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`
}

func (r *CreateRemediationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRemediationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "RemediationType")
	delete(f, "RemediationTemplateId")
	delete(f, "InvokeType")
	delete(f, "SourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRemediationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRemediationResponseParams struct {
	// <p>修正设置Id</p>
	RemediationId *string `json:"RemediationId,omitnil,omitempty" name:"RemediationId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRemediationResponse struct {
	*tchttp.BaseResponse
	Response *CreateRemediationResponseParams `json:"Response"`
}

func (r *CreateRemediationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRemediationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAggregateCompliancePackRequestParams struct {
	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

type DeleteAggregateCompliancePackRequest struct {
	*tchttp.BaseRequest
	
	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

func (r *DeleteAggregateCompliancePackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAggregateCompliancePackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompliancePackId")
	delete(f, "AccountGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAggregateCompliancePackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAggregateCompliancePackResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAggregateCompliancePackResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAggregateCompliancePackResponseParams `json:"Response"`
}

func (r *DeleteAggregateCompliancePackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAggregateCompliancePackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAggregateConfigRuleRequestParams struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

type DeleteAggregateConfigRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

func (r *DeleteAggregateConfigRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAggregateConfigRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "AccountGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAggregateConfigRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAggregateConfigRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAggregateConfigRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAggregateConfigRuleResponseParams `json:"Response"`
}

func (r *DeleteAggregateConfigRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAggregateConfigRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAlarmPolicyRequestParams struct {
	// 告警策略id
	AlarmPolicyId *uint64 `json:"AlarmPolicyId,omitnil,omitempty" name:"AlarmPolicyId"`
}

type DeleteAlarmPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 告警策略id
	AlarmPolicyId *uint64 `json:"AlarmPolicyId,omitnil,omitempty" name:"AlarmPolicyId"`
}

func (r *DeleteAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmPolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAlarmPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAlarmPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAlarmPolicyResponseParams `json:"Response"`
}

func (r *DeleteAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCompliancePackRequestParams struct {
	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`
}

type DeleteCompliancePackRequest struct {
	*tchttp.BaseRequest
	
	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`
}

func (r *DeleteCompliancePackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCompliancePackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompliancePackId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCompliancePackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCompliancePackResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCompliancePackResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCompliancePackResponseParams `json:"Response"`
}

func (r *DeleteCompliancePackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCompliancePackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigRuleRequestParams struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type DeleteConfigRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

func (r *DeleteConfigRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConfigRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteConfigRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConfigRuleResponseParams `json:"Response"`
}

func (r *DeleteConfigRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRemediationsRequestParams struct {
	// 修正设置ID
	RemediationIds []*string `json:"RemediationIds,omitnil,omitempty" name:"RemediationIds"`
}

type DeleteRemediationsRequest struct {
	*tchttp.BaseRequest
	
	// 修正设置ID
	RemediationIds []*string `json:"RemediationIds,omitnil,omitempty" name:"RemediationIds"`
}

func (r *DeleteRemediationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRemediationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RemediationIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRemediationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRemediationsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRemediationsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRemediationsResponseParams `json:"Response"`
}

func (r *DeleteRemediationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRemediationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAggregateCompliancePackRequestParams struct {
	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`
}

type DescribeAggregateCompliancePackRequest struct {
	*tchttp.BaseRequest
	
	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`
}

func (r *DescribeAggregateCompliancePackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAggregateCompliancePackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountGroupId")
	delete(f, "CompliancePackId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAggregateCompliancePackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAggregateCompliancePackResponseParams struct {
	// 合规包名称
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 风险等级
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 合规包规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigRules []*ComplianceConfigRule `json:"ConfigRules,omitnil,omitempty" name:"ConfigRules"`

	// 合规包id
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 合规包状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAggregateCompliancePackResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAggregateCompliancePackResponseParams `json:"Response"`
}

func (r *DescribeAggregateCompliancePackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAggregateCompliancePackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAggregateConfigDeliverRequestParams struct {
	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

type DescribeAggregateConfigDeliverRequest struct {
	*tchttp.BaseRequest
	
	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

func (r *DescribeAggregateConfigDeliverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAggregateConfigDeliverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAggregateConfigDeliverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAggregateConfigDeliverResponseParams struct {
	// 投递名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeliverName *string `json:"DeliverName,omitnil,omitempty" name:"DeliverName"`

	// 资源六段式
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetArn *string `json:"TargetArn,omitnil,omitempty" name:"TargetArn"`

	// 投递状态 DeliverStatus：0 关闭  1 开启
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 日志前缀
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeliverPrefix *string `json:"DeliverPrefix,omitnil,omitempty" name:"DeliverPrefix"`

	// 投递类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeliverType *string `json:"DeliverType,omitnil,omitempty" name:"DeliverType"`

	// 支持跨账号投递的成员账号uin，只能是委派管理员。默认为0，即投递到管理员账号下
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeliverUin *int64 `json:"DeliverUin,omitnil,omitempty" name:"DeliverUin"`

	// 1：配置变更 2： 资源列表 3：全部
	DeliverContentType *uint64 `json:"DeliverContentType,omitnil,omitempty" name:"DeliverContentType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAggregateConfigDeliverResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAggregateConfigDeliverResponseParams `json:"Response"`
}

func (r *DescribeAggregateConfigDeliverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAggregateConfigDeliverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAggregateConfigRuleRequestParams struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 账号组
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

type DescribeAggregateConfigRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 账号组
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

func (r *DescribeAggregateConfigRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAggregateConfigRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "AccountGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAggregateConfigRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAggregateConfigRuleResponseParams struct {
	// 规则详情
	ConfigRule *ConfigRule `json:"ConfigRule,omitnil,omitempty" name:"ConfigRule"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAggregateConfigRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAggregateConfigRuleResponseParams `json:"Response"`
}

func (r *DescribeAggregateConfigRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAggregateConfigRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAggregateDiscoveredResourceRequestParams struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源地域
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 资源所属用户ID
	ResourceOwnerId *uint64 `json:"ResourceOwnerId,omitnil,omitempty" name:"ResourceOwnerId"`
}

type DescribeAggregateDiscoveredResourceRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源地域
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 资源所属用户ID
	ResourceOwnerId *uint64 `json:"ResourceOwnerId,omitnil,omitempty" name:"ResourceOwnerId"`
}

func (r *DescribeAggregateDiscoveredResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAggregateDiscoveredResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "ResourceType")
	delete(f, "ResourceRegion")
	delete(f, "AccountGroupId")
	delete(f, "ResourceOwnerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAggregateDiscoveredResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAggregateDiscoveredResourceResponseParams struct {
	// 资源Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源名
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 资源地域
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`

	// 资源可用区
	ResourceZone *string `json:"ResourceZone,omitnil,omitempty" name:"ResourceZone"`

	// 资源配置
	Configuration *string `json:"Configuration,omitnil,omitempty" name:"Configuration"`

	// 资源创建时间
	ResourceCreateTime *string `json:"ResourceCreateTime,omitnil,omitempty" name:"ResourceCreateTime"`

	// 资源标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 资源更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAggregateDiscoveredResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAggregateDiscoveredResourceResponseParams `json:"Response"`
}

func (r *DescribeAggregateDiscoveredResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAggregateDiscoveredResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAggregatorRequestParams struct {
	// <p>账号组ID</p>
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// <p>账号组创建者ID</p>
	OwnerUin *uint64 `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

type DescribeAggregatorRequest struct {
	*tchttp.BaseRequest
	
	// <p>账号组ID</p>
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// <p>账号组创建者ID</p>
	OwnerUin *uint64 `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

func (r *DescribeAggregatorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAggregatorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccountGroupId")
	delete(f, "OwnerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAggregatorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAggregatorResponseParams struct {
	// <p>账号组名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>账号组描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>账号组类型</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>成员信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AggregatorAccounts []*AggregatorAccount `json:"AggregatorAccounts,omitnil,omitempty" name:"AggregatorAccounts"`

	// <p>创建状态</p>
	AggregatorStatus *uint64 `json:"AggregatorStatus,omitnil,omitempty" name:"AggregatorStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAggregatorResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAggregatorResponseParams `json:"Response"`
}

func (r *DescribeAggregatorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAggregatorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompliancePackRequestParams struct {
	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`
}

type DescribeCompliancePackRequest struct {
	*tchttp.BaseRequest
	
	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`
}

func (r *DescribeCompliancePackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompliancePackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompliancePackId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCompliancePackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompliancePackResponseParams struct {
	// 合规包名称
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 风险等级
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 合规包规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigRules []*ComplianceConfigRule `json:"ConfigRules,omitnil,omitempty" name:"ConfigRules"`

	// 合规包id
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 合规包状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCompliancePackResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCompliancePackResponseParams `json:"Response"`
}

func (r *DescribeCompliancePackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompliancePackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigDeliverRequestParams struct {

}

type DescribeConfigDeliverRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeConfigDeliverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigDeliverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigDeliverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigDeliverResponseParams struct {
	// 投递名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeliverName *string `json:"DeliverName,omitnil,omitempty" name:"DeliverName"`

	// 资源六段式
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetArn *string `json:"TargetArn,omitnil,omitempty" name:"TargetArn"`

	// 投递状态 DeliverStatus：0 关闭  1 开启
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 日志前缀
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeliverPrefix *string `json:"DeliverPrefix,omitnil,omitempty" name:"DeliverPrefix"`

	// 投递类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeliverType *string `json:"DeliverType,omitnil,omitempty" name:"DeliverType"`

	// 1：配置变更   2： 资源列表 3：全部
	DeliverContentType *uint64 `json:"DeliverContentType,omitnil,omitempty" name:"DeliverContentType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigDeliverResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigDeliverResponseParams `json:"Response"`
}

func (r *DescribeConfigDeliverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigDeliverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigRecorderRequestParams struct {

}

type DescribeConfigRecorderRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeConfigRecorderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigRecorderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigRecorderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigRecorderResponseParams struct {
	// 用户监控资源列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*UserConfigResource `json:"Items,omitnil,omitempty" name:"Items"`

	// 0 关闭1 打开
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 当日快照次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerCount *uint64 `json:"TriggerCount,omitnil,omitempty" name:"TriggerCount"`

	// 当日打开监控次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenCount *uint64 `json:"OpenCount,omitnil,omitempty" name:"OpenCount"`

	// 当日修改监控范围次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateCount *uint64 `json:"UpdateCount,omitnil,omitempty" name:"UpdateCount"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigRecorderResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigRecorderResponseParams `json:"Response"`
}

func (r *DescribeConfigRecorderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigRecorderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigRuleRequestParams struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type DescribeConfigRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

func (r *DescribeConfigRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigRuleResponseParams struct {
	// 规则详情
	ConfigRule *ConfigRule `json:"ConfigRule,omitnil,omitempty" name:"ConfigRule"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigRuleResponseParams `json:"Response"`
}

func (r *DescribeConfigRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiscoveredResourceRequestParams struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源地域
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`
}

type DescribeDiscoveredResourceRequest struct {
	*tchttp.BaseRequest
	
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源地域
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`
}

func (r *DescribeDiscoveredResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiscoveredResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "ResourceType")
	delete(f, "ResourceRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiscoveredResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiscoveredResourceResponseParams struct {
	// 资源Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源名
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 资源地域
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`

	// 资源可用区
	ResourceZone *string `json:"ResourceZone,omitnil,omitempty" name:"ResourceZone"`

	// 资源配置
	Configuration *string `json:"Configuration,omitnil,omitempty" name:"Configuration"`

	// 资源创建时间
	ResourceCreateTime *string `json:"ResourceCreateTime,omitnil,omitempty" name:"ResourceCreateTime"`

	// 资源标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 资源更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDiscoveredResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiscoveredResourceResponseParams `json:"Response"`
}

func (r *DescribeDiscoveredResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiscoveredResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSystemCompliancePackRequestParams struct {
	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`
}

type DescribeSystemCompliancePackRequest struct {
	*tchttp.BaseRequest
	
	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`
}

func (r *DescribeSystemCompliancePackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSystemCompliancePackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompliancePackId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSystemCompliancePackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSystemCompliancePackResponseParams struct {
	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 合规包名称
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// 合规包描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 风险等级
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 合规包规则列表
	ConfigRules []*CompliancePackRuleForManage `json:"ConfigRules,omitnil,omitempty" name:"ConfigRules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSystemCompliancePackResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSystemCompliancePackResponseParams `json:"Response"`
}

func (r *DescribeSystemCompliancePackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSystemCompliancePackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSystemRuleRequestParams struct {
	// 规则唯一标识
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`
}

type DescribeSystemRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则唯一标识
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`
}

func (r *DescribeSystemRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSystemRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Identifier")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSystemRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSystemRuleResponseParams struct {
	// 详情
	SystemConfigRule *SystemConfigRule `json:"SystemConfigRule,omitnil,omitempty" name:"SystemConfigRule"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSystemRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSystemRuleResponseParams `json:"Response"`
}

func (r *DescribeSystemRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSystemRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachAggregateConfigRuleToCompliancePackRequestParams struct {
	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 规则ID
	ConfigRuleId *string `json:"ConfigRuleId,omitnil,omitempty" name:"ConfigRuleId"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

type DetachAggregateConfigRuleToCompliancePackRequest struct {
	*tchttp.BaseRequest
	
	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 规则ID
	ConfigRuleId *string `json:"ConfigRuleId,omitnil,omitempty" name:"ConfigRuleId"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

func (r *DetachAggregateConfigRuleToCompliancePackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachAggregateConfigRuleToCompliancePackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompliancePackId")
	delete(f, "ConfigRuleId")
	delete(f, "AccountGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachAggregateConfigRuleToCompliancePackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachAggregateConfigRuleToCompliancePackResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetachAggregateConfigRuleToCompliancePackResponse struct {
	*tchttp.BaseResponse
	Response *DetachAggregateConfigRuleToCompliancePackResponseParams `json:"Response"`
}

func (r *DetachAggregateConfigRuleToCompliancePackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachAggregateConfigRuleToCompliancePackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachConfigRuleToCompliancePackRequestParams struct {
	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 规则ID
	ConfigRuleId *string `json:"ConfigRuleId,omitnil,omitempty" name:"ConfigRuleId"`
}

type DetachConfigRuleToCompliancePackRequest struct {
	*tchttp.BaseRequest
	
	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 规则ID
	ConfigRuleId *string `json:"ConfigRuleId,omitnil,omitempty" name:"ConfigRuleId"`
}

func (r *DetachConfigRuleToCompliancePackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachConfigRuleToCompliancePackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompliancePackId")
	delete(f, "ConfigRuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachConfigRuleToCompliancePackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachConfigRuleToCompliancePackResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetachConfigRuleToCompliancePackResponse struct {
	*tchttp.BaseResponse
	Response *DetachConfigRuleToCompliancePackResponseParams `json:"Response"`
}

func (r *DetachConfigRuleToCompliancePackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachConfigRuleToCompliancePackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type EvaluationResult struct {
	// 资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`

	// 资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 规则id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigRuleId *string `json:"ConfigRuleId,omitnil,omitempty" name:"ConfigRuleId"`

	// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigRuleName *string `json:"ConfigRuleName,omitnil,omitempty" name:"ConfigRuleName"`

	// 合规包id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 风险等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 评估结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Annotation *Annotation `json:"Annotation,omitnil,omitempty" name:"Annotation"`

	// 合规类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComplianceType *string `json:"ComplianceType,omitnil,omitempty" name:"ComplianceType"`

	// 规则发起类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokingEventMessageType *string `json:"InvokingEventMessageType,omitnil,omitempty" name:"InvokingEventMessageType"`

	// 评估发起时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigRuleInvokedTime *string `json:"ConfigRuleInvokedTime,omitnil,omitempty" name:"ConfigRuleInvokedTime"`

	// 评估实际时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultRecordedTime *string `json:"ResultRecordedTime,omitnil,omitempty" name:"ResultRecordedTime"`
}

type Filter struct {
	// 查询字段名称 资源名称：resourceName  资源ID：resourceId 资源类型：resourceType 资源地域：resourceRegion    删除状态：resourceDelete 0未删除，1已删除  resourceRegionAndZone地域/可用区
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 查询字段值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type InputParameter struct {
	// 参数名
	ParameterKey *string `json:"ParameterKey,omitnil,omitempty" name:"ParameterKey"`

	// 参数类型。必填类型：Require，可选类型：Optional。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 参数值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type InputParameterForManage struct {
	// 值类型。数值：Integer， 字符串：String
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// 参数Key
	ParameterKey *string `json:"ParameterKey,omitnil,omitempty" name:"ParameterKey"`

	// 参数类型。必填类型：Require，可选类型：Optional。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 默认值
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

// Predefined struct for user
type ListAggregateCompliancePacksRequestParams struct {
	// <p>数量</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>账号组ID</p>
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// <p>合规包名称</p>
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// <p>风险等级<br>1：高风险。<br>2：中风险。<br>3：低风险。</p>
	RiskLevel []*uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// <p>合规包状态 ACTIVE、NO_ACTIVE</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>评估状态合规： &#39;COMPLIANT&#39;<br>不合规： &#39;NON_COMPLIANT&#39;</p>
	ComplianceResult []*string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// <p>排序类型, 倒序：desc，顺序：asc</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>包含合规包结果定义</p><p>枚举值：</p><ul><li>NO： 不包含</li></ul><p>默认值：空</p><p>此字段为新增，因此不传或者传了YES都会默认返回包含合规结果数据，其他枚举值后面视情况丰富</p>
	IncludeCompliancePackRuleResult *string `json:"IncludeCompliancePackRuleResult,omitnil,omitempty" name:"IncludeCompliancePackRuleResult"`
}

type ListAggregateCompliancePacksRequest struct {
	*tchttp.BaseRequest
	
	// <p>数量</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>账号组ID</p>
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// <p>合规包名称</p>
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// <p>风险等级<br>1：高风险。<br>2：中风险。<br>3：低风险。</p>
	RiskLevel []*uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// <p>合规包状态 ACTIVE、NO_ACTIVE</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>评估状态合规： &#39;COMPLIANT&#39;<br>不合规： &#39;NON_COMPLIANT&#39;</p>
	ComplianceResult []*string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// <p>排序类型, 倒序：desc，顺序：asc</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>包含合规包结果定义</p><p>枚举值：</p><ul><li>NO： 不包含</li></ul><p>默认值：空</p><p>此字段为新增，因此不传或者传了YES都会默认返回包含合规结果数据，其他枚举值后面视情况丰富</p>
	IncludeCompliancePackRuleResult *string `json:"IncludeCompliancePackRuleResult,omitnil,omitempty" name:"IncludeCompliancePackRuleResult"`
}

func (r *ListAggregateCompliancePacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAggregateCompliancePacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "AccountGroupId")
	delete(f, "CompliancePackName")
	delete(f, "RiskLevel")
	delete(f, "Status")
	delete(f, "ComplianceResult")
	delete(f, "OrderType")
	delete(f, "IncludeCompliancePackRuleResult")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAggregateCompliancePacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAggregateCompliancePacksResponseParams struct {
	// <p>总数</p>
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// <p>详情</p>
	Items []*ConfigCompliancePack `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAggregateCompliancePacksResponse struct {
	*tchttp.BaseResponse
	Response *ListAggregateCompliancePacksResponseParams `json:"Response"`
}

func (r *ListAggregateCompliancePacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAggregateCompliancePacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAggregateConfigRuleEvaluationResultsRequestParams struct {
	// <p>规则ID</p>
	ConfigRuleId *string `json:"ConfigRuleId,omitnil,omitempty" name:"ConfigRuleId"`

	// <p>偏移量</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>当前页</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>账号组ID</p>
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// <p>类型</p>
	ResourceType []*string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>评估结果 COMPLIANT：合规   NON_COMPLIANT：不合规</p>
	ComplianceType []*string `json:"ComplianceType,omitnil,omitempty" name:"ComplianceType"`

	// <p>资源所属用户ID</p>
	ResourceOwnerId *uint64 `json:"ResourceOwnerId,omitnil,omitempty" name:"ResourceOwnerId"`

	// <p>筛选的资源拥有者uin集合</p>
	ResourceOwnerIds []*uint64 `json:"ResourceOwnerIds,omitnil,omitempty" name:"ResourceOwnerIds"`
}

type ListAggregateConfigRuleEvaluationResultsRequest struct {
	*tchttp.BaseRequest
	
	// <p>规则ID</p>
	ConfigRuleId *string `json:"ConfigRuleId,omitnil,omitempty" name:"ConfigRuleId"`

	// <p>偏移量</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>当前页</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>账号组ID</p>
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// <p>类型</p>
	ResourceType []*string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>评估结果 COMPLIANT：合规   NON_COMPLIANT：不合规</p>
	ComplianceType []*string `json:"ComplianceType,omitnil,omitempty" name:"ComplianceType"`

	// <p>资源所属用户ID</p>
	ResourceOwnerId *uint64 `json:"ResourceOwnerId,omitnil,omitempty" name:"ResourceOwnerId"`

	// <p>筛选的资源拥有者uin集合</p>
	ResourceOwnerIds []*uint64 `json:"ResourceOwnerIds,omitnil,omitempty" name:"ResourceOwnerIds"`
}

func (r *ListAggregateConfigRuleEvaluationResultsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAggregateConfigRuleEvaluationResultsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigRuleId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "AccountGroupId")
	delete(f, "ResourceType")
	delete(f, "ComplianceType")
	delete(f, "ResourceOwnerId")
	delete(f, "ResourceOwnerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAggregateConfigRuleEvaluationResultsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAggregateConfigRuleEvaluationResultsResponseParams struct {
	// <p>总数</p>
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// <p>详情</p>
	Items []*AggregateEvaluationResult `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAggregateConfigRuleEvaluationResultsResponse struct {
	*tchttp.BaseResponse
	Response *ListAggregateConfigRuleEvaluationResultsResponseParams `json:"Response"`
}

func (r *ListAggregateConfigRuleEvaluationResultsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAggregateConfigRuleEvaluationResultsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAggregateConfigRulesRequestParams struct {
	// <p>每页限制</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>账号组ID</p>
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// <p>排序类型, 倒序：desc，顺序：asc</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>风险等级<br>1：高风险。<br>2：中风险。<br>3：低风险。</p>
	RiskLevel []*uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// <p>规则状态</p>
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// <p>评估结果</p>
	ComplianceResult []*string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// <p>规则名</p>
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// <p>规则所属账号ID</p>
	RuleOwnerId *uint64 `json:"RuleOwnerId,omitnil,omitempty" name:"RuleOwnerId"`
}

type ListAggregateConfigRulesRequest struct {
	*tchttp.BaseRequest
	
	// <p>每页限制</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>账号组ID</p>
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// <p>排序类型, 倒序：desc，顺序：asc</p>
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// <p>风险等级<br>1：高风险。<br>2：中风险。<br>3：低风险。</p>
	RiskLevel []*uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// <p>规则状态</p>
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// <p>评估结果</p>
	ComplianceResult []*string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// <p>规则名</p>
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// <p>规则所属账号ID</p>
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
	// <p>总数</p>
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// <p>详情</p>
	Items []*ConfigRule `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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
type ListAggregateDiscoveredResourcesRequestParams struct {
	// 每页显示数量
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// resourceName：资源名  resourceId ：资源ID resourceType：资源类型
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 下一页token
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 排序方式 asc、desc
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

type ListAggregateDiscoveredResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 每页显示数量
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// resourceName：资源名  resourceId ：资源ID resourceType：资源类型
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 下一页token
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 排序方式 asc、desc
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

func (r *ListAggregateDiscoveredResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAggregateDiscoveredResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MaxResults")
	delete(f, "AccountGroupId")
	delete(f, "Filters")
	delete(f, "Tags")
	delete(f, "NextToken")
	delete(f, "OrderType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAggregateDiscoveredResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAggregateDiscoveredResourcesResponseParams struct {
	// 详情
	Items []*AggregateResourceInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 下一页
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAggregateDiscoveredResourcesResponse struct {
	*tchttp.BaseResponse
	Response *ListAggregateDiscoveredResourcesResponseParams `json:"Response"`
}

func (r *ListAggregateDiscoveredResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAggregateDiscoveredResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAggregatorsRequestParams struct {
	// 每页显示数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 起始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type ListAggregatorsRequest struct {
	*tchttp.BaseRequest
	
	// 每页显示数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 起始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *ListAggregatorsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAggregatorsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAggregatorsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAggregatorsResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 账号组列表
	Items []*Aggregator `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAggregatorsResponse struct {
	*tchttp.BaseResponse
	Response *ListAggregatorsResponseParams `json:"Response"`
}

func (r *ListAggregatorsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAggregatorsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAlarmPolicyRequestParams struct {
	// 页码
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type ListAlarmPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *ListAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAlarmPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAlarmPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAlarmPolicyResponseParams struct {
	// 返回记录的数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 告警策略返回值
	AlarmPolicyList []*AlarmPolicyRsp `json:"AlarmPolicyList,omitnil,omitempty" name:"AlarmPolicyList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ListAlarmPolicyResponseParams `json:"Response"`
}

func (r *ListAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCompliancePacksRequestParams struct {
	// 数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 合规包名称
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// 风险等级
	// 1：高风险。
	// 2：中风险。
	// 3：低风险。
	RiskLevel []*uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 合规包状态 ACTIVE、NO_ACTIVE
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 评估状态合规： 'COMPLIANT'
	// 不合规： 'NON_COMPLIANT'
	ComplianceResult []*string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// 排序类型, 倒序：desc，顺序：asc
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

type ListCompliancePacksRequest struct {
	*tchttp.BaseRequest
	
	// 数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 合规包名称
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// 风险等级
	// 1：高风险。
	// 2：中风险。
	// 3：低风险。
	RiskLevel []*uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 合规包状态 ACTIVE、NO_ACTIVE
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 评估状态合规： 'COMPLIANT'
	// 不合规： 'NON_COMPLIANT'
	ComplianceResult []*string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// 排序类型, 倒序：desc，顺序：asc
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

func (r *ListCompliancePacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCompliancePacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "CompliancePackName")
	delete(f, "RiskLevel")
	delete(f, "Status")
	delete(f, "ComplianceResult")
	delete(f, "OrderType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListCompliancePacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCompliancePacksResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 详情
	Items []*ConfigCompliancePack `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListCompliancePacksResponse struct {
	*tchttp.BaseResponse
	Response *ListCompliancePacksResponseParams `json:"Response"`
}

func (r *ListCompliancePacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCompliancePacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListConfigRuleEvaluationResultsRequestParams struct {
	// 规则ID
	ConfigRuleId *string `json:"ConfigRuleId,omitnil,omitempty" name:"ConfigRuleId"`

	// 偏移量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 当前页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 类型
	ResourceType []*string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 评估结果 COMPLIANT：合规   NON_COMPLIANT：不合规
	ComplianceType []*string `json:"ComplianceType,omitnil,omitempty" name:"ComplianceType"`
}

type ListConfigRuleEvaluationResultsRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	ConfigRuleId *string `json:"ConfigRuleId,omitnil,omitempty" name:"ConfigRuleId"`

	// 偏移量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 当前页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 类型
	ResourceType []*string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 评估结果 COMPLIANT：合规   NON_COMPLIANT：不合规
	ComplianceType []*string `json:"ComplianceType,omitnil,omitempty" name:"ComplianceType"`
}

func (r *ListConfigRuleEvaluationResultsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListConfigRuleEvaluationResultsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigRuleId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ResourceType")
	delete(f, "ComplianceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListConfigRuleEvaluationResultsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListConfigRuleEvaluationResultsResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 详情
	Items []*EvaluationResult `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListConfigRuleEvaluationResultsResponse struct {
	*tchttp.BaseResponse
	Response *ListConfigRuleEvaluationResultsResponseParams `json:"Response"`
}

func (r *ListConfigRuleEvaluationResultsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListConfigRuleEvaluationResultsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListConfigRulesRequestParams struct {
	// 每页数量。
	// 取值范围：1~200
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。
	// 取值范围：最小值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型(规则名称)。
	// 倒序：desc，
	// 顺序：asc
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 风险等级。
	// 1：高风险，
	// 2：中风险，
	// 3：低风险。
	RiskLevel []*uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 规则状态。
	// ACTIVE：启用
	// UN_ACTIVE：停用
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 评估结果。
	// COMPLIANT：合规
	// NON_COMPLIANT：不合规
	ComplianceResult []*string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`

	// 规则名
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

type ListConfigRulesRequest struct {
	*tchttp.BaseRequest
	
	// 每页数量。
	// 取值范围：1~200
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。
	// 取值范围：最小值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序类型(规则名称)。
	// 倒序：desc，
	// 顺序：asc
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 风险等级。
	// 1：高风险，
	// 2：中风险，
	// 3：低风险。
	RiskLevel []*uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 规则状态。
	// ACTIVE：启用
	// UN_ACTIVE：停用
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 评估结果。
	// COMPLIANT：合规
	// NON_COMPLIANT：不合规
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

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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
type ListDiscoveredResourcesRequestParams struct {
	// 每页显示数量
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// resourceName：资源名  resourceId ：资源ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 下一页token
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 排序方式 asc、desc
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

type ListDiscoveredResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 每页显示数量
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// resourceName：资源名  resourceId ：资源ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 标签
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 下一页token
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 排序方式 asc、desc
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

func (r *ListDiscoveredResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDiscoveredResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MaxResults")
	delete(f, "Filters")
	delete(f, "Tags")
	delete(f, "NextToken")
	delete(f, "OrderType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDiscoveredResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDiscoveredResourcesResponseParams struct {
	// 详情
	Items []*ResourceListInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 下一页
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDiscoveredResourcesResponse struct {
	*tchttp.BaseResponse
	Response *ListDiscoveredResourcesResponseParams `json:"Response"`
}

func (r *ListDiscoveredResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDiscoveredResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRemediationExecutionsRequestParams struct {
	// 规则 ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 分页条数。默认20， 取值1~200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 修正状态 1：运行中 2：成功 3：失败
	ExecutionStatus *uint64 `json:"ExecutionStatus,omitnil,omitempty" name:"ExecutionStatus"`
}

type ListRemediationExecutionsRequest struct {
	*tchttp.BaseRequest
	
	// 规则 ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 分页条数。默认20， 取值1~200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 修正状态 1：运行中 2：成功 3：失败
	ExecutionStatus *uint64 `json:"ExecutionStatus,omitnil,omitempty" name:"ExecutionStatus"`
}

func (r *ListRemediationExecutionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRemediationExecutionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ExecutionStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRemediationExecutionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRemediationExecutionsResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 修复记录
	RemediationExecutions []*RemediationExecutions `json:"RemediationExecutions,omitnil,omitempty" name:"RemediationExecutions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRemediationExecutionsResponse struct {
	*tchttp.BaseResponse
	Response *ListRemediationExecutionsResponseParams `json:"Response"`
}

func (r *ListRemediationExecutionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRemediationExecutionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRemediationsRequestParams struct {
	// 分页条数。默认20， 取值1~200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 规则ID
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

type ListRemediationsRequest struct {
	*tchttp.BaseRequest
	
	// 分页条数。默认20， 取值1~200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 规则ID
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

func (r *ListRemediationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRemediationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRemediationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRemediationsResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 修正设置
	Remediations []*Remediation `json:"Remediations,omitnil,omitempty" name:"Remediations"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRemediationsResponse struct {
	*tchttp.BaseResponse
	Response *ListRemediationsResponseParams `json:"Response"`
}

func (r *ListRemediationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRemediationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListResourceTypesRequestParams struct {

}

type ListResourceTypesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListResourceTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListResourceTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListResourceTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListResourceTypesResponseParams struct {
	// 详情
	Items []*ConfigResource `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListResourceTypesResponse struct {
	*tchttp.BaseResponse
	Response *ListResourceTypesResponseParams `json:"Response"`
}

func (r *ListResourceTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListResourceTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSystemCompliancePacksRequestParams struct {
	// 每页限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type ListSystemCompliancePacksRequest struct {
	*tchttp.BaseRequest
	
	// 每页限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *ListSystemCompliancePacksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSystemCompliancePacksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSystemCompliancePacksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSystemCompliancePacksResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*SystemCompliancePack `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSystemCompliancePacksResponse struct {
	*tchttp.BaseResponse
	Response *ListSystemCompliancePacksResponseParams `json:"Response"`
}

func (r *ListSystemCompliancePacksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSystemCompliancePacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSystemRulesRequestParams struct {
	// 每页数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 当前页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 搜索关键字。支持标识/名称/标签/描述搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 风险等级
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`
}

type ListSystemRulesRequest struct {
	*tchttp.BaseRequest
	
	// 每页数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 当前页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 搜索关键字。支持标识/名称/标签/描述搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 风险等级
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`
}

func (r *ListSystemRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSystemRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Keyword")
	delete(f, "RiskLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSystemRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSystemRulesResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 详情
	Items []*SystemConfigRule `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSystemRulesResponse struct {
	*tchttp.BaseResponse
	Response *ListSystemRulesResponseParams `json:"Response"`
}

func (r *ListSystemRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSystemRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenAggregateConfigRuleRequestParams struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

type OpenAggregateConfigRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

func (r *OpenAggregateConfigRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAggregateConfigRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "AccountGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenAggregateConfigRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenAggregateConfigRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenAggregateConfigRuleResponse struct {
	*tchttp.BaseResponse
	Response *OpenAggregateConfigRuleResponseParams `json:"Response"`
}

func (r *OpenAggregateConfigRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAggregateConfigRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenConfigRecorderRequestParams struct {

}

type OpenConfigRecorderRequest struct {
	*tchttp.BaseRequest
	
}

func (r *OpenConfigRecorderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenConfigRecorderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenConfigRecorderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenConfigRecorderResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenConfigRecorderResponse struct {
	*tchttp.BaseResponse
	Response *OpenConfigRecorderResponseParams `json:"Response"`
}

func (r *OpenConfigRecorderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenConfigRecorderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenConfigRuleRequestParams struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type OpenConfigRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

func (r *OpenConfigRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenConfigRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenConfigRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenConfigRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenConfigRuleResponse struct {
	*tchttp.BaseResponse
	Response *OpenConfigRuleResponseParams `json:"Response"`
}

func (r *OpenConfigRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenConfigRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PutEvaluationsRequestParams struct {
	// 回调令牌。从自定义规则所选的scf云函数入参中取参数ResultToken值
	// <a href="https://cloud.tencent.com/document/product/583/9210#.E5.87.BD.E6.95.B0.E5.85.A5.E5.8F.82.3Ca-id.3D.22input.22.3E.3C.2Fa.3E" target="_blank">云函数入参说明</a>
	ResultToken *string `json:"ResultToken,omitnil,omitempty" name:"ResultToken"`

	// 自定义规则评估结果信息。
	Evaluations []*Evaluation `json:"Evaluations,omitnil,omitempty" name:"Evaluations"`
}

type PutEvaluationsRequest struct {
	*tchttp.BaseRequest
	
	// 回调令牌。从自定义规则所选的scf云函数入参中取参数ResultToken值
	// <a href="https://cloud.tencent.com/document/product/583/9210#.E5.87.BD.E6.95.B0.E5.85.A5.E5.8F.82.3Ca-id.3D.22input.22.3E.3C.2Fa.3E" target="_blank">云函数入参说明</a>
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
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

type Remediation struct {
	// 修正模板 ID。
	RemediationTemplateId *string `json:"RemediationTemplateId,omitnil,omitempty" name:"RemediationTemplateId"`

	// 修正设置 ID。
	RemediationId *string `json:"RemediationId,omitnil,omitempty" name:"RemediationId"`

	// 执行修正的模板来源
	RemediationSourceType *string `json:"RemediationSourceType,omitnil,omitempty" name:"RemediationSourceType"`

	// 修正类型。取值：
	// SCF：函数计算（自定义修正）。
	RemediationType *string `json:"RemediationType,omitnil,omitempty" name:"RemediationType"`

	// 账号ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 修正执行方式。取值：
	// MANUAL_EXECUTION：手动执行。
	InvokeType *string `json:"InvokeType,omitnil,omitempty" name:"InvokeType"`

	// 规则 ID。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type RemediationExecutions struct {
	// 修正状态 1：运行中 2：成功 3：失败
	ExecutionStatus *uint64 `json:"ExecutionStatus,omitnil,omitempty" name:"ExecutionStatus"`

	// 资源类型
	ExecutionResourceType *string `json:"ExecutionResourceType,omitnil,omitempty" name:"ExecutionResourceType"`

	// 修复时间
	ExecutionCreateDate *string `json:"ExecutionCreateDate,omitnil,omitempty" name:"ExecutionCreateDate"`

	// 错误信息
	ExecutionStatusMessage *string `json:"ExecutionStatusMessage,omitnil,omitempty" name:"ExecutionStatusMessage"`

	// 资源ID
	ExecutionResourceIds *string `json:"ExecutionResourceIds,omitnil,omitempty" name:"ExecutionResourceIds"`
}

type ResourceListInfo struct {
	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源名称
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 地域
	ResourceRegion *string `json:"ResourceRegion,omitnil,omitempty" name:"ResourceRegion"`

	// 资源状态
	ResourceStatus *string `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`

	// 1 :已删除 2：未删除
	ResourceDelete *uint64 `json:"ResourceDelete,omitnil,omitempty" name:"ResourceDelete"`

	// 资源创建时间
	ResourceCreateTime *string `json:"ResourceCreateTime,omitnil,omitempty" name:"ResourceCreateTime"`

	// 标签信息
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 可用区
	ResourceZone *string `json:"ResourceZone,omitnil,omitempty" name:"ResourceZone"`

	// 合规状态
	ComplianceResult *string `json:"ComplianceResult,omitnil,omitempty" name:"ComplianceResult"`
}

type SourceConditionForManage struct {
	// 条件为空，合规：COMPLIANT，不合规：NON_COMPLIANT，无法应用：NOT_APPLICABLE
	EmptyAs *string `json:"EmptyAs,omitnil,omitempty" name:"EmptyAs"`

	// 配置路径
	SelectPath *string `json:"SelectPath,omitnil,omitempty" name:"SelectPath"`

	// 操作运算符
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 是否必须
	Required *bool `json:"Required,omitnil,omitempty" name:"Required"`

	// 期望值
	DesiredValue *string `json:"DesiredValue,omitnil,omitempty" name:"DesiredValue"`
}

// Predefined struct for user
type StartAggregateConfigRuleEvaluationRequestParams struct {
	// 手动触发：MANUAL  周期触发：SCHEDULE
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`
}

type StartAggregateConfigRuleEvaluationRequest struct {
	*tchttp.BaseRequest
	
	// 手动触发：MANUAL  周期触发：SCHEDULE
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`
}

func (r *StartAggregateConfigRuleEvaluationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartAggregateConfigRuleEvaluationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TriggerType")
	delete(f, "AccountGroupId")
	delete(f, "RuleId")
	delete(f, "CompliancePackId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartAggregateConfigRuleEvaluationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartAggregateConfigRuleEvaluationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartAggregateConfigRuleEvaluationResponse struct {
	*tchttp.BaseResponse
	Response *StartAggregateConfigRuleEvaluationResponseParams `json:"Response"`
}

func (r *StartAggregateConfigRuleEvaluationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartAggregateConfigRuleEvaluationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartConfigRuleEvaluationRequestParams struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`
}

type StartConfigRuleEvaluationRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`
}

func (r *StartConfigRuleEvaluationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartConfigRuleEvaluationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "CompliancePackId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartConfigRuleEvaluationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartConfigRuleEvaluationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartConfigRuleEvaluationResponse struct {
	*tchttp.BaseResponse
	Response *StartConfigRuleEvaluationResponseParams `json:"Response"`
}

func (r *StartConfigRuleEvaluationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartConfigRuleEvaluationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartRemediationRequestParams struct {
	// 规则 ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type StartRemediationRequest struct {
	*tchttp.BaseRequest
	
	// 规则 ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

func (r *StartRemediationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartRemediationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartRemediationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartRemediationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartRemediationResponse struct {
	*tchttp.BaseResponse
	Response *StartRemediationResponseParams `json:"Response"`
}

func (r *StartRemediationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartRemediationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemCompliancePack struct {
	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 名称
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 合规包风险等级1：高风险。
	// 2：中风险。
	// 3：低风险。
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 合规包规则信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigRules []*CompliancePackRuleForManage `json:"ConfigRules,omitnil,omitempty" name:"ConfigRules"`
}

type SystemConfigRule struct {
	// 规则标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Identifier *string `json:"Identifier,omitnil,omitempty" name:"Identifier"`

	// 规则名
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputParameter []*InputParameterForManage `json:"InputParameter,omitnil,omitempty" name:"InputParameter"`

	// 规则触发条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceCondition []*SourceConditionForManage `json:"SourceCondition,omitnil,omitempty" name:"SourceCondition"`

	// 支持的资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType []*string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label []*string `json:"Label,omitnil,omitempty" name:"Label"`

	// 风险等级，1，2，3
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 对应的函数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceFunction *string `json:"ServiceFunction,omitnil,omitempty" name:"ServiceFunction"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 触发类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerType []*string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 使用次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReferenceCount *uint64 `json:"ReferenceCount,omitnil,omitempty" name:"ReferenceCount"`

	// 规则类型
	IdentifierType *string `json:"IdentifierType,omitnil,omitempty" name:"IdentifierType"`
}

type Tag struct {
	// 标签key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签value
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TriggerType struct {
	// 触发类型
	MessageType *string `json:"MessageType,omitnil,omitempty" name:"MessageType"`

	// 触发时间周期
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitnil,omitempty" name:"MaximumExecutionFrequency"`
}

// Predefined struct for user
type UpdateAggregateCompliancePackRequestParams struct {
	// 合规包名称
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// 风险等级
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 合规包规则
	ConfigRules []*CompliancePackRule `json:"ConfigRules,omitnil,omitempty" name:"ConfigRules"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateAggregateCompliancePackRequest struct {
	*tchttp.BaseRequest
	
	// 合规包名称
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// 风险等级
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 合规包规则
	ConfigRules []*CompliancePackRule `json:"ConfigRules,omitnil,omitempty" name:"ConfigRules"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateAggregateCompliancePackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAggregateCompliancePackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompliancePackName")
	delete(f, "RiskLevel")
	delete(f, "CompliancePackId")
	delete(f, "ConfigRules")
	delete(f, "AccountGroupId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAggregateCompliancePackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAggregateCompliancePackResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAggregateCompliancePackResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAggregateCompliancePackResponseParams `json:"Response"`
}

func (r *UpdateAggregateCompliancePackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAggregateCompliancePackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAggregateCompliancePackStatusRequestParams struct {
	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// ACTIVE：启用
	// UN_ACTIVE ：停用
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

type UpdateAggregateCompliancePackStatusRequest struct {
	*tchttp.BaseRequest
	
	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// ACTIVE：启用
	// UN_ACTIVE ：停用
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`
}

func (r *UpdateAggregateCompliancePackStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAggregateCompliancePackStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompliancePackId")
	delete(f, "Status")
	delete(f, "AccountGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAggregateCompliancePackStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAggregateCompliancePackStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAggregateCompliancePackStatusResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAggregateCompliancePackStatusResponseParams `json:"Response"`
}

func (r *UpdateAggregateCompliancePackStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAggregateCompliancePackStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAggregateConfigDeliverRequestParams struct {
	// 0 关闭  1 开启
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 投递服务名称
	DeliverName *string `json:"DeliverName,omitnil,omitempty" name:"DeliverName"`

	// 资源六段式  
	// COS：qcs::cos:$region:$account:prefix/$appid/$BucketName
	// CLS:
	// qcs::cls:$region:$account:cls/topicId
	TargetArn *string `json:"TargetArn,omitnil,omitempty" name:"TargetArn"`

	// 资源前缀
	DeliverPrefix *string `json:"DeliverPrefix,omitnil,omitempty" name:"DeliverPrefix"`

	// 投递类型  COS CLS
	DeliverType *string `json:"DeliverType,omitnil,omitempty" name:"DeliverType"`

	// 支持跨账号投递的成员账号uin，只能是委派管理员。默认为0，即投递到管理员账号下
	DeliverUin *int64 `json:"DeliverUin,omitnil,omitempty" name:"DeliverUin"`

	// 1：配置变更 2： 资源列表 3：全选
	DeliverContentType *uint64 `json:"DeliverContentType,omitnil,omitempty" name:"DeliverContentType"`
}

type UpdateAggregateConfigDeliverRequest struct {
	*tchttp.BaseRequest
	
	// 0 关闭  1 开启
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 投递服务名称
	DeliverName *string `json:"DeliverName,omitnil,omitempty" name:"DeliverName"`

	// 资源六段式  
	// COS：qcs::cos:$region:$account:prefix/$appid/$BucketName
	// CLS:
	// qcs::cls:$region:$account:cls/topicId
	TargetArn *string `json:"TargetArn,omitnil,omitempty" name:"TargetArn"`

	// 资源前缀
	DeliverPrefix *string `json:"DeliverPrefix,omitnil,omitempty" name:"DeliverPrefix"`

	// 投递类型  COS CLS
	DeliverType *string `json:"DeliverType,omitnil,omitempty" name:"DeliverType"`

	// 支持跨账号投递的成员账号uin，只能是委派管理员。默认为0，即投递到管理员账号下
	DeliverUin *int64 `json:"DeliverUin,omitnil,omitempty" name:"DeliverUin"`

	// 1：配置变更 2： 资源列表 3：全选
	DeliverContentType *uint64 `json:"DeliverContentType,omitnil,omitempty" name:"DeliverContentType"`
}

func (r *UpdateAggregateConfigDeliverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAggregateConfigDeliverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "AccountGroupId")
	delete(f, "DeliverName")
	delete(f, "TargetArn")
	delete(f, "DeliverPrefix")
	delete(f, "DeliverType")
	delete(f, "DeliverUin")
	delete(f, "DeliverContentType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAggregateConfigDeliverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAggregateConfigDeliverResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAggregateConfigDeliverResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAggregateConfigDeliverResponseParams `json:"Response"`
}

func (r *UpdateAggregateConfigDeliverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAggregateConfigDeliverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAggregateConfigRuleRequestParams struct {
	// 触发类型，最多支持两种
	TriggerType []*TriggerType `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 风险等级
	// 1：高风险。
	// 2：中风险。
	// 3：低风险。
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 入参
	InputParameter []*InputParameter `json:"InputParameter,omitnil,omitempty" name:"InputParameter"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 关联地域
	RegionScope []*string `json:"RegionScope,omitnil,omitempty" name:"RegionScope"`

	// 关联标签
	TagsScope []*Tag `json:"TagsScope,omitnil,omitempty" name:"TagsScope"`

	// 排除的资源ID
	ExcludeResourceIdsScope []*string `json:"ExcludeResourceIdsScope,omitnil,omitempty" name:"ExcludeResourceIdsScope"`
}

type UpdateAggregateConfigRuleRequest struct {
	*tchttp.BaseRequest
	
	// 触发类型，最多支持两种
	TriggerType []*TriggerType `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 风险等级
	// 1：高风险。
	// 2：中风险。
	// 3：低风险。
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 账号组ID
	AccountGroupId *string `json:"AccountGroupId,omitnil,omitempty" name:"AccountGroupId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 入参
	InputParameter []*InputParameter `json:"InputParameter,omitnil,omitempty" name:"InputParameter"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 关联地域
	RegionScope []*string `json:"RegionScope,omitnil,omitempty" name:"RegionScope"`

	// 关联标签
	TagsScope []*Tag `json:"TagsScope,omitnil,omitempty" name:"TagsScope"`

	// 排除的资源ID
	ExcludeResourceIdsScope []*string `json:"ExcludeResourceIdsScope,omitnil,omitempty" name:"ExcludeResourceIdsScope"`
}

func (r *UpdateAggregateConfigRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAggregateConfigRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TriggerType")
	delete(f, "RiskLevel")
	delete(f, "RuleId")
	delete(f, "AccountGroupId")
	delete(f, "RuleName")
	delete(f, "InputParameter")
	delete(f, "Description")
	delete(f, "RegionScope")
	delete(f, "TagsScope")
	delete(f, "ExcludeResourceIdsScope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAggregateConfigRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAggregateConfigRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAggregateConfigRuleResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAggregateConfigRuleResponseParams `json:"Response"`
}

func (r *UpdateAggregateConfigRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAggregateConfigRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAlarmPolicyRequestParams struct {
	// 告警策略id
	AlarmPolicyId *uint64 `json:"AlarmPolicyId,omitnil,omitempty" name:"AlarmPolicyId"`

	// 告警策略名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 事件范围  1：当前账号  2：多账号
	EventScope []*int64 `json:"EventScope,omitnil,omitempty" name:"EventScope"`

	// 风险等级 1：高风险  2：中风险 3：低风险
	RiskLevel []*int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 通知时间段
	NoticeTime *string `json:"NoticeTime,omitnil,omitempty" name:"NoticeTime"`

	// 通知机制
	NotificationMechanism *string `json:"NotificationMechanism,omitnil,omitempty" name:"NotificationMechanism"`

	// 状态 1：启用 2：停用
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 通知周期,1-7数字分别代表周一到周日
	NoticePeriod []*int64 `json:"NoticePeriod,omitnil,omitempty" name:"NoticePeriod"`

	// 策略描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateAlarmPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 告警策略id
	AlarmPolicyId *uint64 `json:"AlarmPolicyId,omitnil,omitempty" name:"AlarmPolicyId"`

	// 告警策略名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 事件范围  1：当前账号  2：多账号
	EventScope []*int64 `json:"EventScope,omitnil,omitempty" name:"EventScope"`

	// 风险等级 1：高风险  2：中风险 3：低风险
	RiskLevel []*int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 通知时间段
	NoticeTime *string `json:"NoticeTime,omitnil,omitempty" name:"NoticeTime"`

	// 通知机制
	NotificationMechanism *string `json:"NotificationMechanism,omitnil,omitempty" name:"NotificationMechanism"`

	// 状态 1：启用 2：停用
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 通知周期,1-7数字分别代表周一到周日
	NoticePeriod []*int64 `json:"NoticePeriod,omitnil,omitempty" name:"NoticePeriod"`

	// 策略描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAlarmPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmPolicyId")
	delete(f, "Name")
	delete(f, "EventScope")
	delete(f, "RiskLevel")
	delete(f, "NoticeTime")
	delete(f, "NotificationMechanism")
	delete(f, "Status")
	delete(f, "NoticePeriod")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAlarmPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAlarmPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAlarmPolicyResponseParams `json:"Response"`
}

func (r *UpdateAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCompliancePackRequestParams struct {
	// 合规包名称
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// 风险等级
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 合规包规则
	ConfigRules []*CompliancePackRule `json:"ConfigRules,omitnil,omitempty" name:"ConfigRules"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateCompliancePackRequest struct {
	*tchttp.BaseRequest
	
	// 合规包名称
	CompliancePackName *string `json:"CompliancePackName,omitnil,omitempty" name:"CompliancePackName"`

	// 风险等级
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// 合规包规则
	ConfigRules []*CompliancePackRule `json:"ConfigRules,omitnil,omitempty" name:"ConfigRules"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateCompliancePackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCompliancePackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompliancePackName")
	delete(f, "RiskLevel")
	delete(f, "CompliancePackId")
	delete(f, "ConfigRules")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCompliancePackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCompliancePackResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCompliancePackResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCompliancePackResponseParams `json:"Response"`
}

func (r *UpdateCompliancePackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCompliancePackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCompliancePackStatusRequestParams struct {
	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// ACTIVE：启用
	// UN_ACTIVE ：停用
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type UpdateCompliancePackStatusRequest struct {
	*tchttp.BaseRequest
	
	// 合规包ID
	CompliancePackId *string `json:"CompliancePackId,omitnil,omitempty" name:"CompliancePackId"`

	// ACTIVE：启用
	// UN_ACTIVE ：停用
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *UpdateCompliancePackStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCompliancePackStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompliancePackId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCompliancePackStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCompliancePackStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCompliancePackStatusResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCompliancePackStatusResponseParams `json:"Response"`
}

func (r *UpdateCompliancePackStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCompliancePackStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateConfigDeliverRequestParams struct {
	// 0 关闭  1 开启
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 投递服务名称
	DeliverName *string `json:"DeliverName,omitnil,omitempty" name:"DeliverName"`

	// 资源六段式  
	// COS：qcs::cos:$region:$account:prefix/$appid/$BucketName
	// CLS:
	// qcs::cls:$region:$account:cls/topicId
	TargetArn *string `json:"TargetArn,omitnil,omitempty" name:"TargetArn"`

	// clonfig_fix
	DeliverPrefix *string `json:"DeliverPrefix,omitnil,omitempty" name:"DeliverPrefix"`

	// 投递类型
	DeliverType *string `json:"DeliverType,omitnil,omitempty" name:"DeliverType"`

	// 1：配置变更 2： 资源列表 3：全选
	DeliverContentType *uint64 `json:"DeliverContentType,omitnil,omitempty" name:"DeliverContentType"`
}

type UpdateConfigDeliverRequest struct {
	*tchttp.BaseRequest
	
	// 0 关闭  1 开启
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 投递服务名称
	DeliverName *string `json:"DeliverName,omitnil,omitempty" name:"DeliverName"`

	// 资源六段式  
	// COS：qcs::cos:$region:$account:prefix/$appid/$BucketName
	// CLS:
	// qcs::cls:$region:$account:cls/topicId
	TargetArn *string `json:"TargetArn,omitnil,omitempty" name:"TargetArn"`

	// clonfig_fix
	DeliverPrefix *string `json:"DeliverPrefix,omitnil,omitempty" name:"DeliverPrefix"`

	// 投递类型
	DeliverType *string `json:"DeliverType,omitnil,omitempty" name:"DeliverType"`

	// 1：配置变更 2： 资源列表 3：全选
	DeliverContentType *uint64 `json:"DeliverContentType,omitnil,omitempty" name:"DeliverContentType"`
}

func (r *UpdateConfigDeliverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateConfigDeliverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "DeliverName")
	delete(f, "TargetArn")
	delete(f, "DeliverPrefix")
	delete(f, "DeliverType")
	delete(f, "DeliverContentType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateConfigDeliverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateConfigDeliverResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateConfigDeliverResponse struct {
	*tchttp.BaseResponse
	Response *UpdateConfigDeliverResponseParams `json:"Response"`
}

func (r *UpdateConfigDeliverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateConfigDeliverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateConfigRecorderRequestParams struct {
	// 资源类型列表
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`
}

type UpdateConfigRecorderRequest struct {
	*tchttp.BaseRequest
	
	// 资源类型列表
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`
}

func (r *UpdateConfigRecorderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateConfigRecorderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateConfigRecorderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateConfigRecorderResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateConfigRecorderResponse struct {
	*tchttp.BaseResponse
	Response *UpdateConfigRecorderResponseParams `json:"Response"`
}

func (r *UpdateConfigRecorderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateConfigRecorderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateConfigRuleRequestParams struct {
	// 触发类型，最多支持两种
	TriggerType []*TriggerType `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 风险等级
	// 1：高风险。
	// 2：中风险。
	// 3：低风险。
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 入参
	InputParameter []*InputParameter `json:"InputParameter,omitnil,omitempty" name:"InputParameter"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 规则评估地域范围，规则仅对指定地域中的资源生效。
	// 支持的地域范围config:ListResourceRegions返回的地域
	RegionsScope []*string `json:"RegionsScope,omitnil,omitempty" name:"RegionsScope"`

	// 规则评估标签范围，规则仅对绑定指定标签的资源生效。
	TagsScope []*Tag `json:"TagsScope,omitnil,omitempty" name:"TagsScope"`

	// 规则对指定资源ID无效，即不对该资源执行评估。
	ExcludeResourceIdsScope []*string `json:"ExcludeResourceIdsScope,omitnil,omitempty" name:"ExcludeResourceIdsScope"`
}

type UpdateConfigRuleRequest struct {
	*tchttp.BaseRequest
	
	// 触发类型，最多支持两种
	TriggerType []*TriggerType `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 风险等级
	// 1：高风险。
	// 2：中风险。
	// 3：低风险。
	RiskLevel *uint64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 入参
	InputParameter []*InputParameter `json:"InputParameter,omitnil,omitempty" name:"InputParameter"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 规则评估地域范围，规则仅对指定地域中的资源生效。
	// 支持的地域范围config:ListResourceRegions返回的地域
	RegionsScope []*string `json:"RegionsScope,omitnil,omitempty" name:"RegionsScope"`

	// 规则评估标签范围，规则仅对绑定指定标签的资源生效。
	TagsScope []*Tag `json:"TagsScope,omitnil,omitempty" name:"TagsScope"`

	// 规则对指定资源ID无效，即不对该资源执行评估。
	ExcludeResourceIdsScope []*string `json:"ExcludeResourceIdsScope,omitnil,omitempty" name:"ExcludeResourceIdsScope"`
}

func (r *UpdateConfigRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateConfigRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TriggerType")
	delete(f, "RiskLevel")
	delete(f, "RuleId")
	delete(f, "RuleName")
	delete(f, "InputParameter")
	delete(f, "Description")
	delete(f, "RegionsScope")
	delete(f, "TagsScope")
	delete(f, "ExcludeResourceIdsScope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateConfigRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateConfigRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateConfigRuleResponse struct {
	*tchttp.BaseResponse
	Response *UpdateConfigRuleResponseParams `json:"Response"`
}

func (r *UpdateConfigRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateConfigRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRemediationRequestParams struct {
	// 修正设置 ID
	RemediationId *string `json:"RemediationId,omitnil,omitempty" name:"RemediationId"`

	// 修正类型。取值： SCF：函数计算（自定义修正）。
	RemediationType *string `json:"RemediationType,omitnil,omitempty" name:"RemediationType"`

	// 修正模板 ID
	RemediationTemplateId *string `json:"RemediationTemplateId,omitnil,omitempty" name:"RemediationTemplateId"`

	// 修正执行方式。取值：  NON_EXECUTION：不执行。  AUTO_EXECUTION：自动执行。  MANUAL_EXECUTION：手动执行。  NOT_CONFIG：未设置。
	InvokeType *string `json:"InvokeType,omitnil,omitempty" name:"InvokeType"`

	// 执行修正来源
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`
}

type UpdateRemediationRequest struct {
	*tchttp.BaseRequest
	
	// 修正设置 ID
	RemediationId *string `json:"RemediationId,omitnil,omitempty" name:"RemediationId"`

	// 修正类型。取值： SCF：函数计算（自定义修正）。
	RemediationType *string `json:"RemediationType,omitnil,omitempty" name:"RemediationType"`

	// 修正模板 ID
	RemediationTemplateId *string `json:"RemediationTemplateId,omitnil,omitempty" name:"RemediationTemplateId"`

	// 修正执行方式。取值：  NON_EXECUTION：不执行。  AUTO_EXECUTION：自动执行。  MANUAL_EXECUTION：手动执行。  NOT_CONFIG：未设置。
	InvokeType *string `json:"InvokeType,omitnil,omitempty" name:"InvokeType"`

	// 执行修正来源
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`
}

func (r *UpdateRemediationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRemediationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RemediationId")
	delete(f, "RemediationType")
	delete(f, "RemediationTemplateId")
	delete(f, "InvokeType")
	delete(f, "SourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRemediationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRemediationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateRemediationResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRemediationResponseParams `json:"Response"`
}

func (r *UpdateRemediationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRemediationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserConfigResource struct {
	// 产品
	// 注意：此字段可能返回 null，表示取不到有效值。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 产品名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源类型名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceTypeName *string `json:"ResourceTypeName,omitnil,omitempty" name:"ResourceTypeName"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}