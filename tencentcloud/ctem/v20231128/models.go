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

package v20231128

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type CreateAppRequestParams struct {
	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 移动端名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 图片地址
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// 平台，ios或android
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 版本
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// 下载地址
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 安装包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 开发者
	Developer *string `json:"Developer,omitnil,omitempty" name:"Developer"`

	// 移动端地址
	ServerUrl *string `json:"ServerUrl,omitnil,omitempty" name:"ServerUrl"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 子公司ID
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`
}

type CreateAppRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 移动端名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 图片地址
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// 平台，ios或android
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 版本
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`

	// 下载地址
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 安装包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 开发者
	Developer *string `json:"Developer,omitnil,omitempty" name:"Developer"`

	// 移动端地址
	ServerUrl *string `json:"ServerUrl,omitnil,omitempty" name:"ServerUrl"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 子公司ID
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`
}

func (r *CreateAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerId")
	delete(f, "Name")
	delete(f, "Logo")
	delete(f, "Platform")
	delete(f, "AppVersion")
	delete(f, "DownloadUrl")
	delete(f, "PackageName")
	delete(f, "Developer")
	delete(f, "ServerUrl")
	delete(f, "Description")
	delete(f, "EnterpriseUid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppResponseParams struct {
	// Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAppResponse struct {
	*tchttp.BaseResponse
	Response *CreateAppResponseParams `json:"Response"`
}

func (r *CreateAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetRequestParams struct {
	// IP地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 企业Id
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 类型
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 国家
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// 省份
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// 城市
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 子公司Id
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`
}

type CreateAssetRequest struct {
	*tchttp.BaseRequest
	
	// IP地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 企业Id
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 类型
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 国家
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// 省份
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// 城市
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 子公司Id
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`
}

func (r *CreateAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ip")
	delete(f, "CustomerId")
	delete(f, "Os")
	delete(f, "Country")
	delete(f, "Province")
	delete(f, "City")
	delete(f, "Isp")
	delete(f, "EnterpriseUid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAssetResponseParams struct {
	// Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAssetResponse struct {
	*tchttp.BaseResponse
	Response *CreateAssetResponseParams `json:"Response"`
}

func (r *CreateAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomerRequestParams struct {
	// 企业名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资产收集、漏洞信息、弱口令、目录爆破、暗网泄露、Github泄露、文库网盘泄露、敏感信息泄露，其中资产收集必包含，多个用英文逗号隔离，例如：资产收集,漏洞信息
	ScanType *string `json:"ScanType,omitnil,omitempty" name:"ScanType"`

	// 百分比取值范围为30-100
	Percent *int64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 周期测绘时间
	ScanCron *string `json:"ScanCron,omitnil,omitempty" name:"ScanCron"`

	// 是否立即启动
	IsScanNow *bool `json:"IsScanNow,omitnil,omitempty" name:"IsScanNow"`

	// 是否启用周期测绘
	EnableCron *bool `json:"EnableCron,omitnil,omitempty" name:"EnableCron"`

	// 是否扫描子公司
	EnableScanSubEnterprise *bool `json:"EnableScanSubEnterprise,omitnil,omitempty" name:"EnableScanSubEnterprise"`

	// 是否授权
	EnableAuth *bool `json:"EnableAuth,omitnil,omitempty" name:"EnableAuth"`

	// 授权开始时间
	AuthStartAt *string `json:"AuthStartAt,omitnil,omitempty" name:"AuthStartAt"`

	// 授权结束时间
	AuthEndAt *string `json:"AuthEndAt,omitnil,omitempty" name:"AuthEndAt"`

	// 授权文件id
	AuthFile *string `json:"AuthFile,omitnil,omitempty" name:"AuthFile"`

	// 测绘时间配置项，采用json字符串格式
	ScanTime *string `json:"ScanTime,omitnil,omitempty" name:"ScanTime"`

	// 企业相关的关键字
	Keywords *string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// 图标
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`

	// 并发设置
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// 限制子公司层级，-1表示不限制
	SubCompanyLevel *int64 `json:"SubCompanyLevel,omitnil,omitempty" name:"SubCompanyLevel"`

	// 是否以企业名称为起点做完整扫描(包含手动上传),如只想扫描特定的某几个域名，则传false。
	IsIncludeFullScan *bool `json:"IsIncludeFullScan,omitnil,omitempty" name:"IsIncludeFullScan"`
}

type CreateCustomerRequest struct {
	*tchttp.BaseRequest
	
	// 企业名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资产收集、漏洞信息、弱口令、目录爆破、暗网泄露、Github泄露、文库网盘泄露、敏感信息泄露，其中资产收集必包含，多个用英文逗号隔离，例如：资产收集,漏洞信息
	ScanType *string `json:"ScanType,omitnil,omitempty" name:"ScanType"`

	// 百分比取值范围为30-100
	Percent *int64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 周期测绘时间
	ScanCron *string `json:"ScanCron,omitnil,omitempty" name:"ScanCron"`

	// 是否立即启动
	IsScanNow *bool `json:"IsScanNow,omitnil,omitempty" name:"IsScanNow"`

	// 是否启用周期测绘
	EnableCron *bool `json:"EnableCron,omitnil,omitempty" name:"EnableCron"`

	// 是否扫描子公司
	EnableScanSubEnterprise *bool `json:"EnableScanSubEnterprise,omitnil,omitempty" name:"EnableScanSubEnterprise"`

	// 是否授权
	EnableAuth *bool `json:"EnableAuth,omitnil,omitempty" name:"EnableAuth"`

	// 授权开始时间
	AuthStartAt *string `json:"AuthStartAt,omitnil,omitempty" name:"AuthStartAt"`

	// 授权结束时间
	AuthEndAt *string `json:"AuthEndAt,omitnil,omitempty" name:"AuthEndAt"`

	// 授权文件id
	AuthFile *string `json:"AuthFile,omitnil,omitempty" name:"AuthFile"`

	// 测绘时间配置项，采用json字符串格式
	ScanTime *string `json:"ScanTime,omitnil,omitempty" name:"ScanTime"`

	// 企业相关的关键字
	Keywords *string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// 图标
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`

	// 并发设置
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// 限制子公司层级，-1表示不限制
	SubCompanyLevel *int64 `json:"SubCompanyLevel,omitnil,omitempty" name:"SubCompanyLevel"`

	// 是否以企业名称为起点做完整扫描(包含手动上传),如只想扫描特定的某几个域名，则传false。
	IsIncludeFullScan *bool `json:"IsIncludeFullScan,omitnil,omitempty" name:"IsIncludeFullScan"`
}

func (r *CreateCustomerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ScanType")
	delete(f, "Percent")
	delete(f, "ScanCron")
	delete(f, "IsScanNow")
	delete(f, "EnableCron")
	delete(f, "EnableScanSubEnterprise")
	delete(f, "EnableAuth")
	delete(f, "AuthStartAt")
	delete(f, "AuthEndAt")
	delete(f, "AuthFile")
	delete(f, "ScanTime")
	delete(f, "Keywords")
	delete(f, "Icon")
	delete(f, "Qps")
	delete(f, "SubCompanyLevel")
	delete(f, "IsIncludeFullScan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCustomerResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomerResponseParams `json:"Response"`
}

func (r *CreateCustomerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainRequestParams struct {
	// 企业Id
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 主域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// ICP
	ICP *string `json:"ICP,omitnil,omitempty" name:"ICP"`

	// 注册时间
	RegisteredTime *string `json:"RegisteredTime,omitnil,omitempty" name:"RegisteredTime"`

	// 过期时间
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 公司
	Company *string `json:"Company,omitnil,omitempty" name:"Company"`

	// 子公司
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`
}

type CreateDomainRequest struct {
	*tchttp.BaseRequest
	
	// 企业Id
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 主域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// ICP
	ICP *string `json:"ICP,omitnil,omitempty" name:"ICP"`

	// 注册时间
	RegisteredTime *string `json:"RegisteredTime,omitnil,omitempty" name:"RegisteredTime"`

	// 过期时间
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 公司
	Company *string `json:"Company,omitnil,omitempty" name:"Company"`

	// 子公司
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`
}

func (r *CreateDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerId")
	delete(f, "Domain")
	delete(f, "ICP")
	delete(f, "RegisteredTime")
	delete(f, "ExpiredTime")
	delete(f, "Company")
	delete(f, "EnterpriseUid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDomainResponseParams struct {
	// Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateDomainResponseParams `json:"Response"`
}

func (r *CreateDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnterpriseRequestParams struct {
	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 上一级企业
	ParentEnterpriseUid *string `json:"ParentEnterpriseUid,omitnil,omitempty" name:"ParentEnterpriseUid"`

	// 统一社会信用代码
	CreditCode *string `json:"CreditCode,omitnil,omitempty" name:"CreditCode"`

	// 企业状态:存续、已注销
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 注册资本（单位:元）
	RegisteredCapital *string `json:"RegisteredCapital,omitnil,omitempty" name:"RegisteredCapital"`

	// 持股比例
	ShareholdingRatio *string `json:"ShareholdingRatio,omitnil,omitempty" name:"ShareholdingRatio"`

	// 法人代表
	LegalPerson *string `json:"LegalPerson,omitnil,omitempty" name:"LegalPerson"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 行业类型
	Industry *string `json:"Industry,omitnil,omitempty" name:"Industry"`

	// 子公司ID
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`
}

type CreateEnterpriseRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 上一级企业
	ParentEnterpriseUid *string `json:"ParentEnterpriseUid,omitnil,omitempty" name:"ParentEnterpriseUid"`

	// 统一社会信用代码
	CreditCode *string `json:"CreditCode,omitnil,omitempty" name:"CreditCode"`

	// 企业状态:存续、已注销
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 注册资本（单位:元）
	RegisteredCapital *string `json:"RegisteredCapital,omitnil,omitempty" name:"RegisteredCapital"`

	// 持股比例
	ShareholdingRatio *string `json:"ShareholdingRatio,omitnil,omitempty" name:"ShareholdingRatio"`

	// 法人代表
	LegalPerson *string `json:"LegalPerson,omitnil,omitempty" name:"LegalPerson"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 行业类型
	Industry *string `json:"Industry,omitnil,omitempty" name:"Industry"`

	// 子公司ID
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`
}

func (r *CreateEnterpriseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnterpriseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerId")
	delete(f, "Name")
	delete(f, "ParentEnterpriseUid")
	delete(f, "CreditCode")
	delete(f, "Status")
	delete(f, "RegisteredCapital")
	delete(f, "ShareholdingRatio")
	delete(f, "LegalPerson")
	delete(f, "Type")
	delete(f, "Industry")
	delete(f, "EnterpriseUid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEnterpriseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnterpriseResponseParams struct {
	// Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEnterpriseResponse struct {
	*tchttp.BaseResponse
	Response *CreateEnterpriseResponseParams `json:"Response"`
}

func (r *CreateEnterpriseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnterpriseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHttpRequestParams struct {
	// 企业Id
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// Url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 子公司
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`

	// 标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 报文长度
	ContentLength *int64 `json:"ContentLength,omitnil,omitempty" name:"ContentLength"`

	// 报文内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 缩略图Url
	ScreenshotUrl *string `json:"ScreenshotUrl,omitnil,omitempty" name:"ScreenshotUrl"`

	// 标签
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 状态码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 解析的IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 证书信息
	Ssl *string `json:"Ssl,omitnil,omitempty" name:"Ssl"`

	// ssl证书过期时间
	SslExpiredTime *string `json:"SslExpiredTime,omitnil,omitempty" name:"SslExpiredTime"`
}

