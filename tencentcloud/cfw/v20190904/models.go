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
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AcListsData struct {

	// 规则id
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 访问源
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 访问目的
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`

	// 协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *string `json:"Port,omitempty" name:"Port"`

	// 策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	Strategy *uint64 `json:"Strategy,omitempty" name:"Strategy"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// 命中次数
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// 告警规则id
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogId *string `json:"LogId,omitempty" name:"LogId"`
}

type AssociatedInstanceInfo struct {

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例类型，3是cvm实例,4是clb实例,5是eni实例,6是云数据库
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 公网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 关联安全组数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroupCount *uint64 `json:"SecurityGroupCount,omitempty" name:"SecurityGroupCount"`
}

type CfwNatDnatRule struct {

	// 网络协议，可选值：TCP、UDP。
	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`

	// 弹性IP。
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" name:"PublicIpAddress"`

	// 公网端口。
	PublicPort *int64 `json:"PublicPort,omitempty" name:"PublicPort"`

	// 内网地址。
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" name:"PrivateIpAddress"`

	// 内网端口。
	PrivatePort *int64 `json:"PrivatePort,omitempty" name:"PrivatePort"`

	// NAT防火墙转发规则描述。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateAcRulesRequest struct {
	*tchttp.BaseRequest

	// 创建规则数据
	Data []*RuleInfoData `json:"Data,omitempty" name:"Data" list`

	// 0：添加（默认），1：插入
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 边id
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// 访问控制规则状态
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// 0：添加，1：覆盖
	Overwrite *uint64 `json:"Overwrite,omitempty" name:"Overwrite"`

	// NAT实例ID, 参数Area存在的时候这个必传
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// portScan: 来自于端口扫描, patchImport: 来自于批量导入
	From *string `json:"From,omitempty" name:"From"`

	// NAT地域
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("CreateAcRulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAcRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态值，0:操作成功
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 返回多余的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Info *string `json:"Info,omitempty" name:"Info"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAcRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupApiRulesRequest struct {
	*tchttp.BaseRequest

	// 创建规则数据
	Data []*SecurityGroupApiRuleData `json:"Data,omitempty" name:"Data" list`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// 插入类型，0：后插，1：前插，2：中插
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 腾讯云地域的英文简写
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupApiRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	delete(f, "Direction")
	delete(f, "Type")
	delete(f, "Area")
	if len(f) > 0 {
		return errors.New("CreateSecurityGroupApiRulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSecurityGroupApiRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态值，0:添加成功，非0：添加失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityGroupApiRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAcRuleRequest struct {
	*tchttp.BaseRequest

	// 删除规则对应的id值, 对应获取规则列表接口的Id 值
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// EdgeId值两个vpc间的边id
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// NAT地域， 如ap-shanghai/ap-guangzhou/ap-chongqing等
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DeleteAcRuleRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态值 0: 删除成功, !0: 删除失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 返回多余的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Info *string `json:"Info,omitempty" name:"Info"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAllAccessControlRuleRequest struct {
	*tchttp.BaseRequest

	// 方向，0：出站，1：入站  默认值是 0
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// VPC间防火墙开关ID  全部删除 EdgeId和Area只填写一个，不填写则不删除vpc间防火墙开关 ，默认值为‘’
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// nat地域 全部删除 EdgeId和Area只填写一个，不填写则不删除nat防火墙开关 默认值为‘’
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DeleteAllAccessControlRuleRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAllAccessControlRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态值 0: 修改成功, !0: 修改失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 删除了几条访问控制规则
	// 注意：此字段可能返回 null，表示取不到有效值。
		Info *int64 `json:"Info,omitempty" name:"Info"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAllAccessControlRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupAllRuleRequest struct {
	*tchttp.BaseRequest

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// 腾讯云地域的英文简写
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupAllRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "Area")
	if len(f) > 0 {
		return errors.New("DeleteSecurityGroupAllRuleRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupAllRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0: 操作成功，非0：操作失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 返回数据的json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
		Info *int64 `json:"Info,omitempty" name:"Info"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupAllRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupRuleRequest struct {
	*tchttp.BaseRequest

	// 所需要删除规则的ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 腾讯云地域的英文简写
	Area *string `json:"Area,omitempty" name:"Area"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// 是否删除反向规则，0：否，1：是
	IsDelReverse *uint64 `json:"IsDelReverse,omitempty" name:"IsDelReverse"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DeleteSecurityGroupRuleRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSecurityGroupRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态值，0：成功，非0：失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 返回多余的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Info *string `json:"Info,omitempty" name:"Info"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityGroupRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAcListsRequest struct {
	*tchttp.BaseRequest

	// 协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 策略
	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// 每页条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 出站还是入站，0：入站，1：出站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// EdgeId值
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// 规则是否开启，'0': 未开启，'1': 开启, 默认为'0'
	Status *string `json:"Status,omitempty" name:"Status"`

	// 地域
	Area *string `json:"Area,omitempty" name:"Area"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeAcListsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAcListsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总条数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 访问控制列表数据
		Data []*AcListsData `json:"Data,omitempty" name:"Data" list`

		// 不算筛选条数的总条数
		AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`

		// 访问控制规则全部启用/全部停用
	// 注意：此字段可能返回 null，表示取不到有效值。
		Enable *uint64 `json:"Enable,omitempty" name:"Enable"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAcListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssociatedInstanceListRequest struct {
	*tchttp.BaseRequest

	// 列表偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页记录条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 地域代码（例：ap-guangzhou）,支持腾讯云全地域
	Area *string `json:"Area,omitempty" name:"Area"`

	// 额外检索条件（JSON字符串）
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`

	// 排序方式（asc:升序,desc:降序）
	Order *string `json:"Order,omitempty" name:"Order"`

	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`

	// 实例类型,'3'是cvm实例,'4'是clb实例,'5'是eni实例,'6'是云数据库
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeAssociatedInstanceListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAssociatedInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*AssociatedInstanceInfo `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssociatedInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwEipsRequest struct {
	*tchttp.BaseRequest

	// 0：cfw新增模式，1：cfw接入模式
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// ALL：查询所有弹性公网ip; nat-xxxxx：接入模式场景指定网关的弹性公网ip
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeCfwEipsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCfwEipsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回值信息
		NatFwEipList []*NatFwEipsInfo `json:"NatFwEipList,omitempty" name:"NatFwEipList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCfwEipsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatRuleOverviewRequest struct {
	*tchttp.BaseRequest

	// 方向，0：出站，1：入站 默认值：0
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// NAT地域  这个是必填项，填入相关的英文，'ap-beijing-fsi': '北京金融',
	//         'ap-beijing': '北京',
	//         'ap-changsha-ec': '长沙EC',
	//         'ap-chengdu': '成都',
	//         'ap-chongqing': '重庆',
	//         'ap-fuzhou-ec': '福州EC',
	//         'ap-guangzhou-open': '广州Open',
	//         'ap-guangzhou': '广州',
	//         'ap-hangzhou-ec': '杭州EC',
	//         'ap-jinan-ec': '济南EC',
	//         'ap-nanjing': '南京',
	//         'ap-shanghai-fsi': '上海金融',
	//         'ap-shanghai': '上海',
	//         'ap-shenzhen-fsi': '深圳金融',
	//         'ap-shenzhen': '深圳',
	//         'ap-wuhan-ec': '武汉EC'
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatRuleOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "Area")
	if len(f) > 0 {
		return errors.New("DescribeNatRuleOverviewRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNatRuleOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例ID
		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

		// 实例名称
		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

		// 弹性IP列表
		EipList []*string `json:"EipList,omitempty" name:"EipList" list`

		// 端口转发规则数量
		DnatNum *int64 `json:"DnatNum,omitempty" name:"DnatNum"`

		// 访问控制规则总数
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 访问控制规则剩余配额
		RemainNum *int64 `json:"RemainNum,omitempty" name:"RemainNum"`

		// 阻断规则条数
		BlockNum *int64 `json:"BlockNum,omitempty" name:"BlockNum"`

		// 启用规则条数
		EnableNum *int64 `json:"EnableNum,omitempty" name:"EnableNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNatRuleOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleOverviewRequest struct {
	*tchttp.BaseRequest

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	if len(f) > 0 {
		return errors.New("DescribeRuleOverviewRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`

		// 阻断策略规则数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		StrategyNum *uint64 `json:"StrategyNum,omitempty" name:"StrategyNum"`

		// 启用规则数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		StartRuleNum *uint64 `json:"StartRuleNum,omitempty" name:"StartRuleNum"`

		// 停用规则数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		StopRuleNum *uint64 `json:"StopRuleNum,omitempty" name:"StopRuleNum"`

		// 剩余配额
	// 注意：此字段可能返回 null，表示取不到有效值。
		RemainingNum *int64 `json:"RemainingNum,omitempty" name:"RemainingNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupListRequest struct {
	*tchttp.BaseRequest

	// 0: 出站规则，1：入站规则
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// 地域代码（例: ap-guangzhou),支持腾讯云全部地域
	Area *string `json:"Area,omitempty" name:"Area"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// 每页条数，默认为10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移值，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 状态，'': 全部，'0'：筛选停用规则，'1'：筛选启用规则
	Status *string `json:"Status,omitempty" name:"Status"`

	// 0: 不过滤，1：过滤掉正常规则，保留下发异常规则
	Filter *uint64 `json:"Filter,omitempty" name:"Filter"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeSecurityGroupListRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSecurityGroupListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表当前规则总条数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 安全组规则列表数据
		Data []*SecurityGroupListData `json:"Data,omitempty" name:"Data" list`

		// 不算筛选条数的总条数
		AllTotal *uint64 `json:"AllTotal,omitempty" name:"AllTotal"`

		// 访问控制规则全部启用/全部停用
	// 注意：此字段可能返回 null，表示取不到有效值。
		Enable *uint64 `json:"Enable,omitempty" name:"Enable"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchListsRequest struct {
	*tchttp.BaseRequest

	// 防火墙状态  0: 关闭，1：开启
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 资产类型 CVM/NAT/VPN/CLB/其它
	Type *string `json:"Type,omitempty" name:"Type"`

	// 地域 上海/重庆/广州，等等
	Area *string `json:"Area,omitempty" name:"Area"`

	// 搜索值  例子："{"common":"106.54.189.45"}"
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// 条数  默认值:10
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移值 默认值: 0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序，desc：降序，asc：升序
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段 PortTimes(风险端口数)
	By *string `json:"By,omitempty" name:"By"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeSwitchListsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchListsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总条数
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 列表数据
		Data []*SwitchListsData `json:"Data,omitempty" name:"Data" list`

		// 区域列表
		AreaLists []*string `json:"AreaLists,omitempty" name:"AreaLists" list`

		// 打开个数
	// 注意：此字段可能返回 null，表示取不到有效值。
		OnNum *uint64 `json:"OnNum,omitempty" name:"OnNum"`

		// 关闭个数
	// 注意：此字段可能返回 null，表示取不到有效值。
		OffNum *uint64 `json:"OffNum,omitempty" name:"OffNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSwitchListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncAssetStatusRequest struct {
	*tchttp.BaseRequest

	// 0: 互联网防火墙开关，1：vpc 防火墙开关
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSyncAssetStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	if len(f) > 0 {
		return errors.New("DescribeSyncAssetStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncAssetStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 1-更新中 2-更新完成 3、4-更新失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSyncAssetStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTableStatusRequest struct {
	*tchttp.BaseRequest

	// EdgeId值两个vpc间的边id vpc填Edgeid，不要填Area；
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// 状态值，0：检查表的状态 确实只有一个默认值
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Nat所在地域 NAT填Area，不要填Edgeid；
	Area *string `json:"Area,omitempty" name:"Area"`

	// 方向，0：出站，1：入站 默认值为 0
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("DescribeTableStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTableStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0：正常，其它：不正常
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcRuleOverviewRequest struct {
	*tchttp.BaseRequest

	// EdgeId值两个vpc间的边id  不是必填项可以为空，就是所有vpc间的访问控制规则
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcRuleOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EdgeId")
	if len(f) > 0 {
		return errors.New("DescribeVpcRuleOverviewRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcRuleOverviewResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 阻断策略规则数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		StrategyNum *uint64 `json:"StrategyNum,omitempty" name:"StrategyNum"`

		// 启用规则数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		StartRuleNum *uint64 `json:"StartRuleNum,omitempty" name:"StartRuleNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVpcRuleOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpandCfwVerticalRequest struct {
	*tchttp.BaseRequest

	// nat：nat防火墙，ew：东西向防火墙
	FwType *string `json:"FwType,omitempty" name:"FwType"`

	// 带宽值
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 防火墙实例id
	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("ExpandCfwVerticalRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ExpandCfwVerticalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExpandCfwVerticalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAcRuleRequest struct {
	*tchttp.BaseRequest

	// 规则数组
	Data []*RuleInfoData `json:"Data,omitempty" name:"Data" list`

	// EdgeId值
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// 访问规则状态
	Enable *int64 `json:"Enable,omitempty" name:"Enable"`

	// NAT地域
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("ModifyAcRuleRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态值，0:操作成功，非0：操作失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 返回多余的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Info *string `json:"Info,omitempty" name:"Info"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAcRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 状态，0：全部停用，1：全部启用
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Edge ID值
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// NAT地域
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("ModifyAllRuleStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0: 修改成功, 其他: 修改失败
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// 状态，0：关闭，1：开启
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 0: 互联网边界防火墙开关，1：vpc防火墙开关
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 选中的防火墙开关Id
	Ids []*string `json:"Ids,omitempty" name:"Ids" list`

	// NAT开关切换类型，1,单个子网，2，同开同关，3，全部
	ChangeType *int64 `json:"ChangeType,omitempty" name:"ChangeType"`

	// NAT实例所在地域
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllSwitchStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "Type")
	delete(f, "Ids")
	delete(f, "ChangeType")
	delete(f, "Area")
	if len(f) > 0 {
		return errors.New("ModifyAllSwitchStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAllSwitchStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改成功与否的状态值 0：修改成功，非 0：修改失败
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAllSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyItemSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// id值
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 状态值，0: 关闭 ,1:开启
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 0: 互联网边界边界防火墙开关，1：vpc防火墙开关
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyItemSwitchStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "Status")
	delete(f, "Type")
	if len(f) > 0 {
		return errors.New("ModifyItemSwitchStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyItemSwitchStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改成功与否状态值 0：修改成功，非 0：修改失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyItemSwitchStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupAllRuleStatusRequest struct {
	*tchttp.BaseRequest

	// 列表规则状态，0：全部停用，1：全部启用
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// Edge ID值
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// NAT地域, 腾讯云地域的英文简写
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupAllRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "Direction")
	delete(f, "EdgeId")
	delete(f, "Area")
	if len(f) > 0 {
		return errors.New("ModifySecurityGroupAllRuleStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySecurityGroupAllRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0: 修改成功, 其他: 修改失败
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityGroupAllRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifySequenceRulesRequest struct {
	*tchttp.BaseRequest

	// 边Id值
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// 修改数据
	Data []*SequenceData `json:"Data,omitempty" name:"Data" list`

	// NAT地域
	Area *string `json:"Area,omitempty" name:"Area"`

	// 方向，0：出向，1：入向
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("ModifySequenceRulesRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifySequenceRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0: 修改成功, 非0: 修改失败
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySequenceRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTableStatusRequest struct {
	*tchttp.BaseRequest

	// EdgeId值两个vpc间的边id
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// 状态值，1：锁表，2：解锁表
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Nat所在地域
	Area *string `json:"Area,omitempty" name:"Area"`

	// 0： 出向，1：入向
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("ModifyTableStatusRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTableStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0：正常，-1：不正常
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTableStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NatFwEipsInfo struct {

	// 弹性公网ip
	Eip *string `json:"Eip,omitempty" name:"Eip"`

	// 所属的Nat网关Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatGatewayId *string `json:"NatGatewayId,omitempty" name:"NatGatewayId"`

	// Nat网关名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NatGatewayName *string `json:"NatGatewayName,omitempty" name:"NatGatewayName"`
}

type RuleInfoData struct {

	// 执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// 访问源
	SourceIp *string `json:"SourceIp,omitempty" name:"SourceIp"`

	// 访问目的
	TargetIp *string `json:"TargetIp,omitempty" name:"TargetIp"`

	// 协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 策略
	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`

	// 访问源类型，1是IP，3是域名，4是IP地址模版，5是域名地址模版
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// 描述
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// 访问目的类型，1是IP，3是域名，4是IP地址模版，5是域名地址模版
	TargetType *uint64 `json:"TargetType,omitempty" name:"TargetType"`

	// 端口
	Port *string `json:"Port,omitempty" name:"Port"`

	// id值
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 日志id，从告警处创建必传，其它为空
	LogId *string `json:"LogId,omitempty" name:"LogId"`

	// 城市Code
	City *uint64 `json:"City,omitempty" name:"City"`

	// 国家Code
	Country *uint64 `json:"Country,omitempty" name:"Country"`

	// 云厂商，支持多个，以逗号分隔， 1:腾讯云（仅中国香港及海外）,2:阿里云,3:亚马逊云,4:华为云,5:微软云
	CloudCode *string `json:"CloudCode,omitempty" name:"CloudCode"`

	// 是否为地域
	IsRegion *uint64 `json:"IsRegion,omitempty" name:"IsRegion"`

	// 城市名
	CityName *string `json:"CityName,omitempty" name:"CityName"`

	// 国家名
	CountryName *string `json:"CountryName,omitempty" name:"CountryName"`
}

type RunSyncAssetRequest struct {
	*tchttp.BaseRequest

	// 0: 互联网防火墙开关，1：vpc 防火墙开关
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunSyncAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	if len(f) > 0 {
		return errors.New("RunSyncAssetRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RunSyncAssetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0：同步成功，1：资产更新中，2：后台同步调用失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunSyncAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupApiRuleData struct {

	// 访问源，入站时为Ip/Cidr，默认为0.0.0.0/0； 出站时当RuleType为1时，支持内网Ip/Cidr, 当RuleType为2时，填实例ID
	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`

	// 访问目的，出站时为Ip/Cidr，默认为0.0.0.0/0；入站时当RuleType为1时，支持内网Ip/Cidr, 当RuleType为2时，填实例ID
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// 协议，支持ANY/TCP/UDP/ICMP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 端口, 当Protocol为ANY或ICMP时，Port为-1/-1
	Port *string `json:"Port,omitempty" name:"Port"`

	// 策略, 1：阻断，2：放行
	Strategy *string `json:"Strategy,omitempty" name:"Strategy"`

	// 描述
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// 规则类型，1：VpcId+Ip/Cidr, 2: 实例ID，入站时为访问目的类型，出站时为访问源类型
	RuleType *uint64 `json:"RuleType,omitempty" name:"RuleType"`

	// 执行顺序，中间插入必传，前插、后插非必传
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// 私有网络ID，当RuleType为1时必传
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type SecurityGroupListData struct {

	// 规则ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// 访问源
	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`

	// 访问源类型，默认为0，1: VPC, 2: SUBNET, 3: CVM, 4: CLB, 5: ENI, 6: CDB
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 访问目的
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// 访问目的类型，默认为0，1: VPC, 2: SUBNET, 3: CVM, 4: CLB, 5: ENI, 6: CDB
	TargetType *uint64 `json:"TargetType,omitempty" name:"TargetType"`

	// 协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 目的端口
	Port *string `json:"Port,omitempty" name:"Port"`

	// 策略, 1：阻断，2：放行
	Strategy *uint64 `json:"Strategy,omitempty" name:"Strategy"`

	// 描述
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// 是否开关开启，0：未开启，1：开启
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 是否是正常规则，0：正常，1：异常
	IsNew *uint64 `json:"IsNew,omitempty" name:"IsNew"`

	// 单/双向下发，0:单向下发，1：双向下发
	BothWay *uint64 `json:"BothWay,omitempty" name:"BothWay"`

	// 私有网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 公网IP，多个以英文逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 内网IP，多个以英文逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`

	// 掩码地址，多个以英文逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cidr *string `json:"Cidr,omitempty" name:"Cidr"`
}

type SequenceData struct {

	// 规则Id值
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 修改前执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// 修改后执行顺序
	NewOrderIndex *uint64 `json:"NewOrderIndex,omitempty" name:"NewOrderIndex"`
}

type SetNatFwDnatRuleRequest struct {
	*tchttp.BaseRequest

	// 0：cfw新增模式，1：cfw接入模式。
	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`

	// 操作类型，可选值：add，del，modify。
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 防火墙实例id。
	CfwInstance *string `json:"CfwInstance,omitempty" name:"CfwInstance"`

	// 添加或删除操作的Dnat规则列表。
	AddOrDelDnatRules []*CfwNatDnatRule `json:"AddOrDelDnatRules,omitempty" name:"AddOrDelDnatRules" list`

	// 修改操作的原始Dnat规则
	OriginDnat *CfwNatDnatRule `json:"OriginDnat,omitempty" name:"OriginDnat"`

	// 修改操作的新的Dnat规则
	NewDnat *CfwNatDnatRule `json:"NewDnat,omitempty" name:"NewDnat"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
		return errors.New("SetNatFwDnatRuleRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SetNatFwDnatRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *%(obj)s) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetNatFwDnatRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchListsData struct {

	// 公网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIp *string `json:"PublicIp,omitempty" name:"PublicIp"`

	// 内网IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntranetIp *string `json:"IntranetIp,omitempty" name:"IntranetIp"`

	// 实例名
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 资产类型
	AssetType *string `json:"AssetType,omitempty" name:"AssetType"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Area *string `json:"Area,omitempty" name:"Area"`

	// 防火墙开关
	Switch *int64 `json:"Switch,omitempty" name:"Switch"`

	// id值
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 公网 IP 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpType *uint64 `json:"PublicIpType,omitempty" name:"PublicIpType"`

	// 风险端口数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortTimes *uint64 `json:"PortTimes,omitempty" name:"PortTimes"`

	// 最近扫描时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`

	// 扫描深度
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanMode *string `json:"ScanMode,omitempty" name:"ScanMode"`

	// 扫描状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanStatus *uint64 `json:"ScanStatus,omitempty" name:"ScanStatus"`
}
