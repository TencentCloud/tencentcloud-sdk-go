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

package v20180709

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ActionSummaryOverviewItem struct {
	// 交易类型编码
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// 交易类型：如包年包月新购、包年包月续费、按量计费扣费等类型
	ActionTypeName *string `json:"ActionTypeName,omitnil" name:"ActionTypeName"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil" name:"RealTotalCostRatio"`

	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil" name:"IncentivePayAmount"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil" name:"TransferPayAmount"`

	// 账单月份，格式2019-08
	BillMonth *string `json:"BillMonth,omitnil" name:"BillMonth"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitnil" name:"TotalCost"`
}

type ApplicableProducts struct {
	// 适用商品名称，值为“全产品通用”或商品名称组成的string，以","分割。
	GoodsName *string `json:"GoodsName,omitnil" name:"GoodsName"`

	// postPay后付费/prePay预付费/riPay预留实例/空字符串或者"*"表示全部模式。如GoodsName为多个商品名以","分割组成的string，而PayMode为"*"，表示每一件商品的模式都为"*"。
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`
}

type BillDetail struct {
	// 产品名称：用户所采购的各类云产品，例如：云服务器 CVM
	BusinessCodeName *string `json:"BusinessCodeName,omitnil" name:"BusinessCodeName"`

	// 子产品名称：用户采购的具体产品细分类型，例如：云服务器 CVM-标准型 S1
	ProductCodeName *string `json:"ProductCodeName,omitnil" name:"ProductCodeName"`

	// 计费模式：资源的计费模式，区分为包年包月和按量计费
	PayModeName *string `json:"PayModeName,omitnil" name:"PayModeName"`

	// 项目名称：资源归属的项目，用户在控制台给资源自主分配项目，未分配则是默认项目
	ProjectName *string `json:"ProjectName,omitnil" name:"ProjectName"`

	// 地域：资源所属地域，如华南地区（广州）
	RegionName *string `json:"RegionName,omitnil" name:"RegionName"`

	// 可用区：资源所属可用区，如广州三区
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// 资源 ID：账单中出账对象 ID，不同产品因资源形态不同，资源内容不完全相同，如云服务器 CVM 为对应的实例 ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 资源别名：用户在控制台为资源设置的名称，如果未设置，则默认为空
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`

	// 交易类型，如包年包月新购、包年包月续费、按量计费扣费等类型
	ActionTypeName *string `json:"ActionTypeName,omitnil" name:"ActionTypeName"`

	// 订单ID：包年包月计费模式下订购的订单号
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// 交易ID：结算扣费单号
	BillId *string `json:"BillId,omitnil" name:"BillId"`

	// 扣费时间：结算扣费时间
	PayTime *string `json:"PayTime,omitnil" name:"PayTime"`

	// 开始使用时间：产品服务开始使用时间
	FeeBeginTime *string `json:"FeeBeginTime,omitnil" name:"FeeBeginTime"`

	// 结束使用时间：产品服务结束使用时间
	FeeEndTime *string `json:"FeeEndTime,omitnil" name:"FeeEndTime"`

	// 组件列表
	ComponentSet []*BillDetailComponent `json:"ComponentSet,omitnil" name:"ComponentSet"`

	// 支付者UIN：支付者的账号 ID，账号 ID 是用户在腾讯云的唯一账号标识
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`

	// 使用者UIN：实际使用资源的账号 ID
	OwnerUin *string `json:"OwnerUin,omitnil" name:"OwnerUin"`

	// 操作者UIN：操作者账号 ID（预付费资源下单或后付费操作开通资源账号的 ID 或者角色 ID ）
	OperateUin *string `json:"OperateUin,omitnil" name:"OperateUin"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*BillTagInfo `json:"Tags,omitnil" name:"Tags"`

	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessCode *string `json:"BusinessCode,omitnil" name:"BusinessCode"`

	// 子产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitnil" name:"ProductCode"`

	// 交易类型编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *string `json:"RegionId,omitnil" name:"RegionId"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 价格属性：该组件除单价、时长外的其他影响折扣定价的属性信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PriceInfo []*string `json:"PriceInfo,omitnil" name:"PriceInfo"`

	// 关联交易单据ID：和本笔交易关联单据 ID，如，冲销订单，记录原订单、重结订单，退费单记录对应的原购买订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociatedOrder *BillDetailAssociatedOrder `json:"AssociatedOrder,omitnil" name:"AssociatedOrder"`

	// 计算说明：特殊交易类型计费结算的详细计算说明，如退费及变配
	// 注意：此字段可能返回 null，表示取不到有效值。
	Formula *string `json:"Formula,omitnil" name:"Formula"`

	// 计费规则：各产品详细的计费规则官网说明链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	FormulaUrl *string `json:"FormulaUrl,omitnil" name:"FormulaUrl"`

	// 账单归属日
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillDay *string `json:"BillDay,omitnil" name:"BillDay"`

	// 账单归属月
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillMonth *string `json:"BillMonth,omitnil" name:"BillMonth"`
}

type BillDetailAssociatedOrder struct {
	// 新购订单
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrepayPurchase *string `json:"PrepayPurchase,omitnil" name:"PrepayPurchase"`

	// 续费订单
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrepayRenew *string `json:"PrepayRenew,omitnil" name:"PrepayRenew"`

	// 升配订单
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrepayModifyUp *string `json:"PrepayModifyUp,omitnil" name:"PrepayModifyUp"`

	// 冲销订单
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReverseOrder *string `json:"ReverseOrder,omitnil" name:"ReverseOrder"`

	// 优惠调整后订单
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewOrder *string `json:"NewOrder,omitnil" name:"NewOrder"`

	// 优惠调整前订单
	// 注意：此字段可能返回 null，表示取不到有效值。
	Original *string `json:"Original,omitnil" name:"Original"`
}

type BillDetailComponent struct {
	// 组件类型：用户购买的产品或服务对应的组件大类，例如：云服务器 CVM 的组件：CPU、内存等
	ComponentCodeName *string `json:"ComponentCodeName,omitnil" name:"ComponentCodeName"`

	// 组件名称：用户购买的产品或服务，所包含的具体组件
	ItemCodeName *string `json:"ItemCodeName,omitnil" name:"ItemCodeName"`

	// 组件刊例价：组件的官网原始单价（如果客户享受一口价/合同价则默认不展示）
	SinglePrice *string `json:"SinglePrice,omitnil" name:"SinglePrice"`

	// 组件指定价（已废弃）
	//
	// Deprecated: SpecifiedPrice is deprecated.
	SpecifiedPrice *string `json:"SpecifiedPrice,omitnil" name:"SpecifiedPrice"`

	// 组件价格单位：组件价格的单位，单位构成：元/用量单位/时长单位
	PriceUnit *string `json:"PriceUnit,omitnil" name:"PriceUnit"`

	// 组件用量：该组件实际结算用量，组件用量 = 组件原始用量 - 抵扣用量（含资源包
	UsedAmount *string `json:"UsedAmount,omitnil" name:"UsedAmount"`

	// 组件用量单位：组件用量对应的单位
	UsedAmountUnit *string `json:"UsedAmountUnit,omitnil" name:"UsedAmountUnit"`

	// 原始用量/时长：组件被资源包抵扣前的原始用量/时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalMeasure *string `json:"RealTotalMeasure,omitnil" name:"RealTotalMeasure"`

	// 抵扣用量/时长（含资源包）：组件被资源包抵扣的用量/时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeductedMeasure *string `json:"DeductedMeasure,omitnil" name:"DeductedMeasure"`

	// 使用时长：资源使用的时长
	TimeSpan *string `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 时长单位：资源使用时长的单位
	TimeUnitName *string `json:"TimeUnitName,omitnil" name:"TimeUnitName"`

	// 组件原价：原价 = 组件刊例价 * 组件用量 * 使用时长（如果客户享受一口价/合同价则默认不展示，退费类场景也默认不展示）
	Cost *string `json:"Cost,omitnil" name:"Cost"`

	// 折扣率：本资源享受的折扣率（如果客户享受一口价/合同价则默认不展示，退费场景也默认不展示）
	Discount *string `json:"Discount,omitnil" name:"Discount"`

	// 优惠类型
	ReduceType *string `json:"ReduceType,omitnil" name:"ReduceType"`

	// 优惠后总价：优惠后总价=（原价 - 预留实例抵扣原价 - 节省计划抵扣原价）* 折扣率
	RealCost *string `json:"RealCost,omitnil" name:"RealCost"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil" name:"IncentivePayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil" name:"TransferPayAmount"`

	// 组件类型编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemCode *string `json:"ItemCode,omitnil" name:"ItemCode"`

	// 组件名称编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentCode *string `json:"ComponentCode,omitnil" name:"ComponentCode"`

	// 组件单价：组件的折后单价，组件单价 = 刊例价 * 折扣
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractPrice *string `json:"ContractPrice,omitnil" name:"ContractPrice"`

	// 实例类型：购买的产品服务对应的实例类型，包括资源包、RI、SP、竞价实例。正常的实例展示默认为不展示
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 预留实例抵扣的使用时长：本产品或服务使用预留实例抵扣的使用时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiTimeSpan *string `json:"RiTimeSpan,omitnil" name:"RiTimeSpan"`

	// 预留实例抵扣组件原价：本产品或服务使用预留实例抵扣的组件原价金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalCostWithRI *string `json:"OriginalCostWithRI,omitnil" name:"OriginalCostWithRI"`

	// 节省计划抵扣率：节省计划可用余额额度范围内，节省计划对于此组件打的折扣率
	// 注意：此字段可能返回 null，表示取不到有效值。
	SPDeductionRate *string `json:"SPDeductionRate,omitnil" name:"SPDeductionRate"`

	// 节省计划抵扣金额（已废弃）
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: SPDeduction is deprecated.
	SPDeduction *string `json:"SPDeduction,omitnil" name:"SPDeduction"`

	// 节省计划抵扣组件原价：节省计划抵扣原价=节省计划包抵扣金额/节省计划抵扣率
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalCostWithSP *string `json:"OriginalCostWithSP,omitnil" name:"OriginalCostWithSP"`

	// 混合折扣率：综合各类折扣抵扣信息后的最终折扣率，混合折扣率 = 优惠后总价 / 组件原价
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlendedDiscount *string `json:"BlendedDiscount,omitnil" name:"BlendedDiscount"`

	// 配置描述：资源配置规格信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentConfig []*BillDetailComponentConfig `json:"ComponentConfig,omitnil" name:"ComponentConfig"`
}

type BillDetailComponentConfig struct {
	// 配置描述名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 配置描述值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type BillDistributionResourceSummary struct {
	// 产品名称：用户所采购的各类云产品，例如：云服务器 CVM
	BusinessCodeName *string `json:"BusinessCodeName,omitnil" name:"BusinessCodeName"`

	// 子产品名称：用户采购的具体产品细分类型，例如：云服务器 CVM-标准型 S1
	ProductCodeName *string `json:"ProductCodeName,omitnil" name:"ProductCodeName"`

	// 计费模式：资源的计费模式，区分为包年包月和按量计费
	PayModeName *string `json:"PayModeName,omitnil" name:"PayModeName"`

	// 项目名称：资源归属的项目，用户在控制台给资源自主分配项目，未分配则是默认项目
	ProjectName *string `json:"ProjectName,omitnil" name:"ProjectName"`

	// 地域：资源所属地域，如华南地区（广州）
	RegionName *string `json:"RegionName,omitnil" name:"RegionName"`

	// 可用区：资源所属可用区，如广州三区
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// 资源 ID：账单中出账对象 ID，不同产品因资源形态不同，资源内容不完全相同，如云服务器 CVM 为对应的实例 ID	
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 资源别名：用户在控制台为资源设置的名称，如果未设置，则默认为空
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`

	// 交易类型：如包年包月新购、包年包月续费、按量计费扣费等类型
	ActionTypeName *string `json:"ActionTypeName,omitnil" name:"ActionTypeName"`

	// 订单ID：包年包月计费模式下订购的订单号
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// 扣费时间：结算扣费时间
	PayTime *string `json:"PayTime,omitnil" name:"PayTime"`

	// 开始使用时间：产品服务开始使用时间
	FeeBeginTime *string `json:"FeeBeginTime,omitnil" name:"FeeBeginTime"`

	// 结束使用时间：产品服务结束使用时间
	FeeEndTime *string `json:"FeeEndTime,omitnil" name:"FeeEndTime"`

	// 配置描述：该资源下的计费项名称和用量合并展示，仅在资源账单体现
	ConfigDesc *string `json:"ConfigDesc,omitnil" name:"ConfigDesc"`

	// 扩展字段1：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField1 *string `json:"ExtendField1,omitnil" name:"ExtendField1"`

	// 扩展字段2：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField2 *string `json:"ExtendField2,omitnil" name:"ExtendField2"`

	// 原价：原价 = 组件刊例价 * 组件用量 * 使用时长（如果客户享受一口价/合同价则默认不展示，退费类场景也默认不展示）
	TotalCost *string `json:"TotalCost,omitnil" name:"TotalCost"`

	// 折扣率：本资源享受的折扣率（如果客户享受一口价/合同价则默认不展示，退费场景也默认不展示）
	Discount *string `json:"Discount,omitnil" name:"Discount"`

	// 优惠类型
	ReduceType *string `json:"ReduceType,omitnil" name:"ReduceType"`

	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil" name:"IncentivePayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil" name:"TransferPayAmount"`

	// 扩展字段3：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField3 *string `json:"ExtendField3,omitnil" name:"ExtendField3"`

	// 扩展字段4：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField4 *string `json:"ExtendField4,omitnil" name:"ExtendField4"`

	// 扩展字段5：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField5 *string `json:"ExtendField5,omitnil" name:"ExtendField5"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*BillTagInfo `json:"Tags,omitnil" name:"Tags"`

	// 使用者UIN：实际使用资源的账号 ID
	OwnerUin *string `json:"OwnerUin,omitnil" name:"OwnerUin"`

	// 操作者UIN：操作者账号 ID（预付费资源下单或后付费操作开通资源账号的 ID 或者角色 ID ）
	OperateUin *string `json:"OperateUin,omitnil" name:"OperateUin"`

	// 产品编码
	BusinessCode *string `json:"BusinessCode,omitnil" name:"BusinessCode"`

	// 子产品编码
	ProductCode *string `json:"ProductCode,omitnil" name:"ProductCode"`

	// 地域ID
	RegionId *int64 `json:"RegionId,omitnil" name:"RegionId"`

	// 实例类型：购买的产品服务对应的实例类型，包括资源包、RI、SP、竞价实例。正常的实例展示默认为不展示
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 预留实例抵扣组件原价：本产品或服务使用预留实例抵扣的组件原价金额	
	OriginalCostWithRI *string `json:"OriginalCostWithRI,omitnil" name:"OriginalCostWithRI"`

	// 节省计划抵扣金额（已废弃）
	//
	// Deprecated: SPDeduction is deprecated.
	SPDeduction *string `json:"SPDeduction,omitnil" name:"SPDeduction"`

	// 节省计划抵扣组件原价：节省计划抵扣原价=节省计划包抵扣金额/节省计划抵扣率	
	OriginalCostWithSP *string `json:"OriginalCostWithSP,omitnil" name:"OriginalCostWithSP"`

	// 账单归属月
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillMonth *string `json:"BillMonth,omitnil" name:"BillMonth"`
}

