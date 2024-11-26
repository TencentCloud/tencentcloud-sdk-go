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

package v20210811

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AffectedComponent struct {
	// 受漏洞影响的组件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 受漏洞影响的版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	AffectedVersionList []*string `json:"AffectedVersionList,omitnil,omitempty" name:"AffectedVersionList"`

	// 修复此漏洞的版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	FixedVersionList []*string `json:"FixedVersionList,omitnil,omitempty" name:"FixedVersionList"`
}

type CVSSV2Info struct {
	// CVE评分。
	CVSS *float64 `json:"CVSS,omitnil,omitempty" name:"CVSS"`

	// AccessVector 攻击途径。
	// 取值范围：
	// <li>NETWORK 远程</li>
	// <li>ADJACENT_NETWORK 近邻</li>
	// <li>LOCAL 本地</li>
	AccessVector *string `json:"AccessVector,omitnil,omitempty" name:"AccessVector"`

	// AccessComplexity 攻击复杂度。
	// 取值范围：
	// <li>HIGH 高</li>
	// <li>MEDIUM 中</li>
	// <li>LOW 低</li>
	AccessComplexity *string `json:"AccessComplexity,omitnil,omitempty" name:"AccessComplexity"`

	// Authentication 身份验证。
	// 取值范围：
	// <li>MULTIPLE 多系统认证</li>
	// <li>SINGLE 单系统认证</li>
	// <li>NONE 无</li>
	Authentication *string `json:"Authentication,omitnil,omitempty" name:"Authentication"`

	// ConfidentialityImpact 机密性影响。
	// 取值范围：
	// <li>NONE 无</li>
	// <li>PARTIAL 部分</li>
	// <li>COMPLETE 完整</li>
	ConImpact *string `json:"ConImpact,omitnil,omitempty" name:"ConImpact"`

	// IntegrityImpact 完整性影响。
	// 取值范围：
	// <li>NONE 无</li>
	// <li>PARTIAL 部分</li>
	// <li>COMPLETE 完整</li>
	IntegrityImpact *string `json:"IntegrityImpact,omitnil,omitempty" name:"IntegrityImpact"`

	// AvailabilityImpact 可用性影响。
	// 取值范围：
	// <li>NONE 无</li>
	// <li>PARTIAL 部分</li>
	// <li>COMPLETE 完整</li>
	AvailabilityImpact *string `json:"AvailabilityImpact,omitnil,omitempty" name:"AvailabilityImpact"`
}

type CVSSV3Info struct {
	// CVE评分。
	CVSS *float64 `json:"CVSS,omitnil,omitempty" name:"CVSS"`

	// AttackVector 攻击途径。
	// 取值范围：
	// <li>NETWORK 远程</li>
	// <li>ADJACENT_NETWORK 近邻</li>
	// <li>LOCAL 本地</li>
	// <li>PHYSICAL 物理</li>
	AttackVector *string `json:"AttackVector,omitnil,omitempty" name:"AttackVector"`

	// AttackComplexity 攻击复杂度。
	// 取值范围：
	// <li>HIGH 高</li>
	// <li>LOW 低</li>
	AttackComplexity *string `json:"AttackComplexity,omitnil,omitempty" name:"AttackComplexity"`

	// PrivilegesRequired 触发特权。
	// 取值范围：
	// <li>HIGH 高</li>
	// <li>LOW 低</li>
	// <li>NONE 无</li>
	PrivilegesRequired *string `json:"PrivilegesRequired,omitnil,omitempty" name:"PrivilegesRequired"`

	// UserInteraction 交互必要性。
	// 取值范围：
	// <li>NONE 无</li>
	// <li>REQUIRED 需要</li>
	UserInteraction *string `json:"UserInteraction,omitnil,omitempty" name:"UserInteraction"`

	// Scope 绕过安全边界。
	// 取值范围：
	// <li>UNCHANGED 否</li>
	// <li>CHANGED 能</li>
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// ConfidentialityImpact 机密性影响。
	// 取值范围：
	// <li>NONE 无</li>
	// <li>LOW 低</li>
	// <li>HIGH 高</li>
	ConImpact *string `json:"ConImpact,omitnil,omitempty" name:"ConImpact"`

	// IntegrityImpact 完整性影响。
	// 取值范围：
	// <li>NONE 无</li>
	// <li>LOW 低</li>
	// <li>HIGH 高</li>
	IntegrityImpact *string `json:"IntegrityImpact,omitnil,omitempty" name:"IntegrityImpact"`

	// AvailabilityImpact 可用性影响。
	// 取值范围：
	// <li>NONE 无</li>
	// <li>LOW 低</li>
	// <li>HIGH 高</li>
	AvailabilityImpact *string `json:"AvailabilityImpact,omitnil,omitempty" name:"AvailabilityImpact"`
}

