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

package v20240125

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type CreateSavingPlanOrderRequestParams struct {
	// 地域编码
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 区域编码
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 预付费类型
	PrePayType *string `json:"PrePayType,omitnil,omitempty" name:"PrePayType"`

	// 时长
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 时长单位
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 商品唯一标识
	CommodityCode *string `json:"CommodityCode,omitnil,omitempty" name:"CommodityCode"`

	// 承诺时长内的小额金额（单位：元）
	PromiseUseAmount *int64 `json:"PromiseUseAmount,omitnil,omitempty" name:"PromiseUseAmount"`

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
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 时长单位
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 商品唯一标识
	CommodityCode *string `json:"CommodityCode,omitnil,omitempty" name:"CommodityCode"`

	// 承诺时长内的小额金额（单位：元）
	PromiseUseAmount *int64 `json:"PromiseUseAmount,omitnil,omitempty" name:"PromiseUseAmount"`

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