type BillResourceSummary struct {
	// 产品名称：用户所采购的各类云产品，例如：云服务器 CVM
	BusinessCodeName *string `json:"BusinessCodeName,omitnil" name:"BusinessCodeName"`

	// 子产品名称：用户采购的具体产品细分类型，例如：云服务器 CVM-标准型 S1
	ProductCodeName *string `json:"ProductCodeName,omitnil" name:"ProductCodeName"`

	// 计费模式：资源的计费模式，区分为包年包月和按量计费
	PayModeName *string `json:"PayModeName,omitnil" name:"PayModeName"`

	// 项目名称：资源归属的项目，用户在控制台给资源自主分配项目，未分配则是默认项目
	ProjectName *string `json:"ProjectName,omitnil" name:"ProjectName"`

	// 地域：资源所属地域，如华南地区（广州）
	RegionName *string `json:"RegionName,omitnil" name:"RegionName"`

	// 可用区：资源所属可用区，如广州三区
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// 资源 ID：账单中出账对象 ID，不同产品因资源形态不同，资源内容不完全相同，如云服务器 CVM 为对应的实例 ID	
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 资源别名：用户在控制台为资源设置的名称，如果未设置，则默认为空
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`

	// 交易类型：如包年包月新购、包年包月续费、按量计费扣费等类型
	ActionTypeName *string `json:"ActionTypeName,omitnil" name:"ActionTypeName"`

	// 订单ID：包年包月计费模式下订购的订单号
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// 扣费时间：结算扣费时间
	PayTime *string `json:"PayTime,omitnil" name:"PayTime"`

	// 开始使用时间：产品服务开始使用时间
	FeeBeginTime *string `json:"FeeBeginTime,omitnil" name:"FeeBeginTime"`

	// 结束使用时间：产品服务结束使用时间
	FeeEndTime *string `json:"FeeEndTime,omitnil" name:"FeeEndTime"`

	// 配置描述：该资源下的计费项名称和用量合并展示，仅在资源账单体现
	ConfigDesc *string `json:"ConfigDesc,omitnil" name:"ConfigDesc"`

	// 扩展字段1：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField1 *string `json:"ExtendField1,omitnil" name:"ExtendField1"`

	// 扩展字段2：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField2 *string `json:"ExtendField2,omitnil" name:"ExtendField2"`

	// 原价：原价 = 组件刊例价 * 组件用量 * 使用时长（如果客户享受一口价/合同价则默认不展示，退费类场景也默认不展示）
	TotalCost *string `json:"TotalCost,omitnil" name:"TotalCost"`

	// 折扣率：本资源享受的折扣率（如果客户享受一口价/合同价则默认不展示，退费场景也默认不展示）
	Discount *string `json:"Discount,omitnil" name:"Discount"`

	// 优惠类型
	ReduceType *string `json:"ReduceType,omitnil" name:"ReduceType"`

	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil" name:"IncentivePayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil" name:"TransferPayAmount"`

	// 扩展字段3：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField3 *string `json:"ExtendField3,omitnil" name:"ExtendField3"`

	// 扩展字段4：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField4 *string `json:"ExtendField4,omitnil" name:"ExtendField4"`

	// 扩展字段5：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField5 *string `json:"ExtendField5,omitnil" name:"ExtendField5"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*BillTagInfo `json:"Tags,omitnil" name:"Tags"`

	// 支付者UIN：支付者的账号 ID，账号 ID 是用户在腾讯云的唯一账号标识
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`

	// 使用者UIN：实际使用资源的账号 ID
	OwnerUin *string `json:"OwnerUin,omitnil" name:"OwnerUin"`

	// 操作者UIN：操作者账号 ID（预付费资源下单或后付费操作开通资源账号的 ID 或者角色 ID ）
	OperateUin *string `json:"OperateUin,omitnil" name:"OperateUin"`

	// 产品编码
	BusinessCode *string `json:"BusinessCode,omitnil" name:"BusinessCode"`

	// 子产品编码
	ProductCode *string `json:"ProductCode,omitnil" name:"ProductCode"`

	// 地域ID
	RegionId *int64 `json:"RegionId,omitnil" name:"RegionId"`

	// 实例类型：购买的产品服务对应的实例类型，包括资源包、RI、SP、竞价实例。正常的实例展示默认为不展示
	InstanceType *string `json:"InstanceType,omitnil" name:"InstanceType"`

	// 预留实例抵扣组件原价：本产品或服务使用预留实例抵扣的组件原价金额	
	OriginalCostWithRI *string `json:"OriginalCostWithRI,omitnil" name:"OriginalCostWithRI"`

	// 节省计划抵扣金额（已废弃）
	//
	// Deprecated: SPDeduction is deprecated.
	SPDeduction *string `json:"SPDeduction,omitnil" name:"SPDeduction"`

	// 节省计划抵扣组件原价：节省计划抵扣原价=节省计划包抵扣金额/节省计划抵扣率	
	OriginalCostWithSP *string `json:"OriginalCostWithSP,omitnil" name:"OriginalCostWithSP"`

	// 账单归属月
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillMonth *string `json:"BillMonth,omitnil" name:"BillMonth"`
}

type BillTagInfo struct {
	// 分账标签键
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type BillTransactionInfo struct {
	// 收支类型：deduct 扣费, recharge 充值, return 退费， block 冻结, unblock 解冻
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// 流水金额，单位（分）；正数表示入账，负数表示出账
	Amount *int64 `json:"Amount,omitnil" name:"Amount"`

	// 可用余额，单位（分）；正数表示入账，负数表示出账
	Balance *int64 `json:"Balance,omitnil" name:"Balance"`

	// 流水号，如20190131020000236005203583326401
	BillId *string `json:"BillId,omitnil" name:"BillId"`

	// 描述信息
	OperationInfo *string `json:"OperationInfo,omitnil" name:"OperationInfo"`

	// 操作时间"2019-01-31 23:35:10.000"
	OperationTime *string `json:"OperationTime,omitnil" name:"OperationTime"`

	// 现金账户余额，单位（分）
	Cash *int64 `json:"Cash,omitnil" name:"Cash"`

	// 赠送金余额，单位（分）
	Incentive *int64 `json:"Incentive,omitnil" name:"Incentive"`

	// 冻结余额，单位（分）
	Freezing *int64 `json:"Freezing,omitnil" name:"Freezing"`

	// 交易渠道
	PayChannel *string `json:"PayChannel,omitnil" name:"PayChannel"`

	// 扣费模式：trade 包年包月(预付费)，hourh  按量-小时结，hourd 按量-日结，hourm 按量-月结，month 按量-月结
	DeductMode *string `json:"DeductMode,omitnil" name:"DeductMode"`
}

type BusinessSummaryInfo struct {
	// 产品编码
	BusinessCode *string `json:"BusinessCode,omitnil" name:"BusinessCode"`

	// 产品名称：用户所采购的各类云产品，例如：云服务器 CVM
	BusinessCodeName *string `json:"BusinessCodeName,omitnil" name:"BusinessCodeName"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCost *string `json:"TotalCost,omitnil" name:"TotalCost"`

	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil" name:"IncentivePayAmount"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil" name:"TransferPayAmount"`
}

type BusinessSummaryOverviewItem struct {
	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessCode *string `json:"BusinessCode,omitnil" name:"BusinessCode"`

	// 产品名称：用户所采购的各类云产品，例如：云服务器 CVM
	BusinessCodeName *string `json:"BusinessCodeName,omitnil" name:"BusinessCodeName"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil" name:"RealTotalCostRatio"`

	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil" name:"IncentivePayAmount"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	TransferPayAmount *string `json:"TransferPayAmount,omitnil" name:"TransferPayAmount"`

	// 账单月份，格式2019-08
	BillMonth *string `json:"BillMonth,omitnil" name:"BillMonth"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitnil" name:"TotalCost"`
}

type BusinessSummaryTotal struct {
	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil" name:"IncentivePayAmount"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil" name:"CashPayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	TransferPayAmount *string `json:"TransferPayAmount,omitnil" name:"TransferPayAmount"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitnil" name:"TotalCost"`
}

type ConditionBusiness struct {
	// 产品名称代码
	BusinessCode *string `json:"BusinessCode,omitnil" name:"BusinessCode"`

	// 产品名称
	BusinessCodeName *string `json:"BusinessCodeName,omitnil" name:"BusinessCodeName"`
}

type ConditionPayMode struct {
	// 付费模式
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 付费模式名称
	PayModeName *string `json:"PayModeName,omitnil" name:"PayModeName"`
}

type ConditionProject struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil" name:"ProjectName"`
}

type ConditionRegion struct {
	// 地域ID
	RegionId *string `json:"RegionId,omitnil" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitnil" name:"RegionName"`
}

type Conditions struct {
	// 只支持6和12两个值
	TimeRange *uint64 `json:"TimeRange,omitnil" name:"TimeRange"`

	// 产品名称代码
	BusinessCode *string `json:"BusinessCode,omitnil" name:"BusinessCode"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 地域ID
	RegionId *int64 `json:"RegionId,omitnil" name:"RegionId"`

	// 付费模式，可选prePay和postPay
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 资源关键字
	ResourceKeyword *string `json:"ResourceKeyword,omitnil" name:"ResourceKeyword"`

	// 产品名称代码
	BusinessCodes []*string `json:"BusinessCodes,omitnil" name:"BusinessCodes"`

	// 子产品名称代码
	ProductCodes []*string `json:"ProductCodes,omitnil" name:"ProductCodes"`

	// 地域ID
	RegionIds []*int64 `json:"RegionIds,omitnil" name:"RegionIds"`

	// 项目ID
	ProjectIds []*int64 `json:"ProjectIds,omitnil" name:"ProjectIds"`

	// 付费模式，可选prePay和postPay
	PayModes []*string `json:"PayModes,omitnil" name:"PayModes"`

	// 交易类型
	ActionTypes []*string `json:"ActionTypes,omitnil" name:"ActionTypes"`

	// 是否隐藏0元流水
	HideFreeCost *int64 `json:"HideFreeCost,omitnil" name:"HideFreeCost"`

	// 排序规则，可选desc和asc
	OrderByCost *string `json:"OrderByCost,omitnil" name:"OrderByCost"`

	// 交易ID
	BillIds []*string `json:"BillIds,omitnil" name:"BillIds"`

	// 组件编码
	ComponentCodes []*string `json:"ComponentCodes,omitnil" name:"ComponentCodes"`

	// 文件ID
	FileIds []*string `json:"FileIds,omitnil" name:"FileIds"`

	// 文件类型
	FileTypes []*string `json:"FileTypes,omitnil" name:"FileTypes"`

	// 状态
	Status []*uint64 `json:"Status,omitnil" name:"Status"`
}

type ConsumptionBusinessSummaryDataItem struct {
	// 产品名称代码
	BusinessCode *string `json:"BusinessCode,omitnil" name:"BusinessCode"`

	// 产品名称
	BusinessCodeName *string `json:"BusinessCodeName,omitnil" name:"BusinessCodeName"`

	// 折后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// 费用趋势
	Trend *ConsumptionSummaryTrend `json:"Trend,omitnil" name:"Trend"`

	// 现金
	// 注意：此字段可能返回 null，表示取不到有效值。
	CashPayAmount *string `json:"CashPayAmount,omitnil" name:"CashPayAmount"`

	// 赠送金
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil" name:"IncentivePayAmount"`

	// 代金券
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// 分成金
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil" name:"TransferPayAmount"`
}

type ConsumptionProjectSummaryDataItem struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil" name:"ProjectName"`

	// 折后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// 趋势
	Trend *ConsumptionSummaryTrend `json:"Trend,omitnil" name:"Trend"`

	// 产品消耗详情
	Business []*ConsumptionBusinessSummaryDataItem `json:"Business,omitnil" name:"Business"`

	// 现金
	// 注意：此字段可能返回 null，表示取不到有效值。
	CashPayAmount *string `json:"CashPayAmount,omitnil" name:"CashPayAmount"`

	// 赠送金
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil" name:"IncentivePayAmount"`

	// 代金券
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// 分成金
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil" name:"TransferPayAmount"`
}

type ConsumptionRegionSummaryDataItem struct {
	// 地域ID
	RegionId *string `json:"RegionId,omitnil" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitnil" name:"RegionName"`

	// 折后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// 趋势
	Trend *ConsumptionSummaryTrend `json:"Trend,omitnil" name:"Trend"`

	// 产品消费详情
	Business []*ConsumptionBusinessSummaryDataItem `json:"Business,omitnil" name:"Business"`
}

type ConsumptionResourceSummaryConditionValue struct {
	// 产品列表
	Business []*ConditionBusiness `json:"Business,omitnil" name:"Business"`

	// 项目列表
	Project []*ConditionProject `json:"Project,omitnil" name:"Project"`

	// 地域列表
	Region []*ConditionRegion `json:"Region,omitnil" name:"Region"`

	// 付费模式列表
	PayMode []*ConditionPayMode `json:"PayMode,omitnil" name:"PayMode"`
}

type ConsumptionResourceSummaryDataItem struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 资源名称
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`

	// 折后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// 现金花费
	CashPayAmount *string `json:"CashPayAmount,omitnil" name:"CashPayAmount"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil" name:"ProjectName"`

	// 地域ID
	RegionId *string `json:"RegionId,omitnil" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitnil" name:"RegionName"`

	// 付费模式
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 付费模式名称
	PayModeName *string `json:"PayModeName,omitnil" name:"PayModeName"`

	// 产品名称代码
	BusinessCode *string `json:"BusinessCode,omitnil" name:"BusinessCode"`

	// 产品名称
	BusinessCodeName *string `json:"BusinessCodeName,omitnil" name:"BusinessCodeName"`

	// 消耗类型
	ConsumptionTypeName *string `json:"ConsumptionTypeName,omitnil" name:"ConsumptionTypeName"`

	// 折前价
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealCost *string `json:"RealCost,omitnil" name:"RealCost"`

	// 费用起始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeeBeginTime *string `json:"FeeBeginTime,omitnil" name:"FeeBeginTime"`

	// 费用结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeeEndTime *string `json:"FeeEndTime,omitnil" name:"FeeEndTime"`

	// 天数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DayDiff *string `json:"DayDiff,omitnil" name:"DayDiff"`

	// 每日消耗
	// 注意：此字段可能返回 null，表示取不到有效值。
	DailyTotalCost *string `json:"DailyTotalCost,omitnil" name:"DailyTotalCost"`

	// 订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// 代金券
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// 赠送金
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil" name:"IncentivePayAmount"`

	// 分成金
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil" name:"TransferPayAmount"`
}

