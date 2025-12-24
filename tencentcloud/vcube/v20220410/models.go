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

package v20220410

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AppInfo struct {
	// 应用Id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Ios 包名
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Andorid 包名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 应用详情
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// Mac 进程名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MacBundleId *string `json:"MacBundleId,omitnil,omitempty" name:"MacBundleId"`

	// windows 进程名
	// 注意：此字段可能返回 null，表示取不到有效值。
	WinProcessName *string `json:"WinProcessName,omitnil,omitempty" name:"WinProcessName"`

	// 允许的web域名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`
}

type ApplicationInfo struct {
	// 应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Ios应用的唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 应用类型，formal： 正式应用，test：测试应用
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// license数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Licenses []*License `json:"Licenses,omitnil,omitempty" name:"Licenses"`

	// license 秘钥
	LicenseKey *string `json:"LicenseKey,omitnil,omitempty" name:"LicenseKey"`

	// 安卓应用的唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 用户应用Id
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 视立方下载license的url
	LicenseUrl *string `json:"LicenseUrl,omitnil,omitempty" name:"LicenseUrl"`

	// 优图美视信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	XMagics []*XMagicInfo `json:"XMagics,omitnil,omitempty" name:"XMagics"`

	// Mac  进程名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MacBundleId *string `json:"MacBundleId,omitnil,omitempty" name:"MacBundleId"`

	// windows 进程名
	// 注意：此字段可能返回 null，表示取不到有效值。
	WinProcessName *string `json:"WinProcessName,omitnil,omitempty" name:"WinProcessName"`

	// web端Domain列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// 账号AppId
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 扩展包名数量上限
	NameLimit *uint64 `json:"NameLimit,omitnil,omitempty" name:"NameLimit"`
}

// Predefined struct for user
type CreateActivityLicenseRequestParams struct {
	// 活动Id
	Activity *string `json:"Activity,omitnil,omitempty" name:"Activity"`
}

type CreateActivityLicenseRequest struct {
	*tchttp.BaseRequest
	
	// 活动Id
	Activity *string `json:"Activity,omitnil,omitempty" name:"Activity"`
}

