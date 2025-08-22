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

package v20180321

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AgentAuditedClient struct {
	// 代理商账号ID
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 代客账号ID
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// 代客审核通过时间戳
	AgentTime *string `json:"AgentTime,omitnil,omitempty" name:"AgentTime"`

	// 代客类型，可能值为a/b/c
	ClientFlag *string `json:"ClientFlag,omitnil,omitempty" name:"ClientFlag"`

	// 代客备注
	ClientRemark *string `json:"ClientRemark,omitnil,omitempty" name:"ClientRemark"`

	// 代客名称（首选实名认证名称）
	ClientName *string `json:"ClientName,omitnil,omitempty" name:"ClientName"`

	// 认证类型, 0：个人，1：企业；其他：未认证或无效值
	AuthType *string `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// 代客APPID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 上月消费金额
	LastMonthAmt *int64 `json:"LastMonthAmt,omitnil,omitempty" name:"LastMonthAmt"`

	// 本月消费金额
	ThisMonthAmt *int64 `json:"ThisMonthAmt,omitnil,omitempty" name:"ThisMonthAmt"`

	// 是否欠费,0：不欠费；1：欠费
	HasOverdueBill *uint64 `json:"HasOverdueBill,omitnil,omitempty" name:"HasOverdueBill"`

	// 客户类型：可以为new(自拓)/assign(指派)/old(官网)/direct(直销)/direct_newopp(直销(新商机))/空
	ClientType *string `json:"ClientType,omitnil,omitempty" name:"ClientType"`

	// 项目类型：可以为self(自拓项目)/platform(合作项目)/repeat(复算项目  )/空
	ProjectType *string `json:"ProjectType,omitnil,omitempty" name:"ProjectType"`

	// 业务员ID
	SalesUin *string `json:"SalesUin,omitnil,omitempty" name:"SalesUin"`

	// 业务员姓名
	SalesName *string `json:"SalesName,omitnil,omitempty" name:"SalesName"`

	// 代客邮箱
	Mail *string `json:"Mail,omitnil,omitempty" name:"Mail"`

	// 交易类型:交易类型 1-原类型 2-代理型  3-代采型
	TransactionType *string `json:"TransactionType,omitnil,omitempty" name:"TransactionType"`
}

type AgentBillElem struct {
	// 代理商账号ID
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 订单号，仅对预付费账单有意义
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 代客账号ID
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// 代客备注名称
	ClientRemark *string `json:"ClientRemark,omitnil,omitempty" name:"ClientRemark"`

	// 支付时间
	PayTime *string `json:"PayTime,omitnil,omitempty" name:"PayTime"`

	// 云产品名称
	GoodsType *string `json:"GoodsType,omitnil,omitempty" name:"GoodsType"`

	// 预付费/后付费
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 支付月份
	SettleMonth *string `json:"SettleMonth,omitnil,omitempty" name:"SettleMonth"`

	// 支付金额，单位分
	Amt *int64 `json:"Amt,omitnil,omitempty" name:"Amt"`

	// agentpay：代付；selfpay：自付
	PayerMode *string `json:"PayerMode,omitnil,omitempty" name:"PayerMode"`

	// 客户类型：可以为new(自拓)/assign(指定)/old(官网)/direct(直销)/direct_newopp(直销(新商机))/空
	ClientType *string `json:"ClientType,omitnil,omitempty" name:"ClientType"`

	// 项目类型：可以为self(自拓项目)/platform(合作项目)/repeat(复算项目  )/空
	ProjectType *string `json:"ProjectType,omitnil,omitempty" name:"ProjectType"`

	// 活动ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`
}

type AgentClientElem struct {
	// 代理商账号ID
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 代客账号ID
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// 代客申请时间戳
	ApplyTime *uint64 `json:"ApplyTime,omitnil,omitempty" name:"ApplyTime"`

	// 代客类型，可能值为a/b/c/other
	ClientFlag *string `json:"ClientFlag,omitnil,omitempty" name:"ClientFlag"`

	// 代客邮箱，打码显示
	Mail *string `json:"Mail,omitnil,omitempty" name:"Mail"`

	// 代客手机，打码显示
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 0表示不欠费，1表示欠费
	HasOverdueBill *uint64 `json:"HasOverdueBill,omitnil,omitempty" name:"HasOverdueBill"`

	// 1:待代理商审核;2:待腾讯云审核4:待腾讯云渠道审批
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 业务员ID
	SalesUin *string `json:"SalesUin,omitnil,omitempty" name:"SalesUin"`

	// 业务员姓名
	SalesName *string `json:"SalesName,omitnil,omitempty" name:"SalesName"`

	// 客户名称，此字段和控制台返回一致。
	ClientName *string `json:"ClientName,omitnil,omitempty" name:"ClientName"`

	// 增量目标金额(分)
	IncreaseGoal *string `json:"IncreaseGoal,omitnil,omitempty" name:"IncreaseGoal"`
}

