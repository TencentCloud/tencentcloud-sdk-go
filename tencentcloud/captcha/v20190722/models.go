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

package v20190722

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CaptchaOperDataInterceptUnit struct {
	// 时间
	DateKey *string `json:"DateKey,omitempty" name:"DateKey"`

	// 停止验证数量
	AllStopCnt *float64 `json:"AllStopCnt,omitempty" name:"AllStopCnt"`

	// 图片停止加载数量
	PicStopCnt *float64 `json:"PicStopCnt,omitempty" name:"PicStopCnt"`

	// 策略拦截数量
	StrategyStopCnt *float64 `json:"StrategyStopCnt,omitempty" name:"StrategyStopCnt"`
}

type CaptchaOperDataLoadTimeUnit struct {
	// 时间
	DateKey *string `json:"DateKey,omitempty" name:"DateKey"`

	// Market加载时间
	MarketLoadTime *float64 `json:"MarketLoadTime,omitempty" name:"MarketLoadTime"`

	// AppId加载时间
	AppIdLoadTime *float64 `json:"AppIdLoadTime,omitempty" name:"AppIdLoadTime"`
}

type CaptchaOperDataRes struct {
	// 验证码加载耗时数据返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperDataLoadTimeUnitArray []*CaptchaOperDataLoadTimeUnit `json:"OperDataLoadTimeUnitArray,omitempty" name:"OperDataLoadTimeUnitArray"`

	// 验证码拦截情况数据返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperDataInterceptUnitArray []*CaptchaOperDataInterceptUnit `json:"OperDataInterceptUnitArray,omitempty" name:"OperDataInterceptUnitArray"`

	// 验证码尝试次数数据返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperDataTryTimesUnitArray []*CaptchaOperDataTryTimesUnit `json:"OperDataTryTimesUnitArray,omitempty" name:"OperDataTryTimesUnitArray"`

	// 验证码尝试次数分布数据返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperDataTryTimesDistributeUnitArray []*CaptchaOperDataTryTimesDistributeUnit `json:"OperDataTryTimesDistributeUnitArray,omitempty" name:"OperDataTryTimesDistributeUnitArray"`
}

type CaptchaOperDataTryTimesDistributeUnit struct {
	// 尝试次数
	TryCount *int64 `json:"TryCount,omitempty" name:"TryCount"`

	// 用户请求数量
	UserCount *int64 `json:"UserCount,omitempty" name:"UserCount"`
}

type CaptchaOperDataTryTimesUnit struct {
	// 时间
	DateKey *string `json:"DateKey,omitempty" name:"DateKey"`

	// 平均尝试次数
	CntPerPass []*float64 `json:"CntPerPass,omitempty" name:"CntPerPass"`

	// market平均尝试次数
	MarketCntPerPass *float64 `json:"MarketCntPerPass,omitempty" name:"MarketCntPerPass"`
}

type CaptchaQueryData struct {
	// 数量
	Cnt *int64 `json:"Cnt,omitempty" name:"Cnt"`

	// 时间
	Date *string `json:"Date,omitempty" name:"Date"`
}

type CaptchaStatisticObj struct {
	// 请求总量
	ActionTotal *int64 `json:"ActionTotal,omitempty" name:"ActionTotal"`

	// 验证总量
	VerifyTotal *int64 `json:"VerifyTotal,omitempty" name:"VerifyTotal"`

	// 验证通过总量
	VerifyThroughTotal *int64 `json:"VerifyThroughTotal,omitempty" name:"VerifyThroughTotal"`

	// 验证拦截总量
	VerifyInterceptTotal *int64 `json:"VerifyInterceptTotal,omitempty" name:"VerifyInterceptTotal"`

	// 票据校验总量
	TicketTotal *int64 `json:"TicketTotal,omitempty" name:"TicketTotal"`

	// 票据通过总量
	TicketThroughTotal *int64 `json:"TicketThroughTotal,omitempty" name:"TicketThroughTotal"`

	// 票据拦截总量
	TicketInterceptTotal *int64 `json:"TicketInterceptTotal,omitempty" name:"TicketInterceptTotal"`

	// 请求趋势图
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestTrend []*RequestTrendObj `json:"RequestTrend,omitempty" name:"RequestTrend"`

	// 拦截率趋势图
	// 注意：此字段可能返回 null，表示取不到有效值。
	InterceptPerTrend []*InterceptPerTrendObj `json:"InterceptPerTrend,omitempty" name:"InterceptPerTrend"`

	// 票据校验趋势图
	// 注意：此字段可能返回 null，表示取不到有效值。
	TicketCheckTrend []*TicketCheckTrendObj `json:"TicketCheckTrend,omitempty" name:"TicketCheckTrend"`
}

type CaptchaTicketDataRes struct {
	// 票据验证总量返回
	TicketAmountArray []*TicketAmountUnit `json:"TicketAmountArray,omitempty" name:"TicketAmountArray"`

	// 票据验证通过量返回
	TicketThroughArray []*TicketThroughUnit `json:"TicketThroughArray,omitempty" name:"TicketThroughArray"`

	// 票据验证拦截量返回
	TicketInterceptArray []*TicketInterceptUnit `json:"TicketInterceptArray,omitempty" name:"TicketInterceptArray"`
}

type CaptchaUserAllAppId struct {
	// 验证码应用ID
	CaptchaAppId *int64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 注册应用名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 腾讯云APPID
	TcAppId *int64 `json:"TcAppId,omitempty" name:"TcAppId"`

	// 渠道信息
	ChannelInfo *string `json:"ChannelInfo,omitempty" name:"ChannelInfo"`
}

// Predefined struct for user
type DescribeCaptchaAppIdInfoRequestParams struct {
	// 验证码应用注册APPID
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`
}

type DescribeCaptchaAppIdInfoRequest struct {
	*tchttp.BaseRequest
	
	// 验证码应用注册APPID
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`
}

func (r *DescribeCaptchaAppIdInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaAppIdInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaptchaAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCaptchaAppIdInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaAppIdInfoResponseParams struct {
	// 界面风格
	SchemeColor *string `json:"SchemeColor,omitempty" name:"SchemeColor"`

	// 语言
	Language *int64 `json:"Language,omitempty" name:"Language"`

	// 场景
	SceneType *int64 `json:"SceneType,omitempty" name:"SceneType"`

	// 防控风险等级
	EvilInterceptGrade *int64 `json:"EvilInterceptGrade,omitempty" name:"EvilInterceptGrade"`

	// 智能验证
	SmartVerify *int64 `json:"SmartVerify,omitempty" name:"SmartVerify"`

	// 智能引擎
	SmartEngine *int64 `json:"SmartEngine,omitempty" name:"SmartEngine"`

	// 验证码类型
	CapType *int64 `json:"CapType,omitempty" name:"CapType"`

	// 应用名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 域名限制
	DomainLimit *string `json:"DomainLimit,omitempty" name:"DomainLimit"`

	// 邮件告警
	// 注意：此字段可能返回 null，表示取不到有效值。
	MailAlarm []*string `json:"MailAlarm,omitempty" name:"MailAlarm"`

	// 流量控制
	TrafficThreshold *int64 `json:"TrafficThreshold,omitempty" name:"TrafficThreshold"`

	// 加密key
	EncryptKey *string `json:"EncryptKey,omitempty" name:"EncryptKey"`

	// 是否全屏
	TopFullScreen *int64 `json:"TopFullScreen,omitempty" name:"TopFullScreen"`

	// 成功返回0 其它失败
	CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

	// 返回操作信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCaptchaAppIdInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCaptchaAppIdInfoResponseParams `json:"Response"`
}

func (r *DescribeCaptchaAppIdInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaAppIdInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaDataRequestParams struct {
	// 验证码应用ID
	CaptchaAppId *int64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 查询开始时间
	Start *int64 `json:"Start,omitempty" name:"Start"`

	// 查询结束时间
	End *int64 `json:"End,omitempty" name:"End"`

	// 查询类型
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type DescribeCaptchaDataRequest struct {
	*tchttp.BaseRequest
	
	// 验证码应用ID
	CaptchaAppId *int64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 查询开始时间
	Start *int64 `json:"Start,omitempty" name:"Start"`

	// 查询结束时间
	End *int64 `json:"End,omitempty" name:"End"`

	// 查询类型
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeCaptchaDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaptchaAppId")
	delete(f, "Start")
	delete(f, "End")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCaptchaDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaDataResponseParams struct {
	// 返回码 0 成功 其它失败
	CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

	// 数据数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*CaptchaQueryData `json:"Data,omitempty" name:"Data"`

	// 返回信息描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCaptchaDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCaptchaDataResponseParams `json:"Response"`
}

func (r *DescribeCaptchaDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaDataSumRequestParams struct {
	// 验证码应用ID
	CaptchaAppId *int64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 查询开始时间
	Start *int64 `json:"Start,omitempty" name:"Start"`

	// 查询结束时间
	End *int64 `json:"End,omitempty" name:"End"`
}

type DescribeCaptchaDataSumRequest struct {
	*tchttp.BaseRequest
	
	// 验证码应用ID
	CaptchaAppId *int64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 查询开始时间
	Start *int64 `json:"Start,omitempty" name:"Start"`

	// 查询结束时间
	End *int64 `json:"End,omitempty" name:"End"`
}

func (r *DescribeCaptchaDataSumRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaDataSumRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaptchaAppId")
	delete(f, "Start")
	delete(f, "End")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCaptchaDataSumRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaDataSumResponseParams struct {
	// 请求总量
	GetSum *int64 `json:"GetSum,omitempty" name:"GetSum"`

	// 请求验证成功量
	VfySuccSum *int64 `json:"VfySuccSum,omitempty" name:"VfySuccSum"`

	// 请求验证量
	VfySum *int64 `json:"VfySum,omitempty" name:"VfySum"`

	// 拦截攻击量
	AttackSum *int64 `json:"AttackSum,omitempty" name:"AttackSum"`

	// 返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

	// 成功返回0  其它失败
	CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

	// 票据校验总量
	CheckTicketSum *int64 `json:"CheckTicketSum,omitempty" name:"CheckTicketSum"`

	// 票据验证通过量
	TicketThroughputSum *int64 `json:"TicketThroughputSum,omitempty" name:"TicketThroughputSum"`

	// 票据验证拦截量
	TicketInterceptSum *int64 `json:"TicketInterceptSum,omitempty" name:"TicketInterceptSum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCaptchaDataSumResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCaptchaDataSumResponseParams `json:"Response"`
}

func (r *DescribeCaptchaDataSumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaDataSumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaMiniDataRequestParams struct {
	// 验证码应用ID
	CaptchaAppId *int64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 查询开始时间 例如：2019112900
	Start *int64 `json:"Start,omitempty" name:"Start"`

	// 查询结束时间 例如：2019112902
	End *int64 `json:"End,omitempty" name:"End"`

	// 查询类型 安全验证码小程序插件分类查询数据接口，请求量type=0、通过量type=1、验证量type=2、拦截量type=3 小时级查询（五小时左右延迟）
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type DescribeCaptchaMiniDataRequest struct {
	*tchttp.BaseRequest
	
	// 验证码应用ID
	CaptchaAppId *int64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 查询开始时间 例如：2019112900
	Start *int64 `json:"Start,omitempty" name:"Start"`

	// 查询结束时间 例如：2019112902
	End *int64 `json:"End,omitempty" name:"End"`

	// 查询类型 安全验证码小程序插件分类查询数据接口，请求量type=0、通过量type=1、验证量type=2、拦截量type=3 小时级查询（五小时左右延迟）
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeCaptchaMiniDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaMiniDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaptchaAppId")
	delete(f, "Start")
	delete(f, "End")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCaptchaMiniDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaMiniDataResponseParams struct {
	// 返回码 0 成功 其它失败
	CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

	// 数据数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*CaptchaQueryData `json:"Data,omitempty" name:"Data"`

	// 返回信息描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCaptchaMiniDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCaptchaMiniDataResponseParams `json:"Response"`
}

func (r *DescribeCaptchaMiniDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaMiniDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaMiniDataSumRequestParams struct {
	// 验证码应用ID
	CaptchaAppId *int64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 查询开始时间
	Start *int64 `json:"Start,omitempty" name:"Start"`

	// 查询结束时间
	End *int64 `json:"End,omitempty" name:"End"`
}

type DescribeCaptchaMiniDataSumRequest struct {
	*tchttp.BaseRequest
	
	// 验证码应用ID
	CaptchaAppId *int64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 查询开始时间
	Start *int64 `json:"Start,omitempty" name:"Start"`

	// 查询结束时间
	End *int64 `json:"End,omitempty" name:"End"`
}

func (r *DescribeCaptchaMiniDataSumRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaMiniDataSumRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaptchaAppId")
	delete(f, "Start")
	delete(f, "End")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCaptchaMiniDataSumRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaMiniDataSumResponseParams struct {
	// 请求总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	GetSum *int64 `json:"GetSum,omitempty" name:"GetSum"`

	// 请求验证成功量
	// 注意：此字段可能返回 null，表示取不到有效值。
	VfySuccSum *int64 `json:"VfySuccSum,omitempty" name:"VfySuccSum"`

	// 请求验证量
	// 注意：此字段可能返回 null，表示取不到有效值。
	VfySum *int64 `json:"VfySum,omitempty" name:"VfySum"`

	// 拦截攻击量
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttackSum *int64 `json:"AttackSum,omitempty" name:"AttackSum"`

	// 返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

	// 成功返回0  其它失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

	// 票据校验总量
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckTicketSum *int64 `json:"CheckTicketSum,omitempty" name:"CheckTicketSum"`

	// 票据验证通过量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TicketThroughputSum *int64 `json:"TicketThroughputSum,omitempty" name:"TicketThroughputSum"`

	// 票据验证拦截量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TicketInterceptSum *int64 `json:"TicketInterceptSum,omitempty" name:"TicketInterceptSum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCaptchaMiniDataSumResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCaptchaMiniDataSumResponseParams `json:"Response"`
}

func (r *DescribeCaptchaMiniDataSumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaMiniDataSumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaMiniOperDataRequestParams struct {
	// 验证码应用ID
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 查询开始时间
	Start *uint64 `json:"Start,omitempty" name:"Start"`

	// 查询类型
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 查询结束时间
	End *uint64 `json:"End,omitempty" name:"End"`
}

type DescribeCaptchaMiniOperDataRequest struct {
	*tchttp.BaseRequest
	
	// 验证码应用ID
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 查询开始时间
	Start *uint64 `json:"Start,omitempty" name:"Start"`

	// 查询类型
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 查询结束时间
	End *uint64 `json:"End,omitempty" name:"End"`
}

func (r *DescribeCaptchaMiniOperDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaMiniOperDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaptchaAppId")
	delete(f, "Start")
	delete(f, "Type")
	delete(f, "End")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCaptchaMiniOperDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaMiniOperDataResponseParams struct {
	// 成功返回 0 其它失败
	CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

	// 返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

	// 用户操作数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CaptchaOperDataRes `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCaptchaMiniOperDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCaptchaMiniOperDataResponseParams `json:"Response"`
}

func (r *DescribeCaptchaMiniOperDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaMiniOperDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaMiniResultRequestParams struct {
	// 固定填值：9（滑块验证码）
	CaptchaType *uint64 `json:"CaptchaType,omitempty" name:"CaptchaType"`

	// 验证码返回给用户的票据
	Ticket *string `json:"Ticket,omitempty" name:"Ticket"`

	// 业务侧获取到的验证码使用者的外网IP
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 验证码应用ID。登录 [验证码控制台](https://console.cloud.tencent.com/captcha/graphical)，在验证列表的【密钥】列，即可查看到CaptchaAppId。
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 验证码应用密钥。登录 [验证码控制台](https://console.cloud.tencent.com/captcha/graphical)，在验证列表的【密钥】列，即可查看到AppSecretKey。AppSecretKey属于服务器端校验验证码票据的密钥，请妥善保密，请勿泄露给第三方。
	AppSecretKey *string `json:"AppSecretKey,omitempty" name:"AppSecretKey"`

	// 业务 ID，网站或应用在多个业务中使用此服务，通过此 ID 区分统计数据
	BusinessId *uint64 `json:"BusinessId,omitempty" name:"BusinessId"`

	// 场景 ID，网站或应用的业务下有多个场景使用此服务，通过此 ID 区分统计数据
	SceneId *uint64 `json:"SceneId,omitempty" name:"SceneId"`

	// mac 地址或设备唯一标识
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// 手机设备号
	Imei *string `json:"Imei,omitempty" name:"Imei"`
}

type DescribeCaptchaMiniResultRequest struct {
	*tchttp.BaseRequest
	
	// 固定填值：9（滑块验证码）
	CaptchaType *uint64 `json:"CaptchaType,omitempty" name:"CaptchaType"`

	// 验证码返回给用户的票据
	Ticket *string `json:"Ticket,omitempty" name:"Ticket"`

	// 业务侧获取到的验证码使用者的外网IP
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 验证码应用ID。登录 [验证码控制台](https://console.cloud.tencent.com/captcha/graphical)，在验证列表的【密钥】列，即可查看到CaptchaAppId。
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 验证码应用密钥。登录 [验证码控制台](https://console.cloud.tencent.com/captcha/graphical)，在验证列表的【密钥】列，即可查看到AppSecretKey。AppSecretKey属于服务器端校验验证码票据的密钥，请妥善保密，请勿泄露给第三方。
	AppSecretKey *string `json:"AppSecretKey,omitempty" name:"AppSecretKey"`

	// 业务 ID，网站或应用在多个业务中使用此服务，通过此 ID 区分统计数据
	BusinessId *uint64 `json:"BusinessId,omitempty" name:"BusinessId"`

	// 场景 ID，网站或应用的业务下有多个场景使用此服务，通过此 ID 区分统计数据
	SceneId *uint64 `json:"SceneId,omitempty" name:"SceneId"`

	// mac 地址或设备唯一标识
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// 手机设备号
	Imei *string `json:"Imei,omitempty" name:"Imei"`
}

func (r *DescribeCaptchaMiniResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaMiniResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaptchaType")
	delete(f, "Ticket")
	delete(f, "UserIp")
	delete(f, "CaptchaAppId")
	delete(f, "AppSecretKey")
	delete(f, "BusinessId")
	delete(f, "SceneId")
	delete(f, "MacAddress")
	delete(f, "Imei")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCaptchaMiniResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaMiniResultResponseParams struct {
	// 1       ticket verification succeeded     票据验证成功
	// 7       CaptchaAppId does not match     票据与验证码应用APPID不匹配
	// 8       ticket expired     票据超时
	// 10     ticket format error     票据格式不正确
	// 15     ticket decryption failed     票据解密失败
	// 16     CaptchaAppId wrong format     检查验证码应用APPID错误
	// 21     ticket error     票据验证错误
	// 25     invalid ticket     无效票据
	// 26     system internal error     系统内部错误
	// 31 	   UnauthorizedOperation.Unauthorized	无有效套餐包/账户已欠费
	// 100   param err     参数校验错误
	CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

	// 状态描述及验证错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCaptchaMiniResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCaptchaMiniResultResponseParams `json:"Response"`
}

func (r *DescribeCaptchaMiniResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaMiniResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaMiniRiskResultRequestParams struct {
	// 固定填值：9（滑块验证码）
	CaptchaType *uint64 `json:"CaptchaType,omitempty" name:"CaptchaType"`

	// 验证码返回给用户的票据
	Ticket *string `json:"Ticket,omitempty" name:"Ticket"`

	// 业务侧获取到的验证码使用者的外网IP
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 验证码应用APPID
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 用于服务器端校验验证码票据的验证密钥，请妥善保密，请勿泄露给第三方
	AppSecretKey *string `json:"AppSecretKey,omitempty" name:"AppSecretKey"`

	// 业务 ID，网站或应用在多个业务中使用此服务，通过此 ID 区分统计数据
	BusinessId *uint64 `json:"BusinessId,omitempty" name:"BusinessId"`

	// 场景 ID，网站或应用的业务下有多个场景使用此服务，通过此 ID 区分统计数据
	SceneId *uint64 `json:"SceneId,omitempty" name:"SceneId"`

	// mac 地址或设备唯一标识
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// 手机设备号
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// 验证场景：1 活动防刷场景，2 登录保护场景，3 注册保护场景。根据需求选择场景参数。
	SceneCode *int64 `json:"SceneCode,omitempty" name:"SceneCode"`

	// 用户操作来源的微信开放账号
	WeChatOpenId *string `json:"WeChatOpenId,omitempty" name:"WeChatOpenId"`
}

type DescribeCaptchaMiniRiskResultRequest struct {
	*tchttp.BaseRequest
	
	// 固定填值：9（滑块验证码）
	CaptchaType *uint64 `json:"CaptchaType,omitempty" name:"CaptchaType"`

	// 验证码返回给用户的票据
	Ticket *string `json:"Ticket,omitempty" name:"Ticket"`

	// 业务侧获取到的验证码使用者的外网IP
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 验证码应用APPID
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 用于服务器端校验验证码票据的验证密钥，请妥善保密，请勿泄露给第三方
	AppSecretKey *string `json:"AppSecretKey,omitempty" name:"AppSecretKey"`

	// 业务 ID，网站或应用在多个业务中使用此服务，通过此 ID 区分统计数据
	BusinessId *uint64 `json:"BusinessId,omitempty" name:"BusinessId"`

	// 场景 ID，网站或应用的业务下有多个场景使用此服务，通过此 ID 区分统计数据
	SceneId *uint64 `json:"SceneId,omitempty" name:"SceneId"`

	// mac 地址或设备唯一标识
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// 手机设备号
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// 验证场景：1 活动防刷场景，2 登录保护场景，3 注册保护场景。根据需求选择场景参数。
	SceneCode *int64 `json:"SceneCode,omitempty" name:"SceneCode"`

	// 用户操作来源的微信开放账号
	WeChatOpenId *string `json:"WeChatOpenId,omitempty" name:"WeChatOpenId"`
}

func (r *DescribeCaptchaMiniRiskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaMiniRiskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaptchaType")
	delete(f, "Ticket")
	delete(f, "UserIp")
	delete(f, "CaptchaAppId")
	delete(f, "AppSecretKey")
	delete(f, "BusinessId")
	delete(f, "SceneId")
	delete(f, "MacAddress")
	delete(f, "Imei")
	delete(f, "SceneCode")
	delete(f, "WeChatOpenId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCaptchaMiniRiskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaMiniRiskResultResponseParams struct {
	// 1 ticket verification succeeded 票据验证成功
	// 7 CaptchaAppId does not match 票据与验证码应用APPID不匹配
	// 8 ticket expired 票据超时
	// 10 ticket format error 票据格式不正确
	// 15 ticket decryption failed 票据解密失败
	// 16 CaptchaAppId wrong format 检查验证码应用APPID错误
	// 21 ticket error 票据验证错误
	// 25 bad visitor 策略拦截
	// 26 system internal error 系统内部错误
	// 100 param err 参数校验错误
	CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

	// 状态描述及验证错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

	// 拦截策略返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManageMarketingRiskValue *OutputManageMarketingRiskValue `json:"ManageMarketingRiskValue,omitempty" name:"ManageMarketingRiskValue"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCaptchaMiniRiskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCaptchaMiniRiskResultResponseParams `json:"Response"`
}

func (r *DescribeCaptchaMiniRiskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaMiniRiskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaOperDataRequestParams struct {
	// 验证码应用ID
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 查询开始时间
	Start *uint64 `json:"Start,omitempty" name:"Start"`

	// 查询类型
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 查询结束时间
	End *uint64 `json:"End,omitempty" name:"End"`
}

type DescribeCaptchaOperDataRequest struct {
	*tchttp.BaseRequest
	
	// 验证码应用ID
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 查询开始时间
	Start *uint64 `json:"Start,omitempty" name:"Start"`

	// 查询类型
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 查询结束时间
	End *uint64 `json:"End,omitempty" name:"End"`
}

func (r *DescribeCaptchaOperDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaOperDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaptchaAppId")
	delete(f, "Start")
	delete(f, "Type")
	delete(f, "End")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCaptchaOperDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaOperDataResponseParams struct {
	// 成功返回 0 其它失败
	CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

	// 返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

	// 用户操作数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CaptchaOperDataRes `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCaptchaOperDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCaptchaOperDataResponseParams `json:"Response"`
}

func (r *DescribeCaptchaOperDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaOperDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaResultRequestParams struct {
	// 固定填值：9。可在控制台配置不同验证码类型。
	CaptchaType *uint64 `json:"CaptchaType,omitempty" name:"CaptchaType"`

	// 前端回调函数返回的用户验证票据
	Ticket *string `json:"Ticket,omitempty" name:"Ticket"`

	// 业务侧获取到的验证码使用者的外网IP
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 前端回调函数返回的随机字符串
	Randstr *string `json:"Randstr,omitempty" name:"Randstr"`

	// 验证码应用ID。登录 [验证码控制台](https://console.cloud.tencent.com/captcha/graphical)，在验证列表的【密钥】列，即可查看到CaptchaAppId。
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 验证码应用密钥。登录 [验证码控制台](https://console.cloud.tencent.com/captcha/graphical)，在验证列表的【密钥】列，即可查看到AppSecretKey。AppSecretKey属于服务器端校验验证码票据的密钥，请妥善保密，请勿泄露给第三方。
	AppSecretKey *string `json:"AppSecretKey,omitempty" name:"AppSecretKey"`

	// 业务 ID，网站或应用在多个业务中使用此服务，通过此 ID 区分统计数据
	BusinessId *uint64 `json:"BusinessId,omitempty" name:"BusinessId"`

	// 场景 ID，网站或应用的业务下有多个场景使用此服务，通过此 ID 区分统计数据
	SceneId *uint64 `json:"SceneId,omitempty" name:"SceneId"`

	// mac 地址或设备唯一标识
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// 手机设备号
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// 是否返回前端获取验证码时间，取值1：需要返回
	NeedGetCaptchaTime *int64 `json:"NeedGetCaptchaTime,omitempty" name:"NeedGetCaptchaTime"`
}

type DescribeCaptchaResultRequest struct {
	*tchttp.BaseRequest
	
	// 固定填值：9。可在控制台配置不同验证码类型。
	CaptchaType *uint64 `json:"CaptchaType,omitempty" name:"CaptchaType"`

	// 前端回调函数返回的用户验证票据
	Ticket *string `json:"Ticket,omitempty" name:"Ticket"`

	// 业务侧获取到的验证码使用者的外网IP
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 前端回调函数返回的随机字符串
	Randstr *string `json:"Randstr,omitempty" name:"Randstr"`

	// 验证码应用ID。登录 [验证码控制台](https://console.cloud.tencent.com/captcha/graphical)，在验证列表的【密钥】列，即可查看到CaptchaAppId。
	CaptchaAppId *uint64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 验证码应用密钥。登录 [验证码控制台](https://console.cloud.tencent.com/captcha/graphical)，在验证列表的【密钥】列，即可查看到AppSecretKey。AppSecretKey属于服务器端校验验证码票据的密钥，请妥善保密，请勿泄露给第三方。
	AppSecretKey *string `json:"AppSecretKey,omitempty" name:"AppSecretKey"`

	// 业务 ID，网站或应用在多个业务中使用此服务，通过此 ID 区分统计数据
	BusinessId *uint64 `json:"BusinessId,omitempty" name:"BusinessId"`

	// 场景 ID，网站或应用的业务下有多个场景使用此服务，通过此 ID 区分统计数据
	SceneId *uint64 `json:"SceneId,omitempty" name:"SceneId"`

	// mac 地址或设备唯一标识
	MacAddress *string `json:"MacAddress,omitempty" name:"MacAddress"`

	// 手机设备号
	Imei *string `json:"Imei,omitempty" name:"Imei"`

	// 是否返回前端获取验证码时间，取值1：需要返回
	NeedGetCaptchaTime *int64 `json:"NeedGetCaptchaTime,omitempty" name:"NeedGetCaptchaTime"`
}

func (r *DescribeCaptchaResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaptchaType")
	delete(f, "Ticket")
	delete(f, "UserIp")
	delete(f, "Randstr")
	delete(f, "CaptchaAppId")
	delete(f, "AppSecretKey")
	delete(f, "BusinessId")
	delete(f, "SceneId")
	delete(f, "MacAddress")
	delete(f, "Imei")
	delete(f, "NeedGetCaptchaTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCaptchaResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaResultResponseParams struct {
	// 1 OK 验证通过
	// 7 captcha no match 传入的Randstr不合法，请检查Randstr是否与前端返回的Randstr一致
	// 8 ticket expired 传入的Ticket已过期（Ticket有效期5分钟），请重新生成Ticket、Randstr进行校验
	// 9 ticket reused 传入的Ticket被重复使用，请重新生成Ticket、Randstr进行校验
	// 15 decrypt fail 传入的Ticket不合法，请检查Ticket是否与前端返回的Ticket一致
	// 16 appid-ticket mismatch 传入的CaptchaAppId错误，请检查CaptchaAppId是否与前端传入的CaptchaAppId一致，并且保障CaptchaAppId是从验证码控制台【验证管理】->【基础配置】中获取
	// 21 diff 票据校验异常，可能的原因是（1）若Ticket包含terror前缀，一般是由于用户网络较差，导致前端自动容灾，而生成了容灾票据，业务侧可根据需要进行跳过或二次处理。（2）若Ticket不包含terror前缀，则是由于验证码风控系统发现请求有安全风险，业务侧可根据需要进行拦截。
	// 100 appid-secretkey-ticket mismatch 参数校验错误，（1）请检查CaptchaAppId与AppSecretKey是否正确，CaptchaAppId、AppSecretKey需要在验证码控制台【验证管理】>【基础配置】中获取（2）请检查传入的Ticket是否由传入的CaptchaAppId生成
	CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

	// 状态描述及验证错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

	// 无感验证模式下，该参数返回验证结果：
	// EvilLevel=0 请求无恶意
	// EvilLevel=100 请求有恶意
	// 注意：此字段可能返回 null，表示取不到有效值。
	EvilLevel *int64 `json:"EvilLevel,omitempty" name:"EvilLevel"`

	// 前端获取验证码时间，时间戳格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	GetCaptchaTime *int64 `json:"GetCaptchaTime,omitempty" name:"GetCaptchaTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCaptchaResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCaptchaResultResponseParams `json:"Response"`
}

func (r *DescribeCaptchaResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaTicketDataRequestParams struct {
	// 验证码应用ID
	CaptchaAppId *int64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 查询开始时间 例如：20200909
	Start *int64 `json:"Start,omitempty" name:"Start"`

	// 查询结束时间 例如：20220314
	End *int64 `json:"End,omitempty" name:"End"`
}

type DescribeCaptchaTicketDataRequest struct {
	*tchttp.BaseRequest
	
	// 验证码应用ID
	CaptchaAppId *int64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 查询开始时间 例如：20200909
	Start *int64 `json:"Start,omitempty" name:"Start"`

	// 查询结束时间 例如：20220314
	End *int64 `json:"End,omitempty" name:"End"`
}

func (r *DescribeCaptchaTicketDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaTicketDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaptchaAppId")
	delete(f, "Start")
	delete(f, "End")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCaptchaTicketDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaTicketDataResponseParams struct {
	// 成功返回 0 其它失败
	CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

	// 返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

	// 验证码票据信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CaptchaTicketDataRes `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCaptchaTicketDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCaptchaTicketDataResponseParams `json:"Response"`
}

func (r *DescribeCaptchaTicketDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaTicketDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaUserAllAppIdRequestParams struct {

}

type DescribeCaptchaUserAllAppIdRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeCaptchaUserAllAppIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaUserAllAppIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCaptchaUserAllAppIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaptchaUserAllAppIdResponseParams struct {
	// 用户注册的所有Appid和应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*CaptchaUserAllAppId `json:"Data,omitempty" name:"Data"`

	// 成功返回 0  其它失败
	CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

	// 返回操作信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCaptchaUserAllAppIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCaptchaUserAllAppIdResponseParams `json:"Response"`
}

func (r *DescribeCaptchaUserAllAppIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaptchaUserAllAppIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRequestStatisticsRequestParams struct {
	// 验证码AppId
	CaptchaAppId *string `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 开始时间字符串
	StartTimeStr *string `json:"StartTimeStr,omitempty" name:"StartTimeStr"`

	// 结束时间字符串
	EndTimeStr *string `json:"EndTimeStr,omitempty" name:"EndTimeStr"`

	// 查询粒度
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
}

type GetRequestStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 验证码AppId
	CaptchaAppId *string `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 开始时间字符串
	StartTimeStr *string `json:"StartTimeStr,omitempty" name:"StartTimeStr"`

	// 结束时间字符串
	EndTimeStr *string `json:"EndTimeStr,omitempty" name:"EndTimeStr"`

	// 查询粒度
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
}

func (r *GetRequestStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRequestStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaptchaAppId")
	delete(f, "StartTimeStr")
	delete(f, "EndTimeStr")
	delete(f, "Dimension")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetRequestStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetRequestStatisticsResponseParams struct {
	// 查询后数据块
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CaptchaStatisticObj `json:"Data,omitempty" name:"Data"`

	// 验证码返回码
	CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

	// 验证码返回信息
	CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetRequestStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *GetRequestStatisticsResponseParams `json:"Response"`
}

func (r *GetRequestStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetRequestStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTicketStatisticsRequestParams struct {
	// 验证码AppId
	CaptchaAppId *string `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 开始时间字符串
	StartTimeStr *string `json:"StartTimeStr,omitempty" name:"StartTimeStr"`

	// 结束时间字符串
	EndTimeStr *string `json:"EndTimeStr,omitempty" name:"EndTimeStr"`

	// 查询粒度
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
}

type GetTicketStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 验证码AppId
	CaptchaAppId *string `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 开始时间字符串
	StartTimeStr *string `json:"StartTimeStr,omitempty" name:"StartTimeStr"`

	// 结束时间字符串
	EndTimeStr *string `json:"EndTimeStr,omitempty" name:"EndTimeStr"`

	// 查询粒度
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
}

func (r *GetTicketStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTicketStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaptchaAppId")
	delete(f, "StartTimeStr")
	delete(f, "EndTimeStr")
	delete(f, "Dimension")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTicketStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTicketStatisticsResponseParams struct {
	// 查询后数据块
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CaptchaStatisticObj `json:"Data,omitempty" name:"Data"`

	// 验证码返回码
	CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

	// 验证码返回信息
	CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTicketStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *GetTicketStatisticsResponseParams `json:"Response"`
}

func (r *GetTicketStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTicketStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTotalRequestStatisticsRequestParams struct {
	// 开始时间字符串
	StartTimeStr *string `json:"StartTimeStr,omitempty" name:"StartTimeStr"`

	// 结束时间字符串
	EndTimeStr *string `json:"EndTimeStr,omitempty" name:"EndTimeStr"`

	// 查询粒度
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
}

type GetTotalRequestStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间字符串
	StartTimeStr *string `json:"StartTimeStr,omitempty" name:"StartTimeStr"`

	// 结束时间字符串
	EndTimeStr *string `json:"EndTimeStr,omitempty" name:"EndTimeStr"`

	// 查询粒度
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
}

func (r *GetTotalRequestStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTotalRequestStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTimeStr")
	delete(f, "EndTimeStr")
	delete(f, "Dimension")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTotalRequestStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTotalRequestStatisticsResponseParams struct {
	// 查询后数据块
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CaptchaStatisticObj `json:"Data,omitempty" name:"Data"`

	// 验证码返回码
	CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

	// 验证码返回信息
	CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTotalRequestStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *GetTotalRequestStatisticsResponseParams `json:"Response"`
}

func (r *GetTotalRequestStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTotalRequestStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTotalTicketStatisticsRequestParams struct {
	// 开始时间
	StartTimeStr *string `json:"StartTimeStr,omitempty" name:"StartTimeStr"`

	// 结束时间
	EndTimeStr *string `json:"EndTimeStr,omitempty" name:"EndTimeStr"`

	// 查询粒度
	// 分钟：“1”
	// 小时：“2”
	// 天：“3”
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
}

type GetTotalTicketStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTimeStr *string `json:"StartTimeStr,omitempty" name:"StartTimeStr"`

	// 结束时间
	EndTimeStr *string `json:"EndTimeStr,omitempty" name:"EndTimeStr"`

	// 查询粒度
	// 分钟：“1”
	// 小时：“2”
	// 天：“3”
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
}

func (r *GetTotalTicketStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTotalTicketStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTimeStr")
	delete(f, "EndTimeStr")
	delete(f, "Dimension")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTotalTicketStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTotalTicketStatisticsResponseParams struct {
	// 返回数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CaptchaStatisticObj `json:"Data,omitempty" name:"Data"`

	// 返回码
	CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

	// 返回信息
	CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetTotalTicketStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *GetTotalTicketStatisticsResponseParams `json:"Response"`
}

func (r *GetTotalTicketStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTotalTicketStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InterceptPerTrendObj struct {
	// 时间参数
	Ftime *string `json:"Ftime,omitempty" name:"Ftime"`

	// 拦截率
	RequestInterceptPer *float64 `json:"RequestInterceptPer,omitempty" name:"RequestInterceptPer"`

	// 答案拦截率
	AnswerInterceptPer *float64 `json:"AnswerInterceptPer,omitempty" name:"AnswerInterceptPer"`

	// 策略拦截率
	PolicyInterceptPer *float64 `json:"PolicyInterceptPer,omitempty" name:"PolicyInterceptPer"`
}

type OutputManageMarketingRiskValue struct {
	// 账号 ID。对应输入参数： AccountType 是 1 时，对应 QQ 的 OpenID。
	// AccountType 是 2 时，对应微信的 OpenID/UnionID。
	// AccountType 是 4 时，对应手机号。
	// AccountType 是 8 时，对应 imei、idfa、imeiMD5 或者 idfaMD5。
	// AccountType 是 0 时，对应账号信息。
	// AccountType 是 10004 时，对应手机号的 MD5。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 操作时间戳，单位秒（对应输入参数）。 
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostTime *int64 `json:"PostTime,omitempty" name:"PostTime"`

	// 对应输入参数，AccountType 是 QQ 或微信开放账号时，用于标识 QQ 或微信用户登录 后关联业务自身的账号 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociateAccount *string `json:"AssociateAccount,omitempty" name:"AssociateAccount"`

	// 业务详情。 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 风险值 pass : 无恶意
	// review：需要人工审核
	// reject：拒绝，高风险恶意
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskLevel *string `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 风险类型，请查看下面详细说明 注意：此字段可能返回 null，表示取不到有效值。
	// 账号风险	
	//         账号信用低	1	账号近期存在因恶意被处罚历史，网络低活跃，被举报等因素
	// 	疑似 低活跃账号	11	账号活跃度与正常用户有差异
	// 	垃圾账号	2	疑似批量注册小号，近期存在严重违规或大量举报
	// 	疑似小号	21	账号有疑似线上养号，小号等行为
	// 	疑似 违规账号	22	账号曾有违规行为、曾被举报过、曾因违规被处罚过等
	// 	无效账号	3	送检账号参数无法成功解析，请检查微信 openid 是否有
	// 	黑名单	4	该账号在业务侧有过拉黑记录
	// 	白名单 	5	业务自行有添加过白名单记录
	// 行为风险	
	//         批量操作	101	存在 ip/设备/环境等因素的聚集性异常
	// 	疑似 IP 属性聚集 	1011	出现 IP 聚集
	// 	疑似 设备属性聚集 	1012	出现设备聚集
	// 	自动机 	103	疑似自动机批量请求
	// 	微信登录态无效 	104	检查 wxtoken 参数，是否已经失效
	// 环境风险	
	//         环境异常 	201	操作 ip/设备/环境存在异常。当前 ip 为非常用 ip 或恶意 ip 段
	// 	疑似 非常用IP请求 	2011	当前请求 IP 非该账号常用 IP
	// 	疑似 IP 异常 	2012	使用 idc 机房 ip 或 使用代理 ip 或 使用恶意 ip 
	// 	非公网有效 ip 	205	传进来的 IP 地址为内网 ip 地址或者 ip 保留地
	// 注意：此字段可能返回 null，表示取不到有效值。
	RiskType []*int64 `json:"RiskType,omitempty" name:"RiskType"`
}

type RequestTrendObj struct {
	// 时间参数
	Ftime *string `json:"Ftime,omitempty" name:"Ftime"`

	// 请求量
	RequestAction *int64 `json:"RequestAction,omitempty" name:"RequestAction"`

	// 验证量
	RequestVerify *int64 `json:"RequestVerify,omitempty" name:"RequestVerify"`

	// 通过量
	RequestThroughput *int64 `json:"RequestThroughput,omitempty" name:"RequestThroughput"`

	// 拦截量
	RequestIntercept *uint64 `json:"RequestIntercept,omitempty" name:"RequestIntercept"`
}

type TicketAmountUnit struct {
	// 时间
	DateKey *string `json:"DateKey,omitempty" name:"DateKey"`

	// 票据验证总量
	Amount *int64 `json:"Amount,omitempty" name:"Amount"`
}

type TicketCheckTrendObj struct {
	// 时间参数
	Ftime *string `json:"Ftime,omitempty" name:"Ftime"`

	// 票据校验量
	TicketCount *int64 `json:"TicketCount,omitempty" name:"TicketCount"`

	// 票据通过量
	TicketThroughput *int64 `json:"TicketThroughput,omitempty" name:"TicketThroughput"`

	// 票据拦截量
	TicketIntercept *int64 `json:"TicketIntercept,omitempty" name:"TicketIntercept"`
}

type TicketInterceptUnit struct {
	// 时间
	DateKey *string `json:"DateKey,omitempty" name:"DateKey"`

	// 票据验证拦截量
	Intercept *int64 `json:"Intercept,omitempty" name:"Intercept"`
}

type TicketThroughUnit struct {
	// 时间
	DateKey *string `json:"DateKey,omitempty" name:"DateKey"`

	// 票据验证的通过量
	Through *int64 `json:"Through,omitempty" name:"Through"`
}

// Predefined struct for user
type UpdateCaptchaAppIdInfoRequestParams struct {
	// 验证码应用ID
	CaptchaAppId *int64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 应用名
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 域名限制
	DomainLimit *string `json:"DomainLimit,omitempty" name:"DomainLimit"`

	// 场景类型
	SceneType *int64 `json:"SceneType,omitempty" name:"SceneType"`

	// 验证码类型
	CapType *int64 `json:"CapType,omitempty" name:"CapType"`

	// 风险级别
	EvilInterceptGrade *int64 `json:"EvilInterceptGrade,omitempty" name:"EvilInterceptGrade"`

	// 智能检测
	SmartVerify *int64 `json:"SmartVerify,omitempty" name:"SmartVerify"`

	// 开启智能引擎
	SmartEngine *int64 `json:"SmartEngine,omitempty" name:"SmartEngine"`

	// web风格
	SchemeColor *string `json:"SchemeColor,omitempty" name:"SchemeColor"`

	// 语言
	CaptchaLanguage *int64 `json:"CaptchaLanguage,omitempty" name:"CaptchaLanguage"`

	// 告警邮箱
	MailAlarm *string `json:"MailAlarm,omitempty" name:"MailAlarm"`

	// 是否全屏
	TopFullScreen *int64 `json:"TopFullScreen,omitempty" name:"TopFullScreen"`

	// 流量限制
	TrafficThreshold *int64 `json:"TrafficThreshold,omitempty" name:"TrafficThreshold"`
}

type UpdateCaptchaAppIdInfoRequest struct {
	*tchttp.BaseRequest
	
	// 验证码应用ID
	CaptchaAppId *int64 `json:"CaptchaAppId,omitempty" name:"CaptchaAppId"`

	// 应用名
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 域名限制
	DomainLimit *string `json:"DomainLimit,omitempty" name:"DomainLimit"`

	// 场景类型
	SceneType *int64 `json:"SceneType,omitempty" name:"SceneType"`

	// 验证码类型
	CapType *int64 `json:"CapType,omitempty" name:"CapType"`

	// 风险级别
	EvilInterceptGrade *int64 `json:"EvilInterceptGrade,omitempty" name:"EvilInterceptGrade"`

	// 智能检测
	SmartVerify *int64 `json:"SmartVerify,omitempty" name:"SmartVerify"`

	// 开启智能引擎
	SmartEngine *int64 `json:"SmartEngine,omitempty" name:"SmartEngine"`

	// web风格
	SchemeColor *string `json:"SchemeColor,omitempty" name:"SchemeColor"`

	// 语言
	CaptchaLanguage *int64 `json:"CaptchaLanguage,omitempty" name:"CaptchaLanguage"`

	// 告警邮箱
	MailAlarm *string `json:"MailAlarm,omitempty" name:"MailAlarm"`

	// 是否全屏
	TopFullScreen *int64 `json:"TopFullScreen,omitempty" name:"TopFullScreen"`

	// 流量限制
	TrafficThreshold *int64 `json:"TrafficThreshold,omitempty" name:"TrafficThreshold"`
}

func (r *UpdateCaptchaAppIdInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCaptchaAppIdInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaptchaAppId")
	delete(f, "AppName")
	delete(f, "DomainLimit")
	delete(f, "SceneType")
	delete(f, "CapType")
	delete(f, "EvilInterceptGrade")
	delete(f, "SmartVerify")
	delete(f, "SmartEngine")
	delete(f, "SchemeColor")
	delete(f, "CaptchaLanguage")
	delete(f, "MailAlarm")
	delete(f, "TopFullScreen")
	delete(f, "TrafficThreshold")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCaptchaAppIdInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCaptchaAppIdInfoResponseParams struct {
	// 返回码 0 成功，其它失败
	CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

	// 返回操作信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateCaptchaAppIdInfoResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCaptchaAppIdInfoResponseParams `json:"Response"`
}

func (r *UpdateCaptchaAppIdInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCaptchaAppIdInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}