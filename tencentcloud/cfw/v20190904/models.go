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

package v20190904

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AcListsData struct {
	// 规则id
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 访问源
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceIp *string `json:"SourceIp,omitnil" name:"SourceIp"`

	// 访问目的
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetIp *string `json:"TargetIp,omitnil" name:"TargetIp"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitnil" name:"Port"`

	// 策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	Strategy *uint64 `json:"Strategy,omitnil" name:"Strategy"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *string `json:"Detail,omitnil" name:"Detail"`

	// 命中次数
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitnil" name:"OrderIndex"`

	// 告警规则id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogId *string `json:"LogId,omitnil" name:"LogId"`
}

// Predefined struct for user
type AddAcRuleRequestParams struct {
	// -1表示优先级最低，1表示优先级最高
	OrderIndex *string `json:"OrderIndex,omitnil" name:"OrderIndex"`

	// 访问控制策略中设置的流量通过云防火墙的方式。取值：
	// accept：放行
	// drop：拒绝
	// log：观察
	RuleAction *string `json:"RuleAction,omitnil" name:"RuleAction"`

	// 访问控制策略的流量方向。取值：
	// in：外对内流量访问控制
	// out：内对外流量访问控制
	Direction *string `json:"Direction,omitnil" name:"Direction"`

	// 访问控制策略的描述信息
	Description *string `json:"Description,omitnil" name:"Description"`

	// 访问控制策略中的源地址类型。取值：
	// net：源IP或网段（IP或者CIDR）
	// location：源区域
	// template：云防火墙地址模板
	// instance：实例id
	// vendor：云厂商
	SourceType *string `json:"SourceType,omitnil" name:"SourceType"`

	// 访问控制策略中的源地址。取值：
	// 当SourceType为net时，SourceContent为源IP地址或者CIDR地址。
	// 例如：1.1.1.0/24
	// 
	// 当SourceType为template时，SourceContent为源地址模板id。
	// 
	// 当SourceType为location时，SourceContent为源区域。
	// 例如["BJ11", "ZB"]
	// 
	// 当SourceType为instance时，SourceContent为该实例id对应的公网ip。
	// 例如ins-xxxxx
	// 
	// 当SourceType为vendor时，SourceContent为所选择厂商的公网ip列表。
	// 例如：aws,huawei,tencent,aliyun,azure,all代表以上五个
	SourceContent *string `json:"SourceContent,omitnil" name:"SourceContent"`

	// 访问控制策略中的目的地址类型。取值：
	// net：目的IP或者网段（IP或者CIDR）
	// location：源区域
	// template：云防火墙地址模板
	// instance：实例id
	// vendor：云厂商
	// domain: 域名或者ip
	DestType *string `json:"DestType,omitnil" name:"DestType"`

	// 访问控制策略中的目的地址。取值：
	// 当DestType为net时，DestContent为源IP地址或者CIDR地址。
	// 例如：1.1.1.0/24
	// 
	// 当DestType为template时，DestContent为源地址模板id。
	// 
	// 当DestType为location时，DestContent为源区域。
	// 例如["BJ11", "ZB"]
	// 
	// 当DestType为instance时，DestContent为该实例id对应的公网ip。
	// 例如ins-xxxxx
	// 
	// 当DestType为domain时，DestContent为该实例id对应的域名规则。
	// 例如*.qq.com
	// 
	// 当DestType为vendor时，DestContent为所选择厂商的公网ip列表。
	// 例如：aws,huawei,tencent,aliyun,azure,all代表以上五个
	DestContent *string `json:"DestContent,omitnil" name:"DestContent"`

	// 访问控制策略的端口。取值：
	// -1/-1：全部端口
	// 80,443：80或者443
	Port *string `json:"Port,omitnil" name:"Port"`

	// 访问控制策略中流量访问的协议类型。取值：TCP，目前互联网边界规则只能支持TCP，不传参数默认就是TCP
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 七层协议，取值：
	// HTTP/HTTPS
	// TLS/SSL
	ApplicationName *string `json:"ApplicationName,omitnil" name:"ApplicationName"`

	// 是否启用规则，默认为启用，取值：
	// true为启用，false为不启用
	Enable *string `json:"Enable,omitnil" name:"Enable"`
}

type AddAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// -1表示优先级最低，1表示优先级最高
	OrderIndex *string `json:"OrderIndex,omitnil" name:"OrderIndex"`

	// 访问控制策略中设置的流量通过云防火墙的方式。取值：
	// accept：放行
	// drop：拒绝
	// log：观察
	RuleAction *string `json:"RuleAction,omitnil" name:"RuleAction"`

	// 访问控制策略的流量方向。取值：
	// in：外对内流量访问控制
	// out：内对外流量访问控制
	Direction *string `json:"Direction,omitnil" name:"Direction"`

	// 访问控制策略的描述信息
	Description *string `json:"Description,omitnil" name:"Description"`

	// 访问控制策略中的源地址类型。取值：
	// net：源IP或网段（IP或者CIDR）
	// location：源区域
	// template：云防火墙地址模板
	// instance：实例id
	// vendor：云厂商
	SourceType *string `json:"SourceType,omitnil" name:"SourceType"`

	// 访问控制策略中的源地址。取值：
	// 当SourceType为net时，SourceContent为源IP地址或者CIDR地址。
	// 例如：1.1.1.0/24
	// 
	// 当SourceType为template时，SourceContent为源地址模板id。
	// 
	// 当SourceType为location时，SourceContent为源区域。
	// 例如["BJ11", "ZB"]
	// 
	// 当SourceType为instance时，SourceContent为该实例id对应的公网ip。
	// 例如ins-xxxxx
	// 
	// 当SourceType为vendor时，SourceContent为所选择厂商的公网ip列表。
	// 例如：aws,huawei,tencent,aliyun,azure,all代表以上五个
	SourceContent *string `json:"SourceContent,omitnil" name:"SourceContent"`

	// 访问控制策略中的目的地址类型。取值：
	// net：目的IP或者网段（IP或者CIDR）
	// location：源区域
	// template：云防火墙地址模板
	// instance：实例id
	// vendor：云厂商
	// domain: 域名或者ip
	DestType *string `json:"DestType,omitnil" name:"DestType"`

	// 访问控制策略中的目的地址。取值：
	// 当DestType为net时，DestContent为源IP地址或者CIDR地址。
	// 例如：1.1.1.0/24
	// 
	// 当DestType为template时，DestContent为源地址模板id。
	// 
	// 当DestType为location时，DestContent为源区域。
	// 例如["BJ11", "ZB"]
	// 
	// 当DestType为instance时，DestContent为该实例id对应的公网ip。
	// 例如ins-xxxxx
	// 
	// 当DestType为domain时，DestContent为该实例id对应的域名规则。
	// 例如*.qq.com
	// 
	// 当DestType为vendor时，DestContent为所选择厂商的公网ip列表。
	// 例如：aws,huawei,tencent,aliyun,azure,all代表以上五个
	DestContent *string `json:"DestContent,omitnil" name:"DestContent"`

	// 访问控制策略的端口。取值：
	// -1/-1：全部端口
	// 80,443：80或者443
	Port *string `json:"Port,omitnil" name:"Port"`

	// 访问控制策略中流量访问的协议类型。取值：TCP，目前互联网边界规则只能支持TCP，不传参数默认就是TCP
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 七层协议，取值：
	// HTTP/HTTPS
	// TLS/SSL
	ApplicationName *string `json:"ApplicationName,omitnil" name:"ApplicationName"`

	// 是否启用规则，默认为启用，取值：
	// true为启用，false为不启用
	Enable *string `json:"Enable,omitnil" name:"Enable"`
}

