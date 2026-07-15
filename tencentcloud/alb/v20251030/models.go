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

package v20251030

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AccessLogConfig struct {
	// 负载均衡日志服务(CLS)的日志集ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogSetId *string `json:"LogSetId,omitnil,omitempty" name:"LogSetId"`

	// 负载均衡日志服务(CLS)的日志主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogTopicId *string `json:"LogTopicId,omitnil,omitempty" name:"LogTopicId"`
}

// Predefined struct for user
type AddTargetsToTargetGroupRequestParams struct {
	// 目标组 ID，格式为 lbtg- 后接 8 位字母数字。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 需要添加至目标组的后端服务列表。单次请求最多支持添加 **50** 个后端服务。
	Targets []*TargetToAdd `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 是否预览此次请求。 
	// - **false**（默认）：发送普通请求，直接添加后端服务至目标组。 
	// - **true**：发送预览请求，检查添加后端服务的参数、格式、业务限制等是否符合要求。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type AddTargetsToTargetGroupRequest struct {
	*tchttp.BaseRequest
	
	// 目标组 ID，格式为 lbtg- 后接 8 位字母数字。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 需要添加至目标组的后端服务列表。单次请求最多支持添加 **50** 个后端服务。
	Targets []*TargetToAdd `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 是否预览此次请求。 
	// - **false**（默认）：发送普通请求，直接添加后端服务至目标组。 
	// - **true**：发送预览请求，检查添加后端服务的参数、格式、业务限制等是否符合要求。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *AddTargetsToTargetGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddTargetsToTargetGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "Targets")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddTargetsToTargetGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddTargetsToTargetGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddTargetsToTargetGroupResponse struct {
	*tchttp.BaseResponse
	Response *AddTargetsToTargetGroupResponseParams `json:"Response"`
}

