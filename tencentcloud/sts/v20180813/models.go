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

package v20180813

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ApiKey struct {
	// 密钥ID
	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

	// 创建时间(时间戳)
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 状态(2:有效, 3:禁用, 4:已删除)
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

// Predefined struct for user
type AssumeRoleRequestParams struct {
	// 角色的资源描述，可在[访问管理](https://console.cloud.tencent.com/cam/role)，点击角色名获取。
	// 普通角色：
	// qcs::cam::uin/12345678:role/4611686018427397919、qcs::cam::uin/12345678:roleName/testRoleName
	// 服务角色：
	// qcs::cam::uin/12345678:role/tencentcloudServiceRole/4611686018427397920、qcs::cam::uin/12345678:role/tencentcloudServiceRoleName/testServiceRoleName
	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`

	// 临时会话名称，由用户自定义名称。
	// 长度在2到128之间，可包含大小写字符，数字以及特殊字符：=,.@_-。 正则为：[\w+=,.@_-]*
	RoleSessionName *string `json:"RoleSessionName,omitempty" name:"RoleSessionName"`

	// 指定临时证书的有效期，单位：秒，默认 7200 秒，最长可设定有效期为 43200 秒
	DurationSeconds *uint64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`

	// 策略描述
	// 注意：
	// 1、policy 需要做 urlencode（如果通过 GET 方法请求云 API，发送请求前，所有参数都需要按照[云 API 规范](https://cloud.tencent.com/document/api/598/33159#1.-.E6.8B.BC.E6.8E.A5.E8.A7.84.E8.8C.83.E8.AF.B7.E6.B1.82.E4.B8.B2)再 urlencode 一次）。
	// 2、策略语法参照[ CAM 策略语法](https://cloud.tencent.com/document/product/598/10603)。
	// 3、策略中不能包含 principal 元素。
	Policy *string `json:"Policy,omitempty" name:"Policy"`

	// 角色外部ID，可在[访问管理](https://console.cloud.tencent.com/cam/role)，点击角色名获取。
	// 长度在2到128之间，可包含大小写字符，数字以及特殊字符：=,.@:/-。 正则为：[\w+=,.@:\/-]*
	ExternalId *string `json:"ExternalId,omitempty" name:"ExternalId"`

	// 会话标签列表。最多可以传递 50 个会话标签，不支持包含相同标签键。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 调用者身份uin
	SourceIdentity *string `json:"SourceIdentity,omitempty" name:"SourceIdentity"`
}

type AssumeRoleRequest struct {
	*tchttp.BaseRequest
	
	// 角色的资源描述，可在[访问管理](https://console.cloud.tencent.com/cam/role)，点击角色名获取。
	// 普通角色：
	// qcs::cam::uin/12345678:role/4611686018427397919、qcs::cam::uin/12345678:roleName/testRoleName
	// 服务角色：
	// qcs::cam::uin/12345678:role/tencentcloudServiceRole/4611686018427397920、qcs::cam::uin/12345678:role/tencentcloudServiceRoleName/testServiceRoleName
	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`

	// 临时会话名称，由用户自定义名称。
	// 长度在2到128之间，可包含大小写字符，数字以及特殊字符：=,.@_-。 正则为：[\w+=,.@_-]*
	RoleSessionName *string `json:"RoleSessionName,omitempty" name:"RoleSessionName"`

	// 指定临时证书的有效期，单位：秒，默认 7200 秒，最长可设定有效期为 43200 秒
	DurationSeconds *uint64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`

	// 策略描述
	// 注意：
	// 1、policy 需要做 urlencode（如果通过 GET 方法请求云 API，发送请求前，所有参数都需要按照[云 API 规范](https://cloud.tencent.com/document/api/598/33159#1.-.E6.8B.BC.E6.8E.A5.E8.A7.84.E8.8C.83.E8.AF.B7.E6.B1.82.E4.B8.B2)再 urlencode 一次）。
	// 2、策略语法参照[ CAM 策略语法](https://cloud.tencent.com/document/product/598/10603)。
	// 3、策略中不能包含 principal 元素。
	Policy *string `json:"Policy,omitempty" name:"Policy"`

	// 角色外部ID，可在[访问管理](https://console.cloud.tencent.com/cam/role)，点击角色名获取。
	// 长度在2到128之间，可包含大小写字符，数字以及特殊字符：=,.@:/-。 正则为：[\w+=,.@:\/-]*
	ExternalId *string `json:"ExternalId,omitempty" name:"ExternalId"`

	// 会话标签列表。最多可以传递 50 个会话标签，不支持包含相同标签键。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 调用者身份uin
	SourceIdentity *string `json:"SourceIdentity,omitempty" name:"SourceIdentity"`
}

func (r *AssumeRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssumeRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleArn")
	delete(f, "RoleSessionName")
	delete(f, "DurationSeconds")
	delete(f, "Policy")
	delete(f, "ExternalId")
	delete(f, "Tags")
	delete(f, "SourceIdentity")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssumeRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssumeRoleResponseParams struct {
	// 临时安全证书
	Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

	// 证书无效的时间，返回 Unix 时间戳，精确到秒
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 证书无效的时间，以 iso8601 格式的 UTC 时间表示
	Expiration *string `json:"Expiration,omitempty" name:"Expiration"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssumeRoleResponse struct {
	*tchttp.BaseResponse
	Response *AssumeRoleResponseParams `json:"Response"`
}

func (r *AssumeRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssumeRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssumeRoleWithSAMLRequestParams struct {
	// base64 编码的 SAML 断言信息
	SAMLAssertion *string `json:"SAMLAssertion,omitempty" name:"SAMLAssertion"`

	// 扮演者访问描述名
	PrincipalArn *string `json:"PrincipalArn,omitempty" name:"PrincipalArn"`

	// 角色访问描述名
	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`

	// 会话名称
	RoleSessionName *string `json:"RoleSessionName,omitempty" name:"RoleSessionName"`

	// 指定临时证书的有效期，单位：秒，默认 7200 秒，最长可设定有效期为 43200 秒
	DurationSeconds *uint64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`
}

type AssumeRoleWithSAMLRequest struct {
	*tchttp.BaseRequest
	
	// base64 编码的 SAML 断言信息
	SAMLAssertion *string `json:"SAMLAssertion,omitempty" name:"SAMLAssertion"`

	// 扮演者访问描述名
	PrincipalArn *string `json:"PrincipalArn,omitempty" name:"PrincipalArn"`

	// 角色访问描述名
	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`

	// 会话名称
	RoleSessionName *string `json:"RoleSessionName,omitempty" name:"RoleSessionName"`

	// 指定临时证书的有效期，单位：秒，默认 7200 秒，最长可设定有效期为 43200 秒
	DurationSeconds *uint64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`
}

func (r *AssumeRoleWithSAMLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssumeRoleWithSAMLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SAMLAssertion")
	delete(f, "PrincipalArn")
	delete(f, "RoleArn")
	delete(f, "RoleSessionName")
	delete(f, "DurationSeconds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssumeRoleWithSAMLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssumeRoleWithSAMLResponseParams struct {
	// 对象里面包含 Token，TmpSecretId，TmpSecretKey 三元组
	Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

	// 证书无效的时间，返回 Unix 时间戳，精确到秒
	ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 证书无效的时间，以 ISO8601 格式的 UTC 时间表示
	Expiration *string `json:"Expiration,omitempty" name:"Expiration"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssumeRoleWithSAMLResponse struct {
	*tchttp.BaseResponse
	Response *AssumeRoleWithSAMLResponseParams `json:"Response"`
}

func (r *AssumeRoleWithSAMLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssumeRoleWithSAMLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssumeRoleWithWebIdentityRequestParams struct {
	// 身份提供商名称
	ProviderId *string `json:"ProviderId,omitempty" name:"ProviderId"`

	// IdP签发的OIDC令牌
	WebIdentityToken *string `json:"WebIdentityToken,omitempty" name:"WebIdentityToken"`

	// 角色访问描述名
	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`

	// 会话名称
	RoleSessionName *string `json:"RoleSessionName,omitempty" name:"RoleSessionName"`

	// 指定临时证书的有效期，单位：秒，默认 7200 秒，最长可设定有效期为 43200 秒
	DurationSeconds *int64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`
}

type AssumeRoleWithWebIdentityRequest struct {
	*tchttp.BaseRequest
	
	// 身份提供商名称
	ProviderId *string `json:"ProviderId,omitempty" name:"ProviderId"`

	// IdP签发的OIDC令牌
	WebIdentityToken *string `json:"WebIdentityToken,omitempty" name:"WebIdentityToken"`

	// 角色访问描述名
	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`

	// 会话名称
	RoleSessionName *string `json:"RoleSessionName,omitempty" name:"RoleSessionName"`

	// 指定临时证书的有效期，单位：秒，默认 7200 秒，最长可设定有效期为 43200 秒
	DurationSeconds *int64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`
}

func (r *AssumeRoleWithWebIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssumeRoleWithWebIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProviderId")
	delete(f, "WebIdentityToken")
	delete(f, "RoleArn")
	delete(f, "RoleSessionName")
	delete(f, "DurationSeconds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssumeRoleWithWebIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssumeRoleWithWebIdentityResponseParams struct {
	// 临时密钥过期时间(时间戳)
	ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 临时密钥过期时间
	Expiration *string `json:"Expiration,omitempty" name:"Expiration"`

	// 临时密钥
	Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AssumeRoleWithWebIdentityResponse struct {
	*tchttp.BaseResponse
	Response *AssumeRoleWithWebIdentityResponseParams `json:"Response"`
}

func (r *AssumeRoleWithWebIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssumeRoleWithWebIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Credentials struct {
	// token。token长度和绑定的策略有关，最长不超过4096字节。
	Token *string `json:"Token,omitempty" name:"Token"`

	// 临时证书密钥ID。最长不超过1024字节。
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// 临时证书密钥Key。最长不超过1024字节。
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`
}

// Predefined struct for user
type GetCallerIdentityRequestParams struct {

}

type GetCallerIdentityRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetCallerIdentityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCallerIdentityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCallerIdentityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCallerIdentityResponseParams struct {
	// 当前调用者ARN。
	Arn *string `json:"Arn,omitempty" name:"Arn"`

	// 当前调用者所属主账号Uin。
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`

	// 身份标识。
	// 1. 调用者是云账号时，返回的是当前账号Uin
	// 2. 调用者是角色时，返回的是roleId:roleSessionName
	// 3. 调用者是联合身份时，返回的是uin:federatedUserName
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 密钥所属账号Uin。
	// 1. 调用者是云账号，返回的当前账号Uin
	// 2, 调用者是角色，返回的申请角色密钥的账号Uin
	PrincipalId *string `json:"PrincipalId,omitempty" name:"PrincipalId"`

	// 身份类型。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetCallerIdentityResponse struct {
	*tchttp.BaseResponse
	Response *GetCallerIdentityResponseParams `json:"Response"`
}

func (r *GetCallerIdentityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCallerIdentityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFederationTokenRequestParams struct {
	// 您可以自定义调用方英文名称，由字母组成。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 授予该临时证书权限的CAM策略
	// 注意：
	// 1、策略语法参照[ CAM 策略语法](https://cloud.tencent.com/document/product/598/10603)。
	// 2、策略中不能包含 principal 元素。
	// 3、该参数需要做urlencode。
	Policy *string `json:"Policy,omitempty" name:"Policy"`

	// 指定临时证书的有效期，单位：秒，默认1800秒，主账号最长可设定有效期为7200秒，子账号最长可设定有效期为129600秒。
	DurationSeconds *uint64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`
}

type GetFederationTokenRequest struct {
	*tchttp.BaseRequest
	
	// 您可以自定义调用方英文名称，由字母组成。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 授予该临时证书权限的CAM策略
	// 注意：
	// 1、策略语法参照[ CAM 策略语法](https://cloud.tencent.com/document/product/598/10603)。
	// 2、策略中不能包含 principal 元素。
	// 3、该参数需要做urlencode。
	Policy *string `json:"Policy,omitempty" name:"Policy"`

	// 指定临时证书的有效期，单位：秒，默认1800秒，主账号最长可设定有效期为7200秒，子账号最长可设定有效期为129600秒。
	DurationSeconds *uint64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`
}

func (r *GetFederationTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFederationTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Policy")
	delete(f, "DurationSeconds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFederationTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFederationTokenResponseParams struct {
	// 临时证书
	Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

	// 临时证书有效的时间，返回 Unix 时间戳，精确到秒
	ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

	// 证书有效的时间，以 iso8601 格式的 UTC 时间表示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Expiration *string `json:"Expiration,omitempty" name:"Expiration"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetFederationTokenResponse struct {
	*tchttp.BaseResponse
	Response *GetFederationTokenResponseParams `json:"Response"`
}

func (r *GetFederationTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFederationTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryApiKeyRequestParams struct {
	// 待查询的账号(不填默认查当前账号)
	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

type QueryApiKeyRequest struct {
	*tchttp.BaseRequest
	
	// 待查询的账号(不填默认查当前账号)
	TargetUin *uint64 `json:"TargetUin,omitempty" name:"TargetUin"`
}

func (r *QueryApiKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryApiKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryApiKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryApiKeyResponseParams struct {
	// 密钥ID列表
	IdKeys []*ApiKey `json:"IdKeys,omitempty" name:"IdKeys"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type QueryApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *QueryApiKeyResponseParams `json:"Response"`
}

func (r *QueryApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryApiKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签键，最长128个字符，区分大小写。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签值，最长256个字符，区分大小写。
	Value *string `json:"Value,omitempty" name:"Value"`
}