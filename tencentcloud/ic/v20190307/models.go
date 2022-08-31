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

package v20190307

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AppInfo struct {
	// 应用ID
	Sdkappid *string `json:"Sdkappid,omitempty" name:"Sdkappid"`

	// 应用key
	Appkey *string `json:"Appkey,omitempty" name:"Appkey"`

	// 用户appid
	CloudAppid *string `json:"CloudAppid,omitempty" name:"CloudAppid"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 应用描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 应用类型
	BizType *int64 `json:"BizType,omitempty" name:"BizType"`

	// 用户Uin
	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type CardInfo struct {
	// 卡片ID
	Iccid *string `json:"Iccid,omitempty" name:"Iccid"`

	// 卡电话号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msisdn *string `json:"Msisdn,omitempty" name:"Msisdn"`

	// 卡imsi
	// 注意：此字段可能返回 null，表示取不到有效值。
	Imsi *string `json:"Imsi,omitempty" name:"Imsi"`

	// 卡imei
	// 注意：此字段可能返回 null，表示取不到有效值。
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// 应用ID
	Sdkappid *string `json:"Sdkappid,omitempty" name:"Sdkappid"`

	// 运营商编号
	Teleoperator *int64 `json:"Teleoperator,omitempty" name:"Teleoperator"`

	// 卡片状态 1:未激活 2：激活 3：停卡 5：销卡
	CardStatus *int64 `json:"CardStatus,omitempty" name:"CardStatus"`

	// 网络状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkStatus *int64 `json:"NetworkStatus,omitempty" name:"NetworkStatus"`

	// 激活时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivitedTime *string `json:"ActivitedTime,omitempty" name:"ActivitedTime"`

	// 资费类型，1 单卡，2 流量池
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// 套餐类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 流量池ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PoolId *string `json:"PoolId,omitempty" name:"PoolId"`

	// 周期套餐流量使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataUsedInPeriod *float64 `json:"DataUsedInPeriod,omitempty" name:"DataUsedInPeriod"`

	// 周期套餐总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataTotalInPeriod *float64 `json:"DataTotalInPeriod,omitempty" name:"DataTotalInPeriod"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductExpiredTime *string `json:"ProductExpiredTime,omitempty" name:"ProductExpiredTime"`

	// 描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 修改时间
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 套餐周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreorderCnt *int64 `json:"PreorderCnt,omitempty" name:"PreorderCnt"`

	// 激活被回调标志
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsActivated *int64 `json:"IsActivated,omitempty" name:"IsActivated"`

	// 订单ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderId *string `json:"OrderId,omitempty" name:"OrderId"`

	// 是否自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenew *int64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 0 不需要开通达量不停卡 1 需要开通达量不停卡
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowArrears *int64 `json:"AllowArrears,omitempty" name:"AllowArrears"`

	// 是否开通短信0:未开短信 1:开通短信
	NeedSms *int64 `json:"NeedSms,omitempty" name:"NeedSms"`

	// 供应商
	Provider *int64 `json:"Provider,omitempty" name:"Provider"`

	// 实名认证 0:无 1:未实名 2:已实名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificationState *int64 `json:"CertificationState,omitempty" name:"CertificationState"`

	// 其他流量信息,流量分离统计其他流量
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherData *float64 `json:"OtherData,omitempty" name:"OtherData"`
}

type CardList struct {
	// 卡片总数
	Total *string `json:"Total,omitempty" name:"Total"`

	// 卡片列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*CardInfo `json:"List,omitempty" name:"List"`
}

// Predefined struct for user
type DescribeAppRequestParams struct {
	// 物联卡应用ID
	Sdkappid *int64 `json:"Sdkappid,omitempty" name:"Sdkappid"`
}

type DescribeAppRequest struct {
	*tchttp.BaseRequest
	
	// 物联卡应用ID
	Sdkappid *int64 `json:"Sdkappid,omitempty" name:"Sdkappid"`
}

