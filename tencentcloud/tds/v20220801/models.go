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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type DataAuthorizationInfo struct {
	// 数据委托方、需求方：客户主体名称。
	DataProviderName *string `json:"DataProviderName,omitnil,omitempty" name:"DataProviderName"`

	// 数据受托方、提供方：腾讯云主体名称。
	// 
	// 固定填：腾讯云计算（北京）有限责任公司
	DataRecipientName *string `json:"DataRecipientName,omitnil,omitempty" name:"DataRecipientName"`

	// 客户请求所涉及的用户个人信息类型，支持多选。实际以接口请求传参为准。
	// 1-手机号；
	// 2-微信开放账号；
	// 3-QQ开放账号；
	// 4-IP地址；
	UserDataType []*uint64 `json:"UserDataType,omitnil,omitempty" name:"UserDataType"`

	// 客户是否已按合规指南要求获取用户授权，同意客户委托腾讯云处理入参信息，结合已合法收集的用户数据进行必要处理得出服务结果，并返回给客户。
	// 
	// 1-已授权；其它值为未授权。
	IsAuthorize *uint64 `json:"IsAuthorize,omitnil,omitempty" name:"IsAuthorize"`

	// 客户获得的用户授权期限时间戳（单位秒）。
	// 
	// 不填或0默认无固定期限。
	AuthorizationTerm *uint64 `json:"AuthorizationTerm,omitnil,omitempty" name:"AuthorizationTerm"`

	// 客户获得用户授权所依赖的协议地址。
	PrivacyPolicyLink *string `json:"PrivacyPolicyLink,omitnil,omitempty" name:"PrivacyPolicyLink"`
}

// Predefined struct for user
type DescribeFraudBaseRequestParams struct {
	// 客户端通过SDK获取的设备Token
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`
}

type DescribeFraudBaseRequest struct {
	*tchttp.BaseRequest
	
	// 客户端通过SDK获取的设备Token
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`
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
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 客户端IP
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// 机型
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 网络类型
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// 应用包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 平台（2-Android，3-iOS，4-H5，5-微信小程序）
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 系统版本
	SystemVersion *string `json:"SystemVersion,omitnil,omitempty" name:"SystemVersion"`

	// SDK版本号
	SdkBuildNo *string `json:"SdkBuildNo,omitnil,omitempty" name:"SdkBuildNo"`

	// 实时风险信息
	RiskInfos []*RiskInfo `json:"RiskInfos,omitnil,omitempty" name:"RiskInfos"`

	// 离线风险信息
	HistRiskInfos []*RiskInfo `json:"HistRiskInfos,omitnil,omitempty" name:"HistRiskInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`
}

type DescribeFraudPremiumRequest struct {
	*tchttp.BaseRequest
	
	// 客户端通过SDK获取的设备Token
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`
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
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 客户端IP
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// 机型
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 网络类型
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// 应用包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 平台（2-Android，3-iOS，4-H5，5-微信小程序）
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 系统版本
	SystemVersion *string `json:"SystemVersion,omitnil,omitempty" name:"SystemVersion"`

	// SDK版本号
	SdkBuildNo *string `json:"SdkBuildNo,omitnil,omitempty" name:"SdkBuildNo"`

	// 实时风险信息
	RiskInfos []*RiskInfo `json:"RiskInfos,omitnil,omitempty" name:"RiskInfos"`

	// 离线风险信息
	HistRiskInfos []*RiskInfo `json:"HistRiskInfos,omitnil,omitempty" name:"HistRiskInfos"`

	// 设备匿名标识
	Openid *string `json:"Openid,omitnil,omitempty" name:"Openid"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// 使用场景。目前仅支持login-登录场景、register-注册场景
	SceneCode *string `json:"SceneCode,omitnil,omitempty" name:"SceneCode"`

	// 用户唯一标识
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 事件时间戳（毫秒）
	EventTime *uint64 `json:"EventTime,omitnil,omitempty" name:"EventTime"`

	// 事件耗时（毫秒），例如进入登录界面到点击登录按钮耗时
	ElapsedTime *uint64 `json:"ElapsedTime,omitnil,omitempty" name:"ElapsedTime"`

	// 微信的OpenId
	WeChatOpenId *string `json:"WeChatOpenId,omitnil,omitempty" name:"WeChatOpenId"`

	// 手机号码（注：不需要带国家代码 例如：13430421011）。可以传入原文或MD5
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 客户端IP
	ClientIP *string `json:"ClientIP,omitnil,omitempty" name:"ClientIP"`

	// QQ的OpenId
	QQOpenId *string `json:"QQOpenId,omitnil,omitempty" name:"QQOpenId"`

	// 数据授权信息
	DataAuthorization *DataAuthorizationInfo `json:"DataAuthorization,omitnil,omitempty" name:"DataAuthorization"`
}

