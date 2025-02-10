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

// Predefined struct for user
type ActivateDeviceCertificateRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设备证书序列号
	DeviceCertificateSn *string `json:"DeviceCertificateSn,omitnil,omitempty" name:"DeviceCertificateSn"`
}

type ActivateDeviceCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设备证书序列号
	DeviceCertificateSn *string `json:"DeviceCertificateSn,omitnil,omitempty" name:"DeviceCertificateSn"`
}

func (r *ActivateDeviceCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateDeviceCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DeviceCertificateSn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ActivateDeviceCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActivateDeviceCertificateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ActivateDeviceCertificateResponse struct {
	*tchttp.BaseResponse
	Response *ActivateDeviceCertificateResponseParams `json:"Response"`
}

func (r *ActivateDeviceCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateDeviceCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuthorizationPolicyItem struct {
	// 规则ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 集群ID
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

	// IP地址
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

	// IP地址
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
type CreateInsPublicEndpointRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 带宽,单位Mbps
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网访问规则
	Rules []*PublicAccessRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type CreateInsPublicEndpointRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 带宽,单位Mbps
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网访问规则
	Rules []*PublicAccessRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *CreateInsPublicEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInsPublicEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Bandwidth")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInsPublicEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInsPublicEndpointResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInsPublicEndpointResponse struct {
	*tchttp.BaseResponse
	Response *CreateInsPublicEndpointResponseParams `json:"Response"`
}

func (r *CreateInsPublicEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInsPublicEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceRequestParams struct {
	// 实例类型，
	// BASIC 基础版
	// PRO  专业版
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 商品规格，可用规格可通过接口DescribeProductSKUList查询
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 实例绑定的VPC信息
	VpcList []*VpcInfo `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// 是否开启公网
	EnablePublic *bool `json:"EnablePublic,omitnil,omitempty" name:"EnablePublic"`

	// 公网带宽（单位：兆）
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网访问白名单
	IpRules []*IpRule `json:"IpRules,omitnil,omitempty" name:"IpRules"`

	// 是否自动续费（0: 不自动续费；1: 自动续费）
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 购买时长（单位：月）
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 付费模式（0: 后付费；1: 预付费）
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

type CreateInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例类型，
	// BASIC 基础版
	// PRO  专业版
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 商品规格，可用规格可通过接口DescribeProductSKUList查询
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 标签列表
	TagList []*Tag `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 实例绑定的VPC信息
	VpcList []*VpcInfo `json:"VpcList,omitnil,omitempty" name:"VpcList"`

	// 是否开启公网
	EnablePublic *bool `json:"EnablePublic,omitnil,omitempty" name:"EnablePublic"`

	// 公网带宽（单位：兆）
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网访问白名单
	IpRules []*IpRule `json:"IpRules,omitnil,omitempty" name:"IpRules"`

	// 是否自动续费（0: 不自动续费；1: 自动续费）
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 购买时长（单位：月）
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 付费模式（0: 后付费；1: 预付费）
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`
}

func (r *CreateInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceType")
	delete(f, "Name")
	delete(f, "SkuCode")
	delete(f, "Remark")
	delete(f, "TagList")
	delete(f, "VpcList")
	delete(f, "EnablePublic")
	delete(f, "Bandwidth")
	delete(f, "IpRules")
	delete(f, "RenewFlag")
	delete(f, "TimeSpan")
	delete(f, "PayMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInstanceResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateInstanceResponseParams `json:"Response"`
}

func (r *CreateInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInstanceResponse) FromJsonString(s string) error {
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

	// 设备连接时传递jwt的key；
	// username-使用用户名字段传递；
	// password-使用密码字段传递
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

	// 设备连接时传递jwt的key；
	// username-使用用户名字段传递；
	// password-使用密码字段传递
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
type CreateUserRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 密码，该字段为空时候则后端会默认生成
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 密码，该字段为空时候则后端会默认生成
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
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
	delete(f, "InstanceId")
	delete(f, "Username")
	delete(f, "Password")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserResponseParams struct {
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
type DeactivateDeviceCertificateRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设备证书序列号
	DeviceCertificateSn *string `json:"DeviceCertificateSn,omitnil,omitempty" name:"DeviceCertificateSn"`
}

type DeactivateDeviceCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设备证书序列号
	DeviceCertificateSn *string `json:"DeviceCertificateSn,omitnil,omitempty" name:"DeviceCertificateSn"`
}

func (r *DeactivateDeviceCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeactivateDeviceCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DeviceCertificateSn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeactivateDeviceCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeactivateDeviceCertificateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeactivateDeviceCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DeactivateDeviceCertificateResponseParams `json:"Response"`
}

func (r *DeactivateDeviceCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeactivateDeviceCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuthenticatorRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 认证器类型:
	// JWT：JWT认证器
	// JWKS：JWKS认证器
	// BYOC：一端一证认证器
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DeleteAuthenticatorRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 认证器类型:
	// JWT：JWT认证器
	// JWKS：JWKS认证器
	// BYOC：一端一证认证器
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
type DeleteDeviceCertificateRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设备证书序列号
	DeviceCertificateSn *string `json:"DeviceCertificateSn,omitnil,omitempty" name:"DeviceCertificateSn"`
}

type DeleteDeviceCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设备证书序列号
	DeviceCertificateSn *string `json:"DeviceCertificateSn,omitnil,omitempty" name:"DeviceCertificateSn"`
}

func (r *DeleteDeviceCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DeviceCertificateSn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDeviceCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDeviceCertificateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDeviceCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDeviceCertificateResponseParams `json:"Response"`
}

func (r *DeleteDeviceCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDeviceCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInsPublicEndpointRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteInsPublicEndpointRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteInsPublicEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInsPublicEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInsPublicEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInsPublicEndpointResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInsPublicEndpointResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInsPublicEndpointResponseParams `json:"Response"`
}

func (r *DeleteInsPublicEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInsPublicEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInstanceResponseParams `json:"Response"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInstanceResponse) FromJsonString(s string) error {
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
type DeleteUserRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`
}

func (r *DeleteUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Username")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserResponseParams `json:"Response"`
}

func (r *DeleteUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuthenticatorRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 认证器类型: JWT：JWT认证器 JWKS：JWKS认证器 HTTP:HTTP认证器
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeAuthenticatorRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 认证器类型: JWT：JWT认证器 JWKS：JWKS认证器 HTTP:HTTP认证器
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
type DescribeDeviceCertificateRequestParams struct {
	// 设备证书sn
	DeviceCertificateSn *string `json:"DeviceCertificateSn,omitnil,omitempty" name:"DeviceCertificateSn"`

	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeDeviceCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 设备证书sn
	DeviceCertificateSn *string `json:"DeviceCertificateSn,omitnil,omitempty" name:"DeviceCertificateSn"`

	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeDeviceCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeviceCertificateSn")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceCertificateResponseParams struct {
	// 创建时间
	CreatedTime *int64 `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 上次更新时间
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 证书失效日期
	NotAfterTime *int64 `json:"NotAfterTime,omitnil,omitempty" name:"NotAfterTime"`

	// 上次激活时间
	LastActivationTime *int64 `json:"LastActivationTime,omitnil,omitempty" name:"LastActivationTime"`

	// 上次取消激活时间
	LastInactivationTime *int64 `json:"LastInactivationTime,omitnil,omitempty" name:"LastInactivationTime"`

	// 证书状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Ca证书序列号
	CaSn *string `json:"CaSn,omitnil,omitempty" name:"CaSn"`

	// 设备证书序列号
	DeviceCertificateSn *string `json:"DeviceCertificateSn,omitnil,omitempty" name:"DeviceCertificateSn"`

	// 设备证书内容
	DeviceCertificate *string `json:"DeviceCertificate,omitnil,omitempty" name:"DeviceCertificate"`

	// 设备证书common name
	DeviceCertificateCn *string `json:"DeviceCertificateCn,omitnil,omitempty" name:"DeviceCertificateCn"`

	// 证书格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 客户端id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	//     API, 手动注册   
	//     JITP 自动注册
	CertificateSource *string `json:"CertificateSource,omitnil,omitempty" name:"CertificateSource"`

	// 证书生效开始时间
	NotBeforeTime *int64 `json:"NotBeforeTime,omitnil,omitempty" name:"NotBeforeTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceCertificateResponseParams `json:"Response"`
}

func (r *DescribeDeviceCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceCertificatesRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 过滤器支持ClientId、CaSn、DeviceCertificateSn、Status搜索
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// CREATE_TIME_DESC, 创建时间降序
	//     CREATE_TIME_ASC,创建时间升序
	//     UPDATE_TIME_DESC,更新时间降序
	//     UPDATE_TIME_ASC,更新时间升序
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
}

type DescribeDeviceCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 过滤器支持ClientId、CaSn、DeviceCertificateSn、Status搜索
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分页limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// CREATE_TIME_DESC, 创建时间降序
	//     CREATE_TIME_ASC,创建时间升序
	//     UPDATE_TIME_DESC,更新时间降序
	//     UPDATE_TIME_ASC,更新时间升序
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`
}

func (r *DescribeDeviceCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceCertificatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "OrderBy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDeviceCertificatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDeviceCertificatesResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 设备证书
	Data []*DeviceCertificateItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDeviceCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDeviceCertificatesResponseParams `json:"Response"`
}

func (r *DescribeDeviceCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDeviceCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInsPublicEndpointsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInsPublicEndpointsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInsPublicEndpointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInsPublicEndpointsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInsPublicEndpointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInsPublicEndpointsResponseParams struct {
	// 接入点
	Endpoints []*MQTTEndpointItem `json:"Endpoints,omitnil,omitempty" name:"Endpoints"`

	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 带宽
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网访问规则
	Rules []*PublicAccessRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 公网状态：
	//     NORMAL-正常
	//     CLOSING-关闭中
	//     MODIFYING-修改中
	//     CREATING-开启中
	//     CLOSE-关闭
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInsPublicEndpointsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInsPublicEndpointsResponseParams `json:"Response"`
}

func (r *DescribeInsPublicEndpointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInsPublicEndpointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceListRequestParams struct {
	// 查询条件列表,支持以下子弹
	// InstanceName：集群名模糊搜索
	// InstanceId：集群id精确搜索
	// InstanceStatus：集群状态搜索
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
	
	// 查询条件列表,支持以下子弹
	// InstanceName：集群名模糊搜索
	// InstanceId：集群id精确搜索
	// InstanceStatus：集群状态搜索
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
	// 实例类型
	// BASIC 基础版
	// PRO  专业版
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

	// 客户端证书注册方式：
	// JITP：自动注册
	// API：通过API手动注册
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

	// TLS,单向认证    mTLS,双向认证    BYOC;一机一证
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

	// 查询条件列表:
	// 支持TopicName模糊查询
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

	// 查询条件列表:
	// 支持TopicName模糊查询
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

// Predefined struct for user
type DescribeUserListRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询条件列表支持字段
	// Username：Username模糊查询
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeUserListRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询条件列表支持字段
	// Username：Username模糊查询
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
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
	delete(f, "InstanceId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserListResponseParams struct {
	// 查询总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 角色信息列表
	Data []*MQTTUserItem `json:"Data,omitnil,omitempty" name:"Data"`

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

type DeviceCertificateItem struct {
	// 客户端id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 设备证书
	DeviceCertificate *string `json:"DeviceCertificate,omitnil,omitempty" name:"DeviceCertificate"`

	// 设备证书Sn
	DeviceCertificateSn *string `json:"DeviceCertificateSn,omitnil,omitempty" name:"DeviceCertificateSn"`

	// 设备证书Cn
	DeviceCertificateCn *string `json:"DeviceCertificateCn,omitnil,omitempty" name:"DeviceCertificateCn"`

	// 签发ca的序列号
	CaSn *string `json:"CaSn,omitnil,omitempty" name:"CaSn"`

	// 证书格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 证书状态
	//     ACTIVE,//激活
	//     INACTIVE,//未激活
	//     REVOKED,//吊销
	//     PENDING_ACTIVATION,//注册待激活
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 上次激活时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastActivationTime *int64 `json:"LastActivationTime,omitnil,omitempty" name:"LastActivationTime"`

	// 上次取消激活时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastInactivationTime *int64 `json:"LastInactivationTime,omitnil,omitempty" name:"LastInactivationTime"`

	// 创建时间
	CreatedTime *int64 `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 预销毁时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 证书来源：
	// API, 手动注册   
	// JITP 自动注册
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificateSource *string `json:"CertificateSource,omitnil,omitempty" name:"CertificateSource"`

	// 证书失效日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotAfterTime *int64 `json:"NotAfterTime,omitnil,omitempty" name:"NotAfterTime"`

	// 证书生效开始日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotBeforeTime *int64 `json:"NotBeforeTime,omitnil,omitempty" name:"NotBeforeTime"`
}

type Filter struct {
	// 过滤条件名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤条件的值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type IpRule struct {
	// IP地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 是否允许放行
	Allow *bool `json:"Allow,omitnil,omitempty" name:"Allow"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
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

type MQTTEndpointItem struct {
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 接入点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// vpc信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 主机
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 接入点ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`
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

type MQTTUserItem struct {
	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 创建时间，秒为单位
	CreatedTime *int64 `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 修改时间，秒为单位
	ModifiedTime *int64 `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`
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

	// IP
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

	// IP
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
type ModifyInsPublicEndpointRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 带宽，单位：Mbps
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网访问规则
	Rules []*PublicAccessRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type ModifyInsPublicEndpointRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 带宽，单位：Mbps
	Bandwidth *int64 `json:"Bandwidth,omitnil,omitempty" name:"Bandwidth"`

	// 公网访问规则
	Rules []*PublicAccessRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *ModifyInsPublicEndpointRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInsPublicEndpointRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Bandwidth")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInsPublicEndpointRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInsPublicEndpointResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInsPublicEndpointResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInsPublicEndpointResponseParams `json:"Response"`
}

func (r *ModifyInsPublicEndpointResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInsPublicEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要修改实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 要修改的备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 要变更的配置规格
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// 客户端证书注册方式：
	// JITP：自动注册
	// API：手动通过API注册
	DeviceCertificateProvisionType *string `json:"DeviceCertificateProvisionType,omitnil,omitempty" name:"DeviceCertificateProvisionType"`

	// 自动注册证书是否自动激活
	AutomaticActivation *bool `json:"AutomaticActivation,omitnil,omitempty" name:"AutomaticActivation"`
}

type ModifyInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要修改实例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 要修改的备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 要变更的配置规格
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// 客户端证书注册方式：
	// JITP：自动注册
	// API：手动通过API注册
	DeviceCertificateProvisionType *string `json:"DeviceCertificateProvisionType,omitnil,omitempty" name:"DeviceCertificateProvisionType"`

	// 自动注册证书是否自动激活
	AutomaticActivation *bool `json:"AutomaticActivation,omitnil,omitempty" name:"AutomaticActivation"`
}

func (r *ModifyInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "Remark")
	delete(f, "SkuCode")
	delete(f, "DeviceCertificateProvisionType")
	delete(f, "AutomaticActivation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceResponseParams `json:"Response"`
}

func (r *ModifyInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceResponse) FromJsonString(s string) error {
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

	// 设备连接时传递jwt的key；
	// username-使用用户名字段传递；
	// password-使用密码字段传递
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

	// 设备连接时传递jwt的key；
	// username-使用用户名字段传递；
	// password-使用密码字段传递
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
type ModifyUserRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type ModifyUserRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
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
	delete(f, "InstanceId")
	delete(f, "Username")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserResponseParams struct {
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

type PublicAccessRule struct {
	// ip网段信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpRule *string `json:"IpRule,omitnil,omitempty" name:"IpRule"`

	// 允许或者拒绝
	// 注意：此字段可能返回 null，表示取不到有效值。
	Allow *bool `json:"Allow,omitnil,omitempty" name:"Allow"`

	// 备注信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
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

// Predefined struct for user
type RevokedDeviceCertificateRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设备证书序列号
	DeviceCertificateSn *string `json:"DeviceCertificateSn,omitnil,omitempty" name:"DeviceCertificateSn"`
}

type RevokedDeviceCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 设备证书序列号
	DeviceCertificateSn *string `json:"DeviceCertificateSn,omitnil,omitempty" name:"DeviceCertificateSn"`
}

func (r *RevokedDeviceCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevokedDeviceCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "DeviceCertificateSn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RevokedDeviceCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RevokedDeviceCertificateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RevokedDeviceCertificateResponse struct {
	*tchttp.BaseResponse
	Response *RevokedDeviceCertificateResponseParams `json:"Response"`
}

func (r *RevokedDeviceCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevokedDeviceCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {
	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
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

type VpcInfo struct {
	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}