type AgentDealNewElem struct {
	// 订单自增 ID【请勿依赖该字段作为唯一标识】
	DealId *string `json:"DealId,omitnil,omitempty" name:"DealId"`

	// 订单号【订单唯一键】
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 商品类型 ID
	GoodsCategoryId *string `json:"GoodsCategoryId,omitnil,omitempty" name:"GoodsCategoryId"`

	// 订单所有者
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 订单所有者对应 appId
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 商品数量
	GoodsNum *string `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 价格详情
	GoodsPrice *DealGoodsPriceNewElem `json:"GoodsPrice,omitnil,omitempty" name:"GoodsPrice"`

	// 下单人
	Creater *string `json:"Creater,omitnil,omitempty" name:"Creater"`

	// 下单时间
	CreatTime *string `json:"CreatTime,omitnil,omitempty" name:"CreatTime"`

	// 支付结束时间
	PayEndTime *string `json:"PayEndTime,omitnil,omitempty" name:"PayEndTime"`

	// 扣费流水号
	BillId *string `json:"BillId,omitnil,omitempty" name:"BillId"`

	// 支付人
	Payer *string `json:"Payer,omitnil,omitempty" name:"Payer"`

	// 订单状态，中文描述
	DealStatus *string `json:"DealStatus,omitnil,omitempty" name:"DealStatus"`

	// 订单的状态(1：未支付;2：已支付;3：发货中;4：已发货;5：发货失败;6：已退款;7：已关单;8：订单过期;9：订单已失效;10：产品已失效;11：代付拒绝;12：支付中)
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 产品名称
	GoodsName *string `json:"GoodsName,omitnil,omitempty" name:"GoodsName"`

	// 客户备注
	ClientRemark *string `json:"ClientRemark,omitnil,omitempty" name:"ClientRemark"`

	// 订单操作类型，"purchase":"新购","upgrade":"升配","upConvertExpire":"升配","downgrade":"降配","downConvertExpire":"降配","renew":"续费","refund":"退货","modifyNetworkMode":"调整带宽模式","modifyNetworkSize":"调整带宽大小","preMoveOut":"资源迁出","preMoveIn":"资源迁入","preToPost":"包年包月转按量","modify":"变配","postMoveOut":"资源迁出","postMoveIn":"资源迁入","recoverRefundForward":"调账补偿","recoverPayReserve":"调账补偿","recoverPayForward":"调账扣费","recoverRefundReserve":"调账扣费"
	ActionType *string `json:"ActionType,omitnil,omitempty" name:"ActionType"`

	// 代金券抵扣金额，单位分
	VoucherDecline *string `json:"VoucherDecline,omitnil,omitempty" name:"VoucherDecline"`

	// 大订单号
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`

	// 客户类型（new：自拓；old：官网；assign：指派；direct：直销；direct_newopp：直销(新商机)）
	ClientType *string `json:"ClientType,omitnil,omitempty" name:"ClientType"`

	// 项目类型（self：自拓；repeat：直销；platform：官网合作）
	ProjectType *string `json:"ProjectType,omitnil,omitempty" name:"ProjectType"`

	// 业务员账号ID
	SalesUin *string `json:"SalesUin,omitnil,omitempty" name:"SalesUin"`

	// 支付方式，0：自付；1：代付
	PayerMode *string `json:"PayerMode,omitnil,omitempty" name:"PayerMode"`

	// 活动ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 订单过期时间
	OverdueTime *string `json:"OverdueTime,omitnil,omitempty" name:"OverdueTime"`

	// 产品详情
	ProductInfo []*ProductInfoElem `json:"ProductInfo,omitnil,omitempty" name:"ProductInfo"`

	// 付款方式
	PaymentMethod *string `json:"PaymentMethod,omitnil,omitempty" name:"PaymentMethod"`

	// 订单更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 资源id
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 退款单的原订单信息。当前仅 DescribeClientDealsByCache 接口会返回该字段
	RefundMap []*RefundMap `json:"RefundMap,omitnil,omitempty" name:"RefundMap"`

	// 子产品名称
	SubGoodsName *string `json:"SubGoodsName,omitnil,omitempty" name:"SubGoodsName"`
}

// Predefined struct for user
type AgentPayDealsRequestParams struct {
	// 订单所有者uin
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 代付标志，1：代付；0：自付
	AgentPay *uint64 `json:"AgentPay,omitnil,omitempty" name:"AgentPay"`

	// 订单号数组
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
}

type AgentPayDealsRequest struct {
	*tchttp.BaseRequest
	
	// 订单所有者uin
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 代付标志，1：代付；0：自付
	AgentPay *uint64 `json:"AgentPay,omitnil,omitempty" name:"AgentPay"`

	// 订单号数组
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 业务员ID
	SalesUin *string `json:"SalesUin,omitnil,omitempty" name:"SalesUin"`

	// 业务员姓名
	SalesName *string `json:"SalesName,omitnil,omitempty" name:"SalesName"`

	// 业务员创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type AgentTransferMoneyRequestParams struct {
	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// 转账金额，单位分
	Amount *uint64 `json:"Amount,omitnil,omitempty" name:"Amount"`
}

