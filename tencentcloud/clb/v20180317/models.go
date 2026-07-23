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

package v20180317

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddModelKeyRequestParams struct {
	// <p>服务提供商ID</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>Key 列表，至少 1 个，最多 10 个</p>
	Keys []*KeyItem `json:"Keys,omitnil,omitempty" name:"Keys"`
}

type AddModelKeyRequest struct {
	*tchttp.BaseRequest
	
	// <p>服务提供商ID</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>Key 列表，至少 1 个，最多 10 个</p>
	Keys []*KeyItem `json:"Keys,omitnil,omitempty" name:"Keys"`
}

func (r *AddModelKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddModelKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceProviderId")
	delete(f, "Keys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddModelKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddModelKeyResponseParams struct {
	// <p>生成的 Key ID 列表</p>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddModelKeyResponse struct {
	*tchttp.BaseResponse
	Response *AddModelKeyResponseParams `json:"Response"`
}

func (r *AddModelKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddModelKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddModelRewriteRequestParams struct {
	// <p>模型路由实例 ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>源模型名（重写规则的 key）。</p><p>长度 1-255 字符；支持特殊值 <code>default</code> 表示兜底规则（命中所有未显式列出的源模型）。</p><p>不允许使用 <code>IntentRouter/</code> 前缀（大小写不敏感），即 IntentRouter 不能作为 source。</p>
	SourceModel *string `json:"SourceModel,omitnil,omitempty" name:"SourceModel"`

	// <p>目标模型名（重写规则的 value）。</p><p>长度 1-255 字符；必须是已关联到该模型路由实例的模型（含 IntentRouter/* 也需先通过 AssociateModels 关联）。</p><p>不允许使用 <code>default</code>；不允许与 SourceModel 相同（大小写不敏感）。</p>
	TargetModel *string `json:"TargetModel,omitnil,omitempty" name:"TargetModel"`
}

type AddModelRewriteRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例 ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>源模型名（重写规则的 key）。</p><p>长度 1-255 字符；支持特殊值 <code>default</code> 表示兜底规则（命中所有未显式列出的源模型）。</p><p>不允许使用 <code>IntentRouter/</code> 前缀（大小写不敏感），即 IntentRouter 不能作为 source。</p>
	SourceModel *string `json:"SourceModel,omitnil,omitempty" name:"SourceModel"`

	// <p>目标模型名（重写规则的 value）。</p><p>长度 1-255 字符；必须是已关联到该模型路由实例的模型（含 IntentRouter/* 也需先通过 AssociateModels 关联）。</p><p>不允许使用 <code>default</code>；不允许与 SourceModel 相同（大小写不敏感）。</p>
	TargetModel *string `json:"TargetModel,omitnil,omitempty" name:"TargetModel"`
}

func (r *AddModelRewriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddModelRewriteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "SourceModel")
	delete(f, "TargetModel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddModelRewriteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddModelRewriteResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddModelRewriteResponse struct {
	*tchttp.BaseResponse
	Response *AddModelRewriteResponseParams `json:"Response"`
}

func (r *AddModelRewriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddModelRewriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateBudgetRequestParams struct {
	// <p>Budget ID。</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>要关联的资源列表。</p><p>仅支持企业型模型路由实例和企业型实例下的Key。同一请求内不允许重复资源；资源已关联其他Budget时将替换为新的Budget。</p>
	Resources []*BudgetResource `json:"Resources,omitnil,omitempty" name:"Resources"`
}

type AssociateBudgetRequest struct {
	*tchttp.BaseRequest
	
	// <p>Budget ID。</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>要关联的资源列表。</p><p>仅支持企业型模型路由实例和企业型实例下的Key。同一请求内不允许重复资源；资源已关联其他Budget时将替换为新的Budget。</p>
	Resources []*BudgetResource `json:"Resources,omitnil,omitempty" name:"Resources"`
}

func (r *AssociateBudgetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateBudgetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BudgetId")
	delete(f, "Resources")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateBudgetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateBudgetResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateBudgetResponse struct {
	*tchttp.BaseResponse
	Response *AssociateBudgetResponseParams `json:"Response"`
}

func (r *AssociateBudgetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateBudgetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateCustomizedConfigRequestParams struct {
	// 配置ID
	UconfigId *string `json:"UconfigId,omitnil,omitempty" name:"UconfigId"`

	// 关联的server或location
	BindList []*BindItem `json:"BindList,omitnil,omitempty" name:"BindList"`
}

type AssociateCustomizedConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置ID
	UconfigId *string `json:"UconfigId,omitnil,omitempty" name:"UconfigId"`

	// 关联的server或location
	BindList []*BindItem `json:"BindList,omitnil,omitempty" name:"BindList"`
}

func (r *AssociateCustomizedConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateCustomizedConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UconfigId")
	delete(f, "BindList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateCustomizedConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateCustomizedConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateCustomizedConfigResponse struct {
	*tchttp.BaseResponse
	Response *AssociateCustomizedConfigResponseParams `json:"Response"`
}

func (r *AssociateCustomizedConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateCustomizedConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociateGuardrailConfig struct {
	// <p>Guardrail 防护类型。</p><p>枚举值：</p><ul><li>WAF：使用腾讯云 WAF LLM SDK 接入配置对模型路由请求进行安全防护。</li></ul><p>当前仅支持 WAF；不传时默认为 WAF。</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>关联的腾讯云 WAF 实例 ID。</p><p>当 Type 为 WAF 时必填。接口会校验该 WAF 实例存在且属于当前账号。</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>WAF LLM SDK 接入服务 ID。</p><p>该字段对应 WAF LLM SDK 接入配置中的服务标识，用于指定模型路由请求要绑定的 WAF 防护配置。当 Type 为 WAF 时必填。接口会校验该服务配置存在于指定的 WAF 实例下。</p>
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// <p>最大检测对话轮数。</p><p>当 Type 为 WAF 时选填；未传时默认取值为 5。若传入，取值必须为正整数。</p>
	InputCheckDepth *uint64 `json:"InputCheckDepth,omitnil,omitempty" name:"InputCheckDepth"`
}

// Predefined struct for user
type AssociateModelRouterGuardrailsRequestParams struct {
	// <p>待关联的 Guardrail 防护配置列表。</p><p>当前最多支持 1 个元素。每个元素必须填写 InstanceId、ServiceId；Type 和 InputCheckDepth 为选填，不传时分别使用默认值 WAF 和 5。本结构不包含 GuardrailId，关联成功后由系统生成。</p>
	Guardrails []*AssociateGuardrailConfig `json:"Guardrails,omitnil,omitempty" name:"Guardrails"`

	// <p>模型路由实例 ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`
}

type AssociateModelRouterGuardrailsRequest struct {
	*tchttp.BaseRequest
	
	// <p>待关联的 Guardrail 防护配置列表。</p><p>当前最多支持 1 个元素。每个元素必须填写 InstanceId、ServiceId；Type 和 InputCheckDepth 为选填，不传时分别使用默认值 WAF 和 5。本结构不包含 GuardrailId，关联成功后由系统生成。</p>
	Guardrails []*AssociateGuardrailConfig `json:"Guardrails,omitnil,omitempty" name:"Guardrails"`

	// <p>模型路由实例 ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`
}

func (r *AssociateModelRouterGuardrailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateModelRouterGuardrailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Guardrails")
	delete(f, "ModelRouterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateModelRouterGuardrailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateModelRouterGuardrailsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateModelRouterGuardrailsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateModelRouterGuardrailsResponseParams `json:"Response"`
}

func (r *AssociateModelRouterGuardrailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateModelRouterGuardrailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateModelsToModelRouterRequestParams struct {
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>需要关联的模型信息</p>
	Models []*ModelRouterModel `json:"Models,omitnil,omitempty" name:"Models"`
}

type AssociateModelsToModelRouterRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>需要关联的模型信息</p>
	Models []*ModelRouterModel `json:"Models,omitnil,omitempty" name:"Models"`
}

func (r *AssociateModelsToModelRouterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateModelsToModelRouterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "Models")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateModelsToModelRouterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateModelsToModelRouterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateModelsToModelRouterResponse struct {
	*tchttp.BaseResponse
	Response *AssociateModelsToModelRouterResponseParams `json:"Response"`
}

func (r *AssociateModelsToModelRouterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateModelsToModelRouterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateTargetGroupsRequestParams struct {
	// 绑定的关系数组，目标组类型需要一致。
	// 一次请求最多支持20个。
	Associations []*TargetGroupAssociation `json:"Associations,omitnil,omitempty" name:"Associations"`
}

type AssociateTargetGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 绑定的关系数组，目标组类型需要一致。
	// 一次请求最多支持20个。
	Associations []*TargetGroupAssociation `json:"Associations,omitnil,omitempty" name:"Associations"`
}

func (r *AssociateTargetGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateTargetGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Associations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateTargetGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateTargetGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateTargetGroupsResponse struct {
	*tchttp.BaseResponse
	Response *AssociateTargetGroupsResponseParams `json:"Response"`
}

func (r *AssociateTargetGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateTargetGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AssociatedModelRouterItem struct {
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`
}

type AssociationItem struct {
	// 关联到的负载均衡ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 关联到的监听器ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 关联到的转发规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 关联到的监听器协议类型，如HTTP,TCP,
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 关联到的监听器端口
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 关联到的转发规则域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 关联到的转发规则URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 负载均衡名称
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 监听器名称
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// 关联目标组的权重， 该参数只有v2新版目标组生效。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 高级路由规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

// Predefined struct for user
type AutoRewriteRequestParams struct {
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// HTTPS:443监听器的ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// HTTPS:443监听器下需要重定向的域名，若不填，则对HTTPS:443监听器下的所有域名都设置重定向。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 重定向状态码，可取值301,302,307。
	RewriteCodes []*int64 `json:"RewriteCodes,omitnil,omitempty" name:"RewriteCodes"`

	// 重定向是否携带匹配的URL。
	TakeUrls []*bool `json:"TakeUrls,omitnil,omitempty" name:"TakeUrls"`
}

type AutoRewriteRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// HTTPS:443监听器的ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// HTTPS:443监听器下需要重定向的域名，若不填，则对HTTPS:443监听器下的所有域名都设置重定向。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 重定向状态码，可取值301,302,307。
	RewriteCodes []*int64 `json:"RewriteCodes,omitnil,omitempty" name:"RewriteCodes"`

	// 重定向是否携带匹配的URL。
	TakeUrls []*bool `json:"TakeUrls,omitnil,omitempty" name:"TakeUrls"`
}

func (r *AutoRewriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AutoRewriteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "Domains")
	delete(f, "RewriteCodes")
	delete(f, "TakeUrls")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AutoRewriteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AutoRewriteResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AutoRewriteResponse struct {
	*tchttp.BaseResponse
	Response *AutoRewriteResponseParams `json:"Response"`
}

func (r *AutoRewriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AutoRewriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AvailableZoneAffinityInfo struct {
	// 是否开启可用区转发亲和。true：开启可用区转发亲和；false：开启可用区转发亲和。
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 可用区转发亲和失效阈值，当可用区内后端服务健康比例小于该阈值时，负载均衡会退出可用区转发亲和，转为全可用区转发。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExitRatio *uint64 `json:"ExitRatio,omitnil,omitempty" name:"ExitRatio"`

	// 可用区转发亲和的重新生效阈值，当处于全可用区转发，且负载均衡可用区内的后端服务健康比例大于等于该阈值时，负载均衡会重新进入可用区转发亲和。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReentryRatio *uint64 `json:"ReentryRatio,omitnil,omitempty" name:"ReentryRatio"`
}

type Backend struct {
	// 后端服务的类型，可取：CVM、ENI、CCN、EVM、GLOBALROUTE、NAT、SRV等
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 后端服务的唯一 ID，如 ins-abcd1234
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 后端服务的监听端口，如果是全端口段监听器绑定的全监听目标组场景，此端口返回0，表示无效端口，绑定的后端服务的端口随监听器端口。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 后端服务的转发权重，取值范围：[0, 100]，默认为 10。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 后端服务的外网 IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// 后端服务的内网 IP
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// 后端服务的实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 后端服务被绑定的时间
	RegisteredTime *string `json:"RegisteredTime,omitnil,omitempty" name:"RegisteredTime"`

	// 弹性网卡唯一ID，如 eni-1234abcd
	// 注意：此字段可能返回 null，表示取不到有效值。
	EniId *string `json:"EniId,omitnil,omitempty" name:"EniId"`

	// 标签。
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 后端服务所在的可用区，如ap-guangzhou-1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type BasicTargetGroupInfo struct {
	// 目标组ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 目标组名称
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`

	// 目标组权重
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

// Predefined struct for user
type BatchDeregisterTargetsRequestParams struct {
	// 负载均衡ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 解绑目标。
	Targets []*BatchTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

type BatchDeregisterTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 解绑目标。
	Targets []*BatchTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

func (r *BatchDeregisterTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeregisterTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "Targets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeregisterTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeregisterTargetsResponseParams struct {
	// 解绑失败的监听器ID。
	FailListenerIdSet []*string `json:"FailListenerIdSet,omitnil,omitempty" name:"FailListenerIdSet"`

	// 解绑失败错误原因信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchDeregisterTargetsResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeregisterTargetsResponseParams `json:"Response"`
}

func (r *BatchDeregisterTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeregisterTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyTargetTagRequestParams struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 要批量修改标签的列表。
	ModifyList []*RsTagRule `json:"ModifyList,omitnil,omitempty" name:"ModifyList"`
}

type BatchModifyTargetTagRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 要批量修改标签的列表。
	ModifyList []*RsTagRule `json:"ModifyList,omitnil,omitempty" name:"ModifyList"`
}

func (r *BatchModifyTargetTagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyTargetTagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ModifyList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchModifyTargetTagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyTargetTagResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchModifyTargetTagResponse struct {
	*tchttp.BaseResponse
	Response *BatchModifyTargetTagResponseParams `json:"Response"`
}

func (r *BatchModifyTargetTagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyTargetTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyTargetWeightRequestParams struct {
	// <p>负载均衡实例 ID。</p>
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// <p>要批量修改权重的列表。ModifyList数组最多100个元素，ModifyList[i].Targets最多50个，全部Targets累加不超过500。</p>
	ModifyList []*RsWeightRule `json:"ModifyList,omitnil,omitempty" name:"ModifyList"`
}

type BatchModifyTargetWeightRequest struct {
	*tchttp.BaseRequest
	
	// <p>负载均衡实例 ID。</p>
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// <p>要批量修改权重的列表。ModifyList数组最多100个元素，ModifyList[i].Targets最多50个，全部Targets累加不超过500。</p>
	ModifyList []*RsWeightRule `json:"ModifyList,omitnil,omitempty" name:"ModifyList"`
}

func (r *BatchModifyTargetWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyTargetWeightRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ModifyList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchModifyTargetWeightRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyTargetWeightResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchModifyTargetWeightResponse struct {
	*tchttp.BaseResponse
	Response *BatchModifyTargetWeightResponseParams `json:"Response"`
}

func (r *BatchModifyTargetWeightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyTargetWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchRegisterTargetsRequestParams struct {
	// 负载均衡ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 绑定目标。
	Targets []*BatchTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

type BatchRegisterTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 绑定目标。
	Targets []*BatchTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

func (r *BatchRegisterTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchRegisterTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "Targets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchRegisterTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchRegisterTargetsResponseParams struct {
	// 绑定失败的监听器ID，如为空表示全部绑定成功。
	FailListenerIdSet []*string `json:"FailListenerIdSet,omitnil,omitempty" name:"FailListenerIdSet"`

	// 绑定失败错误原因信息。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchRegisterTargetsResponse struct {
	*tchttp.BaseResponse
	Response *BatchRegisterTargetsResponseParams `json:"Response"`
}

func (r *BatchRegisterTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchRegisterTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchTarget struct {
	// 监听器 ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 绑定端口。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 子机 ID。表示绑定主网卡主 IP。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 绑定 IP 时需要传入此参数，支持弹性网卡的 IP 和其他内网 IP，如果是弹性网卡则必须先绑定至CVM，然后才能绑定到负载均衡实例。
	// 注意：参数 InstanceId、EniIp 只能传入一个且必须传入一个。如果绑定双栈IPV6子机，必须传该参数。
	EniIp *string `json:"EniIp,omitnil,omitempty" name:"EniIp"`

	// 子机权重，范围[0, 100]。绑定时如果不存在，则默认为10。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 七层规则 ID。7层负载均衡该参数必填
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 标签。
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`
}

type BindDetailItem struct {
	// 配置绑定的CLB ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 配置绑定的监听器ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 配置绑定的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 配置绑定的规则
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 监听器名字
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// 监听器协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 监听器端口
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// location的url
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 配置ID
	UconfigId *string `json:"UconfigId,omitnil,omitempty" name:"UconfigId"`
}

type BindItem struct {
	// 配置绑定的CLB ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 配置绑定的监听器ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 配置绑定的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 配置绑定的规则
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`
}

type BlockedIP struct {
	// 黑名单IP
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// 加入黑名单的时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type BudgetAssociation struct {
	// <p>Budget ID。</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>关联创建时间。</p>
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// <p>Key ID。仅当Type为Key时返回。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// <p>模型路由实例ID。</p><p>当Type为ModelRouter时表示关联资源本身；当Type为Key时表示Key所属实例。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>资源对象的名称。</p>
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// <p>关联关系的状态</p><p>枚举值：</p><ul><li>Active： 已生效</li><li>Configuring： 配置中</li><li>ConfigureFailed： 配置失败</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>关联资源类型。</p><p>枚举值：</p><ul><li>ModelRouter：模型路由实例</li><li>Key：模型路由Key</li><li>UserGroup：用户组</li></ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>关联的用户组id</p>
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`
}

type BudgetConfig struct {
	// <p>预算刷新周期。</p><p>枚举值：</p><ul><li>1d：按天刷新</li><li>7d：按周刷新</li><li>30d：按月刷新</li></ul><p>不传时默认30d。同一个Budget下每种刷新周期最多配置一次。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BudgetDuration *string `json:"BudgetDuration,omitnil,omitempty" name:"BudgetDuration"`

	// <p>下一次刷新的时间</p>
	BudgetResetAt *string `json:"BudgetResetAt,omitnil,omitempty" name:"BudgetResetAt"`

	// <p>最大预算。</p><p>单位：credit。取值需大于0且不超过10000000000；不传时默认100000。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxBudget *float64 `json:"MaxBudget,omitnil,omitempty" name:"MaxBudget"`
}

type BudgetConfigInput struct {
	// <p>预算刷新周期。</p><p>支持取值：</p><ul><li>1d：按天刷新</li><li>7d：按周刷新</li><li>30d：按月刷新</li></ul><p>不传时默认使用30d。同一个Budget下每种刷新周期最多配置一次。</p>
	BudgetDuration *string `json:"BudgetDuration,omitnil,omitempty" name:"BudgetDuration"`

	// <p>最大预算。</p><p>单位：credit。取值需大于0且不超过10000000000；不传时默认100000。</p>
	MaxBudget *float64 `json:"MaxBudget,omitnil,omitempty" name:"MaxBudget"`
}

type BudgetInfo struct {
	// <p>关联的key数量</p>
	AssociationKeyCount *uint64 `json:"AssociationKeyCount,omitnil,omitempty" name:"AssociationKeyCount"`

	// <p>关联的模型路由数量</p>
	AssociationModelRouterCount *uint64 `json:"AssociationModelRouterCount,omitnil,omitempty" name:"AssociationModelRouterCount"`

	// <p>关联的用户组数量</p>
	AssociationUserGroupCount *uint64 `json:"AssociationUserGroupCount,omitnil,omitempty" name:"AssociationUserGroupCount"`

	// <p>Budget预算配置数组。</p><p>最多返回3个元素，每种刷新周期（1d/7d/30d）各一个。</p>
	BudgetConfigs []*BudgetConfig `json:"BudgetConfigs,omitnil,omitempty" name:"BudgetConfigs"`

	// <p>Budget ID。</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>Budget名称。</p>
	BudgetName *string `json:"BudgetName,omitnil,omitempty" name:"BudgetName"`

	// <p>创建时间。</p>
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// <p>修改时间。</p>
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// <p>Budget限速信息。</p>
	RateLimitConfig *RateLimitConfigForBudget `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`

	// <p>Budget状态。</p><p>枚举值：</p><ul><li>Provisioning：创建中</li><li>Active：运行中</li><li>Configuring：变配中</li><li>Deleting：删除中</li><li>ProvisionFailed：创建失败</li><li>ConfigureFailed：变配失败</li><li>DeletionFailed：删除失败</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type BudgetResource struct {
	// <p>模型路由实例ID。</p><p>当Type为ModelRouter时表示要关联的实例；当Type为Key时表示Key所属实例。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>资源类型。</p><p>枚举值：</p><ul><li>ModelRouter：模型路由实例</li><li>Key：模型路由Key</li><li>UserGroup：用户组（Type 为 UserGroup 时需传 UserGroupId）</li></ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Key ID。</p><p>字段本身选填；当Type为Key时必填，当Type为ModelRouter时不传。</p>
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// <p>用户组ID</p>
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`
}

type CertIdRelatedWithLoadBalancers struct {
	// 证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 与证书关联的负载均衡实例列表
	LoadBalancers []*LoadBalancer `json:"LoadBalancers,omitnil,omitempty" name:"LoadBalancers"`
}

type CertInfo struct {
	// 证书 ID，如果不填写此项则必须上传证书内容，包括CertName, CertContent，若为服务端证书必须包含CertKey。
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 上传证书的名称，如果没有 CertId，则此项必传。
	CertName *string `json:"CertName,omitnil,omitempty" name:"CertName"`

	// 上传证书的公钥；如果没有 CertId，则此项必传。
	CertContent *string `json:"CertContent,omitnil,omitempty" name:"CertContent"`

	// 上传服务端证书的私钥；如果没有 CertId，则此项必传。
	CertKey *string `json:"CertKey,omitnil,omitempty" name:"CertKey"`
}

type CertificateInput struct {
	// 认证类型，UNIDIRECTIONAL：单向认证，MUTUAL：双向认证。
	// 默认为 UNIDIRECTIONAL。
	SSLMode *string `json:"SSLMode,omitnil,omitempty" name:"SSLMode"`

	// 双向认证时，是否开启客户端认证，ON:开启，OPTIONAL:自适应，默认ON。
	SSLVerifyClient *string `json:"SSLVerifyClient,omitnil,omitempty" name:"SSLVerifyClient"`

	// 服务端证书的 ID，如果不填写此项则必须上传证书，包括 CertContent（服务端证书内容），CertKey（服务端证书密钥），CertName（服务端证书名称）。
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 客户端证书的 ID，当监听器采用双向认证，即 SSLMode=MUTUAL 时，如果不填写此项则必须上传客户端证书，包括 CertCaContent，CertCaName。
	CertCaId *string `json:"CertCaId,omitnil,omitempty" name:"CertCaId"`

	// 上传服务端证书的名称，如果没有 CertId，则此项必传。
	CertName *string `json:"CertName,omitnil,omitempty" name:"CertName"`

	// 上传服务端证书的 key，如果没有 CertId，则此项必传。
	CertKey *string `json:"CertKey,omitnil,omitempty" name:"CertKey"`

	// 上传服务端证书的内容，如果没有 CertId，则此项必传。
	CertContent *string `json:"CertContent,omitnil,omitempty" name:"CertContent"`

	// 上传客户端 CA 证书的名称，如果 SSLMode=MUTUAL，如果没有 CertCaId，则此项必传。
	CertCaName *string `json:"CertCaName,omitnil,omitempty" name:"CertCaName"`

	// 上传客户端证书的内容，如果 SSLMode=MUTUAL，如果没有 CertCaId，则此项必传。
	CertCaContent *string `json:"CertCaContent,omitnil,omitempty" name:"CertCaContent"`
}

type CertificateOutput struct {
	// 认证类型，UNIDIRECTIONAL：单向认证，MUTUAL：双向认证
	SSLMode *string `json:"SSLMode,omitnil,omitempty" name:"SSLMode"`

	// 是否开启客户端证书验证，只在双向认证时生效。
	SSLVerifyClient *string `json:"SSLVerifyClient,omitnil,omitempty" name:"SSLVerifyClient"`

	// 服务端证书的ID。
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 客户端证书的 ID。
	CertCaId *string `json:"CertCaId,omitnil,omitempty" name:"CertCaId"`

	// 多本服务器证书场景扩展的服务器证书ID。
	ExtCertIds []*string `json:"ExtCertIds,omitnil,omitempty" name:"ExtCertIds"`
}

// Predefined struct for user
type ChatCompletionsRequestParams struct {
	// <p>virtual key，用于向代理网关鉴权</p>
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`

	// <p>多模态附件列表</p>
	Attachments []*MultiModalityAttachments `json:"Attachments,omitnil,omitempty" name:"Attachments"`

	// <p>聊天内容</p>
	ChatContent *string `json:"ChatContent,omitnil,omitempty" name:"ChatContent"`

	// <p>模型名称，配置的模型标识</p><p>示例：gpt-4o、deepseek-chat</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>请求路径</p><p>默认值：/v1/chat/completions</p>
	RequestPath *string `json:"RequestPath,omitnil,omitempty" name:"RequestPath"`
}

type ChatCompletionsRequest struct {
	*tchttp.BaseRequest
	
	// <p>virtual key，用于向代理网关鉴权</p>
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`

	// <p>多模态附件列表</p>
	Attachments []*MultiModalityAttachments `json:"Attachments,omitnil,omitempty" name:"Attachments"`

	// <p>聊天内容</p>
	ChatContent *string `json:"ChatContent,omitnil,omitempty" name:"ChatContent"`

	// <p>模型名称，配置的模型标识</p><p>示例：gpt-4o、deepseek-chat</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>请求路径</p><p>默认值：/v1/chat/completions</p>
	RequestPath *string `json:"RequestPath,omitnil,omitempty" name:"RequestPath"`
}

func (r *ChatCompletionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatCompletionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApiKey")
	delete(f, "Attachments")
	delete(f, "ChatContent")
	delete(f, "Model")
	delete(f, "ModelRouterId")
	delete(f, "RequestPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChatCompletionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChatCompletionsResponseParams struct {
	// <p>聊天的返回信息</p>
	ChatResponseMessage *string `json:"ChatResponseMessage,omitnil,omitempty" name:"ChatResponseMessage"`

	// <p>聊天请求发送过程中的失败信息</p>
	ErrorInChat *string `json:"ErrorInChat,omitnil,omitempty" name:"ErrorInChat"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ChatCompletionsResponse struct {
	*tchttp.BaseResponse
	Response *ChatCompletionsResponseParams `json:"Response"`
}

func (r *ChatCompletionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatCompletionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClassicalHealth struct {
	// 后端服务的内网 IP
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// 后端服务的端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 负载均衡的监听端口
	ListenerPort *int64 `json:"ListenerPort,omitnil,omitempty" name:"ListenerPort"`

	// 转发协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 健康检查结果，1 表示健康，0 表示不健康
	HealthStatus *int64 `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`
}

type ClassicalListener struct {
	// <p>负载均衡监听器ID</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>负载均衡监听器端口</p>
	ListenerPort *int64 `json:"ListenerPort,omitnil,omitempty" name:"ListenerPort"`

	// <p>监听器后端转发端口</p>
	InstancePort *int64 `json:"InstancePort,omitnil,omitempty" name:"InstancePort"`

	// <p>监听器名称</p>
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// <p>监听器协议类型</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>会话保持时间</p>
	SessionExpire *int64 `json:"SessionExpire,omitnil,omitempty" name:"SessionExpire"`

	// <p>是否开启了健康检查：1（开启）、0（关闭）</p>
	HealthSwitch *int64 `json:"HealthSwitch,omitnil,omitempty" name:"HealthSwitch"`

	// <p>响应超时时间</p><p>单位：秒</p>
	TimeOut *int64 `json:"TimeOut,omitnil,omitempty" name:"TimeOut"`

	// <p>检查间隔</p><p>单位：秒</p>
	IntervalTime *int64 `json:"IntervalTime,omitnil,omitempty" name:"IntervalTime"`

	// <p>健康阈值</p>
	HealthNum *int64 `json:"HealthNum,omitnil,omitempty" name:"HealthNum"`

	// <p>不健康阈值</p>
	UnhealthNum *int64 `json:"UnhealthNum,omitnil,omitempty" name:"UnhealthNum"`

	// <p>传统型公网负载均衡 监听器的请求均衡方法。空字符串或wrr 表示按权重轮询，ip_hash 表示根据访问的源 IP 进行一致性哈希方式来分发，least_conn表示按最小连接数。</p>
	HttpHash *string `json:"HttpHash,omitnil,omitempty" name:"HttpHash"`

	// <p>传统型公网负载均衡的 HTTP、HTTPS 监听器的健康检查返回码。具体可参考创建监听器中对该字段的解释</p>
	HttpCode *int64 `json:"HttpCode,omitnil,omitempty" name:"HttpCode"`

	// <p>传统型公网负载均衡的 HTTP、HTTPS 监听器的健康检查路径</p>
	HttpCheckPath *string `json:"HttpCheckPath,omitnil,omitempty" name:"HttpCheckPath"`

	// <p>传统型公网负载均衡的 HTTPS 监听器的认证方式</p>
	SSLMode *string `json:"SSLMode,omitnil,omitempty" name:"SSLMode"`

	// <p>传统型公网负载均衡的 HTTPS 监听器的服务端证书 ID</p>
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// <p>传统型公网负载均衡的 HTTPS 监听器的客户端证书 ID</p>
	CertCaId *string `json:"CertCaId,omitnil,omitempty" name:"CertCaId"`

	// <p>监听器的状态，0 表示创建中，1 表示运行中</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ClassicalLoadBalancerInfo struct {
	// 后端实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 负载均衡实例ID列表
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`
}

type ClassicalTarget struct {
	// 后端服务的类型，可取值：CVM、ENI（即将支持）
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 后端服务的唯一 ID，可通过 DescribeInstances 接口返回字段中的 unInstanceId 字段获取
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 后端服务的转发权重，取值范围：[0, 100]，默认为 10。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 后端服务的外网 IP
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// 后端服务的内网 IP
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// 后端服务的实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 后端服务的状态
	// 1：故障，2：运行中，3：创建中，4：已关机，5：已退还，6：退还中， 7：重启中，8：开机中，9：关机中，10：密码重置中，11：格式化中，12：镜像制作中，13：带宽设置中，14：重装系统中，19：升级中，21：热迁移中
	RunFlag *int64 `json:"RunFlag,omitnil,omitempty" name:"RunFlag"`
}

type ClassicalTargetInfo struct {
	// 后端实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 权重，取值范围 [0, 100]
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

// Predefined struct for user
type CloneLoadBalancerRequestParams struct {
	// 负载均衡ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口获取。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 克隆出负载均衡实例的名称，规则：1-60 个英文、汉字、数字、连接线“-”或下划线“_”。
	// 注意：如果名称与系统中已有负载均衡实例的名称相同，则系统将会自动生成此次创建的负载均衡实例的名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 负载均衡实例所属的项目 ID，默认项目 ID 为0，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口获取。不传此参数则视为默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 仅适用于公网负载均衡。设置跨可用区容灾时的主可用区ID，可用区 ID 和名称均支持，例如 100001 或 ap-guangzhou-1。
	// 注：主可用区是需要承载流量的可用区，备可用区默认不承载流量，主可用区不可用时才使用备可用区，平台将为您自动选择最佳备可用区。可通过 [DescribeResources](https://cloud.tencent.com/document/api/214/70213) 接口查询一个地域的主可用区的列表。
	MasterZoneId *string `json:"MasterZoneId,omitnil,omitempty" name:"MasterZoneId"`

	// 仅适用于公网负载均衡。设置跨可用区容灾时的备可用区ID，可用区 ID 和名称均支持，例如 100001 或 ap-guangzhou-1。
	// 注：备可用区是主可用区故障后，需要承载流量的可用区。可通过 [DescribeResources](https://cloud.tencent.com/document/api/214/70213) 接口查询一个地域的主/备可用区的列表。
	SlaveZoneId *string `json:"SlaveZoneId,omitnil,omitempty" name:"SlaveZoneId"`

	// 仅适用于公网负载均衡。可用区ID，可用区 ID 和名称均支持，指定可用区以创建负载均衡实例。如：100001 或 ap-guangzhou-1。不传则查询所有可用区的 CVM 实例。如需指定可用区，可调用查询可用区列表 [DescribeZones](https://cloud.tencent.com/document/product/213/15707) 接口查询。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 仅适用于公网负载均衡。负载均衡的网络计费模式。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 仅适用于公网负载均衡。目前仅广州、上海、南京、济南、杭州、福州、北京、石家庄、武汉、长沙、成都、重庆地域支持静态单线 IP 线路类型，如需体验，请联系商务经理申请。申请通过后，即可选择中国移动（CMCC）、中国联通（CUCC）或中国电信（CTCC）的运营商类型，网络计费模式只能使用按带宽包计费(BANDWIDTH_PACKAGE)。 如果不指定本参数，则默认使用BGP。可通过 DescribeResources 接口查询一个地域所支持的Isp。
	VipIsp *string `json:"VipIsp,omitnil,omitempty" name:"VipIsp"`

	// 指定Vip申请负载均衡。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 购买负载均衡同时，给负载均衡打上标签。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 独占集群信息。
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitnil,omitempty" name:"ExclusiveCluster"`

	// 带宽包ID，可以通过 [DescribeBandwidthPackages](https://cloud.tencent.com/document/api/215/19209) 接口获取。指定此参数时，网络计费方式（InternetAccessible.InternetChargeType）只支持按带宽包计费（BANDWIDTH_PACKAGE）。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// 是否支持绑定跨地域/跨Vpc绑定IP的功能。
	SnatPro *bool `json:"SnatPro,omitnil,omitempty" name:"SnatPro"`

	// 开启绑定跨地域/跨Vpc绑定IP的功能后，创建SnatIp。
	SnatIps []*SnatIp `json:"SnatIps,omitnil,omitempty" name:"SnatIps"`

	// 公网独占集群ID或者CDCId，可以通过 [DescribeExclusiveClusters](https://cloud.tencent.com/document/product/214/49278) 接口获取。
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 性能容量型规格。<li>clb.c2.medium（标准型）</li><li>clb.c3.small（高阶型1）</li><li>clb.c3.medium（高阶型2）</li><li>clb.c4.small（超强型1）</li><li>clb.c4.medium（超强型2）</li><li>clb.c4.large（超强型3）</li><li>clb.c4.xlarge（超强型4）</li>
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`

	// Stgw独占集群的标签。
	ClusterTag *string `json:"ClusterTag,omitnil,omitempty" name:"ClusterTag"`

	// 仅适用于私有网络内网负载均衡。内网就近接入时，选择可用区下发。可调用[DescribeZones](https://cloud.tencent.com/document/product/213/15707)接口查询可用区列表。
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// EIP 的唯一 ID，形如：eip-qhx8udkc，仅适用于内网负载均衡绑定EIP，可以通过 [DescribeAddresses](https://cloud.tencent.com/document/product/215/16702) 接口查询。
	EipAddressId *string `json:"EipAddressId,omitnil,omitempty" name:"EipAddressId"`
}

type CloneLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口获取。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 克隆出负载均衡实例的名称，规则：1-60 个英文、汉字、数字、连接线“-”或下划线“_”。
	// 注意：如果名称与系统中已有负载均衡实例的名称相同，则系统将会自动生成此次创建的负载均衡实例的名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 负载均衡实例所属的项目 ID，默认项目 ID 为0，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口获取。不传此参数则视为默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 仅适用于公网负载均衡。设置跨可用区容灾时的主可用区ID，可用区 ID 和名称均支持，例如 100001 或 ap-guangzhou-1。
	// 注：主可用区是需要承载流量的可用区，备可用区默认不承载流量，主可用区不可用时才使用备可用区，平台将为您自动选择最佳备可用区。可通过 [DescribeResources](https://cloud.tencent.com/document/api/214/70213) 接口查询一个地域的主可用区的列表。
	MasterZoneId *string `json:"MasterZoneId,omitnil,omitempty" name:"MasterZoneId"`

	// 仅适用于公网负载均衡。设置跨可用区容灾时的备可用区ID，可用区 ID 和名称均支持，例如 100001 或 ap-guangzhou-1。
	// 注：备可用区是主可用区故障后，需要承载流量的可用区。可通过 [DescribeResources](https://cloud.tencent.com/document/api/214/70213) 接口查询一个地域的主/备可用区的列表。
	SlaveZoneId *string `json:"SlaveZoneId,omitnil,omitempty" name:"SlaveZoneId"`

	// 仅适用于公网负载均衡。可用区ID，可用区 ID 和名称均支持，指定可用区以创建负载均衡实例。如：100001 或 ap-guangzhou-1。不传则查询所有可用区的 CVM 实例。如需指定可用区，可调用查询可用区列表 [DescribeZones](https://cloud.tencent.com/document/product/213/15707) 接口查询。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 仅适用于公网负载均衡。负载均衡的网络计费模式。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 仅适用于公网负载均衡。目前仅广州、上海、南京、济南、杭州、福州、北京、石家庄、武汉、长沙、成都、重庆地域支持静态单线 IP 线路类型，如需体验，请联系商务经理申请。申请通过后，即可选择中国移动（CMCC）、中国联通（CUCC）或中国电信（CTCC）的运营商类型，网络计费模式只能使用按带宽包计费(BANDWIDTH_PACKAGE)。 如果不指定本参数，则默认使用BGP。可通过 DescribeResources 接口查询一个地域所支持的Isp。
	VipIsp *string `json:"VipIsp,omitnil,omitempty" name:"VipIsp"`

	// 指定Vip申请负载均衡。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 购买负载均衡同时，给负载均衡打上标签。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 独占集群信息。
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitnil,omitempty" name:"ExclusiveCluster"`

	// 带宽包ID，可以通过 [DescribeBandwidthPackages](https://cloud.tencent.com/document/api/215/19209) 接口获取。指定此参数时，网络计费方式（InternetAccessible.InternetChargeType）只支持按带宽包计费（BANDWIDTH_PACKAGE）。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// 是否支持绑定跨地域/跨Vpc绑定IP的功能。
	SnatPro *bool `json:"SnatPro,omitnil,omitempty" name:"SnatPro"`

	// 开启绑定跨地域/跨Vpc绑定IP的功能后，创建SnatIp。
	SnatIps []*SnatIp `json:"SnatIps,omitnil,omitempty" name:"SnatIps"`

	// 公网独占集群ID或者CDCId，可以通过 [DescribeExclusiveClusters](https://cloud.tencent.com/document/product/214/49278) 接口获取。
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 性能容量型规格。<li>clb.c2.medium（标准型）</li><li>clb.c3.small（高阶型1）</li><li>clb.c3.medium（高阶型2）</li><li>clb.c4.small（超强型1）</li><li>clb.c4.medium（超强型2）</li><li>clb.c4.large（超强型3）</li><li>clb.c4.xlarge（超强型4）</li>
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`

	// Stgw独占集群的标签。
	ClusterTag *string `json:"ClusterTag,omitnil,omitempty" name:"ClusterTag"`

	// 仅适用于私有网络内网负载均衡。内网就近接入时，选择可用区下发。可调用[DescribeZones](https://cloud.tencent.com/document/product/213/15707)接口查询可用区列表。
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// EIP 的唯一 ID，形如：eip-qhx8udkc，仅适用于内网负载均衡绑定EIP，可以通过 [DescribeAddresses](https://cloud.tencent.com/document/product/215/16702) 接口查询。
	EipAddressId *string `json:"EipAddressId,omitnil,omitempty" name:"EipAddressId"`
}

func (r *CloneLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "LoadBalancerName")
	delete(f, "ProjectId")
	delete(f, "MasterZoneId")
	delete(f, "SlaveZoneId")
	delete(f, "ZoneId")
	delete(f, "InternetAccessible")
	delete(f, "VipIsp")
	delete(f, "Vip")
	delete(f, "Tags")
	delete(f, "ExclusiveCluster")
	delete(f, "BandwidthPackageId")
	delete(f, "SnatPro")
	delete(f, "SnatIps")
	delete(f, "ClusterIds")
	delete(f, "SlaType")
	delete(f, "ClusterTag")
	delete(f, "Zones")
	delete(f, "EipAddressId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloneLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloneLoadBalancerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloneLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *CloneLoadBalancerResponseParams `json:"Response"`
}

func (r *CloneLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloneLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Cluster struct {
	// 集群唯一ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 集群类型，如TGW，STGW，VPCGW
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 集群标签，只有TGW/STGW集群有标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterTag *string `json:"ClusterTag,omitnil,omitempty" name:"ClusterTag"`

	// 集群所在可用区，如ap-guangzhou-1
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 集群网络类型，如Public，Private
	Network *string `json:"Network,omitnil,omitempty" name:"Network"`

	// 最大连接数（个/秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxConn *int64 `json:"MaxConn,omitnil,omitempty" name:"MaxConn"`

	// 最大入带宽Mbps
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxInFlow *int64 `json:"MaxInFlow,omitnil,omitempty" name:"MaxInFlow"`

	// 最大入包量（个/秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxInPkg *int64 `json:"MaxInPkg,omitnil,omitempty" name:"MaxInPkg"`

	// 最大出带宽Mbps
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxOutFlow *int64 `json:"MaxOutFlow,omitnil,omitempty" name:"MaxOutFlow"`

	// 最大出包量（个/秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxOutPkg *int64 `json:"MaxOutPkg,omitnil,omitempty" name:"MaxOutPkg"`

	// 最大新建连接数（个/秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxNewConn *int64 `json:"MaxNewConn,omitnil,omitempty" name:"MaxNewConn"`

	// http最大新建连接数（个/秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	HTTPMaxNewConn *int64 `json:"HTTPMaxNewConn,omitnil,omitempty" name:"HTTPMaxNewConn"`

	// https最大新建连接数（个/秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	HTTPSMaxNewConn *int64 `json:"HTTPSMaxNewConn,omitnil,omitempty" name:"HTTPSMaxNewConn"`

	// http QPS
	// 注意：此字段可能返回 null，表示取不到有效值。
	HTTPQps *int64 `json:"HTTPQps,omitnil,omitempty" name:"HTTPQps"`

	// https QPS
	// 注意：此字段可能返回 null，表示取不到有效值。
	HTTPSQps *int64 `json:"HTTPSQps,omitnil,omitempty" name:"HTTPSQps"`

	// 集群内资源总数目
	ResourceCount *int64 `json:"ResourceCount,omitnil,omitempty" name:"ResourceCount"`

	// 集群内空闲资源数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdleResourceCount *int64 `json:"IdleResourceCount,omitnil,omitempty" name:"IdleResourceCount"`

	// 集群内转发机的数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalanceDirectorCount *int64 `json:"LoadBalanceDirectorCount,omitnil,omitempty" name:"LoadBalanceDirectorCount"`

	// 集群的Isp属性，如："BGP","CMCC","CUCC","CTCC","INTERNAL"。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 集群所在的可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClustersZone *ClustersZone `json:"ClustersZone,omitnil,omitempty" name:"ClustersZone"`

	// 集群版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClustersVersion *string `json:"ClustersVersion,omitnil,omitempty" name:"ClustersVersion"`

	// 集群容灾类型，如SINGLE-ZONE，DISASTER-RECOVERY，MUTUAL-DISASTER-RECOVERY
	DisasterRecoveryType *string `json:"DisasterRecoveryType,omitnil,omitempty" name:"DisasterRecoveryType"`

	// 网络出口
	Egress *string `json:"Egress,omitnil,omitempty" name:"Egress"`

	// IP版本
	IPVersion *string `json:"IPVersion,omitnil,omitempty" name:"IPVersion"`

	// 标签信息
	Tag []*TagInfo `json:"Tag,omitnil,omitempty" name:"Tag"`
}

type ClusterInfo struct {
	// <p>独占集群ID</p>
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// <p>集群类型</p><p>枚举值：</p><ul><li>Public： 公有云集群</li><li>Exclusive： 独占集群</li></ul><p>默认值：Public</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>集群名称</p>
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`
}

type ClusterItem struct {
	// 集群唯一ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 集群所在可用区，如ap-guangzhou-1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`
}

type ClusterResource struct {
	// 集群唯一ID，如tgw-12345678。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// ip地址。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 负载均衡唯一ID，如lb-12345678。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 资源是否闲置。
	Idle *string `json:"Idle,omitnil,omitempty" name:"Idle"`

	// 集群名称。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 集群的Isp属性，如："BGP","CMCC","CUCC","CTCC","INTERNAL"。
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 集群所在的可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClustersZone *ClustersZone `json:"ClustersZone,omitnil,omitempty" name:"ClustersZone"`
}

type ClustersZone struct {
	// 集群所在的主可用区。
	MasterZone []*string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// 集群所在的备可用区。
	SlaveZone []*string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`
}

type Coefficient struct {
	// <p>缓存命中输入积分系数。</p><p>用于 provider prompt cache 命中的输入 token。</p><p>取值范围：[0, 5000]</p><p>默认值：3</p>
	InputCachedCoefficient *float64 `json:"InputCachedCoefficient,omitnil,omitempty" name:"InputCachedCoefficient"`

	// <p>输入积分系数。</p><p>取值范围：[1, 5000]</p><p>默认值：25</p>
	InputCoefficient *float64 `json:"InputCoefficient,omitnil,omitempty" name:"InputCoefficient"`

	// <p>输出积分系数。</p><p>取值范围：[1, 5000]</p><p>默认值：100</p>
	OutputCoefficient *float64 `json:"OutputCoefficient,omitnil,omitempty" name:"OutputCoefficient"`
}

type ConfigListItem struct {
	// 配置ID
	UconfigId *string `json:"UconfigId,omitnil,omitempty" name:"UconfigId"`

	// 配置类型， 可选值：CLB（实例维度配置）， SERVER（服务维度配置），LOCATION（规则维度配置）
	ConfigType *string `json:"ConfigType,omitnil,omitempty" name:"ConfigType"`

	// 配置名字
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// 配置内容
	ConfigContent *string `json:"ConfigContent,omitnil,omitempty" name:"ConfigContent"`

	// 配置的创建时间。
	// 格式：YYYY-MM-DD HH:mm:ss
	CreateTimestamp *string `json:"CreateTimestamp,omitnil,omitempty" name:"CreateTimestamp"`

	// 配置的修改时间。
	// 格式：YYYY-MM-DD HH:mm:ss
	UpdateTimestamp *string `json:"UpdateTimestamp,omitnil,omitempty" name:"UpdateTimestamp"`
}

// Predefined struct for user
type CreateBYOKNetworkRequestParams struct {
	// <p>子网 ID</p><p>参数格式：subnet-xxxxxxxx</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>VPC 实例 ID</p><p>参数格式：vpc-xxxxxxxx</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>BYOK 的自定义名字</p><p>入参限制：1～256个字符，可选</p>
	ServiceProviderName *string `json:"ServiceProviderName,omitnil,omitempty" name:"ServiceProviderName"`

	// <p>标签</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateBYOKNetworkRequest struct {
	*tchttp.BaseRequest
	
	// <p>子网 ID</p><p>参数格式：subnet-xxxxxxxx</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>VPC 实例 ID</p><p>参数格式：vpc-xxxxxxxx</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>BYOK 的自定义名字</p><p>入参限制：1～256个字符，可选</p>
	ServiceProviderName *string `json:"ServiceProviderName,omitnil,omitempty" name:"ServiceProviderName"`

	// <p>标签</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateBYOKNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBYOKNetworkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubnetId")
	delete(f, "VpcId")
	delete(f, "ServiceProviderName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBYOKNetworkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBYOKNetworkResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBYOKNetworkResponse struct {
	*tchttp.BaseResponse
	Response *CreateBYOKNetworkResponseParams `json:"Response"`
}

func (r *CreateBYOKNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBYOKNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBudgetRequestParams struct {
	// <p>预算配置数组。</p><p>数组长度最大为3，最多可同时配置1d、7d、30d三个刷新周期，且每种刷新周期只能出现一次。BudgetResetAt不支持作为入参设置，系统会按配置的刷新周期自动维护刷新时间。</p>
	BudgetConfigs []*BudgetConfigInput `json:"BudgetConfigs,omitnil,omitempty" name:"BudgetConfigs"`

	// <p>Budget名称。</p><p>不传默认为 '-'。</p>
	BudgetName *string `json:"BudgetName,omitnil,omitempty" name:"BudgetName"`

	// <p>Budget限速配置。</p>
	RateLimitConfig *RateLimitConfigForBudget `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`

	// <p>创建Budget时同时关联的资源列表。</p><p>仅支持企业型模型路由实例和企业型实例下的Key。如果资源不存在或不可关联，创建请求失败；资源已关联其他Budget时将替换为新创建的Budget。</p>
	Resources []*BudgetResource `json:"Resources,omitnil,omitempty" name:"Resources"`
}

type CreateBudgetRequest struct {
	*tchttp.BaseRequest
	
	// <p>预算配置数组。</p><p>数组长度最大为3，最多可同时配置1d、7d、30d三个刷新周期，且每种刷新周期只能出现一次。BudgetResetAt不支持作为入参设置，系统会按配置的刷新周期自动维护刷新时间。</p>
	BudgetConfigs []*BudgetConfigInput `json:"BudgetConfigs,omitnil,omitempty" name:"BudgetConfigs"`

	// <p>Budget名称。</p><p>不传默认为 '-'。</p>
	BudgetName *string `json:"BudgetName,omitnil,omitempty" name:"BudgetName"`

	// <p>Budget限速配置。</p>
	RateLimitConfig *RateLimitConfigForBudget `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`

	// <p>创建Budget时同时关联的资源列表。</p><p>仅支持企业型模型路由实例和企业型实例下的Key。如果资源不存在或不可关联，创建请求失败；资源已关联其他Budget时将替换为新创建的Budget。</p>
	Resources []*BudgetResource `json:"Resources,omitnil,omitempty" name:"Resources"`
}

func (r *CreateBudgetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBudgetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BudgetConfigs")
	delete(f, "BudgetName")
	delete(f, "RateLimitConfig")
	delete(f, "Resources")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBudgetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBudgetResponseParams struct {
	// <p>Budget ID。</p><p>创建请求提交后返回，可通过DescribeBudgets查询状态。</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBudgetResponse struct {
	*tchttp.BaseResponse
	Response *CreateBudgetResponseParams `json:"Response"`
}

func (r *CreateBudgetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBudgetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClsLogSetRequestParams struct {
	// 日志集的名字，不能和cls其他日志集重名。不填默认为clb_logset。
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// 日志集的保存周期，单位：天。
	//
	// Deprecated: Period is deprecated.
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 日志集类型，ACCESS：访问日志，HEALTH：健康检查日志，默认ACCESS。
	LogsetType *string `json:"LogsetType,omitnil,omitempty" name:"LogsetType"`
}

type CreateClsLogSetRequest struct {
	*tchttp.BaseRequest
	
	// 日志集的名字，不能和cls其他日志集重名。不填默认为clb_logset。
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// 日志集的保存周期，单位：天。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 日志集类型，ACCESS：访问日志，HEALTH：健康检查日志，默认ACCESS。
	LogsetType *string `json:"LogsetType,omitnil,omitempty" name:"LogsetType"`
}

func (r *CreateClsLogSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClsLogSetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetName")
	delete(f, "Period")
	delete(f, "LogsetType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClsLogSetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClsLogSetResponseParams struct {
	// 日志集的 ID。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateClsLogSetResponse struct {
	*tchttp.BaseResponse
	Response *CreateClsLogSetResponseParams `json:"Response"`
}

func (r *CreateClsLogSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClsLogSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntentRouterRequestParams struct {
	// <p>模型路由实例ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>路由名称，用作LiteLLM的model_name。</p><p>必须以"IntentRouter/"为前缀，后缀仅支持字母、数字、连字符和下划线，后缀长度1-128个字符。</p>
	RouteName *string `json:"RouteName,omitnil,omitempty" name:"RouteName"`

	// <p>Tier配置列表。</p><p>每个Tier至少包含一个模型，模型名称必须是已关联到该实例的模型。</p>
	Tiers []*TierItem `json:"Tiers,omitnil,omitempty" name:"Tiers"`

	// <p>意图路由描述。</p>
	RouterDescribe *string `json:"RouterDescribe,omitnil,omitempty" name:"RouterDescribe"`
}

type CreateIntentRouterRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>路由名称，用作LiteLLM的model_name。</p><p>必须以"IntentRouter/"为前缀，后缀仅支持字母、数字、连字符和下划线，后缀长度1-128个字符。</p>
	RouteName *string `json:"RouteName,omitnil,omitempty" name:"RouteName"`

	// <p>Tier配置列表。</p><p>每个Tier至少包含一个模型，模型名称必须是已关联到该实例的模型。</p>
	Tiers []*TierItem `json:"Tiers,omitnil,omitempty" name:"Tiers"`

	// <p>意图路由描述。</p>
	RouterDescribe *string `json:"RouterDescribe,omitnil,omitempty" name:"RouterDescribe"`
}

func (r *CreateIntentRouterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntentRouterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "RouteName")
	delete(f, "Tiers")
	delete(f, "RouterDescribe")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIntentRouterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntentRouterResponseParams struct {
	// <p>创建的意图路由ID（ir-xxx格式）。</p>
	IntentRouterId *string `json:"IntentRouterId,omitnil,omitempty" name:"IntentRouterId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIntentRouterResponse struct {
	*tchttp.BaseResponse
	Response *CreateIntentRouterResponseParams `json:"Response"`
}

func (r *CreateIntentRouterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntentRouterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKeyRequestParams struct {
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>Key名称</p>
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// <p>限速配置</p>
	RateLimitConfig *RateLimitConfigForKey `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`

	// <p>关联的积分预算ID</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>需要关联的用户组ID</p>
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// <p>标签</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateKeyRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>Key名称</p>
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// <p>限速配置</p>
	RateLimitConfig *RateLimitConfigForKey `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`

	// <p>关联的积分预算ID</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>需要关联的用户组ID</p>
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// <p>标签</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "KeyName")
	delete(f, "RateLimitConfig")
	delete(f, "BudgetId")
	delete(f, "UserGroupId")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKeyResponseParams struct {
	// <p>Key的ID</p>
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// <p>返回的真实Key</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateKeyResponseParams `json:"Response"`
}

func (r *CreateKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKeysRequestParams struct {
	// <p>模型路由ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>需要绑定的预算信息，所有Key共用</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>Key列表</p>
	Keys []*InputKeyInfo `json:"Keys,omitnil,omitempty" name:"Keys"`

	// <p>批量创建Key的模式</p><p>枚举值：</p><ul><li>Generate： 平台生成Key</li><li>Import： 导入自带Key</li></ul><p>默认值：Generate</p>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>限速信息，所有Key共用</p>
	RateLimitConfig *RateLimitConfigForKey `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`

	// <p>标签。所有Key都会绑定该标签。</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>需要关联的用户组ID</p>
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`
}

type CreateKeysRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>需要绑定的预算信息，所有Key共用</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>Key列表</p>
	Keys []*InputKeyInfo `json:"Keys,omitnil,omitempty" name:"Keys"`

	// <p>批量创建Key的模式</p><p>枚举值：</p><ul><li>Generate： 平台生成Key</li><li>Import： 导入自带Key</li></ul><p>默认值：Generate</p>
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// <p>限速信息，所有Key共用</p>
	RateLimitConfig *RateLimitConfigForKey `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`

	// <p>标签。所有Key都会绑定该标签。</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>需要关联的用户组ID</p>
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`
}

func (r *CreateKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "BudgetId")
	delete(f, "Keys")
	delete(f, "Mode")
	delete(f, "RateLimitConfig")
	delete(f, "Tags")
	delete(f, "UserGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKeysResponseParams struct {
	// <p>创建的Key的信息</p>
	CreatedKeys []*CreatedKey `json:"CreatedKeys,omitnil,omitempty" name:"CreatedKeys"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateKeysResponse struct {
	*tchttp.BaseResponse
	Response *CreateKeysResponseParams `json:"Response"`
}

func (r *CreateKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateListenerRequestParams struct {
	// <p>负载均衡实例 ID，可以通过 <a href="https://cloud.tencent.com/document/product/214/30685">DescribeLoadBalancers</a> 接口获取。</p>
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// <p>要将监听器创建到哪些端口，每个端口对应一个新的监听器。<br>端口范围：1~65535</p>
	Ports []*int64 `json:"Ports,omitnil,omitempty" name:"Ports"`

	// <p>监听器协议： TCP | UDP | HTTP | HTTPS | TCP_SSL | QUIC。</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>要创建的监听器名称列表，名称与Ports数组按序一一对应，如不需立即命名，则无需提供此参数。命名规则：1-80 个英文字母、汉字等国际通用语言字符，数字，连接线“-”、下划线“_”等常见字符（禁止Unicode补充字符，如emoji表情、生僻汉字等）。</p>
	ListenerNames []*string `json:"ListenerNames,omitnil,omitempty" name:"ListenerNames"`

	// <p>健康检查相关参数，此参数仅适用于TCP/UDP/TCP_SSL/QUIC监听器。</p>
	HealthCheck *HealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// <p>证书相关信息。参数限制如下：</p><li>此参数仅适用于TCP_SSL监听器和未开启SNI特性的HTTPS监听器。</li><li>创建TCP_SSL监听器和未开启SNI特性的HTTPS监听器时，此参数和参数MultiCertInfo至少需要传一个， 但不能同时传入。</li>
	Certificate *CertificateInput `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// <p>会话保持时间，单位：秒。可选值：30~3600，默认为0，默认不开启。此参数仅适用于TCP/UDP监听器。</p>
	SessionExpireTime *int64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// <p>监听器转发的方式。可选值：WRR（按权重轮询）、LEAST_CONN（按最小连接数）<br>默认为 WRR。此参数仅适用于TCP/UDP/TCP_SSL/QUIC监听器。</p>
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// <p>是否开启SNI特性，此参数仅适用于HTTPS监听器。0表示未开启，1表示开启。</p>
	SniSwitch *int64 `json:"SniSwitch,omitnil,omitempty" name:"SniSwitch"`

	// <p>后端目标类型，NODE表示绑定普通节点，TARGETGROUP表示绑定目标组。此参数仅适用于TCP/UDP监听器。七层监听器应在转发规则中设置。</p>
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// <p>会话保持类型。不传或传NORMAL表示默认会话保持类型。QUIC_CID 表示根据Quic Connection ID做会话保持。QUIC_CID只支持UDP协议。此参数仅适用于TCP/UDP监听器。七层监听器应在转发规则中设置。（若选择QUIC_CID，则Protocol必须为UDP，Scheduler必须为WRR，同时只支持ipv4）</p>
	SessionType *string `json:"SessionType,omitnil,omitempty" name:"SessionType"`

	// <p>是否开启长连接，此参数仅适用于HTTP/HTTPS监听器，0:关闭；1:开启， 默认关闭。<br>若后端服务对连接数上限有限制，则建议谨慎开启。此功能目前处于内测中，如需使用，请提交 <a href="https://cloud.tencent.com/apply/p/tsodp6qm21">内测申请</a>。</p>
	KeepaliveEnable *int64 `json:"KeepaliveEnable,omitnil,omitempty" name:"KeepaliveEnable"`

	// <p>创建端口段监听器时必须传入此参数，用以标识结束端口。同时，入参Ports只允许传入一个成员，用以标识开始端口。【如果您需要体验端口段功能，请通过 <a href="https://console.cloud.tencent.com/workorder/category">工单申请</a>】。</p>
	EndPort *uint64 `json:"EndPort,omitnil,omitempty" name:"EndPort"`

	// <p>重新调度功能，解绑后端服务开关，打开此开关，当解绑后端服务时触发重新调度。仅TCP/UDP监听器支持。</p>
	DeregisterTargetRst *bool `json:"DeregisterTargetRst,omitnil,omitempty" name:"DeregisterTargetRst"`

	// <p>证书信息，支持同时传入不同算法类型的多本服务端证书，参数限制如下：</p><li>此参数仅适用于TCP_SSL监听器和未开启SNI特性的HTTPS监听器。</li><li>创建TCP_SSL监听器和未开启SNI特性的HTTPS监听器时，此参数和参数Certificate至少需要传一个， 但不能同时传入。</li>
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitnil,omitempty" name:"MultiCertInfo"`

	// <p>监听器最大连接数，当前仅性能容量型实例且仅TCP/UDP/TCP_SSL/QUIC监听器支持，不传或者传-1表示监听器维度不限速。基础网络实例不支持该参数。</p>
	MaxConn *int64 `json:"MaxConn,omitnil,omitempty" name:"MaxConn"`

	// <p>监听器最大新增连接数，当前仅性能容量型实例且仅TCP/UDP/TCP_SSL/QUIC监听器支持，不传或者传-1表示监听器维度不限速。基础网络实例不支持该参数。</p>
	MaxCps *int64 `json:"MaxCps,omitnil,omitempty" name:"MaxCps"`

	// <p>空闲连接超时时间，此参数仅适用于TCP/UDP监听器，单位：秒。默认值：TCP监听器默认值为900s，UDP监听器默认值为300s。取值范围：共享型实例和独占型实例支持：10-900，性能容量型实例支持：10-1980。如需设置超过取值范围的值请通过 <a href="https://console.cloud.tencent.com/workorder/category">工单申请</a>。</p><p>取值范围：[10, 1980]</p><p>单位：秒</p><p>默认值：900</p><p>TCP监听器默认值为900s，UDP监听器默认值为300s。取值范围：共享型实例和独占型实例支持：10-900，性能容量型实例支持：10-1980。</p>
	IdleConnectTimeout *int64 `json:"IdleConnectTimeout,omitnil,omitempty" name:"IdleConnectTimeout"`

	// <p>TCP_SSL和QUIC是否支持PP</p>
	ProxyProtocol *bool `json:"ProxyProtocol,omitnil,omitempty" name:"ProxyProtocol"`

	// <p>是否开启SNAT（源IP替换），True（开启）、False（关闭）。默认为关闭。注意：SnatEnable开启时会替换客户端源IP，此时<code>透传客户端源IP</code>选项关闭，反之亦然。</p>
	SnatEnable *bool `json:"SnatEnable,omitnil,omitempty" name:"SnatEnable"`

	// <p>全端口段监听器的结束端口，端口范围：2 - 65535</p>
	FullEndPorts []*int64 `json:"FullEndPorts,omitnil,omitempty" name:"FullEndPorts"`

	// <p>内网 HTTP 监听器开启 h2c 开关。<br>True（开启）、False（关闭）。默认为关闭。<br>开启后，监听器仅支持创建后端转发类型为 GRPC 或 GRPCS 的七层规则；创建规则时需在 Rules.N.ForwardType 中显式传入 GRPC 或 GRPCS。</p>
	H2cSwitch *bool `json:"H2cSwitch,omitnil,omitempty" name:"H2cSwitch"`

	// <p>控制 TCP_SSL 类型的监听器是否移除 SSL 加密层。开启后，监听器将作为普通 TCP 协议运行。 可选值：</p><ul><li>True： 关闭 SSL 功能（协议降级为纯文本 TCP）。</li><li>False（默认）： 保持 SSL 功能开启。</li></ul>
	SslCloseSwitch *bool `json:"SslCloseSwitch,omitnil,omitempty" name:"SslCloseSwitch"`

	// <p>数据压缩模式。可选值：transparent（透传模式）、compatibility（兼容模式）</p>
	DataCompressMode *string `json:"DataCompressMode,omitnil,omitempty" name:"DataCompressMode"`

	// <p>重新调度功能，权重调为0开关，打开此开关，后端服务器权重调为0时触发重新调度。仅TCP/UDP监听器支持。</p>
	RescheduleTargetZeroWeight *bool `json:"RescheduleTargetZeroWeight,omitnil,omitempty" name:"RescheduleTargetZeroWeight"`

	// <p>重新调度功能，健康检查异常开关，打开此开关，后端服务器健康检查异常时触发重新调度。仅TCP/UDP监听器支持。</p>
	RescheduleUnhealthy *bool `json:"RescheduleUnhealthy,omitnil,omitempty" name:"RescheduleUnhealthy"`

	// <p>重新调度功能，扩容后端服务开关，打开此开关，后端服务器增加或者减少时触发重新调度。仅TCP/UDP监听器支持。</p>
	RescheduleExpandTarget *bool `json:"RescheduleExpandTarget,omitnil,omitempty" name:"RescheduleExpandTarget"`

	// <p>重新调度触发开始时间，取值0~3600s。仅TCP/UDP监听器支持。</p>
	RescheduleStartTime *int64 `json:"RescheduleStartTime,omitnil,omitempty" name:"RescheduleStartTime"`

	// <p>重新调度触发持续时间，取值0~3600s。仅TCP/UDP监听器支持。</p>
	RescheduleInterval *int64 `json:"RescheduleInterval,omitnil,omitempty" name:"RescheduleInterval"`
}

type CreateListenerRequest struct {
	*tchttp.BaseRequest
	
	// <p>负载均衡实例 ID，可以通过 <a href="https://cloud.tencent.com/document/product/214/30685">DescribeLoadBalancers</a> 接口获取。</p>
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// <p>要将监听器创建到哪些端口，每个端口对应一个新的监听器。<br>端口范围：1~65535</p>
	Ports []*int64 `json:"Ports,omitnil,omitempty" name:"Ports"`

	// <p>监听器协议： TCP | UDP | HTTP | HTTPS | TCP_SSL | QUIC。</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>要创建的监听器名称列表，名称与Ports数组按序一一对应，如不需立即命名，则无需提供此参数。命名规则：1-80 个英文字母、汉字等国际通用语言字符，数字，连接线“-”、下划线“_”等常见字符（禁止Unicode补充字符，如emoji表情、生僻汉字等）。</p>
	ListenerNames []*string `json:"ListenerNames,omitnil,omitempty" name:"ListenerNames"`

	// <p>健康检查相关参数，此参数仅适用于TCP/UDP/TCP_SSL/QUIC监听器。</p>
	HealthCheck *HealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// <p>证书相关信息。参数限制如下：</p><li>此参数仅适用于TCP_SSL监听器和未开启SNI特性的HTTPS监听器。</li><li>创建TCP_SSL监听器和未开启SNI特性的HTTPS监听器时，此参数和参数MultiCertInfo至少需要传一个， 但不能同时传入。</li>
	Certificate *CertificateInput `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// <p>会话保持时间，单位：秒。可选值：30~3600，默认为0，默认不开启。此参数仅适用于TCP/UDP监听器。</p>
	SessionExpireTime *int64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// <p>监听器转发的方式。可选值：WRR（按权重轮询）、LEAST_CONN（按最小连接数）<br>默认为 WRR。此参数仅适用于TCP/UDP/TCP_SSL/QUIC监听器。</p>
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// <p>是否开启SNI特性，此参数仅适用于HTTPS监听器。0表示未开启，1表示开启。</p>
	SniSwitch *int64 `json:"SniSwitch,omitnil,omitempty" name:"SniSwitch"`

	// <p>后端目标类型，NODE表示绑定普通节点，TARGETGROUP表示绑定目标组。此参数仅适用于TCP/UDP监听器。七层监听器应在转发规则中设置。</p>
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// <p>会话保持类型。不传或传NORMAL表示默认会话保持类型。QUIC_CID 表示根据Quic Connection ID做会话保持。QUIC_CID只支持UDP协议。此参数仅适用于TCP/UDP监听器。七层监听器应在转发规则中设置。（若选择QUIC_CID，则Protocol必须为UDP，Scheduler必须为WRR，同时只支持ipv4）</p>
	SessionType *string `json:"SessionType,omitnil,omitempty" name:"SessionType"`

	// <p>是否开启长连接，此参数仅适用于HTTP/HTTPS监听器，0:关闭；1:开启， 默认关闭。<br>若后端服务对连接数上限有限制，则建议谨慎开启。此功能目前处于内测中，如需使用，请提交 <a href="https://cloud.tencent.com/apply/p/tsodp6qm21">内测申请</a>。</p>
	KeepaliveEnable *int64 `json:"KeepaliveEnable,omitnil,omitempty" name:"KeepaliveEnable"`

	// <p>创建端口段监听器时必须传入此参数，用以标识结束端口。同时，入参Ports只允许传入一个成员，用以标识开始端口。【如果您需要体验端口段功能，请通过 <a href="https://console.cloud.tencent.com/workorder/category">工单申请</a>】。</p>
	EndPort *uint64 `json:"EndPort,omitnil,omitempty" name:"EndPort"`

	// <p>重新调度功能，解绑后端服务开关，打开此开关，当解绑后端服务时触发重新调度。仅TCP/UDP监听器支持。</p>
	DeregisterTargetRst *bool `json:"DeregisterTargetRst,omitnil,omitempty" name:"DeregisterTargetRst"`

	// <p>证书信息，支持同时传入不同算法类型的多本服务端证书，参数限制如下：</p><li>此参数仅适用于TCP_SSL监听器和未开启SNI特性的HTTPS监听器。</li><li>创建TCP_SSL监听器和未开启SNI特性的HTTPS监听器时，此参数和参数Certificate至少需要传一个， 但不能同时传入。</li>
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitnil,omitempty" name:"MultiCertInfo"`

	// <p>监听器最大连接数，当前仅性能容量型实例且仅TCP/UDP/TCP_SSL/QUIC监听器支持，不传或者传-1表示监听器维度不限速。基础网络实例不支持该参数。</p>
	MaxConn *int64 `json:"MaxConn,omitnil,omitempty" name:"MaxConn"`

	// <p>监听器最大新增连接数，当前仅性能容量型实例且仅TCP/UDP/TCP_SSL/QUIC监听器支持，不传或者传-1表示监听器维度不限速。基础网络实例不支持该参数。</p>
	MaxCps *int64 `json:"MaxCps,omitnil,omitempty" name:"MaxCps"`

	// <p>空闲连接超时时间，此参数仅适用于TCP/UDP监听器，单位：秒。默认值：TCP监听器默认值为900s，UDP监听器默认值为300s。取值范围：共享型实例和独占型实例支持：10-900，性能容量型实例支持：10-1980。如需设置超过取值范围的值请通过 <a href="https://console.cloud.tencent.com/workorder/category">工单申请</a>。</p><p>取值范围：[10, 1980]</p><p>单位：秒</p><p>默认值：900</p><p>TCP监听器默认值为900s，UDP监听器默认值为300s。取值范围：共享型实例和独占型实例支持：10-900，性能容量型实例支持：10-1980。</p>
	IdleConnectTimeout *int64 `json:"IdleConnectTimeout,omitnil,omitempty" name:"IdleConnectTimeout"`

	// <p>TCP_SSL和QUIC是否支持PP</p>
	ProxyProtocol *bool `json:"ProxyProtocol,omitnil,omitempty" name:"ProxyProtocol"`

	// <p>是否开启SNAT（源IP替换），True（开启）、False（关闭）。默认为关闭。注意：SnatEnable开启时会替换客户端源IP，此时<code>透传客户端源IP</code>选项关闭，反之亦然。</p>
	SnatEnable *bool `json:"SnatEnable,omitnil,omitempty" name:"SnatEnable"`

	// <p>全端口段监听器的结束端口，端口范围：2 - 65535</p>
	FullEndPorts []*int64 `json:"FullEndPorts,omitnil,omitempty" name:"FullEndPorts"`

	// <p>内网 HTTP 监听器开启 h2c 开关。<br>True（开启）、False（关闭）。默认为关闭。<br>开启后，监听器仅支持创建后端转发类型为 GRPC 或 GRPCS 的七层规则；创建规则时需在 Rules.N.ForwardType 中显式传入 GRPC 或 GRPCS。</p>
	H2cSwitch *bool `json:"H2cSwitch,omitnil,omitempty" name:"H2cSwitch"`

	// <p>控制 TCP_SSL 类型的监听器是否移除 SSL 加密层。开启后，监听器将作为普通 TCP 协议运行。 可选值：</p><ul><li>True： 关闭 SSL 功能（协议降级为纯文本 TCP）。</li><li>False（默认）： 保持 SSL 功能开启。</li></ul>
	SslCloseSwitch *bool `json:"SslCloseSwitch,omitnil,omitempty" name:"SslCloseSwitch"`

	// <p>数据压缩模式。可选值：transparent（透传模式）、compatibility（兼容模式）</p>
	DataCompressMode *string `json:"DataCompressMode,omitnil,omitempty" name:"DataCompressMode"`

	// <p>重新调度功能，权重调为0开关，打开此开关，后端服务器权重调为0时触发重新调度。仅TCP/UDP监听器支持。</p>
	RescheduleTargetZeroWeight *bool `json:"RescheduleTargetZeroWeight,omitnil,omitempty" name:"RescheduleTargetZeroWeight"`

	// <p>重新调度功能，健康检查异常开关，打开此开关，后端服务器健康检查异常时触发重新调度。仅TCP/UDP监听器支持。</p>
	RescheduleUnhealthy *bool `json:"RescheduleUnhealthy,omitnil,omitempty" name:"RescheduleUnhealthy"`

	// <p>重新调度功能，扩容后端服务开关，打开此开关，后端服务器增加或者减少时触发重新调度。仅TCP/UDP监听器支持。</p>
	RescheduleExpandTarget *bool `json:"RescheduleExpandTarget,omitnil,omitempty" name:"RescheduleExpandTarget"`

	// <p>重新调度触发开始时间，取值0~3600s。仅TCP/UDP监听器支持。</p>
	RescheduleStartTime *int64 `json:"RescheduleStartTime,omitnil,omitempty" name:"RescheduleStartTime"`

	// <p>重新调度触发持续时间，取值0~3600s。仅TCP/UDP监听器支持。</p>
	RescheduleInterval *int64 `json:"RescheduleInterval,omitnil,omitempty" name:"RescheduleInterval"`
}

func (r *CreateListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "Ports")
	delete(f, "Protocol")
	delete(f, "ListenerNames")
	delete(f, "HealthCheck")
	delete(f, "Certificate")
	delete(f, "SessionExpireTime")
	delete(f, "Scheduler")
	delete(f, "SniSwitch")
	delete(f, "TargetType")
	delete(f, "SessionType")
	delete(f, "KeepaliveEnable")
	delete(f, "EndPort")
	delete(f, "DeregisterTargetRst")
	delete(f, "MultiCertInfo")
	delete(f, "MaxConn")
	delete(f, "MaxCps")
	delete(f, "IdleConnectTimeout")
	delete(f, "ProxyProtocol")
	delete(f, "SnatEnable")
	delete(f, "FullEndPorts")
	delete(f, "H2cSwitch")
	delete(f, "SslCloseSwitch")
	delete(f, "DataCompressMode")
	delete(f, "RescheduleTargetZeroWeight")
	delete(f, "RescheduleUnhealthy")
	delete(f, "RescheduleExpandTarget")
	delete(f, "RescheduleStartTime")
	delete(f, "RescheduleInterval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateListenerResponseParams struct {
	// <p>创建的监听器的唯一标识数组。</p>
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateListenerResponse struct {
	*tchttp.BaseResponse
	Response *CreateListenerResponseParams `json:"Response"`
}

func (r *CreateListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadBalancerRequestParams struct {
	// <p>负载均衡实例的网络类型：<br>OPEN：公网属性， INTERNAL：内网属性。</p>
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// <p>负载均衡实例的类型。1：通用的负载均衡实例，目前只支持传入1。</p>
	Forward *int64 `json:"Forward,omitnil,omitempty" name:"Forward"`

	// <p>负载均衡实例的名称，只在创建一个实例的时候才会生效。规则：1-80 个英文字母、汉字等国际通用语言字符，数字，连接线“-”、下划线“_”等常见字符（禁止Unicode补充字符，如emoji表情、生僻汉字等）。注意：如果名称与系统中已有负载均衡实例的名称相同，则系统将会自动生成此次创建的负载均衡实例的名称。</p>
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// <p>负载均衡后端目标设备所属的网络 ID，如vpc-12345678，可以通过 <a href="https://cloud.tencent.com/document/product/215/15778">DescribeVpcs</a> 接口获取。 不填此参数则默认为DefaultVPC。创建内网负载均衡实例时，此参数必填。</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>在私有网络内购买内网负载均衡实例的情况下，必须指定子网 ID，内网负载均衡实例的 VIP 将从这个子网中产生。<br>创建内网负载均衡实例，或者创建 IPv6FullChain 版本的负载均衡实例，此参数必填。<br>创建公网IPv4负载均衡实例时，不支持指定该参数。</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>负载均衡实例所属的项目 ID，默认项目 ID 为0。可以通过 <a href="https://cloud.tencent.com/document/api/651/78725">DescribeProject</a> 接口获取。不填此参数则视为默认项目。</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>仅适用于公网负载均衡。IP版本，可取值：IPV4、IPV6、IPv6FullChain，不区分大小写，默认值 IPV4。说明：取值为IPV6表示为IPV6 NAT64版本；取值为IPv6FullChain，表示为IPv6版本。</p>
	AddressIPVersion *string `json:"AddressIPVersion,omitnil,omitempty" name:"AddressIPVersion"`

	// <p>创建负载均衡的个数，默认值 1。创建个数不能超过账号所能创建的最大值，默认创建最大值为20。</p>
	Number *uint64 `json:"Number,omitnil,omitempty" name:"Number"`

	// <p>仅适用于公网且IP版本为IPv4的负载均衡。设置跨可用区容灾时的主可用区ID， 可用区 ID 和名称均支持，例如 100001 或 ap-guangzhou-1<br>注：主可用区是需要承载流量的可用区，备可用区默认不承载流量，主可用区不可用时才使用备可用区。</p>
	MasterZoneId *string `json:"MasterZoneId,omitnil,omitempty" name:"MasterZoneId"`

	// <p>仅适用于公网且IP版本为IPv4的负载均衡。可用区ID，可用区 ID 和名称均支持，指定可用区以创建负载均衡实例。如：100001 或 ap-guangzhou-1。</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>网络计费模式，其中的最大出带宽，仅对内网属性的性能容量型实例和公网属性的所有实例生效。</p>
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// <p>仅适用于公网负载均衡。目前仅广州、上海、南京、济南、杭州、福州、北京、石家庄、武汉、长沙、成都、重庆地域支持静态单线 IP 线路类型，如需体验，请联系商务经理申请。申请通过后，即可选择中国移动（CMCC）、中国联通（CUCC）或中国电信（CTCC）的运营商类型，网络计费模式只能使用按带宽包计费(BANDWIDTH_PACKAGE)。 如果不指定本参数，则默认使用BGP。可通过 <a href="https://cloud.tencent.com/document/api/214/70213">DescribeResources</a>  接口查询一个地域所支持的Isp。</p>
	VipIsp *string `json:"VipIsp,omitnil,omitempty" name:"VipIsp"`

	// <p>购买负载均衡的同时，给负载均衡打上标签，最大支持20个标签键值对。</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>指定VIP申请负载均衡。此参数选填，不填写此参数时自动分配VIP。IPv4和IPv6类型支持此参数，IPv6 NAT64类型不支持。<br>注意：当指定VIP创建内网实例、或公网IPv6 BGP实例时，若VIP不属于指定VPC子网的网段内时，会创建失败；若VIP已被占用，也会创建失败。</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>带宽包ID，可以通过 <a href="https://cloud.tencent.com/document/api/215/19209">DescribeBandwidthPackages</a> 接口获取。指定此参数时，网络计费方式（InternetAccessible.InternetChargeType）只支持按带宽包计费（BANDWIDTH_PACKAGE），带宽包的属性即为其结算方式。非上移用户购买的 IPv6 负载均衡实例，且运营商类型非 BGP 时 ，不支持指定具体带宽包id。</p>
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// <p>独占型实例信息。若创建独占型的内网负载均衡实例，则此参数必填。</p>
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitnil,omitempty" name:"ExclusiveCluster"`

	// <p>性能容量型规格。</p><ul><li>若需要创建性能容量型实例，则此参数必填，取值范围：<ul><li> clb.c2.medium：标准型规格 </li><li> clb.c3.small：高阶型1规格 </li><li> clb.c3.medium：高阶型2规格 </li><li> clb.c4.small：超强型1规格 </li><li> clb.c4.medium：超强型2规格 </li><li> clb.c4.large：超强型3规格 </li><li> clb.c4.xlarge：超强型4规格 </li></ul></li><li>中国站用户若需要创建共享型实例，则无需填写此参数。国际站用户不传该参数默认购买的是标准型实例。</li></ul>如需了解规格详情，请参见[实例规格对比](https://cloud.tencent.com/document/product/214/84689)。
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`

	// <p>集群ID，集群标识，在需要配置公有云独占集群或本地专有集群时使用。公有云独占集群申请请<a href="https://console.cloud.tencent.com/workorder/category">提交工单</a>，本地专有集群请参考<a href="https://cloud.tencent.com/document/product/1346">本地专有集群</a>描述。</p>
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// <p>用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。</p>
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// <p>是否支持绑定跨地域/跨Vpc绑定IP的功能。</p>
	SnatPro *bool `json:"SnatPro,omitnil,omitempty" name:"SnatPro"`

	// <p>开启绑定跨地域/跨Vpc绑定IP的功能后，创建SnatIp。</p>
	SnatIps []*SnatIp `json:"SnatIps,omitnil,omitempty" name:"SnatIps"`

	// <p>Stgw独占集群的标签。</p>
	ClusterTag *string `json:"ClusterTag,omitnil,omitempty" name:"ClusterTag"`

	// <p>仅适用于公网且IP版本为IPv4的负载均衡。设置跨可用区容灾时的备可用区ID，可用区 ID 和名称均支持，例如 100001 或 ap-guangzhou-1<br>注：备可用区是主可用区故障后，需要承载流量的可用区。可通过 <a href="https://cloud.tencent.com/document/api/214/70213">DescribeResources</a> 接口查询一个地域的主/备可用区的列表。【如果您需要体验该功能，请通过 <a href="https://console.cloud.tencent.com/workorder/category">工单申请</a>】</p>
	SlaveZoneId *string `json:"SlaveZoneId,omitnil,omitempty" name:"SlaveZoneId"`

	// <p>EIP 的唯一 ID，可以通过 <a href="https://cloud.tencent.com/document/product/215/16702">DescribeAddresses</a> 接口查询。形如：eip-qhx8udkc，仅适用于内网负载均衡绑定EIP。</p>
	EipAddressId *string `json:"EipAddressId,omitnil,omitempty" name:"EipAddressId"`

	// <p>Target是否放通来自CLB的流量。开启放通（true）：只验证CLB上的安全组；不开启放通（false）：需同时验证CLB和后端实例上的安全组。IPv6 CLB安全组默认放通，不需要传此参数。</p>
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitnil,omitempty" name:"LoadBalancerPassToTarget"`

	// <p>创建域名化负载均衡。</p>
	DynamicVip *bool `json:"DynamicVip,omitnil,omitempty" name:"DynamicVip"`

	// <p>网络出口</p>
	Egress *string `json:"Egress,omitnil,omitempty" name:"Egress"`

	// <p>负载均衡实例的预付费相关属性</p>
	LBChargePrepaid *LBChargePrepaid `json:"LBChargePrepaid,omitnil,omitempty" name:"LBChargePrepaid"`

	// <p>负载均衡实例计费类型，取值：POSTPAID_BY_HOUR，PREPAID，默认是POSTPAID_BY_HOUR。</p><p>枚举值：</p><ul><li>POSTPAID_BY_HOUR： 按量计费</li><li>PREPAID： 包年包月</li></ul>
	LBChargeType *string `json:"LBChargeType,omitnil,omitempty" name:"LBChargeType"`

	// <p>七层访问日志主题ID</p>
	AccessLogTopicId *string `json:"AccessLogTopicId,omitnil,omitempty" name:"AccessLogTopicId"`

	// <p>是否开启七层高级路由</p>
	AdvancedRoute *bool `json:"AdvancedRoute,omitnil,omitempty" name:"AdvancedRoute"`

	// <p>可用区亲和信息</p>
	AvailableZoneAffinityInfo *AvailableZoneAffinityInfo `json:"AvailableZoneAffinityInfo,omitnil,omitempty" name:"AvailableZoneAffinityInfo"`
}

type CreateLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// <p>负载均衡实例的网络类型：<br>OPEN：公网属性， INTERNAL：内网属性。</p>
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// <p>负载均衡实例的类型。1：通用的负载均衡实例，目前只支持传入1。</p>
	Forward *int64 `json:"Forward,omitnil,omitempty" name:"Forward"`

	// <p>负载均衡实例的名称，只在创建一个实例的时候才会生效。规则：1-80 个英文字母、汉字等国际通用语言字符，数字，连接线“-”、下划线“_”等常见字符（禁止Unicode补充字符，如emoji表情、生僻汉字等）。注意：如果名称与系统中已有负载均衡实例的名称相同，则系统将会自动生成此次创建的负载均衡实例的名称。</p>
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// <p>负载均衡后端目标设备所属的网络 ID，如vpc-12345678，可以通过 <a href="https://cloud.tencent.com/document/product/215/15778">DescribeVpcs</a> 接口获取。 不填此参数则默认为DefaultVPC。创建内网负载均衡实例时，此参数必填。</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>在私有网络内购买内网负载均衡实例的情况下，必须指定子网 ID，内网负载均衡实例的 VIP 将从这个子网中产生。<br>创建内网负载均衡实例，或者创建 IPv6FullChain 版本的负载均衡实例，此参数必填。<br>创建公网IPv4负载均衡实例时，不支持指定该参数。</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>负载均衡实例所属的项目 ID，默认项目 ID 为0。可以通过 <a href="https://cloud.tencent.com/document/api/651/78725">DescribeProject</a> 接口获取。不填此参数则视为默认项目。</p>
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// <p>仅适用于公网负载均衡。IP版本，可取值：IPV4、IPV6、IPv6FullChain，不区分大小写，默认值 IPV4。说明：取值为IPV6表示为IPV6 NAT64版本；取值为IPv6FullChain，表示为IPv6版本。</p>
	AddressIPVersion *string `json:"AddressIPVersion,omitnil,omitempty" name:"AddressIPVersion"`

	// <p>创建负载均衡的个数，默认值 1。创建个数不能超过账号所能创建的最大值，默认创建最大值为20。</p>
	Number *uint64 `json:"Number,omitnil,omitempty" name:"Number"`

	// <p>仅适用于公网且IP版本为IPv4的负载均衡。设置跨可用区容灾时的主可用区ID， 可用区 ID 和名称均支持，例如 100001 或 ap-guangzhou-1<br>注：主可用区是需要承载流量的可用区，备可用区默认不承载流量，主可用区不可用时才使用备可用区。</p>
	MasterZoneId *string `json:"MasterZoneId,omitnil,omitempty" name:"MasterZoneId"`

	// <p>仅适用于公网且IP版本为IPv4的负载均衡。可用区ID，可用区 ID 和名称均支持，指定可用区以创建负载均衡实例。如：100001 或 ap-guangzhou-1。</p>
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>网络计费模式，其中的最大出带宽，仅对内网属性的性能容量型实例和公网属性的所有实例生效。</p>
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// <p>仅适用于公网负载均衡。目前仅广州、上海、南京、济南、杭州、福州、北京、石家庄、武汉、长沙、成都、重庆地域支持静态单线 IP 线路类型，如需体验，请联系商务经理申请。申请通过后，即可选择中国移动（CMCC）、中国联通（CUCC）或中国电信（CTCC）的运营商类型，网络计费模式只能使用按带宽包计费(BANDWIDTH_PACKAGE)。 如果不指定本参数，则默认使用BGP。可通过 <a href="https://cloud.tencent.com/document/api/214/70213">DescribeResources</a>  接口查询一个地域所支持的Isp。</p>
	VipIsp *string `json:"VipIsp,omitnil,omitempty" name:"VipIsp"`

	// <p>购买负载均衡的同时，给负载均衡打上标签，最大支持20个标签键值对。</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>指定VIP申请负载均衡。此参数选填，不填写此参数时自动分配VIP。IPv4和IPv6类型支持此参数，IPv6 NAT64类型不支持。<br>注意：当指定VIP创建内网实例、或公网IPv6 BGP实例时，若VIP不属于指定VPC子网的网段内时，会创建失败；若VIP已被占用，也会创建失败。</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>带宽包ID，可以通过 <a href="https://cloud.tencent.com/document/api/215/19209">DescribeBandwidthPackages</a> 接口获取。指定此参数时，网络计费方式（InternetAccessible.InternetChargeType）只支持按带宽包计费（BANDWIDTH_PACKAGE），带宽包的属性即为其结算方式。非上移用户购买的 IPv6 负载均衡实例，且运营商类型非 BGP 时 ，不支持指定具体带宽包id。</p>
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// <p>独占型实例信息。若创建独占型的内网负载均衡实例，则此参数必填。</p>
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitnil,omitempty" name:"ExclusiveCluster"`

	// <p>性能容量型规格。</p><ul><li>若需要创建性能容量型实例，则此参数必填，取值范围：<ul><li> clb.c2.medium：标准型规格 </li><li> clb.c3.small：高阶型1规格 </li><li> clb.c3.medium：高阶型2规格 </li><li> clb.c4.small：超强型1规格 </li><li> clb.c4.medium：超强型2规格 </li><li> clb.c4.large：超强型3规格 </li><li> clb.c4.xlarge：超强型4规格 </li></ul></li><li>中国站用户若需要创建共享型实例，则无需填写此参数。国际站用户不传该参数默认购买的是标准型实例。</li></ul>如需了解规格详情，请参见[实例规格对比](https://cloud.tencent.com/document/product/214/84689)。
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`

	// <p>集群ID，集群标识，在需要配置公有云独占集群或本地专有集群时使用。公有云独占集群申请请<a href="https://console.cloud.tencent.com/workorder/category">提交工单</a>，本地专有集群请参考<a href="https://cloud.tencent.com/document/product/1346">本地专有集群</a>描述。</p>
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// <p>用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。</p>
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// <p>是否支持绑定跨地域/跨Vpc绑定IP的功能。</p>
	SnatPro *bool `json:"SnatPro,omitnil,omitempty" name:"SnatPro"`

	// <p>开启绑定跨地域/跨Vpc绑定IP的功能后，创建SnatIp。</p>
	SnatIps []*SnatIp `json:"SnatIps,omitnil,omitempty" name:"SnatIps"`

	// <p>Stgw独占集群的标签。</p>
	ClusterTag *string `json:"ClusterTag,omitnil,omitempty" name:"ClusterTag"`

	// <p>仅适用于公网且IP版本为IPv4的负载均衡。设置跨可用区容灾时的备可用区ID，可用区 ID 和名称均支持，例如 100001 或 ap-guangzhou-1<br>注：备可用区是主可用区故障后，需要承载流量的可用区。可通过 <a href="https://cloud.tencent.com/document/api/214/70213">DescribeResources</a> 接口查询一个地域的主/备可用区的列表。【如果您需要体验该功能，请通过 <a href="https://console.cloud.tencent.com/workorder/category">工单申请</a>】</p>
	SlaveZoneId *string `json:"SlaveZoneId,omitnil,omitempty" name:"SlaveZoneId"`

	// <p>EIP 的唯一 ID，可以通过 <a href="https://cloud.tencent.com/document/product/215/16702">DescribeAddresses</a> 接口查询。形如：eip-qhx8udkc，仅适用于内网负载均衡绑定EIP。</p>
	EipAddressId *string `json:"EipAddressId,omitnil,omitempty" name:"EipAddressId"`

	// <p>Target是否放通来自CLB的流量。开启放通（true）：只验证CLB上的安全组；不开启放通（false）：需同时验证CLB和后端实例上的安全组。IPv6 CLB安全组默认放通，不需要传此参数。</p>
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitnil,omitempty" name:"LoadBalancerPassToTarget"`

	// <p>创建域名化负载均衡。</p>
	DynamicVip *bool `json:"DynamicVip,omitnil,omitempty" name:"DynamicVip"`

	// <p>网络出口</p>
	Egress *string `json:"Egress,omitnil,omitempty" name:"Egress"`

	// <p>负载均衡实例的预付费相关属性</p>
	LBChargePrepaid *LBChargePrepaid `json:"LBChargePrepaid,omitnil,omitempty" name:"LBChargePrepaid"`

	// <p>负载均衡实例计费类型，取值：POSTPAID_BY_HOUR，PREPAID，默认是POSTPAID_BY_HOUR。</p><p>枚举值：</p><ul><li>POSTPAID_BY_HOUR： 按量计费</li><li>PREPAID： 包年包月</li></ul>
	LBChargeType *string `json:"LBChargeType,omitnil,omitempty" name:"LBChargeType"`

	// <p>七层访问日志主题ID</p>
	AccessLogTopicId *string `json:"AccessLogTopicId,omitnil,omitempty" name:"AccessLogTopicId"`

	// <p>是否开启七层高级路由</p>
	AdvancedRoute *bool `json:"AdvancedRoute,omitnil,omitempty" name:"AdvancedRoute"`

	// <p>可用区亲和信息</p>
	AvailableZoneAffinityInfo *AvailableZoneAffinityInfo `json:"AvailableZoneAffinityInfo,omitnil,omitempty" name:"AvailableZoneAffinityInfo"`
}

func (r *CreateLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerType")
	delete(f, "Forward")
	delete(f, "LoadBalancerName")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ProjectId")
	delete(f, "AddressIPVersion")
	delete(f, "Number")
	delete(f, "MasterZoneId")
	delete(f, "ZoneId")
	delete(f, "InternetAccessible")
	delete(f, "VipIsp")
	delete(f, "Tags")
	delete(f, "Vip")
	delete(f, "BandwidthPackageId")
	delete(f, "ExclusiveCluster")
	delete(f, "SlaType")
	delete(f, "ClusterIds")
	delete(f, "ClientToken")
	delete(f, "SnatPro")
	delete(f, "SnatIps")
	delete(f, "ClusterTag")
	delete(f, "SlaveZoneId")
	delete(f, "EipAddressId")
	delete(f, "LoadBalancerPassToTarget")
	delete(f, "DynamicVip")
	delete(f, "Egress")
	delete(f, "LBChargePrepaid")
	delete(f, "LBChargeType")
	delete(f, "AccessLogTopicId")
	delete(f, "AdvancedRoute")
	delete(f, "AvailableZoneAffinityInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadBalancerResponseParams struct {
	// <p>由负载均衡实例唯一 ID 组成的数组。<br>存在某些场景，如创建出现延迟时，此字段可能返回为空；此时可以根据接口返回的RequestId或DealName参数，通过DescribeTaskStatus接口查询创建的资源ID。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// <p>订单号。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *CreateLoadBalancerResponseParams `json:"Response"`
}

func (r *CreateLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadBalancerSnatIpsRequestParams struct {
	// 负载均衡唯一性ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口查询。例如：lb-12345678。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 添加的SnatIp信息，可指定IP申请，或者指定子网自动申请。可以通过 [DescribeSubnets](https://cloud.tencent.com/document/api/215/15784) 查询获取，单个CLB实例可申请的默认上限为10个。
	SnatIps []*SnatIp `json:"SnatIps,omitnil,omitempty" name:"SnatIps"`

	// 添加的SnatIp的个数，可与SnatIps一起使用，但若指定IP时，则不能指定创建的SnatIp个数。默认值为1，数量上限与用户配置有关，默认上限为10。
	Number *uint64 `json:"Number,omitnil,omitempty" name:"Number"`
}

type CreateLoadBalancerSnatIpsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡唯一性ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口查询。例如：lb-12345678。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 添加的SnatIp信息，可指定IP申请，或者指定子网自动申请。可以通过 [DescribeSubnets](https://cloud.tencent.com/document/api/215/15784) 查询获取，单个CLB实例可申请的默认上限为10个。
	SnatIps []*SnatIp `json:"SnatIps,omitnil,omitempty" name:"SnatIps"`

	// 添加的SnatIp的个数，可与SnatIps一起使用，但若指定IP时，则不能指定创建的SnatIp个数。默认值为1，数量上限与用户配置有关，默认上限为10。
	Number *uint64 `json:"Number,omitnil,omitempty" name:"Number"`
}

func (r *CreateLoadBalancerSnatIpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancerSnatIpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "SnatIps")
	delete(f, "Number")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoadBalancerSnatIpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadBalancerSnatIpsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLoadBalancerSnatIpsResponse struct {
	*tchttp.BaseResponse
	Response *CreateLoadBalancerSnatIpsResponseParams `json:"Response"`
}

func (r *CreateLoadBalancerSnatIpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLoadBalancerSnatIpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelRequestParams struct {
	// <p>接入类型：PublicBYOK/PublicCustom/PrivateCustom</p>
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// <p>模型提供商</p>
	ModelProvider *string `json:"ModelProvider,omitnil,omitempty" name:"ModelProvider"`

	// <p>通用模型标识列表</p>
	ModelIds []*ModelItem `json:"ModelIds,omitnil,omitempty" name:"ModelIds"`

	// <p>Key 列表</p>
	Keys []*KeyItem `json:"Keys,omitnil,omitempty" name:"Keys"`

	// <p>BYOK ID(在自定义模型时在部署网络后必须填写)</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>服务供应商(创建BYOK自定义名称)。</p>
	ServiceProviderName *string `json:"ServiceProviderName,omitnil,omitempty" name:"ServiceProviderName"`

	// <p>模型协议</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>API Base URL</p>
	ApiBase *string `json:"ApiBase,omitnil,omitempty" name:"ApiBase"`

	// <p>VPC ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>子网 ID</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>转发请求时添加的Host请求头</p>
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`

	// <p>标签信息</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>是否校验服务提供商的SSL证书</p>
	VerifySSL *bool `json:"VerifySSL,omitnil,omitempty" name:"VerifySSL"`
}

type CreateModelRequest struct {
	*tchttp.BaseRequest
	
	// <p>接入类型：PublicBYOK/PublicCustom/PrivateCustom</p>
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// <p>模型提供商</p>
	ModelProvider *string `json:"ModelProvider,omitnil,omitempty" name:"ModelProvider"`

	// <p>通用模型标识列表</p>
	ModelIds []*ModelItem `json:"ModelIds,omitnil,omitempty" name:"ModelIds"`

	// <p>Key 列表</p>
	Keys []*KeyItem `json:"Keys,omitnil,omitempty" name:"Keys"`

	// <p>BYOK ID(在自定义模型时在部署网络后必须填写)</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>服务供应商(创建BYOK自定义名称)。</p>
	ServiceProviderName *string `json:"ServiceProviderName,omitnil,omitempty" name:"ServiceProviderName"`

	// <p>模型协议</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>API Base URL</p>
	ApiBase *string `json:"ApiBase,omitnil,omitempty" name:"ApiBase"`

	// <p>VPC ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>子网 ID</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>转发请求时添加的Host请求头</p>
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`

	// <p>标签信息</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>是否校验服务提供商的SSL证书</p>
	VerifySSL *bool `json:"VerifySSL,omitnil,omitempty" name:"VerifySSL"`
}

func (r *CreateModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessType")
	delete(f, "ModelProvider")
	delete(f, "ModelIds")
	delete(f, "Keys")
	delete(f, "ServiceProviderId")
	delete(f, "ServiceProviderName")
	delete(f, "Protocol")
	delete(f, "ApiBase")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "HostHeader")
	delete(f, "Tags")
	delete(f, "VerifySSL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelResponseParams struct {
	// <p>服务供应商ID</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>生成的 Key ID 列表</p>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateModelResponse struct {
	*tchttp.BaseResponse
	Response *CreateModelResponseParams `json:"Response"`
}

func (r *CreateModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelRouterRequestParams struct {
	// <p>模型路由类型</p><p>枚举值：</p><ul><li>Shared： 共享型</li><li>Enterprise： 企业级</li></ul>
	ModelRouterType *string `json:"ModelRouterType,omitnil,omitempty" name:"ModelRouterType"`

	// <p>关联的积分预算ID</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>证书ID</p><p>入参限制：当Schema为HTTPS时，该参数必传</p>
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// <p>集群信息</p>
	ClusterInfo *ClusterInfo `json:"ClusterInfo,omitnil,omitempty" name:"ClusterInfo"`

	// <p>模型路由实例名称</p><p>默认值：-</p>
	ModelRouterName *string `json:"ModelRouterName,omitnil,omitempty" name:"ModelRouterName"`

	// <p>网络类型</p><p>枚举值：</p><ul><li>Internet： 公网</li><li>Intranet： 内网</li></ul>
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// <p>模型路由的监听端口</p><p>取值范围：[1, 65535]</p>
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>限速配置</p>
	RateLimitConfig *RateLimitConfigForModelRouter `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`

	// <p>路由配置</p>
	RouterSetting *RouterSettingWithoutFallBack `json:"RouterSetting,omitnil,omitempty" name:"RouterSetting"`

	// <p>模型路由实例的网络协议</p><p>枚举值：</p><ul><li>HTTP： HTTP协议</li><li>HTTPS： HTTPS协议</li></ul>
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// <p>模型路由实例所属子网的ID</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>标签</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>模型路由实例所属VPC的ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type CreateModelRouterRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由类型</p><p>枚举值：</p><ul><li>Shared： 共享型</li><li>Enterprise： 企业级</li></ul>
	ModelRouterType *string `json:"ModelRouterType,omitnil,omitempty" name:"ModelRouterType"`

	// <p>关联的积分预算ID</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>证书ID</p><p>入参限制：当Schema为HTTPS时，该参数必传</p>
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// <p>集群信息</p>
	ClusterInfo *ClusterInfo `json:"ClusterInfo,omitnil,omitempty" name:"ClusterInfo"`

	// <p>模型路由实例名称</p><p>默认值：-</p>
	ModelRouterName *string `json:"ModelRouterName,omitnil,omitempty" name:"ModelRouterName"`

	// <p>网络类型</p><p>枚举值：</p><ul><li>Internet： 公网</li><li>Intranet： 内网</li></ul>
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// <p>模型路由的监听端口</p><p>取值范围：[1, 65535]</p>
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>限速配置</p>
	RateLimitConfig *RateLimitConfigForModelRouter `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`

	// <p>路由配置</p>
	RouterSetting *RouterSettingWithoutFallBack `json:"RouterSetting,omitnil,omitempty" name:"RouterSetting"`

	// <p>模型路由实例的网络协议</p><p>枚举值：</p><ul><li>HTTP： HTTP协议</li><li>HTTPS： HTTPS协议</li></ul>
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// <p>模型路由实例所属子网的ID</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>标签</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>模型路由实例所属VPC的ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

func (r *CreateModelRouterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelRouterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterType")
	delete(f, "BudgetId")
	delete(f, "CertId")
	delete(f, "ClusterInfo")
	delete(f, "ModelRouterName")
	delete(f, "NetworkType")
	delete(f, "Port")
	delete(f, "RateLimitConfig")
	delete(f, "RouterSetting")
	delete(f, "Schema")
	delete(f, "SubnetId")
	delete(f, "Tags")
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateModelRouterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelRouterResourcePackageRequestParams struct {
	// <p>模型路由资源包容量</p><p>取值范围：[1000, 10000000]</p><p>单次购买的模型路由资源包容量下限为1000，上限为10000000</p>
	ModelRouterResourcePackageAmount *uint64 `json:"ModelRouterResourcePackageAmount,omitnil,omitempty" name:"ModelRouterResourcePackageAmount"`

	// <p>是否自动续订。</p><p>0:不自动续订, 1:用尽到期续订</p>
	AutoPurchaseFlag *uint64 `json:"AutoPurchaseFlag,omitnil,omitempty" name:"AutoPurchaseFlag"`

	// <p>该笔订单是否自动选择代金券</p><p>默认值：false（不自动选择代金券）</p><p>true时会为本笔订单自动匹配满足条件、最优惠的代金券</p>
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

type CreateModelRouterResourcePackageRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由资源包容量</p><p>取值范围：[1000, 10000000]</p><p>单次购买的模型路由资源包容量下限为1000，上限为10000000</p>
	ModelRouterResourcePackageAmount *uint64 `json:"ModelRouterResourcePackageAmount,omitnil,omitempty" name:"ModelRouterResourcePackageAmount"`

	// <p>是否自动续订。</p><p>0:不自动续订, 1:用尽到期续订</p>
	AutoPurchaseFlag *uint64 `json:"AutoPurchaseFlag,omitnil,omitempty" name:"AutoPurchaseFlag"`

	// <p>该笔订单是否自动选择代金券</p><p>默认值：false（不自动选择代金券）</p><p>true时会为本笔订单自动匹配满足条件、最优惠的代金券</p>
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

func (r *CreateModelRouterResourcePackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelRouterResourcePackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterResourcePackageAmount")
	delete(f, "AutoPurchaseFlag")
	delete(f, "AutoVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateModelRouterResourcePackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelRouterResourcePackageResponseParams struct {
	// <p>模型路由资源包Id</p>
	ModelRouterResourcePackageIds []*string `json:"ModelRouterResourcePackageIds,omitnil,omitempty" name:"ModelRouterResourcePackageIds"`

	// <p>订单号</p>
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateModelRouterResourcePackageResponse struct {
	*tchttp.BaseResponse
	Response *CreateModelRouterResourcePackageResponseParams `json:"Response"`
}

func (r *CreateModelRouterResourcePackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelRouterResourcePackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModelRouterResponseParams struct {
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateModelRouterResponse struct {
	*tchttp.BaseResponse
	Response *CreateModelRouterResponseParams `json:"Response"`
}

func (r *CreateModelRouterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModelRouterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleRequestParams struct {
	// 负载均衡实例 ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口获取
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 监听器 ID，可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口获取
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 新建转发规则的信息。
	Rules []*RuleInput `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type CreateRuleRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口获取
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 监听器 ID，可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口获取
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 新建转发规则的信息。
	Rules []*RuleInput `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *CreateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleResponseParams struct {
	// 创建的转发规则的唯一标识数组。
	LocationIds []*string `json:"LocationIds,omitnil,omitempty" name:"LocationIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateRuleResponseParams `json:"Response"`
}

func (r *CreateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTargetGroupRequestParams struct {
	// <p>目标组名称。命名规则：1-80 个英文字母、汉字等国际通用语言字符，数字，连接线“-”、下划线“_”等常见字符（禁止Unicode补充字符，如emoji表情、生僻汉字等）。</p>
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`

	// <p>目标组的vpcId属性，不填则使用默认vpc。</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>目标组的默认端口， 后续添加服务器时可使用该默认端口。全监听目标组不支持此参数，非全监听目标组Port和TargetGroupInstances.N中的port二者必填其一。</p>
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>目标组绑定的后端服务器，单次最多支持50个。</p>
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`

	// <p>目标组类型，当前支持v1(旧版目标组), v2(新版目标组), 默认为v1(旧版目标组)。</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>目标组后端转发协议。v2新版目标组该项必填。目前支持TCP、UDP、HTTP、HTTPS、GRPC。</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>健康检查。</p>
	HealthCheck *TargetGroupHealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// <p>调度算法，仅V2新版目标组，且后端转发协议为(HTTP|HTTPS|GRPC)时该参数有效。可选值：</p><li>WRR:按权重轮询。</li><li>LEAST_CONN:最小连接数。</li><li>IP_HASH:按IP哈希。</li><li>默认为 WRR。</li>
	ScheduleAlgorithm *string `json:"ScheduleAlgorithm,omitnil,omitempty" name:"ScheduleAlgorithm"`

	// <p>标签。</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>后端服务默认权重, 其中：</p><ul><li>取值范围[0, 100]</li><li>设置该值后，添加后端服务到目标组时， 若后端服务不单独设置权重， 则使用这里的默认权重。 </li><li>v1 目标组类型不支持设置 Weight 参数。</li></ul>
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// <p>全监听目标组标识，true表示是全监听目标组，false表示不是全监听目标组。仅V2新版类型目标组支持该参数。</p>
	FullListenSwitch *bool `json:"FullListenSwitch,omitnil,omitempty" name:"FullListenSwitch"`

	// <p>是否开启长连接，此参数仅适用于HTTP/HTTPS目标组，0:关闭；1:开启， 默认关闭。</p>
	KeepaliveEnable *bool `json:"KeepaliveEnable,omitnil,omitempty" name:"KeepaliveEnable"`

	// <p>会话保持时间，单位：秒。可选值：30~3600，默认 0，表示不开启。仅V2新版且后端转发协议为HTTP/HTTPS/GRPC目标组支持该参数。</p>
	SessionExpireTime *uint64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// <p>IP版本类型。</p>
	IpVersion *string `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`

	// <p>是否开启SNAT（源IP替换），True（开启）、False（关闭）。默认为关闭。注意：SnatEnable开启时会替换客户端源IP，此时<code>透传客户端源IP</code>选项关闭，反之亦然。</p>
	SnatEnable *bool `json:"SnatEnable,omitnil,omitempty" name:"SnatEnable"`
}

type CreateTargetGroupRequest struct {
	*tchttp.BaseRequest
	
	// <p>目标组名称。命名规则：1-80 个英文字母、汉字等国际通用语言字符，数字，连接线“-”、下划线“_”等常见字符（禁止Unicode补充字符，如emoji表情、生僻汉字等）。</p>
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`

	// <p>目标组的vpcId属性，不填则使用默认vpc。</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>目标组的默认端口， 后续添加服务器时可使用该默认端口。全监听目标组不支持此参数，非全监听目标组Port和TargetGroupInstances.N中的port二者必填其一。</p>
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>目标组绑定的后端服务器，单次最多支持50个。</p>
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`

	// <p>目标组类型，当前支持v1(旧版目标组), v2(新版目标组), 默认为v1(旧版目标组)。</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>目标组后端转发协议。v2新版目标组该项必填。目前支持TCP、UDP、HTTP、HTTPS、GRPC。</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>健康检查。</p>
	HealthCheck *TargetGroupHealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// <p>调度算法，仅V2新版目标组，且后端转发协议为(HTTP|HTTPS|GRPC)时该参数有效。可选值：</p><li>WRR:按权重轮询。</li><li>LEAST_CONN:最小连接数。</li><li>IP_HASH:按IP哈希。</li><li>默认为 WRR。</li>
	ScheduleAlgorithm *string `json:"ScheduleAlgorithm,omitnil,omitempty" name:"ScheduleAlgorithm"`

	// <p>标签。</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>后端服务默认权重, 其中：</p><ul><li>取值范围[0, 100]</li><li>设置该值后，添加后端服务到目标组时， 若后端服务不单独设置权重， 则使用这里的默认权重。 </li><li>v1 目标组类型不支持设置 Weight 参数。</li></ul>
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// <p>全监听目标组标识，true表示是全监听目标组，false表示不是全监听目标组。仅V2新版类型目标组支持该参数。</p>
	FullListenSwitch *bool `json:"FullListenSwitch,omitnil,omitempty" name:"FullListenSwitch"`

	// <p>是否开启长连接，此参数仅适用于HTTP/HTTPS目标组，0:关闭；1:开启， 默认关闭。</p>
	KeepaliveEnable *bool `json:"KeepaliveEnable,omitnil,omitempty" name:"KeepaliveEnable"`

	// <p>会话保持时间，单位：秒。可选值：30~3600，默认 0，表示不开启。仅V2新版且后端转发协议为HTTP/HTTPS/GRPC目标组支持该参数。</p>
	SessionExpireTime *uint64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// <p>IP版本类型。</p>
	IpVersion *string `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`

	// <p>是否开启SNAT（源IP替换），True（开启）、False（关闭）。默认为关闭。注意：SnatEnable开启时会替换客户端源IP，此时<code>透传客户端源IP</code>选项关闭，反之亦然。</p>
	SnatEnable *bool `json:"SnatEnable,omitnil,omitempty" name:"SnatEnable"`
}

func (r *CreateTargetGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTargetGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupName")
	delete(f, "VpcId")
	delete(f, "Port")
	delete(f, "TargetGroupInstances")
	delete(f, "Type")
	delete(f, "Protocol")
	delete(f, "HealthCheck")
	delete(f, "ScheduleAlgorithm")
	delete(f, "Tags")
	delete(f, "Weight")
	delete(f, "FullListenSwitch")
	delete(f, "KeepaliveEnable")
	delete(f, "SessionExpireTime")
	delete(f, "IpVersion")
	delete(f, "SnatEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTargetGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTargetGroupResponseParams struct {
	// <p>创建目标组后生成的id</p>
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTargetGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateTargetGroupResponseParams `json:"Response"`
}

func (r *CreateTargetGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTargetGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicRequestParams struct {
	// 日志主题的名称。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 主题分区Partition的数量，不传参默认创建1个，最大创建允许10个，分裂/合并操作会改变分区数量，整体上限50个。
	PartitionCount *uint64 `json:"PartitionCount,omitnil,omitempty" name:"PartitionCount"`

	// 日志类型，ACCESS：访问日志，HEALTH：健康检查日志，默认ACCESS。
	TopicType *string `json:"TopicType,omitnil,omitempty" name:"TopicType"`

	// 存储时间，单位天，默认为 30。
	// - 日志接入标准存储时，支持1至3600天，值为3640时代表永久保存。
	// - 日志接入低频存储时，支持7至3600天，值为3640时代表永久保存。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 日志主题的存储类型，可选值 HOT（标准存储），COLD（低频存储）；默认为HOT。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题的名称。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 主题分区Partition的数量，不传参默认创建1个，最大创建允许10个，分裂/合并操作会改变分区数量，整体上限50个。
	PartitionCount *uint64 `json:"PartitionCount,omitnil,omitempty" name:"PartitionCount"`

	// 日志类型，ACCESS：访问日志，HEALTH：健康检查日志，默认ACCESS。
	TopicType *string `json:"TopicType,omitnil,omitempty" name:"TopicType"`

	// 存储时间，单位天，默认为 30。
	// - 日志接入标准存储时，支持1至3600天，值为3640时代表永久保存。
	// - 日志接入低频存储时，支持7至3600天，值为3640时代表永久保存。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 日志主题的存储类型，可选值 HOT（标准存储），COLD（低频存储）；默认为HOT。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`
}

func (r *CreateTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "PartitionCount")
	delete(f, "TopicType")
	delete(f, "Period")
	delete(f, "StorageType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicResponseParams struct {
	// 日志主题的 ID。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse
	Response *CreateTopicResponseParams `json:"Response"`
}

func (r *CreateTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserGroupRequestParams struct {
	// <p>模型路由实例ID。用户组将创建在该实例下。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>用户组名称。必填。同一模型路由实例下名称唯一，长度不超过255个字符。</p>
	UserGroupName *string `json:"UserGroupName,omitnil,omitempty" name:"UserGroupName"`

	// <p>建组时直接关联的预算 ID（须为已存在的 Budget）。关联后由该 Budget 统一管理本组的消费上限与限速。不传表示不关联预算，可建组后再用 AssociateBudget 关联。</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>用户组意图路由白名单（ir-xxx）。每一项须为该实例已创建的意图路由名。命中意图路由名时其内部真实模型自动豁免白名单。为空表示不授权任何意图路由。</p>
	IntentRouters []*string `json:"IntentRouters,omitnil,omitempty" name:"IntentRouters"`

	// <p>建组时同时绑定的已有 Key ID 列表，最多100个。每个 Key 须属于同一模型路由实例。建组与绑定为一个原子异步任务，任一 Key 失败则整组创建回滚。不传表示建空组。</p>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// <p>用户组真实模型白名单。每一项须为该实例已关联的模型名。为空表示不在用户组层限制真实模型（按实例层白名单生效）。</p>
	Models []*string `json:"Models,omitnil,omitempty" name:"Models"`

	// <p>标签列表，最多50个。</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID。用户组将创建在该实例下。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>用户组名称。必填。同一模型路由实例下名称唯一，长度不超过255个字符。</p>
	UserGroupName *string `json:"UserGroupName,omitnil,omitempty" name:"UserGroupName"`

	// <p>建组时直接关联的预算 ID（须为已存在的 Budget）。关联后由该 Budget 统一管理本组的消费上限与限速。不传表示不关联预算，可建组后再用 AssociateBudget 关联。</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>用户组意图路由白名单（ir-xxx）。每一项须为该实例已创建的意图路由名。命中意图路由名时其内部真实模型自动豁免白名单。为空表示不授权任何意图路由。</p>
	IntentRouters []*string `json:"IntentRouters,omitnil,omitempty" name:"IntentRouters"`

	// <p>建组时同时绑定的已有 Key ID 列表，最多100个。每个 Key 须属于同一模型路由实例。建组与绑定为一个原子异步任务，任一 Key 失败则整组创建回滚。不传表示建空组。</p>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// <p>用户组真实模型白名单。每一项须为该实例已关联的模型名。为空表示不在用户组层限制真实模型（按实例层白名单生效）。</p>
	Models []*string `json:"Models,omitnil,omitempty" name:"Models"`

	// <p>标签列表，最多50个。</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "UserGroupName")
	delete(f, "BudgetId")
	delete(f, "IntentRouters")
	delete(f, "KeyIds")
	delete(f, "Models")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserGroupResponseParams struct {
	// <p>新建用户组的ID。</p>
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserGroupResponseParams `json:"Response"`
}

func (r *CreateUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatedKey struct {
	// <p>明文Key</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>Key的ID</p>
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// <p>Key的名称</p>
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`
}

type CreditUsage struct {
	// <p>Budget刷新周期。</p><p>枚举值：</p><ul><li>1d：按天刷新</li><li>7d：按周刷新</li><li>30d：按月刷新</li></ul><p>仅在 CreditUsageSet 数组元素中返回。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BudgetDuration *string `json:"BudgetDuration,omitnil,omitempty" name:"BudgetDuration"`

	// <p>下次刷新时间。</p><p>用户组关联Budget且Budget设置重置周期时返回；未关联Budget或未设置重置周期时为空。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BudgetResetAt *string `json:"BudgetResetAt,omitnil,omitempty" name:"BudgetResetAt"`

	// <p>Credit上限。</p><p>用户组关联Budget且Budget设置最大预算时返回；未设置最大预算时为空。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Limit *float64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>用户组已使用的Credit数量。</p>
	Used *float64 `json:"Used,omitnil,omitempty" name:"Used"`
}

type CrossTargets struct {
	// 本地私有网络ID，即负载均衡的VpcId。
	LocalVpcId *string `json:"LocalVpcId,omitnil,omitempty" name:"LocalVpcId"`

	// 子机或网卡所属的私有网络ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子机或网卡的IP地址
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// 子机或网卡所属的私有网络名称。
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 子机的网卡ID。
	EniId *string `json:"EniId,omitnil,omitempty" name:"EniId"`

	// 子机实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 子机实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 子机或者网卡所属的地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

// Predefined struct for user
type DeleteBudgetsRequestParams struct {
	// <p>要删除的Budget ID列表。</p>
	BudgetIds []*string `json:"BudgetIds,omitnil,omitempty" name:"BudgetIds"`
}

type DeleteBudgetsRequest struct {
	*tchttp.BaseRequest
	
	// <p>要删除的Budget ID列表。</p>
	BudgetIds []*string `json:"BudgetIds,omitnil,omitempty" name:"BudgetIds"`
}

func (r *DeleteBudgetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBudgetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BudgetIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteBudgetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteBudgetsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteBudgetsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteBudgetsResponseParams `json:"Response"`
}

func (r *DeleteBudgetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteBudgetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntentRouterRequestParams struct {
	// <p>模型路由实例ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>意图路由ID（ir-xxx格式）。</p>
	IntentRouterId *string `json:"IntentRouterId,omitnil,omitempty" name:"IntentRouterId"`
}

type DeleteIntentRouterRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>意图路由ID（ir-xxx格式）。</p>
	IntentRouterId *string `json:"IntentRouterId,omitnil,omitempty" name:"IntentRouterId"`
}

func (r *DeleteIntentRouterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntentRouterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "IntentRouterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIntentRouterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntentRouterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteIntentRouterResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIntentRouterResponseParams `json:"Response"`
}

func (r *DeleteIntentRouterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntentRouterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteKeysRequestParams struct {
	// <p>模型路由ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>key的ID列表</p>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

type DeleteKeysRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>key的ID列表</p>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

func (r *DeleteKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "KeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteKeysResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteKeysResponse struct {
	*tchttp.BaseResponse
	Response *DeleteKeysResponseParams `json:"Response"`
}

func (r *DeleteKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteListenerRequestParams struct {
	// 负载均衡实例ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 要删除的监听器ID，可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口查询。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`
}

type DeleteListenerRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 要删除的监听器ID，可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口查询。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`
}

func (r *DeleteListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteListenerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteListenerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteListenerResponseParams `json:"Response"`
}

func (r *DeleteListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancerListenersRequestParams struct {
	// 负载均衡实例ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 指定删除的监听器ID数组，最大为20个。若不填则删除负载均衡的所有监听器，可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口查询。
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`
}

type DeleteLoadBalancerListenersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 指定删除的监听器ID数组，最大为20个。若不填则删除负载均衡的所有监听器，可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口查询。
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`
}

func (r *DeleteLoadBalancerListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoadBalancerListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancerListenersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLoadBalancerListenersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoadBalancerListenersResponseParams `json:"Response"`
}

func (r *DeleteLoadBalancerListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancerRequestParams struct {
	// 要删除的负载均衡实例 ID 数组，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口获取，数组大小最大支持20。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 是否强制删除clb。true表示强制删除，false表示不是强制删除，需要做拦截校验。
	// 默认为false。
	// 以下几种情况会默认拦截删除操作，如果触发情况1、2但确认强制删除则需要传强制校验参数ForceDelete为true。
	// 1、删除后端绑定大于等于 20 个 RS 的实例时。
	// 2、删除后端有 RS 且 5 分钟 内“出/入带宽”峰值取大 > 10Mbps 的实例时。
	// 3、单地域内 5 分钟 内删除大于等于 30 个实例时。
	ForceDelete *bool `json:"ForceDelete,omitnil,omitempty" name:"ForceDelete"`
}

type DeleteLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的负载均衡实例 ID 数组，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口获取，数组大小最大支持20。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 是否强制删除clb。true表示强制删除，false表示不是强制删除，需要做拦截校验。
	// 默认为false。
	// 以下几种情况会默认拦截删除操作，如果触发情况1、2但确认强制删除则需要传强制校验参数ForceDelete为true。
	// 1、删除后端绑定大于等于 20 个 RS 的实例时。
	// 2、删除后端有 RS 且 5 分钟 内“出/入带宽”峰值取大 > 10Mbps 的实例时。
	// 3、单地域内 5 分钟 内删除大于等于 30 个实例时。
	ForceDelete *bool `json:"ForceDelete,omitnil,omitempty" name:"ForceDelete"`
}

func (r *DeleteLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	delete(f, "ForceDelete")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoadBalancerResponseParams `json:"Response"`
}

func (r *DeleteLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancerSnatIpsRequestParams struct {
	// 负载均衡唯一ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。例如：lb-12345678。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 删除SnatIp地址数组，最大支持删除数量为20个。
	Ips []*string `json:"Ips,omitnil,omitempty" name:"Ips"`
}

type DeleteLoadBalancerSnatIpsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡唯一ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。例如：lb-12345678。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 删除SnatIp地址数组，最大支持删除数量为20个。
	Ips []*string `json:"Ips,omitnil,omitempty" name:"Ips"`
}

func (r *DeleteLoadBalancerSnatIpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerSnatIpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "Ips")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoadBalancerSnatIpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancerSnatIpsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLoadBalancerSnatIpsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoadBalancerSnatIpsResponseParams `json:"Response"`
}

func (r *DeleteLoadBalancerSnatIpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancerSnatIpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelRequestParams struct {
	// <p>服务提供商ID列表</p>
	ServiceProviderIds []*string `json:"ServiceProviderIds,omitnil,omitempty" name:"ServiceProviderIds"`
}

type DeleteModelRequest struct {
	*tchttp.BaseRequest
	
	// <p>服务提供商ID列表</p>
	ServiceProviderIds []*string `json:"ServiceProviderIds,omitnil,omitempty" name:"ServiceProviderIds"`
}

func (r *DeleteModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceProviderIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteModelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteModelResponseParams `json:"Response"`
}

func (r *DeleteModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelRoutersRequestParams struct {
	// <p>模型路由实例ID列表</p>
	ModelRouterIds []*string `json:"ModelRouterIds,omitnil,omitempty" name:"ModelRouterIds"`
}

type DeleteModelRoutersRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID列表</p>
	ModelRouterIds []*string `json:"ModelRouterIds,omitnil,omitempty" name:"ModelRouterIds"`
}

func (r *DeleteModelRoutersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelRoutersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteModelRoutersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteModelRoutersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteModelRoutersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteModelRoutersResponseParams `json:"Response"`
}

func (r *DeleteModelRoutersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteModelRoutersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRewriteRequestParams struct {
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 源监听器ID。
	SourceListenerId *string `json:"SourceListenerId,omitnil,omitempty" name:"SourceListenerId"`

	// 目标监听器ID。
	TargetListenerId *string `json:"TargetListenerId,omitnil,omitempty" name:"TargetListenerId"`

	// 转发规则之间的重定向关系。
	RewriteInfos []*RewriteLocationMap `json:"RewriteInfos,omitnil,omitempty" name:"RewriteInfos"`
}

type DeleteRewriteRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 源监听器ID。
	SourceListenerId *string `json:"SourceListenerId,omitnil,omitempty" name:"SourceListenerId"`

	// 目标监听器ID。
	TargetListenerId *string `json:"TargetListenerId,omitnil,omitempty" name:"TargetListenerId"`

	// 转发规则之间的重定向关系。
	RewriteInfos []*RewriteLocationMap `json:"RewriteInfos,omitnil,omitempty" name:"RewriteInfos"`
}

func (r *DeleteRewriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRewriteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "SourceListenerId")
	delete(f, "TargetListenerId")
	delete(f, "RewriteInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRewriteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRewriteResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRewriteResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRewriteResponseParams `json:"Response"`
}

func (r *DeleteRewriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRewriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRuleRequestParams struct {
	// 负载均衡实例ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID，可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口查询。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 要删除的转发规则的ID组成的数组，可以通过 [DescribeLoadBalancersDetail](https://cloud.tencent.com/document/api/214/46916) 接口查询。
	LocationIds []*string `json:"LocationIds,omitnil,omitempty" name:"LocationIds"`

	// 要删除的转发规则的域名，如果是多域名，可以指定多域名列表中的任意一个，可以通过 [DescribeLoadBalancersDetail](https://cloud.tencent.com/document/api/214/46916) 接口查询。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 要删除的转发规则的转发路径，可以通过 [DescribeLoadBalancersDetail](https://cloud.tencent.com/document/api/214/46916) 接口查询。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 监听器下必须配置一个默认域名，当需要删除默认域名时，可以指定另一个域名作为新的默认域名，如果新的默认域名是多域名，可以指定多域名列表中的任意一个，可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口查询。
	NewDefaultServerDomain *string `json:"NewDefaultServerDomain,omitnil,omitempty" name:"NewDefaultServerDomain"`
}

type DeleteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID，可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口查询。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 要删除的转发规则的ID组成的数组，可以通过 [DescribeLoadBalancersDetail](https://cloud.tencent.com/document/api/214/46916) 接口查询。
	LocationIds []*string `json:"LocationIds,omitnil,omitempty" name:"LocationIds"`

	// 要删除的转发规则的域名，如果是多域名，可以指定多域名列表中的任意一个，可以通过 [DescribeLoadBalancersDetail](https://cloud.tencent.com/document/api/214/46916) 接口查询。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 要删除的转发规则的转发路径，可以通过 [DescribeLoadBalancersDetail](https://cloud.tencent.com/document/api/214/46916) 接口查询。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 监听器下必须配置一个默认域名，当需要删除默认域名时，可以指定另一个域名作为新的默认域名，如果新的默认域名是多域名，可以指定多域名列表中的任意一个，可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口查询。
	NewDefaultServerDomain *string `json:"NewDefaultServerDomain,omitnil,omitempty" name:"NewDefaultServerDomain"`
}

func (r *DeleteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "LocationIds")
	delete(f, "Domain")
	delete(f, "Url")
	delete(f, "NewDefaultServerDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRuleResponseParams `json:"Response"`
}

func (r *DeleteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTargetGroupsRequestParams struct {
	// 目标组的ID数组，单次最多支持删除20个。
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`
}

type DeleteTargetGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 目标组的ID数组，单次最多支持删除20个。
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`
}

func (r *DeleteTargetGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTargetGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTargetGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTargetGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTargetGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTargetGroupsResponseParams `json:"Response"`
}

func (r *DeleteTargetGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTargetGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserGroupsRequestParams struct {
	// <p>模型路由实例ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>待删除的用户组ID列表，单次1-100个。不可包含「未分组」虚拟分组 ugrp-ungrouped。组内仍有 Key 时将拒绝删除，需先将 Key 移出或迁移到其他组。</p>
	UserGroupIds []*string `json:"UserGroupIds,omitnil,omitempty" name:"UserGroupIds"`
}

type DeleteUserGroupsRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>待删除的用户组ID列表，单次1-100个。不可包含「未分组」虚拟分组 ugrp-ungrouped。组内仍有 Key 时将拒绝删除，需先将 Key 移出或迁移到其他组。</p>
	UserGroupIds []*string `json:"UserGroupIds,omitnil,omitempty" name:"UserGroupIds"`
}

func (r *DeleteUserGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "UserGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserGroupsResponseParams `json:"Response"`
}

func (r *DeleteUserGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeregisterFunctionTargetsRequestParams struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 待解绑的云函数列表。
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitnil,omitempty" name:"FunctionTargets"`

	// 目标转发规则的 ID，当将云函数从七层转发规则上解绑时，必须输入此参数或 Domain+Url 参数。
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 目标转发规则的域名，若已经输入 LocationId 参数，则本参数不生效。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 目标转发规则的 URL，若已经输入 LocationId 参数，则本参数不生效。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type DeregisterFunctionTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 待解绑的云函数列表。
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitnil,omitempty" name:"FunctionTargets"`

	// 目标转发规则的 ID，当将云函数从七层转发规则上解绑时，必须输入此参数或 Domain+Url 参数。
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 目标转发规则的域名，若已经输入 LocationId 参数，则本参数不生效。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 目标转发规则的 URL，若已经输入 LocationId 参数，则本参数不生效。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

func (r *DeregisterFunctionTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterFunctionTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "FunctionTargets")
	delete(f, "LocationId")
	delete(f, "Domain")
	delete(f, "Url")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeregisterFunctionTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeregisterFunctionTargetsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeregisterFunctionTargetsResponse struct {
	*tchttp.BaseResponse
	Response *DeregisterFunctionTargetsResponseParams `json:"Response"`
}

func (r *DeregisterFunctionTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterFunctionTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeregisterModelsFromServiceProviderRequestParams struct {
	// <p>BYOK的ID</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>模型别名列表</p>
	ModelAliases []*string `json:"ModelAliases,omitnil,omitempty" name:"ModelAliases"`
}

type DeregisterModelsFromServiceProviderRequest struct {
	*tchttp.BaseRequest
	
	// <p>BYOK的ID</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>模型别名列表</p>
	ModelAliases []*string `json:"ModelAliases,omitnil,omitempty" name:"ModelAliases"`
}

func (r *DeregisterModelsFromServiceProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterModelsFromServiceProviderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceProviderId")
	delete(f, "ModelAliases")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeregisterModelsFromServiceProviderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeregisterModelsFromServiceProviderResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeregisterModelsFromServiceProviderResponse struct {
	*tchttp.BaseResponse
	Response *DeregisterModelsFromServiceProviderResponseParams `json:"Response"`
}

func (r *DeregisterModelsFromServiceProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterModelsFromServiceProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeregisterTargetGroupInstancesRequestParams struct {
	// 目标组ID。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 待解绑的服务器信息，支持批量解除绑定，单次批量解除数量最多为20个。
	// 在这个接口 Port 参数为必填项。
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
}

type DeregisterTargetGroupInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 目标组ID。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 待解绑的服务器信息，支持批量解除绑定，单次批量解除数量最多为20个。
	// 在这个接口 Port 参数为必填项。
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
}

func (r *DeregisterTargetGroupInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterTargetGroupInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "TargetGroupInstances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeregisterTargetGroupInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeregisterTargetGroupInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeregisterTargetGroupInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DeregisterTargetGroupInstancesResponseParams `json:"Response"`
}

func (r *DeregisterTargetGroupInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterTargetGroupInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeregisterTargetsFromClassicalLBRequestParams struct {
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 后端服务的实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DeregisterTargetsFromClassicalLBRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 后端服务的实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DeregisterTargetsFromClassicalLBRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterTargetsFromClassicalLBRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeregisterTargetsFromClassicalLBRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeregisterTargetsFromClassicalLBResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeregisterTargetsFromClassicalLBResponse struct {
	*tchttp.BaseResponse
	Response *DeregisterTargetsFromClassicalLBResponseParams `json:"Response"`
}

func (r *DeregisterTargetsFromClassicalLBResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterTargetsFromClassicalLBResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeregisterTargetsRequestParams struct {
	// 负载均衡实例 ID，格式如 lb-12345678。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 监听器 ID，格式如 lbl-12345678。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 要解绑的后端服务列表，数组长度最大支持20。
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 转发规则的ID，格式如 loc-12345678，当从七层转发规则解绑机器时，必须提供此参数或Domain+URL两者之一。
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 目标规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 目标规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type DeregisterTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID，格式如 lb-12345678。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 监听器 ID，格式如 lbl-12345678。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 要解绑的后端服务列表，数组长度最大支持20。
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 转发规则的ID，格式如 loc-12345678，当从七层转发规则解绑机器时，必须提供此参数或Domain+URL两者之一。
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 目标规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 目标规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

func (r *DeregisterTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "Targets")
	delete(f, "LocationId")
	delete(f, "Domain")
	delete(f, "Url")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeregisterTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeregisterTargetsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeregisterTargetsResponse struct {
	*tchttp.BaseResponse
	Response *DeregisterTargetsResponseParams `json:"Response"`
}

func (r *DeregisterTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeregisterTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssociatedModelAvailabilityRequestParams struct {
	// <p>模型路由ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>模型列表</p>
	Models []*string `json:"Models,omitnil,omitempty" name:"Models"`
}

type DescribeAssociatedModelAvailabilityRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>模型列表</p>
	Models []*string `json:"Models,omitnil,omitempty" name:"Models"`
}

func (r *DescribeAssociatedModelAvailabilityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssociatedModelAvailabilityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "Models")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAssociatedModelAvailabilityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAssociatedModelAvailabilityResponseParams struct {
	// <p>模型可用性列表</p>
	ModelAvailability []*ModelAvailability `json:"ModelAvailability,omitnil,omitempty" name:"ModelAvailability"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAssociatedModelAvailabilityResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAssociatedModelAvailabilityResponseParams `json:"Response"`
}

func (r *DescribeAssociatedModelAvailabilityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAssociatedModelAvailabilityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsyncJobsRequestParams struct {
	// <p>请求ID列表</p>
	RequestIds []*string `json:"RequestIds,omitnil,omitempty" name:"RequestIds"`

	// <p>分页游标</p>
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// <p>本次查询最大数量</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`
}

type DescribeAsyncJobsRequest struct {
	*tchttp.BaseRequest
	
	// <p>请求ID列表</p>
	RequestIds []*string `json:"RequestIds,omitnil,omitempty" name:"RequestIds"`

	// <p>分页游标</p>
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// <p>本次查询最大数量</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`
}

func (r *DescribeAsyncJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RequestIds")
	delete(f, "NextToken")
	delete(f, "MaxResults")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAsyncJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsyncJobsResponseParams struct {
	// <p>异步任务列表</p>
	Jobs []*Job `json:"Jobs,omitnil,omitempty" name:"Jobs"`

	// <p>分页游标</p>
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// <p>本次查询最大数量</p>
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// <p>本次查询总数</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAsyncJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAsyncJobsResponseParams `json:"Response"`
}

func (r *DescribeAsyncJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAsyncJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockIPListRequestParams struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 数据偏移量，默认为 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回IP的最大个数，默认为 100000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeBlockIPListRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 数据偏移量，默认为 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回IP的最大个数，默认为 100000。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeBlockIPListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockIPListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlockIPListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockIPListResponseParams struct {
	// 返回的IP的数量
	BlockedIPCount *uint64 `json:"BlockedIPCount,omitnil,omitempty" name:"BlockedIPCount"`

	// 获取用户真实IP的字段
	ClientIPField *string `json:"ClientIPField,omitnil,omitempty" name:"ClientIPField"`

	// 加入了12360黑名单的IP列表
	BlockedIPList []*BlockedIP `json:"BlockedIPList,omitnil,omitempty" name:"BlockedIPList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBlockIPListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBlockIPListResponseParams `json:"Response"`
}

func (r *DescribeBlockIPListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockIPTaskRequestParams struct {
	// ModifyBlockIPList 接口返回的异步任务的ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeBlockIPTaskRequest struct {
	*tchttp.BaseRequest
	
	// ModifyBlockIPList 接口返回的异步任务的ID。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeBlockIPTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockIPTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBlockIPTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBlockIPTaskResponseParams struct {
	// 1 running，2 fail，6 succ
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBlockIPTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBlockIPTaskResponseParams `json:"Response"`
}

func (r *DescribeBlockIPTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBlockIPTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBudgetAssociationsRequestParams struct {
	// <p>Budget ID。</p><p>一次只允许查询一个Budget。</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>本次查询限制的数量</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>本次查询偏移量</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>资源类型。</p><p>枚举值：</p><ul><li>ModelRouter：模型路由实例</li><li>Key：模型路由Key</li></ul><p>不传时返回全部资源类型。</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeBudgetAssociationsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Budget ID。</p><p>一次只允许查询一个Budget。</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>本次查询限制的数量</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>本次查询偏移量</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>资源类型。</p><p>枚举值：</p><ul><li>ModelRouter：模型路由实例</li><li>Key：模型路由Key</li></ul><p>不传时返回全部资源类型。</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeBudgetAssociationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBudgetAssociationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BudgetId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBudgetAssociationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBudgetAssociationsResponseParams struct {
	// <p>Budget关联资源列表。</p>
	AssociationSet []*BudgetAssociation `json:"AssociationSet,omitnil,omitempty" name:"AssociationSet"`

	// <p>符合条件的总数。</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBudgetAssociationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBudgetAssociationsResponseParams `json:"Response"`
}

func (r *DescribeBudgetAssociationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBudgetAssociationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBudgetsRequestParams struct {
	// <p>Budget ID列表。</p>
	BudgetIds []*string `json:"BudgetIds,omitnil,omitempty" name:"BudgetIds"`

	// <p>过滤列表。</p><p>支持：BudgetId、BudgetName、Status。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>本次查询限制的数量。</p><p>取值范围：[1, 100]</p><p>默认值：20。</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>本次查询偏移量。</p><p>默认值：0。</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeBudgetsRequest struct {
	*tchttp.BaseRequest
	
	// <p>Budget ID列表。</p>
	BudgetIds []*string `json:"BudgetIds,omitnil,omitempty" name:"BudgetIds"`

	// <p>过滤列表。</p><p>支持：BudgetId、BudgetName、Status。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>本次查询限制的数量。</p><p>取值范围：[1, 100]</p><p>默认值：20。</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>本次查询偏移量。</p><p>默认值：0。</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeBudgetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBudgetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BudgetIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBudgetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBudgetsResponseParams struct {
	// <p>Budget列表。</p>
	BudgetSet []*BudgetInfo `json:"BudgetSet,omitnil,omitempty" name:"BudgetSet"`

	// <p>符合条件的总数。</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBudgetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBudgetsResponseParams `json:"Response"`
}

func (r *DescribeBudgetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBudgetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassicalLBByInstanceIdRequestParams struct {
	// 后端实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DescribeClassicalLBByInstanceIdRequest struct {
	*tchttp.BaseRequest
	
	// 后端实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DescribeClassicalLBByInstanceIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassicalLBByInstanceIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClassicalLBByInstanceIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassicalLBByInstanceIdResponseParams struct {
	// 负载均衡相关信息列表。
	LoadBalancerInfoList []*ClassicalLoadBalancerInfo `json:"LoadBalancerInfoList,omitnil,omitempty" name:"LoadBalancerInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClassicalLBByInstanceIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClassicalLBByInstanceIdResponseParams `json:"Response"`
}

func (r *DescribeClassicalLBByInstanceIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassicalLBByInstanceIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassicalLBHealthStatusRequestParams struct {
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`
}

type DescribeClassicalLBHealthStatusRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`
}

func (r *DescribeClassicalLBHealthStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassicalLBHealthStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClassicalLBHealthStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassicalLBHealthStatusResponseParams struct {
	// 后端健康状态列表。
	HealthList []*ClassicalHealth `json:"HealthList,omitnil,omitempty" name:"HealthList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClassicalLBHealthStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClassicalLBHealthStatusResponseParams `json:"Response"`
}

func (r *DescribeClassicalLBHealthStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassicalLBHealthStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassicalLBListenersRequestParams struct {
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID列表。
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// 负载均衡监听的协议：'TCP', 'UDP', 'HTTP', 'HTTPS'。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 负载均衡监听端口，范围为[1-65535]。
	ListenerPort *int64 `json:"ListenerPort,omitnil,omitempty" name:"ListenerPort"`

	// 监听器的状态，0：创建中，1：运行中。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type DescribeClassicalLBListenersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID列表。
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// 负载均衡监听的协议：'TCP', 'UDP', 'HTTP', 'HTTPS'。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 负载均衡监听端口，范围为[1-65535]。
	ListenerPort *int64 `json:"ListenerPort,omitnil,omitempty" name:"ListenerPort"`

	// 监听器的状态，0：创建中，1：运行中。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *DescribeClassicalLBListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassicalLBListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerIds")
	delete(f, "Protocol")
	delete(f, "ListenerPort")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClassicalLBListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassicalLBListenersResponseParams struct {
	// 监听器列表。
	Listeners []*ClassicalListener `json:"Listeners,omitnil,omitempty" name:"Listeners"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClassicalLBListenersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClassicalLBListenersResponseParams `json:"Response"`
}

func (r *DescribeClassicalLBListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassicalLBListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassicalLBTargetsRequestParams struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`
}

type DescribeClassicalLBTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`
}

func (r *DescribeClassicalLBTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassicalLBTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClassicalLBTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClassicalLBTargetsResponseParams struct {
	// 后端服务列表。
	Targets []*ClassicalTarget `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClassicalLBTargetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClassicalLBTargetsResponseParams `json:"Response"`
}

func (r *DescribeClassicalLBTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClassicalLBTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClsLogSetRequestParams struct {

}

type DescribeClsLogSetRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeClsLogSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClsLogSetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClsLogSetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClsLogSetResponseParams struct {
	// 日志集的 ID。
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// 健康检查日志集的 ID。
	HealthLogsetId *string `json:"HealthLogsetId,omitnil,omitempty" name:"HealthLogsetId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClsLogSetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClsLogSetResponseParams `json:"Response"`
}

func (r *DescribeClsLogSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClsLogSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterResourcesRequestParams struct {
	// 返回集群中资源列表数目，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回集群中资源列表起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询集群中资源列表条件，详细的过滤条件如下：
	// <li> cluster-id - String - 是否必填：否 - （过滤条件）按照 集群 的唯一ID过滤，如 ："tgw-12345678","stgw-12345678","vpcgw-12345678"。</li>
	// <li> vip - String - 是否必填：否 - （过滤条件）按照vip过滤。</li>
	// <li> loadbalancer-id - String - 是否必填：否 - （过滤条件）按照负载均衡唯一ID过滤。</li>
	// <li> idle - String 是否必填：否 - （过滤条件）按照是否闲置过滤，如"True","False"。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeClusterResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 返回集群中资源列表数目，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回集群中资源列表起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询集群中资源列表条件，详细的过滤条件如下：
	// <li> cluster-id - String - 是否必填：否 - （过滤条件）按照 集群 的唯一ID过滤，如 ："tgw-12345678","stgw-12345678","vpcgw-12345678"。</li>
	// <li> vip - String - 是否必填：否 - （过滤条件）按照vip过滤。</li>
	// <li> loadbalancer-id - String - 是否必填：否 - （过滤条件）按照负载均衡唯一ID过滤。</li>
	// <li> idle - String 是否必填：否 - （过滤条件）按照是否闲置过滤，如"True","False"。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeClusterResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterResourcesResponseParams struct {
	// 集群中资源列表。
	ClusterResourceSet []*ClusterResource `json:"ClusterResourceSet,omitnil,omitempty" name:"ClusterResourceSet"`

	// 集群中资源总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterResourcesResponseParams `json:"Response"`
}

func (r *DescribeClusterResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCrossTargetsRequestParams struct {
	// 返回后端服务列表数目，默认20，最大值100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回后端服务列表起始偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询跨域2.0版本云联网后端子机和网卡服务列表条件，详细的过滤条件如下：
	// <li> vpc-id - String - 是否必填：否 - （过滤条件）按照 本地私有网络ID，即负载均衡的VpcId 过滤，如："vpc-12345678"。</li>
	// <li> ip - String - 是否必填：否 - （过滤条件）按照 后端服务ip 过滤，如："192.168.0.1"。</li>
	// <li> listener-id - String - 是否必填：否 - （过滤条件）按照 监听器ID 过滤，如："lbl-12345678"。</li>
	// <li> location-id - String - 是否必填：否 - （过滤条件）按照 七层监听器规则ID 过滤，如："loc-12345678"。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeCrossTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 返回后端服务列表数目，默认20，最大值100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回后端服务列表起始偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询跨域2.0版本云联网后端子机和网卡服务列表条件，详细的过滤条件如下：
	// <li> vpc-id - String - 是否必填：否 - （过滤条件）按照 本地私有网络ID，即负载均衡的VpcId 过滤，如："vpc-12345678"。</li>
	// <li> ip - String - 是否必填：否 - （过滤条件）按照 后端服务ip 过滤，如："192.168.0.1"。</li>
	// <li> listener-id - String - 是否必填：否 - （过滤条件）按照 监听器ID 过滤，如："lbl-12345678"。</li>
	// <li> location-id - String - 是否必填：否 - （过滤条件）按照 七层监听器规则ID 过滤，如："loc-12345678"。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeCrossTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCrossTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCrossTargetsResponseParams struct {
	// 后端服务列表总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 后端服务列表。
	CrossTargetSet []*CrossTargets `json:"CrossTargetSet,omitnil,omitempty" name:"CrossTargetSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCrossTargetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCrossTargetsResponseParams `json:"Response"`
}

func (r *DescribeCrossTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCrossTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomizedConfigAssociateListRequestParams struct {
	// 配置ID，可以通过 [DescribeCustomizedConfigList](https://cloud.tencent.com/document/product/214/60009) 接口获取。
	UconfigId *string `json:"UconfigId,omitnil,omitempty" name:"UconfigId"`

	// 拉取绑定关系列表开始位置，默认值 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 拉取绑定关系列表数目，默认值 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索域名，可以通过 [DescribeLoadBalancersDetail](https://cloud.tencent.com/document/product/214/46916) 接口返回值的 `Domain` 字段查询。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type DescribeCustomizedConfigAssociateListRequest struct {
	*tchttp.BaseRequest
	
	// 配置ID，可以通过 [DescribeCustomizedConfigList](https://cloud.tencent.com/document/product/214/60009) 接口获取。
	UconfigId *string `json:"UconfigId,omitnil,omitempty" name:"UconfigId"`

	// 拉取绑定关系列表开始位置，默认值 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 拉取绑定关系列表数目，默认值 20
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 搜索域名，可以通过 [DescribeLoadBalancersDetail](https://cloud.tencent.com/document/product/214/46916) 接口返回值的 `Domain` 字段查询。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

func (r *DescribeCustomizedConfigAssociateListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizedConfigAssociateListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UconfigId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomizedConfigAssociateListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomizedConfigAssociateListResponseParams struct {
	// 绑定关系列表
	BindList []*BindDetailItem `json:"BindList,omitnil,omitempty" name:"BindList"`

	// 绑定关系总数目
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomizedConfigAssociateListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomizedConfigAssociateListResponseParams `json:"Response"`
}

func (r *DescribeCustomizedConfigAssociateListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizedConfigAssociateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomizedConfigListRequestParams struct {
	// 配置类型:CLB 负载均衡维度。 SERVER 域名维度。 LOCATION 规则维度。
	ConfigType *string `json:"ConfigType,omitnil,omitempty" name:"ConfigType"`

	// 拉取页偏移，默认值0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 拉取数目，默认值20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 拉取指定配置名字，模糊匹配。
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// 配置ID，可以通过 [DescribeCustomizedConfigList](https://cloud.tencent.com/document/api/214/60009) 接口查询。
	UconfigIds []*string `json:"UconfigIds,omitnil,omitempty" name:"UconfigIds"`

	// 过滤条件如下：
	// - loadbalancer-id
	// 按照【负载均衡 ID】进行过滤。实例计费模式例如：lb-9vxezxza。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459)
	// - vip
	// 按照【负载均衡VIP】进行过滤。网络计费模式例如："1.1.1.1","2204::22:3"。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459)
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeCustomizedConfigListRequest struct {
	*tchttp.BaseRequest
	
	// 配置类型:CLB 负载均衡维度。 SERVER 域名维度。 LOCATION 规则维度。
	ConfigType *string `json:"ConfigType,omitnil,omitempty" name:"ConfigType"`

	// 拉取页偏移，默认值0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 拉取数目，默认值20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 拉取指定配置名字，模糊匹配。
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// 配置ID，可以通过 [DescribeCustomizedConfigList](https://cloud.tencent.com/document/api/214/60009) 接口查询。
	UconfigIds []*string `json:"UconfigIds,omitnil,omitempty" name:"UconfigIds"`

	// 过滤条件如下：
	// - loadbalancer-id
	// 按照【负载均衡 ID】进行过滤。实例计费模式例如：lb-9vxezxza。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459)
	// - vip
	// 按照【负载均衡VIP】进行过滤。网络计费模式例如："1.1.1.1","2204::22:3"。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459)
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeCustomizedConfigListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizedConfigListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConfigType")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ConfigName")
	delete(f, "UconfigIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCustomizedConfigListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCustomizedConfigListResponseParams struct {
	// 配置列表
	ConfigList []*ConfigListItem `json:"ConfigList,omitnil,omitempty" name:"ConfigList"`

	// 配置数目
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCustomizedConfigListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCustomizedConfigListResponseParams `json:"Response"`
}

func (r *DescribeCustomizedConfigListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCustomizedConfigListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExclusiveClustersRequestParams struct {
	// 返回集群列表数目，默认值为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回集群列表起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询集群列表条件，详细的过滤条件如下：
	// <li> cluster-type - String - 是否必填：否 - （过滤条件）按照 集群 的类型过滤，包括"TGW","STGW","VPCGW"。</li>
	// <li> cluster-id - String - 是否必填：否 - （过滤条件）按照 集群 的唯一ID过滤，如 ："tgw-12345678","stgw-12345678","vpcgw-12345678"。</li>
	// <li> cluster-name - String - 是否必填：否 - （过滤条件）按照 集群 的名称过滤。</li>
	// <li> cluster-tag - String - 是否必填：否 - （过滤条件）按照 集群 的标签过滤。（只有TGW/STGW集群有集群标签） </li>
	// <li> vip - String - 是否必填：否 - （过滤条件）按照 集群 内的vip过滤。</li>
	// <li> loadbalancer-id - String - 是否必填：否 - （过滤条件）按照 集群 内的负载均衡唯一ID过滤。</li>
	// <li> network - String - 是否必填：否 - （过滤条件）按照 集群 的网络类型过滤，如："Public","Private"。</li>
	// <li> zone - String - 是否必填：否 - （过滤条件）按照 集群 所在可用区过滤，如："ap-guangzhou-1"（广州一区）。</li>
	// <li> isp -- String - 是否必填：否 - （过滤条件）按照TGW集群的 Isp 类型过滤，如："BGP","CMCC","CUCC","CTCC","INTERNAL"。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeExclusiveClustersRequest struct {
	*tchttp.BaseRequest
	
	// 返回集群列表数目，默认值为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回集群列表起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询集群列表条件，详细的过滤条件如下：
	// <li> cluster-type - String - 是否必填：否 - （过滤条件）按照 集群 的类型过滤，包括"TGW","STGW","VPCGW"。</li>
	// <li> cluster-id - String - 是否必填：否 - （过滤条件）按照 集群 的唯一ID过滤，如 ："tgw-12345678","stgw-12345678","vpcgw-12345678"。</li>
	// <li> cluster-name - String - 是否必填：否 - （过滤条件）按照 集群 的名称过滤。</li>
	// <li> cluster-tag - String - 是否必填：否 - （过滤条件）按照 集群 的标签过滤。（只有TGW/STGW集群有集群标签） </li>
	// <li> vip - String - 是否必填：否 - （过滤条件）按照 集群 内的vip过滤。</li>
	// <li> loadbalancer-id - String - 是否必填：否 - （过滤条件）按照 集群 内的负载均衡唯一ID过滤。</li>
	// <li> network - String - 是否必填：否 - （过滤条件）按照 集群 的网络类型过滤，如："Public","Private"。</li>
	// <li> zone - String - 是否必填：否 - （过滤条件）按照 集群 所在可用区过滤，如："ap-guangzhou-1"（广州一区）。</li>
	// <li> isp -- String - 是否必填：否 - （过滤条件）按照TGW集群的 Isp 类型过滤，如："BGP","CMCC","CUCC","CTCC","INTERNAL"。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeExclusiveClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExclusiveClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExclusiveClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExclusiveClustersResponseParams struct {
	// 集群列表。
	ClusterSet []*Cluster `json:"ClusterSet,omitnil,omitempty" name:"ClusterSet"`

	// 集群总数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExclusiveClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExclusiveClustersResponseParams `json:"Response"`
}

func (r *DescribeExclusiveClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExclusiveClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdleLoadBalancersRequestParams struct {
	// 数据偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回负载均衡实例的数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 负载均衡所在地域，可以通过 [DescribeRegions](https://cloud.tencent.com/document/product/1596/77930) 接口返回值 `RegionSet.Region` 字段获取。
	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitnil,omitempty" name:"LoadBalancerRegion"`
}

type DescribeIdleLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// 数据偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回负载均衡实例的数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 负载均衡所在地域，可以通过 [DescribeRegions](https://cloud.tencent.com/document/product/1596/77930) 接口返回值 `RegionSet.Region` 字段获取。
	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitnil,omitempty" name:"LoadBalancerRegion"`
}

func (r *DescribeIdleLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdleLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "LoadBalancerRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIdleLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIdleLoadBalancersResponseParams struct {
	// 闲置实例列表
	IdleLoadBalancers []*IdleLoadBalancer `json:"IdleLoadBalancers,omitnil,omitempty" name:"IdleLoadBalancers"`

	// 所有闲置实例数目
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIdleLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIdleLoadBalancersResponseParams `json:"Response"`
}

func (r *DescribeIdleLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIdleLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntentRouterTiersRequestParams struct {

}

type DescribeIntentRouterTiersRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeIntentRouterTiersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntentRouterTiersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntentRouterTiersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntentRouterTiersResponseParams struct {
	// <p>Tier 字典列表（按 tier_id 升序排列）</p>
	TierSet []*IntentRouterTierDictItem `json:"TierSet,omitnil,omitempty" name:"TierSet"`

	// <p>Tier 总条数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIntentRouterTiersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntentRouterTiersResponseParams `json:"Response"`
}

func (r *DescribeIntentRouterTiersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntentRouterTiersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntentRoutersRequestParams struct {
	// <p>模型路由实例ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>意图路由ID列表</p>
	IntentRouterIds []*string `json:"IntentRouterIds,omitnil,omitempty" name:"IntentRouterIds"`
}

type DescribeIntentRoutersRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>意图路由ID列表</p>
	IntentRouterIds []*string `json:"IntentRouterIds,omitnil,omitempty" name:"IntentRouterIds"`
}

func (r *DescribeIntentRoutersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntentRoutersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "IntentRouterIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntentRoutersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntentRoutersResponseParams struct {
	// <p>意图路由列表。</p>
	IntentRouterSet []*IntentRouterItem `json:"IntentRouterSet,omitnil,omitempty" name:"IntentRouterSet"`

	// <p>意图路由总数。</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIntentRoutersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntentRoutersResponseParams `json:"Response"`
}

func (r *DescribeIntentRoutersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntentRoutersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeysRequestParams struct {
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>过滤列表</p><p>支持：KeyName、BudgetId、tag-key、tag:&lt;tag-key&gt;。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>API Key的ID列表</p>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// <p>本次查询限制的数量</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>本次查询偏移量</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeKeysRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>过滤列表</p><p>支持：KeyName、BudgetId、tag-key、tag:&lt;tag-key&gt;。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>API Key的ID列表</p>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// <p>本次查询限制的数量</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>本次查询偏移量</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "Filters")
	delete(f, "KeyIds")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKeysResponseParams struct {
	// <p>API Key列表</p>
	Keys []*KeyInfo `json:"Keys,omitnil,omitempty" name:"Keys"`

	// <p>符合条件的总数</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKeysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKeysResponseParams `json:"Response"`
}

func (r *DescribeKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLBListenersRequestParams struct {
	// 需要查询的内网ip列表
	Backends []*LbRsItem `json:"Backends,omitnil,omitempty" name:"Backends"`
}

type DescribeLBListenersRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的内网ip列表
	Backends []*LbRsItem `json:"Backends,omitnil,omitempty" name:"Backends"`
}

func (r *DescribeLBListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLBListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Backends")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLBListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLBListenersResponseParams struct {
	// 绑定的后端规则
	LoadBalancers []*LBItem `json:"LoadBalancers,omitnil,omitempty" name:"LoadBalancers"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLBListenersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLBListenersResponseParams `json:"Response"`
}

func (r *DescribeLBListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLBListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLBOperateProtectRequestParams struct {
	// 负载均衡实例ID。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`
}

type DescribeLBOperateProtectRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`
}

func (r *DescribeLBOperateProtectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLBOperateProtectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLBOperateProtectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLBOperateProtectResponseParams struct {
	// 返回的负载均衡操作保护信息数组。
	LoadBalancerSet []*LBOperateProtectInfo `json:"LoadBalancerSet,omitnil,omitempty" name:"LoadBalancerSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLBOperateProtectResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLBOperateProtectResponseParams `json:"Response"`
}

func (r *DescribeLBOperateProtectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLBOperateProtectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenersRequestParams struct {
	// 负载均衡实例 ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/api/214/30685) 接口获取。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 要查询的负载均衡监听器 ID 数组，最大为100个，可以通过 [DescribeListeners](https://cloud.tencent.com/document/api/214/30686) 接口获取。
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// 要查询的监听器协议类型，取值 TCP | UDP | HTTP | HTTPS | TCP_SSL | QUIC。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 要查询的监听器的端口，端口范围：1-65535
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type DescribeListenersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/api/214/30685) 接口获取。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 要查询的负载均衡监听器 ID 数组，最大为100个，可以通过 [DescribeListeners](https://cloud.tencent.com/document/api/214/30686) 接口获取。
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// 要查询的监听器协议类型，取值 TCP | UDP | HTTP | HTTPS | TCP_SSL | QUIC。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 要查询的监听器的端口，端口范围：1-65535
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

func (r *DescribeListenersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerIds")
	delete(f, "Protocol")
	delete(f, "Port")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenersResponseParams struct {
	// 监听器列表。
	Listeners []*Listener `json:"Listeners,omitnil,omitempty" name:"Listeners"`

	// 总的监听器个数（根据端口、协议、监听器ID过滤后）。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeListenersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListenersResponseParams `json:"Response"`
}

func (r *DescribeListenersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancerListByCertIdRequestParams struct {
	// 服务端证书的ID，或客户端证书的ID。可以通过 [DescribeCertificate](https://cloud.tencent.com/document/api/400/41674) 接口查询。
	// 数组最大长度为20。
	CertIds []*string `json:"CertIds,omitnil,omitempty" name:"CertIds"`
}

type DescribeLoadBalancerListByCertIdRequest struct {
	*tchttp.BaseRequest
	
	// 服务端证书的ID，或客户端证书的ID。可以通过 [DescribeCertificate](https://cloud.tencent.com/document/api/400/41674) 接口查询。
	// 数组最大长度为20。
	CertIds []*string `json:"CertIds,omitnil,omitempty" name:"CertIds"`
}

func (r *DescribeLoadBalancerListByCertIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerListByCertIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancerListByCertIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancerListByCertIdResponseParams struct {
	// 证书ID，以及与该证书ID关联的负载均衡实例列表
	CertSet []*CertIdRelatedWithLoadBalancers `json:"CertSet,omitnil,omitempty" name:"CertSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLoadBalancerListByCertIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancerListByCertIdResponseParams `json:"Response"`
}

func (r *DescribeLoadBalancerListByCertIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerListByCertIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancerOverviewRequestParams struct {

}

type DescribeLoadBalancerOverviewRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLoadBalancerOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancerOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancerOverviewResponseParams struct {
	// 负载均衡总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 运行中的负载均衡数目
	RunningCount *int64 `json:"RunningCount,omitnil,omitempty" name:"RunningCount"`

	// 隔离中的负载均衡数目
	IsolationCount *int64 `json:"IsolationCount,omitnil,omitempty" name:"IsolationCount"`

	// 即将到期的负载均衡数目
	WillExpireCount *int64 `json:"WillExpireCount,omitnil,omitempty" name:"WillExpireCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLoadBalancerOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancerOverviewResponseParams `json:"Response"`
}

func (r *DescribeLoadBalancerOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancerTrafficRequestParams struct {
	// 负载均衡所在地域，不传默认返回所有地域负载均衡。
	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitnil,omitempty" name:"LoadBalancerRegion"`
}

type DescribeLoadBalancerTrafficRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡所在地域，不传默认返回所有地域负载均衡。
	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitnil,omitempty" name:"LoadBalancerRegion"`
}

func (r *DescribeLoadBalancerTrafficRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerTrafficRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancerTrafficRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancerTrafficResponseParams struct {
	// 按出带宽从高到低排序后的负载均衡信息。
	LoadBalancerTraffic []*LoadBalancerTraffic `json:"LoadBalancerTraffic,omitnil,omitempty" name:"LoadBalancerTraffic"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLoadBalancerTrafficResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancerTrafficResponseParams `json:"Response"`
}

func (r *DescribeLoadBalancerTrafficResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerTrafficResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancersDetailRequestParams struct {
	// 返回负载均衡列表数目，默认20，最大值100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回负载均衡列表起始偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 选择返回的Fields列表，系统仅会返回Fileds中填写的字段，可填写的字段详情请参见<a href="https://cloud.tencent.com/document/api/214/30694#LoadBalancerDetail">LoadBalancerDetail</a>。若未在Fileds填写相关字段，则此字段返回null。Fileds中默认添加LoadBalancerId和LoadBalancerName字段。
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 当Fields包含TargetId、TargetAddress、TargetPort、TargetWeight、ListenerId、Protocol、Port、LocationId、Domain、Url等Fields时，必选选择导出目标组的Target或者非目标组Target，取值范围NODE、GROUP。
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 查询负载均衡详细信息列表条件，详细的过滤条件如下：
	// - loadbalancer-id
	// 按照【负载均衡ID】进行过滤。例如：lb-rbw5skde。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459)
	// - project-id
	// 按照【项目ID】进行过滤。例如： "0"、"123"。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeProject](https://cloud.tencent.com/document/api/651/78725)
	// - network
	// 按照【负载均衡网络类型】进行过滤。例如：Public。
	// 类型：String
	// 必选：否
	// 可选值：Private（内网）、Public（公网）
	// - vip
	// 按照【负载均衡 VIP】进行过滤。例如："1.1.1.1","2204::22:3"。
	// 类型：String
	// 必选：否
	// - vpcid
	// 按照【负载均衡所属 VPCID】进行过滤。例如："vpc-12345678"。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeZones](https://cloud.tencent.com/document/product/213/15707)
	// - target-ip
	// 按照【后端目标内网 IP】进行过滤。例如："1.1.1.1","2203::214:4"。
	// 类型：String
	// 必选：否
	// - zone
	// 按照【负载均衡所属的可用区】进行过滤。例如："ap-guangzhou-1"。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeZones](https://cloud.tencent.com/document/product/213/15707)
	// - tag-key
	// 按照【负载均衡标签的标签键】进行过滤，例如："name"。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeTags](https://cloud.tencent.com/document/api/651/35316)
	// - tag:*
	// 按照【负载均衡的标签】进行过滤，':' 后面跟的是标签键。如过滤标签键name，标签值zhangsan,lisi，{"Name": "tag:name","Values": ["zhangsan", "lisi"]}。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeTagKeys](https://cloud.tencent.com/document/api/651/35318)
	// - fuzzy-search
	// 按照【负载均衡VIP，负载均衡名称】模糊搜索，例如："1.1"。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeLoadBalancersDetailRequest struct {
	*tchttp.BaseRequest
	
	// 返回负载均衡列表数目，默认20，最大值100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回负载均衡列表起始偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 选择返回的Fields列表，系统仅会返回Fileds中填写的字段，可填写的字段详情请参见<a href="https://cloud.tencent.com/document/api/214/30694#LoadBalancerDetail">LoadBalancerDetail</a>。若未在Fileds填写相关字段，则此字段返回null。Fileds中默认添加LoadBalancerId和LoadBalancerName字段。
	Fields []*string `json:"Fields,omitnil,omitempty" name:"Fields"`

	// 当Fields包含TargetId、TargetAddress、TargetPort、TargetWeight、ListenerId、Protocol、Port、LocationId、Domain、Url等Fields时，必选选择导出目标组的Target或者非目标组Target，取值范围NODE、GROUP。
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 查询负载均衡详细信息列表条件，详细的过滤条件如下：
	// - loadbalancer-id
	// 按照【负载均衡ID】进行过滤。例如：lb-rbw5skde。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459)
	// - project-id
	// 按照【项目ID】进行过滤。例如： "0"、"123"。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeProject](https://cloud.tencent.com/document/api/651/78725)
	// - network
	// 按照【负载均衡网络类型】进行过滤。例如：Public。
	// 类型：String
	// 必选：否
	// 可选值：Private（内网）、Public（公网）
	// - vip
	// 按照【负载均衡 VIP】进行过滤。例如："1.1.1.1","2204::22:3"。
	// 类型：String
	// 必选：否
	// - vpcid
	// 按照【负载均衡所属 VPCID】进行过滤。例如："vpc-12345678"。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeZones](https://cloud.tencent.com/document/product/213/15707)
	// - target-ip
	// 按照【后端目标内网 IP】进行过滤。例如："1.1.1.1","2203::214:4"。
	// 类型：String
	// 必选：否
	// - zone
	// 按照【负载均衡所属的可用区】进行过滤。例如："ap-guangzhou-1"。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeZones](https://cloud.tencent.com/document/product/213/15707)
	// - tag-key
	// 按照【负载均衡标签的标签键】进行过滤，例如："name"。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeTags](https://cloud.tencent.com/document/api/651/35316)
	// - tag:*
	// 按照【负载均衡的标签】进行过滤，':' 后面跟的是标签键。如过滤标签键name，标签值zhangsan,lisi，{"Name": "tag:name","Values": ["zhangsan", "lisi"]}。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeTagKeys](https://cloud.tencent.com/document/api/651/35318)
	// - fuzzy-search
	// 按照【负载均衡VIP，负载均衡名称】模糊搜索，例如："1.1"。
	// 类型：String
	// 必选：否
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeLoadBalancersDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancersDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Fields")
	delete(f, "TargetType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancersDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancersDetailResponseParams struct {
	// 负载均衡详情列表总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 负载均衡详情列表。
	LoadBalancerDetailSet []*LoadBalancerDetail `json:"LoadBalancerDetailSet,omitnil,omitempty" name:"LoadBalancerDetailSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLoadBalancersDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancersDetailResponseParams `json:"Response"`
}

func (r *DescribeLoadBalancersDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancersDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancersRequestParams struct {
	// 负载均衡实例ID。实例ID数量上限为20个。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 负载均衡实例的网络类型：
	// OPEN：公网属性， INTERNAL：内网属性。
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// 负载均衡实例的类型。1：通用的负载均衡实例，0：传统型负载均衡实例。如果不传此参数，则查询所有类型的负载均衡实例。
	Forward *int64 `json:"Forward,omitnil,omitempty" name:"Forward"`

	// 负载均衡实例的名称，支持模糊查询。
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 腾讯云为负载均衡实例分配的域名，支持模糊查询。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 负载均衡实例的 VIP 地址，支持多个。
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitnil,omitempty" name:"LoadBalancerVips"`

	// 负载均衡绑定的后端服务的外网 IP，只支持查询云服务器的公网 IP。
	BackendPublicIps []*string `json:"BackendPublicIps,omitnil,omitempty" name:"BackendPublicIps"`

	// 负载均衡绑定的后端服务的内网 IP，只支持查询云服务器的内网 IP。
	BackendPrivateIps []*string `json:"BackendPrivateIps,omitnil,omitempty" name:"BackendPrivateIps"`

	// 数据偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回负载均衡实例的数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序参数，支持以下字段：
	// - LoadBalancerName
	// - CreateTime
	// - Domain
	// - LoadBalancerType
	// 
	// 默认为 CreateTime。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 1：倒序，0：顺序，默认为1，按照创建时间倒序。
	OrderType *int64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 模糊搜索字段，模糊匹配负载均衡实例的名称、域名、负载均衡实例的 VIP 地址，负载均衡实例ID。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 负载均衡实例所属的项目 ID，可以通过[DescribeProject](https://cloud.tencent.com/document/api/651/78725)接口获取，不传默认所有项目。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 负载均衡是否绑定后端服务，0：没有绑定后端服务，1：绑定后端服务，-1：查询全部。
	WithRs *int64 `json:"WithRs,omitnil,omitempty" name:"WithRs"`

	// 负载均衡实例所属私有网络唯一ID，如 vpc-bhqkbhdx，可以通过[DescribeVpcs](https://cloud.tencent.com/document/api/215/15778)接口获取。
	// 查找基础网络类型的负载均衡可传入'0'。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 安全组ID，如 sg-m1cc****，可以通过接口[DescribeSecurityGroups](https://cloud.tencent.com/document/product/215/15808)获取。
	SecurityGroup *string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 主可用区ID，如 ："100001" （对应的是广州一区）。可通过[DescribeZones](https://cloud.tencent.com/document/product/213/15707)获取可用区列表。
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为100。<br/>`Filter.Name`和`Filter.Values`皆为必填项。详细的过滤条件如下：
	// - charge-type
	// 按照【实例计费模式】进行过滤。实例计费模式例如：PREPAID。
	// 类型：String
	// 必选：否
	// 可选项：PREPAID(预付费)、POSTPAID_BY_HOUR(后付费)
	// - internet-charge-type
	// 按照【网络计费模式】进行过滤。网络计费模式例如：BANDWIDTH_PREPAID。
	// 类型：String
	// 必选：否
	// 可选项：BANDWIDTH_PREPAID(预付费按带宽结算)、 TRAFFIC_POSTPAID_BY_HOUR(流量按小时后付费)、BANDWIDTH_POSTPAID_BY_HOUR(带宽按小时后付费)、BANDWIDTH_PACKAGE(带宽包用户)
	// - master-zone-id
	// 按照【CLB主可用区ID】进行过滤。例如：100001（对应的是广州一区）。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeZones](https://cloud.tencent.com/document/product/213/15707)
	// - tag-key
	// 按照【CLB 标签的键】进行过滤，例如：tag-key。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeTags](https://cloud.tencent.com/document/api/651/35316)
	// - tag:tag-key
	// 按照【CLB标签键值】进行过滤，例如：tag-test。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeTagKeys](https://cloud.tencent.com/document/api/651/35318)
	// - function-name
	// 按照【后端绑定SCF云函数的函数名称】进行过滤，例如：helloworld-1744958255。
	// 类型：String
	// 必选：否
	// 获取方式：[ListFunctions](https://cloud.tencent.com/document/api/583/18582)
	// - vip-isp
	// 按照【CLB VIP的运营商类型】进行过滤，例如：BGP。
	// 类型：String
	// 必选：否
	// 公网类型可选项：BGP(多线)、CMCC(中国移动)、CTCC(中国电信)、CUCC(中国联通)
	// 内网类型可选项：INTERNAL(内网)
	// - sla-type
	// 按照【CLB 的性能容量型规格】进行过滤，例如：clb.c4.xlarge。
	// 类型：String
	// 必选：否
	// 可选项：clb.c2.medium(标准型)、clb.c3.small(高阶型1)、clb.c3.medium(高阶型2)、clb.c4.small(超强型1)、clb.c4.medium(超强型2)、clb.c4.large(超强型3)、clb.c4.xlarge(超强型4)
	// 具体规格参数参考：
	// - exclusive
	// 按照【独占实例】进行过滤。例如：1，代表筛选独占型实例。
	// 类型：String
	// 必选：否
	// 可选项：0、1
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 选择返回的扩充字段，不指定时，扩充字段默认不返回。详细支持的扩充字段如下：
	// <li> TargetCount：绑定的后端服务数量</li>
	AdditionalFields []*string `json:"AdditionalFields,omitnil,omitempty" name:"AdditionalFields"`
}

type DescribeLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。实例ID数量上限为20个。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 负载均衡实例的网络类型：
	// OPEN：公网属性， INTERNAL：内网属性。
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// 负载均衡实例的类型。1：通用的负载均衡实例，0：传统型负载均衡实例。如果不传此参数，则查询所有类型的负载均衡实例。
	Forward *int64 `json:"Forward,omitnil,omitempty" name:"Forward"`

	// 负载均衡实例的名称，支持模糊查询。
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 腾讯云为负载均衡实例分配的域名，支持模糊查询。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 负载均衡实例的 VIP 地址，支持多个。
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitnil,omitempty" name:"LoadBalancerVips"`

	// 负载均衡绑定的后端服务的外网 IP，只支持查询云服务器的公网 IP。
	BackendPublicIps []*string `json:"BackendPublicIps,omitnil,omitempty" name:"BackendPublicIps"`

	// 负载均衡绑定的后端服务的内网 IP，只支持查询云服务器的内网 IP。
	BackendPrivateIps []*string `json:"BackendPrivateIps,omitnil,omitempty" name:"BackendPrivateIps"`

	// 数据偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回负载均衡实例的数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序参数，支持以下字段：
	// - LoadBalancerName
	// - CreateTime
	// - Domain
	// - LoadBalancerType
	// 
	// 默认为 CreateTime。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 1：倒序，0：顺序，默认为1，按照创建时间倒序。
	OrderType *int64 `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 模糊搜索字段，模糊匹配负载均衡实例的名称、域名、负载均衡实例的 VIP 地址，负载均衡实例ID。
	SearchKey *string `json:"SearchKey,omitnil,omitempty" name:"SearchKey"`

	// 负载均衡实例所属的项目 ID，可以通过[DescribeProject](https://cloud.tencent.com/document/api/651/78725)接口获取，不传默认所有项目。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 负载均衡是否绑定后端服务，0：没有绑定后端服务，1：绑定后端服务，-1：查询全部。
	WithRs *int64 `json:"WithRs,omitnil,omitempty" name:"WithRs"`

	// 负载均衡实例所属私有网络唯一ID，如 vpc-bhqkbhdx，可以通过[DescribeVpcs](https://cloud.tencent.com/document/api/215/15778)接口获取。
	// 查找基础网络类型的负载均衡可传入'0'。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 安全组ID，如 sg-m1cc****，可以通过接口[DescribeSecurityGroups](https://cloud.tencent.com/document/product/215/15808)获取。
	SecurityGroup *string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 主可用区ID，如 ："100001" （对应的是广州一区）。可通过[DescribeZones](https://cloud.tencent.com/document/product/213/15707)获取可用区列表。
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为100。<br/>`Filter.Name`和`Filter.Values`皆为必填项。详细的过滤条件如下：
	// - charge-type
	// 按照【实例计费模式】进行过滤。实例计费模式例如：PREPAID。
	// 类型：String
	// 必选：否
	// 可选项：PREPAID(预付费)、POSTPAID_BY_HOUR(后付费)
	// - internet-charge-type
	// 按照【网络计费模式】进行过滤。网络计费模式例如：BANDWIDTH_PREPAID。
	// 类型：String
	// 必选：否
	// 可选项：BANDWIDTH_PREPAID(预付费按带宽结算)、 TRAFFIC_POSTPAID_BY_HOUR(流量按小时后付费)、BANDWIDTH_POSTPAID_BY_HOUR(带宽按小时后付费)、BANDWIDTH_PACKAGE(带宽包用户)
	// - master-zone-id
	// 按照【CLB主可用区ID】进行过滤。例如：100001（对应的是广州一区）。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeZones](https://cloud.tencent.com/document/product/213/15707)
	// - tag-key
	// 按照【CLB 标签的键】进行过滤，例如：tag-key。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeTags](https://cloud.tencent.com/document/api/651/35316)
	// - tag:tag-key
	// 按照【CLB标签键值】进行过滤，例如：tag-test。
	// 类型：String
	// 必选：否
	// 获取方式：[DescribeTagKeys](https://cloud.tencent.com/document/api/651/35318)
	// - function-name
	// 按照【后端绑定SCF云函数的函数名称】进行过滤，例如：helloworld-1744958255。
	// 类型：String
	// 必选：否
	// 获取方式：[ListFunctions](https://cloud.tencent.com/document/api/583/18582)
	// - vip-isp
	// 按照【CLB VIP的运营商类型】进行过滤，例如：BGP。
	// 类型：String
	// 必选：否
	// 公网类型可选项：BGP(多线)、CMCC(中国移动)、CTCC(中国电信)、CUCC(中国联通)
	// 内网类型可选项：INTERNAL(内网)
	// - sla-type
	// 按照【CLB 的性能容量型规格】进行过滤，例如：clb.c4.xlarge。
	// 类型：String
	// 必选：否
	// 可选项：clb.c2.medium(标准型)、clb.c3.small(高阶型1)、clb.c3.medium(高阶型2)、clb.c4.small(超强型1)、clb.c4.medium(超强型2)、clb.c4.large(超强型3)、clb.c4.xlarge(超强型4)
	// 具体规格参数参考：
	// - exclusive
	// 按照【独占实例】进行过滤。例如：1，代表筛选独占型实例。
	// 类型：String
	// 必选：否
	// 可选项：0、1
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 选择返回的扩充字段，不指定时，扩充字段默认不返回。详细支持的扩充字段如下：
	// <li> TargetCount：绑定的后端服务数量</li>
	AdditionalFields []*string `json:"AdditionalFields,omitnil,omitempty" name:"AdditionalFields"`
}

func (r *DescribeLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	delete(f, "LoadBalancerType")
	delete(f, "Forward")
	delete(f, "LoadBalancerName")
	delete(f, "Domain")
	delete(f, "LoadBalancerVips")
	delete(f, "BackendPublicIps")
	delete(f, "BackendPrivateIps")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "OrderType")
	delete(f, "SearchKey")
	delete(f, "ProjectId")
	delete(f, "WithRs")
	delete(f, "VpcId")
	delete(f, "SecurityGroup")
	delete(f, "MasterZone")
	delete(f, "Filters")
	delete(f, "AdditionalFields")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancersResponseParams struct {
	// 满足过滤条件的负载均衡实例总数。此数值与入参中的Limit无关。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 返回的负载均衡实例数组。
	LoadBalancerSet []*LoadBalancer `json:"LoadBalancerSet,omitnil,omitempty" name:"LoadBalancerSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancersResponseParams `json:"Response"`
}

func (r *DescribeLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelAliasesRequestParams struct {
	// <p>过滤条件</p><p>支持的过滤键：</p><ul><li>ModelAliasName：按模型别名过滤。</li></ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>每页数量，取值范围：[1, 100]，默认值：20。</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>分页偏移量，默认值：0。</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>排序条件。支持按 InputCoefficient、InputCachedCoefficient 或 OutputCoefficient 排序，Order 支持 ASC、DESC。不传或传空数组时，默认按 OutputCoefficient 降序排列。最多支持 3 个排序条件，排序字段不可重复。</p>
	Sort []*Sort `json:"Sort,omitnil,omitempty" name:"Sort"`
}

type DescribeModelAliasesRequest struct {
	*tchttp.BaseRequest
	
	// <p>过滤条件</p><p>支持的过滤键：</p><ul><li>ModelAliasName：按模型别名过滤。</li></ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>每页数量，取值范围：[1, 100]，默认值：20。</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>分页偏移量，默认值：0。</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>排序条件。支持按 InputCoefficient、InputCachedCoefficient 或 OutputCoefficient 排序，Order 支持 ASC、DESC。不传或传空数组时，默认按 OutputCoefficient 降序排列。最多支持 3 个排序条件，排序字段不可重复。</p>
	Sort []*Sort `json:"Sort,omitnil,omitempty" name:"Sort"`
}

func (r *DescribeModelAliasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelAliasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelAliasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelAliasesResponseParams struct {
	// <p>模型别名列表。</p>
	ModelAliasSet []*ModelAlias `json:"ModelAliasSet,omitnil,omitempty" name:"ModelAliasSet"`

	// <p>符合条件的总数。</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelAliasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelAliasesResponseParams `json:"Response"`
}

func (r *DescribeModelAliasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelAliasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelAssociationsRequestParams struct {
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>翻页限制</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>翻页偏移量</p><p>默认值：0</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeModelAssociationsRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>翻页限制</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>翻页偏移量</p><p>默认值：0</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeModelAssociationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelAssociationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelAssociationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelAssociationsResponseParams struct {
	// <p>模型路由实例与模型的关联关系集合</p>
	ModelAssociationSet []*ModelAssociation `json:"ModelAssociationSet,omitnil,omitempty" name:"ModelAssociationSet"`

	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>符合条件的总数</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelAssociationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelAssociationsResponseParams `json:"Response"`
}

func (r *DescribeModelAssociationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelAssociationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelKeysRequestParams struct {
	// <p>接入类型过滤：PublicBYOK/PublicCustom/PrivateCustom</p>
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// <p>过滤条件</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>返回数量限制</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>翻页启始索引</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>服务提供商ID</p>
	ServiceProviderIds []*string `json:"ServiceProviderIds,omitnil,omitempty" name:"ServiceProviderIds"`
}

type DescribeModelKeysRequest struct {
	*tchttp.BaseRequest
	
	// <p>接入类型过滤：PublicBYOK/PublicCustom/PrivateCustom</p>
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// <p>过滤条件</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>返回数量限制</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>翻页启始索引</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>服务提供商ID</p>
	ServiceProviderIds []*string `json:"ServiceProviderIds,omitnil,omitempty" name:"ServiceProviderIds"`
}

func (r *DescribeModelKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessType")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ServiceProviderIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelKeysResponseParams struct {
	// <p>模型列表（含 Key 信息）</p>
	Models []*ModelKeyInfoItem `json:"Models,omitnil,omitempty" name:"Models"`

	// <p>模型总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelKeysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelKeysResponseParams `json:"Response"`
}

func (r *DescribeModelKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelNamesRequestParams struct {
	// <p>分页偏移量（&gt;=0）</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量（1-100）</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤PrivateCustom类型自建模型。如果传递了此参数，则只返回具有相同VPC Id的模型。</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type DescribeModelNamesRequest struct {
	*tchttp.BaseRequest
	
	// <p>分页偏移量（&gt;=0）</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页数量（1-100）</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>过滤PrivateCustom类型自建模型。如果传递了此参数，则只返回具有相同VPC Id的模型。</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

func (r *DescribeModelNamesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelNamesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelNamesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelNamesResponseParams struct {
	// <p>模型标识聚合列表</p>
	ModelNames []*ModelNameAggregatedItem `json:"ModelNames,omitnil,omitempty" name:"ModelNames"`

	// <p>聚合后的模型标识总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelNamesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelNamesResponseParams `json:"Response"`
}

func (r *DescribeModelNamesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelRewriteRequestParams struct {
	// <p>模型路由实例 ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>选填，按源模型名精确过滤（大小写敏感）。</p><p>长度 1-255 字符；不传则返回该实例的全部重写规则；命中至多 1 条；未命中返回空列表（不报错）。</p>
	SourceModel *string `json:"SourceModel,omitnil,omitempty" name:"SourceModel"`
}

type DescribeModelRewriteRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例 ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>选填，按源模型名精确过滤（大小写敏感）。</p><p>长度 1-255 字符；不传则返回该实例的全部重写规则；命中至多 1 条；未命中返回空列表（不报错）。</p>
	SourceModel *string `json:"SourceModel,omitnil,omitempty" name:"SourceModel"`
}

func (r *DescribeModelRewriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelRewriteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "SourceModel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelRewriteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelRewriteResponseParams struct {
	// <p>重写规则列表，按 SourceModel 字典序排序。无规则或过滤未命中时为空数组。</p>
	Rewrites []*RewriteItem `json:"Rewrites,omitnil,omitempty" name:"Rewrites"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelRewriteResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelRewriteResponseParams `json:"Response"`
}

func (r *DescribeModelRewriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelRewriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelRouterDetailRequestParams struct {
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`
}

type DescribeModelRouterDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`
}

func (r *DescribeModelRouterDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelRouterDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelRouterDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelRouterDetailResponseParams struct {
	// <p>模型路由实例详情</p>
	ModelRouter *ModelRouterDetail `json:"ModelRouter,omitnil,omitempty" name:"ModelRouter"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelRouterDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelRouterDetailResponseParams `json:"Response"`
}

func (r *DescribeModelRouterDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelRouterDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelRouterGuardrailsRequestParams struct {
	// <p>模型路由实例 ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`
}

type DescribeModelRouterGuardrailsRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例 ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`
}

func (r *DescribeModelRouterGuardrailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelRouterGuardrailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelRouterGuardrailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelRouterGuardrailsResponseParams struct {
	// <p>当前已关联的 Guardrail 防护配置列表。</p><p>当前最多返回 1 个元素。</p>
	Guardrails []*GuardrailConfig `json:"Guardrails,omitnil,omitempty" name:"Guardrails"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelRouterGuardrailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelRouterGuardrailsResponseParams `json:"Response"`
}

func (r *DescribeModelRouterGuardrailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelRouterGuardrailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelRouterLogsRequestParams struct {
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>Key的ID</p>
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// <p>模型名称</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>请求状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>开始时间，与EndTime需要同时传入、开始时间不得早于24小时前，默认仅查询近5分钟日志</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，与StartTime需要同时传入、开始时间不得早于24小时前，默认仅查询近5分钟日志</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>游标NextToken</p>
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// <p>本次查询最大数量</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`
}

type DescribeModelRouterLogsRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>Key的ID</p>
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// <p>模型名称</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>请求状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>开始时间，与EndTime需要同时传入、开始时间不得早于24小时前，默认仅查询近5分钟日志</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，与StartTime需要同时传入、开始时间不得早于24小时前，默认仅查询近5分钟日志</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>游标NextToken</p>
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// <p>本次查询最大数量</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`
}

func (r *DescribeModelRouterLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelRouterLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "KeyId")
	delete(f, "Model")
	delete(f, "Status")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "NextToken")
	delete(f, "MaxResults")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelRouterLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelRouterLogsResponseParams struct {
	// <p>日志列表</p>
	Logs []*ModelRouterLog `json:"Logs,omitnil,omitempty" name:"Logs"`

	// <p>满足条件的数量</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>游标NextToken</p>
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelRouterLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelRouterLogsResponseParams `json:"Response"`
}

func (r *DescribeModelRouterLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelRouterLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelRouterQuotaRequestParams struct {
	// <p>配额类型</p>
	QuotaTypes []*string `json:"QuotaTypes,omitnil,omitempty" name:"QuotaTypes"`

	// <p>要查询的资源ID</p>
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// <p>需要展示的字段</p><p>枚举值：</p><ul><li>Used： 已使用的配额数量</li><li>Available： 剩余的配额数量</li></ul>
	DisplayFields []*string `json:"DisplayFields,omitnil,omitempty" name:"DisplayFields"`
}

type DescribeModelRouterQuotaRequest struct {
	*tchttp.BaseRequest
	
	// <p>配额类型</p>
	QuotaTypes []*string `json:"QuotaTypes,omitnil,omitempty" name:"QuotaTypes"`

	// <p>要查询的资源ID</p>
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// <p>需要展示的字段</p><p>枚举值：</p><ul><li>Used： 已使用的配额数量</li><li>Available： 剩余的配额数量</li></ul>
	DisplayFields []*string `json:"DisplayFields,omitnil,omitempty" name:"DisplayFields"`
}

func (r *DescribeModelRouterQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelRouterQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QuotaTypes")
	delete(f, "ResourceIds")
	delete(f, "DisplayFields")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelRouterQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelRouterQuotaResponseParams struct {
	// <p>配额信息</p>
	Quotas []*ModelRouterQuota `json:"Quotas,omitnil,omitempty" name:"Quotas"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelRouterQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelRouterQuotaResponseParams `json:"Response"`
}

func (r *DescribeModelRouterQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelRouterQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelRouterResourcePackageDeductionRequestParams struct {
	// <p>抵扣起始时间</p>
	DeductionTimeBegin *string `json:"DeductionTimeBegin,omitnil,omitempty" name:"DeductionTimeBegin"`

	// <p>抵扣截止时间</p>
	DeductionTimeEnd *string `json:"DeductionTimeEnd,omitnil,omitempty" name:"DeductionTimeEnd"`

	// <p>模型路由资源包Id</p>
	ModelRouterResourcePackageId *string `json:"ModelRouterResourcePackageId,omitnil,omitempty" name:"ModelRouterResourcePackageId"`

	// <p>返回的数量</p><p>取值范围：[0, 100]</p><p>默认值：20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>数据偏移量</p><p>默认值：0</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>排序方式：asc，desc</p>
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`
}

type DescribeModelRouterResourcePackageDeductionRequest struct {
	*tchttp.BaseRequest
	
	// <p>抵扣起始时间</p>
	DeductionTimeBegin *string `json:"DeductionTimeBegin,omitnil,omitempty" name:"DeductionTimeBegin"`

	// <p>抵扣截止时间</p>
	DeductionTimeEnd *string `json:"DeductionTimeEnd,omitnil,omitempty" name:"DeductionTimeEnd"`

	// <p>模型路由资源包Id</p>
	ModelRouterResourcePackageId *string `json:"ModelRouterResourcePackageId,omitnil,omitempty" name:"ModelRouterResourcePackageId"`

	// <p>返回的数量</p><p>取值范围：[0, 100]</p><p>默认值：20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>数据偏移量</p><p>默认值：0</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>排序方式：asc，desc</p>
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`
}

func (r *DescribeModelRouterResourcePackageDeductionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelRouterResourcePackageDeductionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeductionTimeBegin")
	delete(f, "DeductionTimeEnd")
	delete(f, "ModelRouterResourcePackageId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SortBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelRouterResourcePackageDeductionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelRouterResourcePackageDeductionResponseParams struct {
	// <p>模型路由资源包抵扣信息</p>
	ModelRouterResourcePackageDeductionSet []*ModelRouterResourcePackageDeduction `json:"ModelRouterResourcePackageDeductionSet,omitnil,omitempty" name:"ModelRouterResourcePackageDeductionSet"`

	// <p>符合查询条件的详情信息总数</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelRouterResourcePackageDeductionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelRouterResourcePackageDeductionResponseParams `json:"Response"`
}

func (r *DescribeModelRouterResourcePackageDeductionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelRouterResourcePackageDeductionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelRouterResourcePackagesRequestParams struct {
	// <p>模型路由资源包Id。</p>
	ModelRouterResourcePackageIds []*string `json:"ModelRouterResourcePackageIds,omitnil,omitempty" name:"ModelRouterResourcePackageIds"`

	// <p>数据偏移量。</p><p>默认值：0</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回的数量，最大值为100。</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>排序参数，支持以下字段：&quot;buyTime&quot;, &quot;startTime&quot;, &quot;endTime&quot;</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>排序方式：asc，desc，默认asc</p>
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// <p>查询的过滤条件。</p><p>每次请求的Filters的上限为10，Filter.Values的上限为100。 Filter.Name和Filter.Values皆为必填项。详细的过滤条件如下： status - Integer - 是否必填：否 - 状态：0-有效 1-已退款 2-已过期 3-已用完。 effect_time_start - String - 是否必填：否 - 生效起始时间,YYYY-MM-DD HH:MM:SS格式。 effect_time_end - String - 是否必填：否 - 生效截止时间。 expire_time_start - String - 是否必填：否 - 失效起始时间。 expire_time_end - String - 是否必填：否 - 失效截止时间。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeModelRouterResourcePackagesRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由资源包Id。</p>
	ModelRouterResourcePackageIds []*string `json:"ModelRouterResourcePackageIds,omitnil,omitempty" name:"ModelRouterResourcePackageIds"`

	// <p>数据偏移量。</p><p>默认值：0</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>返回的数量，最大值为100。</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>排序参数，支持以下字段：&quot;buyTime&quot;, &quot;startTime&quot;, &quot;endTime&quot;</p>
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// <p>排序方式：asc，desc，默认asc</p>
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// <p>查询的过滤条件。</p><p>每次请求的Filters的上限为10，Filter.Values的上限为100。 Filter.Name和Filter.Values皆为必填项。详细的过滤条件如下： status - Integer - 是否必填：否 - 状态：0-有效 1-已退款 2-已过期 3-已用完。 effect_time_start - String - 是否必填：否 - 生效起始时间,YYYY-MM-DD HH:MM:SS格式。 effect_time_end - String - 是否必填：否 - 生效截止时间。 expire_time_start - String - 是否必填：否 - 失效起始时间。 expire_time_end - String - 是否必填：否 - 失效截止时间。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeModelRouterResourcePackagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelRouterResourcePackagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterResourcePackageIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "SortBy")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelRouterResourcePackagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelRouterResourcePackagesResponseParams struct {
	// <p>模型路由资源包信息</p>
	ModelRouterResourcePackageSet []*ModelRouterPackage `json:"ModelRouterResourcePackageSet,omitnil,omitempty" name:"ModelRouterResourcePackageSet"`

	// <p>符合查询条件的模型路由资源包数量</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>资源包的剩余总量</p>
	TotalDosage *uint64 `json:"TotalDosage,omitnil,omitempty" name:"TotalDosage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelRouterResourcePackagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelRouterResourcePackagesResponseParams `json:"Response"`
}

func (r *DescribeModelRouterResourcePackagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelRouterResourcePackagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelRoutersRequestParams struct {
	// <p>过滤条件</p><p>支持：ModelRouterName、ModelRouterType、Status、BudgetId、tag-key、tag:&lt;tag-key&gt;。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>每页数量，1-100，默认 20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>模型路由实例ID列表</p>
	ModelRouterIds []*string `json:"ModelRouterIds,omitnil,omitempty" name:"ModelRouterIds"`

	// <p>分页偏移量，默认 0</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeModelRoutersRequest struct {
	*tchttp.BaseRequest
	
	// <p>过滤条件</p><p>支持：ModelRouterName、ModelRouterType、Status、BudgetId、tag-key、tag:&lt;tag-key&gt;。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>每页数量，1-100，默认 20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>模型路由实例ID列表</p>
	ModelRouterIds []*string `json:"ModelRouterIds,omitnil,omitempty" name:"ModelRouterIds"`

	// <p>分页偏移量，默认 0</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeModelRoutersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelRoutersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "ModelRouterIds")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelRoutersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelRoutersResponseParams struct {
	// <p>模型路由实例列表</p>
	ModelRouterSet []*ModelRouterSet `json:"ModelRouterSet,omitnil,omitempty" name:"ModelRouterSet"`

	// <p>符合条件的总数</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelRoutersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelRoutersResponseParams `json:"Response"`
}

func (r *DescribeModelRoutersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelRoutersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuotaRequestParams struct {

}

type DescribeQuotaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuotaResponseParams struct {
	// 配额列表
	QuotaSet []*Quota `json:"QuotaSet,omitnil,omitempty" name:"QuotaSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeQuotaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQuotaResponseParams `json:"Response"`
}

func (r *DescribeQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesRequestParams struct {
	// 返回可用区资源列表数目，默认20，最大值100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回可用区资源列表起始偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询可用区资源列表条件，详细的过滤条件如下：
	// - master-zone
	// 按照【地域可用区】进行过滤，例如：ap-guangzhou-2。
	// 类型：String
	// 必选：否
	// - ip-version
	// 按照【IP 类型】进行过滤，例如：IPv4。
	// 类型：String
	// 必选：否
	// 可选项：IPv4、IPv6、IPv6_Nat
	// - isp
	// 按照【ISP 类型】进行过滤，例如：BGP。
	// 类型：String
	// 必选：否
	// 可选项：BGP、CMCC（中国移动）、CUCC（中国联通）、CTCC（中国电信）、BGP_PRO、INTERNAL（内网）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 返回可用区资源列表数目，默认20，最大值100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回可用区资源列表起始偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询可用区资源列表条件，详细的过滤条件如下：
	// - master-zone
	// 按照【地域可用区】进行过滤，例如：ap-guangzhou-2。
	// 类型：String
	// 必选：否
	// - ip-version
	// 按照【IP 类型】进行过滤，例如：IPv4。
	// 类型：String
	// 必选：否
	// 可选项：IPv4、IPv6、IPv6_Nat
	// - isp
	// 按照【ISP 类型】进行过滤，例如：BGP。
	// 类型：String
	// 必选：否
	// 可选项：BGP、CMCC（中国移动）、CUCC（中国联通）、CTCC（中国电信）、BGP_PRO、INTERNAL（内网）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcesResponseParams struct {
	// 可用区支持的资源列表。
	ZoneResourceSet []*ZoneResource `json:"ZoneResourceSet,omitnil,omitempty" name:"ZoneResourceSet"`

	// 可用区资源列表数目。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcesResponseParams `json:"Response"`
}

func (r *DescribeResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRewriteRequestParams struct {
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID数组。
	SourceListenerIds []*string `json:"SourceListenerIds,omitnil,omitempty" name:"SourceListenerIds"`

	// 负载均衡转发规则的ID数组。
	SourceLocationIds []*string `json:"SourceLocationIds,omitnil,omitempty" name:"SourceLocationIds"`
}

type DescribeRewriteRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID数组。
	SourceListenerIds []*string `json:"SourceListenerIds,omitnil,omitempty" name:"SourceListenerIds"`

	// 负载均衡转发规则的ID数组。
	SourceLocationIds []*string `json:"SourceLocationIds,omitnil,omitempty" name:"SourceLocationIds"`
}

func (r *DescribeRewriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRewriteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "SourceListenerIds")
	delete(f, "SourceLocationIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRewriteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRewriteResponseParams struct {
	// 重定向转发规则构成的数组，若无重定向规则，则返回空数组。
	RewriteSet []*RuleOutput `json:"RewriteSet,omitnil,omitempty" name:"RewriteSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRewriteResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRewriteResponseParams `json:"Response"`
}

func (r *DescribeRewriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRewriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceProviderHealthStatusRequestParams struct {
	// <p>BYOK的ID</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>本次查询的限制数量</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>本次查询翻页的偏移量</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeServiceProviderHealthStatusRequest struct {
	*tchttp.BaseRequest
	
	// <p>BYOK的ID</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>本次查询的限制数量</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>本次查询翻页的偏移量</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeServiceProviderHealthStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceProviderHealthStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceProviderId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceProviderHealthStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceProviderHealthStatusResponseParams struct {
	// <p>健康检查的结果</p>
	HealthCheckResults []*ModelHealthCheckResults `json:"HealthCheckResults,omitnil,omitempty" name:"HealthCheckResults"`

	// <p>本次请求总数</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServiceProviderHealthStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceProviderHealthStatusResponseParams `json:"Response"`
}

func (r *DescribeServiceProviderHealthStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceProviderHealthStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportedProvidersRequestParams struct {

}

type DescribeSupportedProvidersRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSupportedProvidersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportedProvidersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSupportedProvidersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSupportedProvidersResponseParams struct {
	// <p>Provider 列表</p>
	Providers []*ProviderItem `json:"Providers,omitnil,omitempty" name:"Providers"`

	// <p>Provider 总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSupportedProvidersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSupportedProvidersResponseParams `json:"Response"`
}

func (r *DescribeSupportedProvidersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSupportedProvidersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupInstanceStatusRequestParams struct {
	// 目标组唯一id
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 目标组绑定的后端服务ip列表
	TargetGroupInstanceIps []*string `json:"TargetGroupInstanceIps,omitnil,omitempty" name:"TargetGroupInstanceIps"`
}

type DescribeTargetGroupInstanceStatusRequest struct {
	*tchttp.BaseRequest
	
	// 目标组唯一id
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 目标组绑定的后端服务ip列表
	TargetGroupInstanceIps []*string `json:"TargetGroupInstanceIps,omitnil,omitempty" name:"TargetGroupInstanceIps"`
}

func (r *DescribeTargetGroupInstanceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupInstanceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "TargetGroupInstanceIps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetGroupInstanceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupInstanceStatusResponseParams struct {
	// 健康检查后端rs状态列表
	TargetGroupInstanceSet []*TargetGroupInstanceStatus `json:"TargetGroupInstanceSet,omitnil,omitempty" name:"TargetGroupInstanceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTargetGroupInstanceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetGroupInstanceStatusResponseParams `json:"Response"`
}

func (r *DescribeTargetGroupInstanceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupInstanceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupInstancesRequestParams struct {
	// 过滤条件，当前支持按照 TargetGroupId，BindIP，InstanceId 多个条件组合过滤。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 显示数量限制，默认20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 显示的偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeTargetGroupInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，当前支持按照 TargetGroupId，BindIP，InstanceId 多个条件组合过滤。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 显示数量限制，默认20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 显示的偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeTargetGroupInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetGroupInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupInstancesResponseParams struct {
	// 本次查询的结果数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 绑定的服务器信息。
	TargetGroupInstanceSet []*TargetGroupBackend `json:"TargetGroupInstanceSet,omitnil,omitempty" name:"TargetGroupInstanceSet"`

	// 实际统计数量，不受Limit、Offset、CAM的影响。
	RealCount *uint64 `json:"RealCount,omitnil,omitempty" name:"RealCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTargetGroupInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetGroupInstancesResponseParams `json:"Response"`
}

func (r *DescribeTargetGroupInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupListRequestParams struct {
	// <p>目标组ID数组。</p>
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`

	// <p>过滤条件数组，支持TargetGroupVpcId和TargetGroupName。与TargetGroupIds互斥，优先使用目标组ID。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>显示的偏移起始量。</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页显示条目数。</p><p>取值范围：[0, 100]</p><p>默认值：20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTargetGroupListRequest struct {
	*tchttp.BaseRequest
	
	// <p>目标组ID数组。</p>
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`

	// <p>过滤条件数组，支持TargetGroupVpcId和TargetGroupName。与TargetGroupIds互斥，优先使用目标组ID。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>显示的偏移起始量。</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>每页显示条目数。</p><p>取值范围：[0, 100]</p><p>默认值：20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeTargetGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupListResponseParams struct {
	// <p>显示的结果数量。</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>显示的目标组信息集合。</p>
	TargetGroupSet []*TargetGroupInfo `json:"TargetGroupSet,omitnil,omitempty" name:"TargetGroupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTargetGroupListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetGroupListResponseParams `json:"Response"`
}

func (r *DescribeTargetGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupsRequestParams struct {
	// 目标组ID，与Filters互斥。
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`

	// 显示条数限制，默认为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 显示的偏移起始量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件数组，与TargetGroupIds互斥，支持 TargetGroupVpcId（私有网络 ID）和 TargetGroupName（目标组名称）以及 Tag（标签）。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTargetGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 目标组ID，与Filters互斥。
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`

	// 显示条数限制，默认为20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 显示的偏移起始量。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件数组，与TargetGroupIds互斥，支持 TargetGroupVpcId（私有网络 ID）和 TargetGroupName（目标组名称）以及 Tag（标签）。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeTargetGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupIds")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupsResponseParams struct {
	// 显示的结果数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 显示的目标组信息集合。
	TargetGroupSet []*TargetGroupInfo `json:"TargetGroupSet,omitnil,omitempty" name:"TargetGroupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTargetGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetGroupsResponseParams `json:"Response"`
}

func (r *DescribeTargetGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetHealthRequestParams struct {
	// <p>要查询的负载均衡实例ID列表。数组大小最大支持30。</p>
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// <p>要查询的监听器ID列表。</p>
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// <p>要查询的转发规则ID列表。</p>
	LocationIds []*string `json:"LocationIds,omitnil,omitempty" name:"LocationIds"`
}

type DescribeTargetHealthRequest struct {
	*tchttp.BaseRequest
	
	// <p>要查询的负载均衡实例ID列表。数组大小最大支持30。</p>
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// <p>要查询的监听器ID列表。</p>
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// <p>要查询的转发规则ID列表。</p>
	LocationIds []*string `json:"LocationIds,omitnil,omitempty" name:"LocationIds"`
}

func (r *DescribeTargetHealthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetHealthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	delete(f, "ListenerIds")
	delete(f, "LocationIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetHealthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetHealthResponseParams struct {
	// <p>负载均衡实例列表。</p>
	LoadBalancers []*LoadBalancerHealth `json:"LoadBalancers,omitnil,omitempty" name:"LoadBalancers"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTargetHealthResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetHealthResponseParams `json:"Response"`
}

func (r *DescribeTargetHealthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetHealthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetsRequestParams struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 监听器 ID 列表。ID 数量上限为20个。
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// 监听器协议类型。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 监听器端口。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 查询负载均衡绑定的后端服务列表，过滤条件如下：
	// <li> location-id - String - 是否必填：否 - （过滤条件）按照 规则ID 过滤，如："loc-12345678"。</li>
	// <li> private-ip-address - String - 是否必填：否 - （过滤条件）按照 后端服务内网IP 过滤，如："172.16.1.1"。</li>
	// <li> tag - String - 是否必填：否 - （过滤条件）按照 标签 过滤，如："tag-test"。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 监听器 ID 列表。ID 数量上限为20个。
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// 监听器协议类型。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 监听器端口。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 查询负载均衡绑定的后端服务列表，过滤条件如下：
	// <li> location-id - String - 是否必填：否 - （过滤条件）按照 规则ID 过滤，如："loc-12345678"。</li>
	// <li> private-ip-address - String - 是否必填：否 - （过滤条件）按照 后端服务内网IP 过滤，如："172.16.1.1"。</li>
	// <li> tag - String - 是否必填：否 - （过滤条件）按照 标签 过滤，如："tag-test"。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerIds")
	delete(f, "Protocol")
	delete(f, "Port")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetsResponseParams struct {
	// 监听器后端绑定的机器信息。
	Listeners []*ListenerBackend `json:"Listeners,omitnil,omitempty" name:"Listeners"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTargetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetsResponseParams `json:"Response"`
}

func (r *DescribeTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusRequestParams struct {
	// 请求ID，即接口返回的 RequestId 参数。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 订单ID。
	// 注意：参数TaskId和DealName必须传一个。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 请求ID，即接口返回的 RequestId 参数。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 订单ID。
	// 注意：参数TaskId和DealName必须传一个。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`
}

func (r *DescribeTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "DealName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusResponseParams struct {
	// 任务的当前状态。 0：成功，1：失败，2：进行中。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 由负载均衡实例唯一 ID 组成的数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 辅助描述信息，如失败原因等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUpperModelsRequestParams struct {
	// <p>接入类型：PublicBYOK/PublicCustom/PrivateCustom</p>
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// <p>上游 Provider API 地址</p><p>示例：https://api.moonshot.cn</p>
	ApiBase *string `json:"ApiBase,omitnil,omitempty" name:"ApiBase"`

	// <p>上游 Provider API Key</p><p>用于鉴权访问上游模型列表接口</p>
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`

	// <p>自定义 Host Header，可选</p><p>仅 VPC 内网场景需要，用于指定请求的 Host 头</p>
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`

	// <p>Key Id 配合ServiceProviderId一同输入，不指定则默认选用最近创建的Key</p>
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// <p>模型列表端点路径，可选</p><p>默认值：/v1/models</p>
	ModelPath *string `json:"ModelPath,omitnil,omitempty" name:"ModelPath"`

	// <p>模型协议</p>
	ModelProtocol *string `json:"ModelProtocol,omitnil,omitempty" name:"ModelProtocol"`

	// <p>模型提供商</p>
	ModelProvider *string `json:"ModelProvider,omitnil,omitempty" name:"ModelProvider"`

	// <p>BYOK 业务 ID，可选</p><p>格式：byok-xxxxxxxx</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`
}

type DescribeUpperModelsRequest struct {
	*tchttp.BaseRequest
	
	// <p>接入类型：PublicBYOK/PublicCustom/PrivateCustom</p>
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// <p>上游 Provider API 地址</p><p>示例：https://api.moonshot.cn</p>
	ApiBase *string `json:"ApiBase,omitnil,omitempty" name:"ApiBase"`

	// <p>上游 Provider API Key</p><p>用于鉴权访问上游模型列表接口</p>
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`

	// <p>自定义 Host Header，可选</p><p>仅 VPC 内网场景需要，用于指定请求的 Host 头</p>
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`

	// <p>Key Id 配合ServiceProviderId一同输入，不指定则默认选用最近创建的Key</p>
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// <p>模型列表端点路径，可选</p><p>默认值：/v1/models</p>
	ModelPath *string `json:"ModelPath,omitnil,omitempty" name:"ModelPath"`

	// <p>模型协议</p>
	ModelProtocol *string `json:"ModelProtocol,omitnil,omitempty" name:"ModelProtocol"`

	// <p>模型提供商</p>
	ModelProvider *string `json:"ModelProvider,omitnil,omitempty" name:"ModelProvider"`

	// <p>BYOK 业务 ID，可选</p><p>格式：byok-xxxxxxxx</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`
}

func (r *DescribeUpperModelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpperModelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AccessType")
	delete(f, "ApiBase")
	delete(f, "ApiKey")
	delete(f, "HostHeader")
	delete(f, "KeyId")
	delete(f, "ModelPath")
	delete(f, "ModelProtocol")
	delete(f, "ModelProvider")
	delete(f, "ServiceProviderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUpperModelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUpperModelsResponseParams struct {
	// <p>上游模型列表</p>
	Models []*string `json:"Models,omitnil,omitempty" name:"Models"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUpperModelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUpperModelsResponseParams `json:"Response"`
}

func (r *DescribeUpperModelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpperModelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserGroupsRequestParams struct {
	// <p>模型路由实例ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>用户组ID列表，用于按ID过滤，单次最多100个；可包含「未分组」虚拟分组 ugrp-ungrouped。</p>
	UserGroupIds []*string `json:"UserGroupIds,omitnil,omitempty" name:"UserGroupIds"`

	// <p>过滤列表。支持：UserGroupName、Status、tag-key、tag:&lt;tag-key&gt;。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>本次查询限制的数量</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>本次查询偏移量</p><p>默认值：0</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeUserGroupsRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>用户组ID列表，用于按ID过滤，单次最多100个；可包含「未分组」虚拟分组 ugrp-ungrouped。</p>
	UserGroupIds []*string `json:"UserGroupIds,omitnil,omitempty" name:"UserGroupIds"`

	// <p>过滤列表。支持：UserGroupName、Status、tag-key、tag:&lt;tag-key&gt;。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>本次查询限制的数量</p><p>取值范围：[1, 100]</p><p>默认值：20</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>本次查询偏移量</p><p>默认值：0</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeUserGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "UserGroupIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserGroupsResponseParams struct {
	// <p>用户组列表。</p>
	UserGroups []*UserGroupInfo `json:"UserGroups,omitnil,omitempty" name:"UserGroups"`

	// <p>符合条件的总数（含「未分组」逻辑组 ugrp-ungrouped：当其未被过滤条件排除时计入，即 TotalCount = 真实用户组数 + 1）。</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserGroupsResponseParams `json:"Response"`
}

func (r *DescribeUserGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateBudgetRequestParams struct {
	// <p>Budget ID。</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>要解除关联的资源列表。</p>
	Resources []*BudgetResource `json:"Resources,omitnil,omitempty" name:"Resources"`
}

type DisassociateBudgetRequest struct {
	*tchttp.BaseRequest
	
	// <p>Budget ID。</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>要解除关联的资源列表。</p>
	Resources []*BudgetResource `json:"Resources,omitnil,omitempty" name:"Resources"`
}

func (r *DisassociateBudgetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateBudgetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BudgetId")
	delete(f, "Resources")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateBudgetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateBudgetResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateBudgetResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateBudgetResponseParams `json:"Response"`
}

func (r *DisassociateBudgetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateBudgetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateCustomizedConfigRequestParams struct {
	// 配置ID
	UconfigId *string `json:"UconfigId,omitnil,omitempty" name:"UconfigId"`

	// 解绑的列表
	BindList []*BindItem `json:"BindList,omitnil,omitempty" name:"BindList"`
}

type DisassociateCustomizedConfigRequest struct {
	*tchttp.BaseRequest
	
	// 配置ID
	UconfigId *string `json:"UconfigId,omitnil,omitempty" name:"UconfigId"`

	// 解绑的列表
	BindList []*BindItem `json:"BindList,omitnil,omitempty" name:"BindList"`
}

func (r *DisassociateCustomizedConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateCustomizedConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UconfigId")
	delete(f, "BindList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateCustomizedConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateCustomizedConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateCustomizedConfigResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateCustomizedConfigResponseParams `json:"Response"`
}

func (r *DisassociateCustomizedConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateCustomizedConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisassociateGuardrailConfig struct {
	// <p>Guardrail 防护配置 ID。</p><p>可通过 DescribeModelRouterGuardrails 获取；DisassociateModelRouterGuardrails 使用该字段定位要解除关联的防护配置。</p>
	GuardrailId *string `json:"GuardrailId,omitnil,omitempty" name:"GuardrailId"`
}

// Predefined struct for user
type DisassociateModelRouterGuardrailsRequestParams struct {
	// <p>待解除关联的 Guardrail 防护配置列表。</p><p>每个元素只需要填写 GuardrailId。</p>
	Guardrails []*DisassociateGuardrailConfig `json:"Guardrails,omitnil,omitempty" name:"Guardrails"`

	// <p>模型路由实例 ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`
}

type DisassociateModelRouterGuardrailsRequest struct {
	*tchttp.BaseRequest
	
	// <p>待解除关联的 Guardrail 防护配置列表。</p><p>每个元素只需要填写 GuardrailId。</p>
	Guardrails []*DisassociateGuardrailConfig `json:"Guardrails,omitnil,omitempty" name:"Guardrails"`

	// <p>模型路由实例 ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`
}

func (r *DisassociateModelRouterGuardrailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateModelRouterGuardrailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Guardrails")
	delete(f, "ModelRouterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateModelRouterGuardrailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateModelRouterGuardrailsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateModelRouterGuardrailsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateModelRouterGuardrailsResponseParams `json:"Response"`
}

func (r *DisassociateModelRouterGuardrailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateModelRouterGuardrailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateModelsFromModelRouterRequestParams struct {
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>需要解除关联的模型信息</p>
	Models []*ModelRouterModel `json:"Models,omitnil,omitempty" name:"Models"`
}

type DisassociateModelsFromModelRouterRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>需要解除关联的模型信息</p>
	Models []*ModelRouterModel `json:"Models,omitnil,omitempty" name:"Models"`
}

func (r *DisassociateModelsFromModelRouterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateModelsFromModelRouterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "Models")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateModelsFromModelRouterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateModelsFromModelRouterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateModelsFromModelRouterResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateModelsFromModelRouterResponseParams `json:"Response"`
}

func (r *DisassociateModelsFromModelRouterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateModelsFromModelRouterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateTargetGroupsRequestParams struct {
	// 待解绑的规则关系数组，支持批量解绑多个监听器，单次批量解除最多20个。
	Associations []*TargetGroupAssociation `json:"Associations,omitnil,omitempty" name:"Associations"`
}

type DisassociateTargetGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 待解绑的规则关系数组，支持批量解绑多个监听器，单次批量解除最多20个。
	Associations []*TargetGroupAssociation `json:"Associations,omitnil,omitempty" name:"Associations"`
}

func (r *DisassociateTargetGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateTargetGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Associations")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateTargetGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateTargetGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateTargetGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateTargetGroupsResponseParams `json:"Response"`
}

func (r *DisassociateTargetGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateTargetGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExclusiveCluster struct {
	// 4层独占集群列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	L4Clusters []*ClusterItem `json:"L4Clusters,omitnil,omitempty" name:"L4Clusters"`

	// 7层独占集群列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	L7Clusters []*ClusterItem `json:"L7Clusters,omitnil,omitempty" name:"L7Clusters"`

	// vpcgw集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassicalCluster *ClusterItem `json:"ClassicalCluster,omitnil,omitempty" name:"ClassicalCluster"`
}

type ExtraInfo struct {
	// 是否开通VIP直通
	ZhiTong *bool `json:"ZhiTong,omitnil,omitempty" name:"ZhiTong"`

	// TgwGroup名称
	TgwGroupName *string `json:"TgwGroupName,omitnil,omitempty" name:"TgwGroupName"`
}

type FallBackItem struct {
	// <p>默认回退模型列表</p>
	DefaultFallBackModels []*string `json:"DefaultFallBackModels,omitnil,omitempty" name:"DefaultFallBackModels"`
}

type Filter struct {
	// <p>过滤器的名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>过滤器的值数组</p>
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FunctionInfo struct {
	// 函数命名空间
	FunctionNamespace *string `json:"FunctionNamespace,omitnil,omitempty" name:"FunctionNamespace"`

	// 函数名称
	FunctionName *string `json:"FunctionName,omitnil,omitempty" name:"FunctionName"`

	// 函数的版本名称或别名
	FunctionQualifier *string `json:"FunctionQualifier,omitnil,omitempty" name:"FunctionQualifier"`

	// 标识 FunctionQualifier 参数的类型，可取值： VERSION（版本）、ALIAS（别名）
	// 注意：此字段可能返回 null，表示取不到有效值。
	FunctionQualifierType *string `json:"FunctionQualifierType,omitnil,omitempty" name:"FunctionQualifierType"`
}

type FunctionTarget struct {
	// 云函数相关信息
	Function *FunctionInfo `json:"Function,omitnil,omitempty" name:"Function"`

	// 权重
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type GuardrailConfig struct {
	// <p>Guardrail 防护配置 ID。</p><p>DescribeModelRouterGuardrails 会返回该字段；DisassociateModelRouterGuardrails 和 ModifyModelRouterGuardrails 需要使用该字段定位要操作的防护配置。</p>
	GuardrailId *string `json:"GuardrailId,omitnil,omitempty" name:"GuardrailId"`

	// <p>Guardrail 防护类型。</p><p>枚举值：</p><ul><li>WAF：使用腾讯云 WAF LLM SDK 接入配置对模型路由请求进行安全防护。</li></ul><p>当前仅支持 WAF；AssociateModelRouterGuardrails 不传时默认为 WAF，ModifyModelRouterGuardrails 不传时沿用当前已关联 Guardrail 的 Type。</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>关联的腾讯云 WAF 实例 ID。</p><p>ModifyModelRouterGuardrails 在 Type 为 WAF 时必填。DescribeModelRouterGuardrails 返回。接口会校验该 WAF 实例存在且属于当前账号。</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>WAF LLM SDK 接入服务 ID。</p><p>该字段对应 WAF LLM SDK 接入配置中的服务标识，用于指定模型路由请求要绑定的 WAF 防护配置。ModifyModelRouterGuardrails 在 Type 为 WAF 时必填。DescribeModelRouterGuardrails 返回。接口会校验该服务配置存在于指定的 WAF 实例下。</p>
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// <p>最大检测对话轮数。</p><p>ModifyModelRouterGuardrails 选填；未传时沿用当前已关联 Guardrail 的 InputCheckDepth。DescribeModelRouterGuardrails 返回。若传入，取值必须为正整数。</p>
	InputCheckDepth *uint64 `json:"InputCheckDepth,omitnil,omitempty" name:"InputCheckDepth"`
}

type HealthCheck struct {
	// 是否开启健康检查：1（开启）、0（关闭）。
	// 默认为开启。
	HealthSwitch *int64 `json:"HealthSwitch,omitnil,omitempty" name:"HealthSwitch"`

	// 健康检查的响应超时时间，可选值：2~60，默认值：2，单位：秒。响应超时时间要小于检查间隔时间。
	TimeOut *int64 `json:"TimeOut,omitnil,omitempty" name:"TimeOut"`

	// 健康检查探测间隔时间，默认值：5，IPv4 CLB实例的取值范围为：2-300，IPv6 CLB 实例的取值范围为：5-300。单位：秒。
	// 说明：部分老旧 IPv4 CLB实例的取值范围为：5-300。
	IntervalTime *int64 `json:"IntervalTime,omitnil,omitempty" name:"IntervalTime"`

	// 健康阈值，默认值：3，表示当连续探测三次健康则表示该转发正常，可选值：2~10，单位：次。
	HealthNum *int64 `json:"HealthNum,omitnil,omitempty" name:"HealthNum"`

	// 不健康阈值，默认值：3，表示当连续探测三次不健康则表示该转发异常，可选值：2~10，单位：次。
	UnHealthNum *int64 `json:"UnHealthNum,omitnil,omitempty" name:"UnHealthNum"`

	// 健康检查状态码（仅适用于HTTP/HTTPS转发规则、TCP监听器的HTTP健康检查方式）。可选值：1~31，默认 31。
	// 1 表示探测后返回值 1xx 代表健康，2 表示返回 2xx 代表健康，4 表示返回 3xx 代表健康，8 表示返回 4xx 代表健康，16 表示返回 5xx 代表健康。若希望多种返回码都可代表健康，则将相应的值相加。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpCode *int64 `json:"HttpCode,omitnil,omitempty" name:"HttpCode"`

	// 健康检查路径（仅适用于HTTP/HTTPS转发规则、TCP监听器的HTTP健康检查方式）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpCheckPath *string `json:"HttpCheckPath,omitnil,omitempty" name:"HttpCheckPath"`

	// 健康检查域名，将在HTTP协议 Host 头字段中携带。（仅适用于HTTP/HTTPS监听器和TCP监听器的HTTP健康检查方式。针对TCP监听器，当使用HTTP健康检查方式时，该参数为必填项）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpCheckDomain *string `json:"HttpCheckDomain,omitnil,omitempty" name:"HttpCheckDomain"`

	// 健康检查方法（仅适用于HTTP/HTTPS转发规则、TCP监听器的HTTP健康检查方式），默认值：HEAD，可选值HEAD或GET。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpCheckMethod *string `json:"HttpCheckMethod,omitnil,omitempty" name:"HttpCheckMethod"`

	// 自定义探测相关参数。健康检查端口，默认为后端服务的端口，除非您希望指定特定端口，否则建议留空。传参数值-1可恢复默认设置。（仅适用于TCP/UDP监听器）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckPort *int64 `json:"CheckPort,omitnil,omitempty" name:"CheckPort"`

	// 自定义探测相关参数。健康检查协议CheckType的值取CUSTOM时，必填此字段，代表健康检查的输入格式，可取值：HEX或TEXT；取值为HEX时，SendContext和RecvContext的字符只能在0123456789ABCDEF中选取且长度必须是偶数位。（仅适用于TCP/UDP监听器）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContextType *string `json:"ContextType,omitnil,omitempty" name:"ContextType"`

	// 自定义探测相关参数。健康检查协议CheckType的值取CUSTOM时，必填此字段，代表健康检查发送的请求内容，只允许ASCII可见字符，最大长度限制500。（仅适用于TCP/UDP监听器）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SendContext *string `json:"SendContext,omitnil,omitempty" name:"SendContext"`

	// 自定义探测相关参数。健康检查协议CheckType的值取CUSTOM时，必填此字段，代表健康检查返回的结果，只允许ASCII可见字符，最大长度限制500。（仅适用于TCP/UDP监听器）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecvContext *string `json:"RecvContext,omitnil,omitempty" name:"RecvContext"`

	// 健康检查使用的协议。取值 TCP | HTTP | HTTPS | GRPC | PING | CUSTOM，UDP监听器支持PING/CUSTOM，TCP监听器支持TCP/HTTP/CUSTOM，TCP_SSL/QUIC监听器支持TCP/HTTP，HTTP规则支持HTTP/GRPC，HTTPS规则支持HTTP/HTTPS/GRPC。HTTP监听器默认值为HTTP;TCP、TCP_SSL、QUIC监听器默认值为TCP;UDP监听器默认为PING;HTTPS监听器的CheckType默认值与后端转发协议一致。
	CheckType *string `json:"CheckType,omitnil,omitempty" name:"CheckType"`

	// HTTP版本。健康检查协议CheckType的值取HTTP时，必传此字段，代表后端服务的HTTP版本：HTTP/1.0、HTTP/1.1；（仅适用于TCP监听器）
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpVersion *string `json:"HttpVersion,omitnil,omitempty" name:"HttpVersion"`

	// 健康检查源IP类型：0（使用LB的VIP作为源IP），1（使用100.64网段IP作为源IP）。
	SourceIpType *int64 `json:"SourceIpType,omitnil,omitempty" name:"SourceIpType"`

	// GRPC健康检查状态码（仅适用于后端转发协议为GRPC的规则）。默认值为 12，可输入值为数值、多个数值、或者范围，例如 20 或 20,25 或 0-99
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtendedCode *string `json:"ExtendedCode,omitnil,omitempty" name:"ExtendedCode"`
}

type IdleLoadBalancer struct {
	// 负载均衡ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡名字
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 负载均衡所在地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 负载均衡的vip
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 闲置原因。NO_RULES：没有规则，NO_RS：有规则没有绑定子机。
	IdleReason *string `json:"IdleReason,omitnil,omitempty" name:"IdleReason"`

	// 负载均衡实例的状态，包括
	// 0：创建中，1：正常运行。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 负载均衡类型标识，1：负载均衡，0：传统型负载均衡。
	Forward *uint64 `json:"Forward,omitnil,omitempty" name:"Forward"`

	// 负载均衡域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type InputKeyInfo struct {
	// <p>Key的名称</p>
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// <p>导入的明文Key</p><p>仅允许导入Key模式下输入</p>
	PlainKey *string `json:"PlainKey,omitnil,omitempty" name:"PlainKey"`
}

// Predefined struct for user
type InquirePriceCreateModelRouterResourcePackageRequestParams struct {
	// <p>模型路由资源包容量</p><p>取值范围：[1000, 10000000]</p><p>单次购买的模型路由资源包容量下限为1000，上限为10000000</p>
	ModelRouterResourcePackageAmount *uint64 `json:"ModelRouterResourcePackageAmount,omitnil,omitempty" name:"ModelRouterResourcePackageAmount"`
}

type InquirePriceCreateModelRouterResourcePackageRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由资源包容量</p><p>取值范围：[1000, 10000000]</p><p>单次购买的模型路由资源包容量下限为1000，上限为10000000</p>
	ModelRouterResourcePackageAmount *uint64 `json:"ModelRouterResourcePackageAmount,omitnil,omitempty" name:"ModelRouterResourcePackageAmount"`
}

func (r *InquirePriceCreateModelRouterResourcePackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateModelRouterResourcePackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterResourcePackageAmount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceCreateModelRouterResourcePackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateModelRouterResourcePackageResponseParams struct {
	// <p>模型路由资源包价格</p>
	ModelRouterResourcePackagePrice *ItemPrice `json:"ModelRouterResourcePackagePrice,omitnil,omitempty" name:"ModelRouterResourcePackagePrice"`

	// <p>本次购买资源包是否可享受首购优惠</p><p>1:可享受首购优惠，0:不可享受首购优惠</p>
	FirstBuy *uint64 `json:"FirstBuy,omitnil,omitempty" name:"FirstBuy"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePriceCreateModelRouterResourcePackageResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceCreateModelRouterResourcePackageResponseParams `json:"Response"`
}

func (r *InquirePriceCreateModelRouterResourcePackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateModelRouterResourcePackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRefundModelRouterResourcePackageRequestParams struct {
	// <p>待退款的模型路由资源包Id</p><p>非有效状态或者设置了自动续订且自动续订已生效的资源包不允许退款。</p>
	ModelRouterResourcePackageIds []*string `json:"ModelRouterResourcePackageIds,omitnil,omitempty" name:"ModelRouterResourcePackageIds"`
}

type InquirePriceRefundModelRouterResourcePackageRequest struct {
	*tchttp.BaseRequest
	
	// <p>待退款的模型路由资源包Id</p><p>非有效状态或者设置了自动续订且自动续订已生效的资源包不允许退款。</p>
	ModelRouterResourcePackageIds []*string `json:"ModelRouterResourcePackageIds,omitnil,omitempty" name:"ModelRouterResourcePackageIds"`
}

func (r *InquirePriceRefundModelRouterResourcePackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRefundModelRouterResourcePackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterResourcePackageIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceRefundModelRouterResourcePackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceRefundModelRouterResourcePackageResponseParams struct {
	// <p>待退款的模型路由资源包可退价格</p>
	ModelRouterResourcePackageRefundPrice []*ModelRouterResourcePackageRefundPrice `json:"ModelRouterResourcePackageRefundPrice,omitnil,omitempty" name:"ModelRouterResourcePackageRefundPrice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePriceRefundModelRouterResourcePackageResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceRefundModelRouterResourcePackageResponseParams `json:"Response"`
}

func (r *InquirePriceRefundModelRouterResourcePackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceRefundModelRouterResourcePackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateLoadBalancerRequestParams struct {
	// 询价的负载均衡类型，OPEN为公网类型，INTERNAL为内网类型
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// 询价的收费类型，POSTPAID为按量计费，"PREPAID"为预付费包年包月
	LoadBalancerChargeType *string `json:"LoadBalancerChargeType,omitnil,omitempty" name:"LoadBalancerChargeType"`

	// 询价的收费周期。（仅包年包月支持该参数）
	LoadBalancerChargePrepaid *LBChargePrepaid `json:"LoadBalancerChargePrepaid,omitnil,omitempty" name:"LoadBalancerChargePrepaid"`

	// 询价的网络计费方式
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 询价的负载均衡实例个数，默认为1
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 指定可用区询价。如：ap-guangzhou-1
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 包年包月询价时传性能容量型规格，如：<li>clb.c2.medium（标准型）</li><li>clb.c3.small（高阶型1）</li><li>clb.c3.medium（高阶型2）</li>
	// <li>clb.c4.small（超强型1）</li><li>clb.c4.medium（超强型2）</li><li>clb.c4.large（超强型3）</li><li>clb.c4.xlarge（超强型4）</li>
	// 按量付费询价时传SLA
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`

	// IP版本，可取值：IPV4、IPV6、IPv6FullChain，不区分大小写，默认值 IPV4。说明：取值为IPV6表示为IPV6 NAT64版本；取值为IPv6FullChain，表示为IPv6版本。
	AddressIPVersion *string `json:"AddressIPVersion,omitnil,omitempty" name:"AddressIPVersion"`

	// 仅适用于公网负载均衡。目前仅广州、上海、南京、济南、杭州、福州、北京、石家庄、武汉、长沙、成都、重庆地域支持静态单线 IP 线路类型，如需体验，请联系商务经理申请。申请通过后，即可选择中国移动（CMCC）、中国联通（CUCC）或中国电信（CTCC）的运营商类型，网络计费模式只能使用按带宽包计费(BANDWIDTH_PACKAGE)。 如果不指定本参数，则默认使用BGP。可通过 DescribeResources 接口查询一个地域所支持的Isp。
	VipIsp *string `json:"VipIsp,omitnil,omitempty" name:"VipIsp"`
}

type InquiryPriceCreateLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 询价的负载均衡类型，OPEN为公网类型，INTERNAL为内网类型
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// 询价的收费类型，POSTPAID为按量计费，"PREPAID"为预付费包年包月
	LoadBalancerChargeType *string `json:"LoadBalancerChargeType,omitnil,omitempty" name:"LoadBalancerChargeType"`

	// 询价的收费周期。（仅包年包月支持该参数）
	LoadBalancerChargePrepaid *LBChargePrepaid `json:"LoadBalancerChargePrepaid,omitnil,omitempty" name:"LoadBalancerChargePrepaid"`

	// 询价的网络计费方式
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 询价的负载均衡实例个数，默认为1
	GoodsNum *uint64 `json:"GoodsNum,omitnil,omitempty" name:"GoodsNum"`

	// 指定可用区询价。如：ap-guangzhou-1
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 包年包月询价时传性能容量型规格，如：<li>clb.c2.medium（标准型）</li><li>clb.c3.small（高阶型1）</li><li>clb.c3.medium（高阶型2）</li>
	// <li>clb.c4.small（超强型1）</li><li>clb.c4.medium（超强型2）</li><li>clb.c4.large（超强型3）</li><li>clb.c4.xlarge（超强型4）</li>
	// 按量付费询价时传SLA
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`

	// IP版本，可取值：IPV4、IPV6、IPv6FullChain，不区分大小写，默认值 IPV4。说明：取值为IPV6表示为IPV6 NAT64版本；取值为IPv6FullChain，表示为IPv6版本。
	AddressIPVersion *string `json:"AddressIPVersion,omitnil,omitempty" name:"AddressIPVersion"`

	// 仅适用于公网负载均衡。目前仅广州、上海、南京、济南、杭州、福州、北京、石家庄、武汉、长沙、成都、重庆地域支持静态单线 IP 线路类型，如需体验，请联系商务经理申请。申请通过后，即可选择中国移动（CMCC）、中国联通（CUCC）或中国电信（CTCC）的运营商类型，网络计费模式只能使用按带宽包计费(BANDWIDTH_PACKAGE)。 如果不指定本参数，则默认使用BGP。可通过 DescribeResources 接口查询一个地域所支持的Isp。
	VipIsp *string `json:"VipIsp,omitnil,omitempty" name:"VipIsp"`
}

func (r *InquiryPriceCreateLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerType")
	delete(f, "LoadBalancerChargeType")
	delete(f, "LoadBalancerChargePrepaid")
	delete(f, "InternetAccessible")
	delete(f, "GoodsNum")
	delete(f, "ZoneId")
	delete(f, "SlaType")
	delete(f, "AddressIPVersion")
	delete(f, "VipIsp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceCreateLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceCreateLoadBalancerResponseParams struct {
	// 该参数表示对应的价格。
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceCreateLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceCreateLoadBalancerResponseParams `json:"Response"`
}

func (r *InquiryPriceCreateLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceCreateLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceModifyLoadBalancerRequestParams struct {
	// 负载均衡实例 ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 修改后的网络带宽信息
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`
}

type InquiryPriceModifyLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 修改后的网络带宽信息
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`
}

func (r *InquiryPriceModifyLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceModifyLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "InternetAccessible")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceModifyLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceModifyLoadBalancerResponseParams struct {
	// 描述价格信息
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceModifyLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceModifyLoadBalancerResponseParams `json:"Response"`
}

func (r *InquiryPriceModifyLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceModifyLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRefundLoadBalancerRequestParams struct {
	// 负载均衡实例ID。可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`
}

type InquiryPriceRefundLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`
}

func (r *InquiryPriceRefundLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRefundLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceRefundLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRefundLoadBalancerResponseParams struct {
	// 该参数表示对应的价格。
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceRefundLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceRefundLoadBalancerResponseParams `json:"Response"`
}

func (r *InquiryPriceRefundLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRefundLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewLoadBalancerRequestParams struct {
	// 负载均衡实例 ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 续费周期
	LoadBalancerChargePrepaid *LBChargePrepaid `json:"LoadBalancerChargePrepaid,omitnil,omitempty" name:"LoadBalancerChargePrepaid"`
}

type InquiryPriceRenewLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 续费周期
	LoadBalancerChargePrepaid *LBChargePrepaid `json:"LoadBalancerChargePrepaid,omitnil,omitempty" name:"LoadBalancerChargePrepaid"`
}

func (r *InquiryPriceRenewLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "LoadBalancerChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquiryPriceRenewLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquiryPriceRenewLoadBalancerResponseParams struct {
	// 表示续费价格
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquiryPriceRenewLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *InquiryPriceRenewLoadBalancerResponseParams `json:"Response"`
}

func (r *InquiryPriceRenewLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquiryPriceRenewLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IntentRouterItem struct {
	// <p>创建时间（ISO 8601格式）。</p>
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// <p>意图路由ID（ir-xxx格式）。</p>
	IntentRouterId *string `json:"IntentRouterId,omitnil,omitempty" name:"IntentRouterId"`

	// <p>路由名称（例如 IntentRouter/customer-support）。</p>
	RouteName *string `json:"RouteName,omitnil,omitempty" name:"RouteName"`

	// <p>意图路由描述。</p>
	RouterDescribe *string `json:"RouterDescribe,omitnil,omitempty" name:"RouterDescribe"`

	// <p>状态。</p><p>枚举值：</p><ul><li>Provisioning：创建中</li><li>Active：正常</li><li>Configuring：配置中</li><li>ConfigureFailed：配置失败</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>分层配置列表。</p>
	Tiers []*IntentRouterTierItem `json:"Tiers,omitnil,omitempty" name:"Tiers"`

	// <p>更新时间（ISO 8601格式）。</p>
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`
}

type IntentRouterTierDictItem struct {
	// <p>Tier 标识</p><p>枚举值：</p><ul><li>default： 默认</li><li>general_chat： 通用对话</li><li>transformation_rewrite： 转换与改写</li><li>knowledge_qa： 知识问答</li><li>summarization： 摘要</li><li>extraction_structuring： 抽取与结构化输出</li><li>content_generation： 内容生成</li><li>coding_technical： 编码与技术</li><li>data_math_analysis： 数据、数学与分析</li><li>reasoning_planning： 推理与规划</li><li>tool_agentic_workflow： 工具与智能体工作流</li></ul>
	TierId *string `json:"TierId,omitnil,omitempty" name:"TierId"`

	// <p>Tier 显示名称（已国际化）</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>Tier 描述（已国际化）</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type IntentRouterTierItem struct {
	// <p>该分层下的模型显示名称列表。</p>
	Models []*string `json:"Models,omitnil,omitempty" name:"Models"`

	// <p>Tier 标识。<br>枚举值：</p><ul><li>复杂度分层（4 个固定值，需全部包含）：SIMPLE、MEDIUM、COMPLEX、REASONING</li><li>default： 默认</li><li>general_chat： 通用对话</li><li>transformation_rewrite： 转换与改写</li><li>knowledge_qa： 知识问答</li><li>summarization： 摘要</li><li>extraction_structuring： 抽取与结构化输出</li><li>content_generation： 内容生成</li><li>coding_technical： 编码与技术</li><li>data_math_analysis： 数据、数学与分析</li><li>reasoning_planning： 推理与规划</li><li>tool_agentic_workflow： 工具与智能体工作流</li></ul>
	TierName *string `json:"TierName,omitnil,omitempty" name:"TierName"`
}

type InternetAccessible struct {
	// TRAFFIC_POSTPAID_BY_HOUR 按流量按小时后计费 ; BANDWIDTH_POSTPAID_BY_HOUR 按带宽按小时后计费，国际站用户不支持该计费模式; BANDWIDTH_PACKAGE 按带宽包计费;BANDWIDTH_PREPAID按带宽预付费。
	InternetChargeType *string `json:"InternetChargeType,omitnil,omitempty" name:"InternetChargeType"`

	// 最大出带宽，单位Mbps，仅对公网属性的共享型、性能容量型和独占型 CLB 实例、以及内网属性的性能容量型 CLB 实例生效。
	// - 对于公网属性的共享型和独占型 CLB 实例，最大出带宽的范围为1Mbps-2048Mbps。
	// - 对于公网属性和内网属性的性能容量型 CLB实例，最大出带宽的范围为1Mbps-61440Mbps。
	// （调用CreateLoadBalancer创建LB时不指定此参数则设置为默认值10Mbps。此上限可调整）
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 带宽包的类型，如 BGP（多线）。
	// 类型如下：
	// SINGLEISP: 单线
	// BGP: 多线
	// HIGH_QUALITY_BGP: 精品BGP共享带宽包
	// SINGLEISP_CMCC: 中国移动共享带宽包
	// SINGLEISP_CTCC: 中国电信共享带宽包
	// SINGLEISP_CUCC: 中国联通共享带宽包
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthpkgSubType *string `json:"BandwidthpkgSubType,omitnil,omitempty" name:"BandwidthpkgSubType"`
}

type ItemPrice struct {
	// 后付费单价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPrice *float64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// 后续计价单元，可取值范围： 
	// HOUR：表示计价单元是按每小时来计算。当前涉及该计价单元的场景有：实例按小时后付费（POSTPAID_BY_HOUR）、带宽按小时后付费（BANDWIDTH_POSTPAID_BY_HOUR）；
	// GB：表示计价单元是按每GB来计算。当前涉及该计价单元的场景有：流量按小时后付费（TRAFFIC_POSTPAID_BY_HOUR）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeUnit *string `json:"ChargeUnit,omitnil,omitempty" name:"ChargeUnit"`

	// 预支费用的原价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 预支费用的折扣价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// 后付费的折扣单价，单位:元
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitnil,omitempty" name:"UnitPriceDiscount"`

	// 折扣，如20.0代表2折。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`
}

type Job struct {
	// <p>接口名称</p>
	ApiName *string `json:"ApiName,omitnil,omitempty" name:"ApiName"`

	// <p>请求ID</p>
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// <p>异步任务状态</p><p>枚举值：</p><ul><li>Processing： 进行中</li><li>Succeeded： 成功</li><li>Failed： 失败</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>资源ID</p>
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`
}

type KeyDetailItem struct {
	// Key 业务 ID
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// Key 创建时间（ISO 8601）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// Key 显示名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type KeyInfo struct {
	// <p>是否禁用Key</p>
	Blocked *bool `json:"Blocked,omitnil,omitempty" name:"Blocked"`

	// <p>Key关联的Budget ID。</p><p>未关联Budget时返回空字符串。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>Key关联的Budget名称。</p><p>未关联Budget时返回空字符串。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BudgetName *string `json:"BudgetName,omitnil,omitempty" name:"BudgetName"`

	// <p>创建时间</p>
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// <p>Key按Budget刷新周期划分的Credit使用情况。</p><p>当关联Budget配置多个刷新周期时，按1d、7d、30d顺序返回各周期用量；未关联Budget时返回空数组。</p>
	CreditUsageSet []*CreditUsage `json:"CreditUsageSet,omitnil,omitempty" name:"CreditUsageSet"`

	// <p>Key的值</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>Key的ID</p>
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// <p>Key名称</p>
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// <p>修改时间</p>
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// <p>限速信息</p>
	RateLimitConfig *RateLimitConfigForKey `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`

	// <p>Key状态</p><p>枚举值：</p><ul><li>Active： 正常</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>标签</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>所属的用户组ID</p>
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// <p>所属的用户组名称</p>
	UserGroupName *string `json:"UserGroupName,omitnil,omitempty" name:"UserGroupName"`
}

type KeyItem struct {
	// Provider API Key
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`

	// Key 标识名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type LBChargePrepaid struct {
	// 续费类型：AUTO_RENEW 自动续费，  MANUAL_RENEW 手动续费
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 购买时长，单位：月
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`
}

type LBItem struct {
	// lb的字符串id
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// lb的vip
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 监听器规则
	Listeners []*ListenerItem `json:"Listeners,omitnil,omitempty" name:"Listeners"`

	// LB所在地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type LBOperateProtectInfo struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 保护状态，true：表示开启了操作保护，false：表示未开启操作保护。
	ProtectState *bool `json:"ProtectState,omitnil,omitempty" name:"ProtectState"`

	// 操作保护的设置uin。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorUin *string `json:"OperatorUin,omitnil,omitempty" name:"OperatorUin"`

	// 设置操作保护时的描述信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 最后修改时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`
}

type LbRsItem struct {
	// vpc的字符串id，只支持字符串id。
	// 可以通过 [DescribeVpcs](https://cloud.tencent.com/document/api/215/15778) 接口查询。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 需要查询后端的内网 IP，可以是 CVM 和弹性网卡。
	// 可以通过 [DescribeAddresses](https://cloud.tencent.com/document/product/215/16702) 接口查询。
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`
}

type LbRsTargets struct {
	// 内网ip类型。“cvm”或“eni”
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 后端实例的内网ip。
	PrivateIp *string `json:"PrivateIp,omitnil,omitempty" name:"PrivateIp"`

	// 绑定后端实例的端口。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// rs的vpcId
	VpcId *int64 `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// rs的权重
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type Listener struct {
	// <p>负载均衡监听器 ID</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>监听器协议，可选值：TCP、UDP、HTTP、HTTPS、TCP_SSL、QUIC</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>监听器端口，端口范围：1-65535</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>监听器绑定的证书信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Certificate *CertificateOutput `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// <p>监听器的健康检查信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheck *HealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// <p>请求的调度方式。 WRR、LEAST_CONN、IP_HASH分别表示按权重轮询、最小连接数、IP Hash。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// <p>会话保持时间，单位：秒。可选值：30~3600，默认 0，默认不开启。此参数仅适用于TCP/UDP监听器。</p><p>单位：秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionExpireTime *int64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// <p>是否开启SNI特性，1：表示开启，0：表示不开启（本参数仅对于HTTPS监听器有意义）</p>
	SniSwitch *int64 `json:"SniSwitch,omitnil,omitempty" name:"SniSwitch"`

	// <p>监听器下的全部转发规则（本参数仅对于HTTP/HTTPS监听器有意义）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*RuleOutput `json:"Rules,omitnil,omitempty" name:"Rules"`

	// <p>监听器的名称</p>
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// <p>监听器的创建时间。</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>端口段结束端口，端口范围：2-65535</p>
	EndPort *int64 `json:"EndPort,omitnil,omitempty" name:"EndPort"`

	// <p>后端服务器类型，可选值：NODE、POLARIS、TARGETGROUP、TARGETGROUP-V2</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// <p>绑定的目标组基本信息；当监听器绑定目标组时，会返回该字段</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetGroup *BasicTargetGroupInfo `json:"TargetGroup,omitnil,omitempty" name:"TargetGroup"`

	// <p>会话保持类型。NORMAL表示默认会话保持类型。QUIC_CID 表示根据Quic Connection ID做会话保持。</p>
	SessionType *string `json:"SessionType,omitnil,omitempty" name:"SessionType"`

	// <p>是否开启长连接，1开启，0关闭，（本参数仅对于HTTP/HTTPS监听器有意义）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeepaliveEnable *int64 `json:"KeepaliveEnable,omitnil,omitempty" name:"KeepaliveEnable"`

	// <p>仅支持Nat64 CLB TCP监听器</p>
	Toa *bool `json:"Toa,omitnil,omitempty" name:"Toa"`

	// <p>重新调度功能，解绑后端服务开关，打开此开关，当解绑后端服务时触发重新调度。仅TCP/UDP监听器支持。</p>
	DeregisterTargetRst *bool `json:"DeregisterTargetRst,omitnil,omitempty" name:"DeregisterTargetRst"`

	// <p>监听器的属性</p>
	AttrFlags []*string `json:"AttrFlags,omitnil,omitempty" name:"AttrFlags"`

	// <p>绑定的目标组列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetGroupList []*BasicTargetGroupInfo `json:"TargetGroupList,omitnil,omitempty" name:"TargetGroupList"`

	// <p>监听器最大连接数，-1表示监听器维度不限速。</p>
	MaxConn *int64 `json:"MaxConn,omitnil,omitempty" name:"MaxConn"`

	// <p>监听器最大新增连接数，-1表示监听器维度不限速。</p>
	MaxCps *int64 `json:"MaxCps,omitnil,omitempty" name:"MaxCps"`

	// <p>空闲连接超时时间，仅支持TCP监听器。默认值:900；共享型实例和独占型实例取值范围：300～900，性能容量型实例取值范围:300～1980。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdleConnectTimeout *int64 `json:"IdleConnectTimeout,omitnil,omitempty" name:"IdleConnectTimeout"`

	// <p>重新调度触发持续时间，取值0~3600s。仅TCP/UDP监听器支持。触发重新调度后，长连接将会在设置的调度时间内断开并完成重新分配。</p><p>单位：秒</p>
	RescheduleInterval *uint64 `json:"RescheduleInterval,omitnil,omitempty" name:"RescheduleInterval"`

	// <p>数据压缩模式</p>
	DataCompressMode *string `json:"DataCompressMode,omitnil,omitempty" name:"DataCompressMode"`

	// <p>重新调度启动时间，配置了重新调度启动时间后，会在启动时间到达时触发重新调度。</p>
	RescheduleStartTime *int64 `json:"RescheduleStartTime,omitnil,omitempty" name:"RescheduleStartTime"`
}

type ListenerBackend struct {
	// 监听器 ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 监听器的协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 监听器的端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 监听器下的规则信息（仅适用于HTTP/HTTPS监听器）
	Rules []*RuleTargets `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 监听器上绑定的后端服务列表（仅适用于TCP/UDP/TCP_SSL监听器）
	Targets []*Backend `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 若支持端口段，则为端口段结束端口；若不支持端口段，则为0
	EndPort *int64 `json:"EndPort,omitnil,omitempty" name:"EndPort"`
}

type ListenerHealth struct {
	// 监听器ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 监听器名称
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// 监听器的协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 监听器的端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 监听器的转发规则列表
	Rules []*RuleHealth `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type ListenerItem struct {
	// 监听器ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 监听器协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 监听器端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 绑定规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*RulesItems `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 四层绑定对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Targets []*LbRsTargets `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 端口段监听器的结束端口
	EndPort *int64 `json:"EndPort,omitnil,omitempty" name:"EndPort"`
}

type LoadBalancer struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡实例的名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 负载均衡实例的网络类型：
	// OPEN：公网属性， INTERNAL：内网属性；对于内网属性的负载均衡，可通过绑定EIP出公网，具体可参考EIP文档[绑定弹性公网IP](https://cloud.tencent.com/document/product/215/16700)。
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// 负载均衡类型标识，1：负载均衡，0：传统型负载均衡。
	Forward *uint64 `json:"Forward,omitnil,omitempty" name:"Forward"`

	// 负载均衡实例的域名，仅公网传统型和域名型负载均衡实例才提供该字段。逐步下线中，建议用LoadBalancerDomain替代。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 负载均衡实例的 VIP 列表。
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitnil,omitempty" name:"LoadBalancerVips"`

	// 负载均衡实例的状态，包括
	// 0：创建中，1：正常运行。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 负载均衡实例的创建时间。
	// 格式：YYYY-MM-DD HH:mm:ss
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 负载均衡实例的上次状态转换时间。
	// 格式：YYYY-MM-DD HH:mm:ss
	StatusTime *string `json:"StatusTime,omitnil,omitempty" name:"StatusTime"`

	// 负载均衡实例所属的项目 ID， 0 表示默认项目。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 私有网络的 ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 高防 LB 的标识，1：高防负载均衡 0：非高防负载均衡。
	OpenBgp *uint64 `json:"OpenBgp,omitnil,omitempty" name:"OpenBgp"`

	// 是否开启 SNAT，在 2016 年 12 月份之前的传统型内网负载均衡都是开启了 SNAT 的。
	Snat *bool `json:"Snat,omitnil,omitempty" name:"Snat"`

	// 是否被隔离，0：表示未被隔离，1：表示被隔离。
	Isolation *uint64 `json:"Isolation,omitnil,omitempty" name:"Isolation"`

	// 用户开启日志的信息，日志只有公网属性创建了 HTTP 、HTTPS 监听器的负载均衡才会有日志。
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Log is deprecated.
	Log *string `json:"Log,omitnil,omitempty" name:"Log"`

	// 负载均衡实例所在的子网（仅对内网VPC型LB有意义）
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 负载均衡实例的标签信息
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 负载均衡实例的安全组
	SecureGroups []*string `json:"SecureGroups,omitnil,omitempty" name:"SecureGroups"`

	// 负载均衡实例绑定的后端设备的基本信息
	TargetRegionInfo *TargetRegionInfo `json:"TargetRegionInfo,omitnil,omitempty" name:"TargetRegionInfo"`

	// anycast负载均衡的发布域，对于非anycast的负载均衡，此字段返回为空字符串
	AnycastZone *string `json:"AnycastZone,omitnil,omitempty" name:"AnycastZone"`

	// IP版本，ipv4 | ipv6
	AddressIPVersion *string `json:"AddressIPVersion,omitnil,omitempty" name:"AddressIPVersion"`

	// 数值形式的私有网络 ID，可以通过[DescribeVpcs](https://cloud.tencent.com/document/product/215/15778)接口获取。
	NumericalVpcId *uint64 `json:"NumericalVpcId,omitnil,omitempty" name:"NumericalVpcId"`

	// 负载均衡IP地址所属的运营商。
	// 
	// - BGP :  BGP（多线）
	// - CMCC：中国移动单线
	// - CTCC：中国电信单线
	// - CUCC：中国联通单线
	// 注意：此字段可能返回 null，表示取不到有效值。
	VipIsp *string `json:"VipIsp,omitnil,omitempty" name:"VipIsp"`

	// 主可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterZone *ZoneInfo `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// 备可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupZoneSet []*ZoneInfo `json:"BackupZoneSet,omitnil,omitempty" name:"BackupZoneSet"`

	// 负载均衡实例被隔离的时间。
	// 格式：YYYY-MM-DD HH:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolatedTime *string `json:"IsolatedTime,omitnil,omitempty" name:"IsolatedTime"`

	// 负载均衡实例的过期时间，仅对预付费负载均衡生效。
	// 格式：YYYY-MM-DD HH:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 负载均衡实例的计费类型，PREPAID：包年包月，POSTPAID_BY_HOUR：按量计费
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 负载均衡实例的网络属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkAttributes *InternetAccessible `json:"NetworkAttributes,omitnil,omitempty" name:"NetworkAttributes"`

	// 负载均衡实例的预付费相关属性，仅在 ChargeType=PREPAID 时显示。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrepaidAttributes *LBChargePrepaid `json:"PrepaidAttributes,omitnil,omitempty" name:"PrepaidAttributes"`

	// 负载均衡日志服务(CLS)的日志集ID
	LogSetId *string `json:"LogSetId,omitnil,omitempty" name:"LogSetId"`

	// 负载均衡日志服务(CLS)的日志主题ID
	LogTopicId *string `json:"LogTopicId,omitnil,omitempty" name:"LogTopicId"`

	// 负载均衡实例的IPv6地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressIPv6 *string `json:"AddressIPv6,omitnil,omitempty" name:"AddressIPv6"`

	// 暂做保留，一般用户无需关注。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo *ExtraInfo `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 是否可绑定高防包
	IsDDos *bool `json:"IsDDos,omitnil,omitempty" name:"IsDDos"`

	// 负载均衡维度的个性化配置ID
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// 后端服务是否放通来自LB的流量
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitnil,omitempty" name:"LoadBalancerPassToTarget"`

	// 内网独占集群
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitnil,omitempty" name:"ExclusiveCluster"`

	// IP地址版本为ipv6时此字段有意义，IPv6Nat64 | IPv6FullChain。
	// IPv6Nat64: 基于 NAT64 IPv6 过渡技术实现的负载均衡器。
	// IPv6FullChain：基于 IPv6 单栈技术实现的负载均衡。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPv6Mode *string `json:"IPv6Mode,omitnil,omitempty" name:"IPv6Mode"`

	// 是否开启SnatPro。
	SnatPro *bool `json:"SnatPro,omitnil,omitempty" name:"SnatPro"`

	// 开启SnatPro负载均衡后，SnatIp列表。
	SnatIps []*SnatIp `json:"SnatIps,omitnil,omitempty" name:"SnatIps"`

	// 性能容量型规格。<ul><li> clb.c1.small：简约型规格 </li><li> clb.c2.medium：标准型规格 </li><li> clb.c3.small：高阶型1规格 </li><li> clb.c3.medium：高阶型2规格 </li><li> clb.c4.small：超强型1规格 </li><li> clb.c4.medium：超强型2规格 </li><li> clb.c4.large：超强型3规格 </li><li> clb.c4.xlarge：超强型4规格 </li><li>""：非性能容量型实例</li></ul>
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`

	// vip是否被封堵
	IsBlock *bool `json:"IsBlock,omitnil,omitempty" name:"IsBlock"`

	// 封堵或解封时间。
	// 格式：YYYY-MM-DD HH:mm:ss。
	IsBlockTime *string `json:"IsBlockTime,omitnil,omitempty" name:"IsBlockTime"`

	// IP类型是否是本地BGP
	LocalBgp *bool `json:"LocalBgp,omitnil,omitempty" name:"LocalBgp"`

	// 7层独占标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterTag *string `json:"ClusterTag,omitnil,omitempty" name:"ClusterTag"`

	// 开启IPv6FullChain负载均衡7层监听器支持混绑IPv4/IPv6目标功能。
	MixIpTarget *bool `json:"MixIpTarget,omitnil,omitempty" name:"MixIpTarget"`

	// 私有网络内网负载均衡，就近接入模式下规则所落在的可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// CLB是否为NFV，空：不是，l7nfv：七层是NFV。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NfvInfo *string `json:"NfvInfo,omitnil,omitempty" name:"NfvInfo"`

	// 负载均衡日志服务(CLS)的健康检查日志集ID
	HealthLogSetId *string `json:"HealthLogSetId,omitnil,omitempty" name:"HealthLogSetId"`

	// 负载均衡日志服务(CLS)的健康检查日志主题ID
	HealthLogTopicId *string `json:"HealthLogTopicId,omitnil,omitempty" name:"HealthLogTopicId"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterIds []*string `json:"ClusterIds,omitnil,omitempty" name:"ClusterIds"`

	// 负载均衡的属性，传入字符串数组来决定是否开启
	// DeleteProtect: 删除保护，开启后防止负载均衡被误删除。 
	// UserInVisible: 用户不可见，控制负载均衡对用户的可见性。 
	// BlockStatus: 阻塞状态，用于限制负载均衡的某些操作或流量。 
	// NoLBNat: 禁用负载均衡的NAT功能，用于特定场景下的流量直接转发。 
	// BanStatus: 封禁状态，用于暂停负载均衡服务或限制访问。 
	// ShiftupFlag: 升配标志，用于标识负载均衡需要升级配置或性能。 
	// Stop: 停止状态，开启后负载均衡暂停服务。 
	// NoVpcGw: 不使用VPC网关，用于绕过VPC网关直接处理流量。 
	// SgInTgw: 安全组在TGW（Transit Gateway）中，涉及网络安全策略配置。 
	// SharedLimitFlag: 共享限制标志，用于控制负载均衡的共享资源限制。 
	// WafFlag: Web应用防火墙（WAF）标志，开启后启用WAF保护。 
	// IsDomainCLB: 域名型负载均衡，标识负载均衡是否基于域名进行流量分发。 
	// IPv6Snat: IPv6源地址转换（SNAT），用于IPv6网络的源地址处理。 
	// HideDomain: 隐藏域名，用于隐私保护或特定场景下不暴露域名。 
	// JumboFrame: 巨型帧支持，开启后支持更大的数据帧以提高网络效率。 
	// NoLBNatL4IPdc: 四层IP直连无NAT，用于四层负载均衡直接转发IP流量。 
	// VpcGwL3Service: VPC网关三层服务，涉及三层网络服务的网关功能。 
	// Ipv62Flag: IPv6扩展标志，用于特定的IPv6功能支持。 
	// Ipv62ExclusiveFlag: IPv6独占标志，用于专属IPv6流量处理。 
	// BgpPro: BGP Pro 支持。 
	// ToaClean: TOA（TCP Option Address）清理，清除TCP选项中的地址信息。 
	AttributeFlags []*string `json:"AttributeFlags,omitnil,omitempty" name:"AttributeFlags"`

	// 负载均衡实例的域名。
	LoadBalancerDomain *string `json:"LoadBalancerDomain,omitnil,omitempty" name:"LoadBalancerDomain"`

	// 网络出口
	Egress *string `json:"Egress,omitnil,omitempty" name:"Egress"`

	// 实例类型是否为独占型。1：独占型实例。0：非独占型实例。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Exclusive *uint64 `json:"Exclusive,omitnil,omitempty" name:"Exclusive"`

	// 已绑定的后端服务数量。
	TargetCount *uint64 `json:"TargetCount,omitnil,omitempty" name:"TargetCount"`

	// 负载均衡实例关联的Endpoint id。
	AssociateEndpoint *string `json:"AssociateEndpoint,omitnil,omitempty" name:"AssociateEndpoint"`

	// 可用区转发亲和信息
	AvailableZoneAffinityInfo *AvailableZoneAffinityInfo `json:"AvailableZoneAffinityInfo,omitnil,omitempty" name:"AvailableZoneAffinityInfo"`
}

type LoadBalancerDetail struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡实例的名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 负载均衡实例的网络类型：
	// Public：公网属性，Private：内网属性；对于内网属性的负载均衡，可通过绑定EIP出公网，具体可参考EIP文档。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// 负载均衡实例的状态，包括
	// 0：创建中，1：正常运行。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 负载均衡实例的 VIP 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 负载均衡实例 VIP 的IPv6地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressIPv6 *string `json:"AddressIPv6,omitnil,omitempty" name:"AddressIPv6"`

	// 负载均衡实例IP版本，IPv4 | IPv6。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressIPVersion *string `json:"AddressIPVersion,omitnil,omitempty" name:"AddressIPVersion"`

	// 负载均衡实例IPv6地址类型，IPv6Nat64 | IPv6FullChain。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPv6Mode *string `json:"IPv6Mode,omitnil,omitempty" name:"IPv6Mode"`

	// 负载均衡实例所在可用区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 负载均衡实例IP地址所属的ISP。取值范围：BGP（多线）、CMCC（中国移动）、CUCC（中国联通）、CTCC（中国电信）、INTERNAL（内网）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressIsp *string `json:"AddressIsp,omitnil,omitempty" name:"AddressIsp"`

	// 负载均衡实例所属私有网络的 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 负载均衡实例所属的项目 ID， 0 表示默认项目。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 负载均衡实例的创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 负载均衡实例的计费类型。取值范围：PREPAID预付费、POSTPAID_BY_HOUR按量付费。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 负载均衡实例的网络属性。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkAttributes *InternetAccessible `json:"NetworkAttributes,omitnil,omitempty" name:"NetworkAttributes"`

	// 负载均衡实例的预付费相关属性。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrepaidAttributes *LBChargePrepaid `json:"PrepaidAttributes,omitnil,omitempty" name:"PrepaidAttributes"`

	// 暂做保留，一般用户无需关注。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo *ExtraInfo `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 负载均衡维度的个性化配置ID，多个配置用逗号隔开。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// 负载均衡实例的标签信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 负载均衡监听器 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 监听器协议。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 监听器端口。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 转发规则的 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 转发规则的域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 转发规则的路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 后端目标ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 后端目标的IP地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetAddress *string `json:"TargetAddress,omitnil,omitempty" name:"TargetAddress"`

	// 后端目标监听端口。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetPort *uint64 `json:"TargetPort,omitnil,omitempty" name:"TargetPort"`

	// 后端目标转发权重。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetWeight *uint64 `json:"TargetWeight,omitnil,omitempty" name:"TargetWeight"`

	// 0：表示未被隔离，1：表示被隔离。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Isolation *uint64 `json:"Isolation,omitnil,omitempty" name:"Isolation"`

	// 负载均衡绑定的安全组列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroup []*string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// 负载均衡安全组上移特性是否开启标识。取值范围：1表示开启、0表示未开启。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerPassToTarget *uint64 `json:"LoadBalancerPassToTarget,omitnil,omitempty" name:"LoadBalancerPassToTarget"`

	// 后端目标健康状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetHealth *string `json:"TargetHealth,omitnil,omitempty" name:"TargetHealth"`

	// 转发规则的域名列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domains *string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 多可用区负载均衡实例所选备区
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaveZone []*string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// 内网负载均衡实例所在可用区，由白名单CLB_Internal_Zone控制
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 是否开启SNI特性，1：表示开启，0：表示不开启（本参数仅对于HTTPS监听器有意义）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SniSwitch *int64 `json:"SniSwitch,omitnil,omitempty" name:"SniSwitch"`

	// 负载均衡实例的域名。
	LoadBalancerDomain *string `json:"LoadBalancerDomain,omitnil,omitempty" name:"LoadBalancerDomain"`

	// 网络出口
	Egress *string `json:"Egress,omitnil,omitempty" name:"Egress"`

	// 负载均衡的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttributeFlags []*string `json:"AttributeFlags,omitnil,omitempty" name:"AttributeFlags"`

	// 负载均衡实例的规格类型信息<ul><li> clb.c1.small：简约型规格 </li><li>clb.c2.medium：标准型规格 </li><li> clb.c3.small：高阶型1规格 </li><li> clb.c3.medium：高阶型2规格 </li><li> clb.c4.small：超强型1规格 </li><li> clb.c4.medium：超强型2规格 </li><li> clb.c4.large：超强型3规格 </li><li> clb.c4.xlarge：超强型4规格 </li><li>""：非性能容量型实例</li></ul>
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`

	// 0：表示非独占型实例，1：表示独占型态实例。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Exclusive *uint64 `json:"Exclusive,omitnil,omitempty" name:"Exclusive"`

	// 可用区转发亲和信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailableZoneAffinityInfo *AvailableZoneAffinityInfo `json:"AvailableZoneAffinityInfo,omitnil,omitempty" name:"AvailableZoneAffinityInfo"`
}

type LoadBalancerHealth struct {
	// 负载均衡实例ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡实例名称
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 监听器列表
	Listeners []*ListenerHealth `json:"Listeners,omitnil,omitempty" name:"Listeners"`
}

type LoadBalancerTraffic struct {
	// 负载均衡ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡名字
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 负载均衡所在地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 负载均衡的vip
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 最大出带宽，单位：Mbps
	OutBandwidth *float64 `json:"OutBandwidth,omitnil,omitempty" name:"OutBandwidth"`

	// CLB域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`
}

// Predefined struct for user
type ManualRewriteRequestParams struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 源监听器 ID。
	SourceListenerId *string `json:"SourceListenerId,omitnil,omitempty" name:"SourceListenerId"`

	// 目标监听器 ID。
	TargetListenerId *string `json:"TargetListenerId,omitnil,omitempty" name:"TargetListenerId"`

	// 转发规则之间的重定向关系。
	RewriteInfos []*RewriteLocationMap `json:"RewriteInfos,omitnil,omitempty" name:"RewriteInfos"`
}

type ManualRewriteRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 源监听器 ID。
	SourceListenerId *string `json:"SourceListenerId,omitnil,omitempty" name:"SourceListenerId"`

	// 目标监听器 ID。
	TargetListenerId *string `json:"TargetListenerId,omitnil,omitempty" name:"TargetListenerId"`

	// 转发规则之间的重定向关系。
	RewriteInfos []*RewriteLocationMap `json:"RewriteInfos,omitnil,omitempty" name:"RewriteInfos"`
}

func (r *ManualRewriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManualRewriteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "SourceListenerId")
	delete(f, "TargetListenerId")
	delete(f, "RewriteInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ManualRewriteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ManualRewriteResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ManualRewriteResponse struct {
	*tchttp.BaseResponse
	Response *ManualRewriteResponseParams `json:"Response"`
}

func (r *ManualRewriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ManualRewriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MigrateClassicalLoadBalancersRequestParams struct {
	// 传统型负载均衡ID数组
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 独占集群信息
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitnil,omitempty" name:"ExclusiveCluster"`
}

type MigrateClassicalLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// 传统型负载均衡ID数组
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 独占集群信息
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitnil,omitempty" name:"ExclusiveCluster"`
}

func (r *MigrateClassicalLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigrateClassicalLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	delete(f, "ExclusiveCluster")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MigrateClassicalLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MigrateClassicalLoadBalancersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type MigrateClassicalLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *MigrateClassicalLoadBalancersResponseParams `json:"Response"`
}

func (r *MigrateClassicalLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MigrateClassicalLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModalityProbeDetail struct {
	// <p>探测的模态</p>
	Modality *string `json:"Modality,omitnil,omitempty" name:"Modality"`

	// <p>探测结果</p><p>枚举值：</p><ul><li>Supported： 模型支持该输入模态</li><li>Unsupported： 模型不支持该输入模态</li><li>Inconclusive： 模型未明确是否支持该模态，待重新探测</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>探测该模态请求的报错详情</p>
	ErrorInfo *ProviderTestConnectionErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`
}

type ModelAlias struct {
	// <p>模型积分系数配置，包含 <code>InputCoefficient</code>、<code>InputCachedCoefficient</code> 和 <code>OutputCoefficient</code>。</p><p>未配置时输入系数默认为 25，缓存命中输入系数默认为 3，输出系数默认为 100。</p>
	Coefficient *Coefficient `json:"Coefficient,omitnil,omitempty" name:"Coefficient"`

	// <p>模型别名名称。</p><p>若用户配置了模型别名，则为该别名；未配置时为原始模型名称。</p>
	ModelAliasName *string `json:"ModelAliasName,omitnil,omitempty" name:"ModelAliasName"`

	// <p>该模型别名下各 BYOK 实例（ServiceProvider）的积分系数明细，体现 ModelAlias 与 ServiceProvider 的层级关系。</p><p>默认返回该别名引用的全部实例；某实例返回 <code>Coefficient</code> 表示其单独配置了 ServiceProvider 维度系数，否则继承顶层 ModelAlias 的 <code>Coefficient</code>。</p><p>该别名当前无有效 BYOK 引用时返回空数组。</p>
	ServiceProviderCoefficientSet []*ServiceProviderCoefficient `json:"ServiceProviderCoefficientSet,omitnil,omitempty" name:"ServiceProviderCoefficientSet"`

	// <p>模型来源。</p><p>枚举值：</p><ul><li>BYOK：用户 BYOK 配置的模型。</li></ul>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>状态</p><p>枚举值：</p><ul><li>Active： 正常可用</li><li>Configuring： 变配中</li><li>ConfigureFailed： 变配失败</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModelAssociation struct {
	// <p>该模型最大可支持的输入多模态能力列表</p><p>枚举值：</p><ul><li>text： 支持文本输入</li><li>image： 支持图像输入</li><li>file： 支持文件输入（当前仅支持pdf）</li></ul>
	InputModalitiesUnion []*string `json:"InputModalitiesUnion,omitnil,omitempty" name:"InputModalitiesUnion"`

	// <p>模型名称</p>
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// <p>BYOK列表</p>
	ServiceProviders []*ServiceProvider `json:"ServiceProviders,omitnil,omitempty" name:"ServiceProviders"`

	// <p>模型类型</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ModelAvailability struct {
	// <p>该模型所有健康BYOK实例下支持的输入多模态能力的并集。模型不健康时返回空列表。</p><p>枚举值：</p><ul><li>text： 支持文本输入</li><li>image： 支持图像输入</li><li>file： 支持文件输入（当前仅支持pdf）</li></ul>
	InputModalities []*string `json:"InputModalities,omitnil,omitempty" name:"InputModalities"`

	// <p>模型</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>可用性状态</p><p>枚举值：</p><ul><li>Available： 可用</li><li>Unavailable： 不可用</li><li>Unknown： 未探测</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModelHealthCheckResults struct {
	// <p>BYOK的KeyID</p>
	ProviderKeyId *string `json:"ProviderKeyId,omitnil,omitempty" name:"ProviderKeyId"`

	// <p>模型</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>健康检查状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModelItem struct {
	// <p>模型唯一标识, 用于实际访问</p>
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// <p>该模型当前支持的输入多模态能力列表</p><p>枚举值：</p><ul><li>text： 支持文本输入</li><li>image： 支持图像输入</li><li>file： 支持文件输入（当前仅支持pdf）</li></ul><p>默认值：text</p>
	InputModalities []*string `json:"InputModalities,omitnil,omitempty" name:"InputModalities"`

	// <p>模型别名, 可以用于实际访问</p>
	ModelAlias *string `json:"ModelAlias,omitnil,omitempty" name:"ModelAlias"`
}

type ModelKeyInfoItem struct {
	// <p>接入类型</p>
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// <p>API Base URL</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiBase *string `json:"ApiBase,omitnil,omitempty" name:"ApiBase"`

	// <p>模型创建时间（ISO 8601）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// <p>自定义host header</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`

	// <p>Key 数量</p>
	KeyCount *int64 `json:"KeyCount,omitnil,omitempty" name:"KeyCount"`

	// <p>Key 详情列表</p>
	Keys []*KeyDetailItem `json:"Keys,omitnil,omitempty" name:"Keys"`

	// <p>model信息</p>
	ModelIdsWithAlias []*ServiceProviderModelItem `json:"ModelIdsWithAlias,omitnil,omitempty" name:"ModelIdsWithAlias"`

	// <p>模型供应商</p>
	ModelProvider *string `json:"ModelProvider,omitnil,omitempty" name:"ModelProvider"`

	// <p>模型协议</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>内部通信占用IP</p>
	ServiceIps []*string `json:"ServiceIps,omitnil,omitempty" name:"ServiceIps"`

	// <p>服务提供商ID</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>服务提供商自定义名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceProviderName *string `json:"ServiceProviderName,omitnil,omitempty" name:"ServiceProviderName"`

	// <p>模型状态</p><p>枚举值：</p><ul><li>Active： 运行中</li><li>Provisioning： 创建中</li><li>Configuring： 变配中</li><li>Deleting： 删除中</li><li>ProvisionFailed： 创建失败</li><li>ConfigureFailed： 变配失败</li><li>DeletionFailed： 删除失败</li><li>Disabled： 已禁用</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>子网 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>标签信息</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>是否校验上游SSL</p>
	VerifySSL *bool `json:"VerifySSL,omitnil,omitempty" name:"VerifySSL"`

	// <p>VPC 实例 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type ModelNameAggregatedItem struct {
	// <p>模型标识显示名称（优先使用 model_alias，否则使用 model_name）</p>
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// <p>关联的服务商列表</p>
	ServiceProviders []*ServiceProviderItem `json:"ServiceProviders,omitnil,omitempty" name:"ServiceProviders"`

	// <p>该模型最大可支持的输入多模态能力列表。</p><p>枚举值：</p><ul><li>text： 支持文本输入</li><li>image： 支持图像输入</li><li>file： 支持文件输入（当前仅支持pdf）</li></ul>
	InputModalitiesUnion []*string `json:"InputModalitiesUnion,omitnil,omitempty" name:"InputModalitiesUnion"`
}

type ModelRouterDetail struct {
	// <p>模型路由实例关联的Budget ID。</p><p>未关联Budget时返回空字符串。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>模型路由实例关联的Budget名称。</p><p>未关联Budget时返回空字符串。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BudgetName *string `json:"BudgetName,omitnil,omitempty" name:"BudgetName"`

	// <p>集群信息</p>
	ClusterInfo *ClusterInfo `json:"ClusterInfo,omitnil,omitempty" name:"ClusterInfo"`

	// <p>创建时间</p>
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// <p>模型路由实例按Budget刷新周期划分的Credit使用情况。</p><p>当关联Budget配置多个刷新周期时，按1d、7d、30d顺序返回各周期用量；未关联Budget时返回空数组。</p>
	CreditUsageSet []*CreditUsage `json:"CreditUsageSet,omitnil,omitempty" name:"CreditUsageSet"`

	// <p>模型路由实例域名</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>模型路由ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>模型路由名称</p><p>默认值：-</p>
	ModelRouterName *string `json:"ModelRouterName,omitnil,omitempty" name:"ModelRouterName"`

	// <p>模型路由类型</p><p>枚举值：</p><ul><li>Shared： 共享型</li><li>Enterprise： 企业级</li></ul>
	ModelRouterType *string `json:"ModelRouterType,omitnil,omitempty" name:"ModelRouterType"`

	// <p>修改时间</p>
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// <p>模型路由实例网络类型</p><p>枚举值：</p><ul><li>Internet： 公网</li><li>Intranet： 内网</li></ul>
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// <p>模型路由限速信息</p>
	RateLimitConfig *RateLimitConfigForModelRouter `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`

	// <p>模型路由的路由配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouterSetting *RouterSettingWithFallBack `json:"RouterSetting,omitnil,omitempty" name:"RouterSetting"`

	// <p>安全组ID列表</p>
	SecurityGroups []*string `json:"SecurityGroups,omitnil,omitempty" name:"SecurityGroups"`

	// <p>模型路由实例的安全状态</p><p>枚举值：</p><ul><li>Normal： 正常</li><li>Banned： 已封禁</li><li>Frozen： 已冻结</li></ul>
	SecurityStatus *string `json:"SecurityStatus,omitnil,omitempty" name:"SecurityStatus"`

	// <p>模型路由网络配置列表</p>
	ServiceEndPoints []*ServiceEndPoints `json:"ServiceEndPoints,omitnil,omitempty" name:"ServiceEndPoints"`

	// <p>模型路由实例状态</p><p>枚举值：</p><ul><li>Active： 运行中</li><li>Provisioning： 创建中</li><li>Configuring： 变配中</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>模型路由实例所属子网的ID</p>
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>标签</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>模型路由实例的计费状态</p><p>枚举值：</p><ul><li>Normal： 正常</li><li>Isolated： 已隔离</li></ul>
	TradeStatus *string `json:"TradeStatus,omitnil,omitempty" name:"TradeStatus"`

	// <p>模型路由实例VIP</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>模型路由实例所属VPC的ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type ModelRouterLog struct {
	// <p>API Key的ID</p>
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// <p>模型</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>所属厂商</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>请求状态</p><p>枚举值：</p><ul><li>failure： 失败</li><li>success： 成功</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>最大重试次数</p>
	MaxRetries *uint64 `json:"MaxRetries,omitnil,omitempty" name:"MaxRetries"`

	// <p>单次请求消耗的总Token数量</p>
	TotalTokens *uint64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`

	// <p>单次请求输入消耗的Token数量</p>
	InputTokens *uint64 `json:"InputTokens,omitnil,omitempty" name:"InputTokens"`

	// <p>单次请求输出消耗的Token数量</p>
	OutputTokens *uint64 `json:"OutputTokens,omitnil,omitempty" name:"OutputTokens"`

	// <p>请求耗时</p><p>单位：ms</p>
	RequestDuration *uint64 `json:"RequestDuration,omitnil,omitempty" name:"RequestDuration"`

	// <p>请求IP</p>
	RequesterIp *string `json:"RequesterIp,omitnil,omitempty" name:"RequesterIp"`

	// <p>日志查询起始时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>日志查询结束时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ModelRouterModel struct {
	// <p>模型名称</p>
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// <p>所属厂商</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>模型类型。</p><p>枚举值：</p><ul><li>BYOK： BYOK类型</li><li>Platform： 平台类型</li></ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>服务商/模型 ID（byok_model.model_id，形如 model-xxxxxxxx；Platform 类型不传）</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`
}

type ModelRouterPackage struct {
	// <p>模型路由资源包总容量</p>
	CapacitySize *string `json:"CapacitySize,omitnil,omitempty" name:"CapacitySize"`

	// <p>模型路由资源包总余量</p>
	CapacityRemain *string `json:"CapacityRemain,omitnil,omitempty" name:"CapacityRemain"`

	// <p>模型路由资源包容量精确值</p>
	CapacitySizePrecise *string `json:"CapacitySizePrecise,omitnil,omitempty" name:"CapacitySizePrecise"`

	// <p>模型路由资源包总余量精确值</p>
	CapacityRemainPrecise *string `json:"CapacityRemainPrecise,omitnil,omitempty" name:"CapacityRemainPrecise"`

	// <p>模型路由资源包设置用尽续购标志位 0:未设置 1:用尽到期新购</p><p>取值范围：[0, 1]</p>
	AutoPurchaseFlag *uint64 `json:"AutoPurchaseFlag,omitnil,omitempty" name:"AutoPurchaseFlag"`

	// <p>模型路由资源包Id</p>
	ModelRouterResourcePackageId *string `json:"ModelRouterResourcePackageId,omitnil,omitempty" name:"ModelRouterResourcePackageId"`

	// <p>模型路由资源包创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>模型路由资源包抵扣开始时间</p>
	DeductionStartTime *string `json:"DeductionStartTime,omitnil,omitempty" name:"DeductionStartTime"`

	// <p>模型路由资源包抵扣截止时间</p>
	DeductionEndTime *string `json:"DeductionEndTime,omitnil,omitempty" name:"DeductionEndTime"`

	// <p>模型路由资源包失效时间</p>
	ExpiredTime *string `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// <p>模型路由资源包状态</p><p>枚举值：</p><ul><li>0： 有效</li><li>1： 已退款</li><li>2： 已过期</li><li>3： 已用完</li></ul>
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModelRouterQuota struct {
	// <p>配额名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuotaType *string `json:"QuotaType,omitnil,omitempty" name:"QuotaType"`

	// <p>资源ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// <p>配额上限</p><p>单位：个</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>已使用配额数量</p><p>单位：个</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Used *uint64 `json:"Used,omitnil,omitempty" name:"Used"`

	// <p>剩余配额数量</p><p>单位：个</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Available *uint64 `json:"Available,omitnil,omitempty" name:"Available"`
}

type ModelRouterResourcePackageDeduction struct {
	// <p>实际抵扣量</p>
	ActualDosage *string `json:"ActualDosage,omitnil,omitempty" name:"ActualDosage"`

	// <p>抵扣后包剩余量</p>
	AfterDeductionRemain *string `json:"AfterDeductionRemain,omitnil,omitempty" name:"AfterDeductionRemain"`

	// <p>抵扣前包剩余量</p>
	BeforeDeductionRemain *string `json:"BeforeDeductionRemain,omitnil,omitempty" name:"BeforeDeductionRemain"`

	// <p>抵扣时间</p>
	DeductionTime *string `json:"DeductionTime,omitnil,omitempty" name:"DeductionTime"`

	// <p>原始用量</p>
	Dosage *string `json:"Dosage,omitnil,omitempty" name:"Dosage"`

	// <p>用量结束时间</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>产生用量的模型路由实例Id</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>模型路由资源包Id</p>
	ModelRouterResourcePackageId *string `json:"ModelRouterResourcePackageId,omitnil,omitempty" name:"ModelRouterResourcePackageId"`

	// <p>用量开始时间</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type ModelRouterResourcePackageRefundPrice struct {
	// <p>模型路由资源包Id</p>
	ModelRouterPackageId *string `json:"ModelRouterPackageId,omitnil,omitempty" name:"ModelRouterPackageId"`

	// <p>可退还金额</p>
	Price *float64 `json:"Price,omitnil,omitempty" name:"Price"`
}

type ModelRouterSet struct {
	// <p>模型路由实例关联的Budget ID。</p><p>未关联Budget时返回空字符串。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>模型路由实例关联的Budget名称。</p><p>未关联Budget时返回空字符串。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BudgetName *string `json:"BudgetName,omitnil,omitempty" name:"BudgetName"`

	// <p>集群信息</p>
	ClusterInfo *ClusterInfo `json:"ClusterInfo,omitnil,omitempty" name:"ClusterInfo"`

	// <p>创建时间</p>
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// <p>模型路由实例按Budget刷新周期划分的Credit使用情况。</p><p>当关联Budget配置多个刷新周期时，按1d、7d、30d顺序返回各周期用量；未关联Budget时返回空数组。</p>
	CreditUsageSet []*CreditUsage `json:"CreditUsageSet,omitnil,omitempty" name:"CreditUsageSet"`

	// <p>模型路由实例域名</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>模型路由ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>模型路由名称</p><p>默认值：-</p>
	ModelRouterName *string `json:"ModelRouterName,omitnil,omitempty" name:"ModelRouterName"`

	// <p>模型路由类型</p><p>枚举值：</p><ul><li>Shared： 共享型</li><li>Enterprise： 企业级</li></ul>
	ModelRouterType *string `json:"ModelRouterType,omitnil,omitempty" name:"ModelRouterType"`

	// <p>修改时间</p>
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// <p>模型路由实例网络类型</p><p>枚举值：</p><ul><li>Internet： 公网</li><li>Intranet： 内网</li></ul>
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// <p>模型路由实例的安全状态</p><p>枚举值：</p><ul><li>Normal： 正常</li><li>Banned： 已封禁</li><li>Frozen： 已冻结</li></ul>
	SecurityStatus *string `json:"SecurityStatus,omitnil,omitempty" name:"SecurityStatus"`

	// <p>模型路由实例状态</p><p>枚举值：</p><ul><li>Active： 运行中</li><li>Provisioning： 创建中</li><li>Configuring： 变配中</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>标签</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>模型路由实例的计费状态</p><p>枚举值：</p><ul><li>Normal： 正常</li><li>Isolated： 已隔离</li></ul>
	TradeStatus *string `json:"TradeStatus,omitnil,omitempty" name:"TradeStatus"`

	// <p>模型路由实例VIP</p>
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// <p>模型路由实例所属VPC的ID</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type ModelTestResult struct {
	// <p>模型</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>健康状况</p><p>枚举值：</p><ul><li>Success： 健康</li><li>Error： 不健康</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>错误信息</p>
	ErrorInfo *ProviderTestConnectionErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>探测请求</p>
	TestConnectionRequest *TestConnectionRequestInfo `json:"TestConnectionRequest,omitnil,omitempty" name:"TestConnectionRequest"`
}

// Predefined struct for user
type ModifyBlockIPListRequestParams struct {
	// 负载均衡实例ID
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 操作类型，可取：
	// <li> add_customized_field（首次设置header，开启黑名单功能）</li>
	// <li> set_customized_field（修改header）</li>
	// <li> del_customized_field（删除header）</li>
	// <li> add_blocked（添加黑名单）</li>
	// <li> del_blocked（删除黑名单）</li>
	// <li> flush_blocked（清空黑名单）</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 客户端真实IP存放的header字段名
	ClientIPField *string `json:"ClientIPField,omitnil,omitempty" name:"ClientIPField"`

	// 封禁IP列表，单次操作数组最大长度支持200000
	BlockIPList []*string `json:"BlockIPList,omitnil,omitempty" name:"BlockIPList"`

	// 过期时间，单位秒，默认值3600
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 添加IP的策略，可取：fifo（如果黑名单容量已满，新加入黑名单的IP采用先进先出策略）
	AddStrategy *string `json:"AddStrategy,omitnil,omitempty" name:"AddStrategy"`
}

type ModifyBlockIPListRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 操作类型，可取：
	// <li> add_customized_field（首次设置header，开启黑名单功能）</li>
	// <li> set_customized_field（修改header）</li>
	// <li> del_customized_field（删除header）</li>
	// <li> add_blocked（添加黑名单）</li>
	// <li> del_blocked（删除黑名单）</li>
	// <li> flush_blocked（清空黑名单）</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 客户端真实IP存放的header字段名
	ClientIPField *string `json:"ClientIPField,omitnil,omitempty" name:"ClientIPField"`

	// 封禁IP列表，单次操作数组最大长度支持200000
	BlockIPList []*string `json:"BlockIPList,omitnil,omitempty" name:"BlockIPList"`

	// 过期时间，单位秒，默认值3600
	ExpireTime *uint64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 添加IP的策略，可取：fifo（如果黑名单容量已满，新加入黑名单的IP采用先进先出策略）
	AddStrategy *string `json:"AddStrategy,omitnil,omitempty" name:"AddStrategy"`
}

func (r *ModifyBlockIPListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockIPListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	delete(f, "Type")
	delete(f, "ClientIPField")
	delete(f, "BlockIPList")
	delete(f, "ExpireTime")
	delete(f, "AddStrategy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBlockIPListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBlockIPListResponseParams struct {
	// 异步任务的ID
	JodId *string `json:"JodId,omitnil,omitempty" name:"JodId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBlockIPListResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBlockIPListResponseParams `json:"Response"`
}

func (r *ModifyBlockIPListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBlockIPListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBudgetAttributesRequestParams struct {
	// <p>Budget ID。</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>预算配置数组。</p><p>数组长度最大为3，最多可同时配置1d、7d、30d三个刷新周期，且每种刷新周期只能出现一次。BudgetResetAt不支持作为入参设置，系统会按配置的刷新周期自动维护刷新时间。</p>
	BudgetConfigs []*BudgetConfigInput `json:"BudgetConfigs,omitnil,omitempty" name:"BudgetConfigs"`

	// <p>Budget名称。</p>
	BudgetName *string `json:"BudgetName,omitnil,omitempty" name:"BudgetName"`

	// <p>Budget限速配置。</p>
	RateLimitConfig *RateLimitConfigForBudget `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`
}

type ModifyBudgetAttributesRequest struct {
	*tchttp.BaseRequest
	
	// <p>Budget ID。</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>预算配置数组。</p><p>数组长度最大为3，最多可同时配置1d、7d、30d三个刷新周期，且每种刷新周期只能出现一次。BudgetResetAt不支持作为入参设置，系统会按配置的刷新周期自动维护刷新时间。</p>
	BudgetConfigs []*BudgetConfigInput `json:"BudgetConfigs,omitnil,omitempty" name:"BudgetConfigs"`

	// <p>Budget名称。</p>
	BudgetName *string `json:"BudgetName,omitnil,omitempty" name:"BudgetName"`

	// <p>Budget限速配置。</p>
	RateLimitConfig *RateLimitConfigForBudget `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`
}

func (r *ModifyBudgetAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBudgetAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BudgetId")
	delete(f, "BudgetConfigs")
	delete(f, "BudgetName")
	delete(f, "RateLimitConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyBudgetAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyBudgetAttributesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyBudgetAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyBudgetAttributesResponseParams `json:"Response"`
}

func (r *ModifyBudgetAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyBudgetAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainAttributesRequestParams struct {
	// 负载均衡实例ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID，可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口查询。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 域名（必须是已经创建的转发规则下的域名），如果是多域名，可以指定多域名列表中的任意一个，可以通过[DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口查询。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 要修改的新域名。NewDomain和NewDomains只能传一个。
	NewDomain *string `json:"NewDomain,omitnil,omitempty" name:"NewDomain"`

	// 域名相关的证书信息，注意，仅对启用SNI的监听器适用，不可和MultiCertInfo 同时传入。
	Certificate *CertificateInput `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// 是否开启HTTP2，注意，只有HTTPS域名才能开启HTTP2。
	// True: 开启HTTP2，Fasle: 不开启HTTP2。
	Http2 *bool `json:"Http2,omitnil,omitempty" name:"Http2"`

	// 是否设为默认域名，注意，一个监听器下只能设置一个默认域名。
	// True: 设为默认域名，Fasle: 不设置为默认域名。
	DefaultServer *bool `json:"DefaultServer,omitnil,omitempty" name:"DefaultServer"`

	// 是否开启 QUIC，注意，只有 HTTPS 域名才能开启 QUIC。
	// True: 开启 QUIC，False: 不开启QUIC。
	Quic *bool `json:"Quic,omitnil,omitempty" name:"Quic"`

	// 监听器下必须配置一个默认域名，若要关闭原默认域名，必须同时指定另一个域名作为新的默认域名，如果新的默认域名是多域名，可以指定多域名列表中的任意一个。
	NewDefaultServerDomain *string `json:"NewDefaultServerDomain,omitnil,omitempty" name:"NewDefaultServerDomain"`

	// 要修改的新域名列表。NewDomain和NewDomains只能传一个。
	NewDomains []*string `json:"NewDomains,omitnil,omitempty" name:"NewDomains"`

	// 域名相关的证书信息，注意，仅对启用SNI的监听器适用；支持同时传入多本算法类型不同的服务器证书，不可和Certificate 同时传入。
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitnil,omitempty" name:"MultiCertInfo"`
}

type ModifyDomainAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID，可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口查询。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 域名（必须是已经创建的转发规则下的域名），如果是多域名，可以指定多域名列表中的任意一个，可以通过[DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口查询。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 要修改的新域名。NewDomain和NewDomains只能传一个。
	NewDomain *string `json:"NewDomain,omitnil,omitempty" name:"NewDomain"`

	// 域名相关的证书信息，注意，仅对启用SNI的监听器适用，不可和MultiCertInfo 同时传入。
	Certificate *CertificateInput `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// 是否开启HTTP2，注意，只有HTTPS域名才能开启HTTP2。
	// True: 开启HTTP2，Fasle: 不开启HTTP2。
	Http2 *bool `json:"Http2,omitnil,omitempty" name:"Http2"`

	// 是否设为默认域名，注意，一个监听器下只能设置一个默认域名。
	// True: 设为默认域名，Fasle: 不设置为默认域名。
	DefaultServer *bool `json:"DefaultServer,omitnil,omitempty" name:"DefaultServer"`

	// 是否开启 QUIC，注意，只有 HTTPS 域名才能开启 QUIC。
	// True: 开启 QUIC，False: 不开启QUIC。
	Quic *bool `json:"Quic,omitnil,omitempty" name:"Quic"`

	// 监听器下必须配置一个默认域名，若要关闭原默认域名，必须同时指定另一个域名作为新的默认域名，如果新的默认域名是多域名，可以指定多域名列表中的任意一个。
	NewDefaultServerDomain *string `json:"NewDefaultServerDomain,omitnil,omitempty" name:"NewDefaultServerDomain"`

	// 要修改的新域名列表。NewDomain和NewDomains只能传一个。
	NewDomains []*string `json:"NewDomains,omitnil,omitempty" name:"NewDomains"`

	// 域名相关的证书信息，注意，仅对启用SNI的监听器适用；支持同时传入多本算法类型不同的服务器证书，不可和Certificate 同时传入。
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitnil,omitempty" name:"MultiCertInfo"`
}

func (r *ModifyDomainAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "Domain")
	delete(f, "NewDomain")
	delete(f, "Certificate")
	delete(f, "Http2")
	delete(f, "DefaultServer")
	delete(f, "Quic")
	delete(f, "NewDefaultServerDomain")
	delete(f, "NewDomains")
	delete(f, "MultiCertInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainAttributesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDomainAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainAttributesResponseParams `json:"Response"`
}

func (r *ModifyDomainAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainRequestParams struct {
	// 负载均衡实例 ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID， 可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口查询。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 监听器下的某个旧域名， 可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 查询。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 新域名，	长度限制为：1-120。有三种使用格式：非正则表达式格式，通配符格式，正则表达式格式。非正则表达式格式只能使用字母、数字、‘-’、‘.’。通配符格式的使用 ‘*’ 只能在开头或者结尾。正则表达式以'~'开头。
	NewDomain *string `json:"NewDomain,omitnil,omitempty" name:"NewDomain"`
}

type ModifyDomainRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID， 可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口查询。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 监听器下的某个旧域名， 可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 查询。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 新域名，	长度限制为：1-120。有三种使用格式：非正则表达式格式，通配符格式，正则表达式格式。非正则表达式格式只能使用字母、数字、‘-’、‘.’。通配符格式的使用 ‘*’ 只能在开头或者结尾。正则表达式以'~'开头。
	NewDomain *string `json:"NewDomain,omitnil,omitempty" name:"NewDomain"`
}

func (r *ModifyDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "Domain")
	delete(f, "NewDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDomainResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDomainResponseParams `json:"Response"`
}

func (r *ModifyDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFunctionTargetsRequestParams struct {
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 要修改的后端云函数服务列表，仅支持 Event 函数类型。
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitnil,omitempty" name:"FunctionTargets"`

	// 转发规则的ID，当绑定机器到七层转发规则时，必须提供此参数或Domain+Url两者之一。
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 目标规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 目标规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type ModifyFunctionTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 要修改的后端云函数服务列表，仅支持 Event 函数类型。
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitnil,omitempty" name:"FunctionTargets"`

	// 转发规则的ID，当绑定机器到七层转发规则时，必须提供此参数或Domain+Url两者之一。
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 目标规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 目标规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

func (r *ModifyFunctionTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFunctionTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "FunctionTargets")
	delete(f, "LocationId")
	delete(f, "Domain")
	delete(f, "Url")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFunctionTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFunctionTargetsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyFunctionTargetsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFunctionTargetsResponseParams `json:"Response"`
}

func (r *ModifyFunctionTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFunctionTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIntentRouterAttributeRequestParams struct {
	// <p>意图路由ID（ir-xxx格式）。</p>
	IntentRouterId *string `json:"IntentRouterId,omitnil,omitempty" name:"IntentRouterId"`

	// <p>模型路由实例ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>新的路由名称。</p><p>选填；必须以"IntentRouter/"为前缀，后缀仅支持字母、数字、连字符和下划线，后缀长度1-128个字符。不传则不修改。</p>
	RouteName *string `json:"RouteName,omitnil,omitempty" name:"RouteName"`

	// <p>意图路由描述。</p>
	RouterDescribe *string `json:"RouterDescribe,omitnil,omitempty" name:"RouterDescribe"`

	// <p>新的分层配置列表（全量替换）。</p><p>选填；不传则不修改。传入时必须为完整分层集合：复杂度分层须包含全部 4 个分层 SIMPLE/MEDIUM/COMPLEX/REASONING；语义分层须包含 default 及各语义 Tier（取决于实例所用协议，且不可跨协议变更）。每个分层至少包含一个模型，模型名称必须是已关联到该实例的模型。</p>
	Tiers []*TierItem `json:"Tiers,omitnil,omitempty" name:"Tiers"`
}

type ModifyIntentRouterAttributeRequest struct {
	*tchttp.BaseRequest
	
	// <p>意图路由ID（ir-xxx格式）。</p>
	IntentRouterId *string `json:"IntentRouterId,omitnil,omitempty" name:"IntentRouterId"`

	// <p>模型路由实例ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>新的路由名称。</p><p>选填；必须以"IntentRouter/"为前缀，后缀仅支持字母、数字、连字符和下划线，后缀长度1-128个字符。不传则不修改。</p>
	RouteName *string `json:"RouteName,omitnil,omitempty" name:"RouteName"`

	// <p>意图路由描述。</p>
	RouterDescribe *string `json:"RouterDescribe,omitnil,omitempty" name:"RouterDescribe"`

	// <p>新的分层配置列表（全量替换）。</p><p>选填；不传则不修改。传入时必须为完整分层集合：复杂度分层须包含全部 4 个分层 SIMPLE/MEDIUM/COMPLEX/REASONING；语义分层须包含 default 及各语义 Tier（取决于实例所用协议，且不可跨协议变更）。每个分层至少包含一个模型，模型名称必须是已关联到该实例的模型。</p>
	Tiers []*TierItem `json:"Tiers,omitnil,omitempty" name:"Tiers"`
}

func (r *ModifyIntentRouterAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIntentRouterAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IntentRouterId")
	delete(f, "ModelRouterId")
	delete(f, "RouteName")
	delete(f, "RouterDescribe")
	delete(f, "Tiers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIntentRouterAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIntentRouterAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyIntentRouterAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIntentRouterAttributeResponseParams `json:"Response"`
}

func (r *ModifyIntentRouterAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIntentRouterAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKeyAttributesRequestParams struct {
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>API Key的ID</p>
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// <p>Key的名称</p>
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// <p>限速配置</p>
	RateLimitConfig *RateLimitConfigForKey `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`
}

type ModifyKeyAttributesRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>API Key的ID</p>
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// <p>Key的名称</p>
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// <p>限速配置</p>
	RateLimitConfig *RateLimitConfigForKey `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`
}

func (r *ModifyKeyAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKeyAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "KeyId")
	delete(f, "KeyName")
	delete(f, "RateLimitConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyKeyAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKeyAttributesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyKeyAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyKeyAttributesResponseParams `json:"Response"`
}

func (r *ModifyKeyAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKeyAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKeysBlockStatusRequestParams struct {
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>是否停止使用</p>
	Blocked *bool `json:"Blocked,omitnil,omitempty" name:"Blocked"`

	// <p>需要修改的Key的ID列表</p>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

type ModifyKeysBlockStatusRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>是否停止使用</p>
	Blocked *bool `json:"Blocked,omitnil,omitempty" name:"Blocked"`

	// <p>需要修改的Key的ID列表</p>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

func (r *ModifyKeysBlockStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKeysBlockStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "Blocked")
	delete(f, "KeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyKeysBlockStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKeysBlockStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyKeysBlockStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyKeysBlockStatusResponseParams `json:"Response"`
}

func (r *ModifyKeysBlockStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKeysBlockStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKeysUserGroupRequestParams struct {
	// <p>模型路由实例ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>目标归属用户组ID。传真实用户组ID表示批量入组或跨组移动（Key 已属其它组则改为目标组）；传 ugrp-ungrouped 表示批量移出到未分组。</p>
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// <p>待变更归属的 Key ID 列表，单次1-100个。</p>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

type ModifyKeysUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>目标归属用户组ID。传真实用户组ID表示批量入组或跨组移动（Key 已属其它组则改为目标组）；传 ugrp-ungrouped 表示批量移出到未分组。</p>
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// <p>待变更归属的 Key ID 列表，单次1-100个。</p>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

func (r *ModifyKeysUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKeysUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "UserGroupId")
	delete(f, "KeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyKeysUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyKeysUserGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyKeysUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyKeysUserGroupResponseParams `json:"Response"`
}

func (r *ModifyKeysUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyKeysUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyListenerRequestParams struct {
	// <p>负载均衡实例ID，可以通过 <a href="https://cloud.tencent.com/document/product/214/30685">DescribeLoadBalancers</a> 接口查询。</p>
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// <p>负载均衡监听器ID，可以通过 <a href="https://cloud.tencent.com/document/product/214/30686">DescribeListeners</a> 接口查询。</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>新的监听器名称。命名规则：1-80 个英文字母、汉字等国际通用语言字符，数字，连接线“-”、下划线“_”等常见字符（禁止Unicode补充字符，如emoji表情、生僻汉字等）。</p>
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// <p>会话保持时间，单位：秒。可选值：30~3600，默认 0，表示不开启。此参数仅适用于TCP/UDP监听器。</p>
	SessionExpireTime *int64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// <p>健康检查相关参数，此参数仅适用于TCP/UDP/TCP_SSL/QUIC监听器。</p>
	HealthCheck *HealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// <p>证书相关信息，此参数仅适用于HTTPS/TCP_SSL/QUIC监听器；此参数和MultiCertInfo不能同时传入。</p>
	Certificate *CertificateInput `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// <p>监听器转发的方式。可选值：WRR（按权重轮询）、LEAST_CONN（按最小连接数）、IP_HASH（按 IP 地址哈希）<br>分别表示按权重轮询、最小连接数， 默认为 WRR。<br>使用场景：适用于TCP/UDP/TCP_SSL/QUIC监听器。七层监听器的均衡方式应在转发规则中修改。</p>
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// <p>是否开启SNI特性，此参数仅适用于HTTPS监听器。默认0，表示不开启，1表示开启。注意：未开启SNI的监听器可以开启SNI；已开启SNI的监听器不能关闭SNI。</p>
	SniSwitch *int64 `json:"SniSwitch,omitnil,omitempty" name:"SniSwitch"`

	// <p>后端目标类型，NODE表示绑定普通节点，TARGETGROUP表示绑定目标组。</p>
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// <p>是否开启长连接，此参数仅适用于HTTP/HTTPS监听器。<br>默认值0表示不开启，1表示开启。<br>若后端服务对连接数上限有限制，则建议谨慎开启。此功能目前处于内测中，如需使用，请提交 <a href="https://cloud.tencent.com/apply/p/tsodp6qm21">内测申请</a>。</p>
	KeepaliveEnable *int64 `json:"KeepaliveEnable,omitnil,omitempty" name:"KeepaliveEnable"`

	// <p>重新调度功能，解绑后端服务开关，打开此开关，当解绑后端服务时触发重新调度。仅TCP/UDP监听器支持。</p>
	DeregisterTargetRst *bool `json:"DeregisterTargetRst,omitnil,omitempty" name:"DeregisterTargetRst"`

	// <p>会话保持类型。NORMAL表示默认会话保持类型。QUIC_CID表示根据Quic Connection ID做会话保持。QUIC_CID只支持UDP协议。<br>使用场景：适用于TCP/UDP/TCP_SSL/QUIC监听器。<br>默认为 NORMAL。</p>
	SessionType *string `json:"SessionType,omitnil,omitempty" name:"SessionType"`

	// <p>证书信息，支持同时传入不同算法类型的多本服务端证书；此参数仅适用于未开启SNI特性的HTTPS监听器。此参数和Certificate不能同时传入。</p>
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitnil,omitempty" name:"MultiCertInfo"`

	// <p>监听器粒度并发连接数上限，当前仅性能容量型实例且仅TCP/UDP/TCP_SSL/QUIC监听器支持。取值范围：1-实例规格并发连接上限，其中-1表示关闭监听器粒度并发连接数限速。基础网络实例不支持该参数。<br>默认为 -1，表示不限速。</p>
	MaxConn *int64 `json:"MaxConn,omitnil,omitempty" name:"MaxConn"`

	// <p>监听器粒度新建连接数上限，当前仅性能容量型实例且仅TCP/UDP/TCP_SSL/QUIC监听器支持。取值范围：1-实例规格新建连接上限，其中-1表示关闭监听器粒度新建连接数限速。基础网络实例不支持该参数。<br>默认为 -1 表示不限速。</p>
	MaxCps *int64 `json:"MaxCps,omitnil,omitempty" name:"MaxCps"`

	// <p>空闲连接超时时间，此参数仅适用于TCP/UDP监听器。如需设置超过1980s，请通过 <a href="https://console.cloud.tencent.com/workorder/category">工单申请</a>,最大可设置到3600s。</p><p>取值范围：[10, 1980]</p><p>单位：秒</p><p>默认值：900</p><p>TCP监听器默认值：900，UDP监听器默认值：300s。取值范围：共享型实例和独占型实例支持：10～900，性能容量型实例支持：10~1980。</p>
	IdleConnectTimeout *int64 `json:"IdleConnectTimeout,omitnil,omitempty" name:"IdleConnectTimeout"`

	// <p>TCP_SSL和QUIC是否支持PP</p>
	ProxyProtocol *bool `json:"ProxyProtocol,omitnil,omitempty" name:"ProxyProtocol"`

	// <p>是否开启SNAT（源IP替换），True（开启）、False（关闭）。默认为关闭。注意：SnatEnable开启时会替换客户端源IP，此时<code>透传客户端源IP</code>选项关闭，反之亦然。</p>
	SnatEnable *bool `json:"SnatEnable,omitnil,omitempty" name:"SnatEnable"`

	// <p>数据压缩模式</p><p>枚举值：</p><ul><li>transparent： 透明模式（默认值）</li><li>compatibility： 兼容模式（开启 gzip 兼容压缩配置）</li></ul>
	DataCompressMode *string `json:"DataCompressMode,omitnil,omitempty" name:"DataCompressMode"`

	// <p>重新调度功能，权重调为0开关，打开此开关，后端服务器权重调为0时触发重新调度。仅TCP/UDP监听器支持。</p>
	RescheduleTargetZeroWeight *bool `json:"RescheduleTargetZeroWeight,omitnil,omitempty" name:"RescheduleTargetZeroWeight"`

	// <p>重新调度功能，健康检查异常开关，打开此开关，后端服务器健康检查异常时触发重新调度。仅TCP/UDP监听器支持。</p>
	RescheduleUnhealthy *bool `json:"RescheduleUnhealthy,omitnil,omitempty" name:"RescheduleUnhealthy"`

	// <p>重新调度功能，扩容后端服务开关，打开此开关，后端服务器增加或者减少时触发重新调度。仅TCP/UDP监听器支持。</p>
	RescheduleExpandTarget *bool `json:"RescheduleExpandTarget,omitnil,omitempty" name:"RescheduleExpandTarget"`

	// <p>重新调度触发开始时间，取值0~3600s。仅TCP/UDP监听器支持。</p>
	RescheduleStartTime *int64 `json:"RescheduleStartTime,omitnil,omitempty" name:"RescheduleStartTime"`

	// <p>重新调度触发持续时间，取值0~3600s。仅TCP/UDP监听器支持。</p>
	RescheduleInterval *int64 `json:"RescheduleInterval,omitnil,omitempty" name:"RescheduleInterval"`
}

type ModifyListenerRequest struct {
	*tchttp.BaseRequest
	
	// <p>负载均衡实例ID，可以通过 <a href="https://cloud.tencent.com/document/product/214/30685">DescribeLoadBalancers</a> 接口查询。</p>
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// <p>负载均衡监听器ID，可以通过 <a href="https://cloud.tencent.com/document/product/214/30686">DescribeListeners</a> 接口查询。</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>新的监听器名称。命名规则：1-80 个英文字母、汉字等国际通用语言字符，数字，连接线“-”、下划线“_”等常见字符（禁止Unicode补充字符，如emoji表情、生僻汉字等）。</p>
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// <p>会话保持时间，单位：秒。可选值：30~3600，默认 0，表示不开启。此参数仅适用于TCP/UDP监听器。</p>
	SessionExpireTime *int64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// <p>健康检查相关参数，此参数仅适用于TCP/UDP/TCP_SSL/QUIC监听器。</p>
	HealthCheck *HealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// <p>证书相关信息，此参数仅适用于HTTPS/TCP_SSL/QUIC监听器；此参数和MultiCertInfo不能同时传入。</p>
	Certificate *CertificateInput `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// <p>监听器转发的方式。可选值：WRR（按权重轮询）、LEAST_CONN（按最小连接数）、IP_HASH（按 IP 地址哈希）<br>分别表示按权重轮询、最小连接数， 默认为 WRR。<br>使用场景：适用于TCP/UDP/TCP_SSL/QUIC监听器。七层监听器的均衡方式应在转发规则中修改。</p>
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// <p>是否开启SNI特性，此参数仅适用于HTTPS监听器。默认0，表示不开启，1表示开启。注意：未开启SNI的监听器可以开启SNI；已开启SNI的监听器不能关闭SNI。</p>
	SniSwitch *int64 `json:"SniSwitch,omitnil,omitempty" name:"SniSwitch"`

	// <p>后端目标类型，NODE表示绑定普通节点，TARGETGROUP表示绑定目标组。</p>
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// <p>是否开启长连接，此参数仅适用于HTTP/HTTPS监听器。<br>默认值0表示不开启，1表示开启。<br>若后端服务对连接数上限有限制，则建议谨慎开启。此功能目前处于内测中，如需使用，请提交 <a href="https://cloud.tencent.com/apply/p/tsodp6qm21">内测申请</a>。</p>
	KeepaliveEnable *int64 `json:"KeepaliveEnable,omitnil,omitempty" name:"KeepaliveEnable"`

	// <p>重新调度功能，解绑后端服务开关，打开此开关，当解绑后端服务时触发重新调度。仅TCP/UDP监听器支持。</p>
	DeregisterTargetRst *bool `json:"DeregisterTargetRst,omitnil,omitempty" name:"DeregisterTargetRst"`

	// <p>会话保持类型。NORMAL表示默认会话保持类型。QUIC_CID表示根据Quic Connection ID做会话保持。QUIC_CID只支持UDP协议。<br>使用场景：适用于TCP/UDP/TCP_SSL/QUIC监听器。<br>默认为 NORMAL。</p>
	SessionType *string `json:"SessionType,omitnil,omitempty" name:"SessionType"`

	// <p>证书信息，支持同时传入不同算法类型的多本服务端证书；此参数仅适用于未开启SNI特性的HTTPS监听器。此参数和Certificate不能同时传入。</p>
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitnil,omitempty" name:"MultiCertInfo"`

	// <p>监听器粒度并发连接数上限，当前仅性能容量型实例且仅TCP/UDP/TCP_SSL/QUIC监听器支持。取值范围：1-实例规格并发连接上限，其中-1表示关闭监听器粒度并发连接数限速。基础网络实例不支持该参数。<br>默认为 -1，表示不限速。</p>
	MaxConn *int64 `json:"MaxConn,omitnil,omitempty" name:"MaxConn"`

	// <p>监听器粒度新建连接数上限，当前仅性能容量型实例且仅TCP/UDP/TCP_SSL/QUIC监听器支持。取值范围：1-实例规格新建连接上限，其中-1表示关闭监听器粒度新建连接数限速。基础网络实例不支持该参数。<br>默认为 -1 表示不限速。</p>
	MaxCps *int64 `json:"MaxCps,omitnil,omitempty" name:"MaxCps"`

	// <p>空闲连接超时时间，此参数仅适用于TCP/UDP监听器。如需设置超过1980s，请通过 <a href="https://console.cloud.tencent.com/workorder/category">工单申请</a>,最大可设置到3600s。</p><p>取值范围：[10, 1980]</p><p>单位：秒</p><p>默认值：900</p><p>TCP监听器默认值：900，UDP监听器默认值：300s。取值范围：共享型实例和独占型实例支持：10～900，性能容量型实例支持：10~1980。</p>
	IdleConnectTimeout *int64 `json:"IdleConnectTimeout,omitnil,omitempty" name:"IdleConnectTimeout"`

	// <p>TCP_SSL和QUIC是否支持PP</p>
	ProxyProtocol *bool `json:"ProxyProtocol,omitnil,omitempty" name:"ProxyProtocol"`

	// <p>是否开启SNAT（源IP替换），True（开启）、False（关闭）。默认为关闭。注意：SnatEnable开启时会替换客户端源IP，此时<code>透传客户端源IP</code>选项关闭，反之亦然。</p>
	SnatEnable *bool `json:"SnatEnable,omitnil,omitempty" name:"SnatEnable"`

	// <p>数据压缩模式</p><p>枚举值：</p><ul><li>transparent： 透明模式（默认值）</li><li>compatibility： 兼容模式（开启 gzip 兼容压缩配置）</li></ul>
	DataCompressMode *string `json:"DataCompressMode,omitnil,omitempty" name:"DataCompressMode"`

	// <p>重新调度功能，权重调为0开关，打开此开关，后端服务器权重调为0时触发重新调度。仅TCP/UDP监听器支持。</p>
	RescheduleTargetZeroWeight *bool `json:"RescheduleTargetZeroWeight,omitnil,omitempty" name:"RescheduleTargetZeroWeight"`

	// <p>重新调度功能，健康检查异常开关，打开此开关，后端服务器健康检查异常时触发重新调度。仅TCP/UDP监听器支持。</p>
	RescheduleUnhealthy *bool `json:"RescheduleUnhealthy,omitnil,omitempty" name:"RescheduleUnhealthy"`

	// <p>重新调度功能，扩容后端服务开关，打开此开关，后端服务器增加或者减少时触发重新调度。仅TCP/UDP监听器支持。</p>
	RescheduleExpandTarget *bool `json:"RescheduleExpandTarget,omitnil,omitempty" name:"RescheduleExpandTarget"`

	// <p>重新调度触发开始时间，取值0~3600s。仅TCP/UDP监听器支持。</p>
	RescheduleStartTime *int64 `json:"RescheduleStartTime,omitnil,omitempty" name:"RescheduleStartTime"`

	// <p>重新调度触发持续时间，取值0~3600s。仅TCP/UDP监听器支持。</p>
	RescheduleInterval *int64 `json:"RescheduleInterval,omitnil,omitempty" name:"RescheduleInterval"`
}

func (r *ModifyListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "ListenerName")
	delete(f, "SessionExpireTime")
	delete(f, "HealthCheck")
	delete(f, "Certificate")
	delete(f, "Scheduler")
	delete(f, "SniSwitch")
	delete(f, "TargetType")
	delete(f, "KeepaliveEnable")
	delete(f, "DeregisterTargetRst")
	delete(f, "SessionType")
	delete(f, "MultiCertInfo")
	delete(f, "MaxConn")
	delete(f, "MaxCps")
	delete(f, "IdleConnectTimeout")
	delete(f, "ProxyProtocol")
	delete(f, "SnatEnable")
	delete(f, "DataCompressMode")
	delete(f, "RescheduleTargetZeroWeight")
	delete(f, "RescheduleUnhealthy")
	delete(f, "RescheduleExpandTarget")
	delete(f, "RescheduleStartTime")
	delete(f, "RescheduleInterval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyListenerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyListenerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyListenerResponseParams `json:"Response"`
}

func (r *ModifyListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerAttributesRequestParams struct {
	// <p>负载均衡的唯一ID，可以通过 <a href="https://cloud.tencent.com/document/product/214/30685">DescribeLoadBalancers</a> 接口获取。</p>
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// <p>负载均衡实例名称，规则：1-80 个英文字母、汉字等国际通用语言字符，数字，连接线“-”、下划线“_”等常见字符（禁止Unicode补充字符，如emoji表情、生僻汉字等）。</p>
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// <p>设置负载均衡跨地域绑定1.0的后端服务信息</p>
	TargetRegionInfo *TargetRegionInfo `json:"TargetRegionInfo,omitnil,omitempty" name:"TargetRegionInfo"`

	// <p>网络计费相关参数</p>
	InternetChargeInfo *InternetAccessible `json:"InternetChargeInfo,omitnil,omitempty" name:"InternetChargeInfo"`

	// <p>Target是否放通来自CLB的流量。<br>开启放通（true）：只验证CLB上的安全组；<br>不开启放通（false）：需同时验证CLB和后端实例上的安全组。<br>不填则不修改。</p>
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitnil,omitempty" name:"LoadBalancerPassToTarget"`

	// <p>不同计费模式之间的切换：0表示不切换，1表示预付费和后付费切换，2表示后付费之间切换。默认值：0</p>
	SwitchFlag *uint64 `json:"SwitchFlag,omitnil,omitempty" name:"SwitchFlag"`

	// <p>是否开启跨地域绑定2.0功能。不填则不修改。</p>
	SnatPro *bool `json:"SnatPro,omitnil,omitempty" name:"SnatPro"`

	// <p>是否开启删除保护，不填则不修改。</p>
	DeleteProtect *bool `json:"DeleteProtect,omitnil,omitempty" name:"DeleteProtect"`

	// <p>将负载均衡二级域名由mycloud.com改为tencentclb.com，子域名也会变换，修改后mycloud.com域名将失效。不填则不修改。</p>
	ModifyClassicDomain *bool `json:"ModifyClassicDomain,omitnil,omitempty" name:"ModifyClassicDomain"`

	// <p>关联的终端节点Id，可通过<a href="https://cloud.tencent.com/document/product/215/54679">DescribeVpcEndPoint</a>接口查询。传空字符串代表解除关联。</p>
	AssociateEndpoint *string `json:"AssociateEndpoint,omitnil,omitempty" name:"AssociateEndpoint"`
}

type ModifyLoadBalancerAttributesRequest struct {
	*tchttp.BaseRequest
	
	// <p>负载均衡的唯一ID，可以通过 <a href="https://cloud.tencent.com/document/product/214/30685">DescribeLoadBalancers</a> 接口获取。</p>
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// <p>负载均衡实例名称，规则：1-80 个英文字母、汉字等国际通用语言字符，数字，连接线“-”、下划线“_”等常见字符（禁止Unicode补充字符，如emoji表情、生僻汉字等）。</p>
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// <p>设置负载均衡跨地域绑定1.0的后端服务信息</p>
	TargetRegionInfo *TargetRegionInfo `json:"TargetRegionInfo,omitnil,omitempty" name:"TargetRegionInfo"`

	// <p>网络计费相关参数</p>
	InternetChargeInfo *InternetAccessible `json:"InternetChargeInfo,omitnil,omitempty" name:"InternetChargeInfo"`

	// <p>Target是否放通来自CLB的流量。<br>开启放通（true）：只验证CLB上的安全组；<br>不开启放通（false）：需同时验证CLB和后端实例上的安全组。<br>不填则不修改。</p>
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitnil,omitempty" name:"LoadBalancerPassToTarget"`

	// <p>不同计费模式之间的切换：0表示不切换，1表示预付费和后付费切换，2表示后付费之间切换。默认值：0</p>
	SwitchFlag *uint64 `json:"SwitchFlag,omitnil,omitempty" name:"SwitchFlag"`

	// <p>是否开启跨地域绑定2.0功能。不填则不修改。</p>
	SnatPro *bool `json:"SnatPro,omitnil,omitempty" name:"SnatPro"`

	// <p>是否开启删除保护，不填则不修改。</p>
	DeleteProtect *bool `json:"DeleteProtect,omitnil,omitempty" name:"DeleteProtect"`

	// <p>将负载均衡二级域名由mycloud.com改为tencentclb.com，子域名也会变换，修改后mycloud.com域名将失效。不填则不修改。</p>
	ModifyClassicDomain *bool `json:"ModifyClassicDomain,omitnil,omitempty" name:"ModifyClassicDomain"`

	// <p>关联的终端节点Id，可通过<a href="https://cloud.tencent.com/document/product/215/54679">DescribeVpcEndPoint</a>接口查询。传空字符串代表解除关联。</p>
	AssociateEndpoint *string `json:"AssociateEndpoint,omitnil,omitempty" name:"AssociateEndpoint"`
}

func (r *ModifyLoadBalancerAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "LoadBalancerName")
	delete(f, "TargetRegionInfo")
	delete(f, "InternetChargeInfo")
	delete(f, "LoadBalancerPassToTarget")
	delete(f, "SwitchFlag")
	delete(f, "SnatPro")
	delete(f, "DeleteProtect")
	delete(f, "ModifyClassicDomain")
	delete(f, "AssociateEndpoint")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerAttributesResponseParams struct {
	// <p>切换负载均衡计费方式时，可用此参数查询切换任务是否成功。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLoadBalancerAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancerAttributesResponseParams `json:"Response"`
}

func (r *ModifyLoadBalancerAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerMixIpTargetRequestParams struct {
	// 负载均衡实例ID数组，默认支持20个负载均衡实例ID。
	// 可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 开启/关闭IPv6FullChain负载均衡7层监听器支持混绑IPv4/IPv6目标特性。
	MixIpTarget *bool `json:"MixIpTarget,omitnil,omitempty" name:"MixIpTarget"`
}

type ModifyLoadBalancerMixIpTargetRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID数组，默认支持20个负载均衡实例ID。
	// 可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 开启/关闭IPv6FullChain负载均衡7层监听器支持混绑IPv4/IPv6目标特性。
	MixIpTarget *bool `json:"MixIpTarget,omitnil,omitempty" name:"MixIpTarget"`
}

func (r *ModifyLoadBalancerMixIpTargetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerMixIpTargetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	delete(f, "MixIpTarget")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerMixIpTargetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerMixIpTargetResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLoadBalancerMixIpTargetResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancerMixIpTargetResponseParams `json:"Response"`
}

func (r *ModifyLoadBalancerMixIpTargetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerMixIpTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerSlaRequestParams struct {
	// 负载均衡实例信息。
	LoadBalancerSla []*SlaUpdateParam `json:"LoadBalancerSla,omitnil,omitempty" name:"LoadBalancerSla"`

	// 是否强制升级，默认否。
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

type ModifyLoadBalancerSlaRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例信息。
	LoadBalancerSla []*SlaUpdateParam `json:"LoadBalancerSla,omitnil,omitempty" name:"LoadBalancerSla"`

	// 是否强制升级，默认否。
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

func (r *ModifyLoadBalancerSlaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerSlaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerSla")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerSlaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerSlaResponseParams struct {
	// 订单号。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLoadBalancerSlaResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancerSlaResponseParams `json:"Response"`
}

func (r *ModifyLoadBalancerSlaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerSlaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancersProjectRequestParams struct {
	// 一个或多个待操作的负载均衡实例ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	// 列表支持最大长度为20。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 项目ID。可以通过 [DescribeProject](https://cloud.tencent.com/document/api/651/78725) 接口获取。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type ModifyLoadBalancersProjectRequest struct {
	*tchttp.BaseRequest
	
	// 一个或多个待操作的负载均衡实例ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	// 列表支持最大长度为20。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 项目ID。可以通过 [DescribeProject](https://cloud.tencent.com/document/api/651/78725) 接口获取。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *ModifyLoadBalancersProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancersProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancersProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancersProjectResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLoadBalancersProjectResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancersProjectResponseParams `json:"Response"`
}

func (r *ModifyLoadBalancersProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancersProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelAliasAttributesRequestParams struct {
	// <p>模型积分系数配置。</p><p>必填，至少包含 <code>InputCoefficient</code>、<code>InputCachedCoefficient</code>、<code>OutputCoefficient</code> 中的一个字段，未传字段保持原值。</p><p><code>InputCoefficient</code> 为非缓存命中输入积分系数。</p><p><code>InputCachedCoefficient</code> 为缓存命中输入积分系数，用于 provider prompt cache 命中的输入 token。</p><p><code>OutputCoefficient</code> 为输出积分系数。</p><p>各字段取值范围：[0, 5000]，仅支持整数，0 表示该类 token 不计积分。</p>
	Coefficient *Coefficient `json:"Coefficient,omitnil,omitempty" name:"Coefficient"`

	// <p>模型别名列表。</p><p>不传 <code>ServiceProviderIds</code>（按 ModelAlias 账号维度修改）时支持数组批量，同一份 Coefficient 应用到多个别名。</p><p>传入 <code>ServiceProviderIds</code>（按 ServiceProvider 维度修改）时只能传 1 个别名，锁定唯一 model 别名；去重后不等于 1 个将返回 InvalidParameter。</p>
	ModelAliasNames []*string `json:"ModelAliasNames,omitnil,omitempty" name:"ModelAliasNames"`

	// <p>BYOK 实例（ServiceProvider）ID 列表。</p><p>可选，数组。传入时按 ServiceProvider 维度修改：把同一份 Coefficient 批量应用到数组内每一个实例（覆盖配置，仅作用于这些实例），此时 <code>ModelAliasNames</code> 只能传 1 个别名（即 1 别名 × N ServiceProvider）；数组需去重、非空、上限 100，任一实例不归属/不存在/该实例下无该别名将整批返回错误。不传时按 ModelAlias（账号）维度修改，作用于该别名下未单独配置覆盖的全部实例。</p>
	ServiceProviderIds []*string `json:"ServiceProviderIds,omitnil,omitempty" name:"ServiceProviderIds"`
}

type ModifyModelAliasAttributesRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型积分系数配置。</p><p>必填，至少包含 <code>InputCoefficient</code>、<code>InputCachedCoefficient</code>、<code>OutputCoefficient</code> 中的一个字段，未传字段保持原值。</p><p><code>InputCoefficient</code> 为非缓存命中输入积分系数。</p><p><code>InputCachedCoefficient</code> 为缓存命中输入积分系数，用于 provider prompt cache 命中的输入 token。</p><p><code>OutputCoefficient</code> 为输出积分系数。</p><p>各字段取值范围：[0, 5000]，仅支持整数，0 表示该类 token 不计积分。</p>
	Coefficient *Coefficient `json:"Coefficient,omitnil,omitempty" name:"Coefficient"`

	// <p>模型别名列表。</p><p>不传 <code>ServiceProviderIds</code>（按 ModelAlias 账号维度修改）时支持数组批量，同一份 Coefficient 应用到多个别名。</p><p>传入 <code>ServiceProviderIds</code>（按 ServiceProvider 维度修改）时只能传 1 个别名，锁定唯一 model 别名；去重后不等于 1 个将返回 InvalidParameter。</p>
	ModelAliasNames []*string `json:"ModelAliasNames,omitnil,omitempty" name:"ModelAliasNames"`

	// <p>BYOK 实例（ServiceProvider）ID 列表。</p><p>可选，数组。传入时按 ServiceProvider 维度修改：把同一份 Coefficient 批量应用到数组内每一个实例（覆盖配置，仅作用于这些实例），此时 <code>ModelAliasNames</code> 只能传 1 个别名（即 1 别名 × N ServiceProvider）；数组需去重、非空、上限 100，任一实例不归属/不存在/该实例下无该别名将整批返回错误。不传时按 ModelAlias（账号）维度修改，作用于该别名下未单独配置覆盖的全部实例。</p>
	ServiceProviderIds []*string `json:"ServiceProviderIds,omitnil,omitempty" name:"ServiceProviderIds"`
}

func (r *ModifyModelAliasAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelAliasAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Coefficient")
	delete(f, "ModelAliasNames")
	delete(f, "ServiceProviderIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModelAliasAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelAliasAttributesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyModelAliasAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModelAliasAttributesResponseParams `json:"Response"`
}

func (r *ModifyModelAliasAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelAliasAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelAttributesRequestParams struct {
	// <p>BYOK的ID</p><p>参数格式：byok-kot39u7j</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>BYOK的自定义名字</p><p>入参限制：1～256个字符</p>
	ServiceProviderName *string `json:"ServiceProviderName,omitnil,omitempty" name:"ServiceProviderName"`
}

type ModifyModelAttributesRequest struct {
	*tchttp.BaseRequest
	
	// <p>BYOK的ID</p><p>参数格式：byok-kot39u7j</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>BYOK的自定义名字</p><p>入参限制：1～256个字符</p>
	ServiceProviderName *string `json:"ServiceProviderName,omitnil,omitempty" name:"ServiceProviderName"`
}

func (r *ModifyModelAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceProviderId")
	delete(f, "ServiceProviderName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModelAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelAttributesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyModelAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModelAttributesResponseParams `json:"Response"`
}

func (r *ModifyModelAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelRouterAttributesRequestParams struct {
	// <p>模型路由ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>新的 HTTPS 证书ID，用于替换实例 HTTPS 服务端点当前绑定的证书。常用于证书到期前的更换场景。</p><p>使用限制：</p><ul><li>仅企业型（Enterprise）且服务端点协议为 HTTPS 的实例支持修改证书。</li><li>证书须为 SSL 证书控制台中状态为“已签发”（可用）且未过期的服务器证书（SVR 类型）。可在 <a href="https://console.cloud.tencent.com/ssl">SSL 证书控制台</a> 查看证书ID。</li><li>替换后新证书立即生效，过程中不会中断业务流量。</li><li>若传入的证书与当前绑定的证书相同，接口直接返回成功，不做任何变更。</li></ul><p>不传则证书保持不变。可通过 <code>DescribeModelRouterDetail</code> 接口的 <code>ServiceEndPoints.CertId</code> 字段查询当前绑定的证书。</p>
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// <p>模型路由名称</p>
	ModelRouterName *string `json:"ModelRouterName,omitnil,omitempty" name:"ModelRouterName"`

	// <p>限速配置</p>
	RateLimitConfig *RateLimitConfigForModelRouter `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`

	// <p>路由配置</p>
	RouterSetting *RouterSettingWithFallBack `json:"RouterSetting,omitnil,omitempty" name:"RouterSetting"`
}

type ModifyModelRouterAttributesRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>新的 HTTPS 证书ID，用于替换实例 HTTPS 服务端点当前绑定的证书。常用于证书到期前的更换场景。</p><p>使用限制：</p><ul><li>仅企业型（Enterprise）且服务端点协议为 HTTPS 的实例支持修改证书。</li><li>证书须为 SSL 证书控制台中状态为“已签发”（可用）且未过期的服务器证书（SVR 类型）。可在 <a href="https://console.cloud.tencent.com/ssl">SSL 证书控制台</a> 查看证书ID。</li><li>替换后新证书立即生效，过程中不会中断业务流量。</li><li>若传入的证书与当前绑定的证书相同，接口直接返回成功，不做任何变更。</li></ul><p>不传则证书保持不变。可通过 <code>DescribeModelRouterDetail</code> 接口的 <code>ServiceEndPoints.CertId</code> 字段查询当前绑定的证书。</p>
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// <p>模型路由名称</p>
	ModelRouterName *string `json:"ModelRouterName,omitnil,omitempty" name:"ModelRouterName"`

	// <p>限速配置</p>
	RateLimitConfig *RateLimitConfigForModelRouter `json:"RateLimitConfig,omitnil,omitempty" name:"RateLimitConfig"`

	// <p>路由配置</p>
	RouterSetting *RouterSettingWithFallBack `json:"RouterSetting,omitnil,omitempty" name:"RouterSetting"`
}

func (r *ModifyModelRouterAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelRouterAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "CertId")
	delete(f, "ModelRouterName")
	delete(f, "RateLimitConfig")
	delete(f, "RouterSetting")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModelRouterAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelRouterAttributesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyModelRouterAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModelRouterAttributesResponseParams `json:"Response"`
}

func (r *ModifyModelRouterAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelRouterAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelRouterGuardrailsRequestParams struct {
	// <p>待修改的 Guardrail 防护配置列表。</p><p>当前最多支持 1 个元素。每个元素必须填写 GuardrailId；当 Type 为 WAF 或未传按 WAF 处理时，InstanceId 和 ServiceId 必填；InputCheckDepth 为选填，不传时沿用当前已关联 Guardrail 的取值。</p>
	Guardrails []*GuardrailConfig `json:"Guardrails,omitnil,omitempty" name:"Guardrails"`

	// <p>模型路由实例 ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`
}

type ModifyModelRouterGuardrailsRequest struct {
	*tchttp.BaseRequest
	
	// <p>待修改的 Guardrail 防护配置列表。</p><p>当前最多支持 1 个元素。每个元素必须填写 GuardrailId；当 Type 为 WAF 或未传按 WAF 处理时，InstanceId 和 ServiceId 必填；InputCheckDepth 为选填，不传时沿用当前已关联 Guardrail 的取值。</p>
	Guardrails []*GuardrailConfig `json:"Guardrails,omitnil,omitempty" name:"Guardrails"`

	// <p>模型路由实例 ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`
}

func (r *ModifyModelRouterGuardrailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelRouterGuardrailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Guardrails")
	delete(f, "ModelRouterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModelRouterGuardrailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelRouterGuardrailsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyModelRouterGuardrailsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModelRouterGuardrailsResponseParams `json:"Response"`
}

func (r *ModifyModelRouterGuardrailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelRouterGuardrailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelRouterSecurityGroupsRequestParams struct {
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>需要绑定的安全组ID列表</p>
	SecurityGroups []*string `json:"SecurityGroups,omitnil,omitempty" name:"SecurityGroups"`
}

type ModifyModelRouterSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>需要绑定的安全组ID列表</p>
	SecurityGroups []*string `json:"SecurityGroups,omitnil,omitempty" name:"SecurityGroups"`
}

func (r *ModifyModelRouterSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelRouterSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "SecurityGroups")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyModelRouterSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyModelRouterSecurityGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyModelRouterSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *ModifyModelRouterSecurityGroupsResponseParams `json:"Response"`
}

func (r *ModifyModelRouterSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyModelRouterSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleRequestParams struct {
	// 负载均衡实例 ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口获取。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID，可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口获取。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 要修改的转发规则的 ID， 可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口获取。
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 转发规则的新的转发路径，如不需修改Url，则不需提供此参数。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 健康检查信息。
	HealthCheck *HealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// 规则的请求转发方式，可选值：WRR、LEAST_CONN、IP_HASH
	// 分别表示按权重轮询、最小连接数、按IP哈希， 默认为 WRR。
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// 会话保持时间。取值范围0或30-86400（单位：秒）。
	// 默认为0。
	SessionExpireTime *int64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// 负载均衡实例与后端服务之间的转发协议，默认HTTP，可取值：HTTP、HTTPS、GRPC。仅HTTPS监听器该参数有效。
	ForwardType *string `json:"ForwardType,omitnil,omitempty" name:"ForwardType"`

	// TRPC被调服务器路由，ForwardType为TRPC时必填。目前暂未对外开放。
	TrpcCallee *string `json:"TrpcCallee,omitnil,omitempty" name:"TrpcCallee"`

	// TRPC调用服务接口，ForwardType为TRPC时必填。目前暂未对外开放。
	TrpcFunc *string `json:"TrpcFunc,omitnil,omitempty" name:"TrpcFunc"`

	// OAuth配置信息。
	OAuth *OAuth `json:"OAuth,omitnil,omitempty" name:"OAuth"`

	// 自定义cookie名
	CookieName *string `json:"CookieName,omitnil,omitempty" name:"CookieName"`
}

type ModifyRuleRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/214/30685) 接口获取。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID，可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口获取。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 要修改的转发规则的 ID， 可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口获取。
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 转发规则的新的转发路径，如不需修改Url，则不需提供此参数。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 健康检查信息。
	HealthCheck *HealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// 规则的请求转发方式，可选值：WRR、LEAST_CONN、IP_HASH
	// 分别表示按权重轮询、最小连接数、按IP哈希， 默认为 WRR。
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// 会话保持时间。取值范围0或30-86400（单位：秒）。
	// 默认为0。
	SessionExpireTime *int64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// 负载均衡实例与后端服务之间的转发协议，默认HTTP，可取值：HTTP、HTTPS、GRPC。仅HTTPS监听器该参数有效。
	ForwardType *string `json:"ForwardType,omitnil,omitempty" name:"ForwardType"`

	// TRPC被调服务器路由，ForwardType为TRPC时必填。目前暂未对外开放。
	TrpcCallee *string `json:"TrpcCallee,omitnil,omitempty" name:"TrpcCallee"`

	// TRPC调用服务接口，ForwardType为TRPC时必填。目前暂未对外开放。
	TrpcFunc *string `json:"TrpcFunc,omitnil,omitempty" name:"TrpcFunc"`

	// OAuth配置信息。
	OAuth *OAuth `json:"OAuth,omitnil,omitempty" name:"OAuth"`

	// 自定义cookie名
	CookieName *string `json:"CookieName,omitnil,omitempty" name:"CookieName"`
}

func (r *ModifyRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "LocationId")
	delete(f, "Url")
	delete(f, "HealthCheck")
	delete(f, "Scheduler")
	delete(f, "SessionExpireTime")
	delete(f, "ForwardType")
	delete(f, "TrpcCallee")
	delete(f, "TrpcFunc")
	delete(f, "OAuth")
	delete(f, "CookieName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRuleResponseParams `json:"Response"`
}

func (r *ModifyRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceProviderModelAttributesRequestParams struct {
	// <p>BYOK 实例 ID</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>待修改的模型的名称（原始模型名称）</p>
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// <p>该模型支持的输入多模态能力列表</p><p>枚举值：</p><ul><li>text： 支持文本输入</li><li>image： 支持图像输入</li><li>file： 支持文件输入（当前仅支持pdf）</li></ul>
	InputModalities []*string `json:"InputModalities,omitnil,omitempty" name:"InputModalities"`
}

type ModifyServiceProviderModelAttributesRequest struct {
	*tchttp.BaseRequest
	
	// <p>BYOK 实例 ID</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>待修改的模型的名称（原始模型名称）</p>
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// <p>该模型支持的输入多模态能力列表</p><p>枚举值：</p><ul><li>text： 支持文本输入</li><li>image： 支持图像输入</li><li>file： 支持文件输入（当前仅支持pdf）</li></ul>
	InputModalities []*string `json:"InputModalities,omitnil,omitempty" name:"InputModalities"`
}

func (r *ModifyServiceProviderModelAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceProviderModelAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceProviderId")
	delete(f, "ModelName")
	delete(f, "InputModalities")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyServiceProviderModelAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyServiceProviderModelAttributesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyServiceProviderModelAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyServiceProviderModelAttributesResponseParams `json:"Response"`
}

func (r *ModifyServiceProviderModelAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyServiceProviderModelAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetGroupAttributeRequestParams struct {
	// <p>目标组的ID。</p>
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// <p>目标组的新名称。命名规则：1-80 个英文字母、汉字等国际通用语言字符，数字，连接线“-”、下划线“_”等常见字符（禁止Unicode补充字符，如emoji表情、生僻汉字等）。</p>
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`

	// <p>目标组的新默认端口。全监听目标组不支持此参数。</p>
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>调度算法，仅V2新版目标组，且后端转发协议为(HTTP|HTTPS|GRPC)时该参数有效。可选值：<br>&lt;ur&gt;<li>WRR:按权重轮询。</li><li>LEAST_CONN:最小连接数。</li><li>IP_HASH:按IP哈希。</li><li>默认为 WRR。</li>&lt;ur&gt;</p>
	ScheduleAlgorithm *string `json:"ScheduleAlgorithm,omitnil,omitempty" name:"ScheduleAlgorithm"`

	// <p>健康检查详情。</p>
	HealthCheck *TargetGroupHealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// <p>后端服务默认权重, 其中：<ul><li>取值范围[0, 100]</li><li>设置该值后，添加后端服务到目标组时， 若后端服务不单独设置权重， 则使用这里的默认权重。 </li><li>v1目标组类型不支持设置Weight参数。</li> </ul></p>
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// <p>是否开启长连接，此参数仅适用于HTTP/HTTPS目标组，true:关闭；false:开启， 默认关闭。</p>
	KeepaliveEnable *bool `json:"KeepaliveEnable,omitnil,omitempty" name:"KeepaliveEnable"`

	// <p>会话保持时间，单位：秒。可选值：30~3600，默认 0，表示不开启。TCP/UDP目标组不支持该参数。</p>
	SessionExpireTime *uint64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// <p>是否开启SNAT（源IP替换），True（开启）、False（关闭）。默认为关闭。注意：SnatEnable开启时会替换客户端源IP，此时<code>透传客户端源IP</code>选项关闭，反之亦然。</p>
	SnatEnable *bool `json:"SnatEnable,omitnil,omitempty" name:"SnatEnable"`
}

type ModifyTargetGroupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// <p>目标组的ID。</p>
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// <p>目标组的新名称。命名规则：1-80 个英文字母、汉字等国际通用语言字符，数字，连接线“-”、下划线“_”等常见字符（禁止Unicode补充字符，如emoji表情、生僻汉字等）。</p>
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`

	// <p>目标组的新默认端口。全监听目标组不支持此参数。</p>
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>调度算法，仅V2新版目标组，且后端转发协议为(HTTP|HTTPS|GRPC)时该参数有效。可选值：<br>&lt;ur&gt;<li>WRR:按权重轮询。</li><li>LEAST_CONN:最小连接数。</li><li>IP_HASH:按IP哈希。</li><li>默认为 WRR。</li>&lt;ur&gt;</p>
	ScheduleAlgorithm *string `json:"ScheduleAlgorithm,omitnil,omitempty" name:"ScheduleAlgorithm"`

	// <p>健康检查详情。</p>
	HealthCheck *TargetGroupHealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// <p>后端服务默认权重, 其中：<ul><li>取值范围[0, 100]</li><li>设置该值后，添加后端服务到目标组时， 若后端服务不单独设置权重， 则使用这里的默认权重。 </li><li>v1目标组类型不支持设置Weight参数。</li> </ul></p>
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// <p>是否开启长连接，此参数仅适用于HTTP/HTTPS目标组，true:关闭；false:开启， 默认关闭。</p>
	KeepaliveEnable *bool `json:"KeepaliveEnable,omitnil,omitempty" name:"KeepaliveEnable"`

	// <p>会话保持时间，单位：秒。可选值：30~3600，默认 0，表示不开启。TCP/UDP目标组不支持该参数。</p>
	SessionExpireTime *uint64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// <p>是否开启SNAT（源IP替换），True（开启）、False（关闭）。默认为关闭。注意：SnatEnable开启时会替换客户端源IP，此时<code>透传客户端源IP</code>选项关闭，反之亦然。</p>
	SnatEnable *bool `json:"SnatEnable,omitnil,omitempty" name:"SnatEnable"`
}

func (r *ModifyTargetGroupAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetGroupAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "TargetGroupName")
	delete(f, "Port")
	delete(f, "ScheduleAlgorithm")
	delete(f, "HealthCheck")
	delete(f, "Weight")
	delete(f, "KeepaliveEnable")
	delete(f, "SessionExpireTime")
	delete(f, "SnatEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetGroupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetGroupAttributeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTargetGroupAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTargetGroupAttributeResponseParams `json:"Response"`
}

func (r *ModifyTargetGroupAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetGroupInstancesPortRequestParams struct {
	// 目标组ID。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 待修改端口的服务器数组，在这个接口 NewPort 和 Port 为必填项。
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
}

type ModifyTargetGroupInstancesPortRequest struct {
	*tchttp.BaseRequest
	
	// 目标组ID。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 待修改端口的服务器数组，在这个接口 NewPort 和 Port 为必填项。
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
}

func (r *ModifyTargetGroupInstancesPortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetGroupInstancesPortRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "TargetGroupInstances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetGroupInstancesPortRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetGroupInstancesPortResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTargetGroupInstancesPortResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTargetGroupInstancesPortResponseParams `json:"Response"`
}

func (r *ModifyTargetGroupInstancesPortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetGroupInstancesPortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetGroupInstancesWeightRequestParams struct {
	// 目标组ID。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 待修改权重的服务器数组，在这个接口 Port 为必填项。
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
}

type ModifyTargetGroupInstancesWeightRequest struct {
	*tchttp.BaseRequest
	
	// 目标组ID。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 待修改权重的服务器数组，在这个接口 Port 为必填项。
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
}

func (r *ModifyTargetGroupInstancesWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetGroupInstancesWeightRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "TargetGroupInstances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetGroupInstancesWeightRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetGroupInstancesWeightResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTargetGroupInstancesWeightResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTargetGroupInstancesWeightResponseParams `json:"Response"`
}

func (r *ModifyTargetGroupInstancesWeightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetGroupInstancesWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetPortRequestParams struct {
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 要修改端口的后端服务列表。
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 后端服务绑定到监听器或转发规则的新端口。
	NewPort *int64 `json:"NewPort,omitnil,omitempty" name:"NewPort"`

	// 转发规则的ID，当后端服务绑定到七层转发规则时，必须提供此参数或Domain+Url两者之一。
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 目标规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 目标规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type ModifyTargetPortRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 要修改端口的后端服务列表。
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 后端服务绑定到监听器或转发规则的新端口。
	NewPort *int64 `json:"NewPort,omitnil,omitempty" name:"NewPort"`

	// 转发规则的ID，当后端服务绑定到七层转发规则时，必须提供此参数或Domain+Url两者之一。
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 目标规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 目标规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

func (r *ModifyTargetPortRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetPortRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "Targets")
	delete(f, "NewPort")
	delete(f, "LocationId")
	delete(f, "Domain")
	delete(f, "Url")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetPortRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetPortResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTargetPortResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTargetPortResponseParams `json:"Response"`
}

func (r *ModifyTargetPortResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetPortResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetWeightRequestParams struct {
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 要修改权重的后端服务列表。
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 转发规则的ID，当绑定机器到七层转发规则时，必须提供此参数或Domain+Url两者之一。
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 目标规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 目标规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 后端服务新的转发权重，取值范围：0~100，默认值10。如果设置了 Targets.Weight 参数，则此参数不生效。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type ModifyTargetWeightRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 要修改权重的后端服务列表。
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 转发规则的ID，当绑定机器到七层转发规则时，必须提供此参数或Domain+Url两者之一。
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 目标规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 目标规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 后端服务新的转发权重，取值范围：0~100，默认值10。如果设置了 Targets.Weight 参数，则此参数不生效。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

func (r *ModifyTargetWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetWeightRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "Targets")
	delete(f, "LocationId")
	delete(f, "Domain")
	delete(f, "Url")
	delete(f, "Weight")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetWeightRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetWeightResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTargetWeightResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTargetWeightResponseParams `json:"Response"`
}

func (r *ModifyTargetWeightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserGroupAttributesRequestParams struct {
	// <p>模型路由实例ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>待修改的用户组ID。不可为「未分组」虚拟分组 ugrp-ungrouped。</p>
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// <p>用户组关联的预算ID。不传则不修改预算关联；传入有效 budget-xxx 则将该用户组关联到此预算（若已关联其它预算则替换为本预算）。仅支持关联/替换，不支持解绑——解绑请用 DisassociateBudget。预算与组内 Key、所属实例的预算各自独立判定。</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>用户组意图路由白名单（ir-xxx）。每一项须为该实例已创建的意图路由名。不传则不修改；传入即整体覆盖。</p>
	IntentRouters []*string `json:"IntentRouters,omitnil,omitempty" name:"IntentRouters"`

	// <p>用户组真实模型白名单。每一项须为该实例已关联的模型名。不传则不修改；传入即整体覆盖。</p>
	Models []*string `json:"Models,omitnil,omitempty" name:"Models"`

	// <p>用户组名称。不传则不修改；传入时长度不超过255个字符、同实例下唯一。</p>
	UserGroupName *string `json:"UserGroupName,omitnil,omitempty" name:"UserGroupName"`
}

type ModifyUserGroupAttributesRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>待修改的用户组ID。不可为「未分组」虚拟分组 ugrp-ungrouped。</p>
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// <p>用户组关联的预算ID。不传则不修改预算关联；传入有效 budget-xxx 则将该用户组关联到此预算（若已关联其它预算则替换为本预算）。仅支持关联/替换，不支持解绑——解绑请用 DisassociateBudget。预算与组内 Key、所属实例的预算各自独立判定。</p>
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>用户组意图路由白名单（ir-xxx）。每一项须为该实例已创建的意图路由名。不传则不修改；传入即整体覆盖。</p>
	IntentRouters []*string `json:"IntentRouters,omitnil,omitempty" name:"IntentRouters"`

	// <p>用户组真实模型白名单。每一项须为该实例已关联的模型名。不传则不修改；传入即整体覆盖。</p>
	Models []*string `json:"Models,omitnil,omitempty" name:"Models"`

	// <p>用户组名称。不传则不修改；传入时长度不超过255个字符、同实例下唯一。</p>
	UserGroupName *string `json:"UserGroupName,omitnil,omitempty" name:"UserGroupName"`
}

func (r *ModifyUserGroupAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserGroupAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "UserGroupId")
	delete(f, "BudgetId")
	delete(f, "IntentRouters")
	delete(f, "Models")
	delete(f, "UserGroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserGroupAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserGroupAttributesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserGroupAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserGroupAttributesResponseParams `json:"Response"`
}

func (r *ModifyUserGroupAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserGroupAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MultiCertInfo struct {
	// 认证类型，UNIDIRECTIONAL：单向认证，MUTUAL：双向认证
	SSLMode *string `json:"SSLMode,omitnil,omitempty" name:"SSLMode"`

	// 监听器或规则证书列表，单双向认证，多本服务端证书算法类型不能重复;若SSLMode为双向认证，证书列表必须包含一本ca证书。
	CertList []*CertInfo `json:"CertList,omitnil,omitempty" name:"CertList"`

	// 双向认证时，是否开启客户端认证，ON:开启，OPTIONAL:自适应，默认ON
	SSLVerifyClient *string `json:"SSLVerifyClient,omitnil,omitempty" name:"SSLVerifyClient"`
}

type MultiModalityAttachments struct {
	// <p>base64 url编码的文件内容</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>附件类型</p><p>枚举值：</p><ul><li>image： 图像</li><li>pdf： pdf（文件）</li></ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type OAuth struct {
	// 开启或关闭鉴权。
	// True: 开启;
	// False: 关闭
	// 默认为关闭。
	OAuthEnable *bool `json:"OAuthEnable,omitnil,omitempty" name:"OAuthEnable"`

	// IAP全部故障后，拒绝请求还是放行。
	// BYPASS: 通过
	// REJECT: 拒绝
	// 默认为 BYPASS
	OAuthFailureStatus *string `json:"OAuthFailureStatus,omitnil,omitempty" name:"OAuthFailureStatus"`
}

type Price struct {
	// 描述了实例价格。
	InstancePrice *ItemPrice `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`

	// 描述了网络价格。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthPrice *ItemPrice `json:"BandwidthPrice,omitnil,omitempty" name:"BandwidthPrice"`

	// 描述了lcu价格。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LcuPrice *ItemPrice `json:"LcuPrice,omitnil,omitempty" name:"LcuPrice"`
}

type ProviderItem struct {
	// <p>Provider 标识（如 openai）</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>显示名称（如 OpenAI）</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>模型协议列表</p>
	Protocols []*string `json:"Protocols,omitnil,omitempty" name:"Protocols"`

	// <p>英文显示名称</p>
	EnglishDisplayName *string `json:"EnglishDisplayName,omitnil,omitempty" name:"EnglishDisplayName"`
}

type ProviderTestConnectionErrorInfo struct {
	// <p>上游模型侧返回的HTTP状态码</p>
	HttpCode *uint64 `json:"HttpCode,omitnil,omitempty" name:"HttpCode"`

	// <p>错误状态码</p>
	ErrorStatus *string `json:"ErrorStatus,omitnil,omitempty" name:"ErrorStatus"`

	// <p>探测请求错误信息</p>
	OriginalMessage *string `json:"OriginalMessage,omitnil,omitempty" name:"OriginalMessage"`
}

type Quota struct {
	// 配额名称，取值范围：
	// <li> TOTAL_OPEN_CLB_QUOTA：用户当前地域下的公网CLB配额 </li>
	// <li> TOTAL_INTERNAL_CLB_QUOTA：用户当前地域下的内网CLB配额 </li>
	// <li> TOTAL_LISTENER_QUOTA：一个CLB下的监听器配额 </li>
	// <li> TOTAL_LISTENER_RULE_QUOTA：一个监听器下的转发规则配额 </li>
	// <li> TOTAL_TARGET_BIND_QUOTA：一条转发规则下可绑定设备的配额 </li>
	// <li> TOTAL_SNAT_IP_QUOTA： 一个CLB实例下跨地域2.0的SNAT IP配额 </li>
	// <li>TOTAL_ISP_CLB_QUOTA：用户当前地域下的三网CLB配额 </li>
	// <li>TOTAL_FULL_PORT_RANGE_LISTENER_QUOTA：一个CLB实例下的单个协议全端口段监听器配额</li>
	QuotaId *string `json:"QuotaId,omitnil,omitempty" name:"QuotaId"`

	// 当前使用数量，为 null 时表示无意义。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuotaCurrent *int64 `json:"QuotaCurrent,omitnil,omitempty" name:"QuotaCurrent"`

	// 配额数量。
	QuotaLimit *int64 `json:"QuotaLimit,omitnil,omitempty" name:"QuotaLimit"`
}

type RateLimitConfigForBudget struct {
	// <p>每分钟限制的请求数量。</p><p>单位：次/分钟。</p>
	RPM *uint64 `json:"RPM,omitnil,omitempty" name:"RPM"`

	// <p>每分钟限制的Token数量。</p><p>单位：个/分钟。</p>
	TPM *uint64 `json:"TPM,omitnil,omitempty" name:"TPM"`
}

type RateLimitConfigForKey struct {
	// <p>最大并发请求数量</p><p>单位：次</p>
	MaxParallelRequest *uint64 `json:"MaxParallelRequest,omitnil,omitempty" name:"MaxParallelRequest"`

	// <p>每分钟限制的请求数量</p><p>单位：次/分钟</p>
	RPM *uint64 `json:"RPM,omitnil,omitempty" name:"RPM"`

	// <p>每分钟限制的Token数量</p><p>单位：个/分钟</p>
	TPM *uint64 `json:"TPM,omitnil,omitempty" name:"TPM"`
}

type RateLimitConfigForModelRouter struct {
	// <p>每分钟限制的请求数量</p><p>单位：次/分钟</p>
	RPM *uint64 `json:"RPM,omitnil,omitempty" name:"RPM"`

	// <p>每分钟限制的Token数量</p><p>单位：个/分钟</p>
	TPM *uint64 `json:"TPM,omitnil,omitempty" name:"TPM"`
}

// Predefined struct for user
type RefundModelRouterResourcePackageRequestParams struct {
	// <p>待退还的模型路由资源包Id</p><p>非有效状态或者设置了自动续订且自动续订已生效的资源包不允许退款。</p>
	ModelRouterResourcePackageIds []*string `json:"ModelRouterResourcePackageIds,omitnil,omitempty" name:"ModelRouterResourcePackageIds"`
}

type RefundModelRouterResourcePackageRequest struct {
	*tchttp.BaseRequest
	
	// <p>待退还的模型路由资源包Id</p><p>非有效状态或者设置了自动续订且自动续订已生效的资源包不允许退款。</p>
	ModelRouterResourcePackageIds []*string `json:"ModelRouterResourcePackageIds,omitnil,omitempty" name:"ModelRouterResourcePackageIds"`
}

func (r *RefundModelRouterResourcePackageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundModelRouterResourcePackageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterResourcePackageIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RefundModelRouterResourcePackageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RefundModelRouterResourcePackageResponseParams struct {
	// <p>退还模型路由资源包的订单号</p>
	DealNames []*string `json:"DealNames,omitnil,omitempty" name:"DealNames"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RefundModelRouterResourcePackageResponse struct {
	*tchttp.BaseResponse
	Response *RefundModelRouterResourcePackageResponseParams `json:"Response"`
}

func (r *RefundModelRouterResourcePackageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RefundModelRouterResourcePackageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegenerateKeysRequestParams struct {
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>Key的ID列表</p>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

type RegenerateKeysRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例ID</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>Key的ID列表</p>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

func (r *RegenerateKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegenerateKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "KeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegenerateKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegenerateKeysResponseParams struct {
	// <p>重新生成失败的Key的ID列表</p>
	FailedKeyIds []*string `json:"FailedKeyIds,omitnil,omitempty" name:"FailedKeyIds"`

	// <p>重新生成后的Key的信息</p>
	RegeneratedKeys []*RegeneratedKey `json:"RegeneratedKeys,omitnil,omitempty" name:"RegeneratedKeys"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RegenerateKeysResponse struct {
	*tchttp.BaseResponse
	Response *RegenerateKeysResponseParams `json:"Response"`
}

func (r *RegenerateKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegenerateKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegeneratedKey struct {
	// <p>重新生成的明文Key</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>Key的ID</p>
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

// Predefined struct for user
type RegisterFunctionTargetsRequestParams struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 待绑定的云函数列表。
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitnil,omitempty" name:"FunctionTargets"`

	// 目标转发规则的 ID，当将云函数绑定到七层转发规则时，必须输入此参数或 Domain+Url 参数。
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 目标转发规则的域名，若已经输入 LocationId 参数，则本参数不生效。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 目标转发规则的 URL，若已经输入 LocationId 参数，则本参数不生效。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type RegisterFunctionTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 待绑定的云函数列表。
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitnil,omitempty" name:"FunctionTargets"`

	// 目标转发规则的 ID，当将云函数绑定到七层转发规则时，必须输入此参数或 Domain+Url 参数。
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 目标转发规则的域名，若已经输入 LocationId 参数，则本参数不生效。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 目标转发规则的 URL，若已经输入 LocationId 参数，则本参数不生效。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

func (r *RegisterFunctionTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterFunctionTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "FunctionTargets")
	delete(f, "LocationId")
	delete(f, "Domain")
	delete(f, "Url")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterFunctionTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterFunctionTargetsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RegisterFunctionTargetsResponse struct {
	*tchttp.BaseResponse
	Response *RegisterFunctionTargetsResponseParams `json:"Response"`
}

func (r *RegisterFunctionTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterFunctionTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterModelsToServiceProviderRequestParams struct {
	// <p>BYOK的ID</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>需要关联的模型信息</p>
	Models []*ModelItem `json:"Models,omitnil,omitempty" name:"Models"`
}

type RegisterModelsToServiceProviderRequest struct {
	*tchttp.BaseRequest
	
	// <p>BYOK的ID</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>需要关联的模型信息</p>
	Models []*ModelItem `json:"Models,omitnil,omitempty" name:"Models"`
}

func (r *RegisterModelsToServiceProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterModelsToServiceProviderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceProviderId")
	delete(f, "Models")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterModelsToServiceProviderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterModelsToServiceProviderResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RegisterModelsToServiceProviderResponse struct {
	*tchttp.BaseResponse
	Response *RegisterModelsToServiceProviderResponseParams `json:"Response"`
}

func (r *RegisterModelsToServiceProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterModelsToServiceProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterTargetGroupInstancesRequestParams struct {
	// 目标组ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 服务器实例数组，服务器和目标组的 VPC 需相同。
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
}

type RegisterTargetGroupInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 目标组ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 服务器实例数组，服务器和目标组的 VPC 需相同。
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitnil,omitempty" name:"TargetGroupInstances"`
}

func (r *RegisterTargetGroupInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterTargetGroupInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "TargetGroupInstances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterTargetGroupInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterTargetGroupInstancesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RegisterTargetGroupInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RegisterTargetGroupInstancesResponseParams `json:"Response"`
}

func (r *RegisterTargetGroupInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterTargetGroupInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterTargetsRequestParams struct {
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 待绑定的后端服务列表，数组长度最大支持20。
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 转发规则的ID，可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口获取，当绑定后端服务到七层转发规则时，必须提供此参数或Domain+Url两者之一。
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 目标转发规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 目标转发规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type RegisterTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 待绑定的后端服务列表，数组长度最大支持20。
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 转发规则的ID，可以通过 [DescribeListeners](https://cloud.tencent.com/document/product/214/30686) 接口获取，当绑定后端服务到七层转发规则时，必须提供此参数或Domain+Url两者之一。
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 目标转发规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 目标转发规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

func (r *RegisterTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ListenerId")
	delete(f, "Targets")
	delete(f, "LocationId")
	delete(f, "Domain")
	delete(f, "Url")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterTargetsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RegisterTargetsResponse struct {
	*tchttp.BaseResponse
	Response *RegisterTargetsResponseParams `json:"Response"`
}

func (r *RegisterTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterTargetsWithClassicalLBRequestParams struct {
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 后端服务信息。
	Targets []*ClassicalTargetInfo `json:"Targets,omitnil,omitempty" name:"Targets"`
}

type RegisterTargetsWithClassicalLBRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 后端服务信息。
	Targets []*ClassicalTargetInfo `json:"Targets,omitnil,omitempty" name:"Targets"`
}

func (r *RegisterTargetsWithClassicalLBRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterTargetsWithClassicalLBRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "Targets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterTargetsWithClassicalLBRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterTargetsWithClassicalLBResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RegisterTargetsWithClassicalLBResponse struct {
	*tchttp.BaseResponse
	Response *RegisterTargetsWithClassicalLBResponseParams `json:"Response"`
}

func (r *RegisterTargetsWithClassicalLBResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterTargetsWithClassicalLBResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveModelKeyRequestParams struct {
	// <p>服务提供商ID</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>Key 业务 ID 列表，至少 1 个，最多 10 个</p>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

type RemoveModelKeyRequest struct {
	*tchttp.BaseRequest
	
	// <p>服务提供商ID</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>Key 业务 ID 列表，至少 1 个，最多 10 个</p>
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

func (r *RemoveModelKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveModelKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceProviderId")
	delete(f, "KeyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveModelKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveModelKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveModelKeyResponse struct {
	*tchttp.BaseResponse
	Response *RemoveModelKeyResponseParams `json:"Response"`
}

func (r *RemoveModelKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveModelKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveModelRewriteRequestParams struct {
	// <p>模型路由实例 ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>要删除的源模型名（重写规则的 key）。</p><p>长度 1-255 字符；支持特殊值 <code>default</code> 表示删除兜底规则。</p><p>当指定的 SourceModel 当前不存在重写规则时，请求幂等成功。</p>
	SourceModel *string `json:"SourceModel,omitnil,omitempty" name:"SourceModel"`
}

type RemoveModelRewriteRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型路由实例 ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>要删除的源模型名（重写规则的 key）。</p><p>长度 1-255 字符；支持特殊值 <code>default</code> 表示删除兜底规则。</p><p>当指定的 SourceModel 当前不存在重写规则时，请求幂等成功。</p>
	SourceModel *string `json:"SourceModel,omitnil,omitempty" name:"SourceModel"`
}

func (r *RemoveModelRewriteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveModelRewriteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelRouterId")
	delete(f, "SourceModel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveModelRewriteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveModelRewriteResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveModelRewriteResponse struct {
	*tchttp.BaseResponse
	Response *RemoveModelRewriteResponseParams `json:"Response"`
}

func (r *RemoveModelRewriteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveModelRewriteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewLoadBalancersRequestParams struct {
	// 负载均衡实例唯一ID数组，最多支持20个。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 负载均衡实例的预付费相关属性。
	LBChargePrepaid *LBChargePrepaid `json:"LBChargePrepaid,omitnil,omitempty" name:"LBChargePrepaid"`
}

type RenewLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例唯一ID数组，最多支持20个。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 负载均衡实例的预付费相关属性。
	LBChargePrepaid *LBChargePrepaid `json:"LBChargePrepaid,omitnil,omitempty" name:"LBChargePrepaid"`
}

func (r *RenewLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	delete(f, "LBChargePrepaid")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewLoadBalancersResponseParams struct {
	// 订单号。
	DealName *string `json:"DealName,omitnil,omitempty" name:"DealName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *RenewLoadBalancersResponseParams `json:"Response"`
}

func (r *RenewLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceCertForLoadBalancersRequestParams struct {
	// 需要被替换的证书的ID，可以是服务端证书或客户端证书
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// 新证书的内容等相关信息
	Certificate *CertificateInput `json:"Certificate,omitnil,omitempty" name:"Certificate"`
}

type ReplaceCertForLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// 需要被替换的证书的ID，可以是服务端证书或客户端证书
	OldCertificateId *string `json:"OldCertificateId,omitnil,omitempty" name:"OldCertificateId"`

	// 新证书的内容等相关信息
	Certificate *CertificateInput `json:"Certificate,omitnil,omitempty" name:"Certificate"`
}

func (r *ReplaceCertForLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceCertForLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OldCertificateId")
	delete(f, "Certificate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReplaceCertForLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReplaceCertForLoadBalancersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReplaceCertForLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *ReplaceCertForLoadBalancersResponseParams `json:"Response"`
}

func (r *ReplaceCertForLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReplaceCertForLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Resource struct {
	// 运营商内具体资源信息，如"CMCC", "CUCC", "CTCC", "BGP", "INTERNAL"。
	Type []*string `json:"Type,omitnil,omitempty" name:"Type"`

	// 运营商信息，如"CMCC", "CUCC", "CTCC", "BGP", "INTERNAL"。
	Isp *string `json:"Isp,omitnil,omitempty" name:"Isp"`

	// 可用资源。
	AvailabilitySet []*ResourceAvailability `json:"AvailabilitySet,omitnil,omitempty" name:"AvailabilitySet"`

	// 运营商类型信息
	TypeSet []*TypeInfo `json:"TypeSet,omitnil,omitempty" name:"TypeSet"`
}

type ResourceAvailability struct {
	// 运营商内具体资源信息，如"CMCC", "CUCC", "CTCC", "BGP"。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 资源可用性，"Available"：可用，"Unavailable"：不可用
	Availability *string `json:"Availability,omitnil,omitempty" name:"Availability"`
}

type RewriteItem struct {
	// <p>源模型名（重写规则的 key）。</p><p>特殊值 <code>default</code> 表示兜底规则（命中所有未显式列出的源模型）。</p>
	SourceModel *string `json:"SourceModel,omitnil,omitempty" name:"SourceModel"`

	// <p>目标模型名（重写规则的 value）。</p>
	TargetModel *string `json:"TargetModel,omitnil,omitempty" name:"TargetModel"`
}

type RewriteLocationMap struct {
	// 源转发规则ID
	SourceLocationId *string `json:"SourceLocationId,omitnil,omitempty" name:"SourceLocationId"`

	// 重定向目标转发规则的ID
	TargetLocationId *string `json:"TargetLocationId,omitnil,omitempty" name:"TargetLocationId"`

	// 重定向状态码，可取值301,302,307
	RewriteCode *int64 `json:"RewriteCode,omitnil,omitempty" name:"RewriteCode"`

	// 重定向是否携带匹配的url，配置RewriteCode时必填
	TakeUrl *bool `json:"TakeUrl,omitnil,omitempty" name:"TakeUrl"`

	// 源转发的域名，必须是SourceLocationId对应的域名，配置RewriteCode时必填
	SourceDomain *string `json:"SourceDomain,omitnil,omitempty" name:"SourceDomain"`
}

type RewriteTarget struct {
	// 重定向目标的监听器ID，该字段仅配置了重定向时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetListenerId *string `json:"TargetListenerId,omitnil,omitempty" name:"TargetListenerId"`

	// 重定向目标的转发规则ID，该字段仅配置了重定向时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetLocationId *string `json:"TargetLocationId,omitnil,omitempty" name:"TargetLocationId"`

	// 重定向状态码
	RewriteCode *int64 `json:"RewriteCode,omitnil,omitempty" name:"RewriteCode"`

	// 重定向是否携带匹配的url
	TakeUrl *bool `json:"TakeUrl,omitnil,omitempty" name:"TakeUrl"`

	// 重定向类型，Manual: 手动重定向，Auto:  自动重定向
	RewriteType *string `json:"RewriteType,omitnil,omitempty" name:"RewriteType"`
}

type RouterSettingWithFallBack struct {
	// <p>模型间路由策略。</p><p>枚举值：</p><ul><li>SimpleShuffle： 简单随机路由</li><li>LowestCost： 最低积分路由</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrossModelGroupRoutingStrategy *string `json:"CrossModelGroupRoutingStrategy,omitnil,omitempty" name:"CrossModelGroupRoutingStrategy"`

	// <p>回退策略</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FallBack *FallBackItem `json:"FallBack,omitnil,omitempty" name:"FallBack"`

	// <p>模型内路由策略</p><p>枚举值：</p><ul><li>SimpleShuffle： 简单随机路由</li><li>LeastBusy： 最低繁忙路由</li><li>LatencyBasedRouting： 最低延迟路由</li><li>UsageBasedRouting： 用量均衡路由</li><li>CostBasedRouting： 最低积分路由</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoutingStrategy *string `json:"RoutingStrategy,omitnil,omitempty" name:"RoutingStrategy"`
}

type RouterSettingWithoutFallBack struct {
	// <p>路由策略</p><p>枚举值：</p><ul><li>SimpleShuffle： 简单随机路由</li><li>LeastBusy： 最低繁忙路由</li><li>LatencyBasedRouting： 最低延迟路由</li><li>UsageBasedRouting： 用量均衡路由</li><li>CostBasedRouting： 最低积分路由</li></ul>
	RoutingStrategy *string `json:"RoutingStrategy,omitnil,omitempty" name:"RoutingStrategy"`
}

type RsTagRule struct {
	// <p>负载均衡监听器 ID。</p>
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// <p>要修改标签的后端机器列表。</p>
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// <p>转发规则的ID，七层规则时需要此参数，4层规则不需要。</p>
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// <p>后端服务修改后的标签。此参数的优先级低于前述<a href="https://cloud.tencent.com/document/api/214/30694#Target">Target</a>中的Tag参数，即最终的标签以Target中的Tag参数值为准，仅当Target中的Tag参数为空时，才以RsTagRule中的Tag参数为准。</p>
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`
}

type RsWeightRule struct {
	// 负载均衡监听器 ID。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 要修改权重的后端机器列表。
	Targets []*Target `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 转发规则的ID，七层规则时需要此参数，4层规则不需要。
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 目标规则的域名，提供LocationId参数时本参数不生效。
	//
	// Deprecated: Domain is deprecated.
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 目标规则的URL，提供LocationId参数时本参数不生效。
	//
	// Deprecated: Url is deprecated.
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 后端服务修改后的转发权重，取值范围：[0，100]。此参数的优先级低于前述[Target](https://cloud.tencent.com/document/api/214/30694#Target)中的Weight参数，即最终的权重值以Target中的Weight参数值为准，仅当Target中的Weight参数为空时，才以RsWeightRule中的Weight参数为准。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type RuleHealth struct {
	// 转发规则ID
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 转发规则的域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 转发规则的Url
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 高级路由规则ID
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 本规则上绑定的后端服务的健康检查状态
	Targets []*TargetHealth `json:"Targets,omitnil,omitempty" name:"Targets"`
}

type RuleInput struct {
	// 转发规则的路径。长度限制为：1~200。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 转发规则的域名。长度限制为：1~80。Domain和Domains只需要传一个，单域名规则传Domain，多域名规则传Domains。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 会话保持时间。设置为0表示关闭会话保持，开启会话保持可取值30~86400，单位：秒。
	SessionExpireTime *int64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// 健康检查信息。详情请参见：[健康检查](https://cloud.tencent.com/document/product/214/6097)
	HealthCheck *HealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// 证书信息；此参数和MultiCertInfo不能同时传入。
	Certificate *CertificateInput `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// 规则的请求转发方式，可选值：WRR、LEAST_CONN、IP_HASH
	// 分别表示按权重轮询、最小连接数、按IP哈希， 默认为 WRR。
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// 负载均衡与后端服务之间的转发协议，目前支持 HTTP/HTTPS/GRPC/GRPCS/TRPC，TRPC暂未对外开放，默认HTTP。
	ForwardType *string `json:"ForwardType,omitnil,omitempty" name:"ForwardType"`

	// 是否将该域名设为默认域名，注意，一个监听器下只能设置一个默认域名。
	DefaultServer *bool `json:"DefaultServer,omitnil,omitempty" name:"DefaultServer"`

	// 是否开启Http2，注意，只有HTTPS域名才能开启Http2。
	Http2 *bool `json:"Http2,omitnil,omitempty" name:"Http2"`

	// 后端目标类型，NODE表示绑定普通节点，TARGETGROUP表示绑定目标组
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// TRPC被调服务器路由，ForwardType为TRPC时必填。目前暂未对外开放。
	TrpcCallee *string `json:"TrpcCallee,omitnil,omitempty" name:"TrpcCallee"`

	// TRPC调用服务接口，ForwardType为TRPC时必填。目前暂未对外开放
	TrpcFunc *string `json:"TrpcFunc,omitnil,omitempty" name:"TrpcFunc"`

	// 是否开启QUIC，注意，只有HTTPS域名才能开启QUIC
	Quic *bool `json:"Quic,omitnil,omitempty" name:"Quic"`

	// 转发规则的域名列表。每个域名的长度限制为：1~80。Domain和Domains只需要传一个，单域名规则传Domain，多域名规则传Domains。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 证书信息，支持同时传入不同算法类型的多本服务端证书；此参数和Certificate不能同时传入。
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitnil,omitempty" name:"MultiCertInfo"`

	// 自定义cookie名
	CookieName *string `json:"CookieName,omitnil,omitempty" name:"CookieName"`
}

type RuleOutput struct {
	// 转发规则的 ID
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 转发规则的域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 转发规则的路径。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 会话保持时间
	SessionExpireTime *int64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// 健康检查信息
	HealthCheck *HealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// 证书信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Certificate *CertificateOutput `json:"Certificate,omitnil,omitempty" name:"Certificate"`

	// 规则的请求转发方式。
	// WRR、LEAST_CONN、IP_HASH分别表示按权重轮询、最小连接数、IP Hash。
	Scheduler *string `json:"Scheduler,omitnil,omitempty" name:"Scheduler"`

	// 转发规则所属的监听器 ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 转发规则的重定向目标信息
	RewriteTarget *RewriteTarget `json:"RewriteTarget,omitnil,omitempty" name:"RewriteTarget"`

	// 是否开启gzip
	HttpGzip *bool `json:"HttpGzip,omitnil,omitempty" name:"HttpGzip"`

	// 转发规则是否为自动创建
	BeAutoCreated *bool `json:"BeAutoCreated,omitnil,omitempty" name:"BeAutoCreated"`

	// 是否作为默认域名
	DefaultServer *bool `json:"DefaultServer,omitnil,omitempty" name:"DefaultServer"`

	// 是否开启Http2
	Http2 *bool `json:"Http2,omitnil,omitempty" name:"Http2"`

	// 负载均衡与后端服务之间的转发协议
	ForwardType *string `json:"ForwardType,omitnil,omitempty" name:"ForwardType"`

	// 转发规则的创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 后端服务器类型。NODE表示绑定普通节点，TARGETGROUP表示绑定目标组。
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 绑定的目标组基本信息；当规则绑定目标组时，会返回该字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetGroup *BasicTargetGroupInfo `json:"TargetGroup,omitnil,omitempty" name:"TargetGroup"`

	// WAF实例ID
	WafDomainId *string `json:"WafDomainId,omitnil,omitempty" name:"WafDomainId"`

	// TRPC被调服务器路由，ForwardType为TRPC时有效。目前暂未对外开放。
	TrpcCallee *string `json:"TrpcCallee,omitnil,omitempty" name:"TrpcCallee"`

	// TRPC调用服务接口，ForwardType为TRPC时有效。目前暂未对外开放。
	TrpcFunc *string `json:"TrpcFunc,omitnil,omitempty" name:"TrpcFunc"`

	// QUIC状态。QUIC_ACTIVE表示开启，QUIC_INACTIVE表示未开启。注意，只有HTTPS域名才能开启QUIC。
	QuicStatus *string `json:"QuicStatus,omitnil,omitempty" name:"QuicStatus"`

	// 转发规则的域名列表。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 绑定的目标组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetGroupList []*BasicTargetGroupInfo `json:"TargetGroupList,omitnil,omitempty" name:"TargetGroupList"`

	// OAuth配置状态信息。
	OAuth *OAuth `json:"OAuth,omitnil,omitempty" name:"OAuth"`

	// 自定义cookie名。
	CookieName *string `json:"CookieName,omitnil,omitempty" name:"CookieName"`
}

type RuleTargets struct {
	// 转发规则的 ID
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 转发规则的域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 转发规则的路径。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 后端服务的信息
	Targets []*Backend `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 后端云函数的信息
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitnil,omitempty" name:"FunctionTargets"`
}

type RulesItems struct {
	// 规则id
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// uri
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 绑定的后端对象
	Targets []*LbRsTargets `json:"Targets,omitnil,omitempty" name:"Targets"`
}

type ServiceEndPoints struct {
	// <p>证书ID</p>
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// <p>监听端口</p>
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>网络协议</p>
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`
}

type ServiceProvider struct {
	// <p>BYOK类型</p>
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// <p>单个byok实例下该模型可支持的输入多模态能力列表。</p><p>枚举值：</p><ul><li>text： 支持文本输入</li><li>file： 支持文件输入（当前仅支持pdf）</li><li>image： 支持图像输入</li></ul>
	InputModalities []*string `json:"InputModalities,omitnil,omitempty" name:"InputModalities"`

	// <p>模型协议</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>BYOK的所属厂商</p>
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// <p>BYOK实例ID</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>BYOK名称</p>
	ServiceProviderName *string `json:"ServiceProviderName,omitnil,omitempty" name:"ServiceProviderName"`
}

type ServiceProviderCoefficient struct {
	// <p>该 BYOK 实例（ServiceProvider）维度的积分系数。</p><p>可选字段：仅当该实例单独配置了 ServiceProvider 维度系数时返回，返回值即该实例的生效系数；未返回时表示该实例继承所属 ModelAlias 的 <code>Coefficient</code>。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Coefficient *Coefficient `json:"Coefficient,omitnil,omitempty" name:"Coefficient"`

	// <p>BYOK 实例（ServiceProvider）ID。</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>BYOK 实例（ServiceProvider）名称。</p>
	ServiceProviderName *string `json:"ServiceProviderName,omitnil,omitempty" name:"ServiceProviderName"`
}

type ServiceProviderItem struct {
	// <p>服务提供商 ID</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>用户自定义服务提供商名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceProviderName *string `json:"ServiceProviderName,omitnil,omitempty" name:"ServiceProviderName"`

	// <p>模型供应商</p>
	ModelProvider *string `json:"ModelProvider,omitnil,omitempty" name:"ModelProvider"`

	// <p>该byok实例下该模型可支持的输入多模态能力列表。</p><p>枚举值：</p><ul><li>text： 支持文本输入</li><li>file： 支持文件输入（当前仅支持pdf）</li><li>image： 支持图像输入</li></ul>
	InputModalities []*string `json:"InputModalities,omitnil,omitempty" name:"InputModalities"`
}

type ServiceProviderModelItem struct {
	// <p>关联的模型路由实例列表</p>
	AssociatedModelRouters []*AssociatedModelRouterItem `json:"AssociatedModelRouters,omitnil,omitempty" name:"AssociatedModelRouters"`

	// <p>该模型当前支持的输入多模态能力列表</p><p>枚举值：</p><ul><li>text： 支持文本输入</li><li>image： 支持图像输入</li><li>file： 支持文件输入（当前仅支持pdf）</li></ul><p>默认值：text</p>
	InputModalities []*string `json:"InputModalities,omitnil,omitempty" name:"InputModalities"`

	// <p>模型别名, 可以用于实际访问</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAlias *string `json:"ModelAlias,omitnil,omitempty" name:"ModelAlias"`

	// <p>模型唯一标识, 原始模型名称</p>
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// <p>该模型经探测最多支持的输入多模态能力列表</p><p>枚举值：</p><ul><li>text： 支持文本输入</li><li>file： 支持文件输入（当前仅支持pdf）</li><li>image： 支持图像输入</li></ul><p>模型不健康时列表为空</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProbedInputModalities []*string `json:"ProbedInputModalities,omitnil,omitempty" name:"ProbedInputModalities"`
}

// Predefined struct for user
type SetCustomizedConfigForLoadBalancerRequestParams struct {
	// 操作类型。
	// - ADD：创建
	// - DELETE：删除
	// - UPDATE：修改
	// - BIND：绑定
	// - UNBIND：解绑
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 个性化配置ID。除了创建个性化配置外，必传此字段，如：pz-1234abcd
	UconfigId *string `json:"UconfigId,omitnil,omitempty" name:"UconfigId"`

	// 个性化配置内容。创建个性化配置或修改个性化配置的内容时，必传此字段。
	// 具体限制查看 [七层个性化配置](https://cloud.tencent.com/document/product/214/15171)
	ConfigContent *string `json:"ConfigContent,omitnil,omitempty" name:"ConfigContent"`

	// 个性化配置名称。创建个性化配置或修改个性化配置的名字时，必传此字段。
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// 负载均衡实例ID。绑定解绑时，必传此字段。
	// 可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 标签
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type SetCustomizedConfigForLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 操作类型。
	// - ADD：创建
	// - DELETE：删除
	// - UPDATE：修改
	// - BIND：绑定
	// - UNBIND：解绑
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 个性化配置ID。除了创建个性化配置外，必传此字段，如：pz-1234abcd
	UconfigId *string `json:"UconfigId,omitnil,omitempty" name:"UconfigId"`

	// 个性化配置内容。创建个性化配置或修改个性化配置的内容时，必传此字段。
	// 具体限制查看 [七层个性化配置](https://cloud.tencent.com/document/product/214/15171)
	ConfigContent *string `json:"ConfigContent,omitnil,omitempty" name:"ConfigContent"`

	// 个性化配置名称。创建个性化配置或修改个性化配置的名字时，必传此字段。
	ConfigName *string `json:"ConfigName,omitnil,omitempty" name:"ConfigName"`

	// 负载均衡实例ID。绑定解绑时，必传此字段。
	// 可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 标签
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *SetCustomizedConfigForLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCustomizedConfigForLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OperationType")
	delete(f, "UconfigId")
	delete(f, "ConfigContent")
	delete(f, "ConfigName")
	delete(f, "LoadBalancerIds")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetCustomizedConfigForLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetCustomizedConfigForLoadBalancerResponseParams struct {
	// 个性化配置ID，如：pz-1234abcd
	ConfigId *string `json:"ConfigId,omitnil,omitempty" name:"ConfigId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetCustomizedConfigForLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *SetCustomizedConfigForLoadBalancerResponseParams `json:"Response"`
}

func (r *SetCustomizedConfigForLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetCustomizedConfigForLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetLoadBalancerClsLogRequestParams struct {
	// 负载均衡实例 ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 日志服务(CLS)的日志集 ID。
	// <li>增加和更新日志主题时可调用 [DescribeLogsets](https://cloud.tencent.com/document/product/614/58624) 接口获取日志集 ID。</li>
	// <li>删除日志主题时，此参数填写为**空字符串**即可。</li>
	LogSetId *string `json:"LogSetId,omitnil,omitempty" name:"LogSetId"`

	// 日志服务(CLS)的日志主题 ID。
	// <li>增加和更新日志主题时可调用 [DescribeTopics](https://cloud.tencent.com/document/product/614/56454) 接口获取日志主题 ID。</li>
	// <li>删除日志主题时，此参数填写为**空字符串**即可。</li>
	LogTopicId *string `json:"LogTopicId,omitnil,omitempty" name:"LogTopicId"`

	// 日志类型：
	// <li>ACCESS：访问日志</li>
	// <li>HEALTH：健康检查日志</li>
	// 默认为ACCESS。
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`
}

type SetLoadBalancerClsLogRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 日志服务(CLS)的日志集 ID。
	// <li>增加和更新日志主题时可调用 [DescribeLogsets](https://cloud.tencent.com/document/product/614/58624) 接口获取日志集 ID。</li>
	// <li>删除日志主题时，此参数填写为**空字符串**即可。</li>
	LogSetId *string `json:"LogSetId,omitnil,omitempty" name:"LogSetId"`

	// 日志服务(CLS)的日志主题 ID。
	// <li>增加和更新日志主题时可调用 [DescribeTopics](https://cloud.tencent.com/document/product/614/56454) 接口获取日志主题 ID。</li>
	// <li>删除日志主题时，此参数填写为**空字符串**即可。</li>
	LogTopicId *string `json:"LogTopicId,omitnil,omitempty" name:"LogTopicId"`

	// 日志类型：
	// <li>ACCESS：访问日志</li>
	// <li>HEALTH：健康检查日志</li>
	// 默认为ACCESS。
	LogType *string `json:"LogType,omitnil,omitempty" name:"LogType"`
}

func (r *SetLoadBalancerClsLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetLoadBalancerClsLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "LogSetId")
	delete(f, "LogTopicId")
	delete(f, "LogType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetLoadBalancerClsLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetLoadBalancerClsLogResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetLoadBalancerClsLogResponse struct {
	*tchttp.BaseResponse
	Response *SetLoadBalancerClsLogResponseParams `json:"Response"`
}

func (r *SetLoadBalancerClsLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetLoadBalancerClsLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetLoadBalancerSecurityGroupsRequestParams struct {
	// 负载均衡实例 ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 安全组ID构成的数组，一个负载均衡实例最多可绑定50个安全组，如果要解绑所有安全组，可不传此参数。
	// 可以通过 [DescribeSecurityGroups](https://cloud.tencent.com/document/product/215/15808) 接口查询。
	SecurityGroups []*string `json:"SecurityGroups,omitnil,omitempty" name:"SecurityGroups"`
}

type SetLoadBalancerSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 安全组ID构成的数组，一个负载均衡实例最多可绑定50个安全组，如果要解绑所有安全组，可不传此参数。
	// 可以通过 [DescribeSecurityGroups](https://cloud.tencent.com/document/product/215/15808) 接口查询。
	SecurityGroups []*string `json:"SecurityGroups,omitnil,omitempty" name:"SecurityGroups"`
}

func (r *SetLoadBalancerSecurityGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetLoadBalancerSecurityGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "SecurityGroups")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetLoadBalancerSecurityGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetLoadBalancerSecurityGroupsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetLoadBalancerSecurityGroupsResponse struct {
	*tchttp.BaseResponse
	Response *SetLoadBalancerSecurityGroupsResponseParams `json:"Response"`
}

func (r *SetLoadBalancerSecurityGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetLoadBalancerSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetLoadBalancerStartStatusRequestParams struct {
	// 操作类型。Start：启动实例，Stop：停止实例。
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 负载均衡实例ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 监听器ID。如果该字段为空，则表示操作负载均衡实例，如果不为空，则表示操作监听器。
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`
}

type SetLoadBalancerStartStatusRequest struct {
	*tchttp.BaseRequest
	
	// 操作类型。Start：启动实例，Stop：停止实例。
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 负载均衡实例ID，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 监听器ID。如果该字段为空，则表示操作负载均衡实例，如果不为空，则表示操作监听器。
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`
}

func (r *SetLoadBalancerStartStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetLoadBalancerStartStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OperationType")
	delete(f, "LoadBalancerId")
	delete(f, "ListenerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetLoadBalancerStartStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetLoadBalancerStartStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetLoadBalancerStartStatusResponse struct {
	*tchttp.BaseResponse
	Response *SetLoadBalancerStartStatusResponseParams `json:"Response"`
}

func (r *SetLoadBalancerStartStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetLoadBalancerStartStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetSecurityGroupForLoadbalancersRequestParams struct {
	// 安全组ID，如 sg-12345678。可以通过 [DescribeSecurityGroups](https://cloud.tencent.com/document/product/215/15808) 接口获取。
	SecurityGroup *string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// ADD 绑定安全组；
	// DEL 解绑安全组
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 负载均衡实例ID数组，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	// 列表支持的最大长度为20。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`
}

type SetSecurityGroupForLoadbalancersRequest struct {
	*tchttp.BaseRequest
	
	// 安全组ID，如 sg-12345678。可以通过 [DescribeSecurityGroups](https://cloud.tencent.com/document/product/215/15808) 接口获取。
	SecurityGroup *string `json:"SecurityGroup,omitnil,omitempty" name:"SecurityGroup"`

	// ADD 绑定安全组；
	// DEL 解绑安全组
	OperationType *string `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 负载均衡实例ID数组，可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	// 列表支持的最大长度为20。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`
}

func (r *SetSecurityGroupForLoadbalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetSecurityGroupForLoadbalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityGroup")
	delete(f, "OperationType")
	delete(f, "LoadBalancerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetSecurityGroupForLoadbalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetSecurityGroupForLoadbalancersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetSecurityGroupForLoadbalancersResponse struct {
	*tchttp.BaseResponse
	Response *SetSecurityGroupForLoadbalancersResponseParams `json:"Response"`
}

func (r *SetSecurityGroupForLoadbalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetSecurityGroupForLoadbalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SlaUpdateParam struct {
	// 负载均衡实例 ID。
	// 可以通过 [DescribeLoadBalancers](https://cloud.tencent.com/document/product/1108/48459) 接口查询。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 性能容量型规格，取值范围：
	// <li> clb.c2.medium：标准型规格 </li>
	// <li> clb.c3.small：高阶型1规格 </li>
	// <li> clb.c3.medium：高阶型2规格 </li>
	// <li> clb.c4.small：超强型1规格 </li>
	// <li> clb.c4.medium：超强型2规格 </li>
	// <li> clb.c4.large：超强型3规格 </li>
	// <li> clb.c4.xlarge：超强型4规格 </li>如需了解规格详情，请参见[实例规格对比](https://cloud.tencent.com/document/product/214/84689)
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`
}

type SnatIp struct {
	// 私有网络子网的唯一性id，如subnet-12345678
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// IP地址，如192.168.0.1
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`
}

type Sort struct {
	// <p>排序的字段</p>
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// <p>排序方式，支持ASC、DESC</p>
	Order *string `json:"Order,omitnil,omitempty" name:"Order"`
}

type SpecAvailability struct {
	// 规格类型。
	// <li>clb.c2.medium（标准型）</li><li>clb.c3.small（高阶型1）</li><li>clb.c3.medium（高阶型2）</li>
	// <li>clb.c4.small（超强型1）</li><li>clb.c4.medium（超强型2）</li><li>clb.c4.large（超强型3）</li><li>clb.c4.xlarge（超强型4）</li><li>shared（共享型）</li>
	SpecType *string `json:"SpecType,omitnil,omitempty" name:"SpecType"`

	// 规格可用性。资源可用性，"Available"：可用，"Unavailable"：不可用
	Availability *string `json:"Availability,omitnil,omitempty" name:"Availability"`
}

type TagInfo struct {
	// 标签的键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type Target struct {
	// 后端服务的监听端口。
	// 注意：绑定CVM（云服务器）或ENI（弹性网卡）时必传此参数
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 后端服务的类型，可取：CVM（云服务器）、ENI（弹性网卡）；作为入参时，目前本参数暂不生效。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 绑定CVM时需要传入此参数，代表CVM的唯一 ID，可通过 DescribeInstances 接口返回字段中的 InstanceId 字段获取。表示绑定主网卡主IPv4地址；以下场景都不支持指定InstanceId：绑定非CVM，绑定CVM上的辅助网卡IP，通过跨域2.0绑定CVM，以及绑定CVM的IPv6地址等。
	// 注意：参数 InstanceId、EniIp 有且只能传入其中一个参数。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 后端服务修改后的转发权重，取值范围：[0, 100]，默认为 10。此参数的优先级高于[RsWeightRule](https://cloud.tencent.com/document/api/214/30694#RsWeightRule)中的Weight参数，即最终的权重值以此Weight参数值为准，仅当此Weight参数为空时，才以RsWeightRule中的Weight参数为准。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 绑定IP时需要传入此参数，支持弹性网卡的IP和其他内网IP，如果是弹性网卡则必须先绑定至CVM，然后才能绑定到负载均衡实例。
	// 注意：参数 InstanceId、EniIp 有且只能传入其中一个参数。如果绑定双栈IPV6子机，则必须传该参数。如果是跨地域绑定，则必须传该参数，不支持传InstanceId参数。
	EniIp *string `json:"EniIp,omitnil,omitempty" name:"EniIp"`

	// 标签。
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`
}

type TargetGroupAssociation struct {
	// 负载均衡ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 目标组ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 监听器ID。访问AssociateTargetGroups和DisassociateTargetGroups接口时必传此参数。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 转发规则ID
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 目标组权重，范围[0, 100]。仅绑定v2目标组时生效，如果不存在，则默认为10。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type TargetGroupBackend struct {
	// 目标组ID
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 后端服务的类型，可取：CVM、ENI（即将支持）
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 后端服务的唯一 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 后端服务的监听端口，全端口段监听器此字段返回0，代表无效端口，即不支持设置。
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 后端服务的转发权重，取值范围：[0, 100]，默认为 10。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 后端服务的外网 IP
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// 后端服务的内网 IP
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitnil,omitempty" name:"PrivateIpAddresses"`

	// 后端服务的实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 后端服务被绑定的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegisteredTime *string `json:"RegisteredTime,omitnil,omitempty" name:"RegisteredTime"`

	// 弹性网卡唯一ID
	EniId *string `json:"EniId,omitnil,omitempty" name:"EniId"`

	// 后端服务的可用区ID
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`
}

type TargetGroupHealthCheck struct {
	// 是否开启健康检查。
	HealthSwitch *bool `json:"HealthSwitch,omitnil,omitempty" name:"HealthSwitch"`

	// 健康检查方式， 其中仅V2新版目标组类型支持该参数， 支持取值 TCP | HTTP | HTTPS | PING | CUSTOM，其中:
	// <ur><li>当目标组后端转发协议为TCP时， 健康检查方式支持 TCP/HTTP/CUSTOM， 默认为TCP。</li><li>当目标组后端转发协议为UDP时， 健康检查方式支持 PING/CUSTOM，默认为PING。</li><li>当目标组后端转发协议为HTTP时， 健康检查方式支持 HTTP/TCP， 默认为HTTP。</li><li>当目标组后端转发协议为HTTPS时， 健康检查方式支持 HTTPS/TCP， 默认为HTTPS。</li><li>当目标组后端转发协议为GRPC时， 健康检查方式支持GRPC/TCP， 默认为GRPC。</li></ur>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 自定义探测相关参数。健康检查端口，默认为后端服务的端口，除非您希望指定特定端口，否则建议留空。（仅适用于TCP/UDP目标组）。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 健康检查超时时间。 默认为2秒。 可配置范围：2 - 30秒。
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 检测间隔时间。 默认为5秒。 可配置范围：2 - 300秒。
	GapTime *int64 `json:"GapTime,omitnil,omitempty" name:"GapTime"`

	// 检测健康阈值。 默认为3秒。 可配置范围：2 - 10次。
	GoodLimit *int64 `json:"GoodLimit,omitnil,omitempty" name:"GoodLimit"`

	// 检测不健康阈值。 默认为3秒。 可配置范围：2 - 10次。
	BadLimit *int64 `json:"BadLimit,omitnil,omitempty" name:"BadLimit"`

	// 目标组下的所有rs的探测包是否开启巨帧。默认开启。仅GWLB类型目标组支持该参数。
	JumboFrame *bool `json:"JumboFrame,omitnil,omitempty" name:"JumboFrame"`

	// 健康检查状态码（仅适用于HTTP/HTTPS目标组、TCP目标组的HTTP健康检查方式）。可选值：1~31，默认 31，其中：<url> <li>1 表示探测后返回值 1xx 代表健康。</li><li>2 表示返回 2xx 代表健康。</li><li>4 表示返回 3xx 代表健康。</li><li>8 表示返回 4xx 代表健康。</li><li>16 表示返回 5xx 代表健康。</li></url>若希望多种返回码都可代表健康，则将相应的值相加。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpCode *int64 `json:"HttpCode,omitnil,omitempty" name:"HttpCode"`

	// 健康检查域名， 其中：<ur><li>仅适用于HTTP/HTTPS目标组和TCP目标组的HTTP健康检查方式。</li><li>针对HTTP/HTTPS目标组，当使用HTTP健康检查方式时，该参数为必填项。</li></ur>
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpCheckDomain *string `json:"HttpCheckDomain,omitnil,omitempty" name:"HttpCheckDomain"`

	// 健康检查路径（仅适用于HTTP/HTTPS转发规则、TCP监听器的HTTP健康检查方式）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpCheckPath *string `json:"HttpCheckPath,omitnil,omitempty" name:"HttpCheckPath"`

	// 健康检查方法（仅适用于HTTP/HTTPS转发规则、TCP监听器的HTTP健康检查方式），默认值：HEAD，可选值HEAD或GET。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpCheckMethod *string `json:"HttpCheckMethod,omitnil,omitempty" name:"HttpCheckMethod"`

	// 健康检查的输入格式，健康检查方式取CUSTOM时，必填此字段，可取值：HEX或TEXT，其中：<ur><li>TEXT：文本格式。</li><li>HEX：十六进制格式， SendContext和RecvContext的字符只能在0123456789ABCDEF中选取且长度必须是偶数位。</li><li>仅适用于TCP/UDP目标组。</li></ur>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContextType *string `json:"ContextType,omitnil,omitempty" name:"ContextType"`

	// 自定义探测相关参数。健康检查协议CheckType的值取CUSTOM时，必填此字段，代表健康检查发送的请求内容，只允许ASCII可见字符，最大长度限制500。（仅适用于TCP/UDP目标组）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SendContext *string `json:"SendContext,omitnil,omitempty" name:"SendContext"`

	// 自定义探测相关参数。健康检查协议CheckType的值取CUSTOM时，必填此字段，代表健康检查返回的结果，只允许ASCII可见字符，最大长度限制500。（仅适用于TCP/UDP目标组）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecvContext *string `json:"RecvContext,omitnil,omitempty" name:"RecvContext"`

	// HTTP版本, 其中：<ur><li>健康检查协议CheckType的值取HTTP时，必传此字段。</li><li>支持配置选项：HTTP/1.0, HTTP/1.1。</li><li>仅适用于TCP目标组。</li></ur>
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpVersion *string `json:"HttpVersion,omitnil,omitempty" name:"HttpVersion"`

	// GRPC健康检查状态码（仅适用于后端转发协议为GRPC的目标组）。默认值为 12，可输入值为数值、多个数值、或者范围，例如 20 或 20,25 或 0-99。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtendedCode *string `json:"ExtendedCode,omitnil,omitempty" name:"ExtendedCode"`
}

type TargetGroupInfo struct {
	// <p>目标组ID</p>
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// <p>目标组的vpcid</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>目标组的名字</p>
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`

	// <p>目标组的默认端口，全监听目标组此字段返回0，表示无效端口。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>目标组的创建时间</p>
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// <p>目标组的修改时间</p>
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// <p>关联到的规则数组。在DescribeTargetGroupList接口调用时无法获取到该参数。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociatedRule []*AssociationItem `json:"AssociatedRule,omitnil,omitempty" name:"AssociatedRule"`

	// <p>目标组后端转发协议, 仅v2新版目标组返回有效值。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>调度算法，仅后端转发协议为(HTTP、HTTPS、GRPC)的目标组返回有效值， 可选值：<br><ur></p><li>WRR:按权重轮询。</li><li>LEAST_CONN:最小连接数。</li><li>IP_HASH:按IP哈希。</li></ur>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleAlgorithm *string `json:"ScheduleAlgorithm,omitnil,omitempty" name:"ScheduleAlgorithm"`

	// <p>健康检查详情。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheck *TargetGroupHealthCheck `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// <p>目标组类型，当前支持v1(旧版目标组), v2(新版目标组)。默认为v1旧版目标组。</p>
	TargetGroupType *string `json:"TargetGroupType,omitnil,omitempty" name:"TargetGroupType"`

	// <p>目标组已关联的规则数。</p>
	AssociatedRuleCount *int64 `json:"AssociatedRuleCount,omitnil,omitempty" name:"AssociatedRuleCount"`

	// <p>目标组内的实例数量。</p>
	RegisteredInstancesCount *int64 `json:"RegisteredInstancesCount,omitnil,omitempty" name:"RegisteredInstancesCount"`

	// <p>标签。</p>
	Tag []*TagInfo `json:"Tag,omitnil,omitempty" name:"Tag"`

	// <p>默认权重。只有v2类型目标组返回该字段。当返回为NULL时， 表示未设置默认权重。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// <p>是否全监听目标组。</p>
	FullListenSwitch *bool `json:"FullListenSwitch,omitnil,omitempty" name:"FullListenSwitch"`

	// <p>是否开启长连接,  仅后端转发协议为HTTP/HTTPS/GRPC目标组返回有效值。</p>
	KeepaliveEnable *bool `json:"KeepaliveEnable,omitnil,omitempty" name:"KeepaliveEnable"`

	// <p>会话保持时间，仅后端转发协议为HTTP/HTTPS/GRPC目标组返回有效值。</p>
	SessionExpireTime *int64 `json:"SessionExpireTime,omitnil,omitempty" name:"SessionExpireTime"`

	// <p>IP版本。</p>
	IpVersion *string `json:"IpVersion,omitnil,omitempty" name:"IpVersion"`

	// <p>是否开启SNAT</p>
	SnatEnable *bool `json:"SnatEnable,omitnil,omitempty" name:"SnatEnable"`
}

type TargetGroupInstance struct {
	// 目标组实例的内网IP
	BindIP *string `json:"BindIP,omitnil,omitempty" name:"BindIP"`

	// 目标组实例的端口，全监听目标组不支持传此字段。
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 目标组实例的权重
	// v2目标组需要配置权重，调用CreateTargetGroup接口创建目标组时该参数与创建接口中的Weight参数必填其一。
	// 取值范围：0-100
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 目标组实例的新端口，全监听目标组不支持传此字段。
	NewPort *uint64 `json:"NewPort,omitnil,omitempty" name:"NewPort"`
}

type TargetGroupInstanceStatus struct {
	// 后端rs的IP
	InstanceIp *string `json:"InstanceIp,omitnil,omitempty" name:"InstanceIp"`

	// 健康检查状态，参数值及含义如下：
	// ● on：表示探测中。
	// ● off：表示健康检查关闭。
	// ● health：表示健康。
	// ● unhealth：表示异常。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 端口
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 网卡ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EniId *string `json:"EniId,omitnil,omitempty" name:"EniId"`
}

type TargetHealth struct {
	// Target的内网IP
	IP *string `json:"IP,omitnil,omitempty" name:"IP"`

	// Target绑定的端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 当前健康状态，true：健康，false：不健康（包括尚未开始探测、探测中、状态异常等几种状态）。只有处于健康状态（且权重大于0），负载均衡才会向其转发流量。
	HealthStatus *bool `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// Target的实例ID，如 ins-12345678
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 当前健康状态的详细信息。如：Alive、Dead、Unknown、Close。Alive状态为健康，Dead状态为异常，Unknown状态包括尚未开始探测、探测中、状态未知，Close表示健康检查关闭或监听器状态停止。
	HealthStatusDetail *string `json:"HealthStatusDetail,omitnil,omitempty" name:"HealthStatusDetail"`

	// (**该参数对象即将下线，不推荐使用，请使用HealthStatusDetail获取健康详情**) 当前健康状态的详细信息。如：Alive、Dead、Unknown。Alive状态为健康，Dead状态为异常，Unknown状态包括尚未开始探测、探测中、状态未知。
	//
	// Deprecated: HealthStatusDetial is deprecated.
	HealthStatusDetial *string `json:"HealthStatusDetial,omitnil,omitempty" name:"HealthStatusDetial"`

	// 目标组唯一ID。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// Target的权重。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type TargetRegionInfo struct {
	// Target所属地域，如 ap-guangzhou
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Target所属网络，私有网络格式如 vpc-abcd1234，如果是基础网络，则为"0"
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Target所属网络，私有网络格式如86323，如果是基础网络，则为0
	NumericalVpcId *int64 `json:"NumericalVpcId,omitnil,omitempty" name:"NumericalVpcId"`
}

type TestConnectionRequestInfo struct {
	// <p>请求URL</p>
	RequestUrl *string `json:"RequestUrl,omitnil,omitempty" name:"RequestUrl"`

	// <p>请求体</p>
	RequestBody *string `json:"RequestBody,omitnil,omitempty" name:"RequestBody"`

	// <p>请求头</p>
	RequestHeaders *string `json:"RequestHeaders,omitnil,omitempty" name:"RequestHeaders"`
}

// Predefined struct for user
type TestModelInputModalitiesRequestParams struct {
	// <p>待探测的模型（原始模型名称）</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>待探测的API Key（明文）</p>
	ProviderKey *string `json:"ProviderKey,omitnil,omitempty" name:"ProviderKey"`

	// <p>已创建的BYOK API Key ID（与ProviderKey二选一传入）</p>
	ProviderKeyId *string `json:"ProviderKeyId,omitnil,omitempty" name:"ProviderKeyId"`

	// <p>BYOK类型，当ProviderKey传入时必填</p>
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// <p>模型厂商协议，当ProviderKey传入时必填</p>
	ModelProtocol *string `json:"ModelProtocol,omitnil,omitempty" name:"ModelProtocol"`

	// <p>模型的厂商</p>
	ModelProvider *string `json:"ModelProvider,omitnil,omitempty" name:"ModelProvider"`

	// <p>自定义ApiBase，当ProviderKey传入且AccessType且PrivateCustom/PublicCustom时必填</p>
	ApiBase *string `json:"ApiBase,omitnil,omitempty" name:"ApiBase"`

	// <p>请求携带的Host头部，当AccessType为PrivateCustom时生效</p>
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`

	// <p>BYOK实例ID，当AccessType为PrivateCustom时生效，ProviderKey传入时必填</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>是否校验服务提供商的SSL证书</p><p>PublicBYOK时为True且禁止传入；若传入VerifySSL，则优先同步入参逻辑；若传入了ServiceProviderId则同步已创建的Byok实例该Model的逻辑；否则PublicCustom模式下为True，PrivateCustom模式下为False。</p>
	VerifySSL *bool `json:"VerifySSL,omitnil,omitempty" name:"VerifySSL"`
}

type TestModelInputModalitiesRequest struct {
	*tchttp.BaseRequest
	
	// <p>待探测的模型（原始模型名称）</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>待探测的API Key（明文）</p>
	ProviderKey *string `json:"ProviderKey,omitnil,omitempty" name:"ProviderKey"`

	// <p>已创建的BYOK API Key ID（与ProviderKey二选一传入）</p>
	ProviderKeyId *string `json:"ProviderKeyId,omitnil,omitempty" name:"ProviderKeyId"`

	// <p>BYOK类型，当ProviderKey传入时必填</p>
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// <p>模型厂商协议，当ProviderKey传入时必填</p>
	ModelProtocol *string `json:"ModelProtocol,omitnil,omitempty" name:"ModelProtocol"`

	// <p>模型的厂商</p>
	ModelProvider *string `json:"ModelProvider,omitnil,omitempty" name:"ModelProvider"`

	// <p>自定义ApiBase，当ProviderKey传入且AccessType且PrivateCustom/PublicCustom时必填</p>
	ApiBase *string `json:"ApiBase,omitnil,omitempty" name:"ApiBase"`

	// <p>请求携带的Host头部，当AccessType为PrivateCustom时生效</p>
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`

	// <p>BYOK实例ID，当AccessType为PrivateCustom时生效，ProviderKey传入时必填</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>是否校验服务提供商的SSL证书</p><p>PublicBYOK时为True且禁止传入；若传入VerifySSL，则优先同步入参逻辑；若传入了ServiceProviderId则同步已创建的Byok实例该Model的逻辑；否则PublicCustom模式下为True，PrivateCustom模式下为False。</p>
	VerifySSL *bool `json:"VerifySSL,omitnil,omitempty" name:"VerifySSL"`
}

func (r *TestModelInputModalitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TestModelInputModalitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Model")
	delete(f, "ProviderKey")
	delete(f, "ProviderKeyId")
	delete(f, "AccessType")
	delete(f, "ModelProtocol")
	delete(f, "ModelProvider")
	delete(f, "ApiBase")
	delete(f, "HostHeader")
	delete(f, "ServiceProviderId")
	delete(f, "VerifySSL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TestModelInputModalitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TestModelInputModalitiesResponseParams struct {
	// <p>探测的模型</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>该模型确认支持的输入模态列表</p><p>枚举值：</p><ul><li>text： 支持文本输入</li><li>image： 支持图像输入</li><li>file： 支持文件输入（当前仅支持pdf）</li></ul><p>收到上游大模型对于输入模态的响应即为“确认支持”</p>
	SupportedModalities []*string `json:"SupportedModalities,omitnil,omitempty" name:"SupportedModalities"`

	// <p>每个待探测模态的详细请求结果</p>
	ProbeDetails []*ModalityProbeDetail `json:"ProbeDetails,omitnil,omitempty" name:"ProbeDetails"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TestModelInputModalitiesResponse struct {
	*tchttp.BaseResponse
	Response *TestModelInputModalitiesResponseParams `json:"Response"`
}

func (r *TestModelInputModalitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TestModelInputModalitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TestServiceProviderConnectionRequestParams struct {
	// <p>需要探测的模型列表</p><p>入参限制：上限为20个模型</p>
	Models []*string `json:"Models,omitnil,omitempty" name:"Models"`

	// <p>需要探测的Key</p>
	ProviderKey *string `json:"ProviderKey,omitnil,omitempty" name:"ProviderKey"`

	// <p>需要探测的KeyId，和ProviderKey二者传一个即可</p>
	ProviderKeyId *string `json:"ProviderKeyId,omitnil,omitempty" name:"ProviderKeyId"`

	// <p>BYOK类型，当ProviderKey存在时必传</p>
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// <p>模型的厂商</p>
	ModelProvider *string `json:"ModelProvider,omitnil,omitempty" name:"ModelProvider"`

	// <p>模型厂商协议，当ProviderKey存在时必传</p>
	ModelProtocol *string `json:"ModelProtocol,omitnil,omitempty" name:"ModelProtocol"`

	// <p>BYOK类型，当AccessType为PublicCustom时生效</p>
	ApiBase *string `json:"ApiBase,omitnil,omitempty" name:"ApiBase"`

	// <p>请求携带的Host头部，当AccessType为PrivateCustom时生效</p>
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`

	// <p>BYOK的ID，当AccessType为PrivateCustom时生效</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>是否校验服务提供商的SSL证书</p><p>默认值：AccessType取值为：</p><ul><li>PublicBYOK时，该参数无效；</li><li>PublicCustom时，该参数默认为true；</li><li>PrivateCustom时，该参数默认为false；</li></ul>
	VerifySSL *bool `json:"VerifySSL,omitnil,omitempty" name:"VerifySSL"`
}

type TestServiceProviderConnectionRequest struct {
	*tchttp.BaseRequest
	
	// <p>需要探测的模型列表</p><p>入参限制：上限为20个模型</p>
	Models []*string `json:"Models,omitnil,omitempty" name:"Models"`

	// <p>需要探测的Key</p>
	ProviderKey *string `json:"ProviderKey,omitnil,omitempty" name:"ProviderKey"`

	// <p>需要探测的KeyId，和ProviderKey二者传一个即可</p>
	ProviderKeyId *string `json:"ProviderKeyId,omitnil,omitempty" name:"ProviderKeyId"`

	// <p>BYOK类型，当ProviderKey存在时必传</p>
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// <p>模型的厂商</p>
	ModelProvider *string `json:"ModelProvider,omitnil,omitempty" name:"ModelProvider"`

	// <p>模型厂商协议，当ProviderKey存在时必传</p>
	ModelProtocol *string `json:"ModelProtocol,omitnil,omitempty" name:"ModelProtocol"`

	// <p>BYOK类型，当AccessType为PublicCustom时生效</p>
	ApiBase *string `json:"ApiBase,omitnil,omitempty" name:"ApiBase"`

	// <p>请求携带的Host头部，当AccessType为PrivateCustom时生效</p>
	HostHeader *string `json:"HostHeader,omitnil,omitempty" name:"HostHeader"`

	// <p>BYOK的ID，当AccessType为PrivateCustom时生效</p>
	ServiceProviderId *string `json:"ServiceProviderId,omitnil,omitempty" name:"ServiceProviderId"`

	// <p>是否校验服务提供商的SSL证书</p><p>默认值：AccessType取值为：</p><ul><li>PublicBYOK时，该参数无效；</li><li>PublicCustom时，该参数默认为true；</li><li>PrivateCustom时，该参数默认为false；</li></ul>
	VerifySSL *bool `json:"VerifySSL,omitnil,omitempty" name:"VerifySSL"`
}

func (r *TestServiceProviderConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TestServiceProviderConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Models")
	delete(f, "ProviderKey")
	delete(f, "ProviderKeyId")
	delete(f, "AccessType")
	delete(f, "ModelProvider")
	delete(f, "ModelProtocol")
	delete(f, "ApiBase")
	delete(f, "HostHeader")
	delete(f, "ServiceProviderId")
	delete(f, "VerifySSL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TestServiceProviderConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TestServiceProviderConnectionResponseParams struct {
	// <p>探测结果</p>
	Results []*ModelTestResult `json:"Results,omitnil,omitempty" name:"Results"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TestServiceProviderConnectionResponse struct {
	*tchttp.BaseResponse
	Response *TestServiceProviderConnectionResponseParams `json:"Response"`
}

func (r *TestServiceProviderConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TestServiceProviderConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TierItem struct {
	// <p>该分层下的模型显示名称列表。</p><p>至少包含一个模型，模型名称必须是已关联到该模型路由实例的模型。同一分层内不允许重复模型名称。</p>
	Models []*string `json:"Models,omitnil,omitempty" name:"Models"`

	// <p>Tier 标识。<br>枚举值：</p><ul><li>复杂度分层（4 个固定值，需全部包含）：SIMPLE、MEDIUM、COMPLEX、REASONING</li><li>default：默认</li><li>general_chat：通用对话</li><li>transformation_rewrite：转换与改写</li><li>knowledge_qa：知识问答</li><li>summarization：摘要</li><li>extraction_structuring：抽取与结构化输出</li><li>content_generation：内容生成</li><li>coding_technical：编码与技术</li><li>data_math_analysis：数据、数学与分析</li><li>reasoning_planning：推理与规划</li><li>tool_agentic_workflow：工具与智能体工作流</li></ul>
	TierName *string `json:"TierName,omitnil,omitempty" name:"TierName"`
}

type TypeInfo struct {
	// 运营商类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 规格可用性
	SpecAvailabilitySet []*SpecAvailability `json:"SpecAvailabilitySet,omitnil,omitempty" name:"SpecAvailabilitySet"`
}

type UserGroupInfo struct {
	// <p>用户组ID。「未分组」虚拟分组固定为 ugrp-ungrouped。</p>
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// <p>用户组名称。「未分组」虚拟分组固定为 ungrouped。</p>
	UserGroupName *string `json:"UserGroupName,omitnil,omitempty" name:"UserGroupName"`

	// <p>所属模型路由实例ID。</p>
	ModelRouterId *string `json:"ModelRouterId,omitnil,omitempty" name:"ModelRouterId"`

	// <p>用户组状态。</p><p>枚举值：</p><ul><li>Creating：创建中</li><li>Active：正常</li><li>Configuring：配置中</li><li>Deleting：删除中</li></ul><p>「未分组」虚拟分组（ugrp-ungrouped）恒为 Active。</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>用户组真实模型白名单。「未分组」虚拟分组为空数组。</p>
	Models []*string `json:"Models,omitnil,omitempty" name:"Models"`

	// <p>用户组意图路由白名单（ir-xxx）。「未分组」虚拟分组为空数组。</p>
	IntentRouters []*string `json:"IntentRouters,omitnil,omitempty" name:"IntentRouters"`

	// <p>关联的Budget ID。</p><p>未关联时为空；「未分组」虚拟分组恒为空。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BudgetId *string `json:"BudgetId,omitnil,omitempty" name:"BudgetId"`

	// <p>关联的Budget名称。</p><p>未关联时为空；「未分组」虚拟分组恒为空。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	BudgetName *string `json:"BudgetName,omitnil,omitempty" name:"BudgetName"`

	// <p>关联的提示词 ID。</p><p>未关联时为空；「未分组」虚拟分组恒为空。</p><p>参数格式：pmt-1a2b3c4d</p>
	PromptId *string `json:"PromptId,omitnil,omitempty" name:"PromptId"`

	// <p>关联的提示词版本。</p><p>未关联时为空；「未分组」虚拟分组恒为空。</p>
	PromptVersion *string `json:"PromptVersion,omitnil,omitempty" name:"PromptVersion"`

	// <p>关联的提示词名称。</p><p>未关联时为空；「未分组」虚拟分组恒为空。</p>
	PromptName *string `json:"PromptName,omitnil,omitempty" name:"PromptName"`

	// <p>用户组多刷新周期 Credit 使用情况。</p><p>无多周期预算时为空数组。</p>
	CreditUsageSet []*CreditUsage `json:"CreditUsageSet,omitnil,omitempty" name:"CreditUsageSet"`

	// <p>用户组当前包含的 Key 数量。「未分组」虚拟分组（ugrp-ungrouped）返回该模型路由实例下未归属任何用户组的 Key 数量。</p>
	KeyCount *int64 `json:"KeyCount,omitnil,omitempty" name:"KeyCount"`

	// <p>标签列表。「未分组」虚拟分组为空数组。</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>创建时间。「未分组」虚拟分组不返回此字段。</p>
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// <p>修改时间。「未分组」虚拟分组不返回此字段。</p>
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`
}

type ZoneInfo struct {
	// 可用区数值形式的唯一ID，如：100001
	ZoneId *uint64 `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 可用区字符串形式的唯一ID，如：ap-guangzhou-1
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 可用区名称，如：广州一区
	ZoneName *string `json:"ZoneName,omitnil,omitempty" name:"ZoneName"`

	// 可用区所属地域，如：ap-guangzhou
	ZoneRegion *string `json:"ZoneRegion,omitnil,omitempty" name:"ZoneRegion"`

	// 可用区是否是LocalZone可用区，如：false
	LocalZone *bool `json:"LocalZone,omitnil,omitempty" name:"LocalZone"`

	// 可用区是否是EdgeZone可用区，如：false
	EdgeZone *bool `json:"EdgeZone,omitnil,omitempty" name:"EdgeZone"`
}

type ZoneResource struct {
	// 主可用区，如"ap-guangzhou-1"。
	MasterZone *string `json:"MasterZone,omitnil,omitempty" name:"MasterZone"`

	// 资源列表。
	ResourceSet []*Resource `json:"ResourceSet,omitnil,omitempty" name:"ResourceSet"`

	// 备可用区，如"ap-guangzhou-2"，单可用区时，备可用区为null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaveZone *string `json:"SlaveZone,omitnil,omitempty" name:"SlaveZone"`

	// IP版本，如IPv4，IPv6，IPv6_Nat。
	IPVersion *string `json:"IPVersion,omitnil,omitempty" name:"IPVersion"`

	// 可用区所属地域，如：ap-guangzhou
	ZoneRegion *string `json:"ZoneRegion,omitnil,omitempty" name:"ZoneRegion"`

	// 可用区是否是LocalZone可用区，如：false
	LocalZone *bool `json:"LocalZone,omitnil,omitempty" name:"LocalZone"`

	// 可用区资源的类型，SHARED表示共享资源，EXCLUSIVE表示独占资源。
	ZoneResourceType *string `json:"ZoneResourceType,omitnil,omitempty" name:"ZoneResourceType"`

	// 可用区是否是EdgeZone可用区，如：false
	EdgeZone *bool `json:"EdgeZone,omitnil,omitempty" name:"EdgeZone"`

	// 网络出口
	Egress *string `json:"Egress,omitnil,omitempty" name:"Egress"`
}