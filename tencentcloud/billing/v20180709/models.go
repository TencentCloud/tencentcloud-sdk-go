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
	ComponentSet []*BillDetailComponent `json:"ComponentSet,omitempty" name:"ComponentSet" list`

	// 支付者UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 使用者UIN
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 操作者UIN
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// Tag 信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*BillTagInfo `json:"Tags,omitempty" name:"Tags" list`

	// 商品名称代码（未开放的字段）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// 子商品名称代码 （未开放的字段）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 交易类型代码（未开放的字段）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 区域ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
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

	// 组件类型代码（未开放的字段）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemCode *string `json:"ItemCode,omitempty" name:"ItemCode"`

	// 组件名称代码（未开放的字段）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComponentCode *string `json:"ComponentCode,omitempty" name:"ComponentCode"`

	// 合同价
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContractPrice *string `json:"ContractPrice,omitempty" name:"ContractPrice"`
}

type BillResourceSummary struct {

	// 产品名称：云产品大类，如云服务器CVM、云数据库MySQL
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`

	// 子产品：云产品子类，如云服务器CVM-标准型S1， 当没有获取到子产品名称时，返回"-"
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
	Tags []*BillTagInfo `json:"Tags,omitempty" name:"Tags" list`

	// 付款方uin
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 资源所有者uin,无值则返回"-"
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 操作者uin,无值则返回"-"
	OperateUin *string `json:"OperateUin,omitempty" name:"OperateUin"`

	// 商品名称代码
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// 子商品名称代码
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 区域ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
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

	// 产品码
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
}

type ConditionBusiness struct {

	// 产品码
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

	// 产品编码
	BusinessCode *string `json:"BusinessCode,omitempty" name:"BusinessCode"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 地域ID
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 付费模式，可选prePay和postPay
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 资源关键字
	ResourceKeyword *string `json:"ResourceKeyword,omitempty" name:"ResourceKeyword"`

	// 产品编码
	BusinessCodes []*string `json:"BusinessCodes,omitempty" name:"BusinessCodes" list`

	// 子产品编码
	ProductCodes []*string `json:"ProductCodes,omitempty" name:"ProductCodes" list`

	// 地域ID
	RegionIds []*int64 `json:"RegionIds,omitempty" name:"RegionIds" list`

	// 项目ID
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// 付费模式，可选prePay和postPay
	PayModes []*string `json:"PayModes,omitempty" name:"PayModes" list`

	// 交易类型
	ActionTypes []*string `json:"ActionTypes,omitempty" name:"ActionTypes" list`

	// 是否隐藏0元流水
	HideFreeCost *int64 `json:"HideFreeCost,omitempty" name:"HideFreeCost"`

	// 排序规则，可选desc和asc
	OrderByCost *string `json:"OrderByCost,omitempty" name:"OrderByCost"`

	// 交易ID
	BillIds []*string `json:"BillIds,omitempty" name:"BillIds" list`

	// 组件编码
	ComponentCodes []*string `json:"ComponentCodes,omitempty" name:"ComponentCodes" list`

	// 文件ID
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds" list`

	// 文件类型
	FileTypes []*string `json:"FileTypes,omitempty" name:"FileTypes" list`

	// 状态
	Status []*uint64 `json:"Status,omitempty" name:"Status" list`
}

type ConsumptionBusinessSummaryDataItem struct {

	// 产品码
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
	Business []*ConsumptionBusinessSummaryDataItem `json:"Business,omitempty" name:"Business" list`
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
	Business []*ConsumptionBusinessSummaryDataItem `json:"Business,omitempty" name:"Business" list`
}

