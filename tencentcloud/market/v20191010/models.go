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

package v20191010

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Error struct {

	// 一级错误描述
	Code *string `json:"Code,omitempty" name:"Code"`

	// 二级错误描述
	Message *string `json:"Message,omitempty" name:"Message"`
}

type FlowProductRemindRequest struct {
	*tchttp.BaseRequest

	// 服务商uin
	ProviderUin *string `json:"ProviderUin,omitempty" name:"ProviderUin"`

	// 服务商实例ID
	SignId *string `json:"SignId,omitempty" name:"SignId"`

	// 云市场实例ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 实例总流量
	TotalFlow *string `json:"TotalFlow,omitempty" name:"TotalFlow"`

	// 剩余流量
	LeftFlow *string `json:"LeftFlow,omitempty" name:"LeftFlow"`

	// 流量单位
	FlowUnit *string `json:"FlowUnit,omitempty" name:"FlowUnit"`
}

func (r *FlowProductRemindRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlowProductRemindRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProviderUin")
	delete(f, "SignId")
	delete(f, "ResourceId")
	delete(f, "TotalFlow")
	delete(f, "LeftFlow")
	delete(f, "FlowUnit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FlowProductRemindRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type FlowProductRemindResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否成功
		Success *string `json:"Success,omitempty" name:"Success"`

		// 流水号
		FlowId *string `json:"FlowId,omitempty" name:"FlowId"`

		// 消息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Info *string `json:"Info,omitempty" name:"Info"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FlowProductRemindResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FlowProductRemindResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUsagePlanUsageAmountRequest struct {
	*tchttp.BaseRequest

	// 用于查询实例的Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *GetUsagePlanUsageAmountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUsagePlanUsageAmountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetUsagePlanUsageAmountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetUsagePlanUsageAmountResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 最大调用量
		MaxRequestNum *int64 `json:"MaxRequestNum,omitempty" name:"MaxRequestNum"`

		// 已经调用量
		InUseRequestNum *int64 `json:"InUseRequestNum,omitempty" name:"InUseRequestNum"`

		// 剩余调用量
		RemainingRequestNum *int64 `json:"RemainingRequestNum,omitempty" name:"RemainingRequestNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUsagePlanUsageAmountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetUsagePlanUsageAmountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OfflineProductDeal struct {

	// 产品标签
	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`

	// 子产品标签
	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`

	// 计费项
	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`

	// 计费细项
	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`

	// 单价（单位：分）
	UnitPrice *float64 `json:"UnitPrice,omitempty" name:"UnitPrice"`

	// 用量
	Dosage *float64 `json:"Dosage,omitempty" name:"Dosage"`

	// 用量单位
	DosageUnit *string `json:"DosageUnit,omitempty" name:"DosageUnit"`

	// 商品的时间大小，当TimeUnit 是package时，timeSpan 只能传1。当TimeUnit是year；month；day；hour；minute；second，传具体时长。
	TimeSpan *float64 `json:"TimeSpan,omitempty" name:"TimeSpan"`

	// 商品的时间单位：year:年；month:月；day:日；hour:小时；minute:分钟；second:秒; package:与价格无关,一次性购买的产品传package
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// 原价 （单位：分）OriginPrice=round(UnitPrice * Dosage * TimeSpan)
	OriginPrice *int64 `json:"OriginPrice,omitempty" name:"OriginPrice"`

	// 折扣百分比，传入0-100的值
	Discount *float64 `json:"Discount,omitempty" name:"Discount"`

	// 最终扣费金额（单位：分）Fee=round(OriginPrice*Discount/100)
	Fee *int64 `json:"Fee,omitempty" name:"Fee"`
}

type SyncUserAndOrderInfoDetail struct {

	// 腾讯云订单总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 腾讯云订单列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MarketOrders []*string `json:"MarketOrders,omitempty" name:"MarketOrders"`

	// 腾讯云用户Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

type SyncUserAndOrderInfoRequest struct {
	*tchttp.BaseRequest

	// 企业微信用户信息
	UserInfo *WeChatUserInfo `json:"UserInfo,omitempty" name:"UserInfo"`

	// 企业微信订单信息
	OrderInfo *WeChatOrderInfo `json:"OrderInfo,omitempty" name:"OrderInfo"`
}

func (r *SyncUserAndOrderInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncUserAndOrderInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserInfo")
	delete(f, "OrderInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SyncUserAndOrderInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SyncUserAndOrderInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 同步用户信息订单信息详情
		Details *SyncUserAndOrderInfoDetail `json:"Details,omitempty" name:"Details"`

		// 返回信息 成功返回 0 success
		Info *Error `json:"Info,omitempty" name:"Info"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SyncUserAndOrderInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SyncUserAndOrderInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WeChatOrderInfo struct {

	// 企业微信订单号
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 订单状态。0-未⽀支付，1-已⽀支付，2-已关闭， 3-未⽀支付且已过期， 4-申请退款中， 5-申请退款成功， 6-退款被拒绝
	OrderStatus *int64 `json:"OrderStatus,omitempty" name:"OrderStatus"`

	// 订单类型。0-普通订单，1-扩容订单，2-续期，3-版本变更更
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 应⽤id
	SuiteId *string `json:"SuiteId,omitempty" name:"SuiteId"`

	// 购买版本ID
	EditionId *string `json:"EditionId,omitempty" name:"EditionId"`

	// 购买版本名称
	EditionName *string `json:"EditionName,omitempty" name:"EditionName"`

	// 实付款金额，单位分
	Price *int64 `json:"Price,omitempty" name:"Price"`

	// 下单时间
	OrderTime *int64 `json:"OrderTime,omitempty" name:"OrderTime"`

	// 付款时间
	PaidTime *int64 `json:"PaidTime,omitempty" name:"PaidTime"`

	// 备注
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 资源使用开始时间
	UseBeginTime *int64 `json:"UseBeginTime,omitempty" name:"UseBeginTime"`

	// 资源使用结束时间
	UseEndTime *int64 `json:"UseEndTime,omitempty" name:"UseEndTime"`

	// 磐石详细的四层订单
	Deals *OfflineProductDeal `json:"Deals,omitempty" name:"Deals"`
}

type WeChatUserInfo struct {

	// 客户企业的corpid
	PaidCorpId *string `json:"PaidCorpId,omitempty" name:"PaidCorpId"`

	// 客户企业的Corp全称
	PaidCorpName *string `json:"PaidCorpName,omitempty" name:"PaidCorpName"`
}
