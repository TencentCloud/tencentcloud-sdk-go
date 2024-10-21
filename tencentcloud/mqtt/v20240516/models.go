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

package v20240516

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AuthorizationPolicyItem struct {
	// 规则id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 规则名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 规则语法版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *int64 `json:"Version,omitnil,omitempty" name:"Version"`

	// 越小越优先
	// 注意：此字段可能返回 null，表示取不到有效值。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// allow/deny
	// 注意：此字段可能返回 null，表示取不到有效值。
	Effect *string `json:"Effect,omitnil,omitempty" name:"Effect"`

	// connect、pub、sub
	// 注意：此字段可能返回 null，表示取不到有效值。
	Actions *string `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resources *string `json:"Resources,omitnil,omitempty" name:"Resources"`

	// client
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 0，1，2
	// 注意：此字段可能返回 null，表示取不到有效值。
	Qos *string `json:"Qos,omitnil,omitempty" name:"Qos"`

	// 1：表示匹配retain消息
	// 2：表示匹配非retain消息
	// 3：表示匹配retain和非retain消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Retain *int64 `json:"Retain,omitnil,omitempty" name:"Retain"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 1713164969433
	CreatedTime *int64 `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 1713164969433
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type AuthorizationPolicyPriority struct {
	// 策略id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 优先级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

// Predefined struct for user
type CreateAuthorizationPolicyRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略版本
	PolicyVersion *int64 `json:"PolicyVersion,omitnil,omitempty" name:"PolicyVersion"`

	// 策略优先级，越小越优先
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// allow、deny
	Effect *string `json:"Effect,omitnil,omitempty" name:"Effect"`

	// connect、pub、sub
	Actions *string `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 1,匹配保留消息；2,匹配非保留消息，3.匹配所有消息
	Retain *int64 `json:"Retain,omitnil,omitempty" name:"Retain"`

	// 0、1、2
	Qos *string `json:"Qos,omitnil,omitempty" name:"Qos"`

	// 资源
	Resources *string `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 客户端
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// ip
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateAuthorizationPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略版本
	PolicyVersion *int64 `json:"PolicyVersion,omitnil,omitempty" name:"PolicyVersion"`

	// 策略优先级，越小越优先
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// allow、deny
	Effect *string `json:"Effect,omitnil,omitempty" name:"Effect"`

	// connect、pub、sub
	Actions *string `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 1,匹配保留消息；2,匹配非保留消息，3.匹配所有消息
	Retain *int64 `json:"Retain,omitnil,omitempty" name:"Retain"`

	// 0、1、2
	Qos *string `json:"Qos,omitnil,omitempty" name:"Qos"`

	// 资源
	Resources *string `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 客户端
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// ip
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreateAuthorizationPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuthorizationPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "PolicyName")
	delete(f, "PolicyVersion")
	delete(f, "Priority")
	delete(f, "Effect")
	delete(f, "Actions")
	delete(f, "Retain")
	delete(f, "Qos")
	delete(f, "Resources")
	delete(f, "Username")
	delete(f, "ClientId")
	delete(f, "Ip")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuthorizationPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuthorizationPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAuthorizationPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuthorizationPolicyResponseParams `json:"Response"`
}

func (r *CreateAuthorizationPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuthorizationPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateJWKSAuthenticatorRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// jwks端点
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// jwks刷新间隔,单位：秒
	RefreshInterval *int64 `json:"RefreshInterval,omitnil,omitempty" name:"RefreshInterval"`

	// jwks文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 认证器是否开启：open-启用；close-关闭
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 设备连接时传递jwt的key；username-使用用户名字段传递；password-使用密码字段传递
	From *string `json:"From,omitnil,omitempty" name:"From"`
}

type CreateJWKSAuthenticatorRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// jwks端点
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// jwks刷新间隔,单位：秒
	RefreshInterval *int64 `json:"RefreshInterval,omitnil,omitempty" name:"RefreshInterval"`

	// jwks文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 认证器是否开启：open-启用；close-关闭
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 设备连接时传递jwt的key；username-使用用户名字段传递；password-使用密码字段传递
	From *string `json:"From,omitnil,omitempty" name:"From"`
}

func (r *CreateJWKSAuthenticatorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateJWKSAuthenticatorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Endpoint")
	delete(f, "RefreshInterval")
	delete(f, "Text")
	delete(f, "Status")
	delete(f, "Remark")
	delete(f, "From")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateJWKSAuthenticatorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateJWKSAuthenticatorResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateJWKSAuthenticatorResponse struct {
	*tchttp.BaseResponse
	Response *CreateJWKSAuthenticatorResponseParams `json:"Response"`
}

func (r *CreateJWKSAuthenticatorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateJWKSAuthenticatorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateJWTAuthenticatorRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 算法：hmac-based，public-key
	Algorithm *string `json:"Algorithm,omitnil,omitempty" name:"Algorithm"`

	// 设备连接时传递jwt的key；username-使用用户名字段传递；password-使用密码字段传递
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 密码
	Secret *string `json:"Secret,omitnil,omitempty" name:"Secret"`

	// 公钥
	PublicKey *string `json:"PublicKey,omitnil,omitempty" name:"PublicKey"`

	// 认证器是否开启：open-启用；close-关闭
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateJWTAuthenticatorRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 算法：hmac-based，public-key
	Algorithm *string `json:"Algorithm,omitnil,omitempty" name:"Algorithm"`

	// 设备连接时传递jwt的key；username-使用用户名字段传递；password-使用密码字段传递
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 密码
	Secret *string `json:"Secret,omitnil,omitempty" name:"Secret"`

	// 公钥
	PublicKey *string `json:"PublicKey,omitnil,omitempty" name:"PublicKey"`

	// 认证器是否开启：open-启用；close-关闭
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *CreateJWTAuthenticatorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateJWTAuthenticatorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Algorithm")
	delete(f, "From")
	delete(f, "Secret")
	delete(f, "PublicKey")
	delete(f, "Status")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateJWTAuthenticatorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateJWTAuthenticatorResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateJWTAuthenticatorResponse struct {
	*tchttp.BaseResponse
	Response *CreateJWTAuthenticatorResponseParams `json:"Response"`
}

func (r *CreateJWTAuthenticatorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateJWTAuthenticatorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
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
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTopicResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

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
type DeleteAuthenticatorRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 认证器类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DeleteAuthenticatorRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 认证器类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DeleteAuthenticatorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuthenticatorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuthenticatorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuthenticatorResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAuthenticatorResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuthenticatorResponseParams `json:"Response"`
}

func (r *DeleteAuthenticatorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuthenticatorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuthorizationPolicyRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略规则id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteAuthorizationPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略规则id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DeleteAuthorizationPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuthorizationPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuthorizationPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuthorizationPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAuthorizationPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuthorizationPolicyResponseParams `json:"Response"`
}

func (r *DeleteAuthorizationPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuthorizationPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
}

type DeleteTopicRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
}

func (r *DeleteTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTopicResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTopicResponseParams `json:"Response"`
}

func (r *DeleteTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuthenticatorRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 认证器类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeAuthenticatorRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 认证器类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeAuthenticatorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuthenticatorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuthenticatorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuthenticatorResponseParams struct {
	// 集群认证器列表
	Authenticators []*MQTTAuthenticatorItem `json:"Authenticators,omitnil,omitempty" name:"Authenticators"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuthenticatorResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuthenticatorResponseParams `json:"Response"`
}

