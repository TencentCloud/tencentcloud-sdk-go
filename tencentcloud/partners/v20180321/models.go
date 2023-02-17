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

package v20180321

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AgentAuditedClient struct {
	// 代理商账号ID
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 代客账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 代客审核通过时间戳
	AgentTime *string `json:"AgentTime,omitempty" name:"AgentTime"`

	// 代客类型，可能值为a/b/c
	ClientFlag *string `json:"ClientFlag,omitempty" name:"ClientFlag"`

	// 代客备注
	ClientRemark *string `json:"ClientRemark,omitempty" name:"ClientRemark"`

	// 代客名称（首选实名认证名称）
	ClientName *string `json:"ClientName,omitempty" name:"ClientName"`

	// 认证类型, 0：个人，1：企业；其他：未认证或无效值
	AuthType *string `json:"AuthType,omitempty" name:"AuthType"`

	// 代客APPID
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 上月消费金额
	LastMonthAmt *int64 `json:"LastMonthAmt,omitempty" name:"LastMonthAmt"`

	// 本月消费金额
	ThisMonthAmt *int64 `json:"ThisMonthAmt,omitempty" name:"ThisMonthAmt"`

	// 是否欠费,0：不欠费；1：欠费
	HasOverdueBill *uint64 `json:"HasOverdueBill,omitempty" name:"HasOverdueBill"`

	// 客户类型：可以为new(自拓)/assign(指派)/old(官网)/direct(直销)/direct_newopp(直销(新商机))/空
	ClientType *string `json:"ClientType,omitempty" name:"ClientType"`

	// 项目类型：可以为self(自拓项目)/platform(合作项目)/repeat(复算项目  )/空
	ProjectType *string `json:"ProjectType,omitempty" name:"ProjectType"`

	// 业务员ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SalesUin *string `json:"SalesUin,omitempty" name:"SalesUin"`

	// 业务员姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SalesName *string `json:"SalesName,omitempty" name:"SalesName"`

	// 代客邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mail *string `json:"Mail,omitempty" name:"Mail"`
}

type AgentBillElem struct {
	// 代理商账号ID
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 订单号，仅对预付费账单有意义
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 代客账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 代客备注名称
	ClientRemark *string `json:"ClientRemark,omitempty" name:"ClientRemark"`

	// 支付时间
	PayTime *string `json:"PayTime,omitempty" name:"PayTime"`

	// 云产品名称
	GoodsType *string `json:"GoodsType,omitempty" name:"GoodsType"`

	// 预付费/后付费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 支付月份
	SettleMonth *string `json:"SettleMonth,omitempty" name:"SettleMonth"`

	// 支付金额，单位分
	Amt *int64 `json:"Amt,omitempty" name:"Amt"`

	// agentpay：代付；selfpay：自付
	PayerMode *string `json:"PayerMode,omitempty" name:"PayerMode"`

	// 客户类型：可以为new(自拓)/assign(指定)/old(官网)/direct(直销)/direct_newopp(直销(新商机))/空
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientType *string `json:"ClientType,omitempty" name:"ClientType"`

	// 项目类型：可以为self(自拓项目)/platform(合作项目)/repeat(复算项目  )/空
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectType *string `json:"ProjectType,omitempty" name:"ProjectType"`

	// 活动ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`
}

type AgentClientElem struct {
	// 代理商账号ID
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 代客账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 代客申请时间戳
	ApplyTime *uint64 `json:"ApplyTime,omitempty" name:"ApplyTime"`

	// 代客类型，可能值为a/b/c
	ClientFlag *string `json:"ClientFlag,omitempty" name:"ClientFlag"`

	// 代客邮箱，打码显示
	Mail *string `json:"Mail,omitempty" name:"Mail"`

	// 代客手机，打码显示
	Phone *string `json:"Phone,omitempty" name:"Phone"`

	// 0表示不欠费，1表示欠费
	HasOverdueBill *uint64 `json:"HasOverdueBill,omitempty" name:"HasOverdueBill"`

	// 1:待代理商审核;2:待腾讯云审核4:待腾讯云渠道审批
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 业务员ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SalesUin *string `json:"SalesUin,omitempty" name:"SalesUin"`

	// 业务员姓名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SalesName *string `json:"SalesName,omitempty" name:"SalesName"`

	// 客户名称，此字段和控制台返回一致。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientName *string `json:"ClientName,omitempty" name:"ClientName"`
}