func (r *AddAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAcRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderIndex")
	delete(f, "RuleAction")
	delete(f, "Direction")
	delete(f, "Description")
	delete(f, "SourceType")
	delete(f, "SourceContent")
	delete(f, "DestType")
	delete(f, "DestContent")
	delete(f, "Port")
	delete(f, "Protocol")
	delete(f, "ApplicationName")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAcRuleResponseParams struct {
	// 创建成功后返回新策略的uuid
	RuleUuid *int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`

	// 0代表成功，-1代表失败
	ReturnCode *int64 `json:"ReturnCode,omitnil" name:"ReturnCode"`

	// success代表成功，failed代表失败
	ReturnMsg *string `json:"ReturnMsg,omitnil" name:"ReturnMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AddAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *AddAcRuleResponseParams `json:"Response"`
}

func (r *AddAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddAclRuleRequestParams struct {
	// 需要添加的访问控制规则列表
	Rules []*CreateRuleItem `json:"Rules,omitnil" name:"Rules"`

	// 添加规则的来源，一般不需要使用，值insert_rule 表示插入指定位置的规则；值batch_import 表示批量导入规则；为空时表示添加规则
	From *string `json:"From,omitnil" name:"From"`
}

type AddAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要添加的访问控制规则列表
	Rules []*CreateRuleItem `json:"Rules,omitnil" name:"Rules"`

	// 添加规则的来源，一般不需要使用，值insert_rule 表示插入指定位置的规则；值batch_import 表示批量导入规则；为空时表示添加规则
	From *string `json:"From,omitnil" name:"From"`
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
	RuleUuid []*int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Data []*SecurityGroupRule `json:"Data,omitnil" name:"Data"`

	// 添加类型，0：添加到最后，1：添加到最前；2：中间插入；默认0添加到最后
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符，且不能超过64个字符。
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 是否延迟下发，1则延迟下发，否则立即下发
	IsDelay *uint64 `json:"IsDelay,omitnil" name:"IsDelay"`
}

type AddEnterpriseSecurityGroupRulesRequest struct {
	*tchttp.BaseRequest
	
	// 创建规则数据
	Data []*SecurityGroupRule `json:"Data,omitnil" name:"Data"`

	// 添加类型，0：添加到最后，1：添加到最前；2：中间插入；默认0添加到最后
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符，且不能超过64个字符。
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`

	// 是否延迟下发，1则延迟下发，否则立即下发
	IsDelay *uint64 `json:"IsDelay,omitnil" name:"IsDelay"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddEnterpriseSecurityGroupRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddEnterpriseSecurityGroupRulesResponseParams struct {
	// 状态值，0：添加成功，非0：添加失败
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 规则uuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*SecurityGroupSimplifyRule `json:"Rules,omitnil" name:"Rules"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Rules []*CreateNatRuleItem `json:"Rules,omitnil" name:"Rules"`

	// 添加规则的来源，一般不需要使用，值insert_rule 表示插入指定位置的规则；值batch_import 表示批量导入规则；为空时表示添加规则
	From *string `json:"From,omitnil" name:"From"`
}

type AddNatAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要添加的nat访问控制规则列表
	Rules []*CreateNatRuleItem `json:"Rules,omitnil" name:"Rules"`

	// 添加规则的来源，一般不需要使用，值insert_rule 表示插入指定位置的规则；值batch_import 表示批量导入规则；为空时表示添加规则
	From *string `json:"From,omitnil" name:"From"`
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
	RuleUuid []*int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Rules []*VpcRuleItem `json:"Rules,omitnil" name:"Rules"`

	// 添加规则的来源，一般不需要使用，值insert_rule 表示插入指定位置的规则；值batch_import 表示批量导入规则；为空时表示添加规则
	From *string `json:"From,omitnil" name:"From"`
}

type AddVpcAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要添加的vpc内网间规则列表
	Rules []*VpcRuleItem `json:"Rules,omitnil" name:"Rules"`

	// 添加规则的来源，一般不需要使用，值insert_rule 表示插入指定位置的规则；值batch_import 表示批量导入规则；为空时表示添加规则
	From *string `json:"From,omitnil" name:"From"`
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
	RuleUuids []*int64 `json:"RuleUuids,omitnil" name:"RuleUuids"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 地域英文
	ZoneEng *string `json:"ZoneEng,omitnil" name:"ZoneEng"`
}

type AssociatedInstanceInfo struct {
	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 实例类型，3是cvm实例,4是clb实例,5是eni实例,6是云数据库
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 私有网络名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// 公网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 关联安全组数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroupCount *uint64 `json:"SecurityGroupCount,omitnil" name:"SecurityGroupCount"`

	// 关联安全组规则数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroupRuleCount *uint64 `json:"SecurityGroupRuleCount,omitnil" name:"SecurityGroupRuleCount"`

	// 关联数据库代理Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdbId *string `json:"CdbId,omitnil" name:"CdbId"`
}

type BetaInfoByACL struct {
	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *int64 `json:"TaskId,omitnil" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 上次执行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTime *string `json:"LastTime,omitnil" name:"LastTime"`
}

type BlockIgnoreRule struct {
	// 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	IP *string `json:"IP,omitnil" name:"IP"`

	// 规则ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ioc *string `json:"Ioc,omitnil" name:"Ioc"`

	// 危险等级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *string `json:"Level,omitnil" name:"Level"`

	// 来源事件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventName *string `json:"EventName,omitnil" name:"EventName"`

	// 方向：1入站，0出站
	// 注意：此字段可能返回 null，表示取不到有效值。
	Direction *int64 `json:"Direction,omitnil" name:"Direction"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 地理位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitnil" name:"Address"`

	// 规则类型：1封禁，2放通
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *int64 `json:"Action,omitnil" name:"Action"`

	// 规则生效开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 规则生效结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 忽略原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreReason *string `json:"IgnoreReason,omitnil" name:"IgnoreReason"`

	// 安全事件来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil" name:"Source"`

	// 规则id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqueId *string `json:"UniqueId,omitnil" name:"UniqueId"`

	// 规则命中次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchTimes *int64 `json:"MatchTimes,omitnil" name:"MatchTimes"`

	// 国家
	// 注意：此字段可能返回 null，表示取不到有效值。
	Country *string `json:"Country,omitnil" name:"Country"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitnil" name:"Comment"`
}

type CfwNatDnatRule struct {
	// 网络协议，可选值：TCP、UDP。
	IpProtocol *string `json:"IpProtocol,omitnil" name:"IpProtocol"`

	// 弹性IP。
	PublicIpAddress *string `json:"PublicIpAddress,omitnil" name:"PublicIpAddress"`

	// 公网端口。
	PublicPort *int64 `json:"PublicPort,omitnil" name:"PublicPort"`

	// 内网地址。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitnil" name:"PrivateIpAddress"`

	// 内网端口。
	PrivatePort *int64 `json:"PrivatePort,omitnil" name:"PrivatePort"`

	// NAT防火墙转发规则描述。
	Description *string `json:"Description,omitnil" name:"Description"`
}

type CommonFilter struct {
	// 检索的键值
	Name *string `json:"Name,omitnil" name:"Name"`

	// 检索的值，各检索值间为OR关系
	Values []*string `json:"Values,omitnil" name:"Values"`

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
	OperatorType *int64 `json:"OperatorType,omitnil" name:"OperatorType"`
}

// Predefined struct for user
type CreateAcRulesRequestParams struct {
	// 创建规则数据
	Data []*RuleInfoData `json:"Data,omitnil" name:"Data"`

	// 0：添加（默认），1：插入
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 边id
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// 访问控制规则状态
	Enable *int64 `json:"Enable,omitnil" name:"Enable"`

	// 0：添加，1：覆盖
	Overwrite *uint64 `json:"Overwrite,omitnil" name:"Overwrite"`

	// NAT实例ID, 参数Area存在的时候这个必传
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// portScan: 来自于端口扫描, patchImport: 来自于批量导入
	From *string `json:"From,omitnil" name:"From"`

	// NAT地域
	Area *string `json:"Area,omitnil" name:"Area"`
}

type CreateAcRulesRequest struct {
	*tchttp.BaseRequest
	
	// 创建规则数据
	Data []*RuleInfoData `json:"Data,omitnil" name:"Data"`

	// 0：添加（默认），1：插入
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 边id
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// 访问控制规则状态
	Enable *int64 `json:"Enable,omitnil" name:"Enable"`

	// 0：添加，1：覆盖
	Overwrite *uint64 `json:"Overwrite,omitnil" name:"Overwrite"`

	// NAT实例ID, 参数Area存在的时候这个必传
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// portScan: 来自于端口扫描, patchImport: 来自于批量导入
	From *string `json:"From,omitnil" name:"From"`

	// NAT地域
	Area *string `json:"Area,omitnil" name:"Area"`
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
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 返回多余的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Info *string `json:"Info,omitnil" name:"Info"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板描述
	Detail *string `json:"Detail,omitnil" name:"Detail"`

	// Type为1，ip模板eg：1.1.1.1,2.2.2.2；
	// Type为5，域名模板eg：www.qq.com,www.tencent.com
	IpString *string `json:"IpString,omitnil" name:"IpString"`

	// 1 ip模板
	// 5 域名模板
	Type *int64 `json:"Type,omitnil" name:"Type"`
}

type CreateAddressTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板描述
	Detail *string `json:"Detail,omitnil" name:"Detail"`

	// Type为1，ip模板eg：1.1.1.1,2.2.2.2；
	// Type为5，域名模板eg：www.qq.com,www.tencent.com
	IpString *string `json:"IpString,omitnil" name:"IpString"`

	// 1 ip模板
	// 5 域名模板
	Type *int64 `json:"Type,omitnil" name:"Type"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAddressTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAddressTemplateResponseParams struct {
	// 创建结果,0成功
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 唯一Id
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type CreateBlockIgnoreRuleListRequestParams struct {
	// 规则列表
	Rules []*IntrusionDefenseRule `json:"Rules,omitnil" name:"Rules"`

	// 规则类型，1封禁，2放通，不支持域名封禁
	RuleType *int64 `json:"RuleType,omitnil" name:"RuleType"`

	// 是否覆盖重复数据，1覆盖，非1不覆盖，跳过重复数据
	CoverDuplicate *int64 `json:"CoverDuplicate,omitnil" name:"CoverDuplicate"`
}

type CreateBlockIgnoreRuleListRequest struct {
	*tchttp.BaseRequest
	
	// 规则列表
	Rules []*IntrusionDefenseRule `json:"Rules,omitnil" name:"Rules"`

	// 规则类型，1封禁，2放通，不支持域名封禁
	RuleType *int64 `json:"RuleType,omitnil" name:"RuleType"`

	// 是否覆盖重复数据，1覆盖，非1不覆盖，跳过重复数据
	CoverDuplicate *int64 `json:"CoverDuplicate,omitnil" name:"CoverDuplicate"`
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
	List []*IocListData `json:"List,omitnil" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type CreateChooseVpcsRequestParams struct {
	// vpc列表
	VpcList []*string `json:"VpcList,omitnil" name:"VpcList"`

	// zone列表
	AllZoneList []*VpcZoneData `json:"AllZoneList,omitnil" name:"AllZoneList"`
}

type CreateChooseVpcsRequest struct {
	*tchttp.BaseRequest
	
	// vpc列表
	VpcList []*string `json:"VpcList,omitnil" name:"VpcList"`

	// zone列表
	AllZoneList []*VpcZoneData `json:"AllZoneList,omitnil" name:"AllZoneList"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	DatabaseWhiteListRuleData []*DatabaseWhiteListRuleData `json:"DatabaseWhiteListRuleData,omitnil" name:"DatabaseWhiteListRuleData"`
}

type CreateDatabaseWhiteListRulesRequest struct {
	*tchttp.BaseRequest
	
	// 创建白名单数据
	DatabaseWhiteListRuleData []*DatabaseWhiteListRuleData `json:"DatabaseWhiteListRuleData,omitnil" name:"DatabaseWhiteListRuleData"`
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
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 带宽
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 模式 1：接入模式；0：新增模式
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 新增模式传递参数，其中NewModeItems和NatgwList至少传递一种。
	NewModeItems *NewModeItems `json:"NewModeItems,omitnil" name:"NewModeItems"`

	// 接入模式接入的nat网关列表，其中NewModeItems和NatgwList至少传递一种。
	NatGwList []*string `json:"NatGwList,omitnil" name:"NatGwList"`

	// 主可用区，为空则选择默认可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 备可用区，为空则选择默认可用区
	ZoneBak *string `json:"ZoneBak,omitnil" name:"ZoneBak"`

	// 异地灾备 1：使用异地灾备；0：不使用异地灾备；为空则默认不使用异地灾备
	CrossAZone *int64 `json:"CrossAZone,omitnil" name:"CrossAZone"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil" name:"FwCidrInfo"`
}

type CreateNatFwInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙实例名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 带宽
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 模式 1：接入模式；0：新增模式
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 新增模式传递参数，其中NewModeItems和NatgwList至少传递一种。
	NewModeItems *NewModeItems `json:"NewModeItems,omitnil" name:"NewModeItems"`

	// 接入模式接入的nat网关列表，其中NewModeItems和NatgwList至少传递一种。
	NatGwList []*string `json:"NatGwList,omitnil" name:"NatGwList"`

	// 主可用区，为空则选择默认可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 备可用区，为空则选择默认可用区
	ZoneBak *string `json:"ZoneBak,omitnil" name:"ZoneBak"`

	// 异地灾备 1：使用异地灾备；0：不使用异地灾备；为空则默认不使用异地灾备
	CrossAZone *int64 `json:"CrossAZone,omitnil" name:"CrossAZone"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil" name:"FwCidrInfo"`
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
	CfwInsId *string `json:"CfwInsId,omitnil" name:"CfwInsId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 带宽
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 模式 1：接入模式；0：新增模式
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 新增模式传递参数，其中NewModeItems和NatgwList至少传递一种。
	NewModeItems *NewModeItems `json:"NewModeItems,omitnil" name:"NewModeItems"`

	// 接入模式接入的nat网关列表，其中NewModeItems和NatgwList至少传递一种。
	NatGwList []*string `json:"NatGwList,omitnil" name:"NatGwList"`

	// 主可用区，为空则选择默认可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 备可用区，为空则选择默认可用区
	ZoneBak *string `json:"ZoneBak,omitnil" name:"ZoneBak"`

	// 异地灾备 1：使用异地灾备；0：不使用异地灾备；为空则默认不使用异地灾备
	CrossAZone *int64 `json:"CrossAZone,omitnil" name:"CrossAZone"`

	// 0不创建域名,1创建域名
	IsCreateDomain *int64 `json:"IsCreateDomain,omitnil" name:"IsCreateDomain"`

	// 如果要创建域名则必填
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil" name:"FwCidrInfo"`
}

type CreateNatFwInstanceWithDomainRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙实例名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 带宽
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 模式 1：接入模式；0：新增模式
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 新增模式传递参数，其中NewModeItems和NatgwList至少传递一种。
	NewModeItems *NewModeItems `json:"NewModeItems,omitnil" name:"NewModeItems"`

	// 接入模式接入的nat网关列表，其中NewModeItems和NatgwList至少传递一种。
	NatGwList []*string `json:"NatGwList,omitnil" name:"NatGwList"`

	// 主可用区，为空则选择默认可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 备可用区，为空则选择默认可用区
	ZoneBak *string `json:"ZoneBak,omitnil" name:"ZoneBak"`

	// 异地灾备 1：使用异地灾备；0：不使用异地灾备；为空则默认不使用异地灾备
	CrossAZone *int64 `json:"CrossAZone,omitnil" name:"CrossAZone"`

	// 0不创建域名,1创建域名
	IsCreateDomain *int64 `json:"IsCreateDomain,omitnil" name:"IsCreateDomain"`

	// 如果要创建域名则必填
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil" name:"FwCidrInfo"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	CfwInsId *string `json:"CfwInsId,omitnil" name:"CfwInsId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	SourceContent *string `json:"SourceContent,omitnil" name:"SourceContent"`

	// 访问源类型：入向规则时类型可以为 ip,net,template,location；出向规则时可以为 ip,net,template,instance,group,tag
	SourceType *string `json:"SourceType,omitnil" name:"SourceType"`

	// 访问目的示例： net：IP/CIDR(192.168.0.2) domain：域名规则，例如*.qq.com
	TargetContent *string `json:"TargetContent,omitnil" name:"TargetContent"`

	// 访问目的类型：入向规则时类型可以为ip,net,template,instance,group,tag；出向规则时可以为  ip,net,domain,template,location
	TargetType *string `json:"TargetType,omitnil" name:"TargetType"`

	// 协议，可选的值： TCP UDP ICMP ANY HTTP HTTPS HTTP/HTTPS SMTP SMTPS SMTP/SMTPS FTP DNS
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 访问控制策略中设置的流量通过云防火墙的方式。取值： accept：放行 drop：拒绝 log：观察
	RuleAction *string `json:"RuleAction,omitnil" name:"RuleAction"`

	// 访问控制策略的端口。取值： -1/-1：全部端口 80：80端口
	Port *string `json:"Port,omitnil" name:"Port"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 规则序号
	OrderIndex *int64 `json:"OrderIndex,omitnil" name:"OrderIndex"`

	// 规则状态，true表示启用，false表示禁用
	Enable *string `json:"Enable,omitnil" name:"Enable"`

	// 规则对应的唯一id，创建规则时无需填写
	Uuid *int64 `json:"Uuid,omitnil" name:"Uuid"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`
}

type CreateRuleItem struct {
	// 访问源示例： net：IP/CIDR(192.168.0.2)
	SourceContent *string `json:"SourceContent,omitnil" name:"SourceContent"`

	// 访问源类型：入向规则时类型可以为 ip,net,template,location；出向规则时可以为 ip,net,template,instance,group,tag
	SourceType *string `json:"SourceType,omitnil" name:"SourceType"`

	// 访问目的示例： net：IP/CIDR(192.168.0.2) domain：域名规则，例如*.qq.com
	TargetContent *string `json:"TargetContent,omitnil" name:"TargetContent"`

	// 访问目的类型：入向规则时类型可以为ip,net,template,instance,group,tag；出向规则时可以为  ip,net,domain,template,location
	TargetType *string `json:"TargetType,omitnil" name:"TargetType"`

	// 协议，可选的值： TCP UDP ICMP ANY HTTP HTTPS HTTP/HTTPS SMTP SMTPS SMTP/SMTPS FTP DNS
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 访问控制策略中设置的流量通过云防火墙的方式。取值： accept：放行 drop：拒绝 log：观察
	RuleAction *string `json:"RuleAction,omitnil" name:"RuleAction"`

	// 访问控制策略的端口。取值： -1/-1：全部端口 80：80端口
	Port *string `json:"Port,omitnil" name:"Port"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 规则序号
	OrderIndex *int64 `json:"OrderIndex,omitnil" name:"OrderIndex"`

	// 规则对应的唯一id，创建规则时无需填写
	Uuid *int64 `json:"Uuid,omitnil" name:"Uuid"`

	// 规则状态，true表示启用，false表示禁用
	Enable *string `json:"Enable,omitnil" name:"Enable"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// all
	Scope *string `json:"Scope,omitnil" name:"Scope"`

	// 0，正常规则添加；1，入侵检测添加
	RuleSource *int64 `json:"RuleSource,omitnil" name:"RuleSource"`

	// 告警Id
	LogId *string `json:"LogId,omitnil" name:"LogId"`
}

// Predefined struct for user
type CreateSecurityGroupRulesRequestParams struct {
	// 添加的企业安全组规则数据
	Data []*SecurityGroupListData `json:"Data,omitnil" name:"Data"`

	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 0：后插，1：前插，2：中插，默认0
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 添加后是否启用规则，0：不启用，1：启用，默认1
	Enable *uint64 `json:"Enable,omitnil" name:"Enable"`
}

type CreateSecurityGroupRulesRequest struct {
	*tchttp.BaseRequest
	
	// 添加的企业安全组规则数据
	Data []*SecurityGroupListData `json:"Data,omitnil" name:"Data"`

	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 0：后插，1：前插，2：中插，默认0
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// 添加后是否启用规则，0：不启用，1：启用，默认1
	Enable *uint64 `json:"Enable,omitnil" name:"Enable"`
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
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模式 1：CCN云联网模式；0：私有网络模式 2: sase 模式 3：ccn 高级模式 4: 私有网络(跨租户单边模式)
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 防火墙(组)下的防火墙实例列表
	VpcFwInstances []*VpcFwInstance `json:"VpcFwInstances,omitnil" name:"VpcFwInstances"`

	// 防火墙实例的开关模式
	// 1: 单点互通
	// 2: 多点互通
	// 3: 全互通
	// 4: 自定义路由
	SwitchMode *int64 `json:"SwitchMode,omitnil" name:"SwitchMode"`

	// auto 自动选择防火墙网段
	// 10.10.10.0/24 用户输入的防火墙网段
	FwVpcCidr *string `json:"FwVpcCidr,omitnil" name:"FwVpcCidr"`

	// 云联网id ，适用于云联网模式
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil" name:"FwCidrInfo"`

	// 跨租户管理员模式  1管理员 2多账号
	CrossUserMode *string `json:"CrossUserMode,omitnil" name:"CrossUserMode"`
}

type CreateVpcFwGroupRequest struct {
	*tchttp.BaseRequest
	
	// VPC防火墙(组)名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模式 1：CCN云联网模式；0：私有网络模式 2: sase 模式 3：ccn 高级模式 4: 私有网络(跨租户单边模式)
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 防火墙(组)下的防火墙实例列表
	VpcFwInstances []*VpcFwInstance `json:"VpcFwInstances,omitnil" name:"VpcFwInstances"`

	// 防火墙实例的开关模式
	// 1: 单点互通
	// 2: 多点互通
	// 3: 全互通
	// 4: 自定义路由
	SwitchMode *int64 `json:"SwitchMode,omitnil" name:"SwitchMode"`

	// auto 自动选择防火墙网段
	// 10.10.10.0/24 用户输入的防火墙网段
	FwVpcCidr *string `json:"FwVpcCidr,omitnil" name:"FwVpcCidr"`

	// 云联网id ，适用于云联网模式
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil" name:"FwCidrInfo"`

	// 跨租户管理员模式  1管理员 2多账号
	CrossUserMode *string `json:"CrossUserMode,omitnil" name:"CrossUserMode"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwGroupId *string `json:"FwGroupId,omitnil" name:"FwGroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type DatabaseWhiteListRuleData struct {
	// 访问源
	SourceIp *string `json:"SourceIp,omitnil" name:"SourceIp"`

	// 访问源类型，1 ip；6 实例；100 资源分组
	SourceType *int64 `json:"SourceType,omitnil" name:"SourceType"`

	// 访问目的
	TargetIp *string `json:"TargetIp,omitnil" name:"TargetIp"`

	// 访问目的类型，1 ip；6 实例；100 资源分组
	TargetType *int64 `json:"TargetType,omitnil" name:"TargetType"`

	// 规则描述
	Detail *string `json:"Detail,omitnil" name:"Detail"`

	// 是否地域规则，0不是 1是
	IsRegionRule *int64 `json:"IsRegionRule,omitnil" name:"IsRegionRule"`

	// 是否云厂商规则，0不是 1 时
	IsCloudRule *int64 `json:"IsCloudRule,omitnil" name:"IsCloudRule"`

	// 是否启用，0 不启用，1启用
	Enable *int64 `json:"Enable,omitnil" name:"Enable"`

	// 地域码1
	FirstLevelRegionCode *int64 `json:"FirstLevelRegionCode,omitnil" name:"FirstLevelRegionCode"`

	// 地域码2
	SecondLevelRegionCode *int64 `json:"SecondLevelRegionCode,omitnil" name:"SecondLevelRegionCode"`

	// 地域名称1
	FirstLevelRegionName *string `json:"FirstLevelRegionName,omitnil" name:"FirstLevelRegionName"`

	// 地域名称2
	SecondLevelRegionName *string `json:"SecondLevelRegionName,omitnil" name:"SecondLevelRegionName"`

	// 云厂商码
	CloudCode *string `json:"CloudCode,omitnil" name:"CloudCode"`
}

// Predefined struct for user
type DeleteAcRuleRequestParams struct {
	// 删除规则对应的id值, 对应获取规则列表接口的Id 值
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// EdgeId值两个vpc间的边id
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// NAT地域， 如ap-shanghai/ap-guangzhou/ap-chongqing等
	Area *string `json:"Area,omitnil" name:"Area"`
}

type DeleteAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 删除规则对应的id值, 对应获取规则列表接口的Id 值
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// EdgeId值两个vpc间的边id
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// NAT地域， 如ap-shanghai/ap-guangzhou/ap-chongqing等
	Area *string `json:"Area,omitnil" name:"Area"`
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
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 返回多余的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Info *string `json:"Info,omitnil" name:"Info"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

type DeleteAddressTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板id
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
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
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DeleteAllAccessControlRuleRequestParams struct {
	// 方向，0：出站，1：入站  默认值是 0
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// VPC间防火墙开关ID  全部删除 EdgeId和Area只填写一个，不填写则不删除vpc间防火墙开关 ，默认值为‘’
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// nat地域 全部删除 EdgeId和Area只填写一个，不填写则不删除nat防火墙开关 默认值为‘’
	Area *string `json:"Area,omitnil" name:"Area"`
}

type DeleteAllAccessControlRuleRequest struct {
	*tchttp.BaseRequest
	
	// 方向，0：出站，1：入站  默认值是 0
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// VPC间防火墙开关ID  全部删除 EdgeId和Area只填写一个，不填写则不删除vpc间防火墙开关 ，默认值为‘’
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// nat地域 全部删除 EdgeId和Area只填写一个，不填写则不删除nat防火墙开关 默认值为‘’
	Area *string `json:"Area,omitnil" name:"Area"`
}

func (r *DeleteAllAccessControlRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAllAccessControlRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "EdgeId")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAllAccessControlRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAllAccessControlRuleResponseParams struct {
	// 状态值 0: 修改成功, 非0: 修改失败
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 删除了几条访问控制规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Info *int64 `json:"Info,omitnil" name:"Info"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAllAccessControlRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAllAccessControlRuleResponseParams `json:"Response"`
}

func (r *DeleteAllAccessControlRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAllAccessControlRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBlockIgnoreRuleListRequestParams struct {
	// 规则列表
	Rules []*IocListData `json:"Rules,omitnil" name:"Rules"`

	// 规则类型，1封禁，2放通，不支持域名封禁
	RuleType *int64 `json:"RuleType,omitnil" name:"RuleType"`
}

type DeleteBlockIgnoreRuleListRequest struct {
	*tchttp.BaseRequest
	
	// 规则列表
	Rules []*IocListData `json:"Rules,omitnil" name:"Rules"`

	// 规则类型，1封禁，2放通，不支持域名封禁
	RuleType *int64 `json:"RuleType,omitnil" name:"RuleType"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DeleteNatFwInstanceRequestParams struct {
	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitnil" name:"CfwInstance"`
}

type DeleteNatFwInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitnil" name:"CfwInstance"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DeleteResourceGroupRequestParams struct {
	// 组id
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`
}

type DeleteResourceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 组id
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 腾讯云地域的英文简写
	Area *string `json:"Area,omitnil" name:"Area"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 是否删除反向规则，0：否，1：是
	IsDelReverse *uint64 `json:"IsDelReverse,omitnil" name:"IsDelReverse"`
}

type DeleteSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// 所需要删除规则的ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 腾讯云地域的英文简写
	Area *string `json:"Area,omitnil" name:"Area"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 是否删除反向规则，0：否，1：是
	IsDelReverse *uint64 `json:"IsDelReverse,omitnil" name:"IsDelReverse"`
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
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 返回多余的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Info *string `json:"Info,omitnil" name:"Info"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	FwGroupId *string `json:"FwGroupId,omitnil" name:"FwGroupId"`

	// 是否删除整个防火墙(组)
	// 0：不删除防火墙(组)，只删除单独实例
	// 1：删除整个防火墙(组)
	DeleteFwGroup *int64 `json:"DeleteFwGroup,omitnil" name:"DeleteFwGroup"`

	// 待删除的防火墙实例数组
	VpcFwInsList []*string `json:"VpcFwInsList,omitnil" name:"VpcFwInsList"`
}

type DeleteVpcFwGroupRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙(组)Id
	FwGroupId *string `json:"FwGroupId,omitnil" name:"FwGroupId"`

	// 是否删除整个防火墙(组)
	// 0：不删除防火墙(组)，只删除单独实例
	// 1：删除整个防火墙(组)
	DeleteFwGroup *int64 `json:"DeleteFwGroup,omitnil" name:"DeleteFwGroup"`

	// 待删除的防火墙实例数组
	VpcFwInsList []*string `json:"VpcFwInsList,omitnil" name:"VpcFwInsList"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

// Predefined struct for user
type DeleteVpcInstanceRequestParams struct {

}

type DeleteVpcInstanceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DeleteVpcInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVpcInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVpcInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteVpcInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVpcInstanceResponseParams `json:"Response"`
}

func (r *DeleteVpcInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVpcInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescAcItem struct {
	// 访问源
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceContent *string `json:"SourceContent,omitnil" name:"SourceContent"`

	// 访问目的
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetContent *string `json:"TargetContent,omitnil" name:"TargetContent"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitnil" name:"Port"`

	// 访问控制策略中设置的流量通过云防火墙的方式。取值： accept：放行 drop：拒绝 log：观察
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleAction *string `json:"RuleAction,omitnil" name:"RuleAction"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 命中次数
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitnil" name:"OrderIndex"`

	// 访问源类型：入向规则时类型可以为 ip,net,template,location；出向规则时可以为 ip,net,template,instance,group,tag
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceType *string `json:"SourceType,omitnil" name:"SourceType"`

	// 访问目的类型：入向规则时类型可以为ip,net,template,instance,group,tag；出向规则时可以为 ip,net,domain,template,location
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetType *string `json:"TargetType,omitnil" name:"TargetType"`

	// 规则对应的唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uuid *uint64 `json:"Uuid,omitnil" name:"Uuid"`

	// 规则有效性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Invalid *uint64 `json:"Invalid,omitnil" name:"Invalid"`

	// 0为正常规则,1为地域规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsRegion *uint64 `json:"IsRegion,omitnil" name:"IsRegion"`

	// 国家id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountryCode *uint64 `json:"CountryCode,omitnil" name:"CountryCode"`

	// 城市id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CityCode *uint64 `json:"CityCode,omitnil" name:"CityCode"`

	// 国家名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountryName *string `json:"CountryName,omitnil" name:"CountryName"`

	// 省名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CityName *string `json:"CityName,omitnil" name:"CityName"`

	// 云厂商code
	// 注意：此字段可能返回 null，表示取不到有效值。
	CloudCode *string `json:"CloudCode,omitnil" name:"CloudCode"`

	// 0为正常规则,1为云厂商规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCloud *uint64 `json:"IsCloud,omitnil" name:"IsCloud"`

	// 规则状态，true表示启用，false表示禁用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *string `json:"Enable,omitnil" name:"Enable"`

	// 规则方向：1，入向；0，出向
	// 注意：此字段可能返回 null，表示取不到有效值。
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 内部使用的uuid，一般情况下不会使用到该字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternalUuid *int64 `json:"InternalUuid,omitnil" name:"InternalUuid"`

	// 规则状态，查询规则命中详情时该字段有效，0：新增，1: 已删除, 2: 编辑删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 关联任务详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	BetaList []*BetaInfoByACL `json:"BetaList,omitnil" name:"BetaList"`

	// 生效范围：serial，串行；side，旁路；all，全局
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scope *string `json:"Scope,omitnil" name:"Scope"`

	// 互联网边界防火墙使用的内部规则id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetBorderUuid *string `json:"InternetBorderUuid,omitnil" name:"InternetBorderUuid"`
}

// Predefined struct for user
type DescribeAcListsRequestParams struct {
	// 协议
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 策略
	Strategy *string `json:"Strategy,omitnil" name:"Strategy"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`

	// 每页条数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 出站还是入站，1：入站，0：出站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// EdgeId值
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// 规则是否开启，'0': 未开启，'1': 开启, 默认为'0'
	Status *string `json:"Status,omitnil" name:"Status"`

	// 地域
	Area *string `json:"Area,omitnil" name:"Area"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DescribeAcListsRequest struct {
	*tchttp.BaseRequest
	
	// 协议
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 策略
	Strategy *string `json:"Strategy,omitnil" name:"Strategy"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`

	// 每页条数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 出站还是入站，1：入站，0：出站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// EdgeId值
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// 规则是否开启，'0': 未开启，'1': 开启, 默认为'0'
	Status *string `json:"Status,omitnil" name:"Status"`

	// 地域
	Area *string `json:"Area,omitnil" name:"Area"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
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
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 访问控制列表数据
	Data []*AcListsData `json:"Data,omitnil" name:"Data"`

	// 不算筛选条数的总条数
	AllTotal *uint64 `json:"AllTotal,omitnil" name:"AllTotal"`

	// 访问控制规则全部启用/全部停用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *uint64 `json:"Enable,omitnil" name:"Enable"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要查询的索引，特定场景使用，可不填
	Index *string `json:"Index,omitnil" name:"Index"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// 每页条数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要查询的索引，特定场景使用，可不填
	Index *string `json:"Index,omitnil" name:"Index"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil" name:"By"`
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
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// nat访问控制列表数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DescAcItem `json:"Data,omitnil" name:"Data"`

	// 未过滤的总条数
	AllTotal *uint64 `json:"AllTotal,omitnil" name:"AllTotal"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 条数，分页用
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段，取值 'UpdateTime' | 'RulesNum'
	By *string `json:"By,omitnil" name:"By"`

	// 排序，取值 'asc'|'desc'
	Order *string `json:"Order,omitnil" name:"Order"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`

	// 检索地址模板唯一id
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
}

type DescribeAddressTemplateListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，分页用
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 条数，分页用
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 排序字段，取值 'UpdateTime' | 'RulesNum'
	By *string `json:"By,omitnil" name:"By"`

	// 排序，取值 'asc'|'desc'
	Order *string `json:"Order,omitnil" name:"Order"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`

	// 检索地址模板唯一id
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAddressTemplateListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAddressTemplateListResponseParams struct {
	// 模板总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 模板列表数据
	Data []*TemplateListInfo `json:"Data,omitnil" name:"Data"`

	// 模板名称列表
	NameList []*string `json:"NameList,omitnil" name:"NameList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// success 成功
	// 其他失败
	ReturnMsg *string `json:"ReturnMsg,omitnil" name:"ReturnMsg"`

	// 0 成功
	// 非0 失败
	ReturnCode *uint64 `json:"ReturnCode,omitnil" name:"ReturnCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页记录条数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 地域代码（例：ap-guangzhou）,支持腾讯云全地域
	Area *string `json:"Area,omitnil" name:"Area"`

	// 额外检索条件（JSON字符串）
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`

	// 排序字段
	By *string `json:"By,omitnil" name:"By"`

	// 排序方式（asc:升序,desc:降序）
	Order *string `json:"Order,omitnil" name:"Order"`

	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 实例类型,'3'是cvm实例,'4'是clb实例,'5'是eni实例,'6'是云数据库
	Type *string `json:"Type,omitnil" name:"Type"`
}

type DescribeAssociatedInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 列表偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页记录条数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 地域代码（例：ap-guangzhou）,支持腾讯云全地域
	Area *string `json:"Area,omitnil" name:"Area"`

	// 额外检索条件（JSON字符串）
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`

	// 排序字段
	By *string `json:"By,omitnil" name:"By"`

	// 排序方式（asc:升序,desc:降序）
	Order *string `json:"Order,omitnil" name:"Order"`

	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitnil" name:"SecurityGroupId"`

	// 实例类型,'3'是cvm实例,'4'是clb实例,'5'是eni实例,'6'是云数据库
	Type *string `json:"Type,omitnil" name:"Type"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*AssociatedInstanceInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// ip查询条件
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 地域
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 方向
	Direction *string `json:"Direction,omitnil" name:"Direction"`

	// 来源
	Source *string `json:"Source,omitnil" name:"Source"`

	// vpc间防火墙开关边id
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// 日志来源 move：vpc间防火墙
	LogSource *string `json:"LogSource,omitnil" name:"LogSource"`
}

type DescribeBlockByIpTimesListRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// ip查询条件
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 地域
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 方向
	Direction *string `json:"Direction,omitnil" name:"Direction"`

	// 来源
	Source *string `json:"Source,omitnil" name:"Source"`

	// vpc间防火墙开关边id
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// 日志来源 move：vpc间防火墙
	LogSource *string `json:"LogSource,omitnil" name:"LogSource"`
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
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Ip")
	delete(f, "Zone")
	delete(f, "Direction")
	delete(f, "Source")
	delete(f, "EdgeId")
	delete(f, "LogSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlockByIpTimesListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockByIpTimesListResponseParams struct {
	// 返回数据
	Data []*IpStatic `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 页偏移量
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 方向：1互联网入站，0互联网出站，3内网，空 全部方向
	Direction *string `json:"Direction,omitnil" name:"Direction"`

	// 规则类型：1封禁，2放通
	RuleType *uint64 `json:"RuleType,omitnil" name:"RuleType"`

	// 排序类型：desc降序，asc正序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序列：EndTime结束时间，StartTime开始时间，MatchTimes命中次数
	By *string `json:"By,omitnil" name:"By"`

	// 搜索参数，json格式字符串，空则传"{}"，域名：domain，危险等级：level，放通原因：ignore_reason，安全事件来源：rule_source，地理位置：address，模糊搜索：common
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`
}

type DescribeBlockIgnoreListRequest struct {
	*tchttp.BaseRequest
	
	// 单页数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 页偏移量
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 方向：1互联网入站，0互联网出站，3内网，空 全部方向
	Direction *string `json:"Direction,omitnil" name:"Direction"`

	// 规则类型：1封禁，2放通
	RuleType *uint64 `json:"RuleType,omitnil" name:"RuleType"`

	// 排序类型：desc降序，asc正序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序列：EndTime结束时间，StartTime开始时间，MatchTimes命中次数
	By *string `json:"By,omitnil" name:"By"`

	// 搜索参数，json格式字符串，空则传"{}"，域名：domain，危险等级：level，放通原因：ignore_reason，安全事件来源：rule_source，地理位置：address，模糊搜索：common
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`
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
	delete(f, "RuleType")
	delete(f, "Order")
	delete(f, "By")
	delete(f, "SearchValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlockIgnoreListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockIgnoreListResponseParams struct {
	// 列表数据
	Data []*BlockIgnoreRule `json:"Data,omitnil" name:"Data"`

	// 查询结果总数，用于分页
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 状态值，0：查询成功，非0：查询失败
	ReturnCode *int64 `json:"ReturnCode,omitnil" name:"ReturnCode"`

	// 状态信息，success：查询成功，fail：查询失败
	ReturnMsg *string `json:"ReturnMsg,omitnil" name:"ReturnMsg"`

	// 安全事件来源下拉框
	SourceList []*string `json:"SourceList,omitnil" name:"SourceList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 列表类型，只能是下面三种之一：port、address、ip
	QueryType *string `json:"QueryType,omitnil" name:"QueryType"`

	// top数
	Top *int64 `json:"Top,omitnil" name:"Top"`

	// 查询条件
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`
}

type DescribeBlockStaticListRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 列表类型，只能是下面三种之一：port、address、ip
	QueryType *string `json:"QueryType,omitnil" name:"QueryType"`

	// top数
	Top *int64 `json:"Top,omitnil" name:"Top"`

	// 查询条件
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`
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
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QueryType")
	delete(f, "Top")
	delete(f, "SearchValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlockStaticListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockStaticListResponseParams struct {
	// 无
	Data []*StaticInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeCfwEipsRequestParams struct {
	// 1：cfw接入模式，目前仅支持接入模式实例
	Mode *uint64 `json:"Mode,omitnil" name:"Mode"`

	// ALL：查询所有弹性公网ip; nat-xxxxx：接入模式场景指定网关的弹性公网ip
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// 防火墙实例id，当前仅支持接入模式的实例
	CfwInstance *string `json:"CfwInstance,omitnil" name:"CfwInstance"`
}

type DescribeCfwEipsRequest struct {
	*tchttp.BaseRequest
	
	// 1：cfw接入模式，目前仅支持接入模式实例
	Mode *uint64 `json:"Mode,omitnil" name:"Mode"`

	// ALL：查询所有弹性公网ip; nat-xxxxx：接入模式场景指定网关的弹性公网ip
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// 防火墙实例id，当前仅支持接入模式的实例
	CfwInstance *string `json:"CfwInstance,omitnil" name:"CfwInstance"`
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
	NatFwEipList []*NatFwEipsInfo `json:"NatFwEipList,omitnil" name:"NatFwEipList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	BasicRuleSwitch *int64 `json:"BasicRuleSwitch,omitnil" name:"BasicRuleSwitch"`

	// 安全基线开关
	BaselineAllSwitch *int64 `json:"BaselineAllSwitch,omitnil" name:"BaselineAllSwitch"`

	// 威胁情报开关
	TiSwitch *int64 `json:"TiSwitch,omitnil" name:"TiSwitch"`

	// 虚拟补丁开关
	VirtualPatchSwitch *int64 `json:"VirtualPatchSwitch,omitnil" name:"VirtualPatchSwitch"`

	// 是否历史开启
	HistoryOpen *int64 `json:"HistoryOpen,omitnil" name:"HistoryOpen"`

	// 状态值，0：查询成功，非0：查询失败
	ReturnCode *int64 `json:"ReturnCode,omitnil" name:"ReturnCode"`

	// 状态信息，success：查询成功，fail：查询失败
	ReturnMsg *string `json:"ReturnMsg,omitnil" name:"ReturnMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Progress *int64 `json:"Progress,omitnil" name:"Progress"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeEnterpriseSecurityGroupRuleRequestParams struct {
	// 分页查询时，显示的当前页的页码。
	// 
	// 默认值为1。
	PageNo *string `json:"PageNo,omitnil" name:"PageNo"`

	// 分页查询时，显示的每页数据的最大条数。
	// 
	// 可设置值最大为50。
	PageSize *string `json:"PageSize,omitnil" name:"PageSize"`

	// 访问源示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板(ipm-dyodhpby)
	// instance：资产实例(ins-123456)
	// resourcegroup：资产分组(/全部分组/分组1/子分组1)
	// tag：资源标签({"Key":"标签key值","Value":"标签Value值"})
	// region：地域(ap-gaungzhou)
	// 支持通配
	SourceContent *string `json:"SourceContent,omitnil" name:"SourceContent"`

	// 访问目的示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板(ipm-dyodhpby)
	// instance：资产实例(ins-123456)
	// resourcegroup：资产分组(/全部分组/分组1/子分组1)
	// tag：资源标签({"Key":"标签key值","Value":"标签Value值"})
	// region：地域(ap-gaungzhou)
	// 支持通配
	DestContent *string `json:"DestContent,omitnil" name:"DestContent"`

	// 规则描述，支持通配
	Description *string `json:"Description,omitnil" name:"Description"`

	// 访问控制策略中设置的流量通过云防火墙的方式。取值：
	// accept：放行
	// drop：拒绝
	RuleAction *string `json:"RuleAction,omitnil" name:"RuleAction"`

	// 是否启用规则，默认为启用，取值：
	// true为启用，false为不启用
	Enable *string `json:"Enable,omitnil" name:"Enable"`

	// 访问控制策略的端口。取值：
	// -1/-1：全部端口
	// 80：80端口
	Port *string `json:"Port,omitnil" name:"Port"`

	// 协议；TCP/UDP/ICMP/ANY
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 端口协议类型参数模板id；协议端口模板id；与Protocol,Port互斥
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil" name:"ServiceTemplateId"`

	// 规则的uuid
	RuleUuid *int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`
}

type DescribeEnterpriseSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询时，显示的当前页的页码。
	// 
	// 默认值为1。
	PageNo *string `json:"PageNo,omitnil" name:"PageNo"`

	// 分页查询时，显示的每页数据的最大条数。
	// 
	// 可设置值最大为50。
	PageSize *string `json:"PageSize,omitnil" name:"PageSize"`

	// 访问源示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板(ipm-dyodhpby)
	// instance：资产实例(ins-123456)
	// resourcegroup：资产分组(/全部分组/分组1/子分组1)
	// tag：资源标签({"Key":"标签key值","Value":"标签Value值"})
	// region：地域(ap-gaungzhou)
	// 支持通配
	SourceContent *string `json:"SourceContent,omitnil" name:"SourceContent"`

	// 访问目的示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板(ipm-dyodhpby)
	// instance：资产实例(ins-123456)
	// resourcegroup：资产分组(/全部分组/分组1/子分组1)
	// tag：资源标签({"Key":"标签key值","Value":"标签Value值"})
	// region：地域(ap-gaungzhou)
	// 支持通配
	DestContent *string `json:"DestContent,omitnil" name:"DestContent"`

	// 规则描述，支持通配
	Description *string `json:"Description,omitnil" name:"Description"`

	// 访问控制策略中设置的流量通过云防火墙的方式。取值：
	// accept：放行
	// drop：拒绝
	RuleAction *string `json:"RuleAction,omitnil" name:"RuleAction"`

	// 是否启用规则，默认为启用，取值：
	// true为启用，false为不启用
	Enable *string `json:"Enable,omitnil" name:"Enable"`

	// 访问控制策略的端口。取值：
	// -1/-1：全部端口
	// 80：80端口
	Port *string `json:"Port,omitnil" name:"Port"`

	// 协议；TCP/UDP/ICMP/ANY
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 端口协议类型参数模板id；协议端口模板id；与Protocol,Port互斥
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil" name:"ServiceTemplateId"`

	// 规则的uuid
	RuleUuid *int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`
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
	PageNo *string `json:"PageNo,omitnil" name:"PageNo"`

	// 分页查询时，显示的每页数据的最大条数。
	PageSize *string `json:"PageSize,omitnil" name:"PageSize"`

	// 访问控制策略列表
	Rules []*SecurityGroupRule `json:"Rules,omitnil" name:"Rules"`

	// 访问控制策略的总数量。
	TotalCount *string `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Filters []*CommonFilter `json:"Filters,omitnil" name:"Filters"`

	// 每页条数
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeFwEdgeIpsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil" name:"Filters"`

	// 每页条数
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil" name:"By"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*EdgeIpInfo `json:"Data,omitnil" name:"Data"`

	// ip 开关列表个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 地域列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionLst []*string `json:"RegionLst,omitnil" name:"RegionLst"`

	// 实例类型列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTypeLst []*string `json:"InstanceTypeLst,omitnil" name:"InstanceTypeLst"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeFwGroupInstanceInfoRequest struct {
	*tchttp.BaseRequest
	
	// 每页条数
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil" name:"By"`
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
	// 防火墙(组)
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcFwGroupLst []*VpcFwGroupInfo `json:"VpcFwGroupLst,omitnil" name:"VpcFwGroupLst"`

	// 防火墙(组)个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	SyncType *string `json:"SyncType,omitnil" name:"SyncType"`
}

