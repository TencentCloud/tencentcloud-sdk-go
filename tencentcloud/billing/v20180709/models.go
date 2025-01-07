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
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 交易类型：如包年包月新购、包年包月续费、按量计费扣费等类型
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil,omitempty" name:"RealTotalCostRatio"`

	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// 账单月份，格式2019-08
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type AdjustInfoDetail struct {
	// 支付者UIN：支付者的账号 ID，账号 ID 是用户在腾讯云的唯一账号标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 账单月份，格式：yyyy-MM
	// 注意：此字段可能返回 null，表示取不到有效值。
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 调整类型
	// 调账：manualAdjustment
	// 补结算：supplementarySettlement
	// 重结算：reSettlement
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdjustType *string `json:"AdjustType,omitnil,omitempty" name:"AdjustType"`

	// 调整单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdjustNum *string `json:"AdjustNum,omitnil,omitempty" name:"AdjustNum"`

	// 异常调整完成时间，格式：yyyy-MM-dd HH:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdjustCompletionTime *string `json:"AdjustCompletionTime,omitnil,omitempty" name:"AdjustCompletionTime"`

	// 调整金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdjustAmount *float64 `json:"AdjustAmount,omitnil,omitempty" name:"AdjustAmount"`
}

type AllocationAverageData struct {
	// 起始月份
	BeginMonth *string `json:"BeginMonth,omitnil,omitempty" name:"BeginMonth"`

	// 结束月份
	EndMonth *string `json:"EndMonth,omitnil,omitempty" name:"EndMonth"`

	// 合计费用(折后总额)平均值
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`
}

type AllocationBillTrendDetail struct {
	// 账单月份
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 账单月份展示名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 合计费用(折后总额)：分账单元总费用，归集费用(折后总额) + 分摊费用(折后总额)
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`
}

type AllocationDetail struct {
	// 分账单元唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// 分账单元名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreeNodeUniqKeyName *string `json:"TreeNodeUniqKeyName,omitnil,omitempty" name:"TreeNodeUniqKeyName"`

	// 日期：结算日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillDate *string `json:"BillDate,omitnil,omitempty" name:"BillDate"`

	// 支付者 UIN：支付者的账号 ID，账号 ID 是用户在腾讯云的唯一账号标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 使用者 UIN：实际使用资源的账号 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 操作者 UIN：操作者账号 ID（预付费资源下单或后付费操作开通资源账号的ID或者角色 ID）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 产品名称：用户所采购的各类云产品
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// 计费模式编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 计费模式：资源的计费模式，区分为包年包月和按量计费
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称：资源归属的项目，用户在控制台给资源自主分配项目，未分配则是默认项目
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域名称：资源所属地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 可用区：资源所属可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 资源ID：不同产品因资源形态不同，资源内容不完全相同，如云服务器 CVM 为对应的实例 ID； 若该产品被分拆，则展示产品分拆后的分拆项 ID，如 COS 桶 ID，CDN 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 实例名称：用户在控制台为资源设置的名称，如未设置默认为空；若该产品被分拆，则展示分拆产品分拆后的分拆项资源别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 实例类型编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例类型：购买的产品服务对应的实例类型，包括资源包、RI、SP、竞价实例。常规实例默认展示“-”
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTypeName *string `json:"InstanceTypeName,omitnil,omitempty" name:"InstanceTypeName"`

	// 分拆项 ID：涉及分拆产品的分拆后的分拆项 ID，如 COS 桶 ID，CDN 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: SplitItemId is deprecated.
	SplitItemId *string `json:"SplitItemId,omitnil,omitempty" name:"SplitItemId"`

	// 分拆项名称：涉及分拆产品的分拆后的分拆项
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: SplitItemName is deprecated.
	SplitItemName *string `json:"SplitItemName,omitnil,omitempty" name:"SplitItemName"`

	// 子产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 子产品名称：用户采购的具体产品细分类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// 交易类型编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 交易类型：明细交易类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// 订单 ID：包年包月计费模式下订购的订单号
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 交易 ID：结算扣费单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// 扣费时间：结算扣费时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayTime *string `json:"PayTime,omitnil,omitempty" name:"PayTime"`

	// 开始使用时间：产品服务开始使用时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeeBeginTime *string `json:"FeeBeginTime,omitnil,omitempty" name:"FeeBeginTime"`

	// 结束使用时间：产品服务结束使用时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeeEndTime *string `json:"FeeEndTime,omitnil,omitempty" name:"FeeEndTime"`

	// 组件类型编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentCode *string `json:"ComponentCode,omitnil,omitempty" name:"ComponentCode"`

	// 组件类型：用户购买的产品或服务对应的组件大类
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentCodeName *string `json:"ComponentCodeName,omitnil,omitempty" name:"ComponentCodeName"`

	// 组件刊例价：组件的官网原始单价（如客户享受一口价/合同价则默认不展示）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SinglePrice *string `json:"SinglePrice,omitnil,omitempty" name:"SinglePrice"`

	// 组件单价：组件的折后单价，组件单价 = 刊例价 * 折扣
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractPrice *string `json:"ContractPrice,omitnil,omitempty" name:"ContractPrice"`

	// 组件价格单位：组件价格的单位，单位构成：元/用量单位/时长单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	SinglePriceUnit *string `json:"SinglePriceUnit,omitnil,omitempty" name:"SinglePriceUnit"`

	// 组件用量：该组件实际结算用量，组件用量=组件原始用量-抵扣用量（含资源包）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedAmount *string `json:"UsedAmount,omitnil,omitempty" name:"UsedAmount"`

	// 组件用量单位：组件用量对应的单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedAmountUnit *string `json:"UsedAmountUnit,omitnil,omitempty" name:"UsedAmountUnit"`

	// 使用时长：资源使用的时长，组件用量=组件原始使用时长-抵扣时长（含资源包）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSpan *string `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 时长单位：资源使用时长的单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 备注属性（实例配置）：其他备注信息，如预留实例的预留实例类型和交易类型、CCN 产品的两端地域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReserveDetail *string `json:"ReserveDetail,omitnil,omitempty" name:"ReserveDetail"`

	// 分拆项用量/时长占比：分拆项用量（时长）占比，分拆项用量（时长）/ 拆分前合计用量（时长）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SplitRatio *string `json:"SplitRatio,omitnil,omitempty" name:"SplitRatio"`

	// 组件原价：原价 = 组件刊例价 * 组件用量 * 使用时长（如客户享受一口价/合同价则默认不展示，退费类场景也默认不展示），指定价模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 预留实例抵扣时长：本产品或服务使用预留实例抵扣的使用时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	RITimeSpan *string `json:"RITimeSpan,omitnil,omitempty" name:"RITimeSpan"`

	// 预留实例抵扣原价：本产品或服务使用预留实例抵扣的组件原价金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	RICost *string `json:"RICost,omitnil,omitempty" name:"RICost"`

	// 节省计划抵扣原价：节省计划抵扣原价 = 节省计划包抵扣面值 / 节省计划抵扣率
	// 注意：此字段可能返回 null，表示取不到有效值。
	SPCost *string `json:"SPCost,omitnil,omitempty" name:"SPCost"`

	// 折扣率：本资源享受的折扣率（如客户享受一口价/合同价则默认不展示，退费场景也默认不展示）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Discount *string `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 混合折扣率：综合各类折扣抵扣信息后的最终折扣率，混合折扣率=优惠后总价/原价
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlendedDiscount *string `json:"BlendedDiscount,omitnil,omitempty" name:"BlendedDiscount"`

	// 优惠后总价：优惠后总价 =（原价 - 预留实例抵扣原价 - 节省计划抵扣原价）* 折扣率
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 现金账户支出(元)：通过现金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 代金券支出(元)：使用各类优惠券（如代金券、现金券等）支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 赠送账户支出(元)：使用赠送金支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 分成账户支出(元)：通过分成金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// 分账标签：资源绑定的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*BillTag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 国内国际编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionType *string `json:"RegionType,omitnil,omitempty" name:"RegionType"`

	// 国内国际：资源所属区域类型（国内、国际）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionTypeName *string `json:"RegionTypeName,omitnil,omitempty" name:"RegionTypeName"`

	// 组件名称编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemCode *string `json:"ItemCode,omitnil,omitempty" name:"ItemCode"`

	// 组件名称：用户购买的产品或服务，所包含的具体组件
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemCodeName *string `json:"ItemCodeName,omitnil,omitempty" name:"ItemCodeName"`

	// 关联单据ID：和本笔交易关联单据ID，如退费订单对应的原新购订单等
	AssociatedOrder *string `json:"AssociatedOrder,omitnil,omitempty" name:"AssociatedOrder"`

	// 价格属性：该组件除单价、时长外的其他影响折扣定价的属性信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PriceInfo []*string `json:"PriceInfo,omitnil,omitempty" name:"PriceInfo"`

	// 计算规则说明：特殊交易类型计费结算的详细计算说明，如退费及变配
	// 注意：此字段可能返回 null，表示取不到有效值。
	Formula *string `json:"Formula,omitnil,omitempty" name:"Formula"`

	// 计费规则：各产品详细的计费规则官网说明链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	FormulaUrl *string `json:"FormulaUrl,omitnil,omitempty" name:"FormulaUrl"`

	// 原始用量/时长：组件被资源包抵扣前的原始用量
	// （目前仅实时音视频、弹性微服务、云呼叫中心及专属可用区产品支持该信息外显，其他产品尚在接入中）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalMeasure *string `json:"RealTotalMeasure,omitnil,omitempty" name:"RealTotalMeasure"`

	// 抵扣用量/时长（含资源包）：组件被资源包抵扣的用量
	// （目前仅实时音视频、弹性微服务、云呼叫中心及专属可用区产品支持该信息外显，其他产品尚在接入中）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeductedMeasure *string `json:"DeductedMeasure,omitnil,omitempty" name:"DeductedMeasure"`

	// 配置描述：资源配置规格信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentConfig *string `json:"ComponentConfig,omitnil,omitempty" name:"ComponentConfig"`

	// 费用归集类型：费用来源类型，分摊、归集、未分配
	// 0 - 分摊
	// 1 - 归集
	// 2 - 未分配
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllocationType *uint64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`

	// 当前消费项的优惠对象，例如：官网折扣、用户折扣、活动折扣。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountObject *string `json:"DiscountObject,omitnil,omitempty" name:"DiscountObject"`

	// 当前消费项的优惠类型，例如：折扣、合同价。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountType *string `json:"DiscountType,omitnil,omitempty" name:"DiscountType"`

	// 对优惠类型的补充描述，例如：商务折扣8折，则优惠类型为“折扣”，优惠内容为“0.8”。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountContent *string `json:"DiscountContent,omitnil,omitempty" name:"DiscountContent"`

	// SPDeduction
	// 注意：此字段可能返回 null，表示取不到有效值。
	SPDeduction *string `json:"SPDeduction,omitnil,omitempty" name:"SPDeduction"`

	// SPDeduction
	// 注意：此字段可能返回 null，表示取不到有效值。
	SPDeductionRate *string `json:"SPDeductionRate,omitnil,omitempty" name:"SPDeductionRate"`

	// 账单月
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`
}

type AllocationMonthOverviewDetail struct {
	// 归集费用(现金)：基于归集规则直接归集到分账单元的现金
	GatherCashPayAmount *string `json:"GatherCashPayAmount,omitnil,omitempty" name:"GatherCashPayAmount"`

	// 归集费用(优惠券)：基于归集规则直接归集到分账单元的资源优惠券
	GatherVoucherPayAmount *string `json:"GatherVoucherPayAmount,omitnil,omitempty" name:"GatherVoucherPayAmount"`

	// 归集费用(赠送金)：基于归集规则直接归集到分账单元的资源赠送金
	GatherIncentivePayAmount *string `json:"GatherIncentivePayAmount,omitnil,omitempty" name:"GatherIncentivePayAmount"`

	// 归集费用(分成金)：基于归集规则直接归集到分账单元的资源分成金
	GatherTransferPayAmount *string `json:"GatherTransferPayAmount,omitnil,omitempty" name:"GatherTransferPayAmount"`

	// 分摊费用(现金)：基于分摊规则分摊到分账单元的资源现金
	AllocateCashPayAmount *string `json:"AllocateCashPayAmount,omitnil,omitempty" name:"AllocateCashPayAmount"`

	// 分摊费用(优惠券)：基于分摊规则分摊到分账单元的资源优惠券
	AllocateVoucherPayAmount *string `json:"AllocateVoucherPayAmount,omitnil,omitempty" name:"AllocateVoucherPayAmount"`

	// 分摊费用(赠送金)：基于分摊规则分摊到分账单元的资源赠送金
	AllocateIncentivePayAmount *string `json:"AllocateIncentivePayAmount,omitnil,omitempty" name:"AllocateIncentivePayAmount"`

	// 分摊费用(分成金)：基于分摊规则分摊到分账单元的资源分成金
	AllocateTransferPayAmount *string `json:"AllocateTransferPayAmount,omitnil,omitempty" name:"AllocateTransferPayAmount"`

	// 合计费用(现金)：分账单元总费用，归集费用(现金) + 分摊费用(现金)
	TotalCashPayAmount *string `json:"TotalCashPayAmount,omitnil,omitempty" name:"TotalCashPayAmount"`

	// 合计费用(优惠券)：分账单元总费用，归集费用(优惠券) + 分摊费用(优惠券)
	TotalVoucherPayAmount *string `json:"TotalVoucherPayAmount,omitnil,omitempty" name:"TotalVoucherPayAmount"`

	// 合计费用(赠送金)：分账单元总费用，归集费用(赠送金) + 分摊费用(赠送金)
	TotalIncentivePayAmount *string `json:"TotalIncentivePayAmount,omitnil,omitempty" name:"TotalIncentivePayAmount"`

	// 合计费用(分成金)：分账单元总费用，归集费用(分成金)+分摊费用(分成金)
	TotalTransferPayAmount *string `json:"TotalTransferPayAmount,omitnil,omitempty" name:"TotalTransferPayAmount"`

	// 归集费用(折后总额)：基于归集规则直接归集到分账单元的资源优惠后总价
	GatherRealCost *string `json:"GatherRealCost,omitnil,omitempty" name:"GatherRealCost"`

	// 分摊费用(折后总额)：基于分摊规则分摊到分账单元的资源优惠后总价
	AllocateRealCost *string `json:"AllocateRealCost,omitnil,omitempty" name:"AllocateRealCost"`

	// 合计费用(折后总额)：分账单元总费用，归集费用(折后总额) + 分摊费用(折后总额)
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 占比(折后总额)：本分账单元合计费用(折后总额)/合计费用(折后总额)*100%
	Ratio *string `json:"Ratio,omitnil,omitempty" name:"Ratio"`

	// 环比(折后总额)：[本月分账单元合计费用(折后总额) - 上月分账单元合计费用(折后总额)] / 上月分账单元合计费用(折后总额) * 100%
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trend *string `json:"Trend,omitnil,omitempty" name:"Trend"`

	// 环比箭头
	// upward -上升
	// downward - 下降
	// none - 平稳
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrendType *string `json:"TrendType,omitnil,omitempty" name:"TrendType"`
}

type AllocationOverviewDetail struct {
	// 分账单元唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// 分账单元名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreeNodeUniqKeyName *string `json:"TreeNodeUniqKeyName,omitnil,omitempty" name:"TreeNodeUniqKeyName"`

	// 日期：结算日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillDate *string `json:"BillDate,omitnil,omitempty" name:"BillDate"`

	// 归集费用(现金)：基于归集规则直接归集到分账单元的现金
	GatherCashPayAmount *string `json:"GatherCashPayAmount,omitnil,omitempty" name:"GatherCashPayAmount"`

	// 归集费用(优惠券)：基于归集规则直接归集到分账单元的资源优惠券
	GatherVoucherPayAmount *string `json:"GatherVoucherPayAmount,omitnil,omitempty" name:"GatherVoucherPayAmount"`

	// 归集费用(赠送金)：基于归集规则直接归集到分账单元的资源赠送金
	GatherIncentivePayAmount *string `json:"GatherIncentivePayAmount,omitnil,omitempty" name:"GatherIncentivePayAmount"`

	// 归集费用(分成金)：基于归集规则直接归集到分账单元的资源分成金
	GatherTransferPayAmount *string `json:"GatherTransferPayAmount,omitnil,omitempty" name:"GatherTransferPayAmount"`

	// 分摊费用(现金)：基于分摊规则分摊到分账单元的资源现金
	AllocateCashPayAmount *string `json:"AllocateCashPayAmount,omitnil,omitempty" name:"AllocateCashPayAmount"`

	// 分摊费用(优惠券)：基于分摊规则分摊到分账单元的资源优惠券
	AllocateVoucherPayAmount *string `json:"AllocateVoucherPayAmount,omitnil,omitempty" name:"AllocateVoucherPayAmount"`

	// 分摊费用(赠送金)：基于分摊规则分摊到分账单元的资源赠送金
	AllocateIncentivePayAmount *string `json:"AllocateIncentivePayAmount,omitnil,omitempty" name:"AllocateIncentivePayAmount"`

	// 分摊费用(分成金)：基于分摊规则分摊到分账单元的资源分成金
	AllocateTransferPayAmount *string `json:"AllocateTransferPayAmount,omitnil,omitempty" name:"AllocateTransferPayAmount"`

	// 合计费用(现金)：分账单元总费用，归集费用(现金) + 分摊费用(现金)
	TotalCashPayAmount *string `json:"TotalCashPayAmount,omitnil,omitempty" name:"TotalCashPayAmount"`

	// 合计费用(优惠券)：分账单元总费用，归集费用(优惠券) + 分摊费用(优惠券)
	TotalVoucherPayAmount *string `json:"TotalVoucherPayAmount,omitnil,omitempty" name:"TotalVoucherPayAmount"`

	// 合计费用(赠送金)：分账单元总费用，归集费用(赠送金) + 分摊费用(赠送金)
	TotalIncentivePayAmount *string `json:"TotalIncentivePayAmount,omitnil,omitempty" name:"TotalIncentivePayAmount"`

	// 合计费用(分成金)：分账单元总费用，归集费用(分成金)+分摊费用(分成金)
	TotalTransferPayAmount *string `json:"TotalTransferPayAmount,omitnil,omitempty" name:"TotalTransferPayAmount"`

	// 归集费用(折后总额)：基于归集规则直接归集到分账单元的资源优惠后总价
	GatherRealCost *string `json:"GatherRealCost,omitnil,omitempty" name:"GatherRealCost"`

	// 分摊费用(折后总额)：基于分摊规则分摊到分账单元的资源优惠后总价
	AllocateRealCost *string `json:"AllocateRealCost,omitnil,omitempty" name:"AllocateRealCost"`

	// 合计费用(折后总额)：分账单元总费用，归集费用(折后总额) + 分摊费用(折后总额)
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 占比(折后总额)：本分账单元合计费用(折后总额)/合计费用(折后总额)*100%
	Ratio *string `json:"Ratio,omitnil,omitempty" name:"Ratio"`

	// 环比(折后总额)：[本月分账单元合计费用(折后总额) - 上月分账单元合计费用(折后总额)] / 上月分账单元合计费用(折后总额) * 100%
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trend *string `json:"Trend,omitnil,omitempty" name:"Trend"`

	// 环比箭头
	// upward -上升
	// downward - 下降
	// none - 平稳
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrendType *string `json:"TrendType,omitnil,omitempty" name:"TrendType"`
}

type AllocationOverviewNode struct {
	// 分账单元ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 分账单元名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分账单元唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// 分账单元包含规则标志
	// 0 - 不存在规则
	// 1 - 同时存在归集规则和公摊规则
	// 2 - 仅存在归集规则
	// 3 - 仅存在公摊规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Symbol *uint64 `json:"Symbol,omitnil,omitempty" name:"Symbol"`

	// 子单元月概览详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Children []*AllocationOverviewNode `json:"Children,omitnil,omitempty" name:"Children"`

	// 分账账单月概览金额明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *AllocationMonthOverviewDetail `json:"Detail,omitnil,omitempty" name:"Detail"`
}

