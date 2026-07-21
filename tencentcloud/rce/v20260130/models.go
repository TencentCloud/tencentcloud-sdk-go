// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20260130

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AssessDeviceRiskPremiumProRequestParams struct {
	// <p>用户设备指纹token标识，在您的网站或者应用程序中集成设备指纹的SDK后获取</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>客户端 IP 地址</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

type AssessDeviceRiskPremiumProRequest struct {
	*tchttp.BaseRequest
	
	// <p>用户设备指纹token标识，在您的网站或者应用程序中集成设备指纹的SDK后获取</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>客户端 IP 地址</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

func (r *AssessDeviceRiskPremiumProRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessDeviceRiskPremiumProRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceToken")
	delete(f, "UserIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssessDeviceRiskPremiumProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssessDeviceRiskPremiumProResponseParams struct {
	// <p>设备风险评估高级版返回结果</p>
	Data *AssessDeviceRiskPremiumRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssessDeviceRiskPremiumProResponse struct {
	*tchttp.BaseResponse
	Response *AssessDeviceRiskPremiumProResponseParams `json:"Response"`
}

func (r *AssessDeviceRiskPremiumProResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessDeviceRiskPremiumProResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssessDeviceRiskPremiumRsp struct {
	// <p>决策信息</p>
	Decision *Decision `json:"Decision,omitnil,omitempty" name:"Decision"`

	// <p>设备风险分信息</p>
	Score *DataScore `json:"Score,omitnil,omitempty" name:"Score"`

	// <p>设备基础信息</p>
	Device *Device `json:"Device,omitnil,omitempty" name:"Device"`

	// <p>IP环境基础信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Environment *Environment `json:"Environment,omitnil,omitempty" name:"Environment"`
}

// Predefined struct for user
type AssessDeviceRiskProRequestParams struct {
	// <p>用户设备指纹token标识，在您的网站或者应用程序中集成设备指纹的SDK后获取</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>客户端 IP 地址</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

type AssessDeviceRiskProRequest struct {
	*tchttp.BaseRequest
	
	// <p>用户设备指纹token标识，在您的网站或者应用程序中集成设备指纹的SDK后获取</p>
	DeviceToken *string `json:"DeviceToken,omitnil,omitempty" name:"DeviceToken"`

	// <p>客户端 IP 地址</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

func (r *AssessDeviceRiskProRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessDeviceRiskProRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceToken")
	delete(f, "UserIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssessDeviceRiskProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssessDeviceRiskProResponseParams struct {
	// <p>设备风险评估基础版返回结果</p>
	Data *AssessDeviceRiskRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssessDeviceRiskProResponse struct {
	*tchttp.BaseResponse
	Response *AssessDeviceRiskProResponseParams `json:"Response"`
}

func (r *AssessDeviceRiskProResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessDeviceRiskProResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssessDeviceRiskRsp struct {
	// <p>设备风险分信息</p>
	Score *DataScore `json:"Score,omitnil,omitempty" name:"Score"`

	// <p>设备基础信息</p>
	Device *Device `json:"Device,omitnil,omitempty" name:"Device"`
}

// Predefined struct for user
type AssessEnvironmentRiskRequestParams struct {
	// <p>客户端 IP 地址</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

type AssessEnvironmentRiskRequest struct {
	*tchttp.BaseRequest
	
	// <p>客户端 IP 地址</p>
	UserIp *string `json:"UserIp,omitnil,omitempty" name:"UserIp"`
}

func (r *AssessEnvironmentRiskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessEnvironmentRiskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserIp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssessEnvironmentRiskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssessEnvironmentRiskResponseParams struct {
	// <p>环境风险评估返回结果</p>
	Data *AssessEnvironmentRiskRsp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssessEnvironmentRiskResponse struct {
	*tchttp.BaseResponse
	Response *AssessEnvironmentRiskResponseParams `json:"Response"`
}

func (r *AssessEnvironmentRiskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssessEnvironmentRiskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssessEnvironmentRiskRsp struct {
	// <p>IP环境风险分信息</p>
	Score *DataScore `json:"Score,omitnil,omitempty" name:"Score"`

	// <p>IP环境基础信息</p>
	Environment *Environment `json:"Environment,omitnil,omitempty" name:"Environment"`
}

type DataScore struct {
	// <p>风险等级</p>
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// <p>风险标签</p>
	RiskLabels []*RiskLabel `json:"RiskLabels,omitnil,omitempty" name:"RiskLabels"`
}

type Decision struct {
	// <p>决策结果</p><ul><li>pass：通过</li><li>review：复审</li><li>reject：拒绝</li></ul>
	DecisionResult *string `json:"DecisionResult,omitnil,omitempty" name:"DecisionResult"`
}

type Device struct {
	// <p>设备ID</p>
	DeviceId *string `json:"DeviceId,omitnil,omitempty" name:"DeviceId"`

	// <p>App版本信息</p>
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// <p>品牌</p>
	Brand *string `json:"Brand,omitnil,omitempty" name:"Brand"`

	// <p>客户端IP</p>
	ClientIp *string `json:"ClientIp,omitnil,omitempty" name:"ClientIp"`

	// <p>机型</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>网络类型</p>
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// <p>应用包名</p>
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// <p>平台</p><p>枚举值：</p><ul><li>2： Android</li><li>3： IOS</li><li>4： H5</li><li>5： 微信小程序</li></ul>
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>系统版本</p>
	SystemVersion *string `json:"SystemVersion,omitnil,omitempty" name:"SystemVersion"`

	// <p>SDK版本</p>
	SdkBuildVersion *string `json:"SdkBuildVersion,omitnil,omitempty" name:"SdkBuildVersion"`
}

type Environment struct {
	// <p>IP地理位置信息</p>
	Location *IPLocation `json:"Location,omitnil,omitempty" name:"Location"`

	// <p>IP基础网络信息</p>
	Network *IPNetwork `json:"Network,omitnil,omitempty" name:"Network"`
}

type IPLocation struct {
	// <p>IP地址所属国家</p>
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// <p>IP地址所属省份</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>IP地址所属城市</p>
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// <p>IP地址所属地区</p>
	District *string `json:"District,omitnil,omitempty" name:"District"`

	// <p>IP地址的经度</p>
	Longitude *string `json:"Longitude,omitnil,omitempty" name:"Longitude"`

	// <p>IP地址的纬度</p>
	Latitude *string `json:"Latitude,omitnil,omitempty" name:"Latitude"`

	// <p>IP地址所属时区</p>
	Timezone *string `json:"Timezone,omitnil,omitempty" name:"Timezone"`

	// <p>IP地址的邮政编码</p>
	ZipCode *string `json:"ZipCode,omitnil,omitempty" name:"ZipCode"`
}

type IPNetwork struct {
	// <p>互联网服务提供商</p>
	ISP *string `json:"ISP,omitnil,omitempty" name:"ISP"`

	// <p>自治系统号</p>
	ASN *string `json:"ASN,omitnil,omitempty" name:"ASN"`

	// <p>IP注册组织名称</p>
	Organization *string `json:"Organization,omitnil,omitempty" name:"Organization"`

	// <p>是否保留IP</p>
	IsReserved *bool `json:"IsReserved,omitnil,omitempty" name:"IsReserved"`

	// <p>是否网关IP</p>
	IsGateway *bool `json:"IsGateway,omitnil,omitempty" name:"IsGateway"`

	// <p>是否任播网络</p>
	IsAnycast *bool `json:"IsAnycast,omitnil,omitempty" name:"IsAnycast"`

	// <p>是否移动网络</p>
	IsMobile *bool `json:"IsMobile,omitnil,omitempty" name:"IsMobile"`

	// <p>是否动态IP</p>
	IsDynamic *bool `json:"IsDynamic,omitnil,omitempty" name:"IsDynamic"`

	// <p>是否网络出口</p>
	IsEgress *bool `json:"IsEgress,omitnil,omitempty" name:"IsEgress"`

	// <p>是否域名解析</p>
	IsDNS *bool `json:"IsDNS,omitnil,omitempty" name:"IsDNS"`

	// <p>是否教育机构</p>
	IsEducation *bool `json:"IsEducation,omitnil,omitempty" name:"IsEducation"`

	// <p>是否组织机构</p>
	IsInstitution *bool `json:"IsInstitution,omitnil,omitempty" name:"IsInstitution"`

	// <p>是否企业专线</p>
	IsCompany *bool `json:"IsCompany,omitnil,omitempty" name:"IsCompany"`

	// <p>是否家用宽带</p>
	IsResidence *bool `json:"IsResidence,omitnil,omitempty" name:"IsResidence"`

	// <p>是否云服务</p>
	IsCloudService *bool `json:"IsCloudService,omitnil,omitempty" name:"IsCloudService"`

	// <p>是否基础设施</p>
	IsInfrastructure *bool `json:"IsInfrastructure,omitnil,omitempty" name:"IsInfrastructure"`

	// <p>是否邮箱服务</p>
	IsMXServer *bool `json:"IsMXServer,omitnil,omitempty" name:"IsMXServer"`
}

type RiskLabel struct {
	// <p>风险ID</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>风险描述</p>
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}