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

package v20180608

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AIModel struct {
	// 模型名
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 是否开启MCP
	EnableMCP *bool `json:"EnableMCP,omitnil,omitempty" name:"EnableMCP"`

	// 标签
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type AIModelGroup struct {
	// <p>模型分组</p><p>枚举值：</p><ul><li>hunyuan-exp： 内置 hunyuan 分组，Models 中包含混元生文模型</li><li>hunyuan-image： 内置 hunyuan 分组，Models 中包含混元生图模型</li><li>deepseek： 内置 deepseek 分组，Models 中包含Deepseek生文模型</li><li>cloudbase： 内置 cloudbase 分组，Models 中包含云开发提供的模型，支持的所有模型可从 DescribeManagedAIModelList 获取</li><li>custom-xxxx： 自定义模型分组，Models 中包含用户自行配置的模型</li></ul>
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// <p>模型列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Models []*AIModel `json:"Models,omitnil,omitempty" name:"Models"`

	// <p>模型类型</p><p>枚举值：</p><ul><li>builtin： 内置模型分组类别</li><li>custom： 用户自定义模型分组类别</li></ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>原始模型类型</p><p>枚举值：</p><ul><li>builtin： 内置模型类型</li><li>custom： 用户自定义模型类型</li></ul>
	OriginType *string `json:"OriginType,omitnil,omitempty" name:"OriginType"`

	// <p>备注</p>
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// <p>模型地址</p>
	BaseUrl *string `json:"BaseUrl,omitnil,omitempty" name:"BaseUrl"`

	// <p>模型状态, 1: 开启, 2: 关闭</p>
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>模型密钥</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Secret *AIModelSecret `json:"Secret,omitnil,omitempty" name:"Secret"`

	// <p>创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type AIModelSecret struct {
	// 密钥来源
	SecretSource *string `json:"SecretSource,omitnil,omitempty" name:"SecretSource"`

	// 密钥ID, 和SecretKey一一对应
	SecretId *string `json:"SecretId,omitnil,omitempty" name:"SecretId"`

	// 密钥Key, 和SecretId一一对应
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// ApiKey,SecretKey和ApiKey二选一
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`
}

// Predefined struct for user
type AddProviderRequestParams struct {
	// 云开发环境 ID，用于唯一标识当前操作所属的云开发环境。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 身份源的显示名称，支持国际化多语言配置。用户在登录页面看到的身份源名称将使用该字段，建议根据实际业务场景填写易于识别的名称，例如：企业微信、GitHub 等。
	Name *LocalizedMessage `json:"Name,omitnil,omitempty" name:"Name"`

	// 身份源协议类型，决定该身份源使用何种认证协议与第三方平台对接。可选值：
	// OAUTH：标准 OAuth 2.0 协议
	// OIDC：OpenID Connect 协议
	// SAML：SAML 2.0 协议
	// WX_MICRO_APP：微信小程序登录
	// WX_QRCODE_MICRO_APP：微信小程序扫码登录
	// WX_CLOUDBASE_MICRO_APP：云开发托管小程序登录
	// WX_MP：微信公众号网页授权登录
	// WX_OPEN：微信开放平台扫码登录
	// WX_WORK_INTERNAL：企业微信自建应用登录
	// WX_WORK_AGENT：企业微信代开发应用登录
	// WX_WORK_THIRD_PARTY：企业微信第三方应用登录
	// WX_WORK_THIRD_PARTY_ASSOCIATION：企业微信第三方应用关联登录
	// CUSTOM：自定义登录
	// EMAIL：邮箱登录
	ProviderType *string `json:"ProviderType,omitnil,omitempty" name:"ProviderType"`

	// 身份源的唯一标识符，用于在系统内区分不同的身份源。格式要求：2~32 位，仅支持小写英文字母和数字，不可包含空格或特殊字符。若不填写，系统将自动生成。例如：github、google。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 身份源图标的访问地址，将展示在登录页的身份源按钮上。建议使用 64×64 像素的 SVG 格式图片以保证清晰度，支持 HTTP/HTTPS 公网可访问的图片链接。
	Picture *string `json:"Picture,omitnil,omitempty" name:"Picture"`

	// 身份源对应的官方主页地址。该信息将在用户查看自己的第三方账号绑定列表时展示，帮助用户识别已绑定的身份源来源。例如 GitHub 身份源可填写：https://github.com。
	Homepage *string `json:"Homepage,omitnil,omitempty" name:"Homepage"`

	// 身份认证源协议连接配置，包含与第三方平台对接所需的核心参数，如 ClientId、ClientSecret、授权端点、Token 端点、回调地址、Scope、SAML Metadata、请求和响应参数映射等。不同 ProviderType 对应不同的配置项要求。
	Config *ProviderConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 是否开启透传登录模式。可选值：TRUE（开启）、FALSE（关闭）、UNSPECIFIED（默认为 FALSE，企业微信代开发应用 WX_WORK_AGENT 类型默认为 TRUE）。开启后，平台不会持久化存储用户数据，仅将第三方身份源返回的用户信息透传给业务方，适用于不希望平台留存用户数据的场景。注意：开启透传模式时，ReuseUserId 将自动设为 TRUE，AutoSignUpWithProviderUser 将自动设为 FALSE。
	TransparentMode *string `json:"TransparentMode,omitnil,omitempty" name:"TransparentMode"`

	// 身份源的详细描述信息，支持国际化多语言配置。可用于向用户说明该身份源的用途或使用场景，例如：谷歌授权登录。
	Description *LocalizedMessage `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否直接复用第三方身份源的用户 ID 作为平台的用户 ID。可选值：TRUE（直接复用，适用于已有用户体系迁移场景）、FALSE（不复用，由平台生成独立用户 ID）、UNSPECIFIED（默认为 FALSE，但当 TransparentMode 为 TRUE 时自动设为 TRUE）。注意：开启后需确保第三方用户 ID 的唯一性，避免 ID 冲突。
	ReuseUserId *string `json:"ReuseUserId,omitnil,omitempty" name:"ReuseUserId"`

	// 身份源的启用状态。可选值：TRUE（启用，用户可通过该身份源登录）、FALSE（禁用，登录入口将被隐藏，已有绑定关系不受影响）、UNSPECIFIED（默认为 TRUE）。
	On *string `json:"On,omitnil,omitempty" name:"On"`

	// 是否开启邮箱自动关联登录。可选值：TRUE（开启）、FALSE（关闭）、UNSPECIFIED（默认为 FALSE）。开启后，若第三方身份源返回的邮箱与系统中已有用户的邮箱一致，则自动将该第三方账号与已有用户关联并完成登录，无需用户手动绑定。
	AutoSignInWhenEmailMatch *string `json:"AutoSignInWhenEmailMatch,omitnil,omitempty" name:"AutoSignInWhenEmailMatch"`

	// 是否开启手机号自动关联登录。可选值：TRUE（开启）、FALSE（关闭）、UNSPECIFIED（默认行为等同 TRUE）。开启后，若第三方身份源返回的手机号与系统中已有用户的手机号一致，则自动将该第三方账号与已有用户关联并完成登录，无需用户手动绑定。注意：该字段默认行为（UNSPECIFIED）与 AutoSignInWhenEmailMatch 不同，手机号匹配在未显式关闭时默认生效。
	AutoSignInWhenPhoneNumberMatch *string `json:"AutoSignInWhenPhoneNumberMatch,omitnil,omitempty" name:"AutoSignInWhenPhoneNumberMatch"`
}

type AddProviderRequest struct {
	*tchttp.BaseRequest
	
	// 云开发环境 ID，用于唯一标识当前操作所属的云开发环境。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 身份源的显示名称，支持国际化多语言配置。用户在登录页面看到的身份源名称将使用该字段，建议根据实际业务场景填写易于识别的名称，例如：企业微信、GitHub 等。
	Name *LocalizedMessage `json:"Name,omitnil,omitempty" name:"Name"`

	// 身份源协议类型，决定该身份源使用何种认证协议与第三方平台对接。可选值：
	// OAUTH：标准 OAuth 2.0 协议
	// OIDC：OpenID Connect 协议
	// SAML：SAML 2.0 协议
	// WX_MICRO_APP：微信小程序登录
	// WX_QRCODE_MICRO_APP：微信小程序扫码登录
	// WX_CLOUDBASE_MICRO_APP：云开发托管小程序登录
	// WX_MP：微信公众号网页授权登录
	// WX_OPEN：微信开放平台扫码登录
	// WX_WORK_INTERNAL：企业微信自建应用登录
	// WX_WORK_AGENT：企业微信代开发应用登录
	// WX_WORK_THIRD_PARTY：企业微信第三方应用登录
	// WX_WORK_THIRD_PARTY_ASSOCIATION：企业微信第三方应用关联登录
	// CUSTOM：自定义登录
	// EMAIL：邮箱登录
	ProviderType *string `json:"ProviderType,omitnil,omitempty" name:"ProviderType"`

	// 身份源的唯一标识符，用于在系统内区分不同的身份源。格式要求：2~32 位，仅支持小写英文字母和数字，不可包含空格或特殊字符。若不填写，系统将自动生成。例如：github、google。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 身份源图标的访问地址，将展示在登录页的身份源按钮上。建议使用 64×64 像素的 SVG 格式图片以保证清晰度，支持 HTTP/HTTPS 公网可访问的图片链接。
	Picture *string `json:"Picture,omitnil,omitempty" name:"Picture"`

	// 身份源对应的官方主页地址。该信息将在用户查看自己的第三方账号绑定列表时展示，帮助用户识别已绑定的身份源来源。例如 GitHub 身份源可填写：https://github.com。
	Homepage *string `json:"Homepage,omitnil,omitempty" name:"Homepage"`

	// 身份认证源协议连接配置，包含与第三方平台对接所需的核心参数，如 ClientId、ClientSecret、授权端点、Token 端点、回调地址、Scope、SAML Metadata、请求和响应参数映射等。不同 ProviderType 对应不同的配置项要求。
	Config *ProviderConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 是否开启透传登录模式。可选值：TRUE（开启）、FALSE（关闭）、UNSPECIFIED（默认为 FALSE，企业微信代开发应用 WX_WORK_AGENT 类型默认为 TRUE）。开启后，平台不会持久化存储用户数据，仅将第三方身份源返回的用户信息透传给业务方，适用于不希望平台留存用户数据的场景。注意：开启透传模式时，ReuseUserId 将自动设为 TRUE，AutoSignUpWithProviderUser 将自动设为 FALSE。
	TransparentMode *string `json:"TransparentMode,omitnil,omitempty" name:"TransparentMode"`

	// 身份源的详细描述信息，支持国际化多语言配置。可用于向用户说明该身份源的用途或使用场景，例如：谷歌授权登录。
	Description *LocalizedMessage `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否直接复用第三方身份源的用户 ID 作为平台的用户 ID。可选值：TRUE（直接复用，适用于已有用户体系迁移场景）、FALSE（不复用，由平台生成独立用户 ID）、UNSPECIFIED（默认为 FALSE，但当 TransparentMode 为 TRUE 时自动设为 TRUE）。注意：开启后需确保第三方用户 ID 的唯一性，避免 ID 冲突。
	ReuseUserId *string `json:"ReuseUserId,omitnil,omitempty" name:"ReuseUserId"`

	// 身份源的启用状态。可选值：TRUE（启用，用户可通过该身份源登录）、FALSE（禁用，登录入口将被隐藏，已有绑定关系不受影响）、UNSPECIFIED（默认为 TRUE）。
	On *string `json:"On,omitnil,omitempty" name:"On"`

	// 是否开启邮箱自动关联登录。可选值：TRUE（开启）、FALSE（关闭）、UNSPECIFIED（默认为 FALSE）。开启后，若第三方身份源返回的邮箱与系统中已有用户的邮箱一致，则自动将该第三方账号与已有用户关联并完成登录，无需用户手动绑定。
	AutoSignInWhenEmailMatch *string `json:"AutoSignInWhenEmailMatch,omitnil,omitempty" name:"AutoSignInWhenEmailMatch"`

	// 是否开启手机号自动关联登录。可选值：TRUE（开启）、FALSE（关闭）、UNSPECIFIED（默认行为等同 TRUE）。开启后，若第三方身份源返回的手机号与系统中已有用户的手机号一致，则自动将该第三方账号与已有用户关联并完成登录，无需用户手动绑定。注意：该字段默认行为（UNSPECIFIED）与 AutoSignInWhenEmailMatch 不同，手机号匹配在未显式关闭时默认生效。
	AutoSignInWhenPhoneNumberMatch *string `json:"AutoSignInWhenPhoneNumberMatch,omitnil,omitempty" name:"AutoSignInWhenPhoneNumberMatch"`
}

func (r *AddProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddProviderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Name")
	delete(f, "ProviderType")
	delete(f, "Id")
	delete(f, "Picture")
	delete(f, "Homepage")
	delete(f, "Config")
	delete(f, "TransparentMode")
	delete(f, "Description")
	delete(f, "ReuseUserId")
	delete(f, "On")
	delete(f, "AutoSignInWhenEmailMatch")
	delete(f, "AutoSignInWhenPhoneNumberMatch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddProviderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddProviderResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddProviderResponse struct {
	*tchttp.BaseResponse
	Response *AddProviderResponseParams `json:"Response"`
}

func (r *AddProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AllocateEnvRequestParams struct {
	// <p>分配请求ID，会按这个值做幂等</p><p>入参限制：长度不超过64</p>
	AllocateId *string `json:"AllocateId,omitnil,omitempty" name:"AllocateId"`

	// <p>客户平台的应用标识，如果没有则不传</p>
	ExternalAppId *string `json:"ExternalAppId,omitnil,omitempty" name:"ExternalAppId"`
}

type AllocateEnvRequest struct {
	*tchttp.BaseRequest
	
	// <p>分配请求ID，会按这个值做幂等</p><p>入参限制：长度不超过64</p>
	AllocateId *string `json:"AllocateId,omitnil,omitempty" name:"AllocateId"`

	// <p>客户平台的应用标识，如果没有则不传</p>
	ExternalAppId *string `json:"ExternalAppId,omitnil,omitempty" name:"ExternalAppId"`
}

func (r *AllocateEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AllocateId")
	delete(f, "ExternalAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AllocateEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AllocateEnvResponseParams struct {
	// <p>环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>回显    客户平台的应用标识，如果没有则不传</p>
	ExternalAppId *string `json:"ExternalAppId,omitnil,omitempty" name:"ExternalAppId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AllocateEnvResponse struct {
	*tchttp.BaseResponse
	Response *AllocateEnvResponseParams `json:"Response"`
}

func (r *AllocateEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AllocateEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApiKeyToken struct {
	// API Key 的唯一标识符，由系统基于 UUID 自动生成的 Base64 URL 编码字符串。后续对该 API Key 进行删除、修改名称或精确查询操作时，均需使用该值作为定位参数
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// API Key 的名称，即创建时传入的 KeyName 参数值。对于 publish_key 类型，该值固定为 publish_key
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// API Key 的令牌值（JWT 格式），用于服务端接口调用时的身份认证。出于安全考虑，仅在创建时返回一次完整明文；后续通过列表查询接口获取时，api_key 类型将进行脱敏处理；publish_key 类型始终返回完整明文。请在创建后妥善保存
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`

	// API Key 的过期时间，格式遵循 ISO 8601 标准。对于 api_key 类型：若创建时未指定有效期（ExpireIn），则该字段不返回，表示永不过期；若指定了有效期，则返回具体的过期时间。对于 publish_key 类型：始终返回，固定为约 2099 年
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireAt *string `json:"ExpireAt,omitnil,omitempty" name:"ExpireAt"`

	// API Key 的创建时间，格式遵循 ISO 8601 标准。对于 api_key 类型：为该 Key 实际创建时的时间。对于 publish_key 类型：若环境下已存在 publish_key 记录，则返回首次创建的时间而非本次调用时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`
}

// Predefined struct for user
type AssumeRoleForAllocatedEnvRequestParams struct {
	// <p>环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type AssumeRoleForAllocatedEnvRequest struct {
	*tchttp.BaseRequest
	
	// <p>环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *AssumeRoleForAllocatedEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssumeRoleForAllocatedEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssumeRoleForAllocatedEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssumeRoleForAllocatedEnvResponseParams struct {
	// <p>SecretId</p>
	SecretId *string `json:"SecretId,omitnil,omitempty" name:"SecretId"`

	// <p>SecretKey</p>
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// <p>Token值</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// <p>过期时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// <p>是否从缓存中加载。标明该值是否实时从sts服务获取，还是从缓存中获取。调用方可不关心</p>
	IsCache *bool `json:"IsCache,omitnil,omitempty" name:"IsCache"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssumeRoleForAllocatedEnvResponse struct {
	*tchttp.BaseResponse
	Response *AssumeRoleForAllocatedEnvResponseParams `json:"Response"`
}

func (r *AssumeRoleForAllocatedEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssumeRoleForAllocatedEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuthDomain struct {
	// 域名ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名类型。包含以下取值：
	// <li>SYSTEM</li>
	// <li>USER</li>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 状态。包含以下取值：
	// <li>ENABLE</li>
	// <li>DISABLE</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type BaasPackageInfo struct {
	// DAU产品套餐ID
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// DAU套餐中文名称
	PackageTitle *string `json:"PackageTitle,omitnil,omitempty" name:"PackageTitle"`

	// 套餐分组
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 套餐分组中文名
	GroupTitle *string `json:"GroupTitle,omitnil,omitempty" name:"GroupTitle"`

	// json格式化计费标签，例如：
	// {"pid":2, "cids":{"create": 2, "renew": 2, "modify": 2}, "productCode":"p_tcb_mp", "subProductCode":"sp_tcb_mp_cloudbase_dau"}
	BillTags *string `json:"BillTags,omitnil,omitempty" name:"BillTags"`

	// json格式化用户资源限制，例如：
	// {"Qps":1000,"InvokeNum":{"TimeUnit":"m", "Unit":"万次", "MaxSize": 100},"Capacity":{"TimeUnit":"m", "Unit":"GB", "MaxSize": 100}, "Cdn":{"Flux":{"TimeUnit":"m", "Unit":"GB", "MaxSize": 100}, "BackFlux":{"TimeUnit":"m", "Unit":"GB", "MaxSize": 100}},"Scf":{"Concurrency":1000,"OutFlux":{"TimeUnit":"m", "Unit":"GB", "MaxSize": 100},"MemoryUse":{"TimeUnit":"m", "Unit":"WGBS", "MaxSize": 100000}}}
	ResourceLimit *string `json:"ResourceLimit,omitnil,omitempty" name:"ResourceLimit"`

	// json格式化高级限制，例如：
	// {"CMSEnable":false,"ProvisionedConcurrencyMem":512000, "PictureProcessing":false, "SecurityAudit":false, "RealTimePush":false, "TemplateMessageBatchPush":false, "Payment":false}
	AdvanceLimit *string `json:"AdvanceLimit,omitnil,omitempty" name:"AdvanceLimit"`

	// 套餐描述
	PackageDescription *string `json:"PackageDescription,omitnil,omitempty" name:"PackageDescription"`

	// 是否对外展示
	IsExternal *bool `json:"IsExternal,omitnil,omitempty" name:"IsExternal"`
}

type BanConfig struct {
	// ip白名单，支持ipv4、ipv6，支持CIDR
	IpWhiteList []*string `json:"IpWhiteList,omitnil,omitempty" name:"IpWhiteList"`

	// ip黑名单，支持ipv4、ipv6，支持CIDR
	IpBlackList []*string `json:"IpBlackList,omitnil,omitempty" name:"IpBlackList"`

	// 地域白名单（国家英文名）
	CountryWhiteList []*string `json:"CountryWhiteList,omitnil,omitempty" name:"CountryWhiteList"`

	// 地域黑名单（国家英文名）
	CountryBlackList []*string `json:"CountryBlackList,omitnil,omitempty" name:"CountryBlackList"`
}

// Predefined struct for user
type BindStorageSourceRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 存储源
	StorageConfig *ExternalStorage `json:"StorageConfig,omitnil,omitempty" name:"StorageConfig"`
}

type BindStorageSourceRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 存储源
	StorageConfig *ExternalStorage `json:"StorageConfig,omitnil,omitempty" name:"StorageConfig"`
}

func (r *BindStorageSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindStorageSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "StorageConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindStorageSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindStorageSourceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindStorageSourceResponse struct {
	*tchttp.BaseResponse
	Response *BindStorageSourceResponseParams `json:"Response"`
}

func (r *BindStorageSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindStorageSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckTcbServiceRequestParams struct {

}

type CheckTcbServiceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *CheckTcbServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckTcbServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckTcbServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckTcbServiceResponseParams struct {
	// true表示已开通
	Initialized *bool `json:"Initialized,omitnil,omitempty" name:"Initialized"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckTcbServiceResponse struct {
	*tchttp.BaseResponse
	Response *CheckTcbServiceResponseParams `json:"Response"`
}

func (r *CheckTcbServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckTcbServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClsInfo struct {
	// cls所属地域
	ClsRegion *string `json:"ClsRegion,omitnil,omitempty" name:"ClsRegion"`

	// cls日志集ID
	ClsLogsetId *string `json:"ClsLogsetId,omitnil,omitempty" name:"ClsLogsetId"`

	// cls日志主题ID
	ClsTopicId *string `json:"ClsTopicId,omitnil,omitempty" name:"ClsTopicId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type ClusterDetail struct {
	// 是否开启公网访问
	IsOpenPubNetAccess *bool `json:"IsOpenPubNetAccess,omitnil,omitempty" name:"IsOpenPubNetAccess"`

	// 最大算力
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxCpu *float64 `json:"MaxCpu,omitnil,omitempty" name:"MaxCpu"`

	// 最小算力
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinCpu *float64 `json:"MinCpu,omitnil,omitempty" name:"MinCpu"`

	// TDSQL-C集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 存储用量（单位：MB）
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedStorage *int64 `json:"UsedStorage,omitnil,omitempty" name:"UsedStorage"`

	// 最大存储量（单位：GB）
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageLimit *int64 `json:"StorageLimit,omitnil,omitempty" name:"StorageLimit"`

	// 数据库类型
	DbType *string `json:"DbType,omitnil,omitempty" name:"DbType"`

	// 数据库类型
	DbVersion *string `json:"DbVersion,omitnil,omitempty" name:"DbVersion"`

	// 公网访问状态；open开启，opening开启中，closed关闭，closing关闭中
	WanStatus *string `json:"WanStatus,omitnil,omitempty" name:"WanStatus"`

	// 数据库集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterStatus *string `json:"ClusterStatus,omitnil,omitempty" name:"ClusterStatus"`

	// serverless状态
	ServerlessStatus *string `json:"ServerlessStatus,omitnil,omitempty" name:"ServerlessStatus"`
}

// Predefined struct for user
type CreateAIModelRequestParams struct {
	// <p>环境id</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>分组名</p><p>入参限制：不允许以 cloudbase 为前缀</p>
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// <p>模型服务地址</p>
	BaseUrl *string `json:"BaseUrl,omitnil,omitempty" name:"BaseUrl"`

	// <p>模型名列表</p>
	Models []*AIModel `json:"Models,omitnil,omitempty" name:"Models"`

	// <p>分组备注</p>
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// <p>模型状态,  1: 开启, 2: 关闭</p>
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>模型密钥</p>
	Secret *AIModelSecret `json:"Secret,omitnil,omitempty" name:"Secret"`
}

type CreateAIModelRequest struct {
	*tchttp.BaseRequest
	
	// <p>环境id</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>分组名</p><p>入参限制：不允许以 cloudbase 为前缀</p>
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// <p>模型服务地址</p>
	BaseUrl *string `json:"BaseUrl,omitnil,omitempty" name:"BaseUrl"`

	// <p>模型名列表</p>
	Models []*AIModel `json:"Models,omitnil,omitempty" name:"Models"`

	// <p>分组备注</p>
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// <p>模型状态,  1: 开启, 2: 关闭</p>
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>模型密钥</p>
	Secret *AIModelSecret `json:"Secret,omitnil,omitempty" name:"Secret"`
}

func (r *CreateAIModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "GroupName")
	delete(f, "BaseUrl")
	delete(f, "Models")
	delete(f, "Remark")
	delete(f, "Status")
	delete(f, "Secret")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAIModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAIModelResponseParams struct {
	// <p>创建数量</p>
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAIModelResponse struct {
	*tchttp.BaseResponse
	Response *CreateAIModelResponseParams `json:"Response"`
}

func (r *CreateAIModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiKeyRequestParams struct {
	// 环境 ID，用于标识该密钥归属的云开发环境，不同环境之间的数据相互隔离
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 密钥类型。可选值：api_key（服务端调用使用的 API 密钥，具有完整权限，请勿暴露在客户端）、publish_key（客户端使用的公开密钥，权限受限，可安全用于前端或移动端）。
	KeyType *string `json:"KeyType,omitnil,omitempty" name:"KeyType"`

	// 密钥的自定义名称，用于在管理列表中标识和区分不同的密钥，建议填写能体现用途或归属的描述性名称，例如：server-prod、mobile-test
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// 密钥的有效期，单位为秒，最短不得低于 7200 秒。超过有效期后密钥将自动失效。不设置或设置为 0 则表示永不过期，建议根据安全需求合理设置有效期
	ExpireIn *int64 `json:"ExpireIn,omitnil,omitempty" name:"ExpireIn"`
}

type CreateApiKeyRequest struct {
	*tchttp.BaseRequest
	
	// 环境 ID，用于标识该密钥归属的云开发环境，不同环境之间的数据相互隔离
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 密钥类型。可选值：api_key（服务端调用使用的 API 密钥，具有完整权限，请勿暴露在客户端）、publish_key（客户端使用的公开密钥，权限受限，可安全用于前端或移动端）。
	KeyType *string `json:"KeyType,omitnil,omitempty" name:"KeyType"`

	// 密钥的自定义名称，用于在管理列表中标识和区分不同的密钥，建议填写能体现用途或归属的描述性名称，例如：server-prod、mobile-test
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// 密钥的有效期，单位为秒，最短不得低于 7200 秒。超过有效期后密钥将自动失效。不设置或设置为 0 则表示永不过期，建议根据安全需求合理设置有效期
	ExpireIn *int64 `json:"ExpireIn,omitnil,omitempty" name:"ExpireIn"`
}

func (r *CreateApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "KeyType")
	delete(f, "KeyName")
	delete(f, "ExpireIn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApiKeyResponseParams struct {
	// API Key 的唯一标识符，由系统基于 JWT Access Token Hash 自动生成。后续对该 API Key 进行查询、修改名称或删除操作时，均需使用该值作为定位参数
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`

	// API Key 的名称，即创建时传入的 KeyName 参数值。对于 publish_key 类型，该值固定为 publish_key
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// API Key 的令牌值（JWT 格式），用于服务端接口调用时的身份认证。出于安全考虑，仅在创建时返回一次完整明文；后续通过列表查询接口获取时，api_key 类型将进行脱敏处理；publish_key 类型始终返回完整明文。请在创建后妥善保存
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`

	// API Key 的过期时间。对于 api_key 类型：若创建时未指定有效期，则该字段不返回，表示永不过期；若指定了有效期，则返回具体的过期时间。对于 publish_key 类型：始终返回，固定为 2099 年
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireAt *string `json:"ExpireAt,omitnil,omitempty" name:"ExpireAt"`

	// API Key 的创建时间。对于 api_key 类型：为实际创建该 Key 时的时间。对于 publish_key 类型：若环境下已存在 publish_key，则返回首次创建的时间而非本次调用时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateAt *string `json:"CreateAt,omitnil,omitempty" name:"CreateAt"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateApiKeyResponseParams `json:"Response"`
}

func (r *CreateApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuthDomainRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 安全域名
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`
}

type CreateAuthDomainRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 安全域名
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`
}

func (r *CreateAuthDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuthDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Domains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuthDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuthDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAuthDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuthDomainResponseParams `json:"Response"`
}

func (r *CreateAuthDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuthDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBillDealRequestParams struct {
	// 当前下单的操作类型，可取[purchase,renew,modify]三种值，分别代表新购，续费，变配。
	DealType *string `json:"DealType,omitnil,omitempty" name:"DealType"`

	// 购买的产品类型，可取[tcb-baas,tcb-promotion,tcb-package], 分别代表baas套餐、大促包、资源包
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 目标下单产品/套餐Id。
	// 对于云开发环境套餐，可通过 DescribeBaasPackageList 接口获取，对应其出参的PackageName
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 默认只下单不支付，为ture则下单并支付。
	// 如果需要下单并支付，请确保账户下有足够的余额，否则会导致下单失败。
	CreateAndPay *bool `json:"CreateAndPay,omitnil,omitempty" name:"CreateAndPay"`

	// 购买时长，与TimeUnit字段搭配使用。
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 购买时长单位,按各产品规则可选d(天),m(月),y(年),p(一次性)。
	// 对于 云开发环境的 新购和续费，目前仅支持 按月购买（即 TimeUnit=m）。
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 资源唯一标识。
	// 在云开发环境 续费和变配 场景下必传，取值为环境ID。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 来源可选[qcloud,miniapp]，默认qcloud。
	// miniapp表示微信云开发，主要适用于[小程序云开发](https://developers.weixin.qq.com/miniprogram/dev/wxcloudservice/wxcloud/billing/price.html)。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 环境别名，用于新购云开发环境时，给云开发环境起别名。
	// 仅当 新购云开发环境（DealType=purchase 并且 ProductType=tcb-baas ）时有效。
	// 
	// ### 格式要求
	// - 可选字符： 小写字母(a~z)、数字、减号(-)
	// - 不能以 减号(-) 开头或结尾
	// - 不能有连个连续的 减号(-)
	// - 长度不超过20位
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 环境id，当购买资源包和大促包时（ProductType取值为tcb-promotion 或 tcb-package）必传，表示资源包在哪个环境下生效。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 开启超限按量。
	// 开启后，当 套餐内的资源点 和 资源包 都用尽后，会自动按量计费。
	// 详见 [计费说明](https://cloud.tencent.com/document/product/876/127357)。
	EnableExcess *bool `json:"EnableExcess,omitnil,omitempty" name:"EnableExcess"`

	// 变配目标套餐id，对于云开发环境变配场景下必传。
	// 对于云开发环境套餐，可通过 DescribeBaasPackageList 接口获取，对应其出参的PackageName
	ModifyPackageId *string `json:"ModifyPackageId,omitnil,omitempty" name:"ModifyPackageId"`

	// jsonstr附加信息
	Extension *string `json:"Extension,omitnil,omitempty" name:"Extension"`

	// 是否自动选择代金券支付。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 资源类型。
	// 代表新购环境（DealType=purchase 并且 ProductType=tcb-baas ）时需要发货哪些资源。
	// 可取值：flexdb, cos, cdn, scf
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// 环境标签。
	//  代表新购环境（DealType=purchase 并且 ProductType=tcb-baas ）时需要打的标签。
	EnvTags []*Tag `json:"EnvTags,omitnil,omitempty" name:"EnvTags"`
}

type CreateBillDealRequest struct {
	*tchttp.BaseRequest
	
	// 当前下单的操作类型，可取[purchase,renew,modify]三种值，分别代表新购，续费，变配。
	DealType *string `json:"DealType,omitnil,omitempty" name:"DealType"`

	// 购买的产品类型，可取[tcb-baas,tcb-promotion,tcb-package], 分别代表baas套餐、大促包、资源包
	ProductType *string `json:"ProductType,omitnil,omitempty" name:"ProductType"`

	// 目标下单产品/套餐Id。
	// 对于云开发环境套餐，可通过 DescribeBaasPackageList 接口获取，对应其出参的PackageName
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 默认只下单不支付，为ture则下单并支付。
	// 如果需要下单并支付，请确保账户下有足够的余额，否则会导致下单失败。
	CreateAndPay *bool `json:"CreateAndPay,omitnil,omitempty" name:"CreateAndPay"`

	// 购买时长，与TimeUnit字段搭配使用。
	TimeSpan *uint64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 购买时长单位,按各产品规则可选d(天),m(月),y(年),p(一次性)。
	// 对于 云开发环境的 新购和续费，目前仅支持 按月购买（即 TimeUnit=m）。
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 资源唯一标识。
	// 在云开发环境 续费和变配 场景下必传，取值为环境ID。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 来源可选[qcloud,miniapp]，默认qcloud。
	// miniapp表示微信云开发，主要适用于[小程序云开发](https://developers.weixin.qq.com/miniprogram/dev/wxcloudservice/wxcloud/billing/price.html)。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 环境别名，用于新购云开发环境时，给云开发环境起别名。
	// 仅当 新购云开发环境（DealType=purchase 并且 ProductType=tcb-baas ）时有效。
	// 
	// ### 格式要求
	// - 可选字符： 小写字母(a~z)、数字、减号(-)
	// - 不能以 减号(-) 开头或结尾
	// - 不能有连个连续的 减号(-)
	// - 长度不超过20位
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 环境id，当购买资源包和大促包时（ProductType取值为tcb-promotion 或 tcb-package）必传，表示资源包在哪个环境下生效。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 开启超限按量。
	// 开启后，当 套餐内的资源点 和 资源包 都用尽后，会自动按量计费。
	// 详见 [计费说明](https://cloud.tencent.com/document/product/876/127357)。
	EnableExcess *bool `json:"EnableExcess,omitnil,omitempty" name:"EnableExcess"`

	// 变配目标套餐id，对于云开发环境变配场景下必传。
	// 对于云开发环境套餐，可通过 DescribeBaasPackageList 接口获取，对应其出参的PackageName
	ModifyPackageId *string `json:"ModifyPackageId,omitnil,omitempty" name:"ModifyPackageId"`

	// jsonstr附加信息
	Extension *string `json:"Extension,omitnil,omitempty" name:"Extension"`

	// 是否自动选择代金券支付。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// 资源类型。
	// 代表新购环境（DealType=purchase 并且 ProductType=tcb-baas ）时需要发货哪些资源。
	// 可取值：flexdb, cos, cdn, scf
	ResourceTypes []*string `json:"ResourceTypes,omitnil,omitempty" name:"ResourceTypes"`

	// 环境标签。
	//  代表新购环境（DealType=purchase 并且 ProductType=tcb-baas ）时需要打的标签。
	EnvTags []*Tag `json:"EnvTags,omitnil,omitempty" name:"EnvTags"`
}

func (r *CreateBillDealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBillDealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DealType")
	delete(f, "ProductType")
	delete(f, "PackageId")
	delete(f, "CreateAndPay")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "ResourceId")
	delete(f, "Source")
	delete(f, "Alias")
	delete(f, "EnvId")
	delete(f, "EnableExcess")
	delete(f, "ModifyPackageId")
	delete(f, "Extension")
	delete(f, "AutoVoucher")
	delete(f, "ResourceTypes")
	delete(f, "EnvTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateBillDealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateBillDealResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateBillDealResponse struct {
	*tchttp.BaseResponse
	Response *CreateBillDealResponseParams `json:"Response"`
}

func (r *CreateBillDealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateBillDealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomLoginKeyRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type CreateCustomLoginKeyRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *CreateCustomLoginKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomLoginKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomLoginKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomLoginKeyResponseParams struct {
	// 自定义登录的 RSA 私钥（1024 位），PEM 编码格式（PKCS#1）。调用方需使用该私钥对包含用户身份信息的 JSON 数据进行 JWS 签名，生成 JWT Token 后传入自定义登录接口完成身份认证。出于安全考虑，系统仅存储公钥，私钥仅在创建时返回一次且无法恢复，请妥善保存。创建新密钥后，该环境下原有未设置过期时间的旧密钥将被自动标记为 2 小时后过期
	PrivateKey *string `json:"PrivateKey,omitnil,omitempty" name:"PrivateKey"`

	// 密钥对的唯一标识符（UUID 格式），由系统自动生成。在自定义登录时，需将该 KeyID 拼接到 ProviderToken 参数中（格式：{KeyID}/{algorithm}/{signedJWT}），服务端通过 KeyID 查找对应的公钥以验证签名
	KeyID *string `json:"KeyID,omitnil,omitempty" name:"KeyID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCustomLoginKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomLoginKeyResponseParams `json:"Response"`
}

func (r *CreateCustomLoginKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomLoginKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnvRequestParams struct {
	// <p>环境别名。</p><h3 id=".E6.A0.BC.E5.BC.8F.E8.A6.81.E6.B1.82">格式要求</h3><ul><li>可选字符： 小写字母(a~z)、数字、减号(-)</li><li>不能以 减号(-) 开头或结尾</li><li>不能有连个连续的 减号(-)</li><li>长度不超过20位<br>示例值：cloud</li></ul>
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// <p>云开发环境套餐Id。<br>对于云开发环境套餐，可通过 <a href="https://cloud.tencent.com/document/product/876/78167">DescribeBaasPackageList</a> 接口获取，对应其出参的PackageName。</p>
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// <p>资源类型。代表新购环境时需要发货哪些资源。<br>可取值以及含义：</p><ul><li>flexdb : 表示文档型数据库</li><li>storage : 表示云存储</li><li>function : 表示云函数</li><li>postgresql：表示postgresql数据库</li></ul><p><strong>该字段不可为空</strong></p>
	Resources []*string `json:"Resources,omitnil,omitempty" name:"Resources"`

	// <p>购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24。<br>默认值为1，即1个月。</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>是否自动选择代金券支付。</p>
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>环境标签。<br>可取值通过接口 <a href="https://cloud.tencent.com/document/product/651/35316">tag:DescribeTags</a> 可获取到。<br>不传或为空则默认不打任何标签。</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>自动续费标识。取值范围：</p><ul><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费（需要手动续费，可通过接口 <a href="https://cloud.tencent.com/document/product/876/128590">RenewEnv</a> 来续费）</li></ul><p>默认取值：NOTIFY_AND_MANUAL_RENEW。<br>若该参数指定为NOTIFY_AND_AUTO_RENEW（即：自动续费），在账户余额充足的情况下，实例到期后将按月自动续费；但如果账户余额不足，将无法自动续费。请留意腾讯云短信和邮件通知。</p>
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// <p>云存储外部存储源。<br>表示该环境下不再自动分配云存储cos桶，而是由指定的bucket作为该环境的云存储介质。</p><p>仅当 Resources 中包含storage时有效。</p>
	ExternalStorage *ExternalStorage `json:"ExternalStorage,omitnil,omitempty" name:"ExternalStorage"`

	// <p>开启或关闭 超限转按量。 可取值： TRUE/FALSE （字符串类型） 非法制、不传、为空 则不变更该字段。</p>
	EnableOverrun *string `json:"EnableOverrun,omitnil,omitempty" name:"EnableOverrun"`
}

type CreateEnvRequest struct {
	*tchttp.BaseRequest
	
	// <p>环境别名。</p><h3 id=".E6.A0.BC.E5.BC.8F.E8.A6.81.E6.B1.82">格式要求</h3><ul><li>可选字符： 小写字母(a~z)、数字、减号(-)</li><li>不能以 减号(-) 开头或结尾</li><li>不能有连个连续的 减号(-)</li><li>长度不超过20位<br>示例值：cloud</li></ul>
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// <p>云开发环境套餐Id。<br>对于云开发环境套餐，可通过 <a href="https://cloud.tencent.com/document/product/876/78167">DescribeBaasPackageList</a> 接口获取，对应其出参的PackageName。</p>
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// <p>资源类型。代表新购环境时需要发货哪些资源。<br>可取值以及含义：</p><ul><li>flexdb : 表示文档型数据库</li><li>storage : 表示云存储</li><li>function : 表示云函数</li><li>postgresql：表示postgresql数据库</li></ul><p><strong>该字段不可为空</strong></p>
	Resources []*string `json:"Resources,omitnil,omitempty" name:"Resources"`

	// <p>购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24。<br>默认值为1，即1个月。</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>是否自动选择代金券支付。</p>
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`

	// <p>环境标签。<br>可取值通过接口 <a href="https://cloud.tencent.com/document/product/651/35316">tag:DescribeTags</a> 可获取到。<br>不传或为空则默认不打任何标签。</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>自动续费标识。取值范围：</p><ul><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费（需要手动续费，可通过接口 <a href="https://cloud.tencent.com/document/product/876/128590">RenewEnv</a> 来续费）</li></ul><p>默认取值：NOTIFY_AND_MANUAL_RENEW。<br>若该参数指定为NOTIFY_AND_AUTO_RENEW（即：自动续费），在账户余额充足的情况下，实例到期后将按月自动续费；但如果账户余额不足，将无法自动续费。请留意腾讯云短信和邮件通知。</p>
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// <p>云存储外部存储源。<br>表示该环境下不再自动分配云存储cos桶，而是由指定的bucket作为该环境的云存储介质。</p><p>仅当 Resources 中包含storage时有效。</p>
	ExternalStorage *ExternalStorage `json:"ExternalStorage,omitnil,omitempty" name:"ExternalStorage"`

	// <p>开启或关闭 超限转按量。 可取值： TRUE/FALSE （字符串类型） 非法制、不传、为空 则不变更该字段。</p>
	EnableOverrun *string `json:"EnableOverrun,omitnil,omitempty" name:"EnableOverrun"`
}

func (r *CreateEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Alias")
	delete(f, "PackageId")
	delete(f, "Resources")
	delete(f, "Period")
	delete(f, "AutoVoucher")
	delete(f, "Tags")
	delete(f, "RenewFlag")
	delete(f, "ExternalStorage")
	delete(f, "EnableOverrun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnvResourceRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 资源类型。代表本次开通哪些资源。 可取值以及含义： - log : 表示日志资源，当前仅支持 log（日志资源类型），后续版本可能扩展，该数组不能为空，且每个元素必须为合法的资源类型值
	Resources []*string `json:"Resources,omitnil,omitempty" name:"Resources"`
}

type CreateEnvResourceRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 资源类型。代表本次开通哪些资源。 可取值以及含义： - log : 表示日志资源，当前仅支持 log（日志资源类型），后续版本可能扩展，该数组不能为空，且每个元素必须为合法的资源类型值
	Resources []*string `json:"Resources,omitnil,omitempty" name:"Resources"`
}

func (r *CreateEnvResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Resources")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEnvResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnvResourceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEnvResourceResponse struct {
	*tchttp.BaseResponse
	Response *CreateEnvResourceResponseParams `json:"Response"`
}

func (r *CreateEnvResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnvResponseParams struct {
	// <p>自动生成的环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateEnvResponse struct {
	*tchttp.BaseResponse
	Response *CreateEnvResponseParams `json:"Response"`
}

func (r *CreateEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHTTPServiceRouteRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 域名路由信息
	Domain *HTTPServiceDomainParam `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type CreateHTTPServiceRouteRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 域名路由信息
	Domain *HTTPServiceDomainParam `json:"Domain,omitnil,omitempty" name:"Domain"`
}

func (r *CreateHTTPServiceRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHTTPServiceRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHTTPServiceRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHTTPServiceRouteResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateHTTPServiceRouteResponse struct {
	*tchttp.BaseResponse
	Response *CreateHTTPServiceRouteResponseParams `json:"Response"`
}

func (r *CreateHTTPServiceRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHTTPServiceRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHostingDomainRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`
}

type CreateHostingDomainRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`
}

func (r *CreateHostingDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHostingDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Domain")
	delete(f, "CertId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHostingDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHostingDomainResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateHostingDomainResponse struct {
	*tchttp.BaseResponse
	Response *CreateHostingDomainResponseParams `json:"Response"`
}

func (r *CreateHostingDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHostingDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateIndex struct {
	// 索引名称
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`

	// 索引结构
	MgoKeySchema *MgoKeySchema `json:"MgoKeySchema,omitnil,omitempty" name:"MgoKeySchema"`
}

// Predefined struct for user
type CreateMySQLRequestParams struct {
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// Db类型: MYSQL
	DbInstanceType *string `json:"DbInstanceType,omitnil,omitempty" name:"DbInstanceType"`

	// mysql版本
	MysqlVersion *string `json:"MysqlVersion,omitnil,omitempty" name:"MysqlVersion"`

	// vpc Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 0 区分表名大小写；1 不区分表名大小写(默认)
	LowerCaseTableNames *string `json:"LowerCaseTableNames,omitnil,omitempty" name:"LowerCaseTableNames"`
}

type CreateMySQLRequest struct {
	*tchttp.BaseRequest
	
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// Db类型: MYSQL
	DbInstanceType *string `json:"DbInstanceType,omitnil,omitempty" name:"DbInstanceType"`

	// mysql版本
	MysqlVersion *string `json:"MysqlVersion,omitnil,omitempty" name:"MysqlVersion"`

	// vpc Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 0 区分表名大小写；1 不区分表名大小写(默认)
	LowerCaseTableNames *string `json:"LowerCaseTableNames,omitnil,omitempty" name:"LowerCaseTableNames"`
}

func (r *CreateMySQLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMySQLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "DbInstanceType")
	delete(f, "MysqlVersion")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "LowerCaseTableNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMySQLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMySQLResponseParams struct {
	// 开通结果
	Data *CreateMySQLResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMySQLResponse struct {
	*tchttp.BaseResponse
	Response *CreateMySQLResponseParams `json:"Response"`
}

func (r *CreateMySQLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMySQLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMySQLResult struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type CreateStaticStoreRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 是否启用统一域名
	EnableUnion *bool `json:"EnableUnion,omitnil,omitempty" name:"EnableUnion"`

	// 外部存储源。
	ExternalStorage *ExternalStorage `json:"ExternalStorage,omitnil,omitempty" name:"ExternalStorage"`
}

type CreateStaticStoreRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 是否启用统一域名
	EnableUnion *bool `json:"EnableUnion,omitnil,omitempty" name:"EnableUnion"`

	// 外部存储源。
	ExternalStorage *ExternalStorage `json:"ExternalStorage,omitnil,omitempty" name:"ExternalStorage"`
}

func (r *CreateStaticStoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStaticStoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "EnableUnion")
	delete(f, "ExternalStorage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStaticStoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStaticStoreResponseParams struct {
	// 创建静态资源结果(succ/fail)
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStaticStoreResponse struct {
	*tchttp.BaseResponse
	Response *CreateStaticStoreResponseParams `json:"Response"`
}

func (r *CreateStaticStoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStaticStoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTableRequestParams struct {
	// 数据表名；长度不超过96个字符，可以为英文字母、数字、下划线(_)和短横线(-)的组合，且不能以下划线开头
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// FlexDB实例ID，如：tnt-nl7hjzasw
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// FlexDB数据库权限信息
	PermissionInfo *PermissionInfo `json:"PermissionInfo,omitnil,omitempty" name:"PermissionInfo"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoDB连接器配置
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

type CreateTableRequest struct {
	*tchttp.BaseRequest
	
	// 数据表名；长度不超过96个字符，可以为英文字母、数字、下划线(_)和短横线(-)的组合，且不能以下划线开头
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// FlexDB实例ID，如：tnt-nl7hjzasw
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// FlexDB数据库权限信息
	PermissionInfo *PermissionInfo `json:"PermissionInfo,omitnil,omitempty" name:"PermissionInfo"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoDB连接器配置
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

func (r *CreateTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableName")
	delete(f, "Tag")
	delete(f, "PermissionInfo")
	delete(f, "EnvId")
	delete(f, "MongoConnector")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTableResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateTableResponseParams `json:"Response"`
}

func (r *CreateTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 用户名，用户名规则：1. 长度1-64字符 2. 只能包含大小写英文字母、数字和符号 . _ - 3. 只能以字母或数字开头 4. 不能重复
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户ID，最多64字符，如不传则系统自动生成
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 用户类型：internalUser-内部用户、externalUser-外部用户，默认internalUser（内部用户）
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 密码，传入Uid时密码可不传。密码规则：1. 长度8-32字符（推荐12位以上） 2. 不能以特殊字符开头 3. 至少包含以下四项中的三项：小写字母a-z、大写字母A-Z、数字0-9、特殊字符()!@#$%^&*\|?><_-
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 用户状态：ACTIVE（激活）、BLOCKED（冻结），默认激活
	UserStatus *string `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 用户昵称，长度2-64字符
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 手机号，不能重复
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 邮箱地址，不能重复
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 头像链接，可访问的公网URL
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`

	// 用户描述，最多200字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 用户名，用户名规则：1. 长度1-64字符 2. 只能包含大小写英文字母、数字和符号 . _ - 3. 只能以字母或数字开头 4. 不能重复
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户ID，最多64字符，如不传则系统自动生成
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 用户类型：internalUser-内部用户、externalUser-外部用户，默认internalUser（内部用户）
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 密码，传入Uid时密码可不传。密码规则：1. 长度8-32字符（推荐12位以上） 2. 不能以特殊字符开头 3. 至少包含以下四项中的三项：小写字母a-z、大写字母A-Z、数字0-9、特殊字符()!@#$%^&*\|?><_-
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 用户状态：ACTIVE（激活）、BLOCKED（冻结），默认激活
	UserStatus *string `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 用户昵称，长度2-64字符
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 手机号，不能重复
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 邮箱地址，不能重复
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 头像链接，可访问的公网URL
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`

	// 用户描述，最多200字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Name")
	delete(f, "Uid")
	delete(f, "Type")
	delete(f, "Password")
	delete(f, "UserStatus")
	delete(f, "NickName")
	delete(f, "Phone")
	delete(f, "Email")
	delete(f, "AvatarUrl")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserResp struct {
	// 用户ID
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`
}

// Predefined struct for user
type CreateUserResponseParams struct {
	// 结果返回
	Data *CreateUserResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserResponseParams `json:"Response"`
}

func (r *CreateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVmInstanceRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务器类型：
	// LightHouse = 轻量云服务器
	// CVM = 云服务器
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 轻量云服务器套餐ID。 当Type=LightHouse时必传
	LightHouseBundleId *string `json:"LightHouseBundleId,omitnil,omitempty" name:"LightHouseBundleId"`

	// 轻量云服务器镜像ID。当Type=LightHouse时必传
	LightHouseBlueprintId *string `json:"LightHouseBlueprintId,omitnil,omitempty" name:"LightHouseBlueprintId"`

	// 服务器别名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 登录方式
	LoginConfiguration *VMLoginConfiguration `json:"LoginConfiguration,omitnil,omitempty" name:"LoginConfiguration"`
}

type CreateVmInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务器类型：
	// LightHouse = 轻量云服务器
	// CVM = 云服务器
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 轻量云服务器套餐ID。 当Type=LightHouse时必传
	LightHouseBundleId *string `json:"LightHouseBundleId,omitnil,omitempty" name:"LightHouseBundleId"`

	// 轻量云服务器镜像ID。当Type=LightHouse时必传
	LightHouseBlueprintId *string `json:"LightHouseBlueprintId,omitnil,omitempty" name:"LightHouseBlueprintId"`

	// 服务器别名
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 登录方式
	LoginConfiguration *VMLoginConfiguration `json:"LoginConfiguration,omitnil,omitempty" name:"LoginConfiguration"`
}

func (r *CreateVmInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVmInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Type")
	delete(f, "LightHouseBundleId")
	delete(f, "LightHouseBlueprintId")
	delete(f, "InstanceName")
	delete(f, "LoginConfiguration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVmInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVmInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVmInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateVmInstanceResponseParams `json:"Response"`
}

func (r *CreateVmInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVmInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomLogConfig struct {
	// 是否需要请求体
	NeedReqBodyLog *bool `json:"NeedReqBodyLog,omitnil,omitempty" name:"NeedReqBodyLog"`

	// 是否需要请求头
	NeedReqHeaderLog *bool `json:"NeedReqHeaderLog,omitnil,omitempty" name:"NeedReqHeaderLog"`

	// 是否需要回包体
	NeedRspBodyLog *bool `json:"NeedRspBodyLog,omitnil,omitempty" name:"NeedRspBodyLog"`

	// 是否需要回包头部信息
	NeedRspHeaderLog *bool `json:"NeedRspHeaderLog,omitnil,omitempty" name:"NeedRspHeaderLog"`

	// cls set信息
	LogSetId *string `json:"LogSetId,omitnil,omitempty" name:"LogSetId"`

	// cls topicId
	LogTopicId *string `json:"LogTopicId,omitnil,omitempty" name:"LogTopicId"`
}

type DatabasesInfo struct {
	// 数据库唯一标识
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 状态。包含以下取值：
	// <li>INITIALIZING：资源初始化中</li>
	// <li>RUNNING：运行中，可正常使用的状态</li>
	// <li>UNUSABLE：禁用，不可用</li>
	// <li>OVERDUE：资源过期</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 所属地域。
	// 当前支持ap-shanghai
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type DbInstance struct {
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MySQL 连接器实例 ID；`"default"` 或为空表示使用 TCB 环境的默认连接器
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名；为空时使用连接器配置的默认数据库名
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`
}

// Predefined struct for user
type DeleteAIModelRequestParams struct {
	// <p>环境id</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>分组名列表</p>
	GroupNames []*string `json:"GroupNames,omitnil,omitempty" name:"GroupNames"`
}

type DeleteAIModelRequest struct {
	*tchttp.BaseRequest
	
	// <p>环境id</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>分组名列表</p>
	GroupNames []*string `json:"GroupNames,omitnil,omitempty" name:"GroupNames"`
}

func (r *DeleteAIModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAIModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "GroupNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAIModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAIModelResponseParams struct {
	// <p>成功删除数量</p>
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAIModelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAIModelResponseParams `json:"Response"`
}

func (r *DeleteAIModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAIModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApiKeyRequestParams struct {
	// 环境 ID，用于标识该密钥归属的云开发环境，不同环境之间的数据相互隔离
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 密钥的唯一标识符，用于精确定位指定的 API 密钥。可通过查询密钥列表接口获取
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

type DeleteApiKeyRequest struct {
	*tchttp.BaseRequest
	
	// 环境 ID，用于标识该密钥归属的云开发环境，不同环境之间的数据相互隔离
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 密钥的唯一标识符，用于精确定位指定的 API 密钥。可通过查询密钥列表接口获取
	KeyId *string `json:"KeyId,omitnil,omitempty" name:"KeyId"`
}

func (r *DeleteApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "KeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApiKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApiKeyResponseParams `json:"Response"`
}

func (r *DeleteApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuthDomainRequestParams struct {
	// 开发者的环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 域名ID列表，支持批量传递
	DomainIds []*string `json:"DomainIds,omitnil,omitempty" name:"DomainIds"`
}

type DeleteAuthDomainRequest struct {
	*tchttp.BaseRequest
	
	// 开发者的环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 域名ID列表，支持批量传递
	DomainIds []*string `json:"DomainIds,omitnil,omitempty" name:"DomainIds"`
}

func (r *DeleteAuthDomainRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuthDomainRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "DomainIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuthDomainRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuthDomainResponseParams struct {
	// 删除的域名个数
	Deleted *int64 `json:"Deleted,omitnil,omitempty" name:"Deleted"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAuthDomainResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuthDomainResponseParams `json:"Response"`
}

func (r *DeleteAuthDomainResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuthDomainResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHTTPServiceRouteRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 路径列表。为空则表示删除此域名和所有路由
	Paths []*string `json:"Paths,omitnil,omitempty" name:"Paths"`
}

type DeleteHTTPServiceRouteRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 路径列表。为空则表示删除此域名和所有路由
	Paths []*string `json:"Paths,omitnil,omitempty" name:"Paths"`
}

func (r *DeleteHTTPServiceRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHTTPServiceRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Domain")
	delete(f, "Paths")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteHTTPServiceRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHTTPServiceRouteResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteHTTPServiceRouteResponse struct {
	*tchttp.BaseResponse
	Response *DeleteHTTPServiceRouteResponseParams `json:"Response"`
}

func (r *DeleteHTTPServiceRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHTTPServiceRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProviderRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 认证源ID，比如：github, 格式必须为：2-32位小写英文字符串或数字
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteProviderRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 认证源ID，比如：github, 格式必须为：2-32位小写英文字符串或数字
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DeleteProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProviderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProviderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProviderResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteProviderResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProviderResponseParams `json:"Response"`
}

func (r *DeleteProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableRequestParams struct {
	// 待删除的表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// FlexDB实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoDB连接器配置
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

type DeleteTableRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// FlexDB实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoDB连接器配置
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

func (r *DeleteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableName")
	delete(f, "Tag")
	delete(f, "EnvId")
	delete(f, "MongoConnector")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTableResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTableResponseParams `json:"Response"`
}

func (r *DeleteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsersRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// tcb用户id列表, 一次最多支持删除100个
	Uids []*string `json:"Uids,omitnil,omitempty" name:"Uids"`
}

type DeleteUsersRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// tcb用户id列表, 一次最多支持删除100个
	Uids []*string `json:"Uids,omitnil,omitempty" name:"Uids"`
}

func (r *DeleteUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Uids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUsersResp struct {
	// 成功个数
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 失败个数
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`
}

// Predefined struct for user
type DeleteUsersResponseParams struct {
	// 删除用户结果
	Data *DeleteUsersResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUsersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUsersResponseParams `json:"Response"`
}

func (r *DeleteUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVmInstanceRequestParams struct {
	// 服务器实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DeleteVmInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 服务器实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *DeleteVmInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVmInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVmInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVmInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVmInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVmInstanceResponseParams `json:"Response"`
}

func (r *DeleteVmInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVmInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIModelsRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DescribeAIModelsRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *DescribeAIModelsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIModelsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIModelsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIModelsResponseParams struct {
	// 模型列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AIModels []*AIModelGroup `json:"AIModels,omitnil,omitempty" name:"AIModels"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIModelsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIModelsResponseParams `json:"Response"`
}

func (r *DescribeAIModelsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIModelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiKeyListRequestParams struct {
	// 环境 ID，用于标识该密钥归属的云开发环境，不同环境之间的数据相互隔离
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 分页查询的页码，从 1 开始。与 PageSize 配合使用，不传则默认返回第 1 页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页查询每页返回的记录条数。与 PageNumber 配合使用，不传则使用系统默认值
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 密钥类型过滤条件。可选值：api_key（服务端调用使用的 API 密钥，具有完整权限）、publish_key（客户端使用的公开密钥，权限受限）密钥类型过滤条件。不传默认值为api_key
	KeyType *string `json:"KeyType,omitnil,omitempty" name:"KeyType"`
}

type DescribeApiKeyListRequest struct {
	*tchttp.BaseRequest
	
	// 环境 ID，用于标识该密钥归属的云开发环境，不同环境之间的数据相互隔离
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 分页查询的页码，从 1 开始。与 PageSize 配合使用，不传则默认返回第 1 页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页查询每页返回的记录条数。与 PageNumber 配合使用，不传则使用系统默认值
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 密钥类型过滤条件。可选值：api_key（服务端调用使用的 API 密钥，具有完整权限）、publish_key（客户端使用的公开密钥，权限受限）密钥类型过滤条件。不传默认值为api_key
	KeyType *string `json:"KeyType,omitnil,omitempty" name:"KeyType"`
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
	delete(f, "EnvId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "KeyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApiKeyListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApiKeyListResponseParams struct {
	// API Key列表
	Data []*ApiKeyToken `json:"Data,omitnil,omitempty" name:"Data"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

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
type DescribeAuthDomainsRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DescribeAuthDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *DescribeAuthDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuthDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuthDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuthDomainsResponseParams struct {
	// 安全域名列表
	Domains []*AuthDomain `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuthDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuthDomainsResponseParams `json:"Response"`
}

func (r *DescribeAuthDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuthDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaasPackageListRequestParams struct {
	// tcb产品套餐ID，不填拉取全量package信息。
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 套餐归属方，填写后只返回对应的套餐 包含miniapp与qcloud两种 默认为miniapp
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 套餐归属环境渠道
	EnvChannel *string `json:"EnvChannel,omitnil,omitempty" name:"EnvChannel"`

	// 拉取套餐用途：
	// 1）new 新购
	// 2）modify变配
	// 3）renew续费
	TargetAction *string `json:"TargetAction,omitnil,omitempty" name:"TargetAction"`

	// 预留字段，同一商品会对应多个类型套餐，对指标有不同侧重。
	// 计算型calculation
	// 流量型flux
	// 容量型capactiy
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 类型分组过滤。默认为["default"]
	PackageTypeList []*string `json:"PackageTypeList,omitnil,omitempty" name:"PackageTypeList"`

	// 付费渠道，与回包billTags中的计费参数相关，不填返回默认值。
	PaymentChannel *string `json:"PaymentChannel,omitnil,omitempty" name:"PaymentChannel"`
}

type DescribeBaasPackageListRequest struct {
	*tchttp.BaseRequest
	
	// tcb产品套餐ID，不填拉取全量package信息。
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 套餐归属方，填写后只返回对应的套餐 包含miniapp与qcloud两种 默认为miniapp
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 套餐归属环境渠道
	EnvChannel *string `json:"EnvChannel,omitnil,omitempty" name:"EnvChannel"`

	// 拉取套餐用途：
	// 1）new 新购
	// 2）modify变配
	// 3）renew续费
	TargetAction *string `json:"TargetAction,omitnil,omitempty" name:"TargetAction"`

	// 预留字段，同一商品会对应多个类型套餐，对指标有不同侧重。
	// 计算型calculation
	// 流量型flux
	// 容量型capactiy
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 类型分组过滤。默认为["default"]
	PackageTypeList []*string `json:"PackageTypeList,omitnil,omitempty" name:"PackageTypeList"`

	// 付费渠道，与回包billTags中的计费参数相关，不填返回默认值。
	PaymentChannel *string `json:"PaymentChannel,omitnil,omitempty" name:"PaymentChannel"`
}

func (r *DescribeBaasPackageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaasPackageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PackageName")
	delete(f, "EnvId")
	delete(f, "Source")
	delete(f, "EnvChannel")
	delete(f, "TargetAction")
	delete(f, "GroupName")
	delete(f, "PackageTypeList")
	delete(f, "PaymentChannel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBaasPackageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBaasPackageListResponseParams struct {
	// 套餐列表
	PackageList []*BaasPackageInfo `json:"PackageList,omitnil,omitempty" name:"PackageList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBaasPackageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBaasPackageListResponseParams `json:"Response"`
}

func (r *DescribeBaasPackageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBaasPackageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingInfoRequestParams struct {
	// <p>环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>环境列表，当环境列表不为空时，查询的环境以该参数为准</p>
	EnvIds []*string `json:"EnvIds,omitnil,omitempty" name:"EnvIds"`

	// <p>每页条数（用于拉取列表时分页）</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeBillingInfoRequest struct {
	*tchttp.BaseRequest
	
	// <p>环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>环境列表，当环境列表不为空时，查询的环境以该参数为准</p>
	EnvIds []*string `json:"EnvIds,omitnil,omitempty" name:"EnvIds"`

	// <p>每页条数（用于拉取列表时分页）</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeBillingInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "EnvIds")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBillingInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBillingInfoResponseParams struct {
	// <p>环境计费信息列表</p>
	EnvBillingInfoList []*EnvBillingInfoItem `json:"EnvBillingInfoList,omitnil,omitempty" name:"EnvBillingInfoList"`

	// <p>总个数</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBillingInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBillingInfoResponseParams `json:"Response"`
}

func (r *DescribeBillingInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBillingInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 客户端的唯一标识符（Client ID），在 OAuth/OIDC 授权流程中作为 client_id 参数使用，创建后不可修改，一般使用环境id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DescribeClientRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 客户端的唯一标识符（Client ID），在 OAuth/OIDC 授权流程中作为 client_id 参数使用，创建后不可修改，一般使用环境id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DescribeClientRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClientRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientResponseParams struct {
	// 客户端的唯一标识符（Client ID），在 OAuth/OIDC 授权流程中作为 client_id 参数使用。创建时仅可传入环境 ID 或留空：传入环境 ID 时将直接使用该值作为客户端 ID（一个环境仅允许一个）；留空时由系统自动生成与环境 ID 关联的唯一 ID。创建后不可修改。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 客户端的创建时间，格式遵循 ISO 8601 标准（如：2024-01-01T00:00:00Z），由系统自动生成，不可手动修改。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 客户端信息的最后修改时间，格式遵循 ISO 8601 标准（如：2024-01-01T00:00:00Z），每次更新应用配置时由系统自动更新。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *string `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// Refresh Token 的有效期，单位为秒。超过该时间后 Refresh Token 将失效，用户需重新登录。取值范围：最小 1800 秒（30 分钟），最大 2592000 秒（30 天），超出上限将自动截断为 30 天。若不设置则默认为 30 天。当该值小于等于 7200 秒时，系统会自动将 AccessTokenExpiresIn 调整为 RefreshTokenExpiresIn - 660 秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RefreshTokenExpiresIn *int64 `json:"RefreshTokenExpiresIn,omitnil,omitempty" name:"RefreshTokenExpiresIn"`

	// Access Token 的有效期，单位为秒。超过该时间后 Access Token 将失效，需通过 Refresh Token 换取新的 Access Token。若不设置则默认为 7200 秒（2 小时）。设置值小于 1800 秒时将被忽略，使用系统默认值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessTokenExpiresIn *int64 `json:"AccessTokenExpiresIn,omitnil,omitempty" name:"AccessTokenExpiresIn"`

	// 单个用户在该客户端下允许同时登录的最大会话数量。取值范围：-1 至 50。-1 表示不限制设备数量；0 或不填默认按 User-Agent 区分设备（相同 User-Agent 共享一个会话，不同 User-Agent 各独立一个会话）；1 表示单设备登录，新登录将踢出旧会话；大于 1 时按真实设备 ID 限制，超出限制后最早登录的会话将被自动踢出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDevice *int64 `json:"MaxDevice,omitnil,omitempty" name:"MaxDevice"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClientResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClientResponseParams `json:"Response"`
}

func (r *DescribeClientResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseBuildServiceRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// build类型,枚举值有: cloudbaserun, framework-ci
	CIBusiness *string `json:"CIBusiness,omitnil,omitempty" name:"CIBusiness"`

	// 服务版本
	ServiceVersion *string `json:"ServiceVersion,omitnil,omitempty" name:"ServiceVersion"`

	// 文件后缀
	Suffix *string `json:"Suffix,omitnil,omitempty" name:"Suffix"`
}

type DescribeCloudBaseBuildServiceRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// build类型,枚举值有: cloudbaserun, framework-ci
	CIBusiness *string `json:"CIBusiness,omitnil,omitempty" name:"CIBusiness"`

	// 服务版本
	ServiceVersion *string `json:"ServiceVersion,omitnil,omitempty" name:"ServiceVersion"`

	// 文件后缀
	Suffix *string `json:"Suffix,omitnil,omitempty" name:"Suffix"`
}

func (r *DescribeCloudBaseBuildServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseBuildServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServiceName")
	delete(f, "CIBusiness")
	delete(f, "ServiceVersion")
	delete(f, "Suffix")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseBuildServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseBuildServiceResponseParams struct {
	// 上传url
	UploadUrl *string `json:"UploadUrl,omitnil,omitempty" name:"UploadUrl"`

	// 上传header
	UploadHeaders []*KVPair `json:"UploadHeaders,omitnil,omitempty" name:"UploadHeaders"`

	// 包名
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 包版本
	PackageVersion *string `json:"PackageVersion,omitnil,omitempty" name:"PackageVersion"`

	// 下载链接
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 下载Httpheader
	DownloadHeaders []*KVPair `json:"DownloadHeaders,omitnil,omitempty" name:"DownloadHeaders"`

	// 下载链接是否过期
	OutDate *bool `json:"OutDate,omitnil,omitempty" name:"OutDate"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudBaseBuildServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseBuildServiceResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseBuildServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseBuildServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunServerVersionRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名称
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 版本名称
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`
}

type DescribeCloudBaseRunServerVersionRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务名称
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 版本名称
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`
}

func (r *DescribeCloudBaseRunServerVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunServerVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ServerName")
	delete(f, "VersionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudBaseRunServerVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudBaseRunServerVersionResponseParams struct {
	// 版本名称
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// Dockerfile的路径
	DockerfilePath *string `json:"DockerfilePath,omitnil,omitempty" name:"DockerfilePath"`

	// DockerBuild的目录
	BuildDir *string `json:"BuildDir,omitnil,omitempty" name:"BuildDir"`

	// 请使用CPUSize
	Cpu *float64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 请使用MemSize
	Mem *float64 `json:"Mem,omitnil,omitempty" name:"Mem"`

	// 副本最小值
	MinNum *int64 `json:"MinNum,omitnil,omitempty" name:"MinNum"`

	// 副本最大值
	MaxNum *int64 `json:"MaxNum,omitnil,omitempty" name:"MaxNum"`

	// 策略类型
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 策略阈值
	PolicyThreshold *float64 `json:"PolicyThreshold,omitnil,omitempty" name:"PolicyThreshold"`

	// 环境变量
	EnvParams *string `json:"EnvParams,omitnil,omitempty" name:"EnvParams"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 更新时间
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// 版本的IP
	VersionIP *string `json:"VersionIP,omitnil,omitempty" name:"VersionIP"`

	// 版本的端口号
	VersionPort *int64 `json:"VersionPort,omitnil,omitempty" name:"VersionPort"`

	// 版本状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 代码包的名字
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// 代码版本的名字
	PackageVersion *string `json:"PackageVersion,omitnil,omitempty" name:"PackageVersion"`

	// 枚举（package/repository/image)
	UploadType *string `json:"UploadType,omitnil,omitempty" name:"UploadType"`

	// Repo的类型(gitlab/github/coding)
	RepoType *string `json:"RepoType,omitnil,omitempty" name:"RepoType"`

	// 地址
	Repo *string `json:"Repo,omitnil,omitempty" name:"Repo"`

	// 分支
	Branch *string `json:"Branch,omitnil,omitempty" name:"Branch"`

	// 服务名字
	ServerName *string `json:"ServerName,omitnil,omitempty" name:"ServerName"`

	// 是否对于外网开放
	IsPublic *bool `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`

	// vpc id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 日志采集路径
	CustomLogs *string `json:"CustomLogs,omitnil,omitempty" name:"CustomLogs"`

	// 监听端口
	ContainerPort *int64 `json:"ContainerPort,omitnil,omitempty" name:"ContainerPort"`

	// 延迟多长时间开始健康检查（单位s）
	InitialDelaySeconds *int64 `json:"InitialDelaySeconds,omitnil,omitempty" name:"InitialDelaySeconds"`

	// 镜像地址
	ImageUrl *string `json:"ImageUrl,omitnil,omitempty" name:"ImageUrl"`

	// CPU 大小
	CpuSize *float64 `json:"CpuSize,omitnil,omitempty" name:"CpuSize"`

	// MEM 大小
	MemSize *float64 `json:"MemSize,omitnil,omitempty" name:"MemSize"`

	// 是否有Dockerfile：0-default has, 1-has, 2-has not
	HasDockerfile *int64 `json:"HasDockerfile,omitnil,omitempty" name:"HasDockerfile"`

	// 基础镜像
	BaseImage *string `json:"BaseImage,omitnil,omitempty" name:"BaseImage"`

	// 容器启动入口命令
	EntryPoint *string `json:"EntryPoint,omitnil,omitempty" name:"EntryPoint"`

	// 仓库语言
	RepoLanguage *string `json:"RepoLanguage,omitnil,omitempty" name:"RepoLanguage"`

	// 自动扩缩容策略组
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyDetail []*HpaPolicy `json:"PolicyDetail,omitnil,omitempty" name:"PolicyDetail"`

	// Tke集群信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TkeClusterInfo *TkeClusterInfo `json:"TkeClusterInfo,omitnil,omitempty" name:"TkeClusterInfo"`

	// 版本工作负载类型；deployment/deamonset
	TkeWorkloadType *string `json:"TkeWorkloadType,omitnil,omitempty" name:"TkeWorkloadType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudBaseRunServerVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudBaseRunServerVersionResponseParams `json:"Response"`
}

func (r *DescribeCloudBaseRunServerVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudBaseRunServerVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCreateMySQLResult struct {
	// 状态 notexist | init | doing | success | fail
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailReason *string `json:"FailReason,omitnil,omitempty" name:"FailReason"`

	// 是否已被冻结（只在 Status=success时有效）
	FreezeStatus *bool `json:"FreezeStatus,omitnil,omitempty" name:"FreezeStatus"`
}

// Predefined struct for user
type DescribeCreateMySQLResultRequestParams struct {
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// OpenMysql 返回任务 Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeCreateMySQLResultRequest struct {
	*tchttp.BaseRequest
	
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// OpenMysql 返回任务 Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeCreateMySQLResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCreateMySQLResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCreateMySQLResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCreateMySQLResultResponseParams struct {
	// 查询开通结果
	Data *DescribeCreateMySQLResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCreateMySQLResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCreateMySQLResultResponseParams `json:"Response"`
}

func (r *DescribeCreateMySQLResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCreateMySQLResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCurveDataRequestParams struct {
	// <p>环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <h4>文档型数据库相关指标</h4><li> DbRead: 数据库读请求数 </li><li> DbWrite: 数据库写请求数 </li><li> DbCostTime10ms: 数据库耗时在10ms-50ms请求数 </li><li> DbCostTime50ms: 数据库耗时在50ms-100ms请求数 </li><li> DbCostTime100ms: 数据库耗时在100ms以上请求数 </li><li> DbSizepkg: 数据库容量，单位MB </li><h4>SQL型数据库相关指标</h4><li> MysqlStorageUsage: 关系型数据库容量，单位MB </li><li> MysqlCCU: CCU </li><li> MysqlCpuUsageRate:CPU利用率 </li><li> MysqlDbConnections:数据库连接数 </li><li> MysqlMemoryUse: 内存使用量，单位MB </li><li> MysqlSlowQueries:慢查询数 </li><li> MysqlTps: 提交数 </li><li> MysqlQps: QPS </li><h4>云函数相关指标</h4><li> FunctionCU: 资源用量</li><li> FunctionInvocation: 调用次数 </li><li> FunctionFlux: 外网出流量, 单位千字节(KB) </li><li> FunctionThrottle: 受限次数 </li><li> FunctionConcurrentExecutions: 并发执行个数</li><li> FunctionTimeout: 函数执行超时次数</li><li> FunctionGBs: 资源用量, 单位Mb*Ms </li><li> FunctionError: 云错误次数 </li><li> FunctionDuration: 运行时间, 单位毫秒 </li><li> FunctionConcurrencyMemoryMB: 并发执行内存量 </li><li>FunctionMemOverFlow：内存超限次数</li><li> FunctionIdleProvisioned: 预置并发闲置量 </li><li> FunctionProvisionedConcurrency: 预置并发个数 </li><h4>云托管相关指标</h4><li>TkeRspTimeService ： 响应时间，单位毫秒</li><li>TkeCpuUsedService ： CPU使用量</li><li>TkeMemUsedService ： 内存使用量</li><li>TkeQPSService ： QPS</li><li>TkePodNumService ： 实例个数</li><li>TkeHttpServiceNatPkg ： 外网出流量，单位byte</li><li>TkeCUUsedService ： 内存使用量(CU单位)</li><li>TkeInvokeNumService ： 调用量</li><li>TkeHttpErrorService ： 错误响应（404、500等）</li><h4>静态网站托管相关指标</h4><li>StaticFsFluxPkg：流量，单位byte</li><li>StaticFsSizePkg：存储容量，单位MB</li><h4>身份认证相关指标</h4><li>AuthInvocationNumPkg：调用次数</li><h4>API调用相关指标</h4><li>GwCloudDevelopmentSecureCallsInvocation：云开发API调用次数</li><li>GwWXInvocation：小程序API调用次数</li><h4>HTTP网关相关指标</h4><li>GwCloudDevelopmentStandardCallsInvocation：HTTP调用次数</li><h4>大模型相关指标</h4><li>AIPromptTokenNumPkg：输入Token</li><li>AICompletionTokenNumPkg：输出Token</li><li>AITotalTokenNumPkg：总Token</li><h4>知识库相关指标</h4><li>KnowledgeBaseCapacity：容量，单位bytes</li><h4>用户登录相关指标</h4><li>DayActiveLoginAnonymousUser：匿名用户登录日活</li><li>DayActiveLoginAllUser ： 全部用户登录日活</li><li>DayActiveLoginExternalUser ： 外部用户登录日活</li><li>DayActiveLoginInternalUser ： 内部用户登录日活</li><h4>环境QPS相关指标</h4><li>EnvQPSAll：环境总QPS</li><h4>数据库连接器相关指标</h4><li> MongoConnectorRead: 数据库连接器读请求数 </li><li> MongoConnectorWrite: 数据库连接器写请求数 </li><li> MongoConnectorCostTime10ms: 数据库连接器耗时在10ms-50ms请求数 </li><li> MongoConnectorCostTime50ms: 数据库连接器耗时在50ms-100ms请求数 </li><li> MongoConnectorCostTime100ms: 数据库连接器耗时在100ms以上请求数 </li><li> MongoConnectorInvokeNum: 数据库连接器调用次数</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// <p>开始时间，如2018-08-24 10:50:00, 开始时间需要早于结束时间至少五分钟(原因是因为目前统计粒度最小是5分钟)</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，如2018-08-24 10:50:00, 结束时间需要晚于开始时间至少五分钟(原因是因为目前统计粒度最小是5分钟)</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>资源ID, 目前仅对文档型数据库、云函数、云托管、数据库连接器相关的指标有意义。<br>如果想查询某个具体云函数/具体数据库集合的指标，则需传入对应的云函数名称/集合名称；如果只想查询整个namespace的指标, 则留空或不传。<br>云托管相关指标的查询，必须传入云托管服务名称。<br>数据库连接器相关指标的查询，必须传入数据库连接器实例id</p>
	ResourceID *string `json:"ResourceID,omitnil,omitempty" name:"ResourceID"`

	// <p>微信AppId，微信必传</p>
	WxAppId *string `json:"WxAppId,omitnil,omitempty" name:"WxAppId"`

	// <p>子资源信息。<br>查询云托管相关指标的具体版本的监控数据，需传入。</p>
	SubresourceID *string `json:"SubresourceID,omitnil,omitempty" name:"SubresourceID"`

	// <p>网关路由</p>
	ThirdResource *string `json:"ThirdResource,omitnil,omitempty" name:"ThirdResource"`

	// <p>统计周期(单位秒)，非必传，传入时仅支持传入300，3600，86400。不传采用默认以下默认规则：当时间区间为1天内, 统计周期为300；当时间区间选择为1天以上, 15天以下, 统计周期为3600； 当时间区间选择为15天以上, 180天以下, 统计周期为86400。<br>如果传入period，需遵循以下规则。EndTime-StartTime的时间范围不超过1 天，Period可以取300或3600；EndTime-StartTime的时间范围满足超过1天且不超过3 天，Period可以取300或3600或86400；EndTime-StartTime的时间范围超过3天时，Period可以取3600或86400。</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`
}

type DescribeCurveDataRequest struct {
	*tchttp.BaseRequest
	
	// <p>环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <h4>文档型数据库相关指标</h4><li> DbRead: 数据库读请求数 </li><li> DbWrite: 数据库写请求数 </li><li> DbCostTime10ms: 数据库耗时在10ms-50ms请求数 </li><li> DbCostTime50ms: 数据库耗时在50ms-100ms请求数 </li><li> DbCostTime100ms: 数据库耗时在100ms以上请求数 </li><li> DbSizepkg: 数据库容量，单位MB </li><h4>SQL型数据库相关指标</h4><li> MysqlStorageUsage: 关系型数据库容量，单位MB </li><li> MysqlCCU: CCU </li><li> MysqlCpuUsageRate:CPU利用率 </li><li> MysqlDbConnections:数据库连接数 </li><li> MysqlMemoryUse: 内存使用量，单位MB </li><li> MysqlSlowQueries:慢查询数 </li><li> MysqlTps: 提交数 </li><li> MysqlQps: QPS </li><h4>云函数相关指标</h4><li> FunctionCU: 资源用量</li><li> FunctionInvocation: 调用次数 </li><li> FunctionFlux: 外网出流量, 单位千字节(KB) </li><li> FunctionThrottle: 受限次数 </li><li> FunctionConcurrentExecutions: 并发执行个数</li><li> FunctionTimeout: 函数执行超时次数</li><li> FunctionGBs: 资源用量, 单位Mb*Ms </li><li> FunctionError: 云错误次数 </li><li> FunctionDuration: 运行时间, 单位毫秒 </li><li> FunctionConcurrencyMemoryMB: 并发执行内存量 </li><li>FunctionMemOverFlow：内存超限次数</li><li> FunctionIdleProvisioned: 预置并发闲置量 </li><li> FunctionProvisionedConcurrency: 预置并发个数 </li><h4>云托管相关指标</h4><li>TkeRspTimeService ： 响应时间，单位毫秒</li><li>TkeCpuUsedService ： CPU使用量</li><li>TkeMemUsedService ： 内存使用量</li><li>TkeQPSService ： QPS</li><li>TkePodNumService ： 实例个数</li><li>TkeHttpServiceNatPkg ： 外网出流量，单位byte</li><li>TkeCUUsedService ： 内存使用量(CU单位)</li><li>TkeInvokeNumService ： 调用量</li><li>TkeHttpErrorService ： 错误响应（404、500等）</li><h4>静态网站托管相关指标</h4><li>StaticFsFluxPkg：流量，单位byte</li><li>StaticFsSizePkg：存储容量，单位MB</li><h4>身份认证相关指标</h4><li>AuthInvocationNumPkg：调用次数</li><h4>API调用相关指标</h4><li>GwCloudDevelopmentSecureCallsInvocation：云开发API调用次数</li><li>GwWXInvocation：小程序API调用次数</li><h4>HTTP网关相关指标</h4><li>GwCloudDevelopmentStandardCallsInvocation：HTTP调用次数</li><h4>大模型相关指标</h4><li>AIPromptTokenNumPkg：输入Token</li><li>AICompletionTokenNumPkg：输出Token</li><li>AITotalTokenNumPkg：总Token</li><h4>知识库相关指标</h4><li>KnowledgeBaseCapacity：容量，单位bytes</li><h4>用户登录相关指标</h4><li>DayActiveLoginAnonymousUser：匿名用户登录日活</li><li>DayActiveLoginAllUser ： 全部用户登录日活</li><li>DayActiveLoginExternalUser ： 外部用户登录日活</li><li>DayActiveLoginInternalUser ： 内部用户登录日活</li><h4>环境QPS相关指标</h4><li>EnvQPSAll：环境总QPS</li><h4>数据库连接器相关指标</h4><li> MongoConnectorRead: 数据库连接器读请求数 </li><li> MongoConnectorWrite: 数据库连接器写请求数 </li><li> MongoConnectorCostTime10ms: 数据库连接器耗时在10ms-50ms请求数 </li><li> MongoConnectorCostTime50ms: 数据库连接器耗时在50ms-100ms请求数 </li><li> MongoConnectorCostTime100ms: 数据库连接器耗时在100ms以上请求数 </li><li> MongoConnectorInvokeNum: 数据库连接器调用次数</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// <p>开始时间，如2018-08-24 10:50:00, 开始时间需要早于结束时间至少五分钟(原因是因为目前统计粒度最小是5分钟)</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间，如2018-08-24 10:50:00, 结束时间需要晚于开始时间至少五分钟(原因是因为目前统计粒度最小是5分钟)</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>资源ID, 目前仅对文档型数据库、云函数、云托管、数据库连接器相关的指标有意义。<br>如果想查询某个具体云函数/具体数据库集合的指标，则需传入对应的云函数名称/集合名称；如果只想查询整个namespace的指标, 则留空或不传。<br>云托管相关指标的查询，必须传入云托管服务名称。<br>数据库连接器相关指标的查询，必须传入数据库连接器实例id</p>
	ResourceID *string `json:"ResourceID,omitnil,omitempty" name:"ResourceID"`

	// <p>微信AppId，微信必传</p>
	WxAppId *string `json:"WxAppId,omitnil,omitempty" name:"WxAppId"`

	// <p>子资源信息。<br>查询云托管相关指标的具体版本的监控数据，需传入。</p>
	SubresourceID *string `json:"SubresourceID,omitnil,omitempty" name:"SubresourceID"`

	// <p>网关路由</p>
	ThirdResource *string `json:"ThirdResource,omitnil,omitempty" name:"ThirdResource"`

	// <p>统计周期(单位秒)，非必传，传入时仅支持传入300，3600，86400。不传采用默认以下默认规则：当时间区间为1天内, 统计周期为300；当时间区间选择为1天以上, 15天以下, 统计周期为3600； 当时间区间选择为15天以上, 180天以下, 统计周期为86400。<br>如果传入period，需遵循以下规则。EndTime-StartTime的时间范围不超过1 天，Period可以取300或3600；EndTime-StartTime的时间范围满足超过1天且不超过3 天，Period可以取300或3600或86400；EndTime-StartTime的时间范围超过3天时，Period可以取3600或86400。</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`
}

func (r *DescribeCurveDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCurveDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "MetricName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ResourceID")
	delete(f, "WxAppId")
	delete(f, "SubresourceID")
	delete(f, "ThirdResource")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCurveDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCurveDataResponseParams struct {
	// <p>开始时间, 会根据数据的统计周期进行取整</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>结束时间, 会根据数据的统计周期进行取整</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>指标名</p>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// <p>统计周期(单位秒), 当时间区间为1天内, 统计周期为5分钟; 当时间区间选择为1天以上, 15天以下, 统计周期为1小时; 当时间区间选择为15天以上, 180天以下, 统计周期为1天</p>
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// <p>有效的监控数据, 每个有效监控数据的上报时间可以从时间数组中的对应位置上获取到</p>
	Values []*int64 `json:"Values,omitnil,omitempty" name:"Values"`

	// <p>各数据点的时间戳数组（Unix 时间戳，秒级），与  ⁠Values⁠  一一对应</p>
	Time []*int64 `json:"Time,omitnil,omitempty" name:"Time"`

	// <p>有效的监控数据, 每个有效监控数据的上报时间可以从时间数组中的对应位置上获取到</p>
	NewValues []*float64 `json:"NewValues,omitnil,omitempty" name:"NewValues"`

	// <p>聚合方式， ⁠&quot;last&quot;⁠  表示取时间段内最后一个值，“max”表示取时间段内最大值，“avg”表示取时间段内的平均值</p>
	Statistics *string `json:"Statistics,omitnil,omitempty" name:"Statistics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCurveDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCurveDataResponseParams `json:"Response"`
}

func (r *DescribeCurveDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCurveDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseACLRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitnil,omitempty" name:"CollectionName"`
}

type DescribeDatabaseACLRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitnil,omitempty" name:"CollectionName"`
}

func (r *DescribeDatabaseACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "CollectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseACLResponseParams struct {
	// 权限标签。包含以下取值：
	// <li> READONLY：所有用户可读，仅创建者和管理员可写</li>
	// <li> PRIVATE：仅创建者及管理员可读写</li>
	// <li> ADMINWRITE：所有用户可读，仅管理员可写</li>
	// <li> ADMINONLY：仅管理员可读写</li>
	AclTag *string `json:"AclTag,omitnil,omitempty" name:"AclTag"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabaseACLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseACLResponseParams `json:"Response"`
}

func (r *DescribeDatabaseACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvAccountCircleRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DescribeEnvAccountCircleRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *DescribeEnvAccountCircleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvAccountCircleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvAccountCircleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvAccountCircleResponseParams struct {
	// 环境计费周期开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 环境计费周期结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEnvAccountCircleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvAccountCircleResponseParams `json:"Response"`
}

func (r *DescribeEnvAccountCircleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvAccountCircleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvLimitRequestParams struct {

}

type DescribeEnvLimitRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeEnvLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvLimitResponseParams struct {
	// 环境总数上限
	MaxEnvNum *int64 `json:"MaxEnvNum,omitnil,omitempty" name:"MaxEnvNum"`

	// 目前环境总数
	CurrentEnvNum *int64 `json:"CurrentEnvNum,omitnil,omitempty" name:"CurrentEnvNum"`

	// 免费环境数量上限
	MaxFreeEnvNum *int64 `json:"MaxFreeEnvNum,omitnil,omitempty" name:"MaxFreeEnvNum"`

	// 目前免费环境数量
	CurrentFreeEnvNum *int64 `json:"CurrentFreeEnvNum,omitnil,omitempty" name:"CurrentFreeEnvNum"`

	// 总计允许销毁环境次数上限
	MaxDeleteTotal *int64 `json:"MaxDeleteTotal,omitnil,omitempty" name:"MaxDeleteTotal"`

	// 目前已销毁环境次数
	CurrentDeleteTotal *int64 `json:"CurrentDeleteTotal,omitnil,omitempty" name:"CurrentDeleteTotal"`

	// 每月允许销毁环境次数上限
	MaxDeleteMonthly *int64 `json:"MaxDeleteMonthly,omitnil,omitempty" name:"MaxDeleteMonthly"`

	// 本月已销毁环境次数
	CurrentDeleteMonthly *int64 `json:"CurrentDeleteMonthly,omitnil,omitempty" name:"CurrentDeleteMonthly"`

	// 微信网关体验版可购买月份数
	MaxFreeTrialNum *int64 `json:"MaxFreeTrialNum,omitnil,omitempty" name:"MaxFreeTrialNum"`

	// 微信网关体验版已购买月份数
	CurrentFreeTrialNum *int64 `json:"CurrentFreeTrialNum,omitnil,omitempty" name:"CurrentFreeTrialNum"`

	// 转支付限额总数
	ChangePayTotal *int64 `json:"ChangePayTotal,omitnil,omitempty" name:"ChangePayTotal"`

	// 当前已用转支付次数
	CurrentChangePayTotal *int64 `json:"CurrentChangePayTotal,omitnil,omitempty" name:"CurrentChangePayTotal"`

	// 转支付每月限额
	ChangePayMonthly *int64 `json:"ChangePayMonthly,omitnil,omitempty" name:"ChangePayMonthly"`

	// 本月已用转支付额度
	CurrentChangePayMonthly *int64 `json:"CurrentChangePayMonthly,omitnil,omitempty" name:"CurrentChangePayMonthly"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEnvLimitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvLimitResponseParams `json:"Response"`
}

func (r *DescribeEnvLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvsRequestParams struct {
	// <p>环境ID，如果传了这个参数则只返回该环境的相关信息</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>指定Channels字段为可见渠道列表或不可见渠道列表<br>如只想获取渠道A的环境 就填写IsVisible= true,Channels = [&quot;A&quot;], 过滤渠道A拉取其他渠道环境时填写IsVisible= false,Channels = [&quot;A&quot;]</p>
	IsVisible *bool `json:"IsVisible,omitnil,omitempty" name:"IsVisible"`

	// <p>渠道列表，代表可见或不可见渠道由IsVisible参数指定</p>
	Channels []*string `json:"Channels,omitnil,omitempty" name:"Channels"`

	// <p>分页参数，单页限制个数</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>分页参数，偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeEnvsRequest struct {
	*tchttp.BaseRequest
	
	// <p>环境ID，如果传了这个参数则只返回该环境的相关信息</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>指定Channels字段为可见渠道列表或不可见渠道列表<br>如只想获取渠道A的环境 就填写IsVisible= true,Channels = [&quot;A&quot;], 过滤渠道A拉取其他渠道环境时填写IsVisible= false,Channels = [&quot;A&quot;]</p>
	IsVisible *bool `json:"IsVisible,omitnil,omitempty" name:"IsVisible"`

	// <p>渠道列表，代表可见或不可见渠道由IsVisible参数指定</p>
	Channels []*string `json:"Channels,omitnil,omitempty" name:"Channels"`

	// <p>分页参数，单页限制个数</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>分页参数，偏移量</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeEnvsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "IsVisible")
	delete(f, "Channels")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvsResponseParams struct {
	// <p>环境信息列表</p>
	EnvList []*EnvInfo `json:"EnvList,omitnil,omitempty" name:"EnvList"`

	// <p>环境个数</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEnvsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvsResponseParams `json:"Response"`
}

func (r *DescribeEnvsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayVersionsRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 网关id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 版本名
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`
}

type DescribeGatewayVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 网关id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 版本名
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`
}

func (r *DescribeGatewayVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "GatewayId")
	delete(f, "VersionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGatewayVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGatewayVersionsResponseParams struct {
	// 网关id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 版本总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 版本信息详情
	GatewayVersionItems []*GatewayVersionItem `json:"GatewayVersionItems,omitnil,omitempty" name:"GatewayVersionItems"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGatewayVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGatewayVersionsResponseParams `json:"Response"`
}

func (r *DescribeGatewayVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGatewayVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHTTPServiceRouteRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 过滤条件。Key的含义参考对应字段，Value精确匹配。可过滤: Domain、Path、DomainType、UpstreamResourceType。可过滤的Values单条不超过100
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页偏移量。默认 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制。默认20，最大值1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeHTTPServiceRouteRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 过滤条件。Key的含义参考对应字段，Value精确匹配。可过滤: Domain、Path、DomainType、UpstreamResourceType。可过滤的Values单条不超过100
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页偏移量。默认 0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制。默认20，最大值1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeHTTPServiceRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHTTPServiceRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHTTPServiceRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHTTPServiceRouteResponseParams struct {
	// 域名路由信息列表
	Domains []*HTTPServiceDomain `json:"Domains,omitnil,omitempty" name:"Domains"`

	// 自定义接入的源站域名（HTTPService接入层域名）
	OriginDomain *string `json:"OriginDomain,omitnil,omitempty" name:"OriginDomain"`

	// 域名总数，分页查询使用总数判断是否已经拉取到所有数据
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHTTPServiceRouteResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHTTPServiceRouteResponseParams `json:"Response"`
}

func (r *DescribeHTTPServiceRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHTTPServiceRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostingDomainTaskRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DescribeHostingDomainTaskRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *DescribeHostingDomainTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostingDomainTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHostingDomainTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHostingDomainTaskResponseParams struct {
	// todo/doing/done/error
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHostingDomainTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHostingDomainTaskResponseParams `json:"Response"`
}

func (r *DescribeHostingDomainTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHostingDomainTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoginConfigRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DescribeLoginConfigRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *DescribeLoginConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoginConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoginConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoginConfigResponseParams struct {
	// 是否开启邮箱登录方式。true 表示已开启，允许用户使用邮箱和密码进行登录；false 表示已关闭。
	EmailLogin *bool `json:"EmailLogin,omitnil,omitempty" name:"EmailLogin"`

	// 是否开启匿名登录方式。true 表示已开启，允许用户无需注册即可以匿名身份登录；false 表示已关闭。
	AnonymousLogin *bool `json:"AnonymousLogin,omitnil,omitempty" name:"AnonymousLogin"`

	// 是否开启用户名密码登录方式。true 表示已开启，允许用户使用用户名和密码进行登录；false 表示已关闭。
	UserNameLogin *bool `json:"UserNameLogin,omitnil,omitempty" name:"UserNameLogin"`

	// 短信验证码发送配置，包含短信发送通道类型、自定义 APIs 数据源、调用方法及每日发送限额等信息。
	SmsVerificationConfig *VerificationConfig `json:"SmsVerificationConfig,omitnil,omitempty" name:"SmsVerificationConfig"`

	// 是否开启手机号短信登录方式。true 表示已开启，允许用户使用手机号和短信验证码进行登录和注册；false 表示已关闭。
	PhoneNumberLogin *bool `json:"PhoneNumberLogin,omitnil,omitempty" name:"PhoneNumberLogin"`

	// MFA 多因子认证登录配置，包含 MFA 开关及各验证方式（短信、邮箱、TOTP、强制绑定手机号）的启用状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MfaConfig *MFALoginConfig `json:"MfaConfig,omitnil,omitempty" name:"MfaConfig"`

	// 密码修改策略配置，包含首次登录强制修改密码开关及定期修改密码策略（周期和时间单位）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PwdUpdateStrategy *PasswordUpdateLoginConfig `json:"PwdUpdateStrategy,omitnil,omitempty" name:"PwdUpdateStrategy"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLoginConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoginConfigResponseParams `json:"Response"`
}

func (r *DescribeLoginConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoginConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeManagedAIModelListRequestParams struct {
	// <p>环境id</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DescribeManagedAIModelListRequest struct {
	*tchttp.BaseRequest
	
	// <p>环境id</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *DescribeManagedAIModelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeManagedAIModelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeManagedAIModelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeManagedAIModelListResponseParams struct {
	// <p>托管模型列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ManagedAIModelList []*ManagedAIModelGroup `json:"ManagedAIModelList,omitnil,omitempty" name:"ManagedAIModelList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeManagedAIModelListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeManagedAIModelListResponseParams `json:"Response"`
}

func (r *DescribeManagedAIModelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeManagedAIModelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMySQLClusterDetailRequestParams struct {
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DescribeMySQLClusterDetailRequest struct {
	*tchttp.BaseRequest
	
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *DescribeMySQLClusterDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMySQLClusterDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMySQLClusterDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMySQLClusterDetailResponseParams struct {
	// 集群详情
	Data *MySQLClusterDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMySQLClusterDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMySQLClusterDetailResponseParams `json:"Response"`
}

func (r *DescribeMySQLClusterDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMySQLClusterDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMySQLTaskStatusRequestParams struct {
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`
}

type DescribeMySQLTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`
}

func (r *DescribeMySQLTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMySQLTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "TaskId")
	delete(f, "TaskName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMySQLTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMySQLTaskStatusResponseParams struct {
	// 任务状态
	Data *MySQLTaskStatus `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMySQLTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMySQLTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeMySQLTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMySQLTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePGUserMigrationRequestParams struct {
	// <p>云开发环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>版本号</p><p>参数格式：14位时间格式</p><p>入参限制：纯数字</p>
	MigrationVersion *string `json:"MigrationVersion,omitnil,omitempty" name:"MigrationVersion"`
}

type DescribePGUserMigrationRequest struct {
	*tchttp.BaseRequest
	
	// <p>云开发环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>版本号</p><p>参数格式：14位时间格式</p><p>入参限制：纯数字</p>
	MigrationVersion *string `json:"MigrationVersion,omitnil,omitempty" name:"MigrationVersion"`
}

func (r *DescribePGUserMigrationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePGUserMigrationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "MigrationVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePGUserMigrationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePGUserMigrationResponseParams struct {
	// <p>版本号</p><p>参数格式：纯数字，14位时间格式</p>
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// <p>版本名</p><p>参数格式：只允许小写字母和下划线</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>要执行的migration sql 语句</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>回滚的sql 语句</p>
	Rollback *string `json:"Rollback,omitnil,omitempty" name:"Rollback"`

	// <p>migration query 语句的checksum值</p><p>由服务端自动生成，同版本 checksum 不一致会拒绝执行</p>
	Checksum *string `json:"Checksum,omitnil,omitempty" name:"Checksum"`

	// <p>用于标记调用来源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>用于标记该条migration由谁创建，目前默认调用的用户uin</p>
	CreatedBy *string `json:"CreatedBy,omitnil,omitempty" name:"CreatedBy"`

	// <p>该migration创建时间</p>
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// <p>该migration应用时间</p>
	AppliedAt *string `json:"AppliedAt,omitnil,omitempty" name:"AppliedAt"`

	// <p>该migration执行耗时</p><p>单位：毫秒</p>
	DurationMs *int64 `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePGUserMigrationResponse struct {
	*tchttp.BaseResponse
	Response *DescribePGUserMigrationResponseParams `json:"Response"`
}

func (r *DescribePGUserMigrationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePGUserMigrationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuotaDataRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <li> 指标名: </li>
	// <li> StorageSizepkg: 当月存储空间容量, 单位MB </li>
	// <li> StorageReadpkg: 当月存储读请求次数 </li>
	// <li> StorageWritepkg: 当月存储写请求次数 </li>
	// <li> StorageCdnOriginFluxpkg: 当月CDN回源流量, 单位字节 </li>
	// <li> StorageCdnOriginFluxpkgDay: 当日CDN回源流量, 单位字节 </li>
	// <li> StorageReadpkgDay: 当日存储读请求次数 </li>
	// <li> StorageWritepkgDay: 当日写请求次数 </li>
	// <li> CDNFluxpkg: 当月CDN流量, 单位为字节 </li>
	// <li> CDNFluxpkgDay: 当日CDN流量, 单位为字节 </li>
	// <li> FunctionInvocationpkg: 当月云函数调用次数 </li>
	// <li> FunctionGBspkg: 当月云函数资源使用量, 单位Mb*Ms </li>
	// <li> FunctionFluxpkg: 当月云函数流量, 单位千字节(KB) </li>
	// <li> FunctionInvocationpkgDay: 当日云函数调用次数 </li>
	// <li> FunctionGBspkgDay: 当日云函数资源使用量, 单位Mb*Ms </li>
	// <li> FunctionFluxpkgDay: 当日云函数流量, 单位千字节(KB) </li>
	// <li> DbSizepkg: 当月数据库容量大小, 单位MB </li>
	// <li> DbReadpkg: 当日数据库读请求数 </li>
	// <li> DbWritepkg: 当日数据库写请求数 </li>
	// <li> StaticFsFluxPkgDay: 当日静态托管流量 </li>
	// <li> StaticFsFluxPkg: 当月静态托管流量</li>
	// <li> StaticFsSizePkg: 当月静态托管容量 </li>
	// <li> TkeCpuUsedPkg: 当月容器托管CPU使用量，单位核*秒 </li>
	// <li> TkeCpuUsedPkgDay: 当天容器托管CPU使用量，单位核*秒 </li>
	// <li> TkeMemUsedPkg: 当月容器托管内存使用量，单位MB*秒 </li>
	// <li> TkeMemUsedPkgDay: 当天容器托管内存使用量，单位MB*秒 </li>
	// <li> CodingBuildTimePkgDay: 当天容器托管构建时间使用量，单位毫秒 </li>
	// <li> TkeHttpServiceNatPkgDay: 当天容器托管流量使用量，单位B </li>
	// <li> CynosdbCcupkg: 当月微信云托管MySQL CCU使用量，单位个  （需要除以1000）</li>
	// <li> CynosdbStoragepkg: 当月微信云托管MySQL 存储使用量，单位MB  （需要除以1000）</li>
	// <li> CynosdbCcupkgDay: 当天微信云托管MySQL 存储使用量，单位个 （需要除以1000） </li>
	// <li> CynosdbStoragepkgDay: 当天微信云托管MySQL 存储使用量，单位MB （需要除以1000） </li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 资源ID, 目前仅对云函数、容器托管相关的指标有意义。云函数(FunctionInvocationpkg, FunctionGBspkg, FunctionFluxpkg)、容器托管（服务名称）。如果想查询某个云函数的指标则在ResourceId中传入函数名; 如果只想查询整个namespace的指标, 则留空或不传。
	ResourceID *string `json:"ResourceID,omitnil,omitempty" name:"ResourceID"`
}

type DescribeQuotaDataRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <li> 指标名: </li>
	// <li> StorageSizepkg: 当月存储空间容量, 单位MB </li>
	// <li> StorageReadpkg: 当月存储读请求次数 </li>
	// <li> StorageWritepkg: 当月存储写请求次数 </li>
	// <li> StorageCdnOriginFluxpkg: 当月CDN回源流量, 单位字节 </li>
	// <li> StorageCdnOriginFluxpkgDay: 当日CDN回源流量, 单位字节 </li>
	// <li> StorageReadpkgDay: 当日存储读请求次数 </li>
	// <li> StorageWritepkgDay: 当日写请求次数 </li>
	// <li> CDNFluxpkg: 当月CDN流量, 单位为字节 </li>
	// <li> CDNFluxpkgDay: 当日CDN流量, 单位为字节 </li>
	// <li> FunctionInvocationpkg: 当月云函数调用次数 </li>
	// <li> FunctionGBspkg: 当月云函数资源使用量, 单位Mb*Ms </li>
	// <li> FunctionFluxpkg: 当月云函数流量, 单位千字节(KB) </li>
	// <li> FunctionInvocationpkgDay: 当日云函数调用次数 </li>
	// <li> FunctionGBspkgDay: 当日云函数资源使用量, 单位Mb*Ms </li>
	// <li> FunctionFluxpkgDay: 当日云函数流量, 单位千字节(KB) </li>
	// <li> DbSizepkg: 当月数据库容量大小, 单位MB </li>
	// <li> DbReadpkg: 当日数据库读请求数 </li>
	// <li> DbWritepkg: 当日数据库写请求数 </li>
	// <li> StaticFsFluxPkgDay: 当日静态托管流量 </li>
	// <li> StaticFsFluxPkg: 当月静态托管流量</li>
	// <li> StaticFsSizePkg: 当月静态托管容量 </li>
	// <li> TkeCpuUsedPkg: 当月容器托管CPU使用量，单位核*秒 </li>
	// <li> TkeCpuUsedPkgDay: 当天容器托管CPU使用量，单位核*秒 </li>
	// <li> TkeMemUsedPkg: 当月容器托管内存使用量，单位MB*秒 </li>
	// <li> TkeMemUsedPkgDay: 当天容器托管内存使用量，单位MB*秒 </li>
	// <li> CodingBuildTimePkgDay: 当天容器托管构建时间使用量，单位毫秒 </li>
	// <li> TkeHttpServiceNatPkgDay: 当天容器托管流量使用量，单位B </li>
	// <li> CynosdbCcupkg: 当月微信云托管MySQL CCU使用量，单位个  （需要除以1000）</li>
	// <li> CynosdbStoragepkg: 当月微信云托管MySQL 存储使用量，单位MB  （需要除以1000）</li>
	// <li> CynosdbCcupkgDay: 当天微信云托管MySQL 存储使用量，单位个 （需要除以1000） </li>
	// <li> CynosdbStoragepkgDay: 当天微信云托管MySQL 存储使用量，单位MB （需要除以1000） </li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 资源ID, 目前仅对云函数、容器托管相关的指标有意义。云函数(FunctionInvocationpkg, FunctionGBspkg, FunctionFluxpkg)、容器托管（服务名称）。如果想查询某个云函数的指标则在ResourceId中传入函数名; 如果只想查询整个namespace的指标, 则留空或不传。
	ResourceID *string `json:"ResourceID,omitnil,omitempty" name:"ResourceID"`
}

func (r *DescribeQuotaDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotaDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "MetricName")
	delete(f, "ResourceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQuotaDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuotaDataResponseParams struct {
	// 指标名
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 指标的值
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`

	// 指标的附加值信息
	SubValue *string `json:"SubValue,omitnil,omitempty" name:"SubValue"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeQuotaDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQuotaDataResponseParams `json:"Response"`
}

func (r *DescribeQuotaDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQuotaDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePermissionRequestParams struct {
	// 环境 ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 资源类型：`function`-云函数、`storage`-云存储、`table`-SQL型数据库表、`collection`-文档型数据库表 `<br>`示例值：`table`。
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源标识列表。云函数不传或传空数组、云存储传存储桶名、数据库表传表名，不能超过100条。
	Resources []*string `json:"Resources,omitnil,omitempty" name:"Resources"`
}

type DescribeResourcePermissionRequest struct {
	*tchttp.BaseRequest
	
	// 环境 ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 资源类型：`function`-云函数、`storage`-云存储、`table`-SQL型数据库表、`collection`-文档型数据库表 `<br>`示例值：`table`。
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源标识列表。云函数不传或传空数组、云存储传存储桶名、数据库表传表名，不能超过100条。
	Resources []*string `json:"Resources,omitnil,omitempty" name:"Resources"`
}

func (r *DescribeResourcePermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ResourceType")
	delete(f, "Resources")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourcePermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourcePermissionResponseParams struct {
	// 查询资源权限返回结果
	Data *DescribeResourcePermissionResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResourcePermissionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourcePermissionResponseParams `json:"Response"`
}

func (r *DescribeResourcePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourcePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcePermissionResult struct {
	// 查询到的资源总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 资源权限列表
	PermissionList []*ResourcePermission `json:"PermissionList,omitnil,omitempty" name:"PermissionList"`
}

// Predefined struct for user
type DescribeSafeRuleRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitnil,omitempty" name:"CollectionName"`

	// 微信AppId，微信必传
	WxAppId *string `json:"WxAppId,omitnil,omitempty" name:"WxAppId"`
}

type DescribeSafeRuleRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitnil,omitempty" name:"CollectionName"`

	// 微信AppId，微信必传
	WxAppId *string `json:"WxAppId,omitnil,omitempty" name:"WxAppId"`
}

func (r *DescribeSafeRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSafeRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "CollectionName")
	delete(f, "WxAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSafeRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSafeRuleResponseParams struct {
	// 规则内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 权限标签。包含以下取值：
	// <li> READONLY：所有用户可读，仅创建者和管理员可写</li>
	// <li> PRIVATE：仅创建者及管理员可读写</li>
	// <li> ADMINWRITE：所有用户可读，仅管理员可写</li>
	// <li> ADMINONLY：仅管理员可读写</li>
	// <li> CUSTOM：自定义安全规则</li>
	AclTag *string `json:"AclTag,omitnil,omitempty" name:"AclTag"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSafeRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSafeRuleResponseParams `json:"Response"`
}

func (r *DescribeSafeRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSafeRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStaticStoreRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DescribeStaticStoreRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *DescribeStaticStoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaticStoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStaticStoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStaticStoreResponseParams struct {
	// 静态托管资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*StaticStoreInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStaticStoreResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStaticStoreResponseParams `json:"Response"`
}

func (r *DescribeStaticStoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStaticStoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableRequestParams struct {
	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// FlecDB实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoDB连接器配置
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

type DescribeTableRequest struct {
	*tchttp.BaseRequest
	
	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// FlecDB实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoDB连接器配置
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

func (r *DescribeTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableName")
	delete(f, "Tag")
	delete(f, "EnvId")
	delete(f, "MongoConnector")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableResponseParams struct {
	// 索引相关信息
	Indexes []*IndexInfo `json:"Indexes,omitnil,omitempty" name:"Indexes"`

	// 索引个数
	IndexNum *int64 `json:"IndexNum,omitnil,omitempty" name:"IndexNum"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTableResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableResponseParams `json:"Response"`
}

func (r *DescribeTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesRequestParams struct {
	// 分页条件
	MgoLimit *int64 `json:"MgoLimit,omitnil,omitempty" name:"MgoLimit"`

	// 实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 分页条件
	MgoOffset *int64 `json:"MgoOffset,omitnil,omitempty" name:"MgoOffset"`

	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoConnector
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`

	// 指定表名过滤，为空时返回所有表
	TableNames []*string `json:"TableNames,omitnil,omitempty" name:"TableNames"`
}

type DescribeTablesRequest struct {
	*tchttp.BaseRequest
	
	// 分页条件
	MgoLimit *int64 `json:"MgoLimit,omitnil,omitempty" name:"MgoLimit"`

	// 实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 分页条件
	MgoOffset *int64 `json:"MgoOffset,omitnil,omitempty" name:"MgoOffset"`

	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoConnector
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`

	// 指定表名过滤，为空时返回所有表
	TableNames []*string `json:"TableNames,omitnil,omitempty" name:"TableNames"`
}

func (r *DescribeTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MgoLimit")
	delete(f, "Tag")
	delete(f, "MgoOffset")
	delete(f, "EnvId")
	delete(f, "MongoConnector")
	delete(f, "TableNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesResponseParams struct {
	// 表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tables []*TableInfo `json:"Tables,omitnil,omitempty" name:"Tables"`

	// 分页信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pager *Pager `json:"Pager,omitnil,omitempty" name:"Pager"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTablesResponseParams `json:"Response"`
}

func (r *DescribeTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserListRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 页码，从1开始，默认1
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 每页数量，默认20，最大100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 用户名，模糊查询
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户昵称，模糊查询
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 手机号，模糊查询
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 邮箱，模糊查询
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`
}

type DescribeUserListRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 页码，从1开始，默认1
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// 每页数量，默认20，最大100
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 用户名，模糊查询
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户昵称，模糊查询
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 手机号，模糊查询
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 邮箱，模糊查询
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`
}

func (r *DescribeUserListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "Name")
	delete(f, "NickName")
	delete(f, "Phone")
	delete(f, "Email")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserListResp struct {
	// 用户总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 用户列表
	UserList []*User `json:"UserList,omitnil,omitempty" name:"UserList"`
}

// Predefined struct for user
type DescribeUserListResponseParams struct {
	// 结果返回
	Data *DescribeUserListResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserListResponseParams `json:"Response"`
}

func (r *DescribeUserListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVmInstancesRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务器类型： LightHouse = 轻量云服务器 CVM = 云服务器
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeVmInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 服务器类型： LightHouse = 轻量云服务器 CVM = 云服务器
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeVmInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVmInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVmInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVmInstancesResponseParams struct {
	// 主机实例列表
	InstanceList []*VmInstance `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVmInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVmInstancesResponseParams `json:"Response"`
}

func (r *DescribeVmInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVmInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVmSpecRequestParams struct {
	// 类型：
	// LightHouse = 轻量云服务器
	// CVM = 云服务器
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeVmSpecRequest struct {
	*tchttp.BaseRequest
	
	// 类型：
	// LightHouse = 轻量云服务器
	// CVM = 云服务器
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeVmSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVmSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVmSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVmSpecResponseParams struct {
	// 规格列表
	SpecList []*VMSpec `json:"SpecList,omitnil,omitempty" name:"SpecList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVmSpecResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVmSpecResponseParams `json:"Response"`
}

func (r *DescribeVmSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVmSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyEnvRequestParams struct {
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 针对预付费 删除隔离中的环境时要传true 正常环境直接跳过隔离期删除
	IsForce *bool `json:"IsForce,omitnil,omitempty" name:"IsForce"`

	// 是否绕过资源检查，资源包等额外资源，默认为false，如果为true，则不检查资源是否有数据，直接删除。
	BypassCheck *bool `json:"BypassCheck,omitnil,omitempty" name:"BypassCheck"`
}

type DestroyEnvRequest struct {
	*tchttp.BaseRequest
	
	// 环境Id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 针对预付费 删除隔离中的环境时要传true 正常环境直接跳过隔离期删除
	IsForce *bool `json:"IsForce,omitnil,omitempty" name:"IsForce"`

	// 是否绕过资源检查，资源包等额外资源，默认为false，如果为true，则不检查资源是否有数据，直接删除。
	BypassCheck *bool `json:"BypassCheck,omitnil,omitempty" name:"BypassCheck"`
}

func (r *DestroyEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "IsForce")
	delete(f, "BypassCheck")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyEnvResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyEnvResponse struct {
	*tchttp.BaseResponse
	Response *DestroyEnvResponseParams `json:"Response"`
}

func (r *DestroyEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyMySQLRequestParams struct {
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type DestroyMySQLRequest struct {
	*tchttp.BaseRequest
	
	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *DestroyMySQLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyMySQLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyMySQLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyMySQLResponseParams struct {
	// 销毁结果
	Data *DestroyMySQLResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyMySQLResponse struct {
	*tchttp.BaseResponse
	Response *DestroyMySQLResponseParams `json:"Response"`
}

func (r *DestroyMySQLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyMySQLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DestroyMySQLResult struct {
	// 是否成功
	IsSuccess *bool `json:"IsSuccess,omitnil,omitempty" name:"IsSuccess"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`
}

// Predefined struct for user
type DestroyStaticStoreRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// cdn域名
	CdnDomain *string `json:"CdnDomain,omitnil,omitempty" name:"CdnDomain"`
}

type DestroyStaticStoreRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// cdn域名
	CdnDomain *string `json:"CdnDomain,omitnil,omitempty" name:"CdnDomain"`
}

func (r *DestroyStaticStoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyStaticStoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "CdnDomain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyStaticStoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyStaticStoreResponseParams struct {
	// 条件任务结果(succ/fail)
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DestroyStaticStoreResponse struct {
	*tchttp.BaseResponse
	Response *DestroyStaticStoreResponseParams `json:"Response"`
}

func (r *DestroyStaticStoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyStaticStoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DropIndex struct {
	// 索引名称
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`
}

type EmailProviderConfig struct {
	// <p>smtp配置</p>
	SmtpConfig *EmailSmtpConfig `json:"SmtpConfig,omitnil,omitempty" name:"SmtpConfig"`

	// <p>可选：TRUE，FALSE，如果On为TRUE，则表示采用默认代发。</p>
	On *string `json:"On,omitnil,omitempty" name:"On"`

	// <p>邮件模板配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateConfig *EmailTemplateConfig `json:"TemplateConfig,omitnil,omitempty" name:"TemplateConfig"`
}

type EmailSmtpConfig struct {
	// 邮件发送者的邮箱地址，即收件人看到的发件人地址。需为有效的邮箱格式，且须与 SMTP 服务器的授权账号一致，否则可能被邮件服务商拒绝发送。例如：abc@example.com
	SenderAddress *string `json:"SenderAddress,omitnil,omitempty" name:"SenderAddress"`

	// SMTP 邮件服务器的域名或 IP 地址，用于建立邮件发送连接。不同邮件服务商的 SMTP 地址不同，例如 QQ 邮箱为 smtp.qq.com，Gmail 为 smtp.gmail.com，请以实际服务商提供的地址为准。
	ServerHost *string `json:"ServerHost,omitnil,omitempty" name:"ServerHost"`

	// SMTP 邮件服务器的端口号，需与所选安全模式（SecurityMode）匹配。常用端口：465（SSL 加密）、587（STARTTLS 加密）、25（无加密，不推荐）。建议优先使用 465 或 587 以保障传输安全。
	ServerPort *uint64 `json:"ServerPort,omitnil,omitempty" name:"ServerPort"`

	// SMTP 服务器的登录账号，通常为发件人的完整邮箱地址。部分邮件服务商支持使用独立的 SMTP 授权账号，请以实际服务商的要求为准。
	AccountUsername *string `json:"AccountUsername,omitnil,omitempty" name:"AccountUsername"`

	// SMTP 服务器的登录密码。注意：部分邮件服务商（如 QQ 邮箱、163 邮箱）不支持直接使用账号登录密码，需在邮箱设置中开启 SMTP 服务并生成专用的授权码，请以实际服务商的要求为准。
	AccountPassword *string `json:"AccountPassword,omitnil,omitempty" name:"AccountPassword"`

	// SMTP 连接的加密模式，用于保障邮件传输安全。可选值：AUTO（自动选择，优先使用安全连接）、SSL（全程 SSL/TLS 加密，通常配合端口 465 使用）、STARTSSL（通过 STARTTLS 命令升级为加密连接，通常配合端口 587 使用）、NO_SSL（不使用加密，仅建议在内网或测试环境中使用）。推荐使用 AUTO 或 SSL 以确保传输安全。
	SecurityMode *string `json:"SecurityMode,omitnil,omitempty" name:"SecurityMode"`
}

type EmailTemplateConfig struct {
	// <p>注册登录模板</p><p>入参限制：模板中必须包含{{.VerificationCode}}变量，用于邮件中验证码的展示，可选变量有{{.Usage}}、{{.ExpireMinutes}}、{{.Email}}。邮件模板中禁止包含 script、javascript、onclick、onload、iframe、link 标签及 CSS expression、CSS url() 等</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegisterSignIn *LocalizedTemplate `json:"RegisterSignIn,omitnil,omitempty" name:"RegisterSignIn"`

	// <p>默认模板</p><p>入参限制：模板中必须包含{{.VerificationCode}}变量，用于邮件中验证码的展示，可选变量有{{.Usage}}、{{.ExpireMinutes}}、{{.Email}}。邮件模板中禁止包含 script、javascript、onclick、onload、iframe、link 标签及 CSS expression、CSS url() 等</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultTpl *LocalizedTemplate `json:"DefaultTpl,omitnil,omitempty" name:"DefaultTpl"`
}

type EnvBillingInfoItem struct {
	// <p>环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>tcb产品套餐ID，参考DescribePackages接口的返回值。</p>
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// <p>自动续费标记</p>
	IsAutoRenew *bool `json:"IsAutoRenew,omitnil,omitempty" name:"IsAutoRenew"`

	// <p>状态。包含以下取值：</p><li> 空字符串：初始化中</li><li> NORMAL：正常</li><li> ISOLATE：隔离</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>支付方式。包含以下取值：</p><li> PREPAYMENT：预付费</li><li> POSTPAID：后付费</li>
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>隔离时间，最近一次隔离的时间</p>
	IsolatedTime *string `json:"IsolatedTime,omitnil,omitempty" name:"IsolatedTime"`

	// <p>过期时间，套餐即将到期的时间</p>
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// <p>创建时间，第一次接入计费方案的时间。</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间，计费信息最近一次更新的时间。</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>true表示从未升级过付费版。</p>
	IsAlwaysFree *bool `json:"IsAlwaysFree,omitnil,omitempty" name:"IsAlwaysFree"`

	// <p>付费渠道。</p><li> miniapp：小程序</li><li> qcloud：腾讯云</li>
	PaymentChannel *string `json:"PaymentChannel,omitnil,omitempty" name:"PaymentChannel"`

	// <p>最新的订单信息</p>
	OrderInfo *OrderInfo `json:"OrderInfo,omitnil,omitempty" name:"OrderInfo"`

	// <p>免费配额信息。</p>
	FreeQuota *string `json:"FreeQuota,omitnil,omitempty" name:"FreeQuota"`

	// <p>是否开启 <code>超过套餐额度部分转按量付费</code></p>
	EnableOverrun *bool `json:"EnableOverrun,omitnil,omitempty" name:"EnableOverrun"`

	// <p>环境套餐类型</p>
	ExtPackageType *string `json:"ExtPackageType,omitnil,omitempty" name:"ExtPackageType"`

	// <p>是否付费期环境，可取值：yes/no。</p>
	EnvCharged *string `json:"EnvCharged,omitnil,omitempty" name:"EnvCharged"`

	// <p>是否已激活，可取值：yes/no。</p>
	EnvActivated *string `json:"EnvActivated,omitnil,omitempty" name:"EnvActivated"`
}

type EnvInfo struct {
	// <p>账户下该环境唯一标识</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>环境来源。包含以下取值：</p><li>miniapp：微信小程序</li><li>qcloud ：腾讯云</li>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>环境别名，要以a-z开头，不能包含 a-zA-z0-9- 以外的字符</p>
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// <p>创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>最后修改时间</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>环境状态。包含以下取值：</p><li>NORMAL：正常可用</li><li>UNAVAILABLE：服务不可用，可能是尚未初始化或者初始化过程中</li>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>数据库列表</p>
	Databases []*DatabasesInfo `json:"Databases,omitnil,omitempty" name:"Databases"`

	// <p>存储列表</p>
	Storages []*StorageInfo `json:"Storages,omitnil,omitempty" name:"Storages"`

	// <p>函数列表</p>
	Functions []*FunctionInfo `json:"Functions,omitnil,omitempty" name:"Functions"`

	// <p>tcb产品套餐ID，参考DescribePackages接口的返回值。</p>
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// <p>套餐中文名称，参考DescribePackages接口的返回值。</p>
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`

	// <p>云日志服务列表</p>
	LogServices []*LogServiceInfo `json:"LogServices,omitnil,omitempty" name:"LogServices"`

	// <p>静态资源信息</p>
	StaticStorages []*StaticStorageInfo `json:"StaticStorages,omitnil,omitempty" name:"StaticStorages"`

	// <p>是否到期自动降为免费版</p>
	IsAutoDegrade *bool `json:"IsAutoDegrade,omitnil,omitempty" name:"IsAutoDegrade"`

	// <p>环境渠道</p>
	EnvChannel *string `json:"EnvChannel,omitnil,omitempty" name:"EnvChannel"`

	// <p>支付方式。包含以下取值：</p><li> prepayment：预付费</li><li> postpaid：后付费</li>
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// <p>是否为默认环境</p>
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// <p>环境所属地域</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>环境标签列表</p>
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>自定义日志服务</p>
	CustomLogServices []*ClsInfo `json:"CustomLogServices,omitnil,omitempty" name:"CustomLogServices"`

	// <p>环境类型：baas, run, hoting, weda</p>
	EnvType *string `json:"EnvType,omitnil,omitempty" name:"EnvType"`

	// <p>是否是dau新套餐</p>
	IsDauPackage *bool `json:"IsDauPackage,omitnil,omitempty" name:"IsDauPackage"`

	// <p>套餐类型:空\baas\tcbr</p>
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// <p>架构类型</p>
	ArchitectureType *string `json:"ArchitectureType,omitnil,omitempty" name:"ArchitectureType"`

	// <p>回收标志，默认为空</p>
	Recycle *string `json:"Recycle,omitnil,omitempty" name:"Recycle"`

	// <p>环境meta信息列表</p>
	Meta []*KVPair `json:"Meta,omitnil,omitempty" name:"Meta"`

	// <p>pg信息</p>
	PostgreSQL []*PostgreSQLInfo `json:"PostgreSQL,omitnil,omitempty" name:"PostgreSQL"`
}

// Predefined struct for user
type ExecutePGSqlRequestParams struct {
	// <p>云开发环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>要执行的SQL语句</p>
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// <p>指定 role 执行 SQL</p>
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`
}

type ExecutePGSqlRequest struct {
	*tchttp.BaseRequest
	
	// <p>云开发环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>要执行的SQL语句</p>
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// <p>指定 role 执行 SQL</p>
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`
}

func (r *ExecutePGSqlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecutePGSqlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Sql")
	delete(f, "Role")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecutePGSqlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecutePGSqlResponseParams struct {
	// <p>影响行数</p>
	AffectedRows *int64 `json:"AffectedRows,omitnil,omitempty" name:"AffectedRows"`

	// <p>字段名列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*string `json:"Columns,omitnil,omitempty" name:"Columns"`

	// <p>数据行。每一行数据都是一个JSON串，将JSON进行反序列化将得到了每列的值。值可能是 null 或者 字符串，如果是 null 说明该列的值为 &lt;null&gt;，如果是字符串则为该列的值的字符串表示形式。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rows []*string `json:"Rows,omitnil,omitempty" name:"Rows"`

	// <p>SQL执行耗时</p><p>单位：毫秒</p>
	ExecutionTimeMs *int64 `json:"ExecutionTimeMs,omitnil,omitempty" name:"ExecutionTimeMs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExecutePGSqlResponse struct {
	*tchttp.BaseResponse
	Response *ExecutePGSqlResponseParams `json:"Response"`
}

func (r *ExecutePGSqlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecutePGSqlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExternalStorage struct {
	// 桶名。
	// 当 Provider=cos 时，表示腾讯云对象存储桶。
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// Bucket所属地域。
	// 当 Provider=cos 时，表示腾讯云对象存储桶的所属地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 基础路径。
	// 绑定之后，用户访问云存储内的文件，后台会自动以BasePath作为前缀，拼接到所访问的文件中。
	// 例如：
	//   BasePath=my-cloudbase-path ， 当用户访问云存储内的 /tencentcloud.png 时，实际访问的完整路径是：/my-cloudbase-path/tencentcloud.png
	BasePath *string `json:"BasePath,omitnil,omitempty" name:"BasePath"`

	// 是否启用外部存储
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type Filter struct {
	// 需要过滤的字段。过滤条件数量限制为10。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FunctionInfo struct {
	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 所属地域。
	// 当前支持ap-shanghai
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type GatewayVersionItem struct {
	// 版本名
	VersionName *string `json:"VersionName,omitnil,omitempty" name:"VersionName"`

	// 版本流量权重
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 创建状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 更新时间
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// 构建ID
	BuildId *uint64 `json:"BuildId,omitnil,omitempty" name:"BuildId"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 优先级
	Priority *uint64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 是否默认版本
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 网关版本自定义配置
	CustomConfig *WxGatewayCustomConfig `json:"CustomConfig,omitnil,omitempty" name:"CustomConfig"`
}

// Predefined struct for user
type GetProvidersRequestParams struct {
	// 环境 ID，用于指定需要查询配置第三方身份源的云开发环境。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type GetProvidersRequest struct {
	*tchttp.BaseRequest
	
	// 环境 ID，用于指定需要查询配置第三方身份源的云开发环境。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *GetProvidersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProvidersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetProvidersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetProvidersResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 三方认证源列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*Provider `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetProvidersResponse struct {
	*tchttp.BaseResponse
	Response *GetProvidersResponseParams `json:"Response"`
}

func (r *GetProvidersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetProvidersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HTTPServiceDomain struct {
	// 域名
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 域名类型。 HTTPSERVICE: HTTP访问服务，CBR: 云托管服务，ANYSERVICE: 任意服务，AI_AGENT: AI agent，VM: 主机，INTEGRATION_CALLBACK: 集成回调
	DomainType *string `json:"DomainType,omitnil,omitempty" name:"DomainType"`

	// 绑定类型。默认DIRECT。DIRECT: 直连到HTTP访问服务， CDN: 接入云开发CDN，CUSTOM: 自定义接入类型（其他CDN或者WAF）
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// 证书ID。当前账户下SSL平台的证书ID
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 协议类型。默认HTTP_AND_HTTPS。HTTP_AND_HTTPS: 同时开启http和https，HTTP_TO_HTTPS: http重定向成https，HTTPS_TO_HTTP: https重定向成http。如果未配置证书无法访问https或者进行重定向
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 配置DNS解析的CNAME。根据AccessType返回不同的CNAME值。
	Cname *string `json:"Cname,omitnil,omitempty" name:"Cname"`

	// 是否是默认域名
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 域名开启状态
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 状态。PROCESSING、FAIL，SUCCESS。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// DNS解析状态。OK： 解析正常，INVALID：解析不正确，域名未解析到当前Cname域名。
	DNSStatus *string `json:"DNSStatus,omitnil,omitempty" name:"DNSStatus"`

	// HTTP访问服务路由信息
	Routes []*HTTPServiceRoute `json:"Routes,omitnil,omitempty" name:"Routes"`

	// 扩展字段，内部包含headers处理等
	Extension *HTTPServiceExtension `json:"Extension,omitnil,omitempty" name:"Extension"`

	// 域名创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 域名更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type HTTPServiceDomainParam struct {
	// <p>域名。全局唯一。如果域名在其他环境下占用或者腾讯云CDN占用，可能会导致创建失败</p>
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>绑定类型。默认DIRECT。DIRECT: 直连到HTTP访问服务， CDN: 接入云开发CDN，CUSTOM: 自定义接入类型（其他CDN或者WAF）</p>
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// <p>证书ID。当前账户下SSL平台的证书ID</p>
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// <p>协议类型。默认HTTP_AND_HTTPS。HTTP_AND_HTTPS: 同时开启http和https，HTTP_TO_HTTPS: http重定向成https，HTTPS_TO_HTTP: https重定向成http。如果未配置证书无法访问https或者进行重定向</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>自定义CNAME。对应AccessType: Custom</p>
	CustomCname *string `json:"CustomCname,omitnil,omitempty" name:"CustomCname"`

	// <p>域名开启状态，不传默认开启</p>
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// <p>创建/修改的HTTP访问服务路由列表。如果不传，仅创建或修改域名信息。列表最大支持传入20个</p>
	Routes []*HTTPServiceRouteParam `json:"Routes,omitnil,omitempty" name:"Routes"`

	// <p>扩展字段，内部包含headers处理等</p>
	Extension *HTTPServiceExtension `json:"Extension,omitnil,omitempty" name:"Extension"`
}

type HTTPServiceExtension struct {
	// 添加请求头列表
	HeadersHandler *HTTPServiceHeadersHandler `json:"HeadersHandler,omitnil,omitempty" name:"HeadersHandler"`
}

type HTTPServiceHeaderToAdd struct {
	// 添加头部的key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 添加头部的值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 添加头部的处理行为。默认：OVERWRITE_IF_EXISTS_OR_ADD。APPEND_IF_EXISTS_OR_ADD: 已存在时追加值，不存在时添加，ADD_IF_ABSENT:  仅在 header 不存在时添加，已存在时不做任何操作，OVERWRITE_IF_EXISTS_OR_ADD: 已存在时覆盖值，不存在时添加（默认值），OVERWRITE_IF_EXISTS: 仅在 header 已存在时覆盖值，不存在时不做任何操作
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`
}

type HTTPServiceHeadersHandler struct {
	// 添加请求头列表
	RequestHeadersToAdd []*HTTPServiceHeaderToAdd `json:"RequestHeadersToAdd,omitnil,omitempty" name:"RequestHeadersToAdd"`

	// 删除请求头列表
	RequestHeadersToRemove []*string `json:"RequestHeadersToRemove,omitnil,omitempty" name:"RequestHeadersToRemove"`

	// 添加返回头列表
	ResponseHeadersToAdd []*HTTPServiceHeaderToAdd `json:"ResponseHeadersToAdd,omitnil,omitempty" name:"ResponseHeadersToAdd"`

	// 删除返回头列表
	ResponseHeadersToRemove []*string `json:"ResponseHeadersToRemove,omitnil,omitempty" name:"ResponseHeadersToRemove"`
}

type HTTPServicePathRewrite struct {
	// 路径前缀重写。StaticStorePrefix、Prefix只能填一个
	Prefix *string `json:"Prefix,omitnil,omitempty" name:"Prefix"`
}

type HTTPServiceQPSPerClient struct {
	// 客户端维度限频标识。限制当前资源被单个客户端调用的频率，客户端标识支持 用户ID（UserID） 或 客户端 IP（ClientIP）。UserID 包括 云开发用户 ID 或 微信 openid，如果请求无 UserID 信息，则不会限制。
	LimitBy *string `json:"LimitBy,omitnil,omitempty" name:"LimitBy"`

	// 限制QPS值，每秒请求次数
	LimitValue *int64 `json:"LimitValue,omitnil,omitempty" name:"LimitValue"`
}

type HTTPServiceRoute struct {
	// 路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 路径重写
	PathRewrite *HTTPServicePathRewrite `json:"PathRewrite,omitnil,omitempty" name:"PathRewrite"`

	// 上游服务类型。SCF: 云函数，CBR: 云托管，STATIC_STORE: 静态托管，WEB_SCF: WEB云函数，LH: Lighthouse
	UpstreamResourceType *string `json:"UpstreamResourceType,omitnil,omitempty" name:"UpstreamResourceType"`

	// 上游服务名
	UpstreamResourceName *string `json:"UpstreamResourceName,omitnil,omitempty" name:"UpstreamResourceName"`

	// 是否开启安全域名
	EnableSafeDomain *bool `json:"EnableSafeDomain,omitnil,omitempty" name:"EnableSafeDomain"`

	// 是否开启身份认证
	EnableAuth *bool `json:"EnableAuth,omitnil,omitempty" name:"EnableAuth"`

	// 是否开启路径透传
	EnablePathTransmission *bool `json:"EnablePathTransmission,omitnil,omitempty" name:"EnablePathTransmission"`

	// QPS限频策略
	QPSPolicy *HTTPServiceRouteQPSPolicy `json:"QPSPolicy,omitnil,omitempty" name:"QPSPolicy"`

	// 是否开启路由
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 扩展字段，内部包含headers处理等
	Extension *HTTPServiceExtension `json:"Extension,omitnil,omitempty" name:"Extension"`

	// 路由创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 路由更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type HTTPServiceRouteParam struct {
	// 路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 上游服务类型。创建时必填，修改时可选填。SCF: 云函数，CBR: 云托管，STATIC_STORE: 静态托管，WEB_SCF: WEB云函数，LH: Lighthouse
	UpstreamResourceType *string `json:"UpstreamResourceType,omitnil,omitempty" name:"UpstreamResourceType"`

	// 上游服务名。创建时必填，修改时可选填
	UpstreamResourceName *string `json:"UpstreamResourceName,omitnil,omitempty" name:"UpstreamResourceName"`

	// 路径重写
	PathRewrite *HTTPServicePathRewrite `json:"PathRewrite,omitnil,omitempty" name:"PathRewrite"`

	// 是否开启安全域名。默认开启
	EnableSafeDomain *bool `json:"EnableSafeDomain,omitnil,omitempty" name:"EnableSafeDomain"`

	// 是否开启身份认证。默认关闭
	EnableAuth *bool `json:"EnableAuth,omitnil,omitempty" name:"EnableAuth"`

	// 是否开启路径透传。默认关闭
	EnablePathTransmission *bool `json:"EnablePathTransmission,omitnil,omitempty" name:"EnablePathTransmission"`

	// QPS限频策略
	QPSPolicy *HTTPServiceRouteQPSPolicy `json:"QPSPolicy,omitnil,omitempty" name:"QPSPolicy"`

	// 是否开启路由
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 扩展字段，内部包含headers处理等
	Extension *HTTPServiceExtension `json:"Extension,omitnil,omitempty" name:"Extension"`
}

type HTTPServiceRouteQPSPolicy struct {
	// QPS值，每秒请求次数
	QPSTotal *int64 `json:"QPSTotal,omitnil,omitempty" name:"QPSTotal"`

	// 客户端限频配置
	QPSPerClient *HTTPServiceQPSPerClient `json:"QPSPerClient,omitnil,omitempty" name:"QPSPerClient"`
}

type HpaPolicy struct {
	// 策略类型
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 策略阈值
	PolicyThreshold *int64 `json:"PolicyThreshold,omitnil,omitempty" name:"PolicyThreshold"`
}

type IndexAccesses struct {
	// 索引命中次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ops *int64 `json:"Ops,omitnil,omitempty" name:"Ops"`

	// 命中次数从何时开始计数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Since *string `json:"Since,omitnil,omitempty" name:"Since"`
}

type IndexInfo struct {
	// 索引名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 索引大小，单位: 字节
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 索引键值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keys []*Indexkey `json:"Keys,omitnil,omitempty" name:"Keys"`

	// 索引使用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Accesses *IndexAccesses `json:"Accesses,omitnil,omitempty" name:"Accesses"`

	// 是否为唯一索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unique *bool `json:"Unique,omitnil,omitempty" name:"Unique"`
}

type Indexkey struct {
	// 键名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 方向：specify 1 for ascending or -1 for descending
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`
}

// Predefined struct for user
type InquireVmPriceRequestParams struct {
	// 服务器类型：
	// LightHouse = 轻量云服务器
	// CVM = 云服务器
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 轻量云服务器套餐ID。
	// 当Type=LightHouse时必传
	LightHouseBundleId *string `json:"LightHouseBundleId,omitnil,omitempty" name:"LightHouseBundleId"`

	// 轻量云服务器镜像ID。当Type=LightHouse时必传
	LightHouseBlueprintId *string `json:"LightHouseBlueprintId,omitnil,omitempty" name:"LightHouseBlueprintId"`
}

type InquireVmPriceRequest struct {
	*tchttp.BaseRequest
	
	// 服务器类型：
	// LightHouse = 轻量云服务器
	// CVM = 云服务器
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 轻量云服务器套餐ID。
	// 当Type=LightHouse时必传
	LightHouseBundleId *string `json:"LightHouseBundleId,omitnil,omitempty" name:"LightHouseBundleId"`

	// 轻量云服务器镜像ID。当Type=LightHouse时必传
	LightHouseBlueprintId *string `json:"LightHouseBlueprintId,omitnil,omitempty" name:"LightHouseBlueprintId"`
}

func (r *InquireVmPriceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquireVmPriceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "LightHouseBundleId")
	delete(f, "LightHouseBlueprintId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquireVmPriceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquireVmPriceResponseParams struct {
	// 价格货币单位。取值范围CNY:人民币。USD:美元。
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 原价（主机原始每月价格）
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 折扣率
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 折扣后每月价格
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// 折扣前每天资源点
	OriginalCredits *float64 `json:"OriginalCredits,omitnil,omitempty" name:"OriginalCredits"`

	// 折扣后每天资源点
	DiscountCredits *float64 `json:"DiscountCredits,omitnil,omitempty" name:"DiscountCredits"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquireVmPriceResponse struct {
	*tchttp.BaseResponse
	Response *InquireVmPriceResponseParams `json:"Response"`
}

func (r *InquireVmPriceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquireVmPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KVPair struct {
	// 键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type ListPGUserMigrationsRequestParams struct {
	// <p>云开发环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>查询条数</p><p>取值范围：[1, 500]</p><p>默认值：100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>分页偏移</p><p>默认值：0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type ListPGUserMigrationsRequest struct {
	*tchttp.BaseRequest
	
	// <p>云开发环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>查询条数</p><p>取值范围：[1, 500]</p><p>默认值：100</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>分页偏移</p><p>默认值：0</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *ListPGUserMigrationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPGUserMigrationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListPGUserMigrationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListPGUserMigrationsResponseParams struct {
	// <p>总数量</p>
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// <p>已应用最新版本号</p><p>参数格式：纯数字，14位时间格式</p>
	LatestVersion *string `json:"LatestVersion,omitnil,omitempty" name:"LatestVersion"`

	// <p>已应用migration列表</p>
	Migrations []*MigrationSummary `json:"Migrations,omitnil,omitempty" name:"Migrations"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListPGUserMigrationsResponse struct {
	*tchttp.BaseResponse
	Response *ListPGUserMigrationsResponseParams `json:"Response"`
}

func (r *ListPGUserMigrationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListPGUserMigrationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTablesRequestParams struct {
	// 每页返回数量（0-1000)
	MgoLimit *int64 `json:"MgoLimit,omitnil,omitempty" name:"MgoLimit"`

	// FlexDB实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 分页偏移量
	MgoOffset *int64 `json:"MgoOffset,omitnil,omitempty" name:"MgoOffset"`

	// 过滤标签数组，用于过滤表名，可选值如：HIDDEN、WEDA、WEDA_SYSTEM
	Filters []*string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 模糊搜索查询值
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 是否展示隐藏表
	ShowHidden *bool `json:"ShowHidden,omitnil,omitempty" name:"ShowHidden"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// mongo连接器信息
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

type ListTablesRequest struct {
	*tchttp.BaseRequest
	
	// 每页返回数量（0-1000)
	MgoLimit *int64 `json:"MgoLimit,omitnil,omitempty" name:"MgoLimit"`

	// FlexDB实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 分页偏移量
	MgoOffset *int64 `json:"MgoOffset,omitnil,omitempty" name:"MgoOffset"`

	// 过滤标签数组，用于过滤表名，可选值如：HIDDEN、WEDA、WEDA_SYSTEM
	Filters []*string `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 模糊搜索查询值
	SearchValue *string `json:"SearchValue,omitnil,omitempty" name:"SearchValue"`

	// 是否展示隐藏表
	ShowHidden *bool `json:"ShowHidden,omitnil,omitempty" name:"ShowHidden"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// mongo连接器信息
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

func (r *ListTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MgoLimit")
	delete(f, "Tag")
	delete(f, "MgoOffset")
	delete(f, "Filters")
	delete(f, "SearchValue")
	delete(f, "ShowHidden")
	delete(f, "EnvId")
	delete(f, "MongoConnector")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTablesResponseParams struct {
	// 表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tables []*TableInfo `json:"Tables,omitnil,omitempty" name:"Tables"`

	// 分页信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pager *Pager `json:"Pager,omitnil,omitempty" name:"Pager"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTablesResponse struct {
	*tchttp.BaseResponse
	Response *ListTablesResponseParams `json:"Response"`
}

func (r *ListTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LocalizedMessage struct {
	// 默认展示的文本
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 针对每种语言展示的文字
	Localized []*MessageLocalized `json:"Localized,omitnil,omitempty" name:"Localized"`
}

type LocalizedTemplate struct {
	// <p>中文</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZhCN *string `json:"ZhCN,omitnil,omitempty" name:"ZhCN"`

	// <p>英文</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnUS *string `json:"EnUS,omitnil,omitempty" name:"EnUS"`
}

type LogObject struct {
	// 日志属于的 topic ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志主题的名字
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 日志时间
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 日志内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 采集路径
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 日志来源设备
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`
}

type LogResObject struct {
	// 获取更多检索结果的游标
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 搜索结果是否已经全部返回
	ListOver *bool `json:"ListOver,omitnil,omitempty" name:"ListOver"`

	// 日志内容信息
	Results []*LogObject `json:"Results,omitnil,omitempty" name:"Results"`

	// 日志聚合结果
	AnalysisRecords []*string `json:"AnalysisRecords,omitnil,omitempty" name:"AnalysisRecords"`
}

type LogServiceInfo struct {
	// log名
	LogsetName *string `json:"LogsetName,omitnil,omitempty" name:"LogsetName"`

	// log-id
	LogsetId *string `json:"LogsetId,omitnil,omitempty" name:"LogsetId"`

	// topic名
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// topic-id
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// cls日志所属地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// topic保存时长 默认7天
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`
}

type MFALoginConfig struct {
	// MFA 多因子认证开关。取值范围：
	// TRUE：开启 MFA 多因子认证
	// FALSE：关闭 MFA 多因子认证
	// 不传则不修改当前配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	On *string `json:"On,omitnil,omitempty" name:"On"`

	// 短信验证开关，控制是否在 MFA 流程中启用短信验证码校验。取值范围：
	// TRUE：开启短信验证
	// FALSE：关闭短信验证
	// 不传则不修改当前配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sms *string `json:"Sms,omitnil,omitempty" name:"Sms"`

	// 邮箱验证开关，控制是否在 MFA 流程中启用邮箱验证码校验。取值范围：
	// TRUE：开启邮箱验证
	// FALSE：关闭邮箱验证
	// 不传则不修改当前配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 强制绑定手机号开关，控制用户在完成 MFA 认证前是否必须绑定手机号。取值范围：
	// TRUE：要求绑定手机号
	// FALSE：不要求绑定手机号
	// 不传则不修改当前配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequiredBindPhone *string `json:"RequiredBindPhone,omitnil,omitempty" name:"RequiredBindPhone"`
}

type ManagedAIModel struct {
	// <p>模型名</p>
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>是否开启MCP</p>
	EnableMCP *bool `json:"EnableMCP,omitnil,omitempty" name:"EnableMCP"`

	// <p>模型规格</p>
	ModelSpec *ManagedAIModelSpec `json:"ModelSpec,omitnil,omitempty" name:"ModelSpec"`

	// <p>模型计费信息</p>
	ModelChargingInfo []*ManagedAIModelChargingInfo `json:"ModelChargingInfo,omitnil,omitempty" name:"ModelChargingInfo"`
}

type ManagedAIModelChargingInfo struct {
	// <p>计费类型</p><p>枚举值：</p><ul><li>Uniform： 固定计费</li><li>Tiered： 分段计费</li></ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>分组名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>输入 Token 价格</p>
	InputPrice *string `json:"InputPrice,omitnil,omitempty" name:"InputPrice"`

	// <p>输出 Token 价格</p>
	OutputPrice *string `json:"OutputPrice,omitnil,omitempty" name:"OutputPrice"`

	// <p>命中缓存价格</p>
	CachePrice *string `json:"CachePrice,omitnil,omitempty" name:"CachePrice"`

	// <p>计费单位</p>
	InputOutputUnit *string `json:"InputOutputUnit,omitnil,omitempty" name:"InputOutputUnit"`
}

type ManagedAIModelGroup struct {
	// <p>模型分组</p>
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// <p>模型列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Models []*ManagedAIModel `json:"Models,omitnil,omitempty" name:"Models"`

	// <p>备注</p>
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ManagedAIModelSpec struct {
	// <p>最大输入 Token</p>
	MaxInputToken *string `json:"MaxInputToken,omitnil,omitempty" name:"MaxInputToken"`

	// <p>最大输出 Token</p>
	MaxOutputToken *string `json:"MaxOutputToken,omitnil,omitempty" name:"MaxOutputToken"`

	// <p>上下文长度</p>
	ContextLength *string `json:"ContextLength,omitnil,omitempty" name:"ContextLength"`
}

type MessageLocalized struct {
	// 字符串
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 在该语言中
	Locale *string `json:"Locale,omitnil,omitempty" name:"Locale"`
}

type MgoCommandParam struct {
	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 操作类型，可选类型为：UPDATE/QUERY/INSERT/DELETE/COMMAND，本操作必须按实际填写，监控会依赖该字段统计本次操作的类型，并实时减少用户配额，如果填写错误会误扣用户请求配额
	CommandType *string `json:"CommandType,omitnil,omitempty" name:"CommandType"`

	// 待执行命令
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`
}

type MgoIndexKeys struct {
	// 无
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 无
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type MgoKeySchema struct {
	// 索引字段
	MgoIndexKeys []*MgoIndexKeys `json:"MgoIndexKeys,omitnil,omitempty" name:"MgoIndexKeys"`

	// 是否唯一索引
	MgoIsUnique *bool `json:"MgoIsUnique,omitnil,omitempty" name:"MgoIsUnique"`

	// 是否稀疏索引
	MgoIsSparse *bool `json:"MgoIsSparse,omitnil,omitempty" name:"MgoIsSparse"`
}

type MigrationConflict struct {
	// <p>migration 版本号</p><p>参数格式：纯数字，14位时间格式</p>
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// <p>migration 版本名</p><p>参数格式：仅允许小写字母和下划线</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>数据库已应用migration的版本名</p><p>参数格式：仅允许小写字母和下划线</p>
	RemoteName *string `json:"RemoteName,omitnil,omitempty" name:"RemoteName"`

	// <p>本次sql计算出来的checksum</p>
	LocalChecksum *string `json:"LocalChecksum,omitnil,omitempty" name:"LocalChecksum"`

	// <p>已应用的migration，数据库存储的checksum</p>
	RemoteChecksum *string `json:"RemoteChecksum,omitnil,omitempty" name:"RemoteChecksum"`

	// <p>归入该分组的原因</p>
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// <p>冲突信息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type MigrationInput struct {
	// <p>migration 版本号</p><p>参数格式：纯数字，14位时间格式</p>
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// <p>migration 版本名</p><p>入参限制：仅允许小写字母和下划线</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>migration 应用 sql 语句</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>migration 回滚 sql 语句</p>
	Rollback *string `json:"Rollback,omitnil,omitempty" name:"Rollback"`
}

type MigrationPlanItem struct {
	// <p>migration 版本号</p><p>参数格式：纯数字，14位时间格式</p>
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// <p>migration 版本名</p><p>参数格式：仅允许小写字母和下划线</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>migration query sql checksum</p><p>服务端自动生成，同版本不同checksum会拒绝执行</p>
	Checksum *string `json:"Checksum,omitnil,omitempty" name:"Checksum"`

	// <p>状态</p><p>枚举值：</p><ul><li>applied： 已应用</li><li>pending： 待执行</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>标记请求来源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>被归入该分组的原因，比如not_applied、checksum_matched</p>
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type MigrationSummary struct {
	// <p>migration 版本号</p><p>参数格式：纯数字，14位时间格式</p>
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// <p>migration 版本名</p><p>参数格式：仅允许小写字母和下划线</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>migration query sql 语句checksum</p><p>服务端自动生成，同版本不同checksum会拒绝执行</p>
	Checksum *string `json:"Checksum,omitnil,omitempty" name:"Checksum"`

	// <p>应用时间</p>
	AppliedAt *string `json:"AppliedAt,omitnil,omitempty" name:"AppliedAt"`

	// <p>请求来源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>migration 创建时间</p>
	CreatedBy *string `json:"CreatedBy,omitnil,omitempty" name:"CreatedBy"`
}

// Predefined struct for user
type ModifyClientRequestParams struct {
	// 客户端所属的云开发环境 ID，用于标识该应用归属的云开发环境。不同环境之间的应用数据相互隔离。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 需要修改的客户端唯一标识符（Client ID），在 OAuth/OIDC 授权流程中作为 client_id 参数使用。该字段为定位参数，仅用于指定目标客户端，不可修改。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Refresh Token 的有效期，单位为秒。超过该时间后 Refresh Token 将失效，用户需重新登录。取值范围：1800~2592000（即 30 分钟至 30 天），超出上限将被截断为 2592000。默认值为 2592000（即 30 天）。注意：当该值 ≤ 7200 时，AccessTokenExpiresIn 将被自动设为该值减去 660 秒。
	RefreshTokenExpiresIn *int64 `json:"RefreshTokenExpiresIn,omitnil,omitempty" name:"RefreshTokenExpiresIn"`

	// 单个用户在该应用下允许同时登录的最大会话数量。取值范围：-1~50。特殊值说明：-1 表示不限制设备数；0 表示按客户端 User-Agent 区分设备（相同 User-Agent 视为同一设备）；1~50 为精确的最大会话数限制，超出限制后最早登录的会话将被自动踢出。不传则保持原有配置不变。
	MaxDevice *int64 `json:"MaxDevice,omitnil,omitempty" name:"MaxDevice"`

	// Access Token 的有效期，单位为秒。超过该时间后 Access Token 将失效，需使用 Refresh Token 重新换取。最小有效值为 1800 秒（小于 1800 将被忽略，使用默认值），默认值为 7200（即 2 小时）。该值应小于 RefreshTokenExpiresIn。
	AccessTokenExpiresIn *int64 `json:"AccessTokenExpiresIn,omitnil,omitempty" name:"AccessTokenExpiresIn"`
}

type ModifyClientRequest struct {
	*tchttp.BaseRequest
	
	// 客户端所属的云开发环境 ID，用于标识该应用归属的云开发环境。不同环境之间的应用数据相互隔离。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 需要修改的客户端唯一标识符（Client ID），在 OAuth/OIDC 授权流程中作为 client_id 参数使用。该字段为定位参数，仅用于指定目标客户端，不可修改。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// Refresh Token 的有效期，单位为秒。超过该时间后 Refresh Token 将失效，用户需重新登录。取值范围：1800~2592000（即 30 分钟至 30 天），超出上限将被截断为 2592000。默认值为 2592000（即 30 天）。注意：当该值 ≤ 7200 时，AccessTokenExpiresIn 将被自动设为该值减去 660 秒。
	RefreshTokenExpiresIn *int64 `json:"RefreshTokenExpiresIn,omitnil,omitempty" name:"RefreshTokenExpiresIn"`

	// 单个用户在该应用下允许同时登录的最大会话数量。取值范围：-1~50。特殊值说明：-1 表示不限制设备数；0 表示按客户端 User-Agent 区分设备（相同 User-Agent 视为同一设备）；1~50 为精确的最大会话数限制，超出限制后最早登录的会话将被自动踢出。不传则保持原有配置不变。
	MaxDevice *int64 `json:"MaxDevice,omitnil,omitempty" name:"MaxDevice"`

	// Access Token 的有效期，单位为秒。超过该时间后 Access Token 将失效，需使用 Refresh Token 重新换取。最小有效值为 1800 秒（小于 1800 将被忽略，使用默认值），默认值为 7200（即 2 小时）。该值应小于 RefreshTokenExpiresIn。
	AccessTokenExpiresIn *int64 `json:"AccessTokenExpiresIn,omitnil,omitempty" name:"AccessTokenExpiresIn"`
}

func (r *ModifyClientRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClientRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Id")
	delete(f, "RefreshTokenExpiresIn")
	delete(f, "MaxDevice")
	delete(f, "AccessTokenExpiresIn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClientRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClientResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClientResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClientResponseParams `json:"Response"`
}

func (r *ModifyClientResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClientResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClsTopicRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 日志生命周期，单位天，可取值范围1~3600，取值为3640时代表永久保存
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`
}

type ModifyClsTopicRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 日志生命周期，单位天，可取值范围1~3600，取值为3640时代表永久保存
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`
}

func (r *ModifyClsTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClsTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClsTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyClsTopicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyClsTopicResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClsTopicResponseParams `json:"Response"`
}

func (r *ModifyClsTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClsTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseACLRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitnil,omitempty" name:"CollectionName"`

	// 权限标签。包含以下取值：
	// <li> READONLY：所有用户可读，仅创建者和管理员可写</li>
	// <li> PRIVATE：仅创建者及管理员可读写</li>
	// <li> ADMINWRITE：所有用户可读，仅管理员可写</li>
	// <li> ADMINONLY：仅管理员可读写</li>
	AclTag *string `json:"AclTag,omitnil,omitempty" name:"AclTag"`
}

type ModifyDatabaseACLRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitnil,omitempty" name:"CollectionName"`

	// 权限标签。包含以下取值：
	// <li> READONLY：所有用户可读，仅创建者和管理员可写</li>
	// <li> PRIVATE：仅创建者及管理员可读写</li>
	// <li> ADMINWRITE：所有用户可读，仅管理员可写</li>
	// <li> ADMINONLY：仅管理员可读写</li>
	AclTag *string `json:"AclTag,omitnil,omitempty" name:"AclTag"`
}

func (r *ModifyDatabaseACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "CollectionName")
	delete(f, "AclTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDatabaseACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDatabaseACLResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDatabaseACLResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDatabaseACLResponseParams `json:"Response"`
}

func (r *ModifyDatabaseACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDatabaseACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnvPlanRequestParams struct {
	// 所需变更套餐的环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 目标套餐Id。
	// 对于云开发环境套餐，可通过 [DescribeBaasPackageList](https://cloud.tencent.com/document/product/876/78167) 接口获取，对应其出参的PackageName
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 是否自动选择代金券支付。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

type ModifyEnvPlanRequest struct {
	*tchttp.BaseRequest
	
	// 所需变更套餐的环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 目标套餐Id。
	// 对于云开发环境套餐，可通过 [DescribeBaasPackageList](https://cloud.tencent.com/document/product/876/78167) 接口获取，对应其出参的PackageName
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 是否自动选择代金券支付。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

func (r *ModifyEnvPlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvPlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "PackageId")
	delete(f, "AutoVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEnvPlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnvPlanResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEnvPlanResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEnvPlanResponseParams `json:"Response"`
}

func (r *ModifyEnvPlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvPlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnvRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 环境备注名，要以a-z开头，不能包含 a-zA-z0-9- 以外的字符
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`
}

type ModifyEnvRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 环境备注名，要以a-z开头，不能包含 a-zA-z0-9- 以外的字符
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`
}

func (r *ModifyEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Alias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnvResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyEnvResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEnvResponseParams `json:"Response"`
}

func (r *ModifyEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHTTPServiceRouteRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 域名路由信息
	Domain *HTTPServiceDomainParam `json:"Domain,omitnil,omitempty" name:"Domain"`
}

type ModifyHTTPServiceRouteRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 域名路由信息
	Domain *HTTPServiceDomainParam `json:"Domain,omitnil,omitempty" name:"Domain"`
}

func (r *ModifyHTTPServiceRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHTTPServiceRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHTTPServiceRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHTTPServiceRouteResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyHTTPServiceRouteResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHTTPServiceRouteResponseParams `json:"Response"`
}

func (r *ModifyHTTPServiceRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHTTPServiceRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoginConfigRequestParams struct {
	// 环境 ID，用于指定需要修改登录策略的云开发环境。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 手机号短信登录开关。设置为 true 开启手机号短信登录，允许用户使用手机号和短信验证码进行登录和注册；设置为 false 关闭手机号短信登录。
	PhoneNumberLogin *bool `json:"PhoneNumberLogin,omitnil,omitempty" name:"PhoneNumberLogin"`

	// 邮箱登录开关。设置为 true 开启邮箱登录，允许用户使用邮箱和密码进行登录和注册；设置为 false 关闭邮箱登录。
	EmailLogin *bool `json:"EmailLogin,omitnil,omitempty" name:"EmailLogin"`

	// 用户名密码登录开关。设置为 true 开启用户名密码登录，允许用户使用用户名和密码进行登录和注册；设置为 false 关闭用户名密码登录。
	UserNameLogin *bool `json:"UserNameLogin,omitnil,omitempty" name:"UserNameLogin"`

	// 匿名登录开关。设置为 true 开启匿名登录，允许用户无需注册即可以匿名身份访问应用；设置为 false 关闭匿名登录。
	AnonymousLogin *bool `json:"AnonymousLogin,omitnil,omitempty" name:"AnonymousLogin"`

	// 短信验证码发送配置，用于设置短信验证码的发送通道类型和日发送限额。不传则不修改当前配置。
	SmsVerificationConfig *VerificationConfig `json:"SmsVerificationConfig,omitnil,omitempty" name:"SmsVerificationConfig"`

	// MFA 多因子认证登录配置，用于设置多因子认证开关及验证方式（短信、邮箱、TOTP、强制绑定手机号）。不传则不修改当前配置。
	MfaConfig *MFALoginConfig `json:"MfaConfig,omitnil,omitempty" name:"MfaConfig"`

	// 密码更新策略配置，用于设置首次登录强制修改密码和定期强制修改密码策略。不传则不修改当前配置。
	PwdUpdateStrategy *PasswordUpdateLoginConfig `json:"PwdUpdateStrategy,omitnil,omitempty" name:"PwdUpdateStrategy"`
}

type ModifyLoginConfigRequest struct {
	*tchttp.BaseRequest
	
	// 环境 ID，用于指定需要修改登录策略的云开发环境。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 手机号短信登录开关。设置为 true 开启手机号短信登录，允许用户使用手机号和短信验证码进行登录和注册；设置为 false 关闭手机号短信登录。
	PhoneNumberLogin *bool `json:"PhoneNumberLogin,omitnil,omitempty" name:"PhoneNumberLogin"`

	// 邮箱登录开关。设置为 true 开启邮箱登录，允许用户使用邮箱和密码进行登录和注册；设置为 false 关闭邮箱登录。
	EmailLogin *bool `json:"EmailLogin,omitnil,omitempty" name:"EmailLogin"`

	// 用户名密码登录开关。设置为 true 开启用户名密码登录，允许用户使用用户名和密码进行登录和注册；设置为 false 关闭用户名密码登录。
	UserNameLogin *bool `json:"UserNameLogin,omitnil,omitempty" name:"UserNameLogin"`

	// 匿名登录开关。设置为 true 开启匿名登录，允许用户无需注册即可以匿名身份访问应用；设置为 false 关闭匿名登录。
	AnonymousLogin *bool `json:"AnonymousLogin,omitnil,omitempty" name:"AnonymousLogin"`

	// 短信验证码发送配置，用于设置短信验证码的发送通道类型和日发送限额。不传则不修改当前配置。
	SmsVerificationConfig *VerificationConfig `json:"SmsVerificationConfig,omitnil,omitempty" name:"SmsVerificationConfig"`

	// MFA 多因子认证登录配置，用于设置多因子认证开关及验证方式（短信、邮箱、TOTP、强制绑定手机号）。不传则不修改当前配置。
	MfaConfig *MFALoginConfig `json:"MfaConfig,omitnil,omitempty" name:"MfaConfig"`

	// 密码更新策略配置，用于设置首次登录强制修改密码和定期强制修改密码策略。不传则不修改当前配置。
	PwdUpdateStrategy *PasswordUpdateLoginConfig `json:"PwdUpdateStrategy,omitnil,omitempty" name:"PwdUpdateStrategy"`
}

func (r *ModifyLoginConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoginConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "PhoneNumberLogin")
	delete(f, "EmailLogin")
	delete(f, "UserNameLogin")
	delete(f, "AnonymousLogin")
	delete(f, "SmsVerificationConfig")
	delete(f, "MfaConfig")
	delete(f, "PwdUpdateStrategy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoginConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoginConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLoginConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoginConfigResponseParams `json:"Response"`
}

func (r *ModifyLoginConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoginConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProviderRequestParams struct {
	// 云开发环境 ID，用于唯一标识当前操作所属的云开发环境。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 身份源的唯一标识符，用于指定需要修改的目标身份源。格式要求：2~32 位，仅支持小写英文字母和数字，不可包含空格或特殊字符。例如：github、google。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 身份源的显示名称，支持国际化多语言配置。用户在登录页面看到的身份源名称将使用该字段，建议根据实际业务场景填写易于识别的名称，例如：GitHub、Google 等。
	Name *LocalizedMessage `json:"Name,omitnil,omitempty" name:"Name"`

	// 身份源图标的访问地址，将展示在登录页的身份源按钮上。建议使用 64×64 像素的 SVG 格式图片以保证清晰度，支持 HTTP/HTTPS 公网可访问的图片链接。
	Picture *string `json:"Picture,omitnil,omitempty" name:"Picture"`

	// 身份源对应的官方主页地址。该信息将在用户查看自己的第三方账号绑定列表时展示，帮助用户识别已绑定的身份源来源。例如 GitHub 身份源可填写：https://github.com。
	Homepage *string `json:"Homepage,omitnil,omitempty" name:"Homepage"`

	// 身份源协议类型，决定该身份源使用何种认证协议与第三方平台对接。可选值：
	// OAUTH：标准 OAuth 2.0 协议
	// OIDC：OpenID Connect 协议
	// SAML：SAML 2.0 协议
	// WX_MICRO_APP：微信小程序登录
	// WX_QRCODE_MICRO_APP：微信小程序扫码登录
	// WX_CLOUDBASE_MICRO_APP：云开发托管小程序登录
	// WX_MP：微信公众号网页授权登录
	// WX_OPEN：微信开放平台扫码登录
	// WX_WORK_INTERNAL：企业微信自建应用登录
	// WX_WORK_AGENT：企业微信代开发应用登录
	// WX_WORK_THIRD_PARTY：企业微信第三方应用登录
	// WX_WORK_THIRD_PARTY_ASSOCIATION：企业微信第三方应用关联登录
	// CUSTOM：自定义登录
	// EMAIL：邮箱登录
	ProviderType *string `json:"ProviderType,omitnil,omitempty" name:"ProviderType"`

	// 身份认证源协议连接配置，包含与第三方平台对接所需的核心参数，如 ClientId、ClientSecret、授权端点、Token 端点、回调地址、Scope、SAML Metadata、请求和响应参数映射等。不同 ProviderType 对应不同的配置项要求。注意：CUSTOM 和 EMAIL 类型的身份源，其存储后端类型（StorageDb）不可修改。
	Config *ProviderConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 是否开启透传登录模式。可选值：TRUE（开启）、FALSE（关闭，默认值）。开启后，平台不会持久化存储用户数据，仅将第三方身份源返回的用户信息透传给业务方，适用于不希望平台留存用户数据的场景。注意：开启透传模式时，ReuseUserId 将被强制设为 TRUE，AutoSignUpWithProviderUser 将被强制设为 FALSE。
	TransparentMode *string `json:"TransparentMode,omitnil,omitempty" name:"TransparentMode"`

	// 身份源的启用状态。可选值：TRUE（启用，用户可通过该身份源登录）、FALSE（禁用，登录入口将被隐藏，已有绑定关系不受影响）、UNSPECIFIED（默认为 TRUE）。
	On *string `json:"On,omitnil,omitempty" name:"On"`

	// 身份源的详细描述信息，支持国际化多语言配置。可用于向用户说明该身份源的用途或使用场景，例如：谷歌授权登录。
	Description *LocalizedMessage `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否直接复用第三方身份源的用户 ID 作为平台用户 ID。可选值：TRUE（开启，返回的用户 ID 将直接使用第三方身份源的用户 ID，适用于已有用户体系迁移场景）、FALSE（关闭，由平台生成独立用户 ID）、UNSPECIFIED（默认为 FALSE，但当 TransparentMode 为 TRUE 时将被强制设为 TRUE）。注意：开启后需确保第三方用户 ID 的全局唯一性，避免 ID 冲突。
	ReuseUserId *string `json:"ReuseUserId,omitnil,omitempty" name:"ReuseUserId"`

	// 邮箱身份源的专项配置，包含邮件服务商、发件人地址、SMTP 配置等参数，用于支持通过邮箱验证码或邮箱密码方式进行身份认证。仅当身份源 ID 为 email 时有效。若该身份源不存在，系统将自动创建一个默认关闭的邮箱身份源。
	EmailConfig *EmailProviderConfig `json:"EmailConfig,omitnil,omitempty" name:"EmailConfig"`

	// 是否开启邮箱自动关联登录。可选值：TRUE（开启）、FALSE（关闭）、UNSPECIFIED（默认为 FALSE）。开启后，若第三方身份源返回的邮箱与系统中已有用户的邮箱一致，则自动将该第三方账号与已有用户关联并完成登录，无需用户手动绑定。
	AutoSignInWhenEmailMatch *string `json:"AutoSignInWhenEmailMatch,omitnil,omitempty" name:"AutoSignInWhenEmailMatch"`

	// 是否开启手机号自动关联登录。可选值：TRUE（开启）、FALSE（关闭）、UNSPECIFIED（默认行为等同 TRUE）。开启后，若第三方身份源返回的手机号与系统中已有用户的手机号一致，则自动将该第三方账号与已有用户关联并完成登录，无需用户手动绑定。注意：该字段默认行为（UNSPECIFIED）与 AutoSignInWhenEmailMatch 不同，手机号匹配在未显式关闭时默认生效。
	AutoSignInWhenPhoneNumberMatch *string `json:"AutoSignInWhenPhoneNumberMatch,omitnil,omitempty" name:"AutoSignInWhenPhoneNumberMatch"`
}

type ModifyProviderRequest struct {
	*tchttp.BaseRequest
	
	// 云开发环境 ID，用于唯一标识当前操作所属的云开发环境。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 身份源的唯一标识符，用于指定需要修改的目标身份源。格式要求：2~32 位，仅支持小写英文字母和数字，不可包含空格或特殊字符。例如：github、google。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 身份源的显示名称，支持国际化多语言配置。用户在登录页面看到的身份源名称将使用该字段，建议根据实际业务场景填写易于识别的名称，例如：GitHub、Google 等。
	Name *LocalizedMessage `json:"Name,omitnil,omitempty" name:"Name"`

	// 身份源图标的访问地址，将展示在登录页的身份源按钮上。建议使用 64×64 像素的 SVG 格式图片以保证清晰度，支持 HTTP/HTTPS 公网可访问的图片链接。
	Picture *string `json:"Picture,omitnil,omitempty" name:"Picture"`

	// 身份源对应的官方主页地址。该信息将在用户查看自己的第三方账号绑定列表时展示，帮助用户识别已绑定的身份源来源。例如 GitHub 身份源可填写：https://github.com。
	Homepage *string `json:"Homepage,omitnil,omitempty" name:"Homepage"`

	// 身份源协议类型，决定该身份源使用何种认证协议与第三方平台对接。可选值：
	// OAUTH：标准 OAuth 2.0 协议
	// OIDC：OpenID Connect 协议
	// SAML：SAML 2.0 协议
	// WX_MICRO_APP：微信小程序登录
	// WX_QRCODE_MICRO_APP：微信小程序扫码登录
	// WX_CLOUDBASE_MICRO_APP：云开发托管小程序登录
	// WX_MP：微信公众号网页授权登录
	// WX_OPEN：微信开放平台扫码登录
	// WX_WORK_INTERNAL：企业微信自建应用登录
	// WX_WORK_AGENT：企业微信代开发应用登录
	// WX_WORK_THIRD_PARTY：企业微信第三方应用登录
	// WX_WORK_THIRD_PARTY_ASSOCIATION：企业微信第三方应用关联登录
	// CUSTOM：自定义登录
	// EMAIL：邮箱登录
	ProviderType *string `json:"ProviderType,omitnil,omitempty" name:"ProviderType"`

	// 身份认证源协议连接配置，包含与第三方平台对接所需的核心参数，如 ClientId、ClientSecret、授权端点、Token 端点、回调地址、Scope、SAML Metadata、请求和响应参数映射等。不同 ProviderType 对应不同的配置项要求。注意：CUSTOM 和 EMAIL 类型的身份源，其存储后端类型（StorageDb）不可修改。
	Config *ProviderConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 是否开启透传登录模式。可选值：TRUE（开启）、FALSE（关闭，默认值）。开启后，平台不会持久化存储用户数据，仅将第三方身份源返回的用户信息透传给业务方，适用于不希望平台留存用户数据的场景。注意：开启透传模式时，ReuseUserId 将被强制设为 TRUE，AutoSignUpWithProviderUser 将被强制设为 FALSE。
	TransparentMode *string `json:"TransparentMode,omitnil,omitempty" name:"TransparentMode"`

	// 身份源的启用状态。可选值：TRUE（启用，用户可通过该身份源登录）、FALSE（禁用，登录入口将被隐藏，已有绑定关系不受影响）、UNSPECIFIED（默认为 TRUE）。
	On *string `json:"On,omitnil,omitempty" name:"On"`

	// 身份源的详细描述信息，支持国际化多语言配置。可用于向用户说明该身份源的用途或使用场景，例如：谷歌授权登录。
	Description *LocalizedMessage `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否直接复用第三方身份源的用户 ID 作为平台用户 ID。可选值：TRUE（开启，返回的用户 ID 将直接使用第三方身份源的用户 ID，适用于已有用户体系迁移场景）、FALSE（关闭，由平台生成独立用户 ID）、UNSPECIFIED（默认为 FALSE，但当 TransparentMode 为 TRUE 时将被强制设为 TRUE）。注意：开启后需确保第三方用户 ID 的全局唯一性，避免 ID 冲突。
	ReuseUserId *string `json:"ReuseUserId,omitnil,omitempty" name:"ReuseUserId"`

	// 邮箱身份源的专项配置，包含邮件服务商、发件人地址、SMTP 配置等参数，用于支持通过邮箱验证码或邮箱密码方式进行身份认证。仅当身份源 ID 为 email 时有效。若该身份源不存在，系统将自动创建一个默认关闭的邮箱身份源。
	EmailConfig *EmailProviderConfig `json:"EmailConfig,omitnil,omitempty" name:"EmailConfig"`

	// 是否开启邮箱自动关联登录。可选值：TRUE（开启）、FALSE（关闭）、UNSPECIFIED（默认为 FALSE）。开启后，若第三方身份源返回的邮箱与系统中已有用户的邮箱一致，则自动将该第三方账号与已有用户关联并完成登录，无需用户手动绑定。
	AutoSignInWhenEmailMatch *string `json:"AutoSignInWhenEmailMatch,omitnil,omitempty" name:"AutoSignInWhenEmailMatch"`

	// 是否开启手机号自动关联登录。可选值：TRUE（开启）、FALSE（关闭）、UNSPECIFIED（默认行为等同 TRUE）。开启后，若第三方身份源返回的手机号与系统中已有用户的手机号一致，则自动将该第三方账号与已有用户关联并完成登录，无需用户手动绑定。注意：该字段默认行为（UNSPECIFIED）与 AutoSignInWhenEmailMatch 不同，手机号匹配在未显式关闭时默认生效。
	AutoSignInWhenPhoneNumberMatch *string `json:"AutoSignInWhenPhoneNumberMatch,omitnil,omitempty" name:"AutoSignInWhenPhoneNumberMatch"`
}

func (r *ModifyProviderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProviderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "Picture")
	delete(f, "Homepage")
	delete(f, "ProviderType")
	delete(f, "Config")
	delete(f, "TransparentMode")
	delete(f, "On")
	delete(f, "Description")
	delete(f, "ReuseUserId")
	delete(f, "EmailConfig")
	delete(f, "AutoSignInWhenEmailMatch")
	delete(f, "AutoSignInWhenPhoneNumberMatch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyProviderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyProviderResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyProviderResponse struct {
	*tchttp.BaseResponse
	Response *ModifyProviderResponseParams `json:"Response"`
}

func (r *ModifyProviderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyProviderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePermissionRequestParams struct {
	// 环境 ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 资源类型：`function`-云函数、`storage`-云存储、`table`-SQL型数据库表、`collection`-文档型数据库表。
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 权限级别。可选值：- SQL型数据库表：`READONLY`-读取全部数据，修改本人数据；`PRIVATE`-读取和修改本人数据；`ADMINWRITE`-读取全部数据，不可修改数据；`ADMINONLY`-无权限 。- 文档型数据库表：`READONLY`-读取全部数据，修改本人数据；`PRIVATE`-读取和修改本人数据；`ADMINWRITE`-读取全部数据，不可修改数据；`ADMINONLY`-无权限；`CUSTOM`-自定义安全规则 。- 云函数：`CUSTOM`-自定义安全规则 。- 云存储（权限标签）：`READONLY`-所有用户可读，仅创建者和管理员可写；`PRIVATE`-仅创建者及管理员可读写；`ADMINWRITE`-所有用户可读，仅管理员可写；`ADMINONLY`-仅管理员可读写；`CUSTOM`-自定义安全规则。
	Permission *string `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 资源标识。云函数可不传、云存储传存储桶名、数据库表传表名。
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 自定义安全规则配置，当Permission为 `CUSTOM`时必传。JSON字符串格式的规则表达式。配置参考：[云函数安全规则](https://docs.cloudbase.net/cloud-function/security-rules)、[云存储安全规则](https://docs.cloudbase.net/storage/security-rules)、[文档型数据库安全规则](https://docs.cloudbase.net/database/security-rules)。
	SecurityRule *string `json:"SecurityRule,omitnil,omitempty" name:"SecurityRule"`
}

type ModifyResourcePermissionRequest struct {
	*tchttp.BaseRequest
	
	// 环境 ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 资源类型：`function`-云函数、`storage`-云存储、`table`-SQL型数据库表、`collection`-文档型数据库表。
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 权限级别。可选值：- SQL型数据库表：`READONLY`-读取全部数据，修改本人数据；`PRIVATE`-读取和修改本人数据；`ADMINWRITE`-读取全部数据，不可修改数据；`ADMINONLY`-无权限 。- 文档型数据库表：`READONLY`-读取全部数据，修改本人数据；`PRIVATE`-读取和修改本人数据；`ADMINWRITE`-读取全部数据，不可修改数据；`ADMINONLY`-无权限；`CUSTOM`-自定义安全规则 。- 云函数：`CUSTOM`-自定义安全规则 。- 云存储（权限标签）：`READONLY`-所有用户可读，仅创建者和管理员可写；`PRIVATE`-仅创建者及管理员可读写；`ADMINWRITE`-所有用户可读，仅管理员可写；`ADMINONLY`-仅管理员可读写；`CUSTOM`-自定义安全规则。
	Permission *string `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 资源标识。云函数可不传、云存储传存储桶名、数据库表传表名。
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 自定义安全规则配置，当Permission为 `CUSTOM`时必传。JSON字符串格式的规则表达式。配置参考：[云函数安全规则](https://docs.cloudbase.net/cloud-function/security-rules)、[云存储安全规则](https://docs.cloudbase.net/storage/security-rules)、[文档型数据库安全规则](https://docs.cloudbase.net/database/security-rules)。
	SecurityRule *string `json:"SecurityRule,omitnil,omitempty" name:"SecurityRule"`
}

func (r *ModifyResourcePermissionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePermissionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "ResourceType")
	delete(f, "Permission")
	delete(f, "Resource")
	delete(f, "SecurityRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyResourcePermissionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyResourcePermissionResponseParams struct {
	// 修改结果
	Data *ModifyResourcePermissionResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyResourcePermissionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyResourcePermissionResponseParams `json:"Response"`
}

func (r *ModifyResourcePermissionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyResourcePermissionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourcePermissionResult struct {
	// 是否成功
	Success *bool `json:"Success,omitnil,omitempty" name:"Success"`
}

// Predefined struct for user
type ModifySafeRuleRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitnil,omitempty" name:"CollectionName"`

	// 权限标签。包含以下取值：
	// <li> READONLY：所有用户可读，仅创建者和管理员可写</li>
	// <li> PRIVATE：仅创建者及管理员可读写</li>
	// <li> ADMINWRITE：所有用户可读，仅管理员可写</li>
	// <li> ADMINONLY：仅管理员可读写</li>
	// <li> CUSTOM：自定义安全规则</li>
	AclTag *string `json:"AclTag,omitnil,omitempty" name:"AclTag"`

	// 安全规则内容。
	// 当 AclTag=CUSTOM 时，此参数必填。
	// 详情参考：[文档型数据库安全规则](https://docs.cloudbase.net/database/security-rules)
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`
}

type ModifySafeRuleRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitnil,omitempty" name:"CollectionName"`

	// 权限标签。包含以下取值：
	// <li> READONLY：所有用户可读，仅创建者和管理员可写</li>
	// <li> PRIVATE：仅创建者及管理员可读写</li>
	// <li> ADMINWRITE：所有用户可读，仅管理员可写</li>
	// <li> ADMINONLY：仅管理员可读写</li>
	// <li> CUSTOM：自定义安全规则</li>
	AclTag *string `json:"AclTag,omitnil,omitempty" name:"AclTag"`

	// 安全规则内容。
	// 当 AclTag=CUSTOM 时，此参数必填。
	// 详情参考：[文档型数据库安全规则](https://docs.cloudbase.net/database/security-rules)
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`
}

func (r *ModifySafeRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySafeRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "CollectionName")
	delete(f, "AclTag")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySafeRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySafeRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySafeRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifySafeRuleResponseParams `json:"Response"`
}

func (r *ModifySafeRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySafeRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStorageSourceRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 存储源
	StorageConfig *ExternalStorage `json:"StorageConfig,omitnil,omitempty" name:"StorageConfig"`
}

type ModifyStorageSourceRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 存储源
	StorageConfig *ExternalStorage `json:"StorageConfig,omitnil,omitempty" name:"StorageConfig"`
}

func (r *ModifyStorageSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStorageSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "StorageConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyStorageSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyStorageSourceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyStorageSourceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyStorageSourceResponseParams `json:"Response"`
}

func (r *ModifyStorageSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyStorageSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserRequestParams struct {
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 用户Id, 不做修改
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 用户名，用户名规则：1. 长度1-64字符 2. 只能包含大小写英文字母、数字和符号 . _ - 3. 只能以字母或数字开头 4. 不能重复，不传该字段或传空字符不修改
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户类型：internalUser-内部用户、externalUser-外部用户，不传该字段或传空字符串不修改。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 密码，传入Uid时密码可不传。密码规则：1. 长度8-32字符（推荐12位以上） 2. 不能以特殊字符开头 3. 至少包含以下四项中的三项：小写字母a-z、大写字母A-Z、数字0-9、特殊字符()!@#$%^&*\|?><_-，不传该字段或传空字符串不修改
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 用户状态：ACTIVE（激活）、BLOCKED（冻结），默认冻结，不传该字段或传空字符串不修改
	UserStatus *string `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 用户昵称，长度2-64字符，不传该字段不修改，传空字符修改为空
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 手机号，11位数字，不传该字段不修改，传空字符串修改为空
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 邮箱地址，不传该字段不修改，传空字符修改为空
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 头像链接，可访问的公网URL，不传该字段不修改，传空字符串修改为空
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`

	// 用户描述，最多200字符，不传该字段不修改，传空字符修改为空
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyUserRequest struct {
	*tchttp.BaseRequest
	
	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 用户Id, 不做修改
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 用户名，用户名规则：1. 长度1-64字符 2. 只能包含大小写英文字母、数字和符号 . _ - 3. 只能以字母或数字开头 4. 不能重复，不传该字段或传空字符不修改
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户类型：internalUser-内部用户、externalUser-外部用户，不传该字段或传空字符串不修改。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 密码，传入Uid时密码可不传。密码规则：1. 长度8-32字符（推荐12位以上） 2. 不能以特殊字符开头 3. 至少包含以下四项中的三项：小写字母a-z、大写字母A-Z、数字0-9、特殊字符()!@#$%^&*\|?><_-，不传该字段或传空字符串不修改
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 用户状态：ACTIVE（激活）、BLOCKED（冻结），默认冻结，不传该字段或传空字符串不修改
	UserStatus *string `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 用户昵称，长度2-64字符，不传该字段不修改，传空字符修改为空
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 手机号，11位数字，不传该字段不修改，传空字符串修改为空
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 邮箱地址，不传该字段不修改，传空字符修改为空
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 头像链接，可访问的公网URL，不传该字段不修改，传空字符串修改为空
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`

	// 用户描述，最多200字符，不传该字段不修改，传空字符修改为空
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Uid")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "Password")
	delete(f, "UserStatus")
	delete(f, "NickName")
	delete(f, "Phone")
	delete(f, "Email")
	delete(f, "AvatarUrl")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserResp struct {
	// 是否成功
	Success *bool `json:"Success,omitnil,omitempty" name:"Success"`
}

// Predefined struct for user
type ModifyUserResponseParams struct {
	// 修改用户返回值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *ModifyUserResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserResponseParams `json:"Response"`
}

func (r *ModifyUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MongoConnector struct {
	// 连接器实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// MongoDB数据库名
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`
}

type MySQLClusterDetail struct {
	// 集群ID
	DbClusterId *string `json:"DbClusterId,omitnil,omitempty" name:"DbClusterId"`

	// 网络详情
	NetInfo *MySQLNetDetail `json:"NetInfo,omitnil,omitempty" name:"NetInfo"`

	// 数据库详情
	DbInfo *ClusterDetail `json:"DbInfo,omitnil,omitempty" name:"DbInfo"`
}

type MySQLNetDetail struct {
	// 内网地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateNetAddress *string `json:"PrivateNetAddress,omitnil,omitempty" name:"PrivateNetAddress"`

	// 外网地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PubNetAddress *string `json:"PubNetAddress,omitnil,omitempty" name:"PubNetAddress"`

	// 网络信息（VPCID/SubnetID）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Net *string `json:"Net,omitnil,omitempty" name:"Net"`

	// 是否开通公网
	PubNetAccessEnabled *bool `json:"PubNetAccessEnabled,omitnil,omitempty" name:"PubNetAccessEnabled"`

	// vpc id 
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vpc name
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 子网名
	SubnetName *string `json:"SubnetName,omitnil,omitempty" name:"SubnetName"`
}

type MySQLTaskStatus struct {
	// SUCCESS | FAILED | PENDING
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`
}

type OrderInfo struct {
	// 订单号
	TranId *string `json:"TranId,omitnil,omitempty" name:"TranId"`

	// 订单要切换的套餐ID
	PackageId *string `json:"PackageId,omitnil,omitempty" name:"PackageId"`

	// 订单类型
	// <li>1 购买</li>
	// <li>2 续费</li>
	// <li>3 变配</li>
	TranType *string `json:"TranType,omitnil,omitempty" name:"TranType"`

	// 订单状态。
	// <li>1未支付</li>
	// <li>2 支付中</li>
	// <li>3 发货中</li>
	// <li>4 发货成功</li>
	// <li>5 发货失败</li>
	// <li>6 已退款</li>
	// <li>7 已取消</li>
	// <li>100 已删除</li>
	TranStatus *string `json:"TranStatus,omitnil,omitempty" name:"TranStatus"`

	// 订单更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 订单创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 付费模式.
	// <li>prepayment 预付费</li>
	// <li>postpaid 后付费</li>
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 订单绑定的扩展ID
	ExtensionId *string `json:"ExtensionId,omitnil,omitempty" name:"ExtensionId"`

	// 资源初始化结果(仅当ExtensionId不为空时有效): successful(初始化成功), failed(初始化失败), doing(初始化进行中), init(准备初始化)
	ResourceReady *string `json:"ResourceReady,omitnil,omitempty" name:"ResourceReady"`

	// 安装标记。建议使用方统一转大小写之后再判断。
	// <li>QuickStart：快速启动来源</li>
	// <li>Activity：活动来源</li>
	Flag *string `json:"Flag,omitnil,omitempty" name:"Flag"`

	// 下单时的参数
	ReqBody *string `json:"ReqBody,omitnil,omitempty" name:"ReqBody"`
}

type Pager struct {
	// 分页偏移量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 每页返回记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 文档集合总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type PasswordUpdateLoginConfig struct {
	// 首次登录强制修改密码开关。开启后，用户首次登录时将强制要求修改密码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstLoginUpdate *bool `json:"FirstLoginUpdate,omitnil,omitempty" name:"FirstLoginUpdate"`

	// 定期强制修改密码开关。开启后，用户需按照 PeriodValue 和 PeriodType 指定的周期定期修改密码，超过周期未修改将在登录时强制要求修改。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeriodUpdate *bool `json:"PeriodUpdate,omitnil,omitempty" name:"PeriodUpdate"`

	// 定期修改密码的周期数值，与 PeriodType 配合使用。例如 PeriodValue 为 6，PeriodType 为 MONTH，表示每 6 个月需修改一次密码。当 PeriodUpdate 为 true 时必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeriodValue *int64 `json:"PeriodValue,omitnil,omitempty" name:"PeriodValue"`

	// 定期修改密码的周期时间单位，与 PeriodValue 配合使用。取值范围：
	// WEEK：周
	// MONTH：月
	// YEAR：年
	// 当 PeriodUpdate 为 true 时必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeriodType *string `json:"PeriodType,omitnil,omitempty" name:"PeriodType"`
}

type PermissionInfo struct {
	// "READONLY",   //公有读，私有写。所有用户可读，仅创建者及管理员可写  
	// "PRIVATE",    //私有读写，仅创建者及管理员可读写  
	// "ADMINWRITE", //所有用户可读，仅管理员可写  
	// "ADMINONLY",  //仅管理员可操作  
	// "CUSTOM",     // 安全规则
	AclTag *string `json:"AclTag,omitnil,omitempty" name:"AclTag"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 自定义规则
	Rule *string `json:"Rule,omitnil,omitempty" name:"Rule"`
}

type PostgreSQLInfo struct {
	// <p>数据库名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>实例id</p>
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// <p>实例状态</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>地域</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>数据库引擎版本</p>
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`
}

// Predefined struct for user
type PreviewPGUserMigrationsRequestParams struct {
	// <p>云开发环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>预览要执行的migration 列表</p>
	Migrations []*MigrationInput `json:"Migrations,omitnil,omitempty" name:"Migrations"`

	// <p>标记请求来源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`
}

type PreviewPGUserMigrationsRequest struct {
	*tchttp.BaseRequest
	
	// <p>云开发环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>预览要执行的migration 列表</p>
	Migrations []*MigrationInput `json:"Migrations,omitnil,omitempty" name:"Migrations"`

	// <p>标记请求来源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`
}

func (r *PreviewPGUserMigrationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PreviewPGUserMigrationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Migrations")
	delete(f, "Source")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PreviewPGUserMigrationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PreviewPGUserMigrationsResponseParams struct {
	// <p>将要执行的migration列表</p>
	Pending []*MigrationPlanItem `json:"Pending,omitnil,omitempty" name:"Pending"`

	// <p>已经应用的migration列表</p>
	Applied []*MigrationPlanItem `json:"Applied,omitnil,omitempty" name:"Applied"`

	// <p>版本相同但 checksum 不一致冲突的migration列表</p>
	Conflicts []*MigrationConflict `json:"Conflicts,omitnil,omitempty" name:"Conflicts"`

	// <p>是否可直接执行；当前仅表示没有 checksum 冲突</p>
	Executable *bool `json:"Executable,omitnil,omitempty" name:"Executable"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PreviewPGUserMigrationsResponse struct {
	*tchttp.BaseResponse
	Response *PreviewPGUserMigrationsResponseParams `json:"Response"`
}

func (r *PreviewPGUserMigrationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PreviewPGUserMigrationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Provider struct {
	// 身份源的唯一标识符，用于在系统内区分不同的身份源。格式要求：2~32 位，仅支持小写英文字母和数字，不可包含空格或特殊字符。创建后不可修改
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 身份源的安全认证配置，包含与第三方平台对接所需的核心参数，如 ClientId、ClientSecret、授权端点、Token 端点、回调地址、Scope 等。不同 ProviderType 对应不同的配置项。CUSTOM 类型无需手动配置（系统自动填充），OIDC 类型会根据 Issuer 自动补全端点信息，SAML 类型需提供 SamlMetadata（最大 10KB）
	Config *ProviderConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 身份源的显示名称，支持国际化多语言配置。用户在登录页面看到的身份源名称将使用该字段，建议根据实际业务场景填写易于识别的名称。未传入时默认使用 Id 值作为显示名称
	Name *LocalizedMessage `json:"Name,omitnil,omitempty" name:"Name"`

	// 身份源图标的访问地址，将展示在登录页的身份源按钮上。建议使用 64×64 像素的 SVG 格式图片以保证清晰度，支持 HTTP/HTTPS 公网可访问的图片链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	Picture *string `json:"Picture,omitnil,omitempty" name:"Picture"`

	// 身份源对应的官方主页地址。该信息将在用户查看自己的第三方账号绑定列表时展示，帮助用户识别已绑定的身份源来源。例如 GitHub 身份源可填写：https://github.com
	// 注意：此字段可能返回 null，表示取不到有效值。
	Homepage *string `json:"Homepage,omitnil,omitempty" name:"Homepage"`

	// 身份源协议类型，决定该身份源使用何种认证协议与第三方平台对接。可选值：OAUTH（标准 OAuth 2.0 协议）、OIDC（OpenID Connect 协议）、SAML（SAML 2.0 协议）、CUSTOM（自定义登录，使用 RSA 密钥对签名验证）、EMAIL（邮箱登录，需配合 EmailConfig 使用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProviderType *string `json:"ProviderType,omitnil,omitempty" name:"ProviderType"`

	// 控制第三方身份源登录时是否自动注册系统用户。可选值：TRUE（始终自动注册，无论第三方返回的用户信息是否包含手机号或邮箱）、FALSE（不自动注册，需用户手动绑定）、UNSPECIFIED（默认行为：仅当第三方身份源返回的用户信息中包含手机号或邮箱时才自动注册，否则登录完成后要求用户绑定手机号方可继续使用）。注意：企业微信类型（WX_WORK_AGENT/WX_WORK_INTERNAL/WX_WORK_THIRD_PARTY/WX_WORK_THIRD_PARTY_ASSOCIATION）和微信小程序类型（WX_MICRO_APP/WX_QRCODE_MICRO_APP/WX_OPEN）在 UNSPECIFIED 时会自动设为 TRUE。当 TransparentMode 为 TRUE 时，该字段将被强制设为 FALSE
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoSignUpWithProviderUser *string `json:"AutoSignUpWithProviderUser,omitnil,omitempty" name:"AutoSignUpWithProviderUser"`

	// 身份源的启用状态。可选值：TRUE（启用，用户可通过该身份源登录）、FALSE（禁用，已有绑定关系不受影响）。未传入时默认为 TRUE（启用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	On *string `json:"On,omitnil,omitempty" name:"On"`

	// 身份源的详细描述信息，支持国际化多语言配置。可用于向用户说明该身份源的用途或使用场景。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *LocalizedMessage `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否开启信息透传模式。可选值：TRUE（仅登录模式：平台不持久化存储用户数据，仅将第三方身份源返回的用户信息透传给业务方，适用于不希望平台留存用户数据的场景）、FALSE（登录且注册模式：平台正常注册并存储用户信息，默认值）。注意：开启透传模式时，AutoSignUpWithProviderUser 将被强制设为 FALSE；若 ReuseUserId 为 UNSPECIFIED，将被自动设为 TRUE。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransparentMode *string `json:"TransparentMode,omitnil,omitempty" name:"TransparentMode"`

	// 是否直接复用第三方身份源返回的用户标识（如 OpenID、UnionID 等）作为平台用户 ID。可选值：TRUE（开启，平台用户 ID 将直接使用第三方身份源返回的用户标识，适用于已有用户体系迁移场景）、FALSE（关闭，由平台生成独立用户 ID）。注意：开启后需确保第三方用户标识的全局唯一性，避免 ID 冲突。当 TransparentMode 为 TRUE 且该字段为 UNSPECIFIED 时，将被自动设为 TRUE
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReuseUserId *string `json:"ReuseUserId,omitnil,omitempty" name:"ReuseUserId"`

	// 邮箱身份源的专项配置，仅当 ProviderType 为 EMAIL 时有效且必填。包含邮件服务商、发件人地址、SMTP 配置等参数，用于支持通过邮箱验证码方式进行身份认证。支持两种模式：自有 SMTP 服务器（需填写完整的 SMTP 配置）和平台代发（EmailConfig.On 设为 TRUE 时由平台随机分配 SMTP 服务器）
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmailConfig *EmailProviderConfig `json:"EmailConfig,omitnil,omitempty" name:"EmailConfig"`

	// 是否开启邮箱自动关联登录。可选值：TRUE（开启）、FALSE（关闭）、UNSPECIFIED（默认为 FALSE）。开启后，若第三方身份源返回的邮箱与系统中已有用户的邮箱一致，则自动将该第三方账号与已有用户关联绑定并完成登录，无需用户手动绑定
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoSignInWhenEmailMatch *string `json:"AutoSignInWhenEmailMatch,omitnil,omitempty" name:"AutoSignInWhenEmailMatch"`

	// 是否开启手机号自动关联登录。可选值：TRUE（开启）、FALSE（关闭）、UNSPECIFIED（默认行为等同于 TRUE，即默认开启）。开启后，若第三方身份源返回的手机号与系统中已有用户的手机号一致，则自动将该第三方账号与已有用户关联绑定并完成登录，无需用户手动绑定
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoSignInWhenPhoneNumberMatch *string `json:"AutoSignInWhenPhoneNumberMatch,omitnil,omitempty" name:"AutoSignInWhenPhoneNumberMatch"`
}

type ProviderConfig struct {
	// 身份提供方的唯一标识符（Issuer URL），用于验证 ID Token 中的 iss 字段。仅当 ProviderType 为 OIDC 时需要填写，值通常为第三方 OIDC 服务的根地址，例如：https://accounts.google.com。填写后平台将自动通过 /.well-known/openid-configuration 发现并填充 AuthorizationEndpoint、TokenEndpoint、UserinfoEndpoint、JwksUri 等端点地址。详情参考 OpenID Connect Discovery 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Issuer *string `json:"Issuer,omitnil,omitempty" name:"Issuer"`

	// 第三方身份提供方的 JSON Web Key Set 地址，用于获取公钥以验证 ID Token 签名。仅当 ProviderType 为 OIDC 时需要填写。若已填写 Issuer，该字段将通过 OpenID Connect Discovery 自动获取，无需手动填写。详情参考 OpenID Connect Discovery 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	JwksUri *string `json:"JwksUri,omitnil,omitempty" name:"JwksUri"`

	// 在第三方身份提供方注册的应用客户端 ID，用于标识当前接入应用。当 ProviderType 为 OIDC 或 OAUTH 时必须填写，可在对应平台的开发者控制台中获取。详情参考 OAuth 2.0 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 在第三方身份提供方注册的应用客户端密钥，与 ClientId 配合使用，用于在 Token 端点进行身份验证。当 ProviderType 为 OIDC 或 OAUTH 时必须填写，请妥善保管，避免泄露。详情参考 OAuth 2.0 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientSecret *string `json:"ClientSecret,omitnil,omitempty" name:"ClientSecret"`

	// OAuth 授权完成后第三方平台回调的地址，需与在第三方平台注册的回调地址完全一致，否则授权将失败。当 ProviderType 为 OIDC 或 OAUTH 时必须填写，并需在对应平台的开发者控制台中配置该地址为合法回调地址。详情参考 OAuth 2.0 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedirectUri *string `json:"RedirectUri,omitnil,omitempty" name:"RedirectUri"`

	// 向第三方身份提供方申请的权限范围，多个 scope 之间用空格分隔。当 ProviderType 为 OIDC 或 OAUTH 时必须填写，OIDC 场景下通常至少包含 openid，如需获取用户邮箱或手机号可追加 email、phone 等。若已填写 Issuer 且未指定 Scope，将自动使用 OpenID Connect Discovery 返回的 scopes_supported。详情参考 OAuth 2.0 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scope *string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 第三方身份提供方的授权端点地址，用于发起 OAuth/OIDC 授权请求，引导用户跳转至第三方登录页面。当 ProviderType 为 OIDC 或 OAUTH 时必须填写。若已填写 Issuer，该字段将通过 OpenID Connect Discovery 自动获取，无需手动填写。详情参考 OAuth 2.0 / OIDC 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitnil,omitempty" name:"AuthorizationEndpoint"`

	// 第三方身份提供方的 Token 端点地址，用于通过授权码（code）换取 Access Token 和 ID Token。当 ProviderType 为 OIDC 或 OAUTH 时必须填写。若已填写 Issuer，该字段将通过 OpenID Connect Discovery 自动获取，无需手动填写。详情参考 OAuth 2.0 / OIDC 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenEndpoint *string `json:"TokenEndpoint,omitnil,omitempty" name:"TokenEndpoint"`

	// 第三方身份提供方的用户信息端点地址，用于通过 Access Token 获取用户的基本信息（如昵称、头像、邮箱等）。当 ProviderType 为 OIDC 或 OAUTH 且需要获取用户详细信息时填写。若已填写 Issuer，该字段将通过 OpenID Connect Discovery 自动获取，无需手动填写。详情参考 OIDC 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserinfoEndpoint *string `json:"UserinfoEndpoint,omitnil,omitempty" name:"UserinfoEndpoint"`

	// OAuth/OIDC 授权请求的响应类型，决定授权端点返回的内容。可选值：code（授权码模式，推荐）、token（隐式模式，直接返回 Access Token）、id_token（直接返回 ID Token）。当 ProviderType 为 OIDC 时默认使用 id_token，其他类型默认使用 code。当 ProviderType 为 OIDC 或 OAUTH 时可选填写。详情参考 OAuth 2.0 / OIDC 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseType *string `json:"ResponseType,omitnil,omitempty" name:"ResponseType"`

	// 第三方身份提供方的单点退出端点地址。配置后，用户退出当前应用时将被跳转至该地址，使第三方 IDP 的登录态也一并失效，实现单点退出（SLO）。适用于 OIDC、OAUTH、SAML 等所有支持单点退出的身份源类型。不填则退出时仅清除本平台登录态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignoutEndpoint *string `json:"SignoutEndpoint,omitnil,omitempty" name:"SignoutEndpoint"`

	// Token 端点的客户端身份验证方式，决定请求 Token 时如何传递 ClientId 和 ClientSecret。可选值：CLIENT_SECRET_POST（将凭证放在请求 Body 中传递）、CLIENT_SECRET_BASIC（将凭证通过 HTTP Basic Auth Header 传递）。当 ProviderType 为 OIDC 或 OAUTH 时可选填写，默认使用 CLIENT_SECRET_POST。详情参考 OIDC 标准。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenEndpointAuthMethod *string `json:"TokenEndpointAuthMethod,omitnil,omitempty" name:"TokenEndpointAuthMethod"`

	// SAML 身份提供方的 Metadata XML 内容，包含 IDP 的实体 ID、SSO 端点地址、签名证书等关键信息，平台将据此完成 SAML 协议的对接配置。仅当 ProviderType 为 SAML 时可填写，通常可从第三方 IDP 的管理控制台中下载获取。详情参考 SAML 2.0 标准。
	SamlMetadata *string `json:"SamlMetadata,omitnil,omitempty" name:"SamlMetadata"`

	// 请求参数映射配置，用于处理非标准 OAuth 协议的参数转换。默认情况下平台严格遵循 OAuth 2.0 标准进行参数传递，若对接的第三方平台（如微信、企业微信等）使用了非标准的参数名称或传参方式，可通过该字段配置自定义的参数映射规则，以确保请求参数与第三方平台的要求一致。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestParametersMap *ProviderRequestParametersMap `json:"RequestParametersMap,omitnil,omitempty" name:"RequestParametersMap"`

	// 响应参数映射配置，用于处理非标准 OAuth 协议的响应参数转换。默认情况下平台严格遵循 OAuth 2.0 标准解析响应参数，若对接的第三方平台（如微信、企业微信等）返回了非标准的字段名称或数据结构，可通过该字段配置自定义的响应参数映射规则，将第三方返回的字段映射为平台标准字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseParametersMap *ProviderResponseParametersMap `json:"ResponseParametersMap,omitnil,omitempty" name:"ResponseParametersMap"`
}

type ProviderRequestParametersMap struct {
	// OAuth 标准协议中的 client_id。不同第三方平台的字段名称可能不同，例如微信平台对应 appid、新浪微博对应 app_id。
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// OAuth 标准协议中的 client_secret，用于身份认证源的密钥鉴权。请妥善保管，避免泄露。
	ClientSecret *string `json:"ClientSecret,omitnil,omitempty" name:"ClientSecret"`

	// OAuth 标准协议中的 redirect_uri，即授权回调地址。用户完成第三方认证后将重定向至该地址。
	RedirectUri *string `json:"RedirectUri,omitnil,omitempty" name:"RedirectUri"`

	// 身份源注册用户时自动绑定的角色 ID。配置后，通过该身份源注册的新用户将自动关联指定角色。
	RegisterUserRoleId *string `json:"RegisterUserRoleId,omitnil,omitempty" name:"RegisterUserRoleId"`

	// 身份源注册用户时是否自动授予许可证。取值范围：
	// TRUE：自动授权许可证
	// FALSE：不自动授权（默认值）
	RegisterUserAutoLicense *string `json:"RegisterUserAutoLicense,omitnil,omitempty" name:"RegisterUserAutoLicense"`

	// OAuth 获取 Token 时认证信息的请求位置。取值范围：
	// URL：将认证信息放在请求 URL 参数中
	// Headers：将认证信息放在请求 Header 中
	// Body：将认证信息放在请求 Body 中
	AuthPosition *string `json:"AuthPosition,omitnil,omitempty" name:"AuthPosition"`

	// OAuth 授权模式匹配的参数字段名。用于指定获取 Token 请求中 grant_type 参数对应的字段名称。
	GrantType *string `json:"GrantType,omitnil,omitempty" name:"GrantType"`

	// OAuth 授权模式类型。用于指定 grant_type 的值，例如 client_credentials 表示客户端凭证模式。
	ClientCredentials *string `json:"ClientCredentials,omitnil,omitempty" name:"ClientCredentials"`

	// OAuth 返回中 access_token 的映射字段名。若第三方平台返回的 Token 字段名不是标准的 access_token，可通过此字段指定实际字段名。
	AccessToken *string `json:"AccessToken,omitnil,omitempty" name:"AccessToken"`

	// OAuth 返回中 Token 有效期的映射字段名。若第三方平台返回的有效期字段名不是标准的 expires_in，可通过此字段指定实际字段名。
	ExpiresIn *string `json:"ExpiresIn,omitnil,omitempty" name:"ExpiresIn"`

	// 身份源注册用户时的用户类型。取值范围：
	// externalUser：外部用户
	// internalUser：内部用户
	// 默认值为 externalUser。
	RegisterUserType *string `json:"RegisterUserType,omitnil,omitempty" name:"RegisterUserType"`
}

type ProviderResponseParametersMap struct {
	// 用户唯一标识（sub）的映射字段名。对应 OIDC 标准中的 sub 字段，值为第三方平台返回的用户信息 JSON 中表示用户 ID 的字段路径。例如github平台填sub。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sub *string `json:"Sub,omitnil,omitempty" name:"Sub"`

	// 用户名称（name）的映射字段名。对应 OIDC 标准中的 name 字段，值为第三方平台返回的用户信息 JSON 中表示用户昵称或姓名的字段路径。例如github平台填 name。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户头像（picture）的映射字段名。对应 OIDC 标准中的 picture 字段，值为第三方平台返回的用户信息 JSON 中表示用户头像 URL 的字段路径。需要公网可访问的url。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Picture *string `json:"Picture,omitnil,omitempty" name:"Picture"`

	// 用户登录名（username）的映射字段名。对应 OIDC 标准中的 preferred_username 字段，值为第三方平台返回的用户信息 JSON 中表示用户唯一登录名的字段, 例如可使用sub或email等唯一值的字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 用户邮箱（email）的映射字段名。对应 OIDC 标准中的 email 字段，值为第三方平台返回的用户信息 JSON 中表示用户邮箱地址的字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 用户手机号（phone_number）的映射字段名。对应 OIDC 标准中的 phone_number 字段，值为第三方平台返回的用户信息 JSON 中表示用户手机号的字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNumber *string `json:"PhoneNumber,omitnil,omitempty" name:"PhoneNumber"`

	// 用户角色/分组（groups）的映射字段名。对应 OIDC 标准中的 groups 字段，值为第三方平台返回的用户信息 JSON 中表示用户所属角色或分组的字段路径。支持字符串数组类型的返回值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Groups *string `json:"Groups,omitnil,omitempty" name:"Groups"`
}

// Predefined struct for user
type PushPGUserMigrationsRequestParams struct {
	// <p>云开发环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>结构化 SQL migration 列表；每项包含 Query SQL 内容</p>
	Migrations []*MigrationInput `json:"Migrations,omitnil,omitempty" name:"Migrations"`

	// <p>等待获取数据库锁的最长时间</p><p>单位：毫秒</p><p>默认值：5000</p>
	LockTimeoutMs *int64 `json:"LockTimeoutMs,omitnil,omitempty" name:"LockTimeoutMs"`

	// <p>单条 SQL 执行最长时间，超过后由 PostgreSQL 取消该语句</p><p>单位：毫秒</p><p>默认值：300000</p>
	StatementTimeoutMs *int64 `json:"StatementTimeoutMs,omitnil,omitempty" name:"StatementTimeoutMs"`

	// <p>标记请求来源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`
}

type PushPGUserMigrationsRequest struct {
	*tchttp.BaseRequest
	
	// <p>云开发环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>结构化 SQL migration 列表；每项包含 Query SQL 内容</p>
	Migrations []*MigrationInput `json:"Migrations,omitnil,omitempty" name:"Migrations"`

	// <p>等待获取数据库锁的最长时间</p><p>单位：毫秒</p><p>默认值：5000</p>
	LockTimeoutMs *int64 `json:"LockTimeoutMs,omitnil,omitempty" name:"LockTimeoutMs"`

	// <p>单条 SQL 执行最长时间，超过后由 PostgreSQL 取消该语句</p><p>单位：毫秒</p><p>默认值：300000</p>
	StatementTimeoutMs *int64 `json:"StatementTimeoutMs,omitnil,omitempty" name:"StatementTimeoutMs"`

	// <p>标记请求来源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`
}

func (r *PushPGUserMigrationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushPGUserMigrationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Migrations")
	delete(f, "LockTimeoutMs")
	delete(f, "StatementTimeoutMs")
	delete(f, "Source")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PushPGUserMigrationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PushPGUserMigrationsResponseParams struct {
	// <p>任务ID</p><p>可通过DescribeTaskResult 接口查询进度</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PushPGUserMigrationsResponse struct {
	*tchttp.BaseResponse
	Response *PushPGUserMigrationsResponseParams `json:"Response"`
}

func (r *PushPGUserMigrationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PushPGUserMigrationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseEnvRequestParams struct {
	// <p>环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>分配请求ID</p>
	AllocateId *string `json:"AllocateId,omitnil,omitempty" name:"AllocateId"`
}

type ReleaseEnvRequest struct {
	*tchttp.BaseRequest
	
	// <p>环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>分配请求ID</p>
	AllocateId *string `json:"AllocateId,omitnil,omitempty" name:"AllocateId"`
}

func (r *ReleaseEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "AllocateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseEnvResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReleaseEnvResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseEnvResponseParams `json:"Response"`
}

func (r *ReleaseEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewEnvRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 续费周期，单位：月。
	// 默认值为 1，即续费1个月。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 是否自动选择代金券支付。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

type RenewEnvRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 续费周期，单位：月。
	// 默认值为 1，即续费1个月。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 是否自动选择代金券支付。
	AutoVoucher *bool `json:"AutoVoucher,omitnil,omitempty" name:"AutoVoucher"`
}

func (r *RenewEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewEnvRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "Period")
	delete(f, "AutoVoucher")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewEnvRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewEnvResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewEnvResponse struct {
	*tchttp.BaseResponse
	Response *RenewEnvResponseParams `json:"Response"`
}

func (r *RenewEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewEnvResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RepairPGUserMigrationHistoryRequestParams struct {
	// <p>云开发环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>migration版本</p><p>参数格式：14位时间格式</p><p>入参限制：纯数字</p>
	MigrationVersion *string `json:"MigrationVersion,omitnil,omitempty" name:"MigrationVersion"`

	// <p>migration 版本名</p><p>入参限制：限制小写字母和下划线</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>状态</p><p>枚举值：</p><ul><li>applied： 已应用</li><li>reverted： 表示删除 history 记录</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>修复原因</p>
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// <p>applied的时候填写，记录应用的sql语句</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type RepairPGUserMigrationHistoryRequest struct {
	*tchttp.BaseRequest
	
	// <p>云开发环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>migration版本</p><p>参数格式：14位时间格式</p><p>入参限制：纯数字</p>
	MigrationVersion *string `json:"MigrationVersion,omitnil,omitempty" name:"MigrationVersion"`

	// <p>migration 版本名</p><p>入参限制：限制小写字母和下划线</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>状态</p><p>枚举值：</p><ul><li>applied： 已应用</li><li>reverted： 表示删除 history 记录</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>修复原因</p>
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// <p>applied的时候填写，记录应用的sql语句</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *RepairPGUserMigrationHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RepairPGUserMigrationHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "MigrationVersion")
	delete(f, "Name")
	delete(f, "Status")
	delete(f, "Reason")
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RepairPGUserMigrationHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RepairPGUserMigrationHistoryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RepairPGUserMigrationHistoryResponse struct {
	*tchttp.BaseResponse
	Response *RepairPGUserMigrationHistoryResponseParams `json:"Response"`
}

func (r *RepairPGUserMigrationHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RepairPGUserMigrationHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourcePermission struct {
	// 资源类型。
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源标识
	Resource *string `json:"Resource,omitnil,omitempty" name:"Resource"`

	// 权限级别。取值：READONLY、PRIVATE、ADMINWRITE、ADMINONLY、CUSTOM。
	Permission *string `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 自定义安全规则配置，当 Permission 为 CUSTOM 时返回。
	SecurityRule *string `json:"SecurityRule,omitnil,omitempty" name:"SecurityRule"`
}

// Predefined struct for user
type RollbackPGUserMigrationsRequestParams struct {
	// <p>云开发环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>要回滚的条数</p><p>按照逆序回滚最近N条migration</p>
	LastN *int64 `json:"LastN,omitnil,omitempty" name:"LastN"`

	// <p>等待获取数据库锁的最长时间</p><p>单位：毫秒</p><p>默认值：5000</p>
	LockTimeoutMs *int64 `json:"LockTimeoutMs,omitnil,omitempty" name:"LockTimeoutMs"`

	// <p>单条 SQL 执行最长时间，超过后由 PostgreSQL 取消该语句</p><p>单位：毫秒</p><p>默认值：300000</p>
	StatementTimeoutMs *int64 `json:"StatementTimeoutMs,omitnil,omitempty" name:"StatementTimeoutMs"`

	// <p>标记API调用来源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`
}

type RollbackPGUserMigrationsRequest struct {
	*tchttp.BaseRequest
	
	// <p>云开发环境ID</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>要回滚的条数</p><p>按照逆序回滚最近N条migration</p>
	LastN *int64 `json:"LastN,omitnil,omitempty" name:"LastN"`

	// <p>等待获取数据库锁的最长时间</p><p>单位：毫秒</p><p>默认值：5000</p>
	LockTimeoutMs *int64 `json:"LockTimeoutMs,omitnil,omitempty" name:"LockTimeoutMs"`

	// <p>单条 SQL 执行最长时间，超过后由 PostgreSQL 取消该语句</p><p>单位：毫秒</p><p>默认值：300000</p>
	StatementTimeoutMs *int64 `json:"StatementTimeoutMs,omitnil,omitempty" name:"StatementTimeoutMs"`

	// <p>标记API调用来源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`
}

func (r *RollbackPGUserMigrationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackPGUserMigrationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "LastN")
	delete(f, "LockTimeoutMs")
	delete(f, "StatementTimeoutMs")
	delete(f, "Source")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackPGUserMigrationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackPGUserMigrationsResponseParams struct {
	// <p>任务ID</p><p>可通过DescribeTaskResult 接口查询进度</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>已成功回滚并删除 history 的 migration</p>
	RolledBack []*MigrationSummary `json:"RolledBack,omitnil,omitempty" name:"RolledBack"`

	// <p>未提供 Rollback SQL、视为成功并删除 history 的 migration</p>
	SkippedRollbackSql []*MigrationSummary `json:"SkippedRollbackSql,omitnil,omitempty" name:"SkippedRollbackSql"`

	// <p>执行 Rollback SQL 失败的 migration，可为空</p>
	Failed *MigrationSummary `json:"Failed,omitnil,omitempty" name:"Failed"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RollbackPGUserMigrationsResponse struct {
	*tchttp.BaseResponse
	Response *RollbackPGUserMigrationsResponseParams `json:"Response"`
}

func (r *RollbackPGUserMigrationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackPGUserMigrationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunCommandsRequestParams struct {
	// 待执行命令
	MgoCommands []*MgoCommandParam `json:"MgoCommands,omitnil,omitempty" name:"MgoCommands"`

	// 实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// Mongo连接器实例信息
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

type RunCommandsRequest struct {
	*tchttp.BaseRequest
	
	// 待执行命令
	MgoCommands []*MgoCommandParam `json:"MgoCommands,omitnil,omitempty" name:"MgoCommands"`

	// 实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 环境id
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// Mongo连接器实例信息
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

func (r *RunCommandsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunCommandsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MgoCommands")
	delete(f, "Tag")
	delete(f, "EnvId")
	delete(f, "MongoConnector")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunCommandsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunCommandsResponseParams struct {
	// 返回结果，返回结果为一个json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunCommandsResponse struct {
	*tchttp.BaseResponse
	Response *RunCommandsResponseParams `json:"Response"`
}

func (r *RunCommandsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunCommandsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunSqlRequestParams struct {
	// 要执行的SQL语句
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 数据库连接器实例信息
	DbInstance *DbInstance `json:"DbInstance,omitnil,omitempty" name:"DbInstance"`

	// 是否只读；当 `true` 时仅允许以 `SELECT/WITH/SHOW/DESCRIBE/DESC/EXPLAIN` 开头的 SQL
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`
}

type RunSqlRequest struct {
	*tchttp.BaseRequest
	
	// 要执行的SQL语句
	Sql *string `json:"Sql,omitnil,omitempty" name:"Sql"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 数据库连接器实例信息
	DbInstance *DbInstance `json:"DbInstance,omitnil,omitempty" name:"DbInstance"`

	// 是否只读；当 `true` 时仅允许以 `SELECT/WITH/SHOW/DESCRIBE/DESC/EXPLAIN` 开头的 SQL
	ReadOnly *bool `json:"ReadOnly,omitnil,omitempty" name:"ReadOnly"`
}

func (r *RunSqlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunSqlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Sql")
	delete(f, "EnvId")
	delete(f, "DbInstance")
	delete(f, "ReadOnly")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunSqlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunSqlResponseParams struct {
	// 查询结果行，每个元素为 JSON 字符串
	Items []*string `json:"Items,omitnil,omitempty" name:"Items"`

	// 列元数据信息，每个元素为 JSON 字符串，字段包含 `name/databaseType/nullable/length/precision/scale`
	Infos []*string `json:"Infos,omitnil,omitempty" name:"Infos"`

	// 受影响的行数（INSERT/UPDATE/DELETE 等语句）
	RowsAffected *int64 `json:"RowsAffected,omitnil,omitempty" name:"RowsAffected"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunSqlResponse struct {
	*tchttp.BaseResponse
	Response *RunSqlResponseParams `json:"Response"`
}

func (r *RunSqlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunSqlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SMSProviderTemplateConfig struct {
	// <p>短信服务商类型</p><p>枚举值：</p><ul><li>TENCENT_CN： 腾讯云国内短信</li><li>TENCENT_INTL： 腾讯云国际短信</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vendor *string `json:"Vendor,omitnil,omitempty" name:"Vendor"`

	// <p>短信服务商侧申请并审核通过的模板ID。</p><ul><li>腾讯云短信模板ID可前往 <a href="https://console.cloud.tencent.com/smsv2/csms-template">腾讯云国内短信</a> 或 <a href="https://console.cloud.tencent.com/smsv2/isms-template">国际/港澳台短信</a> 的正文模板管理查看，若向境外手机号发送短信，仅支持使用国际/港澳台短信模板。</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// <p>短信服务商侧为应用分配的ID，按照服务商文档和要求需要此参数的，填写该参数。</p><ul><li>腾讯云国内短信和国际短信此参数必填，可以在<a href="https://console.cloud.tencent.com/smsv2/app-manage">短信控制台</a>的应用列表中查看对应的应用id。</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SdkAppId *string `json:"SdkAppId,omitnil,omitempty" name:"SdkAppId"`

	// <p>短信服务商侧申请并审核通过的签名，按照服务商的文档和要求填写。</p><ul><li>腾讯云短信服务商，签名信息可前往 <a href="https://console.cloud.tencent.com/smsv2/csms-sign">国内短信</a> 或 <a href="https://console.cloud.tencent.com/smsv2/isms-sign">国际/港澳台短信</a> 的签名管理查看。<br> 注意：<ol><li>发送国内短信该参数必填，且需填写签名内容而非签名ID。</li><li>发送国际/港澳台短信该参数非必填。</li></ol></li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SignName *string `json:"SignName,omitnil,omitempty" name:"SignName"`

	// <p>调用短信服务商发送短信接口的调用秘钥对应的ID。</p><ul><li>调用api秘钥会保存在云开发平台控制台—扩展功能—授权管理中，如果对于短信调用的api秘钥有删除需求，可在此处进行删除，删除后，短信将无法正常发送。</li><li>腾讯云的调用api秘钥在腾讯云控制台获取，建议使用子账号的秘钥ID，并且按照最小权限配置。</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretId *string `json:"SecretId,omitnil,omitempty" name:"SecretId"`

	// <p>调用短信服务商发送短信接口的调用api秘钥对应的秘钥Key。</p><ul><li>腾讯云的调用api秘钥在腾讯云控制台获取，建议使用子账号的秘钥ID, 并且按照最小权限配置。平台对于调用api秘钥key是加密存储的，不会明文存储。</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// <p>短信服务商侧自定义短信发送的senderID，按照服务商文档和要求传参。</p><ul><li>仅国际化场景使用。部分国家/运营商支持自定义 Sender ID替代默认主叫号码。国内站点忽略此字段。</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SenderId *string `json:"SenderId,omitnil,omitempty" name:"SenderId"`

	// <p>当短信自定义模板含多个占位符时，平台只负责生成验证码值，其余占位符由调用方在此提供。</p><ul><li>无需提供验证码对应的占位的值，验证码由云开发平台侧生成。</li><li>如果是命名占位的服务商的短信模板，这里的参数按照需要对应的占位的key和value，会按照对应的key和value在发送短信时，填充到模板中。</li><li>如果是序号占位的服务商的短信模板，这里的参数不需要key, 只需要填写对应的value, 会按照填写的顺序依次填充到模板中。</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateExtendParam []*SMSTemplateParams `json:"TemplateExtendParam,omitnil,omitempty" name:"TemplateExtendParam"`
}

type SMSTemplateParams struct {
	// <p>短信模板的自定义参数的key。如果短信厂商的自定义参数按照命名占位的，才需要此参数；如果按照序号占位的， 不需要此参数。</p><p>腾讯云短信是按照序号占位的，不需要此参数。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>短信模板的自定义参数对应的value</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type SearchClsLogRequestParams struct {
	// 环境唯一ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 查询起始时间条件
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间条件
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询语句， 例如查询云函数：(src:app OR src:system) AND log:\"START RequestId*\"，  聚合云函数请求状态：* | select request_id, max(status_code) as status  where ((request_id='44738f94-16dd-11f1-****' AND retry_num=0) AND retry_num=0)) AND status_code!=202 group by request_id, retry_num 查询云数据库[文档型]：module:database， 查询云数据库[文档型]事件：module:database AND eventType:(MongoSlowQuery)，MongoSlowQuery为文档型数据库慢查询事件 查询云数据库[SQL型]：module:rdb， 查询云数据库[SQL型]事件：module:rdb AND eventType:(MysqlFreeze OR MysqlRecover OR MysqlSlowQuery)，MysqlFreeze为mysql数据库冻结事件、MysqlRecover为mysql数据库恢复事件、MysqlSlowQuery为mysql数据库慢查询事件 查询审批流：module:workflow， 查询模型：module:model， 查询用户权限：module:auth， 查询大模型：module:llm AND logType:llm-tracelog 查询网关服务调用：logType:accesslog 查询应用发布/删除事件：module:app AND eventType:(AppProdPub OR AppProdDel)，AppProdPub为应用发布事件，AppProdDel为应用删除事件 以上仅为示例语句，实际使用时请根据具体日志内容进行调整，查询语句需严格遵循CLS（Cloud Log Service）语法规范 详细的语法规则请参考官方档：https://cloud.tencent.com/document/product/614/47044
	QueryString *string `json:"QueryString,omitnil,omitempty" name:"QueryString"`

	// 单次要返回的日志条数，单次返回的最大条数为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 加载更多使用，透传上次返回的 context 值，获取后续的日志内容，通过游标最多可获取10000条，请尽可能缩小时间范围
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 按时间排序 asc（升序）或者 desc（降序），默认为 desc
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 是否使用Lucene语法，默认为false
	UseLucene *bool `json:"UseLucene,omitnil,omitempty" name:"UseLucene"`
}

type SearchClsLogRequest struct {
	*tchttp.BaseRequest
	
	// 环境唯一ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 查询起始时间条件
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询结束时间条件
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询语句， 例如查询云函数：(src:app OR src:system) AND log:\"START RequestId*\"，  聚合云函数请求状态：* | select request_id, max(status_code) as status  where ((request_id='44738f94-16dd-11f1-****' AND retry_num=0) AND retry_num=0)) AND status_code!=202 group by request_id, retry_num 查询云数据库[文档型]：module:database， 查询云数据库[文档型]事件：module:database AND eventType:(MongoSlowQuery)，MongoSlowQuery为文档型数据库慢查询事件 查询云数据库[SQL型]：module:rdb， 查询云数据库[SQL型]事件：module:rdb AND eventType:(MysqlFreeze OR MysqlRecover OR MysqlSlowQuery)，MysqlFreeze为mysql数据库冻结事件、MysqlRecover为mysql数据库恢复事件、MysqlSlowQuery为mysql数据库慢查询事件 查询审批流：module:workflow， 查询模型：module:model， 查询用户权限：module:auth， 查询大模型：module:llm AND logType:llm-tracelog 查询网关服务调用：logType:accesslog 查询应用发布/删除事件：module:app AND eventType:(AppProdPub OR AppProdDel)，AppProdPub为应用发布事件，AppProdDel为应用删除事件 以上仅为示例语句，实际使用时请根据具体日志内容进行调整，查询语句需严格遵循CLS（Cloud Log Service）语法规范 详细的语法规则请参考官方档：https://cloud.tencent.com/document/product/614/47044
	QueryString *string `json:"QueryString,omitnil,omitempty" name:"QueryString"`

	// 单次要返回的日志条数，单次返回的最大条数为100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 加载更多使用，透传上次返回的 context 值，获取后续的日志内容，通过游标最多可获取10000条，请尽可能缩小时间范围
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 按时间排序 asc（升序）或者 desc（降序），默认为 desc
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 是否使用Lucene语法，默认为false
	UseLucene *bool `json:"UseLucene,omitnil,omitempty" name:"UseLucene"`
}

func (r *SearchClsLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClsLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "QueryString")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "Sort")
	delete(f, "UseLucene")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchClsLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchClsLogResponseParams struct {
	// 日志内容结果
	LogResults *LogResObject `json:"LogResults,omitnil,omitempty" name:"LogResults"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SearchClsLogResponse struct {
	*tchttp.BaseResponse
	Response *SearchClsLogResponseParams `json:"Response"`
}

func (r *SearchClsLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchClsLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StaticStorageInfo struct {
	// <p>静态CDN域名</p>
	StaticDomain *string `json:"StaticDomain,omitnil,omitempty" name:"StaticDomain"`

	// <p>静态CDN默认文件夹，当前为根目录</p>
	DefaultDirName *string `json:"DefaultDirName,omitnil,omitempty" name:"DefaultDirName"`

	// <p>资源状态(process/online/offline/init)</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>cos所属区域</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>bucket信息</p>
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// <p>到期时间（秒级时间戳）</p>
	AccessExpire *int64 `json:"AccessExpire,omitnil,omitempty" name:"AccessExpire"`

	// <p>外部存储。</p>
	ExternalStorage *ExternalStorage `json:"ExternalStorage,omitnil,omitempty" name:"ExternalStorage"`
}

type StaticStoreInfo struct {
	// 环境ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// 静态域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdnDomain *string `json:"CdnDomain,omitnil,omitempty" name:"CdnDomain"`

	// COS桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// cos区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Regoin is deprecated.
	Regoin *string `json:"Regoin,omitnil,omitempty" name:"Regoin"`

	// 资源状态:init(初始化)/process(处理中)/online(上线)/destroying(销毁中)/offline(下线))
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type StorageInfo struct {
	// <p>资源所属地域。<br>当前支持ap-shanghai</p>
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>桶名，存储资源的唯一标识</p>
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// <p>cdn 域名</p>
	CdnDomain *string `json:"CdnDomain,omitnil,omitempty" name:"CdnDomain"`

	// <p>资源所属用户的腾讯云appId</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>外部存储介质相关信息。</p>
	ExternalStorage *ExternalStorage `json:"ExternalStorage,omitnil,omitempty" name:"ExternalStorage"`
}

type TableInfo struct {
	// 表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 表中文档数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 表的大小（即表中文档总大小），单位：字节
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 索引数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexCount *int64 `json:"IndexCount,omitnil,omitempty" name:"IndexCount"`

	// 索引占用空间，单位：字节
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexSize *int64 `json:"IndexSize,omitnil,omitempty" name:"IndexSize"`
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TkeClusterInfo struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群的vpcId
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 版本内网CLB所在子网Id
	VersionClbSubnetId *string `json:"VersionClbSubnetId,omitnil,omitempty" name:"VersionClbSubnetId"`
}

// Predefined struct for user
type UnbindStorageSourceRequestParams struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

type UnbindStorageSourceRequest struct {
	*tchttp.BaseRequest
	
	// 环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`
}

func (r *UnbindStorageSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindStorageSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindStorageSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindStorageSourceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindStorageSourceResponse struct {
	*tchttp.BaseResponse
	Response *UnbindStorageSourceResponseParams `json:"Response"`
}

func (r *UnbindStorageSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindStorageSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAIModelRequestParams struct {
	// <p>环境id</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>分组名</p>
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// <p>模型地址</p><p>枚举值：</p><ul><li>http://default.tcb： 默认模型地址，custom模型切换为builtin模型时使用</li></ul>
	BaseUrl *string `json:"BaseUrl,omitnil,omitempty" name:"BaseUrl"`

	// <p>模型名列表</p><p>Models 列表更新采用全量替换</p>
	Models []*AIModel `json:"Models,omitnil,omitempty" name:"Models"`

	// <p>备注</p>
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// <p>模型状态, 1: 开启, 2: 关闭</p>
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>模型密钥</p>
	Secret *AIModelSecret `json:"Secret,omitnil,omitempty" name:"Secret"`
}

type UpdateAIModelRequest struct {
	*tchttp.BaseRequest
	
	// <p>环境id</p>
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// <p>分组名</p>
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// <p>模型地址</p><p>枚举值：</p><ul><li>http://default.tcb： 默认模型地址，custom模型切换为builtin模型时使用</li></ul>
	BaseUrl *string `json:"BaseUrl,omitnil,omitempty" name:"BaseUrl"`

	// <p>模型名列表</p><p>Models 列表更新采用全量替换</p>
	Models []*AIModel `json:"Models,omitnil,omitempty" name:"Models"`

	// <p>备注</p>
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// <p>模型状态, 1: 开启, 2: 关闭</p>
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>模型密钥</p>
	Secret *AIModelSecret `json:"Secret,omitnil,omitempty" name:"Secret"`
}

func (r *UpdateAIModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAIModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvId")
	delete(f, "GroupName")
	delete(f, "BaseUrl")
	delete(f, "Models")
	delete(f, "Remark")
	delete(f, "Status")
	delete(f, "Secret")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAIModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAIModelResponseParams struct {
	// <p>更新数量</p>
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAIModelResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAIModelResponseParams `json:"Response"`
}

func (r *UpdateAIModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAIModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTableRequestParams struct {
	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// FlexDB实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 待删除索引信息
	DropIndexes []*DropIndex `json:"DropIndexes,omitnil,omitempty" name:"DropIndexes"`

	// 待创建索引信息
	CreateIndexes []*CreateIndex `json:"CreateIndexes,omitnil,omitempty" name:"CreateIndexes"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoDB连接器配置
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

type UpdateTableRequest struct {
	*tchttp.BaseRequest
	
	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// FlexDB实例ID
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 待删除索引信息
	DropIndexes []*DropIndex `json:"DropIndexes,omitnil,omitempty" name:"DropIndexes"`

	// 待创建索引信息
	CreateIndexes []*CreateIndex `json:"CreateIndexes,omitnil,omitempty" name:"CreateIndexes"`

	// 云开发环境ID
	EnvId *string `json:"EnvId,omitnil,omitempty" name:"EnvId"`

	// MongoDB连接器配置
	MongoConnector *MongoConnector `json:"MongoConnector,omitnil,omitempty" name:"MongoConnector"`
}

func (r *UpdateTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableName")
	delete(f, "Tag")
	delete(f, "DropIndexes")
	delete(f, "CreateIndexes")
	delete(f, "EnvId")
	delete(f, "MongoConnector")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateTableResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateTableResponse struct {
	*tchttp.BaseResponse
	Response *UpdateTableResponseParams `json:"Response"`
}

func (r *UpdateTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type User struct {
	// 用户ID
	Uid *string `json:"Uid,omitnil,omitempty" name:"Uid"`

	// 用户名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 用户类型：internalUser-内部用户、externalUser-外部用户
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 用户状态：ACTIVE（激活）、BLOCKED（冻结）
	UserStatus *string `json:"UserStatus,omitnil,omitempty" name:"UserStatus"`

	// 用户昵称
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 手机号
	Phone *string `json:"Phone,omitnil,omitempty" name:"Phone"`

	// 邮箱
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`

	// 头像链接
	AvatarUrl *string `json:"AvatarUrl,omitnil,omitempty" name:"AvatarUrl"`

	// 用户描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type VMLoginConfiguration struct {
	// 登录方式。扫码登录时指定为 SCAN_LOGIN
	LoginType *string `json:"LoginType,omitnil,omitempty" name:"LoginType"`

	// 是否自动生成密码
	AutoGeneratePassword *string `json:"AutoGeneratePassword,omitnil,omitempty" name:"AutoGeneratePassword"`

	// 指定密码登录
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 绑定密钥ID
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

type VMPrice struct {
	// 价格货币单位。取值范围CNY:人民币。USD:美元。
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// 原始价格
	OriginalPrice *float64 `json:"OriginalPrice,omitnil,omitempty" name:"OriginalPrice"`

	// 折扣率
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 折扣后的价格
	DiscountPrice *float64 `json:"DiscountPrice,omitnil,omitempty" name:"DiscountPrice"`

	// 折扣前每天资源点
	OriginalCredits *float64 `json:"OriginalCredits,omitnil,omitempty" name:"OriginalCredits"`

	// 折扣后每天所需资源点
	DiscountCredits *float64 `json:"DiscountCredits,omitnil,omitempty" name:"DiscountCredits"`
}

type VMSpec struct {
	// LightHouse=轻量云服务器
	// CVM=云服务器
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 轻量云服务器规格。
	// 当Type=LightHouse时有效
	LightHouseSpec *VMSpecLightHouse `json:"LightHouseSpec,omitnil,omitempty" name:"LightHouseSpec"`

	// 价格信息
	Price *VMPrice `json:"Price,omitnil,omitempty" name:"Price"`
}

type VMSpecLightHouse struct {
	// LH主机的BundleId
	BundleId *string `json:"BundleId,omitnil,omitempty" name:"BundleId"`

	// 主机配置详情json
	BundleConfig *string `json:"BundleConfig,omitnil,omitempty" name:"BundleConfig"`
}

type VerificationConfig struct {
	// <p>短信验证码发送通道类型。</p><p>枚举值：</p><ul><li>default： 使用默认云开发短信包发送短信</li><li>apis： 使用云开发自定义 APIs 作为短信发送通道，需配合 Name 和 Method 参数使用。不传则不修改当前配置。</li><li>template： 自定义短信模板配置，需要配置TemplateProvider</li></ul>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>自定义 APIs 数据源唯一标识，当 Type 为 apis 时必填。用于定位微搭 APIs 中对应的数据源。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>自定义 APIs 方法名，当 Type 为 apis 时必填。指定微搭 APIs 中用于发送验证码的方法。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// <p>单个手机号每日短信发送上限。默认值为 30，传 -1 表示不限制，如果设置为不限制，需要注意恶意攻击，导致短信套餐用量计费问题。仅支持正整数或 -1。不传则不修改当前配置。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmsDayLimit *int64 `json:"SmsDayLimit,omitnil,omitempty" name:"SmsDayLimit"`

	// <p>自定义短信服务商模板配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateProvider *SMSProviderTemplateConfig `json:"TemplateProvider,omitnil,omitempty" name:"TemplateProvider"`
}

type VmInstance struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type WxGatewayCustomConfig struct {
	// 是否开启x-real-ip
	IsOpenXRealIp *bool `json:"IsOpenXRealIp,omitnil,omitempty" name:"IsOpenXRealIp"`

	// 封禁配置
	BanConfig *BanConfig `json:"BanConfig,omitnil,omitempty" name:"BanConfig"`

	// 获取源ip方式，PPV1(Proxy Protocol V1)、PPV2(Proxy Protocol V2)、TOA(tcp option address)
	SourceIpType *string `json:"SourceIpType,omitnil,omitempty" name:"SourceIpType"`

	// 日志信息
	LogConfig *CustomLogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// 是否开启http1.0
	IsAcceptHttpOne *bool `json:"IsAcceptHttpOne,omitnil,omitempty" name:"IsAcceptHttpOne"`
}