type AgentTransferMoneyRequest struct {
	*tchttp.BaseRequest
	
	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// 转账金额，单位分
	Amount *uint64 `json:"Amount,omitnil,omitempty" name:"Amount"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClientUins []*string `json:"ClientUins,omitnil,omitempty" name:"ClientUins"`

	// 业务员uin
	SalesUin *string `json:"SalesUin,omitnil,omitempty" name:"SalesUin"`

	// 代客类型:normal-代客 apply-申请中代客
	AssignClientStatus *string `json:"AssignClientStatus,omitnil,omitempty" name:"AssignClientStatus"`

	// 操作类型:assign-执行分派 cancel-取消分派
	AssignActionType *string `json:"AssignActionType,omitnil,omitempty" name:"AssignActionType"`
}

type AssignClientsToSalesRequest struct {
	*tchttp.BaseRequest
	
	// 代客/申请中代客uin列表，最大50条
	ClientUins []*string `json:"ClientUins,omitnil,omitempty" name:"ClientUins"`

	// 业务员uin
	SalesUin *string `json:"SalesUin,omitnil,omitempty" name:"SalesUin"`

	// 代客类型:normal-代客 apply-申请中代客
	AssignClientStatus *string `json:"AssignClientStatus,omitnil,omitempty" name:"AssignClientStatus"`

	// 操作类型:assign-执行分派 cancel-取消分派
	AssignActionType *string `json:"AssignActionType,omitnil,omitempty" name:"AssignActionType"`
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
	SucceedUins []*string `json:"SucceedUins,omitnil,omitempty" name:"SucceedUins"`

	// 处理失败的代客uin列表
	FailedUins []*string `json:"FailedUins,omitnil,omitempty" name:"FailedUins"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// 审核结果，可能的取值：accept/reject
	AuditResult *string `json:"AuditResult,omitnil,omitempty" name:"AuditResult"`

	// 申请理由，B类客户审核通过时必须填写申请理由
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`
}

type AuditApplyClientRequest struct {
	*tchttp.BaseRequest
	
	// 待审核客户账号ID
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// 审核结果，可能的取值：accept/reject
	AuditResult *string `json:"AuditResult,omitnil,omitempty" name:"AuditResult"`

	// 申请理由，B类客户审核通过时必须填写申请理由
	Note *string `json:"Note,omitnil,omitempty" name:"Note"`
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
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// 审核结果，包括accept/reject/qcloudaudit（腾讯云审核）
	AuditResult *string `json:"AuditResult,omitnil,omitempty" name:"AuditResult"`

	// 关联时间对应的时间戳
	AgentTime *uint64 `json:"AgentTime,omitnil,omitempty" name:"AgentTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type ClientIncreaseInfoList struct {
	// 客户UIN
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// 是否参与增量政策，
	// Y：是，N：否
	IsJoinIncrease *string `json:"IsJoinIncrease,omitnil,omitempty" name:"IsJoinIncrease"`

	// 增量考核关联时间
	IncreaseUseAssociateDate *string `json:"IncreaseUseAssociateDate,omitnil,omitempty" name:"IncreaseUseAssociateDate"`

	// 参与增量考核的原始客户等级
	TLevel *string `json:"TLevel,omitnil,omitempty" name:"TLevel"`

	// 增量考核目标,分
	IncreaseGoal *string `json:"IncreaseGoal,omitnil,omitempty" name:"IncreaseGoal"`

	// 完成订单金额,分
	TotalBaseAmt *string `json:"TotalBaseAmt,omitnil,omitempty" name:"TotalBaseAmt"`
}

