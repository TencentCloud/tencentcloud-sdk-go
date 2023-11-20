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

package v20180125

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AccessFullTextInfo struct {
	// 是否大小写敏感
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaseSensitive *bool `json:"CaseSensitive,omitnil" name:"CaseSensitive"`

	// 全文索引的分词符，字符串中每个字符代表一个分词符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tokenizer *string `json:"Tokenizer,omitnil" name:"Tokenizer"`

	// 是否包含中文
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainZH *bool `json:"ContainZH,omitnil" name:"ContainZH"`
}

type AccessHistogramItem struct {
	// 时间，单位ms
	//
	// Deprecated: BTime is deprecated.
	BTime *int64 `json:"BTime,omitnil" name:"BTime"`

	// 日志条数
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 时间，单位ms
	BeginTime *int64 `json:"BeginTime,omitnil" name:"BeginTime"`
}

type AccessKeyValueInfo struct {
	// 需要配置键值或者元字段索引的字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 字段的索引描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *AccessValueInfo `json:"Value,omitnil" name:"Value"`
}

type AccessLogInfo struct {
	// 日志时间，单位ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *int64 `json:"Time,omitnil" name:"Time"`

	// 日志主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 日志主题名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// 日志来源IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil" name:"Source"`

	// 日志文件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// 日志上报请求包的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgId *string `json:"PkgId,omitnil" name:"PkgId"`

	// 请求包内日志的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgLogId *string `json:"PkgLogId,omitnil" name:"PkgLogId"`

	// 日志内容的Json序列化字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogJson *string `json:"LogJson,omitnil" name:"LogJson"`
}

type AccessLogItem struct {
	// 日记Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 日志Value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type AccessLogItems struct {
	// 分析结果返回的KV数据对
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*AccessLogItem `json:"Data,omitnil" name:"Data"`
}

type AccessRuleInfo struct {
	// 全文索引配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullText *AccessFullTextInfo `json:"FullText,omitnil" name:"FullText"`

	// 键值索引配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyValue *AccessRuleKeyValueInfo `json:"KeyValue,omitnil" name:"KeyValue"`

	// 元字段索引配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *AccessRuleTagInfo `json:"Tag,omitnil" name:"Tag"`
}

type AccessRuleKeyValueInfo struct {
	// 是否大小写敏感
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaseSensitive *bool `json:"CaseSensitive,omitnil" name:"CaseSensitive"`

	// 需要建立索引的键值对信息；最大只能配置100个键值对
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyValues []*AccessKeyValueInfo `json:"KeyValues,omitnil" name:"KeyValues"`
}

type AccessRuleTagInfo struct {
	// 是否大小写敏感
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaseSensitive *bool `json:"CaseSensitive,omitnil" name:"CaseSensitive"`

	// 标签索引配置中的字段信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyValues []*AccessKeyValueInfo `json:"KeyValues,omitnil" name:"KeyValues"`
}

type AccessValueInfo struct {
	// 字段类型，目前支持的类型有：long、text、double
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 字段的分词符，只有当字段类型为text时才有意义；输入字符串中的每个字符代表一个分词符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tokenizer *string `json:"Tokenizer,omitnil" name:"Tokenizer"`

	// 字段是否开启分析功能
	// 注意：此字段可能返回 null，表示取不到有效值。
	SqlFlag *bool `json:"SqlFlag,omitnil" name:"SqlFlag"`

	// 是否包含中文
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainZH *bool `json:"ContainZH,omitnil" name:"ContainZH"`
}

// Predefined struct for user
type AddAntiFakeUrlRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// uri
	Uri *string `json:"Uri,omitnil" name:"Uri"`
}

type AddAntiFakeUrlRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// uri
	Uri *string `json:"Uri,omitnil" name:"Uri"`
}

