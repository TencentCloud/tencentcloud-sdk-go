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

package v20180317

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AssociateTargetGroupsRequestParams struct {
	// 绑定的关系数组。
	Associations []*TargetGroupAssociation `json:"Associations,omitempty" name:"Associations"`
}

type AssociateTargetGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 绑定的关系数组。
	Associations []*TargetGroupAssociation `json:"Associations,omitempty" name:"Associations"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type AssociationItem struct {
	// 关联到的负载均衡ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 关联到的监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 关联到的转发规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 关联到的监听器协议类型，如HTTP,TCP,
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 关联到的监听器端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 关联到的转发规则域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 关联到的转发规则URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 负载均衡名称
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 监听器名称
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`
}

// Predefined struct for user
type AutoRewriteRequestParams struct {
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// HTTPS:443监听器的ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// HTTPS:443监听器下需要重定向的域名，若不填，则对HTTPS:443监听器下的所有域名都设置重定向。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 重定向状态码，可取值301,302,307。
	RewriteCodes []*int64 `json:"RewriteCodes,omitempty" name:"RewriteCodes"`

	// 重定向是否携带匹配的URL。
	TakeUrls []*bool `json:"TakeUrls,omitempty" name:"TakeUrls"`
}

type AutoRewriteRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// HTTPS:443监听器的ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// HTTPS:443监听器下需要重定向的域名，若不填，则对HTTPS:443监听器下的所有域名都设置重定向。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 重定向状态码，可取值301,302,307。
	RewriteCodes []*int64 `json:"RewriteCodes,omitempty" name:"RewriteCodes"`

	// 重定向是否携带匹配的URL。
	TakeUrls []*bool `json:"TakeUrls,omitempty" name:"TakeUrls"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type Backend struct {
	// 后端服务的类型，可取：CVM、ENI
	Type *string `json:"Type,omitempty" name:"Type"`

	// 后端服务的唯一 ID，如 ins-abcd1234
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 后端服务的监听端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 后端服务的转发权重，取值范围：[0, 100]，默认为 10。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 后端服务的外网 IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`

	// 后端服务的内网 IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// 后端服务的实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 后端服务被绑定的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegisteredTime *string `json:"RegisteredTime,omitempty" name:"RegisteredTime"`

	// 弹性网卡唯一ID，如 eni-1234abcd
	// 注意：此字段可能返回 null，表示取不到有效值。
	EniId *string `json:"EniId,omitempty" name:"EniId"`
}

type BasicTargetGroupInfo struct {
	// 目标组ID
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// 目标组名称
	TargetGroupName *string `json:"TargetGroupName,omitempty" name:"TargetGroupName"`
}

// Predefined struct for user
type BatchDeregisterTargetsRequestParams struct {
	// 负载均衡ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 解绑目标。
	Targets []*BatchTarget `json:"Targets,omitempty" name:"Targets"`
}

type BatchDeregisterTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 解绑目标。
	Targets []*BatchTarget `json:"Targets,omitempty" name:"Targets"`
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
	FailListenerIdSet []*string `json:"FailListenerIdSet,omitempty" name:"FailListenerIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type BatchModifyTargetWeightRequestParams struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 要批量修改权重的列表。
	ModifyList []*RsWeightRule `json:"ModifyList,omitempty" name:"ModifyList"`
}

type BatchModifyTargetWeightRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 要批量修改权重的列表。
	ModifyList []*RsWeightRule `json:"ModifyList,omitempty" name:"ModifyList"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 绑定目标。
	Targets []*BatchTarget `json:"Targets,omitempty" name:"Targets"`
}

type BatchRegisterTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 绑定目标。
	Targets []*BatchTarget `json:"Targets,omitempty" name:"Targets"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailListenerIdSet []*string `json:"FailListenerIdSet,omitempty" name:"FailListenerIdSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 绑定端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 子机 ID。表示绑定主网卡主 IP。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 绑定 IP 时需要传入此参数，支持弹性网卡的 IP 和其他内网 IP，如果是弹性网卡则必须先绑定至CVM，然后才能绑定到负载均衡实例。
	// 注意：参数 InstanceId、EniIp 只能传入一个且必须传入一个。如果绑定双栈IPV6子机，必须传该参数。
	EniIp *string `json:"EniIp,omitempty" name:"EniIp"`

	// 子机权重，范围[0, 100]。绑定时如果不存在，则默认为10。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 七层规则 ID。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
}

type BindDetailItem struct {
	// 配置绑定的CLB ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 配置绑定的监听器ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 配置绑定的域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 配置绑定的规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 监听器名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// location的url
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 配置ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UconfigId *string `json:"UconfigId,omitempty" name:"UconfigId"`
}

type BlockedIP struct {
	// 黑名单IP
	IP *string `json:"IP,omitempty" name:"IP"`

	// 加入黑名单的时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type CertIdRelatedWithLoadBalancers struct {
	// 证书ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 与证书关联的负载均衡实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancers []*LoadBalancer `json:"LoadBalancers,omitempty" name:"LoadBalancers"`
}

type CertInfo struct {
	// 证书 ID，如果不填写此项则必须上传证书内容，包括CertName, CertContent，若为服务端证书必须包含CertKey。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 上传证书的名称，如果没有 CertId，则此项必传。
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// 上传证书的公钥；如果没有 CertId，则此项必传。
	CertContent *string `json:"CertContent,omitempty" name:"CertContent"`

	// 上传服务端证书的私钥；如果没有 CertId，则此项必传。
	CertKey *string `json:"CertKey,omitempty" name:"CertKey"`
}

type CertificateInput struct {
	// 认证类型，UNIDIRECTIONAL：单向认证，MUTUAL：双向认证
	SSLMode *string `json:"SSLMode,omitempty" name:"SSLMode"`

	// 服务端证书的 ID，如果不填写此项则必须上传证书，包括 CertContent，CertKey，CertName。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 客户端证书的 ID，当监听器采用双向认证，即 SSLMode=MUTUAL 时，如果不填写此项则必须上传客户端证书，包括 CertCaContent，CertCaName。
	CertCaId *string `json:"CertCaId,omitempty" name:"CertCaId"`

	// 上传服务端证书的名称，如果没有 CertId，则此项必传。
	CertName *string `json:"CertName,omitempty" name:"CertName"`

	// 上传服务端证书的 key，如果没有 CertId，则此项必传。
	CertKey *string `json:"CertKey,omitempty" name:"CertKey"`

	// 上传服务端证书的内容，如果没有 CertId，则此项必传。
	CertContent *string `json:"CertContent,omitempty" name:"CertContent"`

	// 上传客户端 CA 证书的名称，如果 SSLMode=mutual，如果没有 CertCaId，则此项必传。
	CertCaName *string `json:"CertCaName,omitempty" name:"CertCaName"`

	// 上传客户端证书的内容，如果 SSLMode=mutual，如果没有 CertCaId，则此项必传。
	CertCaContent *string `json:"CertCaContent,omitempty" name:"CertCaContent"`
}

type CertificateOutput struct {
	// 认证类型，UNIDIRECTIONAL：单向认证，MUTUAL：双向认证
	SSLMode *string `json:"SSLMode,omitempty" name:"SSLMode"`

	// 服务端证书的ID。
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 客户端证书的 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertCaId *string `json:"CertCaId,omitempty" name:"CertCaId"`

	// 多本服务器证书场景扩展的服务器证书ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtCertIds []*string `json:"ExtCertIds,omitempty" name:"ExtCertIds"`
}

type ClassicalHealth struct {
	// 后端服务的内网 IP
	IP *string `json:"IP,omitempty" name:"IP"`

	// 后端服务的端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 负载均衡的监听端口
	ListenerPort *int64 `json:"ListenerPort,omitempty" name:"ListenerPort"`

	// 转发协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 健康检查结果，1 表示健康，0 表示不健康
	HealthStatus *int64 `json:"HealthStatus,omitempty" name:"HealthStatus"`
}

type ClassicalListener struct {
	// 负载均衡监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 负载均衡监听器端口
	ListenerPort *int64 `json:"ListenerPort,omitempty" name:"ListenerPort"`

	// 监听器后端转发端口
	InstancePort *int64 `json:"InstancePort,omitempty" name:"InstancePort"`

	// 监听器名称
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器协议类型
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 会话保持时间
	SessionExpire *int64 `json:"SessionExpire,omitempty" name:"SessionExpire"`

	// 是否开启了健康检查：1（开启）、0（关闭）
	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`

	// 响应超时时间
	TimeOut *int64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// 检查间隔
	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`

	// 健康阈值
	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`

	// 不健康阈值
	UnhealthNum *int64 `json:"UnhealthNum,omitempty" name:"UnhealthNum"`

	// 传统型公网负载均衡的 HTTP、HTTPS 监听器的请求均衡方法。wrr 表示按权重轮询，ip_hash 表示根据访问的源 IP 进行一致性哈希方式来分发
	HttpHash *string `json:"HttpHash,omitempty" name:"HttpHash"`

	// 传统型公网负载均衡的 HTTP、HTTPS 监听器的健康检查返回码。具体可参考创建监听器中对该字段的解释
	HttpCode *int64 `json:"HttpCode,omitempty" name:"HttpCode"`

	// 传统型公网负载均衡的 HTTP、HTTPS 监听器的健康检查路径
	HttpCheckPath *string `json:"HttpCheckPath,omitempty" name:"HttpCheckPath"`

	// 传统型公网负载均衡的 HTTPS 监听器的认证方式
	SSLMode *string `json:"SSLMode,omitempty" name:"SSLMode"`

	// 传统型公网负载均衡的 HTTPS 监听器的服务端证书 ID
	CertId *string `json:"CertId,omitempty" name:"CertId"`

	// 传统型公网负载均衡的 HTTPS 监听器的客户端证书 ID
	CertCaId *string `json:"CertCaId,omitempty" name:"CertCaId"`

	// 监听器的状态，0 表示创建中，1 表示运行中
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ClassicalLoadBalancerInfo struct {
	// 后端实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 负载均衡实例ID列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

type ClassicalTarget struct {
	// 后端服务的类型，可取值：CVM、ENI（即将支持）
	Type *string `json:"Type,omitempty" name:"Type"`

	// 后端服务的唯一 ID，可通过 DescribeInstances 接口返回字段中的 unInstanceId 字段获取
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 后端服务的转发权重，取值范围：[0, 100]，默认为 10。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 后端服务的外网 IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`

	// 后端服务的内网 IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// 后端服务的实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 后端服务的状态
	// 1：故障，2：运行中，3：创建中，4：已关机，5：已退还，6：退还中， 7：重启中，8：开机中，9：关机中，10：密码重置中，11：格式化中，12：镜像制作中，13：带宽设置中，14：重装系统中，19：升级中，21：热迁移中
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunFlag *int64 `json:"RunFlag,omitempty" name:"RunFlag"`
}

type ClassicalTargetInfo struct {
	// 后端实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 权重，取值范围 [0, 100]
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

// Predefined struct for user
type CloneLoadBalancerRequestParams struct {
	// 负载均衡ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 克隆出负载均衡实例的名称，规则：1-60 个英文、汉字、数字、连接线“-”或下划线“_”。
	// 注意：如果名称与系统中已有负载均衡实例的名称相同，则系统将会自动生成此次创建的负载均衡实例的名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡实例所属的项目 ID，可以通过 [DescribeProject](https://cloud.tencent.com/document/product/378/4400) 接口获取。不传此参数则视为默认项目。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 仅适用于公网负载均衡。设置跨可用区容灾时的主可用区ID，例如 100001 或 ap-guangzhou-1
	// 注：主可用区是需要承载流量的可用区，备可用区默认不承载流量，主可用区不可用时才使用备可用区，平台将为您自动选择最佳备可用区。可通过 [DescribeResources](https://cloud.tencent.com/document/api/214/70213) 接口查询一个地域的主可用区的列表。
	MasterZoneId *string `json:"MasterZoneId,omitempty" name:"MasterZoneId"`

	// 仅适用于公网负载均衡。设置跨可用区容灾时的备可用区ID，例如 100001 或 ap-guangzhou-1
	// 注：备可用区是主可用区故障后，需要承载流量的可用区。可通过 [DescribeResources](https://cloud.tencent.com/document/api/214/70213) 接口查询一个地域的主/备可用区的列表。
	SlaveZoneId *string `json:"SlaveZoneId,omitempty" name:"SlaveZoneId"`

	// 仅适用于公网负载均衡。可用区ID，指定可用区以创建负载均衡实例。如：ap-guangzhou-1。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 仅适用于公网负载均衡。负载均衡的网络计费模式。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 仅适用于公网负载均衡。CMCC | CTCC | CUCC，分别对应 移动 | 电信 | 联通，如果不指定本参数，则默认使用BGP。可通过 DescribeSingleIsp 接口查询一个地域所支持的Isp。如果指定运营商，则网络计费式只能使用按带宽包计费(BANDWIDTH_PACKAGE)。
	VipIsp *string `json:"VipIsp,omitempty" name:"VipIsp"`

	// 指定Vip申请负载均衡。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 购买负载均衡同时，给负载均衡打上标签。
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// 独占集群信息。
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitempty" name:"ExclusiveCluster"`

	// 带宽包ID，指定此参数时，网络计费方式（InternetAccessible.InternetChargeType）只支持按带宽包计费（BANDWIDTH_PACKAGE）。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`

	// 是否支持绑定跨地域/跨Vpc绑定IP的功能。
	SnatPro *bool `json:"SnatPro,omitempty" name:"SnatPro"`

	// 开启绑定跨地域/跨Vpc绑定IP的功能后，创建SnatIp。
	SnatIps []*SnatIp `json:"SnatIps,omitempty" name:"SnatIps"`

	// 公网独占集群ID或者CDCId。
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// 性能容量型规格。
	SlaType *string `json:"SlaType,omitempty" name:"SlaType"`

	// Stgw独占集群的标签。
	ClusterTag *string `json:"ClusterTag,omitempty" name:"ClusterTag"`

	// 仅适用于私有网络内网负载均衡。内网就近接入时，选择可用区下发。
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// EIP 的唯一 ID，形如：eip-11112222，仅适用于内网负载均衡绑定EIP。
	EipAddressId *string `json:"EipAddressId,omitempty" name:"EipAddressId"`
}

type CloneLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 克隆出负载均衡实例的名称，规则：1-60 个英文、汉字、数字、连接线“-”或下划线“_”。
	// 注意：如果名称与系统中已有负载均衡实例的名称相同，则系统将会自动生成此次创建的负载均衡实例的名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡实例所属的项目 ID，可以通过 [DescribeProject](https://cloud.tencent.com/document/product/378/4400) 接口获取。不传此参数则视为默认项目。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 仅适用于公网负载均衡。设置跨可用区容灾时的主可用区ID，例如 100001 或 ap-guangzhou-1
	// 注：主可用区是需要承载流量的可用区，备可用区默认不承载流量，主可用区不可用时才使用备可用区，平台将为您自动选择最佳备可用区。可通过 [DescribeResources](https://cloud.tencent.com/document/api/214/70213) 接口查询一个地域的主可用区的列表。
	MasterZoneId *string `json:"MasterZoneId,omitempty" name:"MasterZoneId"`

	// 仅适用于公网负载均衡。设置跨可用区容灾时的备可用区ID，例如 100001 或 ap-guangzhou-1
	// 注：备可用区是主可用区故障后，需要承载流量的可用区。可通过 [DescribeResources](https://cloud.tencent.com/document/api/214/70213) 接口查询一个地域的主/备可用区的列表。
	SlaveZoneId *string `json:"SlaveZoneId,omitempty" name:"SlaveZoneId"`

	// 仅适用于公网负载均衡。可用区ID，指定可用区以创建负载均衡实例。如：ap-guangzhou-1。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 仅适用于公网负载均衡。负载均衡的网络计费模式。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 仅适用于公网负载均衡。CMCC | CTCC | CUCC，分别对应 移动 | 电信 | 联通，如果不指定本参数，则默认使用BGP。可通过 DescribeSingleIsp 接口查询一个地域所支持的Isp。如果指定运营商，则网络计费式只能使用按带宽包计费(BANDWIDTH_PACKAGE)。
	VipIsp *string `json:"VipIsp,omitempty" name:"VipIsp"`

	// 指定Vip申请负载均衡。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 购买负载均衡同时，给负载均衡打上标签。
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// 独占集群信息。
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitempty" name:"ExclusiveCluster"`

	// 带宽包ID，指定此参数时，网络计费方式（InternetAccessible.InternetChargeType）只支持按带宽包计费（BANDWIDTH_PACKAGE）。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`

	// 是否支持绑定跨地域/跨Vpc绑定IP的功能。
	SnatPro *bool `json:"SnatPro,omitempty" name:"SnatPro"`

	// 开启绑定跨地域/跨Vpc绑定IP的功能后，创建SnatIp。
	SnatIps []*SnatIp `json:"SnatIps,omitempty" name:"SnatIps"`

	// 公网独占集群ID或者CDCId。
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// 性能容量型规格。
	SlaType *string `json:"SlaType,omitempty" name:"SlaType"`

	// Stgw独占集群的标签。
	ClusterTag *string `json:"ClusterTag,omitempty" name:"ClusterTag"`

	// 仅适用于私有网络内网负载均衡。内网就近接入时，选择可用区下发。
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// EIP 的唯一 ID，形如：eip-11112222，仅适用于内网负载均衡绑定EIP。
	EipAddressId *string `json:"EipAddressId,omitempty" name:"EipAddressId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群类型，如TGW，STGW，VPCGW
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 集群标签，只有STGW集群有标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterTag *string `json:"ClusterTag,omitempty" name:"ClusterTag"`

	// 集群所在可用区，如ap-guangzhou-1
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 集群网络类型，如Public，Private
	Network *string `json:"Network,omitempty" name:"Network"`

	// 最大连接数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxConn *int64 `json:"MaxConn,omitempty" name:"MaxConn"`

	// 最大入带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxInFlow *int64 `json:"MaxInFlow,omitempty" name:"MaxInFlow"`

	// 最大入包量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxInPkg *int64 `json:"MaxInPkg,omitempty" name:"MaxInPkg"`

	// 最大出带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxOutFlow *int64 `json:"MaxOutFlow,omitempty" name:"MaxOutFlow"`

	// 最大出包量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxOutPkg *int64 `json:"MaxOutPkg,omitempty" name:"MaxOutPkg"`

	// 最大新建连接数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxNewConn *int64 `json:"MaxNewConn,omitempty" name:"MaxNewConn"`

	// http最大新建连接数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HTTPMaxNewConn *int64 `json:"HTTPMaxNewConn,omitempty" name:"HTTPMaxNewConn"`

	// https最大新建连接数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HTTPSMaxNewConn *int64 `json:"HTTPSMaxNewConn,omitempty" name:"HTTPSMaxNewConn"`

	// http QPS
	// 注意：此字段可能返回 null，表示取不到有效值。
	HTTPQps *int64 `json:"HTTPQps,omitempty" name:"HTTPQps"`

	// https QPS
	// 注意：此字段可能返回 null，表示取不到有效值。
	HTTPSQps *int64 `json:"HTTPSQps,omitempty" name:"HTTPSQps"`

	// 集群内资源总数目
	ResourceCount *int64 `json:"ResourceCount,omitempty" name:"ResourceCount"`

	// 集群内空闲资源数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdleResourceCount *int64 `json:"IdleResourceCount,omitempty" name:"IdleResourceCount"`

	// 集群内转发机的数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalanceDirectorCount *int64 `json:"LoadBalanceDirectorCount,omitempty" name:"LoadBalanceDirectorCount"`

	// 集群的Isp属性，如："BGP","CMCC","CUCC","CTCC","INTERNAL"。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// 集群所在的可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClustersZone *ClustersZone `json:"ClustersZone,omitempty" name:"ClustersZone"`

	// 集群版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClustersVersion *string `json:"ClustersVersion,omitempty" name:"ClustersVersion"`

	// 集群容灾类型，如SINGLE-ZONE，DISASTER-RECOVERY，MUTUAL-DISASTER-RECOVERY
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisasterRecoveryType *string `json:"DisasterRecoveryType,omitempty" name:"DisasterRecoveryType"`
}

type ClusterItem struct {
	// 集群唯一ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群所在可用区，如ap-guangzhou-1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type ClusterResource struct {
	// 集群唯一ID，如tgw-12345678。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// ip地址。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 负载均衡唯一ID，如lb-12345678。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 资源是否闲置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Idle *string `json:"Idle,omitempty" name:"Idle"`

	// 集群名称。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群的Isp属性，如："BGP","CMCC","CUCC","CTCC","INTERNAL"。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// 集群所在的可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClustersZone *ClustersZone `json:"ClustersZone,omitempty" name:"ClustersZone"`
}

type ClustersZone struct {
	// 集群所在的主可用区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterZone []*string `json:"MasterZone,omitempty" name:"MasterZone"`

	// 集群所在的备可用区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaveZone []*string `json:"SlaveZone,omitempty" name:"SlaveZone"`
}

type ConfigListItem struct {
	// 配置ID
	UconfigId *string `json:"UconfigId,omitempty" name:"UconfigId"`

	// 配置类型
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`

	// 配置名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置内容
	ConfigContent *string `json:"ConfigContent,omitempty" name:"ConfigContent"`

	// 增加配置时间
	CreateTimestamp *string `json:"CreateTimestamp,omitempty" name:"CreateTimestamp"`

	// 修改配置时间
	UpdateTimestamp *string `json:"UpdateTimestamp,omitempty" name:"UpdateTimestamp"`
}

