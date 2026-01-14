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

package v20190904

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AcListsData struct {
	// 规则id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 访问源
	SourceIp *string `json:"SourceIp,omitnil,omitempty" name:"SourceIp"`

	// 访问目的
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 策略
	Strategy *uint64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// 描述
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 命中次数
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// 告警规则id
	LogId *string `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 规则开关状态 1打开 0关闭
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 规则源类型
	SrcType *int64 `json:"SrcType,omitnil,omitempty" name:"SrcType"`

	// 规则目的类型
	DstType *int64 `json:"DstType,omitnil,omitempty" name:"DstType"`

	// 规则唯一ID
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// 规则有效性
	// 1 有效
	// 0 无效
	Invalid *int64 `json:"Invalid,omitnil,omitempty" name:"Invalid"`

	// 是否地域规则
	IsRegion *int64 `json:"IsRegion,omitnil,omitempty" name:"IsRegion"`

	// 云厂商代码
	CloudCode *string `json:"CloudCode,omitnil,omitempty" name:"CloudCode"`

	// 自动化助手信息
	AutoTask *string `json:"AutoTask,omitnil,omitempty" name:"AutoTask"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 地域码信息
	RegionCode *string `json:"RegionCode,omitnil,omitempty" name:"RegionCode"`

	// 国家代码
	Country *int64 `json:"Country,omitnil,omitempty" name:"Country"`

	// 城市代码
	City *int64 `json:"City,omitnil,omitempty" name:"City"`

	// 国家名称
	RegName1 *string `json:"RegName1,omitnil,omitempty" name:"RegName1"`

	// 城市名称
	RegName2 *string `json:"RegName2,omitnil,omitempty" name:"RegName2"`
}

type AccessInstanceInfo struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例类型VPC or DIRECTCONNECT等类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例所在地域
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`

	// 接入防火墙的网段模式：0-不接入，1-接入实例关联的所有网段，2-接入用户自定义的网段
	AccessCidrMode *int64 `json:"AccessCidrMode,omitnil,omitempty" name:"AccessCidrMode"`

	// 接入防火墙的网段列表
	AccessCidrList []*string `json:"AccessCidrList,omitnil,omitempty" name:"AccessCidrList"`
}

// Predefined struct for user
type AddAclRuleRequestParams struct {
	// 需要添加的访问控制规则列表
	Rules []*CreateRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 添加规则的来源，一般不需要使用，值insert_rule 表示插入指定位置的规则；值batch_import 表示批量导入规则；为空时表示添加规则
	From *string `json:"From,omitnil,omitempty" name:"From"`
}

type AddAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要添加的访问控制规则列表
	Rules []*CreateRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 添加规则的来源，一般不需要使用，值insert_rule 表示插入指定位置的规则；值batch_import 表示批量导入规则；为空时表示添加规则
	From *string `json:"From,omitnil,omitempty" name:"From"`
}

func (r *AddAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	delete(f, "From")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAclRuleResponseParams struct {
	// 创建成功后返回新策略ID列表
	RuleUuid []*int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddAclRuleResponseParams `json:"Response"`
}

func (r *AddAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEnterpriseSecurityGroupRulesRequestParams struct {
	// 创建规则数据
	Data []*SecurityGroupRule `json:"Data,omitnil,omitempty" name:"Data"`

	// 添加类型，0：添加到最后，1：添加到最前；2：中间插入；默认0添加到最后
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符，且不能超过64个字符。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// （IsDelay为老版参数，新版无需输入）是否延迟下发，1则延迟下发，否则立即下发
	IsDelay *uint64 `json:"IsDelay,omitnil,omitempty" name:"IsDelay"`

	// 来源 默认空 覆盖导入是 batch_import_cover
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 是否复用rule id，1为是，默认不需要
	IsUseId *int64 `json:"IsUseId,omitnil,omitempty" name:"IsUseId"`
}

type AddEnterpriseSecurityGroupRulesRequest struct {
	*tchttp.BaseRequest
	
	// 创建规则数据
	Data []*SecurityGroupRule `json:"Data,omitnil,omitempty" name:"Data"`

	// 添加类型，0：添加到最后，1：添加到最前；2：中间插入；默认0添加到最后
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符，且不能超过64个字符。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// （IsDelay为老版参数，新版无需输入）是否延迟下发，1则延迟下发，否则立即下发
	IsDelay *uint64 `json:"IsDelay,omitnil,omitempty" name:"IsDelay"`

	// 来源 默认空 覆盖导入是 batch_import_cover
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 是否复用rule id，1为是，默认不需要
	IsUseId *int64 `json:"IsUseId,omitnil,omitempty" name:"IsUseId"`
}

func (r *AddEnterpriseSecurityGroupRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEnterpriseSecurityGroupRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	delete(f, "Type")
	delete(f, "ClientToken")
	delete(f, "IsDelay")
	delete(f, "From")
	delete(f, "IsUseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddEnterpriseSecurityGroupRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEnterpriseSecurityGroupRulesResponseParams struct {
	// 状态值，0：添加成功，非0：添加失败
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 添加成功的规则详情
	Rules []*SecurityGroupSimplifyRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddEnterpriseSecurityGroupRulesResponse struct {
	*tchttp.BaseResponse
	Response *AddEnterpriseSecurityGroupRulesResponseParams `json:"Response"`
}

func (r *AddEnterpriseSecurityGroupRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddEnterpriseSecurityGroupRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNatAcRuleRequestParams struct {
	// 需要添加的nat访问控制规则列表
	Rules []*CreateNatRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 添加规则的来源，一般不需要使用，值insert_rule 表示插入指定位置的规则；值batch_import 表示批量导入规则；为空时表示添加规则
	From *string `json:"From,omitnil,omitempty" name:"From"`
}

type AddNatAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要添加的nat访问控制规则列表
	Rules []*CreateNatRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 添加规则的来源，一般不需要使用，值insert_rule 表示插入指定位置的规则；值batch_import 表示批量导入规则；为空时表示添加规则
	From *string `json:"From,omitnil,omitempty" name:"From"`
}

func (r *AddNatAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNatAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	delete(f, "From")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddNatAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNatAcRuleResponseParams struct {
	// 创建成功后返回新策略ID列表
	RuleUuid []*int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddNatAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddNatAcRuleResponseParams `json:"Response"`
}

func (r *AddNatAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNatAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddVpcAcRuleRequestParams struct {
	// 需要添加的vpc内网间规则列表
	Rules []*VpcRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 添加规则的来源，一般不需要使用，值insert_rule 表示插入指定位置的规则；值batch_import 表示批量导入规则；为空时表示添加规则
	From *string `json:"From,omitnil,omitempty" name:"From"`
}

type AddVpcAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要添加的vpc内网间规则列表
	Rules []*VpcRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 添加规则的来源，一般不需要使用，值insert_rule 表示插入指定位置的规则；值batch_import 表示批量导入规则；为空时表示添加规则
	From *string `json:"From,omitnil,omitempty" name:"From"`
}

func (r *AddVpcAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddVpcAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	delete(f, "From")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddVpcAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddVpcAcRuleResponseParams struct {
	// 创建成功后返回新策略ID列表
	RuleUuids []*int64 `json:"RuleUuids,omitnil,omitempty" name:"RuleUuids"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddVpcAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddVpcAcRuleResponseParams `json:"Response"`
}

func (r *AddVpcAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddVpcAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssetZone struct {
	// 地域
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 地域英文
	ZoneEng *string `json:"ZoneEng,omitnil,omitempty" name:"ZoneEng"`
}

type AssociatedInstanceInfo struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例类型，3是cvm实例,4是clb实例,5是eni实例,6是云数据库
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 私有网络名称
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 公网IP
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 内网IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 关联安全组数量
	SecurityGroupCount *uint64 `json:"SecurityGroupCount,omitnil,omitempty" name:"SecurityGroupCount"`

	// 关联安全组规则数量
	SecurityGroupRuleCount *uint64 `json:"SecurityGroupRuleCount,omitnil,omitempty" name:"SecurityGroupRuleCount"`

	// 关联数据库代理Id
	CdbId *string `json:"CdbId,omitnil,omitempty" name:"CdbId"`
}

type AttachInsInfo struct {
	// 实例对象可以是cvm类型:ins-ad21xuds1形式;路由表类型:rtb-da12daxd形式;vpc类型:vpc-1dxdad2d形式
	InsId *string `json:"InsId,omitnil,omitempty" name:"InsId"`

	// 实例对象名称
	InsName *string `json:"InsName,omitnil,omitempty" name:"InsName"`

	// 实例的cidr
	Cidr *string `json:"Cidr,omitnil,omitempty" name:"Cidr"`
}

type BanAndAllowRule struct {
	// 规则评论
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 自定义白名单规则
	CustomRule *CustomWhiteRule `json:"CustomRule,omitnil,omitempty" name:"CustomRule"`

	// 0互联网出站 1互联网入站 5内网访问源 6内网访问目的
	DirectionList *string `json:"DirectionList,omitnil,omitempty" name:"DirectionList"`

	// 规则截止时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 放通的引擎: 1针对互联网边界 2针对nat防火墙 4针对vpc防火墙
	FwType *int64 `json:"FwType,omitnil,omitempty" name:"FwType"`

	// 封禁和放通对象
	Ioc *string `json:"Ioc,omitnil,omitempty" name:"Ioc"`
}

type BanAndAllowRuleDel struct {
	// 封禁和放通对象
	Ioc *string `json:"Ioc,omitnil,omitempty" name:"Ioc"`

	// 0互联网出站 1互联网入站 5内网访问源 6内网访问目的 （DeleteBlockIgnoreRuleNew接口，该字段无效）
	DirectionList *string `json:"DirectionList,omitnil,omitempty" name:"DirectionList"`

	// 规则类型
	// RuleType: 1黑名单 2外部IP 3域名 4情报 5资产 6自定义规则  7入侵防御规则
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`
}

type BetaInfoByACL struct {
	// 任务id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 上次执行时间
	LastTime *string `json:"LastTime,omitnil,omitempty" name:"LastTime"`
}

type BlockIgnoreRule struct {
	// 规则类型，取值：1 封禁，2外部IP，3域名，4情报，5assets，6自定义策略，7入侵防御规则id （2-7属于白名单类型）
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 规则ip或白名单内容
	Ioc *string `json:"Ioc,omitnil,omitempty" name:"Ioc"`

	// 资产实例名称、自定义策略名称等
	IocName *string `json:"IocName,omitnil,omitempty" name:"IocName"`

	// 白名单信息
	IocInfo *string `json:"IocInfo,omitnil,omitempty" name:"IocInfo"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// IP
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// 危险等级
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 来源事件名称
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 方向：1入站，0出站
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 所有方向聚合成字符串
	DirectionList *string `json:"DirectionList,omitnil,omitempty" name:"DirectionList"`

	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 地理位置
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 规则类型：1封禁，2放通
	Action *int64 `json:"Action,omitnil,omitempty" name:"Action"`

	// 规则生效开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 规则生效结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 忽略原因
	IgnoreReason *string `json:"IgnoreReason,omitnil,omitempty" name:"IgnoreReason"`

	// 安全事件来源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 规则id
	UniqueId *string `json:"UniqueId,omitnil,omitempty" name:"UniqueId"`

	// 规则命中次数
	MatchTimes *int64 `json:"MatchTimes,omitnil,omitempty" name:"MatchTimes"`

	// 国家
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// 备注
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 上次命中时间
	LastHitTime *string `json:"LastHitTime,omitnil,omitempty" name:"LastHitTime"`

	// 自定义规则细节
	CustomRule *CustomWhiteRule `json:"CustomRule,omitnil,omitempty" name:"CustomRule"`

	// 1 border 2 nat 4 vpc 8 border-serial
	FwType *int64 `json:"FwType,omitnil,omitempty" name:"FwType"`
}

type CcnAssociatedInstance struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例类型
	InsType *string `json:"InsType,omitnil,omitempty" name:"InsType"`

	// 实例的网段列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CidrLst []*string `json:"CidrLst,omitnil,omitempty" name:"CidrLst"`

	// 实例所属地域
	InstanceRegion *string `json:"InstanceRegion,omitnil,omitempty" name:"InstanceRegion"`
}

type CcnSwitchInfo struct {
	// ccn的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// 开关接入模式，1:自动接入,2:手动接入
	// 注意：此字段可能返回 null，表示取不到有效值。
	SwitchMode *uint64 `json:"SwitchMode,omitnil,omitempty" name:"SwitchMode"`

	// 引流路由方法 0:多路由表, 1:策略路由
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoutingMode *int64 `json:"RoutingMode,omitnil,omitempty" name:"RoutingMode"`

	// 地域级别CIDR配置
	RegionCidrConfigs []*RegionCidrConfig `json:"RegionCidrConfigs,omitnil,omitempty" name:"RegionCidrConfigs"`

	// 互联集合对列表
	InterconnectPairs []*InterconnectPair `json:"InterconnectPairs,omitnil,omitempty" name:"InterconnectPairs"`

	// 引流通用CIDR(废弃)
	FwVpcCidr *string `json:"FwVpcCidr,omitnil,omitempty" name:"FwVpcCidr"`
}

type CfwInsStatus struct {
	// 防火墙实例id
	CfwInsId *string `json:"CfwInsId,omitnil,omitempty" name:"CfwInsId"`

	// 防火墙类型，nat：nat防火墙；ew：vpc间防火墙
	FwType *string `json:"FwType,omitnil,omitempty" name:"FwType"`

	// 实例所属地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 实例运行状态，Running：正常运行；BypassAutoFix：bypass修复；Updating：升级中；Expand：扩容中；BypassManual：手动触发bypass中；BypassAuto：自动触发bypass中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 事件时间
	EventTime *string `json:"EventTime,omitnil,omitempty" name:"EventTime"`

	// 恢复时间
	RecoverTime *string `json:"RecoverTime,omitnil,omitempty" name:"RecoverTime"`

	// 实例名称
	CfwInsName *string `json:"CfwInsName,omitnil,omitempty" name:"CfwInsName"`

	// Normal: 正常模式
	// OnlyRoute: 透明模式
	TrafficMode *string `json:"TrafficMode,omitnil,omitempty" name:"TrafficMode"`
}

type CfwNatDnatRule struct {
	// 网络协议，可选值：TCP、UDP。
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// 弹性IP。
	PublicIpAddress *string `json:"PublicIpAddress,omitnil,omitempty" name:"PublicIpAddress"`

	// 公网端口。
	PublicPort *int64 `json:"PublicPort,omitnil,omitempty" name:"PublicPort"`

	// 内网地址。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil,omitempty" name:"PrivateIpAddress"`

	// 内网端口。
	PrivatePort *int64 `json:"PrivatePort,omitnil,omitempty" name:"PrivatePort"`

	// NAT防火墙转发规则描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ClusterSwitchDetail struct {
	// 实例对象可以是ccnid类型:ccn-ad21xuds形式;nat网关类型:nat-da12daxd形式;ip类型:1.1.1.1形式等
	InsObj *string `json:"InsObj,omitnil,omitempty" name:"InsObj"`

	// 实例对象名称
	ObjName *string `json:"ObjName,omitnil,omitempty" name:"ObjName"`

	// 防火墙类型，ew：vpc间防火墙；nat：nat防火墙；border：互联网边界防火墙
	FwType *string `json:"FwType,omitnil,omitempty" name:"FwType"`

	// 资产类型，ccn：ccn实例类型；nat：nat网关类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 开关状态
	// 0 : 关闭
	// 1 : 开启
	// 2 : 开启中
	// 3 : 关闭中
	// 4 : 异常
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 开关接入模式，1：自动接入；2，手动接入，0：未选择
	SwitchMode *uint64 `json:"SwitchMode,omitnil,omitempty" name:"SwitchMode"`

	// 实例对象是否处于非集群接入场景（主备模式）
	NonCluster *int64 `json:"NonCluster,omitnil,omitempty" name:"NonCluster"`

	// ip版本，0：ipv4；1：ipv6
	IpVersion *int64 `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`

	// 关联实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachIns []*AttachInsInfo `json:"AttachIns,omitnil,omitempty" name:"AttachIns"`

	// 引流私有网络端点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Endpoints []*EndpointInfo `json:"Endpoints,omitnil,omitempty" name:"Endpoints"`

	// 入侵防护模式,0:观察;1:拦截;2:严格;3:关闭
	Idpsaction *uint64 `json:"Idpsaction,omitnil,omitempty" name:"Idpsaction"`

	// //透明模式开关,0:未开启,1:已开启
	TransEnable *uint64 `json:"TransEnable,omitnil,omitempty" name:"TransEnable"`

	// 开关状态 0关闭 1开启
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 路由模式：0：多路由表，1：策略路由
	RoutingMode *int64 `json:"RoutingMode,omitnil,omitempty" name:"RoutingMode"`

	// 是否跨租户开关 1是 0不是
	IsPeer *int64 `json:"IsPeer,omitnil,omitempty" name:"IsPeer"`

	// 跨租户appid
	PeerAppid *string `json:"PeerAppid,omitnil,omitempty" name:"PeerAppid"`

	// 跨租户操作状态 1不允许操作 0可以
	PeerStatus *int64 `json:"PeerStatus,omitnil,omitempty" name:"PeerStatus"`
}