// Predefined struct for user
type CreatePayRelationForClientRequestParams struct {
	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`
}

type CreatePayRelationForClientRequest struct {
	*tchttp.BaseRequest
	
	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type DealGoodsPriceNewElem struct {
	// 实付金额（单位：分）
	RealTotalCost *int64 `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`

	// 原始金额（不含折扣，单位：分）
	OriginalTotalCost *int64 `json:"OriginalTotalCost,omitnil,omitempty" name:"OriginalTotalCost"`
}

type DealPriceDetail struct {
	// 子订单号
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 订单归属人uin（代客uin）
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 子产品价格详情列表
	SubProductPriceDetail []*SubProductPriceDetail `json:"SubProductPriceDetail,omitnil,omitempty" name:"SubProductPriceDetail"`
}

// Predefined struct for user
type DescribeAgentAuditedClientsRequestParams struct {
	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// 客户名称。由于涉及隐私，名称打码显示，故名称仅支持打码后的模糊搜索
	ClientName *string `json:"ClientName,omitnil,omitempty" name:"ClientName"`

	// 客户类型，a/b，类型定义参考代理商相关政策文档
	ClientFlag *string `json:"ClientFlag,omitnil,omitempty" name:"ClientFlag"`

	// ASC/DESC， 不区分大小写，按审核通过时间排序
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// 客户账号ID列表
	ClientUins []*string `json:"ClientUins,omitnil,omitempty" name:"ClientUins"`

	// 是否欠费。0：不欠费；1：欠费
	HasOverdueBill *uint64 `json:"HasOverdueBill,omitnil,omitempty" name:"HasOverdueBill"`

	// 客户备注
	ClientRemark *string `json:"ClientRemark,omitnil,omitempty" name:"ClientRemark"`

	// *偏移量 【请保持必传】
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// *限制数目 【请保持必传】最大2000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 可以为new(自拓)/assign(指派)/old(官网)/direct(直销)/direct_newopp(直销(新商机))/空
	ClientType *string `json:"ClientType,omitnil,omitempty" name:"ClientType"`

	// 项目类型：可以为self(自拓项目)/platform(合作项目)/repeat(复算项目  )/空
	ProjectType *string `json:"ProjectType,omitnil,omitempty" name:"ProjectType"`

	// 业务员ID
	SalesUin *string `json:"SalesUin,omitnil,omitempty" name:"SalesUin"`

	// 业务员姓名（模糊查询）
	SalesName *string `json:"SalesName,omitnil,omitempty" name:"SalesName"`
}

type DescribeAgentAuditedClientsRequest struct {
	*tchttp.BaseRequest
	
	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// 客户名称。由于涉及隐私，名称打码显示，故名称仅支持打码后的模糊搜索
	ClientName *string `json:"ClientName,omitnil,omitempty" name:"ClientName"`

	// 客户类型，a/b，类型定义参考代理商相关政策文档
	ClientFlag *string `json:"ClientFlag,omitnil,omitempty" name:"ClientFlag"`

	// ASC/DESC， 不区分大小写，按审核通过时间排序
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// 客户账号ID列表
	ClientUins []*string `json:"ClientUins,omitnil,omitempty" name:"ClientUins"`

	// 是否欠费。0：不欠费；1：欠费
	HasOverdueBill *uint64 `json:"HasOverdueBill,omitnil,omitempty" name:"HasOverdueBill"`

	// 客户备注
	ClientRemark *string `json:"ClientRemark,omitnil,omitempty" name:"ClientRemark"`

	// *偏移量 【请保持必传】
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// *限制数目 【请保持必传】最大2000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 可以为new(自拓)/assign(指派)/old(官网)/direct(直销)/direct_newopp(直销(新商机))/空
	ClientType *string `json:"ClientType,omitnil,omitempty" name:"ClientType"`

	// 项目类型：可以为self(自拓项目)/platform(合作项目)/repeat(复算项目  )/空
	ProjectType *string `json:"ProjectType,omitnil,omitempty" name:"ProjectType"`

	// 业务员ID
	SalesUin *string `json:"SalesUin,omitnil,omitempty" name:"SalesUin"`

	// 业务员姓名（模糊查询）
	SalesName *string `json:"SalesName,omitnil,omitempty" name:"SalesName"`
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
	AgentClientSet []*AgentAuditedClient `json:"AgentClientSet,omitnil,omitempty" name:"AgentClientSet"`

	// 符合条件的代客总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	SettleMonth *string `json:"SettleMonth,omitnil,omitempty" name:"SettleMonth"`

	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// 支付方式，prepay/postpay
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 预付费订单号
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 客户备注名称
	ClientRemark *string `json:"ClientRemark,omitnil,omitempty" name:"ClientRemark"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAgentBillsRequest struct {
	*tchttp.BaseRequest
	
	// 支付月份，如2018-02
	SettleMonth *string `json:"SettleMonth,omitnil,omitempty" name:"SettleMonth"`

	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// 支付方式，prepay/postpay
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 预付费订单号
	OrderId *string `json:"OrderId,omitnil,omitempty" name:"OrderId"`

	// 客户备注名称
	ClientRemark *string `json:"ClientRemark,omitnil,omitempty" name:"ClientRemark"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 业务明细列表
	AgentBillSet []*AgentBillElem `json:"AgentBillSet,omitnil,omitempty" name:"AgentBillSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`
}

type DescribeAgentClientGradeRequest struct {
	*tchttp.BaseRequest
	
	// 代客uin
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`
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
	AuditStatus *uint64 `json:"AuditStatus,omitnil,omitempty" name:"AuditStatus"`

	// 实名认证状态：0，未实名认证，1实名认证
	AuthState *uint64 `json:"AuthState,omitnil,omitempty" name:"AuthState"`

	// 客户级别
	ClientGrade *string `json:"ClientGrade,omitnil,omitempty" name:"ClientGrade"`

	// 客户类型：1，个人；2，企业；3，其他
	ClientType *uint64 `json:"ClientType,omitnil,omitempty" name:"ClientType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// 客户名称。由于涉及隐私，名称打码显示，故名称仅支持打码后的模糊搜索
	ClientName *string `json:"ClientName,omitnil,omitempty" name:"ClientName"`

	// 客户类型，a/b，类型定义参考代理商相关政策文档
	ClientFlag *string `json:"ClientFlag,omitnil,omitempty" name:"ClientFlag"`

	// ASC/DESC， 不区分大小写，按申请时间排序
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 业务员ID
	SalesUin *string `json:"SalesUin,omitnil,omitempty" name:"SalesUin"`

	// 业务员姓名（模糊查询）
	SalesName *string `json:"SalesName,omitnil,omitempty" name:"SalesName"`
}

type DescribeAgentClientsRequest struct {
	*tchttp.BaseRequest
	
	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// 客户名称。由于涉及隐私，名称打码显示，故名称仅支持打码后的模糊搜索
	ClientName *string `json:"ClientName,omitnil,omitempty" name:"ClientName"`

	// 客户类型，a/b，类型定义参考代理商相关政策文档
	ClientFlag *string `json:"ClientFlag,omitnil,omitempty" name:"ClientFlag"`

	// ASC/DESC， 不区分大小写，按申请时间排序
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 业务员ID
	SalesUin *string `json:"SalesUin,omitnil,omitempty" name:"SalesUin"`

	// 业务员姓名（模糊查询）
	SalesName *string `json:"SalesName,omitnil,omitempty" name:"SalesName"`
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
	AgentClientSet []*AgentClientElem `json:"AgentClientSet,omitnil,omitempty" name:"AgentClientSet"`

	// 符合条件的代客总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目 最大200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 下单时间范围起始点【*请必传并控制时间范围最大90天，避免出现超时】
	CreatTimeRangeStart *string `json:"CreatTimeRangeStart,omitnil,omitempty" name:"CreatTimeRangeStart"`

	// 下单时间范围终止点【*请必传并控制时间范围最大90天，避免出现超时】
	CreatTimeRangeEnd *string `json:"CreatTimeRangeEnd,omitnil,omitempty" name:"CreatTimeRangeEnd"`

	// 0:下单时间降序；其他：下单时间升序
	Order *uint64 `json:"Order,omitnil,omitempty" name:"Order"`

	// 子订单状态(1-待支付,2-已支付,3-发货中,4-已发货,5-发货失败,6-已退款,7-已取消,8-已过期,9-已失效,12-支付中,13-退款中,30-处理中)
	// 
	// 控制台订单状态为以上状态的组合：未支付(1) 处理中(2,3,5,12,13,30) 已取消(7) 交易成功(4) 已过期(8) 已退款(6) 订单错误(9)
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 下单人账号ID列表
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// 子订单号列表
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 大订单号列表
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// 支付方式，0：自付；1：代付
	PayerMode *uint64 `json:"PayerMode,omitnil,omitempty" name:"PayerMode"`
}

type DescribeAgentDealsByCacheRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目 最大200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 下单时间范围起始点【*请必传并控制时间范围最大90天，避免出现超时】
	CreatTimeRangeStart *string `json:"CreatTimeRangeStart,omitnil,omitempty" name:"CreatTimeRangeStart"`

	// 下单时间范围终止点【*请必传并控制时间范围最大90天，避免出现超时】
	CreatTimeRangeEnd *string `json:"CreatTimeRangeEnd,omitnil,omitempty" name:"CreatTimeRangeEnd"`

	// 0:下单时间降序；其他：下单时间升序
	Order *uint64 `json:"Order,omitnil,omitempty" name:"Order"`

	// 子订单状态(1-待支付,2-已支付,3-发货中,4-已发货,5-发货失败,6-已退款,7-已取消,8-已过期,9-已失效,12-支付中,13-退款中,30-处理中)
	// 
	// 控制台订单状态为以上状态的组合：未支付(1) 处理中(2,3,5,12,13,30) 已取消(7) 交易成功(4) 已过期(8) 已退款(6) 订单错误(9)
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 下单人账号ID列表
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// 子订单号列表
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 大订单号列表
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`

	// 支付方式，0：自付；1：代付
	PayerMode *uint64 `json:"PayerMode,omitnil,omitempty" name:"PayerMode"`
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
	AgentDealSet []*AgentDealNewElem `json:"AgentDealSet,omitnil,omitempty" name:"AgentDealSet"`

	// 符合条件的订单总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeAgentDealsPriceDetailByDealNameRequestParams struct {
	// 下单年份（订单创建时间归属年份）
	DealCreatYear *uint64 `json:"DealCreatYear,omitnil,omitempty" name:"DealCreatYear"`

	// 子订单号，每个请求最多查询100条
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 订单归属代客uin
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

type DescribeAgentDealsPriceDetailByDealNameRequest struct {
	*tchttp.BaseRequest
	
	// 下单年份（订单创建时间归属年份）
	DealCreatYear *uint64 `json:"DealCreatYear,omitnil,omitempty" name:"DealCreatYear"`

	// 子订单号，每个请求最多查询100条
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 订单归属代客uin
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

func (r *DescribeAgentDealsPriceDetailByDealNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentDealsPriceDetailByDealNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealCreatYear")
	delete(f, "DealNames")
	delete(f, "OwnerUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentDealsPriceDetailByDealNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentDealsPriceDetailByDealNameResponseParams struct {
	// 子订单的费用详情
	DealList []*DealPriceDetail `json:"DealList,omitnil,omitempty" name:"DealList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentDealsPriceDetailByDealNameResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentDealsPriceDetailByDealNameResponseParams `json:"Response"`
}