type ConsumptionResourceSummaryConditionValue struct {

	// 产品列表
	Business []*ConditionBusiness `json:"Business,omitempty" name:"Business" list`

	// 项目列表
	Project []*ConditionProject `json:"Project,omitempty" name:"Project" list`

	// 地域列表
	Region []*ConditionRegion `json:"Region,omitempty" name:"Region" list`

	// 付费模式列表
	PayMode []*ConditionPayMode `json:"PayMode,omitempty" name:"PayMode" list`
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

	// 产品码
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

	// 业务名称
	BusinessCodeName *string `json:"BusinessCodeName,omitempty" name:"BusinessCodeName"`

	// 产品名称
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
	ComponentSet []*CostComponentSet `json:"ComponentSet,omitempty" name:"ComponentSet" list`

	// 产品代码
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
	ProductInfo []*ProductInfo `json:"ProductInfo,omitempty" name:"ProductInfo" list`

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
}

type DescribeAccountBalanceRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccountBalanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountBalanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountBalanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 云账户信息中的”展示可用余额”字段，单位为"分"
		Balance *int64 `json:"Balance,omitempty" name:"Balance"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountBalanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountBalanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

	// 查询指定产品信息（暂时未开放获取）
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 付费模式 prePay/postPay
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 查询指定资源信息
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 查询交易类型。如 按量计费日结，按量计费小时结 等
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

func (r *DescribeBillDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 详情列表
		DetailSet []*BillDetail `json:"DetailSet,omitempty" name:"DetailSet" list`

		// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	PayType []*string `json:"PayType,omitempty" name:"PayType" list`

	// 扣费模式，当所选的交易类型中包含扣费deduct时有意义： all所有扣费类型，trade预付费支付，hour_h按量小时结，hour_d按量日结，hour_m按量月结，decompensate调账扣费，other其他扣费
	SubPayType []*string `json:"SubPayType,omitempty" name:"SubPayType" list`

	// 是否返回0元交易金额的交易项，取值：0-不返回，1-返回。不传该参数则不返回
	WithZeroAmount *uint64 `json:"WithZeroAmount,omitempty" name:"WithZeroAmount"`
}

func (r *DescribeBillListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 收支明细列表
		TransactionList []*BillTransactionInfo `json:"TransactionList,omitempty" name:"TransactionList" list`

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

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillResourceSummaryRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 数量，最大值为1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期。需要与费用中心该月份账单的周期保持一致。您可前往[账单概览](https://console.cloud.tencent.com/expense/bill/overview)页面顶部查看确认您的账单统计周期类型。
	PeriodType *string `json:"PeriodType,omitempty" name:"PeriodType"`

	// 月份，格式为yyyy-mm。不能早于开通账单2.0的月份，最多可拉取24个月内的数据。
	Month *string `json:"Month,omitempty" name:"Month"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *int64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`

	// 查询交易类型。如 按量计费日结，按量计费小时结 等
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`
}

func (r *DescribeBillResourceSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillResourceSummaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillResourceSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 资源汇总列表
		ResourceSummarySet []*BillResourceSummary `json:"ResourceSummarySet,omitempty" name:"ResourceSummarySet" list`

		// 资源汇总列表总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillResourceSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillResourceSummaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByPayModeRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillSummaryByPayModeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByPayModeRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByPayModeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据是否准备好，0未准备好，1准备好
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 各付费模式花费分布详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		SummaryOverview []*PayModeSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByPayModeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByPayModeResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByProductRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillSummaryByProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据是否准备好，0未准备好，1准备好
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 总花费详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		SummaryTotal *BusinessSummaryTotal `json:"SummaryTotal,omitempty" name:"SummaryTotal"`

		// 各产品花费分布
	// 注意：此字段可能返回 null，表示取不到有效值。
		SummaryOverview []*BusinessSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByProjectRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillSummaryByProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据是否准备好，0未准备好，1准备好
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 各项目花费分布详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		SummaryOverview []*ProjectSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByRegionRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBillSummaryByRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByRegionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据是否准备好，0未准备好，1准备好
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 各地域花费分布详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		SummaryOverview []*RegionSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByRegionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByTagRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月开始，且必须和EndTime为相同月份，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月结束，且必须和BeginTime为相同月份，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 分账标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
}