type DescribeFraudUltimateRequest struct {
	*tchttp.BaseRequest
	
	// 客户端通过SDK获取的设备Token
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// 使用场景。目前仅支持login-登录场景、register-注册场景
	SceneCode *string `json:"SceneCode,omitnil,omitempty" name:"SceneCode"`

	// 用户唯一标识
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 事件时间戳（毫秒）
	EventTime *uint64 `json:"EventTime,omitnil,omitempty" name:"EventTime"`

	// 事件耗时（毫秒），例如进入登录界面到点击登录按钮耗时
	ElapsedTime *uint64 `json:"ElapsedTime,omitnil,omitempty" name:"ElapsedTime"`

	// 微信的OpenId
	WeChatOpenId *string `json:"WeChatOpenId,omitnil,omitempty" name:"WeChatOpenId"`

	// 手机号码（注：不需要带国家代码 例如：13430421011）。可以传入原文或MD5
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 客户端IP
	ClientIP *string `json:"ClientIP,omitnil,omitempty" name:"ClientIP"`

	// QQ的OpenId
	QQOpenId *string `json:"QQOpenId,omitnil,omitempty" name:"QQOpenId"`

	// 数据授权信息
	DataAuthorization *DataAuthorizationInfo `json:"DataAuthorization,omitnil,omitempty" name:"DataAuthorization"`
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
	delete(f, "DataAuthorization")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFraudUltimateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFraudUltimateResponseParams struct {
	// App版本信息
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 客户端IP
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// 机型
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 网络类型
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// 应用包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 平台（2-Android，3-iOS，4-H5，5-微信小程序）
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 系统版本
	SystemVersion *string `json:"SystemVersion,omitnil,omitempty" name:"SystemVersion"`

	// SDK版本号
	SdkBuildNo *string `json:"SdkBuildNo,omitnil,omitempty" name:"SdkBuildNo"`

	// 实时风险信息
	RiskInfos []*RiskInfo `json:"RiskInfos,omitnil,omitempty" name:"RiskInfos"`

	// 离线风险信息
	HistRiskInfos []*RiskInfo `json:"HistRiskInfos,omitnil,omitempty" name:"HistRiskInfos"`

	// 设备匿名标识
	Openid *string `json:"Openid,omitnil,omitempty" name:"Openid"`

	// 场景风险信息
	SceneRiskInfos []*RiskInfo `json:"SceneRiskInfos,omitnil,omitempty" name:"SceneRiskInfos"`

	// 建议等级。1-极差，2-较差，3-中等，4-良好，5-优秀
	SuggestionLevel *uint64 `json:"SuggestionLevel,omitnil,omitempty" name:"SuggestionLevel"`

	// 图灵盾统一ID
	Unionid *string `json:"Unionid,omitnil,omitempty" name:"Unionid"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`
}

type DescribeTrustedIDRequest struct {
	*tchttp.BaseRequest
	
	// 客户端通过SDK获取的设备Token
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`
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
	Openid *string `json:"Openid,omitnil,omitempty" name:"Openid"`

	// App版本信息
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// 品牌
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// 客户端IP
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// 机型
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 网络类型
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// 应用包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 平台（2-Android，3-iOS，4-H5，5-微信小程序）
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 系统版本
	SystemVersion *string `json:"SystemVersion,omitnil,omitempty" name:"SystemVersion"`

	// SDK版本号
	SdkBuildNo *string `json:"SdkBuildNo,omitnil,omitempty" name:"SdkBuildNo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 风险类型。更多详情请参见：[Android](https://cloud.tencent.com/document/product/1628/85898)、[iOS](https://cloud.tencent.com/document/product/1628/85896)、[H5](https://cloud.tencent.com/document/product/1628/85897)、[小程序](https://cloud.tencent.com/document/product/1628/85895)、[场景风险](https://cloud.tencent.com/document/product/1628/88912)
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// [风险等级](https://cloud.tencent.com/document/product/1628/85308)
	Level *uint64 `json:"Level,omitnil,omitempty" name:"Level"`
}