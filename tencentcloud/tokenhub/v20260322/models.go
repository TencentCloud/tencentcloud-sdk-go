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

package v20260322

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ApiKeyDetail struct {
	// API 密钥 ID。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`

	// 名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// API 密钥值。接口返回脱敏值
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 平台类型。当前支持填值：maas
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 主账号。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 子账号。
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// 状态。取值：enable（启用）、disable（禁用）。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 绑定类型。取值：all（全部模型和服务）、model_all_endpoint_custom（全部模型+自定义服务）、model_custom_endpoint_all（自定义模型+全部服务）、model_custom_endpoint_custom（自定义模型+自定义服务）。
	BindType *string `json:"BindType,omitnil,omitempty" name:"BindType"`

	// 创建时间。格式：YYYY-MM-DD HH:mm:ss。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。格式：YYYY-MM-DD HH:mm:ss。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 应用 ID。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 是否可编辑。true 表示可编辑，false 表示不可编辑。
	Editable *bool `json:"Editable,omitnil,omitempty" name:"Editable"`

	// 绑定资源列表，区分 endpoint 和 model 类型。
	BindingItems []*BindingItem `json:"BindingItems,omitnil,omitempty" name:"BindingItems"`

	// IP 白名单列表。支持 IPv4 和 CIDR 格式。空数组表示不限制 IP。
	IpWhitelist []*string `json:"IpWhitelist,omitnil,omitempty" name:"IpWhitelist"`

	// 当Platform为maas时该字段为空
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// Token 限额信息多维度列表。未配置限额时不返回该字段。
	QuotaSet []*QuotaInfo `json:"QuotaSet,omitnil,omitempty" name:"QuotaSet"`

	// Token 限额状态。空字符串表示未配置任何限额包；active 表示已配置且当前可用；inactive 表示已配置但额度耗尽
	QuotaStatus *string `json:"QuotaStatus,omitnil,omitempty" name:"QuotaStatus"`
}

