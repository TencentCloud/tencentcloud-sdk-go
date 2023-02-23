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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ActionSummaryOverviewItem struct {
	// 交易类型：包年包月新购/续费/升降配/退款、按量计费扣费、调账补偿/扣费等类型
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 交易类型名称
	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`

	// 实际花费
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`

	// 现金金额
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// 赠送金金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// 代金券金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// 账单月份，格式2019-08
	BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`
}

type ApplicableProducts struct {
	// 适用商品名称，值为“全产品通用”或商品名称组成的string，以","分割。
	GoodsName *string `json:"GoodsName,omitempty" name:"GoodsName"`

	// postPay后付费/prePay预付费/riPay预留实例/空字符串或者"*"表示全部模式。如GoodsName为多个商品名以","分割组成的string，而PayMode为"*"，表示每一件商品的模式都为"*"。
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
}

type BillDetail struct {
	// 产品名称：云产品大类，如云服务器CVM、云数据库MySQL
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`

	// 子产品名称：云产品子类，如云服务器CVM-标准型S1
	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`

	// 计费模式：包年包月和按量计费
	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`

	// 项目:资源所属项目
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 区域：资源所属地域，如华南地区（广州）
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 可用区：资源所属可用区，如广州三区
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 资源实例ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 实例名称
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 交易类型
	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`

	// 订单ID
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 交易ID
	BillId *string `json:"BillId,omitempty" name:"BillId"`

	// 扣费时间
	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`

	// 开始使用时间
	FeeBeginTime *string `json:"FeeBeginTime,omitempty" name:"FeeBeginTime"`

	// 结束使用时间
	FeeEndTime *string `json:"FeeEndTime,omitempty" name:"FeeEndTime"`

	// 组件列表
	ComponentSet []*BillDetailComponent `json:"ComponentSet,omitempty" name:"ComponentSet"`

	// 支付者UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 使用者UIN
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 操作者UIN
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// Tag 信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*BillTagInfo `json:"Tags,omitempty" name:"Tags"`

	// 产品名称代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// 子产品名称代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 交易类型代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 区域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 项目ID:资源所属项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

type BillDetailComponent struct {
	// 组件类型:资源组件类型的名称，如内存、硬盘等
	ComponentCodeName *string `json:"ComponentCodeName,omitempty" name:"ComponentCodeName"`

	// 组件名称:资源组件的名称，如云数据库MySQL-内存等
	ItemCodeName *string `json:"ItemCodeName,omitempty" name:"ItemCodeName"`

	// 组件刊例价:资源组件的原始价格，保持原始粒度
	SinglePrice *string `json:"SinglePrice,omitempty" name:"SinglePrice"`

	// 组件指定价
	SpecifiedPrice *string `json:"SpecifiedPrice,omitempty" name:"SpecifiedPrice"`

	// 价格单位
	PriceUnit *string `json:"PriceUnit,omitempty" name:"PriceUnit"`

	// 组件用量
	UsedAmount *string `json:"UsedAmount,omitempty" name:"UsedAmount"`

	// 组件用量单位
	UsedAmountUnit *string `json:"UsedAmountUnit,omitempty" name:"UsedAmountUnit"`

	// 使用时长
	TimeSpan *string `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 时长单位
	TimeUnitName *string `json:"TimeUnitName,omitempty" name:"TimeUnitName"`

	// 组件原价
	Cost *string `json:"Cost,omitempty" name:"Cost"`

	// 折扣率
	Discount *string `json:"Discount,omitempty" name:"Discount"`

	// 优惠类型
	ReduceType *string `json:"ReduceType,omitempty" name:"ReduceType"`

	// 优惠后总价
	RealCost *string `json:"RealCost,omitempty" name:"RealCost"`

	// 代金券支付金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// 现金支付金额
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// 赠送账户支付金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// 组件类型代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemCode *string `json:"ItemCode,omitempty" name:"ItemCode"`

	// 组件名称代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentCode *string `json:"ComponentCode,omitempty" name:"ComponentCode"`

	// 合同价
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractPrice *string `json:"ContractPrice,omitempty" name:"ContractPrice"`

	// 资源包、预留实例、节省计划、竞价实例这四类特殊实例本身的扣费行为，此字段体现对应的实例类型。枚举值如下：
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 预留实例抵扣的使用时长，时长单位与被抵扣的时长单位保持一致
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiTimeSpan *string `json:"RiTimeSpan,omitempty" name:"RiTimeSpan"`

	// 按组件原价的口径换算的预留实例抵扣金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalCostWithRI *string `json:"OriginalCostWithRI,omitempty" name:"OriginalCostWithRI"`

	// 节省计划可用余额额度范围内，节省计划对于此组件打的折扣率
	// 注意：此字段可能返回 null，表示取不到有效值。
	SPDeductionRate *string `json:"SPDeductionRate,omitempty" name:"SPDeductionRate"`

	// 节省计划抵扣的SP包面值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SPDeduction *string `json:"SPDeduction,omitempty" name:"SPDeduction"`

	// 按组件原价的口径换算的节省计划抵扣金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalCostWithSP *string `json:"OriginalCostWithSP,omitempty" name:"OriginalCostWithSP"`

	// 综合了官网折扣、预留实例抵扣、节省计划抵扣的混合折扣率。若没有预留实例抵扣、节省计划抵扣,混合折扣率等于折扣率
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlendedDiscount *string `json:"BlendedDiscount,omitempty" name:"BlendedDiscount"`
}

type BillResourceSummary struct {
	// 产品名称：云产品大类，如云服务器CVM、云数据库MySQL
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`

	// 子产品名称：云产品子类，如云服务器CVM-标准型S1， 当没有获取到子产品名称时，返回"-"
	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`

	// 计费模式：包年包月和按量计费
	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`

	// 项目
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 地域
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 可用区
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 资源实例ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源实例名称
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 交易类型：包年包月新购/续费/升降配/退款、按量计费扣费、调账补偿/扣费等类型
	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`

	// 订单ID
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 扣费时间
	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`

	// 开始使用时间
	FeeBeginTime *string `json:"FeeBeginTime,omitempty" name:"FeeBeginTime"`

	// 结束使用时间
	FeeEndTime *string `json:"FeeEndTime,omitempty" name:"FeeEndTime"`

	// 配置描述
	ConfigDesc *string `json:"ConfigDesc,omitempty" name:"ConfigDesc"`

	// 扩展字段1
	ExtendField1 *string `json:"ExtendField1,omitempty" name:"ExtendField1"`

	// 扩展字段2
	ExtendField2 *string `json:"ExtendField2,omitempty" name:"ExtendField2"`

	// 原价，单位为元
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`

	// 折扣率
	// 当聚合之后折扣不唯一或者合同价的情况下，返回“-”
	Discount *string `json:"Discount,omitempty" name:"Discount"`

	// 优惠类型
	ReduceType *string `json:"ReduceType,omitempty" name:"ReduceType"`

	// 优惠后总价，单位为元
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 代金券支付金额，单位为元
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// 现金账户支付金额，单位为元
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// 赠送账户支付金额，单位为元
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// 扩展字段3
	ExtendField3 *string `json:"ExtendField3,omitempty" name:"ExtendField3"`

	// 扩展字段4
	ExtendField4 *string `json:"ExtendField4,omitempty" name:"ExtendField4"`

	// 扩展字段5
	ExtendField5 *string `json:"ExtendField5,omitempty" name:"ExtendField5"`

	// Tag 信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*BillTagInfo `json:"Tags,omitempty" name:"Tags"`

	// 付款方uin
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 资源所有者uin,无值则返回"-"
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 操作者uin,无值则返回"-"
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// 产品名称代码
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// 子产品名称代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 区域ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 资源包、预留实例、节省计划、竞价实例这四类特殊实例本身的扣费行为，此字段体现对应的实例类型。枚举值如下：
	// 
	// ri=Standard RI
	// 
	// svp=Savings Plan
	// 
	// si=Spot Instances
	// 
	// rp=Resource Pack
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 按组件原价的口径换算的预留实例抵扣金额
	OriginalCostWithRI *string `json:"OriginalCostWithRI,omitempty" name:"OriginalCostWithRI"`

	// 节省计划抵扣的SP包面值
	SPDeduction *string `json:"SPDeduction,omitempty" name:"SPDeduction"`

	// 按组件原价的口径换算的节省计划抵扣金额
	OriginalCostWithSP *string `json:"OriginalCostWithSP,omitempty" name:"OriginalCostWithSP"`
}

type BillTagInfo struct {
	// 分账标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type BillTransactionInfo struct {
	// 收支类型：deduct 扣费, recharge 充值, return 退费， block 冻结, unblock 解冻
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 流水金额，单位（分）；正数表示入账，负数表示出账
	Amount *int64 `json:"Amount,omitempty" name:"Amount"`

	// 可用余额，单位（分）；正数表示入账，负数表示出账
	Balance *int64 `json:"Balance,omitempty" name:"Balance"`

	// 流水号，如20190131020000236005203583326401
	BillId *string `json:"BillId,omitempty" name:"BillId"`

	// 描述信息
	OperationInfo *string `json:"OperationInfo,omitempty" name:"OperationInfo"`

	// 操作时间"2019-01-31 23:35:10.000"
	OperationTime *string `json:"OperationTime,omitempty" name:"OperationTime"`

	// 现金账户余额，单位（分）
	Cash *int64 `json:"Cash,omitempty" name:"Cash"`

	// 赠送金余额，单位（分）
	Incentive *int64 `json:"Incentive,omitempty" name:"Incentive"`

	// 冻结余额，单位（分）
	Freezing *int64 `json:"Freezing,omitempty" name:"Freezing"`

	// 交易渠道
	PayChannel *string `json:"PayChannel,omitempty" name:"PayChannel"`

	// 扣费模式：trade 包年包月(预付费)，hourh  按量-小时结，hourd 按量-日结，hourm 按量-月结，month 按量-月结
	DeductMode *string `json:"DeductMode,omitempty" name:"DeductMode"`
}

type BusinessSummaryOverviewItem struct {
	// 产品名称代码
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// 产品名称：云产品大类，如云服务器CVM、云数据库MySQL
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`

	// 实际花费
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`

	// 现金金额
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// 赠送金金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// 代金券金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// 账单月份，格式2019-08
	BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`

	// 分成金金额
	TransferPayAmount *string `json:"TransferPayAmount,omitempty" name:"TransferPayAmount"`
}

type BusinessSummaryTotal struct {
	// 总花费
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 代金券金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// 赠送金金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// 现金金额
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`

	// 分成金金额
	TransferPayAmount *string `json:"TransferPayAmount,omitempty" name:"TransferPayAmount"`
}

type ConditionBusiness struct {
	// 产品名称代码
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// 产品名称
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`
}

type ConditionPayMode struct {
	// 付费模式
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 付费模式名称
	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`
}

type ConditionProject struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

type ConditionRegion struct {
	// 地域ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
}

type Conditions struct {
	// 只支持6和12两个值
	TimeRange *uint64 `json:"TimeRange,omitempty" name:"TimeRange"`

	// 产品名称代码
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 地域ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 付费模式，可选prePay和postPay
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 资源关键字
	ResourceKeyword *string `json:"ResourceKeyword,omitempty" name:"ResourceKeyword"`

	// 产品名称代码
	BusinessCodes []*string `json:"BusinessCodes,omitempty" name:"BusinessCodes"`

	// 子产品名称代码
	ProductCodes []*string `json:"ProductCodes,omitempty" name:"ProductCodes"`

	// 地域ID
	RegionIds []*int64 `json:"RegionIds,omitempty" name:"RegionIds"`

	// 项目ID
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// 付费模式，可选prePay和postPay
	PayModes []*string `json:"PayModes,omitempty" name:"PayModes"`

	// 交易类型
	ActionTypes []*string `json:"ActionTypes,omitempty" name:"ActionTypes"`

	// 是否隐藏0元流水
	HideFreeCost *int64 `json:"HideFreeCost,omitempty" name:"HideFreeCost"`

	// 排序规则，可选desc和asc
	OrderByCost *string `json:"OrderByCost,omitempty" name:"OrderByCost"`

	// 交易ID
	BillIds []*string `json:"BillIds,omitempty" name:"BillIds"`

	// 组件编码
	ComponentCodes []*string `json:"ComponentCodes,omitempty" name:"ComponentCodes"`

	// 文件ID
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds"`

	// 文件类型
	FileTypes []*string `json:"FileTypes,omitempty" name:"FileTypes"`

	// 状态
	Status []*uint64 `json:"Status,omitempty" name:"Status"`
}

type ConsumptionBusinessSummaryDataItem struct {
	// 产品名称代码
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// 产品名称
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`

	// 折后总价
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 费用趋势
	Trend *ConsumptionSummaryTrend `json:"Trend,omitempty" name:"Trend"`
}

type ConsumptionProjectSummaryDataItem struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 折后总价
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 趋势
	Trend *ConsumptionSummaryTrend `json:"Trend,omitempty" name:"Trend"`

	// 产品消耗详情
	Business []*ConsumptionBusinessSummaryDataItem `json:"Business,omitempty" name:"Business"`
}

type ConsumptionRegionSummaryDataItem struct {
	// 地域ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 折后总价
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 趋势
	Trend *ConsumptionSummaryTrend `json:"Trend,omitempty" name:"Trend"`

	// 产品消费详情
	Business []*ConsumptionBusinessSummaryDataItem `json:"Business,omitempty" name:"Business"`
}

type ConsumptionResourceSummaryConditionValue struct {
	// 产品列表
	Business []*ConditionBusiness `json:"Business,omitempty" name:"Business"`

	// 项目列表
	Project []*ConditionProject `json:"Project,omitempty" name:"Project"`

	// 地域列表
	Region []*ConditionRegion `json:"Region,omitempty" name:"Region"`

	// 付费模式列表
	PayMode []*ConditionPayMode `json:"PayMode,omitempty" name:"PayMode"`
}

type ConsumptionResourceSummaryDataItem struct {
	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源名称
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 折后总价
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 现金花费
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 地域ID
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 付费模式
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 付费模式名称
	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`

	// 产品名称代码
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// 产品名称
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`

	// 消耗类型
	ConsumptionTypeName *string `json:"ConsumptionTypeName,omitempty" name:"ConsumptionTypeName"`
}

type ConsumptionSummaryTotal struct {
	// 折后总价
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
}

type ConsumptionSummaryTrend struct {
	// 趋势类型，upward上升/downward下降/none无
	Type *string `json:"Type,omitempty" name:"Type"`

	// 趋势值，Type为none是该字段值为null
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type CosDetailSets struct {
	// 存储桶名称
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

	// 用量开始时间
	DosageBeginTime *string `json:"DosageBeginTime,omitempty" name:"DosageBeginTime"`

	// 用量结束时间
	DosageEndTime *string `json:"DosageEndTime,omitempty" name:"DosageEndTime"`

	// 子产品名称
	SubProductCodeName *string `json:"SubProductCodeName,omitempty" name:"SubProductCodeName"`

	// 计费项名称
	BillingItemCodeName *string `json:"BillingItemCodeName,omitempty" name:"BillingItemCodeName"`

	// 用量
	DosageValue *string `json:"DosageValue,omitempty" name:"DosageValue"`

	// 单位
	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

type CostComponentSet struct {
	// 组件类型名称
	ComponentCodeName *string `json:"ComponentCodeName,omitempty" name:"ComponentCodeName"`

	// 组件名称
	ItemCodeName *string `json:"ItemCodeName,omitempty" name:"ItemCodeName"`

	// 刊例价
	SinglePrice *string `json:"SinglePrice,omitempty" name:"SinglePrice"`

	// 刊例价单位
	PriceUnit *string `json:"PriceUnit,omitempty" name:"PriceUnit"`

	// 用量
	UsedAmount *string `json:"UsedAmount,omitempty" name:"UsedAmount"`

	// 用量单位
	UsedAmountUnit *string `json:"UsedAmountUnit,omitempty" name:"UsedAmountUnit"`

	// 原价
	Cost *string `json:"Cost,omitempty" name:"Cost"`

	// 折扣
	Discount *string `json:"Discount,omitempty" name:"Discount"`

	// 折后价
	RealCost *string `json:"RealCost,omitempty" name:"RealCost"`

	// 代金券支付金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// 现金支付金额
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// 赠送金支付金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`
}

type CostDetail struct {
	// 支付者uin
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 产品名称
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`

	// 子产品名称
	ProductCodeName *string `json:"ProductCodeName,omitempty" name:"ProductCodeName"`

	// 计费模式名称
	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 区域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 地区名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源名称
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionTypeName *string `json:"ActionTypeName,omitempty" name:"ActionTypeName"`

	// 订单id
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 交易id
	BillId *string `json:"BillId,omitempty" name:"BillId"`

	// 费用开始时间
	FeeBeginTime *string `json:"FeeBeginTime,omitempty" name:"FeeBeginTime"`

	// 费用结束时间
	FeeEndTime *string `json:"FeeEndTime,omitempty" name:"FeeEndTime"`

	// 组件明细
	ComponentSet []*CostComponentSet `json:"ComponentSet,omitempty" name:"ComponentSet"`

	// 子产品名称代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
}

type Deal struct {
	// 订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 订单状态
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 支付者
	Payer *string `json:"Payer,omitempty" name:"Payer"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 创建人
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 实际支付金额（分）
	RealTotalCost *int64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 代金券抵扣金额（分）
	VoucherDecline *int64 `json:"VoucherDecline,omitempty" name:"VoucherDecline"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 产品分类ID
	GoodsCategoryId *int64 `json:"GoodsCategoryId,omitempty" name:"GoodsCategoryId"`

	// 产品详情
	ProductInfo []*ProductInfo `json:"ProductInfo,omitempty" name:"ProductInfo"`

	// 时长
	TimeSpan *float64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 时间单位
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 货币单位
	Currency *string `json:"Currency,omitempty" name:"Currency"`

	// 折扣率
	Policy *float64 `json:"Policy,omitempty" name:"Policy"`

	// 单价（分）
	Price *float64 `json:"Price,omitempty" name:"Price"`

	// 原价（分）
	TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`

	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子产品编码
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 大订单号
	BigDealId *string `json:"BigDealId,omitempty" name:"BigDealId"`

	// 退费公式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Formula *string `json:"Formula,omitempty" name:"Formula"`

	// 退费涉及订单信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefReturnDeals *string `json:"RefReturnDeals,omitempty" name:"RefReturnDeals"`

	// 付费模式：prePay 预付费 postPay后付费 riPay预留实例
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

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
	Action *string `json:"Action,omitempty" name:"Action"`

	// 产品编码中文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 子产品编码中文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`

	// 订单对应的资源id, 查询参数Limit超过200，将返回null
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId []*string `json:"ResourceId,omitempty" name:"ResourceId"`
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
	Balance *int64 `json:"Balance,omitempty" name:"Balance"`

	// 查询的用户Uin
	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`

	// 当前真实可用余额,单位 分
	RealBalance *float64 `json:"RealBalance,omitempty" name:"RealBalance"`

	// 现金账户余额,单位 分
	CashAccountBalance *float64 `json:"CashAccountBalance,omitempty" name:"CashAccountBalance"`

	// 收益转入账户余额,单位 分
	IncomeIntoAccountBalance *float64 `json:"IncomeIntoAccountBalance,omitempty" name:"IncomeIntoAccountBalance"`

	// 赠送账户余额,单位 分
	PresentAccountBalance *float64 `json:"PresentAccountBalance,omitempty" name:"PresentAccountBalance"`

	// 冻结金额,单位 分
	FreezeAmount *float64 `json:"FreezeAmount,omitempty" name:"FreezeAmount"`

	// 欠费金额,单位 分
	OweAmount *float64 `json:"OweAmount,omitempty" name:"OweAmount"`

	// 是否允许欠费消费
	IsAllowArrears *bool `json:"IsAllowArrears,omitempty" name:"IsAllowArrears"`

	// 是否限制信用额度
	IsCreditLimited *bool `json:"IsCreditLimited,omitempty" name:"IsCreditLimited"`

	// 信用额度
	CreditAmount *float64 `json:"CreditAmount,omitempty" name:"CreditAmount"`

	// 可用信用额度
	CreditBalance *float64 `json:"CreditBalance,omitempty" name:"CreditBalance"`

	// 真实可用信用额度
	RealCreditBalance *float64 `json:"RealCreditBalance,omitempty" name:"RealCreditBalance"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeBillDetailRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	PeriodType *string `json:"PeriodType,omitempty" name:"PeriodType"`

	// 月份，格式为yyyy-mm，Month和BeginTime&EndTime必传一个，如果有传BeginTime&EndTime则Month字段无效。不能早于开通账单2.0的月份，最多可拉取24个月内的数据。
	Month *string `json:"Month,omitempty" name:"Month"`

	// 周期开始时间，格式为Y-m-d H:i:s，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传。不能早于开通账单2.0的月份，最多可拉取24个月内的数据。(不支持跨月查询)
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 周期结束时间，格式为Y-m-d H:i:s，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传。不能早于开通账单2.0的月份，最多可拉取24个月内的数据。（不支持跨月查询）
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`

	// 已废弃参数，未开放
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 付费模式 prePay/postPay
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 查询交易类型（请使用交易类型名称入参），入参示例枚举如下：
	// 包年包月新购
	// 包年包月续费
	// 包年包月配置变更
	// 包年包月退款
	// 按量计费扣费
	// 按量计费小时结
	// 按量计费日结
	// 按量计费月结
	// 线下项目扣费
	// 线下产品扣费
	// 调账扣费
	// 调账补偿
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
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 项目ID:资源所属项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`
}

type DescribeBillDetailRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	PeriodType *string `json:"PeriodType,omitempty" name:"PeriodType"`

	// 月份，格式为yyyy-mm，Month和BeginTime&EndTime必传一个，如果有传BeginTime&EndTime则Month字段无效。不能早于开通账单2.0的月份，最多可拉取24个月内的数据。
	Month *string `json:"Month,omitempty" name:"Month"`

	// 周期开始时间，格式为Y-m-d H:i:s，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传。不能早于开通账单2.0的月份，最多可拉取24个月内的数据。(不支持跨月查询)
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 周期结束时间，格式为Y-m-d H:i:s，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传。不能早于开通账单2.0的月份，最多可拉取24个月内的数据。（不支持跨月查询）
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`

	// 已废弃参数，未开放
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 付费模式 prePay/postPay
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 查询交易类型（请使用交易类型名称入参），入参示例枚举如下：
	// 包年包月新购
	// 包年包月续费
	// 包年包月配置变更
	// 包年包月退款
	// 按量计费扣费
	// 按量计费小时结
	// 按量计费日结
	// 按量计费月结
	// 线下项目扣费
	// 线下产品扣费
	// 调账扣费
	// 调账补偿
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
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 项目ID:资源所属项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillDetailResponseParams struct {
	// 详情列表
	DetailSet []*BillDetail `json:"DetailSet,omitempty" name:"DetailSet"`

	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeBillListRequestParams struct {
	// 查询范围的起始时间（包含）
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询范围的结束时间（包含）
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 翻页偏移量，初始值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页的限制数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 交易类型： all所有交易类型，recharge充值，return退款，unblock解冻，agentin资金转入，advanced垫付，cash提现，deduct扣费，block冻结，agentout资金转出，repay垫付回款，repayment还款(仅国际信用账户)，adj_refund调增(仅国际信用账户)，adj_deduct调减(仅国际信用账户)
	PayType []*string `json:"PayType,omitempty" name:"PayType"`

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
	SubPayType []*string `json:"SubPayType,omitempty" name:"SubPayType"`

	// 是否返回0元交易金额的交易项，取值：0-不返回，1-返回。不传该参数则不返回
	WithZeroAmount *uint64 `json:"WithZeroAmount,omitempty" name:"WithZeroAmount"`
}

type DescribeBillListRequest struct {
	*tchttp.BaseRequest
	
	// 查询范围的起始时间（包含）
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询范围的结束时间（包含）
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 翻页偏移量，初始值为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页的限制数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 交易类型： all所有交易类型，recharge充值，return退款，unblock解冻，agentin资金转入，advanced垫付，cash提现，deduct扣费，block冻结，agentout资金转出，repay垫付回款，repayment还款(仅国际信用账户)，adj_refund调增(仅国际信用账户)，adj_deduct调减(仅国际信用账户)
	PayType []*string `json:"PayType,omitempty" name:"PayType"`

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
	SubPayType []*string `json:"SubPayType,omitempty" name:"SubPayType"`

	// 是否返回0元交易金额的交易项，取值：0-不返回，1-返回。不传该参数则不返回
	WithZeroAmount *uint64 `json:"WithZeroAmount,omitempty" name:"WithZeroAmount"`
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
	TransactionList []*BillTransactionInfo `json:"TransactionList,omitempty" name:"TransactionList"`

	// 总条数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 退费总额，单位（分）
	ReturnAmount *float64 `json:"ReturnAmount,omitempty" name:"ReturnAmount"`

	// 充值总额，单位（分）
	RechargeAmount *float64 `json:"RechargeAmount,omitempty" name:"RechargeAmount"`

	// 冻结总额，单位（分）
	BlockAmount *float64 `json:"BlockAmount,omitempty" name:"BlockAmount"`

	// 解冻总额，单位（分）
	UnblockAmount *float64 `json:"UnblockAmount,omitempty" name:"UnblockAmount"`

	// 扣费总额，单位（分）
	DeductAmount *float64 `json:"DeductAmount,omitempty" name:"DeductAmount"`

	// 资金转入总额，单位（分）
	AgentInAmount *float64 `json:"AgentInAmount,omitempty" name:"AgentInAmount"`

	// 垫付充值总额，单位（分）
	AdvanceRechargeAmount *float64 `json:"AdvanceRechargeAmount,omitempty" name:"AdvanceRechargeAmount"`

	// 提现扣减总额，单位（分）
	WithdrawAmount *float64 `json:"WithdrawAmount,omitempty" name:"WithdrawAmount"`

	// 资金转出总额，单位（分）
	AgentOutAmount *float64 `json:"AgentOutAmount,omitempty" name:"AgentOutAmount"`

	// 还垫付总额，单位（分）
	AdvancePayAmount *float64 `json:"AdvancePayAmount,omitempty" name:"AdvancePayAmount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeBillResourceSummaryRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 月份，格式为yyyy-mm。不能早于开通账单2.0的月份，最多可拉取24个月内的数据。
	Month *string `json:"Month,omitempty" name:"Month"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	PeriodType *string `json:"PeriodType,omitempty" name:"PeriodType"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`

	// 查询交易类型，如下：
	// 包年包月新购
	// 包年包月续费
	// 包年包月配置变更
	// 包年包月退款
	// 按量计费扣费
	// 按量计费小时结
	// 按量计费日结
	// 按量计费月结
	// 线下项目扣费
	// 线下产品扣费
	// 调账扣费
	// 调账补偿
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
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 付费模式 prePay/postPay
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`
}

type DescribeBillResourceSummaryRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 月份，格式为yyyy-mm。不能早于开通账单2.0的月份，最多可拉取24个月内的数据。
	Month *string `json:"Month,omitempty" name:"Month"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	PeriodType *string `json:"PeriodType,omitempty" name:"PeriodType"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`

	// 查询交易类型，如下：
	// 包年包月新购
	// 包年包月续费
	// 包年包月配置变更
	// 包年包月退款
	// 按量计费扣费
	// 按量计费小时结
	// 按量计费日结
	// 按量计费月结
	// 线下项目扣费
	// 线下产品扣费
	// 调账扣费
	// 调账补偿
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
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 付费模式 prePay/postPay
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 产品名称代码
	// 备注：如需获取当月使用过的BusinessCode，请调用API：<a href="https://cloud.tencent.com/document/product/555/35761">获取产品汇总费用分布</a>
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillResourceSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillResourceSummaryResponseParams struct {
	// 资源汇总列表
	ResourceSummarySet []*BillResourceSummary `json:"ResourceSummarySet,omitempty" name:"ResourceSummarySet"`

	// 资源汇总列表总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
}

type DescribeBillSummaryByPayModeRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
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
	// 数据是否准备好，0未准备好，1准备好
	Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

	// 各付费模式花费分布详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryOverview []*PayModeSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 款项类别，与L0账单上的汇总类别对应。
	// 此参数自账单3.0（即2021-05）之后开始生效。
	// 枚举值：
	// consume-消费
	// refund-退款
	// adjustment-调账
	PayType *string `json:"PayType,omitempty" name:"PayType"`
}

type DescribeBillSummaryByProductRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 款项类别，与L0账单上的汇总类别对应。
	// 此参数自账单3.0（即2021-05）之后开始生效。
	// 枚举值：
	// consume-消费
	// refund-退款
	// adjustment-调账
	PayType *string `json:"PayType,omitempty" name:"PayType"`
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
	// 数据是否准备好，0未准备好，1准备好
	Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

	// 总花费详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryTotal *BusinessSummaryTotal `json:"SummaryTotal,omitempty" name:"SummaryTotal"`

	// 各产品花费分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryOverview []*BusinessSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
}

type DescribeBillSummaryByProjectRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
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
	// 数据是否准备好，0未准备好，1准备好
	Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

	// 各项目花费分布详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryOverview []*ProjectSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
}

type DescribeBillSummaryByRegionRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`
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
	// 数据是否准备好，0未准备好，1准备好
	Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

	// 各地域花费分布详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryOverview []*RegionSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分账标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 分账标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type DescribeBillSummaryByTagRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分账标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 分账标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
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
	// 数据是否准备好，0未准备好，1准备好
	Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

	// 各标签值花费分布详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryOverview []*TagSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SummaryTotal *SummaryTotal `json:"SummaryTotal,omitempty" name:"SummaryTotal"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeCostDetailRequestParams struct {
	// 数量，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 周期开始时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为同一月份，暂不支持跨月拉取。可拉取的数据是开通成本分析后，且距今 24 个月内的数据。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 周期结束时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为同一月份，暂不支持跨月拉取。可拉取的数据是开通成本分析后，且距今 24 个月内的数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`

	// 月份，格式为yyyy-mm，Month和BeginTime&EndTime必传一个，如果有传BeginTime&EndTime则Month字段无效。不能早于开通成本分析的月份，最多可拉取24个月内的数据。
	Month *string `json:"Month,omitempty" name:"Month"`

	// 查询指定产品信息（暂时未开放获取）
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 付费模式 prePay/postPay
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type DescribeCostDetailRequest struct {
	*tchttp.BaseRequest
	
	// 数量，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 周期开始时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为同一月份，暂不支持跨月拉取。可拉取的数据是开通成本分析后，且距今 24 个月内的数据。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 周期结束时间，格式为yyyy-mm-dd hh:ii:ss，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传，且为同一月份，暂不支持跨月拉取。可拉取的数据是开通成本分析后，且距今 24 个月内的数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`

	// 月份，格式为yyyy-mm，Month和BeginTime&EndTime必传一个，如果有传BeginTime&EndTime则Month字段无效。不能早于开通成本分析的月份，最多可拉取24个月内的数据。
	Month *string `json:"Month,omitempty" name:"Month"`

	// 查询指定产品信息（暂时未开放获取）
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 付费模式 prePay/postPay
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
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
	DetailSet []*CostDetail `json:"DetailSet,omitempty" name:"DetailSet"`

	// 记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每次获取数据量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
}

type DescribeCostSummaryByProductRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每次获取数据量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
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
	Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

	// 消耗详情
	Total *ConsumptionSummaryTotal `json:"Total,omitempty" name:"Total"`

	// 消耗按产品汇总详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ConsumptionBusinessSummaryDataItem `json:"Data,omitempty" name:"Data"`

	// 记录数量，NeedRecordNum为0是返回null
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每次获取数据量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
}

type DescribeCostSummaryByProjectRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每次获取数据量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
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
	Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

	// 消耗详情
	Total *ConsumptionSummaryTotal `json:"Total,omitempty" name:"Total"`

	// 消耗按业务汇总详情
	Data []*ConsumptionProjectSummaryDataItem `json:"Data,omitempty" name:"Data"`

	// 记录数量，NeedRecordNum为0是返回null
	RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每次获取数据量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
}

type DescribeCostSummaryByRegionRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每次获取数据量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
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
	Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

	// 消耗详情
	Total *ConsumptionSummaryTotal `json:"Total,omitempty" name:"Total"`

	// 消耗按地域汇总详情
	Data []*ConsumptionRegionSummaryDataItem `json:"Data,omitempty" name:"Data"`

	// 记录数量，NeedRecordNum为0是返回null
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每次获取数据量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`

	// 是否需要返回过滤条件，0不需要，1需要，默认不需要
	NeedConditionValue *uint64 `json:"NeedConditionValue,omitempty" name:"NeedConditionValue"`

	// 过滤条件，只支持ResourceKeyword(资源关键字，支持资源id及资源名称模糊查询)，ProjectIds（项目id），RegionIds(地域id)，PayModes(付费模式，可选prePay和postPay)，HideFreeCost（是否隐藏0元流水，可选0和1），OrderByCost（按费用排序规则，可选desc和asc）
	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
}

type DescribeCostSummaryByResourceRequest struct {
	*tchttp.BaseRequest
	
	// 目前必须和EndTime相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前必须和BeginTime为相同月份，不支持跨月查询，且查询结果是整月数据，例如 BeginTime为2018-09，EndTime 为 2018-09，查询结果是 2018 年 9 月数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每次获取数据量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`

	// 是否需要返回过滤条件，0不需要，1需要，默认不需要
	NeedConditionValue *uint64 `json:"NeedConditionValue,omitempty" name:"NeedConditionValue"`

	// 过滤条件，只支持ResourceKeyword(资源关键字，支持资源id及资源名称模糊查询)，ProjectIds（项目id），RegionIds(地域id)，PayModes(付费模式，可选prePay和postPay)，HideFreeCost（是否隐藏0元流水，可选0和1），OrderByCost（按费用排序规则，可选desc和asc）
	Conditions *Conditions `json:"Conditions,omitempty" name:"Conditions"`
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
	Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

	// 消耗详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *ConsumptionSummaryTotal `json:"Total,omitempty" name:"Total"`

	// 过滤条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionValue *ConsumptionResourceSummaryConditionValue `json:"ConditionValue,omitempty" name:"ConditionValue"`

	// 记录数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`

	// 资源消耗详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ConsumptionResourceSummaryDataItem `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 一页多少条数据，默认是20条，最大不超过1000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 第多少页，从0开始，默认是0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

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
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 大订单号
	BigDealId *string `json:"BigDealId,omitempty" name:"BigDealId"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type DescribeDealsByCondRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 一页多少条数据，默认是20条，最大不超过1000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 第多少页，从0开始，默认是0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

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
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 大订单号
	BigDealId *string `json:"BigDealId,omitempty" name:"BigDealId"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
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
	Deals []*Deal `json:"Deals,omitempty" name:"Deals"`

	// 订单总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 查询用量结束时间，例如：2020-09-30（与开始时间同月，不支持跨月查询）
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// COS 存储桶名称，可通过Get Service 接口是用来获取请求者名下的所有存储空间列表（Bucket list）https://cloud.tencent.com/document/product/436/8291
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
}

