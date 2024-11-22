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

package v20240713

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type CreateIAPUserOIDCConfigRequestParams struct {
	// 身份提供商URL。OpenID Connect身份提供商标识。对应企业IdP提供的Openid-configuration中"issuer"字段的值。
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// 客户端ID，在OpenID Connect身份提供商注册的客户端ID。
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 授权请求Endpoint，OpenID Connect身份提供商授权地址。对应企业IdP提供的Openid-configuration中"authorization_endpoint"字段的值。
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitnil,omitempty" name:"AuthorizationEndpoint"`

	// 授权请求Response type，固定值id_token
	ResponseType *string `json:"ResponseType,omitnil,omitempty" name:"ResponseType"`

	// 授权请求Response mode。授权请求返回模式，form_post和fragment两种可选模式，推荐选择form_post模式。
	ResponseMode *string `json:"ResponseMode,omitnil,omitempty" name:"ResponseMode"`

	// 映射字段名称。IdP的id_token中哪一个字段映射到子用户的用户名，通常是sub或者name字段
	MappingFiled *string `json:"MappingFiled,omitnil,omitempty" name:"MappingFiled"`

	// 签名公钥，需要base64_encode。验证OpenID Connect身份提供商ID Token签名的公钥。为了您的账号安全，建议您定期轮换签名公钥。
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// 授权请求Scope。openid; email;profile。授权请求信息范围。默认必选openid。
	Scope []*string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateIAPUserOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
	// 身份提供商URL。OpenID Connect身份提供商标识。对应企业IdP提供的Openid-configuration中"issuer"字段的值。
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// 客户端ID，在OpenID Connect身份提供商注册的客户端ID。
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 授权请求Endpoint，OpenID Connect身份提供商授权地址。对应企业IdP提供的Openid-configuration中"authorization_endpoint"字段的值。
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitnil,omitempty" name:"AuthorizationEndpoint"`

	// 授权请求Response type，固定值id_token
	ResponseType *string `json:"ResponseType,omitnil,omitempty" name:"ResponseType"`

	// 授权请求Response mode。授权请求返回模式，form_post和fragment两种可选模式，推荐选择form_post模式。
	ResponseMode *string `json:"ResponseMode,omitnil,omitempty" name:"ResponseMode"`

	// 映射字段名称。IdP的id_token中哪一个字段映射到子用户的用户名，通常是sub或者name字段
	MappingFiled *string `json:"MappingFiled,omitnil,omitempty" name:"MappingFiled"`

	// 签名公钥，需要base64_encode。验证OpenID Connect身份提供商ID Token签名的公钥。为了您的账号安全，建议您定期轮换签名公钥。
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// 授权请求Scope。openid; email;profile。授权请求信息范围。默认必选openid。
	Scope []*string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateIAPUserOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIAPUserOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdentityUrl")
	delete(f, "ClientId")
	delete(f, "AuthorizationEndpoint")
	delete(f, "ResponseType")
	delete(f, "ResponseMode")
	delete(f, "MappingFiled")
	delete(f, "IdentityKey")
	delete(f, "Scope")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIAPUserOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIAPUserOIDCConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateIAPUserOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateIAPUserOIDCConfigResponseParams `json:"Response"`
}