func (r *AddAntiFakeUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAntiFakeUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Name")
	delete(f, "Uri")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddAntiFakeUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAntiFakeUrlResponseParams struct {
	// 结果
	Result *string `json:"Result,omitnil" name:"Result"`

	// 规则ID
	Id *string `json:"Id,omitnil" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddAntiFakeUrlResponse struct {
	*tchttp.BaseResponse
	Response *AddAntiFakeUrlResponseParams `json:"Response"`
}

func (r *AddAntiFakeUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAntiFakeUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAntiInfoLeakRulesRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 规则名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 动作，0（告警）、1（替换）、2（仅显示前四位）、3（仅显示后四位）、4（阻断）
	ActionType *uint64 `json:"ActionType,omitnil" name:"ActionType"`

	// 策略详情
	Strategies []*StrategyForAntiInfoLeak `json:"Strategies,omitnil" name:"Strategies"`

	// 网址
	Uri *string `json:"Uri,omitnil" name:"Uri"`
}

type AddAntiInfoLeakRulesRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 规则名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 动作，0（告警）、1（替换）、2（仅显示前四位）、3（仅显示后四位）、4（阻断）
	ActionType *uint64 `json:"ActionType,omitnil" name:"ActionType"`

	// 策略详情
	Strategies []*StrategyForAntiInfoLeak `json:"Strategies,omitnil" name:"Strategies"`

	// 网址
	Uri *string `json:"Uri,omitnil" name:"Uri"`
}

func (r *AddAntiInfoLeakRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAntiInfoLeakRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Name")
	delete(f, "ActionType")
	delete(f, "Strategies")
	delete(f, "Uri")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddAntiInfoLeakRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAntiInfoLeakRulesResponseParams struct {
	// 规则ID
	RuleId *int64 `json:"RuleId,omitnil" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddAntiInfoLeakRulesResponse struct {
	*tchttp.BaseResponse
	Response *AddAntiInfoLeakRulesResponseParams `json:"Response"`
}

func (r *AddAntiInfoLeakRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAntiInfoLeakRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAttackWhiteRuleRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 规则Id
	SignatureId *string `json:"SignatureId,omitnil" name:"SignatureId"`

	// 规则状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 匹配规则项列表
	Rules []*UserWhiteRuleItem `json:"Rules,omitnil" name:"Rules"`

	// 规则序号
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`
}

type AddAttackWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 规则Id
	SignatureId *string `json:"SignatureId,omitnil" name:"SignatureId"`

	// 规则状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 匹配规则项列表
	Rules []*UserWhiteRuleItem `json:"Rules,omitnil" name:"Rules"`

	// 规则序号
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`
}

func (r *AddAttackWhiteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAttackWhiteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "SignatureId")
	delete(f, "Status")
	delete(f, "Rules")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddAttackWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAttackWhiteRuleResponseParams struct {
	// 规则总数
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddAttackWhiteRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddAttackWhiteRuleResponseParams `json:"Response"`
}

func (r *AddAttackWhiteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAttackWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCustomRuleRequestParams struct {
	// 规则名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 优先级
	SortId *string `json:"SortId,omitnil" name:"SortId"`

	// 过期时间，单位为秒级时间戳，例如1677254399表示过期时间为2023-02-24 23:59:59. 0表示永不过期
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 策略详情
	Strategies []*Strategy `json:"Strategies,omitnil" name:"Strategies"`

	// 需要添加策略的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 动作类型，1代表阻断，2代表人机识别，3代表观察，4代表重定向
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// 如果动作是重定向，则表示重定向的地址；其他情况可以为空
	Redirect *string `json:"Redirect,omitnil" name:"Redirect"`

	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 放行的详情
	Bypass *string `json:"Bypass,omitnil" name:"Bypass"`

	// 添加规则的来源，默认为空
	EventId *string `json:"EventId,omitnil" name:"EventId"`
}

type AddCustomRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 优先级
	SortId *string `json:"SortId,omitnil" name:"SortId"`

	// 过期时间，单位为秒级时间戳，例如1677254399表示过期时间为2023-02-24 23:59:59. 0表示永不过期
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 策略详情
	Strategies []*Strategy `json:"Strategies,omitnil" name:"Strategies"`

	// 需要添加策略的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 动作类型，1代表阻断，2代表人机识别，3代表观察，4代表重定向
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// 如果动作是重定向，则表示重定向的地址；其他情况可以为空
	Redirect *string `json:"Redirect,omitnil" name:"Redirect"`

	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 放行的详情
	Bypass *string `json:"Bypass,omitnil" name:"Bypass"`

	// 添加规则的来源，默认为空
	EventId *string `json:"EventId,omitnil" name:"EventId"`
}

func (r *AddCustomRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCustomRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "SortId")
	delete(f, "ExpireTime")
	delete(f, "Strategies")
	delete(f, "Domain")
	delete(f, "ActionType")
	delete(f, "Redirect")
	delete(f, "Edition")
	delete(f, "Bypass")
	delete(f, "EventId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCustomRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCustomRuleResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	Success *ResponseCode `json:"Success,omitnil" name:"Success"`

	// 添加成功的规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *int64 `json:"RuleId,omitnil" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddCustomRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddCustomRuleResponseParams `json:"Response"`
}

func (r *AddCustomRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCustomRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCustomWhiteRuleRequestParams struct {
	// 规则名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 优先级
	SortId *string `json:"SortId,omitnil" name:"SortId"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 策略详情
	Strategies []*Strategy `json:"Strategies,omitnil" name:"Strategies"`

	// 需要添加策略的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 放行的详情
	Bypass *string `json:"Bypass,omitnil" name:"Bypass"`
}

type AddCustomWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 优先级
	SortId *string `json:"SortId,omitnil" name:"SortId"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 策略详情
	Strategies []*Strategy `json:"Strategies,omitnil" name:"Strategies"`

	// 需要添加策略的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 放行的详情
	Bypass *string `json:"Bypass,omitnil" name:"Bypass"`
}

func (r *AddCustomWhiteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCustomWhiteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "SortId")
	delete(f, "ExpireTime")
	delete(f, "Strategies")
	delete(f, "Domain")
	delete(f, "Bypass")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCustomWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCustomWhiteRuleResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	Success *ResponseCode `json:"Success,omitnil" name:"Success"`

	// 添加成功的规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddCustomWhiteRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddCustomWhiteRuleResponseParams `json:"Response"`
}

func (r *AddCustomWhiteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCustomWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddDomainWhiteRuleRequestParams struct {
	// 需要添加的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 需要添加的规则
	Rules []*uint64 `json:"Rules,omitnil" name:"Rules"`

	// 需要添加的规则url
	Url *string `json:"Url,omitnil" name:"Url"`

	// 规则的方法
	Function *string `json:"Function,omitnil" name:"Function"`

	// 规则的开关，0表示规则关闭，1表示规则打开
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

type AddDomainWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要添加的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 需要添加的规则
	Rules []*uint64 `json:"Rules,omitnil" name:"Rules"`

	// 需要添加的规则url
	Url *string `json:"Url,omitnil" name:"Url"`

	// 规则的方法
	Function *string `json:"Function,omitnil" name:"Function"`

	// 规则的开关，0表示规则关闭，1表示规则打开
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

func (r *AddDomainWhiteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDomainWhiteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Rules")
	delete(f, "Url")
	delete(f, "Function")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddDomainWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddDomainWhiteRuleResponseParams struct {
	// 规则id
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddDomainWhiteRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddDomainWhiteRuleResponseParams `json:"Response"`
}

func (r *AddDomainWhiteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDomainWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddSpartaProtectionRequestParams struct {
	// 需要防护的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 证书类型。
	// 0：仅配置HTTP监听端口，没有证书
	// 1：证书来源为自有证书
	// 2：证书来源为托管证书
	CertType *int64 `json:"CertType,omitnil" name:"CertType"`

	// waf前是否部署有七层代理服务。
	// 0：没有部署代理服务
	// 1：有部署代理服务，waf将使用XFF获取客户端IP
	// 2：有部署代理服务，waf将使用remote_addr获取客户端IP
	// 3：有部署代理服务，waf将使用ip_headers中的自定义header获取客户端IP
	IsCdn *int64 `json:"IsCdn,omitnil" name:"IsCdn"`

	// 回源类型。
	// 0：通过IP回源
	// 1：通过域名回源
	UpstreamType *int64 `json:"UpstreamType,omitnil" name:"UpstreamType"`

	// 是否开启WebSocket支持。
	// 0：关闭
	// 1：开启
	IsWebsocket *int64 `json:"IsWebsocket,omitnil" name:"IsWebsocket"`

	// 回源负载均衡策略。
	// 0：轮询
	// 1：IP hash
	// 2：加权轮询
	LoadBalance *string `json:"LoadBalance,omitnil" name:"LoadBalance"`

	// CertType为1时，需要填充此参数，表示自有证书的证书链
	Cert *string `json:"Cert,omitnil" name:"Cert"`

	// CertType为1时，需要填充此参数，表示自有证书的私钥
	PrivateKey *string `json:"PrivateKey,omitnil" name:"PrivateKey"`

	// CertType为2时，需要填充此参数，表示腾讯云SSL平台托管的证书id
	SSLId *string `json:"SSLId,omitnil" name:"SSLId"`

	// 待废弃，可不填。Waf的资源ID。
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// IsCdn为3时，需要填此参数，表示自定义header
	IpHeaders []*string `json:"IpHeaders,omitnil" name:"IpHeaders"`

	// 服务配置有HTTPS端口时，HTTPS的回源协议。
	// http：使用http协议回源，和HttpsUpstreamPort配合使用
	// https：使用https协议回源
	UpstreamScheme *string `json:"UpstreamScheme,omitnil" name:"UpstreamScheme"`

	// HTTPS回源端口,仅UpstreamScheme为http时需要填当前字段
	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitnil" name:"HttpsUpstreamPort"`

	// 待废弃，可不填。是否开启灰度，0表示不开启灰度。
	IsGray *int64 `json:"IsGray,omitnil" name:"IsGray"`

	// 待废弃，可不填。灰度的地区
	GrayAreas []*string `json:"GrayAreas,omitnil" name:"GrayAreas"`

	// 是否开启HTTP强制跳转到HTTPS。
	// 0：不强制跳转
	// 1：开启强制跳转
	HttpsRewrite *int64 `json:"HttpsRewrite,omitnil" name:"HttpsRewrite"`

	// 域名回源时的回源域名。UpstreamType为1时，需要填充此字段
	UpstreamDomain *string `json:"UpstreamDomain,omitnil" name:"UpstreamDomain"`

	// IP回源时的回源IP列表。UpstreamType为0时，需要填充此字段
	SrcList []*string `json:"SrcList,omitnil" name:"SrcList"`

	// 是否开启HTTP2，需要开启HTTPS协议支持。
	// 0：关闭
	// 1：开启
	IsHttp2 *int64 `json:"IsHttp2,omitnil" name:"IsHttp2"`

	// 服务端口列表配置。
	// NginxServerId：新增域名时填'0'
	// Port：监听端口号
	// Protocol：端口协议
	// UpstreamPort：与Port相同
	// UpstreamProtocol：与Protocol相同
	Ports []*PortItem `json:"Ports,omitnil" name:"Ports"`

	// 待废弃，可不填。WAF实例类型。
	// sparta-waf：SAAS型WAF
	// clb-waf：负载均衡型WAF
	// cdn-waf：CDN上的Web防护能力
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 是否开启长连接。
	// 0： 短连接
	// 1： 长连接
	IsKeepAlive *string `json:"IsKeepAlive,omitnil" name:"IsKeepAlive"`

	// 域名所属实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 待废弃，目前填0即可。anycast IP类型开关： 0 普通IP 1 Anycast IP
	Anycast *int64 `json:"Anycast,omitnil" name:"Anycast"`

	// 回源IP列表各IP的权重，和SrcList一一对应。当且仅当UpstreamType为0，并且SrcList有多个IP，并且LoadBalance为2时需要填写，否则填 []
	Weights []*int64 `json:"Weights,omitnil" name:"Weights"`

	// 是否开启主动健康检测。
	// 0：不开启
	// 1：开启
	ActiveCheck *int64 `json:"ActiveCheck,omitnil" name:"ActiveCheck"`

	// TLS版本信息
	TLSVersion *int64 `json:"TLSVersion,omitnil" name:"TLSVersion"`

	// 加密套件模板。
	// 0：不支持选择，使用默认模版  
	// 1：通用型模版 
	// 2：安全型模版 
	// 3：自定义模版
	CipherTemplate *int64 `json:"CipherTemplate,omitnil" name:"CipherTemplate"`

	// 自定义的加密套件列表。CipherTemplate为3时需要填此字段，表示自定义的加密套件，值通过DescribeCiphersDetail接口获取。
	Ciphers []*int64 `json:"Ciphers,omitnil" name:"Ciphers"`

	// WAF与源站的读超时时间，默认300s。
	ProxyReadTimeout *int64 `json:"ProxyReadTimeout,omitnil" name:"ProxyReadTimeout"`

	// WAF与源站的写超时时间，默认300s。
	ProxySendTimeout *int64 `json:"ProxySendTimeout,omitnil" name:"ProxySendTimeout"`

	// WAF回源时的SNI类型。
	// 0：关闭SNI，不配置client_hello中的server_name
	// 1：开启SNI，client_hello中的server_name为防护域名
	// 2：开启SNI，SNI为域名回源时的源站域名
	// 3：开启SNI，SNI为自定义域名
	SniType *int64 `json:"SniType,omitnil" name:"SniType"`

	// SniType为3时，需要填此参数，表示自定义的SNI；
	SniHost *string `json:"SniHost,omitnil" name:"SniHost"`

	// 是否开启XFF重置。
	// 0：关闭
	// 1：开启
	XFFReset *int64 `json:"XFFReset,omitnil" name:"XFFReset"`
}

type AddSpartaProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 需要防护的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 证书类型。
	// 0：仅配置HTTP监听端口，没有证书
	// 1：证书来源为自有证书
	// 2：证书来源为托管证书
	CertType *int64 `json:"CertType,omitnil" name:"CertType"`

	// waf前是否部署有七层代理服务。
	// 0：没有部署代理服务
	// 1：有部署代理服务，waf将使用XFF获取客户端IP
	// 2：有部署代理服务，waf将使用remote_addr获取客户端IP
	// 3：有部署代理服务，waf将使用ip_headers中的自定义header获取客户端IP
	IsCdn *int64 `json:"IsCdn,omitnil" name:"IsCdn"`

	// 回源类型。
	// 0：通过IP回源
	// 1：通过域名回源
	UpstreamType *int64 `json:"UpstreamType,omitnil" name:"UpstreamType"`

	// 是否开启WebSocket支持。
	// 0：关闭
	// 1：开启
	IsWebsocket *int64 `json:"IsWebsocket,omitnil" name:"IsWebsocket"`

	// 回源负载均衡策略。
	// 0：轮询
	// 1：IP hash
	// 2：加权轮询
	LoadBalance *string `json:"LoadBalance,omitnil" name:"LoadBalance"`

	// CertType为1时，需要填充此参数，表示自有证书的证书链
	Cert *string `json:"Cert,omitnil" name:"Cert"`

	// CertType为1时，需要填充此参数，表示自有证书的私钥
	PrivateKey *string `json:"PrivateKey,omitnil" name:"PrivateKey"`

	// CertType为2时，需要填充此参数，表示腾讯云SSL平台托管的证书id
	SSLId *string `json:"SSLId,omitnil" name:"SSLId"`

	// 待废弃，可不填。Waf的资源ID。
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// IsCdn为3时，需要填此参数，表示自定义header
	IpHeaders []*string `json:"IpHeaders,omitnil" name:"IpHeaders"`

	// 服务配置有HTTPS端口时，HTTPS的回源协议。
	// http：使用http协议回源，和HttpsUpstreamPort配合使用
	// https：使用https协议回源
	UpstreamScheme *string `json:"UpstreamScheme,omitnil" name:"UpstreamScheme"`

	// HTTPS回源端口,仅UpstreamScheme为http时需要填当前字段
	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitnil" name:"HttpsUpstreamPort"`

	// 待废弃，可不填。是否开启灰度，0表示不开启灰度。
	IsGray *int64 `json:"IsGray,omitnil" name:"IsGray"`

	// 待废弃，可不填。灰度的地区
	GrayAreas []*string `json:"GrayAreas,omitnil" name:"GrayAreas"`

	// 是否开启HTTP强制跳转到HTTPS。
	// 0：不强制跳转
	// 1：开启强制跳转
	HttpsRewrite *int64 `json:"HttpsRewrite,omitnil" name:"HttpsRewrite"`

	// 域名回源时的回源域名。UpstreamType为1时，需要填充此字段
	UpstreamDomain *string `json:"UpstreamDomain,omitnil" name:"UpstreamDomain"`

	// IP回源时的回源IP列表。UpstreamType为0时，需要填充此字段
	SrcList []*string `json:"SrcList,omitnil" name:"SrcList"`

	// 是否开启HTTP2，需要开启HTTPS协议支持。
	// 0：关闭
	// 1：开启
	IsHttp2 *int64 `json:"IsHttp2,omitnil" name:"IsHttp2"`

	// 服务端口列表配置。
	// NginxServerId：新增域名时填'0'
	// Port：监听端口号
	// Protocol：端口协议
	// UpstreamPort：与Port相同
	// UpstreamProtocol：与Protocol相同
	Ports []*PortItem `json:"Ports,omitnil" name:"Ports"`

	// 待废弃，可不填。WAF实例类型。
	// sparta-waf：SAAS型WAF
	// clb-waf：负载均衡型WAF
	// cdn-waf：CDN上的Web防护能力
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 是否开启长连接。
	// 0： 短连接
	// 1： 长连接
	IsKeepAlive *string `json:"IsKeepAlive,omitnil" name:"IsKeepAlive"`

	// 域名所属实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 待废弃，目前填0即可。anycast IP类型开关： 0 普通IP 1 Anycast IP
	Anycast *int64 `json:"Anycast,omitnil" name:"Anycast"`

	// 回源IP列表各IP的权重，和SrcList一一对应。当且仅当UpstreamType为0，并且SrcList有多个IP，并且LoadBalance为2时需要填写，否则填 []
	Weights []*int64 `json:"Weights,omitnil" name:"Weights"`

	// 是否开启主动健康检测。
	// 0：不开启
	// 1：开启
	ActiveCheck *int64 `json:"ActiveCheck,omitnil" name:"ActiveCheck"`

	// TLS版本信息
	TLSVersion *int64 `json:"TLSVersion,omitnil" name:"TLSVersion"`

	// 加密套件模板。
	// 0：不支持选择，使用默认模版  
	// 1：通用型模版 
	// 2：安全型模版 
	// 3：自定义模版
	CipherTemplate *int64 `json:"CipherTemplate,omitnil" name:"CipherTemplate"`

	// 自定义的加密套件列表。CipherTemplate为3时需要填此字段，表示自定义的加密套件，值通过DescribeCiphersDetail接口获取。
	Ciphers []*int64 `json:"Ciphers,omitnil" name:"Ciphers"`

	// WAF与源站的读超时时间，默认300s。
	ProxyReadTimeout *int64 `json:"ProxyReadTimeout,omitnil" name:"ProxyReadTimeout"`

	// WAF与源站的写超时时间，默认300s。
	ProxySendTimeout *int64 `json:"ProxySendTimeout,omitnil" name:"ProxySendTimeout"`

	// WAF回源时的SNI类型。
	// 0：关闭SNI，不配置client_hello中的server_name
	// 1：开启SNI，client_hello中的server_name为防护域名
	// 2：开启SNI，SNI为域名回源时的源站域名
	// 3：开启SNI，SNI为自定义域名
	SniType *int64 `json:"SniType,omitnil" name:"SniType"`

	// SniType为3时，需要填此参数，表示自定义的SNI；
	SniHost *string `json:"SniHost,omitnil" name:"SniHost"`

	// 是否开启XFF重置。
	// 0：关闭
	// 1：开启
	XFFReset *int64 `json:"XFFReset,omitnil" name:"XFFReset"`
}

func (r *AddSpartaProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSpartaProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "CertType")
	delete(f, "IsCdn")
	delete(f, "UpstreamType")
	delete(f, "IsWebsocket")
	delete(f, "LoadBalance")
	delete(f, "Cert")
	delete(f, "PrivateKey")
	delete(f, "SSLId")
	delete(f, "ResourceId")
	delete(f, "IpHeaders")
	delete(f, "UpstreamScheme")
	delete(f, "HttpsUpstreamPort")
	delete(f, "IsGray")
	delete(f, "GrayAreas")
	delete(f, "HttpsRewrite")
	delete(f, "UpstreamDomain")
	delete(f, "SrcList")
	delete(f, "IsHttp2")
	delete(f, "Ports")
	delete(f, "Edition")
	delete(f, "IsKeepAlive")
	delete(f, "InstanceID")
	delete(f, "Anycast")
	delete(f, "Weights")
	delete(f, "ActiveCheck")
	delete(f, "TLSVersion")
	delete(f, "CipherTemplate")
	delete(f, "Ciphers")
	delete(f, "ProxyReadTimeout")
	delete(f, "ProxySendTimeout")
	delete(f, "SniType")
	delete(f, "SniHost")
	delete(f, "XFFReset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddSpartaProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddSpartaProtectionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddSpartaProtectionResponse struct {
	*tchttp.BaseResponse
	Response *AddSpartaProtectionResponseParams `json:"Response"`
}

func (r *AddSpartaProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddSpartaProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiPkg struct {
	// 资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceIds *string `json:"ResourceIds,omitnil" name:"ResourceIds"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *int64 `json:"Region,omitnil" name:"Region"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 申请数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InquireNum *int64 `json:"InquireNum,omitnil" name:"InquireNum"`

	// 使用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedNum *int64 `json:"UsedNum,omitnil" name:"UsedNum"`

	// 续费标志
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *uint64 `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// 计费项
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingItem *string `json:"BillingItem,omitnil" name:"BillingItem"`

	// api安全7天试用标识。1试用。0没试用
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAPISecurityTrial *int64 `json:"IsAPISecurityTrial,omitnil" name:"IsAPISecurityTrial"`
}

type AttackLogInfo struct {
	// 攻击日志的详情内容
	Content *string `json:"Content,omitnil" name:"Content"`

	// CLS返回内容
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// CLS返回内容
	Source *string `json:"Source,omitnil" name:"Source"`

	// CLS返回内容
	TimeStamp *string `json:"TimeStamp,omitnil" name:"TimeStamp"`
}

type AutoDenyDetail struct {
	// 攻击封禁类型标签
	AttackTags []*string `json:"AttackTags,omitnil" name:"AttackTags"`

	// 攻击次数阈值
	AttackThreshold *int64 `json:"AttackThreshold,omitnil" name:"AttackThreshold"`

	// 自动封禁状态
	DefenseStatus *int64 `json:"DefenseStatus,omitnil" name:"DefenseStatus"`

	// 攻击时间阈值
	TimeThreshold *int64 `json:"TimeThreshold,omitnil" name:"TimeThreshold"`

	// 自动封禁时间
	DenyTimeThreshold *int64 `json:"DenyTimeThreshold,omitnil" name:"DenyTimeThreshold"`

	// 最后更新时间
	LastUpdateTime *string `json:"LastUpdateTime,omitnil" name:"LastUpdateTime"`
}

type BatchIpAccessControlData struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 黑白名单条目
	Res []*BatchIpAccessControlItem `json:"Res,omitnil" name:"Res"`
}

type BatchIpAccessControlItem struct {
	// mongo表自增Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil" name:"Id"`

	// 黑名单42或白名单40
	ActionType *int64 `json:"ActionType,omitnil" name:"ActionType"`

	// 黑白名单的IP
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 备注
	Note *string `json:"Note,omitnil" name:"Note"`

	// 添加路径
	Source *string `json:"Source,omitnil" name:"Source"`

	// 修改时间
	TsVersion *uint64 `json:"TsVersion,omitnil" name:"TsVersion"`

	// 超时时间
	ValidTs *int64 `json:"ValidTs,omitnil" name:"ValidTs"`

	// 域名列表
	Hosts []*string `json:"Hosts,omitnil" name:"Hosts"`
}

type BotPkg struct {
	// 资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceIds *string `json:"ResourceIds,omitnil" name:"ResourceIds"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *int64 `json:"Region,omitnil" name:"Region"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 申请数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InquireNum *int64 `json:"InquireNum,omitnil" name:"InquireNum"`

	// 使用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedNum *int64 `json:"UsedNum,omitnil" name:"UsedNum"`

	// 子产品code
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 续费标志	
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *uint64 `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// 购买页bot6折
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotCPWaf *int64 `json:"BotCPWaf,omitnil" name:"BotCPWaf"`

	// 控制台买bot5折
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotNPWaf *int64 `json:"BotNPWaf,omitnil" name:"BotNPWaf"`

	// 7天bot试用标识 1 试用 0 没有试用
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsBotTrial *int64 `json:"IsBotTrial,omitnil" name:"IsBotTrial"`
}

type BotQPS struct {
	// 资源id
	ResourceIds *string `json:"ResourceIds,omitnil" name:"ResourceIds"`

	// 有效时间
	ValidTime *string `json:"ValidTime,omitnil" name:"ValidTime"`

	// 资源数量
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 资源所在地区
	Region *string `json:"Region,omitnil" name:"Region"`

	// 使用qps的最大值
	MaxBotQPS *uint64 `json:"MaxBotQPS,omitnil" name:"MaxBotQPS"`

	// 续费标志
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *uint64 `json:"RenewFlag,omitnil" name:"RenewFlag"`
}

type BotStatPointItem struct {
	// 横坐标
	TimeStamp *string `json:"TimeStamp,omitnil" name:"TimeStamp"`

	// value的所属对象
	Key *string `json:"Key,omitnil" name:"Key"`

	// 纵列表
	Value *int64 `json:"Value,omitnil" name:"Value"`

	// Key对应的页面展示内容
	Label *string `json:"Label,omitnil" name:"Label"`
}

type CCRuleData struct {
	// cc规则
	Res []*CCRuleItem `json:"Res,omitnil" name:"Res"`

	// 规则数目
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`
}

type CCRuleItem struct {
	// 动作
	ActionType *uint64 `json:"ActionType,omitnil" name:"ActionType"`

	// 高级模式
	Advance *uint64 `json:"Advance,omitnil" name:"Advance"`

	// 时间周期
	Interval *uint64 `json:"Interval,omitnil" name:"Interval"`

	// 限制次数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 匹配方法
	MatchFunc *uint64 `json:"MatchFunc,omitnil" name:"MatchFunc"`

	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 优先级
	Priority *uint64 `json:"Priority,omitnil" name:"Priority"`

	// 状态
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 更新时间戳
	TsVersion *uint64 `json:"TsVersion,omitnil" name:"TsVersion"`

	// 匹配url
	Url *string `json:"Url,omitnil" name:"Url"`

	// 策略动作有效时间
	ValidTime *uint64 `json:"ValidTime,omitnil" name:"ValidTime"`

	// 高级参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OptionsArr *string `json:"OptionsArr,omitnil" name:"OptionsArr"`
}

type CCRuleItems struct {
	// 名字
	Name *string `json:"Name,omitnil" name:"Name"`

	// 状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 模式
	Advance *uint64 `json:"Advance,omitnil" name:"Advance"`

	// 限制
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 范围
	Interval *uint64 `json:"Interval,omitnil" name:"Interval"`

	// 网址
	Url *string `json:"Url,omitnil" name:"Url"`

	// 匹配类型
	MatchFunc *uint64 `json:"MatchFunc,omitnil" name:"MatchFunc"`

	// 动作
	ActionType *uint64 `json:"ActionType,omitnil" name:"ActionType"`

	// 优先级
	Priority *uint64 `json:"Priority,omitnil" name:"Priority"`

	// 有效时间
	ValidTime *uint64 `json:"ValidTime,omitnil" name:"ValidTime"`

	// 版本
	TsVersion *uint64 `json:"TsVersion,omitnil" name:"TsVersion"`

	// 规则详情
	Options *string `json:"Options,omitnil" name:"Options"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 事件id
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventId *string `json:"EventId,omitnil" name:"EventId"`

	// 关联的Session规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionApplied []*int64 `json:"SessionApplied,omitnil" name:"SessionApplied"`
}

type CCRuleLists struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Res []*CCRuleItems `json:"Res,omitnil" name:"Res"`
}

type CacheUrlItem struct {
	// Id
	Id *string `json:"Id,omitnil" name:"Id"`

	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// uri
	Uri *string `json:"Uri,omitnil" name:"Uri"`

	// 协议
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 状态
	Status *string `json:"Status,omitnil" name:"Status"`
}

type CacheUrlItems struct {
	// 标识
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 名字
	Name *string `json:"Name,omitnil" name:"Name"`

	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 网址
	Uri *string `json:"Uri,omitnil" name:"Uri"`

	// 协议
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

type CdcCluster struct {
	// cdc的集群id
	Id *string `json:"Id,omitnil" name:"Id"`

	// cdc的集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`
}

type CdcRegion struct {
	// 地域
	Region *string `json:"Region,omitnil" name:"Region"`

	// 该地域对应的集群信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Clusters []*CdcCluster `json:"Clusters,omitnil" name:"Clusters"`
}

type ClbDomainsInfo struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名id
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// waf类型
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 是否是cdn
	IsCdn *uint64 `json:"IsCdn,omitnil" name:"IsCdn"`

	// 负载均衡算法
	LoadBalancerSet []*LoadBalancerPackageNew `json:"LoadBalancerSet,omitnil" name:"LoadBalancerSet"`

	// 镜像模式
	FlowMode *uint64 `json:"FlowMode,omitnil" name:"FlowMode"`

	// 绑定clb状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *int64 `json:"State,omitnil" name:"State"`

	// 负载均衡类型，clb或者apisix
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlbType *string `json:"AlbType,omitnil" name:"AlbType"`

	// IsCdn=3时，表示自定义header
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpHeaders []*string `json:"IpHeaders,omitnil" name:"IpHeaders"`

	// cdc类型会增加集群信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdcClusters *string `json:"CdcClusters,omitnil" name:"CdcClusters"`

	// 云类型:public:公有云；private:私有云;hybrid:混合云
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloudType *string `json:"CloudType,omitnil" name:"CloudType"`
}

type ClbHostResult struct {
	// WAF绑定的监听器实例
	LoadBalancer *LoadBalancer `json:"LoadBalancer,omitnil" name:"LoadBalancer"`

	// WAF绑定的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// WAF绑定的实例ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 是否有绑定WAF，1：绑定了WAF，0：没有绑定WAF
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 绑定了WAF的情况下，WAF流量模式，1：清洗模式，0：镜像模式（默认）
	FlowMode *uint64 `json:"FlowMode,omitnil" name:"FlowMode"`
}

type ClbHostsParams struct {
	// 负载均衡实例ID，如果不传次参数则默认认为操作的是整个AppId的监听器，如果此参数不为空则认为操作的是对应负载均衡实例。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil" name:"LoadBalancerId"`

	// 负载均衡监听器ID，，如果不传次参数则默认认为操作的是整个负载均衡实例，如果此参数不为空则认为操作的是对应负载均衡监听器。
	ListenerId *string `json:"ListenerId,omitnil" name:"ListenerId"`

	// WAF实例ID，，如果不传次参数则默认认为操作的是整个负载均衡监听器实例，如果此参数不为空则认为操作的是对应负载均衡监听器的某一个具体的域名。
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`
}

type ClbObject struct {
	// 对象ID
	ObjectId *string `json:"ObjectId,omitnil" name:"ObjectId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 精准域名列表
	PreciseDomains []*string `json:"PreciseDomains,omitnil" name:"PreciseDomains"`

	// WAF功能开关状态，0关闭1开启
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// WAF日志开关状态，0关闭1开启
	ClsStatus *int64 `json:"ClsStatus,omitnil" name:"ClsStatus"`

	// CLB对象对应的虚拟域名
	VirtualDomain *string `json:"VirtualDomain,omitnil" name:"VirtualDomain"`

	// 对象名称
	ObjectName *string `json:"ObjectName,omitnil" name:"ObjectName"`

	// 公网地址
	PublicIp []*string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 内网地址
	PrivateIp []*string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// VPC名称
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// VPC ID
	Vpc *string `json:"Vpc,omitnil" name:"Vpc"`

	// waf实例等级，如果未绑定实例为0
	InstanceLevel *int64 `json:"InstanceLevel,omitnil" name:"InstanceLevel"`

	// clb投递开关
	PostCLSStatus *int64 `json:"PostCLSStatus,omitnil" name:"PostCLSStatus"`

	// kafka投递开关
	PostCKafkaStatus *int64 `json:"PostCKafkaStatus,omitnil" name:"PostCKafkaStatus"`

	// 对象类型：CLB:负载均衡器，TSE:云原生网关
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 对象地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`
}

type ClbWafRegionItem struct {
	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil" name:"Id"`

	// 地域中文说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 地域英文全拼
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 地域编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitnil" name:"Code"`
}

// Predefined struct for user
type CreateAccessExportRequestParams struct {
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitnil" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitnil" name:"To"`

	// 日志导出检索语句
	Query *string `json:"Query,omitnil" name:"Query"`

	// 日志导出数量，最大值100w
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 日志导出数据格式。json，csv，默认为json
	Format *string `json:"Format,omitnil" name:"Format"`

	// 日志导出时间排序。desc，asc，默认为desc
	Order *string `json:"Order,omitnil" name:"Order"`
}

type CreateAccessExportRequest struct {
	*tchttp.BaseRequest
	
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitnil" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitnil" name:"To"`

	// 日志导出检索语句
	Query *string `json:"Query,omitnil" name:"Query"`

	// 日志导出数量，最大值100w
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 日志导出数据格式。json，csv，默认为json
	Format *string `json:"Format,omitnil" name:"Format"`

	// 日志导出时间排序。desc，asc，默认为desc
	Order *string `json:"Order,omitnil" name:"Order"`
}

func (r *CreateAccessExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "Count")
	delete(f, "Format")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccessExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccessExportResponseParams struct {
	// 日志导出ID。
	ExportId *string `json:"ExportId,omitnil" name:"ExportId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAccessExportResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccessExportResponseParams `json:"Response"`
}

func (r *CreateAccessExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccessExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHostRequestParams struct {
	// 防护域名配置信息
	Host *HostRecord `json:"Host,omitnil" name:"Host"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

type CreateHostRequest struct {
	*tchttp.BaseRequest
	
	// 防护域名配置信息
	Host *HostRecord `json:"Host,omitnil" name:"Host"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

func (r *CreateHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Host")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHostResponseParams struct {
	// 新增防护域名ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateHostResponse struct {
	*tchttp.BaseResponse
	Response *CreateHostResponseParams `json:"Response"`
}

func (r *CreateHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DealData struct {
	// 订单号列表，元素个数与请求包的goods数组的元素个数一致，商品详情与订单按顺序对应
	DealNames []*string `json:"DealNames,omitnil" name:"DealNames"`

	// 大订单号，一个大订单号下可以有多个子订单，说明是同一次下单[{},{}]
	BigDealId *string `json:"BigDealId,omitnil" name:"BigDealId"`
}

// Predefined struct for user
type DeleteAccessExportRequestParams struct {
	// 日志导出ID
	ExportId *string `json:"ExportId,omitnil" name:"ExportId"`

	// 日志主题
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`
}

type DeleteAccessExportRequest struct {
	*tchttp.BaseRequest
	
	// 日志导出ID
	ExportId *string `json:"ExportId,omitnil" name:"ExportId"`

	// 日志主题
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`
}

func (r *DeleteAccessExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccessExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportId")
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccessExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccessExportResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAccessExportResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccessExportResponseParams `json:"Response"`
}

func (r *DeleteAccessExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccessExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAntiFakeUrlRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Id
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

type DeleteAntiFakeUrlRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Id
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

func (r *DeleteAntiFakeUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAntiFakeUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAntiFakeUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAntiFakeUrlResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAntiFakeUrlResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAntiFakeUrlResponseParams `json:"Response"`
}

func (r *DeleteAntiFakeUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAntiFakeUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAntiInfoLeakRuleRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 规则id
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`
}

type DeleteAntiInfoLeakRuleRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 规则id
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`
}

func (r *DeleteAntiInfoLeakRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAntiInfoLeakRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAntiInfoLeakRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAntiInfoLeakRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAntiInfoLeakRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAntiInfoLeakRuleResponseParams `json:"Response"`
}

func (r *DeleteAntiInfoLeakRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAntiInfoLeakRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAttackDownloadRecordRequestParams struct {
	// 下载任务记录唯一标记
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

type DeleteAttackDownloadRecordRequest struct {
	*tchttp.BaseRequest
	
	// 下载任务记录唯一标记
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

func (r *DeleteAttackDownloadRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttackDownloadRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAttackDownloadRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAttackDownloadRecordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAttackDownloadRecordResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAttackDownloadRecordResponseParams `json:"Response"`
}

func (r *DeleteAttackDownloadRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttackDownloadRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAttackWhiteRuleRequestParams struct {
	// 规则序号组
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`

	// 用户域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type DeleteAttackWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则序号组
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`

	// 用户域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *DeleteAttackWhiteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttackWhiteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAttackWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAttackWhiteRuleResponseParams struct {
	// 删除失败的规则序号组
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailIds []*uint64 `json:"FailIds,omitnil" name:"FailIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAttackWhiteRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAttackWhiteRuleResponseParams `json:"Response"`
}

func (r *DeleteAttackWhiteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttackWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCCRuleRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 规则名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// clb-waf或者sparta-waf
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 规则Id
	RuleId *int64 `json:"RuleId,omitnil" name:"RuleId"`
}

type DeleteCCRuleRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 规则名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// clb-waf或者sparta-waf
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 规则Id
	RuleId *int64 `json:"RuleId,omitnil" name:"RuleId"`
}

func (r *DeleteCCRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Name")
	delete(f, "Edition")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCCRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCCRuleResponseParams struct {
	// 一般为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil" name:"Data"`

	// 操作的规则Id
	RuleId *int64 `json:"RuleId,omitnil" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCCRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCCRuleResponseParams `json:"Response"`
}

func (r *DeleteCCRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCCRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomRuleRequestParams struct {
	// 删除的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 删除的规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil" name:"Edition"`
}

type DeleteCustomRuleRequest struct {
	*tchttp.BaseRequest
	
	// 删除的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 删除的规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil" name:"Edition"`
}

func (r *DeleteCustomRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RuleId")
	delete(f, "Edition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCustomRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCustomRuleResponseParams `json:"Response"`
}

func (r *DeleteCustomRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomWhiteRuleRequestParams struct {
	// 删除的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 删除的规则ID
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`
}

type DeleteCustomWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 删除的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 删除的规则ID
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`
}

func (r *DeleteCustomWhiteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomWhiteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomWhiteRuleResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	Success *ResponseCode `json:"Success,omitnil" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCustomWhiteRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCustomWhiteRuleResponseParams `json:"Response"`
}

func (r *DeleteCustomWhiteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainWhiteRulesRequestParams struct {
	// 需要删除的规则域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 需要删除的白名单规则
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type DeleteDomainWhiteRulesRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的规则域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 需要删除的白名单规则
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *DeleteDomainWhiteRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainWhiteRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDomainWhiteRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainWhiteRulesResponseParams struct {
	// 出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteDomainWhiteRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDomainWhiteRulesResponseParams `json:"Response"`
}

func (r *DeleteDomainWhiteRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainWhiteRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDownloadRecordRequestParams struct {
	// 记录id
	Flow *string `json:"Flow,omitnil" name:"Flow"`
}

type DeleteDownloadRecordRequest struct {
	*tchttp.BaseRequest
	
	// 记录id
	Flow *string `json:"Flow,omitnil" name:"Flow"`
}

func (r *DeleteDownloadRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDownloadRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Flow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDownloadRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDownloadRecordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteDownloadRecordResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDownloadRecordResponseParams `json:"Response"`
}

func (r *DeleteDownloadRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDownloadRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHostRequestParams struct {
	// 删除的域名列表
	HostsDel []*HostDel `json:"HostsDel,omitnil" name:"HostsDel"`
}

type DeleteHostRequest struct {
	*tchttp.BaseRequest
	
	// 删除的域名列表
	HostsDel []*HostDel `json:"HostsDel,omitnil" name:"HostsDel"`
}

func (r *DeleteHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HostsDel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHostResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	Success *ResponseCode `json:"Success,omitnil" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteHostResponse struct {
	*tchttp.BaseResponse
	Response *DeleteHostResponseParams `json:"Response"`
}

func (r *DeleteHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIpAccessControlRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 删除的ip数组
	Items []*string `json:"Items,omitnil" name:"Items"`

	// 若IsId字段为True，则Items列表元素需为Id，否则为IP
	IsId *bool `json:"IsId,omitnil" name:"IsId"`

	// 是否删除对应的域名下的所有黑/白IP名单，true表示全部删除，false表示只删除指定ip名单
	DeleteAll *bool `json:"DeleteAll,omitnil" name:"DeleteAll"`

	// 是否为多域名黑白名单
	SourceType *string `json:"SourceType,omitnil" name:"SourceType"`
}

type DeleteIpAccessControlRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 删除的ip数组
	Items []*string `json:"Items,omitnil" name:"Items"`

	// 若IsId字段为True，则Items列表元素需为Id，否则为IP
	IsId *bool `json:"IsId,omitnil" name:"IsId"`

	// 是否删除对应的域名下的所有黑/白IP名单，true表示全部删除，false表示只删除指定ip名单
	DeleteAll *bool `json:"DeleteAll,omitnil" name:"DeleteAll"`

	// 是否为多域名黑白名单
	SourceType *string `json:"SourceType,omitnil" name:"SourceType"`
}

func (r *DeleteIpAccessControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIpAccessControlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Items")
	delete(f, "IsId")
	delete(f, "DeleteAll")
	delete(f, "SourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIpAccessControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIpAccessControlResponseParams struct {
	// 删除失败的条目
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedItems *string `json:"FailedItems,omitnil" name:"FailedItems"`

	// 删除失败的条目数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedCount *int64 `json:"FailedCount,omitnil" name:"FailedCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteIpAccessControlResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIpAccessControlResponseParams `json:"Response"`
}

func (r *DeleteIpAccessControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIpAccessControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSessionRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// clb-waf 或者 sprta-waf
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 要删除的SessionID
	SessionID *int64 `json:"SessionID,omitnil" name:"SessionID"`
}

type DeleteSessionRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// clb-waf 或者 sprta-waf
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 要删除的SessionID
	SessionID *int64 `json:"SessionID,omitnil" name:"SessionID"`
}

func (r *DeleteSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Edition")
	delete(f, "SessionID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSessionResponseParams struct {
	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteSessionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSessionResponseParams `json:"Response"`
}

func (r *DeleteSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSpartaProtectionRequestParams struct {
	// 域名列表
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 实例类型
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

type DeleteSpartaProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 域名列表
	Domains []*string `json:"Domains,omitnil" name:"Domains"`

	// 实例类型
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

func (r *DeleteSpartaProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSpartaProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "Edition")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSpartaProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSpartaProtectionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteSpartaProtectionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSpartaProtectionResponseParams `json:"Response"`
}

func (r *DeleteSpartaProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSpartaProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessExportsRequestParams struct {
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeAccessExportsRequest struct {
	*tchttp.BaseRequest
	
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeAccessExportsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessExportsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessExportsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessExportsResponseParams struct {
	// 日志导出ID。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 日志导出列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Exports []*ExportAccessInfo `json:"Exports,omitnil" name:"Exports"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAccessExportsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessExportsResponseParams `json:"Response"`
}

func (r *DescribeAccessExportsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessExportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessFastAnalysisRequestParams struct {
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitnil" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitnil" name:"To"`

	// 查询语句，语句长度最大为4096，由于本接口是分析接口，如果无过滤条件，必须传 * 表示匹配所有，参考CLS的分析统计语句的文档
	Query *string `json:"Query,omitnil" name:"Query"`

	// 需要分析统计的字段名
	FieldName *string `json:"FieldName,omitnil" name:"FieldName"`

	// 排序字段,升序asc,降序desc，默认降序desc 
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 返回的top数，默认返回top5
	Count *int64 `json:"Count,omitnil" name:"Count"`
}

type DescribeAccessFastAnalysisRequest struct {
	*tchttp.BaseRequest
	
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitnil" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitnil" name:"To"`

	// 查询语句，语句长度最大为4096，由于本接口是分析接口，如果无过滤条件，必须传 * 表示匹配所有，参考CLS的分析统计语句的文档
	Query *string `json:"Query,omitnil" name:"Query"`

	// 需要分析统计的字段名
	FieldName *string `json:"FieldName,omitnil" name:"FieldName"`

	// 排序字段,升序asc,降序desc，默认降序desc 
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 返回的top数，默认返回top5
	Count *int64 `json:"Count,omitnil" name:"Count"`
}

func (r *DescribeAccessFastAnalysisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessFastAnalysisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "FieldName")
	delete(f, "Sort")
	delete(f, "Count")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessFastAnalysisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessFastAnalysisResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAccessFastAnalysisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessFastAnalysisResponseParams `json:"Response"`
}

func (r *DescribeAccessFastAnalysisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessFastAnalysisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessHistogramRequestParams struct {
	// 老版本查询的日志主题ID，新版本传空字符串即可
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitnil" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitnil" name:"To"`

	// 查询语句，语句长度最大为4096
	Query *string `json:"Query,omitnil" name:"Query"`

	// 柱状图间隔时间差，单位ms
	Interval *int64 `json:"Interval,omitnil" name:"Interval"`
}

type DescribeAccessHistogramRequest struct {
	*tchttp.BaseRequest
	
	// 老版本查询的日志主题ID，新版本传空字符串即可
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitnil" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitnil" name:"To"`

	// 查询语句，语句长度最大为4096
	Query *string `json:"Query,omitnil" name:"Query"`

	// 柱状图间隔时间差，单位ms
	Interval *int64 `json:"Interval,omitnil" name:"Interval"`
}

func (r *DescribeAccessHistogramRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessHistogramRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "Interval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessHistogramRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessHistogramResponseParams struct {
	// 柱状图间隔时间差，单位ms
	Interval *int64 `json:"Interval,omitnil" name:"Interval"`

	// 满足条件的日志条数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 注意：此字段可能返回 null，表示取不到有效值
	// 注意：此字段可能返回 null，表示取不到有效值。
	HistogramInfos []*AccessHistogramItem `json:"HistogramInfos,omitnil" name:"HistogramInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAccessHistogramResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessHistogramResponseParams `json:"Response"`
}

func (r *DescribeAccessHistogramResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessHistogramResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessIndexRequestParams struct {

}

type DescribeAccessIndexRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAccessIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessIndexResponseParams struct {
	// 是否生效，true表示生效，false表示未生效
	Status *bool `json:"Status,omitnil" name:"Status"`

	// 索引配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rule *AccessRuleInfo `json:"Rule,omitnil" name:"Rule"`

	// 索引修改时间，初始值为索引创建时间。
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAccessIndexResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccessIndexResponseParams `json:"Response"`
}

func (r *DescribeAccessIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccessIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAntiFakeRulesRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 偏移
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤数组,name可以是如下的值： RuleID,ParamName,Url,Action,Method,Source,Status
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`

	// asc或者desc
	Order *string `json:"Order,omitnil" name:"Order"`

	// 目前支持根据ts排序
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeAntiFakeRulesRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 偏移
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤数组,name可以是如下的值： RuleID,ParamName,Url,Action,Method,Source,Status
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`

	// asc或者desc
	Order *string `json:"Order,omitnil" name:"Order"`

	// 目前支持根据ts排序
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeAntiFakeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAntiFakeRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAntiFakeRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAntiFakeRulesResponseParams struct {
	// 返回值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*CacheUrlItems `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAntiFakeRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAntiFakeRulesResponseParams `json:"Response"`
}

func (r *DescribeAntiFakeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAntiFakeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAntiFakeUrlRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 翻页参数
	PageInfo *PageInfo `json:"PageInfo,omitnil" name:"PageInfo"`
}

type DescribeAntiFakeUrlRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 翻页参数
	PageInfo *PageInfo `json:"PageInfo,omitnil" name:"PageInfo"`
}

func (r *DescribeAntiFakeUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAntiFakeUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "PageInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAntiFakeUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAntiFakeUrlResponseParams struct {
	// 总数
	Total *string `json:"Total,omitnil" name:"Total"`

	// 信息
	List []*CacheUrlItem `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAntiFakeUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAntiFakeUrlResponseParams `json:"Response"`
}

func (r *DescribeAntiFakeUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAntiFakeUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAntiInfoLeakRulesRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 动作类型
	ActionType *int64 `json:"ActionType,omitnil" name:"ActionType"`

	// 翻页
	PageInfo *PageInfo `json:"PageInfo,omitnil" name:"PageInfo"`
}

type DescribeAntiInfoLeakRulesRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 动作类型
	ActionType *int64 `json:"ActionType,omitnil" name:"ActionType"`

	// 翻页
	PageInfo *PageInfo `json:"PageInfo,omitnil" name:"PageInfo"`
}

func (r *DescribeAntiInfoLeakRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAntiInfoLeakRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "ActionType")
	delete(f, "PageInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAntiInfoLeakRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAntiInfoLeakRulesResponseParams struct {
	// 记录条数
	TotalCount *string `json:"TotalCount,omitnil" name:"TotalCount"`

	// 规则列表
	RuleList []*DescribeAntiInfoLeakRulesRuleItem `json:"RuleList,omitnil" name:"RuleList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAntiInfoLeakRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAntiInfoLeakRulesResponseParams `json:"Response"`
}

func (r *DescribeAntiInfoLeakRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAntiInfoLeakRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAntiInfoLeakRulesRuleItem struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 规则名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 规则状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// 规则动作类型
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// 规则创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 详细的规则
	Strategies []*DescribeAntiInfoLeakRulesStrategyItem `json:"Strategies,omitnil" name:"Strategies"`
}

type DescribeAntiInfoLeakRulesStrategyItem struct {
	// 字段
	Field *string `json:"Field,omitnil" name:"Field"`

	// 条件
	CompareFunc *string `json:"CompareFunc,omitnil" name:"CompareFunc"`

	// 内容
	Content *string `json:"Content,omitnil" name:"Content"`
}

// Predefined struct for user
type DescribeAntiInfoLeakageRulesRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 翻页支持，读取偏移
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 翻页支持，读取长度限制
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc或者desc
	Order *string `json:"Order,omitnil" name:"Order"`

	// 过滤器,可以允许如下的值：
	// RuleId,Match_field,Name,Action,Status
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`
}

type DescribeAntiInfoLeakageRulesRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 翻页支持，读取偏移
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 翻页支持，读取长度限制
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序方式，asc或者desc
	Order *string `json:"Order,omitnil" name:"Order"`

	// 过滤器,可以允许如下的值：
	// RuleId,Match_field,Name,Action,Status
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeAntiInfoLeakageRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAntiInfoLeakageRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAntiInfoLeakageRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAntiInfoLeakageRulesResponseParams struct {
	// 记录条数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 规则列表
	RuleList []*DescribeAntiLeakageItem `json:"RuleList,omitnil" name:"RuleList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAntiInfoLeakageRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAntiInfoLeakageRulesResponseParams `json:"Response"`
}

func (r *DescribeAntiInfoLeakageRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAntiInfoLeakageRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAntiLeakageItem struct {
	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 状态值
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 动作
	Action *string `json:"Action,omitnil" name:"Action"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 匹配条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Strategies []*DescribeAntiInfoLeakRulesStrategyItem `json:"Strategies,omitnil" name:"Strategies"`

	// 匹配的URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uri *string `json:"Uri,omitnil" name:"Uri"`
}

// Predefined struct for user
type DescribeAttackOverviewRequestParams struct {
	// 查询开始时间
	FromTime *string `json:"FromTime,omitnil" name:"FromTime"`

	// 查询结束时间
	ToTime *string `json:"ToTime,omitnil" name:"ToTime"`

	// 客户的Appid
	Appid *uint64 `json:"Appid,omitnil" name:"Appid"`

	// 被查询的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

type DescribeAttackOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始时间
	FromTime *string `json:"FromTime,omitnil" name:"FromTime"`

	// 查询结束时间
	ToTime *string `json:"ToTime,omitnil" name:"ToTime"`

	// 客户的Appid
	Appid *uint64 `json:"Appid,omitnil" name:"Appid"`

	// 被查询的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

func (r *DescribeAttackOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromTime")
	delete(f, "ToTime")
	delete(f, "Appid")
	delete(f, "Domain")
	delete(f, "Edition")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAttackOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackOverviewResponseParams struct {
	// 访问请求总数
	AccessCount *uint64 `json:"AccessCount,omitnil" name:"AccessCount"`

	// Web攻击总数
	AttackCount *uint64 `json:"AttackCount,omitnil" name:"AttackCount"`

	// 访问控制总数
	ACLCount *uint64 `json:"ACLCount,omitnil" name:"ACLCount"`

	// CC攻击总数
	CCCount *uint64 `json:"CCCount,omitnil" name:"CCCount"`

	// Bot攻击总数
	BotCount *uint64 `json:"BotCount,omitnil" name:"BotCount"`

	// api资产总数
	ApiAssetsCount *uint64 `json:"ApiAssetsCount,omitnil" name:"ApiAssetsCount"`

	// api风险事件数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiRiskEventCount *uint64 `json:"ApiRiskEventCount,omitnil" name:"ApiRiskEventCount"`

	// 黑名单总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPBlackCount *uint64 `json:"IPBlackCount,omitnil" name:"IPBlackCount"`

	// 防篡改总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TamperCount *uint64 `json:"TamperCount,omitnil" name:"TamperCount"`

	// 信息泄露总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LeakCount *uint64 `json:"LeakCount,omitnil" name:"LeakCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAttackOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAttackOverviewResponseParams `json:"Response"`
}

func (r *DescribeAttackOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackTypeRequestParams struct {
	// 起始时间
	FromTime *string `json:"FromTime,omitnil" name:"FromTime"`

	// 结束时间
	ToTime *string `json:"ToTime,omitnil" name:"ToTime"`

	// 兼容Host，逐步淘汰Host字段
	Host *string `json:"Host,omitnil" name:"Host"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 域名过滤，不传则不过滤，用于替代Host字段，逐步淘汰Host
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type DescribeAttackTypeRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间
	FromTime *string `json:"FromTime,omitnil" name:"FromTime"`

	// 结束时间
	ToTime *string `json:"ToTime,omitnil" name:"ToTime"`

	// 兼容Host，逐步淘汰Host字段
	Host *string `json:"Host,omitnil" name:"Host"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 域名过滤，不传则不过滤，用于替代Host字段，逐步淘汰Host
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *DescribeAttackTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromTime")
	delete(f, "ToTime")
	delete(f, "Host")
	delete(f, "Edition")
	delete(f, "InstanceID")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAttackTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackTypeResponseParams struct {
	// 数量
	Piechart []*PiechartItem `json:"Piechart,omitnil" name:"Piechart"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAttackTypeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAttackTypeResponseParams `json:"Response"`
}

func (r *DescribeAttackTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackWhiteRuleRequestParams struct {
	// 需要查询的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 分页
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页容量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段，支持user_id, signature_id, modify_time
	By *string `json:"By,omitnil" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitnil" name:"Order"`

	// 筛选条件，支持SignatureId, MatchContent
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`
}

type DescribeAttackWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 分页
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页容量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段，支持user_id, signature_id, modify_time
	By *string `json:"By,omitnil" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitnil" name:"Order"`

	// 筛选条件，支持SignatureId, MatchContent
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeAttackWhiteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackWhiteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "By")
	delete(f, "Order")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAttackWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackWhiteRuleResponseParams struct {
	// 规则总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 规则白名单列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*UserWhiteRule `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAttackWhiteRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAttackWhiteRuleResponseParams `json:"Response"`
}

func (r *DescribeAttackWhiteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttackWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoDenyIPRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 查询IP自动封禁状态
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 计数标识
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 类别
	Category *string `json:"Category,omitnil" name:"Category"`

	// 有效时间最小时间戳
	VtsMin *uint64 `json:"VtsMin,omitnil" name:"VtsMin"`

	// 有效时间最大时间戳
	VtsMax *uint64 `json:"VtsMax,omitnil" name:"VtsMax"`

	// 创建时间最小时间戳
	CtsMin *uint64 `json:"CtsMin,omitnil" name:"CtsMin"`

	// 创建时间最大时间戳
	CtsMax *uint64 `json:"CtsMax,omitnil" name:"CtsMax"`

	// 偏移量
	Skip *uint64 `json:"Skip,omitnil" name:"Skip"`

	// 限制条数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 策略名字
	Name *string `json:"Name,omitnil" name:"Name"`

	// 排序参数
	Sort *string `json:"Sort,omitnil" name:"Sort"`
}

type DescribeAutoDenyIPRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 查询IP自动封禁状态
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 计数标识
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 类别
	Category *string `json:"Category,omitnil" name:"Category"`

	// 有效时间最小时间戳
	VtsMin *uint64 `json:"VtsMin,omitnil" name:"VtsMin"`

	// 有效时间最大时间戳
	VtsMax *uint64 `json:"VtsMax,omitnil" name:"VtsMax"`

	// 创建时间最小时间戳
	CtsMin *uint64 `json:"CtsMin,omitnil" name:"CtsMin"`

	// 创建时间最大时间戳
	CtsMax *uint64 `json:"CtsMax,omitnil" name:"CtsMax"`

	// 偏移量
	Skip *uint64 `json:"Skip,omitnil" name:"Skip"`

	// 限制条数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 策略名字
	Name *string `json:"Name,omitnil" name:"Name"`

	// 排序参数
	Sort *string `json:"Sort,omitnil" name:"Sort"`
}

func (r *DescribeAutoDenyIPRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoDenyIPRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Ip")
	delete(f, "Count")
	delete(f, "Category")
	delete(f, "VtsMin")
	delete(f, "VtsMax")
	delete(f, "CtsMin")
	delete(f, "CtsMax")
	delete(f, "Skip")
	delete(f, "Limit")
	delete(f, "Name")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoDenyIPRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoDenyIPResponseParams struct {
	// 查询IP封禁状态返回结果
	Data *IpHitItemsData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAutoDenyIPResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoDenyIPResponseParams `json:"Response"`
}

func (r *DescribeAutoDenyIPResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoDenyIPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchIpAccessControlRequestParams struct {
	// 筛选条件，支持 ActionType 40/42，Ip：ip地址，Domain：域名三类
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`

	// 偏移
	OffSet *uint64 `json:"OffSet,omitnil" name:"OffSet"`

	// 限制
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序参数
	Sort *string `json:"Sort,omitnil" name:"Sort"`
}

type DescribeBatchIpAccessControlRequest struct {
	*tchttp.BaseRequest
	
	// 筛选条件，支持 ActionType 40/42，Ip：ip地址，Domain：域名三类
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`

	// 偏移
	OffSet *uint64 `json:"OffSet,omitnil" name:"OffSet"`

	// 限制
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序参数
	Sort *string `json:"Sort,omitnil" name:"Sort"`
}

func (r *DescribeBatchIpAccessControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchIpAccessControlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "OffSet")
	delete(f, "Limit")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBatchIpAccessControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchIpAccessControlResponseParams struct {
	// 输出
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *BatchIpAccessControlData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBatchIpAccessControlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBatchIpAccessControlResponseParams `json:"Response"`
}

func (r *DescribeBatchIpAccessControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchIpAccessControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCAutoStatusRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type DescribeCCAutoStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *DescribeCCAutoStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCAutoStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCAutoStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCAutoStatusResponseParams struct {
	// 配置状态
	AutoCCSwitch *int64 `json:"AutoCCSwitch,omitnil" name:"AutoCCSwitch"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCCAutoStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCAutoStatusResponseParams `json:"Response"`
}

func (r *DescribeCCAutoStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCAutoStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCRuleListRequestParams struct {
	// 需要查询的API所属的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 偏移
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 目前支持根据ts_version排序
	By *string `json:"By,omitnil" name:"By"`

	// 过滤数组,name可以是如下的值： RuleID,ParamName,Url,Action,Method,Source,Status
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`

	// asc或者desc
	Order *string `json:"Order,omitnil" name:"Order"`
}

type DescribeCCRuleListRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的API所属的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 偏移
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 目前支持根据ts_version排序
	By *string `json:"By,omitnil" name:"By"`

	// 过滤数组,name可以是如下的值： RuleID,ParamName,Url,Action,Method,Source,Status
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`

	// asc或者desc
	Order *string `json:"Order,omitnil" name:"Order"`
}

func (r *DescribeCCRuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCRuleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "By")
	delete(f, "Filters")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCRuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCRuleListResponseParams struct {
	// 查询到的CC规则的列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CCRuleLists `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCCRuleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCRuleListResponseParams `json:"Response"`
}

func (r *DescribeCCRuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCRuleRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 页码
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 页的数目
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序参数
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// clb-waf 或者 sparta-waf
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 过滤条件
	Name *string `json:"Name,omitnil" name:"Name"`
}

type DescribeCCRuleRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 页码
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 页的数目
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序参数
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// clb-waf 或者 sparta-waf
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 过滤条件
	Name *string `json:"Name,omitnil" name:"Name"`
}

func (r *DescribeCCRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Sort")
	delete(f, "Edition")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCCRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCCRuleResponseParams struct {
	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CCRuleData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCCRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCCRuleResponseParams `json:"Response"`
}

func (r *DescribeCCRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCCRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateVerifyResultRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 证书类型
	CertType *int64 `json:"CertType,omitnil" name:"CertType"`

	// 证书公钥
	Certificate *string `json:"Certificate,omitnil" name:"Certificate"`

	// 证书ID
	CertID *string `json:"CertID,omitnil" name:"CertID"`

	// 私钥信息
	PrivateKey *string `json:"PrivateKey,omitnil" name:"PrivateKey"`
}

type DescribeCertificateVerifyResultRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 证书类型
	CertType *int64 `json:"CertType,omitnil" name:"CertType"`

	// 证书公钥
	Certificate *string `json:"Certificate,omitnil" name:"Certificate"`

	// 证书ID
	CertID *string `json:"CertID,omitnil" name:"CertID"`

	// 私钥信息
	PrivateKey *string `json:"PrivateKey,omitnil" name:"PrivateKey"`
}

func (r *DescribeCertificateVerifyResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateVerifyResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "CertType")
	delete(f, "Certificate")
	delete(f, "CertID")
	delete(f, "PrivateKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificateVerifyResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateVerifyResultResponseParams struct {
	// 状态码
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 错误详情
	Detail []*string `json:"Detail,omitnil" name:"Detail"`

	// 过期时间
	NotAfter *string `json:"NotAfter,omitnil" name:"NotAfter"`

	// 证书是否改变:1有改变，0没有改变
	// 注意：此字段可能返回 null，表示取不到有效值。
	Changed *int64 `json:"Changed,omitnil" name:"Changed"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCertificateVerifyResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCertificateVerifyResultResponseParams `json:"Response"`
}

func (r *DescribeCertificateVerifyResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCertificateVerifyResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCiphersDetailRequestParams struct {

}

type DescribeCiphersDetailRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCiphersDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCiphersDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCiphersDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCiphersDetailResponseParams struct {
	// 加密套件信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ciphers []*TLSCiphers `json:"Ciphers,omitnil" name:"Ciphers"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCiphersDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCiphersDetailResponseParams `json:"Response"`
}

func (r *DescribeCiphersDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCiphersDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomRuleListRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 偏移
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤数组,name可以是如下的值： RuleID,RuleName,Match
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`

	// asc或者desc
	Order *string `json:"Order,omitnil" name:"Order"`

	// exp_ts或者mod_ts
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeCustomRuleListRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 偏移
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤数组,name可以是如下的值： RuleID,RuleName,Match
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`

	// asc或者desc
	Order *string `json:"Order,omitnil" name:"Order"`

	// exp_ts或者mod_ts
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeCustomRuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomRuleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomRuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomRuleListResponseParams struct {
	// 规则详情
	RuleList []*DescribeCustomRulesRspRuleListItem `json:"RuleList,omitnil" name:"RuleList"`

	// 规则条数
	TotalCount *string `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCustomRuleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomRuleListResponseParams `json:"Response"`
}

func (r *DescribeCustomRuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomRulesRspRuleListItem struct {
	// 动作类型
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// 跳过的策略
	Bypass *string `json:"Bypass,omitnil" name:"Bypass"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 策略名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 重定向地址
	Redirect *string `json:"Redirect,omitnil" name:"Redirect"`

	// 策略ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`

	// 优先级
	SortId *string `json:"SortId,omitnil" name:"SortId"`

	// 状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// 策略详情
	Strategies []*Strategy `json:"Strategies,omitnil" name:"Strategies"`

	// 事件id
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventId *string `json:"EventId,omitnil" name:"EventId"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 生效状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidStatus *int64 `json:"ValidStatus,omitnil" name:"ValidStatus"`

	// 来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil" name:"Source"`
}

// Predefined struct for user
type DescribeCustomWhiteRuleRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤数组,name可以是如下的值： RuleID,RuleName,Match
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`

	// asc或者desc
	Order *string `json:"Order,omitnil" name:"Order"`

	// exp_ts或者mod_ts
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeCustomWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤数组,name可以是如下的值： RuleID,RuleName,Match
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`

	// asc或者desc
	Order *string `json:"Order,omitnil" name:"Order"`

	// exp_ts或者mod_ts
	By *string `json:"By,omitnil" name:"By"`
}

func (r *DescribeCustomWhiteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomWhiteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomWhiteRuleResponseParams struct {
	// 规则详情
	RuleList []*DescribeCustomRulesRspRuleListItem `json:"RuleList,omitnil" name:"RuleList"`

	// 规则条数
	TotalCount *string `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCustomWhiteRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomWhiteRuleResponseParams `json:"Response"`
}

func (r *DescribeCustomWhiteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainCountInfoRequestParams struct {

}

type DescribeDomainCountInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDomainCountInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainCountInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainCountInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainCountInfoResponseParams struct {
	// 域名总数
	AllDomain *uint64 `json:"AllDomain,omitnil" name:"AllDomain"`

	// 最近发现时间
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 接入域名总数
	WafDomainCount *uint64 `json:"WafDomainCount,omitnil" name:"WafDomainCount"`

	// 剩下配额
	LeftDomainCount *uint64 `json:"LeftDomainCount,omitnil" name:"LeftDomainCount"`

	// 开启防护域名数
	OpenWafDomain *uint64 `json:"OpenWafDomain,omitnil" name:"OpenWafDomain"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDomainCountInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainCountInfoResponseParams `json:"Response"`
}

func (r *DescribeDomainCountInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainCountInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainDetailsClbRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名id
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeDomainDetailsClbRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名id
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeDomainDetailsClbRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainDetailsClbRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainDetailsClbRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainDetailsClbResponseParams struct {
	// clb域名详情
	DomainsClbPartInfo *ClbDomainsInfo `json:"DomainsClbPartInfo,omitnil" name:"DomainsClbPartInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDomainDetailsClbResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainDetailsClbResponseParams `json:"Response"`
}

func (r *DescribeDomainDetailsClbResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainDetailsClbResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainDetailsSaasRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名id
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeDomainDetailsSaasRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名id
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeDomainDetailsSaasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainDetailsSaasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainDetailsSaasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainDetailsSaasResponseParams struct {
	// 域名详情
	DomainsPartInfo *DomainsPartInfo `json:"DomainsPartInfo,omitnil" name:"DomainsPartInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDomainDetailsSaasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainDetailsSaasResponseParams `json:"Response"`
}

func (r *DescribeDomainDetailsSaasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainDetailsSaasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainRulesRequestParams struct {
	// 需要查询的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type DescribeDomainRulesRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *DescribeDomainRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainRulesResponseParams struct {
	// 规则列表详情
	Rules []*Rule `json:"Rules,omitnil" name:"Rules"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDomainRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainRulesResponseParams `json:"Response"`
}

func (r *DescribeDomainRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainVerifyResultRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

type DescribeDomainVerifyResultRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

func (r *DescribeDomainVerifyResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainVerifyResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainVerifyResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainVerifyResultResponseParams struct {
	// 结果描述；如果可以添加返回空字符串
	Msg *string `json:"Msg,omitnil" name:"Msg"`

	// 检验状态：0表示可以添加，大于0为不能添加
	VerifyCode *int64 `json:"VerifyCode,omitnil" name:"VerifyCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDomainVerifyResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainVerifyResultResponseParams `json:"Response"`
}

func (r *DescribeDomainVerifyResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainVerifyResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainWhiteRulesRequestParams struct {
	// 需要查询的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 请求的白名单匹配路径
	Url *string `json:"Url,omitnil" name:"Url"`

	// 翻到多少页
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// 每页展示的条数
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 排序方式,desc表示降序，asc表示升序
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`
}

type DescribeDomainWhiteRulesRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 请求的白名单匹配路径
	Url *string `json:"Url,omitnil" name:"Url"`

	// 翻到多少页
	Page *uint64 `json:"Page,omitnil" name:"Page"`

	// 每页展示的条数
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 排序方式,desc表示降序，asc表示升序
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil" name:"RuleId"`
}

func (r *DescribeDomainWhiteRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainWhiteRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Url")
	delete(f, "Page")
	delete(f, "Count")
	delete(f, "Sort")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainWhiteRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainWhiteRulesResponseParams struct {
	// 规则列表
	RuleList []*RuleList `json:"RuleList,omitnil" name:"RuleList"`

	// 规则的数量
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDomainWhiteRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainWhiteRulesResponseParams `json:"Response"`
}

func (r *DescribeDomainWhiteRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainWhiteRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainsRequestParams struct {
	// 分页偏移量，取Limit整数倍。最小值为0，最大值= Total/Limit向上取整
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回域名的数量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤数组
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`
}

type DescribeDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，取Limit整数倍。最小值为0，最大值= Total/Limit向上取整
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回域名的数量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤数组
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainsResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// domain列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domains []*DomainInfo `json:"Domains,omitnil" name:"Domains"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainsResponseParams `json:"Response"`
}

func (r *DescribeDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFindDomainListRequestParams struct {
	// 分页
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页容量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件
	Key *string `json:"Key,omitnil" name:"Key"`

	// 是否接入waf
	IsWafDomain *string `json:"IsWafDomain,omitnil" name:"IsWafDomain"`

	// 排序参数
	By *string `json:"By,omitnil" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitnil" name:"Order"`
}

type DescribeFindDomainListRequest struct {
	*tchttp.BaseRequest
	
	// 分页
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页容量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤条件
	Key *string `json:"Key,omitnil" name:"Key"`

	// 是否接入waf
	IsWafDomain *string `json:"IsWafDomain,omitnil" name:"IsWafDomain"`

	// 排序参数
	By *string `json:"By,omitnil" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitnil" name:"Order"`
}

func (r *DescribeFindDomainListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFindDomainListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Key")
	delete(f, "IsWafDomain")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFindDomainListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFindDomainListResponseParams struct {
	// 域名总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 域名信息列表
	List []*FindAllDomainDetail `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeFindDomainListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFindDomainListResponseParams `json:"Response"`
}

func (r *DescribeFindDomainListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFindDomainListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowTrendRequestParams struct {
	// 需要获取流量趋势的域名, all表示所有域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 起始时间戳，精度秒
	StartTs *int64 `json:"StartTs,omitnil" name:"StartTs"`

	// 结束时间戳，精度秒
	EndTs *int64 `json:"EndTs,omitnil" name:"EndTs"`
}

type DescribeFlowTrendRequest struct {
	*tchttp.BaseRequest
	
	// 需要获取流量趋势的域名, all表示所有域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 起始时间戳，精度秒
	StartTs *int64 `json:"StartTs,omitnil" name:"StartTs"`

	// 结束时间戳，精度秒
	EndTs *int64 `json:"EndTs,omitnil" name:"EndTs"`
}

func (r *DescribeFlowTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "StartTs")
	delete(f, "EndTs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlowTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlowTrendResponseParams struct {
	// 流量趋势数据
	Data []*BotStatPointItem `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeFlowTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlowTrendResponseParams `json:"Response"`
}

func (r *DescribeFlowTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlowTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHistogramRequestParams struct {
	// 起始时间
	FromTime *string `json:"FromTime,omitnil" name:"FromTime"`

	// 结束时间
	ToTime *string `json:"ToTime,omitnil" name:"ToTime"`

	// 聚类字段，ip为ip聚合，art为响应耗时聚合，url为url聚合，local为ip转化的城市聚合
	QueryField *string `json:"QueryField,omitnil" name:"QueryField"`

	// 条件，access为访问日志，attack为攻击日志
	Source *string `json:"Source,omitnil" name:"Source"`

	// 兼容Host，逐步淘汰Host字段
	Host *string `json:"Host,omitnil" name:"Host"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 域名过滤，不传则不过滤，用于替代Host字段，逐步淘汰Host
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type DescribeHistogramRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间
	FromTime *string `json:"FromTime,omitnil" name:"FromTime"`

	// 结束时间
	ToTime *string `json:"ToTime,omitnil" name:"ToTime"`

	// 聚类字段，ip为ip聚合，art为响应耗时聚合，url为url聚合，local为ip转化的城市聚合
	QueryField *string `json:"QueryField,omitnil" name:"QueryField"`

	// 条件，access为访问日志，attack为攻击日志
	Source *string `json:"Source,omitnil" name:"Source"`

	// 兼容Host，逐步淘汰Host字段
	Host *string `json:"Host,omitnil" name:"Host"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 域名过滤，不传则不过滤，用于替代Host字段，逐步淘汰Host
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *DescribeHistogramRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistogramRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromTime")
	delete(f, "ToTime")
	delete(f, "QueryField")
	delete(f, "Source")
	delete(f, "Host")
	delete(f, "Edition")
	delete(f, "InstanceID")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHistogramRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHistogramResponseParams struct {
	// 统计数据
	Histogram []*string `json:"Histogram,omitnil" name:"Histogram"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeHistogramResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHistogramResponseParams `json:"Response"`
}

func (r *DescribeHistogramResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistogramResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostLimitRequestParams struct {
	// 添加的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 流量来源
	AlbType *string `json:"AlbType,omitnil" name:"AlbType"`
}

type DescribeHostLimitRequest struct {
	*tchttp.BaseRequest
	
	// 添加的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 流量来源
	AlbType *string `json:"AlbType,omitnil" name:"AlbType"`
}

func (r *DescribeHostLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "InstanceID")
	delete(f, "AlbType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostLimitResponseParams struct {
	// 成功返回的状态码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Success *ResponseCode `json:"Success,omitnil" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeHostLimitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostLimitResponseParams `json:"Response"`
}

func (r *DescribeHostLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

type DescribeHostRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

func (r *DescribeHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostResponseParams struct {
	// 域名详情
	Host *HostRecord `json:"Host,omitnil" name:"Host"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeHostResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostResponseParams `json:"Response"`
}

func (r *DescribeHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostsRequestParams struct {
	// 防护域名，如果是要查询某一具体的防护域名则传入此参数，要求是准确的域名，此参数不支持模糊搜索
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 防护域名ID，如果是要查询某一具体的防护域名则传入此参数，要求是准确的域名ID，此参数不支持模糊搜索
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 搜索条件，根据此参数对域名做模糊搜索
	Search *string `json:"Search,omitnil" name:"Search"`

	// 复杂的搜索条件
	Item *SearchItem `json:"Item,omitnil" name:"Item"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

type DescribeHostsRequest struct {
	*tchttp.BaseRequest
	
	// 防护域名，如果是要查询某一具体的防护域名则传入此参数，要求是准确的域名，此参数不支持模糊搜索
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 防护域名ID，如果是要查询某一具体的防护域名则传入此参数，要求是准确的域名ID，此参数不支持模糊搜索
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 搜索条件，根据此参数对域名做模糊搜索
	Search *string `json:"Search,omitnil" name:"Search"`

	// 复杂的搜索条件
	Item *SearchItem `json:"Item,omitnil" name:"Item"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

func (r *DescribeHostsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "Search")
	delete(f, "Item")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostsResponseParams struct {
	// 防护域名列表的长度
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 防护域名的列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostList []*HostRecord `json:"HostList,omitnil" name:"HostList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeHostsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostsResponseParams `json:"Response"`
}

func (r *DescribeHostsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤数组
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤数组
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// instance列表
	Instances []*InstanceInfo `json:"Instances,omitnil" name:"Instances"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesResponseParams `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpAccessControlRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 计数标识
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 动作，40表示查询白名单，42表示查询黑名单
	ActionType *uint64 `json:"ActionType,omitnil" name:"ActionType"`

	// 最小有效时间的时间戳
	//
	// Deprecated: VtsMin is deprecated.
	VtsMin *uint64 `json:"VtsMin,omitnil" name:"VtsMin"`

	// 最大有效时间的时间戳
	//
	// Deprecated: VtsMax is deprecated.
	VtsMax *uint64 `json:"VtsMax,omitnil" name:"VtsMax"`

	// 最小创建时间的时间戳
	CtsMin *uint64 `json:"CtsMin,omitnil" name:"CtsMin"`

	// 最大创建时间的时间戳
	CtsMax *uint64 `json:"CtsMax,omitnil" name:"CtsMax"`

	// 分页开始条数
	OffSet *uint64 `json:"OffSet,omitnil" name:"OffSet"`

	// 每页的条数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 来源
	Source *string `json:"Source,omitnil" name:"Source"`

	// 排序参数
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// ip
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 生效状态
	ValidStatus *int64 `json:"ValidStatus,omitnil" name:"ValidStatus"`

	// 最小有效时间的时间戳
	ValidTimeStampMin *string `json:"ValidTimeStampMin,omitnil" name:"ValidTimeStampMin"`

	// 最大有效时间的时间戳
	ValidTimeStampMax *string `json:"ValidTimeStampMax,omitnil" name:"ValidTimeStampMax"`
}

type DescribeIpAccessControlRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 计数标识
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 动作，40表示查询白名单，42表示查询黑名单
	ActionType *uint64 `json:"ActionType,omitnil" name:"ActionType"`

	// 最小有效时间的时间戳
	VtsMin *uint64 `json:"VtsMin,omitnil" name:"VtsMin"`

	// 最大有效时间的时间戳
	VtsMax *uint64 `json:"VtsMax,omitnil" name:"VtsMax"`

	// 最小创建时间的时间戳
	CtsMin *uint64 `json:"CtsMin,omitnil" name:"CtsMin"`

	// 最大创建时间的时间戳
	CtsMax *uint64 `json:"CtsMax,omitnil" name:"CtsMax"`

	// 分页开始条数
	OffSet *uint64 `json:"OffSet,omitnil" name:"OffSet"`

	// 每页的条数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 来源
	Source *string `json:"Source,omitnil" name:"Source"`

	// 排序参数
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// ip
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 生效状态
	ValidStatus *int64 `json:"ValidStatus,omitnil" name:"ValidStatus"`

	// 最小有效时间的时间戳
	ValidTimeStampMin *string `json:"ValidTimeStampMin,omitnil" name:"ValidTimeStampMin"`

	// 最大有效时间的时间戳
	ValidTimeStampMax *string `json:"ValidTimeStampMax,omitnil" name:"ValidTimeStampMax"`
}

func (r *DescribeIpAccessControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpAccessControlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Count")
	delete(f, "ActionType")
	delete(f, "VtsMin")
	delete(f, "VtsMax")
	delete(f, "CtsMin")
	delete(f, "CtsMax")
	delete(f, "OffSet")
	delete(f, "Limit")
	delete(f, "Source")
	delete(f, "Sort")
	delete(f, "Ip")
	delete(f, "ValidStatus")
	delete(f, "ValidTimeStampMin")
	delete(f, "ValidTimeStampMax")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpAccessControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpAccessControlResponseParams struct {
	// 输出
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *IpAccessControlData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeIpAccessControlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpAccessControlResponseParams `json:"Response"`
}

func (r *DescribeIpAccessControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpAccessControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpHitItemsRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 计数标识
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 类别
	Category *string `json:"Category,omitnil" name:"Category"`

	// 有效时间最小时间戳
	//
	// Deprecated: VtsMin is deprecated.
	VtsMin *uint64 `json:"VtsMin,omitnil" name:"VtsMin"`

	// 有效时间最大时间戳
	//
	// Deprecated: VtsMax is deprecated.
	VtsMax *uint64 `json:"VtsMax,omitnil" name:"VtsMax"`

	// 创建时间最小时间戳
	CtsMin *uint64 `json:"CtsMin,omitnil" name:"CtsMin"`

	// 创建时间最大时间戳
	CtsMax *uint64 `json:"CtsMax,omitnil" name:"CtsMax"`

	// 偏移参数
	Skip *uint64 `json:"Skip,omitnil" name:"Skip"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 策略名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 排序参数
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// IP
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 有效时间最小时间戳
	ValidTimeStampMin *uint64 `json:"ValidTimeStampMin,omitnil" name:"ValidTimeStampMin"`

	// 有效时间最大时间戳
	ValidTimeStampMax *uint64 `json:"ValidTimeStampMax,omitnil" name:"ValidTimeStampMax"`
}

type DescribeIpHitItemsRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 计数标识
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 类别
	Category *string `json:"Category,omitnil" name:"Category"`

	// 有效时间最小时间戳
	VtsMin *uint64 `json:"VtsMin,omitnil" name:"VtsMin"`

	// 有效时间最大时间戳
	VtsMax *uint64 `json:"VtsMax,omitnil" name:"VtsMax"`

	// 创建时间最小时间戳
	CtsMin *uint64 `json:"CtsMin,omitnil" name:"CtsMin"`

	// 创建时间最大时间戳
	CtsMax *uint64 `json:"CtsMax,omitnil" name:"CtsMax"`

	// 偏移参数
	Skip *uint64 `json:"Skip,omitnil" name:"Skip"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 策略名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 排序参数
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// IP
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 有效时间最小时间戳
	ValidTimeStampMin *uint64 `json:"ValidTimeStampMin,omitnil" name:"ValidTimeStampMin"`

	// 有效时间最大时间戳
	ValidTimeStampMax *uint64 `json:"ValidTimeStampMax,omitnil" name:"ValidTimeStampMax"`
}

func (r *DescribeIpHitItemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpHitItemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Count")
	delete(f, "Category")
	delete(f, "VtsMin")
	delete(f, "VtsMax")
	delete(f, "CtsMin")
	delete(f, "CtsMax")
	delete(f, "Skip")
	delete(f, "Limit")
	delete(f, "Name")
	delete(f, "Sort")
	delete(f, "Ip")
	delete(f, "ValidTimeStampMin")
	delete(f, "ValidTimeStampMax")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpHitItemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpHitItemsResponseParams struct {
	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *IpHitItemsData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeIpHitItemsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpHitItemsResponseParams `json:"Response"`
}

func (r *DescribeIpHitItemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpHitItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModuleStatusRequestParams struct {
	// 要查询状态的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type DescribeModuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 要查询状态的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *DescribeModuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModuleStatusResponseParams struct {
	// WEB安全规则是否开启
	WebSecurity *uint64 `json:"WebSecurity,omitnil" name:"WebSecurity"`

	// 访问控制规则是否开启
	AccessControl *int64 `json:"AccessControl,omitnil" name:"AccessControl"`

	// CC防护是否开启
	CcProtection *uint64 `json:"CcProtection,omitnil" name:"CcProtection"`

	// 网页防篡改是否开启
	AntiTamper *uint64 `json:"AntiTamper,omitnil" name:"AntiTamper"`

	// 信息防泄漏是否开启
	AntiLeakage *uint64 `json:"AntiLeakage,omitnil" name:"AntiLeakage"`

	// API安全是否开启
	ApiProtection *uint64 `json:"ApiProtection,omitnil" name:"ApiProtection"`

	// 限流模块开关
	RateLimit *uint64 `json:"RateLimit,omitnil" name:"RateLimit"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeModuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModuleStatusResponseParams `json:"Response"`
}

func (r *DescribeModuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeObjectsRequestParams struct {
	// 支持的过滤器:
	// 	ObjectId: clb实例ID
	// 	VIP: clb实例的公网IP
	// 	InstanceId: waf实例ID
	// 	Domain: 精准域名
	// 	Status: waf防护开关状态: 0关闭，1开启
	// 	ClsStatus: waf日志开关: 0关闭，1开启
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`
}

type DescribeObjectsRequest struct {
	*tchttp.BaseRequest
	
	// 支持的过滤器:
	// 	ObjectId: clb实例ID
	// 	VIP: clb实例的公网IP
	// 	InstanceId: waf实例ID
	// 	Domain: 精准域名
	// 	Status: waf防护开关状态: 0关闭，1开启
	// 	ClsStatus: waf日志开关: 0关闭，1开启
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeObjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeObjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeObjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeObjectsResponseParams struct {
	// 对象列表
	ClbObjects []*ClbObject `json:"ClbObjects,omitnil" name:"ClbObjects"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeObjectsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeObjectsResponseParams `json:"Response"`
}

func (r *DescribeObjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeObjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePeakPointsRequestParams struct {
	// 查询起始时间
	FromTime *string `json:"FromTime,omitnil" name:"FromTime"`

	// 查询终止时间
	ToTime *string `json:"ToTime,omitnil" name:"ToTime"`

	// 查询的域名，如果查询所有域名数据，该参数不填写
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 十二个值可选：
	// access-峰值qps趋势图
	// botAccess- bot峰值qps趋势图
	// down-下行峰值带宽趋势图
	// up-上行峰值带宽趋势图
	// attack-Web攻击总数趋势图
	// cc-CC攻击总数趋势图
	// bw-黑IP攻击总数趋势图
	// tamper-防篡改攻击总数趋势图
	// leak-防泄露攻击总数趋势图
	// acl-访问控制攻击总数趋势图
	// http_status-状态码各次数趋势图
	// wx_access-微信小程序峰值qps趋势图
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`
}

type DescribePeakPointsRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间
	FromTime *string `json:"FromTime,omitnil" name:"FromTime"`

	// 查询终止时间
	ToTime *string `json:"ToTime,omitnil" name:"ToTime"`

	// 查询的域名，如果查询所有域名数据，该参数不填写
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 十二个值可选：
	// access-峰值qps趋势图
	// botAccess- bot峰值qps趋势图
	// down-下行峰值带宽趋势图
	// up-上行峰值带宽趋势图
	// attack-Web攻击总数趋势图
	// cc-CC攻击总数趋势图
	// bw-黑IP攻击总数趋势图
	// tamper-防篡改攻击总数趋势图
	// leak-防泄露攻击总数趋势图
	// acl-访问控制攻击总数趋势图
	// http_status-状态码各次数趋势图
	// wx_access-微信小程序峰值qps趋势图
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`
}

func (r *DescribePeakPointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePeakPointsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromTime")
	delete(f, "ToTime")
	delete(f, "Domain")
	delete(f, "Edition")
	delete(f, "InstanceID")
	delete(f, "MetricName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePeakPointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePeakPointsResponseParams struct {
	// 数据点
	Points []*PeakPointsItem `json:"Points,omitnil" name:"Points"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePeakPointsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePeakPointsResponseParams `json:"Response"`
}

func (r *DescribePeakPointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePeakPointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePeakValueRequestParams struct {
	// 查询起始时间
	FromTime *string `json:"FromTime,omitnil" name:"FromTime"`

	// 查询结束时间
	ToTime *string `json:"ToTime,omitnil" name:"ToTime"`

	// 需要查询的域名，当前用户所有域名可以不传
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 五个值可选：
	// access-峰值qps
	// down-下行峰值带宽
	// up-上行峰值带宽
	// attack-Web攻击总数
	// cc-CC攻击总数趋势图
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`
}

type DescribePeakValueRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间
	FromTime *string `json:"FromTime,omitnil" name:"FromTime"`

	// 查询结束时间
	ToTime *string `json:"ToTime,omitnil" name:"ToTime"`

	// 需要查询的域名，当前用户所有域名可以不传
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 五个值可选：
	// access-峰值qps
	// down-下行峰值带宽
	// up-上行峰值带宽
	// attack-Web攻击总数
	// cc-CC攻击总数趋势图
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`
}

func (r *DescribePeakValueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePeakValueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromTime")
	delete(f, "ToTime")
	delete(f, "Domain")
	delete(f, "Edition")
	delete(f, "InstanceID")
	delete(f, "MetricName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePeakValueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePeakValueResponseParams struct {
	// QPS峰值
	Access *uint64 `json:"Access,omitnil" name:"Access"`

	// 上行带宽峰值，单位B
	Up *uint64 `json:"Up,omitnil" name:"Up"`

	// 下行带宽峰值，单位B
	Down *uint64 `json:"Down,omitnil" name:"Down"`

	// Web攻击总数
	Attack *uint64 `json:"Attack,omitnil" name:"Attack"`

	// CC攻击总数
	Cc *uint64 `json:"Cc,omitnil" name:"Cc"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePeakValueResponse struct {
	*tchttp.BaseResponse
	Response *DescribePeakValueResponseParams `json:"Response"`
}

func (r *DescribePeakValueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePeakValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyStatusRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// clb-waf或者saas-waf
	Edition *string `json:"Edition,omitnil" name:"Edition"`
}

type DescribePolicyStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// clb-waf或者saas-waf
	Edition *string `json:"Edition,omitnil" name:"Edition"`
}

func (r *DescribePolicyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Edition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePolicyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyStatusResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 防护状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePolicyStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribePolicyStatusResponseParams `json:"Response"`
}

func (r *DescribePolicyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePortsRequestParams struct {
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 实例类型
	Edition *string `json:"Edition,omitnil" name:"Edition"`
}

type DescribePortsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 实例类型
	Edition *string `json:"Edition,omitnil" name:"Edition"`
}

func (r *DescribePortsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePortsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceID")
	delete(f, "Edition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePortsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePortsResponseParams struct {
	// http端口列表
	HttpPorts []*string `json:"HttpPorts,omitnil" name:"HttpPorts"`

	// https端口列表
	HttpsPorts []*string `json:"HttpsPorts,omitnil" name:"HttpsPorts"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePortsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePortsResponseParams `json:"Response"`
}

func (r *DescribePortsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePortsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleLimitRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeRuleLimitRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeRuleLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleLimitResponseParams struct {
	// waf模块的规格
	Res *WafRuleLimit `json:"Res,omitnil" name:"Res"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRuleLimitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleLimitResponseParams `json:"Response"`
}

func (r *DescribeRuleLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// clb-waf或者sparta-waf
	Edition *string `json:"Edition,omitnil" name:"Edition"`
}

type DescribeSessionRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// clb-waf或者sparta-waf
	Edition *string `json:"Edition,omitnil" name:"Edition"`
}

func (r *DescribeSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Edition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionResponseParams struct {
	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *SessionData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSessionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSessionResponseParams `json:"Response"`
}

func (r *DescribeSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpartaProtectionInfoRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 版本
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 实例
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

type DescribeSpartaProtectionInfoRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 版本
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 实例
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

func (r *DescribeSpartaProtectionInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpartaProtectionInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Edition")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpartaProtectionInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpartaProtectionInfoResponseParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// cname取值
	Cname *string `json:"Cname,omitnil" name:"Cname"`

	// 状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// 源IP地址列表
	SrcList []*string `json:"SrcList,omitnil" name:"SrcList"`

	// 证书类型
	CertType *string `json:"CertType,omitnil" name:"CertType"`

	// 证书
	Cert *string `json:"Cert,omitnil" name:"Cert"`

	// 私有密钥
	PrivateKey *string `json:"PrivateKey,omitnil" name:"PrivateKey"`

	// ssl的id
	Sslid *string `json:"Sslid,omitnil" name:"Sslid"`

	// 是否是cdn
	IsCdn *string `json:"IsCdn,omitnil" name:"IsCdn"`

	// 灰度区域列表
	GrayAreas []*string `json:"GrayAreas,omitnil" name:"GrayAreas"`

	// 引擎
	Engine *string `json:"Engine,omitnil" name:"Engine"`

	// HTTPS重写
	HttpsRewrite *string `json:"HttpsRewrite,omitnil" name:"HttpsRewrite"`

	// upstreamType取值
	UpstreamType *string `json:"UpstreamType,omitnil" name:"UpstreamType"`

	// upstreamDomain取值
	UpstreamDomain *string `json:"UpstreamDomain,omitnil" name:"UpstreamDomain"`

	// upstreamScheme取值
	UpstreamScheme *string `json:"UpstreamScheme,omitnil" name:"UpstreamScheme"`

	// 是否是HTTP2
	IsHttp2 *string `json:"IsHttp2,omitnil" name:"IsHttp2"`

	// 是否含有websocket
	IsWebsocket *string `json:"IsWebsocket,omitnil" name:"IsWebsocket"`

	// loadBalance信息
	LoadBalance *string `json:"LoadBalance,omitnil" name:"LoadBalance"`

	// httpsUpstreamPort取值
	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitnil" name:"HttpsUpstreamPort"`

	// port信息
	Ports []*PortItem `json:"Ports,omitnil" name:"Ports"`

	// 是否灰度
	IsGray *string `json:"IsGray,omitnil" name:"IsGray"`

	// 模式
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 防御等级,100,200,300
	Level *string `json:"Level,omitnil" name:"Level"`

	// 与源站是否保持长连接
	IsKeepAlive *string `json:"IsKeepAlive,omitnil" name:"IsKeepAlive"`

	// 0：BGP 1：Anycast
	// 注意：此字段可能返回 null，表示取不到有效值。
	Anycast *string `json:"Anycast,omitnil" name:"Anycast"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSpartaProtectionInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpartaProtectionInfoResponseParams `json:"Response"`
}

func (r *DescribeSpartaProtectionInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpartaProtectionInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTlsVersionRequestParams struct {

}

type DescribeTlsVersionRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeTlsVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTlsVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTlsVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTlsVersionResponseParams struct {
	// TLS key value
	TLS []*TLSVersion `json:"TLS,omitnil" name:"TLS"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTlsVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTlsVersionResponseParams `json:"Response"`
}

func (r *DescribeTlsVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTlsVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopAttackDomainRequestParams struct {
	// 查询起始时间
	FromTime *string `json:"FromTime,omitnil" name:"FromTime"`

	// 查询结束时间
	ToTime *string `json:"ToTime,omitnil" name:"ToTime"`

	// TOP N,可从0-10选择，默认是10
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

type DescribeTopAttackDomainRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间
	FromTime *string `json:"FromTime,omitnil" name:"FromTime"`

	// 查询结束时间
	ToTime *string `json:"ToTime,omitnil" name:"ToTime"`

	// TOP N,可从0-10选择，默认是10
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

func (r *DescribeTopAttackDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopAttackDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FromTime")
	delete(f, "ToTime")
	delete(f, "Count")
	delete(f, "Edition")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopAttackDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopAttackDomainResponseParams struct {
	// CC攻击域名列表
	CC []*KVInt `json:"CC,omitnil" name:"CC"`

	// Web攻击域名列表
	Web []*KVInt `json:"Web,omitnil" name:"Web"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTopAttackDomainResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopAttackDomainResponseParams `json:"Response"`
}

func (r *DescribeTopAttackDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopAttackDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserCdcClbWafRegionsRequestParams struct {

}

type DescribeUserCdcClbWafRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserCdcClbWafRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserCdcClbWafRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserCdcClbWafRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserCdcClbWafRegionsResponseParams struct {
	// CdcRegion的类型描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*CdcRegion `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUserCdcClbWafRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserCdcClbWafRegionsResponseParams `json:"Response"`
}

func (r *DescribeUserCdcClbWafRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserCdcClbWafRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserClbWafRegionsRequestParams struct {

}

type DescribeUserClbWafRegionsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserClbWafRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserClbWafRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserClbWafRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserClbWafRegionsResponseParams struct {
	// 地域（标准的ap-格式）列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*string `json:"Data,omitnil" name:"Data"`

	// 包含详细属性的地域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RichDatas []*ClbWafRegionItem `json:"RichDatas,omitnil" name:"RichDatas"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUserClbWafRegionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserClbWafRegionsResponseParams `json:"Response"`
}

func (r *DescribeUserClbWafRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserClbWafRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserDomainInfoRequestParams struct {

}

type DescribeUserDomainInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserDomainInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDomainInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserDomainInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserDomainInfoResponseParams struct {
	// saas和clb域名信息
	UsersInfo []*UserDomainInfo `json:"UsersInfo,omitnil" name:"UsersInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUserDomainInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserDomainInfoResponseParams `json:"Response"`
}

func (r *DescribeUserDomainInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDomainInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserLevelRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type DescribeUserLevelRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *DescribeUserLevelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserLevelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserLevelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserLevelResponseParams struct {
	// 300:正常 400:严格
	Level *uint64 `json:"Level,omitnil" name:"Level"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUserLevelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserLevelResponseParams `json:"Response"`
}

func (r *DescribeUserLevelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserSignatureRuleRequestParams struct {
	// 需要查询的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 分页
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页容量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段，支持 signature_id, modify_time
	By *string `json:"By,omitnil" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitnil" name:"Order"`

	// 筛选条件，支持 MainClassName，SubClassID ,CveID, Status, ID;  ID为规则id
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`
}

type DescribeUserSignatureRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 分页
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页容量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段，支持 signature_id, modify_time
	By *string `json:"By,omitnil" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitnil" name:"Order"`

	// 筛选条件，支持 MainClassName，SubClassID ,CveID, Status, ID;  ID为规则id
	Filters []*FiltersItemNew `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeUserSignatureRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserSignatureRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "By")
	delete(f, "Order")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserSignatureRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserSignatureRuleResponseParams struct {
	// 规则总数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*UserSignatureRule `json:"Rules,omitnil" name:"Rules"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeUserSignatureRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserSignatureRuleResponseParams `json:"Response"`
}

func (r *DescribeUserSignatureRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserSignatureRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVipInfoRequestParams struct {
	// waf实例id列表
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

type DescribeVipInfoRequest struct {
	*tchttp.BaseRequest
	
	// waf实例id列表
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`
}

func (r *DescribeVipInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVipInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVipInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVipInfoResponseParams struct {
	// VIP信息
	VipInfo []*VipInfo `json:"VipInfo,omitnil" name:"VipInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVipInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVipInfoResponseParams `json:"Response"`
}

func (r *DescribeVipInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVipInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafAutoDenyRulesRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeWafAutoDenyRulesRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DescribeWafAutoDenyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafAutoDenyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWafAutoDenyRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafAutoDenyRulesResponseParams struct {
	// 攻击次数阈值
	AttackThreshold *int64 `json:"AttackThreshold,omitnil" name:"AttackThreshold"`

	// 攻击时间阈值
	TimeThreshold *int64 `json:"TimeThreshold,omitnil" name:"TimeThreshold"`

	// 自动封禁时间
	DenyTimeThreshold *int64 `json:"DenyTimeThreshold,omitnil" name:"DenyTimeThreshold"`

	// 自动封禁状态
	DefenseStatus *int64 `json:"DefenseStatus,omitnil" name:"DefenseStatus"`

	// 重保护网域名状态
	HWState *int64 `json:"HWState,omitnil" name:"HWState"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWafAutoDenyRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWafAutoDenyRulesResponseParams `json:"Response"`
}

func (r *DescribeWafAutoDenyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafAutoDenyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafAutoDenyStatusRequestParams struct {

}

type DescribeWafAutoDenyStatusRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeWafAutoDenyStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafAutoDenyStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWafAutoDenyStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafAutoDenyStatusResponseParams struct {
	// WAF 自动封禁详情
	WafAutoDenyDetails *AutoDenyDetail `json:"WafAutoDenyDetails,omitnil" name:"WafAutoDenyDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWafAutoDenyStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWafAutoDenyStatusResponseParams `json:"Response"`
}

func (r *DescribeWafAutoDenyStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafAutoDenyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafInfoRequestParams struct {
	// CLB回调WAF接口（获取、删除）的参数
	Params []*ClbHostsParams `json:"Params,omitnil" name:"Params"`
}

type DescribeWafInfoRequest struct {
	*tchttp.BaseRequest
	
	// CLB回调WAF接口（获取、删除）的参数
	Params []*ClbHostsParams `json:"Params,omitnil" name:"Params"`
}

func (r *DescribeWafInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Params")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWafInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafInfoResponseParams struct {
	// 返回的WAF信息数组的长度，为0则表示没有查询到对应的信息
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 对应的WAF信息的数组。
	HostList []*ClbHostResult `json:"HostList,omitnil" name:"HostList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWafInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWafInfoResponseParams `json:"Response"`
}

func (r *DescribeWafInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafThreatenIntelligenceRequestParams struct {

}

type DescribeWafThreatenIntelligenceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeWafThreatenIntelligenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafThreatenIntelligenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWafThreatenIntelligenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafThreatenIntelligenceResponseParams struct {
	// WAF 威胁情报封禁信息
	WafThreatenIntelligenceDetails *WafThreatenIntelligenceDetails `json:"WafThreatenIntelligenceDetails,omitnil" name:"WafThreatenIntelligenceDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWafThreatenIntelligenceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWafThreatenIntelligenceResponseParams `json:"Response"`
}

func (r *DescribeWafThreatenIntelligenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafThreatenIntelligenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebshellStatusRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type DescribeWebshellStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *DescribeWebshellStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebshellStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWebshellStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWebshellStatusResponseParams struct {
	// webshell域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 开关状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeWebshellStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWebshellStatusResponseParams `json:"Response"`
}

func (r *DescribeWebshellStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWebshellStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainInfo struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// cname地址
	Cname *string `json:"Cname,omitnil" name:"Cname"`

	// 实例类型,sparta-waf表示saaswaf实例域名,clb-waf表示clbwaf实例域名,cdc-clb-waf表示CDC环境下clbwaf实例域名,cdn-waf表示cdnwaf实例域名
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 地域
	Region *string `json:"Region,omitnil" name:"Region"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 日志包
	ClsStatus *uint64 `json:"ClsStatus,omitnil" name:"ClsStatus"`

	// clbwaf使用模式,0镜像模式 1清洗模式
	FlowMode *uint64 `json:"FlowMode,omitnil" name:"FlowMode"`

	// waf开关,0关闭 1开启
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 规则引擎防护模式,0观察模式 1拦截模式
	Mode *uint64 `json:"Mode,omitnil" name:"Mode"`

	// 规则引擎和AI引擎防护模式联合状态,10规则引擎观察&&AI引擎关闭模式 11规则引擎观察&&AI引擎观察模式 12规则引擎观察&&AI引擎拦截模式 20规则引擎拦截&&AI引擎关闭模式 21规则引擎拦截&&AI引擎观察模式 22规则引擎拦截&&AI引擎拦截模式
	Engine *uint64 `json:"Engine,omitnil" name:"Engine"`

	// CC列表
	CCList []*string `json:"CCList,omitnil" name:"CCList"`

	// 回源ip
	RsList []*string `json:"RsList,omitnil" name:"RsList"`

	// 服务端口配置
	Ports []*PortInfo `json:"Ports,omitnil" name:"Ports"`

	// 负载均衡器
	LoadBalancerSet []*LoadBalancerPackageNew `json:"LoadBalancerSet,omitnil" name:"LoadBalancerSet"`

	// 用户id
	AppId *uint64 `json:"AppId,omitnil" name:"AppId"`

	// clbwaf域名监听器状态,0操作成功 4正在绑定LB 6正在解绑LB 7解绑LB失败 8绑定LB失败 10内部错误
	State *int64 `json:"State,omitnil" name:"State"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Ipv6开关状态,0关闭 1开启
	Ipv6Status *int64 `json:"Ipv6Status,omitnil" name:"Ipv6Status"`

	// BOT开关状态,0关闭 1开启
	BotStatus *int64 `json:"BotStatus,omitnil" name:"BotStatus"`

	// 版本信息
	Level *int64 `json:"Level,omitnil" name:"Level"`

	// 是否开启投递CLS功能,0关闭 1开启
	PostCLSStatus *int64 `json:"PostCLSStatus,omitnil" name:"PostCLSStatus"`

	// 是否开启投递CKafka功能,0关闭 1开启
	PostCKafkaStatus *int64 `json:"PostCKafkaStatus,omitnil" name:"PostCKafkaStatus"`

	// cdc实例域名接入的集群信息,非cdc实例忽略
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdcClusters *string `json:"CdcClusters,omitnil" name:"CdcClusters"`

	// api安全开关状态,0关闭 1开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiStatus *int64 `json:"ApiStatus,omitnil" name:"ApiStatus"`

	// 应用型负载均衡类型,clb或者apisix，默认clb
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlbType *string `json:"AlbType,omitnil" name:"AlbType"`

	// 安全组状态,0不展示 1非腾讯云源站 2安全组绑定失败 3安全组发生变更
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgState *int64 `json:"SgState,omitnil" name:"SgState"`

	// 安全组状态的详细解释
	// 注意：此字段可能返回 null，表示取不到有效值。
	SgDetail *string `json:"SgDetail,omitnil" name:"SgDetail"`

	// 域名类型:hybrid表示混合云域名，public表示公有云域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloudType *string `json:"CloudType,omitnil" name:"CloudType"`
}

type DomainPackageNew struct {
	// 资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceIds *string `json:"ResourceIds,omitnil" name:"ResourceIds"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidTime *string `json:"ValidTime,omitnil" name:"ValidTime"`

	// 是否自动续费，1：自动续费，0：不自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *uint64 `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// 套餐购买个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 套餐购买地域，clb-waf暂时没有用到
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`
}

type DomainURI struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 版本
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

type DomainsPartInfo struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名id
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 类型
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 证书
	Cert *string `json:"Cert,omitnil" name:"Cert"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// AI防御模式
	Engine *uint64 `json:"Engine,omitnil" name:"Engine"`

	// 是否开启httpRewrite
	HttpsRewrite *uint64 `json:"HttpsRewrite,omitnil" name:"HttpsRewrite"`

	// https回源端口
	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitnil" name:"HttpsUpstreamPort"`

	// 是否是cdn
	IsCdn *uint64 `json:"IsCdn,omitnil" name:"IsCdn"`

	// 是否开启gray
	IsGray *uint64 `json:"IsGray,omitnil" name:"IsGray"`

	// 是否是http2
	IsHttp2 *uint64 `json:"IsHttp2,omitnil" name:"IsHttp2"`

	// 是否开启websocket
	IsWebsocket *uint64 `json:"IsWebsocket,omitnil" name:"IsWebsocket"`

	// 负载均衡
	LoadBalance *uint64 `json:"LoadBalance,omitnil" name:"LoadBalance"`

	// 防御模式
	Mode *uint64 `json:"Mode,omitnil" name:"Mode"`

	// 私钥
	PrivateKey *string `json:"PrivateKey,omitnil" name:"PrivateKey"`

	// ssl id
	SSLId *string `json:"SSLId,omitnil" name:"SSLId"`

	// 回源域名
	UpstreamDomain *string `json:"UpstreamDomain,omitnil" name:"UpstreamDomain"`

	// 回源类型
	UpstreamType *uint64 `json:"UpstreamType,omitnil" name:"UpstreamType"`

	// 回源ip
	SrcList []*string `json:"SrcList,omitnil" name:"SrcList"`

	// 服务端口配置
	Ports []*PortInfo `json:"Ports,omitnil" name:"Ports"`

	// 证书类型
	CertType *uint64 `json:"CertType,omitnil" name:"CertType"`

	// 回源方式
	UpstreamScheme *string `json:"UpstreamScheme,omitnil" name:"UpstreamScheme"`

	// 日志包
	Cls *uint64 `json:"Cls,omitnil" name:"Cls"`

	// 一级cname
	Cname *string `json:"Cname,omitnil" name:"Cname"`

	// 是否长连接
	IsKeepAlive *uint64 `json:"IsKeepAlive,omitnil" name:"IsKeepAlive"`

	// 是否开启主动健康检测，1表示开启，0表示不开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveCheck *uint64 `json:"ActiveCheck,omitnil" name:"ActiveCheck"`

	// TLS版本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TLSVersion *int64 `json:"TLSVersion,omitnil" name:"TLSVersion"`

	// 加密套件信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ciphers []*int64 `json:"Ciphers,omitnil" name:"Ciphers"`

	// 模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	CipherTemplate *int64 `json:"CipherTemplate,omitnil" name:"CipherTemplate"`

	// 300s
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxyReadTimeout *int64 `json:"ProxyReadTimeout,omitnil" name:"ProxyReadTimeout"`

	// 300s
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProxySendTimeout *int64 `json:"ProxySendTimeout,omitnil" name:"ProxySendTimeout"`

	// 0:关闭SNI；1:开启SNI，SNI=源请求host；2:开启SNI，SNI=修改为源站host；3：开启SNI，自定义host，SNI=SniHost；
	// 注意：此字段可能返回 null，表示取不到有效值。
	SniType *int64 `json:"SniType,omitnil" name:"SniType"`

	// SniType=3时，需要填此参数，表示自定义的host；
	// 注意：此字段可能返回 null，表示取不到有效值。
	SniHost *string `json:"SniHost,omitnil" name:"SniHost"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weights []*string `json:"Weights,omitnil" name:"Weights"`

	// IsCdn=3时，表示自定义header
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpHeaders []*string `json:"IpHeaders,omitnil" name:"IpHeaders"`

	// 0:关闭xff重置；1:开启xff重置
	// 注意：此字段可能返回 null，表示取不到有效值。
	XFFReset *int64 `json:"XFFReset,omitnil" name:"XFFReset"`
}

type DownloadAttackRecordInfo struct {
	// 记录ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 下载任务名
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`

	// 域名
	Host *string `json:"Host,omitnil" name:"Host"`

	// 当前下载任务的日志条数
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 下载任务运行状态：-1-下载超时，0-下载等待，1-下载完成，2-下载失败，4-正在下载
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 下载文件URL
	Url *string `json:"Url,omitnil" name:"Url"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 最后更新修改时间
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 下载任务需下载的日志总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`
}

type ExportAccessInfo struct {
	// 日志导出任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExportId *string `json:"ExportId,omitnil" name:"ExportId"`

	// 日志导出查询语句
	// 注意：此字段可能返回 null，表示取不到有效值。
	Query *string `json:"Query,omitnil" name:"Query"`

	// 日志导出文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil" name:"FileName"`

	// 日志文件大小
	FileSize *int64 `json:"FileSize,omitnil" name:"FileSize"`

	// 日志导出时间排序
	// 注意：此字段可能返回 null，表示取不到有效值。
	Order *string `json:"Order,omitnil" name:"Order"`

	// 日志导出格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Format *string `json:"Format,omitnil" name:"Format"`

	// 日志导出数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 日志下载状态。Processing:导出正在进行中，Complete:导出完成，Failed:导出失败，Expired:日志导出已过期（三天有效期）
	Status *string `json:"Status,omitnil" name:"Status"`

	// 日志导出起始时间
	From *int64 `json:"From,omitnil" name:"From"`

	// 日志导出结束时间
	To *int64 `json:"To,omitnil" name:"To"`

	// 日志导出路径
	CosPath *string `json:"CosPath,omitnil" name:"CosPath"`

	// 日志导出创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`
}

type FiltersItemNew struct {
	// 字段名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 过滤值
	Values []*string `json:"Values,omitnil" name:"Values"`

	// 是否精确查找
	ExactMatch *bool `json:"ExactMatch,omitnil" name:"ExactMatch"`
}

type FindAllDomainDetail struct {
	// 用户id
	Appid *uint64 `json:"Appid,omitnil" name:"Appid"`

	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名ip
	Ips []*string `json:"Ips,omitnil" name:"Ips"`

	// 发现时间
	FindTime *string `json:"FindTime,omitnil" name:"FindTime"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 域名id
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// waf类型
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 是否接入waf
	IsWafDomain *uint64 `json:"IsWafDomain,omitnil" name:"IsWafDomain"`
}

type FraudPkg struct {
	// 资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceIds *string `json:"ResourceIds,omitnil" name:"ResourceIds"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *int64 `json:"Region,omitnil" name:"Region"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 申请数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InquireNum *int64 `json:"InquireNum,omitnil" name:"InquireNum"`

	// 使用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedNum *int64 `json:"UsedNum,omitnil" name:"UsedNum"`

	// 续费标志
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *uint64 `json:"RenewFlag,omitnil" name:"RenewFlag"`
}

// Predefined struct for user
type FreshAntiFakeUrlRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Id
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

type FreshAntiFakeUrlRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Id
	Id *uint64 `json:"Id,omitnil" name:"Id"`
}

func (r *FreshAntiFakeUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FreshAntiFakeUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FreshAntiFakeUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FreshAntiFakeUrlResponseParams struct {
	// 结果成功与否
	Result *string `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type FreshAntiFakeUrlResponse struct {
	*tchttp.BaseResponse
	Response *FreshAntiFakeUrlResponseParams `json:"Response"`
}

func (r *FreshAntiFakeUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FreshAntiFakeUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateDealsAndPayNewRequestParams struct {
	// 计费下单入参
	Goods []*GoodNews `json:"Goods,omitnil" name:"Goods"`
}

type GenerateDealsAndPayNewRequest struct {
	*tchttp.BaseRequest
	
	// 计费下单入参
	Goods []*GoodNews `json:"Goods,omitnil" name:"Goods"`
}

func (r *GenerateDealsAndPayNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateDealsAndPayNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Goods")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateDealsAndPayNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateDealsAndPayNewResponseParams struct {
	// 计费下单响应结构体
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DealData `json:"Data,omitnil" name:"Data"`

	// 1:成功，0:失败
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 返回message
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnMessage *string `json:"ReturnMessage,omitnil" name:"ReturnMessage"`

	// 购买的实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GenerateDealsAndPayNewResponse struct {
	*tchttp.BaseResponse
	Response *GenerateDealsAndPayNewResponseParams `json:"Response"`
}

func (r *GenerateDealsAndPayNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateDealsAndPayNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAttackDownloadRecordsRequestParams struct {

}

type GetAttackDownloadRecordsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetAttackDownloadRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAttackDownloadRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAttackDownloadRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAttackDownloadRecordsResponseParams struct {
	// 下载攻击日志记录数组
	Records []*DownloadAttackRecordInfo `json:"Records,omitnil" name:"Records"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetAttackDownloadRecordsResponse struct {
	*tchttp.BaseResponse
	Response *GetAttackDownloadRecordsResponseParams `json:"Response"`
}

func (r *GetAttackDownloadRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAttackDownloadRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAttackHistogramRequestParams struct {
	// 查询的域名，所有域名使用all
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Lucene语法
	QueryString *string `json:"QueryString,omitnil" name:"QueryString"`
}

type GetAttackHistogramRequest struct {
	*tchttp.BaseRequest
	
	// 查询的域名，所有域名使用all
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Lucene语法
	QueryString *string `json:"QueryString,omitnil" name:"QueryString"`
}

func (r *GetAttackHistogramRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAttackHistogramRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QueryString")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAttackHistogramRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAttackHistogramResponseParams struct {
	// 统计详情
	Data []*LogHistogramInfo `json:"Data,omitnil" name:"Data"`

	// 时间段大小
	Period *uint64 `json:"Period,omitnil" name:"Period"`

	// 统计的条目数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetAttackHistogramResponse struct {
	*tchttp.BaseResponse
	Response *GetAttackHistogramResponseParams `json:"Response"`
}

func (r *GetAttackHistogramResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAttackHistogramResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAttackTotalCountRequestParams struct {
	// 起始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询的域名，全部域名不指定
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 查询条件，默认为""
	QueryString *string `json:"QueryString,omitnil" name:"QueryString"`
}

type GetAttackTotalCountRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询的域名，全部域名不指定
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 查询条件，默认为""
	QueryString *string `json:"QueryString,omitnil" name:"QueryString"`
}

func (r *GetAttackTotalCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAttackTotalCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Domain")
	delete(f, "QueryString")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAttackTotalCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAttackTotalCountResponseParams struct {
	// 攻击总次数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetAttackTotalCountResponse struct {
	*tchttp.BaseResponse
	Response *GetAttackTotalCountResponseParams `json:"Response"`
}

func (r *GetAttackTotalCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAttackTotalCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetInstanceQpsLimitRequestParams struct {
	// 套餐实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 套餐类型
	Type *string `json:"Type,omitnil" name:"Type"`
}

type GetInstanceQpsLimitRequest struct {
	*tchttp.BaseRequest
	
	// 套餐实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 套餐类型
	Type *string `json:"Type,omitnil" name:"Type"`
}

func (r *GetInstanceQpsLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetInstanceQpsLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetInstanceQpsLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetInstanceQpsLimitResponseParams struct {
	// 弹性qps相关值集合
	QpsData *QpsData `json:"QpsData,omitnil" name:"QpsData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetInstanceQpsLimitResponse struct {
	*tchttp.BaseResponse
	Response *GetInstanceQpsLimitResponseParams `json:"Response"`
}

func (r *GetInstanceQpsLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetInstanceQpsLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GoodNews struct {
	// 商品数量
	GoodsNum *int64 `json:"GoodsNum,omitnil" name:"GoodsNum"`

	// 商品明细
	GoodsDetail *GoodsDetailNew `json:"GoodsDetail,omitnil" name:"GoodsDetail"`

	// 订单类型ID，用来唯一标识一个业务的一种场景（总共三种场景：新购、配置变更、续费）
	// 高级版: 102375(新购),102376(续费),102377(变配)
	// 企业版 : 102378(新购),102379(续费),102380(变配)
	// 旗舰版 : 102369(新购),102370(续费),102371(变配)
	// 域名包 : 102372(新购),102373(续费),102374(变配)
	// 业务扩展包 : 101040(新购),101041(续费),101042(变配)
	// 
	// 高级版-CLB: 新购 101198  续费 101199 变配 101200
	// 企业版-CLB 101204(新购),101205(续费),101206(变配)
	// 旗舰版-CLB : 101201(新购),101202(续费),101203(变配)
	// 域名包-CLB: 101207(新购),101208(续费),101209(变配)
	// 业务扩展包-CLB: 101210(新购),101211(续费),101212(变配)
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodsCategoryId *int64 `json:"GoodsCategoryId,omitnil" name:"GoodsCategoryId"`

	// 购买waf实例区域ID
	// 1 表示购买大陆资源;
	// 9表示购买非中国大陆资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitnil" name:"RegionId"`
}

type Goods struct {
	// 付费类型，1:预付费，0:后付费
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 商品数量
	GoodsNum *int64 `json:"GoodsNum,omitnil" name:"GoodsNum"`

	// 商品明细
	GoodsDetail *GoodsDetail `json:"GoodsDetail,omitnil" name:"GoodsDetail"`

	// 默认为0
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 计费类目ID，对应cid
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodsCategoryId *int64 `json:"GoodsCategoryId,omitnil" name:"GoodsCategoryId"`

	// 平台类型，默认1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`

	// 购买waf实例区域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitnil" name:"RegionId"`
}

type GoodsDetail struct {
	// 时间间隔
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 单位，支持m、y、d
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 产品码
	ProductCode *string `json:"ProductCode,omitnil" name:"ProductCode"`

	// 二级产品码
	SubProductCode *string `json:"SubProductCode,omitnil" name:"SubProductCode"`

	// 计费策略id
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`

	// waf产品码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductInfo []*ProductInfo `json:"ProductInfo,omitnil" name:"ProductInfo"`

	// waf实例名
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// QPS数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ElasticQps *int64 `json:"ElasticQps,omitnil" name:"ElasticQps"`

	// 弹性账单
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlexBill *int64 `json:"FlexBill,omitnil" name:"FlexBill"`

	// 1:自动续费，0:不自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// waf购买的实际地域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealRegion *int64 `json:"RealRegion,omitnil" name:"RealRegion"`

	// Waf实例对应的二级产品码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 计费细项标签数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelTypes []*string `json:"LabelTypes,omitnil" name:"LabelTypes"`

	// 计费细项标签数量，一般和SvLabelType一一对应
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelCounts []*int64 `json:"LabelCounts,omitnil" name:"LabelCounts"`

	// 变配使用，实例到期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurDeadline *string `json:"CurDeadline,omitnil" name:"CurDeadline"`

	// 对存在的实例购买bot 或api 安全
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type GoodsDetailNew struct {
	// 时间间隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 单位，支持购买d、m、y 即（日、月、年）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 子产品标签,。新购，续费必传，变配时放在oldConfig newConfig里面
	// 
	// Saas 高级版 ：sp_wsm_waf_premium
	// Saas企业版 ：sp_wsm_waf_enterprise
	// Saas旗舰版 ：sp_wsm_waf_ultimate
	// Saas 业务扩展包：sp_wsm_waf_qpsep
	// Saas 域名扩展包：sp_wsm_waf_domain
	// 
	// 高级版-CLB:sp_wsm_waf_premium_clb
	// 企业版-CLB : sp_wsm_waf_enterprise_clb
	// 旗舰版-CLB:sp_wsm_waf_ultimate_clb
	//  业务扩展包-CLB：sp_wsm_waf_qpsep_clb
	// 域名扩展包-CLB：sp_wsm_waf_domain_clb
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductCode *string `json:"SubProductCode,omitnil" name:"SubProductCode"`

	// 业务产品申请的pid（对应一个定价公式），通过pid计费查询到定价模型
	// 高级版 ：1000827
	// 企业版 ：1000830
	// 旗舰版 ：1000832
	// 域名包 : 1000834
	// 业务扩展包 : 1000481
	// 高级版-CLB:1001150
	// 企业版-CLB : 1001152
	// 旗舰版-CLB:1001154
	// 域名包-CLB: 1001156
	// 业务扩展包-CLB : 1001160
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pid *int64 `json:"Pid,omitnil" name:"Pid"`

	// waf实例名
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 1:自动续费，0:不自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// waf购买的实际地域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealRegion *int64 `json:"RealRegion,omitnil" name:"RealRegion"`

	// 计费细项标签数组
	// Saas 高级版  sv_wsm_waf_package_premium 
	// Saas 企业版  sv_wsm_waf_package_enterprise
	// Saas 旗舰版  sv_wsm_waf_package_ultimate 
	// Saas 非中国大陆高级版  sv_wsm_waf_package_premium_intl
	// Saas 非中国大陆企业版   sv_wsm_waf_package_enterprise_intl
	// Saas 非中国大陆旗舰版  sv_wsm_waf_package_ultimate _intl
	// Saas 业务扩展包  sv_wsm_waf_qps_ep
	// Saas 域名扩展包  sv_wsm_waf_domain
	// 
	// 高级版CLB   sv_wsm_waf_package_premium_clb
	// 企业版CLB  sv_wsm_waf_package_enterprise_clb
	// 旗舰版CLB   sv_wsm_waf_package_ultimate_clb
	// 非中国大陆高级版 CLB sv_wsm_waf_package_premium_clb_intl
	// 非中国大陆企业版CLB   sv_wsm_waf_package_premium_clb_intl
	// 非中国大陆旗舰版CLB  sv_wsm_waf_package_ultimate_clb _intl
	// 业务扩展包CLB sv_wsm_waf_qps_ep_clb
	// 域名扩展包CLB  sv_wsm_waf_domain_clb
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelTypes []*string `json:"LabelTypes,omitnil" name:"LabelTypes"`

	// 计费细项标签数量，一般和SvLabelType一一对应
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelCounts []*int64 `json:"LabelCounts,omitnil" name:"LabelCounts"`

	// 变配使用，实例到期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurDeadline *string `json:"CurDeadline,omitnil" name:"CurDeadline"`

	// 对存在的实例购买bot 或api 安全
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`
}

type HostDel struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 实例类型
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

type HostRecord struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 主域名，入参时为空
	MainDomain *string `json:"MainDomain,omitnil" name:"MainDomain"`

	// 规则引擎防护模式，0 观察模式，1拦截模式
	Mode *uint64 `json:"Mode,omitnil" name:"Mode"`

	// waf和LD的绑定，0：没有绑定，1：绑定
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 域名状态，0：正常，1：未检测到流量，2：即将过期，3：过期
	State *uint64 `json:"State,omitnil" name:"State"`

	// 规则引擎和AI引擎防护模式联合状态,10规则引擎观察&&AI引擎关闭模式 11规则引擎观察&&AI引擎观察模式 12规则引擎观察&&AI引擎拦截模式 20规则引擎拦截&&AI引擎关闭模式 21规则引擎拦截&&AI引擎观察模式 22规则引擎拦截&&AI引擎拦截模式
	Engine *uint64 `json:"Engine,omitnil" name:"Engine"`

	// 是否开启代理，0：不开启，1：开启
	IsCdn *uint64 `json:"IsCdn,omitnil" name:"IsCdn"`

	// 绑定的LB列表
	LoadBalancerSet []*LoadBalancer `json:"LoadBalancerSet,omitnil" name:"LoadBalancerSet"`

	// 域名绑定的LB的地域，以,分割多个地域
	Region *string `json:"Region,omitnil" name:"Region"`

	// 产品分类，取值为：sparta-waf、clb-waf、cdn-waf
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// WAF的流量模式，1：清洗模式，0：镜像模式
	FlowMode *uint64 `json:"FlowMode,omitnil" name:"FlowMode"`

	// 是否开启访问日志，1：开启，0：关闭
	ClsStatus *uint64 `json:"ClsStatus,omitnil" name:"ClsStatus"`

	// 防护等级，可选值100,200,300
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *uint64 `json:"Level,omitnil" name:"Level"`

	// 域名需要下发到的cdc集群列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdcClusters []*string `json:"CdcClusters,omitnil" name:"CdcClusters"`

	// 应用型负载均衡类型: clb或者apisix，默认clb
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlbType *string `json:"AlbType,omitnil" name:"AlbType"`

	// IsCdn=3时，需要填此参数，表示自定义header
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpHeaders []*string `json:"IpHeaders,omitnil" name:"IpHeaders"`

	// 规则引擎类型， 1: menshen,   2:tiga
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineType *int64 `json:"EngineType,omitnil" name:"EngineType"`

	// 云类型:public:公有云；private:私有云;hybrid:混合云
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloudType *string `json:"CloudType,omitnil" name:"CloudType"`
}

type HostStatus struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// WAF的开关，1：开，0：关
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

type HybridPkg struct {
	// 资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceIds *string `json:"ResourceIds,omitnil" name:"ResourceIds"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *int64 `json:"Region,omitnil" name:"Region"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 申请数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InquireNum *int64 `json:"InquireNum,omitnil" name:"InquireNum"`

	// 使用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedNum *int64 `json:"UsedNum,omitnil" name:"UsedNum"`

	// 续费标志
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *uint64 `json:"RenewFlag,omitnil" name:"RenewFlag"`
}

type InstanceInfo struct {
	// id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 资源id
	ResourceIds *string `json:"ResourceIds,omitnil" name:"ResourceIds"`

	// 地域
	Region *string `json:"Region,omitnil" name:"Region"`

	// 付费模式
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// 自动续费
	RenewFlag *uint64 `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// 弹性计费
	Mode *uint64 `json:"Mode,omitnil" name:"Mode"`

	// 套餐版本
	Level *uint64 `json:"Level,omitnil" name:"Level"`

	// 过期时间
	ValidTime *string `json:"ValidTime,omitnil" name:"ValidTime"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 已用
	DomainCount *uint64 `json:"DomainCount,omitnil" name:"DomainCount"`

	// 上限
	SubDomainLimit *uint64 `json:"SubDomainLimit,omitnil" name:"SubDomainLimit"`

	// 已用
	MainDomainCount *uint64 `json:"MainDomainCount,omitnil" name:"MainDomainCount"`

	// 上限
	MainDomainLimit *uint64 `json:"MainDomainLimit,omitnil" name:"MainDomainLimit"`

	// 峰值
	MaxQPS *uint64 `json:"MaxQPS,omitnil" name:"MaxQPS"`

	// qps套餐
	QPS *QPSPackageNew `json:"QPS,omitnil" name:"QPS"`

	// 域名套餐
	DomainPkg *DomainPackageNew `json:"DomainPkg,omitnil" name:"DomainPkg"`

	// 用户appid
	AppId *uint64 `json:"AppId,omitnil" name:"AppId"`

	// clb或saas
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 业务安全包
	// 注意：此字段可能返回 null，表示取不到有效值。
	FraudPkg *FraudPkg `json:"FraudPkg,omitnil" name:"FraudPkg"`

	// Bot资源包
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotPkg *BotPkg `json:"BotPkg,omitnil" name:"BotPkg"`

	// bot的qps详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	BotQPS *BotQPS `json:"BotQPS,omitnil" name:"BotQPS"`

	// qps弹性计费上限
	// 注意：此字段可能返回 null，表示取不到有效值。
	ElasticBilling *uint64 `json:"ElasticBilling,omitnil" name:"ElasticBilling"`

	// 攻击日志投递开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackLogPost *int64 `json:"AttackLogPost,omitnil" name:"AttackLogPost"`

	// 带宽峰值，单位为B/s(字节每秒)
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxBandwidth *uint64 `json:"MaxBandwidth,omitnil" name:"MaxBandwidth"`

	// api安全是否购买
	APISecurity *uint64 `json:"APISecurity,omitnil" name:"APISecurity"`

	// 购买的qps规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	QpsStandard *uint64 `json:"QpsStandard,omitnil" name:"QpsStandard"`

	// 购买的带宽规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthStandard *uint64 `json:"BandwidthStandard,omitnil" name:"BandwidthStandard"`

	// 实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 实例沙箱值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SandboxQps *uint64 `json:"SandboxQps,omitnil" name:"SandboxQps"`

	// 是否api 安全试用
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAPISecurityTrial *uint64 `json:"IsAPISecurityTrial,omitnil" name:"IsAPISecurityTrial"`

	// 重保包
	// 注意：此字段可能返回 null，表示取不到有效值。
	MajorEventsPkg *MajorEventsPkg `json:"MajorEventsPkg,omitnil" name:"MajorEventsPkg"`

	// 混合云子节点包
	// 注意：此字段可能返回 null，表示取不到有效值。
	HybridPkg *HybridPkg `json:"HybridPkg,omitnil" name:"HybridPkg"`

	// API安全资源包
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiPkg *ApiPkg `json:"ApiPkg,omitnil" name:"ApiPkg"`
}

type IpAccessControlData struct {
	// ip黑白名单
	// 注意：此字段可能返回 null，表示取不到有效值。
	Res []*IpAccessControlItem `json:"Res,omitnil" name:"Res"`

	// 计数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`
}

type IpAccessControlItem struct {
	// mongo表自增Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil" name:"Id"`

	// 动作
	ActionType *uint64 `json:"ActionType,omitnil" name:"ActionType"`

	// ip
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 备注
	Note *string `json:"Note,omitnil" name:"Note"`

	// 来源
	Source *string `json:"Source,omitnil" name:"Source"`

	// 更新时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	TsVersion *uint64 `json:"TsVersion,omitnil" name:"TsVersion"`

	// 有效截止时间戳
	ValidTs *uint64 `json:"ValidTs,omitnil" name:"ValidTs"`

	// 生效状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidStatus *int64 `json:"ValidStatus,omitnil" name:"ValidStatus"`
}

type IpHitItem struct {
	// 动作
	Action *uint64 `json:"Action,omitnil" name:"Action"`

	// 类别
	Category *string `json:"Category,omitnil" name:"Category"`

	// ip
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 规则名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 时间戳
	TsVersion *uint64 `json:"TsVersion,omitnil" name:"TsVersion"`

	// 有效截止时间戳
	ValidTs *uint64 `json:"ValidTs,omitnil" name:"ValidTs"`
}

type IpHitItemsData struct {
	// 数组封装
	Res []*IpHitItem `json:"Res,omitnil" name:"Res"`

	// 总数目
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`
}

type KVInt struct {
	// Key
	Key *string `json:"Key,omitnil" name:"Key"`

	// Value
	Value *uint64 `json:"Value,omitnil" name:"Value"`
}

type LoadBalancer struct {
	// 负载均衡LD的ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil" name:"LoadBalancerId"`

	// 负载均衡LD的名称
	LoadBalancerName *string `json:"LoadBalancerName,omitnil" name:"LoadBalancerName"`

	// 负载均衡监听器的ID
	ListenerId *string `json:"ListenerId,omitnil" name:"ListenerId"`

	// 负载均衡监听器的名称
	ListenerName *string `json:"ListenerName,omitnil" name:"ListenerName"`

	// 负载均衡实例的IP
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// 负载均衡实例的端口
	Vport *uint64 `json:"Vport,omitnil" name:"Vport"`

	// 负载均衡LD的地域
	Region *string `json:"Region,omitnil" name:"Region"`

	// 监听器协议，http、https
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 负载均衡监听器所在的zone
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 负载均衡的VPCID，公网为-1，内网按实际填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumericalVpcId *int64 `json:"NumericalVpcId,omitnil" name:"NumericalVpcId"`

	// 负载均衡的网络类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerType *string `json:"LoadBalancerType,omitnil" name:"LoadBalancerType"`

	// 负载均衡的域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerDomain *string `json:"LoadBalancerDomain,omitnil" name:"LoadBalancerDomain"`
}

type LoadBalancerPackageNew struct {
	// 监听id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerId *string `json:"ListenerId,omitnil" name:"ListenerId"`

	// 监听名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerName *string `json:"ListenerName,omitnil" name:"ListenerName"`

	// 负载均衡id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil" name:"LoadBalancerId"`

	// 负载均衡名
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerName *string `json:"LoadBalancerName,omitnil" name:"LoadBalancerName"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 地区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 接入IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// 接入端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vport *uint64 `json:"Vport,omitnil" name:"Vport"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// VPCID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumericalVpcId *int64 `json:"NumericalVpcId,omitnil" name:"NumericalVpcId"`

	// CLB类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerType *string `json:"LoadBalancerType,omitnil" name:"LoadBalancerType"`

	// 负载均衡器的域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerDomain *string `json:"LoadBalancerDomain,omitnil" name:"LoadBalancerDomain"`
}

type LogHistogramInfo struct {
	// 日志条数
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 时间戳
	TimeStamp *int64 `json:"TimeStamp,omitnil" name:"TimeStamp"`
}

type MajorEventsPkg struct {
	// 资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceIds *string `json:"ResourceIds,omitnil" name:"ResourceIds"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *int64 `json:"Region,omitnil" name:"Region"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 申请数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InquireNum *int64 `json:"InquireNum,omitnil" name:"InquireNum"`

	// 使用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedNum *int64 `json:"UsedNum,omitnil" name:"UsedNum"`

	// 续费标志
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *uint64 `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// 计费项
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingItem *string `json:"BillingItem,omitnil" name:"BillingItem"`

	// 护网包状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	HWState *int64 `json:"HWState,omitnil" name:"HWState"`
}

// Predefined struct for user
type ModifyAccessPeriodRequestParams struct {
	// 访问日志保存期限，范围为[1, 180]
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// 日志主题，新版本不需要再传
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`
}

type ModifyAccessPeriodRequest struct {
	*tchttp.BaseRequest
	
	// 访问日志保存期限，范围为[1, 180]
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// 日志主题，新版本不需要再传
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`
}

func (r *ModifyAccessPeriodRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessPeriodRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Period")
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAccessPeriodRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAccessPeriodResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAccessPeriodResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAccessPeriodResponseParams `json:"Response"`
}

func (r *ModifyAccessPeriodResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAccessPeriodResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAntiFakeUrlRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// uri
	Uri *string `json:"Uri,omitnil" name:"Uri"`

	// ID
	Id *int64 `json:"Id,omitnil" name:"Id"`
}

type ModifyAntiFakeUrlRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// uri
	Uri *string `json:"Uri,omitnil" name:"Uri"`

	// ID
	Id *int64 `json:"Id,omitnil" name:"Id"`
}

func (r *ModifyAntiFakeUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAntiFakeUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Name")
	delete(f, "Uri")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAntiFakeUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAntiFakeUrlResponseParams struct {
	// 结果
	Result *string `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAntiFakeUrlResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAntiFakeUrlResponseParams `json:"Response"`
}

func (r *ModifyAntiFakeUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAntiFakeUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAntiFakeUrlStatusRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// Id列表
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

type ModifyAntiFakeUrlStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// Id列表
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`
}

func (r *ModifyAntiFakeUrlStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAntiFakeUrlStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Status")
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAntiFakeUrlStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAntiFakeUrlStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAntiFakeUrlStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAntiFakeUrlStatusResponseParams `json:"Response"`
}

func (r *ModifyAntiFakeUrlStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAntiFakeUrlStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAntiInfoLeakRuleStatusRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 规则
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

type ModifyAntiInfoLeakRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 规则
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

func (r *ModifyAntiInfoLeakRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAntiInfoLeakRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RuleId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAntiInfoLeakRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAntiInfoLeakRuleStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAntiInfoLeakRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAntiInfoLeakRuleStatusResponseParams `json:"Response"`
}

func (r *ModifyAntiInfoLeakRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAntiInfoLeakRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAntiInfoLeakRulesRequestParams struct {
	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 规则名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Action 值
	ActionType *uint64 `json:"ActionType,omitnil" name:"ActionType"`

	// 策略数组
	Strategies []*StrategyForAntiInfoLeak `json:"Strategies,omitnil" name:"Strategies"`
}

type ModifyAntiInfoLeakRulesRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 规则名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// Action 值
	ActionType *uint64 `json:"ActionType,omitnil" name:"ActionType"`

	// 策略数组
	Strategies []*StrategyForAntiInfoLeak `json:"Strategies,omitnil" name:"Strategies"`
}

func (r *ModifyAntiInfoLeakRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAntiInfoLeakRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "Name")
	delete(f, "Domain")
	delete(f, "ActionType")
	delete(f, "Strategies")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAntiInfoLeakRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAntiInfoLeakRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAntiInfoLeakRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAntiInfoLeakRulesResponseParams `json:"Response"`
}

func (r *ModifyAntiInfoLeakRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAntiInfoLeakRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApiAnalyzeStatusRequestParams struct {
	// 开关状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 需要批量开启的实体列表
	TargetList []*TargetEntity `json:"TargetList,omitnil" name:"TargetList"`
}

type ModifyApiAnalyzeStatusRequest struct {
	*tchttp.BaseRequest
	
	// 开关状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 需要批量开启的实体列表
	TargetList []*TargetEntity `json:"TargetList,omitnil" name:"TargetList"`
}

func (r *ModifyApiAnalyzeStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiAnalyzeStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "Domain")
	delete(f, "InstanceId")
	delete(f, "TargetList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApiAnalyzeStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApiAnalyzeStatusResponseParams struct {
	// 已经开启的数量,如果返回值为3（大于支持的域名开启数量），则表示开启失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 不支持开启的域名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnSupportedList []*string `json:"UnSupportedList,omitnil" name:"UnSupportedList"`

	// 开启/关闭失败的域名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailDomainList []*string `json:"FailDomainList,omitnil" name:"FailDomainList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyApiAnalyzeStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApiAnalyzeStatusResponseParams `json:"Response"`
}

func (r *ModifyApiAnalyzeStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiAnalyzeStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAreaBanStatusRequestParams struct {
	// 需要修改的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 状态值，0表示关闭，1表示开启
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type ModifyAreaBanStatusRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 状态值，0表示关闭，1表示开启
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

func (r *ModifyAreaBanStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAreaBanStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAreaBanStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAreaBanStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAreaBanStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAreaBanStatusResponseParams `json:"Response"`
}

func (r *ModifyAreaBanStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAreaBanStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAttackWhiteRuleRequestParams struct {
	// 规则序号
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 规则Id
	SignatureId *string `json:"SignatureId,omitnil" name:"SignatureId"`

	// 规则状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 匹配规则项列表
	Rules []*UserWhiteRuleItem `json:"Rules,omitnil" name:"Rules"`
}

type ModifyAttackWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则序号
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 规则Id
	SignatureId *string `json:"SignatureId,omitnil" name:"SignatureId"`

	// 规则状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 匹配规则项列表
	Rules []*UserWhiteRuleItem `json:"Rules,omitnil" name:"Rules"`
}

func (r *ModifyAttackWhiteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAttackWhiteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "Domain")
	delete(f, "SignatureId")
	delete(f, "Status")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAttackWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAttackWhiteRuleResponseParams struct {
	// 规则总数
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAttackWhiteRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAttackWhiteRuleResponseParams `json:"Response"`
}

func (r *ModifyAttackWhiteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAttackWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBotStatusRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 类别
	Category *string `json:"Category,omitnil" name:"Category"`

	// 状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 是否是bot4.0版本
	IsVersionFour *bool `json:"IsVersionFour,omitnil" name:"IsVersionFour"`

	// 传入Bot版本号，场景化版本为"4.1.0"
	BotVersion *string `json:"BotVersion,omitnil" name:"BotVersion"`
}

type ModifyBotStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 类别
	Category *string `json:"Category,omitnil" name:"Category"`

	// 状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 是否是bot4.0版本
	IsVersionFour *bool `json:"IsVersionFour,omitnil" name:"IsVersionFour"`

	// 传入Bot版本号，场景化版本为"4.1.0"
	BotVersion *string `json:"BotVersion,omitnil" name:"BotVersion"`
}

func (r *ModifyBotStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBotStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Category")
	delete(f, "Status")
	delete(f, "InstanceID")
	delete(f, "IsVersionFour")
	delete(f, "BotVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBotStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBotStatusResponseParams struct {
	// 正常情况为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyBotStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBotStatusResponseParams `json:"Response"`
}

func (r *ModifyBotStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBotStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomRuleRequestParams struct {
	// 编辑的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 编辑的规则ID
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 编辑的规则名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 执行动作，0：放行、1：阻断、2：人机识别、3：观察、4：重定向
	RuleAction *string `json:"RuleAction,omitnil" name:"RuleAction"`

	// 匹配条件数组
	Strategies []*Strategy `json:"Strategies,omitnil" name:"Strategies"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 动作为重定向的时候重定向URL，默认为"/"
	Redirect *string `json:"Redirect,omitnil" name:"Redirect"`

	// 放行时是否继续执行其它检查逻辑，继续执行地域封禁防护：geoip、继续执行CC策略防护：cc、继续执行WEB应用防护：owasp、继续执行AI引擎防护：ai、继续执行信息防泄漏防护：antileakage。如果多个勾选那么以,串接。
	// 默认是"geoip,cc,owasp,ai,antileakage"
	Bypass *string `json:"Bypass,omitnil" name:"Bypass"`

	// 优先级，1~100的整数，数字越小，代表这条规则的执行优先级越高。
	// 默认是100
	SortId *uint64 `json:"SortId,omitnil" name:"SortId"`

	// 规则生效截止时间，0：永久生效，其它值为对应时间的时间戳。
	// 默认是0
	ExpireTime *uint64 `json:"ExpireTime,omitnil" name:"ExpireTime"`
}

type ModifyCustomRuleRequest struct {
	*tchttp.BaseRequest
	
	// 编辑的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 编辑的规则ID
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 编辑的规则名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 执行动作，0：放行、1：阻断、2：人机识别、3：观察、4：重定向
	RuleAction *string `json:"RuleAction,omitnil" name:"RuleAction"`

	// 匹配条件数组
	Strategies []*Strategy `json:"Strategies,omitnil" name:"Strategies"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 动作为重定向的时候重定向URL，默认为"/"
	Redirect *string `json:"Redirect,omitnil" name:"Redirect"`

	// 放行时是否继续执行其它检查逻辑，继续执行地域封禁防护：geoip、继续执行CC策略防护：cc、继续执行WEB应用防护：owasp、继续执行AI引擎防护：ai、继续执行信息防泄漏防护：antileakage。如果多个勾选那么以,串接。
	// 默认是"geoip,cc,owasp,ai,antileakage"
	Bypass *string `json:"Bypass,omitnil" name:"Bypass"`

	// 优先级，1~100的整数，数字越小，代表这条规则的执行优先级越高。
	// 默认是100
	SortId *uint64 `json:"SortId,omitnil" name:"SortId"`

	// 规则生效截止时间，0：永久生效，其它值为对应时间的时间戳。
	// 默认是0
	ExpireTime *uint64 `json:"ExpireTime,omitnil" name:"ExpireTime"`
}

func (r *ModifyCustomRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RuleId")
	delete(f, "RuleName")
	delete(f, "RuleAction")
	delete(f, "Strategies")
	delete(f, "Edition")
	delete(f, "Redirect")
	delete(f, "Bypass")
	delete(f, "SortId")
	delete(f, "ExpireTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCustomRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomRuleResponseParams `json:"Response"`
}

func (r *ModifyCustomRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomRuleStatusRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 开关的状态，1是开启、0是关闭
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil" name:"Edition"`
}

type ModifyCustomRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 开关的状态，1是开启、0是关闭
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil" name:"Edition"`
}

func (r *ModifyCustomRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RuleId")
	delete(f, "Status")
	delete(f, "Edition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomRuleStatusResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	Success *ResponseCode `json:"Success,omitnil" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCustomRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomRuleStatusResponseParams `json:"Response"`
}

func (r *ModifyCustomRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomWhiteRuleRequestParams struct {
	// 编辑的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 编辑的规则ID
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 编辑的规则名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 放行时是否继续执行其它检查逻辑，继续执行地域封禁防护：geoip、继续执行CC策略防护：cc、继续执行WEB应用防护：owasp、继续执行AI引擎防护：ai、继续执行信息防泄漏防护：antileakage。如果勾选多个，则以“，”串接。
	Bypass *string `json:"Bypass,omitnil" name:"Bypass"`

	// 优先级，1~100的整数，数字越小，代表这条规则的执行优先级越高。
	SortId *uint64 `json:"SortId,omitnil" name:"SortId"`

	// 规则生效截止时间，0：永久生效，其它值为对应时间的时间戳。
	ExpireTime *uint64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 匹配条件数组
	Strategies []*Strategy `json:"Strategies,omitnil" name:"Strategies"`
}

type ModifyCustomWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 编辑的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 编辑的规则ID
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 编辑的规则名称
	RuleName *string `json:"RuleName,omitnil" name:"RuleName"`

	// 放行时是否继续执行其它检查逻辑，继续执行地域封禁防护：geoip、继续执行CC策略防护：cc、继续执行WEB应用防护：owasp、继续执行AI引擎防护：ai、继续执行信息防泄漏防护：antileakage。如果勾选多个，则以“，”串接。
	Bypass *string `json:"Bypass,omitnil" name:"Bypass"`

	// 优先级，1~100的整数，数字越小，代表这条规则的执行优先级越高。
	SortId *uint64 `json:"SortId,omitnil" name:"SortId"`

	// 规则生效截止时间，0：永久生效，其它值为对应时间的时间戳。
	ExpireTime *uint64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 匹配条件数组
	Strategies []*Strategy `json:"Strategies,omitnil" name:"Strategies"`
}

func (r *ModifyCustomWhiteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomWhiteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RuleId")
	delete(f, "RuleName")
	delete(f, "Bypass")
	delete(f, "SortId")
	delete(f, "ExpireTime")
	delete(f, "Strategies")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomWhiteRuleResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	Success *ResponseCode `json:"Success,omitnil" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCustomWhiteRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomWhiteRuleResponseParams `json:"Response"`
}

func (r *ModifyCustomWhiteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomWhiteRuleStatusRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 开关的状态，1是开启、0是关闭
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

type ModifyCustomWhiteRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil" name:"RuleId"`

	// 开关的状态，1是开启、0是关闭
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

func (r *ModifyCustomWhiteRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomWhiteRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RuleId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomWhiteRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomWhiteRuleStatusResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	Success *ResponseCode `json:"Success,omitnil" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCustomWhiteRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomWhiteRuleStatusResponseParams `json:"Response"`
}

func (r *ModifyCustomWhiteRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomWhiteRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainIpv6StatusRequestParams struct {
	// 需要修改的域名所属的实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 需要修改的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 需要修改的域名ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 修改域名的Ipv6开关为Status （1:开启 2:关闭）
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type ModifyDomainIpv6StatusRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的域名所属的实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 需要修改的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 需要修改的域名ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 修改域名的Ipv6开关为Status （1:开启 2:关闭）
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

func (r *ModifyDomainIpv6StatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainIpv6StatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainIpv6StatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainIpv6StatusResponseParams struct {
	// 返回的状态 （0: 操作失败 1:操作成功 2:企业版以上不支持 3:企业版以下不支持 ）
	Ipv6Status *int64 `json:"Ipv6Status,omitnil" name:"Ipv6Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDomainIpv6StatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainIpv6StatusResponseParams `json:"Response"`
}

func (r *ModifyDomainIpv6StatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainIpv6StatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainWhiteRuleRequestParams struct {
	// 需要更改的规则的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 白名单id
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 规则的id列表
	Rules []*uint64 `json:"Rules,omitnil" name:"Rules"`

	// 规则匹配路径
	Url *string `json:"Url,omitnil" name:"Url"`

	// 规则匹配方法
	Function *string `json:"Function,omitnil" name:"Function"`

	// 规则的开关状态，0表示关闭开关，1表示打开开关
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

type ModifyDomainWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要更改的规则的域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 白名单id
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 规则的id列表
	Rules []*uint64 `json:"Rules,omitnil" name:"Rules"`

	// 规则匹配路径
	Url *string `json:"Url,omitnil" name:"Url"`

	// 规则匹配方法
	Function *string `json:"Function,omitnil" name:"Function"`

	// 规则的开关状态，0表示关闭开关，1表示打开开关
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

func (r *ModifyDomainWhiteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainWhiteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Id")
	delete(f, "Rules")
	delete(f, "Url")
	delete(f, "Function")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainWhiteRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDomainWhiteRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainWhiteRuleResponseParams `json:"Response"`
}

func (r *ModifyDomainWhiteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainWhiteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainsCLSStatusRequestParams struct {
	// 需要修改的域名列表
	Domains []*DomainURI `json:"Domains,omitnil" name:"Domains"`

	// 修改域名的访问日志开关为Status
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

type ModifyDomainsCLSStatusRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的域名列表
	Domains []*DomainURI `json:"Domains,omitnil" name:"Domains"`

	// 修改域名的访问日志开关为Status
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

func (r *ModifyDomainsCLSStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainsCLSStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domains")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainsCLSStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainsCLSStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyDomainsCLSStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainsCLSStatusResponseParams `json:"Response"`
}

func (r *ModifyDomainsCLSStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainsCLSStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGenerateDealsRequestParams struct {
	// 计费下单入参
	Goods []*Goods `json:"Goods,omitnil" name:"Goods"`
}

type ModifyGenerateDealsRequest struct {
	*tchttp.BaseRequest
	
	// 计费下单入参
	Goods []*Goods `json:"Goods,omitnil" name:"Goods"`
}

func (r *ModifyGenerateDealsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGenerateDealsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Goods")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGenerateDealsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGenerateDealsResponseParams struct {
	// 计费下单响应结构体
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DealData `json:"Data,omitnil" name:"Data"`

	// 1:成功，0:失败
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 返回message
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnMessage *string `json:"ReturnMessage,omitnil" name:"ReturnMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyGenerateDealsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGenerateDealsResponseParams `json:"Response"`
}

func (r *ModifyGenerateDealsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGenerateDealsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostFlowModeRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// WAF流量模式，1：清洗模式，0：镜像模式（默认）
	FlowMode *uint64 `json:"FlowMode,omitnil" name:"FlowMode"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

type ModifyHostFlowModeRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// WAF流量模式，1：清洗模式，0：镜像模式（默认）
	FlowMode *uint64 `json:"FlowMode,omitnil" name:"FlowMode"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

func (r *ModifyHostFlowModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostFlowModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "FlowMode")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHostFlowModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostFlowModeResponseParams struct {
	// 成功的状态码
	Success *ResponseCode `json:"Success,omitnil" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyHostFlowModeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHostFlowModeResponseParams `json:"Response"`
}

func (r *ModifyHostFlowModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostFlowModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostModeRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 防护状态：
	// 10：规则观察&&AI关闭模式，11：规则观察&&AI观察模式，12：规则观察&&AI拦截模式
	// 20：规则拦截&&AI关闭模式，21：规则拦截&&AI观察模式，22：规则拦截&&AI拦截模式
	Mode *uint64 `json:"Mode,omitnil" name:"Mode"`

	// 0:修改防护模式，1:修改AI
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 实例类型
	Edition *string `json:"Edition,omitnil" name:"Edition"`
}

type ModifyHostModeRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 防护状态：
	// 10：规则观察&&AI关闭模式，11：规则观察&&AI观察模式，12：规则观察&&AI拦截模式
	// 20：规则拦截&&AI关闭模式，21：规则拦截&&AI观察模式，22：规则拦截&&AI拦截模式
	Mode *uint64 `json:"Mode,omitnil" name:"Mode"`

	// 0:修改防护模式，1:修改AI
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 实例类型
	Edition *string `json:"Edition,omitnil" name:"Edition"`
}

func (r *ModifyHostModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "Mode")
	delete(f, "Type")
	delete(f, "InstanceID")
	delete(f, "Edition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHostModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostModeResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Success *ResponseCode `json:"Success,omitnil" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyHostModeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHostModeResponseParams `json:"Response"`
}

func (r *ModifyHostModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostRequestParams struct {
	// 编辑的域名配置信息
	Host *HostRecord `json:"Host,omitnil" name:"Host"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

type ModifyHostRequest struct {
	*tchttp.BaseRequest
	
	// 编辑的域名配置信息
	Host *HostRecord `json:"Host,omitnil" name:"Host"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

func (r *ModifyHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Host")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostResponseParams struct {
	// 编辑的域名ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyHostResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHostResponseParams `json:"Response"`
}

func (r *ModifyHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostStatusRequestParams struct {
	// 域名状态列表
	HostsStatus []*HostStatus `json:"HostsStatus,omitnil" name:"HostsStatus"`
}

type ModifyHostStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名状态列表
	HostsStatus []*HostStatus `json:"HostsStatus,omitnil" name:"HostsStatus"`
}

func (r *ModifyHostStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HostsStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHostStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHostStatusResponseParams struct {
	// 成功的状态码，需要JSON解码后再使用，返回的格式是{"域名":"状态"}，成功的状态码为Success，其它的为失败的状态码（yunapi定义的错误码）
	Success *ResponseCode `json:"Success,omitnil" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyHostStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHostStatusResponseParams `json:"Response"`
}

func (r *ModifyHostStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHostStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceElasticModeRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 弹性计费开关
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`
}

type ModifyInstanceElasticModeRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 弹性计费开关
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`
}

func (r *ModifyInstanceElasticModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceElasticModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceElasticModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceElasticModeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstanceElasticModeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceElasticModeResponseParams `json:"Response"`
}

func (r *ModifyInstanceElasticModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceElasticModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceNameRequestParams struct {
	// 新名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 版本
	Edition *string `json:"Edition,omitnil" name:"Edition"`
}

type ModifyInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// 新名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 版本
	Edition *string `json:"Edition,omitnil" name:"Edition"`
}

func (r *ModifyInstanceNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "InstanceID")
	delete(f, "Edition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceNameResponseParams struct {
	// 修改状态：0为成功
	ModifyCode *int64 `json:"ModifyCode,omitnil" name:"ModifyCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstanceNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceNameResponseParams `json:"Response"`
}

func (r *ModifyInstanceNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceQpsLimitRequestParams struct {
	// 套餐实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// qps上限
	QpsLimit *int64 `json:"QpsLimit,omitnil" name:"QpsLimit"`
}

type ModifyInstanceQpsLimitRequest struct {
	*tchttp.BaseRequest
	
	// 套餐实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// qps上限
	QpsLimit *int64 `json:"QpsLimit,omitnil" name:"QpsLimit"`
}

func (r *ModifyInstanceQpsLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceQpsLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "QpsLimit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceQpsLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceQpsLimitResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstanceQpsLimitResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceQpsLimitResponseParams `json:"Response"`
}

func (r *ModifyInstanceQpsLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceQpsLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceRenewFlagRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 续费开关
	RenewFlag *int64 `json:"RenewFlag,omitnil" name:"RenewFlag"`
}

type ModifyInstanceRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 续费开关
	RenewFlag *int64 `json:"RenewFlag,omitnil" name:"RenewFlag"`
}

func (r *ModifyInstanceRenewFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceRenewFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceRenewFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceRenewFlagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyInstanceRenewFlagResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceRenewFlagResponseParams `json:"Response"`
}

func (r *ModifyInstanceRenewFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceRenewFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleStatusRequestParams struct {
	// 需要设置的domain
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// WEB 安全模块开关，0或1
	WebSecurity *uint64 `json:"WebSecurity,omitnil" name:"WebSecurity"`

	// 访问控制模块开关，0或者1
	AccessControl *uint64 `json:"AccessControl,omitnil" name:"AccessControl"`

	// CC模块开关，0或者1
	CcProtection *uint64 `json:"CcProtection,omitnil" name:"CcProtection"`

	// API安全模块开关，0或者1
	ApiProtection *uint64 `json:"ApiProtection,omitnil" name:"ApiProtection"`

	// 防篡改模块开关，0或者1
	AntiTamper *uint64 `json:"AntiTamper,omitnil" name:"AntiTamper"`

	// 防泄漏模块开关，0或者1
	AntiLeakage *uint64 `json:"AntiLeakage,omitnil" name:"AntiLeakage"`

	// 限流模块开关，0或1
	RateLimit *uint64 `json:"RateLimit,omitnil" name:"RateLimit"`
}

type ModifyModuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 需要设置的domain
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// WEB 安全模块开关，0或1
	WebSecurity *uint64 `json:"WebSecurity,omitnil" name:"WebSecurity"`

	// 访问控制模块开关，0或者1
	AccessControl *uint64 `json:"AccessControl,omitnil" name:"AccessControl"`

	// CC模块开关，0或者1
	CcProtection *uint64 `json:"CcProtection,omitnil" name:"CcProtection"`

	// API安全模块开关，0或者1
	ApiProtection *uint64 `json:"ApiProtection,omitnil" name:"ApiProtection"`

	// 防篡改模块开关，0或者1
	AntiTamper *uint64 `json:"AntiTamper,omitnil" name:"AntiTamper"`

	// 防泄漏模块开关，0或者1
	AntiLeakage *uint64 `json:"AntiLeakage,omitnil" name:"AntiLeakage"`

	// 限流模块开关，0或1
	RateLimit *uint64 `json:"RateLimit,omitnil" name:"RateLimit"`
}

func (r *ModifyModuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "WebSecurity")
	delete(f, "AccessControl")
	delete(f, "CcProtection")
	delete(f, "ApiProtection")
	delete(f, "AntiTamper")
	delete(f, "AntiLeakage")
	delete(f, "RateLimit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyModuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModuleStatusResponseParams `json:"Response"`
}

func (r *ModifyModuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyObjectRequestParams struct {
	// 修改对象标识
	ObjectId *string `json:"ObjectId,omitnil" name:"ObjectId"`

	// 改动作类型:Status修改开关，InstanceId绑定实例
	OpType *string `json:"OpType,omitnil" name:"OpType"`

	// 新的Waf开关状态，如果和已有状态相同认为修改成功
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 新的实例ID，如果和已绑定的实例相同认为修改成功
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type ModifyObjectRequest struct {
	*tchttp.BaseRequest
	
	// 修改对象标识
	ObjectId *string `json:"ObjectId,omitnil" name:"ObjectId"`

	// 改动作类型:Status修改开关，InstanceId绑定实例
	OpType *string `json:"OpType,omitnil" name:"OpType"`

	// 新的Waf开关状态，如果和已有状态相同认为修改成功
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 新的实例ID，如果和已绑定的实例相同认为修改成功
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *ModifyObjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyObjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ObjectId")
	delete(f, "OpType")
	delete(f, "Status")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyObjectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyObjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyObjectResponseParams `json:"Response"`
}

func (r *ModifyObjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProtectionStatusRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil" name:"Edition"`
}

type ModifyProtectionStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil" name:"Edition"`
}

func (r *ModifyProtectionStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProtectionStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Status")
	delete(f, "Edition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProtectionStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProtectionStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyProtectionStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProtectionStatusResponseParams `json:"Response"`
}

func (r *ModifyProtectionStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProtectionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySpartaProtectionModeRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 防护状态：
	// 10：规则观察&&AI关闭模式，11：规则观察&&AI观察模式，12：规则观察&&AI拦截模式
	// 20：规则拦截&&AI关闭模式，21：规则拦截&&AI观察模式，22：规则拦截&&AI拦截模式
	Mode *uint64 `json:"Mode,omitnil" name:"Mode"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 0是修改规则引擎状态，1是修改AI的状态
	Type *uint64 `json:"Type,omitnil" name:"Type"`
}

type ModifySpartaProtectionModeRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 防护状态：
	// 10：规则观察&&AI关闭模式，11：规则观察&&AI观察模式，12：规则观察&&AI拦截模式
	// 20：规则拦截&&AI关闭模式，21：规则拦截&&AI观察模式，22：规则拦截&&AI拦截模式
	Mode *uint64 `json:"Mode,omitnil" name:"Mode"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 0是修改规则引擎状态，1是修改AI的状态
	Type *uint64 `json:"Type,omitnil" name:"Type"`
}

func (r *ModifySpartaProtectionModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySpartaProtectionModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Mode")
	delete(f, "Edition")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySpartaProtectionModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySpartaProtectionModeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySpartaProtectionModeResponse struct {
	*tchttp.BaseResponse
	Response *ModifySpartaProtectionModeResponseParams `json:"Response"`
}

func (r *ModifySpartaProtectionModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySpartaProtectionModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySpartaProtectionRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 证书类型，0表示没有证书，CertType=1表示自有证书,2 为托管证书
	CertType *int64 `json:"CertType,omitnil" name:"CertType"`

	// CertType=1时，需要填次参数，表示证书内容
	Cert *string `json:"Cert,omitnil" name:"Cert"`

	// CertType=1时，需要填次参数，表示证书的私钥
	PrivateKey *string `json:"PrivateKey,omitnil" name:"PrivateKey"`

	// CertType=2时，需要填次参数，表示证书的ID
	SSLId *string `json:"SSLId,omitnil" name:"SSLId"`

	// 表示是否开启了CDN代理
	IsCdn *int64 `json:"IsCdn,omitnil" name:"IsCdn"`

	// HTTPS回源协议
	UpstreamScheme *string `json:"UpstreamScheme,omitnil" name:"UpstreamScheme"`

	// HTTPS回源端口,仅UpstreamScheme为http时需要填当前字段
	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitnil" name:"HttpsUpstreamPort"`

	// 表示是否强制跳转到HTTPS，1表示开启，0表示不开启
	HttpsRewrite *uint64 `json:"HttpsRewrite,omitnil" name:"HttpsRewrite"`

	// 回源类型，0表示通过IP回源,1 表示通过域名回源
	UpstreamType *int64 `json:"UpstreamType,omitnil" name:"UpstreamType"`

	// UpstreamType=1时，填次字段表示回源域名
	UpstreamDomain *string `json:"UpstreamDomain,omitnil" name:"UpstreamDomain"`

	// UpstreamType=0时，填次字段表示回源ip
	SrcList []*string `json:"SrcList,omitnil" name:"SrcList"`

	// 是否开启HTTP2，1表示开启，0表示不开启http2。开启HTTP2需要HTTPS支持
	IsHttp2 *int64 `json:"IsHttp2,omitnil" name:"IsHttp2"`

	// 是否开启WebSocket， 1：开启WebSocket，0：不开启WebSocket
	IsWebsocket *int64 `json:"IsWebsocket,omitnil" name:"IsWebsocket"`

	// 负载均衡策略，0表示轮徇，1表示IP hash
	LoadBalance *int64 `json:"LoadBalance,omitnil" name:"LoadBalance"`

	// 是否灰度
	IsGray *int64 `json:"IsGray,omitnil" name:"IsGray"`

	// WAF版本
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 端口信息
	Ports []*SpartaProtectionPort `json:"Ports,omitnil" name:"Ports"`

	// 长短连接标志，仅IP回源时有效
	IsKeepAlive *string `json:"IsKeepAlive,omitnil" name:"IsKeepAlive"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 是否为Anycast ip类型：1 Anycast 0 普通ip
	Anycast *int64 `json:"Anycast,omitnil" name:"Anycast"`

	// src的权重
	Weights []*int64 `json:"Weights,omitnil" name:"Weights"`

	// 是否开启源站的主动健康检测，1表示开启，0表示不开启
	ActiveCheck *int64 `json:"ActiveCheck,omitnil" name:"ActiveCheck"`

	// TLS版本信息
	TLSVersion *int64 `json:"TLSVersion,omitnil" name:"TLSVersion"`

	// 加密套件信息
	Ciphers []*int64 `json:"Ciphers,omitnil" name:"Ciphers"`

	// 0:不支持选择：默认模板  1:通用型模板 2:安全型模板 3:自定义模板
	CipherTemplate *int64 `json:"CipherTemplate,omitnil" name:"CipherTemplate"`

	// 300s
	ProxyReadTimeout *int64 `json:"ProxyReadTimeout,omitnil" name:"ProxyReadTimeout"`

	// 300s
	ProxySendTimeout *int64 `json:"ProxySendTimeout,omitnil" name:"ProxySendTimeout"`

	// 0:关闭SNI；1:开启SNI，SNI=源请求host；2:开启SNI，SNI=修改为源站host；3：开启SNI，自定义host，SNI=SniHost；
	SniType *int64 `json:"SniType,omitnil" name:"SniType"`

	// SniType=3时，需要填此参数，表示自定义的host；
	SniHost *string `json:"SniHost,omitnil" name:"SniHost"`

	// IsCdn=3时，需要填此参数，表示自定义header
	IpHeaders []*string `json:"IpHeaders,omitnil" name:"IpHeaders"`

	// 0:关闭xff重置；1:开启xff重置，只有在IsCdn=0时可以开启
	XFFReset *int64 `json:"XFFReset,omitnil" name:"XFFReset"`
}

type ModifySpartaProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 证书类型，0表示没有证书，CertType=1表示自有证书,2 为托管证书
	CertType *int64 `json:"CertType,omitnil" name:"CertType"`

	// CertType=1时，需要填次参数，表示证书内容
	Cert *string `json:"Cert,omitnil" name:"Cert"`

	// CertType=1时，需要填次参数，表示证书的私钥
	PrivateKey *string `json:"PrivateKey,omitnil" name:"PrivateKey"`

	// CertType=2时，需要填次参数，表示证书的ID
	SSLId *string `json:"SSLId,omitnil" name:"SSLId"`

	// 表示是否开启了CDN代理
	IsCdn *int64 `json:"IsCdn,omitnil" name:"IsCdn"`

	// HTTPS回源协议
	UpstreamScheme *string `json:"UpstreamScheme,omitnil" name:"UpstreamScheme"`

	// HTTPS回源端口,仅UpstreamScheme为http时需要填当前字段
	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitnil" name:"HttpsUpstreamPort"`

	// 表示是否强制跳转到HTTPS，1表示开启，0表示不开启
	HttpsRewrite *uint64 `json:"HttpsRewrite,omitnil" name:"HttpsRewrite"`

	// 回源类型，0表示通过IP回源,1 表示通过域名回源
	UpstreamType *int64 `json:"UpstreamType,omitnil" name:"UpstreamType"`

	// UpstreamType=1时，填次字段表示回源域名
	UpstreamDomain *string `json:"UpstreamDomain,omitnil" name:"UpstreamDomain"`

	// UpstreamType=0时，填次字段表示回源ip
	SrcList []*string `json:"SrcList,omitnil" name:"SrcList"`

	// 是否开启HTTP2，1表示开启，0表示不开启http2。开启HTTP2需要HTTPS支持
	IsHttp2 *int64 `json:"IsHttp2,omitnil" name:"IsHttp2"`

	// 是否开启WebSocket， 1：开启WebSocket，0：不开启WebSocket
	IsWebsocket *int64 `json:"IsWebsocket,omitnil" name:"IsWebsocket"`

	// 负载均衡策略，0表示轮徇，1表示IP hash
	LoadBalance *int64 `json:"LoadBalance,omitnil" name:"LoadBalance"`

	// 是否灰度
	IsGray *int64 `json:"IsGray,omitnil" name:"IsGray"`

	// WAF版本
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 端口信息
	Ports []*SpartaProtectionPort `json:"Ports,omitnil" name:"Ports"`

	// 长短连接标志，仅IP回源时有效
	IsKeepAlive *string `json:"IsKeepAlive,omitnil" name:"IsKeepAlive"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`

	// 是否为Anycast ip类型：1 Anycast 0 普通ip
	Anycast *int64 `json:"Anycast,omitnil" name:"Anycast"`

	// src的权重
	Weights []*int64 `json:"Weights,omitnil" name:"Weights"`

	// 是否开启源站的主动健康检测，1表示开启，0表示不开启
	ActiveCheck *int64 `json:"ActiveCheck,omitnil" name:"ActiveCheck"`

	// TLS版本信息
	TLSVersion *int64 `json:"TLSVersion,omitnil" name:"TLSVersion"`

	// 加密套件信息
	Ciphers []*int64 `json:"Ciphers,omitnil" name:"Ciphers"`

	// 0:不支持选择：默认模板  1:通用型模板 2:安全型模板 3:自定义模板
	CipherTemplate *int64 `json:"CipherTemplate,omitnil" name:"CipherTemplate"`

	// 300s
	ProxyReadTimeout *int64 `json:"ProxyReadTimeout,omitnil" name:"ProxyReadTimeout"`

	// 300s
	ProxySendTimeout *int64 `json:"ProxySendTimeout,omitnil" name:"ProxySendTimeout"`

	// 0:关闭SNI；1:开启SNI，SNI=源请求host；2:开启SNI，SNI=修改为源站host；3：开启SNI，自定义host，SNI=SniHost；
	SniType *int64 `json:"SniType,omitnil" name:"SniType"`

	// SniType=3时，需要填此参数，表示自定义的host；
	SniHost *string `json:"SniHost,omitnil" name:"SniHost"`

	// IsCdn=3时，需要填此参数，表示自定义header
	IpHeaders []*string `json:"IpHeaders,omitnil" name:"IpHeaders"`

	// 0:关闭xff重置；1:开启xff重置，只有在IsCdn=0时可以开启
	XFFReset *int64 `json:"XFFReset,omitnil" name:"XFFReset"`
}

func (r *ModifySpartaProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySpartaProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "DomainId")
	delete(f, "CertType")
	delete(f, "Cert")
	delete(f, "PrivateKey")
	delete(f, "SSLId")
	delete(f, "IsCdn")
	delete(f, "UpstreamScheme")
	delete(f, "HttpsUpstreamPort")
	delete(f, "HttpsRewrite")
	delete(f, "UpstreamType")
	delete(f, "UpstreamDomain")
	delete(f, "SrcList")
	delete(f, "IsHttp2")
	delete(f, "IsWebsocket")
	delete(f, "LoadBalance")
	delete(f, "IsGray")
	delete(f, "Edition")
	delete(f, "Ports")
	delete(f, "IsKeepAlive")
	delete(f, "InstanceID")
	delete(f, "Anycast")
	delete(f, "Weights")
	delete(f, "ActiveCheck")
	delete(f, "TLSVersion")
	delete(f, "Ciphers")
	delete(f, "CipherTemplate")
	delete(f, "ProxyReadTimeout")
	delete(f, "ProxySendTimeout")
	delete(f, "SniType")
	delete(f, "SniHost")
	delete(f, "IpHeaders")
	delete(f, "XFFReset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySpartaProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySpartaProtectionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySpartaProtectionResponse struct {
	*tchttp.BaseResponse
	Response *ModifySpartaProtectionResponseParams `json:"Response"`
}

func (r *ModifySpartaProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySpartaProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserLevelRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 防护规则等级 300=standard，400=extended
	Level *uint64 `json:"Level,omitnil" name:"Level"`
}

type ModifyUserLevelRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 防护规则等级 300=standard，400=extended
	Level *uint64 `json:"Level,omitnil" name:"Level"`
}

func (r *ModifyUserLevelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserLevelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Level")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserLevelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserLevelResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyUserLevelResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserLevelResponseParams `json:"Response"`
}

func (r *ModifyUserLevelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserLevelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserSignatureRuleRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 主类id
	MainClassID *string `json:"MainClassID,omitnil" name:"MainClassID"`

	// 主类开关0=关闭，1=开启，2=只告警
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 下发修改的规则列表
	RuleID []*ReqUserRule `json:"RuleID,omitnil" name:"RuleID"`
}

type ModifyUserSignatureRuleRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 主类id
	MainClassID *string `json:"MainClassID,omitnil" name:"MainClassID"`

	// 主类开关0=关闭，1=开启，2=只告警
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 下发修改的规则列表
	RuleID []*ReqUserRule `json:"RuleID,omitnil" name:"RuleID"`
}

func (r *ModifyUserSignatureRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserSignatureRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "MainClassID")
	delete(f, "Status")
	delete(f, "RuleID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserSignatureRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserSignatureRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyUserSignatureRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserSignatureRuleResponseParams `json:"Response"`
}

func (r *ModifyUserSignatureRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserSignatureRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWafAutoDenyRulesRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 触发IP封禁的攻击次数阈值，范围为2~100次
	AttackThreshold *int64 `json:"AttackThreshold,omitnil" name:"AttackThreshold"`

	// IP封禁统计时间，范围为1-60分钟
	TimeThreshold *int64 `json:"TimeThreshold,omitnil" name:"TimeThreshold"`

	// 触发IP封禁后的封禁时间，范围为5~360分钟
	DenyTimeThreshold *int64 `json:"DenyTimeThreshold,omitnil" name:"DenyTimeThreshold"`

	// 自动封禁状态，0表示关闭，1表示打开
	DefenseStatus *int64 `json:"DefenseStatus,omitnil" name:"DefenseStatus"`
}

type ModifyWafAutoDenyRulesRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 触发IP封禁的攻击次数阈值，范围为2~100次
	AttackThreshold *int64 `json:"AttackThreshold,omitnil" name:"AttackThreshold"`

	// IP封禁统计时间，范围为1-60分钟
	TimeThreshold *int64 `json:"TimeThreshold,omitnil" name:"TimeThreshold"`

	// 触发IP封禁后的封禁时间，范围为5~360分钟
	DenyTimeThreshold *int64 `json:"DenyTimeThreshold,omitnil" name:"DenyTimeThreshold"`

	// 自动封禁状态，0表示关闭，1表示打开
	DefenseStatus *int64 `json:"DefenseStatus,omitnil" name:"DefenseStatus"`
}

func (r *ModifyWafAutoDenyRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWafAutoDenyRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "AttackThreshold")
	delete(f, "TimeThreshold")
	delete(f, "DenyTimeThreshold")
	delete(f, "DefenseStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWafAutoDenyRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWafAutoDenyRulesResponseParams struct {
	// 成功的状态码，需要JSON解码后再使用，返回的格式是{"域名":"状态"}，成功的状态码为Success，其它的为失败的状态码（yunapi定义的错误码）
	Success *ResponseCode `json:"Success,omitnil" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyWafAutoDenyRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWafAutoDenyRulesResponseParams `json:"Response"`
}

func (r *ModifyWafAutoDenyRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWafAutoDenyRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWafThreatenIntelligenceRequestParams struct {
	// 配置WAF威胁情报封禁模块详情
	WafThreatenIntelligenceDetails *WafThreatenIntelligenceDetails `json:"WafThreatenIntelligenceDetails,omitnil" name:"WafThreatenIntelligenceDetails"`
}

type ModifyWafThreatenIntelligenceRequest struct {
	*tchttp.BaseRequest
	
	// 配置WAF威胁情报封禁模块详情
	WafThreatenIntelligenceDetails *WafThreatenIntelligenceDetails `json:"WafThreatenIntelligenceDetails,omitnil" name:"WafThreatenIntelligenceDetails"`
}

func (r *ModifyWafThreatenIntelligenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWafThreatenIntelligenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WafThreatenIntelligenceDetails")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWafThreatenIntelligenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWafThreatenIntelligenceResponseParams struct {
	// 当前WAF威胁情报封禁模块详情
	WafThreatenIntelligenceDetails *WafThreatenIntelligenceDetails `json:"WafThreatenIntelligenceDetails,omitnil" name:"WafThreatenIntelligenceDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyWafThreatenIntelligenceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWafThreatenIntelligenceResponseParams `json:"Response"`
}

func (r *ModifyWafThreatenIntelligenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWafThreatenIntelligenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWebshellStatusRequestParams struct {
	// 域名webshell状态
	Webshell *WebshellStatus `json:"Webshell,omitnil" name:"Webshell"`
}

type ModifyWebshellStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名webshell状态
	Webshell *WebshellStatus `json:"Webshell,omitnil" name:"Webshell"`
}

func (r *ModifyWebshellStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebshellStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Webshell")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWebshellStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWebshellStatusResponseParams struct {
	// 成功的状态码，需要JSON解码后再使用，返回的格式是{"域名":"状态"}，成功的状态码为Success，其它的为失败的状态码（yunapi定义的错误码）
	Success *ResponseCode `json:"Success,omitnil" name:"Success"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyWebshellStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWebshellStatusResponseParams `json:"Response"`
}

func (r *ModifyWebshellStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWebshellStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PageInfo struct {
	// 页码
	PageNumber *string `json:"PageNumber,omitnil" name:"PageNumber"`

	// 页条目数量
	PageSize *string `json:"PageSize,omitnil" name:"PageSize"`
}

type PeakPointsItem struct {
	// 秒级别时间戳
	Time *uint64 `json:"Time,omitnil" name:"Time"`

	// QPS
	Access *uint64 `json:"Access,omitnil" name:"Access"`

	// 上行带宽峰值，单位B
	Up *uint64 `json:"Up,omitnil" name:"Up"`

	// 下行带宽峰值，单位B
	Down *uint64 `json:"Down,omitnil" name:"Down"`

	// Web攻击次数
	Attack *uint64 `json:"Attack,omitnil" name:"Attack"`

	// CC攻击次数
	Cc *uint64 `json:"Cc,omitnil" name:"Cc"`

	// Bot qps
	BotAccess *uint64 `json:"BotAccess,omitnil" name:"BotAccess"`

	// WAF返回给客户端状态码5xx次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusServerError *uint64 `json:"StatusServerError,omitnil" name:"StatusServerError"`

	// WAF返回给客户端状态码4xx次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusClientError *uint64 `json:"StatusClientError,omitnil" name:"StatusClientError"`

	// WAF返回给客户端状态码302次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusRedirect *uint64 `json:"StatusRedirect,omitnil" name:"StatusRedirect"`

	// WAF返回给客户端状态码202次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusOk *uint64 `json:"StatusOk,omitnil" name:"StatusOk"`

	// 源站返回给WAF状态码5xx次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamServerError *uint64 `json:"UpstreamServerError,omitnil" name:"UpstreamServerError"`

	// 源站返回给WAF状态码4xx次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamClientError *uint64 `json:"UpstreamClientError,omitnil" name:"UpstreamClientError"`

	// 源站返回给WAF状态码302次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamRedirect *uint64 `json:"UpstreamRedirect,omitnil" name:"UpstreamRedirect"`

	// 黑名单次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlackIP *uint64 `json:"BlackIP,omitnil" name:"BlackIP"`

	// 防篡改次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tamper *uint64 `json:"Tamper,omitnil" name:"Tamper"`

	// 信息防泄露次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Leak *uint64 `json:"Leak,omitnil" name:"Leak"`

	// 访问控制 
	// 注意：此字段可能返回 null，表示取不到有效值。
	ACL *uint64 `json:"ACL,omitnil" name:"ACL"`

	// 小程序 qps
	// 注意：此字段可能返回 null，表示取不到有效值。
	WxAccess *uint64 `json:"WxAccess,omitnil" name:"WxAccess"`
}

type PiechartItem struct {
	// 类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 数量
	Count *uint64 `json:"Count,omitnil" name:"Count"`
}

type PortInfo struct {
	// Nginx的服务器id
	NginxServerId *uint64 `json:"NginxServerId,omitnil" name:"NginxServerId"`

	// 监听端口配置
	Port *string `json:"Port,omitnil" name:"Port"`

	// 与端口对应的协议
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 回源端口
	UpstreamPort *string `json:"UpstreamPort,omitnil" name:"UpstreamPort"`

	// 回源协议
	UpstreamProtocol *string `json:"UpstreamProtocol,omitnil" name:"UpstreamProtocol"`
}

type PortItem struct {
	// 监听端口配置
	Port *string `json:"Port,omitnil" name:"Port"`

	// 与Port一一对应，表示端口对应的协议
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 与Port一一对应,  表示回源端口
	UpstreamPort *string `json:"UpstreamPort,omitnil" name:"UpstreamPort"`

	// 与Port一一对应,  表示回源协议
	UpstreamProtocol *string `json:"UpstreamProtocol,omitnil" name:"UpstreamProtocol"`

	// Nginx的服务器ID
	NginxServerId *string `json:"NginxServerId,omitnil" name:"NginxServerId"`
}

// Predefined struct for user
type PostAttackDownloadTaskRequestParams struct {
	// 查询的域名，所有域名使用all
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Lucene语法
	QueryString *string `json:"QueryString,omitnil" name:"QueryString"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 默认为desc，可以取值desc和asc
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 下载的日志条数
	Count *int64 `json:"Count,omitnil" name:"Count"`
}

type PostAttackDownloadTaskRequest struct {
	*tchttp.BaseRequest
	
	// 查询的域名，所有域名使用all
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// Lucene语法
	QueryString *string `json:"QueryString,omitnil" name:"QueryString"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 默认为desc，可以取值desc和asc
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 下载的日志条数
	Count *int64 `json:"Count,omitnil" name:"Count"`
}

func (r *PostAttackDownloadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PostAttackDownloadTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QueryString")
	delete(f, "TaskName")
	delete(f, "Sort")
	delete(f, "Count")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PostAttackDownloadTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PostAttackDownloadTaskResponseParams struct {
	// 任务task id
	Flow *string `json:"Flow,omitnil" name:"Flow"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PostAttackDownloadTaskResponse struct {
	*tchttp.BaseResponse
	Response *PostAttackDownloadTaskResponseParams `json:"Response"`
}

func (r *PostAttackDownloadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PostAttackDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductInfo struct {
	// 产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type QPSPackageNew struct {
	// 资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceIds *string `json:"ResourceIds,omitnil" name:"ResourceIds"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValidTime *string `json:"ValidTime,omitnil" name:"ValidTime"`

	// 是否自动续费，1：自动续费，0：不自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitnil" name:"RenewFlag"`

	// 套餐购买个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 套餐购买地域，clb-waf暂时没有用到
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 计费项
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingItem *string `json:"BillingItem,omitnil" name:"BillingItem"`
}

type QpsData struct {
	// 弹性qps默认值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ElasticBillingDefault *uint64 `json:"ElasticBillingDefault,omitnil" name:"ElasticBillingDefault"`

	// 弹性qps最小值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ElasticBillingMin *uint64 `json:"ElasticBillingMin,omitnil" name:"ElasticBillingMin"`

	// 弹性qps最大值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ElasticBillingMax *uint64 `json:"ElasticBillingMax,omitnil" name:"ElasticBillingMax"`

	// 业务扩展包最大qps
	// 注意：此字段可能返回 null，表示取不到有效值。
	QPSExtendMax *uint64 `json:"QPSExtendMax,omitnil" name:"QPSExtendMax"`

	// 海外业务扩展包最大qps
	// 注意：此字段可能返回 null，表示取不到有效值。
	QPSExtendIntlMax *uint64 `json:"QPSExtendIntlMax,omitnil" name:"QPSExtendIntlMax"`
}

// Predefined struct for user
type RefreshAccessCheckResultRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type RefreshAccessCheckResultRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *RefreshAccessCheckResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshAccessCheckResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefreshAccessCheckResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefreshAccessCheckResultResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RefreshAccessCheckResultResponse struct {
	*tchttp.BaseResponse
	Response *RefreshAccessCheckResultResponseParams `json:"Response"`
}

func (r *RefreshAccessCheckResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefreshAccessCheckResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReqUserRule struct {
	// 特征序号
	Id *string `json:"Id,omitnil" name:"Id"`

	// 规则开关
	// 0：关
	// 1：开
	// 2：只告警
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 修改原因
	// 0：无(兼容记录为空)
	// 1：业务自身特性误报避免
	// 2：规则误报上报
	// 3：核心业务规则灰度
	// 4：其它
	Reason *int64 `json:"Reason,omitnil" name:"Reason"`
}

type ResponseCode struct {
	// 如果成功则返回Success，失败则返回云api定义的错误码
	Code *string `json:"Code,omitnil" name:"Code"`

	// 如果成功则返回Success，失败则返回WAF定义的二级错误码
	Message *string `json:"Message,omitnil" name:"Message"`
}

type Rule struct {
	// 规则id
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 规则类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 规则等级
	Level *string `json:"Level,omitnil" name:"Level"`

	// 规则描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 规则防护的CVE编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVE *string `json:"CVE,omitnil" name:"CVE"`

	// 规则的状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 规则修改的时间
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 门神规则新增/更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddTime *string `json:"AddTime,omitnil" name:"AddTime"`
}

type RuleList struct {
	// 规则Id
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 规则列表的id
	Rules []*uint64 `json:"Rules,omitnil" name:"Rules"`

	// 请求url
	Url *string `json:"Url,omitnil" name:"Url"`

	// 请求的方法
	Function *string `json:"Function,omitnil" name:"Function"`

	// 时间戳
	Time *string `json:"Time,omitnil" name:"Time"`

	// 开关状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

// Predefined struct for user
type SearchAccessLogRequestParams struct {
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题，新版本此字段填空字符串
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitnil" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitnil" name:"To"`

	// 查询语句，语句长度最大为4096
	Query *string `json:"Query,omitnil" name:"Query"`

	// 单次查询返回的日志条数，最大值为100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 新版本此字段失效，填空字符串，翻页使用Page
	Context *string `json:"Context,omitnil" name:"Context"`

	// 日志接口是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 第几页，从0开始。新版本接口字段
	Page *int64 `json:"Page,omitnil" name:"Page"`
}

type SearchAccessLogRequest struct {
	*tchttp.BaseRequest
	
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题，新版本此字段填空字符串
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitnil" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitnil" name:"To"`

	// 查询语句，语句长度最大为4096
	Query *string `json:"Query,omitnil" name:"Query"`

	// 单次查询返回的日志条数，最大值为100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 新版本此字段失效，填空字符串，翻页使用Page
	Context *string `json:"Context,omitnil" name:"Context"`

	// 日志接口是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 第几页，从0开始。新版本接口字段
	Page *int64 `json:"Page,omitnil" name:"Page"`
}

func (r *SearchAccessLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchAccessLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "Sort")
	delete(f, "Page")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchAccessLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchAccessLogResponseParams struct {
	// 新接口此字段失效，默认返回空字符串
	Context *string `json:"Context,omitnil" name:"Context"`

	// 日志查询结果是否全部返回，其中，“true”表示结果返回，“false”表示结果为返回
	ListOver *bool `json:"ListOver,omitnil" name:"ListOver"`

	// 返回的是否为分析结果，其中，“true”表示返回分析结果，“false”表示未返回分析结果
	Analysis *bool `json:"Analysis,omitnil" name:"Analysis"`

	// 如果Analysis为True，则返回分析结果的列名，否则为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColNames []*string `json:"ColNames,omitnil" name:"ColNames"`

	// 日志查询结果；当Analysis为True时，可能返回为null
	// 注意：此字段可能返回 null，表示取不到有效值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*AccessLogInfo `json:"Results,omitnil" name:"Results"`

	// 日志分析结果；当Analysis为False时，可能返回为null
	// 注意：此字段可能返回 null，表示取不到有效值
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnalysisResults []*AccessLogItems `json:"AnalysisResults,omitnil" name:"AnalysisResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SearchAccessLogResponse struct {
	*tchttp.BaseResponse
	Response *SearchAccessLogResponseParams `json:"Response"`
}

func (r *SearchAccessLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchAccessLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchAttackLogRequestParams struct {
	// 查询的域名，所有域名使用all
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 接口升级，这个字段传空字符串,翻页使用Page字段
	Context *string `json:"Context,omitnil" name:"Context"`

	// Lucene语法
	QueryString *string `json:"QueryString,omitnil" name:"QueryString"`

	// 查询的数量，默认10条，最多100条
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 默认为desc，可以取值desc和asc
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 第几页，从0开始
	Page *int64 `json:"Page,omitnil" name:"Page"`
}

type SearchAttackLogRequest struct {
	*tchttp.BaseRequest
	
	// 查询的域名，所有域名使用all
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 接口升级，这个字段传空字符串,翻页使用Page字段
	Context *string `json:"Context,omitnil" name:"Context"`

	// Lucene语法
	QueryString *string `json:"QueryString,omitnil" name:"QueryString"`

	// 查询的数量，默认10条，最多100条
	Count *int64 `json:"Count,omitnil" name:"Count"`

	// 默认为desc，可以取值desc和asc
	Sort *string `json:"Sort,omitnil" name:"Sort"`

	// 第几页，从0开始
	Page *int64 `json:"Page,omitnil" name:"Page"`
}

func (r *SearchAttackLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchAttackLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Context")
	delete(f, "QueryString")
	delete(f, "Count")
	delete(f, "Sort")
	delete(f, "Page")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchAttackLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchAttackLogResponseParams struct {
	// 当前返回的攻击日志条数
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 接口升级，此字段无效，默认返回空字符串
	Context *string `json:"Context,omitnil" name:"Context"`

	// 攻击日志数组条目内容
	Data []*AttackLogInfo `json:"Data,omitnil" name:"Data"`

	// CLS接口返回内容
	ListOver *bool `json:"ListOver,omitnil" name:"ListOver"`

	// CLS接口返回内容，标志是否启动新版本索引
	SqlFlag *bool `json:"SqlFlag,omitnil" name:"SqlFlag"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SearchAttackLogResponse struct {
	*tchttp.BaseResponse
	Response *SearchAttackLogResponseParams `json:"Response"`
}

func (r *SearchAttackLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchAttackLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchItem struct {
	// 日志开关
	ClsStatus *string `json:"ClsStatus,omitnil" name:"ClsStatus"`

	// waf开关
	Status *string `json:"Status,omitnil" name:"Status"`

	// 流量模式
	FlowMode *string `json:"FlowMode,omitnil" name:"FlowMode"`
}

type SessionData struct {
	// session定义
	Res []*SessionItem `json:"Res,omitnil" name:"Res"`
}

type SessionItem struct {
	// 匹配类型
	Category *string `json:"Category,omitnil" name:"Category"`

	// 起始模式
	KeyOrStartMat *string `json:"KeyOrStartMat,omitnil" name:"KeyOrStartMat"`

	// 结束模式
	EndMat *string `json:"EndMat,omitnil" name:"EndMat"`

	// 起始偏移
	StartOffset *string `json:"StartOffset,omitnil" name:"StartOffset"`

	// 结束偏移
	EndOffset *string `json:"EndOffset,omitnil" name:"EndOffset"`

	// 数据源
	Source *string `json:"Source,omitnil" name:"Source"`

	// 更新时间戳
	TsVersion *string `json:"TsVersion,omitnil" name:"TsVersion"`

	// SessionID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *int64 `json:"SessionId,omitnil" name:"SessionId"`

	// Session名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionName *string `json:"SessionName,omitnil" name:"SessionName"`

	// Session是否正在被启用
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionInUsed *bool `json:"SessionInUsed,omitnil" name:"SessionInUsed"`

	// Session关联的CC规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelatedRuleID []*int64 `json:"RelatedRuleID,omitnil" name:"RelatedRuleID"`
}

type SpartaProtectionPort struct {
	// nginx Id
	NginxServerId *uint64 `json:"NginxServerId,omitnil" name:"NginxServerId"`

	// 端口
	Port *string `json:"Port,omitnil" name:"Port"`

	// 协议
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 后端端口
	UpstreamPort *string `json:"UpstreamPort,omitnil" name:"UpstreamPort"`

	// 后端协议
	UpstreamProtocol *string `json:"UpstreamProtocol,omitnil" name:"UpstreamProtocol"`
}

type Strategy struct {
	// 匹配字段
	// 
	//     匹配字段不同，相应的匹配参数、逻辑符号、匹配内容有所不同
	// 具体如下所示：
	// <table><thead><tr><th>匹配字段</th><th>匹配参数</th><th>逻辑符号</th><th>匹配内容</th></tr></thead><tbody><tr><td>IP（来源IP）</td><td>不支持参数</td><td>ipmatch（匹配）<br/>ipnmatch（不匹配）</td><td>多个IP以英文逗号隔开,最多20个</td></tr><tr><td>IPV6（来源IPv6）</td><td>不支持参数</td><td>ipmatch（匹配）<br/>ipnmatch（不匹配）</td><td>支持单个IPV6地址</td></tr><tr><td>Referer（Referer）</td><td>不支持参数</td><td>empty（内容为空）<br/>null（不存在）<br/>eq（等于）<br/>neq（不等于）<br/>contains（包含）<br/>ncontains（不包含）<br/>len_eq（长度等于）<br/>len_gt（长度大于）<br/>len_lt（长度小于）<br/>strprefix（前缀匹配）<br/>strsuffix（后缀匹配）<br/>rematch（正则匹配）</td><td>请输入内容,512个字符以内</td></tr><tr><td>URL（请求路径）</td><td>不支持参数</td><td>eq（等于）<br/>neq（不等于）<br/>contains（包含）<br/>ncontains（不包含）<br/>len_eq（长度等于）<br/>len_gt（长度大于）<br/>len_lt（长度小于）<br/>strprefix（前缀匹配）<br/>strsuffix（后缀匹配）<br/>rematch（正则匹配）<br/></td><td>请以/开头,512个字符以内</td></tr><tr><td>UserAgent（UserAgent）</td><td>不支持参数</td><td>同匹配字段<font color="Red">Referer</font>逻辑符号</td><td>请输入内容,512个字符以内</td></tr><tr><td>HTTP_METHOD（HTTP请求方法）</td><td>不支持参数</td><td>eq（等于）<br/>neq（不等于）</td><td>请输入方法名称,建议大写</td></tr><tr><td>QUERY_STRING（请求字符串）</td><td>不支持参数</td><td>同匹配字段<font color="Red">请求路径</font>逻辑符号</td><td>请输入内容,512个字符以内</td></tr><tr><td>GET（GET参数值）</td><td>支持参数录入</td><td>contains（包含）<br/>ncontains（不包含）<br/>len_eq（长度等于）<br/>len_gt（长度大于）<br/>len_lt（长度小于）<br/>strprefix（前缀匹配）<br/>strsuffix（后缀匹配）</td><td>请输入内容,512个字符以内</td></tr><tr><td>GET_PARAMS_NAMES（GET参数名）</td><td>不支持参数</td><td>exsit（存在参数）<br/>nexsit（不存在参数）<br/>len_eq（长度等于）<br/>len_gt（长度大于）<br/>len_lt（长度小于）<br/>strprefix（前缀匹配）<br/>strsuffix（后缀匹配）</td><td>请输入内容,512个字符以内</td></tr><tr><td>POST（POST参数值）</td><td>支持参数录入</td><td>同匹配字段<font color="Red">GET参数值</font>逻辑符号</td><td>请输入内容,512个字符以内</td></tr><tr><td>GET_POST_NAMES（POST参数名）</td><td>不支持参数</td><td>同匹配字段<font color="Red">GET参数名</font>逻辑符号</td><td>请输入内容,512个字符以内</td></tr><tr><td>POST_BODY（完整BODY）</td><td>不支持参数</td><td>同匹配字段<font color="Red">请求路径</font>逻辑符号</td><td>请输入BODY内容,512个字符以内</td></tr><tr><td>COOKIE（Cookie）</td><td>不支持参数</td><td>empty（内容为空）<br/>null（不存在）<br/>rematch（正则匹配）</td><td><font color="Red">暂不支持</font></td></tr><tr><td>GET_COOKIES_NAMES（Cookie参数名）</td><td>不支持参数</td><td>同匹配字段<font color="Red">GET参数名</font>逻辑符号</td><td>请输入内容,512个字符以内</td></tr><tr><td>ARGS_COOKIE（Cookie参数值）</td><td>支持参数录入</td><td>同匹配字段<font color="Red">GET参数值</font>逻辑符号</td><td>请输入内容,512个字符以内</td></tr><tr><td>GET_HEADERS_NAMES（Header参数名）</td><td>不支持参数</td><td>exsit（存在参数）<br/>nexsit（不存在参数）<br/>len_eq（长度等于）<br/>len_gt（长度大于）<br/>len_lt（长度小于）<br/>strprefix（前缀匹配）<br/>strsuffix（后缀匹配）<br/>rematch（正则匹配）</td><td>请输入内容,建议小写,512个字符以内</td></tr><tr><td>ARGS_HEADER（Header参数值）</td><td>支持参数录入</td><td>contains（包含）<br/>ncontains（不包含）<br/>len_eq（长度等于）<br/>len_gt（长度大于）<br/>len_lt（长度小于）<br/>strprefix（前缀匹配）<br/>strsuffix（后缀匹配）<br/>rematch（正则匹配）</td><td>请输入内容,512个字符以内</td></tr></tbody></table>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Field *string `json:"Field,omitnil" name:"Field"`

	// 逻辑符号 
	// 
	//     逻辑符号一共分为以下几种类型：
	//         empty （ 内容为空）
	//         null （不存在）
	//         eq （ 等于）
	//         neq （ 不等于）
	//         contains （ 包含）
	//         ncontains （ 不包含）
	//         strprefix （ 前缀匹配）
	//         strsuffix （ 后缀匹配）
	//         len_eq （ 长度等于）
	//         len_gt （ 长度大于）
	//         len_lt （ 长度小于）
	//         ipmatch （ 属于）
	//         ipnmatch （ 不属于）
	//     各匹配字段对应的逻辑符号不同，详见上述匹配字段表格
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareFunc *string `json:"CompareFunc,omitnil" name:"CompareFunc"`

	// 匹配内容
	// 
	//     目前 当匹配字段为COOKIE（Cookie）时，不需要输入 匹配内容
	// 其他都需要
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil" name:"Content"`

	// 匹配参数
	// 
	//     配置参数一共分2种类型 不支持参数与支持参数
	//     当匹配字段为以下4个时，匹配参数才能录入，否则不支持该参数
	//         GET（GET参数值）
	// 		
	//         POST（POST参数值）
	// 		
	//         ARGS_COOKIE（Cookie参数值）
	// 		
	//         ARGS_HEADER（Header参数值）
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Arg *string `json:"Arg,omitnil" name:"Arg"`
}

type StrategyForAntiInfoLeak struct {
	// 匹配条件，returncode（响应码）、keywords（关键字）、information（敏感信息）
	Field *string `json:"Field,omitnil" name:"Field"`

	// 逻辑符号，固定取值为contains
	CompareFunc *string `json:"CompareFunc,omitnil" name:"CompareFunc"`

	// 匹配内容。
	// 以下三个对应Field为information时可取的匹配内容：
	// idcard（身份证）、phone（手机号）、bankcard（银行卡）。
	// 以下为对应Field为returncode时可取的匹配内容：
	// 400（状态码400）、403（状态码403）、404（状态码404）、4xx（其它4xx状态码）、500（状态码500）、501（状态码501）、502（状态码502）、504（状态码504）、5xx（其它5xx状态码）。
	// 当对应Field为keywords时由用户自己输入匹配内容。
	Content *string `json:"Content,omitnil" name:"Content"`
}

// Predefined struct for user
type SwitchDomainRulesRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 规则列表
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`

	// 开关状态，0表示关闭，1表示开启，2表示只观察
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 设置为观察模式原因，
	// 1表示业务自身原因观察，2表示系统规则误报上报，3表示核心业务灰度观察，4表示其他
	Reason *uint64 `json:"Reason,omitnil" name:"Reason"`
}

type SwitchDomainRulesRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 规则列表
	Ids []*uint64 `json:"Ids,omitnil" name:"Ids"`

	// 开关状态，0表示关闭，1表示开启，2表示只观察
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 设置为观察模式原因，
	// 1表示业务自身原因观察，2表示系统规则误报上报，3表示核心业务灰度观察，4表示其他
	Reason *uint64 `json:"Reason,omitnil" name:"Reason"`
}

func (r *SwitchDomainRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDomainRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Ids")
	delete(f, "Status")
	delete(f, "Reason")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchDomainRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDomainRulesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SwitchDomainRulesResponse struct {
	*tchttp.BaseResponse
	Response *SwitchDomainRulesResponseParams `json:"Response"`
}

func (r *SwitchDomainRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDomainRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchElasticModeRequestParams struct {
	// 版本，只能是sparta-waf, clb-waf, cdn-waf
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 0代表关闭，1代表打开
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

type SwitchElasticModeRequest struct {
	*tchttp.BaseRequest
	
	// 版本，只能是sparta-waf, clb-waf, cdn-waf
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 0代表关闭，1代表打开
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

func (r *SwitchElasticModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchElasticModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Edition")
	delete(f, "Mode")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchElasticModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchElasticModeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SwitchElasticModeResponse struct {
	*tchttp.BaseResponse
	Response *SwitchElasticModeResponseParams `json:"Response"`
}

func (r *SwitchElasticModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchElasticModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TLSCiphers struct {
	// TLS版本ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionId *int64 `json:"VersionId,omitnil" name:"VersionId"`

	// 加密套件ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CipherId *int64 `json:"CipherId,omitnil" name:"CipherId"`

	// 加密套件
	// 注意：此字段可能返回 null，表示取不到有效值。
	CipherName *string `json:"CipherName,omitnil" name:"CipherName"`
}

type TLSVersion struct {
	// TLSVERSION的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionId *int64 `json:"VersionId,omitnil" name:"VersionId"`

	// TLSVERSION的NAME
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionName *string `json:"VersionName,omitnil" name:"VersionName"`
}

type TargetEntity struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

// Predefined struct for user
type UpsertCCAutoStatusRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 状态值
	Value *int64 `json:"Value,omitnil" name:"Value"`

	// 版本：clb-waf, spart-waf
	Edition *string `json:"Edition,omitnil" name:"Edition"`
}

type UpsertCCAutoStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 状态值
	Value *int64 `json:"Value,omitnil" name:"Value"`

	// 版本：clb-waf, spart-waf
	Edition *string `json:"Edition,omitnil" name:"Edition"`
}

func (r *UpsertCCAutoStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpsertCCAutoStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Value")
	delete(f, "Edition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpsertCCAutoStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpsertCCAutoStatusResponseParams struct {
	// 正常情况为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpsertCCAutoStatusResponse struct {
	*tchttp.BaseResponse
	Response *UpsertCCAutoStatusResponseParams `json:"Response"`
}

func (r *UpsertCCAutoStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpsertCCAutoStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpsertCCRuleRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 状态
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 高级模式
	Advance *string `json:"Advance,omitnil" name:"Advance"`

	// CC检测阈值
	Limit *string `json:"Limit,omitnil" name:"Limit"`

	// CC检测周期
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 检测Url
	Url *string `json:"Url,omitnil" name:"Url"`

	// 匹配方法
	MatchFunc *int64 `json:"MatchFunc,omitnil" name:"MatchFunc"`

	// 动作
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// 优先级
	Priority *int64 `json:"Priority,omitnil" name:"Priority"`

	// 动作有效时间
	ValidTime *int64 `json:"ValidTime,omitnil" name:"ValidTime"`

	// 附加参数
	OptionsArr *string `json:"OptionsArr,omitnil" name:"OptionsArr"`

	// waf版本
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 操作类型
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 添加规则的来源事件id
	EventId *string `json:"EventId,omitnil" name:"EventId"`

	// 规则需要启用的SessionID
	SessionApplied []*int64 `json:"SessionApplied,omitnil" name:"SessionApplied"`

	// 规则ID，新增时填0
	RuleId *int64 `json:"RuleId,omitnil" name:"RuleId"`
}

type UpsertCCRuleRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 状态
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 高级模式
	Advance *string `json:"Advance,omitnil" name:"Advance"`

	// CC检测阈值
	Limit *string `json:"Limit,omitnil" name:"Limit"`

	// CC检测周期
	Interval *string `json:"Interval,omitnil" name:"Interval"`

	// 检测Url
	Url *string `json:"Url,omitnil" name:"Url"`

	// 匹配方法
	MatchFunc *int64 `json:"MatchFunc,omitnil" name:"MatchFunc"`

	// 动作
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// 优先级
	Priority *int64 `json:"Priority,omitnil" name:"Priority"`

	// 动作有效时间
	ValidTime *int64 `json:"ValidTime,omitnil" name:"ValidTime"`

	// 附加参数
	OptionsArr *string `json:"OptionsArr,omitnil" name:"OptionsArr"`

	// waf版本
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 操作类型
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 添加规则的来源事件id
	EventId *string `json:"EventId,omitnil" name:"EventId"`

	// 规则需要启用的SessionID
	SessionApplied []*int64 `json:"SessionApplied,omitnil" name:"SessionApplied"`

	// 规则ID，新增时填0
	RuleId *int64 `json:"RuleId,omitnil" name:"RuleId"`
}

func (r *UpsertCCRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpsertCCRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Name")
	delete(f, "Status")
	delete(f, "Advance")
	delete(f, "Limit")
	delete(f, "Interval")
	delete(f, "Url")
	delete(f, "MatchFunc")
	delete(f, "ActionType")
	delete(f, "Priority")
	delete(f, "ValidTime")
	delete(f, "OptionsArr")
	delete(f, "Edition")
	delete(f, "Type")
	delete(f, "EventId")
	delete(f, "SessionApplied")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpsertCCRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpsertCCRuleResponseParams struct {
	// 一般为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil" name:"Data"`

	// 操作的RuleId
	RuleId *int64 `json:"RuleId,omitnil" name:"RuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpsertCCRuleResponse struct {
	*tchttp.BaseResponse
	Response *UpsertCCRuleResponseParams `json:"Response"`
}

func (r *UpsertCCRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpsertCCRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpsertIpAccessControlRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// ip 参数列表，json数组由ip，source，note，action，valid_ts组成。ip对应配置的ip地址，source固定为custom值，note为注释，action值42为黑名单，40为白名单，valid_ts为有效日期，值为秒级时间戳（（如1680570420代表2023-04-04 09:07:00））
	Items []*string `json:"Items,omitnil" name:"Items"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 是否为多域名黑白名单，当为多域名的黑白名单时，取值为batch，否则为空
	SourceType *string `json:"SourceType,omitnil" name:"SourceType"`
}

type UpsertIpAccessControlRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// ip 参数列表，json数组由ip，source，note，action，valid_ts组成。ip对应配置的ip地址，source固定为custom值，note为注释，action值42为黑名单，40为白名单，valid_ts为有效日期，值为秒级时间戳（（如1680570420代表2023-04-04 09:07:00））
	Items []*string `json:"Items,omitnil" name:"Items"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 是否为多域名黑白名单，当为多域名的黑白名单时，取值为batch，否则为空
	SourceType *string `json:"SourceType,omitnil" name:"SourceType"`
}

func (r *UpsertIpAccessControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpsertIpAccessControlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Items")
	delete(f, "InstanceId")
	delete(f, "Edition")
	delete(f, "SourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpsertIpAccessControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpsertIpAccessControlResponseParams struct {
	// 添加或修改失败的条目
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedItems *string `json:"FailedItems,omitnil" name:"FailedItems"`

	// 添加或修改失败的数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedCount *int64 `json:"FailedCount,omitnil" name:"FailedCount"`

	// 添加或修改的IP数据Id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ids []*string `json:"Ids,omitnil" name:"Ids"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpsertIpAccessControlResponse struct {
	*tchttp.BaseResponse
	Response *UpsertIpAccessControlResponseParams `json:"Response"`
}

func (r *UpsertIpAccessControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpsertIpAccessControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpsertSessionRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// session来源位置
	Source *string `json:"Source,omitnil" name:"Source"`

	// 提取类别
	Category *string `json:"Category,omitnil" name:"Category"`

	// 提取key或者起始匹配模式
	KeyOrStartMat *string `json:"KeyOrStartMat,omitnil" name:"KeyOrStartMat"`

	// 结束匹配模式
	EndMat *string `json:"EndMat,omitnil" name:"EndMat"`

	// 起始偏移位置
	StartOffset *string `json:"StartOffset,omitnil" name:"StartOffset"`

	// 结束偏移位置
	EndOffset *string `json:"EndOffset,omitnil" name:"EndOffset"`

	// 版本
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// Session名
	SessionName *string `json:"SessionName,omitnil" name:"SessionName"`

	// Session对应ID
	SessionID *int64 `json:"SessionID,omitnil" name:"SessionID"`
}

type UpsertSessionRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// session来源位置
	Source *string `json:"Source,omitnil" name:"Source"`

	// 提取类别
	Category *string `json:"Category,omitnil" name:"Category"`

	// 提取key或者起始匹配模式
	KeyOrStartMat *string `json:"KeyOrStartMat,omitnil" name:"KeyOrStartMat"`

	// 结束匹配模式
	EndMat *string `json:"EndMat,omitnil" name:"EndMat"`

	// 起始偏移位置
	StartOffset *string `json:"StartOffset,omitnil" name:"StartOffset"`

	// 结束偏移位置
	EndOffset *string `json:"EndOffset,omitnil" name:"EndOffset"`

	// 版本
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// Session名
	SessionName *string `json:"SessionName,omitnil" name:"SessionName"`

	// Session对应ID
	SessionID *int64 `json:"SessionID,omitnil" name:"SessionID"`
}

func (r *UpsertSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpsertSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Source")
	delete(f, "Category")
	delete(f, "KeyOrStartMat")
	delete(f, "EndMat")
	delete(f, "StartOffset")
	delete(f, "EndOffset")
	delete(f, "Edition")
	delete(f, "SessionName")
	delete(f, "SessionID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpsertSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpsertSessionResponseParams struct {
	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil" name:"Data"`

	// SessionID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionID *int64 `json:"SessionID,omitnil" name:"SessionID"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpsertSessionResponse struct {
	*tchttp.BaseResponse
	Response *UpsertSessionResponseParams `json:"Response"`
}

func (r *UpsertSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpsertSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserDomainInfo struct {
	// 用户id
	Appid *uint64 `json:"Appid,omitnil" name:"Appid"`

	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 域名id
	DomainId *string `json:"DomainId,omitnil" name:"DomainId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// waf类型
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *string `json:"Level,omitnil" name:"Level"`

	// 指定域名访问日志字段的开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	WriteConfig *string `json:"WriteConfig,omitnil" name:"WriteConfig"`

	// 指定域名是否写cls的开关 1:写 0:不写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cls *uint64 `json:"Cls,omitnil" name:"Cls"`

	// 标记是否是混合云接入。hybrid表示混合云接入域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloudType *string `json:"CloudType,omitnil" name:"CloudType"`
}

type UserSignatureRule struct {
	// 特征ID
	ID *string `json:"ID,omitnil" name:"ID"`

	// 规则开关
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 主类ID
	MainClassID *string `json:"MainClassID,omitnil" name:"MainClassID"`

	// 子类ID
	SubClassID *string `json:"SubClassID,omitnil" name:"SubClassID"`

	// CVE ID
	CveID *string `json:"CveID,omitnil" name:"CveID"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 更新时间
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 主类名字，根据Language字段输出中文/英文
	MainClassName *string `json:"MainClassName,omitnil" name:"MainClassName"`

	// 子类名字，根据Language字段输出中文/英文，若子类id为00000000，此字段为空
	SubClassName *string `json:"SubClassName,omitnil" name:"SubClassName"`

	// 规则描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 0/1
	Reason *int64 `json:"Reason,omitnil" name:"Reason"`
}

type UserWhiteRule struct {
	// 白名单的id
	WhiteRuleId *uint64 `json:"WhiteRuleId,omitnil" name:"WhiteRuleId"`

	// 规则id
	SignatureId *string `json:"SignatureId,omitnil" name:"SignatureId"`

	// 状态
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 匹配域
	MatchField *string `json:"MatchField,omitnil" name:"MatchField"`

	// 匹配方法
	MatchMethod *string `json:"MatchMethod,omitnil" name:"MatchMethod"`

	// 匹配内容
	MatchContent *string `json:"MatchContent,omitnil" name:"MatchContent"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`
}

type UserWhiteRuleItem struct {
	// 匹配域
	MatchField *string `json:"MatchField,omitnil" name:"MatchField"`

	// 匹配方法
	MatchMethod *string `json:"MatchMethod,omitnil" name:"MatchMethod"`

	// 匹配内容
	MatchContent *string `json:"MatchContent,omitnil" name:"MatchContent"`
}

type VipInfo struct {
	// Virtual IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// waf实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type WafRuleLimit struct {
	// 自定义CC的规格
	CC *uint64 `json:"CC,omitnil" name:"CC"`

	// 自定义策略的规格
	CustomRule *uint64 `json:"CustomRule,omitnil" name:"CustomRule"`

	// 黑白名单的规格
	IPControl *uint64 `json:"IPControl,omitnil" name:"IPControl"`

	// 信息防泄漏的规格
	AntiLeak *uint64 `json:"AntiLeak,omitnil" name:"AntiLeak"`

	// 防篡改的规格
	AntiTamper *uint64 `json:"AntiTamper,omitnil" name:"AntiTamper"`

	// 紧急CC的规格
	AutoCC *uint64 `json:"AutoCC,omitnil" name:"AutoCC"`

	// 地域封禁的规格
	AreaBan *uint64 `json:"AreaBan,omitnil" name:"AreaBan"`

	// 自定义CC中配置session
	CCSession *uint64 `json:"CCSession,omitnil" name:"CCSession"`

	// AI的规格
	AI *uint64 `json:"AI,omitnil" name:"AI"`

	// 精准白名单的规格
	CustomWhite *uint64 `json:"CustomWhite,omitnil" name:"CustomWhite"`

	// api安全的规格
	ApiSecurity *uint64 `json:"ApiSecurity,omitnil" name:"ApiSecurity"`

	// 客户端流量标记的规格
	ClientMsg *uint64 `json:"ClientMsg,omitnil" name:"ClientMsg"`

	// 流量标记的规格
	TrafficMarking *uint64 `json:"TrafficMarking,omitnil" name:"TrafficMarking"`
}

type WafThreatenIntelligenceDetails struct {
	// 封禁属性标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 封禁模组启用状态
	DefenseStatus *int64 `json:"DefenseStatus,omitnil" name:"DefenseStatus"`

	// 最后更新时间
	LastUpdateTime *string `json:"LastUpdateTime,omitnil" name:"LastUpdateTime"`
}

type WebshellStatus struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// webshell开关，1：开。0：关。2：观察
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}