type CreateHttpRequest struct {
	*tchttp.BaseRequest
	
	// 企业Id
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// Url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 子公司
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`

	// 标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 报文长度
	ContentLength *int64 `json:"ContentLength,omitnil,omitempty" name:"ContentLength"`

	// 报文内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 缩略图Url
	ScreenshotUrl *string `json:"ScreenshotUrl,omitnil,omitempty" name:"ScreenshotUrl"`

	// 标签
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 状态码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 解析的IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 证书信息
	Ssl *string `json:"Ssl,omitnil,omitempty" name:"Ssl"`

	// ssl证书过期时间
	SslExpiredTime *string `json:"SslExpiredTime,omitnil,omitempty" name:"SslExpiredTime"`
}

func (r *CreateHttpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHttpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerId")
	delete(f, "Url")
	delete(f, "EnterpriseUid")
	delete(f, "Title")
	delete(f, "ContentLength")
	delete(f, "Content")
	delete(f, "ScreenshotUrl")
	delete(f, "Tags")
	delete(f, "Code")
	delete(f, "Ip")
	delete(f, "Ssl")
	delete(f, "SslExpiredTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHttpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHttpResponseParams struct {
	// Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateHttpResponse struct {
	*tchttp.BaseResponse
	Response *CreateHttpResponseParams `json:"Response"`
}

func (r *CreateHttpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHttpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateJobRecordRequestParams struct {
	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 任务类型：即时任务
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 资产收集、漏洞信息、弱口令、目录爆破、暗网泄露、Github泄露、文库网盘泄露、敏感信息泄露，其中资产收集必包含，多个用英文逗号隔离，例如：资产收集,漏洞信息
	ScanType *string `json:"ScanType,omitnil,omitempty" name:"ScanType"`

	// qps设置
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// 是否包含完整扫描
	IsIncludeFullScan *bool `json:"IsIncludeFullScan,omitnil,omitempty" name:"IsIncludeFullScan"`
}

type CreateJobRecordRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 任务类型：即时任务
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 资产收集、漏洞信息、弱口令、目录爆破、暗网泄露、Github泄露、文库网盘泄露、敏感信息泄露，其中资产收集必包含，多个用英文逗号隔离，例如：资产收集,漏洞信息
	ScanType *string `json:"ScanType,omitnil,omitempty" name:"ScanType"`

	// qps设置
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// 是否包含完整扫描
	IsIncludeFullScan *bool `json:"IsIncludeFullScan,omitnil,omitempty" name:"IsIncludeFullScan"`
}

func (r *CreateJobRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateJobRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerId")
	delete(f, "TaskType")
	delete(f, "ScanType")
	delete(f, "Qps")
	delete(f, "IsIncludeFullScan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateJobRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateJobRecordResponseParams struct {
	// 任务Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateJobRecordResponse struct {
	*tchttp.BaseResponse
	Response *CreateJobRecordResponseParams `json:"Response"`
}

func (r *CreateJobRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateJobRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateManageRequestParams struct {
	// 企业Id
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// Url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Screenshot
	Screenshot *string `json:"Screenshot,omitnil,omitempty" name:"Screenshot"`

	// 状态码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 子公司
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`
}

type CreateManageRequest struct {
	*tchttp.BaseRequest
	
	// 企业Id
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// Url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Screenshot
	Screenshot *string `json:"Screenshot,omitnil,omitempty" name:"Screenshot"`

	// 状态码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 子公司
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`
}

func (r *CreateManageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateManageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerId")
	delete(f, "Url")
	delete(f, "Title")
	delete(f, "Screenshot")
	delete(f, "Code")
	delete(f, "EnterpriseUid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateManageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateManageResponseParams struct {
	// Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateManageResponse struct {
	*tchttp.BaseResponse
	Response *CreateManageResponseParams `json:"Response"`
}

func (r *CreateManageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateManageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePortRequestParams struct {
	// 企业Id
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// IP或域名地址
	Asset *string `json:"Asset,omitnil,omitempty" name:"Asset"`

	// 是否高危
	IsHighRisk *bool `json:"IsHighRisk,omitnil,omitempty" name:"IsHighRisk"`

	// 子公司
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`

	// base64编码后的指纹
	Banner *string `json:"Banner,omitnil,omitempty" name:"Banner"`

	// 解析的IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 组件名称
	App *string `json:"App,omitnil,omitempty" name:"App"`

	// 服务名称
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`
}

type CreatePortRequest struct {
	*tchttp.BaseRequest
	
	// 企业Id
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// IP或域名地址
	Asset *string `json:"Asset,omitnil,omitempty" name:"Asset"`

	// 是否高危
	IsHighRisk *bool `json:"IsHighRisk,omitnil,omitempty" name:"IsHighRisk"`

	// 子公司
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`

	// base64编码后的指纹
	Banner *string `json:"Banner,omitnil,omitempty" name:"Banner"`

	// 解析的IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 组件名称
	App *string `json:"App,omitnil,omitempty" name:"App"`

	// 服务名称
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`
}

func (r *CreatePortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePortRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerId")
	delete(f, "Port")
	delete(f, "Asset")
	delete(f, "IsHighRisk")
	delete(f, "EnterpriseUid")
	delete(f, "Banner")
	delete(f, "Ip")
	delete(f, "App")
	delete(f, "Service")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePortRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePortResponseParams struct {
	// Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePortResponse struct {
	*tchttp.BaseResponse
	Response *CreatePortResponseParams `json:"Response"`
}

func (r *CreatePortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSeedsRequestParams struct {
	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// ip种子数组
	Ips []*string `json:"Ips,omitnil,omitempty" name:"Ips"`

	// icon种子数组
	Icons []*string `json:"Icons,omitnil,omitempty" name:"Icons"`

	// 主域名种子数组
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// title种子数组
	Titles []*string `json:"Titles,omitnil,omitempty" name:"Titles"`

	// 子域名种子数组
	SubDomains []*string `json:"SubDomains,omitnil,omitempty" name:"SubDomains"`

	// 关键词种子数组
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// 母公司种子数组
	ParentCompanies []*string `json:"ParentCompanies,omitnil,omitempty" name:"ParentCompanies"`
}

type CreateSeedsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// ip种子数组
	Ips []*string `json:"Ips,omitnil,omitempty" name:"Ips"`

	// icon种子数组
	Icons []*string `json:"Icons,omitnil,omitempty" name:"Icons"`

	// 主域名种子数组
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// title种子数组
	Titles []*string `json:"Titles,omitnil,omitempty" name:"Titles"`

	// 子域名种子数组
	SubDomains []*string `json:"SubDomains,omitnil,omitempty" name:"SubDomains"`

	// 关键词种子数组
	Keywords []*string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// 母公司种子数组
	ParentCompanies []*string `json:"ParentCompanies,omitnil,omitempty" name:"ParentCompanies"`
}

func (r *CreateSeedsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSeedsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerId")
	delete(f, "Ips")
	delete(f, "Icons")
	delete(f, "Domains")
	delete(f, "Titles")
	delete(f, "SubDomains")
	delete(f, "Keywords")
	delete(f, "ParentCompanies")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSeedsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSeedsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSeedsResponse struct {
	*tchttp.BaseResponse
	Response *CreateSeedsResponseParams `json:"Response"`
}

func (r *CreateSeedsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSeedsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubDomainRequestParams struct {
	// 企业Id
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 子域名
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// Ip
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 国家
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// 省
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// 城市
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// Isp
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 子公司
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`
}

type CreateSubDomainRequest struct {
	*tchttp.BaseRequest
	
	// 企业Id
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 子域名
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// Ip
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 国家
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// 省
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// 城市
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// Isp
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 子公司
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`
}

func (r *CreateSubDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerId")
	delete(f, "SubDomain")
	delete(f, "Ip")
	delete(f, "Country")
	delete(f, "Province")
	delete(f, "City")
	delete(f, "Isp")
	delete(f, "EnterpriseUid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSubDomainResponseParams struct {
	// Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSubDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateSubDomainResponseParams `json:"Response"`
}

func (r *CreateSubDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSuspiciousAssetRequestParams struct {
	// 企业Id
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// Url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 子公司
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`

	// 标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 标签
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 所属者
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 来源类型
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 来源值
	SourceValue *string `json:"SourceValue,omitnil,omitempty" name:"SourceValue"`

	// 是否信任
	Trusted *bool `json:"Trusted,omitnil,omitempty" name:"Trusted"`
}