func (r *CreateIAPUserOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIAPUserOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIAPLoginSessionDurationRequestParams struct {

}

type DescribeIAPLoginSessionDurationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeIAPLoginSessionDurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIAPLoginSessionDurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIAPLoginSessionDurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIAPLoginSessionDurationResponseParams struct {
	// 登录会话时长
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIAPLoginSessionDurationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIAPLoginSessionDurationResponseParams `json:"Response"`
}

func (r *DescribeIAPLoginSessionDurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIAPLoginSessionDurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIAPUserOIDCConfigRequestParams struct {

}

type DescribeIAPUserOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeIAPUserOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIAPUserOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIAPUserOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIAPUserOIDCConfigResponseParams struct {
	// 身份提供商类型。 13：IAP用户OIDC身份提供商
	ProviderType *uint64 `json:"ProviderType,omitnil,omitempty" name:"ProviderType"`

	// 身份提供商URL
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// 签名公钥
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// 客户端id
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 状态：0:未设置，11:已开启，2:已禁用
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// HTTPS CA证书的验证指纹，允许英文字母和数字，每个指纹长度为40个字符，最多5个指纹。
	Fingerprints []*string `json:"Fingerprints,omitnil,omitempty" name:"Fingerprints"`

	// 是否需要开启自动使用OIDC签名公钥，1:需要，2:不需要，默认不需要
	EnableAutoPublicKey *uint64 `json:"EnableAutoPublicKey,omitnil,omitempty" name:"EnableAutoPublicKey"`

	// 授权请求Endpoint
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitnil,omitempty" name:"AuthorizationEndpoint"`

	// 授权请求Scope
	Scope []*string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 授权请求Response type
	ResponseType *string `json:"ResponseType,omitnil,omitempty" name:"ResponseType"`

	// 授权请求Response mode
	ResponseMode *string `json:"ResponseMode,omitnil,omitempty" name:"ResponseMode"`

	// 映射字段名称
	MappingFiled *string `json:"MappingFiled,omitnil,omitempty" name:"MappingFiled"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIAPUserOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIAPUserOIDCConfigResponseParams `json:"Response"`
}

func (r *DescribeIAPUserOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIAPUserOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableIAPUserSSORequestParams struct {

}

type DisableIAPUserSSORequest struct {
	*tchttp.BaseRequest
	
}

func (r *DisableIAPUserSSORequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableIAPUserSSORequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableIAPUserSSORequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableIAPUserSSOResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableIAPUserSSOResponse struct {
	*tchttp.BaseResponse
	Response *DisableIAPUserSSOResponseParams `json:"Response"`
}

func (r *DisableIAPUserSSOResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableIAPUserSSOResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIAPLoginSessionDurationRequestParams struct {
	// 登录会话时长
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`
}

type ModifyIAPLoginSessionDurationRequest struct {
	*tchttp.BaseRequest
	
	// 登录会话时长
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`
}

func (r *ModifyIAPLoginSessionDurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIAPLoginSessionDurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Duration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIAPLoginSessionDurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIAPLoginSessionDurationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyIAPLoginSessionDurationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIAPLoginSessionDurationResponseParams `json:"Response"`
}

func (r *ModifyIAPLoginSessionDurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIAPLoginSessionDurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateIAPUserOIDCConfigRequestParams struct {
	// 身份提供商URL。OpenID Connect身份提供商标识。对应企业IdP提供的Openid-configuration中"issuer"字段的值。
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// 客户端ID，在OpenID Connect身份提供商注册的客户端ID。
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 授权请求Endpoint，OpenID Connect身份提供商授权地址。对应企业IdP提供的Openid-configuration中"authorization_endpoint"字段的值。
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitnil,omitempty" name:"AuthorizationEndpoint"`

	// 授权请求Response type，固定值id_token
	ResponseType *string `json:"ResponseType,omitnil,omitempty" name:"ResponseType"`

	// 授权请求Response mode。授权请求返回模式，form_post和fragment两种可选模式，推荐选择form_post模式。
	ResponseMode *string `json:"ResponseMode,omitnil,omitempty" name:"ResponseMode"`

	// 映射字段名称。IdP的id_token中哪一个字段映射到子用户的用户名，通常是sub或者name字段
	MappingFiled *string `json:"MappingFiled,omitnil,omitempty" name:"MappingFiled"`

	// RSA签名公钥，JWKS格式，需要进行base64_encode。验证OpenID Connect身份提供商ID Token签名的公钥。为了您的账号安全，建议您定期轮换签名公钥。
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// 授权请求Scope。openid; email;profile。授权请求信息范围。默认必选openid。
	Scope []*string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 描述，长度为1~255个英文或中文字符，默认值为空。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type UpdateIAPUserOIDCConfigRequest struct {
	*tchttp.BaseRequest
	
	// 身份提供商URL。OpenID Connect身份提供商标识。对应企业IdP提供的Openid-configuration中"issuer"字段的值。
	IdentityUrl *string `json:"IdentityUrl,omitnil,omitempty" name:"IdentityUrl"`

	// 客户端ID，在OpenID Connect身份提供商注册的客户端ID。
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 授权请求Endpoint，OpenID Connect身份提供商授权地址。对应企业IdP提供的Openid-configuration中"authorization_endpoint"字段的值。
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitnil,omitempty" name:"AuthorizationEndpoint"`

	// 授权请求Response type，固定值id_token
	ResponseType *string `json:"ResponseType,omitnil,omitempty" name:"ResponseType"`

	// 授权请求Response mode。授权请求返回模式，form_post和fragment两种可选模式，推荐选择form_post模式。
	ResponseMode *string `json:"ResponseMode,omitnil,omitempty" name:"ResponseMode"`

	// 映射字段名称。IdP的id_token中哪一个字段映射到子用户的用户名，通常是sub或者name字段
	MappingFiled *string `json:"MappingFiled,omitnil,omitempty" name:"MappingFiled"`

	// RSA签名公钥，JWKS格式，需要进行base64_encode。验证OpenID Connect身份提供商ID Token签名的公钥。为了您的账号安全，建议您定期轮换签名公钥。
	IdentityKey *string `json:"IdentityKey,omitnil,omitempty" name:"IdentityKey"`

	// 授权请求Scope。openid; email;profile。授权请求信息范围。默认必选openid。
	Scope []*string `json:"Scope,omitnil,omitempty" name:"Scope"`

	// 描述，长度为1~255个英文或中文字符，默认值为空。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *UpdateIAPUserOIDCConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateIAPUserOIDCConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IdentityUrl")
	delete(f, "ClientId")
	delete(f, "AuthorizationEndpoint")
	delete(f, "ResponseType")
	delete(f, "ResponseMode")
	delete(f, "MappingFiled")
	delete(f, "IdentityKey")
	delete(f, "Scope")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateIAPUserOIDCConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateIAPUserOIDCConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateIAPUserOIDCConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateIAPUserOIDCConfigResponseParams `json:"Response"`
}

func (r *UpdateIAPUserOIDCConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateIAPUserOIDCConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}