func (r *DescribeAgentDealsPriceDetailByDealNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentDealsPriceDetailByDealNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentPayDealsV2RequestParams struct {
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目 最大100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 下单时间范围起始点(不传时会默认查最近15天内订单，传值时需要传最近15天内的起始时间)
	CreatTimeRangeStart *string `json:"CreatTimeRangeStart,omitnil,omitempty" name:"CreatTimeRangeStart"`

	// 下单时间范围终止点
	CreatTimeRangeEnd *string `json:"CreatTimeRangeEnd,omitnil,omitempty" name:"CreatTimeRangeEnd"`

	// 0:下单时间降序；其他：下单时间升序
	Order *uint64 `json:"Order,omitnil,omitempty" name:"Order"`

	// 订单的状态(1：未支付;2：已支付;3：发货中;4：已发货;5：发货失败;6：已退款;7：已关单;8：订单过期;9：订单已失效;10：产品已失效;11：代付拒绝;12：支付中)
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 下单人账号ID列表
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// 子订单号列表
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 大订单号列表
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`
}

type DescribeAgentPayDealsV2Request struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目 最大100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 下单时间范围起始点(不传时会默认查最近15天内订单，传值时需要传最近15天内的起始时间)
	CreatTimeRangeStart *string `json:"CreatTimeRangeStart,omitnil,omitempty" name:"CreatTimeRangeStart"`

	// 下单时间范围终止点
	CreatTimeRangeEnd *string `json:"CreatTimeRangeEnd,omitnil,omitempty" name:"CreatTimeRangeEnd"`

	// 0:下单时间降序；其他：下单时间升序
	Order *uint64 `json:"Order,omitnil,omitempty" name:"Order"`

	// 订单的状态(1：未支付;2：已支付;3：发货中;4：已发货;5：发货失败;6：已退款;7：已关单;8：订单过期;9：订单已失效;10：产品已失效;11：代付拒绝;12：支付中)
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 下单人账号ID列表
	OwnerUins []*string `json:"OwnerUins,omitnil,omitempty" name:"OwnerUins"`

	// 子订单号列表
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 大订单号列表
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`
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
	AgentPayDealSet []*AgentDealNewElem `json:"AgentPayDealSet,omitnil,omitempty" name:"AgentPayDealSet"`

	// 符合条件的订单总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeAgentRelateBigDealIdsRequestParams struct {
	// 大订单号
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`
}

type DescribeAgentRelateBigDealIdsRequest struct {
	*tchttp.BaseRequest
	
	// 大订单号
	BigDealId *string `json:"BigDealId,omitnil,omitempty" name:"BigDealId"`
}

func (r *DescribeAgentRelateBigDealIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentRelateBigDealIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BigDealId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentRelateBigDealIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentRelateBigDealIdsResponseParams struct {
	// 申请合并支付的关联大订单号列表（不包含请求的订单号）
	BigDealIdList []*string `json:"BigDealIdList,omitnil,omitempty" name:"BigDealIdList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentRelateBigDealIdsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentRelateBigDealIdsResponseParams `json:"Response"`
}

func (r *DescribeAgentRelateBigDealIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentRelateBigDealIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentSelfPayDealsV2RequestParams struct {
	// 下单人账号ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目 最大100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 下单时间范围起始点(不传时会默认查最近15天内订单，传值时需要传最近15天内的起始时间)
	CreatTimeRangeStart *string `json:"CreatTimeRangeStart,omitnil,omitempty" name:"CreatTimeRangeStart"`

	// 下单时间范围终止点
	CreatTimeRangeEnd *string `json:"CreatTimeRangeEnd,omitnil,omitempty" name:"CreatTimeRangeEnd"`

	// 0:下单时间降序；其他：下单时间升序
	Order *uint64 `json:"Order,omitnil,omitempty" name:"Order"`

	// 订单的状态(1：未支付;2：已支付;3：发货中;4：已发货;5：发货失败;6：已退款;7：已关单;8：订单过期;9：订单已失效;10：产品已失效;11：代付拒绝;12：支付中)
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 子订单号列表
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 大订单号列表
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`
}

type DescribeAgentSelfPayDealsV2Request struct {
	*tchttp.BaseRequest
	
	// 下单人账号ID
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目 最大100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 下单时间范围起始点(不传时会默认查最近15天内订单，传值时需要传最近15天内的起始时间)
	CreatTimeRangeStart *string `json:"CreatTimeRangeStart,omitnil,omitempty" name:"CreatTimeRangeStart"`

	// 下单时间范围终止点
	CreatTimeRangeEnd *string `json:"CreatTimeRangeEnd,omitnil,omitempty" name:"CreatTimeRangeEnd"`

	// 0:下单时间降序；其他：下单时间升序
	Order *uint64 `json:"Order,omitnil,omitempty" name:"Order"`

	// 订单的状态(1：未支付;2：已支付;3：发货中;4：已发货;5：发货失败;6：已退款;7：已关单;8：订单过期;9：订单已失效;10：产品已失效;11：代付拒绝;12：支付中)
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 子订单号列表
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 大订单号列表
	BigDealIds []*string `json:"BigDealIds,omitnil,omitempty" name:"BigDealIds"`
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
	AgentPayDealSet []*AgentDealNewElem `json:"AgentPayDealSet,omitnil,omitempty" name:"AgentPayDealSet"`

	// 符合条件的订单总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`
}

type DescribeClientBalanceNewRequest struct {
	*tchttp.BaseRequest
	
	// 客户(代客)账号ID
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`
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
	Balance *int64 `json:"Balance,omitnil,omitempty" name:"Balance"`

	// 账户现金余额，单位分
	Cash *int64 `json:"Cash,omitnil,omitempty" name:"Cash"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeClientJoinIncreaseListRequestParams struct {
	// 客户UIN列表
	ClientUins []*string `json:"ClientUins,omitnil,omitempty" name:"ClientUins"`
}

type DescribeClientJoinIncreaseListRequest struct {
	*tchttp.BaseRequest
	
	// 客户UIN列表
	ClientUins []*string `json:"ClientUins,omitnil,omitempty" name:"ClientUins"`
}

func (r *DescribeClientJoinIncreaseListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientJoinIncreaseListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUins")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClientJoinIncreaseListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientJoinIncreaseListResponseParams struct {
	// 已审核代客列表
	List []*ClientIncreaseInfoList `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClientJoinIncreaseListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClientJoinIncreaseListResponseParams `json:"Response"`
}

func (r *DescribeClientJoinIncreaseListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientJoinIncreaseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientSwitchTraTaskInfoRequestParams struct {
	// 代客UIN
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// 1：代理，2：代采
	SwitchType *int64 `json:"SwitchType,omitnil,omitempty" name:"SwitchType"`
}

type DescribeClientSwitchTraTaskInfoRequest struct {
	*tchttp.BaseRequest
	
	// 代客UIN
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// 1：代理，2：代采
	SwitchType *int64 `json:"SwitchType,omitnil,omitempty" name:"SwitchType"`
}

func (r *DescribeClientSwitchTraTaskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientSwitchTraTaskInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClientUin")
	delete(f, "SwitchType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClientSwitchTraTaskInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientSwitchTraTaskInfoResponseParams struct {
	// 客户UIN
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`

	// 切换类型：代理,代采
	SwitchType *string `json:"SwitchType,omitnil,omitempty" name:"SwitchType"`

	// ok，符合，fail，不符合
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 切换链接
	SwitchUrl *string `json:"SwitchUrl,omitnil,omitempty" name:"SwitchUrl"`

	// 不符合的原因
	ResultMsg *string `json:"ResultMsg,omitnil,omitempty" name:"ResultMsg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClientSwitchTraTaskInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClientSwitchTraTaskInfoResponseParams `json:"Response"`
}

func (r *DescribeClientSwitchTraTaskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientSwitchTraTaskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRebateInfosNewRequestParams struct {
	// 返佣月份，如2018-02
	RebateMonth *string `json:"RebateMonth,omitnil,omitempty" name:"RebateMonth"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeRebateInfosNewRequest struct {
	*tchttp.BaseRequest
	
	// 返佣月份，如2018-02
	RebateMonth *string `json:"RebateMonth,omitnil,omitempty" name:"RebateMonth"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	RebateInfoSet []*RebateInfoElemNew `json:"RebateInfoSet,omitnil,omitempty" name:"RebateInfoSet"`

	// 符合查询条件返佣信息数目
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	RebateMonth *string `json:"RebateMonth,omitnil,omitempty" name:"RebateMonth"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeRebateInfosRequest struct {
	*tchttp.BaseRequest
	
	// 返佣月份，如2018-02
	RebateMonth *string `json:"RebateMonth,omitnil,omitempty" name:"RebateMonth"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	RebateInfoSet []*RebateInfoElem `json:"RebateInfoSet,omitnil,omitempty" name:"RebateInfoSet"`

	// 符合查询条件返佣信息数目
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 业务员姓名(模糊查询)
	SalesName *string `json:"SalesName,omitnil,omitempty" name:"SalesName"`

	// 业务员ID
	SalesUin *string `json:"SalesUin,omitnil,omitempty" name:"SalesUin"`

	// ASC/DESC， 不区分大小写，按创建通过时间排序
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

type DescribeSalesmansRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 业务员姓名(模糊查询)
	SalesName *string `json:"SalesName,omitnil,omitempty" name:"SalesName"`

	// 业务员ID
	SalesUin *string `json:"SalesUin,omitnil,omitempty" name:"SalesUin"`

	// ASC/DESC， 不区分大小写，按创建通过时间排序
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
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
	AgentSalesmanSet []*AgentSalesmanElem `json:"AgentSalesmanSet,omitnil,omitempty" name:"AgentSalesmanSet"`

	// 符合条件的代客总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 解绑账号ID
	UnbindUin *string `json:"UnbindUin,omitnil,omitempty" name:"UnbindUin"`

	// 解绑申请时间范围起始点
	ApplyTimeStart *string `json:"ApplyTimeStart,omitnil,omitempty" name:"ApplyTimeStart"`

	// 解绑申请时间范围终止点
	ApplyTimeEnd *string `json:"ApplyTimeEnd,omitnil,omitempty" name:"ApplyTimeEnd"`

	// 对申请时间的升序降序，值：asc，desc
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

type DescribeUnbindClientListRequest struct {
	*tchttp.BaseRequest
	
	// 解绑状态：0:所有,1:审核中,2已解绑
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限制数目
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 解绑账号ID
	UnbindUin *string `json:"UnbindUin,omitnil,omitempty" name:"UnbindUin"`

	// 解绑申请时间范围起始点
	ApplyTimeStart *string `json:"ApplyTimeStart,omitnil,omitempty" name:"ApplyTimeStart"`

	// 解绑申请时间范围终止点
	ApplyTimeEnd *string `json:"ApplyTimeEnd,omitnil,omitempty" name:"ApplyTimeEnd"`

	// 对申请时间的升序降序，值：asc，desc
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合条件的解绑客户列表
	UnbindClientList []*UnbindClientElem `json:"UnbindClientList,omitnil,omitempty" name:"UnbindClientList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ClientRemark *string `json:"ClientRemark,omitnil,omitempty" name:"ClientRemark"`

	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`
}

type ModifyClientRemarkRequest struct {
	*tchttp.BaseRequest
	
	// 客户备注名称
	ClientRemark *string `json:"ClientRemark,omitnil,omitempty" name:"ClientRemark"`

	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 产品属性值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type RebateInfoElem struct {
	// 代理商账号ID
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 返佣月份，如2018-02
	RebateMonth *string `json:"RebateMonth,omitnil,omitempty" name:"RebateMonth"`

	// 返佣金额，单位分
	Amt *uint64 `json:"Amt,omitnil,omitempty" name:"Amt"`

	// 月度业绩，单位分
	MonthSales *uint64 `json:"MonthSales,omitnil,omitempty" name:"MonthSales"`

	// 季度业绩，单位分
	QuarterSales *uint64 `json:"QuarterSales,omitnil,omitempty" name:"QuarterSales"`

	// NORMAL(正常)/HAS_OVERDUE_BILL(欠费)/NO_CONTRACT(缺合同)
	ExceptionFlag *string `json:"ExceptionFlag,omitnil,omitempty" name:"ExceptionFlag"`
}

type RebateInfoElemNew struct {
	// 代理商账号ID
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 返佣月份，如2018-02
	RebateMonth *string `json:"RebateMonth,omitnil,omitempty" name:"RebateMonth"`

	// 返佣金额，单位分
	Amt *int64 `json:"Amt,omitnil,omitempty" name:"Amt"`

	// 月度业绩，单位分
	MonthSales *int64 `json:"MonthSales,omitnil,omitempty" name:"MonthSales"`

	// 季度业绩，单位分
	QuarterSales *int64 `json:"QuarterSales,omitnil,omitempty" name:"QuarterSales"`

	// NORMAL(正常)/HAS_OVERDUE_BILL(欠费)/NO_CONTRACT(缺合同)
	ExceptionFlag *string `json:"ExceptionFlag,omitnil,omitempty" name:"ExceptionFlag"`
}

type RefundMap struct {
	// 退款单关联的原始子订单号
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 退款金额，单位分
	RefundAmount *int64 `json:"RefundAmount,omitnil,omitempty" name:"RefundAmount"`
}

// Predefined struct for user
type RemovePayRelationForClientRequestParams struct {
	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`
}

type RemovePayRelationForClientRequest struct {
	*tchttp.BaseRequest
	
	// 客户账号ID
	ClientUin *string `json:"ClientUin,omitnil,omitempty" name:"ClientUin"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type SubProductPriceDetail struct {
	// 子产品名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 折扣值，=100时表示无折扣，=85时表示8.5折
	DiscountValue *float64 `json:"DiscountValue,omitnil,omitempty" name:"DiscountValue"`

	// 原价，折扣前价格，单位：分
	TotalCost *int64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 折后价，单位：分
	RealTotalCost *int64 `json:"RealTotalCost,omitnil,omitempty" name:"RealTotalCost"`
}

type UnbindClientElem struct {
	// 解绑账号ID
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 状态：0:审核中；1：已解绑；2：已撤销 3：关联撤销 4: 已驳回
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 申请时间
	ApplyTime *string `json:"ApplyTime,omitnil,omitempty" name:"ApplyTime"`

	// 解绑/撤销时间
	ActionTime *string `json:"ActionTime,omitnil,omitempty" name:"ActionTime"`
}