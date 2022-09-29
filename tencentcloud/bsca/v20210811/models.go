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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CVSSV2Info struct {
	// CVE评分。
	CVSS *float64 `json:"CVSS,omitempty" name:"CVSS"`

	// AccessVector 攻击途径。
	// 取值范围：
	// <li>NETWORK 远程</li>
	// <li>ADJACENT_NETWORK 近邻</li>
	// <li>LOCAL 本地</li>
	AccessVector *string `json:"AccessVector,omitempty" name:"AccessVector"`

	// AccessComplexity 攻击复杂度。
	// 取值范围：
	// <li>HIGH 高</li>
	// <li>MEDIUM 中</li>
	// <li>LOW 低</li>
	AccessComplexity *string `json:"AccessComplexity,omitempty" name:"AccessComplexity"`

	// Authentication 身份验证。
	// 取值范围：
	// <li>MULTIPLE 多系统认证</li>
	// <li>SINGLE 单系统认证</li>
	// <li>NONE 无</li>
	Authentication *string `json:"Authentication,omitempty" name:"Authentication"`

	// ConfidentialityImpact 机密性影响。
	// 取值范围：
	// <li>NONE 无</li>
	// <li>PARTIAL 部分</li>
	// <li>COMPLETE 完整</li>
	ConImpact *string `json:"ConImpact,omitempty" name:"ConImpact"`

	// IntegrityImpact 完整性影响。
	// 取值范围：
	// <li>NONE 无</li>
	// <li>PARTIAL 部分</li>
	// <li>COMPLETE 完整</li>
	IntegrityImpact *string `json:"IntegrityImpact,omitempty" name:"IntegrityImpact"`

	// AvailabilityImpact 可用性影响。
	// 取值范围：
	// <li>NONE 无</li>
	// <li>PARTIAL 部分</li>
	// <li>COMPLETE 完整</li>
	AvailabilityImpact *string `json:"AvailabilityImpact,omitempty" name:"AvailabilityImpact"`
}

type CVSSV3Info struct {
	// CVE评分。
	CVSS *float64 `json:"CVSS,omitempty" name:"CVSS"`

	// AttackVector 攻击途径。
	// 取值范围：
	// <li>NETWORK 远程</li>
	// <li>ADJACENT_NETWORK 近邻</li>
	// <li>LOCAL 本地</li>
	// <li>PHYSICAL 物理</li>
	AttackVector *string `json:"AttackVector,omitempty" name:"AttackVector"`

	// AttackComplexity 攻击复杂度。
	// 取值范围：
	// <li>HIGH 高</li>
	// <li>LOW 低</li>
	AttackComplexity *string `json:"AttackComplexity,omitempty" name:"AttackComplexity"`

	// PrivilegesRequired 触发特权。
	// 取值范围：
	// <li>HIGH 高</li>
	// <li>LOW 低</li>
	// <li>NONE 无</li>
	PrivilegesRequired *string `json:"PrivilegesRequired,omitempty" name:"PrivilegesRequired"`

	// UserInteraction 交互必要性。
	// 取值范围：
	// <li>NONE 无</li>
	// <li>REQUIRED 需要</li>
	UserInteraction *string `json:"UserInteraction,omitempty" name:"UserInteraction"`

	// Scope 绕过安全边界。
	// 取值范围：
	// <li>UNCHANGED 否</li>
	// <li>CHANGED 能</li>
	Scope *string `json:"Scope,omitempty" name:"Scope"`

	// ConfidentialityImpact 机密性影响。
	// 取值范围：
	// <li>NONE 无</li>
	// <li>LOW 低</li>
	// <li>HIGH 高</li>
	ConImpact *string `json:"ConImpact,omitempty" name:"ConImpact"`

	// IntegrityImpact 完整性影响。
	// 取值范围：
	// <li>NONE 无</li>
	// <li>LOW 低</li>
	// <li>HIGH 高</li>
	IntegrityImpact *string `json:"IntegrityImpact,omitempty" name:"IntegrityImpact"`

	// AvailabilityImpact 可用性影响。
	// 取值范围：
	// <li>NONE 无</li>
	// <li>LOW 低</li>
	// <li>HIGH 高</li>
	AvailabilityImpact *string `json:"AvailabilityImpact,omitempty" name:"AvailabilityImpact"`
}

type Component struct {
	// 第三方组件的PURL
	PURL *PURL `json:"PURL,omitempty" name:"PURL"`

	// 第三方组件的主页
	Homepage *string `json:"Homepage,omitempty" name:"Homepage"`

	// 第三方组件的简介
	Summary *string `json:"Summary,omitempty" name:"Summary"`

	// 第三方组件的别名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NicknameList []*string `json:"NicknameList,omitempty" name:"NicknameList"`

	// 第三方组件的代码位置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeLocationList []*string `json:"CodeLocationList,omitempty" name:"CodeLocationList"`

	// 第三方组件的许可证表达式
	LicenseExpression *string `json:"LicenseExpression,omitempty" name:"LicenseExpression"`
}

