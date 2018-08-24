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

package v20180228

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DescribePersonInfoRequest struct {
	*tchttp.BaseRequest
	// 公司ID
	CompanyId *string `json:"CompanyId" name:"CompanyId"`
	// 门店ID
	ShopId *uint64 `json:"ShopId" name:"ShopId"`
	// 起始ID，第一次拉取时StartPersonId传0，后续送入的值为上一页最后一条数据项的PersonId
	StartPersonId *uint64 `json:"StartPersonId" name:"StartPersonId"`
	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset *uint64 `json:"Offset" name:"Offset"`
	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit *uint64 `json:"Limit" name:"Limit"`
}

func (r *DescribePersonInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePersonInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePersonInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 公司ID
		CompanyId *string `json:"CompanyId" name:"CompanyId"`
		// 门店ID
		ShopId *uint64 `json:"ShopId" name:"ShopId"`
		// 总数
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 用户信息
		PersonInfoSet []*PersonInfo `json:"PersonInfoSet" name:"PersonInfoSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePersonInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePersonInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePersonVisitInfoRequest struct {
	*tchttp.BaseRequest
	// 公司ID
	CompanyId *string `json:"CompanyId" name:"CompanyId"`
	// 门店ID
	ShopId *uint64 `json:"ShopId" name:"ShopId"`
	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate" name:"StartDate"`
	// 结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate" name:"EndDate"`
	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset *uint64 `json:"Offset" name:"Offset"`
	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit *uint64 `json:"Limit" name:"Limit"`
}

func (r *DescribePersonVisitInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePersonVisitInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePersonVisitInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 公司ID
		CompanyId *string `json:"CompanyId" name:"CompanyId"`
		// 门店ID
		ShopId *uint64 `json:"ShopId" name:"ShopId"`
		// 总数
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 用户到访明细
		PersonVisitInfoSet []*PersonVisitInfo `json:"PersonVisitInfoSet" name:"PersonVisitInfoSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePersonVisitInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePersonVisitInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeShopHourTrafficInfoRequest struct {
	*tchttp.BaseRequest
	// 公司ID
	CompanyId *string `json:"CompanyId" name:"CompanyId"`
	// 门店ID
	ShopId *uint64 `json:"ShopId" name:"ShopId"`
	// 开始日期，格式：yyyy-MM-dd
	StartDate *string `json:"StartDate" name:"StartDate"`
	// 结束日期，格式：yyyy-MM-dd
	EndDate *string `json:"EndDate" name:"EndDate"`
	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset *uint64 `json:"Offset" name:"Offset"`
	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit *uint64 `json:"Limit" name:"Limit"`
}

func (r *DescribeShopHourTrafficInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeShopHourTrafficInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeShopHourTrafficInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 公司ID
		CompanyId *string `json:"CompanyId" name:"CompanyId"`
		// 门店ID
		ShopId *uint64 `json:"ShopId" name:"ShopId"`
		// 查询结果总数
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 分时客流信息
		ShopHourTrafficInfoSet []*ShopHourTrafficInfo `json:"ShopHourTrafficInfoSet" name:"ShopHourTrafficInfoSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShopHourTrafficInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeShopHourTrafficInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeShopInfoRequest struct {
	*tchttp.BaseRequest
	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset *uint64 `json:"Offset" name:"Offset"`
	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit *uint64 `json:"Limit" name:"Limit"`
}

func (r *DescribeShopInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeShopInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeShopInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 门店总数
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 门店列表信息
		ShopInfoSet []*ShopInfo `json:"ShopInfoSet" name:"ShopInfoSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShopInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeShopInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeShopTrafficInfoRequest struct {
	*tchttp.BaseRequest
	// 公司ID
	CompanyId *string `json:"CompanyId" name:"CompanyId"`
	// 门店ID
	ShopId *uint64 `json:"ShopId" name:"ShopId"`
	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate" name:"StartDate"`
	// 介绍日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate" name:"EndDate"`
	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset *uint64 `json:"Offset" name:"Offset"`
	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit *uint64 `json:"Limit" name:"Limit"`
}

func (r *DescribeShopTrafficInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeShopTrafficInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeShopTrafficInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 公司ID
		CompanyId *string `json:"CompanyId" name:"CompanyId"`
		// 门店ID
		ShopId *uint64 `json:"ShopId" name:"ShopId"`
		// 查询结果总数
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 客流信息列表
		ShopDayTrafficInfoSet []*ShopDayTrafficInfo `json:"ShopDayTrafficInfoSet" name:"ShopDayTrafficInfoSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShopTrafficInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeShopTrafficInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneTrafficInfoRequest struct {
	*tchttp.BaseRequest
	// 公司ID
	CompanyId *string `json:"CompanyId" name:"CompanyId"`
	// 店铺ID
	ShopId *uint64 `json:"ShopId" name:"ShopId"`
	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate" name:"StartDate"`
	// 结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate" name:"EndDate"`
	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset *uint64 `json:"Offset" name:"Offset"`
	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit *uint64 `json:"Limit" name:"Limit"`
}

