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

type AssumeRoleRequest struct {
	*tchttp.BaseRequest

	// 角色的资源描述。例如：qcs::cam::uin/12345678:role/4611686018427397919、qcs::cam::uin/12345678:roleName/testRoleName
	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`

	// 临时会话名称，由用户自定义名称
	RoleSessionName *string `json:"RoleSessionName,omitempty" name:"RoleSessionName"`

	// 指定临时证书的有效期，单位：秒，默认 7200 秒，最长可设定有效期为 43200 秒
	DurationSeconds *uint64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`

	// 策略描述
	// 注意：
	// 1、policy 需要做 urlencode（如果通过 GET 方法请求云 API，发送请求前，所有参数都需要按照[云 API 规范](https://cloud.tencent.com/document/api/598/33159#1.-.E6.8B.BC.E6.8E.A5.E8.A7.84.E8.8C.83.E8.AF.B7.E6.B1.82.E4.B8.B2)再 urlencode 一次）。
	// 2、策略语法参照[ CAM 策略语法](https://cloud.tencent.com/document/product/598/10603)。
	// 3、策略中不能包含 principal 元素。
	Policy *string `json:"Policy,omitempty" name:"Policy"`
}

func (r *AssumeRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssumeRoleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssumeRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 临时安全证书
		Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

		// 证书无效的时间，返回 Unix 时间戳，精确到秒
		ExpiredTime *int64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

		// 证书无效的时间，以 iso8601 格式的 UTC 时间表示
		Expiration *string `json:"Expiration,omitempty" name:"Expiration"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssumeRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssumeRoleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

	// 指定临时证书的有效期，单位：秒，默认 7200 秒，最长可设定有效期为 7200 秒
	DurationSeconds *uint64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`
}

func (r *AssumeRoleWithSAMLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssumeRoleWithSAMLRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssumeRoleWithSAMLResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 对象里面包含 Token，TmpSecretId，TmpSecretKey 三元组
		Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

		// 证书无效的时间，返回 Unix 时间戳，精确到秒
		ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

		// 证书无效的时间，以 ISO8601 格式的 UTC 时间表示
		Expiration *string `json:"Expiration,omitempty" name:"Expiration"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssumeRoleWithSAMLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssumeRoleWithSAMLResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Credentials struct {

	// token
	Token *string `json:"Token,omitempty" name:"Token"`

	// 临时证书密钥ID
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// 临时证书密钥Key
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`
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

	// 指定临时证书的有效期，单位：秒，默认1800秒，最长可设定有效期为7200秒。
	DurationSeconds *uint64 `json:"DurationSeconds,omitempty" name:"DurationSeconds"`
}

func (r *GetFederationTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetFederationTokenRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetFederationTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 临时证书
		Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

		// 临时证书有效的时间，返回 Unix 时间戳，精确到秒
		ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

		// 证书有效的时间，以 iso8601 格式的 UTC 时间表示
	// 注意：此字段可能返回 null，表示取不到有效值。
		Expiration *string `json:"Expiration,omitempty" name:"Expiration"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetFederationTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetFederationTokenResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
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

func (r *QueryApiKeyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QueryApiKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 密钥ID列表
		IdKeys []*ApiKey `json:"IdKeys,omitempty" name:"IdKeys" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryApiKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *QueryApiKeyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