func (r *CreateActivityLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateActivityLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Activity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateActivityLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateActivityLicenseResponseParams struct {
	// 用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// app名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// ios包名
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 安卓包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 有效时长
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// license秘钥
	LicenseKey *string `json:"LicenseKey,omitnil,omitempty" name:"LicenseKey"`

	// license 授权文件下载链接
	LicenseUrl *string `json:"LicenseUrl,omitnil,omitempty" name:"LicenseUrl"`

	// license剩余天数，最后一天以及过期显示0
	ResidueDay *uint64 `json:"ResidueDay,omitnil,omitempty" name:"ResidueDay"`

	// license剩余秒数
	Residue *uint64 `json:"Residue,omitnil,omitempty" name:"Residue"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateActivityLicenseResponse struct {
	*tchttp.BaseResponse
	Response *CreateActivityLicenseResponseParams `json:"Response"`
}

func (r *CreateActivityLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateActivityLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationAndBindLicenseRequestParams struct {
	// 应用名
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 应用ID
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 资源包ID
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 营业执照
	CompanyPermit *string `json:"CompanyPermit,omitnil,omitempty" name:"CompanyPermit"`

	// 公司类型
	CompanyType *string `json:"CompanyType,omitnil,omitempty" name:"CompanyType"`

	// 公司名称
	CompanyName *string `json:"CompanyName,omitnil,omitempty" name:"CompanyName"`

	// 优图资源id列表
	XMagicResourceIds []*string `json:"XMagicResourceIds,omitnil,omitempty" name:"XMagicResourceIds"`

	// Mac 进程名
	MacBundleId *string `json:"MacBundleId,omitnil,omitempty" name:"MacBundleId"`

	// Windows 进程名
	WinProcessName *string `json:"WinProcessName,omitnil,omitempty" name:"WinProcessName"`

	// 要开通的域名列表
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// 要开通的端，web/mobile/pc
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type CreateApplicationAndBindLicenseRequest struct {
	*tchttp.BaseRequest
	
	// 应用名
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 应用ID
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 资源包ID
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 营业执照
	CompanyPermit *string `json:"CompanyPermit,omitnil,omitempty" name:"CompanyPermit"`

	// 公司类型
	CompanyType *string `json:"CompanyType,omitnil,omitempty" name:"CompanyType"`

	// 公司名称
	CompanyName *string `json:"CompanyName,omitnil,omitempty" name:"CompanyName"`

	// 优图资源id列表
	XMagicResourceIds []*string `json:"XMagicResourceIds,omitnil,omitempty" name:"XMagicResourceIds"`

	// Mac 进程名
	MacBundleId *string `json:"MacBundleId,omitnil,omitempty" name:"MacBundleId"`

	// Windows 进程名
	WinProcessName *string `json:"WinProcessName,omitnil,omitempty" name:"WinProcessName"`

	// 要开通的域名列表
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`

	// 要开通的端，web/mobile/pc
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`
}

func (r *CreateApplicationAndBindLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationAndBindLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "BundleId")
	delete(f, "PackageName")
	delete(f, "ResourceIds")
	delete(f, "CompanyPermit")
	delete(f, "CompanyType")
	delete(f, "CompanyName")
	delete(f, "XMagicResourceIds")
	delete(f, "MacBundleId")
	delete(f, "WinProcessName")
	delete(f, "DomainList")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationAndBindLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationAndBindLicenseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateApplicationAndBindLicenseResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationAndBindLicenseResponseParams `json:"Response"`
}

func (r *CreateApplicationAndBindLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationAndBindLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationAndVideoRequestParams struct {
	// 应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Ios 包名
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Android 包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type CreateApplicationAndVideoRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Ios 包名
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Android 包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

func (r *CreateApplicationAndVideoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationAndVideoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "BundleId")
	delete(f, "PackageName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationAndVideoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationAndVideoResponseParams struct {
	// license唯一标识
	LicenseId *uint64 `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateApplicationAndVideoResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationAndVideoResponseParams `json:"Response"`
}

func (r *CreateApplicationAndVideoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationAndVideoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationAndWebPlayerLicenseRequestParams struct {
	// 应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 域名列表
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`
}

type CreateApplicationAndWebPlayerLicenseRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 域名列表
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`
}

func (r *CreateApplicationAndWebPlayerLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationAndWebPlayerLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "DomainList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApplicationAndWebPlayerLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApplicationAndWebPlayerLicenseResponseParams struct {
	// license的唯一标识
	LicenseId *uint64 `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateApplicationAndWebPlayerLicenseResponse struct {
	*tchttp.BaseResponse
	Response *CreateApplicationAndWebPlayerLicenseResponseParams `json:"Response"`
}

func (r *CreateApplicationAndWebPlayerLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApplicationAndWebPlayerLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLicenseRequestParams struct {
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 资源包ID
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// url地址分组
	Group *uint64 `json:"Group,omitnil,omitempty" name:"Group"`
}

type CreateLicenseRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 资源包ID
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// url地址分组
	Group *uint64 `json:"Group,omitnil,omitempty" name:"Group"`
}

func (r *CreateLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "ResourceIds")
	delete(f, "Group")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLicenseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLicenseResponse struct {
	*tchttp.BaseResponse
	Response *CreateLicenseResponseParams `json:"Response"`
}

func (r *CreateLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTestXMagicRequestParams struct {
	// 将腾讯特效创建到哪个应用下
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 营业执照
	CompanyPermit *string `json:"CompanyPermit,omitnil,omitempty" name:"CompanyPermit"`

	// 公司类型
	CompanyType *string `json:"CompanyType,omitnil,omitempty" name:"CompanyType"`

	// 公司名称
	CompanyName *string `json:"CompanyName,omitnil,omitempty" name:"CompanyName"`

	// 测试套餐名称
	PlanList []*string `json:"PlanList,omitnil,omitempty" name:"PlanList"`
}

type CreateTestXMagicRequest struct {
	*tchttp.BaseRequest
	
	// 将腾讯特效创建到哪个应用下
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 营业执照
	CompanyPermit *string `json:"CompanyPermit,omitnil,omitempty" name:"CompanyPermit"`

	// 公司类型
	CompanyType *string `json:"CompanyType,omitnil,omitempty" name:"CompanyType"`

	// 公司名称
	CompanyName *string `json:"CompanyName,omitnil,omitempty" name:"CompanyName"`

	// 测试套餐名称
	PlanList []*string `json:"PlanList,omitnil,omitempty" name:"PlanList"`
}

func (r *CreateTestXMagicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTestXMagicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "CompanyPermit")
	delete(f, "CompanyType")
	delete(f, "CompanyName")
	delete(f, "PlanList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTestXMagicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTestXMagicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTestXMagicResponse struct {
	*tchttp.BaseResponse
	Response *CreateTestXMagicResponseParams `json:"Response"`
}

func (r *CreateTestXMagicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTestXMagicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrialApplicationAndLicenseRequestParams struct {
	// 应用名
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 应用ID
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 功能 id 数组
	FeatureIds []*int64 `json:"FeatureIds,omitnil,omitempty" name:"FeatureIds"`

	// 是否要开通优图功能
	XMagic *bool `json:"XMagic,omitnil,omitempty" name:"XMagic"`

	// 营业执照
	CompanyPermit *string `json:"CompanyPermit,omitnil,omitempty" name:"CompanyPermit"`

	// 公司类型
	CompanyType *string `json:"CompanyType,omitnil,omitempty" name:"CompanyType"`

	// 公司名称
	CompanyName *string `json:"CompanyName,omitnil,omitempty" name:"CompanyName"`

	// 要开通的测试功能名称，基础套餐只能有一个
	PlanList []*string `json:"PlanList,omitnil,omitempty" name:"PlanList"`

	// Mac 进程名
	MacBundleId *string `json:"MacBundleId,omitnil,omitempty" name:"MacBundleId"`

	// Windows 进程名
	WinProcessName *string `json:"WinProcessName,omitnil,omitempty" name:"WinProcessName"`

	// 要创建到哪个平台，web、mobile、pc，默认mobile
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 授权域名列表
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`
}

type CreateTrialApplicationAndLicenseRequest struct {
	*tchttp.BaseRequest
	
	// 应用名
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 应用ID
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 功能 id 数组
	FeatureIds []*int64 `json:"FeatureIds,omitnil,omitempty" name:"FeatureIds"`

	// 是否要开通优图功能
	XMagic *bool `json:"XMagic,omitnil,omitempty" name:"XMagic"`

	// 营业执照
	CompanyPermit *string `json:"CompanyPermit,omitnil,omitempty" name:"CompanyPermit"`

	// 公司类型
	CompanyType *string `json:"CompanyType,omitnil,omitempty" name:"CompanyType"`

	// 公司名称
	CompanyName *string `json:"CompanyName,omitnil,omitempty" name:"CompanyName"`

	// 要开通的测试功能名称，基础套餐只能有一个
	PlanList []*string `json:"PlanList,omitnil,omitempty" name:"PlanList"`

	// Mac 进程名
	MacBundleId *string `json:"MacBundleId,omitnil,omitempty" name:"MacBundleId"`

	// Windows 进程名
	WinProcessName *string `json:"WinProcessName,omitnil,omitempty" name:"WinProcessName"`

	// 要创建到哪个平台，web、mobile、pc，默认mobile
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 授权域名列表
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`
}

func (r *CreateTrialApplicationAndLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrialApplicationAndLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "BundleId")
	delete(f, "PackageName")
	delete(f, "FeatureIds")
	delete(f, "XMagic")
	delete(f, "CompanyPermit")
	delete(f, "CompanyType")
	delete(f, "CompanyName")
	delete(f, "PlanList")
	delete(f, "MacBundleId")
	delete(f, "WinProcessName")
	delete(f, "Platform")
	delete(f, "DomainList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTrialApplicationAndLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrialApplicationAndLicenseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTrialApplicationAndLicenseResponse struct {
	*tchttp.BaseResponse
	Response *CreateTrialApplicationAndLicenseResponseParams `json:"Response"`
}

func (r *CreateTrialApplicationAndLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrialApplicationAndLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrialLicenseRequestParams struct {
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 功能 ID 数组
	FeatureIds []*uint64 `json:"FeatureIds,omitnil,omitempty" name:"FeatureIds"`

	// Url分组
	Group *uint64 `json:"Group,omitnil,omitempty" name:"Group"`
}

type CreateTrialLicenseRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApplicationId *string `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 功能 ID 数组
	FeatureIds []*uint64 `json:"FeatureIds,omitnil,omitempty" name:"FeatureIds"`

	// Url分组
	Group *uint64 `json:"Group,omitnil,omitempty" name:"Group"`
}

func (r *CreateTrialLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrialLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "FeatureIds")
	delete(f, "Group")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTrialLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTrialLicenseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTrialLicenseResponse struct {
	*tchttp.BaseResponse
	Response *CreateTrialLicenseResponseParams `json:"Response"`
}

func (r *CreateTrialLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTrialLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateXMagicRequestParams struct {
	// 应用ApplicationId
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 营业执照
	CompanyPermit *string `json:"CompanyPermit,omitnil,omitempty" name:"CompanyPermit"`

	// 公司类型
	CompanyType *string `json:"CompanyType,omitnil,omitempty" name:"CompanyType"`

	// 公司名称
	CompanyName *string `json:"CompanyName,omitnil,omitempty" name:"CompanyName"`

	// 优图资源id列表
	XMagicResourceIds []*string `json:"XMagicResourceIds,omitnil,omitempty" name:"XMagicResourceIds"`
}

type CreateXMagicRequest struct {
	*tchttp.BaseRequest
	
	// 应用ApplicationId
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 营业执照
	CompanyPermit *string `json:"CompanyPermit,omitnil,omitempty" name:"CompanyPermit"`

	// 公司类型
	CompanyType *string `json:"CompanyType,omitnil,omitempty" name:"CompanyType"`

	// 公司名称
	CompanyName *string `json:"CompanyName,omitnil,omitempty" name:"CompanyName"`

	// 优图资源id列表
	XMagicResourceIds []*string `json:"XMagicResourceIds,omitnil,omitempty" name:"XMagicResourceIds"`
}

func (r *CreateXMagicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateXMagicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "CompanyPermit")
	delete(f, "CompanyType")
	delete(f, "CompanyName")
	delete(f, "XMagicResourceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateXMagicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateXMagicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateXMagicResponse struct {
	*tchttp.BaseResponse
	Response *CreateXMagicResponseParams `json:"Response"`
}

func (r *CreateXMagicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateXMagicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationAndVideoLicenseRequestParams struct {
	// license唯一标识
	LicenseId *uint64 `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`
}

type DeleteApplicationAndVideoLicenseRequest struct {
	*tchttp.BaseRequest
	
	// license唯一标识
	LicenseId *uint64 `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`
}

func (r *DeleteApplicationAndVideoLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationAndVideoLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LicenseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationAndVideoLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationAndVideoLicenseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteApplicationAndVideoLicenseResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApplicationAndVideoLicenseResponseParams `json:"Response"`
}

func (r *DeleteApplicationAndVideoLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationAndVideoLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationAndWebPlayerLicenseRequestParams struct {
	// license唯一标识
	LicenseId *uint64 `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`
}

type DeleteApplicationAndWebPlayerLicenseRequest struct {
	*tchttp.BaseRequest
	
	// license唯一标识
	LicenseId *uint64 `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`
}

func (r *DeleteApplicationAndWebPlayerLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationAndWebPlayerLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LicenseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApplicationAndWebPlayerLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApplicationAndWebPlayerLicenseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteApplicationAndWebPlayerLicenseResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApplicationAndWebPlayerLicenseResponseParams `json:"Response"`
}

func (r *DeleteApplicationAndWebPlayerLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApplicationAndWebPlayerLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFeatureListRequestParams struct {

}

type DescribeFeatureListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeFeatureListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFeatureListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFeatureListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFeatureListResponseParams struct {
	// 功能列表
	FeatureList []*VideoFeature `json:"FeatureList,omitnil,omitempty" name:"FeatureList"`

	// 优图功能列表
	XMagicFeatureList []*XMagicFeature `json:"XMagicFeatureList,omitnil,omitempty" name:"XMagicFeatureList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFeatureListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFeatureListResponseParams `json:"Response"`
}

func (r *DescribeFeatureListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFeatureListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLicenseListRequestParams struct {

}

type DescribeLicenseListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLicenseListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLicenseListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLicenseListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLicenseListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLicenseListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLicenseListResponseParams `json:"Response"`
}

func (r *DescribeLicenseListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLicenseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNewsRequestParams struct {
	// 页码，从0开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数据量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeNewsRequest struct {
	*tchttp.BaseRequest
	
	// 页码，从0开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数据量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeNewsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNewsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNewsResponseParams struct {
	// 产品动态详情列表
	NewsList []*NewsInfo `json:"NewsList,omitnil,omitempty" name:"NewsList"`

	// 总数据量
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNewsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNewsResponseParams `json:"Response"`
}

func (r *DescribeNewsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNewsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSTSRequestParams struct {

}

type DescribeSTSRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSTSRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSTSRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSTSRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSTSResponseParams struct {
	// 桶名称
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 地区
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 开始生效时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 临时token过期时间
	ExpiredTime *uint64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 临时token
	SessionToken *string `json:"SessionToken,omitnil,omitempty" name:"SessionToken"`

	// 临时SecretId
	TmpSecretId *string `json:"TmpSecretId,omitnil,omitempty" name:"TmpSecretId"`

	// 临时秘钥
	TmpSecretKey *string `json:"TmpSecretKey,omitnil,omitempty" name:"TmpSecretKey"`

	// 上传的根路径，底下可以多层
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSTSResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSTSResponseParams `json:"Response"`
}

func (r *DescribeSTSResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSTSResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrialFeatureRequestParams struct {

}

type DescribeTrialFeatureRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeTrialFeatureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrialFeatureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrialFeatureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrialFeatureResponseParams struct {
	// 功能列表
	FeatureList []*Feature `json:"FeatureList,omitnil,omitempty" name:"FeatureList"`

	// 可以开通的优图测试功能
	// 注意：此字段可能返回 null，表示取不到有效值。
	XMagicTrial *XMagicTrial `json:"XMagicTrial,omitnil,omitempty" name:"XMagicTrial"`

	// 可以开通的优图测试功能列表
	XMagicTrialList []*XMagicTrial `json:"XMagicTrialList,omitnil,omitempty" name:"XMagicTrialList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrialFeatureResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrialFeatureResponseParams `json:"Response"`
}

func (r *DescribeTrialFeatureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrialFeatureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserConfigRequestParams struct {

}

type DescribeUserConfigRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserConfigResponseParams `json:"Response"`
}

func (r *DescribeUserConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVcubeApplicationAndLicenseRequestParams struct {
	// 查询对应业务方的license，直播:live 点播：vod
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 应用类型测试：test，正式：formal
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用名称，模糊查询
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Ios包名，模糊查询
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Android 包名，模糊查询
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 平台信息，pc 或者 mobile
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// Mac 进程名
	MacBundleId *string `json:"MacBundleId,omitnil,omitempty" name:"MacBundleId"`

	// Windows 进程名
	WinProcessName *string `json:"WinProcessName,omitnil,omitempty" name:"WinProcessName"`
}

type DescribeVcubeApplicationAndLicenseRequest struct {
	*tchttp.BaseRequest
	
	// 查询对应业务方的license，直播:live 点播：vod
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 应用类型测试：test，正式：formal
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用名称，模糊查询
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Ios包名，模糊查询
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Android 包名，模糊查询
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 平台信息，pc 或者 mobile
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// Mac 进程名
	MacBundleId *string `json:"MacBundleId,omitnil,omitempty" name:"MacBundleId"`

	// Windows 进程名
	WinProcessName *string `json:"WinProcessName,omitnil,omitempty" name:"WinProcessName"`
}

func (r *DescribeVcubeApplicationAndLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVcubeApplicationAndLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizType")
	delete(f, "AppType")
	delete(f, "AppName")
	delete(f, "BundleId")
	delete(f, "PackageName")
	delete(f, "Platform")
	delete(f, "MacBundleId")
	delete(f, "WinProcessName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVcubeApplicationAndLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVcubeApplicationAndLicenseResponseParams struct {
	// 应用license列表
	ApplicationInfoList []*ApplicationInfo `json:"ApplicationInfoList,omitnil,omitempty" name:"ApplicationInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVcubeApplicationAndLicenseResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVcubeApplicationAndLicenseResponseParams `json:"Response"`
}

func (r *DescribeVcubeApplicationAndLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVcubeApplicationAndLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVcubeApplicationAndPlayListRequestParams struct {
	// 查询对应业务方的license，直播:live 点播：vod
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 应用类型测试：test，正式：formal
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用名称，模糊查询
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Ios包名，模糊查询
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Android 包名，模糊查询
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 平台信息，pc 或者 mobile
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// Mac 进程名
	MacBundleId *string `json:"MacBundleId,omitnil,omitempty" name:"MacBundleId"`

	// Windows 进程名
	WinProcessName *string `json:"WinProcessName,omitnil,omitempty" name:"WinProcessName"`

	// license唯一标识
	LicenseId *uint64 `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`
}

type DescribeVcubeApplicationAndPlayListRequest struct {
	*tchttp.BaseRequest
	
	// 查询对应业务方的license，直播:live 点播：vod
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 应用类型测试：test，正式：formal
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用名称，模糊查询
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Ios包名，模糊查询
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Android 包名，模糊查询
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 平台信息，pc 或者 mobile
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// Mac 进程名
	MacBundleId *string `json:"MacBundleId,omitnil,omitempty" name:"MacBundleId"`

	// Windows 进程名
	WinProcessName *string `json:"WinProcessName,omitnil,omitempty" name:"WinProcessName"`

	// license唯一标识
	LicenseId *uint64 `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`
}

func (r *DescribeVcubeApplicationAndPlayListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVcubeApplicationAndPlayListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizType")
	delete(f, "AppType")
	delete(f, "AppName")
	delete(f, "BundleId")
	delete(f, "PackageName")
	delete(f, "Platform")
	delete(f, "MacBundleId")
	delete(f, "WinProcessName")
	delete(f, "LicenseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVcubeApplicationAndPlayListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVcubeApplicationAndPlayListResponseParams struct {
	// 应用license列表
	ApplicationInfoList []*ApplicationInfo `json:"ApplicationInfoList,omitnil,omitempty" name:"ApplicationInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVcubeApplicationAndPlayListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVcubeApplicationAndPlayListResponseParams `json:"Response"`
}

func (r *DescribeVcubeApplicationAndPlayListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVcubeApplicationAndPlayListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVcubeApplicationAndXMagicListRequestParams struct {
	// 查询对应业务方的license，直播:live 点播：vod
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 应用类型测试：test，正式：formal
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用名称，模糊查询
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Ios包名，模糊查询
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Android 包名，模糊查询
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 平台信息，pc 或者 mobile
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// Mac 进程名
	MacBundleId *string `json:"MacBundleId,omitnil,omitempty" name:"MacBundleId"`

	// Windows 进程名
	WinProcessName *string `json:"WinProcessName,omitnil,omitempty" name:"WinProcessName"`
}

type DescribeVcubeApplicationAndXMagicListRequest struct {
	*tchttp.BaseRequest
	
	// 查询对应业务方的license，直播:live 点播：vod
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 应用类型测试：test，正式：formal
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用名称，模糊查询
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Ios包名，模糊查询
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Android 包名，模糊查询
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 平台信息，pc 或者 mobile
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// Mac 进程名
	MacBundleId *string `json:"MacBundleId,omitnil,omitempty" name:"MacBundleId"`

	// Windows 进程名
	WinProcessName *string `json:"WinProcessName,omitnil,omitempty" name:"WinProcessName"`
}

func (r *DescribeVcubeApplicationAndXMagicListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVcubeApplicationAndXMagicListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizType")
	delete(f, "AppType")
	delete(f, "AppName")
	delete(f, "BundleId")
	delete(f, "PackageName")
	delete(f, "Platform")
	delete(f, "MacBundleId")
	delete(f, "WinProcessName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVcubeApplicationAndXMagicListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVcubeApplicationAndXMagicListResponseParams struct {
	// 应用license列表
	ApplicationInfoList []*ApplicationInfo `json:"ApplicationInfoList,omitnil,omitempty" name:"ApplicationInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVcubeApplicationAndXMagicListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVcubeApplicationAndXMagicListResponseParams `json:"Response"`
}

func (r *DescribeVcubeApplicationAndXMagicListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVcubeApplicationAndXMagicListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVcubeResourcesListRequestParams struct {
	// 查询资源页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页的数据量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用名称，模块查询
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 资源Id，资源包id或者license资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源Id、应用名称、包名
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`

	// 查询资源所属平台，web、mobile
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type DescribeVcubeResourcesListRequest struct {
	*tchttp.BaseRequest
	
	// 查询资源页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页的数据量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用名称，模块查询
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 资源Id，资源包id或者license资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源Id、应用名称、包名
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`

	// 查询资源所属平台，web、mobile
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`
}

func (r *DescribeVcubeResourcesListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVcubeResourcesListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "AppName")
	delete(f, "ResourceId")
	delete(f, "Word")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVcubeResourcesListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVcubeResourcesListResponseParams struct {
	// 资源列表
	ResourceList []*LicenseResourceInfo `json:"ResourceList,omitnil,omitempty" name:"ResourceList"`

	// 总数据量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVcubeResourcesListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVcubeResourcesListResponseParams `json:"Response"`
}

func (r *DescribeVcubeResourcesListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVcubeResourcesListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVcubeResourcesRequestParams struct {
	// 查找那个平台的资源
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type DescribeVcubeResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 查找那个平台的资源
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`
}

func (r *DescribeVcubeResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVcubeResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVcubeResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVcubeResourcesResponseParams struct {
	// 资源列表
	ResourceList []*LicenseResourceSimpleInfo `json:"ResourceList,omitnil,omitempty" name:"ResourceList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVcubeResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVcubeResourcesResponseParams `json:"Response"`
}

func (r *DescribeVcubeResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVcubeResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeXMagicResourceListRequestParams struct {
	// 分页页码 第一页是0
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 关联的应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 优图资源Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 查询关键字，资源Id、应用名称、包名
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`

	// 平台信息，pc 或者 mobile
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 所属业务，xmagic：优图， avatar：虚拟人，不传查全部
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`
}

type DescribeXMagicResourceListRequest struct {
	*tchttp.BaseRequest
	
	// 分页页码 第一页是0
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 关联的应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 优图资源Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 查询关键字，资源Id、应用名称、包名
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`

	// 平台信息，pc 或者 mobile
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 所属业务，xmagic：优图， avatar：虚拟人，不传查全部
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`
}

func (r *DescribeXMagicResourceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeXMagicResourceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "AppName")
	delete(f, "ResourceId")
	delete(f, "Word")
	delete(f, "Platform")
	delete(f, "BizType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeXMagicResourceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeXMagicResourceListResponseParams struct {
	// 优图资源包信息
	Resources []*XMagicResourceSimpleInfo `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 资源数量
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeXMagicResourceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeXMagicResourceListResponseParams `json:"Response"`
}

func (r *DescribeXMagicResourceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeXMagicResourceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeXMagicResourceRequestParams struct {
	// 分页页码 第一页是0
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeXMagicResourceRequest struct {
	*tchttp.BaseRequest
	
	// 分页页码 第一页是0
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeXMagicResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeXMagicResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeXMagicResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeXMagicResourceResponseParams struct {
	// 优图资源包信息
	Resources []*XMagicResource `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 资源数量
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeXMagicResourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeXMagicResourceResponseParams `json:"Response"`
}

func (r *DescribeXMagicResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeXMagicResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Feature struct {
	// 索引
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 功能ID
	FeatureId *uint64 `json:"FeatureId,omitnil,omitempty" name:"FeatureId"`

	// 功能名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 功能类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否可以申请试用
	Trial *bool `json:"Trial,omitnil,omitempty" name:"Trial"`

	// 可以试用的次数
	TrialCount *uint64 `json:"TrialCount,omitnil,omitempty" name:"TrialCount"`

	// 可以试用的时长，单位天
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`
}

type License struct {
	// license类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 老系统迁移备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// license生效时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// license失效时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// license对应的功能Id
	FeatureId *uint64 `json:"FeatureId,omitnil,omitempty" name:"FeatureId"`

	// license是测试：test还是正式：formal
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 测试license 是否可以续期
	Renewal *bool `json:"Renewal,omitnil,omitempty" name:"Renewal"`

	// license索引
	LicenseId *uint64 `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`

	// license名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 测试license 是否升级
	Update *bool `json:"Update,omitnil,omitempty" name:"Update"`

	// 兼容老的licenseUrl
	// 注意：此字段可能返回 null，表示取不到有效值。
	OldLicenseUrl *string `json:"OldLicenseUrl,omitnil,omitempty" name:"OldLicenseUrl"`

	// 视立方url分组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Group *uint64 `json:"Group,omitnil,omitempty" name:"Group"`

	// 过期：true ，未过期：其它
	Expired *bool `json:"Expired,omitnil,omitempty" name:"Expired"`

	// 返回还有多少秒过期，过期返回0
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestTime *uint64 `json:"RestTime,omitnil,omitempty" name:"RestTime"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 计费资源相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *RenewResource `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 是否为马甲包
	IsVest *bool `json:"IsVest,omitnil,omitempty" name:"IsVest"`
}

type LicenseResourceInfo struct {
	// 授权功能Id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 有效期时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 功能Id
	FeatureId *uint64 `json:"FeatureId,omitnil,omitempty" name:"FeatureId"`

	// 此license资源的开始生效时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 此license资源的生效结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 此资源是否可用于续期和更新有效期
	IsUse *bool `json:"IsUse,omitnil,omitempty" name:"IsUse"`

	// 此资源的状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 销毁时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolatedTimestamp *string `json:"IsolatedTimestamp,omitnil,omitempty" name:"IsolatedTimestamp"`

	// 功能模块名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 功能模块类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 资源包信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Package *Package `json:"Package,omitnil,omitempty" name:"Package"`

	// 应用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Application *AppInfo `json:"Application,omitnil,omitempty" name:"Application"`

	// 资源Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 自动续费状态
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

type LicenseResourceSimpleInfo struct {
	// 授权功能Id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 有效期时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 功能Id
	FeatureId *uint64 `json:"FeatureId,omitnil,omitempty" name:"FeatureId"`

	// 此license资源的开始生效时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 此license资源的生效结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 此资源是否可用于续期和更新有效期
	IsUse *bool `json:"IsUse,omitnil,omitempty" name:"IsUse"`

	// 此资源的状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 销毁时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolatedTimestamp *string `json:"IsolatedTimestamp,omitnil,omitempty" name:"IsolatedTimestamp"`

	// 功能模块名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 功能模块类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 资源包信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Package *Package `json:"Package,omitnil,omitempty" name:"Package"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 自动续费状态
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

// Predefined struct for user
type ModifyApplicationRequestParams struct {
	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 应用名
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 应用ID
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// Windows 进程名
	WinProcessName *string `json:"WinProcessName,omitnil,omitempty" name:"WinProcessName"`

	// Mac 进程名
	MacBundleId *string `json:"MacBundleId,omitnil,omitempty" name:"MacBundleId"`

	// 要追加的web域名列表
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`
}

type ModifyApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 应用名
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 应用ID
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// Windows 进程名
	WinProcessName *string `json:"WinProcessName,omitnil,omitempty" name:"WinProcessName"`

	// Mac 进程名
	MacBundleId *string `json:"MacBundleId,omitnil,omitempty" name:"MacBundleId"`

	// 要追加的web域名列表
	DomainList []*string `json:"DomainList,omitnil,omitempty" name:"DomainList"`
}

func (r *ModifyApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "AppName")
	delete(f, "BundleId")
	delete(f, "PackageName")
	delete(f, "WinProcessName")
	delete(f, "MacBundleId")
	delete(f, "DomainList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApplicationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApplicationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApplicationResponseParams `json:"Response"`
}

func (r *ModifyApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFormalApplicationRequestParams struct {
	// 应用Id
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Ios 包名称
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Android  包名称
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type ModifyFormalApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 应用Id
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// Ios 包名称
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// Android  包名称
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

func (r *ModifyFormalApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFormalApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "AppName")
	delete(f, "BundleId")
	delete(f, "PackageName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFormalApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFormalApplicationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyFormalApplicationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFormalApplicationResponseParams `json:"Response"`
}

func (r *ModifyFormalApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFormalApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLicenseRequestParams struct {
	// License 功能模块 Id
	LicenseId *string `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`

	// 资源包 Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type ModifyLicenseRequest struct {
	*tchttp.BaseRequest
	
	// License 功能模块 Id
	LicenseId *string `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`

	// 资源包 Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *ModifyLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LicenseId")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLicenseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLicenseResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLicenseResponseParams `json:"Response"`
}

func (r *ModifyLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPresetApplicationRequestParams struct {
	// 应用Id
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 您的app名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// ios包名
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 安卓包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type ModifyPresetApplicationRequest struct {
	*tchttp.BaseRequest
	
	// 应用Id
	ApplicationId *uint64 `json:"ApplicationId,omitnil,omitempty" name:"ApplicationId"`

	// 您的app名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// ios包名
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 安卓包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

func (r *ModifyPresetApplicationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPresetApplicationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplicationId")
	delete(f, "AppName")
	delete(f, "BundleId")
	delete(f, "PackageName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPresetApplicationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPresetApplicationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPresetApplicationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPresetApplicationResponseParams `json:"Response"`
}

func (r *ModifyPresetApplicationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPresetApplicationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTrialLicenseRequestParams struct {
	// 授权ID
	LicenseId *string `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`
}

type ModifyTrialLicenseRequest struct {
	*tchttp.BaseRequest
	
	// 授权ID
	LicenseId *string `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`
}

func (r *ModifyTrialLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTrialLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LicenseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTrialLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTrialLicenseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTrialLicenseResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTrialLicenseResponseParams `json:"Response"`
}

func (r *ModifyTrialLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTrialLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyXMagicRequestParams struct {
	// 资源Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 优图licenseId
	XMagicId *string `json:"XMagicId,omitnil,omitempty" name:"XMagicId"`
}

type ModifyXMagicRequest struct {
	*tchttp.BaseRequest
	
	// 资源Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 优图licenseId
	XMagicId *string `json:"XMagicId,omitnil,omitempty" name:"XMagicId"`
}

func (r *ModifyXMagicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyXMagicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "XMagicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyXMagicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyXMagicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyXMagicResponse struct {
	*tchttp.BaseResponse
	Response *ModifyXMagicResponseParams `json:"Response"`
}

func (r *ModifyXMagicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyXMagicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NewsInfo struct {
	// 新闻Id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type Package struct {
	// 资源包Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 资源包资源Id
	BizResourceId *string `json:"BizResourceId,omitnil,omitempty" name:"BizResourceId"`

	// 站点 中国站还是国际站
	Site *string `json:"Site,omitnil,omitempty" name:"Site"`

	// 资源包开始生效时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 资源包过期时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 资源包退费时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefundTime *string `json:"RefundTime,omitnil,omitempty" name:"RefundTime"`

	// 资源包名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资源包类型ID,对应点播：PackageId 对应直播：TypeId
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type RenewLicenseRequestParams struct {
	// License 功能模块 Id
	LicenseId *string `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`

	// 资源包 Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type RenewLicenseRequest struct {
	*tchttp.BaseRequest
	
	// License 功能模块 Id
	LicenseId *string `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`

	// 资源包 Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *RenewLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LicenseId")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewLicenseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewLicenseResponse struct {
	*tchttp.BaseResponse
	Response *RenewLicenseResponseParams `json:"Response"`
}

func (r *RenewLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewResource struct {
	// 资源Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 自动续期标记，0：默认，1：自动续期，2：不续期
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 资源冻结时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolatedTimestamp *string `json:"IsolatedTimestamp,omitnil,omitempty" name:"IsolatedTimestamp"`

	// 资源销毁状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Refund *bool `json:"Refund,omitnil,omitempty" name:"Refund"`

	// 子产品码
	SubProductCode *string `json:"SubProductCode,omitnil,omitempty" name:"SubProductCode"`

	// 资源到期时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

// Predefined struct for user
type RenewTestXMagicRequestParams struct {
	// 优图美视Id
	XMagicId *uint64 `json:"XMagicId,omitnil,omitempty" name:"XMagicId"`
}

type RenewTestXMagicRequest struct {
	*tchttp.BaseRequest
	
	// 优图美视Id
	XMagicId *uint64 `json:"XMagicId,omitnil,omitempty" name:"XMagicId"`
}

func (r *RenewTestXMagicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewTestXMagicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "XMagicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewTestXMagicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewTestXMagicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewTestXMagicResponse struct {
	*tchttp.BaseResponse
	Response *RenewTestXMagicResponseParams `json:"Response"`
}

func (r *RenewTestXMagicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewTestXMagicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewVideoRequestParams struct {
	// 要续期的视频播放license Id
	LicenseId *uint64 `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`
}

type RenewVideoRequest struct {
	*tchttp.BaseRequest
	
	// 要续期的视频播放license Id
	LicenseId *uint64 `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`
}

func (r *RenewVideoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewVideoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LicenseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewVideoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewVideoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewVideoResponse struct {
	*tchttp.BaseResponse
	Response *RenewVideoResponseParams `json:"Response"`
}

func (r *RenewVideoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewVideoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTestXMagicRequestParams struct {
	// 优图美视功能Id
	XMagicId *uint64 `json:"XMagicId,omitnil,omitempty" name:"XMagicId"`

	// XMagic套餐包Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type UpdateTestXMagicRequest struct {
	*tchttp.BaseRequest
	
	// 优图美视功能Id
	XMagicId *uint64 `json:"XMagicId,omitnil,omitempty" name:"XMagicId"`

	// XMagic套餐包Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *UpdateTestXMagicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTestXMagicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "XMagicId")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateTestXMagicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTestXMagicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateTestXMagicResponse struct {
	*tchttp.BaseResponse
	Response *UpdateTestXMagicResponseParams `json:"Response"`
}

func (r *UpdateTestXMagicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTestXMagicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTrialLicenseRequestParams struct {
	// 执照ID
	LicenseId *string `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`

	// 资源包ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type UpdateTrialLicenseRequest struct {
	*tchttp.BaseRequest
	
	// 执照ID
	LicenseId *string `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`

	// 资源包ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *UpdateTrialLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTrialLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LicenseId")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateTrialLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTrialLicenseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateTrialLicenseResponse struct {
	*tchttp.BaseResponse
	Response *UpdateTrialLicenseResponseParams `json:"Response"`
}

func (r *UpdateTrialLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTrialLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateXMagicRequestParams struct {
	// 要修改的XMagic Id
	XMagicId *uint64 `json:"XMagicId,omitnil,omitempty" name:"XMagicId"`

	// 资源id
	XMagicResourceId *string `json:"XMagicResourceId,omitnil,omitempty" name:"XMagicResourceId"`

	// 营业执照地址
	CompanyPermit *string `json:"CompanyPermit,omitnil,omitempty" name:"CompanyPermit"`

	// 公司类型
	CompanyType *string `json:"CompanyType,omitnil,omitempty" name:"CompanyType"`

	// 公司名称
	CompanyName *string `json:"CompanyName,omitnil,omitempty" name:"CompanyName"`
}

type UpdateXMagicRequest struct {
	*tchttp.BaseRequest
	
	// 要修改的XMagic Id
	XMagicId *uint64 `json:"XMagicId,omitnil,omitempty" name:"XMagicId"`

	// 资源id
	XMagicResourceId *string `json:"XMagicResourceId,omitnil,omitempty" name:"XMagicResourceId"`

	// 营业执照地址
	CompanyPermit *string `json:"CompanyPermit,omitnil,omitempty" name:"CompanyPermit"`

	// 公司类型
	CompanyType *string `json:"CompanyType,omitnil,omitempty" name:"CompanyType"`

	// 公司名称
	CompanyName *string `json:"CompanyName,omitnil,omitempty" name:"CompanyName"`
}

func (r *UpdateXMagicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateXMagicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "XMagicId")
	delete(f, "XMagicResourceId")
	delete(f, "CompanyPermit")
	delete(f, "CompanyType")
	delete(f, "CompanyName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateXMagicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateXMagicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateXMagicResponse struct {
	*tchttp.BaseResponse
	Response *UpdateXMagicResponseParams `json:"Response"`
}

func (r *UpdateXMagicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateXMagicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VideoFeature struct {
	// 索引
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 功能ID
	FeatureId *uint64 `json:"FeatureId,omitnil,omitempty" name:"FeatureId"`

	// 功能名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 功能类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 是否可以申请试用
	Trial *bool `json:"Trial,omitnil,omitempty" name:"Trial"`

	// 可以试用的次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrialCount *uint64 `json:"TrialCount,omitnil,omitempty" name:"TrialCount"`

	// 可以试用的时长，单位天
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 功能所属业务方
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 平台类型
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type XMagicFeature struct {
	// 功能名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 可以申请的次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrialCount *uint64 `json:"TrialCount,omitnil,omitempty" name:"TrialCount"`

	// 每次申请的时长单位：天
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 功能类别
	Plan *string `json:"Plan,omitnil,omitempty" name:"Plan"`

	// single: 原子能力,combined:套餐
	XMagicType *string `json:"XMagicType,omitnil,omitempty" name:"XMagicType"`

	// 此功能是否支持开通测试
	Trial *bool `json:"Trial,omitnil,omitempty" name:"Trial"`

	// 功能所属业务方
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`
}

type XMagicInfo struct {
	// 优图Id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户公司名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompanyName *string `json:"CompanyName,omitnil,omitempty" name:"CompanyName"`

	// https://cos.xxx
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompanyPermit *string `json:"CompanyPermit,omitnil,omitempty" name:"CompanyPermit"`

	// 用户公司行业类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompanyType *string `json:"CompanyType,omitnil,omitempty" name:"CompanyType"`

	// 优图套餐类型
	Plan *string `json:"Plan,omitnil,omitempty" name:"Plan"`

	// 测试版还是正式版 test | formal
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// 0: 预申请，需要补充申请信息
	//    1: 审批通过，xmagic已签发，正在生效中
	//     2: 提交完申请资料后待运营审核状态
	//     3: 申请被驳回，需要重新修改申请资料
	//     4: 应用包名被修改后触发xmagic审批，当前xmagic已暂停生效
	//     5: 应用修改包名后，审批未通过状态，可以重新修改应用PB，状态会回到4
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 测试license是否已经升级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Update *bool `json:"Update,omitnil,omitempty" name:"Update"`

	// 优图生效开始时间 Status为1的时候
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 优图生效结束时间 Status为1的时候
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 续期次数， LicenseType=test时有此字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewalCount *uint64 `json:"RenewalCount,omitnil,omitempty" name:"RenewalCount"`

	// 历次审批的回复
	Reply []*string `json:"Reply,omitnil,omitempty" name:"Reply"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 修改时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 用户更新信息的时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 是否过期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Expired *bool `json:"Expired,omitnil,omitempty" name:"Expired"`

	// 返回还剩多少秒过期，过期后返回0
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestTime *uint64 `json:"RestTime,omitnil,omitempty" name:"RestTime"`

	// single: 原子能力,combined:套餐
	XMagicType *string `json:"XMagicType,omitnil,omitempty" name:"XMagicType"`

	// 优图模块名称，自动中英文
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 优图资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resource *RenewResource `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 是否是马甲包
	IsVest *bool `json:"IsVest,omitnil,omitempty" name:"IsVest"`
}

type XMagicResource struct {
	// 资源Id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 套餐类别
	Plan *string `json:"Plan,omitnil,omitempty" name:"Plan"`

	// 单位：年
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 资源是否已使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	XMagic *bool `json:"XMagic,omitnil,omitempty" name:"XMagic"`

	// 资源开始生效时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 资源结束生效时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 资源包是否过期
	Expired *bool `json:"Expired,omitnil,omitempty" name:"Expired"`

	// 功能模块名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// single: 原子能力,combined:套餐
	XMagicType *string `json:"XMagicType,omitnil,omitempty" name:"XMagicType"`

	// xmagic:优图， avatar：虚拟人
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 资源Id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源是否可以使用
	IsUse *bool `json:"IsUse,omitnil,omitempty" name:"IsUse"`
}

type XMagicResourceSimpleInfo struct {
	// 资源Id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 套餐类别
	Plan *string `json:"Plan,omitnil,omitempty" name:"Plan"`

	// 单位：年
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新时间
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 资源开始生效时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 资源结束生效时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Application *AppInfo `json:"Application,omitnil,omitempty" name:"Application"`

	// 开通的优图功能信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	XMagic *XMagicSimpleInfo `json:"XMagic,omitnil,omitempty" name:"XMagic"`

	// 优图资源状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 操作日期记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation []*string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 是否可以使用
	IsUse *bool `json:"IsUse,omitnil,omitempty" name:"IsUse"`

	// single: 原子能力,combined:套餐
	XMagicType *string `json:"XMagicType,omitnil,omitempty" name:"XMagicType"`

	// 功能模块名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资源所属业务方 xmagic：优图，avatar：虚拟形象 
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源自动续费状态
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

type XMagicSimpleInfo struct {
	// XMagic 的Id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// XMagic 状态
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type XMagicTrial struct {
	// 功能名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 可以申请的次数
	TrialCount *uint64 `json:"TrialCount,omitnil,omitempty" name:"TrialCount"`

	// 每次申请的时长单位：天
	Duration *uint64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 功能类别
	Plan *string `json:"Plan,omitnil,omitempty" name:"Plan"`

	// single: 原子能力,combined:套餐
	XMagicType *string `json:"XMagicType,omitnil,omitempty" name:"XMagicType"`

	// vod：点播 live：直播
	BizType *string `json:"BizType,omitnil,omitempty" name:"BizType"`
}