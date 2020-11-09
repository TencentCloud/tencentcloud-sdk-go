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

type CreateAcRulesRequest struct {
	*tchttp.BaseRequest

	// 创建规则数据
	Data []*RuleInfoData `json:"Data,omitempty" name:"Data" list`

	// 0：添加，1：插入
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

func (r *CreateAcRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAcRulesRequest) FromJsonString(s string) error {
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

func (r *CreateAcRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAcRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAcRuleRequest struct {
	*tchttp.BaseRequest

	// id值
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 出站还是入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// EdgeId值
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// NAT地域
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DeleteAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAcRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAcRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态值
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 返回多余的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Info *string `json:"Info,omitempty" name:"Info"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAcRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAllAccessControlRuleRequest struct {
	*tchttp.BaseRequest

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// VPC间防火墙开关ID
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// nat地域
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DeleteAllAccessControlRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAllAccessControlRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAllAccessControlRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 状态值
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 返回多余信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Info *int64 `json:"Info,omitempty" name:"Info"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAllAccessControlRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAllAccessControlRuleResponse) FromJsonString(s string) error {
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

func (r *DescribeAcListsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAcListsRequest) FromJsonString(s string) error {
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

func (r *DescribeAcListsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAcListsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNatRuleOverviewRequest struct {
	*tchttp.BaseRequest

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// NAT地域
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *DescribeNatRuleOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNatRuleOverviewRequest) FromJsonString(s string) error {
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

		// 访问规则剩余条数
		RemainNum *int64 `json:"RemainNum,omitempty" name:"RemainNum"`

		// 阻断规则条数
		BlockNum *int64 `json:"BlockNum,omitempty" name:"BlockNum"`

		// 启用规则条数
		EnableNum *int64 `json:"EnableNum,omitempty" name:"EnableNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNatRuleOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNatRuleOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRuleOverviewRequest struct {
	*tchttp.BaseRequest

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *DescribeRuleOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRuleOverviewRequest) FromJsonString(s string) error {
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

func (r *DescribeRuleOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRuleOverviewResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSwitchListsRequest struct {
	*tchttp.BaseRequest

	// 防火墙状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 资产类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 地域
	Area *string `json:"Area,omitempty" name:"Area"`

	// 搜索值
	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`

	// 条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移值
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序，desc：降序，asc：升序
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序字段
	By *string `json:"By,omitempty" name:"By"`
}

func (r *DescribeSwitchListsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSwitchListsRequest) FromJsonString(s string) error {
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

func (r *DescribeSwitchListsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSwitchListsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncAssetStatusRequest struct {
	*tchttp.BaseRequest

	// 0: 互联网防火墙开关，1：vpc 防火墙开关
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeSyncAssetStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSyncAssetStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSyncAssetStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0：同步成功，1：资产更新中，2：后台同步调用失败
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSyncAssetStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSyncAssetStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTableStatusRequest struct {
	*tchttp.BaseRequest

	// EdgeId值
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// 状态值，0：检查表的状态
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Nat所在地域
	Area *string `json:"Area,omitempty" name:"Area"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *DescribeTableStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTableStatusRequest) FromJsonString(s string) error {
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

func (r *DescribeTableStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTableStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcRuleOverviewRequest struct {
	*tchttp.BaseRequest

	// 边id
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`
}

func (r *DescribeVpcRuleOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcRuleOverviewRequest) FromJsonString(s string) error {
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

func (r *DescribeVpcRuleOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeVpcRuleOverviewResponse) FromJsonString(s string) error {
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

func (r *ModifyAcRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAcRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAcRuleResponse struct {
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

func (r *ModifyAcRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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

func (r *ModifyAllRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAllRuleStatusRequest) FromJsonString(s string) error {
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

func (r *ModifyAllRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAllRuleStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAllSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// 状态，0：关闭，1：开启
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 0: 边界防火墙开关，1：vpc防火墙开关
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 选中的防火墙开关Id
	Ids []*string `json:"Ids,omitempty" name:"Ids" list`

	// NAT开关切换类型，1,单个子网，2，同开同关，3，全部
	ChangeType *int64 `json:"ChangeType,omitempty" name:"ChangeType"`

	// NAT实例所在地域
	Area *string `json:"Area,omitempty" name:"Area"`
}

func (r *ModifyAllSwitchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAllSwitchStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAllSwitchStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 开启或者关闭成功与否状态值
		Status *uint64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAllSwitchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAllSwitchStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyItemSwitchStatusRequest struct {
	*tchttp.BaseRequest

	// id值
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 状态值，0: 关闭 ,1:开启
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 0: 边界防火墙开关，1：vpc防火墙开关
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *ModifyItemSwitchStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyItemSwitchStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyItemSwitchStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 修改成功与否状态值
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyItemSwitchStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyItemSwitchStatusResponse) FromJsonString(s string) error {
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

	// 0：出向，1：入向
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *ModifySequenceRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySequenceRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySequenceRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 0: 修改成功, 其他: 修改失败
	// 注意：此字段可能返回 null，表示取不到有效值。
		Status *int64 `json:"Status,omitempty" name:"Status"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySequenceRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySequenceRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTableStatusRequest struct {
	*tchttp.BaseRequest

	// EdgeId值
	EdgeId *string `json:"EdgeId,omitempty" name:"EdgeId"`

	// 状态值，1：锁表，2：解锁表
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// Nat所在地域
	Area *string `json:"Area,omitempty" name:"Area"`

	// 0： 出向，1：入向
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`
}

func (r *ModifyTableStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTableStatusRequest) FromJsonString(s string) error {
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

func (r *ModifyTableStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTableStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

	// 描述
	Detail *string `json:"Detail,omitempty" name:"Detail"`

	// 方向，0：出站，1：入站
	Direction *uint64 `json:"Direction,omitempty" name:"Direction"`

	// 源类型,1是ip,2是域名,3是ip地址簿，4是ip组地址簿
	SourceType *uint64 `json:"SourceType,omitempty" name:"SourceType"`

	// 目的类型,1是ip,2是域名,3是ip地址簿，4是ip组地址簿
	TargetType *uint64 `json:"TargetType,omitempty" name:"TargetType"`

	// 端口
	Port *string `json:"Port,omitempty" name:"Port"`

	// id值
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// log
	LogId *string `json:"LogId,omitempty" name:"LogId"`

	// 城市Code
	City *uint64 `json:"City,omitempty" name:"City"`

	// 国家Code
	Country *uint64 `json:"Country,omitempty" name:"Country"`

	// 云厂商，支持多个，以逗号分隔， 1:腾讯云（仅海外）,2:阿里云,3:亚马逊云,4:华为云,5:微软云
	CloudCode *string `json:"CloudCode,omitempty" name:"CloudCode"`

	// 是否为地域
	IsRegion *uint64 `json:"IsRegion,omitempty" name:"IsRegion"`

	// 地域名
	CityName *string `json:"CityName,omitempty" name:"CityName"`

	// 地域名
	CountryName *string `json:"CountryName,omitempty" name:"CountryName"`
}

type RunSyncAssetRequest struct {
	*tchttp.BaseRequest

	// 0: 互联网防火墙开关，1：vpc 防火墙开关
	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

func (r *RunSyncAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunSyncAssetRequest) FromJsonString(s string) error {
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

func (r *RunSyncAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RunSyncAssetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SequenceData struct {

	// 规则Id值
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 修改前执行顺序
	OrderIndex *uint64 `json:"OrderIndex,omitempty" name:"OrderIndex"`

	// 修改后执行顺序
	NewOrderIndex *uint64 `json:"NewOrderIndex,omitempty" name:"NewOrderIndex"`
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