func (r *DescribeAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Sdkappid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppResponseParams struct {
	// 应用信息详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *AppInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAppResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppResponseParams `json:"Response"`
}

func (r *DescribeAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCardRequestParams struct {
	// 应用ID
	Sdkappid *int64 `json:"Sdkappid,omitempty" name:"Sdkappid"`

	// 卡片ID
	Iccid *string `json:"Iccid,omitempty" name:"Iccid"`
}

type DescribeCardRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	Sdkappid *int64 `json:"Sdkappid,omitempty" name:"Sdkappid"`

	// 卡片ID
	Iccid *string `json:"Iccid,omitempty" name:"Iccid"`
}

func (r *DescribeCardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCardRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Sdkappid")
	delete(f, "Iccid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCardRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCardResponseParams struct {
	// 卡片详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CardInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCardResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCardResponseParams `json:"Response"`
}

func (r *DescribeCardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCardsRequestParams struct {
	// 应用ID
	Sdkappid *string `json:"Sdkappid,omitempty" name:"Sdkappid"`

	// 偏移值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 列表限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeCardsRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	Sdkappid *string `json:"Sdkappid,omitempty" name:"Sdkappid"`

	// 偏移值
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 列表限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCardsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCardsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Sdkappid")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCardsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCardsResponseParams struct {
	// 卡片列表信息
	Data *CardList `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCardsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCardsResponseParams `json:"Response"`
}

func (r *DescribeCardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmsRequestParams struct {
	// 应用ID
	Sdkappid *int64 `json:"Sdkappid,omitempty" name:"Sdkappid"`

	// 卡片ID
	Iccid *string `json:"Iccid,omitempty" name:"Iccid"`

	// 卡片号码
	Msisdn *string `json:"Msisdn,omitempty" name:"Msisdn"`

	// 短信类型
	SmsType *int64 `json:"SmsType,omitempty" name:"SmsType"`

	// 开始时间  YYYY-MM-DD HH:mm:ss
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间  YYYY-MM-DD HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 小于200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeSmsRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	Sdkappid *int64 `json:"Sdkappid,omitempty" name:"Sdkappid"`

	// 卡片ID
	Iccid *string `json:"Iccid,omitempty" name:"Iccid"`

	// 卡片号码
	Msisdn *string `json:"Msisdn,omitempty" name:"Msisdn"`

	// 短信类型
	SmsType *int64 `json:"SmsType,omitempty" name:"SmsType"`

	// 开始时间  YYYY-MM-DD HH:mm:ss
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间  YYYY-MM-DD HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 小于200
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Sdkappid")
	delete(f, "Iccid")
	delete(f, "Msisdn")
	delete(f, "SmsType")
	delete(f, "BeginTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSmsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSmsResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 短信列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*ResSms `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeSmsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSmsResponseParams `json:"Response"`
}

func (r *DescribeSmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserCardRemarkRequestParams struct {
	// 应用ID
	Sdkappid *int64 `json:"Sdkappid,omitempty" name:"Sdkappid"`

	// 物联卡ICCID
	Iccid *string `json:"Iccid,omitempty" name:"Iccid"`

	// 备注信息，限50字
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type ModifyUserCardRemarkRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	Sdkappid *int64 `json:"Sdkappid,omitempty" name:"Sdkappid"`

	// 物联卡ICCID
	Iccid *string `json:"Iccid,omitempty" name:"Iccid"`

	// 备注信息，限50字
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyUserCardRemarkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserCardRemarkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Sdkappid")
	delete(f, "Iccid")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserCardRemarkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserCardRemarkResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyUserCardRemarkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserCardRemarkResponseParams `json:"Response"`
}

func (r *ModifyUserCardRemarkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserCardRemarkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PayForExtendDataRequestParams struct {
	// 卡片ICCID
	Iccid *string `json:"Iccid,omitempty" name:"Iccid"`

	// 套外流量,单位MB
	ExtentData *uint64 `json:"ExtentData,omitempty" name:"ExtentData"`

	// 应用ID
	Sdkappid *uint64 `json:"Sdkappid,omitempty" name:"Sdkappid"`
}

type PayForExtendDataRequest struct {
	*tchttp.BaseRequest
	
	// 卡片ICCID
	Iccid *string `json:"Iccid,omitempty" name:"Iccid"`

	// 套外流量,单位MB
	ExtentData *uint64 `json:"ExtentData,omitempty" name:"ExtentData"`

	// 应用ID
	Sdkappid *uint64 `json:"Sdkappid,omitempty" name:"Sdkappid"`
}

func (r *PayForExtendDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PayForExtendDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Iccid")
	delete(f, "ExtentData")
	delete(f, "Sdkappid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PayForExtendDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PayForExtendDataResponseParams struct {
	// 订单号
	Data *ResOrderIds `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type PayForExtendDataResponse struct {
	*tchttp.BaseResponse
	Response *PayForExtendDataResponseParams `json:"Response"`
}

func (r *PayForExtendDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PayForExtendDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewCardsRequestParams struct {
	// 应用ID
	Sdkappid *uint64 `json:"Sdkappid,omitempty" name:"Sdkappid"`

	// 续费的iccid
	Iccids []*string `json:"Iccids,omitempty" name:"Iccids"`

	// 续费的周期（单位：月）
	RenewNum *uint64 `json:"RenewNum,omitempty" name:"RenewNum"`
}

type RenewCardsRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	Sdkappid *uint64 `json:"Sdkappid,omitempty" name:"Sdkappid"`

	// 续费的iccid
	Iccids []*string `json:"Iccids,omitempty" name:"Iccids"`

	// 续费的周期（单位：月）
	RenewNum *uint64 `json:"RenewNum,omitempty" name:"RenewNum"`
}

func (r *RenewCardsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewCardsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Sdkappid")
	delete(f, "Iccids")
	delete(f, "RenewNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewCardsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewCardsResponseParams struct {
	// 续费成功的订单id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ResRenew `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RenewCardsResponse struct {
	*tchttp.BaseResponse
	Response *RenewCardsResponseParams `json:"Response"`
}

func (r *RenewCardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewCardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResOrderIds struct {
	// 每一张续费卡片的订单ID数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderIds []*string `json:"OrderIds,omitempty" name:"OrderIds"`
}

type ResRenew struct {
	// 每一张续费卡片的订单ID数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderIds []*string `json:"OrderIds,omitempty" name:"OrderIds"`
}

type ResSms struct {
	// 卡片ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Iccid *string `json:"Iccid,omitempty" name:"Iccid"`

	// 卡片号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msisdn *string `json:"Msisdn,omitempty" name:"Msisdn"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SdkAppid *int64 `json:"SdkAppid,omitempty" name:"SdkAppid"`

	// 短信内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`

	// 短信类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmsType *int64 `json:"SmsType,omitempty" name:"SmsType"`

	// 发送时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SendTime *string `json:"SendTime,omitempty" name:"SendTime"`

	// 推送时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportTime *string `json:"ReportTime,omitempty" name:"ReportTime"`

	// SUCC：成功  FAIL 失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 回执状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

// Predefined struct for user
type SendMultiSmsRequestParams struct {
	// 应用ID
	Sdkappid *string `json:"Sdkappid,omitempty" name:"Sdkappid"`

	// 卡片列表
	Iccids []*string `json:"Iccids,omitempty" name:"Iccids"`

	// 短信内容 长度限制 70
	Content *string `json:"Content,omitempty" name:"Content"`
}

type SendMultiSmsRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	Sdkappid *string `json:"Sdkappid,omitempty" name:"Sdkappid"`

	// 卡片列表
	Iccids []*string `json:"Iccids,omitempty" name:"Iccids"`

	// 短信内容 长度限制 70
	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *SendMultiSmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendMultiSmsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Sdkappid")
	delete(f, "Iccids")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendMultiSmsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendMultiSmsResponseParams struct {
	// 短信流水数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*SmsRet `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SendMultiSmsResponse struct {
	*tchttp.BaseResponse
	Response *SendMultiSmsResponseParams `json:"Response"`
}

func (r *SendMultiSmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendMultiSmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendSmsRequestParams struct {
	// 应用ID
	Sdkappid *int64 `json:"Sdkappid,omitempty" name:"Sdkappid"`

	// 卡片ID
	Iccid *string `json:"Iccid,omitempty" name:"Iccid"`

	// 短信内容长度70限制
	Content *string `json:"Content,omitempty" name:"Content"`
}

type SendSmsRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	Sdkappid *int64 `json:"Sdkappid,omitempty" name:"Sdkappid"`

	// 卡片ID
	Iccid *string `json:"Iccid,omitempty" name:"Iccid"`

	// 短信内容长度70限制
	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *SendSmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendSmsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Sdkappid")
	delete(f, "Iccid")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendSmsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendSmsResponseParams struct {
	// 短信流水信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *SmsSid `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SendSmsResponse struct {
	*tchttp.BaseResponse
	Response *SendSmsResponseParams `json:"Response"`
}

func (r *SendSmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendSmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SmsRet struct {
	// 该iccid请求状态
	Code *string `json:"Code,omitempty" name:"Code"`

	// 短信发送返回信息
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 卡片ID
	Iccid *string `json:"Iccid,omitempty" name:"Iccid"`

	// 流水ID
	Sid *string `json:"Sid,omitempty" name:"Sid"`
}

type SmsSid struct {
	// 卡片ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Iccid *string `json:"Iccid,omitempty" name:"Iccid"`

	// 信息流水ID
	Sid *string `json:"Sid,omitempty" name:"Sid"`
}