type Component struct {
	// 第三方组件的PURL
	PURL *PURL `json:"PURL,omitnil,omitempty" name:"PURL"`

	// 第三方组件的主页
	Homepage *string `json:"Homepage,omitnil,omitempty" name:"Homepage"`

	// 第三方组件的简介
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 第三方组件的别名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NicknameList []*string `json:"NicknameList,omitnil,omitempty" name:"NicknameList"`

	// 第三方组件的代码位置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeLocationList []*string `json:"CodeLocationList,omitnil,omitempty" name:"CodeLocationList"`

	// 第三方组件的许可证表达式
	LicenseExpression *string `json:"LicenseExpression,omitnil,omitempty" name:"LicenseExpression"`

	// 第三方组件的版本信息(如果匹配到版本)
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionInfo *ComponentVersionInfo `json:"VersionInfo,omitnil,omitempty" name:"VersionInfo"`

	// 第三方组件的最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// 第三方组件的类型标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*string `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type ComponentTagFilter struct {
	// 包括的Tag
	IncludeTags []*string `json:"IncludeTags,omitnil,omitempty" name:"IncludeTags"`

	// 排除的Tag
	ExcludeTags []*string `json:"ExcludeTags,omitnil,omitempty" name:"ExcludeTags"`
}

type ComponentVersion struct {
	// 该组件的PURL
	// 注意：此字段可能返回 null，表示取不到有效值。
	PURL *PURL `json:"PURL,omitnil,omitempty" name:"PURL"`

	// 该组件版本的许可证表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	LicenseExpression *string `json:"LicenseExpression,omitnil,omitempty" name:"LicenseExpression"`

	// 组件的版本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionInfo *ComponentVersionInfo `json:"VersionInfo,omitnil,omitempty" name:"VersionInfo"`
}

type ComponentVersionInfo struct {
	// 版本发布时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublishTime *string `json:"PublishTime,omitnil,omitempty" name:"PublishTime"`

	// 当前版本的所有copyright
	// 注意：此字段可能返回 null，表示取不到有效值。
	CopyrightList []*string `json:"CopyrightList,omitnil,omitempty" name:"CopyrightList"`

	// 版本标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*string `json:"TagList,omitnil,omitempty" name:"TagList"`
}

type ComponentVulnerabilitySummary struct {
	// 用于匹配漏洞的PURL
	// 注意：此字段可能返回 null，表示取不到有效值。
	PURL *PURL `json:"PURL,omitnil,omitempty" name:"PURL"`

	// 该组件是否包含修复漏洞的官方补丁
	CanBeFixed *bool `json:"CanBeFixed,omitnil,omitempty" name:"CanBeFixed"`

	// 修复漏洞的组件版本号
	FixedVersion *string `json:"FixedVersion,omitnil,omitempty" name:"FixedVersion"`

	// 漏洞影响的组件版本号
	AffectedVersion *string `json:"AffectedVersion,omitnil,omitempty" name:"AffectedVersion"`

	// 漏洞影响组件
	AffectedComponent *string `json:"AffectedComponent,omitnil,omitempty" name:"AffectedComponent"`

	// 漏洞在该产品中的风险等级
	// <li>Critical</li>
	// <li>High</li>
	// <li>Medium</li>
	// <li>Low</li>
	RiskLevel *string `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`
}

type ComponentVulnerabilityUnion struct {
	// 漏洞概览信息
	Summary *VulnerabilitySummary `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 与组件相关的漏洞概览信息
	SummaryInComponent *ComponentVulnerabilitySummary `json:"SummaryInComponent,omitnil,omitempty" name:"SummaryInComponent"`
}

// Predefined struct for user
type DescribeKBComponentRequestParams struct {
	// 组件的PURL
	PURL *PURL `json:"PURL,omitnil,omitempty" name:"PURL"`
}

type DescribeKBComponentRequest struct {
	*tchttp.BaseRequest
	
	// 组件的PURL
	PURL *PURL `json:"PURL,omitnil,omitempty" name:"PURL"`
}

func (r *DescribeKBComponentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKBComponentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PURL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKBComponentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKBComponentResponseParams struct {
	// 匹配的组件信息
	Component *Component `json:"Component,omitnil,omitempty" name:"Component"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKBComponentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKBComponentResponseParams `json:"Response"`
}

