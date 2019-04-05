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

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AppInfo struct {

	// 应用id
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

	// 卡片id
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

	// 应用id
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

	// 流量池id
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

	// 订单id
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

	// 服务
	Provider *int64 `json:"Provider,omitempty" name:"Provider"`
}

type CardList struct {

	// 卡片总数
	Total *string `json:"Total,omitempty" name:"Total"`

	// 卡片列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*CardInfo `json:"List,omitempty" name:"List" list`
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

func (r *DescribeAppRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAppResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 应用信息详情
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *AppInfo `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAppResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCardRequest struct {
	*tchttp.BaseRequest

	// 应用id
	Sdkappid *int64 `json:"Sdkappid,omitempty" name:"Sdkappid"`

	// 卡片id
	Iccid *string `json:"Iccid,omitempty" name:"Iccid"`
}

func (r *DescribeCardRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCardRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCardResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 卡片详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data *CardInfo `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCardResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCardResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCardsRequest struct {
	*tchttp.BaseRequest

	// 应用id
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

func (r *DescribeCardsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCardsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 卡片列表信息
		Data *CardList `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCardsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCardsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendMultiSmsRequest struct {
	*tchttp.BaseRequest

	// 应用id
	Sdkappid *string `json:"Sdkappid,omitempty" name:"Sdkappid"`

	// 卡片列表
	Iccids []*string `json:"Iccids,omitempty" name:"Iccids" list`

	// 短信内容
	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *SendMultiSmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendMultiSmsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendMultiSmsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 短信流水数组
		Data []*SmsRet `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendMultiSmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendMultiSmsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendSmsRequest struct {
	*tchttp.BaseRequest

	// 应用id
	Sdkappid *int64 `json:"Sdkappid,omitempty" name:"Sdkappid"`

	// 卡片id
	Iccid *string `json:"Iccid,omitempty" name:"Iccid"`

	// 短信内容
	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *SendSmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendSmsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SendSmsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 短信流水信息
		Data *SmsSid `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendSmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SendSmsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SmsRet struct {

	// 该iccid请求状态
	Code *string `json:"Code,omitempty" name:"Code"`

	// 短信发送返回信息
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 卡片id
	Iccid *string `json:"Iccid,omitempty" name:"Iccid"`

	// 流水id
	Sid *string `json:"Sid,omitempty" name:"Sid"`
}

type SmsSid struct {

	// 卡片id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Iccid *string `json:"Iccid,omitempty" name:"Iccid"`

	// 信息流水id
	Sid *string `json:"Sid,omitempty" name:"Sid"`
}