// Predefined struct for user
type CreateClsLogSetRequestParams struct {
	// 日志集的名字，不能和cls其他日志集重名。不填默认为clb_logset。
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// 日志集的保存周期，单位：天。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 日志集类型，ACCESS：访问日志，HEALTH：健康检查日志，默认ACCESS。
	LogsetType *string `json:"LogsetType,omitempty" name:"LogsetType"`
}

type CreateClsLogSetRequest struct {
	*tchttp.BaseRequest
	
	// 日志集的名字，不能和cls其他日志集重名。不填默认为clb_logset。
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// 日志集的保存周期，单位：天。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 日志集类型，ACCESS：访问日志，HEALTH：健康检查日志，默认ACCESS。
	LogsetType *string `json:"LogsetType,omitempty" name:"LogsetType"`
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
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateListenerRequestParams struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 要将监听器创建到哪些端口，每个端口对应一个新的监听器。
	Ports []*int64 `json:"Ports,omitempty" name:"Ports"`

	// 监听器协议： TCP | UDP | HTTP | HTTPS | TCP_SSL | QUIC。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 要创建的监听器名称列表，名称与Ports数组按序一一对应，如不需立即命名，则无需提供此参数。
	ListenerNames []*string `json:"ListenerNames,omitempty" name:"ListenerNames"`

	// 健康检查相关参数，此参数仅适用于TCP/UDP/TCP_SSL监听器。
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 证书相关信息，此参数仅适用于TCP_SSL监听器和未开启SNI特性的HTTPS监听器。此参数和MultiCertInfo不能同时传入。
	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`

	// 会话保持时间，单位：秒。可选值：30~3600，默认 0，表示不开启。此参数仅适用于TCP/UDP监听器。
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// 监听器转发的方式。可选值：WRR、LEAST_CONN
	// 分别表示按权重轮询、最小连接数， 默认为 WRR。此参数仅适用于TCP/UDP/TCP_SSL监听器。
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 是否开启SNI特性，此参数仅适用于HTTPS监听器。
	SniSwitch *int64 `json:"SniSwitch,omitempty" name:"SniSwitch"`

	// 后端目标类型，NODE表示绑定普通节点，TARGETGROUP表示绑定目标组。
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// 会话保持类型。不传或传NORMAL表示默认会话保持类型。QUIC_CID 表示根据Quic Connection ID做会话保持。QUIC_CID只支持UDP协议。
	SessionType *string `json:"SessionType,omitempty" name:"SessionType"`

	// 是否开启长连接，此参数仅适用于HTTP/HTTPS监听器，0:关闭；1:开启， 默认关闭。
	KeepaliveEnable *int64 `json:"KeepaliveEnable,omitempty" name:"KeepaliveEnable"`

	// 创建端口段监听器时必须传入此参数，用以标识结束端口。同时，入参Ports只允许传入一个成员，用以标识开始端口。【如果您需要体验端口段功能，请通过 [工单申请](https://console.cloud.tencent.com/workorder/category)】。
	EndPort *uint64 `json:"EndPort,omitempty" name:"EndPort"`

	// 解绑后端目标时，是否发RST给客户端，此参数仅适用于TCP监听器。
	DeregisterTargetRst *bool `json:"DeregisterTargetRst,omitempty" name:"DeregisterTargetRst"`

	// 证书信息，支持同时传入不同算法类型的多本服务端证书；此参数仅适用于未开启SNI特性的HTTPS监听器。此参数和Certificate不能同时传入。
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitempty" name:"MultiCertInfo"`

	// 监听器最大连接数，只有TCP/UDP/TCP_SSL/QUIC监听器支持，不传或者传-1表示监听器维度不限速。
	MaxConn *int64 `json:"MaxConn,omitempty" name:"MaxConn"`

	// 监听器最大新增连接数，只有TCP/UDP/TCP_SSL/QUIC监听器支持，不传或者传-1表示监听器维度不限速。
	MaxCps *int64 `json:"MaxCps,omitempty" name:"MaxCps"`
}

type CreateListenerRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 要将监听器创建到哪些端口，每个端口对应一个新的监听器。
	Ports []*int64 `json:"Ports,omitempty" name:"Ports"`

	// 监听器协议： TCP | UDP | HTTP | HTTPS | TCP_SSL | QUIC。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 要创建的监听器名称列表，名称与Ports数组按序一一对应，如不需立即命名，则无需提供此参数。
	ListenerNames []*string `json:"ListenerNames,omitempty" name:"ListenerNames"`

	// 健康检查相关参数，此参数仅适用于TCP/UDP/TCP_SSL监听器。
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 证书相关信息，此参数仅适用于TCP_SSL监听器和未开启SNI特性的HTTPS监听器。此参数和MultiCertInfo不能同时传入。
	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`

	// 会话保持时间，单位：秒。可选值：30~3600，默认 0，表示不开启。此参数仅适用于TCP/UDP监听器。
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// 监听器转发的方式。可选值：WRR、LEAST_CONN
	// 分别表示按权重轮询、最小连接数， 默认为 WRR。此参数仅适用于TCP/UDP/TCP_SSL监听器。
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 是否开启SNI特性，此参数仅适用于HTTPS监听器。
	SniSwitch *int64 `json:"SniSwitch,omitempty" name:"SniSwitch"`

	// 后端目标类型，NODE表示绑定普通节点，TARGETGROUP表示绑定目标组。
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// 会话保持类型。不传或传NORMAL表示默认会话保持类型。QUIC_CID 表示根据Quic Connection ID做会话保持。QUIC_CID只支持UDP协议。
	SessionType *string `json:"SessionType,omitempty" name:"SessionType"`

	// 是否开启长连接，此参数仅适用于HTTP/HTTPS监听器，0:关闭；1:开启， 默认关闭。
	KeepaliveEnable *int64 `json:"KeepaliveEnable,omitempty" name:"KeepaliveEnable"`

	// 创建端口段监听器时必须传入此参数，用以标识结束端口。同时，入参Ports只允许传入一个成员，用以标识开始端口。【如果您需要体验端口段功能，请通过 [工单申请](https://console.cloud.tencent.com/workorder/category)】。
	EndPort *uint64 `json:"EndPort,omitempty" name:"EndPort"`

	// 解绑后端目标时，是否发RST给客户端，此参数仅适用于TCP监听器。
	DeregisterTargetRst *bool `json:"DeregisterTargetRst,omitempty" name:"DeregisterTargetRst"`

	// 证书信息，支持同时传入不同算法类型的多本服务端证书；此参数仅适用于未开启SNI特性的HTTPS监听器。此参数和Certificate不能同时传入。
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitempty" name:"MultiCertInfo"`

	// 监听器最大连接数，只有TCP/UDP/TCP_SSL/QUIC监听器支持，不传或者传-1表示监听器维度不限速。
	MaxConn *int64 `json:"MaxConn,omitempty" name:"MaxConn"`

	// 监听器最大新增连接数，只有TCP/UDP/TCP_SSL/QUIC监听器支持，不传或者传-1表示监听器维度不限速。
	MaxCps *int64 `json:"MaxCps,omitempty" name:"MaxCps"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateListenerResponseParams struct {
	// 创建的监听器的唯一标识数组。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 负载均衡实例的网络类型：
	// OPEN：公网属性， INTERNAL：内网属性。
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// 负载均衡实例的类型。1：通用的负载均衡实例，目前只支持传入1。
	Forward *int64 `json:"Forward,omitempty" name:"Forward"`

	// 负载均衡实例的名称，只在创建一个实例的时候才会生效。规则：1-60 个英文、汉字、数字、连接线“-”或下划线“_”。
	// 注意：如果名称与系统中已有负载均衡实例的名称相同，则系统将会自动生成此次创建的负载均衡实例的名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡后端目标设备所属的网络 ID，如vpc-12345678，可以通过 [DescribeVpcEx](https://cloud.tencent.com/document/product/215/1372) 接口获取。 不填此参数则默认为DefaultVPC。创建内网负载均衡实例时，此参数必填。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 在私有网络内购买内网负载均衡实例的情况下，必须指定子网 ID，内网负载均衡实例的 VIP 将从这个子网中产生。创建内网负载均衡实例时，此参数必填。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 负载均衡实例所属的项目 ID，可以通过 [DescribeProject](https://cloud.tencent.com/document/product/378/4400) 接口获取。不填此参数则视为默认项目。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 仅适用于公网负载均衡。IP版本，可取值：IPV4、IPV6、IPv6FullChain，默认值 IPV4。说明：取值为IPV6表示为IPV6 NAT64版本；取值为IPv6FullChain，表示为IPv6版本。
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`

	// 创建负载均衡的个数，默认值 1。
	Number *uint64 `json:"Number,omitempty" name:"Number"`

	// 仅适用于公网负载均衡。设置跨可用区容灾时的主可用区ID，例如 100001 或 ap-guangzhou-1
	// 注：主可用区是需要承载流量的可用区，备可用区默认不承载流量，主可用区不可用时才使用备可用区，平台将为您自动选择最佳备可用区。可通过 [DescribeResources](https://cloud.tencent.com/document/api/214/70213) 接口查询一个地域的主可用区的列表。
	MasterZoneId *string `json:"MasterZoneId,omitempty" name:"MasterZoneId"`

	// 仅适用于公网负载均衡。可用区ID，指定可用区以创建负载均衡实例。如：ap-guangzhou-1。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 仅适用于公网负载均衡。负载均衡的网络计费模式。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 仅适用于公网负载均衡。CMCC | CTCC | CUCC，分别对应 移动 | 电信 | 联通，如果不指定本参数，则默认使用BGP。可通过 DescribeSingleIsp 接口查询一个地域所支持的Isp。如果指定运营商，则网络计费式只能使用按带宽包计费(BANDWIDTH_PACKAGE)。
	VipIsp *string `json:"VipIsp,omitempty" name:"VipIsp"`

	// 购买负载均衡的同时，给负载均衡打上标签，最大支持20个标签键值对。
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// 指定VIP申请负载均衡。此参数选填，不填写此参数时自动分配VIP。IPv4和IPv6类型支持此参数，IPv6 NAT64类型不支持。
	// 注意：当指定VIP创建内网实例、或公网IPv6 BGP实例时，若VIP不属于指定VPC子网的网段内时，会创建失败；若VIP已被占用，也会创建失败。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 带宽包ID，指定此参数时，网络计费方式（InternetAccessible.InternetChargeType）只支持按带宽包计费（BANDWIDTH_PACKAGE）。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`

	// 独占型实例信息。若创建独占型的内网负载均衡实例，则此参数必填。
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitempty" name:"ExclusiveCluster"`

	// 创建性能容量型实例。
	// <ul><li>若需要创建性能容量型实例，则此参数必填，且取值为：SLA，表示创建按量计费模式下的默认规格的性能容量型实例。
	// <ul><li>当您开通了普通规格的性能容量型时，SLA对应超强型1规格。普通规格的性能容量型正在内测中，请提交 [内测申请](https://cloud.tencent.com/apply/p/hf45esx99lf)。</li>
	// <li>当您开通了超大型规格的性能容量型时，SLA对应超强型4规格。超大型规格的性能容量型正在内测中，请提交 [工单申请](https://console.cloud.tencent.com/workorder/category)。</li></ul></li><li>若需要创建共享型实例，则无需填写此参数。</li></ul>
	SlaType *string `json:"SlaType,omitempty" name:"SlaType"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 是否支持绑定跨地域/跨Vpc绑定IP的功能。
	SnatPro *bool `json:"SnatPro,omitempty" name:"SnatPro"`

	// 开启绑定跨地域/跨Vpc绑定IP的功能后，创建SnatIp。
	SnatIps []*SnatIp `json:"SnatIps,omitempty" name:"SnatIps"`

	// Stgw独占集群的标签。
	ClusterTag *string `json:"ClusterTag,omitempty" name:"ClusterTag"`

	// 仅适用于公网负载均衡。设置跨可用区容灾时的备可用区ID，例如 100001 或 ap-guangzhou-1
	// 注：备可用区是主可用区故障后，需要承载流量的可用区。可通过 [DescribeResources](https://cloud.tencent.com/document/api/214/70213) 接口查询一个地域的主/备可用区的列表。
	SlaveZoneId *string `json:"SlaveZoneId,omitempty" name:"SlaveZoneId"`

	// EIP 的唯一 ID，形如：eip-11112222，仅适用于内网负载均衡绑定EIP。
	EipAddressId *string `json:"EipAddressId,omitempty" name:"EipAddressId"`

	// Target是否放通来自CLB的流量。开启放通（true）：只验证CLB上的安全组；不开启放通（false）：需同时验证CLB和后端实例上的安全组。
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitempty" name:"LoadBalancerPassToTarget"`
}

type CreateLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例的网络类型：
	// OPEN：公网属性， INTERNAL：内网属性。
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// 负载均衡实例的类型。1：通用的负载均衡实例，目前只支持传入1。
	Forward *int64 `json:"Forward,omitempty" name:"Forward"`

	// 负载均衡实例的名称，只在创建一个实例的时候才会生效。规则：1-60 个英文、汉字、数字、连接线“-”或下划线“_”。
	// 注意：如果名称与系统中已有负载均衡实例的名称相同，则系统将会自动生成此次创建的负载均衡实例的名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡后端目标设备所属的网络 ID，如vpc-12345678，可以通过 [DescribeVpcEx](https://cloud.tencent.com/document/product/215/1372) 接口获取。 不填此参数则默认为DefaultVPC。创建内网负载均衡实例时，此参数必填。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 在私有网络内购买内网负载均衡实例的情况下，必须指定子网 ID，内网负载均衡实例的 VIP 将从这个子网中产生。创建内网负载均衡实例时，此参数必填。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 负载均衡实例所属的项目 ID，可以通过 [DescribeProject](https://cloud.tencent.com/document/product/378/4400) 接口获取。不填此参数则视为默认项目。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 仅适用于公网负载均衡。IP版本，可取值：IPV4、IPV6、IPv6FullChain，默认值 IPV4。说明：取值为IPV6表示为IPV6 NAT64版本；取值为IPv6FullChain，表示为IPv6版本。
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`

	// 创建负载均衡的个数，默认值 1。
	Number *uint64 `json:"Number,omitempty" name:"Number"`

	// 仅适用于公网负载均衡。设置跨可用区容灾时的主可用区ID，例如 100001 或 ap-guangzhou-1
	// 注：主可用区是需要承载流量的可用区，备可用区默认不承载流量，主可用区不可用时才使用备可用区，平台将为您自动选择最佳备可用区。可通过 [DescribeResources](https://cloud.tencent.com/document/api/214/70213) 接口查询一个地域的主可用区的列表。
	MasterZoneId *string `json:"MasterZoneId,omitempty" name:"MasterZoneId"`

	// 仅适用于公网负载均衡。可用区ID，指定可用区以创建负载均衡实例。如：ap-guangzhou-1。
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 仅适用于公网负载均衡。负载均衡的网络计费模式。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 仅适用于公网负载均衡。CMCC | CTCC | CUCC，分别对应 移动 | 电信 | 联通，如果不指定本参数，则默认使用BGP。可通过 DescribeSingleIsp 接口查询一个地域所支持的Isp。如果指定运营商，则网络计费式只能使用按带宽包计费(BANDWIDTH_PACKAGE)。
	VipIsp *string `json:"VipIsp,omitempty" name:"VipIsp"`

	// 购买负载均衡的同时，给负载均衡打上标签，最大支持20个标签键值对。
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// 指定VIP申请负载均衡。此参数选填，不填写此参数时自动分配VIP。IPv4和IPv6类型支持此参数，IPv6 NAT64类型不支持。
	// 注意：当指定VIP创建内网实例、或公网IPv6 BGP实例时，若VIP不属于指定VPC子网的网段内时，会创建失败；若VIP已被占用，也会创建失败。
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 带宽包ID，指定此参数时，网络计费方式（InternetAccessible.InternetChargeType）只支持按带宽包计费（BANDWIDTH_PACKAGE）。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`

	// 独占型实例信息。若创建独占型的内网负载均衡实例，则此参数必填。
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitempty" name:"ExclusiveCluster"`

	// 创建性能容量型实例。
	// <ul><li>若需要创建性能容量型实例，则此参数必填，且取值为：SLA，表示创建按量计费模式下的默认规格的性能容量型实例。
	// <ul><li>当您开通了普通规格的性能容量型时，SLA对应超强型1规格。普通规格的性能容量型正在内测中，请提交 [内测申请](https://cloud.tencent.com/apply/p/hf45esx99lf)。</li>
	// <li>当您开通了超大型规格的性能容量型时，SLA对应超强型4规格。超大型规格的性能容量型正在内测中，请提交 [工单申请](https://console.cloud.tencent.com/workorder/category)。</li></ul></li><li>若需要创建共享型实例，则无需填写此参数。</li></ul>
	SlaType *string `json:"SlaType,omitempty" name:"SlaType"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 是否支持绑定跨地域/跨Vpc绑定IP的功能。
	SnatPro *bool `json:"SnatPro,omitempty" name:"SnatPro"`

	// 开启绑定跨地域/跨Vpc绑定IP的功能后，创建SnatIp。
	SnatIps []*SnatIp `json:"SnatIps,omitempty" name:"SnatIps"`

	// Stgw独占集群的标签。
	ClusterTag *string `json:"ClusterTag,omitempty" name:"ClusterTag"`

	// 仅适用于公网负载均衡。设置跨可用区容灾时的备可用区ID，例如 100001 或 ap-guangzhou-1
	// 注：备可用区是主可用区故障后，需要承载流量的可用区。可通过 [DescribeResources](https://cloud.tencent.com/document/api/214/70213) 接口查询一个地域的主/备可用区的列表。
	SlaveZoneId *string `json:"SlaveZoneId,omitempty" name:"SlaveZoneId"`

	// EIP 的唯一 ID，形如：eip-11112222，仅适用于内网负载均衡绑定EIP。
	EipAddressId *string `json:"EipAddressId,omitempty" name:"EipAddressId"`

	// Target是否放通来自CLB的流量。开启放通（true）：只验证CLB上的安全组；不开启放通（false）：需同时验证CLB和后端实例上的安全组。
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitempty" name:"LoadBalancerPassToTarget"`
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
	delete(f, "ClientToken")
	delete(f, "SnatPro")
	delete(f, "SnatIps")
	delete(f, "ClusterTag")
	delete(f, "SlaveZoneId")
	delete(f, "EipAddressId")
	delete(f, "LoadBalancerPassToTarget")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadBalancerResponseParams struct {
	// 由负载均衡实例唯一 ID 组成的数组。
	// 存在某些场景，如创建出现延迟时，此字段可能返回为空；此时可以根据接口返回的RequestId或DealName参数，通过DescribeTaskStatus接口查询创建的资源ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// 订单号。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 负载均衡唯一性ID，例如：lb-12345678。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 添加的SnatIp信息，可指定IP申请，或者指定子网自动申请。单个CLB实例可申请的默认上限为10个。
	SnatIps []*SnatIp `json:"SnatIps,omitempty" name:"SnatIps"`

	// 添加的SnatIp的个数，可与SnatIps一起使用，但若指定IP时，则不能指定创建的SnatIp个数。默认值为1，数量上限与用户配置有关，默认上限为10。
	Number *uint64 `json:"Number,omitempty" name:"Number"`
}

type CreateLoadBalancerSnatIpsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡唯一性ID，例如：lb-12345678。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 添加的SnatIp信息，可指定IP申请，或者指定子网自动申请。单个CLB实例可申请的默认上限为10个。
	SnatIps []*SnatIp `json:"SnatIps,omitempty" name:"SnatIps"`

	// 添加的SnatIp的个数，可与SnatIps一起使用，但若指定IP时，则不能指定创建的SnatIp个数。默认值为1，数量上限与用户配置有关，默认上限为10。
	Number *uint64 `json:"Number,omitempty" name:"Number"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateRuleRequestParams struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 监听器 ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 新建转发规则的信息。
	Rules []*RuleInput `json:"Rules,omitempty" name:"Rules"`
}

type CreateRuleRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 监听器 ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 新建转发规则的信息。
	Rules []*RuleInput `json:"Rules,omitempty" name:"Rules"`
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
	LocationIds []*string `json:"LocationIds,omitempty" name:"LocationIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 目标组名称，限定50个字符
	TargetGroupName *string `json:"TargetGroupName,omitempty" name:"TargetGroupName"`

	// 目标组的vpcid属性，不填则使用默认vpc
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 目标组的默认端口， 后续添加服务器时可使用该默认端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 目标组绑定的后端服务器
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances"`
}

type CreateTargetGroupRequest struct {
	*tchttp.BaseRequest
	
	// 目标组名称，限定50个字符
	TargetGroupName *string `json:"TargetGroupName,omitempty" name:"TargetGroupName"`

	// 目标组的vpcid属性，不填则使用默认vpc
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 目标组的默认端口， 后续添加服务器时可使用该默认端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 目标组绑定的后端服务器
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTargetGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTargetGroupResponseParams struct {
	// 创建目标组后生成的id
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 主题分区Partition的数量，不传参默认创建1个，最大创建允许10个，分裂/合并操作会改变分区数量，整体上限50个。
	PartitionCount *uint64 `json:"PartitionCount,omitempty" name:"PartitionCount"`

	// 日志类型，ACCESS：访问日志，HEALTH：健康检查日志，默认ACCESS。
	TopicType *string `json:"TopicType,omitempty" name:"TopicType"`

	// 日志集的保存周期，单位：天，默认30天。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 日志主题的存储类型，可选值 HOT（标准存储），COLD（低频存储）；默认为HOT。
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest
	
	// 日志主题的名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 主题分区Partition的数量，不传参默认创建1个，最大创建允许10个，分裂/合并操作会改变分区数量，整体上限50个。
	PartitionCount *uint64 `json:"PartitionCount,omitempty" name:"PartitionCount"`

	// 日志类型，ACCESS：访问日志，HEALTH：健康检查日志，默认ACCESS。
	TopicType *string `json:"TopicType,omitempty" name:"TopicType"`

	// 日志集的保存周期，单位：天，默认30天。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 日志主题的存储类型，可选值 HOT（标准存储），COLD（低频存储）；默认为HOT。
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
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
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type CrossTargets struct {
	// 本地私有网络ID，即负载均衡的VpcId。
	LocalVpcId *string `json:"LocalVpcId,omitempty" name:"LocalVpcId"`

	// 子机或网卡所属的私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子机或网卡的IP地址
	IP *string `json:"IP,omitempty" name:"IP"`

	// 子机或网卡所属的私有网络名称。
	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`

	// 子机的网卡ID。
	EniId *string `json:"EniId,omitempty" name:"EniId"`

	// 子机实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 子机实例名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 子机或者网卡所属的地域。
	Region *string `json:"Region,omitempty" name:"Region"`
}

// Predefined struct for user
type DeleteListenerRequestParams struct {
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 要删除的监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

type DeleteListenerRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 要删除的监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 指定删除的监听器ID数组，最大为20个。若不填则删除负载均衡的所有监听器。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
}

type DeleteLoadBalancerListenersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 指定删除的监听器ID数组，最大为20个。若不填则删除负载均衡的所有监听器。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 要删除的负载均衡实例 ID数组，数组大小最大支持20。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

type DeleteLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的负载均衡实例 ID数组，数组大小最大支持20。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancerResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 负载均衡唯一ID，例如：lb-12345678。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 删除SnatIp地址数组。
	Ips []*string `json:"Ips,omitempty" name:"Ips"`
}

type DeleteLoadBalancerSnatIpsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡唯一ID，例如：lb-12345678。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 删除SnatIp地址数组。
	Ips []*string `json:"Ips,omitempty" name:"Ips"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeleteRewriteRequestParams struct {
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 源监听器ID。
	SourceListenerId *string `json:"SourceListenerId,omitempty" name:"SourceListenerId"`

	// 目标监听器ID。
	TargetListenerId *string `json:"TargetListenerId,omitempty" name:"TargetListenerId"`

	// 转发规则之间的重定向关系。
	RewriteInfos []*RewriteLocationMap `json:"RewriteInfos,omitempty" name:"RewriteInfos"`
}

type DeleteRewriteRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 源监听器ID。
	SourceListenerId *string `json:"SourceListenerId,omitempty" name:"SourceListenerId"`

	// 目标监听器ID。
	TargetListenerId *string `json:"TargetListenerId,omitempty" name:"TargetListenerId"`

	// 转发规则之间的重定向关系。
	RewriteInfos []*RewriteLocationMap `json:"RewriteInfos,omitempty" name:"RewriteInfos"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 要删除的转发规则的ID组成的数组。
	LocationIds []*string `json:"LocationIds,omitempty" name:"LocationIds"`

	// 要删除的转发规则的域名，如果是多域名，可以指定多域名列表中的任意一个。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 要删除的转发规则的转发路径。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 监听器下必须配置一个默认域名，当需要删除默认域名时，可以指定另一个域名作为新的默认域名，如果新的默认域名是多域名，可以指定多域名列表中的任意一个。
	NewDefaultServerDomain *string `json:"NewDefaultServerDomain,omitempty" name:"NewDefaultServerDomain"`
}

type DeleteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 要删除的转发规则的ID组成的数组。
	LocationIds []*string `json:"LocationIds,omitempty" name:"LocationIds"`

	// 要删除的转发规则的域名，如果是多域名，可以指定多域名列表中的任意一个。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 要删除的转发规则的转发路径。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 监听器下必须配置一个默认域名，当需要删除默认域名时，可以指定另一个域名作为新的默认域名，如果新的默认域名是多域名，可以指定多域名列表中的任意一个。
	NewDefaultServerDomain *string `json:"NewDefaultServerDomain,omitempty" name:"NewDefaultServerDomain"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 目标组的ID数组。
	TargetGroupIds []*string `json:"TargetGroupIds,omitempty" name:"TargetGroupIds"`
}

type DeleteTargetGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 目标组的ID数组。
	TargetGroupIds []*string `json:"TargetGroupIds,omitempty" name:"TargetGroupIds"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeregisterFunctionTargetsRequestParams struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 待解绑的云函数列表。
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitempty" name:"FunctionTargets"`

	// 目标转发规则的 ID，当将云函数从七层转发规则上解绑时，必须输入此参数或 Domain+Url 参数。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 目标转发规则的域名，若已经输入 LocationId 参数，则本参数不生效。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 目标转发规则的 URL，若已经输入 LocationId 参数，则本参数不生效。
	Url *string `json:"Url,omitempty" name:"Url"`
}

type DeregisterFunctionTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 待解绑的云函数列表。
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitempty" name:"FunctionTargets"`

	// 目标转发规则的 ID，当将云函数从七层转发规则上解绑时，必须输入此参数或 Domain+Url 参数。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 目标转发规则的域名，若已经输入 LocationId 参数，则本参数不生效。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 目标转发规则的 URL，若已经输入 LocationId 参数，则本参数不生效。
	Url *string `json:"Url,omitempty" name:"Url"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DeregisterTargetGroupInstancesRequestParams struct {
	// 目标组ID。
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// 待解绑的服务器信息。
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances"`
}

type DeregisterTargetGroupInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 目标组ID。
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// 待解绑的服务器信息。
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 后端服务的实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DeregisterTargetsFromClassicalLBRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 后端服务的实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 监听器 ID，格式如 lbl-12345678。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 要解绑的后端服务列表，数组长度最大支持20。
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// 转发规则的ID，格式如 loc-12345678，当从七层转发规则解绑机器时，必须提供此参数或Domain+URL两者之一。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 目标规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 目标规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitempty" name:"Url"`
}

type DeregisterTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID，格式如 lb-12345678。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 监听器 ID，格式如 lbl-12345678。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 要解绑的后端服务列表，数组长度最大支持20。
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// 转发规则的ID，格式如 loc-12345678，当从七层转发规则解绑机器时，必须提供此参数或Domain+URL两者之一。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 目标规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 目标规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitempty" name:"Url"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeBlockIPListRequestParams struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 数据偏移量，默认为 0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回IP的最大个数，默认为 100000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeBlockIPListRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 数据偏移量，默认为 0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回IP的最大个数，默认为 100000。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
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
	BlockedIPCount *uint64 `json:"BlockedIPCount,omitempty" name:"BlockedIPCount"`

	// 获取用户真实IP的字段
	ClientIPField *string `json:"ClientIPField,omitempty" name:"ClientIPField"`

	// 加入了12360黑名单的IP列表
	BlockedIPList []*BlockedIP `json:"BlockedIPList,omitempty" name:"BlockedIPList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeBlockIPTaskRequest struct {
	*tchttp.BaseRequest
	
	// ModifyBlockIPList 接口返回的异步任务的ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
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
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeClassicalLBByInstanceIdRequestParams struct {
	// 后端实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DescribeClassicalLBByInstanceIdRequest struct {
	*tchttp.BaseRequest
	
	// 后端实例ID列表。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
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
	LoadBalancerInfoList []*ClassicalLoadBalancerInfo `json:"LoadBalancerInfoList,omitempty" name:"LoadBalancerInfoList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
}

type DescribeClassicalLBHealthStatusRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthList []*ClassicalHealth `json:"HealthList,omitempty" name:"HealthList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID列表。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// 负载均衡监听的协议：'TCP', 'UDP', 'HTTP', 'HTTPS'。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 负载均衡监听端口，范围为[1-65535]。
	ListenerPort *int64 `json:"ListenerPort,omitempty" name:"ListenerPort"`

	// 监听器的状态，0：创建中，1：运行中。
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeClassicalLBListenersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID列表。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// 负载均衡监听的协议：'TCP', 'UDP', 'HTTP', 'HTTPS'。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 负载均衡监听端口，范围为[1-65535]。
	ListenerPort *int64 `json:"ListenerPort,omitempty" name:"ListenerPort"`

	// 监听器的状态，0：创建中，1：运行中。
	Status *int64 `json:"Status,omitempty" name:"Status"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Listeners []*ClassicalListener `json:"Listeners,omitempty" name:"Listeners"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
}

type DescribeClassicalLBTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Targets []*ClassicalTarget `json:"Targets,omitempty" name:"Targets"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 健康检查日志集的 ID。
	HealthLogsetId *string `json:"HealthLogsetId,omitempty" name:"HealthLogsetId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回集群中资源列表起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询集群中资源列表条件，详细的过滤条件如下：
	// <li> cluster-id - String - 是否必填：否 - （过滤条件）按照 集群 的唯一ID过滤，如 ："tgw-12345678","stgw-12345678","vpcgw-12345678"。</li>
	// <li> vip - String - 是否必填：否 - （过滤条件）按照vip过滤。</li>
	// <li> loadblancer-id - String - 是否必填：否 - （过滤条件）按照负载均衡唯一ID过滤。</li>
	// <li> idle - String 是否必填：否 - （过滤条件）按照是否闲置过滤，如"True","False"。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeClusterResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 返回集群中资源列表数目，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回集群中资源列表起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询集群中资源列表条件，详细的过滤条件如下：
	// <li> cluster-id - String - 是否必填：否 - （过滤条件）按照 集群 的唯一ID过滤，如 ："tgw-12345678","stgw-12345678","vpcgw-12345678"。</li>
	// <li> vip - String - 是否必填：否 - （过滤条件）按照vip过滤。</li>
	// <li> loadblancer-id - String - 是否必填：否 - （过滤条件）按照负载均衡唯一ID过滤。</li>
	// <li> idle - String 是否必填：否 - （过滤条件）按照是否闲置过滤，如"True","False"。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	ClusterResourceSet []*ClusterResource `json:"ClusterResourceSet,omitempty" name:"ClusterResourceSet"`

	// 集群中资源总数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回后端服务列表起始偏移量，默认0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询跨域2.0版本云联网后端子机和网卡服务列表条件，详细的过滤条件如下：
	// <li> vpc-id - String - 是否必填：否 - （过滤条件）按照 本地私有网络ID，即负载均衡的VpcId 过滤，如："vpc-12345678"。</li>
	// <li> ip - String - 是否必填：否 - （过滤条件）按照 后端服务ip 过滤，如："192.168.0.1"。</li>
	// <li> listener-id - String - 是否必填：否 - （过滤条件）按照 监听器ID 过滤，如："lbl-12345678"。</li>
	// <li> location-id - String - 是否必填：否 - （过滤条件）按照 七层监听器规则ID 过滤，如："loc-12345678"。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeCrossTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 返回后端服务列表数目，默认20，最大值100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回后端服务列表起始偏移量，默认0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询跨域2.0版本云联网后端子机和网卡服务列表条件，详细的过滤条件如下：
	// <li> vpc-id - String - 是否必填：否 - （过滤条件）按照 本地私有网络ID，即负载均衡的VpcId 过滤，如："vpc-12345678"。</li>
	// <li> ip - String - 是否必填：否 - （过滤条件）按照 后端服务ip 过滤，如："192.168.0.1"。</li>
	// <li> listener-id - String - 是否必填：否 - （过滤条件）按照 监听器ID 过滤，如："lbl-12345678"。</li>
	// <li> location-id - String - 是否必填：否 - （过滤条件）按照 七层监听器规则ID 过滤，如："loc-12345678"。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 后端服务列表。
	CrossTargetSet []*CrossTargets `json:"CrossTargetSet,omitempty" name:"CrossTargetSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 配置ID
	UconfigId *string `json:"UconfigId,omitempty" name:"UconfigId"`

	// 拉取绑定关系列表开始位置，默认值 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 拉取绑定关系列表数目，默认值 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type DescribeCustomizedConfigAssociateListRequest struct {
	*tchttp.BaseRequest
	
	// 配置ID
	UconfigId *string `json:"UconfigId,omitempty" name:"UconfigId"`

	// 拉取绑定关系列表开始位置，默认值 0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 拉取绑定关系列表数目，默认值 20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`
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
	BindList []*BindDetailItem `json:"BindList,omitempty" name:"BindList"`

	// 绑定关系总数目
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`

	// 拉取页偏移，默认值0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 拉取数目，默认值20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 拉取指定配置名字，模糊匹配。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置ID
	UconfigIds []*string `json:"UconfigIds,omitempty" name:"UconfigIds"`

	// 过滤条件如下：
	// <li> loadbalancer-id - String - 是否必填：否 - （过滤条件）按照 负载均衡ID 过滤，如："lb-12345678"。</li>
	// <li> vip - String - 是否必填：否 - （过滤条件）按照 负载均衡Vip 过滤，如："1.1.1.1","2204::22:3"。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeCustomizedConfigListRequest struct {
	*tchttp.BaseRequest
	
	// 配置类型:CLB 负载均衡维度。 SERVER 域名维度。 LOCATION 规则维度。
	ConfigType *string `json:"ConfigType,omitempty" name:"ConfigType"`

	// 拉取页偏移，默认值0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 拉取数目，默认值20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 拉取指定配置名字，模糊匹配。
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 配置ID
	UconfigIds []*string `json:"UconfigIds,omitempty" name:"UconfigIds"`

	// 过滤条件如下：
	// <li> loadbalancer-id - String - 是否必填：否 - （过滤条件）按照 负载均衡ID 过滤，如："lb-12345678"。</li>
	// <li> vip - String - 是否必填：否 - （过滤条件）按照 负载均衡Vip 过滤，如："1.1.1.1","2204::22:3"。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	ConfigList []*ConfigListItem `json:"ConfigList,omitempty" name:"ConfigList"`

	// 配置数目
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回集群列表起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询集群列表条件，详细的过滤条件如下：
	// <li> cluster-type - String - 是否必填：否 - （过滤条件）按照 集群 的类型过滤，包括"TGW","STGW","VPCGW"。</li>
	// <li> cluster-id - String - 是否必填：否 - （过滤条件）按照 集群 的唯一ID过滤，如 ："tgw-12345678","stgw-12345678","vpcgw-12345678"。</li>
	// <li> cluster-name - String - 是否必填：否 - （过滤条件）按照 集群 的名称过滤。</li>
	// <li> cluster-tag - String - 是否必填：否 - （过滤条件）按照 集群 的标签过滤。（只有TGW/STGW集群有集群标签） </li>
	// <li> vip - String - 是否必填：否 - （过滤条件）按照 集群 内的vip过滤。</li>
	// <li> loadblancer-id - String - 是否必填：否 - （过滤条件）按照 集群 内的负载均衡唯一ID过滤。</li>
	// <li> network - String - 是否必填：否 - （过滤条件）按照 集群 的网络类型过滤，如："Public","Private"。</li>
	// <li> zone - String - 是否必填：否 - （过滤条件）按照 集群 所在可用区过滤，如："ap-guangzhou-1"（广州一区）。</li>
	// <li> isp -- String - 是否必填：否 - （过滤条件）按照TGW集群的 Isp 类型过滤，如："BGP","CMCC","CUCC","CTCC","INTERNAL"。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeExclusiveClustersRequest struct {
	*tchttp.BaseRequest
	
	// 返回集群列表数目，默认值为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回集群列表起始偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询集群列表条件，详细的过滤条件如下：
	// <li> cluster-type - String - 是否必填：否 - （过滤条件）按照 集群 的类型过滤，包括"TGW","STGW","VPCGW"。</li>
	// <li> cluster-id - String - 是否必填：否 - （过滤条件）按照 集群 的唯一ID过滤，如 ："tgw-12345678","stgw-12345678","vpcgw-12345678"。</li>
	// <li> cluster-name - String - 是否必填：否 - （过滤条件）按照 集群 的名称过滤。</li>
	// <li> cluster-tag - String - 是否必填：否 - （过滤条件）按照 集群 的标签过滤。（只有TGW/STGW集群有集群标签） </li>
	// <li> vip - String - 是否必填：否 - （过滤条件）按照 集群 内的vip过滤。</li>
	// <li> loadblancer-id - String - 是否必填：否 - （过滤条件）按照 集群 内的负载均衡唯一ID过滤。</li>
	// <li> network - String - 是否必填：否 - （过滤条件）按照 集群 的网络类型过滤，如："Public","Private"。</li>
	// <li> zone - String - 是否必填：否 - （过滤条件）按照 集群 所在可用区过滤，如："ap-guangzhou-1"（广州一区）。</li>
	// <li> isp -- String - 是否必填：否 - （过滤条件）按照TGW集群的 Isp 类型过滤，如："BGP","CMCC","CUCC","CTCC","INTERNAL"。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	ClusterSet []*Cluster `json:"ClusterSet,omitempty" name:"ClusterSet"`

	// 集群总数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回负载均衡实例的数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 负载均衡所在地域。
	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitempty" name:"LoadBalancerRegion"`
}

type DescribeIdleLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// 数据偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回负载均衡实例的数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 负载均衡所在地域。
	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitempty" name:"LoadBalancerRegion"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	IdleLoadBalancers []*IdleLoadBalancer `json:"IdleLoadBalancers,omitempty" name:"IdleLoadBalancers"`

	// 所有闲置实例数目
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeLBListenersRequestParams struct {
	// 需要查询的内网ip列表
	Backends []*LbRsItem `json:"Backends,omitempty" name:"Backends"`
}

type DescribeLBListenersRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的内网ip列表
	Backends []*LbRsItem `json:"Backends,omitempty" name:"Backends"`
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
	LoadBalancers []*LBItem `json:"LoadBalancers,omitempty" name:"LoadBalancers"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeListenersRequestParams struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 要查询的负载均衡监听器 ID 数组，最大为100个。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// 要查询的监听器协议类型，取值 TCP | UDP | HTTP | HTTPS | TCP_SSL。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 要查询的监听器的端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`
}

type DescribeListenersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 要查询的负载均衡监听器 ID 数组，最大为100个。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// 要查询的监听器协议类型，取值 TCP | UDP | HTTP | HTTPS | TCP_SSL。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 要查询的监听器的端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`
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
	Listeners []*Listener `json:"Listeners,omitempty" name:"Listeners"`

	// 总的监听器个数（根据端口、协议、监听器ID过滤后）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 服务端证书的ID，或客户端证书的ID
	CertIds []*string `json:"CertIds,omitempty" name:"CertIds"`
}

type DescribeLoadBalancerListByCertIdRequest struct {
	*tchttp.BaseRequest
	
	// 服务端证书的ID，或客户端证书的ID
	CertIds []*string `json:"CertIds,omitempty" name:"CertIds"`
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
	CertSet []*CertIdRelatedWithLoadBalancers `json:"CertSet,omitempty" name:"CertSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 运行中的负载均衡数目
	RunningCount *int64 `json:"RunningCount,omitempty" name:"RunningCount"`

	// 隔离中的负载均衡数目
	IsolationCount *int64 `json:"IsolationCount,omitempty" name:"IsolationCount"`

	// 即将到期的负载均衡数目
	WillExpireCount *int64 `json:"WillExpireCount,omitempty" name:"WillExpireCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitempty" name:"LoadBalancerRegion"`
}

type DescribeLoadBalancerTrafficRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡所在地域，不传默认返回所有地域负载均衡。
	LoadBalancerRegion *string `json:"LoadBalancerRegion,omitempty" name:"LoadBalancerRegion"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerTraffic []*LoadBalancerTraffic `json:"LoadBalancerTraffic,omitempty" name:"LoadBalancerTraffic"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回负载均衡列表起始偏移量，默认0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 选择返回的Fields列表，系统仅会返回Fileds中填写的字段，可填写的字段详情请参见<a href="https://cloud.tencent.com/document/api/214/30694#LoadBalancerDetail">LoadBalancerDetail</a>。若未在Fileds填写相关字段，则此字段返回null。Fileds中默认添加LoadBalancerId和LoadBalancerName字段。
	Fields []*string `json:"Fields,omitempty" name:"Fields"`

	// 当Fields包含TargetId、TargetAddress、TargetPort、TargetWeight等Fields时，必选选择导出目标组的Target或者非目标组Target，值范围NODE、GROUP。
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// 查询负载均衡详细信息列表条件，详细的过滤条件如下：
	// <li> loadbalancer-id - String - 是否必填：否 - （过滤条件）按照 负载均衡ID 过滤，如："lb-12345678"。</li>
	// <li> project-id - String - 是否必填：否 - （过滤条件）按照 项目ID 过滤，如："0","123"。</li>
	// <li> network - String - 是否必填：否 - （过滤条件）按照 负载均衡网络类型 过滤，如："Public","Private"。</li>
	// <li> vip - String - 是否必填：否 - （过滤条件）按照 负载均衡Vip 过滤，如："1.1.1.1","2204::22:3"。</li>
	// <li> target-ip - String - 是否必填：否 - （过滤条件）按照 后端目标内网Ip 过滤，如："1.1.1.1","2203::214:4"。</li>
	// <li> vpcid - String - 是否必填：否 - （过滤条件）按照 负载均衡所属vpcId 过滤，如："vpc-12345678"。</li>
	// <li> zone - String - 是否必填：否 - （过滤条件）按照 负载均衡所属的可用区 过滤，如："ap-guangzhou-1"。</li>
	// <li> tag-key - String - 是否必填：否 - （过滤条件）按照 负载均衡标签的标签键 过滤，如："name"。</li>
	// <li> tag:* - String - 是否必填：否 - （过滤条件）按照 负载均衡的标签 过滤，':' 后面跟的是标签键。如：过滤标签键name，标签值zhangsan,lisi，{"Name": "tag:name","Values": ["zhangsan", "lisi"]}。</li>
	// <li> fuzzy-search - String - 是否必填：否 - （过滤条件）按照 负载均衡Vip，负载均衡名称 模糊搜索，如："1.1"。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeLoadBalancersDetailRequest struct {
	*tchttp.BaseRequest
	
	// 返回负载均衡列表数目，默认20，最大值100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回负载均衡列表起始偏移量，默认0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 选择返回的Fields列表，系统仅会返回Fileds中填写的字段，可填写的字段详情请参见<a href="https://cloud.tencent.com/document/api/214/30694#LoadBalancerDetail">LoadBalancerDetail</a>。若未在Fileds填写相关字段，则此字段返回null。Fileds中默认添加LoadBalancerId和LoadBalancerName字段。
	Fields []*string `json:"Fields,omitempty" name:"Fields"`

	// 当Fields包含TargetId、TargetAddress、TargetPort、TargetWeight等Fields时，必选选择导出目标组的Target或者非目标组Target，值范围NODE、GROUP。
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// 查询负载均衡详细信息列表条件，详细的过滤条件如下：
	// <li> loadbalancer-id - String - 是否必填：否 - （过滤条件）按照 负载均衡ID 过滤，如："lb-12345678"。</li>
	// <li> project-id - String - 是否必填：否 - （过滤条件）按照 项目ID 过滤，如："0","123"。</li>
	// <li> network - String - 是否必填：否 - （过滤条件）按照 负载均衡网络类型 过滤，如："Public","Private"。</li>
	// <li> vip - String - 是否必填：否 - （过滤条件）按照 负载均衡Vip 过滤，如："1.1.1.1","2204::22:3"。</li>
	// <li> target-ip - String - 是否必填：否 - （过滤条件）按照 后端目标内网Ip 过滤，如："1.1.1.1","2203::214:4"。</li>
	// <li> vpcid - String - 是否必填：否 - （过滤条件）按照 负载均衡所属vpcId 过滤，如："vpc-12345678"。</li>
	// <li> zone - String - 是否必填：否 - （过滤条件）按照 负载均衡所属的可用区 过滤，如："ap-guangzhou-1"。</li>
	// <li> tag-key - String - 是否必填：否 - （过滤条件）按照 负载均衡标签的标签键 过滤，如："name"。</li>
	// <li> tag:* - String - 是否必填：否 - （过滤条件）按照 负载均衡的标签 过滤，':' 后面跟的是标签键。如：过滤标签键name，标签值zhangsan,lisi，{"Name": "tag:name","Values": ["zhangsan", "lisi"]}。</li>
	// <li> fuzzy-search - String - 是否必填：否 - （过滤条件）按照 负载均衡Vip，负载均衡名称 模糊搜索，如："1.1"。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 负载均衡详情列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerDetailSet []*LoadBalancerDetail `json:"LoadBalancerDetailSet,omitempty" name:"LoadBalancerDetailSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// 负载均衡实例的网络类型：
	// OPEN：公网属性， INTERNAL：内网属性。
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// 负载均衡实例的类型。1：通用的负载均衡实例，0：传统型负载均衡实例。如果不传此参数，则查询所有类型的负载均衡实例。
	Forward *int64 `json:"Forward,omitempty" name:"Forward"`

	// 负载均衡实例的名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 腾讯云为负载均衡实例分配的域名，本参数仅对传统型公网负载均衡才有意义。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 负载均衡实例的 VIP 地址，支持多个。
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`

	// 负载均衡绑定的后端服务的外网 IP，只支持查询云服务器的公网 IP。
	BackendPublicIps []*string `json:"BackendPublicIps,omitempty" name:"BackendPublicIps"`

	// 负载均衡绑定的后端服务的内网 IP，只支持查询云服务器的内网 IP。
	BackendPrivateIps []*string `json:"BackendPrivateIps,omitempty" name:"BackendPrivateIps"`

	// 数据偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回负载均衡实例的数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序参数，支持以下字段：LoadBalancerName，CreateTime，Domain，LoadBalancerType。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 1：倒序，0：顺序，默认按照创建时间倒序。
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 搜索字段，模糊匹配名称、域名、VIP。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 负载均衡实例所属的项目 ID，可以通过 DescribeProject 接口获取。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 负载均衡是否绑定后端服务，0：没有绑定后端服务，1：绑定后端服务，-1：查询全部。
	WithRs *int64 `json:"WithRs,omitempty" name:"WithRs"`

	// 负载均衡实例所属私有网络唯一ID，如 vpc-bhqkbhdx，
	// 基础网络可传入'0'。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 安全组ID，如 sg-m1cc****。
	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// 主可用区ID，如 ："100001" （对应的是广州一区）。
	MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`

	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为100。<br/>`Filter.Name`和`Filter.Values`皆为必填项。详细的过滤条件如下：
	// <li> charge-type - String - 是否必填：否 - （过滤条件）按照 CLB 的实例计费模式过滤，包括"PREPAID","POSTPAID_BY_HOUR"。</li>
	// <li> internet-charge-type - String - 是否必填：否 - （过滤条件）按照 CLB 的网络计费模式过滤，包括"BANDWIDTH_PREPAID","TRAFFIC_POSTPAID_BY_HOUR","BANDWIDTH_POSTPAID_BY_HOUR","BANDWIDTH_PACKAGE"。</li>
	// <li> master-zone-id - String - 是否必填：否 - （过滤条件）按照 CLB 的主可用区ID过滤，如 ："100001" （对应的是广州一区）。</li>
	// <li> tag-key - String - 是否必填：否 - （过滤条件）按照 CLB 标签的键过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 - （过滤条件）按照CLB标签键值对进行过滤，tag-key使用具体的标签键进行替换。</li>
	// <li> function-name - String - 是否必填：否 - （过滤条件）按照 CLB 后端绑定的SCF云函数的函数名称过滤。</li>
	// <li> function-name - String - 是否必填：否 - （过滤条件）按照 CLB 后端绑定的SCF云函数的函数名称过滤。</li>
	// <li> vip-isp - String - 是否必填：否 - （过滤条件）按照 CLB VIP的运营商类型过滤，如："BGP","INTERNAL","CMCC","CTCC","CUCC"等。</li>
	// <li> sla-type - String - 是否必填：否 - （过滤条件）按照 CLB 的性能容量型规格过滤，包括"clb.c2.medium","clb.c3.small","clb.c3.medium","clb.c4.small","clb.c4.medium","clb.c4.large","clb.c4.xlarge"。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。实例ID数量上限为20个。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// 负载均衡实例的网络类型：
	// OPEN：公网属性， INTERNAL：内网属性。
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// 负载均衡实例的类型。1：通用的负载均衡实例，0：传统型负载均衡实例。如果不传此参数，则查询所有类型的负载均衡实例。
	Forward *int64 `json:"Forward,omitempty" name:"Forward"`

	// 负载均衡实例的名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 腾讯云为负载均衡实例分配的域名，本参数仅对传统型公网负载均衡才有意义。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 负载均衡实例的 VIP 地址，支持多个。
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`

	// 负载均衡绑定的后端服务的外网 IP，只支持查询云服务器的公网 IP。
	BackendPublicIps []*string `json:"BackendPublicIps,omitempty" name:"BackendPublicIps"`

	// 负载均衡绑定的后端服务的内网 IP，只支持查询云服务器的内网 IP。
	BackendPrivateIps []*string `json:"BackendPrivateIps,omitempty" name:"BackendPrivateIps"`

	// 数据偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回负载均衡实例的数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 排序参数，支持以下字段：LoadBalancerName，CreateTime，Domain，LoadBalancerType。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 1：倒序，0：顺序，默认按照创建时间倒序。
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 搜索字段，模糊匹配名称、域名、VIP。
	SearchKey *string `json:"SearchKey,omitempty" name:"SearchKey"`

	// 负载均衡实例所属的项目 ID，可以通过 DescribeProject 接口获取。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 负载均衡是否绑定后端服务，0：没有绑定后端服务，1：绑定后端服务，-1：查询全部。
	WithRs *int64 `json:"WithRs,omitempty" name:"WithRs"`

	// 负载均衡实例所属私有网络唯一ID，如 vpc-bhqkbhdx，
	// 基础网络可传入'0'。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 安全组ID，如 sg-m1cc****。
	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// 主可用区ID，如 ："100001" （对应的是广州一区）。
	MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`

	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为100。<br/>`Filter.Name`和`Filter.Values`皆为必填项。详细的过滤条件如下：
	// <li> charge-type - String - 是否必填：否 - （过滤条件）按照 CLB 的实例计费模式过滤，包括"PREPAID","POSTPAID_BY_HOUR"。</li>
	// <li> internet-charge-type - String - 是否必填：否 - （过滤条件）按照 CLB 的网络计费模式过滤，包括"BANDWIDTH_PREPAID","TRAFFIC_POSTPAID_BY_HOUR","BANDWIDTH_POSTPAID_BY_HOUR","BANDWIDTH_PACKAGE"。</li>
	// <li> master-zone-id - String - 是否必填：否 - （过滤条件）按照 CLB 的主可用区ID过滤，如 ："100001" （对应的是广州一区）。</li>
	// <li> tag-key - String - 是否必填：否 - （过滤条件）按照 CLB 标签的键过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 - （过滤条件）按照CLB标签键值对进行过滤，tag-key使用具体的标签键进行替换。</li>
	// <li> function-name - String - 是否必填：否 - （过滤条件）按照 CLB 后端绑定的SCF云函数的函数名称过滤。</li>
	// <li> function-name - String - 是否必填：否 - （过滤条件）按照 CLB 后端绑定的SCF云函数的函数名称过滤。</li>
	// <li> vip-isp - String - 是否必填：否 - （过滤条件）按照 CLB VIP的运营商类型过滤，如："BGP","INTERNAL","CMCC","CTCC","CUCC"等。</li>
	// <li> sla-type - String - 是否必填：否 - （过滤条件）按照 CLB 的性能容量型规格过滤，包括"clb.c2.medium","clb.c3.small","clb.c3.medium","clb.c4.small","clb.c4.medium","clb.c4.large","clb.c4.xlarge"。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancersResponseParams struct {
	// 满足过滤条件的负载均衡实例总数。此数值与入参中的Limit无关。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回的负载均衡实例数组。
	LoadBalancerSet []*LoadBalancer `json:"LoadBalancerSet,omitempty" name:"LoadBalancerSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	QuotaSet []*Quota `json:"QuotaSet,omitempty" name:"QuotaSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回可用区资源列表起始偏移量，默认0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询可用区资源列表条件，详细的过滤条件如下：
	// <li> zone - String - 是否必填：否 - （过滤条件）按照 可用区 过滤，如："ap-guangzhou-1"（广州一区）。</li>
	// <li> isp -- String - 是否必填：否 - （过滤条件）按照 Isp 类型过滤，如："BGP","CMCC","CUCC","CTCC"。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeResourcesRequest struct {
	*tchttp.BaseRequest
	
	// 返回可用区资源列表数目，默认20，最大值100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回可用区资源列表起始偏移量，默认0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询可用区资源列表条件，详细的过滤条件如下：
	// <li> zone - String - 是否必填：否 - （过滤条件）按照 可用区 过滤，如："ap-guangzhou-1"（广州一区）。</li>
	// <li> isp -- String - 是否必填：否 - （过滤条件）按照 Isp 类型过滤，如："BGP","CMCC","CUCC","CTCC"。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	ZoneResourceSet []*ZoneResource `json:"ZoneResourceSet,omitempty" name:"ZoneResourceSet"`

	// 可用区资源列表数目。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID数组。
	SourceListenerIds []*string `json:"SourceListenerIds,omitempty" name:"SourceListenerIds"`

	// 负载均衡转发规则的ID数组。
	SourceLocationIds []*string `json:"SourceLocationIds,omitempty" name:"SourceLocationIds"`
}

type DescribeRewriteRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID数组。
	SourceListenerIds []*string `json:"SourceListenerIds,omitempty" name:"SourceListenerIds"`

	// 负载均衡转发规则的ID数组。
	SourceLocationIds []*string `json:"SourceLocationIds,omitempty" name:"SourceLocationIds"`
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
	RewriteSet []*RuleOutput `json:"RewriteSet,omitempty" name:"RewriteSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeTargetGroupInstancesRequestParams struct {
	// 过滤条件，当前仅支持TargetGroupId，BindIP，InstanceId过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 显示数量限制，默认20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 显示的偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeTargetGroupInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，当前仅支持TargetGroupId，BindIP，InstanceId过滤。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 显示数量限制，默认20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 显示的偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 绑定的服务器信息。
	TargetGroupInstanceSet []*TargetGroupBackend `json:"TargetGroupInstanceSet,omitempty" name:"TargetGroupInstanceSet"`

	// 实际统计数量，不受Limit、Offset、CAM的影响。
	RealCount *uint64 `json:"RealCount,omitempty" name:"RealCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 目标组ID数组。
	TargetGroupIds []*string `json:"TargetGroupIds,omitempty" name:"TargetGroupIds"`

	// 过滤条件数组，支持TargetGroupVpcId和TargetGroupName。与TargetGroupIds互斥，优先使用目标组ID。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 显示的偏移起始量。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 显示条数限制，默认为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeTargetGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 目标组ID数组。
	TargetGroupIds []*string `json:"TargetGroupIds,omitempty" name:"TargetGroupIds"`

	// 过滤条件数组，支持TargetGroupVpcId和TargetGroupName。与TargetGroupIds互斥，优先使用目标组ID。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 显示的偏移起始量。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 显示条数限制，默认为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
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
	// 显示的结果数量。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 显示的目标组信息集合。
	TargetGroupSet []*TargetGroupInfo `json:"TargetGroupSet,omitempty" name:"TargetGroupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TargetGroupIds []*string `json:"TargetGroupIds,omitempty" name:"TargetGroupIds"`

	// 显示条数限制，默认为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 显示的偏移起始量。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件数组，与TargetGroupIds互斥，支持TargetGroupVpcId和TargetGroupName。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeTargetGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 目标组ID，与Filters互斥。
	TargetGroupIds []*string `json:"TargetGroupIds,omitempty" name:"TargetGroupIds"`

	// 显示条数限制，默认为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 显示的偏移起始量。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件数组，与TargetGroupIds互斥，支持TargetGroupVpcId和TargetGroupName。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 显示的目标组信息集合。
	TargetGroupSet []*TargetGroupInfo `json:"TargetGroupSet,omitempty" name:"TargetGroupSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 要查询的负载均衡实例ID列表。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

type DescribeTargetHealthRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的负载均衡实例ID列表。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetHealthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetHealthResponseParams struct {
	// 负载均衡实例列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancers []*LoadBalancerHealth `json:"LoadBalancers,omitempty" name:"LoadBalancers"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 监听器 ID 列表。ID 数量上限为20个。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// 监听器协议类型。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 查询负载均衡绑定的后端服务列表，过滤条件如下：
	// <li> location-id - String - 是否必填：否 - （过滤条件）按照 规则ID 过滤，如："loc-12345678"。</li>
	// <li> private-ip-address - String - 是否必填：否 - （过滤条件）按照 后端服务内网IP 过滤，如："172.16.1.1"。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 监听器 ID 列表。ID 数量上限为20个。
	ListenerIds []*string `json:"ListenerIds,omitempty" name:"ListenerIds"`

	// 监听器协议类型。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 查询负载均衡绑定的后端服务列表，过滤条件如下：
	// <li> location-id - String - 是否必填：否 - （过滤条件）按照 规则ID 过滤，如："loc-12345678"。</li>
	// <li> private-ip-address - String - 是否必填：否 - （过滤条件）按照 后端服务内网IP 过滤，如："172.16.1.1"。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Listeners []*ListenerBackend `json:"Listeners,omitempty" name:"Listeners"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 订单ID。
	// 注意：参数TaskId和DealName必须传一个。
	DealName *string `json:"DealName,omitempty" name:"DealName"`
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 请求ID，即接口返回的 RequestId 参数。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 订单ID。
	// 注意：参数TaskId和DealName必须传一个。
	DealName *string `json:"DealName,omitempty" name:"DealName"`
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
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 由负载均衡实例唯一 ID 组成的数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DisassociateTargetGroupsRequestParams struct {
	// 待解绑的规则关系数组。
	Associations []*TargetGroupAssociation `json:"Associations,omitempty" name:"Associations"`
}

type DisassociateTargetGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 待解绑的规则关系数组。
	Associations []*TargetGroupAssociation `json:"Associations,omitempty" name:"Associations"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	L4Clusters []*ClusterItem `json:"L4Clusters,omitempty" name:"L4Clusters"`

	// 7层独占集群列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	L7Clusters []*ClusterItem `json:"L7Clusters,omitempty" name:"L7Clusters"`

	// vpcgw集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassicalCluster *ClusterItem `json:"ClassicalCluster,omitempty" name:"ClassicalCluster"`
}

type ExtraInfo struct {
	// 是否开通VIP直通
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZhiTong *bool `json:"ZhiTong,omitempty" name:"ZhiTong"`

	// TgwGroup名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TgwGroupName *string `json:"TgwGroupName,omitempty" name:"TgwGroupName"`
}

type Filter struct {
	// 过滤器的名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤器的值数组
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type FunctionInfo struct {
	// 函数命名空间
	FunctionNamespace *string `json:"FunctionNamespace,omitempty" name:"FunctionNamespace"`

	// 函数名称
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// 函数的版本名称或别名
	FunctionQualifier *string `json:"FunctionQualifier,omitempty" name:"FunctionQualifier"`

	// 标识 FunctionQualifier 参数的类型，可取值： VERSION（版本）、ALIAS（别名）
	// 注意：此字段可能返回 null，表示取不到有效值。
	FunctionQualifierType *string `json:"FunctionQualifierType,omitempty" name:"FunctionQualifierType"`
}

type FunctionTarget struct {
	// 云函数相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Function *FunctionInfo `json:"Function,omitempty" name:"Function"`

	// 权重
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
}

type HealthCheck struct {
	// 是否开启健康检查：1（开启）、0（关闭）。
	HealthSwitch *int64 `json:"HealthSwitch,omitempty" name:"HealthSwitch"`

	// 健康检查的响应超时时间（仅适用于四层监听器），可选值：2~60，默认值：2，单位：秒。响应超时时间要小于检查间隔时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeOut *int64 `json:"TimeOut,omitempty" name:"TimeOut"`

	// 健康检查探测间隔时间，默认值：5，可选值：5~300，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntervalTime *int64 `json:"IntervalTime,omitempty" name:"IntervalTime"`

	// 健康阈值，默认值：3，表示当连续探测三次健康则表示该转发正常，可选值：2~10，单位：次。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthNum *int64 `json:"HealthNum,omitempty" name:"HealthNum"`

	// 不健康阈值，默认值：3，表示当连续探测三次不健康则表示该转发异常，可选值：2~10，单位：次。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnHealthNum *int64 `json:"UnHealthNum,omitempty" name:"UnHealthNum"`

	// 健康检查状态码（仅适用于HTTP/HTTPS转发规则、TCP监听器的HTTP健康检查方式）。可选值：1~31，默认 31。
	// 1 表示探测后返回值 1xx 代表健康，2 表示返回 2xx 代表健康，4 表示返回 3xx 代表健康，8 表示返回 4xx 代表健康，16 表示返回 5xx 代表健康。若希望多种返回码都可代表健康，则将相应的值相加。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpCode *int64 `json:"HttpCode,omitempty" name:"HttpCode"`

	// 健康检查路径（仅适用于HTTP/HTTPS转发规则、TCP监听器的HTTP健康检查方式）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpCheckPath *string `json:"HttpCheckPath,omitempty" name:"HttpCheckPath"`

	// 健康检查域名（仅适用于HTTP/HTTPS转发规则、TCP监听器的HTTP健康检查方式，当监听器是TCP类型时，该参数为必填项）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpCheckDomain *string `json:"HttpCheckDomain,omitempty" name:"HttpCheckDomain"`

	// 健康检查方法（仅适用于HTTP/HTTPS转发规则、TCP监听器的HTTP健康检查方式），默认值：HEAD，可选值HEAD或GET。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpCheckMethod *string `json:"HttpCheckMethod,omitempty" name:"HttpCheckMethod"`

	// 自定义探测相关参数。健康检查端口，默认为后端服务的端口，除非您希望指定特定端口，否则建议留空。（仅适用于TCP/UDP监听器）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckPort *int64 `json:"CheckPort,omitempty" name:"CheckPort"`

	// 自定义探测相关参数。健康检查协议CheckType的值取CUSTOM时，必填此字段，代表健康检查的输入格式，可取值：HEX或TEXT；取值为HEX时，SendContext和RecvContext的字符只能在0123456789ABCDEF中选取且长度必须是偶数位。（仅适用于TCP/UDP监听器）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContextType *string `json:"ContextType,omitempty" name:"ContextType"`

	// 自定义探测相关参数。健康检查协议CheckType的值取CUSTOM时，必填此字段，代表健康检查发送的请求内容，只允许ASCII可见字符，最大长度限制500。（仅适用于TCP/UDP监听器）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SendContext *string `json:"SendContext,omitempty" name:"SendContext"`

	// 自定义探测相关参数。健康检查协议CheckType的值取CUSTOM时，必填此字段，代表健康检查返回的结果，只允许ASCII可见字符，最大长度限制500。（仅适用于TCP/UDP监听器）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecvContext *string `json:"RecvContext,omitempty" name:"RecvContext"`

	// 自定义探测相关参数。健康检查使用的协议：TCP | HTTP | CUSTOM（仅适用于TCP/UDP监听器，其中UDP监听器只支持CUSTOM；如果使用自定义健康检查功能，则必传）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckType *string `json:"CheckType,omitempty" name:"CheckType"`

	// 自定义探测相关参数。健康检查协议CheckType的值取HTTP时，必传此字段，代表后端服务的HTTP版本：HTTP/1.0、HTTP/1.1；（仅适用于TCP监听器）
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpVersion *string `json:"HttpVersion,omitempty" name:"HttpVersion"`

	// 自定义探测相关参数。健康检查源IP类型：0（使用LB的VIP作为源IP），1（使用100.64网段IP作为源IP），默认值：0
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceIpType *int64 `json:"SourceIpType,omitempty" name:"SourceIpType"`

	// GRPC健康检查状态码（仅适用于后端转发协议为GRPC的规则）。默认值为 12，可输入值为数值、多个数值、或者范围，例如 20 或 20,25 或 0-99
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtendedCode *string `json:"ExtendedCode,omitempty" name:"ExtendedCode"`
}

type IdleLoadBalancer struct {
	// 负载均衡ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡名字
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡所在地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 负载均衡的vip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 闲置原因。NO_RULES：没有规则，NO_RS：有规则没有绑定子机。
	IdleReason *string `json:"IdleReason,omitempty" name:"IdleReason"`

	// 负载均衡实例的状态，包括
	// 0：创建中，1：正常运行。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 负载均衡类型标识，1：负载均衡，0：传统型负载均衡。
	Forward *uint64 `json:"Forward,omitempty" name:"Forward"`
}

type InternetAccessible struct {
	// TRAFFIC_POSTPAID_BY_HOUR 按流量按小时后计费 ; BANDWIDTH_POSTPAID_BY_HOUR 按带宽按小时后计费;
	// BANDWIDTH_PACKAGE 按带宽包计费;
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// 最大出带宽，单位Mbps，范围支持0到2048，仅对公网属性的LB生效，默认值 10
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// 带宽包的类型，如SINGLEISP
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthpkgSubType *string `json:"BandwidthpkgSubType,omitempty" name:"BandwidthpkgSubType"`
}

type LBChargePrepaid struct {
	// 续费类型：AUTO_RENEW 自动续费，  MANUAL_RENEW 手动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`

	// 购买时长，单位：月
	// 注意：此字段可能返回 null，表示取不到有效值。
	Period *int64 `json:"Period,omitempty" name:"Period"`
}

type LBItem struct {
	// lb的字符串id
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// lb的vip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 监听器规则
	Listeners []*ListenerItem `json:"Listeners,omitempty" name:"Listeners"`

	// LB所在地域
	Region *string `json:"Region,omitempty" name:"Region"`
}

