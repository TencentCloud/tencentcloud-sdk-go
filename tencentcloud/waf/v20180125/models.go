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
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AddCustomRuleRequest struct {
	*tchttp.BaseRequest

	// 规则名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 优先级
	SortId *string `json:"SortId,omitempty" name:"SortId"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 策略详情
	Strategies []*Strategy `json:"Strategies,omitempty" name:"Strategies" list`

	// 需要添加策略的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 动作类型
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 如果动作是重定向，则表示重定向的地址；其他情况可以为空
	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`

	// "clb-waf"或者"sparta-waf"
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 放行的详情
	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`
}

func (r *AddCustomRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
	if len(f) > 0 {
		return errors.New("AddCustomRuleRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AddCustomRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
		Success *ResponseCode `json:"Success,omitempty" name:"Success"`

		// 添加成功的规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddCustomRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCustomRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAttackDownloadTaskRequest struct {
	*tchttp.BaseRequest

	// 域名，所有域名填写all
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 查询起始时间
	FromTime *string `json:"FromTime,omitempty" name:"FromTime"`

	// 查询结束时间
	ToTime *string `json:"ToTime,omitempty" name:"ToTime"`

	// 下载任务名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 风险等级
	RiskLevel *uint64 `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 拦截状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 自定义策略ID
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// 攻击者IP
	AttackIp *string `json:"AttackIp,omitempty" name:"AttackIp"`

	// 攻击类型
	AttackType *string `json:"AttackType,omitempty" name:"AttackType"`
}

func (r *CreateAttackDownloadTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAttackDownloadTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "FromTime")
	delete(f, "ToTime")
	delete(f, "Name")
	delete(f, "RiskLevel")
	delete(f, "Status")
	delete(f, "RuleId")
	delete(f, "AttackIp")
	delete(f, "AttackType")
	if len(f) > 0 {
		return errors.New("CreateAttackDownloadTaskRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAttackDownloadTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		Flow *string `json:"Flow,omitempty" name:"Flow"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAttackDownloadTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAttackDownloadTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAttackDownloadRecordRequest struct {
	*tchttp.BaseRequest

	// 下载任务记录唯一标记
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteAttackDownloadRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttackDownloadRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return errors.New("DeleteAttackDownloadRecordRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAttackDownloadRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAttackDownloadRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttackDownloadRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDownloadRecordRequest struct {
	*tchttp.BaseRequest

	// 记录id
	Flow *string `json:"Flow,omitempty" name:"Flow"`
}

func (r *DeleteDownloadRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDownloadRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Flow")
	if len(f) > 0 {
		return errors.New("DeleteDownloadRecordRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDownloadRecordResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDownloadRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDownloadRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSessionRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// clb-waf 或者 sprta-waf
	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *DeleteSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Edition")
	if len(f) > 0 {
		return errors.New("DeleteSessionRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomRulesPagingInfo struct {

	// 当前页码
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 当前页的最大数据条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeCustomRulesRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 分页参数
	Paging *DescribeCustomRulesPagingInfo `json:"Paging,omitempty" name:"Paging"`

	// clb-waf或者sparta-waf
	Edition *string `json:"Edition,omitempty" name:"Edition"`

	// 过滤参数：动作类型：0放行，1阻断，2人机识别，3观察，4重定向
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 过滤参数：规则名称过滤条件
	Search *string `json:"Search,omitempty" name:"Search"`
}

func (r *DescribeCustomRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Domain")
	delete(f, "Paging")
	delete(f, "Edition")
	delete(f, "ActionType")
	delete(f, "Search")
	if len(f) > 0 {
		return errors.New("DescribeCustomRulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则详情
		RuleList []*DescribeCustomRulesRspRuleListItem `json:"RuleList,omitempty" name:"RuleList" list`

		// 规则条数
		TotalCount *string `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomRulesRspRuleListItem struct {

	// 动作类型
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 跳过的策略
	Bypass *string `json:"Bypass,omitempty" name:"Bypass"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 重定向地址
	Redirect *string `json:"Redirect,omitempty" name:"Redirect"`

	// 策略ID
	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`

	// 优先级
	SortId *string `json:"SortId,omitempty" name:"SortId"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 策略详情
	Strategies []*Strategy `json:"Strategies,omitempty" name:"Strategies" list`
}

type DescribeUserClbWafRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserClbWafRegionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserClbWafRegionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("DescribeUserClbWafRegionsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserClbWafRegionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 地域（标准的ap-格式）列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*string `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserClbWafRegionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserClbWafRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitempty" name:"RuleId"`

	// 开关的状态，1是开启、0是关闭
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// WAF的版本，clb-waf代表负载均衡WAF、sparta-waf代表SaaS WAF，默认是sparta-waf。
	Edition *string `json:"Edition,omitempty" name:"Edition"`
}

func (r *ModifyCustomRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("ModifyCustomRuleStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCustomRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 操作的状态码，如果所有的资源操作成功则返回的是成功的状态码，如果有资源操作失败则需要解析Message的内容来查看哪个资源失败
		Success *ResponseCode `json:"Success,omitempty" name:"Success"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCustomRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResponseCode struct {

	// 如果成功则返回Success，失败则返回yunapi定义的错误码
	Code *string `json:"Code,omitempty" name:"Code"`

	// 如果成功则返回Success，失败则返回WAF定义的二级错误码
	Message *string `json:"Message,omitempty" name:"Message"`
}

type Strategy struct {

	// 匹配字段
	Field *string `json:"Field,omitempty" name:"Field"`

	// 逻辑符号
	CompareFunc *string `json:"CompareFunc,omitempty" name:"CompareFunc"`

	// 匹配内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 匹配参数
	Arg *string `json:"Arg,omitempty" name:"Arg"`
}