func (r *AddTargetsToTargetGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddTargetsToTargetGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateBandwidthPackageWithLoadBalancerRequestParams struct {
	// 共享带宽包 ID。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 客户端Token，用于保证请求的幂等性。
	// 
	// 从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符。
	// 
	// > 若您未指定，则系统自动使用API请求的**RequestId**作为**ClientToken**标识。每次API请求的**RequestId**不一样。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否只预检此次请求。取值：
	// - **true**：发送检查请求，不会将共享带宽包绑定到负载均衡实例。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回错误码`DryRunOperation`。
	// - **false**（默认值）：发送正常请求，通过检查后返回HTTP 2xx状态码并直接进行操作。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type AssociateBandwidthPackageWithLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 共享带宽包 ID。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 客户端Token，用于保证请求的幂等性。
	// 
	// 从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符。
	// 
	// > 若您未指定，则系统自动使用API请求的**RequestId**作为**ClientToken**标识。每次API请求的**RequestId**不一样。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否只预检此次请求。取值：
	// - **true**：发送检查请求，不会将共享带宽包绑定到负载均衡实例。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回错误码`DryRunOperation`。
	// - **false**（默认值）：发送正常请求，通过检查后返回HTTP 2xx状态码并直接进行操作。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *AssociateBandwidthPackageWithLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateBandwidthPackageWithLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BandwidthPackageId")
	delete(f, "LoadBalancerId")
	delete(f, "ClientToken")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateBandwidthPackageWithLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateBandwidthPackageWithLoadBalancerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateBandwidthPackageWithLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *AssociateBandwidthPackageWithLoadBalancerResponseParams `json:"Response"`
}

func (r *AssociateBandwidthPackageWithLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateBandwidthPackageWithLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateListenerAdditionalCertificatesRequestParams struct {
	// 扩展证书 ID 列表。
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 客户端 Token，用于保证请求的幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken 只支持 ASCII 字符。
	// 若您未指定，则系统自动使用 API 请求的 RequestId 作为 ClientToken 标识。每次 API 请求的 RequestId 不一样。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否只预检此次请求，取值：
	// true：发送检查请求，不会为HTTPS和QUIC监听器添加扩展证书。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回错误码DryRunOperation。
	// false（默认值）：发送正常请求，通过检查后返回HTTP2xx状态码并直接进行操作。
	DryRun *string `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type AssociateListenerAdditionalCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// 扩展证书 ID 列表。
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 客户端 Token，用于保证请求的幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken 只支持 ASCII 字符。
	// 若您未指定，则系统自动使用 API 请求的 RequestId 作为 ClientToken 标识。每次 API 请求的 RequestId 不一样。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否只预检此次请求，取值：
	// true：发送检查请求，不会为HTTPS和QUIC监听器添加扩展证书。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回错误码DryRunOperation。
	// false（默认值）：发送正常请求，通过检查后返回HTTP2xx状态码并直接进行操作。
	DryRun *string `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *AssociateListenerAdditionalCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateListenerAdditionalCertificatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateIds")
	delete(f, "ListenerId")
	delete(f, "LoadBalancerId")
	delete(f, "ClientToken")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateListenerAdditionalCertificatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateListenerAdditionalCertificatesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateListenerAdditionalCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *AssociateListenerAdditionalCertificatesResponseParams `json:"Response"`
}

func (r *AssociateListenerAdditionalCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateListenerAdditionalCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CertificateInfo struct {
	// 证书绑定时间。
	AssociatedTime *string `json:"AssociatedTime,omitnil,omitempty" name:"AssociatedTime"`

	// 证书 ID。
	CertificateId *string `json:"CertificateId,omitnil,omitempty" name:"CertificateId"`

	// 证书类型。取值：CA或SVR（服务器证书）。
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// 是否为监听器默认证书。取值：
	// true：默认证书。
	// false：扩展证书。
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 证书与监听器的绑定状态。取值：Associated（已关联）、Associating（关联中）、Disassociating（解除关联中）、Error（异常）。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type CreateHealthCheckTemplateRequestParams struct {
	// 是否预览此次请求。
	// - **false**（默认）：发送普通请求，直接修改健康检查模板。
	// - **true**：发送预览请求，检查修改健康检查模板的参数、格式、业务限制等是否符合要求。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 健康检查状态码。取值：
	// - 当健康检查协议为**HTTP/HTTPS**时：
	// 	- **http_1xx**
	// 	- **http_2xx**（默认值）
	// 	-  **http_3xx**
	// 	-  **http_4xx**
	// 	-  **http_5xx**
	// - 当健康检查协议为**GRPC/GRPCS**时：默认值为**12**，数值范围为**0-99**，输入值可为数值、多个数值或者范围以及相互组合，如：
	// 	- **"20"**
	// 	- **"0-99"**
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitnil,omitempty" name:"HealthCheckCodes"`

	// 判定后端服务健康的阈值，当健康检查连续成功多少次后，后端服务的状态由**不健康**变为**健康**。
	// 取值范围：**2**-**10**。
	// 默认值：**2**。
	HealthCheckHealthyThreshold *uint64 `json:"HealthCheckHealthyThreshold,omitnil,omitempty" name:"HealthCheckHealthyThreshold"`

	// 健康检查域名。
	// 长度限制为 **1-255** 个字符。
	// 可包含小写字母、数字、短划线（-）和半角句号（.）。
	// 
	// > 仅当 **HealthCheckProtocol** 设置为 **HTTP/HTTPS/GRPC/GRPCS** 时，该参数生效。
	HealthCheckHost *string `json:"HealthCheckHost,omitnil,omitempty" name:"HealthCheckHost"`

	// 健康检查 HTTP 协议版本，取值：
	// - **HTTP1.1**（默认）
	// - **HTTP1.0** 
	// > 仅当**HealthCheckProtocol**设置为**HTTP** 或 **HTTPS** 时，该参数生效。
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitnil,omitempty" name:"HealthCheckHttpVersion"`

	// 健康检查的时间间隔。单位：秒。 取值范围：**2**-**300**。 默认值：**5**。
	HealthCheckInterval *uint64 `json:"HealthCheckInterval,omitnil,omitempty" name:"HealthCheckInterval"`

	// 健康检查方法，取值： - **GET** - **HEAD**（默认值） 
	// > 仅当**HealthCheckProtocol**设置为**HTTP** 或 **HTTPS** 时，该参数生效。
	HealthCheckMethod *string `json:"HealthCheckMethod,omitnil,omitempty" name:"HealthCheckMethod"`

	// 健康检查的转发规则路径。 长度为 **1-80** 个字符，只能使用字母、数字、字符`-/.%?#&=`以及扩展字符`_;~!（)*[]@$^:',+`。 URL 必须以正斜线（/）开头。 
	// > 仅当**HealthCheckProtocol**为**HTTP/HTTPS/GRPC/GRPCS**时，转发规则路径参数生效。
	HealthCheckPath *string `json:"HealthCheckPath,omitnil,omitempty" name:"HealthCheckPath"`

	// 健康检查访问后端服务器的端口。  取值范围：**0-65535**。  默认值：**0**，表示后端服务器的端口。
	HealthCheckPort *uint64 `json:"HealthCheckPort,omitnil,omitempty" name:"HealthCheckPort"`

	// 健康检查协议。取值：
	// - **HTTP**（默认）：通过发送 HEAD 或 GET 请求模拟浏览器的访问行为来检查服务器应用是否健康。
	// - **HTTPS**：通过发送 HEAD 或 GET 请求模拟浏览器的访问行为来检查服务器应用是否健康。（数据加密，相比 HTTP 更安全。）
	// - **TCP**：通过发送 SYN 握手报文来检测服务器端口是否存活。
	// - **GRPC**：通过发送 POST 或 GET 请求来检查服务器应用是否健康。
	// - **GRPCS**：通过发送 POST 或 GET 请求来检查服务器应用是否健康。
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitnil,omitempty" name:"HealthCheckProtocol"`

	// 健康检查模板名称。长度为 **1-255** 个字符，可包含数字、大小写字母、中文、半角句号（.）、下划线（_）和短划线（-）。
	HealthCheckTemplateName *string `json:"HealthCheckTemplateName,omitnil,omitempty" name:"HealthCheckTemplateName"`

	// 健康检查的响应超时时间。单位：秒。
	// 取值范围：**2**-**60**。
	// 默认值：**2**。
	HealthCheckTimeout *uint64 `json:"HealthCheckTimeout,omitnil,omitempty" name:"HealthCheckTimeout"`

	// 判定后端服务不健康的阈值，当健康检查连续失败多少次后，后端服务的状态由**健康**变为**不健康**。
	// 取值范围：**2**-**10**。
	// 默认值：**2**。
	HealthCheckUnhealthyThreshold *uint64 `json:"HealthCheckUnhealthyThreshold,omitnil,omitempty" name:"HealthCheckUnhealthyThreshold"`

	// 标签。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateHealthCheckTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 是否预览此次请求。
	// - **false**（默认）：发送普通请求，直接修改健康检查模板。
	// - **true**：发送预览请求，检查修改健康检查模板的参数、格式、业务限制等是否符合要求。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 健康检查状态码。取值：
	// - 当健康检查协议为**HTTP/HTTPS**时：
	// 	- **http_1xx**
	// 	- **http_2xx**（默认值）
	// 	-  **http_3xx**
	// 	-  **http_4xx**
	// 	-  **http_5xx**
	// - 当健康检查协议为**GRPC/GRPCS**时：默认值为**12**，数值范围为**0-99**，输入值可为数值、多个数值或者范围以及相互组合，如：
	// 	- **"20"**
	// 	- **"0-99"**
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitnil,omitempty" name:"HealthCheckCodes"`

	// 判定后端服务健康的阈值，当健康检查连续成功多少次后，后端服务的状态由**不健康**变为**健康**。
	// 取值范围：**2**-**10**。
	// 默认值：**2**。
	HealthCheckHealthyThreshold *uint64 `json:"HealthCheckHealthyThreshold,omitnil,omitempty" name:"HealthCheckHealthyThreshold"`

	// 健康检查域名。
	// 长度限制为 **1-255** 个字符。
	// 可包含小写字母、数字、短划线（-）和半角句号（.）。
	// 
	// > 仅当 **HealthCheckProtocol** 设置为 **HTTP/HTTPS/GRPC/GRPCS** 时，该参数生效。
	HealthCheckHost *string `json:"HealthCheckHost,omitnil,omitempty" name:"HealthCheckHost"`

	// 健康检查 HTTP 协议版本，取值：
	// - **HTTP1.1**（默认）
	// - **HTTP1.0** 
	// > 仅当**HealthCheckProtocol**设置为**HTTP** 或 **HTTPS** 时，该参数生效。
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitnil,omitempty" name:"HealthCheckHttpVersion"`

	// 健康检查的时间间隔。单位：秒。 取值范围：**2**-**300**。 默认值：**5**。
	HealthCheckInterval *uint64 `json:"HealthCheckInterval,omitnil,omitempty" name:"HealthCheckInterval"`

	// 健康检查方法，取值： - **GET** - **HEAD**（默认值） 
	// > 仅当**HealthCheckProtocol**设置为**HTTP** 或 **HTTPS** 时，该参数生效。
	HealthCheckMethod *string `json:"HealthCheckMethod,omitnil,omitempty" name:"HealthCheckMethod"`

	// 健康检查的转发规则路径。 长度为 **1-80** 个字符，只能使用字母、数字、字符`-/.%?#&=`以及扩展字符`_;~!（)*[]@$^:',+`。 URL 必须以正斜线（/）开头。 
	// > 仅当**HealthCheckProtocol**为**HTTP/HTTPS/GRPC/GRPCS**时，转发规则路径参数生效。
	HealthCheckPath *string `json:"HealthCheckPath,omitnil,omitempty" name:"HealthCheckPath"`

	// 健康检查访问后端服务器的端口。  取值范围：**0-65535**。  默认值：**0**，表示后端服务器的端口。
	HealthCheckPort *uint64 `json:"HealthCheckPort,omitnil,omitempty" name:"HealthCheckPort"`

	// 健康检查协议。取值：
	// - **HTTP**（默认）：通过发送 HEAD 或 GET 请求模拟浏览器的访问行为来检查服务器应用是否健康。
	// - **HTTPS**：通过发送 HEAD 或 GET 请求模拟浏览器的访问行为来检查服务器应用是否健康。（数据加密，相比 HTTP 更安全。）
	// - **TCP**：通过发送 SYN 握手报文来检测服务器端口是否存活。
	// - **GRPC**：通过发送 POST 或 GET 请求来检查服务器应用是否健康。
	// - **GRPCS**：通过发送 POST 或 GET 请求来检查服务器应用是否健康。
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitnil,omitempty" name:"HealthCheckProtocol"`

	// 健康检查模板名称。长度为 **1-255** 个字符，可包含数字、大小写字母、中文、半角句号（.）、下划线（_）和短划线（-）。
	HealthCheckTemplateName *string `json:"HealthCheckTemplateName,omitnil,omitempty" name:"HealthCheckTemplateName"`

	// 健康检查的响应超时时间。单位：秒。
	// 取值范围：**2**-**60**。
	// 默认值：**2**。
	HealthCheckTimeout *uint64 `json:"HealthCheckTimeout,omitnil,omitempty" name:"HealthCheckTimeout"`

	// 判定后端服务不健康的阈值，当健康检查连续失败多少次后，后端服务的状态由**健康**变为**不健康**。
	// 取值范围：**2**-**10**。
	// 默认值：**2**。
	HealthCheckUnhealthyThreshold *uint64 `json:"HealthCheckUnhealthyThreshold,omitnil,omitempty" name:"HealthCheckUnhealthyThreshold"`

	// 标签。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateHealthCheckTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHealthCheckTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DryRun")
	delete(f, "HealthCheckCodes")
	delete(f, "HealthCheckHealthyThreshold")
	delete(f, "HealthCheckHost")
	delete(f, "HealthCheckHttpVersion")
	delete(f, "HealthCheckInterval")
	delete(f, "HealthCheckMethod")
	delete(f, "HealthCheckPath")
	delete(f, "HealthCheckPort")
	delete(f, "HealthCheckProtocol")
	delete(f, "HealthCheckTemplateName")
	delete(f, "HealthCheckTimeout")
	delete(f, "HealthCheckUnhealthyThreshold")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHealthCheckTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHealthCheckTemplateResponseParams struct {
	// 健康检查模板 ID，格式为 hct- 后接字母数字。所有接口（创建、查询、修改、删除）均使用 hct- 前缀。
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitnil,omitempty" name:"HealthCheckTemplateId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateHealthCheckTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateHealthCheckTemplateResponseParams `json:"Response"`
}

func (r *CreateHealthCheckTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHealthCheckTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateListenerRequestParams struct {
	// 默认转发规则动作列表。目前监听器仅支持添加 1 个默认转发规则动作。
	DefaultActions []*DefaultAction `json:"DefaultActions,omitnil,omitempty" name:"DefaultActions"`

	// 负载均衡实例前端使用的端口。  取值：1~65535。
	ListenerPort *uint64 `json:"ListenerPort,omitnil,omitempty" name:"ListenerPort"`

	// 监听协议。  取值：HTTP、HTTPS 或 QUIC。
	ListenerProtocol *string `json:"ListenerProtocol,omitnil,omitempty" name:"ListenerProtocol"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 监听器配置的CA证书ID列表。目前监听器仅支持添加 1 个 CA 证书。
	// 当 CaEnabled 参数取值为 true 时，此参数必填。
	CaCertificateIds []*string `json:"CaCertificateIds,omitnil,omitempty" name:"CaCertificateIds"`

	// 是否开启双向认证。
	// 取值：
	// true：开启。
	// false（默认值）：不开启。
	CaEnabled *bool `json:"CaEnabled,omitnil,omitempty" name:"CaEnabled"`

	// 服务器证书 ID 列表。
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 客户端Token，用于保证请求的幂等性。  
	// 
	// 从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否开启Gzip压缩。取值:true(默认值):是。false:否
	GzipEnabled *bool `json:"GzipEnabled,omitnil,omitempty" name:"GzipEnabled"`

	// 是否开启HTTP/2特性。HTTP 协议默认 false，HTTPS 协议默认 true。只有 HTTPS 协议支持此参数。
	Http2Enabled *bool `json:"Http2Enabled,omitnil,omitempty" name:"Http2Enabled"`

	// 连接空闲超时时间。单位：秒。
	// 取值范围：1~600。
	// 默认值：15。
	// 如果在超时时间内一直没有访问请求，负载均衡会断开当前连接，在下次请求到来时创建新的连接。
	IdleTimeout *uint64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`

	// 自定义监听名称。  长度为 1~255 个字符，必须是中文和无害字符串中的字符，  可包含中文、字母、数字、短划线（-）、正斜线（/）、半角句号（.）、下划线（_）。
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// 请求超时时间。单位：秒。
	// 取值：1~600。
	// 默认值：60。
	// 如果在超时时间内后端服务器没有返回响应，负载均衡将放弃等待，并给客户端返回HTTP 504错误码。
	RequestTimeout *uint64 `json:"RequestTimeout,omitnil,omitempty" name:"RequestTimeout"`

	// 安全策略 ID，格式为 tls- 后接 8 位字母数字。
	SecurityPolicyId *string `json:"SecurityPolicyId,omitnil,omitempty" name:"SecurityPolicyId"`

	// 标签列表。最大支持20个。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// X-Forwarded-For配置
	XForwardedForConfig *XForwardedForConfig `json:"XForwardedForConfig,omitnil,omitempty" name:"XForwardedForConfig"`
}

type CreateListenerRequest struct {
	*tchttp.BaseRequest
	
	// 默认转发规则动作列表。目前监听器仅支持添加 1 个默认转发规则动作。
	DefaultActions []*DefaultAction `json:"DefaultActions,omitnil,omitempty" name:"DefaultActions"`

	// 负载均衡实例前端使用的端口。  取值：1~65535。
	ListenerPort *uint64 `json:"ListenerPort,omitnil,omitempty" name:"ListenerPort"`

	// 监听协议。  取值：HTTP、HTTPS 或 QUIC。
	ListenerProtocol *string `json:"ListenerProtocol,omitnil,omitempty" name:"ListenerProtocol"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 监听器配置的CA证书ID列表。目前监听器仅支持添加 1 个 CA 证书。
	// 当 CaEnabled 参数取值为 true 时，此参数必填。
	CaCertificateIds []*string `json:"CaCertificateIds,omitnil,omitempty" name:"CaCertificateIds"`

	// 是否开启双向认证。
	// 取值：
	// true：开启。
	// false（默认值）：不开启。
	CaEnabled *bool `json:"CaEnabled,omitnil,omitempty" name:"CaEnabled"`

	// 服务器证书 ID 列表。
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 客户端Token，用于保证请求的幂等性。  
	// 
	// 从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否开启Gzip压缩。取值:true(默认值):是。false:否
	GzipEnabled *bool `json:"GzipEnabled,omitnil,omitempty" name:"GzipEnabled"`

	// 是否开启HTTP/2特性。HTTP 协议默认 false，HTTPS 协议默认 true。只有 HTTPS 协议支持此参数。
	Http2Enabled *bool `json:"Http2Enabled,omitnil,omitempty" name:"Http2Enabled"`

	// 连接空闲超时时间。单位：秒。
	// 取值范围：1~600。
	// 默认值：15。
	// 如果在超时时间内一直没有访问请求，负载均衡会断开当前连接，在下次请求到来时创建新的连接。
	IdleTimeout *uint64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`

	// 自定义监听名称。  长度为 1~255 个字符，必须是中文和无害字符串中的字符，  可包含中文、字母、数字、短划线（-）、正斜线（/）、半角句号（.）、下划线（_）。
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// 请求超时时间。单位：秒。
	// 取值：1~600。
	// 默认值：60。
	// 如果在超时时间内后端服务器没有返回响应，负载均衡将放弃等待，并给客户端返回HTTP 504错误码。
	RequestTimeout *uint64 `json:"RequestTimeout,omitnil,omitempty" name:"RequestTimeout"`

	// 安全策略 ID，格式为 tls- 后接 8 位字母数字。
	SecurityPolicyId *string `json:"SecurityPolicyId,omitnil,omitempty" name:"SecurityPolicyId"`

	// 标签列表。最大支持20个。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// X-Forwarded-For配置
	XForwardedForConfig *XForwardedForConfig `json:"XForwardedForConfig,omitnil,omitempty" name:"XForwardedForConfig"`
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
	delete(f, "DefaultActions")
	delete(f, "ListenerPort")
	delete(f, "ListenerProtocol")
	delete(f, "LoadBalancerId")
	delete(f, "CaCertificateIds")
	delete(f, "CaEnabled")
	delete(f, "CertificateIds")
	delete(f, "ClientToken")
	delete(f, "GzipEnabled")
	delete(f, "Http2Enabled")
	delete(f, "IdleTimeout")
	delete(f, "ListenerName")
	delete(f, "RequestTimeout")
	delete(f, "SecurityPolicyId")
	delete(f, "Tags")
	delete(f, "XForwardedForConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateListenerResponseParams struct {
	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

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
	// 应用型负载均衡的地址类型。取值：
	// 
	// - **Internet**：负载均衡具有公网IP地址，DNS域名被解析到公网IP，因此可以在公网环境访问。
	// 
	// - **Intranet**：负载均衡只有私网IP地址，DNS域名被解析到私网IP，因此只能被负载均衡所在VPC的内网环境访问。
	AddressType *string `json:"AddressType,omitnil,omitempty" name:"AddressType"`

	// 应用型负载均衡实例计费配置。
	LoadBalancerBillingConfig *LoadBalancerBillingConfig `json:"LoadBalancerBillingConfig,omitnil,omitempty" name:"LoadBalancerBillingConfig"`

	// 私有网络 ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 可用区及私有网络子网映射列表，最多支持添加10个可用区。若当前地域支持2个及以上的可用区，至少需要添加2个可用区。
	ZoneMappings []*ZoneMappingsItem `json:"ZoneMappings,omitnil,omitempty" name:"ZoneMappings"`

	// IP 地址版本，取值 IPv4 或 IPv6。
	AddressIpVersion *string `json:"AddressIpVersion,omitnil,omitempty" name:"AddressIpVersion"`

	// 客户端Token，用于保证请求的幂等性。
	// 
	// 从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 删除保护配置。
	DeleteProtection *DeletionProtectionConfig `json:"DeleteProtection,omitnil,omitempty" name:"DeleteProtection"`

	// 是否只预检此次请求，取值：
	// 
	// - **true**：发送检查请求，不会创建应用型负载均衡实例。检查项包括是否填写了必需参数、请求格式和业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回错误码`DryRunOperation`。
	// 
	// - **false**（默认值）：发送正常请求，通过检查后返回HTTP 2xx状态码并直接进行操作。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// EIP 地址类型，可取值：
	// - **EIP**: 普通弹性公网 IP
	// - **AntiDDoSEIP**: 高防EIP
	// - **AnycastEIP**: 加速 EIP
	// - **HighQualityEIP**: 精品 IP。仅新加坡和中国香港支持精品IP。
	// - **ResidentialEIP**: 原生 IP
	// 
	// 不传默认是EIP。
	InternetAddressType *string `json:"InternetAddressType,omitnil,omitempty" name:"InternetAddressType"`

	// 应用型负载均衡实例名称。长度为1~80个字符，可包含中文、字母、数字、短划线（-）、正斜线（/）、半角句号（.）和下划线（_）。
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 标签。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 应用型负载均衡的地址类型。取值：
	// 
	// - **Internet**：负载均衡具有公网IP地址，DNS域名被解析到公网IP，因此可以在公网环境访问。
	// 
	// - **Intranet**：负载均衡只有私网IP地址，DNS域名被解析到私网IP，因此只能被负载均衡所在VPC的内网环境访问。
	AddressType *string `json:"AddressType,omitnil,omitempty" name:"AddressType"`

	// 应用型负载均衡实例计费配置。
	LoadBalancerBillingConfig *LoadBalancerBillingConfig `json:"LoadBalancerBillingConfig,omitnil,omitempty" name:"LoadBalancerBillingConfig"`

	// 私有网络 ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 可用区及私有网络子网映射列表，最多支持添加10个可用区。若当前地域支持2个及以上的可用区，至少需要添加2个可用区。
	ZoneMappings []*ZoneMappingsItem `json:"ZoneMappings,omitnil,omitempty" name:"ZoneMappings"`

	// IP 地址版本，取值 IPv4 或 IPv6。
	AddressIpVersion *string `json:"AddressIpVersion,omitnil,omitempty" name:"AddressIpVersion"`

	// 客户端Token，用于保证请求的幂等性。
	// 
	// 从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 删除保护配置。
	DeleteProtection *DeletionProtectionConfig `json:"DeleteProtection,omitnil,omitempty" name:"DeleteProtection"`

	// 是否只预检此次请求，取值：
	// 
	// - **true**：发送检查请求，不会创建应用型负载均衡实例。检查项包括是否填写了必需参数、请求格式和业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回错误码`DryRunOperation`。
	// 
	// - **false**（默认值）：发送正常请求，通过检查后返回HTTP 2xx状态码并直接进行操作。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// EIP 地址类型，可取值：
	// - **EIP**: 普通弹性公网 IP
	// - **AntiDDoSEIP**: 高防EIP
	// - **AnycastEIP**: 加速 EIP
	// - **HighQualityEIP**: 精品 IP。仅新加坡和中国香港支持精品IP。
	// - **ResidentialEIP**: 原生 IP
	// 
	// 不传默认是EIP。
	InternetAddressType *string `json:"InternetAddressType,omitnil,omitempty" name:"InternetAddressType"`

	// 应用型负载均衡实例名称。长度为1~80个字符，可包含中文、字母、数字、短划线（-）、正斜线（/）、半角句号（.）和下划线（_）。
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 标签。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
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
	delete(f, "AddressType")
	delete(f, "LoadBalancerBillingConfig")
	delete(f, "VpcId")
	delete(f, "ZoneMappings")
	delete(f, "AddressIpVersion")
	delete(f, "ClientToken")
	delete(f, "DeleteProtection")
	delete(f, "DryRun")
	delete(f, "InternetAddressType")
	delete(f, "LoadBalancerName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLoadBalancerResponseParams struct {
	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

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
type CreateRulesRequestParams struct {
	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 转发规则列表。
	Rules []*RuleInput `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 客户端Token，用于保证请求的幂等性。  从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符。  若您未指定，则系统自动使用API请求的RequestId作为ClientToken标识。每次API请求的RequestId不一样。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否只预检查此次请求。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type CreateRulesRequest struct {
	*tchttp.BaseRequest
	
	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 转发规则列表。
	Rules []*RuleInput `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 客户端Token，用于保证请求的幂等性。  从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符。  若您未指定，则系统自动使用API请求的RequestId作为ClientToken标识。每次API请求的RequestId不一样。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否只预检查此次请求。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *CreateRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "LoadBalancerId")
	delete(f, "Rules")
	delete(f, "ClientToken")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRulesResponseParams struct {
	// 转发规则 ID 列表，ID 格式为 rule- 后接 8 位字母数字。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRulesResponse struct {
	*tchttp.BaseResponse
	Response *CreateRulesResponseParams `json:"Response"`
}

func (r *CreateRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityPolicyRequestParams struct {
	// <p>安全策略支持的加密套件列表。加密套件用于协商客户端与服务端之间的加密算法。</p><p><strong>配置说明：</strong></p><ul><li>加密套件的可选范围取决于所选的 TLS 协议版本（TLSVersions 参数）。</li><li>只要加密套件被任意一个已选 TLS 版本支持，即可添加到列表中。</li><li>若 TLSVersions 包含 TLSv1.3：可不指定 TLSv1.3 专属加密套件（系统将自动补全全部 TLSv1.3 套件）；若指定，则必须包含全部 TLSv1.3 专属加密套件，不支持仅指定部分。</li></ul><p><strong>获取可用加密套件：</strong><br>请调用 <a href="https://cloud.tencent.com/document/api/1822/133718">DescribeSecurityPolicyCapabilities</a> 接口查询各 TLS 版本支持的加密套件列表。</p>
	Ciphers []*string `json:"Ciphers,omitnil,omitempty" name:"Ciphers"`

	// <p>安全策略支持的 TLS 协议版本列表。TLS（Transport Layer Security）协议用于保障客户端与负载均衡之间的通信安全。</p><p><strong>可选值：</strong></p><ul><li><strong>TLSv1.0</strong>：兼容性最好，但安全性较低，不推荐在生产环境使用。</li><li><strong>TLSv1.1</strong>：安全性略优于 TLSv1.0，但仍不推荐。</li><li><strong>TLSv1.2</strong>：目前主流的安全协议版本，兼顾安全性与兼容性。</li><li><strong>TLSv1.3</strong>：最新版本，安全性最高，性能更优，推荐优先使用。</li></ul><p><strong>建议：</strong> 生产环境建议至少选择 TLSv1.2，若客户端支持，优先启用 TLSv1.3。</p>
	TLSVersions []*string `json:"TLSVersions,omitnil,omitempty" name:"TLSVersions"`

	// <p>客户端幂等性令牌。</p><p>用于保证请求的幂等性，防止因网络超时或客户端重试导致的重复创建。建议使用 UUID 作为令牌值。相同的 ClientToken 在有效期内重复请求时，服务端将返回相同的结果。</p>
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// <p>是否仅执行预检请求。取值：</p><ul><li><strong>true</strong>：仅执行预检请求，不实际创建资源。预检请求将验证参数格式、权限及资源配额等，帮助您在正式操作前发现潜在问题。</li><li><strong>false</strong>（默认）：执行正常请求，通过预检后将直接创建安全策略。</li></ul>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// <p>安全策略名称。用于标识和区分不同的安全策略。</p><p><strong>命名规则：</strong></p><ul><li>长度为 2~128 个字符。</li><li>必须以英文字母或中文开头。</li><li>可包含英文字母、中文、数字、半角句号（.）、下划线（_）和短划线（-）。</li></ul><p><strong>建议：</strong> 使用具有业务含义的名称，例如 &quot;prod-high-security&quot; 或 &quot;测试环境策略&quot;。</p>
	SecurityPolicyName *string `json:"SecurityPolicyName,omitnil,omitempty" name:"SecurityPolicyName"`

	// <p>安全策略的标签列表。标签用于对资源进行分类和管理，便于按业务、环境、部门等维度筛选和组织资源。</p><p>每个标签由键值对（Key-Value）组成，同一资源下标签键不可重复。</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateSecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// <p>安全策略支持的加密套件列表。加密套件用于协商客户端与服务端之间的加密算法。</p><p><strong>配置说明：</strong></p><ul><li>加密套件的可选范围取决于所选的 TLS 协议版本（TLSVersions 参数）。</li><li>只要加密套件被任意一个已选 TLS 版本支持，即可添加到列表中。</li><li>若 TLSVersions 包含 TLSv1.3：可不指定 TLSv1.3 专属加密套件（系统将自动补全全部 TLSv1.3 套件）；若指定，则必须包含全部 TLSv1.3 专属加密套件，不支持仅指定部分。</li></ul><p><strong>获取可用加密套件：</strong><br>请调用 <a href="https://cloud.tencent.com/document/api/1822/133718">DescribeSecurityPolicyCapabilities</a> 接口查询各 TLS 版本支持的加密套件列表。</p>
	Ciphers []*string `json:"Ciphers,omitnil,omitempty" name:"Ciphers"`

	// <p>安全策略支持的 TLS 协议版本列表。TLS（Transport Layer Security）协议用于保障客户端与负载均衡之间的通信安全。</p><p><strong>可选值：</strong></p><ul><li><strong>TLSv1.0</strong>：兼容性最好，但安全性较低，不推荐在生产环境使用。</li><li><strong>TLSv1.1</strong>：安全性略优于 TLSv1.0，但仍不推荐。</li><li><strong>TLSv1.2</strong>：目前主流的安全协议版本，兼顾安全性与兼容性。</li><li><strong>TLSv1.3</strong>：最新版本，安全性最高，性能更优，推荐优先使用。</li></ul><p><strong>建议：</strong> 生产环境建议至少选择 TLSv1.2，若客户端支持，优先启用 TLSv1.3。</p>
	TLSVersions []*string `json:"TLSVersions,omitnil,omitempty" name:"TLSVersions"`

	// <p>客户端幂等性令牌。</p><p>用于保证请求的幂等性，防止因网络超时或客户端重试导致的重复创建。建议使用 UUID 作为令牌值。相同的 ClientToken 在有效期内重复请求时，服务端将返回相同的结果。</p>
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// <p>是否仅执行预检请求。取值：</p><ul><li><strong>true</strong>：仅执行预检请求，不实际创建资源。预检请求将验证参数格式、权限及资源配额等，帮助您在正式操作前发现潜在问题。</li><li><strong>false</strong>（默认）：执行正常请求，通过预检后将直接创建安全策略。</li></ul>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// <p>安全策略名称。用于标识和区分不同的安全策略。</p><p><strong>命名规则：</strong></p><ul><li>长度为 2~128 个字符。</li><li>必须以英文字母或中文开头。</li><li>可包含英文字母、中文、数字、半角句号（.）、下划线（_）和短划线（-）。</li></ul><p><strong>建议：</strong> 使用具有业务含义的名称，例如 &quot;prod-high-security&quot; 或 &quot;测试环境策略&quot;。</p>
	SecurityPolicyName *string `json:"SecurityPolicyName,omitnil,omitempty" name:"SecurityPolicyName"`

	// <p>安全策略的标签列表。标签用于对资源进行分类和管理，便于按业务、环境、部门等维度筛选和组织资源。</p><p>每个标签由键值对（Key-Value）组成，同一资源下标签键不可重复。</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ciphers")
	delete(f, "TLSVersions")
	delete(f, "ClientToken")
	delete(f, "DryRun")
	delete(f, "SecurityPolicyName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityPolicyResponseParams struct {
	// <p>安全策略 ID，格式为 tls- 后接 8 位字母数字。</p>
	SecurityPolicyId *string `json:"SecurityPolicyId,omitnil,omitempty" name:"SecurityPolicyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityPolicyResponseParams `json:"Response"`
}

func (r *CreateSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTargetGroupRequestParams struct {
	// <p>目标组类型。取值：</p><ul><li><strong>Instance</strong>（默认）：Cvm服务器类型或者Eni网卡类型。</li></ul>
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// <p>私有网络 ID。</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>是否预览此次请求。</p><ul><li><strong>false</strong>（默认）：发送普通请求，直接创建目标组。</li><li><strong>true</strong>：发送预览请求，检查创建目标组的参数、格式、业务限制等是否符合要求。</li></ul>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// <p>健康检查配置。</p>
	HealthCheckConfig *HealthCheckConfig `json:"HealthCheckConfig,omitnil,omitempty" name:"HealthCheckConfig"`

	// <p>是否开启长连接。</p>
	KeepaliveEnabled *bool `json:"KeepaliveEnabled,omitnil,omitempty" name:"KeepaliveEnabled"`

	// <p>后端服务协议类型。取值：</p><ul><li><strong>HTTP</strong>（默认）：支持绑定HTTP、HTTPS的监听器</li><li><strong>HTTPS</strong>：支持绑定HTTPS类型的监听器</li><li><strong>GRPC</strong>：支持绑定HTTPS类型的监听器</li><li><strong>GRPCS</strong>：支持绑定HTTPS类型的监听器</li></ul>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>调度算法。取值：</p><ul><li><strong>wrr</strong>（默认）：加权轮询，按照权重选择后端服务器，权重越高的服务器被轮询到的概率越高。</li><li><strong>wlc</strong>：加权最小连接数，当不同后端服务器权重值相同时，当前连接数越小的后端服务器被轮询到的概率越高。</li></ul>
	SchedulerAlgorithm *string `json:"SchedulerAlgorithm,omitnil,omitempty" name:"SchedulerAlgorithm"`

	// <p>会话保持配置。</p>
	StickySessionConfig *StickySessionConfig `json:"StickySessionConfig,omitnil,omitempty" name:"StickySessionConfig"`

	// <p>标签。</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>目标组名称。默认为目标组ID。长度为 <strong>1-255</strong> 个字符，可包含数字、大小写字母、中文、半角句号（.）、下划线（_）和短划线（-）。</p>
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`
}

type CreateTargetGroupRequest struct {
	*tchttp.BaseRequest
	
	// <p>目标组类型。取值：</p><ul><li><strong>Instance</strong>（默认）：Cvm服务器类型或者Eni网卡类型。</li></ul>
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// <p>私有网络 ID。</p>
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// <p>是否预览此次请求。</p><ul><li><strong>false</strong>（默认）：发送普通请求，直接创建目标组。</li><li><strong>true</strong>：发送预览请求，检查创建目标组的参数、格式、业务限制等是否符合要求。</li></ul>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// <p>健康检查配置。</p>
	HealthCheckConfig *HealthCheckConfig `json:"HealthCheckConfig,omitnil,omitempty" name:"HealthCheckConfig"`

	// <p>是否开启长连接。</p>
	KeepaliveEnabled *bool `json:"KeepaliveEnabled,omitnil,omitempty" name:"KeepaliveEnabled"`

	// <p>后端服务协议类型。取值：</p><ul><li><strong>HTTP</strong>（默认）：支持绑定HTTP、HTTPS的监听器</li><li><strong>HTTPS</strong>：支持绑定HTTPS类型的监听器</li><li><strong>GRPC</strong>：支持绑定HTTPS类型的监听器</li><li><strong>GRPCS</strong>：支持绑定HTTPS类型的监听器</li></ul>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>调度算法。取值：</p><ul><li><strong>wrr</strong>（默认）：加权轮询，按照权重选择后端服务器，权重越高的服务器被轮询到的概率越高。</li><li><strong>wlc</strong>：加权最小连接数，当不同后端服务器权重值相同时，当前连接数越小的后端服务器被轮询到的概率越高。</li></ul>
	SchedulerAlgorithm *string `json:"SchedulerAlgorithm,omitnil,omitempty" name:"SchedulerAlgorithm"`

	// <p>会话保持配置。</p>
	StickySessionConfig *StickySessionConfig `json:"StickySessionConfig,omitnil,omitempty" name:"StickySessionConfig"`

	// <p>标签。</p>
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>目标组名称。默认为目标组ID。长度为 <strong>1-255</strong> 个字符，可包含数字、大小写字母、中文、半角句号（.）、下划线（_）和短划线（-）。</p>
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`
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
	delete(f, "TargetType")
	delete(f, "VpcId")
	delete(f, "DryRun")
	delete(f, "HealthCheckConfig")
	delete(f, "KeepaliveEnabled")
	delete(f, "Protocol")
	delete(f, "SchedulerAlgorithm")
	delete(f, "StickySessionConfig")
	delete(f, "Tags")
	delete(f, "TargetGroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTargetGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTargetGroupResponseParams struct {
	// <p>目标组 ID，格式为 lbtg- 后接 8 位字母数字。</p>
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

type DefaultAction struct {
	// 转发目标组配置。创建监听器时转发动作中的目标组配置仅支持单个目标组。
	TargetGroupConfig *TargetGroupConfig `json:"TargetGroupConfig,omitnil,omitempty" name:"TargetGroupConfig"`

	// 转发动作类型。创建监听器时，默认转发动作类型仅支持转发至目标组。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type DeleteHealthCheckTemplatesRequestParams struct {
	// 健康检查模板 ID 列表，ID 格式为 hct- 后接字母数字。
	HealthCheckTemplateIds []*string `json:"HealthCheckTemplateIds,omitnil,omitempty" name:"HealthCheckTemplateIds"`

	// 是否预览此次请求。
	// - **false**（默认）：发送普通请求，直接删除模板。
	// - **true**：发送预览请求，检查删除模板的参数、格式、业务限制等是否符合要求。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type DeleteHealthCheckTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 健康检查模板 ID 列表，ID 格式为 hct- 后接字母数字。
	HealthCheckTemplateIds []*string `json:"HealthCheckTemplateIds,omitnil,omitempty" name:"HealthCheckTemplateIds"`

	// 是否预览此次请求。
	// - **false**（默认）：发送普通请求，直接删除模板。
	// - **true**：发送预览请求，检查删除模板的参数、格式、业务限制等是否符合要求。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *DeleteHealthCheckTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHealthCheckTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HealthCheckTemplateIds")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteHealthCheckTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteHealthCheckTemplatesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteHealthCheckTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteHealthCheckTemplatesResponseParams `json:"Response"`
}

func (r *DeleteHealthCheckTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteHealthCheckTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteListenerRequestParams struct {
	// 监听器 ID 列表，ID 格式为 lst- 后接 8 位字母数字。
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 客户端Token，用于保证请求的幂等性。
	// 
	// 从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`
}

type DeleteListenerRequest struct {
	*tchttp.BaseRequest
	
	// 监听器 ID 列表，ID 格式为 lst- 后接 8 位字母数字。
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 客户端Token，用于保证请求的幂等性。
	// 
	// 从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`
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
	delete(f, "ListenerIds")
	delete(f, "LoadBalancerId")
	delete(f, "ClientToken")
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
type DeleteLoadBalancersRequestParams struct {
	// 负载均衡实例 ID 列表，ID 格式为 alb- 后接 8 位字母数字。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 客户端Token，用于保证请求的幂等性。
	// 
	// 从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符。
	// 
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否只预检此次请求，取值：
	// 
	// - **true**：发送检查请求，不会删除应用型负载均衡实例。检查项包括是否填写了必需参数、请求格式和业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回错误码`DryRunOperation`。
	// 
	// - **false**（默认值）：发送正常请求，通过检查后返回`HTTP 2xx`状态码并直接进行操作。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type DeleteLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID 列表，ID 格式为 alb- 后接 8 位字母数字。
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 客户端Token，用于保证请求的幂等性。
	// 
	// 从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符。
	// 
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否只预检此次请求，取值：
	// 
	// - **true**：发送检查请求，不会删除应用型负载均衡实例。检查项包括是否填写了必需参数、请求格式和业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回错误码`DryRunOperation`。
	// 
	// - **false**（默认值）：发送正常请求，通过检查后返回`HTTP 2xx`状态码并直接进行操作。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *DeleteLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerIds")
	delete(f, "ClientToken")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLoadBalancersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLoadBalancersResponseParams `json:"Response"`
}

func (r *DeleteLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRulesRequestParams struct {
	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 转发规则 ID 列表，ID 格式为 rule- 后接 8 位字母数字。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// 是否只预检查此次请求。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type DeleteRulesRequest struct {
	*tchttp.BaseRequest
	
	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 转发规则 ID 列表，ID 格式为 rule- 后接 8 位字母数字。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`

	// 是否只预检查此次请求。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *DeleteRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "LoadBalancerId")
	delete(f, "RuleIds")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRulesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRulesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRulesResponseParams `json:"Response"`
}

func (r *DeleteRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityPolicyRequestParams struct {
	// 安全策略 ID 列表，ID 格式为 tls- 后接 8 位字母数字。
	SecurityPolicyIds []*string `json:"SecurityPolicyIds,omitnil,omitempty" name:"SecurityPolicyIds"`

	// 是否仅执行预检请求。取值：
	// - **true**：仅执行预检请求，不实际删除资源。预检请求将验证参数格式、权限及安全策略是否被引用等，帮助您在正式操作前发现潜在问题。
	// - **false**（默认）：执行正常请求，通过预检后将直接删除安全策略。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type DeleteSecurityPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 安全策略 ID 列表，ID 格式为 tls- 后接 8 位字母数字。
	SecurityPolicyIds []*string `json:"SecurityPolicyIds,omitnil,omitempty" name:"SecurityPolicyIds"`

	// 是否仅执行预检请求。取值：
	// - **true**：仅执行预检请求，不实际删除资源。预检请求将验证参数格式、权限及安全策略是否被引用等，帮助您在正式操作前发现潜在问题。
	// - **false**（默认）：执行正常请求，通过预检后将直接删除安全策略。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *DeleteSecurityPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityPolicyIds")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSecurityPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityPolicyResponseParams `json:"Response"`
}

func (r *DeleteSecurityPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTargetGroupsRequestParams struct {
	// 是否预览此次请求。
	// - **false**（默认）：发送普通请求，直接删除目标组。
	// - **true**：发送预览请求，检查删除目标组的参数、格式、业务限制等是否符合要求。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 目标组 ID 列表，ID 格式为 lbtg- 后接 8 位字母数字。
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`
}

type DeleteTargetGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 是否预览此次请求。
	// - **false**（默认）：发送普通请求，直接删除目标组。
	// - **true**：发送预览请求，检查删除目标组的参数、格式、业务限制等是否符合要求。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 目标组 ID 列表，ID 格式为 lbtg- 后接 8 位字母数字。
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
	delete(f, "DryRun")
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

type DeletionProtectionConfig struct {
	// 是否开启删除保护。开启后，可防止实例被意外删除。
	// - true：开启删除保护
	// - false：关闭删除保护
	DeletionProtectionEnabled *bool `json:"DeletionProtectionEnabled,omitnil,omitempty" name:"DeletionProtectionEnabled"`

	// 开启修改保护的原因说明。
	// 长度为 1~255 个字符，必须是中文和无害字符串中的字符， 可包含中文、字母、数字、短划线（-）、正斜线（/）、半角句号（.）、下划线（_）。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

// Predefined struct for user
type DescribeAsyncJobsRequestParams struct {
	// 分批次查询时每次显示的条目数。取值范围：1~100，默认值：20。
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 是否拥有下一次查询的令牌（Token）。取值：  第一次查询和没有下一次查询时，均无需填写。 如果有下一次查询，取值为上一次 API 调用返回的 NextToken 值。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 异步请求返回的RequestId列表
	RequestIds []*string `json:"RequestIds,omitnil,omitempty" name:"RequestIds"`
}

type DescribeAsyncJobsRequest struct {
	*tchttp.BaseRequest
	
	// 分批次查询时每次显示的条目数。取值范围：1~100，默认值：20。
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 是否拥有下一次查询的令牌（Token）。取值：  第一次查询和没有下一次查询时，均无需填写。 如果有下一次查询，取值为上一次 API 调用返回的 NextToken 值。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 异步请求返回的RequestId列表
	RequestIds []*string `json:"RequestIds,omitnil,omitempty" name:"RequestIds"`
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
	delete(f, "MaxResults")
	delete(f, "NextToken")
	delete(f, "RequestIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAsyncJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAsyncJobsResponseParams struct {
	// 任务列表。
	Jobs []*Job `json:"Jobs,omitnil,omitempty" name:"Jobs"`

	// 分批次查询时每次显示的条目数。
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 是否拥有下一次查询的令牌（Token）。取值：  如果 NextToken 为空表示没有下一次查询。 如果 NextToken 有返回值，该取值表示下一次查询开始的令牌。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 列表条目数。
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
type DescribeHealthCheckTemplatesRequestParams struct {
	// <p>过滤器。通过指定的过滤条件来查询健康检查模板，支持：</p><ul><li>Name的值为<strong>HealthCheckTemplateName</strong>。通过名称来筛选健康检查模板。<strong>Values</strong>的值为模板名称列表。</li><li>Name的值为<strong>HealthCheckProtocol</strong>。通过健康检查协议来筛选健康检查模板。<strong>Values</strong>的值为协议列表。</li><li>通过标签方式筛选。</li></ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>健康检查模板 ID 列表，ID 格式为 hct- 后接字母数字。</p>
	HealthCheckTemplateIds []*string `json:"HealthCheckTemplateIds,omitnil,omitempty" name:"HealthCheckTemplateIds"`

	// <p>返回列表的数量，默认为20，最大值为100。</p>
	MaxResults *string `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// <p>下一次查询的Token值。第一次查询和没有下一次查询时，无需填写。<br>如果有下一次查询，取值为上一次 API 调用返回的 NextToken 值。</p>
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`
}

type DescribeHealthCheckTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// <p>过滤器。通过指定的过滤条件来查询健康检查模板，支持：</p><ul><li>Name的值为<strong>HealthCheckTemplateName</strong>。通过名称来筛选健康检查模板。<strong>Values</strong>的值为模板名称列表。</li><li>Name的值为<strong>HealthCheckProtocol</strong>。通过健康检查协议来筛选健康检查模板。<strong>Values</strong>的值为协议列表。</li><li>通过标签方式筛选。</li></ul>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>健康检查模板 ID 列表，ID 格式为 hct- 后接字母数字。</p>
	HealthCheckTemplateIds []*string `json:"HealthCheckTemplateIds,omitnil,omitempty" name:"HealthCheckTemplateIds"`

	// <p>返回列表的数量，默认为20，最大值为100。</p>
	MaxResults *string `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// <p>下一次查询的Token值。第一次查询和没有下一次查询时，无需填写。<br>如果有下一次查询，取值为上一次 API 调用返回的 NextToken 值。</p>
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`
}

func (r *DescribeHealthCheckTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHealthCheckTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "HealthCheckTemplateIds")
	delete(f, "MaxResults")
	delete(f, "NextToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHealthCheckTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHealthCheckTemplatesResponseParams struct {
	// <p>健康检查模板列表。</p>
	HealthCheckTemplates []*HealthCheckTemplate `json:"HealthCheckTemplates,omitnil,omitempty" name:"HealthCheckTemplates"`

	// <p>下一次查询的Token值，如果当前是最后一页，返回为空。</p>
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// <p>经过筛选后查询到的健康检查模板总数。</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHealthCheckTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHealthCheckTemplatesResponseParams `json:"Response"`
}

func (r *DescribeHealthCheckTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHealthCheckTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenerCertificatesRequestParams struct {
	// 证书类型。取值：CA或SVR（服务器证书）。
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 本次读取的最大数据记录数量。取值: 1~100。默认值: 20。
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 下一次查询的令牌（Token）。取值：
	// 第一次查询和没有下一次查询时，均无需填写。
	// 如果有下一次查询，取值为上一次 API 调用返回的 NextToken 值。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`
}

type DescribeListenerCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// 证书类型。取值：CA或SVR（服务器证书）。
	CertificateType *string `json:"CertificateType,omitnil,omitempty" name:"CertificateType"`

	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 本次读取的最大数据记录数量。取值: 1~100。默认值: 20。
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 下一次查询的令牌（Token）。取值：
	// 第一次查询和没有下一次查询时，均无需填写。
	// 如果有下一次查询，取值为上一次 API 调用返回的 NextToken 值。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`
}

func (r *DescribeListenerCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerCertificatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateType")
	delete(f, "ListenerId")
	delete(f, "LoadBalancerId")
	delete(f, "MaxResults")
	delete(f, "NextToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListenerCertificatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenerCertificatesResponseParams struct {
	// 监听器绑定的证书信息列表。
	Certificates []*CertificateInfo `json:"Certificates,omitnil,omitempty" name:"Certificates"`

	// 本次读取的最大数据记录数量。	
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 下一次查询的令牌。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 监听器绑定的证书总量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeListenerCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListenerCertificatesResponseParams `json:"Response"`
}

func (r *DescribeListenerCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenerDetailRequestParams struct {
	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`
}

type DescribeListenerDetailRequest struct {
	*tchttp.BaseRequest
	
	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`
}

func (r *DescribeListenerDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "LoadBalancerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListenerDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenerDetailResponseParams struct {
	// 监听器绑定的CA证书ID列表。
	CaCertificateIds []*string `json:"CaCertificateIds,omitnil,omitempty" name:"CaCertificateIds"`

	// 是否开启双向认证。
	CaEnabled *bool `json:"CaEnabled,omitnil,omitempty" name:"CaEnabled"`

	// 服务器证书 ID 列表。
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 监听器实例的创建时间。格式：ISO 8601（例如 2025-01-01T08:30:00+08:00）
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 规则动作列表。
	DefaultActions []*DefaultAction `json:"DefaultActions,omitnil,omitempty" name:"DefaultActions"`

	// 是否启用 Gzip 压缩。
	GzipEnabled *bool `json:"GzipEnabled,omitnil,omitempty" name:"GzipEnabled"`

	// 是否开启HTTP/2特性。
	Http2Enabled *bool `json:"Http2Enabled,omitnil,omitempty" name:"Http2Enabled"`

	// 指定连接空闲超时时间。单位：秒。
	IdleTimeout *uint64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`

	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 自定义监听名称。
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// 负载均衡实例前端使用的端口。
	ListenerPort *uint64 `json:"ListenerPort,omitnil,omitempty" name:"ListenerPort"`

	// 监听协议。
	ListenerProtocol *string `json:"ListenerProtocol,omitnil,omitempty" name:"ListenerProtocol"`

	// 监听器状态。取值:=
	// 
	// - **Active**: 运行中。
	// - **Provisioning**：创建中。
	// - **Configuring**：变配中。
	// - **ProvisionFailed**：创建失败
	ListenerStatus *string `json:"ListenerStatus,omitnil,omitempty" name:"ListenerStatus"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 监听器实例的最后变更时间。格式：ISO 8601（例如 2025-01-01T08:30:00+08:00）
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 请求超时时间。单位：秒。
	RequestTimeout *uint64 `json:"RequestTimeout,omitnil,omitempty" name:"RequestTimeout"`

	// 安全策略 ID，格式为 tls- 后接 8 位字母数字。
	SecurityPolicyId *string `json:"SecurityPolicyId,omitnil,omitempty" name:"SecurityPolicyId"`

	// 标签。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// XForwardedFor配置。
	XForwardedForConfig *XForwardedForConfig `json:"XForwardedForConfig,omitnil,omitempty" name:"XForwardedForConfig"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeListenerDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListenerDetailResponseParams `json:"Response"`
}

func (r *DescribeListenerDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenerHealthStatusRequestParams struct {
	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 是否包含转发规则的健康检查结果。为false时只返回默认转发规则的健康状态，为true时返回全部规则的健康状态(包含默认规则)。
	// 取值：
	// true：是。
	// false（默认值）：否。
	IncludeRule *bool `json:"IncludeRule,omitnil,omitempty" name:"IncludeRule"`

	// 本次读取的最大数据记录数量。
	// 取值: 1~100。
	// 默认值: 20
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 下一页查询的Token值。第一次查询时，无需填写。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`
}

type DescribeListenerHealthStatusRequest struct {
	*tchttp.BaseRequest
	
	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 是否包含转发规则的健康检查结果。为false时只返回默认转发规则的健康状态，为true时返回全部规则的健康状态(包含默认规则)。
	// 取值：
	// true：是。
	// false（默认值）：否。
	IncludeRule *bool `json:"IncludeRule,omitnil,omitempty" name:"IncludeRule"`

	// 本次读取的最大数据记录数量。
	// 取值: 1~100。
	// 默认值: 20
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 下一页查询的Token值。第一次查询时，无需填写。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`
}

func (r *DescribeListenerHealthStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerHealthStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "LoadBalancerId")
	delete(f, "IncludeRule")
	delete(f, "MaxResults")
	delete(f, "NextToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListenerHealthStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenerHealthStatusResponseParams struct {
	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 监听器端口。
	ListenerPort *string `json:"ListenerPort,omitnil,omitempty" name:"ListenerPort"`

	// 监听器协议。
	ListenerProtocol *string `json:"ListenerProtocol,omitnil,omitempty" name:"ListenerProtocol"`

	// 下一次查询的令牌（Token）。为空时表示这是最后一页。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 转发规则健康状态。
	RuleHealthStatusInfos []*RuleHealthStatusInfo `json:"RuleHealthStatusInfos,omitnil,omitempty" name:"RuleHealthStatusInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeListenerHealthStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeListenerHealthStatusResponseParams `json:"Response"`
}

func (r *DescribeListenerHealthStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeListenerHealthStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenersRequestParams struct {
	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 过滤条件列表，最大支持20个。支持以下几个字段
	// - **Protocol**: 协议类型
	// - **Tags**: 标签
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 监听器 ID 列表，ID 格式为 lst- 后接 8 位字母数字。
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// 本次读取的最大数据记录数量。
	// 取值: 1~100。
	// 默认值: 20
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 下一次查询的令牌（Token）。为空时查询第一页。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`
}

type DescribeListenersRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 过滤条件列表，最大支持20个。支持以下几个字段
	// - **Protocol**: 协议类型
	// - **Tags**: 标签
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 监听器 ID 列表，ID 格式为 lst- 后接 8 位字母数字。
	ListenerIds []*string `json:"ListenerIds,omitnil,omitempty" name:"ListenerIds"`

	// 本次读取的最大数据记录数量。
	// 取值: 1~100。
	// 默认值: 20
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 下一次查询的令牌（Token）。为空时查询第一页。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`
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
	delete(f, "Filters")
	delete(f, "ListenerIds")
	delete(f, "MaxResults")
	delete(f, "NextToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeListenersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeListenersResponseParams struct {
	// 监听器信息。
	Listeners []*ListenerOutput `json:"Listeners,omitnil,omitempty" name:"Listeners"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 本次读取的最大数据记录数量。
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 下一次查询的令牌。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 总条目数。
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
type DescribeLoadBalancerDetailRequestParams struct {
	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`
}

type DescribeLoadBalancerDetailRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`
}

func (r *DescribeLoadBalancerDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancerDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancerDetailResponseParams struct {
	// 负载均衡详细信息
	LoadBalancerDetail *LoadBalancerDetail `json:"LoadBalancerDetail,omitnil,omitempty" name:"LoadBalancerDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLoadBalancerDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLoadBalancerDetailResponseParams `json:"Response"`
}

func (r *DescribeLoadBalancerDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLoadBalancerDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancersRequestParams struct {
	// 查询过滤条件，支持以下几个字段
	// - **LoadBalancerId**: 负载均衡实例 ID
	// - **LoadBalancerName**: 负载均衡名称
	// - **LoadBalancerStatus**: 负载均衡状态
	// - **VpcId**: 私有网络 ID
	// - **tag:tag-key**：按标签键值对筛选，tag-key 请替换为实际的标签键。例如 `tag:env` 表示按标签键 `env` 筛选。
	// - **AddressType**: 网络类型
	//     - **Intranet**: 内网
	//     - **Internet**: 公网 
	// - **AddressIpVersion**:
	//     - **IPv4**: IPv4 地址
	//     - **IPv6** IPv6 地址
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分批次查询时每次显示的条目数。取值范围：**1**~**100**，默认值：**20**。
	// 
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 是否拥有下一次查询的令牌（Token）。取值：
	// - 第一次查询和没有下一次查询时，均无需填写。
	// - 如果有下一次查询，取值为上一次API调用返回的**NextToken**值。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`
}

type DescribeLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// 查询过滤条件，支持以下几个字段
	// - **LoadBalancerId**: 负载均衡实例 ID
	// - **LoadBalancerName**: 负载均衡名称
	// - **LoadBalancerStatus**: 负载均衡状态
	// - **VpcId**: 私有网络 ID
	// - **tag:tag-key**：按标签键值对筛选，tag-key 请替换为实际的标签键。例如 `tag:env` 表示按标签键 `env` 筛选。
	// - **AddressType**: 网络类型
	//     - **Intranet**: 内网
	//     - **Internet**: 公网 
	// - **AddressIpVersion**:
	//     - **IPv4**: IPv4 地址
	//     - **IPv6** IPv6 地址
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 分批次查询时每次显示的条目数。取值范围：**1**~**100**，默认值：**20**。
	// 
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 是否拥有下一次查询的令牌（Token）。取值：
	// - 第一次查询和没有下一次查询时，均无需填写。
	// - 如果有下一次查询，取值为上一次API调用返回的**NextToken**值。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`
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
	delete(f, "Filters")
	delete(f, "MaxResults")
	delete(f, "NextToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLoadBalancersResponseParams struct {
	// 应用型负载均衡实例列表。
	LoadBalancers []*LoadBalancer `json:"LoadBalancers,omitnil,omitempty" name:"LoadBalancers"`

	// 分批次查询时每次显示的条目数。
	// 
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 是否拥有下一次查询的令牌（Token）。取值：
	// - 如果**NextToken**为空表示没有下一次查询。
	// - 如果**NextToken**有返回值，该取值表示下一次查询开始的令牌。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 列表条目数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

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
type DescribeQuotaRequestParams struct {
	// 配额类型列表。支持同时传入多个配额类型。查询资源级配额时，可配合 ResourceIds 传入对应资源ID；如需返回已使用量和可用量，可在 DisplayFields 中传入 used、available。
	// 
	// 枚举说明：
	// - alb_quota_loadbalancers_num：每个地域可创建的 ALB 实例数。
	// - alb_quota_targetgroups_num：每个地域可创建的 ALB 目标组数。
	// - alb_quota_loadbalancer_listeners_num：每个 ALB 实例可创建的监听器数，ResourceIds 填写 ALB 实例 ID。
	// - alb_quota_loadbalancer_rules_num：每个 ALB 实例可添加的转发规则数，不计入默认规则，ResourceIds 填写 ALB 实例 ID。
	// - alb_quota_loadbalancer_certificates_num：每个 ALB 实例可添加的扩展证书数，不计入默认证书，ResourceIds 填写 ALB 实例 ID。
	// - alb_quota_loadbalancer_targetgroup_num：每个 ALB 实例可绑定的目标组数，ResourceIds 填写 ALB 实例 ID。
	// - alb_quota_loadbalancer_servers_num：每个 ALB 实例可添加的后端服务器数，ResourceIds 填写 ALB 实例 ID。
	// - alb_quota_server_added_num：单个后端服务器 IP 可被添加到 ALB 后端目标组的次数。
	// - alb_quota_targetgroup_attached_num：每个目标组可被 ALB 转发规则关联的次数，ResourceIds 填写目标组 ID。
	// - alb_quota_targetgroup_targets_num：每个目标组支持的后端服务器数，适用于 IP 和端口类型后端，ResourceIds 填写目标组 ID。
	// - alb_quota_targetgroup_targets_num_scf：每个目标组支持的 SCF 函数后端数，ResourceIds 填写目标组 ID。
	// - alb_quota_max_request_timeout：创建监听器时可配置的连接请求最大超时时间。
	// - alb_quota_max_idle_timeout：创建监听器时可配置的连接空闲最大超时时间。
	// - alb_quota_listener_certificates_num：单个监听器可添加的证书数量，ResourceIds 填写监听器 ID。
	// - alb_quota_rule_targetgroups_num：单条转发规则可绑定的目标组数量。
	// - alb_quota_rule_conditions_num：单条转发规则可添加的匹配条件条目数。
	// - alb_quota_rule_wildcards_num：单条转发规则可添加的包含通配符的匹配条目数。
	// - alb_quota_rule_actions_num：单条转发规则可添加的动作条目数。
	// - alb_quota_cipher_template_listeners_num：单个加密套件模板可关联的监听器数量。
	// - alb_quota_healthcheck_templates_num：每个地域可创建的健康检查模板数。
	// - alb_quota_securitygroup_templates_num：单个 ALB 实例支持绑定的安全组数量。
	// - alb_quota_securitygroup_rules_per_sg_num：单个 ALB 实例中单个安全组支持的规则条目数。
	// - alb_quota_security_policies_num：每个地域可创建的自定义安全策略数。
	QuotaTypes []*string `json:"QuotaTypes,omitnil,omitempty" name:"QuotaTypes"`

	// 显示字段列表，用于控制是否额外返回用量信息。支持 used、available：used 表示返回当前已使用量，available 表示返回当前剩余可用量。QuotaType 和 Limit 总是返回；ResourceId 会在请求传入 ResourceIds 时返回。
	DisplayFields []*string `json:"DisplayFields,omitnil,omitempty" name:"DisplayFields"`

	// 资源ID列表。用于查询具体资源维度的配额和用量；不传时查询账号或地域维度的默认配额配置。资源ID的类型由 QuotaTypes 决定，例如 ALB 实例级配额填写 ALB 实例 ID，监听器级配额填写监听器 ID，目标组级配额填写目标组 ID。
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`
}

type DescribeQuotaRequest struct {
	*tchttp.BaseRequest
	
	// 配额类型列表。支持同时传入多个配额类型。查询资源级配额时，可配合 ResourceIds 传入对应资源ID；如需返回已使用量和可用量，可在 DisplayFields 中传入 used、available。
	// 
	// 枚举说明：
	// - alb_quota_loadbalancers_num：每个地域可创建的 ALB 实例数。
	// - alb_quota_targetgroups_num：每个地域可创建的 ALB 目标组数。
	// - alb_quota_loadbalancer_listeners_num：每个 ALB 实例可创建的监听器数，ResourceIds 填写 ALB 实例 ID。
	// - alb_quota_loadbalancer_rules_num：每个 ALB 实例可添加的转发规则数，不计入默认规则，ResourceIds 填写 ALB 实例 ID。
	// - alb_quota_loadbalancer_certificates_num：每个 ALB 实例可添加的扩展证书数，不计入默认证书，ResourceIds 填写 ALB 实例 ID。
	// - alb_quota_loadbalancer_targetgroup_num：每个 ALB 实例可绑定的目标组数，ResourceIds 填写 ALB 实例 ID。
	// - alb_quota_loadbalancer_servers_num：每个 ALB 实例可添加的后端服务器数，ResourceIds 填写 ALB 实例 ID。
	// - alb_quota_server_added_num：单个后端服务器 IP 可被添加到 ALB 后端目标组的次数。
	// - alb_quota_targetgroup_attached_num：每个目标组可被 ALB 转发规则关联的次数，ResourceIds 填写目标组 ID。
	// - alb_quota_targetgroup_targets_num：每个目标组支持的后端服务器数，适用于 IP 和端口类型后端，ResourceIds 填写目标组 ID。
	// - alb_quota_targetgroup_targets_num_scf：每个目标组支持的 SCF 函数后端数，ResourceIds 填写目标组 ID。
	// - alb_quota_max_request_timeout：创建监听器时可配置的连接请求最大超时时间。
	// - alb_quota_max_idle_timeout：创建监听器时可配置的连接空闲最大超时时间。
	// - alb_quota_listener_certificates_num：单个监听器可添加的证书数量，ResourceIds 填写监听器 ID。
	// - alb_quota_rule_targetgroups_num：单条转发规则可绑定的目标组数量。
	// - alb_quota_rule_conditions_num：单条转发规则可添加的匹配条件条目数。
	// - alb_quota_rule_wildcards_num：单条转发规则可添加的包含通配符的匹配条目数。
	// - alb_quota_rule_actions_num：单条转发规则可添加的动作条目数。
	// - alb_quota_cipher_template_listeners_num：单个加密套件模板可关联的监听器数量。
	// - alb_quota_healthcheck_templates_num：每个地域可创建的健康检查模板数。
	// - alb_quota_securitygroup_templates_num：单个 ALB 实例支持绑定的安全组数量。
	// - alb_quota_securitygroup_rules_per_sg_num：单个 ALB 实例中单个安全组支持的规则条目数。
	// - alb_quota_security_policies_num：每个地域可创建的自定义安全策略数。
	QuotaTypes []*string `json:"QuotaTypes,omitnil,omitempty" name:"QuotaTypes"`

	// 显示字段列表，用于控制是否额外返回用量信息。支持 used、available：used 表示返回当前已使用量，available 表示返回当前剩余可用量。QuotaType 和 Limit 总是返回；ResourceId 会在请求传入 ResourceIds 时返回。
	DisplayFields []*string `json:"DisplayFields,omitnil,omitempty" name:"DisplayFields"`

	// 资源ID列表。用于查询具体资源维度的配额和用量；不传时查询账号或地域维度的默认配额配置。资源ID的类型由 QuotaTypes 决定，例如 ALB 实例级配额填写 ALB 实例 ID，监听器级配额填写监听器 ID，目标组级配额填写目标组 ID。
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`
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
	delete(f, "QuotaTypes")
	delete(f, "DisplayFields")
	delete(f, "ResourceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQuotaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQuotaResponseParams struct {
	// 配额列表。每个元素表示一个配额类型的查询结果；当请求传入 ResourceIds 时，每个元素表示一个配额类型和一个资源ID组合的查询结果。
	Quotas []*QuotaInfo `json:"Quotas,omitnil,omitempty" name:"Quotas"`

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
type DescribeRulesRequestParams struct {
	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 支持的过滤条件如下：
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回列表的数量，默认为20，最大值为100。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 下一次查询的Token值。第一次查询和没有下一次查询时，无需填写。如果有下一次查询，取值为上一次 API 调用返回的 NextToken 值。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 转发规则 ID 列表，ID 格式为 rule- 后接 8 位字母数字。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

type DescribeRulesRequest struct {
	*tchttp.BaseRequest
	
	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 支持的过滤条件如下：
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回列表的数量，默认为20，最大值为100。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 下一次查询的Token值。第一次查询和没有下一次查询时，无需填写。如果有下一次查询，取值为上一次 API 调用返回的 NextToken 值。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 转发规则 ID 列表，ID 格式为 rule- 后接 8 位字母数字。
	RuleIds []*string `json:"RuleIds,omitnil,omitempty" name:"RuleIds"`
}

func (r *DescribeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "LoadBalancerId")
	delete(f, "Filters")
	delete(f, "MaxResults")
	delete(f, "NextToken")
	delete(f, "RuleIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRulesResponseParams struct {
	// 下一次查询的Token值，如果当前是最后一页，返回为空。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 转发规则列表。
	Rules []*RuleOutput `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 总的转发规则个数（根据监听器ID、规则ID等条件过滤后）。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRulesResponseParams `json:"Response"`
}

func (r *DescribeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPoliciesRequestParams struct {
	// 过滤条件列表，用于筛选符合指定条件的安全策略。多个过滤条件之间为"与"关系。
	// 
	// **支持的过滤条件：**
	// - **SecurityPolicyNames**：按安全策略名称筛选，支持模糊匹配。
	// - **tag:tag-key**：按标签键值对筛选，tag-key 请替换为实际的标签键。例如 `tag:env` 表示按标签键 `env` 筛选。
	// 
	// **说明：** 每个过滤条件最多支持 10 个取值。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 单次请求返回的最大结果数。用于分页查询，与 NextToken 配合使用。
	// 
	// **取值范围：** 1~100。
	// 
	// **默认值：** 20。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 分页查询的起始令牌。用于获取下一页结果数据。
	// 
	// **使用说明：**
	// - 首次查询时无需设置此参数。
	// - 如果上一次查询返回了 NextToken，表示还有更多数据，请将该值传入此参数以获取下一页。
	// - 若上一次查询未返回 NextToken 或返回为空，表示已是最后一页。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 安全策略 ID 列表，ID 格式为 tls- 后接 8 位字母数字。
	SecurityPolicyIds []*string `json:"SecurityPolicyIds,omitnil,omitempty" name:"SecurityPolicyIds"`
}

type DescribeSecurityPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件列表，用于筛选符合指定条件的安全策略。多个过滤条件之间为"与"关系。
	// 
	// **支持的过滤条件：**
	// - **SecurityPolicyNames**：按安全策略名称筛选，支持模糊匹配。
	// - **tag:tag-key**：按标签键值对筛选，tag-key 请替换为实际的标签键。例如 `tag:env` 表示按标签键 `env` 筛选。
	// 
	// **说明：** 每个过滤条件最多支持 10 个取值。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 单次请求返回的最大结果数。用于分页查询，与 NextToken 配合使用。
	// 
	// **取值范围：** 1~100。
	// 
	// **默认值：** 20。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 分页查询的起始令牌。用于获取下一页结果数据。
	// 
	// **使用说明：**
	// - 首次查询时无需设置此参数。
	// - 如果上一次查询返回了 NextToken，表示还有更多数据，请将该值传入此参数以获取下一页。
	// - 若上一次查询未返回 NextToken 或返回为空，表示已是最后一页。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 安全策略 ID 列表，ID 格式为 tls- 后接 8 位字母数字。
	SecurityPolicyIds []*string `json:"SecurityPolicyIds,omitnil,omitempty" name:"SecurityPolicyIds"`
}

func (r *DescribeSecurityPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "MaxResults")
	delete(f, "NextToken")
	delete(f, "SecurityPolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPoliciesResponseParams struct {
	// 下一页查询的起始令牌。
	// 
	// - 若返回值不为空，表示还有更多数据，可将此值作为下一次请求的 NextToken 参数继续查询。
	// - 若返回值为空或未返回此字段，表示已是最后一页。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 安全策略信息列表。包含每个安全策略的详细配置，如策略 ID、名称、TLS 版本、加密套件等。
	SecurityPolicies []*SecurityPolicyInfo `json:"SecurityPolicies,omitnil,omitempty" name:"SecurityPolicies"`

	// 符合过滤条件的安全策略总数。
	// 
	// **说明：** 此值表示满足查询条件的总记录数，而非本次返回的记录数。可用于计算分页信息。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityPoliciesResponseParams `json:"Response"`
}

func (r *DescribeSecurityPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyCapabilitiesRequestParams struct {

}

type DescribeSecurityPolicyCapabilitiesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSecurityPolicyCapabilitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyCapabilitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPolicyCapabilitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyCapabilitiesResponseParams struct {
	// 安全策略配置能力列表。返回当前地域支持的所有 TLS 版本及其对应的加密套件信息。
	// 
	// **返回内容包含：**
	// - 支持的 TLS 协议版本（如 TLSv1.0、TLSv1.1、TLSv1.2、TLSv1.3）。
	// - 每个 TLS 版本支持的加密套件列表。
	// 
	// **使用场景：**
	// - 在创建安全策略（CreateSecurityPolicy）前，调用此接口获取可选的加密套件。
	// - 在修改安全策略（ModifySecurityPolicyAttributes）前，确认新配置的有效性。
	SecurityPolicyCapabilities []*SecurityPolicyCapability `json:"SecurityPolicyCapabilities,omitnil,omitempty" name:"SecurityPolicyCapabilities"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityPolicyCapabilitiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityPolicyCapabilitiesResponseParams `json:"Response"`
}

func (r *DescribeSecurityPolicyCapabilitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyCapabilitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyRelationsRequestParams struct {
	// 安全策略 ID 列表，ID 格式为 tls- 后接 8 位字母数字。
	SecurityPolicyIds []*string `json:"SecurityPolicyIds,omitnil,omitempty" name:"SecurityPolicyIds"`
}

type DescribeSecurityPolicyRelationsRequest struct {
	*tchttp.BaseRequest
	
	// 安全策略 ID 列表，ID 格式为 tls- 后接 8 位字母数字。
	SecurityPolicyIds []*string `json:"SecurityPolicyIds,omitnil,omitempty" name:"SecurityPolicyIds"`
}

func (r *DescribeSecurityPolicyRelationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyRelationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityPolicyIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityPolicyRelationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityPolicyRelationsResponseParams struct {
	// 安全策略关联的监听器列表。返回每个安全策略所关联的 HTTPS 监听器信息。
	SecurityPolicyRelations []*SecurityPolicyRelations `json:"SecurityPolicyRelations,omitnil,omitempty" name:"SecurityPolicyRelations"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityPolicyRelationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityPolicyRelationsResponseParams `json:"Response"`
}

func (r *DescribeSecurityPolicyRelationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityPolicyRelationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSystemSecurityPoliciesRequestParams struct {

}

type DescribeSystemSecurityPoliciesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSystemSecurityPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSystemSecurityPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSystemSecurityPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSystemSecurityPoliciesResponseParams struct {
	// 系统安全策略列表。
	SecurityPolicies []*SecurityPolicyInfo `json:"SecurityPolicies,omitnil,omitempty" name:"SecurityPolicies"`

	// 安全策略总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSystemSecurityPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSystemSecurityPoliciesResponseParams `json:"Response"`
}

func (r *DescribeSystemSecurityPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSystemSecurityPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupTargetsRequestParams struct {
	// 目标组 ID，格式为 lbtg- 后接 8 位字母数字。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 过滤器。通过指定的过滤条件来查询后端服务，支持：
	// - Name的值为**TargetId**。通过资源ID来筛选后端服务，当目标组后端类型为**Instance**时生效。**Values**的值为Cvm或Eni的资源ID。
	// - Name的值为**TargetIp**。通过资源IP来筛选后端服务，当目标组后端类型为**Ip**时生效。**Values**的值为后端服务的IP。
	// - 通过标签方式筛选。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回列表的数量，默认为**20**，最大值为**100**。
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 下一次查询的Token值。第一次查询和没有下一次查询时，无需填写。
	// 如果有下一次查询，取值为上一次 API 调用返回的 NextToken 值。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`
}

type DescribeTargetGroupTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 目标组 ID，格式为 lbtg- 后接 8 位字母数字。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 过滤器。通过指定的过滤条件来查询后端服务，支持：
	// - Name的值为**TargetId**。通过资源ID来筛选后端服务，当目标组后端类型为**Instance**时生效。**Values**的值为Cvm或Eni的资源ID。
	// - Name的值为**TargetIp**。通过资源IP来筛选后端服务，当目标组后端类型为**Ip**时生效。**Values**的值为后端服务的IP。
	// - 通过标签方式筛选。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回列表的数量，默认为**20**，最大值为**100**。
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 下一次查询的Token值。第一次查询和没有下一次查询时，无需填写。
	// 如果有下一次查询，取值为上一次 API 调用返回的 NextToken 值。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`
}

func (r *DescribeTargetGroupTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "Filters")
	delete(f, "MaxResults")
	delete(f, "NextToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetGroupTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupTargetsResponseParams struct {
	// 下一次查询的Token值，如果当前是最后一页，返回为空。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 后端服务信息。
	Targets []*TargetOutput `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 目标组内后端服务总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTargetGroupTargetsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetGroupTargetsResponseParams `json:"Response"`
}

func (r *DescribeTargetGroupTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupsByTargetRequestParams struct {
	// 后端服务实例 ID，CVM 实例格式为 ins- 后接 8 位字母数字。
	TargetId []*string `json:"TargetId,omitnil,omitempty" name:"TargetId"`
}

type DescribeTargetGroupsByTargetRequest struct {
	*tchttp.BaseRequest
	
	// 后端服务实例 ID，CVM 实例格式为 ins- 后接 8 位字母数字。
	TargetId []*string `json:"TargetId,omitnil,omitempty" name:"TargetId"`
}

func (r *DescribeTargetGroupsByTargetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupsByTargetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetGroupsByTargetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupsByTargetResponseParams struct {
	// 总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTargetGroupsByTargetResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTargetGroupsByTargetResponseParams `json:"Response"`
}

func (r *DescribeTargetGroupsByTargetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTargetGroupsByTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupsRequestParams struct {
	//  过滤器。通过指定的过滤条件来查询后端服务，支持：
	// - Name的值为**VpcId**。通过VPC实例来筛选目标组。**Values**的值为VPC唯一ID列表。
	// - Name的值为**TargetType**。通过后端服务类型来筛选目标组。**Values**的值可以取为**Instance**。
	// - Name的值为**TargetGroupName**。通过目标组名称来筛选目标组。**Values**的值为目标组名称列表。
	// - Name的值为**Protocol**。通过目标组后端服务协议来筛选目标组。**Values**的值为目标组后端服务协议列表。
	// - 通过标签方式筛选。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回列表的数量，默认为**20**，最大值为**100**。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 下一次查询的Token值。第一次查询和没有下一次查询时，无需填写。
	// 如果有下一次查询，取值为上一次 API 调用返回的 NextToken 值。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 目标组 ID 列表，ID 格式为 lbtg- 后接 8 位字母数字。
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`
}

type DescribeTargetGroupsRequest struct {
	*tchttp.BaseRequest
	
	//  过滤器。通过指定的过滤条件来查询后端服务，支持：
	// - Name的值为**VpcId**。通过VPC实例来筛选目标组。**Values**的值为VPC唯一ID列表。
	// - Name的值为**TargetType**。通过后端服务类型来筛选目标组。**Values**的值可以取为**Instance**。
	// - Name的值为**TargetGroupName**。通过目标组名称来筛选目标组。**Values**的值为目标组名称列表。
	// - Name的值为**Protocol**。通过目标组后端服务协议来筛选目标组。**Values**的值为目标组后端服务协议列表。
	// - 通过标签方式筛选。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回列表的数量，默认为**20**，最大值为**100**。
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 下一次查询的Token值。第一次查询和没有下一次查询时，无需填写。
	// 如果有下一次查询，取值为上一次 API 调用返回的 NextToken 值。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 目标组 ID 列表，ID 格式为 lbtg- 后接 8 位字母数字。
	TargetGroupIds []*string `json:"TargetGroupIds,omitnil,omitempty" name:"TargetGroupIds"`
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
	delete(f, "Filters")
	delete(f, "MaxResults")
	delete(f, "NextToken")
	delete(f, "TargetGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTargetGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTargetGroupsResponseParams struct {
	// 下一次查询的Token值，如果当前是最后一页，返回为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 目标组信息。
	TargetGroups []*TargetGroupOutput `json:"TargetGroups,omitnil,omitempty" name:"TargetGroups"`

	// 目标组总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

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
type DescribeZonesRequestParams struct {

}

type DescribeZonesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZonesResponseParams struct {
	// 可用区列表
	Zones []*Zone `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZonesResponseParams `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateBandwidthPackageFromLoadBalancerRequestParams struct {
	// 共享带宽包 ID。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 客户端Token，用于保证请求的幂等性。
	// 
	// 从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符。
	// 
	// > 若您未指定，则系统自动使用API请求的**RequestId**作为**ClientToken**标识。每次API请求的**RequestId**不一样。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否只预检此次请求。取值：
	// - **true**：发送检查请求，不会将共享带宽包从负载均衡实例中移除。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回错误码`DryRunOperation`。
	// - **false**（默认值）：发送正常请求，通过检查后返回HTTP 2xx状态码并直接进行操作。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type DisassociateBandwidthPackageFromLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 共享带宽包 ID。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 客户端Token，用于保证请求的幂等性。
	// 
	// 从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符。
	// 
	// > 若您未指定，则系统自动使用API请求的**RequestId**作为**ClientToken**标识。每次API请求的**RequestId**不一样。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否只预检此次请求。取值：
	// - **true**：发送检查请求，不会将共享带宽包从负载均衡实例中移除。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回错误码`DryRunOperation`。
	// - **false**（默认值）：发送正常请求，通过检查后返回HTTP 2xx状态码并直接进行操作。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *DisassociateBandwidthPackageFromLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateBandwidthPackageFromLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BandwidthPackageId")
	delete(f, "LoadBalancerId")
	delete(f, "ClientToken")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateBandwidthPackageFromLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateBandwidthPackageFromLoadBalancerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateBandwidthPackageFromLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateBandwidthPackageFromLoadBalancerResponseParams `json:"Response"`
}

func (r *DisassociateBandwidthPackageFromLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateBandwidthPackageFromLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateListenerAdditionalCertificatesRequestParams struct {
	// 待解绑的扩展证书 ID 列表。
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 客户端 Token，用于保证请求的幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken 只支持 ASCII 字符。
	// 若您未指定，则系统自动使用 API 请求的 RequestId 作为 ClientToken 标识。每次 API 请求的 RequestId 不一样。  
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否只预检此次请求，取值：
	// true：发送检查请求，不会从 HTTPS和QUIC监听器解绑扩展证书。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回错误码DryRunOperation。
	// false（默认值）：发送正常请求，通过检查后返回HTTP 2xx状态码并直接进行操作。
	DryRun *string `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type DisassociateListenerAdditionalCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// 待解绑的扩展证书 ID 列表。
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 客户端 Token，用于保证请求的幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken 只支持 ASCII 字符。
	// 若您未指定，则系统自动使用 API 请求的 RequestId 作为 ClientToken 标识。每次 API 请求的 RequestId 不一样。  
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 是否只预检此次请求，取值：
	// true：发送检查请求，不会从 HTTPS和QUIC监听器解绑扩展证书。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回错误码DryRunOperation。
	// false（默认值）：发送正常请求，通过检查后返回HTTP 2xx状态码并直接进行操作。
	DryRun *string `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *DisassociateListenerAdditionalCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateListenerAdditionalCertificatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CertificateIds")
	delete(f, "ListenerId")
	delete(f, "LoadBalancerId")
	delete(f, "ClientToken")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisassociateListenerAdditionalCertificatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisassociateListenerAdditionalCertificatesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisassociateListenerAdditionalCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *DisassociateListenerAdditionalCertificatesResponseParams `json:"Response"`
}

func (r *DisassociateListenerAdditionalCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisassociateListenerAdditionalCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 过滤器的名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤器的值数组
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FixedResponseInfo struct {
	// 返回的HTTP响应码，支持 2xx、4xx、5xx。
	HttpCode *int64 `json:"HttpCode,omitnil,omitempty" name:"HttpCode"`

	// 返回的固定内容。只支持 ASCII 字符，最大1KB。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 返回固定内容的格式。
	// 取值：text/plain、text/css、text/html、application/javascript或application/json。
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`
}

type HTTPCookieInfo struct {
	// Cookie的键，长度1~64个字符，支持字母、数字、下划线。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Cookie的值，长度1~128个字符，支持可打印字符。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type HTTPHeaderInfo struct {
	// HTTP Header的键，长度1 ~ 40个字符，支持的字符集为：a-z A-Z 0-9 - _ 。
	// 不支持中文，不支持Host，Cookie。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// HTTP Header的值，长度1 ~ 128个字符，支持可打印字符。
	// 不支持"，开头和结尾不能是空格，结尾不能是\。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type HTTPQueryStringInfo struct {
	// 查询字符串的键，长度1 ~ 16个字符。支持可打印字符，不支持空格和#[]{}\|<>&。
	// 支持 * 多字符通配，? 单字符通配。
	// 
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 查询字符串的值，长度1 ~ 128字符，支持可打印字符，不支持空格和#[]{}\|<>&。
	// 支持 * 多字符通配，? 单字符通配。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type HTTPRedirectInfo struct {
	// <p>重定向的HTTP码，支持301、302、303、307、 308。</p>
	HttpCode *int64 `json:"HttpCode,omitnil,omitempty" name:"HttpCode"`

	// <p>重定向的主机地址，默认值${host}。长度3 ~ 128个字符，支持的字符集为：a-z 0-9 _ . -。</p>
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// <p>重定向的路径，默认值${path}。长度1 ~ 128个字符，支持的字符集为：a-z A-Z 0-9  ? =  _  . - / : 。</p>
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// <p>重定向的端口，默认值 ${port}。取值1 ~ 65535。</p>
	Port *string `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>重定向的协议，取值：HTTP,HTTPS，默认值${protocol}。</p>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>重定向的查询字符串，默认值${query}。长度1 ~ 128字符，支持可打印字符，不支持 #[]{}|&lt;&gt;&amp; 和空格。</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type HTTPRewriteInfo struct {
	// <p>重写的主机地址，默认值${host}。长度3 ~ 128个字符，支持的字符集为：a-z 0-9 _ . -。</p>
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// <p>重写的路径，默认值${path}。长度1 ~ 128个字符，支持的字符集为：a-z A-Z 0-9 ? = _ . - / : 。</p>
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// <p>重写的查询字符串，默认值${query}。长度1 ~ 128字符，支持可打印字符，不支持 #[]{}|&lt;&gt;&amp; 和空格。</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type HealthCheckConfig struct {
	// 是否开启健康检查。
	// - **true**：开启。
	// - **false**：不开启。
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitnil,omitempty" name:"HealthCheckEnabled"`

	// 健康检查状态码。取值：
	// - 当健康检查协议为**HTTP/HTTPS**时：
	// 	- **http_1xx**
	// 	- **http_2xx**（默认值）
	// 	-  **http_3xx**
	// 	-  **http_4xx**
	// 	-  **http_5xx**
	// - 当健康检查协议为**gRPC**时：默认值为12，数值范围为0-99，输入值可为数值、多个数值或者范围以及相互组合，如：
	// 	- **"20"**
	// 	- **"0-99"**
	// > 仅当**HealthCheckProtocol**设置为**HTTP** 、**HTTPS**、**GRPC** 或者**GRPCS**时，该参数生效。
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitnil,omitempty" name:"HealthCheckCodes"`

	// 判定后端服务健康的阈值，当健康检查连续成功多少次后，后端服务的状态由**不健康**变为**健康**。
	// 取值范围：**2**-**10**。
	// 默认值：**2**。
	HealthCheckHealthyThreshold *uint64 `json:"HealthCheckHealthyThreshold,omitnil,omitempty" name:"HealthCheckHealthyThreshold"`

	// 健康检查域名。该参数不设置时默认使用后端服务的内网IP作为健康检查地址。
	// 域名限制：
	// - 长度限制为 **1-255** 个字符。
	// - 可包含小写字母、数字、短划线（-）和半角句号（.）。
	// - 至少包含一个半角句号（.），半角句号（.）不能出现在开头或结尾。
	// - 最右侧的域标签，只能包含字母，不能包含数字或短划线（-）。
	// - 短划线（-）不能出现在开头或结尾。
	// >仅当 **HealthCheckProtocol** 设置为 **HTTP、HTTPS** 、**GRPC**、**GRPCS** 时，该参数生效。
	HealthCheckHost *string `json:"HealthCheckHost,omitnil,omitempty" name:"HealthCheckHost"`

	// 健康检查 HTTP 协议版本，取值：
	// - **HTTP1.1**（默认）
	// - **HTTP1.0** 
	// > 仅当**HealthCheckProtocol**设置为**HTTP** 或 **HTTPS** 时，该参数生效。
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitnil,omitempty" name:"HealthCheckHttpVersion"`

	// 健康检查的时间间隔。单位：秒。
	// 取值范围：**2**-**300**。
	// 默认值：**5**。
	HealthCheckInterval *uint64 `json:"HealthCheckInterval,omitnil,omitempty" name:"HealthCheckInterval"`

	// 健康检查方法，取值：
	// - **GET**
	// - **HEAD**（默认值）
	// > 仅当**HealthCheckProtocol**设置为**HTTP** 或 **HTTPS** 时，该参数生效。
	HealthCheckMethod *string `json:"HealthCheckMethod,omitnil,omitempty" name:"HealthCheckMethod"`

	// 健康检查的转发规则路径。
	// 长度为 1~80 个字符，只能使用字母、数字、字符`-/.%?#&=`以及扩展字符`_;~!（)*[]@$^:',+`。 URL 必须以正斜线（/）开头。
	// > 仅当**HealthCheckProtocol**为**HTTP**、**HTTPS** 、**GRPC**、**GRPCS**时，转发规则路径参数生效。
	HealthCheckPath *string `json:"HealthCheckPath,omitnil,omitempty" name:"HealthCheckPath"`

	// 健康检查访问后端服务器的端口。
	// 
	// 取值范围：**0-65535**。
	// 
	// 默认值：**0**，表示后端服务器的端口。
	HealthCheckPort *uint64 `json:"HealthCheckPort,omitnil,omitempty" name:"HealthCheckPort"`

	// 健康检查协议。取值：
	// - **HTTP**（默认）：通过发送 HEAD 或 GET 请求模拟浏览器的访问行为来检查服务器应用是否健康。
	// - **HTTPS**：通过发送 HEAD 或 GET 请求模拟浏览器的访问行为来检查服务器应用是否健康。（数据加密，相比 HTTP 更安全。）
	// - **TCP**：通过发送 SYN 握手报文来检测服务器端口是否存活。
	// - **GRPC**：通过发送 POST 请求来检查服务器应用是否健康。
	// - **GRPCS**：通过发送 POST 请求来检查服务器应用是否健康。
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitnil,omitempty" name:"HealthCheckProtocol"`

	// 健康检查的响应超时时间。单位：秒。
	// 取值范围：**2**-**60**。
	// 默认值：**2**。
	HealthCheckTimeout *uint64 `json:"HealthCheckTimeout,omitnil,omitempty" name:"HealthCheckTimeout"`

	// 判定后端服务不健康的阈值，当健康检查连续失败多少次后，后端服务的状态由**健康**变为**不健康**。
	// 取值范围：**2**-**10**。
	// 默认值：**2**。
	HealthCheckUnhealthyThreshold *uint64 `json:"HealthCheckUnhealthyThreshold,omitnil,omitempty" name:"HealthCheckUnhealthyThreshold"`
}

type HealthCheckTemplate struct {
	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 健康检查状态码。取值：
	// - 当健康检查协议为**HTTP/HTTPS**时：
	// 	- **http_1xx**
	// 	- **http_2xx**（默认值）
	// 	-  **http_3xx**
	// 	-  **http_4xx**
	// 	-  **http_5xx**
	// - 当健康检查协议为**GRPC/GRPCS**时：默认值为**12**，数值范围为**0-99**，输入值可为数值、多个数值或者范围以及相互组合，如：
	// 	- **"20"**
	// 	- **"0-99"**
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitnil,omitempty" name:"HealthCheckCodes"`

	// 判定后端服务健康的阈值，当健康检查连续成功多少次后，后端服务的状态由**不健康**变为**健康**。
	// 取值范围：**2**-**10**。
	// 默认值：**2**。
	HealthCheckHealthyThreshold *uint64 `json:"HealthCheckHealthyThreshold,omitnil,omitempty" name:"HealthCheckHealthyThreshold"`

	// 健康检查域名。
	// 长度限制为 **1-255** 个字符。
	// 可包含小写字母、数字、短划线（-）和半角句号（.）。
	// 
	// > 仅当 **HealthCheckProtocol** 设置为 **HTTP/HTTPS/GRPC/GRPCS** 时，该参数生效。
	HealthCheckHost *string `json:"HealthCheckHost,omitnil,omitempty" name:"HealthCheckHost"`

	// 健康检查 HTTP 协议版本，取值：
	// - **HTTP1.1**（默认）
	// - **HTTP1.0** 
	// > 仅当**HealthCheckProtocol**设置为**HTTP** 或 **HTTPS** 时，该参数生效。
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitnil,omitempty" name:"HealthCheckHttpVersion"`

	// 健康检查的时间间隔。单位：秒。
	// 取值范围：**2**-**300**。
	// 默认值：**5**。
	HealthCheckInterval *uint64 `json:"HealthCheckInterval,omitnil,omitempty" name:"HealthCheckInterval"`

	// 健康检查方法，取值：
	// - **GET**
	// - **HEAD**（默认值）
	// > 仅当**HealthCheckProtocol**设置为**HTTP** 或 **HTTPS** 时，该参数生效。
	HealthCheckMethod *string `json:"HealthCheckMethod,omitnil,omitempty" name:"HealthCheckMethod"`

	// 健康检查的转发规则路径。 长度为 **1-80** 个字符，只能使用字母、数字、字符`-/.%?#&=`以及扩展字符`_;~!（)*[]@$^:',+`。 URL 必须以正斜线（/）开头。 
	// > 仅当**HealthCheckProtocol**为**HTTP/HTTPS/GRPC/GRPCS**时，转发规则路径参数生效。
	HealthCheckPath *string `json:"HealthCheckPath,omitnil,omitempty" name:"HealthCheckPath"`

	// 健康检查访问后端服务器的端口。
	// 
	// 取值范围：**0-65535**。
	// 
	// 默认值：**0**，表示后端服务器的端口。
	HealthCheckPort *uint64 `json:"HealthCheckPort,omitnil,omitempty" name:"HealthCheckPort"`

	// 健康检查协议。取值：
	// - **HTTP**（默认）：通过发送 HEAD 或 GET 请求模拟浏览器的访问行为来检查服务器应用是否健康。
	// - **HTTPS**：通过发送 HEAD 或 GET 请求模拟浏览器的访问行为来检查服务器应用是否健康。（数据加密，相比 HTTP 更安全。）
	// - **TCP**：通过发送 SYN 握手报文来检测服务器端口是否存活。
	// - **GRPC**：通过发送 POST 或 GET 请求来检查服务器应用是否健康。
	// - **GRPCS**：通过发送 POST 或 GET 请求来检查服务器应用是否健康。
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitnil,omitempty" name:"HealthCheckProtocol"`

	// 健康检查模板 ID，格式为 hct- 后接字母数字。所有接口（创建、查询、修改、删除）均使用 hct- 前缀。
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitnil,omitempty" name:"HealthCheckTemplateId"`

	// 健康检查模板名称。长度为 **1-255** 个字符，可包含数字、大小写字母、中文、半角句号（.）、下划线（_）和短划线（-）。
	HealthCheckTemplateName *string `json:"HealthCheckTemplateName,omitnil,omitempty" name:"HealthCheckTemplateName"`

	// 健康检查的响应超时时间。单位：秒。
	// 取值范围：**2**-**60**。
	// 默认值：**2**。
	HealthCheckTimeout *uint64 `json:"HealthCheckTimeout,omitnil,omitempty" name:"HealthCheckTimeout"`

	// 判定后端服务不健康的阈值，当健康检查连续失败多少次后，后端服务的状态由**健康**变为**不健康**。
	// 取值范围：**2**-**10**。
	// 默认值：**2**。
	HealthCheckUnhealthyThreshold *uint64 `json:"HealthCheckUnhealthyThreshold,omitnil,omitempty" name:"HealthCheckUnhealthyThreshold"`

	// 修改时间。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 标签。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type IPAddressInfo struct {
	// IP 地址
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// EIP AddressId
	AddressId *string `json:"AddressId,omitnil,omitempty" name:"AddressId"`
}

// Predefined struct for user
type InquirePriceCreateLoadBalancerRequestParams struct {
	// 实例的计费类型。默认POSTPAID_BY_HOUR，仅取值 POSTPAID_BY_HOUR：表示按量计费。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`
}

type InquirePriceCreateLoadBalancerRequest struct {
	*tchttp.BaseRequest
	
	// 实例的计费类型。默认POSTPAID_BY_HOUR，仅取值 POSTPAID_BY_HOUR：表示按量计费。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`
}

func (r *InquirePriceCreateLoadBalancerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateLoadBalancerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChargeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquirePriceCreateLoadBalancerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquirePriceCreateLoadBalancerResponseParams struct {
	// 询价结果。
	Price *Price `json:"Price,omitnil,omitempty" name:"Price"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type InquirePriceCreateLoadBalancerResponse struct {
	*tchttp.BaseResponse
	Response *InquirePriceCreateLoadBalancerResponseParams `json:"Response"`
}

func (r *InquirePriceCreateLoadBalancerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquirePriceCreateLoadBalancerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InsertHTTPHeaderInfo struct {
	// 插入的HTTP Header键，长度1 ~ 40个字符，支持的字符集为：a-z A-Z 0-9 - _ 。
	// 不支持中文，不支持Cookie,Host,Content-Length,Connection,Upgrade,transfer-encoding,keep-alive,te,authority,x-forwarded-for,x-forwarded-proto,x-forwarded-host,x-forwarded-port。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// HTTP Header值的类型。
	// ValueType为SystemDefined时，取值范围 ClientPort：客户端端口，ClientIp：客户端 IP 地址，Protocol：客户端请求的协议，CLBPort：负载均衡实例监听端口。
	// ValueType为UserDefined时，长度1 ~ 128的可打印字符，不支持"，开头和结尾不能为空格，结尾不能为\。
	// ValueType为ReferenceHeader时，引用请求头中的某一个header，长度1~128的可打印字符，不支持"，开头和结尾不能为空格，结尾不能为\。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// HTTP Header值的类型，取值：
	// SystemDefined：系统定义的header。
	// UserDefined：用户自定义的header。
	// ReferenceHeader：引用请求头中的某一个header。
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`
}

type Job struct {
	// 操作接口名称。
	ApiName *string `json:"ApiName,omitnil,omitempty" name:"ApiName"`

	// 任务流Id
	FlowId *uint64 `json:"FlowId,omitnil,omitempty" name:"FlowId"`

	// 任务请求Id。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// 资源ID列表
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 任务状态。取值：Processing、Succeeded、Failed。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ListenerOutput struct {
	// 是否开启双向认证。
	CaEnable *bool `json:"CaEnable,omitnil,omitempty" name:"CaEnable"`

	// 监听器实例的创建时间。格式：ISO 8601（例如 2025-01-01T08:30:00+08:00）
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否启用 Gzip 压缩。
	GzipEnabled *bool `json:"GzipEnabled,omitnil,omitempty" name:"GzipEnabled"`

	// 是否启用http2。
	Http2Enable *bool `json:"Http2Enable,omitnil,omitempty" name:"Http2Enable"`

	// 空闲超时时间。
	IdleTimeout *uint64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`

	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 监听器名称。
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// 监听器端口。
	ListenerPort *uint64 `json:"ListenerPort,omitnil,omitempty" name:"ListenerPort"`

	// 监听器协议。
	ListenerProtocol *string `json:"ListenerProtocol,omitnil,omitempty" name:"ListenerProtocol"`

	// 监听器状态。取值:=
	// 
	// - **Active**: 运行中。
	// - **Provisioning**：创建中。
	// - **Configuring**：变配中。
	// - **ProvisionFailed**：创建失败
	ListenerStatus *string `json:"ListenerStatus,omitnil,omitempty" name:"ListenerStatus"`

	// 监听器实例的最后变更时间。格式：ISO 8601（例如 2025-01-01T08:30:00+08:00）
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 请求超时时间。
	RequestTimeout *uint64 `json:"RequestTimeout,omitnil,omitempty" name:"RequestTimeout"`

	// 标签。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 安全策略 ID。
	TlsSecurityPolicyId *string `json:"TlsSecurityPolicyId,omitnil,omitempty" name:"TlsSecurityPolicyId"`

	// XForwardedFor配置。
	XForwardedForConfig *XForwardedForConfig `json:"XForwardedForConfig,omitnil,omitempty" name:"XForwardedForConfig"`
}

type LoadBalancer struct {
	// 访问日志配置结构。
	AccessLogConfig *AccessLogConfig `json:"AccessLogConfig,omitnil,omitempty" name:"AccessLogConfig"`

	// IP 地址版本，取值 IPv4 或 IPv6。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressIpVersion *string `json:"AddressIpVersion,omitnil,omitempty" name:"AddressIpVersion"`

	// 负载均衡的地址类型。取值：
	// 
	// - **Internet**：负载均衡具有公网IP地址，DNS域名被解析到公网IP，因此可以在公网环境访问。
	// 
	// - **Intranet**：负载均衡只有私网IP地址，DNS域名被解析到私网IP，因此只能被负载均衡所在VPC的内网环境访问。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddressType *string `json:"AddressType,omitnil,omitempty" name:"AddressType"`

	// 资源创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 删除保护设置信息。
	DeletionProtection *DeletionProtectionConfig `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`

	// DNS域名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 负载均衡实例计费配置。
	LoadBalancerBillingConfig *LoadBalancerBillingConfig `json:"LoadBalancerBillingConfig,omitnil,omitempty" name:"LoadBalancerBillingConfig"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 负载均衡实例名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 负载均衡操作锁配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerOperationLocks []*LoadBalancerOperationLocksItem `json:"LoadBalancerOperationLocks,omitnil,omitempty" name:"LoadBalancerOperationLocks"`

	// 应用型负载均衡实例状态。取值：
	// 
	// - **Provisioning**：创建中。
	// - **Active**: 运行中。
	// - **Configuring**: 变配中。
	// - **Deleting**：删除中。
	// - **ProvisionFailed**：创建失败。
	// - **ConfigureFailed**：变配失败。
	// - **DeletionFailed**：删除失败。
	// - **Abnormal**：异常状态，具体异常原因参见LoadBalancerOperationLocks字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitnil,omitempty" name:"LoadBalancerStatus"`

	// 修改保护设置信息。
	ModificationProtection *ModificationProtectionInfo `json:"ModificationProtection,omitnil,omitempty" name:"ModificationProtection"`

	// 标签列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 私有网络 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type LoadBalancerAddress struct {
	// IPv4 地址列表
	IPv4Address []*IPAddressInfo `json:"IPv4Address,omitnil,omitempty" name:"IPv4Address"`

	// IPv6 地址列表
	IPv6Address []*IPAddressInfo `json:"IPv6Address,omitnil,omitempty" name:"IPv6Address"`
}

type LoadBalancerBillingConfig struct {
	// 实例的计费类型。
	// 
	// 取值**POSTPAID_BY_HOUR**：表示按量计费。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 共享带宽包 ID。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`
}

type LoadBalancerDetail struct {
	// 访问日志配置。
	AccessLogConfig *AccessLogConfig `json:"AccessLogConfig,omitnil,omitempty" name:"AccessLogConfig"`

	// IP 地址版本，取值 IPv4 或 IPv6。
	AddressIpVersion *string `json:"AddressIpVersion,omitnil,omitempty" name:"AddressIpVersion"`

	// 应用型负载均衡实例的网络地址类型。取值：
	// 
	// - **Internet/Public**：负载均衡具有公网IP地址，DNS域名被解析到公网IP，因此可以在公网环境访问。
	// 
	// - **Intranet/Internal**：负载均衡只有私网IP地址，DNS域名被解析到私网IP，因此只能被负载均衡所在VPC的内网环境访问。
	// 
	AddressType *string `json:"AddressType,omitnil,omitempty" name:"AddressType"`

	// 资源创建时间，格式为`yyyy-MM-ddTHH:mm:ss±hh:mm`。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 删除保护设置信息。
	DeletionProtection *DeletionProtectionConfig `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`

	// DNS域名。
	Domain *string `json:"Domain,omitnil,omitempty" name:"Domain"`

	// 负载均衡实例付计费配置信息
	LoadBalancerBillingConfig *LoadBalancerBillingConfig `json:"LoadBalancerBillingConfig,omitnil,omitempty" name:"LoadBalancerBillingConfig"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 实例名称。
	// 
	// 长度为1~80个字符，可包含中文、字母、数字、短划线（-）、正斜线（/）、半角句号（.）和下划线（_）。
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`

	// 应用型负载均衡操作锁配置。
	LoadBalancerOperationLocks []*LoadBalancerOperationLocksItem `json:"LoadBalancerOperationLocks,omitnil,omitempty" name:"LoadBalancerOperationLocks"`

	// 应用型负载均衡实例状态。取值：
	// 
	// - **Provisioning**：创建中。
	// - **Active**: 运行中。
	// - **Configuring**: 变配中。
	// - **Deleting**：删除中。
	// - **ProvisionFailed**：创建失败。
	// - **ConfigureFailed**：变配失败。
	// - **DeletionFailed**：删除失败。
	// - **Abnormal**：异常状态，具体异常原因参见LoadBalancerOperationLocks字段。
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitnil,omitempty" name:"LoadBalancerStatus"`

	// 修改保护设置信息。
	ModificationProtection *ModificationProtectionInfo `json:"ModificationProtection,omitnil,omitempty" name:"ModificationProtection"`

	// 应用型负载均衡实例绑定的安全组ID集合。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 标签。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 私有网络 ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 可用区及子网映射列表，最多返回10个可用区。若当前地域支持2个及以上可用区，至少返回2个及以上可用区。
	ZoneMappings []*ZoneMappingInfo `json:"ZoneMappings,omitnil,omitempty" name:"ZoneMappings"`
}

type LoadBalancerOperationLocksItem struct {
	// 锁定的原因。在**LoadBalancerStatus**为**Abnormal**时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LockReason *string `json:"LockReason,omitnil,omitempty" name:"LockReason"`

	// 锁定的类型。取值 ：
	// 
	// - **SecurityLocked**：安全锁定。
	// 
	// - **RelatedResourceLocked**：关联锁定。
	// 
	// - **FinancialLocked**：欠费锁定。
	// 
	// - **ResidualLocked**：残留锁定。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LockType *string `json:"LockType,omitnil,omitempty" name:"LockType"`
}

type ModificationProtectionInfo struct {
	// 是否开启修改保护。开启后，可防止实例被意外修改或删除。
	// - true：开启修改保护
	// - false：关闭修改保护
	ModificationProtectionEnabled *bool `json:"ModificationProtectionEnabled,omitnil,omitempty" name:"ModificationProtectionEnabled"`

	// 1238716123
	OperatorUin *string `json:"OperatorUin,omitnil,omitempty" name:"OperatorUin"`

	// 开启修改保护的原因说明。
	// 长度为 1~255 个字符，必须是中文和无害字符串中的字符， 可包含中文、字母、数字、短划线（-）、正斜线（/）、半角句号（.）、下划线（_）。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

// Predefined struct for user
type ModifyHealthCheckTemplateRequestParams struct {
	// <p>健康检查模板 ID，格式为 hct- 后接字母数字。</p>
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitnil,omitempty" name:"HealthCheckTemplateId"`

	// <p>是否预览此次请求。</p><ul><li><strong>false</strong>（默认）：发送普通请求，直接修改健康检查模板。</li><li><strong>true</strong>：发送预览请求，检查修改健康检查模板的参数、格式、业务限制等是否符合要求。</li></ul>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// <p>健康检查状态码。取值：</p><ul><li>当健康检查协议为<strong>HTTP/HTTPS</strong>时：<ul><li><strong>http_1xx</strong></li><li><strong>http_2xx</strong>（默认值）</li><li><strong>http_3xx</strong></li><li><strong>http_4xx</strong></li><li><strong>http_5xx</strong></li></ul></li><li>当健康检查协议为<strong>GRPC/GRPCS</strong>时：默认值为<strong>12</strong>，数值范围为<strong>0-99</strong>，输入值可为数值、多个数值或者范围以及相互组合，如：<ul><li><strong>&quot;20&quot;</strong></li><li><strong>&quot;0-99&quot;</strong></li></ul></li></ul>
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitnil,omitempty" name:"HealthCheckCodes"`

	// <p>判定后端服务健康的阈值，当健康检查连续成功多少次后，后端服务的状态由<strong>不健康</strong>变为<strong>健康</strong>。<br>取值范围：<strong>2</strong>-<strong>10</strong>。<br>默认值：<strong>2</strong>。</p>
	HealthCheckHealthyThreshold *uint64 `json:"HealthCheckHealthyThreshold,omitnil,omitempty" name:"HealthCheckHealthyThreshold"`

	// <p>健康检查域名。<br>长度限制为 <strong>1-255</strong> 个字符。<br>可包含小写字母、数字、短划线（-）和半角句号（.）。</p><blockquote><p>仅当 <strong>HealthCheckProtocol</strong> 设置为 <strong>HTTP/HTTPS/GRPC/GRPCS</strong> 时，该参数生效。</p></blockquote>
	HealthCheckHost *string `json:"HealthCheckHost,omitnil,omitempty" name:"HealthCheckHost"`

	// <p>健康检查 HTTP 协议版本，取值：</p><ul><li><strong>HTTP1.1</strong>（默认）</li><li><strong>HTTP1.0</strong> <blockquote><p>仅当<strong>HealthCheckProtocol</strong>设置为<strong>HTTP</strong> 或 <strong>HTTPS</strong> 时，该参数生效。</p></blockquote></li></ul>
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitnil,omitempty" name:"HealthCheckHttpVersion"`

	// <p>健康检查的时间间隔。单位：秒。 取值范围：<strong>2</strong>-<strong>300</strong>。 默认值：<strong>5</strong>。</p>
	HealthCheckInterval *uint64 `json:"HealthCheckInterval,omitnil,omitempty" name:"HealthCheckInterval"`

	// <p>健康检查方法，取值： - <strong>GET</strong> - <strong>HEAD</strong>（默认值） </p><blockquote><p>仅当<strong>HealthCheckProtocol</strong>设置为<strong>HTTP</strong> 或 <strong>HTTPS</strong> 时，该参数生效。</p></blockquote>
	HealthCheckMethod *string `json:"HealthCheckMethod,omitnil,omitempty" name:"HealthCheckMethod"`

	// <p>健康检查的转发规则路径。 长度为 <strong>1-80</strong> 个字符，只能使用字母、数字、字符<code>-/.%?#&amp;=</code>以及扩展字符<code>_;~!（)*[]@$^:&#39;,+</code>。 URL 必须以正斜线（/）开头。 </p><blockquote><p>仅当<strong>HealthCheckProtocol</strong>为<strong>HTTP/HTTPS/GRPC/GRPCS</strong>时，转发规则路径参数生效。</p></blockquote>
	HealthCheckPath *string `json:"HealthCheckPath,omitnil,omitempty" name:"HealthCheckPath"`

	// <p>健康检查访问后端服务器的端口。  取值范围：<strong>0-65535</strong>。  默认值：<strong>0</strong>，表示后端服务器的端口。</p>
	HealthCheckPort *uint64 `json:"HealthCheckPort,omitnil,omitempty" name:"HealthCheckPort"`

	// <p>健康检查协议。取值：</p><ul><li><strong>HTTP</strong>（默认）：通过发送 HEAD 或 GET 请求模拟浏览器的访问行为来检查服务器应用是否健康。</li><li><strong>HTTPS</strong>：通过发送 HEAD 或 GET 请求模拟浏览器的访问行为来检查服务器应用是否健康。（数据加密，相比 HTTP 更安全。）</li><li><strong>TCP</strong>：通过发送 SYN 握手报文来检测服务器端口是否存活。</li><li><strong>GRPC</strong>：通过发送 POST 或 GET 请求来检查服务器应用是否健康。</li><li><strong>GRPCS</strong>：通过发送 POST 或 GET 请求来检查服务器应用是否健康。</li></ul>
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitnil,omitempty" name:"HealthCheckProtocol"`

	// <p>健康检查模板名称。长度为 <strong>1-255</strong> 个字符，可包含数字、大小写字母、中文、半角句号（.）、下划线（_）和短划线（-）。</p>
	HealthCheckTemplateName *string `json:"HealthCheckTemplateName,omitnil,omitempty" name:"HealthCheckTemplateName"`

	// <p>健康检查的响应超时时间。单位：秒。<br>取值范围：<strong>2</strong>-<strong>60</strong>。<br>默认值：<strong>2</strong>。</p>
	HealthCheckTimeout *uint64 `json:"HealthCheckTimeout,omitnil,omitempty" name:"HealthCheckTimeout"`

	// <p>判定后端服务不健康的阈值，当健康检查连续失败多少次后，后端服务的状态由<strong>健康</strong>变为<strong>不健康</strong>。<br>取值范围：<strong>2</strong>-<strong>10</strong>。<br>默认值：<strong>2</strong>。</p>
	HealthCheckUnhealthyThreshold *uint64 `json:"HealthCheckUnhealthyThreshold,omitnil,omitempty" name:"HealthCheckUnhealthyThreshold"`
}

type ModifyHealthCheckTemplateRequest struct {
	*tchttp.BaseRequest
	
	// <p>健康检查模板 ID，格式为 hct- 后接字母数字。</p>
	HealthCheckTemplateId *string `json:"HealthCheckTemplateId,omitnil,omitempty" name:"HealthCheckTemplateId"`

	// <p>是否预览此次请求。</p><ul><li><strong>false</strong>（默认）：发送普通请求，直接修改健康检查模板。</li><li><strong>true</strong>：发送预览请求，检查修改健康检查模板的参数、格式、业务限制等是否符合要求。</li></ul>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// <p>健康检查状态码。取值：</p><ul><li>当健康检查协议为<strong>HTTP/HTTPS</strong>时：<ul><li><strong>http_1xx</strong></li><li><strong>http_2xx</strong>（默认值）</li><li><strong>http_3xx</strong></li><li><strong>http_4xx</strong></li><li><strong>http_5xx</strong></li></ul></li><li>当健康检查协议为<strong>GRPC/GRPCS</strong>时：默认值为<strong>12</strong>，数值范围为<strong>0-99</strong>，输入值可为数值、多个数值或者范围以及相互组合，如：<ul><li><strong>&quot;20&quot;</strong></li><li><strong>&quot;0-99&quot;</strong></li></ul></li></ul>
	HealthCheckCodes []*string `json:"HealthCheckCodes,omitnil,omitempty" name:"HealthCheckCodes"`

	// <p>判定后端服务健康的阈值，当健康检查连续成功多少次后，后端服务的状态由<strong>不健康</strong>变为<strong>健康</strong>。<br>取值范围：<strong>2</strong>-<strong>10</strong>。<br>默认值：<strong>2</strong>。</p>
	HealthCheckHealthyThreshold *uint64 `json:"HealthCheckHealthyThreshold,omitnil,omitempty" name:"HealthCheckHealthyThreshold"`

	// <p>健康检查域名。<br>长度限制为 <strong>1-255</strong> 个字符。<br>可包含小写字母、数字、短划线（-）和半角句号（.）。</p><blockquote><p>仅当 <strong>HealthCheckProtocol</strong> 设置为 <strong>HTTP/HTTPS/GRPC/GRPCS</strong> 时，该参数生效。</p></blockquote>
	HealthCheckHost *string `json:"HealthCheckHost,omitnil,omitempty" name:"HealthCheckHost"`

	// <p>健康检查 HTTP 协议版本，取值：</p><ul><li><strong>HTTP1.1</strong>（默认）</li><li><strong>HTTP1.0</strong> <blockquote><p>仅当<strong>HealthCheckProtocol</strong>设置为<strong>HTTP</strong> 或 <strong>HTTPS</strong> 时，该参数生效。</p></blockquote></li></ul>
	HealthCheckHttpVersion *string `json:"HealthCheckHttpVersion,omitnil,omitempty" name:"HealthCheckHttpVersion"`

	// <p>健康检查的时间间隔。单位：秒。 取值范围：<strong>2</strong>-<strong>300</strong>。 默认值：<strong>5</strong>。</p>
	HealthCheckInterval *uint64 `json:"HealthCheckInterval,omitnil,omitempty" name:"HealthCheckInterval"`

	// <p>健康检查方法，取值： - <strong>GET</strong> - <strong>HEAD</strong>（默认值） </p><blockquote><p>仅当<strong>HealthCheckProtocol</strong>设置为<strong>HTTP</strong> 或 <strong>HTTPS</strong> 时，该参数生效。</p></blockquote>
	HealthCheckMethod *string `json:"HealthCheckMethod,omitnil,omitempty" name:"HealthCheckMethod"`

	// <p>健康检查的转发规则路径。 长度为 <strong>1-80</strong> 个字符，只能使用字母、数字、字符<code>-/.%?#&amp;=</code>以及扩展字符<code>_;~!（)*[]@$^:&#39;,+</code>。 URL 必须以正斜线（/）开头。 </p><blockquote><p>仅当<strong>HealthCheckProtocol</strong>为<strong>HTTP/HTTPS/GRPC/GRPCS</strong>时，转发规则路径参数生效。</p></blockquote>
	HealthCheckPath *string `json:"HealthCheckPath,omitnil,omitempty" name:"HealthCheckPath"`

	// <p>健康检查访问后端服务器的端口。  取值范围：<strong>0-65535</strong>。  默认值：<strong>0</strong>，表示后端服务器的端口。</p>
	HealthCheckPort *uint64 `json:"HealthCheckPort,omitnil,omitempty" name:"HealthCheckPort"`

	// <p>健康检查协议。取值：</p><ul><li><strong>HTTP</strong>（默认）：通过发送 HEAD 或 GET 请求模拟浏览器的访问行为来检查服务器应用是否健康。</li><li><strong>HTTPS</strong>：通过发送 HEAD 或 GET 请求模拟浏览器的访问行为来检查服务器应用是否健康。（数据加密，相比 HTTP 更安全。）</li><li><strong>TCP</strong>：通过发送 SYN 握手报文来检测服务器端口是否存活。</li><li><strong>GRPC</strong>：通过发送 POST 或 GET 请求来检查服务器应用是否健康。</li><li><strong>GRPCS</strong>：通过发送 POST 或 GET 请求来检查服务器应用是否健康。</li></ul>
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitnil,omitempty" name:"HealthCheckProtocol"`

	// <p>健康检查模板名称。长度为 <strong>1-255</strong> 个字符，可包含数字、大小写字母、中文、半角句号（.）、下划线（_）和短划线（-）。</p>
	HealthCheckTemplateName *string `json:"HealthCheckTemplateName,omitnil,omitempty" name:"HealthCheckTemplateName"`

	// <p>健康检查的响应超时时间。单位：秒。<br>取值范围：<strong>2</strong>-<strong>60</strong>。<br>默认值：<strong>2</strong>。</p>
	HealthCheckTimeout *uint64 `json:"HealthCheckTimeout,omitnil,omitempty" name:"HealthCheckTimeout"`

	// <p>判定后端服务不健康的阈值，当健康检查连续失败多少次后，后端服务的状态由<strong>健康</strong>变为<strong>不健康</strong>。<br>取值范围：<strong>2</strong>-<strong>10</strong>。<br>默认值：<strong>2</strong>。</p>
	HealthCheckUnhealthyThreshold *uint64 `json:"HealthCheckUnhealthyThreshold,omitnil,omitempty" name:"HealthCheckUnhealthyThreshold"`
}

func (r *ModifyHealthCheckTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHealthCheckTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "HealthCheckTemplateId")
	delete(f, "DryRun")
	delete(f, "HealthCheckCodes")
	delete(f, "HealthCheckHealthyThreshold")
	delete(f, "HealthCheckHost")
	delete(f, "HealthCheckHttpVersion")
	delete(f, "HealthCheckInterval")
	delete(f, "HealthCheckMethod")
	delete(f, "HealthCheckPath")
	delete(f, "HealthCheckPort")
	delete(f, "HealthCheckProtocol")
	delete(f, "HealthCheckTemplateName")
	delete(f, "HealthCheckTimeout")
	delete(f, "HealthCheckUnhealthyThreshold")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyHealthCheckTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyHealthCheckTemplateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyHealthCheckTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyHealthCheckTemplateResponseParams `json:"Response"`
}

func (r *ModifyHealthCheckTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyHealthCheckTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyListenerAttributesRequestParams struct {
	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 监听器配置的CA证书ID列表。目前仅支持添加1个CA证书。
	CaCertificateIds []*string `json:"CaCertificateIds,omitnil,omitempty" name:"CaCertificateIds"`

	// 是否开启双向认证。
	// 取值：
	// true：开启。
	// false（默认值）：不开启。
	CaEnabled *bool `json:"CaEnabled,omitnil,omitempty" name:"CaEnabled"`

	// 服务器证书 ID 列表。
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 客户端Token，用于保证请求的幂等性。  
	// 
	// 从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 默认转发规则动作列表。目前监听器仅支持添加 1 个默认转发规则动作。
	DefaultActions []*DefaultAction `json:"DefaultActions,omitnil,omitempty" name:"DefaultActions"`

	// 是否启用 Gzip 压缩。
	GzipEnabled *bool `json:"GzipEnabled,omitnil,omitempty" name:"GzipEnabled"`

	// 是否开启HTTP/2特性。只有 HTTPS 协议支持此参数。
	Http2Enabled *bool `json:"Http2Enabled,omitnil,omitempty" name:"Http2Enabled"`

	// 指定连接空闲超时时间。单位：秒。
	// 取值范围：1~600。
	// 默认值：15。
	// 如果在设定时间内一直没有访问请求，负载均衡会暂时断开当前连接，下次请求来临时重新建立新的连接。
	IdleTimeout *uint64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`

	// 自定义监听名称。  长度为 1~255 个字符，必须是中文和无害字符串中的字符，  可包含中文、字母、数字、短划线（-）、正斜线（/）、半角句号（.）、下划线（_）。
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// 指定请求超时时间。单位：秒。
	// 取值：1~600。
	// 默认值：60。
	// 如果在超时时间内后端服务器一直没有响应，负载均衡将放弃等待，并给客户端返回HTTP 504错误码。
	RequestTimeout *uint64 `json:"RequestTimeout,omitnil,omitempty" name:"RequestTimeout"`

	// 安全策略 ID，格式为 tls- 后接 8 位字母数字。
	SecurityPolicyId *string `json:"SecurityPolicyId,omitnil,omitempty" name:"SecurityPolicyId"`

	// XForwardedFor配置。
	XForwardedForConfig *XForwardedForConfig `json:"XForwardedForConfig,omitnil,omitempty" name:"XForwardedForConfig"`
}

type ModifyListenerAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 监听器配置的CA证书ID列表。目前仅支持添加1个CA证书。
	CaCertificateIds []*string `json:"CaCertificateIds,omitnil,omitempty" name:"CaCertificateIds"`

	// 是否开启双向认证。
	// 取值：
	// true：开启。
	// false（默认值）：不开启。
	CaEnabled *bool `json:"CaEnabled,omitnil,omitempty" name:"CaEnabled"`

	// 服务器证书 ID 列表。
	CertificateIds []*string `json:"CertificateIds,omitnil,omitempty" name:"CertificateIds"`

	// 客户端Token，用于保证请求的幂等性。  
	// 
	// 从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 默认转发规则动作列表。目前监听器仅支持添加 1 个默认转发规则动作。
	DefaultActions []*DefaultAction `json:"DefaultActions,omitnil,omitempty" name:"DefaultActions"`

	// 是否启用 Gzip 压缩。
	GzipEnabled *bool `json:"GzipEnabled,omitnil,omitempty" name:"GzipEnabled"`

	// 是否开启HTTP/2特性。只有 HTTPS 协议支持此参数。
	Http2Enabled *bool `json:"Http2Enabled,omitnil,omitempty" name:"Http2Enabled"`

	// 指定连接空闲超时时间。单位：秒。
	// 取值范围：1~600。
	// 默认值：15。
	// 如果在设定时间内一直没有访问请求，负载均衡会暂时断开当前连接，下次请求来临时重新建立新的连接。
	IdleTimeout *uint64 `json:"IdleTimeout,omitnil,omitempty" name:"IdleTimeout"`

	// 自定义监听名称。  长度为 1~255 个字符，必须是中文和无害字符串中的字符，  可包含中文、字母、数字、短划线（-）、正斜线（/）、半角句号（.）、下划线（_）。
	ListenerName *string `json:"ListenerName,omitnil,omitempty" name:"ListenerName"`

	// 指定请求超时时间。单位：秒。
	// 取值：1~600。
	// 默认值：60。
	// 如果在超时时间内后端服务器一直没有响应，负载均衡将放弃等待，并给客户端返回HTTP 504错误码。
	RequestTimeout *uint64 `json:"RequestTimeout,omitnil,omitempty" name:"RequestTimeout"`

	// 安全策略 ID，格式为 tls- 后接 8 位字母数字。
	SecurityPolicyId *string `json:"SecurityPolicyId,omitnil,omitempty" name:"SecurityPolicyId"`

	// XForwardedFor配置。
	XForwardedForConfig *XForwardedForConfig `json:"XForwardedForConfig,omitnil,omitempty" name:"XForwardedForConfig"`
}

func (r *ModifyListenerAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyListenerAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "LoadBalancerId")
	delete(f, "CaCertificateIds")
	delete(f, "CaEnabled")
	delete(f, "CertificateIds")
	delete(f, "ClientToken")
	delete(f, "DefaultActions")
	delete(f, "GzipEnabled")
	delete(f, "Http2Enabled")
	delete(f, "IdleTimeout")
	delete(f, "ListenerName")
	delete(f, "RequestTimeout")
	delete(f, "SecurityPolicyId")
	delete(f, "XForwardedForConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyListenerAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyListenerAttributesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyListenerAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyListenerAttributesResponseParams `json:"Response"`
}

func (r *ModifyListenerAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyListenerAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerAddressTypeRequestParams struct {
	// 目标网络类型。取值：
	// - **Internet**（公网型）
	// 负载均衡实例分配公网 IP 地址，域名（DNS）解析至公网 IP，可在公网环境中直接访问，适用于对外提供服务的业务场景。
	// - **Intranet**（内网型）
	// 负载均衡实例仅分配私网 IP 地址，域名（DNS）解析至私网 IP，仅支持在负载均衡实例所属 VPC 内的内网环境访问，适用于内部业务或对安全性要求较高的场景。
	AddressType *string `json:"AddressType,omitnil,omitempty" name:"AddressType"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 共享带宽包 ID。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// 是否只预检此次请求。取值：
	// - **true**：发送检查请求，不会更新实例的网络类型。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回错误码DryRunOperation。
	// - **false**（默认值）：发送正常请求，通过检查后返回 HTTP 2xx 状态码并直接进行操作。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 可用区及子网映射结构体。
	// 若当前地域支持 2 个及以上的可用区，至少需要添加 2 个可用区。
	ZoneMappings []*ZoneMappingsItem `json:"ZoneMappings,omitnil,omitempty" name:"ZoneMappings"`
}

type ModifyLoadBalancerAddressTypeRequest struct {
	*tchttp.BaseRequest
	
	// 目标网络类型。取值：
	// - **Internet**（公网型）
	// 负载均衡实例分配公网 IP 地址，域名（DNS）解析至公网 IP，可在公网环境中直接访问，适用于对外提供服务的业务场景。
	// - **Intranet**（内网型）
	// 负载均衡实例仅分配私网 IP 地址，域名（DNS）解析至私网 IP，仅支持在负载均衡实例所属 VPC 内的内网环境访问，适用于内部业务或对安全性要求较高的场景。
	AddressType *string `json:"AddressType,omitnil,omitempty" name:"AddressType"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 共享带宽包 ID。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`

	// 是否只预检此次请求。取值：
	// - **true**：发送检查请求，不会更新实例的网络类型。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回错误码DryRunOperation。
	// - **false**（默认值）：发送正常请求，通过检查后返回 HTTP 2xx 状态码并直接进行操作。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 可用区及子网映射结构体。
	// 若当前地域支持 2 个及以上的可用区，至少需要添加 2 个可用区。
	ZoneMappings []*ZoneMappingsItem `json:"ZoneMappings,omitnil,omitempty" name:"ZoneMappings"`
}

func (r *ModifyLoadBalancerAddressTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerAddressTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddressType")
	delete(f, "LoadBalancerId")
	delete(f, "BandwidthPackageId")
	delete(f, "DryRun")
	delete(f, "ZoneMappings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerAddressTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerAddressTypeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLoadBalancerAddressTypeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancerAddressTypeResponseParams `json:"Response"`
}

func (r *ModifyLoadBalancerAddressTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerAddressTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerAttributesRequestParams struct {
	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 客户端Token，用于保证请求的幂等性。
	// 
	// 从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符。
	// 
	// > 若您未指定，则系统自动使用API请求的**RequestId**作为**ClientToken**标识。每次API请求的**RequestId**不一样。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 删除保护配置
	DeletionProtection *DeletionProtectionConfig `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`

	// 是否只预检此次请求，取值：
	// 
	// - **true**：发送检查请求，不会修改应用型负载均衡实例的属性。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回错误码`DryRunOperation`。
	// 
	// - **false**（默认值）：发送正常请求，通过检查后返回`HTTP_2xx`状态码并直接进行操作。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 应用型负载均衡实例名称。长度为1~80个字符，可包含中文、字母、数字、短划线（-）、正斜线（/）、半角句号（.）和下划线（_）。
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`
}

type ModifyLoadBalancerAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 客户端Token，用于保证请求的幂等性。
	// 
	// 从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符。
	// 
	// > 若您未指定，则系统自动使用API请求的**RequestId**作为**ClientToken**标识。每次API请求的**RequestId**不一样。
	ClientToken *string `json:"ClientToken,omitnil,omitempty" name:"ClientToken"`

	// 删除保护配置
	DeletionProtection *DeletionProtectionConfig `json:"DeletionProtection,omitnil,omitempty" name:"DeletionProtection"`

	// 是否只预检此次请求，取值：
	// 
	// - **true**：发送检查请求，不会修改应用型负载均衡实例的属性。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回错误码`DryRunOperation`。
	// 
	// - **false**（默认值）：发送正常请求，通过检查后返回`HTTP_2xx`状态码并直接进行操作。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 应用型负载均衡实例名称。长度为1~80个字符，可包含中文、字母、数字、短划线（-）、正斜线（/）、半角句号（.）和下划线（_）。
	LoadBalancerName *string `json:"LoadBalancerName,omitnil,omitempty" name:"LoadBalancerName"`
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
	delete(f, "ClientToken")
	delete(f, "DeletionProtection")
	delete(f, "DryRun")
	delete(f, "LoadBalancerName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerAttributesResponseParams struct {
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
type ModifyLoadBalancerModificationProtectionRequestParams struct {
	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 是否开启修改保护。开启后，可防止实例被意外修改或删除。\n- true：开启修改保护\n- false：关闭修改保护
	ModificationProtectionEnabled *bool `json:"ModificationProtectionEnabled,omitnil,omitempty" name:"ModificationProtectionEnabled"`

	// 是否只预检此次请求。取值：
	// - true：仅执行预检，不实际操作资源。检查参数完整性、请求格式及业务限制，通过返回 DryRunOperation，不通过返回对应错误。
	// - false（默认）：执行正常请求，检查通过后直接操作资源。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 开启修改保护的原因说明。
	// 长度为 1~255 个字符，必须是中文和无害字符串中的字符， 可包含中文、字母、数字、短划线（-）、正斜线（/）、半角句号（.）、下划线（_）。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type ModifyLoadBalancerModificationProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 是否开启修改保护。开启后，可防止实例被意外修改或删除。\n- true：开启修改保护\n- false：关闭修改保护
	ModificationProtectionEnabled *bool `json:"ModificationProtectionEnabled,omitnil,omitempty" name:"ModificationProtectionEnabled"`

	// 是否只预检此次请求。取值：
	// - true：仅执行预检，不实际操作资源。检查参数完整性、请求格式及业务限制，通过返回 DryRunOperation，不通过返回对应错误。
	// - false（默认）：执行正常请求，检查通过后直接操作资源。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 开启修改保护的原因说明。
	// 长度为 1~255 个字符，必须是中文和无害字符串中的字符， 可包含中文、字母、数字、短划线（-）、正斜线（/）、半角句号（.）、下划线（_）。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

func (r *ModifyLoadBalancerModificationProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerModificationProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoadBalancerId")
	delete(f, "ModificationProtectionEnabled")
	delete(f, "DryRun")
	delete(f, "Reason")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerModificationProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerModificationProtectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLoadBalancerModificationProtectionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancerModificationProtectionResponseParams `json:"Response"`
}

func (r *ModifyLoadBalancerModificationProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerModificationProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRulesAttributesRequestParams struct {
	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 转发规则列表。
	Rules []*RuleModify `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 是否只预检查此次请求。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type ModifyRulesAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 转发规则列表。
	Rules []*RuleModify `json:"Rules,omitnil,omitempty" name:"Rules"`

	// 是否只预检查此次请求。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *ModifyRulesAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRulesAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ListenerId")
	delete(f, "LoadBalancerId")
	delete(f, "Rules")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRulesAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRulesAttributesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRulesAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRulesAttributesResponseParams `json:"Response"`
}

func (r *ModifyRulesAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRulesAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityPolicyAttributesRequestParams struct {
	// <p>安全策略 ID，格式为 tls- 后接 8 位字母数字。</p>
	SecurityPolicyId *string `json:"SecurityPolicyId,omitnil,omitempty" name:"SecurityPolicyId"`

	// <p>修改后的加密套件列表。加密套件用于协商客户端与服务端之间的加密算法。</p><p><strong>配置说明：</strong></p><ul><li>加密套件的可选范围取决于所选的 TLS 协议版本（TLSVersions 参数）。</li><li>只要加密套件被任意一个已选 TLS 版本支持，即可添加到列表中。</li><li>若 TLSVersions 包含 TLSv1.3：可不指定 TLSv1.3 专属加密套件（系统将自动补全全部 TLSv1.3 套件）；若指定，则必须包含全部 TLSv1.3 专属加密套件，不支持仅指定部分。</li></ul><p><strong>获取可用加密套件：</strong><br>请调用 <a href="https://cloud.tencent.com/document/api/1822/133718">DescribeSecurityPolicyCapabilities</a> 接口查询各 TLS 版本支持的加密套件列表。</p><p><strong>注意：</strong> 若不传此参数，则保持原有配置不变。</p>
	Ciphers []*string `json:"Ciphers,omitnil,omitempty" name:"Ciphers"`

	// <p>是否仅执行预检请求。取值：</p><ul><li><strong>true</strong>：仅执行预检请求，不实际修改资源。预检请求将验证参数格式、权限及配置有效性等，帮助您在正式操作前发现潜在问题。</li><li><strong>false</strong>（默认）：执行正常请求，通过预检后将直接修改安全策略。</li></ul>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// <p>修改后的安全策略名称。用于标识和区分不同的安全策略。</p><p><strong>命名规则：</strong></p><ul><li>长度为 2~128 个字符。</li><li>必须以英文字母或中文开头。</li><li>可包含英文字母、中文、数字、半角句号（.）、下划线（_）和短划线（-）。</li></ul><p><strong>注意：</strong> 若不传此参数，则保持原有名称不变。</p>
	SecurityPolicyName *string `json:"SecurityPolicyName,omitnil,omitempty" name:"SecurityPolicyName"`

	// <p>修改后的 TLS 协议版本列表。TLS（Transport Layer Security）协议用于保障客户端与负载均衡之间的通信安全。</p><p><strong>可选值：</strong></p><ul><li><strong>TLSv1.0</strong>：兼容性最好，但安全性较低，不推荐在生产环境使用。</li><li><strong>TLSv1.1</strong>：安全性略优于 TLSv1.0，但仍不推荐。</li><li><strong>TLSv1.2</strong>：目前主流的安全协议版本，兼顾安全性与兼容性。</li><li><strong>TLSv1.3</strong>：最新版本，安全性最高，性能更优，推荐优先使用。</li></ul><p><strong>注意：</strong> </p><ul><li>若不传此参数，则保持原有配置不变。</li><li>修改 TLS 版本时，请同步检查 Ciphers 参数的配置是否兼容。</li></ul>
	TLSVersions []*string `json:"TLSVersions,omitnil,omitempty" name:"TLSVersions"`
}

type ModifySecurityPolicyAttributesRequest struct {
	*tchttp.BaseRequest
	
	// <p>安全策略 ID，格式为 tls- 后接 8 位字母数字。</p>
	SecurityPolicyId *string `json:"SecurityPolicyId,omitnil,omitempty" name:"SecurityPolicyId"`

	// <p>修改后的加密套件列表。加密套件用于协商客户端与服务端之间的加密算法。</p><p><strong>配置说明：</strong></p><ul><li>加密套件的可选范围取决于所选的 TLS 协议版本（TLSVersions 参数）。</li><li>只要加密套件被任意一个已选 TLS 版本支持，即可添加到列表中。</li><li>若 TLSVersions 包含 TLSv1.3：可不指定 TLSv1.3 专属加密套件（系统将自动补全全部 TLSv1.3 套件）；若指定，则必须包含全部 TLSv1.3 专属加密套件，不支持仅指定部分。</li></ul><p><strong>获取可用加密套件：</strong><br>请调用 <a href="https://cloud.tencent.com/document/api/1822/133718">DescribeSecurityPolicyCapabilities</a> 接口查询各 TLS 版本支持的加密套件列表。</p><p><strong>注意：</strong> 若不传此参数，则保持原有配置不变。</p>
	Ciphers []*string `json:"Ciphers,omitnil,omitempty" name:"Ciphers"`

	// <p>是否仅执行预检请求。取值：</p><ul><li><strong>true</strong>：仅执行预检请求，不实际修改资源。预检请求将验证参数格式、权限及配置有效性等，帮助您在正式操作前发现潜在问题。</li><li><strong>false</strong>（默认）：执行正常请求，通过预检后将直接修改安全策略。</li></ul>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// <p>修改后的安全策略名称。用于标识和区分不同的安全策略。</p><p><strong>命名规则：</strong></p><ul><li>长度为 2~128 个字符。</li><li>必须以英文字母或中文开头。</li><li>可包含英文字母、中文、数字、半角句号（.）、下划线（_）和短划线（-）。</li></ul><p><strong>注意：</strong> 若不传此参数，则保持原有名称不变。</p>
	SecurityPolicyName *string `json:"SecurityPolicyName,omitnil,omitempty" name:"SecurityPolicyName"`

	// <p>修改后的 TLS 协议版本列表。TLS（Transport Layer Security）协议用于保障客户端与负载均衡之间的通信安全。</p><p><strong>可选值：</strong></p><ul><li><strong>TLSv1.0</strong>：兼容性最好，但安全性较低，不推荐在生产环境使用。</li><li><strong>TLSv1.1</strong>：安全性略优于 TLSv1.0，但仍不推荐。</li><li><strong>TLSv1.2</strong>：目前主流的安全协议版本，兼顾安全性与兼容性。</li><li><strong>TLSv1.3</strong>：最新版本，安全性最高，性能更优，推荐优先使用。</li></ul><p><strong>注意：</strong> </p><ul><li>若不传此参数，则保持原有配置不变。</li><li>修改 TLS 版本时，请同步检查 Ciphers 参数的配置是否兼容。</li></ul>
	TLSVersions []*string `json:"TLSVersions,omitnil,omitempty" name:"TLSVersions"`
}

func (r *ModifySecurityPolicyAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityPolicyAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecurityPolicyId")
	delete(f, "Ciphers")
	delete(f, "DryRun")
	delete(f, "SecurityPolicyName")
	delete(f, "TLSVersions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySecurityPolicyAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySecurityPolicyAttributesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySecurityPolicyAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifySecurityPolicyAttributesResponseParams `json:"Response"`
}

func (r *ModifySecurityPolicyAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySecurityPolicyAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetGroupAttributesRequestParams struct {
	// <p>是否预览此次请求。</p><ul><li><strong>false</strong>（默认）：发送普通请求，直接修改目标组。</li><li><strong>true</strong>：发送预览请求，检查修改目标组的参数、格式、业务限制等是否符合要求。</li></ul>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// <p>健康检查配置。</p>
	HealthCheckConfig *HealthCheckConfig `json:"HealthCheckConfig,omitnil,omitempty" name:"HealthCheckConfig"`

	// <p>是否开启长连接。</p>
	KeepaliveEnabled *bool `json:"KeepaliveEnabled,omitnil,omitempty" name:"KeepaliveEnabled"`

	// <p>调度算法。取值：</p><ul><li><strong>wrr</strong>：加权轮询，按照权重选择后端服务器，权重越高的服务器被轮询到的概率越高。</li><li><strong>wlc</strong>：加权最小连接数，当不同后端服务器权重值相同时，当前连接数越小的后端服务器被轮询到的概率越高。</li></ul>
	SchedulerAlgorithm *string `json:"SchedulerAlgorithm,omitnil,omitempty" name:"SchedulerAlgorithm"`

	// <p>会话保持配置。</p>
	StickySessionConfig *StickySessionConfig `json:"StickySessionConfig,omitnil,omitempty" name:"StickySessionConfig"`

	// <p>目标组 ID，格式为 lbtg- 后接 8 位字母数字。</p>
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// <p>目标组名称。长度为 1~255 个字符，可包含数字、大小写字母、中文、半角句号（.）、下划线（_）和短划线（-）。不传目标组名称时默认使用ID作为目标组名称。</p>
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`
}

type ModifyTargetGroupAttributesRequest struct {
	*tchttp.BaseRequest
	
	// <p>是否预览此次请求。</p><ul><li><strong>false</strong>（默认）：发送普通请求，直接修改目标组。</li><li><strong>true</strong>：发送预览请求，检查修改目标组的参数、格式、业务限制等是否符合要求。</li></ul>
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// <p>健康检查配置。</p>
	HealthCheckConfig *HealthCheckConfig `json:"HealthCheckConfig,omitnil,omitempty" name:"HealthCheckConfig"`

	// <p>是否开启长连接。</p>
	KeepaliveEnabled *bool `json:"KeepaliveEnabled,omitnil,omitempty" name:"KeepaliveEnabled"`

	// <p>调度算法。取值：</p><ul><li><strong>wrr</strong>：加权轮询，按照权重选择后端服务器，权重越高的服务器被轮询到的概率越高。</li><li><strong>wlc</strong>：加权最小连接数，当不同后端服务器权重值相同时，当前连接数越小的后端服务器被轮询到的概率越高。</li></ul>
	SchedulerAlgorithm *string `json:"SchedulerAlgorithm,omitnil,omitempty" name:"SchedulerAlgorithm"`

	// <p>会话保持配置。</p>
	StickySessionConfig *StickySessionConfig `json:"StickySessionConfig,omitnil,omitempty" name:"StickySessionConfig"`

	// <p>目标组 ID，格式为 lbtg- 后接 8 位字母数字。</p>
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// <p>目标组名称。长度为 1~255 个字符，可包含数字、大小写字母、中文、半角句号（.）、下划线（_）和短划线（-）。不传目标组名称时默认使用ID作为目标组名称。</p>
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`
}

func (r *ModifyTargetGroupAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetGroupAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DryRun")
	delete(f, "HealthCheckConfig")
	delete(f, "KeepaliveEnabled")
	delete(f, "SchedulerAlgorithm")
	delete(f, "StickySessionConfig")
	delete(f, "TargetGroupId")
	delete(f, "TargetGroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetGroupAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetGroupAttributesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTargetGroupAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTargetGroupAttributesResponseParams `json:"Response"`
}

func (r *ModifyTargetGroupAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetGroupAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetsInTargetGroupRequestParams struct {
	// 目标组 ID，格式为 lbtg- 后接 8 位字母数字。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 需要修改的后端服务列表。
	Targets []*TargetToModify `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 是否预览此次请求。 
	// - **false**（默认）：发送普通请求，直接修改后端服务信息。 
	// - **true**：发送预览请求，检查修改后端服务的参数、格式、业务限制等是否符合要求。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type ModifyTargetsInTargetGroupRequest struct {
	*tchttp.BaseRequest
	
	// 目标组 ID，格式为 lbtg- 后接 8 位字母数字。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 需要修改的后端服务列表。
	Targets []*TargetToModify `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 是否预览此次请求。 
	// - **false**（默认）：发送普通请求，直接修改后端服务信息。 
	// - **true**：发送预览请求，检查修改后端服务的参数、格式、业务限制等是否符合要求。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *ModifyTargetsInTargetGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetsInTargetGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "Targets")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTargetsInTargetGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTargetsInTargetGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyTargetsInTargetGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTargetsInTargetGroupResponseParams `json:"Response"`
}

func (r *ModifyTargetsInTargetGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTargetsInTargetGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type NotifyUnbindTargetRequestParams struct {
	// 后端服务的IP列表。
	// > **VpcId**（**NumericVpcId**）、**Ips**必须同时设置。
	Ips []*string `json:"Ips,omitnil,omitempty" name:"Ips"`

	// 后端服务所属VPC的数字ID。
	// > **VpcId**（**NumericVpcId**）、**Ips**必须同时设置。
	NumericVpcId *uint64 `json:"NumericVpcId,omitnil,omitempty" name:"NumericVpcId"`
}

type NotifyUnbindTargetRequest struct {
	*tchttp.BaseRequest
	
	// 后端服务的IP列表。
	// > **VpcId**（**NumericVpcId**）、**Ips**必须同时设置。
	Ips []*string `json:"Ips,omitnil,omitempty" name:"Ips"`

	// 后端服务所属VPC的数字ID。
	// > **VpcId**（**NumericVpcId**）、**Ips**必须同时设置。
	NumericVpcId *uint64 `json:"NumericVpcId,omitnil,omitempty" name:"NumericVpcId"`
}

func (r *NotifyUnbindTargetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *NotifyUnbindTargetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ips")
	delete(f, "NumericVpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "NotifyUnbindTargetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type NotifyUnbindTargetResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type NotifyUnbindTargetResponse struct {
	*tchttp.BaseResponse
	Response *NotifyUnbindTargetResponseParams `json:"Response"`
}

func (r *NotifyUnbindTargetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *NotifyUnbindTargetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PostPayPriceInfo struct {
	// 折扣，如20.0代表2折。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Discount *float64 `json:"Discount,omitnil,omitempty" name:"Discount"`

	// 单价，单位：元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPrice *float64 `json:"UnitPrice,omitnil,omitempty" name:"UnitPrice"`

	// 折扣单价，单位:元。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UnitPriceDiscount *float64 `json:"UnitPriceDiscount,omitnil,omitempty" name:"UnitPriceDiscount"`
}

type Price struct {
	// 描述了实例价格，单位：元/小时。
	InstancePrice *PostPayPriceInfo `json:"InstancePrice,omitnil,omitempty" name:"InstancePrice"`

	// 描述了lcu价格，单位：元/个。
	LcuPrice *PostPayPriceInfo `json:"LcuPrice,omitnil,omitempty" name:"LcuPrice"`
}

type QuotaInfo struct {
	// 当前剩余可用量，计算方式为 Limit - Used。仅当请求参数 DisplayFields 包含 available 时返回有效值；未请求时不返回或为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Available *int64 `json:"Available,omitnil,omitempty" name:"Available"`

	// 配额上限值。不同配额类型的单位不同，通常表示资源个数；超时时间类配额表示秒。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 配额类型，与请求参数 QuotaTypes 中的取值对应。每种配额类型的含义请参考 QuotaTypes 参数说明。
	QuotaType *string `json:"QuotaType,omitnil,omitempty" name:"QuotaType"`

	// 资源 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 当前已使用量。仅当请求参数 DisplayFields 包含 used 时返回有效值；未请求时不返回或为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Used *uint64 `json:"Used,omitnil,omitempty" name:"Used"`
}

type RelatedListener struct {
	// 监听器 ID，格式为 lst- 后接 8 位字母数字。
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 监听器端口。
	ListenerPort *int64 `json:"ListenerPort,omitnil,omitempty" name:"ListenerPort"`

	// 监听器协议。
	ListenerProtocol *string `json:"ListenerProtocol,omitnil,omitempty" name:"ListenerProtocol"`

	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`
}

type RemoveHTTPHeaderInfo struct {
	// 要删除的HTTP Header的键，长度1 ~ 40个字符，支持的字符集为：a-z A-Z 0-9 - _ 。
	// 不支持Cookie,Host,Content-Length,Connection,Upgrade,transfer-encoding,keep-alive,te,authority,x-forwarded-for,x-forwarded-proto,x-forwarded-host,x-forwarded-port,server
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`
}

// Predefined struct for user
type RemoveTargetsFromTargetGroupRequestParams struct {
	// 目标组 ID，格式为 lbtg- 后接 8 位字母数字。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 需要从目标组移除的后端服务列表。单次请求最多移除 **50** 个后端服务。
	Targets []*TargetToRemove `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 是否预览此次请求。 
	// - **false**（默认）：发送普通请求，直接移除后端服务。 
	// - **true**：发送预览请求，检查移除后端服务的参数、格式、业务限制等是否符合要求。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

type RemoveTargetsFromTargetGroupRequest struct {
	*tchttp.BaseRequest
	
	// 目标组 ID，格式为 lbtg- 后接 8 位字母数字。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 需要从目标组移除的后端服务列表。单次请求最多移除 **50** 个后端服务。
	Targets []*TargetToRemove `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 是否预览此次请求。 
	// - **false**（默认）：发送普通请求，直接移除后端服务。 
	// - **true**：发送预览请求，检查移除后端服务的参数、格式、业务限制等是否符合要求。
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`
}

func (r *RemoveTargetsFromTargetGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveTargetsFromTargetGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TargetGroupId")
	delete(f, "Targets")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveTargetsFromTargetGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveTargetsFromTargetGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveTargetsFromTargetGroupResponse struct {
	*tchttp.BaseResponse
	Response *RemoveTargetsFromTargetGroupResponseParams `json:"Response"`
}

func (r *RemoveTargetsFromTargetGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveTargetsFromTargetGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleAction struct {
	// 转发动作执行顺序，不能重复，按从小到大顺序执行。取值范围：1 ~ 50000。
	Order *int64 `json:"Order,omitnil,omitempty" name:"Order"`

	// 转发动作类型。取值：
	// TargetGroup：转发至目标组。
	// Redirect：重定向。
	// FixedResponse：返回固定内容。
	// Rewrite：重写。
	// InsertHeader：写入HTTP Header。
	// RemoveHeader：删除HTTP Header。
	// 转发动作必选包含TargetGroup,Redirect,FixedResponse其中一个，并且执行顺序放在最后。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 固定响应内容配置。
	FixedResponseConfig *FixedResponseInfo `json:"FixedResponseConfig,omitnil,omitempty" name:"FixedResponseConfig"`

	// 插入HTTP Header配置。
	InsertHeaderConfig *InsertHTTPHeaderInfo `json:"InsertHeaderConfig,omitnil,omitempty" name:"InsertHeaderConfig"`

	// 重定向配置。除HttpCode外，其他配置不能都使用默认值。
	RedirectConfig *HTTPRedirectInfo `json:"RedirectConfig,omitnil,omitempty" name:"RedirectConfig"`

	// 删除HTTP Header配置。
	RemoveHeaderConfig *RemoveHTTPHeaderInfo `json:"RemoveHeaderConfig,omitnil,omitempty" name:"RemoveHeaderConfig"`

	// 重写配置。
	RewriteConfig *HTTPRewriteInfo `json:"RewriteConfig,omitnil,omitempty" name:"RewriteConfig"`

	// 转发目标组配置。
	TargetGroupConfig *TargetGroupConfig `json:"TargetGroupConfig,omitnil,omitempty" name:"TargetGroupConfig"`
}

type RuleCondition struct {
	// 转发条件类型。取值：
	// Host：主机。
	// Path：路径。
	// Header：HTTP Header字段。
	// QueryString：HTPP查询字符串。
	// Method：请求方法。
	// Cookie：Cookie。
	// SourceIp：源 IP。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Cookie配置。
	CookieConfig []*HTTPCookieInfo `json:"CookieConfig,omitnil,omitempty" name:"CookieConfig"`

	// HTTP Header配置。
	HeaderConfig *HTTPHeaderInfo `json:"HeaderConfig,omitnil,omitempty" name:"HeaderConfig"`

	// 主机名。主机配置在一个规则中只能出现一次，长度3 ~ 128个字符，支持精确匹配，正则匹配，通配匹配。
	// 不能以半角句号（.）、下划线（_）开头或结尾。
	// 精确匹配，支持的字符集为：a-z 0-9 . - _ 。
	// 正则匹配，波浪线（~）开头为正则匹配，支持的字符集为：a-z 0-9 . - ? = ~ _ - + \ ^ * ! $ & | ( ) [ ] 。
	// 通配匹配，星号（*）多字符通配，半角问号（?）单字符通配，支持的字符集为：a-z 0-9 . - _ * ?。
	HostConfig []*string `json:"HostConfig,omitnil,omitempty" name:"HostConfig"`

	// 请求方法。取值：HEAD、GET、POST、OPTIONS、PUT、PATCH、DELETE。
	MethodConfig []*string `json:"MethodConfig,omitnil,omitempty" name:"MethodConfig"`

	// 转发路径。长度1 ~ 128个字符，支持精确匹配，正则匹配，通配匹配。
	// 精确匹配，支持的字符集为：a-z A-Z 0-9 . - _ / =  :。
	// 正则匹配，需以 ~ 开头，~ 开头表示区分大小写， ~* 开头表示不区分大小写，支持的字符集为： a-z A-Z 0-9 . - _ / = ?  ~ ^ * $ : ( ) [ ] + |。
	// 通配匹配，* 表示多字符通配，? 表示单字符通配，支持的字符集为：a-z A-Z 0-9 . - _ / =  :。
	PathConfig []*string `json:"PathConfig,omitnil,omitempty" name:"PathConfig"`

	// 查询字符串配置。
	QueryStringConfig []*HTTPQueryStringInfo `json:"QueryStringConfig,omitnil,omitempty" name:"QueryStringConfig"`

	// 源IP匹配配置。CIDR格式，IP地址x.x.x.x/32，IP网段x.x.x.x/24。
	SourceIpConfig []*string `json:"SourceIpConfig,omitnil,omitempty" name:"SourceIpConfig"`
}

type RuleHealthStatusInfo struct {
	// 是否为默认转发规则。
	IsDefaultRule *string `json:"IsDefaultRule,omitnil,omitempty" name:"IsDefaultRule"`

	// 转发规则 ID，格式为 rule- 后接 8 位字母数字。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 目标组健康状态。
	TargetGroupHealthInfos []*TargetGroupHealthInfo `json:"TargetGroupHealthInfos,omitnil,omitempty" name:"TargetGroupHealthInfos"`
}

type RuleInput struct {
	// 转发规则动作列表。
	Actions []*RuleAction `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 转发规则条件列表。
	Conditions []*RuleCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// 优先级，值越小优先级越高，不能重复，取值范围：1~10000。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 转发规则的方向。Request：客户端到负载均衡的请求方向，Response：后端服务器到负载均衡的响应方向。默认Request。
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 转发规则名称。长度为 1~255 个字符，可包含数字、大小写字母、中文、半角句号（.）、下划线（_）和短划线（-）。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 标签。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type RuleModify struct {
	// 转发规则动作列表。
	Actions []*RuleAction `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 转发规则条件列表。
	Conditions []*RuleCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// 优先级，值越小优先级越高，取值范围：1~10000。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 转发规则 ID，格式为 rule- 后接 8 位字母数字。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 转发规则名称。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

type RuleOutput struct {
	// 转发规则动作列表。	
	Actions []*RuleAction `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 转发规则条件列表。
	Conditions []*RuleCondition `json:"Conditions,omitnil,omitempty" name:"Conditions"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 转发规则的方向。Request：客户端到负载均衡的请求方向，Response：后端服务器到负载均衡的响应方向。
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 最后修改时间。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 优先级，值越小优先级越高，取值范围：1~10000。
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 转发规则 ID，格式为 rule- 后接 8 位字母数字。
	RuleId *string `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 转发规则名称。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 转发规则状态，Provisioning：创建中，Active：运行中，Configuring：配置中。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 标签列表。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type SecurityPolicyCapability struct {
	// 支持的加密套件列表。
	Ciphers []*string `json:"Ciphers,omitnil,omitempty" name:"Ciphers"`

	// 支持的 TLS 协议版本。可选值包括：TLSv1.0、TLSv1.1、TLSv1.2、TLSv1.3。
	TLSVersion *string `json:"TLSVersion,omitnil,omitempty" name:"TLSVersion"`
}

type SecurityPolicyInfo struct {
	// 支持的加密套件列表。
	// 支持的加密套件，具体依赖 TLSVersions 的取值。
	// Cipher 只要被任何一个传入的 TLSVersions 支持即可。
	// 
	// 说明：若选择了 TLSv1.3，那么 Cipher 列表必须包含 TLSv1.3 支持的 Cipher。
	// 
	// 请调用 DescribeSecurityPolicyCapabilities 接口获取支持的加密套件列表。
	Ciphers []*string `json:"Ciphers,omitnil,omitempty" name:"Ciphers"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 安全策略 ID，格式为 tls- 后接 8 位字母数字。
	SecurityPolicyId *string `json:"SecurityPolicyId,omitnil,omitempty" name:"SecurityPolicyId"`

	// 安全策略名称。长度为2~128个英文或中文字符，必须以字母或中文开头，可包含数字、半角句号（.）、下划线（_）和短划线（-）。
	SecurityPolicyName *string `json:"SecurityPolicyName,omitnil,omitempty" name:"SecurityPolicyName"`

	// 安全策略状态。当前接口最常返回 Active，表示安全策略处于可用状态。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 支持的 TLS 协议版本列表。可选值包括：TLSv1.0、TLSv1.1、TLSv1.2、TLSv1.3。
	TLSVersions []*string `json:"TLSVersions,omitnil,omitempty" name:"TLSVersions"`

	// 标签信息。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type SecurityPolicyRelations struct {
	// 安全策略与监听的关联关系列表。
	RelatedListeners []*RelatedListener `json:"RelatedListeners,omitnil,omitempty" name:"RelatedListeners"`

	// 安全策略 ID，格式为 tls- 后接 8 位字母数字。
	SecurityPolicyId *string `json:"SecurityPolicyId,omitnil,omitempty" name:"SecurityPolicyId"`
}

// Predefined struct for user
type SetLoadBalancerSecurityGroupsRequestParams struct {
	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 安全组 ID 列表。
	SecurityGroups []*string `json:"SecurityGroups,omitnil,omitempty" name:"SecurityGroups"`
}

type SetLoadBalancerSecurityGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 负载均衡实例 ID，格式为 alb- 后接 8 位字母数字。
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 安全组 ID 列表。
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

type StickySessionConfig struct {
	// 是否开启会话保持。
	// - **true**：开启。
	// - **false**：不开启。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StickySessionEnabled *bool `json:"StickySessionEnabled,omitnil,omitempty" name:"StickySessionEnabled"`

	// 自定义 Cookie 名称。
	// 长度为 1~255 个字符，只能包含英文字母和数字字符，且不能为`tgw_l7_tg_route`，该字段为目标组间会话保持Cookie保留字段。
	// >仅当 **StickySessionEnabled** 为 **true** 时该参数生效。
	Cookie *string `json:"Cookie,omitnil,omitempty" name:"Cookie"`

	// 会话保持时间。
	// 取值范围：**1-86400**，单位：**秒**。
	// 默认值：**1000**。
	// >仅当 **StickySessionEnabled** 为 **true**时，该参数生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CookieTimeout *int64 `json:"CookieTimeout,omitnil,omitempty" name:"CookieTimeout"`

	// 会话保持类型（处理Cookie的方式）。
	// - **Insert**（默认值）：植入 Cookie。 客户端第一次访问后端服务时，应用型负载均衡会在返回请求中植入 Cookie。下次客户端请求时携带该 Cookie，负载均衡将请求转发到上次请求的相同后端服务。
	// - **Rewrite**：重写 Cookie。 负载均衡会对用户自定义的 Cookie 进行重写，下次客户端请求时携带该 Cookie，负载均衡将请求转发到上次请求的相同后端服务。
	// >仅当 **StickySessionEnabled** 为 **true** 时该参数生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StickySessionType *string `json:"StickySessionType,omitnil,omitempty" name:"StickySessionType"`
}

type TagInfo struct {
	// 标签的键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签的值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type TargetGroupConfig struct {
	// 目标组列表。
	TargetGroups []*TargetGroupTuple `json:"TargetGroups,omitnil,omitempty" name:"TargetGroups"`

	// 目标组间会话保持
	TargetGroupStickySession *TargetGroupStickySession `json:"TargetGroupStickySession,omitnil,omitempty" name:"TargetGroupStickySession"`
}

type TargetGroupHealthInfo struct {
	// 是否开启健康检查。
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitnil,omitempty" name:"HealthCheckEnabled"`

	// 目标组 ID，格式为 lbtg- 后接 8 位字母数字。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 服务健康检查状态列表。
	TargetHealthStatusInfos []*TargetHealthStatusInfo `json:"TargetHealthStatusInfos,omitnil,omitempty" name:"TargetHealthStatusInfos"`

	// 转发动作类型。取值：
	// TargetGroup：转发至目标组。
	// Redirect：重定向。
	// FixedResponse：返回固定内容。
	// Rewrite：重写。
	// InsertHeader：写入HTTP Header。
	// RemoveHeader：删除HTTP Header。
	// 转发动作必选包含TargetGroup,Redirect,FixedResponse其中一个，并且执行顺序放在最后。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type TargetGroupOutput struct {
	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 健康检查配置。
	HealthCheckConfig *HealthCheckConfig `json:"HealthCheckConfig,omitnil,omitempty" name:"HealthCheckConfig"`

	// 是否开启长连接。
	KeepaliveEnabled *bool `json:"KeepaliveEnabled,omitnil,omitempty" name:"KeepaliveEnabled"`

	// 后端服务协议类型。取值：
	// - **HTTP**（默认）：支持绑定HTTP、HTTPS的监听器
	// - **HTTPS**：支持绑定HTTPS类型的监听器
	// - **GRPC**：支持绑定HTTPS类型的监听器
	// - **GRPCS**：支持绑定HTTPS类型的监听器
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 目标组关联的负载均衡数量。
	RelatedLoadBalancersCount *int64 `json:"RelatedLoadBalancersCount,omitnil,omitempty" name:"RelatedLoadBalancersCount"`

	// 调度算法。
	SchedulerAlgorithm *string `json:"SchedulerAlgorithm,omitnil,omitempty" name:"SchedulerAlgorithm"`

	// 会话保持配置。
	StickySessionConfig *StickySessionConfig `json:"StickySessionConfig,omitnil,omitempty" name:"StickySessionConfig"`

	// 标签。
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 目标组 ID，格式为 lbtg- 后接 8 位字母数字。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 目标组名称。默认为目标组ID。长度为 **1-255** 个字符，可包含数字、大小写字母、中文、半角句号（.）、下划线（_）和短划线（-）。
	TargetGroupName *string `json:"TargetGroupName,omitnil,omitempty" name:"TargetGroupName"`

	// 目标组状态。取值：
	// - **Provisioning**：创建中。
	// - **ProvisionFailed**：创建失败。
	// - **Active**: 运行中。
	// - **Configuring**：配置变更中。
	TargetGroupStatus *string `json:"TargetGroupStatus,omitnil,omitempty" name:"TargetGroupStatus"`

	// 目标组类型。取值：
	// - **Instance**：Cvm服务器类型或Eni弹性网卡类型
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 私有网络 ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type TargetGroupStickySession struct {
	// 是否开启会话保持，默认关闭。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 超时时间，单位秒，取值范围：1~86400，默认值：1000。
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`
}

type TargetGroupTuple struct {
	// 目标组 ID，格式为 lbtg- 后接 8 位字母数字。
	TargetGroupId *string `json:"TargetGroupId,omitnil,omitempty" name:"TargetGroupId"`

	// 权重，取值范围：[0, 100]，默认为 10。
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type TargetHealthStatusInfo struct {
	// 后端服务健康状态。DescribeListenerHealthStatus 仅返回非健康后端时，该值为 UnHealthy。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 后端服务实例 ID，CVM 实例格式为 ins- 后接 8 位字母数字。
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 后端目标服务IP。
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// 后端服务器端口。
	TargetPort *uint64 `json:"TargetPort,omitnil,omitempty" name:"TargetPort"`
}

type TargetOutput struct {
	// 弹性网卡 ID。
	EniId *string `json:"EniId,omitnil,omitempty" name:"EniId"`

	// 后端服务器使用的端口。取值范围：**1 - 65535**。
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 后端服务实例 ID，CVM 实例格式为 ins- 后接 8 位字母数字。
	TargetId *string `json:"TargetId,omitnil,omitempty" name:"TargetId"`

	// 后端服务IP。**TargetIp**和**TargetId**需要至少传一个。
	// 
	// - 当服务器组为 **Instance** 类型时，该参数为 **Eni** 的主内网 IP 或辅助内网 IP。
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// 后端服务名称。当前只有CVM后端类型后端服务返回有效名称。
	TargetName *string `json:"TargetName,omitnil,omitempty" name:"TargetName"`

	// 后端服务的状态。取值：
	// - **Adding**：添加中。
	// - **Active**：正常可用状态。
	// - **Configuring**：配置中。
	// - **Removing**：移除中。
	TargetStatus *string `json:"TargetStatus,omitnil,omitempty" name:"TargetStatus"`

	// 后端服务类型。
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 后端服务的权重，取值范围：**0 - 100**。默认值为**100**。如果设置权重为**0**，则不会将请求转发给该后端服务。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type TargetToAdd struct {
	// 后端服务器使用的端口。取值范围：**1 - 65535**。
	// 
	// >当目标组的 **targetType** 取值为 **Instance** 时，该参数必传。
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 后端服务IP。**TargetIp**和**TargetId**需要至少传一个。
	// 
	// - 当服务器组为 **Instance** 类型时，该参数为 **Eni** 的主内网 IP 或辅助内网 IP。
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// 后端服务的权重，取值范围：**0 - 100**。默认值为**10**。如果设置权重为**0**，则不会将请求转发给该后端服务。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type TargetToModify struct {
	// 后端服务IP。**TargetIp**和**TargetId**需要至少传一个。
	// 
	// - 当服务器组为 **Instance** 类型时，该参数为 **Eni** 的主内网 IP 或辅助内网 IP。
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`

	// 后端服务器使用的端口。取值范围：**1 - 65535**。
	// 
	// >当目标组的 **targetType** 取值为 **Instance** 时，该参数必传。
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 后端服务的权重，取值范围：**0 - 100**。如果设置权重为**0**，则不会将请求转发给该后端服务。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

type TargetToRemove struct {
	// 后端服务器使用的端口。取值范围：**1 - 65535**。
	// 
	// >当目标组的 **targetType** 取值为 **Instance** 时，该参数必传。
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 后端服务IP。**TargetIp**和**TargetId**需要至少传一个。
	// 
	// - 当服务器组为 **Instance** 类型时，该参数为 **Eni** 的主内网 IP 或辅助内网 IP。
	TargetIp *string `json:"TargetIp,omitnil,omitempty" name:"TargetIp"`
}

type XForwardedForConfig struct {
	// 是否通过 ALB-ID 头字段获取负载均衡实例ID。
	// - **true**：是。
	// - **false**：否。
	XForwardedForAlbIdEnabled *bool `json:"XForwardedForAlbIdEnabled,omitnil,omitempty" name:"XForwardedForAlbIdEnabled"`

	// 是否通过X-Forwarded-Client-srcport头字段获取访问负载均衡实例客户端的端口。
	// - **true**：是。
	// - **false**：否。
	XForwardedForClientSrcPortEnabled *bool `json:"XForwardedForClientSrcPortEnabled,omitnil,omitempty" name:"XForwardedForClientSrcPortEnabled"`

	// 是否开启通过X-Forwarded-Host头字段获取访问负载均衡实例客户端的域名。
	// - **true**：是。
	// - **false**：否。
	XForwardedForHostEnabled *bool `json:"XForwardedForHostEnabled,omitnil,omitempty" name:"XForwardedForHostEnabled"`

	// 指定如何处理 X-Forwarded-For（XFF）HTTP 头字段。
	// - **append**:  附加模式（默认），将客户端的真实 IP 地址附加到 X-Forwarded-For 头的末尾，保留原有的 XFF 链路信息
	// - **remove**:  删除模式，移除 X-Forwarded-For 头字段，不将该头传递给后端服务器
	// - **passthrough**: 透传模式，保持 X-Forwarded-For 头不变，直接透传给后端服务器，不做任何修改
	XForwardedForMode *string `json:"XForwardedForMode,omitnil,omitempty" name:"XForwardedForMode"`

	// 是否通过X-Forwarded-Port头字段获取负载均衡实例的监听端口。
	// - **true**：是。
	// - **false**：否。
	XForwardedForPortEnabled *bool `json:"XForwardedForPortEnabled,omitnil,omitempty" name:"XForwardedForPortEnabled"`

	// 是否通过X-Forwarded-Proto头字段获取负载均衡实例的监听协议。
	// - **true**：是。
	// - **false**：否。
	XForwardedForProtoEnabled *bool `json:"XForwardedForProtoEnabled,omitnil,omitempty" name:"XForwardedForProtoEnabled"`

	// 是否通过 X-Tencent-Client-IDN 头访问 客户端证书的颁发者 $ssl_client_i_dn。
	// - **true**：是。
	// - **false**：否。
	XTencentClientIDNEnabled *bool `json:"XTencentClientIDNEnabled,omitnil,omitempty" name:"XTencentClientIDNEnabled"`

	// 是否通过 X-Tencent-Client-SDN 头访问客户端证书的主题$ssl_client_s_dn。
	// - **true**：是。
	// - **false**：否。
	XTencentClientSDNEnabled *bool `json:"XTencentClientSDNEnabled,omitnil,omitempty" name:"XTencentClientSDNEnabled"`

	// 是否通过 X-Tencent-Client-Serial 头访问 客户端证书的序列号 $ssl_client_serial。
	// - **true**：是。
	// - **false**：否。
	XTencentClientSerialEnabled *bool `json:"XTencentClientSerialEnabled,omitnil,omitempty" name:"XTencentClientSerialEnabled"`

	// 是通过 X-Tencent-Client-Verify 头访问 客户端证书的验证结果 $ssl_client_verify。
	// - **true**：是。
	// - **false**：否。
	XTencentClientVerifyEnabled *bool `json:"XTencentClientVerifyEnabled,omitnil,omitempty" name:"XTencentClientVerifyEnabled"`
}

type Zone struct {
	// 可用区名称
	LocalName *string `json:"LocalName,omitnil,omitempty" name:"LocalName"`

	// 可用区 ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 可用区状态
	ZoneStatus *string `json:"ZoneStatus,omitnil,omitempty" name:"ZoneStatus"`
}

type ZoneMappingInfo struct {
	// <p>子网 ID。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>可用区ID。最多支持添加10个可用区。若当前地域支持2个及以上的可用区，至少需要添加2个可用区。<br>您可以通过调用<a href="https://cloud.tencent.com/document/api/1822/133727">DescribeZones</a>接口获取可用区ID对应的可用区的信息。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>负载均衡 VIP/EIP 信息</p>
	LoadBalancerAddress *LoadBalancerAddress `json:"LoadBalancerAddress,omitnil,omitempty" name:"LoadBalancerAddress"`

	// <p>可用区状态。取值：</p><ul><li><strong>Active</strong>：运行中。</li><li><strong>Stopped</strong>：已停止。</li><li><strong>Shifted</strong>：已移除。</li><li><strong>Starting</strong>：启动中。</li><li><strong>Stopping</strong>：停止中。</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ZoneMappingsItem struct {
	// <p>子网 ID。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// <p>可用区ID。最多支持添加10个可用区。若当前地域支持2个及以上的可用区，至少需要添加2个可用区。<br>您可以通过调用<a href="https://cloud.tencent.com/document/api/1822/133727">DescribeZones</a>接口获取可用区ID对应的可用区的信息。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// <p>公网实例绑定的EIP实例ID。</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerAddress *LoadBalancerAddress `json:"LoadBalancerAddress,omitnil,omitempty" name:"LoadBalancerAddress"`
}