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

type AccessFieldValueRatioInfo struct {
	// 日志条数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 对应的Value值的百分比
	Ratio *float64 `json:"Ratio,omitnil,omitempty" name:"Ratio"`

	// 字段对应的值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AccessFullTextInfo struct {
	// 是否大小写敏感
	CaseSensitive *bool `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// 全文索引的分词符，字符串中每个字符代表一个分词符
	Tokenizer *string `json:"Tokenizer,omitnil,omitempty" name:"Tokenizer"`

	// 是否包含中文
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainZH *bool `json:"ContainZH,omitnil,omitempty" name:"ContainZH"`
}

type AccessHistogramItem struct {
	// 时间，单位ms
	//
	// Deprecated: BTime is deprecated.
	BTime *int64 `json:"BTime,omitnil,omitempty" name:"BTime"`

	// 日志条数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 时间，单位ms
	BeginTime *int64 `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`
}

type AccessKeyValueInfo struct {
	// 需要配置键值或者元字段索引的字段
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 字段的索引描述信息
	Value *AccessValueInfo `json:"Value,omitnil,omitempty" name:"Value"`
}

type AccessLogInfo struct {
	// 日志时间，单位ms
	Time *int64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 日志来源IP
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 日志文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 日志上报请求包的ID
	PkgId *string `json:"PkgId,omitnil,omitempty" name:"PkgId"`

	// 请求包内日志的ID
	PkgLogId *string `json:"PkgLogId,omitnil,omitempty" name:"PkgLogId"`

	// 日志内容的Json序列化字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogJson *string `json:"LogJson,omitnil,omitempty" name:"LogJson"`
}

type AccessLogItem struct {
	// 日记Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 日志Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AccessLogItems struct {
	// 分析结果返回的KV数据对
	Data []*AccessLogItem `json:"Data,omitnil,omitempty" name:"Data"`
}

type AccessRuleInfo struct {
	// 全文索引配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullText *AccessFullTextInfo `json:"FullText,omitnil,omitempty" name:"FullText"`

	// 键值索引配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyValue *AccessRuleKeyValueInfo `json:"KeyValue,omitnil,omitempty" name:"KeyValue"`

	// 元字段索引配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *AccessRuleTagInfo `json:"Tag,omitnil,omitempty" name:"Tag"`
}

type AccessRuleKeyValueInfo struct {
	// 是否大小写敏感
	CaseSensitive *bool `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// 需要建立索引的键值对信息；最大只能配置100个键值对
	KeyValues []*AccessKeyValueInfo `json:"KeyValues,omitnil,omitempty" name:"KeyValues"`
}

type AccessRuleTagInfo struct {
	// 是否大小写敏感
	CaseSensitive *bool `json:"CaseSensitive,omitnil,omitempty" name:"CaseSensitive"`

	// 标签索引配置中的字段信息
	KeyValues []*AccessKeyValueInfo `json:"KeyValues,omitnil,omitempty" name:"KeyValues"`
}

type AccessValueInfo struct {
	// 字段类型，目前支持的类型有：long、text、double
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 字段的分词符，只有当字段类型为text时才有意义；输入字符串中的每个字符代表一个分词符
	Tokenizer *string `json:"Tokenizer,omitnil,omitempty" name:"Tokenizer"`

	// 字段是否开启分析功能
	SqlFlag *bool `json:"SqlFlag,omitnil,omitempty" name:"SqlFlag"`

	// 是否包含中文
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainZH *bool `json:"ContainZH,omitnil,omitempty" name:"ContainZH"`
}

// Predefined struct for user
type AddAntiFakeUrlRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// uri
	Uri *string `json:"Uri,omitnil,omitempty" name:"Uri"`
}

type AddAntiFakeUrlRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// uri
	Uri *string `json:"Uri,omitnil,omitempty" name:"Uri"`
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
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 规则ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 动作，0（告警）、1（替换）、2（仅显示前四位）、3（仅显示后四位）、4（阻断）
	ActionType *uint64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 策略详情
	Strategies []*StrategyForAntiInfoLeak `json:"Strategies,omitnil,omitempty" name:"Strategies"`

	// 网址
	Uri *string `json:"Uri,omitnil,omitempty" name:"Uri"`
}

type AddAntiInfoLeakRulesRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 动作，0（告警）、1（替换）、2（仅显示前四位）、3（仅显示后四位）、4（阻断）
	ActionType *uint64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 策略详情
	Strategies []*StrategyForAntiInfoLeak `json:"Strategies,omitnil,omitempty" name:"Strategies"`

	// 网址
	Uri *string `json:"Uri,omitnil,omitempty" name:"Uri"`
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
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type AddAreaBanAreasRequestParams struct {
	// 需要修改的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 需要新增的封禁地域
	Areas []*string `json:"Areas,omitnil,omitempty" name:"Areas"`

	// waf版本信息，spart-waf或者clb-waf，其他无效，请一定填写
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 规则执行的方式，TimedJob为定时执行，CronJob为周期执行
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 定时任务配置
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`
}

type AddAreaBanAreasRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 需要新增的封禁地域
	Areas []*string `json:"Areas,omitnil,omitempty" name:"Areas"`

	// waf版本信息，spart-waf或者clb-waf，其他无效，请一定填写
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 规则执行的方式，TimedJob为定时执行，CronJob为周期执行
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 定时任务配置
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`
}