type DescribeFwSyncStatusRequest struct {
	*tchttp.BaseRequest
	
	// 查询的同步状态类型：Route,同步路由状态
	SyncType *string `json:"SyncType,omitnil" name:"SyncType"`
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
	SyncStatus *int64 `json:"SyncStatus,omitnil" name:"SyncStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Data *ScanInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	IPList []*string `json:"IPList,omitnil" name:"IPList"`
}

type DescribeIPStatusListRequest struct {
	*tchttp.BaseRequest
	
	// 资产Id
	IPList []*string `json:"IPList,omitnil" name:"IPList"`
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
	// ip状态信息
	StatusList []*IPDefendStatus `json:"StatusList,omitnil" name:"StatusList"`

	// 状态码
	ReturnCode *int64 `json:"ReturnCode,omitnil" name:"ReturnCode"`

	// 状态信息
	ReturnMsg *string `json:"ReturnMsg,omitnil" name:"ReturnMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeLogsRequestParams struct {
	// 日志类型标识
	// 流量日志：互联网边界防火墙netflow_border，NAT边界防火墙netflow_nat，VPC间防火墙vpcnetflow，内网流量日志netflow_fl
	// 入侵防御日志rule_threatinfo
	// 访问控制日志：互联网边界规则rule_acl，NAT边界规则rule_acl，内网间规则rule_vpcacl，企业安全组rule_sg
	// 操作日志：防火墙开关-开关操作operate_switch，防火墙开关-实例配置operate_instance，资产中心操作operate_assetgroup，访问控制操作operate_acl，零信任防护操作operate_identity，入侵防御操作-入侵防御operate_ids，入侵防御操作-安全基线operate_baseline，常用工具操作operate_tool，网络蜜罐操作operate_honeypot，日志投递操作operate_logdelivery，通用设置操作operate_logstorage，登录日志operate_login
	Index *string `json:"Index,omitnil" name:"Index"`

	// 每页条数，最大支持2000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值，最大支持60000
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 筛选开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 筛选结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 过滤条件组合，各数组元素间为AND关系，查询字段名Name参考文档https://cloud.tencent.com/document/product/1132/87894，数值类型字段不支持模糊匹配
	Filters []*CommonFilter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeLogsRequest struct {
	*tchttp.BaseRequest
	
	// 日志类型标识
	// 流量日志：互联网边界防火墙netflow_border，NAT边界防火墙netflow_nat，VPC间防火墙vpcnetflow，内网流量日志netflow_fl
	// 入侵防御日志rule_threatinfo
	// 访问控制日志：互联网边界规则rule_acl，NAT边界规则rule_acl，内网间规则rule_vpcacl，企业安全组rule_sg
	// 操作日志：防火墙开关-开关操作operate_switch，防火墙开关-实例配置operate_instance，资产中心操作operate_assetgroup，访问控制操作operate_acl，零信任防护操作operate_identity，入侵防御操作-入侵防御operate_ids，入侵防御操作-安全基线operate_baseline，常用工具操作operate_tool，网络蜜罐操作operate_honeypot，日志投递操作operate_logdelivery，通用设置操作operate_logstorage，登录日志operate_login
	Index *string `json:"Index,omitnil" name:"Index"`

	// 每页条数，最大支持2000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值，最大支持60000
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 筛选开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 筛选结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 过滤条件组合，各数组元素间为AND关系，查询字段名Name参考文档https://cloud.tencent.com/document/product/1132/87894，数值类型字段不支持模糊匹配
	Filters []*CommonFilter `json:"Filters,omitnil" name:"Filters"`
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
	Data *string `json:"Data,omitnil" name:"Data"`

	// 总条数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 返回状态码 0 成功 非0不成功
	ReturnCode *int64 `json:"ReturnCode,omitnil" name:"ReturnCode"`

	// 返回信息  success 成功 其他 不成功
	ReturnMsg *string `json:"ReturnMsg,omitnil" name:"ReturnMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要查询的索引，特定场景使用，可不填
	Index *string `json:"Index,omitnil" name:"Index"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeNatAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 每页条数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要查询的索引，特定场景使用，可不填
	Index *string `json:"Index,omitnil" name:"Index"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil" name:"By"`
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
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// nat访问控制列表数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DescAcItem `json:"Data,omitnil" name:"Data"`

	// 未过滤的总条数
	AllTotal *uint64 `json:"AllTotal,omitnil" name:"AllTotal"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 返回参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnMsg *string `json:"ReturnMsg,omitnil" name:"ReturnMsg"`

	// 当前租户的nat实例个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatFwInsCount *int64 `json:"NatFwInsCount,omitnil" name:"NatFwInsCount"`

	// 当前租户接入子网个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetCount *int64 `json:"SubnetCount,omitnil" name:"SubnetCount"`

	// 打开开关个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenSwitchCount *int64 `json:"OpenSwitchCount,omitnil" name:"OpenSwitchCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	NatinsLst []*NatFwInstance `json:"NatinsLst,omitnil" name:"NatinsLst"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatinsLst []*NatFwInstance `json:"NatinsLst,omitnil" name:"NatinsLst"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Filter []*NatFwFilter `json:"Filter,omitnil" name:"Filter"`

	// 第几页
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 每页长度
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeNatFwInstancesInfoRequest struct {
	*tchttp.BaseRequest
	
	// 获取实例列表过滤字段
	Filter []*NatFwFilter `json:"Filter,omitnil" name:"Filter"`

	// 第几页
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 每页长度
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatinsLst []*NatInstanceInfo `json:"NatinsLst,omitnil" name:"NatinsLst"`

	// nat 防火墙个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeNatFwVpcDnsLstRequestParams struct {
	// natfw 防火墙实例id
	NatFwInsId *string `json:"NatFwInsId,omitnil" name:"NatFwInsId"`

	// natfw 过滤，以','分隔
	NatInsIdFilter *string `json:"NatInsIdFilter,omitnil" name:"NatInsIdFilter"`

	// 分页页数
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 每页最多个数
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeNatFwVpcDnsLstRequest struct {
	*tchttp.BaseRequest
	
	// natfw 防火墙实例id
	NatFwInsId *string `json:"NatFwInsId,omitnil" name:"NatFwInsId"`

	// natfw 过滤，以','分隔
	NatInsIdFilter *string `json:"NatInsIdFilter,omitnil" name:"NatInsIdFilter"`

	// 分页页数
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 每页最多个数
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcDnsSwitchLst []*VpcDnsInfo `json:"VpcDnsSwitchLst,omitnil" name:"VpcDnsSwitchLst"`

	// 返回参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnMsg *string `json:"ReturnMsg,omitnil" name:"ReturnMsg"`

	// 开关总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeNatSwitchListRequestParams struct {
	// 偏移量，分页用
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 条数，分页用
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`

	// 开关，1打开，0关闭
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 筛选NAT防火墙子网开关所属VPC
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 筛选NAT防火墙子网开关所属NAT网关
	NatId *string `json:"NatId,omitnil" name:"NatId"`

	// 筛选NAT防火墙子网开关所属NAT防火墙实例
	NatInsId *string `json:"NatInsId,omitnil" name:"NatInsId"`

	// 筛选NAT防火墙子网开关所属地域
	Area *string `json:"Area,omitnil" name:"Area"`
}

type DescribeNatSwitchListRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，分页用
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 条数，分页用
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`

	// 开关，1打开，0关闭
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 筛选NAT防火墙子网开关所属VPC
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 筛选NAT防火墙子网开关所属NAT网关
	NatId *string `json:"NatId,omitnil" name:"NatId"`

	// 筛选NAT防火墙子网开关所属NAT防火墙实例
	NatInsId *string `json:"NatInsId,omitnil" name:"NatInsId"`

	// 筛选NAT防火墙子网开关所属地域
	Area *string `json:"Area,omitnil" name:"Area"`
}

func (r *DescribeNatSwitchListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatSwitchListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SearchValue")
	delete(f, "Status")
	delete(f, "VpcId")
	delete(f, "NatId")
	delete(f, "NatInsId")
	delete(f, "Area")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNatSwitchListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNatSwitchListResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// NAT边界防火墙开关列表数据
	Data []*NatSwitchListData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNatSwitchListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNatSwitchListResponseParams `json:"Response"`
}

func (r *DescribeNatSwitchListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatSwitchListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceGroupNewRequestParams struct {
	// 查询类型 网络结构-vpc，业务识别-resource ，资源标签-tag
	QueryType *string `json:"QueryType,omitnil" name:"QueryType"`

	// 资产组id  全部传0
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// all  包含子组 own自己
	ShowType *string `json:"ShowType,omitnil" name:"ShowType"`
}

type DescribeResourceGroupNewRequest struct {
	*tchttp.BaseRequest
	
	// 查询类型 网络结构-vpc，业务识别-resource ，资源标签-tag
	QueryType *string `json:"QueryType,omitnil" name:"QueryType"`

	// 资产组id  全部传0
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// all  包含子组 own自己
	ShowType *string `json:"ShowType,omitnil" name:"ShowType"`
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
	Data *string `json:"Data,omitnil" name:"Data"`

	// 未分类实例数量
	UnResourceNum *int64 `json:"UnResourceNum,omitnil" name:"UnResourceNum"`

	// 接口返回消息
	ReturnMsg *string `json:"ReturnMsg,omitnil" name:"ReturnMsg"`

	// 返回码；0为请求成功
	ReturnCode *int64 `json:"ReturnCode,omitnil" name:"ReturnCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	QueryType *string `json:"QueryType,omitnil" name:"QueryType"`

	// 资产组id  全部传0
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`
}

type DescribeResourceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 查询类型 网络结构 vpc，业务识别- resource ，资源标签-tag
	QueryType *string `json:"QueryType,omitnil" name:"QueryType"`

	// 资产组id  全部传0
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceGroupResponseParams struct {
	// 返回树形结构
	Data *string `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`
}

type DescribeRuleOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllTotal *uint64 `json:"AllTotal,omitnil" name:"AllTotal"`

	// 阻断策略规则数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyNum *uint64 `json:"StrategyNum,omitnil" name:"StrategyNum"`

	// 启用规则数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartRuleNum *uint64 `json:"StartRuleNum,omitnil" name:"StartRuleNum"`

	// 停用规则数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	StopRuleNum *uint64 `json:"StopRuleNum,omitnil" name:"StopRuleNum"`

	// 剩余配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemainingNum *int64 `json:"RemainingNum,omitnil" name:"RemainingNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 地域代码（例: ap-guangzhou),支持腾讯云全部地域
	Area *string `json:"Area,omitnil" name:"Area"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`

	// 每页条数，默认为10
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 状态，'': 全部，'0'：筛选停用规则，'1'：筛选启用规则
	Status *string `json:"Status,omitnil" name:"Status"`

	// 0: 不过滤，1：过滤掉正常规则，保留下发异常规则
	Filter *uint64 `json:"Filter,omitnil" name:"Filter"`
}

type DescribeSecurityGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 0: 出站规则，1：入站规则
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 地域代码（例: ap-guangzhou),支持腾讯云全部地域
	Area *string `json:"Area,omitnil" name:"Area"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`

	// 每页条数，默认为10
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 状态，'': 全部，'0'：筛选停用规则，'1'：筛选启用规则
	Status *string `json:"Status,omitnil" name:"Status"`

	// 0: 不过滤，1：过滤掉正常规则，保留下发异常规则
	Filter *uint64 `json:"Filter,omitnil" name:"Filter"`
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
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 安全组规则列表数据
	Data []*SecurityGroupListData `json:"Data,omitnil" name:"Data"`

	// 不算筛选条数的总条数
	AllTotal *uint64 `json:"AllTotal,omitnil" name:"AllTotal"`

	// 访问控制规则全部启用/全部停用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *uint64 `json:"Enable,omitnil" name:"Enable"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 模糊查询
	FuzzySearch *string `json:"FuzzySearch,omitnil" name:"FuzzySearch"`

	// 资产类型 1公网 2内网
	InsType *string `json:"InsType,omitnil" name:"InsType"`

	// ChooseType为1，查询已经分组的资产；ChooseType不为1查询没有分组的资产
	ChooseType *string `json:"ChooseType,omitnil" name:"ChooseType"`

	// 地域
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 查询单页的最大值；eg：10；则最多返回10条结果
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询结果的偏移量
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeSourceAssetRequest struct {
	*tchttp.BaseRequest
	
	// 模糊查询
	FuzzySearch *string `json:"FuzzySearch,omitnil" name:"FuzzySearch"`

	// 资产类型 1公网 2内网
	InsType *string `json:"InsType,omitnil" name:"InsType"`

	// ChooseType为1，查询已经分组的资产；ChooseType不为1查询没有分组的资产
	ChooseType *string `json:"ChooseType,omitnil" name:"ChooseType"`

	// 地域
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 查询单页的最大值；eg：10；则最多返回10条结果
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 查询结果的偏移量
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
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
	delete(f, "FuzzySearch")
	delete(f, "InsType")
	delete(f, "ChooseType")
	delete(f, "Zone")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSourceAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSourceAssetResponseParams struct {
	// 地域集合
	ZoneList []*AssetZone `json:"ZoneList,omitnil" name:"ZoneList"`

	// 数据
	Data []*InstanceInfo `json:"Data,omitnil" name:"Data"`

	// 返回数据总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeSwitchListsRequestParams struct {
	// 防火墙状态  0: 关闭，1：开启
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 资产类型 CVM/NAT/VPN/CLB/其它
	Type *string `json:"Type,omitnil" name:"Type"`

	// 地域 上海/重庆/广州，等等
	Area *string `json:"Area,omitnil" name:"Area"`

	// 搜索值  例子："{"common":"106.54.189.45"}"
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`

	// 条数  默认值:10
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值 默认值: 0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序，desc：降序，asc：升序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段 PortTimes(风险端口数)
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeSwitchListsRequest struct {
	*tchttp.BaseRequest
	
	// 防火墙状态  0: 关闭，1：开启
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 资产类型 CVM/NAT/VPN/CLB/其它
	Type *string `json:"Type,omitnil" name:"Type"`

	// 地域 上海/重庆/广州，等等
	Area *string `json:"Area,omitnil" name:"Area"`

	// 搜索值  例子："{"common":"106.54.189.45"}"
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`

	// 条数  默认值:10
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值 默认值: 0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 排序，desc：降序，asc：升序
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序字段 PortTimes(风险端口数)
	By *string `json:"By,omitnil" name:"By"`
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
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 列表数据
	Data []*SwitchListsData `json:"Data,omitnil" name:"Data"`

	// 区域列表
	AreaLists []*string `json:"AreaLists,omitnil" name:"AreaLists"`

	// 打开个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OnNum *uint64 `json:"OnNum,omitnil" name:"OnNum"`

	// 关闭个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OffNum *uint64 `json:"OffNum,omitnil" name:"OffNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 类型 1 告警 2阻断
	QueryType *string `json:"QueryType,omitnil" name:"QueryType"`

	// 查询条件
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`
}

type DescribeTLogInfoRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 类型 1 告警 2阻断
	QueryType *string `json:"QueryType,omitnil" name:"QueryType"`

	// 查询条件
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`
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
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QueryType")
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
	Data *TLogInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 类型 1 告警 2阻断
	QueryType *string `json:"QueryType,omitnil" name:"QueryType"`

	// top数
	Top *int64 `json:"Top,omitnil" name:"Top"`

	// 查询条件
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`
}

type DescribeTLogIpListRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 类型 1 告警 2阻断
	QueryType *string `json:"QueryType,omitnil" name:"QueryType"`

	// top数
	Top *int64 `json:"Top,omitnil" name:"Top"`

	// 查询条件
	SearchValue *string `json:"SearchValue,omitnil" name:"SearchValue"`
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
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QueryType")
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
	Data []*StaticInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// 状态值，0：检查表的状态 确实只有一个默认值
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// Nat所在地域 NAT填Area，不要填Edgeid；
	Area *string `json:"Area,omitnil" name:"Area"`

	// 方向，0：出站，1：入站 默认值为 0
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`
}

type DescribeTableStatusRequest struct {
	*tchttp.BaseRequest
	
	// EdgeId值两个vpc间的边id vpc填Edgeid，不要填Area；
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// 状态值，0：检查表的状态 确实只有一个默认值
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// Nat所在地域 NAT填Area，不要填Edgeid；
	Area *string `json:"Area,omitnil" name:"Area"`

	// 方向，0：出站，1：入站 默认值为 0
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询示例ID
	AssetID *string `json:"AssetID,omitnil" name:"AssetID"`
}

type DescribeUnHandleEventTabListRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询示例ID
	AssetID *string `json:"AssetID,omitnil" name:"AssetID"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *UnHandleEvent `json:"Data,omitnil" name:"Data"`

	// 错误码，0成功 非0错误
	ReturnCode *int64 `json:"ReturnCode,omitnil" name:"ReturnCode"`

	// 返回信息 success成功
	ReturnMsg *string `json:"ReturnMsg,omitnil" name:"ReturnMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要查询的索引，特定场景使用，可不填
	Index *string `json:"Index,omitnil" name:"Index"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeVpcAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 每页条数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 需要查询的索引，特定场景使用，可不填
	Index *string `json:"Index,omitnil" name:"Index"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil" name:"By"`
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
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 内网间访问控制列表数据
	Data []*VpcRuleItem `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeVpcFwGroupSwitchRequestParams struct {
	// 每页条数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil" name:"By"`
}

type DescribeVpcFwGroupSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 每页条数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件组合
	Filters []*CommonFilter `json:"Filters,omitnil" name:"Filters"`

	// 检索的起始时间，可不传
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 检索的截止时间，可不传
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// desc：降序；asc：升序。根据By字段的值进行排序，这里传参的话则By也必须有值
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序所用到的字段
	By *string `json:"By,omitnil" name:"By"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	SwitchList []*FwGroupSwitchShow `json:"SwitchList,omitnil" name:"SwitchList"`

	// 开关总个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 0：设置为关闭 1:设置为打开
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type EdgeIpInfo struct {
	// 公网IP
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 公网 IP 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpType *int64 `json:"PublicIpType,omitnil" name:"PublicIpType"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntranetIp *string `json:"IntranetIp,omitnil" name:"IntranetIp"`

	// 资产类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetType *string `json:"AssetType,omitnil" name:"AssetType"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 风险端口数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortRiskCount *int64 `json:"PortRiskCount,omitnil" name:"PortRiskCount"`

	// 最近扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastScanTime *string `json:"LastScanTime,omitnil" name:"LastScanTime"`

	// 是否为region eip
	// 0 不为region eip，不能选择串行
	// 1 为region eip 可以选择串行
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsRegionEip *int64 `json:"IsRegionEip,omitnil" name:"IsRegionEip"`

	// EIP 所关联的VPC
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 0: 该地域暂未支持串行
	// 1: 该用户未在该地域配置串行带宽
	// 2: 该用户已在该地域配置串行带宽，可以开启串行开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSerialRegion *int64 `json:"IsSerialRegion,omitnil" name:"IsSerialRegion"`

	// 0: 不是公网CLB 可以开启串行开关
	// 1: 是公网CLB 不可以开启串行开关
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPublicClb *int64 `json:"IsPublicClb,omitnil" name:"IsPublicClb"`

	// 0: 开启开关时提示要创建私有连接。
	// 1: 关闭该开关是提示删除私有连接。
	// 如果大于 1: 关闭开关 、开启开关不需提示创建删除私有连接。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndpointBindEipNum *int64 `json:"EndpointBindEipNum,omitnil" name:"EndpointBindEipNum"`

	// 扫描深度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanMode *string `json:"ScanMode,omitnil" name:"ScanMode"`

	// 扫描状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanStatus *int64 `json:"ScanStatus,omitnil" name:"ScanStatus"`

	// 开关状态
	// 0 : 关闭
	// 1 : 开启
	// 2 : 开启中
	// 3 : 关闭中
	// 4 : 异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 私有连接ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndpointId *string `json:"EndpointId,omitnil" name:"EndpointId"`

	// 私有连接IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndpointIp *string `json:"EndpointIp,omitnil" name:"EndpointIp"`

	// 0 : 旁路
	// 1 : 串行
	// 2 : 正在模式切换
	SwitchMode *uint64 `json:"SwitchMode,omitnil" name:"SwitchMode"`

	// 开关权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	SwitchWeight *int64 `json:"SwitchWeight,omitnil" name:"SwitchWeight"`
}

type EdgeIpSwitch struct {
	// 公网IP
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// vpc 中第一个EIP开关打开，需要指定子网创建私有连接
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 创建私有连接指定IP
	EndpointIp *string `json:"EndpointIp,omitnil" name:"EndpointIp"`

	// 0 : 旁路 1 : 串行
	SwitchMode *int64 `json:"SwitchMode,omitnil" name:"SwitchMode"`
}

// Predefined struct for user
type ExpandCfwVerticalRequestParams struct {
	// nat：nat防火墙，ew：东西向防火墙
	FwType *string `json:"FwType,omitnil" name:"FwType"`

	// 带宽值
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitnil" name:"CfwInstance"`
}

type ExpandCfwVerticalRequest struct {
	*tchttp.BaseRequest
	
	// nat：nat防火墙，ew：东西向防火墙
	FwType *string `json:"FwType,omitnil" name:"FwType"`

	// 带宽值
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitnil" name:"CfwInstance"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExpandCfwVerticalRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExpandCfwVerticalResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	FwCidrType *string `json:"FwCidrType,omitnil" name:"FwCidrType"`

	// 为每个vpc指定防火墙的网段
	FwCidrLst []*FwVpcCidr `json:"FwCidrLst,omitnil" name:"FwCidrLst"`

	// 其他防火墙占用网段，一般是防火墙需要独占vpc时指定的网段
	ComFwCidr *string `json:"ComFwCidr,omitnil" name:"ComFwCidr"`
}

type FwDeploy struct {
	// 防火墙部署地域
	DeployRegion *string `json:"DeployRegion,omitnil" name:"DeployRegion"`

	// 带宽，单位：Mbps
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 异地灾备 1：使用异地灾备；0：不使用异地灾备；为空则默认不使用异地灾备
	CrossAZone *int64 `json:"CrossAZone,omitnil" name:"CrossAZone"`

	// 主可用区，为空则选择默认可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 备可用区，为空则选择默认可用区
	ZoneBak *string `json:"ZoneBak,omitnil" name:"ZoneBak"`

	// 若为cdc防火墙时填充该id
	CdcId *string `json:"CdcId,omitnil" name:"CdcId"`
}

type FwGateway struct {
	// 防火墙网关id
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 网关所属vpc id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 网关ip地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpAddress *string `json:"IpAddress,omitnil" name:"IpAddress"`
}

type FwGroupSwitch struct {
	// 防火墙实例的开关模式 1: 单点互通 2: 多点互通 3: 全互通 4: 自定义路由
	SwitchMode *int64 `json:"SwitchMode,omitnil" name:"SwitchMode"`

	// 防火墙开关ID
	// 支持三种类型
	// 1. 边开关(单点互通)
	// 2. 点开关(多点互通)
	// 3. 全开关(全互通)
	SwitchId *string `json:"SwitchId,omitnil" name:"SwitchId"`
}

type FwGroupSwitchShow struct {
	// 防火墙开关ID
	SwitchId *string `json:"SwitchId,omitnil" name:"SwitchId"`

	// 防火墙开关NAME
	// 注意：此字段可能返回 null，表示取不到有效值。
	SwitchName *string `json:"SwitchName,omitnil" name:"SwitchName"`

	// 互通模式
	SwitchMode *int64 `json:"SwitchMode,omitnil" name:"SwitchMode"`

	// 开关边连接类型 0：对等连接， 1：云连网
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectType *int64 `json:"ConnectType,omitnil" name:"ConnectType"`

	// 连接ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectId *string `json:"ConnectId,omitnil" name:"ConnectId"`

	// 连接名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectName *string `json:"ConnectName,omitnil" name:"ConnectName"`

	// 源实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcInstancesInfo []*NetInstancesInfo `json:"SrcInstancesInfo,omitnil" name:"SrcInstancesInfo"`

	// 目的实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstInstancesInfo []*NetInstancesInfo `json:"DstInstancesInfo,omitnil" name:"DstInstancesInfo"`

	// 防火墙(组)数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwGroupId *string `json:"FwGroupId,omitnil" name:"FwGroupId"`

	// 防火墙(组)名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwGroupName *string `json:"FwGroupName,omitnil" name:"FwGroupName"`

	// 开关状态 0：关 ， 1：开
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *int64 `json:"Enable,omitnil" name:"Enable"`

	// 开关的状态 0：正常， 1：转换中
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 0-非sase实例，忽略，1-未绑定状态，2-已绑定
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttachWithEdge *int64 `json:"AttachWithEdge,omitnil" name:"AttachWithEdge"`

	// 对等防火墙和开关状态 0：正常， 1：对等未创建防火墙，2：对等已创建防火墙，未打开开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrossEdgeStatus *int64 `json:"CrossEdgeStatus,omitnil" name:"CrossEdgeStatus"`

	// 网络经过VPC防火墙CVM所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwInsRegion []*string `json:"FwInsRegion,omitnil" name:"FwInsRegion"`

	// 0 观察 1 拦截 2 严格 3 关闭
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpsAction *int64 `json:"IpsAction,omitnil" name:"IpsAction"`

	// 开关关联的防火墙实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwInsLst []*VpcFwInstanceShow `json:"FwInsLst,omitnil" name:"FwInsLst"`

	// 开关是否处于bypass状态
	// 0：正常状态
	// 1：bypass状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	BypassStatus *int64 `json:"BypassStatus,omitnil" name:"BypassStatus"`
}

type FwVpcCidr struct {
	// vpc的id
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 防火墙网段，最少/24的网段
	FwCidr *string `json:"FwCidr,omitnil" name:"FwCidr"`
}

type IPDefendStatus struct {
	// ip地址
	IP *string `json:"IP,omitnil" name:"IP"`

	// 防护状态   1:防护打开; -1:地址错误; 其他:未防护
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type InstanceInfo struct {
	// appid信息
	AppId *string `json:"AppId,omitnil" name:"AppId"`

	// 地域
	Region *string `json:"Region,omitnil" name:"Region"`

	// vpcid信息
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// vpc名称
	VPCName *string `json:"VPCName,omitnil" name:"VPCName"`

	// 子网id
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 资产id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 资产名
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 资产类型
	//  3是cvm实例,4是clb实例,5是eni实例,6是mysql,7是redis,8是NAT,9是VPN,10是ES,11是MARIADB,12是KAFKA 13 NATFW
	InsType *int64 `json:"InsType,omitnil" name:"InsType"`

	// 公网ip
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 内网ip
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 端口数
	PortNum *string `json:"PortNum,omitnil" name:"PortNum"`

	// 漏洞数
	LeakNum *string `json:"LeakNum,omitnil" name:"LeakNum"`

	// 1，公网 2内网
	InsSource *string `json:"InsSource,omitnil" name:"InsSource"`

	// [a,b]
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourcePath []*string `json:"ResourcePath,omitnil" name:"ResourcePath"`

	// 扫描结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Server []*string `json:"Server,omitnil" name:"Server"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionKey *string `json:"RegionKey,omitnil" name:"RegionKey"`
}

type IntrusionDefenseRule struct {
	// 规则方向，0出站，1入站，3内网间
	Direction *int64 `json:"Direction,omitnil" name:"Direction"`

	// 规则结束时间，格式：2006-01-02 15:04:05，必须大于当前时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 规则IP地址，IP与Domain必填其中之一
	IP *string `json:"IP,omitnil" name:"IP"`

	// 规则域名，IP与Domain必填其中之一
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 备注信息，长度不能超过50
	Comment *string `json:"Comment,omitnil" name:"Comment"`

	// 规则开始时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`
}

type IocListData struct {
	// 待处置IP地址，IP/Domain字段二选一
	// 注意：此字段可能返回 null，表示取不到有效值。
	IP *string `json:"IP,omitnil" name:"IP"`

	// 只能为0或者1   0代表出站 1代表入站
	// 注意：此字段可能返回 null，表示取不到有效值。
	Direction *int64 `json:"Direction,omitnil" name:"Direction"`

	// 待处置域名，IP/Domain字段二选一
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type IpStatic struct {
	// 值
	Num *int64 `json:"Num,omitnil" name:"Num"`

	// 折线图横坐标时间
	StatTime *string `json:"StatTime,omitnil" name:"StatTime"`
}

// Predefined struct for user
type ModifyAcRuleRequestParams struct {
	// 规则数组
	Data []*RuleInfoData `json:"Data,omitnil" name:"Data"`

	// EdgeId值
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// 访问规则状态
	Enable *int64 `json:"Enable,omitnil" name:"Enable"`

	// NAT地域
	Area *string `json:"Area,omitnil" name:"Area"`
}

type ModifyAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则数组
	Data []*RuleInfoData `json:"Data,omitnil" name:"Data"`

	// EdgeId值
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// 访问规则状态
	Enable *int64 `json:"Enable,omitnil" name:"Enable"`

	// NAT地域
	Area *string `json:"Area,omitnil" name:"Area"`
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
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 返回多余的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Info *string `json:"Info,omitnil" name:"Info"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 需要编辑的规则数组
	Rules []*CreateRuleItem `json:"Rules,omitnil" name:"Rules"`
}

type ModifyAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要编辑的规则数组
	Rules []*CreateRuleItem `json:"Rules,omitnil" name:"Rules"`
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
	RuleUuid []*int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 模板名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板描述
	Detail *string `json:"Detail,omitnil" name:"Detail"`

	// Type为1，ip模板eg：1.1.1.1,2.2.2.2；
	// Type为5，域名模板eg：www.qq.com,www.tencent.com
	IpString *string `json:"IpString,omitnil" name:"IpString"`

	// 1 ip模板
	// 5 域名模板
	Type *int64 `json:"Type,omitnil" name:"Type"`
}

type ModifyAddressTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 地址模板唯一Id
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 模板名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 模板描述
	Detail *string `json:"Detail,omitnil" name:"Detail"`

	// Type为1，ip模板eg：1.1.1.1,2.2.2.2；
	// Type为5，域名模板eg：www.qq.com,www.tencent.com
	IpString *string `json:"IpString,omitnil" name:"IpString"`

	// 1 ip模板
	// 5 域名模板
	Type *int64 `json:"Type,omitnil" name:"Type"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAddressTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAddressTemplateResponseParams struct {
	// 创建结果,0成功
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 唯一Id
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 选中的防火墙开关Id
	FireWallPublicIPs []*string `json:"FireWallPublicIPs,omitnil" name:"FireWallPublicIPs"`
}

type ModifyAllPublicIPSwitchStatusRequest struct {
	*tchttp.BaseRequest
	
	// 状态，0：关闭，1：开启
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 选中的防火墙开关Id
	FireWallPublicIPs []*string `json:"FireWallPublicIPs,omitnil" name:"FireWallPublicIPs"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnMsg *string `json:"ReturnMsg,omitnil" name:"ReturnMsg"`

	// 接口返回错误码，0请求成功  非0失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnCode *int64 `json:"ReturnCode,omitnil" name:"ReturnCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// Edge ID值
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// NAT地域
	Area *string `json:"Area,omitnil" name:"Area"`
}

type ModifyAllRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 状态，0：全部停用，1：全部启用
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// Edge ID值
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// NAT地域
	Area *string `json:"Area,omitnil" name:"Area"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type ModifyAllVPCSwitchStatusRequestParams struct {
	// 状态，0：关闭，1：开启
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 选中的防火墙开关Id
	FireWallVpcIds []*string `json:"FireWallVpcIds,omitnil" name:"FireWallVpcIds"`
}

type ModifyAllVPCSwitchStatusRequest struct {
	*tchttp.BaseRequest
	
	// 状态，0：关闭，1：开启
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 选中的防火墙开关Id
	FireWallVpcIds []*string `json:"FireWallVpcIds,omitnil" name:"FireWallVpcIds"`
}

func (r *ModifyAllVPCSwitchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllVPCSwitchStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "FireWallVpcIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAllVPCSwitchStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAllVPCSwitchStatusResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyAllVPCSwitchStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAllVPCSwitchStatusResponseParams `json:"Response"`
}

func (r *ModifyAllVPCSwitchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllVPCSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAssetScanRequestParams struct {
	// 扫描范围：1端口, 2端口+漏扫
	ScanRange *int64 `json:"ScanRange,omitnil" name:"ScanRange"`

	// 扫描深度：'heavy', 'medium', 'light'
	ScanDeep *string `json:"ScanDeep,omitnil" name:"ScanDeep"`

	// 扫描类型：1立即扫描 2 周期任务
	RangeType *int64 `json:"RangeType,omitnil" name:"RangeType"`

	// RangeType为2 是必须添加，定时任务时间
	ScanPeriod *string `json:"ScanPeriod,omitnil" name:"ScanPeriod"`

	// 立即扫描这个字段传过滤的扫描集合
	ScanFilterIp []*string `json:"ScanFilterIp,omitnil" name:"ScanFilterIp"`

	// 1全量2单个
	ScanType *int64 `json:"ScanType,omitnil" name:"ScanType"`
}

type ModifyAssetScanRequest struct {
	*tchttp.BaseRequest
	
	// 扫描范围：1端口, 2端口+漏扫
	ScanRange *int64 `json:"ScanRange,omitnil" name:"ScanRange"`

	// 扫描深度：'heavy', 'medium', 'light'
	ScanDeep *string `json:"ScanDeep,omitnil" name:"ScanDeep"`

	// 扫描类型：1立即扫描 2 周期任务
	RangeType *int64 `json:"RangeType,omitnil" name:"RangeType"`

	// RangeType为2 是必须添加，定时任务时间
	ScanPeriod *string `json:"ScanPeriod,omitnil" name:"ScanPeriod"`

	// 立即扫描这个字段传过滤的扫描集合
	ScanFilterIp []*string `json:"ScanFilterIp,omitnil" name:"ScanFilterIp"`

	// 1全量2单个
	ScanType *int64 `json:"ScanType,omitnil" name:"ScanType"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnMsg *string `json:"ReturnMsg,omitnil" name:"ReturnMsg"`

	// 接口返回错误码，0请求成功  非0失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnCode *int64 `json:"ReturnCode,omitnil" name:"ReturnCode"`

	// 状态值 0：成功，1 执行扫描中,其他：失败
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// success 成功
	// 其他失败
	ReturnMsg *string `json:"ReturnMsg,omitnil" name:"ReturnMsg"`

	// 0 成功
	// 非0 失败
	ReturnCode *uint64 `json:"ReturnCode,omitnil" name:"ReturnCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RuleType *int64 `json:"RuleType,omitnil" name:"RuleType"`

	// IP、Domain二选一（注：封禁列表，只能填写IP），不能同时为空
	IOC []*IocListData `json:"IOC,omitnil" name:"IOC"`

	// 可选值：delete（删除）、edit（编辑）、add（添加）  其他值无效
	IocAction *string `json:"IocAction,omitnil" name:"IocAction"`

	// 时间格式：yyyy-MM-dd HH:mm:ss，IocAction 为edit或add时必填
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 时间格式：yyyy-MM-dd HH:mm:ss，IocAction 为edit或add时必填，必须大于当前时间且大于StartTime
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type ModifyBlockIgnoreListRequest struct {
	*tchttp.BaseRequest
	
	// 1封禁列表 2 放通列表
	RuleType *int64 `json:"RuleType,omitnil" name:"RuleType"`

	// IP、Domain二选一（注：封禁列表，只能填写IP），不能同时为空
	IOC []*IocListData `json:"IOC,omitnil" name:"IOC"`

	// 可选值：delete（删除）、edit（编辑）、add（添加）  其他值无效
	IocAction *string `json:"IocAction,omitnil" name:"IocAction"`

	// 时间格式：yyyy-MM-dd HH:mm:ss，IocAction 为edit或add时必填
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 时间格式：yyyy-MM-dd HH:mm:ss，IocAction 为edit或add时必填，必须大于当前时间且大于StartTime
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
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
	ReturnMsg *string `json:"ReturnMsg,omitnil" name:"ReturnMsg"`

	// 接口返回错误码，0请求成功  非0失败
	ReturnCode *uint64 `json:"ReturnCode,omitnil" name:"ReturnCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type ModifyBlockIgnoreRuleRequestParams struct {
	// 规则
	Rule *IntrusionDefenseRule `json:"Rule,omitnil" name:"Rule"`

	// 规则类型，1封禁，2放通
	RuleType *int64 `json:"RuleType,omitnil" name:"RuleType"`
}

type ModifyBlockIgnoreRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则
	Rule *IntrusionDefenseRule `json:"Rule,omitnil" name:"Rule"`

	// 规则类型，1封禁，2放通
	RuleType *int64 `json:"RuleType,omitnil" name:"RuleType"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 记录id
	UniqueId *string `json:"UniqueId,omitnil" name:"UniqueId"`

	// 操作类型 1 置顶 0取消
	OpeType *string `json:"OpeType,omitnil" name:"OpeType"`
}

type ModifyBlockTopRequest struct {
	*tchttp.BaseRequest
	
	// 记录id
	UniqueId *string `json:"UniqueId,omitnil" name:"UniqueId"`

	// 操作类型 1 置顶 0取消
	OpeType *string `json:"OpeType,omitnil" name:"OpeType"`
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
	delete(f, "UniqueId")
	delete(f, "OpeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBlockTopRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBlockTopResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type ModifyEWRuleStatusRequestParams struct {
	// vpc规则必填，边id
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// 是否开关开启，0：未开启，1：开启
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 规则方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 更改的规则当前执行顺序
	RuleSequence *uint64 `json:"RuleSequence,omitnil" name:"RuleSequence"`

	// 规则类型，vpc：VPC间规则、nat：Nat边界规则
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`
}

type ModifyEWRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// vpc规则必填，边id
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// 是否开关开启，0：未开启，1：开启
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 规则方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 更改的规则当前执行顺序
	RuleSequence *uint64 `json:"RuleSequence,omitnil" name:"RuleSequence"`

	// 规则类型，vpc：VPC间规则、nat：Nat边界规则
	RuleType *string `json:"RuleType,omitnil" name:"RuleType"`
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
	ReturnCode *int64 `json:"ReturnCode,omitnil" name:"ReturnCode"`

	// 状态信息，success：查询成功，fail：查询失败
	ReturnMsg *string `json:"ReturnMsg,omitnil" name:"ReturnMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Enable *int64 `json:"Enable,omitnil" name:"Enable"`

	// 操作开关详情
	EdgeIpSwitchLst []*EdgeIpSwitch `json:"EdgeIpSwitchLst,omitnil" name:"EdgeIpSwitchLst"`

	// 0 不自动选择子网
	// 1 自动选择子网创建私有连接
	AutoChooseSubnet *int64 `json:"AutoChooseSubnet,omitnil" name:"AutoChooseSubnet"`

	// 0 切换为旁路
	// 1 切换为串行
	// 2 不切换模式，此次操作开关
	SwitchMode *int64 `json:"SwitchMode,omitnil" name:"SwitchMode"`
}

type ModifyEdgeIpSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 0 关闭开关
	// 1 打开开关
	// 2 不操作开关，此次切换模式
	Enable *int64 `json:"Enable,omitnil" name:"Enable"`

	// 操作开关详情
	EdgeIpSwitchLst []*EdgeIpSwitch `json:"EdgeIpSwitchLst,omitnil" name:"EdgeIpSwitchLst"`

	// 0 不自动选择子网
	// 1 自动选择子网创建私有连接
	AutoChooseSubnet *int64 `json:"AutoChooseSubnet,omitnil" name:"AutoChooseSubnet"`

	// 0 切换为旁路
	// 1 切换为串行
	// 2 不切换模式，此次操作开关
	SwitchMode *int64 `json:"SwitchMode,omitnil" name:"SwitchMode"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Status *uint64 `json:"Status,omitnil" name:"Status"`
}