type DescribeDosageCosDetailByDateRequest struct {
	*tchttp.BaseRequest
	
	// 查询用量开始时间，例如：2020-09-01
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 查询用量结束时间，例如：2020-09-30（与开始时间同月，不支持跨月查询）
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// COS 存储桶名称，可通过Get Service 接口是用来获取请求者名下的所有存储空间列表（Bucket list）https://cloud.tencent.com/document/product/436/8291
	BucketName *string `json:"BucketName,omitempty" name:"BucketName"`
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
	DetailSets []*CosDetailSets `json:"DetailSets,omitempty" name:"DetailSets"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 查询账单结束日期，如 2019-01-01， 时间跨度不超过7天
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

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
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 查询域名 例如 www.qq.com
	// 非CDN业务查询时传入空字符串，返回的值为空
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 1、如果为空，则返回EIP或CLB所有实例的明细；
	// 2、如果传入实例名，则返回该实例明细
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type DescribeDosageDetailByDateRequest struct {
	*tchttp.BaseRequest
	
	// 查询账单开始日期，如 2019-01-01
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 查询账单结束日期，如 2019-01-01， 时间跨度不超过7天
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

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
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 查询域名 例如 www.qq.com
	// 非CDN业务查询时传入空字符串，返回的值为空
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 1、如果为空，则返回EIP或CLB所有实例的明细；
	// 2、如果传入实例名，则返回该实例明细
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
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
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 用量数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetailSets []*DetailSet `json:"DetailSets,omitempty" name:"DetailSets"`

	// 错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetMsg *string `json:"RetMsg,omitempty" name:"RetMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeVoucherInfoRequestParams struct {
	// 一页多少条数据，默认是20条，最大不超过1000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 第多少页，默认是1
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 券状态：待使用：unUsed，已使用： used，已发货：delivered，已作废： cancel，已过期：overdue
	Status *string `json:"Status,omitempty" name:"Status"`

	// 代金券id
	VoucherId *string `json:"VoucherId,omitempty" name:"VoucherId"`

	// 代金券订单id
	CodeId *string `json:"CodeId,omitempty" name:"CodeId"`

	// 商品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 活动id
	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

	// 代金券名称
	VoucherName *string `json:"VoucherName,omitempty" name:"VoucherName"`

	// 发放开始时间
	TimeFrom *string `json:"TimeFrom,omitempty" name:"TimeFrom"`

	// 发放结束时间
	TimeTo *string `json:"TimeTo,omitempty" name:"TimeTo"`

	// 指定排序字段：BeginTime开始时间、EndTime到期时间、CreateTime创建时间
	SortField *string `json:"SortField,omitempty" name:"SortField"`

	// 指定升序降序：desc、asc
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`

	// 付费模式，postPay后付费/prePay预付费/riPay预留实例/""或者"*"表示全部模式，如果payMode为""或"*"，那么productCode与subProductCode必须传空
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 付费场景PayMode=postPay时：spotpay-竞价实例,"settle account"-普通后付费PayMode=prePay时：purchase-包年包月新购，renew-包年包月续费（自动续费），modify-包年包月配置变更(变配）PayMode=riPay时：oneOffFee-预留实例预付，hourlyFee-预留实例每小时扣费，*-支持全部付费场景
	PayScene *string `json:"PayScene,omitempty" name:"PayScene"`

	// 操作人，默认就是用户uin
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type DescribeVoucherInfoRequest struct {
	*tchttp.BaseRequest
	
	// 一页多少条数据，默认是20条，最大不超过1000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 第多少页，默认是1
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 券状态：待使用：unUsed，已使用： used，已发货：delivered，已作废： cancel，已过期：overdue
	Status *string `json:"Status,omitempty" name:"Status"`

	// 代金券id
	VoucherId *string `json:"VoucherId,omitempty" name:"VoucherId"`

	// 代金券订单id
	CodeId *string `json:"CodeId,omitempty" name:"CodeId"`

	// 商品码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 活动id
	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

	// 代金券名称
	VoucherName *string `json:"VoucherName,omitempty" name:"VoucherName"`

	// 发放开始时间
	TimeFrom *string `json:"TimeFrom,omitempty" name:"TimeFrom"`

	// 发放结束时间
	TimeTo *string `json:"TimeTo,omitempty" name:"TimeTo"`

	// 指定排序字段：BeginTime开始时间、EndTime到期时间、CreateTime创建时间
	SortField *string `json:"SortField,omitempty" name:"SortField"`

	// 指定升序降序：desc、asc
	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`

	// 付费模式，postPay后付费/prePay预付费/riPay预留实例/""或者"*"表示全部模式，如果payMode为""或"*"，那么productCode与subProductCode必须传空
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 付费场景PayMode=postPay时：spotpay-竞价实例,"settle account"-普通后付费PayMode=prePay时：purchase-包年包月新购，renew-包年包月续费（自动续费），modify-包年包月配置变更(变配）PayMode=riPay时：oneOffFee-预留实例预付，hourlyFee-预留实例每小时扣费，*-支持全部付费场景
	PayScene *string `json:"PayScene,omitempty" name:"PayScene"`

	// 操作人，默认就是用户uin
	Operator *string `json:"Operator,omitempty" name:"Operator"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVoucherInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVoucherInfoResponseParams struct {
	// 券总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 总余额（微分）
	TotalBalance *int64 `json:"TotalBalance,omitempty" name:"TotalBalance"`

	// 代金券相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoucherInfos []*VoucherInfos `json:"VoucherInfos,omitempty" name:"VoucherInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 第多少页，默认是1
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 代金券id
	VoucherId *string `json:"VoucherId,omitempty" name:"VoucherId"`

	// 操作人，默认就是用户uin
	Operator *string `json:"Operator,omitempty" name:"Operator"`
}

type DescribeVoucherUsageDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 一页多少条数据，默认是20条，最大不超过1000
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 第多少页，默认是1
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 代金券id
	VoucherId *string `json:"VoucherId,omitempty" name:"VoucherId"`

	// 操作人，默认就是用户uin
	Operator *string `json:"Operator,omitempty" name:"Operator"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 总已用金额（微分）
	TotalUsedAmount *int64 `json:"TotalUsedAmount,omitempty" name:"TotalUsedAmount"`

	// 代金券使用记录细节
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsageRecords []*UsageRecords `json:"UsageRecords,omitempty" name:"UsageRecords"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Time *string `json:"Time,omitempty" name:"Time"`

	// 值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type DetailSet struct {
	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 使用数据明细
	DetailPoints []*DetailPoint `json:"DetailPoints,omitempty" name:"DetailPoints"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type ExcludedProducts struct {
	// 不适用商品名称
	GoodsName *string `json:"GoodsName,omitempty" name:"GoodsName"`

	// postPay后付费/prePay预付费/riPay预留实例/空字符串或者"*"表示全部模式。
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
}

// Predefined struct for user
type PayDealsRequestParams struct {
	// 需要支付的一个或者多个子订单号，与BigDealIds字段两者必须且仅传一个参数
	OrderIds []*string `json:"OrderIds,omitempty" name:"OrderIds"`

	// 是否自动使用代金券,1:是,0否,默认0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表,目前仅支持指定一张代金券
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 需要支付的一个或者多个大订单号，与OrderIds字段两者必须且仅传一个参数
	BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds"`
}