func (r *DescribeZoneTrafficInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZoneTrafficInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneTrafficInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 公司ID
		CompanyId *string `json:"CompanyId" name:"CompanyId"`
		// 门店ID
		ShopId *uint64 `json:"ShopId" name:"ShopId"`
		// 查询结果总数
		TotalCount *uint64 `json:"TotalCount" name:"TotalCount"`
		// 区域客流信息列表
		ZoneTrafficInfoSet []*ZoneTrafficInfo `json:"ZoneTrafficInfoSet" name:"ZoneTrafficInfoSet" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneTrafficInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeZoneTrafficInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GenderAgeTrafficDetail struct {
	// 性别: 0男1女
	Gender *uint64 `json:"Gender" name:"Gender"`
	// 年龄区间，枚举值：0-17、18-23、24-30、31-40、41-50、51-60、>60
	AgeGap *string `json:"AgeGap" name:"AgeGap"`
	// 客流量
	TrafficCount *uint64 `json:"TrafficCount" name:"TrafficCount"`
}

type HourTrafficInfoDetail struct {
	// 小时 取值为：0，1，2，3，4，5，6，7，8，9，10，11，12，13，14，15，16，17，18，19，20，21，22，23
	Hour *uint64 `json:"Hour" name:"Hour"`
	// 分时客流量
	HourTrafficTotalCount *uint64 `json:"HourTrafficTotalCount" name:"HourTrafficTotalCount"`
}

type PersonInfo struct {
	// 用户ID
	PersonId *uint64 `json:"PersonId" name:"PersonId"`
	// 人脸图片，这里返回的是图片内容的Base64编码
	PersonPicture *string `json:"PersonPicture" name:"PersonPicture"`
	// 性别：0男1女
	Gender *uint64 `json:"Gender" name:"Gender"`
	// 年龄
	Age *uint64 `json:"Age" name:"Age"`
	// 身份类型：0-普通顾客，1~10黑名单，11~20白名单，11店员
	PersonType *uint64 `json:"PersonType" name:"PersonType"`
}

type PersonVisitInfo struct {
	// 用户ID
	PersonId *uint64 `json:"PersonId" name:"PersonId"`
	// 用户到访ID
	VisitId *uint64 `json:"VisitId" name:"VisitId"`
	// 到访时间：Unix时间戳
	InTime *uint64 `json:"InTime" name:"InTime"`
	// 抓拍到的头像，这里返回的是图片内容的Base64编码
	CapturedPicture *string `json:"CapturedPicture" name:"CapturedPicture"`
	// 口罩类型：0不戴口罩，1戴口罩
	MaskType *uint64 `json:"MaskType" name:"MaskType"`
	// 眼镜类型：0不戴眼镜，1普通眼镜 , 2墨镜
	GlassType *uint64 `json:"GlassType" name:"GlassType"`
	// 发型：0 短发,  1长发
	HairType *uint64 `json:"HairType" name:"HairType"`
}

type RegisterCallbackRequest struct {
	*tchttp.BaseRequest
	// 集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId *string `json:"CompanyId" name:"CompanyId"`
	// 通知回调地址，完整url，示例（http://youmall.tencentcloudapi.com/）
	BackUrl *string `json:"BackUrl" name:"BackUrl"`
	// 请求时间戳
	Time *uint64 `json:"Time" name:"Time"`
	// 是否需要顾客图片，1-需要图片，其它-不需要图片
	NeedFacePic *uint64 `json:"NeedFacePic" name:"NeedFacePic"`
}

func (r *RegisterCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RegisterCallbackRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RegisterCallbackResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *RegisterCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RegisterCallbackResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ShopDayTrafficInfo struct {
	// 日期
	Date *string `json:"Date" name:"Date"`
	// 客流量
	DayTrafficTotalCount *uint64 `json:"DayTrafficTotalCount" name:"DayTrafficTotalCount"`
	// 性别年龄分组下的客流信息
	GenderAgeTrafficDetailSet []*GenderAgeTrafficDetail `json:"GenderAgeTrafficDetailSet" name:"GenderAgeTrafficDetailSet" list`
}

type ShopHourTrafficInfo struct {
	// 日期，格式yyyy-MM-dd
	Date *string `json:"Date" name:"Date"`
	// 分时客流详细信息
	HourTrafficInfoDetailSet []*HourTrafficInfoDetail `json:"HourTrafficInfoDetailSet" name:"HourTrafficInfoDetailSet" list`
}

type ShopInfo struct {
	// 公司ID
	CompanyId *string `json:"CompanyId" name:"CompanyId"`
	// 门店ID
	ShopId *uint64 `json:"ShopId" name:"ShopId"`
	// 门店名称
	ShopName *string `json:"ShopName" name:"ShopName"`
	// 客户门店编码
	ShopCode *string `json:"ShopCode" name:"ShopCode"`
	// 省
	Province *string `json:"Province" name:"Province"`
	// 市
	City *string `json:"City" name:"City"`
	// 公司名称
	CompanyName *string `json:"CompanyName" name:"CompanyName"`
}

type ZoneTrafficInfo struct {
	// 日期
	Date *string `json:"Date" name:"Date"`
	// 门店区域客流详细信息
	ZoneTrafficInfoDetailSet []*ZoneTrafficInfoDetail `json:"ZoneTrafficInfoDetailSet" name:"ZoneTrafficInfoDetailSet" list`
}

type ZoneTrafficInfoDetail struct {
	// 区域ID
	ZoneId *uint64 `json:"ZoneId" name:"ZoneId"`
	// 区域名称
	ZoneName *string `json:"ZoneName" name:"ZoneName"`
	// 客流量
	TrafficTotalCount *uint64 `json:"TrafficTotalCount" name:"TrafficTotalCount"`
	// 平均停留时间
	AvgStayTime *uint64 `json:"AvgStayTime" name:"AvgStayTime"`
}