type CreateSuspiciousAssetRequest struct {
	*tchttp.BaseRequest
	
	// 企业Id
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// Url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 子公司
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`

	// 标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 标签
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 所属者
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 来源类型
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 来源值
	SourceValue *string `json:"SourceValue,omitnil,omitempty" name:"SourceValue"`

	// 是否信任
	Trusted *bool `json:"Trusted,omitnil,omitempty" name:"Trusted"`
}

func (r *CreateSuspiciousAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSuspiciousAssetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerId")
	delete(f, "Url")
	delete(f, "EnterpriseUid")
	delete(f, "Title")
	delete(f, "Tags")
	delete(f, "Owner")
	delete(f, "SourceType")
	delete(f, "SourceValue")
	delete(f, "Trusted")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSuspiciousAssetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSuspiciousAssetResponseParams struct {
	// Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSuspiciousAssetResponse struct {
	*tchttp.BaseResponse
	Response *CreateSuspiciousAssetResponseParams `json:"Response"`
}

func (r *CreateSuspiciousAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSuspiciousAssetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWechatAppletRequestParams struct {
	// 企业Id
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 图片地址
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// 账号
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`

	// 二维码
	QrCode *string `json:"QrCode,omitnil,omitempty" name:"QrCode"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 子公司
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`

	// 账号Appid
	AccountAppid *string `json:"AccountAppid,omitnil,omitempty" name:"AccountAppid"`

	// 认证主体
	RecordSubject *string `json:"RecordSubject,omitnil,omitempty" name:"RecordSubject"`
}

type CreateWechatAppletRequest struct {
	*tchttp.BaseRequest
	
	// 企业Id
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 图片地址
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// 账号
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`

	// 二维码
	QrCode *string `json:"QrCode,omitnil,omitempty" name:"QrCode"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 子公司
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`

	// 账号Appid
	AccountAppid *string `json:"AccountAppid,omitnil,omitempty" name:"AccountAppid"`

	// 认证主体
	RecordSubject *string `json:"RecordSubject,omitnil,omitempty" name:"RecordSubject"`
}

func (r *CreateWechatAppletRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWechatAppletRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerId")
	delete(f, "Name")
	delete(f, "Logo")
	delete(f, "AccountId")
	delete(f, "QrCode")
	delete(f, "Description")
	delete(f, "EnterpriseUid")
	delete(f, "AccountAppid")
	delete(f, "RecordSubject")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWechatAppletRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWechatAppletResponseParams struct {
	// Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWechatAppletResponse struct {
	*tchttp.BaseResponse
	Response *CreateWechatAppletResponseParams `json:"Response"`
}

func (r *CreateWechatAppletResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWechatAppletResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWechatOfficialAccountRequestParams struct {
	// 企业Id
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 图片地址
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// 账号
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`

	// 二维码
	QrCode *string `json:"QrCode,omitnil,omitempty" name:"QrCode"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 子公司
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`

	// 认证主体
	RecordSubject *string `json:"RecordSubject,omitnil,omitempty" name:"RecordSubject"`
}

type CreateWechatOfficialAccountRequest struct {
	*tchttp.BaseRequest
	
	// 企业Id
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 图片地址
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// 账号
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`

	// 二维码
	QrCode *string `json:"QrCode,omitnil,omitempty" name:"QrCode"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 子公司
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`

	// 认证主体
	RecordSubject *string `json:"RecordSubject,omitnil,omitempty" name:"RecordSubject"`
}

func (r *CreateWechatOfficialAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWechatOfficialAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerId")
	delete(f, "Name")
	delete(f, "Logo")
	delete(f, "AccountId")
	delete(f, "QrCode")
	delete(f, "Description")
	delete(f, "EnterpriseUid")
	delete(f, "RecordSubject")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWechatOfficialAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWechatOfficialAccountResponseParams struct {
	// Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWechatOfficialAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateWechatOfficialAccountResponseParams `json:"Response"`
}

func (r *CreateWechatOfficialAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWechatOfficialAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Customer struct {
	// 企业ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 企业名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 股权占比
	Percent *int64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 资产收集、漏洞信息、弱口令、目录爆破、暗网泄露、Github泄露、文库网盘泄露、敏感信息泄露，其中资产收集必包含，多个用英文逗号隔离，例如：资产收集,漏洞信息
	ScanType *string `json:"ScanType,omitnil,omitempty" name:"ScanType"`

	// 创建账号
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 腾讯云客户AppId
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 腾讯云客户Uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 创建时间
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// 更新时间
	UpdateAt *string `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`

	// 周期测绘时间
	ScanCron *string `json:"ScanCron,omitnil,omitempty" name:"ScanCron"`

	// 是否启用周期测绘
	EnableCron *bool `json:"EnableCron,omitnil,omitempty" name:"EnableCron"`

	// 是否扫描子公司
	EnableScanSubEnterprise *bool `json:"EnableScanSubEnterprise,omitnil,omitempty" name:"EnableScanSubEnterprise"`

	// 是否授权
	EnableAuth *bool `json:"EnableAuth,omitnil,omitempty" name:"EnableAuth"`

	// 授权开始时间
	AuthStartAt *string `json:"AuthStartAt,omitnil,omitempty" name:"AuthStartAt"`

	// 授权结束时间
	AuthEndAt *string `json:"AuthEndAt,omitnil,omitempty" name:"AuthEndAt"`

	// 授权文件id
	AuthFile *string `json:"AuthFile,omitnil,omitempty" name:"AuthFile"`

	// 测绘时间配置项
	ScanTime *string `json:"ScanTime,omitnil,omitempty" name:"ScanTime"`

	// 图标
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`

	// 关键字
	Keywords *string `json:"Keywords,omitnil,omitempty" name:"Keywords"`

	// Qps设置，10-500，默认100
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// 子公司拓展层次
	SubCompanyLevel *int64 `json:"SubCompanyLevel,omitnil,omitempty" name:"SubCompanyLevel"`

	// 是否包含完整扫描
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsIncludeFullScan *bool `json:"IsIncludeFullScan,omitnil,omitempty" name:"IsIncludeFullScan"`

	// 是否识别集团成员
	EnableGroupMemberDiscovered *bool `json:"EnableGroupMemberDiscovered,omitnil,omitempty" name:"EnableGroupMemberDiscovered"`
}