type PayDealsRequest struct {
	*tchttp.BaseRequest
	
	// 需要支付的一个或者多个子订单号，与BigDealIds字段两者必须且仅传一个参数
	OrderIds []*string `json:"OrderIds,omitempty" name:"OrderIds"`

	// 是否自动使用代金券,1:是,0否,默认0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表,目前仅支持指定一张代金券
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds"`

	// 需要支付的一个或者多个大订单号，与OrderIds字段两者必须且仅传一个参数
	BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds"`
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
	OrderIds []*string `json:"OrderIds,omitempty" name:"OrderIds"`

	// 此次操作支付成功的资源Id数组
	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`

	// 此次操作支付成功的大订单号数组
	BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 付费模式
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 付费模式名称
	PayModeName *string `json:"PayModeName,omitempty" name:"PayModeName"`

	// 实际花费
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`

	// 按交易类型：包年包月新购/续费/升降配/退款、按量计费扣费、调账补偿/扣费等类型汇总消费详情
	Detail []*ActionSummaryOverviewItem `json:"Detail,omitempty" name:"Detail"`

	// 现金金额
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// 赠送金金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// 代金券金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`

	// 分成金金额
	TransferPayAmount *string `json:"TransferPayAmount,omitempty" name:"TransferPayAmount"`
}