type Column struct {
	// 列的名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 列的属性
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type CommonFilter struct {
	// 检索的键值
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 枚举类型，代表Name与Values之间的匹配关系
	// enum FilterOperatorType {
	//     //等于
	//     FILTER_OPERATOR_TYPE_EQUAL = 1;
	//     //大于
	//     FILTER_OPERATOR_TYPE_GREATER = 2;
	//     //小于
	//     FILTER_OPERATOR_TYPE_LESS = 3;
	//     //大于等于
	//     FILTER_OPERATOR_TYPE_GREATER_EQ = 4;
	//     //小于等于
	//     FILTER_OPERATOR_TYPE_LESS_EQ = 5;
	//     //不等于
	//     FILTER_OPERATOR_TYPE_NO_EQ = 6;
	//     //not in
	//     FILTER_OPERATOR_TYPE_NOT_IN = 8;
	//     //模糊匹配
	//     FILTER_OPERATOR_TYPE_FUZZINESS = 9;
	// }
	OperatorType *int64 `json:"OperatorType,omitnil,omitempty" name:"OperatorType"`

	// 检索的值，各检索值间为OR关系
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type CommonIdName struct {
	// 资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 资源名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type CommonIdNameStatus struct {
	// 资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 资源name
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type CreateAcRulesRequestParams struct {
	// 创建规则数据
	Data []*RuleInfoData `json:"Data,omitnil,omitempty" name:"Data"`

	// 0：添加（默认），1：插入
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 边id
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 访问控制规则状态
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 0：添加，1：覆盖
	Overwrite *uint64 `json:"Overwrite,omitnil,omitempty" name:"Overwrite"`

	// NAT实例ID, 参数Area存在的时候这个必传
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// portScan: 来自于端口扫描, patchImport: 来自于批量导入
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// NAT地域
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type CreateAcRulesRequest struct {
	*tchttp.BaseRequest
	
	// 创建规则数据
	Data []*RuleInfoData `json:"Data,omitnil,omitempty" name:"Data"`

	// 0：添加（默认），1：插入
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 边id
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 访问控制规则状态
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 0：添加，1：覆盖
	Overwrite *uint64 `json:"Overwrite,omitnil,omitempty" name:"Overwrite"`

	// NAT实例ID, 参数Area存在的时候这个必传
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// portScan: 来自于端口扫描, patchImport: 来自于批量导入
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// NAT地域
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

func (r *CreateAcRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAcRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	delete(f, "Type")
	delete(f, "EdgeId")
	delete(f, "Enable")
	delete(f, "Overwrite")
	delete(f, "InstanceId")
	delete(f, "From")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAcRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAcRulesResponseParams struct {
	// 状态值，0:操作成功
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 返回多余的信息
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAcRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateAcRulesResponseParams `json:"Response"`
}

func (r *CreateAcRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAcRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAddressTemplateRequestParams struct {
	// 模板名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// Type为1，ip模板eg：1.1.1.1,2.2.2.2；
	// Type为5，域名模板eg：www.qq.com,www.tencent.com
	IpString *string `json:"IpString,omitnil,omitempty" name:"IpString"`

	// 1 ip模板
	// 5 域名模板
	// 6 协议端口模板
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 协议端口模板，协议类型，4:4层协议，7:7层协议，Type=6时必填
	ProtocolType *string `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`

	// IP版本,0 IPV4;1 IPV6
	IpVersion *int64 `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`
}

type CreateAddressTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// Type为1，ip模板eg：1.1.1.1,2.2.2.2；
	// Type为5，域名模板eg：www.qq.com,www.tencent.com
	IpString *string `json:"IpString,omitnil,omitempty" name:"IpString"`

	// 1 ip模板
	// 5 域名模板
	// 6 协议端口模板
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 协议端口模板，协议类型，4:4层协议，7:7层协议，Type=6时必填
	ProtocolType *string `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`

	// IP版本,0 IPV4;1 IPV6
	IpVersion *int64 `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`
}

func (r *CreateAddressTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAddressTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Detail")
	delete(f, "IpString")
	delete(f, "Type")
	delete(f, "ProtocolType")
	delete(f, "IpVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAddressTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAddressTemplateResponseParams struct {
	// 创建结果,0成功
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一Id
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAddressTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateAddressTemplateResponseParams `json:"Response"`
}

func (r *CreateAddressTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAddressTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlertCenterIsolateRequestParams struct {
	// 处置对象,资产列表
	HandleAssetList []*string `json:"HandleAssetList,omitnil,omitempty" name:"HandleAssetList"`

	// 处置时间
	// 1  1天
	// 7   7天
	// -2 永久
	HandleTime *int64 `json:"HandleTime,omitnil,omitempty" name:"HandleTime"`

	// 当前日志方向： 0 出向 1 入向
	AlertDirection *int64 `json:"AlertDirection,omitnil,omitempty" name:"AlertDirection"`

	// 隔离类型 
	// 1 互联网入站
	// 2 互联网出站
	// 4 内网访问
	IsolateType []*int64 `json:"IsolateType,omitnil,omitempty" name:"IsolateType"`

	// 运维模式 1 IP白名单 2 身份认证  0 非运维模式
	OmMode *int64 `json:"OmMode,omitnil,omitempty" name:"OmMode"`
}

type CreateAlertCenterIsolateRequest struct {
	*tchttp.BaseRequest
	
	// 处置对象,资产列表
	HandleAssetList []*string `json:"HandleAssetList,omitnil,omitempty" name:"HandleAssetList"`

	// 处置时间
	// 1  1天
	// 7   7天
	// -2 永久
	HandleTime *int64 `json:"HandleTime,omitnil,omitempty" name:"HandleTime"`

	// 当前日志方向： 0 出向 1 入向
	AlertDirection *int64 `json:"AlertDirection,omitnil,omitempty" name:"AlertDirection"`

	// 隔离类型 
	// 1 互联网入站
	// 2 互联网出站
	// 4 内网访问
	IsolateType []*int64 `json:"IsolateType,omitnil,omitempty" name:"IsolateType"`

	// 运维模式 1 IP白名单 2 身份认证  0 非运维模式
	OmMode *int64 `json:"OmMode,omitnil,omitempty" name:"OmMode"`
}

func (r *CreateAlertCenterIsolateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlertCenterIsolateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HandleAssetList")
	delete(f, "HandleTime")
	delete(f, "AlertDirection")
	delete(f, "IsolateType")
	delete(f, "OmMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlertCenterIsolateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlertCenterIsolateResponseParams struct {
	// 返回状态码：
	// 0 成功
	// 非0 失败
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 返回信息：
	// success 成功
	// 其他
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 处置状态码：
	// 0  处置成功
	// -1 通用错误，不用处理
	// -3 表示重复，需重新刷新列表
	// 其他
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAlertCenterIsolateResponse struct {
	*tchttp.BaseResponse
	Response *CreateAlertCenterIsolateResponseParams `json:"Response"`
}

func (r *CreateAlertCenterIsolateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlertCenterIsolateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlertCenterOmitRequestParams struct {
	// 处置对象,ID列表，  IdLists和IpList二选一
	HandleIdList []*string `json:"HandleIdList,omitnil,omitempty" name:"HandleIdList"`

	// 忽略数据来源：
	// AlertTable 告警中心  InterceptionTable拦截列表
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`

	// 处置对象,事件ID列表
	HandleEventIdList []*string `json:"HandleEventIdList,omitnil,omitempty" name:"HandleEventIdList"`
}

type CreateAlertCenterOmitRequest struct {
	*tchttp.BaseRequest
	
	// 处置对象,ID列表，  IdLists和IpList二选一
	HandleIdList []*string `json:"HandleIdList,omitnil,omitempty" name:"HandleIdList"`

	// 忽略数据来源：
	// AlertTable 告警中心  InterceptionTable拦截列表
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`

	// 处置对象,事件ID列表
	HandleEventIdList []*string `json:"HandleEventIdList,omitnil,omitempty" name:"HandleEventIdList"`
}

func (r *CreateAlertCenterOmitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlertCenterOmitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HandleIdList")
	delete(f, "TableType")
	delete(f, "HandleEventIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlertCenterOmitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlertCenterOmitResponseParams struct {
	// 返回状态码：
	// 0 成功
	// 非0 失败
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 返回信息：
	// success 成功
	// 其他
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 处置状态码：
	// 0  处置成功
	// -1 通用错误，不用处理
	// -3 表示重复，需重新刷新列表
	// 其他
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAlertCenterOmitResponse struct {
	*tchttp.BaseResponse
	Response *CreateAlertCenterOmitResponseParams `json:"Response"`
}

func (r *CreateAlertCenterOmitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlertCenterOmitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlertCenterRuleRequestParams struct {
	// 处置时间
	// 1  1天
	// 7   7天
	// -2 永久
	HandleTime *int64 `json:"HandleTime,omitnil,omitempty" name:"HandleTime"`

	// 处置类型
	// 当HandleIdList 不为空时：1封禁 2放通  
	// 当HandleIpList 不为空时：3放通 4封禁
	HandleType *int64 `json:"HandleType,omitnil,omitempty" name:"HandleType"`

	// 当前日志方向： 0 出向 1 入向
	AlertDirection *int64 `json:"AlertDirection,omitnil,omitempty" name:"AlertDirection"`

	// 处置方向： 0出向 1入向 0,1出入向 3内网
	HandleDirection *string `json:"HandleDirection,omitnil,omitempty" name:"HandleDirection"`

	// 处置对象,ID列表，  IdLists和IpList二选一
	HandleIdList []*string `json:"HandleIdList,omitnil,omitempty" name:"HandleIdList"`

	// 处置对象,IP列表，  IdLists和IpList二选一
	HandleIpList []*string `json:"HandleIpList,omitnil,omitempty" name:"HandleIpList"`

	// 处置描述
	HandleComment *string `json:"HandleComment,omitnil,omitempty" name:"HandleComment"`

	// 放通原因:
	// 0默认 1重复 2误报 3紧急放通
	IgnoreReason *int64 `json:"IgnoreReason,omitnil,omitempty" name:"IgnoreReason"`

	// 封禁域名-保留字段
	BlockDomain *string `json:"BlockDomain,omitnil,omitempty" name:"BlockDomain"`
}

type CreateAlertCenterRuleRequest struct {
	*tchttp.BaseRequest
	
	// 处置时间
	// 1  1天
	// 7   7天
	// -2 永久
	HandleTime *int64 `json:"HandleTime,omitnil,omitempty" name:"HandleTime"`

	// 处置类型
	// 当HandleIdList 不为空时：1封禁 2放通  
	// 当HandleIpList 不为空时：3放通 4封禁
	HandleType *int64 `json:"HandleType,omitnil,omitempty" name:"HandleType"`

	// 当前日志方向： 0 出向 1 入向
	AlertDirection *int64 `json:"AlertDirection,omitnil,omitempty" name:"AlertDirection"`

	// 处置方向： 0出向 1入向 0,1出入向 3内网
	HandleDirection *string `json:"HandleDirection,omitnil,omitempty" name:"HandleDirection"`

	// 处置对象,ID列表，  IdLists和IpList二选一
	HandleIdList []*string `json:"HandleIdList,omitnil,omitempty" name:"HandleIdList"`

	// 处置对象,IP列表，  IdLists和IpList二选一
	HandleIpList []*string `json:"HandleIpList,omitnil,omitempty" name:"HandleIpList"`

	// 处置描述
	HandleComment *string `json:"HandleComment,omitnil,omitempty" name:"HandleComment"`

	// 放通原因:
	// 0默认 1重复 2误报 3紧急放通
	IgnoreReason *int64 `json:"IgnoreReason,omitnil,omitempty" name:"IgnoreReason"`

	// 封禁域名-保留字段
	BlockDomain *string `json:"BlockDomain,omitnil,omitempty" name:"BlockDomain"`
}

func (r *CreateAlertCenterRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlertCenterRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HandleTime")
	delete(f, "HandleType")
	delete(f, "AlertDirection")
	delete(f, "HandleDirection")
	delete(f, "HandleIdList")
	delete(f, "HandleIpList")
	delete(f, "HandleComment")
	delete(f, "IgnoreReason")
	delete(f, "BlockDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlertCenterRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAlertCenterRuleResponseParams struct {
	// 返回状态码：
	// 0 成功
	// 非0 失败
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 返回信息：
	// success 成功
	// 其他
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 处置状态码：
	// 0  处置成功
	// -1 通用错误，不用处理
	// -3 表示重复，需重新刷新列表
	// 其他
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAlertCenterRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateAlertCenterRuleResponseParams `json:"Response"`
}

func (r *CreateAlertCenterRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlertCenterRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBlockIgnoreRuleListRequestParams struct {
	// 规则列表
	Rules []*IntrusionDefenseRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 规则类型，1封禁，2放通，不支持域名封禁
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 是否覆盖重复数据，1覆盖，非1不覆盖，跳过重复数据
	CoverDuplicate *int64 `json:"CoverDuplicate,omitnil,omitempty" name:"CoverDuplicate"`
}

type CreateBlockIgnoreRuleListRequest struct {
	*tchttp.BaseRequest
	
	// 规则列表
	Rules []*IntrusionDefenseRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 规则类型，1封禁，2放通，不支持域名封禁
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 是否覆盖重复数据，1覆盖，非1不覆盖，跳过重复数据
	CoverDuplicate *int64 `json:"CoverDuplicate,omitnil,omitempty" name:"CoverDuplicate"`
}

func (r *CreateBlockIgnoreRuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBlockIgnoreRuleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	delete(f, "RuleType")
	delete(f, "CoverDuplicate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBlockIgnoreRuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBlockIgnoreRuleListResponseParams struct {
	// 成功返回
	List []*IocListData `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBlockIgnoreRuleListResponse struct {
	*tchttp.BaseResponse
	Response *CreateBlockIgnoreRuleListResponseParams `json:"Response"`
}

func (r *CreateBlockIgnoreRuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBlockIgnoreRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBlockIgnoreRuleNewRequestParams struct {
	// 非自定义类型规则列表
	Rules []*BanAndAllowRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// RuleType: 1黑名单 2外部IP 3域名 4情报 5资产 6自定义规则  7入侵防御规则
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 删除白名单冲突地址并继续添加/删除封禁列表冲突地址并继续添加；表示是否覆盖重复数据，1为覆盖，非1不覆盖，跳过重复数据
	CoverDuplicate *int64 `json:"CoverDuplicate,omitnil,omitempty" name:"CoverDuplicate"`
}

type CreateBlockIgnoreRuleNewRequest struct {
	*tchttp.BaseRequest
	
	// 非自定义类型规则列表
	Rules []*BanAndAllowRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// RuleType: 1黑名单 2外部IP 3域名 4情报 5资产 6自定义规则  7入侵防御规则
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 删除白名单冲突地址并继续添加/删除封禁列表冲突地址并继续添加；表示是否覆盖重复数据，1为覆盖，非1不覆盖，跳过重复数据
	CoverDuplicate *int64 `json:"CoverDuplicate,omitnil,omitempty" name:"CoverDuplicate"`
}

func (r *CreateBlockIgnoreRuleNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBlockIgnoreRuleNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	delete(f, "RuleType")
	delete(f, "CoverDuplicate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBlockIgnoreRuleNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBlockIgnoreRuleNewResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBlockIgnoreRuleNewResponse struct {
	*tchttp.BaseResponse
	Response *CreateBlockIgnoreRuleNewResponseParams `json:"Response"`
}

func (r *CreateBlockIgnoreRuleNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBlockIgnoreRuleNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateChooseVpcsRequestParams struct {
	// vpc列表
	VpcList []*string `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// zone列表
	AllZoneList []*VpcZoneData `json:"AllZoneList,omitnil,omitempty" name:"AllZoneList"`
}

type CreateChooseVpcsRequest struct {
	*tchttp.BaseRequest
	
	// vpc列表
	VpcList []*string `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// zone列表
	AllZoneList []*VpcZoneData `json:"AllZoneList,omitnil,omitempty" name:"AllZoneList"`
}

func (r *CreateChooseVpcsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChooseVpcsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcList")
	delete(f, "AllZoneList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateChooseVpcsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateChooseVpcsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateChooseVpcsResponse struct {
	*tchttp.BaseResponse
	Response *CreateChooseVpcsResponseParams `json:"Response"`
}

func (r *CreateChooseVpcsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateChooseVpcsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatabaseWhiteListRulesRequestParams struct {
	// 创建白名单数据
	DatabaseWhiteListRuleData []*DatabaseWhiteListRuleData `json:"DatabaseWhiteListRuleData,omitnil,omitempty" name:"DatabaseWhiteListRuleData"`
}

type CreateDatabaseWhiteListRulesRequest struct {
	*tchttp.BaseRequest
	
	// 创建白名单数据
	DatabaseWhiteListRuleData []*DatabaseWhiteListRuleData `json:"DatabaseWhiteListRuleData,omitnil,omitempty" name:"DatabaseWhiteListRuleData"`
}

func (r *CreateDatabaseWhiteListRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatabaseWhiteListRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseWhiteListRuleData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatabaseWhiteListRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatabaseWhiteListRulesResponseParams struct {
	// 状态值，0:添加成功，非0：添加失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDatabaseWhiteListRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateDatabaseWhiteListRulesResponseParams `json:"Response"`
}

func (r *CreateDatabaseWhiteListRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatabaseWhiteListRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatFwInstanceRequestParams struct {
	// 防火墙实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 带宽
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 模式 1：接入模式；0：新增模式
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 新增模式传递参数，其中NewModeItems和NatgwList至少传递一种。
	NewModeItems *NewModeItems `json:"NewModeItems,omitnil,omitempty" name:"NewModeItems"`

	// 接入模式接入的nat网关列表，其中NewModeItems和NatgwList至少传递一种。
	NatGwList []*string `json:"NatGwList,omitnil,omitempty" name:"NatGwList"`

	// 主可用区，为空则选择默认可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 备可用区，为空则选择默认可用区
	ZoneBak *string `json:"ZoneBak,omitnil,omitempty" name:"ZoneBak"`

	// 异地灾备 1：使用异地灾备；0：不使用异地灾备；为空则默认不使用异地灾备
	CrossAZone *int64 `json:"CrossAZone,omitnil,omitempty" name:"CrossAZone"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil,omitempty" name:"FwCidrInfo"`
}

type CreateNatFwInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 带宽
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 模式 1：接入模式；0：新增模式
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 新增模式传递参数，其中NewModeItems和NatgwList至少传递一种。
	NewModeItems *NewModeItems `json:"NewModeItems,omitnil,omitempty" name:"NewModeItems"`

	// 接入模式接入的nat网关列表，其中NewModeItems和NatgwList至少传递一种。
	NatGwList []*string `json:"NatGwList,omitnil,omitempty" name:"NatGwList"`

	// 主可用区，为空则选择默认可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 备可用区，为空则选择默认可用区
	ZoneBak *string `json:"ZoneBak,omitnil,omitempty" name:"ZoneBak"`

	// 异地灾备 1：使用异地灾备；0：不使用异地灾备；为空则默认不使用异地灾备
	CrossAZone *int64 `json:"CrossAZone,omitnil,omitempty" name:"CrossAZone"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil,omitempty" name:"FwCidrInfo"`
}

func (r *CreateNatFwInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatFwInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Width")
	delete(f, "Mode")
	delete(f, "NewModeItems")
	delete(f, "NatGwList")
	delete(f, "Zone")
	delete(f, "ZoneBak")
	delete(f, "CrossAZone")
	delete(f, "FwCidrInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNatFwInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatFwInstanceResponseParams struct {
	// 防火墙实例id
	CfwInsId *string `json:"CfwInsId,omitnil,omitempty" name:"CfwInsId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNatFwInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateNatFwInstanceResponseParams `json:"Response"`
}

func (r *CreateNatFwInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatFwInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatFwInstanceWithDomainRequestParams struct {
	// 防火墙实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 带宽
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 模式 1：接入模式；0：新增模式
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 新增模式传递参数，其中NewModeItems和NatgwList至少传递一种。
	NewModeItems *NewModeItems `json:"NewModeItems,omitnil,omitempty" name:"NewModeItems"`

	// 接入模式接入的nat网关列表，其中NewModeItems和NatgwList至少传递一种。
	NatGwList []*string `json:"NatGwList,omitnil,omitempty" name:"NatGwList"`

	// 主可用区，为空则选择默认可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 备可用区，为空则选择默认可用区
	ZoneBak *string `json:"ZoneBak,omitnil,omitempty" name:"ZoneBak"`

	// 异地灾备 1：使用异地灾备；0：不使用异地灾备；为空则默认不使用异地灾备
	CrossAZone *int64 `json:"CrossAZone,omitnil,omitempty" name:"CrossAZone"`

	// 0不创建域名,1创建域名
	IsCreateDomain *int64 `json:"IsCreateDomain,omitnil,omitempty" name:"IsCreateDomain"`

	// 如果要创建域名则必填
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil,omitempty" name:"FwCidrInfo"`
}

type CreateNatFwInstanceWithDomainRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 带宽
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 模式 1：接入模式；0：新增模式
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 新增模式传递参数，其中NewModeItems和NatgwList至少传递一种。
	NewModeItems *NewModeItems `json:"NewModeItems,omitnil,omitempty" name:"NewModeItems"`

	// 接入模式接入的nat网关列表，其中NewModeItems和NatgwList至少传递一种。
	NatGwList []*string `json:"NatGwList,omitnil,omitempty" name:"NatGwList"`

	// 主可用区，为空则选择默认可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 备可用区，为空则选择默认可用区
	ZoneBak *string `json:"ZoneBak,omitnil,omitempty" name:"ZoneBak"`

	// 异地灾备 1：使用异地灾备；0：不使用异地灾备；为空则默认不使用异地灾备
	CrossAZone *int64 `json:"CrossAZone,omitnil,omitempty" name:"CrossAZone"`

	// 0不创建域名,1创建域名
	IsCreateDomain *int64 `json:"IsCreateDomain,omitnil,omitempty" name:"IsCreateDomain"`

	// 如果要创建域名则必填
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil,omitempty" name:"FwCidrInfo"`
}

func (r *CreateNatFwInstanceWithDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatFwInstanceWithDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Width")
	delete(f, "Mode")
	delete(f, "NewModeItems")
	delete(f, "NatGwList")
	delete(f, "Zone")
	delete(f, "ZoneBak")
	delete(f, "CrossAZone")
	delete(f, "IsCreateDomain")
	delete(f, "Domain")
	delete(f, "FwCidrInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNatFwInstanceWithDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNatFwInstanceWithDomainResponseParams struct {
	// nat实例信息
	CfwInsId *string `json:"CfwInsId,omitnil,omitempty" name:"CfwInsId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNatFwInstanceWithDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateNatFwInstanceWithDomainResponseParams `json:"Response"`
}

func (r *CreateNatFwInstanceWithDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNatFwInstanceWithDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNatRuleItem struct {
	// 访问源示例： net：IP/CIDR(192.168.0.2)
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// 访问源类型：入向规则时类型可以为 ip,net,template,location；出向规则时可以为 ip,net,template,instance,group,tag
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 访问目的示例： net：IP/CIDR(192.168.0.2) domain：域名规则，例如*.qq.com
	TargetContent *string `json:"TargetContent,omitnil,omitempty" name:"TargetContent"`

	// 访问目的类型：入向规则时类型可以为ip,net,template,instance,group,tag；出向规则时可以为  ip,net,domain,template,location
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 协议，可选的值： TCP UDP ICMP ANY HTTP HTTPS HTTP/HTTPS SMTP SMTPS SMTP/SMTPS FTP DNS
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 访问控制策略中设置的流量通过云防火墙的方式。取值： accept：放行 drop：拒绝 log：观察
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// 访问控制策略的端口。取值： -1/-1：全部端口 80：80端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 规则序号
	OrderIndex *int64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// 规则状态，true表示启用，false表示禁用
	Enable *string `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 规则对应的唯一id，创建规则时无需填写
	Uuid *int64 `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 端口协议组ID
	ParamTemplateId *string `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// 内部id
	InternalUuid *int64 `json:"InternalUuid,omitnil,omitempty" name:"InternalUuid"`

	// 规则生效的范围：ALL，全局生效；ap-guangzhou，生效的地域；cfwnat-xxx，生效基于实例维度
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`
}

type CreateRuleItem struct {
	// 访问源示例： net：IP/CIDR(192.168.0.2)
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// 访问源类型：入向规则时类型可以为 ip,net,template,location；出向规则时可以为 ip,net,template,instance,group,tag
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 访问目的示例： net：IP/CIDR(192.168.0.2) domain：域名规则，例如*.qq.com
	TargetContent *string `json:"TargetContent,omitnil,omitempty" name:"TargetContent"`

	// 访问目的类型：入向规则时类型可以为ip,net,template,instance,group,tag；出向规则时可以为  ip,net,domain,template,location
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 协议，可选的值： TCP UDP ICMP ANY HTTP HTTPS HTTP/HTTPS SMTP SMTPS SMTP/SMTPS FTP DNS
	// 1. 入方向  旁路防火墙/全局规则 仅支持TCP
	// 
	// 2.出方向  旁路防火墙/全局规则 仅支持TCP HTTP/HTTPS TLS/SSL
	// 
	// 3.domain  请选择七层协议 如HTTP/HTTPS
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 访问控制策略中设置的流量通过云防火墙的方式。取值： accept：放行 drop：拒绝 log：观察
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// 访问控制策略的端口。取值： -1/-1：全部端口 80：80端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 规则序号
	OrderIndex *int64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// 规则对应的唯一id，创建规则时无需填写
	Uuid *int64 `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// 规则状态，true表示启用，false表示禁用
	Enable *string `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// all
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 0，正常规则添加；1，入侵检测添加
	RuleSource *int64 `json:"RuleSource,omitnil,omitempty" name:"RuleSource"`

	// 告警Id
	LogId *string `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 端都协议组ID
	ParamTemplateId *string `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`
}

// Predefined struct for user
type CreateSecurityGroupRulesRequestParams struct {
	// 添加的企业安全组规则数据
	Data []*SecurityGroupListData `json:"Data,omitnil,omitempty" name:"Data"`

	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 0：后插，1：前插，2：中插，默认0
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 添加后是否启用规则，0：不启用，1：启用，默认1
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type CreateSecurityGroupRulesRequest struct {
	*tchttp.BaseRequest
	
	// 添加的企业安全组规则数据
	Data []*SecurityGroupListData `json:"Data,omitnil,omitempty" name:"Data"`

	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 0：后插，1：前插，2：中插，默认0
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 添加后是否启用规则，0：不启用，1：启用，默认1
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

func (r *CreateSecurityGroupRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	delete(f, "Direction")
	delete(f, "Type")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityGroupRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityGroupRulesResponseParams struct {
	// 状态值，0：添加成功，非0：添加失败
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSecurityGroupRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityGroupRulesResponseParams `json:"Response"`
}

func (r *CreateSecurityGroupRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcFwGroupRequestParams struct {
	// VPC防火墙(组)名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模式 1：CCN云联网模式；0：私有网络模式 2: sase 模式 3：ccn 高级模式 4: 私有网络(跨租户单边模式)
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 防火墙(组)下的防火墙实例列表
	VpcFwInstances []*VpcFwInstance `json:"VpcFwInstances,omitnil,omitempty" name:"VpcFwInstances"`

	// 防火墙实例的开关模式
	// 1: 单点互通
	// 2: 多点互通
	// 3: 全互通
	// 4: 自定义路由
	SwitchMode *int64 `json:"SwitchMode,omitnil,omitempty" name:"SwitchMode"`

	// auto 自动选择防火墙网段
	// 10.10.10.0/24 用户输入的防火墙网段
	FwVpcCidr *string `json:"FwVpcCidr,omitnil,omitempty" name:"FwVpcCidr"`

	// 云联网id ，适用于云联网模式
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil,omitempty" name:"FwCidrInfo"`

	// 跨租户管理员模式  1管理员 2多账号
	CrossUserMode *string `json:"CrossUserMode,omitnil,omitempty" name:"CrossUserMode"`
}

type CreateVpcFwGroupRequest struct {
	*tchttp.BaseRequest
	
	// VPC防火墙(组)名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模式 1：CCN云联网模式；0：私有网络模式 2: sase 模式 3：ccn 高级模式 4: 私有网络(跨租户单边模式)
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 防火墙(组)下的防火墙实例列表
	VpcFwInstances []*VpcFwInstance `json:"VpcFwInstances,omitnil,omitempty" name:"VpcFwInstances"`

	// 防火墙实例的开关模式
	// 1: 单点互通
	// 2: 多点互通
	// 3: 全互通
	// 4: 自定义路由
	SwitchMode *int64 `json:"SwitchMode,omitnil,omitempty" name:"SwitchMode"`

	// auto 自动选择防火墙网段
	// 10.10.10.0/24 用户输入的防火墙网段
	FwVpcCidr *string `json:"FwVpcCidr,omitnil,omitempty" name:"FwVpcCidr"`

	// 云联网id ，适用于云联网模式
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil,omitempty" name:"FwCidrInfo"`

	// 跨租户管理员模式  1管理员 2多账号
	CrossUserMode *string `json:"CrossUserMode,omitnil,omitempty" name:"CrossUserMode"`
}

func (r *CreateVpcFwGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcFwGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Mode")
	delete(f, "VpcFwInstances")
	delete(f, "SwitchMode")
	delete(f, "FwVpcCidr")
	delete(f, "CcnId")
	delete(f, "FwCidrInfo")
	delete(f, "CrossUserMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVpcFwGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVpcFwGroupResponseParams struct {
	// 防火墙组ID
	FwGroupId *string `json:"FwGroupId,omitnil,omitempty" name:"FwGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVpcFwGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateVpcFwGroupResponseParams `json:"Response"`
}

func (r *CreateVpcFwGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVpcFwGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomWhiteRule struct {
	// 访问目的
	DstIP *string `json:"DstIP,omitnil,omitempty" name:"DstIP"`

	// 规则ID
	IdsRuleId *string `json:"IdsRuleId,omitnil,omitempty" name:"IdsRuleId"`

	// 规则名称
	IdsRuleName *string `json:"IdsRuleName,omitnil,omitempty" name:"IdsRuleName"`

	// 访问源
	SrcIP *string `json:"SrcIP,omitnil,omitempty" name:"SrcIP"`
}

type DatabaseWhiteListRuleData struct {
	// 访问源
	SourceIp *string `json:"SourceIp,omitnil,omitempty" name:"SourceIp"`

	// 访问源类型，1 ip；6 实例；100 资源分组
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 访问目的
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// 访问目的类型，1 ip；6 实例；100 资源分组
	TargetType *int64 `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 规则描述
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 是否地域规则，0不是 1是
	IsRegionRule *int64 `json:"IsRegionRule,omitnil,omitempty" name:"IsRegionRule"`

	// 是否云厂商规则，0不是 1 时
	IsCloudRule *int64 `json:"IsCloudRule,omitnil,omitempty" name:"IsCloudRule"`

	// 是否启用，0 不启用，1启用
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 地域码1
	FirstLevelRegionCode *int64 `json:"FirstLevelRegionCode,omitnil,omitempty" name:"FirstLevelRegionCode"`

	// 地域码2
	SecondLevelRegionCode *int64 `json:"SecondLevelRegionCode,omitnil,omitempty" name:"SecondLevelRegionCode"`

	// 地域名称1
	FirstLevelRegionName *string `json:"FirstLevelRegionName,omitnil,omitempty" name:"FirstLevelRegionName"`

	// 地域名称2
	SecondLevelRegionName *string `json:"SecondLevelRegionName,omitnil,omitempty" name:"SecondLevelRegionName"`

	// 云厂商码
	CloudCode *string `json:"CloudCode,omitnil,omitempty" name:"CloudCode"`
}

// Predefined struct for user
type DeleteAcRuleRequestParams struct {
	// 删除规则对应的id值, 对应获取规则列表接口的Id 值
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// EdgeId值两个vpc间的边id
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// NAT地域， 如ap-shanghai/ap-guangzhou/ap-chongqing等
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type DeleteAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 删除规则对应的id值, 对应获取规则列表接口的Id 值
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// EdgeId值两个vpc间的边id
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// NAT地域， 如ap-shanghai/ap-guangzhou/ap-chongqing等
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

func (r *DeleteAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Direction")
	delete(f, "EdgeId")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAcRuleResponseParams struct {
	// 状态值 0: 删除成功, !0: 删除失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 返回多余的信息
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAcRuleResponseParams `json:"Response"`
}

func (r *DeleteAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAddressTemplateRequestParams struct {
	// 模板id
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`
}

type DeleteAddressTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板id
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`
}

func (r *DeleteAddressTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAddressTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAddressTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAddressTemplateResponseParams struct {
	// 删除结果,0成功
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAddressTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAddressTemplateResponseParams `json:"Response"`
}

func (r *DeleteAddressTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAddressTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBlockIgnoreRuleListRequestParams struct {
	// 规则列表
	Rules []*IocListData `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 规则类型，1封禁，2放通，不支持域名封禁
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`
}

type DeleteBlockIgnoreRuleListRequest struct {
	*tchttp.BaseRequest
	
	// 规则列表
	Rules []*IocListData `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 规则类型，1封禁，2放通，不支持域名封禁
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`
}

func (r *DeleteBlockIgnoreRuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlockIgnoreRuleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	delete(f, "RuleType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBlockIgnoreRuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBlockIgnoreRuleListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBlockIgnoreRuleListResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBlockIgnoreRuleListResponseParams `json:"Response"`
}

func (r *DeleteBlockIgnoreRuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlockIgnoreRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBlockIgnoreRuleNewRequestParams struct {
	// 是否删除全部
	DeleteAll *int64 `json:"DeleteAll,omitnil,omitempty" name:"DeleteAll"`

	// 规则列表
	Rules []*BanAndAllowRuleDel `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 封禁：1，放通：100，
	// 主要用于全部删除时区分列表类型
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// blocklist 封禁列表 whitelist 白名单列表
	ShowType *string `json:"ShowType,omitnil,omitempty" name:"ShowType"`
}

type DeleteBlockIgnoreRuleNewRequest struct {
	*tchttp.BaseRequest
	
	// 是否删除全部
	DeleteAll *int64 `json:"DeleteAll,omitnil,omitempty" name:"DeleteAll"`

	// 规则列表
	Rules []*BanAndAllowRuleDel `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 封禁：1，放通：100，
	// 主要用于全部删除时区分列表类型
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// blocklist 封禁列表 whitelist 白名单列表
	ShowType *string `json:"ShowType,omitnil,omitempty" name:"ShowType"`
}

func (r *DeleteBlockIgnoreRuleNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlockIgnoreRuleNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeleteAll")
	delete(f, "Rules")
	delete(f, "RuleType")
	delete(f, "ShowType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBlockIgnoreRuleNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBlockIgnoreRuleNewResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBlockIgnoreRuleNewResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBlockIgnoreRuleNewResponseParams `json:"Response"`
}

func (r *DeleteBlockIgnoreRuleNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBlockIgnoreRuleNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNatFwInstanceRequestParams struct {
	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`
}

type DeleteNatFwInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`
}

func (r *DeleteNatFwInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatFwInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CfwInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNatFwInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNatFwInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNatFwInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNatFwInstanceResponseParams `json:"Response"`
}

func (r *DeleteNatFwInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNatFwInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRemoteAccessDomainRequestParams struct {
	// 域名列表
	AccessDomainList []*string `json:"AccessDomainList,omitnil,omitempty" name:"AccessDomainList"`
}

type DeleteRemoteAccessDomainRequest struct {
	*tchttp.BaseRequest
	
	// 域名列表
	AccessDomainList []*string `json:"AccessDomainList,omitnil,omitempty" name:"AccessDomainList"`
}

func (r *DeleteRemoteAccessDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRemoteAccessDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessDomainList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRemoteAccessDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRemoteAccessDomainResponseParams struct {
	// 状态值 0：删除成功，非 0：删除失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRemoteAccessDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRemoteAccessDomainResponseParams `json:"Response"`
}

func (r *DeleteRemoteAccessDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRemoteAccessDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceGroupRequestParams struct {
	// 组id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DeleteResourceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 组id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *DeleteResourceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteResourceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteResourceGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteResourceGroupResponseParams `json:"Response"`
}

func (r *DeleteResourceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupRuleRequestParams struct {
	// 所需要删除规则的ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 腾讯云地域的英文简写
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 是否删除反向规则，0：否，1：是
	IsDelReverse *uint64 `json:"IsDelReverse,omitnil,omitempty" name:"IsDelReverse"`
}

type DeleteSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// 所需要删除规则的ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 腾讯云地域的英文简写
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 是否删除反向规则，0：否，1：是
	IsDelReverse *uint64 `json:"IsDelReverse,omitnil,omitempty" name:"IsDelReverse"`
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
	delete(f, "Id")
	delete(f, "Area")
	delete(f, "Direction")
	delete(f, "IsDelReverse")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityGroupRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityGroupRuleResponseParams struct {
	// 状态值，0：成功，非0：失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 返回多余的信息
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

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
type DeleteVpcFwGroupRequestParams struct {
	// 防火墙(组)Id
	FwGroupId *string `json:"FwGroupId,omitnil,omitempty" name:"FwGroupId"`

	// 是否删除整个防火墙(组)
	// 0：不删除防火墙(组)，只删除单独实例
	// 1：删除整个防火墙(组)
	DeleteFwGroup *int64 `json:"DeleteFwGroup,omitnil,omitempty" name:"DeleteFwGroup"`

	// 待删除的防火墙实例数组
	VpcFwInsList []*string `json:"VpcFwInsList,omitnil,omitempty" name:"VpcFwInsList"`
}

type DeleteVpcFwGroupRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙(组)Id
	FwGroupId *string `json:"FwGroupId,omitnil,omitempty" name:"FwGroupId"`

	// 是否删除整个防火墙(组)
	// 0：不删除防火墙(组)，只删除单独实例
	// 1：删除整个防火墙(组)
	DeleteFwGroup *int64 `json:"DeleteFwGroup,omitnil,omitempty" name:"DeleteFwGroup"`

	// 待删除的防火墙实例数组
	VpcFwInsList []*string `json:"VpcFwInsList,omitnil,omitempty" name:"VpcFwInsList"`
}

func (r *DeleteVpcFwGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcFwGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FwGroupId")
	delete(f, "DeleteFwGroup")
	delete(f, "VpcFwInsList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpcFwGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcFwGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVpcFwGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpcFwGroupResponseParams `json:"Response"`
}

func (r *DeleteVpcFwGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcFwGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescAcItem struct {
	// 访问源
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// 访问目的
	TargetContent *string `json:"TargetContent,omitnil,omitempty" name:"TargetContent"`

	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 访问控制策略中设置的流量通过云防火墙的方式。取值： accept：放行 drop：拒绝 log：观察
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 命中次数
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// 访问源类型：入向规则时类型可以为 ip,net,template,location；出向规则时可以为 ip,net,template,instance,group,tag
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 访问目的类型：入向规则时类型可以为ip,net,template,instance,group,tag；出向规则时可以为 ip,net,domain,template,location,dnsparse
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 规则对应的唯一id
	Uuid *uint64 `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// 规则有效性
	Invalid *uint64 `json:"Invalid,omitnil,omitempty" name:"Invalid"`

	// 0为正常规则,1为地域规则
	IsRegion *uint64 `json:"IsRegion,omitnil,omitempty" name:"IsRegion"`

	// 国家id
	CountryCode *uint64 `json:"CountryCode,omitnil,omitempty" name:"CountryCode"`

	// 城市id
	CityCode *uint64 `json:"CityCode,omitnil,omitempty" name:"CityCode"`

	// 国家名称
	CountryName *string `json:"CountryName,omitnil,omitempty" name:"CountryName"`

	// 省名称
	CityName *string `json:"CityName,omitnil,omitempty" name:"CityName"`

	// 云厂商code
	CloudCode *string `json:"CloudCode,omitnil,omitempty" name:"CloudCode"`

	// 0为正常规则,1为云厂商规则
	IsCloud *uint64 `json:"IsCloud,omitnil,omitempty" name:"IsCloud"`

	// 规则状态，true表示启用，false表示禁用
	Enable *string `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 规则方向：1，入向；0，出向
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 内部使用的uuid，一般情况下不会使用到该字段
	InternalUuid *int64 `json:"InternalUuid,omitnil,omitempty" name:"InternalUuid"`

	// 规则状态，查询规则命中详情时该字段有效，0：新增，1: 已删除, 2: 编辑删除
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 关联任务详情
	BetaList []*BetaInfoByACL `json:"BetaList,omitnil,omitempty" name:"BetaList"`

	// （1）互联网边界防火墙，生效范围：serial，串行；side，旁路；all，全局；
	// （2）NAT边界防火墙：ALL，全局生效；ap-guangzhou，生效的地域；cfwnat-xxx，生效基于实例维度
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 生效范围描述
	ScopeDesc *string `json:"ScopeDesc,omitnil,omitempty" name:"ScopeDesc"`

	// 互联网边界防火墙使用的内部规则id
	InternetBorderUuid *string `json:"InternetBorderUuid,omitnil,omitempty" name:"InternetBorderUuid"`

	// 协议端口组名称
	ParamTemplateName *string `json:"ParamTemplateName,omitnil,omitempty" name:"ParamTemplateName"`

	// 协议端口组ID
	ParamTemplateId *string `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// 访问源名称
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// 访问目的名称
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// 规则最近命中时间
	LastHitTime *string `json:"LastHitTime,omitnil,omitempty" name:"LastHitTime"`

	// 地区简称
	CountryKey *string `json:"CountryKey,omitnil,omitempty" name:"CountryKey"`

	// 省份、城市简称
	CityKey *string `json:"CityKey,omitnil,omitempty" name:"CityKey"`

	// 规则创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 规则最近更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type DescNatDnatRule struct {
	// id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 网络协议，可选值：TCP、UDP。
	IpProtocol *string `json:"IpProtocol,omitnil,omitempty" name:"IpProtocol"`

	// 弹性IP。
	PublicIpAddress *string `json:"PublicIpAddress,omitnil,omitempty" name:"PublicIpAddress"`

	// 公网端口。
	PublicPort *int64 `json:"PublicPort,omitnil,omitempty" name:"PublicPort"`

	// 内网地址。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil,omitempty" name:"PrivateIpAddress"`

	// 内网端口。
	PrivatePort *int64 `json:"PrivatePort,omitnil,omitempty" name:"PrivatePort"`

	// NAT防火墙转发规则描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否被关联引用，如被远程运维使用
	IsReferenced *int64 `json:"IsReferenced,omitnil,omitempty" name:"IsReferenced"`

	// 所属防火墙实例id
	FwInsId *string `json:"FwInsId,omitnil,omitempty" name:"FwInsId"`

	// 关联的nat网关Id
	NatGwId *string `json:"NatGwId,omitnil,omitempty" name:"NatGwId"`
}

// Predefined struct for user
type DescribeAcListsRequestParams struct {
	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 策略
	Strategy *string `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 出站还是入站，1：入站，0：出站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// EdgeId值
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 规则是否开启，'0': 未开启，'1': 开启, 默认为'0'
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 地域
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeAcListsRequest struct {
	*tchttp.BaseRequest
	
	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 策略
	Strategy *string `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 出站还是入站，1：入站，0：出站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// EdgeId值
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 规则是否开启，'0': 未开启，'1': 开启, 默认为'0'
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 地域
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeAcListsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAcListsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Protocol")
	delete(f, "Strategy")
	delete(f, "SearchValue")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Direction")
	delete(f, "EdgeId")
	delete(f, "Status")
	delete(f, "Area")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAcListsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAcListsResponseParams struct {
	// 总条数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 访问控制列表数据
	Data []*AcListsData `json:"Data,omitnil,omitempty" name:"Data"`

	// 不算筛选条数的总条数
	AllTotal *uint64 `json:"AllTotal,omitnil,omitempty" name:"AllTotal"`

	// 访问控制规则全部启用/全部停用
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAcListsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAcListsResponseParams `json:"Response"`
}

func (r *DescribeAcListsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAcListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAclRuleRequestParams struct {
	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 需要查询的索引，特定场景使用，可不填
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值，默认为asc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序所用到的字段，默认为sequence
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 需要查询的索引，特定场景使用，可不填
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值，默认为asc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序所用到的字段，默认为sequence
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Index")
	delete(f, "Filters")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAclRuleResponseParams struct {
	// 总条数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// nat访问控制列表数据
	Data []*DescAcItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 未过滤的总条数
	AllTotal *uint64 `json:"AllTotal,omitnil,omitempty" name:"AllTotal"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAclRuleResponseParams `json:"Response"`
}

func (r *DescribeAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressTemplateListRequestParams struct {
	// 偏移量，分页用
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 条数，分页用
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段，取值：UpdateTime最近更新时间，RulesNum关联规则数
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 排序，取值 ：asc正序，desc逆序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 检索地址模板唯一id
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// 模板类型，取值：1：ip模板，5：域名模板，6：协议端口模板
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// 模板Id
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeAddressTemplateListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，分页用
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 条数，分页用
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段，取值：UpdateTime最近更新时间，RulesNum关联规则数
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 排序，取值 ：asc正序，desc逆序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 检索地址模板唯一id
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// 模板类型，取值：1：ip模板，5：域名模板，6：协议端口模板
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// 模板Id
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribeAddressTemplateListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressTemplateListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "By")
	delete(f, "Order")
	delete(f, "SearchValue")
	delete(f, "Uuid")
	delete(f, "TemplateType")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddressTemplateListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressTemplateListResponseParams struct {
	// 模板总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 模板列表数据
	Data []*TemplateListInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 模板名称列表
	NameList []*string `json:"NameList,omitnil,omitempty" name:"NameList"`

	// Ip地址模板数量
	IpTemplateCount *int64 `json:"IpTemplateCount,omitnil,omitempty" name:"IpTemplateCount"`

	// 域名地址模板数量
	DomainTemplateCount *int64 `json:"DomainTemplateCount,omitnil,omitempty" name:"DomainTemplateCount"`

	// 协议端口模板数量
	PortTemplateCount *int64 `json:"PortTemplateCount,omitnil,omitempty" name:"PortTemplateCount"`

	// 已使用的地址模板数
	UsedTemplateCount *int64 `json:"UsedTemplateCount,omitnil,omitempty" name:"UsedTemplateCount"`

	// 地址模板配额数量
	TemplateQuotaCount *int64 `json:"TemplateQuotaCount,omitnil,omitempty" name:"TemplateQuotaCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAddressTemplateListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAddressTemplateListResponseParams `json:"Response"`
}

func (r *DescribeAddressTemplateListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAddressTemplateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetSyncRequestParams struct {

}

type DescribeAssetSyncRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAssetSyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetSyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetSyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetSyncResponseParams struct {
	// 返回状态
	// 1 更新中
	// 2 更新完成
	// 3 更新失败
	// 4 更新失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// success 成功
	// 其他失败
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 0 成功
	// 非0 失败
	ReturnCode *uint64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAssetSyncResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetSyncResponseParams `json:"Response"`
}

func (r *DescribeAssetSyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetSyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssociatedInstanceListRequestParams struct {
	// 列表偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页记录条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 地域代码（例：ap-guangzhou）,支持腾讯云全地域
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 额外检索条件（JSON字符串）
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 排序方式（asc:升序,desc:降序）
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 实例类型,'3'是cvm实例,'4'是clb实例,'5'是eni实例,'6'是云数据库
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeAssociatedInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 列表偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页记录条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 地域代码（例：ap-guangzhou）,支持腾讯云全地域
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 额外检索条件（JSON字符串）
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 排序字段
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 排序方式（asc:升序,desc:降序）
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil,omitempty" name:"SecurityGroupId"`

	// 实例类型,'3'是cvm实例,'4'是clb实例,'5'是eni实例,'6'是云数据库
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeAssociatedInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssociatedInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Area")
	delete(f, "SearchValue")
	delete(f, "By")
	delete(f, "Order")
	delete(f, "SecurityGroupId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssociatedInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssociatedInstanceListResponseParams struct {
	// 实例数量
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 实例列表
	Data []*AssociatedInstanceInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAssociatedInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssociatedInstanceListResponseParams `json:"Response"`
}

func (r *DescribeAssociatedInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssociatedInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockByIpTimesListRequestParams struct {
	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// ip查询条件
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 方向
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// vpc间防火墙开关边id
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 日志来源 move：vpc间防火墙
	LogSource *string `json:"LogSource,omitnil,omitempty" name:"LogSource"`

	// 来源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 地域
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type DescribeBlockByIpTimesListRequest struct {
	*tchttp.BaseRequest
	
	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// ip查询条件
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 方向
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// vpc间防火墙开关边id
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 日志来源 move：vpc间防火墙
	LogSource *string `json:"LogSource,omitnil,omitempty" name:"LogSource"`

	// 来源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 地域
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

func (r *DescribeBlockByIpTimesListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockByIpTimesListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndTime")
	delete(f, "Ip")
	delete(f, "StartTime")
	delete(f, "Direction")
	delete(f, "EdgeId")
	delete(f, "LogSource")
	delete(f, "Source")
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlockByIpTimesListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockByIpTimesListResponseParams struct {
	// 返回数据
	Data []*IpStatic `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBlockByIpTimesListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBlockByIpTimesListResponseParams `json:"Response"`
}

func (r *DescribeBlockByIpTimesListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockByIpTimesListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockIgnoreListRequestParams struct {
	// 单页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 方向：1互联网入站，0互联网出站，3内网，空 全部方向
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 排序类型：desc降序，asc正序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序列：EndTime结束时间，StartTime开始时间，MatchTimes命中次数
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 搜索参数，json格式字符串，空则传"{}"，域名：domain，危险等级：level，放通原因：ignore_reason，安全事件来源：rule_source，地理位置：address，模糊搜索：common
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 规则类型：1封禁，2放通
	RuleType *uint64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// blocklist 封禁列表
	// whitelist 白名单列表
	ShowType *string `json:"ShowType,omitnil,omitempty" name:"ShowType"`
}

type DescribeBlockIgnoreListRequest struct {
	*tchttp.BaseRequest
	
	// 单页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 方向：1互联网入站，0互联网出站，3内网，空 全部方向
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 排序类型：desc降序，asc正序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序列：EndTime结束时间，StartTime开始时间，MatchTimes命中次数
	By *string `json:"By,omitnil,omitempty" name:"By"`

	// 搜索参数，json格式字符串，空则传"{}"，域名：domain，危险等级：level，放通原因：ignore_reason，安全事件来源：rule_source，地理位置：address，模糊搜索：common
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 规则类型：1封禁，2放通
	RuleType *uint64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// blocklist 封禁列表
	// whitelist 白名单列表
	ShowType *string `json:"ShowType,omitnil,omitempty" name:"ShowType"`
}

func (r *DescribeBlockIgnoreListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockIgnoreListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Direction")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "SearchValue")
	delete(f, "RuleType")
	delete(f, "ShowType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlockIgnoreListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockIgnoreListResponseParams struct {
	// 列表数据
	Data []*BlockIgnoreRule `json:"Data,omitnil,omitempty" name:"Data"`

	// 查询结果总数，用于分页
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 状态值，0：查询成功，非0：查询失败
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 状态信息，success：查询成功，fail：查询失败
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 安全事件来源下拉框
	SourceList []*string `json:"SourceList,omitnil,omitempty" name:"SourceList"`

	// 对应规则类型的数量，示例：[0,122,30,55,12,232,0]，封禁0个，IP地址122个，域名30个，威胁情报55个，资产实例12个，自定义策略232个，入侵防御规则0个
	RuleTypeDataList []*int64 `json:"RuleTypeDataList,omitnil,omitempty" name:"RuleTypeDataList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBlockIgnoreListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBlockIgnoreListResponseParams `json:"Response"`
}

func (r *DescribeBlockIgnoreListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockIgnoreListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockStaticListRequestParams struct {
	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 列表类型，只能是下面三种之一：port、address、ip
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// top数
	Top *int64 `json:"Top,omitnil,omitempty" name:"Top"`

	// 查询条件
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`
}

type DescribeBlockStaticListRequest struct {
	*tchttp.BaseRequest
	
	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 列表类型，只能是下面三种之一：port、address、ip
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// top数
	Top *int64 `json:"Top,omitnil,omitempty" name:"Top"`

	// 查询条件
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`
}

func (r *DescribeBlockStaticListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockStaticListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndTime")
	delete(f, "QueryType")
	delete(f, "StartTime")
	delete(f, "Top")
	delete(f, "SearchValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlockStaticListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockStaticListResponseParams struct {
	// 查询结果
	Data []*StaticInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 异步查询状态，1查询执行中，0查询已结束
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBlockStaticListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBlockStaticListResponseParams `json:"Response"`
}

func (r *DescribeBlockStaticListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockStaticListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnAssociatedInstancesRequestParams struct {
	// 云联网ID
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`
}

type DescribeCcnAssociatedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 云联网ID
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`
}

func (r *DescribeCcnAssociatedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnAssociatedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcnAssociatedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnAssociatedInstancesResponseParams struct {
	// 实例总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 云联网关联的实例信息
	CcnAssociatedInstances []*CcnAssociatedInstance `json:"CcnAssociatedInstances,omitnil,omitempty" name:"CcnAssociatedInstances"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCcnAssociatedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCcnAssociatedInstancesResponseParams `json:"Response"`
}

func (r *DescribeCcnAssociatedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnAssociatedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnInstanceRegionStatusRequestParams struct {
	// 云联网ID
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// 要查询引流网络部署状态的云联网关联的实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 引流路由方法 0:多路由表, 1:策略路由
	RoutingMode *uint64 `json:"RoutingMode,omitnil,omitempty" name:"RoutingMode"`
}

type DescribeCcnInstanceRegionStatusRequest struct {
	*tchttp.BaseRequest
	
	// 云联网ID
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// 要查询引流网络部署状态的云联网关联的实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 引流路由方法 0:多路由表, 1:策略路由
	RoutingMode *uint64 `json:"RoutingMode,omitnil,omitempty" name:"RoutingMode"`
}

func (r *DescribeCcnInstanceRegionStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnInstanceRegionStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	delete(f, "InstanceIds")
	delete(f, "RoutingMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcnInstanceRegionStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnInstanceRegionStatusResponseParams struct {
	// 地域总数量
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 地域防火墙引流网络状态列表
	RegionFwStatus []*RegionFwStatus `json:"RegionFwStatus,omitnil,omitempty" name:"RegionFwStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCcnInstanceRegionStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCcnInstanceRegionStatusResponseParams `json:"Response"`
}

func (r *DescribeCcnInstanceRegionStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnInstanceRegionStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnVpcFwPolicyLimitRequestParams struct {

}

type DescribeCcnVpcFwPolicyLimitRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCcnVpcFwPolicyLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnVpcFwPolicyLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcnVpcFwPolicyLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnVpcFwPolicyLimitResponseParams struct {
	// 支持的引流策略数量（最外层总条数）
	CcnPolicyInterconnectPairLenLimit *uint64 `json:"CcnPolicyInterconnectPairLenLimit,omitnil,omitempty" name:"CcnPolicyInterconnectPairLenLimit"`

	// 单条引流策略中单组的最大配置数量（内层单组总条数）
	CcnPolicyGroupLenLimit *uint64 `json:"CcnPolicyGroupLenLimit,omitnil,omitempty" name:"CcnPolicyGroupLenLimit"`

	// 接入的实例网段长度（网段数量）限制
	CcnPolicyCidrLenLimit *uint64 `json:"CcnPolicyCidrLenLimit,omitnil,omitempty" name:"CcnPolicyCidrLenLimit"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCcnVpcFwPolicyLimitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCcnVpcFwPolicyLimitResponseParams `json:"Response"`
}

func (r *DescribeCcnVpcFwPolicyLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnVpcFwPolicyLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnVpcFwSwitchRequestParams struct {
	// 云联网ID
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`
}

type DescribeCcnVpcFwSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 云联网ID
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`
}

func (r *DescribeCcnVpcFwSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnVpcFwSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCcnVpcFwSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCcnVpcFwSwitchResponseParams struct {
	// 互联对配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	InterconnectPairs []*InterconnectPair `json:"InterconnectPairs,omitnil,omitempty" name:"InterconnectPairs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCcnVpcFwSwitchResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCcnVpcFwSwitchResponseParams `json:"Response"`
}

func (r *DescribeCcnVpcFwSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCcnVpcFwSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfwEipsRequestParams struct {
	// 1：cfw接入模式，目前仅支持接入模式实例
	Mode *uint64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// ALL：查询所有弹性公网ip; nat-xxxxx：接入模式场景指定网关的弹性公网ip
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// 防火墙实例id，当前仅支持接入模式的实例，该字段必填
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`
}

type DescribeCfwEipsRequest struct {
	*tchttp.BaseRequest
	
	// 1：cfw接入模式，目前仅支持接入模式实例
	Mode *uint64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// ALL：查询所有弹性公网ip; nat-xxxxx：接入模式场景指定网关的弹性公网ip
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// 防火墙实例id，当前仅支持接入模式的实例，该字段必填
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`
}

func (r *DescribeCfwEipsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfwEipsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	delete(f, "NatGatewayId")
	delete(f, "CfwInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfwEipsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfwEipsResponseParams struct {
	// 返回值信息
	NatFwEipList []*NatFwEipsInfo `json:"NatFwEipList,omitnil,omitempty" name:"NatFwEipList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCfwEipsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCfwEipsResponseParams `json:"Response"`
}

func (r *DescribeCfwEipsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfwEipsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfwInsStatusRequestParams struct {

}

type DescribeCfwInsStatusRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCfwInsStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfwInsStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCfwInsStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCfwInsStatusResponseParams struct {
	// 防火墙实例运行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CfwInsStatus []*CfwInsStatus `json:"CfwInsStatus,omitnil,omitempty" name:"CfwInsStatus"`

	// 0
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCfwInsStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCfwInsStatusResponseParams `json:"Response"`
}

func (r *DescribeCfwInsStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfwInsStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterVpcFwSwitchsRequestParams struct {
	// 需要查询的索引，特定场景使用，可不填
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeClusterVpcFwSwitchsRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的索引，特定场景使用，可不填
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeClusterVpcFwSwitchsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterVpcFwSwitchsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Index")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterVpcFwSwitchsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterVpcFwSwitchsResponseParams struct {
	// 总条数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 防火墙开关列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ClusterSwitchDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterVpcFwSwitchsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterVpcFwSwitchsResponseParams `json:"Response"`
}

func (r *DescribeClusterVpcFwSwitchsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterVpcFwSwitchsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefenseSwitchRequestParams struct {

}

type DescribeDefenseSwitchRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDefenseSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefenseSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDefenseSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDefenseSwitchResponseParams struct {
	// 基础防御开关
	BasicRuleSwitch *int64 `json:"BasicRuleSwitch,omitnil,omitempty" name:"BasicRuleSwitch"`

	// 安全基线开关
	BaselineAllSwitch *int64 `json:"BaselineAllSwitch,omitnil,omitempty" name:"BaselineAllSwitch"`

	// 威胁情报开关
	TiSwitch *int64 `json:"TiSwitch,omitnil,omitempty" name:"TiSwitch"`

	// 虚拟补丁开关
	VirtualPatchSwitch *int64 `json:"VirtualPatchSwitch,omitnil,omitempty" name:"VirtualPatchSwitch"`

	// 是否历史开启
	HistoryOpen *int64 `json:"HistoryOpen,omitnil,omitempty" name:"HistoryOpen"`

	// 状态值，0：查询成功，非0：查询失败
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 状态信息，success：查询成功，fail：查询失败
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDefenseSwitchResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDefenseSwitchResponseParams `json:"Response"`
}

func (r *DescribeDefenseSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDefenseSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnterpriseSGRuleProgressRequestParams struct {

}

type DescribeEnterpriseSGRuleProgressRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeEnterpriseSGRuleProgressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnterpriseSGRuleProgressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnterpriseSGRuleProgressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnterpriseSGRuleProgressResponseParams struct {
	// 0-100，代表下发进度百分比
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 是否用户中止 用户中止返回true
	UserStopped *bool `json:"UserStopped,omitnil,omitempty" name:"UserStopped"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEnterpriseSGRuleProgressResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnterpriseSGRuleProgressResponseParams `json:"Response"`
}

func (r *DescribeEnterpriseSGRuleProgressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnterpriseSGRuleProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnterpriseSecurityGroupRuleListRequestParams struct {
	// 分页每页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页当前页
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 启用状态 1启用 0 未启用
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 地域
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 规则下发方式筛选  1 新规则和延迟下发  2  仅看新规则  
	Filter *int64 `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询条件
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 查询条件新
	SearchFilters []*CommonFilter `json:"SearchFilters,omitnil,omitempty" name:"SearchFilters"`
}

type DescribeEnterpriseSecurityGroupRuleListRequest struct {
	*tchttp.BaseRequest
	
	// 分页每页数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页当前页
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 启用状态 1启用 0 未启用
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 地域
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 规则下发方式筛选  1 新规则和延迟下发  2  仅看新规则  
	Filter *int64 `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 查询条件
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 查询条件新
	SearchFilters []*CommonFilter `json:"SearchFilters,omitnil,omitempty" name:"SearchFilters"`
}

func (r *DescribeEnterpriseSecurityGroupRuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnterpriseSecurityGroupRuleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Status")
	delete(f, "Area")
	delete(f, "Filter")
	delete(f, "SearchValue")
	delete(f, "SearchFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnterpriseSecurityGroupRuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnterpriseSecurityGroupRuleListResponseParams struct {
	// 查询结果总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 规则总数
	AllTotal *int64 `json:"AllTotal,omitnil,omitempty" name:"AllTotal"`

	// 规则列表
	Data []*EnterpriseSecurityGroupRuleRuleInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 规则列表整体启用状态 
	// 取值范围 0/1/2
	// 0.表示没有启用的(可以批量启用)  
	// 1.表示没有禁用的(可以批量禁用)    
	// 2 表示混合情况（不可批量操作）
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEnterpriseSecurityGroupRuleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnterpriseSecurityGroupRuleListResponseParams `json:"Response"`
}

func (r *DescribeEnterpriseSecurityGroupRuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnterpriseSecurityGroupRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnterpriseSecurityGroupRuleRequestParams struct {
	// 分页查询时，显示的当前页的页码。
	// 
	// 默认值为1。
	PageNo *string `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 分页查询时，显示的每页数据的最大条数。
	// 
	// 可设置值最大为50。
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 访问源示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板(ipm-dyodhpby)
	// instance：资产实例(ins-123456)
	// resourcegroup：资产分组(/全部分组/分组1/子分组1)
	// tag：资源标签({"Key":"标签key值","Value":"标签Value值"})
	// region：地域(ap-gaungzhou)
	// 支持通配
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// 访问目的示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板(ipm-dyodhpby)
	// instance：资产实例(ins-123456)
	// resourcegroup：资产分组(/全部分组/分组1/子分组1)
	// tag：资源标签({"Key":"标签key值","Value":"标签Value值"})
	// region：地域(ap-gaungzhou)
	// 支持通配
	DestContent *string `json:"DestContent,omitnil,omitempty" name:"DestContent"`

	// 规则描述，支持通配
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 访问控制策略中设置的流量通过云防火墙的方式。取值：
	// accept：放行
	// drop：拒绝
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// 是否启用规则，默认为启用，取值：
	// true为启用，false为不启用
	Enable *string `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 访问控制策略的端口。取值：
	// -1/-1：全部端口
	// 80：80端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 协议；TCP/UDP/ICMP/ANY
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 端口协议类型参数模板id；协议端口模板id
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`

	// 规则的uuid
	RuleUuid *int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`
}

type DescribeEnterpriseSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询时，显示的当前页的页码。
	// 
	// 默认值为1。
	PageNo *string `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 分页查询时，显示的每页数据的最大条数。
	// 
	// 可设置值最大为50。
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 访问源示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板(ipm-dyodhpby)
	// instance：资产实例(ins-123456)
	// resourcegroup：资产分组(/全部分组/分组1/子分组1)
	// tag：资源标签({"Key":"标签key值","Value":"标签Value值"})
	// region：地域(ap-gaungzhou)
	// 支持通配
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// 访问目的示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板(ipm-dyodhpby)
	// instance：资产实例(ins-123456)
	// resourcegroup：资产分组(/全部分组/分组1/子分组1)
	// tag：资源标签({"Key":"标签key值","Value":"标签Value值"})
	// region：地域(ap-gaungzhou)
	// 支持通配
	DestContent *string `json:"DestContent,omitnil,omitempty" name:"DestContent"`

	// 规则描述，支持通配
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 访问控制策略中设置的流量通过云防火墙的方式。取值：
	// accept：放行
	// drop：拒绝
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// 是否启用规则，默认为启用，取值：
	// true为启用，false为不启用
	Enable *string `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 访问控制策略的端口。取值：
	// -1/-1：全部端口
	// 80：80端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 协议；TCP/UDP/ICMP/ANY
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 端口协议类型参数模板id；协议端口模板id
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`

	// 规则的uuid
	RuleUuid *int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`
}

func (r *DescribeEnterpriseSecurityGroupRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnterpriseSecurityGroupRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "SourceContent")
	delete(f, "DestContent")
	delete(f, "Description")
	delete(f, "RuleAction")
	delete(f, "Enable")
	delete(f, "Port")
	delete(f, "Protocol")
	delete(f, "ServiceTemplateId")
	delete(f, "RuleUuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnterpriseSecurityGroupRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnterpriseSecurityGroupRuleResponseParams struct {
	// 分页查询时，显示的当前页的页码。
	PageNo *string `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 分页查询时，显示的每页数据的最大条数。
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 访问控制策略列表
	Rules []*SecurityGroupRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 访问控制策略的总数量。
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEnterpriseSecurityGroupRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnterpriseSecurityGroupRuleResponseParams `json:"Response"`
}

func (r *DescribeEnterpriseSecurityGroupRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnterpriseSecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFwEdgeIpsRequestParams struct {
	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 每页条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeFwEdgeIpsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 每页条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeFwEdgeIpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFwEdgeIpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFwEdgeIpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFwEdgeIpsResponseParams struct {
	// ip 开关列表
	Data []*EdgeIpInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// ip 开关列表个数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 地域列表
	RegionLst []*string `json:"RegionLst,omitnil,omitempty" name:"RegionLst"`

	// 实例类型列表
	InstanceTypeLst []*string `json:"InstanceTypeLst,omitnil,omitempty" name:"InstanceTypeLst"`

	// 串行模式开关个数
	SerilCount *int64 `json:"SerilCount,omitnil,omitempty" name:"SerilCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFwEdgeIpsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFwEdgeIpsResponseParams `json:"Response"`
}

func (r *DescribeFwEdgeIpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFwEdgeIpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFwGroupInstanceInfoRequestParams struct {
	// 每页条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeFwGroupInstanceInfoRequest struct {
	*tchttp.BaseRequest
	
	// 每页条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeFwGroupInstanceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFwGroupInstanceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFwGroupInstanceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFwGroupInstanceInfoResponseParams struct {
	// 防火墙(组)详细信息
	VpcFwGroupLst []*VpcFwGroupInfo `json:"VpcFwGroupLst,omitnil,omitempty" name:"VpcFwGroupLst"`

	// 防火墙(组)个数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFwGroupInstanceInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFwGroupInstanceInfoResponseParams `json:"Response"`
}

func (r *DescribeFwGroupInstanceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFwGroupInstanceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFwSyncStatusRequestParams struct {
	// 查询的同步状态类型：Route,同步路由状态
	SyncType *string `json:"SyncType,omitnil,omitempty" name:"SyncType"`
}

type DescribeFwSyncStatusRequest struct {
	*tchttp.BaseRequest
	
	// 查询的同步状态类型：Route,同步路由状态
	SyncType *string `json:"SyncType,omitnil,omitempty" name:"SyncType"`
}

func (r *DescribeFwSyncStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFwSyncStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SyncType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFwSyncStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFwSyncStatusResponseParams struct {
	// 同步状态：1，同步中；0，同步完成
	SyncStatus *int64 `json:"SyncStatus,omitnil,omitempty" name:"SyncStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFwSyncStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFwSyncStatusResponseParams `json:"Response"`
}

func (r *DescribeFwSyncStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFwSyncStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGuideScanInfoRequestParams struct {

}

type DescribeGuideScanInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeGuideScanInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGuideScanInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGuideScanInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGuideScanInfoResponseParams struct {
	// 扫描信息
	Data *ScanInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGuideScanInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGuideScanInfoResponseParams `json:"Response"`
}

func (r *DescribeGuideScanInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGuideScanInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPStatusListRequestParams struct {
	// 资产Id
	IPList []*string `json:"IPList,omitnil,omitempty" name:"IPList"`
}

type DescribeIPStatusListRequest struct {
	*tchttp.BaseRequest
	
	// 资产Id
	IPList []*string `json:"IPList,omitnil,omitempty" name:"IPList"`
}

func (r *DescribeIPStatusListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPStatusListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IPList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIPStatusListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIPStatusListResponseParams struct {
	// IP状态信息
	StatusList []*IPDefendStatus `json:"StatusList,omitnil,omitempty" name:"StatusList"`

	// 状态码
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 状态信息
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIPStatusListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIPStatusListResponseParams `json:"Response"`
}

func (r *DescribeIPStatusListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIPStatusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpsModeSwitchRequestParams struct {

}

type DescribeIpsModeSwitchRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeIpsModeSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpsModeSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpsModeSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpsModeSwitchResponseParams struct {
	// // Mode 取值校验：0-观察模式, 1-拦截模式, 2-严格模式
	Data *ModeInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 0 成功 非0失败
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// success 成功 其他失败
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIpsModeSwitchResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpsModeSwitchResponseParams `json:"Response"`
}

func (r *DescribeIpsModeSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpsModeSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogStorageStatisticRequestParams struct {

}

type DescribeLogStorageStatisticRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLogStorageStatisticRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogStorageStatisticRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogStorageStatisticRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogStorageStatisticResponseParams struct {
	// 返回状态码 0 成功 非0不成功
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 返回信息  success 成功 其他 不成功
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 已使用存储量，单位B
	UsedSize *int64 `json:"UsedSize,omitnil,omitempty" name:"UsedSize"`

	// 配额存储总量，单位B
	TotalSize *int64 `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`

	// 存储天数
	StorageDay *int64 `json:"StorageDay,omitnil,omitempty" name:"StorageDay"`

	// 访问控制日志存储量，单位B
	AclSize *int64 `json:"AclSize,omitnil,omitempty" name:"AclSize"`

	// 入侵防御日志存储量，单位B
	IdsSize *int64 `json:"IdsSize,omitnil,omitempty" name:"IdsSize"`

	// 流量日志存储量，单位B
	NetFlowSize *int64 `json:"NetFlowSize,omitnil,omitempty" name:"NetFlowSize"`

	// 操作日志存储量，单位B
	OperateSize *int64 `json:"OperateSize,omitnil,omitempty" name:"OperateSize"`

	// 剩余存储量，单位B
	LeftSize *int64 `json:"LeftSize,omitnil,omitempty" name:"LeftSize"`

	// 计费模式，0后付费，1预付费
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 每日增加日志存储量柱状图
	TimeHistogram []*StorageHistogram `json:"TimeHistogram,omitnil,omitempty" name:"TimeHistogram"`

	// 柱形图格式数据
	TimeHistogramShow *StorageHistogramShow `json:"TimeHistogramShow,omitnil,omitempty" name:"TimeHistogramShow"`

	// 后付费模式存储状态，0正常，1欠费停止写入
	ArrearsStopWriting *int64 `json:"ArrearsStopWriting,omitnil,omitempty" name:"ArrearsStopWriting"`

	// NDR流量日志存储量，单位B
	NDRNetFlowSize *int64 `json:"NDRNetFlowSize,omitnil,omitempty" name:"NDRNetFlowSize"`

	// NDR风险日志存储量，单位B
	NDRRiskSize *int64 `json:"NDRRiskSize,omitnil,omitempty" name:"NDRRiskSize"`

	// NDR日志存储天数
	NDRStorageDay *int64 `json:"NDRStorageDay,omitnil,omitempty" name:"NDRStorageDay"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogStorageStatisticResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogStorageStatisticResponseParams `json:"Response"`
}

func (r *DescribeLogStorageStatisticResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogStorageStatisticResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogsRequestParams struct {
	// 日志类型标识
	// 流量日志：互联网边界防火墙netflow_border，NAT边界防火墙netflow_nat，VPC间防火墙vpcnetflow，内网流量日志netflow_fl，流量分析日志netflow_nta
	// 入侵防御日志rule_threatinfo
	// 访问控制日志：互联网边界规则rule_acl，NAT边界规则rule_acl，内网间规则rule_vpcacl，企业安全组rule_sg
	// 操作日志：防火墙开关-开关操作operate_switch，防火墙开关-实例配置operate_instance，资产中心操作operate_assetgroup，访问控制操作operate_acl，零信任防护操作operate_identity，入侵防御操作-入侵防御operate_ids，入侵防御操作-安全基线operate_baseline，常用工具操作operate_tool，网络蜜罐操作operate_honeypot，日志投递操作operate_logdelivery，通用设置操作operate_logstorage，登录日志operate_login
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 每页条数，最大支持1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值，最大支持60000
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 筛选开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 筛选结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 过滤条件组合，各数组元素间为AND关系，查询字段名Name参考文档https://cloud.tencent.com/document/product/1132/87894，数值类型字段不支持模糊匹配
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeLogsRequest struct {
	*tchttp.BaseRequest
	
	// 日志类型标识
	// 流量日志：互联网边界防火墙netflow_border，NAT边界防火墙netflow_nat，VPC间防火墙vpcnetflow，内网流量日志netflow_fl，流量分析日志netflow_nta
	// 入侵防御日志rule_threatinfo
	// 访问控制日志：互联网边界规则rule_acl，NAT边界规则rule_acl，内网间规则rule_vpcacl，企业安全组rule_sg
	// 操作日志：防火墙开关-开关操作operate_switch，防火墙开关-实例配置operate_instance，资产中心操作operate_assetgroup，访问控制操作operate_acl，零信任防护操作operate_identity，入侵防御操作-入侵防御operate_ids，入侵防御操作-安全基线operate_baseline，常用工具操作operate_tool，网络蜜罐操作operate_honeypot，日志投递操作operate_logdelivery，通用设置操作operate_logstorage，登录日志operate_login
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 每页条数，最大支持1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值，最大支持60000
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 筛选开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 筛选结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 过滤条件组合，各数组元素间为AND关系，查询字段名Name参考文档https://cloud.tencent.com/document/product/1132/87894，数值类型字段不支持模糊匹配
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Index")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLogsResponseParams struct {
	// 日志列表
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 总条数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 返回状态码 0 成功 非0不成功
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 返回信息  success 成功 其他 不成功
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 七层协议，NTA日志有效
	AppProtocolList []*string `json:"AppProtocolList,omitnil,omitempty" name:"AppProtocolList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLogsResponseParams `json:"Response"`
}

func (r *DescribeLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatAcRuleRequestParams struct {
	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 需要查询的索引，特定场景使用，可不填
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 过滤条件组合，Direction 为0时表述查询出向规则，为1时表示查询入向规则
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值，默认为asc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序所用到的字段，默认为sequence
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeNatAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 需要查询的索引，特定场景使用，可不填
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 过滤条件组合，Direction 为0时表述查询出向规则，为1时表示查询入向规则
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值，默认为asc
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序所用到的字段，默认为sequence
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeNatAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Index")
	delete(f, "Filters")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatAcRuleResponseParams struct {
	// 总条数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// nat访问控制列表数据
	Data []*DescAcItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 未过滤的总条数
	AllTotal *uint64 `json:"AllTotal,omitnil,omitempty" name:"AllTotal"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNatAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatAcRuleResponseParams `json:"Response"`
}

func (r *DescribeNatAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwDnatRuleRequestParams struct {
	// 需要查询的索引，特定场景使用，可不填
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值，可不传
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序所用到的字段，可不传
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeNatFwDnatRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的索引，特定场景使用，可不填
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值，可不传
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序所用到的字段，可不传
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeNatFwDnatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwDnatRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Index")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatFwDnatRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwDnatRuleResponseParams struct {
	// Dnat规则列表
	Data []*DescNatDnatRule `json:"Data,omitnil,omitempty" name:"Data"`

	// 列表总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNatFwDnatRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatFwDnatRuleResponseParams `json:"Response"`
}

func (r *DescribeNatFwDnatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwDnatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwInfoCountRequestParams struct {

}

type DescribeNatFwInfoCountRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeNatFwInfoCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInfoCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatFwInfoCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwInfoCountResponseParams struct {
	// 返回参数 success 成功 failed 失败
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 当前租户的nat防火墙实例个数
	NatFwInsCount *int64 `json:"NatFwInsCount,omitnil,omitempty" name:"NatFwInsCount"`

	// 当前租户接入防火墙的子网个数
	SubnetCount *int64 `json:"SubnetCount,omitnil,omitempty" name:"SubnetCount"`

	// 打开NAT防火墙开关个数
	OpenSwitchCount *int64 `json:"OpenSwitchCount,omitnil,omitempty" name:"OpenSwitchCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNatFwInfoCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatFwInfoCountResponseParams `json:"Response"`
}

func (r *DescribeNatFwInfoCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInfoCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwInstanceRequestParams struct {

}

type DescribeNatFwInstanceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeNatFwInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatFwInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwInstanceResponseParams struct {
	// 实例数组
	NatinsLst []*NatFwInstance `json:"NatinsLst,omitnil,omitempty" name:"NatinsLst"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNatFwInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatFwInstanceResponseParams `json:"Response"`
}

func (r *DescribeNatFwInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwInstanceWithRegionRequestParams struct {

}

type DescribeNatFwInstanceWithRegionRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeNatFwInstanceWithRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInstanceWithRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatFwInstanceWithRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwInstanceWithRegionResponseParams struct {
	// 实例数组
	NatinsLst []*NatFwInstance `json:"NatinsLst,omitnil,omitempty" name:"NatinsLst"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNatFwInstanceWithRegionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatFwInstanceWithRegionResponseParams `json:"Response"`
}

func (r *DescribeNatFwInstanceWithRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInstanceWithRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwInstancesInfoRequestParams struct {
	// 获取实例列表过滤字段
	Filter []*NatFwFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 第几页
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页长度
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeNatFwInstancesInfoRequest struct {
	*tchttp.BaseRequest
	
	// 获取实例列表过滤字段
	Filter []*NatFwFilter `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 第几页
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页长度
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeNatFwInstancesInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInstancesInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filter")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatFwInstancesInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwInstancesInfoResponseParams struct {
	// 实例卡片信息数组
	NatinsLst []*NatInstanceInfo `json:"NatinsLst,omitnil,omitempty" name:"NatinsLst"`

	// nat 防火墙个数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNatFwInstancesInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatFwInstancesInfoResponseParams `json:"Response"`
}

func (r *DescribeNatFwInstancesInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwInstancesInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwSwitchRequestParams struct {
	// 偏移量，分页用
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 条数，分页用
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeNatFwSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，分页用
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 条数，分页用
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeNatFwSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatFwSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwSwitchResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// NAT边界防火墙开关列表数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*NatSwitchListData `json:"Data,omitnil,omitempty" name:"Data"`

	// 开关相关VPC列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcList []*CommonIdName `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// 开关相关NAT列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatList []*CommonIdName `json:"NatList,omitnil,omitempty" name:"NatList"`

	// 开关相关ROUTE列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteList []*CommonIdName `json:"RouteList,omitnil,omitempty" name:"RouteList"`

	// 开启开关个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OnNum *int64 `json:"OnNum,omitnil,omitempty" name:"OnNum"`

	// 关闭开关个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffNum *int64 `json:"OffNum,omitnil,omitempty" name:"OffNum"`

	// 失败开关状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailData []*CommonIdNameStatus `json:"FailData,omitnil,omitempty" name:"FailData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNatFwSwitchResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatFwSwitchResponseParams `json:"Response"`
}

func (r *DescribeNatFwSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwVpcDnsLstRequestParams struct {
	// natfw 防火墙实例id
	NatFwInsId *string `json:"NatFwInsId,omitnil,omitempty" name:"NatFwInsId"`

	// natfw 过滤，以','分隔
	NatInsIdFilter *string `json:"NatInsIdFilter,omitnil,omitempty" name:"NatInsIdFilter"`

	// 分页页数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页最多个数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeNatFwVpcDnsLstRequest struct {
	*tchttp.BaseRequest
	
	// natfw 防火墙实例id
	NatFwInsId *string `json:"NatFwInsId,omitnil,omitempty" name:"NatFwInsId"`

	// natfw 过滤，以','分隔
	NatInsIdFilter *string `json:"NatInsIdFilter,omitnil,omitempty" name:"NatInsIdFilter"`

	// 分页页数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页最多个数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeNatFwVpcDnsLstRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwVpcDnsLstRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatFwInsId")
	delete(f, "NatInsIdFilter")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatFwVpcDnsLstRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatFwVpcDnsLstResponseParams struct {
	// nat防火墙vpc dns 信息数组
	VpcDnsSwitchLst []*VpcDnsInfo `json:"VpcDnsSwitchLst,omitnil,omitempty" name:"VpcDnsSwitchLst"`

	// 返回参数 success成功 failed 失败
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 开关总条数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNatFwVpcDnsLstResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatFwVpcDnsLstResponseParams `json:"Response"`
}

func (r *DescribeNatFwVpcDnsLstResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatFwVpcDnsLstResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceGroupNewRequestParams struct {
	// 查询类型 网络结构-vpc，业务识别-resource ，资源标签-tag
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 资产组id  全部传0
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// all  包含子组 own自己
	ShowType *string `json:"ShowType,omitnil,omitempty" name:"ShowType"`
}

type DescribeResourceGroupNewRequest struct {
	*tchttp.BaseRequest
	
	// 查询类型 网络结构-vpc，业务识别-resource ，资源标签-tag
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 资产组id  全部传0
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// all  包含子组 own自己
	ShowType *string `json:"ShowType,omitnil,omitempty" name:"ShowType"`
}

func (r *DescribeResourceGroupNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceGroupNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueryType")
	delete(f, "GroupId")
	delete(f, "ShowType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceGroupNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceGroupNewResponseParams struct {
	// 返回树形结构
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 返回码；0为请求成功
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 接口返回消息
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 未分类实例数量
	UnResourceNum *int64 `json:"UnResourceNum,omitnil,omitempty" name:"UnResourceNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourceGroupNewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceGroupNewResponseParams `json:"Response"`
}

func (r *DescribeResourceGroupNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceGroupNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceGroupRequestParams struct {
	// 查询类型 网络结构 vpc，业务识别- resource ，资源标签-tag
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 资产组id  全部传0
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// all  包含子组 own自己
	ShowType *string `json:"ShowType,omitnil,omitempty" name:"ShowType"`
}

type DescribeResourceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 查询类型 网络结构 vpc，业务识别- resource ，资源标签-tag
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 资产组id  全部传0
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// all  包含子组 own自己
	ShowType *string `json:"ShowType,omitnil,omitempty" name:"ShowType"`
}

func (r *DescribeResourceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueryType")
	delete(f, "GroupId")
	delete(f, "ShowType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceGroupResponseParams struct {
	// 返回树形结构
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourceGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceGroupResponseParams `json:"Response"`
}

func (r *DescribeResourceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleOverviewRequestParams struct {
	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type DescribeRuleOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

func (r *DescribeRuleOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleOverviewResponseParams struct {
	// 规则总数
	AllTotal *uint64 `json:"AllTotal,omitnil,omitempty" name:"AllTotal"`

	// 阻断策略规则数量
	StrategyNum *uint64 `json:"StrategyNum,omitnil,omitempty" name:"StrategyNum"`

	// 启用规则数量
	StartRuleNum *uint64 `json:"StartRuleNum,omitnil,omitempty" name:"StartRuleNum"`

	// 停用规则数量
	StopRuleNum *uint64 `json:"StopRuleNum,omitnil,omitempty" name:"StopRuleNum"`

	// 剩余配额
	RemainingNum *int64 `json:"RemainingNum,omitnil,omitempty" name:"RemainingNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRuleOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleOverviewResponseParams `json:"Response"`
}

func (r *DescribeRuleOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupListRequestParams struct {
	// 0: 出站规则，1：入站规则
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 地域代码（例: ap-guangzhou),支持腾讯云全部地域
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 每页条数，默认为10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 状态，'': 全部，'0'：筛选停用规则，'1'：筛选启用规则
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 0: 不过滤，1：过滤掉正常规则，保留下发异常规则
	Filter *uint64 `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeSecurityGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 0: 出站规则，1：入站规则
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 地域代码（例: ap-guangzhou),支持腾讯云全部地域
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 每页条数，默认为10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值，默认为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 状态，'': 全部，'0'：筛选停用规则，'1'：筛选启用规则
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 0: 不过滤，1：过滤掉正常规则，保留下发异常规则
	Filter *uint64 `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeSecurityGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "Area")
	delete(f, "SearchValue")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Status")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityGroupListResponseParams struct {
	// 列表当前规则总条数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 安全组规则列表数据
	Data []*SecurityGroupListData `json:"Data,omitnil,omitempty" name:"Data"`

	// 不算筛选条数的总条数
	AllTotal *uint64 `json:"AllTotal,omitnil,omitempty" name:"AllTotal"`

	// 访问控制规则全部启用/全部停用
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityGroupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityGroupListResponseParams `json:"Response"`
}

func (r *DescribeSecurityGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSourceAssetRequestParams struct {
	// ChooseType为1，查询已经分组的资产；ChooseType不为1查询没有分组的资产
	ChooseType *string `json:"ChooseType,omitnil,omitempty" name:"ChooseType"`

	// 模糊查询
	FuzzySearch *string `json:"FuzzySearch,omitnil,omitempty" name:"FuzzySearch"`

	// 资产类型 1公网 2内网
	InsType *string `json:"InsType,omitnil,omitempty" name:"InsType"`

	// 查询单页的最大值；eg：10；则最多返回10条结果
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询结果的偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 地域
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type DescribeSourceAssetRequest struct {
	*tchttp.BaseRequest
	
	// ChooseType为1，查询已经分组的资产；ChooseType不为1查询没有分组的资产
	ChooseType *string `json:"ChooseType,omitnil,omitempty" name:"ChooseType"`

	// 模糊查询
	FuzzySearch *string `json:"FuzzySearch,omitnil,omitempty" name:"FuzzySearch"`

	// 资产类型 1公网 2内网
	InsType *string `json:"InsType,omitnil,omitempty" name:"InsType"`

	// 查询单页的最大值；eg：10；则最多返回10条结果
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询结果的偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 地域
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

func (r *DescribeSourceAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChooseType")
	delete(f, "FuzzySearch")
	delete(f, "InsType")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Zone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSourceAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSourceAssetResponseParams struct {
	// 数据
	Data []*InstanceInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 返回数据总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 地域集合
	ZoneList []*AssetZone `json:"ZoneList,omitnil,omitempty" name:"ZoneList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSourceAssetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSourceAssetResponseParams `json:"Response"`
}

func (r *DescribeSourceAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSourceAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSwitchErrorRequestParams struct {
	// EDGE_FW : 互联网边界防火墙 , NDR: 流量分析，VPC_FW：VPC边界防火墙
	FwType *string `json:"FwType,omitnil,omitempty" name:"FwType"`
}

type DescribeSwitchErrorRequest struct {
	*tchttp.BaseRequest
	
	// EDGE_FW : 互联网边界防火墙 , NDR: 流量分析，VPC_FW：VPC边界防火墙
	FwType *string `json:"FwType,omitnil,omitempty" name:"FwType"`
}

func (r *DescribeSwitchErrorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSwitchErrorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FwType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSwitchErrorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSwitchErrorResponseParams struct {
	// 错误信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*SwitchError `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSwitchErrorResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSwitchErrorResponseParams `json:"Response"`
}

func (r *DescribeSwitchErrorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSwitchErrorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSwitchListsRequestParams struct {
	// 防火墙状态  0: 关闭，1：开启
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 资产类型 CVM/NAT/VPN/CLB/其它
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 地域 上海/重庆/广州，等等
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 搜索值  例子："{"common":"106.54.189.45"}"
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 条数  默认值:10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值 默认值: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序，desc：降序，asc：升序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段 PortTimes(风险端口数)
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeSwitchListsRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙状态  0: 关闭，1：开启
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 资产类型 CVM/NAT/VPN/CLB/其它
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 地域 上海/重庆/广州，等等
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 搜索值  例子："{"common":"106.54.189.45"}"
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 条数  默认值:10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值 默认值: 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序，desc：降序，asc：升序
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序字段 PortTimes(风险端口数)
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeSwitchListsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSwitchListsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "Type")
	delete(f, "Area")
	delete(f, "SearchValue")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSwitchListsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSwitchListsResponseParams struct {
	// 总条数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 列表数据
	Data []*SwitchListsData `json:"Data,omitnil,omitempty" name:"Data"`

	// 区域列表
	AreaLists []*string `json:"AreaLists,omitnil,omitempty" name:"AreaLists"`

	// 打开个数
	OnNum *uint64 `json:"OnNum,omitnil,omitempty" name:"OnNum"`

	// 关闭个数
	OffNum *uint64 `json:"OffNum,omitnil,omitempty" name:"OffNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSwitchListsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSwitchListsResponseParams `json:"Response"`
}

func (r *DescribeSwitchListsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSwitchListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTLogInfoRequestParams struct {
	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 类型 1 告警 2阻断
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询条件
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`
}

type DescribeTLogInfoRequest struct {
	*tchttp.BaseRequest
	
	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 类型 1 告警 2阻断
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询条件
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`
}

func (r *DescribeTLogInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTLogInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndTime")
	delete(f, "QueryType")
	delete(f, "StartTime")
	delete(f, "SearchValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTLogInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTLogInfoResponseParams struct {
	// "NetworkNum":网络扫描探测
	//  "HandleNum": 待处理事件
	// "BanNum": 
	//   "VulNum": 漏洞利用
	//   "OutNum": 失陷主机
	// "BruteForceNum": 0
	Data *TLogInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTLogInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTLogInfoResponseParams `json:"Response"`
}

func (r *DescribeTLogInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTLogInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTLogIpListRequestParams struct {
	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 类型 1 告警 2阻断
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// top数
	Top *int64 `json:"Top,omitnil,omitempty" name:"Top"`

	// 查询条件
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`
}

type DescribeTLogIpListRequest struct {
	*tchttp.BaseRequest
	
	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 类型 1 告警 2阻断
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// top数
	Top *int64 `json:"Top,omitnil,omitempty" name:"Top"`

	// 查询条件
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`
}

func (r *DescribeTLogIpListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTLogIpListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndTime")
	delete(f, "QueryType")
	delete(f, "StartTime")
	delete(f, "Top")
	delete(f, "SearchValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTLogIpListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTLogIpListResponseParams struct {
	// 数据集合
	Data []*StaticInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTLogIpListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTLogIpListResponseParams `json:"Response"`
}

func (r *DescribeTLogIpListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTLogIpListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableStatusRequestParams struct {
	// EdgeId值两个vpc间的边id vpc填Edgeid，不要填Area；
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 状态值，0：检查表的状态 确实只有一个默认值
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Nat所在地域 NAT填Area，不要填Edgeid；
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 方向，0：出站，1：入站 默认值为 0
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type DescribeTableStatusRequest struct {
	*tchttp.BaseRequest
	
	// EdgeId值两个vpc间的边id vpc填Edgeid，不要填Area；
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 状态值，0：检查表的状态 确实只有一个默认值
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Nat所在地域 NAT填Area，不要填Edgeid；
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 方向，0：出站，1：入站 默认值为 0
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

func (r *DescribeTableStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeId")
	delete(f, "Status")
	delete(f, "Area")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableStatusResponseParams struct {
	// 0：正常，其它：不正常
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTableStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableStatusResponseParams `json:"Response"`
}

func (r *DescribeTableStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnHandleEventTabListRequestParams struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询示例ID
	AssetID *string `json:"AssetID,omitnil,omitempty" name:"AssetID"`
}

type DescribeUnHandleEventTabListRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询示例ID
	AssetID *string `json:"AssetID,omitnil,omitempty" name:"AssetID"`
}

func (r *DescribeUnHandleEventTabListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnHandleEventTabListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AssetID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUnHandleEventTabListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnHandleEventTabListResponseParams struct {
	// 租户伪攻击链未处置事件
	Data *UnHandleEvent `json:"Data,omitnil,omitempty" name:"Data"`

	// 错误码，0成功 非0错误
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 返回信息 success成功
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUnHandleEventTabListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUnHandleEventTabListResponseParams `json:"Response"`
}

func (r *DescribeUnHandleEventTabListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnHandleEventTabListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcAcRuleRequestParams struct {
	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 需要查询的索引，特定场景使用，可不填
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeVpcAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 需要查询的索引，特定场景使用，可不填
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeVpcAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Index")
	delete(f, "Filters")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcAcRuleResponseParams struct {
	// 总条数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 内网间访问控制列表数据
	Data []*VpcRuleItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpcAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcAcRuleResponseParams `json:"Response"`
}

func (r *DescribeVpcAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcFwCcnPolicyWhiteListRequestParams struct {

}

type DescribeVpcFwCcnPolicyWhiteListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeVpcFwCcnPolicyWhiteListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcFwCcnPolicyWhiteListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcFwCcnPolicyWhiteListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcFwCcnPolicyWhiteListResponseParams struct {
	// 支持自动接入和策略路由的CCN列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportCcnPolicy []*string `json:"SupportCcnPolicy,omitnil,omitempty" name:"SupportCcnPolicy"`

	// 自动接入中支持自定义cidr的CCN列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportCcnPolicyCidr []*string `json:"SupportCcnPolicyCidr,omitnil,omitempty" name:"SupportCcnPolicyCidr"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpcFwCcnPolicyWhiteListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcFwCcnPolicyWhiteListResponseParams `json:"Response"`
}

func (r *DescribeVpcFwCcnPolicyWhiteListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcFwCcnPolicyWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcFwGroupSwitchRequestParams struct {
	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

type DescribeVpcFwGroupSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 每页条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil,omitempty" name:"By"`
}

func (r *DescribeVpcFwGroupSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcFwGroupSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Order")
	delete(f, "By")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVpcFwGroupSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVpcFwGroupSwitchResponseParams struct {
	// 开关列表
	SwitchList []*FwGroupSwitchShow `json:"SwitchList,omitnil,omitempty" name:"SwitchList"`

	// 开关总个数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVpcFwGroupSwitchResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVpcFwGroupSwitchResponseParams `json:"Response"`
}

func (r *DescribeVpcFwGroupSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcFwGroupSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DnsVpcSwitch struct {
	// vpc id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 0：设置为关闭 1:设置为打开
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type EdgeIpInfo struct {
	// 公网IP
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 公网 IP 类型 1 公网,2 弹性,3 弹性ipv6,4 anycastIP, 6 HighQualityEIP
	PublicIpType *int64 `json:"PublicIpType,omitnil,omitempty" name:"PublicIpType"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 内网IP
	IntranetIp *string `json:"IntranetIp,omitnil,omitempty" name:"IntranetIp"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 风险端口数
	PortRiskCount *int64 `json:"PortRiskCount,omitnil,omitempty" name:"PortRiskCount"`

	// 最近扫描时间
	LastScanTime *string `json:"LastScanTime,omitnil,omitempty" name:"LastScanTime"`

	// 是否为region eip
	// 0 不为region eip，不能选择串行
	// 1 为region eip 可以选择串行
	IsRegionEip *int64 `json:"IsRegionEip,omitnil,omitempty" name:"IsRegionEip"`

	// EIP 所关联的VPC
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 0: 该地域暂未支持串行
	// 1: 该用户未在该地域配置串行带宽
	// 2: 该用户已在该地域配置串行带宽，可以开启串行开关
	// 3. 该地域可以支持串行，但是未部署公共集群
	IsSerialRegion *int64 `json:"IsSerialRegion,omitnil,omitempty" name:"IsSerialRegion"`

	// 0: 不是公网CLB 可以开启串行开关
	// 1: 是公网CLB 不可以开启串行开关
	IsPublicClb *int64 `json:"IsPublicClb,omitnil,omitempty" name:"IsPublicClb"`

	// 0: 开启开关时提示要创建私有连接。
	// 1: 关闭该开关时提示删除私有连接。
	// 如果大于 1: 关闭开关 、开启开关不需提示创建删除私有连接。
	EndpointBindEipNum *int64 `json:"EndpointBindEipNum,omitnil,omitempty" name:"EndpointBindEipNum"`

	// 扫描深度
	ScanMode *string `json:"ScanMode,omitnil,omitempty" name:"ScanMode"`

	// 扫描状态
	ScanStatus *int64 `json:"ScanStatus,omitnil,omitempty" name:"ScanStatus"`

	// 开关状态
	// 0 : 关闭
	// 1 : 开启
	// 2 : 开启中
	// 3 : 关闭中
	// 4 : 异常
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 私有连接ID
	EndpointId *string `json:"EndpointId,omitnil,omitempty" name:"EndpointId"`

	// 私有连接IP
	EndpointIp *string `json:"EndpointIp,omitnil,omitempty" name:"EndpointIp"`

	// 0 : 旁路
	// 1 : 串行
	// 2 : 正在模式切换
	SwitchMode *uint64 `json:"SwitchMode,omitnil,omitempty" name:"SwitchMode"`

	// 开关权重
	SwitchWeight *int64 `json:"SwitchWeight,omitnil,omitempty" name:"SwitchWeight"`

	// 域名化CLB的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// IP超量状态
	OverUsedStatus *int64 `json:"OverUsedStatus,omitnil,omitempty" name:"OverUsedStatus"`

	// 0 都不支持
	// 1 支持旁路
	// 2 支持串行
	// 3 旁路串行都支持
	SwitchSupportType *int64 `json:"SwitchSupportType,omitnil,omitempty" name:"SwitchSupportType"`
}

type EdgeIpSwitch struct {
	// 公网IP
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// vpc 中第一个EIP开关打开，需要指定子网创建私有连接
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 创建私有连接指定IP
	EndpointIp *string `json:"EndpointIp,omitnil,omitempty" name:"EndpointIp"`

	// 0 : 旁路 1 : 串行
	SwitchMode *int64 `json:"SwitchMode,omitnil,omitempty" name:"SwitchMode"`
}

type EndpointInfo struct {
	// 引流私有连接端点id
	EndpointId *string `json:"EndpointId,omitnil,omitempty" name:"EndpointId"`

	// 引流VpcId
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 所属地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 引流Vpc的Cidr
	VpcCidr *string `json:"VpcCidr,omitnil,omitempty" name:"VpcCidr"`
}

type EnterpriseSecurityGroupRuleBetaInfo struct {
	// 任务id
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 时间
	LastTime *string `json:"LastTime,omitnil,omitempty" name:"LastTime"`
}

type EnterpriseSecurityGroupRuleRuleInfo struct {
	// 排序
	OrderIndex *int64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// 主键id
	RuleUuid *uint64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 规则uuid
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// 源规则内容
	SourceId *string `json:"SourceId,omitnil,omitempty" name:"SourceId"`

	// 源规则类型 
	// 取值范围 0/1/2/3/4/5/6/7/8/9/100
	// 0表示ip(net),
	// 1表示VPC实例(instance)
	// 2表示子网实例(instance)
	// 3表示CVM实例(instance)
	// 4表示CLB实例(instance)
	// 5表示ENI实例(instance)
	// 6表示数据库实例(instance)
	// 7表示模板(template)
	// 8表示标签(tag)
	// 9表示地域(region)
	// 100表示资产分组(resourcegroup)
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 目的规则内容
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 目的规则类型 
	// 取值范围 0/1/2/3/4/5/6/7/8/9/100
	// 0表示ip(net),
	// 1表示VPC实例(instance)
	// 2表示子网实例(instance)
	// 3表示CVM实例(instance)
	// 4表示CLB实例(instance)
	// 5表示ENI实例(instance)
	// 6表示数据库实例(instance)
	// 7表示模板(template)
	// 8表示标签(tag)
	// 9表示地域(region)
	// 100表示资产分组(resourcegroup)
	TargetType *int64 `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 协议名称
	// 取值范围:TCP/ANY/ICMP/UDP
	// ANY:表示所有
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 规则策略
	// 取值范围:1/2
	// 1:阻断
	// 2:放行
	Strategy *int64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// 规则启用状态 
	// 取值范围： 0/1
	// 0:未开启
	// 1:开启
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 描述
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 标签
	AclTags *string `json:"AclTags,omitnil,omitempty" name:"AclTags"`

	// 规则最新一次是否有改动
	// 取值范围：0/1
	// 0:否
	// 1:是
	IsNew *int64 `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 是否延迟下发规则 
	// 取值范围：0/1
	// 0:立即下发 1:延迟下发
	IsDelay *int64 `json:"IsDelay,omitnil,omitempty" name:"IsDelay"`

	// 服务模板id
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`

	// 源资产名称
	SouInstanceName *string `json:"SouInstanceName,omitnil,omitempty" name:"SouInstanceName"`

	// 源资产公网ip
	SouPublicIp *string `json:"SouPublicIp,omitnil,omitempty" name:"SouPublicIp"`

	// 源资产内网ip
	SouPrivateIp *string `json:"SouPrivateIp,omitnil,omitempty" name:"SouPrivateIp"`

	// 源资产网段信息
	SouCidr *string `json:"SouCidr,omitnil,omitempty" name:"SouCidr"`

	// 源模板名称
	SouParameterName *string `json:"SouParameterName,omitnil,omitempty" name:"SouParameterName"`

	// 目的资产名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 目的资产公网ip
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 目的资产内网ip
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 目的资产网段信息
	Cidr *string `json:"Cidr,omitnil,omitempty" name:"Cidr"`

	// 目的模板名称
	ParameterName *string `json:"ParameterName,omitnil,omitempty" name:"ParameterName"`

	// 端口模板名称
	ProtocolPortName *string `json:"ProtocolPortName,omitnil,omitempty" name:"ProtocolPortName"`

	// 自动化任务信息
	BetaList []*EnterpriseSecurityGroupRuleBetaInfo `json:"BetaList,omitnil,omitempty" name:"BetaList"`

	// 规则id  等同RuleUuid
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 域名解析的IP统计
	DnsParseCount *SgDnsParseCount `json:"DnsParseCount,omitnil,omitempty" name:"DnsParseCount"`

	// 规则创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 规则最近更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type ExpandCfwVerticalRequestParams struct {
	// nat：nat防火墙，ew：东西向防火墙
	FwType *string `json:"FwType,omitnil,omitempty" name:"FwType"`

	// 带宽值
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`

	// 弹性开关 1打开 0 关闭
	ElasticSwitch *int64 `json:"ElasticSwitch,omitnil,omitempty" name:"ElasticSwitch"`

	// 弹性带宽上限，单位Mbps
	ElasticBandwidth *int64 `json:"ElasticBandwidth,omitnil,omitempty" name:"ElasticBandwidth"`

	// 按量计费标签
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ExpandCfwVerticalRequest struct {
	*tchttp.BaseRequest
	
	// nat：nat防火墙，ew：东西向防火墙
	FwType *string `json:"FwType,omitnil,omitempty" name:"FwType"`

	// 带宽值
	Width *uint64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`

	// 弹性开关 1打开 0 关闭
	ElasticSwitch *int64 `json:"ElasticSwitch,omitnil,omitempty" name:"ElasticSwitch"`

	// 弹性带宽上限，单位Mbps
	ElasticBandwidth *int64 `json:"ElasticBandwidth,omitnil,omitempty" name:"ElasticBandwidth"`

	// 按量计费标签
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *ExpandCfwVerticalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExpandCfwVerticalRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FwType")
	delete(f, "Width")
	delete(f, "CfwInstance")
	delete(f, "ElasticSwitch")
	delete(f, "ElasticBandwidth")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExpandCfwVerticalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExpandCfwVerticalResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExpandCfwVerticalResponse struct {
	*tchttp.BaseResponse
	Response *ExpandCfwVerticalResponseParams `json:"Response"`
}

func (r *ExpandCfwVerticalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExpandCfwVerticalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FwCidrInfo struct {
	// 防火墙使用的网段类型，值VpcSelf/Assis/Custom分别代表自有网段优先/扩展网段优先/自定义
	FwCidrType *string `json:"FwCidrType,omitnil,omitempty" name:"FwCidrType"`

	// 为每个vpc指定防火墙的网段
	FwCidrLst []*FwVpcCidr `json:"FwCidrLst,omitnil,omitempty" name:"FwCidrLst"`

	// 其他防火墙占用网段，一般是防火墙需要独占vpc时指定的网段
	ComFwCidr *string `json:"ComFwCidr,omitnil,omitempty" name:"ComFwCidr"`
}

type FwDeploy struct {
	// 防火墙部署地域
	DeployRegion *string `json:"DeployRegion,omitnil,omitempty" name:"DeployRegion"`

	// 带宽，单位：Mbps
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 异地灾备 1：使用异地灾备；0：不使用异地灾备；为空则默认不使用异地灾备
	CrossAZone *int64 `json:"CrossAZone,omitnil,omitempty" name:"CrossAZone"`

	// 主可用区，为空则选择默认可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 备可用区，为空则选择默认可用区
	ZoneBak *string `json:"ZoneBak,omitnil,omitempty" name:"ZoneBak"`

	// 若为cdc防火墙时填充该id
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`
}

type FwGateway struct {
	// 防火墙网关id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关所属vpc id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 网关ip地址
	IpAddress *string `json:"IpAddress,omitnil,omitempty" name:"IpAddress"`
}

type FwGroupSwitch struct {
	// 防火墙实例的开关模式 1: 单点互通 2: 多点互通 3: 全互通 4: 自定义路由
	SwitchMode *int64 `json:"SwitchMode,omitnil,omitempty" name:"SwitchMode"`

	// 防火墙开关ID
	// 支持三种类型
	// 1. 边开关(单点互通)
	// 2. 点开关(多点互通)
	// 3. 全开关(全互通)
	SwitchId *string `json:"SwitchId,omitnil,omitempty" name:"SwitchId"`
}

type FwGroupSwitchShow struct {
	// 防火墙开关ID
	SwitchId *string `json:"SwitchId,omitnil,omitempty" name:"SwitchId"`

	// 防火墙开关NAME
	SwitchName *string `json:"SwitchName,omitnil,omitempty" name:"SwitchName"`

	// 互通模式
	SwitchMode *int64 `json:"SwitchMode,omitnil,omitempty" name:"SwitchMode"`

	// 开关边连接类型 0：对等连接， 1：云连网
	ConnectType *int64 `json:"ConnectType,omitnil,omitempty" name:"ConnectType"`

	// 连接ID
	ConnectId *string `json:"ConnectId,omitnil,omitempty" name:"ConnectId"`

	// 连接名称
	ConnectName *string `json:"ConnectName,omitnil,omitempty" name:"ConnectName"`

	// 源实例信息
	SrcInstancesInfo []*NetInstancesInfo `json:"SrcInstancesInfo,omitnil,omitempty" name:"SrcInstancesInfo"`

	// 目的实例信息
	DstInstancesInfo []*NetInstancesInfo `json:"DstInstancesInfo,omitnil,omitempty" name:"DstInstancesInfo"`

	// 防火墙(组)数据
	FwGroupId *string `json:"FwGroupId,omitnil,omitempty" name:"FwGroupId"`

	// 防火墙(组)名称
	FwGroupName *string `json:"FwGroupName,omitnil,omitempty" name:"FwGroupName"`

	// 开关状态 0：关 ， 1：开
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 开关的状态 0：正常， 1：转换中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 0-非sase实例，忽略，1-未绑定状态，2-已绑定
	AttachWithEdge *int64 `json:"AttachWithEdge,omitnil,omitempty" name:"AttachWithEdge"`

	// 对等防火墙和开关状态 0：正常， 1：对等未创建防火墙，2：对等已创建防火墙，未打开开关
	CrossEdgeStatus *int64 `json:"CrossEdgeStatus,omitnil,omitempty" name:"CrossEdgeStatus"`

	// 网络经过VPC防火墙CVM所在地域
	FwInsRegion []*string `json:"FwInsRegion,omitnil,omitempty" name:"FwInsRegion"`

	// 0 观察 1 拦截 2 严格 3 关闭 4 不支持ips 前端展示tag
	IpsAction *int64 `json:"IpsAction,omitnil,omitempty" name:"IpsAction"`

	// 开关关联的防火墙实例列表
	FwInsLst []*VpcFwInstanceShow `json:"FwInsLst,omitnil,omitempty" name:"FwInsLst"`

	// 开关是否处于bypass状态
	// 0：正常状态
	// 1：bypass状态
	BypassStatus *int64 `json:"BypassStatus,omitnil,omitempty" name:"BypassStatus"`

	// 0: ipv4 , 1:ipv6
	IpVersion *int64 `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`
}

type FwVpcCidr struct {
	// vpc的id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 防火墙网段，最少/24的网段
	FwCidr *string `json:"FwCidr,omitnil,omitempty" name:"FwCidr"`
}

type IPDefendStatus struct {
	// ip地址
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// 防护状态   1:防护打开; -1:地址错误; 其他:未防护
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type InstanceInfo struct {
	// appid信息
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 资产来源
	// 1公网 2内网
	InsSource *string `json:"InsSource,omitnil,omitempty" name:"InsSource"`

	// 资产类型
	//  3是cvm实例,4是clb实例,5是eni实例,6是mysql,7是redis,8是NAT,9是VPN,10是ES,11是MARIADB,12是KAFKA 13 NATFW
	InsType *int64 `json:"InsType,omitnil,omitempty" name:"InsType"`

	// 资产id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 资产名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 漏洞数
	LeakNum *string `json:"LeakNum,omitnil,omitempty" name:"LeakNum"`

	// 端口数
	PortNum *string `json:"PortNum,omitnil,omitempty" name:"PortNum"`

	// 内网ip
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 公网ip
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 地域
	RegionKey *string `json:"RegionKey,omitnil,omitempty" name:"RegionKey"`

	// 资产路径
	ResourcePath []*string `json:"ResourcePath,omitnil,omitempty" name:"ResourcePath"`

	// 扫描结果
	Server []*string `json:"Server,omitnil,omitempty" name:"Server"`

	// 子网id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// vpc名称
	VPCName *string `json:"VPCName,omitnil,omitempty" name:"VPCName"`

	// vpcid信息
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type IntArray struct {
	// 数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*int64 `json:"List,omitnil,omitempty" name:"List"`
}

type InterconnectPair struct {
	// 集合A
	GroupA []*AccessInstanceInfo `json:"GroupA,omitnil,omitempty" name:"GroupA"`

	// 集合B
	GroupB []*AccessInstanceInfo `json:"GroupB,omitnil,omitempty" name:"GroupB"`

	// 互联模式："CrossConnect": 交叉互联（组A内每个实例和组B内每个实例互联），"FullMesh": 全互联（组A实际和组B内容一致，相当于组内两两互联）
	InterconnectMode *string `json:"InterconnectMode,omitnil,omitempty" name:"InterconnectMode"`
}

type IntrusionDefenseRule struct {
	// 规则方向，0出站，1入站，3内网间
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 规则结束时间，格式：2006-01-02 15:04:05，必须大于当前时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 规则IP地址，IP与Domain必填其中之一
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// 规则域名，IP与Domain必填其中之一
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 规则开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 备注信息，长度不能超过50
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type IocListData struct {
	// 待处置IP地址，IP/Domain字段二选一
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// 只能为0或者1   0代表出站 1代表入站
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 待处置域名，IP/Domain字段二选一
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type IpStatic struct {
	// 值
	Num *int64 `json:"Num,omitnil,omitempty" name:"Num"`

	// 折线图横坐标时间
	StatTime *string `json:"StatTime,omitnil,omitempty" name:"StatTime"`
}

type LogInfo struct {
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

	// 日志来源主机名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 原始日志(仅在日志创建索引异常时有值)
	// 注意：此字段可能返回 null，表示取不到有效值。
	RawLog *string `json:"RawLog,omitnil,omitempty" name:"RawLog"`

	// 日志创建索引异常原因(仅在日志创建索引异常时有值)
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexStatus *string `json:"IndexStatus,omitnil,omitempty" name:"IndexStatus"`
}

type LogItem struct {
	// 日志Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 日志Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type LogItems struct {
	// 分析结果返回的KV数据对
	Data []*LogItem `json:"Data,omitnil,omitempty" name:"Data"`
}

type ModeInfo struct {
	// 0-观察模式, 1-拦截模式, 2-严格模式
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`
}

// Predefined struct for user
type ModifyAcRuleRequestParams struct {
	// 规则数组
	Data []*RuleInfoData `json:"Data,omitnil,omitempty" name:"Data"`

	// EdgeId值
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 访问规则状态
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// NAT地域
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type ModifyAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则数组
	Data []*RuleInfoData `json:"Data,omitnil,omitempty" name:"Data"`

	// EdgeId值
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 访问规则状态
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// NAT地域
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

func (r *ModifyAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	delete(f, "EdgeId")
	delete(f, "Enable")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAcRuleResponseParams struct {
	// 状态值，0:操作成功，非0：操作失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 返回多余的信息
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAcRuleResponseParams `json:"Response"`
}

func (r *ModifyAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAclRuleRequestParams struct {
	// 需要编辑的规则数组，基于Uuid唯一id修改该规则
	Rules []*CreateRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type ModifyAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要编辑的规则数组，基于Uuid唯一id修改该规则
	Rules []*CreateRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *ModifyAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAclRuleResponseParams struct {
	// 编辑成功后返回新策略ID列表
	RuleUuid []*int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAclRuleResponseParams `json:"Response"`
}

func (r *ModifyAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressTemplateRequestParams struct {
	// 地址模板唯一Id
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// 模板名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// Type为1，ip模板eg：1.1.1.1,2.2.2.2；
	// Type为5，域名模板eg：www.qq.com,www.tencent.com
	IpString *string `json:"IpString,omitnil,omitempty" name:"IpString"`

	// 1 ip模板  5 域名模板
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 协议端口模板，协议类型，4:4层协议，7:7层协议。Type=6时必填。
	ProtocolType *string `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`
}

type ModifyAddressTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 地址模板唯一Id
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// 模板名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模板描述
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// Type为1，ip模板eg：1.1.1.1,2.2.2.2；
	// Type为5，域名模板eg：www.qq.com,www.tencent.com
	IpString *string `json:"IpString,omitnil,omitempty" name:"IpString"`

	// 1 ip模板  5 域名模板
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 协议端口模板，协议类型，4:4层协议，7:7层协议。Type=6时必填。
	ProtocolType *string `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`
}

func (r *ModifyAddressTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Uuid")
	delete(f, "Name")
	delete(f, "Detail")
	delete(f, "IpString")
	delete(f, "Type")
	delete(f, "ProtocolType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAddressTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressTemplateResponseParams struct {
	// 创建结果,0成功
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一Id
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// 规则数上限配置
	RuleLimitNum *int64 `json:"RuleLimitNum,omitnil,omitempty" name:"RuleLimitNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAddressTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAddressTemplateResponseParams `json:"Response"`
}

func (r *ModifyAddressTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAddressTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAllPublicIPSwitchStatusRequestParams struct {
	// 状态，0：关闭，1：开启
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 选中的防火墙开关Id
	FireWallPublicIPs []*string `json:"FireWallPublicIPs,omitnil,omitempty" name:"FireWallPublicIPs"`
}

type ModifyAllPublicIPSwitchStatusRequest struct {
	*tchttp.BaseRequest
	
	// 状态，0：关闭，1：开启
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 选中的防火墙开关Id
	FireWallPublicIPs []*string `json:"FireWallPublicIPs,omitnil,omitempty" name:"FireWallPublicIPs"`
}

func (r *ModifyAllPublicIPSwitchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllPublicIPSwitchStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "FireWallPublicIPs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAllPublicIPSwitchStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAllPublicIPSwitchStatusResponseParams struct {
	// 接口返回信息
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 接口返回错误码，0请求成功  非0失败
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAllPublicIPSwitchStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAllPublicIPSwitchStatusResponseParams `json:"Response"`
}

func (r *ModifyAllPublicIPSwitchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllPublicIPSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAllRuleStatusRequestParams struct {
	// 状态，0：全部停用，1：全部启用
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Edge ID值
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// NAT地域
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

type ModifyAllRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 状态，0：全部停用，1：全部启用
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// Edge ID值
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// NAT地域
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`
}

func (r *ModifyAllRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "Direction")
	delete(f, "EdgeId")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAllRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAllRuleStatusResponseParams struct {
	// 0: 修改成功, 其他: 修改失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 规则限制数量
	RuleLimitNum *int64 `json:"RuleLimitNum,omitnil,omitempty" name:"RuleLimitNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAllRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAllRuleStatusResponseParams `json:"Response"`
}

func (r *ModifyAllRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAssetScanRequestParams struct {
	// 扫描范围：1端口, 2端口+漏扫
	ScanRange *int64 `json:"ScanRange,omitnil,omitempty" name:"ScanRange"`

	// 扫描深度：'heavy', 'medium', 'light'
	ScanDeep *string `json:"ScanDeep,omitnil,omitempty" name:"ScanDeep"`

	// 扫描类型：1立即扫描 2 周期任务
	RangeType *int64 `json:"RangeType,omitnil,omitempty" name:"RangeType"`

	// RangeType为2 是必须添加，定时任务时间
	ScanPeriod *string `json:"ScanPeriod,omitnil,omitempty" name:"ScanPeriod"`

	// 立即扫描这个字段传过滤的扫描集合
	ScanFilterIp []*string `json:"ScanFilterIp,omitnil,omitempty" name:"ScanFilterIp"`

	// 1全量2单个
	ScanType *int64 `json:"ScanType,omitnil,omitempty" name:"ScanType"`
}

type ModifyAssetScanRequest struct {
	*tchttp.BaseRequest
	
	// 扫描范围：1端口, 2端口+漏扫
	ScanRange *int64 `json:"ScanRange,omitnil,omitempty" name:"ScanRange"`

	// 扫描深度：'heavy', 'medium', 'light'
	ScanDeep *string `json:"ScanDeep,omitnil,omitempty" name:"ScanDeep"`

	// 扫描类型：1立即扫描 2 周期任务
	RangeType *int64 `json:"RangeType,omitnil,omitempty" name:"RangeType"`

	// RangeType为2 是必须添加，定时任务时间
	ScanPeriod *string `json:"ScanPeriod,omitnil,omitempty" name:"ScanPeriod"`

	// 立即扫描这个字段传过滤的扫描集合
	ScanFilterIp []*string `json:"ScanFilterIp,omitnil,omitempty" name:"ScanFilterIp"`

	// 1全量2单个
	ScanType *int64 `json:"ScanType,omitnil,omitempty" name:"ScanType"`
}

func (r *ModifyAssetScanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssetScanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScanRange")
	delete(f, "ScanDeep")
	delete(f, "RangeType")
	delete(f, "ScanPeriod")
	delete(f, "ScanFilterIp")
	delete(f, "ScanType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAssetScanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAssetScanResponseParams struct {
	// 接口返回信息
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 接口返回错误码，0请求成功  非0失败
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 状态值 0：成功，1 执行扫描中,其他：失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAssetScanResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAssetScanResponseParams `json:"Response"`
}

func (r *ModifyAssetScanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssetScanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAssetSyncRequestParams struct {

}

type ModifyAssetSyncRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyAssetSyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssetSyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAssetSyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAssetSyncResponseParams struct {
	// 返回状态
	// 0 请求成功
	// 2 请求失败
	// 3 请求失败-频率限制
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// success 成功
	// 其他失败
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 0 成功
	// 非0 失败
	ReturnCode *uint64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAssetSyncResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAssetSyncResponseParams `json:"Response"`
}

func (r *ModifyAssetSyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAssetSyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBlockIgnoreListRequestParams struct {
	// 1封禁列表 2 放通列表
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// IP、Domain二选一（注：封禁列表，只能填写IP），不能同时为空
	IOC []*IocListData `json:"IOC,omitnil,omitempty" name:"IOC"`

	// 可选值：delete（删除）、edit（编辑）、add（添加）  其他值无效
	IocAction *string `json:"IocAction,omitnil,omitempty" name:"IocAction"`

	// 时间格式：yyyy-MM-dd HH:mm:ss，IocAction 为edit或add时必填
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 时间格式：yyyy-MM-dd HH:mm:ss，IocAction 为edit或add时必填，必须大于当前时间且大于StartTime
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ModifyBlockIgnoreListRequest struct {
	*tchttp.BaseRequest
	
	// 1封禁列表 2 放通列表
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// IP、Domain二选一（注：封禁列表，只能填写IP），不能同时为空
	IOC []*IocListData `json:"IOC,omitnil,omitempty" name:"IOC"`

	// 可选值：delete（删除）、edit（编辑）、add（添加）  其他值无效
	IocAction *string `json:"IocAction,omitnil,omitempty" name:"IocAction"`

	// 时间格式：yyyy-MM-dd HH:mm:ss，IocAction 为edit或add时必填
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 时间格式：yyyy-MM-dd HH:mm:ss，IocAction 为edit或add时必填，必须大于当前时间且大于StartTime
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *ModifyBlockIgnoreListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockIgnoreListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleType")
	delete(f, "IOC")
	delete(f, "IocAction")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBlockIgnoreListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBlockIgnoreListResponseParams struct {
	// 接口返回信息
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 接口返回错误码，0请求成功  非0失败
	ReturnCode *uint64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBlockIgnoreListResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBlockIgnoreListResponseParams `json:"Response"`
}

func (r *ModifyBlockIgnoreListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockIgnoreListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBlockIgnoreRuleNewRequestParams struct {
	// 规则
	Rule *BanAndAllowRule `json:"Rule,omitnil,omitempty" name:"Rule"`

	// RuleType: 1放通列表 2外部IP 3域名 4情报 5资产 6自定义规则  7入侵防御规则
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`
}

type ModifyBlockIgnoreRuleNewRequest struct {
	*tchttp.BaseRequest
	
	// 规则
	Rule *BanAndAllowRule `json:"Rule,omitnil,omitempty" name:"Rule"`

	// RuleType: 1放通列表 2外部IP 3域名 4情报 5资产 6自定义规则  7入侵防御规则
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`
}

func (r *ModifyBlockIgnoreRuleNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockIgnoreRuleNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rule")
	delete(f, "RuleType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBlockIgnoreRuleNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBlockIgnoreRuleNewResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBlockIgnoreRuleNewResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBlockIgnoreRuleNewResponseParams `json:"Response"`
}

func (r *ModifyBlockIgnoreRuleNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockIgnoreRuleNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBlockIgnoreRuleRequestParams struct {
	// 规则列表
	Rule *IntrusionDefenseRule `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 规则类型，1封禁，2放通
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`
}

type ModifyBlockIgnoreRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则列表
	Rule *IntrusionDefenseRule `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 规则类型，1封禁，2放通
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`
}

func (r *ModifyBlockIgnoreRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockIgnoreRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rule")
	delete(f, "RuleType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBlockIgnoreRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBlockIgnoreRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBlockIgnoreRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBlockIgnoreRuleResponseParams `json:"Response"`
}

func (r *ModifyBlockIgnoreRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockIgnoreRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBlockTopRequestParams struct {
	// 操作类型 1 置顶 0取消
	OpeType *string `json:"OpeType,omitnil,omitempty" name:"OpeType"`

	// 记录id
	UniqueId *string `json:"UniqueId,omitnil,omitempty" name:"UniqueId"`
}

type ModifyBlockTopRequest struct {
	*tchttp.BaseRequest
	
	// 操作类型 1 置顶 0取消
	OpeType *string `json:"OpeType,omitnil,omitempty" name:"OpeType"`

	// 记录id
	UniqueId *string `json:"UniqueId,omitnil,omitempty" name:"UniqueId"`
}

func (r *ModifyBlockTopRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockTopRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpeType")
	delete(f, "UniqueId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBlockTopRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBlockTopResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBlockTopResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBlockTopResponseParams `json:"Response"`
}

func (r *ModifyBlockTopResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockTopResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterVpcFwSwitchRequestParams struct {
	// 开关，0：关闭，1：开启
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 集群模式vpc间防火墙ccn开关信息
	CcnSwitch []*CcnSwitchInfo `json:"CcnSwitch,omitnil,omitempty" name:"CcnSwitch"`
}

type ModifyClusterVpcFwSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 开关，0：关闭，1：开启
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 集群模式vpc间防火墙ccn开关信息
	CcnSwitch []*CcnSwitchInfo `json:"CcnSwitch,omitnil,omitempty" name:"CcnSwitch"`
}

func (r *ModifyClusterVpcFwSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterVpcFwSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Enable")
	delete(f, "CcnSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterVpcFwSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClusterVpcFwSwitchResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClusterVpcFwSwitchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterVpcFwSwitchResponseParams `json:"Response"`
}

func (r *ModifyClusterVpcFwSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterVpcFwSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEWRuleStatusRequestParams struct {
	// vpc规则必填，边id
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 是否开关开启，0：未开启，1：开启
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 规则方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 更改的规则当前执行顺序
	RuleSequence *uint64 `json:"RuleSequence,omitnil,omitempty" name:"RuleSequence"`

	// 规则类型，vpc：VPC间规则、nat：Nat边界规则
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`
}

type ModifyEWRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// vpc规则必填，边id
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 是否开关开启，0：未开启，1：开启
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 规则方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 更改的规则当前执行顺序
	RuleSequence *uint64 `json:"RuleSequence,omitnil,omitempty" name:"RuleSequence"`

	// 规则类型，vpc：VPC间规则、nat：Nat边界规则
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`
}

func (r *ModifyEWRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEWRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeId")
	delete(f, "Status")
	delete(f, "Direction")
	delete(f, "RuleSequence")
	delete(f, "RuleType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEWRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEWRuleStatusResponseParams struct {
	// 状态值，0：修改成功，非0：修改失败
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// 状态信息，success：查询成功，fail：查询失败
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEWRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEWRuleStatusResponseParams `json:"Response"`
}

func (r *ModifyEWRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEWRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeIpSwitchRequestParams struct {
	// 0 关闭开关
	// 1 打开开关
	// 2 不操作开关，此次切换模式
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 操作开关详情
	EdgeIpSwitchLst []*EdgeIpSwitch `json:"EdgeIpSwitchLst,omitnil,omitempty" name:"EdgeIpSwitchLst"`

	// 0 不自动选择子网
	// 1 自动选择子网创建私有连接
	AutoChooseSubnet *int64 `json:"AutoChooseSubnet,omitnil,omitempty" name:"AutoChooseSubnet"`

	// 0 切换为旁路
	// 1 切换为串行
	// 2 不切换模式，此次操作开关
	SwitchMode *int64 `json:"SwitchMode,omitnil,omitempty" name:"SwitchMode"`
}

type ModifyEdgeIpSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 0 关闭开关
	// 1 打开开关
	// 2 不操作开关，此次切换模式
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 操作开关详情
	EdgeIpSwitchLst []*EdgeIpSwitch `json:"EdgeIpSwitchLst,omitnil,omitempty" name:"EdgeIpSwitchLst"`

	// 0 不自动选择子网
	// 1 自动选择子网创建私有连接
	AutoChooseSubnet *int64 `json:"AutoChooseSubnet,omitnil,omitempty" name:"AutoChooseSubnet"`

	// 0 切换为旁路
	// 1 切换为串行
	// 2 不切换模式，此次操作开关
	SwitchMode *int64 `json:"SwitchMode,omitnil,omitempty" name:"SwitchMode"`
}

func (r *ModifyEdgeIpSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeIpSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Enable")
	delete(f, "EdgeIpSwitchLst")
	delete(f, "AutoChooseSubnet")
	delete(f, "SwitchMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEdgeIpSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEdgeIpSwitchResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEdgeIpSwitchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEdgeIpSwitchResponseParams `json:"Response"`
}

func (r *ModifyEdgeIpSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEdgeIpSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnterpriseSecurityDispatchStatusRequestParams struct {
	// 0：打开立即下发开关；
	// 
	// 1：关闭立即下发开关；
	// 
	// 2：关闭立即下发开关情况下，触发开始下发
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyEnterpriseSecurityDispatchStatusRequest struct {
	*tchttp.BaseRequest
	
	// 0：打开立即下发开关；
	// 
	// 1：关闭立即下发开关；
	// 
	// 2：关闭立即下发开关情况下，触发开始下发
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyEnterpriseSecurityDispatchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnterpriseSecurityDispatchStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEnterpriseSecurityDispatchStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnterpriseSecurityDispatchStatusResponseParams struct {
	// 0: 修改成功, 其他: 修改失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEnterpriseSecurityDispatchStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEnterpriseSecurityDispatchStatusResponseParams `json:"Response"`
}

func (r *ModifyEnterpriseSecurityDispatchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnterpriseSecurityDispatchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnterpriseSecurityGroupRuleRequestParams struct {
	// 规则的uuid，可通过查询规则列表获取
	RuleUuid *uint64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 修改类型，0：修改规则内容；1：修改单条规则开关状态；2：修改所有规则开关状态
	ModifyType *uint64 `json:"ModifyType,omitnil,omitempty" name:"ModifyType"`

	// 编辑后的企业安全组规则数据；修改规则状态不用填该字段
	Data *SecurityGroupRule `json:"Data,omitnil,omitempty" name:"Data"`

	// 0是关闭,1是开启
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type ModifyEnterpriseSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则的uuid，可通过查询规则列表获取
	RuleUuid *uint64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 修改类型，0：修改规则内容；1：修改单条规则开关状态；2：修改所有规则开关状态
	ModifyType *uint64 `json:"ModifyType,omitnil,omitempty" name:"ModifyType"`

	// 编辑后的企业安全组规则数据；修改规则状态不用填该字段
	Data *SecurityGroupRule `json:"Data,omitnil,omitempty" name:"Data"`

	// 0是关闭,1是开启
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

func (r *ModifyEnterpriseSecurityGroupRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnterpriseSecurityGroupRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleUuid")
	delete(f, "ModifyType")
	delete(f, "Data")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEnterpriseSecurityGroupRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnterpriseSecurityGroupRuleResponseParams struct {
	// 状态值，0：编辑成功，非0：编辑失败
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 编辑后新生成规则的Id
	NewRuleUuid *uint64 `json:"NewRuleUuid,omitnil,omitempty" name:"NewRuleUuid"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEnterpriseSecurityGroupRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEnterpriseSecurityGroupRuleResponseParams `json:"Response"`
}

func (r *ModifyEnterpriseSecurityGroupRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnterpriseSecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFwGroupSwitchRequestParams struct {
	// 打开或关闭开关
	// 0：关闭开关
	// 1：打开开关
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 是否操作全部开关 0 不操作全部开关，1 操作全部开关
	AllSwitch *int64 `json:"AllSwitch,omitnil,omitempty" name:"AllSwitch"`

	// 开关列表
	SwitchList []*FwGroupSwitch `json:"SwitchList,omitnil,omitempty" name:"SwitchList"`
}

type ModifyFwGroupSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 打开或关闭开关
	// 0：关闭开关
	// 1：打开开关
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 是否操作全部开关 0 不操作全部开关，1 操作全部开关
	AllSwitch *int64 `json:"AllSwitch,omitnil,omitempty" name:"AllSwitch"`

	// 开关列表
	SwitchList []*FwGroupSwitch `json:"SwitchList,omitnil,omitempty" name:"SwitchList"`
}

func (r *ModifyFwGroupSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFwGroupSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Enable")
	delete(f, "AllSwitch")
	delete(f, "SwitchList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFwGroupSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFwGroupSwitchResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyFwGroupSwitchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFwGroupSwitchResponseParams `json:"Response"`
}

func (r *ModifyFwGroupSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFwGroupSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIpsModeSwitchRequestParams struct {
	// 防护模式：0-观察模式, 1-拦截模式, 2-严格模式
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`
}

type ModifyIpsModeSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 防护模式：0-观察模式, 1-拦截模式, 2-严格模式
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`
}

func (r *ModifyIpsModeSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIpsModeSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIpsModeSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIpsModeSwitchResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyIpsModeSwitchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIpsModeSwitchResponseParams `json:"Response"`
}

func (r *ModifyIpsModeSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIpsModeSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatAcRuleRequestParams struct {
	// 需要编辑的规则数组,基于Uuid唯一id来修改该规则
	Rules []*CreateNatRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type ModifyNatAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要编辑的规则数组,基于Uuid唯一id来修改该规则
	Rules []*CreateNatRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *ModifyNatAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatAcRuleResponseParams struct {
	// 编辑成功后返回新策略ID列表
	RuleUuid []*int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNatAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNatAcRuleResponseParams `json:"Response"`
}

func (r *ModifyNatAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatFwReSelectRequestParams struct {
	// 模式 1：接入模式；0：新增模式
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`

	// 接入模式重新接入的nat网关列表，其中NatGwList和VpcList只能传递一个。
	NatGwList []*string `json:"NatGwList,omitnil,omitempty" name:"NatGwList"`

	// 新增模式重新接入的vpc列表，其中NatGwList和NatgwList只能传递一个。
	VpcList []*string `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil,omitempty" name:"FwCidrInfo"`
}

type ModifyNatFwReSelectRequest struct {
	*tchttp.BaseRequest
	
	// 模式 1：接入模式；0：新增模式
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`

	// 接入模式重新接入的nat网关列表，其中NatGwList和VpcList只能传递一个。
	NatGwList []*string `json:"NatGwList,omitnil,omitempty" name:"NatGwList"`

	// 新增模式重新接入的vpc列表，其中NatGwList和NatgwList只能传递一个。
	VpcList []*string `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil,omitempty" name:"FwCidrInfo"`
}

func (r *ModifyNatFwReSelectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatFwReSelectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	delete(f, "CfwInstance")
	delete(f, "NatGwList")
	delete(f, "VpcList")
	delete(f, "FwCidrInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatFwReSelectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatFwReSelectResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNatFwReSelectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNatFwReSelectResponseParams `json:"Response"`
}

func (r *ModifyNatFwReSelectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatFwReSelectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatFwSwitchRequestParams struct {
	// 开关，0：关闭，1：开启
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 防火墙实例id列表，其中CfwInsIdList，SubnetIdList和RouteTableIdList只能传递一种。
	CfwInsIdList []*string `json:"CfwInsIdList,omitnil,omitempty" name:"CfwInsIdList"`

	// 子网id列表，其中CfwInsIdList，SubnetIdList和RouteTableIdList只能传递一种。
	SubnetIdList []*string `json:"SubnetIdList,omitnil,omitempty" name:"SubnetIdList"`

	// 路由表id列表，其中CfwInsIdList，SubnetIdList和RouteTableIdList只能传递一种。
	RouteTableIdList []*string `json:"RouteTableIdList,omitnil,omitempty" name:"RouteTableIdList"`
}

type ModifyNatFwSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 开关，0：关闭，1：开启
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 防火墙实例id列表，其中CfwInsIdList，SubnetIdList和RouteTableIdList只能传递一种。
	CfwInsIdList []*string `json:"CfwInsIdList,omitnil,omitempty" name:"CfwInsIdList"`

	// 子网id列表，其中CfwInsIdList，SubnetIdList和RouteTableIdList只能传递一种。
	SubnetIdList []*string `json:"SubnetIdList,omitnil,omitempty" name:"SubnetIdList"`

	// 路由表id列表，其中CfwInsIdList，SubnetIdList和RouteTableIdList只能传递一种。
	RouteTableIdList []*string `json:"RouteTableIdList,omitnil,omitempty" name:"RouteTableIdList"`
}

func (r *ModifyNatFwSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatFwSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Enable")
	delete(f, "CfwInsIdList")
	delete(f, "SubnetIdList")
	delete(f, "RouteTableIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatFwSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatFwSwitchResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNatFwSwitchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNatFwSwitchResponseParams `json:"Response"`
}

func (r *ModifyNatFwSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatFwSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatFwVpcDnsSwitchRequestParams struct {
	// nat 防火墙 id
	NatFwInsId *string `json:"NatFwInsId,omitnil,omitempty" name:"NatFwInsId"`

	// DNS 开关切换列表
	DnsVpcSwitchLst []*DnsVpcSwitch `json:"DnsVpcSwitchLst,omitnil,omitempty" name:"DnsVpcSwitchLst"`
}

type ModifyNatFwVpcDnsSwitchRequest struct {
	*tchttp.BaseRequest
	
	// nat 防火墙 id
	NatFwInsId *string `json:"NatFwInsId,omitnil,omitempty" name:"NatFwInsId"`

	// DNS 开关切换列表
	DnsVpcSwitchLst []*DnsVpcSwitch `json:"DnsVpcSwitchLst,omitnil,omitempty" name:"DnsVpcSwitchLst"`
}

func (r *ModifyNatFwVpcDnsSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatFwVpcDnsSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NatFwInsId")
	delete(f, "DnsVpcSwitchLst")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatFwVpcDnsSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatFwVpcDnsSwitchResponseParams struct {
	// 修改成功
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNatFwVpcDnsSwitchResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNatFwVpcDnsSwitchResponseParams `json:"Response"`
}

func (r *ModifyNatFwVpcDnsSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatFwVpcDnsSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatInstanceRequestParams struct {
	// NAT防火墙实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// NAT防火墙实例ID
	NatInstanceId *string `json:"NatInstanceId,omitnil,omitempty" name:"NatInstanceId"`
}

type ModifyNatInstanceRequest struct {
	*tchttp.BaseRequest
	
	// NAT防火墙实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// NAT防火墙实例ID
	NatInstanceId *string `json:"NatInstanceId,omitnil,omitempty" name:"NatInstanceId"`
}

func (r *ModifyNatInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceName")
	delete(f, "NatInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatInstanceResponseParams struct {
	// 0 正常
	// -1 异常
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// nat实例唯一ID
	NatInstanceId *string `json:"NatInstanceId,omitnil,omitempty" name:"NatInstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNatInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNatInstanceResponseParams `json:"Response"`
}

func (r *ModifyNatInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatSequenceRulesRequestParams struct {
	// 规则快速排序：OrderIndex，原始序号；NewOrderIndex：新序号
	RuleChangeItems []*RuleChangeItem `json:"RuleChangeItems,omitnil,omitempty" name:"RuleChangeItems"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type ModifyNatSequenceRulesRequest struct {
	*tchttp.BaseRequest
	
	// 规则快速排序：OrderIndex，原始序号；NewOrderIndex：新序号
	RuleChangeItems []*RuleChangeItem `json:"RuleChangeItems,omitnil,omitempty" name:"RuleChangeItems"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

func (r *ModifyNatSequenceRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatSequenceRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleChangeItems")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNatSequenceRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNatSequenceRulesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNatSequenceRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNatSequenceRulesResponseParams `json:"Response"`
}

func (r *ModifyNatSequenceRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNatSequenceRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceGroupRequestParams struct {
	// 资产组id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 上级组资产组id
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`
}

type ModifyResourceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 资产组id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 上级组资产组id
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`
}

func (r *ModifyResourceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "ParentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyResourceGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourceGroupResponseParams `json:"Response"`
}

func (r *ModifyResourceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRunSyncAssetRequestParams struct {
	// 0: 互联网防火墙开关，1：vpc 防火墙开关
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type ModifyRunSyncAssetRequest struct {
	*tchttp.BaseRequest
	
	// 0: 互联网防火墙开关，1：vpc 防火墙开关
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *ModifyRunSyncAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRunSyncAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRunSyncAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRunSyncAssetResponseParams struct {
	// 0：同步成功，1：资产更新中，2：后台同步调用失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRunSyncAssetResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRunSyncAssetResponseParams `json:"Response"`
}

func (r *ModifyRunSyncAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRunSyncAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupItemRuleStatusRequestParams struct {
	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 是否开关开启，0：未开启，1：开启
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 更改的企业安全组规则的执行顺序
	RuleSequence *uint64 `json:"RuleSequence,omitnil,omitempty" name:"RuleSequence"`
}

type ModifySecurityGroupItemRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 是否开关开启，0：未开启，1：开启
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 更改的企业安全组规则的执行顺序
	RuleSequence *uint64 `json:"RuleSequence,omitnil,omitempty" name:"RuleSequence"`
}

func (r *ModifySecurityGroupItemRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupItemRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "Status")
	delete(f, "RuleSequence")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityGroupItemRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupItemRuleStatusResponseParams struct {
	// 状态值，0：修改成功，非0：修改失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySecurityGroupItemRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityGroupItemRuleStatusResponseParams `json:"Response"`
}

func (r *ModifySecurityGroupItemRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupItemRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupRuleRequestParams struct {
	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 编辑后是否启用规则，0：不启用，1：启用，默认1
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 编辑的企业安全组规则数据
	Data []*SecurityGroupListData `json:"Data,omitnil,omitempty" name:"Data"`

	// 编辑的企业安全组规则的原始执行顺序
	SgRuleOriginSequence *uint64 `json:"SgRuleOriginSequence,omitnil,omitempty" name:"SgRuleOriginSequence"`
}

type ModifySecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 编辑后是否启用规则，0：不启用，1：启用，默认1
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 编辑的企业安全组规则数据
	Data []*SecurityGroupListData `json:"Data,omitnil,omitempty" name:"Data"`

	// 编辑的企业安全组规则的原始执行顺序
	SgRuleOriginSequence *uint64 `json:"SgRuleOriginSequence,omitnil,omitempty" name:"SgRuleOriginSequence"`
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
	delete(f, "Direction")
	delete(f, "Enable")
	delete(f, "Data")
	delete(f, "SgRuleOriginSequence")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityGroupRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupRuleResponseParams struct {
	// 状态值，0：编辑成功，非0：编辑失败
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 编辑后新生成规则的Id
	NewRuleId *uint64 `json:"NewRuleId,omitnil,omitempty" name:"NewRuleId"`

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
type ModifySecurityGroupSequenceRulesRequestParams struct {
	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 企业安全组规则快速排序数据
	Data []*SecurityGroupOrderIndexData `json:"Data,omitnil,omitempty" name:"Data"`
}

type ModifySecurityGroupSequenceRulesRequest struct {
	*tchttp.BaseRequest
	
	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 企业安全组规则快速排序数据
	Data []*SecurityGroupOrderIndexData `json:"Data,omitnil,omitempty" name:"Data"`
}

func (r *ModifySecurityGroupSequenceRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupSequenceRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "Data")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityGroupSequenceRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityGroupSequenceRulesResponseParams struct {
	// 状态值，0：修改成功，非0：修改失败
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySecurityGroupSequenceRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityGroupSequenceRulesResponseParams `json:"Response"`
}

func (r *ModifySecurityGroupSequenceRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupSequenceRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySequenceAclRulesRequestParams struct {
	// 规则快速排序：OrderIndex，原始序号；NewOrderIndex：新序号
	RuleChangeItems []*RuleChangeItem `json:"RuleChangeItems,omitnil,omitempty" name:"RuleChangeItems"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type ModifySequenceAclRulesRequest struct {
	*tchttp.BaseRequest
	
	// 规则快速排序：OrderIndex，原始序号；NewOrderIndex：新序号
	RuleChangeItems []*RuleChangeItem `json:"RuleChangeItems,omitnil,omitempty" name:"RuleChangeItems"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

func (r *ModifySequenceAclRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySequenceAclRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleChangeItems")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySequenceAclRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySequenceAclRulesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySequenceAclRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifySequenceAclRulesResponseParams `json:"Response"`
}

func (r *ModifySequenceAclRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySequenceAclRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySequenceRulesRequestParams struct {
	// 边Id值
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 修改数据
	Data []*SequenceData `json:"Data,omitnil,omitempty" name:"Data"`

	// NAT地域
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 方向，0：出向，1：入向
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type ModifySequenceRulesRequest struct {
	*tchttp.BaseRequest
	
	// 边Id值
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 修改数据
	Data []*SequenceData `json:"Data,omitnil,omitempty" name:"Data"`

	// NAT地域
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 方向，0：出向，1：入向
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

func (r *ModifySequenceRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySequenceRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeId")
	delete(f, "Data")
	delete(f, "Area")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySequenceRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySequenceRulesResponseParams struct {
	// 0: 修改成功, 非0: 修改失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySequenceRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifySequenceRulesResponseParams `json:"Response"`
}

func (r *ModifySequenceRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySequenceRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStorageSettingRequestParams struct {

}

type ModifyStorageSettingRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyStorageSettingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStorageSettingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStorageSettingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStorageSettingResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyStorageSettingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStorageSettingResponseParams `json:"Response"`
}

func (r *ModifyStorageSettingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStorageSettingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableStatusRequestParams struct {
	// EdgeId值两个vpc间的边id
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 状态值，1：锁表，2：解锁表
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Nat所在地域
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 0： 出向，1：入向
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type ModifyTableStatusRequest struct {
	*tchttp.BaseRequest
	
	// EdgeId值两个vpc间的边id
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 状态值，1：锁表，2：解锁表
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// Nat所在地域
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 0： 出向，1：入向
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

func (r *ModifyTableStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeId")
	delete(f, "Status")
	delete(f, "Area")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTableStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTableStatusResponseParams struct {
	// 0：正常，-1：不正常
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTableStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTableStatusResponseParams `json:"Response"`
}

func (r *ModifyTableStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcAcRuleRequestParams struct {
	// 需要编辑的规则数组
	Rules []*VpcRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type ModifyVpcAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要编辑的规则数组
	Rules []*VpcRuleItem `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *ModifyVpcAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpcAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcAcRuleResponseParams struct {
	// 编辑成功后返回新策略ID列表
	RuleUuids []*int64 `json:"RuleUuids,omitnil,omitempty" name:"RuleUuids"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVpcAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpcAcRuleResponseParams `json:"Response"`
}

func (r *ModifyVpcAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcFwGroupRequestParams struct {
	// 编辑的防火墙(组)ID
	FwGroupId *string `json:"FwGroupId,omitnil,omitempty" name:"FwGroupId"`

	// 修改防火墙(组)名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 编辑的防火墙实例列表
	VpcFwInstances []*VpcFwInstance `json:"VpcFwInstances,omitnil,omitempty" name:"VpcFwInstances"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil,omitempty" name:"FwCidrInfo"`
}

type ModifyVpcFwGroupRequest struct {
	*tchttp.BaseRequest
	
	// 编辑的防火墙(组)ID
	FwGroupId *string `json:"FwGroupId,omitnil,omitempty" name:"FwGroupId"`

	// 修改防火墙(组)名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 编辑的防火墙实例列表
	VpcFwInstances []*VpcFwInstance `json:"VpcFwInstances,omitnil,omitempty" name:"VpcFwInstances"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil,omitempty" name:"FwCidrInfo"`
}

func (r *ModifyVpcFwGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcFwGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FwGroupId")
	delete(f, "Name")
	delete(f, "VpcFwInstances")
	delete(f, "FwCidrInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpcFwGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcFwGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVpcFwGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpcFwGroupResponseParams `json:"Response"`
}

func (r *ModifyVpcFwGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcFwGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcFwSequenceRulesRequestParams struct {
	// 规则快速排序：OrderIndex，原始序号；NewOrderIndex：新序号
	RuleChangeItems []*RuleChangeItem `json:"RuleChangeItems,omitnil,omitempty" name:"RuleChangeItems"`
}

type ModifyVpcFwSequenceRulesRequest struct {
	*tchttp.BaseRequest
	
	// 规则快速排序：OrderIndex，原始序号；NewOrderIndex：新序号
	RuleChangeItems []*RuleChangeItem `json:"RuleChangeItems,omitnil,omitempty" name:"RuleChangeItems"`
}

func (r *ModifyVpcFwSequenceRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcFwSequenceRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleChangeItems")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVpcFwSequenceRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVpcFwSequenceRulesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVpcFwSequenceRulesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVpcFwSequenceRulesResponseParams `json:"Response"`
}

func (r *ModifyVpcFwSequenceRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVpcFwSequenceRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MultiTopicSearchInformation struct {
	// 要检索分析的日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 透传上次接口返回的Context值，可获取后续更多日志，总计最多可获取1万条原始日志，过期时间1小时
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`
}

type NatFwEipsInfo struct {
	// 弹性公网ip
	Eip *string `json:"Eip,omitnil,omitempty" name:"Eip"`

	// 所属的Nat网关Id
	NatGatewayId *string `json:"NatGatewayId,omitnil,omitempty" name:"NatGatewayId"`

	// Nat网关名称
	NatGatewayName *string `json:"NatGatewayName,omitnil,omitempty" name:"NatGatewayName"`
}

type NatFwFilter struct {
	// 过滤的类型，例如实例id
	FilterType *string `json:"FilterType,omitnil,omitempty" name:"FilterType"`

	// 过滤的内容，以',' 分隔
	FilterContent *string `json:"FilterContent,omitnil,omitempty" name:"FilterContent"`
}

type NatFwInstance struct {
	// nat实例id
	NatinsId *string `json:"NatinsId,omitnil,omitempty" name:"NatinsId"`

	// nat实例名称
	NatinsName *string `json:"NatinsName,omitnil,omitempty" name:"NatinsName"`

	// 实例所在地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 0:新增模式，1:接入模式
	FwMode *int64 `json:"FwMode,omitnil,omitempty" name:"FwMode"`

	// 0:正常状态， 1: 正在创建
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// nat公网ip
	NatIp *string `json:"NatIp,omitnil,omitempty" name:"NatIp"`
}

type NatInstanceInfo struct {
	// nat实例id
	NatinsId *string `json:"NatinsId,omitnil,omitempty" name:"NatinsId"`

	// nat实例名称
	NatinsName *string `json:"NatinsName,omitnil,omitempty" name:"NatinsName"`

	// 实例所在地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 0: 新增模式，1:接入模式
	FwMode *int64 `json:"FwMode,omitnil,omitempty" name:"FwMode"`

	// 实例带宽大小 Mbps
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// 入向带宽峰值 bps
	InFlowMax *int64 `json:"InFlowMax,omitnil,omitempty" name:"InFlowMax"`

	// 出向带宽峰值 bps
	OutFlowMax *uint64 `json:"OutFlowMax,omitnil,omitempty" name:"OutFlowMax"`

	// 地域中文信息
	RegionZh *string `json:"RegionZh,omitnil,omitempty" name:"RegionZh"`

	// 公网ip数组
	EipAddress []*string `json:"EipAddress,omitnil,omitempty" name:"EipAddress"`

	// 内外使用ip数组
	VpcIp []*string `json:"VpcIp,omitnil,omitempty" name:"VpcIp"`

	// 实例关联子网数组
	Subnets []*string `json:"Subnets,omitnil,omitempty" name:"Subnets"`

	// 0 :正常 1：正在初始化
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 地域区域信息
	RegionDetail *string `json:"RegionDetail,omitnil,omitempty" name:"RegionDetail"`

	// 实例所在可用区
	ZoneZh *string `json:"ZoneZh,omitnil,omitempty" name:"ZoneZh"`

	// 实例所在可用区
	ZoneZhBak *string `json:"ZoneZhBak,omitnil,omitempty" name:"ZoneZhBak"`

	// 已使用规则数
	RuleUsed *uint64 `json:"RuleUsed,omitnil,omitempty" name:"RuleUsed"`

	// 实例的规则限制最大规格数
	RuleMax *uint64 `json:"RuleMax,omitnil,omitempty" name:"RuleMax"`

	// 实例引擎版本
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 引擎是否可升级：0，不可升级；1，可升级
	UpdateEnable *int64 `json:"UpdateEnable,omitnil,omitempty" name:"UpdateEnable"`

	// 是的需要升级引擎 支持 nat拨测 1需要 0不需要
	NeedProbeEngineUpdate *int64 `json:"NeedProbeEngineUpdate,omitnil,omitempty" name:"NeedProbeEngineUpdate"`

	// 引擎运行模式，Normal:正常, OnlyRoute:透明模式
	TrafficMode *string `json:"TrafficMode,omitnil,omitempty" name:"TrafficMode"`

	// 实例主所在可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例备所在可用区
	ZoneBak *string `json:"ZoneBak,omitnil,omitempty" name:"ZoneBak"`

	// 引擎预约升级时间
	ReserveTime *string `json:"ReserveTime,omitnil,omitempty" name:"ReserveTime"`

	// 引擎预约升级版本
	ReserveVersion *string `json:"ReserveVersion,omitnil,omitempty" name:"ReserveVersion"`

	// 引擎预约升级版本状态 stable:稳定版；previewed:预览版
	ReserveVersionState *string `json:"ReserveVersionState,omitnil,omitempty" name:"ReserveVersionState"`

	// 弹性开关
	// 1 打开
	// 0 关闭
	ElasticSwitch *int64 `json:"ElasticSwitch,omitnil,omitempty" name:"ElasticSwitch"`

	// 弹性带宽，单位Mbps
	ElasticBandwidth *int64 `json:"ElasticBandwidth,omitnil,omitempty" name:"ElasticBandwidth"`

	// 是否首次开通按量付费
	// 1 是
	// 0 不是
	IsFirstAfterPay *int64 `json:"IsFirstAfterPay,omitnil,omitempty" name:"IsFirstAfterPay"`
}

type NatSwitchListData struct {
	// 列表ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 子网名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`

	// IPv4 CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetCidr *string `json:"SubnetCidr,omitnil,omitempty" name:"SubnetCidr"`

	// 关联路由ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteId *string `json:"RouteId,omitnil,omitempty" name:"RouteId"`

	// 关联路由名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteName *string `json:"RouteName,omitnil,omitempty" name:"RouteName"`

	// 云服务器个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CvmNum *uint64 `json:"CvmNum,omitnil,omitempty" name:"CvmNum"`

	// 所属VPC ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 所属VPC名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 是否生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 开关状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// NAT网关ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatId *string `json:"NatId,omitnil,omitempty" name:"NatId"`

	// NAT网关名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatName *string `json:"NatName,omitnil,omitempty" name:"NatName"`

	// NAT防火墙实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatInsId *string `json:"NatInsId,omitnil,omitempty" name:"NatInsId"`

	// NAT防火墙实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatInsName *string `json:"NatInsName,omitnil,omitempty" name:"NatInsName"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 开关是否异常,0:正常,1:异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Abnormal *int64 `json:"Abnormal,omitnil,omitempty" name:"Abnormal"`

	// nat防火墙出口路由表id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ORTableId *string `json:"ORTableId,omitnil,omitempty" name:"ORTableId"`

	// nat防火墙出口路由表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ORTableName *string `json:"ORTableName,omitnil,omitempty" name:"ORTableName"`

	// 出口Snat Ip列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ohavips []*string `json:"Ohavips,omitnil,omitempty" name:"Ohavips"`
}

type NetInstancesInfo struct {
	// 网络实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 网络实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 网络cidr (多段以逗号分隔)
	InstanceCidr *string `json:"InstanceCidr,omitnil,omitempty" name:"InstanceCidr"`

	// 网络实例所在地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type NewModeItems struct {
	// 新增模式下接入的vpc列表
	VpcList []*string `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// 新增模式下绑定的出口弹性公网ip列表，其中Eips和AddCount至少传递一个。
	Eips []*string `json:"Eips,omitnil,omitempty" name:"Eips"`

	// 新增模式下新增绑定的出口弹性公网ip个数，其中Eips和AddCount至少传递一个。
	AddCount *int64 `json:"AddCount,omitnil,omitempty" name:"AddCount"`
}

type RegionCidrConfig struct {
	// 引流地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// CIDR模式：0-跳过，1-自动，2-自定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	CidrMode *int64 `json:"CidrMode,omitnil,omitempty" name:"CidrMode"`

	// 自定义CIDR（CidrMode=2时必填），其它时候为空字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomCidr *string `json:"CustomCidr,omitnil,omitempty" name:"CustomCidr"`
}

type RegionFwStatus struct {
	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 引流网络部署状态
	// 1. "NotDeployed"  防火墙集群未部署
	// 2. "Deployed"        防火墙集群已部署，但未创建引流网络
	// 3. "Auto"                防火墙集群已部署，并自动选择网段创建了引流网络
	// 4. "Custom"            防火墙集群已部署，并根据用户自定义网段创建了引流网络
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 引流网络的cidr，如果没有部署引流网络则为空
	Cidr *string `json:"Cidr,omitnil,omitempty" name:"Cidr"`
}

// Predefined struct for user
type RemoveAcRuleRequestParams struct {
	// 规则的uuid，可通过查询规则列表获取
	RuleUuid *int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`
}

type RemoveAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则的uuid，可通过查询规则列表获取
	RuleUuid *int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`
}

func (r *RemoveAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleUuid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveAcRuleResponseParams struct {
	// 删除成功后返回被删除策略的uuid
	RuleUuid *int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 0代表成功，-1代表失败
	ReturnCode *int64 `json:"ReturnCode,omitnil,omitempty" name:"ReturnCode"`

	// success代表成功，failed代表失败
	ReturnMsg *string `json:"ReturnMsg,omitnil,omitempty" name:"ReturnMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *RemoveAcRuleResponseParams `json:"Response"`
}

func (r *RemoveAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveAclRuleRequestParams struct {
	// 规则的uuid列表，可通过查询规则列表获取，注意：如果传入的是[-1]将删除所有规则
	RuleUuid []*int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type RemoveAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则的uuid列表，可通过查询规则列表获取，注意：如果传入的是[-1]将删除所有规则
	RuleUuid []*int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

func (r *RemoveAclRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveAclRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleUuid")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveAclRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveAclRuleResponseParams struct {
	// 删除成功后返回被删除策略的uuid列表
	RuleUuid []*int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveAclRuleResponse struct {
	*tchttp.BaseResponse
	Response *RemoveAclRuleResponseParams `json:"Response"`
}

func (r *RemoveAclRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveAclRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveEnterpriseSecurityGroupRuleRequestParams struct {
	// 规则的uuid，可通过查询规则列表获取
	RuleUuid *int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 删除类型，0是单条删除，RuleUuid填写删除规则id，1为全部删除，RuleUuid填0即可
	RemoveType *int64 `json:"RemoveType,omitnil,omitempty" name:"RemoveType"`
}

type RemoveEnterpriseSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则的uuid，可通过查询规则列表获取
	RuleUuid *int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 删除类型，0是单条删除，RuleUuid填写删除规则id，1为全部删除，RuleUuid填0即可
	RemoveType *int64 `json:"RemoveType,omitnil,omitempty" name:"RemoveType"`
}

func (r *RemoveEnterpriseSecurityGroupRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveEnterpriseSecurityGroupRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleUuid")
	delete(f, "RemoveType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveEnterpriseSecurityGroupRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveEnterpriseSecurityGroupRuleResponseParams struct {
	// 删除成功后返回被删除策略的uuid
	RuleUuid *int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 0代表成功，-1代表失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveEnterpriseSecurityGroupRuleResponse struct {
	*tchttp.BaseResponse
	Response *RemoveEnterpriseSecurityGroupRuleResponseParams `json:"Response"`
}

func (r *RemoveEnterpriseSecurityGroupRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveEnterpriseSecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveNatAcRuleRequestParams struct {
	// 规则的uuid列表，可通过查询规则列表获取，注意：如果传入的是[-1]将删除所有规则
	RuleUuid []*int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type RemoveNatAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则的uuid列表，可通过查询规则列表获取，注意：如果传入的是[-1]将删除所有规则
	RuleUuid []*int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

func (r *RemoveNatAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveNatAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleUuid")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveNatAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveNatAcRuleResponseParams struct {
	// 删除成功后返回被删除策略的uuid列表
	RuleUuid []*int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveNatAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *RemoveNatAcRuleResponseParams `json:"Response"`
}

func (r *RemoveNatAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveNatAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveVpcAcRuleRequestParams struct {
	// 规则的uuid列表，可通过查询规则列表获取，注意：如果传入的是[-1]将删除所有规则
	RuleUuids []*int64 `json:"RuleUuids,omitnil,omitempty" name:"RuleUuids"`

	// 仅当RuleUuids为-1时有效；0：删除Ipv4规则，1：删除Ipv6规则；默认为Ipv4类型规则
	IpVersion *uint64 `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`
}

type RemoveVpcAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则的uuid列表，可通过查询规则列表获取，注意：如果传入的是[-1]将删除所有规则
	RuleUuids []*int64 `json:"RuleUuids,omitnil,omitempty" name:"RuleUuids"`

	// 仅当RuleUuids为-1时有效；0：删除Ipv4规则，1：删除Ipv6规则；默认为Ipv4类型规则
	IpVersion *uint64 `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`
}

func (r *RemoveVpcAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveVpcAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleUuids")
	delete(f, "IpVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveVpcAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveVpcAcRuleResponseParams struct {
	// 删除成功后返回被删除策略的uuid列表
	RuleUuids []*int64 `json:"RuleUuids,omitnil,omitempty" name:"RuleUuids"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveVpcAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *RemoveVpcAcRuleResponseParams `json:"Response"`
}

func (r *RemoveVpcAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveVpcAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleChangeItem struct {
	// 原始sequence 值
	OrderIndex *int64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// 新的sequence 值
	NewOrderIndex *int64 `json:"NewOrderIndex,omitnil,omitempty" name:"NewOrderIndex"`

	// Ip版本，0：IPv4，1：IPv6，默认为IPv4
	IpVersion *int64 `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`
}

type RuleInfoData struct {
	// 执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// 访问源
	SourceIp *string `json:"SourceIp,omitnil,omitempty" name:"SourceIp"`

	// 访问目的
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 策略, 0：观察，1：阻断，2：放行
	Strategy *string `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// 访问源类型，1是IP，3是域名，4是IP地址模板，5是域名地址模板
	SourceType *uint64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 描述
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 访问目的类型，1是IP，3是域名，4是IP地址模板，5是域名地址模板
	TargetType *uint64 `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// id值
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 日志id，从告警处创建必传，其它为空
	LogId *string `json:"LogId,omitnil,omitempty" name:"LogId"`

	// 城市Code
	City *uint64 `json:"City,omitnil,omitempty" name:"City"`

	// 国家Code
	Country *uint64 `json:"Country,omitnil,omitempty" name:"Country"`

	// 云厂商，支持多个，以逗号分隔， 1:腾讯云（仅中国香港及海外）,2:阿里云,3:亚马逊云,4:华为云,5:微软云
	CloudCode *string `json:"CloudCode,omitnil,omitempty" name:"CloudCode"`

	// 是否为地域
	IsRegion *uint64 `json:"IsRegion,omitnil,omitempty" name:"IsRegion"`

	// 城市名
	CityName *string `json:"CityName,omitnil,omitempty" name:"CityName"`

	// 国家名
	CountryName *string `json:"CountryName,omitnil,omitempty" name:"CountryName"`

	// 国家二位iso代码或者省份缩写代码
	RegionIso *string `json:"RegionIso,omitnil,omitempty" name:"RegionIso"`
}

type ScanInfo struct {
	// 进度
	ScanPercent *float64 `json:"ScanPercent,omitnil,omitempty" name:"ScanPercent"`

	// 扫描结果信息
	ScanResultInfo *ScanResultInfo `json:"ScanResultInfo,omitnil,omitempty" name:"ScanResultInfo"`

	// 扫描状态 0扫描中 1完成  2未勾选自动扫描
	ScanStatus *int64 `json:"ScanStatus,omitnil,omitempty" name:"ScanStatus"`

	// 预计完成时间
	ScanTime *string `json:"ScanTime,omitnil,omitempty" name:"ScanTime"`
}

type ScanResultInfo struct {
	// 是否禁封端口
	BanStatus *bool `json:"BanStatus,omitnil,omitempty" name:"BanStatus"`

	// 防护ip数量
	IPNum *uint64 `json:"IPNum,omitnil,omitempty" name:"IPNum"`

	// 是否开启防护
	IPStatus *bool `json:"IPStatus,omitnil,omitempty" name:"IPStatus"`

	// 是否拦截攻击
	IdpStatus *bool `json:"IdpStatus,omitnil,omitempty" name:"IdpStatus"`

	// 暴露漏洞数量
	LeakNum *uint64 `json:"LeakNum,omitnil,omitempty" name:"LeakNum"`

	// 暴露端口数量
	PortNum *uint64 `json:"PortNum,omitnil,omitempty" name:"PortNum"`
}

type SearchLogErrors struct {
	// 日志主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCodeStr *string `json:"ErrorCodeStr,omitnil,omitempty" name:"ErrorCodeStr"`
}

type SearchLogInfos struct {
	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志存储生命周期
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 透传本次接口返回的Context值，可获取后续更多日志，过期时间1小时
	// 注意：此字段可能返回 null，表示取不到有效值。
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`
}

// Predefined struct for user
type SearchLogRequestParams struct {
	// 要检索分析的日志的起始时间，Unix时间戳（毫秒）
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 要检索分析的日志的结束时间，Unix时间戳（毫秒）
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 检索分析语句，最大长度为12KB
	// 语句由 <a href="https://cloud.tencent.com/document/product/614/47044" target="_blank">[检索条件]</a> | <a href="https://cloud.tencent.com/document/product/614/44061" target="_blank">[SQL语句]</a>构成，无需对日志进行统计分析时，可省略其中的管道符<code> | </code>及SQL语句
	// 使用*或空字符串可查询所有日志
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 检索语法规则，默认值为0，推荐使用1 。
	// 
	// - 0：Lucene语法
	// - 1：CQL语法（日志服务专用检索语法，控制台默认也使用该语法规则）。
	// 
	// 详细说明参见<a href="https://cloud.tencent.com/document/product/614/47044#RetrievesConditionalRules" target="_blank">检索条件语法规则</a>
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`

	// - 要检索分析的日志主题ID，仅能指定一个日志主题。
	// - 如需同时检索多个日志主题，请使用Topics参数。
	// - TopicId 和 Topics 不能同时使用，在一次请求中有且只能选择一个。
	// 各日志主题ID如下
	// 访问控制-互联网边界 cfw_rule_acl
	// 访问控制-NAT边界 cfw_rule_nat_acl
	// 访问控制-VPC边界 cfw_rule_vpc_acl
	// 访问控制-DNS开关 cfw_rule_dns_acl
	// 入侵防御 cfw_rule_threatinfo
	// 全流量检测与响应日志-流量分析 cfw_netflow_nta
	// 全流量检测与响应日志-流量告警 cfw_rule_ndr_threatinfo
	// 零信任运维-数据库登录 cfw_operate_db
	// 零信任运维-服务器访问 operate_remote_om
	// 零信任运维-Web服务访问 operate_web_access
	// 零信任运维-行为审计 remoteom_commands
	// 流量日志-互联网边界 cfw_netflow_border
	// 流量日志-NAT边界 cfw_netflow_nat
	// 流量日志-VPC边界 cfw_netflow_vpc
	// 流量日志-DNS开关 cfw_netflow_dns
	// 流量日志-内网流量 cfw_netflow_fl
	// 操作日志 operate_log_all
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// - 要检索分析的日志主题列表，最大支持50个日志主题。
	// - 检索单个日志主题时请使用TopicId。
	// - TopicId 和 Topics 不能同时使用，在一次请求中有且只能选择一个。
	Topics []*MultiTopicSearchInformation `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 原始日志是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc
	// 注意：
	// * 仅当检索分析语句(Query)不包含SQL时有效
	// * SQL结果排序方式参考<a href="https://cloud.tencent.com/document/product/614/58978" target="_blank">SQL ORDER BY语法</a>
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 表示单次查询返回的原始日志条数，默认为100，最大值为1000。
	// 注意：
	// * 仅当检索分析语句(Query)不包含SQL时有效
	// * SQL结果条数指定方式参考<a href="https://cloud.tencent.com/document/product/614/58977" target="_blank">SQL LIMIT语法</a>
	// 
	// 可通过两种方式获取后续更多日志：
	// * Context:透传上次接口返回的Context值，获取后续更多日志，总计最多可获取1万条原始日志
	// * Offset:偏移量，表示从第几行开始返回原始日志，无日志条数限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询原始日志的偏移量，表示从第几行开始返回原始日志，默认为0。 
	// 注意：
	// * 仅当检索分析语句(Query)不包含SQL时有效
	// * 不能与Context参数同时使用
	// * 仅适用于单日志主题检索
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 透传上次接口返回的Context值，可获取后续更多日志，总计最多可获取1万条原始日志，过期时间1小时。
	// 注意：
	// * 透传该参数时，请勿修改除该参数外的其它参数
	// * 仅适用于单日志主题检索，检索多个日志主题时，请使用Topics中的Context
	// * 仅当检索分析语句(Query)不包含SQL时有效，SQL获取后续结果参考<a href="https://cloud.tencent.com/document/product/614/58977" target="_blank">SQL LIMIT语法</a>
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 执行统计分析（Query中包含SQL）时，是否对原始日志先进行采样，再进行统计分析。
	// 0：自动采样;
	// 0～1：按指定采样率采样，例如0.02;
	// 1：不采样，即精确分析
	// 默认值为1
	SamplingRate *float64 `json:"SamplingRate,omitnil,omitempty" name:"SamplingRate"`

	// 为true代表使用新的检索结果返回方式，输出参数AnalysisRecords和Columns有效
	// 为false时代表使用老的检索结果返回方式, 输出AnalysisResults和ColNames有效
	// 两种返回方式在编码格式上有少量区别，建议使用true
	UseNewAnalysis *bool `json:"UseNewAnalysis,omitnil,omitempty" name:"UseNewAnalysis"`
}

type SearchLogRequest struct {
	*tchttp.BaseRequest
	
	// 要检索分析的日志的起始时间，Unix时间戳（毫秒）
	From *int64 `json:"From,omitnil,omitempty" name:"From"`

	// 要检索分析的日志的结束时间，Unix时间戳（毫秒）
	To *int64 `json:"To,omitnil,omitempty" name:"To"`

	// 检索分析语句，最大长度为12KB
	// 语句由 <a href="https://cloud.tencent.com/document/product/614/47044" target="_blank">[检索条件]</a> | <a href="https://cloud.tencent.com/document/product/614/44061" target="_blank">[SQL语句]</a>构成，无需对日志进行统计分析时，可省略其中的管道符<code> | </code>及SQL语句
	// 使用*或空字符串可查询所有日志
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 检索语法规则，默认值为0，推荐使用1 。
	// 
	// - 0：Lucene语法
	// - 1：CQL语法（日志服务专用检索语法，控制台默认也使用该语法规则）。
	// 
	// 详细说明参见<a href="https://cloud.tencent.com/document/product/614/47044#RetrievesConditionalRules" target="_blank">检索条件语法规则</a>
	SyntaxRule *uint64 `json:"SyntaxRule,omitnil,omitempty" name:"SyntaxRule"`

	// - 要检索分析的日志主题ID，仅能指定一个日志主题。
	// - 如需同时检索多个日志主题，请使用Topics参数。
	// - TopicId 和 Topics 不能同时使用，在一次请求中有且只能选择一个。
	// 各日志主题ID如下
	// 访问控制-互联网边界 cfw_rule_acl
	// 访问控制-NAT边界 cfw_rule_nat_acl
	// 访问控制-VPC边界 cfw_rule_vpc_acl
	// 访问控制-DNS开关 cfw_rule_dns_acl
	// 入侵防御 cfw_rule_threatinfo
	// 全流量检测与响应日志-流量分析 cfw_netflow_nta
	// 全流量检测与响应日志-流量告警 cfw_rule_ndr_threatinfo
	// 零信任运维-数据库登录 cfw_operate_db
	// 零信任运维-服务器访问 operate_remote_om
	// 零信任运维-Web服务访问 operate_web_access
	// 零信任运维-行为审计 remoteom_commands
	// 流量日志-互联网边界 cfw_netflow_border
	// 流量日志-NAT边界 cfw_netflow_nat
	// 流量日志-VPC边界 cfw_netflow_vpc
	// 流量日志-DNS开关 cfw_netflow_dns
	// 流量日志-内网流量 cfw_netflow_fl
	// 操作日志 operate_log_all
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// - 要检索分析的日志主题列表，最大支持50个日志主题。
	// - 检索单个日志主题时请使用TopicId。
	// - TopicId 和 Topics 不能同时使用，在一次请求中有且只能选择一个。
	Topics []*MultiTopicSearchInformation `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 原始日志是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc
	// 注意：
	// * 仅当检索分析语句(Query)不包含SQL时有效
	// * SQL结果排序方式参考<a href="https://cloud.tencent.com/document/product/614/58978" target="_blank">SQL ORDER BY语法</a>
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 表示单次查询返回的原始日志条数，默认为100，最大值为1000。
	// 注意：
	// * 仅当检索分析语句(Query)不包含SQL时有效
	// * SQL结果条数指定方式参考<a href="https://cloud.tencent.com/document/product/614/58977" target="_blank">SQL LIMIT语法</a>
	// 
	// 可通过两种方式获取后续更多日志：
	// * Context:透传上次接口返回的Context值，获取后续更多日志，总计最多可获取1万条原始日志
	// * Offset:偏移量，表示从第几行开始返回原始日志，无日志条数限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询原始日志的偏移量，表示从第几行开始返回原始日志，默认为0。 
	// 注意：
	// * 仅当检索分析语句(Query)不包含SQL时有效
	// * 不能与Context参数同时使用
	// * 仅适用于单日志主题检索
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 透传上次接口返回的Context值，可获取后续更多日志，总计最多可获取1万条原始日志，过期时间1小时。
	// 注意：
	// * 透传该参数时，请勿修改除该参数外的其它参数
	// * 仅适用于单日志主题检索，检索多个日志主题时，请使用Topics中的Context
	// * 仅当检索分析语句(Query)不包含SQL时有效，SQL获取后续结果参考<a href="https://cloud.tencent.com/document/product/614/58977" target="_blank">SQL LIMIT语法</a>
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 执行统计分析（Query中包含SQL）时，是否对原始日志先进行采样，再进行统计分析。
	// 0：自动采样;
	// 0～1：按指定采样率采样，例如0.02;
	// 1：不采样，即精确分析
	// 默认值为1
	SamplingRate *float64 `json:"SamplingRate,omitnil,omitempty" name:"SamplingRate"`

	// 为true代表使用新的检索结果返回方式，输出参数AnalysisRecords和Columns有效
	// 为false时代表使用老的检索结果返回方式, 输出AnalysisResults和ColNames有效
	// 两种返回方式在编码格式上有少量区别，建议使用true
	UseNewAnalysis *bool `json:"UseNewAnalysis,omitnil,omitempty" name:"UseNewAnalysis"`
}

func (r *SearchLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "SyntaxRule")
	delete(f, "TopicId")
	delete(f, "Topics")
	delete(f, "Sort")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Context")
	delete(f, "SamplingRate")
	delete(f, "UseNewAnalysis")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchLogResponseParams struct {
	// 透传本次接口返回的Context值，可获取后续更多日志，过期时间1小时。
	// 注意：
	// * 仅适用于单日志主题检索，检索多个日志主题时，请使用Topics中的Context
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 符合检索条件的日志是否已全部返回，如未全部返回可使用Context参数获取后续更多日志
	// 注意：仅当检索分析语句(Query)不包含SQL时有效
	ListOver *bool `json:"ListOver,omitnil,omitempty" name:"ListOver"`

	// 返回的是否为统计分析（即SQL）结果
	Analysis *bool `json:"Analysis,omitnil,omitempty" name:"Analysis"`

	// 匹配检索条件的原始日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*LogInfo `json:"Results,omitnil,omitempty" name:"Results"`

	// 日志统计分析结果的列名
	// 当UseNewAnalysis为false时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColNames []*string `json:"ColNames,omitnil,omitempty" name:"ColNames"`

	// 日志统计分析结果
	// 当UseNewAnalysis为false时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnalysisResults []*LogItems `json:"AnalysisResults,omitnil,omitempty" name:"AnalysisResults"`

	// 日志统计分析结果
	// 当UseNewAnalysis为true时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnalysisRecords []*string `json:"AnalysisRecords,omitnil,omitempty" name:"AnalysisRecords"`

	// 日志统计分析结果的列属性
	// 当UseNewAnalysis为true时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*Column `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 本次统计分析使用的采样率
	// 注意：此字段可能返回 null，表示取不到有效值。
	SamplingRate *float64 `json:"SamplingRate,omitnil,omitempty" name:"SamplingRate"`

	// 使用多日志主题检索时，各个日志主题的基本信息，例如报错信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Topics *SearchLogTopics `json:"Topics,omitnil,omitempty" name:"Topics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchLogResponse struct {
	*tchttp.BaseResponse
	Response *SearchLogResponseParams `json:"Response"`
}

func (r *SearchLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogTopics struct {
	// 多日志主题检索对应的错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Errors []*SearchLogErrors `json:"Errors,omitnil,omitempty" name:"Errors"`

	// 多日志主题检索各日志主题信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Infos []*SearchLogInfos `json:"Infos,omitnil,omitempty" name:"Infos"`
}

type SecurityGroupBothWayInfo struct {
	// 执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// 访问源
	SourceId *string `json:"SourceId,omitnil,omitempty" name:"SourceId"`

	// 访问源类型，默认为0，0: IP, 1: VPC, 2: SUBNET, 3: CVM, 4: CLB, 5: ENI, 6: CDB, 7: 参数模板, 100: 资产分组
	SourceType *uint64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 访问目的
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 访问目的类型，默认为0，0: IP, 1: VPC, 2: SUBNET, 3: CVM, 4: CLB, 5: ENI, 6: CDB, 7: 参数模板, 100: 资产分组
	TargetType *uint64 `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 目的端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 策略, 1：阻断，2：放行
	Strategy *uint64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 描述
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 是否开关开启，0：未开启，1：开启
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否是正常规则，0：正常，1：异常
	IsNew *uint64 `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 单/双向下发，0:单向下发，1：双向下发
	BothWay *uint64 `json:"BothWay,omitnil,omitempty" name:"BothWay"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 公网IP，多个以英文逗号分隔
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 内网IP，多个以英文逗号分隔
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 掩码地址，多个以英文逗号分隔
	Cidr *string `json:"Cidr,omitnil,omitempty" name:"Cidr"`

	// 端口协议类型参数模板id
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`

	// 是否使用端口协议模板，0：否，1：是
	ProtocolPortType *uint64 `json:"ProtocolPortType,omitnil,omitempty" name:"ProtocolPortType"`
}

type SecurityGroupListData struct {
	// 执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// 访问源
	SourceId *string `json:"SourceId,omitnil,omitempty" name:"SourceId"`

	// 访问源类型，默认为0，1: VPC, 2: SUBNET, 3: CVM, 4: CLB, 5: ENI, 6: CDB, 7: 参数模板, 100: 资源组
	SourceType *uint64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 访问目的
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 访问目的类型，默认为0，1: VPC, 2: SUBNET, 3: CVM, 4: CLB, 5: ENI, 6: CDB, 7: 参数模板, 100:资源组
	TargetType *uint64 `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 目的端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 策略, 1：阻断，2：放行
	Strategy *uint64 `json:"Strategy,omitnil,omitempty" name:"Strategy"`

	// 描述
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 单/双向下发，0:单向下发，1：双向下发
	BothWay *uint64 `json:"BothWay,omitnil,omitempty" name:"BothWay"`

	// 规则ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 是否开关开启，0：未开启，1：开启
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否是正常规则，0：正常，1：异常
	IsNew *uint64 `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 私有网络ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 公网IP，多个以英文逗号分隔
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 内网IP，多个以英文逗号分隔
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 掩码地址，多个以英文逗号分隔
	Cidr *string `json:"Cidr,omitnil,omitempty" name:"Cidr"`

	// 端口协议类型参数模板id
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`

	// 生成双向下发规则
	BothWayInfo []*SecurityGroupBothWayInfo `json:"BothWayInfo,omitnil,omitempty" name:"BothWayInfo"`

	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 是否使用端口协议模板，0：否，1：是
	ProtocolPortType *uint64 `json:"ProtocolPortType,omitnil,omitempty" name:"ProtocolPortType"`

	// Uuid
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 资产分组名称
	AssetGroupNameIn *string `json:"AssetGroupNameIn,omitnil,omitempty" name:"AssetGroupNameIn"`

	// 资产分组名称
	AssetGroupNameOut *string `json:"AssetGroupNameOut,omitnil,omitempty" name:"AssetGroupNameOut"`

	// 模板名称
	ParameterName *string `json:"ParameterName,omitnil,omitempty" name:"ParameterName"`

	// 端口协议类型参数模板名称
	ProtocolPortName *string `json:"ProtocolPortName,omitnil,omitempty" name:"ProtocolPortName"`
}

type SecurityGroupOrderIndexData struct {
	// 企业安全组规则当前执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// 企业安全组规则更新目标执行顺序
	NewOrderIndex *uint64 `json:"NewOrderIndex,omitnil,omitempty" name:"NewOrderIndex"`
}

type SecurityGroupRule struct {
	// 访问源示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板id(ipm-dyodhpby)
	// instance：资产实例id(ins-123456)
	// resourcegroup：资产分组id(cfwrg-xxxx)
	// tag：资源标签({\"Key\":\"标签key值\",\"Value\":\"标签Value值\"})
	// region：地域(ap-gaungzhou)
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// 访问源类型，类型可以为以下6种：net|template|instance|resourcegroup|tag|region
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 访问目的示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板id(ipm-dyodhpby)
	// instance：资产实例id(ins-123456)
	// resourcegroup：资产分组id(cfwrg-xxxx)
	// tag：资源标签({\"Key\":\"标签key值\",\"Value\":\"标签Value值\"})
	// region：地域(ap-gaungzhou)
	DestContent *string `json:"DestContent,omitnil,omitempty" name:"DestContent"`

	// 访问目的类型，类型可以为以下6种：net|template|instance|resourcegroup|tag|region
	DestType *string `json:"DestType,omitnil,omitempty" name:"DestType"`

	// 访问控制策略中设置的流量通过云防火墙的方式。取值：
	// accept：放行
	// drop：拒绝
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// 规则描述 用于规则使用或者场景的描述，最多支持50个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 规则顺序，-1表示最低，1表示最高，请勿和外层Type冲突（和外层的Type配合使用，当中间插入时，指定添加位置）
	OrderIndex *string `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// 协议；TCP/UDP/ICMP/ICMPv6/ANY
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 访问控制策略的端口。取值：
	// -1/-1：全部端口
	// 80：80端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 端口协议类型参数模板id；协议端口模板id；与Protocol,Port互斥
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil,omitempty" name:"ServiceTemplateId"`

	// （入参时无需填写，自动生成）规则对应的唯一id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// （入参时Enable无意义；由通用配置中新增规则启用状态控制）
	// 规则状态，true表示启用，false表示禁用
	Enable *string `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 规则对应的唯一内部id
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`
}

type SecurityGroupSimplifyRule struct {
	// 访问源示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板(ipm-dyodhpby)
	// instance：资产实例(ins-123456)
	// resourcegroup：资产分组(/全部分组/分组1/子分组1)
	// tag：资源标签({"Key":"标签key值","Value":"标签Value值"})
	// region：地域(ap-gaungzhou)
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// 访问目的示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板(ipm-dyodhpby)
	// instance：资产实例(ins-123456)
	// resourcegroup：资产分组(/全部分组/分组1/子分组1)
	// tag：资源标签({"Key":"标签key值","Value":"标签Value值"})
	// region：地域(ap-gaungzhou)
	DestContent *string `json:"DestContent,omitnil,omitempty" name:"DestContent"`

	// 协议；TCP/UDP/ICMP/ANY
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 规则对应的唯一id
	RuleUuid *int64 `json:"RuleUuid,omitnil,omitempty" name:"RuleUuid"`

	// 规则序号
	Sequence *int64 `json:"Sequence,omitnil,omitempty" name:"Sequence"`
}

type SequenceData struct {
	// 规则Id值
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 修改前执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// 修改后执行顺序
	NewOrderIndex *uint64 `json:"NewOrderIndex,omitnil,omitempty" name:"NewOrderIndex"`
}

// Predefined struct for user
type SetNatFwDnatRuleRequestParams struct {
	// 0：cfw新增模式，1：cfw接入模式。
	Mode *uint64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 操作类型，可选值：add，del，modify。
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 防火墙实例id，该字段必须传递。
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`

	// 添加或删除操作的Dnat规则列表。
	AddOrDelDnatRules []*CfwNatDnatRule `json:"AddOrDelDnatRules,omitnil,omitempty" name:"AddOrDelDnatRules"`

	// 修改操作的原始Dnat规则
	OriginDnat *CfwNatDnatRule `json:"OriginDnat,omitnil,omitempty" name:"OriginDnat"`

	// 修改操作的新的Dnat规则
	NewDnat *CfwNatDnatRule `json:"NewDnat,omitnil,omitempty" name:"NewDnat"`
}

type SetNatFwDnatRuleRequest struct {
	*tchttp.BaseRequest
	
	// 0：cfw新增模式，1：cfw接入模式。
	Mode *uint64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 操作类型，可选值：add，del，modify。
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 防火墙实例id，该字段必须传递。
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`

	// 添加或删除操作的Dnat规则列表。
	AddOrDelDnatRules []*CfwNatDnatRule `json:"AddOrDelDnatRules,omitnil,omitempty" name:"AddOrDelDnatRules"`

	// 修改操作的原始Dnat规则
	OriginDnat *CfwNatDnatRule `json:"OriginDnat,omitnil,omitempty" name:"OriginDnat"`

	// 修改操作的新的Dnat规则
	NewDnat *CfwNatDnatRule `json:"NewDnat,omitnil,omitempty" name:"NewDnat"`
}

func (r *SetNatFwDnatRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNatFwDnatRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	delete(f, "OperationType")
	delete(f, "CfwInstance")
	delete(f, "AddOrDelDnatRules")
	delete(f, "OriginDnat")
	delete(f, "NewDnat")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetNatFwDnatRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetNatFwDnatRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetNatFwDnatRuleResponse struct {
	*tchttp.BaseResponse
	Response *SetNatFwDnatRuleResponseParams `json:"Response"`
}

func (r *SetNatFwDnatRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNatFwDnatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetNatFwEipRequestParams struct {
	// bind：绑定eip；unbind：解绑eip；newAdd：新增防火墙弹性公网ip
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`

	// 当OperationType 为bind或unbind操作时，使用该字段。
	EipList []*string `json:"EipList,omitnil,omitempty" name:"EipList"`
}

type SetNatFwEipRequest struct {
	*tchttp.BaseRequest
	
	// bind：绑定eip；unbind：解绑eip；newAdd：新增防火墙弹性公网ip
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitnil,omitempty" name:"CfwInstance"`

	// 当OperationType 为bind或unbind操作时，使用该字段。
	EipList []*string `json:"EipList,omitnil,omitempty" name:"EipList"`
}

func (r *SetNatFwEipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNatFwEipRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OperationType")
	delete(f, "CfwInstance")
	delete(f, "EipList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetNatFwEipRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetNatFwEipResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetNatFwEipResponse struct {
	*tchttp.BaseResponse
	Response *SetNatFwEipResponseParams `json:"Response"`
}

func (r *SetNatFwEipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNatFwEipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SgDnsParseCount struct {
	// 有效下发的IP个数，离散数据
	ValidCount *int64 `json:"ValidCount,omitnil,omitempty" name:"ValidCount"`

	// 未下发的IP个数，离散数据
	InvalidCount *int64 `json:"InvalidCount,omitnil,omitempty" name:"InvalidCount"`
}

type StaticInfo struct {
	// 地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 资产id
	InsID *string `json:"InsID,omitnil,omitempty" name:"InsID"`

	// 资产名称
	InsName *string `json:"InsName,omitnil,omitempty" name:"InsName"`

	// ip信息
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 数
	Num *int64 `json:"Num,omitnil,omitempty" name:"Num"`

	// 端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`
}

// Predefined struct for user
type StopSecurityGroupRuleDispatchRequestParams struct {
	// 值为1，中止全部
	StopType *int64 `json:"StopType,omitnil,omitempty" name:"StopType"`
}

type StopSecurityGroupRuleDispatchRequest struct {
	*tchttp.BaseRequest
	
	// 值为1，中止全部
	StopType *int64 `json:"StopType,omitnil,omitempty" name:"StopType"`
}

func (r *StopSecurityGroupRuleDispatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopSecurityGroupRuleDispatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StopType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopSecurityGroupRuleDispatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopSecurityGroupRuleDispatchResponseParams struct {
	// true代表成功，false代表错误
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopSecurityGroupRuleDispatchResponse struct {
	*tchttp.BaseResponse
	Response *StopSecurityGroupRuleDispatchResponseParams `json:"Response"`
}

func (r *StopSecurityGroupRuleDispatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopSecurityGroupRuleDispatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StorageHistogram struct {
	// 访问控制日志存储量，单位B
	// 注意：此字段可能返回 null，表示取不到有效值。
	AclSize *int64 `json:"AclSize,omitnil,omitempty" name:"AclSize"`

	// 入侵防御日志存储量，单位B
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdsSize *int64 `json:"IdsSize,omitnil,omitempty" name:"IdsSize"`

	// 流量日志存储量，单位B
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetFlowSize *int64 `json:"NetFlowSize,omitnil,omitempty" name:"NetFlowSize"`

	// 操作日志存储量，单位B
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateSize *int64 `json:"OperateSize,omitnil,omitempty" name:"OperateSize"`

	// 统计时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// NDR流量日志存储量，单位B
	NDRNetflowSize *int64 `json:"NDRNetflowSize,omitnil,omitempty" name:"NDRNetflowSize"`

	// NDR风险日志存储量，单位B
	NDRRiskSize *int64 `json:"NDRRiskSize,omitnil,omitempty" name:"NDRRiskSize"`
}

type StorageHistogramShow struct {
	// 存储类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType []*string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dates []*string `json:"Dates,omitnil,omitempty" name:"Dates"`

	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*IntArray `json:"Data,omitnil,omitempty" name:"Data"`
}

type SwitchError struct {
	// 开关唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrIns *string `json:"ErrIns,omitnil,omitempty" name:"ErrIns"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// 错误类型区分
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrKey *string `json:"ErrKey,omitnil,omitempty" name:"ErrKey"`

	// 错误时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTime *string `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`
}

type SwitchListsData struct {
	// 公网IP
	PublicIp *string `json:"PublicIp,omitnil,omitempty" name:"PublicIp"`

	// 内网IP
	IntranetIp *string `json:"IntranetIp,omitnil,omitempty" name:"IntranetIp"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil,omitempty" name:"AssetType"`

	// 地域
	Area *string `json:"Area,omitnil,omitempty" name:"Area"`

	// 防火墙开关
	Switch *int64 `json:"Switch,omitnil,omitempty" name:"Switch"`

	// id值
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公网 IP 类型
	PublicIpType *uint64 `json:"PublicIpType,omitnil,omitempty" name:"PublicIpType"`

	// 风险端口数
	PortTimes *uint64 `json:"PortTimes,omitnil,omitempty" name:"PortTimes"`

	// 最近扫描时间
	LastTime *string `json:"LastTime,omitnil,omitempty" name:"LastTime"`

	// 扫描深度
	ScanMode *string `json:"ScanMode,omitnil,omitempty" name:"ScanMode"`

	// 扫描状态
	ScanStatus *uint64 `json:"ScanStatus,omitnil,omitempty" name:"ScanStatus"`
}

// Predefined struct for user
type SyncFwOperateRequestParams struct {
	// 同步操作类型：Route，同步防火墙路由
	SyncType *string `json:"SyncType,omitnil,omitempty" name:"SyncType"`

	// 防火墙类型；nat,nat防火墙;ew,vpc间防火墙
	FwType *string `json:"FwType,omitnil,omitempty" name:"FwType"`
}

type SyncFwOperateRequest struct {
	*tchttp.BaseRequest
	
	// 同步操作类型：Route，同步防火墙路由
	SyncType *string `json:"SyncType,omitnil,omitempty" name:"SyncType"`

	// 防火墙类型；nat,nat防火墙;ew,vpc间防火墙
	FwType *string `json:"FwType,omitnil,omitempty" name:"FwType"`
}

func (r *SyncFwOperateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncFwOperateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SyncType")
	delete(f, "FwType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncFwOperateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SyncFwOperateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SyncFwOperateResponse struct {
	*tchttp.BaseResponse
	Response *SyncFwOperateResponseParams `json:"Response"`
}

func (r *SyncFwOperateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncFwOperateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TLogInfo struct {
	// 封禁列表
	BanNum *int64 `json:"BanNum,omitnil,omitempty" name:"BanNum"`

	// 暴力破解
	BruteForceNum *int64 `json:"BruteForceNum,omitnil,omitempty" name:"BruteForceNum"`

	// 待处置告警
	HandleNum *int64 `json:"HandleNum,omitnil,omitempty" name:"HandleNum"`

	// 网络探测
	NetworkNum *int64 `json:"NetworkNum,omitnil,omitempty" name:"NetworkNum"`

	// 失陷主机
	OutNum *int64 `json:"OutNum,omitnil,omitempty" name:"OutNum"`

	// 漏洞攻击
	VulNum *int64 `json:"VulNum,omitnil,omitempty" name:"VulNum"`
}

type TagInfo struct {
	// 目标key
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 目标值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TemplateListInfo struct {
	// 模板ID
	Uuid *string `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// 模板名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// IP模板
	IpString *string `json:"IpString,omitnil,omitempty" name:"IpString"`

	// 插入时间
	InsertTime *string `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// 修改时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 模板类型
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 关联规则条数
	RulesNum *int64 `json:"RulesNum,omitnil,omitempty" name:"RulesNum"`

	// 模板Id
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 协议端口模板，协议类型，4:4层协议，7:7层协议
	ProtocolType *string `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`

	// 模板包含地址数量
	IPNum *int64 `json:"IPNum,omitnil,omitempty" name:"IPNum"`

	// IP版本,0,IPv4;1,IPv6
	IpVersion *int64 `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`
}

type UnHandleEvent struct {
	// 伪攻击链类型
	EventTableListStruct []*UnHandleEventDetail `json:"EventTableListStruct,omitnil,omitempty" name:"EventTableListStruct"`

	// 1 是  0否
	BaseLineUser *uint64 `json:"BaseLineUser,omitnil,omitempty" name:"BaseLineUser"`

	// 1 打开 0 关闭
	BaseLineInSwitch *uint64 `json:"BaseLineInSwitch,omitnil,omitempty" name:"BaseLineInSwitch"`

	// 1 打开 0 关闭
	BaseLineOutSwitch *uint64 `json:"BaseLineOutSwitch,omitnil,omitempty" name:"BaseLineOutSwitch"`

	// vpc间防火墙实例数量
	VpcFwCount *uint64 `json:"VpcFwCount,omitnil,omitempty" name:"VpcFwCount"`
}

type UnHandleEventDetail struct {
	// 安全事件名称
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 未处置事件数量
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`
}

// Predefined struct for user
type UpdateCheckCcnNonDirectFlagRequestParams struct {
	// 云联网ID
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`
}

type UpdateCheckCcnNonDirectFlagRequest struct {
	*tchttp.BaseRequest
	
	// 云联网ID
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`
}

func (r *UpdateCheckCcnNonDirectFlagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCheckCcnNonDirectFlagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCheckCcnNonDirectFlagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCheckCcnNonDirectFlagResponseParams struct {
	// 检测更新状态
	// "Checked"：重新检测完成
	// "Checking": 正在重新检测中，请稍后刷新状态查看
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateCheckCcnNonDirectFlagResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCheckCcnNonDirectFlagResponseParams `json:"Response"`
}

func (r *UpdateCheckCcnNonDirectFlagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCheckCcnNonDirectFlagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateClusterVpcFwRequestParams struct {
	// ccn防火墙开关配置信息
	CcnSwitch *CcnSwitchInfo `json:"CcnSwitch,omitnil,omitempty" name:"CcnSwitch"`
}

type UpdateClusterVpcFwRequest struct {
	*tchttp.BaseRequest
	
	// ccn防火墙开关配置信息
	CcnSwitch *CcnSwitchInfo `json:"CcnSwitch,omitnil,omitempty" name:"CcnSwitch"`
}

func (r *UpdateClusterVpcFwRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateClusterVpcFwRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CcnSwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateClusterVpcFwRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateClusterVpcFwResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateClusterVpcFwResponse struct {
	*tchttp.BaseResponse
	Response *UpdateClusterVpcFwResponseParams `json:"Response"`
}

func (r *UpdateClusterVpcFwResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateClusterVpcFwResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcDnsInfo struct {
	// vpc id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vpc 名称
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// nat 防火墙模式 0：新增模式， 1: 接入模式
	FwMode *int64 `json:"FwMode,omitnil,omitempty" name:"FwMode"`

	// vpc ipv4网段范围 CIDR（Classless Inter-Domain Routing，无类域间路由选择）
	VpcIpv4Cidr *string `json:"VpcIpv4Cidr,omitnil,omitempty" name:"VpcIpv4Cidr"`

	// 外网弹性ip，防火墙 dns解析地址
	DNSEip *string `json:"DNSEip,omitnil,omitempty" name:"DNSEip"`

	// nat网关id
	NatInsId *string `json:"NatInsId,omitnil,omitempty" name:"NatInsId"`

	// nat网关名称
	NatInsName *string `json:"NatInsName,omitnil,omitempty" name:"NatInsName"`

	// 0：开关关闭 ， 1: 开关打开
	SwitchStatus *int64 `json:"SwitchStatus,omitnil,omitempty" name:"SwitchStatus"`

	// 0：未防护， 1: 已防护，2：忽略此字段
	ProtectedStatus *uint64 `json:"ProtectedStatus,omitnil,omitempty" name:"ProtectedStatus"`

	// 是否支持DNS FW，0-不支持、1-支持
	SupportDNSFW *uint64 `json:"SupportDNSFW,omitnil,omitempty" name:"SupportDNSFW"`
}

type VpcFwCvmInsInfo struct {
	// VPC防火墙实例ID
	FwInsId *string `json:"FwInsId,omitnil,omitempty" name:"FwInsId"`

	// CVM所在地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// CVM所在地域中文
	RegionZh *string `json:"RegionZh,omitnil,omitempty" name:"RegionZh"`

	// CVM所在地域详情
	RegionDetail *string `json:"RegionDetail,omitnil,omitempty" name:"RegionDetail"`

	// 主机所在可用区
	ZoneZh *string `json:"ZoneZh,omitnil,omitempty" name:"ZoneZh"`

	// 备机所在可用区
	ZoneZhBack *string `json:"ZoneZhBack,omitnil,omitempty" name:"ZoneZhBack"`

	// 防火墙CVM带宽值
	BandWidth *int64 `json:"BandWidth,omitnil,omitempty" name:"BandWidth"`

	// 实例主机所在可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例备机所在可用区
	ZoneBak *string `json:"ZoneBak,omitnil,omitempty" name:"ZoneBak"`
}

type VpcFwGroupInfo struct {
	// 防火墙(组)ID
	FwGroupId *string `json:"FwGroupId,omitnil,omitempty" name:"FwGroupId"`

	// 防火墙(组)名称
	FwGroupName *string `json:"FwGroupName,omitnil,omitempty" name:"FwGroupName"`

	// 防火墙组涉及到的开关个数
	FwSwitchNum *int64 `json:"FwSwitchNum,omitnil,omitempty" name:"FwSwitchNum"`

	// 防火墙(组)部署的地域
	RegionLst []*string `json:"RegionLst,omitnil,omitempty" name:"RegionLst"`

	// 模式 1：CCN云联网模式；0：私有网络模式 2: sase 模式 3：ccn 高级模式 4: 私有网络(跨租户单边模式)
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 防火墙实例的开关模式 1: 单点互通 2: 多点互通 3: 全互通 4: 自定义路由
	SwitchMode *int64 `json:"SwitchMode,omitnil,omitempty" name:"SwitchMode"`

	// VPC防火墙实例卡片信息数组
	FwInstanceLst []*VpcFwInstanceInfo `json:"FwInstanceLst,omitnil,omitempty" name:"FwInstanceLst"`

	// 防火墙(状态) 0：正常 1: 初始化或操作中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// auto :自动选择
	// 如果为网段，则为用户自定义 192.168.0.0/20 
	FwVpcCidr *string `json:"FwVpcCidr,omitnil,omitempty" name:"FwVpcCidr"`

	// cdc专用集群场景时表示部署所属的cdc
	CdcId *string `json:"CdcId,omitnil,omitempty" name:"CdcId"`

	// cdc专用集群场景时表示cdc名称
	CdcName *string `json:"CdcName,omitnil,omitempty" name:"CdcName"`

	// 跨租户模式 1管理员 2单边 0 非跨租户
	CrossUserMode *string `json:"CrossUserMode,omitnil,omitempty" name:"CrossUserMode"`

	// 云联网模式下，当前实例是否需要开启重叠路由开关，1：需要开启，0：不需要开启
	NeedSwitchCcnOverlap *int64 `json:"NeedSwitchCcnOverlap,omitnil,omitempty" name:"NeedSwitchCcnOverlap"`

	// 云联网模式下，实例关联的云联网id
	CcnId *string `json:"CcnId,omitnil,omitempty" name:"CcnId"`
}

type VpcFwInstance struct {
	// 防火墙实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 私有网络模式下接入的VpcId列表；仅私有网络模式使用
	VpcIds []*string `json:"VpcIds,omitnil,omitempty" name:"VpcIds"`

	// 部署地域信息
	FwDeploy *FwDeploy `json:"FwDeploy,omitnil,omitempty" name:"FwDeploy"`

	// 防火墙实例ID (编辑场景传)
	FwInsId *string `json:"FwInsId,omitnil,omitempty" name:"FwInsId"`
}

type VpcFwInstanceInfo struct {
	// VPC防火墙实例名称
	FwInsName *string `json:"FwInsName,omitnil,omitempty" name:"FwInsName"`

	// VPC防火墙实例ID
	FwInsId *string `json:"FwInsId,omitnil,omitempty" name:"FwInsId"`

	// VPC防火墙实例模式 0: 旧VPC模式防火墙 1: CCN模式防火墙
	FwMode *int64 `json:"FwMode,omitnil,omitempty" name:"FwMode"`

	// VPC防火墙接入网络实例个数
	JoinInsNum *int64 `json:"JoinInsNum,omitnil,omitempty" name:"JoinInsNum"`

	// VPC防火墙开关个数
	FwSwitchNum *int64 `json:"FwSwitchNum,omitnil,omitempty" name:"FwSwitchNum"`

	// VPC防火墙状态 0:正常 ， 1：创建中 2: 变更中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// VPC防火墙创建时间
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// VPC 相关云联网ID列表
	CcnId []*string `json:"CcnId,omitnil,omitempty" name:"CcnId"`

	// VPC 相关云联网名称列表
	CcnName []*string `json:"CcnName,omitnil,omitempty" name:"CcnName"`

	// VPC 相关对等连接ID列表
	PeerConnectionId []*string `json:"PeerConnectionId,omitnil,omitempty" name:"PeerConnectionId"`

	// VPC 相关对等连接名称列表
	PeerConnectionName []*string `json:"PeerConnectionName,omitnil,omitempty" name:"PeerConnectionName"`

	// VPC防火墙CVM的列表
	FwCvmLst []*VpcFwCvmInsInfo `json:"FwCvmLst,omitnil,omitempty" name:"FwCvmLst"`

	// VPC防火墙接入网络实例类型列表
	JoinInsLst []*VpcFwJoinInstanceType `json:"JoinInsLst,omitnil,omitempty" name:"JoinInsLst"`

	// 防火墙网关信息
	FwGateway []*FwGateway `json:"FwGateway,omitnil,omitempty" name:"FwGateway"`

	// 防火墙(组)ID
	FwGroupId *string `json:"FwGroupId,omitnil,omitempty" name:"FwGroupId"`

	// 已使用规则数
	RuleUsed *int64 `json:"RuleUsed,omitnil,omitempty" name:"RuleUsed"`

	// 最大规则数
	RuleMax *int64 `json:"RuleMax,omitnil,omitempty" name:"RuleMax"`

	// 防火墙实例带宽
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 用户VPC墙总带宽
	UserVpcWidth *int64 `json:"UserVpcWidth,omitnil,omitempty" name:"UserVpcWidth"`

	// 接入的vpc列表
	JoinInsIdLst []*string `json:"JoinInsIdLst,omitnil,omitempty" name:"JoinInsIdLst"`

	// 内网间峰值带宽 (单位 bps )
	FlowMax *int64 `json:"FlowMax,omitnil,omitempty" name:"FlowMax"`

	// 实例引擎版本
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 引擎是否可升级：0，不可升级；1，可升级
	UpdateEnable *int64 `json:"UpdateEnable,omitnil,omitempty" name:"UpdateEnable"`

	// 引擎运行模式，Normal:正常, OnlyRoute:透明模式
	TrafficMode *string `json:"TrafficMode,omitnil,omitempty" name:"TrafficMode"`

	// 引擎预约升级时间
	ReserveTime *string `json:"ReserveTime,omitnil,omitempty" name:"ReserveTime"`

	// 预约引擎升级版本
	ReserveVersion *string `json:"ReserveVersion,omitnil,omitempty" name:"ReserveVersion"`

	// 引擎预约升级版本状态
	ReserveVersionState *string `json:"ReserveVersionState,omitnil,omitempty" name:"ReserveVersionState"`

	// 弹性开关 1打开 0关闭
	ElasticSwitch *int64 `json:"ElasticSwitch,omitnil,omitempty" name:"ElasticSwitch"`

	// 弹性带宽，单位Mbps
	ElasticBandwidth *int64 `json:"ElasticBandwidth,omitnil,omitempty" name:"ElasticBandwidth"`

	// 是否首次开通按量付费
	// 1 是
	// 0 不是
	IsFirstAfterPay *int64 `json:"IsFirstAfterPay,omitnil,omitempty" name:"IsFirstAfterPay"`
}

type VpcFwInstanceShow struct {
	// VPC防火墙实例ID
	FwInsId *string `json:"FwInsId,omitnil,omitempty" name:"FwInsId"`

	// VPC防火墙实例名称
	FwInsName *string `json:"FwInsName,omitnil,omitempty" name:"FwInsName"`

	// 网络经过VPC防火墙CVM所在地域
	FwInsRegion *string `json:"FwInsRegion,omitnil,omitempty" name:"FwInsRegion"`
}

type VpcFwJoinInstanceType struct {
	// 接入实例类型，VPC、DIRECTCONNECT、 VPNGW 等
	JoinType *string `json:"JoinType,omitnil,omitempty" name:"JoinType"`

	// 接入的对应网络实例类型的数量
	Num *int64 `json:"Num,omitnil,omitempty" name:"Num"`
}

type VpcRuleItem struct {
	// 访问源示例：
	// net：IP/CIDR(192.168.0.2)
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// 访问源类型，类型可以为：net
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 访问目的示例：
	// net：IP/CIDR(192.168.0.2)
	// domain：域名规则，例如*.qq.com
	DestContent *string `json:"DestContent,omitnil,omitempty" name:"DestContent"`

	// 访问目的类型，类型可以为：net，domain，dnsparse
	DestType *string `json:"DestType,omitnil,omitempty" name:"DestType"`

	// 协议，可选的值：
	// TCP
	// UDP
	// ICMP
	// ANY
	// HTTP
	// HTTPS
	// HTTP/HTTPS
	// SMTP
	// SMTPS
	// SMTP/SMTPS
	// FTP
	// DNS
	// TLS/SSL
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 访问控制策略中设置的流量通过云防火墙的方式。取值：
	// accept：放行
	// drop：拒绝
	// log：观察
	RuleAction *string `json:"RuleAction,omitnil,omitempty" name:"RuleAction"`

	// 访问控制策略的端口。取值：
	// -1/-1：全部端口
	// 80：80端口
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 规则顺序，-1表示最低，1表示最高
	OrderIndex *int64 `json:"OrderIndex,omitnil,omitempty" name:"OrderIndex"`

	// 规则状态，true表示启用，false表示禁用
	Enable *string `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 规则生效的范围，是在哪对vpc之间还是针对所有vpc间生效
	EdgeId *string `json:"EdgeId,omitnil,omitempty" name:"EdgeId"`

	// 规则对应的唯一id，添加规则时忽略该字段，修改该规则时需要填写Uuid;查询返回时会返回该参数
	Uuid *int64 `json:"Uuid,omitnil,omitempty" name:"Uuid"`

	// 规则的命中次数，增删改查规则时无需传入此参数，主要用于返回查询结果数据
	DetectedTimes *int64 `json:"DetectedTimes,omitnil,omitempty" name:"DetectedTimes"`

	// EdgeId对应的这对VPC间防火墙的描述
	EdgeName *string `json:"EdgeName,omitnil,omitempty" name:"EdgeName"`

	// 内部使用的uuid，一般情况下不会使用到该字段
	InternalUuid *int64 `json:"InternalUuid,omitnil,omitempty" name:"InternalUuid"`

	// 规则被删除：1，已删除；0，未删除
	Deleted *int64 `json:"Deleted,omitnil,omitempty" name:"Deleted"`

	// 规则生效的防火墙实例ID
	FwGroupId *string `json:"FwGroupId,omitnil,omitempty" name:"FwGroupId"`

	// 防火墙名称
	FwGroupName *string `json:"FwGroupName,omitnil,omitempty" name:"FwGroupName"`

	// beta任务详情
	BetaList []*BetaInfoByACL `json:"BetaList,omitnil,omitempty" name:"BetaList"`

	// 端口协议组ID
	ParamTemplateId *string `json:"ParamTemplateId,omitnil,omitempty" name:"ParamTemplateId"`

	// 端口协议组名称
	ParamTemplateName *string `json:"ParamTemplateName,omitnil,omitempty" name:"ParamTemplateName"`

	// 访问目的名称
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// 访问源名称
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// Ip版本，0：IPv4，1：IPv6，默认为IPv4
	IpVersion *int64 `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`

	// 是否是无效规则，0 表示有效规则，1 表示无效规则，出参场景返回使用
	Invalid *int64 `json:"Invalid,omitnil,omitempty" name:"Invalid"`

	// 规则创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 规则最近更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type VpcZoneData struct {
	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// vpc节点地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}