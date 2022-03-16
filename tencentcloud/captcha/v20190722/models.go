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

type DescribeCaptchaAppIdInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeCaptchaDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeCaptchaDataSumResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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

		// 票据校验量
		CheckTicketSum *int64 `json:"CheckTicketSum,omitempty" name:"CheckTicketSum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeCaptchaMiniDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeCaptchaMiniDataSumResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeCaptchaMiniOperDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeCaptchaMiniResultRequest struct {
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

type DescribeCaptchaMiniResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 1       ticket verification succeeded     票据验证成功
	// 7       CaptchaAppId does not match     票据与验证码应用APPID不匹配
	// 8       ticket expired     票据超时
	// 10     ticket format error     票据格式不正确
	// 15     ticket decryption failed     票据解密失败
	// 16     CaptchaAppId wrong format     检查验证码应用APPID错误
	// 21     ticket error     票据验证错误
	// 25     invalid ticket     无效票据
	// 26     system internal error     系统内部错误
	// 100   param err     参数校验错误
		CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

		// 状态描述及验证错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeCaptchaMiniRiskResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeCaptchaOperDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

	// 验证码应用ID
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

type DescribeCaptchaResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 1 OK 验证通过
	// 6 user code len error 验证码长度不匹配，请检查请求是否带Randstr参数，Randstr参数大小写是否有误
	// 7 captcha no match 验证码答案不匹配/Randstr参数不匹配，请重新生成Randstr、Ticket进行校验
	// 8 verify timeout 验证码签名超时，票据已过期，请重新生成Randstr、Ticket票进行校验
	// 9 Sequnce repeat 验证码签名重放，票据重复使用，请重新生成Randstr、Ticket进行校验
	// 10 Sequnce invalid 验证码签名序列
	// 11 Cookie invalid 验证码cookie信息不合法，非法请求，可能存在不规范接入
	// 12 sig len error 签名长度错误
	// 13 verify ip no match ip不匹配，非法请求，可能存在不规范接入
	// 15 decrypt fail 验证码签名解密失败，票据校验失败，请检查Ticket票据是否与前端返回Ticket一致
	// 16 appid no match 验证码强校验appid错误，前端代码 data-appid 和后端 CaptchaAppId 所填写的值，必须和 验证码控制台 中【验证详情】>【基础配置】内的 AppID 一致,请检查CaptchaAppId是否为控制台基础配置界面系统分配的APPID
	// 17 cmd no much 验证码系统命令不匹配
	// 18 uin no match 号码不匹配
	// 19 seq redirect 重定向验证
	// 20 opt no vcode 操作使用pt免验证码校验错误
	// 21 diff 差别，验证错误 
	// 该情况出现原因一般为，当验证码前端生成terror格式票据并进行后端票据校验时，用户网络较差，该情况下仍会返回可用票据，业务侧可以自行根据需要，进行票据生成或做其他处理。详情参见 [验证码Web前端接入-异常处理文档](https://cloud.tencent.com/document/product/1110/36841#.E5.BC.82.E5.B8.B8.E5.A4.84.E7.90.86)。
	// 22 captcha type not match 验证码类型与拉取时不一致
	// 23 verify type error 验证类型错误
	// 24 invalid pkg 非法请求包
	// 25 bad visitor 策略拦截
	// 26 system busy 系统内部错误
	// 100 param err appsecretkey 参数校验错误，CaptchaAppId 与对应 AppSecretKey 不一致，需检查 AppSecretKey 参数是否有误。其中 CaptchaAppId、 AppSecretKey 在 验证码控制台 的【验证详情】>【基础配置】中获取
	// 104 Ticket Reuse 票据重复使用，同个票据验证多次，请重新生成Randstr、Ticket进行校验
		CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

		// 状态描述及验证错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

		// [0,100]，恶意等级
	// 注意：此字段可能返回 null，表示取不到有效值。
		EvilLevel *int64 `json:"EvilLevel,omitempty" name:"EvilLevel"`

		// 前端获取验证码时间，时间戳格式
	// 注意：此字段可能返回 null，表示取不到有效值。
		GetCaptchaTime *int64 `json:"GetCaptchaTime,omitempty" name:"GetCaptchaTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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

type DescribeCaptchaTicketDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type DescribeCaptchaUserAllAppIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
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

type TicketAmountUnit struct {

	// 时间
	DateKey *string `json:"DateKey,omitempty" name:"DateKey"`

	// 票据验证总量
	Amount *int64 `json:"Amount,omitempty" name:"Amount"`
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

type UpdateCaptchaAppIdInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回码 0 成功，其它失败
		CaptchaCode *int64 `json:"CaptchaCode,omitempty" name:"CaptchaCode"`

		// 返回操作信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		CaptchaMsg *string `json:"CaptchaMsg,omitempty" name:"CaptchaMsg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
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