type ProductInfo struct {
	// 商品详情名称标识
	Name *string `json:"Name,omitempty" name:"Name"`

	// 商品详情
	Value *string `json:"Value,omitempty" name:"Value"`
}

type ProjectSummaryOverviewItem struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 实际花费
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`

	// 现金金额
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// 赠送金金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// 代金券金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// 账单月份，格式2019-08
	BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`

	// 分成金金额
	TransferPayAmount *string `json:"TransferPayAmount,omitempty" name:"TransferPayAmount"`
}

type RegionSummaryOverviewItem struct {
	// 地域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`

	// 地域名称
	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`

	// 实际花费
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 费用所占百分比，两位小数
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`

	// 现金金额
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// 赠送金金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// 代金券金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// 账单月份，格式2019-08
	BillMonth *string `json:"BillMonth,omitempty" name:"BillMonth"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`

	// 分成金金额
	TransferPayAmount *string `json:"TransferPayAmount,omitempty" name:"TransferPayAmount"`
}

type SummaryTotal struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`
}

type TagSummaryOverviewItem struct {
	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`

	// 实际花费
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalCost *string `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 费用所占百分比，两位小数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTotalCostRatio *string `json:"RealTotalCostRatio,omitempty" name:"RealTotalCostRatio"`

	// 原价，单位为元。TotalCost字段自账单3.0（即2021-05）之后开始生效，账单3.0之前返回"-"。合同价的情况下，TotalCost字段与官网价格存在差异，也返回“-”。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCost *string `json:"TotalCost,omitempty" name:"TotalCost"`

	// 现金金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// 赠送金金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// 代金券金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`

	// 分成金金额
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransferPayAmount *string `json:"TransferPayAmount,omitempty" name:"TransferPayAmount"`
}

