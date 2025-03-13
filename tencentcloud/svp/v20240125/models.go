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
type DescribeSavingPlanDeductRequestParams struct {
	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，最大值为200
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 节省计划使用开始的查询结束时间
	StartEndDate *string `json:"StartEndDate,omitnil,omitempty" name:"StartEndDate"`

	// 节省计划使用开始的查询开始时间
	StartStartDate *string `json:"StartStartDate,omitnil,omitempty" name:"StartStartDate"`

	// 地域编码
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 区域编码
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 节省计划资源id
	SpId *string `json:"SpId,omitnil,omitempty" name:"SpId"`

	// 抵扣查询结束时间，格式：yyyy-MM-dd HH:mm:ss
	DeductEndDate *string `json:"DeductEndDate,omitnil,omitempty" name:"DeductEndDate"`

	// 抵扣查询开始时间，格式：yyyy-MM-dd HH:mm:ss
	DeductStartDate *string `json:"DeductStartDate,omitnil,omitempty" name:"DeductStartDate"`

	// 节省计划使用结束的查询结束时间
	EndEndDate *string `json:"EndEndDate,omitnil,omitempty" name:"EndEndDate"`

	// 节省计划使用结束的查询开始时间
	EndStartDate *string `json:"EndStartDate,omitnil,omitempty" name:"EndStartDate"`
}

type DescribeSavingPlanDeductRequest struct {
	*tchttp.BaseRequest
	
	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页数量，最大值为200
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 节省计划使用开始的查询结束时间
	StartEndDate *string `json:"StartEndDate,omitnil,omitempty" name:"StartEndDate"`

	// 节省计划使用开始的查询开始时间
	StartStartDate *string `json:"StartStartDate,omitnil,omitempty" name:"StartStartDate"`

	// 地域编码
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 区域编码
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 节省计划资源id
	SpId *string `json:"SpId,omitnil,omitempty" name:"SpId"`

	// 抵扣查询结束时间，格式：yyyy-MM-dd HH:mm:ss
	DeductEndDate *string `json:"DeductEndDate,omitnil,omitempty" name:"DeductEndDate"`

	// 抵扣查询开始时间，格式：yyyy-MM-dd HH:mm:ss
	DeductStartDate *string `json:"DeductStartDate,omitnil,omitempty" name:"DeductStartDate"`

	// 节省计划使用结束的查询结束时间
	EndEndDate *string `json:"EndEndDate,omitnil,omitempty" name:"EndEndDate"`

	// 节省计划使用结束的查询开始时间
	EndStartDate *string `json:"EndStartDate,omitnil,omitempty" name:"EndStartDate"`
}

func (r *DescribeSavingPlanDeductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSavingPlanDeductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "StartEndDate")
	delete(f, "StartStartDate")
	delete(f, "RegionId")
	delete(f, "ZoneId")
	delete(f, "SpId")
	delete(f, "DeductEndDate")
	delete(f, "DeductStartDate")
	delete(f, "EndEndDate")
	delete(f, "EndStartDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSavingPlanDeductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSavingPlanDeductResponseParams struct {
	// 查询命中的节省计划抵扣明细数据总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 查询命中的节省计划抵扣明细数据明细
	Deducts []*SavingPlanDeductDetail `json:"Deducts,omitnil,omitempty" name:"Deducts"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSavingPlanDeductResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSavingPlanDeductResponseParams `json:"Response"`
}

func (r *DescribeSavingPlanDeductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSavingPlanDeductResponse) FromJsonString(s string) error {
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
	// 注意：此字段可能返回 null，表示取不到有效值。
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
	// 查询命中的节省计划总览明细数据总条数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 节省计划使用率数据
	Usages []*SavingPlanUsageDetail `json:"Usages,omitnil,omitempty" name:"Usages"`

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

type SavingPlanCoverageDetail struct {
	// 资源 ID：账单中出账对象 ID，不同产品因资源形态不同，资源内容不完全相同，如云服务器 CVM 为对应的实例 ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 地域ID
	RegionId *uint64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 产品编码
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 子产品编码
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

	// 支付者昵称
	PayerUinName *string `json:"PayerUinName,omitnil,omitempty" name:"PayerUinName"`

	// 使用者昵称
	OwnerUinName *string `json:"OwnerUinName,omitnil,omitempty" name:"OwnerUinName"`

	// 支付者uin
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 计费项名称
	SubBillingItemName *string `json:"SubBillingItemName,omitnil,omitempty" name:"SubBillingItemName"`

	// 计费细项名称
	BillingItemName *string `json:"BillingItemName,omitnil,omitempty" name:"BillingItemName"`

	// 子产品名称
	SubProductName *string `json:"SubProductName,omitnil,omitempty" name:"SubProductName"`
}

type SavingPlanCoverageRate struct {
	// 聚合时间维度，按天聚合格式为yyyy-MM-dd，按月聚合格式为yyyy-MM
	DatePoint *string `json:"DatePoint,omitnil,omitempty" name:"DatePoint"`

	// 覆盖率结果，取值[0, 100]
	Rate *float64 `json:"Rate,omitnil,omitempty" name:"Rate"`
}

type SavingPlanDeductDetail struct {
	// 账号id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 账号名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUinName *string `json:"OwnerUinName,omitnil,omitempty" name:"OwnerUinName"`

	// 抵扣账号id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayerUin *string `json:"PayerUin,omitnil,omitempty" name:"PayerUin"`

	// 抵扣账号名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayerUinName *string `json:"PayerUinName,omitnil,omitempty" name:"PayerUinName"`

	// 节省计划资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpId *string `json:"SpId,omitnil,omitempty" name:"SpId"`

	// 产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductCode *string `json:"ProductCode,omitnil,omitempty" name:"ProductCode"`

	// 产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 子产品编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// 子产品名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubProductName *string `json:"SubProductName,omitnil,omitempty" name:"SubProductName"`

	// 交易ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutTradeNo *string `json:"OutTradeNo,omitnil,omitempty" name:"OutTradeNo"`

	// 地域id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionId *int64 `json:"RegionId,omitnil,omitempty" name:"RegionId"`

	// 地域名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 地区id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *int64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 地区名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 开始使用时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpStartTime *string `json:"SpStartTime,omitnil,omitempty" name:"SpStartTime"`

	// 结束使用时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpEndTime *string `json:"SpEndTime,omitnil,omitempty" name:"SpEndTime"`

	// 折扣时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeductTime *string `json:"DeductTime,omitnil,omitempty" name:"DeductTime"`

	// 抵扣金额，单位分
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeductAmount *string `json:"DeductAmount,omitnil,omitempty" name:"DeductAmount"`

	// 抵扣折扣率
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeductDiscount *string `json:"DeductDiscount,omitnil,omitempty" name:"DeductDiscount"`

	// 抵扣比率
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeductRate *string `json:"DeductRate,omitnil,omitempty" name:"DeductRate"`
}

type SavingPlanOverviewDetail struct {
	// 节省计划类型
	SpType *string `json:"SpType,omitnil,omitempty" name:"SpType"`

	// 支付金额（单位：元）
	PayAmount *string `json:"PayAmount,omitnil,omitempty" name:"PayAmount"`

	// 开始时间 yyyy-mm-dd HH:mm:ss格式
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间 yyyy-mm-dd HH:mm:ss格式
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 累计节省金额（单位：元）
	SavingAmount *string `json:"SavingAmount,omitnil,omitempty" name:"SavingAmount"`

	// 地域
	Region []*string `json:"Region,omitnil,omitempty" name:"Region"`

	// 支付类型
	PayType *uint64 `json:"PayType,omitnil,omitempty" name:"PayType"`

	// 购买时间 yyyy-mm-dd HH:mm:ss格式
	BuyTime *string `json:"BuyTime,omitnil,omitempty" name:"BuyTime"`
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
	Region []*string `json:"Region,omitnil,omitempty" name:"Region"`
}