type BatchCreateFailedItem struct {
	// 失败项的序号（从 1 开始，对应后缀编号）。
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 失败项的名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 失败原因。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type BindingItem struct {
	// 资源 ID（模型 ID 或服务 ID）。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 资源类型。取值：endpoint（服务）、model（模型）。
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type CreateApiKeysResultItem struct {
	// APIKey ID。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`
}

// Predefined struct for user
type CreateGlossaryEntriesRequestParams struct {
	// 术语库 ID。可通过 DescribeGlossaries 接口获取。
	GlossaryId *string `json:"GlossaryId,omitnil,omitempty" name:"GlossaryId"`

	// 术语条目列表。单次最多 100 个。
	Entries []*GlossaryEntryInput `json:"Entries,omitnil,omitempty" name:"Entries"`
}

type CreateGlossaryEntriesRequest struct {
	*tchttp.BaseRequest
	
	// 术语库 ID。可通过 DescribeGlossaries 接口获取。
	GlossaryId *string `json:"GlossaryId,omitnil,omitempty" name:"GlossaryId"`

	// 术语条目列表。单次最多 100 个。
	Entries []*GlossaryEntryInput `json:"Entries,omitnil,omitempty" name:"Entries"`
}

func (r *CreateGlossaryEntriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlossaryEntriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlossaryId")
	delete(f, "Entries")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGlossaryEntriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGlossaryEntriesResponseParams struct {
	// 创建成功的术语条目列表。
	Entries []*GlossaryEntryItem `json:"Entries,omitnil,omitempty" name:"Entries"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGlossaryEntriesResponse struct {
	*tchttp.BaseResponse
	Response *CreateGlossaryEntriesResponseParams `json:"Response"`
}

func (r *CreateGlossaryEntriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlossaryEntriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGlossaryRequestParams struct {
	// 术语库名称。最大 50 字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 源语言代码。最大 16 字符。例如：zh（中文）、en（英文）。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言代码。最大 16 字符。例如：zh（中文）、en（英文）。
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 术语库描述。最大 255 字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateGlossaryRequest struct {
	*tchttp.BaseRequest
	
	// 术语库名称。最大 50 字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 源语言代码。最大 16 字符。例如：zh（中文）、en（英文）。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言代码。最大 16 字符。例如：zh（中文）、en（英文）。
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 术语库描述。最大 255 字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateGlossaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlossaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Source")
	delete(f, "Target")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGlossaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGlossaryResponseParams struct {
	// 术语库 ID。
	GlossaryId *string `json:"GlossaryId,omitnil,omitempty" name:"GlossaryId"`

	// 术语库名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 创建时间。Unix 时间戳（毫秒）。
	CreatedAt *int64 `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGlossaryResponse struct {
	*tchttp.BaseResponse
	Response *CreateGlossaryResponseParams `json:"Response"`
}

func (r *CreateGlossaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGlossaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTokenPlanApiKeysRequestParams struct {
	// 套餐 ID。可通过DescribeTokenPlanList接口获取。
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// API Key 名称，最大 128 字符。如果创建 API Key 的数量超过1个，实际名称格式为 {ApiKeyName}-{序号}（如 mykey-1, mykey-2）。
	ApiKeyName *string `json:"ApiKeyName,omitnil,omitempty" name:"ApiKeyName"`

	// 创建数量。取值范围：1 ~ 10。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 可用模型列表。如果套餐类型为企业版专业套餐，可以指定模型，也可以传入“all”，传入 “all“表示可以使用套餐支持的所有模型，如果要指定具体模型，传入 Model ID，“all“和具体的 Model ID 不能同时传入，如果不传则表示该API Key不支持任何模型，从而影响API Key的正常使用。如果套餐类型为企业版轻享套餐，则无论是否传入该字段，以及该字段传入什么值，都会被强制覆盖为["auto"]。
	AllowedModels []*string `json:"AllowedModels,omitnil,omitempty" name:"AllowedModels"`

	// 独占额度。不传则为0，代表该 API Key 不被分配任何独占配额。单位说明如下：
	// - 套餐类型为专业，单位取值为积分；
	// - 套餐类型为轻享套餐，单位取值为 token。
	ExclusiveQuota *int64 `json:"ExclusiveQuota,omitnil,omitempty" name:"ExclusiveQuota"`

	// 总额度上限。-1 表示不限，必须为 -1 或 >= API Key 当前的 ExclusiveQuota（独占额度）。不传表示不设置上限。单位说明如下：
	// - 套餐类型为专业，单位取值为积分；
	// - 套餐类型为轻享套餐，单位取值为 token。
	TotalQuota *int64 `json:"TotalQuota,omitnil,omitempty" name:"TotalQuota"`

	// TPM（Tokens Per Minute）限制。不传使用套餐级别 TPM。必须 >= 0 且 <= 套餐 TPM。
	TPM *int64 `json:"TPM,omitnil,omitempty" name:"TPM"`
}

type CreateTokenPlanApiKeysRequest struct {
	*tchttp.BaseRequest
	
	// 套餐 ID。可通过DescribeTokenPlanList接口获取。
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// API Key 名称，最大 128 字符。如果创建 API Key 的数量超过1个，实际名称格式为 {ApiKeyName}-{序号}（如 mykey-1, mykey-2）。
	ApiKeyName *string `json:"ApiKeyName,omitnil,omitempty" name:"ApiKeyName"`

	// 创建数量。取值范围：1 ~ 10。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 可用模型列表。如果套餐类型为企业版专业套餐，可以指定模型，也可以传入“all”，传入 “all“表示可以使用套餐支持的所有模型，如果要指定具体模型，传入 Model ID，“all“和具体的 Model ID 不能同时传入，如果不传则表示该API Key不支持任何模型，从而影响API Key的正常使用。如果套餐类型为企业版轻享套餐，则无论是否传入该字段，以及该字段传入什么值，都会被强制覆盖为["auto"]。
	AllowedModels []*string `json:"AllowedModels,omitnil,omitempty" name:"AllowedModels"`

	// 独占额度。不传则为0，代表该 API Key 不被分配任何独占配额。单位说明如下：
	// - 套餐类型为专业，单位取值为积分；
	// - 套餐类型为轻享套餐，单位取值为 token。
	ExclusiveQuota *int64 `json:"ExclusiveQuota,omitnil,omitempty" name:"ExclusiveQuota"`

	// 总额度上限。-1 表示不限，必须为 -1 或 >= API Key 当前的 ExclusiveQuota（独占额度）。不传表示不设置上限。单位说明如下：
	// - 套餐类型为专业，单位取值为积分；
	// - 套餐类型为轻享套餐，单位取值为 token。
	TotalQuota *int64 `json:"TotalQuota,omitnil,omitempty" name:"TotalQuota"`

	// TPM（Tokens Per Minute）限制。不传使用套餐级别 TPM。必须 >= 0 且 <= 套餐 TPM。
	TPM *int64 `json:"TPM,omitnil,omitempty" name:"TPM"`
}

func (r *CreateTokenPlanApiKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTokenPlanApiKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TeamId")
	delete(f, "ApiKeyName")
	delete(f, "Count")
	delete(f, "AllowedModels")
	delete(f, "ExclusiveQuota")
	delete(f, "TotalQuota")
	delete(f, "TPM")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTokenPlanApiKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTokenPlanApiKeysResponseParams struct {
	// 成功创建的项列表。
	Items []*CreateApiKeysResultItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 创建失败的项列表。
	FailedItems []*BatchCreateFailedItem `json:"FailedItems,omitnil,omitempty" name:"FailedItems"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTokenPlanApiKeysResponse struct {
	*tchttp.BaseResponse
	Response *CreateTokenPlanApiKeysResponseParams `json:"Response"`
}

func (r *CreateTokenPlanApiKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTokenPlanApiKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTokenPlanTeamOrderAndBuyRequestParams struct {
	// 套餐类型。取值：enterprise（企业版专业套餐）、enterprise-auto（企业版轻享套餐）。
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 套餐名称。只能包含中文、字母、数字、下划线、连字符，以中文或者字母开头，以中文或字母或数字结尾，2~50个字符
	TeamName *string `json:"TeamName,omitnil,omitempty" name:"TeamName"`

	// 购买时长。单位：月。必须大于 0。
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 购买的套餐规格。套餐类型为专业套餐（enterprise），单位取值为积分；轻享套餐（enterprise-auto），单位取值为 tokens。
	CreditOrToken *int64 `json:"CreditOrToken,omitnil,omitempty" name:"CreditOrToken"`

	// 是否开启自动续费。默认不开启。
	EnableAutoRenew *bool `json:"EnableAutoRenew,omitnil,omitempty" name:"EnableAutoRenew"`
}

type CreateTokenPlanTeamOrderAndBuyRequest struct {
	*tchttp.BaseRequest
	
	// 套餐类型。取值：enterprise（企业版专业套餐）、enterprise-auto（企业版轻享套餐）。
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 套餐名称。只能包含中文、字母、数字、下划线、连字符，以中文或者字母开头，以中文或字母或数字结尾，2~50个字符
	TeamName *string `json:"TeamName,omitnil,omitempty" name:"TeamName"`

	// 购买时长。单位：月。必须大于 0。
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 购买的套餐规格。套餐类型为专业套餐（enterprise），单位取值为积分；轻享套餐（enterprise-auto），单位取值为 tokens。
	CreditOrToken *int64 `json:"CreditOrToken,omitnil,omitempty" name:"CreditOrToken"`

	// 是否开启自动续费。默认不开启。
	EnableAutoRenew *bool `json:"EnableAutoRenew,omitnil,omitempty" name:"EnableAutoRenew"`
}

func (r *CreateTokenPlanTeamOrderAndBuyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTokenPlanTeamOrderAndBuyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductType")
	delete(f, "TeamName")
	delete(f, "TimeSpan")
	delete(f, "CreditOrToken")
	delete(f, "EnableAutoRenew")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTokenPlanTeamOrderAndBuyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTokenPlanTeamOrderAndBuyResponseParams struct {
	// 腾讯云订单 ID。用于关联一次购买操作下的所有子订单。
	BigOrderId *string `json:"BigOrderId,omitnil,omitempty" name:"BigOrderId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTokenPlanTeamOrderAndBuyResponse struct {
	*tchttp.BaseResponse
	Response *CreateTokenPlanTeamOrderAndBuyResponseParams `json:"Response"`
}

func (r *CreateTokenPlanTeamOrderAndBuyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTokenPlanTeamOrderAndBuyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGlossaryEntriesRequestParams struct {
	// 术语库 ID。可通过 DescribeGlossaries 接口获取。
	GlossaryId *string `json:"GlossaryId,omitnil,omitempty" name:"GlossaryId"`

	// 待删除的术语条目列表。单次最多 200 个。
	Entries []*DeleteGlossaryEntryInput `json:"Entries,omitnil,omitempty" name:"Entries"`
}

type DeleteGlossaryEntriesRequest struct {
	*tchttp.BaseRequest
	
	// 术语库 ID。可通过 DescribeGlossaries 接口获取。
	GlossaryId *string `json:"GlossaryId,omitnil,omitempty" name:"GlossaryId"`

	// 待删除的术语条目列表。单次最多 200 个。
	Entries []*DeleteGlossaryEntryInput `json:"Entries,omitnil,omitempty" name:"Entries"`
}

func (r *DeleteGlossaryEntriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGlossaryEntriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlossaryId")
	delete(f, "Entries")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGlossaryEntriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGlossaryEntriesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGlossaryEntriesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGlossaryEntriesResponseParams `json:"Response"`
}

func (r *DeleteGlossaryEntriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGlossaryEntriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGlossaryEntryInput struct {
	// 术语条目 ID。可通过 DescribeGlossaryEntries 接口获取。
	EntryId *string `json:"EntryId,omitnil,omitempty" name:"EntryId"`
}

// Predefined struct for user
type DeleteGlossaryRequestParams struct {
	// 术语库 ID。可通过 DescribeGlossaries 接口获取。
	GlossaryId *string `json:"GlossaryId,omitnil,omitempty" name:"GlossaryId"`
}

type DeleteGlossaryRequest struct {
	*tchttp.BaseRequest
	
	// 术语库 ID。可通过 DescribeGlossaries 接口获取。
	GlossaryId *string `json:"GlossaryId,omitnil,omitempty" name:"GlossaryId"`
}

func (r *DeleteGlossaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGlossaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlossaryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGlossaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGlossaryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGlossaryResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGlossaryResponseParams `json:"Response"`
}

func (r *DeleteGlossaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGlossaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTokenPlanApiKeyRequestParams struct {
	// API Key ID。可通过DescribeTokenPlanApiKeyList接口获取。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`
}

type DeleteTokenPlanApiKeyRequest struct {
	*tchttp.BaseRequest
	
	// API Key ID。可通过DescribeTokenPlanApiKeyList接口获取。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`
}

func (r *DeleteTokenPlanApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTokenPlanApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTokenPlanApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTokenPlanApiKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTokenPlanApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTokenPlanApiKeyResponseParams `json:"Response"`
}

func (r *DeleteTokenPlanApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTokenPlanApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiKeyListRequestParams struct {
	// 平台类型。当前支持填值：maas
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件列表。支持的过滤字段：apikeyId（API 密钥 ID）、apiKeyName（名称）、platform（平台类型）、status（状态）、bindType（绑定类型）。
	Filters []*RequestFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序条件列表。支持的排序字段：apiKeyName
	Sorts []*RequestSort `json:"Sorts,omitnil,omitempty" name:"Sorts"`
}

type DescribeApiKeyListRequest struct {
	*tchttp.BaseRequest
	
	// 平台类型。当前支持填值：maas
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件列表。支持的过滤字段：apikeyId（API 密钥 ID）、apiKeyName（名称）、platform（平台类型）、status（状态）、bindType（绑定类型）。
	Filters []*RequestFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序条件列表。支持的排序字段：apiKeyName
	Sorts []*RequestSort `json:"Sorts,omitnil,omitempty" name:"Sorts"`
}

func (r *DescribeApiKeyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiKeyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Sorts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiKeyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiKeyListResponseParams struct {
	// API 密钥列表。
	ApiKeySet []*ApiKeyDetail `json:"ApiKeySet,omitnil,omitempty" name:"ApiKeySet"`

	// 符合条件的 API 密钥总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApiKeyListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiKeyListResponseParams `json:"Response"`
}

func (r *DescribeApiKeyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiKeyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiKeyRequestParams struct {
	// 平台类型。当前支持填值：maas
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// API 密钥 ID，与 ApiKey 至少需传入其一，优先使用 ApiKeyId。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`

	// API 密钥明文，与 ApiKeyId 至少需传入其一。
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`
}

type DescribeApiKeyRequest struct {
	*tchttp.BaseRequest
	
	// 平台类型。当前支持填值：maas
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// API 密钥 ID，与 ApiKey 至少需传入其一，优先使用 ApiKeyId。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`

	// API 密钥明文，与 ApiKeyId 至少需传入其一。
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`
}

func (r *DescribeApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Platform")
	delete(f, "ApiKeyId")
	delete(f, "ApiKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiKeyResponseParams struct {
	// API 密钥 ID。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`

	// 名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// API 密钥值（明文）。
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`

	// 备注。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 平台类型。枚举：maas
	Platform *string `json:"Platform,omitnil,omitempty" name:"Platform"`

	// 主账号。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 子账号。
	SubUin *string `json:"SubUin,omitnil,omitempty" name:"SubUin"`

	// 状态。取值：enable（启用）、disable（禁用）。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 绑定类型。取值：all（全部模型和接入点）、model_all_endpoint_custom（全部模型+自定义接入点）、model_custom_endpoint_all（自定义模型+全部接入点）、model_custom_endpoint_custom（自定义模型+自定义接入点）。
	BindType *string `json:"BindType,omitnil,omitempty" name:"BindType"`

	// 创建时间。格式：YYYY-MM-DD HH:mm:ss。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间。格式：YYYY-MM-DD HH:mm:ss。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 应用 ID。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 是否可编辑。true 表示可编辑，false 表示不可编辑。
	Editable *bool `json:"Editable,omitnil,omitempty" name:"Editable"`

	// 绑定资源列表，区分 endpoint 和 model 类型。
	BindingItems []*BindingItem `json:"BindingItems,omitnil,omitempty" name:"BindingItems"`

	// IP 白名单列表。支持 IPv4和 CIDR 格式。空数组表示不限制 IP。
	IpWhitelist []*string `json:"IpWhitelist,omitnil,omitempty" name:"IpWhitelist"`

	// 当Platform为maas时该字段为空
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// Token 限额多维度信息。未配置限额时不返回该字段。
	QuotaSet []*QuotaInfo `json:"QuotaSet,omitnil,omitempty" name:"QuotaSet"`

	// Token 限额状态。空字符串表示未配置任何限额包；active 表示已配置且当前可用；inactive 表示已配置但额度耗尽
	QuotaStatus *string `json:"QuotaStatus,omitnil,omitempty" name:"QuotaStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApiKeyResponseParams `json:"Response"`
}

func (r *DescribeApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlossariesRequestParams struct {
	// 返回数量。默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件列表。支持的过滤字段：GlossaryId（术语库 ID）、Name（名称）、Source（源语言代码）、Target（目标语言代码）。
	Filters []*RequestFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序条件列表。支持的排序字段：CreatedTime（创建时间）、UpdatedTime（更新时间）。
	Sorts []*RequestSort `json:"Sorts,omitnil,omitempty" name:"Sorts"`
}

type DescribeGlossariesRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量。默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量。默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件列表。支持的过滤字段：GlossaryId（术语库 ID）、Name（名称）、Source（源语言代码）、Target（目标语言代码）。
	Filters []*RequestFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序条件列表。支持的排序字段：CreatedTime（创建时间）、UpdatedTime（更新时间）。
	Sorts []*RequestSort `json:"Sorts,omitnil,omitempty" name:"Sorts"`
}

func (r *DescribeGlossariesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlossariesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Sorts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGlossariesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlossariesResponseParams struct {
	// 术语库列表。
	Items []*GlossaryItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 符合条件的术语库总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 当前页码。
	Current *int64 `json:"Current,omitnil,omitempty" name:"Current"`

	// 每页大小。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGlossariesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGlossariesResponseParams `json:"Response"`
}

func (r *DescribeGlossariesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlossariesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlossaryEntriesRequestParams struct {
	// 术语库 ID。可通过 DescribeGlossaries 接口获取。
	GlossaryId *string `json:"GlossaryId,omitnil,omitempty" name:"GlossaryId"`

	// 页码。默认为 1。
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页大小。默认为 20，最大值为 200。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeGlossaryEntriesRequest struct {
	*tchttp.BaseRequest
	
	// 术语库 ID。可通过 DescribeGlossaries 接口获取。
	GlossaryId *string `json:"GlossaryId,omitnil,omitempty" name:"GlossaryId"`

	// 页码。默认为 1。
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页大小。默认为 20，最大值为 200。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeGlossaryEntriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlossaryEntriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlossaryId")
	delete(f, "Page")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGlossaryEntriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGlossaryEntriesResponseParams struct {
	// 术语条目列表。
	Entries []*GlossaryEntryItem `json:"Entries,omitnil,omitempty" name:"Entries"`

	// 符合条件的术语条目总数。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 当前页码。
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页大小。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGlossaryEntriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGlossaryEntriesResponseParams `json:"Response"`
}

func (r *DescribeGlossaryEntriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGlossaryEntriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelListRequestParams struct {

}

type DescribeModelListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeModelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelListResponseParams `json:"Response"`
}

func (r *DescribeModelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenPlanApiKeyListRequestParams struct {
	// 套餐 ID。可通过DescribeTokenPlanList接口获取。
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询过滤条件列表。支持的过滤字段：ApiKeyId （API Key ID）、Name （API Key 名称）、Status （API Key是否可用）、StopReason （API Key停用原因）、UseStatus （API Key用户侧开关）。
	Filters []*RequestFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询排序条件列表。支持的排序字段：CreatedAt（创建时间）、UpdatedAt（更新时间）。默认按 CreatedAt 降序。
	Sorts []*RequestSort `json:"Sorts,omitnil,omitempty" name:"Sorts"`
}

type DescribeTokenPlanApiKeyListRequest struct {
	*tchttp.BaseRequest
	
	// 套餐 ID。可通过DescribeTokenPlanList接口获取。
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询过滤条件列表。支持的过滤字段：ApiKeyId （API Key ID）、Name （API Key 名称）、Status （API Key是否可用）、StopReason （API Key停用原因）、UseStatus （API Key用户侧开关）。
	Filters []*RequestFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页查询排序条件列表。支持的排序字段：CreatedAt（创建时间）、UpdatedAt（更新时间）。默认按 CreatedAt 降序。
	Sorts []*RequestSort `json:"Sorts,omitnil,omitempty" name:"Sorts"`
}

func (r *DescribeTokenPlanApiKeyListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenPlanApiKeyListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TeamId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Sorts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTokenPlanApiKeyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenPlanApiKeyListResponseParams struct {
	// API Key 列表。
	ApiKeySet []*TokenPlanApiKeyListItem `json:"ApiKeySet,omitnil,omitempty" name:"ApiKeySet"`

	// API Key总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTokenPlanApiKeyListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTokenPlanApiKeyListResponseParams `json:"Response"`
}

func (r *DescribeTokenPlanApiKeyListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenPlanApiKeyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenPlanApiKeyRequestParams struct {
	// API Key ID。可通过DescribeTokenPlanApiKeyList接口获取。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`
}

type DescribeTokenPlanApiKeyRequest struct {
	*tchttp.BaseRequest
	
	// API Key ID。可通过DescribeTokenPlanApiKeyList接口获取。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`
}

func (r *DescribeTokenPlanApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenPlanApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTokenPlanApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenPlanApiKeyResponseParams struct {
	// API Key 详情信息。
	ApiKey *TokenPlanApiKeyInfo `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`

	// API Key 额度及用量信息。
	Balance *SubPackageBalance `json:"Balance,omitnil,omitempty" name:"Balance"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTokenPlanApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTokenPlanApiKeyResponseParams `json:"Response"`
}

func (r *DescribeTokenPlanApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenPlanApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenPlanApiKeySecretRequestParams struct {
	// API Key ID。可通过DescribeTokenPlanApiKeyList接口获取。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`
}

type DescribeTokenPlanApiKeySecretRequest struct {
	*tchttp.BaseRequest
	
	// API Key ID。可通过DescribeTokenPlanApiKeyList接口获取。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`
}

func (r *DescribeTokenPlanApiKeySecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenPlanApiKeySecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTokenPlanApiKeySecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenPlanApiKeySecretResponseParams struct {
	// APIKey ID。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`

	// APIKey 密钥值（明文）。请妥善保管。
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTokenPlanApiKeySecretResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTokenPlanApiKeySecretResponseParams `json:"Response"`
}

func (r *DescribeTokenPlanApiKeySecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenPlanApiKeySecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenPlanApiKeyUsageDetailRequestParams struct {
	// 套餐 ID。可通过DescribeTokenPlanList接口获取。
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// 起始时间，RFC3339 格式。不传默认为结束时间前 15 分钟。
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 结束时间，RFC3339 格式。不传默认为当前时间。
	To *string `json:"To,omitnil,omitempty" name:"To"`

	// 排序方式。取值：asc（升序）、desc（降序），默认为 desc。
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 返回条数，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 翻页上下文，首次查询不传，后续传入上次返回的 Context，直到 ListOver 为 true。
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 按 API Key ID 精确过滤。最大 128 字符。与 ApiKeyName 至少需传入其一，都传时以 ApiKeyId 为准。可通过 DescribeTokenPlanApiKeyList 接口获取。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`

	// 按 API Key 名称模糊过滤。最大 64 字符。与 ApiKeyId 至少需传入其一，都传时以 ApiKeyId 为准。
	ApiKeyName *string `json:"ApiKeyName,omitnil,omitempty" name:"ApiKeyName"`

	// 按模型 ID (Model ID) 精确过滤。需要按模型名称过滤时传入该字段。
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`
}

type DescribeTokenPlanApiKeyUsageDetailRequest struct {
	*tchttp.BaseRequest
	
	// 套餐 ID。可通过DescribeTokenPlanList接口获取。
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// 起始时间，RFC3339 格式。不传默认为结束时间前 15 分钟。
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 结束时间，RFC3339 格式。不传默认为当前时间。
	To *string `json:"To,omitnil,omitempty" name:"To"`

	// 排序方式。取值：asc（升序）、desc（降序），默认为 desc。
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 返回条数，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 翻页上下文，首次查询不传，后续传入上次返回的 Context，直到 ListOver 为 true。
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 按 API Key ID 精确过滤。最大 128 字符。与 ApiKeyName 至少需传入其一，都传时以 ApiKeyId 为准。可通过 DescribeTokenPlanApiKeyList 接口获取。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`

	// 按 API Key 名称模糊过滤。最大 64 字符。与 ApiKeyId 至少需传入其一，都传时以 ApiKeyId 为准。
	ApiKeyName *string `json:"ApiKeyName,omitnil,omitempty" name:"ApiKeyName"`

	// 按模型 ID (Model ID) 精确过滤。需要按模型名称过滤时传入该字段。
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`
}

func (r *DescribeTokenPlanApiKeyUsageDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenPlanApiKeyUsageDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TeamId")
	delete(f, "From")
	delete(f, "To")
	delete(f, "Sort")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "ApiKeyId")
	delete(f, "ApiKeyName")
	delete(f, "ModelName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTokenPlanApiKeyUsageDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenPlanApiKeyUsageDetailResponseParams struct {
	// 翻页上下文，传入下一次请求的 Context 参数继续翻页。
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 是否已到末尾，为 true 时无需继续翻页。
	ListOver *bool `json:"ListOver,omitnil,omitempty" name:"ListOver"`

	// 调用明细列表。
	List []*UsageDetailItem `json:"List,omitnil,omitempty" name:"List"`

	// 	 套餐类型。取值：enterprise（企业版专业套餐）、enterprise-auto（企业版轻享套餐）
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTokenPlanApiKeyUsageDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTokenPlanApiKeyUsageDetailResponseParams `json:"Response"`
}

func (r *DescribeTokenPlanApiKeyUsageDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenPlanApiKeyUsageDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenPlanListRequestParams struct {
	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询过滤条件列表。支持的过滤字段：TeamId（套餐 ID）、Name（套餐名称）、StopReason（关停原因）、ProductType（套餐类型）。
	Filters []*RequestFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序条件列表。支持的排序字段：CreatedAt（创建时间）、UpdatedAt（更新时间）。默认按 CreatedAt 降序。
	Sorts []*RequestSort `json:"Sorts,omitnil,omitempty" name:"Sorts"`
}

type DescribeTokenPlanListRequest struct {
	*tchttp.BaseRequest
	
	// 分页查询偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页查询过滤条件列表。支持的过滤字段：TeamId（套餐 ID）、Name（套餐名称）、StopReason（关停原因）、ProductType（套餐类型）。
	Filters []*RequestFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序条件列表。支持的排序字段：CreatedAt（创建时间）、UpdatedAt（更新时间）。默认按 CreatedAt 降序。
	Sorts []*RequestSort `json:"Sorts,omitnil,omitempty" name:"Sorts"`
}

func (r *DescribeTokenPlanListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenPlanListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "Sorts")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTokenPlanListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenPlanListResponseParams struct {
	// 套餐列表。
	TokenPlanSet []*TokenPlanListItem `json:"TokenPlanSet,omitnil,omitempty" name:"TokenPlanSet"`

	// 套餐总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTokenPlanListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTokenPlanListResponseParams `json:"Response"`
}

func (r *DescribeTokenPlanListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenPlanListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenPlanRequestParams struct {
	// 套餐 ID。可通过 DescribeTokenPlanList 接口获取。
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`
}

type DescribeTokenPlanRequest struct {
	*tchttp.BaseRequest
	
	// 套餐 ID。可通过 DescribeTokenPlanList 接口获取。
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`
}

func (r *DescribeTokenPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TeamId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTokenPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenPlanResponseParams struct {
	// 套餐 ID。
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// 套餐名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 主账号 APP ID。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 主账号 UIN。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 状态。取值：enable（启用）、disable（停用）。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 关停原因。取值：取值：NORMAL（套餐正常）、ISOLATED（隔离/欠费）、FROZEN（冻结）、EXHAUSTED（额度耗尽）、DESTROYED（已销毁）。
	StopReason *string `json:"StopReason,omitnil,omitempty" name:"StopReason"`

	// 可创建APIKey 上限。
	ApiKeyMax *int64 `json:"ApiKeyMax,omitnil,omitempty" name:"ApiKeyMax"`

	// 云计费预付费资源包 ID。
	PrepayResourceID *string `json:"PrepayResourceID,omitnil,omitempty" name:"PrepayResourceID"`

	// 创建人，子账号创建的套餐将展示子账号UIN。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 创建时间。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新时间。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 套餐包基本信息
	PackageInfo *TokenPlanPackageInfo `json:"PackageInfo,omitnil,omitempty" name:"PackageInfo"`

	// 自动续费标识。取值：0（手动续费）、1（自动续费）、2（明确不自动续费）。未绑定预付费资源时不返回。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 当前已创建的 API Key 数量
	ApiKeyCount *int64 `json:"ApiKeyCount,omitnil,omitempty" name:"ApiKeyCount"`

	// 当前周期 Token 用量明细
	TokenSummary *TokenSummary `json:"TokenSummary,omitnil,omitempty" name:"TokenSummary"`

	// 套餐类型。取值：enterprise（企业版专业套餐）、enterprise-auto（企业版轻享套餐）
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTokenPlanResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTokenPlanResponseParams `json:"Response"`
}

func (r *DescribeTokenPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsageRankListRequestParams struct {
	// <p>统计维度。取值：apikey（按 APIKey 统计）、endpoint（按接入点统计）、model（按模型统计）。</p>
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// <p>起始时间（闭区间），RFC3339 格式。</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间（开区间），RFC3339 格式。与 StartTime 的跨度最大 90 天。</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>指标族切换字段。</p><ul><li>tokens（默认）：Token 消耗图（statistics=sum），支持 Dimension = apikey/endpoint/model</li><li>search【待上线】：联网搜索调用次数（statistics=sum），仅支持 Dimension = model</li><li>其他值返回 InvalidParameter。</li></ul><p>枚举值：</p><ul><li>tokens： tokens</li></ul>
	MetricType *string `json:"MetricType,omitnil,omitempty" name:"MetricType"`

	// <p>维度过滤值。空字符串表示查询全部对象，非空时仅查询指定单个对象（如指定 APIKey ID）。最大 256 字符。</p>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// <p>统计粒度（秒）。取值：60、300、3600、86400。必须不小于跨度对应下限：跨度 ≤ 1 天 → 60；1 ~ 5 天 → 300；5 ~ 10 天 → 3600；&gt; 10 天 → 86400。仅 ShowAll=false 时使用。</p>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>翻页起点，从 0 起，默认 0。ShowAll=true 时忽略。页大小固定为 10。</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>是否返回全量结果。</p><ul><li>false（默认）：按 Offset 分页返回 TopList（每页 10 条），每个对象包含<br>Series 时序点用于绘制曲线。</li><li>true：忽略 Offset，返回全量对象列表，不返回 Series（CSV 导出场景）。</li></ul>
	ShowAll *bool `json:"ShowAll,omitnil,omitempty" name:"ShowAll"`
}

type DescribeUsageRankListRequest struct {
	*tchttp.BaseRequest
	
	// <p>统计维度。取值：apikey（按 APIKey 统计）、endpoint（按接入点统计）、model（按模型统计）。</p>
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// <p>起始时间（闭区间），RFC3339 格式。</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间（开区间），RFC3339 格式。与 StartTime 的跨度最大 90 天。</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>指标族切换字段。</p><ul><li>tokens（默认）：Token 消耗图（statistics=sum），支持 Dimension = apikey/endpoint/model</li><li>search【待上线】：联网搜索调用次数（statistics=sum），仅支持 Dimension = model</li><li>其他值返回 InvalidParameter。</li></ul><p>枚举值：</p><ul><li>tokens： tokens</li></ul>
	MetricType *string `json:"MetricType,omitnil,omitempty" name:"MetricType"`

	// <p>维度过滤值。空字符串表示查询全部对象，非空时仅查询指定单个对象（如指定 APIKey ID）。最大 256 字符。</p>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// <p>统计粒度（秒）。取值：60、300、3600、86400。必须不小于跨度对应下限：跨度 ≤ 1 天 → 60；1 ~ 5 天 → 300；5 ~ 10 天 → 3600；&gt; 10 天 → 86400。仅 ShowAll=false 时使用。</p>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>翻页起点，从 0 起，默认 0。ShowAll=true 时忽略。页大小固定为 10。</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>是否返回全量结果。</p><ul><li>false（默认）：按 Offset 分页返回 TopList（每页 10 条），每个对象包含<br>Series 时序点用于绘制曲线。</li><li>true：忽略 Offset，返回全量对象列表，不返回 Series（CSV 导出场景）。</li></ul>
	ShowAll *bool `json:"ShowAll,omitnil,omitempty" name:"ShowAll"`
}

func (r *DescribeUsageRankListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsageRankListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Dimension")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MetricType")
	delete(f, "Target")
	delete(f, "Period")
	delete(f, "Offset")
	delete(f, "ShowAll")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsageRankListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsageRankListResponseParams struct {
	// <p>回填请求的统计维度。</p>
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// <p>回填请求的指标族：tokens / search 。</p>
	MetricType *string `json:"MetricType,omitnil,omitempty" name:"MetricType"`

	// <p>本次响应中 Stats / Series / PageStats / TotalStats 实际包含的 metric key 列表，按MetricType 区分：tokens=[Total,Input,Output,Cache]、search=[SearchRequestCount,SearchCount]</p>
	MetricKeys []*string `json:"MetricKeys,omitnil,omitempty" name:"MetricKeys"`

	// <p>视图（数据来源）</p>
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// <p>回填请求的统计粒度（秒）。ShowAll=true 时为 0。</p>
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>回填请求的起始时间。</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>回填请求的结束时间。</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>全量对象数。</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// <p>回填请求的翻页起点。ShowAll=true 时为 0。</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>页大小，恒为 10。ShowAll=true 时为 Total。</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>Series 数组对应的时间戳序列（Unix 秒）。ShowAll=true 时为空数组。</p>
	Timestamps []*int64 `json:"Timestamps,omitnil,omitempty" name:"Timestamps"`

	// <p>对象排行列表，按<code>MetricKeys[0]</code>降序排序。ShowAll=false 时为当前页 10 个对象（含 Series）；ShowAll=true 时为全量对象（不含 Series，用于 CSV 导出）。</p>
	TopList []*UsageRankItem `json:"TopList,omitnil,omitempty" name:"TopList"`

	// <p>分页统计结果</p>
	PageStats *UsageStats `json:"PageStats,omitnil,omitempty" name:"PageStats"`

	// <p>总统计结果</p>
	TotalStats *UsageStats `json:"TotalStats,omitnil,omitempty" name:"TotalStats"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUsageRankListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsageRankListResponseParams `json:"Response"`
}

func (r *DescribeUsageRankListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsageRankListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GlossaryEntryInput struct {
	// 源语言术语。最大 1000 字符。
	SourceTerm *string `json:"SourceTerm,omitnil,omitempty" name:"SourceTerm"`

	// 目标语言术语。最大 1000 字符。
	TargetTerm *string `json:"TargetTerm,omitnil,omitempty" name:"TargetTerm"`
}

type GlossaryEntryItem struct {
	// 术语条目 ID。
	EntryId *string `json:"EntryId,omitnil,omitempty" name:"EntryId"`

	// 源语言术语。
	SourceTerm *string `json:"SourceTerm,omitnil,omitempty" name:"SourceTerm"`

	// 目标语言术语。
	TargetTerm *string `json:"TargetTerm,omitnil,omitempty" name:"TargetTerm"`

	// 更新时间。Unix 时间戳（毫秒）。
	UpdatedAt *int64 `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`
}

type GlossaryItem struct {
	// 术语库 ID。
	GlossaryId *string `json:"GlossaryId,omitnil,omitempty" name:"GlossaryId"`

	// 术语库名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 术语库描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 源语言代码。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言代码。
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 更新时间。
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`
}

// Predefined struct for user
type ModifyGlossaryEntriesRequestParams struct {
	// 术语库 ID。可通过 DescribeGlossaries 接口获取。
	GlossaryId *string `json:"GlossaryId,omitnil,omitempty" name:"GlossaryId"`

	// 术语条目列表。单次最多 200 个。
	Entries []*ModifyGlossaryEntryInput `json:"Entries,omitnil,omitempty" name:"Entries"`
}

type ModifyGlossaryEntriesRequest struct {
	*tchttp.BaseRequest
	
	// 术语库 ID。可通过 DescribeGlossaries 接口获取。
	GlossaryId *string `json:"GlossaryId,omitnil,omitempty" name:"GlossaryId"`

	// 术语条目列表。单次最多 200 个。
	Entries []*ModifyGlossaryEntryInput `json:"Entries,omitnil,omitempty" name:"Entries"`
}

func (r *ModifyGlossaryEntriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGlossaryEntriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GlossaryId")
	delete(f, "Entries")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGlossaryEntriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGlossaryEntriesResponseParams struct {
	// 修改后的术语条目列表。
	Entries []*GlossaryEntryItem `json:"Entries,omitnil,omitempty" name:"Entries"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGlossaryEntriesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGlossaryEntriesResponseParams `json:"Response"`
}

func (r *ModifyGlossaryEntriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGlossaryEntriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGlossaryEntryInput struct {
	// 术语条目 ID。可通过 DescribeGlossaryEntries 接口获取。
	EntryId *string `json:"EntryId,omitnil,omitempty" name:"EntryId"`

	// 源语言术语。最大 1000 字符。不传则保持不变。
	SourceTerm *string `json:"SourceTerm,omitnil,omitempty" name:"SourceTerm"`

	// 目标语言术语。最大 1000 字符。不传则保持不变。
	TargetTerm *string `json:"TargetTerm,omitnil,omitempty" name:"TargetTerm"`
}

// Predefined struct for user
type ModifyTokenPlanApiKeyRequestParams struct {
	// API Key ID。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`

	// 可用模型列表。不传则不修改。
	// 
	// - 如果套餐类型为企业版专业套餐：
	// 1）传入“all“ ：使用套餐支持的所有模型
	// 2）传入 Model ID：指定具体模型。“all“和具体的 Model ID 不能同时传入。
	// 
	// - 如果套餐类型为企业版轻享套餐，不允许传该参数。
	AllowedModels []*string `json:"AllowedModels,omitnil,omitempty" name:"AllowedModels"`

	// 独占额度，不传则不修改。单位说明：
	// 
	// - 套餐类型为专业套餐，单位取值为积分；
	// - 套餐类型为轻享套餐，单位取值为 token。
	ExclusiveQuota *int64 `json:"ExclusiveQuota,omitnil,omitempty" name:"ExclusiveQuota"`

	// 总额度上限。-1 表示不限，必须为 -1 或 >= API Key 当前的 ExclusiveQuota（独占额度），不传则不修改。单位说明如下：
	// - 套餐类型为专业套餐，单位取值为积分；
	// - 套餐类型为轻享套餐，单位取值为 token。
	TotalQuota *int64 `json:"TotalQuota,omitnil,omitempty" name:"TotalQuota"`

	// 是否启用该 API Key。取值：enable（启用）、disable（停用），不传则不修改。
	UseStatus *string `json:"UseStatus,omitnil,omitempty" name:"UseStatus"`

	// TPM（Tokens Per Minute）限制。不传则不修改。必须 >= 0 且 <= 套餐 TPM。
	TPM *int64 `json:"TPM,omitnil,omitempty" name:"TPM"`
}

type ModifyTokenPlanApiKeyRequest struct {
	*tchttp.BaseRequest
	
	// API Key ID。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`

	// 可用模型列表。不传则不修改。
	// 
	// - 如果套餐类型为企业版专业套餐：
	// 1）传入“all“ ：使用套餐支持的所有模型
	// 2）传入 Model ID：指定具体模型。“all“和具体的 Model ID 不能同时传入。
	// 
	// - 如果套餐类型为企业版轻享套餐，不允许传该参数。
	AllowedModels []*string `json:"AllowedModels,omitnil,omitempty" name:"AllowedModels"`

	// 独占额度，不传则不修改。单位说明：
	// 
	// - 套餐类型为专业套餐，单位取值为积分；
	// - 套餐类型为轻享套餐，单位取值为 token。
	ExclusiveQuota *int64 `json:"ExclusiveQuota,omitnil,omitempty" name:"ExclusiveQuota"`

	// 总额度上限。-1 表示不限，必须为 -1 或 >= API Key 当前的 ExclusiveQuota（独占额度），不传则不修改。单位说明如下：
	// - 套餐类型为专业套餐，单位取值为积分；
	// - 套餐类型为轻享套餐，单位取值为 token。
	TotalQuota *int64 `json:"TotalQuota,omitnil,omitempty" name:"TotalQuota"`

	// 是否启用该 API Key。取值：enable（启用）、disable（停用），不传则不修改。
	UseStatus *string `json:"UseStatus,omitnil,omitempty" name:"UseStatus"`

	// TPM（Tokens Per Minute）限制。不传则不修改。必须 >= 0 且 <= 套餐 TPM。
	TPM *int64 `json:"TPM,omitnil,omitempty" name:"TPM"`
}

func (r *ModifyTokenPlanApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTokenPlanApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiKeyId")
	delete(f, "AllowedModels")
	delete(f, "ExclusiveQuota")
	delete(f, "TotalQuota")
	delete(f, "UseStatus")
	delete(f, "TPM")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTokenPlanApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTokenPlanApiKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTokenPlanApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTokenPlanApiKeyResponseParams `json:"Response"`
}

func (r *ModifyTokenPlanApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTokenPlanApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTokenPlanApiKeySecretRequestParams struct {
	// API Key ID。可通过DescribeTokenPlanApiKeyList接口获取。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`
}

type ModifyTokenPlanApiKeySecretRequest struct {
	*tchttp.BaseRequest
	
	// API Key ID。可通过DescribeTokenPlanApiKeyList接口获取。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`
}

func (r *ModifyTokenPlanApiKeySecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTokenPlanApiKeySecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTokenPlanApiKeySecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTokenPlanApiKeySecretResponseParams struct {
	// API Key ID。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`

	// 重置后的密钥版本号。
	KeyVersion *int64 `json:"KeyVersion,omitnil,omitempty" name:"KeyVersion"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTokenPlanApiKeySecretResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTokenPlanApiKeySecretResponseParams `json:"Response"`
}

func (r *ModifyTokenPlanApiKeySecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTokenPlanApiKeySecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QuotaInfo struct {
	// 限额包 ID。
	PkgId *string `json:"PkgId,omitnil,omitempty" name:"PkgId"`

	// 限额包状态。取值：1（正常）、3（已耗尽）、4（已销毁）。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 限额周期。取值：d（按日）、m（按月）、lifetime（总额度，不重置）。
	CycleUnit *string `json:"CycleUnit,omitnil,omitempty" name:"CycleUnit"`

	// 维度当期限额总量（Token 数）。使用字符串避免大数精度丢失。
	CycleCredits *string `json:"CycleCredits,omitnil,omitempty" name:"CycleCredits"`

	// 维度当期已使用量（Token 数）。使用字符串避免大数精度丢失。
	CycleUsed *string `json:"CycleUsed,omitnil,omitempty" name:"CycleUsed"`

	// 限额生效起始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 限额过期时间。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

// Predefined struct for user
type RenewTokenPlanTeamOrderRequestParams struct {
	// 套餐 ID。可通过 DescribeTokenPlanList 接口获取。
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// 续费时长。单位：月。必须大于 0。
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`
}

type RenewTokenPlanTeamOrderRequest struct {
	*tchttp.BaseRequest
	
	// 套餐 ID。可通过 DescribeTokenPlanList 接口获取。
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// 续费时长。单位：月。必须大于 0。
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`
}

func (r *RenewTokenPlanTeamOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewTokenPlanTeamOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TeamId")
	delete(f, "TimeSpan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewTokenPlanTeamOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewTokenPlanTeamOrderResponseParams struct {
	// 腾讯云订单 ID。用于关联一次续费操作下的所有子订单。
	BigOrderId *string `json:"BigOrderId,omitnil,omitempty" name:"BigOrderId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewTokenPlanTeamOrderResponse struct {
	*tchttp.BaseResponse
	Response *RenewTokenPlanTeamOrderResponseParams `json:"Response"`
}

func (r *RenewTokenPlanTeamOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewTokenPlanTeamOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RequestFilter struct {
	// 过滤字段名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤操作符。取值：EXACT（精确匹配）、FUZZY（模糊匹配）、NOT（排除匹配）。
	Op *string `json:"Op,omitnil,omitempty" name:"Op"`

	// 过滤值列表。最多支持 10 个。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type RequestSort struct {
	// 排序字段名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 排序方向。取值：ASC（升序）、DESC（降序）。
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type SubPackageBalance struct {
	// 独占额度。单位说明如下：
	// - 套餐类型为专业套餐，单位取值为积分；
	// - 套餐类型为轻享套餐，单位取值为 token。
	ExclusiveQuota *string `json:"ExclusiveQuota,omitnil,omitempty" name:"ExclusiveQuota"`

	// 独占额度已用量。单位说明如下：
	// - 套餐类型为专业套餐，单位取值为积分；
	// - 套餐类型为轻享套餐，单位取值为 token。
	ExclusiveUsed *string `json:"ExclusiveUsed,omitnil,omitempty" name:"ExclusiveUsed"`

	// 独占额度剩余量。单位说明如下：
	// - 套餐类型为专业套餐，单位取值为积分；
	// - 套餐类型为轻享套餐，单位取值为 token。
	ExclusiveRemain *string `json:"ExclusiveRemain,omitnil,omitempty" name:"ExclusiveRemain"`

	// 共享额度上限，-1 表示不限。单位说明如下：
	// - 套餐类型为专业套餐，单位取值为积分；
	// - 套餐类型为轻享套餐，单位取值为 token。
	SharedQuota *string `json:"SharedQuota,omitnil,omitempty" name:"SharedQuota"`

	// 共享额度已用量。单位说明如下：
	// - 套餐类型为专业套餐，单位取值为积分；
	// - 套餐类型为轻享套餐，单位取值为 token。
	SharedUsed *string `json:"SharedUsed,omitnil,omitempty" name:"SharedUsed"`

	// 共享额度剩余量。单位说明如下：
	// - 套餐类型为专业套餐，单位取值为积分；
	// - 套餐类型为轻享套餐，单位取值为 token。
	SharedRemain *string `json:"SharedRemain,omitnil,omitempty" name:"SharedRemain"`

	// API Key 额度包状态。取值：0（正常）、1（耗尽）。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type TokenPlanApiKeyInfo struct {
	// API Key ID。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`

	// API Key 密钥值（脱敏）。
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`

	// API Key 名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 所属套餐 ID。
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// 账号APP ID。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 主账号 UIN。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// API Key 可用模型列表（JSON 数组字符串）。
	AllowedModels *string `json:"AllowedModels,omitnil,omitempty" name:"AllowedModels"`

	// API Key 是否可用。取值：enable（启用）、disable（停用）。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// API Key 停用原因。取值：NORMAL（正常，默认值），QUOTA_EXHAUSTED（API Key额度包耗尽），ABNORMAL（异常，需人工介入）
	StopReason *string `json:"StopReason,omitnil,omitempty" name:"StopReason"`

	// 用户侧开关。取值：enable（启用）、disable（停用）。
	UseStatus *string `json:"UseStatus,omitnil,omitempty" name:"UseStatus"`

	// 密钥版本号。
	KeyVersion *int64 `json:"KeyVersion,omitnil,omitempty" name:"KeyVersion"`

	// 最近一次重置时间。（ISO 8601）
	LastRotatedAt *string `json:"LastRotatedAt,omitnil,omitempty" name:"LastRotatedAt"`

	// 创建人，如果是子账号创建，则该值为子账号UIN。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 创建时间。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新时间。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// TPM 限制（Tokens Per Minute）。
	TPM *int64 `json:"TPM,omitnil,omitempty" name:"TPM"`

	// 套餐类型。取值：enterprise（企业版专业套餐）、enterprise-auto（企业版轻享套餐）
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`
}

type TokenPlanApiKeyListItem struct {
	// API Key ID。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`

	// API Key 密钥值（脱敏）。
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`

	// API Key 名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 所属套餐 ID。
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// 账号APP ID。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 主账号 UIN。最大 128 字符。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// API Key 可用模型列表（JSON 数组字符串）。
	AllowedModels *string `json:"AllowedModels,omitnil,omitempty" name:"AllowedModels"`

	// API Key 是否可用。取值：enable（启用）、disable（停用）。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// API Key 停用原因。取值：NORMAL（正常，默认值），QUOTA_EXHAUSTED（API Key额度包耗尽），ABNORMAL（异常，需人工介入）
	StopReason *string `json:"StopReason,omitnil,omitempty" name:"StopReason"`

	// 用户侧开关。取值：enable（启用）、disable（停用）。
	UseStatus *string `json:"UseStatus,omitnil,omitempty" name:"UseStatus"`

	// 密钥版本号。
	KeyVersion *int64 `json:"KeyVersion,omitnil,omitempty" name:"KeyVersion"`

	// 最近一次重置时间。（ISO 8601）
	LastRotatedAt *string `json:"LastRotatedAt,omitnil,omitempty" name:"LastRotatedAt"`

	// 创建人，如果是子账号创建，则该值为子账号UIN。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 创建时间。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新时间。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// API Key 额度用量信息
	Balance *SubPackageBalance `json:"Balance,omitnil,omitempty" name:"Balance"`

	// 套餐类型。取值：enterprise（企业版专业套餐）、enterprise-auto（企业版轻享套餐）。
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`
}

type TokenPlanListItem struct {
	// 套餐 ID。
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// 套餐类型。取值：enterprise（企业版专业套餐）、enterprise-auto（企业版轻享套餐）
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 套餐名称。最大 128 字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 账号 APP ID。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 主账号 UIN。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 套餐状态。取值：enable（启用）、disable（停用）。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 套餐关停原因。取值：NORMAL（正常）、ISOLATED（隔离/欠费）、FROZEN（冻结）、EXHAUSTED（额度耗尽）、DESTROYED（已销毁）
	StopReason *string `json:"StopReason,omitnil,omitempty" name:"StopReason"`

	// 可创建 API Key 上限。
	ApiKeyMax *int64 `json:"ApiKeyMax,omitnil,omitempty" name:"ApiKeyMax"`

	// 云计费预付费资源包 ID。
	PrepayResourceID *string `json:"PrepayResourceID,omitnil,omitempty" name:"PrepayResourceID"`

	// 创建人。若为子账号创建的套餐，则该值为子账号UIN。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 创建时间。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 更新时间。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// 套餐包基本信息。
	PackageInfo *TokenPlanPackageInfo `json:"PackageInfo,omitnil,omitempty" name:"PackageInfo"`

	// 是否开启自动续费。取值：0（未开启），1（开启）
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`
}

type TokenPlanPackageInfo struct {
	// 总额度。根据套餐类型区分单位：credits（企业版专业套餐），tokens（企业版auto套餐）
	TotalQuota *string `json:"TotalQuota,omitnil,omitempty" name:"TotalQuota"`

	// 总已使用额度。根据套餐类型区分单位：credits（企业版专业套餐），tokens（企业版auto套餐）
	TotalUsed *string `json:"TotalUsed,omitnil,omitempty" name:"TotalUsed"`

	// 总周期数。
	TotalCycles *int64 `json:"TotalCycles,omitnil,omitempty" name:"TotalCycles"`

	// 周期单位。取值：month（月）
	CycleUnit *string `json:"CycleUnit,omitnil,omitempty" name:"CycleUnit"`

	// 套餐包生效时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 套餐包到期时间。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 独占池已分配额度。根据套餐类型区分单位：credits（企业版专业套餐），tokens（企业版auto套餐）
	ExclusiveAllocated *string `json:"ExclusiveAllocated,omitnil,omitempty" name:"ExclusiveAllocated"`

	// 独占池已用额度。根据套餐类型区分单位：credits（企业版专业套餐），tokens（企业版auto套餐）
	ExclusiveUsed *string `json:"ExclusiveUsed,omitnil,omitempty" name:"ExclusiveUsed"`

	// 共享池总额度。根据套餐类型区分单位：credits（企业版专业套餐），tokens（企业版auto套餐）
	SharedPool *string `json:"SharedPool,omitnil,omitempty" name:"SharedPool"`

	// 共享已用额度。根据套餐类型区分单位：credits（企业版专业套餐），tokens（企业版auto套餐）
	SharedUsed *string `json:"SharedUsed,omitnil,omitempty" name:"SharedUsed"`

	// 当期额度。根据套餐类型区分单位：credits（企业版专业套餐），tokens（企业版auto套餐）
	CycleQuota *string `json:"CycleQuota,omitnil,omitempty" name:"CycleQuota"`

	// 当前周期。
	CurrentCycle *int64 `json:"CurrentCycle,omitnil,omitempty" name:"CurrentCycle"`

	// 剩余周期。
	RemainCycles *int64 `json:"RemainCycles,omitnil,omitempty" name:"RemainCycles"`
}

type TokenSummary struct {
	// 套餐包当前计费周期序号
	CycleSeq *int64 `json:"CycleSeq,omitnil,omitempty" name:"CycleSeq"`

	// 套餐包计费周期开始时间（RFC3339）
	CycleStartTime *string `json:"CycleStartTime,omitnil,omitempty" name:"CycleStartTime"`

	// 套餐包当前计费周期结束时间（RFC3339）
	CycleEndTime *string `json:"CycleEndTime,omitnil,omitempty" name:"CycleEndTime"`

	// 按计费项分组的 token 汇总列表
	BillingItems []*TokenSummaryBillingItem `json:"BillingItems,omitnil,omitempty" name:"BillingItems"`
}

type TokenSummaryBillingItem struct {
	// 计费项。取值：input（输入 Token）、output（输出 Token）、cache（缓存 Token）、call_count（调用次数）。
	BillingItem *string `json:"BillingItem,omitnil,omitempty" name:"BillingItem"`

	// 该计费项在周期内的原始用量汇总，单位：tokens。
	TotalQty *int64 `json:"TotalQty,omitnil,omitempty" name:"TotalQty"`
}

// Predefined struct for user
type UpgradeTokenPlanTeamOrderRequestParams struct {
	// 套餐 ID。可通过 DescribeTokenPlanList 接口获取。
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// 升配后的新规格额度。套餐类型为 enterprise 时表示积分额度，套餐类型为 enterprise-auto 时表示 Token 数。必须大于当前额度。
	NewCreditOrToken *int64 `json:"NewCreditOrToken,omitnil,omitempty" name:"NewCreditOrToken"`
}

type UpgradeTokenPlanTeamOrderRequest struct {
	*tchttp.BaseRequest
	
	// 套餐 ID。可通过 DescribeTokenPlanList 接口获取。
	TeamId *string `json:"TeamId,omitnil,omitempty" name:"TeamId"`

	// 升配后的新规格额度。套餐类型为 enterprise 时表示积分额度，套餐类型为 enterprise-auto 时表示 Token 数。必须大于当前额度。
	NewCreditOrToken *int64 `json:"NewCreditOrToken,omitnil,omitempty" name:"NewCreditOrToken"`
}

func (r *UpgradeTokenPlanTeamOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeTokenPlanTeamOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TeamId")
	delete(f, "NewCreditOrToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeTokenPlanTeamOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeTokenPlanTeamOrderResponseParams struct {
	// 腾讯云订单 ID。用于关联一次升配操作下的所有子订单。
	BigOrderId *string `json:"BigOrderId,omitnil,omitempty" name:"BigOrderId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeTokenPlanTeamOrderResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeTokenPlanTeamOrderResponseParams `json:"Response"`
}

func (r *UpgradeTokenPlanTeamOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeTokenPlanTeamOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsageDetailItem struct {
	// 主账号 UIN。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 模型名称。
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// APIKey ID。
	ApiKeyId *string `json:"ApiKeyId,omitnil,omitempty" name:"ApiKeyId"`

	// APIKey 名称。
	ApiKeyName *string `json:"ApiKeyName,omitnil,omitempty" name:"ApiKeyName"`

	// 请求 ID。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// 请求时间（RFC3339 格式）。
	RequestTime *string `json:"RequestTime,omitnil,omitempty" name:"RequestTime"`

	// 输入 token 数。
	InputToken *int64 `json:"InputToken,omitnil,omitempty" name:"InputToken"`

	// 缓存 token 数。
	CacheToken *int64 `json:"CacheToken,omitnil,omitempty" name:"CacheToken"`

	// 输出 token 数。
	OutputToken *int64 `json:"OutputToken,omitnil,omitempty" name:"OutputToken"`

	// 总 token 数。
	TotalToken *int64 `json:"TotalToken,omitnil,omitempty" name:"TotalToken"`

	// 未命中缓存输入消耗额度。单位说明如下：
	// - 套餐类型为专业套餐（enterprise），单位取值为积分；
	// - 套餐类型轻享套餐（enterprise-auto），单位取值为 token。
	InputQuota *string `json:"InputQuota,omitnil,omitempty" name:"InputQuota"`

	// 缓存消耗额度。单位说明如下：
	// - 套餐类型为专业套餐（enterprise），单位取值为积分；
	// - 套餐类型轻享套餐（enterprise-auto），单位取值为 token。
	CacheQuota *string `json:"CacheQuota,omitnil,omitempty" name:"CacheQuota"`

	// 输出消耗额度。单位说明如下：
	// - 套餐类型为专业套餐（enterprise），单位取值为积分；
	// - 套餐类型轻享套餐（enterprise-auto），单位取值为 token。
	OutputQuota *string `json:"OutputQuota,omitnil,omitempty" name:"OutputQuota"`

	// 总消耗额度。单位说明如下：
	// - 套餐类型为专业套餐（enterprise），单位取值为积分；
	// - 套餐类型轻享套餐（enterprise-auto），单位取值为 token。
	TotalQuota *string `json:"TotalQuota,omitnil,omitempty" name:"TotalQuota"`
}

type UsageRankItem struct {
	// 全局排名（从 1 起）。分页场景下仍为全量排序中的位次，非页内序号
	Rank *int64 `json:"Rank,omitnil,omitempty" name:"Rank"`

	// 对象标识。apikey 维度为 APIKey ID；endpoint 维度为接入点；model 维度为模型名。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 对象展示名。apikey 维度返回 APIKey 名称（已删除的 APIKey 仍保留原名）；
	// endpoint / model 维度等于 Key。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 时间周期内的统计值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Stats *UsageStats `json:"Stats,omitnil,omitempty" name:"Stats"`

	// 时间周期内的时序点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Series *UsageSeries `json:"Series,omitnil,omitempty" name:"Series"`
}

type UsageSeries struct {
	// <p>[tokens 族]总 token 数用量时间周期内的 JSON 字符串形式，如 <code>&quot;[12,null,15]&quot;</code>。</p>
	TotalToken *string `json:"TotalToken,omitnil,omitempty" name:"TotalToken"`

	// <p>[tokens 族]输入 token 数用量时间周期内的 JSON 字符串形式，如 <code>&quot;[7,null,9]&quot;</code>。</p>
	InputTotalToken *string `json:"InputTotalToken,omitnil,omitempty" name:"InputTotalToken"`

	// <p>[tokens 族]输出 token 数用量时间周期内的 JSON 字符串形式，如 <code>&quot;[5,null,6]&quot;</code>。</p>
	OutputTotalToken *string `json:"OutputTotalToken,omitnil,omitempty" name:"OutputTotalToken"`

	// <p>[tokens 族]读缓存 token 数用量时间周期内的 JSON 字符串形式，如<code>&quot;[5,null,6]&quot;</code>。</p>
	CacheTotalToken *string `json:"CacheTotalToken,omitnil,omitempty" name:"CacheTotalToken"`

	// <p>[search 族] 搜索请求数用量时间周期内的 JSON 字符串形式，如<code>&quot;[5,null,6]&quot;</code>。</p>
	SearchRequestCount *string `json:"SearchRequestCount,omitnil,omitempty" name:"SearchRequestCount"`

	// <p>[search 族] 搜索引擎调用次数用量时间周期内的 JSON 字符串形式，如<code>&quot;[5,null,6]&quot;</code>。</p>
	SearchCount *string `json:"SearchCount,omitnil,omitempty" name:"SearchCount"`
}

type UsageStats struct {
	// <p>[tokens 族] 时间周期内的累计总 token 数。</p>
	TotalToken *int64 `json:"TotalToken,omitnil,omitempty" name:"TotalToken"`

	// <p>[tokens 族] 时间周期内的累计输入 token 数。</p>
	InputTotalToken *int64 `json:"InputTotalToken,omitnil,omitempty" name:"InputTotalToken"`

	// <p>[tokens 族] 时间周期内的累计输出 token 数。</p>
	OutputTotalToken *int64 `json:"OutputTotalToken,omitnil,omitempty" name:"OutputTotalToken"`

	// <p>[tokens 族] 时间周期内的累计读缓存 token 数（命中缓存部分）</p>
	CacheTotalToken *int64 `json:"CacheTotalToken,omitnil,omitempty" name:"CacheTotalToken"`

	// <p>[search 族] 整段累计联网搜索请求数</p>
	SearchRequestCount *int64 `json:"SearchRequestCount,omitnil,omitempty" name:"SearchRequestCount"`

	// <p>[search 族] 整段累计搜索引擎调用次数</p>
	SearchCount *int64 `json:"SearchCount,omitnil,omitempty" name:"SearchCount"`
}