type AgentDealElem struct {
	// 订单自增 ID【请勿依赖该字段作为唯一标识】
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// 订单号【订单唯一键】
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 商品类型 ID
	GoodsCategoryId *string `json:"GoodsCategoryId,omitempty" name:"GoodsCategoryId"`

	// 订单所有者
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 订单所有者对应 appId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 商品数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodsNum *string `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 价格详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodsPrice *DealGoodsPriceElem `json:"GoodsPrice,omitempty" name:"GoodsPrice"`

	// 下单人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creater *string `json:"Creater,omitempty" name:"Creater"`

	// 下单时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatTime *string `json:"CreatTime,omitempty" name:"CreatTime"`

	// 支付结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayEndTime *string `json:"PayEndTime,omitempty" name:"PayEndTime"`

	// 扣费流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillId *string `json:"BillId,omitempty" name:"BillId"`

	// 支付人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Payer *string `json:"Payer,omitempty" name:"Payer"`

	// 订单状态，中文描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealStatus *string `json:"DealStatus,omitempty" name:"DealStatus"`

	// 订单的状态(1：未支付;2：已支付;3：发货中;4：已发货;5：发货失败;6：已退款;7：已关单;8：订单过期;9：订单已失效;10：产品已失效;11：代付拒绝;12：支付中)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodsName *string `json:"GoodsName,omitempty" name:"GoodsName"`

	// 客户备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientRemark *string `json:"ClientRemark,omitempty" name:"ClientRemark"`

	// 订单操作类型，purchase（新购），renew（续费），modify（配置变更）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 代金券抵扣金额，单位分
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoucherDecline *string `json:"VoucherDecline,omitempty" name:"VoucherDecline"`

	// 大订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BigDealId *string `json:"BigDealId,omitempty" name:"BigDealId"`

	// 客户类型（new：自拓；old：官网；assign：指派；direct：直销；direct_newopp：直销(新商机)）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientType *string `json:"ClientType,omitempty" name:"ClientType"`

	// 项目类型（self：自拓；repeat：直销；platform：官网合作）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectType *string `json:"ProjectType,omitempty" name:"ProjectType"`

	// 业务员账号ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SalesUin *string `json:"SalesUin,omitempty" name:"SalesUin"`

	// 支付方式，0：自付；1：代付
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayerMode *string `json:"PayerMode,omitempty" name:"PayerMode"`

	// 活动ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

	// 订单过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OverdueTime *string `json:"OverdueTime,omitempty" name:"OverdueTime"`

	// 产品详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductInfo []*ProductInfoElem `json:"ProductInfo,omitempty" name:"ProductInfo"`

	// 付款方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 订单更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AgentDealNewElem struct {
	// 订单自增 ID【请勿依赖该字段作为唯一标识】
	DealId *string `json:"DealId,omitempty" name:"DealId"`

	// 订单号【订单唯一键】
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 商品类型 ID
	GoodsCategoryId *string `json:"GoodsCategoryId,omitempty" name:"GoodsCategoryId"`

	// 订单所有者
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 订单所有者对应 appId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 商品数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodsNum *string `json:"GoodsNum,omitempty" name:"GoodsNum"`

	// 价格详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodsPrice *DealGoodsPriceNewElem `json:"GoodsPrice,omitempty" name:"GoodsPrice"`

	// 下单人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creater *string `json:"Creater,omitempty" name:"Creater"`

	// 下单时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatTime *string `json:"CreatTime,omitempty" name:"CreatTime"`

	// 支付结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayEndTime *string `json:"PayEndTime,omitempty" name:"PayEndTime"`

	// 扣费流水号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillId *string `json:"BillId,omitempty" name:"BillId"`

	// 支付人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Payer *string `json:"Payer,omitempty" name:"Payer"`

	// 订单状态，中文描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealStatus *string `json:"DealStatus,omitempty" name:"DealStatus"`

	// 订单的状态(1：未支付;2：已支付;3：发货中;4：已发货;5：发货失败;6：已退款;7：已关单;8：订单过期;9：订单已失效;10：产品已失效;11：代付拒绝;12：支付中)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GoodsName *string `json:"GoodsName,omitempty" name:"GoodsName"`

	// 客户备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientRemark *string `json:"ClientRemark,omitempty" name:"ClientRemark"`

	// 订单操作类型，"purchase":"新购","upgrade":"升配","upConvertExpire":"升配","downgrade":"降配","downConvertExpire":"降配","renew":"续费","refund":"退货","modifyNetworkMode":"调整带宽模式","modifyNetworkSize":"调整带宽大小","preMoveOut":"资源迁出","preMoveIn":"资源迁入","preToPost":"包年包月转按量","modify":"变配","postMoveOut":"资源迁出","postMoveIn":"资源迁入","recoverRefundForward":"调账补偿","recoverPayReserve":"调账补偿","recoverPayForward":"调账扣费","recoverRefundReserve":"调账扣费"
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 代金券抵扣金额，单位分
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoucherDecline *string `json:"VoucherDecline,omitempty" name:"VoucherDecline"`

	// 大订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BigDealId *string `json:"BigDealId,omitempty" name:"BigDealId"`

	// 客户类型（new：自拓；old：官网；assign：指派；direct：直销；direct_newopp：直销(新商机)）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientType *string `json:"ClientType,omitempty" name:"ClientType"`

	// 项目类型（self：自拓；repeat：直销；platform：官网合作）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectType *string `json:"ProjectType,omitempty" name:"ProjectType"`

	// 业务员账号ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SalesUin *string `json:"SalesUin,omitempty" name:"SalesUin"`

	// 支付方式，0：自付；1：代付
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayerMode *string `json:"PayerMode,omitempty" name:"PayerMode"`

	// 活动ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

	// 订单过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	OverdueTime *string `json:"OverdueTime,omitempty" name:"OverdueTime"`

	// 产品详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductInfo []*ProductInfoElem `json:"ProductInfo,omitempty" name:"ProductInfo"`

	// 付款方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PaymentMethod *string `json:"PaymentMethod,omitempty" name:"PaymentMethod"`

	// 订单更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type AgentPayDealsRequestParams struct {
	// 订单所有者uin
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 代付标志，1：代付；0：自付
	AgentPay *uint64 `json:"AgentPay,omitempty" name:"AgentPay"`

	// 订单号数组
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
}

type AgentPayDealsRequest struct {
	*tchttp.BaseRequest
	
	// 订单所有者uin
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 代付标志，1：代付；0：自付
	AgentPay *uint64 `json:"AgentPay,omitempty" name:"AgentPay"`

	// 订单号数组
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
}

func (r *AgentPayDealsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AgentPayDealsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OwnerUin")
	delete(f, "AgentPay")
	delete(f, "DealNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AgentPayDealsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AgentPayDealsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AgentPayDealsResponse struct {
	*tchttp.BaseResponse
	Response *AgentPayDealsResponseParams `json:"Response"`
}

func (r *AgentPayDealsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AgentPayDealsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentSalesmanElem struct {
	// 代理商账号ID
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 业务员ID
	SalesUin *string `json:"SalesUin,omitempty" name:"SalesUin"`

	// 业务员姓名
	SalesName *string `json:"SalesName,omitempty" name:"SalesName"`

	// 业务员创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type AgentTransferMoneyRequestParams struct {
	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 转账金额，单位分
	Amount *uint64 `json:"Amount,omitempty" name:"Amount"`
}

type AgentTransferMoneyRequest struct {
	*tchttp.BaseRequest
	
	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 转账金额，单位分
	Amount *uint64 `json:"Amount,omitempty" name:"Amount"`
}

func (r *AgentTransferMoneyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AgentTransferMoneyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUin")
	delete(f, "Amount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AgentTransferMoneyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AgentTransferMoneyResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AgentTransferMoneyResponse struct {
	*tchttp.BaseResponse
	Response *AgentTransferMoneyResponseParams `json:"Response"`
}

func (r *AgentTransferMoneyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AgentTransferMoneyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignClientsToSalesRequestParams struct {
	// 代客/申请中代客uin列表，最大50条
	ClientUins []*string `json:"ClientUins,omitempty" name:"ClientUins"`

	// 业务员uin
	SalesUin *string `json:"SalesUin,omitempty" name:"SalesUin"`

	// 代客类型:normal-代客 apply-申请中代客
	AssignClientStatus *string `json:"AssignClientStatus,omitempty" name:"AssignClientStatus"`

	// 操作类型:assign-执行分派 cancel-取消分派
	AssignActionType *string `json:"AssignActionType,omitempty" name:"AssignActionType"`
}

type AssignClientsToSalesRequest struct {
	*tchttp.BaseRequest
	
	// 代客/申请中代客uin列表，最大50条
	ClientUins []*string `json:"ClientUins,omitempty" name:"ClientUins"`

	// 业务员uin
	SalesUin *string `json:"SalesUin,omitempty" name:"SalesUin"`

	// 代客类型:normal-代客 apply-申请中代客
	AssignClientStatus *string `json:"AssignClientStatus,omitempty" name:"AssignClientStatus"`

	// 操作类型:assign-执行分派 cancel-取消分派
	AssignActionType *string `json:"AssignActionType,omitempty" name:"AssignActionType"`
}

func (r *AssignClientsToSalesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignClientsToSalesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUins")
	delete(f, "SalesUin")
	delete(f, "AssignClientStatus")
	delete(f, "AssignActionType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssignClientsToSalesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignClientsToSalesResponseParams struct {
	// 处理成功的代客uin列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SucceedUins []*string `json:"SucceedUins,omitempty" name:"SucceedUins"`

	// 处理失败的代客uin列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedUins []*string `json:"FailedUins,omitempty" name:"FailedUins"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssignClientsToSalesResponse struct {
	*tchttp.BaseResponse
	Response *AssignClientsToSalesResponseParams `json:"Response"`
}

func (r *AssignClientsToSalesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignClientsToSalesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AuditApplyClientRequestParams struct {
	// 待审核客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 审核结果，可能的取值：accept/reject
	AuditResult *string `json:"AuditResult,omitempty" name:"AuditResult"`

	// 申请理由，B类客户审核通过时必须填写申请理由
	Note *string `json:"Note,omitempty" name:"Note"`
}

type AuditApplyClientRequest struct {
	*tchttp.BaseRequest
	
	// 待审核客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 审核结果，可能的取值：accept/reject
	AuditResult *string `json:"AuditResult,omitempty" name:"AuditResult"`

	// 申请理由，B类客户审核通过时必须填写申请理由
	Note *string `json:"Note,omitempty" name:"Note"`
}

func (r *AuditApplyClientRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AuditApplyClientRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUin")
	delete(f, "AuditResult")
	delete(f, "Note")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AuditApplyClientRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AuditApplyClientResponseParams struct {
	// 代理商账号ID
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 审核结果，包括accept/reject/qcloudaudit（腾讯云审核）
	AuditResult *string `json:"AuditResult,omitempty" name:"AuditResult"`

	// 关联时间对应的时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentTime *uint64 `json:"AgentTime,omitempty" name:"AgentTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AuditApplyClientResponse struct {
	*tchttp.BaseResponse
	Response *AuditApplyClientResponseParams `json:"Response"`
}

func (r *AuditApplyClientResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AuditApplyClientResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePayRelationForClientRequestParams struct {
	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`
}

type CreatePayRelationForClientRequest struct {
	*tchttp.BaseRequest
	
	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`
}

func (r *CreatePayRelationForClientRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePayRelationForClientRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePayRelationForClientRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePayRelationForClientResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePayRelationForClientResponse struct {
	*tchttp.BaseResponse
	Response *CreatePayRelationForClientResponseParams `json:"Response"`
}

func (r *CreatePayRelationForClientResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePayRelationForClientResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DealGoodsPriceElem struct {
	// 实付金额（单位：分）
	RealTotalCost *uint64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 订单实际金额（不含折扣，单位：分）
	OriginalTotalCost *int64 `json:"OriginalTotalCost,omitempty" name:"OriginalTotalCost"`
}

type DealGoodsPriceNewElem struct {
	// 实付金额（单位：分）
	RealTotalCost *int64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`

	// 原始金额（不含折扣，单位：分）
	OriginalTotalCost *int64 `json:"OriginalTotalCost,omitempty" name:"OriginalTotalCost"`
}

// Predefined struct for user
type DescribeAgentAuditedClientsRequestParams struct {
	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 客户名称。由于涉及隐私，名称打码显示，故名称仅支持打码后的模糊搜索
	ClientName *string `json:"ClientName,omitempty" name:"ClientName"`

	// 客户类型，a/b，类型定义参考代理商相关政策文档
	ClientFlag *string `json:"ClientFlag,omitempty" name:"ClientFlag"`

	// ASC/DESC， 不区分大小写，按审核通过时间排序
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// 客户账号ID列表
	ClientUins []*string `json:"ClientUins,omitempty" name:"ClientUins"`

	// 是否欠费。0：不欠费；1：欠费
	HasOverdueBill *uint64 `json:"HasOverdueBill,omitempty" name:"HasOverdueBill"`

	// 客户备注
	ClientRemark *string `json:"ClientRemark,omitempty" name:"ClientRemark"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 可以为new(自拓)/assign(指派)/old(官网)/direct(直销)/direct_newopp(直销(新商机))/空
	ClientType *string `json:"ClientType,omitempty" name:"ClientType"`

	// 项目类型：可以为self(自拓项目)/platform(合作项目)/repeat(复算项目  )/空
	ProjectType *string `json:"ProjectType,omitempty" name:"ProjectType"`

	// 业务员ID
	SalesUin *string `json:"SalesUin,omitempty" name:"SalesUin"`

	// 业务员姓名（模糊查询）
	SalesName *string `json:"SalesName,omitempty" name:"SalesName"`
}

type DescribeAgentAuditedClientsRequest struct {
	*tchttp.BaseRequest
	
	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 客户名称。由于涉及隐私，名称打码显示，故名称仅支持打码后的模糊搜索
	ClientName *string `json:"ClientName,omitempty" name:"ClientName"`

	// 客户类型，a/b，类型定义参考代理商相关政策文档
	ClientFlag *string `json:"ClientFlag,omitempty" name:"ClientFlag"`

	// ASC/DESC， 不区分大小写，按审核通过时间排序
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// 客户账号ID列表
	ClientUins []*string `json:"ClientUins,omitempty" name:"ClientUins"`

	// 是否欠费。0：不欠费；1：欠费
	HasOverdueBill *uint64 `json:"HasOverdueBill,omitempty" name:"HasOverdueBill"`

	// 客户备注
	ClientRemark *string `json:"ClientRemark,omitempty" name:"ClientRemark"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 可以为new(自拓)/assign(指派)/old(官网)/direct(直销)/direct_newopp(直销(新商机))/空
	ClientType *string `json:"ClientType,omitempty" name:"ClientType"`

	// 项目类型：可以为self(自拓项目)/platform(合作项目)/repeat(复算项目  )/空
	ProjectType *string `json:"ProjectType,omitempty" name:"ProjectType"`

	// 业务员ID
	SalesUin *string `json:"SalesUin,omitempty" name:"SalesUin"`

	// 业务员姓名（模糊查询）
	SalesName *string `json:"SalesName,omitempty" name:"SalesName"`
}

func (r *DescribeAgentAuditedClientsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentAuditedClientsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUin")
	delete(f, "ClientName")
	delete(f, "ClientFlag")
	delete(f, "OrderDirection")
	delete(f, "ClientUins")
	delete(f, "HasOverdueBill")
	delete(f, "ClientRemark")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClientType")
	delete(f, "ProjectType")
	delete(f, "SalesUin")
	delete(f, "SalesName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentAuditedClientsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentAuditedClientsResponseParams struct {
	// 已审核代客列表
	AgentClientSet []*AgentAuditedClient `json:"AgentClientSet,omitempty" name:"AgentClientSet"`

	// 符合条件的代客总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAgentAuditedClientsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentAuditedClientsResponseParams `json:"Response"`
}

func (r *DescribeAgentAuditedClientsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentAuditedClientsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentBillsRequestParams struct {
	// 支付月份，如2018-02
	SettleMonth *string `json:"SettleMonth,omitempty" name:"SettleMonth"`

	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 支付方式，prepay/postpay
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 预付费订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 客户备注名称
	ClientRemark *string `json:"ClientRemark,omitempty" name:"ClientRemark"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeAgentBillsRequest struct {
	*tchttp.BaseRequest
	
	// 支付月份，如2018-02
	SettleMonth *string `json:"SettleMonth,omitempty" name:"SettleMonth"`

	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 支付方式，prepay/postpay
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 预付费订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 客户备注名称
	ClientRemark *string `json:"ClientRemark,omitempty" name:"ClientRemark"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAgentBillsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentBillsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SettleMonth")
	delete(f, "ClientUin")
	delete(f, "PayMode")
	delete(f, "OrderId")
	delete(f, "ClientRemark")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentBillsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentBillsResponseParams struct {
	// 符合查询条件列表总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 业务明细列表
	AgentBillSet []*AgentBillElem `json:"AgentBillSet,omitempty" name:"AgentBillSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAgentBillsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentBillsResponseParams `json:"Response"`
}

func (r *DescribeAgentBillsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentBillsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentClientGradeRequestParams struct {
	// 代客uin
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`
}

type DescribeAgentClientGradeRequest struct {
	*tchttp.BaseRequest
	
	// 代客uin
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`
}

func (r *DescribeAgentClientGradeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentClientGradeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentClientGradeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentClientGradeResponseParams struct {
	// 审核状态：0待审核，1，已审核
	AuditStatus *uint64 `json:"AuditStatus,omitempty" name:"AuditStatus"`

	// 实名认证状态：0，未实名认证，1实名认证
	AuthState *uint64 `json:"AuthState,omitempty" name:"AuthState"`

	// 客户级别
	ClientGrade *string `json:"ClientGrade,omitempty" name:"ClientGrade"`

	// 客户类型：1，个人；2，企业；3，其他
	ClientType *uint64 `json:"ClientType,omitempty" name:"ClientType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAgentClientGradeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentClientGradeResponseParams `json:"Response"`
}

func (r *DescribeAgentClientGradeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentClientGradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentClientsRequestParams struct {
	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 客户名称。由于涉及隐私，名称打码显示，故名称仅支持打码后的模糊搜索
	ClientName *string `json:"ClientName,omitempty" name:"ClientName"`

	// 客户类型，a/b，类型定义参考代理商相关政策文档
	ClientFlag *string `json:"ClientFlag,omitempty" name:"ClientFlag"`

	// ASC/DESC， 不区分大小写，按申请时间排序
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 业务员ID
	SalesUin *string `json:"SalesUin,omitempty" name:"SalesUin"`

	// 业务员姓名（模糊查询）
	SalesName *string `json:"SalesName,omitempty" name:"SalesName"`
}

type DescribeAgentClientsRequest struct {
	*tchttp.BaseRequest
	
	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`

	// 客户名称。由于涉及隐私，名称打码显示，故名称仅支持打码后的模糊搜索
	ClientName *string `json:"ClientName,omitempty" name:"ClientName"`

	// 客户类型，a/b，类型定义参考代理商相关政策文档
	ClientFlag *string `json:"ClientFlag,omitempty" name:"ClientFlag"`

	// ASC/DESC， 不区分大小写，按申请时间排序
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 业务员ID
	SalesUin *string `json:"SalesUin,omitempty" name:"SalesUin"`

	// 业务员姓名（模糊查询）
	SalesName *string `json:"SalesName,omitempty" name:"SalesName"`
}

func (r *DescribeAgentClientsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentClientsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUin")
	delete(f, "ClientName")
	delete(f, "ClientFlag")
	delete(f, "OrderDirection")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SalesUin")
	delete(f, "SalesName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentClientsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentClientsResponseParams struct {
	// 待审核代客列表
	AgentClientSet []*AgentClientElem `json:"AgentClientSet,omitempty" name:"AgentClientSet"`

	// 符合条件的代客总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAgentClientsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentClientsResponseParams `json:"Response"`
}

func (r *DescribeAgentClientsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentClientsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentDealsByCacheRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目 最大200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下单时间范围起始点【请保持时间范围最大90天】
	CreatTimeRangeStart *string `json:"CreatTimeRangeStart,omitempty" name:"CreatTimeRangeStart"`

	// 下单时间范围终止点【请保持时间范围最大90天】
	CreatTimeRangeEnd *string `json:"CreatTimeRangeEnd,omitempty" name:"CreatTimeRangeEnd"`

	// 0:下单时间降序；其他：下单时间升序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 订单的状态(1：未支付;2：已支付;3：发货中;4：已发货;5：发货失败;6：已退款;7：已关单;8：订单过期;9：订单已失效;10：产品已失效;11：代付拒绝;12：支付中)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 下单人账号ID列表
	OwnerUins []*string `json:"OwnerUins,omitempty" name:"OwnerUins"`

	// 子订单号列表
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 大订单号列表
	BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds"`

	// 支付方式，0：自付；1：代付
	PayerMode *uint64 `json:"PayerMode,omitempty" name:"PayerMode"`
}

type DescribeAgentDealsByCacheRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目 最大200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下单时间范围起始点【请保持时间范围最大90天】
	CreatTimeRangeStart *string `json:"CreatTimeRangeStart,omitempty" name:"CreatTimeRangeStart"`

	// 下单时间范围终止点【请保持时间范围最大90天】
	CreatTimeRangeEnd *string `json:"CreatTimeRangeEnd,omitempty" name:"CreatTimeRangeEnd"`

	// 0:下单时间降序；其他：下单时间升序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 订单的状态(1：未支付;2：已支付;3：发货中;4：已发货;5：发货失败;6：已退款;7：已关单;8：订单过期;9：订单已失效;10：产品已失效;11：代付拒绝;12：支付中)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 下单人账号ID列表
	OwnerUins []*string `json:"OwnerUins,omitempty" name:"OwnerUins"`

	// 子订单号列表
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 大订单号列表
	BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds"`

	// 支付方式，0：自付；1：代付
	PayerMode *uint64 `json:"PayerMode,omitempty" name:"PayerMode"`
}

func (r *DescribeAgentDealsByCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentDealsByCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CreatTimeRangeStart")
	delete(f, "CreatTimeRangeEnd")
	delete(f, "Order")
	delete(f, "Status")
	delete(f, "OwnerUins")
	delete(f, "DealNames")
	delete(f, "BigDealIds")
	delete(f, "PayerMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentDealsByCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentDealsByCacheResponseParams struct {
	// 订单数组
	AgentDealSet []*AgentDealNewElem `json:"AgentDealSet,omitempty" name:"AgentDealSet"`

	// 符合条件的订单总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAgentDealsByCacheResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentDealsByCacheResponseParams `json:"Response"`
}

func (r *DescribeAgentDealsByCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentDealsByCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentDealsCacheRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目 最大200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下单时间范围起始点【请保持时间范围最大90天】
	CreatTimeRangeStart *string `json:"CreatTimeRangeStart,omitempty" name:"CreatTimeRangeStart"`

	// 下单时间范围终止点【请保持时间范围最大90天】
	CreatTimeRangeEnd *string `json:"CreatTimeRangeEnd,omitempty" name:"CreatTimeRangeEnd"`

	// 0:下单时间降序；其他：下单时间升序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 订单的状态(1：未支付;2：已支付;3：发货中;4：已发货;5：发货失败;6：已退款;7：已关单;8：订单过期;9：订单已失效;10：产品已失效;11：代付拒绝;12：支付中)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 下单人账号ID列表
	OwnerUins []*string `json:"OwnerUins,omitempty" name:"OwnerUins"`

	// 订单号列表
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 支付方式，0：自付；1：代付
	PayerMode *uint64 `json:"PayerMode,omitempty" name:"PayerMode"`
}

type DescribeAgentDealsCacheRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目 最大200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下单时间范围起始点【请保持时间范围最大90天】
	CreatTimeRangeStart *string `json:"CreatTimeRangeStart,omitempty" name:"CreatTimeRangeStart"`

	// 下单时间范围终止点【请保持时间范围最大90天】
	CreatTimeRangeEnd *string `json:"CreatTimeRangeEnd,omitempty" name:"CreatTimeRangeEnd"`

	// 0:下单时间降序；其他：下单时间升序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 订单的状态(1：未支付;2：已支付;3：发货中;4：已发货;5：发货失败;6：已退款;7：已关单;8：订单过期;9：订单已失效;10：产品已失效;11：代付拒绝;12：支付中)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 下单人账号ID列表
	OwnerUins []*string `json:"OwnerUins,omitempty" name:"OwnerUins"`

	// 订单号列表
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 支付方式，0：自付；1：代付
	PayerMode *uint64 `json:"PayerMode,omitempty" name:"PayerMode"`
}

func (r *DescribeAgentDealsCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentDealsCacheRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CreatTimeRangeStart")
	delete(f, "CreatTimeRangeEnd")
	delete(f, "Order")
	delete(f, "Status")
	delete(f, "OwnerUins")
	delete(f, "DealNames")
	delete(f, "PayerMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentDealsCacheRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentDealsCacheResponseParams struct {
	// 订单数组
	AgentDealSet []*AgentDealElem `json:"AgentDealSet,omitempty" name:"AgentDealSet"`

	// 符合条件的订单总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAgentDealsCacheResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentDealsCacheResponseParams `json:"Response"`
}

func (r *DescribeAgentDealsCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentDealsCacheResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentPayDealsRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目 最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下单时间范围起始点(不传时会默认查15天内订单，传值时需要传15天内的起始时间)
	CreatTimeRangeStart *string `json:"CreatTimeRangeStart,omitempty" name:"CreatTimeRangeStart"`

	// 下单时间范围终止点
	CreatTimeRangeEnd *string `json:"CreatTimeRangeEnd,omitempty" name:"CreatTimeRangeEnd"`

	// 0:下单时间降序；其他：下单时间升序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 订单的状态(1：未支付;2：已支付;3：发货中;4：已发货;5：发货失败;6：已退款;7：已关单;8：订单过期;9：订单已失效;10：产品已失效;11：代付拒绝;12：支付中)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 下单人账号ID列表
	OwnerUins []*string `json:"OwnerUins,omitempty" name:"OwnerUins"`

	// 订单号列表
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
}

type DescribeAgentPayDealsRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目 最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下单时间范围起始点(不传时会默认查15天内订单，传值时需要传15天内的起始时间)
	CreatTimeRangeStart *string `json:"CreatTimeRangeStart,omitempty" name:"CreatTimeRangeStart"`

	// 下单时间范围终止点
	CreatTimeRangeEnd *string `json:"CreatTimeRangeEnd,omitempty" name:"CreatTimeRangeEnd"`

	// 0:下单时间降序；其他：下单时间升序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 订单的状态(1：未支付;2：已支付;3：发货中;4：已发货;5：发货失败;6：已退款;7：已关单;8：订单过期;9：订单已失效;10：产品已失效;11：代付拒绝;12：支付中)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 下单人账号ID列表
	OwnerUins []*string `json:"OwnerUins,omitempty" name:"OwnerUins"`

	// 订单号列表
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
}

func (r *DescribeAgentPayDealsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentPayDealsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CreatTimeRangeStart")
	delete(f, "CreatTimeRangeEnd")
	delete(f, "Order")
	delete(f, "Status")
	delete(f, "OwnerUins")
	delete(f, "DealNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentPayDealsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentPayDealsResponseParams struct {
	// 订单数组
	AgentPayDealSet []*AgentDealElem `json:"AgentPayDealSet,omitempty" name:"AgentPayDealSet"`

	// 符合条件的订单总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAgentPayDealsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentPayDealsResponseParams `json:"Response"`
}

func (r *DescribeAgentPayDealsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentPayDealsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentPayDealsV2RequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目 最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下单时间范围起始点(不传时会默认查15天内订单，传值时需要传15天内的起始时间)
	CreatTimeRangeStart *string `json:"CreatTimeRangeStart,omitempty" name:"CreatTimeRangeStart"`

	// 下单时间范围终止点
	CreatTimeRangeEnd *string `json:"CreatTimeRangeEnd,omitempty" name:"CreatTimeRangeEnd"`

	// 0:下单时间降序；其他：下单时间升序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 订单的状态(1：未支付;2：已支付;3：发货中;4：已发货;5：发货失败;6：已退款;7：已关单;8：订单过期;9：订单已失效;10：产品已失效;11：代付拒绝;12：支付中)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 下单人账号ID列表
	OwnerUins []*string `json:"OwnerUins,omitempty" name:"OwnerUins"`

	// 子订单号列表
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 大订单号列表
	BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds"`
}

type DescribeAgentPayDealsV2Request struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目 最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下单时间范围起始点(不传时会默认查15天内订单，传值时需要传15天内的起始时间)
	CreatTimeRangeStart *string `json:"CreatTimeRangeStart,omitempty" name:"CreatTimeRangeStart"`

	// 下单时间范围终止点
	CreatTimeRangeEnd *string `json:"CreatTimeRangeEnd,omitempty" name:"CreatTimeRangeEnd"`

	// 0:下单时间降序；其他：下单时间升序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 订单的状态(1：未支付;2：已支付;3：发货中;4：已发货;5：发货失败;6：已退款;7：已关单;8：订单过期;9：订单已失效;10：产品已失效;11：代付拒绝;12：支付中)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 下单人账号ID列表
	OwnerUins []*string `json:"OwnerUins,omitempty" name:"OwnerUins"`

	// 子订单号列表
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 大订单号列表
	BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds"`
}

func (r *DescribeAgentPayDealsV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentPayDealsV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CreatTimeRangeStart")
	delete(f, "CreatTimeRangeEnd")
	delete(f, "Order")
	delete(f, "Status")
	delete(f, "OwnerUins")
	delete(f, "DealNames")
	delete(f, "BigDealIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentPayDealsV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentPayDealsV2ResponseParams struct {
	// 订单数组
	AgentPayDealSet []*AgentDealNewElem `json:"AgentPayDealSet,omitempty" name:"AgentPayDealSet"`

	// 符合条件的订单总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAgentPayDealsV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeAgentPayDealsV2ResponseParams `json:"Response"`
}

func (r *DescribeAgentPayDealsV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentPayDealsV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentSelfPayDealsRequestParams struct {
	// 下单人账号ID
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目 最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下单时间范围起始点(不传时会默认查15天内订单，传值时需要传15天内的起始时间)
	CreatTimeRangeStart *string `json:"CreatTimeRangeStart,omitempty" name:"CreatTimeRangeStart"`

	// 下单时间范围终止点
	CreatTimeRangeEnd *string `json:"CreatTimeRangeEnd,omitempty" name:"CreatTimeRangeEnd"`

	// 0:下单时间降序；其他：下单时间升序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 订单的状态(1：未支付;2：已支付;3：发货中;4：已发货;5：发货失败;6：已退款;7：已关单;8：订单过期;9：订单已失效;10：产品已失效;11：代付拒绝;12：支付中)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 订单号列表
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
}

type DescribeAgentSelfPayDealsRequest struct {
	*tchttp.BaseRequest
	
	// 下单人账号ID
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目 最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下单时间范围起始点(不传时会默认查15天内订单，传值时需要传15天内的起始时间)
	CreatTimeRangeStart *string `json:"CreatTimeRangeStart,omitempty" name:"CreatTimeRangeStart"`

	// 下单时间范围终止点
	CreatTimeRangeEnd *string `json:"CreatTimeRangeEnd,omitempty" name:"CreatTimeRangeEnd"`

	// 0:下单时间降序；其他：下单时间升序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 订单的状态(1：未支付;2：已支付;3：发货中;4：已发货;5：发货失败;6：已退款;7：已关单;8：订单过期;9：订单已失效;10：产品已失效;11：代付拒绝;12：支付中)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 订单号列表
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
}

func (r *DescribeAgentSelfPayDealsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentSelfPayDealsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OwnerUin")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CreatTimeRangeStart")
	delete(f, "CreatTimeRangeEnd")
	delete(f, "Order")
	delete(f, "Status")
	delete(f, "DealNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentSelfPayDealsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentSelfPayDealsResponseParams struct {
	// 订单数组
	AgentPayDealSet []*AgentDealElem `json:"AgentPayDealSet,omitempty" name:"AgentPayDealSet"`

	// 符合条件的订单总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAgentSelfPayDealsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentSelfPayDealsResponseParams `json:"Response"`
}

func (r *DescribeAgentSelfPayDealsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentSelfPayDealsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentSelfPayDealsV2RequestParams struct {
	// 下单人账号ID
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目 最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下单时间范围起始点(不传时会默认查15天内订单，传值时需要传15天内的起始时间)
	CreatTimeRangeStart *string `json:"CreatTimeRangeStart,omitempty" name:"CreatTimeRangeStart"`

	// 下单时间范围终止点
	CreatTimeRangeEnd *string `json:"CreatTimeRangeEnd,omitempty" name:"CreatTimeRangeEnd"`

	// 0:下单时间降序；其他：下单时间升序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 订单的状态(1：未支付;2：已支付;3：发货中;4：已发货;5：发货失败;6：已退款;7：已关单;8：订单过期;9：订单已失效;10：产品已失效;11：代付拒绝;12：支付中)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 子订单号列表
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 大订单号列表
	BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds"`
}

type DescribeAgentSelfPayDealsV2Request struct {
	*tchttp.BaseRequest
	
	// 下单人账号ID
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目 最大100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 下单时间范围起始点(不传时会默认查15天内订单，传值时需要传15天内的起始时间)
	CreatTimeRangeStart *string `json:"CreatTimeRangeStart,omitempty" name:"CreatTimeRangeStart"`

	// 下单时间范围终止点
	CreatTimeRangeEnd *string `json:"CreatTimeRangeEnd,omitempty" name:"CreatTimeRangeEnd"`

	// 0:下单时间降序；其他：下单时间升序
	Order *uint64 `json:"Order,omitempty" name:"Order"`

	// 订单的状态(1：未支付;2：已支付;3：发货中;4：已发货;5：发货失败;6：已退款;7：已关单;8：订单过期;9：订单已失效;10：产品已失效;11：代付拒绝;12：支付中)
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 子订单号列表
	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`

	// 大订单号列表
	BigDealIds []*string `json:"BigDealIds,omitempty" name:"BigDealIds"`
}

func (r *DescribeAgentSelfPayDealsV2Request) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentSelfPayDealsV2Request) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OwnerUin")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "CreatTimeRangeStart")
	delete(f, "CreatTimeRangeEnd")
	delete(f, "Order")
	delete(f, "Status")
	delete(f, "DealNames")
	delete(f, "BigDealIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentSelfPayDealsV2Request has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentSelfPayDealsV2ResponseParams struct {
	// 订单数组
	AgentPayDealSet []*AgentDealNewElem `json:"AgentPayDealSet,omitempty" name:"AgentPayDealSet"`

	// 符合条件的订单总数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAgentSelfPayDealsV2Response struct {
	*tchttp.BaseResponse
	Response *DescribeAgentSelfPayDealsV2ResponseParams `json:"Response"`
}

func (r *DescribeAgentSelfPayDealsV2Response) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentSelfPayDealsV2Response) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientBalanceNewRequestParams struct {
	// 客户(代客)账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`
}

type DescribeClientBalanceNewRequest struct {
	*tchttp.BaseRequest
	
	// 客户(代客)账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`
}

func (r *DescribeClientBalanceNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientBalanceNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClientBalanceNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientBalanceNewResponseParams struct {
	// 账户可用余额，单位分 （可用余额 = 现金余额 + 赠送金余额 - 欠费金额 - 冻结金额）
	Balance *int64 `json:"Balance,omitempty" name:"Balance"`

	// 账户现金余额，单位分
	Cash *int64 `json:"Cash,omitempty" name:"Cash"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClientBalanceNewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClientBalanceNewResponseParams `json:"Response"`
}

func (r *DescribeClientBalanceNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientBalanceNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientBalanceRequestParams struct {
	// 客户(代客)账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`
}

type DescribeClientBalanceRequest struct {
	*tchttp.BaseRequest
	
	// 客户(代客)账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`
}

func (r *DescribeClientBalanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientBalanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClientBalanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientBalanceResponseParams struct {
	// 账户可用余额，单位分 （可用余额 = 现金余额 - 冻结金额）  【注：该数据准确性存疑，请切换至DescribeClientBalanceNew取值】
	Balance *uint64 `json:"Balance,omitempty" name:"Balance"`

	// 账户现金余额，单位分
	Cash *int64 `json:"Cash,omitempty" name:"Cash"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClientBalanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClientBalanceResponseParams `json:"Response"`
}

func (r *DescribeClientBalanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientBalanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRebateInfosNewRequestParams struct {
	// 返佣月份，如2018-02
	RebateMonth *string `json:"RebateMonth,omitempty" name:"RebateMonth"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeRebateInfosNewRequest struct {
	*tchttp.BaseRequest
	
	// 返佣月份，如2018-02
	RebateMonth *string `json:"RebateMonth,omitempty" name:"RebateMonth"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRebateInfosNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRebateInfosNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RebateMonth")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRebateInfosNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRebateInfosNewResponseParams struct {
	// 返佣信息列表
	RebateInfoSet []*RebateInfoElemNew `json:"RebateInfoSet,omitempty" name:"RebateInfoSet"`

	// 符合查询条件返佣信息数目
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRebateInfosNewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRebateInfosNewResponseParams `json:"Response"`
}

func (r *DescribeRebateInfosNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRebateInfosNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRebateInfosRequestParams struct {
	// 返佣月份，如2018-02
	RebateMonth *string `json:"RebateMonth,omitempty" name:"RebateMonth"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeRebateInfosRequest struct {
	*tchttp.BaseRequest
	
	// 返佣月份，如2018-02
	RebateMonth *string `json:"RebateMonth,omitempty" name:"RebateMonth"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRebateInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRebateInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RebateMonth")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRebateInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRebateInfosResponseParams struct {
	// 返佣信息列表
	RebateInfoSet []*RebateInfoElem `json:"RebateInfoSet,omitempty" name:"RebateInfoSet"`

	// 符合查询条件返佣信息数目
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRebateInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRebateInfosResponseParams `json:"Response"`
}

func (r *DescribeRebateInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRebateInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSalesmansRequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 业务员姓名(模糊查询)
	SalesName *string `json:"SalesName,omitempty" name:"SalesName"`

	// 业务员ID
	SalesUin *string `json:"SalesUin,omitempty" name:"SalesUin"`

	// ASC/DESC， 不区分大小写，按创建通过时间排序
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeSalesmansRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 业务员姓名(模糊查询)
	SalesName *string `json:"SalesName,omitempty" name:"SalesName"`

	// 业务员ID
	SalesUin *string `json:"SalesUin,omitempty" name:"SalesUin"`

	// ASC/DESC， 不区分大小写，按创建通过时间排序
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeSalesmansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSalesmansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SalesName")
	delete(f, "SalesUin")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSalesmansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSalesmansResponseParams struct {
	// 业务员列表
	AgentSalesmanSet []*AgentSalesmanElem `json:"AgentSalesmanSet,omitempty" name:"AgentSalesmanSet"`

	// 符合条件的代客总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSalesmansResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSalesmansResponseParams `json:"Response"`
}

func (r *DescribeSalesmansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSalesmansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnbindClientListRequestParams struct {
	// 解绑状态：0:所有,1:审核中,2已解绑
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 解绑账号ID
	UnbindUin *string `json:"UnbindUin,omitempty" name:"UnbindUin"`

	// 解绑申请时间范围起始点
	ApplyTimeStart *string `json:"ApplyTimeStart,omitempty" name:"ApplyTimeStart"`

	// 解绑申请时间范围终止点
	ApplyTimeEnd *string `json:"ApplyTimeEnd,omitempty" name:"ApplyTimeEnd"`

	// 对申请时间的升序降序，值：asc，desc
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

type DescribeUnbindClientListRequest struct {
	*tchttp.BaseRequest
	
	// 解绑状态：0:所有,1:审核中,2已解绑
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 解绑账号ID
	UnbindUin *string `json:"UnbindUin,omitempty" name:"UnbindUin"`

	// 解绑申请时间范围起始点
	ApplyTimeStart *string `json:"ApplyTimeStart,omitempty" name:"ApplyTimeStart"`

	// 解绑申请时间范围终止点
	ApplyTimeEnd *string `json:"ApplyTimeEnd,omitempty" name:"ApplyTimeEnd"`

	// 对申请时间的升序降序，值：asc，desc
	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeUnbindClientListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnbindClientListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Status")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "UnbindUin")
	delete(f, "ApplyTimeStart")
	delete(f, "ApplyTimeEnd")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUnbindClientListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnbindClientListResponseParams struct {
	// 符合条件的解绑客户数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 符合条件的解绑客户列表
	UnbindClientList []*UnbindClientElem `json:"UnbindClientList,omitempty" name:"UnbindClientList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUnbindClientListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUnbindClientListResponseParams `json:"Response"`
}

func (r *DescribeUnbindClientListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnbindClientListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClientRemarkRequestParams struct {
	// 客户备注名称
	ClientRemark *string `json:"ClientRemark,omitempty" name:"ClientRemark"`

	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`
}

type ModifyClientRemarkRequest struct {
	*tchttp.BaseRequest
	
	// 客户备注名称
	ClientRemark *string `json:"ClientRemark,omitempty" name:"ClientRemark"`

	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`
}

func (r *ModifyClientRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClientRemarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientRemark")
	delete(f, "ClientUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClientRemarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClientRemarkResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyClientRemarkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClientRemarkResponseParams `json:"Response"`
}

func (r *ModifyClientRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClientRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductInfoElem struct {
	// 产品属性
	Name *string `json:"Name,omitempty" name:"Name"`

	// 产品属性值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type RebateInfoElem struct {
	// 代理商账号ID
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 返佣月份，如2018-02
	RebateMonth *string `json:"RebateMonth,omitempty" name:"RebateMonth"`

	// 返佣金额，单位分
	Amt *uint64 `json:"Amt,omitempty" name:"Amt"`

	// 月度业绩，单位分
	MonthSales *uint64 `json:"MonthSales,omitempty" name:"MonthSales"`

	// 季度业绩，单位分
	QuarterSales *uint64 `json:"QuarterSales,omitempty" name:"QuarterSales"`

	// NORMAL(正常)/HAS_OVERDUE_BILL(欠费)/NO_CONTRACT(缺合同)
	ExceptionFlag *string `json:"ExceptionFlag,omitempty" name:"ExceptionFlag"`
}

type RebateInfoElemNew struct {
	// 代理商账号ID
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 返佣月份，如2018-02
	RebateMonth *string `json:"RebateMonth,omitempty" name:"RebateMonth"`

	// 返佣金额，单位分
	Amt *int64 `json:"Amt,omitempty" name:"Amt"`

	// 月度业绩，单位分
	MonthSales *int64 `json:"MonthSales,omitempty" name:"MonthSales"`

	// 季度业绩，单位分
	QuarterSales *int64 `json:"QuarterSales,omitempty" name:"QuarterSales"`

	// NORMAL(正常)/HAS_OVERDUE_BILL(欠费)/NO_CONTRACT(缺合同)
	ExceptionFlag *string `json:"ExceptionFlag,omitempty" name:"ExceptionFlag"`
}

// Predefined struct for user
type RemovePayRelationForClientRequestParams struct {
	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`
}

type RemovePayRelationForClientRequest struct {
	*tchttp.BaseRequest
	
	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitempty" name:"ClientUin"`
}

func (r *RemovePayRelationForClientRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemovePayRelationForClientRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemovePayRelationForClientRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemovePayRelationForClientResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RemovePayRelationForClientResponse struct {
	*tchttp.BaseResponse
	Response *RemovePayRelationForClientResponseParams `json:"Response"`
}

func (r *RemovePayRelationForClientResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemovePayRelationForClientResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindClientElem struct {
	// 解绑账号ID
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 状态：0:审核中；1：已解绑；2：已撤销 3：关联撤销 4: 已驳回
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 申请时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplyTime *string `json:"ApplyTime,omitempty" name:"ApplyTime"`

	// 解绑/撤销时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionTime *string `json:"ActionTime,omitempty" name:"ActionTime"`
}