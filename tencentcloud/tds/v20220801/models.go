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

package v20220801

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type DescribeFraudBaseRequestParams struct {
	// 客户端通过SDK获取的设备Token
	DeviceToken *string `json:"DeviceToken,omitempty" name:"DeviceToken"`
}

type DescribeFraudBaseRequest struct {
	*tchttp.BaseRequest
	
	// 客户端通过SDK获取的设备Token
	DeviceToken *string `json:"DeviceToken,omitempty" name:"DeviceToken"`
}

func (r *DescribeFraudBaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFraudBaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFraudBaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFraudBaseResponseParams struct {
	// App版本信息
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// 品牌
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// 客户端IP
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 机型
	Model *string `json:"Model,omitempty" name:"Model"`

	// 网络类型
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// 应用包名
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 平台（2-Android，3-iOS，4-H5，5-微信小程序）
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 系统版本
	SystemVersion *string `json:"SystemVersion,omitempty" name:"SystemVersion"`

	// SDK版本号
	SdkBuildNo *string `json:"SdkBuildNo,omitempty" name:"SdkBuildNo"`

	// 实时风险信息
	RiskInfos []*RiskInfo `json:"RiskInfos,omitempty" name:"RiskInfos"`

	// 离线风险信息
	HistRiskInfos []*RiskInfo `json:"HistRiskInfos,omitempty" name:"HistRiskInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFraudBaseResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFraudBaseResponseParams `json:"Response"`
}

func (r *DescribeFraudBaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFraudBaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFraudPremiumRequestParams struct {
	// 客户端通过SDK获取的设备Token
	DeviceToken *string `json:"DeviceToken,omitempty" name:"DeviceToken"`
}

type DescribeFraudPremiumRequest struct {
	*tchttp.BaseRequest
	
	// 客户端通过SDK获取的设备Token
	DeviceToken *string `json:"DeviceToken,omitempty" name:"DeviceToken"`
}

func (r *DescribeFraudPremiumRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFraudPremiumRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFraudPremiumRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFraudPremiumResponseParams struct {
	// App版本信息
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// 品牌
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// 客户端IP
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 机型
	Model *string `json:"Model,omitempty" name:"Model"`

	// 网络类型
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// 应用包名
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 平台（2-Android，3-iOS，4-H5，5-微信小程序）
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 系统版本
	SystemVersion *string `json:"SystemVersion,omitempty" name:"SystemVersion"`

	// SDK版本号
	SdkBuildNo *string `json:"SdkBuildNo,omitempty" name:"SdkBuildNo"`

	// 实时风险信息
	RiskInfos []*RiskInfo `json:"RiskInfos,omitempty" name:"RiskInfos"`

	// 离线风险信息
	HistRiskInfos []*RiskInfo `json:"HistRiskInfos,omitempty" name:"HistRiskInfos"`

	// 设备匿名标识
	Openid *string `json:"Openid,omitempty" name:"Openid"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFraudPremiumResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFraudPremiumResponseParams `json:"Response"`
}

func (r *DescribeFraudPremiumResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFraudPremiumResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFraudUltimateRequestParams struct {
	// 客户端通过SDK获取的设备Token
	DeviceToken *string `json:"DeviceToken,omitempty" name:"DeviceToken"`

	// 使用场景。目前仅支持login-登录场景、register-注册场景
	SceneCode *string `json:"SceneCode,omitempty" name:"SceneCode"`

	// 用户唯一标识
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 事件时间戳（毫秒）
	EventTime *uint64 `json:"EventTime,omitempty" name:"EventTime"`

	// 事件耗时（毫秒），例如进入登录界面到点击登录按钮耗时
	ElapsedTime *uint64 `json:"ElapsedTime,omitempty" name:"ElapsedTime"`

	// 微信的OpenId
	WeChatOpenId *string `json:"WeChatOpenId,omitempty" name:"WeChatOpenId"`

	// 手机号码（注：不需要带国家代码 例如：13430421011）。可以传入原文或MD5
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 客户端IP
	ClientIP *string `json:"ClientIP,omitempty" name:"ClientIP"`

	// QQ的OpenId
	QQOpenId *string `json:"QQOpenId,omitempty" name:"QQOpenId"`
}

type DescribeFraudUltimateRequest struct {
	*tchttp.BaseRequest
	
	// 客户端通过SDK获取的设备Token
	DeviceToken *string `json:"DeviceToken,omitempty" name:"DeviceToken"`

	// 使用场景。目前仅支持login-登录场景、register-注册场景
	SceneCode *string `json:"SceneCode,omitempty" name:"SceneCode"`

	// 用户唯一标识
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 事件时间戳（毫秒）
	EventTime *uint64 `json:"EventTime,omitempty" name:"EventTime"`

	// 事件耗时（毫秒），例如进入登录界面到点击登录按钮耗时
	ElapsedTime *uint64 `json:"ElapsedTime,omitempty" name:"ElapsedTime"`

	// 微信的OpenId
	WeChatOpenId *string `json:"WeChatOpenId,omitempty" name:"WeChatOpenId"`

	// 手机号码（注：不需要带国家代码 例如：13430421011）。可以传入原文或MD5
	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`

	// 客户端IP
	ClientIP *string `json:"ClientIP,omitempty" name:"ClientIP"`

	// QQ的OpenId
	QQOpenId *string `json:"QQOpenId,omitempty" name:"QQOpenId"`
}

func (r *DescribeFraudUltimateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFraudUltimateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceToken")
	delete(f, "SceneCode")
	delete(f, "UserId")
	delete(f, "EventTime")
	delete(f, "ElapsedTime")
	delete(f, "WeChatOpenId")
	delete(f, "PhoneNumber")
	delete(f, "ClientIP")
	delete(f, "QQOpenId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFraudUltimateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFraudUltimateResponseParams struct {
	// App版本信息
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// 品牌
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// 客户端IP
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 机型
	Model *string `json:"Model,omitempty" name:"Model"`

	// 网络类型
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// 应用包名
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 平台（2-Android，3-iOS，4-H5，5-微信小程序）
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 系统版本
	SystemVersion *string `json:"SystemVersion,omitempty" name:"SystemVersion"`

	// SDK版本号
	SdkBuildNo *string `json:"SdkBuildNo,omitempty" name:"SdkBuildNo"`

	// 实时风险信息
	RiskInfos []*RiskInfo `json:"RiskInfos,omitempty" name:"RiskInfos"`

	// 离线风险信息
	HistRiskInfos []*RiskInfo `json:"HistRiskInfos,omitempty" name:"HistRiskInfos"`

	// 设备匿名标识
	Openid *string `json:"Openid,omitempty" name:"Openid"`

	// 场景风险信息
	SceneRiskInfos []*RiskInfo `json:"SceneRiskInfos,omitempty" name:"SceneRiskInfos"`

	// 建议等级。1-极差，2-较差，3-中等，4-良好，5-优秀
	SuggestionLevel *uint64 `json:"SuggestionLevel,omitempty" name:"SuggestionLevel"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFraudUltimateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFraudUltimateResponseParams `json:"Response"`
}

func (r *DescribeFraudUltimateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFraudUltimateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrustedIDRequestParams struct {
	// 客户端通过SDK获取的设备Token
	DeviceToken *string `json:"DeviceToken,omitempty" name:"DeviceToken"`
}

type DescribeTrustedIDRequest struct {
	*tchttp.BaseRequest
	
	// 客户端通过SDK获取的设备Token
	DeviceToken *string `json:"DeviceToken,omitempty" name:"DeviceToken"`
}

func (r *DescribeTrustedIDRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrustedIDRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrustedIDRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrustedIDResponseParams struct {
	// 设备匿名标识
	Openid *string `json:"Openid,omitempty" name:"Openid"`

	// App版本信息
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// 品牌
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// 客户端IP
	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`

	// 机型
	Model *string `json:"Model,omitempty" name:"Model"`

	// 网络类型
	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`

	// 应用包名
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`

	// 平台（2-Android，3-iOS，4-H5，5-微信小程序）
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 系统版本
	SystemVersion *string `json:"SystemVersion,omitempty" name:"SystemVersion"`

	// SDK版本号
	SdkBuildNo *string `json:"SdkBuildNo,omitempty" name:"SdkBuildNo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTrustedIDResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrustedIDResponseParams `json:"Response"`
}

func (r *DescribeTrustedIDResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrustedIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RiskInfo struct {
	// 风险类型
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// 风险等级
	Level *uint64 `json:"Level,omitempty" name:"Level"`
}