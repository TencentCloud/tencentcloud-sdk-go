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

package v20220530

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type DescribeLicenseRequestParams struct {
	// 可选过滤器
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeLicenseRequest struct {
	*tchttp.BaseRequest
	
	// 可选过滤器
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLicenseResponseParams struct {
	// 针对上面raw的签名
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLicenseResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLicenseResponseParams `json:"Response"`
}

func (r *DescribeLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisplayMetadata struct {
	// <p>展示的名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>展示的值</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Filter struct {
	// 需要过滤的字段
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段的过滤值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type License struct {
	// <p>License ID</p>
	LicenseId *string `json:"LicenseId,omitnil,omitempty" name:"LicenseId"`

	// <p>软件授权模式。<table><thead><tr><th>枚举值</th><th>说明</th></tr></thead><tbody><tr><td>Permanent</td><td>永久授权。该授权不受有效期限制。</td></tr><tr><td>Subscription</td><td>订阅授权。授权如果过了有效期，则会进入过期状态。</td></tr><tr><td>Accept</td><td>验收期授权。用于需要验收的软件处于验收期间的授权，授权如果过了验收有效期，则会进入过期状态。</td></tr></tbody></table></p>
	LicenseMode *string `json:"LicenseMode,omitnil,omitempty" name:"LicenseMode"`

	// <p>软件的授权状态。<table><thead><tr><th>枚举值</th><th>说明</th></tr></thead><tbody><tr><td>Issued</td><td>已颁发，等待激活。一般来说，如果软件已经在运行，不会出现该状态。</td></tr><tr><td>Active</td><td>授权在有效期内，这是软件运行期间最常见的状态。</td></tr><tr><td>Expired</td><td>授权已过期。订阅类的软件授权有有效期，如果服务器时间已晚于有效期，则会进入过期状态。</td></tr><tr><td>Isolated</td><td>授权已隔离。有截止日期的授权，当用户授权到期时，先进入此状态，用户可以去续费，超过7天不续费则授权进入Destroyed状态。</td></tr><tr><td>Destroyed</td><td>授权已失效/销毁。用户如果退货软件，则授权会自动失效。</td></tr></tbody></table></p>
	LicenseStatus *string `json:"LicenseStatus,omitnil,omitempty" name:"LicenseStatus"`

	// <p>软件供应方 ID。</p>
	ProviderId *uint64 `json:"ProviderId,omitnil,omitempty" name:"ProviderId"`

	// <p>软件包 ID。</p>
	SoftwarePackageId *string `json:"SoftwarePackageId,omitnil,omitempty" name:"SoftwarePackageId"`

	// <p>软件包版本。</p>
	SoftwarePackageVersion *string `json:"SoftwarePackageVersion,omitnil,omitempty" name:"SoftwarePackageVersion"`

	// <p>被授权的用户 UIN。</p>
	AuthorizedUserUin *string `json:"AuthorizedUserUin,omitnil,omitempty" name:"AuthorizedUserUin"`

	// <p>被授权的应用实例 ID。</p>
	AuthorizedCloudappId *string `json:"AuthorizedCloudappId,omitnil,omitempty" name:"AuthorizedCloudappId"`

	// <p>被授权的角色 ID。</p>
	AuthorizedCloudappRoleId *string `json:"AuthorizedCloudappRoleId,omitnil,omitempty" name:"AuthorizedCloudappRoleId"`

	// <p>被授权的软件规格，具体字段请参考结构SaleParam</p>
	AuthorizedSpecification []*SaleParam `json:"AuthorizedSpecification,omitnil,omitempty" name:"AuthorizedSpecification"`

	// <p>被授权的软件的计费模式。<table><thead><tr><th>枚举值</th><th>说明</th></tr></thead><tbody><tr><td>1</td><td>线上计费，软件的授权从腾讯云线上购买，支持续费、退款等操作。</td></tr><tr><td>2</td><td>线下计费，软件的授权线下签订合同购买，定向客户交付，无法从线上续费和退款。</td></tr><tr><td>4</td><td>免费</td></tr></tbody></table></p>
	BillingMode *int64 `json:"BillingMode,omitnil,omitempty" name:"BillingMode"`

	// <p>授权时长（单位由LifeSpanUnit确定，枚举值有Y年/M月/D日三种）</p>
	LifeSpan *int64 `json:"LifeSpan,omitnil,omitempty" name:"LifeSpan"`

	// <p>授权颁发时间。</p>
	IssueDate *string `json:"IssueDate,omitnil,omitempty" name:"IssueDate"`

	// <p>授权激活时间，如从未激活则返回 null。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivationDate *string `json:"ActivationDate,omitnil,omitempty" name:"ActivationDate"`

	// <p>授权过期时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpirationDate *string `json:"ExpirationDate,omitnil,omitempty" name:"ExpirationDate"`

	// <p>授权时长单位，枚举值有Y年/M月/D日三种</p>
	LifeSpanUnit *string `json:"LifeSpanUnit,omitnil,omitempty" name:"LifeSpanUnit"`

	// <p>授权的类型：Standard正式版/Development开发版/Trial体验版</p>
	LicenseType *string `json:"LicenseType,omitnil,omitempty" name:"LicenseType"`

	// <p>授权的层级：Master 主授权；Child 子授权/增强型授权</p>
	LicenseLevel *string `json:"LicenseLevel,omitnil,omitempty" name:"LicenseLevel"`

	// <p>License 内容信息</p>
	LicenseData *LicenseData `json:"LicenseData,omitnil,omitempty" name:"LicenseData"`

	// <p>License 颁发地址</p>
	IssueURL *string `json:"IssueURL,omitnil,omitempty" name:"IssueURL"`

	// <p>商品供应商 UIN</p>
	ProviderUin *string `json:"ProviderUin,omitnil,omitempty" name:"ProviderUin"`

	// <p>License 创建来源 ID</p>
	CreateSource *string `json:"CreateSource,omitnil,omitempty" name:"CreateSource"`
}

type LicenseData struct {
	// <p>License 文本内容。支持密钥、证书等文本形式，二进制的密钥需要伙伴进行 base64 转码</p>
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// <p>部署服务输出信息，基于部署签发 License 时需要该参数。</p>
	DeploymentOutput *string `json:"DeploymentOutput,omitnil,omitempty" name:"DeploymentOutput"`

	// <p>License 前端展示信息。key、value 形式，比如可传入，颁发机构：XXXX 有限公司</p>
	Metadata []*DisplayMetadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`
}

type SaleParam struct {
	// 售卖参数标识
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// 售卖参数的展示名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamKeyName *string `json:"ParamKeyName,omitnil,omitempty" name:"ParamKeyName"`

	// 参数 Id
	ParamId *string `json:"ParamId,omitnil,omitempty" name:"ParamId"`

	// 参数值 Id
	ParamValueId *string `json:"ParamValueId,omitnil,omitempty" name:"ParamValueId"`

	// 售卖参数值，当ParamType=Quant时，该值有可能为Null
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`

	// 售卖参数值的展示名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamValueName *string `json:"ParamValueName,omitnil,omitempty" name:"ParamValueName"`

	// 售卖参数的类型，目前支持枚举类Enum/数量类Quant
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// 模块ID
	ModuleId *string `json:"ModuleId,omitnil,omitempty" name:"ModuleId"`

	// 模块key
	ModuleKey *string `json:"ModuleKey,omitnil,omitempty" name:"ModuleKey"`

	// 模块名称
	ModuleName *string `json:"ModuleName,omitnil,omitempty" name:"ModuleName"`
}

// Predefined struct for user
type VerifyLicenseRequestParams struct {

}

type VerifyLicenseRequest struct {
	*tchttp.BaseRequest
	
}

func (r *VerifyLicenseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyLicenseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyLicenseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyLicenseResponseParams struct {
	// 软件的详细授权信息。
	License *License `json:"License,omitnil,omitempty" name:"License"`

	// 当前请求服务端的时间戳，格式为RFC3339
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 对License字段对应的json数据的签名
	Signature *string `json:"Signature,omitnil,omitempty" name:"Signature"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VerifyLicenseResponse struct {
	*tchttp.BaseResponse
	Response *VerifyLicenseResponseParams `json:"Response"`
}

func (r *VerifyLicenseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyLicenseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}