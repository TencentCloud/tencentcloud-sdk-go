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