type ConsumptionSummaryTotal struct {
	// 折后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil" name:"RealTotalCost"`
}

type ConsumptionSummaryTrend struct {
	// 趋势类型，upward上升/downward下降/none无
	Type *string `json:"Type,omitnil" name:"Type"`

	// 趋势值，Type为none是该字段值为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type CosDetailSets struct {
	// 存储桶名称
	BucketName *string `json:"BucketName,omitnil" name:"BucketName"`

	// 用量开始时间
	DosageBeginTime *string `json:"DosageBeginTime,omitnil" name:"DosageBeginTime"`

	// 用量结束时间
	DosageEndTime *string `json:"DosageEndTime,omitnil" name:"DosageEndTime"`

	// 子产品名称
	SubProductCodeName *string `json:"SubProductCodeName,omitnil" name:"SubProductCodeName"`

	// 计费项名称
	BillingItemCodeName *string `json:"BillingItemCodeName,omitnil" name:"BillingItemCodeName"`

	// 用量
	DosageValue *string `json:"DosageValue,omitnil" name:"DosageValue"`

	// 单位
	Unit *string `json:"Unit,omitnil" name:"Unit"`
}

type CostComponentSet struct {
	// 组件类型名称
	ComponentCodeName *string `json:"ComponentCodeName,omitnil" name:"ComponentCodeName"`

	// 组件名称
	ItemCodeName *string `json:"ItemCodeName,omitnil" name:"ItemCodeName"`

	// 刊例价
	SinglePrice *string `json:"SinglePrice,omitnil" name:"SinglePrice"`

	// 刊例价单位
	PriceUnit *string `json:"PriceUnit,omitnil" name:"PriceUnit"`

	// 用量
	UsedAmount *string `json:"UsedAmount,omitnil" name:"UsedAmount"`

	// 用量单位
	UsedAmountUnit *string `json:"UsedAmountUnit,omitnil" name:"UsedAmountUnit"`

	// 原价
	Cost *string `json:"Cost,omitnil" name:"Cost"`

	// 折扣
	Discount *string `json:"Discount,omitnil" name:"Discount"`

	// 折后价
	RealCost *string `json:"RealCost,omitnil" name:"RealCost"`

	// 代金券支付金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// 现金支付金额
	CashPayAmount *string `json:"CashPayAmount,omitnil" name:"CashPayAmount"`

	// 赠送金支付金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil" name:"IncentivePayAmount"`
}

type CostDetail struct {
	// 支付者uin
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`

	// 产品名称
	BusinessCodeName *string `json:"BusinessCodeName,omitnil" name:"BusinessCodeName"`

	// 子产品名称
	ProductCodeName *string `json:"ProductCodeName,omitnil" name:"ProductCodeName"`

	// 计费模式名称
	PayModeName *string `json:"PayModeName,omitnil" name:"PayModeName"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil" name:"ProjectName"`

	// 区域名称
	RegionName *string `json:"RegionName,omitnil" name:"RegionName"`

	// 地区名称
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 资源名称
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`

	// 类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionTypeName *string `json:"ActionTypeName,omitnil" name:"ActionTypeName"`

	// 订单id
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// 交易id
	BillId *string `json:"BillId,omitnil" name:"BillId"`

	// 费用开始时间
	FeeBeginTime *string `json:"FeeBeginTime,omitnil" name:"FeeBeginTime"`

	// 费用结束时间
	FeeEndTime *string `json:"FeeEndTime,omitnil" name:"FeeEndTime"`

	// 组件明细
	ComponentSet []*CostComponentSet `json:"ComponentSet,omitnil" name:"ComponentSet"`

	// 子产品名称代码
	ProductCode *string `json:"ProductCode,omitnil" name:"ProductCode"`
}

// Predefined struct for user
type CreateAllocationTagRequestParams struct {
	// 用户分账标签键
	TagKey []*string `json:"TagKey,omitnil" name:"TagKey"`
}

type CreateAllocationTagRequest struct {
	*tchttp.BaseRequest
	
	// 用户分账标签键
	TagKey []*string `json:"TagKey,omitnil" name:"TagKey"`
}

func (r *CreateAllocationTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAllocationTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAllocationTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAllocationTagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAllocationTagResponse struct {
	*tchttp.BaseResponse
	Response *CreateAllocationTagResponseParams `json:"Response"`
}

func (r *CreateAllocationTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAllocationTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSavingPlanOrderRequestParams struct {
	// 地域编码
	RegionId *int64 `json:"RegionId,omitnil" name:"RegionId"`

	// 区域编码
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// 预付费类型
	PrePayType *string `json:"PrePayType,omitnil" name:"PrePayType"`

	// 时长
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 时长单位
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 商品唯一标识
	CommodityCode *string `json:"CommodityCode,omitnil" name:"CommodityCode"`

	// 承诺时长内的小额金额（单位：元）
	PromiseUseAmount *uint64 `json:"PromiseUseAmount,omitnil" name:"PromiseUseAmount"`

	// 节省计划的指定生效时间，若不传则为当前下单时间。传参数格式:"2023-10-01 00:00:00"，仅支持指定日期的0点时刻
	SpecifyEffectTime *string `json:"SpecifyEffectTime,omitnil" name:"SpecifyEffectTime"`

	// 可重入ID
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`
}

type CreateSavingPlanOrderRequest struct {
	*tchttp.BaseRequest
	
	// 地域编码
	RegionId *int64 `json:"RegionId,omitnil" name:"RegionId"`

	// 区域编码
	ZoneId *int64 `json:"ZoneId,omitnil" name:"ZoneId"`

	// 预付费类型
	PrePayType *string `json:"PrePayType,omitnil" name:"PrePayType"`

	// 时长
	TimeSpan *uint64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 时长单位
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 商品唯一标识
	CommodityCode *string `json:"CommodityCode,omitnil" name:"CommodityCode"`

	// 承诺时长内的小额金额（单位：元）
	PromiseUseAmount *uint64 `json:"PromiseUseAmount,omitnil" name:"PromiseUseAmount"`

	// 节省计划的指定生效时间，若不传则为当前下单时间。传参数格式:"2023-10-01 00:00:00"，仅支持指定日期的0点时刻
	SpecifyEffectTime *string `json:"SpecifyEffectTime,omitnil" name:"SpecifyEffectTime"`

	// 可重入ID
	ClientToken *string `json:"ClientToken,omitnil" name:"ClientToken"`
}

func (r *CreateSavingPlanOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSavingPlanOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RegionId")
	delete(f, "ZoneId")
	delete(f, "PrePayType")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "CommodityCode")
	delete(f, "PromiseUseAmount")
	delete(f, "SpecifyEffectTime")
	delete(f, "ClientToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSavingPlanOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSavingPlanOrderResponseParams struct {
	// 订单号
	BigDealId *string `json:"BigDealId,omitnil" name:"BigDealId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSavingPlanOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateSavingPlanOrderResponseParams `json:"Response"`
}

func (r *CreateSavingPlanOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSavingPlanOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Deal struct {
	// 订单号
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// 订单的状态 1：未支付 2：已支付3：发货中 4：已发货 5：发货失败 6：已退款 7：已关单 8：订单过期 9：订单已失效 10：产品已失效 11：代付拒绝 12：支付中
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 支付者
	Payer *string `json:"Payer,omitnil" name:"Payer"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 创建人
	Creator *string `json:"Creator,omitnil" name:"Creator"`

	// 实际支付金额（分）
	RealTotalCost *int64 `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// 代金券抵扣金额（分）
	VoucherDecline *int64 `json:"VoucherDecline,omitnil" name:"VoucherDecline"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 产品分类ID
	GoodsCategoryId *int64 `json:"GoodsCategoryId,omitnil" name:"GoodsCategoryId"`

	// 产品详情
	ProductInfo []*ProductInfo `json:"ProductInfo,omitnil" name:"ProductInfo"`

	// 时长
	TimeSpan *float64 `json:"TimeSpan,omitnil" name:"TimeSpan"`

	// 时间单位
	TimeUnit *string `json:"TimeUnit,omitnil" name:"TimeUnit"`

	// 货币单位
	Currency *string `json:"Currency,omitnil" name:"Currency"`

	// 折扣率
	Policy *float64 `json:"Policy,omitnil" name:"Policy"`

	// 单价（分）
	Price *float64 `json:"Price,omitnil" name:"Price"`

	// 原价（分）
	TotalCost *float64 `json:"TotalCost,omitnil" name:"TotalCost"`

	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitnil" name:"ProductCode"`

	// 子产品编码
	SubProductCode *string `json:"SubProductCode,omitnil" name:"SubProductCode"`

	// 大订单号
	BigDealId *string `json:"BigDealId,omitnil" name:"BigDealId"`

	// 退费公式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Formula *string `json:"Formula,omitnil" name:"Formula"`

	// 退费涉及订单信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefReturnDeals *string `json:"RefReturnDeals,omitnil" name:"RefReturnDeals"`

	// 付费模式：prePay 预付费 postPay后付费 riPay预留实例
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 交易类型
	// modifyNetworkMode 调整带宽模式
	// modifyNetworkSize 调整带宽大小
	// refund 退款
	// downgrade 降配
	// upgrade 升配
	// renew 续费
	// purchase 购买
	// preMoveOut 包年包月迁出资源
	// preMoveIn 包年包月迁入资源
	// preToPost 预付费转后付费
	// postMoveOut 按量计费迁出资源
	// postMoveIn 按量计费迁入资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *string `json:"Action,omitnil" name:"Action"`

	// 产品编码中文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitnil" name:"ProductName"`

	// 子产品编码中文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductName *string `json:"SubProductName,omitnil" name:"SubProductName"`

	// 订单对应的资源id, 查询参数Limit超过200，将返回null
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId []*string `json:"ResourceId,omitnil" name:"ResourceId"`
}

// Predefined struct for user
type DeleteAllocationTagRequestParams struct {
	// 用户分账标签键
	TagKey []*string `json:"TagKey,omitnil" name:"TagKey"`
}

type DeleteAllocationTagRequest struct {
	*tchttp.BaseRequest
	
	// 用户分账标签键
	TagKey []*string `json:"TagKey,omitnil" name:"TagKey"`
}

func (r *DeleteAllocationTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAllocationTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TagKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAllocationTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAllocationTagResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteAllocationTagResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAllocationTagResponseParams `json:"Response"`
}

func (r *DeleteAllocationTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAllocationTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountBalanceRequestParams struct {

}

type DescribeAccountBalanceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAccountBalanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountBalanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountBalanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountBalanceResponseParams struct {
	// 接口做过变更,为兼容老接口,本字段与RealBalance相同,为当前真实可用余额,单位 分
	Balance *int64 `json:"Balance,omitnil" name:"Balance"`

	// 查询的用户Uin
	Uin *uint64 `json:"Uin,omitnil" name:"Uin"`

	// 当前真实可用余额,单位 分
	RealBalance *float64 `json:"RealBalance,omitnil" name:"RealBalance"`

	// 现金账户余额,单位 分
	CashAccountBalance *float64 `json:"CashAccountBalance,omitnil" name:"CashAccountBalance"`

	// 收益转入账户余额,单位 分
	IncomeIntoAccountBalance *float64 `json:"IncomeIntoAccountBalance,omitnil" name:"IncomeIntoAccountBalance"`

	// 赠送账户余额,单位 分
	PresentAccountBalance *float64 `json:"PresentAccountBalance,omitnil" name:"PresentAccountBalance"`

	// 冻结金额,单位 分
	FreezeAmount *float64 `json:"FreezeAmount,omitnil" name:"FreezeAmount"`

	// 欠费金额,单位 分
	OweAmount *float64 `json:"OweAmount,omitnil" name:"OweAmount"`

	// 是否允许欠费消费
	IsAllowArrears *bool `json:"IsAllowArrears,omitnil" name:"IsAllowArrears"`

	// 是否限制信用额度
	IsCreditLimited *bool `json:"IsCreditLimited,omitnil" name:"IsCreditLimited"`

	// 信用额度
	CreditAmount *float64 `json:"CreditAmount,omitnil" name:"CreditAmount"`

	// 可用信用额度
	CreditBalance *float64 `json:"CreditBalance,omitnil" name:"CreditBalance"`

	// 真实可用信用额度
	RealCreditBalance *float64 `json:"RealCreditBalance,omitnil" name:"RealCreditBalance"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeAccountBalanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountBalanceResponseParams `json:"Response"`
}

func (r *DescribeAccountBalanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountBalanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillDetailForOrganizationRequestParams struct {
	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，依次类推
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 数量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	//
	// Deprecated: PeriodType is deprecated.
	PeriodType *string `json:"PeriodType,omitnil" name:"PeriodType"`

	// 月份，格式为yyyy-mm，Month和BeginTime&EndTime必传一个，如果有传BeginTime&EndTime则Month字段无效。最多可拉取近18个月内的数据。
	Month *string `json:"Month,omitnil" name:"Month"`

	// 周期开始时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据。最多可拉取18个月内的数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 周期结束时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据。最多可拉取近18个月内的数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil" name:"NeedRecordNum"`

	// 付费模式 prePay(表示包年包月)/postPay(表示按时按量)
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 查询交易类型（请使用交易类型名称入参），入参示例枚举如下：
	// 包年包月新购
	// 包年包月续费
	// 包年包月配置变更
	// 包年包月退款 
	// 按量计费扣费 
	// 线下项目扣费 
	// 线下产品扣费 
	// 调账扣费 
	// 调账补偿 
	// 按量计费小时结 
	// 按量计费日结 
	// 按量计费月结 
	// 竞价实例小时结 
	// 线下项目调账补偿 
	// 线下产品调账补偿 
	// 优惠扣费 
	// 优惠补偿 
	// 按量计费迁入资源 
	// 按量计费迁出资源 
	// 包年包月迁入资源 
	// 包年包月迁出资源 
	// 预付费用 
	// 小时费用 
	// 预留实例退款 
	// 按量计费冲正 
	// 包年包月转按量 
	// 保底扣款 
	// 节省计划小时费用
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// 项目ID:资源所属项目ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitnil" name:"BusinessCode"`

	// 上一次请求返回的上下文信息，翻页查询Month>=2023-05的月份的数据可加快查询速度，数据量10万级别以上的用户建议使用，查询速度可提升2~10倍
	Context *string `json:"Context,omitnil" name:"Context"`
}

type DescribeBillDetailForOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，依次类推
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 数量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	PeriodType *string `json:"PeriodType,omitnil" name:"PeriodType"`

	// 月份，格式为yyyy-mm，Month和BeginTime&EndTime必传一个，如果有传BeginTime&EndTime则Month字段无效。最多可拉取近18个月内的数据。
	Month *string `json:"Month,omitnil" name:"Month"`

	// 周期开始时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据。最多可拉取18个月内的数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 周期结束时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据。最多可拉取近18个月内的数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil" name:"NeedRecordNum"`

	// 付费模式 prePay(表示包年包月)/postPay(表示按时按量)
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 查询交易类型（请使用交易类型名称入参），入参示例枚举如下：
	// 包年包月新购
	// 包年包月续费
	// 包年包月配置变更
	// 包年包月退款 
	// 按量计费扣费 
	// 线下项目扣费 
	// 线下产品扣费 
	// 调账扣费 
	// 调账补偿 
	// 按量计费小时结 
	// 按量计费日结 
	// 按量计费月结 
	// 竞价实例小时结 
	// 线下项目调账补偿 
	// 线下产品调账补偿 
	// 优惠扣费 
	// 优惠补偿 
	// 按量计费迁入资源 
	// 按量计费迁出资源 
	// 包年包月迁入资源 
	// 包年包月迁出资源 
	// 预付费用 
	// 小时费用 
	// 预留实例退款 
	// 按量计费冲正 
	// 包年包月转按量 
	// 保底扣款 
	// 节省计划小时费用
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// 项目ID:资源所属项目ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitnil" name:"BusinessCode"`

	// 上一次请求返回的上下文信息，翻页查询Month>=2023-05的月份的数据可加快查询速度，数据量10万级别以上的用户建议使用，查询速度可提升2~10倍
	Context *string `json:"Context,omitnil" name:"Context"`
}

