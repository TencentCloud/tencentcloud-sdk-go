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

package v20250611

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Account struct {
	// 租户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 租户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 剩余可用额度
	RemainQuota *int64 `json:"RemainQuota,omitnil,omitempty" name:"RemainQuota"`

	// 租户名称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 下发规则数
	DispatchRuleNum *int64 `json:"DispatchRuleNum,omitnil,omitempty" name:"DispatchRuleNum"`

	// 产品已有规则数
	OriginRuleNum *int64 `json:"OriginRuleNum,omitnil,omitempty" name:"OriginRuleNum"`

	// 总额度
	TotalQuota *int64 `json:"TotalQuota,omitnil,omitempty" name:"TotalQuota"`

	// 成员Id
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type AccountGroupInfo struct {
	// 账户组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 账户组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
}

type AccountGroupQuotaDetail struct {
	// 账号组Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 账号组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 账号组成员数
	MemberCount *int64 `json:"MemberCount,omitnil,omitempty" name:"MemberCount"`

	// 取组内 RemainQuota 最小成员的值
	RemainQuota *int64 `json:"RemainQuota,omitnil,omitempty" name:"RemainQuota"`

	// 同上成员的 TotalQuota
	TotalQuota *int64 `json:"TotalQuota,omitnil,omitempty" name:"TotalQuota"`

	// 同上成员的 DispatchRuleNum
	DispatchRuleNum *int64 `json:"DispatchRuleNum,omitnil,omitempty" name:"DispatchRuleNum"`

	// 同上成员的 OriginRuleNum
	OriginRuleNum *int64 `json:"OriginRuleNum,omitnil,omitempty" name:"OriginRuleNum"`

	// 配额最少的成员 Uin
	BottleneckUin *string `json:"BottleneckUin,omitnil,omitempty" name:"BottleneckUin"`

	// 成员列表
	Members []*Account `json:"Members,omitnil,omitempty" name:"Members"`
}

type AccountProductDetailStats struct {
	// 产品类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 体检策略数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyCount *int64 `json:"PolicyCount,omitnil,omitempty" name:"PolicyCount"`

	// 待整改风险数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UntreatedRiskCount *int64 `json:"UntreatedRiskCount,omitnil,omitempty" name:"UntreatedRiskCount"`

	// 总风险数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalRiskCount *int64 `json:"TotalRiskCount,omitnil,omitempty" name:"TotalRiskCount"`

	// 已处置数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreatedRiskCount *int64 `json:"TreatedRiskCount,omitnil,omitempty" name:"TreatedRiskCount"`

	// 已忽略数
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoredRiskCount *int64 `json:"IgnoredRiskCount,omitnil,omitempty" name:"IgnoredRiskCount"`

	// 整改率，如 50%，无需整改时为 无需整改
	// 注意：此字段可能返回 null，表示取不到有效值。
	RectifyRate *string `json:"RectifyRate,omitnil,omitempty" name:"RectifyRate"`

	// 最近一次体检时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastCheckTime *string `json:"LastCheckTime,omitnil,omitempty" name:"LastCheckTime"`

	// 子类 ID 列表
	SubcategoryIds []*string `json:"SubcategoryIds,omitnil,omitempty" name:"SubcategoryIds"`

	// 是否超时未体检
	IsOverdue *bool `json:"IsOverdue,omitnil,omitempty" name:"IsOverdue"`
}

type AccountStatsGroup struct {
	// 成员账号信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Member *MemberInfo `json:"Member,omitnil,omitempty" name:"Member"`

	// 该账号下所有产品待整改风险数汇总
	// 注意：此字段可能返回 null，表示取不到有效值。
	UntreatedRiskCount *int64 `json:"UntreatedRiskCount,omitnil,omitempty" name:"UntreatedRiskCount"`

	// 该账号下各产品维度的风险统计列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductStats []*AccountProductDetailStats `json:"ProductStats,omitnil,omitempty" name:"ProductStats"`

	// 整改率
	RectifyRate *string `json:"RectifyRate,omitnil,omitempty" name:"RectifyRate"`
}

type AddressTemplateSpecification struct {
	// IP地址ID，例如：ipm-2uw6ujo6。
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`

	//  IP地址组ID，例如：ipmg-2uw6ujo6。
	AddressGroupId *string `json:"AddressGroupId,omitnil,omitempty" name:"AddressGroupId"`
}

type AnalysisSgRuleInfoResp struct {
	// <p>规则id  等同RuleUuid</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>规则Id</p>
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// <p>排序</p>
	OrderIndex *int64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// <p>云防排序</p>
	CfwOrderIndex *int64 `json:"CfwOrderIndex,omitnil,omitempty" name:"CfwOrderIndex"`

	// <p>源规则内容</p>
	SourceId *string `json:"SourceId,omitnil,omitempty" name:"SourceId"`

	// <p>源规则类型<br>取值范围 0/1/2/3/4/5/6/7/8/9<br>0表示ip(net),<br>1表示VPC实例(instance)<br>2表示子网实例(instance)<br>3表示CVM实例(instance)<br>4表示CLB实例(instance)<br>5表示ENI实例(instance)<br>6表示数据库实例(instance)<br>7表示模板(template)<br>8表示标签(tag)<br>9表示地域(region)</p><p>枚举值：</p><ul><li>0： IP / CIDR</li><li>1： VPC 实例</li><li>2： 子网 </li><li>3： CVM 实例</li><li>4： CLB 实例</li><li>5： ENI（弹性网卡）实例</li><li>6： CDB（云数据库）实例</li><li>7： 参数模板</li><li>8： 标签</li><li>9： 地域</li></ul>
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// <p>目的规则内容</p>
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// <p>目的规则类型<br>取值范围 0/1/2/3/4/5/6/7/8/9/100<br>0表示ip(net),<br>1表示VPC实例(instance)<br>2表示子网实例(instance)<br>3表示CVM实例(instance)<br>4表示CLB实例(instance)<br>5表示ENI实例(instance)<br>6表示数据库实例(instance)<br>7表示模板(template)<br>8表示标签(tag)<br>9表示地域(region)<br>100表示资产分组(resourcegroup)</p><p>枚举值：</p><ul><li>0： IP / CIDR</li><li>1： VPC 实例</li><li>2： 子网 </li><li>3： CVM 实例</li><li>4： CLB 实例</li><li>5： ENI（弹性网卡）实例</li><li>6： CDB（云数据库）实例</li><li>7： 参数模板</li><li>8： 标签</li><li>9： 地域</li></ul>
	TargetType *int64 `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// <p>协议名称<br>取值范围:TCP/ANY/ICMP/UDP<br>ANY:表示所有</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>端口</p>
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>规则策略<br>取值范围:1/2<br>1:阻断<br>2:放行</p>
	Strategy *int64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// <p>描述</p>
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// <p>地域</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>服务模板id</p>
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`

	// <p>源资产名称</p>
	SouInstanceName *string `json:"SouInstanceName,omitnil,omitempty" name:"SouInstanceName"`

	// <p>源资产公网ip</p>
	SouPublicIp *string `json:"SouPublicIp,omitnil,omitempty" name:"SouPublicIp"`

	// <p>源资产内网ip</p>
	SouPrivateIp *string `json:"SouPrivateIp,omitnil,omitempty" name:"SouPrivateIp"`

	// <p>源资产网段信息</p>
	SouCidr *string `json:"SouCidr,omitnil,omitempty" name:"SouCidr"`

	// <p>源模板名称</p>
	SouParameterName *string `json:"SouParameterName,omitnil,omitempty" name:"SouParameterName"`

	// <p>目的资产名称</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>目的资产公网ip</p>
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// <p>目的资产内网ip</p>
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// <p>目的资产网段信息</p>
	Cidr *string `json:"Cidr,omitnil,omitempty" name:"Cidr"`

	// <p>目的模板名称</p>
	ParameterName *string `json:"ParameterName,omitnil,omitempty" name:"ParameterName"`

	// <p>端口模板名称</p>
	ProtocolPortName *string `json:"ProtocolPortName,omitnil,omitempty" name:"ProtocolPortName"`

	// <p>域名解析的IP统计</p>
	DnsParseCount *SgDnsParseCount `json:"DnsParseCount,omitnil,omitempty" name:"DnsParseCount"`

	// <p>规则生效范围</p>
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>分区：<br>1防火墙管理最前分区<br>2是云防规则<br>3防火墙管理最后分区</p>
	RulePartition *int64 `json:"RulePartition,omitnil,omitempty" name:"RulePartition"`

	// <p>规则组Id</p>
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// <p>规则组名称</p>
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// <p>规则组内规则id</p>
	GroupRuleId *string `json:"GroupRuleId,omitnil,omitempty" name:"GroupRuleId"`

	// <p>策略Id</p>
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// <p>ip类型</p>
	IpVersion *string `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`

	// <p>成员信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BelongMember *MemberInfo `json:"BelongMember,omitnil,omitempty" name:"BelongMember"`
}

// Predefined struct for user
type CancelIgnorePolicyRiskRequestParams struct {
	// 风险ID
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// 成员Id
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type CancelIgnorePolicyRiskRequest struct {
	*tchttp.BaseRequest
	
	// 风险ID
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// 成员Id
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *CancelIgnorePolicyRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelIgnorePolicyRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RiskId")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelIgnorePolicyRiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelIgnorePolicyRiskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelIgnorePolicyRiskResponse struct {
	*tchttp.BaseResponse
	Response *CancelIgnorePolicyRiskResponseParams `json:"Response"`
}

func (r *CancelIgnorePolicyRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelIgnorePolicyRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommonFilter struct {
	// <p>筛选字段名。支持：SecurityGroupId、FwGroupId、IP（IP地址模糊搜索）、InstanceName（实例名称模糊搜索）、VpcId（VPC ID精确搜索）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>筛选值列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// <p>操作类型。1=等于，7=in，9=模糊匹配</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorType *int64 `json:"OperatorType,omitnil,omitempty" name:"OperatorType"`
}

// Predefined struct for user
type CreateAnalyzePolicyTaskRequestParams struct {
	// 产品类型
	Products []*string `json:"Products,omitnil,omitempty" name:"Products"`

	// 成员Id 列表
	MemberIdSet []*string `json:"MemberIdSet,omitnil,omitempty" name:"MemberIdSet"`
}

type CreateAnalyzePolicyTaskRequest struct {
	*tchttp.BaseRequest
	
	// 产品类型
	Products []*string `json:"Products,omitnil,omitempty" name:"Products"`

	// 成员Id 列表
	MemberIdSet []*string `json:"MemberIdSet,omitnil,omitempty" name:"MemberIdSet"`
}

func (r *CreateAnalyzePolicyTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAnalyzePolicyTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Products")
	delete(f, "MemberIdSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAnalyzePolicyTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAnalyzePolicyTaskResponseParams struct {
	// 任务状态 ，1 表示执行中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAnalyzePolicyTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateAnalyzePolicyTaskResponseParams `json:"Response"`
}

func (r *CreateAnalyzePolicyTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAnalyzePolicyTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeAclRuleGroupRequestParams struct {
	// 规则组名称，长度1-50字符
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 产品类型，固定为 cfw_edge_acl
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 规则列表
	Rules []*EdgeAclRuleInfo `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type CreateEdgeAclRuleGroupRequest struct {
	*tchttp.BaseRequest
	
	// 规则组名称，长度1-50字符
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 产品类型，固定为 cfw_edge_acl
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 规则列表
	Rules []*EdgeAclRuleInfo `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *CreateEdgeAclRuleGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeAclRuleGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "Product")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEdgeAclRuleGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeAclRuleGroupResponseParams struct {
	// 创建的规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEdgeAclRuleGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateEdgeAclRuleGroupResponseParams `json:"Response"`
}

func (r *CreateEdgeAclRuleGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeAclRuleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeAclRuleRequestParams struct {
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则列表
	Rules []*EdgeAclRuleInfo `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type CreateEdgeAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则列表
	Rules []*EdgeAclRuleInfo `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *CreateEdgeAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEdgeAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEdgeAclRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEdgeAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateEdgeAclRuleResponseParams `json:"Response"`
}

func (r *CreateEdgeAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEdgeAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatAclRuleGroupRequestParams struct {
	// 规则组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 产品类型，固定为 cfw_nat_acl
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 规则列表
	Rules []*NatAclRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type CreateNatAclRuleGroupRequest struct {
	*tchttp.BaseRequest
	
	// 规则组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 产品类型，固定为 cfw_nat_acl
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 规则列表
	Rules []*NatAclRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *CreateNatAclRuleGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatAclRuleGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "Product")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNatAclRuleGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatAclRuleGroupResponseParams struct {
	// 创建的规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNatAclRuleGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateNatAclRuleGroupResponseParams `json:"Response"`
}

func (r *CreateNatAclRuleGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatAclRuleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatAclRuleRequestParams struct {
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则列表
	Rules []*NatAclRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type CreateNatAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则列表
	Rules []*NatAclRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *CreateNatAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNatAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatAclRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNatAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateNatAclRuleResponseParams `json:"Response"`
}

func (r *CreateNatAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupRuleGroupRequestParams struct {
	// 规则组ID
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 产品类型
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 规则列表
	Rules []*SecurityGroupRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type CreateSecurityGroupRuleGroupRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 产品类型
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 规则列表
	Rules []*SecurityGroupRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *CreateSecurityGroupRuleGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupRuleGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "Product")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityGroupRuleGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupRuleGroupResponseParams struct {
	// 创建的规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSecurityGroupRuleGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityGroupRuleGroupResponseParams `json:"Response"`
}

func (r *CreateSecurityGroupRuleGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupRuleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupRuleRequestParams struct {
	// 规则组Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则列表
	Rules []*SecurityGroupRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type CreateSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则组Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则列表
	Rules []*SecurityGroupRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *CreateSecurityGroupRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityGroupRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSecurityGroupRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityGroupRuleResponseParams `json:"Response"`
}

func (r *CreateSecurityGroupRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStrategyRequestParams struct {
	// 产品类型
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 下发账号
	ReceiveAccount []*string `json:"ReceiveAccount,omitnil,omitempty" name:"ReceiveAccount"`

	// 前区规则组
	PreStrategy []*StrategyReq `json:"PreStrategy,omitnil,omitempty" name:"PreStrategy"`

	// 后区规则组
	PostStrategy []*StrategyReq `json:"PostStrategy,omitnil,omitempty" name:"PostStrategy"`

	// 下发账号组
	ReceiveGroup []*string `json:"ReceiveGroup,omitnil,omitempty" name:"ReceiveGroup"`
}

type CreateStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 产品类型
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 下发账号
	ReceiveAccount []*string `json:"ReceiveAccount,omitnil,omitempty" name:"ReceiveAccount"`

	// 前区规则组
	PreStrategy []*StrategyReq `json:"PreStrategy,omitnil,omitempty" name:"PreStrategy"`

	// 后区规则组
	PostStrategy []*StrategyReq `json:"PostStrategy,omitnil,omitempty" name:"PostStrategy"`

	// 下发账号组
	ReceiveGroup []*string `json:"ReceiveGroup,omitnil,omitempty" name:"ReceiveGroup"`
}

func (r *CreateStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "ReceiveAccount")
	delete(f, "PreStrategy")
	delete(f, "PostStrategy")
	delete(f, "ReceiveGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStrategyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStrategyResponse struct {
	*tchttp.BaseResponse
	Response *CreateStrategyResponseParams `json:"Response"`
}

func (r *CreateStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcAclRuleGroupRequestParams struct {
	// 规则组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 产品类型，固定为 cfw_vpc_acl
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 规则列表
	Rules []*VpcAclRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type CreateVpcAclRuleGroupRequest struct {
	*tchttp.BaseRequest
	
	// 规则组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 产品类型，固定为 cfw_vpc_acl
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 规则列表
	Rules []*VpcAclRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *CreateVpcAclRuleGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcAclRuleGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "Product")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpcAclRuleGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcAclRuleGroupResponseParams struct {
	// 创建的规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVpcAclRuleGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpcAclRuleGroupResponseParams `json:"Response"`
}

func (r *CreateVpcAclRuleGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcAclRuleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcAclRuleRequestParams struct {
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则列表
	Rules []*VpcAclRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type CreateVpcAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则列表
	Rules []*VpcAclRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *CreateVpcAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpcAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcAclRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVpcAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpcAclRuleResponseParams `json:"Response"`
}

func (r *CreateVpcAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeAclRuleRequestParams struct {
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 要删除的规则ID列表
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

type DeleteEdgeAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 要删除的规则ID列表
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

func (r *DeleteEdgeAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEdgeAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEdgeAclRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEdgeAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEdgeAclRuleResponseParams `json:"Response"`
}

func (r *DeleteEdgeAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEdgeAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNatAclRuleRequestParams struct {
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则ID列表
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

type DeleteNatAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则ID列表
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

func (r *DeleteNatAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNatAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNatAclRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNatAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNatAclRuleResponseParams `json:"Response"`
}

func (r *DeleteNatAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRuleGroupRequestParams struct {
	// 规则组Id列表
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

type DeleteRuleGroupRequest struct {
	*tchttp.BaseRequest
	
	// 规则组Id列表
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

func (r *DeleteRuleGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRuleGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRuleGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRuleGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRuleGroupResponseParams `json:"Response"`
}

func (r *DeleteRuleGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupRuleRequestParams struct {
	// 规则组Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则列表
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

type DeleteSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则组Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则列表
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

func (r *DeleteSecurityGroupRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityGroupRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSecurityGroupRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityGroupRuleResponseParams `json:"Response"`
}

func (r *DeleteSecurityGroupRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStrategyRequestParams struct {
	// 策略Id列表
	StrategyIds []*string `json:"StrategyIds,omitnil,omitempty" name:"StrategyIds"`
}

type DeleteStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 策略Id列表
	StrategyIds []*string `json:"StrategyIds,omitnil,omitempty" name:"StrategyIds"`
}

func (r *DeleteStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StrategyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStrategyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStrategyResponseParams `json:"Response"`
}

func (r *DeleteStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcAclRuleRequestParams struct {
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则ID列表
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

type DeleteVpcAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则ID列表
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

func (r *DeleteVpcAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpcAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcAclRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVpcAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpcAclRuleResponseParams `json:"Response"`
}

func (r *DeleteVpcAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeAclRulesRequestParams struct {
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则方向：0-出向，1-入向
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 过滤条件列表，支持按 RuleId、Direction、Protocol、RuleAction 等字段过滤
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小，默认100，最大1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序顺序，asc:升序 desc:降序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段，支持 Sequence、RuleId 等
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeEdgeAclRulesRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则方向：0-出向，1-入向
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 过滤条件列表，支持按 RuleId、Direction、Protocol、RuleAction 等字段过滤
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小，默认100，最大1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序顺序，asc:升序 desc:降序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段，支持 Sequence、RuleId 等
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeEdgeAclRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeAclRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Direction")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEdgeAclRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEdgeAclRulesResponseParams struct {
	// 规则总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 规则列表
	Rules []*EdgeAclRuleResp `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 不过滤的总数
	AllTotalCount *int64 `json:"AllTotalCount,omitnil,omitempty" name:"AllTotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEdgeAclRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEdgeAclRulesResponseParams `json:"Response"`
}

func (r *DescribeEdgeAclRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEdgeAclRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatAclRulesRequestParams struct {
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则方向：0-出向，1-入向
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 过滤条件
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序顺序，asc:升序 desc:降序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeNatAclRulesRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则方向：0-出向，1-入向
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 过滤条件
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序顺序，asc:升序 desc:降序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeNatAclRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatAclRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Direction")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatAclRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatAclRulesResponseParams struct {
	// 规则列表
	Rules []*NatAclRuleResp `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 不过滤的总数
	AllTotalCount *int64 `json:"AllTotalCount,omitnil,omitempty" name:"AllTotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNatAclRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatAclRulesResponseParams `json:"Response"`
}

func (r *DescribeNatAclRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatAclRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganMembersRequestParams struct {
	// 搜索过滤条件列表，支持按成员 ID、账号名称、身份、纳管状态等字段筛选
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小，默认 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，如 MemberCreateTime
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 排序方式：asc 升序，desc 降序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type DescribeOrganMembersRequest struct {
	*tchttp.BaseRequest
	
	// 搜索过滤条件列表，支持按成员 ID、账号名称、身份、纳管状态等字段筛选
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页大小，默认 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，默认 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，如 MemberCreateTime
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 排序方式：asc 升序，desc 降序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

func (r *DescribeOrganMembersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganMembersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "By")
	delete(f, "Order")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganMembersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganMembersResponseParams struct {
	// 集团成员总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 集团成员列表
	Members []*OrganMemberItem `json:"Members,omitnil,omitempty" name:"Members"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganMembersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganMembersResponseParams `json:"Response"`
}

func (r *DescribeOrganMembersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganMembersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganSummaryRequestParams struct {

}

type DescribeOrganSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeOrganSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganSummaryResponseParams struct {
	// 集团概览
	Summary *OrganSummary `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOrganSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganSummaryResponseParams `json:"Response"`
}

func (r *DescribeOrganSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyRiskAccountProductStatsRequestParams struct {
	// 分页大小，按账号分页，默认20，最大100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 筛选条件列表。支持的筛选字段：AccountName（账号名称模糊搜索）、AccountId（账号Uin精确搜索）、OnlyUntreated（仅看待整改，值为1时生效）、OnlyOverdue（仅超时未体检，值为1时生效）
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribePolicyRiskAccountProductStatsRequest struct {
	*tchttp.BaseRequest
	
	// 分页大小，按账号分页，默认20，最大100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 筛选条件列表。支持的筛选字段：AccountName（账号名称模糊搜索）、AccountId（账号Uin精确搜索）、OnlyUntreated（仅看待整改，值为1时生效）、OnlyOverdue（仅超时未体检，值为1时生效）
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribePolicyRiskAccountProductStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyRiskAccountProductStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePolicyRiskAccountProductStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePolicyRiskAccountProductStatsResponseParams struct {
	// 按账号分组的风险统计列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountStats []*AccountStatsGroup `json:"AccountStats,omitnil,omitempty" name:"AccountStats"`

	// 满足条件的账号总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 超时未体检的账号数
	OverdueAccountCount *int64 `json:"OverdueAccountCount,omitnil,omitempty" name:"OverdueAccountCount"`

	// 超时未体检的产品数
	OverdueProductCount *int64 `json:"OverdueProductCount,omitnil,omitempty" name:"OverdueProductCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePolicyRiskAccountProductStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePolicyRiskAccountProductStatsResponseParams `json:"Response"`
}

func (r *DescribePolicyRiskAccountProductStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePolicyRiskAccountProductStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskAnalysisDetailsRequestParams struct {
	// 风险ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 查询类型，analyze实时数据分析，task定时分析结果
	SearchType *string `json:"SearchType,omitnil,omitempty" name:"SearchType"`

	// 成员Id
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeRiskAnalysisDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 风险ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 查询类型，analyze实时数据分析，task定时分析结果
	SearchType *string `json:"SearchType,omitnil,omitempty" name:"SearchType"`

	// 成员Id
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeRiskAnalysisDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskAnalysisDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "SearchType")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskAnalysisDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskAnalysisDetailsResponseParams struct {
	// 风险企业安全组规则列表
	EnterpriseSecurityGroupRule []*AnalysisSgRuleInfoResp `json:"EnterpriseSecurityGroupRule,omitnil,omitempty" name:"EnterpriseSecurityGroupRule"`

	// 风险安全组规则列表
	SecurityGroupPolicy []*SecurityGroupRiskPolicy `json:"SecurityGroupPolicy,omitnil,omitempty" name:"SecurityGroupPolicy"`

	// 实时分析状态，1分析执行中请轮询，0分析已结束
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 风险企业安全组IPV6规则列表
	EnterpriseSecurityGroupRuleIPV6 []*AnalysisSgRuleInfoResp `json:"EnterpriseSecurityGroupRuleIPV6,omitnil,omitempty" name:"EnterpriseSecurityGroupRuleIPV6"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskAnalysisDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskAnalysisDetailsResponseParams `json:"Response"`
}

func (r *DescribeRiskAnalysisDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskAnalysisDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCategoryStatsRequestParams struct {
	// 分页大小,默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量,默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 产品类型
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 筛选器
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段："RuleCount", "TreatedCount", "IgnoredCount", "UntreatedCount", "DisposalRate"
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 顺序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 成员Id
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeRiskCategoryStatsRequest struct {
	*tchttp.BaseRequest
	
	// 分页大小,默认20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量,默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 产品类型
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 筛选器
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段："RuleCount", "TreatedCount", "IgnoredCount", "UntreatedCount", "DisposalRate"
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 顺序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 成员Id
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeRiskCategoryStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCategoryStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Product")
	delete(f, "Filters")
	delete(f, "By")
	delete(f, "Order")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskCategoryStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskCategoryStatsResponseParams struct {
	// 风险分类总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 风险分类统计列表
	Data []*RiskCategoryItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskCategoryStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskCategoryStatsResponseParams `json:"Response"`
}

func (r *DescribeRiskCategoryStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskCategoryStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskListRequestParams struct {
	// 条数限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 产品类型
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 筛选条件
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 成员Id
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type DescribeRiskListRequest struct {
	*tchttp.BaseRequest
	
	// 条数限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 产品类型
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 筛选条件
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序方式
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 成员Id
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *DescribeRiskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Product")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRiskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRiskListResponseParams struct {
	// 策略问题列表
	PolicyRiskLst []*PolicyRisk `json:"PolicyRiskLst,omitnil,omitempty" name:"PolicyRiskLst"`

	// 策略问题数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRiskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRiskListResponseParams `json:"Response"`
}

func (r *DescribeRiskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRiskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupRuleRequestParams struct {
	// 规则组Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则Id
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type DescribeSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则组Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则Id
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

func (r *DescribeSecurityGroupRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupRuleResponseParams struct {
	// 规则详情
	Rule *SgRuleResp `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupRuleResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupRulesRequestParams struct {
	// 规则组Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 模糊搜索关键词
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSecurityGroupRulesRequest struct {
	*tchttp.BaseRequest
	
	// 规则组Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 模糊搜索关键词
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeSecurityGroupRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupRulesResponseParams struct {
	// 规则列表
	Rules []*SecGroupRuleResp `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 不过滤的规则总数
	AllTotalCount *int64 `json:"AllTotalCount,omitnil,omitempty" name:"AllTotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupRulesResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStrategiesRequestParams struct {
	// 产品类型
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 执行区域：pre是前区，post 是后区
	ExecArea *string `json:"ExecArea,omitnil,omitempty" name:"ExecArea"`

	// 筛选条件
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 条数限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeStrategiesRequest struct {
	*tchttp.BaseRequest
	
	// 产品类型
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 执行区域：pre是前区，post 是后区
	ExecArea *string `json:"ExecArea,omitnil,omitempty" name:"ExecArea"`

	// 筛选条件
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 条数限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeStrategiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStrategiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "ExecArea")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStrategiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStrategiesResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 策略列表
	Strategies []*StrategyResp `json:"Strategies,omitnil,omitempty" name:"Strategies"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStrategiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStrategiesResponseParams `json:"Response"`
}

func (r *DescribeStrategiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStrategiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStrategyAccountsRequestParams struct {
	// 下发产品 secgroup // 企业安全组
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 筛选器
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeStrategyAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 下发产品 secgroup // 企业安全组
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 筛选器
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeStrategyAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStrategyAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStrategyAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStrategyAccountsResponseParams struct {
	// 账号列表
	Accounts []*Account `json:"Accounts,omitnil,omitempty" name:"Accounts"`

	// 账号组列表
	AccountGroups []*AccountGroupQuotaDetail `json:"AccountGroups,omitnil,omitempty" name:"AccountGroups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStrategyAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStrategyAccountsResponseParams `json:"Response"`
}

func (r *DescribeStrategyAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStrategyAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStrategyDispatchStatusRequestParams struct {
	// 产品
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeStrategyDispatchStatusRequest struct {
	*tchttp.BaseRequest
	
	// 产品
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeStrategyDispatchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStrategyDispatchStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStrategyDispatchStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStrategyDispatchStatusResponseParams struct {
	// 进度
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 下发开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 下发结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 下发状态，0无变动，1下发中，2下发成功，3下发失败，4更新待下发
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 下发规则组数量
	RuleGroupNum *int64 `json:"RuleGroupNum,omitnil,omitempty" name:"RuleGroupNum"`

	// 下发失败错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 下发关联策略id列表
	DispatchStrategyList []*string `json:"DispatchStrategyList,omitnil,omitempty" name:"DispatchStrategyList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStrategyDispatchStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStrategyDispatchStatusResponseParams `json:"Response"`
}

func (r *DescribeStrategyDispatchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStrategyDispatchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStrategyRequestParams struct {
	// 策略Id
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`
}

type DescribeStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 策略Id
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`
}

func (r *DescribeStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StrategyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStrategyResponseParams struct {
	// 策略详情
	Strategy *StrategyResp `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStrategyResponseParams `json:"Response"`
}

func (r *DescribeStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcAclRulesRequestParams struct {
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 过滤条件
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeVpcAclRulesRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 过滤条件
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeVpcAclRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcAclRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcAclRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcAclRulesResponseParams struct {
	// 规则列表
	Rules []*VpcAclRuleResp `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 不过滤的总数
	AllTotalCount *int64 `json:"AllTotalCount,omitnil,omitempty" name:"AllTotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpcAclRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcAclRulesResponseParams `json:"Response"`
}

func (r *DescribeVpcAclRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcAclRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DispatchStrategyRequestParams struct {
	// 1:下发，2:中止
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 产品
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DispatchStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 1:下发，2:中止
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 产品
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DispatchStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DispatchStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DispatchStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DispatchStrategyResponseParams struct {
	// 返回状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DispatchStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DispatchStrategyResponseParams `json:"Response"`
}

func (r *DispatchStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DispatchStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EdgeAclRuleInfo struct {
	// 规则ID，修改规则时必填
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则执行顺序，数字越小优先级越高，创建规则组时必须从1开始严格递增
	OrderIndex *int64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// 规则方向：0-出站，1-入站
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 源地址内容，根据 SourceType 不同有不同的格式：ip 时为 IP/CIDR，domain 时为域名，template 时为模板ID，instance 时为实例ID列表（逗号分隔），tag 时为标签键值对（格式：key:value）
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// 源地址类型：ip-IP地址，domain-域名，template-参数模板，instance-实例，tag-标签
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 目标地址内容，格式同 SourceContent
	TargetContent *string `json:"TargetContent,omitnil,omitempty" name:"TargetContent"`

	// 目标地址类型：ip-IP地址，domain-域名，template-参数模板，instance-实例，tag-标签
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 端口，支持单端口、端口范围和逗号分隔的多端口，如：80、1-65535、80,443,8080
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 协议类型：TCP、UDP、ICMP、ANY
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 规则动作：accept-放行，drop-阻断，log-观察
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// 规则描述，长度0-256字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 生效范围：serial，串行；side，旁路；all，全局	
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 规则归属的成员账号ID（多账号场景下使用）。当 SourceType 或 TargetType 为 instance 或 tag 时，此参数必填，用于指定实例/标签所属的成员账号
	BelongMemberId *string `json:"BelongMemberId,omitnil,omitempty" name:"BelongMemberId"`

	// 参数模板
	ParamTemplateId *string `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`
}

type EdgeAclRuleResp struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则执行顺序
	Sequence *int64 `json:"Sequence,omitnil,omitempty" name:"Sequence"`

	// 规则方向：0-出站，1-入站
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 源地址内容
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// 源地址类型
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 源地址名称（当类型为模板/实例/标签时返回对应名称）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// 目标地址内容
	TargetContent *string `json:"TargetContent,omitnil,omitempty" name:"TargetContent"`

	// 目标地址类型
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 目标地址名称（当类型为模板/实例/标签时返回对应名称）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// 目标端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 协议类型
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 规则动作：accept-放行，drop-阻断，log-观察
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// 规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 规则生效范围
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 地域名称1（正则匹配时使用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountryName *string `json:"CountryName,omitnil,omitempty" name:"CountryName"`

	// 地域名称2（正则匹配时使用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CityName *string `json:"CityName,omitnil,omitempty" name:"CityName"`

	// 参数模板ID（当类型为模板时返回）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamTemplateId *string `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// 参数模板名称（当类型为模板时返回）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamTemplateName *string `json:"ParamTemplateName,omitnil,omitempty" name:"ParamTemplateName"`

	// 规则是否失效：0-有效，1-失效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Invalid *int64 `json:"Invalid,omitnil,omitempty" name:"Invalid"`

	// 规则归属的成员账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BelongMember *MemberInfo `json:"BelongMember,omitnil,omitempty" name:"BelongMember"`

	// 国家Id
	CountryCode *uint64 `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 城市Id
	CityCode *uint64 `json:"CityCode,omitnil,omitempty" name:"CityCode"`

	// 0为正常规则,1为地域规则
	IsRegion *uint64 `json:"IsRegion,omitnil,omitempty" name:"IsRegion"`

	// 云厂商code
	CloudCode *string `json:"CloudCode,omitnil,omitempty" name:"CloudCode"`

	// 0为正常规则,1为云厂商规则
	IsCloud *uint64 `json:"IsCloud,omitnil,omitempty" name:"IsCloud"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 地区简称
	CountryKey *string `json:"CountryKey,omitnil,omitempty" name:"CountryKey"`

	// 省份、城市简称
	CityKey *string `json:"CityKey,omitnil,omitempty" name:"CityKey"`

	// 规则创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 规则最近更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 域名数
	DnsParseCnt *int64 `json:"DnsParseCnt,omitnil,omitempty" name:"DnsParseCnt"`
}

// Predefined struct for user
type IgnorePolicyRiskRequestParams struct {
	// 策略问题ID
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// 成员Id
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

type IgnorePolicyRiskRequest struct {
	*tchttp.BaseRequest
	
	// 策略问题ID
	RiskId *string `json:"RiskId,omitnil,omitempty" name:"RiskId"`

	// 成员Id
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

func (r *IgnorePolicyRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnorePolicyRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RiskId")
	delete(f, "MemberId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IgnorePolicyRiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IgnorePolicyRiskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IgnorePolicyRiskResponse struct {
	*tchttp.BaseResponse
	Response *IgnorePolicyRiskResponseParams `json:"Response"`
}

func (r *IgnorePolicyRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnorePolicyRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MemberInfo struct {
	// 成员AppId
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 成员Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 成员昵称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 成员Id
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`
}

// Predefined struct for user
type ModifyEdgeAclRuleRequestParams struct {
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 要修改的规则，必须包含RuleId
	Rule *EdgeAclRuleInfo `json:"Rule,omitnil,omitempty" name:"Rule"`
}

type ModifyEdgeAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 要修改的规则，必须包含RuleId
	Rule *EdgeAclRuleInfo `json:"Rule,omitnil,omitempty" name:"Rule"`
}

func (r *ModifyEdgeAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEdgeAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeAclRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEdgeAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEdgeAclRuleResponseParams `json:"Response"`
}

func (r *ModifyEdgeAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeAclRuleSequenceRequestParams struct {
	// 出入站方向 0=出向，1=入向
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则序号调整列表，必须包含所有受影响的规则
	Sequences []*SequenceIndex `json:"Sequences,omitnil,omitempty" name:"Sequences"`
}

type ModifyEdgeAclRuleSequenceRequest struct {
	*tchttp.BaseRequest
	
	// 出入站方向 0=出向，1=入向
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则序号调整列表，必须包含所有受影响的规则
	Sequences []*SequenceIndex `json:"Sequences,omitnil,omitempty" name:"Sequences"`
}

func (r *ModifyEdgeAclRuleSequenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeAclRuleSequenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "GroupId")
	delete(f, "Sequences")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEdgeAclRuleSequenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeAclRuleSequenceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEdgeAclRuleSequenceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEdgeAclRuleSequenceResponseParams `json:"Response"`
}

func (r *ModifyEdgeAclRuleSequenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeAclRuleSequenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatAclRuleRequestParams struct {
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则
	Rule *NatAclRule `json:"Rule,omitnil,omitempty" name:"Rule"`
}

type ModifyNatAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则
	Rule *NatAclRule `json:"Rule,omitnil,omitempty" name:"Rule"`
}

func (r *ModifyNatAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatAclRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNatAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNatAclRuleResponseParams `json:"Response"`
}

func (r *ModifyNatAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatAclRuleSequenceRequestParams struct {
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 序号调整列表
	Sequences []*SequenceIndex `json:"Sequences,omitnil,omitempty" name:"Sequences"`

	// 规则方向：1-入站规则，0-出站规则
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type ModifyNatAclRuleSequenceRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 序号调整列表
	Sequences []*SequenceIndex `json:"Sequences,omitnil,omitempty" name:"Sequences"`

	// 规则方向：1-入站规则，0-出站规则
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

func (r *ModifyNatAclRuleSequenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatAclRuleSequenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Sequences")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatAclRuleSequenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatAclRuleSequenceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNatAclRuleSequenceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNatAclRuleSequenceResponseParams `json:"Response"`
}

func (r *ModifyNatAclRuleSequenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatAclRuleSequenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleGroupRequestParams struct {
	// 规则组名称
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
}

type ModifyRuleGroupRequest struct {
	*tchttp.BaseRequest
	
	// 规则组名称
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`
}

func (r *ModifyRuleGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRuleGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRuleGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRuleGroupResponseParams `json:"Response"`
}

func (r *ModifyRuleGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupRuleRequestParams struct {
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则
	Rule *SecurityGroupRule `json:"Rule,omitnil,omitempty" name:"Rule"`
}

type ModifySecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则
	Rule *SecurityGroupRule `json:"Rule,omitnil,omitempty" name:"Rule"`
}

func (r *ModifySecurityGroupRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityGroupRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySecurityGroupRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityGroupRuleResponseParams `json:"Response"`
}

func (r *ModifySecurityGroupRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStrategyRequestParams struct {
	// 策略Id
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 下发规则接收账号
	ReceiveAccount []*string `json:"ReceiveAccount,omitnil,omitempty" name:"ReceiveAccount"`

	// 优先级
	Sequence *int64 `json:"Sequence,omitnil,omitempty" name:"Sequence"`

	// 规则组Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 下发规则接收账号组
	ReceiveGroup []*string `json:"ReceiveGroup,omitnil,omitempty" name:"ReceiveGroup"`
}

type ModifyStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 策略Id
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 下发规则接收账号
	ReceiveAccount []*string `json:"ReceiveAccount,omitnil,omitempty" name:"ReceiveAccount"`

	// 优先级
	Sequence *int64 `json:"Sequence,omitnil,omitempty" name:"Sequence"`

	// 规则组Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 下发规则接收账号组
	ReceiveGroup []*string `json:"ReceiveGroup,omitnil,omitempty" name:"ReceiveGroup"`
}

func (r *ModifyStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StrategyId")
	delete(f, "ReceiveAccount")
	delete(f, "Sequence")
	delete(f, "GroupId")
	delete(f, "ReceiveGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStrategyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStrategyResponseParams `json:"Response"`
}

func (r *ModifyStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStrategySequenceRequestParams struct {
	// 优先级列表
	Sequences []*SequenceIndex `json:"Sequences,omitnil,omitempty" name:"Sequences"`

	// 执行区域
	ExecArea *string `json:"ExecArea,omitnil,omitempty" name:"ExecArea"`

	// 产品类型
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type ModifyStrategySequenceRequest struct {
	*tchttp.BaseRequest
	
	// 优先级列表
	Sequences []*SequenceIndex `json:"Sequences,omitnil,omitempty" name:"Sequences"`

	// 执行区域
	ExecArea *string `json:"ExecArea,omitnil,omitempty" name:"ExecArea"`

	// 产品类型
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *ModifyStrategySequenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStrategySequenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Sequences")
	delete(f, "ExecArea")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStrategySequenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStrategySequenceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyStrategySequenceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStrategySequenceResponseParams `json:"Response"`
}

func (r *ModifyStrategySequenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStrategySequenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcAclRuleRequestParams struct {
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则
	Rule *VpcAclRule `json:"Rule,omitnil,omitempty" name:"Rule"`
}

type ModifyVpcAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则
	Rule *VpcAclRule `json:"Rule,omitnil,omitempty" name:"Rule"`
}

func (r *ModifyVpcAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpcAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcAclRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVpcAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpcAclRuleResponseParams `json:"Response"`
}

func (r *ModifyVpcAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcAclRuleSequenceRequestParams struct {
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 序号调整列表
	Sequences []*SequenceIndex `json:"Sequences,omitnil,omitempty" name:"Sequences"`
}

type ModifyVpcAclRuleSequenceRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 序号调整列表
	Sequences []*SequenceIndex `json:"Sequences,omitnil,omitempty" name:"Sequences"`
}

func (r *ModifyVpcAclRuleSequenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcAclRuleSequenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Sequences")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpcAclRuleSequenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcAclRuleSequenceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVpcAclRuleSequenceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpcAclRuleSequenceResponseParams `json:"Response"`
}

func (r *ModifyVpcAclRuleSequenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcAclRuleSequenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NatAclRule struct {
	// <p>源地址内容</p>
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// <p>源类型：ip/url/template/instance/tag</p>
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// <p>目的地址内容</p>
	TargetContent *string `json:"TargetContent,omitnil,omitempty" name:"TargetContent"`

	// <p>目的类型：ip/url/template/instance/tag</p>
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// <p>协议：TCP/UDP/ICMP/ANY/HTTP/HTTPS/DNS/FTP等</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>动作：accept/drop/log</p>
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// <p>优先级（从1开始）</p>
	OrderIndex *int64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// <p>规则生效范围：ALL-全局生效，ap-xxx-地域生效，cfwnat-xxx-NAT防火墙实例生效</p>
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>规则方向：1-入站规则，0-出站规则</p>
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// <p>规则ID（修改时必填）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// <p>端口（ICMP协议时为空）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>规则描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>端口模板ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamTemplateId *string `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// <p>规则归属的成员账号ID(当Scope为cfwnat-xxx或SourceType/DestType为instance/tag时必填)</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BelongMemberId *string `json:"BelongMemberId,omitnil,omitempty" name:"BelongMemberId"`
}

type NatAclRuleResp struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 优先级
	Sequence *int64 `json:"Sequence,omitnil,omitempty" name:"Sequence"`

	// 规则方向：0-出向，1-入向
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 源地址内容
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// 源类型
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 源资产名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// 目的地址内容
	TargetContent *string `json:"TargetContent,omitnil,omitempty" name:"TargetContent"`

	// 目的类型
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 目的资产名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 动作：accept/drop/log
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// 规则描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 规则生效范围
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 规则生效范围描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScopeDesc *string `json:"ScopeDesc,omitnil,omitempty" name:"ScopeDesc"`

	// 防火墙实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwInsId *string `json:"FwInsId,omitnil,omitempty" name:"FwInsId"`

	// 国家名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountryName *string `json:"CountryName,omitnil,omitempty" name:"CountryName"`

	// 城市名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CityName *string `json:"CityName,omitnil,omitempty" name:"CityName"`

	// 国家代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountryCode *int64 `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 城市代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	CityCode *int64 `json:"CityCode,omitnil,omitempty" name:"CityCode"`

	// 国家键值
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountryKey *string `json:"CountryKey,omitnil,omitempty" name:"CountryKey"`

	// 城市键值
	// 注意：此字段可能返回 null，表示取不到有效值。
	CityKey *string `json:"CityKey,omitnil,omitempty" name:"CityKey"`

	// 是否地域规则：0-否，1-是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsRegion *int64 `json:"IsRegion,omitnil,omitempty" name:"IsRegion"`

	// 云厂商代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloudCode *string `json:"CloudCode,omitnil,omitempty" name:"CloudCode"`

	// 是否云厂商规则：0-否，1-是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCloud *int64 `json:"IsCloud,omitnil,omitempty" name:"IsCloud"`

	// 端口模板ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamTemplateId *string `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// 端口模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamTemplateName *string `json:"ParamTemplateName,omitnil,omitempty" name:"ParamTemplateName"`

	// 规则是否失效：0-有效，1-失效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Invalid *int64 `json:"Invalid,omitnil,omitempty" name:"Invalid"`

	// 规则归属的成员账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BelongMember *MemberInfo `json:"BelongMember,omitnil,omitempty" name:"BelongMember"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 域名数
	DnsParseCnt *int64 `json:"DnsParseCnt,omitnil,omitempty" name:"DnsParseCnt"`
}

type OrganMemberItem struct {
	// 成员 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberId *string `json:"MemberId,omitnil,omitempty" name:"MemberId"`

	// 成员账号 AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 账号Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 账号名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 子账号数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountCount *int64 `json:"SubAccountCount,omitnil,omitempty" name:"SubAccountCount"`

	// 所属组织架构节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 成员身份：admin-管理员，delegatedAdmin-委派管理员，member-普通成员
	// 注意：此字段可能返回 null，表示取不到有效值。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 成员身份显示名称（前端展示用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleDisplay *string `json:"RoleDisplay,omitnil,omitempty" name:"RoleDisplay"`

	// 所属账户组 
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountGroup *AccountGroupInfo `json:"AccountGroup,omitnil,omitempty" name:"AccountGroup"`

	// 云防火墙纳管状态：0-未纳管，1-已纳管
	// 注意：此字段可能返回 null，表示取不到有效值。
	CfwManaged *int64 `json:"CfwManaged,omitnil,omitempty" name:"CfwManaged"`

	// 云防火墙共享角色：sharer-共享者，user-使用者，none-未设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	CfwShareRole *string `json:"CfwShareRole,omitnil,omitempty" name:"CfwShareRole"`

	// 云防火墙共享角色显示名称（前端展示用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CfwShareRoleDisplay *string `json:"CfwShareRoleDisplay,omitnil,omitempty" name:"CfwShareRoleDisplay"`

	// 云防火墙共享者 AppId，成员角色为使用者时有值
	// 注意：此字段可能返回 null，表示取不到有效值。
	CfwSharerAppId *string `json:"CfwSharerAppId,omitnil,omitempty" name:"CfwSharerAppId"`

	// 云防火墙计费实例 ID，非空表示已购买云防火墙
	// 注意：此字段可能返回 null，表示取不到有效值。
	CfwInstanceId *string `json:"CfwInstanceId,omitnil,omitempty" name:"CfwInstanceId"`

	// 策略分析权限：0-关闭，1-开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyAnalysisEnabled *int64 `json:"PolicyAnalysisEnabled,omitnil,omitempty" name:"PolicyAnalysisEnabled"`

	// 成员加入集团时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	MemberCreateTime *string `json:"MemberCreateTime,omitnil,omitempty" name:"MemberCreateTime"`

	// 账号加入方式
	JoinType *string `json:"JoinType,omitnil,omitempty" name:"JoinType"`
}

type OrganSummary struct {
	// 集团名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 管理员账号信息
	AdminInfo *MemberInfo `json:"AdminInfo,omitnil,omitempty" name:"AdminInfo"`

	// 管理员/委派管理员数量
	AdminCount *int64 `json:"AdminCount,omitnil,omitempty" name:"AdminCount"`

	// 已接入成员数
	JoinedMemberCount *int64 `json:"JoinedMemberCount,omitnil,omitempty" name:"JoinedMemberCount"`

	// 接入成员上限（-1表示无上限）
	MemberLimit *int64 `json:"MemberLimit,omitnil,omitempty" name:"MemberLimit"`

	// 接入成员上限显示
	MemberLimitDisplay *string `json:"MemberLimitDisplay,omitnil,omitempty" name:"MemberLimitDisplay"`

	// 规格共享者数量
	CfwSharerCount *int64 `json:"CfwSharerCount,omitnil,omitempty" name:"CfwSharerCount"`

	// 规格使用者数量
	CfwUserCount *int64 `json:"CfwUserCount,omitnil,omitempty" name:"CfwUserCount"`

	// 部门名称列表
	Departments []*string `json:"Departments,omitnil,omitempty" name:"Departments"`

	// 纳管账号数
	ManagedMemberCount *int64 `json:"ManagedMemberCount,omitnil,omitempty" name:"ManagedMemberCount"`

	// 纳管产品数
	ManagedProductCount *int64 `json:"ManagedProductCount,omitnil,omitempty" name:"ManagedProductCount"`

	// 纳管账号数
	CfwManageCount *int64 `json:"CfwManageCount,omitnil,omitempty" name:"CfwManageCount"`
}

type PolicyRisk struct {
	// 问题对应的唯一uuid
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 风险大类
	RiskCategory *string `json:"RiskCategory,omitnil,omitempty" name:"RiskCategory"`

	// 风险子类
	RiskSubCategory *string `json:"RiskSubCategory,omitnil,omitempty" name:"RiskSubCategory"`

	// 规则分类
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 风险等级，0：低风险
	// 1：中风险
	// 2：高风险
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 安全组
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 风险包含的企业安全组规则ID
	SgRuleId []*string `json:"SgRuleId,omitnil,omitempty" name:"SgRuleId"`

	// 风险包含安全组ID内的问题规则数
	RuleCount *int64 `json:"RuleCount,omitnil,omitempty" name:"RuleCount"`

	// 风险包含的安全组ID
	SgId []*string `json:"SgId,omitnil,omitempty" name:"SgId"`

	// 风险特征
	RiskFeature *string `json:"RiskFeature,omitnil,omitempty" name:"RiskFeature"`

	// 处置建议
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 处置状态，0：未处理，1：已处理，2：忽略
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 发现时间
	FoundTime *string `json:"FoundTime,omitnil,omitempty" name:"FoundTime"`

	// 处置时间
	DisposalTime *string `json:"DisposalTime,omitnil,omitempty" name:"DisposalTime"`

	// 安全组地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Ingress入站，Egress出站
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 风险原因
	RiskReason *string `json:"RiskReason,omitnil,omitempty" name:"RiskReason"`
}

type ReceiveAccount struct {
	// 租户 uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 租户名称
	Nickname *string `json:"Nickname,omitnil,omitempty" name:"Nickname"`

	// 0=账号uin，1=账号组
	ReceiverType *int64 `json:"ReceiverType,omitnil,omitempty" name:"ReceiverType"`

	// 只有ReceiverType 是 1 时 才返回账号列表
	Members []*MemberInfo `json:"Members,omitnil,omitempty" name:"Members"`
}

type RiskCategoryItem struct {
	// 风险大类ID
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// 风险大类名称
	CategoryName *string `json:"CategoryName,omitnil,omitempty" name:"CategoryName"`

	// 风险子类ID
	SubcategoryId *string `json:"SubcategoryId,omitnil,omitempty" name:"SubcategoryId"`

	// 风险子类名称
	SubcategoryName *string `json:"SubcategoryName,omitnil,omitempty" name:"SubcategoryName"`

	// 风险等级(0-低危,1-中危,2-高危)
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 风险等级名称(低/中/高)
	RiskLevelName *string `json:"RiskLevelName,omitnil,omitempty" name:"RiskLevelName"`

	// 风险描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 处置建议
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 该类风险的规则数量
	RuleCount *int64 `json:"RuleCount,omitnil,omitempty" name:"RuleCount"`

	// 已处置数量
	TreatedCount *int64 `json:"TreatedCount,omitnil,omitempty" name:"TreatedCount"`

	// 已忽略数量
	IgnoredCount *int64 `json:"IgnoredCount,omitnil,omitempty" name:"IgnoredCount"`

	// 待整改数量
	UntreatedCount *int64 `json:"UntreatedCount,omitnil,omitempty" name:"UntreatedCount"`

	// 整改率(百分比字符串)
	DisposalRate *uint64 `json:"DisposalRate,omitnil,omitempty" name:"DisposalRate"`

	// 是否有未处理风险
	// -1: 未体检
	// 0: 无风险
	// 1: 有风险
	HasRisk *int64 `json:"HasRisk,omitnil,omitempty" name:"HasRisk"`

	// 整改状态：
	// Completed： 已整改完成（整改率 100%）
	// Incomplete： 未整改完成（整改率 < 100%）
	// -：未体检/无数据
	RemediationStatus *string `json:"RemediationStatus,omitnil,omitempty" name:"RemediationStatus"`
}

type SecGroupRuleResp struct {
	// <p>排序</p>
	OrderIndex *int64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// <p>主键id</p>
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// <p>ip类型</p>
	IpVersion *string `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`

	// <p>源规则内容</p>
	SourceId *string `json:"SourceId,omitnil,omitempty" name:"SourceId"`

	// <p>源规则类型<br>取值范围 0/1/2/3/4/5/6/7/8/9/100<br>0表示ip(net),<br>1表示VPC实例(instance)<br>2表示子网实例(instance)<br>3表示CVM实例(instance)<br>4表示CLB实例(instance)<br>5表示ENI实例(instance)<br>6表示数据库实例(instance)<br>7表示模板(template)<br>8表示标签(tag)<br>9表示地域(region)<br>100表示资产分组(resourcegroup)</p>
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// <p>目的规则内容</p>
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// <p>目的规则类型<br>取值范围 0/1/2/3/4/5/6/7/8/9/100<br>0表示ip(net),<br>1表示VPC实例(instance)<br>2表示子网实例(instance)<br>3表示CVM实例(instance)<br>4表示CLB实例(instance)<br>5表示ENI实例(instance)<br>6表示数据库实例(instance)<br>7表示模板(template)<br>8表示标签(tag)<br>9表示地域(region)<br>100表示资产分组(resourcegroup)</p>
	TargetType *int64 `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// <p>协议名称<br>取值范围:TCP/ANY/ICMP/UDP<br>ANY:表示所有</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>端口</p>
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>策略</p>
	Strategy *int64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// <p>描述</p>
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// <p>地域</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>服务模板id</p>
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`

	// <p>源资产名称</p>
	SouInstanceName *string `json:"SouInstanceName,omitnil,omitempty" name:"SouInstanceName"`

	// <p>源资产公网ip</p>
	SouPublicIp *string `json:"SouPublicIp,omitnil,omitempty" name:"SouPublicIp"`

	// <p>源资产内网ip</p>
	SouPrivateIp *string `json:"SouPrivateIp,omitnil,omitempty" name:"SouPrivateIp"`

	// <p>源资产网段信息</p>
	SouCidr *string `json:"SouCidr,omitnil,omitempty" name:"SouCidr"`

	// <p>源模板名称</p>
	SouParameterName *string `json:"SouParameterName,omitnil,omitempty" name:"SouParameterName"`

	// <p>目的资产名称</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>目的资产公网ip</p>
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// <p>目的资产内网ip</p>
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// <p>目的资产网段信息</p>
	Cidr *string `json:"Cidr,omitnil,omitempty" name:"Cidr"`

	// <p>目的模板名称</p>
	ParameterName *string `json:"ParameterName,omitnil,omitempty" name:"ParameterName"`

	// <p>端口模板名称</p>
	ProtocolPortName *string `json:"ProtocolPortName,omitnil,omitempty" name:"ProtocolPortName"`

	// <p>规则id  等同RuleUuid</p>
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>域名解析的IP统计</p>
	DnsParseCount *SgDnsParseCount `json:"DnsParseCount,omitnil,omitempty" name:"DnsParseCount"`

	// <p>规则生效范围</p>
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>规则最近一次是否有改动 取值范围：0/1 0:否 1:是</p>
	IsNew *int64 `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// <p>规则归属的成员账号（当FwGroupId为cfwg-xxx或SourceType/DestType为instance/tag时必填)</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BelongMember *MemberInfo `json:"BelongMember,omitnil,omitempty" name:"BelongMember"`
}

type SecurityGroupRiskPolicy struct {
	// 安全组规则索引号
	PolicyIndex *int64 `json:"PolicyIndex,omitnil,omitempty" name:"PolicyIndex"`

	//  协议, 取值: TCP,UDP,ICMP,ICMPv6,ALL。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 端口(all, 离散port,  range)。
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 端口ID或者协议端口组ID。ServiceTemplate和Protocol+Port互斥。
	ServiceTemplate *ServiceTemplateSpecification `json:"ServiceTemplate,omitnil,omitempty" name:"ServiceTemplate"`

	// 网段或IP(互斥)。
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// 网段或IPv6(互斥)。
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitnil,omitempty" name:"Ipv6CidrBlock"`

	// 安全组实例ID，例如：sg-ohuuioma。
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// IP地址ID或者ID地址组ID。
	AddressTemplate *AddressTemplateSpecification `json:"AddressTemplate,omitnil,omitempty" name:"AddressTemplate"`

	// 动作：ACCEPT 或 DROP。
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 安全组规则描述。
	PolicyDescription *string `json:"PolicyDescription,omitnil,omitempty" name:"PolicyDescription"`

	// 安全组规则当前版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 规则方向，Egress出站规则，Ingress入站规则
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 安全组最近修改时间。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 安全组所在地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type SecurityGroupRule struct {
	// ip类型
	IpVersion *string `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`

	// 源地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// 源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 目的地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestContent *string `json:"DestContent,omitnil,omitempty" name:"DestContent"`

	// 目的类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestType *string `json:"DestType,omitnil,omitempty" name:"DestType"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`

	// 动作
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 优先级
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderIndex *int64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// rule id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 生效范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 端口类型
	ProtocolPortType *int64 `json:"ProtocolPortType,omitnil,omitempty" name:"ProtocolPortType"`

	// 规则归属的成员账号ID（当FwGroupId为cfwg-xxx或SourceType/DestType为instance/tag时必填）
	BelongMemberId *string `json:"BelongMemberId,omitnil,omitempty" name:"BelongMemberId"`
}

type SequenceIndex struct {
	// 原规则序号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderIndex *int64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// 新规则序号
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewOrderIndex *int64 `json:"NewOrderIndex,omitnil,omitempty" name:"NewOrderIndex"`
}

type ServiceTemplateSpecification struct {
	//  协议端口ID，例如：ppm-f5n1f8da。
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 协议端口组ID，例如：ppmg-f5n1f8da
	ServiceGroupId *string `json:"ServiceGroupId,omitnil,omitempty" name:"ServiceGroupId"`
}

type SgDnsParseCount struct {
	// 有效下发的IP个数，离散数据
	ValidCount *int64 `json:"ValidCount,omitnil,omitempty" name:"ValidCount"`

	// 未下发的IP个数，离散数据
	InvalidCount *int64 `json:"InvalidCount,omitnil,omitempty" name:"InvalidCount"`
}

type SgRuleResp struct {
	// 规则Id
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 优先级
	Sequence *int64 `json:"Sequence,omitnil,omitempty" name:"Sequence"`

	// 区域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// ip类型
	IpVersion *string `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`

	// 源内容
	SrcContent *string `json:"SrcContent,omitnil,omitempty" name:"SrcContent"`

	// 源类型
	SrcType *string `json:"SrcType,omitnil,omitempty" name:"SrcType"`

	// 目的内容
	DstContent *string `json:"DstContent,omitnil,omitempty" name:"DstContent"`

	// 目的类型
	DstType *string `json:"DstType,omitnil,omitempty" name:"DstType"`

	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 协议端口参数模板
	ProtocolPortType *int64 `json:"ProtocolPortType,omitnil,omitempty" name:"ProtocolPortType"`

	// 协议端口参数模板id
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`

	// 端口段,支持单端口,多端口和端口段
	DstPort *string `json:"DstPort,omitnil,omitempty" name:"DstPort"`

	// 策略，1阻断，2放行
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// 描述
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 规则来源，0为用户控制台添加
	RuleSource *string `json:"RuleSource,omitnil,omitempty" name:"RuleSource"`

	// 生效范围,字节位,1:SG 企业安全组,2:LH 轻量服务器
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`
}

type StrategyReq struct {
	// 规则组Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 优先级
	Sequence *int64 `json:"Sequence,omitnil,omitempty" name:"Sequence"`
}

type StrategyResp struct {
	// 策略Id
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 规则组Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 规则组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 规则数
	RuleCount *int64 `json:"RuleCount,omitnil,omitempty" name:"RuleCount"`

	// 策略状态
	RuleStatus *int64 `json:"RuleStatus,omitnil,omitempty" name:"RuleStatus"`

	// 下发账号
	ReceiveAccount []*ReceiveAccount `json:"ReceiveAccount,omitnil,omitempty" name:"ReceiveAccount"`

	// 优先级
	Sequence *int64 `json:"Sequence,omitnil,omitempty" name:"Sequence"`

	// 下发失败原因
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// 下发失败原因类型
	ErrorType *string `json:"ErrorType,omitnil,omitempty" name:"ErrorType"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 创建人
	CreateBy *string `json:"CreateBy,omitnil,omitempty" name:"CreateBy"`

	// 更新人
	UpdateBy *string `json:"UpdateBy,omitnil,omitempty" name:"UpdateBy"`

	// 执行区域
	ExecArea *string `json:"ExecArea,omitnil,omitempty" name:"ExecArea"`

	// 创建人名称
	CreateName *string `json:"CreateName,omitnil,omitempty" name:"CreateName"`

	// 更新人名称
	UpdateName *string `json:"UpdateName,omitnil,omitempty" name:"UpdateName"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type VpcAclRule struct {
	// 源地址内容
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// 源类型：ip/url/template/instance/tag
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 目的地址内容
	DestContent *string `json:"DestContent,omitnil,omitempty" name:"DestContent"`

	// 目的类型：ip/url/template/instance/tag
	DestType *string `json:"DestType,omitnil,omitempty" name:"DestType"`

	// 协议：TCP/UDP/ICMP/ANY/HTTP/HTTPS/DNS/FTP等
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 动作：accept/drop/log
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// 优先级（从1开始）
	OrderIndex *int64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// 边界防火墙ID：ALL表示全局，CFWS-xxx表示指定边界
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 防火墙实例ID（规则生效范围）：ALL-全局生效，ccn-xxx-云联网实例，cfwg-xxx-防火墙组实例
	FwGroupId *string `json:"FwGroupId,omitnil,omitempty" name:"FwGroupId"`

	// 规则ID（修改时必填）
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 端口（ICMP协议时为空）
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 端口模板ID
	ParamTemplateId *string `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// 规则归属的成员账号ID（当FwGroupId为cfwg-xxx或SourceType/DestType为instance/tag时必填）
	BelongMemberId *string `json:"BelongMemberId,omitnil,omitempty" name:"BelongMemberId"`
}

type VpcAclRuleResp struct {
	// 规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 优先级
	Sequence *int64 `json:"Sequence,omitnil,omitempty" name:"Sequence"`

	// IP版本：ipv4或ipv6
	IpVersion *string `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`

	// 源地址内容
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// 源类型
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 源资产名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// 目的地址内容
	DestContent *string `json:"DestContent,omitnil,omitempty" name:"DestContent"`

	// 目的类型
	DestType *string `json:"DestType,omitnil,omitempty" name:"DestType"`

	// 目的资产名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestName *string `json:"DestName,omitnil,omitempty" name:"DestName"`

	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 端口模板ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamTemplateId *string `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// 端口模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamTemplateName *string `json:"ParamTemplateName,omitnil,omitempty" name:"ParamTemplateName"`

	// 动作：accept/drop/log
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// 规则描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 边界防火墙ID
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 防火墙实例ID
	FwGroupId *string `json:"FwGroupId,omitnil,omitempty" name:"FwGroupId"`

	// 规则是否失效：0-有效，1-失效
	Invalid *int64 `json:"Invalid,omitnil,omitempty" name:"Invalid"`

	// 规则归属的成员账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BelongMember *MemberInfo `json:"BelongMember,omitnil,omitempty" name:"BelongMember"`

	// 规则创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 规则修改时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 域名数
	DnsParseCnt *int64 `json:"DnsParseCnt,omitnil,omitempty" name:"DnsParseCnt"`

	// 防火墙组名称
	FwGroupName *string `json:"FwGroupName,omitnil,omitempty" name:"FwGroupName"`
}