func (r *DescribeBillSummaryByTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByTagRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeBillSummaryByTagResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据是否准备好，0未准备好，1准备好
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 各标签值花费分布详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		SummaryOverview []*TagSummaryOverviewItem `json:"SummaryOverview,omitempty" name:"SummaryOverview" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBillSummaryByTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeBillSummaryByTagResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCostDetailRequest struct {
	*tchttp.BaseRequest

	// 数量，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 周期开始时间，格式为Y-m-d H:i:s，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传。不能早于开通成本分析的月份，最多可拉取24个月内的数据。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 周期结束时间，格式为Y-m-d H:i:s，Month和BeginTime&EndTime必传一个，如果有该字段则Month字段无效。BeginTime和EndTime必须一起传。不能早于开通成本分析的月份，最多可拉取24个月内的数据。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 是否需要访问列表的总记录数，用于前端分页
	// 1-表示需要， 0-表示不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`

	// 月份，格式为yyyy-mm，Month和BeginTime&EndTime必传一个，如果有传BeginTime&EndTime则Month字段无效。不能早于开通成本分析的月份，最多可拉取24个月内的数据。
	Month *string `json:"Month,omitempty" name:"Month"`

	// 查询指定产品信息
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

func (r *DescribeCostDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCostDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 消耗明细
	// 注意：此字段可能返回 null，表示取不到有效值。
		DetailSet []*CostDetail `json:"DetailSet,omitempty" name:"DetailSet" list`

		// 记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCostDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCostDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCostSummaryByProductRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月1号 00:00:00，且必须和EndTime为相同月份，不支持跨月查询，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月最后一天 23:59:59，且必须和BeginTime为相同月份，不支持跨月查询，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每次获取数据量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
}

func (r *DescribeCostSummaryByProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCostSummaryByProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCostSummaryByProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据是否准备好，0未准备好，1准备好
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 消耗详情
		Total *ConsumptionSummaryTotal `json:"Total,omitempty" name:"Total"`

		// 消耗按产品汇总详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*ConsumptionBusinessSummaryDataItem `json:"Data,omitempty" name:"Data" list`

		// 记录数量，NeedRecordNum为0是返回null
	// 注意：此字段可能返回 null，表示取不到有效值。
		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCostSummaryByProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCostSummaryByProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCostSummaryByProjectRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月1号 00:00:00，且必须和EndTime为相同月份，不支持跨月查询，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月最后一天 23:59:59，且必须和BeginTime为相同月份，不支持跨月查询，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每次获取数据量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
}

func (r *DescribeCostSummaryByProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCostSummaryByProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCostSummaryByProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据是否准备好，0未准备好，1准备好
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 消耗详情
		Total *ConsumptionSummaryTotal `json:"Total,omitempty" name:"Total"`

		// 消耗按业务汇总详情
		Data []*ConsumptionProjectSummaryDataItem `json:"Data,omitempty" name:"Data" list`

		// 记录数量，NeedRecordNum为0是返回null
		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCostSummaryByProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCostSummaryByProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCostSummaryByRegionRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月1号 00:00:00，且必须和EndTime为相同月份，不支持跨月查询，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月最后一天 23:59:59，且必须和BeginTime为相同月份，不支持跨月查询，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每次获取数据量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 是否需要返回记录数量，0不需要，1需要，默认不需要
	NeedRecordNum *uint64 `json:"NeedRecordNum,omitempty" name:"NeedRecordNum"`
}

func (r *DescribeCostSummaryByRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCostSummaryByRegionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCostSummaryByRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据是否准备好，0未准备好，1准备好
		Ready *uint64 `json:"Ready,omitempty" name:"Ready"`

		// 消耗详情
		Total *ConsumptionSummaryTotal `json:"Total,omitempty" name:"Total"`

		// 消耗按地域汇总详情
		Data []*ConsumptionRegionSummaryDataItem `json:"Data,omitempty" name:"Data" list`

		// 记录数量，NeedRecordNum为0是返回null
	// 注意：此字段可能返回 null，表示取不到有效值。
		RecordNum *uint64 `json:"RecordNum,omitempty" name:"RecordNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCostSummaryByRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCostSummaryByRegionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCostSummaryByResourceRequest struct {
	*tchttp.BaseRequest

	// 查询账单数据的用户UIN
	PayerUin *string `json:"PayerUin,omitempty" name:"PayerUin"`

	// 目前只支持传当月1号 00:00:00，且必须和EndTime为相同月份，不支持跨月查询，例 2018-09-01 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 目前只支持传当月最后一天 23:59:59，且必须和BeginTime为相同月份，不支持跨月查询，例 2018-09-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每次获取数据量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

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

func (r *DescribeCostSummaryByResourceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCostSummaryByResourceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
		Data []*ConsumptionResourceSummaryDataItem `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCostSummaryByResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCostSummaryByResourceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
}

func (r *DescribeDealsByCondRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDealsByCondRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDealsByCondResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订单列表
		Deals []*Deal `json:"Deals,omitempty" name:"Deals" list`

		// 订单总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDealsByCondResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDealsByCondResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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
	// 10180：CDN静态加速流量(国内)
	// 10181：CDN静态加速带宽(国内)
	// 10182：CDN静态加速普通流量
	// 10183：CDN静态加速普通带宽
	// 10231：CDN静态加速流量(海外)
	// 10232：CDN静态加速带宽(海外)
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
	// 非CDN业务查询时值为空
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 1、如果为空，则返回EIP或CLB所有实例的明细；
	// 2、如果传入实例名，则返回该实例明细
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

func (r *DescribeDosageDetailByDateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDosageDetailByDateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDosageDetailByDateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 计量单位
	// 注意：此字段可能返回 null，表示取不到有效值。
		Unit *string `json:"Unit,omitempty" name:"Unit"`

		// 用量数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		DetailSets []*DetailSet `json:"DetailSets,omitempty" name:"DetailSets" list`

		// 错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
		RetCode *int64 `json:"RetCode,omitempty" name:"RetCode"`

		// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		RetMsg *string `json:"RetMsg,omitempty" name:"RetMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDosageDetailByDateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDosageDetailByDateResponse) FromJsonString(s string) error {
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
	DetailPoints []*DetailPoint `json:"DetailPoints,omitempty" name:"DetailPoints" list`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type PayDealsRequest struct {
	*tchttp.BaseRequest

	// 需要支付的一个或者多个子订单号，与BigDealIds字段两者必须且仅传一个参数
	OrderIds []*string `json:"OrderIds,omitempty" name:"OrderIds" list`

	// 是否自动使用代金券,1:是,0否,默认0
	AutoVoucher *int64 `json:"AutoVoucher,omitempty" name:"AutoVoucher"`

	// 代金券ID列表,目前仅支持指定一张代金券
	VoucherIds []*string `json:"VoucherIds,omitempty" name:"VoucherIds" list`

	// 需要支付的一个或者多个大订单号，与OrderIds字段两者必须且仅传一个参数
	BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds" list`
}

func (r *PayDealsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PayDealsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PayDealsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 此次操作支付成功的子订单号数组
		OrderIds []*string `json:"OrderIds,omitempty" name:"OrderIds" list`

		// 此次操作支付成功的资源Id数组
		ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds" list`

		// 此次操作支付成功的大订单号数组
		BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PayDealsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

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
	Detail []*ActionSummaryOverviewItem `json:"Detail,omitempty" name:"Detail" list`

	// 现金金额
	CashPayAmount *string `json:"CashPayAmount,omitempty" name:"CashPayAmount"`

	// 赠送金金额
	IncentivePayAmount *string `json:"IncentivePayAmount,omitempty" name:"IncentivePayAmount"`

	// 代金券金额
	VoucherPayAmount *string `json:"VoucherPayAmount,omitempty" name:"VoucherPayAmount"`
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
}