func (r *DescribeBillDetailForOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillDetailForOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PeriodType")
	delete(f, "Month")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "NeedRecordNum")
	delete(f, "PayMode")
	delete(f, "ResourceId")
	delete(f, "ActionType")
	delete(f, "ProjectId")
	delete(f, "BusinessCode")
	delete(f, "Context")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillDetailForOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillDetailForOrganizationResponseParams struct {
	// 详情列表
	DetailSet []*DistributionBillDetail `json:"DetailSet,omitnil" name:"DetailSet"`

	// 总记录数，24小时缓存一次，可能比实际总记录数少
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 本次请求的上下文信息，可用于下一次请求的请求参数中，加快查询速度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Context *string `json:"Context,omitnil" name:"Context"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBillDetailForOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillDetailForOrganizationResponseParams `json:"Response"`
}

func (r *DescribeBillDetailForOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillDetailForOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillDetailRequestParams struct {
	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，依次类推
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 数量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	//
	// Deprecated: PeriodType is deprecated.
	PeriodType *string `json:"PeriodType,omitnil" name:"PeriodType"`

	// 月份，格式为yyyy-mm，Month和BeginTime&EndTime必传一个，如果有传BeginTime&EndTime则Month字段无效。最多可拉取近18个月内的数据。
	Month *string `json:"Month,omitnil" name:"Month"`

	// 周期开始时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据。最多可拉取18个月内的数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 周期结束时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据。最多可拉取近18个月内的数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil" name:"NeedRecordNum"`

	// 已废弃参数，未开放
	ProductCode *string `json:"ProductCode,omitnil" name:"ProductCode"`

	// 付费模式 prePay(表示包年包月)/postPay(表示按时按量)
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 查询交易类型（请使用交易类型名称入参），入参示例枚举如下：
	// 包年包月新购
	// 包年包月续费
	// 包年包月配置变更
	// 包年包月退款 
	// 按量计费扣费 
	// 线下项目扣费 
	// 线下产品扣费 
	// 调账扣费 
	// 调账补偿 
	// 按量计费小时结 
	// 按量计费日结 
	// 按量计费月结 
	// 竞价实例小时结 
	// 线下项目调账补偿 
	// 线下产品调账补偿 
	// 优惠扣费 
	// 优惠补偿 
	// 按量计费迁入资源 
	// 按量计费迁出资源 
	// 包年包月迁入资源 
	// 包年包月迁出资源 
	// 预付费用 
	// 小时费用 
	// 预留实例退款 
	// 按量计费冲正 
	// 包年包月转按量 
	// 保底扣款 
	// 节省计划小时费用
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// 项目ID:资源所属项目ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitnil" name:"BusinessCode"`

	// 上一次请求返回的上下文信息，翻页查询Month>=2023-05的月份的数据可加快查询速度，数据量10万级别以上的用户建议使用，查询速度可提升2~10倍
	Context *string `json:"Context,omitnil" name:"Context"`

	// 支付者的账号 ID（账号 ID 是用户在腾讯云的唯一账号标识），默认查询本账号账单，如集团管理账号需查询成员账号自付的账单，该字段需入参成员账号UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`
}

type DescribeBillDetailRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，依次类推
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 数量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	PeriodType *string `json:"PeriodType,omitnil" name:"PeriodType"`

	// 月份，格式为yyyy-mm，Month和BeginTime&EndTime必传一个，如果有传BeginTime&EndTime则Month字段无效。最多可拉取近18个月内的数据。
	Month *string `json:"Month,omitnil" name:"Month"`

	// 周期开始时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据。最多可拉取18个月内的数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 周期结束时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据。最多可拉取近18个月内的数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil" name:"NeedRecordNum"`

	// 已废弃参数，未开放
	ProductCode *string `json:"ProductCode,omitnil" name:"ProductCode"`

	// 付费模式 prePay(表示包年包月)/postPay(表示按时按量)
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 查询交易类型（请使用交易类型名称入参），入参示例枚举如下：
	// 包年包月新购
	// 包年包月续费
	// 包年包月配置变更
	// 包年包月退款 
	// 按量计费扣费 
	// 线下项目扣费 
	// 线下产品扣费 
	// 调账扣费 
	// 调账补偿 
	// 按量计费小时结 
	// 按量计费日结 
	// 按量计费月结 
	// 竞价实例小时结 
	// 线下项目调账补偿 
	// 线下产品调账补偿 
	// 优惠扣费 
	// 优惠补偿 
	// 按量计费迁入资源 
	// 按量计费迁出资源 
	// 包年包月迁入资源 
	// 包年包月迁出资源 
	// 预付费用 
	// 小时费用 
	// 预留实例退款 
	// 按量计费冲正 
	// 包年包月转按量 
	// 保底扣款 
	// 节省计划小时费用
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// 项目ID:资源所属项目ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitnil" name:"BusinessCode"`

	// 上一次请求返回的上下文信息，翻页查询Month>=2023-05的月份的数据可加快查询速度，数据量10万级别以上的用户建议使用，查询速度可提升2~10倍
	Context *string `json:"Context,omitnil" name:"Context"`

	// 支付者的账号 ID（账号 ID 是用户在腾讯云的唯一账号标识），默认查询本账号账单，如集团管理账号需查询成员账号自付的账单，该字段需入参成员账号UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`
}

func (r *DescribeBillDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PeriodType")
	delete(f, "Month")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "NeedRecordNum")
	delete(f, "ProductCode")
	delete(f, "PayMode")
	delete(f, "ResourceId")
	delete(f, "ActionType")
	delete(f, "ProjectId")
	delete(f, "BusinessCode")
	delete(f, "Context")
	delete(f, "PayerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillDetailResponseParams struct {
	// 详情列表
	DetailSet []*BillDetail `json:"DetailSet,omitnil" name:"DetailSet"`

	// 总记录数，24小时缓存一次，可能比实际总记录数少
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 本次请求的上下文信息，可用于下一次请求的请求参数中，加快查询速度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Context *string `json:"Context,omitnil" name:"Context"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBillDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillDetailResponseParams `json:"Response"`
}

func (r *DescribeBillDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillDownloadUrlRequestParams struct {
	// 账单类型，枚举值
	// billOverview=L0-PDF账单
	// billSummary=L1-汇总账单	
	// billResource=L2-资源账单	
	// billDetail=L3-明细账单	
	// billPack=账单包
	FileType *string `json:"FileType,omitnil" name:"FileType"`

	// 账单月份
	// 支持的最早开始月份为2021-01
	// L0-PDF&账单包不支持当月下载，当月账单请在次月1号19:00出账后下载
	Month *string `json:"Month,omitnil" name:"Month"`

	// 下载的账号 ID列表，默认查询本账号账单，如集团管理账号需下载成员账号自付的账单，该字段需入参成员账号UIN
	ChildUin []*string `json:"ChildUin,omitnil" name:"ChildUin"`
}

type DescribeBillDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// 账单类型，枚举值
	// billOverview=L0-PDF账单
	// billSummary=L1-汇总账单	
	// billResource=L2-资源账单	
	// billDetail=L3-明细账单	
	// billPack=账单包
	FileType *string `json:"FileType,omitnil" name:"FileType"`

	// 账单月份
	// 支持的最早开始月份为2021-01
	// L0-PDF&账单包不支持当月下载，当月账单请在次月1号19:00出账后下载
	Month *string `json:"Month,omitnil" name:"Month"`

	// 下载的账号 ID列表，默认查询本账号账单，如集团管理账号需下载成员账号自付的账单，该字段需入参成员账号UIN
	ChildUin []*string `json:"ChildUin,omitnil" name:"ChildUin"`
}

func (r *DescribeBillDownloadUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillDownloadUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FileType")
	delete(f, "Month")
	delete(f, "ChildUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillDownloadUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillDownloadUrlResponseParams struct {
	// 账单文件是否准备就绪，0文件生成中，1文件已生成
	Ready *int64 `json:"Ready,omitnil" name:"Ready"`

	// 账单文件下载链接，有效时长为一小时
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitnil" name:"DownloadUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBillDownloadUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillDownloadUrlResponseParams `json:"Response"`
}

func (r *DescribeBillDownloadUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillDownloadUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillListRequestParams struct {
	// 查询范围的起始时间（包含）时间格式 yyyy-MM-dd HH:mm:ss 开始时间和结束时间差值小于等于六个月
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询范围的结束时间（包含）时间格式 yyyy-MM-dd HH:mm:ss ，开始时间和结束时间差值小于等于六个月
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 翻页偏移量，初始值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页的限制数量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 交易类型： all所有交易类型，recharge充值，return退款，unblock解冻，agentin资金转入，advanced垫付，cash提现，deduct扣费，block冻结，agentout资金转出，repay垫付回款，repayment还款(仅国际信用账户)，adj_refund调增(仅国际信用账户)，adj_deduct调减(仅国际信用账户)
	PayType []*string `json:"PayType,omitnil" name:"PayType"`

	// 扣费模式，
	// 当所选的交易类型为扣费deduct时： 
	// all所有扣费类型;trade预付费支付;hour_h按量小时结;hour_d按量日结;hour_m按量月结;decompensate调账扣费;other第三方扣费;panshi 线下项目扣费;offline 线下产品扣费;
	// 
	// 当所选的交易类型为扣费recharge时： 
	// online 在线充值;bank-enterprice 银企直连;offline 线下充值;transfer 分成充值
	// 
	// 当所选的交易类型为扣费cash时： 
	// online 线上提现;offline 线下提现;panshi 赠送金清零
	// 
	// 当所选的交易类型为扣费advanced时： 
	// advanced 垫付充值
	// 
	// 当所选的交易类型为扣费repay时： 
	// panshi 垫付回款
	// 
	// 当所选的交易类型为扣费block时： 
	// other 第三方冻结;hour 按量冻结;month按月冻结
	// 
	// 当所选的交易类型为扣费return时： 
	// compensate 调账补偿;trade 预付费退款
	// 
	// 当所选的交易类型为扣费unblock时：
	// other 第三方解冻;hour 按量解冻;month 按月解冻
	SubPayType []*string `json:"SubPayType,omitnil" name:"SubPayType"`

	// 是否返回0元交易金额的交易项，取值：0-不返回，1-返回。不传该参数则不返回
	WithZeroAmount *uint64 `json:"WithZeroAmount,omitnil" name:"WithZeroAmount"`
}

type DescribeBillListRequest struct {
	*tchttp.BaseRequest
	
	// 查询范围的起始时间（包含）时间格式 yyyy-MM-dd HH:mm:ss 开始时间和结束时间差值小于等于六个月
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询范围的结束时间（包含）时间格式 yyyy-MM-dd HH:mm:ss ，开始时间和结束时间差值小于等于六个月
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 翻页偏移量，初始值为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页的限制数量
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 交易类型： all所有交易类型，recharge充值，return退款，unblock解冻，agentin资金转入，advanced垫付，cash提现，deduct扣费，block冻结，agentout资金转出，repay垫付回款，repayment还款(仅国际信用账户)，adj_refund调增(仅国际信用账户)，adj_deduct调减(仅国际信用账户)
	PayType []*string `json:"PayType,omitnil" name:"PayType"`

	// 扣费模式，
	// 当所选的交易类型为扣费deduct时： 
	// all所有扣费类型;trade预付费支付;hour_h按量小时结;hour_d按量日结;hour_m按量月结;decompensate调账扣费;other第三方扣费;panshi 线下项目扣费;offline 线下产品扣费;
	// 
	// 当所选的交易类型为扣费recharge时： 
	// online 在线充值;bank-enterprice 银企直连;offline 线下充值;transfer 分成充值
	// 
	// 当所选的交易类型为扣费cash时： 
	// online 线上提现;offline 线下提现;panshi 赠送金清零
	// 
	// 当所选的交易类型为扣费advanced时： 
	// advanced 垫付充值
	// 
	// 当所选的交易类型为扣费repay时： 
	// panshi 垫付回款
	// 
	// 当所选的交易类型为扣费block时： 
	// other 第三方冻结;hour 按量冻结;month按月冻结
	// 
	// 当所选的交易类型为扣费return时： 
	// compensate 调账补偿;trade 预付费退款
	// 
	// 当所选的交易类型为扣费unblock时：
	// other 第三方解冻;hour 按量解冻;month 按月解冻
	SubPayType []*string `json:"SubPayType,omitnil" name:"SubPayType"`

	// 是否返回0元交易金额的交易项，取值：0-不返回，1-返回。不传该参数则不返回
	WithZeroAmount *uint64 `json:"WithZeroAmount,omitnil" name:"WithZeroAmount"`
}

func (r *DescribeBillListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PayType")
	delete(f, "SubPayType")
	delete(f, "WithZeroAmount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillListResponseParams struct {
	// 收支明细列表
	TransactionList []*BillTransactionInfo `json:"TransactionList,omitnil" name:"TransactionList"`

	// 总条数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 退费总额，单位（分）
	ReturnAmount *float64 `json:"ReturnAmount,omitnil" name:"ReturnAmount"`

	// 充值总额，单位（分）
	RechargeAmount *float64 `json:"RechargeAmount,omitnil" name:"RechargeAmount"`

	// 冻结总额，单位（分）
	BlockAmount *float64 `json:"BlockAmount,omitnil" name:"BlockAmount"`

	// 解冻总额，单位（分）
	UnblockAmount *float64 `json:"UnblockAmount,omitnil" name:"UnblockAmount"`

	// 扣费总额，单位（分）
	DeductAmount *float64 `json:"DeductAmount,omitnil" name:"DeductAmount"`

	// 资金转入总额，单位（分）
	AgentInAmount *float64 `json:"AgentInAmount,omitnil" name:"AgentInAmount"`

	// 垫付充值总额，单位（分）
	AdvanceRechargeAmount *float64 `json:"AdvanceRechargeAmount,omitnil" name:"AdvanceRechargeAmount"`

	// 提现扣减总额，单位（分）
	WithdrawAmount *float64 `json:"WithdrawAmount,omitnil" name:"WithdrawAmount"`

	// 资金转出总额，单位（分）
	AgentOutAmount *float64 `json:"AgentOutAmount,omitnil" name:"AgentOutAmount"`

	// 还垫付总额，单位（分）
	AdvancePayAmount *float64 `json:"AdvancePayAmount,omitnil" name:"AdvancePayAmount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBillListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillListResponseParams `json:"Response"`
}