func (r *DescribeKBComponentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKBComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKBComponentVersionListRequestParams struct {
	// 要查询的组件 PURL
	PURL *PURL `json:"PURL,omitnil,omitempty" name:"PURL"`

	// 页号
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 排序方式，可以是"ASC"或"DESC"，默认"DESC"
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段，可能的字段包括“Version”、"PublishTime"
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Tag筛选
	Filter *ComponentTagFilter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeKBComponentVersionListRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的组件 PURL
	PURL *PURL `json:"PURL,omitnil,omitempty" name:"PURL"`

	// 页号
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 排序方式，可以是"ASC"或"DESC"，默认"DESC"
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段，可能的字段包括“Version”、"PublishTime"
	OrderBy []*string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// Tag筛选
	Filter *ComponentTagFilter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeKBComponentVersionListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKBComponentVersionListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PURL")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Order")
	delete(f, "OrderBy")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKBComponentVersionListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKBComponentVersionListResponseParams struct {
	// 该组件的版本列表信息
	VersionList []*ComponentVersion `json:"VersionList,omitnil,omitempty" name:"VersionList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKBComponentVersionListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKBComponentVersionListResponseParams `json:"Response"`
}

func (r *DescribeKBComponentVersionListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKBComponentVersionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKBComponentVulnerabilityRequestParams struct {
	// 组件的PURL，其中Name和Version为必填字段
	PURL *PURL `json:"PURL,omitnil,omitempty" name:"PURL"`

	// 语言，ZH或EN
	Language *string `json:"Language,omitnil,omitempty" name:"Language"`
}

type DescribeKBComponentVulnerabilityRequest struct {
	*tchttp.BaseRequest
	
	// 组件的PURL，其中Name和Version为必填字段
	PURL *PURL `json:"PURL,omitnil,omitempty" name:"PURL"`

	// 语言，ZH或EN
	Language *string `json:"Language,omitnil,omitempty" name:"Language"`
}

func (r *DescribeKBComponentVulnerabilityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKBComponentVulnerabilityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PURL")
	delete(f, "Language")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKBComponentVulnerabilityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKBComponentVulnerabilityResponseParams struct {
	// 漏洞信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityList []*ComponentVulnerabilityUnion `json:"VulnerabilityList,omitnil,omitempty" name:"VulnerabilityList"`

	// 组件purl
	PURL *PURL `json:"PURL,omitnil,omitempty" name:"PURL"`

	// 推荐版本，最小无高危/严重漏洞的版本。无法升级到安全版本时的备选方案。
	RecommendedVersion *string `json:"RecommendedVersion,omitnil,omitempty" name:"RecommendedVersion"`

	// 安全版本（首选），最小无漏洞的版本。当无法升级到安全版本时可考虑使用推荐版本。
	SecureVersion *string `json:"SecureVersion,omitnil,omitempty" name:"SecureVersion"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKBComponentVulnerabilityResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKBComponentVulnerabilityResponseParams `json:"Response"`
}

func (r *DescribeKBComponentVulnerabilityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKBComponentVulnerabilityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKBLicenseRequestParams struct {
	// License表达式
	LicenseExpression *string `json:"LicenseExpression,omitnil,omitempty" name:"LicenseExpression"`
}

type DescribeKBLicenseRequest struct {
	*tchttp.BaseRequest
	
	// License表达式
	LicenseExpression *string `json:"LicenseExpression,omitnil,omitempty" name:"LicenseExpression"`
}

func (r *DescribeKBLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKBLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LicenseExpression")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKBLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKBLicenseResponseParams struct {
	// 许可证列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LicenseList []*LicenseUnion `json:"LicenseList,omitnil,omitempty" name:"LicenseList"`

	// 用于匹配的License表达式
	NormalizedLicenseExpression *string `json:"NormalizedLicenseExpression,omitnil,omitempty" name:"NormalizedLicenseExpression"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKBLicenseResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKBLicenseResponseParams `json:"Response"`
}

func (r *DescribeKBLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKBLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKBVulnerabilityRequestParams struct {
	// 根据CVE ID查询（不能与其他参数同时存在）
	CVEID []*string `json:"CVEID,omitnil,omitempty" name:"CVEID"`

	// 根据Vul ID查询（不能与其他参数同时存在）
	VulID []*string `json:"VulID,omitnil,omitempty" name:"VulID"`

	// 根据CNVD ID查询（不能与其他参数同时存在）
	CNVDID []*string `json:"CNVDID,omitnil,omitempty" name:"CNVDID"`

	// 根据CNNVD ID查询（不能与其他参数同时存在）
	CNNVDID []*string `json:"CNNVDID,omitnil,omitempty" name:"CNNVDID"`

	// 语言，ZH或EN
	Language *string `json:"Language,omitnil,omitempty" name:"Language"`
}

type DescribeKBVulnerabilityRequest struct {
	*tchttp.BaseRequest
	
	// 根据CVE ID查询（不能与其他参数同时存在）
	CVEID []*string `json:"CVEID,omitnil,omitempty" name:"CVEID"`

	// 根据Vul ID查询（不能与其他参数同时存在）
	VulID []*string `json:"VulID,omitnil,omitempty" name:"VulID"`

	// 根据CNVD ID查询（不能与其他参数同时存在）
	CNVDID []*string `json:"CNVDID,omitnil,omitempty" name:"CNVDID"`

	// 根据CNNVD ID查询（不能与其他参数同时存在）
	CNNVDID []*string `json:"CNNVDID,omitnil,omitempty" name:"CNNVDID"`

	// 语言，ZH或EN
	Language *string `json:"Language,omitnil,omitempty" name:"Language"`
}

func (r *DescribeKBVulnerabilityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKBVulnerabilityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CVEID")
	delete(f, "VulID")
	delete(f, "CNVDID")
	delete(f, "CNNVDID")
	delete(f, "Language")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKBVulnerabilityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKBVulnerabilityResponseParams struct {
	// 漏洞详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityDetailList []*VulnerabilityUnion `json:"VulnerabilityDetailList,omitnil,omitempty" name:"VulnerabilityDetailList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKBVulnerabilityResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKBVulnerabilityResponseParams `json:"Response"`
}

func (r *DescribeKBVulnerabilityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKBVulnerabilityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LicenseDetail struct {
	// 许可证内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 许可证允许信息列表
	ConditionSet []*LicenseRestriction `json:"ConditionSet,omitnil,omitempty" name:"ConditionSet"`

	// 许可证要求信息列表
	ForbiddenSet []*LicenseRestriction `json:"ForbiddenSet,omitnil,omitempty" name:"ForbiddenSet"`

	// 许可证禁止信息列表
	PermissionSet []*LicenseRestriction `json:"PermissionSet,omitnil,omitempty" name:"PermissionSet"`
}

type LicenseRestriction struct {
	// license约束的名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// license约束的描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type LicenseSummary struct {
	// 许可证标识符
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 许可证的SPDX标识符，见 https://spdx.org/licenses/
	SPDXKey *string `json:"SPDXKey,omitnil,omitempty" name:"SPDXKey"`

	// 许可证短名称
	ShortName *string `json:"ShortName,omitnil,omitempty" name:"ShortName"`

	// 许可证完整名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// License风险等级
	// <li>NotDefined</li>
	// <li>LowRisk</li>
	// <li>MediumRisk</li>
	// <li>HighRisk</li>
	Risk *string `json:"Risk,omitnil,omitempty" name:"Risk"`

	// 许可证来源URL
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`
}

type LicenseUnion struct {
	// 许可证概览信息
	LicenseSummary *LicenseSummary `json:"LicenseSummary,omitnil,omitempty" name:"LicenseSummary"`

	// 许可证详细信息
	LicenseDetail *LicenseDetail `json:"LicenseDetail,omitnil,omitempty" name:"LicenseDetail"`
}

// Predefined struct for user
type MatchKBPURLListRequestParams struct {
	// SHA1。
	SHA1 *string `json:"SHA1,omitnil,omitempty" name:"SHA1"`
}

type MatchKBPURLListRequest struct {
	*tchttp.BaseRequest
	
	// SHA1。
	SHA1 *string `json:"SHA1,omitnil,omitempty" name:"SHA1"`
}

func (r *MatchKBPURLListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MatchKBPURLListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SHA1")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MatchKBPURLListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MatchKBPURLListResponseParams struct {
	// 组件列表。
	PURLList []*PURL `json:"PURLList,omitnil,omitempty" name:"PURLList"`

	// 是否命中数据库。
	Hit *bool `json:"Hit,omitnil,omitempty" name:"Hit"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type MatchKBPURLListResponse struct {
	*tchttp.BaseResponse
	Response *MatchKBPURLListResponseParams `json:"Response"`
}

func (r *MatchKBPURLListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MatchKBPURLListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PURL struct {
	// 组件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 组件所属的类型，如：github, gitlab, generic, deb, rpm, maven 等
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 组件名的前缀名，如github和gitlab的用户名，deb的操作系统，maven包的group id等
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 修饰组件的额外属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Qualifiers []*Qualifier `json:"Qualifiers,omitnil,omitempty" name:"Qualifiers"`

	// 相对于组件包根位置的子目录
	Subpath *string `json:"Subpath,omitnil,omitempty" name:"Subpath"`

	// 组件版本号
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`
}

type Qualifier struct {
	// 额外属性的名称。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 额外属性的值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type SearchKBComponentRequestParams struct {
	// 需要搜索的组件名
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 需要搜索的组件类型
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 分页参数，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页参数，设置每页返回的结果数量
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type SearchKBComponentRequest struct {
	*tchttp.BaseRequest
	
	// 需要搜索的组件名
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 需要搜索的组件类型
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 分页参数，从 0 开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页参数，设置每页返回的结果数量
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *SearchKBComponentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchKBComponentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Query")
	delete(f, "Protocol")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchKBComponentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchKBComponentResponseParams struct {
	// 满足搜索条件的组件列表
	ComponentList []*Component `json:"ComponentList,omitnil,omitempty" name:"ComponentList"`

	// 满足搜索条件的总个数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchKBComponentResponse struct {
	*tchttp.BaseResponse
	Response *SearchKBComponentResponseParams `json:"Response"`
}

func (r *SearchKBComponentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchKBComponentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VulnerabilityDetail struct {
	// 漏洞类别
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 漏洞分类
	CategoryType *string `json:"CategoryType,omitnil,omitempty" name:"CategoryType"`

	// 漏洞描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 漏洞官方解决方案
	OfficialSolution *string `json:"OfficialSolution,omitnil,omitempty" name:"OfficialSolution"`

	// 漏洞信息参考列表
	ReferenceList []*string `json:"ReferenceList,omitnil,omitempty" name:"ReferenceList"`

	// 漏洞防御方案
	DefenseSolution *string `json:"DefenseSolution,omitnil,omitempty" name:"DefenseSolution"`

	// 漏洞CVSSv2信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVSSv2Info *CVSSV2Info `json:"CVSSv2Info,omitnil,omitempty" name:"CVSSv2Info"`

	// 漏洞CVSSv3信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVSSv3Info *CVSSV3Info `json:"CVSSv3Info,omitnil,omitempty" name:"CVSSv3Info"`

	// 漏洞提交时间
	SubmitTime *string `json:"SubmitTime,omitnil,omitempty" name:"SubmitTime"`

	// 漏洞更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// CWE编号
	CWEID *string `json:"CWEID,omitnil,omitempty" name:"CWEID"`

	// 漏洞CVSSv2向量
	CVSSv2Vector *string `json:"CVSSv2Vector,omitnil,omitempty" name:"CVSSv2Vector"`

	// 漏洞CVSSv3向量
	CVSSv3Vector *string `json:"CVSSv3Vector,omitnil,omitempty" name:"CVSSv3Vector"`

	// 漏洞影响的组件列表，仅当查询单个漏洞时有效
	AffectedComponentList []*AffectedComponent `json:"AffectedComponentList,omitnil,omitempty" name:"AffectedComponentList"`
}

type VulnerabilitySummary struct {
	// 漏洞ID
	VulID *string `json:"VulID,omitnil,omitempty" name:"VulID"`

	// 漏洞所属CVE编号
	CVEID *string `json:"CVEID,omitnil,omitempty" name:"CVEID"`

	// 漏洞所属CNVD编号
	CNVDID *string `json:"CNVDID,omitnil,omitempty" name:"CNVDID"`

	// 漏洞所属CNNVD编号
	CNNVDID *string `json:"CNNVDID,omitnil,omitempty" name:"CNNVDID"`

	// 漏洞名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 该漏洞是否是需重点关注的漏洞
	IsSuggest *bool `json:"IsSuggest,omitnil,omitempty" name:"IsSuggest"`

	// 漏洞风险等级
	// <li>Critical</li>
	// <li>High</li>
	// <li>Medium</li>
	// <li>Low</li>
	Severity *string `json:"Severity,omitnil,omitempty" name:"Severity"`

	// 架构信息，如x86、ARM等，废弃，请使用ArchitectureList
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Architecture is deprecated.
	Architecture []*string `json:"Architecture,omitnil,omitempty" name:"Architecture"`

	// 架构信息，如x86、ARM等
	// 注意：此字段可能返回 null，表示取不到有效值。
	ArchitectureList []*string `json:"ArchitectureList,omitnil,omitempty" name:"ArchitectureList"`

	// patch链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	PatchUrlList []*string `json:"PatchUrlList,omitnil,omitempty" name:"PatchUrlList"`
}

type VulnerabilityUnion struct {
	// 漏洞概览信息
	Summary *VulnerabilitySummary `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 漏洞详细信息
	Detail *VulnerabilityDetail `json:"Detail,omitnil,omitempty" name:"Detail"`
}