type LbRsItem struct {
	// vpc的字符串id，只支持字符串id。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 需要查询后端的内网ip，可以是cvm和弹性网卡。
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`
}

type LbRsTargets struct {
	// 内网ip类型。“cvm”或“eni”
	Type *string `json:"Type,omitempty" name:"Type"`

	// 后端实例的内网ip。
	PrivateIp *string `json:"PrivateIp,omitempty" name:"PrivateIp"`

	// 绑定后端实例的端口。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// rs的vpcId
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// rs的权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type Listener struct {
	// 负载均衡监听器 ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 监听器绑定的证书信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Certificate *CertificateOutput `json:"Certificate,omitempty" name:"Certificate"`

	// 监听器的健康检查信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 请求的调度方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 会话保持时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// 是否开启SNI特性（本参数仅对于HTTPS监听器有意义）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SniSwitch *int64 `json:"SniSwitch,omitempty" name:"SniSwitch"`

	// 监听器下的全部转发规则（本参数仅对于HTTP/HTTPS监听器有意义）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*RuleOutput `json:"Rules,omitempty" name:"Rules"`

	// 监听器的名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器的创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 端口段结束端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndPort *int64 `json:"EndPort,omitempty" name:"EndPort"`

	// 后端服务器类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// 绑定的目标组基本信息；当监听器绑定目标组时，会返回该字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetGroup *BasicTargetGroupInfo `json:"TargetGroup,omitempty" name:"TargetGroup"`

	// 会话保持类型。NORMAL表示默认会话保持类型。QUIC_CID 表示根据Quic Connection ID做会话保持。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionType *string `json:"SessionType,omitempty" name:"SessionType"`

	// 是否开启长连接，1开启，0关闭，（本参数仅对于HTTP/HTTPS监听器有意义）
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeepaliveEnable *int64 `json:"KeepaliveEnable,omitempty" name:"KeepaliveEnable"`

	// 仅支持Nat64 CLB TCP监听器
	// 注意：此字段可能返回 null，表示取不到有效值。
	Toa *bool `json:"Toa,omitempty" name:"Toa"`

	// 解绑后端目标时，是否发RST给客户端，（此参数仅对于TCP监听器有意义）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeregisterTargetRst *bool `json:"DeregisterTargetRst,omitempty" name:"DeregisterTargetRst"`

	// 监听器的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttrFlags []*string `json:"AttrFlags,omitempty" name:"AttrFlags"`

	// 绑定的目标组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetGroupList []*BasicTargetGroupInfo `json:"TargetGroupList,omitempty" name:"TargetGroupList"`

	// 监听器最大连接数，-1表示监听器维度不限速。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxConn *int64 `json:"MaxConn,omitempty" name:"MaxConn"`

	// 监听器最大新增连接数，-1表示监听器维度不限速。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxCps *int64 `json:"MaxCps,omitempty" name:"MaxCps"`
}

type ListenerBackend struct {
	// 监听器 ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器的协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器的端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 监听器下的规则信息（仅适用于HTTP/HTTPS监听器）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*RuleTargets `json:"Rules,omitempty" name:"Rules"`

	// 监听器上绑定的后端服务列表（仅适用于TCP/UDP/TCP_SSL监听器）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Targets []*Backend `json:"Targets,omitempty" name:"Targets"`

	// 若支持端口段，则为端口段结束端口；若不支持端口段，则为0
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndPort *int64 `json:"EndPort,omitempty" name:"EndPort"`
}

type ListenerHealth struct {
	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 监听器的协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器的端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 监听器的转发规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*RuleHealth `json:"Rules,omitempty" name:"Rules"`
}

type ListenerItem struct {
	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器协议
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 绑定规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rules []*RulesItems `json:"Rules,omitempty" name:"Rules"`

	// 四层绑定对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Targets []*LbRsTargets `json:"Targets,omitempty" name:"Targets"`

	// 端口段监听器的结束端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndPort *int64 `json:"EndPort,omitempty" name:"EndPort"`
}

type LoadBalancer struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡实例的名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡实例的网络类型：
	// OPEN：公网属性， INTERNAL：内网属性。
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// 负载均衡类型标识，1：负载均衡，0：传统型负载均衡。
	Forward *uint64 `json:"Forward,omitempty" name:"Forward"`

	// 负载均衡实例的域名，仅公网传统型负载均衡实例才提供该字段。逐步下线中，建议用LoadBalancerDomain替代。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 负载均衡实例的 VIP 列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerVips []*string `json:"LoadBalancerVips,omitempty" name:"LoadBalancerVips"`

	// 负载均衡实例的状态，包括
	// 0：创建中，1：正常运行。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 负载均衡实例的创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 负载均衡实例的上次状态转换时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusTime *string `json:"StatusTime,omitempty" name:"StatusTime"`

	// 负载均衡实例所属的项目 ID， 0 表示默认项目。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 私有网络的 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 高防 LB 的标识，1：高防负载均衡 0：非高防负载均衡。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenBgp *uint64 `json:"OpenBgp,omitempty" name:"OpenBgp"`

	// 在 2016 年 12 月份之前的传统型内网负载均衡都是开启了 snat 的。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Snat *bool `json:"Snat,omitempty" name:"Snat"`

	// 0：表示未被隔离，1：表示被隔离。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Isolation *uint64 `json:"Isolation,omitempty" name:"Isolation"`

	// 用户开启日志的信息，日志只有公网属性创建了 HTTP 、HTTPS 监听器的负载均衡才会有日志。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Log *string `json:"Log,omitempty" name:"Log"`

	// 负载均衡实例所在的子网（仅对内网VPC型LB有意义）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 负载均衡实例的标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// 负载均衡实例的安全组
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecureGroups []*string `json:"SecureGroups,omitempty" name:"SecureGroups"`

	// 负载均衡实例绑定的后端设备的基本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetRegionInfo *TargetRegionInfo `json:"TargetRegionInfo,omitempty" name:"TargetRegionInfo"`

	// anycast负载均衡的发布域，对于非anycast的负载均衡，此字段返回为空字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnycastZone *string `json:"AnycastZone,omitempty" name:"AnycastZone"`

	// IP版本，ipv4 | ipv6
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`

	// 数值形式的私有网络 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumericalVpcId *uint64 `json:"NumericalVpcId,omitempty" name:"NumericalVpcId"`

	// 负载均衡IP地址所属的ISP
	// 注意：此字段可能返回 null，表示取不到有效值。
	VipIsp *string `json:"VipIsp,omitempty" name:"VipIsp"`

	// 主可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterZone *ZoneInfo `json:"MasterZone,omitempty" name:"MasterZone"`

	// 备可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackupZoneSet []*ZoneInfo `json:"BackupZoneSet,omitempty" name:"BackupZoneSet"`

	// 负载均衡实例被隔离的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolatedTime *string `json:"IsolatedTime,omitempty" name:"IsolatedTime"`

	// 负载均衡实例的过期时间，仅对预付费负载均衡生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 负载均衡实例的计费类型，PREPAID：包年包月，POSTPAID_BY_HOUR：按量计费
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// 负载均衡实例的网络属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkAttributes *InternetAccessible `json:"NetworkAttributes,omitempty" name:"NetworkAttributes"`

	// 负载均衡实例的预付费相关属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrepaidAttributes *LBChargePrepaid `json:"PrepaidAttributes,omitempty" name:"PrepaidAttributes"`

	// 负载均衡日志服务(CLS)的日志集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 负载均衡日志服务(CLS)的日志主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogTopicId *string `json:"LogTopicId,omitempty" name:"LogTopicId"`

	// 负载均衡实例的IPv6地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressIPv6 *string `json:"AddressIPv6,omitempty" name:"AddressIPv6"`

	// 暂做保留，一般用户无需关注。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo *ExtraInfo `json:"ExtraInfo,omitempty" name:"ExtraInfo"`

	// 是否可绑定高防包
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDDos *bool `json:"IsDDos,omitempty" name:"IsDDos"`

	// 负载均衡维度的个性化配置ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 后端服务是否放通来自LB的流量
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitempty" name:"LoadBalancerPassToTarget"`

	// 内网独占集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitempty" name:"ExclusiveCluster"`

	// IP地址版本为ipv6时此字段有意义， IPv6Nat64 | IPv6FullChain
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPv6Mode *string `json:"IPv6Mode,omitempty" name:"IPv6Mode"`

	// 是否开启SnatPro。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnatPro *bool `json:"SnatPro,omitempty" name:"SnatPro"`

	// 开启SnatPro负载均衡后，SnatIp列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnatIps []*SnatIp `json:"SnatIps,omitempty" name:"SnatIps"`

	// 性能容量型规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaType *string `json:"SlaType,omitempty" name:"SlaType"`

	// vip是否被封堵
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsBlock *bool `json:"IsBlock,omitempty" name:"IsBlock"`

	// 封堵或解封时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsBlockTime *string `json:"IsBlockTime,omitempty" name:"IsBlockTime"`

	// IP类型是否是本地BGP
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalBgp *bool `json:"LocalBgp,omitempty" name:"LocalBgp"`

	// 7层独占标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterTag *string `json:"ClusterTag,omitempty" name:"ClusterTag"`

	// 开启IPv6FullChain负载均衡7层监听器支持混绑IPv4/IPv6目标功能。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MixIpTarget *bool `json:"MixIpTarget,omitempty" name:"MixIpTarget"`

	// 私有网络内网负载均衡，就近接入模式下规则所落在的可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// CLB是否为NFV，空：不是，l7nfv：七层是NFV。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NfvInfo *string `json:"NfvInfo,omitempty" name:"NfvInfo"`

	// 负载均衡日志服务(CLS)的健康检查日志集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthLogSetId *string `json:"HealthLogSetId,omitempty" name:"HealthLogSetId"`

	// 负载均衡日志服务(CLS)的健康检查日志主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthLogTopicId *string `json:"HealthLogTopicId,omitempty" name:"HealthLogTopicId"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// 负载均衡的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttributeFlags []*string `json:"AttributeFlags,omitempty" name:"AttributeFlags"`

	// 负载均衡实例的域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerDomain *string `json:"LoadBalancerDomain,omitempty" name:"LoadBalancerDomain"`
}

type LoadBalancerDetail struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡实例的名称。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡实例的网络类型：
	// Public：公网属性， Private：内网属性。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// 负载均衡实例的状态，包括
	// 0：创建中，1：正常运行。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 负载均衡实例的 VIP 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitempty" name:"Address"`

	// 负载均衡实例 VIP 的IPv6地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressIPv6 *string `json:"AddressIPv6,omitempty" name:"AddressIPv6"`

	// 负载均衡实例IP版本，IPv4 | IPv6。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" name:"AddressIPVersion"`

	// 负载均衡实例IPv6地址类型，IPv6Nat64 | IPv6FullChain。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IPv6Mode *string `json:"IPv6Mode,omitempty" name:"IPv6Mode"`

	// 负载均衡实例所在可用区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 负载均衡实例IP地址所属的ISP。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressIsp *string `json:"AddressIsp,omitempty" name:"AddressIsp"`

	// 负载均衡实例所属私有网络的 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 负载均衡实例所属的项目 ID， 0 表示默认项目。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 负载均衡实例的创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 负载均衡实例的计费类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`

	// 负载均衡实例的网络属性。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkAttributes *InternetAccessible `json:"NetworkAttributes,omitempty" name:"NetworkAttributes"`

	// 负载均衡实例的预付费相关属性。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrepaidAttributes *LBChargePrepaid `json:"PrepaidAttributes,omitempty" name:"PrepaidAttributes"`

	// 暂做保留，一般用户无需关注。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo *ExtraInfo `json:"ExtraInfo,omitempty" name:"ExtraInfo"`

	// 负载均衡维度的个性化配置ID，多个配置用逗号隔开。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 负载均衡实例的标签信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// 负载均衡监听器 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器协议。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 监听器端口。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 转发规则的 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 转发规则的域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 转发规则的路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 后端目标ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// 后端目标的IP地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetAddress *string `json:"TargetAddress,omitempty" name:"TargetAddress"`

	// 后端目标监听端口。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetPort *uint64 `json:"TargetPort,omitempty" name:"TargetPort"`

	// 后端目标转发权重。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetWeight *uint64 `json:"TargetWeight,omitempty" name:"TargetWeight"`

	// 0：表示未被隔离，1：表示被隔离。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Isolation *uint64 `json:"Isolation,omitempty" name:"Isolation"`

	// 负载均衡绑定的安全组列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecurityGroup []*string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// 负载均衡安全组上移特性是否开启标识。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerPassToTarget *uint64 `json:"LoadBalancerPassToTarget,omitempty" name:"LoadBalancerPassToTarget"`

	// 后端目标健康状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetHealth *string `json:"TargetHealth,omitempty" name:"TargetHealth"`

	// 转发规则的域名列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domains *string `json:"Domains,omitempty" name:"Domains"`

	// 多可用区负载均衡实例所选备区
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaveZone []*string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// 内网负载均衡实例所在可用区，由白名单CLB_Internal_Zone控制
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// 是否开启SNI特性（本参数仅对于HTTPS监听器有意义）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SniSwitch *int64 `json:"SniSwitch,omitempty" name:"SniSwitch"`

	// 负载均衡实例的域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerDomain *string `json:"LoadBalancerDomain,omitempty" name:"LoadBalancerDomain"`
}

type LoadBalancerHealth struct {
	// 负载均衡实例ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 监听器列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Listeners []*ListenerHealth `json:"Listeners,omitempty" name:"Listeners"`
}

type LoadBalancerTraffic struct {
	// 负载均衡ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡名字
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡所在地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 负载均衡的vip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 最大出带宽，单位：Mbps
	OutBandwidth *float64 `json:"OutBandwidth,omitempty" name:"OutBandwidth"`

	// CLB域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

// Predefined struct for user
type ManualRewriteRequestParams struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 源监听器 ID。
	SourceListenerId *string `json:"SourceListenerId,omitempty" name:"SourceListenerId"`

	// 目标监听器 ID。
	TargetListenerId *string `json:"TargetListenerId,omitempty" name:"TargetListenerId"`

	// 转发规则之间的重定向关系。
	RewriteInfos []*RewriteLocationMap `json:"RewriteInfos,omitempty" name:"RewriteInfos"`
}

type ManualRewriteRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 源监听器 ID。
	SourceListenerId *string `json:"SourceListenerId,omitempty" name:"SourceListenerId"`

	// 目标监听器 ID。
	TargetListenerId *string `json:"TargetListenerId,omitempty" name:"TargetListenerId"`

	// 转发规则之间的重定向关系。
	RewriteInfos []*RewriteLocationMap `json:"RewriteInfos,omitempty" name:"RewriteInfos"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// 独占集群信息
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitempty" name:"ExclusiveCluster"`
}

type MigrateClassicalLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// 传统型负载均衡ID数组
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// 独占集群信息
	ExclusiveCluster *ExclusiveCluster `json:"ExclusiveCluster,omitempty" name:"ExclusiveCluster"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type ModifyBlockIPListRequestParams struct {
	// 负载均衡实例ID
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// 操作类型，可取：
	// <li> add_customized_field（首次设置header，开启黑名单功能）</li>
	// <li> set_customized_field（修改header）</li>
	// <li> del_customized_field（删除header）</li>
	// <li> add_blocked（添加黑名单）</li>
	// <li> del_blocked（删除黑名单）</li>
	// <li> flush_blocked（清空黑名单）</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 客户端真实IP存放的header字段名
	ClientIPField *string `json:"ClientIPField,omitempty" name:"ClientIPField"`

	// 封禁IP列表，单次操作数组最大长度支持200000
	BlockIPList []*string `json:"BlockIPList,omitempty" name:"BlockIPList"`

	// 过期时间，单位秒，默认值3600
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 添加IP的策略，可取：fifo（如果黑名单容量已满，新加入黑名单的IP采用先进先出策略）
	AddStrategy *string `json:"AddStrategy,omitempty" name:"AddStrategy"`
}

type ModifyBlockIPListRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// 操作类型，可取：
	// <li> add_customized_field（首次设置header，开启黑名单功能）</li>
	// <li> set_customized_field（修改header）</li>
	// <li> del_customized_field（删除header）</li>
	// <li> add_blocked（添加黑名单）</li>
	// <li> del_blocked（删除黑名单）</li>
	// <li> flush_blocked（清空黑名单）</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 客户端真实IP存放的header字段名
	ClientIPField *string `json:"ClientIPField,omitempty" name:"ClientIPField"`

	// 封禁IP列表，单次操作数组最大长度支持200000
	BlockIPList []*string `json:"BlockIPList,omitempty" name:"BlockIPList"`

	// 过期时间，单位秒，默认值3600
	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 添加IP的策略，可取：fifo（如果黑名单容量已满，新加入黑名单的IP采用先进先出策略）
	AddStrategy *string `json:"AddStrategy,omitempty" name:"AddStrategy"`
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
	JodId *string `json:"JodId,omitempty" name:"JodId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyDomainAttributesRequestParams struct {
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 域名（必须是已经创建的转发规则下的域名），如果是多域名，可以指定多域名列表中的任意一个。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 要修改的新域名。NewDomain和NewDomains只能传一个。
	NewDomain *string `json:"NewDomain,omitempty" name:"NewDomain"`

	// 域名相关的证书信息，注意，仅对启用SNI的监听器适用，不可和MultiCertInfo 同时传入。
	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`

	// 是否开启Http2，注意，只有HTTPS域名才能开启Http2。
	Http2 *bool `json:"Http2,omitempty" name:"Http2"`

	// 是否设为默认域名，注意，一个监听器下只能设置一个默认域名。
	DefaultServer *bool `json:"DefaultServer,omitempty" name:"DefaultServer"`

	// 监听器下必须配置一个默认域名，若要关闭原默认域名，必须同时指定另一个域名作为新的默认域名，如果新的默认域名是多域名，可以指定多域名列表中的任意一个。
	NewDefaultServerDomain *string `json:"NewDefaultServerDomain,omitempty" name:"NewDefaultServerDomain"`

	// 要修改的新域名列表。NewDomain和NewDomains只能传一个。
	NewDomains []*string `json:"NewDomains,omitempty" name:"NewDomains"`

	// 域名相关的证书信息，注意，仅对启用SNI的监听器适用；支持同时传入多本算法类型不同的服务器证书，不可和MultiCertInfo 同时传入。
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitempty" name:"MultiCertInfo"`
}

type ModifyDomainAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 域名（必须是已经创建的转发规则下的域名），如果是多域名，可以指定多域名列表中的任意一个。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 要修改的新域名。NewDomain和NewDomains只能传一个。
	NewDomain *string `json:"NewDomain,omitempty" name:"NewDomain"`

	// 域名相关的证书信息，注意，仅对启用SNI的监听器适用，不可和MultiCertInfo 同时传入。
	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`

	// 是否开启Http2，注意，只有HTTPS域名才能开启Http2。
	Http2 *bool `json:"Http2,omitempty" name:"Http2"`

	// 是否设为默认域名，注意，一个监听器下只能设置一个默认域名。
	DefaultServer *bool `json:"DefaultServer,omitempty" name:"DefaultServer"`

	// 监听器下必须配置一个默认域名，若要关闭原默认域名，必须同时指定另一个域名作为新的默认域名，如果新的默认域名是多域名，可以指定多域名列表中的任意一个。
	NewDefaultServerDomain *string `json:"NewDefaultServerDomain,omitempty" name:"NewDefaultServerDomain"`

	// 要修改的新域名列表。NewDomain和NewDomains只能传一个。
	NewDomains []*string `json:"NewDomains,omitempty" name:"NewDomains"`

	// 域名相关的证书信息，注意，仅对启用SNI的监听器适用；支持同时传入多本算法类型不同的服务器证书，不可和MultiCertInfo 同时传入。
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitempty" name:"MultiCertInfo"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器下的某个旧域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 新域名，	长度限制为：1-120。有三种使用格式：非正则表达式格式，通配符格式，正则表达式格式。非正则表达式格式只能使用字母、数字、‘-’、‘.’。通配符格式的使用 ‘*’ 只能在开头或者结尾。正则表达式以'~'开头。
	NewDomain *string `json:"NewDomain,omitempty" name:"NewDomain"`
}

type ModifyDomainRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 监听器下的某个旧域名。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 新域名，	长度限制为：1-120。有三种使用格式：非正则表达式格式，通配符格式，正则表达式格式。非正则表达式格式只能使用字母、数字、‘-’、‘.’。通配符格式的使用 ‘*’ 只能在开头或者结尾。正则表达式以'~'开头。
	NewDomain *string `json:"NewDomain,omitempty" name:"NewDomain"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 要修改的后端云函数服务列表。
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitempty" name:"FunctionTargets"`

	// 转发规则的ID，当绑定机器到七层转发规则时，必须提供此参数或Domain+Url两者之一。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 目标规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 目标规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitempty" name:"Url"`
}

type ModifyFunctionTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 要修改的后端云函数服务列表。
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitempty" name:"FunctionTargets"`

	// 转发规则的ID，当绑定机器到七层转发规则时，必须提供此参数或Domain+Url两者之一。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 目标规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 目标规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitempty" name:"Url"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyListenerRequestParams struct {
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 新的监听器名称。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 会话保持时间，单位：秒。可选值：30~3600，默认 0，表示不开启。此参数仅适用于TCP/UDP监听器。
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// 健康检查相关参数，此参数仅适用于TCP/UDP/TCP_SSL监听器。
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 证书相关信息，此参数仅适用于HTTPS/TCP_SSL监听器；此参数和MultiCertInfo不能同时传入。
	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`

	// 监听器转发的方式。可选值：WRR、LEAST_CONN
	// 分别表示按权重轮询、最小连接数， 默认为 WRR。
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 是否开启SNI特性，此参数仅适用于HTTPS监听器。注意：未开启SNI的监听器可以开启SNI；已开启SNI的监听器不能关闭SNI。
	SniSwitch *int64 `json:"SniSwitch,omitempty" name:"SniSwitch"`

	// 后端目标类型，NODE表示绑定普通节点，TARGETGROUP表示绑定目标组。
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// 是否开启长连接，此参数仅适用于HTTP/HTTPS监听器。
	KeepaliveEnable *int64 `json:"KeepaliveEnable,omitempty" name:"KeepaliveEnable"`

	// 解绑后端目标时，是否发RST给客户端，此参数仅适用于TCP监听器。
	DeregisterTargetRst *bool `json:"DeregisterTargetRst,omitempty" name:"DeregisterTargetRst"`

	// 会话保持类型。NORMAL表示默认会话保持类型。QUIC_CID表示根据Quic Connection ID做会话保持。QUIC_CID只支持UDP协议。
	SessionType *string `json:"SessionType,omitempty" name:"SessionType"`

	// 证书信息，支持同时传入不同算法类型的多本服务端证书；此参数仅适用于未开启SNI特性的HTTPS监听器。此参数和Certificate不能同时传入。
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitempty" name:"MultiCertInfo"`

	// 监听器粒度并发连接数上限，当前仅性能容量型实例且仅TCP/UDP/TCP_SSL/QUIC监听器支持。取值范围：1-实例规格并发连接上限，其中-1表示关闭监听器粒度并发连接数限速。
	MaxConn *int64 `json:"MaxConn,omitempty" name:"MaxConn"`

	// 监听器粒度新建连接数上限，当前仅性能容量型实例且仅TCP/UDP/TCP_SSL/QUIC监听器支持。取值范围：1-实例规格新建连接上限，其中-1表示关闭监听器粒度新建连接数限速。
	MaxCps *int64 `json:"MaxCps,omitempty" name:"MaxCps"`
}

type ModifyListenerRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 新的监听器名称。
	ListenerName *string `json:"ListenerName,omitempty" name:"ListenerName"`

	// 会话保持时间，单位：秒。可选值：30~3600，默认 0，表示不开启。此参数仅适用于TCP/UDP监听器。
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// 健康检查相关参数，此参数仅适用于TCP/UDP/TCP_SSL监听器。
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 证书相关信息，此参数仅适用于HTTPS/TCP_SSL监听器；此参数和MultiCertInfo不能同时传入。
	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`

	// 监听器转发的方式。可选值：WRR、LEAST_CONN
	// 分别表示按权重轮询、最小连接数， 默认为 WRR。
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 是否开启SNI特性，此参数仅适用于HTTPS监听器。注意：未开启SNI的监听器可以开启SNI；已开启SNI的监听器不能关闭SNI。
	SniSwitch *int64 `json:"SniSwitch,omitempty" name:"SniSwitch"`

	// 后端目标类型，NODE表示绑定普通节点，TARGETGROUP表示绑定目标组。
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// 是否开启长连接，此参数仅适用于HTTP/HTTPS监听器。
	KeepaliveEnable *int64 `json:"KeepaliveEnable,omitempty" name:"KeepaliveEnable"`

	// 解绑后端目标时，是否发RST给客户端，此参数仅适用于TCP监听器。
	DeregisterTargetRst *bool `json:"DeregisterTargetRst,omitempty" name:"DeregisterTargetRst"`

	// 会话保持类型。NORMAL表示默认会话保持类型。QUIC_CID表示根据Quic Connection ID做会话保持。QUIC_CID只支持UDP协议。
	SessionType *string `json:"SessionType,omitempty" name:"SessionType"`

	// 证书信息，支持同时传入不同算法类型的多本服务端证书；此参数仅适用于未开启SNI特性的HTTPS监听器。此参数和Certificate不能同时传入。
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitempty" name:"MultiCertInfo"`

	// 监听器粒度并发连接数上限，当前仅性能容量型实例且仅TCP/UDP/TCP_SSL/QUIC监听器支持。取值范围：1-实例规格并发连接上限，其中-1表示关闭监听器粒度并发连接数限速。
	MaxConn *int64 `json:"MaxConn,omitempty" name:"MaxConn"`

	// 监听器粒度新建连接数上限，当前仅性能容量型实例且仅TCP/UDP/TCP_SSL/QUIC监听器支持。取值范围：1-实例规格新建连接上限，其中-1表示关闭监听器粒度新建连接数限速。
	MaxCps *int64 `json:"MaxCps,omitempty" name:"MaxCps"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyListenerResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 负载均衡的唯一ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡实例名称
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 设置负载均衡跨地域绑定1.0的后端服务信息
	TargetRegionInfo *TargetRegionInfo `json:"TargetRegionInfo,omitempty" name:"TargetRegionInfo"`

	// 网络计费相关参数
	InternetChargeInfo *InternetAccessible `json:"InternetChargeInfo,omitempty" name:"InternetChargeInfo"`

	// Target是否放通来自CLB的流量。开启放通（true）：只验证CLB上的安全组；不开启放通（false）：需同时验证CLB和后端实例上的安全组。
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitempty" name:"LoadBalancerPassToTarget"`

	// 是否开启跨地域绑定2.0功能
	SnatPro *bool `json:"SnatPro,omitempty" name:"SnatPro"`

	// 是否开启删除保护
	DeleteProtect *bool `json:"DeleteProtect,omitempty" name:"DeleteProtect"`
}

type ModifyLoadBalancerAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡的唯一ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡实例名称
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 设置负载均衡跨地域绑定1.0的后端服务信息
	TargetRegionInfo *TargetRegionInfo `json:"TargetRegionInfo,omitempty" name:"TargetRegionInfo"`

	// 网络计费相关参数
	InternetChargeInfo *InternetAccessible `json:"InternetChargeInfo,omitempty" name:"InternetChargeInfo"`

	// Target是否放通来自CLB的流量。开启放通（true）：只验证CLB上的安全组；不开启放通（false）：需同时验证CLB和后端实例上的安全组。
	LoadBalancerPassToTarget *bool `json:"LoadBalancerPassToTarget,omitempty" name:"LoadBalancerPassToTarget"`

	// 是否开启跨地域绑定2.0功能
	SnatPro *bool `json:"SnatPro,omitempty" name:"SnatPro"`

	// 是否开启删除保护
	DeleteProtect *bool `json:"DeleteProtect,omitempty" name:"DeleteProtect"`
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
	delete(f, "SnatPro")
	delete(f, "DeleteProtect")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerAttributesResponseParams struct {
	// 切换负载均衡计费方式时，可用此参数查询切换任务是否成功。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 负载均衡实例ID数组。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// 开启/关闭IPv6FullChain负载均衡7层监听器支持混绑IPv4/IPv6目标特性。
	MixIpTarget *bool `json:"MixIpTarget,omitempty" name:"MixIpTarget"`
}

type ModifyLoadBalancerMixIpTargetRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID数组。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`

	// 开启/关闭IPv6FullChain负载均衡7层监听器支持混绑IPv4/IPv6目标特性。
	MixIpTarget *bool `json:"MixIpTarget,omitempty" name:"MixIpTarget"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LoadBalancerSla []*SlaUpdateParam `json:"LoadBalancerSla,omitempty" name:"LoadBalancerSla"`
}

type ModifyLoadBalancerSlaRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例信息。
	LoadBalancerSla []*SlaUpdateParam `json:"LoadBalancerSla,omitempty" name:"LoadBalancerSla"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerSlaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerSlaResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyRuleRequestParams struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 要修改的转发规则的 ID。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 转发规则的新的转发路径，如不需修改Url，则不需提供此参数。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 健康检查信息。
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 规则的请求转发方式，可选值：WRR、LEAST_CONN、IP_HASH
	// 分别表示按权重轮询、最小连接数、按IP哈希， 默认为 WRR。
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 会话保持时间。
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// 负载均衡实例与后端服务之间的转发协议，默认HTTP，可取值：HTTP、HTTPS、TRPC。
	ForwardType *string `json:"ForwardType,omitempty" name:"ForwardType"`

	// TRPC被调服务器路由，ForwardType为TRPC时必填。
	TrpcCallee *string `json:"TrpcCallee,omitempty" name:"TrpcCallee"`

	// TRPC调用服务接口，ForwardType为TRPC时必填。
	TrpcFunc *string `json:"TrpcFunc,omitempty" name:"TrpcFunc"`
}

type ModifyRuleRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 要修改的转发规则的 ID。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 转发规则的新的转发路径，如不需修改Url，则不需提供此参数。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 健康检查信息。
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 规则的请求转发方式，可选值：WRR、LEAST_CONN、IP_HASH
	// 分别表示按权重轮询、最小连接数、按IP哈希， 默认为 WRR。
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 会话保持时间。
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// 负载均衡实例与后端服务之间的转发协议，默认HTTP，可取值：HTTP、HTTPS、TRPC。
	ForwardType *string `json:"ForwardType,omitempty" name:"ForwardType"`

	// TRPC被调服务器路由，ForwardType为TRPC时必填。
	TrpcCallee *string `json:"TrpcCallee,omitempty" name:"TrpcCallee"`

	// TRPC调用服务接口，ForwardType为TRPC时必填。
	TrpcFunc *string `json:"TrpcFunc,omitempty" name:"TrpcFunc"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyTargetGroupAttributeRequestParams struct {
	// 目标组的ID。
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// 目标组的新名称。
	TargetGroupName *string `json:"TargetGroupName,omitempty" name:"TargetGroupName"`

	// 目标组的新默认端口。
	Port *uint64 `json:"Port,omitempty" name:"Port"`
}

type ModifyTargetGroupAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 目标组的ID。
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// 目标组的新名称。
	TargetGroupName *string `json:"TargetGroupName,omitempty" name:"TargetGroupName"`

	// 目标组的新默认端口。
	Port *uint64 `json:"Port,omitempty" name:"Port"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetGroupAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetGroupAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// 待修改端口的服务器数组。
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances"`
}

type ModifyTargetGroupInstancesPortRequest struct {
	*tchttp.BaseRequest
	
	// 目标组ID。
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// 待修改端口的服务器数组。
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// 待修改权重的服务器数组。
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances"`
}

type ModifyTargetGroupInstancesWeightRequest struct {
	*tchttp.BaseRequest
	
	// 目标组ID。
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// 待修改权重的服务器数组。
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 要修改端口的后端服务列表。
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// 后端服务绑定到监听器或转发规则的新端口。
	NewPort *int64 `json:"NewPort,omitempty" name:"NewPort"`

	// 转发规则的ID，当后端服务绑定到七层转发规则时，必须提供此参数或Domain+Url两者之一。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 目标规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 目标规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitempty" name:"Url"`
}

type ModifyTargetPortRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 要修改端口的后端服务列表。
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// 后端服务绑定到监听器或转发规则的新端口。
	NewPort *int64 `json:"NewPort,omitempty" name:"NewPort"`

	// 转发规则的ID，当后端服务绑定到七层转发规则时，必须提供此参数或Domain+Url两者之一。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 目标规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 目标规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitempty" name:"Url"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发规则的ID，当绑定机器到七层转发规则时，必须提供此参数或Domain+Url两者之一。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 目标规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 目标规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 要修改权重的后端服务列表。
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// 后端服务新的转发权重，取值范围：0~100，默认值10。如果设置了 Targets.Weight 参数，则此参数不生效。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type ModifyTargetWeightRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发规则的ID，当绑定机器到七层转发规则时，必须提供此参数或Domain+Url两者之一。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 目标规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 目标规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 要修改权重的后端服务列表。
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// 后端服务新的转发权重，取值范围：0~100，默认值10。如果设置了 Targets.Weight 参数，则此参数不生效。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
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
	delete(f, "LocationId")
	delete(f, "Domain")
	delete(f, "Url")
	delete(f, "Targets")
	delete(f, "Weight")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetWeightRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetWeightResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type MultiCertInfo struct {
	// 认证类型，UNIDIRECTIONAL：单向认证，MUTUAL：双向认证
	SSLMode *string `json:"SSLMode,omitempty" name:"SSLMode"`

	// 监听器或规则证书列表，单双向认证，多本服务端证书算法类型不能重复;若SSLMode为双向认证，证书列表必须包含一本ca证书。
	CertList []*CertInfo `json:"CertList,omitempty" name:"CertList"`
}

type Quota struct {
	// 配额名称，取值范围：
	// <li> TOTAL_OPEN_CLB_QUOTA：用户当前地域下的公网CLB配额 </li>
	// <li> TOTAL_INTERNAL_CLB_QUOTA：用户当前地域下的内网CLB配额 </li>
	// <li> TOTAL_LISTENER_QUOTA：一个CLB下的监听器配额 </li>
	// <li> TOTAL_LISTENER_RULE_QUOTA：一个监听器下的转发规则配额 </li>
	// <li> TOTAL_TARGET_BIND_QUOTA：一条转发规则下可绑定设备的配额 </li>
	// <li> TOTAL_SNAP_IP_QUOTA： 一个CLB实例下跨地域2.0的SNAT IP配额 </li>
	// <li>TOTAL_ISP_CLB_QUOTA：用户当前地域下的三网CLB配额 </li>
	QuotaId *string `json:"QuotaId,omitempty" name:"QuotaId"`

	// 当前使用数量，为 null 时表示无意义。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuotaCurrent *int64 `json:"QuotaCurrent,omitempty" name:"QuotaCurrent"`

	// 配额数量。
	QuotaLimit *int64 `json:"QuotaLimit,omitempty" name:"QuotaLimit"`
}

// Predefined struct for user
type RegisterFunctionTargetsRequestParams struct {
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 待绑定的云函数列表。
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitempty" name:"FunctionTargets"`

	// 目标转发规则的 ID，当将云函数绑定到七层转发规则时，必须输入此参数或 Domain+Url 参数。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 目标转发规则的域名，若已经输入 LocationId 参数，则本参数不生效。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 目标转发规则的 URL，若已经输入 LocationId 参数，则本参数不生效。
	Url *string `json:"Url,omitempty" name:"Url"`
}

type RegisterFunctionTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器 ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 待绑定的云函数列表。
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitempty" name:"FunctionTargets"`

	// 目标转发规则的 ID，当将云函数绑定到七层转发规则时，必须输入此参数或 Domain+Url 参数。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 目标转发规则的域名，若已经输入 LocationId 参数，则本参数不生效。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 目标转发规则的 URL，若已经输入 LocationId 参数，则本参数不生效。
	Url *string `json:"Url,omitempty" name:"Url"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type RegisterTargetGroupInstancesRequestParams struct {
	// 目标组ID
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// 服务器实例数组
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances"`
}

type RegisterTargetGroupInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 目标组ID
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// 服务器实例数组
	TargetGroupInstances []*TargetGroupInstance `json:"TargetGroupInstances,omitempty" name:"TargetGroupInstances"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 待绑定的后端服务列表，数组长度最大支持20。
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// 转发规则的ID，当绑定后端服务到七层转发规则时，必须提供此参数或Domain+Url两者之一。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 目标转发规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 目标转发规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitempty" name:"Url"`
}

type RegisterTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡监听器ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 待绑定的后端服务列表，数组长度最大支持20。
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// 转发规则的ID，当绑定后端服务到七层转发规则时，必须提供此参数或Domain+Url两者之一。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 目标转发规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 目标转发规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitempty" name:"Url"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 后端服务信息。
	Targets []*ClassicalTargetInfo `json:"Targets,omitempty" name:"Targets"`
}

type RegisterTargetsWithClassicalLBRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 后端服务信息。
	Targets []*ClassicalTargetInfo `json:"Targets,omitempty" name:"Targets"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ReplaceCertForLoadBalancersRequestParams struct {
	// 需要被替换的证书的ID，可以是服务端证书或客户端证书
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`

	// 新证书的内容等相关信息
	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`
}

type ReplaceCertForLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// 需要被替换的证书的ID，可以是服务端证书或客户端证书
	OldCertificateId *string `json:"OldCertificateId,omitempty" name:"OldCertificateId"`

	// 新证书的内容等相关信息
	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Type []*string `json:"Type,omitempty" name:"Type"`

	// 运营商信息，如"CMCC", "CUCC", "CTCC", "BGP", "INTERNAL"。
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// 可用资源。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailabilitySet []*ResourceAvailability `json:"AvailabilitySet,omitempty" name:"AvailabilitySet"`
}

type ResourceAvailability struct {
	// 运营商内具体资源信息，如"CMCC", "CUCC", "CTCC", "BGP"。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 资源可用性，"Available"：可用，"Unavailable"：不可用
	Availability *string `json:"Availability,omitempty" name:"Availability"`
}

type RewriteLocationMap struct {
	// 源转发规则ID
	SourceLocationId *string `json:"SourceLocationId,omitempty" name:"SourceLocationId"`

	// 重定向至的目标转发规则ID
	TargetLocationId *string `json:"TargetLocationId,omitempty" name:"TargetLocationId"`

	// 重定向状态码，可取值301,302,307
	RewriteCode *int64 `json:"RewriteCode,omitempty" name:"RewriteCode"`

	// 重定向是否携带匹配的url，配置RewriteCode时必填
	TakeUrl *bool `json:"TakeUrl,omitempty" name:"TakeUrl"`

	// 源转发的域名，必须是SourceLocationId对应的域名，配置RewriteCode时必填
	SourceDomain *string `json:"SourceDomain,omitempty" name:"SourceDomain"`
}

type RewriteTarget struct {
	// 重定向目标的监听器ID
	// 注意：此字段可能返回 null，表示无重定向。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetListenerId *string `json:"TargetListenerId,omitempty" name:"TargetListenerId"`

	// 重定向目标的转发规则ID
	// 注意：此字段可能返回 null，表示无重定向。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetLocationId *string `json:"TargetLocationId,omitempty" name:"TargetLocationId"`

	// 重定向状态码
	// 注意：此字段可能返回 null，表示取不到有效值。
	RewriteCode *int64 `json:"RewriteCode,omitempty" name:"RewriteCode"`

	// 重定向是否携带匹配的url
	// 注意：此字段可能返回 null，表示取不到有效值。
	TakeUrl *bool `json:"TakeUrl,omitempty" name:"TakeUrl"`

	// 重定向类型，Manual: 手动重定向，Auto:  自动重定向
	// 注意：此字段可能返回 null，表示取不到有效值。
	RewriteType *string `json:"RewriteType,omitempty" name:"RewriteType"`
}

type RsWeightRule struct {
	// 负载均衡监听器 ID。
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 要修改权重的后端机器列表。
	Targets []*Target `json:"Targets,omitempty" name:"Targets"`

	// 转发规则的ID，七层规则时需要此参数，4层规则不需要。
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 目标规则的域名，提供LocationId参数时本参数不生效。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 目标规则的URL，提供LocationId参数时本参数不生效。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 后端服务修改后的转发权重，取值范围：[0，100]。此参数的优先级低于前述[Target](https://cloud.tencent.com/document/api/214/30694#Target)中的Weight参数，即最终的权重值以Target中的Weight参数值为准，仅当Target中的Weight参数为空时，才以RsWeightRule中的Weight参数为准。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type RuleHealth struct {
	// 转发规则ID
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 转发规则的域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 转发规则的Url
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 本规则上绑定的后端服务的健康检查状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Targets []*TargetHealth `json:"Targets,omitempty" name:"Targets"`
}

type RuleInput struct {
	// 转发规则的路径。长度限制为：1~200。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 转发规则的域名。长度限制为：1~80。Domain和Domains只需要传一个，单域名规则传Domain，多域名规则传Domains。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 会话保持时间。设置为0表示关闭会话保持，开启会话保持可取值30~3600，单位：秒。
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// 健康检查信息。详情请参见：[健康检查](https://cloud.tencent.com/document/product/214/6097)
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 证书信息；此参数和MultiCertInfo不能同时传入。
	Certificate *CertificateInput `json:"Certificate,omitempty" name:"Certificate"`

	// 规则的请求转发方式，可选值：WRR、LEAST_CONN、IP_HASH
	// 分别表示按权重轮询、最小连接数、按IP哈希， 默认为 WRR。
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 负载均衡与后端服务之间的转发协议，目前支持 HTTP/HTTPS/TRPC
	ForwardType *string `json:"ForwardType,omitempty" name:"ForwardType"`

	// 是否将该域名设为默认域名，注意，一个监听器下只能设置一个默认域名。
	DefaultServer *bool `json:"DefaultServer,omitempty" name:"DefaultServer"`

	// 是否开启Http2，注意，只有HTTPS域名才能开启Http2。
	Http2 *bool `json:"Http2,omitempty" name:"Http2"`

	// 后端目标类型，NODE表示绑定普通节点，TARGETGROUP表示绑定目标组
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// TRPC被调服务器路由，ForwardType为TRPC时必填
	TrpcCallee *string `json:"TrpcCallee,omitempty" name:"TrpcCallee"`

	// TRPC调用服务接口，ForwardType为TRPC时必填
	TrpcFunc *string `json:"TrpcFunc,omitempty" name:"TrpcFunc"`

	// 是否开启QUIC，注意，只有HTTPS域名才能开启QUIC
	Quic *bool `json:"Quic,omitempty" name:"Quic"`

	// 转发规则的域名列表。每个域名的长度限制为：1~80。Domain和Domains只需要传一个，单域名规则传Domain，多域名规则传Domains。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 证书信息，支持同时传入不同算法类型的多本服务端证书；此参数和Certificate不能同时传入。
	MultiCertInfo *MultiCertInfo `json:"MultiCertInfo,omitempty" name:"MultiCertInfo"`
}

type RuleOutput struct {
	// 转发规则的 ID
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 转发规则的域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 转发规则的路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 会话保持时间
	SessionExpireTime *int64 `json:"SessionExpireTime,omitempty" name:"SessionExpireTime"`

	// 健康检查信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty" name:"HealthCheck"`

	// 证书信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Certificate *CertificateOutput `json:"Certificate,omitempty" name:"Certificate"`

	// 规则的请求转发方式
	Scheduler *string `json:"Scheduler,omitempty" name:"Scheduler"`

	// 转发规则所属的监听器 ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 转发规则的重定向目标信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RewriteTarget *RewriteTarget `json:"RewriteTarget,omitempty" name:"RewriteTarget"`

	// 是否开启gzip
	HttpGzip *bool `json:"HttpGzip,omitempty" name:"HttpGzip"`

	// 转发规则是否为自动创建
	BeAutoCreated *bool `json:"BeAutoCreated,omitempty" name:"BeAutoCreated"`

	// 是否作为默认域名
	DefaultServer *bool `json:"DefaultServer,omitempty" name:"DefaultServer"`

	// 是否开启Http2
	Http2 *bool `json:"Http2,omitempty" name:"Http2"`

	// 负载均衡与后端服务之间的转发协议
	ForwardType *string `json:"ForwardType,omitempty" name:"ForwardType"`

	// 转发规则的创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 后端服务器类型
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// 绑定的目标组基本信息；当规则绑定目标组时，会返回该字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetGroup *BasicTargetGroupInfo `json:"TargetGroup,omitempty" name:"TargetGroup"`

	// WAF实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WafDomainId *string `json:"WafDomainId,omitempty" name:"WafDomainId"`

	// TRPC被调服务器路由，ForwardType为TRPC时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrpcCallee *string `json:"TrpcCallee,omitempty" name:"TrpcCallee"`

	// TRPC调用服务接口，ForwardType为TRPC时有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrpcFunc *string `json:"TrpcFunc,omitempty" name:"TrpcFunc"`

	// QUIC状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuicStatus *string `json:"QuicStatus,omitempty" name:"QuicStatus"`

	// 转发规则的域名列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 绑定的目标组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetGroupList []*BasicTargetGroupInfo `json:"TargetGroupList,omitempty" name:"TargetGroupList"`
}

type RuleTargets struct {
	// 转发规则的 ID
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 转发规则的域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 转发规则的路径。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 后端服务的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Targets []*Backend `json:"Targets,omitempty" name:"Targets"`

	// 后端云函数的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FunctionTargets []*FunctionTarget `json:"FunctionTargets,omitempty" name:"FunctionTargets"`
}

type RulesItems struct {
	// 规则id
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`

	// 域名
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// uri
	Url *string `json:"Url,omitempty" name:"Url"`

	// 绑定的后端对象
	Targets []*LbRsTargets `json:"Targets,omitempty" name:"Targets"`
}

// Predefined struct for user
type SetCustomizedConfigForLoadBalancerRequestParams struct {
	// 操作类型：'ADD', 'DELETE', 'UPDATE', 'BIND', 'UNBIND'
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 除了创建个性化配置外，必传此字段，如：pz-1234abcd
	UconfigId *string `json:"UconfigId,omitempty" name:"UconfigId"`

	// 创建个性化配置或修改个性化配置的内容时，必传此字段
	ConfigContent *string `json:"ConfigContent,omitempty" name:"ConfigContent"`

	// 创建个性化配置或修改个性化配置的名字时，必传此字段
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 绑定解绑时，必传此字段
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

type SetCustomizedConfigForLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 操作类型：'ADD', 'DELETE', 'UPDATE', 'BIND', 'UNBIND'
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 除了创建个性化配置外，必传此字段，如：pz-1234abcd
	UconfigId *string `json:"UconfigId,omitempty" name:"UconfigId"`

	// 创建个性化配置或修改个性化配置的内容时，必传此字段
	ConfigContent *string `json:"ConfigContent,omitempty" name:"ConfigContent"`

	// 创建个性化配置或修改个性化配置的名字时，必传此字段
	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`

	// 绑定解绑时，必传此字段
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetCustomizedConfigForLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetCustomizedConfigForLoadBalancerResponseParams struct {
	// 个性化配置ID，如：pz-1234abcd
	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 日志服务(CLS)的日志集 ID。
	// <li>增加和更新日志主题时可调用 [DescribeLogsets](https://cloud.tencent.com/document/product/614/56454) 接口获取日志集 ID。</li>
	// <li>删除日志主题时，此参数填写为null即可。</li>
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 日志服务(CLS)的日志主题 ID。
	// <li>增加和更新日志主题时可调用 [DescribeTopics](https://cloud.tencent.com/document/product/614/58624) 接口获取日志主题 ID。</li>
	// <li>删除日志主题时，此参数填写为null即可。</li>
	LogTopicId *string `json:"LogTopicId,omitempty" name:"LogTopicId"`

	// 日志类型：
	// <li>ACCESS：访问日志</li>
	// <li>HEALTH：健康检查日志</li>
	// 默认为ACCESS。
	LogType *string `json:"LogType,omitempty" name:"LogType"`
}

type SetLoadBalancerClsLogRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID。
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 日志服务(CLS)的日志集 ID。
	// <li>增加和更新日志主题时可调用 [DescribeLogsets](https://cloud.tencent.com/document/product/614/56454) 接口获取日志集 ID。</li>
	// <li>删除日志主题时，此参数填写为null即可。</li>
	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`

	// 日志服务(CLS)的日志主题 ID。
	// <li>增加和更新日志主题时可调用 [DescribeTopics](https://cloud.tencent.com/document/product/614/58624) 接口获取日志主题 ID。</li>
	// <li>删除日志主题时，此参数填写为null即可。</li>
	LogTopicId *string `json:"LogTopicId,omitempty" name:"LogTopicId"`

	// 日志类型：
	// <li>ACCESS：访问日志</li>
	// <li>HEALTH：健康检查日志</li>
	// 默认为ACCESS。
	LogType *string `json:"LogType,omitempty" name:"LogType"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 安全组ID构成的数组，一个负载均衡实例最多可绑定50个安全组，如果要解绑所有安全组，可不传此参数，或传入空数组。
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`
}

type SetLoadBalancerSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 安全组ID构成的数组，一个负载均衡实例最多可绑定50个安全组，如果要解绑所有安全组，可不传此参数，或传入空数组。
	SecurityGroups []*string `json:"SecurityGroups,omitempty" name:"SecurityGroups"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type SetSecurityGroupForLoadbalancersRequestParams struct {
	// 安全组ID，如 sg-12345678
	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// ADD 绑定安全组；
	// DEL 解绑安全组
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 负载均衡实例ID数组
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
}

type SetSecurityGroupForLoadbalancersRequest struct {
	*tchttp.BaseRequest
	
	// 安全组ID，如 sg-12345678
	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`

	// ADD 绑定安全组；
	// DEL 解绑安全组
	OperationType *string `json:"OperationType,omitempty" name:"OperationType"`

	// 负载均衡实例ID数组
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// lb的字符串ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 升级为性能容量型，固定取值为SLA。SLA表示升级为默认规格的性能容量型实例。
	// <ul><li>当您开通了普通规格的性能容量型时，SLA对应超强型1规格。普通规格的性能容量型正在内测中，请提交 [内测申请](https://cloud.tencent.com/apply/p/hf45esx99lf)。</li>
	// <li>当您开通了超大型规格的性能容量型时，SLA对应超强型4规格。超大型规格的性能容量型正在内测中，请提交 [工单申请](https://console.cloud.tencent.com/workorder/category)。</li></ul>
	SlaType *string `json:"SlaType,omitempty" name:"SlaType"`
}

type SnatIp struct {
	// 私有网络子网的唯一性id，如subnet-12345678
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// IP地址，如192.168.0.1
	Ip *string `json:"Ip,omitempty" name:"Ip"`
}

type TagInfo struct {
	// 标签的键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type Target struct {
	// 后端服务的监听端口。
	// 注意：绑定CVM（云服务器）或ENI（弹性网卡）时必传此参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 后端服务的类型，可取：CVM（云服务器）、ENI（弹性网卡）；作为入参时，目前本参数暂不生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 绑定CVM时需要传入此参数，代表CVM的唯一 ID，可通过 DescribeInstances 接口返回字段中的 InstanceId 字段获取。表示绑定主网卡主IP。
	// 注意：参数 InstanceId、EniIp 有且只能传入其中一个参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 后端服务修改后的转发权重，取值范围：[0, 100]，默认为 10。此参数的优先级高于[RsWeightRule](https://cloud.tencent.com/document/api/214/30694#RsWeightRule)中的Weight参数，即最终的权重值以此Weight参数值为准，仅当此Weight参数为空时，才以RsWeightRule中的Weight参数为准。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 绑定IP时需要传入此参数，支持弹性网卡的IP和其他内网IP，如果是弹性网卡则必须先绑定至CVM，然后才能绑定到负载均衡实例。
	// 注意：参数 InstanceId、EniIp 有且只能传入其中一个参数。如果绑定双栈IPV6子机，则必须传该参数。如果是跨地域绑定，则必须传该参数，不支持传InstanceId参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EniIp *string `json:"EniIp,omitempty" name:"EniIp"`
}

type TargetGroupAssociation struct {
	// 负载均衡ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 监听器ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 目标组ID
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// 转发规则ID
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
}

type TargetGroupBackend struct {
	// 目标组ID
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// 后端服务的类型，可取：CVM、ENI（即将支持）
	Type *string `json:"Type,omitempty" name:"Type"`

	// 后端服务的唯一 ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 后端服务的监听端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 后端服务的转发权重，取值范围：[0, 100]，默认为 10。
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// 后端服务的外网 IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`

	// 后端服务的内网 IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`

	// 后端服务的实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 后端服务被绑定的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegisteredTime *string `json:"RegisteredTime,omitempty" name:"RegisteredTime"`

	// 弹性网卡唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EniId *string `json:"EniId,omitempty" name:"EniId"`

	// 后端服务的可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type TargetGroupInfo struct {
	// 目标组ID
	TargetGroupId *string `json:"TargetGroupId,omitempty" name:"TargetGroupId"`

	// 目标组的vpcid
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 目标组的名字
	TargetGroupName *string `json:"TargetGroupName,omitempty" name:"TargetGroupName"`

	// 目标组的默认端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 目标组的创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 目标组的修改时间
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 关联到的规则数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociatedRule []*AssociationItem `json:"AssociatedRule,omitempty" name:"AssociatedRule"`
}

type TargetGroupInstance struct {
	// 目标组实例的内网IP
	BindIP *string `json:"BindIP,omitempty" name:"BindIP"`

	// 目标组实例的端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 目标组实例的权重
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`

	// 目标组实例的新端口
	NewPort *uint64 `json:"NewPort,omitempty" name:"NewPort"`
}

type TargetHealth struct {
	// Target的内网IP
	IP *string `json:"IP,omitempty" name:"IP"`

	// Target绑定的端口
	Port *int64 `json:"Port,omitempty" name:"Port"`

	// 当前健康状态，true：健康，false：不健康（包括尚未开始探测、探测中、状态异常等几种状态）。只有处于健康状态（且权重大于0），负载均衡才会向其转发流量。
	HealthStatus *bool `json:"HealthStatus,omitempty" name:"HealthStatus"`

	// Target的实例ID，如 ins-12345678
	TargetId *string `json:"TargetId,omitempty" name:"TargetId"`

	// 当前健康状态的详细信息。如：Alive、Dead、Unknown。Alive状态为健康，Dead状态为异常，Unknown状态包括尚未开始探测、探测中、状态未知。
	HealthStatusDetail *string `json:"HealthStatusDetail,omitempty" name:"HealthStatusDetail"`

	// 当前健康状态的详细信息。如：Alive、Dead、Unknown。Alive状态为健康，Dead状态为异常，Unknown状态包括尚未开始探测、探测中、状态未知。(该参数对象即将下线，不推荐使用，请使用HealthStatusDetail获取健康详情)
	HealthStatusDetial *string `json:"HealthStatusDetial,omitempty" name:"HealthStatusDetial"`
}

type TargetRegionInfo struct {
	// Target所属地域，如 ap-guangzhou
	Region *string `json:"Region,omitempty" name:"Region"`

	// Target所属网络，私有网络格式如 vpc-abcd1234，如果是基础网络，则为"0"
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type ZoneInfo struct {
	// 可用区数值形式的唯一ID，如：100001
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 可用区字符串形式的唯一ID，如：ap-guangzhou-1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 可用区名称，如：广州一区
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 可用区所属地域，如：ap-guangzhou
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneRegion *string `json:"ZoneRegion,omitempty" name:"ZoneRegion"`

	// 可用区是否是LocalZone可用区，如：false
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalZone *bool `json:"LocalZone,omitempty" name:"LocalZone"`

	// 可用区是否是EdgeZone可用区，如：false
	// 注意：此字段可能返回 null，表示取不到有效值。
	EdgeZone *bool `json:"EdgeZone,omitempty" name:"EdgeZone"`
}

type ZoneResource struct {
	// 主可用区，如"ap-guangzhou-1"。
	MasterZone *string `json:"MasterZone,omitempty" name:"MasterZone"`

	// 资源列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceSet []*Resource `json:"ResourceSet,omitempty" name:"ResourceSet"`

	// 备可用区，如"ap-guangzhou-2"，单可用区时，备可用区为null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaveZone *string `json:"SlaveZone,omitempty" name:"SlaveZone"`

	// IP版本，如IPv4，IPv6，IPv6_Nat。
	IPVersion *string `json:"IPVersion,omitempty" name:"IPVersion"`

	// 可用区所属地域，如：ap-guangzhou
	ZoneRegion *string `json:"ZoneRegion,omitempty" name:"ZoneRegion"`

	// 可用区是否是LocalZone可用区，如：false
	LocalZone *bool `json:"LocalZone,omitempty" name:"LocalZone"`

	// 可用区资源的类型，SHARED表示共享资源，EXCLUSIVE表示独占资源。
	ZoneResourceType *string `json:"ZoneResourceType,omitempty" name:"ZoneResourceType"`

	// 可用区是否是EdgeZone可用区，如：false
	EdgeZone *bool `json:"EdgeZone,omitempty" name:"EdgeZone"`
}