// Predefined struct for user
type DeleteAppsRequestParams struct {
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

type DeleteAppsRequest struct {
	*tchttp.BaseRequest
	
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

func (r *DeleteAppsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAppsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAppsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAppsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAppsResponseParams `json:"Response"`
}

func (r *DeleteAppsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAssetsRequestParams struct {
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

type DeleteAssetsRequest struct {
	*tchttp.BaseRequest
	
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

func (r *DeleteAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAssetsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAssetsResponseParams `json:"Response"`
}

func (r *DeleteAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainsRequestParams struct {
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

type DeleteDomainsRequest struct {
	*tchttp.BaseRequest
	
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

func (r *DeleteDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDomainsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDomainsResponseParams `json:"Response"`
}

func (r *DeleteDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEnterprisesRequestParams struct {
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

type DeleteEnterprisesRequest struct {
	*tchttp.BaseRequest
	
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

func (r *DeleteEnterprisesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEnterprisesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEnterprisesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEnterprisesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteEnterprisesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEnterprisesResponseParams `json:"Response"`
}

func (r *DeleteEnterprisesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEnterprisesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHttpsRequestParams struct {
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`

	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`
}

type DeleteHttpsRequest struct {
	*tchttp.BaseRequest
	
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`

	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`
}

func (r *DeleteHttpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHttpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	delete(f, "CustomerIdList")
	delete(f, "IsAggregation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteHttpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHttpsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteHttpsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteHttpsResponseParams `json:"Response"`
}

func (r *DeleteHttpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHttpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteManagesRequestParams struct {
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`

	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`
}

type DeleteManagesRequest struct {
	*tchttp.BaseRequest
	
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`

	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`
}

func (r *DeleteManagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteManagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	delete(f, "CustomerIdList")
	delete(f, "IsAggregation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteManagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteManagesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteManagesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteManagesResponseParams `json:"Response"`
}

func (r *DeleteManagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteManagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePortsRequestParams struct {
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`

	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`
}

type DeletePortsRequest struct {
	*tchttp.BaseRequest
	
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`

	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`
}

func (r *DeletePortsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePortsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	delete(f, "CustomerIdList")
	delete(f, "IsAggregation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePortsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePortsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePortsResponse struct {
	*tchttp.BaseResponse
	Response *DeletePortsResponseParams `json:"Response"`
}

func (r *DeletePortsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePortsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSeedsRequestParams struct {
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

type DeleteSeedsRequest struct {
	*tchttp.BaseRequest
	
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

func (r *DeleteSeedsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSeedsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSeedsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSeedsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSeedsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSeedsResponseParams `json:"Response"`
}

func (r *DeleteSeedsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSeedsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSubDomainsRequestParams struct {
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`

	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`
}

type DeleteSubDomainsRequest struct {
	*tchttp.BaseRequest
	
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`

	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`
}

func (r *DeleteSubDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSubDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	delete(f, "CustomerIdList")
	delete(f, "IsAggregation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSubDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSubDomainsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSubDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSubDomainsResponseParams `json:"Response"`
}

func (r *DeleteSubDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSubDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSuspiciousAssetsRequestParams struct {
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`

	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`
}

type DeleteSuspiciousAssetsRequest struct {
	*tchttp.BaseRequest
	
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`

	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`
}

func (r *DeleteSuspiciousAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSuspiciousAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	delete(f, "CustomerIdList")
	delete(f, "IsAggregation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSuspiciousAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSuspiciousAssetsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSuspiciousAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSuspiciousAssetsResponseParams `json:"Response"`
}

func (r *DeleteSuspiciousAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSuspiciousAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWechatAppletsRequestParams struct {
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

type DeleteWechatAppletsRequest struct {
	*tchttp.BaseRequest
	
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

func (r *DeleteWechatAppletsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWechatAppletsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWechatAppletsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWechatAppletsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWechatAppletsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWechatAppletsResponseParams `json:"Response"`
}

func (r *DeleteWechatAppletsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWechatAppletsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWechatOfficialAccountsRequestParams struct {
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

type DeleteWechatOfficialAccountsRequest struct {
	*tchttp.BaseRequest
	
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

func (r *DeleteWechatOfficialAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWechatOfficialAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWechatOfficialAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWechatOfficialAccountsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWechatOfficialAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWechatOfficialAccountsResponseParams `json:"Response"`
}

func (r *DeleteWechatOfficialAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWechatOfficialAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiSecsRequestParams struct {
	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`
}

type DescribeApiSecsRequest struct {
	*tchttp.BaseRequest
	
	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`
}

func (r *DescribeApiSecsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiSecsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsAggregation")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Ignored")
	delete(f, "UpdateAtEnd")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "CreateAtStart")
	delete(f, "Format")
	delete(f, "IsNew")
	delete(f, "CustomerIdList")
	delete(f, "EnterpriseUidList")
	delete(f, "Filters")
	delete(f, "CustomerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiSecsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiSecsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// API安全数组
	List []*DisplayApiSec `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApiSecsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiSecsResponseParams `json:"Response"`
}

func (r *DescribeApiSecsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiSecsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppsRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeAppsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeAppsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "EnterpriseUidList")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 移动资产数组
	List []*DisplayApp `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppsResponseParams `json:"Response"`
}

func (r *DescribeAppsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetsRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssetsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 主机资产数组
	List []*DisplayAsset `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssetsResponseParams `json:"Response"`
}

func (r *DescribeAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigsRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`

	// 支持按照响应长度排序，例如：+ContentLength或-ContentLength，+是递增，-是递减
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
}

type DescribeConfigsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`

	// 支持按照响应长度排序，例如：+ContentLength或-ContentLength，+是递增，-是递减
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
}

func (r *DescribeConfigsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsAggregation")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	delete(f, "OrderBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	List []*DisplayConfig `json:"List,omitnil,omitempty" name:"List"`

	// 全部路径数量
	AllConfigNum *int64 `json:"AllConfigNum,omitnil,omitempty" name:"AllConfigNum"`

	// 高风险路径数量
	HighRiskConfigNum *int64 `json:"HighRiskConfigNum,omitnil,omitempty" name:"HighRiskConfigNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigsResponseParams `json:"Response"`
}

func (r *DescribeConfigsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomersRequestParams struct {
	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 企业名称模糊搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type DescribeCustomersRequest struct {
	*tchttp.BaseRequest
	
	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 企业名称模糊搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *DescribeCustomersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomersResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 企业列表
	List []*Customer `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomersResponseParams `json:"Response"`
}

func (r *DescribeCustomersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDarkWebsRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeDarkWebsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeDarkWebsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDarkWebsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDarkWebsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDarkWebsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	List []*DisplayDarkWeb `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDarkWebsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDarkWebsResponseParams `json:"Response"`
}

func (r *DescribeDarkWebsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDarkWebsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainsRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	List []*DisplayDomain `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainsResponseParams `json:"Response"`
}

func (r *DescribeDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnterprisesRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`

	// 是否展示统计数据
	IsShowStatistics *bool `json:"IsShowStatistics,omitnil,omitempty" name:"IsShowStatistics"`
}

type DescribeEnterprisesRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`

	// 是否展示统计数据
	IsShowStatistics *bool `json:"IsShowStatistics,omitnil,omitempty" name:"IsShowStatistics"`
}

func (r *DescribeEnterprisesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnterprisesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	delete(f, "IsShowStatistics")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnterprisesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnterprisesResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	List []*DisplayEnterprise `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEnterprisesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnterprisesResponseParams `json:"Response"`
}

func (r *DescribeEnterprisesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnterprisesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFakeAppsRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeFakeAppsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeFakeAppsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFakeAppsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFakeAppsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFakeAppsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 仿冒应用
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*DisplayFakeApp `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFakeAppsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFakeAppsResponseParams `json:"Response"`
}

func (r *DescribeFakeAppsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFakeAppsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFakeMiniProgramsRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeFakeMiniProgramsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeFakeMiniProgramsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFakeMiniProgramsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFakeMiniProgramsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFakeMiniProgramsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 仿冒小程序
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*DisplayFakeMiniProgram `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFakeMiniProgramsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFakeMiniProgramsResponseParams `json:"Response"`
}

func (r *DescribeFakeMiniProgramsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFakeMiniProgramsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFakeWebsitesRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeFakeWebsitesRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeFakeWebsitesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFakeWebsitesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFakeWebsitesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFakeWebsitesResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 仿冒网站
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*DisplayFakeWebsite `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFakeWebsitesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFakeWebsitesResponseParams `json:"Response"`
}

func (r *DescribeFakeWebsitesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFakeWebsitesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFakeWechatOfficialsRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeFakeWechatOfficialsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeFakeWechatOfficialsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFakeWechatOfficialsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFakeWechatOfficialsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFakeWechatOfficialsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 仿冒公众号
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*DisplayFakeWechatOfficial `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFakeWechatOfficialsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFakeWechatOfficialsResponseParams `json:"Response"`
}

func (r *DescribeFakeWechatOfficialsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFakeWechatOfficialsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGithubsRequestParams struct {
	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeGithubsRequest struct {
	*tchttp.BaseRequest
	
	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeGithubsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGithubsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGithubsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGithubsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	List []*DisplayGithub `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGithubsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGithubsResponseParams `json:"Response"`
}

func (r *DescribeGithubsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGithubsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHttpsRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`

	// 是否显示变更
	IsShowChange *bool `json:"IsShowChange,omitnil,omitempty" name:"IsShowChange"`

	// 是否仅显示过期风险资产
	HasExpirationRisk *bool `json:"HasExpirationRisk,omitnil,omitempty" name:"HasExpirationRisk"`

	// 是否只查询离线网站
	OnlyOffline *bool `json:"OnlyOffline,omitnil,omitempty" name:"OnlyOffline"`
}

type DescribeHttpsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`

	// 是否显示变更
	IsShowChange *bool `json:"IsShowChange,omitnil,omitempty" name:"IsShowChange"`

	// 是否仅显示过期风险资产
	HasExpirationRisk *bool `json:"HasExpirationRisk,omitnil,omitempty" name:"HasExpirationRisk"`

	// 是否只查询离线网站
	OnlyOffline *bool `json:"OnlyOffline,omitnil,omitempty" name:"OnlyOffline"`
}

func (r *DescribeHttpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHttpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsAggregation")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	delete(f, "IsShowChange")
	delete(f, "HasExpirationRisk")
	delete(f, "OnlyOffline")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHttpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHttpsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	List []*DisplayHttp `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHttpsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHttpsResponseParams `json:"Response"`
}

func (r *DescribeHttpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHttpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobRecordDetailsRequestParams struct {
	// 数据类型，包括：enterprise(企业架构)，domain(主域名)，sub_domain(子域名)，asset(主机资产)，port(端口服务)，http(网站资产)，vul(漏洞信息)，app(APP)，wechat_applet(微信小程序)，wechat_official_account(微信公众号)，github(Github泄露)，manage(后台识别)，config(目录爆破)，dark_web(暗网泄露)，net_disk(文库网盘泄露)，social_engineering(员工信息)，supply_chain(供应链信息)，weak_password(弱口令)，sensitive_info(敏感信息泄露)，suspicious_asset(影子资产)
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 结果id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 任务id
	JobRecordId *int64 `json:"JobRecordId,omitnil,omitempty" name:"JobRecordId"`
}

type DescribeJobRecordDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 数据类型，包括：enterprise(企业架构)，domain(主域名)，sub_domain(子域名)，asset(主机资产)，port(端口服务)，http(网站资产)，vul(漏洞信息)，app(APP)，wechat_applet(微信小程序)，wechat_official_account(微信公众号)，github(Github泄露)，manage(后台识别)，config(目录爆破)，dark_web(暗网泄露)，net_disk(文库网盘泄露)，social_engineering(员工信息)，supply_chain(供应链信息)，weak_password(弱口令)，sensitive_info(敏感信息泄露)，suspicious_asset(影子资产)
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 结果id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 任务id
	JobRecordId *int64 `json:"JobRecordId,omitnil,omitempty" name:"JobRecordId"`
}

func (r *DescribeJobRecordDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobRecordDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Id")
	delete(f, "JobRecordId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobRecordDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobRecordDetailsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	List []*DisplayJobRecordDetail `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeJobRecordDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobRecordDetailsResponseParams `json:"Response"`
}

func (r *DescribeJobRecordDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobRecordDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobRecordsRequestParams struct {
	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeJobRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeJobRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeJobRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeJobRecordsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	List []*DisplayJobRecord `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeJobRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeJobRecordsResponseParams `json:"Response"`
}

func (r *DescribeJobRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeJobRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLeakageCodesRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeLeakageCodesRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeLeakageCodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLeakageCodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsAggregation")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLeakageCodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLeakageCodesResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*DisplayLeakageCode `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLeakageCodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLeakageCodesResponseParams `json:"Response"`
}

func (r *DescribeLeakageCodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLeakageCodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLeakageDatasRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeLeakageDatasRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeLeakageDatasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLeakageDatasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsAggregation")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLeakageDatasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLeakageDatasResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*DisplayLeakageData `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLeakageDatasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLeakageDatasResponseParams `json:"Response"`
}

func (r *DescribeLeakageDatasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLeakageDatasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLeakageEmailsRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeLeakageEmailsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeLeakageEmailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLeakageEmailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsAggregation")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLeakageEmailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLeakageEmailsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*DisplayLeakageEmail `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLeakageEmailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLeakageEmailsResponseParams `json:"Response"`
}

func (r *DescribeLeakageEmailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLeakageEmailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeManagesRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeManagesRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeManagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeManagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsAggregation")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeManagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeManagesResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	List []*DisplayManage `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeManagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeManagesResponseParams `json:"Response"`
}

func (r *DescribeManagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeManagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetDisksRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeNetDisksRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeNetDisksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetDisksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetDisksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetDisksResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	List []*DisplayNetDisk `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNetDisksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetDisksResponseParams `json:"Response"`
}

func (r *DescribeNetDisksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetDisksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePortsRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribePortsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribePortsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePortsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsAggregation")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePortsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePortsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	List []*DisplayPort `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePortsResponse struct {
	*tchttp.BaseResponse
	Response *DescribePortsResponseParams `json:"Response"`
}

func (r *DescribePortsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePortsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSeedsRequestParams struct {
	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeSeedsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeSeedsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSeedsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerId")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSeedsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSeedsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 种子列表
	List []*DisplaySeed `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSeedsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSeedsResponseParams `json:"Response"`
}

func (r *DescribeSeedsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSeedsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSensitiveInfosRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeSensitiveInfosRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeSensitiveInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSensitiveInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsAggregation")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSensitiveInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSensitiveInfosResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	List []*DisplaySensitiveInfo `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSensitiveInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSensitiveInfosResponseParams `json:"Response"`
}

func (r *DescribeSensitiveInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSensitiveInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubDomainsRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`

	// 是否只查询离线子域名
	OnlyOffline *bool `json:"OnlyOffline,omitnil,omitempty" name:"OnlyOffline"`
}

type DescribeSubDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`

	// 是否只查询离线子域名
	OnlyOffline *bool `json:"OnlyOffline,omitnil,omitempty" name:"OnlyOffline"`
}

func (r *DescribeSubDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "IsAggregation")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	delete(f, "OnlyOffline")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubDomainsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	List []*DisplaySubDomain `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSubDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubDomainsResponseParams `json:"Response"`
}

func (r *DescribeSubDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSuspiciousAssetsRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeSuspiciousAssetsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeSuspiciousAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSuspiciousAssetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsAggregation")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSuspiciousAssetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSuspiciousAssetsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	List []*DisplaySuspiciousAsset `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSuspiciousAssetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSuspiciousAssetsResponseParams `json:"Response"`
}

func (r *DescribeSuspiciousAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSuspiciousAssetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulsRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeVulsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeVulsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVulsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVulsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	List []*DisplayVul `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVulsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVulsResponseParams `json:"Response"`
}

func (r *DescribeVulsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVulsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeakPasswordsRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeWeakPasswordsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeWeakPasswordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeakPasswordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWeakPasswordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWeakPasswordsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	List []*DisplayWeakPassword `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWeakPasswordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWeakPasswordsResponseParams `json:"Response"`
}

func (r *DescribeWeakPasswordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWeakPasswordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWechatAppletsRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeWechatAppletsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeWechatAppletsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWechatAppletsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "CustomerId")
	delete(f, "IsNew")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWechatAppletsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWechatAppletsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	List []*DisplayWechatApplet `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWechatAppletsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWechatAppletsResponseParams `json:"Response"`
}

func (r *DescribeWechatAppletsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWechatAppletsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWechatOfficialAccountsRequestParams struct {
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

type DescribeWechatOfficialAccountsRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否新增数据
	IsNew *bool `json:"IsNew,omitnil,omitempty" name:"IsNew"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分页大小
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 子公司ID列表
	EnterpriseUidList []*string `json:"EnterpriseUidList,omitnil,omitempty" name:"EnterpriseUidList"`

	// 数据输出格式：json、file，默认不填为json
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建时间-开始
	CreateAtStart *string `json:"CreateAtStart,omitnil,omitempty" name:"CreateAtStart"`

	// 创建时间-结束
	CreateAtEnd *string `json:"CreateAtEnd,omitnil,omitempty" name:"CreateAtEnd"`

	// 更新时间-开始
	UpdateAtStart *string `json:"UpdateAtStart,omitnil,omitempty" name:"UpdateAtStart"`

	// 更新时间-结束
	UpdateAtEnd *string `json:"UpdateAtEnd,omitnil,omitempty" name:"UpdateAtEnd"`

	// 查询数组
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`
}

func (r *DescribeWechatOfficialAccountsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWechatOfficialAccountsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerIdList")
	delete(f, "IsNew")
	delete(f, "CustomerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "EnterpriseUidList")
	delete(f, "Format")
	delete(f, "CreateAtStart")
	delete(f, "CreateAtEnd")
	delete(f, "UpdateAtStart")
	delete(f, "UpdateAtEnd")
	delete(f, "Filters")
	delete(f, "Ignored")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWechatOfficialAccountsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWechatOfficialAccountsResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 数组
	List []*DisplayWechatOfficialAccount `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWechatOfficialAccountsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWechatOfficialAccountsResponseParams `json:"Response"`
}

func (r *DescribeWechatOfficialAccountsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWechatOfficialAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisplayApiSec struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// Url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// Host地址
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Path路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 方法：POST、GET、DELETE等
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 修复状态：unrepaired:未修复，repaired:已修复, ignore:已忽略,checking:复测中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 请求体
	Request *string `json:"Request,omitnil,omitempty" name:"Request"`

	// 响应体
	Response *string `json:"Response,omitnil,omitempty" name:"Response"`

	// 是否风险API
	IsRiskAPI *bool `json:"IsRiskAPI,omitnil,omitempty" name:"IsRiskAPI"`
}

type DisplayApp struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// APP名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 开发者
	Developer *string `json:"Developer,omitnil,omitempty" name:"Developer"`

	// 下载地址
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 图片
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// 包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 服务端地址
	ServerUrl *string `json:"ServerUrl,omitnil,omitempty" name:"ServerUrl"`

	// app版本
	AppVersion *string `json:"AppVersion,omitnil,omitempty" name:"AppVersion"`
}

type DisplayAsset struct {
	// 主机资产Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 操作系统类型
	Os *string `json:"Os,omitnil,omitempty" name:"Os"`

	// 主机地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 国家
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// 省份
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// 地域
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// 运营商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// 端口数据
	Ports *string `json:"Ports,omitnil,omitempty" name:"Ports"`

	// 服务数据
	Services *string `json:"Services,omitnil,omitempty" name:"Services"`

	// 域名数据
	Domains *string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 端口和服务最近更新时间
	LastModify *string `json:"LastModify,omitnil,omitempty" name:"LastModify"`

	// 是否为云资产
	IsCloudAsset *int64 `json:"IsCloudAsset,omitnil,omitempty" name:"IsCloudAsset"`

	// 云资产状态，-1为下线
	CloudAssetStatus *int64 `json:"CloudAssetStatus,omitnil,omitempty" name:"CloudAssetStatus"`
}

type DisplayConfig struct {
	// 主键Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 地址
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 站点标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 状态码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 响应长度
	ContentLength *int64 `json:"ContentLength,omitnil,omitempty" name:"ContentLength"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// Ip数据
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// AI评分
	AIRating *int64 `json:"AIRating,omitnil,omitempty" name:"AIRating"`

	// AI分析
	AIAnalysis *string `json:"AIAnalysis,omitnil,omitempty" name:"AIAnalysis"`

	// 风险等级: 1-低危, 2-中危, 3-高危, 4-危级, 5-误报
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 建议
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 是否为云资产
	IsCloudAsset *int64 `json:"IsCloudAsset,omitnil,omitempty" name:"IsCloudAsset"`

	// 云资产状态，-1为下线
	CloudAssetStatus *int64 `json:"CloudAssetStatus,omitnil,omitempty" name:"CloudAssetStatus"`
}

type DisplayDarkWeb struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 命中关键字
	MatchedKeywords *string `json:"MatchedKeywords,omitnil,omitempty" name:"MatchedKeywords"`

	// APP地址
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// 状态：unrepaired:未修复，repaired:已修复，ignore:已忽略
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type DisplayDomain struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 主域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// ICP
	ICP *string `json:"ICP,omitnil,omitempty" name:"ICP"`

	// 注册时间
	RegisteredTime *string `json:"RegisteredTime,omitnil,omitempty" name:"RegisteredTime"`

	// 过期时间
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 公司
	Company *string `json:"Company,omitnil,omitempty" name:"Company"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// 是否为云资产
	IsCloudAsset *int64 `json:"IsCloudAsset,omitnil,omitempty" name:"IsCloudAsset"`

	// 云资产状态，-1为下线
	CloudAssetStatus *int64 `json:"CloudAssetStatus,omitnil,omitempty" name:"CloudAssetStatus"`
}

type DisplayEnterprise struct {
	// 主键Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// 子公司上级
	ParentEnterpriseUid *string `json:"ParentEnterpriseUid,omitnil,omitempty" name:"ParentEnterpriseUid"`

	// 企业名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 公司简称
	Abbreviation *string `json:"Abbreviation,omitnil,omitempty" name:"Abbreviation"`

	// 统一社会信用代码
	CreditCode *string `json:"CreditCode,omitnil,omitempty" name:"CreditCode"`

	// 企业状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 注册资本
	RegisteredCapital *string `json:"RegisteredCapital,omitnil,omitempty" name:"RegisteredCapital"`

	// 持股比例
	ShareholdingRatio *string `json:"ShareholdingRatio,omitnil,omitempty" name:"ShareholdingRatio"`

	// 法人代表
	LegalPerson *string `json:"LegalPerson,omitnil,omitempty" name:"LegalPerson"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 行业类型
	Industry *string `json:"Industry,omitnil,omitempty" name:"Industry"`

	// 子公司唯一uid
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`

	// 主域名数量
	DomainCount *int64 `json:"DomainCount,omitnil,omitempty" name:"DomainCount"`

	// 子域名数量
	SubDomainCount *int64 `json:"SubDomainCount,omitnil,omitempty" name:"SubDomainCount"`

	// 网站数量
	HttpCount *int64 `json:"HttpCount,omitnil,omitempty" name:"HttpCount"`

	// 漏洞数量
	VulCount *int64 `json:"VulCount,omitnil,omitempty" name:"VulCount"`
}

type DisplayFakeApp struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// 仿冒应用名称
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 仿冒应用包名称
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 下载链接
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 处置状态：0-待处理 1-处理中 2-已处理
	HandlingStatus *int64 `json:"HandlingStatus,omitnil,omitempty" name:"HandlingStatus"`

	// 关停状态：0-(默认状态) 1-关停审核中 2-已拦截 3-已拒绝 4-下线流程中 5-已下线 6-下线失败
	ShutdownStatus *int64 `json:"ShutdownStatus,omitnil,omitempty" name:"ShutdownStatus"`

	// 关停时间
	ShutdownTime *string `json:"ShutdownTime,omitnil,omitempty" name:"ShutdownTime"`
}

type DisplayFakeMiniProgram struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// 仿冒小程序名称
	ProgramName *string `json:"ProgramName,omitnil,omitempty" name:"ProgramName"`

	// 小程序ID
	ProgramId *string `json:"ProgramId,omitnil,omitempty" name:"ProgramId"`

	// 类别
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 二维码
	QrCode *string `json:"QrCode,omitnil,omitempty" name:"QrCode"`

	// 处置状态：0-待处理 1-处理中 2-已处理
	HandlingStatus *int64 `json:"HandlingStatus,omitnil,omitempty" name:"HandlingStatus"`

	// 关停状态：0-(默认状态) 1-关停审核中 2-已拦截 3-已拒绝 4-下线流程中 5-已下线 6-下线失败
	ShutdownStatus *int64 `json:"ShutdownStatus,omitnil,omitempty" name:"ShutdownStatus"`

	// 关停时间
	ShutdownTime *string `json:"ShutdownTime,omitnil,omitempty" name:"ShutdownTime"`
}

type DisplayFakeWebsite struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// 仿冒网站
	Website *string `json:"Website,omitnil,omitempty" name:"Website"`

	// ip位置
	IPLocation *string `json:"IPLocation,omitnil,omitempty" name:"IPLocation"`

	// 截图
	Screenshot *string `json:"Screenshot,omitnil,omitempty" name:"Screenshot"`

	// 处置状态：0-待处理 1-处理中 2-已处理
	HandlingStatus *int64 `json:"HandlingStatus,omitnil,omitempty" name:"HandlingStatus"`

	// 关停状态：0-(默认状态) 1-关停审核中 2-已拦截 3-已拒绝 4-下线流程中 5-已下线 6-下线失败
	ShutdownStatus *int64 `json:"ShutdownStatus,omitnil,omitempty" name:"ShutdownStatus"`

	// 关停时间
	ShutdownTime *string `json:"ShutdownTime,omitnil,omitempty" name:"ShutdownTime"`
}

type DisplayFakeWechatOfficial struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// 仿冒公众号名称
	AccountName *string `json:"AccountName,omitnil,omitempty" name:"AccountName"`

	// 公众号ID
	WechatId *string `json:"WechatId,omitnil,omitempty" name:"WechatId"`

	// 头像
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 二维码
	QrCode *string `json:"QrCode,omitnil,omitempty" name:"QrCode"`

	// 处置状态：0-待处理 1-处理中 2-已处理
	HandlingStatus *int64 `json:"HandlingStatus,omitnil,omitempty" name:"HandlingStatus"`

	// 关停状态：0-(默认状态) 1-关停审核中 2-已拦截 3-已拒绝 4-下线流程中 5-已下线 6-下线失败
	ShutdownStatus *int64 `json:"ShutdownStatus,omitnil,omitempty" name:"ShutdownStatus"`

	// 关停时间
	ShutdownTime *string `json:"ShutdownTime,omitnil,omitempty" name:"ShutdownTime"`
}

type DisplayGithub struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 命中关键字
	MatchedKeywords *string `json:"MatchedKeywords,omitnil,omitempty" name:"MatchedKeywords"`

	// 泄露地址
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`
}

type DisplayHttp struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// Url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 报文长度
	ContentLength *int64 `json:"ContentLength,omitnil,omitempty" name:"ContentLength"`

	// 报文内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 截图缩略图URL
	ScreenshotThumbUrl *string `json:"ScreenshotThumbUrl,omitnil,omitempty" name:"ScreenshotThumbUrl"`

	// 截图URL
	ScreenshotUrl *string `json:"ScreenshotUrl,omitnil,omitempty" name:"ScreenshotUrl"`

	// 状态码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// Api地址
	Api *string `json:"Api,omitnil,omitempty" name:"Api"`

	// 解析的IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 证书信息
	Ssl *string `json:"Ssl,omitnil,omitempty" name:"Ssl"`

	// ssl证书过期时间
	SslExpiredTime *string `json:"SslExpiredTime,omitnil,omitempty" name:"SslExpiredTime"`

	// 资产是否发生变动
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsChange *bool `json:"IsChange,omitnil,omitempty" name:"IsChange"`

	// 是否为云资产：0-非云资产 1-是云资产
	IsCloudAsset *int64 `json:"IsCloudAsset,omitnil,omitempty" name:"IsCloudAsset"`

	// 云资产是否下线：-1-已下线 0-正常
	CloudAssetStatus *int64 `json:"CloudAssetStatus,omitnil,omitempty" name:"CloudAssetStatus"`

	// 可用率（百分比）
	AvailabilityRate *int64 `json:"AvailabilityRate,omitnil,omitempty" name:"AvailabilityRate"`

	// 可用状态 1:异常 0:正常
	AvailabilityState *int64 `json:"AvailabilityState,omitnil,omitempty" name:"AvailabilityState"`

	// 平均响应时间：单位ms
	ResponseTime *int64 `json:"ResponseTime,omitnil,omitempty" name:"ResponseTime"`

	// 域名解析状态 1:异常 0:正常
	AnalysisState *int64 `json:"AnalysisState,omitnil,omitempty" name:"AnalysisState"`
}

type DisplayJobRecord struct {
	// 任务Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 企业名称
	CustomerName *string `json:"CustomerName,omitnil,omitempty" name:"CustomerName"`

	// 周期任务详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Crontab *string `json:"Crontab,omitnil,omitempty" name:"Crontab"`

	// 状态：2-错误/已停止，3-进行中，1-已完成, 4-停止
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 新增数据
	NewCount *int64 `json:"NewCount,omitnil,omitempty" name:"NewCount"`

	// 创建时间
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// 更新时间
	UpdateAt *string `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`

	// 子任务进度
	Progress *JobRecordProgress `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 并发设置
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 客户Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 客户appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type DisplayJobRecordDetail struct {
	// 发现时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeAt *string `json:"TimeAt,omitnil,omitempty" name:"TimeAt"`

	// 模块
	// 注意：此字段可能返回 null，表示取不到有效值。
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 模块名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleName *string `json:"ModuleName,omitnil,omitempty" name:"ModuleName"`

	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobRecordId *int64 `json:"JobRecordId,omitnil,omitempty" name:"JobRecordId"`

	// 目标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*IdndValue `json:"Data,omitnil,omitempty" name:"Data"`
}

type DisplayLeakageCode struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// 事件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 事件描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 风险等级：1-低危 2-中危 3-高危 4-严重 5-误报
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 仓库名称
	HubName *string `json:"HubName,omitnil,omitempty" name:"HubName"`

	// 链接
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 截图
	Screenshot *string `json:"Screenshot,omitnil,omitempty" name:"Screenshot"`

	// 建议
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 处置状态：0-待处理 1-处理中 2-已处理
	HandlingStatus *int64 `json:"HandlingStatus,omitnil,omitempty" name:"HandlingStatus"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type DisplayLeakageData struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// 事件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 事件描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 风险等级：1-低危 2-中危 3-高危 4-严重 5-误报
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 链接
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 截图
	Screenshot *string `json:"Screenshot,omitnil,omitempty" name:"Screenshot"`

	// 建议
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 处置状态：0-待处理 1-处理中 2-已处理
	HandlingStatus *int64 `json:"HandlingStatus,omitnil,omitempty" name:"HandlingStatus"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type DisplayLeakageEmail struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 数据源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 风险等级：1-低危 2-中危 3-高危 4-严重 5-误报
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 建议
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 处置状态：0-待处理 1-处理中 2-已处理
	HandlingStatus *int64 `json:"HandlingStatus,omitnil,omitempty" name:"HandlingStatus"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type DisplayManage struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// Url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// Icon
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`

	// 缩略图
	Screenshot *string `json:"Screenshot,omitnil,omitempty" name:"Screenshot"`

	// 状态码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// 后台Host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 状态：not_converged:未收敛, converged:已收敛, ignore:已忽略
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否为云资产：0-非云资产 1-是云资产
	IsCloudAsset *int64 `json:"IsCloudAsset,omitnil,omitempty" name:"IsCloudAsset"`

	// 云资产是否下线：-1-已下线 0-正常
	CloudAssetStatus *int64 `json:"CloudAssetStatus,omitnil,omitempty" name:"CloudAssetStatus"`
}

type DisplayNetDisk struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 命中关键字
	MatchedKeywords *string `json:"MatchedKeywords,omitnil,omitempty" name:"MatchedKeywords"`

	// 泄露地址
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// 泄露平台
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`
}

type DisplayPort struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// IP或域名地址
	Asset *string `json:"Asset,omitnil,omitempty" name:"Asset"`

	// 解析的IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 是否高危
	IsHighRisk *bool `json:"IsHighRisk,omitnil,omitempty" name:"IsHighRisk"`

	// 组件名称
	App *string `json:"App,omitnil,omitempty" name:"App"`

	// 服务名称
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 端口响应详情
	Banner *string `json:"Banner,omitnil,omitempty" name:"Banner"`

	// 上次检测时间
	LastCheckTime *string `json:"LastCheckTime,omitnil,omitempty" name:"LastCheckTime"`

	// 状态，close:连接超时，端口可能已关闭，open:端口开放, checking:复测中, ignore:已忽略
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否为云资产：0-非云资产 1-是云资产
	IsCloudAsset *int64 `json:"IsCloudAsset,omitnil,omitempty" name:"IsCloudAsset"`

	// 云资产是否下线：-1-已下线 0-正常
	CloudAssetStatus *int64 `json:"CloudAssetStatus,omitnil,omitempty" name:"CloudAssetStatus"`

	// 域名解析状态 1:异常 0:正常
	AnalysisState *int64 `json:"AnalysisState,omitnil,omitempty" name:"AnalysisState"`
}

type DisplaySeed struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 分类，包括：domain(主域名)、icon(图标)、ip(IP)、
	// keyword(关键词)、parent_company(母公司)、sub_domain(子域名)、title(标题)
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// md5值
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`

	// 来源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 创建时间
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// 是否可信
	IsValid *bool `json:"IsValid,omitnil,omitempty" name:"IsValid"`
}

type DisplaySensitiveInfo struct {
	// 主键Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 来源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// 是否为云资产：0-非云资产 1-是云资产
	IsCloudAsset *int64 `json:"IsCloudAsset,omitnil,omitempty" name:"IsCloudAsset"`

	// 云资产是否下线：-1-已下线 0-正常
	CloudAssetStatus *int64 `json:"CloudAssetStatus,omitnil,omitempty" name:"CloudAssetStatus"`
}

type DisplaySubDomain struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 子域名
	SubDomain *string `json:"SubDomain,omitnil,omitempty" name:"SubDomain"`

	// Ip
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 国家
	Country *string `json:"Country,omitnil,omitempty" name:"Country"`

	// 省份
	Province *string `json:"Province,omitnil,omitempty" name:"Province"`

	// 城市
	City *string `json:"City,omitnil,omitempty" name:"City"`

	// 互联网服务提供商
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// 是否为云资产：0-非云资产 1-是云资产
	IsCloudAsset *int64 `json:"IsCloudAsset,omitnil,omitempty" name:"IsCloudAsset"`

	// 云资产是否下线：-1-已下线 0-正常
	CloudAssetStatus *int64 `json:"CloudAssetStatus,omitnil,omitempty" name:"CloudAssetStatus"`

	// 可用率（百分比）
	AvailabilityRate *int64 `json:"AvailabilityRate,omitnil,omitempty" name:"AvailabilityRate"`

	// 可用状态 1:异常 0:正常
	AvailabilityState *int64 `json:"AvailabilityState,omitnil,omitempty" name:"AvailabilityState"`

	// 域名解析状态 1:异常 0:正常
	AnalysisState *int64 `json:"AnalysisState,omitnil,omitempty" name:"AnalysisState"`

	// 平均时延：单位ms
	AverageDelay *int64 `json:"AverageDelay,omitnil,omitempty" name:"AverageDelay"`

	// 丢包率（百分比）
	LossRate *int64 `json:"LossRate,omitnil,omitempty" name:"LossRate"`
}

type DisplaySuspiciousAsset struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// Url
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 报文长度
	ContentLength *int64 `json:"ContentLength,omitnil,omitempty" name:"ContentLength"`

	// 报文内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 截图缩略图URL
	ScreenshotThumbUrl *string `json:"ScreenshotThumbUrl,omitnil,omitempty" name:"ScreenshotThumbUrl"`

	// 截图URL
	ScreenshotUrl *string `json:"ScreenshotUrl,omitnil,omitempty" name:"ScreenshotUrl"`

	// 状态码
	Code *int64 `json:"Code,omitnil,omitempty" name:"Code"`

	// Api
	Api *string `json:"Api,omitnil,omitempty" name:"Api"`

	// 解析的IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 证书信息
	Ssl *string `json:"Ssl,omitnil,omitempty" name:"Ssl"`

	// ssl证书过期时间
	SslExpiredTime *string `json:"SslExpiredTime,omitnil,omitempty" name:"SslExpiredTime"`

	// 来源类型
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 来源值
	SourceValue *string `json:"SourceValue,omitnil,omitempty" name:"SourceValue"`

	// 是否信任
	Trusted *bool `json:"Trusted,omitnil,omitempty" name:"Trusted"`

	// 所属者
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 根域名
	RootDomain *string `json:"RootDomain,omitnil,omitempty" name:"RootDomain"`
}

type DisplayToolCommon struct {
	// 子公司ID
	EnterpriseUid *string `json:"EnterpriseUid,omitnil,omitempty" name:"EnterpriseUid"`

	// 子公司名称
	EnterpriseName *string `json:"EnterpriseName,omitnil,omitempty" name:"EnterpriseName"`

	// 主任务ID
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 单任务ID
	JobStageId *int64 `json:"JobStageId,omitnil,omitempty" name:"JobStageId"`

	// 是否忽略
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`

	// 子任务ID
	JobRecordId *int64 `json:"JobRecordId,omitnil,omitempty" name:"JobRecordId"`

	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 企业名称
	CustomerName *string `json:"CustomerName,omitnil,omitempty" name:"CustomerName"`

	// 详情
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// Md5值
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`

	// 创建时间
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// 更新时间
	UpdateAt *string `json:"UpdateAt,omitnil,omitempty" name:"UpdateAt"`

	// 标签列表，json格式：{\"tag1\":[\"责任人xxx\"],\"tag2\":[\"测试站\"]}
	Labels *string `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type DisplayVul struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// 解析的IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Url地址
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 风险等级：1：提示, 2:低危, 3:中危, 4:高危, 5: 严重
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 漏洞名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否修复，0:未修复，1:已修复
	RepairStatus *int64 `json:"RepairStatus,omitnil,omitempty" name:"RepairStatus"`

	// 建议
	Suggestion *string `json:"Suggestion,omitnil,omitempty" name:"Suggestion"`

	// 发现时间
	DiscoverTime *string `json:"DiscoverTime,omitnil,omitempty" name:"DiscoverTime"`

	// AI研判
	AiJudge *string `json:"AiJudge,omitnil,omitempty" name:"AiJudge"`

	// 状态：unrepaired:未修复，repaired:已修复, offline:资产已下线, ignore:已忽略
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 上次复测时间
	LastCheckTime *string `json:"LastCheckTime,omitnil,omitempty" name:"LastCheckTime"`

	// 是否为云资产：0-非云资产 1-是云资产
	IsCloudAsset *int64 `json:"IsCloudAsset,omitnil,omitempty" name:"IsCloudAsset"`

	// 云资产是否下线：-1-已下线 0-正常
	CloudAssetStatus *int64 `json:"CloudAssetStatus,omitnil,omitempty" name:"CloudAssetStatus"`

	// 域名解析状态 1:异常 0:正常
	AnalysisState *int64 `json:"AnalysisState,omitnil,omitempty" name:"AnalysisState"`
}

type DisplayWeakPassword struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// 解析的IP
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// Url地址
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 弱口令类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 弱口令账号
	Account *string `json:"Account,omitnil,omitempty" name:"Account"`

	// 弱口令密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 是否为蜜罐
	IsHoneypot *bool `json:"IsHoneypot,omitnil,omitempty" name:"IsHoneypot"`

	// 截图
	ScreenshotUrl *string `json:"ScreenshotUrl,omitnil,omitempty" name:"ScreenshotUrl"`

	// 状态：unrepaired:未修复，repaired:已修复, offline:资产已下线, ignore:已忽略, checking:复测中
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 上次复测时间
	LastCheckTime *string `json:"LastCheckTime,omitnil,omitempty" name:"LastCheckTime"`

	// 是否为云资产：0-非云资产 1-是云资产
	IsCloudAsset *int64 `json:"IsCloudAsset,omitnil,omitempty" name:"IsCloudAsset"`

	// 云资产是否下线：-1-已下线 0-正常
	CloudAssetStatus *int64 `json:"CloudAssetStatus,omitnil,omitempty" name:"CloudAssetStatus"`
}

type DisplayWechatApplet struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 图片地址
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// 账号
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`

	// 二维码
	QrCode *string `json:"QrCode,omitnil,omitempty" name:"QrCode"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 认证主体
	RecordSubject *string `json:"RecordSubject,omitnil,omitempty" name:"RecordSubject"`

	// 账号Appid
	AccountAppid *string `json:"AccountAppid,omitnil,omitempty" name:"AccountAppid"`
}

type DisplayWechatOfficialAccount struct {
	// 主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 公共字段
	DisplayToolCommon *DisplayToolCommon `json:"DisplayToolCommon,omitnil,omitempty" name:"DisplayToolCommon"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 图片地址
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`

	// 账号
	AccountId *string `json:"AccountId,omitnil,omitempty" name:"AccountId"`

	// 二维码
	QrCode *string `json:"QrCode,omitnil,omitempty" name:"QrCode"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 认证主体
	RecordSubject *string `json:"RecordSubject,omitnil,omitempty" name:"RecordSubject"`
}

type Filter struct {
	// 要搜索的字段
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 目标值列表
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type IdndValue struct {
	// 详情ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 目标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type IgnoreDataRequestParams struct {
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`

	// 模块，包括：enterprise：企业架构，domain：主域名，sub_domain：子域名，asset：IP资产，port：端口服务，http：HTTP资产，vul：漏洞信息，app：APP，wechat_applet：微信小程序，wechat_official_account：微信公众号，github：Github信息泄露，manage：管理后台暴露，config：目录爆破，dark_web：暗网泄露，net_disk：文库网盘泄露，supply_chain：供应链，weak_password：弱口令，sensitive_info：敏感信息泄露
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`

	// 是否忽略全部
	IsAffectOther *bool `json:"IsAffectOther,omitnil,omitempty" name:"IsAffectOther"`
}

type IgnoreDataRequest struct {
	*tchttp.BaseRequest
	
	// ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`

	// 模块，包括：enterprise：企业架构，domain：主域名，sub_domain：子域名，asset：IP资产，port：端口服务，http：HTTP资产，vul：漏洞信息，app：APP，wechat_applet：微信小程序，wechat_official_account：微信公众号，github：Github信息泄露，manage：管理后台暴露，config：目录爆破，dark_web：暗网泄露，net_disk：文库网盘泄露，supply_chain：供应链，weak_password：弱口令，sensitive_info：敏感信息泄露
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 是否显示被忽略的数据
	Ignored *bool `json:"Ignored,omitnil,omitempty" name:"Ignored"`

	// 是否忽略全部
	IsAffectOther *bool `json:"IsAffectOther,omitnil,omitempty" name:"IsAffectOther"`
}

func (r *IgnoreDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnoreDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	delete(f, "Module")
	delete(f, "CustomerIdList")
	delete(f, "IsAggregation")
	delete(f, "Ignored")
	delete(f, "IsAffectOther")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IgnoreDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IgnoreDataResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IgnoreDataResponse struct {
	*tchttp.BaseResponse
	Response *IgnoreDataResponseParams `json:"Response"`
}

func (r *IgnoreDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnoreDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JobRecordProgress struct {
	// 正在执行的任务数
	Doing *int64 `json:"Doing,omitnil,omitempty" name:"Doing"`

	// 已完成的任务数
	Done *int64 `json:"Done,omitnil,omitempty" name:"Done"`

	// 发生错误的任务数
	Error *int64 `json:"Error,omitnil,omitempty" name:"Error"`

	// 超时
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 停止
	Stop *int64 `json:"Stop,omitnil,omitempty" name:"Stop"`

	// 暂停
	Todo *int64 `json:"Todo,omitnil,omitempty" name:"Todo"`
}

// Predefined struct for user
type ModifyCustomerRequestParams struct {
	// 企业名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 百分比取值范围为30-100
	Percent *int64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 资产收集、漏洞信息、弱口令、目录爆破、暗网泄露、Github泄露、文库网盘泄露、敏感信息泄露，其中资产收集必包含，多个用英文逗号隔离，例如：资产收集,漏洞信息
	ScanType *string `json:"ScanType,omitnil,omitempty" name:"ScanType"`

	// 企业ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 周期测绘时间
	ScanCron *string `json:"ScanCron,omitnil,omitempty" name:"ScanCron"`

	// 是否立即启动
	IsScanNow *bool `json:"IsScanNow,omitnil,omitempty" name:"IsScanNow"`

	// 是否启用周期测绘
	EnableCron *bool `json:"EnableCron,omitnil,omitempty" name:"EnableCron"`

	// 是否扫描子公司
	EnableScanSubEnterprise *bool `json:"EnableScanSubEnterprise,omitnil,omitempty" name:"EnableScanSubEnterprise"`

	// 是否授权
	EnableAuth *bool `json:"EnableAuth,omitnil,omitempty" name:"EnableAuth"`

	// 授权开始时间
	AuthStartAt *string `json:"AuthStartAt,omitnil,omitempty" name:"AuthStartAt"`

	// 授权结束时间
	AuthEndAt *string `json:"AuthEndAt,omitnil,omitempty" name:"AuthEndAt"`

	// 授权文件id
	AuthFile *string `json:"AuthFile,omitnil,omitempty" name:"AuthFile"`

	// 测绘时间配置项，采用json字符串格式
	ScanTime *string `json:"ScanTime,omitnil,omitempty" name:"ScanTime"`

	// 企业图标
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`

	// 并发
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// 子公司拓展层次
	SubCompanyLevel *int64 `json:"SubCompanyLevel,omitnil,omitempty" name:"SubCompanyLevel"`

	// 是否包含完整的扫描
	IsIncludeFullScan *bool `json:"IsIncludeFullScan,omitnil,omitempty" name:"IsIncludeFullScan"`
}

type ModifyCustomerRequest struct {
	*tchttp.BaseRequest
	
	// 企业名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 百分比取值范围为30-100
	Percent *int64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 资产收集、漏洞信息、弱口令、目录爆破、暗网泄露、Github泄露、文库网盘泄露、敏感信息泄露，其中资产收集必包含，多个用英文逗号隔离，例如：资产收集,漏洞信息
	ScanType *string `json:"ScanType,omitnil,omitempty" name:"ScanType"`

	// 企业ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 周期测绘时间
	ScanCron *string `json:"ScanCron,omitnil,omitempty" name:"ScanCron"`

	// 是否立即启动
	IsScanNow *bool `json:"IsScanNow,omitnil,omitempty" name:"IsScanNow"`

	// 是否启用周期测绘
	EnableCron *bool `json:"EnableCron,omitnil,omitempty" name:"EnableCron"`

	// 是否扫描子公司
	EnableScanSubEnterprise *bool `json:"EnableScanSubEnterprise,omitnil,omitempty" name:"EnableScanSubEnterprise"`

	// 是否授权
	EnableAuth *bool `json:"EnableAuth,omitnil,omitempty" name:"EnableAuth"`

	// 授权开始时间
	AuthStartAt *string `json:"AuthStartAt,omitnil,omitempty" name:"AuthStartAt"`

	// 授权结束时间
	AuthEndAt *string `json:"AuthEndAt,omitnil,omitempty" name:"AuthEndAt"`

	// 授权文件id
	AuthFile *string `json:"AuthFile,omitnil,omitempty" name:"AuthFile"`

	// 测绘时间配置项，采用json字符串格式
	ScanTime *string `json:"ScanTime,omitnil,omitempty" name:"ScanTime"`

	// 企业图标
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`

	// 并发
	Qps *int64 `json:"Qps,omitnil,omitempty" name:"Qps"`

	// 子公司拓展层次
	SubCompanyLevel *int64 `json:"SubCompanyLevel,omitnil,omitempty" name:"SubCompanyLevel"`

	// 是否包含完整的扫描
	IsIncludeFullScan *bool `json:"IsIncludeFullScan,omitnil,omitempty" name:"IsIncludeFullScan"`
}

func (r *ModifyCustomerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Percent")
	delete(f, "ScanType")
	delete(f, "Id")
	delete(f, "ScanCron")
	delete(f, "IsScanNow")
	delete(f, "EnableCron")
	delete(f, "EnableScanSubEnterprise")
	delete(f, "EnableAuth")
	delete(f, "AuthStartAt")
	delete(f, "AuthEndAt")
	delete(f, "AuthFile")
	delete(f, "ScanTime")
	delete(f, "Icon")
	delete(f, "Qps")
	delete(f, "SubCompanyLevel")
	delete(f, "IsIncludeFullScan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomerResponseParams struct {
	// 企业ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCustomerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomerResponseParams `json:"Response"`
}

func (r *ModifyCustomerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLabelRequestParams struct {
	// 模块，包括：enterprise：企业架构，domain：主域名，sub_domain：子域名，asset：IP资产，port：端口服务，http：HTTP资产，vul：漏洞信息，app：APP，wechat_applet：微信小程序，wechat_official_account：微信公众号，github：Github信息泄露，manage：管理后台暴露，config：目录爆破，dark_web：暗网泄露，net_disk：文库网盘泄露，supply_chain：供应链，weak_password：弱口令，sensitive_info：敏感信息泄露
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 资产或风险主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 企业ID，在企业管理页面查看
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 标签详情
	Labels *string `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 资产或风险主键ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

type ModifyLabelRequest struct {
	*tchttp.BaseRequest
	
	// 模块，包括：enterprise：企业架构，domain：主域名，sub_domain：子域名，asset：IP资产，port：端口服务，http：HTTP资产，vul：漏洞信息，app：APP，wechat_applet：微信小程序，wechat_official_account：微信公众号，github：Github信息泄露，manage：管理后台暴露，config：目录爆破，dark_web：暗网泄露，net_disk：文库网盘泄露，supply_chain：供应链，weak_password：弱口令，sensitive_info：敏感信息泄露
	Module *string `json:"Module,omitnil,omitempty" name:"Module"`

	// 企业ID列表，可多选
	CustomerIdList []*int64 `json:"CustomerIdList,omitnil,omitempty" name:"CustomerIdList"`

	// 资产或风险主键ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 企业ID，在企业管理页面查看
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 是否聚合数据
	IsAggregation *bool `json:"IsAggregation,omitnil,omitempty" name:"IsAggregation"`

	// 标签详情
	Labels *string `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 资产或风险主键ID列表
	Ids []*int64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

func (r *ModifyLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "CustomerIdList")
	delete(f, "Id")
	delete(f, "CustomerId")
	delete(f, "IsAggregation")
	delete(f, "Labels")
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLabelResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLabelResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLabelResponseParams `json:"Response"`
}

func (r *ModifyLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySeedStatusRequestParams struct {
	// ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 是否可信
	IsValid *bool `json:"IsValid,omitnil,omitempty" name:"IsValid"`
}

type ModifySeedStatusRequest struct {
	*tchttp.BaseRequest
	
	// ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 是否可信
	IsValid *bool `json:"IsValid,omitnil,omitempty" name:"IsValid"`
}

func (r *ModifySeedStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySeedStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "IsValid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySeedStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySeedStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySeedStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifySeedStatusResponseParams `json:"Response"`
}

func (r *ModifySeedStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySeedStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopJobRecordRequestParams struct {
	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 任务ID
	JobRecordId *int64 `json:"JobRecordId,omitnil,omitempty" name:"JobRecordId"`
}

type StopJobRecordRequest struct {
	*tchttp.BaseRequest
	
	// 企业ID
	CustomerId *int64 `json:"CustomerId,omitnil,omitempty" name:"CustomerId"`

	// 任务ID
	JobRecordId *int64 `json:"JobRecordId,omitnil,omitempty" name:"JobRecordId"`
}

func (r *StopJobRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopJobRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CustomerId")
	delete(f, "JobRecordId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopJobRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopJobRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopJobRecordResponse struct {
	*tchttp.BaseResponse
	Response *StopJobRecordResponseParams `json:"Response"`
}

func (r *StopJobRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopJobRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}