type AllocationOverviewTotal struct {
	// 总费用：现金费用合计+分成金费用合计+赠送金费用合计+优惠券费用合计
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 现金： 现金费用合计
	// 注意：此字段可能返回 null，表示取不到有效值。
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 赠送金：赠送金费用合计
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 优惠券：优惠券费用合计
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 分成金：分成金费用合计
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`
}

type AllocationRule struct {
	// 公摊规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 公摊规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

type AllocationStat struct {
	// 费用平均信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Average *AllocationAverageData `json:"Average,omitnil,omitempty" name:"Average"`
}

type AllocationSummaryByBusiness struct {
	// 分账单元唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// 分账单元名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreeNodeUniqKeyName *string `json:"TreeNodeUniqKeyName,omitnil,omitempty" name:"TreeNodeUniqKeyName"`

	// 日期：结算日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillDate *string `json:"BillDate,omitnil,omitempty" name:"BillDate"`

	// 归集费用(现金)：基于归集规则直接归集到分账单元的现金
	GatherCashPayAmount *string `json:"GatherCashPayAmount,omitnil,omitempty" name:"GatherCashPayAmount"`

	// 归集费用(优惠券)：基于归集规则直接归集到分账单元的资源优惠券
	GatherVoucherPayAmount *string `json:"GatherVoucherPayAmount,omitnil,omitempty" name:"GatherVoucherPayAmount"`

	// 归集费用(赠送金)：基于归集规则直接归集到分账单元的资源赠送金
	GatherIncentivePayAmount *string `json:"GatherIncentivePayAmount,omitnil,omitempty" name:"GatherIncentivePayAmount"`

	// 归集费用(分成金)：基于归集规则直接归集到分账单元的资源分成金
	GatherTransferPayAmount *string `json:"GatherTransferPayAmount,omitnil,omitempty" name:"GatherTransferPayAmount"`

	// 分摊费用(现金)：基于分摊规则分摊到分账单元的资源现金
	AllocateCashPayAmount *string `json:"AllocateCashPayAmount,omitnil,omitempty" name:"AllocateCashPayAmount"`

	// 分摊费用(优惠券)：基于分摊规则分摊到分账单元的资源优惠券
	AllocateVoucherPayAmount *string `json:"AllocateVoucherPayAmount,omitnil,omitempty" name:"AllocateVoucherPayAmount"`

	// 分摊费用(赠送金)：基于分摊规则分摊到分账单元的资源赠送金
	AllocateIncentivePayAmount *string `json:"AllocateIncentivePayAmount,omitnil,omitempty" name:"AllocateIncentivePayAmount"`

	// 分摊费用(分成金)：基于分摊规则分摊到分账单元的资源分成金
	AllocateTransferPayAmount *string `json:"AllocateTransferPayAmount,omitnil,omitempty" name:"AllocateTransferPayAmount"`

	// 合计费用(现金)：分账单元总费用，归集费用(现金) + 分摊费用(现金)
	TotalCashPayAmount *string `json:"TotalCashPayAmount,omitnil,omitempty" name:"TotalCashPayAmount"`

	// 合计费用(优惠券)：分账单元总费用，归集费用(优惠券) + 分摊费用(优惠券)
	TotalVoucherPayAmount *string `json:"TotalVoucherPayAmount,omitnil,omitempty" name:"TotalVoucherPayAmount"`

	// 合计费用(赠送金)：分账单元总费用，归集费用(赠送金) + 分摊费用(赠送金)
	TotalIncentivePayAmount *string `json:"TotalIncentivePayAmount,omitnil,omitempty" name:"TotalIncentivePayAmount"`

	// 合计费用(分成金)：分账单元总费用，归集费用(分成金)+分摊费用(分成金)
	TotalTransferPayAmount *string `json:"TotalTransferPayAmount,omitnil,omitempty" name:"TotalTransferPayAmount"`

	// 归集费用(折后总额)：基于归集规则直接归集到分账单元的资源优惠后总价
	GatherRealCost *string `json:"GatherRealCost,omitnil,omitempty" name:"GatherRealCost"`

	// 分摊费用(折后总额)：基于分摊规则分摊到分账单元的资源优惠后总价
	AllocateRealCost *string `json:"AllocateRealCost,omitnil,omitempty" name:"AllocateRealCost"`

	// 合计费用(折后总额)：分账单元总费用，归集费用(折后总额) + 分摊费用(折后总额)
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 占比(折后总额)：本分账单元合计费用(折后总额)/合计费用(折后总额)*100%
	Ratio *string `json:"Ratio,omitnil,omitempty" name:"Ratio"`

	// 环比(折后总额)：[本月分账单元合计费用(折后总额) - 上月分账单元合计费用(折后总额)] / 上月分账单元合计费用(折后总额) * 100%
	// 注意：此字段可能返回 null，表示取不到有效值。
	Trend *string `json:"Trend,omitnil,omitempty" name:"Trend"`

	// 环比箭头
	// upward -上升
	// downward - 下降
	// none - 平稳
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrendType *string `json:"TrendType,omitnil,omitempty" name:"TrendType"`

	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 产品名称：用户所采购的各类云产品
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// 组件原价：原价 = 组件刊例价 * 组件用量 * 使用时长（如客户享受一口价/合同价则默认不展示，退费类场景也默认不展示），指定价模式
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 预留实例抵扣原价：本产品或服务使用预留实例抵扣的组件原价金额
	RICost *string `json:"RICost,omitnil,omitempty" name:"RICost"`

	// 节省计划抵扣原价：节省计划抵扣原价 = 节省计划包抵扣面值 / 节省计划抵扣率
	SPCost *string `json:"SPCost,omitnil,omitempty" name:"SPCost"`

	// 现金账户支出(元)：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 代金券支出(元)：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 赠送账户支出(元)：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 分成账户支出(元)：通过分成金账户支付的金额
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// 优惠后总价：优惠后总价 =（原价 - 预留实例抵扣原价 - 节省计划抵扣原价）* 折扣率
	AllocationRealTotalCost *string `json:"AllocationRealTotalCost,omitnil,omitempty" name:"AllocationRealTotalCost"`
}

type AllocationSummaryByItem struct {
	// 分账单元唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// 分账单元名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreeNodeUniqKeyName *string `json:"TreeNodeUniqKeyName,omitnil,omitempty" name:"TreeNodeUniqKeyName"`

	// 日期：结算日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillDate *string `json:"BillDate,omitnil,omitempty" name:"BillDate"`

	// 支付者 UIN：支付者的账号 ID，账号 ID 是用户在腾讯云的唯一账号标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 使用者 UIN：实际使用资源的账号 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 操作者 UIN：操作者账号 ID（预付费资源下单或后付费操作开通资源账号的ID或者角色 ID）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// 计费模式编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 计费模式：资源的计费模式，区分为包年包月和按量计费
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// 交易类型编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 交易类型：明细交易类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 产品名称：用户所采购的各类云产品
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// 子产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 子产品名称：用户采购的具体产品细分类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域名称：资源所属地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 可用区：资源所属可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 实例类型编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例类型：购买的产品服务对应的实例类型，包括资源包、RI、SP、竞价实例。常规实例默认展示“-”
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTypeName *string `json:"InstanceTypeName,omitnil,omitempty" name:"InstanceTypeName"`

	// 资源ID：不同产品因资源形态不同，资源内容不完全相同，如云服务器 CVM 为对应的实例 ID； 若该产品被分拆，则展示产品分拆后的分拆项 ID，如 COS 桶 ID，CDN 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 实例名称：用户在控制台为资源设置的名称，如未设置默认为空；若该产品被分拆，则展示分拆产品分拆后的分拆项资源别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 分账标签：资源绑定的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*BillTag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称：资源归属的项目，用户在控制台给资源自主分配项目，未分配则是默认项目
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 费用归集类型：费用来源类型，分摊、归集、未分配
	// 0 - 分摊
	// 1 - 归集
	// -1 - 未分配
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllocationType *int64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`

	// 组件原价：原价 = 组件刊例价 * 组件用量 * 使用时长（如客户享受一口价/合同价则默认不展示，退费类场景也默认不展示），指定价模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 预留实例抵扣时长：本产品或服务使用预留实例抵扣的使用时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiTimeSpan *string `json:"RiTimeSpan,omitnil,omitempty" name:"RiTimeSpan"`

	// 预留实例抵扣原价：本产品或服务使用预留实例抵扣的组件原价金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiCost *string `json:"RiCost,omitnil,omitempty" name:"RiCost"`

	// 优惠后总价：优惠后总价 =（原价 - 预留实例抵扣原价 - 节省计划抵扣原价）* 折扣率
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 现金账户支出(元)：通过现金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 代金券支出(元)：使用各类优惠券（如代金券、现金券等）支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 赠送账户支出(元)：使用赠送金支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 分成账户支出(元)：通过分成金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// 组件名称编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemCode *string `json:"ItemCode,omitnil,omitempty" name:"ItemCode"`

	// 组件名称：用户购买的产品或服务，所包含的具体组件
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemCodeName *string `json:"ItemCodeName,omitnil,omitempty" name:"ItemCodeName"`

	// 组件类型编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentCode *string `json:"ComponentCode,omitnil,omitempty" name:"ComponentCode"`

	// 组件类型：用户购买的产品或服务对应的组件大类
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentCodeName *string `json:"ComponentCodeName,omitnil,omitempty" name:"ComponentCodeName"`

	// 分拆项 ID：涉及分拆产品的分拆后的分拆项 ID，如 COS 桶 ID，CDN 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: SplitItemId is deprecated.
	SplitItemId *string `json:"SplitItemId,omitnil,omitempty" name:"SplitItemId"`

	// 分拆项名称：涉及分拆产品的分拆后的分拆项
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: SplitItemName is deprecated.
	SplitItemName *string `json:"SplitItemName,omitnil,omitempty" name:"SplitItemName"`

	// 开始使用时间：产品服务开始使用时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeeBeginTime *string `json:"FeeBeginTime,omitnil,omitempty" name:"FeeBeginTime"`

	// 结束使用时间：产品服务结束使用时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeeEndTime *string `json:"FeeEndTime,omitnil,omitempty" name:"FeeEndTime"`

	// 节省计划抵扣原价：节省计划抵扣原价 = 节省计划包抵扣面值 / 节省计划抵扣率
	// 注意：此字段可能返回 null，表示取不到有效值。
	SPCost *string `json:"SPCost,omitnil,omitempty" name:"SPCost"`

	// 国内国际编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionType *string `json:"RegionType,omitnil,omitempty" name:"RegionType"`

	// 国内国际：资源所属区域类型（国内、国际）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionTypeName *string `json:"RegionTypeName,omitnil,omitempty" name:"RegionTypeName"`

	// 组件刊例价：组件的官网原始单价（如客户享受一口价/合同价则默认不展示）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SinglePrice *string `json:"SinglePrice,omitnil,omitempty" name:"SinglePrice"`

	// 组件单价：组件的折后单价，组件单价 = 刊例价 * 折扣
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractPrice *string `json:"ContractPrice,omitnil,omitempty" name:"ContractPrice"`

	// 组件价格单位：组件价格的单位，单位构成：元/用量单位/时长单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	SinglePriceUnit *string `json:"SinglePriceUnit,omitnil,omitempty" name:"SinglePriceUnit"`

	// 组件用量：该组件实际结算用量，组件用量=组件原始用量-抵扣用量（含资源包）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedAmount *string `json:"UsedAmount,omitnil,omitempty" name:"UsedAmount"`

	// 组件用量单位：组件用量对应的单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedAmountUnit *string `json:"UsedAmountUnit,omitnil,omitempty" name:"UsedAmountUnit"`

	// 使用时长：资源使用的时长，组件用量=组件原始使用时长-抵扣时长（含资源包）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeSpan *string `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 时长单位：资源使用时长的单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 备注属性（实例配置）：其他备注信息，如预留实例的预留实例类型和交易类型、CCN 产品的两端地域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReserveDetail *string `json:"ReserveDetail,omitnil,omitempty" name:"ReserveDetail"`

	// 原始用量/时长：组件被资源包抵扣前的原始用量
	// （目前仅实时音视频、弹性微服务、云呼叫中心及专属可用区产品支持该信息外显，其他产品尚在接入中）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalMeasure *string `json:"RealTotalMeasure,omitnil,omitempty" name:"RealTotalMeasure"`

	// 抵扣用量/时长（含资源包）：组件被资源包抵扣的用量
	// （目前仅实时音视频、弹性微服务、云呼叫中心及专属可用区产品支持该信息外显，其他产品尚在接入中）
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeductedMeasure *string `json:"DeductedMeasure,omitnil,omitempty" name:"DeductedMeasure"`

	// 折扣率：本资源享受的折扣率（如客户享受一口价/合同价则默认不展示，退费场景也默认不展示）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Discount *string `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 混合折扣率：综合各类折扣抵扣信息后的最终折扣率，混合折扣率=优惠后总价/原价
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlendedDiscount *string `json:"BlendedDiscount,omitnil,omitempty" name:"BlendedDiscount"`

	// 价格属性：该组件除单价、时长外的其他影响折扣定价的属性信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PriceInfo []*string `json:"PriceInfo,omitnil,omitempty" name:"PriceInfo"`

	// 计算规则说明：特殊交易类型计费结算的详细计算说明，如退费及变配
	// 注意：此字段可能返回 null，表示取不到有效值。
	Formula *string `json:"Formula,omitnil,omitempty" name:"Formula"`

	// 计费规则：各产品详细的计费规则官网说明链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	FormulaUrl *string `json:"FormulaUrl,omitnil,omitempty" name:"FormulaUrl"`

	// 配置描述：资源配置规格信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentConfig *string `json:"ComponentConfig,omitnil,omitempty" name:"ComponentConfig"`

	// SPDeduction
	// 注意：此字段可能返回 null，表示取不到有效值。
	SPDeduction *string `json:"SPDeduction,omitnil,omitempty" name:"SPDeduction"`

	// 节省计划抵扣率：节省计划可用余额额度范围内，节省计划对于此组件打的折扣率
	// 注意：此字段可能返回 null，表示取不到有效值。
	SPDeductionRate *string `json:"SPDeductionRate,omitnil,omitempty" name:"SPDeductionRate"`

	// AssociatedOrder
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociatedOrder *string `json:"AssociatedOrder,omitnil,omitempty" name:"AssociatedOrder"`

	// 当前消费项的优惠对象，例如：官网折扣、用户折扣、活动折扣。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountObject *string `json:"DiscountObject,omitnil,omitempty" name:"DiscountObject"`

	// 当前消费项的优惠类型，例如：折扣、合同价。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountType *string `json:"DiscountType,omitnil,omitempty" name:"DiscountType"`

	// 对优惠类型的补充描述，例如：商务折扣8折，则优惠类型为“折扣”，优惠内容为“0.8”。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountContent *string `json:"DiscountContent,omitnil,omitempty" name:"DiscountContent"`

	// 账单月
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`
}

type AllocationSummaryByResource struct {
	// 分账单元唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// 分账单元名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreeNodeUniqKeyName *string `json:"TreeNodeUniqKeyName,omitnil,omitempty" name:"TreeNodeUniqKeyName"`

	// 日期：结算日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillDate *string `json:"BillDate,omitnil,omitempty" name:"BillDate"`

	// 支付者 UIN：支付者的账号 ID，账号 ID 是用户在腾讯云的唯一账号标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 使用者 UIN：实际使用资源的账号 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 操作者 UIN：操作者账号 ID（预付费资源下单或后付费操作开通资源账号的ID或者角色 ID）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// 计费模式编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 计费模式：资源的计费模式，区分为包年包月和按量计费
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// 交易类型编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 交易类型：明细交易类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 产品名称：用户所采购的各类云产品
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// 子产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 子产品名称：用户采购的具体产品细分类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域名称：资源所属地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 可用区：资源所属可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 实例类型编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例类型：购买的产品服务对应的实例类型，包括资源包、RI、SP、竞价实例。常规实例默认展示“-”
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTypeName *string `json:"InstanceTypeName,omitnil,omitempty" name:"InstanceTypeName"`

	// 资源ID：不同产品因资源形态不同，资源内容不完全相同，如云服务器 CVM 为对应的实例 ID； 若该产品被分拆，则展示产品分拆后的分拆项 ID，如 COS 桶 ID，CDN 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 实例名称：用户在控制台为资源设置的名称，如未设置默认为空；若该产品被分拆，则展示分拆产品分拆后的分拆项资源别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 分账标签：资源绑定的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*BillTag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称：资源归属的项目，用户在控制台给资源自主分配项目，未分配则是默认项目
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 费用归集类型：费用来源类型，分摊、归集、未分配
	// 0 - 分摊 
	// 1 - 归集 
	// -1 -  未分配
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllocationType *int64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`

	// 组件原价：原价 = 组件刊例价 * 组件用量 * 使用时长（如客户享受一口价/合同价则默认不展示，退费类场景也默认不展示），指定价模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 预留实例抵扣时长：本产品或服务使用预留实例抵扣的使用时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiTimeSpan *string `json:"RiTimeSpan,omitnil,omitempty" name:"RiTimeSpan"`

	// 预留实例抵扣原价：本产品或服务使用预留实例抵扣的组件原价金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiCost *string `json:"RiCost,omitnil,omitempty" name:"RiCost"`

	// 优惠后总价：优惠后总价 =（原价 - 预留实例抵扣原价 - 节省计划抵扣原价）* 折扣率
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 现金账户支出(元)：通过现金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 代金券支出(元)：使用各类优惠券（如代金券、现金券等）支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 赠送账户支出(元)：使用赠送金支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 分成账户支出(元)：通过分成金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// 分拆项 ID：涉及分拆产品的分拆后的分拆项 ID，如 COS 桶 ID，CDN 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: SplitItemId is deprecated.
	SplitItemId *string `json:"SplitItemId,omitnil,omitempty" name:"SplitItemId"`

	// 分拆项名称：涉及分拆产品的分拆后的分拆项
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: SplitItemName is deprecated.
	SplitItemName *string `json:"SplitItemName,omitnil,omitempty" name:"SplitItemName"`

	// 开始使用时间：产品服务开始使用时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeeBeginTime *string `json:"FeeBeginTime,omitnil,omitempty" name:"FeeBeginTime"`

	// 结束使用时间：产品服务结束使用时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeeEndTime *string `json:"FeeEndTime,omitnil,omitempty" name:"FeeEndTime"`

	// 节省计划抵扣原价：节省计划抵扣原价 = 节省计划包抵扣面值 / 节省计划抵扣率
	// 注意：此字段可能返回 null，表示取不到有效值。
	SPCost *string `json:"SPCost,omitnil,omitempty" name:"SPCost"`

	// 国内国际编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionType *string `json:"RegionType,omitnil,omitempty" name:"RegionType"`

	// 国内国际：资源所属区域类型（国内、国际）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionTypeName *string `json:"RegionTypeName,omitnil,omitempty" name:"RegionTypeName"`

	// 配置描述：对应资源下各组件名称及用量（如组件为用量累加型计费则为合计用量）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentConfig *string `json:"ComponentConfig,omitnil,omitempty" name:"ComponentConfig"`

	// SPDeduction
	// 注意：此字段可能返回 null，表示取不到有效值。
	SPDeduction *string `json:"SPDeduction,omitnil,omitempty" name:"SPDeduction"`

	// 账单月
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`
}

type AllocationTreeNode struct {
	// 分账单元唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// 分账单元名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreeNodeUniqKeyName *string `json:"TreeNodeUniqKeyName,omitnil,omitempty" name:"TreeNodeUniqKeyName"`
}

type AnalyseActionTypeDetail struct {
	// 交易类型code
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 交易类型Name
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`
}

type AnalyseAmountDetail struct {
	// 费用类型
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 是否展示
	Display *int64 `json:"Display,omitnil,omitempty" name:"Display"`
}

type AnalyseBusinessDetail struct {
	// 产品码code
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 产品名称
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`
}

type AnalyseConditionDetail struct {
	// 产品
	// 注意：此字段可能返回 null，表示取不到有效值。
	Business []*AnalyseBusinessDetail `json:"Business,omitnil,omitempty" name:"Business"`

	// 项目
	// 注意：此字段可能返回 null，表示取不到有效值。
	Project []*AnalyseProjectDetail `json:"Project,omitnil,omitempty" name:"Project"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region []*AnalyseRegionDetail `json:"Region,omitnil,omitempty" name:"Region"`

	// 计费模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode []*AnalysePayModeDetail `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 交易类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionType []*AnalyseActionTypeDetail `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone []*AnalyseZoneDetail `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 资源所有者Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin []*AnalyseOwnerUinDetail `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 费用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Amount []*AnalyseAmountDetail `json:"Amount,omitnil,omitempty" name:"Amount"`
}

type AnalyseConditions struct {
	// 产品名称代码
	BusinessCodes *string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// 子产品名称代码
	ProductCodes *string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// 组件类型代码
	ComponentCode *string `json:"ComponentCode,omitnil,omitempty" name:"ComponentCode"`

	// 可用区ID：资源所属可用区ID
	ZoneIds *string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 地域ID:资源所属地域ID
	RegionIds *string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// 项目ID:资源所属项目ID
	ProjectIds *string `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 计费模式 prePay(表示包年包月)/postPay(表示按量计费)
	PayModes *string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// 交易类型，查询交易类型（请使用交易类型code入参）
	ActionTypes *string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// 分账标签键
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 费用类型，查询费用类型（请使用费用类型code入参)入参枚举如下：
	// cashPayAmount:现金 
	// incentivePayAmount:赠送金 
	// voucherPayAmount:优惠券 
	// tax:税金 
	// costBeforeTax:税前价
	FeeType *string `json:"FeeType,omitnil,omitempty" name:"FeeType"`

	// 查询成本分析数据的用户UIN
	PayerUins *string `json:"PayerUins,omitnil,omitempty" name:"PayerUins"`

	// 使用资源的用户UIN
	OwnerUins *string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// 消耗类型，查询消耗类型（请使用消耗类型code入参）
	ConsumptionTypes *string `json:"ConsumptionTypes,omitnil,omitempty" name:"ConsumptionTypes"`
}

type AnalyseDetail struct {
	// 时间
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 金额
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 日期明细金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeDetail []*AnalyseTimeDetail `json:"TimeDetail,omitnil,omitempty" name:"TimeDetail"`
}

type AnalyseHeaderDetail struct {
	// 表头日期
	HeadDetail []*AnalyseHeaderTimeDetail `json:"HeadDetail,omitnil,omitempty" name:"HeadDetail"`

	// 时间
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 总计
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`
}