func (r *AddAreaBanAreasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAreaBanAreasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Areas")
	delete(f, "Edition")
	delete(f, "JobType")
	delete(f, "JobDateTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddAreaBanAreasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAreaBanAreasResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddAreaBanAreasResponse struct {
	*tchttp.BaseResponse
	Response *AddAreaBanAreasResponseParams `json:"Response"`
}

func (r *AddAreaBanAreasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAreaBanAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAttackWhiteRuleRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 匹配规则项列表
	Rules []*UserWhiteRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 规则序号
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则Id
	SignatureId *string `json:"SignatureId,omitnil,omitempty" name:"SignatureId"`

	// 加白的规则ID列表
	SignatureIds []*string `json:"SignatureIds,omitnil,omitempty" name:"SignatureIds"`

	// 加白的大类规则ID
	TypeIds []*string `json:"TypeIds,omitnil,omitempty" name:"TypeIds"`

	// 0:按照特定规则ID加白, 1:按照规则类型加白
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 规则名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type AddAttackWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 匹配规则项列表
	Rules []*UserWhiteRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 规则序号
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则Id
	SignatureId *string `json:"SignatureId,omitnil,omitempty" name:"SignatureId"`

	// 加白的规则ID列表
	SignatureIds []*string `json:"SignatureIds,omitnil,omitempty" name:"SignatureIds"`

	// 加白的大类规则ID
	TypeIds []*string `json:"TypeIds,omitnil,omitempty" name:"TypeIds"`

	// 0:按照特定规则ID加白, 1:按照规则类型加白
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 规则名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	delete(f, "Status")
	delete(f, "Rules")
	delete(f, "RuleId")
	delete(f, "SignatureId")
	delete(f, "SignatureIds")
	delete(f, "TypeIds")
	delete(f, "Mode")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddAttackWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAttackWhiteRuleResponseParams struct {
	// 规则总数
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 优先级
	SortId *string `json:"SortId,omitnil,omitempty" name:"SortId"`

	// 策略详情
	Strategies []*Strategy `json:"Strategies,omitnil,omitempty" name:"Strategies"`

	// 需要添加策略的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 动作类型，1代表阻断，2代表人机识别，3代表观察，4代表重定向，5代表JS校验
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 如果动作是重定向，则表示重定向的地址；其他情况可以为空
	Redirect *string `json:"Redirect,omitnil,omitempty" name:"Redirect"`

	// 过期时间，单位为秒级时间戳，例如1677254399表示过期时间为2023-02-24 23:59:59. 0表示永不过期
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 放行时是否继续执行其它检查逻辑，继续执行地域封禁防护：geoip、继续执行CC策略防护：cc、继续执行WEB应用防护：owasp、继续执行AI引擎防护：ai、继续执行信息防泄漏防护：antileakage。如果多个勾选那么以,串接。默认是"geoip,cc,owasp,ai,antileakage"
	Bypass *string `json:"Bypass,omitnil,omitempty" name:"Bypass"`

	// 添加规则的来源，默认为空
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 规则执行的方式，TimedJob为定时执行，CronJob为周期执行
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 规则执行的时间
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`

	// 规则来源，判断是不是小程序的
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 规则标签，小程序规则用，标识是内置规则还是自定义规则
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 开关状态，小程序风控规则的时候传该值
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 拦截页面id
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`
}

type AddCustomRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 优先级
	SortId *string `json:"SortId,omitnil,omitempty" name:"SortId"`

	// 策略详情
	Strategies []*Strategy `json:"Strategies,omitnil,omitempty" name:"Strategies"`

	// 需要添加策略的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 动作类型，1代表阻断，2代表人机识别，3代表观察，4代表重定向，5代表JS校验
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 如果动作是重定向，则表示重定向的地址；其他情况可以为空
	Redirect *string `json:"Redirect,omitnil,omitempty" name:"Redirect"`

	// 过期时间，单位为秒级时间戳，例如1677254399表示过期时间为2023-02-24 23:59:59. 0表示永不过期
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 放行时是否继续执行其它检查逻辑，继续执行地域封禁防护：geoip、继续执行CC策略防护：cc、继续执行WEB应用防护：owasp、继续执行AI引擎防护：ai、继续执行信息防泄漏防护：antileakage。如果多个勾选那么以,串接。默认是"geoip,cc,owasp,ai,antileakage"
	Bypass *string `json:"Bypass,omitnil,omitempty" name:"Bypass"`

	// 添加规则的来源，默认为空
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 规则执行的方式，TimedJob为定时执行，CronJob为周期执行
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 规则执行的时间
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`

	// 规则来源，判断是不是小程序的
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 规则标签，小程序规则用，标识是内置规则还是自定义规则
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 开关状态，小程序风控规则的时候传该值
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 拦截页面id
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`
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
	delete(f, "Strategies")
	delete(f, "Domain")
	delete(f, "ActionType")
	delete(f, "Redirect")
	delete(f, "ExpireTime")
	delete(f, "Edition")
	delete(f, "Bypass")
	delete(f, "EventId")
	delete(f, "JobType")
	delete(f, "JobDateTime")
	delete(f, "Source")
	delete(f, "Label")
	delete(f, "Status")
	delete(f, "PageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCustomRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCustomRuleResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	Success *ResponseCode `json:"Success,omitnil,omitempty" name:"Success"`

	// 添加成功的规则ID
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 优先级
	SortId *string `json:"SortId,omitnil,omitempty" name:"SortId"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 策略详情
	Strategies []*Strategy `json:"Strategies,omitnil,omitempty" name:"Strategies"`

	// 需要添加策略的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 放行的详情
	Bypass *string `json:"Bypass,omitnil,omitempty" name:"Bypass"`

	// 规则执行的方式，TimedJob为定时执行，CronJob为周期执行
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 定时任务配置
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`
}

type AddCustomWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 优先级
	SortId *string `json:"SortId,omitnil,omitempty" name:"SortId"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 策略详情
	Strategies []*Strategy `json:"Strategies,omitnil,omitempty" name:"Strategies"`

	// 需要添加策略的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 放行的详情
	Bypass *string `json:"Bypass,omitnil,omitempty" name:"Bypass"`

	// 规则执行的方式，TimedJob为定时执行，CronJob为周期执行
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 定时任务配置
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`
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
	delete(f, "JobType")
	delete(f, "JobDateTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCustomWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCustomWhiteRuleResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	Success *ResponseCode `json:"Success,omitnil,omitempty" name:"Success"`

	// 添加成功的规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 需要添加的规则
	Rules []*uint64 `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 需要添加的规则url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 规则的方法
	Function *string `json:"Function,omitnil,omitempty" name:"Function"`

	// 规则的开关，0表示规则关闭，1表示规则打开
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type AddDomainWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要添加的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 需要添加的规则
	Rules []*uint64 `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 需要添加的规则url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 规则的方法
	Function *string `json:"Function,omitnil,omitempty" name:"Function"`

	// 规则的开关，0表示规则关闭，1表示规则打开
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 证书类型。
	// 0：仅配置HTTP监听端口，没有证书
	// 1：证书来源为自有证书
	// 2：证书来源为托管证书
	CertType *int64 `json:"CertType,omitnil,omitempty" name:"CertType"`

	// waf前是否部署有七层代理服务。
	// 0：没有部署代理服务
	// 1：有部署代理服务，waf将使用XFF获取客户端IP
	// 2：有部署代理服务，waf将使用remote_addr获取客户端IP
	// 3：有部署代理服务，waf将使用ip_headers中的自定义header获取客户端IP
	IsCdn *int64 `json:"IsCdn,omitnil,omitempty" name:"IsCdn"`

	// 回源类型。
	// 0：通过IP回源
	// 1：通过域名回源
	UpstreamType *int64 `json:"UpstreamType,omitnil,omitempty" name:"UpstreamType"`

	// 是否开启WebSocket支持。
	// 0：关闭
	// 1：开启
	IsWebsocket *int64 `json:"IsWebsocket,omitnil,omitempty" name:"IsWebsocket"`

	// 回源负载均衡策略。
	// 0：轮询
	// 1：IP hash
	// 2：加权轮询
	LoadBalance *string `json:"LoadBalance,omitnil,omitempty" name:"LoadBalance"`

	// 服务端口列表配置。
	// NginxServerId：新增域名时填'0'
	// Port：监听端口号
	// Protocol：端口协议
	// UpstreamPort：与Port相同
	// UpstreamProtocol：与Protocol相同
	Ports []*PortItem `json:"Ports,omitnil,omitempty" name:"Ports"`

	// 必填项，是否开启长连接。
	// 0： 短连接
	// 1： 长连接
	IsKeepAlive *string `json:"IsKeepAlive,omitnil,omitempty" name:"IsKeepAlive"`

	// 必填项，域名所属实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// CertType为1时，需要填充此参数，表示自有证书的证书链
	Cert *string `json:"Cert,omitnil,omitempty" name:"Cert"`

	// CertType为1时，需要填充此参数，表示自有证书的私钥
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`

	// CertType为2时，需要填充此参数，表示腾讯云SSL平台托管的证书id
	SSLId *string `json:"SSLId,omitnil,omitempty" name:"SSLId"`

	// 待废弃，可不填。Waf的资源ID。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// IsCdn为3时，需要填此参数，表示自定义header
	IpHeaders []*string `json:"IpHeaders,omitnil,omitempty" name:"IpHeaders"`

	// 服务配置有HTTPS端口时，HTTPS的回源协议。
	// http：使用http协议回源，和HttpsUpstreamPort配合使用
	// https：使用https协议回源
	UpstreamScheme *string `json:"UpstreamScheme,omitnil,omitempty" name:"UpstreamScheme"`

	// HTTPS回源端口,仅UpstreamScheme为http时需要填当前字段
	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitnil,omitempty" name:"HttpsUpstreamPort"`

	// 待废弃，可不填。是否开启灰度，0表示不开启灰度。
	IsGray *int64 `json:"IsGray,omitnil,omitempty" name:"IsGray"`

	// 待废弃，可不填。灰度的地区
	GrayAreas []*string `json:"GrayAreas,omitnil,omitempty" name:"GrayAreas"`

	// 必填项，是否开启HTTP强制跳转到HTTPS。
	// 0：不强制跳转
	// 1：开启强制跳转
	HttpsRewrite *int64 `json:"HttpsRewrite,omitnil,omitempty" name:"HttpsRewrite"`

	// 域名回源时的回源域名。UpstreamType为1时，需要填充此字段
	UpstreamDomain *string `json:"UpstreamDomain,omitnil,omitempty" name:"UpstreamDomain"`

	// IP回源时的回源IP列表。UpstreamType为0时，需要填充此字段
	SrcList []*string `json:"SrcList,omitnil,omitempty" name:"SrcList"`

	// 必填项，是否开启HTTP2，需要开启HTTPS协议支持。
	// 0：关闭
	// 1：开启
	IsHttp2 *int64 `json:"IsHttp2,omitnil,omitempty" name:"IsHttp2"`

	// 待废弃，可不填。WAF实例类型。
	// sparta-waf：SAAS型WAF
	// clb-waf：负载均衡型WAF
	// cdn-waf：CDN上的Web防护能力
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 待废弃，目前填0即可。anycast IP类型开关： 0 普通IP 1 Anycast IP
	Anycast *int64 `json:"Anycast,omitnil,omitempty" name:"Anycast"`

	// 回源IP列表各IP的权重，和SrcList一一对应。当且仅当UpstreamType为0，并且SrcList有多个IP，并且LoadBalance为2时需要填写，否则填 []
	Weights []*int64 `json:"Weights,omitnil,omitempty" name:"Weights"`

	// 必填项，是否开启主动健康检测。
	// 0：不开启
	// 1：开启
	ActiveCheck *int64 `json:"ActiveCheck,omitnil,omitempty" name:"ActiveCheck"`

	// TLS版本信息
	TLSVersion *int64 `json:"TLSVersion,omitnil,omitempty" name:"TLSVersion"`

	// 必填项，加密套件模板。
	// 0：不支持选择，使用默认模板  
	// 1：通用型模板 
	// 2：安全型模板
	// 3：自定义模板
	CipherTemplate *int64 `json:"CipherTemplate,omitnil,omitempty" name:"CipherTemplate"`

	// 自定义的加密套件列表。CipherTemplate为3时需要填此字段，表示自定义的加密套件，值通过DescribeCiphersDetail接口获取。
	Ciphers []*int64 `json:"Ciphers,omitnil,omitempty" name:"Ciphers"`

	// WAF与源站的读超时时间，默认300s。
	ProxyReadTimeout *int64 `json:"ProxyReadTimeout,omitnil,omitempty" name:"ProxyReadTimeout"`

	// WAF与源站的写超时时间，默认300s。
	ProxySendTimeout *int64 `json:"ProxySendTimeout,omitnil,omitempty" name:"ProxySendTimeout"`

	// WAF回源时的SNI类型。
	// 0：关闭SNI，不配置client_hello中的server_name
	// 1：开启SNI，client_hello中的server_name为防护域名
	// 2：开启SNI，SNI为域名回源时的源站域名
	// 3：开启SNI，SNI为自定义域名
	SniType *int64 `json:"SniType,omitnil,omitempty" name:"SniType"`

	// SniType为3时，需要填此参数，表示自定义的SNI；
	SniHost *string `json:"SniHost,omitnil,omitempty" name:"SniHost"`

	// 是否开启XFF重置。0：关闭 1：开启
	XFFReset *int64 `json:"XFFReset,omitnil,omitempty" name:"XFFReset"`

	// 域名备注信息
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 自定义回源Host。默认为空字符串，表示使用防护域名作为回源Host。
	UpstreamHost *string `json:"UpstreamHost,omitnil,omitempty" name:"UpstreamHost"`

	// 是否开启缓存。 0：关闭 1：开启
	ProxyBuffer *int64 `json:"ProxyBuffer,omitnil,omitempty" name:"ProxyBuffer"`

	// 是否开启拨测。 0: 禁用拨测  1: 启用拨测。默认启用拨测
	ProbeStatus *int64 `json:"ProbeStatus,omitnil,omitempty" name:"ProbeStatus"`

	// 国密选项。0：不开启国密 1：在原有TLS选项的基础上追加支持国密 2：开启国密并仅支持国密客户端访问
	GmType *int64 `json:"GmType,omitnil,omitempty" name:"GmType"`

	// 国密证书类型。0：无国密证书 1：证书来源为自有国密证书 2：证书来源为托管国密证书
	GmCertType *int64 `json:"GmCertType,omitnil,omitempty" name:"GmCertType"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的证书链
	GmCert *string `json:"GmCert,omitnil,omitempty" name:"GmCert"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的私钥
	GmPrivateKey *string `json:"GmPrivateKey,omitnil,omitempty" name:"GmPrivateKey"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的加密证书
	GmEncCert *string `json:"GmEncCert,omitnil,omitempty" name:"GmEncCert"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的加密证书的私钥
	GmEncPrivateKey *string `json:"GmEncPrivateKey,omitnil,omitempty" name:"GmEncPrivateKey"`

	// GmCertType为2时，需要填充此参数，表示腾讯云SSL平台托管的证书id
	GmSSLId *string `json:"GmSSLId,omitnil,omitempty" name:"GmSSLId"`
}

type AddSpartaProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 需要防护的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 证书类型。
	// 0：仅配置HTTP监听端口，没有证书
	// 1：证书来源为自有证书
	// 2：证书来源为托管证书
	CertType *int64 `json:"CertType,omitnil,omitempty" name:"CertType"`

	// waf前是否部署有七层代理服务。
	// 0：没有部署代理服务
	// 1：有部署代理服务，waf将使用XFF获取客户端IP
	// 2：有部署代理服务，waf将使用remote_addr获取客户端IP
	// 3：有部署代理服务，waf将使用ip_headers中的自定义header获取客户端IP
	IsCdn *int64 `json:"IsCdn,omitnil,omitempty" name:"IsCdn"`

	// 回源类型。
	// 0：通过IP回源
	// 1：通过域名回源
	UpstreamType *int64 `json:"UpstreamType,omitnil,omitempty" name:"UpstreamType"`

	// 是否开启WebSocket支持。
	// 0：关闭
	// 1：开启
	IsWebsocket *int64 `json:"IsWebsocket,omitnil,omitempty" name:"IsWebsocket"`

	// 回源负载均衡策略。
	// 0：轮询
	// 1：IP hash
	// 2：加权轮询
	LoadBalance *string `json:"LoadBalance,omitnil,omitempty" name:"LoadBalance"`

	// 服务端口列表配置。
	// NginxServerId：新增域名时填'0'
	// Port：监听端口号
	// Protocol：端口协议
	// UpstreamPort：与Port相同
	// UpstreamProtocol：与Protocol相同
	Ports []*PortItem `json:"Ports,omitnil,omitempty" name:"Ports"`

	// 必填项，是否开启长连接。
	// 0： 短连接
	// 1： 长连接
	IsKeepAlive *string `json:"IsKeepAlive,omitnil,omitempty" name:"IsKeepAlive"`

	// 必填项，域名所属实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// CertType为1时，需要填充此参数，表示自有证书的证书链
	Cert *string `json:"Cert,omitnil,omitempty" name:"Cert"`

	// CertType为1时，需要填充此参数，表示自有证书的私钥
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`

	// CertType为2时，需要填充此参数，表示腾讯云SSL平台托管的证书id
	SSLId *string `json:"SSLId,omitnil,omitempty" name:"SSLId"`

	// 待废弃，可不填。Waf的资源ID。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// IsCdn为3时，需要填此参数，表示自定义header
	IpHeaders []*string `json:"IpHeaders,omitnil,omitempty" name:"IpHeaders"`

	// 服务配置有HTTPS端口时，HTTPS的回源协议。
	// http：使用http协议回源，和HttpsUpstreamPort配合使用
	// https：使用https协议回源
	UpstreamScheme *string `json:"UpstreamScheme,omitnil,omitempty" name:"UpstreamScheme"`

	// HTTPS回源端口,仅UpstreamScheme为http时需要填当前字段
	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitnil,omitempty" name:"HttpsUpstreamPort"`

	// 待废弃，可不填。是否开启灰度，0表示不开启灰度。
	IsGray *int64 `json:"IsGray,omitnil,omitempty" name:"IsGray"`

	// 待废弃，可不填。灰度的地区
	GrayAreas []*string `json:"GrayAreas,omitnil,omitempty" name:"GrayAreas"`

	// 必填项，是否开启HTTP强制跳转到HTTPS。
	// 0：不强制跳转
	// 1：开启强制跳转
	HttpsRewrite *int64 `json:"HttpsRewrite,omitnil,omitempty" name:"HttpsRewrite"`

	// 域名回源时的回源域名。UpstreamType为1时，需要填充此字段
	UpstreamDomain *string `json:"UpstreamDomain,omitnil,omitempty" name:"UpstreamDomain"`

	// IP回源时的回源IP列表。UpstreamType为0时，需要填充此字段
	SrcList []*string `json:"SrcList,omitnil,omitempty" name:"SrcList"`

	// 必填项，是否开启HTTP2，需要开启HTTPS协议支持。
	// 0：关闭
	// 1：开启
	IsHttp2 *int64 `json:"IsHttp2,omitnil,omitempty" name:"IsHttp2"`

	// 待废弃，可不填。WAF实例类型。
	// sparta-waf：SAAS型WAF
	// clb-waf：负载均衡型WAF
	// cdn-waf：CDN上的Web防护能力
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 待废弃，目前填0即可。anycast IP类型开关： 0 普通IP 1 Anycast IP
	Anycast *int64 `json:"Anycast,omitnil,omitempty" name:"Anycast"`

	// 回源IP列表各IP的权重，和SrcList一一对应。当且仅当UpstreamType为0，并且SrcList有多个IP，并且LoadBalance为2时需要填写，否则填 []
	Weights []*int64 `json:"Weights,omitnil,omitempty" name:"Weights"`

	// 必填项，是否开启主动健康检测。
	// 0：不开启
	// 1：开启
	ActiveCheck *int64 `json:"ActiveCheck,omitnil,omitempty" name:"ActiveCheck"`

	// TLS版本信息
	TLSVersion *int64 `json:"TLSVersion,omitnil,omitempty" name:"TLSVersion"`

	// 必填项，加密套件模板。
	// 0：不支持选择，使用默认模板  
	// 1：通用型模板 
	// 2：安全型模板
	// 3：自定义模板
	CipherTemplate *int64 `json:"CipherTemplate,omitnil,omitempty" name:"CipherTemplate"`

	// 自定义的加密套件列表。CipherTemplate为3时需要填此字段，表示自定义的加密套件，值通过DescribeCiphersDetail接口获取。
	Ciphers []*int64 `json:"Ciphers,omitnil,omitempty" name:"Ciphers"`

	// WAF与源站的读超时时间，默认300s。
	ProxyReadTimeout *int64 `json:"ProxyReadTimeout,omitnil,omitempty" name:"ProxyReadTimeout"`

	// WAF与源站的写超时时间，默认300s。
	ProxySendTimeout *int64 `json:"ProxySendTimeout,omitnil,omitempty" name:"ProxySendTimeout"`

	// WAF回源时的SNI类型。
	// 0：关闭SNI，不配置client_hello中的server_name
	// 1：开启SNI，client_hello中的server_name为防护域名
	// 2：开启SNI，SNI为域名回源时的源站域名
	// 3：开启SNI，SNI为自定义域名
	SniType *int64 `json:"SniType,omitnil,omitempty" name:"SniType"`

	// SniType为3时，需要填此参数，表示自定义的SNI；
	SniHost *string `json:"SniHost,omitnil,omitempty" name:"SniHost"`

	// 是否开启XFF重置。0：关闭 1：开启
	XFFReset *int64 `json:"XFFReset,omitnil,omitempty" name:"XFFReset"`

	// 域名备注信息
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 自定义回源Host。默认为空字符串，表示使用防护域名作为回源Host。
	UpstreamHost *string `json:"UpstreamHost,omitnil,omitempty" name:"UpstreamHost"`

	// 是否开启缓存。 0：关闭 1：开启
	ProxyBuffer *int64 `json:"ProxyBuffer,omitnil,omitempty" name:"ProxyBuffer"`

	// 是否开启拨测。 0: 禁用拨测  1: 启用拨测。默认启用拨测
	ProbeStatus *int64 `json:"ProbeStatus,omitnil,omitempty" name:"ProbeStatus"`

	// 国密选项。0：不开启国密 1：在原有TLS选项的基础上追加支持国密 2：开启国密并仅支持国密客户端访问
	GmType *int64 `json:"GmType,omitnil,omitempty" name:"GmType"`

	// 国密证书类型。0：无国密证书 1：证书来源为自有国密证书 2：证书来源为托管国密证书
	GmCertType *int64 `json:"GmCertType,omitnil,omitempty" name:"GmCertType"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的证书链
	GmCert *string `json:"GmCert,omitnil,omitempty" name:"GmCert"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的私钥
	GmPrivateKey *string `json:"GmPrivateKey,omitnil,omitempty" name:"GmPrivateKey"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的加密证书
	GmEncCert *string `json:"GmEncCert,omitnil,omitempty" name:"GmEncCert"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的加密证书的私钥
	GmEncPrivateKey *string `json:"GmEncPrivateKey,omitnil,omitempty" name:"GmEncPrivateKey"`

	// GmCertType为2时，需要填充此参数，表示腾讯云SSL平台托管的证书id
	GmSSLId *string `json:"GmSSLId,omitnil,omitempty" name:"GmSSLId"`
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
	delete(f, "Ports")
	delete(f, "IsKeepAlive")
	delete(f, "InstanceID")
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
	delete(f, "Edition")
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
	delete(f, "Note")
	delete(f, "UpstreamHost")
	delete(f, "ProxyBuffer")
	delete(f, "ProbeStatus")
	delete(f, "GmType")
	delete(f, "GmCertType")
	delete(f, "GmCert")
	delete(f, "GmPrivateKey")
	delete(f, "GmEncCert")
	delete(f, "GmEncPrivateKey")
	delete(f, "GmSSLId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddSpartaProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddSpartaProtectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type ApiAsset struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 请求方法
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// api名称
	ApiName *string `json:"ApiName,omitnil,omitempty" name:"ApiName"`

	// 场景
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 数据标签
	Label []*string `json:"Label,omitnil,omitempty" name:"Label"`

	// 过去7天是否活跃
	Active *bool `json:"Active,omitnil,omitempty" name:"Active"`

	// 最近更新时间
	Timestamp *uint64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// api发现时间
	InsertTime *uint64 `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// 资产状态，1:新发现，2，确认中，3，已确认，4，已下线，5，已忽略
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 风险等级，100,200,300对应低中高
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 近30天调用量
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否鉴权，1标识是，0表示否
	IsAuth *int64 `json:"IsAuth,omitnil,omitempty" name:"IsAuth"`

	// 如果添加了api入参检测规则，则此id返回值不为0
	ApiRequestRuleId *int64 `json:"ApiRequestRuleId,omitnil,omitempty" name:"ApiRequestRuleId"`

	// 如果添加了api限流规则，则此id返回值不为0
	ApiLimitRuleId *int64 `json:"ApiLimitRuleId,omitnil,omitempty" name:"ApiLimitRuleId"`

	// 对象接入和泛域名接入时，展示host列表
	HostList []*string `json:"HostList,omitnil,omitempty" name:"HostList"`
}

type ApiDataFilter struct {
	// 数据标签，是否活跃，功能场景
	Entity *string `json:"Entity,omitnil,omitempty" name:"Entity"`

	// 等于
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 日期，手机号，邮箱等
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ApiDetailSampleHistory struct {
	// 样例名称
	SampleNme *string `json:"SampleNme,omitnil,omitempty" name:"SampleNme"`

	// 请求样例
	RepLog *string `json:"RepLog,omitnil,omitempty" name:"RepLog"`

	// 响应样例
	RspLog *string `json:"RspLog,omitnil,omitempty" name:"RspLog"`
}

type ApiParameterType struct {
	// 参数名称
	ParameterName *string `json:"ParameterName,omitnil,omitempty" name:"ParameterName"`

	// 参数类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 参数位置
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 数据标签(敏感字段)
	Label []*string `json:"Label,omitnil,omitempty" name:"Label"`

	// 时间戳
	Timestamp *uint64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 来源是请求或者响应
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 是否需要泛化 ，0表示不需要，1表示需要
	IsPan *int64 `json:"IsPan,omitnil,omitempty" name:"IsPan"`

	// 是否鉴权，1表示是，0表示否
	IsAuth *int64 `json:"IsAuth,omitnil,omitempty" name:"IsAuth"`
}

type ApiPkg struct {
	// 资源id
	ResourceIds *string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 地域
	Region *int64 `json:"Region,omitnil,omitempty" name:"Region"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 申请数量
	InquireNum *int64 `json:"InquireNum,omitnil,omitempty" name:"InquireNum"`

	// 使用数量
	UsedNum *int64 `json:"UsedNum,omitnil,omitempty" name:"UsedNum"`

	// 续费标志
	RenewFlag *uint64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 计费项
	BillingItem *string `json:"BillingItem,omitnil,omitempty" name:"BillingItem"`

	// api安全7天试用标识。1试用。0没试用
	IsAPISecurityTrial *int64 `json:"IsAPISecurityTrial,omitnil,omitempty" name:"IsAPISecurityTrial"`
}

type ApiSecKey struct {
	// api名称
	ApiName *string `json:"ApiName,omitnil,omitempty" name:"ApiName"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 请求方法
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`
}

type AttackLogInfo struct {
	// 攻击日志的详情内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// CLS返回内容
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// CLS返回内容
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// CLS返回内容
	TimeStamp *string `json:"TimeStamp,omitnil,omitempty" name:"TimeStamp"`
}

type AutoDenyDetail struct {
	// 攻击封禁类型标签
	AttackTags []*string `json:"AttackTags,omitnil,omitempty" name:"AttackTags"`

	// 攻击次数阈值
	AttackThreshold *int64 `json:"AttackThreshold,omitnil,omitempty" name:"AttackThreshold"`

	// 自动封禁状态
	DefenseStatus *int64 `json:"DefenseStatus,omitnil,omitempty" name:"DefenseStatus"`

	// 攻击时间阈值
	TimeThreshold *int64 `json:"TimeThreshold,omitnil,omitempty" name:"TimeThreshold"`

	// 自动封禁时间
	DenyTimeThreshold *int64 `json:"DenyTimeThreshold,omitnil,omitempty" name:"DenyTimeThreshold"`

	// 最后更新时间
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`
}

type BatchIpAccessControlData struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 黑白名单条目
	Res []*BatchIpAccessControlItem `json:"Res,omitnil,omitempty" name:"Res"`
}

type BatchIpAccessControlItem struct {
	// mongo表自增Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 黑名单42或白名单40
	ActionType *int64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 黑白名单的IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 备注
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 添加路径
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 修改时间
	TsVersion *uint64 `json:"TsVersion,omitnil,omitempty" name:"TsVersion"`

	// 超时时间
	ValidTs *int64 `json:"ValidTs,omitnil,omitempty" name:"ValidTs"`

	// 域名列表
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// 55101145
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// IP列表
	IpList []*string `json:"IpList,omitnil,omitempty" name:"IpList"`

	// 创建时间
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 定时任务类型
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 周期任务类型
	CronType *string `json:"CronType,omitnil,omitempty" name:"CronType"`

	// 定时任务配置详情
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`

	// 生效状态
	ValidStatus *int64 `json:"ValidStatus,omitnil,omitempty" name:"ValidStatus"`
}

type BotPkg struct {
	// 资源id
	ResourceIds *string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 地域
	Region *int64 `json:"Region,omitnil,omitempty" name:"Region"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 申请数量
	InquireNum *int64 `json:"InquireNum,omitnil,omitempty" name:"InquireNum"`

	// 使用数量
	UsedNum *int64 `json:"UsedNum,omitnil,omitempty" name:"UsedNum"`

	// 子产品code
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 续费标志	
	RenewFlag *uint64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 购买页bot6折
	BotCPWaf *int64 `json:"BotCPWaf,omitnil,omitempty" name:"BotCPWaf"`

	// 控制台买bot5折
	BotNPWaf *int64 `json:"BotNPWaf,omitnil,omitempty" name:"BotNPWaf"`

	// 7天bot试用标识 1 试用 0 没有试用
	IsBotTrial *int64 `json:"IsBotTrial,omitnil,omitempty" name:"IsBotTrial"`
}

type BotQPS struct {
	// 资源id
	ResourceIds *string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 有效时间
	ValidTime *string `json:"ValidTime,omitnil,omitempty" name:"ValidTime"`

	// 资源数量
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 资源所在地区
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 使用qps的最大值
	MaxBotQPS *uint64 `json:"MaxBotQPS,omitnil,omitempty" name:"MaxBotQPS"`

	// 续费标志
	RenewFlag *uint64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type BotStatPointItem struct {
	// 横坐标
	TimeStamp *string `json:"TimeStamp,omitnil,omitempty" name:"TimeStamp"`

	// value的所属对象
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 纵列表
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`

	// Key对应的页面展示内容
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`
}

type CCRuleData struct {
	// cc规则
	Res []*CCRuleItem `json:"Res,omitnil,omitempty" name:"Res"`

	// 规则数目
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type CCRuleItem struct {
	// 动作，20表示观察，21表示人机识别，22表示拦截，23表示精准拦截，24表示JS校验
	ActionType *uint64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 高级模式
	Advance *uint64 `json:"Advance,omitnil,omitempty" name:"Advance"`

	// 时间周期
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 限制次数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 匹配方法
	MatchFunc *uint64 `json:"MatchFunc,omitnil,omitempty" name:"MatchFunc"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 优先级
	Priority *uint64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 更新时间戳
	TsVersion *uint64 `json:"TsVersion,omitnil,omitempty" name:"TsVersion"`

	// 匹配url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 策略动作有效时间
	ValidTime *uint64 `json:"ValidTime,omitnil,omitempty" name:"ValidTime"`

	// 高级参数
	OptionsArr *string `json:"OptionsArr,omitnil,omitempty" name:"OptionsArr"`

	// url长度
	Length *uint64 `json:"Length,omitnil,omitempty" name:"Length"`

	// 规则ID
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 事件id
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 关联的Session规则
	SessionApplied []*int64 `json:"SessionApplied,omitnil,omitempty" name:"SessionApplied"`

	// 创建时间
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type CCRuleItems struct {
	// 名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 模式
	Advance *uint64 `json:"Advance,omitnil,omitempty" name:"Advance"`

	// 限制
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 范围
	Interval *uint64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 网址
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 匹配类型
	MatchFunc *uint64 `json:"MatchFunc,omitnil,omitempty" name:"MatchFunc"`

	// 动作，20表示观察，21表示人机识别，22表示拦截，23表示精准拦截，24表示JS校验
	ActionType *uint64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 优先级
	Priority *uint64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 有效时间
	ValidTime *uint64 `json:"ValidTime,omitnil,omitempty" name:"ValidTime"`

	// 版本
	TsVersion *uint64 `json:"TsVersion,omitnil,omitempty" name:"TsVersion"`

	// 规则详情
	Options *string `json:"Options,omitnil,omitempty" name:"Options"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 事件id
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 关联的Session规则
	SessionApplied []*int64 `json:"SessionApplied,omitnil,omitempty" name:"SessionApplied"`

	// 创建时间
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type CCRuleLists struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 规则
	Res []*CCRuleItems `json:"Res,omitnil,omitempty" name:"Res"`
}

type CacheUrlItems struct {
	// 标识
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 网址
	Uri *string `json:"Uri,omitnil,omitempty" name:"Uri"`

	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type CdcCluster struct {
	// cdc的集群id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// cdc的集群名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type CdcRegion struct {
	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 该地域对应的集群信息
	Clusters []*CdcCluster `json:"Clusters,omitnil,omitempty" name:"Clusters"`
}

type ClbDomainsInfo struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名唯一ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 域名所属实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 域名所属实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 域名所属实例类型
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// waf前是否部署有七层代理服务。 0：没有部署代理服务 1：有部署代理服务，waf将使用XFF获取客户端IP 2：有部署代理服务，waf将使用remote_addr获取客户端IP 3：有部署代理服务，waf将使用ip_headers中的自定义header获取客户端IP
	IsCdn *uint64 `json:"IsCdn,omitnil,omitempty" name:"IsCdn"`

	// 负载均衡类型为clb时，对应的负载均衡器信息
	LoadBalancerSet []*LoadBalancerPackageNew `json:"LoadBalancerSet,omitnil,omitempty" name:"LoadBalancerSet"`

	// 负载均衡型WAF的流量模式，1：清洗模式，0：镜像模式
	FlowMode *uint64 `json:"FlowMode,omitnil,omitempty" name:"FlowMode"`

	// 域名绑定负载均衡器状态
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 负载均衡类型，clb或者apisix
	AlbType *string `json:"AlbType,omitnil,omitempty" name:"AlbType"`

	// IsCdn=3时，表示自定义header
	IpHeaders []*string `json:"IpHeaders,omitnil,omitempty" name:"IpHeaders"`

	// cdc-clb-waf类型WAF的CDC集群信息
	CdcClusters *string `json:"CdcClusters,omitnil,omitempty" name:"CdcClusters"`

	// 云类型:public:公有云；private:私有云;hybrid:混合云
	CloudType *string `json:"CloudType,omitnil,omitempty" name:"CloudType"`

	// 域名备注信息
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 域名标签
	Labels []*string `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type ClbObject struct {
	// 对象ID
	ObjectId *string `json:"ObjectId,omitnil,omitempty" name:"ObjectId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 精准域名列表
	PreciseDomains []*string `json:"PreciseDomains,omitnil,omitempty" name:"PreciseDomains"`

	// WAF功能开关状态，0关闭1开启
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// WAF日志开关状态，0关闭1开启
	ClsStatus *int64 `json:"ClsStatus,omitnil,omitempty" name:"ClsStatus"`

	// CLB对象对应的虚拟域名
	VirtualDomain *string `json:"VirtualDomain,omitnil,omitempty" name:"VirtualDomain"`

	// 对象名称
	ObjectName *string `json:"ObjectName,omitnil,omitempty" name:"ObjectName"`

	// 公网地址
	PublicIp []*string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 内网地址
	PrivateIp []*string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// VPC名称
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// VPC ID
	Vpc *string `json:"Vpc,omitnil,omitempty" name:"Vpc"`

	// waf实例等级，如果未绑定实例为0
	InstanceLevel *int64 `json:"InstanceLevel,omitnil,omitempty" name:"InstanceLevel"`

	// clb投递开关
	PostCLSStatus *int64 `json:"PostCLSStatus,omitnil,omitempty" name:"PostCLSStatus"`

	// kafka投递开关
	PostCKafkaStatus *int64 `json:"PostCKafkaStatus,omitnil,omitempty" name:"PostCKafkaStatus"`

	// 对象类型：CLB:负载均衡器，TSE:云原生网关
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 对象地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 代理状态: 0:不开启,1:以XFF的第一个IP地址作为客户端IP,2:以remote_addr作为客户端IP,3:从指定的头部字段获取客户端IP，字段通过IpHeaders字段给出
	Proxy *uint64 `json:"Proxy,omitnil,omitempty" name:"Proxy"`

	// 指定获取客户端IP的头部字段列表。IsCdn为3时有效
	IpHeaders []*string `json:"IpHeaders,omitnil,omitempty" name:"IpHeaders"`

	// bot防护开关
	BotStatus *int64 `json:"BotStatus,omitnil,omitempty" name:"BotStatus"`

	// api防护开关
	ApiStatus *int64 `json:"ApiStatus,omitnil,omitempty" name:"ApiStatus"`

	// 对象接入模式，0表示镜像模式，1表示清洗模式，2表示体检模式，默认为清洗模式
	ObjectFlowMode *int64 `json:"ObjectFlowMode,omitnil,omitempty" name:"ObjectFlowMode"`

	// 数值形式的私有网络 ID
	NumericalVpcId *int64 `json:"NumericalVpcId,omitnil,omitempty" name:"NumericalVpcId"`
}

type ClbWafRegionItem struct {
	// 地域ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 地域中文说明
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 地域英文全拼
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 地域编码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`
}

// Predefined struct for user
type CreateAccessExportRequestParams struct {
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 日志导出检索语句
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 日志导出数量，最大值100w
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 日志导出数据格式。json，csv，默认为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 日志导出时间排序。desc，asc，默认为desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type CreateAccessExportRequest struct {
	*tchttp.BaseRequest
	
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 日志导出检索语句
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 日志导出数量，最大值100w
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 日志导出数据格式。json，csv，默认为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 日志导出时间排序。desc，asc，默认为desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
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
	ExportId *string `json:"ExportId,omitnil,omitempty" name:"ExportId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type CreateDealsGoods struct {
	// 商品数量
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 商品明细
	GoodsDetail *CreateDealsGoodsDetail `json:"GoodsDetail,omitnil,omitempty" name:"GoodsDetail"`

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
	GoodsCategoryId *int64 `json:"GoodsCategoryId,omitnil,omitempty" name:"GoodsCategoryId"`

	// 购买waf实例区域ID
	// 1 表示购买大陆资源;
	// 9表示购买非中国大陆资源
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`
}

type CreateDealsGoodsDetail struct {
	// 时间间隔
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 单位，支持购买d、m、y 即（日、月、年）
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

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
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

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
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// waf实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 1:自动续费，0:不自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// waf购买的实际地域信息
	RealRegion *int64 `json:"RealRegion,omitnil,omitempty" name:"RealRegion"`

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
	LabelTypes []*string `json:"LabelTypes,omitnil,omitempty" name:"LabelTypes"`

	// 计费细项标签数量，一般和SvLabelType一一对应
	LabelCounts []*int64 `json:"LabelCounts,omitnil,omitempty" name:"LabelCounts"`

	// 变配使用，实例到期时间
	CurDeadline *string `json:"CurDeadline,omitnil,omitempty" name:"CurDeadline"`

	// 对存在的实例购买bot 或api 安全
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

// Predefined struct for user
type CreateDealsRequestParams struct {
	// 计费下单入参
	Goods []*CreateDealsGoods `json:"Goods,omitnil,omitempty" name:"Goods"`
}

type CreateDealsRequest struct {
	*tchttp.BaseRequest
	
	// 计费下单入参
	Goods []*CreateDealsGoods `json:"Goods,omitnil,omitempty" name:"Goods"`
}

func (r *CreateDealsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDealsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Goods")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDealsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDealsResponseParams struct {
	// 计费下单响应结构体
	Data *DealData `json:"Data,omitnil,omitempty" name:"Data"`

	// 1:成功，0:失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 返回message
	ReturnMessage *string `json:"ReturnMessage,omitnil,omitempty" name:"ReturnMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDealsResponse struct {
	*tchttp.BaseResponse
	Response *CreateDealsResponseParams `json:"Response"`
}

func (r *CreateDealsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDealsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHostRequestParams struct {
	// 防护域名配置信息
	Host *HostRecord `json:"Host,omitnil,omitempty" name:"Host"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type CreateHostRequest struct {
	*tchttp.BaseRequest
	
	// 防护域名配置信息
	Host *HostRecord `json:"Host,omitnil,omitempty" name:"Host"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
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
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CreateIpAccessControlRequestParams struct {
	// 具体域名如：test.qcloudwaf.com
	// 全局域名为：global
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// ip参数列表
	IpList []*string `json:"IpList,omitnil,omitempty" name:"IpList"`

	// 42为黑名单，40为白名单
	ActionType *int64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// valid_ts为有效日期，值为秒级时间戳（（如1680570420代表2023-04-04 09:07:00））
	//
	// Deprecated: ValidTS is deprecated.
	ValidTS *int64 `json:"ValidTS,omitnil,omitempty" name:"ValidTS"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 可选值为：batch（批量添加）、bot（BOT流量分析中的BOT详情列表中添加时）、cc（在攻击日志列表中对攻击类型为CC的IP添加时）、custom（非批量添加时的默认值）
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 备注
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 规则执行的方式，TimedJob为定时执行，CronJob为周期执行
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 定时配置详情
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`
}

type CreateIpAccessControlRequest struct {
	*tchttp.BaseRequest
	
	// 具体域名如：test.qcloudwaf.com
	// 全局域名为：global
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// ip参数列表
	IpList []*string `json:"IpList,omitnil,omitempty" name:"IpList"`

	// 42为黑名单，40为白名单
	ActionType *int64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// valid_ts为有效日期，值为秒级时间戳（（如1680570420代表2023-04-04 09:07:00））
	ValidTS *int64 `json:"ValidTS,omitnil,omitempty" name:"ValidTS"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 可选值为：batch（批量添加）、bot（BOT流量分析中的BOT详情列表中添加时）、cc（在攻击日志列表中对攻击类型为CC的IP添加时）、custom（非批量添加时的默认值）
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 备注
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 规则执行的方式，TimedJob为定时执行，CronJob为周期执行
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 定时配置详情
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`
}

func (r *CreateIpAccessControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIpAccessControlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "IpList")
	delete(f, "ActionType")
	delete(f, "ValidTS")
	delete(f, "InstanceId")
	delete(f, "Edition")
	delete(f, "SourceType")
	delete(f, "Note")
	delete(f, "JobType")
	delete(f, "JobDateTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIpAccessControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIpAccessControlResponseParams struct {
	// 新增的规则对应的ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIpAccessControlResponse struct {
	*tchttp.BaseResponse
	Response *CreateIpAccessControlResponseParams `json:"Response"`
}

func (r *CreateIpAccessControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIpAccessControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CronJob struct {
	// 每个月的几号执行
	Days []*uint64 `json:"Days,omitnil,omitempty" name:"Days"`

	// 每个星期的星期几执行
	WDays []*uint64 `json:"WDays,omitnil,omitempty" name:"WDays"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DealData struct {
	// 订单号列表，元素个数与请求包的goods数组的元素个数一致，商品详情与订单按顺序对应
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 大订单号，一个大订单号下可以有多个子订单，说明是同一次下单[{},{}]
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`
}

// Predefined struct for user
type DeleteAccessExportRequestParams struct {
	// 日志导出ID
	ExportId *string `json:"ExportId,omitnil,omitempty" name:"ExportId"`

	// 日志主题
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
}

type DeleteAccessExportRequest struct {
	*tchttp.BaseRequest
	
	// 日志导出ID
	ExportId *string `json:"ExportId,omitnil,omitempty" name:"ExportId"`

	// 日志主题
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteAntiFakeUrlRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则id
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type DeleteAntiInfoLeakRuleRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则id
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteAttackDownloadRecordRequest struct {
	*tchttp.BaseRequest
	
	// 下载任务记录唯一标记
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Ids []*uint64 `json:"Ids,omitnil,omitempty" name:"Ids"`

	// 用户域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type DeleteAttackWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则序号组
	Ids []*uint64 `json:"Ids,omitnil,omitempty" name:"Ids"`

	// 用户域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
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
	FailIds []*uint64 `json:"FailIds,omitnil,omitempty" name:"FailIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// clb-waf或者sparta-waf
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 规则Id
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type DeleteCCRuleRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// clb-waf或者sparta-waf
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 规则Id
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
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
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 操作的规则Id
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 删除的规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 批量删除的规则列表
	DomainRuleIdList []*DomainRuleId `json:"DomainRuleIdList,omitnil,omitempty" name:"DomainRuleIdList"`
}

type DeleteCustomRuleRequest struct {
	*tchttp.BaseRequest
	
	// 删除的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 删除的规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 批量删除的规则列表
	DomainRuleIdList []*DomainRuleId `json:"DomainRuleIdList,omitnil,omitempty" name:"DomainRuleIdList"`
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
	delete(f, "DomainRuleIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomRuleResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	Success *ResponseCode `json:"Success,omitnil,omitempty" name:"Success"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 删除的规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type DeleteCustomWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 删除的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 删除的规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
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
	Success *ResponseCode `json:"Success,omitnil,omitempty" name:"Success"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 需要删除的白名单规则
	Ids []*uint64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

type DeleteDomainWhiteRulesRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的规则域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 需要删除的白名单规则
	Ids []*uint64 `json:"Ids,omitnil,omitempty" name:"Ids"`
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
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteHostRequestParams struct {
	// 删除的域名列表
	HostsDel []*HostDel `json:"HostsDel,omitnil,omitempty" name:"HostsDel"`
}

type DeleteHostRequest struct {
	*tchttp.BaseRequest
	
	// 删除的域名列表
	HostsDel []*HostDel `json:"HostsDel,omitnil,omitempty" name:"HostsDel"`
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
	// 域名删除结果。Code表示状态码，Message表示详细信息。
	Success *ResponseCode `json:"Success,omitnil,omitempty" name:"Success"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 删除的ip数组
	Items []*string `json:"Items,omitnil,omitempty" name:"Items"`

	// 若IsId字段为True，则Items列表元素需为Id，否则为IP
	IsId *bool `json:"IsId,omitnil,omitempty" name:"IsId"`

	// 是否删除对应的域名下的所有黑/白IP名单，true表示全部删除，false表示只删除指定ip名单
	DeleteAll *bool `json:"DeleteAll,omitnil,omitempty" name:"DeleteAll"`

	// 是否为多域名黑白名单
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// IP黑白名单类型，40为IP白名单，42为IP黑名单
	ActionType *uint64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`
}

type DeleteIpAccessControlRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 删除的ip数组
	Items []*string `json:"Items,omitnil,omitempty" name:"Items"`

	// 若IsId字段为True，则Items列表元素需为Id，否则为IP
	IsId *bool `json:"IsId,omitnil,omitempty" name:"IsId"`

	// 是否删除对应的域名下的所有黑/白IP名单，true表示全部删除，false表示只删除指定ip名单
	DeleteAll *bool `json:"DeleteAll,omitnil,omitempty" name:"DeleteAll"`

	// 是否为多域名黑白名单
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// IP黑白名单类型，40为IP白名单，42为IP黑名单
	ActionType *uint64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`
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
	delete(f, "ActionType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIpAccessControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIpAccessControlResponseParams struct {
	// 删除失败的条目
	FailedItems *string `json:"FailedItems,omitnil,omitempty" name:"FailedItems"`

	// 删除失败的条目数
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteIpAccessControlV2RequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则ID列表，支持批量删除，在DeleteAll参数为true的时候可以不传
	RuleIds []*uint64 `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// 是否删除对应的域名下的所有黑/白IP名单，true表示全部删除，false表示只删除指定IP名单，批量防护不支持
	DeleteAll *bool `json:"DeleteAll,omitnil,omitempty" name:"DeleteAll"`

	// batch表示为批量防护的IP黑白名单
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// IP黑白名单类型，40为IP白名单，42为IP黑名单，在DeleteAll为true的时候必传此参数
	ActionType *uint64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`
}

type DeleteIpAccessControlV2Request struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则ID列表，支持批量删除，在DeleteAll参数为true的时候可以不传
	RuleIds []*uint64 `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// 是否删除对应的域名下的所有黑/白IP名单，true表示全部删除，false表示只删除指定IP名单，批量防护不支持
	DeleteAll *bool `json:"DeleteAll,omitnil,omitempty" name:"DeleteAll"`

	// batch表示为批量防护的IP黑白名单
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// IP黑白名单类型，40为IP白名单，42为IP黑名单，在DeleteAll为true的时候必传此参数
	ActionType *uint64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`
}

func (r *DeleteIpAccessControlV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIpAccessControlV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "RuleIds")
	delete(f, "DeleteAll")
	delete(f, "SourceType")
	delete(f, "ActionType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIpAccessControlV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIpAccessControlV2ResponseParams struct {
	// 在批量删除的时候表示删除失败的条数
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteIpAccessControlV2Response struct {
	*tchttp.BaseResponse
	Response *DeleteIpAccessControlV2ResponseParams `json:"Response"`
}

func (r *DeleteIpAccessControlV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIpAccessControlV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSessionRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// clb-waf 或者 sprta-waf
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 要删除的SessionID
	SessionID *int64 `json:"SessionID,omitnil,omitempty" name:"SessionID"`
}

type DeleteSessionRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// clb-waf 或者 sprta-waf
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 要删除的SessionID
	SessionID *int64 `json:"SessionID,omitnil,omitempty" name:"SessionID"`
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
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 实例类型
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 必填项。域名所属实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type DeleteSpartaProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 域名列表
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 实例类型
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 必填项。域名所属实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAccessExportsRequest struct {
	*tchttp.BaseRequest
	
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 日志导出列表
	Exports []*ExportAccessInfo `json:"Exports,omitnil,omitempty" name:"Exports"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 查询语句，语句长度最大为4096，由于本接口是分析接口，如果无过滤条件，必须传 * 表示匹配所有，参考CLS的分析统计语句的文档
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 需要分析统计的字段名
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// 客户要查询的日志主题ID，每个客户都有对应的一个主题
	//
	// Deprecated: TopicId is deprecated.
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 排序字段,升序asc,降序desc，默认降序desc 
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 返回的top数，默认返回top5
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type DescribeAccessFastAnalysisRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 查询语句，语句长度最大为4096，由于本接口是分析接口，如果无过滤条件，必须传 * 表示匹配所有，参考CLS的分析统计语句的文档
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 需要分析统计的字段名
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// 客户要查询的日志主题ID，每个客户都有对应的一个主题
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 排序字段,升序asc,降序desc，默认降序desc 
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 返回的top数，默认返回top5
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
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
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "FieldName")
	delete(f, "TopicId")
	delete(f, "Sort")
	delete(f, "Count")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccessFastAnalysisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccessFastAnalysisResponseParams struct {
	// 注意：此字段可能返回 null，表示取不到有效值
	FieldValueRatioInfos []*AccessFieldValueRatioInfo `json:"FieldValueRatioInfos,omitnil,omitempty" name:"FieldValueRatioInfos"`

	// 日志条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 查询语句，语句长度最大为4096
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 柱状图间隔时间差，单位ms
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`
}

type DescribeAccessHistogramRequest struct {
	*tchttp.BaseRequest
	
	// 老版本查询的日志主题ID，新版本传空字符串即可
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 查询语句，语句长度最大为4096
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 柱状图间隔时间差，单位ms
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`
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
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 满足条件的日志条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 注意：此字段可能返回 null，表示取不到有效值
	HistogramInfos []*AccessHistogramItem `json:"HistogramInfos,omitnil,omitempty" name:"HistogramInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 索引配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rule *AccessRuleInfo `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 索引修改时间，初始值为索引创建时间。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤数组,name可以是如下的值： RuleID,ParamName,Url,Action,Method,Source,Status
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`

	// asc或者desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 目前支持根据ts排序
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeAntiFakeRulesRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤数组,name可以是如下的值： RuleID,ParamName,Url,Action,Method,Source,Status
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`

	// asc或者desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 目前支持根据ts排序
	By *string `json:"By,omitnil,omitempty" name:"By"`
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
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 返回值
	Data []*CacheUrlItems `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type DescribeAntiInfoLeakRulesStrategyItem struct {
	// 字段
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 条件
	CompareFunc *string `json:"CompareFunc,omitnil,omitempty" name:"CompareFunc"`

	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

// Predefined struct for user
type DescribeAntiInfoLeakageRulesRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 翻页支持，读取偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 翻页支持，读取长度限制
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方式，asc或者desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 过滤器,可以允许如下的值：
	// RuleId,Match_field,Name,Action,Status
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAntiInfoLeakageRulesRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 翻页支持，读取偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 翻页支持，读取长度限制
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序方式，asc或者desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 过滤器,可以允许如下的值：
	// RuleId,Match_field,Name,Action,Status
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 规则列表
	RuleList []*DescribeAntiLeakageItem `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 状态值
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 动作
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 匹配条件
	Strategies []*DescribeAntiInfoLeakRulesStrategyItem `json:"Strategies,omitnil,omitempty" name:"Strategies"`

	// 匹配的URL
	Uri *string `json:"Uri,omitnil,omitempty" name:"Uri"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`
}

// Predefined struct for user
type DescribeApiDetailRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Api名称
	ApiName *string `json:"ApiName,omitnil,omitempty" name:"ApiName"`

	// 请求方法
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 是否仅查询敏感的，1表示查询，0表示不查询
	IsSensitive *int64 `json:"IsSensitive,omitnil,omitempty" name:"IsSensitive"`

	// 是否仅查询泛化的，1表示查询，0表示不查询
	IsPan *int64 `json:"IsPan,omitnil,omitempty" name:"IsPan"`
}

type DescribeApiDetailRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Api名称
	ApiName *string `json:"ApiName,omitnil,omitempty" name:"ApiName"`

	// 请求方法
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 是否仅查询敏感的，1表示查询，0表示不查询
	IsSensitive *int64 `json:"IsSensitive,omitnil,omitempty" name:"IsSensitive"`

	// 是否仅查询泛化的，1表示查询，0表示不查询
	IsPan *int64 `json:"IsPan,omitnil,omitempty" name:"IsPan"`
}

func (r *DescribeApiDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "ApiName")
	delete(f, "Method")
	delete(f, "IsSensitive")
	delete(f, "IsPan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiDetailResponseParams struct {
	// 请求样例，json字符串格式
	Log *string `json:"Log,omitnil,omitempty" name:"Log"`

	// 请求参数样例列表
	ParameterList []*ApiParameterType `json:"ParameterList,omitnil,omitempty" name:"ParameterList"`

	// 当前场景标签
	Scene *string `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 敏感字段
	SensitiveFields []*string `json:"SensitiveFields,omitnil,omitempty" name:"SensitiveFields"`

	// 7天内是否活跃
	IsActive *bool `json:"IsActive,omitnil,omitempty" name:"IsActive"`

	// 访问ip数
	IpCount *int64 `json:"IpCount,omitnil,omitempty" name:"IpCount"`

	// 访问地域数量
	RegionCount *int64 `json:"RegionCount,omitnil,omitempty" name:"RegionCount"`

	// 关联事件数
	EventCount *int64 `json:"EventCount,omitnil,omitempty" name:"EventCount"`

	// 涉敏数据条数
	SensitiveCount *uint64 `json:"SensitiveCount,omitnil,omitempty" name:"SensitiveCount"`

	// 风险等级
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 响应体
	RspLog *string `json:"RspLog,omitnil,omitempty" name:"RspLog"`

	// 昨日访问峰值QPS
	MaxQPS *uint64 `json:"MaxQPS,omitnil,omitempty" name:"MaxQPS"`

	// 历史样例
	ApiDetailSampleHistory []*ApiDetailSampleHistory `json:"ApiDetailSampleHistory,omitnil,omitempty" name:"ApiDetailSampleHistory"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApiDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiDetailResponseParams `json:"Response"`
}

func (r *DescribeApiDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiListVersionTwoRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 页面索引，第几页
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤条件
	Filters []*ApiDataFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序方法，1 升序，-1 降序
	Sort []*string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 是否进行总数查询
	NeedTotalCount *bool `json:"NeedTotalCount,omitnil,omitempty" name:"NeedTotalCount"`

	// 查询开始时间
	StartTs *int64 `json:"StartTs,omitnil,omitempty" name:"StartTs"`

	// 查询结束时间
	EndTs *int64 `json:"EndTs,omitnil,omitempty" name:"EndTs"`
}

type DescribeApiListVersionTwoRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 页面索引，第几页
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤条件
	Filters []*ApiDataFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序方法，1 升序，-1 降序
	Sort []*string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 是否进行总数查询
	NeedTotalCount *bool `json:"NeedTotalCount,omitnil,omitempty" name:"NeedTotalCount"`

	// 查询开始时间
	StartTs *int64 `json:"StartTs,omitnil,omitempty" name:"StartTs"`

	// 查询结束时间
	EndTs *int64 `json:"EndTs,omitnil,omitempty" name:"EndTs"`
}

func (r *DescribeApiListVersionTwoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiListVersionTwoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "PageIndex")
	delete(f, "PageSize")
	delete(f, "Filters")
	delete(f, "Sort")
	delete(f, "NeedTotalCount")
	delete(f, "StartTs")
	delete(f, "EndTs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiListVersionTwoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiListVersionTwoResponseParams struct {
	// api资产列表
	Data []*ApiAsset `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApiListVersionTwoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiListVersionTwoResponseParams `json:"Response"`
}

func (r *DescribeApiListVersionTwoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiListVersionTwoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAreaBanAreasRequestParams struct {
	// 需要查询的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type DescribeAreaBanAreasRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

func (r *DescribeAreaBanAreasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAreaBanAreasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAreaBanAreasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAreaBanAreasResponseParams struct {
	// 回包内容
	Data *DescribeAreaBanAreasRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAreaBanAreasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAreaBanAreasResponseParams `json:"Response"`
}

func (r *DescribeAreaBanAreasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAreaBanAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAreaBanAreasRsp struct {
	// 状态 "0"：未开启地域封禁 "1"：开启地域封禁
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 数据来源 custom-自定义(默认)、batch-批量防护
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 字符串数据，配置的地域列表
	Areas []*string `json:"Areas,omitnil,omitempty" name:"Areas"`

	// 定时任务类型
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 定时任务详细配置
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`

	// 周期任务配置
	CronType *string `json:"CronType,omitnil,omitempty" name:"CronType"`
}

// Predefined struct for user
type DescribeAreaBanSupportAreasRequestParams struct {

}

type DescribeAreaBanSupportAreasRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAreaBanSupportAreasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAreaBanSupportAreasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAreaBanSupportAreasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAreaBanSupportAreasResponseParams struct {
	// 地域封禁的地域列表，要解析成json后使用
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAreaBanSupportAreasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAreaBanSupportAreasResponseParams `json:"Response"`
}

func (r *DescribeAreaBanSupportAreasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAreaBanSupportAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttackOverviewRequestParams struct {
	// 查询开始时间
	FromTime *string `json:"FromTime,omitnil,omitempty" name:"FromTime"`

	// 查询结束时间
	ToTime *string `json:"ToTime,omitnil,omitempty" name:"ToTime"`

	// 客户的Appid
	Appid *uint64 `json:"Appid,omitnil,omitempty" name:"Appid"`

	// 被查询的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type DescribeAttackOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 查询开始时间
	FromTime *string `json:"FromTime,omitnil,omitempty" name:"FromTime"`

	// 查询结束时间
	ToTime *string `json:"ToTime,omitnil,omitempty" name:"ToTime"`

	// 客户的Appid
	Appid *uint64 `json:"Appid,omitnil,omitempty" name:"Appid"`

	// 被查询的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
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
	AccessCount *uint64 `json:"AccessCount,omitnil,omitempty" name:"AccessCount"`

	// Web攻击总数
	AttackCount *uint64 `json:"AttackCount,omitnil,omitempty" name:"AttackCount"`

	// 访问控制总数
	ACLCount *uint64 `json:"ACLCount,omitnil,omitempty" name:"ACLCount"`

	// CC攻击总数
	CCCount *uint64 `json:"CCCount,omitnil,omitempty" name:"CCCount"`

	// Bot攻击总数
	BotCount *uint64 `json:"BotCount,omitnil,omitempty" name:"BotCount"`

	// api资产总数
	ApiAssetsCount *uint64 `json:"ApiAssetsCount,omitnil,omitempty" name:"ApiAssetsCount"`

	// api风险事件数量
	ApiRiskEventCount *uint64 `json:"ApiRiskEventCount,omitnil,omitempty" name:"ApiRiskEventCount"`

	// 黑名单总数
	IPBlackCount *uint64 `json:"IPBlackCount,omitnil,omitempty" name:"IPBlackCount"`

	// 防篡改总数
	TamperCount *uint64 `json:"TamperCount,omitnil,omitempty" name:"TamperCount"`

	// 信息泄露总数
	LeakCount *uint64 `json:"LeakCount,omitnil,omitempty" name:"LeakCount"`

	// API风险事件周环比
	ApiRiskEventCircleCount *int64 `json:"ApiRiskEventCircleCount,omitnil,omitempty" name:"ApiRiskEventCircleCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	FromTime *string `json:"FromTime,omitnil,omitempty" name:"FromTime"`

	// 结束时间
	ToTime *string `json:"ToTime,omitnil,omitempty" name:"ToTime"`

	// 兼容Host，逐步淘汰Host字段
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 域名过滤，不传则不过滤，用于替代Host字段，逐步淘汰Host
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type DescribeAttackTypeRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间
	FromTime *string `json:"FromTime,omitnil,omitempty" name:"FromTime"`

	// 结束时间
	ToTime *string `json:"ToTime,omitnil,omitempty" name:"ToTime"`

	// 兼容Host，逐步淘汰Host字段
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 域名过滤，不传则不过滤，用于替代Host字段，逐步淘汰Host
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
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
	Piechart []*PiechartItem `json:"Piechart,omitnil,omitempty" name:"Piechart"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页容量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序的字段，支持user_id, signature_id, modify_time
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 筛选条件，支持SignatureId, MatchContent
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAttackWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页容量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序的字段，支持user_id, signature_id, modify_time
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 筛选条件，支持SignatureId, MatchContent
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 规则白名单列表
	List []*UserWhiteRule `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 查询IP自动封禁状态
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 计数标识
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 类别
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 有效时间最小时间戳
	VtsMin *uint64 `json:"VtsMin,omitnil,omitempty" name:"VtsMin"`

	// 有效时间最大时间戳
	VtsMax *uint64 `json:"VtsMax,omitnil,omitempty" name:"VtsMax"`

	// 创建时间最小时间戳
	CtsMin *uint64 `json:"CtsMin,omitnil,omitempty" name:"CtsMin"`

	// 创建时间最大时间戳
	CtsMax *uint64 `json:"CtsMax,omitnil,omitempty" name:"CtsMax"`

	// 偏移量
	Skip *uint64 `json:"Skip,omitnil,omitempty" name:"Skip"`

	// 限制条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 策略名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 排序参数
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`
}

type DescribeAutoDenyIPRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 查询IP自动封禁状态
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 计数标识
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 类别
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 有效时间最小时间戳
	VtsMin *uint64 `json:"VtsMin,omitnil,omitempty" name:"VtsMin"`

	// 有效时间最大时间戳
	VtsMax *uint64 `json:"VtsMax,omitnil,omitempty" name:"VtsMax"`

	// 创建时间最小时间戳
	CtsMin *uint64 `json:"CtsMin,omitnil,omitempty" name:"CtsMin"`

	// 创建时间最大时间戳
	CtsMax *uint64 `json:"CtsMax,omitnil,omitempty" name:"CtsMax"`

	// 偏移量
	Skip *uint64 `json:"Skip,omitnil,omitempty" name:"Skip"`

	// 限制条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 策略名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 排序参数
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`
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
	Data *IpHitItemsData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 筛选条件，支持 ActionType，可选的值为40（白名单）42（黑名单），ValidStatus，可选的值为1（生效）0（过期）
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移
	OffSet *uint64 `json:"OffSet,omitnil,omitempty" name:"OffSet"`

	// 限制
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序参数
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`
}

type DescribeBatchIpAccessControlRequest struct {
	*tchttp.BaseRequest
	
	// 筛选条件，支持 ActionType，可选的值为40（白名单）42（黑名单），ValidStatus，可选的值为1（生效）0（过期）
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移
	OffSet *uint64 `json:"OffSet,omitnil,omitempty" name:"OffSet"`

	// 限制
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序参数
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`
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
	Data *BatchIpAccessControlData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type DescribeCCAutoStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
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
	// 配置状态，0表示关闭，1表示开启
	AutoCCSwitch *int64 `json:"AutoCCSwitch,omitnil,omitempty" name:"AutoCCSwitch"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 目前支持根据ts_version排序
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 过滤数组,name可以是如下的值： RuleID,ParamName,Url,Action,Method,Source,Status
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`

	// asc或者desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeCCRuleListRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的API所属的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 目前支持根据ts_version排序
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 过滤数组,name可以是如下的值： RuleID,ParamName,Url,Action,Method,Source,Status
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`

	// asc或者desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
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
	Data *CCRuleLists `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 页码
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页的数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序参数
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// clb-waf 或者 sparta-waf
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 过滤条件
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeCCRuleRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 页码
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页的数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序参数
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// clb-waf 或者 sparta-waf
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 过滤条件
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	Data *CCRuleData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 证书类型。 0：不检测国际标准证书 1：证书来源为自有证书 2：证书来源为托管证书
	CertType *int64 `json:"CertType,omitnil,omitempty" name:"CertType"`

	// CertType为1时，需要填充此参数，表示自有证书的证书链
	Certificate *string `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// CertType为2时，需要填充此参数，表示腾讯云SSL平台托管的证书id
	CertID *string `json:"CertID,omitnil,omitempty" name:"CertID"`

	// CertType为1时，需要填充此参数，表示自有证书的私钥
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`

	// 国密证书类型。0：不检测国密证书 1：证书来源为自有国密证书 2：证书来源为托管国密证书
	GmCertType *int64 `json:"GmCertType,omitnil,omitempty" name:"GmCertType"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的证书链
	GmCert *string `json:"GmCert,omitnil,omitempty" name:"GmCert"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的私钥
	GmPrivateKey *string `json:"GmPrivateKey,omitnil,omitempty" name:"GmPrivateKey"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的加密证书
	GmEncCert *string `json:"GmEncCert,omitnil,omitempty" name:"GmEncCert"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的加密证书的私钥
	GmEncPrivateKey *string `json:"GmEncPrivateKey,omitnil,omitempty" name:"GmEncPrivateKey"`

	// GmCertType为2时，需要填充此参数，表示腾讯云SSL平台托管的证书id
	GmSSLId *string `json:"GmSSLId,omitnil,omitempty" name:"GmSSLId"`
}

type DescribeCertificateVerifyResultRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 证书类型。 0：不检测国际标准证书 1：证书来源为自有证书 2：证书来源为托管证书
	CertType *int64 `json:"CertType,omitnil,omitempty" name:"CertType"`

	// CertType为1时，需要填充此参数，表示自有证书的证书链
	Certificate *string `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// CertType为2时，需要填充此参数，表示腾讯云SSL平台托管的证书id
	CertID *string `json:"CertID,omitnil,omitempty" name:"CertID"`

	// CertType为1时，需要填充此参数，表示自有证书的私钥
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`

	// 国密证书类型。0：不检测国密证书 1：证书来源为自有国密证书 2：证书来源为托管国密证书
	GmCertType *int64 `json:"GmCertType,omitnil,omitempty" name:"GmCertType"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的证书链
	GmCert *string `json:"GmCert,omitnil,omitempty" name:"GmCert"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的私钥
	GmPrivateKey *string `json:"GmPrivateKey,omitnil,omitempty" name:"GmPrivateKey"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的加密证书
	GmEncCert *string `json:"GmEncCert,omitnil,omitempty" name:"GmEncCert"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的加密证书的私钥
	GmEncPrivateKey *string `json:"GmEncPrivateKey,omitnil,omitempty" name:"GmEncPrivateKey"`

	// GmCertType为2时，需要填充此参数，表示腾讯云SSL平台托管的证书id
	GmSSLId *string `json:"GmSSLId,omitnil,omitempty" name:"GmSSLId"`
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
	delete(f, "GmCertType")
	delete(f, "GmCert")
	delete(f, "GmPrivateKey")
	delete(f, "GmEncCert")
	delete(f, "GmEncPrivateKey")
	delete(f, "GmSSLId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCertificateVerifyResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCertificateVerifyResultResponseParams struct {
	// 状态码。
	// 0：证书正常
	// 310：证书异常
	// 311：证书过期
	// 312：证书即将过期
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误详情
	Detail []*string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 过期时间
	NotAfter *string `json:"NotAfter,omitnil,omitempty" name:"NotAfter"`

	// 证书是否改变。
	// 0：未变化
	// 1：有变化
	Changed *int64 `json:"Changed,omitnil,omitempty" name:"Changed"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Ciphers []*TLSCiphers `json:"Ciphers,omitnil,omitempty" name:"Ciphers"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤数组,name可以是如下的值： RuleID,RuleName,Match
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`

	// asc或者desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// exp_ts或者mod_ts
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 查询的域名列表,访问控制页面不用传
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`
}

type DescribeCustomRuleListRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤数组,name可以是如下的值： RuleID,RuleName,Match
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`

	// asc或者desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// exp_ts或者mod_ts
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 查询的域名列表,访问控制页面不用传
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`
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
	delete(f, "DomainList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomRuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomRuleListResponseParams struct {
	// 规则详情
	RuleList []*DescribeCustomRulesRspRuleListItem `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// 规则条数
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 动作类型，1代表阻断，2代表人机识别，3代表观察，4代表重定向，5代表JS校验
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 跳过的策略
	Bypass *string `json:"Bypass,omitnil,omitempty" name:"Bypass"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 策略名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 重定向地址
	Redirect *string `json:"Redirect,omitnil,omitempty" name:"Redirect"`

	// 策略ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 优先级
	SortId *string `json:"SortId,omitnil,omitempty" name:"SortId"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 策略详情
	Strategies []*Strategy `json:"Strategies,omitnil,omitempty" name:"Strategies"`

	// 事件id
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 生效状态
	ValidStatus *int64 `json:"ValidStatus,omitnil,omitempty" name:"ValidStatus"`

	// 来源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 定时任务类型
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 定时任务配置信息
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`

	// 周期任务粒度
	CronType *string `json:"CronType,omitnil,omitempty" name:"CronType"`

	// 自定义标签，风控规则用，用来表示是内置规则还是用户自定义的
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 拦截页面id
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

// Predefined struct for user
type DescribeCustomWhiteRuleRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤数组,name可以是如下的值： RuleID,RuleName,Match
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`

	// asc或者desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// exp_ts或者mod_ts
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeCustomWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤数组,name可以是如下的值： RuleID,RuleName,Match
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`

	// asc或者desc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// exp_ts或者mod_ts
	By *string `json:"By,omitnil,omitempty" name:"By"`
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
	RuleList []*DescribeCustomRulesRspRuleListItem `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// 规则条数
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	AllDomain *uint64 `json:"AllDomain,omitnil,omitempty" name:"AllDomain"`

	// 最近发现时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 接入域名总数
	WafDomainCount *uint64 `json:"WafDomainCount,omitnil,omitempty" name:"WafDomainCount"`

	// 剩下配额
	LeftDomainCount *uint64 `json:"LeftDomainCount,omitnil,omitempty" name:"LeftDomainCount"`

	// 开启防护域名数
	OpenWafDomain *uint64 `json:"OpenWafDomain,omitnil,omitempty" name:"OpenWafDomain"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名id
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDomainDetailsClbRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名id
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	DomainsClbPartInfo *ClbDomainsInfo `json:"DomainsClbPartInfo,omitnil,omitempty" name:"DomainsClbPartInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名id
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDomainDetailsSaasRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名id
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	DomainsPartInfo *DomainsPartInfo `json:"DomainsPartInfo,omitnil,omitempty" name:"DomainsPartInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type DescribeDomainRulesRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
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
	Rules []*Rule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type DescribeDomainVerifyResultRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
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
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 检验状态：0表示可以添加，大于0为不能添加
	VerifyCode *int64 `json:"VerifyCode,omitnil,omitempty" name:"VerifyCode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 请求的白名单匹配路径
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 翻到多少页
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页展示的条数
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 排序方式,desc表示降序，asc表示升序
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type DescribeDomainWhiteRulesRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 请求的白名单匹配路径
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 翻到多少页
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页展示的条数
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 排序方式,desc表示降序，asc表示升序
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
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
	RuleList []*RuleList `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// 规则的数量
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回域名的数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤数组，过滤字段包括：Edition：实例版本，sparta-waf或clb-waf Domain：域名 DomainId：域名ID InstanceName：实例名称 InstanceId：实例ID FlowMode：流量接入模式，仅支持CLBWAF FlowCheckMode：流量体检模式，仅支持CLBWAF ClsStatus：日志开关 Status：WAF开关BotStatus：BOT开关 ApiStatus：API安全开关 Engine：引擎模式 UpstreamIP：源站IP，仅支持SAAS型WAF UpstreamDomain：源站域名，仅支持SAAS型WAF DomainState：域名状态，仅支持SAAS型WAF SgState：安全组状态，仅支持SAAS型WAF Label：分组标签，同时仅支持一种标签过滤
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，取Limit整数倍。最小值为0，最大值= Total/Limit向上取整
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回域名的数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤数组，过滤字段包括：Edition：实例版本，sparta-waf或clb-waf Domain：域名 DomainId：域名ID InstanceName：实例名称 InstanceId：实例ID FlowMode：流量接入模式，仅支持CLBWAF FlowCheckMode：流量体检模式，仅支持CLBWAF ClsStatus：日志开关 Status：WAF开关BotStatus：BOT开关 ApiStatus：API安全开关 Engine：引擎模式 UpstreamIP：源站IP，仅支持SAAS型WAF UpstreamDomain：源站域名，仅支持SAAS型WAF DomainState：域名状态，仅支持SAAS型WAF SgState：安全组状态，仅支持SAAS型WAF Label：分组标签，同时仅支持一种标签过滤
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// domain列表
	Domains []*DomainInfo `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页容量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 是否接入waf；传"1"返回接入域名的详情，传"0"返回未接入域名的详情，传""返回接入和未接入域名的详情
	IsWafDomain *string `json:"IsWafDomain,omitnil,omitempty" name:"IsWafDomain"`

	// 排序参数
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeFindDomainListRequest struct {
	*tchttp.BaseRequest
	
	// 分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页容量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 是否接入waf；传"1"返回接入域名的详情，传"0"返回未接入域名的详情，传""返回接入和未接入域名的详情
	IsWafDomain *string `json:"IsWafDomain,omitnil,omitempty" name:"IsWafDomain"`

	// 排序参数
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 域名信息列表
	List []*FindAllDomainDetail `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 起始时间戳，精度秒
	StartTs *int64 `json:"StartTs,omitnil,omitempty" name:"StartTs"`

	// 结束时间戳，精度秒
	EndTs *int64 `json:"EndTs,omitnil,omitempty" name:"EndTs"`
}

type DescribeFlowTrendRequest struct {
	*tchttp.BaseRequest
	
	// 需要获取流量趋势的域名, all表示所有域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 起始时间戳，精度秒
	StartTs *int64 `json:"StartTs,omitnil,omitempty" name:"StartTs"`

	// 结束时间戳，精度秒
	EndTs *int64 `json:"EndTs,omitnil,omitempty" name:"EndTs"`
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
	Data []*BotStatPointItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	FromTime *string `json:"FromTime,omitnil,omitempty" name:"FromTime"`

	// 结束时间
	ToTime *string `json:"ToTime,omitnil,omitempty" name:"ToTime"`

	// 聚类字段，ip为ip聚合，art为响应耗时聚合，url为url聚合，local为ip转化的城市聚合,qps为每秒请求数峰值的聚合，up为上行带宽峰值聚合，down为下行带峰值聚合
	QueryField *string `json:"QueryField,omitnil,omitempty" name:"QueryField"`

	// 条件，access为访问日志，attack为攻击日志
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 兼容Host，逐步淘汰Host字段
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 域名过滤，不传则不过滤，用于替代Host字段，逐步淘汰Host
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type DescribeHistogramRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间
	FromTime *string `json:"FromTime,omitnil,omitempty" name:"FromTime"`

	// 结束时间
	ToTime *string `json:"ToTime,omitnil,omitempty" name:"ToTime"`

	// 聚类字段，ip为ip聚合，art为响应耗时聚合，url为url聚合，local为ip转化的城市聚合,qps为每秒请求数峰值的聚合，up为上行带宽峰值聚合，down为下行带峰值聚合
	QueryField *string `json:"QueryField,omitnil,omitempty" name:"QueryField"`

	// 条件，access为访问日志，attack为攻击日志
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 兼容Host，逐步淘汰Host字段
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 域名过滤，不传则不过滤，用于替代Host字段，逐步淘汰Host
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
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
	Histogram []*string `json:"Histogram,omitnil,omitempty" name:"Histogram"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 流量来源
	AlbType *string `json:"AlbType,omitnil,omitempty" name:"AlbType"`
}

type DescribeHostLimitRequest struct {
	*tchttp.BaseRequest
	
	// 添加的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 流量来源
	AlbType *string `json:"AlbType,omitnil,omitempty" name:"AlbType"`
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
	Success *ResponseCode `json:"Success,omitnil,omitempty" name:"Success"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type DescribeHostRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
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
	Host *HostRecord `json:"Host,omitnil,omitempty" name:"Host"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 防护域名ID，如果是要查询某一具体的防护域名则传入此参数，要求是准确的域名ID，此参数不支持模糊搜索
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 搜索条件，根据此参数对域名做模糊搜索
	Search *string `json:"Search,omitnil,omitempty" name:"Search"`

	// 复杂的搜索条件
	Item *SearchItem `json:"Item,omitnil,omitempty" name:"Item"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type DescribeHostsRequest struct {
	*tchttp.BaseRequest
	
	// 防护域名，如果是要查询某一具体的防护域名则传入此参数，要求是准确的域名，此参数不支持模糊搜索
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 防护域名ID，如果是要查询某一具体的防护域名则传入此参数，要求是准确的域名ID，此参数不支持模糊搜索
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 搜索条件，根据此参数对域名做模糊搜索
	Search *string `json:"Search,omitnil,omitempty" name:"Search"`

	// 复杂的搜索条件
	Item *SearchItem `json:"Item,omitnil,omitempty" name:"Item"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 防护域名的列表
	HostList []*HostRecord `json:"HostList,omitnil,omitempty" name:"HostList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤数组
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 释放延期标识
	FreeDelayFlag *uint64 `json:"FreeDelayFlag,omitnil,omitempty" name:"FreeDelayFlag"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 容量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤数组
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 释放延期标识
	FreeDelayFlag *uint64 `json:"FreeDelayFlag,omitnil,omitempty" name:"FreeDelayFlag"`
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
	delete(f, "FreeDelayFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// instance列表
	Instances []*InstanceInfo `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 计数标识
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 动作，40表示查询白名单，42表示查询黑名单
	ActionType *uint64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 最小有效时间的时间戳
	//
	// Deprecated: VtsMin is deprecated.
	VtsMin *uint64 `json:"VtsMin,omitnil,omitempty" name:"VtsMin"`

	// 最大有效时间的时间戳
	//
	// Deprecated: VtsMax is deprecated.
	VtsMax *uint64 `json:"VtsMax,omitnil,omitempty" name:"VtsMax"`

	// 最小创建时间的时间戳
	CtsMin *uint64 `json:"CtsMin,omitnil,omitempty" name:"CtsMin"`

	// 最大创建时间的时间戳
	CtsMax *uint64 `json:"CtsMax,omitnil,omitempty" name:"CtsMax"`

	// 分页偏移量，取Limit整数倍。最小值为0，最大值= Total/Limit向上取整
	OffSet *uint64 `json:"OffSet,omitnil,omitempty" name:"OffSet"`

	// 每页返回的数量，默认为20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 来源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 排序参数
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 生效状态
	ValidStatus *int64 `json:"ValidStatus,omitnil,omitempty" name:"ValidStatus"`

	// 最小有效时间的时间戳
	ValidTimeStampMin *string `json:"ValidTimeStampMin,omitnil,omitempty" name:"ValidTimeStampMin"`

	// 最大有效时间的时间戳
	ValidTimeStampMax *string `json:"ValidTimeStampMax,omitnil,omitempty" name:"ValidTimeStampMax"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 定时任务类型筛选0 1 2 3 4
	TimerType *int64 `json:"TimerType,omitnil,omitempty" name:"TimerType"`
}

type DescribeIpAccessControlRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 计数标识
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 动作，40表示查询白名单，42表示查询黑名单
	ActionType *uint64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 最小有效时间的时间戳
	VtsMin *uint64 `json:"VtsMin,omitnil,omitempty" name:"VtsMin"`

	// 最大有效时间的时间戳
	VtsMax *uint64 `json:"VtsMax,omitnil,omitempty" name:"VtsMax"`

	// 最小创建时间的时间戳
	CtsMin *uint64 `json:"CtsMin,omitnil,omitempty" name:"CtsMin"`

	// 最大创建时间的时间戳
	CtsMax *uint64 `json:"CtsMax,omitnil,omitempty" name:"CtsMax"`

	// 分页偏移量，取Limit整数倍。最小值为0，最大值= Total/Limit向上取整
	OffSet *uint64 `json:"OffSet,omitnil,omitempty" name:"OffSet"`

	// 每页返回的数量，默认为20
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 来源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 排序参数
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 生效状态
	ValidStatus *int64 `json:"ValidStatus,omitnil,omitempty" name:"ValidStatus"`

	// 最小有效时间的时间戳
	ValidTimeStampMin *string `json:"ValidTimeStampMin,omitnil,omitempty" name:"ValidTimeStampMin"`

	// 最大有效时间的时间戳
	ValidTimeStampMax *string `json:"ValidTimeStampMax,omitnil,omitempty" name:"ValidTimeStampMax"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 定时任务类型筛选0 1 2 3 4
	TimerType *int64 `json:"TimerType,omitnil,omitempty" name:"TimerType"`
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
	delete(f, "RuleId")
	delete(f, "TimerType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpAccessControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpAccessControlResponseParams struct {
	// 输出
	Data *IpAccessControlData `json:"Data,omitnil,omitempty" name:"Data"`

	// 已经使用的IP黑白名单的IP总数
	UsedTotal *uint64 `json:"UsedTotal,omitnil,omitempty" name:"UsedTotal"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 计数标识
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 类别，ip封禁传值auto_deny
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 有效时间最小时间戳
	//
	// Deprecated: VtsMin is deprecated.
	VtsMin *uint64 `json:"VtsMin,omitnil,omitempty" name:"VtsMin"`

	// 有效时间最大时间戳
	//
	// Deprecated: VtsMax is deprecated.
	VtsMax *uint64 `json:"VtsMax,omitnil,omitempty" name:"VtsMax"`

	// 创建时间最小时间戳
	CtsMin *uint64 `json:"CtsMin,omitnil,omitempty" name:"CtsMin"`

	// 创建时间最大时间戳
	CtsMax *uint64 `json:"CtsMax,omitnil,omitempty" name:"CtsMax"`

	// 偏移参数
	Skip *uint64 `json:"Skip,omitnil,omitempty" name:"Skip"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 策略名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 排序参数
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 有效时间最小时间戳
	ValidTimeStampMin *uint64 `json:"ValidTimeStampMin,omitnil,omitempty" name:"ValidTimeStampMin"`

	// 有效时间最大时间戳
	ValidTimeStampMax *uint64 `json:"ValidTimeStampMax,omitnil,omitempty" name:"ValidTimeStampMax"`
}

type DescribeIpHitItemsRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 计数标识
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 类别，ip封禁传值auto_deny
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 有效时间最小时间戳
	VtsMin *uint64 `json:"VtsMin,omitnil,omitempty" name:"VtsMin"`

	// 有效时间最大时间戳
	VtsMax *uint64 `json:"VtsMax,omitnil,omitempty" name:"VtsMax"`

	// 创建时间最小时间戳
	CtsMin *uint64 `json:"CtsMin,omitnil,omitempty" name:"CtsMin"`

	// 创建时间最大时间戳
	CtsMax *uint64 `json:"CtsMax,omitnil,omitempty" name:"CtsMax"`

	// 偏移参数
	Skip *uint64 `json:"Skip,omitnil,omitempty" name:"Skip"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 策略名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 排序参数
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 有效时间最小时间戳
	ValidTimeStampMin *uint64 `json:"ValidTimeStampMin,omitnil,omitempty" name:"ValidTimeStampMin"`

	// 有效时间最大时间戳
	ValidTimeStampMax *uint64 `json:"ValidTimeStampMax,omitnil,omitempty" name:"ValidTimeStampMax"`
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
	Data *IpHitItemsData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type DescribeModuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 要查询状态的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
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
	WebSecurity *uint64 `json:"WebSecurity,omitnil,omitempty" name:"WebSecurity"`

	// 访问控制规则是否开启
	AccessControl *int64 `json:"AccessControl,omitnil,omitempty" name:"AccessControl"`

	// CC防护是否开启
	CcProtection *uint64 `json:"CcProtection,omitnil,omitempty" name:"CcProtection"`

	// 网页防篡改是否开启
	AntiTamper *uint64 `json:"AntiTamper,omitnil,omitempty" name:"AntiTamper"`

	// 信息防泄漏是否开启
	AntiLeakage *uint64 `json:"AntiLeakage,omitnil,omitempty" name:"AntiLeakage"`

	// API安全是否开启
	ApiProtection *uint64 `json:"ApiProtection,omitnil,omitempty" name:"ApiProtection"`

	// 限流模块开关
	RateLimit *uint64 `json:"RateLimit,omitnil,omitempty" name:"RateLimit"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 支持的过滤器:	ObjectId: clb实例ID	VIP: clb实例的公网IP	InstanceId: waf实例ID	Domain: 精准域名	Status: waf防护开关状态: 0关闭，1开启	ClsStatus: waf日志开关: 0关闭，1开启   
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeObjectsRequest struct {
	*tchttp.BaseRequest
	
	// 支持的过滤器:	ObjectId: clb实例ID	VIP: clb实例的公网IP	InstanceId: waf实例ID	Domain: 精准域名	Status: waf防护开关状态: 0关闭，1开启	ClsStatus: waf日志开关: 0关闭，1开启   
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	ClbObjects []*ClbObject `json:"ClbObjects,omitnil,omitempty" name:"ClbObjects"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	FromTime *string `json:"FromTime,omitnil,omitempty" name:"FromTime"`

	// 查询终止时间
	ToTime *string `json:"ToTime,omitnil,omitempty" name:"ToTime"`

	// 查询的域名，如果查询所有域名数据，该参数不填写
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

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
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`
}

type DescribePeakPointsRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间
	FromTime *string `json:"FromTime,omitnil,omitempty" name:"FromTime"`

	// 查询终止时间
	ToTime *string `json:"ToTime,omitnil,omitempty" name:"ToTime"`

	// 查询的域名，如果查询所有域名数据，该参数不填写
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

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
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`
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
	Points []*PeakPointsItem `json:"Points,omitnil,omitempty" name:"Points"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	FromTime *string `json:"FromTime,omitnil,omitempty" name:"FromTime"`

	// 查询结束时间
	ToTime *string `json:"ToTime,omitnil,omitempty" name:"ToTime"`

	// 需要查询的域名，当前用户所有域名可以不传
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 五个值可选：
	// access-峰值qps
	// down-下行峰值带宽
	// up-上行峰值带宽
	// attack-Web攻击总数
	// cc-CC攻击总数趋势图
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`
}

type DescribePeakValueRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间
	FromTime *string `json:"FromTime,omitnil,omitempty" name:"FromTime"`

	// 查询结束时间
	ToTime *string `json:"ToTime,omitnil,omitempty" name:"ToTime"`

	// 需要查询的域名，当前用户所有域名可以不传
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 五个值可选：
	// access-峰值qps
	// down-下行峰值带宽
	// up-上行峰值带宽
	// attack-Web攻击总数
	// cc-CC攻击总数趋势图
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`
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
	Access *uint64 `json:"Access,omitnil,omitempty" name:"Access"`

	// 上行带宽峰值，单位B
	Up *uint64 `json:"Up,omitnil,omitempty" name:"Up"`

	// 下行带宽峰值，单位B
	Down *uint64 `json:"Down,omitnil,omitempty" name:"Down"`

	// Web攻击总数
	Attack *uint64 `json:"Attack,omitnil,omitempty" name:"Attack"`

	// CC攻击总数
	Cc *uint64 `json:"Cc,omitnil,omitempty" name:"Cc"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// clb-waf或者saas-waf
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`
}

type DescribePolicyStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// clb-waf或者saas-waf
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 防护状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 实例类型
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`
}

type DescribePortsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 实例类型
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`
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
	HttpPorts []*string `json:"HttpPorts,omitnil,omitempty" name:"HttpPorts"`

	// https端口列表
	HttpsPorts []*string `json:"HttpsPorts,omitnil,omitempty" name:"HttpsPorts"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeProtectionModesRequestParams struct {
	// sparta-waf或clb
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type DescribeProtectionModesRequest struct {
	*tchttp.BaseRequest
	
	// sparta-waf或clb
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

func (r *DescribeProtectionModesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProtectionModesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Edition")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProtectionModesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProtectionModesResponseParams struct {
	// 规则大类ID及防护模式
	Modes []*TigaMainClassMode `json:"Modes,omitnil,omitempty" name:"Modes"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProtectionModesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProtectionModesResponseParams `json:"Response"`
}

func (r *DescribeProtectionModesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProtectionModesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleLimitRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeRuleLimitRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	Res *WafRuleLimit `json:"Res,omitnil,omitempty" name:"Res"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeScanIpRequestParams struct {
	// 要查询的ip地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`
}

type DescribeScanIpRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的ip地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`
}

func (r *DescribeScanIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanIpResponseParams struct {
	// ip列表,当入参Ip为all时，返回所有已知ip列表
	IpList []*ScanIpInfo `json:"IpList,omitnil,omitempty" name:"IpList"`

	// 所属业务
	Bussiness *string `json:"Bussiness,omitnil,omitempty" name:"Bussiness"`

	// 业务特征
	Characteristic *string `json:"Characteristic,omitnil,omitempty" name:"Characteristic"`

	// 扫描说明
	Descibe *string `json:"Descibe,omitnil,omitempty" name:"Descibe"`

	// 官方公告
	Referer *string `json:"Referer,omitnil,omitempty" name:"Referer"`

	// 扫描示例
	Demo *string `json:"Demo,omitnil,omitempty" name:"Demo"`

	// 扫描对象
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 扫描目的
	Purpose *string `json:"Purpose,omitnil,omitempty" name:"Purpose"`

	// 产品文案
	Announcement *string `json:"Announcement,omitnil,omitempty" name:"Announcement"`

	// 更新时间
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// ipowner
	IpOwner *string `json:"IpOwner,omitnil,omitempty" name:"IpOwner"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScanIpResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanIpResponseParams `json:"Response"`
}

func (r *DescribeScanIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// clb-waf或者sparta-waf
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`
}

type DescribeSessionRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// clb-waf或者sparta-waf
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`
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
	Data *SessionData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 版本
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type DescribeSpartaProtectionInfoRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 版本
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 实例
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// cname取值
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 源IP地址列表
	SrcList []*string `json:"SrcList,omitnil,omitempty" name:"SrcList"`

	// 证书类型
	CertType *string `json:"CertType,omitnil,omitempty" name:"CertType"`

	// 证书
	Cert *string `json:"Cert,omitnil,omitempty" name:"Cert"`

	// 私有密钥
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`

	// ssl的id
	Sslid *string `json:"Sslid,omitnil,omitempty" name:"Sslid"`

	// 是否是cdn
	IsCdn *string `json:"IsCdn,omitnil,omitempty" name:"IsCdn"`

	// 灰度区域列表
	GrayAreas []*string `json:"GrayAreas,omitnil,omitempty" name:"GrayAreas"`

	// 引擎
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// HTTPS重写
	HttpsRewrite *string `json:"HttpsRewrite,omitnil,omitempty" name:"HttpsRewrite"`

	// upstreamType取值
	UpstreamType *string `json:"UpstreamType,omitnil,omitempty" name:"UpstreamType"`

	// upstreamDomain取值
	UpstreamDomain *string `json:"UpstreamDomain,omitnil,omitempty" name:"UpstreamDomain"`

	// upstreamScheme取值
	UpstreamScheme *string `json:"UpstreamScheme,omitnil,omitempty" name:"UpstreamScheme"`

	// 是否是HTTP2
	IsHttp2 *string `json:"IsHttp2,omitnil,omitempty" name:"IsHttp2"`

	// 是否含有websocket
	IsWebsocket *string `json:"IsWebsocket,omitnil,omitempty" name:"IsWebsocket"`

	// loadBalance信息
	LoadBalance *string `json:"LoadBalance,omitnil,omitempty" name:"LoadBalance"`

	// httpsUpstreamPort取值
	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitnil,omitempty" name:"HttpsUpstreamPort"`

	// port信息
	Ports []*PortItem `json:"Ports,omitnil,omitempty" name:"Ports"`

	// 是否灰度
	IsGray *string `json:"IsGray,omitnil,omitempty" name:"IsGray"`

	// 模式
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 防御等级,100,200,300
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 与源站是否保持长连接
	IsKeepAlive *string `json:"IsKeepAlive,omitnil,omitempty" name:"IsKeepAlive"`

	// 0：BGP 1：Anycast
	Anycast *string `json:"Anycast,omitnil,omitempty" name:"Anycast"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// TLS信息
	TLS []*TLSVersion `json:"TLS,omitnil,omitempty" name:"TLS"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	FromTime *string `json:"FromTime,omitnil,omitempty" name:"FromTime"`

	// 查询结束时间
	ToTime *string `json:"ToTime,omitnil,omitempty" name:"ToTime"`

	// TOP N,可从0-10选择，默认是10
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type DescribeTopAttackDomainRequest struct {
	*tchttp.BaseRequest
	
	// 查询起始时间
	FromTime *string `json:"FromTime,omitnil,omitempty" name:"FromTime"`

	// 查询结束时间
	ToTime *string `json:"ToTime,omitnil,omitempty" name:"ToTime"`

	// TOP N,可从0-10选择，默认是10
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 只有两个值有效，sparta-waf，clb-waf，不传则不过滤
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// WAF实例ID，不传则不过滤
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
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
	CC []*KVInt `json:"CC,omitnil,omitempty" name:"CC"`

	// Web攻击域名列表
	Web []*KVInt `json:"Web,omitnil,omitempty" name:"Web"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Data []*CdcRegion `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 流量来源，不填默认clb。clb：负载均衡器，tsegw：云原生API网关，scf：云函数，apisix：腾讯云上其他网关
	AlbType *string `json:"AlbType,omitnil,omitempty" name:"AlbType"`
}

type DescribeUserClbWafRegionsRequest struct {
	*tchttp.BaseRequest
	
	// 流量来源，不填默认clb。clb：负载均衡器，tsegw：云原生API网关，scf：云函数，apisix：腾讯云上其他网关
	AlbType *string `json:"AlbType,omitnil,omitempty" name:"AlbType"`
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
	delete(f, "AlbType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserClbWafRegionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserClbWafRegionsResponseParams struct {
	// 地域（标准的ap-格式）列表
	Data []*string `json:"Data,omitnil,omitempty" name:"Data"`

	// 包含详细属性的地域信息
	RichDatas []*ClbWafRegionItem `json:"RichDatas,omitnil,omitempty" name:"RichDatas"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	UsersInfo []*UserDomainInfo `json:"UsersInfo,omitnil,omitempty" name:"UsersInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type DescribeUserLevelRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
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
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页容量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段，支持 signature_id, modify_time
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 筛选条件，支持 MainClassName，SubClassID ,CveID, Status, ID;  ID为规则id
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeUserSignatureRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 分页
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页容量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段，支持 signature_id, modify_time
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 筛选条件，支持 MainClassName，SubClassID ,CveID, Status, ID;  ID为规则id
	Filters []*FiltersItemNew `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 规则列表
	Rules []*UserSignatureRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeVipInfoRequest struct {
	*tchttp.BaseRequest
	
	// waf实例id列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
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
	VipInfo []*VipInfo `json:"VipInfo,omitnil,omitempty" name:"VipInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeWafAutoDenyRulesRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	AttackThreshold *int64 `json:"AttackThreshold,omitnil,omitempty" name:"AttackThreshold"`

	// 攻击时间阈值
	TimeThreshold *int64 `json:"TimeThreshold,omitnil,omitempty" name:"TimeThreshold"`

	// 自动封禁时间
	DenyTimeThreshold *int64 `json:"DenyTimeThreshold,omitnil,omitempty" name:"DenyTimeThreshold"`

	// 自动封禁状态
	DefenseStatus *int64 `json:"DefenseStatus,omitnil,omitempty" name:"DefenseStatus"`

	// 数据来源Source字段 custom-自定义(默认)、batch-domain-批量域名
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 重保护网域名状态
	HWState *int64 `json:"HWState,omitnil,omitempty" name:"HWState"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	WafAutoDenyDetails *AutoDenyDetail `json:"WafAutoDenyDetails,omitnil,omitempty" name:"WafAutoDenyDetails"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	WafThreatenIntelligenceDetails *WafThreatenIntelligenceDetails `json:"WafThreatenIntelligenceDetails,omitnil,omitempty" name:"WafThreatenIntelligenceDetails"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type DescribeWebshellStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 开关状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// cname地址
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// 域名所属实例类型。
	// sparta-waf：SaaS型WAF实例
	// clb-waf：负载均衡型WAF实例
	// cdc-clb-waf：CDC环境下负载均衡型WAF实例
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 地域。
	// "多伦多": "ca"
	// "广州": "gz"
	// "成都": "cd"
	// "福州": "fzec"
	// "深圳": "szx"
	// "印度": "in"
	// "济南": "jnec"
	// "重庆": "cq"
	// "天津": "tsn"
	// "欧洲东北": "ru"
	// "南京": "nj"
	// "美国硅谷": "usw"
	// "泰国": "th"
	// "广州Open": "gzopen"
	// "深圳金融": "szjr"
	// "法兰克福": "de"
	// "日本": "jp"
	// "弗吉尼亚": "use"
	// "北京": "bj"
	// "中国香港": "hk"
	// "杭州": "hzec"
	// "北京金融": "bjjr"
	// "上海金融": "shjr"
	// "台北": "tpe"
	// "首尔": "kr"
	// "上海": "sh"
	// "新加坡": "sg"
	// "清远": "qy"
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 访问日志开关状态。
	// 0：关闭
	// 1：开启
	ClsStatus *uint64 `json:"ClsStatus,omitnil,omitempty" name:"ClsStatus"`

	// 负载均衡型WAF使用模式。
	// 0：镜像模式 
	// 1：清洗模式
	FlowMode *uint64 `json:"FlowMode,omitnil,omitempty" name:"FlowMode"`

	// waf开关状态。
	// 0：关闭 
	// 1：开启
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 规则引擎防护模式。
	// 0：观察模式 
	// 1：拦截模式
	Mode *uint64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 规则引擎和AI引擎防护模式联合状态。
	// 1:初始状态,规则引擎拦截&&AI引擎未操作开关状态
	// 10：规则引擎观察&&AI引擎关闭模式 
	// 11：规则引擎观察&&AI引擎观察模式 
	// 12：规则引擎观察&&AI引擎拦截模式 
	// 20：规则引擎拦截&&AI引擎关闭模式 
	// 21：规则引擎拦截&&AI引擎观察模式 
	// 22：规则引擎拦截&&AI引擎拦截模式
	Engine *uint64 `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 沙箱集群回源出口IP列表
	CCList []*string `json:"CCList,omitnil,omitempty" name:"CCList"`

	// 生产集群回源出口IP列表
	RsList []*string `json:"RsList,omitnil,omitempty" name:"RsList"`

	// 服务端口配置
	Ports []*PortInfo `json:"Ports,omitnil,omitempty" name:"Ports"`

	// 负载均衡器相关配置
	LoadBalancerSet []*LoadBalancerPackageNew `json:"LoadBalancerSet,omitnil,omitempty" name:"LoadBalancerSet"`

	// 用户id
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 负载均衡型WAF域名LB监听器状态。
	// 0：操作成功 
	// 4：正在绑定LB 
	// 6：正在解绑LB 
	// 7：解绑LB失败 
	// 8：绑定LB失败 
	// 10：内部错误
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Ipv6开关状态。
	// 0：关闭 
	// 1：开启
	Ipv6Status *int64 `json:"Ipv6Status,omitnil,omitempty" name:"Ipv6Status"`

	// BOT开关状态。
	// 0：关闭 
	// 1：关闭
	// 2：开启
	// 3：开启
	BotStatus *int64 `json:"BotStatus,omitnil,omitempty" name:"BotStatus"`

	// 实例版本信息。
	// 101：小微敏捷版 
	// 102：小微超轻版
	// 2：高级版
	// 3：企业版
	// 4：旗舰版
	// 6：独享版
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 投递CLS状态。
	// 0：关闭 
	// 1：开启
	PostCLSStatus *int64 `json:"PostCLSStatus,omitnil,omitempty" name:"PostCLSStatus"`

	// 投递CKafka状态。
	// 0：关闭 
	// 1：开启
	PostCKafkaStatus *int64 `json:"PostCKafkaStatus,omitnil,omitempty" name:"PostCKafkaStatus"`

	// cdc实例域名接入的集群信息,非cdc实例忽略。
	CdcClusters *string `json:"CdcClusters,omitnil,omitempty" name:"CdcClusters"`

	// api安全开关状态。
	// 0：关闭 
	// 1：开启
	ApiStatus *int64 `json:"ApiStatus,omitnil,omitempty" name:"ApiStatus"`

	// 应用型负载均衡类型，默认clb。
	// clb：七层负载均衡器类型
	// apisix：apisix网关型
	AlbType *string `json:"AlbType,omitnil,omitempty" name:"AlbType"`

	// 安全组状态。
	// 0：不展示
	// 1：非腾讯云源站
	// 2：安全组绑定失败
	// 3：安全组发生变更
	SgState *int64 `json:"SgState,omitnil,omitempty" name:"SgState"`

	// 安全组状态的详细解释
	SgDetail *string `json:"SgDetail,omitnil,omitempty" name:"SgDetail"`

	// 域名云环境。hybrid：混合云域名
	// public：公有云域名
	CloudType *string `json:"CloudType,omitnil,omitempty" name:"CloudType"`

	// 域名备注信息
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// SAASWAF源站IP列表
	SrcList []*string `json:"SrcList,omitnil,omitempty" name:"SrcList"`

	// SAASWAF源站域名列表
	UpstreamDomainList []*string `json:"UpstreamDomainList,omitnil,omitempty" name:"UpstreamDomainList"`

	// 安全组ID
	SgID *string `json:"SgID,omitnil,omitempty" name:"SgID"`

	// clbwaf接入状态
	AccessStatus *int64 `json:"AccessStatus,omitnil,omitempty" name:"AccessStatus"`

	// 域名标签
	Labels []*string `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type DomainPackageNew struct {
	// 资源ID
	ResourceIds *string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 过期时间
	ValidTime *string `json:"ValidTime,omitnil,omitempty" name:"ValidTime"`

	// 是否自动续费，1：自动续费，0：不自动续费
	RenewFlag *uint64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 套餐购买个数
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 套餐购买地域，clb-waf暂时没有用到
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type DomainRuleId struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则id
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type DomainURI struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 版本
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type DomainsPartInfo struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名唯一ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 域名所属实例唯一ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 域名所属实例类型
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 域名所属实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 证书
	Cert *string `json:"Cert,omitnil,omitempty" name:"Cert"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 规则引擎和AI引擎防护模式联合状态。
	// 1:初始状态,规则引擎拦截&&AI引擎未操作开关状态
	// 10：规则引擎观察&&AI引擎关闭模式 
	// 11：规则引擎观察&&AI引擎观察模式 
	// 12：规则引擎观察&&AI引擎拦截模式 
	// 20：规则引擎拦截&&AI引擎关闭模式 
	// 21：规则引擎拦截&&AI引擎观察模式 
	// 22：规则引擎拦截&&AI引擎拦截模式
	Engine *uint64 `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 是否开启HTTP强制跳转到HTTPS。
	// 0：不强制跳转
	// 1：开启强制跳转
	HttpsRewrite *uint64 `json:"HttpsRewrite,omitnil,omitempty" name:"HttpsRewrite"`

	// HTTPS回源端口
	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitnil,omitempty" name:"HttpsUpstreamPort"`

	// waf前是否部署有七层代理服务。
	// 0：没有部署代理服务
	// 1：有部署代理服务，waf将使用XFF获取客户端IP
	// 2：有部署代理服务，waf将使用remote_addr获取客户端IP
	// 3：有部署代理服务，waf将使用ip_headers中的自定义header获取客户端IP
	IsCdn *uint64 `json:"IsCdn,omitnil,omitempty" name:"IsCdn"`

	// 是否开启灰度，已废弃。
	IsGray *uint64 `json:"IsGray,omitnil,omitempty" name:"IsGray"`

	// 是否开启HTTP2，需要开启HTTPS协议支持。
	// 0：关闭
	// 1：开启
	IsHttp2 *uint64 `json:"IsHttp2,omitnil,omitempty" name:"IsHttp2"`

	// 是否开启WebSocket支持。
	// 0：关闭
	// 1：开启
	IsWebsocket *uint64 `json:"IsWebsocket,omitnil,omitempty" name:"IsWebsocket"`

	// 回源负载均衡策略。
	// 0：轮询
	// 1：IP hash
	// 2：加权轮询
	LoadBalance *uint64 `json:"LoadBalance,omitnil,omitempty" name:"LoadBalance"`

	// 防护模式。
	// 0：观察模式
	// 1：拦截模式
	Mode *uint64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 自有证书的私钥
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`

	// CertType为2时，需要填充此参数，表示腾讯云SSL平台托管的证书id
	SSLId *string `json:"SSLId,omitnil,omitempty" name:"SSLId"`

	// 域名回源时的回源域名。UpstreamType为1时，需要填充此字段
	UpstreamDomain *string `json:"UpstreamDomain,omitnil,omitempty" name:"UpstreamDomain"`

	// 回源类型。
	// 0：通过IP回源
	// 1：通过域名回源
	UpstreamType *uint64 `json:"UpstreamType,omitnil,omitempty" name:"UpstreamType"`

	// IP回源时的回源IP列表。UpstreamType为0时，需要填充此字段
	SrcList []*string `json:"SrcList,omitnil,omitempty" name:"SrcList"`

	// 域名端口配置
	Ports []*PortInfo `json:"Ports,omitnil,omitempty" name:"Ports"`

	// 证书类型。
	// 0：仅配置HTTP监听端口，没有证书
	// 1：证书来源为自有证书
	// 2：证书来源为托管证书
	CertType *uint64 `json:"CertType,omitnil,omitempty" name:"CertType"`

	// 服务配置有HTTPS端口时，HTTPS的回源协议。
	// http：使用http协议回源，和HttpsUpstreamPort配合使用
	// https：使用https协议回源
	UpstreamScheme *string `json:"UpstreamScheme,omitnil,omitempty" name:"UpstreamScheme"`

	// 日志包是否开启。
	// 0：关闭
	// 1：开启
	Cls *uint64 `json:"Cls,omitnil,omitempty" name:"Cls"`

	// 接入Cname，SaaS型域名使用此Cname进行接入
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// 是否开启长连接。
	// 0： 短连接
	// 1： 长连接
	IsKeepAlive *uint64 `json:"IsKeepAlive,omitnil,omitempty" name:"IsKeepAlive"`

	// 是否开启主动健康检测。
	// 0：不开启
	// 1：开启
	ActiveCheck *uint64 `json:"ActiveCheck,omitnil,omitempty" name:"ActiveCheck"`

	// TLS版本信息
	TLSVersion *int64 `json:"TLSVersion,omitnil,omitempty" name:"TLSVersion"`

	// 自定义的加密套件列表。CipherTemplate为3时需要填此字段，表示自定义的加密套件，值通过DescribeCiphersDetail接口获取。
	Ciphers []*int64 `json:"Ciphers,omitnil,omitempty" name:"Ciphers"`

	// 加密套件模板。
	// 0：不支持选择，使用默认模板  
	// 1：通用型模板 
	// 2：安全型模板
	// 3：自定义模板
	CipherTemplate *int64 `json:"CipherTemplate,omitnil,omitempty" name:"CipherTemplate"`

	// WAF与源站的读超时时间，默认300s。
	ProxyReadTimeout *int64 `json:"ProxyReadTimeout,omitnil,omitempty" name:"ProxyReadTimeout"`

	// WAF与源站的写超时时间，默认300s。
	ProxySendTimeout *int64 `json:"ProxySendTimeout,omitnil,omitempty" name:"ProxySendTimeout"`

	// WAF回源时的SNI类型。
	// 0：关闭SNI，不配置client_hello中的server_name
	// 1：开启SNI，client_hello中的server_name为防护域名
	// 2：开启SNI，SNI为域名回源时的源站域名
	// 3：开启SNI，SNI为自定义域名
	SniType *int64 `json:"SniType,omitnil,omitempty" name:"SniType"`

	// SniType为3时，需要填此参数，表示自定义的SNI；
	SniHost *string `json:"SniHost,omitnil,omitempty" name:"SniHost"`

	// 回源IP权重
	Weights []*string `json:"Weights,omitnil,omitempty" name:"Weights"`

	// IsCdn=3时，表示自定义header
	IpHeaders []*string `json:"IpHeaders,omitnil,omitempty" name:"IpHeaders"`

	// 是否开启XFF重置。
	// 0：关闭
	// 1：开启
	XFFReset *int64 `json:"XFFReset,omitnil,omitempty" name:"XFFReset"`

	// 域名备注信息
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 自定义回源Host。默认为空字符串，表示使用防护域名作为回源Host。
	UpstreamHost *string `json:"UpstreamHost,omitnil,omitempty" name:"UpstreamHost"`

	// 防护规则
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 是否开启缓存 0-关闭 1-开启
	ProxyBuffer *int64 `json:"ProxyBuffer,omitnil,omitempty" name:"ProxyBuffer"`

	// 国密选项。0：不开启国密 1：在原有TLS选项的基础上追加支持国密 2：开启国密并仅支持国密客户端访问
	GmType *int64 `json:"GmType,omitnil,omitempty" name:"GmType"`

	// 国密证书类型。0：无国密证书 1：证书来源为自有国密证书 2：证书来源为托管国密证书
	GmCertType *int64 `json:"GmCertType,omitnil,omitempty" name:"GmCertType"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的证书链
	GmCert *string `json:"GmCert,omitnil,omitempty" name:"GmCert"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的私钥
	GmPrivateKey *string `json:"GmPrivateKey,omitnil,omitempty" name:"GmPrivateKey"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的加密证书
	GmEncCert *string `json:"GmEncCert,omitnil,omitempty" name:"GmEncCert"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的加密证书的私钥
	GmEncPrivateKey *string `json:"GmEncPrivateKey,omitnil,omitempty" name:"GmEncPrivateKey"`

	// GmCertType为2时，需要填充此参数，表示腾讯云SSL平台托管的证书id
	GmSSLId *string `json:"GmSSLId,omitnil,omitempty" name:"GmSSLId"`

	// 域名标签
	Labels []*string `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type DownloadAttackRecordInfo struct {
	// 记录ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 下载任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 域名
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 当前下载任务的日志条数
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 下载任务运行状态：-1-下载超时，0-下载等待，1-下载完成，2-下载失败，4-正在下载
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 下载文件URL
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后更新修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 下载任务需下载的日志总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type ExportAccessInfo struct {
	// 日志导出任务ID
	ExportId *string `json:"ExportId,omitnil,omitempty" name:"ExportId"`

	// 日志导出查询语句
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 日志导出文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 日志文件大小
	FileSize *int64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 日志导出时间排序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 日志导出格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 日志导出数量
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 日志下载状态。Processing:导出正在进行中，Complete:导出完成，Failed:导出失败，Expired:日志导出已过期（三天有效期）
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 日志导出起始时间
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 日志导出结束时间
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 日志导出路径
	CosPath *string `json:"CosPath,omitnil,omitempty" name:"CosPath"`

	// 日志导出创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type FiltersItemNew struct {
	// 字段名； 过滤
	// 子订单号过滤通过name 为：DealName； value为子订单号
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 是否精确查找
	ExactMatch *bool `json:"ExactMatch,omitnil,omitempty" name:"ExactMatch"`
}

type FindAllDomainDetail struct {
	// 用户id
	Appid *uint64 `json:"Appid,omitnil,omitempty" name:"Appid"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名ip
	Ips []*string `json:"Ips,omitnil,omitempty" name:"Ips"`

	// 发现时间
	FindTime *string `json:"FindTime,omitnil,omitempty" name:"FindTime"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 域名id
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// waf类型
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 是否接入waf
	IsWafDomain *uint64 `json:"IsWafDomain,omitnil,omitempty" name:"IsWafDomain"`
}

type FraudPkg struct {
	// 资源id
	ResourceIds *string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 地域
	Region *int64 `json:"Region,omitnil,omitempty" name:"Region"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 申请数量
	InquireNum *int64 `json:"InquireNum,omitnil,omitempty" name:"InquireNum"`

	// 使用数量
	UsedNum *int64 `json:"UsedNum,omitnil,omitempty" name:"UsedNum"`

	// 续费标志
	RenewFlag *uint64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

// Predefined struct for user
type FreshAntiFakeUrlRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type FreshAntiFakeUrlRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`
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
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Goods []*GoodNews `json:"Goods,omitnil,omitempty" name:"Goods"`
}

type GenerateDealsAndPayNewRequest struct {
	*tchttp.BaseRequest
	
	// 计费下单入参
	Goods []*GoodNews `json:"Goods,omitnil,omitempty" name:"Goods"`
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
	Data *DealData `json:"Data,omitnil,omitempty" name:"Data"`

	// 1:成功，0:失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 返回message
	ReturnMessage *string `json:"ReturnMessage,omitnil,omitempty" name:"ReturnMessage"`

	// 购买的实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Records []*DownloadAttackRecordInfo `json:"Records,omitnil,omitempty" name:"Records"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Lucene语法
	QueryString *string `json:"QueryString,omitnil,omitempty" name:"QueryString"`
}

type GetAttackHistogramRequest struct {
	*tchttp.BaseRequest
	
	// 查询的域名，所有域名使用all
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Lucene语法
	QueryString *string `json:"QueryString,omitnil,omitempty" name:"QueryString"`
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
	Data []*LogHistogramInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 时间段大小
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 统计的条目数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的域名，全部域名填all
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 查询条件，默认为""
	QueryString *string `json:"QueryString,omitnil,omitempty" name:"QueryString"`
}

type GetAttackTotalCountRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询的域名，全部域名填all
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 查询条件，默认为""
	QueryString *string `json:"QueryString,omitnil,omitempty" name:"QueryString"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 套餐类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type GetInstanceQpsLimitRequest struct {
	*tchttp.BaseRequest
	
	// 套餐实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 套餐类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
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
	QpsData *QpsData `json:"QpsData,omitnil,omitempty" name:"QpsData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 商品明细
	GoodsDetail *GoodsDetailNew `json:"GoodsDetail,omitnil,omitempty" name:"GoodsDetail"`

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
	GoodsCategoryId *int64 `json:"GoodsCategoryId,omitnil,omitempty" name:"GoodsCategoryId"`

	// 购买waf实例区域ID
	// 1 表示购买大陆资源;
	// 9表示购买非中国大陆资源
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`
}

type Goods struct {
	// 付费类型，1:预付费，0:后付费
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 商品数量
	GoodsNum *int64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 商品明细
	GoodsDetail *GoodsDetail `json:"GoodsDetail,omitnil,omitempty" name:"GoodsDetail"`

	// 默认为0
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 计费类目ID，对应cid
	GoodsCategoryId *int64 `json:"GoodsCategoryId,omitnil,omitempty" name:"GoodsCategoryId"`

	// 平台类型，默认1
	Platform *int64 `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 购买waf实例区域ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`
}

type GoodsDetail struct {
	// 时间间隔
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 单位，支持m、y、d
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 产品码
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 二级产品码
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// 计费策略id
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// waf产品码
	ProductInfo []*ProductInfo `json:"ProductInfo,omitnil,omitempty" name:"ProductInfo"`

	// waf实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// QPS数量
	ElasticQps *int64 `json:"ElasticQps,omitnil,omitempty" name:"ElasticQps"`

	// 弹性账单
	FlexBill *int64 `json:"FlexBill,omitnil,omitempty" name:"FlexBill"`

	// 1:自动续费，0:不自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// waf购买的实际地域信息
	RealRegion *int64 `json:"RealRegion,omitnil,omitempty" name:"RealRegion"`

	// Waf实例对应的二级产品码
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 计费细项标签数组
	LabelTypes []*string `json:"LabelTypes,omitnil,omitempty" name:"LabelTypes"`

	// 计费细项标签数量，一般和SvLabelType一一对应
	LabelCounts []*int64 `json:"LabelCounts,omitnil,omitempty" name:"LabelCounts"`

	// 变配使用，实例到期时间
	CurDeadline *string `json:"CurDeadline,omitnil,omitempty" name:"CurDeadline"`

	// 对存在的实例购买bot 或api 安全
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type GoodsDetailNew struct {
	// 时间间隔
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 单位，支持购买d、m、y 即（日、月、年）
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

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
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

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
	Pid *int64 `json:"Pid,omitnil,omitempty" name:"Pid"`

	// waf实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 1:自动续费，0:不自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// waf购买的实际地域信息
	RealRegion *int64 `json:"RealRegion,omitnil,omitempty" name:"RealRegion"`

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
	LabelTypes []*string `json:"LabelTypes,omitnil,omitempty" name:"LabelTypes"`

	// 计费细项标签数量，一般和SvLabelType一一对应
	LabelCounts []*int64 `json:"LabelCounts,omitnil,omitempty" name:"LabelCounts"`

	// 变配使用，实例到期时间
	CurDeadline *string `json:"CurDeadline,omitnil,omitempty" name:"CurDeadline"`

	// 对存在的实例购买bot 或api 安全
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 模式clb-waf或者saas-waf
	MicroVersion *string `json:"MicroVersion,omitnil,omitempty" name:"MicroVersion"`
}

type HostDel struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 实例类型
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type HostRecord struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名唯一ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 主域名，入参时为空
	MainDomain *string `json:"MainDomain,omitnil,omitempty" name:"MainDomain"`

	// 规则引擎防护模式。
	// 0：观察模式
	// 1：拦截模式
	Mode *uint64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// waf和负载均衡器的绑定关系。
	// 0：未绑定
	// 1：已绑定
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// clbwaf域名监听器状态。
	// 0：操作成功
	// 4：正在绑定LB
	// 6：正在解绑LB 
	// 7：解绑LB失败 
	// 8：绑定LB失败 
	// 10：内部错误
	State *uint64 `json:"State,omitnil,omitempty" name:"State"`

	// 规则引擎和AI引擎防护模式联合状态。
	// 1:初始状态,规则引擎拦截&&AI引擎未操作开关状态
	// 10：规则引擎观察&&AI引擎关闭模式 
	// 11：规则引擎观察&&AI引擎观察模式 
	// 12：规则引擎观察&&AI引擎拦截模式 
	// 20：规则引擎拦截&&AI引擎关闭模式 
	// 21：规则引擎拦截&&AI引擎观察模式 
	// 22：规则引擎拦截&&AI引擎拦截模式
	Engine *uint64 `json:"Engine,omitnil,omitempty" name:"Engine"`

	// waf前是否部署有七层代理服务。 0：没有部署代理服务 1：有部署代理服务，waf将使用XFF获取客户端IP 2：有部署代理服务，waf将使用remote_addr获取客户端IP 3：有部署代理服务，waf将使用ip_headers中的自定义header获取客户端IP
	IsCdn *uint64 `json:"IsCdn,omitnil,omitempty" name:"IsCdn"`

	// 绑定的负载均衡器信息列表
	LoadBalancerSet []*LoadBalancer `json:"LoadBalancerSet,omitnil,omitempty" name:"LoadBalancerSet"`

	// 域名绑定的LB的地域，以逗号分割多个地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 域名所属实例类型。负载均衡型WAF为"clb-waf"
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 负载均衡型WAF域名的流量模式。
	// 1：清洗模式
	// 0：镜像模式
	FlowMode *uint64 `json:"FlowMode,omitnil,omitempty" name:"FlowMode"`

	// 是否开启访问日志。
	// 1：开启
	// 0：关闭
	ClsStatus *uint64 `json:"ClsStatus,omitnil,omitempty" name:"ClsStatus"`

	// 防护等级，可选值100,200,300
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 域名需要下发到的cdc集群列表。仅CDC场景下填充
	CdcClusters []*string `json:"CdcClusters,omitnil,omitempty" name:"CdcClusters"`

	// 应用型负载均衡类型，默认clb。 
	// clb：七层负载均衡器类型 
	// apisix：apisix网关型
	// tsegw：云原生API网关
	// scf：云函数
	AlbType *string `json:"AlbType,omitnil,omitempty" name:"AlbType"`

	// IsCdn=3时，需要填此参数，表示自定义header
	IpHeaders []*string `json:"IpHeaders,omitnil,omitempty" name:"IpHeaders"`

	// 规则引擎类型。
	// 1: menshen
	// 2: tiga
	EngineType *int64 `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 云类型。
	// public:公有云
	// private:私有云
	// hybrid:混合云
	CloudType *string `json:"CloudType,omitnil,omitempty" name:"CloudType"`

	// 域名备注信息
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`
}

type HostStatus struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// WAF的开关，1：开，0：关
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type HybridPkg struct {
	// 资源id
	ResourceIds *string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 地域
	Region *int64 `json:"Region,omitnil,omitempty" name:"Region"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 申请数量
	InquireNum *int64 `json:"InquireNum,omitnil,omitempty" name:"InquireNum"`

	// 使用数量
	UsedNum *int64 `json:"UsedNum,omitnil,omitempty" name:"UsedNum"`

	// 续费标志
	RenewFlag *uint64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

// Predefined struct for user
type ImportIpAccessControlRequestParams struct {
	// 导入的IP黑白名单列表
	Data []*IpAccessControlParam `json:"Data,omitnil,omitempty" name:"Data"`

	// 具体域名如：test.qcloudwaf.com
	// 全局域名为：global
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 是否为批量防护IP黑白名单，当为批量防护IP黑白名单时，取值为batch，否则为空
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ImportIpAccessControlRequest struct {
	*tchttp.BaseRequest
	
	// 导入的IP黑白名单列表
	Data []*IpAccessControlParam `json:"Data,omitnil,omitempty" name:"Data"`

	// 具体域名如：test.qcloudwaf.com
	// 全局域名为：global
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 是否为批量防护IP黑白名单，当为批量防护IP黑白名单时，取值为batch，否则为空
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *ImportIpAccessControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportIpAccessControlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	delete(f, "Domain")
	delete(f, "SourceType")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImportIpAccessControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImportIpAccessControlResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImportIpAccessControlResponse struct {
	*tchttp.BaseResponse
	Response *ImportIpAccessControlResponseParams `json:"Response"`
}

func (r *ImportIpAccessControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImportIpAccessControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceInfo struct {
	// 实例唯一ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例对应资源ID，计费使用
	ResourceIds *string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 实例所属地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 付费模式
	PayMode *uint64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 自动续费标识。
	// 0：关闭
	// 1：开启
	RenewFlag *uint64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 弹性计费开关。
	// 0：关闭
	// 1：开启
	Mode *uint64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 实例套餐版本。
	// 101：小微版
	// 102：超轻版
	// 2：高级版
	// 3：企业版
	// 4：旗舰版
	// 6：独享版
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 实例过期时间
	ValidTime *string `json:"ValidTime,omitnil,omitempty" name:"ValidTime"`

	// 实例开始时间
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 已配置域名个数
	DomainCount *uint64 `json:"DomainCount,omitnil,omitempty" name:"DomainCount"`

	// 域名数量上限
	SubDomainLimit *uint64 `json:"SubDomainLimit,omitnil,omitempty" name:"SubDomainLimit"`

	// 已配置主域名个数
	MainDomainCount *uint64 `json:"MainDomainCount,omitnil,omitempty" name:"MainDomainCount"`

	// 主域名数量上限
	MainDomainLimit *uint64 `json:"MainDomainLimit,omitnil,omitempty" name:"MainDomainLimit"`

	// 实例30天内QPS峰值
	MaxQPS *uint64 `json:"MaxQPS,omitnil,omitempty" name:"MaxQPS"`

	// qps扩展包信息
	QPS *QPSPackageNew `json:"QPS,omitnil,omitempty" name:"QPS"`

	// 域名扩展包信息
	DomainPkg *DomainPackageNew `json:"DomainPkg,omitnil,omitempty" name:"DomainPkg"`

	// 用户appid
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// clb或saas
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 业务安全包
	FraudPkg *FraudPkg `json:"FraudPkg,omitnil,omitempty" name:"FraudPkg"`

	// Bot资源包
	BotPkg *BotPkg `json:"BotPkg,omitnil,omitempty" name:"BotPkg"`

	// bot的qps详情
	BotQPS *BotQPS `json:"BotQPS,omitnil,omitempty" name:"BotQPS"`

	// qps弹性计费上限
	ElasticBilling *uint64 `json:"ElasticBilling,omitnil,omitempty" name:"ElasticBilling"`

	// 攻击日志投递开关
	AttackLogPost *int64 `json:"AttackLogPost,omitnil,omitempty" name:"AttackLogPost"`

	// 带宽峰值，单位为B/s(字节每秒)
	MaxBandwidth *uint64 `json:"MaxBandwidth,omitnil,omitempty" name:"MaxBandwidth"`

	// api安全是否购买
	APISecurity *uint64 `json:"APISecurity,omitnil,omitempty" name:"APISecurity"`

	// 购买的qps规格
	QpsStandard *uint64 `json:"QpsStandard,omitnil,omitempty" name:"QpsStandard"`

	// 购买的带宽规格
	BandwidthStandard *uint64 `json:"BandwidthStandard,omitnil,omitempty" name:"BandwidthStandard"`

	// 实例状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例沙箱qps值
	SandboxQps *uint64 `json:"SandboxQps,omitnil,omitempty" name:"SandboxQps"`

	// 是否api 安全试用
	IsAPISecurityTrial *uint64 `json:"IsAPISecurityTrial,omitnil,omitempty" name:"IsAPISecurityTrial"`

	// 重保包
	MajorEventsPkg *MajorEventsPkg `json:"MajorEventsPkg,omitnil,omitempty" name:"MajorEventsPkg"`

	// 混合云子节点包
	HybridPkg *HybridPkg `json:"HybridPkg,omitnil,omitempty" name:"HybridPkg"`

	// API安全资源包
	ApiPkg *ApiPkg `json:"ApiPkg,omitnil,omitempty" name:"ApiPkg"`

	// 小程序安全加速包
	MiniPkg *MiniPkg `json:"MiniPkg,omitnil,omitempty" name:"MiniPkg"`

	// 小程序qps规格
	MiniQpsStandard *uint64 `json:"MiniQpsStandard,omitnil,omitempty" name:"MiniQpsStandard"`

	// 小程序qps峰值
	MiniMaxQPS *uint64 `json:"MiniMaxQPS,omitnil,omitempty" name:"MiniMaxQPS"`

	// 最近一次超量时间
	LastQpsExceedTime *string `json:"LastQpsExceedTime,omitnil,omitempty" name:"LastQpsExceedTime"`

	// 小程序安全接入ID数量扩张包
	MiniExtendPkg *MiniExtendPkg `json:"MiniExtendPkg,omitnil,omitempty" name:"MiniExtendPkg"`

	// 计费项
	BillingItem *string `json:"BillingItem,omitnil,omitempty" name:"BillingItem"`

	// 实例延期释放标识
	FreeDelayFlag *uint64 `json:"FreeDelayFlag,omitnil,omitempty" name:"FreeDelayFlag"`
}

type IpAccessControlData struct {
	// ip黑白名单
	Res []*IpAccessControlItem `json:"Res,omitnil,omitempty" name:"Res"`

	// 计数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type IpAccessControlItem struct {
	// mongo表自增Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 动作
	ActionType *uint64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// ip
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 备注
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 来源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 更新时间戳
	TsVersion *uint64 `json:"TsVersion,omitnil,omitempty" name:"TsVersion"`

	// 有效截止时间戳
	ValidTs *uint64 `json:"ValidTs,omitnil,omitempty" name:"ValidTs"`

	// 生效状态
	ValidStatus *int64 `json:"ValidStatus,omitnil,omitempty" name:"ValidStatus"`

	// 55000001
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// IP列表
	IpList []*string `json:"IpList,omitnil,omitempty" name:"IpList"`

	// 规则创建时间
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 定时任务类型
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 周期任务类型
	CronType *string `json:"CronType,omitnil,omitempty" name:"CronType"`

	// 定时任务配置详情
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`
}

type IpAccessControlParam struct {
	// IP列表
	IpList []*string `json:"IpList,omitnil,omitempty" name:"IpList"`

	// valid_ts为有效日期，值为秒级时间戳（（如1680570420代表2023-04-04 09:07:00））
	ValidTs *uint64 `json:"ValidTs,omitnil,omitempty" name:"ValidTs"`

	// 42为黑名单，40为白名单
	ActionType *int64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 备注
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`
}

type IpHitItem struct {
	// 动作
	Action *uint64 `json:"Action,omitnil,omitempty" name:"Action"`

	// 类别
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// ip
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 时间戳
	TsVersion *uint64 `json:"TsVersion,omitnil,omitempty" name:"TsVersion"`

	// 有效截止时间戳
	ValidTs *uint64 `json:"ValidTs,omitnil,omitempty" name:"ValidTs"`
}

type IpHitItemsData struct {
	// 数组封装
	Res []*IpHitItem `json:"Res,omitnil,omitempty" name:"Res"`

	// 总数目
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type JobDateTime struct {
	// 定时执行的时间参数
	Timed []*TimedJob `json:"Timed,omitnil,omitempty" name:"Timed"`

	// 周期执行的时间参数
	Cron []*CronJob `json:"Cron,omitnil,omitempty" name:"Cron"`

	// 时区
	TimeTZone *string `json:"TimeTZone,omitnil,omitempty" name:"TimeTZone"`
}

type KVInt struct {
	// Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value
	Value *uint64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type LoadBalancer struct {
	// 负载均衡LD的ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡LD的名称
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 负载均衡监听器的ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡监听器的名称
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// 负载均衡实例的IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 负载均衡实例的端口
	Vport *uint64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 负载均衡LD的地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 监听器协议，http、https
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 负载均衡监听器所在的zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 负载均衡的VPCID，公网为-1，内网按实际填写
	NumericalVpcId *int64 `json:"NumericalVpcId,omitnil,omitempty" name:"NumericalVpcId"`

	// 负载均衡的网络类型。OPEN： 公网 INTERNAL ：内网
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// 负载均衡的域名
	LoadBalancerDomain *string `json:"LoadBalancerDomain,omitnil,omitempty" name:"LoadBalancerDomain"`
}

type LoadBalancerPackageNew struct {
	// 监听id
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 监听名
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// 负载均衡id
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡名
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 地区
	// "多伦多": "ca",
	//     "广州": "gz",
	//     "成都": "cd",
	//     "福州": "fzec",
	//     "深圳": "szx",
	//     "印度": "in",
	//     "济南": "jnec",
	//     "重庆": "cq",
	//     "天津": "tsn",
	//     "欧洲东北": "ru",
	//     "南京": "nj",
	//     "美国硅谷": "usw",
	//     "泰国": "th",
	//     "广州Open": "gzopen",
	//     "深圳金融": "szjr",
	//     "法兰克福": "de",
	//     "日本": "jp",
	//     "弗吉尼亚": "use",
	//     "北京": "bj",
	//     "中国香港": "hk",
	//     "杭州": "hzec",
	//     "北京金融": "bjjr",
	//     "上海金融": "shjr",
	//     "台北": "tpe",
	//     "首尔": "kr",
	//     "上海": "sh",
	//     "新加坡": "sg",
	//     "清远": "qy"
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 接入IP
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 接入端口
	Vport *uint64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 地域
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// VPCID
	NumericalVpcId *int64 `json:"NumericalVpcId,omitnil,omitempty" name:"NumericalVpcId"`

	// CLB类型
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// 负载均衡器的域名
	LoadBalancerDomain *string `json:"LoadBalancerDomain,omitnil,omitempty" name:"LoadBalancerDomain"`
}

type LogHistogramInfo struct {
	// 日志条数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 时间戳
	TimeStamp *int64 `json:"TimeStamp,omitnil,omitempty" name:"TimeStamp"`
}

type MajorEventsPkg struct {
	// 资源id
	ResourceIds *string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 地域
	Region *int64 `json:"Region,omitnil,omitempty" name:"Region"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 申请数量
	InquireNum *int64 `json:"InquireNum,omitnil,omitempty" name:"InquireNum"`

	// 使用数量
	UsedNum *int64 `json:"UsedNum,omitnil,omitempty" name:"UsedNum"`

	// 续费标志
	RenewFlag *uint64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 计费项
	BillingItem *string `json:"BillingItem,omitnil,omitempty" name:"BillingItem"`

	// 护网包状态
	HWState *int64 `json:"HWState,omitnil,omitempty" name:"HWState"`
}

type MiniExtendPkg struct {
	// 资源id
	ResourceIds *string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 地域
	Region *int64 `json:"Region,omitnil,omitempty" name:"Region"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 购买数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 续费标志
	RenewFlag *uint64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 计费项
	BillingItem *string `json:"BillingItem,omitnil,omitempty" name:"BillingItem"`
}

type MiniPkg struct {
	// 资源id
	ResourceIds *string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 地域
	Region *int64 `json:"Region,omitnil,omitempty" name:"Region"`

	// 开始时间
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 购买数量
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 续费标志
	RenewFlag *uint64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 计费项
	BillingItem *string `json:"BillingItem,omitnil,omitempty" name:"BillingItem"`
}

// Predefined struct for user
type ModifyAntiFakeUrlRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// uri
	Uri *string `json:"Uri,omitnil,omitempty" name:"Uri"`

	// ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type ModifyAntiFakeUrlRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// uri
	Uri *string `json:"Uri,omitnil,omitempty" name:"Uri"`

	// ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
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
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Id列表
	Ids []*uint64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

type ModifyAntiFakeUrlStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Id列表
	Ids []*uint64 `json:"Ids,omitnil,omitempty" name:"Ids"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyAntiInfoLeakRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Action 值
	ActionType *uint64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 策略数组
	Strategies []*StrategyForAntiInfoLeak `json:"Strategies,omitnil,omitempty" name:"Strategies"`
}

type ModifyAntiInfoLeakRulesRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Action 值
	ActionType *uint64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 策略数组
	Strategies []*StrategyForAntiInfoLeak `json:"Strategies,omitnil,omitempty" name:"Strategies"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要批量开启的实体列表
	TargetList []*TargetEntity `json:"TargetList,omitnil,omitempty" name:"TargetList"`
}

type ModifyApiAnalyzeStatusRequest struct {
	*tchttp.BaseRequest
	
	// 开关状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要批量开启的实体列表
	TargetList []*TargetEntity `json:"TargetList,omitnil,omitempty" name:"TargetList"`
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
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 不支持开启的域名列表
	UnSupportedList []*string `json:"UnSupportedList,omitnil,omitempty" name:"UnSupportedList"`

	// 开启/关闭失败的域名列表
	FailDomainList []*string `json:"FailDomainList,omitnil,omitempty" name:"FailDomainList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyApiSecEventChangeRequestParams struct {
	// 变更状态，1:新发现，2，确认中，3，已确认，4，已下线，5，已忽略
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 处理人
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 备注，有长度显示1k
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 批量操作的事件列表
	EventIdList []*string `json:"EventIdList,omitnil,omitempty" name:"EventIdList"`

	// 批量操作的api列表
	ApiNameList []*ApiSecKey `json:"ApiNameList,omitnil,omitempty" name:"ApiNameList"`

	// 判断是否删除，包括删除事件和删除资产
	IsDelete *bool `json:"IsDelete,omitnil,omitempty" name:"IsDelete"`

	// 判断是否是更新api的备注，更新api备注的时候，为true
	UpdateApiRemark *bool `json:"UpdateApiRemark,omitnil,omitempty" name:"UpdateApiRemark"`
}

type ModifyApiSecEventChangeRequest struct {
	*tchttp.BaseRequest
	
	// 变更状态，1:新发现，2，确认中，3，已确认，4，已下线，5，已忽略
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 处理人
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 备注，有长度显示1k
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 批量操作的事件列表
	EventIdList []*string `json:"EventIdList,omitnil,omitempty" name:"EventIdList"`

	// 批量操作的api列表
	ApiNameList []*ApiSecKey `json:"ApiNameList,omitnil,omitempty" name:"ApiNameList"`

	// 判断是否删除，包括删除事件和删除资产
	IsDelete *bool `json:"IsDelete,omitnil,omitempty" name:"IsDelete"`

	// 判断是否是更新api的备注，更新api备注的时候，为true
	UpdateApiRemark *bool `json:"UpdateApiRemark,omitnil,omitempty" name:"UpdateApiRemark"`
}

func (r *ModifyApiSecEventChangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiSecEventChangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	delete(f, "UserName")
	delete(f, "Remark")
	delete(f, "EventIdList")
	delete(f, "ApiNameList")
	delete(f, "IsDelete")
	delete(f, "UpdateApiRemark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApiSecEventChangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApiSecEventChangeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApiSecEventChangeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApiSecEventChangeResponseParams `json:"Response"`
}

func (r *ModifyApiSecEventChangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApiSecEventChangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAreaBanAreasRequestParams struct {
	// 需要修改的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 需要调整的地域信息，一个字符串数组
	Areas []*string `json:"Areas,omitnil,omitempty" name:"Areas"`

	// 规则执行的方式，TimedJob为定时执行，CronJob为周期执行
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 定时任务配置
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`
}

type ModifyAreaBanAreasRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 需要调整的地域信息，一个字符串数组
	Areas []*string `json:"Areas,omitnil,omitempty" name:"Areas"`

	// 规则执行的方式，TimedJob为定时执行，CronJob为周期执行
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 定时任务配置
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`
}

func (r *ModifyAreaBanAreasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAreaBanAreasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Areas")
	delete(f, "JobType")
	delete(f, "JobDateTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAreaBanAreasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAreaBanAreasResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAreaBanAreasResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAreaBanAreasResponseParams `json:"Response"`
}

func (r *ModifyAreaBanAreasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAreaBanAreasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAreaBanStatusRequestParams struct {
	// 需要修改的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 状态值，0表示关闭，1表示开启
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyAreaBanStatusRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 状态值，0表示关闭，1表示开启
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 匹配规则项列表
	Rules []*UserWhiteRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 规则Id
	SignatureId *string `json:"SignatureId,omitnil,omitempty" name:"SignatureId"`

	// 编辑的加白的规则ID列表
	SignatureIds []*string `json:"SignatureIds,omitnil,omitempty" name:"SignatureIds"`

	// 加白的大类规则ID
	TypeIds []*string `json:"TypeIds,omitnil,omitempty" name:"TypeIds"`

	// 0表示按照特定规则ID加白, 1表示按照规则类型加白
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 规则名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type ModifyAttackWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则序号
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 匹配规则项列表
	Rules []*UserWhiteRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 规则Id
	SignatureId *string `json:"SignatureId,omitnil,omitempty" name:"SignatureId"`

	// 编辑的加白的规则ID列表
	SignatureIds []*string `json:"SignatureIds,omitnil,omitempty" name:"SignatureIds"`

	// 加白的大类规则ID
	TypeIds []*string `json:"TypeIds,omitnil,omitempty" name:"TypeIds"`

	// 0表示按照特定规则ID加白, 1表示按照规则类型加白
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 规则名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	delete(f, "Status")
	delete(f, "Rules")
	delete(f, "SignatureId")
	delete(f, "SignatureIds")
	delete(f, "TypeIds")
	delete(f, "Mode")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAttackWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAttackWhiteRuleResponseParams struct {
	// 规则总数
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 类别
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 是否是bot4.0版本
	IsVersionFour *bool `json:"IsVersionFour,omitnil,omitempty" name:"IsVersionFour"`

	// 传入Bot版本号，场景化版本为"4.1.0"
	BotVersion *string `json:"BotVersion,omitnil,omitempty" name:"BotVersion"`

	// 批量开启BOT开关的域名列表
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`
}

type ModifyBotStatusRequest struct {
	*tchttp.BaseRequest
	
	// 类别
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 是否是bot4.0版本
	IsVersionFour *bool `json:"IsVersionFour,omitnil,omitempty" name:"IsVersionFour"`

	// 传入Bot版本号，场景化版本为"4.1.0"
	BotVersion *string `json:"BotVersion,omitnil,omitempty" name:"BotVersion"`

	// 批量开启BOT开关的域名列表
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`
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
	delete(f, "Category")
	delete(f, "Status")
	delete(f, "Domain")
	delete(f, "InstanceID")
	delete(f, "IsVersionFour")
	delete(f, "BotVersion")
	delete(f, "DomainList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBotStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBotStatusResponseParams struct {
	// 正常情况为null
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 未购买BOT的域名列表
	UnSupportedList []*string `json:"UnSupportedList,omitnil,omitempty" name:"UnSupportedList"`

	// 已购买但操作失败的域名列表
	FailDomainList []*string `json:"FailDomainList,omitnil,omitempty" name:"FailDomainList"`

	// 成功数目
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 编辑的规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 编辑的规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 动作类型，1代表阻断，2代表人机识别，3代表观察，4代表重定向，5代表JS校验
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// 匹配条件数组
	Strategies []*Strategy `json:"Strategies,omitnil,omitempty" name:"Strategies"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 动作为重定向的时候重定向URL，默认为"/"
	Redirect *string `json:"Redirect,omitnil,omitempty" name:"Redirect"`

	// 放行时是否继续执行其它检查逻辑，继续执行地域封禁防护：geoip、继续执行CC策略防护：cc、继续执行WEB应用防护：owasp、继续执行AI引擎防护：ai、继续执行信息防泄漏防护：antileakage。如果多个勾选那么以,串接。
	// 默认是"geoip,cc,owasp,ai,antileakage"
	Bypass *string `json:"Bypass,omitnil,omitempty" name:"Bypass"`

	// 优先级，1~100的整数，数字越小，代表这条规则的执行优先级越高。
	// 默认是100
	SortId *uint64 `json:"SortId,omitnil,omitempty" name:"SortId"`

	// 规则生效截止时间，0：永久生效，其它值为对应时间的时间戳。
	// 默认是0
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 规则执行的方式，TimedJob为定时执行，CronJob为周期执行
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 定时任务配置
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`

	// 规则来源，判断是不是小程序的
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 开关状态，小程序风控规则的时候传该值
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 拦截页面id
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`
}

type ModifyCustomRuleRequest struct {
	*tchttp.BaseRequest
	
	// 编辑的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 编辑的规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 编辑的规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 动作类型，1代表阻断，2代表人机识别，3代表观察，4代表重定向，5代表JS校验
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// 匹配条件数组
	Strategies []*Strategy `json:"Strategies,omitnil,omitempty" name:"Strategies"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 动作为重定向的时候重定向URL，默认为"/"
	Redirect *string `json:"Redirect,omitnil,omitempty" name:"Redirect"`

	// 放行时是否继续执行其它检查逻辑，继续执行地域封禁防护：geoip、继续执行CC策略防护：cc、继续执行WEB应用防护：owasp、继续执行AI引擎防护：ai、继续执行信息防泄漏防护：antileakage。如果多个勾选那么以,串接。
	// 默认是"geoip,cc,owasp,ai,antileakage"
	Bypass *string `json:"Bypass,omitnil,omitempty" name:"Bypass"`

	// 优先级，1~100的整数，数字越小，代表这条规则的执行优先级越高。
	// 默认是100
	SortId *uint64 `json:"SortId,omitnil,omitempty" name:"SortId"`

	// 规则生效截止时间，0：永久生效，其它值为对应时间的时间戳。
	// 默认是0
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 规则执行的方式，TimedJob为定时执行，CronJob为周期执行
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 定时任务配置
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`

	// 规则来源，判断是不是小程序的
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 开关状态，小程序风控规则的时候传该值
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 拦截页面id
	PageId *string `json:"PageId,omitnil,omitempty" name:"PageId"`
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
	delete(f, "JobType")
	delete(f, "JobDateTime")
	delete(f, "Source")
	delete(f, "Status")
	delete(f, "PageId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomRuleResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	Success *ResponseCode `json:"Success,omitnil,omitempty" name:"Success"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 开关的状态，1是开启、0是关闭
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 规则id
	DomainRuleIdList []*DomainRuleId `json:"DomainRuleIdList,omitnil,omitempty" name:"DomainRuleIdList"`
}

type ModifyCustomRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 开关的状态，1是开启、0是关闭
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 规则id
	DomainRuleIdList []*DomainRuleId `json:"DomainRuleIdList,omitnil,omitempty" name:"DomainRuleIdList"`
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
	delete(f, "DomainRuleIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomRuleStatusResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	Success *ResponseCode `json:"Success,omitnil,omitempty" name:"Success"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 编辑的规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 编辑的规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 放行时是否继续执行其它检查逻辑，继续执行地域封禁防护：geoip、继续执行CC策略防护：cc、继续执行WEB应用防护：owasp、继续执行AI引擎防护：ai、继续执行信息防泄漏防护：antileakage。如果勾选多个，则以“，”串接。
	Bypass *string `json:"Bypass,omitnil,omitempty" name:"Bypass"`

	// 优先级，1~100的整数，数字越小，代表这条规则的执行优先级越高。
	SortId *uint64 `json:"SortId,omitnil,omitempty" name:"SortId"`

	// 规则生效截止时间，0：永久生效，其它值为对应时间的时间戳。
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 匹配条件数组
	Strategies []*Strategy `json:"Strategies,omitnil,omitempty" name:"Strategies"`

	// 规则执行的方式，TimedJob为定时执行，CronJob为周期执行
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 定时任务配置
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`
}

type ModifyCustomWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 编辑的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 编辑的规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 编辑的规则名称
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 放行时是否继续执行其它检查逻辑，继续执行地域封禁防护：geoip、继续执行CC策略防护：cc、继续执行WEB应用防护：owasp、继续执行AI引擎防护：ai、继续执行信息防泄漏防护：antileakage。如果勾选多个，则以“，”串接。
	Bypass *string `json:"Bypass,omitnil,omitempty" name:"Bypass"`

	// 优先级，1~100的整数，数字越小，代表这条规则的执行优先级越高。
	SortId *uint64 `json:"SortId,omitnil,omitempty" name:"SortId"`

	// 规则生效截止时间，0：永久生效，其它值为对应时间的时间戳。
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 匹配条件数组
	Strategies []*Strategy `json:"Strategies,omitnil,omitempty" name:"Strategies"`

	// 规则执行的方式，TimedJob为定时执行，CronJob为周期执行
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 定时任务配置
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`
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
	delete(f, "JobType")
	delete(f, "JobDateTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomWhiteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomWhiteRuleResponseParams struct {
	// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
	Success *ResponseCode `json:"Success,omitnil,omitempty" name:"Success"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 开关的状态，1是开启、0是关闭
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyCustomWhiteRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 开关的状态，1是开启、0是关闭
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
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
	Success *ResponseCode `json:"Success,omitnil,omitempty" name:"Success"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要修改的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 需要修改的域名ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 修改域名的Ipv6开关为Status （1:开启 2:关闭）
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyDomainIpv6StatusRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的域名所属的实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要修改的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 需要修改的域名ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 修改域名的Ipv6开关为Status （1:开启 2:关闭）
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
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
	Ipv6Status *int64 `json:"Ipv6Status,omitnil,omitempty" name:"Ipv6Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyDomainPostActionRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 0-关闭投递，1-开启投递
	PostCLSAction *int64 `json:"PostCLSAction,omitnil,omitempty" name:"PostCLSAction"`

	// 0-关闭投递，1-开启投递
	PostCKafkaAction *int64 `json:"PostCKafkaAction,omitnil,omitempty" name:"PostCKafkaAction"`
}

type ModifyDomainPostActionRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 0-关闭投递，1-开启投递
	PostCLSAction *int64 `json:"PostCLSAction,omitnil,omitempty" name:"PostCLSAction"`

	// 0-关闭投递，1-开启投递
	PostCKafkaAction *int64 `json:"PostCKafkaAction,omitnil,omitempty" name:"PostCKafkaAction"`
}

func (r *ModifyDomainPostActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainPostActionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "PostCLSAction")
	delete(f, "PostCKafkaAction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainPostActionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainPostActionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDomainPostActionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainPostActionResponseParams `json:"Response"`
}

func (r *ModifyDomainPostActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainPostActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainWhiteRuleRequestParams struct {
	// 需要更改的规则的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 白名单id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 规则的id列表
	Rules []*uint64 `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 规则匹配路径
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 规则匹配方法
	Function *string `json:"Function,omitnil,omitempty" name:"Function"`

	// 规则的开关状态，0表示关闭开关，1表示打开开关
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyDomainWhiteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要更改的规则的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 白名单id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 规则的id列表
	Rules []*uint64 `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 规则匹配路径
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 规则匹配方法
	Function *string `json:"Function,omitnil,omitempty" name:"Function"`

	// 规则的开关状态，0表示关闭开关，1表示打开开关
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domains []*DomainURI `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 修改域名的访问日志开关为Status
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyDomainsCLSStatusRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的域名列表
	Domains []*DomainURI `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 修改域名的访问日志开关为Status
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Goods []*Goods `json:"Goods,omitnil,omitempty" name:"Goods"`
}

type ModifyGenerateDealsRequest struct {
	*tchttp.BaseRequest
	
	// 计费下单入参
	Goods []*Goods `json:"Goods,omitnil,omitempty" name:"Goods"`
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
	Data *DealData `json:"Data,omitnil,omitempty" name:"Data"`

	// 1:成功，0:失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 返回message
	ReturnMessage *string `json:"ReturnMessage,omitnil,omitempty" name:"ReturnMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// WAF流量模式。
	// 0：镜像模式（默认）
	// 1：清洗模式
	FlowMode *uint64 `json:"FlowMode,omitnil,omitempty" name:"FlowMode"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type ModifyHostFlowModeRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// WAF流量模式。
	// 0：镜像模式（默认）
	// 1：清洗模式
	FlowMode *uint64 `json:"FlowMode,omitnil,omitempty" name:"FlowMode"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
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
	Success *ResponseCode `json:"Success,omitnil,omitempty" name:"Success"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 防护状态：
	// 10：规则观察&&AI关闭模式，11：规则观察&&AI观察模式，12：规则观察&&AI拦截模式
	// 20：规则拦截&&AI关闭模式，21：规则拦截&&AI观察模式，22：规则拦截&&AI拦截模式
	Mode *uint64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 0:修改防护模式，1:修改AI
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 实例类型
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`
}

type ModifyHostModeRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 防护状态：
	// 10：规则观察&&AI关闭模式，11：规则观察&&AI观察模式，12：规则观察&&AI拦截模式
	// 20：规则拦截&&AI关闭模式，21：规则拦截&&AI观察模式，22：规则拦截&&AI拦截模式
	Mode *uint64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 0:修改防护模式，1:修改AI
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 实例ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 实例类型
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`
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
	Success *ResponseCode `json:"Success,omitnil,omitempty" name:"Success"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Host *HostRecord `json:"Host,omitnil,omitempty" name:"Host"`

	// 实例唯一ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type ModifyHostRequest struct {
	*tchttp.BaseRequest
	
	// 编辑的域名配置信息
	Host *HostRecord `json:"Host,omitnil,omitempty" name:"Host"`

	// 实例唯一ID
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
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
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	HostsStatus []*HostStatus `json:"HostsStatus,omitnil,omitempty" name:"HostsStatus"`
}

type ModifyHostStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名状态列表
	HostsStatus []*HostStatus `json:"HostsStatus,omitnil,omitempty" name:"HostsStatus"`
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
	Success *ResponseCode `json:"Success,omitnil,omitempty" name:"Success"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 弹性计费开关
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`
}

type ModifyInstanceElasticModeRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 弹性计费开关
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 实例版本，支持clb-waf、sparta-waf
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`
}

type ModifyInstanceNameRequest struct {
	*tchttp.BaseRequest
	
	// 新名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 实例版本，支持clb-waf、sparta-waf
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`
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
	ModifyCode *int64 `json:"ModifyCode,omitnil,omitempty" name:"ModifyCode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// qps上限
	QpsLimit *int64 `json:"QpsLimit,omitnil,omitempty" name:"QpsLimit"`
}

type ModifyInstanceQpsLimitRequest struct {
	*tchttp.BaseRequest
	
	// 套餐实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// qps上限
	QpsLimit *int64 `json:"QpsLimit,omitnil,omitempty" name:"QpsLimit"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 续费开关
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type ModifyInstanceRenewFlagRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 续费开关
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyIpAccessControlRequestParams struct {
	// 具体域名如：test.qcloudwaf.com
	// 全局域名为：global
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// ip参数列表
	IpList []*string `json:"IpList,omitnil,omitempty" name:"IpList"`

	// 42为黑名单，40为白名单
	ActionType *int64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// valid_ts为有效日期，值为秒级时间戳（（如1680570420代表2023-04-04 09:07:00））
	//
	// Deprecated: ValidTS is deprecated.
	ValidTS *int64 `json:"ValidTS,omitnil,omitempty" name:"ValidTS"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 是否为批量防护IP黑白名单，当为批量防护IP黑白名单时，取值为batch，否则为空
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 备注
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 规则执行的方式，TimedJob为定时执行，CronJob为周期执行
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 定时配置详情
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`
}

type ModifyIpAccessControlRequest struct {
	*tchttp.BaseRequest
	
	// 具体域名如：test.qcloudwaf.com
	// 全局域名为：global
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// ip参数列表
	IpList []*string `json:"IpList,omitnil,omitempty" name:"IpList"`

	// 42为黑名单，40为白名单
	ActionType *int64 `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// valid_ts为有效日期，值为秒级时间戳（（如1680570420代表2023-04-04 09:07:00））
	ValidTS *int64 `json:"ValidTS,omitnil,omitempty" name:"ValidTS"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 是否为批量防护IP黑白名单，当为批量防护IP黑白名单时，取值为batch，否则为空
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 备注
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 规则执行的方式，TimedJob为定时执行，CronJob为周期执行
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 定时配置详情
	JobDateTime *JobDateTime `json:"JobDateTime,omitnil,omitempty" name:"JobDateTime"`
}

func (r *ModifyIpAccessControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIpAccessControlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "IpList")
	delete(f, "ActionType")
	delete(f, "RuleId")
	delete(f, "ValidTS")
	delete(f, "InstanceId")
	delete(f, "Edition")
	delete(f, "SourceType")
	delete(f, "Note")
	delete(f, "JobType")
	delete(f, "JobDateTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIpAccessControlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIpAccessControlResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyIpAccessControlResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIpAccessControlResponseParams `json:"Response"`
}

func (r *ModifyIpAccessControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIpAccessControlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModuleStatusRequestParams struct {
	// 需要设置的domain
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Web 安全模块开关，0或1
	WebSecurity *uint64 `json:"WebSecurity,omitnil,omitempty" name:"WebSecurity"`

	// 访问控制模块开关，0或者1
	AccessControl *uint64 `json:"AccessControl,omitnil,omitempty" name:"AccessControl"`

	// CC模块开关，0或者1
	CcProtection *uint64 `json:"CcProtection,omitnil,omitempty" name:"CcProtection"`

	// API安全模块开关，0或者1
	ApiProtection *uint64 `json:"ApiProtection,omitnil,omitempty" name:"ApiProtection"`

	// 防篡改模块开关，0或者1
	AntiTamper *uint64 `json:"AntiTamper,omitnil,omitempty" name:"AntiTamper"`

	// 防泄漏模块开关，0或者1
	AntiLeakage *uint64 `json:"AntiLeakage,omitnil,omitempty" name:"AntiLeakage"`

	// 限流模块开关，0或1
	RateLimit *uint64 `json:"RateLimit,omitnil,omitempty" name:"RateLimit"`
}

type ModifyModuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 需要设置的domain
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// Web 安全模块开关，0或1
	WebSecurity *uint64 `json:"WebSecurity,omitnil,omitempty" name:"WebSecurity"`

	// 访问控制模块开关，0或者1
	AccessControl *uint64 `json:"AccessControl,omitnil,omitempty" name:"AccessControl"`

	// CC模块开关，0或者1
	CcProtection *uint64 `json:"CcProtection,omitnil,omitempty" name:"CcProtection"`

	// API安全模块开关，0或者1
	ApiProtection *uint64 `json:"ApiProtection,omitnil,omitempty" name:"ApiProtection"`

	// 防篡改模块开关，0或者1
	AntiTamper *uint64 `json:"AntiTamper,omitnil,omitempty" name:"AntiTamper"`

	// 防泄漏模块开关，0或者1
	AntiLeakage *uint64 `json:"AntiLeakage,omitnil,omitempty" name:"AntiLeakage"`

	// 限流模块开关，0或1
	RateLimit *uint64 `json:"RateLimit,omitnil,omitempty" name:"RateLimit"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ObjectId *string `json:"ObjectId,omitnil,omitempty" name:"ObjectId"`

	// 改动作类型:Status修改开关，InstanceId绑定实例, Proxy设置代理状态
	OpType *string `json:"OpType,omitnil,omitempty" name:"OpType"`

	// 新的Waf开关状态，如果和已有状态相同认为修改成功
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 新的实例ID，如果和已绑定的实例相同认为修改成功
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否开启代理，0:不开启,1:以XFF的第一个IP地址作为客户端IP,2:以remote_addr作为客户端IP,3:从指定的头部字段获取客户端IP，字段通过IpHeaders字段给出(OpType为Status或Proxy时，该值有效)
	Proxy *uint64 `json:"Proxy,omitnil,omitempty" name:"Proxy"`

	// IsCdn=3时，需要填此参数，表示自定义header(OpType为Status或Proxy时，该值有效)
	IpHeaders []*string `json:"IpHeaders,omitnil,omitempty" name:"IpHeaders"`
}

type ModifyObjectRequest struct {
	*tchttp.BaseRequest
	
	// 修改对象标识
	ObjectId *string `json:"ObjectId,omitnil,omitempty" name:"ObjectId"`

	// 改动作类型:Status修改开关，InstanceId绑定实例, Proxy设置代理状态
	OpType *string `json:"OpType,omitnil,omitempty" name:"OpType"`

	// 新的Waf开关状态，如果和已有状态相同认为修改成功
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 新的实例ID，如果和已绑定的实例相同认为修改成功
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 是否开启代理，0:不开启,1:以XFF的第一个IP地址作为客户端IP,2:以remote_addr作为客户端IP,3:从指定的头部字段获取客户端IP，字段通过IpHeaders字段给出(OpType为Status或Proxy时，该值有效)
	Proxy *uint64 `json:"Proxy,omitnil,omitempty" name:"Proxy"`

	// IsCdn=3时，需要填此参数，表示自定义header(OpType为Status或Proxy时，该值有效)
	IpHeaders []*string `json:"IpHeaders,omitnil,omitempty" name:"IpHeaders"`
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
	delete(f, "Proxy")
	delete(f, "IpHeaders")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyObjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyObjectResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`
}

type ModifyProtectionStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 防护状态：
	// 10：规则观察&&AI关闭模式，11：规则观察&&AI观察模式，12：规则观察&&AI拦截模式
	// 20：规则拦截&&AI关闭模式，21：规则拦截&&AI观察模式，22：规则拦截&&AI拦截模式
	Mode *uint64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 0是修改规则引擎状态，1是修改AI的状态
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type ModifySpartaProtectionModeRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 防护状态：
	// 10：规则观察&&AI关闭模式，11：规则观察&&AI观察模式，12：规则观察&&AI拦截模式
	// 20：规则拦截&&AI关闭模式，21：规则拦截&&AI观察模式，22：规则拦截&&AI拦截模式
	Mode *uint64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 0是修改规则引擎状态，1是修改AI的状态
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
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
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySpartaProtectionModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySpartaProtectionModeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 必填项。域名唯一ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 必填项。域名所属实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 证书类型。0：仅配置HTTP监听端口，没有证书1：证书来源为自有证书2：证书来源为托管证书
	CertType *int64 `json:"CertType,omitnil,omitempty" name:"CertType"`

	// CertType为1时，需要填充此参数，表示自有证书的证书链
	Cert *string `json:"Cert,omitnil,omitempty" name:"Cert"`

	// CertType为1时，需要填充此参数，表示自有证书的私钥
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`

	// CertType为2时，需要填充此参数，表示腾讯云SSL平台托管的证书id
	SSLId *string `json:"SSLId,omitnil,omitempty" name:"SSLId"`

	// waf前是否部署有七层代理服务。0：没有部署代理服务1：有部署代理服务，waf将使用XFF获取客户端IP2：有部署代理服务，waf将使用remote_addr获取客户端IP3：有部署代理服务，waf将使用ip_headers中的自定义header获取客户端IP
	IsCdn *int64 `json:"IsCdn,omitnil,omitempty" name:"IsCdn"`

	// 服务配置有HTTPS端口时，HTTPS的回源协议。
	// http：使用http协议回源，和HttpsUpstreamPort配合使用
	// https：使用https协议回源
	UpstreamScheme *string `json:"UpstreamScheme,omitnil,omitempty" name:"UpstreamScheme"`

	// HTTPS回源端口,仅UpstreamScheme为http时需要填当前字段
	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitnil,omitempty" name:"HttpsUpstreamPort"`

	// 是否开启HTTP强制跳转到HTTPS。0：不强制跳转1：开启强制跳转
	HttpsRewrite *uint64 `json:"HttpsRewrite,omitnil,omitempty" name:"HttpsRewrite"`

	// 回源类型。0：通过IP回源1：通过域名回源
	UpstreamType *int64 `json:"UpstreamType,omitnil,omitempty" name:"UpstreamType"`

	// 域名回源时的回源域名。UpstreamType为1时，需要填充此字段
	UpstreamDomain *string `json:"UpstreamDomain,omitnil,omitempty" name:"UpstreamDomain"`

	// IP回源时的回源IP列表。UpstreamType为0时，需要填充此字段
	SrcList []*string `json:"SrcList,omitnil,omitempty" name:"SrcList"`

	// 是否开启HTTP2，需要开启HTTPS协议支持。0：关闭1：开启
	IsHttp2 *int64 `json:"IsHttp2,omitnil,omitempty" name:"IsHttp2"`

	// 是否开启WebSocket支持。0：关闭1：开启
	IsWebsocket *int64 `json:"IsWebsocket,omitnil,omitempty" name:"IsWebsocket"`

	// 回源负载均衡策略。0：轮询1：IP hash2：加权轮询
	LoadBalance *int64 `json:"LoadBalance,omitnil,omitempty" name:"LoadBalance"`

	// 待废弃，可不填。是否开启灰度，0表示不开启灰度。
	IsGray *int64 `json:"IsGray,omitnil,omitempty" name:"IsGray"`

	// 域名所属实例类型
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 端口信息，可通过DescribeDomains接口获取具体参数信息。
	Ports []*SpartaProtectionPort `json:"Ports,omitnil,omitempty" name:"Ports"`

	// 是否开启长连接。0： 短连接1： 长连接
	IsKeepAlive *string `json:"IsKeepAlive,omitnil,omitempty" name:"IsKeepAlive"`

	// 待废弃。目前填0即可。anycast IP类型开关： 0 普通IP 1 Anycast IP
	Anycast *int64 `json:"Anycast,omitnil,omitempty" name:"Anycast"`

	// 回源IP列表各IP的权重，和SrcList一一对应。当且仅当UpstreamType为0，并且SrcList有多个IP，并且LoadBalance为2时需要填写，否则填 []
	Weights []*int64 `json:"Weights,omitnil,omitempty" name:"Weights"`

	// 是否开启主动健康检测。0：不开启1：开启
	ActiveCheck *int64 `json:"ActiveCheck,omitnil,omitempty" name:"ActiveCheck"`

	// TLS版本信息
	TLSVersion *int64 `json:"TLSVersion,omitnil,omitempty" name:"TLSVersion"`

	// 加密套件信息
	Ciphers []*int64 `json:"Ciphers,omitnil,omitempty" name:"Ciphers"`

	// 加密套件模板。0：不支持选择，使用默认模板  1：通用型模板 2：安全型模板3：自定义模板
	CipherTemplate *int64 `json:"CipherTemplate,omitnil,omitempty" name:"CipherTemplate"`

	// WAF与源站的读超时时间，默认300s。
	ProxyReadTimeout *int64 `json:"ProxyReadTimeout,omitnil,omitempty" name:"ProxyReadTimeout"`

	// WAF与源站的写超时时间，默认300s。
	ProxySendTimeout *int64 `json:"ProxySendTimeout,omitnil,omitempty" name:"ProxySendTimeout"`

	// WAF回源时的SNI类型。
	// 0：关闭SNI，不配置client_hello中的server_name
	// 1：开启SNI，client_hello中的server_name为防护域名
	// 2：开启SNI，SNI为域名回源时的源站域名
	// 3：开启SNI，SNI为自定义域名
	SniType *int64 `json:"SniType,omitnil,omitempty" name:"SniType"`

	// SniType为3时，需要填此参数，表示自定义的SNI；
	SniHost *string `json:"SniHost,omitnil,omitempty" name:"SniHost"`

	// IsCdn=3时，需要填此参数，表示自定义header
	IpHeaders []*string `json:"IpHeaders,omitnil,omitempty" name:"IpHeaders"`

	// 是否开启XFF重置。0：关闭1：开启
	XFFReset *int64 `json:"XFFReset,omitnil,omitempty" name:"XFFReset"`

	// 域名备注信息
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 自定义回源Host。默认为空字符串，表示使用防护域名作为回源Host。
	UpstreamHost *string `json:"UpstreamHost,omitnil,omitempty" name:"UpstreamHost"`

	// 是否开启缓存。 0：关闭 1：开启
	ProxyBuffer *int64 `json:"ProxyBuffer,omitnil,omitempty" name:"ProxyBuffer"`

	// 是否开启拨测。 0: 禁用拨测 1: 启用拨测。默认启用拨测
	ProbeStatus *int64 `json:"ProbeStatus,omitnil,omitempty" name:"ProbeStatus"`

	// 国密选项。0：不开启国密 1：在原有TLS选项的基础上追加支持国密 2：开启国密并仅支持国密客户端访问
	GmType *int64 `json:"GmType,omitnil,omitempty" name:"GmType"`

	// 国密证书类型。0：无国密证书 1：证书来源为自有国密证书 2：证书来源为托管国密证书
	GmCertType *int64 `json:"GmCertType,omitnil,omitempty" name:"GmCertType"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的证书链
	GmCert *string `json:"GmCert,omitnil,omitempty" name:"GmCert"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的私钥
	GmPrivateKey *string `json:"GmPrivateKey,omitnil,omitempty" name:"GmPrivateKey"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的加密证书
	GmEncCert *string `json:"GmEncCert,omitnil,omitempty" name:"GmEncCert"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的加密证书的私钥
	GmEncPrivateKey *string `json:"GmEncPrivateKey,omitnil,omitempty" name:"GmEncPrivateKey"`

	// GmCertType为2时，需要填充此参数，表示腾讯云SSL平台托管的证书id
	GmSSLId *string `json:"GmSSLId,omitnil,omitempty" name:"GmSSLId"`
}

type ModifySpartaProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 必填项。域名唯一ID
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 必填项。域名所属实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 证书类型。0：仅配置HTTP监听端口，没有证书1：证书来源为自有证书2：证书来源为托管证书
	CertType *int64 `json:"CertType,omitnil,omitempty" name:"CertType"`

	// CertType为1时，需要填充此参数，表示自有证书的证书链
	Cert *string `json:"Cert,omitnil,omitempty" name:"Cert"`

	// CertType为1时，需要填充此参数，表示自有证书的私钥
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`

	// CertType为2时，需要填充此参数，表示腾讯云SSL平台托管的证书id
	SSLId *string `json:"SSLId,omitnil,omitempty" name:"SSLId"`

	// waf前是否部署有七层代理服务。0：没有部署代理服务1：有部署代理服务，waf将使用XFF获取客户端IP2：有部署代理服务，waf将使用remote_addr获取客户端IP3：有部署代理服务，waf将使用ip_headers中的自定义header获取客户端IP
	IsCdn *int64 `json:"IsCdn,omitnil,omitempty" name:"IsCdn"`

	// 服务配置有HTTPS端口时，HTTPS的回源协议。
	// http：使用http协议回源，和HttpsUpstreamPort配合使用
	// https：使用https协议回源
	UpstreamScheme *string `json:"UpstreamScheme,omitnil,omitempty" name:"UpstreamScheme"`

	// HTTPS回源端口,仅UpstreamScheme为http时需要填当前字段
	HttpsUpstreamPort *string `json:"HttpsUpstreamPort,omitnil,omitempty" name:"HttpsUpstreamPort"`

	// 是否开启HTTP强制跳转到HTTPS。0：不强制跳转1：开启强制跳转
	HttpsRewrite *uint64 `json:"HttpsRewrite,omitnil,omitempty" name:"HttpsRewrite"`

	// 回源类型。0：通过IP回源1：通过域名回源
	UpstreamType *int64 `json:"UpstreamType,omitnil,omitempty" name:"UpstreamType"`

	// 域名回源时的回源域名。UpstreamType为1时，需要填充此字段
	UpstreamDomain *string `json:"UpstreamDomain,omitnil,omitempty" name:"UpstreamDomain"`

	// IP回源时的回源IP列表。UpstreamType为0时，需要填充此字段
	SrcList []*string `json:"SrcList,omitnil,omitempty" name:"SrcList"`

	// 是否开启HTTP2，需要开启HTTPS协议支持。0：关闭1：开启
	IsHttp2 *int64 `json:"IsHttp2,omitnil,omitempty" name:"IsHttp2"`

	// 是否开启WebSocket支持。0：关闭1：开启
	IsWebsocket *int64 `json:"IsWebsocket,omitnil,omitempty" name:"IsWebsocket"`

	// 回源负载均衡策略。0：轮询1：IP hash2：加权轮询
	LoadBalance *int64 `json:"LoadBalance,omitnil,omitempty" name:"LoadBalance"`

	// 待废弃，可不填。是否开启灰度，0表示不开启灰度。
	IsGray *int64 `json:"IsGray,omitnil,omitempty" name:"IsGray"`

	// 域名所属实例类型
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 端口信息，可通过DescribeDomains接口获取具体参数信息。
	Ports []*SpartaProtectionPort `json:"Ports,omitnil,omitempty" name:"Ports"`

	// 是否开启长连接。0： 短连接1： 长连接
	IsKeepAlive *string `json:"IsKeepAlive,omitnil,omitempty" name:"IsKeepAlive"`

	// 待废弃。目前填0即可。anycast IP类型开关： 0 普通IP 1 Anycast IP
	Anycast *int64 `json:"Anycast,omitnil,omitempty" name:"Anycast"`

	// 回源IP列表各IP的权重，和SrcList一一对应。当且仅当UpstreamType为0，并且SrcList有多个IP，并且LoadBalance为2时需要填写，否则填 []
	Weights []*int64 `json:"Weights,omitnil,omitempty" name:"Weights"`

	// 是否开启主动健康检测。0：不开启1：开启
	ActiveCheck *int64 `json:"ActiveCheck,omitnil,omitempty" name:"ActiveCheck"`

	// TLS版本信息
	TLSVersion *int64 `json:"TLSVersion,omitnil,omitempty" name:"TLSVersion"`

	// 加密套件信息
	Ciphers []*int64 `json:"Ciphers,omitnil,omitempty" name:"Ciphers"`

	// 加密套件模板。0：不支持选择，使用默认模板  1：通用型模板 2：安全型模板3：自定义模板
	CipherTemplate *int64 `json:"CipherTemplate,omitnil,omitempty" name:"CipherTemplate"`

	// WAF与源站的读超时时间，默认300s。
	ProxyReadTimeout *int64 `json:"ProxyReadTimeout,omitnil,omitempty" name:"ProxyReadTimeout"`

	// WAF与源站的写超时时间，默认300s。
	ProxySendTimeout *int64 `json:"ProxySendTimeout,omitnil,omitempty" name:"ProxySendTimeout"`

	// WAF回源时的SNI类型。
	// 0：关闭SNI，不配置client_hello中的server_name
	// 1：开启SNI，client_hello中的server_name为防护域名
	// 2：开启SNI，SNI为域名回源时的源站域名
	// 3：开启SNI，SNI为自定义域名
	SniType *int64 `json:"SniType,omitnil,omitempty" name:"SniType"`

	// SniType为3时，需要填此参数，表示自定义的SNI；
	SniHost *string `json:"SniHost,omitnil,omitempty" name:"SniHost"`

	// IsCdn=3时，需要填此参数，表示自定义header
	IpHeaders []*string `json:"IpHeaders,omitnil,omitempty" name:"IpHeaders"`

	// 是否开启XFF重置。0：关闭1：开启
	XFFReset *int64 `json:"XFFReset,omitnil,omitempty" name:"XFFReset"`

	// 域名备注信息
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`

	// 自定义回源Host。默认为空字符串，表示使用防护域名作为回源Host。
	UpstreamHost *string `json:"UpstreamHost,omitnil,omitempty" name:"UpstreamHost"`

	// 是否开启缓存。 0：关闭 1：开启
	ProxyBuffer *int64 `json:"ProxyBuffer,omitnil,omitempty" name:"ProxyBuffer"`

	// 是否开启拨测。 0: 禁用拨测 1: 启用拨测。默认启用拨测
	ProbeStatus *int64 `json:"ProbeStatus,omitnil,omitempty" name:"ProbeStatus"`

	// 国密选项。0：不开启国密 1：在原有TLS选项的基础上追加支持国密 2：开启国密并仅支持国密客户端访问
	GmType *int64 `json:"GmType,omitnil,omitempty" name:"GmType"`

	// 国密证书类型。0：无国密证书 1：证书来源为自有国密证书 2：证书来源为托管国密证书
	GmCertType *int64 `json:"GmCertType,omitnil,omitempty" name:"GmCertType"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的证书链
	GmCert *string `json:"GmCert,omitnil,omitempty" name:"GmCert"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的私钥
	GmPrivateKey *string `json:"GmPrivateKey,omitnil,omitempty" name:"GmPrivateKey"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的加密证书
	GmEncCert *string `json:"GmEncCert,omitnil,omitempty" name:"GmEncCert"`

	// GmCertType为1时，需要填充此参数，表示自有国密证书的加密证书的私钥
	GmEncPrivateKey *string `json:"GmEncPrivateKey,omitnil,omitempty" name:"GmEncPrivateKey"`

	// GmCertType为2时，需要填充此参数，表示腾讯云SSL平台托管的证书id
	GmSSLId *string `json:"GmSSLId,omitnil,omitempty" name:"GmSSLId"`
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
	delete(f, "InstanceID")
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
	delete(f, "Note")
	delete(f, "UpstreamHost")
	delete(f, "ProxyBuffer")
	delete(f, "ProbeStatus")
	delete(f, "GmType")
	delete(f, "GmCertType")
	delete(f, "GmCert")
	delete(f, "GmPrivateKey")
	delete(f, "GmEncCert")
	delete(f, "GmEncPrivateKey")
	delete(f, "GmSSLId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySpartaProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySpartaProtectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 防护规则等级 300=standard，400=extended
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`
}

type ModifyUserLevelRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 防护规则等级 300=standard，400=extended
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 主类id
	MainClassID *string `json:"MainClassID,omitnil,omitempty" name:"MainClassID"`

	// 主类开关0=关闭，1=开启，2=只告警
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 下发修改的规则列表
	RuleID []*ReqUserRule `json:"RuleID,omitnil,omitempty" name:"RuleID"`
}

type ModifyUserSignatureRuleRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 主类id
	MainClassID *string `json:"MainClassID,omitnil,omitempty" name:"MainClassID"`

	// 主类开关0=关闭，1=开启，2=只告警
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 下发修改的规则列表
	RuleID []*ReqUserRule `json:"RuleID,omitnil,omitempty" name:"RuleID"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyUserSignatureRuleV2RequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 主类id
	MainClassID *string `json:"MainClassID,omitnil,omitempty" name:"MainClassID"`

	// 主类开关0=关闭，1=开启，2=只告警
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 下发修改的规则列表
	RuleID []*ReqUserRule `json:"RuleID,omitnil,omitempty" name:"RuleID"`
}

type ModifyUserSignatureRuleV2Request struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 主类id
	MainClassID *string `json:"MainClassID,omitnil,omitempty" name:"MainClassID"`

	// 主类开关0=关闭，1=开启，2=只告警
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 下发修改的规则列表
	RuleID []*ReqUserRule `json:"RuleID,omitnil,omitempty" name:"RuleID"`
}

func (r *ModifyUserSignatureRuleV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserSignatureRuleV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "MainClassID")
	delete(f, "Status")
	delete(f, "RuleID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserSignatureRuleV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserSignatureRuleV2ResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserSignatureRuleV2Response struct {
	*tchttp.BaseResponse
	Response *ModifyUserSignatureRuleV2ResponseParams `json:"Response"`
}

func (r *ModifyUserSignatureRuleV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserSignatureRuleV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWafAutoDenyRulesRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 触发IP封禁的攻击次数阈值，范围为2~100次
	AttackThreshold *int64 `json:"AttackThreshold,omitnil,omitempty" name:"AttackThreshold"`

	// IP封禁统计时间，范围为1-60分钟
	TimeThreshold *int64 `json:"TimeThreshold,omitnil,omitempty" name:"TimeThreshold"`

	// 触发IP封禁后的封禁时间，范围为5~360分钟
	DenyTimeThreshold *int64 `json:"DenyTimeThreshold,omitnil,omitempty" name:"DenyTimeThreshold"`

	// 自动封禁状态，0表示关闭，1表示打开
	DefenseStatus *int64 `json:"DefenseStatus,omitnil,omitempty" name:"DefenseStatus"`
}

type ModifyWafAutoDenyRulesRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 触发IP封禁的攻击次数阈值，范围为2~100次
	AttackThreshold *int64 `json:"AttackThreshold,omitnil,omitempty" name:"AttackThreshold"`

	// IP封禁统计时间，范围为1-60分钟
	TimeThreshold *int64 `json:"TimeThreshold,omitnil,omitempty" name:"TimeThreshold"`

	// 触发IP封禁后的封禁时间，范围为5~360分钟
	DenyTimeThreshold *int64 `json:"DenyTimeThreshold,omitnil,omitempty" name:"DenyTimeThreshold"`

	// 自动封禁状态，0表示关闭，1表示打开
	DefenseStatus *int64 `json:"DefenseStatus,omitnil,omitempty" name:"DefenseStatus"`
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
	Success *ResponseCode `json:"Success,omitnil,omitempty" name:"Success"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	WafThreatenIntelligenceDetails *WafThreatenIntelligenceDetails `json:"WafThreatenIntelligenceDetails,omitnil,omitempty" name:"WafThreatenIntelligenceDetails"`
}

type ModifyWafThreatenIntelligenceRequest struct {
	*tchttp.BaseRequest
	
	// 配置WAF威胁情报封禁模块详情
	WafThreatenIntelligenceDetails *WafThreatenIntelligenceDetails `json:"WafThreatenIntelligenceDetails,omitnil,omitempty" name:"WafThreatenIntelligenceDetails"`
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
	WafThreatenIntelligenceDetails *WafThreatenIntelligenceDetails `json:"WafThreatenIntelligenceDetails,omitnil,omitempty" name:"WafThreatenIntelligenceDetails"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Webshell *WebshellStatus `json:"Webshell,omitnil,omitempty" name:"Webshell"`
}

type ModifyWebshellStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名webshell状态
	Webshell *WebshellStatus `json:"Webshell,omitnil,omitempty" name:"Webshell"`
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
	Success *ResponseCode `json:"Success,omitnil,omitempty" name:"Success"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type PeakPointsItem struct {
	// 秒级别时间戳
	Time *uint64 `json:"Time,omitnil,omitempty" name:"Time"`

	// QPS
	Access *uint64 `json:"Access,omitnil,omitempty" name:"Access"`

	// 上行带宽峰值，单位B
	Up *uint64 `json:"Up,omitnil,omitempty" name:"Up"`

	// 下行带宽峰值，单位B
	Down *uint64 `json:"Down,omitnil,omitempty" name:"Down"`

	// Web攻击次数
	Attack *uint64 `json:"Attack,omitnil,omitempty" name:"Attack"`

	// CC攻击次数
	Cc *uint64 `json:"Cc,omitnil,omitempty" name:"Cc"`

	// Bot qps
	BotAccess *uint64 `json:"BotAccess,omitnil,omitempty" name:"BotAccess"`

	// WAF返回给客户端状态码5xx次数
	StatusServerError *uint64 `json:"StatusServerError,omitnil,omitempty" name:"StatusServerError"`

	// WAF返回给客户端状态码4xx次数
	StatusClientError *uint64 `json:"StatusClientError,omitnil,omitempty" name:"StatusClientError"`

	// WAF返回给客户端状态码302次数
	StatusRedirect *uint64 `json:"StatusRedirect,omitnil,omitempty" name:"StatusRedirect"`

	// WAF返回给客户端状态码202次数
	StatusOk *uint64 `json:"StatusOk,omitnil,omitempty" name:"StatusOk"`

	// 源站返回给WAF状态码5xx次数
	UpstreamServerError *uint64 `json:"UpstreamServerError,omitnil,omitempty" name:"UpstreamServerError"`

	// 源站返回给WAF状态码4xx次数
	UpstreamClientError *uint64 `json:"UpstreamClientError,omitnil,omitempty" name:"UpstreamClientError"`

	// 源站返回给WAF状态码302次数
	UpstreamRedirect *uint64 `json:"UpstreamRedirect,omitnil,omitempty" name:"UpstreamRedirect"`

	// 黑名单次数
	BlackIP *uint64 `json:"BlackIP,omitnil,omitempty" name:"BlackIP"`

	// 防篡改次数
	Tamper *uint64 `json:"Tamper,omitnil,omitempty" name:"Tamper"`

	// 信息防泄露次数
	Leak *uint64 `json:"Leak,omitnil,omitempty" name:"Leak"`

	// 访问控制 
	ACL *uint64 `json:"ACL,omitnil,omitempty" name:"ACL"`

	// 小程序 qps
	WxAccess *uint64 `json:"WxAccess,omitnil,omitempty" name:"WxAccess"`

	// 小程序请求数
	WxCount *uint64 `json:"WxCount,omitnil,omitempty" name:"WxCount"`

	// 小程序上行带宽峰值，单位B
	WxUp *uint64 `json:"WxUp,omitnil,omitempty" name:"WxUp"`

	// 小程序下行带宽峰值，单位B
	WxDown *uint64 `json:"WxDown,omitnil,omitempty" name:"WxDown"`
}

type PiechartItem struct {
	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数量
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type PortInfo struct {
	// Nginx的服务器id
	NginxServerId *uint64 `json:"NginxServerId,omitnil,omitempty" name:"NginxServerId"`

	// 监听端口配置
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 与端口对应的协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 回源端口
	UpstreamPort *string `json:"UpstreamPort,omitnil,omitempty" name:"UpstreamPort"`

	// 回源协议
	UpstreamProtocol *string `json:"UpstreamProtocol,omitnil,omitempty" name:"UpstreamProtocol"`
}

type PortItem struct {
	// 监听端口配置
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 与Port一一对应，表示端口对应的协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 与Port一一对应,  表示回源端口
	UpstreamPort *string `json:"UpstreamPort,omitnil,omitempty" name:"UpstreamPort"`

	// 与Port一一对应,  表示回源协议
	UpstreamProtocol *string `json:"UpstreamProtocol,omitnil,omitempty" name:"UpstreamProtocol"`

	// Nginx的服务器ID,新增域名时填"0"
	NginxServerId *string `json:"NginxServerId,omitnil,omitempty" name:"NginxServerId"`
}

// Predefined struct for user
type PostAttackDownloadTaskRequestParams struct {
	// 查询的域名，所有域名使用all
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Lucene语法
	QueryString *string `json:"QueryString,omitnil,omitempty" name:"QueryString"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 默认为desc，可以取值desc和asc
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 下载的日志条数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type PostAttackDownloadTaskRequest struct {
	*tchttp.BaseRequest
	
	// 查询的域名，所有域名使用all
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// Lucene语法
	QueryString *string `json:"QueryString,omitnil,omitempty" name:"QueryString"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 默认为desc，可以取值desc和asc
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 下载的日志条数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
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
	Flow *string `json:"Flow,omitnil,omitempty" name:"Flow"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 版本
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type QPSPackageNew struct {
	// 资源ID
	ResourceIds *string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 过期时间
	ValidTime *string `json:"ValidTime,omitnil,omitempty" name:"ValidTime"`

	// 是否自动续费，1：自动续费，0：不自动续费
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 套餐购买个数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 套餐购买地域，clb-waf暂时没有用到
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 计费项
	BillingItem *string `json:"BillingItem,omitnil,omitempty" name:"BillingItem"`
}

type QpsData struct {
	// 弹性qps默认值
	ElasticBillingDefault *uint64 `json:"ElasticBillingDefault,omitnil,omitempty" name:"ElasticBillingDefault"`

	// 弹性qps最小值
	ElasticBillingMin *uint64 `json:"ElasticBillingMin,omitnil,omitempty" name:"ElasticBillingMin"`

	// 弹性qps最大值
	ElasticBillingMax *uint64 `json:"ElasticBillingMax,omitnil,omitempty" name:"ElasticBillingMax"`

	// 业务扩展包最大qps
	QPSExtendMax *uint64 `json:"QPSExtendMax,omitnil,omitempty" name:"QPSExtendMax"`

	// 境外业务扩展包最大qps
	QPSExtendIntlMax *uint64 `json:"QPSExtendIntlMax,omitnil,omitempty" name:"QPSExtendIntlMax"`
}

// Predefined struct for user
type RefreshAccessCheckResultRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type RefreshAccessCheckResultRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 规则开关
	// 0：关
	// 1：开
	// 2：只告警
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 修改原因
	// 0：无(兼容记录为空)
	// 1：业务自身特性误报避免
	// 2：规则误报上报
	// 3：核心业务规则灰度
	// 4：其它
	Reason *int64 `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type ResponseCode struct {
	// 如果成功则返回Success，失败则返回云api定义的错误码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 如果成功则返回Success，失败则返回WAF定义的二级错误码
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type Rule struct {
	// 规则id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 规则类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 规则等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 规则防护的CVE编号
	CVE *string `json:"CVE,omitnil,omitempty" name:"CVE"`

	// 规则的状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 规则修改的时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 门神规则新增/更新时间
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`
}

type RuleList struct {
	// 规则Id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 规则列表的id
	Rules []*uint64 `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 请求url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 请求的方法
	Function *string `json:"Function,omitnil,omitempty" name:"Function"`

	// 时间戳
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 开关状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type ScanIpInfo struct {
	// 所属业务
	Bussiness *string `json:"Bussiness,omitnil,omitempty" name:"Bussiness"`

	// 扫描对象
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// ip列表
	IpList []*string `json:"IpList,omitnil,omitempty" name:"IpList"`

	// 扫描说明
	Descibe *string `json:"Descibe,omitnil,omitempty" name:"Descibe"`

	// 官方公告
	Referer *string `json:"Referer,omitnil,omitempty" name:"Referer"`

	// 更新时间
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type SearchAccessLogRequestParams struct {
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题，新版本此字段填空字符串
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 查询语句，语句长度最大为4096
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 单次查询返回的日志条数，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 新版本此字段失效，填空字符串，翻页使用Page
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 日志接口是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 第几页，从0开始。新版本接口字段
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type SearchAccessLogRequest struct {
	*tchttp.BaseRequest
	
	// 客户要查询的日志主题ID，每个客户都有对应的一个主题，新版本此字段填空字符串
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 查询语句，语句长度最大为4096
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 单次查询返回的日志条数，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 新版本此字段失效，填空字符串，翻页使用Page
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 日志接口是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 第几页，从0开始。新版本接口字段
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
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
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 日志查询结果是否全部返回，其中，“true”表示结果返回，“false”表示结果为返回
	ListOver *bool `json:"ListOver,omitnil,omitempty" name:"ListOver"`

	// 返回的是否为分析结果，其中，“true”表示返回分析结果，“false”表示未返回分析结果
	Analysis *bool `json:"Analysis,omitnil,omitempty" name:"Analysis"`

	// 如果Analysis为True，则返回分析结果的列名，否则为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ColNames is deprecated.
	ColNames []*string `json:"ColNames,omitnil,omitempty" name:"ColNames"`

	// 日志查询结果；当Analysis为True时，可能返回为null
	// 注意：此字段可能返回 null，表示取不到有效值
	Results []*AccessLogInfo `json:"Results,omitnil,omitempty" name:"Results"`

	// 日志分析结果；当Analysis为False时，可能返回为null
	// 注意：此字段可能返回 null，表示取不到有效值
	//
	// Deprecated: AnalysisResults is deprecated.
	AnalysisResults []*AccessLogItems `json:"AnalysisResults,omitnil,omitempty" name:"AnalysisResults"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 接口升级，这个字段传空字符串,翻页使用Page字段
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// Lucene语法
	QueryString *string `json:"QueryString,omitnil,omitempty" name:"QueryString"`

	// 查询的数量，默认10条，最多100条
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 默认为desc，可以取值desc和asc
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 第几页，从0开始
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
}

type SearchAttackLogRequest struct {
	*tchttp.BaseRequest
	
	// 查询的域名，所有域名使用all
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 接口升级，这个字段传空字符串,翻页使用Page字段
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// Lucene语法
	QueryString *string `json:"QueryString,omitnil,omitempty" name:"QueryString"`

	// 查询的数量，默认10条，最多100条
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 默认为desc，可以取值desc和asc
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 第几页，从0开始
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`
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
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 接口升级，此字段无效，默认返回空字符串
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 攻击日志数组条目内容
	Data []*AttackLogInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// CLS接口返回内容
	ListOver *bool `json:"ListOver,omitnil,omitempty" name:"ListOver"`

	// CLS接口返回内容，标志是否启动新版本索引
	SqlFlag *bool `json:"SqlFlag,omitnil,omitempty" name:"SqlFlag"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClsStatus *string `json:"ClsStatus,omitnil,omitempty" name:"ClsStatus"`

	// waf开关
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 流量模式
	FlowMode *string `json:"FlowMode,omitnil,omitempty" name:"FlowMode"`
}

type SessionData struct {
	// session定义
	Res []*SessionItem `json:"Res,omitnil,omitempty" name:"Res"`
}

type SessionItem struct {
	// 匹配类型
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 起始模式
	KeyOrStartMat *string `json:"KeyOrStartMat,omitnil,omitempty" name:"KeyOrStartMat"`

	// 结束模式
	EndMat *string `json:"EndMat,omitnil,omitempty" name:"EndMat"`

	// 起始偏移
	StartOffset *string `json:"StartOffset,omitnil,omitempty" name:"StartOffset"`

	// 结束偏移
	EndOffset *string `json:"EndOffset,omitnil,omitempty" name:"EndOffset"`

	// 数据源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 更新时间戳
	TsVersion *string `json:"TsVersion,omitnil,omitempty" name:"TsVersion"`

	// SessionID
	SessionId *int64 `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Session名
	SessionName *string `json:"SessionName,omitnil,omitempty" name:"SessionName"`

	// Session是否正在被启用
	SessionInUsed *bool `json:"SessionInUsed,omitnil,omitempty" name:"SessionInUsed"`

	// Session关联的CC规则ID
	RelatedRuleID []*int64 `json:"RelatedRuleID,omitnil,omitempty" name:"RelatedRuleID"`
}

type SpartaProtectionPort struct {
	// 分配的服务器id
	NginxServerId *uint64 `json:"NginxServerId,omitnil,omitempty" name:"NginxServerId"`

	// 端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 后端端口
	UpstreamPort *string `json:"UpstreamPort,omitnil,omitempty" name:"UpstreamPort"`

	// 后端协议
	UpstreamProtocol *string `json:"UpstreamProtocol,omitnil,omitempty" name:"UpstreamProtocol"`
}

type Strategy struct {
	// 匹配字段
	// 
	//     匹配字段不同，相应的匹配参数、逻辑符号、匹配内容有所不同
	// 具体如下所示：
	// <table><thead><tr><th>匹配字段</th><th>匹配参数</th><th>逻辑符号</th><th>匹配内容</th></tr></thead><tbody><tr><td>IP（来源IP）</td><td>不支持参数</td><td>ipmatch（匹配）<br/>ipnmatch（不匹配）</td><td>多个IP以英文逗号隔开,最多20个</td></tr><tr><td>IPV6（来源IPv6）</td><td>不支持参数</td><td>ipmatch（匹配）<br/>ipnmatch（不匹配）</td><td>支持单个IPV6地址</td></tr><tr><td>Referer（Referer）</td><td>不支持参数</td><td>empty（内容为空）<br/>null（不存在）<br/>eq（等于）<br/>neq（不等于）<br/>contains（包含）<br/>ncontains（不包含）<br/>len_eq（长度等于）<br/>len_gt（长度大于）<br/>len_lt（长度小于）<br/>strprefix（前缀匹配）<br/>strsuffix（后缀匹配）<br/>rematch（正则匹配）</td><td>请输入内容,512个字符以内</td></tr><tr><td>URL（请求路径）</td><td>不支持参数</td><td>eq（等于）<br/>neq（不等于）<br/>contains（包含）<br/>ncontains（不包含）<br/>len_eq（长度等于）<br/>len_gt（长度大于）<br/>len_lt（长度小于）<br/>strprefix（前缀匹配）<br/>strsuffix（后缀匹配）<br/>rematch（正则匹配）<br/></td><td>请以/开头,512个字符以内</td></tr><tr><td>UserAgent（UserAgent）</td><td>不支持参数</td><td>同匹配字段<font color="Red">Referer</font>逻辑符号</td><td>请输入内容,512个字符以内</td></tr><tr><td>HTTP_METHOD（HTTP请求方法）</td><td>不支持参数</td><td>eq（等于）<br/>neq（不等于）</td><td>请输入方法名称,建议大写</td></tr><tr><td>QUERY_STRING（请求字符串）</td><td>不支持参数</td><td>同匹配字段<font color="Red">请求路径</font>逻辑符号</td><td>请输入内容,512个字符以内</td></tr><tr><td>GET（GET参数值）</td><td>支持参数录入</td><td>contains（包含）<br/>ncontains（不包含）<br/>len_eq（长度等于）<br/>len_gt（长度大于）<br/>len_lt（长度小于）<br/>strprefix（前缀匹配）<br/>strsuffix（后缀匹配）</td><td>请输入内容,512个字符以内</td></tr><tr><td>GET_PARAMS_NAMES（GET参数名）</td><td>不支持参数</td><td>exsit（存在参数）<br/>nexsit（不存在参数）<br/>len_eq（长度等于）<br/>len_gt（长度大于）<br/>len_lt（长度小于）<br/>strprefix（前缀匹配）<br/>strsuffix（后缀匹配）</td><td>请输入内容,512个字符以内</td></tr><tr><td>POST（POST参数值）</td><td>支持参数录入</td><td>同匹配字段<font color="Red">GET参数值</font>逻辑符号</td><td>请输入内容,512个字符以内</td></tr><tr><td>GET_POST_NAMES（POST参数名）</td><td>不支持参数</td><td>同匹配字段<font color="Red">GET参数名</font>逻辑符号</td><td>请输入内容,512个字符以内</td></tr><tr><td>POST_BODY（完整BODY）</td><td>不支持参数</td><td>同匹配字段<font color="Red">请求路径</font>逻辑符号</td><td>请输入BODY内容,512个字符以内</td></tr><tr><td>COOKIE（Cookie）</td><td>不支持参数</td><td>empty（内容为空）<br/>null（不存在）<br/>rematch（正则匹配）</td><td><font color="Red">暂不支持</font></td></tr><tr><td>GET_COOKIES_NAMES（Cookie参数名）</td><td>不支持参数</td><td>同匹配字段<font color="Red">GET参数名</font>逻辑符号</td><td>请输入内容,512个字符以内</td></tr><tr><td>ARGS_COOKIE（Cookie参数值）</td><td>支持参数录入</td><td>同匹配字段<font color="Red">GET参数值</font>逻辑符号</td><td>请输入内容,512个字符以内</td></tr><tr><td>GET_HEADERS_NAMES（Header参数名）</td><td>不支持参数</td><td>exsit（存在参数）<br/>nexsit（不存在参数）<br/>len_eq（长度等于）<br/>len_gt（长度大于）<br/>len_lt（长度小于）<br/>strprefix（前缀匹配）<br/>strsuffix（后缀匹配）<br/>rematch（正则匹配）</td><td>请输入内容,建议小写,512个字符以内</td></tr><tr><td>ARGS_HEADER（Header参数值）</td><td>支持参数录入</td><td>contains（包含）<br/>ncontains（不包含）<br/>len_eq（长度等于）<br/>len_gt（长度大于）<br/>len_lt（长度小于）<br/>strprefix（前缀匹配）<br/>strsuffix（后缀匹配）<br/>rematch（正则匹配）</td><td>请输入内容,512个字符以内</td></tr></tbody></table>
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

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
	CompareFunc *string `json:"CompareFunc,omitnil,omitempty" name:"CompareFunc"`

	// 匹配内容
	// 
	//     目前 当匹配字段为COOKIE（Cookie）时，不需要输入 匹配内容
	// 其他都需要
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

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
	Arg *string `json:"Arg,omitnil,omitempty" name:"Arg"`

	// 0：大小写敏感
	// 1：大小写不敏感
	CaseNotSensitive *uint64 `json:"CaseNotSensitive,omitnil,omitempty" name:"CaseNotSensitive"`
}

type StrategyForAntiInfoLeak struct {
	// 匹配条件，returncode（响应码）、keywords（关键字）、information（敏感信息）
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 逻辑符号，固定取值为contains
	CompareFunc *string `json:"CompareFunc,omitnil,omitempty" name:"CompareFunc"`

	// 匹配内容。
	// 以下三个对应Field为information时可取的匹配内容：
	// idcard（身份证）、phone（手机号）、bankcard（银行卡）。
	// 以下为对应Field为returncode时可取的匹配内容：
	// 400（状态码400）、403（状态码403）、404（状态码404）、4xx（其它4xx状态码）、500（状态码500）、501（状态码501）、502（状态码502）、504（状态码504）、5xx（其它5xx状态码）。
	// 当对应Field为keywords时由用户自己输入匹配内容。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

// Predefined struct for user
type SwitchDomainRulesRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则列表
	Ids []*uint64 `json:"Ids,omitnil,omitempty" name:"Ids"`

	// 开关状态，0表示关闭，1表示开启，2表示只观察
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 设置为观察模式原因，
	// 1表示业务自身原因观察，2表示系统规则误报上报，3表示核心业务灰度观察，4表示其他
	Reason *uint64 `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type SwitchDomainRulesRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则列表
	Ids []*uint64 `json:"Ids,omitnil,omitempty" name:"Ids"`

	// 开关状态，0表示关闭，1表示开启，2表示只观察
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 设置为观察模式原因，
	// 1表示业务自身原因观察，2表示系统规则误报上报，3表示核心业务灰度观察，4表示其他
	Reason *uint64 `json:"Reason,omitnil,omitempty" name:"Reason"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 0代表关闭，1代表打开
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type SwitchElasticModeRequest struct {
	*tchttp.BaseRequest
	
	// 版本，只能是sparta-waf, clb-waf, cdn-waf
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 0代表关闭，1代表打开
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 实例id
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	VersionId *int64 `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// 加密套件ID
	CipherId *int64 `json:"CipherId,omitnil,omitempty" name:"CipherId"`

	// 加密套件
	CipherName *string `json:"CipherName,omitnil,omitempty" name:"CipherName"`
}

type TLSVersion struct {
	// TLSVERSION的ID
	VersionId *int64 `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// TLSVERSION的NAME
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`
}

type TargetEntity struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type TigaMainClassMode struct {
	// MainclassID
	TypeID *string `json:"TypeID,omitnil,omitempty" name:"TypeID"`

	// 防护模式，0表示观察，1表示拦截
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`
}

type TimedJob struct {
	// 开始时间戳，单位为秒
	StartDateTime *uint64 `json:"StartDateTime,omitnil,omitempty" name:"StartDateTime"`

	// 结束时间戳，单位为秒
	EndDateTime *uint64 `json:"EndDateTime,omitnil,omitempty" name:"EndDateTime"`
}

// Predefined struct for user
type UpsertCCAutoStatusRequestParams struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 状态值
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`

	// 版本：clb-waf, spart-waf
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`
}

type UpsertCCAutoStatusRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 状态值
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`

	// 版本：clb-waf, spart-waf
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`
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
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 高级模式（是否使用Session检测），0表示不启用，1表示启用
	Advance *string `json:"Advance,omitnil,omitempty" name:"Advance"`

	// CC检测阈值
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// CC检测周期
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 检测Url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 匹配方法，0表示等于，1表示前缀匹配，2表示包含
	MatchFunc *int64 `json:"MatchFunc,omitnil,omitempty" name:"MatchFunc"`

	// 动作，20表示观察，21表示人机识别，22表示拦截，23表示精准拦截，26表示精准人机识别，27表示JS校验
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 优先级
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 动作有效时间
	ValidTime *int64 `json:"ValidTime,omitnil,omitempty" name:"ValidTime"`

	// 附加参数
	OptionsArr *string `json:"OptionsArr,omitnil,omitempty" name:"OptionsArr"`

	// waf版本，sparta-waf或者clb-waf
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 操作类型
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 添加规则的来源事件id
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 规则需要启用的SessionID
	SessionApplied []*int64 `json:"SessionApplied,omitnil,omitempty" name:"SessionApplied"`

	// 规则ID，新增时填0
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则创建时间
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// url长度
	Length *uint64 `json:"Length,omitnil,omitempty" name:"Length"`
}

type UpsertCCRuleRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 高级模式（是否使用Session检测），0表示不启用，1表示启用
	Advance *string `json:"Advance,omitnil,omitempty" name:"Advance"`

	// CC检测阈值
	Limit *string `json:"Limit,omitnil,omitempty" name:"Limit"`

	// CC检测周期
	Interval *string `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 检测Url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 匹配方法，0表示等于，1表示前缀匹配，2表示包含
	MatchFunc *int64 `json:"MatchFunc,omitnil,omitempty" name:"MatchFunc"`

	// 动作，20表示观察，21表示人机识别，22表示拦截，23表示精准拦截，26表示精准人机识别，27表示JS校验
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 优先级
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 动作有效时间
	ValidTime *int64 `json:"ValidTime,omitnil,omitempty" name:"ValidTime"`

	// 附加参数
	OptionsArr *string `json:"OptionsArr,omitnil,omitempty" name:"OptionsArr"`

	// waf版本，sparta-waf或者clb-waf
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 操作类型
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 添加规则的来源事件id
	EventId *string `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 规则需要启用的SessionID
	SessionApplied []*int64 `json:"SessionApplied,omitnil,omitempty" name:"SessionApplied"`

	// 规则ID，新增时填0
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则创建时间
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// url长度
	Length *uint64 `json:"Length,omitnil,omitempty" name:"Length"`
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
	delete(f, "CreateTime")
	delete(f, "Length")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpsertCCRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpsertCCRuleResponseParams struct {
	// 一般为null
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 操作的RuleId
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 具体域名如：test.qcloudwaf.com
	// 全局域名为：global
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// IP 参数列表，json数组由IP，source，note，action，valid_ts组成。IP对应配置的IP地址，source固定为custom值，note为注释，action值42为黑名单，40为白名单，valid_ts为有效日期，值为秒级时间戳（（如1680570420代表2023-04-04 09:07:00））
	Items []*string `json:"Items,omitnil,omitempty" name:"Items"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 可选值为：batch（批量添加）、bot、cc、custom（非批量添加时的默认值）
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`
}

type UpsertIpAccessControlRequest struct {
	*tchttp.BaseRequest
	
	// 具体域名如：test.qcloudwaf.com
	// 全局域名为：global
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// IP 参数列表，json数组由IP，source，note，action，valid_ts组成。IP对应配置的IP地址，source固定为custom值，note为注释，action值42为黑名单，40为白名单，valid_ts为有效日期，值为秒级时间戳（（如1680570420代表2023-04-04 09:07:00））
	Items []*string `json:"Items,omitnil,omitempty" name:"Items"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// WAF实例类型，sparta-waf表示SAAS型WAF，clb-waf表示负载均衡型WAF
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 可选值为：batch（批量添加）、bot、cc、custom（非批量添加时的默认值）
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`
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
	FailedItems *string `json:"FailedItems,omitnil,omitempty" name:"FailedItems"`

	// 添加或修改失败的数目
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// 添加或修改的IP数据Id列表
	Ids []*string `json:"Ids,omitnil,omitempty" name:"Ids"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// session来源位置
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 提取类别
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 提取key或者起始匹配模式
	KeyOrStartMat *string `json:"KeyOrStartMat,omitnil,omitempty" name:"KeyOrStartMat"`

	// 结束匹配模式
	EndMat *string `json:"EndMat,omitnil,omitempty" name:"EndMat"`

	// 起始偏移位置
	StartOffset *string `json:"StartOffset,omitnil,omitempty" name:"StartOffset"`

	// 结束偏移位置
	EndOffset *string `json:"EndOffset,omitnil,omitempty" name:"EndOffset"`

	// 版本
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// Session名
	SessionName *string `json:"SessionName,omitnil,omitempty" name:"SessionName"`

	// Session对应ID
	SessionID *int64 `json:"SessionID,omitnil,omitempty" name:"SessionID"`
}

type UpsertSessionRequest struct {
	*tchttp.BaseRequest
	
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// session来源位置
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 提取类别
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 提取key或者起始匹配模式
	KeyOrStartMat *string `json:"KeyOrStartMat,omitnil,omitempty" name:"KeyOrStartMat"`

	// 结束匹配模式
	EndMat *string `json:"EndMat,omitnil,omitempty" name:"EndMat"`

	// 起始偏移位置
	StartOffset *string `json:"StartOffset,omitnil,omitempty" name:"StartOffset"`

	// 结束偏移位置
	EndOffset *string `json:"EndOffset,omitnil,omitempty" name:"EndOffset"`

	// 版本
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// Session名
	SessionName *string `json:"SessionName,omitnil,omitempty" name:"SessionName"`

	// Session对应ID
	SessionID *int64 `json:"SessionID,omitnil,omitempty" name:"SessionID"`
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
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// SessionID
	SessionID *int64 `json:"SessionID,omitnil,omitempty" name:"SessionID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Appid *uint64 `json:"Appid,omitnil,omitempty" name:"Appid"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名id
	DomainId *string `json:"DomainId,omitnil,omitempty" name:"DomainId"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// waf类型
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 版本
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 指定域名访问日志字段的开关
	WriteConfig *string `json:"WriteConfig,omitnil,omitempty" name:"WriteConfig"`

	// 指定域名是否写cls的开关 1:写 0:不写
	Cls *uint64 `json:"Cls,omitnil,omitempty" name:"Cls"`

	// 标记是否是混合云接入。hybrid表示混合云接入域名
	CloudType *string `json:"CloudType,omitnil,omitempty" name:"CloudType"`

	// 标记clbwaf类型
	AlbType *string `json:"AlbType,omitnil,omitempty" name:"AlbType"`

	// BOT开关状态
	BotStatus *int64 `json:"BotStatus,omitnil,omitempty" name:"BotStatus"`

	// API开关状态
	ApiStatus *int64 `json:"ApiStatus,omitnil,omitempty" name:"ApiStatus"`
}

type UserSignatureRule struct {
	// 特征ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 规则开关
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 主类ID
	MainClassID *string `json:"MainClassID,omitnil,omitempty" name:"MainClassID"`

	// 子类ID
	SubClassID *string `json:"SubClassID,omitnil,omitempty" name:"SubClassID"`

	// CVE ID
	CveID *string `json:"CveID,omitnil,omitempty" name:"CveID"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 主类名字，根据Language字段输出中文/英文
	MainClassName *string `json:"MainClassName,omitnil,omitempty" name:"MainClassName"`

	// 子类名字，根据Language字段输出中文/英文，若子类id为00000000，此字段为空
	SubClassName *string `json:"SubClassName,omitnil,omitempty" name:"SubClassName"`

	// 规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 0/1
	Reason *int64 `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 1: 高危 2:中危 3:低危
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`
}

type UserWhiteRule struct {
	// 白名单的id
	WhiteRuleId *uint64 `json:"WhiteRuleId,omitnil,omitempty" name:"WhiteRuleId"`

	// 规则id
	SignatureId *string `json:"SignatureId,omitnil,omitempty" name:"SignatureId"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 匹配域
	MatchField *string `json:"MatchField,omitnil,omitempty" name:"MatchField"`

	// 匹配参数
	MatchParams *string `json:"MatchParams,omitnil,omitempty" name:"MatchParams"`

	// 匹配方法
	MatchMethod *string `json:"MatchMethod,omitnil,omitempty" name:"MatchMethod"`

	// 匹配内容
	MatchContent *string `json:"MatchContent,omitnil,omitempty" name:"MatchContent"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 规则ID列表
	SignatureIds []*string `json:"SignatureIds,omitnil,omitempty" name:"SignatureIds"`

	// 大类规则ID列表
	TypeIds []*string `json:"TypeIds,omitnil,omitempty" name:"TypeIds"`

	// 大类规则ID
	TypeId *string `json:"TypeId,omitnil,omitempty" name:"TypeId"`

	// 0:按照特定规则ID加白, 1:按照规则类型加白
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 规则名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 匹配规则列表
	MatchInfo []*UserWhiteRuleItem `json:"MatchInfo,omitnil,omitempty" name:"MatchInfo"`

	// MatchInfo字符串
	MatchInfoStr *string `json:"MatchInfoStr,omitnil,omitempty" name:"MatchInfoStr"`
}

type UserWhiteRuleItem struct {
	// 匹配域
	MatchField *string `json:"MatchField,omitnil,omitempty" name:"MatchField"`

	// 匹配方法
	MatchMethod *string `json:"MatchMethod,omitnil,omitempty" name:"MatchMethod"`

	// 匹配内容
	MatchContent *string `json:"MatchContent,omitnil,omitempty" name:"MatchContent"`

	// 匹配参数名
	MatchParams *string `json:"MatchParams,omitnil,omitempty" name:"MatchParams"`
}

type VipInfo struct {
	// VIP地址
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// waf实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 创建时间
	InstanceCreateTime *string `json:"InstanceCreateTime,omitnil,omitempty" name:"InstanceCreateTime"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 地域ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// ip运营商类型
	ISP *string `json:"ISP,omitnil,omitempty" name:"ISP"`

	// ip类型
	VipType *string `json:"VipType,omitnil,omitempty" name:"VipType"`

	// 域名信息
	AddressName *string `json:"AddressName,omitnil,omitempty" name:"AddressName"`
}

type WafRuleLimit struct {
	// 自定义CC的规格
	CC *uint64 `json:"CC,omitnil,omitempty" name:"CC"`

	// 自定义策略的规格
	CustomRule *uint64 `json:"CustomRule,omitnil,omitempty" name:"CustomRule"`

	// 黑白名单的规格
	IPControl *uint64 `json:"IPControl,omitnil,omitempty" name:"IPControl"`

	// 信息防泄漏的规格
	AntiLeak *uint64 `json:"AntiLeak,omitnil,omitempty" name:"AntiLeak"`

	// 防篡改的规格
	AntiTamper *uint64 `json:"AntiTamper,omitnil,omitempty" name:"AntiTamper"`

	// 紧急CC的规格
	AutoCC *uint64 `json:"AutoCC,omitnil,omitempty" name:"AutoCC"`

	// 地域封禁的规格
	AreaBan *uint64 `json:"AreaBan,omitnil,omitempty" name:"AreaBan"`

	// 自定义CC中配置session
	CCSession *uint64 `json:"CCSession,omitnil,omitempty" name:"CCSession"`

	// AI的规格
	AI *uint64 `json:"AI,omitnil,omitempty" name:"AI"`

	// 精准白名单的规格
	CustomWhite *uint64 `json:"CustomWhite,omitnil,omitempty" name:"CustomWhite"`

	// api安全的规格
	ApiSecurity *uint64 `json:"ApiSecurity,omitnil,omitempty" name:"ApiSecurity"`

	// 客户端流量标记的规格
	ClientMsg *uint64 `json:"ClientMsg,omitnil,omitempty" name:"ClientMsg"`

	// 流量标记的规格
	TrafficMarking *uint64 `json:"TrafficMarking,omitnil,omitempty" name:"TrafficMarking"`
}

type WafThreatenIntelligenceDetails struct {
	// 封禁属性标签
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 封禁模组启用状态
	DefenseStatus *int64 `json:"DefenseStatus,omitnil,omitempty" name:"DefenseStatus"`

	// 最后更新时间
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`
}

type WebshellStatus struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// webshell开关，1：开。0：关。2：观察
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}