type ModifyEnterpriseSecurityDispatchStatusRequest struct {
	*tchttp.BaseRequest
	
	// 0：打开立即下发开关；
	// 
	// 1：关闭立即下发开关；
	// 
	// 2：关闭立即下发开关情况下，触发开始下发
	Status *uint64 `json:"Status,omitnil" name:"Status"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RuleUuid *uint64 `json:"RuleUuid,omitnil" name:"RuleUuid"`

	// 修改类型，0：修改规则内容；1：修改单条规则开关状态；2：修改所有规则开关状态
	ModifyType *uint64 `json:"ModifyType,omitnil" name:"ModifyType"`

	// 编辑后的企业安全组规则数据；修改规则状态不用填该字段
	Data *SecurityGroupRule `json:"Data,omitnil" name:"Data"`

	// 0是关闭,1是开启
	Enable *uint64 `json:"Enable,omitnil" name:"Enable"`
}

type ModifyEnterpriseSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则的uuid，可通过查询规则列表获取
	RuleUuid *uint64 `json:"RuleUuid,omitnil" name:"RuleUuid"`

	// 修改类型，0：修改规则内容；1：修改单条规则开关状态；2：修改所有规则开关状态
	ModifyType *uint64 `json:"ModifyType,omitnil" name:"ModifyType"`

	// 编辑后的企业安全组规则数据；修改规则状态不用填该字段
	Data *SecurityGroupRule `json:"Data,omitnil" name:"Data"`

	// 0是关闭,1是开启
	Enable *uint64 `json:"Enable,omitnil" name:"Enable"`
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
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 编辑后新生成规则的Id
	NewRuleUuid *uint64 `json:"NewRuleUuid,omitnil" name:"NewRuleUuid"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Enable *int64 `json:"Enable,omitnil" name:"Enable"`

	// 是否操作全部开关 0 不操作全部开关，1 操作全部开关
	AllSwitch *int64 `json:"AllSwitch,omitnil" name:"AllSwitch"`

	// 开关列表
	SwitchList []*FwGroupSwitch `json:"SwitchList,omitnil" name:"SwitchList"`
}

type ModifyFwGroupSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 打开或关闭开关
	// 0：关闭开关
	// 1：打开开关
	Enable *int64 `json:"Enable,omitnil" name:"Enable"`

	// 是否操作全部开关 0 不操作全部开关，1 操作全部开关
	AllSwitch *int64 `json:"AllSwitch,omitnil" name:"AllSwitch"`

	// 开关列表
	SwitchList []*FwGroupSwitch `json:"SwitchList,omitnil" name:"SwitchList"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type ModifyNatAcRuleRequestParams struct {
	// 需要编辑的规则数组
	Rules []*CreateNatRuleItem `json:"Rules,omitnil" name:"Rules"`
}

type ModifyNatAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要编辑的规则数组
	Rules []*CreateNatRuleItem `json:"Rules,omitnil" name:"Rules"`
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
	RuleUuid []*int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitnil" name:"CfwInstance"`

	// 接入模式重新接入的nat网关列表，其中NatGwList和VpcList只能传递一个。
	NatGwList []*string `json:"NatGwList,omitnil" name:"NatGwList"`

	// 新增模式重新接入的vpc列表，其中NatGwList和NatgwList只能传递一个。
	VpcList []*string `json:"VpcList,omitnil" name:"VpcList"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil" name:"FwCidrInfo"`
}

type ModifyNatFwReSelectRequest struct {
	*tchttp.BaseRequest
	
	// 模式 1：接入模式；0：新增模式
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitnil" name:"CfwInstance"`

	// 接入模式重新接入的nat网关列表，其中NatGwList和VpcList只能传递一个。
	NatGwList []*string `json:"NatGwList,omitnil" name:"NatGwList"`

	// 新增模式重新接入的vpc列表，其中NatGwList和NatgwList只能传递一个。
	VpcList []*string `json:"VpcList,omitnil" name:"VpcList"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil" name:"FwCidrInfo"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Enable *int64 `json:"Enable,omitnil" name:"Enable"`

	// 防火墙实例id列表，其中CfwInsIdList，SubnetIdList和RouteTableIdList只能传递一种。
	CfwInsIdList []*string `json:"CfwInsIdList,omitnil" name:"CfwInsIdList"`

	// 子网id列表，其中CfwInsIdList，SubnetIdList和RouteTableIdList只能传递一种。
	SubnetIdList []*string `json:"SubnetIdList,omitnil" name:"SubnetIdList"`

	// 路由表id列表，其中CfwInsIdList，SubnetIdList和RouteTableIdList只能传递一种。
	RouteTableIdList []*string `json:"RouteTableIdList,omitnil" name:"RouteTableIdList"`
}

type ModifyNatFwSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 开关，0：关闭，1：开启
	Enable *int64 `json:"Enable,omitnil" name:"Enable"`

	// 防火墙实例id列表，其中CfwInsIdList，SubnetIdList和RouteTableIdList只能传递一种。
	CfwInsIdList []*string `json:"CfwInsIdList,omitnil" name:"CfwInsIdList"`

	// 子网id列表，其中CfwInsIdList，SubnetIdList和RouteTableIdList只能传递一种。
	SubnetIdList []*string `json:"SubnetIdList,omitnil" name:"SubnetIdList"`

	// 路由表id列表，其中CfwInsIdList，SubnetIdList和RouteTableIdList只能传递一种。
	RouteTableIdList []*string `json:"RouteTableIdList,omitnil" name:"RouteTableIdList"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	NatFwInsId *string `json:"NatFwInsId,omitnil" name:"NatFwInsId"`

	// DNS 开关切换列表
	DnsVpcSwitchLst []*DnsVpcSwitch `json:"DnsVpcSwitchLst,omitnil" name:"DnsVpcSwitchLst"`
}

type ModifyNatFwVpcDnsSwitchRequest struct {
	*tchttp.BaseRequest
	
	// nat 防火墙 id
	NatFwInsId *string `json:"NatFwInsId,omitnil" name:"NatFwInsId"`

	// DNS 开关切换列表
	DnsVpcSwitchLst []*DnsVpcSwitch `json:"DnsVpcSwitchLst,omitnil" name:"DnsVpcSwitchLst"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnMsg *string `json:"ReturnMsg,omitnil" name:"ReturnMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// NAT防火墙实例ID
	NatInstanceId *string `json:"NatInstanceId,omitnil" name:"NatInstanceId"`
}

type ModifyNatInstanceRequest struct {
	*tchttp.BaseRequest
	
	// NAT防火墙实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// NAT防火墙实例ID
	NatInstanceId *string `json:"NatInstanceId,omitnil" name:"NatInstanceId"`
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
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// nat实例唯一ID
	NatInstanceId *string `json:"NatInstanceId,omitnil" name:"NatInstanceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RuleChangeItems []*RuleChangeItem `json:"RuleChangeItems,omitnil" name:"RuleChangeItems"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`
}

type ModifyNatSequenceRulesRequest struct {
	*tchttp.BaseRequest
	
	// 规则快速排序：OrderIndex，原始序号；NewOrderIndex：新序号
	RuleChangeItems []*RuleChangeItem `json:"RuleChangeItems,omitnil" name:"RuleChangeItems"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type ModifyPublicIPSwitchStatusRequestParams struct {
	// 公网IP
	FireWallPublicIP *string `json:"FireWallPublicIP,omitnil" name:"FireWallPublicIP"`

	// 状态值，0: 关闭 ,1:开启
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type ModifyPublicIPSwitchStatusRequest struct {
	*tchttp.BaseRequest
	
	// 公网IP
	FireWallPublicIP *string `json:"FireWallPublicIP,omitnil" name:"FireWallPublicIP"`

	// 状态值，0: 关闭 ,1:开启
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

func (r *ModifyPublicIPSwitchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPublicIPSwitchStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FireWallPublicIP")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPublicIPSwitchStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPublicIPSwitchStatusResponseParams struct {
	// 接口返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnMsg *string `json:"ReturnMsg,omitnil" name:"ReturnMsg"`

	// 接口返回错误码，0请求成功  非0失败
	ReturnCode *int64 `json:"ReturnCode,omitnil" name:"ReturnCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyPublicIPSwitchStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPublicIPSwitchStatusResponseParams `json:"Response"`
}

func (r *ModifyPublicIPSwitchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPublicIPSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourceGroupRequestParams struct {
	// 组id
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 组名称
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 上级组id
	ParentId *string `json:"ParentId,omitnil" name:"ParentId"`
}

type ModifyResourceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 组id
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 组名称
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 上级组id
	ParentId *string `json:"ParentId,omitnil" name:"ParentId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Type *uint64 `json:"Type,omitnil" name:"Type"`
}

type ModifyRunSyncAssetRequest struct {
	*tchttp.BaseRequest
	
	// 0: 互联网防火墙开关，1：vpc 防火墙开关
	Type *uint64 `json:"Type,omitnil" name:"Type"`
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
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 是否开关开启，0：未开启，1：开启
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 更改的企业安全组规则的执行顺序
	RuleSequence *uint64 `json:"RuleSequence,omitnil" name:"RuleSequence"`
}

type ModifySecurityGroupItemRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 是否开关开启，0：未开启，1：开启
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 更改的企业安全组规则的执行顺序
	RuleSequence *uint64 `json:"RuleSequence,omitnil" name:"RuleSequence"`
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
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 编辑后是否启用规则，0：不启用，1：启用，默认1
	Enable *uint64 `json:"Enable,omitnil" name:"Enable"`

	// 编辑的企业安全组规则数据
	Data []*SecurityGroupListData `json:"Data,omitnil" name:"Data"`

	// 编辑的企业安全组规则的原始执行顺序
	SgRuleOriginSequence *uint64 `json:"SgRuleOriginSequence,omitnil" name:"SgRuleOriginSequence"`
}

type ModifySecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 编辑后是否启用规则，0：不启用，1：启用，默认1
	Enable *uint64 `json:"Enable,omitnil" name:"Enable"`

	// 编辑的企业安全组规则数据
	Data []*SecurityGroupListData `json:"Data,omitnil" name:"Data"`

	// 编辑的企业安全组规则的原始执行顺序
	SgRuleOriginSequence *uint64 `json:"SgRuleOriginSequence,omitnil" name:"SgRuleOriginSequence"`
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
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 编辑后新生成规则的Id
	NewRuleId *uint64 `json:"NewRuleId,omitnil" name:"NewRuleId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 企业安全组规则快速排序数据
	Data []*SecurityGroupOrderIndexData `json:"Data,omitnil" name:"Data"`
}

type ModifySecurityGroupSequenceRulesRequest struct {
	*tchttp.BaseRequest
	
	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 企业安全组规则快速排序数据
	Data []*SecurityGroupOrderIndexData `json:"Data,omitnil" name:"Data"`
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
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RuleChangeItems []*RuleChangeItem `json:"RuleChangeItems,omitnil" name:"RuleChangeItems"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`
}

type ModifySequenceAclRulesRequest struct {
	*tchttp.BaseRequest
	
	// 规则快速排序：OrderIndex，原始序号；NewOrderIndex：新序号
	RuleChangeItems []*RuleChangeItem `json:"RuleChangeItems,omitnil" name:"RuleChangeItems"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// 修改数据
	Data []*SequenceData `json:"Data,omitnil" name:"Data"`

	// NAT地域
	Area *string `json:"Area,omitnil" name:"Area"`

	// 方向，0：出向，1：入向
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`
}

type ModifySequenceRulesRequest struct {
	*tchttp.BaseRequest
	
	// 边Id值
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// 修改数据
	Data []*SequenceData `json:"Data,omitnil" name:"Data"`

	// NAT地域
	Area *string `json:"Area,omitnil" name:"Area"`

	// 方向，0：出向，1：入向
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// 状态值，1：锁表，2：解锁表
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// Nat所在地域
	Area *string `json:"Area,omitnil" name:"Area"`

	// 0： 出向，1：入向
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`
}

type ModifyTableStatusRequest struct {
	*tchttp.BaseRequest
	
	// EdgeId值两个vpc间的边id
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// 状态值，1：锁表，2：解锁表
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// Nat所在地域
	Area *string `json:"Area,omitnil" name:"Area"`

	// 0： 出向，1：入向
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Rules []*VpcRuleItem `json:"Rules,omitnil" name:"Rules"`
}

type ModifyVpcAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要编辑的规则数组
	Rules []*VpcRuleItem `json:"Rules,omitnil" name:"Rules"`
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
	RuleUuids []*int64 `json:"RuleUuids,omitnil" name:"RuleUuids"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	FwGroupId *string `json:"FwGroupId,omitnil" name:"FwGroupId"`

	// 修改防火墙(组)名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 编辑的防火墙实例列表
	VpcFwInstances []*VpcFwInstance `json:"VpcFwInstances,omitnil" name:"VpcFwInstances"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil" name:"FwCidrInfo"`
}