type AnalyseHeaderTimeDetail struct {
	// 日期
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type AnalyseOwnerUinDetail struct {
	// 使用者uin
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

type AnalysePayModeDetail struct {
	// 计费模式code
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 计费模式Name
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`
}

type AnalyseProjectDetail struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 默认项目
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`
}

type AnalyseRegionDetail struct {
	// 地域id
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`
}

type AnalyseTimeDetail struct {
	// 日期
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 金额
	Money *string `json:"Money,omitnil,omitempty" name:"Money"`
}

type AnalyseZoneDetail struct {
	// 可用区id
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 可用区Name
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`
}

type ApplicableProducts struct {
	// 适用商品名称，值为“全产品通用”或商品名称组成的string，以","分割。
	GoodsName *string `json:"GoodsName,omitnil,omitempty" name:"GoodsName"`

	// postPay后付费/prePay预付费/riPay预留实例/空字符串或者"*"表示全部模式。如GoodsName为多个商品名以","分割组成的string，而PayMode为"*"，表示每一件商品的模式都为"*"。
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

type BillActionType struct {
	// 交易类型编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 交易类型：明细交易类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`
}

type BillBusiness struct {
	// 产品编码
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 产品名称：用户所采购的各类云产品
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`
}

type BillBusinessLink struct {
	// 产品编码
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 产品名称
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// 子产品
	Children []*BillProductLink `json:"Children,omitnil,omitempty" name:"Children"`
}

type BillComponent struct {
	// 组件类型编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentCode *string `json:"ComponentCode,omitnil,omitempty" name:"ComponentCode"`

	// 组件类型：用户购买的产品或服务对应的组件大类
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentCodeName *string `json:"ComponentCodeName,omitnil,omitempty" name:"ComponentCodeName"`
}

type BillDays struct {
	// 日期：结算日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillDay *string `json:"BillDay,omitnil,omitempty" name:"BillDay"`
}

type BillDetail struct {
	// 产品名称：用户所采购的各类云产品，例如：云服务器 CVM
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// 子产品名称：用户采购的具体产品细分类型，例如：云服务器 CVM-标准型 S1
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// 计费模式：资源的计费模式，区分为包年包月和按量计费
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// 项目名称：资源归属的项目，用户在控制台给资源自主分配项目，未分配则是默认项目
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 地域：资源所属地域，如华南地区（广州）
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 可用区：资源所属可用区，如广州三区
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 资源 ID：账单中出账对象 ID，不同产品因资源形态不同，资源内容不完全相同，如云服务器 CVM 为对应的实例 ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源别名：用户在控制台为资源设置的名称，如果未设置，则默认为空
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 交易类型，如包年包月新购、包年包月续费、按量计费扣费等类型
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// 订单ID：包年包月计费模式下订购的订单号
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 交易ID：结算扣费单号
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// 扣费时间：结算扣费时间
	PayTime *string `json:"PayTime,omitnil,omitempty" name:"PayTime"`

	// 开始使用时间：产品服务开始使用时间
	FeeBeginTime *string `json:"FeeBeginTime,omitnil,omitempty" name:"FeeBeginTime"`

	// 结束使用时间：产品服务结束使用时间
	FeeEndTime *string `json:"FeeEndTime,omitnil,omitempty" name:"FeeEndTime"`

	// 组件列表
	ComponentSet []*BillDetailComponent `json:"ComponentSet,omitnil,omitempty" name:"ComponentSet"`

	// 支付者UIN：支付者的账号 ID，账号 ID 是用户在腾讯云的唯一账号标识
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 使用者UIN：实际使用资源的账号 ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 操作者UIN：操作者账号 ID（预付费资源下单或后付费操作开通资源账号的 ID 或者角色 ID ）
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*BillTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 子产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 交易类型编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 价格属性：该组件除单价、时长外的其他影响折扣定价的属性信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PriceInfo []*string `json:"PriceInfo,omitnil,omitempty" name:"PriceInfo"`

	// 关联交易单据ID：和本笔交易关联单据 ID，如，冲销订单，记录原订单、重结订单，退费单记录对应的原购买订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociatedOrder *BillDetailAssociatedOrder `json:"AssociatedOrder,omitnil,omitempty" name:"AssociatedOrder"`

	// 计算说明：特殊交易类型计费结算的详细计算说明，如退费及变配
	// 注意：此字段可能返回 null，表示取不到有效值。
	Formula *string `json:"Formula,omitnil,omitempty" name:"Formula"`

	// 计费规则：各产品详细的计费规则官网说明链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	FormulaUrl *string `json:"FormulaUrl,omitnil,omitempty" name:"FormulaUrl"`

	// 账单归属日
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillDay *string `json:"BillDay,omitnil,omitempty" name:"BillDay"`

	// 账单归属月
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// 账单记录ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 国内国际编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionType *string `json:"RegionType,omitnil,omitempty" name:"RegionType"`

	// 国内国际：资源所属区域类型（国内、国际）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionTypeName *string `json:"RegionTypeName,omitnil,omitempty" name:"RegionTypeName"`

	// 备注属性（实例配置）：其他备注信息，如预留实例的预留实例类型和交易类型、CCN 产品的两端地域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReserveDetail *string `json:"ReserveDetail,omitnil,omitempty" name:"ReserveDetail"`

	// 优惠对象
	DiscountObject *string `json:"DiscountObject,omitnil,omitempty" name:"DiscountObject"`

	// 优惠类型
	DiscountType *string `json:"DiscountType,omitnil,omitempty" name:"DiscountType"`

	// 优惠内容
	DiscountContent *string `json:"DiscountContent,omitnil,omitempty" name:"DiscountContent"`
}

type BillDetailAssociatedOrder struct {
	// 新购订单
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrepayPurchase *string `json:"PrepayPurchase,omitnil,omitempty" name:"PrepayPurchase"`

	// 续费订单
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrepayRenew *string `json:"PrepayRenew,omitnil,omitempty" name:"PrepayRenew"`

	// 升配订单
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrepayModifyUp *string `json:"PrepayModifyUp,omitnil,omitempty" name:"PrepayModifyUp"`

	// 冲销订单
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReverseOrder *string `json:"ReverseOrder,omitnil,omitempty" name:"ReverseOrder"`

	// 优惠调整后订单
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewOrder *string `json:"NewOrder,omitnil,omitempty" name:"NewOrder"`

	// 优惠调整前订单
	// 注意：此字段可能返回 null，表示取不到有效值。
	Original *string `json:"Original,omitnil,omitempty" name:"Original"`
}

type BillDetailComponent struct {
	// 组件类型：用户购买的产品或服务对应的组件大类，例如：云服务器 CVM 的组件：CPU、内存等
	ComponentCodeName *string `json:"ComponentCodeName,omitnil,omitempty" name:"ComponentCodeName"`

	// 组件名称：用户购买的产品或服务，所包含的具体组件
	ItemCodeName *string `json:"ItemCodeName,omitnil,omitempty" name:"ItemCodeName"`

	// 组件刊例价：组件的官网原始单价（如果客户享受一口价/合同价则默认不展示）
	SinglePrice *string `json:"SinglePrice,omitnil,omitempty" name:"SinglePrice"`

	// 组件指定价（已废弃）
	//
	// Deprecated: SpecifiedPrice is deprecated.
	SpecifiedPrice *string `json:"SpecifiedPrice,omitnil,omitempty" name:"SpecifiedPrice"`

	// 组件价格单位：组件价格的单位，单位构成：元/用量单位/时长单位
	PriceUnit *string `json:"PriceUnit,omitnil,omitempty" name:"PriceUnit"`

	// 组件用量：该组件实际结算用量，组件用量 = 组件原始用量 - 抵扣用量（含资源包
	UsedAmount *string `json:"UsedAmount,omitnil,omitempty" name:"UsedAmount"`

	// 组件用量单位：组件用量对应的单位
	UsedAmountUnit *string `json:"UsedAmountUnit,omitnil,omitempty" name:"UsedAmountUnit"`

	// 原始用量/时长：组件被资源包抵扣前的原始用量/时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalMeasure *string `json:"RealTotalMeasure,omitnil,omitempty" name:"RealTotalMeasure"`

	// 抵扣用量/时长（含资源包）：组件被资源包抵扣的用量/时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeductedMeasure *string `json:"DeductedMeasure,omitnil,omitempty" name:"DeductedMeasure"`

	// 使用时长：资源使用的时长
	TimeSpan *string `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 时长单位：资源使用时长的单位
	TimeUnitName *string `json:"TimeUnitName,omitnil,omitempty" name:"TimeUnitName"`

	// 组件原价：原价 = 组件刊例价 * 组件用量 * 使用时长（如果客户享受一口价/合同价则默认不展示，退费类场景也默认不展示）
	Cost *string `json:"Cost,omitnil,omitempty" name:"Cost"`

	// 折扣率：本资源享受的折扣率（如果客户享受一口价/合同价则默认不展示，退费场景也默认不展示）
	Discount *string `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 优惠类型
	ReduceType *string `json:"ReduceType,omitnil,omitempty" name:"ReduceType"`

	// 优惠后总价：优惠后总价=（原价 - 预留实例抵扣原价 - 节省计划抵扣原价）* 折扣率
	RealCost *string `json:"RealCost,omitnil,omitempty" name:"RealCost"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// 组件类型编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemCode *string `json:"ItemCode,omitnil,omitempty" name:"ItemCode"`

	// 组件名称编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentCode *string `json:"ComponentCode,omitnil,omitempty" name:"ComponentCode"`

	// 组件单价：组件的折后单价，组件单价 = 刊例价 * 折扣
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractPrice *string `json:"ContractPrice,omitnil,omitempty" name:"ContractPrice"`

	// 实例类型：购买的产品服务对应的实例类型，包括资源包、RI、SP、竞价实例。正常的实例展示默认为不展示
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 预留实例抵扣的使用时长：本产品或服务使用预留实例抵扣的使用时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiTimeSpan *string `json:"RiTimeSpan,omitnil,omitempty" name:"RiTimeSpan"`

	// 预留实例抵扣组件原价：本产品或服务使用预留实例抵扣的组件原价金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalCostWithRI *string `json:"OriginalCostWithRI,omitnil,omitempty" name:"OriginalCostWithRI"`

	// 节省计划抵扣率：节省计划可用余额额度范围内，节省计划对于此组件打的折扣率
	// 注意：此字段可能返回 null，表示取不到有效值。
	SPDeductionRate *string `json:"SPDeductionRate,omitnil,omitempty" name:"SPDeductionRate"`

	// 节省计划抵扣金额（已废弃）
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: SPDeduction is deprecated.
	SPDeduction *string `json:"SPDeduction,omitnil,omitempty" name:"SPDeduction"`

	// 节省计划抵扣组件原价：节省计划抵扣原价=节省计划包抵扣金额/节省计划抵扣率
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalCostWithSP *string `json:"OriginalCostWithSP,omitnil,omitempty" name:"OriginalCostWithSP"`

	// 混合折扣率：综合各类折扣抵扣信息后的最终折扣率，混合折扣率 = 优惠后总价 / 组件原价
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlendedDiscount *string `json:"BlendedDiscount,omitnil,omitempty" name:"BlendedDiscount"`

	// 配置描述：资源配置规格信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentConfig []*BillDetailComponentConfig `json:"ComponentConfig,omitnil,omitempty" name:"ComponentConfig"`
}

type BillDetailComponentConfig struct {
	// 配置描述名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 配置描述值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type BillDistributionResourceSummary struct {
	// 产品名称：用户所采购的各类云产品，例如：云服务器 CVM
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// 子产品名称：用户采购的具体产品细分类型，例如：云服务器 CVM-标准型 S1
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// 计费模式：资源的计费模式，区分为包年包月和按量计费
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// 项目名称：资源归属的项目，用户在控制台给资源自主分配项目，未分配则是默认项目
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 地域：资源所属地域，如华南地区（广州）
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 可用区：资源所属可用区，如广州三区
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 资源 ID：账单中出账对象 ID，不同产品因资源形态不同，资源内容不完全相同，如云服务器 CVM 为对应的实例 ID	
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源别名：用户在控制台为资源设置的名称，如果未设置，则默认为空
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 交易类型：如包年包月新购、包年包月续费、按量计费扣费等类型
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// 订单ID：包年包月计费模式下订购的订单号
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 扣费时间：结算扣费时间
	PayTime *string `json:"PayTime,omitnil,omitempty" name:"PayTime"`

	// 开始使用时间：产品服务开始使用时间
	FeeBeginTime *string `json:"FeeBeginTime,omitnil,omitempty" name:"FeeBeginTime"`

	// 结束使用时间：产品服务结束使用时间
	FeeEndTime *string `json:"FeeEndTime,omitnil,omitempty" name:"FeeEndTime"`

	// 配置描述：该资源下的计费项名称和用量合并展示，仅在资源账单体现
	ConfigDesc *string `json:"ConfigDesc,omitnil,omitempty" name:"ConfigDesc"`

	// 扩展字段1：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField1 *string `json:"ExtendField1,omitnil,omitempty" name:"ExtendField1"`

	// 扩展字段2：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField2 *string `json:"ExtendField2,omitnil,omitempty" name:"ExtendField2"`

	// 原价：原价 = 组件刊例价 * 组件用量 * 使用时长（如果客户享受一口价/合同价则默认不展示，退费类场景也默认不展示）
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 折扣率：本资源享受的折扣率（如果客户享受一口价/合同价则默认不展示，退费场景也默认不展示）
	Discount *string `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 优惠类型
	ReduceType *string `json:"ReduceType,omitnil,omitempty" name:"ReduceType"`

	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// 扩展字段3：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField3 *string `json:"ExtendField3,omitnil,omitempty" name:"ExtendField3"`

	// 扩展字段4：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField4 *string `json:"ExtendField4,omitnil,omitempty" name:"ExtendField4"`

	// 扩展字段5：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField5 *string `json:"ExtendField5,omitnil,omitempty" name:"ExtendField5"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*BillTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 使用者UIN：实际使用资源的账号 ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 操作者UIN：操作者账号 ID（预付费资源下单或后付费操作开通资源账号的 ID 或者角色 ID ）
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// 产品编码
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 子产品编码
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 地域ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 实例类型：购买的产品服务对应的实例类型，包括资源包、RI、SP、竞价实例。正常的实例展示默认为不展示
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 预留实例抵扣组件原价：本产品或服务使用预留实例抵扣的组件原价金额	
	OriginalCostWithRI *string `json:"OriginalCostWithRI,omitnil,omitempty" name:"OriginalCostWithRI"`

	// 节省计划抵扣金额（已废弃）
	//
	// Deprecated: SPDeduction is deprecated.
	SPDeduction *string `json:"SPDeduction,omitnil,omitempty" name:"SPDeduction"`

	// 节省计划抵扣组件原价：节省计划抵扣原价=节省计划包抵扣金额/节省计划抵扣率	
	OriginalCostWithSP *string `json:"OriginalCostWithSP,omitnil,omitempty" name:"OriginalCostWithSP"`

	// 账单归属月
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`
}

type BillInstanceType struct {
	// 实例类型编码
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例类型：购买的产品服务对应的实例类型，包括资源包、RI、SP、竞价实例。常规实例默认展示“-”
	InstanceTypeName *string `json:"InstanceTypeName,omitnil,omitempty" name:"InstanceTypeName"`
}

type BillItem struct {
	// 组件名称编码
	ItemCode *string `json:"ItemCode,omitnil,omitempty" name:"ItemCode"`

	// 组件名称：用户购买的产品或服务，所包含的具体组件
	ItemCodeName *string `json:"ItemCodeName,omitnil,omitempty" name:"ItemCodeName"`
}

type BillOperateUin struct {
	// 操作者 UIN：操作者账号 ID（预付费资源下单或后付费操作开通资源账号的ID或者角色 ID）
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`
}

type BillOwnerUin struct {
	// 使用者 UIN：实际使用资源的账号 ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

type BillPayMode struct {
	// 计费模式编码
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 计费模式：资源的计费模式，区分为包年包月和按量计费
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`
}

type BillProduct struct {
	// 子产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 子产品名称：用户采购的具体产品细分类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`
}

type BillProductLink struct {
	// 子产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 子产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// 组件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Children []*BillItem `json:"Children,omitnil,omitempty" name:"Children"`
}

type BillProject struct {
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称：资源归属的项目，用户在控制台给资源自主分配项目，未分配则是默认项目
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`
}

type BillRegion struct {
	// 地域ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域名称：资源所属地域
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`
}

type BillResourceSummary struct {
	// 产品名称：用户所采购的各类云产品，例如：云服务器 CVM
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// 子产品名称：用户采购的具体产品细分类型，例如：云服务器 CVM-标准型 S1
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// 计费模式：资源的计费模式，区分为包年包月和按量计费
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// 项目名称：资源归属的项目，用户在控制台给资源自主分配项目，未分配则是默认项目
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 地域：资源所属地域，如华南地区（广州）
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 可用区：资源所属可用区，如广州三区
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 资源 ID：账单中出账对象 ID，不同产品因资源形态不同，资源内容不完全相同，如云服务器 CVM 为对应的实例 ID	
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源别名：用户在控制台为资源设置的名称，如果未设置，则默认为空
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 交易类型：如包年包月新购、包年包月续费、按量计费扣费等类型
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// 订单ID：包年包月计费模式下订购的订单号
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 扣费时间：结算扣费时间
	PayTime *string `json:"PayTime,omitnil,omitempty" name:"PayTime"`

	// 开始使用时间：产品服务开始使用时间
	FeeBeginTime *string `json:"FeeBeginTime,omitnil,omitempty" name:"FeeBeginTime"`

	// 结束使用时间：产品服务结束使用时间
	FeeEndTime *string `json:"FeeEndTime,omitnil,omitempty" name:"FeeEndTime"`

	// 配置描述：该资源下的计费项名称和用量合并展示，仅在资源账单体现
	ConfigDesc *string `json:"ConfigDesc,omitnil,omitempty" name:"ConfigDesc"`

	// 扩展字段1：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField1 *string `json:"ExtendField1,omitnil,omitempty" name:"ExtendField1"`

	// 扩展字段2：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField2 *string `json:"ExtendField2,omitnil,omitempty" name:"ExtendField2"`

	// 原价：原价 = 组件刊例价 * 组件用量 * 使用时长（如果客户享受一口价/合同价则默认不展示，退费类场景也默认不展示）
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 折扣率：本资源享受的折扣率（如果客户享受一口价/合同价则默认不展示，退费场景也默认不展示）
	Discount *string `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 优惠类型
	ReduceType *string `json:"ReduceType,omitnil,omitempty" name:"ReduceType"`

	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// 扩展字段3：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField3 *string `json:"ExtendField3,omitnil,omitempty" name:"ExtendField3"`

	// 扩展字段4：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField4 *string `json:"ExtendField4,omitnil,omitempty" name:"ExtendField4"`

	// 扩展字段5：产品对应的扩展属性信息，仅在资源账单体现
	ExtendField5 *string `json:"ExtendField5,omitnil,omitempty" name:"ExtendField5"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*BillTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 支付者UIN：支付者的账号 ID，账号 ID 是用户在腾讯云的唯一账号标识
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 使用者UIN：实际使用资源的账号 ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 操作者UIN：操作者账号 ID（预付费资源下单或后付费操作开通资源账号的 ID 或者角色 ID ）
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// 产品编码
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 子产品编码
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 地域ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 实例类型：购买的产品服务对应的实例类型，包括资源包、RI、SP、竞价实例。正常的实例展示默认为不展示
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 预留实例抵扣组件原价：本产品或服务使用预留实例抵扣的组件原价金额	
	OriginalCostWithRI *string `json:"OriginalCostWithRI,omitnil,omitempty" name:"OriginalCostWithRI"`

	// 节省计划抵扣金额（已废弃）
	//
	// Deprecated: SPDeduction is deprecated.
	SPDeduction *string `json:"SPDeduction,omitnil,omitempty" name:"SPDeduction"`

	// 节省计划抵扣组件原价：节省计划抵扣原价=节省计划包抵扣金额/节省计划抵扣率	
	OriginalCostWithSP *string `json:"OriginalCostWithSP,omitnil,omitempty" name:"OriginalCostWithSP"`

	// 账单归属月
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`
}

type BillTag struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type BillTagInfo struct {
	// 分账标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type BillTransactionInfo struct {
	// 收支类型：deduct 扣费, recharge 充值, return 退费， block 冻结, unblock 解冻
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 流水金额，单位（分）；正数表示入账，负数表示出账
	Amount *int64 `json:"Amount,omitnil,omitempty" name:"Amount"`

	// 可用余额，单位（分）；正数表示入账，负数表示出账
	Balance *int64 `json:"Balance,omitnil,omitempty" name:"Balance"`

	// 流水号，如20190131020000236005203583326401
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// 描述信息
	OperationInfo *string `json:"OperationInfo,omitnil,omitempty" name:"OperationInfo"`

	// 操作时间"2019-01-31 23:35:10.000"
	OperationTime *string `json:"OperationTime,omitnil,omitempty" name:"OperationTime"`

	// 现金账户余额，单位（分）
	Cash *int64 `json:"Cash,omitnil,omitempty" name:"Cash"`

	// 赠送金余额，单位（分）
	Incentive *int64 `json:"Incentive,omitnil,omitempty" name:"Incentive"`

	// 冻结余额，单位（分）
	Freezing *int64 `json:"Freezing,omitnil,omitempty" name:"Freezing"`

	// 交易渠道
	PayChannel *string `json:"PayChannel,omitnil,omitempty" name:"PayChannel"`

	// 扣费模式：trade 包年包月(预付费)，hourh  按量-小时结，hourd 按量-日结，hourm 按量-月结，month 按量-月结
	DeductMode *string `json:"DeductMode,omitnil,omitempty" name:"DeductMode"`
}

type BillZoneId struct {
	// 可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 可用区：资源所属可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`
}

type BusinessSummaryInfo struct {
	// 产品编码
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 产品名称：用户所采购的各类云产品，例如：云服务器 CVM
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`
}

type BusinessSummaryOverviewItem struct {
	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 产品名称：用户所采购的各类云产品，例如：云服务器 CVM
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil,omitempty" name:"RealTotalCostRatio"`

	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// 账单月份，格式2019-08
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type BusinessSummaryTotal struct {
	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type ConditionBusiness struct {
	// 产品名称代码
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 产品名称
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`
}

type ConditionPayMode struct {
	// 付费模式
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 付费模式名称
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`
}

type ConditionProject struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`
}

type ConditionRegion struct {
	// 地域ID
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`
}

type Conditions struct {
	// 只支持6和12两个值
	TimeRange *uint64 `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// 产品名称代码
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 地域ID
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 付费模式，可选prePay和postPay
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 资源关键字
	ResourceKeyword *string `json:"ResourceKeyword,omitnil,omitempty" name:"ResourceKeyword"`

	// 产品名称代码
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// 子产品名称代码
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// 地域ID
	RegionIds []*int64 `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// 项目ID
	ProjectIds []*int64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 付费模式，可选prePay和postPay
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// 交易类型
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// 是否隐藏0元流水
	HideFreeCost *int64 `json:"HideFreeCost,omitnil,omitempty" name:"HideFreeCost"`

	// 排序规则，可选desc和asc
	OrderByCost *string `json:"OrderByCost,omitnil,omitempty" name:"OrderByCost"`

	// 交易ID
	BillIds []*string `json:"BillIds,omitnil,omitempty" name:"BillIds"`

	// 组件编码
	ComponentCodes []*string `json:"ComponentCodes,omitnil,omitempty" name:"ComponentCodes"`

	// 文件ID
	FileIds []*string `json:"FileIds,omitnil,omitempty" name:"FileIds"`

	// 文件类型
	FileTypes []*string `json:"FileTypes,omitnil,omitempty" name:"FileTypes"`

	// 状态
	Status []*uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ConsumptionBusinessSummaryDataItem struct {
	// 产品名称代码
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 产品名称
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// 折后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 费用趋势
	Trend *ConsumptionSummaryTrend `json:"Trend,omitnil,omitempty" name:"Trend"`

	// 现金
	// 注意：此字段可能返回 null，表示取不到有效值。
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 赠送金
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 代金券
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 分成金
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// 地域名称（仅在地域汇总总展示）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`
}

type ConsumptionProjectSummaryDataItem struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 折后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 趋势
	Trend *ConsumptionSummaryTrend `json:"Trend,omitnil,omitempty" name:"Trend"`

	// 产品消耗详情
	Business []*ConsumptionBusinessSummaryDataItem `json:"Business,omitnil,omitempty" name:"Business"`

	// 现金
	// 注意：此字段可能返回 null，表示取不到有效值。
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 赠送金
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 代金券
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 分成金
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`
}

type ConsumptionRegionSummaryDataItem struct {
	// 地域ID
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 折后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 趋势
	Trend *ConsumptionSummaryTrend `json:"Trend,omitnil,omitempty" name:"Trend"`

	// 产品消费详情
	Business []*ConsumptionBusinessSummaryDataItem `json:"Business,omitnil,omitempty" name:"Business"`

	// 现金
	// 注意：此字段可能返回 null，表示取不到有效值。
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 代金券
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 赠送金
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 分成金
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`
}

type ConsumptionResourceSummaryConditionValue struct {
	// 产品列表
	Business []*ConditionBusiness `json:"Business,omitnil,omitempty" name:"Business"`

	// 项目列表
	Project []*ConditionProject `json:"Project,omitnil,omitempty" name:"Project"`

	// 地域列表
	Region []*ConditionRegion `json:"Region,omitnil,omitempty" name:"Region"`

	// 付费模式列表
	PayMode []*ConditionPayMode `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

type ConsumptionResourceSummaryDataItem struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源名称
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 折后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 现金花费
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 地域ID
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 付费模式
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 付费模式名称
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// 产品名称代码
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 产品名称
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// 消耗类型
	ConsumptionTypeName *string `json:"ConsumptionTypeName,omitnil,omitempty" name:"ConsumptionTypeName"`

	// 折前价
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealCost *string `json:"RealCost,omitnil,omitempty" name:"RealCost"`

	// 费用起始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeeBeginTime *string `json:"FeeBeginTime,omitnil,omitempty" name:"FeeBeginTime"`

	// 费用结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeeEndTime *string `json:"FeeEndTime,omitnil,omitempty" name:"FeeEndTime"`

	// 天数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DayDiff *string `json:"DayDiff,omitnil,omitempty" name:"DayDiff"`

	// 每日消耗
	// 注意：此字段可能返回 null，表示取不到有效值。
	DailyTotalCost *string `json:"DailyTotalCost,omitnil,omitempty" name:"DailyTotalCost"`

	// 订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 代金券
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 赠送金
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 分成金
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// 支付者UIN：支付者的账号 ID，账号 ID 是用户在腾讯云的唯一账号标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 使用者UIN：实际使用资源的账号 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 操作者UIN：操作者账号 ID（预付费资源下单或后付费操作开通资源账号的 ID 或者角色 ID ）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// 子产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 子产品名称：用户采购的具体产品细分类型，例如：云服务器 CVM-标准型 S1
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// 地域类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionType *string `json:"RegionType,omitnil,omitempty" name:"RegionType"`

	// 地域类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionTypeName *string `json:"RegionTypeName,omitnil,omitempty" name:"RegionTypeName"`

	// 扩展字段1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extend1 *string `json:"Extend1,omitnil,omitempty" name:"Extend1"`

	// 扩展字段2
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extend2 *string `json:"Extend2,omitnil,omitempty" name:"Extend2"`

	// 扩展字段3
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extend3 *string `json:"Extend3,omitnil,omitempty" name:"Extend3"`

	// 扩展字段4
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extend4 *string `json:"Extend4,omitnil,omitempty" name:"Extend4"`

	// 扩展字段5
	// 注意：此字段可能返回 null，表示取不到有效值。
	Extend5 *string `json:"Extend5,omitnil,omitempty" name:"Extend5"`

	// 实例类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTypeName *string `json:"InstanceTypeName,omitnil,omitempty" name:"InstanceTypeName"`

	// 扣费时间：结算扣费时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayTime *string `json:"PayTime,omitnil,omitempty" name:"PayTime"`

	// 可用区：资源所属可用区，如广州三区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 配置描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentConfig *string `json:"ComponentConfig,omitnil,omitempty" name:"ComponentConfig"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ConsumptionSummaryTotal struct {
	// 折后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`
}

type ConsumptionSummaryTrend struct {
	// 趋势类型，upward上升/downward下降/none无
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 趋势值，Type为none是该字段值为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type CosDetailSets struct {
	// 存储桶名称
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 用量开始时间
	DosageBeginTime *string `json:"DosageBeginTime,omitnil,omitempty" name:"DosageBeginTime"`

	// 用量结束时间
	DosageEndTime *string `json:"DosageEndTime,omitnil,omitempty" name:"DosageEndTime"`

	// 子产品名称
	SubProductCodeName *string `json:"SubProductCodeName,omitnil,omitempty" name:"SubProductCodeName"`

	// 计费项名称
	BillingItemCodeName *string `json:"BillingItemCodeName,omitnil,omitempty" name:"BillingItemCodeName"`

	// 用量
	DosageValue *string `json:"DosageValue,omitnil,omitempty" name:"DosageValue"`

	// 单位
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`
}

type CostComponentSet struct {
	// 组件类型名称
	ComponentCodeName *string `json:"ComponentCodeName,omitnil,omitempty" name:"ComponentCodeName"`

	// 组件名称
	ItemCodeName *string `json:"ItemCodeName,omitnil,omitempty" name:"ItemCodeName"`

	// 刊例价
	SinglePrice *string `json:"SinglePrice,omitnil,omitempty" name:"SinglePrice"`

	// 刊例价单位
	PriceUnit *string `json:"PriceUnit,omitnil,omitempty" name:"PriceUnit"`

	// 用量
	UsedAmount *string `json:"UsedAmount,omitnil,omitempty" name:"UsedAmount"`

	// 用量单位
	UsedAmountUnit *string `json:"UsedAmountUnit,omitnil,omitempty" name:"UsedAmountUnit"`

	// 原价
	Cost *string `json:"Cost,omitnil,omitempty" name:"Cost"`

	// 折扣
	Discount *string `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 折后价
	RealCost *string `json:"RealCost,omitnil,omitempty" name:"RealCost"`

	// 代金券支付金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 现金支付金额
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 赠送金支付金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`
}

type CostDetail struct {
	// 支付者uin
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 产品名称
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// 子产品名称
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// 计费模式名称
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 区域名称
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 地区名称
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源名称
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// 订单id
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 交易id
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// 费用开始时间
	FeeBeginTime *string `json:"FeeBeginTime,omitnil,omitempty" name:"FeeBeginTime"`

	// 费用结束时间
	FeeEndTime *string `json:"FeeEndTime,omitnil,omitempty" name:"FeeEndTime"`

	// 组件明细
	ComponentSet []*CostComponentSet `json:"ComponentSet,omitnil,omitempty" name:"ComponentSet"`

	// 子产品名称代码
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`
}

// Predefined struct for user
type CreateAllocationTagRequestParams struct {
	// 用户分账标签键
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
}

type CreateAllocationTagRequest struct {
	*tchttp.BaseRequest
	
	// 用户分账标签键
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 区域编码
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 预付费类型
	PrePayType *string `json:"PrePayType,omitnil,omitempty" name:"PrePayType"`

	// 时长
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 时长单位
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 商品唯一标识
	CommodityCode *string `json:"CommodityCode,omitnil,omitempty" name:"CommodityCode"`

	// 承诺时长内的小额金额（单位：元）
	PromiseUseAmount *uint64 `json:"PromiseUseAmount,omitnil,omitempty" name:"PromiseUseAmount"`

	// 节省计划的指定生效时间，若不传则为当前下单时间。传参数格式:"2023-10-01 00:00:00"，仅支持指定日期的0点时刻
	SpecifyEffectTime *string `json:"SpecifyEffectTime,omitnil,omitempty" name:"SpecifyEffectTime"`

	// 可重入ID
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`
}

type CreateSavingPlanOrderRequest struct {
	*tchttp.BaseRequest
	
	// 地域编码
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 区域编码
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 预付费类型
	PrePayType *string `json:"PrePayType,omitnil,omitempty" name:"PrePayType"`

	// 时长
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 时长单位
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 商品唯一标识
	CommodityCode *string `json:"CommodityCode,omitnil,omitempty" name:"CommodityCode"`

	// 承诺时长内的小额金额（单位：元）
	PromiseUseAmount *uint64 `json:"PromiseUseAmount,omitnil,omitempty" name:"PromiseUseAmount"`

	// 节省计划的指定生效时间，若不传则为当前下单时间。传参数格式:"2023-10-01 00:00:00"，仅支持指定日期的0点时刻
	SpecifyEffectTime *string `json:"SpecifyEffectTime,omitnil,omitempty" name:"SpecifyEffectTime"`

	// 可重入ID
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`
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
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 订单的状态 1：未支付 2：已支付3：发货中 4：已发货 5：发货失败 6：已退款 7：已关单 8：订单过期 9：订单已失效 10：产品已失效 11：代付拒绝 12：支付中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 支付者
	Payer *string `json:"Payer,omitnil,omitempty" name:"Payer"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 创建人
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 实际支付金额（分）
	RealTotalCost *int64 `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 代金券抵扣金额（分）
	VoucherDecline *int64 `json:"VoucherDecline,omitnil,omitempty" name:"VoucherDecline"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 产品分类ID
	GoodsCategoryId *int64 `json:"GoodsCategoryId,omitnil,omitempty" name:"GoodsCategoryId"`

	// 产品详情
	ProductInfo []*ProductInfo `json:"ProductInfo,omitnil,omitempty" name:"ProductInfo"`

	// 时长
	TimeSpan *float64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 时间单位
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 货币单位
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 折扣率
	Policy *float64 `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 单价（分）
	Price *float64 `json:"Price,omitnil,omitempty" name:"Price"`

	// 原价（分）
	TotalCost *float64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 子产品编码
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// 大订单号
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`

	// 退费公式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Formula *string `json:"Formula,omitnil,omitempty" name:"Formula"`

	// 退费涉及订单信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefReturnDeals *string `json:"RefReturnDeals,omitnil,omitempty" name:"RefReturnDeals"`

	// 付费模式：prePay 预付费 postPay后付费 riPay预留实例
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

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
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// 产品编码中文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 子产品编码中文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductName *string `json:"SubProductName,omitnil,omitempty" name:"SubProductName"`

	// 订单对应的资源id, 查询参数Limit超过200，将返回null
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId []*string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

// Predefined struct for user
type DeleteAllocationTagRequestParams struct {
	// 用户分账标签键
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
}

type DeleteAllocationTagRequest struct {
	*tchttp.BaseRequest
	
	// 用户分账标签键
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Balance *int64 `json:"Balance,omitnil,omitempty" name:"Balance"`

	// 查询的用户Uin
	Uin *uint64 `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 当前真实可用余额,单位 分
	RealBalance *float64 `json:"RealBalance,omitnil,omitempty" name:"RealBalance"`

	// 现金账户余额,单位 分
	CashAccountBalance *float64 `json:"CashAccountBalance,omitnil,omitempty" name:"CashAccountBalance"`

	// 收益转入账户余额,单位 分
	IncomeIntoAccountBalance *float64 `json:"IncomeIntoAccountBalance,omitnil,omitempty" name:"IncomeIntoAccountBalance"`

	// 赠送账户余额,单位 分
	PresentAccountBalance *float64 `json:"PresentAccountBalance,omitnil,omitempty" name:"PresentAccountBalance"`

	// 冻结金额,单位 分
	FreezeAmount *float64 `json:"FreezeAmount,omitnil,omitempty" name:"FreezeAmount"`

	// 欠费金额,单位 分
	OweAmount *float64 `json:"OweAmount,omitnil,omitempty" name:"OweAmount"`

	// 是否允许欠费消费
	//
	// Deprecated: IsAllowArrears is deprecated.
	IsAllowArrears *bool `json:"IsAllowArrears,omitnil,omitempty" name:"IsAllowArrears"`

	// 是否限制信用额度
	//
	// Deprecated: IsCreditLimited is deprecated.
	IsCreditLimited *bool `json:"IsCreditLimited,omitnil,omitempty" name:"IsCreditLimited"`

	// 信用额度,单位 分
	CreditAmount *float64 `json:"CreditAmount,omitnil,omitempty" name:"CreditAmount"`

	// 可用信用额度,单位 分
	CreditBalance *float64 `json:"CreditBalance,omitnil,omitempty" name:"CreditBalance"`

	// 真实可用信用额度,单位 分
	RealCreditBalance *float64 `json:"RealCreditBalance,omitnil,omitempty" name:"RealCreditBalance"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeAllocateConditionsRequestParams struct {
	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

type DescribeAllocateConditionsRequest struct {
	*tchttp.BaseRequest
	
	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

func (r *DescribeAllocateConditionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocateConditionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocateConditionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocateConditionsResponseParams struct {
	// 产品筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Business []*BillBusinessLink `json:"Business,omitnil,omitempty" name:"Business"`

	// 子产品筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Product []*BillProduct `json:"Product,omitnil,omitempty" name:"Product"`

	// 组件名称筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Item []*BillItem `json:"Item,omitnil,omitempty" name:"Item"`

	// 地域筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region []*BillRegion `json:"Region,omitnil,omitempty" name:"Region"`

	// 实例类型筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType []*BillInstanceType `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 计费模式筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode []*BillPayMode `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 项目筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Project []*BillProject `json:"Project,omitnil,omitempty" name:"Project"`

	// 标签筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*BillTag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 使用者 UIN 筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin []*BillOwnerUin `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 操作者 UIN 筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateUin []*BillOperateUin `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// 交易类型筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionType []*BillActionType `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocateConditionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocateConditionsResponseParams `json:"Response"`
}

func (r *DescribeAllocateConditionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocateConditionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationBillConditionsRequestParams struct {
	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 分账单元唯一标识，用作筛选
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// 日期
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`

	// 产品编码
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// 使用者UIN
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// 操作者UIN
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// 计费模式编码
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// 交易类型编码
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// 子产品编码
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// 地域ID
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// 可用区ID
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 实例类型编码
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 标签
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 组件类型编码
	ComponentCodes []*string `json:"ComponentCodes,omitnil,omitempty" name:"ComponentCodes"`

	// 组件名称编码
	ItemCodes []*string `json:"ItemCodes,omitnil,omitempty" name:"ItemCodes"`

	// 模糊搜索条件
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 项目id
	ProjectIds []*uint64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 费用归集类型
	AllocationType []*int64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`
}

type DescribeAllocationBillConditionsRequest struct {
	*tchttp.BaseRequest
	
	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 分账单元唯一标识，用作筛选
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// 日期
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`

	// 产品编码
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// 使用者UIN
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// 操作者UIN
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// 计费模式编码
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// 交易类型编码
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// 子产品编码
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// 地域ID
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// 可用区ID
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 实例类型编码
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 标签
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 组件类型编码
	ComponentCodes []*string `json:"ComponentCodes,omitnil,omitempty" name:"ComponentCodes"`

	// 组件名称编码
	ItemCodes []*string `json:"ItemCodes,omitnil,omitempty" name:"ItemCodes"`

	// 模糊搜索条件
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 项目id
	ProjectIds []*uint64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 费用归集类型
	AllocationType []*int64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`
}

func (r *DescribeAllocationBillConditionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationBillConditionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "TreeNodeUniqKeys")
	delete(f, "BillDates")
	delete(f, "BusinessCodes")
	delete(f, "OwnerUins")
	delete(f, "OperateUins")
	delete(f, "PayModes")
	delete(f, "ActionTypes")
	delete(f, "ProductCodes")
	delete(f, "RegionIds")
	delete(f, "ZoneIds")
	delete(f, "InstanceTypes")
	delete(f, "Tag")
	delete(f, "ComponentCodes")
	delete(f, "ItemCodes")
	delete(f, "SearchKey")
	delete(f, "ProjectIds")
	delete(f, "AllocationType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationBillConditionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationBillConditionsResponseParams struct {
	// 产品筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Business []*BillBusiness `json:"Business,omitnil,omitempty" name:"Business"`

	// 子产品筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Product []*BillProduct `json:"Product,omitnil,omitempty" name:"Product"`

	// 组件名称筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Item []*BillItem `json:"Item,omitnil,omitempty" name:"Item"`

	// 地域筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region []*BillRegion `json:"Region,omitnil,omitempty" name:"Region"`

	// 实例类型筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType []*BillInstanceType `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 计费模式筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode []*BillPayMode `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 项目筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Project []*BillProject `json:"Project,omitnil,omitempty" name:"Project"`

	// 标签筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*BillTag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 使用者 UIN 筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin []*BillOwnerUin `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 操作者 UIN 筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateUin []*BillOperateUin `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// 日期筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillDay []*BillDays `json:"BillDay,omitnil,omitempty" name:"BillDay"`

	// 交易类型筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionType []*BillActionType `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 组件类型筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Component []*BillComponent `json:"Component,omitnil,omitempty" name:"Component"`

	// 可用区筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone []*BillZoneId `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 分账单元筛选列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllocationTreeNode []*AllocationTreeNode `json:"AllocationTreeNode,omitnil,omitempty" name:"AllocationTreeNode"`

	// 分账标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationBillConditionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationBillConditionsResponseParams `json:"Response"`
}

func (r *DescribeAllocationBillConditionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationBillConditionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationBillDetailRequestParams struct {
	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，以此类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 分账单元唯一标识，用作筛选
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// 排序字段，枚举值如下：
	// RiTimeSpan - 预留实例抵扣时长
	// ExtendPayAmount1 - 预留实例抵扣组件原价
	// RealCost - 折后总价
	// CashPayAmount - 现金金额
	// VoucherPayAmount - 代金券金额
	// IncentivePayAmount - 赠送金金额
	// TransferPayAmount -分成金金额
	// Cost - 组件原价
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序类型，枚举值如下：
	// asc - 升序
	// desc - 降序
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 产品编码，用作筛选
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// 使用者UIN，用作筛选
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// 操作者UIN，用作筛选
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// 计费模式编码，用作筛选
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// 交易类型编码，用作筛选
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// 子产品编码，用作筛选
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// 地域ID，用作筛选
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// 可用区ID，用作筛选
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 实例类型编码，用作筛选
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 标签，用作筛选
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 组件类型编码，用作筛选
	ComponentCodes []*string `json:"ComponentCodes,omitnil,omitempty" name:"ComponentCodes"`

	// 组件名称编码，用作筛选
	ItemCodes []*string `json:"ItemCodes,omitnil,omitempty" name:"ItemCodes"`

	// 模糊搜索：支持标签、资源id、资源别名
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 项目ID，用作筛选
	ProjectIds []*uint64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`
}

type DescribeAllocationBillDetailRequest struct {
	*tchttp.BaseRequest
	
	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，以此类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 分账单元唯一标识，用作筛选
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// 排序字段，枚举值如下：
	// RiTimeSpan - 预留实例抵扣时长
	// ExtendPayAmount1 - 预留实例抵扣组件原价
	// RealCost - 折后总价
	// CashPayAmount - 现金金额
	// VoucherPayAmount - 代金券金额
	// IncentivePayAmount - 赠送金金额
	// TransferPayAmount -分成金金额
	// Cost - 组件原价
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序类型，枚举值如下：
	// asc - 升序
	// desc - 降序
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 产品编码，用作筛选
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// 使用者UIN，用作筛选
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// 操作者UIN，用作筛选
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// 计费模式编码，用作筛选
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// 交易类型编码，用作筛选
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// 子产品编码，用作筛选
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// 地域ID，用作筛选
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// 可用区ID，用作筛选
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 实例类型编码，用作筛选
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 标签，用作筛选
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 组件类型编码，用作筛选
	ComponentCodes []*string `json:"ComponentCodes,omitnil,omitempty" name:"ComponentCodes"`

	// 组件名称编码，用作筛选
	ItemCodes []*string `json:"ItemCodes,omitnil,omitempty" name:"ItemCodes"`

	// 模糊搜索：支持标签、资源id、资源别名
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 项目ID，用作筛选
	ProjectIds []*uint64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`
}

func (r *DescribeAllocationBillDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationBillDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Month")
	delete(f, "TreeNodeUniqKeys")
	delete(f, "Sort")
	delete(f, "SortType")
	delete(f, "BusinessCodes")
	delete(f, "OwnerUins")
	delete(f, "OperateUins")
	delete(f, "PayModes")
	delete(f, "ActionTypes")
	delete(f, "ProductCodes")
	delete(f, "RegionIds")
	delete(f, "ZoneIds")
	delete(f, "InstanceTypes")
	delete(f, "Tag")
	delete(f, "ComponentCodes")
	delete(f, "ItemCodes")
	delete(f, "SearchKey")
	delete(f, "ProjectIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationBillDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationBillDetailResponseParams struct {
	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// 分账账单概览金额汇总
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *AllocationOverviewTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// 分账账单明细
	Detail []*AllocationDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationBillDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationBillDetailResponseParams `json:"Response"`
}

func (r *DescribeAllocationBillDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationBillDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationMonthOverviewRequestParams struct {
	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

type DescribeAllocationMonthOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`
}

func (r *DescribeAllocationMonthOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationMonthOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationMonthOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationMonthOverviewResponseParams struct {
	// 分账账单月概览详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*AllocationOverviewNode `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 分账账单概览金额汇总
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *AllocationOverviewTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationMonthOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationMonthOverviewResponseParams `json:"Response"`
}

func (r *DescribeAllocationMonthOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationMonthOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationOverviewRequestParams struct {
	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，以此类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 统计周期，枚举值如下
	// month - 月
	// day - 日
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 分账单元唯一标识，用作筛选
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// 排序字段，枚举值如下： 
	// GatherCashPayAmount - 归集费用(现金)
	// GatherVoucherPayAmount- 归集费用(优惠券)
	// GatherIncentivePayAmount -  归集费用(赠送金)
	// GatherTransferPayAmount - 归集费用(分成金)
	// AllocateCashPayAmount - 分摊费用(现金)
	// AllocateVoucherPayAmount - 分摊费用(优惠券)
	// AllocateIncentivePayAmount - 分摊费用(赠送金)
	// AllocateTransferPayAmount - 分摊费用(分成金)
	// TotalCashPayAmount - 合计费用(现金)
	// TotalVoucherPayAmount - 合计费用(优惠券)
	// TotalIncentivePayAmount - 合计费用(赠送金)
	// TotalTransferPayAmount - 合计费用(分成金)
	// GatherRealCost - 归集费用(折后总额)
	// AllocateRealCost - 分摊费用(折后总额)
	// RealTotalCost - 合计费用(折后总额)
	// Ratio  - 占比(折后总额)
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序类型，枚举值如下：
	// asc - 升序
	// desc - 降序
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 日期，用作筛选
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`
}

type DescribeAllocationOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，以此类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 统计周期，枚举值如下
	// month - 月
	// day - 日
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 分账单元唯一标识，用作筛选
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// 排序字段，枚举值如下： 
	// GatherCashPayAmount - 归集费用(现金)
	// GatherVoucherPayAmount- 归集费用(优惠券)
	// GatherIncentivePayAmount -  归集费用(赠送金)
	// GatherTransferPayAmount - 归集费用(分成金)
	// AllocateCashPayAmount - 分摊费用(现金)
	// AllocateVoucherPayAmount - 分摊费用(优惠券)
	// AllocateIncentivePayAmount - 分摊费用(赠送金)
	// AllocateTransferPayAmount - 分摊费用(分成金)
	// TotalCashPayAmount - 合计费用(现金)
	// TotalVoucherPayAmount - 合计费用(优惠券)
	// TotalIncentivePayAmount - 合计费用(赠送金)
	// TotalTransferPayAmount - 合计费用(分成金)
	// GatherRealCost - 归集费用(折后总额)
	// AllocateRealCost - 分摊费用(折后总额)
	// RealTotalCost - 合计费用(折后总额)
	// Ratio  - 占比(折后总额)
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序类型，枚举值如下：
	// asc - 升序
	// desc - 降序
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 日期，用作筛选
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`
}

func (r *DescribeAllocationOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Month")
	delete(f, "PeriodType")
	delete(f, "TreeNodeUniqKeys")
	delete(f, "Sort")
	delete(f, "SortType")
	delete(f, "BillDates")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationOverviewResponseParams struct {
	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// 分账账单概览金额汇总
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *AllocationOverviewTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// 分账概览明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*AllocationOverviewDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationOverviewResponseParams `json:"Response"`
}

func (r *DescribeAllocationOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationSummaryByBusinessRequestParams struct {
	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，以此类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 统计周期，枚举值如下
	// month - 月
	// day - 日
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 分账单元唯一标识，用作筛选
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// 排序类型，枚举值如下：
	// asc - 升序
	// desc - 降序
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 排序字段，枚举值如下：
	// GatherCashPayAmount - 归集费用(现金)
	// GatherVoucherPayAmount- 归集费用(优惠券)
	// GatherIncentivePayAmount - 归集费用(赠送金)
	// GatherTransferPayAmount - 归集费用(分成金)
	// AllocateCashPayAmount - 分摊费用(现金)
	// AllocateVoucherPayAmount - 分摊费用(优惠券)
	// AllocateIncentivePayAmount - 分摊费用(赠送金)
	// AllocateTransferPayAmount - 分摊费用(分成金)
	// TotalCashPayAmount - 合计费用(现金)
	// TotalVoucherPayAmount - 合计费用(优惠券)
	// TotalIncentivePayAmount - 合计费用(赠送金)
	// TotalTransferPayAmount - 合计费用(分成金)
	// GatherRealCost - 归集费用(折后总额)
	// AllocateRealCost - 分摊费用(折后总额)
	// RealTotalCost - 合计费用(折后总额)
	// BusinessCode - 产品代码
	// Ratio - 占比(折后总额)
	// Trend - 环比(折后总额)
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 日期，用作筛选，PeriodType=day时可传
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`

	// 产品编码，用作筛选
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// 模糊搜索条件
	//
	// Deprecated: SearchKey is deprecated.
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

type DescribeAllocationSummaryByBusinessRequest struct {
	*tchttp.BaseRequest
	
	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，以此类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 统计周期，枚举值如下
	// month - 月
	// day - 日
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 分账单元唯一标识，用作筛选
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// 排序类型，枚举值如下：
	// asc - 升序
	// desc - 降序
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 排序字段，枚举值如下：
	// GatherCashPayAmount - 归集费用(现金)
	// GatherVoucherPayAmount- 归集费用(优惠券)
	// GatherIncentivePayAmount - 归集费用(赠送金)
	// GatherTransferPayAmount - 归集费用(分成金)
	// AllocateCashPayAmount - 分摊费用(现金)
	// AllocateVoucherPayAmount - 分摊费用(优惠券)
	// AllocateIncentivePayAmount - 分摊费用(赠送金)
	// AllocateTransferPayAmount - 分摊费用(分成金)
	// TotalCashPayAmount - 合计费用(现金)
	// TotalVoucherPayAmount - 合计费用(优惠券)
	// TotalIncentivePayAmount - 合计费用(赠送金)
	// TotalTransferPayAmount - 合计费用(分成金)
	// GatherRealCost - 归集费用(折后总额)
	// AllocateRealCost - 分摊费用(折后总额)
	// RealTotalCost - 合计费用(折后总额)
	// BusinessCode - 产品代码
	// Ratio - 占比(折后总额)
	// Trend - 环比(折后总额)
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 日期，用作筛选，PeriodType=day时可传
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`

	// 产品编码，用作筛选
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// 模糊搜索条件
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`
}

func (r *DescribeAllocationSummaryByBusinessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationSummaryByBusinessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Month")
	delete(f, "PeriodType")
	delete(f, "TreeNodeUniqKeys")
	delete(f, "SortType")
	delete(f, "Sort")
	delete(f, "BillDates")
	delete(f, "BusinessCodes")
	delete(f, "SearchKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationSummaryByBusinessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationSummaryByBusinessResponseParams struct {
	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// 分账账单概览金额汇总
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *AllocationOverviewTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// 分账账单按产品汇总明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*AllocationSummaryByBusiness `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationSummaryByBusinessResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationSummaryByBusinessResponseParams `json:"Response"`
}

func (r *DescribeAllocationSummaryByBusinessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationSummaryByBusinessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationSummaryByItemRequestParams struct {
	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，以此类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 统计周期，枚举值如下
	// month - 月
	// day - 日
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 分账单元唯一标识，用作筛选
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// 排序字段，枚举值如下：
	// RiTimeSpan - 预留实例抵扣时长
	// ExtendPayAmount1 - 预留实例抵扣组件原价
	// RealCost - 折后总价
	// CashPayAmount - 现金金额
	// VoucherPayAmount - 代金券金额
	// IncentivePayAmount - 赠送金金额
	// TransferPayAmount -分成金金额
	// Cost - 组件原价
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序类型，枚举值如下：
	// asc - 升序
	// desc - 降序
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 日期，用作筛选
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`

	// 产品编码，用作筛选
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// 使用者UIN，用作筛选
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// 操作者UIN，用作筛选
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// 计费模式编码，用作筛选
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// 交易类型编码，用作筛选
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// 子产品编码，用作筛选
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// 地域ID，用作筛选
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// 可用区ID，用作筛选
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 实例类型编码，用作筛选
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 标签，用作筛选
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 组件类型编码，用作筛选
	ComponentCodes []*string `json:"ComponentCodes,omitnil,omitempty" name:"ComponentCodes"`

	// 组件名称编码，用作筛选
	ItemCodes []*string `json:"ItemCodes,omitnil,omitempty" name:"ItemCodes"`

	// 模糊搜索：支持标签、资源id、资源别名
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 项目ID，用作筛选
	ProjectIds []*uint64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 费用归集类型，枚举值如下：
	// 0 - 分摊
	// 1 - 归集
	// -1 - 未分配
	AllocationType []*int64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`
}

type DescribeAllocationSummaryByItemRequest struct {
	*tchttp.BaseRequest
	
	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，以此类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 统计周期，枚举值如下
	// month - 月
	// day - 日
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 分账单元唯一标识，用作筛选
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// 排序字段，枚举值如下：
	// RiTimeSpan - 预留实例抵扣时长
	// ExtendPayAmount1 - 预留实例抵扣组件原价
	// RealCost - 折后总价
	// CashPayAmount - 现金金额
	// VoucherPayAmount - 代金券金额
	// IncentivePayAmount - 赠送金金额
	// TransferPayAmount -分成金金额
	// Cost - 组件原价
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序类型，枚举值如下：
	// asc - 升序
	// desc - 降序
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 日期，用作筛选
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`

	// 产品编码，用作筛选
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// 使用者UIN，用作筛选
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// 操作者UIN，用作筛选
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// 计费模式编码，用作筛选
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// 交易类型编码，用作筛选
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// 子产品编码，用作筛选
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// 地域ID，用作筛选
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// 可用区ID，用作筛选
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 实例类型编码，用作筛选
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 标签，用作筛选
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 组件类型编码，用作筛选
	ComponentCodes []*string `json:"ComponentCodes,omitnil,omitempty" name:"ComponentCodes"`

	// 组件名称编码，用作筛选
	ItemCodes []*string `json:"ItemCodes,omitnil,omitempty" name:"ItemCodes"`

	// 模糊搜索：支持标签、资源id、资源别名
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 项目ID，用作筛选
	ProjectIds []*uint64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 费用归集类型，枚举值如下：
	// 0 - 分摊
	// 1 - 归集
	// -1 - 未分配
	AllocationType []*int64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`
}

func (r *DescribeAllocationSummaryByItemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationSummaryByItemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Month")
	delete(f, "PeriodType")
	delete(f, "TreeNodeUniqKeys")
	delete(f, "Sort")
	delete(f, "SortType")
	delete(f, "BillDates")
	delete(f, "BusinessCodes")
	delete(f, "OwnerUins")
	delete(f, "OperateUins")
	delete(f, "PayModes")
	delete(f, "ActionTypes")
	delete(f, "ProductCodes")
	delete(f, "RegionIds")
	delete(f, "ZoneIds")
	delete(f, "InstanceTypes")
	delete(f, "Tag")
	delete(f, "ComponentCodes")
	delete(f, "ItemCodes")
	delete(f, "SearchKey")
	delete(f, "ProjectIds")
	delete(f, "AllocationType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationSummaryByItemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationSummaryByItemResponseParams struct {
	// 总条数
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// 分账账单概览金额汇总
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *AllocationOverviewTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// 分账账单按组件汇总明细
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*AllocationSummaryByItem `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationSummaryByItemResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationSummaryByItemResponseParams `json:"Response"`
}

func (r *DescribeAllocationSummaryByItemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationSummaryByItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationSummaryByResourceRequestParams struct {
	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，以此类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 统计周期，枚举值如下
	// month - 月
	// day - 日
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 分账单元唯一标识，用作筛选
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// 排序字段，枚举值如下：
	// RiTimeSpan - 预留实例抵扣时长
	// ExtendPayAmount1 - 预留实例抵扣组件原价
	// RealCost - 折后总价
	// CashPayAmount - 现金金额
	// VoucherPayAmount - 代金券金额
	// IncentivePayAmount - 赠送金金额
	// TransferPayAmount -分成金金额
	// Cost - 组件原价
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序类型，枚举值如下：
	// asc - 升序
	// desc - 降序
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 日期，用作筛选
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`

	// 产品编码，用作筛选
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// 使用者UIN，用作筛选
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// 操作者UIN，用作筛选
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// 计费模式编码，用作筛选
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// 交易类型编码，用作筛选
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// 子产品编码，用作筛选
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// 地域ID，用作筛选
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// 可用区ID，用作筛选
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 实例类型编码，用作筛选
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 标签，用作筛选
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 模糊搜索：支持标签、资源id、资源别名
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 项目ID，用作筛选
	ProjectIds []*uint64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 费用归集类型，枚举值如下：
	// 0 - 分摊 
	// 1 - 归集 
	// -1 -  未分配
	AllocationType []*int64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`
}

type DescribeAllocationSummaryByResourceRequest struct {
	*tchttp.BaseRequest
	
	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，以此类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 统计周期，枚举值如下
	// month - 月
	// day - 日
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 分账单元唯一标识，用作筛选
	TreeNodeUniqKeys []*string `json:"TreeNodeUniqKeys,omitnil,omitempty" name:"TreeNodeUniqKeys"`

	// 排序字段，枚举值如下：
	// RiTimeSpan - 预留实例抵扣时长
	// ExtendPayAmount1 - 预留实例抵扣组件原价
	// RealCost - 折后总价
	// CashPayAmount - 现金金额
	// VoucherPayAmount - 代金券金额
	// IncentivePayAmount - 赠送金金额
	// TransferPayAmount -分成金金额
	// Cost - 组件原价
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序类型，枚举值如下：
	// asc - 升序
	// desc - 降序
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 日期，用作筛选
	BillDates []*string `json:"BillDates,omitnil,omitempty" name:"BillDates"`

	// 产品编码，用作筛选
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// 使用者UIN，用作筛选
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// 操作者UIN，用作筛选
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// 计费模式编码，用作筛选
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// 交易类型编码，用作筛选
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`

	// 子产品编码，用作筛选
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// 地域ID，用作筛选
	RegionIds []*string `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// 可用区ID，用作筛选
	ZoneIds []*string `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 实例类型编码，用作筛选
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 标签，用作筛选
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 模糊搜索：支持标签、资源id、资源别名
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 项目ID，用作筛选
	ProjectIds []*uint64 `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 费用归集类型，枚举值如下：
	// 0 - 分摊 
	// 1 - 归集 
	// -1 -  未分配
	AllocationType []*int64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`
}

func (r *DescribeAllocationSummaryByResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationSummaryByResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Month")
	delete(f, "PeriodType")
	delete(f, "TreeNodeUniqKeys")
	delete(f, "Sort")
	delete(f, "SortType")
	delete(f, "BillDates")
	delete(f, "BusinessCodes")
	delete(f, "OwnerUins")
	delete(f, "OperateUins")
	delete(f, "PayModes")
	delete(f, "ActionTypes")
	delete(f, "ProductCodes")
	delete(f, "RegionIds")
	delete(f, "ZoneIds")
	delete(f, "InstanceTypes")
	delete(f, "Tag")
	delete(f, "SearchKey")
	delete(f, "ProjectIds")
	delete(f, "AllocationType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationSummaryByResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationSummaryByResourceResponseParams struct {
	// 总条数
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// 分账账单概览金额汇总
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *AllocationOverviewTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// 分账账单按资源汇总明细
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*AllocationSummaryByResource `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationSummaryByResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationSummaryByResourceResponseParams `json:"Response"`
}

func (r *DescribeAllocationSummaryByResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationSummaryByResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationTrendByMonthRequestParams struct {
	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 分账单元唯一标识
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// 产品编码，用作筛选
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`
}

type DescribeAllocationTrendByMonthRequest struct {
	*tchttp.BaseRequest
	
	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 分账单元唯一标识
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// 产品编码，用作筛选
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`
}

func (r *DescribeAllocationTrendByMonthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationTrendByMonthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "TreeNodeUniqKey")
	delete(f, "BusinessCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllocationTrendByMonthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllocationTrendByMonthResponseParams struct {
	// 当月费用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Current *AllocationBillTrendDetail `json:"Current,omitnil,omitempty" name:"Current"`

	// 之前月份费用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Previous []*AllocationBillTrendDetail `json:"Previous,omitnil,omitempty" name:"Previous"`

	// 费用统计信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Stat *AllocationStat `json:"Stat,omitnil,omitempty" name:"Stat"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllocationTrendByMonthResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllocationTrendByMonthResponseParams `json:"Response"`
}

func (r *DescribeAllocationTrendByMonthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllocationTrendByMonthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillAdjustInfoRequestParams struct {
	// 格式：yyyy-MM
	// 账单月份，month和timeFrom&timeTo必传一个，如果有传timeFrom&timeTo则month字段无效
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 格式：yyyy-MM-dd
	// 开始时间，month和timeFrom&timeTo必传一个，如果有该字段则month字段无效。timeFrom和timeTo必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据
	TimeFrom *string `json:"TimeFrom,omitnil,omitempty" name:"TimeFrom"`

	// 格式：yyyy-MM-dd
	// 截止时间，month和timeFrom&timeTo必传一个，如果有该字段则month字段无效。timeFrom和timeTo必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据
	TimeTo *string `json:"TimeTo,omitnil,omitempty" name:"TimeTo"`
}

type DescribeBillAdjustInfoRequest struct {
	*tchttp.BaseRequest
	
	// 格式：yyyy-MM
	// 账单月份，month和timeFrom&timeTo必传一个，如果有传timeFrom&timeTo则month字段无效
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 格式：yyyy-MM-dd
	// 开始时间，month和timeFrom&timeTo必传一个，如果有该字段则month字段无效。timeFrom和timeTo必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据
	TimeFrom *string `json:"TimeFrom,omitnil,omitempty" name:"TimeFrom"`

	// 格式：yyyy-MM-dd
	// 截止时间，month和timeFrom&timeTo必传一个，如果有该字段则month字段无效。timeFrom和timeTo必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据
	TimeTo *string `json:"TimeTo,omitnil,omitempty" name:"TimeTo"`
}

func (r *DescribeBillAdjustInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillAdjustInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Month")
	delete(f, "TimeFrom")
	delete(f, "TimeTo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillAdjustInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillAdjustInfoResponseParams struct {
	// 数据总量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 明细数据
	Data []*AdjustInfoDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillAdjustInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillAdjustInfoResponseParams `json:"Response"`
}

func (r *DescribeBillAdjustInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillAdjustInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillDetailForOrganizationRequestParams struct {
	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，依次类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量，最大值为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	//
	// Deprecated: PeriodType is deprecated.
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 月份，格式为yyyy-mm，Month和BeginTime&EndTime必传一个，如果有传BeginTime&EndTime则Month字段无效。最多可拉取近18个月内的数据。
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 周期开始时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据。最多可拉取18个月内的数据。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 周期结束时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据。最多可拉取近18个月内的数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// 付费模式 prePay(表示包年包月)/postPay(表示按时按量)
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

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
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 项目ID:资源所属项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 上一次请求返回的上下文信息，翻页查询Month>=2023-05的月份的数据可加快查询速度，数据量10万级别以上的用户建议使用，查询速度可提升2~10倍
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`
}

type DescribeBillDetailForOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，依次类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量，最大值为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 月份，格式为yyyy-mm，Month和BeginTime&EndTime必传一个，如果有传BeginTime&EndTime则Month字段无效。最多可拉取近18个月内的数据。
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 周期开始时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据。最多可拉取18个月内的数据。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 周期结束时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据。最多可拉取近18个月内的数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// 付费模式 prePay(表示包年包月)/postPay(表示按时按量)
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

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
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 项目ID:资源所属项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 上一次请求返回的上下文信息，翻页查询Month>=2023-05的月份的数据可加快查询速度，数据量10万级别以上的用户建议使用，查询速度可提升2~10倍
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`
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
	DetailSet []*DistributionBillDetail `json:"DetailSet,omitnil,omitempty" name:"DetailSet"`

	// 总记录数，24小时缓存一次，可能比实际总记录数少
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 本次请求的上下文信息，可用于下一次请求的请求参数中，加快查询速度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量，最大值为300
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	//
	// Deprecated: PeriodType is deprecated.
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 月份，格式为yyyy-mm，Month和BeginTime&EndTime必传一个，如果有传BeginTime&EndTime则Month字段无效。最多可拉取近18个月内的数据。
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 周期开始时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据。最多可拉取18个月内的数据。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 周期结束时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据。最多可拉取近18个月内的数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// 已废弃参数，未开放
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 付费模式 prePay(表示包年包月)/postPay(表示按时按量)
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

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
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 项目ID:资源所属项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 上一次请求返回的上下文信息，翻页查询Month>=2023-05的月份的数据可加快查询速度，数据量10万级别以上的用户建议使用，查询速度可提升2~10倍
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 支付者的账号 ID（账号 ID 是用户在腾讯云的唯一账号标识），默认查询本账号账单，如集团管理账号需查询成员账号自付的账单，该字段需入参成员账号UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
}

type DescribeBillDetailRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，依次类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量，最大值为300
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 月份，格式为yyyy-mm，Month和BeginTime&EndTime必传一个，如果有传BeginTime&EndTime则Month字段无效。最多可拉取近18个月内的数据。
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 周期开始时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据。最多可拉取18个月内的数据。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 周期结束时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为相同月份，不支持跨月查询，查询结果是整月数据。最多可拉取近18个月内的数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// 已废弃参数，未开放
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 付费模式 prePay(表示包年包月)/postPay(表示按时按量)
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

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
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 项目ID:资源所属项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 上一次请求返回的上下文信息，翻页查询Month>=2023-05的月份的数据可加快查询速度，数据量10万级别以上的用户建议使用，查询速度可提升2~10倍
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 支付者的账号 ID（账号 ID 是用户在腾讯云的唯一账号标识），默认查询本账号账单，如集团管理账号需查询成员账号自付的账单，该字段需入参成员账号UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
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
	DetailSet []*BillDetail `json:"DetailSet,omitnil,omitempty" name:"DetailSet"`

	// 总记录数，24小时缓存一次，可能比实际总记录数少
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 本次请求的上下文信息，可用于下一次请求的请求参数中，加快查询速度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 账单月份
	// 支持的最早开始月份为2021-01
	// L0-PDF&账单包不支持当月下载，当月账单请在次月1号19:00出账后下载
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 下载的账号 ID列表，默认查询本账号账单，如集团管理账号需下载成员账号自付的账单，该字段需入参成员账号UIN
	ChildUin []*string `json:"ChildUin,omitnil,omitempty" name:"ChildUin"`
}

type DescribeBillDownloadUrlRequest struct {
	*tchttp.BaseRequest
	
	// 账单类型，枚举值
	// billOverview=L0-PDF账单
	// billSummary=L1-汇总账单	
	// billResource=L2-资源账单	
	// billDetail=L3-明细账单	
	// billPack=账单包
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 账单月份
	// 支持的最早开始月份为2021-01
	// L0-PDF&账单包不支持当月下载，当月账单请在次月1号19:00出账后下载
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 下载的账号 ID列表，默认查询本账号账单，如集团管理账号需下载成员账号自付的账单，该字段需入参成员账号UIN
	ChildUin []*string `json:"ChildUin,omitnil,omitempty" name:"ChildUin"`
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
	Ready *int64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// 账单文件下载链接，有效时长为一天
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询范围的结束时间（包含）时间格式 yyyy-MM-dd HH:mm:ss ，开始时间和结束时间差值小于等于六个月
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 翻页偏移量，初始值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页的限制数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 交易类型： all所有交易类型，recharge充值，return退款，unblock解冻，agentin资金转入，advanced垫付，cash提现，deduct扣费，block冻结，agentout资金转出，repay垫付回款，repayment还款(仅国际信用账户)，adj_refund调增(仅国际信用账户)，adj_deduct调减(仅国际信用账户)
	PayType []*string `json:"PayType,omitnil,omitempty" name:"PayType"`

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
	SubPayType []*string `json:"SubPayType,omitnil,omitempty" name:"SubPayType"`

	// 是否返回0元交易金额的交易项，取值：0-不返回，1-返回。不传该参数则不返回
	WithZeroAmount *uint64 `json:"WithZeroAmount,omitnil,omitempty" name:"WithZeroAmount"`
}

type DescribeBillListRequest struct {
	*tchttp.BaseRequest
	
	// 查询范围的起始时间（包含）时间格式 yyyy-MM-dd HH:mm:ss 开始时间和结束时间差值小于等于六个月
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询范围的结束时间（包含）时间格式 yyyy-MM-dd HH:mm:ss ，开始时间和结束时间差值小于等于六个月
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 翻页偏移量，初始值为0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页的限制数量
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 交易类型： all所有交易类型，recharge充值，return退款，unblock解冻，agentin资金转入，advanced垫付，cash提现，deduct扣费，block冻结，agentout资金转出，repay垫付回款，repayment还款(仅国际信用账户)，adj_refund调增(仅国际信用账户)，adj_deduct调减(仅国际信用账户)
	PayType []*string `json:"PayType,omitnil,omitempty" name:"PayType"`

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
	SubPayType []*string `json:"SubPayType,omitnil,omitempty" name:"SubPayType"`

	// 是否返回0元交易金额的交易项，取值：0-不返回，1-返回。不传该参数则不返回
	WithZeroAmount *uint64 `json:"WithZeroAmount,omitnil,omitempty" name:"WithZeroAmount"`
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
	TransactionList []*BillTransactionInfo `json:"TransactionList,omitnil,omitempty" name:"TransactionList"`

	// 总条数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 退费总额，单位（分）
	ReturnAmount *float64 `json:"ReturnAmount,omitnil,omitempty" name:"ReturnAmount"`

	// 充值总额，单位（分）
	RechargeAmount *float64 `json:"RechargeAmount,omitnil,omitempty" name:"RechargeAmount"`

	// 冻结总额，单位（分）
	BlockAmount *float64 `json:"BlockAmount,omitnil,omitempty" name:"BlockAmount"`

	// 解冻总额，单位（分）
	UnblockAmount *float64 `json:"UnblockAmount,omitnil,omitempty" name:"UnblockAmount"`

	// 扣费总额，单位（分）
	DeductAmount *float64 `json:"DeductAmount,omitnil,omitempty" name:"DeductAmount"`

	// 资金转入总额，单位（分）
	AgentInAmount *float64 `json:"AgentInAmount,omitnil,omitempty" name:"AgentInAmount"`

	// 垫付充值总额，单位（分）
	AdvanceRechargeAmount *float64 `json:"AdvanceRechargeAmount,omitnil,omitempty" name:"AdvanceRechargeAmount"`

	// 提现扣减总额，单位（分）
	WithdrawAmount *float64 `json:"WithdrawAmount,omitnil,omitempty" name:"WithdrawAmount"`

	// 资金转出总额，单位（分）
	AgentOutAmount *float64 `json:"AgentOutAmount,omitnil,omitempty" name:"AgentOutAmount"`

	// 还垫付总额，单位（分）
	AdvancePayAmount *float64 `json:"AdvancePayAmount,omitnil,omitempty" name:"AdvancePayAmount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 月份，格式为yyyy-mm。不能早于开通账单2.0的月份
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	//
	// Deprecated: PeriodType is deprecated.
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

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
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 付费模式 prePay/postPay
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 分账标签键，用户自定义（支持2021-01以后账单查询）
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 分账标签值，该参数为空表示该标签键下未设置标签值的记录
	// （支持2021-01以后账单查询）
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type DescribeBillResourceSummaryForOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，依次类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 月份，格式为yyyy-mm。不能早于开通账单2.0的月份
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

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
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 付费模式 prePay/postPay
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 分账标签键，用户自定义（支持2021-01以后账单查询）
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 分账标签值，该参数为空表示该标签键下未设置标签值的记录
	// （支持2021-01以后账单查询）
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
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
	ResourceSummarySet []*BillDistributionResourceSummary `json:"ResourceSummarySet,omitnil,omitempty" name:"ResourceSummarySet"`

	// 资源汇总列表总数，入参NeedRecordNum为0时不返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 月份，格式为yyyy-mm。不能早于开通账单2.0的月份
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	//
	// Deprecated: PeriodType is deprecated.
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

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
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 付费模式 prePay/postPay
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 支付者的账号 ID（账号 ID 是用户在腾讯云的唯一账号标识），默认查询本账号账单，如集团管理账号需查询成员账号自付的账单，该字段需入参成员账号UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 分账标签键，用户自定义（支持2021-01以后账单查询）
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 分账标签值，该参数为空表示该标签键下未设置标签值的记录
	// （支持2021-01以后账单查询）
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type DescribeBillResourceSummaryRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，依次类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 月份，格式为yyyy-mm。不能早于开通账单2.0的月份
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

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
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 付费模式 prePay/postPay
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 支付者的账号 ID（账号 ID 是用户在腾讯云的唯一账号标识），默认查询本账号账单，如集团管理账号需查询成员账号自付的账单，该字段需入参成员账号UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 分账标签键，用户自定义（支持2021-01以后账单查询）
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 分账标签值，该参数为空表示该标签键下未设置标签值的记录
	// （支持2021-01以后账单查询）
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
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
	ResourceSummarySet []*BillResourceSummary `json:"ResourceSummarySet,omitnil,omitempty" name:"ResourceSummarySet"`

	// 资源汇总列表总数，入参NeedRecordNum为0时不返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
}

type DescribeBillSummaryByPayModeRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
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
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// 各付费模式花费分布详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryOverview []*PayModeSummaryOverviewItem `json:"SummaryOverview,omitnil,omitempty" name:"SummaryOverview"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 款项类别，与L0账单上的汇总类别对应。
	// 此参数自账单3.0（即2021-05）之后开始生效。
	// 枚举值：
	// consume-消费
	// refund-退款
	// adjustment-调账
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`
}

type DescribeBillSummaryByProductRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 款项类别，与L0账单上的汇总类别对应。
	// 此参数自账单3.0（即2021-05）之后开始生效。
	// 枚举值：
	// consume-消费
	// refund-退款
	// adjustment-调账
	PayType *string `json:"PayType,omitnil,omitempty" name:"PayType"`
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
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// 总花费详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryTotal *BusinessSummaryTotal `json:"SummaryTotal,omitnil,omitempty" name:"SummaryTotal"`

	// 各产品花费分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryOverview []*BusinessSummaryOverviewItem `json:"SummaryOverview,omitnil,omitempty" name:"SummaryOverview"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
}

type DescribeBillSummaryByProjectRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
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
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// 各项目花费分布详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryOverview []*ProjectSummaryOverviewItem `json:"SummaryOverview,omitnil,omitempty" name:"SummaryOverview"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
}

type DescribeBillSummaryByRegionRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
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
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// 各地域花费分布详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryOverview []*RegionSummaryOverviewItem `json:"SummaryOverview,omitnil,omitempty" name:"SummaryOverview"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分账标签键，用户自定义
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 分账标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type DescribeBillSummaryByTagRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分账标签键，用户自定义
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 分账标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
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
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// 各标签值花费分布详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryOverview []*TagSummaryOverviewItem `json:"SummaryOverview,omitnil,omitempty" name:"SummaryOverview"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryTotal *SummaryTotal `json:"SummaryTotal,omitnil,omitempty" name:"SummaryTotal"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 账单维度类型，枚举值如下：business、project、region、payMode、tag
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 标签键，GroupType=tag获取标签维度账单时传
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
}

type DescribeBillSummaryForOrganizationRequest struct {
	*tchttp.BaseRequest
	
	// 账单月份，格式为2023-04
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 账单维度类型，枚举值如下：business、project、region、payMode、tag
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 标签键，GroupType=tag获取标签维度账单时传
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
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
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// 账单多维度汇总消费详情
	SummaryDetail []*SummaryDetail `json:"SummaryDetail,omitnil,omitempty" name:"SummaryDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 账单维度类型，枚举值如下：business、project、region、payMode、tag
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 标签键，GroupType=tag获取标签维度账单时传
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
}

type DescribeBillSummaryRequest struct {
	*tchttp.BaseRequest
	
	// 账单月份，格式为2023-04
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 账单维度类型，枚举值如下：business、project、region、payMode、tag
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 标签键，GroupType=tag获取标签维度账单时传
	TagKey []*string `json:"TagKey,omitnil,omitempty" name:"TagKey"`
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
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// 账单多维度汇总消费详情
	SummaryDetail []*SummaryDetail `json:"SummaryDetail,omitnil,omitempty" name:"SummaryDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 周期开始时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为同一月份，暂不支持跨月拉取。可拉取的数据是开通成本分析后，且距今 24 个月内的数据。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 周期结束时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为同一月份，暂不支持跨月拉取。可拉取的数据是开通成本分析后，且距今 24 个月内的数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// 月份，格式为yyyy-mm，Month和BeginTime&EndTime必传一个，如果有传BeginTime&EndTime则Month字段无效。不能早于开通成本分析的月份，最多可拉取24个月内的数据。
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 查询指定产品信息（暂时未开放获取）
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 付费模式 prePay/postPay
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type DescribeCostDetailRequest struct {
	*tchttp.BaseRequest
	
	// 数量，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 周期开始时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为同一月份，暂不支持跨月拉取。可拉取的数据是开通成本分析后，且距今 24 个月内的数据。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 周期结束时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为同一月份，暂不支持跨月拉取。可拉取的数据是开通成本分析后，且距今 24 个月内的数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// 月份，格式为yyyy-mm，Month和BeginTime&EndTime必传一个，如果有传BeginTime&EndTime则Month字段无效。不能早于开通成本分析的月份，最多可拉取24个月内的数据。
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 查询指定产品信息（暂时未开放获取）
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 付费模式 prePay/postPay
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
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
	DetailSet []*CostDetail `json:"DetailSet,omitnil,omitempty" name:"DetailSet"`

	// 记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeCostExplorerSummaryRequestParams struct {
	// 周期开始时间，格式为yyyy-mm-dd hh:ii:ss
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 周期结束时间，格式为yyyy-mm-dd hh:ii:ss
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 账单类型：1-费用账单、2-消耗账单
	BillType *string `json:"BillType,omitnil,omitempty" name:"BillType"`

	// 统计周期：日-day，月-month；
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 分类维度（数据汇总维度），查询分类维度（请使用分类维度code入参）入参枚举值：
	// default=仅总计
	// feeType=费用类型
	// billType=账单类型
	// business=产品
	// product=子产品
	// region=地域
	// zone=可用区
	// actionType=交易类型
	// payMode =计费模式
	// tags=标签
	// project =项目
	// payerUin=支付者账号
	// ownerUin=使用者账号
	Dimensions *string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// 费用类型：cost-总费用，totalCost-原价费用
	FeeType *string `json:"FeeType,omitnil,omitempty" name:"FeeType"`

	// 数量，每页最大值为100
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 起始页，当PageNo=1表示第一页， PageNo=2表示第二页，依次类推。
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 分账标签值
	TagKeyStr *string `json:"TagKeyStr,omitnil,omitempty" name:"TagKeyStr"`

	// 是否需要筛选框， 1-表示需要， 0-表示不需要，若不传默认不需要。
	NeedConditionValue *string `json:"NeedConditionValue,omitnil,omitempty" name:"NeedConditionValue"`

	// 筛选参数
	Conditions *AnalyseConditions `json:"Conditions,omitnil,omitempty" name:"Conditions"`
}

type DescribeCostExplorerSummaryRequest struct {
	*tchttp.BaseRequest
	
	// 周期开始时间，格式为yyyy-mm-dd hh:ii:ss
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 周期结束时间，格式为yyyy-mm-dd hh:ii:ss
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 账单类型：1-费用账单、2-消耗账单
	BillType *string `json:"BillType,omitnil,omitempty" name:"BillType"`

	// 统计周期：日-day，月-month；
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`

	// 分类维度（数据汇总维度），查询分类维度（请使用分类维度code入参）入参枚举值：
	// default=仅总计
	// feeType=费用类型
	// billType=账单类型
	// business=产品
	// product=子产品
	// region=地域
	// zone=可用区
	// actionType=交易类型
	// payMode =计费模式
	// tags=标签
	// project =项目
	// payerUin=支付者账号
	// ownerUin=使用者账号
	Dimensions *string `json:"Dimensions,omitnil,omitempty" name:"Dimensions"`

	// 费用类型：cost-总费用，totalCost-原价费用
	FeeType *string `json:"FeeType,omitnil,omitempty" name:"FeeType"`

	// 数量，每页最大值为100
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 起始页，当PageNo=1表示第一页， PageNo=2表示第二页，依次类推。
	PageNo *uint64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 分账标签值
	TagKeyStr *string `json:"TagKeyStr,omitnil,omitempty" name:"TagKeyStr"`

	// 是否需要筛选框， 1-表示需要， 0-表示不需要，若不传默认不需要。
	NeedConditionValue *string `json:"NeedConditionValue,omitnil,omitempty" name:"NeedConditionValue"`

	// 筛选参数
	Conditions *AnalyseConditions `json:"Conditions,omitnil,omitempty" name:"Conditions"`
}

func (r *DescribeCostExplorerSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostExplorerSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "BillType")
	delete(f, "PeriodType")
	delete(f, "Dimensions")
	delete(f, "FeeType")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "TagKeyStr")
	delete(f, "NeedConditionValue")
	delete(f, "Conditions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCostExplorerSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostExplorerSummaryResponseParams struct {
	// 数据条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 表头信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Header *AnalyseHeaderDetail `json:"Header,omitnil,omitempty" name:"Header"`

	// 数据明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail []*AnalyseDetail `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 数据总计
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalDetail *AnalyseDetail `json:"TotalDetail,omitnil,omitempty" name:"TotalDetail"`

	// 筛选框
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionValue *AnalyseConditionDetail `json:"ConditionValue,omitnil,omitempty" name:"ConditionValue"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCostExplorerSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCostExplorerSummaryResponseParams `json:"Response"`
}

func (r *DescribeCostExplorerSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCostExplorerSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCostSummaryByProductRequestParams struct {
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 每次获取数据量，最大值为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量,默认从0开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`
}

type DescribeCostSummaryByProductRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 每次获取数据量，最大值为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量,默认从0开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`
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
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// 消耗详情
	Total *ConsumptionSummaryTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// 消耗按产品汇总详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ConsumptionBusinessSummaryDataItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 记录数量，NeedRecordNum为0时返回null
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 每次获取数据量，最大值为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量,默认从0开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`
}

type DescribeCostSummaryByProjectRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 每次获取数据量，最大值为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量,默认从0开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`
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
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// 消耗详情
	Total *ConsumptionSummaryTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// 消耗按业务汇总详情
	Data []*ConsumptionProjectSummaryDataItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 记录数量，NeedRecordNum为0时返回null
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 每次获取数据量，最大值为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量,默认从0开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`
}

type DescribeCostSummaryByRegionRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 每次获取数据量，最大值为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量,默认从0开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`
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
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// 消耗详情
	Total *ConsumptionSummaryTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// 消耗按地域汇总详情
	Data []*ConsumptionRegionSummaryDataItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 记录数量，NeedRecordNum为0时返回null
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 每次获取数据量，最大值为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量,默认从0开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// 是否需要返回过滤条件，0不需要，1需要，默认不需要
	NeedConditionValue *uint64 `json:"NeedConditionValue,omitnil,omitempty" name:"NeedConditionValue"`

	// 过滤条件，只支持ResourceKeyword(资源关键字，支持资源id及资源名称模糊查询)，ProjectIds（项目id），RegionIds(地域id)，PayModes(付费模式，可选prePay和postPay)，HideFreeCost（是否隐藏0元流水，可选0和1），OrderByCost（按费用排序规则，可选desc和asc）
	Conditions *Conditions `json:"Conditions,omitnil,omitempty" name:"Conditions"`
}

type DescribeCostSummaryByResourceRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 每次获取数据量，最大值为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量,默认从0开始
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitnil,omitempty" name:"NeedRecordNum"`

	// 是否需要返回过滤条件，0不需要，1需要，默认不需要
	NeedConditionValue *uint64 `json:"NeedConditionValue,omitnil,omitempty" name:"NeedConditionValue"`

	// 过滤条件，只支持ResourceKeyword(资源关键字，支持资源id及资源名称模糊查询)，ProjectIds（项目id），RegionIds(地域id)，PayModes(付费模式，可选prePay和postPay)，HideFreeCost（是否隐藏0元流水，可选0和1），OrderByCost（按费用排序规则，可选desc和asc）
	Conditions *Conditions `json:"Conditions,omitnil,omitempty" name:"Conditions"`
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
	Ready *uint64 `json:"Ready,omitnil,omitempty" name:"Ready"`

	// 消耗详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *ConsumptionSummaryTotal `json:"Total,omitnil,omitempty" name:"Total"`

	// 过滤条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionValue *ConsumptionResourceSummaryConditionValue `json:"ConditionValue,omitnil,omitempty" name:"ConditionValue"`

	// 记录数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// 资源消耗详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ConsumptionResourceSummaryDataItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间 2016-02-01 00:00:00 建议跨度不超过3个月
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 一页多少条数据，默认是20条，最大不超过1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 第多少页，从0开始，默认是0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 子订单号
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 大订单号
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type DescribeDealsByCondRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间 2016-01-01 00:00:00
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间 2016-02-01 00:00:00 建议跨度不超过3个月
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 一页多少条数据，默认是20条，最大不超过1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 第多少页，从0开始，默认是0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

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
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 子订单号
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 大订单号
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
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
	Deals []*Deal `json:"Deals,omitnil,omitempty" name:"Deals"`

	// 订单总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 查询用量开始时间，格式为yyyy-mm-dd，例如：2020-09-01
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 查询用量结束时间，格式为yyyy-mm-dd，例如：2020-09-30（与开始时间同月，不支持跨月查询）
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// COS 存储桶名称，可通过Get Service 接口是用来获取请求者名下的所有存储空间列表（Bucket list）https://cloud.tencent.com/document/product/436/8291
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`
}

type DescribeDosageCosDetailByDateRequest struct {
	*tchttp.BaseRequest
	
	// 查询用量开始时间，格式为yyyy-mm-dd，例如：2020-09-01
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 查询用量结束时间，格式为yyyy-mm-dd，例如：2020-09-30（与开始时间同月，不支持跨月查询）
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// COS 存储桶名称，可通过Get Service 接口是用来获取请求者名下的所有存储空间列表（Bucket list）https://cloud.tencent.com/document/product/436/8291
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`
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
	DetailSets []*CosDetailSets `json:"DetailSets,omitnil,omitempty" name:"DetailSets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type DescribeDosageDetail struct {
	// 日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 账号 ID 是用户在腾讯云的唯一账号标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 用量统计类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DosageType *string `json:"DosageType,omitnil,omitempty" name:"DosageType"`

	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 子产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// 组件类型编码
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingItemCode *string `json:"BillingItemCode,omitnil,omitempty" name:"BillingItemCode"`

	// 组件编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubBillingItemCode *string `json:"SubBillingItemCode,omitnil,omitempty" name:"SubBillingItemCode"`

	// 产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// 子产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductCodeName *string `json:"SubProductCodeName,omitnil,omitempty" name:"SubProductCodeName"`

	// 组件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingItemCodeName *string `json:"BillingItemCodeName,omitnil,omitempty" name:"BillingItemCodeName"`

	// 组件
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubBillingItemCodeName *string `json:"SubBillingItemCodeName,omitnil,omitempty" name:"SubBillingItemCodeName"`

	// 用量单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	DosageUnit *string `json:"DosageUnit,omitnil,omitempty" name:"DosageUnit"`

	// 用量起始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DosageBeginTime *string `json:"DosageBeginTime,omitnil,omitempty" name:"DosageBeginTime"`

	// 用量截止时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DosageEndTime *string `json:"DosageEndTime,omitnil,omitempty" name:"DosageEndTime"`

	// 标准用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DosageValue *float64 `json:"DosageValue,omitnil,omitempty" name:"DosageValue"`

	// 抵扣用量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeductValue *float64 `json:"DeductValue,omitnil,omitempty" name:"DeductValue"`

	// 抵扣余量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemainValue *float64 `json:"RemainValue,omitnil,omitempty" name:"RemainValue"`

	// sdkAppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// 其他信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttrStr []*JsonObject `json:"AttrStr,omitnil,omitempty" name:"AttrStr"`

	// 用量模板名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SheetName []*string `json:"SheetName,omitnil,omitempty" name:"SheetName"`
}

// Predefined struct for user
type DescribeDosageDetailByDateRequestParams struct {
	// 查询账单开始日期，如 2019-01-01
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 查询账单结束日期，如 2019-01-01， 时间跨度不超过7天
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

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
	// 
	// 仅支持以上产品
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 查询域名 例如 www.qq.com
	// 非CDN业务查询时传入空字符串，返回的值为空
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 1、如果为空，则返回EIP或CLB所有实例的明细；
	// 2、如果传入实例名，则返回该实例明细
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 支付者的账号 ID（账号 ID 是用户在腾讯云的唯一账号标识），默认查询本账号账单，如集团管理账号需查询成员账号自付的账单，该字段需入参成员账号UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
}

type DescribeDosageDetailByDateRequest struct {
	*tchttp.BaseRequest
	
	// 查询账单开始日期，如 2019-01-01
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 查询账单结束日期，如 2019-01-01， 时间跨度不超过7天
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

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
	// 
	// 仅支持以上产品
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 查询域名 例如 www.qq.com
	// 非CDN业务查询时传入空字符串，返回的值为空
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 1、如果为空，则返回EIP或CLB所有实例的明细；
	// 2、如果传入实例名，则返回该实例明细
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`

	// 支付者的账号 ID（账号 ID 是用户在腾讯云的唯一账号标识），默认查询本账号账单，如集团管理账号需查询成员账号自付的账单，该字段需入参成员账号UIN
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`
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
	delete(f, "PayerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDosageDetailByDateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDosageDetailByDateResponseParams struct {
	// 计量单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 用量数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetailSets []*DetailSet `json:"DetailSets,omitnil,omitempty" name:"DetailSets"`

	// 错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetCode *int64 `json:"RetCode,omitnil,omitempty" name:"RetCode"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetMsg *string `json:"RetMsg,omitnil,omitempty" name:"RetMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeDosageDetailListRequestParams struct {
	// 用量起始时间，如：2023-02-01
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 用量截止时间，如：2023-02-28
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 产品编码，已支持查询的产品如下：
	// p_ccc（云联络中心）
	// p_rav（实时音视频）
	// p_pstn（号码保护）
	// p_smh（智能媒资托管）
	// p_coding_devops（CODING DevOps）
	// p_dsa（全球IP应用加速）
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 数据偏移量（从0开始）
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 单次数据量（最大3000）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 用量统计类型：用量明细的数据统计汇总周期类型，包括minute-按5分钟汇总、hour-按小时汇总、day-按天汇总、month-按月汇总、comm-其他，默认查询所有类型明细，目前各产品已支持的统计类型如下：
	// p_ccc（云联络中心）：comm、day
	// p_rav（实时音视频）：minute、day
	// p_pstn（号码保护）：comm
	// p_smh（智能媒资托管）：day
	// p_coding_devops（CODING DevOps）：comm、day
	// p_dsa（全球IP应用加速）：minute
	DosageType *string `json:"DosageType,omitnil,omitempty" name:"DosageType"`
}

type DescribeDosageDetailListRequest struct {
	*tchttp.BaseRequest
	
	// 用量起始时间，如：2023-02-01
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 用量截止时间，如：2023-02-28
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 产品编码，已支持查询的产品如下：
	// p_ccc（云联络中心）
	// p_rav（实时音视频）
	// p_pstn（号码保护）
	// p_smh（智能媒资托管）
	// p_coding_devops（CODING DevOps）
	// p_dsa（全球IP应用加速）
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 数据偏移量（从0开始）
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 单次数据量（最大3000）
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 用量统计类型：用量明细的数据统计汇总周期类型，包括minute-按5分钟汇总、hour-按小时汇总、day-按天汇总、month-按月汇总、comm-其他，默认查询所有类型明细，目前各产品已支持的统计类型如下：
	// p_ccc（云联络中心）：comm、day
	// p_rav（实时音视频）：minute、day
	// p_pstn（号码保护）：comm
	// p_smh（智能媒资托管）：day
	// p_coding_devops（CODING DevOps）：comm、day
	// p_dsa（全球IP应用加速）：minute
	DosageType *string `json:"DosageType,omitnil,omitempty" name:"DosageType"`
}

func (r *DescribeDosageDetailListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDosageDetailListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ProductCode")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DosageType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDosageDetailListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDosageDetailListResponseParams struct {
	// 用量明细集合
	Record []*DescribeDosageDetail `json:"Record,omitnil,omitempty" name:"Record"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDosageDetailListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDosageDetailListResponseParams `json:"Response"`
}

func (r *DescribeDosageDetailListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDosageDetailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatherResourceRequestParams struct {
	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，以此类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 分账单元唯一标识，用作筛选
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// 资源目录类别，枚举值如下：
	// all - 全部 
	// none - 未归集
	GatherType *string `json:"GatherType,omitnil,omitempty" name:"GatherType"`

	// 排序字段，枚举值如下：
	// realCost  - 折后总价
	// cashPayAmount - 现金金额
	// voucherPayAmount - 代金券金额
	// incentivePayAmount  - 赠送金金额
	// transferPayAmount -分成金金额
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序类型，枚举值如下：
	// asc - 升序
	// desc - 降序
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 产品编码，用作筛选
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// 子产品编码，用作筛选
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// 组件名称编码，用作筛选
	ItemCodes []*string `json:"ItemCodes,omitnil,omitempty" name:"ItemCodes"`

	// 地域ID，用作筛选
	RegionIds []*uint64 `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// 实例类型编码，用作筛选
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 计费模式编码，用作筛选
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// 操作者UIN，用作筛选
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// 使用者UIN，用作筛选
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// 模糊搜索：支持标签、资源id、资源别名
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 标签，用作筛选
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 项目ID，用作筛选
	ProjectIds []*string `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 交易类型编码，用作筛选
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`
}

type DescribeGatherResourceRequest struct {
	*tchttp.BaseRequest
	
	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，以此类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 账单月份，格式为2024-02，不传默认当前月
	Month *string `json:"Month,omitnil,omitempty" name:"Month"`

	// 分账单元唯一标识，用作筛选
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// 资源目录类别，枚举值如下：
	// all - 全部 
	// none - 未归集
	GatherType *string `json:"GatherType,omitnil,omitempty" name:"GatherType"`

	// 排序字段，枚举值如下：
	// realCost  - 折后总价
	// cashPayAmount - 现金金额
	// voucherPayAmount - 代金券金额
	// incentivePayAmount  - 赠送金金额
	// transferPayAmount -分成金金额
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序类型，枚举值如下：
	// asc - 升序
	// desc - 降序
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 产品编码，用作筛选
	BusinessCodes []*string `json:"BusinessCodes,omitnil,omitempty" name:"BusinessCodes"`

	// 子产品编码，用作筛选
	ProductCodes []*string `json:"ProductCodes,omitnil,omitempty" name:"ProductCodes"`

	// 组件名称编码，用作筛选
	ItemCodes []*string `json:"ItemCodes,omitnil,omitempty" name:"ItemCodes"`

	// 地域ID，用作筛选
	RegionIds []*uint64 `json:"RegionIds,omitnil,omitempty" name:"RegionIds"`

	// 实例类型编码，用作筛选
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 计费模式编码，用作筛选
	PayModes []*string `json:"PayModes,omitnil,omitempty" name:"PayModes"`

	// 操作者UIN，用作筛选
	OperateUins []*string `json:"OperateUins,omitnil,omitempty" name:"OperateUins"`

	// 使用者UIN，用作筛选
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// 模糊搜索：支持标签、资源id、资源别名
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 标签，用作筛选
	Tag []*string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 项目ID，用作筛选
	ProjectIds []*string `json:"ProjectIds,omitnil,omitempty" name:"ProjectIds"`

	// 交易类型编码，用作筛选
	ActionTypes []*string `json:"ActionTypes,omitnil,omitempty" name:"ActionTypes"`
}

func (r *DescribeGatherResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatherResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Month")
	delete(f, "TreeNodeUniqKey")
	delete(f, "GatherType")
	delete(f, "Sort")
	delete(f, "SortType")
	delete(f, "BusinessCodes")
	delete(f, "ProductCodes")
	delete(f, "ItemCodes")
	delete(f, "RegionIds")
	delete(f, "InstanceTypes")
	delete(f, "PayModes")
	delete(f, "OperateUins")
	delete(f, "OwnerUins")
	delete(f, "SearchKey")
	delete(f, "Tag")
	delete(f, "ProjectIds")
	delete(f, "ActionTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatherResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatherResourceResponseParams struct {
	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordNum *int64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// 资源归集汇总
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatherResourceSummary []*GatherResourceSummary `json:"GatherResourceSummary,omitnil,omitempty" name:"GatherResourceSummary"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdateTime *string `json:"LastUpdateTime,omitnil,omitempty" name:"LastUpdateTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGatherResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatherResourceResponseParams `json:"Response"`
}

func (r *DescribeGatherResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatherResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSavingPlanCoverageRequestParams struct {
	// 费用起始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 费用结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，以此类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量，最大值为200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 取值包括1（缺省值）和2，1表示按天统计覆盖率，2表示按月统计覆盖率，此参数仅影响返回的RateSet聚合粒度，不影响返回的DetailSet
	PeriodType *uint64 `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`
}

type DescribeSavingPlanCoverageRequest struct {
	*tchttp.BaseRequest
	
	// 费用起始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 费用结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，以此类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量，最大值为200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 取值包括1（缺省值）和2，1表示按天统计覆盖率，2表示按月统计覆盖率，此参数仅影响返回的RateSet聚合粒度，不影响返回的DetailSet
	PeriodType *uint64 `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`
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
	DetailSet []*SavingPlanCoverageDetail `json:"DetailSet,omitnil,omitempty" name:"DetailSet"`

	// 节省计划覆盖率聚合数据
	RateSet []*SavingPlanCoverageRate `json:"RateSet,omitnil,omitempty" name:"RateSet"`

	// 查询命中的节省计划覆盖率明细数据总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 结束时间，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，最大值为200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSavingPlanOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，格式yyyy-MM-dd 注：查询范围请勿超过6个月
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 结束时间，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，最大值为200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	Overviews []*SavingPlanOverviewDetail `json:"Overviews,omitnil,omitempty" name:"Overviews"`

	// 查询命中的节省计划总览明细数据总条数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeSavingPlanResourceInfoRequestParams struct {
	// 数量，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 购买开始时间，格式yyyy-MM-dd
	CreateStartDate *string `json:"CreateStartDate,omitnil,omitempty" name:"CreateStartDate"`

	// 购买结束时间，格式yyyy-MM-dd
	CreateEndDate *string `json:"CreateEndDate,omitnil,omitempty" name:"CreateEndDate"`
}

type DescribeSavingPlanResourceInfoRequest struct {
	*tchttp.BaseRequest
	
	// 数量，最大值为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 购买开始时间，格式yyyy-MM-dd
	CreateStartDate *string `json:"CreateStartDate,omitnil,omitempty" name:"CreateStartDate"`

	// 购买结束时间，格式yyyy-MM-dd
	CreateEndDate *string `json:"CreateEndDate,omitnil,omitempty" name:"CreateEndDate"`
}

func (r *DescribeSavingPlanResourceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSavingPlanResourceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "CreateStartDate")
	delete(f, "CreateEndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSavingPlanResourceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSavingPlanResourceInfoResponseParams struct {
	// 记录数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSavingPlanResourceInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSavingPlanResourceInfoResponseParams `json:"Response"`
}

func (r *DescribeSavingPlanResourceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSavingPlanResourceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSavingPlanUsageRequestParams struct {
	// 开始时间，格式yyyy-MM-dd 注：查询范围请勿超过6个月
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 结束时间，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，最大值为200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询结果数据的时间间隔
	TimeInterval *string `json:"TimeInterval,omitnil,omitempty" name:"TimeInterval"`
}

type DescribeSavingPlanUsageRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，格式yyyy-MM-dd 注：查询范围请勿超过6个月
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 结束时间，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，最大值为200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询结果数据的时间间隔
	TimeInterval *string `json:"TimeInterval,omitnil,omitempty" name:"TimeInterval"`
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
	Usages []*SavingPlanUsageDetail `json:"Usages,omitnil,omitempty" name:"Usages"`

	// 查询命中的节省计划总览明细数据总条数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，依次类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分账标签键，用作模糊搜索
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签类型，枚举值：0普通标签，1分账标签，用作筛选，不传获取全部标签键
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 排序方式，枚举值：asc排升序，desc排降序
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

type DescribeTagListRequest struct {
	*tchttp.BaseRequest
	
	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量，Offset=0表示第一页，如果Limit=100，则Offset=100表示第二页，Offset=200表示第三页，依次类推
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分账标签键，用作模糊搜索
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签类型，枚举值：0普通标签，1分账标签，用作筛选，不传获取全部标签键
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 排序方式，枚举值：asc排升序，desc排降序
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
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
	RecordNum *uint64 `json:"RecordNum,omitnil,omitempty" name:"RecordNum"`

	// 标签信息
	Data []*TagDataInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 第多少页，默认是1
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 券状态：待使用：unUsed，已使用： used，已发货：delivered，已作废： cancel，已过期：overdue
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 代金券id
	VoucherId *string `json:"VoucherId,omitnil,omitempty" name:"VoucherId"`

	// 代金券订单id
	CodeId *string `json:"CodeId,omitnil,omitempty" name:"CodeId"`

	// 商品码
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 活动id
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 代金券名称
	VoucherName *string `json:"VoucherName,omitnil,omitempty" name:"VoucherName"`

	// 发放开始时间,例：2021-01-01
	TimeFrom *string `json:"TimeFrom,omitnil,omitempty" name:"TimeFrom"`

	// 发放结束时间，例：2021-01-01
	TimeTo *string `json:"TimeTo,omitnil,omitempty" name:"TimeTo"`

	// 指定排序字段：BeginTime开始时间、EndTime到期时间、CreateTime创建时间
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// 指定升序降序：desc、asc
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// 付费模式，postPay后付费/prePay预付费/riPay预留实例/""或者"*"表示全部模式，如果payMode为""或"*"，那么productCode与subProductCode必须传空
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 付费场景PayMode=postPay时：spotpay-竞价实例,"settle account"-普通后付费PayMode=prePay时：purchase-包年包月新购，renew-包年包月续费（自动续费），modify-包年包月配置变更(变配）PayMode=riPay时：oneOffFee-预留实例预付，hourlyFee-预留实例每小时扣费，*-支持全部付费场景
	PayScene *string `json:"PayScene,omitnil,omitempty" name:"PayScene"`

	// 操作人，默认就是用户uin
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代金券主类型 has_price 为有价现金券 no_price 为无价代金券
	VoucherMainType *string `json:"VoucherMainType,omitnil,omitempty" name:"VoucherMainType"`

	// 代金券副类型 discount 为折扣券 deduct 为抵扣券
	VoucherSubType *string `json:"VoucherSubType,omitnil,omitempty" name:"VoucherSubType"`

	// 券有效时间开始时间
	StartTimeFrom *string `json:"StartTimeFrom,omitnil,omitempty" name:"StartTimeFrom"`

	// 券有效时间结束时间
	StartTimeTo *string `json:"StartTimeTo,omitnil,omitempty" name:"StartTimeTo"`

	// 券失效时间开始时间
	EndTimeFrom *string `json:"EndTimeFrom,omitnil,omitempty" name:"EndTimeFrom"`

	// 券失效时间结束时间
	EndTimeTo *string `json:"EndTimeTo,omitnil,omitempty" name:"EndTimeTo"`

	// 发券时间开始时间
	CreateTimeFrom *string `json:"CreateTimeFrom,omitnil,omitempty" name:"CreateTimeFrom"`

	// 发券时间结束时间
	CreateTimeTo *string `json:"CreateTimeTo,omitnil,omitempty" name:"CreateTimeTo"`
}

type DescribeVoucherInfoRequest struct {
	*tchttp.BaseRequest
	
	// 一页多少条数据，默认是20条，最大不超过1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 第多少页，默认是1
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 券状态：待使用：unUsed，已使用： used，已发货：delivered，已作废： cancel，已过期：overdue
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 代金券id
	VoucherId *string `json:"VoucherId,omitnil,omitempty" name:"VoucherId"`

	// 代金券订单id
	CodeId *string `json:"CodeId,omitnil,omitempty" name:"CodeId"`

	// 商品码
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 活动id
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 代金券名称
	VoucherName *string `json:"VoucherName,omitnil,omitempty" name:"VoucherName"`

	// 发放开始时间,例：2021-01-01
	TimeFrom *string `json:"TimeFrom,omitnil,omitempty" name:"TimeFrom"`

	// 发放结束时间，例：2021-01-01
	TimeTo *string `json:"TimeTo,omitnil,omitempty" name:"TimeTo"`

	// 指定排序字段：BeginTime开始时间、EndTime到期时间、CreateTime创建时间
	SortField *string `json:"SortField,omitnil,omitempty" name:"SortField"`

	// 指定升序降序：desc、asc
	SortOrder *string `json:"SortOrder,omitnil,omitempty" name:"SortOrder"`

	// 付费模式，postPay后付费/prePay预付费/riPay预留实例/""或者"*"表示全部模式，如果payMode为""或"*"，那么productCode与subProductCode必须传空
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 付费场景PayMode=postPay时：spotpay-竞价实例,"settle account"-普通后付费PayMode=prePay时：purchase-包年包月新购，renew-包年包月续费（自动续费），modify-包年包月配置变更(变配）PayMode=riPay时：oneOffFee-预留实例预付，hourlyFee-预留实例每小时扣费，*-支持全部付费场景
	PayScene *string `json:"PayScene,omitnil,omitempty" name:"PayScene"`

	// 操作人，默认就是用户uin
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 代金券主类型 has_price 为有价现金券 no_price 为无价代金券
	VoucherMainType *string `json:"VoucherMainType,omitnil,omitempty" name:"VoucherMainType"`

	// 代金券副类型 discount 为折扣券 deduct 为抵扣券
	VoucherSubType *string `json:"VoucherSubType,omitnil,omitempty" name:"VoucherSubType"`

	// 券有效时间开始时间
	StartTimeFrom *string `json:"StartTimeFrom,omitnil,omitempty" name:"StartTimeFrom"`

	// 券有效时间结束时间
	StartTimeTo *string `json:"StartTimeTo,omitnil,omitempty" name:"StartTimeTo"`

	// 券失效时间开始时间
	EndTimeFrom *string `json:"EndTimeFrom,omitnil,omitempty" name:"EndTimeFrom"`

	// 券失效时间结束时间
	EndTimeTo *string `json:"EndTimeTo,omitnil,omitempty" name:"EndTimeTo"`

	// 发券时间开始时间
	CreateTimeFrom *string `json:"CreateTimeFrom,omitnil,omitempty" name:"CreateTimeFrom"`

	// 发券时间结束时间
	CreateTimeTo *string `json:"CreateTimeTo,omitnil,omitempty" name:"CreateTimeTo"`
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
	delete(f, "StartTimeFrom")
	delete(f, "StartTimeTo")
	delete(f, "EndTimeFrom")
	delete(f, "EndTimeTo")
	delete(f, "CreateTimeFrom")
	delete(f, "CreateTimeTo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVoucherInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVoucherInfoResponseParams struct {
	// 券总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总余额（微分）
	TotalBalance *int64 `json:"TotalBalance,omitnil,omitempty" name:"TotalBalance"`

	// 代金券相关信息
	VoucherInfos []*VoucherInfos `json:"VoucherInfos,omitnil,omitempty" name:"VoucherInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 第多少页，默认是1
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 代金券id
	VoucherId *string `json:"VoucherId,omitnil,omitempty" name:"VoucherId"`

	// 操作人，默认就是用户uin
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type DescribeVoucherUsageDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 一页多少条数据，默认是20条，最大不超过1000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 第多少页，默认是1
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 代金券id
	VoucherId *string `json:"VoucherId,omitnil,omitempty" name:"VoucherId"`

	// 操作人，默认就是用户uin
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总已用金额（微分）
	TotalUsedAmount *int64 `json:"TotalUsedAmount,omitnil,omitempty" name:"TotalUsedAmount"`

	// 代金券使用记录细节
	UsageRecords []*UsageRecords `json:"UsageRecords,omitnil,omitempty" name:"UsageRecords"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type DetailSet struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 使用数据明细
	DetailPoints []*DetailPoint `json:"DetailPoints,omitnil,omitempty" name:"DetailPoints"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceID *string `json:"InstanceID,omitnil,omitempty" name:"InstanceID"`
}

type DistributionBillDetail struct {
	// 产品名称：用户所采购的各类云产品，例如：云服务器 CVM
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// 子产品名称：用户采购的具体产品细分类型，例如：云服务器 CVM-标准型 S1
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// 计费模式：资源的计费模式，区分为包年包月和按量计费
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// 项目名称：资源归属的项目，用户在控制台给资源自主分配项目，未分配则是默认项目
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 地域：资源所属地域，如华南地区（广州）
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 可用区：资源所属可用区，如广州三区
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 资源 ID：账单中出账对象 ID，不同产品因资源形态不同，资源内容不完全相同，如云服务器 CVM 为对应的实例 ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源别名：用户在控制台为资源设置的名称，如果未设置，则默认为空
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 交易类型，如包年包月新购、包年包月续费、按量计费扣费等类型
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// 订单ID：包年包月计费模式下订购的订单号
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 交易ID：结算扣费单号
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// 扣费时间：结算扣费时间
	PayTime *string `json:"PayTime,omitnil,omitempty" name:"PayTime"`

	// 开始使用时间：产品服务开始使用时间
	FeeBeginTime *string `json:"FeeBeginTime,omitnil,omitempty" name:"FeeBeginTime"`

	// 结束使用时间：产品服务结束使用时间
	FeeEndTime *string `json:"FeeEndTime,omitnil,omitempty" name:"FeeEndTime"`

	// 组件列表
	ComponentSet []*BillDetailComponent `json:"ComponentSet,omitnil,omitempty" name:"ComponentSet"`

	// 使用者UIN：实际使用资源的账号 ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 操作者UIN：操作者账号 ID（预付费资源下单或后付费操作开通资源账号的 ID 或者角色 ID ）
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*BillTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 子产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 交易类型编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 价格属性：该组件除单价、时长外的其他影响折扣定价的属性信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PriceInfo []*string `json:"PriceInfo,omitnil,omitempty" name:"PriceInfo"`

	// 关联交易单据ID：和本笔交易关联单据 ID，如，冲销订单，记录原订单、重结订单，退费单记录对应的原购买订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociatedOrder *BillDetailAssociatedOrder `json:"AssociatedOrder,omitnil,omitempty" name:"AssociatedOrder"`

	// 计算说明：特殊交易类型计费结算的详细计算说明，如退费及变配
	// 注意：此字段可能返回 null，表示取不到有效值。
	Formula *string `json:"Formula,omitnil,omitempty" name:"Formula"`

	// 计费规则：各产品详细的计费规则官网说明链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	FormulaUrl *string `json:"FormulaUrl,omitnil,omitempty" name:"FormulaUrl"`

	// 账单归属月
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// 账单归属日
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillDay *string `json:"BillDay,omitnil,omitempty" name:"BillDay"`
}

type ExcludedProducts struct {
	// 不适用商品名称
	GoodsName *string `json:"GoodsName,omitnil,omitempty" name:"GoodsName"`

	// postPay后付费/prePay预付费/riPay预留实例/空字符串或者"*"表示全部模式。
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

type GatherResourceSummary struct {
	// 支付者 UIN：支付者的账号 ID，账号 ID 是用户在腾讯云的唯一账号标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 使用者 UIN：实际使用资源的账号 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 操作者 UIN：操作者账号 ID（预付费资源下单或后付费操作开通资源账号的ID或者角色 ID）
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// 实例类型编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例类型：购买的产品服务对应的实例类型，包括资源包、RI、SP、竞价实例。常规实例默认展示“-”
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTypeName *string `json:"InstanceTypeName,omitnil,omitempty" name:"InstanceTypeName"`

	// 资源ID：不同产品因资源形态不同，资源内容不完全相同，如云服务器 CVM 为对应的实例 ID； 若该产品被分拆，则展示产品分拆后的分拆项 ID，如 COS 桶 ID，CDN 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 实例名称：用户在控制台为资源设置的名称，如未设置默认为空；若该产品被分拆，则展示分拆产品分拆后的分拆项资源别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 分账单元唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreeNodeUniqKey *string `json:"TreeNodeUniqKey,omitnil,omitempty" name:"TreeNodeUniqKey"`

	// 分账单元名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TreeNodeUniqKeyName *string `json:"TreeNodeUniqKeyName,omitnil,omitempty" name:"TreeNodeUniqKeyName"`

	// 资源命中公摊规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 资源命中公摊规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessCode *string `json:"BusinessCode,omitnil,omitempty" name:"BusinessCode"`

	// 产品名称：用户所采购的各类云产品
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessCodeName *string `json:"BusinessCodeName,omitnil,omitempty" name:"BusinessCodeName"`

	// 组件名称编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemCode *string `json:"ItemCode,omitnil,omitempty" name:"ItemCode"`

	// 组件名称：用户购买的产品或服务，所包含的具体组件
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemCodeName *string `json:"ItemCodeName,omitnil,omitempty" name:"ItemCodeName"`

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域名称：资源所属地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 分账标签：资源绑定的标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag []*BillTag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 优惠后总价：优惠后总价 =（原价 - 预留实例抵扣原价 - 节省计划抵扣原价）* 折扣率
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 现金账户支出(元)：通过现金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 代金券支出(元)：使用各类优惠券（如代金券、现金券等）支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 赠送账户支出(元)：使用赠送金支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 分成账户支出(元)：通过分成金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// 费用归集类型：费用来源类型，分摊、归集、未分配
	// 0 - 分摊
	// 1 - 归集
	// -1 - 未分配
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllocationType *int64 `json:"AllocationType,omitnil,omitempty" name:"AllocationType"`

	// 当前归属单元信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BelongTreeNodeUniqKey *AllocationTreeNode `json:"BelongTreeNodeUniqKey,omitnil,omitempty" name:"BelongTreeNodeUniqKey"`

	// 当前资源命中公摊规则信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BelongRule *AllocationRule `json:"BelongRule,omitnil,omitempty" name:"BelongRule"`

	// 其它归属单元信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherTreeNodeUniqKeys []*AllocationTreeNode `json:"OtherTreeNodeUniqKeys,omitnil,omitempty" name:"OtherTreeNodeUniqKeys"`

	// 其他命中规则信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherRules []*AllocationRule `json:"OtherRules,omitnil,omitempty" name:"OtherRules"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称：资源归属的项目，用户在控制台给资源自主分配项目，未分配则是默认项目
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 子产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 子产品名称：用户采购的具体产品细分类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCodeName *string `json:"ProductCodeName,omitnil,omitempty" name:"ProductCodeName"`

	// 计费模式编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 计费模式：资源的计费模式，区分为包年包月和按量计费
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// 交易类型编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 交易类型：明细交易类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionTypeName *string `json:"ActionTypeName,omitnil,omitempty" name:"ActionTypeName"`

	// 分拆项 ID：涉及分拆产品的分拆后的分拆项 ID，如 COS 桶 ID，CDN 域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: SplitItemId is deprecated.
	SplitItemId *string `json:"SplitItemId,omitnil,omitempty" name:"SplitItemId"`

	// 分拆项名称：涉及分拆产品的分拆后的分拆项
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: SplitItemName is deprecated.
	SplitItemName *string `json:"SplitItemName,omitnil,omitempty" name:"SplitItemName"`
}

type JsonObject struct {
	// key值
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// value值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type PayDealsRequestParams struct {
	// 需要支付的一个或者多个子订单号，与BigDealIds字段两者必须且仅传一个参数
	OrderIds []*string `json:"OrderIds,omitnil,omitempty" name:"OrderIds"`

	// 是否自动使用代金券,1:是,0否,默认0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表,目前仅支持指定一张代金券
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 需要支付的一个或者多个大订单号，与OrderIds字段两者必须且仅传一个参数
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`
}

type PayDealsRequest struct {
	*tchttp.BaseRequest
	
	// 需要支付的一个或者多个子订单号，与BigDealIds字段两者必须且仅传一个参数
	OrderIds []*string `json:"OrderIds,omitnil,omitempty" name:"OrderIds"`

	// 是否自动使用代金券,1:是,0否,默认0
	AutoVoucher *int64 `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 代金券ID列表,目前仅支持指定一张代金券
	VoucherIds []*string `json:"VoucherIds,omitnil,omitempty" name:"VoucherIds"`

	// 需要支付的一个或者多个大订单号，与OrderIds字段两者必须且仅传一个参数
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`
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
	OrderIds []*string `json:"OrderIds,omitnil,omitempty" name:"OrderIds"`

	// 此次操作支付成功的资源Id数组
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 此次操作支付成功的大订单号数组
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 计费模式：区分为包年包月和按量计费
	PayModeName *string `json:"PayModeName,omitnil,omitempty" name:"PayModeName"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil,omitempty" name:"RealTotalCostRatio"`

	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 按交易类型汇总消费详情
	Detail []*ActionSummaryOverviewItem `json:"Detail,omitnil,omitempty" name:"Detail"`
}

type ProductInfo struct {
	// 商品详情名称标识
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 商品详情
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ProjectSummaryOverviewItem struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称：资源归属的项目，用户在控制台给资源自主分配项目，未分配则是默认项目
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil,omitempty" name:"RealTotalCostRatio"`

	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// 账单月份，格式2019-08
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type RegionSummaryOverviewItem struct {
	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *string `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域名称：资源所属地域，例如华南地区（广州）
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil,omitempty" name:"RealTotalCostRatio"`

	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// 账单月份，格式2019-08
	BillMonth *string `json:"BillMonth,omitnil,omitempty" name:"BillMonth"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type SavingPlanCoverageDetail struct {
	// 资源 ID：账单中出账对象 ID，不同产品因资源形态不同，资源内容不完全相同，如云服务器 CVM 为对应的实例 ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 子产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// 费用起始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 费用结束日期，格式yyyy-MM-dd，目前与StartDate相等
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 节省计划覆盖金额（即节省计划支付金额）
	SpCoveredAmount *float64 `json:"SpCoveredAmount,omitnil,omitempty" name:"SpCoveredAmount"`

	// 节省计划未覆盖金额（即优惠后总价）
	SpUncoveredAmount *float64 `json:"SpUncoveredAmount,omitnil,omitempty" name:"SpUncoveredAmount"`

	// 总支出（即节省计划未覆盖金额 + 节省计划覆盖金额）
	TotalRealAmount *float64 `json:"TotalRealAmount,omitnil,omitempty" name:"TotalRealAmount"`

	// 按量计费预期金额（即折前价 * 折扣）
	ExpectedAmount *float64 `json:"ExpectedAmount,omitnil,omitempty" name:"ExpectedAmount"`

	// 覆盖率结果，取值[0, 100]
	SpCoverage *float64 `json:"SpCoverage,omitnil,omitempty" name:"SpCoverage"`
}

type SavingPlanCoverageRate struct {
	// 聚合时间维度，按天聚合格式为yyyy-MM-dd，按月聚合格式为yyyy-MM
	DatePoint *string `json:"DatePoint,omitnil,omitempty" name:"DatePoint"`

	// 覆盖率结果，取值[0, 100]
	Rate *float64 `json:"Rate,omitnil,omitempty" name:"Rate"`
}

type SavingPlanOverviewDetail struct {
	// 节省计划类型
	SpType *string `json:"SpType,omitnil,omitempty" name:"SpType"`

	// 支付类型
	PayType *uint64 `json:"PayType,omitnil,omitempty" name:"PayType"`

	// 支付金额（单位：元）
	PayAmount *string `json:"PayAmount,omitnil,omitempty" name:"PayAmount"`

	// 开始时间 yyyy-mm-dd HH:mm:ss格式
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间 yyyy-mm-dd HH:mm:ss格式
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 购买时间 yyyy-mm-dd HH:mm:ss格式
	BuyTime *string `json:"BuyTime,omitnil,omitempty" name:"BuyTime"`

	// 状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 累计节省金额（单位：元）
	SavingAmount *string `json:"SavingAmount,omitnil,omitempty" name:"SavingAmount"`

	// 地域
	Region []*string `json:"Region,omitnil,omitempty" name:"Region"`
}

type SavingPlanUsageDetail struct {
	// 节省计划类型
	SpType *string `json:"SpType,omitnil,omitempty" name:"SpType"`

	// 节省计划状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 累计抵扣的金额（单位：元）
	DeductAmount *string `json:"DeductAmount,omitnil,omitempty" name:"DeductAmount"`

	// 累计承诺消费金额（单位：元）
	PromiseAmount *string `json:"PromiseAmount,omitnil,omitempty" name:"PromiseAmount"`

	// 累计净节省金额（单位：元）
	NetSavings *string `json:"NetSavings,omitnil,omitempty" name:"NetSavings"`

	// 使用率
	UtilizationRate *float64 `json:"UtilizationRate,omitnil,omitempty" name:"UtilizationRate"`

	// 累计流失金额（单位：元）
	LossAmount *string `json:"LossAmount,omitnil,omitempty" name:"LossAmount"`

	// 累计按量计费预期金额（单位：元）
	DosageAmount *string `json:"DosageAmount,omitnil,omitempty" name:"DosageAmount"`

	// 累计成本金额（单位：元）
	CostAmount *string `json:"CostAmount,omitnil,omitempty" name:"CostAmount"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region []*string `json:"Region,omitnil,omitempty" name:"Region"`
}

type SummaryDetail struct {
	// 账单维度编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupKey *string `json:"GroupKey,omitnil,omitempty" name:"GroupKey"`

	// 账单维度值
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupValue *string `json:"GroupValue,omitnil,omitempty" name:"GroupValue"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 优惠后总价
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 现金账户支出：通过现金账户支付的金额
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// 产品汇总信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Business []*BusinessSummaryInfo `json:"Business,omitnil,omitempty" name:"Business"`
}

type SummaryTotal struct {
	// 优惠后总价
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type TagDataInfo struct {
	// 分账标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签类型，0普通标签，1分账标签
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 设置分账标签时间，普通标签不返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type TagSummaryOverviewItem struct {
	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`

	// 费用所占百分比，两位小数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitnil,omitempty" name:"RealTotalCostRatio"`

	// 优惠后总价
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalCost *string `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 现金账户支出：通过现金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	CashPayAmount *string `json:"CashPayAmount,omitnil,omitempty" name:"CashPayAmount"`

	// 赠送账户支出：使用赠送金支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncentivePayAmount *string `json:"IncentivePayAmount,omitnil,omitempty" name:"IncentivePayAmount"`

	// 优惠券支出：使用各类优惠券（如代金券、现金券等）支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoucherPayAmount *string `json:"VoucherPayAmount,omitnil,omitempty" name:"VoucherPayAmount"`

	// 分成金账户支出：通过分成金账户支付的金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitnil,omitempty" name:"TransferPayAmount"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type UsageDetails struct {
	// 商品名
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 商品细节
	SubProductName *string `json:"SubProductName,omitnil,omitempty" name:"SubProductName"`

	// 产品码	
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 子产品码	
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// 计费项码	
	BillingItemCode *string `json:"BillingItemCode,omitnil,omitempty" name:"BillingItemCode"`

	// 计费细项码	
	SubBillingItemCode *string `json:"SubBillingItemCode,omitnil,omitempty" name:"SubBillingItemCode"`

	// 产品英文名	
	ProductEnName *string `json:"ProductEnName,omitnil,omitempty" name:"ProductEnName"`

	// 子产品英文名	
	SubProductEnName *string `json:"SubProductEnName,omitnil,omitempty" name:"SubProductEnName"`

	// 结算周期	
	CalcUnit *string `json:"CalcUnit,omitnil,omitempty" name:"CalcUnit"`

	// payMode为prepay 且 payScene为common的情况下存在
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`
}

type UsageRecords struct {
	// 使用金额（微分）
	UsedAmount *int64 `json:"UsedAmount,omitnil,omitempty" name:"UsedAmount"`

	// 使用时间
	UsedTime *string `json:"UsedTime,omitnil,omitempty" name:"UsedTime"`

	// 使用记录细节
	UsageDetails []*UsageDetails `json:"UsageDetails,omitnil,omitempty" name:"UsageDetails"`

	// 付费模式
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 查询的券id
	VoucherId *string `json:"VoucherId,omitnil,omitempty" name:"VoucherId"`

	// 交易场景：（adjust：调账、common：正常交易场景）
	PayScene *string `json:"PayScene,omitnil,omitempty" name:"PayScene"`

	// 唯一id,对应交易:预付费的dealName,调账/后付费的outTradeNo
	SeqId *string `json:"SeqId,omitnil,omitempty" name:"SeqId"`
}

type VoucherInfos struct {
	// 代金券拥有者
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 券状态：待使用：unUsed，已使用： used，已发货：delivered，已作废： cancel，已过期：overdue
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 代金券面额（微分）
	NominalValue *int64 `json:"NominalValue,omitnil,omitempty" name:"NominalValue"`

	// 剩余金额（微分）
	Balance *int64 `json:"Balance,omitnil,omitempty" name:"Balance"`

	// 代金券id
	VoucherId *string `json:"VoucherId,omitnil,omitempty" name:"VoucherId"`

	// postPay后付费/prePay预付费/riPay预留实例/空字符串或者'*'表示全部模式
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 付费场景PayMode=postPay时：spotpay-竞价实例,"settle account"-普通后付费PayMode=prePay时：purchase-包年包月新购，renew-包年包月续费（自动续费），modify-包年包月配置变更(变配）PayMode=riPay时：oneOffFee-预留实例预付，hourlyFee-预留实例每小时扣费，*-支持全部付费场景
	PayScene *string `json:"PayScene,omitnil,omitempty" name:"PayScene"`

	// 有效期生效时间
	BeginTime *string `json:"BeginTime,omitnil,omitempty" name:"BeginTime"`

	// 有效期截止时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 适用商品信息
	ApplicableProducts *ApplicableProducts `json:"ApplicableProducts,omitnil,omitempty" name:"ApplicableProducts"`

	// 不适用商品信息
	ExcludedProducts []*ExcludedProducts `json:"ExcludedProducts,omitnil,omitempty" name:"ExcludedProducts"`

	// 使用说明/批次备注
	PolicyRemark *string `json:"PolicyRemark,omitnil,omitempty" name:"PolicyRemark"`

	// 发券时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}