type UsageDetails struct {
	// 商品名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 商品细节
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductName *string `json:"SubProductName,omitempty" name:"SubProductName"`
}

type UsageRecords struct {
	// 使用金额（微分）
	UsedAmount *int64 `json:"UsedAmount,omitempty" name:"UsedAmount"`

	// 使用时间
	UsedTime *string `json:"UsedTime,omitempty" name:"UsedTime"`

	// 使用记录细节
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsageDetails []*UsageDetails `json:"UsageDetails,omitempty" name:"UsageDetails"`
}

type VoucherInfos struct {
	// 代金券拥有者
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 券状态：待使用：unUsed，已使用： used，已发货：delivered，已作废： cancel，已过期：overdue
	Status *string `json:"Status,omitempty" name:"Status"`

	// 代金券面额（微分）
	NominalValue *int64 `json:"NominalValue,omitempty" name:"NominalValue"`

	// 剩余金额（微分）
	Balance *int64 `json:"Balance,omitempty" name:"Balance"`

	// 代金券id
	VoucherId *string `json:"VoucherId,omitempty" name:"VoucherId"`

	// postPay后付费/prePay预付费/riPay预留实例/空字符串或者'*'表示全部模式
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 付费场景PayMode=postPay时：spotpay-竞价实例,"settle account"-普通后付费PayMode=prePay时：purchase-包年包月新购，renew-包年包月续费（自动续费），modify-包年包月配置变更(变配）PayMode=riPay时：oneOffFee-预留实例预付，hourlyFee-预留实例每小时扣费，*-支持全部付费场景
	PayScene *string `json:"PayScene,omitempty" name:"PayScene"`

	// 有效期生效时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 有效期截止时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 适用商品信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplicableProducts *ApplicableProducts `json:"ApplicableProducts,omitempty" name:"ApplicableProducts"`

	// 不适用商品信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludedProducts []*ExcludedProducts `json:"ExcludedProducts,omitempty" name:"ExcludedProducts"`
}