func (r *DescribeAuthenticatorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuthenticatorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuthorizationPoliciesRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeAuthorizationPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeAuthorizationPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuthorizationPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuthorizationPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuthorizationPoliciesResponseParams struct {
	// 规则
	Data []*AuthorizationPolicyItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuthorizationPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuthorizationPoliciesResponseParams `json:"Response"`
}

func (r *DescribeAuthorizationPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuthorizationPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceListRequestParams struct {
	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 标签过滤器
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

type DescribeInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 标签过滤器
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

func (r *DescribeInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceListResponseParams struct {
	// 查询总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例列表
	Data []*MQTTInstanceItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceListResponseParams `json:"Response"`
}

func (r *DescribeInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceResponseParams struct {
	// 实例类型，
	// EXPERIMENT 体验版
	// BASIC 基础版
	// PRO  专业版
	// PLATINUM 铂金版
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 主题数量
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// 实例最大主题数量
	TopicNumLimit *int64 `json:"TopicNumLimit,omitnil,omitempty" name:"TopicNumLimit"`

	// TPS限流值
	TpsLimit *int64 `json:"TpsLimit,omitnil,omitempty" name:"TpsLimit"`

	// 创建时间，秒为单位
	CreatedTime *int64 `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 实例状态
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 实例规格
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// 单客户端最大订阅数
	MaxSubscriptionPerClient *int64 `json:"MaxSubscriptionPerClient,omitnil,omitempty" name:"MaxSubscriptionPerClient"`

	// 授权规则条数
	AuthorizationPolicyLimit *int64 `json:"AuthorizationPolicyLimit,omitnil,omitempty" name:"AuthorizationPolicyLimit"`

	// 客户端数量上限
	ClientNumLimit *int64 `json:"ClientNumLimit,omitnil,omitempty" name:"ClientNumLimit"`

	// 客户端证书注册方式：JITP，API
	DeviceCertificateProvisionType *string `json:"DeviceCertificateProvisionType,omitnil,omitempty" name:"DeviceCertificateProvisionType"`

	// 自动注册设备证书时是否自动激活
	AutomaticActivation *bool `json:"AutomaticActivation,omitnil,omitempty" name:"AutomaticActivation"`

	// 是否自动续费
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 计费模式， POSTPAID，按量计费 PREPAID，包年包月
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 到期时间，秒为单位
	ExpiryTime *int64 `json:"ExpiryTime,omitnil,omitempty" name:"ExpiryTime"`

	// 预销毁时间
	DestroyTime *int64 `json:"DestroyTime,omitnil,omitempty" name:"DestroyTime"`

	//     TLS,单向认证
	//     mTLS,双向认证
	//     BYOC;一机一证
	X509Mode *string `json:"X509Mode,omitnil,omitempty" name:"X509Mode"`

	// 最大Ca配额
	MaxCaNum *int64 `json:"MaxCaNum,omitnil,omitempty" name:"MaxCaNum"`

	// 证书注册码
	RegistrationCode *string `json:"RegistrationCode,omitnil,omitempty" name:"RegistrationCode"`

	// 集群最大订阅数
	MaxSubscription *int64 `json:"MaxSubscription,omitnil,omitempty" name:"MaxSubscription"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceResponseParams `json:"Response"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicListRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeTopicListRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询条件列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeTopicListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicListResponseParams struct {
	// 查询总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 主题列表
	Data []*MQTTTopicItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicListResponseParams `json:"Response"`
}

func (r *DescribeTopicListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
}

type DescribeTopicRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
}

func (r *DescribeTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 创建时间，秒为单位
	CreatedTime *int64 `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopicResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicResponseParams `json:"Response"`
}

func (r *DescribeTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 过滤条件名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤条件的值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type MQTTAuthenticatorItem struct {
	// 认证器类型: JWT：JWT认证器 JWKS：JWKS认证器 BYOC：一端一证认证器
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 认证器配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 认证器状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type MQTTInstanceItem struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 实例类型，
	// EXPERIMENT，体验版
	// BASIC，基础版
	// PRO，专业版
	// PLATINUM，铂金版
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例状态，
	// RUNNING, 运行中
	// MAINTAINING，维护中
	// ABNORMAL，异常
	// OVERDUE，欠费
	// DESTROYED，已删除
	// CREATING，创建中
	// MODIFYING，变配中
	// CREATE_FAILURE，创建失败
	// MODIFY_FAILURE，变配失败
	// DELETING，删除中
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 实例主题数上限
	TopicNumLimit *int64 `json:"TopicNumLimit,omitnil,omitempty" name:"TopicNumLimit"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 主题数量
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// 商品规格
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// 弹性TPS限流值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TpsLimit *int64 `json:"TpsLimit,omitnil,omitempty" name:"TpsLimit"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 单客户端最大订阅数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxSubscriptionPerClient *int64 `json:"MaxSubscriptionPerClient,omitnil,omitempty" name:"MaxSubscriptionPerClient"`

	// 客户端连接数上线
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientNumLimit *int64 `json:"ClientNumLimit,omitnil,omitempty" name:"ClientNumLimit"`

	// 是否自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 计费模式， POSTPAID，按量计费 PREPAID，包年包月
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 到期时间，秒为单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiryTime *int64 `json:"ExpiryTime,omitnil,omitempty" name:"ExpiryTime"`

	// 预销毁时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestroyTime *int64 `json:"DestroyTime,omitnil,omitempty" name:"DestroyTime"`

	// 授权规则条数限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorizationPolicyLimit *int64 `json:"AuthorizationPolicyLimit,omitnil,omitempty" name:"AuthorizationPolicyLimit"`

	// 最大ca配额
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxCaNum *int64 `json:"MaxCaNum,omitnil,omitempty" name:"MaxCaNum"`

	// 最大订阅数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxSubscription *int64 `json:"MaxSubscription,omitnil,omitempty" name:"MaxSubscription"`
}

type MQTTTopicItem struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 主题描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

// Predefined struct for user
type ModifyAuthorizationPolicyRequestParams struct {
	// 策略
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略版本
	PolicyVersion *int64 `json:"PolicyVersion,omitnil,omitempty" name:"PolicyVersion"`

	// 策略优先级，越小越优先
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// allow、deny
	Effect *string `json:"Effect,omitnil,omitempty" name:"Effect"`

	// connect、pub、sub
	Actions *string `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 资源
	Resources *string `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 1.匹配保留消息；2.匹配非保留消息；3.匹配所有消息
	Retain *int64 `json:"Retain,omitnil,omitempty" name:"Retain"`

	// 客户端
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// ip
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 0、1、2
	Qos *string `json:"Qos,omitnil,omitempty" name:"Qos"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyAuthorizationPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 策略版本
	PolicyVersion *int64 `json:"PolicyVersion,omitnil,omitempty" name:"PolicyVersion"`

	// 策略优先级，越小越优先
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// allow、deny
	Effect *string `json:"Effect,omitnil,omitempty" name:"Effect"`

	// connect、pub、sub
	Actions *string `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 资源
	Resources *string `json:"Resources,omitnil,omitempty" name:"Resources"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 1.匹配保留消息；2.匹配非保留消息；3.匹配所有消息
	Retain *int64 `json:"Retain,omitnil,omitempty" name:"Retain"`

	// 客户端
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// ip
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 0、1、2
	Qos *string `json:"Qos,omitnil,omitempty" name:"Qos"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyAuthorizationPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuthorizationPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "InstanceId")
	delete(f, "PolicyName")
	delete(f, "PolicyVersion")
	delete(f, "Priority")
	delete(f, "Effect")
	delete(f, "Actions")
	delete(f, "Resources")
	delete(f, "Username")
	delete(f, "Retain")
	delete(f, "ClientId")
	delete(f, "Ip")
	delete(f, "Qos")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuthorizationPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuthorizationPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAuthorizationPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAuthorizationPolicyResponseParams `json:"Response"`
}

func (r *ModifyAuthorizationPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuthorizationPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyJWKSAuthenticatorRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 端点
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 认证器状态：open-启用；close-关闭
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 刷新时间
	RefreshInterval *int64 `json:"RefreshInterval,omitnil,omitempty" name:"RefreshInterval"`

	// JSKS文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 设备连接时传递jwt的key；username-使用用户名字段传递；password-使用密码字段传递
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyJWKSAuthenticatorRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 端点
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 认证器状态：open-启用；close-关闭
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 刷新时间
	RefreshInterval *int64 `json:"RefreshInterval,omitnil,omitempty" name:"RefreshInterval"`

	// JSKS文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 设备连接时传递jwt的key；username-使用用户名字段传递；password-使用密码字段传递
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyJWKSAuthenticatorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyJWKSAuthenticatorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Endpoint")
	delete(f, "Status")
	delete(f, "RefreshInterval")
	delete(f, "Text")
	delete(f, "From")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyJWKSAuthenticatorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyJWKSAuthenticatorResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyJWKSAuthenticatorResponse struct {
	*tchttp.BaseResponse
	Response *ModifyJWKSAuthenticatorResponseParams `json:"Response"`
}

func (r *ModifyJWKSAuthenticatorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyJWKSAuthenticatorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyJWTAuthenticatorRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 算法：hmac-based，public-key
	Algorithm *string `json:"Algorithm,omitnil,omitempty" name:"Algorithm"`

	// 设备连接时传递jwt的key；username-使用用户名字段传递；password-使用密码字段传递
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 密码
	Secret *string `json:"Secret,omitnil,omitempty" name:"Secret"`

	// 公钥
	PublicKey *string `json:"PublicKey,omitnil,omitempty" name:"PublicKey"`

	// JSKS文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyJWTAuthenticatorRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 算法：hmac-based，public-key
	Algorithm *string `json:"Algorithm,omitnil,omitempty" name:"Algorithm"`

	// 设备连接时传递jwt的key；username-使用用户名字段传递；password-使用密码字段传递
	From *string `json:"From,omitnil,omitempty" name:"From"`

	// 密码
	Secret *string `json:"Secret,omitnil,omitempty" name:"Secret"`

	// 公钥
	PublicKey *string `json:"PublicKey,omitnil,omitempty" name:"PublicKey"`

	// JSKS文本
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyJWTAuthenticatorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyJWTAuthenticatorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Algorithm")
	delete(f, "From")
	delete(f, "Secret")
	delete(f, "PublicKey")
	delete(f, "Text")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyJWTAuthenticatorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyJWTAuthenticatorResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyJWTAuthenticatorResponse struct {
	*tchttp.BaseResponse
	Response *ModifyJWTAuthenticatorResponseParams `json:"Response"`
}

func (r *ModifyJWTAuthenticatorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyJWTAuthenticatorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyTopicRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

func (r *ModifyTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTopicResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTopicResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTopicResponseParams `json:"Response"`
}

func (r *ModifyTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterDeviceCertificateRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设备证书
	DeviceCertificate *string `json:"DeviceCertificate,omitnil,omitempty" name:"DeviceCertificate"`

	// 关联的CA证书SN
	CaSn *string `json:"CaSn,omitnil,omitempty" name:"CaSn"`

	// 客户端ID
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 证书格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	//     ACTIVE,//激活     INACTIVE,//未激活     REVOKED,//吊销     PENDING_ACTIVATION,//注册待激活
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type RegisterDeviceCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设备证书
	DeviceCertificate *string `json:"DeviceCertificate,omitnil,omitempty" name:"DeviceCertificate"`

	// 关联的CA证书SN
	CaSn *string `json:"CaSn,omitnil,omitempty" name:"CaSn"`

	// 客户端ID
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 证书格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	//     ACTIVE,//激活     INACTIVE,//未激活     REVOKED,//吊销     PENDING_ACTIVATION,//注册待激活
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *RegisterDeviceCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterDeviceCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DeviceCertificate")
	delete(f, "CaSn")
	delete(f, "ClientId")
	delete(f, "Format")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterDeviceCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterDeviceCertificateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RegisterDeviceCertificateResponse struct {
	*tchttp.BaseResponse
	Response *RegisterDeviceCertificateResponseParams `json:"Response"`
}

func (r *RegisterDeviceCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterDeviceCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagFilter struct {
	// 标签键名称
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签键名称
	TagValues []*string `json:"TagValues,omitnil,omitempty" name:"TagValues"`
}

// Predefined struct for user
type UpdateAuthorizationPolicyPriorityRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略ID和优先级
	Priorities []*AuthorizationPolicyPriority `json:"Priorities,omitnil,omitempty" name:"Priorities"`
}

type UpdateAuthorizationPolicyPriorityRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 策略ID和优先级
	Priorities []*AuthorizationPolicyPriority `json:"Priorities,omitnil,omitempty" name:"Priorities"`
}

func (r *UpdateAuthorizationPolicyPriorityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAuthorizationPolicyPriorityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Priorities")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAuthorizationPolicyPriorityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAuthorizationPolicyPriorityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAuthorizationPolicyPriorityResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAuthorizationPolicyPriorityResponseParams `json:"Response"`
}

func (r *UpdateAuthorizationPolicyPriorityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAuthorizationPolicyPriorityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}