type ModifyVpcFwGroupRequest struct {
	*tchttp.BaseRequest
	
	// 编辑的防火墙(组)ID
	FwGroupId *string `json:"FwGroupId,omitnil" name:"FwGroupId"`

	// 修改防火墙(组)名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 编辑的防火墙实例列表
	VpcFwInstances []*VpcFwInstance `json:"VpcFwInstances,omitnil" name:"VpcFwInstances"`

	// 指定防火墙使用网段信息
	FwCidrInfo *FwCidrInfo `json:"FwCidrInfo,omitnil" name:"FwCidrInfo"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RuleChangeItems []*RuleChangeItem `json:"RuleChangeItems,omitnil" name:"RuleChangeItems"`
}

type ModifyVpcFwSequenceRulesRequest struct {
	*tchttp.BaseRequest
	
	// 规则快速排序：OrderIndex，原始序号；NewOrderIndex：新序号
	RuleChangeItems []*RuleChangeItem `json:"RuleChangeItems,omitnil" name:"RuleChangeItems"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type NatFwEipsInfo struct {
	// 弹性公网ip
	Eip *string `json:"Eip,omitnil" name:"Eip"`

	// 所属的Nat网关Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatGatewayId *string `json:"NatGatewayId,omitnil" name:"NatGatewayId"`

	// Nat网关名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatGatewayName *string `json:"NatGatewayName,omitnil" name:"NatGatewayName"`
}

type NatFwFilter struct {
	// 过滤的类型，例如实例id
	FilterType *string `json:"FilterType,omitnil" name:"FilterType"`

	// 过滤的内容，以',' 分隔
	FilterContent *string `json:"FilterContent,omitnil" name:"FilterContent"`
}

type NatFwInstance struct {
	// nat实例id
	NatinsId *string `json:"NatinsId,omitnil" name:"NatinsId"`

	// nat实例名称
	NatinsName *string `json:"NatinsName,omitnil" name:"NatinsName"`

	// 实例所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 0:新增模式，1:接入模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwMode *int64 `json:"FwMode,omitnil" name:"FwMode"`

	// 0:正常状态， 1: 正在创建
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// nat公网ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatIp *string `json:"NatIp,omitnil" name:"NatIp"`
}

type NatInstanceInfo struct {
	// nat实例id
	NatinsId *string `json:"NatinsId,omitnil" name:"NatinsId"`

	// nat实例名称
	NatinsName *string `json:"NatinsName,omitnil" name:"NatinsName"`

	// 实例所在地域
	Region *string `json:"Region,omitnil" name:"Region"`

	// 0: 新增模式，1:接入模式
	FwMode *int64 `json:"FwMode,omitnil" name:"FwMode"`

	// 实例带宽大小 Mbps
	BandWidth *int64 `json:"BandWidth,omitnil" name:"BandWidth"`

	// 入向带宽峰值 bps
	InFlowMax *int64 `json:"InFlowMax,omitnil" name:"InFlowMax"`

	// 出向带宽峰值 bps
	OutFlowMax *uint64 `json:"OutFlowMax,omitnil" name:"OutFlowMax"`

	// 地域中文信息
	RegionZh *string `json:"RegionZh,omitnil" name:"RegionZh"`

	// 公网ip数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	EipAddress []*string `json:"EipAddress,omitnil" name:"EipAddress"`

	// 内外使用ip数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcIp []*string `json:"VpcIp,omitnil" name:"VpcIp"`

	// 实例关联子网数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Subnets []*string `json:"Subnets,omitnil" name:"Subnets"`

	// 0 :正常 1：正在初始化
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 地域区域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionDetail *string `json:"RegionDetail,omitnil" name:"RegionDetail"`

	// 实例所在可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneZh *string `json:"ZoneZh,omitnil" name:"ZoneZh"`

	// 实例所在可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneZhBak *string `json:"ZoneZhBak,omitnil" name:"ZoneZhBak"`

	// 已使用规则数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleUsed *uint64 `json:"RuleUsed,omitnil" name:"RuleUsed"`

	// 实例的规则限制最大规格数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleMax *uint64 `json:"RuleMax,omitnil" name:"RuleMax"`

	// 实例引擎版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineVersion *string `json:"EngineVersion,omitnil" name:"EngineVersion"`

	// 引擎是否可升级：0，不可升级；1，可升级
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateEnable *int64 `json:"UpdateEnable,omitnil" name:"UpdateEnable"`

	// 是的需要升级引擎 支持 nat拨测 1需要 0不需要
	// 注意：此字段可能返回 null，表示取不到有效值。
	NeedProbeEngineUpdate *int64 `json:"NeedProbeEngineUpdate,omitnil" name:"NeedProbeEngineUpdate"`
}

type NatSwitchListData struct {
	// 列表ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 子网名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetName *string `json:"SubnetName,omitnil" name:"SubnetName"`

	// IPv4 CIDR
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetCidr *string `json:"SubnetCidr,omitnil" name:"SubnetCidr"`

	// 关联路由ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteId *string `json:"RouteId,omitnil" name:"RouteId"`

	// 关联路由名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteName *string `json:"RouteName,omitnil" name:"RouteName"`

	// 云服务器个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CvmNum *uint64 `json:"CvmNum,omitnil" name:"CvmNum"`

	// 所属VPC ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 所属VPC名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// 是否生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enable *uint64 `json:"Enable,omitnil" name:"Enable"`

	// 开关状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// NAT网关ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatId *string `json:"NatId,omitnil" name:"NatId"`

	// NAT网关名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatName *string `json:"NatName,omitnil" name:"NatName"`

	// NAT防火墙实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatInsId *string `json:"NatInsId,omitnil" name:"NatInsId"`

	// NAT防火墙实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatInsName *string `json:"NatInsName,omitnil" name:"NatInsName"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 开关是否异常,0:正常,1:异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	Abnormal *int64 `json:"Abnormal,omitnil" name:"Abnormal"`
}

type NetInstancesInfo struct {
	// 网络实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 网络实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 网络cidr (多段以逗号分隔)
	InstanceCidr *string `json:"InstanceCidr,omitnil" name:"InstanceCidr"`

	// 网络实例所在地域
	Region *string `json:"Region,omitnil" name:"Region"`
}

type NewModeItems struct {
	// 新增模式下接入的vpc列表
	VpcList []*string `json:"VpcList,omitnil" name:"VpcList"`

	// 新增模式下绑定的出口弹性公网ip列表，其中Eips和AddCount至少传递一个。
	Eips []*string `json:"Eips,omitnil" name:"Eips"`

	// 新增模式下新增绑定的出口弹性公网ip个数，其中Eips和AddCount至少传递一个。
	AddCount *int64 `json:"AddCount,omitnil" name:"AddCount"`
}

// Predefined struct for user
type RemoveAcRuleRequestParams struct {
	// 规则的uuid，可通过查询规则列表获取
	RuleUuid *int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`
}

type RemoveAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则的uuid，可通过查询规则列表获取
	RuleUuid *int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`
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
	RuleUuid *int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`

	// 0代表成功，-1代表失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnCode *int64 `json:"ReturnCode,omitnil" name:"ReturnCode"`

	// success代表成功，failed代表失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnMsg *string `json:"ReturnMsg,omitnil" name:"ReturnMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RuleUuid []*int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`
}

type RemoveAclRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则的uuid列表，可通过查询规则列表获取，注意：如果传入的是[-1]将删除所有规则
	RuleUuid []*int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`
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
	RuleUuid []*int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RuleUuid *int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`

	// 删除类型，0是单条删除，RuleUuid填写删除规则id，1为全部删除，RuleUuid填0即可
	RemoveType *int64 `json:"RemoveType,omitnil" name:"RemoveType"`
}

type RemoveEnterpriseSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则的uuid，可通过查询规则列表获取
	RuleUuid *int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`

	// 删除类型，0是单条删除，RuleUuid填写删除规则id，1为全部删除，RuleUuid填0即可
	RemoveType *int64 `json:"RemoveType,omitnil" name:"RemoveType"`
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
	RuleUuid *int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`

	// 0代表成功，-1代表失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RuleUuid []*int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`
}

type RemoveNatAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则的uuid列表，可通过查询规则列表获取，注意：如果传入的是[-1]将删除所有规则
	RuleUuid []*int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`

	// 规则方向：1，入站；0，出站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`
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
	RuleUuid []*int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	RuleUuids []*int64 `json:"RuleUuids,omitnil" name:"RuleUuids"`
}

type RemoveVpcAcRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则的uuid列表，可通过查询规则列表获取，注意：如果传入的是[-1]将删除所有规则
	RuleUuids []*int64 `json:"RuleUuids,omitnil" name:"RuleUuids"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveVpcAcRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveVpcAcRuleResponseParams struct {
	// 删除成功后返回被删除策略的uuid列表
	RuleUuids []*int64 `json:"RuleUuids,omitnil" name:"RuleUuids"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	OrderIndex *int64 `json:"OrderIndex,omitnil" name:"OrderIndex"`

	// 新的sequence 值
	NewOrderIndex *int64 `json:"NewOrderIndex,omitnil" name:"NewOrderIndex"`
}

type RuleInfoData struct {
	// 执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitnil" name:"OrderIndex"`

	// 访问源
	SourceIp *string `json:"SourceIp,omitnil" name:"SourceIp"`

	// 访问目的
	TargetIp *string `json:"TargetIp,omitnil" name:"TargetIp"`

	// 协议
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 策略, 0：观察，1：阻断，2：放行
	Strategy *string `json:"Strategy,omitnil" name:"Strategy"`

	// 访问源类型，1是IP，3是域名，4是IP地址模版，5是域名地址模版
	SourceType *uint64 `json:"SourceType,omitnil" name:"SourceType"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 描述
	Detail *string `json:"Detail,omitnil" name:"Detail"`

	// 访问目的类型，1是IP，3是域名，4是IP地址模版，5是域名地址模版
	TargetType *uint64 `json:"TargetType,omitnil" name:"TargetType"`

	// 端口
	Port *string `json:"Port,omitnil" name:"Port"`

	// id值
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 日志id，从告警处创建必传，其它为空
	LogId *string `json:"LogId,omitnil" name:"LogId"`

	// 城市Code
	City *uint64 `json:"City,omitnil" name:"City"`

	// 国家Code
	Country *uint64 `json:"Country,omitnil" name:"Country"`

	// 云厂商，支持多个，以逗号分隔， 1:腾讯云（仅中国香港及海外）,2:阿里云,3:亚马逊云,4:华为云,5:微软云
	CloudCode *string `json:"CloudCode,omitnil" name:"CloudCode"`

	// 是否为地域
	IsRegion *uint64 `json:"IsRegion,omitnil" name:"IsRegion"`

	// 城市名
	CityName *string `json:"CityName,omitnil" name:"CityName"`

	// 国家名
	CountryName *string `json:"CountryName,omitnil" name:"CountryName"`

	// 国家二位iso代码或者省份缩写代码
	RegionIso *string `json:"RegionIso,omitnil" name:"RegionIso"`
}

type ScanInfo struct {
	// 扫描结果信息
	ScanResultInfo *ScanResultInfo `json:"ScanResultInfo,omitnil" name:"ScanResultInfo"`

	// 扫描状态 0扫描中 1完成  2未勾选自动扫描
	ScanStatus *int64 `json:"ScanStatus,omitnil" name:"ScanStatus"`

	// 进度
	ScanPercent *float64 `json:"ScanPercent,omitnil" name:"ScanPercent"`

	// 预计完成时间
	ScanTime *string `json:"ScanTime,omitnil" name:"ScanTime"`
}

type ScanResultInfo struct {
	// 暴露漏洞数量
	LeakNum *uint64 `json:"LeakNum,omitnil" name:"LeakNum"`

	// 防护ip数量
	IPNum *uint64 `json:"IPNum,omitnil" name:"IPNum"`

	// 暴露端口数量
	PortNum *uint64 `json:"PortNum,omitnil" name:"PortNum"`

	// 是否开启防护
	IPStatus *bool `json:"IPStatus,omitnil" name:"IPStatus"`

	// 是否拦截攻击
	IdpStatus *bool `json:"IdpStatus,omitnil" name:"IdpStatus"`

	// 是否禁封端口
	BanStatus *bool `json:"BanStatus,omitnil" name:"BanStatus"`
}

type SecurityGroupBothWayInfo struct {
	// 执行顺序
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderIndex *uint64 `json:"OrderIndex,omitnil" name:"OrderIndex"`

	// 访问源
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceId *string `json:"SourceId,omitnil" name:"SourceId"`

	// 访问源类型，默认为0，0: IP, 1: VPC, 2: SUBNET, 3: CVM, 4: CLB, 5: ENI, 6: CDB, 7: 参数模板, 100: 资产分组
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceType *uint64 `json:"SourceType,omitnil" name:"SourceType"`

	// 访问目的
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetId *string `json:"TargetId,omitnil" name:"TargetId"`

	// 访问目的类型，默认为0，0: IP, 1: VPC, 2: SUBNET, 3: CVM, 4: CLB, 5: ENI, 6: CDB, 7: 参数模板, 100: 资产分组
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetType *uint64 `json:"TargetType,omitnil" name:"TargetType"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 目的端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitnil" name:"Port"`

	// 策略, 1：阻断，2：放行
	// 注意：此字段可能返回 null，表示取不到有效值。
	Strategy *uint64 `json:"Strategy,omitnil" name:"Strategy"`

	// 方向，0：出站，1：入站，默认1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 地域
	Region *string `json:"Region,omitnil" name:"Region"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *string `json:"Detail,omitnil" name:"Detail"`

	// 是否开关开启，0：未开启，1：开启
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 是否是正常规则，0：正常，1：异常
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsNew *uint64 `json:"IsNew,omitnil" name:"IsNew"`

	// 单/双向下发，0:单向下发，1：双向下发
	// 注意：此字段可能返回 null，表示取不到有效值。
	BothWay *uint64 `json:"BothWay,omitnil" name:"BothWay"`

	// 私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 公网IP，多个以英文逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 内网IP，多个以英文逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 掩码地址，多个以英文逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cidr *string `json:"Cidr,omitnil" name:"Cidr"`

	// 端口协议类型参数模板id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil" name:"ServiceTemplateId"`

	// 是否使用端口协议模板，0：否，1：是
	ProtocolPortType *uint64 `json:"ProtocolPortType,omitnil" name:"ProtocolPortType"`
}

type SecurityGroupListData struct {
	// 执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitnil" name:"OrderIndex"`

	// 访问源
	SourceId *string `json:"SourceId,omitnil" name:"SourceId"`

	// 访问源类型，默认为0，1: VPC, 2: SUBNET, 3: CVM, 4: CLB, 5: ENI, 6: CDB, 7: 参数模板, 100: 资源组
	SourceType *uint64 `json:"SourceType,omitnil" name:"SourceType"`

	// 访问目的
	TargetId *string `json:"TargetId,omitnil" name:"TargetId"`

	// 访问目的类型，默认为0，1: VPC, 2: SUBNET, 3: CVM, 4: CLB, 5: ENI, 6: CDB, 7: 参数模板, 100:资源组
	TargetType *uint64 `json:"TargetType,omitnil" name:"TargetType"`

	// 协议
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 目的端口
	Port *string `json:"Port,omitnil" name:"Port"`

	// 策略, 1：阻断，2：放行
	Strategy *uint64 `json:"Strategy,omitnil" name:"Strategy"`

	// 描述
	Detail *string `json:"Detail,omitnil" name:"Detail"`

	// 单/双向下发，0:单向下发，1：双向下发
	BothWay *uint64 `json:"BothWay,omitnil" name:"BothWay"`

	// 规则ID
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 是否开关开启，0：未开启，1：开启
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 是否是正常规则，0：正常，1：异常
	IsNew *uint64 `json:"IsNew,omitnil" name:"IsNew"`

	// 私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 公网IP，多个以英文逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 内网IP，多个以英文逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIp *string `json:"PrivateIp,omitnil" name:"PrivateIp"`

	// 掩码地址，多个以英文逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cidr *string `json:"Cidr,omitnil" name:"Cidr"`

	// 端口协议类型参数模板id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil" name:"ServiceTemplateId"`

	// 生成双向下发规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	BothWayInfo []*SecurityGroupBothWayInfo `json:"BothWayInfo,omitnil" name:"BothWayInfo"`

	// 方向，0：出站，1：入站，默认1
	Direction *uint64 `json:"Direction,omitnil" name:"Direction"`

	// 是否使用端口协议模板，0：否，1：是
	ProtocolPortType *uint64 `json:"ProtocolPortType,omitnil" name:"ProtocolPortType"`

	// Uuid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 资产分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetGroupNameIn *string `json:"AssetGroupNameIn,omitnil" name:"AssetGroupNameIn"`

	// 资产分组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssetGroupNameOut *string `json:"AssetGroupNameOut,omitnil" name:"AssetGroupNameOut"`

	// 模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParameterName *string `json:"ParameterName,omitnil" name:"ParameterName"`
}

type SecurityGroupOrderIndexData struct {
	// 企业安全组规则当前执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitnil" name:"OrderIndex"`

	// 企业安全组规则更新目标执行顺序
	NewOrderIndex *uint64 `json:"NewOrderIndex,omitnil" name:"NewOrderIndex"`
}

type SecurityGroupRule struct {
	// 访问源示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板id(ipm-dyodhpby)
	// instance：资产实例id(ins-123456)
	// resourcegroup：资产分组id(cfwrg-xxxx)
	// tag：资源标签({\"Key\":\"标签key值\",\"Value\":\"标签Value值\"})
	// region：地域(ap-gaungzhou)
	SourceContent *string `json:"SourceContent,omitnil" name:"SourceContent"`

	// 访问源类型，类型可以为以下6种：net|template|instance|resourcegroup|tag|region
	SourceType *string `json:"SourceType,omitnil" name:"SourceType"`

	// 访问目的示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板id(ipm-dyodhpby)
	// instance：资产实例id(ins-123456)
	// resourcegroup：资产分组id(cfwrg-xxxx)
	// tag：资源标签({\"Key\":\"标签key值\",\"Value\":\"标签Value值\"})
	// region：地域(ap-gaungzhou)
	DestContent *string `json:"DestContent,omitnil" name:"DestContent"`

	// 访问目的类型，类型可以为以下6种：net|template|instance|resourcegroup|tag|region
	DestType *string `json:"DestType,omitnil" name:"DestType"`

	// 访问控制策略中设置的流量通过云防火墙的方式。取值：
	// accept：放行
	// drop：拒绝
	RuleAction *string `json:"RuleAction,omitnil" name:"RuleAction"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 规则顺序，-1表示最低，1表示最高，请勿和外层Type冲突（和外层的Type配合使用，当中间插入时，指定添加位置）
	OrderIndex *string `json:"OrderIndex,omitnil" name:"OrderIndex"`

	// 协议；TCP/UDP/ICMP/ANY
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 访问控制策略的端口。取值：
	// -1/-1：全部端口
	// 80：80端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitnil" name:"Port"`

	// 端口协议类型参数模板id；协议端口模板id；与Protocol,Port互斥
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceTemplateId *string `json:"ServiceTemplateId,omitnil" name:"ServiceTemplateId"`

	// （入参时无需填写，自动生成）规则对应的唯一id
	Id *string `json:"Id,omitnil" name:"Id"`

	// （入参时、Enable已弃用；由通用配置中新增规则启用状态控制）
	// 规则状态，true表示启用，false表示禁用
	Enable *string `json:"Enable,omitnil" name:"Enable"`
}

type SecurityGroupSimplifyRule struct {
	// 访问源示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板(ipm-dyodhpby)
	// instance：资产实例(ins-123456)
	// resourcegroup：资产分组(/全部分组/分组1/子分组1)
	// tag：资源标签({"Key":"标签key值","Value":"标签Value值"})
	// region：地域(ap-gaungzhou)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceContent *string `json:"SourceContent,omitnil" name:"SourceContent"`

	// 访问目的示例：
	// net：IP/CIDR(192.168.0.2)
	// template：参数模板(ipm-dyodhpby)
	// instance：资产实例(ins-123456)
	// resourcegroup：资产分组(/全部分组/分组1/子分组1)
	// tag：资源标签({"Key":"标签key值","Value":"标签Value值"})
	// region：地域(ap-gaungzhou)
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestContent *string `json:"DestContent,omitnil" name:"DestContent"`

	// 协议；TCP/UDP/ICMP/ANY
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 规则对应的唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleUuid *int64 `json:"RuleUuid,omitnil" name:"RuleUuid"`

	// 规则序号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sequence *int64 `json:"Sequence,omitnil" name:"Sequence"`
}

