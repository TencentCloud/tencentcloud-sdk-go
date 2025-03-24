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
type ActivateCaCertificateRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 证书序列号
	CaSn *string `json:"CaSn,omitnil,omitempty" name:"CaSn"`
}

type ActivateCaCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 证书序列号
	CaSn *string `json:"CaSn,omitnil,omitempty" name:"CaSn"`
}

func (r *ActivateCaCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateCaCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CaSn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ActivateCaCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ActivateCaCertificateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ActivateCaCertificateResponse struct {
	*tchttp.BaseResponse
	Response *ActivateCaCertificateResponseParams `json:"Response"`
}

func (r *ActivateCaCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ActivateCaCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

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

// Predefined struct for user
type ApplyRegistrationCodeRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type ApplyRegistrationCodeRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *ApplyRegistrationCodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyRegistrationCodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyRegistrationCodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyRegistrationCodeResponseParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 注册码
	RegistrationCode *string `json:"RegistrationCode,omitnil,omitempty" name:"RegistrationCode"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ApplyRegistrationCodeResponse struct {
	*tchttp.BaseResponse
	Response *ApplyRegistrationCodeResponseParams `json:"Response"`
}

func (r *ApplyRegistrationCodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyRegistrationCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuthorizationPolicyItem struct {
	// 规则ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 规则名
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// 规则语法版本
	Version *int64 `json:"Version,omitnil,omitempty" name:"Version"`

	// 越小越优先
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// allow/deny
	Effect *string `json:"Effect,omitnil,omitempty" name:"Effect"`

	// connect、pub、sub
	Actions *string `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 资源
	Resources *string `json:"Resources,omitnil,omitempty" name:"Resources"`

	// client
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 用户
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// IP地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 0，1，2
	Qos *string `json:"Qos,omitnil,omitempty" name:"Qos"`

	// 1：表示匹配retain消息
	// 2：表示匹配非retain消息
	// 3：表示匹配retain和非retain消息
	Retain *int64 `json:"Retain,omitnil,omitempty" name:"Retain"`

	// 描述
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 1713164969433
	CreatedTime *int64 `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 1713164969433
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type AuthorizationPolicyPriority struct {
	// 策略id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 优先级
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`
}

type BodyItem struct {
	// body key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// body key
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type CaCertificateItem struct {
	// common name
	CaCn *string `json:"CaCn,omitnil,omitempty" name:"CaCn"`

	// 证书内容
	CaCertificate *string `json:"CaCertificate,omitnil,omitempty" name:"CaCertificate"`

	// 证书序列号
	CaSn *string `json:"CaSn,omitnil,omitempty" name:"CaSn"`

	// 证书格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 验证证书内容
	VerificationCertificate *string `json:"VerificationCertificate,omitnil,omitempty" name:"VerificationCertificate"`

	// ca状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 上次激活时间
	LastActivationTime *int64 `json:"LastActivationTime,omitnil,omitempty" name:"LastActivationTime"`

	// 创建时间
	CreatedTime *int64 `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 预销毁时间
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 上次去激活时间
	LastInactivationTime *int64 `json:"LastInactivationTime,omitnil,omitempty" name:"LastInactivationTime"`

	// Ca证书颁发者CN
	CaIssuerCn *string `json:"CaIssuerCn,omitnil,omitempty" name:"CaIssuerCn"`

	// 生效时间
	NotBeforeTime *int64 `json:"NotBeforeTime,omitnil,omitempty" name:"NotBeforeTime"`

	// 失效时间
	NotAfterTime *int64 `json:"NotAfterTime,omitnil,omitempty" name:"NotAfterTime"`
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
type CreateHttpAuthenticatorRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// jwks端点
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 最大并发连接数，默认8，范围：1-20
	Concurrency *int64 `json:"Concurrency,omitnil,omitempty" name:"Concurrency"`

	// 网络请求方法 Get 或 Post，默认post
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 认证器是否开启：open-启用；close-关闭
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 连接超时时间，单位：秒，范围：1-30
	ConnectTimeout *int64 `json:"ConnectTimeout,omitnil,omitempty" name:"ConnectTimeout"`

	// 请求超时时间，单位：秒，范围：1-30
	ReadTimeout *int64 `json:"ReadTimeout,omitnil,omitempty" name:"ReadTimeout"`

	// 转发请求header
	Header []*HeaderItem `json:"Header,omitnil,omitempty" name:"Header"`

	// 转发请求body
	Body []*BodyItem `json:"Body,omitnil,omitempty" name:"Body"`
}

type CreateHttpAuthenticatorRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// jwks端点
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 最大并发连接数，默认8，范围：1-20
	Concurrency *int64 `json:"Concurrency,omitnil,omitempty" name:"Concurrency"`

	// 网络请求方法 Get 或 Post，默认post
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 认证器是否开启：open-启用；close-关闭
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 连接超时时间，单位：秒，范围：1-30
	ConnectTimeout *int64 `json:"ConnectTimeout,omitnil,omitempty" name:"ConnectTimeout"`

	// 请求超时时间，单位：秒，范围：1-30
	ReadTimeout *int64 `json:"ReadTimeout,omitnil,omitempty" name:"ReadTimeout"`

	// 转发请求header
	Header []*HeaderItem `json:"Header,omitnil,omitempty" name:"Header"`

	// 转发请求body
	Body []*BodyItem `json:"Body,omitnil,omitempty" name:"Body"`
}

func (r *CreateHttpAuthenticatorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHttpAuthenticatorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Endpoint")
	delete(f, "Concurrency")
	delete(f, "Method")
	delete(f, "Status")
	delete(f, "Remark")
	delete(f, "ConnectTimeout")
	delete(f, "ReadTimeout")
	delete(f, "Header")
	delete(f, "Body")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHttpAuthenticatorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHttpAuthenticatorResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateHttpAuthenticatorResponse struct {
	*tchttp.BaseResponse
	Response *CreateHttpAuthenticatorResponseParams `json:"Response"`
}

func (r *CreateHttpAuthenticatorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHttpAuthenticatorResponse) FromJsonString(s string) error {
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
type DeactivateCaCertificateRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 证书序列号
	CaSn *string `json:"CaSn,omitnil,omitempty" name:"CaSn"`
}

type DeactivateCaCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 证书序列号
	CaSn *string `json:"CaSn,omitnil,omitempty" name:"CaSn"`
}

func (r *DeactivateCaCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeactivateCaCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CaSn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeactivateCaCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeactivateCaCertificateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeactivateCaCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DeactivateCaCertificateResponseParams `json:"Response"`
}

func (r *DeactivateCaCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeactivateCaCertificateResponse) FromJsonString(s string) error {
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
type DeleteCaCertificateRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 证书序列号
	CaSn *string `json:"CaSn,omitnil,omitempty" name:"CaSn"`
}

type DeleteCaCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 证书序列号
	CaSn *string `json:"CaSn,omitnil,omitempty" name:"CaSn"`
}

func (r *DeleteCaCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCaCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CaSn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCaCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCaCertificateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCaCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCaCertificateResponseParams `json:"Response"`
}

func (r *DeleteCaCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCaCertificateResponse) FromJsonString(s string) error {
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
type DescribeCaCertificateRequestParams struct {
	// ca证书sn
	CaSn *string `json:"CaSn,omitnil,omitempty" name:"CaSn"`

	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeCaCertificateRequest struct {
	*tchttp.BaseRequest
	
	// ca证书sn
	CaSn *string `json:"CaSn,omitnil,omitempty" name:"CaSn"`

	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeCaCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CaSn")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCaCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaCertificateResponseParams struct {
	// 创建时间
	CreatedTime *int64 `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 上次更新时间
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 失效日期
	NotAfterTime *int64 `json:"NotAfterTime,omitnil,omitempty" name:"NotAfterTime"`

	// 上次激活时间
	LastActivationTime *int64 `json:"LastActivationTime,omitnil,omitempty" name:"LastActivationTime"`

	// 上次吊销时间
	LastInactivationTime *int64 `json:"LastInactivationTime,omitnil,omitempty" name:"LastInactivationTime"`

	// 证书状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 证书序列号
	CaSn *string `json:"CaSn,omitnil,omitempty" name:"CaSn"`

	// common name
	CaCn *string `json:"CaCn,omitnil,omitempty" name:"CaCn"`

	// 证书内容
	CaCertificate *string `json:"CaCertificate,omitnil,omitempty" name:"CaCertificate"`

	// 证书格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// Ca证书颁发者CN
	CaIssuerCn *string `json:"CaIssuerCn,omitnil,omitempty" name:"CaIssuerCn"`

	// 生效开始时间
	NotBeforeTime *int64 `json:"NotBeforeTime,omitnil,omitempty" name:"NotBeforeTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCaCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCaCertificateResponseParams `json:"Response"`
}

func (r *DescribeCaCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaCertificatesRequestParams struct {
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeCaCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeCaCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaCertificatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCaCertificatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCaCertificatesResponseParams struct {
	// ca证书列表
	Data []*CaCertificateItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCaCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCaCertificatesResponseParams `json:"Response"`
}

func (r *DescribeCaCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCaCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientListRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 客户端名
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 客户端数量限制,最大1024，默认1024
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`
}

type DescribeClientListRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 客户端名
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 客户端数量限制,最大1024，默认1024
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`
}

func (r *DescribeClientListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ClientId")
	delete(f, "Number")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClientListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClientListResponseParams struct {
	// 客户端列表
	Clients []*MQTTClientInfo `json:"Clients,omitnil,omitempty" name:"Clients"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClientListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClientListResponseParams `json:"Response"`
}

func (r *DescribeClientListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClientListResponse) FromJsonString(s string) error {
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
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	//     API, 手动注册   
	//     JITP 自动注册
	CertificateSource *string `json:"CertificateSource,omitnil,omitempty" name:"CertificateSource"`

	// 证书生效开始时间
	NotBeforeTime *int64 `json:"NotBeforeTime,omitnil,omitempty" name:"NotBeforeTime"`

	// 组织单位
	OrganizationalUnit *string `json:"OrganizationalUnit,omitnil,omitempty" name:"OrganizationalUnit"`

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
type DescribeInsVPCEndpointsRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInsVPCEndpointsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInsVPCEndpointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInsVPCEndpointsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInsVPCEndpointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInsVPCEndpointsResponseParams struct {
	// 接入点
	Endpoints []*MQTTEndpointItem `json:"Endpoints,omitnil,omitempty" name:"Endpoints"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInsVPCEndpointsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInsVPCEndpointsResponseParams `json:"Response"`
}

func (r *DescribeInsVPCEndpointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInsVPCEndpointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceListRequestParams struct {
	// 查询条件列表,支持以下字段
	// InstanceName：集群名模糊搜索
	// InstanceId：集群id精确搜索
	// InstanceStatus：集群状态搜索（RUNNING-运行中，CREATING-创建中，MODIFYING-变配中，DELETING-删除中）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20，最大100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 标签过滤器
	TagFilters []*TagFilter `json:"TagFilters,omitnil,omitempty" name:"TagFilters"`
}

type DescribeInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 查询条件列表,支持以下字段
	// InstanceName：集群名模糊搜索
	// InstanceId：集群id精确搜索
	// InstanceStatus：集群状态搜索（RUNNING-运行中，CREATING-创建中，MODIFYING-变配中，DELETING-删除中）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 查询起始位置，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量，默认20，最大100
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
type DescribeMessageListRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 请求任务id
	TaskRequestId *string `json:"TaskRequestId,omitnil,omitempty" name:"TaskRequestId"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeMessageListRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 请求任务id
	TaskRequestId *string `json:"TaskRequestId,omitnil,omitempty" name:"TaskRequestId"`

	// 查询起始位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询结果限制数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeMessageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Topic")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskRequestId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMessageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMessageListResponseParams struct {
	// 查询总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 消息记录列表
	Data []*MQTTMessageItem `json:"Data,omitnil,omitempty" name:"Data"`

	// 请求任务id
	TaskRequestId *string `json:"TaskRequestId,omitnil,omitempty" name:"TaskRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMessageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMessageListResponseParams `json:"Response"`
}

func (r *DescribeMessageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMessageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductSKUListRequestParams struct {

}

type DescribeProductSKUListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeProductSKUListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductSKUListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProductSKUListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProductSKUListResponseParams struct {
	// 查询总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// mqtt商品配置信息
	MQTTProductSkuList []*ProductSkuItem `json:"MQTTProductSkuList,omitnil,omitempty" name:"MQTTProductSkuList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProductSKUListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProductSKUListResponseParams `json:"Response"`
}

func (r *DescribeProductSKUListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProductSKUListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSharedSubscriptionLagRequestParams struct {
	// 集群id	
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 共享订阅表达式
	SharedSubscription *string `json:"SharedSubscription,omitnil,omitempty" name:"SharedSubscription"`
}

type DescribeSharedSubscriptionLagRequest struct {
	*tchttp.BaseRequest
	
	// 集群id	
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 共享订阅表达式
	SharedSubscription *string `json:"SharedSubscription,omitnil,omitempty" name:"SharedSubscription"`
}

func (r *DescribeSharedSubscriptionLagRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSharedSubscriptionLagRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SharedSubscription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSharedSubscriptionLagRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSharedSubscriptionLagResponseParams struct {
	// 堆积值
	Lag *int64 `json:"Lag,omitnil,omitempty" name:"Lag"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSharedSubscriptionLagResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSharedSubscriptionLagResponseParams `json:"Response"`
}

func (r *DescribeSharedSubscriptionLagResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSharedSubscriptionLagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicListRequestParams struct {
	// 实例ID，从 [DescribeInstanceList](https://cloud.tencent.com/document/api/1778/111029)接口或控制台获得。
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
	
	// 实例ID，从 [DescribeInstanceList](https://cloud.tencent.com/document/api/1778/111029)接口或控制台获得。
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

	// 组织单位
	OrganizationalUnit *string `json:"OrganizationalUnit,omitnil,omitempty" name:"OrganizationalUnit"`

	// 上次激活时间
	LastActivationTime *int64 `json:"LastActivationTime,omitnil,omitempty" name:"LastActivationTime"`

	// 上次取消激活时间
	LastInactivationTime *int64 `json:"LastInactivationTime,omitnil,omitempty" name:"LastInactivationTime"`

	// 创建时间
	CreatedTime *int64 `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 预销毁时间
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 证书来源：
	// API, 手动注册   
	// JITP 自动注册
	CertificateSource *string `json:"CertificateSource,omitnil,omitempty" name:"CertificateSource"`

	// 证书失效日期
	NotAfterTime *int64 `json:"NotAfterTime,omitnil,omitempty" name:"NotAfterTime"`

	// 证书生效开始日期
	NotBeforeTime *int64 `json:"NotBeforeTime,omitnil,omitempty" name:"NotBeforeTime"`
}

type Filter struct {
	// 过滤条件名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤条件的值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type HeaderItem struct {
	// header key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// header value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type IpRule struct {
	// IP地址
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 是否允许放行
	Allow *bool `json:"Allow,omitnil,omitempty" name:"Allow"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type MQTTAuthenticatorItem struct {
	// 认证器类型: JWT：JWT认证器 JWKS：JWKS认证器 BYOC：一端一证认证器
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 认证器配置
	Config *string `json:"Config,omitnil,omitempty" name:"Config"`

	// 认证器状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

type MQTTClientInfo struct {
	// 客户端唯一标识
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 客户端网络地址
	ClientAddress *string `json:"ClientAddress,omitnil,omitempty" name:"ClientAddress"`

	// MQTT 协议版本，4 表示 MQTT 3.1.1
	ProtocolVersion *int64 `json:"ProtocolVersion,omitnil,omitempty" name:"ProtocolVersion"`

	// 保持连接时间，单位：秒
	Keepalive *int64 `json:"Keepalive,omitnil,omitempty" name:"Keepalive"`

	// 连接状态，CONNECTED 已连接，DISCONNECTED 未连接
	ConnectionStatus *string `json:"ConnectionStatus,omitnil,omitempty" name:"ConnectionStatus"`

	// 客户端创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 上次建立连接时间
	ConnectTime *int64 `json:"ConnectTime,omitnil,omitempty" name:"ConnectTime"`

	// 上次断开连接时间，仅对持久会话（cleanSession=false）并且客户端当前未连接时有意义
	DisconnectTime *int64 `json:"DisconnectTime,omitnil,omitempty" name:"DisconnectTime"`

	// 客户端的订阅列表
	MQTTClientSubscriptions []*MQTTClientSubscription `json:"MQTTClientSubscriptions,omitnil,omitempty" name:"MQTTClientSubscriptions"`
}

type MQTTClientSubscription struct {
	// topic 订阅
	TopicFilter *string `json:"TopicFilter,omitnil,omitempty" name:"TopicFilter"`

	// 服务质量等级
	Qos *int64 `json:"Qos,omitnil,omitempty" name:"Qos"`

	// 堆积数量
	Lag *int64 `json:"Lag,omitnil,omitempty" name:"Lag"`

	// 投递未确认数量
	Inflight *int64 `json:"Inflight,omitnil,omitempty" name:"Inflight"`
}

type MQTTEndpointItem struct {
	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 接入点
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// vpc信息
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网信息
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 主机
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 接入点ip
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`
}

type MQTTInstanceItem struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 实例类型
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
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 主题数量
	TopicNum *int64 `json:"TopicNum,omitnil,omitempty" name:"TopicNum"`

	// 商品规格
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// 弹性TPS限流值
	TpsLimit *int64 `json:"TpsLimit,omitnil,omitempty" name:"TpsLimit"`

	// 创建时间，毫秒级时间戳
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 单客户端最大订阅数量
	MaxSubscriptionPerClient *int64 `json:"MaxSubscriptionPerClient,omitnil,omitempty" name:"MaxSubscriptionPerClient"`

	// 客户端连接数上线
	ClientNumLimit *int64 `json:"ClientNumLimit,omitnil,omitempty" name:"ClientNumLimit"`

	// 是否自动续费。仅包年包月就去那生效。
	// 1:自动续费
	// 0:非自动续费
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 计费模式， POSTPAID，按量计费 PREPAID，包年包月
	PayMode *string `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 到期时间，毫秒级时间戳
	ExpiryTime *int64 `json:"ExpiryTime,omitnil,omitempty" name:"ExpiryTime"`

	// 预销毁时间，毫秒级时间戳
	DestroyTime *int64 `json:"DestroyTime,omitnil,omitempty" name:"DestroyTime"`

	// 授权规则条数限制
	AuthorizationPolicyLimit *int64 `json:"AuthorizationPolicyLimit,omitnil,omitempty" name:"AuthorizationPolicyLimit"`

	// 最大ca配额
	MaxCaNum *int64 `json:"MaxCaNum,omitnil,omitempty" name:"MaxCaNum"`

	// 最大订阅数
	MaxSubscription *int64 `json:"MaxSubscription,omitnil,omitempty" name:"MaxSubscription"`
}

type MQTTMessageItem struct {
	// 消息ID
	MsgId *string `json:"MsgId,omitnil,omitempty" name:"MsgId"`

	// 消息tag
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 消息key
	Keys *string `json:"Keys,omitnil,omitempty" name:"Keys"`

	// 客户端地址	
	ProducerAddr *string `json:"ProducerAddr,omitnil,omitempty" name:"ProducerAddr"`

	// 消息发送时间	
	ProduceTime *string `json:"ProduceTime,omitnil,omitempty" name:"ProduceTime"`

	// 死信重发次数	
	DeadLetterResendTimes *int64 `json:"DeadLetterResendTimes,omitnil,omitempty" name:"DeadLetterResendTimes"`

	// 死信重发成功次数
	DeadLetterResendSuccessTimes *int64 `json:"DeadLetterResendSuccessTimes,omitnil,omitempty" name:"DeadLetterResendSuccessTimes"`

	// 子topic
	SubTopic *string `json:"SubTopic,omitnil,omitempty" name:"SubTopic"`

	// 消息质量等级
	Qos *string `json:"Qos,omitnil,omitempty" name:"Qos"`
}

type MQTTTopicItem struct {
	// 实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 主题名称
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 主题描述
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
type ModifyHttpAuthenticatorRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 端点
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 认证器状态：open-启用；close-关闭
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 最大并发连接数，默认8，范围：1-20
	Concurrency *int64 `json:"Concurrency,omitnil,omitempty" name:"Concurrency"`

	// 连接超时时间，单位：秒，范围：1-30
	ConnectTimeout *int64 `json:"ConnectTimeout,omitnil,omitempty" name:"ConnectTimeout"`

	// 请求超时时间，单位：秒，范围：1-30
	ReadTimeout *int64 `json:"ReadTimeout,omitnil,omitempty" name:"ReadTimeout"`

	// 说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 请求方法，GET 或者 POST
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 请求header
	Header []*HeaderItem `json:"Header,omitnil,omitempty" name:"Header"`

	// 请求body
	Body []*BodyItem `json:"Body,omitnil,omitempty" name:"Body"`
}

type ModifyHttpAuthenticatorRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 端点
	Endpoint *string `json:"Endpoint,omitnil,omitempty" name:"Endpoint"`

	// 认证器状态：open-启用；close-关闭
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 最大并发连接数，默认8，范围：1-20
	Concurrency *int64 `json:"Concurrency,omitnil,omitempty" name:"Concurrency"`

	// 连接超时时间，单位：秒，范围：1-30
	ConnectTimeout *int64 `json:"ConnectTimeout,omitnil,omitempty" name:"ConnectTimeout"`

	// 请求超时时间，单位：秒，范围：1-30
	ReadTimeout *int64 `json:"ReadTimeout,omitnil,omitempty" name:"ReadTimeout"`

	// 说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 请求方法，GET 或者 POST
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 请求header
	Header []*HeaderItem `json:"Header,omitnil,omitempty" name:"Header"`

	// 请求body
	Body []*BodyItem `json:"Body,omitnil,omitempty" name:"Body"`
}

func (r *ModifyHttpAuthenticatorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHttpAuthenticatorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Endpoint")
	delete(f, "Status")
	delete(f, "Concurrency")
	delete(f, "ConnectTimeout")
	delete(f, "ReadTimeout")
	delete(f, "Remark")
	delete(f, "Method")
	delete(f, "Header")
	delete(f, "Body")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHttpAuthenticatorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHttpAuthenticatorResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyHttpAuthenticatorResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHttpAuthenticatorResponseParams `json:"Response"`
}

func (r *ModifyHttpAuthenticatorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHttpAuthenticatorResponse) FromJsonString(s string) error {
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
type ModifyInstanceCertBindingRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务端证书id
	SSLServerCertId *string `json:"SSLServerCertId,omitnil,omitempty" name:"SSLServerCertId"`

	// CA证书id
	SSLCaCertId *string `json:"SSLCaCertId,omitnil,omitempty" name:"SSLCaCertId"`

	// 加密通信方式
	// TLS：单向证书认证
	// mTLS：双向证书认证
	// BYOC：一设备一证书认证
	X509Mode *string `json:"X509Mode,omitnil,omitempty" name:"X509Mode"`

	// 设备证书注册类型：
	// JITP，自动注册；
	// MANUAL 手动注册
	DeviceCertificateProvisionType *string `json:"DeviceCertificateProvisionType,omitnil,omitempty" name:"DeviceCertificateProvisionType"`

	// 是否自动激活，默认为false
	AutomaticActivation *bool `json:"AutomaticActivation,omitnil,omitempty" name:"AutomaticActivation"`
}

type ModifyInstanceCertBindingRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务端证书id
	SSLServerCertId *string `json:"SSLServerCertId,omitnil,omitempty" name:"SSLServerCertId"`

	// CA证书id
	SSLCaCertId *string `json:"SSLCaCertId,omitnil,omitempty" name:"SSLCaCertId"`

	// 加密通信方式
	// TLS：单向证书认证
	// mTLS：双向证书认证
	// BYOC：一设备一证书认证
	X509Mode *string `json:"X509Mode,omitnil,omitempty" name:"X509Mode"`

	// 设备证书注册类型：
	// JITP，自动注册；
	// MANUAL 手动注册
	DeviceCertificateProvisionType *string `json:"DeviceCertificateProvisionType,omitnil,omitempty" name:"DeviceCertificateProvisionType"`

	// 是否自动激活，默认为false
	AutomaticActivation *bool `json:"AutomaticActivation,omitnil,omitempty" name:"AutomaticActivation"`
}

func (r *ModifyInstanceCertBindingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceCertBindingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SSLServerCertId")
	delete(f, "SSLCaCertId")
	delete(f, "X509Mode")
	delete(f, "DeviceCertificateProvisionType")
	delete(f, "AutomaticActivation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyInstanceCertBindingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyInstanceCertBindingResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyInstanceCertBindingResponse struct {
	*tchttp.BaseResponse
	Response *ModifyInstanceCertBindingResponseParams `json:"Response"`
}

func (r *ModifyInstanceCertBindingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyInstanceCertBindingResponse) FromJsonString(s string) error {
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

type PriceTag struct {
	// 计价名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 计价类别
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 计费项标签
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 步长
	Step *int64 `json:"Step,omitnil,omitempty" name:"Step"`
}

type ProductSkuItem struct {
	// 规格类型
	// BASIC：基础版
	// PRO ：专业版
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 规格代码
	SkuCode *string `json:"SkuCode,omitnil,omitempty" name:"SkuCode"`

	// 是否售卖
	OnSale *bool `json:"OnSale,omitnil,omitempty" name:"OnSale"`

	// topic num限制
	TopicNumLimit *int64 `json:"TopicNumLimit,omitnil,omitempty" name:"TopicNumLimit"`

	// tps
	TpsLimit *int64 `json:"TpsLimit,omitnil,omitempty" name:"TpsLimit"`

	// 客户端连接数
	ClientNumLimit *int64 `json:"ClientNumLimit,omitnil,omitempty" name:"ClientNumLimit"`

	// 单客户端最大订阅数
	MaxSubscriptionPerClient *int64 `json:"MaxSubscriptionPerClient,omitnil,omitempty" name:"MaxSubscriptionPerClient"`

	// 授权规则条数
	AuthorizationPolicyLimit *int64 `json:"AuthorizationPolicyLimit,omitnil,omitempty" name:"AuthorizationPolicyLimit"`

	// 计费项信息
	PriceTags []*PriceTag `json:"PriceTags,omitnil,omitempty" name:"PriceTags"`
}

type PublicAccessRule struct {
	// ip网段信息
	IpRule *string `json:"IpRule,omitnil,omitempty" name:"IpRule"`

	// 允许或者拒绝
	Allow *bool `json:"Allow,omitnil,omitempty" name:"Allow"`

	// 备注信息
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`
}

// Predefined struct for user
type PublishMessageRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消息 payload，需要按 encoding 指定的编码方式进行编码
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// 消息目的主题，该参数与 TargetClientId 二选一
	TargetTopic *string `json:"TargetTopic,omitnil,omitempty" name:"TargetTopic"`

	// 消息目的客户端 ID，该参数与 TargetTopic 二选一
	TargetClientId *string `json:"TargetClientId,omitnil,omitempty" name:"TargetClientId"`

	// 消息 payload 编码，可选 plain 或 base64，默认为 plain（即不编码）
	Encoding *string `json:"Encoding,omitnil,omitempty" name:"Encoding"`

	// 消息的服务质量等级，默认为 1
	Qos *int64 `json:"Qos,omitnil,omitempty" name:"Qos"`

	// 是否为保留消息，默认为 false，且仅支持发布到主题的消息设置为 true
	Retain *bool `json:"Retain,omitnil,omitempty" name:"Retain"`
}

type PublishMessageRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 消息 payload，需要按 encoding 指定的编码方式进行编码
	Payload *string `json:"Payload,omitnil,omitempty" name:"Payload"`

	// 消息目的主题，该参数与 TargetClientId 二选一
	TargetTopic *string `json:"TargetTopic,omitnil,omitempty" name:"TargetTopic"`

	// 消息目的客户端 ID，该参数与 TargetTopic 二选一
	TargetClientId *string `json:"TargetClientId,omitnil,omitempty" name:"TargetClientId"`

	// 消息 payload 编码，可选 plain 或 base64，默认为 plain（即不编码）
	Encoding *string `json:"Encoding,omitnil,omitempty" name:"Encoding"`

	// 消息的服务质量等级，默认为 1
	Qos *int64 `json:"Qos,omitnil,omitempty" name:"Qos"`

	// 是否为保留消息，默认为 false，且仅支持发布到主题的消息设置为 true
	Retain *bool `json:"Retain,omitnil,omitempty" name:"Retain"`
}

func (r *PublishMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Payload")
	delete(f, "TargetTopic")
	delete(f, "TargetClientId")
	delete(f, "Encoding")
	delete(f, "Qos")
	delete(f, "Retain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PublishMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PublishMessageResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PublishMessageResponse struct {
	*tchttp.BaseResponse
	Response *PublishMessageResponseParams `json:"Response"`
}

func (r *PublishMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterCaCertificateRequestParams struct {
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// CA证书
	CaCertificate *string `json:"CaCertificate,omitnil,omitempty" name:"CaCertificate"`

	// 验证证书
	VerificationCertificate *string `json:"VerificationCertificate,omitnil,omitempty" name:"VerificationCertificate"`

	// 证书格式，不传默认PEM格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 证书状态，不传默认ACTIVE状态
	//     ACTIVE,//激活
	//     INACTIVE,//未激活
	//     REVOKED,//吊销
	//     PENDING_ACTIVATION,//注册待激活
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type RegisterCaCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// CA证书
	CaCertificate *string `json:"CaCertificate,omitnil,omitempty" name:"CaCertificate"`

	// 验证证书
	VerificationCertificate *string `json:"VerificationCertificate,omitnil,omitempty" name:"VerificationCertificate"`

	// 证书格式，不传默认PEM格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 证书状态，不传默认ACTIVE状态
	//     ACTIVE,//激活
	//     INACTIVE,//未激活
	//     REVOKED,//吊销
	//     PENDING_ACTIVATION,//注册待激活
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *RegisterCaCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterCaCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "CaCertificate")
	delete(f, "VerificationCertificate")
	delete(f, "Format")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterCaCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterCaCertificateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RegisterCaCertificateResponse struct {
	*tchttp.BaseResponse
	Response *RegisterCaCertificateResponseParams `json:"Response"`
}

func (r *RegisterCaCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterCaCertificateResponse) FromJsonString(s string) error {
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
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
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