func (r *DescribeBillListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillResourceSummaryForOrganizationRequestParams struct {
	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，依次类推
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 月份，格式为yyyy-mm。不能早于开通账单2.0的月份
	Month *string `json:"Month,omitnil" name:"Month"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	//
	// Deprecated: PeriodType is deprecated.
	PeriodType *string `json:"PeriodType,omitnil" name:"PeriodType"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil" name:"NeedRecordNum"`

	// 查询交易类型（请使用交易类型名称入参），入参示例枚举如下：
	// 包年包月新购
	// 包年包月续费
	// 包年包月配置变更
	// 包年包月退款 
	// 按量计费扣费 
	// 线下项目扣费 
	// 线下产品扣费 
	// 调账扣费 
	// 调账补偿 
	// 按量计费小时结 
	// 按量计费日结 
	// 按量计费月结 
	// 竞价实例小时结 
	// 线下项目调账补偿 
	// 线下产品调账补偿 
	// 优惠扣费 
	// 优惠补偿 
	// 按量计费迁入资源 
	// 按量计费迁出资源 
	// 包年包月迁入资源 
	// 包年包月迁出资源 
	// 预付费用 
	// 小时费用 
	// 预留实例退款 
	// 按量计费冲正 
	// 包年包月转按量 
	// 保底扣款 
	// 节省计划小时费用
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 付费模式 prePay/postPay
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitnil" name:"BusinessCode"`

	// 分账标签键，用户自定义（支持2021-01以后账单查询）
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 分账标签值，该参数为空表示该标签键下未设置标签值的记录
	// （支持2021-01以后账单查询）
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type DescribeBillResourceSummaryForOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，依次类推
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 月份，格式为yyyy-mm。不能早于开通账单2.0的月份
	Month *string `json:"Month,omitnil" name:"Month"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	PeriodType *string `json:"PeriodType,omitnil" name:"PeriodType"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil" name:"NeedRecordNum"`

	// 查询交易类型（请使用交易类型名称入参），入参示例枚举如下：
	// 包年包月新购
	// 包年包月续费
	// 包年包月配置变更
	// 包年包月退款 
	// 按量计费扣费 
	// 线下项目扣费 
	// 线下产品扣费 
	// 调账扣费 
	// 调账补偿 
	// 按量计费小时结 
	// 按量计费日结 
	// 按量计费月结 
	// 竞价实例小时结 
	// 线下项目调账补偿 
	// 线下产品调账补偿 
	// 优惠扣费 
	// 优惠补偿 
	// 按量计费迁入资源 
	// 按量计费迁出资源 
	// 包年包月迁入资源 
	// 包年包月迁出资源 
	// 预付费用 
	// 小时费用 
	// 预留实例退款 
	// 按量计费冲正 
	// 包年包月转按量 
	// 保底扣款 
	// 节省计划小时费用
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 付费模式 prePay/postPay
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitnil" name:"BusinessCode"`

	// 分账标签键，用户自定义（支持2021-01以后账单查询）
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 分账标签值，该参数为空表示该标签键下未设置标签值的记录
	// （支持2021-01以后账单查询）
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

func (r *DescribeBillResourceSummaryForOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillResourceSummaryForOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Month")
	delete(f, "PeriodType")
	delete(f, "NeedRecordNum")
	delete(f, "ActionType")
	delete(f, "ResourceId")
	delete(f, "PayMode")
	delete(f, "BusinessCode")
	delete(f, "TagKey")
	delete(f, "TagValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillResourceSummaryForOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillResourceSummaryForOrganizationResponseParams struct {
	// 资源汇总列表
	ResourceSummarySet []*BillDistributionResourceSummary `json:"ResourceSummarySet,omitnil" name:"ResourceSummarySet"`

	// 资源汇总列表总数，入参NeedRecordNum为0时不返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBillResourceSummaryForOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillResourceSummaryForOrganizationResponseParams `json:"Response"`
}

func (r *DescribeBillResourceSummaryForOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillResourceSummaryForOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillResourceSummaryRequestParams struct {
	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，依次类推
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 月份，格式为yyyy-mm。不能早于开通账单2.0的月份
	Month *string `json:"Month,omitnil" name:"Month"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	//
	// Deprecated: PeriodType is deprecated.
	PeriodType *string `json:"PeriodType,omitnil" name:"PeriodType"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil" name:"NeedRecordNum"`

	// 查询交易类型（请使用交易类型名称入参），入参示例枚举如下：
	// 包年包月新购
	// 包年包月续费
	// 包年包月配置变更
	// 包年包月退款 
	// 按量计费扣费 
	// 线下项目扣费 
	// 线下产品扣费 
	// 调账扣费 
	// 调账补偿 
	// 按量计费小时结 
	// 按量计费日结 
	// 按量计费月结 
	// 竞价实例小时结 
	// 线下项目调账补偿 
	// 线下产品调账补偿 
	// 优惠扣费 
	// 优惠补偿 
	// 按量计费迁入资源 
	// 按量计费迁出资源 
	// 包年包月迁入资源 
	// 包年包月迁出资源 
	// 预付费用 
	// 小时费用 
	// 预留实例退款 
	// 按量计费冲正 
	// 包年包月转按量 
	// 保底扣款 
	// 节省计划小时费用
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 付费模式 prePay/postPay
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitnil" name:"BusinessCode"`

	// 支付者的账号 ID（账号 ID 是用户在腾讯云的唯一账号标识），默认查询本账号账单，如集团管理账号需查询成员账号自付的账单，该字段需入参成员账号UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`

	// 分账标签键，用户自定义（支持2021-01以后账单查询）
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 分账标签值，该参数为空表示该标签键下未设置标签值的记录
	// （支持2021-01以后账单查询）
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type DescribeBillResourceSummaryRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，依次类推
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 月份，格式为yyyy-mm。不能早于开通账单2.0的月份
	Month *string `json:"Month,omitnil" name:"Month"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	PeriodType *string `json:"PeriodType,omitnil" name:"PeriodType"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil" name:"NeedRecordNum"`

	// 查询交易类型（请使用交易类型名称入参），入参示例枚举如下：
	// 包年包月新购
	// 包年包月续费
	// 包年包月配置变更
	// 包年包月退款 
	// 按量计费扣费 
	// 线下项目扣费 
	// 线下产品扣费 
	// 调账扣费 
	// 调账补偿 
	// 按量计费小时结 
	// 按量计费日结 
	// 按量计费月结 
	// 竞价实例小时结 
	// 线下项目调账补偿 
	// 线下产品调账补偿 
	// 优惠扣费 
	// 优惠补偿 
	// 按量计费迁入资源 
	// 按量计费迁出资源 
	// 包年包月迁入资源 
	// 包年包月迁出资源 
	// 预付费用 
	// 小时费用 
	// 预留实例退款 
	// 按量计费冲正 
	// 包年包月转按量 
	// 保底扣款 
	// 节省计划小时费用
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 付费模式 prePay/postPay
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitnil" name:"BusinessCode"`

	// 支付者的账号 ID（账号 ID 是用户在腾讯云的唯一账号标识），默认查询本账号账单，如集团管理账号需查询成员账号自付的账单，该字段需入参成员账号UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`

	// 分账标签键，用户自定义（支持2021-01以后账单查询）
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 分账标签值，该参数为空表示该标签键下未设置标签值的记录
	// （支持2021-01以后账单查询）
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

func (r *DescribeBillResourceSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillResourceSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Month")
	delete(f, "PeriodType")
	delete(f, "NeedRecordNum")
	delete(f, "ActionType")
	delete(f, "ResourceId")
	delete(f, "PayMode")
	delete(f, "BusinessCode")
	delete(f, "PayerUin")
	delete(f, "TagKey")
	delete(f, "TagValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillResourceSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillResourceSummaryResponseParams struct {
	// 资源汇总列表
	ResourceSummarySet []*BillResourceSummary `json:"ResourceSummarySet,omitnil" name:"ResourceSummarySet"`

	// 资源汇总列表总数，入参NeedRecordNum为0时不返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBillResourceSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillResourceSummaryResponseParams `json:"Response"`
}

func (r *DescribeBillResourceSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillResourceSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByPayModeRequestParams struct {
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`
}

type DescribeBillSummaryByPayModeRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`
}

func (r *DescribeBillSummaryByPayModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByPayModeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "PayerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryByPayModeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByPayModeResponseParams struct {
	// 数据是否准备好，0准备中，1已就绪。（Ready=0，为当前UIN首次进行初始化出账，预计需要5~10分钟出账，请于10分钟后重试即可）
	Ready *uint64 `json:"Ready,omitnil" name:"Ready"`

	// 各付费模式花费分布详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryOverview []*PayModeSummaryOverviewItem `json:"SummaryOverview,omitnil" name:"SummaryOverview"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBillSummaryByPayModeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillSummaryByPayModeResponseParams `json:"Response"`
}

func (r *DescribeBillSummaryByPayModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByPayModeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByProductRequestParams struct {
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`

	// 款项类别，与L0账单上的汇总类别对应。
	// 此参数自账单3.0（即2021-05）之后开始生效。
	// 枚举值：
	// consume-消费
	// refund-退款
	// adjustment-调账
	PayType *string `json:"PayType,omitnil" name:"PayType"`
}

type DescribeBillSummaryByProductRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`

	// 款项类别，与L0账单上的汇总类别对应。
	// 此参数自账单3.0（即2021-05）之后开始生效。
	// 枚举值：
	// consume-消费
	// refund-退款
	// adjustment-调账
	PayType *string `json:"PayType,omitnil" name:"PayType"`
}

func (r *DescribeBillSummaryByProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "PayerUin")
	delete(f, "PayType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryByProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByProductResponseParams struct {
	// 数据是否准备好，0准备中，1已就绪。（Ready=0，为当前UIN首次进行初始化出账，预计需要5~10分钟出账，请于10分钟后重试即可）
	Ready *uint64 `json:"Ready,omitnil" name:"Ready"`

	// 总花费详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryTotal *BusinessSummaryTotal `json:"SummaryTotal,omitnil" name:"SummaryTotal"`

	// 各产品花费分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryOverview []*BusinessSummaryOverviewItem `json:"SummaryOverview,omitnil" name:"SummaryOverview"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBillSummaryByProductResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillSummaryByProductResponseParams `json:"Response"`
}

func (r *DescribeBillSummaryByProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByProjectRequestParams struct {
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`
}

type DescribeBillSummaryByProjectRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`
}

func (r *DescribeBillSummaryByProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "PayerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryByProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByProjectResponseParams struct {
	// 数据是否准备好，0准备中，1已就绪。（Ready=0，为当前UIN首次进行初始化出账，预计需要5~10分钟出账，请于10分钟后重试即可）
	Ready *uint64 `json:"Ready,omitnil" name:"Ready"`

	// 各项目花费分布详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryOverview []*ProjectSummaryOverviewItem `json:"SummaryOverview,omitnil" name:"SummaryOverview"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBillSummaryByProjectResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillSummaryByProjectResponseParams `json:"Response"`
}

func (r *DescribeBillSummaryByProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByRegionRequestParams struct {
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`
}

type DescribeBillSummaryByRegionRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`
}

func (r *DescribeBillSummaryByRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "PayerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryByRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByRegionResponseParams struct {
	// 数据是否准备好，0准备中，1已就绪。（Ready=0，为当前UIN首次进行初始化出账，预计需要5~10分钟出账，请于10分钟后重试即可）
	Ready *uint64 `json:"Ready,omitnil" name:"Ready"`

	// 各地域花费分布详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryOverview []*RegionSummaryOverviewItem `json:"SummaryOverview,omitnil" name:"SummaryOverview"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBillSummaryByRegionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillSummaryByRegionResponseParams `json:"Response"`
}

func (r *DescribeBillSummaryByRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByTagRequestParams struct {
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分账标签键，用户自定义
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`

	// 分账标签值
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type DescribeBillSummaryByTagRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 分账标签键，用户自定义
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`

	// 分账标签值
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

func (r *DescribeBillSummaryByTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "TagKey")
	delete(f, "PayerUin")
	delete(f, "TagValue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryByTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryByTagResponseParams struct {
	// 数据是否准备好，0准备中，1已就绪。（Ready=0，为当前UIN首次进行初始化出账，预计需要5~10分钟出账，请于10分钟后重试即可）
	Ready *uint64 `json:"Ready,omitnil" name:"Ready"`

	// 各标签值花费分布详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryOverview []*TagSummaryOverviewItem `json:"SummaryOverview,omitnil" name:"SummaryOverview"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryTotal *SummaryTotal `json:"SummaryTotal,omitnil" name:"SummaryTotal"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBillSummaryByTagResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillSummaryByTagResponseParams `json:"Response"`
}

func (r *DescribeBillSummaryByTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryByTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryForOrganizationRequestParams struct {
	// 账单月份，格式为2023-04
	Month *string `json:"Month,omitnil" name:"Month"`

	// 账单维度类型，枚举值如下：business、project、region、payMode、tag
	GroupType *string `json:"GroupType,omitnil" name:"GroupType"`

	// 标签键，GroupType=tag获取标签维度账单时传
	TagKey []*string `json:"TagKey,omitnil" name:"TagKey"`
}

type DescribeBillSummaryForOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 账单月份，格式为2023-04
	Month *string `json:"Month,omitnil" name:"Month"`

	// 账单维度类型，枚举值如下：business、project、region、payMode、tag
	GroupType *string `json:"GroupType,omitnil" name:"GroupType"`

	// 标签键，GroupType=tag获取标签维度账单时传
	TagKey []*string `json:"TagKey,omitnil" name:"TagKey"`
}

func (r *DescribeBillSummaryForOrganizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryForOrganizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "GroupType")
	delete(f, "TagKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryForOrganizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryForOrganizationResponseParams struct {
	// 数据是否准备好，0准备中，1已就绪。（Ready=0，为当前UIN首次进行初始化出账，预计需要5~10分钟出账，请于10分钟后重试即可）
	Ready *uint64 `json:"Ready,omitnil" name:"Ready"`

	// 账单多维度汇总消费详情
	SummaryDetail []*SummaryDetail `json:"SummaryDetail,omitnil" name:"SummaryDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBillSummaryForOrganizationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillSummaryForOrganizationResponseParams `json:"Response"`
}

func (r *DescribeBillSummaryForOrganizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryForOrganizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryRequestParams struct {
	// 账单月份，格式为2023-04
	Month *string `json:"Month,omitnil" name:"Month"`

	// 账单维度类型，枚举值如下：business、project、region、payMode、tag
	GroupType *string `json:"GroupType,omitnil" name:"GroupType"`

	// 标签键，GroupType=tag获取标签维度账单时传
	TagKey []*string `json:"TagKey,omitnil" name:"TagKey"`
}

type DescribeBillSummaryRequest struct {
	*tchttp.BaseRequest
	
	// 账单月份，格式为2023-04
	Month *string `json:"Month,omitnil" name:"Month"`

	// 账单维度类型，枚举值如下：business、project、region、payMode、tag
	GroupType *string `json:"GroupType,omitnil" name:"GroupType"`

	// 标签键，GroupType=tag获取标签维度账单时传
	TagKey []*string `json:"TagKey,omitnil" name:"TagKey"`
}

func (r *DescribeBillSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "GroupType")
	delete(f, "TagKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillSummaryResponseParams struct {
	// 数据是否准备好，0准备中，1已就绪。（Ready=0，为当前UIN首次进行初始化出账，预计需要5~10分钟出账，请于10分钟后重试即可）
	Ready *uint64 `json:"Ready,omitnil" name:"Ready"`

	// 账单多维度汇总消费详情
	SummaryDetail []*SummaryDetail `json:"SummaryDetail,omitnil" name:"SummaryDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBillSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillSummaryResponseParams `json:"Response"`
}

func (r *DescribeBillSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostDetailRequestParams struct {
	// 数量，最大值为100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 周期开始时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为同一月份，暂不支持跨月拉取。可拉取的数据是开通成本分析后，且距今 24 个月内的数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 周期结束时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为同一月份，暂不支持跨月拉取。可拉取的数据是开通成本分析后，且距今 24 个月内的数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil" name:"NeedRecordNum"`

	// 月份，格式为yyyy-mm，Month和BeginTime&EndTime必传一个，如果有传BeginTime&EndTime则Month字段无效。不能早于开通成本分析的月份，最多可拉取24个月内的数据。
	Month *string `json:"Month,omitnil" name:"Month"`

	// 查询指定产品信息（暂时未开放获取）
	ProductCode *string `json:"ProductCode,omitnil" name:"ProductCode"`

	// 付费模式 prePay/postPay
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`
}

type DescribeCostDetailRequest struct {
	*tchttp.BaseRequest
	
	// 数量，最大值为100
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 周期开始时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为同一月份，暂不支持跨月拉取。可拉取的数据是开通成本分析后，且距今 24 个月内的数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 周期结束时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为同一月份，暂不支持跨月拉取。可拉取的数据是开通成本分析后，且距今 24 个月内的数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil" name:"NeedRecordNum"`

	// 月份，格式为yyyy-mm，Month和BeginTime&EndTime必传一个，如果有传BeginTime&EndTime则Month字段无效。不能早于开通成本分析的月份，最多可拉取24个月内的数据。
	Month *string `json:"Month,omitnil" name:"Month"`

	// 查询指定产品信息（暂时未开放获取）
	ProductCode *string `json:"ProductCode,omitnil" name:"ProductCode"`

	// 付费模式 prePay/postPay
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`
}

func (r *DescribeCostDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "NeedRecordNum")
	delete(f, "Month")
	delete(f, "ProductCode")
	delete(f, "PayMode")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCostDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostDetailResponseParams struct {
	// 消耗明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetailSet []*CostDetail `json:"DetailSet,omitnil" name:"DetailSet"`

	// 记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCostDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCostDetailResponseParams `json:"Response"`
}

func (r *DescribeCostDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostSummaryByProductRequestParams struct {
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 每次获取数据量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认从0开始
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil" name:"NeedRecordNum"`
}

type DescribeCostSummaryByProductRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 每次获取数据量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认从0开始
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil" name:"NeedRecordNum"`
}

func (r *DescribeCostSummaryByProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostSummaryByProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "PayerUin")
	delete(f, "NeedRecordNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCostSummaryByProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostSummaryByProductResponseParams struct {
	// 数据是否准备好，0未准备好，1准备好
	Ready *uint64 `json:"Ready,omitnil" name:"Ready"`

	// 消耗详情
	Total *ConsumptionSummaryTotal `json:"Total,omitnil" name:"Total"`

	// 消耗按产品汇总详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ConsumptionBusinessSummaryDataItem `json:"Data,omitnil" name:"Data"`

	// 记录数量，NeedRecordNum为0是返回null
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordNum *uint64 `json:"RecordNum,omitnil" name:"RecordNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCostSummaryByProductResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCostSummaryByProductResponseParams `json:"Response"`
}

func (r *DescribeCostSummaryByProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostSummaryByProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostSummaryByProjectRequestParams struct {
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 每次获取数据量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认从0开始
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil" name:"NeedRecordNum"`
}

type DescribeCostSummaryByProjectRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 每次获取数据量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认从0开始
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil" name:"NeedRecordNum"`
}

func (r *DescribeCostSummaryByProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostSummaryByProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "PayerUin")
	delete(f, "NeedRecordNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCostSummaryByProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostSummaryByProjectResponseParams struct {
	// 数据是否准备好，0未准备好，1准备好
	Ready *uint64 `json:"Ready,omitnil" name:"Ready"`

	// 消耗详情
	Total *ConsumptionSummaryTotal `json:"Total,omitnil" name:"Total"`

	// 消耗按业务汇总详情
	Data []*ConsumptionProjectSummaryDataItem `json:"Data,omitnil" name:"Data"`

	// 记录数量，NeedRecordNum为0是返回null
	RecordNum *uint64 `json:"RecordNum,omitnil" name:"RecordNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCostSummaryByProjectResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCostSummaryByProjectResponseParams `json:"Response"`
}

func (r *DescribeCostSummaryByProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostSummaryByProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostSummaryByRegionRequestParams struct {
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 每次获取数据量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认从0开始
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil" name:"NeedRecordNum"`
}

type DescribeCostSummaryByRegionRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 每次获取数据量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认从0开始
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil" name:"NeedRecordNum"`
}

func (r *DescribeCostSummaryByRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostSummaryByRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "PayerUin")
	delete(f, "NeedRecordNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCostSummaryByRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostSummaryByRegionResponseParams struct {
	// 数据是否准备好，0未准备好，1准备好
	Ready *uint64 `json:"Ready,omitnil" name:"Ready"`

	// 消耗详情
	Total *ConsumptionSummaryTotal `json:"Total,omitnil" name:"Total"`

	// 消耗按地域汇总详情
	Data []*ConsumptionRegionSummaryDataItem `json:"Data,omitnil" name:"Data"`

	// 记录数量，NeedRecordNum为0是返回null
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordNum *uint64 `json:"RecordNum,omitnil" name:"RecordNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCostSummaryByRegionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCostSummaryByRegionResponseParams `json:"Response"`
}

func (r *DescribeCostSummaryByRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostSummaryByRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostSummaryByResourceRequestParams struct {
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 每次获取数据量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认从0开始
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil" name:"NeedRecordNum"`

	// 是否需要返回过滤条件，0不需要，1需要，默认不需要
	NeedConditionValue *uint64 `json:"NeedConditionValue,omitnil" name:"NeedConditionValue"`

	// 过滤条件，只支持ResourceKeyword(资源关键字，支持资源id及资源名称模糊查询)，ProjectIds（项目id），RegionIds(地域id)，PayModes(付费模式，可选prePay和postPay)，HideFreeCost（是否隐藏0元流水，可选0和1），OrderByCost（按费用排序规则，可选desc和asc）
	Conditions *Conditions `json:"Conditions,omitnil" name:"Conditions"`
}

type DescribeCostSummaryByResourceRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 每次获取数据量，最大值为100
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量,默认从0开始
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil" name:"NeedRecordNum"`

	// 是否需要返回过滤条件，0不需要，1需要，默认不需要
	NeedConditionValue *uint64 `json:"NeedConditionValue,omitnil" name:"NeedConditionValue"`

	// 过滤条件，只支持ResourceKeyword(资源关键字，支持资源id及资源名称模糊查询)，ProjectIds（项目id），RegionIds(地域id)，PayModes(付费模式，可选prePay和postPay)，HideFreeCost（是否隐藏0元流水，可选0和1），OrderByCost（按费用排序规则，可选desc和asc）
	Conditions *Conditions `json:"Conditions,omitnil" name:"Conditions"`
}

func (r *DescribeCostSummaryByResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostSummaryByResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "PayerUin")
	delete(f, "NeedRecordNum")
	delete(f, "NeedConditionValue")
	delete(f, "Conditions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCostSummaryByResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostSummaryByResourceResponseParams struct {
	// 数据是否准备好，0未准备好，1准备好
	Ready *uint64 `json:"Ready,omitnil" name:"Ready"`

	// 消耗详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *ConsumptionSummaryTotal `json:"Total,omitnil" name:"Total"`

	// 过滤条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionValue *ConsumptionResourceSummaryConditionValue `json:"ConditionValue,omitnil" name:"ConditionValue"`

	// 记录数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordNum *uint64 `json:"RecordNum,omitnil" name:"RecordNum"`

	// 资源消耗详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ConsumptionResourceSummaryDataItem `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCostSummaryByResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCostSummaryByResourceResponseParams `json:"Response"`
}

func (r *DescribeCostSummaryByResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostSummaryByResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDealsByCondRequestParams struct {
	// 开始时间 2016-01-01 00:00:00
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间 2016-02-01 00:00:00 建议跨度不超过3个月
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 一页多少条数据，默认是20条，最大不超过1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 第多少页，从0开始，默认是0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 订单状态,默认为4（成功的订单）
	// 订单的状态
	// 1：未支付
	// 2：已支付3：发货中
	// 4：已发货
	// 5：发货失败
	// 6：已退款
	// 7：已关单
	// 8：订单过期
	// 9：订单已失效
	// 10：产品已失效
	// 11：代付拒绝
	// 12：支付中
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 订单号
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// 大订单号
	BigDealId *string `json:"BigDealId,omitnil" name:"BigDealId"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`
}

type DescribeDealsByCondRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间 2016-01-01 00:00:00
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间 2016-02-01 00:00:00 建议跨度不超过3个月
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 一页多少条数据，默认是20条，最大不超过1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 第多少页，从0开始，默认是0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 订单状态,默认为4（成功的订单）
	// 订单的状态
	// 1：未支付
	// 2：已支付3：发货中
	// 4：已发货
	// 5：发货失败
	// 6：已退款
	// 7：已关单
	// 8：订单过期
	// 9：订单已失效
	// 10：产品已失效
	// 11：代付拒绝
	// 12：支付中
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 订单号
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// 大订单号
	BigDealId *string `json:"BigDealId,omitnil" name:"BigDealId"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`
}

func (r *DescribeDealsByCondRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDealsByCondRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Status")
	delete(f, "OrderId")
	delete(f, "BigDealId")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDealsByCondRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDealsByCondResponseParams struct {
	// 订单列表
	Deals []*Deal `json:"Deals,omitnil" name:"Deals"`

	// 订单总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDealsByCondResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDealsByCondResponseParams `json:"Response"`
}

func (r *DescribeDealsByCondResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDealsByCondResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDosageCosDetailByDateRequestParams struct {
	// 查询用量开始时间，例如：2020-09-01
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 查询用量结束时间，例如：2020-09-30（与开始时间同月，不支持跨月查询）
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// COS 存储桶名称，可通过Get Service 接口是用来获取请求者名下的所有存储空间列表（Bucket list）https://cloud.tencent.com/document/product/436/8291
	BucketName *string `json:"BucketName,omitnil" name:"BucketName"`
}

type DescribeDosageCosDetailByDateRequest struct {
	*tchttp.BaseRequest
	
	// 查询用量开始时间，例如：2020-09-01
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 查询用量结束时间，例如：2020-09-30（与开始时间同月，不支持跨月查询）
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// COS 存储桶名称，可通过Get Service 接口是用来获取请求者名下的所有存储空间列表（Bucket list）https://cloud.tencent.com/document/product/436/8291
	BucketName *string `json:"BucketName,omitnil" name:"BucketName"`
}

func (r *DescribeDosageCosDetailByDateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDosageCosDetailByDateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "BucketName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDosageCosDetailByDateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDosageCosDetailByDateResponseParams struct {
	// 用量数组
	DetailSets []*CosDetailSets `json:"DetailSets,omitnil" name:"DetailSets"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDosageCosDetailByDateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDosageCosDetailByDateResponseParams `json:"Response"`
}

func (r *DescribeDosageCosDetailByDateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDosageCosDetailByDateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDosageDetailByDateRequestParams struct {
	// 查询账单开始日期，如 2019-01-01
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 查询账单结束日期，如 2019-01-01， 时间跨度不超过7天
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 互动直播：
	// 10194   互动直播-核心机房           :
	// 10195   互动直播-边缘机房
	// 
	// cdn业务：
	// 102383：CDN静态加速流量(国内)
	// 102384：CDN静态加速带宽(国内)
	// 102385：CDN静态加速流量(海外)
	// 102386：CDN静态加速带宽(海外)
	// 
	// 100967：弹性公网IP-按流量计费
	// 101065：公网负载均衡-按流量计费
	// 
	// 视频直播
	// 10226 视频直播流量(国内)
	// 10227 视频直播带宽(国内)
	// 100763 视频直播流量(海外)
	// 100762 视频直播宽带(海外)
	ProductCode *string `json:"ProductCode,omitnil" name:"ProductCode"`

	// 查询域名 例如 www.qq.com
	// 非CDN业务查询时传入空字符串，返回的值为空
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 1、如果为空，则返回EIP或CLB所有实例的明细；
	// 2、如果传入实例名，则返回该实例明细
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

type DescribeDosageDetailByDateRequest struct {
	*tchttp.BaseRequest
	
	// 查询账单开始日期，如 2019-01-01
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 查询账单结束日期，如 2019-01-01， 时间跨度不超过7天
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 互动直播：
	// 10194   互动直播-核心机房           :
	// 10195   互动直播-边缘机房
	// 
	// cdn业务：
	// 102383：CDN静态加速流量(国内)
	// 102384：CDN静态加速带宽(国内)
	// 102385：CDN静态加速流量(海外)
	// 102386：CDN静态加速带宽(海外)
	// 
	// 100967：弹性公网IP-按流量计费
	// 101065：公网负载均衡-按流量计费
	// 
	// 视频直播
	// 10226 视频直播流量(国内)
	// 10227 视频直播带宽(国内)
	// 100763 视频直播流量(海外)
	// 100762 视频直播宽带(海外)
	ProductCode *string `json:"ProductCode,omitnil" name:"ProductCode"`

	// 查询域名 例如 www.qq.com
	// 非CDN业务查询时传入空字符串，返回的值为空
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 1、如果为空，则返回EIP或CLB所有实例的明细；
	// 2、如果传入实例名，则返回该实例明细
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

func (r *DescribeDosageDetailByDateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDosageDetailByDateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "ProductCode")
	delete(f, "Domain")
	delete(f, "InstanceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDosageDetailByDateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDosageDetailByDateResponseParams struct {
	// 计量单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 用量数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetailSets []*DetailSet `json:"DetailSets,omitnil" name:"DetailSets"`

	// 错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetCode *int64 `json:"RetCode,omitnil" name:"RetCode"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetMsg *string `json:"RetMsg,omitnil" name:"RetMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeDosageDetailByDateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDosageDetailByDateResponseParams `json:"Response"`
}

func (r *DescribeDosageDetailByDateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDosageDetailByDateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSavingPlanCoverageRequestParams struct {
	// 费用起始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 费用结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，以此类推
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 数量，最大值为200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 取值包括1（缺省值）和2，1表示按天统计覆盖率，2表示按月统计覆盖率，此参数仅影响返回的RateSet聚合粒度，不影响返回的DetailSet
	PeriodType *uint64 `json:"PeriodType,omitnil" name:"PeriodType"`
}

type DescribeSavingPlanCoverageRequest struct {
	*tchttp.BaseRequest
	
	// 费用起始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 费用结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，以此类推
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 数量，最大值为200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 取值包括1（缺省值）和2，1表示按天统计覆盖率，2表示按月统计覆盖率，此参数仅影响返回的RateSet聚合粒度，不影响返回的DetailSet
	PeriodType *uint64 `json:"PeriodType,omitnil" name:"PeriodType"`
}

func (r *DescribeSavingPlanCoverageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSavingPlanCoverageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PeriodType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSavingPlanCoverageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSavingPlanCoverageResponseParams struct {
	// 节省计划覆盖率明细数据
	DetailSet []*SavingPlanCoverageDetail `json:"DetailSet,omitnil" name:"DetailSet"`

	// 节省计划覆盖率聚合数据
	RateSet []*SavingPlanCoverageRate `json:"RateSet,omitnil" name:"RateSet"`

	// 查询命中的节省计划覆盖率明细数据总条数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSavingPlanCoverageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSavingPlanCoverageResponseParams `json:"Response"`
}

func (r *DescribeSavingPlanCoverageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSavingPlanCoverageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSavingPlanOverviewRequestParams struct {
	// 开始时间，格式yyyy-MM-dd 注：查询范围请勿超过6个月
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 结束时间，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页数量，最大值为200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeSavingPlanOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，格式yyyy-MM-dd 注：查询范围请勿超过6个月
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 结束时间，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页数量，最大值为200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeSavingPlanOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSavingPlanOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSavingPlanOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSavingPlanOverviewResponseParams struct {
	// 节省计划总览明细数据	
	Overviews []*SavingPlanOverviewDetail `json:"Overviews,omitnil" name:"Overviews"`

	// 查询命中的节省计划总览明细数据总条数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSavingPlanOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSavingPlanOverviewResponseParams `json:"Response"`
}

func (r *DescribeSavingPlanOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSavingPlanOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSavingPlanUsageRequestParams struct {
	// 开始时间，格式yyyy-MM-dd 注：查询范围请勿超过6个月
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 结束时间，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页数量，最大值为200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 查询结果数据的时间间隔
	TimeInterval *string `json:"TimeInterval,omitnil" name:"TimeInterval"`
}

type DescribeSavingPlanUsageRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，格式yyyy-MM-dd 注：查询范围请勿超过6个月
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 结束时间，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 每页数量，最大值为200
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 查询结果数据的时间间隔
	TimeInterval *string `json:"TimeInterval,omitnil" name:"TimeInterval"`
}

func (r *DescribeSavingPlanUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSavingPlanUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TimeInterval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSavingPlanUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSavingPlanUsageResponseParams struct {
	// 节省计划使用率数据
	Usages []*SavingPlanUsageDetail `json:"Usages,omitnil" name:"Usages"`

	// 查询命中的节省计划总览明细数据总条数
	Total *uint64 `json:"Total,omitnil" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSavingPlanUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSavingPlanUsageResponseParams `json:"Response"`
}

func (r *DescribeSavingPlanUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSavingPlanUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagListRequestParams struct {
	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，依次类推
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分账标签键，用作模糊搜索
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签类型，枚举值：0普通标签，1分账标签，用作筛选，不传获取全部标签键
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 排序方式，枚举值：asc排升序，desc排降序
	OrderType *string `json:"OrderType,omitnil" name:"OrderType"`
}

type DescribeTagListRequest struct {
	*tchttp.BaseRequest
	
	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，依次类推
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 分账标签键，用作模糊搜索
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签类型，枚举值：0普通标签，1分账标签，用作筛选，不传获取全部标签键
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 排序方式，枚举值：asc排升序，desc排降序
	OrderType *string `json:"OrderType,omitnil" name:"OrderType"`
}

func (r *DescribeTagListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "TagKey")
	delete(f, "Status")
	delete(f, "OrderType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagListResponseParams struct {
	// 总记录数
	RecordNum *uint64 `json:"RecordNum,omitnil" name:"RecordNum"`

	// 标签信息
	Data []*TagDataInfo `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTagListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagListResponseParams `json:"Response"`
}

func (r *DescribeTagListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVoucherInfoRequestParams struct {
	// 一页多少条数据，默认是20条，最大不超过1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 第多少页，默认是1
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 券状态：待使用：unUsed，已使用： used，已发货：delivered，已作废： cancel，已过期：overdue
	Status *string `json:"Status,omitnil" name:"Status"`

	// 代金券id
	VoucherId *string `json:"VoucherId,omitnil" name:"VoucherId"`

	// 代金券订单id
	CodeId *string `json:"CodeId,omitnil" name:"CodeId"`

	// 商品码
	ProductCode *string `json:"ProductCode,omitnil" name:"ProductCode"`

	// 活动id
	ActivityId *string `json:"ActivityId,omitnil" name:"ActivityId"`

	// 代金券名称
	VoucherName *string `json:"VoucherName,omitnil" name:"VoucherName"`

	// 发放开始时间,例：2021-01-01
	TimeFrom *string `json:"TimeFrom,omitnil" name:"TimeFrom"`

	// 发放结束时间，例：2021-01-01
	TimeTo *string `json:"TimeTo,omitnil" name:"TimeTo"`

	// 指定排序字段：BeginTime开始时间、EndTime到期时间、CreateTime创建时间
	SortField *string `json:"SortField,omitnil" name:"SortField"`

	// 指定升序降序：desc、asc
	SortOrder *string `json:"SortOrder,omitnil" name:"SortOrder"`

	// 付费模式，postPay后付费/prePay预付费/riPay预留实例/""或者"*"表示全部模式，如果payMode为""或"*"，那么productCode与subProductCode必须传空
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 付费场景PayMode=postPay时：spotpay-竞价实例,"settle account"-普通后付费PayMode=prePay时：purchase-包年包月新购，renew-包年包月续费（自动续费），modify-包年包月配置变更(变配）PayMode=riPay时：oneOffFee-预留实例预付，hourlyFee-预留实例每小时扣费，*-支持全部付费场景
	PayScene *string `json:"PayScene,omitnil" name:"PayScene"`

	// 操作人，默认就是用户uin
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 代金券主类型 has_price 为有价现金券 no_price 为无价代金券
	VoucherMainType *string `json:"VoucherMainType,omitnil" name:"VoucherMainType"`

	// 代金券副类型 discount 为折扣券 deduct 为抵扣券
	VoucherSubType *string `json:"VoucherSubType,omitnil" name:"VoucherSubType"`
}

type DescribeVoucherInfoRequest struct {
	*tchttp.BaseRequest
	
	// 一页多少条数据，默认是20条，最大不超过1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 第多少页，默认是1
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 券状态：待使用：unUsed，已使用： used，已发货：delivered，已作废： cancel，已过期：overdue
	Status *string `json:"Status,omitnil" name:"Status"`

	// 代金券id
	VoucherId *string `json:"VoucherId,omitnil" name:"VoucherId"`

	// 代金券订单id
	CodeId *string `json:"CodeId,omitnil" name:"CodeId"`

	// 商品码
	ProductCode *string `json:"ProductCode,omitnil" name:"ProductCode"`

	// 活动id
	ActivityId *string `json:"ActivityId,omitnil" name:"ActivityId"`

	// 代金券名称
	VoucherName *string `json:"VoucherName,omitnil" name:"VoucherName"`

	// 发放开始时间,例：2021-01-01
	TimeFrom *string `json:"TimeFrom,omitnil" name:"TimeFrom"`

	// 发放结束时间，例：2021-01-01
	TimeTo *string `json:"TimeTo,omitnil" name:"TimeTo"`

	// 指定排序字段：BeginTime开始时间、EndTime到期时间、CreateTime创建时间
	SortField *string `json:"SortField,omitnil" name:"SortField"`

	// 指定升序降序：desc、asc
	SortOrder *string `json:"SortOrder,omitnil" name:"SortOrder"`

	// 付费模式，postPay后付费/prePay预付费/riPay预留实例/""或者"*"表示全部模式，如果payMode为""或"*"，那么productCode与subProductCode必须传空
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 付费场景PayMode=postPay时：spotpay-竞价实例,"settle account"-普通后付费PayMode=prePay时：purchase-包年包月新购，renew-包年包月续费（自动续费），modify-包年包月配置变更(变配）PayMode=riPay时：oneOffFee-预留实例预付，hourlyFee-预留实例每小时扣费，*-支持全部付费场景
	PayScene *string `json:"PayScene,omitnil" name:"PayScene"`

	// 操作人，默认就是用户uin
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 代金券主类型 has_price 为有价现金券 no_price 为无价代金券
	VoucherMainType *string `json:"VoucherMainType,omitnil" name:"VoucherMainType"`

	// 代金券副类型 discount 为折扣券 deduct 为抵扣券
	VoucherSubType *string `json:"VoucherSubType,omitnil" name:"VoucherSubType"`
}

func (r *DescribeVoucherInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVoucherInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Status")
	delete(f, "VoucherId")
	delete(f, "CodeId")
	delete(f, "ProductCode")
	delete(f, "ActivityId")
	delete(f, "VoucherName")
	delete(f, "TimeFrom")
	delete(f, "TimeTo")
	delete(f, "SortField")
	delete(f, "SortOrder")
	delete(f, "PayMode")
	delete(f, "PayScene")
	delete(f, "Operator")
	delete(f, "VoucherMainType")
	delete(f, "VoucherSubType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVoucherInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVoucherInfoResponseParams struct {
	// 券总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 总余额（微分）
	TotalBalance *int64 `json:"TotalBalance,omitnil" name:"TotalBalance"`

	// 代金券相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoucherInfos []*VoucherInfos `json:"VoucherInfos,omitnil" name:"VoucherInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVoucherInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVoucherInfoResponseParams `json:"Response"`
}

func (r *DescribeVoucherInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVoucherInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVoucherUsageDetailsRequestParams struct {
	// 一页多少条数据，默认是20条，最大不超过1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 第多少页，默认是1
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 代金券id
	VoucherId *string `json:"VoucherId,omitnil" name:"VoucherId"`

	// 操作人，默认就是用户uin
	Operator *string `json:"Operator,omitnil" name:"Operator"`
}

type DescribeVoucherUsageDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 一页多少条数据，默认是20条，最大不超过1000
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 第多少页，默认是1
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 代金券id
	VoucherId *string `json:"VoucherId,omitnil" name:"VoucherId"`

	// 操作人，默认就是用户uin
	Operator *string `json:"Operator,omitnil" name:"Operator"`
}

func (r *DescribeVoucherUsageDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVoucherUsageDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "VoucherId")
	delete(f, "Operator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVoucherUsageDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVoucherUsageDetailsResponseParams struct {
	// 券总数
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 总已用金额（微分）
	TotalUsedAmount *int64 `json:"TotalUsedAmount,omitnil" name:"TotalUsedAmount"`

	// 代金券使用记录细节
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsageRecords []*UsageRecords `json:"UsageRecords,omitnil" name:"UsageRecords"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeVoucherUsageDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVoucherUsageDetailsResponseParams `json:"Response"`
}

func (r *DescribeVoucherUsageDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVoucherUsageDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailPoint struct {
	// 时间
	Time *string `json:"Time,omitnil" name:"Time"`

	// 值
	Value *string `json:"Value,omitnil" name:"Value"`
}

type DetailSet struct {
	// 域名
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 使用数据明细
	DetailPoints []*DetailPoint `json:"DetailPoints,omitnil" name:"DetailPoints"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceID *string `json:"InstanceID,omitnil" name:"InstanceID"`
}

type DistributionBillDetail struct {
	// 产品名称：用户所采购的各类云产品，例如：云服务器 CVM
	BusinessCodeName *string `json:"BusinessCodeName,omitnil" name:"BusinessCodeName"`

	// 子产品名称：用户采购的具体产品细分类型，例如：云服务器 CVM-标准型 S1
	ProductCodeName *string `json:"ProductCodeName,omitnil" name:"ProductCodeName"`

	// 计费模式：资源的计费模式，区分为包年包月和按量计费
	PayModeName *string `json:"PayModeName,omitnil" name:"PayModeName"`

	// 项目名称：资源归属的项目，用户在控制台给资源自主分配项目，未分配则是默认项目
	ProjectName *string `json:"ProjectName,omitnil" name:"ProjectName"`

	// 地域：资源所属地域，如华南地区（广州）
	RegionName *string `json:"RegionName,omitnil" name:"RegionName"`

	// 可用区：资源所属可用区，如广州三区
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// 资源 ID：账单中出账对象 ID，不同产品因资源形态不同，资源内容不完全相同，如云服务器 CVM 为对应的实例 ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 资源别名：用户在控制台为资源设置的名称，如果未设置，则默认为空
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`

	// 交易类型，如包年包月新购、包年包月续费、按量计费扣费等类型
	ActionTypeName *string `json:"ActionTypeName,omitnil" name:"ActionTypeName"`

	// 订单ID：包年包月计费模式下订购的订单号
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// 交易ID：结算扣费单号
	BillId *string `json:"BillId,omitnil" name:"BillId"`

	// 扣费时间：结算扣费时间
	PayTime *string `json:"PayTime,omitnil" name:"PayTime"`

	// 开始使用时间：产品服务开始使用时间
	FeeBeginTime *string `json:"FeeBeginTime,omitnil" name:"FeeBeginTime"`

	// 结束使用时间：产品服务结束使用时间
	FeeEndTime *string `json:"FeeEndTime,omitnil" name:"FeeEndTime"`

	// 组件列表
	ComponentSet []*BillDetailComponent `json:"ComponentSet,omitnil" name:"ComponentSet"`

	// 使用者UIN：实际使用资源的账号 ID
	OwnerUin *string `json:"OwnerUin,omitnil" name:"OwnerUin"`

	// 操作者UIN：操作者账号 ID（预付费资源下单或后付费操作开通资源账号的 ID 或者角色 ID ）
	OperateUin *string `json:"OperateUin,omitnil" name:"OperateUin"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*BillTagInfo `json:"Tags,omitnil" name:"Tags"`

	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessCode *string `json:"BusinessCode,omitnil" name:"BusinessCode"`

	// 子产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitnil" name:"ProductCode"`

	// 交易类型编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionType *string `json:"ActionType,omitnil" name:"ActionType"`

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *string `json:"RegionId,omitnil" name:"RegionId"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil" name:"ProjectId"`

	// 价格属性：该组件除单价、时长外的其他影响折扣定价的属性信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PriceInfo []*string `json:"PriceInfo,omitnil" name:"PriceInfo"`

	// 关联交易单据ID：和本笔交易关联单据 ID，如，冲销订单，记录原订单、重结订单，退费单记录对应的原购买订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociatedOrder *BillDetailAssociatedOrder `json:"AssociatedOrder,omitnil" name:"AssociatedOrder"`

	// 计算说明：特殊交易类型计费结算的详细计算说明，如退费及变配
	// 注意：此字段可能返回 null，表示取不到有效值。
	Formula *string `json:"Formula,omitnil" name:"Formula"`

	// 计费规则：各产品详细的计费规则官网说明链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	FormulaUrl *string `json:"FormulaUrl,omitnil" name:"FormulaUrl"`

	// 账单归属月
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillMonth *string `json:"BillMonth,omitnil" name:"BillMonth"`

	// 账单归属日
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillDay *string `json:"BillDay,omitnil" name:"BillDay"`
}

type ExcludedProducts struct {
	// 不适用商品名称
	GoodsName *string `json:"GoodsName,omitnil" name:"GoodsName"`

	// postPay后付费/prePay预付费/riPay预留实例/空字符串或者"*"表示全部模式。
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`
}

// Predefined struct for user
type PayDealsRequestParams struct {
	// 需要支付的一个或者多个子订单号，与BigDealIds字段两者必须且仅传一个参数
	OrderIds []*string `json:"OrderIds,omitnil" name:"OrderIds"`

	// 是否自动使用代金券,1:是,0否,默认0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil" name:"AutoVoucher"`

	// 代金券ID列表,目前仅支持指定一张代金券
	VoucherIds []*string `json:"VoucherIds,omitnil" name:"VoucherIds"`

	// 需要支付的一个或者多个大订单号，与OrderIds字段两者必须且仅传一个参数
	BigDealIds []*string `json:"BigDealIds,omitnil" name:"BigDealIds"`
}

type PayDealsRequest struct {
	*tchttp.BaseRequest
	
	// 需要支付的一个或者多个子订单号，与BigDealIds字段两者必须且仅传一个参数
	OrderIds []*string `json:"OrderIds,omitnil" name:"OrderIds"`

	// 是否自动使用代金券,1:是,0否,默认0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil" name:"AutoVoucher"`

	// 代金券ID列表,目前仅支持指定一张代金券
	VoucherIds []*string `json:"VoucherIds,omitnil" name:"VoucherIds"`

	// 需要支付的一个或者多个大订单号，与OrderIds字段两者必须且仅传一个参数
	BigDealIds []*string `json:"BigDealIds,omitnil" name:"BigDealIds"`
}

func (r *PayDealsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PayDealsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderIds")
	delete(f, "AutoVoucher")
	delete(f, "VoucherIds")
	delete(f, "BigDealIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PayDealsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PayDealsResponseParams struct {
	// 此次操作支付成功的子订单号数组
	OrderIds []*string `json:"OrderIds,omitnil" name:"OrderIds"`

	// 此次操作支付成功的资源Id数组
	ResourceIds []*string `json:"ResourceIds,omitnil" name:"ResourceIds"`

	// 此次操作支付成功的大订单号数组
	BigDealIds []*string `json:"BigDealIds,omitnil" name:"BigDealIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PayDealsResponse struct {
	*tchttp.BaseResponse
	Response *PayDealsResponseParams `json:"Response"`
}

func (r *PayDealsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PayDealsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PayModeSummaryOverviewItem struct {
	// 计费模式编码
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 计费模式：区分为包年包月和按量计费
	PayModeName *string `json:"PayModeName,omitnil" name:"PayModeName"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil" name:"RealTotalCostRatio"`

	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil" name:"IncentivePayAmount"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	TransferPayAmount *string `json:"TransferPayAmount,omitnil" name:"TransferPayAmount"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitnil" name:"TotalCost"`

	// 按交易类型汇总消费详情
	Detail []*ActionSummaryOverviewItem `json:"Detail,omitnil" name:"Detail"`
}

type ProductInfo struct {
	// 商品详情名称标识
	Name *string `json:"Name,omitnil" name:"Name"`

	// 商品详情
	Value *string `json:"Value,omitnil" name:"Value"`
}

type ProjectSummaryOverviewItem struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil" name:"ProjectId"`

	// 项目名称：资源归属的项目，用户在控制台给资源自主分配项目，未分配则是默认项目
	ProjectName *string `json:"ProjectName,omitnil" name:"ProjectName"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil" name:"RealTotalCostRatio"`

	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil" name:"IncentivePayAmount"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	TransferPayAmount *string `json:"TransferPayAmount,omitnil" name:"TransferPayAmount"`

	// 账单月份，格式2019-08
	BillMonth *string `json:"BillMonth,omitnil" name:"BillMonth"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitnil" name:"TotalCost"`
}

type RegionSummaryOverviewItem struct {
	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *string `json:"RegionId,omitnil" name:"RegionId"`

	// 地域名称：资源所属地域，例如华南地区（广州）
	RegionName *string `json:"RegionName,omitnil" name:"RegionName"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil" name:"RealTotalCostRatio"`

	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil" name:"IncentivePayAmount"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	TransferPayAmount *string `json:"TransferPayAmount,omitnil" name:"TransferPayAmount"`

	// 账单月份，格式2019-08
	BillMonth *string `json:"BillMonth,omitnil" name:"BillMonth"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitnil" name:"TotalCost"`
}

type SavingPlanCoverageDetail struct {
	// 资源 ID：账单中出账对象 ID，不同产品因资源形态不同，资源内容不完全相同，如云服务器 CVM 为对应的实例 ID
	ResourceId *string `json:"ResourceId,omitnil" name:"ResourceId"`

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *uint64 `json:"RegionId,omitnil" name:"RegionId"`

	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitnil" name:"ProductCode"`

	// 子产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductCode *string `json:"SubProductCode,omitnil" name:"SubProductCode"`

	// 费用起始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitnil" name:"StartDate"`

	// 费用结束日期，格式yyyy-MM-dd，目前与StartDate相等
	EndDate *string `json:"EndDate,omitnil" name:"EndDate"`

	// 节省计划覆盖金额（即节省计划支付金额）
	SpCoveredAmount *float64 `json:"SpCoveredAmount,omitnil" name:"SpCoveredAmount"`

	// 节省计划未覆盖金额（即优惠后总价）
	SpUncoveredAmount *float64 `json:"SpUncoveredAmount,omitnil" name:"SpUncoveredAmount"`

	// 总支出（即节省计划未覆盖金额 + 节省计划覆盖金额）
	TotalRealAmount *float64 `json:"TotalRealAmount,omitnil" name:"TotalRealAmount"`

	// 按量计费预期金额（即折前价 * 折扣）
	ExpectedAmount *float64 `json:"ExpectedAmount,omitnil" name:"ExpectedAmount"`

	// 覆盖率结果，取值[0, 100]
	SpCoverage *float64 `json:"SpCoverage,omitnil" name:"SpCoverage"`
}

type SavingPlanCoverageRate struct {
	// 聚合时间维度，按天聚合格式为yyyy-MM-dd，按月聚合格式为yyyy-MM
	DatePoint *string `json:"DatePoint,omitnil" name:"DatePoint"`

	// 覆盖率结果，取值[0, 100]
	Rate *float64 `json:"Rate,omitnil" name:"Rate"`
}

type SavingPlanOverviewDetail struct {
	// 节省计划类型
	SpType *string `json:"SpType,omitnil" name:"SpType"`

	// 支付类型
	PayType *uint64 `json:"PayType,omitnil" name:"PayType"`

	// 支付金额（单位：元）
	PayAmount *string `json:"PayAmount,omitnil" name:"PayAmount"`

	// 开始时间 yyyy-mm-dd HH:mm:ss格式
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间 yyyy-mm-dd HH:mm:ss格式
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 购买时间 yyyy-mm-dd HH:mm:ss格式
	BuyTime *string `json:"BuyTime,omitnil" name:"BuyTime"`

	// 状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 累计节省金额（单位：元）
	SavingAmount *string `json:"SavingAmount,omitnil" name:"SavingAmount"`

	// 地域
	Region []*string `json:"Region,omitnil" name:"Region"`
}

type SavingPlanUsageDetail struct {
	// 节省计划类型
	SpType *string `json:"SpType,omitnil" name:"SpType"`

	// 节省计划状态
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// 累计抵扣的金额（单位：元）
	DeductAmount *string `json:"DeductAmount,omitnil" name:"DeductAmount"`

	// 累计承诺消费金额（单位：元）
	PromiseAmount *string `json:"PromiseAmount,omitnil" name:"PromiseAmount"`

	// 累计净节省金额（单位：元）
	NetSavings *string `json:"NetSavings,omitnil" name:"NetSavings"`

	// 使用率
	UtilizationRate *float64 `json:"UtilizationRate,omitnil" name:"UtilizationRate"`

	// 累计流失金额（单位：元）
	LossAmount *string `json:"LossAmount,omitnil" name:"LossAmount"`

	// 累计按量计费预期金额（单位：元）
	DosageAmount *string `json:"DosageAmount,omitnil" name:"DosageAmount"`

	// 累计成本金额（单位：元）
	CostAmount *string `json:"CostAmount,omitnil" name:"CostAmount"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region []*string `json:"Region,omitnil" name:"Region"`
}

type SummaryDetail struct {
	// 账单维度编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupKey *string `json:"GroupKey,omitnil" name:"GroupKey"`

	// 账单维度值
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupValue *string `json:"GroupValue,omitnil" name:"GroupValue"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitnil" name:"TotalCost"`

	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil" name:"IncentivePayAmount"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil" name:"TransferPayAmount"`

	// 产品汇总信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Business []*BusinessSummaryInfo `json:"Business,omitnil" name:"Business"`
}

type SummaryTotal struct {
	// 优惠后总价
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalCost *string `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCost *string `json:"TotalCost,omitnil" name:"TotalCost"`
}

type TagDataInfo struct {
	// 分账标签键
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签类型，0普通标签，1分账标签
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 设置分账标签时间，普通标签不返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type TagSummaryOverviewItem struct {
	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`

	// 费用所占百分比，两位小数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil" name:"RealTotalCostRatio"`

	// 优惠后总价
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalCost *string `json:"RealTotalCost,omitnil" name:"RealTotalCost"`

	// 现金账户支出：通过现金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	CashPayAmount *string `json:"CashPayAmount,omitnil" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil" name:"IncentivePayAmount"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil" name:"VoucherPayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil" name:"TransferPayAmount"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCost *string `json:"TotalCost,omitnil" name:"TotalCost"`
}

type UsageDetails struct {
	// 商品名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitnil" name:"ProductName"`

	// 商品细节
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductName *string `json:"SubProductName,omitnil" name:"SubProductName"`
}

type UsageRecords struct {
	// 使用金额（微分）
	UsedAmount *int64 `json:"UsedAmount,omitnil" name:"UsedAmount"`

	// 使用时间
	UsedTime *string `json:"UsedTime,omitnil" name:"UsedTime"`

	// 使用记录细节
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsageDetails []*UsageDetails `json:"UsageDetails,omitnil" name:"UsageDetails"`
}

type VoucherInfos struct {
	// 代金券拥有者
	OwnerUin *string `json:"OwnerUin,omitnil" name:"OwnerUin"`

	// 券状态：待使用：unUsed，已使用： used，已发货：delivered，已作废： cancel，已过期：overdue
	Status *string `json:"Status,omitnil" name:"Status"`

	// 代金券面额（微分）
	NominalValue *int64 `json:"NominalValue,omitnil" name:"NominalValue"`

	// 剩余金额（微分）
	Balance *int64 `json:"Balance,omitnil" name:"Balance"`

	// 代金券id
	VoucherId *string `json:"VoucherId,omitnil" name:"VoucherId"`

	// postPay后付费/prePay预付费/riPay预留实例/空字符串或者'*'表示全部模式
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 付费场景PayMode=postPay时：spotpay-竞价实例,"settle account"-普通后付费PayMode=prePay时：purchase-包年包月新购，renew-包年包月续费（自动续费），modify-包年包月配置变更(变配）PayMode=riPay时：oneOffFee-预留实例预付，hourlyFee-预留实例每小时扣费，*-支持全部付费场景
	PayScene *string `json:"PayScene,omitnil" name:"PayScene"`

	// 有效期生效时间
	BeginTime *string `json:"BeginTime,omitnil" name:"BeginTime"`

	// 有效期截止时间
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 适用商品信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicableProducts *ApplicableProducts `json:"ApplicableProducts,omitnil" name:"ApplicableProducts"`

	// 不适用商品信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludedProducts []*ExcludedProducts `json:"ExcludedProducts,omitnil" name:"ExcludedProducts"`
}