type ComponentVulnerabilitySummary struct {
	// 用于匹配漏洞的PURL
	// 注意：此字段可能返回 null，表示取不到有效值。
	PURL *PURL `json:"PURL,omitempty" name:"PURL"`

	// 该组件是否包含修复漏洞的官方补丁
	CanBeFixed *bool `json:"CanBeFixed,omitempty" name:"CanBeFixed"`

	// 修复漏洞的组件版本号
	FixedVersion *string `json:"FixedVersion,omitempty" name:"FixedVersion"`

	// 漏洞影响的组件版本号
	AffectedVersion *string `json:"AffectedVersion,omitempty" name:"AffectedVersion"`

	// 漏洞影响组件
	AffectedComponent *string `json:"AffectedComponent,omitempty" name:"AffectedComponent"`

	// 漏洞在该产品中的风险等级
	// <li>Critical</li>
	// <li>High</li>
	// <li>Medium</li>
	// <li>Low</li>
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`
}

type ComponentVulnerabilityUnion struct {
	// 漏洞概览信息
	Summary *VulnerabilitySummary `json:"Summary,omitempty" name:"Summary"`

	// 与组件相关的漏洞概览信息
	SummaryInComponent *ComponentVulnerabilitySummary `json:"SummaryInComponent,omitempty" name:"SummaryInComponent"`
}

// Predefined struct for user
type DescribeKBComponentRequestParams struct {
	// 组件的PURL
	PURL *PURL `json:"PURL,omitempty" name:"PURL"`
}

type DescribeKBComponentRequest struct {
	*tchttp.BaseRequest
	
	// 组件的PURL
	PURL *PURL `json:"PURL,omitempty" name:"PURL"`
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
	Component *Component `json:"Component,omitempty" name:"Component"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeKBComponentVulnerabilityRequestParams struct {
	// 组件的PURL，其中Name和Version为必填字段
	PURL *PURL `json:"PURL,omitempty" name:"PURL"`
}

type DescribeKBComponentVulnerabilityRequest struct {
	*tchttp.BaseRequest
	
	// 组件的PURL，其中Name和Version为必填字段
	PURL *PURL `json:"PURL,omitempty" name:"PURL"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKBComponentVulnerabilityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKBComponentVulnerabilityResponseParams struct {
	// 漏洞信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityList []*ComponentVulnerabilityUnion `json:"VulnerabilityList,omitempty" name:"VulnerabilityList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LicenseExpression *string `json:"LicenseExpression,omitempty" name:"LicenseExpression"`
}

type DescribeKBLicenseRequest struct {
	*tchttp.BaseRequest
	
	// License表达式
	LicenseExpression *string `json:"LicenseExpression,omitempty" name:"LicenseExpression"`
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
	LicenseList []*LicenseUnion `json:"LicenseList,omitempty" name:"LicenseList"`

	// 用于匹配的License表达式
	NormalizedLicenseExpression *string `json:"NormalizedLicenseExpression,omitempty" name:"NormalizedLicenseExpression"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// CVE ID列表（不能与Vul ID同时存在）
	CVEID []*string `json:"CVEID,omitempty" name:"CVEID"`

	// Vul ID列表（不能与CVE ID 同时存在）
	VulID []*string `json:"VulID,omitempty" name:"VulID"`
}

type DescribeKBVulnerabilityRequest struct {
	*tchttp.BaseRequest
	
	// CVE ID列表（不能与Vul ID同时存在）
	CVEID []*string `json:"CVEID,omitempty" name:"CVEID"`

	// Vul ID列表（不能与CVE ID 同时存在）
	VulID []*string `json:"VulID,omitempty" name:"VulID"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKBVulnerabilityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKBVulnerabilityResponseParams struct {
	// 漏洞详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VulnerabilityDetailList []*VulnerabilityUnion `json:"VulnerabilityDetailList,omitempty" name:"VulnerabilityDetailList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Content *string `json:"Content,omitempty" name:"Content"`

	// 许可证允许信息列表
	ConditionSet []*LicenseRestriction `json:"ConditionSet,omitempty" name:"ConditionSet"`

	// 许可证要求信息列表
	ForbiddenSet []*LicenseRestriction `json:"ForbiddenSet,omitempty" name:"ForbiddenSet"`

	// 许可证禁止信息列表
	PermissionSet []*LicenseRestriction `json:"PermissionSet,omitempty" name:"PermissionSet"`
}

type LicenseRestriction struct {
	// license约束的名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// license约束的描述。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type LicenseSummary struct {
	// 许可证标识符
	Key *string `json:"Key,omitempty" name:"Key"`

	// 许可证的SPDX标识符，见 https://spdx.org/licenses/
	SPDXKey *string `json:"SPDXKey,omitempty" name:"SPDXKey"`

	// 许可证短名称
	ShortName *string `json:"ShortName,omitempty" name:"ShortName"`

	// 许可证完整名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// License风险等级
	// <li>NotDefined</li>
	// <li>LowRisk</li>
	// <li>MediumRisk</li>
	// <li>HighRisk</li>
	Risk *string `json:"Risk,omitempty" name:"Risk"`

	// 许可证来源URL
	Source *string `json:"Source,omitempty" name:"Source"`
}

type LicenseUnion struct {
	// 许可证概览信息
	LicenseSummary *LicenseSummary `json:"LicenseSummary,omitempty" name:"LicenseSummary"`

	// 许可证详细信息
	LicenseDetail *LicenseDetail `json:"LicenseDetail,omitempty" name:"LicenseDetail"`
}

// Predefined struct for user
type MatchKBPURLListRequestParams struct {
	// SHA1。
	SHA1 *string `json:"SHA1,omitempty" name:"SHA1"`
}

type MatchKBPURLListRequest struct {
	*tchttp.BaseRequest
	
	// SHA1。
	SHA1 *string `json:"SHA1,omitempty" name:"SHA1"`
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
	PURLList []*PURL `json:"PURLList,omitempty" name:"PURLList"`

	// 是否命中数据库。
	Hit *bool `json:"Hit,omitempty" name:"Hit"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitempty" name:"Name"`

	// 组件所属的类型，如：github, gitlab, generic, deb, rpm, maven 等
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 组件名的前缀名，如github和gitlab的用户名，deb的操作系统，maven包的group id等
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 修饰组件的额外属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Qualifiers []*Qualifier `json:"Qualifiers,omitempty" name:"Qualifiers"`

	// 相对于组件包根位置的子目录
	Subpath *string `json:"Subpath,omitempty" name:"Subpath"`

	// 组件版本号
	Version *string `json:"Version,omitempty" name:"Version"`
}

type Qualifier struct {
	// 额外属性的名称。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 额外属性的值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type VulnerabilityDetail struct {
	// 漏洞类别
	Category *string `json:"Category,omitempty" name:"Category"`

	// 漏洞分类
	CategoryType *string `json:"CategoryType,omitempty" name:"CategoryType"`

	// 漏洞描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 漏洞官方解决方案
	OfficialSolution *string `json:"OfficialSolution,omitempty" name:"OfficialSolution"`

	// 漏洞信息参考列表
	ReferenceList []*string `json:"ReferenceList,omitempty" name:"ReferenceList"`

	// 漏洞防御方案
	DefenseSolution *string `json:"DefenseSolution,omitempty" name:"DefenseSolution"`

	// 漏洞CVSSv2信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVSSv2Info *CVSSV2Info `json:"CVSSv2Info,omitempty" name:"CVSSv2Info"`

	// 漏洞CVSSv3信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CVSSv3Info *CVSSV3Info `json:"CVSSv3Info,omitempty" name:"CVSSv3Info"`

	// 漏洞提交时间
	SubmitTime *string `json:"SubmitTime,omitempty" name:"SubmitTime"`

	// CWE编号
	CWEID *string `json:"CWEID,omitempty" name:"CWEID"`

	// 漏洞CVSSv2向量
	CVSSv2Vector *string `json:"CVSSv2Vector,omitempty" name:"CVSSv2Vector"`

	// 漏洞CVSSv3向量
	CVSSv3Vector *string `json:"CVSSv3Vector,omitempty" name:"CVSSv3Vector"`
}

type VulnerabilitySummary struct {
	// 漏洞ID
	VulID *string `json:"VulID,omitempty" name:"VulID"`

	// 漏洞所属CVE编号
	CVEID *string `json:"CVEID,omitempty" name:"CVEID"`

	// 漏洞所属CNVD编号
	CNVDID *string `json:"CNVDID,omitempty" name:"CNVDID"`

	// 漏洞所属CNNVD编号
	CNNVDID *string `json:"CNNVDID,omitempty" name:"CNNVDID"`

	// 漏洞名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 该漏洞是否是需重点关注的漏洞
	IsSuggest *bool `json:"IsSuggest,omitempty" name:"IsSuggest"`

	// 漏洞风险等级
	// <li>Critical</li>
	// <li>High</li>
	// <li>Medium</li>
	// <li>Low</li>
	Severity *string `json:"Severity,omitempty" name:"Severity"`
}

type VulnerabilityUnion struct {
	// 漏洞概览信息
	Summary *VulnerabilitySummary `json:"Summary,omitempty" name:"Summary"`

	// 漏洞详细信息
	Detail *VulnerabilityDetail `json:"Detail,omitempty" name:"Detail"`
}