type SequenceData struct {
	// 规则Id值
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 修改前执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitnil" name:"OrderIndex"`

	// 修改后执行顺序
	NewOrderIndex *uint64 `json:"NewOrderIndex,omitnil" name:"NewOrderIndex"`
}

// Predefined struct for user
type SetNatFwDnatRuleRequestParams struct {
	// 0：cfw新增模式，1：cfw接入模式。
	Mode *uint64 `json:"Mode,omitnil" name:"Mode"`

	// 操作类型，可选值：add，del，modify。
	OperationType *string `json:"OperationType,omitnil" name:"OperationType"`

	// 防火墙实例id，该字段必须传递。
	CfwInstance *string `json:"CfwInstance,omitnil" name:"CfwInstance"`

	// 添加或删除操作的Dnat规则列表。
	AddOrDelDnatRules []*CfwNatDnatRule `json:"AddOrDelDnatRules,omitnil" name:"AddOrDelDnatRules"`

	// 修改操作的原始Dnat规则
	OriginDnat *CfwNatDnatRule `json:"OriginDnat,omitnil" name:"OriginDnat"`

	// 修改操作的新的Dnat规则
	NewDnat *CfwNatDnatRule `json:"NewDnat,omitnil" name:"NewDnat"`
}

type SetNatFwDnatRuleRequest struct {
	*tchttp.BaseRequest
	
	// 0：cfw新增模式，1：cfw接入模式。
	Mode *uint64 `json:"Mode,omitnil" name:"Mode"`

	// 操作类型，可选值：add，del，modify。
	OperationType *string `json:"OperationType,omitnil" name:"OperationType"`

	// 防火墙实例id，该字段必须传递。
	CfwInstance *string `json:"CfwInstance,omitnil" name:"CfwInstance"`

	// 添加或删除操作的Dnat规则列表。
	AddOrDelDnatRules []*CfwNatDnatRule `json:"AddOrDelDnatRules,omitnil" name:"AddOrDelDnatRules"`

	// 修改操作的原始Dnat规则
	OriginDnat *CfwNatDnatRule `json:"OriginDnat,omitnil" name:"OriginDnat"`

	// 修改操作的新的Dnat规则
	NewDnat *CfwNatDnatRule `json:"NewDnat,omitnil" name:"NewDnat"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	OperationType *string `json:"OperationType,omitnil" name:"OperationType"`

	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitnil" name:"CfwInstance"`

	// 当OperationType 为bind或unbind操作时，使用该字段。
	EipList []*string `json:"EipList,omitnil" name:"EipList"`
}

type SetNatFwEipRequest struct {
	*tchttp.BaseRequest
	
	// bind：绑定eip；unbind：解绑eip；newAdd：新增防火墙弹性公网ip
	OperationType *string `json:"OperationType,omitnil" name:"OperationType"`

	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitnil" name:"CfwInstance"`

	// 当OperationType 为bind或unbind操作时，使用该字段。
	EipList []*string `json:"EipList,omitnil" name:"EipList"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type StaticInfo struct {
	// 数
	Num *int64 `json:"Num,omitnil" name:"Num"`

	// 端口
	Port *string `json:"Port,omitnil" name:"Port"`

	// ip信息
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 地址
	Address *string `json:"Address,omitnil" name:"Address"`

	// 资产id
	InsID *string `json:"InsID,omitnil" name:"InsID"`

	// 资产名称
	InsName *string `json:"InsName,omitnil" name:"InsName"`
}

// Predefined struct for user
type StopSecurityGroupRuleDispatchRequestParams struct {
	// 值为1，中止全部
	StopType *int64 `json:"StopType,omitnil" name:"StopType"`
}

type StopSecurityGroupRuleDispatchRequest struct {
	*tchttp.BaseRequest
	
	// 值为1，中止全部
	StopType *int64 `json:"StopType,omitnil" name:"StopType"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *bool `json:"Status,omitnil" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type SwitchListsData struct {
	// 公网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitnil" name:"PublicIp"`

	// 内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntranetIp *string `json:"IntranetIp,omitnil" name:"IntranetIp"`

	// 实例名
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 资产类型
	AssetType *string `json:"AssetType,omitnil" name:"AssetType"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Area *string `json:"Area,omitnil" name:"Area"`

	// 防火墙开关
	Switch *int64 `json:"Switch,omitnil" name:"Switch"`

	// id值
	Id *uint64 `json:"Id,omitnil" name:"Id"`

	// 公网 IP 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpType *uint64 `json:"PublicIpType,omitnil" name:"PublicIpType"`

	// 风险端口数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortTimes *uint64 `json:"PortTimes,omitnil" name:"PortTimes"`

	// 最近扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTime *string `json:"LastTime,omitnil" name:"LastTime"`

	// 扫描深度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanMode *string `json:"ScanMode,omitnil" name:"ScanMode"`

	// 扫描状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanStatus *uint64 `json:"ScanStatus,omitnil" name:"ScanStatus"`
}

// Predefined struct for user
type SyncFwOperateRequestParams struct {
	// 同步操作类型：Route，同步防火墙路由
	SyncType *string `json:"SyncType,omitnil" name:"SyncType"`

	// 防火墙类型；nat,nat防火墙;ew,vpc间防火墙
	FwType *string `json:"FwType,omitnil" name:"FwType"`
}

type SyncFwOperateRequest struct {
	*tchttp.BaseRequest
	
	// 同步操作类型：Route，同步防火墙路由
	SyncType *string `json:"SyncType,omitnil" name:"SyncType"`

	// 防火墙类型；nat,nat防火墙;ew,vpc间防火墙
	FwType *string `json:"FwType,omitnil" name:"FwType"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 失陷主机
	OutNum *int64 `json:"OutNum,omitnil" name:"OutNum"`

	// 待处置告警
	HandleNum *int64 `json:"HandleNum,omitnil" name:"HandleNum"`

	// 漏洞攻击
	VulNum *int64 `json:"VulNum,omitnil" name:"VulNum"`

	// 网络探测
	NetworkNum *int64 `json:"NetworkNum,omitnil" name:"NetworkNum"`

	// 封禁列表
	BanNum *int64 `json:"BanNum,omitnil" name:"BanNum"`

	// 暴力破解
	BruteForceNum *int64 `json:"BruteForceNum,omitnil" name:"BruteForceNum"`
}

type TemplateListInfo struct {
	// 模版ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uuid *string `json:"Uuid,omitnil" name:"Uuid"`

	// 模版名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *string `json:"Detail,omitnil" name:"Detail"`

	// IP模版
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpString *string `json:"IpString,omitnil" name:"IpString"`

	// 插入时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	InsertTime *string `json:"InsertTime,omitnil" name:"InsertTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 模版类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil" name:"Type"`

	// 关联规则条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RulesNum *int64 `json:"RulesNum,omitnil" name:"RulesNum"`
}

type UnHandleEvent struct {
	// 伪攻击链类型
	EventTableListStruct []*UnHandleEventDetail `json:"EventTableListStruct,omitnil" name:"EventTableListStruct"`

	// 1 是  0否
	BaseLineUser *uint64 `json:"BaseLineUser,omitnil" name:"BaseLineUser"`

	// 1 打开 0 关闭
	BaseLineInSwitch *uint64 `json:"BaseLineInSwitch,omitnil" name:"BaseLineInSwitch"`

	// 1 打开 0 关闭
	BaseLineOutSwitch *uint64 `json:"BaseLineOutSwitch,omitnil" name:"BaseLineOutSwitch"`

	// vpc间防火墙实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcFwCount *uint64 `json:"VpcFwCount,omitnil" name:"VpcFwCount"`
}

type UnHandleEventDetail struct {
	// 安全事件名称
	EventName *string `json:"EventName,omitnil" name:"EventName"`

	// 未处置事件数量
	Total *uint64 `json:"Total,omitnil" name:"Total"`
}

type VpcDnsInfo struct {
	// vpc id
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// vpc 名称
	VpcName *string `json:"VpcName,omitnil" name:"VpcName"`

	// nat 防火墙模式 0：新增模式， 1: 接入模式
	FwMode *int64 `json:"FwMode,omitnil" name:"FwMode"`

	// vpc ipv4网段范围 CIDR（Classless Inter-Domain Routing，无类域间路由选择）
	VpcIpv4Cidr *string `json:"VpcIpv4Cidr,omitnil" name:"VpcIpv4Cidr"`

	// 外网弹性ip，防火墙 dns解析地址
	DNSEip *string `json:"DNSEip,omitnil" name:"DNSEip"`

	// nat网关id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatInsId *string `json:"NatInsId,omitnil" name:"NatInsId"`

	// nat网关名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatInsName *string `json:"NatInsName,omitnil" name:"NatInsName"`

	// 0：开关关闭 ， 1: 开关打开
	SwitchStatus *int64 `json:"SwitchStatus,omitnil" name:"SwitchStatus"`

	// 0：未防护， 1: 已防护，2：忽略此字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectedStatus *uint64 `json:"ProtectedStatus,omitnil" name:"ProtectedStatus"`

	// 是否支持DNS FW，0-不支持、1-支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportDNSFW *uint64 `json:"SupportDNSFW,omitnil" name:"SupportDNSFW"`
}

type VpcFwCvmInsInfo struct {
	// VPC防火墙实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwInsId *string `json:"FwInsId,omitnil" name:"FwInsId"`

	// CVM所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// CVM所在地域中文
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionZh *string `json:"RegionZh,omitnil" name:"RegionZh"`

	// CVM所在地域详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionDetail *string `json:"RegionDetail,omitnil" name:"RegionDetail"`

	// 主机所在可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneZh *string `json:"ZoneZh,omitnil" name:"ZoneZh"`

	// 备机所在可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneZhBack *string `json:"ZoneZhBack,omitnil" name:"ZoneZhBack"`

	// 防火墙CVM带宽值
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandWidth *int64 `json:"BandWidth,omitnil" name:"BandWidth"`
}

type VpcFwGroupInfo struct {
	// 防火墙(组)ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwGroupId *string `json:"FwGroupId,omitnil" name:"FwGroupId"`

	// 防火墙(组)名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwGroupName *string `json:"FwGroupName,omitnil" name:"FwGroupName"`

	// 防火墙组涉及到的开关个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwSwitchNum *int64 `json:"FwSwitchNum,omitnil" name:"FwSwitchNum"`

	// 防火墙(组)部署的地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionLst []*string `json:"RegionLst,omitnil" name:"RegionLst"`

	// 模式 1：CCN云联网模式；0：私有网络模式 2: sase 模式 3：ccn 高级模式 4: 私有网络(跨租户单边模式)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *int64 `json:"Mode,omitnil" name:"Mode"`

	// 防火墙实例的开关模式 1: 单点互通 2: 多点互通 3: 全互通 4: 自定义路由
	// 注意：此字段可能返回 null，表示取不到有效值。
	SwitchMode *int64 `json:"SwitchMode,omitnil" name:"SwitchMode"`

	// VPC防火墙实例卡片信息数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwInstanceLst []*VpcFwInstanceInfo `json:"FwInstanceLst,omitnil" name:"FwInstanceLst"`

	// 防火墙(状态) 0：正常 1: 初始化或操作中
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// auto :自动选择
	// 如果为网段，则为用户自定义 192.168.0.0/20 
	FwVpcCidr *string `json:"FwVpcCidr,omitnil" name:"FwVpcCidr"`

	// cdc专用集群场景时表示部署所属的cdc
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdcId *string `json:"CdcId,omitnil" name:"CdcId"`

	// cdc专用集群场景时表示cdc名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdcName *string `json:"CdcName,omitnil" name:"CdcName"`

	// 跨租户模式 1管理员 2单边 0 非跨租户
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrossUserMode *string `json:"CrossUserMode,omitnil" name:"CrossUserMode"`
}

type VpcFwInstance struct {
	// 防火墙实例名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 私有网络模式下接入的VpcId列表；仅私有网络模式使用
	VpcIds []*string `json:"VpcIds,omitnil" name:"VpcIds"`

	// 部署地域信息
	FwDeploy *FwDeploy `json:"FwDeploy,omitnil" name:"FwDeploy"`

	// 防火墙实例ID (编辑场景传)
	FwInsId *string `json:"FwInsId,omitnil" name:"FwInsId"`
}

type VpcFwInstanceInfo struct {
	// VPC防火墙实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwInsName *string `json:"FwInsName,omitnil" name:"FwInsName"`

	// VPC防火墙实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwInsId *string `json:"FwInsId,omitnil" name:"FwInsId"`

	// VPC防火墙实例模式 0: 旧VPC模式防火墙 1: CCN模式防火墙
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwMode *int64 `json:"FwMode,omitnil" name:"FwMode"`

	// VPC防火墙接入网络实例个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	JoinInsNum *int64 `json:"JoinInsNum,omitnil" name:"JoinInsNum"`

	// VPC防火墙开关个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwSwitchNum *int64 `json:"FwSwitchNum,omitnil" name:"FwSwitchNum"`

	// VPC防火墙状态 0:正常 ， 1：创建中 2: 变更中
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// VPC防火墙创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *string `json:"Time,omitnil" name:"Time"`

	// VPC 相关云联网ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnId []*string `json:"CcnId,omitnil" name:"CcnId"`

	// VPC 相关云联网名称列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnName []*string `json:"CcnName,omitnil" name:"CcnName"`

	// VPC 相关对等连接ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeerConnectionId []*string `json:"PeerConnectionId,omitnil" name:"PeerConnectionId"`

	// VPC 相关对等连接名称列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeerConnectionName []*string `json:"PeerConnectionName,omitnil" name:"PeerConnectionName"`

	// VPC防火墙CVM的列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwCvmLst []*VpcFwCvmInsInfo `json:"FwCvmLst,omitnil" name:"FwCvmLst"`

	// VPC防火墙接入网络实例类型列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	JoinInsLst []*VpcFwJoinInstanceType `json:"JoinInsLst,omitnil" name:"JoinInsLst"`

	// 防火墙网关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwGateway []*FwGateway `json:"FwGateway,omitnil" name:"FwGateway"`

	// 防火墙(组)ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwGroupId *string `json:"FwGroupId,omitnil" name:"FwGroupId"`

	// 已使用规则数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleUsed *int64 `json:"RuleUsed,omitnil" name:"RuleUsed"`

	// 最大规则数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleMax *int64 `json:"RuleMax,omitnil" name:"RuleMax"`

	// 防火墙实例带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 用户VPC墙总带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserVpcWidth *int64 `json:"UserVpcWidth,omitnil" name:"UserVpcWidth"`

	// 接入的vpc列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	JoinInsIdLst []*string `json:"JoinInsIdLst,omitnil" name:"JoinInsIdLst"`

	// 内网间峰值带宽 (单位 bps )
	FlowMax *int64 `json:"FlowMax,omitnil" name:"FlowMax"`

	// 实例引擎版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineVersion *string `json:"EngineVersion,omitnil" name:"EngineVersion"`

	// 引擎是否可升级：0，不可升级；1，可升级
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateEnable *int64 `json:"UpdateEnable,omitnil" name:"UpdateEnable"`
}

type VpcFwInstanceShow struct {
	// VPC防火墙实例ID
	FwInsId *string `json:"FwInsId,omitnil" name:"FwInsId"`

	// VPC防火墙实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwInsName *string `json:"FwInsName,omitnil" name:"FwInsName"`

	// 网络经过VPC防火墙CVM所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwInsRegion *string `json:"FwInsRegion,omitnil" name:"FwInsRegion"`
}

type VpcFwJoinInstanceType struct {
	// 接入实例类型，VPC、DIRECTCONNECT、 VPNGW 等
	// 注意：此字段可能返回 null，表示取不到有效值。
	JoinType *string `json:"JoinType,omitnil" name:"JoinType"`

	// 接入的对应网络实例类型的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Num *int64 `json:"Num,omitnil" name:"Num"`
}

type VpcRuleItem struct {
	// 访问源示例：
	// net：IP/CIDR(192.168.0.2)
	SourceContent *string `json:"SourceContent,omitnil" name:"SourceContent"`

	// 访问源类型，类型可以为：net
	SourceType *string `json:"SourceType,omitnil" name:"SourceType"`

	// 访问目的示例：
	// net：IP/CIDR(192.168.0.2)
	// domain：域名规则，例如*.qq.com
	DestContent *string `json:"DestContent,omitnil" name:"DestContent"`

	// 访问目的类型，类型可以为：net，domain
	DestType *string `json:"DestType,omitnil" name:"DestType"`

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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 访问控制策略中设置的流量通过云防火墙的方式。取值：
	// accept：放行
	// drop：拒绝
	// log：观察
	RuleAction *string `json:"RuleAction,omitnil" name:"RuleAction"`

	// 访问控制策略的端口。取值：
	// -1/-1：全部端口
	// 80：80端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitnil" name:"Port"`

	// 描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 规则顺序，-1表示最低，1表示最高
	OrderIndex *int64 `json:"OrderIndex,omitnil" name:"OrderIndex"`

	// 规则对应的唯一id
	Uuid *int64 `json:"Uuid,omitnil" name:"Uuid"`

	// 规则状态，true表示启用，false表示禁用
	Enable *string `json:"Enable,omitnil" name:"Enable"`

	// 规则生效的范围，是在哪对vpc之间还是针对所有vpc间生效
	EdgeId *string `json:"EdgeId,omitnil" name:"EdgeId"`

	// 规则的命中次数，增删改查规则时无需传入此参数，主要用于返回查询结果数据
	DetectedTimes *int64 `json:"DetectedTimes,omitnil" name:"DetectedTimes"`

	// EdgeId对应的这对VPC间防火墙的描述
	EdgeName *string `json:"EdgeName,omitnil" name:"EdgeName"`

	// 内部使用的uuid，一般情况下不会使用到该字段
	InternalUuid *int64 `json:"InternalUuid,omitnil" name:"InternalUuid"`

	// 规则被删除：1，已删除；0，未删除
	Deleted *int64 `json:"Deleted,omitnil" name:"Deleted"`

	// 规则生效的防火墙实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwGroupId *string `json:"FwGroupId,omitnil" name:"FwGroupId"`

	// 防火墙名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FwGroupName *string `json:"FwGroupName,omitnil" name:"FwGroupName"`

	// beta任务详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	BetaList []*BetaInfoByACL `json:"BetaList,omitnil" name:"BetaList"`
}

type VpcZoneData struct {
	// 可用区
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// vpc节点地域
	Region *string `json:"Region,omitnil" name:"Region"`
}